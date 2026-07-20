## PLAN-REVIEW — s9 pair PLAN approved: rev11 adoption hash exact, ten-row fence all-in at named-seam grain, owner holds and fail-closed invariants preserved

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-build-plan-review-r1
PARENT_DISPATCH_ID: s9-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — same-owner plan review only; implementation still requires a later lineage-valid token, T7/T8 retain their owner holds, and merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev11-s9.1
PLAN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s9-build-plan/PLAN-planner-20260713-183000.md
FROM: s9.implementer
TO: s9.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve — the adopted plan-of-record is hash-exact; the execution delta adds no mechanic or authority; all planned loci are fenced and seam-named; T7/T8 and the blocked ledger remain closed on their named gates

PLAN_REVIEW_VERDICT: approve

I reviewed `master/relays/s9-build-plan/PLAN-planner-20260713-183000.md` against both s9 boots, the Option-A staffing ruling, m-3's approved plan-of-record and r12 review, the three master rulings, conditions (a)–(g), and live `frank@39474d0` source. The PLAN is directly addressed, lint-clean, and contains no implementation token.

This approval closes only the fresh pair's PLAN-REVIEW gate. It does not open T7/T8, authorize an owner-byte adaptation, authorize any blocked-ledger item, or grant merge authority.

## Findings

### F1 — adoption and lineage are faithful

- The adopted plan-of-record hash is exactly `f2a3cb3a5dc63a0814352f654d30de162b8af43816c6b7688d05960ab7b8853b`.
- The instrument correctly consumes the r12-approved rev11 mechanics plus the master-granted rev12 team-shape correction: `s9.planner` and `s9.implementer` are the build pair; m-3 remains outside as guide and owner-fidelity.
- The void `s9-build-impl/IMPL-planner-20260713-214500.md` token is excluded from this chain. The coming delegated token must parent through this review.
- No design lock is self-authored by the fresh pair. `DESIGN_DOC_ID` is context; the reviewed design and plan lineage remains in the m-3/master chain.

### F2 — the ten-row fence is all-in and complete at the approved task grain

Every production, test, and evidence locus named by T1–T9 resolves inside the block: the whole m-3 observe package; the m-7 `config.go` and `main.go` loci; the m-2 registry and validator files; the standing executor fence; the exit-fixture directory; and the slice evidence root. `internal/observe/fs_worker.go` is correctly a planned-new file under the observe directory row.

Live source confirms the promised integration points at `39474d0`: `Supply`/engine-v2 validation and cloning live in `internal/config/config.go`; `RegistryEnv`/`Evaluator` and the existing map clones live in `internal/observe/registry.go`; the serialized-path filesystem probe is present in `internal/observe/checks_base.go`; and the production registry composition is in `cmd/frank/main.go`. The plan therefore has writers for every behavior it promises.

SCOPE_DIFF:
- frank/internal/observe/ -> in
- frank/internal/config/config.go -> in
- frank/internal/fieldspec/registry.json -> in
- frank/internal/fieldspec/registry_test.go -> in
- frank/internal/fieldspec/validate.go -> in
- frank/internal/fieldspec/validate_test.go -> in
- frank/internal/executor/executor.go -> in
- frank/cmd/frank/main.go -> in
- frank/test/fixtures/ -> in
- frank/.relays/s9/ -> in
SCOPE_DIFF_RESULT: all-in

The file-wide rows do not erase the PLAN's seam restrictions. In particular, `executor.go` has no licensed s9 edit seam and any diff there stops; `config.go` and `main.go` are limited to the master-activated m-7 owner bytes; registry/validator edits are limited to their held T7/T8 tasks; and `test/fixtures/` is limited to the named s9 fixture/regression work.

### F3 — shared-file ordering is mechanically reviewable

