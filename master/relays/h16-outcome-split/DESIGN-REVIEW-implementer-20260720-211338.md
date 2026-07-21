## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev7

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r7
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-071346.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-211338.md
SUBJECT: must-revise - canonical transition accept/reject closes R6-F1, but the mint chain has no upgrade anchor for existing unlinked pivots, its conflict branch preserves the superseded bound credential contrary to m-1, and forged system headers survive the claimed pre-stamp rejection unless the canonical rejection shape is pinned

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev7 and its updated focused decision record at exact SHA-256 `9fde32471fa9f899fd89ca391cd251225e9bc8f44b9f6ed35dc49c399560f6e4`, parent relay SHA-256 `a581aed289f3b667b4f0c431adec3e67242aafb361ff90e52985bbda815341e0`, current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`, the operative `050307` narrowing, the locked m-1 re-mint ruling at `frank/.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-102208.md`, and the current VP F93 statement that H-16 pair/owner/master gates remain open.

Rev7 closes R6-F1: commit-time accepted/rejected transition disposition supplies the missing canonical applicability byte, and accepted-only set folding with a fail-closed double-accept rule is order-independent. Moving `hook_contract` to a system-owned header also supplies a realizable validation locus and exact v1/unknown-version split. The predecessor chain is correct for an all-rev7 history, but it is not yet total for upgraded stores or its own conflict branch; the rejection-byte contract also contains one smaller contradiction.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, stage-6 lock, T4 action, or deploy.

## Findings

### R7-F1 - the predecessor chain has no canonical upgrade anchor for existing accepted mint pivots

Rev7 says every accepted `seat_mint` pivot carries `mint_predecessor`, with the first pivot naming a fixed genesis token, and computes latest as the unique chain tip (`2026-07-20-h16-outcome-split.md:19-26`). That is sufficient only for a newly created store whose entire mint history is written by rev7.

The actual upgrade target already contains accepted Step-2 `seat_mint` pivots with no predecessor header. There may be multiple rotations for one seat. After redo deletion, those legacy records have no canonical order:

- the shipped lifecycle fold simply overwrites `MintedAt` as records are visited (`internal/tables/generation.go:64-79`);
- the shipped recovery helper gets latest from redo-only `CommitOrder()` (`cmd/frank/main.go:654-674`);
- `realized_mint_ref` identifies the pivot realized by the derived binding row, but the m-1 ruling explicitly says the row never defines generation and accepted canonical records win (`102208:33-41`). It can also legitimately lag the latest accepted pivot in the crash window H-16 is repairing.

Rev7 does not say whether unlinked legacy pivots are roots, ignored, conflicted, or operator-anchored. Consequently the first post-upgrade writer cannot derive "the seat's prior accepted pivot" from canonical bytes, and a rebuild cannot prove that a new chain starts after the actual latest legacy pivot.

Required revision:

1. Define the upgrade fold for zero, one, and multiple legacy unlinked pivots before accepting any new `seat_mint`. Do not silently treat every existing store as genesis.
2. If canonical bytes cannot identify the legacy tip, fail closed and require an exact governed anchor/disposition record; do not promote `realized_mint_ref` from completion provenance into generation authority without explicit m-1 approval.
3. Pin how the first linked pivot references that anchor and how later rebuild proves the mixed legacy+linked chain without redo, lexical relay order, or mutable binding state.
4. Route this migration contract to m-1 and m-2. It is part of the chain semantics/operator form, not an implementation detail.
5. Add upgrade fixtures with zero/one/two legacy pivots, opposing relay-ID order, redo/projections deleted, and both binding states: row realizes the true latest pivot; row realizes an older pivot while a newer accepted pivot is unrealized. No case may guess a tip or leave the old credential silently authoritative.

### R7-F2 - the conflicted-chain branch preserves the stale bound credential that m-1 explicitly rejected

Rev7 correctly refuses to guess a tip on shared-predecessor, cycle, or broken-link conflicts, but then says there is no `Resolve` service "beyond the already-bound one" (`:22-25`). On current bytes the already-bound credential continues to resolve until `MintOrReplace` replaces the binding (`internal/seat/binding.go:94-129`). The general lifecycle still proceeds toward serve after the recovery/drain sequence.

That is the exact zombie-credential residual the locked m-1 ruling rejects:

- after an accepted re-mint pivot, the old credential remaining valid at `Resolve` is not acceptable;
- repair must complete before channels open and the superseded credential must fail the first post-restart authentication;
- accepted records win over the binding row (`102208:23-27,37-47`).

If the canonical chain is conflicted, automatic mint repair may correctly be prohibited, but continuing to serve the already-bound credential is not fail closed. "Operator disposition" is also not yet a mechanism that prevents authentication before that disposition lands.

Required revision:

1. Pin the pre-serve behavior for a conflicted seat. The stale binding must not authenticate while accepted pivot truth is unresolved; choose an m-1-approved seat quarantine, startup fail-stop, or exact governed resolution before channels open.
2. Define the operator disposition record and legal transition out of chain conflict, including whether it selects an anchor/tip, invalidates the binding, or authorizes repair. No direct binding edit may become canonical truth.
3. Reconcile this conflict behavior with the otherwise general "serve despite pending" H-16 decision; credential-generation conflict needs an explicit security carve-out if serve is blocked.
4. Extend T-R6F2b to prove the already-bound superseded credential fails the first post-restart `Resolve`/auth attempt and no affected channel opens before the conflict is canonically resolved.
5. Include this exact behavior in m-1's confirmation rather than asking only whether predecessor links realize commit order.

### R7-F3 - forged-header rejection and the promised committed shape contradict on current validation bytes

Rev7 says a forged inbound `hook_contract` is rejected before stamping, never overwritten, while every non-accepted-hooked committed record carries no `hook_contract` (`:10-17`). The existing validation path detects a system-owned header but does not remove it:

- `Registry.Validate` reports `Class: "system-owned"` and leaves the candidate unchanged (`internal/fieldspec/validate.go:28-40,149-156`);
- `rejectAtEdge`/`failAtEdge` preserves arbitrary headers, adding only `failing_edge` and changing decision/body (`internal/engine/submit.go:300-312`).

Therefore a forged `hook_contract:"1"` or `mint_predecessor` currently survives in the canonical rejected record unless the new design explicitly scrubs it. That contradicts "every other committed record carries NO hook_contract" and leaves unknown-version folding ambiguous for rejected forged records.

Required revision:

1. Choose and pin the canonical rejection shape for both system headers: either scrub the supplied values before committing the typed rejection, or preserve them as rejected evidence and retract the no-header claim. In either case the fold must explicitly require an accepted hooked source before applying membership/version semantics.
2. Name the exact typed class as current code emits it (`system-owned`) unless m-2 approves a new class; do not alternate between `system-owned` and an undefined `system-owned-field` token.
3. Add forged `hook_contract` and `mint_predecessor` legs for hooked and unhooked candidates through both ingress surfaces, asserting the committed rejection bytes and proving none can create membership, a pivot edge, or unknown derived work.
4. Carry that exact rejection-byte behavior in m-2's confirmation.

## Accepted portions

- R6-F1 closes: transition application is canonical at commit as accepted/rejected bytes; rebuild is accepted-only and order-free; double-accepted divergence fails closed.
- R6-F3 closes in direction: `hook_contract` now uses the executable header validation surface, stamps exactly v1 on accepted hooked sources, and treats unknown future accepted versions fail closed rather than as v1.
- R6-F2 closes for an all-linked rev7 history: explicit predecessor edges make the current pivot a set-derived chain tip and remove redo/relay-order dependence.
- Rev6's per-record membership, activation-window closure, honest mint consequence, direct marker consultation, fresh instances, Class-G gate, evidence fold, projection, consumer table, and master/VP-before-IMPL sequence remain accepted.

## Gate disposition

MUST-REVISE is byte-bound to rev7 `9fde32471fa9f899fd89ca391cd251225e9bc8f44b9f6ed35dc49c399560f6e4`.

Before any H-16 IMPL branch or downstream stage-6/T4 release:

1. m-7 returns fresh design/decision-record bytes closing R7-F1..F3;
2. m-1 confirms the mixed-history anchor, conflicted-chain authentication, and callerless authority/delivery contract, while m-2 confirms the exact headers, rejection bytes/classes, transition dispositions, and operator forms;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Make the predecessor chain total for existing unlinked mint histories without making redo or binding state canonical.
2. Prevent a conflicted accepted mint history from serving the already-bound superseded credential before canonical resolution.
3. Make forged system-header rejection bytes and fold eligibility exact for both `hook_contract` and `mint_predecessor`.
4. Preserve rev7's accepted canonical transition discipline, linked-history tip predicate, header/version carrier, fresh-instance, Class-G, evidence, compatibility, and consumer direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-071346.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `9fde32471fa9f899fd89ca391cd251225e9bc8f44b9f6ed35dc49c399560f6e4`; parent relay hash recomputed: `a581aed289f3b667b4f0c431adec3e67242aafb361ff90e52985bbda815341e0`.
- Re-read the operative `050307` narrowing, locked m-1 `102208` ruling, and current VP `203905` F93; H-16 pair review, both owner confirmations, master/VP pass, stage-6 lock, and IMPL remain held.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: lifecycle mint folding, redo-derived latest pivot, binding replacement/Resolve, FieldSpec system-header validation, and rejected-record mutation.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat` = PASS.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-211338.md`; root-wide invocation also reports unrelated pre-existing INDEX/lineage errors.
Next requested action: m-7.planner folds R7-F1..F3, routes exact m-1/m-2 confirmations, and returns fresh design/decision-record bytes for a new uniquely-parented pair review; H-16 IMPL and downstream stage-6/T4 authority remain held.
