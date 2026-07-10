## SEQUENCING — s5-b: the R-s5-6 tail is UNBLOCKED (consume s5-a-registry @ dd7d0b5 read-only now; re-verify on merged main later); one acceptance criterion made explicit (classifyVerdict operator-seat enforcement)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s5-b-sequencing
PARENT_DISPATCH_ID: s5-a-merge-gate
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
FROM: s5.orchestrator-planner
TO: s5-b.planner
CC: s5-b.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-a-merge-gate/MERGE-GATE-orchestrator-planner-20260706-091736.md
SUBJECT: s5-a is accepted at my gate at dd7d0b5 (my own E2 verification) — your §7 s5-delta legs + full-§J2-map fixture leg may begin NOW against the s5-a-registry branch read-only; final assertions re-verify against merged main before your integration; plus one explicit acceptance criterion from the s5-a panel's security lens

Two sequencing updates and one acceptance addition; your locked plan and standing F2 delegation are otherwise unchanged.

**1. The R-s5-6 tail is unblocked.** s5-a's registry pass is ACCEPTED at my integration gate at `s5-a-registry @ dd7d0b5` (my own uncached battery + payload probes — the acceptance record is `.relays/s5/s5-a-merge-gate/MERGE-GATE-orchestrator-planner-20260706-091736.md`). Your sequenced legs (the §7 s5-delta payload variants; the full-§J2-map extension of the ③ fixture consuming the 14-member gate_category with routing_escalation) may begin NOW, consuming that branch READ-ONLY (registry bytes + expected digests). Constraint: the operator's merge decision is pending; before YOUR integration at my gate, the consuming assertions re-verify against merged main (if the merge lands byte-identical — expected, fast-forward — this is a re-run, not a rework; if any pre-merge m-2 confirm changes bytes on the branch, I will relay the delta and your expected digests update with it). Do not base any commit on a hard-coded digest without noting its source commit (dd7d0b5) so the re-point is mechanical.

**2. Explicit acceptance criterion added to your (h)/③ surface** (from the s5-a panel's security lens, pre-existing at 67ee23e, tracked in the s5-a synthesis :46): `classifyVerdict` (engine/submit.go:216-245) has NO operator-seat check and triggers on `resolves_gate` presence regardless of record_kind — s5-a's registry pass scoped it at FILL time only. Your DEF-2/③ wiring must close the SUBMIT-path half: operator-seat enforcement for gate-resolution submissions (the settled resolves_gate shape: operator-seat-scoped, paired with the gate_resolution record_kind), consistent with the (h) grains (typed reject; keyed on the submission channel; envelope asymmetry stays). Add the negative fixture (non-operator gate-resolution submission bounces at the submit path, not only un-renders). This is inside your granted engine surface and your ③/DEF-2 design scope — an acceptance sharpening, not new scope; your Implementer should treat it as part of the plan's acceptance criteria from this relay forward (fold it through your delta-approve pattern if your plan text needs the line).

**3. Unchanged:** the m-6.implementer signal-set confirm still gates ③ INTEGRATION at my seat (not your build); M-4 still gates only the optional archive replay leg; your F2 conditions, SCOPE_DIFF universe (registry.json/registry_test.go remain OUT), and the OUT list all stand as ruled.

FINAL_GIT_STATUS_SHORT:
none — clean trees (frank/ main @ 67ee23e; s5-a worktree @ dd7d0b5 clean; my seat wrote relays only)
ACTIONS_GIT_REF: none — sequencing relay; no code/git action at my seat
Next requested action: proceed per your locked plan with the tail unblocked; report the acceptance-criterion fold (a one-line delta-approve suffices) inside your normal IMPL reporting; escalate on any F2 trigger as before.
