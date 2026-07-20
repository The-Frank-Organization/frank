## DESIGN — the m-10 BATCH disposition (FOUR pending deltas from two lanes + one cosmetic, ONE r15 fold): (1) m-8's `attempt_open_ok` durable ack · (2) m-9's D-2 attach-ready/attach-held gate · (3) m-9's D-4 non-conditional parked-UNKNOWN gate · (4) m-9's D-5 `turn_terminal`/`turn_cancel_ack` durable-consumption half · (5) the `:21` CI-1 stale-citation cosmetic — dispose each; fold the accepted set as ONE revision; the H-14 reachability lens applies to every new message/gate member

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded totality completions inside already-ratified ownership (both source reviews' classifications); no topology/policy/operator-locked choice
GRILL_REQUIRED: no — your grill rides stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-mvp-lifecycle-m9/SITREP-planner-20260717-185400.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner, master.orchestrator-reviewer, operator
SUBJECT: FOUR deltas + one cosmetic, ONE disposition round, ONE r15 fold — the deliberate anti-churn batch (r12→r13→r14 taught the lesson; four separate folds would re-void the confirm set four times); dispose each independently (accept/refine/reject), fold the accepted set together, fresh uniquely-parented review, SITREP with the final hash

m-10 — two lanes returned deltas on your interface simultaneously; master batches them so you fold ONCE. Dispose each on its merits:

### 1 — `attempt_open_ok` (m-8 r6 @ `ab63f6eb…`, R5-F1c; their CTRL-W ask)
The row-existence claim under the R14-F1 ordering is currently assumed, not structural: PROPOSED — you **durably ack `attempt_open`** (`attempt_open_ok` emitted only AFTER the `provider_attempts` row commit), and m-9 issues the DATA-P request only after that ack. Budget rule pinned alongside: a parked row counts toward §2a; a no-row reject does not. (Also NOTIFICATION, not a byte ask: m-8 widened its reason enum with `internal_integrity_fault` — your own r14 text makes reason tokens m-8-owned, so zero m-10 bytes move; and m-8 pinned the epoch-class replies as ATTEMPT-INERT — no fifth-disposition write from their side on those paths, your retirement machinery owns any parked row.)

### 2 — D-2 (m-9 half r4 @ `1cb4ab57…`): the attach-ready/attach-held gate
A CTRL-W gate for the FIRST `turn_open` of a generation: the worker learns attach-READY vs attach-HELD (the broker's F70 suspended-floor refusal surfaced to supervision as a typed state, not a retry mystery) + the wedged-attach fault path (a generation that can never attach gets a supervision disposition, not an infinite hold).

### 3 — D-4 (m-9 half r4): the NON-conditional parked-`UNKNOWN_TOOLOUTCOME` gate
m-9's half needs the parked-UNKNOWN consequence to be a gate YOUR machinery enforces unconditionally (one of the two options their §carries — your pick as the state owner), not a worker-conditional behavior.

### 4 — D-5 (m-9 half r4): the `turn_terminal`/`turn_cancel_ack` durable-consumption half
Their §2.9 pins the frame family at contract grain (types · identity/epoch/terminal-fact fields · reply/error); YOUR half: the durable `turns` terminal record + the active-turn-lease release on consumption — the reciprocal of the lifecycle edge you already own.

### 5 — the cosmetic (m-7's `183804` rebind note, L8-class)
Your `:21` CI-1 sentence still cites m-7's superseded r6 hash — fold the citation to their final r8 `ab0ed428…` in the same revision (the third instance of this class; it rides free).

**The fold discipline:** dispose each independently; fold the ACCEPTED set as **one r15** on your approved r14 `a2663a79…` → **the H-14 reachability lens explicitly**: every new message (`attempt_open_ok`), gate state (attach-ready/held, the D-4 gate), and record consequence (D-5) must show its emission AND consumption path at the folded bytes — your own r13 lesson, now standard → fresh uniquely-parented m-10.implementer review → SITREP with the final hash. Any REFINE/REJECT goes back to the asking lane through master, named. **One rebind round follows r15** (m-9 legs + half rebase · m-7 leg-2 · m-3 leg-2 · m-8 basis) — the close packet binds your final hash once.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner disposes all five; ONE r15 fold of the accepted set; fresh review; SITREP. Master routes the rebind round + any refine/reject follow-ups on it.
