## REVIEW-FOLD — round 3 (sanctioned addition): the `-mint <seat>` admin-time flag, one bounded commit under the orchestrator's six conditions

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-r3
PARENT_DISPATCH_ID: s2-core-review-r2
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the 13.5 submit stays operator-executed; merge/close stay the operator's
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/SITREP-orchestrator-planner-20260704-143000.md
FROM: s2-core.planner
TO: s2-core.implementer
CC: s2.orchestrator-planner, operator
SUBJECT: fold-sized addition SANCTIONED by `SITREP-orchestrator-planner-20260704-143000.md` (option (a); option (b) rejected on D5-normalization grounds) — implement `-mint <seat>` per the six binding conditions below, red-first, FOLD_SCOPE pre-filed; report back for my targeted check

**Authority trail:** the Task-13.5 executability gap I routed (`SITREP-planner-20260704-141200.md`) → the orchestrator's sanction relay (parent of this one) rules option (a) IN with binding conditions. cmd/frank/main.go is in-fence (plan file list); the mint MECHANISM is S1-built + m-1-approved — this exposes it through conductor-internal tooling only. Ride the round-1/2 fold discipline; no new panel; no fresh dispatch token needed (standing s2-core-impl dispatch + this sanctioned fold).

**The six binding conditions (verbatim from the sanction; each is a check in my verification):**
1. ONE fold-sized commit on `s2-core-impl`, red-first fixture; FOLD_SCOPE artifact pre-filed naming exactly the files (cmd/frank/main.go + its test surface; any NEW file cites the sanction relay `SITREP-orchestrator-planner-20260704-143000.md` as its fence-amendment evidence).
2. **Admin-time posture:** `-mint <seat>` runs against a store root while the conductor is NOT serving (the `-init` operational class); never a seat-facing verb; never in any rendered tool registry (assert: registries unchanged in both surfaces).
3. **Credential handling:** printed ONCE to the invoking admin's stdout (the operator's DI-2 provisioning channel; not a seat-delivered surface, so I-PH's seat-grain holds — state this grain in the fixture); never persisted outside the binding table; never in any seat-deliverable string. Re-run the P1/I-PH + SWEEP legs in the same commit.
4. **Fixture invariants:** `-mint system` ⇒ typed `ErrReservedSeatName`, nonzero exit, binding table unchanged; re-mint of a bound seat ⇒ typed `ErrSeatAlreadyBound`, nonzero exit, table byte-identical (the S1 single-generation invariant at the CLI grain); a successful mint's credential resolves on connect (drive an authenticated call with it).
5. **No binding-table SHAPE change.** If the implementation seems to need one: STOP, relay to me — that re-opens the m-1 surface (as specced it does not).
6. Report back TO me (fold report shape: FOLD_SCOPE above ACTIONS_GIT_REF, red-first evidence, your verification runs); my targeted check + battery at the new head closes the round.

**Suggested flag semantics (yours to refine within the conditions):** `frank -root <store> -mint <seat> [-role <role>] [-operator]` — errors if the socket is live (admin-time check), prints `credential=<hex>` once, exits 0. Roles/IsOperator per existing `seat.Mint` signature.

After your fold lands green: I prepare the exact operator command lines + the OI-S1-F11-SWEEP payload file (guide-confirmed shape) for the operator's 13.5 execution; then the exit-gate pass.

ACTIONS_GIT_REF: none — instruction relay only; no source edits by me; this relay + its INDEX row under gitignored .relays/ (`.gitignore:1`)
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout; s2-core-impl worktree clean at 16342e0)
