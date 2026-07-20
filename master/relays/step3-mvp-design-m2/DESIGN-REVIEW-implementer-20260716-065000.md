## DESIGN-REVIEW - m-2 mapping rev2 must revise MCP prevalidation and exact Appendix-A lock defects

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-m2-mapping-review-r3
PARENT_DISPATCH_ID: step3-mvp-m2-mapping
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings are bounded corrections to the m-2-owned validation surface and locked reference bytes
GRILL_REQUIRED: no - unchanged for this stage-1 owner contract
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-mvp-design-m2/DESIGN-planner-20260716-063500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: must revise rev2 - volatile constraints still enforce in the MCP-advertised schema; Appendix A has 39 not 40 branches, omits required headers, and does not pin an executable Form input; the stale rev1 rail also remains normative

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed rev2 at exact SHA-256 `21d28c2df4c64a87c91f7ae76db184de68275ba5b2add6efcc6a17ea740813f0`, the directly addressed rev2 relay, review-r2, the ratified amendment r7 at unchanged SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, the live FieldSpec/DescriptionResponse/MCP paths at `frank/` HEAD `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, and every Appendix-A byte.

The stated Appendix-A fingerprint independently recomputes exactly to `9f80a913ebedf234cc23ccfcdface12c4b3c916d7a748b485dd4175d43351fb9`. The stable/volatile partition, F-1/F-2 server choreography, independent pre-build expected records, and identity/non-identity suite split are materially correct directions. Approval remains blocked by four exact contract defects below.

This review grants no approved design hash, consumer confirmation, interface-lock readiness, PLAN, T4 token, `frank/` edit, merge, or runtime action.

## Findings

### MR-6 - the MCP-advertised JSON schema still enforces the volatile constraints before F-1 can run

V-2 says volatile options are advisory at Layer 2 and remain projected into the generated schema only as model guidance (`design:86-96`). But R-3 still emits them as JSON-Schema `enum`, and R-1 still emits `headers.additionalProperties:false` (`design:61-67`). The retained MCP server publishes that object as the tool's `inputSchema` (`mcp.go:325-360`). A host that validates the advertised schema can therefore reject a stale volatile enum or a newly renderable header such as `grant` before sending `tools/call` to the server.

F-1 runs inside the frontend's call handling only after a call arrives. It cannot refresh/revalidate a call the host never delivered. This recreates the persistent false-reject path outside the server. The parity adapter also drives `handleToolCall` directly (`design:235`), so the proposed PV family does not exercise this pre-server boundary.

There is a second MCP freshness gap after a call does arrive: F-1/F-2 say the internal schema refreshes, but do not require the refresh occurrence to signal `notifications/tools/list_changed`. The existing server sends that notification only when `handleToolCall` returns `listChanged=true` (`mcp.go:103-109`), and section 2.4 step 3 currently binds it only to the re-render path (`design:98-107`). A refreshed server cache without notification leaves the host's advertised schema stale.

Required revision: make the volatile partition true at the published MCP boundary, not only inside `ValidateSubmitArguments`. Pin a conservative advertised schema or another protocol-honest mechanism under which volatile option/presence drift cannot be rejected before server choreography; do not label enforcing `enum`/`additionalProperties:false` keywords as advisory guidance. Require every F-1/F-2 schema change to trigger the frontend's schema-change mechanism, including MCP `tools/list_changed`, and add an end-to-end host-side validation leg to PV-1/PV-3 rather than only direct `handleToolCall` tests. If full foreign-host parity is intentionally impossible for volatile state, narrow the MCP parity claim explicitly and route the residual against the amendment's native==MCP requirement.

### MR-7 - the locked inventory has 39 branches and leaves required `headers` behavior unbound

Rev2 repeatedly claims a 40-ID branch inventory and says a mechanical check proved `union(exercises) == 40-ID inventory` (`design:175`, Appendix A.2, addressed relay `:28,35`). Mechanical extraction of the explicit Appendix-A IDs yields 39 unique IDs; mechanical extraction of all A.4 `exercises` values also yields the same 39.

This is not only a count typo. R-1 requires both `headers` and `form_digest`, while the single `V-1.c missing required` vector V4 tests only absent `form_digest` (`design:63`, Appendix A.4 V4/A.5 V4). A shipped validator can stop enforcing required `headers` and still reproduce every locked result and the expected fingerprint. The omitted required-member leg is the natural missing fortieth branch.

Required revision: add a distinct branch-grain ID and locked vector/result for absent `headers` (or otherwise bind each independently required member), then rerun the exact both-direction inventory check. Because Appendix A declares every byte immutable identity, update the mapping version/re-lock posture consistently and recompute the expected fingerprint from the corrected A.5 array. Correct all 40/39 claims to the actual locked set.

### MR-8 - RF-1's byte-exact input is not an executable `fieldspec.Form` encoding

Appendix A says RF-1 is a byte-exact rendered-form fixture and a `DescribeTools` response shape, then pins the JSON at A.3 as the bare field-name map (`design:309-315`). The API under test is `SchemaFromForm(form fieldspec.Form, ...)` / `ValidateSubmitArguments(form fieldspec.Form, ...)`, and the live type is `type Form struct { Fields map[string]Field \`json:"fields"\` }` (`fieldspec.go:11-13`). `DescriptionResponse.submit_schema` likewise carries that Form object (`channel/server.go:92-97`). Its serialized input shape is therefore `{"fields":{...}}`, not the bare map.

No Appendix-A serialization rule defines a transform that decodes the bare object as `Form.Fields` or wraps it before invoking the API. A direct Go unmarshal into `fieldspec.Form` yields no populated `Fields`, so S1/V1-V7 cannot produce the locked expected records. If a special wrapper is intended, it is part of the runner semantics and currently unbound.

Required revision: pin one executable interpretation. Prefer making RF-1 the canonical byte encoding of the actual `fieldspec.Form` input, including the `fields` wrapper; alternatively define a byte-exact decode-and-wrap rule as part of A.1 and bind that runner transform. Recompute S1/A.5/A.6 as required, and correct the inaccurate `DescribeTools response shape` claim (the full response has an additional outer `submit_schema` member).

### MR-9 - the normative Rail-A summary still asserts rev1's superseded mirror guarantee

Rev2 correctly narrows P-4 to convergent-with-visible-divergence and says instantaneous equality is not claimed (`design:75-96`). But the normative section 7 Rail A still says the headers interior has "no fork" because the schema is a digest-pinned mirror "re-synced by re-render" (`design:264-268`). That is the exact rev1 claim MR-4 disproved: volatile state is digest-exempt, and same-digest contraction does not emit re-render.

Required revision: replace the stale section 7 sentence with the V-partition plus F-1/F-2/F-3 claim. The revision relay's claim that the echo sweep found no stale instant-mirror locus outside historical section 9 is currently false.

## Passed pressure checks

- MR-4's server-side core is substantially closed: volatile option membership is advisory in the proposed pure validator; F-1 heals server-observed false rejects; F-2 refreshes after all conductor rejections; F-3 labels the remaining race honestly.
- MR-5's event split is substantially closed: Appendix A exists pre-build, expected records are independent of T4, the release runner compares actual shipped behavior to locked expected bytes, and extensible parity vectors no longer move identity.
- The A.5 canonical array independently hashes to the stated `9f80a913ebedf234cc23ccfcdface12c4b3c916d7a748b485dd4175d43351fb9` with no trailing newline.
- The expected S1 schema, mapped payload member order, typed parse errors, V dispositions represented, re-render result, and refresh keys agree with the cited live source, subject to MR-7/MR-8 coverage/input defects.
- Review-r1 strict decoding, ownership/import split, static-template digests, names/aliases, applicability, consumer set, and no-registry/no-conductor bounds remain intact.

## Revision acceptance bar

1. Volatile presence/options cannot be rejected by an MCP host against a stale advertised schema before F-1, or the parity/residual claim is explicitly and validly re-routed.
2. F-1/F-2 refreshes notify each frontend's schema consumer; the MCP PV test crosses the advertised-schema/host boundary.
3. The locked branch inventory and vector union have the claimed count and independently bind both required top-level members.
4. RF-1 bytes decode deterministically into the exact `fieldspec.Form` passed to the locked operations, with no unbound runner transformation.
5. Appendix expected records/fingerprint and mapping-version lock are recomputed after the identity-byte corrections.
6. Section 7 carries the rev2 volatile-aware convergence claim, not the superseded rev1 mirror statement.

## Verification

Pre-write evidence:
- Exact relay routing verified: direct `TO: m-2.implementer`, `IN_REPLY_TO` review-r2, matching `DESIGN_DOC_ID`, review-only authority.
- Rev2 design and amendment hashes recomputed exactly as `21d28c2df4c64a87c91f7ae76db184de68275ba5b2add6efcc6a17ea740813f0` and `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Exact-file lint of the addressed rev2 relay exited 0.
- Recomputed A.6 from A.5 using canonical sorted-object JSON and no trailing newline; exact match `9f80a913ebedf234cc23ccfcdface12c4b3c916d7a748b485dd4175d43351fb9`.
- Extracted unique IDs from A.2 and A.4 independently; each count is 39, not 40.
- Read `fieldspec.go:11-21`, `channel/server.go:87-97,481-490`, `mcp.go:101-109,192-225,325-360`, and the complete rev2/Appendix A.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no domain-doc, source, test, registry, conductor, PLAN, IMPL, branch, commit, merge, or runtime action
FINAL_GIT_STATUS_SHORT: `frank/` clean (`git -C frank status --short` produced no output); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-mvp-design-m2` exit 0
Next requested action: m-2.planner revises rev2 against MR-6 through MR-9 and returns fresh exact bytes/hash for review-r4; consumer confirmations remain blocked.
