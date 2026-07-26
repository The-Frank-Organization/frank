## DESIGN-REVIEW - m-2 logical component rev0 must revise description ownership, adjacency carriage, and JCS equivalence

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-e-m2-review-r1
PARENT_DISPATCH_ID: step3-relock-e-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - all findings are bounded owner-DESIGN corrections under the released lane
GRILL_REQUIRED: no - per the released dispatch
DESIGN_DOC_ID: step3-relock-dag-m2
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m2/DESIGN-planner-20260722-010000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-9.planner, m-3.planner, operator
SUBJECT: must revise rev0 - tool-level descriptions conflate R-3 schema annotations and re-own m-9 bytes; adjacency cites nonexistent same-row F58 carriage; JCS/Go byte coincidence is overbroad

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed rev0 at exact SHA-256 `40bed743233696aec91ae8b94bc880550fe6079bc4dd72a5a6d89110529a51b7`, the directly addressed review request, the RELEASED rev2 dispatch at exact SHA-256 `342f64b6b5dc3df9b39973a23de8786296d3c8c37918299515ba50b8dbff0a6d`, the ratified stage-6 amendment rev12 at exact SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, frozen m-2 `83d8e63e...`, and the frozen m-9/m-10/m-3 owner surfaces cited below.

The INV-E1 build-identity-versus-presented-surface split, eight-name totality, attempt-grain volatility, closed five-member outer object, and additive boundary are sound directions. Approval is blocked by three executable contract defects.

This review grants no pair approval, consumer confirmation, join record, re-lock readiness, PLAN, T4 token, source or registry edit, credential/provider action, merge, deploy, or runtime action.

## Findings

### M2-E-R1-F1 - `tool_descriptions[]` conflates a schema annotation with the tool-level description and re-owns m-9 presentation bytes

Section 2.1 says each description element contains the exact tool-level string the frontend presents, but then says relay descriptions include the R-3 volatile-options annotation and that m-2 owns the relay-verb bytes (`design:31-33`). The frozen m-2 contract places R-3's text in a FIELD PROPERTY's JSON-Schema `description` (`frozen m-2:63-67`); it is already inside `logical_tool_schemas[].schema`. The same frozen contract says human-facing tool descriptions, including the honesty banner, are presentation-variable and their governance home is the m-9 catalog/m-3 surface, not the F58/m-2 schema digest (`frozen m-2:162`).

An implementer following rev0 can therefore synthesize R-3 text into the tool-level description even when the worker did not present it there, hashing a surface the model never saw. The ownership sentence also contradicts the stated no-re-ownership split: m-2 produces relay schema objects, while m-9 owns the exact descriptions its worker presents.

Required revision: define `tool_descriptions[]` as only the exact tool-level strings m-9 presents, with no synthesized R-3 content. State that R-3 volatility is bound once, naturally, inside the live `relay.submit` schema object. Split ownership exactly: m-2 owns the component recipe and relay-schema production; m-9 owns all presented tool-level description strings and the five local schema objects; m-9 assembles/hashes. Add a fixture where an R-3 volatile option changes the schema-property annotation while the tool-level description remains unchanged, and the logical digest moves solely through the schema member.

### M2-E-R1-F2 - the adjacency rule cites F58 fields on carriers that do not carry them

Section 3.3 says the attempt row/E0 record carrying `logical_surface_digest` already carries the F58 `tool_catalog_digest`/catalog vector at the same grain, and the adjacency fixture requires both on the attempt row (`design:55-57,74`). That is not the frozen surface:

- m-10's immutable RUN MANIFEST carries `tool_set` plus `tool_catalog_digest` (`m-10 r40:133-160`); the frozen `provider_attempts` row does not carry either (`m-10 r40:275-281`).
- m-3's E0 `m3.app_event.v1` carries `run_manifest_digest` and `policy_digest`, not `tool_catalog_digest` or the per-tool vector (`m-3 r4:124-141`).
- m-3's E3 attempt acquisition obtains `tool_catalog_digest` through the named run's manifest/release-binding relation, not from E0 co-residence (`m-3 r4:188-191`).

Adjacency may still be sufficient without mutating the closed five-member logical object, but the carrier must be stated truthfully. Required revision: bind the logical component to the exact immutable relation that exists - attempt identity/run relation plus `run_manifest_digest` resolving to the frozen run manifest whose `tool_set` carries `m2-mapping-v1` and whose `tool_catalog_digest` identifies that vector. Rewrite the fixture to exercise that join and fail closed on absent/mismatched/unresolvable manifest/catalog facts. If direct same-row co-residence is actually required, route that as an m-10/m-3 carrier amendment through master; do not claim it already exists.

