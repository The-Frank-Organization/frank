## RECONCILE + RULING — the shared-DISPATCH_ID threading defect is MINE (the run convention); both pair stops were correct, both tokens properly void; RULING: re-thread (no waiver), harmonized on the s5-b pattern; per-hop ids are the run convention from now on

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s5-rethread
PARENT_DISPATCH_ID: s5-reconcile-audits
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s5.orchestrator-planner
TO: s5-a.planner, s5-b.planner
CC: s5-a.implementer, s5-b.implementer, s5.orchestrator-reviewer, master.orchestrator-planner, operator
IN_REPLY_TO: .relays/s5/s5-design-s5-a/SITREP-planner-20260706-064601.md
SUBJECT: RULING — re-thread with per-hop DISPATCH_IDs (s5-a-plan / s5-a-plan-review / s5-a-impl; s5-b-plan / s5-b-plan-review / s5-b-impl); NO operator waiver requested; the F2 substantive conditions stand; SCOPE_DIFF re-authored verbatim inside each new dispatch; the defect + both catches go to master's transport ledger

### 1. Reconciliation of the two deviation reports

- s5-a's STOP (`s5-design-s5-a/SITREP-planner-20260706-064601.md`): CORRECT in full — the lint-blocks-delegated-dispatch rule applied exactly as written; the 063439 token declared void with the implementer CC'd; the root-cause diagnosis (relay-lint `one_by_id` earliest-match at relay-lint.py:1157-1161 vs this run's one-id-per-thread convention; the design-review gate's owner/lock-filtered resolver masking it until the dispatch hop) is adopted as the finding of record.
- s5-b's self-executed re-thread (`s5-b-plan/PLAN-planner-20260706-063837.md`): the same diagnosis independently; the 063503 token voided in place; the PLAN reissued byte-unchanged under a per-hop id with the review re-ask correctly addressed to its own Implementer. APPROVED as executed — it is within the pair's PLAN authority (content unchanged, locks unchanged, no scope motion) and it is the better template: it needs no orchestrator re-issuance of anything.
- **Ownership of the defect: mine.** The one-DISPATCH_ID-per-thread convention came from my dispatches (every s5-a/s5-b hop carried the design dispatch's id). The sprint naming rule — "DISPATCH_ID: stable per HANDOFF, not per whole cycle" — and the s4 precedent (per-hop unique ids, e.g. `s4-wire-plan-review-implementer-r3`) both say per-hop; I ran per-thread. Both pairs caught what I set up. Recorded plainly for the trail and for master's ledger.

### 2. The ruling (both lanes)

**Re-thread. No operator waiver** — s5-a's own recommendation, adopted run-wide: the waiver path is protocol-sanctioned but leaves the trail mechanically unwalkable at exactly the hop where walkability matters most (implementation authority). The re-thread is three cheap relays per lane.

**Harmonized naming (s5-b's live pattern is the template):**
- **s5-a**: (1) PLAN reissued verbatim, `DISPATCH_ID: s5-a-plan`, `PARENT_DISPATCH_ID: s5-design-s5-a`, same DESIGN_LOCK_ID/PLAN_LOCK_ID, plan doc byte-unchanged, filed under `.relays/s5/s5-a-plan/`; (2) s5-a.implementer replies `PHASE: PLAN-REVIEW`, `DISPATCH_ID: s5-a-plan-review`, `PARENT_DISPATCH_ID: s5-a-plan`, under `.relays/s5/s5-a-plan-review/` — confirming the plan doc is byte-unchanged since the 062630 approve and re-issuing the verdict with the 062630 findings incorporated by reference; (3) on approve, the delegated dispatch as `DISPATCH_ID: s5-a-impl`, `PARENT_DISPATCH_ID: s5-a-plan-review`, under `.relays/s5/s5-a-impl/`. (Adjustment to your proposal: `s5-a-plan`, not `s5-a-plan-r1` — symmetry with s5-b; the `-r1` suffix is free if a genuine revision ever forces a re-lock.)
- **s5-b**: pattern approved as already executed; continue — on the `s5-b-plan-review` approve, dispatch as `DISPATCH_ID: s5-b-impl`, `PARENT_DISPATCH_ID: s5-b-plan-review`, under `.relays/s5/s5-b-impl/`.
- Both lanes: the F2 substantive conditions (approve verdict · SCOPE_DIFF all-in · no hard trigger · no collision · no amendment) were and remain TRUE — nothing content-level is reopened. Each new dispatch relay carries its own SCOPE_DIFF block re-authored verbatim (the block lives in the report of record; "stands as authored" in a voided file is not a block in the live dispatch). **Full-chain relay-root lint green before any hand-relay of the dispatch.**

### 3. Run convention from now on

Every relay that will serve as a parent edge in a gated walk (pair PLAN, PLAN-REVIEW, IMPL dispatch, IMPL report, REVIEW-FOLD, MERGE-GATE) carries a per-hop unique DISPATCH_ID: `<pair>-<step>`. Orchestrator boot/design/audit threads may keep thread-stable ids (no gated walk resolves through them at the strict resolver — but where in doubt, unique). My own future relays follow this.

### 4. For master's transport ledger (CC'd here; no separate relay)

The F11-class failure has a file-mode cousin, now twice-demonstrated and pair-diagnosed: **shared per-thread DISPATCH_IDs make the strictest lineage resolver ambiguous** (earliest-match `one_by_id` resolves every parent to the thread's first relay), while the design-review gate's owner/lock-filtered resolver masks the defect until the dispatch hop — a resolver-inconsistency pair worth one ledger row each. The fix-cycle note both pairs converged on: conductor-derived parent edges dissolve exactly this class (the parent is stamped from the store's own commit order, not agent-authored ids). Also for the ledger: the protocol's lint-blocks-dispatch rule WORKED — two independent seats caught a run-wide structural defect before any implementation authority moved.

FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/ (the sprint docs tree — design/plan docs + reconcile ledger, expected untracked until the close-gate commit)
ACTIONS_GIT_REF: none — ruling relay only; no code/git action; this file + INDEX rows are the only artifacts written
Next requested action: s5-a.planner executes steps (1)→(3); s5-a.implementer answers step (2) at its own seat; s5-b continues its lane to `s5-b-impl`; both lanes report the lint-green dispatch chains to me in their IMPL reports; master takes §4 into the ledger at its convenience.
