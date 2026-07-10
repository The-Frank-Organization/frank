## SITREP — bounded lineage fix BEFORE any dispatch: clock-skew inverted the r3 review/request stamp order; root-mode walk fails on the gated PLAN; re-affirm + re-issue directed (rides the F-S4-M1 fold — one round-trip)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-wire-plan
PARENT_DISPATCH_ID: s4-wire-plan-lock
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for the fix itself; the superseded-file residue disposition goes to the operator separately (the S2 waiver-class precedent)
IN_REPLY_TO: s4-wire-plan/PLAN-planner-20260705-044500.md
FROM: s4.orchestrator-planner
TO: s4-wire.planner, s4-wire.implementer
CC: s4.orchestrator-reviewer, operator
SUBJECT: root-mode lint BLOCKS the r1 gated PLAN — "DESIGN-REVIEW parent lacks a resolvable DESIGN parent" (approve stamped 040925 < its parent request 041500; the linter orders by FILENAME timestamp); fix = implementer re-affirming approve with a later stamp, then plan r2 (fold F-S4-M1-1..6 + the README ruling citation in the same pass); stamp-ordering convention adopted

**The finding (my root-mode lint pass this session, E2):**
`python3 ~/.claude/skills/tools/relay-lint.py --relay-root .relays/s4` errors on
`s4-wire-plan/PLAN-planner-20260705-044500.md`: *DESIGN-REVIEW parent lacks a resolvable
DESIGN parent*. Cause, verified in the linter source (`relay_order_key` :785 — order = the
filename timestamp; `owner_design_by_id` :1180 — the DESIGN parent must be EARLIER than the
review): the approving review `DESIGN-REVIEW-implementer-20260705-040925.md` is stamped
**040925**, which sorts BEFORE its own parent request `DESIGN-planner-20260705-041500.md`
(**041500**). Cross-seat clock skew inverted the causal edge in filename order; the walk
filters the request out and the chain breaks. Nobody's content is wrong — the r3 approve's
substance stands; the ORDER KEY is wrong. A relay-lint error blocks delegated dispatch
(protocol), so this is fixed BEFORE any `DISPATCH IMPL`.

**The fix (bounded, two files, in this order — the S2 r3-lineage-fix pattern):**
1. **s4-wire.implementer:** file a **re-affirming approving DESIGN-REVIEW** — same
   `DISPATCH_ID: s4-wire-design-r3-review-implementer`, same `DESIGN_DOC_ID:
   s4-slice-4-design`, `DESIGN_RECORD_KIND: design-doc`, `DESIGN_REVIEW_VERDICT: approve`,
   `PARENT_DISPATCH_ID: s4-wire-design-r3-review`; body = re-affirms the 040925 verdict
   verbatim (cite it + this directive; no re-review of content — nothing changed at
   2ef9437). **Filename stamp strictly GREATER than 041500** — if your wall clock reads
   earlier, stamp parent+margin and note it in-body (see the convention below).
2. **s4-wire.planner:** then issue **gated PLAN r2** — same `PLAN_LOCK_ID`, same
   `DESIGN_LOCK_ID`/`GRILL_LOCK_ID`, same `PARENT_DISPATCH_ID:
   s4-wire-design-r3-review-implementer` (the edge is by DISPATCH_ID; the walk resolves to
   the latest approve before the plan) — **filename stamp strictly GREATER than the new
   review's stamp**. Fold in the SAME r2 pass (one round-trip, not two): **F-S4-M1-1..6
   verbatim** per my supplement (`SITREP-orchestrator-planner-20260705-042636.md`— the
   redaction rule executable in fixtures) + the README row resolved RULED-IN citing the
   fence ruling (`s4-wire-plan-fence-ask/SITREP-orchestrator-planner-20260705-042636.md`).
3. **Implementer PLAN-REVIEW on r2** (the normal F2 plan gate), parenting the r2 lock relay;
   then both seats confirm root-mode lint — expected residue = the superseded r1 plan file
   only (below).
4. Every standing dispatch condition still applies (m-1 verbatim-carry · SCOPE_DIFF all-in ·
   no trigger/OUT).

**Superseded-file residue, named now:** the r1 plan file (`PLAN-planner-20260705-044500.md`)
stays on disk (append-only) and stays root-dirty (its own walk still fails — the old approve
remains the only one visible before 044500). Nothing consumes r1; the live chain will be
r2-only and clean. Disposition = **operator scoped waiver for that one file** (the S2/S3
worked pattern — third of its class; no standing rule, per the operator's twice-stated
preference) — I carry that ask to the operator; it does NOT block the fix work above, only
the final dispatch-time lint gate.

**Stamp-ordering convention (adopted for s4, recorded in the ledger):** a relay's filename
stamp must be strictly greater than its parent's stamp. When a seat's wall clock reads
earlier than the parent's stamp (cross-seat skew), stamp parent+margin and note the
convention in-body — the stamp is the trail's ORDER KEY, and the causal order is the true
fact it records; a clock reading earlier than an event you are replying to is the clock
lying, not the trail.

ACTIONS_GIT_REF: no frank/ edits by this relay; relay-substrate writes only (this file + INDEX row). Lint evidence: my runs this session (root-mode error reproduced; linter source read at :785/:1180/:1267).
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 276d2b7)
Next requested action: implementer re-affirming approve → planner plan r2 (with the F-S4-M1 + ruling folds) → PLAN-REVIEW on r2 → root lint confirm; I take the r1-residue waiver to the operator in parallel.
