## DESIGN-REVIEW - lane-2 r4 MUST REVISE: normalized classifier tokens have no exact projection, the epoch no-row branch is unrepresentable, and the m-9 residue was not routed to Master

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r5
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded pair corrections plus producer work routed through Master
GRILL_REQUIRED: no - no new product-semantic choice is required from this pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-121500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r4 589964063f3fa21a4985d6c1cc7dfbec23b714b43aa9a2af35943afd74d10918 must revise - classifier values are undocumented projections rather than producer literals, N3 cannot represent no provider_attempts row, and the m-9 residue has no addressed Master escalation

## Verdict

**MUST REVISE.** R4 closes the prior D4 hold, corrects N3's m-9 E0 state to `no_message`, removes `no_message` from a durable row column, keeps DAG status out of the runtime schema, and honestly returns the logical-surface binding to partial with no m-9 hash bound. The per-point comparison strategy and producer-root derivability remain sound.

Two local mechanics still prevent approval. The classifier enums are syntactically closed but are not exact producer literals and have no specified projection from the actual m-8 messages. The required `m10_row_digest: hex|absent` cannot represent the legitimate N3 branch where no `provider_attempts` row exists. In addition, the three outer-member recipes are described as escalated to Master, but no relay addresses Master; this DESIGN relay addresses only the Implementer and gives Master CC context.

## Findings

### M3-L2-R4-F1 - BLOCKER - the classifier is closed over sink-invented categories, not mechanically derived producer facts

Section 3.2a labels its values "producer literals," but several are normalized or invented sink values:

- m-8 r5 says only `m8.dataP_reply.v2` "typed reject" for rows 1/2a; it does not author a reply kind named `pre_freeze_typed_reject`.
- The actual epoch values are `STALE_EPOCH` / `EPOCH_AHEAD`, while the classifier uses lowercase underscore forms.
- CTRL-C values carry payloads such as `rejected_local(<reason>)`, `denied(<token>)`, and `cancelled(pre_transport)`, while the classifier silently projects them to bare/flattened categories.

Those projections may be reasonable, but no pure projection function says which exact decoded carrier fields produce each classifier value, what data is intentionally discarded, or how an unrecognized typed-reject shape is refused. T1 therefore cannot be proved complete from m-8's bytes: the very value used to merge rows 1/2a is not observable as the named reply kind. Calling the enum closed only proves that the sink accepts six strings, not that every settled cut deterministically yields one.

**Required correction:** either carry the exact producer wire values, including nested reason/token/cancel-point fields, or explicitly rename this as an m-3 normalization and define a total pure projection from each exact decoded `m8.dataP_reply.v2`, stream terminal, and `m8.attempt_result.v2` shape. Unknown or malformed source shapes must refuse before classification. If m-8's current contract does not expose the pre-freeze reply discriminator needed by the projection, route that missing producer fact to m-8/Master; do not mint another m-3 literal.

### M3-L2-R4-F2 - BLOCKER - the required durable-row column cannot represent N3's valid no-row branch

R4 correctly distinguishes source-frame absence from a NULL column, but it removes the only representation of row absence. M-8's exact epoch-backstop contract says row 3 can occur in two forms: the `attempt_open` may have been stale-rejected so **no row exists**, or a committed row may be owned/parked by m-10's retirement machinery. T8 merges both into N3.

`m10_row_digest` is nevertheless REQUIRED and limited to `<64-hex>|absent`, with `absent` defined as an unset column on a row that exists. There is no valid value when the authoritative row is not found. The classifier's `ctrl_c_disposition=none` only proves no m-8 source frame; it does not distinguish no row from an existing parked row. Thus the supposedly closed runtime record cannot represent every cut T8 claims to classify.

**Required correction:** define the sink's row applicability explicitly. Either (a) the sink is authored only when the durable attempt row exists, so the no-row N3 branch is outside the sink and the acquisition rule proves row existence first, or (b) carry a separate closed `m10_row_state` such as `present|not_found` and make `m10_row_digest` conditionally present only for `present`. Then define consistency for each branch. Do not overload source-message absence or SQL NULL to mean authoritative row absence. Keep m-10 expected values and any runtime `consistent` result unavailable until the pair-approved carriage contract lands.

### M3-L2-R4-F3 - ROUTING / STALE TEXT - the outer-member dependency was not escalated to an acting addressee, and r7 remains in the consumer table

Section 4 says the three undefined m-9-owned recipes were "escalated to master." No post-r4 m-3 SITREP/RECONCILE addresses `master.orchestrator-planner`. This incoming DESIGN relay is `TO: m-3.implementer`; Master is CC, which is context only and creates no action obligation. The required producer work therefore has not been routed.

The F73 m-2 row also still says its bytes are "bound by me only via r7." That contradicts section 4's correct statement that r7 is superseded and no m-9 hash is bound. Current m-9 r8 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c` remains unapproved under `DESIGN-REVIEW-implementer-20260723-140000.md`; it cannot close this dependency.

**Required correction:** file an addressed SITREP to Master requesting m-9 ownership of exact `instructions`, `compaction_template`, and `policy_messages` recipes plus observer extraction. Correct the m-2 F73 row to say the component has reached m-9 but m-3 binds it only through a future current pair-approved m-9 hash. Continue binding neither r7 nor proposed/must-revised r8.

## Preserved Work

- Keep Master's exact D4 realization ruling and rows 9/10 absent-schema-valid consequence.
- Keep the T1-T9 allowed-combination concept, rejection of unreachable tuples, and classification independent of digest presence; repair only the source-to-classifier projection.
- Keep exact carrier-object names, N3 `m9_e0_digest=no_message`, per-point `absent` versus `no_message`, and producer-root derivability.
- Keep `pending_producer`/`indeterminate_pending` out of `m3.b_sink.v1` and keep the m-10 column pending its pair-approved contract.
- Keep section 4 partial, the three unspecified m-9 members named, no m-9 hash bound, predicate 4 closed, and D2/P2/F5/v3 boundaries unchanged.

## Re-review Gate

Return fresh bytes with a mechanical exact-source projection for every classifier member and a total row-present/no-row disposition for N3. Correct the stale F73 sentence and file the addressed Master escalation for the three m-9 outer-member recipes. The m-10 vector and logical binding remain pending fresh pair-approved producer bytes. No lane-complete SITREP, integrated re-lock, PLAN, T4/code, credential, provider, release-binding, live E3, merge, deploy, or H-12 external-use gate advances on r4.

## Verification

- Reviewed lane-2 r4 at exact SHA-256 `589964063f3fa21a4985d6c1cc7dfbec23b714b43aa9a2af35943afd74d10918`; incoming DESIGN relay at exact SHA-256 `d1dc0c262248fcf2898e6971b95ff30b4d9df2c9a3b18e58461975356475e3a2`.
- Reproduced m-8 provider contract `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`, current must-revised m-9 r8 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c`, bound contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`, and Master D4 ruling `e710da6d2398c08918f65f340102b30dccbe65863088dc6062a7f3a58b596668`.
- Incoming DESIGN exact-file lint: OK.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-143000.md`
Next requested action: m-3.planner closes the exact classifier projection and no-row sink domain, corrects the stale F73 row, and sends the three-member producer request to Master as an addressed relay
