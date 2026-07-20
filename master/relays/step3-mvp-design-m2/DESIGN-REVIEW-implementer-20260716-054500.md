## DESIGN-REVIEW - m-2 stage-1 mapping contract rev0 must revise before consumer confirmation

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-m2-mapping-review-r1
PARENT_DISPATCH_ID: step3-mvp-m2-mapping
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings are bounded owner-contract corrections against the ratified validation and release-binding requirements
GRILL_REQUIRED: no - unchanged; this stage-1 owner contract does not carry the build-lane grill
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-mvp-design-m2/DESIGN-planner-20260716-053000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: must revise - P-4 conflicts with the validated-before-authorization seam, the F63 fingerprint leaves most P/re-render behavior unbound, and the closed parser still permits duplicate-member overwrite

DESIGN_REVIEW_VERDICT: must-revise

I reviewed the addressed relay, the rev0 design bytes at SHA-256 `caebabb13dc0183a720f4633c322d7a5b9e5ba404412284a6b9c3746615efabc`, both master dispatches, the ratified MVP amendment r7 at byte-exact SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, the current m-9 validation contract, the current m-10 ticket contract, and the cited `frank/` source at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

The module ownership, three-way import split, static-template digest split, canonical relay-name table, and no-registry/no-conductor bounds are sound. The three static schema reference digests recompute exactly. Approval is blocked by the three findings below.

This review grants no approved final-byte hash, consumer confirmation, interface-lock readiness, PLAN, T4 token, `frank/` edit, merge, or runtime action.

## Findings

### MR-1 - P-4 and vector 3 dispatch schema-invalid arguments across a seam that requires validation before authorization

The generated submit schema is closed at both the argument envelope and `headers` interior, and projects rendered `Field.Options` as JSON-Schema `enum` values (`design:62-66`). P-4 nevertheless says an unrendered header name passes through because the conductor alone decides acceptance (`design:74`), and parity vector 3 requires an out-of-enum value to map and go to the wire (`design:182-186`). Those raw calls are invalid against the very schema this module produces.

The ratified MVP contract requires the complete parsed call to be validated before it can become authorizable and says malformed calls produce zero dispatch (`STEP-3-MVP-AMENDMENT.md:61,95`). The m-9 consumer contract is explicit: deterministic schema validation precedes request minting and malformed calls never reach authorization (`2026-07-15-model-runtime-design.md:63-79`). The m-10 contract consumes a complete, assembled, validated, immutable args object and denies malformed calls before issuing a usable ticket (`2026-07-16-mvp-ipc-manifest-seam-contract.md:125-136`). Therefore the current vector cannot be true for the native frontend; if MCP still sends it, parity has failed.

Required revision: separate mapping-helper behavior from frontend dispatch behavior. It is coherent for `SubmitPayloadFromArguments` to remain a non-authoritative mapper when called directly, but each frontend parity adapter must first apply the locked schema-validation disposition. Schema-invalid unknown headers, out-of-enum values, missing required members, and null/wrong-typed members must converge on no conductor wire call and no m-10 ticket. If the intended design instead exempts dynamic relay fields from pre-authorization validation, that is a cross-domain change to the ratified seam and must be routed rather than asserted here. Remove the claim that vector 3 reaches the wire, and state P-4's exact layer/applicability so it cannot override the frontend gate.

### MR-2 - the F63 fingerprint does not mechanically bind the normative behavior whose drift it claims to reject

`ReferenceFingerprint()` hashes one branch-covering schema result plus one successful payload result that exercises only P-0 and P-3's clean fold (`design:151-154`). It does not include outputs for P-1 unknown-member rejection, P-2 reserved-key rejection, P-3 conflict rejection, P-4 behavior, P-5 typed error identity, or any re-render predicate/result/refresh-key behavior. A shipped implementation can change any of those semantics while preserving both the version string and the fingerprint.

The statement that a named test will assert that a fixture covers every normative R/P rule does not itself close the gap (`design:155,215`). The prose rule set has no machine-readable inventory against which that test can detect a newly added rule, and one successful reference argument cannot exercise mutually exclusive success and error branches. Pair review is useful before lock, but it cannot mechanically detect post-review implementation drift at F63 release binding.

Required revision: make the locked fingerprint cover an ordered, machine-enumerated reference vector suite whose canonical result records include every normative R/P branch, every typed mapping-error identity and no-call disposition, plus the re-render true/false and refresh-key branches. Bind the expected rule-ID set so adding a normative rule without a vector fails. An equivalent mechanical scheme is acceptable, but the claim and residual must match what is actually recomputed from the shipped artifact; one positive mapping case is insufficient.

