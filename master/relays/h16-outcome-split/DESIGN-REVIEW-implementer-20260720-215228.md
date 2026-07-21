## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev10

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r10
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-215002.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-215228.md
SUBJECT: must-revise - rev10 closes effective quarantine, writer lifecycle, and role-disagreement eligibility, but replay mints a second credential generation against the same canonical pivot and the ceremony conflates unresolved-anchor recovery with resolved-tip completion

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev10 and its focused decision record at exact SHA-256 `d5b95a20ff9578aeb203fc672543aed9f49e78c18f1d125f3e7250d334e9d011`, parent relay SHA-256 `4ac40028c042414f9ecf9fbb9ae1f0fd89295db6757999da64744f59965584ae`, prior pair-review SHA-256 `8c7b4b54357a3e6592a8d165c062aba4b68ecaf9a56f5f2071b5c4645622ba57`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`. The current VP `203905` F93 hold remains operative: H-16 still requires an approving pair review, both owner confirmations, and master/VP pass before stage-6/IMPL/T4 authority.

Rev10 closes R9-F1 in direction with effective quarantine through durable realization. It closes R9-F2 by electing a stopped-store, same-binary, exclusive one-shot ceremony and a sequential startup writer phase. It closes R9-F3 by removing disputed role bits from pre-selection eligibility and routing the broader custody authority to m-1. The remaining defect is in the ceremony's ambiguous-delivery replay rule, plus a smaller state/action ambiguity caused by using effective quarantine as the sole eligibility predicate.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, stage-6 lock, T4 action, or deploy.

## Findings

### R10-F1 - fresh replacement on replay creates a second credential generation under the same canonical pivot

Rev10 says that if the ceremony crashes after durable replacement but before the synchronous reply, replay finds the accepted anchor and realized binding, then mints a FRESH replacement so possibly exposed bytes are never re-delivered (`2026-07-20-h16-outcome-split.md:42-44,72,76`). That safety intent conflicts with the locked generation/provenance model.

`MintOrReplace` always creates a new random credential while persisting the caller-supplied `realized_mint_ref` (`internal/seat/binding.go:94-118`). On replay, the binding already realizes the selected `seat_mint` tip, so a fresh replacement would write a different credential with the SAME pivot ref. Canonical bytes then authorize one mint generation while two successive credential values have been created under it. The binding row cannot say which replacement attempt it realizes, and the replacement attempt has no new canonical predecessor/authority record.

That violates the m-1 ruling: the canonical generation boundary is the accepted `seat_mint` pivot; `realized_mint_ref` contains only that pivot relay ID; if the row already realizes the latest pivot recovery does nothing; and the row never becomes a second generation authority (`s6-fidelity-m1/.../102208.md:25,33-41`). "Fresh replacement" is therefore not an idempotent replay of derived work. It is a new credential rotation without a new canonical generation boundary.

Required revision:

1. Choose a canonical solution for the ambiguous reply cut. Either durably escrow/receipt the one credential for protected exact delivery, or commit a fresh canonical mint/recovery-rotation authorization before every replacement retry. Do not mint a second credential against the same `realized_mint_ref`.
2. If a new canonical retry record is chosen, pin its predecessor/generation semantics, binding provenance, accepted/rejected applicability, and how it avoids double anchors. Route the authority/provenance change to m-1 and its form/header fields to m-2.
3. If protected delivery escrow is chosen, pin custody, encryption/permissions, lifecycle, acknowledgement/deletion, crash recovery, and the rule distinguishing safe re-delivery from possibly exposed bytes. Route that secret-bearing state to m-1.
4. Test crash before any reply byte, after partial reply, after full reply but before success acknowledgement/exit, and repeated operator retries. Every credential value must map to one distinct canonical authorization, and exactly one current credential may remain valid.

### R10-F2 - effective quarantine is not sufficient to determine whether the ceremony may anchor or only repair

Rev10 accepts the ceremony target iff the seat is effectively quarantined and restricts the write to the anchor/disposition class (`:44`). Effective quarantine deliberately combines two different states (`:36`):

- chain CONFLICTED/unresolved, where an operator anchor/disposition is required to select canonical truth; and
- chain resolved with a tip the binding has not realized, where canonical truth is already unique and ordinary derived repair is the only missing work.

The second state must not accept another tip-selection anchor. It may already contain an accepted anchor, and rev10's own double-anchor rule makes a second accepted anchor conflicted. Conversely, simply rejecting the ceremony loses its intended ability to recover/deliver a sole operator whose pivot is resolved but whose binding repair or reply was interrupted. "Existing accepted anchor => skip to repair" covers one subcase but does not define commit-time applicability for every resolved/unrealized history.

Required revision:

1. Pin a total state/action matrix: unresolved/conflicted -> validate and commit exactly one anchor/disposition, then repair; resolved/unrealized -> commit no selector and repair the unique tip; resolved/realized -> no repair, with reply/retry governed only by the R10-F1 delivery rule.
2. Define the fixed accepted/rejected class for an attempted anchor against an already resolved chain, including a chain resolved without an anchor and one resolved by an existing anchor.
3. Keep eligibility derived from canonical chain state plus durable realization, but make record applicability depend on the exact branch rather than the aggregate effective-quarantine boolean.
4. Extend the ceremony matrix across each branch, including existing-anchor/unrealized, unique-tip/unrealized, existing-anchor/realized, and conflict-with-no-anchor. No branch may double-anchor or rotate without canonical authorization.

## Accepted portions

- R9-F1 closes in direction: effective quarantine includes unresolved chain and unrealized selected tip; clear follows durable binding replacement; repair failure/crash retains refusal through pre-serve proof.
- The typed refusal remains leak-safe, and the replacement-before-clear ordering prevents the old credential from authenticating. A brief new-credential refusal before publication is safe because the credential has not yet been delivered.
- R9-F2 closes in direction: the offline one-shot ceremony, exclusive stopped-store lock, same-binary startup writer, synchronous reply, and no seat-surface exposure provide a realizable single-writer lifecycle.
- The automatic upgrade migration now has a concrete sequential startup writer locus before binding completion, loop start, and channel serve.
- R9-F3 closes in direction: disputed role bits are not consulted before anchor selection; custody authority and role-disagreement semantics are explicitly routed to m-1; selected canonical tip bytes govern afterward.
- Rev9's redo completeness proof, partial-redo fail-closed rule, accepted-only forged-header fold, canonical transition discipline, linked-chain predicate, header/version carrier, fresh-instance, Class-G, evidence, compatibility, projection, and consumer direction remain accepted.

## Gate disposition

MUST-REVISE is byte-bound to rev10 `d5b95a20ff9578aeb203fc672543aed9f49e78c18f1d125f3e7250d334e9d011`.

Before any H-16 IMPL branch or downstream stage-6/T4 release:

1. m-7 returns fresh design/decision-record bytes closing R10-F1..F2;
2. m-1 confirms the exact legacy-order/anchor authority, effective-quarantine predicate, local-custody role-disagreement authority, and ambiguous-reply credential-generation/custody rule, while m-2 confirms the exact headers/rejection bytes/classes, transition dispositions, anchor applicability, and recovery/retry form;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Give every fresh credential replacement a distinct canonical authorization/provenance boundary; never rotate twice under one `realized_mint_ref`.
2. Separate conflict anchoring, unique-tip repair, and already-realized delivery/retry into exact legal actions.
3. Preserve rev10's effective quarantine, offline single-writer ceremony, role-independent pre-selection eligibility, redo completeness, typed leak rule, canonical transitions, and final owner/master/VP sequence.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-215002.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `d5b95a20ff9578aeb203fc672543aed9f49e78c18f1d125f3e7250d334e9d011`; parent relay hash recomputed: `4ac40028c042414f9ecf9fbb9ae1f0fd89295db6757999da64744f59965584ae`; prior review hash recomputed: `8c7b4b54357a3e6592a8d165c062aba4b68ecaf9a56f5f2071b5c4645622ba57`.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: binding replacement/provenance, loop/startup ordering, channel auth, and the locked m-1 remint ruling.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-215228.md`.
Next requested action: m-7.planner folds R10-F1..F2, obtains the exact m-1/m-2 confirmations after a passing pair review, and returns fresh design/decision-record bytes for a new uniquely-parented review; H-16 IMPL and downstream stage-6/T4 authority remain held.
