## SITREP — s10 T1 complete at commit 94f8ab1: v8 owner bytes, phase-0 refusal, faithful v8-to-v7 genesis inverse, and historical s8 fixture truing; full uncached battery GREEN; request task review plus m-2/m-7 Step-5 confirms

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s10-build-t1-review
PARENT_DISPATCH_ID: s10-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — task review and two owner confirms are required before T2; merge remains operator-only
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
TASK_ID: T1
IN_REPLY_TO: s10-build-t1-review/SITREP-planner-20260713-002732.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: one T1 commit is ready for review; owner bytes are verbatim, T1-g is exactly the leading v8-to-v7 inverse, T1-f restores historical fixture construction without changing the proven s8 behaviors, pinned v5 remains byte-identical, and the unedited recursive dogfood leg recovered with the full suite

Summary: T1 is implemented in the single commit `94f8ab1274a7a884da19dd1941ff274173e79d25` over base `8941889a432a0a02a4183a16fb0015c230c9561b`. The initial marker test was observed RED before owner bytes; the first full battery exposed seven sequence-honest failures; the planner-authorized T1-g/T1-f repair then made the complete uncached battery GREEN. No later task has started.

SCOPE_DIFF:
- frank/internal/fieldspec/registry.json -> in
- frank/internal/fieldspec/registry_test.go -> in
- frank/internal/config/config.go -> in
- frank/internal/store/genesis.go -> in
- frank/test/fixtures/s10_fieldspec_test.go -> in
- frank/test/fixtures/s8_registry_changeset_test.go -> in
- frank/test/fixtures/s8_exit_gate_test.go -> in
- frank/test/fixtures/s8_supply_test.go -> in
- frank/.relays/s10/t1-red-20260712-225520.txt -> in
- frank/.relays/s10/t1-full-battery-blocker-20260713-002343.txt -> in
- frank/.relays/s10/t1-green-20260713-003445.txt -> in
SCOPE_DIFF_RESULT: all-in

ACTIONS_GIT_REF: s10-comms-spine@94f8ab1274a7a884da19dd1941ff274173e79d25; base 8941889a432a0a02a4183a16fb0015c230c9561b; one T1 commit; worktree clean after commit

## Diff-to-license / task seam table

| Diff surface | Task seam | License | Bounded result |
|---|---|---|---|
| `internal/fieldspec/registry.json` marker + two record-kind tokens | T1 owner bytes | m-2 `COORD-planner-20260712-233000.md`; m-7 `SITREP-planner-20260712-230402.md` | v8 marker; append only `odb` and `resummon_command`; seat_scope and required_when untouched |
| `internal/config/config.go` reader ceiling | T1 owner bytes | m-7 return above | append v8 capability only |
| `internal/config/config.go` successor relation | T1 owner bytes | identical m-2 + m-7 byte | append only adjacent v7-to-v8 successor |
| `internal/fieldspec/registry_test.go` | T1 owner pin | m-2 return above | version assertion only |
| `test/fixtures/s10_fieldspec_test.go` | T1 RED/GREEN | locked plan + `SITREP-planner-20260712-233011.md` | v7 reader refuses v8 at phase 0; v8 reader accepts marker before invalid content is parsed |
| `internal/store/genesis.go` | T1-g, before future T2 seam | `SITREP-planner-20260713-002732.md` item 1 | one leading v8-to-v7 mechanical inverse; existing v7-to-v6-to-v5 block byte-untouched |
| three historical s8 fixture files | T1-f | `SITREP-planner-20260713-002732.md` item 2 | constructors now derive intended v7/v6 bytes from live v8; asserted behavior unchanged |
| three `.relays/s10` captures | T1 evidence | plan acceptance + ruling | RED, blocker, full GREEN records |

No T2 fresh-v8 genesis work, T3+, T9/T10, executor, main composition, merge, push, or PR change is present.

## Constructor / fixture semantic-invariance proof lines

