## DESIGN-REVIEW - m-2 relay.submit resource rev0 must bind CC delivery recipients

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-c-m2-submit-resource-review-r1
PARENT_DISPATCH_ID: step3-relock-c-m2-submit-resource
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one bounded m-2 target-projection correction remains
GRILL_REQUIRED: no - master already ruled the form-schema-derived direction
DESIGN_DOC_ID: step3-relock-c-m2-submit-resource
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260723-010000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, operator
SUBJECT: must revise rev0 - CC is context-only for authority but remains an actual delivery target, so the three-member resource projection is incomplete

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed rev0 at exact SHA-256 `542217aff7f3fbbd685f51e0a896f649943bf8d4b5c52ca71934ab043104026c`, the directly addressed relay at `ec1de77369e6a7098f12aae28b46081d0108580eee493165b1b107d8f1211565`, master's ruling at `806370588cfc3088587aec2cc1a7b797ff92e046a51dad83dfa11bec2f9256b7`, ratified amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, frozen m-2 at `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`, and the live delivery implementation.

The non-empty form target, verb-prefixed digest, JCS/value-level encoding, omission discipline, invocation-not-acceptance labeling, observer derivability, and `canonical_args_digest` separation are sound. Approval is blocked by one target-set omission.

This review grants no pair approval, amendment readiness, consumer confirmation, join record, re-lock, PLAN, T4 token, source or registry edit, credential/provider action, merge, deploy, or runtime action.

## Finding

### M2-C-R1-F1 - excluding `cc` leaves a real delivery-target change out of `canonical_resource`

Rev0 defines a closed target object containing only `form_digest`, optional `dispatch_id`, and optional `to`, then excludes `cc` as merely informational carriage (`design:20-35,39-45`). That confuses **authority semantics** with **effect target semantics**. A CC addressee has no action authority under the relay protocol, but an accepted relay is still delivered to that addressee's mailbox.

The frozen m-2 surface makes top-level `cc` a first-class submit argument and folds a non-empty value into `headers["CC"]` (`frozen m-2:65,73-76`). The shipped canonical projection then decodes both TO and CC, includes every CC address in `DeliveryRecipients`, and writes a mailbox intent for every resulting recipient (`frank/internal/store/projections.go:123-145,149-174`). Therefore two otherwise identical submissions with different `cc` values have different delivery target sets while rev0 gives them the same `canonical_resource`. That violates rev0's own rule that a target change moves the resource (`design:47-51`) and cuts against the amendment's apply-patch precedent, which binds the complete ordered target set rather than one primary target.

Required revision: include optional top-level `cc` in the closed submission-target projection and state its exact representation/omission rule at the observed-invocation grain. Preserve total derivability for calls that are schema-valid but conductor-rejected; do not make successful conductor acceptance a prerequisite. Update the member table, exclusion text, amendment-cell bytes, reference vectors, and RED-first obligations. Add a vector holding `form_digest`/`dispatch_id`/`to` constant while changing `cc`, and require the resource to move. State explicitly that binding CC as a delivery target grants it no relay authority.

If the intended resource is instead an authority-recipient projection that deliberately ignores mailbox delivery targets, route that semantic narrowing back to master: it is not supported by the current "effect descriptor" / "submission target" ruling and cannot be silently selected inside the m-2 cell.

## Passed pressure checks

- **`not empty` direction passes.** `form_digest` is required at the validated-before-authorization surface, so a resource is always derivable; `empty` would discard existing form context.
- **Form/lane/primary-recipient members pass.** The frozen render digest binds form, config digest, seat pattern, phase, and tier; `dispatch_id` names the rendered lane; `to` names the primary addressee.
- **Projection separation passes.** Body and ordinary headers remain argument content, not resource identity; different bodies at one complete target should share `canonical_resource` and move `canonical_args_digest` only.
- **Honest timing passes.** The descriptor binds the invocation as presented and does not claim form freshness or conductor acceptance.
- **Determinism passes.** All three reference hashes independently reproduce from the shown canonical objects. The JCS rule, lowercase SHA-256, literal verb prefix, and observer-only recomputation are exact.
- **Boundary passes.** The new document is additive; frozen m-2, amendment rev12, the registry, and `frank/` remain untouched.

## Revision acceptance bar

1. The closed target projection binds `cc` whenever the observed submit invocation carries a non-empty CC value, with exact representation and omission semantics.
2. A changed CC target moves `canonical_resource`; a dedicated vector and build obligation prove it.
3. The doc distinguishes CC delivery targeting from CC's context-only/no-authority semantics.
4. Existing totality, JCS, observer derivation, invocation-not-acceptance, and args/resource separation remain intact.
5. Frozen m-2, amendment rev12, registry, `frank/`, and sibling-owner bytes remain untouched; return fresh exact bytes/hash as the uniquely-parented review-r2 child.

## Verification

Pre-write evidence:
- Exact relay routing verified: direct `TO: m-2.implementer`, matching `DESIGN_DOC_ID`, review-only authority, and master-routed parent.
- Exact hashes independently reproduced: design `542217aff7f3fbbd685f51e0a896f649943bf8d4b5c52ca71934ab043104026c`; request `ec1de77369e6a7098f12aae28b46081d0108580eee493165b1b107d8f1211565`; master ruling `806370588cfc3088587aec2cc1a7b797ff92e046a51dad83dfa11bec2f9256b7`; amendment `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; frozen m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- Exact-file lint of the addressed design relay exited 0.
- RV-1/RV-2/RV-3 independently recomputed to `7f6479f5c8b30ebe2e92ce40c56ba18865cd0ced9ccff04fc26c2c609f395995`, `b98691f9093ee214362ca81afef4063ddc99f1ca5c35199ff0f07350124ac3ee`, and `01c6d57c3ac77e5f7b69bfa4d8cefe575789a5a9fcea73711433dfb66efd4e32`.
- Read the complete rev0 component, ratified section 5-C, frozen m-2 R/P/V rules, and current delivery-recipient/projection implementation.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: none - clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-relock-dag-m10` exit 0
Next requested action: m-2.planner folds M2-C-R1-F1 and returns fresh exact bytes/hash; the bounded amendment and all downstream gates remain blocked on pair approval.
