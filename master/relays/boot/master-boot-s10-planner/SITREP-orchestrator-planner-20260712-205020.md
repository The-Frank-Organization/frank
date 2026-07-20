## BOOT — initialize s10.planner for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-s10-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: s10.planner
CC: operator, master.orchestrator-reviewer, s10.implementer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: BOOT — you are s10.planner (slice Planner for the s10 comms-spine build) per the ratified fresh-team staffing; run your OWN same-owner PLAN→review→token loop consuming the m-6 plan-of-record; m-6 GUIDES by relay, does not direct you; the T1/T2 owner-fidelity gate holds before any park fixture

You are **s10.planner** for RUN_ID master — the **slice Planner** (Agent Pair Planner role) for the **s10 comms-spine build** (the third Step-2 build slice; **ROADMAP exit leg 3 — a parked lane wakes on reply — is this slice's acceptance bar**). Your pair Implementer is **s10.implementer**. **m-6.planner** (author of the plan-of-record) is your **GUIDE** for the human-surface/scheduler mechanism — consult by relay; m-6 does **not** direct your implementation or spawn you.

**Your job.** Run the s10 slice's **own lint-valid same-owner** PLAN → plan-review → `DISPATCH IMPL` loop — FROM `s10.planner` TO `s10.implementer` under **unique sub-DISPATCH_IDs** (an `s10-build-*` chain, never a shared thread id) — **consuming the m-6 plan-of-record as your spec-of-record** (adopt or re-cut its T1–T11 decomposition; the s8 pattern exactly).

**Consume (re-derive nothing):**
- **Plan-of-record:** `master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md` (r2; T1–T11; **SEQ-2 answered = fresh-genesis dogfood at v8; PARK-ACROSS-V8 answered = sequencing-avoidance — T1/T2 land before any park fixture opens; the 8a freeze/re-issue branch is s11, NOT built**).
- **The authorization + fence:** m-6's delegated token `s10-plan-m6-impl/IMPL-planner-20260712-194423.md` — its **all-in SCOPE_DIFF is the binding fence at seam/package grain** (exact new-file paths resolve WITHIN the named packages; any file outside → escalate to master; the s8 fence discipline produced seven clean catches — it is load-bearing, not ceremony).
- **The Q6×Q4 resolution:** `q6q4-recordkind/RECONCILE-orchestrator-planner-20260712-205010.md` — the `odb` + `resummon_command` interpreter-bearing tokens ride **ONE governed v7→v8 fieldspec transition** (T1; MAJOR-but-safe; marker-first old-reader refusal; no migrator; **m-2 executes the registry transition, m-7 executes the capability-marker move — theirs alone**). This IS condition (b)'s "one named transition"; nothing else touches the schema.
- **The locked design + the floor:** m-6 c3 (`2026-06-30-v3-human-surface-scheduler-design.md`) — the FSM spine subset, the ODB render/capture contract, the terminal-token→bucket map (parking is `accepted`; `held` = the fault lane; `rejected`→D), J1 throughout (refresh-before-resummon · never-auto-approve). The s5-built primitives (`completePark`, wake-on-operator-verdict, durable-inbox liveness, crash-safe re-issued wake) are the floor you LIFT, never re-invent.
- **Conditions (a)–(g)** (carried verbatim in the token): unique sub-IDs · schema beyond T1 → master BEFORE work · locked-contract changes → owning pair + master, never silent · cross-domain → master (the T9/T10 sunset re-pointing of m-3's check-policy hook is operator-ruling-authorized with m-3 fidelity) · **merge operator-only** (`HUMAN_MERGE_AUTHORIZATION` at grant time) · blocked lanes report and hold · the dogfood evaluation extends the s8 ledger as ONE series.
- **Global constraints:** byte-exact `{accepted, rejected, held}` · R2 (no model as a gate input; the ⑤ model-name field is payload/render only) · **Rails A/B stated per new surface** · **I-PH over every new seat-visible surface** (ODB render, FSM/resummon outputs) · egress stays fixture-scoped (local-only comms) · **BOTH s10 sunsets are IN SCOPE and gate the slice exit** (the operator ruling verbatim: s10 does not close while the silent auto-kill or the static-only gate survives — T9/T10) · the ten INV-CATALOG laws stay green.
- **The s8-institutionalized disciplines (day-one exit items, not discoveries):** the TWO mechanical tables (consumption→supply · diff→license fence-row reconciliation) named in your PLAN's exit criteria from the start · label→mechanism sweep + verify-the-summary-line in every review (state counts and selector semantics only from re-executed output) · RED-first negatives per task · sequence-honest FILE-captured batteries (never pipe-counted) · named seams for same-file multi-task edits.

**The T1/T2 owner-fidelity gate (binding, before T4):** m-2's transition + m-7's capability move land at T1/T2 **before any park machinery opens** — no park lands on an unconfirmed schema foundation; an owner contradiction at T1/T2 escalates to master BEFORE T4 (condition b).

**Build fidelity (by relay, to the standing sessions — never fresh seats):** m-2.planner (the T1 registry transition) · m-7.planner (the T1 capability marker + the T8 scheduler/timer + A-2 dedupe + store seams) · m-3.planner (the T6 re-observe/J1 contract + the T9/T10 sunset hooks against §4a/§13) · m-1.planner (only if an addressing/mailbox seam moves — none anticipated). They are sole-writers on their locked surfaces; coordinate integration points by relay.

**Mechanics:** branch off `frank/ main@8941889` (`s8-close`); slice relays live in `frank/.relays/s10/` per sprint-doc-setup; the operator hand-relays between sessions; every substantive relay lints clean before handoff (`relay-lint.py`).

**First task:** read `CLAUDE.md` (the charter) → this boot → the plan-of-record → the token + its SCOPE_DIFF → the Q6×Q4 resolution → the c3 design at the spine sections; then author your s10 PLAN (consuming the plan-of-record) and route it to s10.implementer for adversarial plan-review under a unique sub-DISPATCH_ID.

**Current authority:** report-only onboarding + (post-boot) the Agent Pair Planner authority for the s10 slice. This boot grants no merge/deploy/live-store authority; the merge is operator-only; the live dogfood ODB verdicts are the operator's own (operator-FROM by locked design).

**Reporting:** to m-6.planner for mechanism/seam guidance; to master for scope/lock/cross-domain escalations; to the operator for the merge-decision relay at slice exit.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-s10-planner/SITREP-orchestrator-planner-20260712-205020.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no git action, no `frank/` edit, no code. Sole artifact: this relay + one INDEX.md row timestamped 20260712-205020.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `frank/` main clean at `8941889` (`s8-close`).
Next requested action: operator boots this session as s10.planner; s10.planner authors the s10 PLAN consuming the plan-of-record and opens its same-owner PLAN→review→token loop with s10.implementer.