- `s8FieldspecV7Bytes` — new shared historical constructor removes exactly `odb` and `resummon_command` and reverts only the marker; process-substitution `cmp` against `8941889:internal/fieldspec/registry.json` exits 0, proving byte identity with the prior live v7.
- `s8FieldspecV6Bytes` — now starts from the byte-identical historical v7 constructor, then retains its prior v7-to-v6 marker replacement and executable-claims-row removal unchanged; v5-to-v6 and v6-to-v7 transition assertions pass.
- `TestS8RegistryChangesetIsExactV7OptionB` — parses the historical v7 constructor instead of the now-v8 live file; its exact v7 member/scope/predicate assertions are unchanged and pass.
- `TestS8FXCFG10MemberTransitionRelation` — uses historical v7 bytes; lawful v5-to-v6/v6-to-v7 remains accepted and v5-to-v7 skip/v7-to-v6 rollback remain refused.
- `TestS8ExitGateFreshGenesisActivationAndDogfoodLegs` — uses historical v7 for its v6-to-v7 activation step; accepted/rejected evidence and projection outcomes remain unchanged.
- `TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate` — uses historical v7 for the same governed transition; all original false-done and named-predicate outcomes pass. The recursive dogfood script/test was not edited and recovered when the suite became green.
- `TestS8V6ReaderRefusesV7MarkerBeforeContent` — plants invalid content in historical v7; v6 reader still refuses at marker preflight and v7 reader still accepts the marker without interpreting content.
- `TestS8ExecutableClaimTypedRejects` — uses historical v7 before testing claims; all original typed rejection classes/outcomes pass.
- `TestS8FXSUP3MemberTransitionRefusesDuplicateMarkers/fieldspec` — duplicates the current v8 marker instead of the stale v7 literal; the proven behavior remains duplicate-marker transition refusal.

## Acceptance and evidence

- RED-first: `TestS10V7ReaderRefusesV8MarkerBeforeContent` failed on the old v7 marker; captured at `frank/.relays/s10/t1-red-20260712-225520.txt`.
- Owner-byte targeted/package GREEN was established before the full-battery blocker.
- Pinned-v5 tripwire: `TestS8FXCFG7GenesisComposesThreePinnedMembers` passes with the existing expected hash `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`; neither that constant nor the fixture expectation changed.
- Faithful inverse: `cmp <(git show 8941889:internal/fieldspec/registry.json) <(sed -e 's/s10-fieldspec-v8/s8-fieldspec-v7/' -e 's/, "odb", "resummon_command"//' internal/fieldspec/registry.json)` — exit 0.
- Exact former failures: the six direct failures from the blocker capture pass uncached.
- Full battery: `go test ./... -count=1` — exit 0; `test/fixtures` GREEN in 109.092s, including the unedited recursive dogfood leg; full output at `frank/.relays/s10/t1-green-20260713-003445.txt`.
- `git diff --check 8941889..94f8ab1` — exit 0.
- Final worktree: `git status --short --branch` prints only `## s10-comms-spine`.

Evidence level: E2 for RED/GREEN, targeted regression set, full battery, byte comparison, and diff check; E1 for the whole diff and seam mapping.

## Step-5 confirmations requested

- m-2.planner: confirm the landed marker/record-kind/successor/version-pin bytes and confirm T1-g is the faithful inverse of items 1+4 with the pinned-v5 tripwire intact.
- m-7.planner: confirm the v8 reader-capability move and the bounded `genesis.go` v8-to-v7 derivation arm on the owned store-genesis surface.

T4 remains hard-held behind both confirmations. T2+ remain sequence-blocked until s10.planner closes this T1 task review.

FINAL_GIT_STATUS_SHORT:
## s10-comms-spine

## Whole commit diff — 8941889..94f8ab1

