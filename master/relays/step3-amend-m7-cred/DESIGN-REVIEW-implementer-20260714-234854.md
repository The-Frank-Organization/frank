## DESIGN-REVIEW - credential contract r3 closes the mechanics but leaves three transition/binding contradictions

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m7-cred-review-r3
PARENT_DISPATCH_ID: step3-amend-m7-cred
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the operator provisioning and first-E3 gates remain routed and are not decided by this review
GRILL_REQUIRED: yes
DESIGN_DOC_ID: step3-amend-m7-cred-credential-contract
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: step3-amend-m7-cred-grill-r3
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-amend-m7-cred/DESIGN-planner-20260714-233829.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-8.planner, m-3.planner
SUBJECT: r3 must revise - F8 is closed and the F7/F9/F10 directions are correct, but catalog-v2 activation contradicts the byte-exact drift oracle, Selected binds only an endpoint rather than the authorized freeze, and schema-vs-composition timing remains internally inconsistent

DESIGN_REVIEW_VERDICT: must-revise

R3 makes the right corrections: a conditional secret directory does not belong in the exact-live-equality family; component-wise no-follow/nonblocking descriptor access is now complete; a free credential-ID attach API is removed; auth-header bytes are pinned; and engine-v5 gets a closed structural contract. Three contradictions in the new text still prevent technical approval.

## Findings

### F11 - Catalog s8-v2 is neither schema-complete nor "never-red" under the locked source/runtime drift law

Section 2.5 introduces an optional `forbidden_only` list without naming its exact location in the catalog object and then says its v1 content is exactly the secrets row (`2026-07-15-step3-credential-contract.md:61-65`). A v2 candidate with the field absent would therefore satisfy the stated optional schema while omitting the only new guarantee. More importantly, the inherited drift law requires the one source artifact's bytes to equal the live runtime catalog bytes exactly and makes any divergence red (`2026-07-11-s8-config-host.md:109-114,174`). A reader-first deployment containing s8-v2 source bytes against an existing s8-v1 runtime store is red before the catalog record; retaining v1 source bytes until after the record is red afterward. Independent catalog/engine records "in either order" do not solve this source-artifact/version interval.

Required fold: pin the full field path (recommended `discovery.forbidden_only`) and make it REQUIRED for catalog s8-v2, with the exact closed singleton row `{id:"secrets", forbidden:"/secrets/"}` plus uniqueness/order rules. Then reconcile reader-first version adoption with the drift oracle. One coherent option is versioned immutable source artifacts (`catalog.s8-v1.json`, `catalog.s8-v2.json`) and a drift law selecting the artifact whose marker equals the loaded runtime member; new genesis selects the latest supported artifact, while an existing v1 store remains green until its accepted v2 record. If retaining a single mutable source artifact, define an equally exact migration transaction and the bounded state in which byte equality is evaluated. Extend FX-CRED-10/FX-CFG-5 with reader-upgrade-before-record, v1 runtime, accepted v2 transition, restart, and rollback/skip legs; every state claimed lawful must satisfy the named drift oracle.

### F12 - `Selected` claims full-freeze binding but carries only `frozenEndpoint`

The final authorization covers method + canonical URL + headers-sans-auth + body (`design:88`), but `Derive(lane, frozenEndpoint)` and `Selected` carry only the endpoint (`design:103,111`). Two requests with the same endpoint but different method, headers, or body therefore share the stated binding input. `Attach(req, Selected)` also receives a mutable request object; nothing specified prevents a same-endpoint mutation between authorization and attachment. The assertion that attach refuses "any bytes other than the freeze" is stronger than the capability actually described.

Required fold: define one immutable `FrozenRequest`/freeze-identity object covering all authorization bytes, not merely the endpoint. `Derive` consumes that identity and returns `Selected` bound to it; the m-3 authorization token and `Attach` must refer to the same identity. Attach may produce the one auth-header-bearing wire object, but must not accept an independently mutable request whose method/body/headers can diverge. The exact m-8 envelope representation may remain a B14 input; m-7 must pin the equality/capability relation now. Add same-endpoint/different-method, body, and non-auth-header substitutions plus mutate-after-authorize negatives.

### F13 - The v5 table conflates schema descriptors with semantic composition and leaves candidate acceptance timing unclear

The locked descriptor contract validates node kinds, closed key sets, and array element shapes, explicitly with **no value constraints beyond kind** (`2026-07-11-s8-config-host.md:21-28`). R3 instead says every table rule, including sorting, uniqueness, enum membership, non-empty sets, and generation monotonicity, is enforced at both "the descriptor layer" and composition (`credential-contract.md:39-55`). Current acceptance calls `validateEngineSchema(candidate)` and commits when it succeeds (`frank/internal/config/config.go:263-284`; `frank/internal/engine/submit.go:477-488`); a composition rule run only by later `config.Load` could accept a record that bricks the next start. The table also conflicts on endpoint port presence: section 4 says port is explicit while the v5 row marks it optional, and `lane_ids` constrains only array cardinality/uniqueness, not each string's local ID grammar.

Required fold: split the table into (A) schema(V5): exact keys, node kinds, array element shapes, and required/optional fields; and (B) semantic composition: local value, sorting, uniqueness, endpoint, and monotonic-transition rules. State that both validators run on load **and** on the transition candidate before pivot (or define one candidate validator that invokes both), so no accepted candidate is unloadable for an m-7-local rule. Resolve port as either required or absent-means-443, and pin each `lane_ids` element's local syntax/non-empty rule while leaving only catalog existence provisional to B14. Add a fixture that each semantic-invalid candidate is rejected pre-pivot and the prior config remains bootable.

## Confirmed

- R2/F7's direction is accepted: `secrets` belongs in a forbidden-only leak corpus, not the twelve-row live-root census. F11 concerns its exact schema and transition law only.
- R2/F8 is fully closed by root/parent/leaf descriptor confinement, nonblocking leaf open, dot-ID refusal, descriptor-bounded read, and the expanded fixture set.
- R2/F9's auth profiles, reserved-header refusal, Host/authority check, and opaque selection direction are accepted. F12 concerns full-freeze identity only.
- R2/F10's v5 structural delta and local rule set are accepted. F13 concerns validator ownership/timing and two residual grammar contradictions only.
- All r1 F1-F6 confirmations remain intact. The B14 consumer packet still blocks final approval/close, not this technical re-review. F11-F13 require no operator decision.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no credential provisioning/use, no E3 call, no lock, no merge, and no operator decision inferred.

ACTIONS_GIT_REF: wrote this r3 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`
Next requested action: m-7.planner folds F11-F13 into credential-contract r4 and returns it for technical re-review; final amendment approval remains held for the B14 consumer packet.
