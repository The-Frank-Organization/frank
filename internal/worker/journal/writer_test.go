package journal

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"testing"
)

func TestWriterFreshCreateResumeAndCrossGenerationHandoff(t *testing.T) {
	runtimeDir := privateRuntimeDir(t)
	commits := 0
	writer, err := Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunFresh,
		Identity: expectedIdentity(), GenerationID: "generation-a", TurnEpoch: "1",
		OnGenesisDurable: func(GenesisCommit) error { commits++; return nil },
	})
	if err != nil {
		t.Fatalf("Open(fresh): %v", err)
	}
	if commits != 1 {
		t.Fatalf("fresh genesis commits = %d, want 1", commits)
	}
	roundCommit, err := writer.AppendRound("turn-1", "0", []Record{
		recordForKind(t, KindToolCall),
		recordForKind(t, KindToolResult),
	})
	if err != nil {
		t.Fatalf("AppendRound: %v", err)
	}
	if err := RequireToolOutcome(roundCommit, "call-1"); err != nil {
		t.Fatalf("RequireToolOutcome: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Close first generation: %v", err)
	}

	writer, err = Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunResume,
		Identity: expectedIdentity(), GenerationID: "generation-b", TurnEpoch: "2",
		OnGenesisDurable: func(GenesisCommit) error { commits++; return nil },
	})
	if err != nil {
		t.Fatalf("Open(resume): %v", err)
	}
	if commits != 2 {
		t.Fatalf("resume genesis commits = %d, want total 2", commits)
	}
	if _, err := writer.Append(recordForKind(t, KindTurnScope)); err != nil {
		t.Fatalf("Append replacement generation: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Close replacement: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(runtimeDir, SessionLogName))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	result := Recover(data, expectedIdentity())
	if result.GenesisFault {
		t.Fatalf("Recover after handoff: %+v", result)
	}
	last, err := DecodeRecord(result.Records[len(result.Records)-1])
	if err != nil {
		t.Fatalf("Decode last record: %v", err)
	}
	if last.GenerationID != "generation-b" {
		t.Fatalf("last generation = %q, want generation-b", last.GenerationID)
	}
}

func TestSecondWriterFenceViolation(t *testing.T) {
	runtimeDir := privateRuntimeDir(t)
	first := openFreshWriter(t, runtimeDir)
	defer first.Close()
	second, err := Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunResume,
		Identity: expectedIdentity(), GenerationID: "generation-b", TurnEpoch: "2",
	})
	if second != nil {
		second.Close()
	}
	if !errors.Is(err, ErrFenceHeld) {
		t.Fatalf("second Open error = %v, want ErrFenceHeld", err)
	}
}

func TestResumeAbsentNeverRecreates(t *testing.T) {
	runtimeDir := privateRuntimeDir(t)
	writer, err := Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunResume,
		Identity: expectedIdentity(), GenerationID: "generation-b", TurnEpoch: "2",
	})
	if writer != nil {
		writer.Close()
	}
	if !errors.Is(err, ErrResumeLogAbsent) {
		t.Fatalf("Open error = %v, want ErrResumeLogAbsent", err)
	}
	if _, statErr := os.Lstat(filepath.Join(runtimeDir, SessionLogName)); !os.IsNotExist(statErr) {
		t.Fatalf("resume absent created a path: %v", statErr)
	}
}

func TestDuplicateFreshPlusJournalLossCannotConsumeCreateTwice(t *testing.T) {
	var intake IntakeHighWater
	admitted, err := intake.Admit("0")
	if err != nil || !admitted {
		t.Fatalf("first fresh admission = %v, %v", admitted, err)
	}
	runtimeDir := privateRuntimeDir(t)
	writer := openFreshWriter(t, runtimeDir)
	if err := writer.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if err := os.Remove(filepath.Join(runtimeDir, SessionLogName)); err != nil {
		t.Fatalf("Remove journal: %v", err)
	}
	if admitted, err := intake.Admit("0"); err != nil || admitted {
		t.Fatalf("duplicate fresh admission = %v, %v", admitted, err)
	}
	if _, err := os.Lstat(filepath.Join(runtimeDir, SessionLogName)); !os.IsNotExist(err) {
		t.Fatalf("duplicate admission recreated journal: %v", err)
	}

	var replacementIncarnation IntakeHighWater
	if admitted, err := replacementIncarnation.Admit("0"); err != nil || !admitted {
		t.Fatalf("replacement resume admission = %v, %v", admitted, err)
	}
	resumed, err := Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunResume,
		Identity: expectedIdentity(), GenerationID: "generation-b", TurnEpoch: "2",
	})
	if resumed != nil {
		resumed.Close()
	}
	if !errors.Is(err, ErrResumeLogAbsent) {
		t.Fatalf("replacement Open error = %v, want ErrResumeLogAbsent", err)
	}
}

