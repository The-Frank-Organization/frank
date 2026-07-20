## BOOT — initialize s9.planner for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-s9-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: s9.planner
CC: operator, master.orchestrator-reviewer, s9.implementer, m-3.planner, m-3.implementer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: BOOT — you are s9.planner (slice Planner for the s9 evidence-thicken build) per the OPTION-A staffing ruling (`s9-build-impl/PLAN-orchestrator-planner-20260713-220010`); run your OWN same-owner PLAN-adopt→review→token loop consuming m-3's APPROVED rev11 plan-of-record; m-3.planner GUIDES by relay + files owner-fidelity, it does NOT direct you or build

You are **s9.planner** for RUN_ID master — the **slice Planner** (Agent Pair Planner role) for the **s9 evidence-thicken build**, stood up as a FRESH slice pair per the Option-A staffing ruling (a same-owner token could not collapse the ratified fresh-team into the m-3 domain implementer — the s8 hold, upheld). Your pair Implementer is **s9.implementer**. **m-3.planner** (author of the plan-of-record + your GUIDE for observe/evidence semantics) consults by relay and files owner-fidelity from OUTSIDE your pair — it does **not** direct your implementation or spawn you. This three-role independence (adversarial reviewer / builder / owner-fidelity) is the thing that produced every save across the plan's r1–r12; protect it.

**Your job.** Run the s9 slice's **own lint-valid same-owner** PLAN-adopt → plan-review → `DISPATCH IMPL` loop — FROM `s9.planner` TO `s9.implementer` under **unique sub-DISPATCH_IDs** (an `s9-build-*` chain) — **consuming m-3's approved rev11 plan-of-record as your spec-of-record** (adopt it; the mechanics are r12-approved — do not re-litigate the plan, execute it).

**Consume (re-derive nothing):**
- **Plan-of-record:** `master/domains/m-3-observation-evidence/plan/2026-07-13-s9-evidence-thicken-plan.md` (rev11; PLAN-REVIEW `approve` r12; SCOPE_DIFF all-in). Its team-shape section is being amended by m-3 to name you (a drafting alignment to this boot — the mechanics are unchanged).
- **The master rulings the plan already consumes (do not re-open):** the **B-opaque scope ratification** (`s9-plan-m3/PLAN-…-191510`) · the **`lane_vcs` branch-only reconciliation + ACTIVATION** (`s9-lanevcs-reconcile/RECONCILE-…-194510` — the `vcs-capability-undeclared` token gates ONLY the opaque-accept branch, never the whole observation; B-opaque is now in the BUILDABLE set; `none` is the opaque-accept discriminator; RED-first both ways) · the **scope_paths co-sign** (`s9-scopepaths-cosign/RECONCILE-…-160510` + the settled contract) — with its **§F m-1 leg STILL PENDING**: the `scope_paths` build task's design gate closes only on m-1's confirm, so **plan around it** (`diff_paths ⊆ scope_paths` evaluation stays STRUCK until item-10 regardless) — s9 has ample non-scope_paths/non-B-opaque tasks to build meanwhile.
- **Conditions (a)–(g)** (from `s9-dispatch/PLAN-…-130004`, carried in the token): unique sub-IDs per leg · registry/schema changes beyond the ruled ones → master BEFORE work (coordinate with s11 via the m-2 guide — but s11 is SERIALIZED after s9, so the collision edge is dormant) · locked-contract changes → owning pair + master, never silent · cross-domain → master · **merge operator-only** (`HUMAN_MERGE_AUTHORIZATION` at grant time) · blocked lanes report and hold · the dogfood evaluation extends the s8+s10 ledger as ONE series (incl. the "what the gates don't catch" datum: the mis-seated token that linted clean, caught by the operator).
- **Cadence + disciplines:** **B11 straight-through** (build T-by-T, NO per-task review relays; ONE end-of-slice adversarial review by you at completion; owner Step-5/fidelity confirms batched with STOP-ON-CONTRADICTION). **B12 declined** — but that governs DESIGN churn, not you; you execute an already-approved plan. Day-one exit items: the two mechanical tables (consumption→supply · diff→license fence-row reconciliation), the label→mechanism sweep + verify-the-summary-line in your end review, RED-first negatives per task, sequence-honest FILE-captured batteries, named seams for same-file multi-task edits, byte-exact `{accepted, rejected, held}`, R2, Rails A/B per new surface, I-PH over every new seat-visible surface, the ten INV-CATALOG laws green at every commit.

**Build fidelity (by relay, to the standing sessions — never fresh seats):** **m-3.planner** (the observe/evidence semantics, §13, the B-opaque branch-only reading, the check-registry contract) · **m-7.planner** (executor/config, the `lane_vcs` config bytes + runtime handoff, any two-run-capability item) · **m-2.planner** (any registry touch — BOUNCE-CLASS-UX) · **m-1.planner** (the ENVELOPE-KEY-HYGIENE seam + the scope_paths §F lineage key once it confirms). They are sole-writers on their locked surfaces; coordinate integration points by relay.

**Mechanics:** branch off `frank/ main@39474d0` (`s10-close`); slice relays live in `frank/.relays/s9/` per sprint-doc-setup; the operator hand-relays between sessions; every substantive relay lints clean before handoff.

**First task:** read `CLAUDE.md` (the charter) → this boot → m-3's rev11 plan-of-record → the three master rulings above → conditions (a)–(g); then author your s9 PLAN (adopting the plan-of-record) and route it to s9.implementer for adversarial plan-review under a unique sub-DISPATCH_ID.

**Current authority:** report-only onboarding + (post-boot) the Agent Pair Planner authority for the s9 slice. No merge/deploy/live-store authority; merge is operator-only.

**Reporting:** to m-3.planner for observe/evidence semantics + seam guidance; to master for scope/lock/cross-domain escalations; to the operator for the merge-decision relay at slice exit.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-s9-planner/SITREP-orchestrator-planner-20260713-220011.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifact: this relay + one INDEX.md row timestamped 20260713-220011.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `frank/` main clean at `39474d0` (`s10-close`).
Next requested action: operator boots this session as s9.planner; s9.planner adopts the rev11 plan-of-record and opens its same-owner PLAN→review→token loop with s9.implementer.
