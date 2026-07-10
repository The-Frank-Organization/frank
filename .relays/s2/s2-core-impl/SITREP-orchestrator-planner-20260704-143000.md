## SITREP — Task-13.5 sanction: option (a) SANCTIONED, bounded — one fold-sized `-mint` commit under the conditions below; option (b) REJECTED (normalizes a tool-surface bypass); merge posture noted for the close

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-core-impl
PARENT_DISPATCH_ID: s2-core-plan-review-implementer-r6
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this sanction is the orchestrator's scope call; the 13.5 SUBMIT itself and the S2-close/merge remain the operator's (restated below)
IN_REPLY_TO: s2-core-impl/SITREP-planner-20260704-141200.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: sanction (a) — a conductor-internal `-mint <seat>` init/admin-time flag on cmd/frank, one bounded commit + pair verification, fence amended by this relay; (b) rejected on D5-normalization grounds; my own E2 at 16342e0 on record

**Reconciled first (my own verification this session):** delegated dispatch lineage clean (`IMPL-planner-20260704-122950.md` — parent `s2-core-plan-review-implementer-r6`, all four standing conditions restated with evidence, SCOPE_DIFF all-in, valid bare token); 12 task commits + 2 scope-matched fold commits; battery at s2-core-impl@16342e0 — **18 packages ok (uncached), vet clean, race pass green on intake/engine/obligation — my runs, E2**; review trail (2 rounds, FOLD_SCOPE pre-filed both) consistent with your report.

**Sanction ruling — (a), SANCTIONED as a bounded post-review addition:**
Grounds: cmd/frank is an in-fence file (plan file list, Task 11); the mint MECHANISM is S1-built and m-1-fidelity-approved (single-generation invariant; now + `ErrReservedSeatName`); the locked m-1 line anticipates exactly this surface — "mint_seat is conductor-internal; operator/spawn tooling asks the conductor to mint" — an init/admin-time flag on the conductor's own binary IS that realization, no store-shape or API-semantics change. **(b) is REJECTED:** documenting an out-of-binary `seat.Manager` drive as the sanctioned recipe normalizes a tool-surface-bypass idiom — the exact D5-shaped path this design spends its claim boundary refusing to normalize.

**Binding conditions on the (a) commit (each a check in your targeted verification):**
1. **One fold-sized commit** on `s2-core-impl`, implementer-authored, red-first fixture; scope artifact pre-filed naming exactly the files (cmd/frank/main.go + its test surface; any NEW file cites THIS relay as its fence-amendment evidence, the README-row pattern).
2. **Admin-time posture:** the flag runs against a store root while the conductor is NOT serving (same operational class as `-init`); it is never a seat-facing verb, never in any rendered tool registry.
3. **Credential handling:** printed ONCE to the invoking admin stdout (the operator's DI-2 provisioning channel — m-1 §6; admin stdout of the conductor's own init tooling is not a seat-delivered surface, so I-PH's seat-grain holds) — never persisted anywhere but the binding table, never in any seat-deliverable string. State this grain in the fixture, and re-run the P1/I-PH + SWEEP legs in the same commit.
4. **Invariants asserted in the fixture:** `-mint system` ⇒ typed `ErrReservedSeatName`, nonzero exit, binding table unchanged; re-mint of a bound seat ⇒ typed `ErrSeatAlreadyBound`, nonzero exit, table byte-identical (the S1 single-generation invariant at the CLI grain); a successful mint resolves on connect.
5. **No binding-table SHAPE change.** If the implementation needs one, STOP and escalate to me before the commit — that would re-open the m-1 fidelity surface. (As specced above it does not; no m-1 re-review is required for exposing the approved mechanism through conductor-internal tooling.)
6. Your targeted planner check + battery green at the new head close it — the round-1/round-2 fold discipline, no new panel.

**Then the 13.5 sequence stands as you laid it out:** prepare the exact operator command lines + the OI-S1-F11-SWEEP payload file (source citing s1 RECONCILE :160-161 + the guide ruling relay, per the guide-confirmed shape) → the OPERATOR executes mint + submit against the S2 acceptance store → record + post-gate-run disposition copies into results/ → your exit-gate pass + gate report → master SITREP. The submit is operator-executed, never pair-authored (unchanged).

**Merge posture, recorded for the close:** noted and agreed — S2-close will involve the first REAL `git merge s2-core-impl → main`. That is the operator's separate gate: a `PHASE: MERGE-GATE` relay from this seat carrying the mechanical truth (branch, head SHA, base, diff-stat) TO the operator; authorization only by the operator's token/field-form grant. Nothing about green fixtures implies it; the close relay will carry it cleanly.

ACTIONS_GIT_REF: none — sanction/report relay; this file + an INDEX row under gitignored .relays/ (ledger entry rides its own commit on main, cited in RECONCILE.md). No branch edits by this seat.
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout; s2-core-impl worktree clean at 16342e0, verified this session)
