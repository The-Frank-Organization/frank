## RECONCILE — operator-raised in the stage-5 grill (G-5), routed UP: courier READ RECEIPTS — iMessage/WhatsApp-style sender-visible "read" status on relays, replacing pure-acknowledgement relays and some SITREP forms; CONDUCTOR-surface owned (m-7 verb/store + m-1 visibility), recommend hardening-backlog entry beside T11; one honesty rail named (read ≠ acted)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-stage5-m10
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an escalation of an operator-raised feature idea; the operator already holds it
GRILL_REQUIRED: no — the stage-5 grill is CLOSED (GRILL_LOCK `m10-stage5-grill-lock-20260720`); this item was raised in it and is severable
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
IN_REPLY_TO: step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-045603.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-7.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: during the G-4 email-semantics ratification the operator extended (verbatim intent): "a read receipt like in imessage/whatsapp so the sender knows they've been read without the recipient necessarily needing to write a relay back, like in the case of an acknowledgement relay or some forms of sitrep relays" — a relay-volume/friction win squarely in the battle-reports-harden-frank channel; NOT mine to design, handed up with the boundary and rails named

Master — the idea and its shape, for your scoping:

1. **The mechanism, one sentence:** when a seat `read`s a relay, the conductor's store records a seat-stamped read fact, and the SENDER can see it via `project` — so "did they get it?" stops costing an authored acknowledgement relay.
2. **The prize:** pure-ACK relays and receipt-shaped SITREPs are real hand-relay/context cost today (the friction-IS-the-roadmap channel); a store-level read fact deletes a whole message class without weakening the trail — arguably STRENGTHENING it, since a read fact is mechanical truth where an ACK relay is authored prose.
3. **Ownership + the questions that come with it (why this routes up):** the read fact is a **conductor store/verb change** — m-7's surface (the `read` verb gaining a durable side effect; note their store is currently sole-GOVERNED-writer with reads free — a read-writes-a-fact change touches that discipline non-trivially) — with **m-1 visibility rules** (who may see who read what; seat-scoped like everything else) and a FieldSpec/lineage question for how the receipt projects.
4. **The honesty rail, named now so it never gets lost:** **read ≠ acted.** A receipt proves delivery-and-sight, never acknowledgement-of-action — my own app-side `dispatched` flag (processing-began) is already a strictly stronger fact, and the two must never be conflated in any surface; a receipt presented as an ACK would be exactly the dishonest-labeling class the constitution forbids.
5. **Recommendation:** a named hardening-backlog entry beside T11 (they compose: T11 wants delivery-of-a-fact upgrades, G-5 wants delivery-CONFIRMATION upgrades; both are courier-evolution, Step-4+). No MVP dependency exists — my stage-5 design consumes nothing from it.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row timestamped 20260720-053000; the G-4/G-5 folds + §15 GRILL_LOCK landed in the UNLOCKED stage-5 working draft (now r1 @ `b04201b6e4f5e2954b4d6d60164a3f9bc470951ef285ca60859fb835c17db11f`); no closed design doc, no `frank/` action, no lock claim.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: master records G-5 as a hardening-backlog entry (or disposes otherwise); m-10.planner proceeds to the final-byte pair review — no return owed to unblock stage 5.
