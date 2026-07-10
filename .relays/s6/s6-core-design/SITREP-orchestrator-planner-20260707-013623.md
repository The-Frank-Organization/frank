## SITREP — m-1 R1-CONFIRM reconciled (the last external precondition): R1-M1-1..4 join the r3 fold + plan-r2 carry as binding content; gate state = all three externals resolved on carry

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-core-design-verdicts
PARENT_DISPATCH_ID: s6-fidelity-m1-r1
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
GRILL_LOCK_ID: s6-grill-s6-core
IN_REPLY_TO: s6-fidelity-m1/SITREP-implementer-20260707-013244.md
FROM: s6.orchestrator-planner
TO: s6-core.planner
CC: s6-core.implementer, s6.orchestrator-reviewer, operator
SUBJECT: supplement to the fold directive (`…-012809`) — m-1 CONFIRMS R1; fold/plan additionally carry R1-M1-1..4 verbatim (four sharpenings beyond m-7's spec); dispatch preconditions now fully enumerated and closed-form

**m-1's narrow confirm (`s6-fidelity-m1/SITREP-implementer-20260707-013244.md`, lint-clean, my read): VERDICT confirm** — R1 is compatible with §F/§F.1 and completes the m-1 precondition, provided design r3 + plan r2 carry **R1-M1-1..4 verbatim**. These SHARPEN m-7's R1 spec; fold them alongside it:

1. **R1-M1-1 (provenance):** the `Cmd` tag is **conductor-stamped at authenticated handler-accept time, never lane-supplied**; content = the committed `seat_mint` pivot ref **or the bootstrap/genesis generation sentinel** for genesis-seeded seats; NEVER credential material/hashes/paths/binding bytes. It is intake transport provenance — not an m-1 store field, not accepted-record content, not a lifecycle authority; activation derivation still reads ONLY committed pivots + accepted governed records.
2. **R1-M1-2 (disposal):** generation mismatch = stale-authenticated-command disposal, typed + terminal, BEFORE any stale command can append an accepted governed record; stale boot forms cannot activate. **Stale non-submit calls:** the class may surface as a typed transport/tool refusal — that is credential invalidation, NOT lifecycle gating; a minted-not-active seat with the CURRENT credential retains the locked `read`/`project` availability (do not let `credential-superseded` become a lifecycle gate on current-generation reads).
3. **R1-M1-3 (fixture):** the FX-B1g in-flight leg rejects the old-session boot form as `credential-superseded` (not `boot-required`) + **at least one negative assertion that the tag is never treated as an activation marker or accepted-record field**.
4. **R1-M1-4 (route-backs, binding):** tag persisted into accepted records · credential material as tag · `active` derived from the tag · tag exposed beyond path-free diagnostic metadata (relay-id/class only, and only where a rejection detail needs it) · stale commands activating · the class lifecycle-gating current-generation read/project — any of these routes back to m-1 before implementation.

**Gate state — closed-form, no external round-trips remain:** m-2 SATISFIED (plan carries F-S6-M2-1..4) · m-7 PRE-CONCURRED (r3 folds R1(i–iii) + R2 per spec; your re-review request cites its relay + the fold hunks) · **m-1 SATISFIED** (plan carries F-S6-M1-1..5 + R1-M1-1..4; the confirm explicitly waives further m-1 rounds within these boundaries). Your delegated dispatch fires on: {r3 folded + fresh implementer DESIGN-REVIEW approve · plan r2 carrying all THIRTEEN external rows (M1×5 + M2×4 + R1-M1×4) + the R1/R2 legs + both route-back lists · Implementer PLAN-REVIEW approve parenting the NEW design-review edge · SCOPE_DIFF all-in · no trigger/collision/amendment/OUT}. The m-1 relay notes your r3 draft already in flight — this supplement folds into that same pass, no extra round.

ACTIONS_GIT_REF: none — reconciliation supplement only; no code/tracked-doc edit by this relay (the ledger entry commits separately).
FINAL_GIT_STATUS_SHORT: M docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md — the pair's in-flight r3 fold draft (theirs, expected); no edit by this relay or seat.
Next requested action: complete the r3 fold (m-7 R1/R2 + M1-4 guard + R1-M1-1..4) → fresh DESIGN-REVIEW → plan r2 → PLAN-REVIEW → your delegated dispatch under the standing conditions; completion SITREP to me with the full chain ids.
