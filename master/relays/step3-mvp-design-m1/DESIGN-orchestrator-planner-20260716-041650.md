## DESIGN dispatch — §7 stage-1 owned contract: the connector SECRET-BOUNDARY delta (F57 strength) + the SEAT IDENTITY / CREDENTIAL-LIFECYCLE semantics (F60: one broker-held credential per LOGICAL seat) (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-1 owned interface contract under the ratified amendment; the operator gates return at the Master+VP interface-lock, not per-artifact
GRILL_REQUIRED: no — stage-1 owner contracts carry pair review + consumer confirmation; the grills ride the stage-4/5 build lanes (§7)
DESIGN_DOC_ID: step3-mvp-design-m1-secret-boundary-seat-identity
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: author the m-1 stage-1 contract: the connector secret-boundary delta at the ratified F57-narrow strength + the F60 seat identity/credential-lifecycle semantics (what a LOGICAL seat is across worker generations; epoch-fenced replacement, no implicit new identity) — pair-reviewed final bytes, consumers confirm

m-1 — the Step-3 MVP amendment is **ratified + operative** (`master/STEP-3-MVP-AMENDMENT.md` r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`; operator-ratified 2026-07-16, VP approve `…-035505`). Your charter carries the delta. This dispatch opens your **§7 stage-1 owned contract** — two coupled artifacts. (Your prior `124031` config-generation leg is OBE with the ceiling stand-down; nothing carries from it.)

### Author (you own these bytes; m-1.implementer pair-reviews the FINAL bytes)
1. **The connector secret-boundary delta (§1/§2, at the ratified F57-NARROW strength):** the m-8 provider-key custody AND the m-9 seat-credential custody stated as **non-injection into the enumerated surfaces** (model context · local-tool arguments · `bash` environment/files/argv · inherited FDs · logs) + **separate-process placement as accidental-disclosure reduction** — explicitly NOT a hard unreadability boundary (same-user inspection of peer-process state = the named unsandboxed MVP residual; the OS boundary is Step-4 H-12) — + the §2 minimum hardening list as design requirements (sanitized env; no secrets in argv/ordinary env; close-on-exec on internal IPC descriptors; tool subprocesses inherit no m-8/m-10/seat-broker channel; no secret-bearing crash dumps; private runtime dir; the operator's ambient provider credentials not inherited by m-9). The §10 sentinel test proves accidental-disclosure absence ONLY.
2. **The seat identity + credential-lifecycle semantics (grill #3 / F60 — the invariant is operator-locked; you author its SEMANTICS):** what a **LOGICAL seat IS across worker generations** — **one stamped credential per LOGICAL seat**, broker-held (the m-7 broker contract, dispatched in parallel `step3-mvp-design-m7`, realizes custody UNDER your semantics), never copied into worker generations; a replacement worker is **`turn_epoch`-fenced and mints NO implicit new identity** (worker-per-seat is rejected: mint-before-serve makes per-generation minting an operator act; identity churn wrecks lineage accountability); m-10 launches/supervises with **opaque references only, never credential bytes** (the launch-custody answer). Define: seat identity persistence across restart/replacement; the mint/rotate relationship to the existing seat_mint/re-mint machinery (your locked s6 §F/F.1 + re-mint-binding work is the substrate — extend, don't fork); what the epoch-bound USE capability is in identity terms (a delegation of the seat's channel, never a second identity).
3. **Consistency rail:** the Standing Identity-Seam Rail applies — channel-stamped FROM unchanged; the broker is custody plumbing, NOT new identity machinery; no home-grown security primitive; structured-so-it-could-later-carry-attestation preserved.

### Boundaries
No conductor byte/member change; no registry/store change in this DESIGN (any such need routes back). Consumers to confirm: m-7 (broker realization), m-8 (key custody), m-9/m-10 (lifecycle fit). No DESIGN-lock, PLAN, T4 token, or code is authorized by this dispatch.

### Return path (§7 stage 1)
m-1.planner authors the DESIGN parented to THIS dispatch → m-1.implementer DESIGN-REVIEW as a uniquely-parented child (fresh review on any byte revision) → report-only SITREP to master naming the approved bytes + hash → consumer confirmations route on master's direction. The Master+VP interface-lock is the gate; no self-declared lock.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-1.planner opens the DESIGN on this dispatch (grounding: the ratified amendment §1/§2/§7 + your charter delta + your locked s6 §F/§F.1 + re-mint machinery + m-7's parallel dispatch); pair review; report-only SITREP to master.
