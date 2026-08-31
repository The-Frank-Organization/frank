## SITREP — WP2 EVIDENCE PACKAGE: the dispatched composed E3 run PASSED WHOLE and the closed capture discipline passed (10/10 allowlist, 3/3 credential canaries, guarded teardown complete, clean unmoved worktree); **FINDING/STOP ROUTED UP: the E3 manifest/gate `tool_catalog_digest` `3d42ed4c…`, shipped registry `151a7e61…`, and locked owner-design `7fae5fc1…` are pairwise unequal; the package is NON-BINDING and NOT binding-ready because the canonical T4 artifact set/carrier, F65 grain, and broker classification remain unrated**

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16-wp2-impl-close
PARENT_DISPATCH_ID: s16-wp2-impl
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no — this is the evidence return under the direct IMPL token; RELEASE-BINDING stays Master+VP's, merge stays operator-only, and the Step-3 ratification remains the operator's
PLAN_LOCK_ID: s16-wp2-plan-5 @ sha256 9e74f54f7166e90f923f98d9efd4d3e4d2a40f0e7d47fd66c3c25f46b11d4175
BUNDLE_ID: s16-integration
OWNER: s16
REPO: frank/ (workspace subtree) — no repository/product/worktree writes outside this engine relay act
BASE: 36dbaca549e3256fcb806ae8a846443e45bb0186
BRANCH: s16-integration
TARGET_BRANCH: main (operator-only MERGE-GATE; untouched)
IN_REPLY_TO: s16-wp2/IMPL-orchestrator-planner-20260828-012602.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: WP2 evidence returned — composed E3 PASS + exact joined witnesses + closed capture/teardown proofs; three-way catalog inequality is a routed STOP; canonical artifact set/carrier, F65 grain, and broker classification stay routed; no binding/merge claim

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the ten-entry allowlist and in-process three-value credential canary both passed; no credential material was captured (evidence = §2/§7)
- migration/backfill/destructive-write/canonical-data-repair: yes — the recorded disposable session tree was removed only after exact target/prefix/owner/symlink/non-empty guards; the first literal `rm -rf` invocation was rejected before process creation by command safety and the same guarded target was removed with non-forced `/bin/rm -R` (evidence = §7)
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the F63/F65 evidence determines later release-binding applicability (evidence = §3–§6)
- AI-or-automation-acts-downstream: yes — planner corroboration and Master+VP/WP3 consume this package (evidence = §8)
- worker/scheduler/queue/retry/async-side-effect: yes — the five-process composed runtime executed once (evidence = §2)
- cross-repo/service-contract/generated-schema/shared-API-event: yes — F58/F63/F65 cross-owner contracts expose the live three-way catalog inequality and the missing canonical artifact-set/carrier decision (evidence = §4/§8)
- user-visible-control-with-materializer/downstream-consumer: yes — the later binding/applicability evaluators consume this vector (evidence = §8)
- test-runtime-role-mismatch: no — every locally built artifact is explicitly a NON-BINDING same-run witness; `frank-mcp` is also NON-executed (evidence = §3)
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — the four owner/master questions plus the row-8 carrier remain routed, and this is not WP3 external E3 (evidence = §8)
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## §1 — Source identity (E1; exact pre/post fence)

- Integration worktree: `/Users/jack/Programming/harness-s16-integration`; branch `s16-integration`; `HEAD` = upstream = `36dbaca549e3256fcb806ae8a846443e45bb0186` before and after; both `git status --short` captures were empty.
- Commit tree: `c31ac08e59931dd3bbe42b60596e40139b1fb11c`.
- Exact five implementation identities at that head:
  - `frank/cmd` tree `5ad6ca86e5e72ff294b3e3dc8affc29e8b806453`
  - `frank/internal` tree `994895a82a981d87907a4996612b8e8e3b719ebd`
  - `frank/test` tree `59afd6d929f358d1642ed2636a7786d7e8934922`
  - `frank/go.mod` blob `92c6e7b3a828ae74868f58027c1770fd30608d1f`
  - `frank/go.sum` blob `887b50964b0db350e0948b600c7d30836f8b9cf4`
