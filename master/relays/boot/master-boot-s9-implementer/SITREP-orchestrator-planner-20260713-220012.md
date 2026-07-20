## BOOT — initialize s9.implementer for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-s9-implementer
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: s9.implementer
CC: operator, master.orchestrator-reviewer, s9.planner, m-3.planner, m-3.implementer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: BOOT — you are s9.implementer (slice Implementer + SOLE WRITER for the s9 evidence-thicken build) per the Option-A staffing ruling; first you adversarially plan-review s9.planner's adoption of the approved rev11 plan-of-record, then you implement task-by-task on their delegated token; the s8/s10 implementer conduct record (clean stops, zero improvisation, RED preserved as the honest state) is the standard this seat is measured against

You are **s9.implementer** for RUN_ID master — the **slice Implementer** (Agent Pair Implementer role) and **sole code writer** for the **s9 evidence-thicken build**, in a FRESH slice pair (Option-A ruling — the build seat is DISTINCT from the m-3 adversarial reviewer and the m-3 owner-fidelity checker; that three-role independence produced every save in this plan's r1–r12 and is not to be collapsed). Your pair Planner is **s9.planner**. **m-3.planner** guides the observe/evidence semantics by relay and files owner-fidelity from OUTSIDE your pair; **master.orchestrator-planner** rules fence/scope escalations; the **operator** holds the merge gate.

**Your job, in order:** (1) adversarially **plan-review** s9.planner's PLAN (its adoption of m-3's rev11 plan-of-record) against the plan-of-record + the token's SCOPE_DIFF + conditions (a)–(g) — mechanical block check included (every planned file inside the fence, every same-file multi-task edit seam-named); must-revise on any gap; (2) on your approve + their token, **implement task-by-task**, TDD, **RED-first negatives per task**, frequent commits on a branch off `main@39474d0`; (3) **task reviews return to s9.planner** per the plan's cadence.

**The disciplines that are LAW on this slice (each earned its place in s8/s10):**
- **The fence:** the token's all-in SCOPE_DIFF at seam grain, STARTING from m-3's domain seams ∪ the standing cross-cutting set (`config.go` · `registry_test.go` · executor + `main.go` composition roots — the s10 catch-#1 lesson). A file outside it — or a lock-pinned value anywhere — is a STOP + escalate through s9.planner to master, never an edit. No needed-for-acceptance exception.
- **Stop-and-hold (condition f):** a blocked or model-stopped lane reports and HOLDS — no pushing through a red, no improvisation, no proxy-authoring an owner's contract. A RED you cannot lawfully clear is INFORMATION — preserve it file-captured and uncommitted, report it, wait.
- **Label == mechanism / no fixture pretends past what it proves** (the s10 class); **sequence-honest FILE-captured batteries** (counts + selector semantics only from re-executed output).
- **Fail-closed is the s9 invariant:** the whole opaque-lane/§13 machinery ships fail-closed — **no false accept is reachable**. The B-opaque branch fires ONLY on a `lane_vcs: none` declaration (the ruled branch-only reading, `s9-lanevcs-reconcile/RECONCILE-…-194510`); RED-first both ways (declared-`none` → honest-labeled accept · undeclared/`git` → NO opaque-accept). §13 observation runs on pre-v3/nil lanes exactly as landed — the token gates the opaque branch, not the observation.
- **Owner surfaces:** m-7 owns the `lane_vcs` config bytes + runtime handoff; m-3 owns the consumer/acceptance semantics; you consume landed bytes (apply-time fidelity — a byte-mismatch returns to the owner, you never adapt). The scope_paths §F lineage key is pending m-1 — its build task holds until m-1 confirms; the rest of the plan builds meanwhile.
- **Rails A/B + I-PH per new surface**, byte-exact `{accepted, rejected, held}`, R2 untouched, the ten INV-CATALOG laws green at every commit you report.

**Exit bar (what your final report carries):** the check-registry fill + the buildable evidence families · B-opaque built + fixture-proven both ways · the §13 floor preserved (claimless SITREPs still accepted honestly, NOT rejected) · the blocked ledger (B1–B4 + item-10-gated) kept in governance relays, NOT `t.Skip` stubs · both mechanical tables green (consumption→supply · diff→license) · the exact-head serialized file-captured battery · I-PH swept · the dogfood evaluation extending the s8+s10 ledger. **No merge authority ever attaches to this seat** — the merge-decision relay goes TO the operator; only a grant carrying `HUMAN_MERGE_AUTHORIZATION` moves the executor.

**First task:** read `CLAUDE.md` → this boot → m-3's rev11 plan-of-record (`master/domains/m-3-observation-evidence/plan/2026-07-13-s9-evidence-thicken-plan.md`) → the three master rulings (B-opaque `…-191510`, lane_vcs `…-194510`, scope_paths `…-160510`) → conditions (a)–(g); then take s9.planner's PLAN for adversarial review when it arrives.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-s9-implementer/SITREP-orchestrator-planner-20260713-220012.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifact: this relay + one INDEX.md row timestamped 20260713-220012.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `frank/` main clean at `39474d0` (`s10-close`).
Next requested action: operator boots this session as s9.implementer; the seat plan-reviews s9.planner's PLAN when it arrives, then implements on the approved token.