func TestResumeTruncatesOnlyUntrustedTail(t *testing.T) {
	runtimeDir := privateRuntimeDir(t)
	writer := openFreshWriter(t, runtimeDir)
	if err := writer.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	logPath := filepath.Join(runtimeDir, SessionLogName)
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		t.Fatalf("open tail: %v", err)
	}
	if _, err := file.Write([]byte(`{"seq":`)); err != nil {
		t.Fatalf("write torn tail: %v", err)
	}
	if err := file.Sync(); err != nil {
		t.Fatalf("sync torn tail: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("close tail: %v", err)
	}

	writer, err = Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunResume,
		Identity: expectedIdentity(), GenerationID: "generation-b", TurnEpoch: "2",
	})
	if err != nil {
		t.Fatalf("Open resume: %v", err)
	}
	defer writer.Close()
	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !strings.HasSuffix(string(data), "\n") || strings.Contains(string(data), `{"seq":`) {
		t.Fatalf("untrusted torn tail was not removed: %q", data)
	}
	result := Recover(data, expectedIdentity())
	if result.Boundary.Kind != BoundaryGenesis || result.NextSeq != 1 {
		t.Fatalf("Recover after truncate = %+v", result)
	}
}

func TestWriterRejectsWrongIdentityAndCreateAuthorization(t *testing.T) {
	runtimeDir := privateRuntimeDir(t)
	writer := openFreshWriter(t, runtimeDir)
	if err := writer.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	for _, testCase := range []struct {
		name     string
		identity Identity
	}{
		{name: "manifest", identity: Identity{RunID: "run-1", RunManifestDigest: strings.Repeat("f", 64), CreateAuthID: strings.Repeat("b", 32)}},
		{name: "create auth", identity: Identity{RunID: "run-1", RunManifestDigest: strings.Repeat("a", 64), CreateAuthID: strings.Repeat("c", 32)}},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			opened, err := Open(Config{
				RuntimeDir: runtimeDir, RunDisposition: RunResume,
				Identity: testCase.identity, GenerationID: "generation-b", TurnEpoch: "2",
			})
			if opened != nil {
				opened.Close()
			}
			if !errors.Is(err, ErrGenesisIdentity) {
				t.Fatalf("Open error = %v, want ErrGenesisIdentity", err)
			}
		})
	}
}

func TestDescriptorBatteryRejectsSymlinkModeHardlinkAndReplacement(t *testing.T) {
	t.Run("symlinked parent", func(t *testing.T) {
		realDir := privateRuntimeDir(t)
		link := filepath.Join(t.TempDir(), "runtime-link")
		if err := os.Symlink(realDir, link); err != nil {
			t.Fatalf("Symlink: %v", err)
		}
		writer, err := Open(Config{RuntimeDir: link, RunDisposition: RunFresh, Identity: expectedIdentity(), GenerationID: "generation-a", TurnEpoch: "1"})
		if writer != nil {
			writer.Close()
		}
		if !errors.Is(err, ErrDescriptorSafety) {
			t.Fatalf("Open error = %v, want ErrDescriptorSafety", err)
		}
	})

	t.Run("parent mode", func(t *testing.T) {
		runtimeDir := privateRuntimeDir(t)
		if err := os.Chmod(runtimeDir, 0o750); err != nil {
			t.Fatalf("Chmod parent: %v", err)
		}
		writer, err := Open(Config{RuntimeDir: runtimeDir, RunDisposition: RunFresh, Identity: expectedIdentity(), GenerationID: "generation-a", TurnEpoch: "1"})
		if writer != nil {
			writer.Close()
		}
		if !errors.Is(err, ErrDescriptorSafety) {
			t.Fatalf("Open error = %v, want ErrDescriptorSafety", err)
		}
	})

	t.Run("symlinked journal", func(t *testing.T) {
		runtimeDir := privateRuntimeDir(t)
		target := filepath.Join(runtimeDir, "target.log")
		if err := os.WriteFile(target, []byte("foreign\n"), 0o600); err != nil {
			t.Fatalf("WriteFile target: %v", err)
		}
		if err := os.Symlink(target, filepath.Join(runtimeDir, SessionLogName)); err != nil {
			t.Fatalf("Symlink journal: %v", err)
		}
		writer, err := Open(Config{RuntimeDir: runtimeDir, RunDisposition: RunResume, Identity: expectedIdentity(), GenerationID: "generation-a", TurnEpoch: "1"})
		if writer != nil {
			writer.Close()
		}
		if !errors.Is(err, ErrDescriptorSafety) {
			t.Fatalf("Open error = %v, want ErrDescriptorSafety", err)
		}
	})

	t.Run("mode drift", func(t *testing.T) {
		runtimeDir := privateRuntimeDir(t)
		writer := openFreshWriter(t, runtimeDir)
		defer writer.Close()
		if err := os.Chmod(filepath.Join(runtimeDir, SessionLogName), 0o644); err != nil {
			t.Fatalf("Chmod: %v", err)
		}
		if _, err := writer.Append(recordForKind(t, KindTurnScope)); !errors.Is(err, ErrDescriptorSafety) {
			t.Fatalf("Append error = %v, want ErrDescriptorSafety", err)
		}
	})

	t.Run("hard link", func(t *testing.T) {
		runtimeDir := privateRuntimeDir(t)
		writer := openFreshWriter(t, runtimeDir)
		if err := writer.Close(); err != nil {
			t.Fatalf("Close: %v", err)
		}
		if err := os.Link(filepath.Join(runtimeDir, SessionLogName), filepath.Join(runtimeDir, "other-link")); err != nil {
			t.Fatalf("Link: %v", err)
		}
		opened, err := Open(Config{RuntimeDir: runtimeDir, RunDisposition: RunResume, Identity: expectedIdentity(), GenerationID: "generation-b", TurnEpoch: "2"})
		if opened != nil {
			opened.Close()
		}
		if !errors.Is(err, ErrDescriptorSafety) {
			t.Fatalf("Open error = %v, want ErrDescriptorSafety", err)
		}
	})

	t.Run("path replacement while descriptor open", func(t *testing.T) {
		runtimeDir := privateRuntimeDir(t)
		writer := openFreshWriter(t, runtimeDir)
		defer writer.Close()
		logPath := filepath.Join(runtimeDir, SessionLogName)
		replaced := filepath.Join(runtimeDir, "replaced.log")
		if err := os.Rename(logPath, replaced); err != nil {
			t.Fatalf("Rename: %v", err)
		}
		data, err := os.ReadFile(replaced)
		if err != nil {
			t.Fatalf("ReadFile: %v", err)
		}
		if err := os.WriteFile(logPath, data, 0o600); err != nil {
			t.Fatalf("WriteFile replacement: %v", err)
		}
		if _, err := writer.Append(recordForKind(t, KindTurnScope)); !errors.Is(err, ErrDescriptorSafety) {
			t.Fatalf("Append error = %v, want ErrDescriptorSafety", err)
		}
	})
}

