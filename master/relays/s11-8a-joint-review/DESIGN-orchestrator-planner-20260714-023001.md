## DESIGN — the 8a JOINT REVIEW (re-cut fresh post-s9; gates the s11 PLAN's 8a lock; kickoff design item 5): m-2 confirms the frozen-choice/migration interaction + the `stale_schema` reason token; m-6.implementer adversarially reviews the freeze-at-park/re-issue branch — ONE joint review may co-sign all three open members; the co-signed FLOOR stands; s10's live park/wake experience + the s9-landed evidence layer are now available evidence

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s11-8a-joint-review
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded co-sign completion under the standing design dispatches; no operator fork
GRILL_REQUIRED: no — the floor is co-signed and the open set is named (VP F2/r3-F1); this completes the contract, it does not re-open it
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/s9-merge-gate/MERGE-GATE-implementer-20260714-022755.md
FROM: master.orchestrator-planner
TO: m-2.planner, m-6.implementer
CC: operator, master.orchestrator-reviewer, m-6.planner, m-2.implementer
SUBJECT: re-cut fresh at s9 close (the pre-serialization version was struck with the s11 dispatch) — the 8a contract (a parked gate across a schema bump) is PARTIALLY co-signed: the FLOOR is signed by both pairs (wake = migrate-then-validate via the §9 read-time migrator registry, then current-form validation; un-migratable → `held`/escalated, NEVER silently dropped, never auto-resolved) — and THREE members remain open, m-6-proposed but not co-signed; your two legs close them before the s11 PLAN locks 8a

**The three open members (m-6-proposed, awaiting exactly your two legs):**
1. **the `stale_schema` reason token** — m-2 confirms it into the bounce/reason vocabulary (registry grammar, their §17-family; now sequential on s9's landed registry — no concurrent contention);
2. **the bounded-choice set FROZEN at park time** — m-2 confirms the frozen-choice/migration interaction (what the migrator may and may not do to a frozen `agent_enum_pick` set across a bump);
3. **the choice-set-changed ⇒ bounce-as-stale + RE-ISSUE branch** — m-6.implementer adversarially reviews it (m-6.planner flagged this branch consequential and reserved it for exactly this review): the re-issue's exactly-once story, its interaction with the **s10-BUILT resummon dedupe keys** (seat + decision + cadence-slot — now landed, so state how a re-issued gate's keys relate to the original's: new decision identity vs same), and the never-silently-dropped guarantee under crash.

**One joint review may co-sign all three together** (the kickoff's allowance) — m-2's two confirms + m-6.implementer's branch review in a single co-signed return. **Available evidence now (wasn't at the pre-serialization cut):** s10 shipped the live park/wake FSM + the resummon keys (real parked gates existed across the v8 bump); the s9 evidence layer landed — cite any observed edges. The G4 cadence-config carry is s11 build scope, not this review's.

**Return path:** full design ritual (B12 declined — take the rounds the branch needs), unique sub-IDs, the co-signed 8a contract returns TO master → the s11 PLAN locks 8a only on it. The s11 dispatch is cut beside this (`s11-dispatch/PLAN-orchestrator-planner-20260714-023000`) with this gate NAMED — the s11 pair plans around it until your return lands.

ACTIONS_GIT_REF: none — a design dispatch (disk refs: this relay + one INDEX.md row timestamped 20260714-023001).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb` (`s9-close`, synced to frank-dev).
Next requested action: operator carries this to m-2.planner + m-6.implementer; the co-signed 8a contract returns TO master.
