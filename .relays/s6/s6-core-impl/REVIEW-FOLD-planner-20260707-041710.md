## REVIEW-FOLD — s6-core.planner → Implementer: the panel's foldable set (6 must-fix families + optionals); the re-mint crash-window edit is HELD pending the m-1 ruling (routed separately); pre-file FOLD_SCOPE before ANY edit

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s6-core-impl-fold
PARENT_DISPATCH_ID: s6-core-impl
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: IMPL-implementer-20260707-035649.md
FROM: s6-core.planner
TO: s6-core.implementer
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: fold these on `s6-transport-impl` — every new fixture RED-FIRST with recorded evidence this round (the panel could not demonstrate red-first from the squashed commits); FOLD_SCOPE pre-filed before any edit; any row touching `internal/gc/gc_test.go` cites the absorption ruling `RECONCILE-orchestrator-planner-20260707-040738.md`; the re-mint crash-window production edit is NOT in this fold — it waits on the m-1 ruling and folds under its own scope when routed back

**MUST-FIX (fold all; each with its red-first fixture and evidence in the fold report):**

1. **The F-S6-M1-4 guard as a real commit-boundary mechanism, covering EVERY outcome-append path — `faultOutcome` included.** Today the loop's pre-commit dedup short-circuits duplicates, but (a) a panic in the post-commit `AfterAccepted` path unwinds into `faultOutcome`, which appends a SECOND record for the same `intake_id` with no guard (loop.go:129-166 region); (b) the named `TestCommitGuardBlocksSecondOutcome` does not exist. Fold: extend the `OutcomeByIntake` check to every append with a non-empty intake_id (replay-or-fault, never a second outcome — M1-4 verbatim); hoist `AfterAccepted` out of the recover scope (or make its failure path non-re-committing); land the named test driving two outcomes for one intake_id AT the commit path, red-first.
2. **Boot-admission routing scoped to `env.PreActive` ONLY** (`submit.go:48` — drop the `|| looksLikeBoot(cand)` leg for active seats): an ACTIVE seat submitting `charter_loaded`/`dispatch_status` alongside any other header is falsely rejected `non-boot-before-active`, contradicting locked design §12 ("once active, ordinary forms"; "already-active boot-shaped = ordinary accept"). Also ensure the boot rows don't render on active forms (the §12 lifecycle-gated render — a `visible_when` on the two rows or the render-env gate, per the F-S6-M2-2 idiom). Red-first: an active seat's submit carrying boot fields + one ordinary header ⇒ ACCEPTED.
3. **The enum floor over REAL bytes** (`sweep_test.go:74-90`): the current test greps synthetic literals authored in the test body — planting `bounced` in the real enum would not fail it. Rewrite over the real sources (the `record` delivery-state consts + `registry.json` enums) AND a corpus of real serialized outcomes from a driven store; plant-a-leak proof that the scanner bites.
4. **The project-params/three-verb floor over the REAL surface** (`sweep_test.go:115`): the current test compares a literal to itself. Rewrite to drive the real `project` handler with `{"view":"audit"}` and `{"view":"roster"}` through the served socket (or fold the assertions into the live lifecycle fixture and delete the vacuous test).
5. **The missing named negatives + under-tested legs:** `TestTagNeverInAcceptedRecords` (grep committed record bytes for the tag field ⇒ zero — R1-M1-3's negative) · `TestStaleNonSubmitRefusalIsNotLifecycleGating` (current-credential minted-not-active seat retains read/project; only the superseded credential's calls refuse — R1-M1-2) · FX-B1e multi-field smuggle legs (boot set + one REGISTERED extra, + one UNREGISTERED ⇒ per-field `<field>:non-boot-before-active` detail for EACH) · FX-B1f exactly-once unit legs in a NEW `internal/tables/generation_test.go` (second accepted record leaves `ActivationRecordRef` unchanged; already-active boot = ordinary, no second edge). All red-first.
6. **FX-A4b's takeover side** (`s6_lock_test.go` has only the refusal legs): a real-child kill-9-then-TAKEOVER leg for the STORE root lock (acquire after holder death, full recovery, the auditable store-visible takeover diagnostics record asserted) + the alias-path leg (symlinked root ⇒ one winner) + the s4-leftover scenario if not subsumed. Red-first where instrumentable.
7. **The report-accuracy correction** (absorption ruling condition 4): your next substantive report states the gc-file scope fact accurately — in-plan, out-of-row, escalated-and-absorbed citing `RECONCILE-orchestrator-planner-20260707-040738.md`.

**HELD — not in this fold:** the re-mint crash-window production fix (panel trust-1 ≡ design-1; recovery can't distinguish rotated-vs-pending, old credential resolves post-crash). Its fix shape crosses m-1's binding-table hard-stop and is routed via the orchestrator (`SITREP-planner-20260707-041709.md`). Do NOT touch `completeMissingSeatMintBindings`/`completeSeatMintBinding`/`seat.Manager` shapes for this item until the ruling lands; it folds under its own FOLD_SCOPE then, with the SIGKILL-between-pivot-and-rotation fixture red-first.

**OPTIONAL (your discretion, protocol-standard):** gate the shim reconnect-retry on an idempotent-verb set or assert it (mcp.go:172-186) · the `hydrateContentHashes` torn-tail/authoritative-map invariant comment · a checked-in miniature dogfood pattern so one G-1 leg runs unconditionally in CI · feed real parent-stamp/roster bytes into the I-PH scanner (replacing the two literal-fed legs) · assert pre-retraction-accepts-stand explicitly in the waiver fixture · strengthen the thin concurrent-race fixture (distinct IntakeIDs; commit the racing parent) · rename/comment the shim's repurposed `submitNeedsReRender`/`reRenderResult` so the retired readback pattern isn't reintroduced · consider hoisting the recovery closure's inline generation check into the shared handler (the §4 locus principle — currently mirrored, drift risk).

**Discipline:** FOLD_SCOPE pre-filed (artifact BEFORE action) listing every file against this findings scope; `-> in` only for files named by the must-fixes or accepted optionals; ANY out row stops the fold — relay and wait. Rows touching `internal/gc/gc_test.go` cite the absorption ruling. Battery + vet green before each fold commit; `-race` on engine/tables/channel/store. Fold report back under this DISPATCH_ID with FOLD_SCOPE above ACTIONS_GIT_REF, red-first evidence per fixture, and the corrected scope statement (item 7).

FOLD_SCOPE:
- none -> in
FOLD_SCOPE_RESULT: all-in

The scope block above is the structural no-edit form: this relay is the findings DIRECTIVE and touches no file; the executing fold's real FOLD_SCOPE is pre-filed by the Implementer before its first edit, per the Discipline section.

ACTIONS_GIT_REF: none — fold directive only; no edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree (main@2903d84; impl worktree clean at bfbbb2d).
Next requested action: s6-core.implementer pre-files FOLD_SCOPE and folds items 1–7 + accepted optionals; the held item folds after the m-1 ruling routes back.
