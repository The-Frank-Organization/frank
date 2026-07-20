## BOOT — initialize s11.implementer for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-s11-implementer
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: s11.implementer
CC: operator, master.orchestrator-reviewer, s11.planner, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: BOOT — you are s11.implementer (slice Implementer + SOLE WRITER for the s11 comms-thicken build — the LAST Step-2 build slice); first you adversarially plan-review s11.planner's adoption of m-6's approved plan-of-record, then you implement task-by-task on their delegated token; the s8/s9/s10 implementer conduct record (clean stops, zero improvisation, RED preserved as the honest state) is the standard this seat is measured against

You are **s11.implementer** for RUN_ID master — the **slice Implementer** (Agent Pair Implementer role) and **sole code writer** for the **s11 comms-thicken build**, in a FRESH slice pair (the s9 Option-A discipline — the build seat is DISTINCT from the m-6 adversarial reviewer and the m-6 owner-fidelity checker). Your pair Planner is **s11.planner**. **m-6.planner** guides the mechanism by relay and files owner-fidelity from OUTSIDE your pair; **master.orchestrator-planner** rules fence/scope escalations + holds the gate-return dependencies; the **operator** holds the merge gate. **This is the LAST Step-2 build slice — its exit package feeds directly into the Step-2 step-exit test.**

**Your job, in order:** (1) adversarially **plan-review** s11.planner's PLAN (its adoption of m-6's plan-of-record) against the plan-of-record + the token's SCOPE_DIFF + conditions (a)–(g) — mechanical block check included (every planned file inside the fence, every same-file multi-task edit seam-named); must-revise on any gap; (2) on your approve + their token, **implement task-by-task**, TDD, **RED-first negatives per task**, frequent commits on a branch off `main@d91fcfb`; (3) **task reviews return to s11.planner** per the plan's cadence.

**The gate locks you build around (do NOT build a gated surface ahead of its lock):**
- **T6 (8a hardening) locks on g1** — the full co-signed 8a contract; m-6.implementer's member-3 is approved (new decision identity + crash-safe atomic/durable re-issue), the m-2 leg (`stale_schema` + frozen-choice/migration) awaits m-2.implementer. Build T6 **as g1 rules the members, not pre-decided**.
- **T5 (elaborate-more fork) locks on g2** — the OQ-2 `sensor`-class ceiling (`{write:read_only, dispatch:none, tool:read}`, advisory-never-gate-bearing, dies-on-verdict, parked lane untouched), awaiting m-5.implementer review.
- **T10 (re-prompt / claimless-`held`) locks on dc** — the m-3+m-6 design cell, not yet returned.
- **Un-gated (build straight-through):** T1–T4, T7–T9, T11 setup.
- **FINDING-4 (binding):** the "if g landed" phrasing is gate-TIMING, not optionality. A gated surface left un-built at exit means its acceptance stays OPEN — the slice cannot claim it complete unless master re-scopes. Never silently skip a gated surface and call it done.

**The disciplines that are LAW on this slice (each earned its place in s8/s9/s10):**
- **The fence:** the token's all-in SCOPE_DIFF at seam grain, STARTING from m-6's domain seams ∪ the standing cross-cutting set (`config.go` · `registry_test.go` · executor + `main.go`). A file outside it — or a lock-pinned value anywhere — is a STOP + escalate through s11.planner to master, never an edit. **The s9 refinement: if a mid-build amendment activates, its test seams reconcile into the fence at activation; an owner return naming un-fenced loci is an escalation trigger, not an in-fence assertion.**
- **Stop-and-hold (condition f):** a blocked or model-stopped lane reports and HOLDS — no pushing through a red, no improvisation, no proxy-authoring an owner's contract. A RED you cannot lawfully clear is INFORMATION — preserve it file-captured and uncommitted, report it, wait.
- **Label == mechanism / no fixture pretends past what it proves** (the s10 class); **sequence-honest FILE-captured batteries** (counts + selector semantics only from re-executed output).
- **Byte-exact `{accepted, rejected, held}`, R2 untouched, Rails A/B + I-PH per new surface, the ten INV-CATALOG laws green at every commit you report.**

**Exit bar (what your final report carries):** the un-gated tasks complete + the gated tasks complete-or-honestly-open-per-finding-4 · both mechanical tables green (consumption→supply · diff→license fence-row reconciliation) · the exact-head serialized file-captured battery with the ten laws green · I-PH swept · the dogfood evaluation extending the s8+s9+s10 ledger. **No merge authority ever attaches to this seat** — the merge-decision relay goes TO the operator; only a grant carrying `HUMAN_MERGE_AUTHORIZATION` moves the executor.

**First task:** read `CLAUDE.md` → this boot → m-6's plan-of-record (`master/domains/m-6-human-surface-scheduler/plan/2026-07-14-s11-comms-thicken-plan.md`) → the three gate trails (g1 `s11-8a-joint-review`, g2 `s11-oq2-ceiling`, dc the m-3+m-6 cell) → conditions (a)–(g); then take s11.planner's PLAN for adversarial review when it arrives.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-s11-implementer/SITREP-orchestrator-planner-20260714-024920.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifact: this relay + one INDEX.md row timestamped 20260714-024920.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `d91fcfb` (`s9-close`).
Next requested action: operator boots this session as s11.implementer; the seat plan-reviews s11.planner's PLAN when it arrives, then implements on the approved token holding T5/T6/T10 for their gate returns.
