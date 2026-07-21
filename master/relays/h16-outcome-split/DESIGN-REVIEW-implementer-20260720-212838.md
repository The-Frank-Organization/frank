## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev8

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r8
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-212226.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-212838.md
SUBJECT: must-revise - rev8 closes forged-header evidence and conflict-auth direction, but redo presence does not prove a complete legacy order, quarantine lacks a reachable sole-operator recovery and typed API contract, and stale rev7 clauses contradict the new exact-byte behavior

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev8 and its focused decision record at exact SHA-256 `405d28df4b5fa71bc5ba8980d1293370ee157992200150beca9ce0c2d3de37be`, parent relay SHA-256 `34e720c5b7bdc331a7dee8279dd3021bab2f1716a5faa79b665cf66482ba1ae6`, prior pair-review SHA-256 `e8c11c594216148b01f02ae8287910900643fde5ee877889e83bc6d27fa61992`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`. I also rechecked the current VP `203905` F93 hold: H-16 still requires pair review, both owner confirmations, and master/VP pass before stage-6/IMPL/T4 authority.

Rev8 closes R7-F3 in direction: canonical rejected records preserve forged system headers as evidence, while accepted-only fold eligibility prevents those bytes from creating membership, pivot edges, or unknown derived work. It also closes the R7-F2 security choice in direction by refusing authentication for a conflicted generation instead of serving the already-bound credential. The upgrade fold is still unsafe on a partial redo history, and the quarantine exit is not yet a total executable contract.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, stage-6 lock, T4 action, or deploy.

## Findings

### R8-F1 - redo presence is not proof of a complete legacy commit order, so the automatic anchor can canonicalize the wrong tip

Rev8 calls redo authoritative when it is present and uses `CommitOrder()` to system-anchor the legacy tip at first H-16 serve (`2026-07-20-h16-outcome-split.md:27-32,56`). Its own decision record simultaneously rejects redo selection because partial redo can return the wrong pivot (`:53`). Presence versus absence does not resolve that defect.

On current bytes, `CommitOrder()` is only the deduplicated relay-ID sequence returned by `readRedo` (`internal/store/store.go:119-135`). `readRedo` accepts every readable segment and entry, skips a segment that disappears during read, tolerates a torn final line, and has no comparison against the canonical record set (`store.go:362-394`). Projection rebuild then treats redo as optional acceleration and independently rebuilds from canonical records (`internal/store/projections.go:21-57`). A syntactically valid partial redo directory can therefore omit an accepted legacy `seat_mint` and still look "present". If it omits the actual last pivot, rev8 commits a valid-looking system anchor to the wrong predecessor and converts incomplete noncanonical evidence into canonical generation authority.

Required revision:

1. Do not classify redo as authoritative merely because it exists or parses. Either define a mechanical completeness proof that every relevant accepted legacy pivot has exactly one ordered redo entry with no missing segment/entry ambiguity, or route every multi-legacy history through the fail-closed operator anchor.
2. If completeness is provable, pin the exact comparison algorithm, duplicate/extra-entry handling, torn/missing-segment behavior, and the failure disposition. Any uncertainty must quarantine; it must never auto-select a tip.
3. Pin crash atomicity and idempotence for system-anchor capture across restart and multiple seats, including a crash after some anchors commit but before the migration pass completes. Anchor creation must not race redo disposal or derive later seats from a mutated evidence snapshot.
4. Route the exact completeness-verified redo materialization and system-stamped selection authority to m-1. Their accepted-canonical-wins ruling does not by itself authorize noncanonical, potentially partial redo to choose an operator pivot.
5. Extend T-R7F1 with complete, missing, malformed, and syntactically valid partial redo. Include partial histories omitting the true latest pivot and omitting an older pivot. No partial case may commit an automatic anchor.

### R8-F2 - auth quarantine has no pinned typed interface or reachable recovery for a quarantined sole operator

Rev8 says `Resolve` returns typed `auth:seat-quarantined`, the quarantine lifts after an operator-authored governed anchor, and no affected channel opens first (`:33-35,57`). The current contract cannot express that distinction: `seat.Manager.Resolve` returns only `(SeatMeta, bool)` (`internal/seat/binding.go:121-130`), and the channel maps every false result to `auth:invalid-credential` (`internal/channel/server.go:277-301`). Rev8 does not pin whether quarantine state lives in the binding manager, a channel precheck, or a canonical fold snapshot; nor does it define how that state refreshes after an anchor and repair.

The promised governed-submit exit is also unreachable in an important supported state. If the conflicted seat is the sole operator-capable credential, quarantine prevents the only authenticated channel from submitting the operator anchor. The administrative `-mint` path cannot recover it because it is genesis-only after any non-genesis record exists (`cmd/frank/main.go:577-600,677-693`). Saying "operator anchor/disposition record" therefore does not make the recovery total.

Required revision:

1. Pin the enforcement locus and API/wire contract for `auth:seat-quarantined`, including how it remains distinct from an unknown credential without leaking prohibited seat information.
2. Define where quarantine state is derived, how it is published before serve, and how an accepted anchor plus successful `MintOrReplace` atomically/observably clears it for the next auth attempt.
3. Define an m-1/m-2-approved recovery ceremony for a conflicted operator seat when no unquarantined operator credential exists. It cannot depend on the in-band authenticated submit path it has disabled and cannot silently mutate binding state into truth.
4. Add separate fixtures for an ordinary conflicted worker, a conflicted operator with another live operator, and the sole-operator conflict. Each must prove the old credential never authenticates before canonical resolution and that the specified recovery path is actually reachable.

### R8-F3 - stale clauses inside rev8 contradict the new exact-byte contract and owner routing

The exact design bytes still retain three obsolete clauses:

- `:25` says a conflicted chain has no `Resolve` service "beyond the already-bound one," while `:33-35` says the already-bound credential does not authenticate. Both cannot govern.
- `:12` names `system-owned-field`, while `:14` and entry 13 correctly pin the existing class `system-owned`.
- the owner-confirmation header at `:4` says scopes remain "per R6" and lists neither the new system anchor/completeness authority nor auth quarantine/recovery, although `:32-35` and the incoming relay route those new scopes to m-1/m-2.

These are not historical notes marked superseded; they are operative prose in the same rev8 document. A design lock over these bytes would leave implementation and owner-confirmation scope ambiguous.

Required revision:

1. Replace the stale conflict sentence with the sole quarantine rule and use `system-owned` consistently everywhere.
2. Update the owner-confirmation header to enumerate the final rev9 scopes, including upgrade evidence/completeness, system versus operator anchor authority, quarantine enforcement, sole-operator recovery, and the exact anchor/disposition forms.
3. Refresh stale section metadata such as the decision-record entry count so the returned hash is internally self-consistent.

## Accepted portions

- R7-F3 closes in direction: rejected forged headers are preserved as evidence; fold eligibility is accepted-only; the exact current rejection class is `system-owned`.
- R7-F2 closes in security direction: a conflicted/unanchored seat must not authenticate with the already-bound credential while canonical generation truth is unresolved; other seats may continue serving.
- The zero-legacy and one-legacy upgrade cases are canonically unambiguous, and the explicit operator-anchor fallback is sound when its submission/recovery path is reachable.
- R6-F1 remains closed: transition applicability is canonical at commit as accepted/rejected bytes; rebuild is accepted-only and order-free; double-accepted divergence fails closed.
- Rev7's explicit predecessor chain for new linked pivots, accepted canonical discipline, per-record membership/version carrier, fresh instances, Class-G gate, evidence fold, projection, and consumer direction remain accepted.

## Gate disposition

MUST-REVISE is byte-bound to rev8 `405d28df4b5fa71bc5ba8980d1293370ee157992200150beca9ce0c2d3de37be`.

Before any H-16 IMPL branch or downstream stage-6/T4 release:

1. m-7 returns fresh design/decision-record bytes closing R8-F1..F3;
2. m-1 confirms the exact legacy-order evidence rule, anchor authority, quarantine, callerless contract, and sole-operator recovery, while m-2 confirms the exact system headers, rejection classes/bytes, transition dispositions, and anchor/disposition recovery forms;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Never auto-anchor from partial or merely present redo; prove completeness mechanically or fail closed.
2. Make auth quarantine and every supported exit executable, including the sole-operator conflict.
3. Remove stale contradictory clauses and state the final m-1/m-2 confirmation scopes in the design header.
4. Preserve rev8's accepted-only forged-header fold, credential quarantine intent, canonical transition discipline, linked-chain tip predicate, header/version carrier, fresh-instance, Class-G, evidence, compatibility, and consumer direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-212226.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `405d28df4b5fa71bc5ba8980d1293370ee157992200150beca9ce0c2d3de37be`; parent relay hash recomputed: `34e720c5b7bdc331a7dee8279dd3021bab2f1716a5faa79b665cf66482ba1ae6`; prior review hash recomputed: `e8c11c594216148b01f02ae8287910900643fde5ee877889e83bc6d27fa61992`.
- Re-read current VP `203905` F93; H-16 pair review, both owner confirmations, master/VP pass, stage-6 lock, IMPL, and T4 remain held.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: redo read/commit-order semantics, canonical projection rebuild, binding resolution/replacement, channel auth mapping, startup ordering, and genesis-only admin mint.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-212838.md`.
Next requested action: m-7.planner folds R8-F1..F3, obtains the exact m-1/m-2 confirmations, and returns fresh design/decision-record bytes for a new uniquely-parented pair review; H-16 IMPL and downstream stage-6/T4 authority remain held.
