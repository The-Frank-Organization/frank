## DESIGN-REVIEW - m-5.implementer review of the Step-3 pinned-ceiling nod

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m5-ceiling-review-r3
PARENT_DISPATCH_ID: step3-amend-m5-ceiling
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded m-5 owner-position review; positive-tool branch may require Master/VP/operator routing outside this relay
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-amend-m5-ceiling-host
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-amend-m5-ceiling/DESIGN-planner-20260715-144500.md
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-1.planner, m-7.planner, m-9.planner
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

DESIGN_REVIEW_VERDICT: must-revise

m-5.planner - I reviewed the addressed DESIGN `DESIGN-planner-20260715-144500.md`, the master ask `133500`, the m-1 owner return `124031`, m-7's earlier return `060542`, the canonical ceiling contract `643dd7c2...`, and the later VP review `step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-144149.md`. Verdict: must-revise.

The governance-model position in item (b) is acceptable only as a proposed backlog carry. The item (a) nod is not approvable as written. Under the unchanged canonical contract, absence of current-active proof still fails closed to `tool -> none`; "single frozen run" prevents later divergence, but it does not prove that the run-start manifest value equals the authoritative committed-chain head at bind time.

This review grants no contract hash change, no design-of-record fold, no first-stage lock, no positive Step-3 tool authority, no PLAN, no T4 token, no `frank/` edit, and no m-10 consumer approval.

## Findings

### F1 - The "pinned == current-active" nod changes freshness semantics without changing the contract

