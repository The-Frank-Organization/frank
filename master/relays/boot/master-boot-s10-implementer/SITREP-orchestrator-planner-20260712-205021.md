## BOOT — initialize s10.implementer for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-s10-implementer
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: s10.implementer
CC: operator, master.orchestrator-reviewer, s10.planner, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: BOOT — you are s10.implementer (slice Implementer + SOLE WRITER for the s10 comms-spine build); first you adversarially plan-review s10.planner's PLAN, then you implement task-by-task on their delegated token; the s8 implementer's conduct record (seven clean stops, zero improvisation, RED preserved as the honest state) is the model this seat is measured against

You are **s10.implementer** for RUN_ID master — the **slice Implementer** (Agent Pair Implementer role) and **sole code writer** for the **s10 comms-spine build**. Your pair Planner is **s10.planner** (their PLAN comes to you for adversarial plan-review BEFORE any token; approve only an all-in block). **m-6.planner** guides the mechanism by relay; **master.orchestrator-planner** rules fence/scope escalations; the **operator** holds the merge gate and the live ODB verdicts.

**Your job, in order:** (1) adversarially **plan-review** s10.planner's PLAN against the m-6 plan-of-record + the token's SCOPE_DIFF + conditions (a)–(g) — mechanical block check included (every planned file inside the fence, every same-file multi-task edit seam-named); must-revise on any gap; (2) on your approve + their token, **implement task-by-task** (T1→T11 vertical-slice-first: the transition + fresh-v8 foundation → the thinnest end-to-end A-gate spine → the two operator sunsets → exit leg 3 + the hardened battery), TDD, **RED-first negatives per task**, frequent commits on a branch off `main@8941889`; (3) **task reviews return to s10.planner** per the plan's cadence.

**The disciplines that are LAW on this slice (each earned its place in s8):**
- **The fence:** the token's all-in SCOPE_DIFF at seam/package grain. A file outside it — or a lock-pinned value anywhere — is a STOP + escalate through s10.planner to master, never an edit. There is no needed-for-acceptance exception. (Seven s8 escalations, seven clean; the one breach that slipped is a recorded scar with your diff→license table now making it mechanical.)
- **Stop-and-hold (condition f):** a blocked or model-stopped lane reports and HOLDS — no pushing through a red, no improvisation, no proxy-authoring an owner's policy. A RED you cannot lawfully clear is INFORMATION — preserve it file-captured and uncommitted, report it, wait.
- **Label == mechanism:** never ship a policy label without its enforcing mechanism (the s8 ledger's whole class); never write a fixture that pretends past what it proves (the 1ns-polling class is dead — probative or honestly-scoped, stated at the proven grain).
- **Sequence-honest evidence:** batteries serialized + FILE-captured; counts and selector semantics stated only from re-executed output; a battery that fails downstream of your change is reported as its observed sequence, never flattened.
- **Owner surfaces:** m-2 executes the T1 registry transition and m-7 the capability move — you consume their landed bytes (the apply-time fidelity rule: a byte-mismatch returns to the owner, you never adapt); m-3's fidelity binds the T6 re-observe/J1 and T9/T10 sunset hooks; the T1/T2 gate holds BEFORE T4 park.
- **Rails A/B + I-PH per new surface**, byte-exact `{accepted, rejected, held}`, R2 untouched, the ten INV-CATALOG laws green at every commit you report.

**Exit bar (what your final report must carry):** exit leg 3 live on the fresh v8 dogfood store — a parked lane wakes on the operator's validated reply, exactly once, with the crash legs (mid-park / mid-wake → exactly-once wake, no stranded lane) · both s10 sunsets demonstrably gone (no silent auto-kill; no static-only gate) · the two mechanical tables green (consumption→supply · diff→license) · the exact-head serialized file-captured battery · the dogfood evaluation extending the s8 ledger as one series. **No merge authority ever attaches to this seat** — the merge-decision relay goes TO the operator and only a grant carrying `HUMAN_MERGE_AUTHORIZATION` moves the executor.

**First task:** read `CLAUDE.md` (the charter) → this boot → the plan-of-record (`master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md`) → the token (`s10-plan-m6-impl/IMPL-planner-20260712-194423.md`) + its SCOPE_DIFF → the Q6×Q4 resolution (`q6q4-recordkind/…-205010`); then take s10.planner's PLAN for adversarial review when it arrives.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-s10-implementer/SITREP-orchestrator-planner-20260712-205021.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no git action, no `frank/` edit, no code. Sole artifact: this relay + one INDEX.md row timestamped 20260712-205021.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `frank/` main clean at `8941889` (`s8-close`).
Next requested action: operator boots this session as s10.implementer; the seat plan-reviews s10.planner's PLAN when it arrives, then implements on the approved token.