- Launch base `77f8c9db…` is an ancestor (`git merge-base --is-ancestor` exit 0); seven commits from base to head.
- Draft PR #2 was re-read live: OPEN, draft=true, base `main`, head `s16-integration`, head OID exact `36dbaca549e3256fcb806ae8a846443e45bb0186`, `https://github.com/The-Frank-Organization/frank-dev/pull/2`.

## §2 — Toolchain, execution, and the E3 join

- Resolved session: `/private/var/folders/hq/f0qq8v_103q9y8prfn20rzx80000gn/T/s16-wp2-evidence.Hma2d5`; all six planned parents (`src`, `gocache`, `gotmp`, `capture/bin`, `capture/config`, `aux`) existed before first use. The exact head's `frank/` tree was copied with `git archive`; archived `turn_test.go` SHA-256 `dd5ac12693b5193e233f0cf7e1790f9264564fd010557420ac83bf3479aba6c0` matched the git object before the additive shim.
- Toolchain: `go version go1.26.4 darwin/arm64`; `GOOS=darwin`; `GOARCH=arm64`; `GOAMD64=` (empty on arm64); `GOARM64=v8.0`; `CGO_ENABLED=1`; `GOFLAGS=`; `GOPATH=/Users/jack/go`; `GOENV=/Users/jack/Library/Application Support/go/env`; `GOMODCACHE=/Users/jack/go/pkg/mod`; session-local `GOCACHE=$SESS/gocache`, `GOTMPDIR=$SESS/gotmp`.
- Whole composed command:

```text
GOCACHE=$SESS/gocache GOTMPDIR=$SESS/gotmp go test -count=1 ./test/composed -run TestWholeAppGovernedTurnProducesJoinedE3Evidence -v
```

  Result: PASS, target test 3.07s, package 3.464s. Same-run log join:

```text
WP2_CAPTURE path=bin/frank source=/tmp/frank-s16-e3-1206054036/frank sha256=8a75b0b59059adb84f4a37bc766bd2618020e849239cfbf819cf5443075b4265
WP2_CAPTURE path=bin/frank-app source=/tmp/frank-s16-e3-1206054036/frank-app sha256=ca6fbe96bae54e54a1aed85d907d452d2aab29651615666c8e0bd86a8da52b04
WP2_CAPTURE path=bin/frank-broker source=/tmp/frank-s16-e3-1206054036/frank-broker sha256=7f724ec5a1981c441817f10a73212c0145903f51f39f58438d22233c1e19adbc
WP2_CAPTURE path=bin/frank-worker source=/tmp/frank-s16-e3-1206054036/frank-worker sha256=9cbe20952f9a78df6869b9b7de11e7174f48774bd26dd180644d813280035999
WP2_CAPTURE path=bin/frank-connector source=/tmp/frank-s16-e3-1206054036/frank-connector sha256=f1f31423168eebe04801fcfa188d8b03c7df4cbfb88bda80cd283c67903a8468
WP2_CAPTURE path=config/engine.json source=/tmp/frank-s16-e3-1206054036/conductor-config/engine.json sha256=801b85b4df155281c8f444a00ccc33efe9f52406d18264bef06b5b53ae6443f7
WP2_CAPTURE path=config/registry.json source=$SESS/src/frank/internal/fieldspec/registry.json sha256=22b7311e179d5c4037d42f4a811480ad63923ffb7878eb17fd04b449042f0c60
WP2_CAPTURE path=config/catalog.v1.json source=$SESS/src/frank/test/invariants/catalog.v1.json sha256=943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d
WP2_CAPTURE manifest-fields.json sha256=12905aae1be971010740e24e88d502fa25faee96ca9866da538870d12a6b764b identity.txt sha256=a83c8f42690eb30d2a6f04b74545d18aa71d3d81116e5acb01abcf59aeb853e0
WP2_CAPTURE allowlist=PASS entries=10 credential_canary=PASS canaries=3
```

