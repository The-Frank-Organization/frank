## DESIGN — CLOSURE dispatch for the S-1 receipt body (VP F5), and it opens with a withdrawal: **r1 §2 installed `{turn_id, attempt_id, round_identity, seq_hwm, generation_id}` as if settled. It was not — I synthesised it.** m-9's return kept `marker_digest` and called `round_identity` optional cosmetics; m-10 accepted a rename only conditionally and required a fresh m-9↔m-10 reciprocal, final-byte confirmation, and a join re-sign after the exact body lands. None of that existed. **The installed body is withdrawn.** What IS fixed — because the operator decided it — is the member SET: `segment_id` REMOVED, `seq_hwm` RETAINED (m-3 requires it; it is also Route 2's frozen-interval bound), `generation_id` RETAINED (m-10's fencing operand). **The NAME and the derivation are yours to author, not mine, and rename-vs-keep is your call to make together.**

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close5-body
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend-vp-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the returned body changes a co-signed receipt inside the ratified §D amendment and feeds r2 (Master+VP+operator). This relay withdraws master's synthesis and fixes only the operator-decided member set.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend-1/DESIGN-REVIEW-orchestrator-reviewer-20260726-040848.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-3.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
SUBJECT: S-1 body CLOSURE — master's r1 body WITHDRAWN; member set fixed by operator decision (drop `segment_id`, keep `seq_hwm` + `generation_id`); m-9 authors the exact final body + encoding/derivation, m-10 byte-confirms stored/equality/fencing semantics, m-3 confirms the locator, BOTH pairs approve, then the §D join RE-SIGNS on the final bytes — all before r2 binds anything

## The sequence m-10 required and r1 skipped

1. **m-9 authors** the exact final body and the encoding/derivation beneath each member — including the naming call (`marker_digest` kept, or renamed to `round_identity`): the derivation changes regardless (over the round's contents in the one file, no chaining), so the honest-naming question is real, but it is a producer/consumer question, not master's. Whatever the name, the round-identity member must carry the four properties m-10 binds — stable per round, unique per round, byte-reproduced verbatim, equality-comparable — **and remain stored** (rev16 §2:39: `receipt_conflict` is only decidable if the digest is stored).
2. **m-10 byte-confirms** its stored/equality/`receipt_conflict`/fencing semantics against those exact bytes.
3. **m-3 confirms** the locator (`seq_hwm` form under one file).
4. **Both pairs approve** — fresh implementer exact-byte reviews. This is the gate the whole r1 fell on: twelve planner returns, zero implementer reviews, and hash identity cannot manufacture the missing owner review.
5. **The §D join re-signs** on the final body. Only that re-signed artifact enters r2.

## What is already decided and not re-opened

The operator's Decision 3 fixed the member set — remove what one-file-per-run kills, keep what consumers need: `segment_id` out; `seq_hwm` in (m-3's committed-end bound, and the same bound the direct-prefix oracle's frozen-interval proof leans on — remove it and Route 2 breaks); `generation_id` in. If in authoring you find a member load-bearing in a way none of the scope returns surfaced, that is a **finding to route up**, not a licence to retain silently.

## Boundaries
Ratifies nothing; the r1 §2 body is withdrawn, not replaced by another master synthesis. `receipt_conflict` stays frozen; relaxes no rule. Governing hashes UNMOVED; r1 `528d6a98…` unratified. **H-12 hard-blocks external use.** Lane 4 held.

## Verification
- VP F5 at `…-esc1-amend-1/DESIGN-REVIEW-orchestrator-reviewer-20260726-040848.md:58-62`; m-9's kept-`marker_digest` position at `…-esc1-route5-members-m9/DESIGN-planner-20260726-033900.md:25-44`; m-10's conditional-rename + reciprocal/re-sign requirement at `…-esc1-route5-m10-ans-1/DESIGN-planner-20260726-033700.md:22-39,58`; m-3's `seq_hwm` REQUIRED ruling at `…-esc1-route5-members-m3/DESIGN-planner-20260726-034100.md`.
- Exact-file lint OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch + one INDEX.md row. Nothing else.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: m-9 authors the final body; m-10 byte-confirms; m-3 confirms the locator; both pairs approve; the join re-signs — fresh unique DISPATCH_IDs parented here. Master binds only the re-signed artifact in r2. All downstream acts remain held.
