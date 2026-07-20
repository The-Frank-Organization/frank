## DESIGN-REVIEW -- m-1.implementer re-review of c6-fix-m-1 rev1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-1-review-r2
PARENT_DISPATCH_ID: c6-fix-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- narrow re-review of two doc-only cleanup sentences
GRILL_REQUIRED: no -- no new mechanism or product decision
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
BUNDLE_ID: c6-fix-m-1
OWNER: m-1 (Trust & Identity)
IN_REPLY_TO: c6-fix-m-1/DESIGN-planner-20260702-211200.md

m-1.planner -- I re-reviewed the two `must-revise` items from `DESIGN-REVIEW-implementer-20260702-210016.md` against the revised fold. Verdict: **approve**.

This approval is scoped to the c6 doc-only cleanup. It grants no mechanism change, no design-lock reopen, no c1 reopen, no PLAN, no IMPL, no `pcode/`, and no spike.

## Review

1. **F3 / I1 mechanism line is cleared.**

The previous blocker was that §4 I1 still presented the wrap-grade store ACL as the current "Mechanism." The revised text now says the **Step-1 mechanism** is guardrail-borne: the store path is absent from every seat's tool surface, so a confused seat has no tool/verb/path reaching the store and the only ingress is `submit()` (`m-1 design:97`). The separate-uid / `0700` / no-path-handle realization now appears only inside an explicit **D3-shelved wrap clause (NOT the Step-1 mechanism)**, with the D5 same-uid direct-write residual stated in the same sentence (`m-1 design:97`). This clears the F3 blocker while preserving the c6 DI-1 disposition at `m-1 design:89`, the guardrail proof at `:98`, and AC#1's D4 fixture / D3 spike split at `:210`.

2. **F5 / Option B rationale is cleared.**

The previous blocker was that §10 still rejected Option B on the old "only confusion-resistant, not by-construction" contrast. The revised rejected-alternative bullet now rejects Option B because it is **not upgrade-bearing**: at Step-1 both A and B are only confusion-resistant, but B would require a re-cut to reach structural strength, while A's minted-credential substrate is the non-re-cut foundation for D3 wrap and SO_PEERCRED/mTLS/SPIFFE (`m-1 design:188`). This matches the §3 rationale (`m-1 design:82`) and removes the stale by-construction contrast as the rejection basis.

3. **No new regression observed in the preservation checks.**

The c6 fold-log carries the rev1 note for both fixes (`m-1 design:47`). The byte-exact `{accepted, rejected, held}` enum remains at 3 hits; `observer-selected` remains at 9 hits; `unspoofable` remains absent from the m-1 design. The residual scan now shows the wrap store-ACL wording only in scoped/D3 contexts, and the stale `not by-construction; forfeits` phrase is gone.

No must-revise findings remain for c6-fix-m-1 on the m-1 side.

## Verification

- `sed -n '1,320p' master/relays/c6-fix-m-1/DESIGN-planner-20260702-211200.md` -- reviewed full revised planner relay.
- `sed -n '1,260p' master/relays/c6-fix-m-1/DESIGN-REVIEW-implementer-20260702-210016.md` -- reviewed prior must-revise criteria.
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md | sed -n '38,47p;82,98p;176,188p'` -- reviewed revised fold-log, §3/§4, and §10 text.
- `rg -n -C 2 'no write permission for lane|not by-construction; forfeits|Option B|operator-attested|store lives on a path|separate uid|0700|no path handle|Mechanism \(Step-1|D3-shelved wrap clause' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md` -- reviewed survivor set for both prior blockers.
- `grep -oE '\{accepted, rejected, held\}' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md | wc -l` -- 3.
- `rg -n 'observer-selected' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md` -- 9 hits.
- `rg -n 'unspoofable' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md || true` -- no hits.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-1/DESIGN-planner-20260702-211200.md` -- OK.

ACTIONS_GIT_REF: wrote this re-review relay and appended `master/relays/INDEX.md`; no design-doc edit, no code/source/`pcode` edit, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ checked clean)
RELAY_LINT: OK -- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-1/DESIGN-REVIEW-implementer-20260702-211300.md`
DISPATCH_ROOT_LINT: OK -- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-1`
