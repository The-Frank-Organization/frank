## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev6

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r6
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-065418.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-070914.md
SUBJECT: must-revise - per-record membership closes the epoch and activation-window defects, but first-wins transitions and latest-mint selection still depend on non-canonical commit order, while the new system-owned envelope field has no realizable inbound-rejection or unknown-version contract

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev6 and its updated focused decision record at exact SHA-256 `a8710ee511277adb02e1fe1aa484810d9453bea6cea07d8f9cf584fc691c22e3`, parent relay SHA-256 `41710d1ceb0b1797cf4d476d1abc4a9b8c83a02ceb9c2cc19d1e438ab636eda0`, current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`, the operative `050307` narrowing, and the locked m-1 re-mint ruling at `frank/.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-102208.md`.

Rev6 closes the specific rev5 epoch defect: per-record canonical membership needs no redo-derived relation, and stamping both live-loop and recovery decisions removes the activation window. It also retracts the false delivery-inert claim and moves unresolved-marker consultation into each callerless path's real startup position. Three remaining contracts still cannot reconstruct or enforce the promised state from current canonical bytes.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, or deploy.

## Findings

### R6-F1 - the transition fold still uses first-wins commit order that canonical records do not contain

Rev6 keeps rev5's transition map unchanged (`2026-07-20-h16-outcome-split.md:31`). That map says the first operator resolution naming an attempt/park instance wins and later duplicate, conflicting, or stale records are fold-inert anomalies (`DESIGN-planner-20260720-064221.md:26`; rev5 design §5).

The new `hook_contract` field only classifies source-record membership. It does not make competing transition records ordered. After redo deletion:

- `Store.CommitOrder()` has no entries because it derives solely from redo (`internal/store/store.go:119-135`);
- `tables.Build` therefore receives no canonical commit order (`internal/tables/tables.go:116-137`);
- two conflicting operator resolutions naming the same marker remain indistinguishable accepted canonical records except for random relay IDs.

The commit-time one-shot check does not by itself solve rebuild. Rev5 says a later conflict "commits as fold-inert/anomaly", but no canonical byte is pinned that records that disposition. A fresh fold can choose the lexically earlier random relay ID and reverse which operator disposition won.

Required revision:

1. Make transition validity order-independent in canonical bytes. For example, commit an invalid later transition as a canonically rejected decision with a fixed `failing_edge`, or add an exact system-stamped applicability disposition; alternatively define an order-independent conflict rule that fails closed. Do not leave "first wins" dependent on lost commit order.
2. Pin the exact commit-time and rebuild predicates for valid, duplicate, conflicting, stale, and nonexistent-target resolutions, including the bytes that distinguish applied from fold-inert.
3. Extend T2/T13/T14/T16 with conflicting resolutions whose relay-ID lexical order opposes commit order, delete redo/projections, rebuild repeatedly from shuffled record slices, and require the same fold/anomaly result.
4. Include these transition disposition bytes in the exact m-2 confirmation if they touch the operator form or system-owned field catalog.

### R6-F2 - callerless mint recovery still cannot identify the latest pivot without redo

Rev6 says each callerless path performs an order-independent canonical-record scan before acting (`2026-07-20-h16-outcome-split.md:23-29`). That closes unresolved-marker consultation only if the target pivot is already known. The general repair path still selects that target through `latestSeatMintPivots`, which iterates `Store.CommitOrder()` and overwrites the per-seat entry as it walks redo order (`cmd/frank/main.go:629-674`).

With redo deleted, the helper returns no pivots. With multiple accepted re-mints and incomplete redo, it can select the wrong pivot. Canonical record bytes carry no per-seat predecessor/sequence relation. This conflicts directly with the locked m-1 redlines:

- recovery must compute the seat's latest accepted `seat_mint` pivot in commit order;
- accepted canonical records win;
- repair must finish before channels open so the superseded credential fails the first post-restart `Resolve`
  (`frank/.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-102208.md:37-47`).

The rev6 T2/T14 redo-deletion cut therefore strands the ruled s6 recovery predicate or leaves an older delivered credential live, even though the H-16 membership fold itself reconstructs.

Required revision:

1. Give accepted `seat_mint` pivots an order-independent canonical per-seat generation relation, such as an exact predecessor pivot reference validated at commit, or design another canonical latest-pivot mechanism. Relay-ID lexical order and redo are not valid substitutes.
2. Route the exact relation and its conflict behavior to m-1 and m-2 because it changes the locked generation/re-mint contract and the operator form.
3. Make both callerless paths identify the same canonical current pivot, then consult that pivot's unresolved-marker/evidence fold before `MintOrReplace`.
4. Add a two-rotation fixture with pivots A then B whose relay IDs sort B before A; delete all redo/projections, restart, and prove B alone is current, an unresolved marker on B blocks both paths, authorized repair realizes B, A's credential fails the first post-restart `Resolve`, and no older pivot is reminted.

### R6-F3 - `hook_contract` is not yet an enforceable system-owned envelope contract

Rev6 calls the m-2 delta exact: `hook_contract` is an `int`, `omitempty`, absent/0 legacy, 1 H-16, system-stamped, and rejected on the seat-form surface (`2026-07-20-h16-outcome-split.md:14-21`). Current validation cannot implement that statement without additional pinned behavior:

- decoded submit payloads contain the candidate `record.Envelope` before stamping (`internal/engine/submit.go:51-64`);
- `seat.Stamp` overwrites only `From` and `Role` (`internal/seat/binding.go:161-164`);
- FieldSpec `valueForSpec` reads only body or headers, not envelope members (`internal/fieldspec/validate.go:203-210`);
- the system-owned rejection helper applies only to header-layer fields (`validate.go:149-156`).

Thus an inbound `hook_contract` value is neither visible to the generic validator nor cleared by the existing identity stamp. Saying the committing writer stamps hooked records does not disposition a forged value on an unhooked record, and it does not pin whether a hooked value is rejected before being overwritten.

The version rule is also ambiguous: only value 1 is defined, but open uses `hook_contract >= 1`. An older binary would silently interpret an unknown future value as H-16 semantics rather than fail closed.

Required revision:

1. Pin the exact raw-candidate rejection locus before system stamping, including hooked and unhooked candidates, native and MCP ingress, and the typed failure class.
2. Pin writer behavior after revalidation: accepted hooked records receive exactly value 1; every other committed record receives zero/absent unless another explicitly owned contract says otherwise.
3. Decide whether this is a membership bit or a version. If a bit, use a closed representation. If a version, define unknown-positive handling; an old binary must fail closed rather than silently fold a future contract as v1.
4. Have m-2 confirm the executable envelope extraction/rejection path, not only the registry row, and test raw forged values `{0,1,2}` through both ingress surfaces plus restart/rollback decoding.

## Accepted portions

- R5-F1 closes in direction: `hook_contract` makes H-16 source membership canonical, per-record, and independent of redo or relay-ID order.
- R5-F2 closes in direction: stamping in both live and recovery commit paths removes the separate activation event and prevents a new-binary hooked decision from becoming legacy merely because startup fails.
- R5-F3's ordering and honesty corrections close: each callerless path now owns its pre-effect marker consultation, the later-drain claim is gone, and the delivered-A invalidation/admin-custody consequence is explicit.
- Fresh attempt identities, predecessor-scoped authority, Class-G retry-before-replay, `realized_mint_ref` evidence classification, projection, consumer table, dropped heal nudge, and master/VP-before-IMPL sequence remain accepted.

## Gate disposition

MUST-REVISE is byte-bound to rev6 `a8710ee511277adb02e1fe1aa484810d9453bea6cea07d8f9cf584fc691c22e3`.

Before any H-16 IMPL branch:

1. m-7 returns fresh design/decision-record bytes closing R6-F1..F3;
2. m-1 confirms the canonical latest-pivot plus callerless authority/delivery contract, and m-2 confirms the exact operator form, transition disposition, and envelope-field ingress/stamp contract;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Make transition application and anomaly classification deterministic with redo absent and adversarial record order.
2. Make the current seat-mint pivot canonical without redo, preserving every m-1 recovery redline.
3. Make `hook_contract` unsuppliable and version-safe through an exact executable m-2-owned ingress contract.
4. Preserve rev6's accepted per-record membership, activation-window closure, honest mint consequence, direct marker scan, fresh-instance, Class-G, evidence, compatibility, and consumer direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-065418.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `a8710ee511277adb02e1fe1aa484810d9453bea6cea07d8f9cf584fc691c22e3`; parent relay hash recomputed: `41710d1ceb0b1797cf4d476d1abc4a9b8c83a02ceb9c2cc19d1e438ab636eda0`.
- Re-read the operative `050307` narrowing and locked m-1 `102208` ruling; H-16 return 2, both owner confirmations, and master/VP review remain open.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: redo-derived commit order, table hydration, transition validation surface, submit decode and seat stamping, FieldSpec envelope handling, both callerless mint paths, latest-pivot selection, binding replacement, and `realized_mint_ref`.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat` = PASS.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-070914.md`; root-wide invocation also reports unrelated pre-existing INDEX/lineage errors.
Next requested action: m-7.planner folds R6-F1..F3, routes exact m-1/m-2 confirmations, and returns fresh design/decision-record bytes for a new uniquely-parented pair review; H-16 IMPL remains held through the master/VP pass.
