## RECONCILE — H-16 NARROWING SUPPLEMENT to m-7 (append-only; incorporates VP F85/F86 + your own pair's R1-F1..F3, which the VP reviewed and ruled SOUND): the fix is NOT implementation-ready — no IMPL branch until (1) the durable post-commit work contract is total over ALL FIVE returned-fault sites + the quarantine/panic/recovery routes, (2) `post_commit_state` regains `unknown` (the dispatch dropped it; canonical H-16 requires it), (3) the caller MIGRATION TABLE exists (additive-ignored-field compat FAILS OPEN — the four state-only success consumers are enumerated), (4) the per-hook idempotency/result-durability table exists (seat-mint rotation under blind retry is the named hazard) — sequence: corrected design + a FOCUSED DECISION RECORD → fresh uniquely-parented pair review → master/VP review → ONLY THEN the IMPL branch

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split
PARENT_DISPATCH_ID: master
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the operator retains the merge grant; and the decision record's API/recovery choices are hard-to-reverse semantics the VP requires reviewed before IMPL
GRILL_REQUIRED: no — the VP requires a focused decision record, not a grill
DESIGN_DOC_ID: h16-outcome-split-design
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-045522.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, master.orchestrator-reviewer, operator
SUBJECT: corrections to my `043321` dispatch, each verified at the VP's loci — (a) `:279-280` is NOT a callHandler path (callHandler is pre-commit `:210-249`; `:279-280` is the post-commit completeTurn branch in supersededCredentialOutcome — five returned-fault sites total, not four+one) · (b) the total census must also DISPOSITION `processQuarantine` (`:122-128`, discards completeTurn failure), the panic path (`:130-134,:376-379`, serves the decision with no derived-work state), and startup/recovery — disposition ≠ change; none may stay unclassified · (c) `{complete,pending,failed,unknown}` — `unknown` is required for a crash after a possibly non-idempotent effect · (d) your R1-F1 (the in-memory slot loses the fault + hook cursor on crash; restart derives complete; your own T2 test unpassable) is CONFIRMED as the F85 heart — the durable work contract below is the fix bar · (e) intake recovery (`internal/recover/recover.go:71-84`) replays only unconsumed commands — it is NOT a post-commit work driver today; your design must name what IS

m-7 — the corrected bar, from the VP verbatim where it is exact:

1. **The durable post-commit work contract (F85):** one durable work identity keyed by decision/intake + hook contract version · an atomically-committed or deterministically-derivable unresolved state · an exact ordered hook cursor (`not_started` / `running_or_unknown` / terminals) · the retry owner, pre-Ready recovery drain, retry ceiling, park rule, and terminal-resolution record · table reconstruction + duplicate/restart behavior · error/panic/hard-crash cuts at EVERY post-commit boundary incl. the superseded and quarantine routes. An in-memory slot may schedule; it may not be the source of truth.
2. **The caller migration table (F86):** additive-ignored-field is FAIL-OPEN — a legacy decoder reads `state:"accepted"` as complete while derived work is pending/failed/unknown. The four in-repo state-only success consumers: the delivery nudge (`cmd/frank/main.go:337-352`) · the gate/approval prompters (`internal/engine/prompter.go:81-99`) · resummon (`internal/engine/resummon.go:228-249`) · the MCP/native/shared-client JSON projections + external clients. Choose: a versioned/negotiated projection (old clients CANNOT observe completed acceptance for non-complete work) or one atomic migration of every success consumer — and prove the choice in the table across pre-commit-failure / committed+complete / pending / failed / unknown.
3. **The per-hook idempotency table (F86):** per hook — idempotent? result durable? replayable blind? The named hazard: `AfterAccepted` → seat mint delivery; `MintOrReplace` mints a NEW credential per invocation and the extras are absent from `outcomeFromRecord` — blind retry rotates a credential no caller receives; conversely a generic retry must not skip mint or run it against an unrelated command (your R1-F2/R1-F3, VP-confirmed).
4. **Sequence (VP-fixed):** corrected design bytes + the FOCUSED DECISION RECORD (the versioning/delivery/durability/secret-result choices, each with the rejected alternative) → fresh uniquely-parented m-7.implementer review → return to master for the master/VP pass → only then the IMPL branch opens. The operator's merge grant is unchanged at the end.

My `043321` stands as history (append-only); this supplement is the operative narrowing. Your pair's independent catch of R1-F1..F3 before the VP's review is the discipline working — carry on exactly so.

ACTIONS_GIT_REF: docs-workspace disk action — this supplement relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: m-7 returns the corrected design + decision record + fresh pair review; master routes the master/VP pass; no IMPL branch before it.
