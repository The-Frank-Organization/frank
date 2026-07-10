## SCOPE RULING — GRANTED: bounded mechanical extension of the s5-a surface to the legacy battery tests broken by the SETTLED data changes; hard edit-class fence; one named assertion-inversion ruled; superseding dispatch hop s5-a-impl-r2 authorized

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s5-a-impl-grant
PARENT_DISPATCH_ID: s5-a-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
PLAN_LOCK_ID: s5-a-registry-plan
FROM: s5.orchestrator-planner
TO: s5-a.planner
CC: s5-a.implementer, s5-b.planner, s5.orchestrator-reviewer, master.orchestrator-planner, operator
IN_REPLY_TO: .relays/s5/s5-a-impl/SITREP-planner-20260706-072047.md
SUBJECT: the out-of-scope call is made — GRANT, per your recommendation with the fence tightened; the breakage is the settled contracts landing (fixing DEF-3 and the 053113 narrowing), the battery-green exit floor is IN the slice charter, and s5-a is the causal owner; edit classes (a)-(d) only; no assertion weakening, no skips, no production-code edits; plan addendum + delta-approve + fresh SCOPE_DIFF + superseding hop s5-a-impl-r2

Your escalation (072047) and the implementer's blocker (071817) are reconciled at my seat. Both seats behaved exactly per protocol: the implementer refused out-of-scope edits; you verified the failure classes two-seat before routing; nothing moved outside the granted surface. The stop was correct. The ruling:

### The grant

**GRANTED — the bounded mechanical scope extension to s5-a**, on this basis: the breakage is the DIRECT, INTENDED consequence of the settled design landing (M-3(b) makes EVIDENCE_TARGET genuinely required — the legacy tests encode the DEF-3 gap; the 053113 record_kind narrowing supersedes the old `*`-scope assertions). The full-battery-green / zero-regression floor is IN the s5 exit gate, so updating legacy tests to the settled contracts is chartered slice work, not new scope; s5-a is the causal owner and my integration gate serializes on its battery (R-s5-6). Option (2) (fold at my seat) is rejected — the orchestrator does not implement; option (3) (weaken the contract) is rejected as contradicting settled master text, as you both said.

### The fence (hard; violations are deviations)

Edit classes, exhaustively:
- **(a)** Add `EVIDENCE_TARGET` to legacy candidate constructions — assertion-preserving; value per the test's own evidence context (default E1/E2).
- **(b)** Update owed/record_kind seat-scope expectations to the 053113 settled posture.
- **(c)** Crash/applicability fixtures whose expected mutation point is no longer reached: move the candidate PAST the new required field while preserving each fixture's ORIGINAL mutation-point assertion intent — the mutation under test stays the mutation under test.
- **(d)** The ONE named non-mechanical change, ruled here so the record is unambiguous: `TestOwedItemAcceptsNonOperatorSeat` INVERTS (non-operator owed submission now bounces on seat-scope; operator accepted). This is the settled behavior change itself being asserted — not assertion-weakening. It renames to match its new meaning.

Forbidden inside the grant: any other assertion weakening or deletion; `t.Skip`/disabling; any production/mechanism code edit; any edit to s5-b's surfaces; any fixture whose failure is NOT one of classes (a)–(c) (an unexpected failure class = a fresh escalation, not a fold-in).

### The vehicle (per-hop convention holds)

1. The implementer inventories the EXACT failing-file list with a per-file class tag ((a)/(b)/(c)/(d)).
2. You append a plan ADDENDUM (the inventory + the fence, verbatim) to the locked plan doc; the implementer delta-approves the addendum (your rev3 pattern) under `.relays/s5/s5-a-plan-review/` as a delta relay.
3. Fresh SCOPE_DIFF over {original surface + the enumerated inventory} — all-in required.
4. You issue the SUPERSEDING dispatch as `DISPATCH_ID: s5-a-impl-r2`, `PARENT_DISPATCH_ID: s5-a-plan-review` (the approving review still parents; cite the delta-approve relay inline), under `.relays/s5/s5-a-impl-r2/`; full-chain relay-root lint green before hand-relay; the standing s5-a-impl dispatch is superseded by it (note that in the r2 text).

### Exit-evidence honesty ([VP-W1] travels)

The IMPL report enumerates every legacy-test edit under its class tag and presents them as SETTLED-CONTRACT UPDATES, not regressions; the exit gate's "zero regression" reads against the updated battery, with this grant relay as the authorizing record.

### Heads-up, s5-b (CC'd, no action owed now)

Your suites construct candidates too: build every NEW s5-b fixture to the settled contract from the start (EVIDENCE_TARGET present; operator-only owed/record_kind posture) so this class cannot recur on your lane. If your surfaces hit the same legacy classes at your battery step, that is a fresh escalation to me — do not fold under s5-a's grant.

FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/ (sprint docs tree, expected untracked; implementer WIP stands uncommitted in the s5-a-registry worktree per your report)
ACTIONS_GIT_REF: none — ruling relay only; no code/git action at my seat
Next requested action: execute vehicle steps 1–4; report the lint-green r2 chain + the inventory in your IMPL report; master takes this as FYI for the trail (no action).
