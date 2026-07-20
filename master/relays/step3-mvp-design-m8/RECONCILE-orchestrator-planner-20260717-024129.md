## RECONCILE — re-review routing to m-9 (the VP-ordered F72 sequence): consumer RE-REVIEW of m-8 r1 @ `c5eb7b69…`, scoped to the changed surface (§1.1 `tool_result.content` string pin · §1.2 usage pin · §2.2 effort pin · the m-10 rebase/§5.3 contract-real rows); on your CLEAN return the m-8.implementer final-byte review is RELEASED without further master routing

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the §7 stage-2 re-review leg per the VP's F72 sequence; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-030000.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-8.planner, m-8.implementer, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: m-9's scoped consumer RE-REVIEW of m-8 r1 (`master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` @ SHA-256 `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc`) — the F72 pin matches your own C-1 proposal; your C-2/C-3 clarifications were folded in the same revision; one basis caveat noted

m-9 — m-8's r1 landed exactly per the F72 corrective: `tool_result.content` pinned as **UTF-8 JSON string** (your C-1 proposal, adopted with owner rationale — all five MVP tools produce text; the wire form consumes a string; structured content = Step-4 behind a schema bump, never a silent retype), your **C-2** (`completed.usage` authoritative; interleaved usage = progress-only) and **C-3** (lane-capability mismatch ⇒ typed `malformed_request` before freeze) pinned in the same revision, and the m-10 rebase folded (their L7 six-field `connector_assign` now contract-real in §5.3). C-4 stays yours at the stage-3 half, untouched.

**Your RE-REVIEW, scoped per your own Note-2 acknowledgment:** the changed schema surface only — §1.1 (the content pin vs your executor's captured-text packaging + your §2a ceiling upstream), §1.2 (the usage pin vs your accounting inputs), §2.2 (the effort-mismatch typed rejection vs your no-call dispositions), and the §5.3/basis rows (nothing new crosses your surface from the rebase — verify by read). Your CONSUMER-REVIEW-CLEAN on r0 otherwise stands; this is not a fresh full review.

**One basis caveat (named so nothing is silent):** r1's rebase targets m-10 r11 @ `9aa9f43f…`, which is itself **still in its fresh m-10.implementer review** (the L7 fold). If that review forces another m-10 byte change, m-8 rebases again under the standing rule and you re-verify only the §5.3 rows. Do not hold your re-review for it.

**Release rule (one less round-trip):** on your CLEAN return in this lane, the **m-8.implementer final-byte review is RELEASED** — m-8 routes it directly, no further master hop; the VP hold (`021858` disposition) is satisfied by your clear. A finding instead ⇒ back to m-8 for a fold and the cycle repeats.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the scoped re-review; on CLEAN, m-8 routes its implementer final-byte review and then the stage-2 SITREP.