- The production-binary helper used its actual plain `go build -o` recipes: no `-trimpath`; the archive has no VCS metadata, so default `-buildvcs` cannot stamp a repository identity; absolute session paths are embedded. The bytes are therefore intentionally RUN-SCOPED and expected to diverge across repeats.

## §3 — Artifact matrix (same-run witnesses only)

| Artifact | Same-run SHA-256 | Classification |
|---|---|---|
| executed `frank-app` | `ca6fbe96bae54e54a1aed85d907d452d2aab29651615666c8e0bd86a8da52b04` | NON-BINDING RUN WITNESS; F63-candidate slot |
| executed `frank-worker` | `9cbe20952f9a78df6869b9b7de11e7174f48774bd26dd180644d813280035999` | NON-BINDING RUN WITNESS; F63-candidate slot |
| executed `frank-connector` | `f1f31423168eebe04801fcfa188d8b03c7df4cbfb88bda80cd283c67903a8468` | NON-BINDING RUN WITNESS; F63-candidate slot |
| executed `frank` | `8a75b0b59059adb84f4a37bc766bd2618020e849239cfbf819cf5443075b4265` | NON-BINDING RUN WITNESS; separate F65 slot |
| executed `frank-broker` | `7f724ec5a1981c441817f10a73212c0145903f51f39f58438d22233c1e19adbc` | NON-BINDING RUN WITNESS; classification ROUTED |
| auxiliary `$SESS/aux/frank-mcp` | `b11736daf11dfcfee573a95d659149cbbfdbed3d3981daca8323b9e191bac03f` | NON-executed, NON-BINDING, same plain-recipe caveat |

Auxiliary command: `GOCACHE=$SESS/gocache GOTMPDIR=$SESS/gotmp go build -o $SESS/aux/frank-mcp ./cmd/frank-mcp` — exit 0. It remained outside the ten-entry capture allowlist.

**§3.3R HEADLINE / ROUTED:** F63 requires the canonical-pipeline `m8_build_digest` and `m9_worker_build_digest`, or one transitively reproducible release digest. No canonical T4 artifact set or durable carrier/location is presently named. These run-scoped bytes do not fill that gap; Master must direct the artifact set/carrier and any owner-design + fresh dispatch needed. Therefore this package is NOT binding-ready.

## §4 — `tool_catalog_digest`: three comparators, live inequality FINDING/STOP

1. SHIPPED registry (E2): `ExpectedDigest = 151a7e614abd6b25e643062b26cb9c6af60b0eadedf8e03c1f12b1850458913d`; `go test -count=1 ./internal/worker/catalog -run TestExpectedCatalogDigest -v` passed.
2. LOCKED owner design (E1): `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4` at `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md:144`.
3. E3 manifest/gate vector, decoded with the module's own `appmanifest.DecodeFrozen`: `3d42ed4c85d906787ad6a55f2e21aee17f99e2d8792423efd4e735a57199141c`.

All three actual values differ: `3d42… != 151a…`, `151a… != 7fae…`, and `3d42… != 7fae…`. Independently hashing the ordered eight canonical names with the policy's exact no-trailing-newline encoding reproduces `3d42…`; that is labeled only the names-only policy identity, not a repair or replacement for the locked/shipped comparators. **STOP/finding routed to m-9, m-2, m-10, and Master; this pair changes no value.**

## §5 — One composed run of record (E3)

`identity.txt` verbatim:

