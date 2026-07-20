## RECONCILE - revise: c5 step-(c) decomposition needs stronger owner-review and decision-4 split

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c5-decomp
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c5-decomp/SITREP-orchestrator-planner-20260702-041701.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer
SUBJECT: Revise the c5 step-(c) dispatch shape before routing

## Verdict

VERDICT: revise

The decomposition is directionally right: step (c) should be a cross-doc claim-hygiene and remaining-decision-fold pass, not a mechanism redesign or code/PLAN/spike opening. But I do not concur with the dispatch shape as written.

Two revisions are required before dispatch:

1. Do not make the domain claim-sweep "CTO-proposes + owning-planner-confirms" only. The global claim-sweep is exactly the NO-GO's failed claim class, and it touches locked domain docs. It needs owning-pair adversarial confirmation, at least an Implementer `DESIGN-REVIEW` / review-only approve of each domain's semantic relabels. Use light ceremony and batching where appropriate, but do not bypass the pair reviewer.

2. Do not defer decision ④ entirely to step (d). The away-token mechanism is a §2C build-step item, but the operator decision itself must be folded now as a non-locking build-carry / recorded constraint in the relevant locked docs. Step (c) explicitly says "fold the five operator decisions"; step (d) says fix §2C items at their build step. Split it: record the decision and boundary now; defer detailed mechanism/proof to the away-bridge build step.

## Answers To Q1-Q4

**Q1 - rigor/lift:** Claim-sweep is relabel hygiene, but it is high-risk hygiene because it corrects the adversarial overclaim that caused the NO-GO. Approved shape: CTO may prepare candidate relabels and a semantic sweep checklist; owning planner folds or confirms; owning implementer gives a review-only semantic approve. Full-pair is mandatory for decision-folds ③/⑤ and the decision-④ build-carry record. No grill unless a new product/mechanism decision appears.

**Q2 - light-domain scope:** Sweep all six domain docs plus `ARCHITECTURE.md`; do not spot-check m-3/m-4/m-5/m-6 out of scope. The lighter domains can run a batched, low-ceremony review if their hit counts stay small, but they still need an owner-pair response because their locked claims are part of the build-ready record. m-1 and m-2 should get the heavier pass because they carry the main trust/identity and schema overclaim load.

**Q3 - decision ④:** Do not fold the away-token mechanism now. Do fold the recorded operator decision now as a non-locking, §2C build-step carry in m-1/m-6 and the integrated ledger: rotate `decision_id`, burn prior nonces, re-observe current state, bounce/hold if changed. Detailed design, fixture proof, and adversarial review remain deferred to step (d), before park/wake or away-bridge ships.

**Q4 - ownership:** CTO may edit `ARCHITECTURE.md` directly and may propose exact domain relabel patches. CTO should not directly mutate locked domain docs as the final owner-confirming act. Domain docs should be owner-folded or explicitly owner-accepted, with implementer review. This preserves the no-proxy-edit bar while still letting the CTO drive the cross-doc consistency pass.

## Required Dispatch Shape

Dispatch c5 as a small set of explicit work lanes:

1. `c5-claim-sweep-architecture`: CTO edits `ARCHITECTURE.md`; VP reviews.
2. `c5-claim-sweep-m-1` and `c5-claim-sweep-m-2`: CTO candidate sweep plus owning planner fold/confirm plus implementer semantic review.
3. `c5-claim-sweep-light`: either one batched relay for m-3/m-4/m-5/m-6 or four tiny relays, but each domain must have owner-pair confirmation of its own claim text.
4. Decision-fold lanes:
   - ③ RAISE-ONLY A/B + known-A detector: CTO/§J + m-6, owner-authored and implementer-reviewed.
   - ⑤ ODB model-name carve-out: m-3 + m-6 + m-4, owner-authored and implementer-reviewed.
   - ④ away-token rotate+re-observe: record now as non-locking §2C build-carry in m-1/m-6 and the integrated ledger; mechanism deferred to step (d).
5. Byte-consistency re-verify after folds: CTO runs the cross-doc check; VP reviews before step (c) is marked closed.

## Evidence

- `master/DESIGN-REVIEW-2026-07-01.md:147-152` defines §2B as a global claim sweep and says only malicious-lane claims collapse, not mechanisms.
- `master/DESIGN-REVIEW-2026-07-01.md:224-227` defines step (c): global claim sweep plus fold the five operator decisions and the confused-agent guardrails into locked docs, then re-verify byte-consistency.
- `master/DESIGN-REVIEW-2026-07-01.md:154-161` defines §2C away-token sibling-burn/restart-binding as a must before its build step, not disappeared.
- `master/READINESS-REGISTER.md:346-361` records decisions ③, ④, and ⑤ with owners/surfaces: ③ CTO/§J + m-6; ④ m-1 + m-6; ⑤ m-3 + m-6 + m-4.
- `master/README.md:34-38` says step (c) remaining folds and full cross-doc sweep still owe, then §2C-at-build-step, then Step-1 PLAN.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-decomp/SITREP-orchestrator-planner-20260702-041701.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-decomp` - OK
- `nl -ba master/DESIGN-REVIEW-2026-07-01.md | sed -n '115,160p'` - reviewed §2A/§2B and start of §2C.
- `nl -ba master/DESIGN-REVIEW-2026-07-01.md | sed -n '154,164p'` - reviewed away-token §2C build-step boundary.
- `nl -ba master/DESIGN-REVIEW-2026-07-01.md | sed -n '217,236p'` - reviewed re-baseline step (c)/(d)/(e) ordering.
- `nl -ba master/READINESS-REGISTER.md | sed -n '330,370p'` - reviewed operator decisions ②-⑤ and retracted gate note.
- `nl -ba master/README.md | sed -n '1,42p'` - reviewed live dashboard state for c4 closed and c5 owed work.
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no domain doc, architecture, source, code, `pcode`, PLAN, IMPL, runtime spike, or design mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner revises the c5 dispatch set to include owner-pair semantic review for domain claim-sweeps and a decision-④ split: record now as a non-locking build-carry, mechanism deferred to step (d).
