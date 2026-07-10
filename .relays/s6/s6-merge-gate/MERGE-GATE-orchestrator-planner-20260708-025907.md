## MERGE-GATE — the s6 execution relay (token-bearing): integrate `s6-transport-impl@58f2233` into `main`, exactly four bounded steps; operator authorization quoted verbatim; on your green report SLICE-6 CLOSES and Step-1 closes with it

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s6-merge-gate
PARENT_DISPATCH_ID: s6-exit-gate
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the human gate is EXERCISED: operator authorization on the decision packet (`MERGE-GATE-orchestrator-planner-20260708-025547.md`), quoted verbatim below
DESIGN_LOCK_ID: s6-slice-6-design
GRILL_LOCK_ID: s6-grill-s6-core
PLAN_LOCK_ID: s6-slice-6-plan
IN_REPLY_TO: MERGE-GATE-orchestrator-planner-20260708-025547.md
FROM: s6.orchestrator-planner
TO: s6-core.implementer
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s6.orchestrator-reviewer, s6-core.planner
SUBJECT: execute the s6 integration — verify-then-merge `--no-ff` of `s6-transport-impl@58f2233` into `main` · post-integration uncached battery on the merge commit before anything else lands · annotated tag `s6-close` · lint-clean execution report under this DISPATCH_ID; conflict ⇒ abort + escalate; fix-forward and push NOT authorized

**Operator authorization of record (interactive session with this seat, 2026-07-08, in response to the decision packet `MERGE-GATE-orchestrator-planner-20260708-025547.md` — quoted verbatim): "authorized, write it for me"** — decision ① granted; decisions ②/③ stand at their stated recommendations (executor = s6-core.implementer; annotated tag `s6-close`); decision ④ (the VP confirmatory pass) was not elected as gating — it may ride after the close at the operator's discretion.

**Acceptance chain (the token's basis, all on record):** the pair's gate record @ 58f2233 → this seat's verification + released exit SITREP (`s6-exit-gate/SITREP-orchestrator-planner-20260708-024558.md`) → master's acceptance with its own battery + probes (`../.relays/s6/s6-exit-gate/RECONCILE-orchestrator-planner-20260708-025218.md`) → the operator's grant above. Nine independent verification stations; verdict everywhere was merge-blocked pending exactly this grant.

**Authorized steps — exactly four, in order, nothing else:**
1. **Verify-then-integrate:** in the `main` checkout, confirm `s6-transport-impl` tip = `58f2233` and `git merge-base main s6-transport-impl` = `2903d84`; then `git merge --no-ff s6-transport-impl` (expect first parent = the ledger head `86dac3a`-or-later docs-only, second parent = `58f2233`). **Any conflict ⇒ `git merge --abort` + escalate to me; no in-flight resolution.**
2. **Post-integration battery on the merge commit, before anything else lands:** `go clean -testcache && go test ./...` (expect 24 packages ok, zero failures) + `go vet ./...` clean. **Red ⇒ STOP + report; fix-forward is NOT authorized.**
3. **Annotated tag `s6-close`** on the integration commit.
4. **Execution report** — lint-clean, under this DISPATCH_ID (`MERGE-GATE-implementer-<ts>.md`), carrying the merge SHA + parents (graph-verified), the battery output summary, the tag object, and `git status --short`. Root-mode lineage resolves against THIS relay (same DISPATCH_ID, the layered-grant rule).

**Not authorized:** any push · any history rewrite · any edit to any file · fix-forward on a red battery · any second merge or tag · anything beyond the four steps.

DISPATCH MERGE

ACTIONS_GIT_REF: none by this relay — the token grant itself; execution and its git refs belong to the executor's report.
FINAL_GIT_STATUS_SHORT: none — clean tree (main@86dac3a; branch of record clean at 58f2233).
Next requested action: execute steps 1–4; report under this DISPATCH_ID; on the green report SLICE-6 CLOSES — the close record + close SITREP to master follow from my seat, then the Step-1 close fold at master.
