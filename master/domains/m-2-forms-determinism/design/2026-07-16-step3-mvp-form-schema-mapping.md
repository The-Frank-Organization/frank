# Step-3 MVP stage-1 owned contract — the form→tool-schema + submit-payload MAPPING module, the F58 relay-verb components, and the MCP↔native parity vectors

**DESIGN_DOC_ID:** `step3-mvp-design-m2-mapping` · **rev5** — the current revision; the full arc (rev0 → rev5: review-r1 MR-1/2/3 · review-r2 MR-4/5 · review-r3 MR-6/7/8/9 · review-r4 MR-10 · review-r5 MR-11) is recorded in the **§9 revision log**, which is the authoritative history — this marker names only the live revision and is updated with every revision.
**Dispatch:** `master/relays/step3-mvp-design-m2/DESIGN-orchestrator-planner-20260716-041620.md` + supplement `…-043520.md` (consumer set = m-9 + m-7 + m-10).
**Authority basis:** the ratified MVP amendment `master/STEP-3-MVP-AMENDMENT.md` r7 @ SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d` — §4 (F58/F55/F63), §5 (F42), §7 stage 1, §10 ("native == MCP relay path").
**Grounding base:** `frank/` main @ `502e06c` (`s11-close`), read-only.
**GRILL_REQUIRED:** no (per dispatch; stage-1 owner contract — pair review + consumer confirmation; grills ride the stage-4/5 build lanes).
**Owner:** m-2.planner authors; m-2.implementer pair-reviews the final bytes; m-9, m-7, m-10 confirm.

---

## §1 — Scope and boundaries

This contract owns exactly three things:

1. **The mapping module** (§2): the m-2-owned home for `SchemaFromForm` / `SubmitPayloadFromArguments` + the re-render semantics, hoisted out of `cmd/frank-mcp` (package `main`).
2. **The m-2 F58 build-identity components** (§3): the relay-verb tool-schema digests (`relay.submit`, `relay.project`, `relay.read`) + the form→schema mapping VERSION — canonical encoding + field applicability.
3. **The parity conformance vectors** (§4): the shared vectors under which the retained MCP frontend and the native tool produce equivalent conductor calls + re-render behavior.

**Explicit non-scope (dispatch bounds, honored):**
- **No FieldSpec registry byte moves.** Nothing here touches `internal/fieldspec/registry.json`; the mapping consumes rendered `Form` values, it defines no field. Any registry need discovered downstream routes back as its own reviewed changeset.
- **No conductor byte/member change.** The conductor wire protocol (`tools/call` with wire verbs `submit`/`project`/`read`, `tools/descriptions`, `session/connect` — `internal/channel/server.go`) is untouched. The canonical `relay.*` names are an **app-side policy identity** (§4/F58 of the amendment), not a conductor rename.
- **Not the transport.** `Call`/reconnect/socket lifecycle/`DialAuthenticated` are m-7's stage-1 contract (`step3-mvp-design-m7`). §2.1 pins the seam.
- **Not the consumer.** The native tool loop, the local tools, and the F59 executor half are m-9's.
- **v1 honesty rail:** at `502e06c` the mapping lives trapped in `cmd/frank-mcp` (`schema.go`, plus the re-render halves of `mcp.go`); the module, the digest producers, the fingerprint, and the parity harness **do not exist**. Every mechanism below is a RED-first build obligation (§6), not a description of extant code.

---

## §2 — The mapping module contract

### §2.1 Home, ownership, and the 3-way seam

- **Home (proposed path):** `frank/internal/formschema` — a new m-2-owned package. The *path token* may be adjusted by the build lane; the **ownership, API surface, and import boundary are normative**.
- **Ownership:** m-2 is the sole semantic owner. Changes to the mapping rules (§2.3) or the re-render semantics (§2.4) are m-2-reviewed changesets and move the mapping VERSION (§3.3).
- **The 3-way seam (amendment §5, mechanically enforced):**
  - the mapping module **must not import `internal/channel`** (it never dials, calls, reconnects, or sees a credential);
  - **`internal/channel` must not import the mapping module** (the transport half carries opaque `json.RawMessage` args and the `DescriptionResponse` it already owns — it neither absorbs nor re-encodes FieldSpec semantics);
  - **consumers** (the retained `cmd/frank-mcp` frontend and the m-9 native tool) import both and wire them together — the *only* place a rendered form meets a transport call.
  - **Enforcement:** a build test asserting the import graph (both directions), named in §6. "Neither absorbed nor stranded" — the module also must not be re-trapped inside either frontend binary's `main` package; both frontends must consume the *same* module (that shared consumption is what makes §4 parity meaningful).
- **Permitted imports:** `internal/fieldspec` (the `Form`/`Field` types + `CanonicalMarshal`) and `internal/record` (the `Record`/`Envelope` payload types) — both m-2-adjacent data contracts with no transport surface.

### §2.2 The API surface (normative signatures; hoist origin cited)

| API | Origin (prior art, `502e06c`) | Contract |
|---|---|---|
| `SchemaFromForm(form fieldspec.Form, digest string) map[string]any` | `cmd/frank-mcp/schema.go:11-47` | Rendered form + form digest → the submit tool input schema, per the R-rules (§2.3.1). |
| `SubmitPayloadFromArguments(args json.RawMessage) (json.RawMessage, error)` | `schema.go:90-120` | Tool arguments → the canonical `fieldspec.SubmitPayload` bytes, per the P-rules (§2.3.2). |
| `ParseSubmitArguments(args json.RawMessage) (SubmitArguments, error)` + exported `SubmitArguments` | `schema.go:122-129` (unexported `submitArguments`) | The single argument-envelope parser both frontends and `SubmitPayloadFromArguments` share; strict-decoded and fail-closed per P-1/P-6. |
| `ValidateSubmitArguments(form fieldspec.Form, digest string, args json.RawMessage) []Disposition` (+ trivial `ValidateProjectArguments`/`ValidateReadArguments`) | new (MR-1) | The dedicated dispatch-gate validator over the module's own generated schema shape (closed objects, string carriers, enum, const, required) — no general JSON-Schema engine. Ships beside the schema generator so both frontends (and m-9, at their option — §5) validate relay-verb calls with the same code; the parity vectors bind the disposition either way. |
| `DeclaredPhaseTier(args SubmitArguments) (phase, tier string)` | `mcp.go:228-238` | `headers["PHASE"]` else `"SITREP"`; `headers["CEREMONY_TIER"]` else `"medium"` — the re-render refresh key. |
| `SubmitNeedsReRender(result json.RawMessage) bool` | `mcp.go:240-253` | True iff the conductor outcome is `state=="rejected"` and the detail carries both `form_digest` and `re-render` (the conductor's stale-digest bounce). |
| `ReRenderResult(original json.RawMessage) json.RawMessage` | `mcp.go:255-276` | The normalized re-render violation payload (field `form_digest`, class `re-render`, the re-read hint; `relay_id`/`intake_id` passed through when present). |
| `ProjectSchema() map[string]any` / `ReadSchema() map[string]any` | `mcp.go:342-360` (inline in `mcpTools`) | The canonical static input schemas for the other two relay verbs (§3.2 pins the bytes). |
| `CanonicalNames() / WireName(canonical string) (string, bool)` | new | The §3.1 name table: canonical ID ↔ conductor wire verb ↔ known frontend aliases. |
| `ToolSchemaDigest(canonical string) (string, bool)` · `MappingVersion() string` · `ReferenceFingerprint() string` | new | The F58 producers (§3). |

The schema-refresh *loop* (`refreshSubmitSchema`, `cachedSubmitSchema`, the `DescribeTools` call and its cache — `mcp.go:192-226`) is **consumer logic**: it composes m-7 transport with this module and stays in each frontend.

### §2.3 The mapping rules (normative)

#### §2.3.1 Schema construction (the R-rules; behavior of `schema.go:11-88` preserved except where a delta is flagged)

**The advertised-schema principle (review-r3 MR-6):** the generated schema is *published* (the MCP `inputSchema`, `mcp.go:325-360`) and a conforming host may validate it client-side, BEFORE any call reaches our choreography — so **enforcing keywords (`enum`, `additionalProperties`, `required`, `const`, `type`) may carry only constraints whose drift forces a `form_digest` move** (those self-heal: a stale const rides the call in, the conductor bounces re-render, the schema refreshes and notifies). Digest-exempt volatile state must never appear in an enforcing keyword — a host-side reject leaves no call for F-1 to heal (the deadlock MR-6 named). **Annotations (`default`, `description`) are non-enforcing in JSON Schema and may carry volatile guidance.**

- **R-1 (envelope):** the submit schema is a closed object at the top level (`additionalProperties:false`) with `required = ["headers","form_digest"]` and exactly the members `headers` · `to` · `cc` · `dispatch_id` · `body` · `form_digest` (all build-static; `schema.go:19-47`). **R-1.b (AMENDED at review-r3 MR-6):** the `headers` interior is **advertised-open** (`additionalProperties:true`) — volatile-presence fields (`grant`, `render_test.go:90-137`: renderable under an unchanged digest) must not be host-rejectable; header-name enforcement lives at Layer 2 under F-1 freshness (server-side, healable, visible), not in the published schema.
- **R-2 (string carrier):** every header property is JSON-typed `string` regardless of FieldSpec type; structured types carry the canonical-JSON description hint — `row_array` / `object` / `address_list` per `structuredDescription` (`schema.go:50-75`). The conductor's `ParseTyped` (`internal/fieldspec/canonical.go:20-62`) remains the authoritative decoder; the schema hint is honest labeling, never validation authority.
- **R-3 (enum/default projection — AMENDED at review-r3 MR-6):** for **non-volatile** fields, `Field.Options` project verbatim, order-preserving, into `enum` (digest-bound, self-healing). For **volatile-flagged** fields (`conductor_volatile`/`digest_exempt`), options are **never emitted as `enum`**; they ride the `description` annotation as `live options (conductor-validated; may change without a form_digest change): <canonical JSON array>` (joined after any structured hint with `"; "`). `Field.Default` projects into `default` for all fields (a JSON-Schema annotation, non-enforcing — stale guidance surfaces as a visible conductor violation, never a host reject). The mapping never filters, reorders, or augments options — the rendered form (seat/phase/grant-scoped upstream, `render.go:96-146`) is the sole source.
- **R-4 (system-field absence):** rendered fields named `from`/`role`/`relay_id`/`delivery_state` (case-insensitive) and exactly `TO`/`CC`/`DISPATCH_ID` are absent from `headers` (`schema.go:77-88`) — the first four are system-stamped, the last three are carried by the top-level envelope members.
- **R-5 (digest pin):** `form_digest` carries `const: <digest>` — the schema is self-invalidating against form drift; the conductor's digest check is the authoritative gate, the const is the client-side mirror.

#### §2.3.2 Payload construction (the P-rules)

- **P-0 (pass-through, no synthesis):** headers copy verbatim; `to`/`dispatch_id` → `record.Envelope`; `body` and `form_digest` carry through (`schema.go:97-114`). The mapping never defaults, injects, or drops a header value — the conductor's form validation is the sole *acceptance* authority (dispatch-layer *validity* is P-4 Layer 2's stable-partition gate under the §2.3.3 freshness contract, not a competing acceptance gate). (An absent `PARENT_DISPATCH_ID` on a store-first relay stays absent — form semantics, not mapping semantics.)
- **P-1 (closed argument envelope, exact-name binding — DELTA, fail-closed):** `ParseSubmitArguments` rejects unknown *top-level* argument members with a typed error and **no conductor call** — where "unknown" includes **case-variant names**: envelope members bind by **byte-exact name only** (`headers`, `to`, `cc`, `dispatch_id`, `body`, `form_digest`). Go's `Unmarshal` documented case-insensitive field fallback would silently bind `CC`/`Body` to members the case-sensitive generated schema rejects — the parser must agree with the schema, so a case-variant name is an unknown member. The trapped implementation silently ignores unknowns (`schema.go:93` plain `Unmarshal`): a misspelled member (`headres`, `bod`) maps to a *thinner* record that the conductor may legally accept — silent data loss invisible to both the author and the gate. Rail A: at this surface, ignore-unknown changes what the author believes was submitted ⇒ CLOSED.
- **P-2 (reserved header keys — DELTA, fail-closed):** `headers` containing any of `from`/`role`/`relay_id`/`delivery_state` (case-insensitive) or `TO`/`DISPATCH_ID` is a typed error, no call. These are system-stamped or envelope-carried (the R-4 set); passing one through invites a stamp/claim collision the conductor may not uniformly name. (`CC` is handled by P-3, not banned.)
- **P-3 (the CC fold — pinned + conflict-closed):** top-level `cc`, when non-empty, folds to `headers["CC"]` (`schema.go:101-103`). **DELTA:** if `headers["CC"]` is *also* present and differs from a non-empty top-level `cc`, that is a typed error, no call (the trapped code silently overwrites — two disagreeing sources resolved silently is exactly the confusion class this team exists to kill). Equal duplicates fold cleanly.
- **P-4 (layered validation — REVISED at review-r1 MR-1; supersedes rev0's "interior openness"):** three layers, each with its own authority, none overriding the one below:
  - **Layer 1 — the mapper is non-authoritative.** `SubmitPayloadFromArguments` called directly is a faithful courier: it never re-checks enum membership, header presence, or required-ness (its P-1/P-2/P-3/P-6 strictness is *argument-envelope integrity*, not form semantics). It exists so validation and mapping cannot drift apart, not so mapping can bypass validation.
  - **Layer 2 — the frontend dispatch gate (normative for BOTH frontends).** The complete assembled call is validated against the generated schema before any wire call — **enforcing only the §2.3.3 STABLE constraint partition, under the §2.3.3 freshness contract (REVISED at review-r2 MR-4)**. Native path: m-9's lifecycle already pins the gate's position — deterministic schema validation precedes request minting, malformed never reaches authorization (`m-9 design §3`), and m-10's issue check denies malformed with zero dispatch and no ISSUED ticket (`m-10 §D.2`, amendment `:61,:95`). MCP path: the frontend **itself** validates arguments against the same generated schema before dialing (it must not rely on host-side schema enforcement) — the `ValidateSubmitArguments` helper (§2.2) is the shared implementation. **Schema-invalid ⇒ typed no-call disposition: no conductor wire call, no m-10 ticket — both frontends converge (parity-bearing, §4).**
  - **Layer 3 — the conductor stays the sole ACCEPTANCE authority** for record semantics. The dispatch gate does not fork acceptance: it enforces only digest-bound stable constraints (a faithful mirror of the digest-stable form), treats volatile state as advisory, and follows the §2.3.3 freshness contract — so the mirror is **convergent with visible divergence**, never a stuck fork: every residual mismatch surfaces as a typed conductor violation or a re-render bounce, and no silent path exists in either direction. (Rev0's worry that client-side validation "forks acceptance semantics" was wrong about the *stable* constraints at this seam — the ratified tool-dispatch contract requires validated-before-authorization — but rev1 over-claimed the mirror: digest-exempt volatile state CAN diverge under an unchanged digest, which is exactly why the partition + freshness contract exist. See §2.3.3.)
- **P-5 (typed errors and dispositions):** every P-rule failure returns a typed mapping error (distinct identities for P-1/P-2/P-3/P-6) and every Layer-2 failure a typed `schema_invalid` disposition — all converging on **no wire call, no ticket**. Error *presentation* (MCP `error_class`, native tool result shape) is the frontend's; the error/disposition *identity* is this module's.
- **P-6 (strict object decoding — NEW at review-r1 MR-3; DELTA, fail-closed):** before any semantic decoding, the raw argument bytes are token-scanned and **duplicate object member names rejected** at every object layer this API treats as closed — the top-level envelope, the `headers` value, and any nested object — plus **trailing data after the top-level value rejected**; typed error, no call. Go's documented decode behavior is last-wins/merge on duplicates, so `DisallowUnknownFields` alone leaves duplicate `body`, duplicate `headers`, or a duplicated key inside `headers` silently selecting one author-supplied value — the same silent-overwrite class P-3 closes. Source-grounded precedent: `internal/store/config_change.go:69-124` (`rejectDuplicateJSONKeys` + strict decode + trailing-JSON reject) — reuse the pattern, not a new one (Rail B). Duplicate comparison is byte-exact per layer; a case-variant *pair* at the top level (`cc` + `CC`) is not a P-6 duplicate but dies as a P-1 unknown member (exact-name binding), so no case-collision survives either way.

#### §2.3.3 The Layer-2 constraint partition + freshness contract (NEW at review-r2 MR-4)

**The problem (verified at source):** the rendered form carries state that changes **without moving `form_digest`** — `ConductorVolatile`/`DigestExempt` options and defaults (parent candidates, recipient candidates, grant, monotonic floors — `render.go:109-147`), stripped from the digest by `formForDigest` (`render.go:249-262`), proven by `TestRenderStableDigestIgnoresConductorVolatileClasses` (`render_test.go:90-137` — equal digests while recipient/monotonic options differ AND while `grant` goes absent→present). And on a volatile *contraction*, the conductor rejects with the **current** `enum`/`seat-scope`/`monotonic-floor` violation, not `form_digest`/`re-render` (`validate.go:23-26,55-65` — the digest is still current, so `staleForm` is false), so the re-render loop never refreshes a stale cache. A Layer-2 gate that enforced the whole rendered schema against a cached copy would therefore fork: a stuck client-side false-reject (expansion/grant, with no conductor bounce to heal it) or a stale accept the conductor visibly rejects but the cache never learns from. The fix is a partition plus a freshness contract; no conductor byte moves (`DescribeTools` and the volatile flags already ride the wire — `fieldspec.go:15-21` serializes `conductor_volatile`/`digest_exempt`).

**The constraint partition (V-rules — the pure validation projection `ValidateSubmitArguments(form, digest, args)`):**
- **V-0 (valid):** no disposition ⇒ proceed to mapping + wire call.
- **V-1 (STABLE constraints — Layer 2 rejects):** digest-bound state only: the required member set (`headers`, `form_digest` — class `required`); the digest pin (`form_digest` ≠ the schema's const — class `digest-mismatch`); headers-interior membership (a header name absent from the form, or in the R-4 skip set — class `unknown-header`); enum membership for fields **not** flagged volatile (class `enum`). (Member *types* and envelope shape are enforced earlier, at the P-1/P-6 parse layer.)
- **V-2 (VOLATILE constraints — advisory, Layer 2 passes):** for fields flagged `conductor_volatile`/`digest_exempt` in the rendered form, Layer 2 validates the carrier shape only (string — already guaranteed by parse) and **never enforces option membership client-side**. The conductor validates live membership and answers with its typed violation (`enum`/`seat-scope`/`monotonic-floor`) — visible, never silent. The generated schema projects volatile options as **description-annotation guidance only, never as enforcing keywords** (R-3 as amended; the advertised-schema principle, §2.3.1) — so the presentation/enforcement split holds at the published boundary too, not only inside `ValidateSubmitArguments` (review-r3 MR-6).

**The freshness contract (F-rules — frontend choreography, both frontends, parity-bearing):**
- **F-1 (refresh-before-reject):** a Layer-2 reject is surfaced as `schema_invalid` only after **one fresh `DescribeTools` + revalidation against the fresh form**; if the fresh form accepts, the call proceeds. A `DescribeTools` failure surfaces as a transport error class, never `schema_invalid`. This kills the stuck false-reject: volatile expansion and grant-absent→present are healed by the refresh even though the digest never moved.
- **F-2 (refresh-on-rejection):** **any** rejected submit outcome — not only the re-render bounce — triggers a schema refresh keyed per RR-3 (the re-render bounce additionally keeps its RR-2 normalized surface). This kills the stuck stale-accept: a volatile contraction reaches the conductor once, fails with a visible typed violation, and the next schema view is fresh.
- **F-3 (race honesty — the claim, stated exactly):** between a refresh and the subsequent submit a window remains in which volatile state may move again. Neither frontend can *remain* on a false rejection (F-1) or a stale acceptance (F-2), and every residual mismatch lands as a typed conductor violation or a re-render bounce — **there is no silent divergence path in either direction; instantaneous equality of the mirror is NOT claimed.**

**Vectors:** the V-branches are pure and live in the locked fingerprint reference set (Appendix A: V1–V8). The F-choreography needs a live/fake conductor and lives in the extensible parity suite (§4: the PV family — same-digest old/new-form pairs for expansion, contraction, grant absent→present, recipient candidates, monotonic floors).

### §2.4 The re-render contract

On a submit outcome where `SubmitNeedsReRender` is true, a conforming frontend must:
1. **(RR-2)** surface `ReRenderResult(outcome)` to the caller *in place of* the raw outcome (normalized violation: field `form_digest`, class `re-render`, re-read hint);
2. **(RR-3)** re-fetch the form via m-7 `DescribeTools` keyed by `DeclaredPhaseTier(args)` (declared headers, with the pinned `SITREP`/`medium` fallback) and re-invoke `SchemaFromForm` on the response;
3. signal schema refresh to its caller by its own frontend mechanism (MCP: `notifications/tools/list_changed`, `mcp.go:104-110`; native: m-9's design). **(AMENDED at review-r3 MR-6): the signal fires on EVERY schema refresh — the re-render path here, F-1's refresh-before-reject, and F-2's refresh-on-rejection — not only the re-render path.** A refreshed server cache without the signal leaves a host's advertised schema stale (the existing code signals only via `handleToolCall`'s `listChanged` return, `mcp.go:103-109` — the F-1/F-2 refresh points must feed the same mechanism).

The detection predicate itself is **RR-1** (`SubmitNeedsReRender`: `state=="rejected"` + the `form_digest`/`re-render` detail pair). RR-1/RR-2/RR-3 carry branches in the §3.3/Appendix-A branch inventory and are fingerprint-bound (both branches of RR-1, the RR-2 result bytes, and both RR-3 key derivations — declared and fallback).

Steps 1–2 are parity-bearing (§4 vectors assert them); step 3's *mechanism* is frontend-owned, its *occurrence* is parity-bearing. The detection predicate (the `form_digest` + `re-render` substring pair, `mcp.go:251-253`) is pinned as the v1 contract with an honest label: it keys on the conductor's rejection detail text; if the conductor ever moves to a structured violation object on this path, the predicate follows it in the same changeset — flagged as a fragility, accepted for the MVP (no conductor byte may move under this contract).

---

## §3 — The m-2 F58 build-identity components

The per-tool identity vector (amendment §4/F58) is `{canonical name, tool-schema digest, tool-implementation/catalog version, form→schema-mapping version}`.
m-2 produces, **for the three relay verbs only**: the tool-schema digests + the mapping version.
m-9 produces the local-tool digests + the catalog version; m-10 verifies the assembled vector at the serve gate (F55) and the F63 release-binding.

### §3.1 Canonical names + alias normalization

| Canonical ID (the policy identity, operator-ratified F58) | Conductor wire verb (m-7 surface, unchanged) | Known aliases (normalize-to-canonical) |
|---|---|---|
| `relay.submit` | `submit` | `submit` (retained MCP frontend tool name, `mcp.go:134,336`) |
| `relay.project` | `project` | `project` |
| `relay.read` | `read` | `read` |

- The canonical→wire map is mechanical (`strip "relay."`), but it is **pinned as a table, not a rule**, so a future verb cannot silently mint an identity by prefix arithmetic.
- Alias normalization happens **before** m-10's exact-set-equality gate (F55: "aliases normalize to canonical IDs first"); this table is the reference input to that normalization for the three relay verbs. The MCP frontend keeping its unprefixed tool names is legal aliasing; whether it renames is an m-9/m-10 lane choice — either way identity is the canonical ID.
- The wire verbs are the conductor's and do not move (§1 bounds).

### §3.2 Tool-schema digests — what is digested, and why the rendered submit schema cannot be

**The problem the encoding must solve:** the rendered submit schema is inherently run-varying — the `form_digest` const embeds the config digest (`render.go:228-247`), and `ConductorVolatile`/`DigestExempt` options (parent candidates, recipient candidates, monotonic floors, grant — `render.go:96-146`) change *without even moving the form digest* (`formForDigest` strips them, `render.go:249-262`). A digest over rendered bytes would churn per run and per submit; it cannot be a member of a build-identity vector that must survive a run immutably (§4 run-binding) and bind at a pre-build interface-lock (F63 event i).

**The split, therefore:**
- the **tool-schema digest** binds the mapping's *static schema surface* — everything about the tool's input contract that does not depend on a live form;
- the **mapping VERSION** (§3.3) binds the *dynamic-fill rules* — how a live form populates that surface (the R-rules).
Together they are the complete m-2 half of the relay-verb identity: surface + generator.

**Digest inputs (normative bytes, pinned):**

- **`relay.submit`** — the digest is over the **submit schema TEMPLATE**: the `SchemaFromForm` output with the two run-varying slots normalized — `properties.headers.properties = {}` (the form-fill slot, empty) and `properties.form_digest = {"type":"string"}` (the const pin removed; member and type retained). Everything else is the mapping's own static envelope (R-1, incl. the R-1.b advertised-open headers interior — review-r3). Canonical bytes:

  ```
  {"additionalProperties":false,"properties":{"body":{"type":"string"},"cc":{"description":"canonical JSON string - array of address strings","type":"string"},"dispatch_id":{"type":"string"},"form_digest":{"type":"string"},"headers":{"additionalProperties":true,"properties":{},"type":"object"},"to":{"type":"string"}},"required":["headers","form_digest"],"type":"object"}
  ```

- **`relay.project`** — the static schema, whole (hoisted from `mcp.go:342-349`):

  ```
  {"additionalProperties":false,"properties":{"view":{"enum":["default","audit","roster"],"type":"string"}},"type":"object"}
  ```

- **`relay.read`** — the static schema, whole (hoisted from `mcp.go:351-360`):

  ```
  {"additionalProperties":false,"properties":{"relay_id":{"type":"string"}},"required":["relay_id"],"type":"object"}
  ```

**Canonical encoding (normative):** digest = **SHA-256 over the canonical JSON encoding** of the schema document, rendered as **64 lowercase hex characters**. Canonical JSON = Go `encoding/json` marshaling of the schema value (object keys sorted lexicographically — the standard-library guarantee `fieldspec.CanonicalMarshal` already leans on, `canonical.go:12-18`), UTF-8, no insignificant whitespace, no trailing newline. No new canonicalization scheme is invented (Rail B; no home-grown primitives — SHA-256 here is labeling/identity, not a security claim).

**Digest scope:** the digest covers the tool's **input schema only**. The canonical *name* rides the vector beside it (not embedded); the human-facing *description* strings (`mcpTools`' honesty banner, `mcp.go:326-336`) are **excluded** — presentation may vary per frontend without moving call semantics; the honesty banner's own governance home is the m-9 catalog/m-3 surface, not this digest.

**Reference values @ the §3.2 pinned bytes** (derived, recomputable by anyone from the bytes above; **the authoritative values bind at the Master+VP interface-lock over the then-approved bytes** — if pair review or a consumer confirmation moves a byte above, these move with it):

| Canonical ID | tool-schema digest (SHA-256, reference) |
|---|---|
| `relay.submit` | `6bb7bbf46d8bf5d210cee410fbd0fa59106145425878c065adf0d54b05ace08e` *(rev3: the R-1.b headers-open amendment moved this — was `c0c7d82f…` at rev1/rev2)* |
| `relay.project` | `be5c41ec848bd7f6a7afd16af5acc56c65cf39bc113041941bb6747153bd582a` |
| `relay.read` | `a84645cb3f57ea1172661ddcc42e8a710f5a320ee3ed6c944f5e469026b3036e` |

### §3.3 The mapping VERSION + its mechanical verification (the F63 hook)

- **Token grammar (normative):** `m2-mapping-v<N>`, `<N>` a positive integer, strictly monotonic, no gaps skipped silently. First built version = `m2-mapping-v1`. (Deliberately the same declared-marker discipline as the FieldSpec registry's `s10-fieldspec-v8` lineage — one team pattern, no new flavor.)
- **Bump discipline:** ANY semantic change to the R-rules, P-rules, or re-render contract (§2.3/§2.4) — including a change that does not move the §3.2 template bytes (e.g. how `Options` project, a new `structuredDescription` arm, a P-rule tightening) — bumps `<N>` in the same reviewed changeset. A §3.2 template-byte change necessarily bumps too (both members move together through the lock).
- **Mechanical verification at the F63 release-binding — the LOCKED reference set (REVISED at review-r2 MR-5; supersedes rev1's shipped-suite fingerprint).** A version string alone is a declaration; F63 demands drift-without-version-change fail closed, and the amendment's two-event split (`:59,:87`) demands a **pre-build anchor with concrete expected bytes** — running a shipped implementation over a suite shipped beside it would be self-consistency, not comparison against the approved design semantics. Realization:
  - **The branch inventory (branch grain, machine-enumerable):** every normative behavior branch carries a stable branch ID — **41 branches (mechanically counted, review-r3 MR-7)**: `R-1.a–b`, `R-2.a–d`, `R-3.a–e`, `R-4.a–b`, `R-5.a` (schema generation); `P-0.a–b`, `P-1.a–c`, `P-2.a–b`, `P-3.a–c`, `P-6.a–c` (strict parse + payload); `V-0.a`, `V-1.a–e`, `V-2.a–b` (validation projection, §2.3.3); `RR-1.a–b`, `RR-2.a–b`, `RR-3.a–b` (re-render helpers) — enumerated in **Appendix A** and mirrored as an exported ordered constant in the module (divergence = build failure). Coverage binds at **branch** grain, both directions: `union(exercises) == inventory`; a branch without a vector fails, a vector citing an unknown branch fails, and mutually exclusive branches have their own vectors by construction. (P-4 and P-5 carry no branches of their own: P-4's Layer-2 projection *is* the V-branches; P-5's typed identities appear in every error record.)
  - **The LOCKED fingerprint reference set (Appendix A — exact, immutable, pre-build):** the ordered vector IDs (`S1, P1–P14, V1–V8, R1–R5` — 28 vectors), their **byte-exact inputs** (the reference form RF-1, the digest sentinel, raw argument byte strings — including deliberately malformed ones a JSON serializer cannot emit, e.g. duplicate members and trailing data), each vector's `exercises` list, the **canonical expected result records**, the serialization shape, and the resulting **expected fingerprint** are pinned in Appendix A of THIS document. They are covered by this doc's hash and bind at the interface-lock. **Immutability posture (clarified at review-r3):** from the interface-lock onward, **changing any Appendix-A byte is a mapping-version bump + re-lock** — the reference set is identity, not a test suite. Before the lock, corrections ride doc revisions through pair review (as this rev does), the expected fingerprint is recomputed with each such revision, and the first built version lands as `m2-mapping-v1` bound to the fingerprint the lock records.
  - **Scope — pure behavior only:** the reference set exercises the module's transport-free operations (`schema` = `SchemaFromForm`, `payload` = strict parse + `SubmitPayloadFromArguments`, `validate` = the §2.3.3 projection, `rr_detect`/`rr_result`/`rr_key`), so the F63 event can execute it hermetically. The F-rule choreography (refresh timing, wire calls, retries) is deliberately **not** fingerprint input — it needs a conductor and lives in the extensible parity suite (§4).
  - **The two events, honestly split:** **(i)** the Master+VP interface-lock records this doc's hash (covering Appendix A) + the expected fingerprint bound to the locked `m2-mapping-v<N>`; **(ii)** the F63 release-binding **executes the shipped artifact's** mapper/validator/helpers over the **locked Appendix-A inputs** and compares the actual canonical result records + fingerprint against the **lock-recorded expected bytes** — a mismatch without a version bump **fails closed** (the manifest does not serve). Both sides are never derived from T4 output: the expected side is fixed pre-build, the actual side is the shipped implementation.
  - **The extensible parity/regression suite is SEPARATE (§4):** the build lane may add vectors there freely — additions never move identity, because the fingerprint input is exactly the Appendix-A set and nothing else.
  - **Provenance + recomputability:** the Appendix-A expected records were computed from a design-side executable expectation of the R/P/V/RR rules (independent of any Go implementation; the serialization conventions are pinned in Appendix A); the pair reviewer can recompute every record and the fingerprint from the pinned inputs + rules. A T4 build that cannot reproduce a record has either a build bug or has found a design defect — the latter routes back here as a reviewed amendment, never a silent expected-bytes edit.
  - **Honesty bound (unchanged in kind, narrowed in grain):** the fingerprint mechanically binds every enumerated branch. The residual is behavior outside the branch inventory — a dimension nobody has named — guarded by the coverage test and pair review of inventory completeness; accepted and labeled, not claimed away.
- **The fingerprint is a verification datum bound to the version member, NOT a fourth vector field** — the vector stays exactly the ratified 4-tuple; this section is the owner-DESIGN realization of "mechanically verified" that F58 delegates.

### §3.4 Field applicability (normative, for m-10's gate and the catalog assembly)

- The `form→schema-mapping version` member is **present iff** the tool's schema is produced by this module — exactly the three relay verbs. For the five local tools the member is **ABSENT** (not empty-string, not `"none"`): an absent member is the honest encoding of "no such producer exists", and it makes a local tool claiming a mapping version (or a relay verb missing one) a *shape* violation at the serve gate, not a value comparison.
- Vector-member encodings m-10 may rely on: canonical name = the §3.1 table byte-exact; digest = 64 lowercase hex; mapping version = the §3.3 grammar. Set-equality over identity (F55) compares all present members byte-exact after alias normalization.

---

## §4 — Parity conformance vectors (amendment §5/§10: "native == MCP relay path")

**Vector shape (each vector is one row, JSON, in the module's `testdata/parity/`):**

```
{id, description,
 exercises: [<branch-IDs from the Appendix-A inventory, when the vector has a pure core>],
 form:      <rendered-form fixture + form_digest>          (submit vectors),
 tool:      <canonical ID, §3.1>,
 arguments: <raw argument bytes as the frontend receives them>,
 expect:    {outcome: mapped | mapping_error | schema_invalid,
             wire_call:   yes {verb, payload_bytes} | no,
             error:       <P-rule / disposition identity>   (when mapping_error | schema_invalid),
             re_render:   none | detected,
             refresh_key: <phase, tier>                     (when applicable)}}
