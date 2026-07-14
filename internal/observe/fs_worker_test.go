package observe

import (
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"
)

func TestFSWorkerDetachesEveryStage(t *testing.T) {
	stages := []FSStage{
		FSStageRootOpen,
		FSStageDirectoryRead,
		FSStageFileOpen,
		FSStageMetadata,
		FSStageRead,
		FSStageClose,
	}
	for _, stage := range stages {
		t.Run(string(stage), func(t *testing.T) {
			root := t.TempDir()
			if err := os.WriteFile(filepath.Join(root, "artifact.txt"), []byte("inside\n"), 0o600); err != nil {
				t.Fatal(err)
			}
			blocked := make(chan struct{})
			reg := NewRegistry(RegistryEnv{
				Lanes:       map[string]string{"repo": root},
				ReadTimeout: 10 * time.Millisecond,
				FSStageHook: func(got FSStage) {
					if got == stage {
						<-blocked
					}
				},
			})
			started := time.Now()
			var result fsResult
			if stage == FSStageDirectoryRead {
				result = reg.executeRootHealth(Selection{CheckID: "root-health", ClaimRef: "root"}, "repo")
			} else {
				result = reg.executeFS(Selection{CheckID: "read-file", ClaimRef: "artifact"}, "repo", func(rootFD int) fsResult {
					return readFileFSOp(rootFD, "artifact.txt", time.Now().Add(time.Second), reg.fsHook())
				})
			}
			if elapsed := time.Since(started); elapsed > time.Second {
				t.Fatalf("serialized path blocked for %v at %s", elapsed, stage)
			}
			if result.kind != fsResultMachinery || result.timing != "timeout" {
				t.Fatalf("result at %s = %#v, want typed timeout machinery fault", stage, result)
			}
			close(blocked)
		})
	}
}

func TestFSWorkerBreakerBoundsWorkers(t *testing.T) {
	root := t.TempDir()
	var launches atomic.Int32
	blocked := make(chan struct{})
	reg := NewRegistry(RegistryEnv{
		Lanes:       map[string]string{"repo": root, "other": root},
		ReadTimeout: 10 * time.Millisecond,
		FSStageHook: func(stage FSStage) {
			if stage == FSStageRootOpen && launches.Add(1) == 1 {
				<-blocked
			}
		},
	})
	first := reg.executeRootHealth(Selection{CheckID: "root-health", ClaimRef: "first"}, "repo")
	if first.kind != fsResultMachinery || first.timing != "timeout" {
		t.Fatalf("first = %#v", first)
	}
	second := reg.executeRootHealth(Selection{CheckID: "root-health", ClaimRef: "second"}, "repo")
	if second.kind != fsResultMachinery || second.detail != "check-machinery-fs-breaker-open" {
		t.Fatalf("second = %#v, want breaker-open refusal", second)
	}
	if got := launches.Load(); got != 1 {
		t.Fatalf("same-lane second call launched a worker: %d", got)
	}
	other := reg.executeRootHealth(Selection{CheckID: "root-health", ClaimRef: "other"}, "other")
	if other.kind != fsResultData {
		t.Fatalf("other lane = %#v, want independently usable", other)
	}
	close(blocked)
}

func TestFSWorkerComponentSwapNoOutsideBytes(t *testing.T) {
	root := t.TempDir()
	outside := t.TempDir()
	insideDir := filepath.Join(root, "component")
	if err := os.Mkdir(insideDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(insideDir, "artifact.txt"), []byte("inside"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(outside, "artifact.txt"), []byte("outside-secret"), 0o600); err != nil {
		t.Fatal(err)
	}
	var opens atomic.Int32
	reg := NewRegistry(RegistryEnv{
		Lanes: map[string]string{"repo": root},
		FSStageHook: func(stage FSStage) {
			if stage != FSStageFileOpen || opens.Add(1) != 2 {
				return
			}
			moved := insideDir + "-moved"
			if err := os.Rename(insideDir, moved); err != nil {
				t.Errorf("rename component: %v", err)
				return
			}
			if err := os.Symlink(outside, insideDir); err != nil {
				t.Errorf("replace component: %v", err)
			}
		},
	})
	result := reg.executeFS(Selection{CheckID: "read-file", ClaimRef: "swap"}, "repo", func(rootFD int) fsResult {
		return readFileFSOp(rootFD, filepath.Join("component", "artifact.txt"), time.Now().Add(time.Second), reg.fsHook())
	})
	if string(result.data) == "outside-secret" {
		t.Fatalf("descriptor-rooted traversal escaped: %#v", result)
	}
	if result.kind != fsResultData || string(result.data) != "inside" {
		t.Fatalf("component swap result = %#v, want original descriptor-rooted bytes", result)
	}
}

func TestReadFileRootFailureKeepsLandedMachineryClass(t *testing.T) {
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": filepath.Join(t.TempDir(), "missing")}})
	verdict := reg.Run(Selection{CheckID: "read-file", ClaimRef: "missing-root", Params: map[string]string{
		"lane_ref": "repo", "path": "artifact.txt", "expect": "line:inside",
	}})
	if verdict.FailingDetail != "check-machinery-read-file" {
		t.Fatalf("root failure detail = %q, want landed read-file machinery class", verdict.FailingDetail)
	}
}