- T1 owns the shared FS-worker extraction before T2 and T4 consume it.
- `registry.go` is ordered T2 entry/params → T3 `Evaluator` binding → T4 `RegistryEnv.LaneVCS` and clone.
- `gate.go` is ordered T3 claim-row thickening → T4 extension.
- T4 separates m-7-authored owner bytes from the m-3 consumer/terminal implementation. The owner bytes land verbatim as their own commit and receive an end-of-slice m-7 confirmation; a mismatch stops rather than becoming local adaptation.
- Per-task RED/GREEN capture, commit boundaries, and the diff→license table make each shared-file arm reviewable under B11 without per-task review relays.

### F4 — holds, acceptance, and authority boundaries remain closed correctly

- T7 opens only on the m-2 co-sign; T8 opens only on m-1 plus m-2 confirmation. Until then they report-and-hold under condition (f), never proceed on assumption and never become `t.Skip` stubs.
- B1, B-diff, B3, B2, and B4 remain in governance relays behind their named gates. The build does not silently implement any of them.
- The fail-closed invariant is preserved: the opaque-accept branch fires only for governed `lane_vcs: none`; nil/undeclared or `git` never opaque-accepts; pre-v3/nil still runs the full §13 observation and retains the Option-2 E0 floor; marker-derived acceptance is prohibited.
- T3 preserves the locked `CheckVerdict` shape and the `output-truncated` passing warning while binding identity and disposition conductor-side. `executor.go` remains untouched.
- Conditions (a)–(g), byte-exact `{accepted, rejected, held}`, R2, Rails A/B, I-PH, ten INV-CATALOG laws, both mechanical tables, sequence-honest file captures, label-equals-mechanism, and the operator-only merge gate remain binding.

The trust-critical and cross-domain surfaces are not new unruled escalation triggers in this execution delta: each is already master-activated and owner-gated at its named seam. Any byte outside those rulings re-engages conditions (b)–(d) and stops through `s9.planner` before edit.

## Token boundary

`s9.planner` may issue the delegated `s9-build-token` only if the relay:

1. is parented to `s9-build-plan-review-r1`, addressed solely to `s9.implementer`, and carries the required live token in the protocol-valid shape;
2. restates this exact ten-row all-in block plus the named-seam licenses, holds, and conditions (a)–(g);
3. cuts the m-7 owner-byte request beside it and does not claim those bytes are in hand until the owner returns them;
4. preserves T7/T8 as separately owner-gated after dispatch and stops on any owner mismatch, locked-contract change, cross-domain change, or out-of-fence diff; and
5. grants no merge authority. Merge remains operator-only under a later `HUMAN_MERGE_AUTHORIZATION` grant.

## Verification

- Exact incoming PLAN lint returned `OK`; relay-root mode also reported the target PLAN `OK` alongside known unrelated `INDEX.md` and historical lineage noise.
- `shasum -a 256` verified the adopted plan-of-record hash exactly.
- Mechanical readback returned `scope_rows=10` and `SCOPE_DIFF_RESULT: all-in`.
- All existing fence anchors resolved at `frank@39474d0`; `internal/observe/fs_worker.go` is absent as expected for T1's planned-new file.
- E1 source inspection: `internal/config/config.go:42-45,249-266,482-520,594-665`; `internal/observe/registry.go:90-145,224-280`; `internal/observe/checks_base.go:47-91`; `cmd/frank/main.go:155-164`; `internal/executor/executor.go:105-118`.
- Live index inspection found this PLAN as the latest relay directly addressed to `s9.implementer`; no superseding token or review exists.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay at `master/relays/s9-build-plan/PLAN-REVIEW-implementer-20260713-183332.md` and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, branch/worktree creation, commit, implementation dispatch, PR, or merge.

FINAL_GIT_STATUS_SHORT: unavailable — cwd `/Users/jack/Programming/harness` is not a git repository; post-write `frank/` main remained clean at `39474d0`.

Next requested action: `s9.planner` issues the lineage-valid same-owner token under `s9-build-token` with the exact fence, seam licenses, owner holds, conditions, and m-7 request above. Implementation remains held until that relay arrives.