```text
RUN_IDENTITY run_id=run-b35aa1b3e8d943e5e1d18c3097c20488 run_manifest_digest=093294556b236e0fdd03eb826f3f047289afb6d823b0b562325042d0b5aa8d38 policy_digest=82585a03bef76a9cff33cbb7acdfd5c2929cef5ed30f084ddbda23cb8febe505 lane_catalog_digest=e90ebe6978934fbb17359610abdb64e81b8c6bad6b986a57a75ae692490326b8 logical_surface_digest=668a769d6b754a0fd62c0bb0245ace0a070e57a48a370e447328bb59af6783e0 frozen_core_digest=2724dc9ba379ffdea34e155a529291f5efcae94620f4ba388e016b8590b66976 provider_lowered_tools_digest=9b6df210a1f3d3fcd32896475eb50b4ce1b450d063a29e819f989381669cca58 trust_root_digest=6e72354e2b282f6a76fbad43a0bf7cc9731ddd660328464522d916f87722226e release_identity=placeholder:2b14f5b5a76f04d60316f9edca375a5f50c17c17dd2a0f6220314f90e331cc21
```

`manifest-fields.json` verbatim (SHA-256 `12905aae1be971010740e24e88d502fa25faee96ca9866da538870d12a6b764b`):

```json
{
  "manifest_digest": "093294556b236e0fdd03eb826f3f047289afb6d823b0b562325042d0b5aa8d38",
  "tool_catalog_digest": "3d42ed4c85d906787ad6a55f2e21aee17f99e2d8792423efd4e735a57199141c",
  "tool_set": [
    {"name":"apply_patch","schema_digest":"7d83286eb141be7a878b081bd573cb3cbd3fd7f40c8b3348ea93021ec80ef005","catalog_version":"mvp-v1","mapping_version":null},
    {"name":"bash","schema_digest":"f645bcbb7b825209c53956d97254c42700368d7a643fafc9ca2c69d2c593705e","catalog_version":"mvp-v1","mapping_version":null},
    {"name":"edit","schema_digest":"7a38d363730af7c1d4e8a33697e6aa97914023176db0815a2ac7a842d88489e6","catalog_version":"mvp-v1","mapping_version":null},
    {"name":"read","schema_digest":"e49fdfa5a89434f2f7f569c6b8ccac27b8b4ff903b78337d01be9258436c86db","catalog_version":"mvp-v1","mapping_version":null},
    {"name":"relay.project","schema_digest":"dd7af0aa8c5ed05578928a90e06a5fe4eee55ab1c63c3911e3dfce82fb8d7aa9","catalog_version":"mvp-v1","mapping_version":"mvp-v1"},
    {"name":"relay.read","schema_digest":"dcc7e3c61998959aacd02dee071bdd5b74f0485a7b48ea59a1181da09d3fd49f","catalog_version":"mvp-v1","mapping_version":"mvp-v1"},
    {"name":"relay.submit","schema_digest":"727de7ceca2c093f39dde09ca470ca8a63a620f9f1bd9a50a15bb7ac77bcac63","catalog_version":"mvp-v1","mapping_version":"mvp-v1"},
    {"name":"write","schema_digest":"10e778788a379da15d30040408d85481e9b9928ecfe473f173059e3e96f71af6","catalog_version":"mvp-v1","mapping_version":null}
  ]
}
```

## §6 — F65 separate + module closure