```

**This suite is EXTENSIBLE and is NOT fingerprint input (review-r2 MR-5 split):** identity lives exclusively in the locked Appendix-A reference set; the build lane may add parity/regression vectors here freely without moving identity. Where a parity vector's pure core duplicates an Appendix-A vector, it references the Appendix-A ID rather than re-pinning bytes.

**The parity property (build-asserted, both frontends over the same vectors):** for every vector, the retained MCP frontend and the m-9 native tool produce (a) the **byte-identical conductor wire call** — same wire verb, same canonical payload bytes — or (b) the **same typed disposition with no wire call and no ticket** (mapping error or Layer-2 schema-invalid, P-4/P-5), and (c) the **same re-render disposition** (detection + refresh key per §2.4; surface shape is frontend-owned). Transport-level reconnect/retry (m-7's contract, incl. the content-hash replay safety noted at `mcp.go:182-184`) must be payload-transparent: a retry never changes the mapped bytes — asserted by vector harness instrumentation, not by trusting the comment.

**Coverage floor (normative minimum; the build lane may add, not subtract):**
1. plain submit — headers + body + to/dispatch_id (P-0 clean path);
2. structured fields — `row_array` + `object` + `address_list` round-trip as canonical-JSON strings (R-2);
3. enum projection + Layer-2 reject — options + default surface in the schema (R-3); an **out-of-enum submitted value fails the dispatch gate: `schema_invalid`, no wire call, no ticket** (P-4 Layer 2 — REVISED at review-r1 MR-1; rev0's reaches-the-wire claim withdrawn, it contradicted the ratified validated-before-authorization seam);
4. unknown header name → `schema_invalid` at the dispatch gate (the generated schema closes the headers interior), no call (P-4 Layer 2);
5. first-relay shape — no `PARENT_DISPATCH_ID` (absence passes through, P-0);
6. the CC fold — top-level `cc` only; equal duplicate; **conflicting duplicate → P-3 error, no call**;
7. unknown top-level member → P-1 error, no call; **case-variant member name (`CC`, `Body`) → P-1 error, no call** (exact-name binding);
8. reserved header key (`relay_id`, `TO`) → P-2 error, no call;
9. **duplicate members (P-6, each its own vector): duplicate top-level `body`; duplicate top-level `headers`; duplicate key inside `headers`; trailing data after the top-level value — each → P-6 error, no call**;
10. missing required member (`form_digest` absent) → `schema_invalid`, no call (P-4 Layer 2);
11. stale `form_digest` → conductor-bounce fixture → re-render detected (RR-1 true branch), normalized violation surfaced (RR-2), refresh keyed by declared PHASE/CEREMONY_TIER and by the fallback pair (RR-3, two vectors); plus a non-re-render rejection fixture (RR-1 false branch);
12. `project` — each view enum value + empty args;
13. `read` — present `relay_id`; missing `relay_id` → `schema_invalid`, no call, both frontends converge;
14. empty/null arguments on each verb;
15. **the PV volatile-freshness family (review-r2 MR-4; each a same-digest old/new-form fixture pair against a fake conductor):**
    - **PV-1** volatile option expansion (parent candidates grew): stale Layer-2 reject → F-1 refresh → fresh form accepts → the call proceeds (no persistent false reject);
    - **PV-2** volatile option contraction: stale Layer-2 pass → conductor rejects with its current typed violation (visible) → F-2 refresh (no silent stale acceptance, no stuck cache);
    - **PV-3** `grant` absent→present under an unchanged digest: unknown-name reject → F-1 refresh → present in the fresh form → the call proceeds;
    - **PV-4** recipient candidates change: top-level `to` is never schema-constrained — the call proceeds on both sides, the conductor decides (proves no false reject exists on either frontend for recipient drift);
    - **PV-5** monotonic floor raise: stale Layer-2 pass → conductor `monotonic-floor` violation (visible) → F-2 refresh.
    Each PV vector asserts BOTH directions: neither frontend remains on a false rejection (F-1) and neither silently retains a stale acceptance path (F-2), with the F-3 race honesty bound stated in the fixture. **(Review-r3 MR-6) PV-1 and PV-3 additionally carry a HOST-SIDE leg:** the MCP variant drives a host simulator that validates the *advertised* `inputSchema` before issuing `tools/call` (not `handleToolCall` directly), proving a conforming host cannot reject volatile drift pre-call (nothing volatile is enforcing per the amended R-1.b/R-3) and that every F-1/F-2/re-render refresh emits `notifications/tools/list_changed` so the host's advertised schema converges.

**The MCP parity boundary, narrowed exactly (review-r3 MR-6):** parity is defined over OUR two frontends' behavior — mapper, validator, choreography, emitted signals. A foreign MCP host's own validation behavior is outside that boundary; the advertised schema is engineered so that a **spec-conforming** host validating it cannot false-reject volatile drift (no volatile enforcing keyword exists), and a host with a stale *stable* schema self-heals through the digest/re-render path (the stale const rides in, the conductor bounces, the refresh signals). The residual — a host that both ignores `tools/list_changed` and hard-caches schemas indefinitely — is a host-conformance defect outside this contract, labeled here rather than claimed away; the §10 "native == MCP relay path" row is satisfied at our boundary with the host-side PV legs as evidence.

**Home + consumption:** the vectors and the harness half that runs the *mapping* live in the m-2 module; each frontend contributes a thin adapter (MCP: drive `handleToolCall` with a fake transport capturing wire calls; native: m-9's equivalent). The vectors are the shared artifact the §10 acceptance row cites; m-9 confirms consumability (§5).

---

## §5 — Consumer confirmations owed (per the dispatch + supplement)

| Consumer | Confirms |
|---|---|
| **m-9** (native-tool consumer) | The §2.2 API is sufficient for the native relay tool (schema surface, payload mapping, typed errors, re-render loop steps 1–2) with no FieldSpec semantics re-implemented on their side; the §4 vectors are consumable in their harness; the §3.1 alias posture works for their catalog assembly (their catalog carries my digests + version for the relay verbs per §3.4). **Additionally (rev1):** the P-4 Layer-2 gate composes with their validated-before-authorization lifecycle (their §3) — schema validation for the relay verbs may consume the offered `ValidateSubmitArguments` helper or be m-9-implemented against the same generated schema; either way the §4 dispositions bind (offered, not imposed — validation stays m-9-owned code per their contract). |
| **m-7** (transport beside) | The §2.1 seam: `internal/channel` neither absorbs nor strands the module (import boundary both directions); `DescribeTools`/`Call` remain the complete transport surface the re-render contract needs; retry/reconnect payload-transparency is theirs to guarantee, mine to assert in vectors. *(Reciprocally, per the supplement, m-2 owes a confirmation of m-7's transport contract when their bytes land — a separate relay on master's routing, not this doc.)* |
| **m-10** (identity-verification host) | The §3.2 canonical encoding, §3.3 version grammar + fingerprint procedure, and §3.4 applicability/absence rules are sufficient inputs for the F55 exact-set-equality gate and the F63 release-binding verification of the m-2 members; the §3.1 table is a usable normalization reference. |

---

## §6 — Build obligations (all RED-first; none extant @ `502e06c`)

1. **Module hoist** — create the module; move `SchemaFromForm`/`SubmitPayloadFromArguments`/re-render halves; `cmd/frank-mcp` becomes a consumer; behavior-preserving except the flagged P-1/P-2/P-3/P-6 deltas and the MCP-side Layer-2 gate (each lands with its own red test).
2. **Import-boundary test** — the §2.1 graph, both directions, fails on violation.
3. **P-rule enforcement** — strict pre-scan (duplicate members + trailing data, the `config_change.go:101-124` pattern) + exact-name binding + unknown-member reject; reserved-key reject; CC-conflict reject; typed error identities for P-1/P-2/P-3/P-6.
4. **The dispatch-gate validator** — `ValidateSubmitArguments` (+ project/read validators) implementing the §2.3.3 V-partition (stable rejects, volatile advisory); typed `schema_invalid` dispositions; the MCP frontend wires it before dialing.
5. **The freshness choreography** — F-1 refresh-before-reject and F-2 refresh-on-rejection in both frontends (MCP here, native = m-9's lane), with the PV fixture family (§4 item 15).
6. **F58 producers** — `ToolSchemaDigest`/`MappingVersion`/`ReferenceFingerprint` + the pinned template/schema byte constants; a test recomputing the §3.2 reference values from the pinned bytes.
7. **Branch inventory + coverage binding** — the exported branch-ID constant mirroring Appendix A; the `union(exercises) == inventory` test at branch grain, failing in both directions (§3.3).
8. **The Appendix-A runner** — the harness that executes the LOCKED reference set against the built module and emits canonical result records + the fingerprint; a build test asserting the shipped behavior reproduces the Appendix-A expected bytes exactly (the vectors themselves are NOT build-authored — they are locked here).
9. **Parity harness** — mapping+validation half in-module; frontend adapters in their lanes (MCP adapter here, native adapter = m-9); fake-conductor fixtures for the PV family; retry payload-transparency instrumentation.
10. **Release-binding hook** — the fingerprint recomputation entry point the F63 event invokes over the locked inputs (invocation/orchestration = m-10's event, the computation = this module).

---

## §7 — Rails + confusion-firewall trace

- **Rail A, per surface (revised at review-r3 — the review-r1 wording here asserted the superseded rev1 mirror guarantee, MR-9):** the top-level argument envelope → **CLOSED** (P-1/P-6: unknown members, case-variant names, duplicate members, trailing data — ignore-or-last-wins changes what the author believes was accepted, the silent-drop/silent-merge confusion class). The headers interior → **CLOSED at Layer 2 for digest-bound STABLE constraints under the §2.3.3 freshness contract; volatile state is ADVISORY at every client-side surface** (the V-partition; F-1 heals false rejects, F-2 heals stale accepts, F-3 states the residual race) — the mirror claim is **convergent-with-visible-divergence, never instantaneous equality**, and the published schema enforces only self-healing digest-bound keywords (the advertised-schema principle, §2.3.1). The **mapper itself stays pass-through** (P-4 Layer 1 — the helper never pre-empts the conductor). The applicability encoding → absence-is-honest (§3.4).
- **Rail B, cut by function:** no new canonicalization, no new hash discipline, no new version flavor — `CanonicalMarshal`-style canonical JSON, SHA-256 hex, and the registry's declared-marker versioning are each reused. The digests and fingerprint are **identity/honest-labeling mechanisms** (drift becomes visible and fails closed at a gate), not security primitives; nothing here claims tamper-resistance against a malicious build — that is F63's release-digest territory and, beyond it, out of the confusion-not-malice scope.
- **Threat model fit:** every mechanism above kills a confusion path — a stale schema silently accepted (R-5 + §2.4), a misspelled argument silently dropped (P-1), two CC sources silently merged (P-3), mapping drift silently shipped (§3.3), two frontends silently diverging (§4).

## §8 — Flagged-open (routed, not guessed)

- **(a)** Final module path token — build lane's, within §2.1 normative ownership/boundary.
- **(b)** Whether the MCP frontend renames its tools to `relay.*` or keeps aliases — m-9/m-10 lanes, against the §3.1 table.
- **(c)** The re-render detection predicate's text-keying fragility — accepted v1, labeled at §2.4; follows any future conductor structured-violation move in that (Step-4+) changeset.
- **(d)** The m-2→m-7 transport-contract confirmation — owed when m-7's bytes land (supplement), separate relay.

## §9 — Revision log

- **rev5** (2026-07-16, m-2.planner): the single `step3-mvp-m2-mapping-review-r5` blocker folded — **MR-11**: the line-3 header still self-identified as rev2 with a history truncated at rev2 (I bumped it at rev2 and never again — the same stale-locus class as MR-9/MR-10, this time at the primary metadata locus consumers read first). Fix: the header now names only the live revision and **defers to §9 as the authoritative history** (no per-rev summary to go stale), with the convention stated so every future revision moves it. Named deviation from the r5 bar's literal wording (owned in the relay): the bar said "identifies these bytes as rev4", but folding MR-11 mints new bytes — stamping them rev4 while the pair-reviewed rev4 hash `62efe963…` names different bytes would recreate the exact ambiguity class MR-11 corrects; these bytes are **rev5**. Metadata sweep: no other live current-revision marker exists (§9 entries are per-rev history; Appendix A carries no revision marker). **Appendix A byte-identical** — no identity byte moved.
- **rev4** (2026-07-16, m-2.planner): the single `step3-mvp-m2-mapping-review-r4` blocker folded — **MR-10**: two live normative loci (§2.3.3's vector line; §3.3's exact locked-set definition — part of the F63 pre-build identity, where a release-binding implementation following the stale sentence could have omitted V8 and defeated the MR-7 correction) still bound `V1–V7`/27 while Appendix A binds `V1–V8`/28. Both corrected to `V1–V8`/28; full-doc range/count sweep run (no other live locus; the §9 rev2 entry stays as history). **Appendix A byte-identical** — the A.5 array re-verified to recompute the unchanged expected fingerprint `306b3149…`; no identity change, no version-posture event.
- **rev3** (2026-07-16, m-2.planner): the four `step3-mvp-m2-mapping-review-r3` blockers folded.
  **MR-6** — the advertised MCP `inputSchema` (`mcp.go:325-360`) still carried volatile constraints as enforcing keywords, so a schema-validating host could reject volatile drift BEFORE any call existed for F-1 to heal; and F-1/F-2 refreshes did not bind to `tools/list_changed`. Fix: the **advertised-schema principle** (§2.3.1 — enforcing keywords carry only digest-bound self-healing constraints; annotations may carry volatile guidance); **R-1.b** headers interior advertised-open; **R-3 amended** (volatile options → description annotation, never `enum`; `default` kept as a non-enforcing annotation); §2.4 step 3 signals on EVERY refresh; PV-1/PV-3 gain host-side legs; the MCP parity boundary narrowed exactly (our frontends' behavior; conforming-host safety engineered; the ignore-notifications-and-hard-cache residual labeled). The submit TEMPLATE digest moved accordingly (`c0c7d82f…` → `6bb7bbf4…`).
  **MR-7** — the "40-ID" inventory claim was FALSE (mechanical count: 39; the generator asserted set-coverage, never the count — owned), and required-`headers` was genuinely unbound (V4 tested only absent `form_digest`; a validator dropping the headers-required check would have reproduced every locked byte). Fix: **V-1.e** (absent `headers`) + vector **V8**; R-1.b joins the inventory; the corrected mechanically-counted inventory is **41**; fingerprint recomputed.
  **MR-8** — RF-1's pinned bytes were the bare field map, not an executable `fieldspec.Form` encoding (`{"fields":{…}}`, `fieldspec.go:11-13`) — a direct unmarshal would populate nothing and no runner transform was bound. Fix: A.3 re-pinned as the canonical `Form` encoding; A.1 binds the runner rule (`"form":"RF-1"` resolves to the A.3 bytes unmarshaled as `fieldspec.Form`); the "DescribeTools response shape" over-claim corrected (RF-1 is the `submit_schema` MEMBER's shape, not the response).
  **MR-9** — §7's Rail-A line still asserted rev1's superseded digest-pinned-mirror/re-render-resync guarantee, and my rev2 relay's "echo sweep clean" claim was therefore FALSE (owned — the sweep greped for the wrong tokens and §7's paraphrase escaped it). Fix: §7 rewritten to the V-partition + F-rules convergence claim; this rev's sweep checks paraphrase sites, not just token matches.
  Appendix A regenerated end-to-end: 41-branch inventory, 28 vectors, expected fingerprint **`306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac`** (pre-lock revision per the §3.3 posture — no version bump owed before the first lock).
- **rev2** (2026-07-16, m-2.planner): the two `step3-mvp-m2-mapping-review-r2` blockers folded, each verified at source before folding.
  **MR-4** — rev1's Layer-2 "current generated schema" gate + digest-pinned-mirror claim broke on digest-exempt volatile state (verified: `render.go:109-147` volatile marking; `formForDigest` strips before hashing `render.go:249-262`; `render_test.go:90-137` equal digests across recipient/monotonic drift AND grant absent→present; `validate.go:23-26,55-65` contraction surfaces as current `enum`/`seat-scope`/`monotonic-floor`, never `re-render`, so RR-1 cannot heal a stale cache; `mcp.go:192-225` refresh only on initial absence + post-re-render). Fix: **§2.3.3** — the stable/volatile constraint partition (V-rules: Layer 2 rejects digest-bound constraints only; volatile options advisory, conductor answers with visible typed violations) + the freshness contract (F-1 refresh-before-reject; F-2 refresh-on-any-rejection; F-3 the exact race-honesty claim — convergent with visible divergence, no silent path, instantaneous equality NOT claimed); the PV same-digest old/new-form vector family (§4 item 15); P-4 Layers 2/3 and the §7 trace re-stated to the narrowed claim; no conductor byte (the volatile flags already ride `DescribeTools`, `fieldspec.go:15-21`).
  **MR-5** — rev1 had no pre-build fingerprint anchor: the suite was a T4 obligation, extensible additions moved the hash, coverage was rule-grain, and shipped-suite-vs-shipped-impl was self-consistency. Fix: **Appendix A** — the LOCKED reference set with exact bytes authored NOW (27 ordered vectors, 40 branch-grain inventory IDs, byte-exact inputs incl. serializer-unproducible malformed ones, canonical expected result records computed from a design-side executable expectation, pinned serialization conventions, expected fingerprint `9f80a913ebedf234cc23ccfcdface12c4b3c916d7a748b485dd4175d43351fb9`); the extensible parity suite split off as non-identity; the two F63 events cleanly separated (lock records expected bytes pre-build; release-binding compares the shipped artifact's actual records against the locked expectation); Appendix-A bytes immutable — any change = version bump + re-lock.
- **rev1** (2026-07-16, m-2.planner): the three `step3-mvp-m2-mapping-review-r1` blockers folded, each verified at source before folding.
  **MR-1** — rev0's P-4 "interior openness" + vector 3 contradicted the ratified validated-before-authorization seam (verified: m-9 design §3 "malformed … never reaches authorization"; m-10 §D.2 check (3) zero-dispatch/no-ticket; amendment `:61,:95`) and rev0's own closed generated schema. Fix: P-4 rewritten as the three-layer split (non-authoritative mapper / normative frontend dispatch gate with the shared `ValidateSubmitArguments` helper / conductor as sole acceptance authority via the digest-pinned mirror); vector 3's reaches-the-wire claim withdrawn; `schema_invalid` disposition added (P-5); rev0's fork-of-acceptance objection retracted with grounds.
  **MR-2** — rev0's fingerprint hashed one schema render + one clean payload, leaving P-1/P-2/P-3/P-5, the P-4 layers, and all re-render behavior unbound (a shipped change to any of them preserved both version and fingerprint). Fix: machine-enumerable rule inventory (R-1…R-5, P-0…P-6, RR-1…RR-3, exported constant), the ordered `exercises`-annotated reference vector suite with canonical result records (shared with §4 — one artifact), both-direction coverage binding, fingerprint = SHA-256 over the ordered `(vector-id, result-record)` pairs recomputed from the shipped artifact; residual narrowed to behavior outside the inventory and labeled.
  **MR-3** — `DisallowUnknownFields` alone left Go's documented duplicate-member last-wins/merge (and the case-insensitive field fallback) silently selecting author values — the same class P-3 closes. Fix: P-6 strict object decoding (duplicate-member + trailing-data reject at every closed layer, byte-exact per layer, the `internal/store/config_change.go:69-124` precedent reused) + P-1 exact-name binding (case-variants are unknown members); vectors + fingerprint coverage added.
- **rev0** (2026-07-16, m-2.planner): initial contract for pair review. Grounded @ `502e06c`; amendment r7 verified byte-exact @ `2f75f2a1…` before authoring.

---

## Appendix A — the LOCKED fingerprint reference set (identity, not a test suite)

**From the interface-lock onward, any byte change here is a mapping-version bump + re-lock; pre-lock corrections ride doc revisions through pair review (§3.3).**
The expected fingerprint binds at the Master+VP interface-lock; the F63 release-binding executes the shipped artifact over the A.4 inputs and compares its actual canonical result records + fingerprint against A.5/A.6.

### A.1 Serialization + runner conventions (normative)

- Canonical JSON for map-typed values = Go `encoding/json` marshaling: object keys sorted lexically, no insignificant whitespace, UTF-8, no trailing newline (the §3.2 encoding).
- **Mapped payload bytes follow Go STRUCT marshal order** (not sorted): `fieldspec.SubmitPayload` embeds `record.Record` (`validate.go:236-239`), so the payload members appear as `envelope, headers, body, checksum, form_digest` with `x_fields`/`intake_id`/`to` omitted when empty and `checksum:""`/`schema_version:0` always present; `record.Envelope` members appear as `relay_id, dispatch_id, from, to?, role, delivery_state, intake_id?, schema_version` (`record.go:16-33`). Header maps serialize key-sorted.
- **Runner input binding (review-r3 MR-8):** a vector's `"form": "RF-1"` resolves to the A.3 canonical bytes **unmarshaled directly as `fieldspec.Form`** (`{"fields": {...}}`, `fieldspec.go:11-13`) — no other transform exists; `"digest"` / `"args"` / `"outcome"` values are the byte-exact inputs to the named operation.
- Each vector's result record is a map-typed value (canonical, key-sorted). The **fingerprint input** is the canonical JSON array of ordered `[vector_id, result_record]` pairs (A.5, exactly that order); the **fingerprint** is its SHA-256, lowercase hex.
- Validation dispositions use the class vocabulary `{required, digest-mismatch, unknown-header, enum}` (§2.3.3 V-rules); parse errors use the typed identities `P-1.a/b/c, P-2.a/b, P-3.c, P-6.a/b/c` (§2.3.2).

### A.2 The branch inventory (41 IDs — mechanically counted; the exported module constant mirrors this list, in this order)

`R-1.a` envelope shape · `R-1.b` headers interior advertised-open · `R-2.a` row_array hint · `R-2.b` object hint · `R-2.c` address_list hint · `R-2.d` plain carrier · `R-3.a` stable options→enum · `R-3.b` no options · `R-3.c` default→default · `R-3.d` no default · `R-3.e` volatile options→description annotation (never `enum`) · `R-4.a` system-name skip (case-insensitive) · `R-4.b` envelope-name skip (exact) · `R-5.a` digest const ·
`P-0.a` verbatim map · `P-0.b` absent optionals stay absent · `P-1.a` unknown member · `P-1.b` case-variant member · `P-1.c` member type mismatch · `P-2.a` reserved system key · `P-2.b` reserved envelope key · `P-3.a` cc fold · `P-3.b` equal duplicate fold · `P-3.c` cc conflict · `P-6.a` duplicate top-level member · `P-6.b` duplicate headers key · `P-6.c` trailing data ·
`V-0.a` valid · `V-1.a` stable out-of-enum · `V-1.b` unknown header name · `V-1.c` missing required form_digest · `V-1.d` digest mismatch · `V-1.e` missing required headers · `V-2.a` volatile enum advisory pass · `V-2.b` volatile bool/monotonic advisory pass ·
`RR-1.a` re-render detected · `RR-1.b` not detected · `RR-2.a` normalized result bytes · `RR-2.b` relay/intake passthrough · `RR-3.a` declared refresh key · `RR-3.b` fallback refresh key

### A.3 The reference form RF-1 (the canonical `fieldspec.Form` encoding) + digest sentinel `ref-digest-1`

```json
{"fields":{"AUTHORITY":{"default":"report-only","options":["read-only","report-only"],"type":"enum"},"CC":{"type":"address_list"},"CEREMONY_TIER":{"options":["small","medium","large"],"type":"enum"},"CONTEXT":{"type":"object"},"DISPATCH_ID":{"type":"id_ref"},"From":{"type":"text"},"HUMAN_GATE_REQUIRED":{"conductor_volatile":true,"digest_exempt":true,"options":["no"],"type":"bool"},"NOTIFY":{"type":"address_list"},"PARENT_DISPATCH_ID":{"conductor_volatile":true,"default":"parent-a","digest_exempt":true,"options":["parent-a"],"type":"id_ref"},"PHASE":{"options":["SITREP","PLAN"],"type":"enum"},"SCOPE_DIFF":{"type":"row_array"},"SUBJECT":{"type":"text"},"TITLE":{"type":"text"},"TO":{"type":"address_list"}}}
```

RF-1 is a synthetic fixture in the shape of the `submit_schema` MEMBER of a `DescribeTools` response (`channel/server.go:92-97` — the full response wraps it in additional members), not a registry claim: `HUMAN_GATE_REQUIRED` and `PARENT_DISPATCH_ID` carry the volatile flags (`conductor_volatile`/`digest_exempt`, `fieldspec.go:15-21`) to exercise V-2 and R-3.e; `From`/`TO`/`CC`/`DISPATCH_ID` exercise the R-4 skip set.

### A.4 The vectors (28, ordered; `[id, op, exercises, inputs]`; `args`/`outcome` values are byte-exact raw inputs)

```json
["S1","schema",["R-1.a","R-1.b","R-2.a","R-2.b","R-2.c","R-2.d","R-3.a","R-3.b","R-3.c","R-3.d","R-3.e","R-4.a","R-4.b","R-5.a"],{"digest":"ref-digest-1","form":"RF-1"}]
["P1","payload",["P-0.a"],{"args":"{\"headers\":{\"AUTHORITY\":\"report-only\",\"PHASE\":\"SITREP\",\"SUBJECT\":\"s\"},\"to\":\"m-2.implementer\",\"dispatch_id\":\"d-1\",\"body\":\"b\",\"form_digest\":\"ref-digest-1\"}"}]
["P2","payload",["P-0.b"],{"args":"{\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P3","payload",["P-1.a"],{"args":"{\"headres\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P4","payload",["P-1.b"],{"args":"{\"Body\":\"b\",\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P5","payload",["P-1.c"],{"args":"{\"headers\":{\"SUBJECT\":123},\"form_digest\":\"ref-digest-1\"}"}]
["P6","payload",["P-2.a"],{"args":"{\"headers\":{\"relay_id\":\"r-9\"},\"form_digest\":\"ref-digest-1\"}"}]
["P7","payload",["P-2.a"],{"args":"{\"headers\":{\"From\":\"x\"},\"form_digest\":\"ref-digest-1\"}"}]
["P8","payload",["P-2.b"],{"args":"{\"headers\":{\"TO\":\"x\"},\"form_digest\":\"ref-digest-1\"}"}]
["P9","payload",["P-3.a"],{"args":"{\"cc\":\"m-7.planner\",\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P10","payload",["P-3.b"],{"args":"{\"cc\":\"m-7.planner\",\"headers\":{\"CC\":\"m-7.planner\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P11","payload",["P-3.c"],{"args":"{\"cc\":\"m-7.planner\",\"headers\":{\"CC\":\"m-9.planner\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P12","payload",["P-6.a"],{"args":"{\"body\":\"b1\",\"body\":\"b2\",\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["P13","payload",["P-6.b"],{"args":"{\"headers\":{\"SUBJECT\":\"s1\",\"SUBJECT\":\"s2\"},\"form_digest\":\"ref-digest-1\"}"}]
["P14","payload",["P-6.c"],{"args":"{\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}{}"}]
["V1","validate",["V-0.a"],{"args":"{\"headers\":{\"AUTHORITY\":\"report-only\",\"PHASE\":\"SITREP\",\"SUBJECT\":\"s\"},\"to\":\"m-2.implementer\",\"dispatch_id\":\"d-1\",\"body\":\"b\",\"form_digest\":\"ref-digest-1\"}","digest":"ref-digest-1","form":"RF-1"}]
["V2","validate",["V-1.a"],{"args":"{\"headers\":{\"AUTHORITY\":\"design-only\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}","digest":"ref-digest-1","form":"RF-1"}]
["V3","validate",["V-1.b"],{"args":"{\"headers\":{\"NOPE\":\"x\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}","digest":"ref-digest-1","form":"RF-1"}]
["V4","validate",["V-1.c"],{"args":"{\"headers\":{\"SUBJECT\":\"s\"}}","digest":"ref-digest-1","form":"RF-1"}]
["V5","validate",["V-1.d"],{"args":"{\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"other-digest\"}","digest":"ref-digest-1","form":"RF-1"}]
["V6","validate",["V-2.a"],{"args":"{\"headers\":{\"PARENT_DISPATCH_ID\":\"parent-zz\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}","digest":"ref-digest-1","form":"RF-1"}]
["V7","validate",["V-2.b"],{"args":"{\"headers\":{\"HUMAN_GATE_REQUIRED\":\"yes\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}","digest":"ref-digest-1","form":"RF-1"}]
["V8","validate",["V-1.e"],{"args":"{\"form_digest\":\"ref-digest-1\"}","digest":"ref-digest-1","form":"RF-1"}]
["R1","rr_detect",["RR-1.a"],{"outcome":"{\"state\":\"rejected\",\"detail\":\"violation form_digest: re-render - stale form digest\"}"}]
["R2","rr_detect",["RR-1.b"],{"outcome":"{\"state\":\"rejected\",\"detail\":\"violation AUTHORITY: enum - unknown AUTHORITY\"}"}]
["R3","rr_result",["RR-2.a","RR-2.b"],{"outcome":"{\"state\":\"rejected\",\"relay_id\":\"r-1\",\"intake_id\":\"i-1\",\"detail\":\"violation form_digest: re-render - stale form digest\"}"}]
["R4","rr_key",["RR-3.a"],{"args":"{\"headers\":{\"CEREMONY_TIER\":\"large\",\"PHASE\":\"PLAN\",\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
["R5","rr_key",["RR-3.b"],{"args":"{\"headers\":{\"SUBJECT\":\"s\"},\"form_digest\":\"ref-digest-1\"}"}]
```

### A.5 The expected result records (ordered; the fingerprint input is the canonical array of exactly these pairs)

```json
["S1",{"bytes":"{\"additionalProperties\":false,\"properties\":{\"body\":{\"type\":\"string\"},\"cc\":{\"description\":\"canonical JSON string - array of address strings\",\"type\":\"string\"},\"dispatch_id\":{\"type\":\"string\"},\"form_digest\":{\"const\":\"ref-digest-1\",\"type\":\"string\"},\"headers\":{\"additionalProperties\":true,\"properties\":{\"AUTHORITY\":{\"default\":\"report-only\",\"enum\":[\"read-only\",\"report-only\"],\"type\":\"string\"},\"CEREMONY_TIER\":{\"enum\":[\"small\",\"medium\",\"large\"],\"type\":\"string\"},\"CONTEXT\":{\"description\":\"canonical JSON string - object\",\"type\":\"string\"},\"HUMAN_GATE_REQUIRED\":{\"description\":\"live options (conductor-validated; may change without a form_digest change): [\\\"no\\\"]\",\"type\":\"string\"},\"NOTIFY\":{\"description\":\"canonical JSON string - array of address strings\",\"type\":\"string\"},\"PARENT_DISPATCH_ID\":{\"default\":\"parent-a\",\"description\":\"live options (conductor-validated; may change without a form_digest change): [\\\"parent-a\\\"]\",\"type\":\"string\"},\"PHASE\":{\"enum\":[\"SITREP\",\"PLAN\"],\"type\":\"string\"},\"SCOPE_DIFF\":{\"description\":\"canonical JSON string - array of row objects\",\"type\":\"string\"},\"SUBJECT\":{\"type\":\"string\"},\"TITLE\":{\"type\":\"string\"}},\"type\":\"object\"},\"to\":{\"type\":\"string\"}},\"required\":[\"headers\",\"form_digest\"],\"type\":\"object\"}","outcome":"schema"}]
["P1",{"outcome":"mapped","payload":"{\"envelope\":{\"relay_id\":\"\",\"dispatch_id\":\"d-1\",\"from\":\"\",\"to\":\"m-2.implementer\",\"role\":\"\",\"delivery_state\":\"\",\"schema_version\":0},\"headers\":{\"AUTHORITY\":\"report-only\",\"PHASE\":\"SITREP\",\"SUBJECT\":\"s\"},\"body\":\"b\",\"checksum\":\"\",\"form_digest\":\"ref-digest-1\"}"}]
["P2",{"outcome":"mapped","payload":"{\"envelope\":{\"relay_id\":\"\",\"dispatch_id\":\"\",\"from\":\"\",\"role\":\"\",\"delivery_state\":\"\",\"schema_version\":0},\"headers\":{\"SUBJECT\":\"s\"},\"body\":\"\",\"checksum\":\"\",\"form_digest\":\"ref-digest-1\"}"}]
["P3",{"error":"P-1.a","outcome":"mapping_error"}]
["P4",{"error":"P-1.b","outcome":"mapping_error"}]
["P5",{"error":"P-1.c","outcome":"mapping_error"}]
["P6",{"error":"P-2.a","outcome":"mapping_error"}]
["P7",{"error":"P-2.a","outcome":"mapping_error"}]
["P8",{"error":"P-2.b","outcome":"mapping_error"}]
["P9",{"outcome":"mapped","payload":"{\"envelope\":{\"relay_id\":\"\",\"dispatch_id\":\"\",\"from\":\"\",\"role\":\"\",\"delivery_state\":\"\",\"schema_version\":0},\"headers\":{\"CC\":\"m-7.planner\",\"SUBJECT\":\"s\"},\"body\":\"\",\"checksum\":\"\",\"form_digest\":\"ref-digest-1\"}"}]
["P10",{"outcome":"mapped","payload":"{\"envelope\":{\"relay_id\":\"\",\"dispatch_id\":\"\",\"from\":\"\",\"role\":\"\",\"delivery_state\":\"\",\"schema_version\":0},\"headers\":{\"CC\":\"m-7.planner\",\"SUBJECT\":\"s\"},\"body\":\"\",\"checksum\":\"\",\"form_digest\":\"ref-digest-1\"}"}]
["P11",{"error":"P-3.c","outcome":"mapping_error"}]
["P12",{"error":"P-6.a","outcome":"mapping_error"}]
["P13",{"error":"P-6.b","outcome":"mapping_error"}]
["P14",{"error":"P-6.c","outcome":"mapping_error"}]
["V1",{"outcome":"valid"}]
["V2",{"dispositions":[{"class":"enum","member":"AUTHORITY"}],"outcome":"schema_invalid"}]
["V3",{"dispositions":[{"class":"unknown-header","member":"NOPE"}],"outcome":"schema_invalid"}]
["V4",{"dispositions":[{"class":"required","member":"form_digest"}],"outcome":"schema_invalid"}]
["V5",{"dispositions":[{"class":"digest-mismatch","member":"form_digest"}],"outcome":"schema_invalid"}]
["V6",{"outcome":"valid"}]
["V7",{"outcome":"valid"}]
["V8",{"dispositions":[{"class":"required","member":"headers"}],"outcome":"schema_invalid"}]
["R1",{"detected":true,"outcome":"re_render_detect"}]
["R2",{"detected":false,"outcome":"re_render_detect"}]
["R3",{"bytes":"{\"intake_id\":\"i-1\",\"relay_id\":\"r-1\",\"state\":\"rejected\",\"violations\":[{\"class\":\"re-render\",\"field\":\"form_digest\",\"hint\":\"form refreshed - re-read the submit tool schema and re-submit\"}]}","outcome":"re_render_result"}]
["R4",{"outcome":"refresh_key","phase":"PLAN","tier":"large"}]
["R5",{"outcome":"refresh_key","phase":"SITREP","tier":"medium"}]
```

### A.6 The expected fingerprint

```
306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac
```
