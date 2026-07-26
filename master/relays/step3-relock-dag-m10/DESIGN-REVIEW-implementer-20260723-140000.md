## DESIGN-REVIEW - approve m-2 relay.submit resource rev1 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-c-m2-submit-resource-review-r2
PARENT_DISPATCH_ID: step3-relock-c-m2-submit-resource
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair approval only; bounded amendment, VP review, and operator ratification remain separate
GRILL_REQUIRED: no - master already ruled the form-schema-derived direction
DESIGN_DOC_ID: step3-relock-c-m2-submit-resource
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260723-120000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, operator
SUBJECT: APPROVE exact-byte m-2 relay.submit canonical_resource rev1 at SHA-256 5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f

DESIGN_REVIEW_VERDICT: approve

I freshly reviewed rev1 at exact SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`, the directly addressed rev1 relay at `1817ad824d7375b9d58e5d4f6d029c123c4e0b4bf33302131509d3ae514dde69`, review-r1, master's ruling, ratified amendment rev12, frozen m-2, the approved stage-6 E component, and the current delivery projection.

**APPROVE** the m-2 stage-6 section 5-C `relay.submit` `canonical_resource` shape at exact SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.

This approval is byte-bound. Any change to the design document, including metadata or revision history, voids it and requires fresh pair review.

## Closed review bar

- **M2-C-R1-F1 closes.** Optional top-level `cc` now participates in the closed target projection. A non-empty JSON string-array binds as logical member `cc`; every other non-empty string binds as `cc_unparsed`; empty or absent omits both. The branches are exhaustive and mutually exclusive, so schema-valid but conductor-rejected input remains derivable without claiming acceptance.
- **CC target movement closes.** RV-4 holds form, lane, and primary recipient fixed while adding CC and moves the resource. RV-5 proves elementwise/cardinality participation. RV-6 proves the unparsed branch is distinct. The build obligations carry the same target-moving, observer-parity, and totality legs.
- **Authority remains unchanged.** The design now distinguishes mailbox delivery effects from action authority: binding a CC delivery target grants no authority and does not alter the standing TO/CC protocol.
- **Owner boundary is honest.** The descriptor binds targets as named by the observed invocation and does not reimplement store deduplication, trimming, TO/CC union, or fallback. The explicitly disclosed consequence that differently named invocations may converge on one eventual delivery set is compatible with the invocation-grain descriptor claim.
- **Value-grain canonicalization closes.** Valid CC arrays are decoded to logical values before JCS, so equivalent JSON serialization spellings converge. Malformed, non-array, and mixed-type strings remain total under `cc_unparsed`; no conductor state or acceptance result enters derivation.
- The prior passed set remains intact: non-empty target direction, form/lane/primary-recipient identity, resource-versus-args separation, invocation-not-acceptance timing, JCS/lowercase-SHA/prefix determinism, observer derivability, and additive frozen-byte scope.

## Exact evidence

- Rev1, request, frozen m-2, approved E component, and amendment hashes independently reproduce as `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`, `1817ad824d7375b9d58e5d4f6d029c123c4e0b4bf33302131509d3ae514dde69`, `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`, `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`, and `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- RV-1 through RV-6 independently recompute to the six published values: `7f6479f5...`, `b98691f9...`, `01c6d57c...`, `cca45c4f...`, `24793638...`, and `e84184b1...`.
- Partition probes place valid/equivalently serialized string arrays in `cc`, and `null`, objects, mixed arrays, and malformed text in `cc_unparsed`; no non-empty case is unclassified or double-classified.
- Full-doc sweep finds no stale three-member definition. The normative object, member table, exclusion text, observer recipe, amendment cell, vectors, obligations, and revision record agree.
- Current source confirms accepted CC addressees are mailbox delivery recipients (`frank/internal/store/projections.go:137-174`) and the submit surface carries CC as a top-level string (`frank/cmd/frank-mcp/schema.go:29-35,101-103,122-128`).

## Authority boundary

This is m-2 pair approval only. It does not itself amend the ratified stage-6 bytes, satisfy VP review or operator ratification, complete the section-D join, authorize an m-9 fold, grant PLAN/T4/implementation, or authorize source, registry, credential, provider, merge, deploy, or runtime action. Master owns the bounded amendment and downstream routing.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: none - clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-relock-dag-m10` exit 0
Next requested action: m-2.planner SITREPs master with the pair-approved exact bytes/hash; master carries the shape in the bounded section-D-settlement amendment through VP exact-byte review and operator ratification.
