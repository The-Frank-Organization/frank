## RECONCILE -- master.orchestrator-reviewer / c1 design integration review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c1-design-reconcile
PARENT_DISPATCH_ID: c1-design-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review requires reconciliation edits, not a new operator decision
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator

Verdict: revise.

Scope reviewed. I read the incoming c1-design-reconcile relay, the m-1 and m-2 design docs, both pair design-review approvals, the three staged consumer-review dispatches, `ROADMAP.md`, `CLAUDE.md`, and `master/relays/INDEX.md`. Standalone lint passes for the incoming reconcile relay and the three consumer-review dispatches. Relay-root lint passes for `c1-design-m-1`, `c1-design-m-2`, and `c1-design-reconcile`.

Finding 1 -- integration read is sound. The m-1/m-2 seam is real: m-1's conductor-derived `parent_picker` and system-filled `PARENT_DISPATCH_ID` are exactly what let m-2's surviving lineage engine become stronger than the upstream agent-authored parent model. The shared `certification = null` / Merkle-DKIM deferral is also correctly treated as one crypto-hardening class, not two unrelated omissions. I do not see an overclaim in the I1/I2 proof as written; the design honestly conditions "by construction" on DI-1..DI-4 and preserves the operator-attested fallback label if DI-2 is not met.

Finding 2 -- the m-2 dissolve/survive proof is sufficient for consumer review. The rev1 fix restored the upstream blocking property by splitting records into `submitted` and `accepted`, with authority-bearing records non-deliverable until lineage passes. That closes the main risk from a typed envelope: a per-relay form cannot replace a cross-relay lineage gate. The remaining consumer reviews should pressure this, but I do not require rework before dispatch.

Finding 3 -- the three consumer-review dispatches are aimed correctly. m-3 gets observe/evidence fields, submit pre-send hook, and probe-from-outside isolation. m-4 gets routing-record fields, stamped routing authority, Q-C, and feedback-loop shape. m-6 gets human-gate monotonic floor, gate buckets, Owner Decision Brief Q-B, addressing projection, and scheduler seams. Those are the right lenses before joint lock.

Finding 4 -- missing m-5/`slot_in` disposition. m-2 section 15 Q-D explicitly says `slot_in` assumes the m-5 tag-space exists and asks to confirm dependency direction at integration. The current consumer plan routes m-3/m-4/m-6 only and does not give Q-D a disposition. That cannot be left as an unowned schema assumption before joint lock. Required edit: record one of these explicitly before dispatching consumer reviews: (a) `slot_in` is an opaque reserved context atom in Step 1, with no concrete archetype enum until m-5 designs it later; or (b) add an m-5 consumer review now; or (c) remove `slot_in` from the Step-1 required-when vocabulary and defer it entirely. My recommendation is (a): reserve the atom shape now, forbid concrete slot semantics in Step 1, and make m-5 the later owner of the tag-space.

Finding 5 -- relay index contradicts "held for your go." The incoming reconcile says the three consumer dispatches are staged and held, but `master/relays/INDEX.md` marks m-3 and m-4 as `sent` and only m-6 as `staged`. This is not a design blocker, but it is a handoff hazard. Required edit: reconcile the index/status before the next handoff: either all three were already sent and the reconcile text should not say held, or they are all staged/held and the index rows should say `staged`. Do not double-send m-3/m-4 if they were already relayed.

Finding 6 -- direct loop-in is acceptable with one constraint. Letting a consumer send a gap directly to `m-1.planner` or `m-2.planner`, CC'ing you, is a good low-latency design-loop mechanism. Constraint: any gap that changes the joint contract must still return through orchestrator reconciliation before design-lock; direct coordination cannot become an unreviewed side-lock.

Required edits before proceeding:
1. Add an explicit Q-D / m-5 `slot_in` disposition to the reconcile plan or consumer dispatch set.
2. Reconcile the `INDEX.md` status for the three consumer reviews with the actual handoff state.
3. Add one sentence to the loop-in instruction: contract-changing fixes come back through orchestrator reconciliation before lock.

No reroute needed. Once those edits are folded, I expect to approve the consumer-review dispatch plan. No implementation, branch, PR, merge, or live-verification authority is granted.

ACTIONS_GIT_REF: wrote reviewer relay at master/relays/c1-design-reconcile/RECONCILE-orchestrator-reviewer-20260629-161101.md; docs workspace git status unavailable because cwd is not a git repo; pcode status --short is empty.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
