## SEQUENCING — s5-b: the ③ live-wiring task is UNGATED (m-6.implementer APPROVED the signal set; master GO at 153721); plus the M-4 archive-replay leg is ACCEPTED (the copy exists, near-zero cost); one bounded hop closes the slice's last work item

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s5-b-wire3
PARENT_DISPATCH_ID: s5-b-merge-gate
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-wire3 (fresh, off main @ b30df4d)
FROM: s5.orchestrator-planner
TO: s5-b.planner
CC: s5-b.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-b-merge-gate/RECONCILE-orchestrator-planner-20260706-153721.md
SUBJECT: the m-6.implementer signal-set confirm is CLOSED/APPROVE (master 153721, confirm 052907, fold 053113 — both now in the s5 trail); your plan's gated ③ live-wiring task is GO on a fresh branch off b30df4d; the M-4 optional replay leg is accepted into the same hop (copy at ~/frank-s5-team/replay-store-dogfood-20260706, 41 records); standing gates unchanged

Two gate changes and one acceptance; your locked plan and design are otherwise untouched.

**1. ③ live wiring — GO.** The m-6.implementer adversarial confirm of the S1/S2/S3 signal set is an unconditional APPROVE (.relays/s5/s5-escalations/DESIGN-REVIEW-implementer-20260706-052907.md; folded at …-053113; GO at …-153721 — read all three, they are your semantics record for this hop). Execute your plan's gated wiring task exactly as designed: the detector-config binding (the §J2-A-set-as-config read, load-once) + the cmd/* integration, on a FRESH branch `s5-b-wire3` off main @ b30df4d. Binding constraints, restated from the confirm: the Step-1 claim boundary — detection = exactly (S1)+(S2)+(S3)+`other`→A — is MANDATORY on every claim surface this wiring adds (code comment, doc line, tool text); the C1+C2 §C4 registrations remain explicit integration fences; no scope beyond the wiring task.
**2. M-4 — ACCEPTED into this hop (my discretion, exercised).** The archived-store copy sits at `~/frank-s5-team/replay-store-dogfood-20260706` (verified present at my seat: binding/config/journal/mailboxes). Run the zero-loss replay's archive leg against it via your store-path-parameterized harness (R-s5-5): READ-ONLY on the copy, count derived from the copy's own journal (not hard-coded), the count/identity/canonical-wins assertions as designed. Real traffic — 41 records incl. rejects, bounces, the waiver, the config history — is exactly the record-shape variety a constructed store never contains. If the copy's contents surface an unexpected incompatibility, that is a FINDING (report it), not a fix-in-place license.
**3. Ledger adjustment (FYI, no action):** master's probes at b30df4d confirm the m-1-driven scope bytes landed complete in s5-a's pass (genesis in NO scope incl. operator; owed/gate_resolution/disposition operator-only; `*` = diagnostics). The riding-confirms ledger is retired; zero open m-x legs stand between this hop and the exit SITREP.

**Process (standing, unchanged):** your pair gate applies (implementer review of the wiring diff per your established pattern — a bounded delta-review suffices if the diff matches the locked design's wiring section); combined-tree discipline applies (your own two-tree runs per your adopted practice — the branch is off the current main tip so branch-tree ≡ combined-tree unless main moves); then your integration-ready report TO me; my gate verification; the written merge-authorization chain (the 150902 pattern) closes it. SCOPE_DIFF in the dispatch as ever; any out-of-surface need is an escalation.

**After this hop (the close sequence, for your visibility):** my exit-gate evidence assembly → the close-gate sprint-docs commit (operator sign-off) → my exit SITREP to master → master's independent battery at the close tip → the master close fold + the operator's integration/tag gate.

FINAL_GIT_STATUS_SHORT:
none — tracked-clean at frank/ main @ b30df4d; sprint docs untracked by design until the close gate
ACTIONS_GIT_REF: none — sequencing relay; no code/git action at my seat
Next requested action: execute the hop; report integration-ready with two-tree evidence + the archive-leg result; the merge packet follows my verification.
