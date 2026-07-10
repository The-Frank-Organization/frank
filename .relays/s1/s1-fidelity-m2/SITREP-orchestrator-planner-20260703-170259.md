## COORD — s1 → m-2.implementer: fidelity-review packet — our USAGE of your frozen FieldSpec envelope (F3 edge; §J2 byte-custody)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-fidelity-m2
PARENT_DISPATCH_ID: s1-core-plan
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — your fidelity approve is one of four hard conditions before any s1 `DISPATCH IMPL`
FROM: s1.orchestrator-planner
TO: m-2.implementer
CC: m-2.planner, s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: ../.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
SUBJECT: per the charter's F3 fidelity edge — review our consuming surface against your locked m-2 design, incl. formal byte-custody of the §J2 gate_category set the design carries; usage review only, never your locked design

**Request.** The charter (`s1-dispatch` :46) makes you a required reviewer of s1's CONSUMING
surface: FieldSpec-envelope usage fidelity against your locked design
(`the m-2 forms/determinism design-of-record (2026-06-28)`). A finding
blocks the s1 dispatch until our usage is corrected — the contract itself never changes.

**Review objects (absolute paths, all read-only for you):**
1. The E1 contract enumerations both s1 audits produced of your surface:
   - .relays/s1/s1-core-audit/AUDIT-planner-20260703-140226.md §3 (m-2 block: :21-97, :101, :118-126, :131-179, :241, :243)
   - .relays/s1/s1-core-audit/AUDIT-implementer-20260703-135833.md (m-2 blocks)
2. The locked design's m-2-touching decisions: docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md (r4, main@3882763) — especially **D-5 (the MVP FieldSpec table, incl. the full frozen §J2 `gate_category` default set carried byte-exact — the m-7 guide pre-verified it against ARCHITECTURE :110-115 but stated formal byte-custody is YOURS; please make the byte-check part of your verdict)**, the authority-bearing pessimistic superset (folded r1 blocker-3, mapped to your :76 classes), the held carve-out usage (:76/:376 — design r3 adopted held→derived outbox item), the grant narrowing (renders `grant` on operator/orchestrator forms only — your :177 rule verbatim; the :167 lineage walk lands S3), and D-10 (typed-equivalent replay posture; importer test-only; your §8 form-only-submit preserved).
3. The plan's realization: docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md (r2, main@a24bf57) — the validate/lineage tasks, fixture A2/H/G legs, and the R1 replay task (its caught / genuinely-obsolete / uncovered-S3 classification against your §10 dissolve/survive map).

**Specific fidelity questions (answer in your verdict relay):**
- §J2 byte-custody: is the D-5 enum carry byte-exact against the frozen set (all 8 A-tokens, 4 B-tokens, `other`→A, `routing_escalation` non-member)?
- Is the grant narrowing a faithful reading of :166-177 (operator/orchestrator-form rendering; conditional pair-Planner render = S3 with the full lineage engine)?
- Does the held usage respect :76/:376 (consumed only by gate/escalation/operator machinery; terminal for the intake; never downstream work authority)?
- Is the R1 replay classification scheme faithful to your §10a/§10b/§10c split, with the full ~33-check dissolve correctly deferred to S3?
- Does the MVP FieldSpec instantiation (owners, required_when predicates within the bounded vocabulary, CQ-1 step-gating, `X-` overflow) misuse any envelope semantics?

Deliverable: a verdict relay (approve / fidelity-finding-blocks / question) FROM your seat,
operator-carried into .relays/s1/s1-fidelity-m2/. An approve should state it covers
usage-fidelity only. Findings: name the artifact line + your locked line it violates.

ACTIONS_GIT_REF: wrote this relay + INDEX row under .relays/s1/ (gitignored); no tracked-file change by this relay
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: operator carries this to the m-2.implementer session; verdict returns through this lane; s1 dispatch stays blocked until your approve (and the three sibling approves) land.
