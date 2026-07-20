## REVIEW-FOLD — the ONE end-of-slice adversarial review (B11): code APPROVE with ZERO blockers/must-fix code defects across my verification + two read-only reviewer lenses; ONE evidence-only must-fix (complete the diff→license table), ONE optional routed to the guide, two stated residuals and one named carry recorded; the merge decision additionally holds on master's fence-row grant + the batched owner confirms

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s9-build-slice-review
PARENT_DISPATCH_ID: s9-build-review-request
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair-internal review/fold; merge remains operator-only via HUMAN_MERGE_AUTHORIZATION
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev13
IN_REPLY_TO: master/relays/s9-build-review-request/SITREP-implementer-20260714-011429.md
FROM: s9.planner
TO: s9.implementer
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: reviewed `39474d0..76179ec` (8 commits, 29 files, +2255/−281) — the build conforms to the rev13 plan-of-record at line grain and no adversarial lens found a reachable defect in the new machinery; fold MF-1 (evidence table totality, one file, no code) and return; AO-1 goes to m-3 as a guide byte, NOT to your editor

s9.implementer — the single B11 end-of-slice review is complete. Method: my own verification (full uncached battery re-run at `76179ec`; owner-byte isolation diffs at all four loci; executor/fieldspec zero-diff proof; per-commit battery evidence walk; the diff→license reconciliation below) plus two read-only reviewer lenses (the charter's one sanctioned spawn): a plan-fidelity lens over every T1–T5 mechanism at file:line grain, and an adversarial lens over confinement/TOCTOU, serialized-path syscalls, breaker races, binding-pass bypass, redaction, opaque-accept reachability, incomplete-scan honesty, and test-fake contract-weakening.

### Verdict

**Code: APPROVE — zero blockers, zero must-fix code defects.** Confinement (descriptor-rooted, per-component O_NOFOLLOW, no fd leaks), the breaker (set-before-return on every kill path, no double-launch), the binding pass (origin conductor-stamped, identity from Selection, forged rung faults, no gate-consumed verdict bypasses `validateBoundVerdict`), redaction (FailingDetail/Timing discarded; rows/bounces/degradation_notes carry constants and symbols only), the opaque-accept branch (reachable ONLY via governed `lane_vcs: none`; no candidate string requests no-vantage), and the rev13 base-refusal row (allowlist closed at three, machinery boundary intact, per-detail timing) all verified sound. The terminal enum is byte-exact `{accepted, rejected, held}` (`internal/record/record.go:11-13`); the marker sweep is zero-hit; T7/T8 and the blocked ledger have no code shadow.

### MF-1 (must-fix, EVIDENCE file only — the one fold)

`.relays/s9/mechanical-tables.md`'s diff→license table does not reconcile the full changed-path set: it omits the four `db9a166` owner-byte paths (`internal/config/config.go`, `cmd/frank/main.go`, `test/fixtures/s2setup_test.go`, `test/invariants/store_recovery_test.go` — license: m-7 owner bytes carried verbatim under the master activation `…-194510`), omits the `.relays/s9/` evidence files themselves, and marks `internal/config/lane_vcs_test.go` "in" when it — like `test/invariants/store_recovery_test.go` — is **outside the tokenized 10-row block**. The table must cover EVERY path in `git diff 39474d0..76179ec --stat` against the BLOCK, with those two rows marked **OUT — escalated to master** (my escalation rides beside this relay under `s9-build-escalate-fence`; both edits are substantively justified owner-directed loci, but an OUT row is master's call, not the pair's — and my own `…-190901` routing relay wrongly asserted "both in-fence", owned there). Fold = amend the table; no source/test byte moves.

### AO-1 (optional — routed to the GUIDE, do not self-implement)

The plan's T2 I-PH surface contract says the find-references verdict carries "the bounded `count` (+ a saturation flag)"; the landed verdict surfaces neither (pass/fail derived internally — strictly MORE conservative than I-PH, nothing leaks, but a consumer cannot distinguish count=1 from count=1000 on a fail). Whether to deliver the bounded scalar now or record it as a named carry is m-3's surface byte — requested in the batched fidelity relay. Hold your editor until the guide answers.

### Recorded, no action by you

- **N-1:** T3's `failing_detail` handling is validate-and-fault against the closed origin-family allowlist rather than literally discard-and-re-derive — equivalent guarantee (every allowlist is a closed set of conductor-known tokens; any foreign string faults typed via origin-class-mismatch and reaches no surface). Conforms-with-note.
- **N-2:** the find-references binary exclusion (NUL in first 8 KiB) is the plan's DECLARED domain scope with the honest lexical-not-semantic label — a UTF-16 file carrying the token is outside the declared textual domain by design, not a silent skip. No deviation.
- **R-1 (stated residual, D5-class):** `O_NOFOLLOW` confinement does not close hardlinks (an in-root hardlink to an out-of-tree same-filesystem inode would be read). Requires prior write access to the governed tree; negligible under the confusion-not-malice threat model; LABELED here rather than engineered away, per the ratified directive.
- **C-1 (named carry, to the m-3/m-7 ledger):** the pre-existing `git status` exec on the serialized commit path (`checks_base.go` §13 machinery — synchronous `exec.CommandContext` + PATH lookup on the loop). NOT an s9 defect: the branch-only ruling required §13 byte-preservation, and detaching it was never licensed. It is now the ONE remaining un-detached path the T1 worker did not close — named so it stops being invisible.

### Verification (mine, this session)

- `go test ./... -count=1` at `76179ec`: green, uncached, all packages (`test/fixtures` 136.074s).
- Owner-byte isolation: `db9a166` touches exactly the four m-7 paths; `main.go` = the one-line E9; both descriptor constructors = exactly E10 (+ the FX-VCS-6-required derived v2 helper in `test/fixtures/`).
- `git diff 39474d0..76179ec -- internal/executor/ internal/fieldspec/` = empty.
- Per-task RED/GREEN batteries + per-commit full-battery captures present and sequence-honest under `.relays/s9/`; the T3 battery pins the rev13 sha + guide-byte relay.
- rev13's bounded re-review APPROVE in hand (`s9-baserefusal-m3/DESIGN-REVIEW-implementer-20260714-011500`); the current plan sha `d16352a15a6164244beb53d99becf0f1982dfb3af50503c43047cd827dedef50` differs from rev13-at-return only by that fold-log entry (verified).
- Lens reports retained in my session; their dispositive claims re-checked at the cited lines before adoption.

FOLD_SCOPE:
- frank/.relays/s9/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: none — a review relay; no `frank/` edit by this seat. Artifacts this turn: this relay + the master fence escalation (`s9-build-escalate-fence/SITREP-planner-20260714-013001`) + the two batched owner-confirm requests (`s9-build-fidelity-m7/…-013002`, `s9-build-fidelity-m3/…-013003`) + four `master/relays/INDEX.md` rows.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); the `s9-evidence-thicken` worktree clean at `76179ec`; `frank/` main clean at `39474d0`.
Next requested action: fold MF-1 (the one evidence file) and return your REVIEW-FOLD report; hold AO-1 for m-3's byte; I author the merge-decision relay TO the operator once your fold, master's fence-row ruling, and the two owner confirms are all in with no contradiction. No merge authority is granted or implied.
