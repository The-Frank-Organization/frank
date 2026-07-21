## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev11

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r11
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-215950.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-221144.md
SUBJECT: must-revise - rev11 gives every retry a canonical rotation pivot and closes the ceremony matrix, but the delivery-only pivot does not preserve role/is_operator and its admin provenance predicate has no exact non-forgeable canonical carrier or m-2 registry contract

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev11 and its focused decision record at exact SHA-256 `53159a4cf788e799015d036f3b2966f0cd24f6152205e68e074df2707676105c`, parent relay SHA-256 `45a873a3c5c11ac1a810bd546cda2cba896b7349b1ca427527837dd89b2f2c4d`, prior pair-review SHA-256 `0edc834f18cafc4856903afce85a9a38f24d72b23973ea7141d680326c1e657b`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`. The current VP `203905` F93 hold remains operative: H-16 still requires an approving pair review, both owner confirmations, and master/VP pass before stage-6/IMPL/T4 authority.

Rev11 closes R10-F1 in direction by committing a fresh canonical rotation pivot before each replacement retry, so every credential maps to a distinct generation boundary. It closes R10-F2 with a total conflict/repair/retry matrix and fixed `anchor-target-resolved` rejection. The remaining defects are both in the exact retry-pivot contract: its authority-bearing body fields and the provenance marker that gates access to this offline rotation path.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, stage-6 lock, T4 action, or deploy.

## Findings

### R11-F1 - a delivery-retry pivot can silently change role or operator authority

Rev11 says a delivery retry commits an ordinary admin-marked rotation `seat_mint` pivot, with `mint_predecessor` naming the current tip, then realizes that new pivot (`2026-07-20-h16-outcome-split.md:49-50,79`). An ordinary `seat_mint` body contains three required authority-bearing values: `seat`, `role`, and `is_operator` (`internal/engine/submit.go:391-421`). Rev11 pins only the predecessor. It does not say where the retry pivot obtains `role` and `is_operator` or whether the ceremony form can supply them.

This operation is scoped as ambiguous credential delivery, not an authority edit. If the operator can choose those fields, the offline retry path can promote/demote or change the role of any ceremony-eligible seat while being described and reviewed only as credential replacement. If fields are omitted, the ordinary parser rejects the pivot. The selected current tip is canonical and unambiguous in matrix branch (3), so there is no reason to accept new authority values.

Required revision:

1. Pin the retry pivot body as system-derived: `seat`, `role`, and `is_operator` are copied exactly from the current canonical tip; only generation/provenance fields change. The ceremony form must not accept authority-field overrides.
2. If the form necessarily carries those fields, require exact equality with the current tip at commit and define a fixed rejection class for any mismatch.
3. Add the no-authority-delta rule to m-1 scope (g) and the exact form/validation behavior to m-2 scope.
4. Test worker, planner, and operator-role tips, plus attempted role and `is_operator` changes. Every retry must preserve authority bytes while changing only the canonical generation pivot and credential.

### R11-F2 - the canonical `admin-marked` retry predicate has no exact carrier or forgery rule

Rev11 scopes offline retry using canonical provenance: the current tip is an "admin-marked rotation pivot," the chain has an "admin-marked anchor," or the terminal disposition is `realized-undelivered` (`:49`). It also says the rotation records a retry reason. Neither the marker nor reason has a named field, type/version, owner, stamp locus, or validation behavior.

That omission is security-relevant because the marker grants access to repeat offline rotations after realization. If it is supplied in the ordinary `seat_mint` body or an unowned header, an in-band submitter can forge ceremony provenance and make a healthy seat satisfy the recovery predicate. If it is system-stamped, it is an additional system-owned canonical member whose registry and forged-input behavior must be explicit. The current m-2 scope still calls `hook_contract` and `mint_predecessor` the TWO system-stamped header rows and mentions only the anchor schema's admin marking (`:4-6`); it does not cover the retry pivot marker/reason used by this new predicate.

Required revision:

1. Name the exact canonical carrier for ceremony provenance and retry reason, including owner, type/version/enum, accepted-only semantics, and whether anchor and rotation reuse one marker or use separate members.
2. Pin the writer/stamp locus and normal-ingress rule. Ordinary native/MCP seat submissions must not be able to create ceremony provenance; forged values must commit the exact rejected shape/class already governed by m-2.
3. Add every new system-owned member and its executable rejection path to m-2's final scope. If existing envelope provenance is sufficient, state the exact predicate instead of introducing an unnamed mark.
4. Add forged-marker legs through both ingress surfaces and prove an ordinary accepted `seat_mint` cannot satisfy the retry predicate, while an offline-ceremony pivot can.

## Accepted portions

- R10-F1 closes in direction: each retry first commits a distinct predecessor-linked canonical rotation pivot, then realizes that pivot; no fresh credential shares the predecessor's `realized_mint_ref`.
- R10-F2 closes: the matrix separates conflict anchoring, unique-tip repair, and realized-tip retry; an anchor against either resolved shape commits rejected `anchor-target-resolved`.
- The retry provenance cases and operator-asserted ambiguity are a viable scope shape once the canonical carrier is exact and non-forgeable.
- Rev10's effective quarantine through durable realization, typed leak rule, offline exclusive one-writer ceremony, startup writer ordering, role-independent conflict eligibility, and role-disagreement recovery remain accepted.
- Rev9's redo completeness proof, partial-redo fail-closed rule, accepted-only forged-header fold, canonical transition discipline, linked-chain predicate, header/version carrier, fresh-instance, Class-G, evidence, compatibility, projection, and consumer direction remain accepted.

## Gate disposition

MUST-REVISE is byte-bound to rev11 `53159a4cf788e799015d036f3b2966f0cd24f6152205e68e074df2707676105c`.

Before any H-16 IMPL branch or downstream stage-6/T4 release:

1. m-7 returns fresh design/decision-record bytes closing R11-F1..F2;
2. m-1 confirms the exact legacy-order/anchor authority, effective quarantine, local-custody authority, retry generation, and no-authority-delta rule, while m-2 confirms every system-owned carrier/rejection path, transition class, anchor applicability, and recovery/retry form;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Preserve the current canonical tip's role and `is_operator` exactly on every delivery-retry pivot.
2. Make ceremony provenance and retry reason exact, canonical, accepted-only, and non-forgeable through ordinary ingress.
3. Preserve rev11's distinct retry pivot, total ceremony matrix, effective quarantine, offline single-writer lifecycle, redo completeness, canonical transitions, and final owner/master/VP sequence.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-215950.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `53159a4cf788e799015d036f3b2966f0cd24f6152205e68e074df2707676105c`; parent relay hash recomputed: `45a873a3c5c11ac1a810bd546cda2cba896b7349b1ca427527837dd89b2f2c4d`; prior review hash recomputed: `0edc834f18cafc4856903afce85a9a38f24d72b23973ea7141d680326c1e657b`.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: `seat_mint` body/parser authority fields, binding replacement/provenance, normal ingress validation, and the locked m-1 remint ruling.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-221144.md`.
Next requested action: m-7.planner folds R11-F1..F2, obtains the exact m-1/m-2 confirmations after a passing pair review, and returns fresh design/decision-record bytes for a new uniquely-parented review; H-16 IMPL and downstream stage-6/T4 authority remain held.