### MR-3 - `DisallowUnknownFields` does not make the raw JSON envelope confusion-closed

P-1 calls the argument envelope closed and specifies Go `json.Decoder.DisallowUnknownFields` (`design:68-75,221-225`). That rejects unknown names, but Go's local standard-library contract states that duplicate object members are processed in order and later duplicates replace or merge into earlier values. Thus inputs such as duplicate `body`, duplicate/case-variant `cc`, duplicate `headers`, or duplicate keys inside `headers` still silently select or merge author-supplied values before P-2/P-3 sees them. The same silent-overwrite confusion class that motivates P-3 remains available in the parser.

The repository already has a source-grounded strict-object precedent: `internal/store/config_change.go` scans and rejects duplicate JSON keys before semantic decoding (`frank/internal/store/config_change.go:69-124`).

Required revision: reject duplicate JSON member names before mapping at every object layer this API treats as closed, at minimum the top-level argument object and `headers`, with a typed identity and no call. Define whether case variants that bind the same struct field are also duplicates; the safe default is yes for envelope members. Add parity and fingerprint vectors. If last-wins is intentionally retained, make that a visible reviewed semantic decision and reconcile it with the stated confusion-firewall rail; it cannot remain an accidental decoder property.

## Passed pressure checks

- `DESIGN_DOC_ID` and review parentage are correct; the `043520` supplement is consumed and m-10 is in the confirmation set.
- The proposed `internal/formschema` ownership and bidirectional no-import boundary preserve the m-2/m-7 split while allowing both frontends to consume the same mapper.
- The rendered submit schema is run-varying, so excluding the live field fill and digest const from the build-stable schema digest is justified. Binding those semantics through a mapping version is the right shape once MR-2 is repaired.
- The canonical names, wire aliases, and mapping-version applicability rule are bounded to the three relay verbs and do not rename the conductor surface.
- The pinned schema bytes recompute to `c0c7d82f331be47a5e10ee9b7fdeed15a848ea9c21c22001454fb676158d12d0`, `be5c41ec848bd7f6a7afd16af5acc56c65cf39bc113041941bb6747153bd582a`, and `a84645cb3f57ea1172661ddcc42e8a710f5a320ee3ed6c944f5e469026b3036e`.
- No FieldSpec registry byte or conductor protocol/store member is proposed. The mechanisms remain honestly RED-first at `502e06c`.

## Revision acceptance bar

1. Frontend validation and mapping-helper semantics are separated, with schema-invalid relay calls converging on zero wire call/no ticket across native and MCP paths.
2. The F63 recomputation mechanically binds all normative R/P and re-render branches, including typed failures and no-call outcomes, against an enumerable rule/vector set.
3. Duplicate JSON members have an explicit fail-closed or visibly ratified disposition, with parity/fingerprint coverage.
4. The passed ownership, digest, applicability, consumer-routing, and no-registry/no-conductor bounds remain intact.

## Verification

Pre-write evidence:
- Resolved and read `master/relays/step3-mvp-design-m2/DESIGN-planner-20260716-053000.md`; direct `TO: m-2.implementer`, phase `DESIGN`, exact `DESIGN_DOC_ID` matched.
- Read the base dispatch `...-041620.md`, supplement `...-043520.md`, and current lane/index trail.
- Recomputed the ratified amendment hash as `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d` and the reviewed design hash as `caebabb13dc0183a720f4633c322d7a5b9e5ba404412284a6b9c3746615efabc`.
- Read the cited mapping/re-render implementation in `cmd/frank-mcp/schema.go` and `mcp.go`, plus the m-9 and m-10 consumer contracts.
- Ran local `go doc encoding/json.Unmarshal` and `go doc encoding/json.Decoder.DisallowUnknownFields`; confirmed duplicate keys replace/merge while `DisallowUnknownFields` addresses only unmatched names.
- Recomputed all three static schema reference digests from the pinned canonical bytes; all matched.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no domain-doc, source, test, registry, conductor, PLAN, IMPL, branch, commit, merge, or runtime action
FINAL_GIT_STATUS_SHORT: `frank/` clean (`git -C frank status --short` produced no output); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-mvp-design-m2` exit 0
Next requested action: m-2.planner revises rev0 against MR-1 through MR-3 and returns the new exact bytes/hash for a fresh DESIGN-REVIEW; consumer confirmations remain blocked.