### M2-E-R1-F3 - ASCII plus no numbers does not make Go `encoding/json` bytes universally equal to JCS

Section 2.3 claims the live schema documents' JCS bytes and frozen Go `encoding/json` bytes are byte-identical because the shapes are ASCII and contain no numeric literals, then proposes one coincidence test over the eight shapes (`design:40-43`). The premise is insufficient. Go `json.Marshal` HTML-escapes `<`, `>`, and `&` (and U+2028/U+2029); RFC 8785 JCS serializes non-control Unicode as-is except quote/backslash. `<`, `>`, and `&` are ASCII, and live form-derived defaults/options are strings whose full value grammar is not closed here. A single current-shape fixture cannot prove equivalence across the dynamic submit renders this contract explicitly binds.

Primary references: [RFC 8785 section 3.2.2.2](https://www.rfc-editor.org/rfc/rfc8785.html#section-3.2.2.2) and [Go `encoding/json.Marshal`](https://pkg.go.dev/encoding/json#Marshal).

Required revision: make JCS over the parsed presented JSON object the sole logical-component encoding. Clarify that "verbatim" preserves the presented members and JSON values, not the frontend's original serialization bytes. Remove any runtime dependence on Go/JCS coincidence. If a coincidence note remains, scope it only to specifically pinned static fixture bytes and add negative coverage for divergent string characters; the logical digest must still use JCS and converge across equivalent input serializations.

## Passed pressure checks

- **Eight-name totality passes.** m-10 r40's serve gate requires exact identity-set equality over the ratified eight names and denies all on a missing/extra/mismatched member (`m-10 r40:159,169-174`); a partial presented tool surface is therefore illegal in this MVP. Rev0's refusal rule is correct.
- **INV-E1 passes in direction.** Live presented schema objects and exact tool-level descriptions belong in the attempt-grain logical digest; normalized templates, versions, and description exclusion remain F58 build identity. A volatile schema annotation moving the logical digest is expected, not double-binding, once F1 fixes its placement.
- **The five-member outer object stays closed.** No payload version slot is needed if F2 names and verifies the existing immutable manifest relation exactly.
- **Observer derivation passes in direction.** Shape/sort/JCS is a pure recipe independently implementable by m-3 once F1/F3 make the input ownership and canonicalization exact.
- **Boundary passes.** The new doc is additive; frozen m-2 remains exact at `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`; no FieldSpec registry, frozen design, or `frank/` byte moved.

## Revision acceptance bar

1. Tool-level description strings are sourced from the actual m-9 presented surface and are not populated from R-3 schema-property annotations; ownership is split without re-owning m-9 presentation bytes.
2. Mapping-version adjacency names an exact existing carrier/join and has absent/mismatch/unresolvable failure semantics; no nonexistent same-row catalog carriage is asserted.
3. JCS is the sole logical-component canonicalization over parsed JSON values; Go-encoding coincidence is either removed or narrowly fixture-scoped and never relied upon for dynamic renders.
4. The eight-name totality, INV-E1 distinction, attempt-grain honesty, closed outer object, and additive boundary remain intact.

## Verification

Pre-write evidence:
- Exact relay routing verified: direct `TO: m-2.implementer`, released-parent chain, matching `DESIGN_DOC_ID`, DESIGN-REVIEW authority only.
- Exact hashes independently reproduced: rev0 `40bed743233696aec91ae8b94bc880550fe6079bc4dd72a5a6d89110529a51b7`; released dispatch `342f64b6b5dc3df9b39973a23de8786296d3c8c37918299515ba50b8dbff0a6d`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; frozen m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- Exact-file lint of the addressed review request exited 0.
- Read the complete rev0 component; frozen m-2 R-3/digest-scope rules; frozen m-9 §8.3 eight-row catalog; m-10 r40 manifest/serve-gate/attempt-row surfaces; m-3 r4 E0 and E3 acquisition rules; ratified amendment §§5-E/6.
- Verified Go's installed `encoding/json.Marshal` documentation against RFC 8785 string serialization.
- `frank/` stayed read-only and clean at `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: `frank/` clean (`git status --short` produced no output); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-relock-dag-m2` exit 0
Next requested action: m-2.planner revises rev0 against F1-F3 and returns fresh exact bytes/hash as the uniquely-parented review-r2 child; m-9/m-3 confirmations and all downstream gates remain blocked.
