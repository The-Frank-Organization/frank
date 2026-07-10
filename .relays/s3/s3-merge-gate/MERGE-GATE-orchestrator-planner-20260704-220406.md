## MERGE-GATE — token relay: bounded four-step S3 integration, executor s3-form.implementer (operator-authorized; VP confirm on record)

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s3-merge-gate
PARENT_DISPATCH_ID: s3-exit-gate
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the human gate is already exercised: the operator's four decisions are on record (quoted below); this relay executes them, bounded, nothing more
FROM: s3.orchestrator-planner
TO: s3-form.implementer
CC: operator, s3-form.planner, s3.orchestrator-reviewer
IN_REPLY_TO: s3-vp-confirm/RECONCILE-orchestrator-reviewer-20260704-215937.md
SUBJECT: execute exactly four bounded steps — verify-then-integrate s3-form-impl@fe7308e into main with --no-ff · post-integration battery before anything else lands · annotated tag s3-close · execution report under this DISPATCH_ID. Conflict or red battery = stop + escalate, never fix-forward

**Authorization chain (all on record):** the operator's four MERGE-GATE decisions (interactive session with this seat, 2026-07-04, selections verbatim in the ledger + the VP request): ① close authorization **"Authorize (Recommended)"** · ② executor **"s3-form.implementer (Recommended)"** — you · ③ tag **"Tag s3-close (Recommended)"** · ④ **"VP pass before merge"** — SATISFIED: `s3-vp-confirm/RECONCILE-orchestrator-reviewer-20260704-215937.md`, VERDICT approve (the fourth independent verification chain; battery/vet/race + centerpiece probes + the r4/r5 judgment independently ratified). This relay is issued from a protocol-sanctioned grantor seat (orchestrator-planner FROM) under the same `DISPATCH_ID: s3-merge-gate` as the decision packet, so your execution claim resolves in root-mode lineage.

**Authorized steps — exactly four, in order:**
1. **Verify-then-integrate:** confirm `main` HEAD = 354718b (or a descendant that is docs-only vs 354718b — verify by diff-stat if moved; any code movement ⇒ STOP + escalate) and `s3-form-impl` head = fe7308ed830e7b11e3ed1b31e694a81d4ae07ad8; then `git merge --no-ff` of `s3-form-impl` into `main`. **Any conflict ⇒ abort the merge + escalate to me; do not resolve in-flight.**
2. **Post-integration battery on the result, before anything else lands:** `go test -count=1 ./...` (expect 20 ok uncached) + `go vet ./...` clean, run on the integration commit. **Red ⇒ STOP + escalate with the output; no fix-forward, no second commit.**
3. **Annotated tag `s3-close`** on the integration commit.
4. **Execution report** under this `DISPATCH_ID` into `.relays/s3/s3-merge-gate/` (FROM your seat): the integration SHA, the battery output summary, the tag object, `ACTIONS_GIT_REF` with the graph facts, lint-clean.

**Not authorized:** any other edit, any history rewrite, any push, any fix-forward on a failed step, any additional commit beyond the integration itself, any tag other than `s3-close`. A failed step leaves the repo as the failure left it and escalates.

DISPATCH MERGE

ACTIONS_GIT_REF: none — token relay only; no git action performed by this relay itself; this file + an INDEX row under gitignored .relays/
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ main at 354718b)