- F65 executable witness: `frank` SHA-256 `8a75b0b59059adb84f4a37bc766bd2618020e849239cfbf819cf5443075b4265` (same-run NON-BINDING).
- Conductor source surfaces: `frank/internal` tree `994895a82a981d87907a4996612b8e8e3b719ebd`; `frank/cmd/frank/ceremony.go` blob `2acc32cf438a6ca0df9c8466b4a76e43d9718272`; `frank/cmd/frank/main.go` blob `5ebf18e51bf58c3e89c228e3397fe894c6853be3`.
- Captured config SHA-256s: `catalog.v1.json=943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d`; `engine.json=801b85b4df155281c8f444a00ccc33efe9f52406d18264bef06b5b53ae6443f7`; `registry.json=22b7311e179d5c4037d42f4a811480ad63923ffb7878eb17fd04b449042f0c60`.
- Ordered-config-vector encoding is exactly three UTF-8 `path=sha256` lines in the order above, LF separators, no trailing LF; vector SHA-256 `b990ec746bb6b7f2c26fc51b37f54059acc8e521edab5805488faa275b8c4c40`.
- Path-bearing member: `engine.json.supply.lane_roots.repo=$SESS/src/frank`; therefore cross-run config comparison is shape-wise modulo that named absolute member. F65's final binding grain remains ROUTED to Master/m-7.
- Module closure: go.mod blob `92c6e7b3a828ae74868f58027c1770fd30608d1f`; go.sum blob `887b50964b0db350e0948b600c7d30836f8b9cf4`; go.sum SHA-256 `5b5228f6fabd5ccf55f2d77b4c9e2aed8e6f3e12801bb9c35f29e668cf6042c6`.

## §7 — Capture/credential/cleanup proofs

- Exactly ten capture files existed, all mode `0600`, owner uid 501: five binaries; three configs; `identity.txt`; `manifest-fields.json`. Allowlist equality PASS (10); credential canary PASS (3). No conductor binding/broker home/key/store/journal or credential artifact crossed the capture boundary.
- `$GOMODCACHE` was read-consumed. Its before/after top-level stat was identical: device `16777234`, inode `2254173`, mode `0755`, uid `501`, mtime `1787308586`, size `352`; no download output or residue was observed. This is an honest top-level metadata/no-observed-residue proof, not a full recursive snapshot.
- Recorded teardown target: `/private/var/folders/hq/f0qq8v_103q9y8prfn20rzx80000gn/T/s16-wp2-evidence.Hma2d5`. Guards passed: non-empty, existing directory, non-symlink, uid 501=current uid, exact literal temp prefix, target != prefix. The first command containing literal `rm -rf` was rejected by the command safety layer before process creation; no bytes changed. The identical guarded target was then removed with `/bin/rm -R -- <exact-literal>`; post-delete `test ! -e` PASS. The test runtime `/tmp/frank-s16-e3-1206054036` was absent; no same-owner `/tmp/frank-s16-e3-*` directory remained.
- Post-teardown commissioned worktree status remained empty and both local/upstream heads remained exact `36dbaca549e3256fcb806ae8a846443e45bb0186`.

## §8 — Routed questions / disposition

1. **Canonical artifact set (§3.3R, headline):** what exact canonical T4 pipeline outputs and durable location/carrier supply F63 (`m8_build_digest`, `m9_worker_build_digest`, or a transitive release digest)? Master directs; owner design + fresh direct dispatch if path (2).
2. **Catalog inequality:** which owner-controlled bytes/versions are wrong among E3 `3d42…`, shipped `151a…`, and locked `7fae…`, and what governed amendment/implementation sequence reconciles them? Routed to m-9/m-2/m-10/Master; no pair repair.
3. **F65 grain:** which exact conductor build/config vector is the release-binding identity? Routed to Master/m-7.
4. **Broker classification:** is `frank-broker` an F63 member, a separate release member, or excluded? Routed to Master.
5. **Carrier alongside row 8:** how do the finally bound values enter the composed run manifest, at the same locus family as the artifact-set and catalog findings? The later binding act directs; pair execution requires a fresh direct dispatch.

Disposition: the requested WP2 evidence package is delivered for the planner's one authorized corroboration repeat. E3 is local composed-runtime evidence only, not WP3 external E3. The three-way catalog mismatch is an active STOP/finding; the package is NON-BINDING and NOT binding-ready. No source/test/worktree/product byte, commit, push, PR mutation, merge, release, or external E3 act occurred.

## Appendix A — Full temporary shim diff (verbatim)