The canonical contract is exact: before any m-9 tool dispatch, m-10 validates the bound ceiling as present and fresh; freshness holds iff `policy_digest` and `config_generation` equal the current-active Layer-1/current trusted-config values, and stale/malformed/unresolvable artifacts fail closed to the m-5 floors, including `tool -> none` (`2026-07-15-ceiling-artifact-contract.md:40-43`). The file hash still recomputes to `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.

Your rev says the pinned generation at run-start "IS the current-active generation" and therefore the current-active comparison is vacuously satisfied (`DESIGN-planner-20260715-144500.md:26-34`). That only proves no later generation change if the frozen-run premise holds. It does not prove the initial bound value came from the authoritative committed config chain, nor does it name the writer/provenance evidence m-10 checks before accepting that value as current-active.

m-1's owner return is the opposite of a positive proof: it confirms the generation property, but says the read path to the current-active stamp has no already-landed, packet-compliant mechanism, and that disposition #2 means m-10 cannot establish current-active generation, so deny all tool dispatch (`SITREP-planner-20260715-124031.md:36-46`, `:48-50`). The VP review makes the same blocker explicit and says the unchanged contract plus m-1 return require deny-all unless m-5 authors changed run-start freshness/provenance bytes or an E1 existing bind-time mechanism is proven (`RECONCILE-orchestrator-reviewer-20260715-144149.md:31-42`).

Required revision: choose one coherent branch.

1. Strict packet-preserving branch: under unchanged `643dd7c2...`, no current-active proof means `tool -> none` for Step-3; no positive pinned ceiling.
2. Positive-tool branch: author a real m-5 contract amendment that replaces or qualifies the current-active freshness rule for the one-turn frozen-config MVP. It must name the run-start provenance/writer, how the value is tied to the authoritative genesis-anchored committed-chain head, what m-10 validates, and fail-closed behavior for missing/mismatched proof; then produce a new exact hash for m-10 and Master/VP to consume.
3. Existing-mechanism branch: provide E1 proof of a currently landed, packet-compliant bind-time mechanism that satisfies the existing `current active` bytes. I did not find one in the owner returns; m-1 explicitly says it is absent.

### F2 - Audit cannot substitute for the stale-ceiling fail-closed requirement

Your residual paragraph says a missed mid-run tightening is bounded partly because every tool call is m-3 evidence and therefore caught downstream (`DESIGN-planner-20260715-144500.md:32`). That may be useful evidence, but it is not the contract's preventive behavior. The contract says stale is fail-closed to zero tool dispatch, not "dispatch and audit later" (`2026-07-15-ceiling-artifact-contract.md:41-43`). The locked m-5 design similarly makes old or partial records under-authorize rather than widen (`2026-06-30-v3-archetype-system-design.md:82-95`).

Required revision: keep audit-universal as evidence, not as a substitute for freshness. If the design accepts a stale-permissive window under an operator frozen-config promise, say that is a deliberate semantic amendment/residual-risk branch and route it through the changed-contract path above.

### F3 - The governance-model carry is fine as backlog, but cannot justify the current `bash`/cwd tool boundary

I agree with the disposition of item (b) as "accept in principle, carry, do not fold" (`DESIGN-planner-20260715-144500.md:36-45`), and the backlog entry confirms it is a proposed m-5 refinement rather than locked source (`FRANK-HARDENING-BACKLOG.md:26`). But the current relay uses audit-universal/capability-coarse/irreversibility-sharp-end as part of the safety story for positive Step-3 tools while hard sandbox and irreversibility gating defer.

That interacts with the VP's F34: an allowed `bash` can leave cwd, use absolute paths, follow symlinks, run network clients, spawn arbitrary programs, and perform destructive/external effects; audit after execution is not prevention (`RECONCILE-orchestrator-reviewer-20260715-144149.md:44-52`). A future m-5 reversibility organizing principle is plausible, but it cannot be consumed now as proof that the present Step-3 tool surface is cwd-scoped or reversibility-bounded.

Required revision: state that item (b) is non-operative backlog and not a basis for approving positive `bash`/ambient-tool execution in this amendment. If the positive-tool branch remains, the actual shell authority must be handled by Master/VP/operator per the VP finding, not buried in the m-5 nod.

## Accepted portions

- The rev correctly preserves m-5 as policy owner and treats the answer as feeding m-10 DESIGN, not bypassing the m-10 DESIGN -> review -> SITREP -> Master/VP lock sequence (`DESIGN-planner-20260715-144500.md:47-60`).
- The rev correctly carries the m-1 requirement that any future comparand be genesis-identity anchored, not a bare integer (`DESIGN-planner-20260715-144500.md:32`; `SITREP-planner-20260715-124031.md:27-28`).
- The governance-model refinement is worth carrying as an m-5 design-backlog item, provided it remains proposed and non-operative until a proper amendment/review/grill path exists.

## Revision bar

Return a revised m-5 answer that:
1. Does not call unchanged `643dd7c2...` positive-tool faithful unless it supplies an E1 bind-time proof satisfying the current-active bytes.
2. Otherwise chooses strict fail-closed `tool -> none` under the unchanged contract, or authors a changed one-run freshness/provenance contract with a new hash.
3. Keeps audit-universal and reversibility semantics out of the preventive freshness proof unless they are routed as an explicit semantic amendment/residual-risk branch.
4. Preserves item (b) as proposed backlog only, with no implied approval of the current unsandboxed `bash` boundary.

## Verification

Pre-write review evidence:
- `test -f master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-144500.md` - resolved the shortened pointer to the exact relay.
- `sed -n '1,220p' CLAUDE.md` - read current charter and Step-3 scope.
- `sed -n '1,260p' master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-144500.md` - read addressed DESIGN.
- `sed -n '1300,1325p' master/relays/INDEX.md` - checked current relay trail through `144500`.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-144500.md` - OK.
- `sed -n '1,220p' master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-144149.md` - read latest VP review.
- `sed -n '1,220p' master/relays/step3-amend-m5-ceiling/RECONCILE-orchestrator-planner-20260715-133500.md` - read master ask.
- `sed -n '1,240p' master/relays/step3-amend-m5-ceiling/SITREP-planner-20260715-124031.md` - read m-1 owner return.
- `sed -n '1,220p' master/domains/m-5-workflows-archetypes/design/2026-07-15-ceiling-artifact-contract.md` - read canonical contract.
- `shasum -a 256 master/domains/m-5-workflows-archetypes/design/2026-07-15-ceiling-artifact-contract.md` - `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- `rg -n "audit-universal|capability|irreversible|H-11|bash|read/write/edit/bash/apply_patch|cwd-scoped" master/FRANK-HARDENING-BACKLOG.md master/ARCHITECTURE.md master/README.md master/RECONCILE.md` - checked backlog/source posture for item (b) and the current tool-boundary claims.

Post-write verification:
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-144650.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/step3-amend-m5-ceiling` - OK.
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-144650.md` - initially found a single section-symbol in `Next requested action`; patched to ASCII and reran clean.
- `rg -n "20260715-144650|step3-amend-m5-ceiling-review-r3|DESIGN-REVIEW-implementer-20260715-144650" master/relays/INDEX.md master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-144650.md` - relay line 6 and INDEX line 1312 present.
- `git status --short` from `/Users/jack/Programming/harness` - `fatal: not a git repository (or any of the parent directories): .git`.
- `git -C frank status --short` - no output.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no domain-doc edit, no source/code/`frank/` edit, no PLAN, no IMPL, no runtime action.
FINAL_GIT_STATUS_SHORT: harness root is not a git repo (`fatal: not a git repository (or any of the parent directories): .git`); `frank/` clean (`git -C frank status --short` produced no output).
RELAY_LINT: OK `master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-144650.md`
DISPATCH_ROOT_LINT: OK `master/relays/step3-amend-m5-ceiling`
Next requested action: m-5.planner revises the section 5 nod. If positive Step-3 tools remain desired, route changed freshness/provenance semantics through a reviewed m-5 contract hash and keep the unsandboxed-shell authority visible to Master/VP/operator.
