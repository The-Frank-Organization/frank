## DESIGN-REVIEW - m-6 implementer re-review after Seam C lock-status correction

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c3-design-m-6
PARENT_DISPATCH_ID: c3-design-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - re-review verdict; no implementation authority
GRILL_REQUIRED: yes - prior folded GRILL_LOCK unchanged
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_REVIEW_VERDICT: approve
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c3-design-m-6/DESIGN-planner-20260630-184921.md
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)
SUBJECT: DESIGN-REVIEW approve - post-Seam-C correction clears stale held/pending status; m-6 design remains lock-ready

Verdict: **approve** for `DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler`.

I re-reviewed the current m-6 design after the VP's c3-lock revise (`c3-lock/RECONCILE-orchestrator-reviewer-20260630-184253.md`) and the m-6 correction confirm (`c3-design-m-6/DESIGN-planner-20260630-184921.md`). The prior m-6 implementer approval (`133839`) still stands, and the narrow post-approval Seam C correction is now internally consistent.

## Findings

No blocking findings.

## Re-review Evidence

1. **VP stale-status blocker is cleared.**

   The current design top status now says Seam C is resolved A, the away-token cell is folded over m-1-owned mint/verify, and all cells are closed; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:6`.

   The previously stale "held/pending" locations are corrected:
   - §8 now says m-6 calls m-1-owned `mint`/`verify` and no binding is left without a settled upstream; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:128-132`.
   - §10 now marks OQ-1 resolved and says the away-token cell locks; `:147-151`.
   - §11 now says no held cell remains and includes the away-mode token bridge in the locked set; `:155-159`.
   - §12 now describes the signed-token bridge as a build carry with Seam C = A, not as pending OQ-1; `:161-173`.

   A grep sweep found no live stale strings for `pending OQ`, `(pending)`, `Held:`, `held cell`, `does NOT design-lock`, `stays unlocked`, `LOCK-BLOCKING`, or `without a settled upstream` in the current m-6 doc.

2. **Seam C = A stays inside the approved boundary.**

   The m-1 answer says A is forced by DI-1 nonce-burn / sole-writer-store and DI-2 signing-key custody, and defines m-1 as owner of `mint`/`verify` while m-6 owns the bridge around it; `master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-182600.md`. The orchestrator accepted that boundary and routed only a bounded fold to m-6; `master/relays/c3-design-m5-m6-coord/COORD-orchestrator-planner-20260630-183008.md`.

   The m-6 design folds the same boundary without re-owning TCB work: m-1 owns `mint(decision_id, seat, choice, expiry)` and `verify(token)`; m-6 owns the trigger, token-bearing render, POST receipt, routing, UX, and supplies expiry; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:80-88`.

3. **Earlier approved seam and egress fixes remain intact.**

   The m-5/m-6 seam remains the four-class non-gate model from `123022` plus `131856`, with gate-bearing records carrying no `surface_intent`; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:111-124`.

   The egress A/D precedence rule remains explicit by `failing_edge` and pipeline stage; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:31-49`.

## Scope Of This Approval

This re-review approves the current m-6 design doc after the Seam C correction. It is **not** a c3-lock co-sign, does **not** seal `ARCHITECTURE.md`, and does **not** authorize PLAN/IMPL. The c3-lock relay still needs orchestrator re-emission and VP co-sign because the latest visible c3-lock review remains `revise` pending the correction.

ACTIONS_GIT_REF: wrote this read-only design-review relay only: `master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-185424.md` and appended the matching `master/relays/INDEX.md` row; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-185424.md` returned `OK`.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m-6/DESIGN-REVIEW-implementer-20260630-185424.md` returned `OK` for this relay and expected root noise for lint-exempt `INDEX.md`.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m-6/DESIGN-planner-20260630-184921.md` returned `OK` for the m-6 correction-confirm relay and expected root noise for lint-exempt `INDEX.md`.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-design-m5-m6-coord/COORD-planner-20260630-183345.md` returned `OK` for the Seam C fold-confirm relay and expected root noise for lint-exempt `INDEX.md`.
- E1: `find master/relays/c3-lock -maxdepth 1 -type f -name '*.md' -print | sort` shows no re-emitted c3-lock relay after VP revise `184253`; latest c3-lock remains pending re-emission/co-sign.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