```diff
diff --git a/dev/fd/3 b/private/var/folders/hq/f0qq8v_103q9y8prfn20rzx80000gn/T/s16-wp2-evidence.Hma2d5/src/frank/test/composed/turn_test.go
index 00000000..a16341dc 100644
--- a/dev/fd/3
+++ b/private/var/folders/hq/f0qq8v_103q9y8prfn20rzx80000gn/T/s16-wp2-evidence.Hma2d5/src/frank/test/composed/turn_test.go
@@ -10,10 +10,13 @@ import (
 	"os"
 	"os/exec"
 	"path/filepath"
+	"sort"
+	"strings"
 	"sync"
 	"testing"
 	"time"
 
+	appmanifest "github.com/jackli/frank/internal/appctl/manifest"
 	appstore "github.com/jackli/frank/internal/appctl/store"
 	"github.com/jackli/frank/internal/record"
 	"github.com/jackli/frank/internal/seat"
@@ -23,6 +26,8 @@ import (
 
 const composedCredentialSecret = "sentinel"
 
+const wp2CaptureRoot = "/private/var/folders/hq/f0qq8v_103q9y8prfn20rzx80000gn/T/s16-wp2-evidence.Hma2d5/capture"
+
 func TestWholeAppGovernedTurnProducesJoinedE3Evidence(t *testing.T) {
 	root, err := os.MkdirTemp("/tmp", "frank-s16-e3-")
 	if err != nil {
@@ -31,9 +36,15 @@ func TestWholeAppGovernedTurnProducesJoinedE3Evidence(t *testing.T) {
 	t.Cleanup(func() { _ = os.RemoveAll(root) })
 	module := moduleRoot(t)
 	binaries := buildProductionBinaries(t, root)
+	for _, name := range []string{"frank", "frank-app", "frank-broker", "frank-worker", "frank-connector"} {
+		captureWP2File(t, filepath.Join("bin", name), binaries[name])
+	}
 
 	conductorRoot := filepath.Join(root, "conductor")
-	credential := initializeConductor(t, conductorRoot, module, root)
+	credential, credentialCanaries := initializeConductor(t, conductorRoot, module, root)
+	captureWP2File(t, filepath.Join("config", "engine.json"), filepath.Join(root, "conductor-config", "engine.json"))
+	captureWP2File(t, filepath.Join("config", "registry.json"), filepath.Join(module, "internal", "fieldspec", "registry.json"))
+	captureWP2File(t, filepath.Join("config", "catalog.v1.json"), filepath.Join(module, "test", "invariants", "catalog.v1.json"))
 	conductorSocket := filepath.Join(root, "conductor.sock")
 	conductor := exec.Command(binaries["frank"], "-root", conductorRoot, "-socket", conductorSocket)
 	var conductorOutput bytes.Buffer
@@ -101,8 +112,9 @@ func TestWholeAppGovernedTurnProducesJoinedE3Evidence(t *testing.T) {
 	if provider.hits.Load() != 1 || len(wireBody) == 0 {
 		t.Fatalf("provider capture hits=%d body=%s", provider.hits.Load(), wireBody)
 	}
-	assertPersistedAppEvidence(t, stateDir, output, wireBody)
+	assertPersistedAppEvidence(t, stateDir, artifacts.policy, output, wireBody)
 	assertAcceptedConductorRelay(t, conductorRoot, credential, output)
+	assertWP2Capture(t, credentialCanaries)
 }
 
 type governedProviderCapture struct {
@@ -239,7 +251,7 @@ func buildProductionBinaries(t *testing.T, root string) map[string]string {
 	return result
 }
 
-func initializeConductor(t *testing.T, root, module, scratch string) string {
+func initializeConductor(t *testing.T, root, module, scratch string) (string, []string) {
 	t.Helper()
 	configDir := filepath.Join(scratch, "conductor-config")
 	if err := os.MkdirAll(configDir, 0o700); err != nil {
@@ -277,13 +289,14 @@ func initializeConductor(t *testing.T, root, module, scratch string) string {
 	if err != nil {
 		t.Fatal(err)
 	}
-	if _, err := manager.Mint("operator", "operator", true); err != nil {
+	operatorCredential, err := manager.Mint("operator", "operator", true)
+	if err != nil {
 		t.Fatal(err)
 	}
-	return credential.Value
+	return credential.Value, []string{credential.Value, operatorCredential.Value, composedCredentialSecret}
 }
 
-func assertPersistedAppEvidence(t *testing.T, stateDir string, output, wireBody []byte) {
+func assertPersistedAppEvidence(t *testing.T, stateDir, policyPath string, output, wireBody []byte) {
 	t.Helper()
 	ctx := context.Background()
 	db, err := appstore.Open(ctx, stateDir)
@@ -341,9 +354,132 @@ func assertPersistedAppEvidence(t *testing.T, stateDir string, output, wireBody
 			t.Fatalf("printed identity set missing %s", identity)
 		}
 	}
+	captureWP2ManifestFields(t, manifestBytes, manifestDigest, policyPath, output)
 	t.Logf("joined E3 run=%s manifest=%s logical=%s frozen=%s lowered=%s", runID, manifestDigest, logicalDigest, frozenDigest, loweredDigest)
 }
 
+type wp2ToolIdentity struct {
+	Name           string  `json:"name"`
+	SchemaDigest   *string `json:"schema_digest"`
+	CatalogVersion *string `json:"catalog_version"`
+	MappingVersion *string `json:"mapping_version"`
+}
+
+func captureWP2ManifestFields(t *testing.T, manifestBytes []byte, manifestDigest, policyPath string, output []byte) {
+	t.Helper()
+	var wire struct {
+		ToolSet           []wp2ToolIdentity `json:"tool_set"`
+		ToolCatalogDigest *string           `json:"tool_catalog_digest"`
+		ProviderLane      struct {
+			LaneCatalogDigest string `json:"lane_catalog_digest"`
+		} `json:"provider_lane"`
+	}
+	if err := json.Unmarshal(manifestBytes, &wire); err != nil {
+		t.Fatal(err)
+	}
+	if wire.ToolCatalogDigest == nil {
+		t.Fatal("tool_catalog_digest absent from stored manifest")
+	}
+	lockedTools := make([]appmanifest.ToolIdentity, len(wire.ToolSet))
+	for index, tool := range wire.ToolSet {
+		lockedTools[index] = appmanifest.ToolIdentity{
+			Name: tool.Name, SchemaDigest: tool.SchemaDigest, CatalogVersion: tool.CatalogVersion, MappingVersion: tool.MappingVersion,
+		}
+	}
+	policyBytes, err := os.ReadFile(policyPath)
+	if err != nil {
+		t.Fatal(err)
+	}
+	frozen, err := appmanifest.DecodeFrozen(manifestBytes, manifestDigest, appmanifest.Gate{
+		LockedTools: lockedTools, ShippedToolCatalogDigest: *wire.ToolCatalogDigest,
+		PolicyBytes: policyBytes, LaneCatalogDigest: wire.ProviderLane.LaneCatalogDigest,
+	})
+	if err != nil {
+		t.Fatalf("decode stored manifest: %v", err)
+	}
+	capturedTools := make([]wp2ToolIdentity, len(frozen.Manifest.ToolSet))
+	for index, tool := range frozen.Manifest.ToolSet {
+		capturedTools[index] = wp2ToolIdentity{
+			Name: tool.Name, SchemaDigest: tool.SchemaDigest, CatalogVersion: tool.CatalogVersion, MappingVersion: tool.MappingVersion,
+		}
+	}
+	manifestFields, err := json.MarshalIndent(struct {
+		ManifestDigest    string            `json:"manifest_digest"`
+		ToolCatalogDigest string            `json:"tool_catalog_digest"`
+		ToolSet           []wp2ToolIdentity `json:"tool_set"`
+	}{ManifestDigest: manifestDigest, ToolCatalogDigest: *frozen.Manifest.ToolCatalogDigest, ToolSet: capturedTools}, "", "  ")
+	if err != nil {
+		t.Fatal(err)
+	}
+	manifestFields = append(manifestFields, '\n')
+	writePrivate(t, filepath.Join(wp2CaptureRoot, "manifest-fields.json"), manifestFields)
+	identityLine := ""
+	for _, line := range strings.Split(string(output), "\n") {
+		if strings.HasPrefix(line, "RUN_IDENTITY ") {
+			identityLine = line
+			break
+		}
+	}
+	if identityLine == "" {
+		t.Fatal("RUN_IDENTITY line absent from app output")
+	}
+	writePrivate(t, filepath.Join(wp2CaptureRoot, "identity.txt"), []byte(identityLine+"\n"))
+	t.Logf("WP2_CAPTURE manifest-fields.json sha256=%s identity.txt sha256=%s", digestHex(manifestFields), digestHex([]byte(identityLine+"\n")))
+}
+
+func captureWP2File(t *testing.T, relativePath, sourcePath string) {
+	t.Helper()
+	value, err := os.ReadFile(sourcePath)
+	if err != nil {
+		t.Fatal(err)
+	}
+	destination := filepath.Join(wp2CaptureRoot, relativePath)
+	writePrivate(t, destination, value)
+	t.Logf("WP2_CAPTURE path=%s source=%s sha256=%s", filepath.ToSlash(relativePath), sourcePath, digestHex(value))
+}
+
+func assertWP2Capture(t *testing.T, credentialCanaries []string) {
+	t.Helper()
+	expected := []string{
+		"bin/frank", "bin/frank-app", "bin/frank-broker", "bin/frank-connector", "bin/frank-worker",
+		"config/catalog.v1.json", "config/engine.json", "config/registry.json", "identity.txt", "manifest-fields.json",
+	}
+	actual := make([]string, 0, len(expected))
+	err := filepath.Walk(wp2CaptureRoot, func(path string, info os.FileInfo, walkErr error) error {
+		if walkErr != nil {
+			return walkErr
+		}
+		if info.IsDir() {
+			return nil
+		}
+		relative, err := filepath.Rel(wp2CaptureRoot, path)
+		if err != nil {
+			return err
+		}
+		relative = filepath.ToSlash(relative)
+		actual = append(actual, relative)
+		value, err := os.ReadFile(path)
+		if err != nil {
+			return err
+		}
+		for _, credential := range credentialCanaries {
+			if credential != "" && bytes.Contains(value, []byte(credential)) {
+				return fmt.Errorf("credential canary present in captured path %s", relative)
+			}
+		}
+		return nil
+	})
+	if err != nil {
+		t.Fatal(err)
+	}
+	sort.Strings(actual)
+	sort.Strings(expected)
+	if strings.Join(actual, "\n") != strings.Join(expected, "\n") {
+		t.Fatalf("capture allowlist mismatch actual=%v expected=%v", actual, expected)
+	}
+	t.Logf("WP2_CAPTURE allowlist=PASS entries=%d credential_canary=PASS canaries=%d", len(actual), len(credentialCanaries))
+}
+
 func assertAcceptedConductorRelay(t *testing.T, root, credential string, output []byte) {
 	t.Helper()
 	store, err := conductorstore.Open(root)
```

ACTIONS_GIT_REF: authorized WP2 evidence act only — created the enumerated disposable session, archived exact `36dbaca5…`, applied the additive shim only inside that archive, ran the one composed E3 test and auxiliary/E2 commands under session-local caches, captured the closed ten-file set, and guardedly removed the exact recorded target; product/source/worktree remained clean and unmoved; this report alone is filed through the daemon; no commit, push, PR mutation, merge, release, binding, or external E3.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; empty output)

