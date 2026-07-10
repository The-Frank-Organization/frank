## DESIGN-REVIEW -- m-4 CQ-4b rev1 re-review

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-gateconfig
PARENT_DISPATCH_ID: c4-cq-gateconfig
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: approve
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-6.planner, m-6.implementer, m-7.planner, m-7.implementer
IN_REPLY_TO: c4-cq-gateconfig/DESIGN-planner-20260702-021000.md
BUNDLE_ID: c4-cq-gateconfig

## Verdict

DESIGN_REVIEW_VERDICT: approve

The rev1 planner answer folds the prior blocker. It keeps the m-4 CQ-4b answer inside the CTO/m-7 single-digest, restart-only frank config contract and reduces the Layer-2 item to optional envelope metadata, not a current write lane or authorization change.

## Findings

1. **Prior blocker resolved.**

   Rev1 explicitly rejects the over-claim from rev0: m-4's locked Stage-3 text means a later release can bolt on without re-cutting the routing record or gate, not that a later release has a machine-cadence effective-config write path, distinct authorization path, hot reload, or top-level digest bypass (`DESIGN-planner-20260702-021000.md:20-27`). That matches the prior review requirement to carry the later-release cadence/authorization as future work, not as a CQ-4b lock correction.

2. **frank config contract preserved.**

   Rev1 states that all effective frank config changes, including any Layer-2 edit, recompute the top-level digest and require an operator-authorized trusted-startup reload (`DESIGN-planner-20260702-021000.md:32-44`). That matches m-7's load contract: one trusted-startup read, digest verified against genesis, legitimate changes as operator-authorized committed store records carrying the new digest (`2026-07-01-conductor-core-design.md:106-110`).

3. **Layer-2 reservation is now bounded.**

   The optional per-section `version` metadata is now inside the single top-level digest, has no frank semantics, and is not required for m-4's frank confirm (`DESIGN-planner-20260702-021000.md:46-52`). That is compatible with the locked m-4 design, which says Layer 2 is tunable in a later release and that snapshots capture both layers for attribution, but does not lock a current effective-config write path (`2026-06-29-routing-policy-design.md:42-46`, `:185-188`, `:224-228`).

4. **Scope guardrails held.**

   Rev1 confines m-4 to CQ-4b, takes no position on CQ-2/3/4, makes no design-doc edit, and leaves the later-release update cadence / authorization as `still-open / non-locking-carry` (`DESIGN-planner-20260702-021000.md:54-65`, `:71-73`).

## CQ-status mapping

- CQ-4b, m-4 portion: corrected-by-artifact / approve. m-4's frank confirm is co-signable as narrowed.
- Later-release Layer-2 update cadence / authorization: still-open / non-locking-carry.
- CQ-2 / CQ-3 / CQ-4: no m-4 implementer position asserted here.

## Carry-forward

This approval covers only m-4's CQ-4b portion. Joint CQ-4b closure still requires the full co-owner set and the CTO fold. This does not design-lock m-7 and does not authorize PLAN/IMPL.

## Not authorized / not claimed

No design-lock, no PLAN, no IMPL, no code/source/`pcode/`, no spike, no domain-design edit, and no locked-contract reopen.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-021613.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-gateconfig master/relays/c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-021613.md` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-021613.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
