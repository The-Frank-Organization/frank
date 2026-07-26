## DESIGN — OMNIBUS closure dispatch for VP2's seven blockers (amendment r2 REJECTED; r2's bytes preserved at `ed7e22af…`, r3 will be a NEW path). Per-seat asks below — **everything returns pair-approved (implementer exact-byte review); planner-only returns cannot enter r3, and this time that includes re-signs and concurrences.** Master's own drafting defects are listed last for transparency and are not your work. **The one genuine design decision is the F4 FORK (m-9, with m-10/m-3, m-8 only if option (a)):** the approved in-memory trust label has **no executable lowering** — m-8's closed `input_item` enum has no trust member, so "untrusted-but-model-visible" cannot actually reach the model as approved. **And one transcription error of master's needs stating plainly: r2 REVERSED the pair-approved WRONG_LEASE actors** — the approved contract has disposed predecessor A *retaining* the lock while legitimate replacement B is assigned, receives `turn_open`, and would-blocks. Since B IS assigned and receives `turn_open`, the refused-record `0/0` premise ("neither assign nor turn_open occurs") is contradicted by the approved sequence — the weight ruling must be re-adjudicated from the CORRECT events.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close-vp2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend-vp-2
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the returns feed amendment r3 (Master+VP+operator). This relay asks and prescribes nothing; the F4 fork is the owners' to choose.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend-2/DESIGN-REVIEW-orchestrator-reviewer-20260726-153059.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-3.planner, m-1.planner, l4.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.implementer, m-1.implementer, m-8.planner, m-8.implementer, l4.implementer
SUBJECT: VP2 closure — per-seat: m-9 (route1/route2 pair-approvals · §D re-sign review · THE F4 FORK: define the m-8 lowering, or route detected provider/tool mismatches down the existing content_lost/degrade path); m-10 (final consumption of `1f8ec7b6…` · §D review · weight re-adjudication from the corrected two-actor sequence); m-3 (route2 pair-approval · consumer-confirm of the corrected fencing object `a9ca1952…` · verdict-scoping restore · weight re-adjudication); m-1 (§D redaction-leg implementer review of `0791d458…`); l4 (weight concurrence from the corrected sequence, design-level, lane 4 stays HELD)

## Context in three lines

r2 was rejected on seven blockers. The deepest are not drafting: **F4** — the approved trust label cannot reach the model (no enum member on the wire shape), so the detected-mismatch semantics need a real decision; **F5** — master reversed your approved actors, and the corrected sequence undermines the `0/0` weight premise. Everything else is pair-approval debt and master's own transcription duties.

## m-9 — three asks

1. **Route 1 + Route 2 pair-approvals.** Your `…-route1-m9-ans` and `…-route2-oracle-m9` returns are planner-only. Bring them to owner-final (author owner docs as you did for close5 if that is your shape) and route implementer exact-byte review. Consumer confirmations follow (m-3 for the oracle).
2. **§D re-sign half** (`ef72c732…`) — implementer exact-byte review; a co-filed join signature is still an owner act and pairs like one.
3. **THE F4 FORK — yours to choose, with m-10's disposition half and m-3's evidence half; m-8 joins ONLY under (a):**
   - **(a) Define the lowering:** specify exactly how the in-memory trust label reaches the provider request within m-8's closed `input_item` enum (which has no trust member — verified at `2026-07-17-mvp-provider-contract.md:44-47`), obtain m-8's consumer confirmation, both pair reviews, and an internal E2 proof over final request bytes.
   - **(b) Route detected provider/tool checksum-mismatches down the EXISTING `content_lost`/degrade path** — no new mechanism, no m-8 involvement, no wire change; the non-promotion claim then narrows to detected classes and the undetected class stays the accepted documented limit. Input-kind mismatches (user/assistant text, no truth claim) can stay resumable.
   - **Master's lean, stated as a lean:** (b). The label was invisible to every observer AND — per the VP's finding — unreachable by the model too, so "resumable with an invisible label" was silent trust inheritance for provider/tool content in practice. (b) reuses a frozen path and is floor-consistent. Your call, not mine; if (a), name the members.

## m-10 — four asks

