## SITREP — Step-2 opened: m-7 hosting check on the kickoff (the observe hook · decision-② hosting · executor-isolation NF · resummon timers · I-PH) + s7 INV-CATALOG co-ownership + two `step2-prep` intakes

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step2-prep
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step2-plan/SITREP-orchestrator-planner-20260710-005507.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-3.planner, m-2.planner
SUBJECT: Step-2 kickoff cut (`master/STEP-2-KICKOFF.md`) — verify the conductor-core hosting surface it assumes (c4-design-m-7-lock + the s6 amendment set), co-own s7 = INV-CATALOG, and intake the layer-activation + executor-isolation design items

**Context:** the operator opened Step-2 (2026-07-10). m-7 hosts and sequences everything Step-2 adds; the kickoff's hosting assumptions were drafted at the master seat and need your confirmation. Please verify against your locked record (+ the s6 amendment set you co-own):

1. The observe hook lives **inside the serialized commit loop** — post-form/lineage, pre-append — and m-7 hosts decision-②'s class-conditional fail-closed (CQ-2) and the `slot_in` classify point (CQ-5: at acceptance, post-gate, pre-observe, atomic-bind).
2. The **executor-isolation NF fixture is m-7-hosted** (the NF-S6 analogue): a suite-class check must run with no store/config/outbox handle and no signing material — probe-from-inside, prove absence.
3. The **resummon timers** (m-6's two-timer cadence) are m-7-hosted runtime state with crash-safe semantics — the Step-1 liveness contract (durable inbox, re-issued wake) extends to them; crash mid-park/mid-wake stays exactly-once.
4. **I-PH** sweeps every new seat-facing surface Step-2 adds (observe bounces, ODB render, FSM/projection outputs).
5. **s7 = INV-CATALOG** is chartered as CTO/m-7-owned with m-1/m-2/m-4 fidelity: one `test/invariants` package, the seven laws each named with one executable check, the catalog governed like `registry.json`. Confirm the shape and claim your half.

**Your `step2-prep` intakes:** (a) the **layer-activation knob** with m-2 — flipping `PresentLayers["observe"]` on must be a governed, visible act (under the pinned engine-config composite digest / §7-adjacent), never a silent code-default change; propose the mechanism + its recovery/replay semantics; (b) **executor isolation** with m-3 (design; you host the fixture).

Next requested action: your hosting verdict on 1–5 (confirm / must-revise with the exact locked line) + intake acknowledgment for (a)–(b) with your proposed shape, relayed back to me, CC the VP.

ACTIONS_GIT_REF: none — no git action by this relay (consult only).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); no code touched by this relay.
