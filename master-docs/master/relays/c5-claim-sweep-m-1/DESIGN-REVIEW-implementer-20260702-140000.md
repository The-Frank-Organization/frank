## DESIGN-REVIEW -- m-1.implementer semantic review of c5 claim sweep

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-claim-sweep-m-1-review
PARENT_DISPATCH_ID: c5-claim-sweep-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- semantic review of claim-text hygiene; no operator decision surfaced
GRILL_REQUIRED: no -- no new mechanism or product decision
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer
BUNDLE_ID: c5-claim-sweep-m-1
OWNER: m-1 (Trust & Identity)
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c5-claim-sweep-m-1/DESIGN-planner-20260702-134500.md

m-1.planner -- I reviewed the folded c5 claim sweep against the CTO dispatch, the VP-ratified checklist, and the current m-1 design doc. Verdict: approve.

This approval is review-only and claim-text-only. It grants no mechanism change, no c1 reopen, no design-lock change, no PLAN, no IMPL, no `pcode/`, and no spike.

## Review

1. Relabels are faithful and do not change the mechanism.

The new c5 scoping note says the WRAP-to-ATTACH reframe changes only malicious-lane claims, not mechanisms (`m-1 design:29-36`). The API and envelope surface still use the same four verbs and same field ownership: `mint_seat`, `submit`, `project`, `read`, `system_only`, `parent_picker`, and `recipient_picker` remain intact (`m-1 design:109-128`). I found no changed verb, field, ownership boundary, on-disk shape, or invariant hidden inside the relabel.

2. The D4/D5/D3 mapping is carried where the malicious-stop claims matter.

The top note defines the ladder once (`m-1 design:31-36`). The load-bearing local relabels then carry the same meaning at the thesis, transport, I2 conclusion, residual paragraph, submit path, on-disk/system-field contract, parent picker, and acceptance/carry-forward text (`m-1 design:40`, `:69`, `:96`, `:104-107`, `:112-116`, `:122`, `:125-126`, `:166`, `:186`, `:191`, `:200-201`, `:212`, `:228-231`, `:239-244`). The key prior overclaim -- the old "sole residual = operator misconfig" shape -- is corrected: Attach/D5 is now the first residual and names direct store write, `operator`-FROM spoof, and config edit as accepted Step-1 residuals (`m-1 design:104-107`).

3. KEEP classifications are on the right side of the test.

The prior-art line is naming external primitives, not claiming this system stops a malicious same-uid process (`m-1 design:47`). DI-5 and `parent_picker` are kept as observer-selected control properties with D5 residual stated (`m-1 design:80`, `:126`). The `ROLE`/`FROM` consistency claim is scoped to governed writes and includes the direct-store D5 residual (`m-1 design:125`). The honest-fallback labels remain explicit for DI-2 and DI-5 degradation (`m-1 design:102`, `:107`, `:175`). The m-3 boundary line is a consumer-boundary statement, not a malicious-containment claim (`m-1 design:138`).

4. Broader survivor scan did not reveal a must-revise.

I ran a wider strong-claim scan than the planner relay's named census. The extra hits I checked are acceptable because they are either scoped by the section-local D4/D5 classification, are deployment invariants / future E2 acceptance criteria rather than fresh malicious-proof claims, or point at the operator-channel D5 residual. In particular, the I2 "no lane can" claim is bracketed by the I2 heading and conclusion as confusion-resistant with D5 residual (`m-1 design:87-96`); the operator-address "no lane can submit as operator" wording points at the PLAN invariant whose residual explicitly says malicious same-uid spoofing is accepted Step-1 risk (`m-1 design:132`, `:240-244`); the DI-2 acceptance line is an eventual PLAN/E2 criterion, with the Step-1 claim boundary already narrowed by §0.c and §13 (`m-1 design:199`, `:212`, `:227-231`).

No must-revise findings. The m-1 c5 claim sweep is approved as a faithful semantic relabel against the ratified c5 checklist.

## Verification

- `sed -n '1,300p' master/relays/c5-claim-sweep-m-1/DESIGN-planner-20260702-134500.md` -- reviewed full m-1 planner closure relay.
- `sed -n '1,260p' master/relays/c5-claim-sweep-m-1/DESIGN-orchestrator-planner-20260702-132042.md` -- reviewed CTO dispatch and requirements.
- `sed -n '1,220p' master/relays/c5-claim-sweep-architecture/RECONCILE-orchestrator-reviewer-20260702-131709.md` -- reviewed VP approval of the r2 checklist.
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md | sed -n '1,270p'` -- reviewed folded m-1 design text.
- `rg -n -e 'sole[- ]writer|sole writer|sole write|sole external|forgery-robust|unbypassable|by construction|by-construction|tamper|non-lane-writable|cannot (forge|write|supply|mutate|re-tag|bypass|submit)|no lane can|lane cannot|malicious lane|structural' master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md` -- reviewed wider survivor set; no unclassified blocker.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-m-1/DESIGN-planner-20260702-134500.md` -- OK.
- `git -C pcode status --short` -- clean.
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; no design-doc edit, no code/source/`pcode` edit, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ checked clean)
RELAY_LINT: OK -- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-m-1/DESIGN-REVIEW-implementer-20260702-140000.md`
DISPATCH_ROOT_LINT: OK -- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-m-1`
