## DESIGN-REVIEW - m-3.implementer re-review of decision-5 m-3 wording revision

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-fold-decision-5
PARENT_DISPATCH_ID: c5-fold-decision-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded re-review of prior claim-wording blocker; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c5-fold-decision-5/DESIGN-planner-20260702-135809.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-4.planner, m-6.implementer, m-4.implementer, m-7.planner, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: approve

I re-reviewed `c5-fold-decision-5/DESIGN-planner-20260702-135809.md` against the prior `must-revise`, the live m-3 design doc, the c5 checklist, and the m-4/m-6 sibling decision-5 seams.

The prior blocker is folded. This approval supersedes my m-3 `must-revise` at `c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134834.md` for the m-3 scan-side wording delta. It is still review-only: no PLAN, no IMPL, no `pcode/`, no runtime spike, no mechanism change, and no orchestrator ledger fold by this seat.

## Resolved finding

The decision-5 fold no longer claims absolute unforgeability. The previously blocked §7/§12/§13/§15 wording now says the ODB model-name exemption is governed-surface scoped: conductor-generated, non-lane-writable through the tool surface, no lane tool can mark arbitrary bytes exempt, and the D5 same-uid direct-store residual remains (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:122`, `:178`, `:204`, `:221`). That matches the ratified c5 checklist's required confusion-resistant/no-tool plus D5-residual shape (`master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md:27-31`; VP carry-forward in `.../RECONCILE-orchestrator-reviewer-20260702-131709.md:37-41`).

## Reconfirmed substance

- The exemption remains typed and narrow: ODB record, `model_name` field, destination `operator`, confidentiality class only.
- Safety/content still scans the whole ODB; secrets, PII, auth URLs, non-ODB model names, and lane-supplied exempt-marked content still block.
- Transport remains m-6-gated: non-away local render only; away-mode via opt-in away bridge plus egress gate.
- R2 stays untouched: this is egress confidentiality, not model-as-gate-input. The m-4 R2 half is already pair-approved, and m-6's typed ODB render half is already pair-approved.

Joint close still belongs to `master.orchestrator-planner`; this relay only approves the revised m-3 decision-5 leg.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-planner-20260702-135809.md` - OK.
- `sed -n '1,240p' master/relays/c5-fold-decision-5/DESIGN-planner-20260702-135809.md` - reviewed revised m-3 planner relay.
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md | sed -n '116,126p;172,181p;200,206p;218,223p'` - reviewed all four revised decision-5 locations.
- `rg -n "non-forgeable|forgery-proof|unforgeable|tamper-proof|tamper-resistant|structurally unbypassable|sole external sender|a lane (cannot|can't)|cannot mark|lane-supplied|non-lane-writable|D5 direct-store|D5 same-uid|confusion-resistant|no lane \\*tool\\*|no tool" master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md` - checked the survivor net; remaining `unforgeable` at §2:37 is the accepted trusted-courier control-flow KEEP, not the decision-5 blocker.
- `rg -n "c5-fold-decision-5|ODB model-name|model_name|R2 untouched|confidentiality|away-bridge|typed exempt" master/domains/m-4-routing-policy master/domains/m-6-human-surface-scheduler master/relays/c5-fold-decision-5` - checked sibling m-4/m-6 seams and approvals.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit by this seat, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
RELAY_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-140748.md`
DISPATCH_ROOT_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5`
Next requested action: m-3.planner may report this approval to the orchestrator; orchestrator owns final decision-5 ledger fold after considering all owner-pair legs.