```diff
commit 94f8ab1274a7a884da19dd1941ff274173e79d25
Author:     Jack Li <jack@Jacks-MacBook-Pro.local>
AuthorDate: Mon Jul 13 00:35:21 2026 -0700
Commit:     Jack Li <jack@Jacks-MacBook-Pro.local>
CommitDate: Mon Jul 13 00:35:21 2026 -0700

    feat: add s10 fieldspec v8 transition

diff --git a/.relays/s10/t1-full-battery-blocker-20260713-002343.txt b/.relays/s10/t1-full-battery-blocker-20260713-002343.txt
new file mode 100644
index 0000000..9b8f0a0
--- /dev/null
+++ b/.relays/s10/t1-full-battery-blocker-20260713-002343.txt
@@ -0,0 +1,27 @@
+$ go test ./... -count=1
+--- FAIL: TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation (0.05s)
+    s5_config_change_test.go:147: stale state/body = accepted/"", want form_digest:re-render
+--- FAIL: TestS8FXCFG7GenesisComposesThreePinnedMembers (0.04s)
+    s8_config_activation_test.go:45: fresh_v2_genesis_config_digest=f8310109a3636cd5fdfadb2a6b43e251f0a050edd8a7fc16700c63ea7048d1e2
+    s8_config_activation_test.go:51: fieldspec hash = f9a15c53871613015f6af8e8f937b47be51271ee9764eb4218d3be459a267599, want 1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485
+--- FAIL: TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate (41.32s)
+    s8_exit_gate_test.go:315: production suite state = "rejected", want accepted; nested dogfood-battery returned fail
+--- FAIL: TestS8V6ReaderRefusesV7MarkerBeforeContent (0.00s)
+    s8_exit_gate_test.go:403: v7 marker preflight interpreted planted content: config-load: fieldspec-marker
+--- FAIL: TestS8RegistryChangesetIsExactV7OptionB (0.00s)
+    s8_registry_changeset_test.go:21: version = "s10-fieldspec-v8", want s8-fieldspec-v7
+--- FAIL: TestS8FXCFG10MemberTransitionRelation (0.03s)
+    s8_registry_changeset_test.go:114: v5-to-v7 skip accepted
+--- FAIL: TestS8FXSUP3MemberTransitionRefusesDuplicateMarkers (0.00s)
+    --- FAIL: TestS8FXSUP3MemberTransitionRefusesDuplicateMarkers/fieldspec (0.00s)
+        s8_supply_test.go:199: fieldspec marker mutation did not apply
+FAIL
+FAIL	github.com/jackli/frank/test/fixtures	61.538s
+ok  	github.com/jackli/frank/test/invariants	3.778s
+ok  	github.com/jackli/frank/test/replay	1.841s
+ok  	github.com/jackli/frank/test/replay/dogfood	1.966s
+ok  	github.com/jackli/frank/test/replay/zeroloss	2.317s
+FAIL
+
+$ go test ./test/fixtures -run '^(TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation|TestS8FXCFG7GenesisComposesThreePinnedMembers|TestS8V6ReaderRefusesV7MarkerBeforeContent|TestS8RegistryChangesetIsExactV7OptionB|TestS8FXCFG10MemberTransitionRelation|TestS8FXSUP3MemberTransitionRefusesDuplicateMarkers)$' -count=1
+The same six direct failures reproduced; the seventh full-suite failure is the governed dogfood suite recursively observing those failures.
diff --git a/.relays/s10/t1-green-20260713-003445.txt b/.relays/s10/t1-green-20260713-003445.txt
new file mode 100644
index 0000000..7670553
--- /dev/null
+++ b/.relays/s10/t1-green-20260713-003445.txt
@@ -0,0 +1,37 @@
+$ go test ./... -count=1
+?   	github.com/jackli/frank/cmd/frank	[no test files]
+ok  	github.com/jackli/frank/cmd/frank-mcp	0.333s
+ok  	github.com/jackli/frank/internal/bounce	0.435s
+ok  	github.com/jackli/frank/internal/channel	0.943s
+ok  	github.com/jackli/frank/internal/config	0.824s
+ok  	github.com/jackli/frank/internal/crashpoint	1.035s
+ok  	github.com/jackli/frank/internal/egress	1.513s
+ok  	github.com/jackli/frank/internal/engine	2.724s
+?   	github.com/jackli/frank/internal/executor	[no test files]
+ok  	github.com/jackli/frank/internal/fieldspec	1.939s
+ok  	github.com/jackli/frank/internal/fsio	2.133s
+ok  	github.com/jackli/frank/internal/gate	2.990s
+ok  	github.com/jackli/frank/internal/gc	3.072s
+ok  	github.com/jackli/frank/internal/intake	5.803s
+ok  	github.com/jackli/frank/internal/lineage	2.732s
+ok  	github.com/jackli/frank/internal/migrate	2.507s
+ok  	github.com/jackli/frank/internal/obligation	2.912s
+?   	github.com/jackli/frank/internal/observe	[no test files]
+ok  	github.com/jackli/frank/internal/record	2.675s
+ok  	github.com/jackli/frank/internal/recover	3.880s
+ok  	github.com/jackli/frank/internal/seat	2.195s
+ok  	github.com/jackli/frank/internal/store	3.948s
+ok  	github.com/jackli/frank/internal/tables	1.909s
+ok  	github.com/jackli/frank/test/fixtures	109.092s
+ok  	github.com/jackli/frank/test/invariants	3.757s
+ok  	github.com/jackli/frank/test/replay	1.751s
+ok  	github.com/jackli/frank/test/replay/dogfood	1.840s
+ok  	github.com/jackli/frank/test/replay/zeroloss	2.361s
+?   	github.com/jackli/frank/test/seatproc	[no test files]
+exit status: 0
+
+$ cmp <(git show 8941889:internal/fieldspec/registry.json) <(sed -e 's/s10-fieldspec-v8/s8-fieldspec-v7/' -e 's/, "odb", "resummon_command"//' internal/fieldspec/registry.json)
+exit status: 0 — the v8 inverse matches 8941889 v7 byte-for-byte
+
+$ git diff --check
+exit status: 0
diff --git a/.relays/s10/t1-red-20260712-225520.txt b/.relays/s10/t1-red-20260712-225520.txt
new file mode 100644
index 0000000..53f2997
--- /dev/null
+++ b/.relays/s10/t1-red-20260712-225520.txt
@@ -0,0 +1,6 @@
+$ go test ./test/fixtures -run '^TestS10V7ReaderRefusesV8MarkerBeforeContent$' -count=1
+--- FAIL: TestS10V7ReaderRefusesV8MarkerBeforeContent (0.00s)
+    s10_fieldspec_test.go:24: registry version = "s8-fieldspec-v7", want s10-fieldspec-v8
+FAIL
+FAIL	github.com/jackli/frank/test/fixtures	0.745s
+FAIL
diff --git a/internal/config/config.go b/internal/config/config.go
index af9eb23..5b88d91 100644
--- a/internal/config/config.go
+++ b/internal/config/config.go
@@ -304,7 +304,7 @@ func preflightMemberMarkers(loaded map[string][]byte) error {
 		}
 	}
 	if data, ok := loaded["fieldspec"]; ok {
-		if err := ValidateFieldspecReaderMarker(data, "s7a-fieldspec-v5", "s8-fieldspec-v6", "s8-fieldspec-v7"); err != nil {
+		if err := ValidateFieldspecReaderMarker(data, "s7a-fieldspec-v5", "s8-fieldspec-v6", "s8-fieldspec-v7", "s10-fieldspec-v8"); err != nil {
 			return err
 		}
 	}
@@ -412,7 +412,7 @@ func validateFieldspecMarkerTransition(current, candidate []byte) error {
 	if err != nil {
 		return ErrConfigVersionTransition
 	}
-	if from == to || from == "s7a-fieldspec-v5" && to == "s8-fieldspec-v6" || from == "s8-fieldspec-v6" && to == "s8-fieldspec-v7" {
+	if from == to || from == "s7a-fieldspec-v5" && to == "s8-fieldspec-v6" || from == "s8-fieldspec-v6" && to == "s8-fieldspec-v7" || from == "s8-fieldspec-v7" && to == "s10-fieldspec-v8" {
 		return nil
 	}
 	return ErrConfigVersionTransition
diff --git a/internal/fieldspec/registry.json b/internal/fieldspec/registry.json
index f175a99..5893743 100644
--- a/internal/fieldspec/registry.json
+++ b/internal/fieldspec/registry.json
@@ -1,5 +1,5 @@
 {
-  "version": "s8-fieldspec-v7",
+  "version": "s10-fieldspec-v8",
   "provenance": {
     "owner": "m-2",
     "design_doc_id": "F-S7-R2-COLGRAIN",
@@ -81,7 +81,7 @@
     "surface_intent": ["progress", "review_checkpoint", "advisory", "result"],
     "DESIGN_RECORD_KIND": ["design-doc", "audit-record", "direct-override"],
     "DESIGN_REVIEW_VERDICT": ["approve", "must-revise", "reject-narrow", "human-decision-required"],
-    "record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint"],
+    "record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint", "odb", "resummon_command"],
     "config_member": ["fieldspec", "engine", "catalog", "adoption"],
     "dispatch_status": ["read", "awaiting"]
   },
diff --git a/internal/fieldspec/registry_test.go b/internal/fieldspec/registry_test.go
index aa387f7..7083d74 100644
--- a/internal/fieldspec/registry_test.go
+++ b/internal/fieldspec/registry_test.go
@@ -15,8 +15,8 @@ import (
 func TestRegistryV2MemberParsesAndExposesLockedEnums(t *testing.T) {
 	reg := loadRegistry(t)
 
-	if reg.Version != "s8-fieldspec-v7" {
-		t.Fatalf("Version = %q, want s8-fieldspec-v7", reg.Version)
+	if reg.Version != "s10-fieldspec-v8" {
+		t.Fatalf("Version = %q, want s10-fieldspec-v8", reg.Version)
 	}
 	wantProvenance := map[string]string{
 		"owner":         "m-2",
diff --git a/internal/store/genesis.go b/internal/store/genesis.go
index 7535272..fd0383a 100644
--- a/internal/store/genesis.go
+++ b/internal/store/genesis.go
@@ -100,6 +100,15 @@ func genesisMemberBytes(name string, source []byte) ([]byte, error) {
 		return source, nil
 	}
 	predecessor := append([]byte(nil), source...)
+	if bytes.Contains(predecessor, []byte(`"version": "s10-fieldspec-v8"`)) {
+		v8Kinds := []byte(`"record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint", "odb", "resummon_command"]`)
+		v7Kinds := []byte(`"record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint"]`)
+		if bytes.Count(predecessor, v8Kinds) != 1 {
+			return nil, fmt.Errorf("fieldspec v8 record_kind mismatch")
+		}
+		predecessor = bytes.Replace(predecessor, []byte(`"version": "s10-fieldspec-v8"`), []byte(`"version": "s8-fieldspec-v7"`), 1)
+		predecessor = bytes.Replace(predecessor, v8Kinds, v7Kinds, 1)
+	}
 	if bytes.Contains(predecessor, []byte(`"version": "s8-fieldspec-v7"`)) {
 		claimRow := []byte("    " + executableClaimsV7Row + ",\n")
 		if bytes.Count(predecessor, claimRow) != 1 {
diff --git a/test/fixtures/s10_fieldspec_test.go b/test/fixtures/s10_fieldspec_test.go
new file mode 100644
index 0000000..2040383
--- /dev/null
+++ b/test/fixtures/s10_fieldspec_test.go
@@ -0,0 +1,51 @@
+package fixtures_test
+
+import (
+	"encoding/json"
+	"errors"
+	"os"
+	"path/filepath"
+	"testing"
+
+	"github.com/jackli/frank/internal/config"
+	"github.com/jackli/frank/internal/fieldspec"
+)
+
+func TestS10V7ReaderRefusesV8MarkerBeforeContent(t *testing.T) {
+	v8, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
+	if err != nil {
+		t.Fatalf("read v8 registry: %v", err)
+	}
+	var planted map[string]any
+	if err := json.Unmarshal(v8, &planted); err != nil {
+		t.Fatalf("decode v8 registry: %v", err)
+	}
+	if got := planted["version"]; got != "s10-fieldspec-v8" {
+		t.Fatalf("registry version = %q, want s10-fieldspec-v8", got)
+	}
+	planted["fields"] = "content-must-not-be-interpreted"
+	plantedBytes, err := json.Marshal(planted)
+	if err != nil {
+		t.Fatalf("marshal planted registry: %v", err)
+	}
+	if _, err := fieldspec.Parse(plantedBytes); err == nil {
+		t.Fatal("planted fieldspec content unexpectedly valid")
+	}
+	if err := config.ValidateFieldspecReaderMarker(
+		plantedBytes,
+		"s7a-fieldspec-v5",
+		"s8-fieldspec-v6",
+		"s8-fieldspec-v7",
+	); !errors.Is(err, config.ErrConfigLoad) {
+		t.Fatalf("v7 reader marker error = %v, want config-load", err)
+	}
+	if err := config.ValidateFieldspecReaderMarker(
+		plantedBytes,
+		"s7a-fieldspec-v5",
+		"s8-fieldspec-v6",
+		"s8-fieldspec-v7",
+		"s10-fieldspec-v8",
+	); err != nil {
+		t.Fatalf("v8 marker preflight interpreted planted content: %v", err)
+	}
+}
diff --git a/test/fixtures/s8_exit_gate_test.go b/test/fixtures/s8_exit_gate_test.go
index 52c687b..bd7a600 100644
--- a/test/fixtures/s8_exit_gate_test.go
+++ b/test/fixtures/s8_exit_gate_test.go
@@ -39,10 +39,7 @@ func TestS8ExitGateFreshGenesisActivationAndDogfoodLegs(t *testing.T) {
 
 	v6 := s8FieldspecV6Bytes(t)
 	s8CommitOperatorConfigChange(t, st, "fieldspec", v6)
-	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
-	if err != nil {
-		t.Fatalf("read v7 registry: %v", err)
-	}
+	v7 := s8FieldspecV7Bytes(t)
 	s8CommitOperatorConfigChange(t, st, "fieldspec", v7)
 	afterV7 := s8PinnedStore(t, root)
 	t.Logf("v7_transition_new_digest=%s", afterV7.Digest)
@@ -173,10 +170,7 @@ func TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate(t *testing.T) {
 		t.Fatalf("DescribeTools v6: %v", err)
 	}
 	_ = v6Lane.Close()
-	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
-	if err != nil {
-		t.Fatalf("read v7 registry: %v", err)
-	}
+	v7 := s8FieldspecV7Bytes(t)
 	h.submit(t, operator, s8ConfigChangeRecord(t, root, "fieldspec", v7))
 	_ = operator.Close()
 	h.stop(t)
@@ -380,10 +374,7 @@ func TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate(t *testing.T) {
 }
 
 func TestS8V6ReaderRefusesV7MarkerBeforeContent(t *testing.T) {
-	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
-	if err != nil {
-		t.Fatalf("read v7 registry: %v", err)
-	}
+	v7 := s8FieldspecV7Bytes(t)
 	var planted map[string]any
 	if err := json.Unmarshal(v7, &planted); err != nil {
 		t.Fatalf("decode v7 registry: %v", err)
@@ -416,10 +407,7 @@ func TestS8ExecutableClaimTypedRejects(t *testing.T) {
 		t.Fatalf("open typed-reject store: %v", err)
 	}
 	s8CommitOperatorConfigChange(t, st, "fieldspec", s8FieldspecV6Bytes(t))
-	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
-	if err != nil {
-		t.Fatalf("read typed-reject v7 registry: %v", err)
-	}
+	v7 := s8FieldspecV7Bytes(t)
 	s8CommitOperatorConfigChange(t, st, "fieldspec", v7)
 	pinned := s8PinnedStore(t, root)
 	meta := seat.SeatMeta{Name: "s8.implementer", Role: "implementer"}
diff --git a/test/fixtures/s8_registry_changeset_test.go b/test/fixtures/s8_registry_changeset_test.go
index a09581a..f909464 100644
--- a/test/fixtures/s8_registry_changeset_test.go
+++ b/test/fixtures/s8_registry_changeset_test.go
@@ -16,7 +16,10 @@ import (
 )
 
 func TestS8RegistryChangesetIsExactV7OptionB(t *testing.T) {
-	reg := loadS5Registry(t)
+	reg, err := fieldspec.Parse(s8FieldspecV7Bytes(t))
+	if err != nil {
+		t.Fatalf("parse v7 registry: %v", err)
+	}
 	if reg.Version != "s8-fieldspec-v7" {
 		t.Fatalf("version = %q, want s8-fieldspec-v7", reg.Version)
 	}
@@ -100,10 +103,7 @@ func TestS8FXCFG10MemberTransitionRelation(t *testing.T) {
 		t.Fatalf("read v5 registry: %v", err)
 	}
 	v6 := s8FieldspecV6Bytes(t)
-	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
-	if err != nil {
-		t.Fatalf("read v7 registry: %v", err)
-	}
+	v7 := s8FieldspecV7Bytes(t)
 	if err := config.ValidateMemberTransition("fieldspec", v5, v6); err != nil {
 		t.Fatalf("lawful v5-to-v6 transition: %v", err)
 	}
@@ -186,10 +186,7 @@ func s8ConfigChangeRecord(t *testing.T, root, member string, body []byte) record
 
 func s8FieldspecV6Bytes(t *testing.T) []byte {
 	t.Helper()
-	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
-	if err != nil {
-		t.Fatalf("read v7 registry: %v", err)
-	}
+	v7 := s8FieldspecV7Bytes(t)
 	v6 := bytes.Replace(v7, []byte(`"version": "s8-fieldspec-v7"`), []byte(`"version": "s8-fieldspec-v6"`), 1)
 	lines := bytes.Split(v6, []byte{'\n'})
 	for i, line := range lines {
@@ -201,3 +198,19 @@ func s8FieldspecV6Bytes(t *testing.T) []byte {
 	t.Fatal("v7 executable_claims row missing")
 	return nil
 }
+
+func s8FieldspecV7Bytes(t *testing.T) []byte {
+	t.Helper()
+	v8, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
+	if err != nil {
+		t.Fatalf("read v8 registry: %v", err)
+	}
+	v8Marker := []byte(`"version": "s10-fieldspec-v8"`)
+	v8Kinds := []byte(`"record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint", "odb", "resummon_command"]`)
+	v7Kinds := []byte(`"record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint"]`)
+	if bytes.Count(v8, v8Marker) != 1 || bytes.Count(v8, v8Kinds) != 1 {
+		t.Fatal("live v8 registry does not match the historical-v7 inverse fixture")
+	}
+	v7 := bytes.Replace(v8, v8Marker, []byte(`"version": "s8-fieldspec-v7"`), 1)
+	return bytes.Replace(v7, v8Kinds, v7Kinds, 1)
+}
diff --git a/test/fixtures/s8_supply_test.go b/test/fixtures/s8_supply_test.go
index 3296f47..b303058 100644
--- a/test/fixtures/s8_supply_test.go
+++ b/test/fixtures/s8_supply_test.go
@@ -191,7 +191,7 @@ func TestS8FXSUP3MemberTransitionRefusesDuplicateMarkers(t *testing.T) {
 			case "engine":
 				candidate = bytes.Replace(current, []byte(`"version":2`), []byte(`"version":2,"version":2`), 1)
 			case "fieldspec":
-				candidate = bytes.Replace(current, []byte(`"version": "s8-fieldspec-v7"`), []byte(`"version": "s8-fieldspec-v7", "version": "s8-fieldspec-v7"`), 1)
+				candidate = bytes.Replace(current, []byte(`"version": "s10-fieldspec-v8"`), []byte(`"version": "s10-fieldspec-v8", "version": "s10-fieldspec-v8"`), 1)
 			case "catalog":
 				candidate = bytes.Replace(current, []byte(`"version": "s8-v1"`), []byte(`"version": "s8-v1", "version": "s8-v1"`), 1)
 			}

```

Next requested action: s10.planner review the whole T1 commit and route the two extended Step-5 confirmations to m-2.planner and m-7.planner. On review close, this seat will resume at T2 in locked order. No merge authority is requested or implied.