func TestSyntheticWrongOwnerDescriptorRejected(t *testing.T) {
	stat := syscall.Stat_t{Uid: uint32(os.Geteuid() + 1), Mode: syscall.S_IFREG | 0o600, Nlink: 1}
	if err := validateJournalStat(&stat, uint32(os.Geteuid())); !errors.Is(err, ErrDescriptorSafety) {
		t.Fatalf("validateJournalStat error = %v, want ErrDescriptorSafety", err)
	}
}

func TestTurnOpenIntakeDedupIsIncarnationScoped(t *testing.T) {
	var oldIncarnation IntakeHighWater
	for _, testCase := range []struct {
		seq  string
		want bool
	}{
		{seq: "0", want: true},
		{seq: "0", want: false},
		{seq: "1", want: true},
		{seq: "1", want: false},
	} {
		got, err := oldIncarnation.Admit(testCase.seq)
		if err != nil || got != testCase.want {
			t.Fatalf("Admit(%q) = %v, %v, want %v", testCase.seq, got, err, testCase.want)
		}
	}
	if _, err := oldIncarnation.Admit("01"); err == nil {
		t.Fatal("Admit accepted non-canonical seq")
	}

	var freshIncarnation IntakeHighWater
	if admitted, err := freshIncarnation.Admit("0"); err != nil || !admitted {
		t.Fatalf("fresh incarnation first seq = %v, %v", admitted, err)
	}
}

func TestDurabilityOrderingRequiresMarkerBeforeOutcome(t *testing.T) {
	if err := RequireToolOutcome(RoundCommit{}, "call-1"); !errors.Is(err, ErrDurabilityOrder) {
		t.Fatalf("zero commit error = %v, want ErrDurabilityOrder", err)
	}
	runtimeDir := privateRuntimeDir(t)
	writer := openFreshWriter(t, runtimeDir)
	defer writer.Close()
	commit, err := writer.AppendRound("turn-1", "0", []Record{recordForKind(t, KindToolResult)})
	if err != nil {
		t.Fatalf("AppendRound: %v", err)
	}
	if err := RequireToolOutcome(commit, "call-1"); err != nil {
		t.Fatalf("RequireToolOutcome after marker fsync: %v", err)
	}
	if err := RequireToolOutcome(commit, "other-call"); !errors.Is(err, ErrDurabilityOrder) {
		t.Fatalf("wrong identity error = %v, want ErrDurabilityOrder", err)
	}
}

func privateRuntimeDir(t *testing.T) string {
	t.Helper()
	dir := filepath.Join(t.TempDir(), "run")
	if err := os.Mkdir(dir, 0o700); err != nil {
		t.Fatalf("Mkdir: %v", err)
	}
	if runtime.GOOS == "windows" {
		t.Skip("descriptor battery requires POSIX file descriptors")
	}
	return dir
}

func openFreshWriter(t *testing.T, runtimeDir string) *Writer {
	t.Helper()
	writer, err := Open(Config{
		RuntimeDir: runtimeDir, RunDisposition: RunFresh,
		Identity: expectedIdentity(), GenerationID: "generation-a", TurnEpoch: "1",
	})
	if err != nil {
		t.Fatalf("Open(fresh): %v", err)
	}
	return writer
}