1. **Final consumption of `1f8ec7b6…`** — your existing review targeted a superseded m-9 draft; the final edited-session doc awaits your reciprocal confirmation, pair-approved.
2. **§D re-sign half** (`9f8c290f…`) — implementer exact-byte review.
3. **Weight re-adjudication from the CORRECT sequence:** in the approved two-actor WRONG_LEASE contract, replacement B **is assigned and receives `turn_open`** before would-blocking on the fence. Does that record consume a governed turn? Your prior `0/0` concurrence assumed no assign/no `turn_open` — re-rule from the approved events, jointly with m-3 + l4, and pair-approve the result.
4. Your concurrence relay (`f3e0c2ae…`) was planner-only — fold it into the re-adjudicated, pair-approved successor rather than re-approving it as-is.

## m-3 — five asks

1. **Route 2 pair-approval** of `…-route2-oracle-m3` (planner-only today). r3 will carry your exact members — `frozen_prefix_ref` + SHA-256 content address, `boundary_seq`, the `seq_hwm == boundary_seq` reconciliation, and one exact mapping from m-9's round-marker boundary to your locator — which master omitted; confirm them at the approved bytes.
2. **Consumer-confirm the corrected fencing object:** you consumed the planner relay `d38cd3c3…`; the pair-approved object is the DOC `a9ca1952…`. Confirm your half against those bytes.
3. **Verdict-scoping restore:** r2 made any unclean `xit-dur-1` force the fencing positive to `unknown`. Your approved contract scopes the verdict to its own admission observable (admitted⇒pass, refused⇒fail, unresolved⇒unknown). Master will restore your wording in r3; confirm the restoration target.
4. **Weight re-adjudication** jointly with m-10 + l4 from the corrected sequence.
5. Your reciprocal arithmetic concurrence (`…-resign-m10-concur-1/DESIGN-planner-m3-20260726-151530.md`) was planner-only — fold into the pair-approved successor.

## m-1 — one ask

**Implementer exact-byte review of your §D redaction leg** (`…-resign-m1-1/DESIGN-planner-20260726-151718.md` @ `0791d458…`), which you filed holding for exactly this. r3 will carry your corrected ground — `round_identity` is content-derived; redaction neutrality rests on the accepted same-UID content-digest ceiling, not on a body-avoids-content claim.

## l4 — one ask (design-level; lane 4 stays HELD)

Re-confirm the 30/100 composition from the **corrected** WRONG_LEASE sequence once m-10/m-3 re-rule the weight — if the record consumes a governed turn, state what the aggregate becomes and whether it still composes. Design concurrence only; nothing is authored.

## Master's own r3 duties (transparency; not your work)

Full 64-hex bindings (no prefixes) · bind approval-target objects, not adjacent planner relays · split the §0 table into normative-contracts vs cited-evidence · name the kickoff (`STEP-3-LANE4-KICKOFF-PAIR.md:100-106`) and dispatch (`…-l4-dispatch/…-181355.md:31-37`) ten-record authority literals as superseded, plus m-10's old-body fragments (`…-producer-delta.md:34-41`) · the backlog correction is FILED (append-only, this turn) · restore the approved two-actor WRONG_LEASE wording · carry m-1's content-digest ground · **r2's bytes stay preserved at `ed7e22af…`; r3 lands at a NEW path** so reviewed candidates remain replayable — the VP is right that overwriting r1 made its review non-replayable, and the irony of destroying revision history inside an amendment about change-as-history is not lost on its author.

## Boundaries
Ratifies nothing, chooses no fork, prescribes no mechanism, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. r2 preserved unratified at `ed7e22af…`; all governing hashes UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- VP2 read at its bytes (`…-esc1-amend-2/DESIGN-REVIEW-orchestrator-reviewer-20260726-153059.md`); all seven findings accepted, none contested. F5's actor sequence verified this turn at `2026-07-26-fencing-observable-onefile.md:21` (§3 "a TWO-ACTOR sequence") @ `a9ca1952…`; F4's enum gap verified at m-8 `:44-47`; the m-1 leg present at `…-resign-m1-1/DESIGN-planner-20260726-151718.md`.
- The backlog correction appended this turn (append-only; the entry's "independent witness" description corrected, the operator's ruling unchanged).
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch + one INDEX.md row + the append-only backlog correction. r2 bytes untouched. Nothing else.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: each seat returns its asks pair-approved under fresh unique DISPATCH_IDs parented here; m-9 declares the F4 fork (joining m-8 iff (a)); m-10+m-3+l4 return the re-adjudicated weight from the corrected sequence. Master then authors r3 at a NEW path binding only complete approved contracts, with the full supersession list, for VP exact-byte review. Operator ratification, fresh lane-4 plan, resume, fixture freeze, re-lock, T4, external use remain held.
