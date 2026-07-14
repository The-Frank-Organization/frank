package observe

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackli/frank/internal/record"
)

func TestDeclaredNoneOpaqueAcceptFiresHonestlyLabeled(t *testing.T) {
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": t.TempDir()}, LaneVCS: map[string]string{"repo": "none"}})
	for _, tc := range []struct {
		authority string
		terminal  string
	}{
		{authority: "no", terminal: record.Accepted},
		{authority: "yes", terminal: record.Held},
	} {
		cand := record.Record{Headers: map[string]string{"authority_class": tc.authority, "FINAL_GIT_STATUS_SHORT": "none - clean tree"}}
		result, terminal := Gate(cand, "seat-a", "SITREP", "report-only", Env{
			PresentLayers: map[string]bool{"observe": true}, Evaluate: reg.EvaluateCandidateClaims,
		})
		if terminal != tc.terminal || result.PredicateResult != Degraded || result.ObservedFields["achieved_evidence"] != "E0" || result.ObservedFields["record_integrity"] != "self_reported" || result.ObservedFields["degradation_notes"] != "opaque-lane-no-vantage" {
			t.Fatalf("authority %s terminal=%q result=%#v", tc.authority, terminal, result)
		}
	}
}

func TestUndeclaredOrGitNeverOpaqueAccepts(t *testing.T) {
	git := fakeClaimlessGit(t, "")
	for _, vcs := range []map[string]string{nil, {"repo": "git"}} {
		reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": t.TempDir()}, LaneVCS: vcs, GitExecutable: git})
		got := reg.EvaluateCandidateClaims(Candidate{Phase: "SITREP", Authority: "report-only", Record: record.Record{Headers: map[string]string{"FINAL_GIT_STATUS_SHORT": "none - clean tree"}}})
		if got.Predicate != Degraded || got.FailureClass != "turn-attribution-unavailable" || got.ObservedFields["degradation_notes"] == "opaque-lane-no-vantage" {
			t.Fatalf("LaneVCS=%#v result=%#v", vcs, got)
		}
	}
}

func TestUngovernedRootIsConfigFaultNotOpaqueLane(t *testing.T) {
	reg := NewRegistry(RegistryEnv{LaneVCS: map[string]string{"repo": "none"}})
	got := reg.EvaluateCandidateClaims(Candidate{Phase: "SITREP", Authority: "report-only"})
	if !got.MachineryFault || got.Predicate != Blocked || got.FailureClass != "lane-ungoverned" {
		t.Fatalf("result=%#v", got)
	}
}

func TestRegistryEnvLaneVCSIsClonedNonAliasing(t *testing.T) {
	vcs := map[string]string{"repo": "none"}
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": t.TempDir()}, LaneVCS: vcs})
	vcs["repo"] = "git"
	if reg.env.LaneVCS["repo"] != "none" {
		t.Fatalf("registry alias retained: %#v", reg.env.LaneVCS)
	}
	if nilReg := NewRegistry(RegistryEnv{}); nilReg.env.LaneVCS != nil {
		t.Fatalf("nil LaneVCS became %#v", nilReg.env.LaneVCS)
	}
}

func TestAbsenceFloorPerformsNoSerializedFSSyscall(t *testing.T) {
	blocked := make(chan struct{})
	reg := NewRegistry(RegistryEnv{
		Lanes: map[string]string{"repo": t.TempDir()}, LaneVCS: map[string]string{"repo": "git"}, ReadTimeout: 10 * time.Millisecond,
		FSStageHook: func(stage FSStage) {
			if stage == FSStageRootOpen {
				<-blocked
			}
		},
	})
	started := time.Now()
	got := reg.EvaluateCandidateClaims(Candidate{Phase: "SITREP", Authority: "report-only"})
	close(blocked)
	if time.Since(started) > 250*time.Millisecond || !got.MachineryFault || got.Predicate != Blocked || got.FailureClass != "check-machinery-root-observability-timeout" {
		t.Fatalf("elapsed=%s result=%#v", time.Since(started), got)
	}
}

func fakeClaimlessGit(t *testing.T, output string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "git")
	if err := os.WriteFile(path, []byte("#!/bin/sh\nprintf '%s' '"+output+"'\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	return path
}
