## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev5

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r5
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-064221.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-064856.md
SUBJECT: must-revise - rev5's canonical direction, fresh attempt identity, and Class-G retry gate are sound, but the epoch relation is not reconstructible after redo deletion, recovery can create hooked decisions before activation, and callerless mint repair neither follows the stated ordering nor preserves an older delivered credential

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev5 and its updated focused decision record at exact SHA-256 `326991dc049b7294951efb385298d9ef1a4a800fd6318a97c31478acc8066911`, parent relay SHA-256 `9e197a394831593b2a5d2b2d3d7cffe8e93746c67131c2e34222f7e07d13786f`, current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`, the operative `050307` narrowing, and the latest VP `051057` statement that H-16 return 2 remains open.

Rev5 closes three important rev4 defects in direction: open state no longer depends on a projection intent, attempt and park transitions have fresh predecessor-scoped identities, and Class-G has a same-process retry-before-replay truth gate. Its `realized_mint_ref` classification is also useful durable evidence. Three load-bearing contracts remain false or under-specified on current bytes.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, or deploy.

## Findings

### R5-F1 - the epoch relation is not reconstructible from canonical records after redo deletion

Rev5 defines open as a hooked record committed after one canonical `h16-epoch` record in commit order, while adding no source field or canonical sequence (`2026-07-20-h16-outcome-split.md:26-32`). T2/T14 then delete all redo and projections and require `tables.Build` to recover that relation from `Store.Records()` alone (`:72-74`).

Current storage semantics cannot do that:

- `tables.Build` gets records from `Store.Records()`, separately gets `Store.CommitOrder()`, and delegates ordering to `recordsInCommitOrder` (`internal/tables/tables.go:116-131`);
- `Store.CommitOrder()` derives its entire order from redo entries (`internal/store/store.go:119-135`);
- with redo absent, `recordsInCommitOrder` returns the record slice unchanged (`tables.go:134-137`);
- the remaining canonical filenames are random relay IDs, not a commit sequence, so their file order cannot establish whether a hooked record preceded or followed the epoch.

An epoch record is canonical, but the load-bearing relation to it is not. Randomized relay IDs can sort a pre-epoch decision after the epoch or a post-epoch decision before it, producing false grandfathering or false open work.

Required revision:

1. Make H-16 membership itself canonical and order-independent, for example with an exact source-record generation/epoch field, or add an explicitly designed canonical sequence relation. Do not rely on redo-derived commit order.
2. Route the exact source-field/schema delta to m-2 if that is the selected carrier; rejecting it as "no gain" is no longer supportable.
3. Keep T2/T14's redo/projection deletion, and additionally construct pre/post records with relay IDs whose lexical order opposes commit order. The fold must still classify both correctly.
4. Specify migration/grandfather behavior from canonical bytes alone, including stores that contain no H-16 membership carrier.

### R5-F2 - H-16 activation occurs after recovery can commit hooked decisions

Rev5's lifecycle is recovery -> Ready -> `engine.New` -> hook binding -> epoch commit + drain (`:54-57`). That matches the current main ordering: `frankrecover.RunWithProcessor` runs first (`cmd/frank/main.go:219-262`), then missing mint bindings are repaired (`:263-268`), and only afterward are the loop and hooks built (`:281-323`).

The recovery processor can commit a new accepted hooked record and run `completeSeatMint(rec, false)` before any epoch exists (`:238-257`). If that post-H16 recovery attempt fails or the process exits after the decision commit but before activation, the next startup sees the record on the legacy side of rev5's boundary and reports `complete`. The design therefore grandfathers work produced by the new binary during its own activation window.

Required revision:

1. Install an order-independent H-16 membership boundary before any H-16-aware recovery path can commit a hooked decision, after the minimum store/genesis validation required to write it; alternatively stamp every recovery-created source record with canonical membership.
2. Define the cuts before/after activation, before/after recovery decision commit, and before/after the recovery hook effect. No new-binary record may become legacy merely because startup failed before serve.
3. Add recovery fixtures for an unconsumed gate resolution and seat mint processed by the first H-16 startup, with failure/crash after the decision and before hook completion. The next restart must project `pending`/`failed`/`unknown` under the H-16 fold, never grandfathered `complete`.

### R5-F3 - the callerless mint exemption relies on false ordering and overstates delivery safety

Rev5 says callerless repair is disjoint because the startup scan occurs before the loop and the §6 drain folds markers first (`:40-42`). The latter is incompatible with both its own lifecycle and current code: `completeMissingSeatMintBindings` runs at `main.go:263-268`, while loop construction, hook binding, and the proposed drain happen later. The recovery processor's `completeSeatMint(rec, false)` is earlier still. Neither callerless path can rely on a later drain having folded unresolved markers.

The claim that callerless repair rotates a credential "NO ONE ever received" and harms no delivered credential is also too broad. If seat credential A was delivered, then accepted pivot B committed but its mint effect did not run, callerless repair of B calls `MintOrReplace` and invalidates delivered A while creating B without returning it to a caller. `realized_mint_ref != B` proves B unrealized; it does not prove that the currently realized credential was never delivered or that replacing it is delivery-inert.

Required revision:

1. Pin one actual startup order. Either fold canonical marker state before every callerless repair, or make each repair path directly consult the canonical marker fold before `MintOrReplace`; do not claim a later drain already happened.
2. State the credential consequence honestly: repair of a newer accepted pivot can invalidate an older delivered credential and leave the replacement available only through the ruled admin-recovery path.
3. Obtain m-1's exact confirmation of that authority and delivery consequence before H-16 lock. If m-1 does not authorize it, route the startup paths through the marker/evidence machine or hold them for operator disposition.
4. Test: delivered credential A -> accepted pivot B -> crash before B's effect -> unresolved marker and restart. Prove callerless repair skips unresolved B; after any authorized repair, prove A's invalidation and B's actual admin recovery path.
5. Reconcile the recovery-processor path separately: it can run before the general scan and before the proposed marker drain, so the disjointness check must cover it too.

## Accepted portions

- R4-F1 closes in direction: rev5 removes `store.Intent` and keeps redo/projections non-canonical. Only the replacement epoch relation remains defective.
- R4-F2 closes: fresh marker relay IDs, explicit predecessor resolutions, and instance-scoped one-shot operator records represent retries and reopened parks without unbounded precedence.
- R4-F3 closes in design: `classGDirty`, retry-before-every-reply, restart pre-serve coverage, and a diagnostic row give truthful same-process and zero-traffic behavior.
- `realized_mint_ref == pivot` is sound completion evidence for the pivot, and `realized-undelivered` is an honest failed-class result for a caller-coupled effect whose extras were lost.
- Canonical outcome enum, fail-closed projection, caller-present marker law, consumer table, dropped heal nudge, focused decision record, INV-CATALOG boundary, and master/VP-before-IMPL sequence remain intact.

## Gate disposition

MUST-REVISE is byte-bound to rev5 `326991dc049b7294951efb385298d9ef1a4a800fd6318a97c31478acc8066911`.

Before any H-16 IMPL branch:

1. m-7 returns fresh design/decision-record bytes closing R5-F1..F3;
2. m-1 confirms the exact callerless mint authority/delivery contract and m-2 confirms the exact operator form plus any selected source membership field;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Make H-16 membership canonical and reconstructible with redo/projections absent and adversarial relay-ID ordering.
2. Activate that membership before recovery can create any new hooked decision, with first-start crash cuts.
3. Make both callerless mint paths consult unresolved marker state in their actual startup order and describe their credential-delivery effect accurately.
4. Carry exact m-1/m-2 confirmations before lock.
5. Preserve rev5's accepted fresh-instance, Class-G, evidence-fold, projection, compatibility, and consumer direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-064221.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `326991dc049b7294951efb385298d9ef1a4a800fd6318a97c31478acc8066911`; parent relay hash recomputed: `9e197a394831593b2a5d2b2d3d7cffe8e93746c67131c2e34222f7e07d13786f`.
- Re-read the operative `050307` narrowing and latest VP `051057`; H-16 return 2 and master/VP review remain open.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: canonical table hydration and redo-derived commit order, recovery processor ordering, loop/hook construction, both callerless mint paths, `MintOrReplace`, and `realized_mint_ref`.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine` = PASS.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-064856.md`; root-wide invocation also reports unrelated pre-existing INDEX/lineage errors.
Next requested action: m-7.planner folds R5-F1..F3, routes exact m-1/m-2 confirmations, and returns fresh design/decision-record bytes for a new uniquely-parented pair review; H-16 IMPL remains held through the master/VP pass.
