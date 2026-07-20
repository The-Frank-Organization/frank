# m-3 Step-3 MVP stage-1 contract — provider-egress policy delta + E0 app-event schema + F62 E3 observation record & applicability evaluator (r4)

**DESIGN_DOC_ID:** `step3-mvp-design-m3-egress-e0-e3`
**Owner:** m-3 (Observation & Evidence) — sole author; m-3.implementer pair-reviews the final bytes.
**Dispatch:** `step3-mvp-design-m3/DESIGN-orchestrator-planner-20260716-041700` + supplement `…-043510` (F68).
**Governing bytes:** `master/STEP-3-MVP-AMENDMENT.md` r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d` (verified this session) — §3, §4 (F58/F63), §7, §10; the reframe packet `2d240eb6…` as amended.
**Status:** r4 — the cancellation-phase owner delta folded per master routing `step3-mvp-design-m3/RECONCILE-…-033727` (requirement + rationale m-9's `212600`; the token mine; r3 was pair-APPROVED byte-bound at `70838f83…`, `…-030100` — this fold reopens pair review per that approval's own byte rule); pending fresh m-3.implementer review, then the scoped rebinds on master's routing. Consumer confirmations + the m-7 F68 consume gate CLOSE, not authoring. No lock is claimed or claimable here (Master+VP interface-lock is the gate).

---

## §0 — Claim ceiling first (what this contract does NOT claim)

This domain's law applied to itself: name the vantage honestly before naming the mechanism.

1. **The conductor cannot see the provider send.** Provider traffic bypasses the conductor by ratified topology. Everything the conductor ever holds about a provider attempt is a **worker-carried report** — E0, `self_reported`, never gate-satisfying, never promoting. The word **"attestation" is avoided** throughout (ratified §3): the conductor's channel stamp establishes `FROM = m-9-seat` and that *the body claims m-8 origin*; it establishes **nothing** about whether m-8 authored those bytes (m-8 is not a seat).
2. **This policy governs the designated m-9→m-8 provider attempt ONLY.** Local tool effects — **including network egress originating from `bash`** — are NOT governed in the MVP (ratified §3, operator-accepted). No sentence in this contract may be read as a claim over bash-originated traffic.
3. **No content-safety scanning of prompt/response bytes is performed in the MVP.** The away-email scanner was content-based; this class is **binding/structure-based** (§1.3). Prompt bytes are opaque to this policy. Content policy over provider payloads is Step-4+. **Consequently P3 (§1.3) proves only reserved-NAME exclusion: values of non-reserved headers and all body bytes remain opaque and are NOT proven credential-free** (F2).
4. **Enforcement is app-side and therefore itself E0-from-the-conductor's-vantage.** m-8 enforces this policy; proof that it enforced it comes from **instrumented tests + the external E3 observer** (§4), never from an app statement written into a conductor record.
5. **Observer provenance is convention-bound, not cryptographic (F1).** `observer_id` is bound to the recorder boundary the record was obtained from (§3.6) — an observer-owned output root / operator handoff, per the confusion-not-malice threat model. **No signature or cryptographic observer authentication is claimed in the MVP.**
6. **The E0–E3 ladder only.** E4 is out of MVP scope. A historical E3 does not survive any bound-identity change (§3.4).

## §1 — Artifact 1: the provider-request egress-policy delta

### §1.1 The new class, and what it does not inherit
**`egress_class = provider-request`** — a NEW policy class, disjoint from the away-email local-outbox class. Explicit non-inheritance (the held-lane floors + VP guardrail, carried):
- **No model-name confidentiality.** `gpt-`/`claude-`/`llama`/`model_name=` bytes (`frank/internal/egress/rules.go:35-40`) are ordinary, legal provider-request content. That classifier is never invoked on this path.
- **No `egress_blocked` park, no outbox scan, no away-email denial vocabulary.** A denied provider attempt is a typed app-side outcome (§1.4), not a parked outbox item. The dormant `internal/egress` package is untouched and un-imported by this path.
- **No fourth relay `delivery_state`.** The relay states remain exactly `{accepted, rejected, held}` (`frank/internal/record/record.go:11-13`). A worker SITREP *reporting* a denial is an ordinary relay (normally `accepted`) — the denial lives in the report content, not the delivery state.

### §1.2 Placement (consumed, not owned)
m-8 enforces this policy at **`freeze → authorize → attach → send`** (ratified §1, preserved): the **authorize** step evaluates §1.3 over the **frozen request core** (method + canonical endpoint + non-auth headers + body). Deny ⇒ **zero send** and **the secret resolver is never invoked** — credential attach is strictly post-authorize, so on the denied path no secret byte is ever resolved, read, or logged (this strong denied-path property follows from the ratified placement contract, NOT from P3; F2). m-8 owns freeze/attach/send mechanics and the enforcement code; m-3 owns only the policy predicates, deny semantics, and policy bytes below. Upstream of this policy, m-10's epoch/lease fence (F59/F64) has already rejected stale-generation attempts — epoch fencing is **not** re-stated as a policy predicate (single owner; no duplicate enforcement surface).

### §1.3 The policy predicates (ALL must hold; any failure ⇒ deny)
Evaluated over the frozen core + the run-manifest facts presented by the enforcement host:

| # | Predicate | Deny token(s) |
|---|---|---|
| P0 | **Policy availability + validity (fail-closed floor):** the policy bytes are present, are exactly their RFC 8785 canonical serialization (§1.5 — non-canonical stored bytes ARE malformed), satisfy every §1.5 validity rule **including the denied-header membership invariant** (R2-F1: `denied_header_names` ⊇ the mandatory four ∪ the pinned lane's declared auth-header name, lowercased — verified at run freeze and re-verified by m-8 at load against the lane fact), and hash to the manifest's `policy_digest` | `policy-unavailable` (absent / malformed / non-canonical / any validity rule fails) · `policy-digest-mismatch` |
| PS | **Core structure:** the frozen core presents a method, a §1.6-valid canonical endpoint, and a body m-8 presents as the pinned lane's single designated `LLMRequest` encoding. Policy checks the envelope facts m-8 presents — it does **not** re-parse the body (§0.3) | `malformed-core` |
| P1 | **Lane pin:** the attempt's `provider_lane_id` equals the run manifest's pinned lane (exactly one lane in the MVP), and the lane fact presented carries a §1.6-valid declared endpoint | `lane-mismatch` (pin fails) · `lane-endpoint-invalid` (the lane's declared endpoint is not §1.6-valid) |
| P2 | **Endpoint binding:** the frozen core's canonical endpoint (§1.6) is byte-equal to the pinned lane's declared canonical endpoint AND is a member of the policy's `endpoint_allowlist` — **byte equality of canonical strings**, never substring/prefix match. The lane names the endpoint; a credential may only authorize it, never supply or alter it | `endpoint-mismatch` · `endpoint-not-allowlisted` |
| P3 | **Reserved-header-name exclusion (F2 — named for exactly what it proves):** no header in the frozen core's non-auth header set has a canonical (lowercased) name in the policy's `denied_header_names` set. This is a **NAME check only**: it keeps reserved auth-header names out of the core so the post-authorize attach is the sole auth-bearing transform. It does **not** inspect header values or body bytes and does **not** prove the core credential-free (§0.3). The lane's auth-header name is IN the set by the P0 membership invariant, so its coverage is digest-bound | `reserved-auth-header-in-core` |
| P4 | **Method:** the frozen core's method equals the pinned lane's catalog-declared method — **the lane catalog fact (m-8-produced) is the sole method authority; the policy document carries no method data** (R2-F1: `method_by_lane` removed as a semantically dead field) | `method-mismatch` |

### §1.3a Deterministic evaluation order + single-reason rule (F3)
The deny-token enum is **closed and totally ordered**; evaluation proceeds in exactly this order and **stops at the first failure — that token is THE deny reason**:

```
1 policy-unavailable → 2 policy-digest-mismatch → 3 malformed-core → 4 lane-mismatch
→ 5 lane-endpoint-invalid → 6 endpoint-mismatch → 7 endpoint-not-allowlisted
→ 8 method-mismatch → 9 reserved-auth-header-in-core
```

Two conforming implementations MUST return the same token for the same input. **Multi-failure vector (normative example):** a frozen core with the wrong lane pin AND a non-allowlisted endpoint AND an `authorization` header denies with `lane-mismatch` (position 4 — first failure in order), never any later token.

### §1.4 Deny semantics
- Deny ⇒ **zero provider network send, zero provider-wire event, zero secret-resolver invocation** (ratified floors).
- m-8 returns a typed `egress_denied{deny_reason}` to m-9 (exactly one token, chosen per §1.3a); m-10 records the attempt outcome `denied` in its app-state store; the worker MAY report it at E0 (§2). No auto-retry (one attempt per invocation, §2a); a user-requested retry is a NEW `attempt_id`.
- An unrecognized deny token at any consumer is a schema violation, not a new category.

### §1.5 Policy bytes + `policy_digest` (R2-F1 — one named algorithm, one byte form)
The policy is a canonical document (schema `m3.egress_policy.v1`):

```json
{
  "schema": "m3.egress_policy.v1",
  "egress_class": "provider-request",
  "pinned_lane": "<provider_lane_id>",
  "endpoint_allowlist": ["<§1.6-canonical-url>", "..."],
  "denied_header_names": ["authorization", "cookie", "proxy-authorization", "x-api-key", "<lane-auth-header>"]
}
```

**Canonical encoding = RFC 8785 (JCS)** — the same single algorithm the adjacent m-10 stage-1 contract pins for every digested byte sequence (`m-10 …-ipc-manifest-seam-contract.md` §A.2); no home-grown serialization rules. JCS settles string escaping, member ordering, and whitespace; **the stored artifact is exactly the JCS output bytes — no terminal LF, no BOM** (JCS output contains no newline). Digest = **lowercase-hex SHA-256 over exactly those stored bytes.**

**Validity preconditions (checked as P0, distinct from encoding):**
- Every string is **NFC-normalized** (an explicit validity rule — JCS does not normalize; non-NFC ⇒ malformed).
- **Closed schema:** duplicate object keys, unknown fields, wrong-typed or out-of-enum values ⇒ malformed. **`method_by_lane` does not exist in v1** — a policy carrying it is malformed (no semantically inert digest field may exist).
- **Set-valued arrays** (`endpoint_allowlist`, `denied_header_names`) MUST be lexicographically sorted by Unicode code point and duplicate-free; `denied_header_names` entries MUST be lowercase; `endpoint_allowlist` entries MUST be §1.6-valid. Violation ⇒ malformed.
- **Membership invariant (R2-F1):** `denied_header_names` ⊇ `{authorization, cookie, proxy-authorization, x-api-key}` ∪ `{the pinned lane's declared auth-header name, lowercased}` — so the digest covers P3's whole effective set; a lane change that changes the auth-header name invalidates the policy (P0), never silently widens P3.

One logical policy therefore has exactly one byte string and one digest. m-3 is the digest's producer (ratified §3). The run manifest carries `policy_digest` at run freeze (m-10 writes it; m-3 produces it), so policy identity is run-bound: changing the policy ⇒ a new run, and every prior E3 bound to the old digest goes non-applicable (§3.4). m-8 verifies the digest of its loaded policy bytes against the manifest's before serving any authorize decision.

**Normative example — exact canonical bytes (255 bytes) + digest:**
```
{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}
```
`policy_digest = ca364710764c3fb5fa6ca0f2faa6795e6aa6f49d3d55877ca4def10d540a3030` (computed over exactly the line above, no trailing newline).

### §1.6 Canonical endpoint form (R2-F1 — a genuinely unique restricted grammar, m-3-defined, m-8-confirmed)
Rather than adopting a full URL-canonicalization algorithm, the MVP **restricts the endpoint grammar to a subset in which every distinct endpoint has exactly one spelling** — anything outside the subset is invalid, not normalizable:

```
endpoint  := "https://" host [":" port] path
host      := lowercase ASCII reg-name: [a-z0-9] with [a-z0-9-] interior, dot-separated labels;
             NO IP literals (v4 or v6), NO trailing dot, NO IDNA/punycode processing (an xn-- label is
             treated as an opaque literal label), NO userinfo
port      := decimal, no leading zeros, 1-65535; MUST be omitted when 443 (":443" is invalid)
path      := "/" or one-or-more "/"-prefixed segments; segment chars ∈ [A-Za-z0-9._~-] only;
             NO percent-encoding anywhere (any "%" is invalid), NO empty segments ("//"),
             NO "." or ".." segments, NO trailing "/" unless the path is exactly "/"
             NO query ("?") and NO fragment ("#")
```

**Comparison everywhere in P2 is byte equality.** There are no transforms and therefore no transform-ordering ambiguity: two byte-distinct valid endpoints ARE distinct endpoints; every ambiguity class the grammar cannot represent (IPv6 spellings, port `0443`, trailing-dot hosts, `%2e` dot-segments) is **rejected at the source that supplied it** — frozen core ⇒ `malformed-core` (PS) · lane fact ⇒ `lane-endpoint-invalid` (P1) · allowlist entry ⇒ malformed policy (P0). m-3 owns this grammar; m-8 (comparison implementer, catalog-row producer) consumer-confirms it. Real provider endpoints (e.g. `https://api.openai.com/v1/responses`) sit inside the subset; a provider endpoint that ever falls outside it is a Step-4 grammar-extension through contract revision, never an ad-hoc normalizer.

**Accept/reject vectors (normative):** ACCEPT `https://api.openai.com/v1/responses` · `https://gateway.internal:8443/v1/chat` · `https://xn--bcher-kva.example/v1`. REJECT `http://api.openai.com/v1` (scheme) · `https://API.openai.com/v1` (case) · `https://api.openai.com.:443/v1` (trailing dot + explicit 443) · `https://[2001:db8::1]/v1` and `https://192.168.0.1/v1` (IP literals) · `https://api.openai.com/v1//responses` (empty segment) · `https://api.openai.com/v1/%2e%2e/admin` (percent) · `https://api.openai.com/v1/` (trailing slash) · `https://user@api.openai.com/v1` (userinfo) · `https://api.openai.com/v1?x=1` (query).

### §1.7 Policy-field semantics table (R2-F1 — every field's owner + authority, no inert bytes)
| Field | Semantic | Authority relationship |
|---|---|---|
| `schema` | version identity | literal; closed |
| `egress_class` | class marker | literal `provider-request`; closed |
| `pinned_lane` | MUST equal the run manifest's pinned lane at freeze (mismatch ⇒ the manifest is rejected at freeze, before any attempt exists) | policy asserts; manifest (m-10-written) is the run-binding authority |
| `endpoint_allowlist` | the P2 membership set | policy is the sole authority |
| `denied_header_names` | the P3 name set | policy is the sole authority, under the P0 membership invariant (which consumes the lane's declared auth-header fact) |
| *(method)* | — no policy field — | the lane catalog fact (m-8) is the sole method authority (P4) |

## §2 — Artifact 2: the E0 app-event schema (the worker-carried provider report)

### §2.1 Carrier + status
Carried in the **m-9 worker's existing `SITREP` relay body** — **no conductor schema/member/field change** (ratified §3a preserved; the F62 tuple applies to app-side/external artifacts ONLY). Every event floors at **E0 / `self_reported`**: never gate-satisfying, never promoting, never an input to any conductor verdict. The conductor treats the block as ordinary relay content; humans and app-side tooling may read it.

### §2.2 Event schema (`m3.app_event.v1`)
One JSON object per event, in a fenced block under the body heading `### APP-EVENTS (E0, self_reported)`:

```json
{
  "schema": "m3.app_event.v1",
  "event_kind": "provider_attempt",
  "phase": "denied | sent | completed | failed | cancelled | unknown",
  "scope": "attempt",
  "run_id": "…", "turn_id": "…", "attempt_id": "…", "turn_epoch": "0",
  "provider_lane_id": "…",
  "run_manifest_digest": "…", "policy_digest": "…",
  "deny_reason": "<§1.3a token — present iff phase=denied>",
  "event_evidence": "E0",
  "event_integrity": "self_reported",
  "reported_by": "m-9-worker/g<generation>",
  "event_ts": "<RFC3339>"
}
```

- `phase=unknown` mirrors m-10's `UNKNOWN_PROVIDER_OUTCOME` park — the report is allowed to say "I don't know," and the schema makes that a first-class value rather than forcing a fabricated outcome.
- **`phase=cancelled` (r4, F-m9-cancel):** the DETERMINATE NON-FAILURE terminal — the attempt was deliberately stopped (cancellation/interruption), mirroring the authoritative facts in m-9's `turn_terminal{turn_cancelled}` + m-10's terminal `CANCELLED` row. Distinct from `failed` (an actual fault occurred) and from `unknown` (genuine indeterminacy) — reporting a cancellation as either would be a false E0 claim; this token completes the enum so honesty never requires silence. `deny_reason` stays absent (present iff `phase=denied`, unchanged). **Reachability (H-14):** emission = m-9's E0 population for a cancelled attempt (their interim no-terminal-phase posture was CORRECT — honest silence over a false claim — and is superseded by this token); consumption = observability/human readers of the E0 report and the mirror-fidelity relation to the m-10 row — **the F62 E3 record has no `phase` field and the §3.4 evaluator never consumes phases** (verified: `phase` appears in §2 only), so no evaluator disposition exists or is owed for any phase member.
- `event_evidence`/`event_integrity` are **fixed literals** in v1 — a v1 event structurally cannot claim more than E0. `reported_by` is a claim, not a proof (§0.1).
- **`turn_epoch` is the canonical-decimal-uint64 STRING** (grammar `^(0|[1-9][0-9]*)$`, value < 2^64, numeric decoded comparison) — the m-10 §A.2 counter encoding, adopted here so the SAME event object carries ONE encoding on every surface it crosses (the CTRL-W `app_event` frame, the `pending_app_events` row, the SITREP copy). The event stays expressly non-trust-bearing (E0); the uniform encoding kills a three-surface variance, it does not promote the field (F-m9-L5-1, branch (a)).
- Events are **not digested artifacts**: they must parse under the closed schema, but JCS byte-canonicality is not a validity condition for an E0 report (it is for §1.5/§3.2, the digested artifacts). The durable copy is m-10's `pending_app_events` app-state table; the SITREP carriage is the courier-visible copy. Neither is conductor evidence.

### §2.3 Redaction rule (the object-typed negative route, applied)
A v1 event carries **identifiers, digests, enum outcomes, and timestamps only** — never prompt bytes, provider response content, headers, wire frames, or credential references. Raw `LLMRequest` objects, provider headers/stream frames/transport envelopes, and raw tool envelopes are **not valid conductor payload types** (ratified §3) and never appear here. Agent-authored prose elsewhere in the SITREP may quote/summarize model output — governed solely as relay content, with no provider-evidence status.

## §3 — Artifact 3: the F62 external E3 observation record (`m3.e3_observation.v1`)

### §3.1 Writer + home
Authored by the **separate integration harness / operator observer** — the writer IS `observer_id`, bound per §3.6. Stored **outside the conductor store**; it never becomes a conductor record. It is the ONLY artifact in the MVP that carries live-provider-turn evidence at E3.

### §3.2 Record schema — full field census
Universal fields (required at every scope): `schema` (literal `m3.e3_observation.v1`) · `scope` · `claim` (the observed claim, prose) · `observed_outcome` (`pass | fail`) · `observer_id` · `observation_ts` (RFC3339).

Identity/locator fields (required/forbidden per §3.3): `run_id` · `turn_id` · `attempt_id` · `artifact_ref` (closed enum: `app_main | m9_worker | m8_connector | release | tool_catalog | policy`) · `relay_id` (structured field — **never** parsed out of `claim`; F1).

Vector fields (required/forbidden per §3.3): `run_manifest_digest` · `tool_catalog_digest` · `app_main_build_digest`/`m9_worker_build_digest`/`m8_build_digest` **XOR** `release_digest` (exactly one form where builds are required) · `policy_digest` · `provider_lane_id`.

**Canonical encoding = RFC 8785 (JCS) + the §1.5 validity preconditions** (NFC, closed schema, no duplicate keys); the stored artifact is exactly the JCS output bytes. **Record identity + digest home (F1):** the record digest = lowercase-hex SHA-256 over the exact stored canonical bytes; it is **NOT a field inside the record** (no self-reference). It lives **outside** the record as the reference handle: the storage filename `e3-<digest>.json` and any cross-reference (e.g. from the Master+VP composite exit-test record) cite the record by that digest. **Verification procedure:** whenever a record is presented under a digest reference, the evaluator recomputes SHA-256 over the presented bytes and requires equality ⇒ else `non_applicable(record-digest-mismatch)`; non-canonical stored bytes ⇒ `non_applicable(malformed)`. A record presented without a digest reference gets no digest check (the digest is a referencing mechanism, not an applicability input).

**Digest sourcing is by canonical producer only** (ratified F62/F63): `run_manifest_digest` ← m-10 at run freeze · `tool_catalog_digest` ← the m-9-owned catalog build, mechanically verified at the F63 release-binding · the build digests / `release_digest` ← the build pipeline at the post-build RELEASE-BINDING event · `policy_digest` ← m-3 (§1.5).

### §3.3 Scope matrix — required / forbidden, total over all six scopes (F1; policy_digest placement corrected per R2-F2)
Rule: **required fields MUST be present; forbidden fields MUST be absent** (absent = the key does not appear; never null/empty); any violation ⇒ malformed. Universal fields are always required and never listed below.

| scope | required identity | required vector | forbidden |
|---|---|---|---|
| `build` | — (the release vector IS the identity: a build-scoped claim targets exactly the release its digests name) | the three build digests XOR `release_digest` · `tool_catalog_digest` | `run_id`, `turn_id`, `attempt_id`, `artifact_ref`, `relay_id`, `run_manifest_digest`, `policy_digest`, `provider_lane_id` — the policy is a run-bound artifact, not a build input (R2-F2); policy claims ride `artifact` scope |
| `artifact` | `artifact_ref` | **exactly the one digest field its `artifact_ref` names**: `app_main`→`app_main_build_digest` · `m9_worker`→`m9_worker_build_digest` · `m8_connector`→`m8_build_digest` · `release`→`release_digest` · `tool_catalog`→`tool_catalog_digest` · `policy`→`policy_digest` | every other identity + vector field (incl. `release_digest` when `artifact_ref` ≠ `release` — a covering digest cannot locate a single artifact) |
| `run` | `run_id` | full vector: `run_manifest_digest` · `tool_catalog_digest` · builds XOR `release_digest` · `policy_digest` · `provider_lane_id` | `turn_id`, `attempt_id`, `artifact_ref`, `relay_id` |
| `turn` | `run_id` + `turn_id` | same as `run` | `attempt_id`, `artifact_ref`, `relay_id` |
| `attempt` | `run_id` + `turn_id` + `attempt_id` | same as `run` | `artifact_ref`, `relay_id` |
| `relay_record` | `relay_id` + `run_id` | `run_manifest_digest` · `tool_catalog_digest` · builds XOR `release_digest` (the app half that produced/read the relay) | `turn_id`, `attempt_id`, `artifact_ref`, `policy_digest`, `provider_lane_id` (a relay observation is not a provider turn) |

`relay_record` scope observes the **app side** of a relay exchange (e.g. "the native relay tool produced stamped relay `<relay_id>`") — the **conductor side** of the same exchange is m-7's separately-bound leg (§3.5); `relay_id` here is a locator, never a conductor-identity binding.

### §3.3b Evaluation-context table (R2-F2 — where the target vector and the observer mapping come from, per scope)
The evaluator is **never invoked against an ambient "currently-running" notion**. Every invocation names a **claim context** — a specific, digest-identified artifact set — and every target-vector element is drawn from that context's canonical producer artifacts:

| scope | claim context (names the comparison target) | target-vector element sources |
|---|---|---|
| `build` | one named **F63 release-binding record** | build digests / `release_digest` + `tool_catalog_digest` ← that record |
| `artifact` | **immutable exact-digest evidence:** the context is the single target artifact digest, drawn per `artifact_ref` from a named release-binding record (`app_main`/`m9_worker`/`m8_connector`/`release`/`tool_catalog`) or a named run manifest (`policy`) | the one corresponding element of the named context artifact |
| `run` / `turn` / `attempt` | one named **run**: its frozen run manifest + the release-binding record that manifest carries | `run_manifest_digest`, `policy_digest`, `provider_lane_id` ← the manifest · builds/`release_digest`, `tool_catalog_digest` ← the release binding |
| `relay_record` | one named **run** (same two artifacts) | `run_manifest_digest` ← the manifest · builds/`release_digest`, `tool_catalog_digest` ← the release binding |

**Observer registry source (all scopes, uniform):** the `observer_id` enum and its recorder-boundary mapping are **pinned in THIS contract (§3.6) and bound at the Master+VP first-stage interface-lock** — not run-scoped, not manifest-carried, not mutable at runtime. Changing the registry is a revision of these contract bytes through the same review machinery (new SHA, fresh pair review); historical provenance therefore cannot silently change (R2-F2).

### §3.4 The applicability evaluator (m-3-owned rule; acquire-then-compare per R2-F3)
**Rule (ratified, realized here):** an E3 record applies to a claim **only while ALL bound digest/lane fields equal the named claim context's vector; any mismatch ⇒ non-applicable (re-observe).**

Deterministic procedure, fail-closed at every step, first failure wins (verdict reasons are a closed ordered enum, like §1.3a):
1. **Reference integrity:** if the record is presented under a digest reference, recompute and compare ⇒ `non_applicable(record-digest-mismatch)`.
2. **Well-formedness:** exact JCS canonical bytes · parses under `m3.e3_observation.v1` · closed parsing (unknown field / duplicate key / wrong type ⇒ malformed) · satisfies the §3.3 required/forbidden matrix ⇒ `non_applicable(malformed)`. (Unknown-field rejection is the structural absorb-refusal — a conductor-identity field in this record is malformed, never silently carried; F65/F68.)
3. **Observer provenance (F1, §3.6):** `observer_id` is in the §3.6 registry AND equals the recorder identity of the boundary the record was obtained from ⇒ `non_applicable(observer-mismatch)`.
4. **Location:** the claim under evaluation names its scope instance; the record's **scope-specific identity fields per §3.3** must equal it ⇒ `non_applicable(wrong-instance)`. Identity fields locate; they are not the applicability test.
5. **Acquisition (R2-F3 — strictly before any comparison):** obtain EVERY vector element required at the record's scope from the named claim context per §3.3b. **Any element unobtainable ⇒ `non_applicable(vector-unavailable)` — never "assumed unchanged," and no comparison is attempted on a partial vector.**
6. **Comparison:** with all elements in hand, compare each for **exact equality**. ANY inequality ⇒ `non_applicable(mismatched: [fields])`, the field list ordered by the **canonical field order** (R2-F3): `run_manifest_digest` → `tool_catalog_digest` → `app_main_build_digest` → `m9_worker_build_digest` → `m8_build_digest` → `release_digest` → `policy_digest` → `provider_lane_id`.
7. Verdict: `applicable` | `non_applicable(reason)`. **No partial credit, no grace window, no transitive inference** (an E3 at `scope=run` says nothing about a different run; `scope=build` evidence never substitutes for an `attempt`-scoped claim; `scope=artifact` speaks only for the one artifact its `artifact_ref` names).

**Mutation behavior (normative, realizing the §10 annex row):** mutating ANY single bound artifact — the app-main/m-10 binary, the m-9 worker binary, the m-8 connector binary (or the covering release bundle), the manifest, the policy, the lane, the catalog — changes its element in every claim context that includes it, and step 6 returns `non_applicable(mismatched: [that field])` for every prior record whose scope requires that field. Mutating the **conductor** changes NO field of this record type — it invalidates m-7's relay-leg exit evidence, not the provider-turn E3 (F65).

### §3.5 The F65 scope boundary + the F68 upstream edge (consume + confirm; authority-corrected per F4; scope-precise per R2-F4)
- **The common invariant, for ALL six scopes (R2-F4): every field, context, and verdict of this record type is APP/PROVIDER-VERTICAL-ONLY — no scope's `applicable` verdict ever validates (or invalidates) conductor identity.** The conductor service identity (build digest + governing config identity) is **deliberately not a field of `m3.e3_observation.v1`** and is never an evaluator input.
- **"Provider-turn leg" names the `run`/`turn`/`attempt` claims specifically** — the live provider turn F65 splits from the relay leg. `build`/`artifact` records are app-vertical build/artifact evidence; **`relay_record` records are the APP-SIDE HALF of a relay exchange only** — the conductor side of that same exchange is m-7's separately-bound leg, and no `relay_record` verdict speaks for it.
- **Absorb-refusal (structural, mine to enforce):** a conductor-identity field creeping into this record type is malformed at §3.4 step 2 — the conductor identity cannot ride the app/provider vector.
- **The omission edge (F4 — stated at the authority this seat actually has):** this schema **cannot and does not** enforce the presence of m-7's half; the composite exit-test record and its join requirements are **owned by Master+VP**, not by m-3 or m-7. What THIS contract binds is the bounded statement: **an m-3 app/provider-vertical E3 alone does not satisfy the §7 exit proof** — the exit test additionally requires the relay-exchange leg evidenced by conductor-captured observe-as-send E1/E2 records bound to the conductor identity produced under **m-7's F68 producer contract** (`step3-mvp-design-m7`, supplement `…-043459`, m-3's named upstream). It is a **close requirement of this lane** — confirmed against m-7's final bytes before any interface-lock readiness claim — and a **requirement m-3 asserts the Master+VP composite join must carry** (both halves present), not a property my schema enforces.
- Mutation consequences split per §3.4: app-vertical mutation ⇒ the app/provider E3 non-applicable; conductor mutation ⇒ the relay-leg exit evidence goes stale — **not** the app/provider E3.

### §3.6 Observer provenance binding (F1; contract-pinned per R2-F2)
The `observer_id` registry is **closed and pinned in this contract** (bound at the Master+VP first-stage interface-lock; immutable for Step-3; any change is a reviewed revision of these bytes):

| `observer_id` | recorder boundary |
|---|---|
| `integration-harness` | the harness process's own output root — the directory named in the harness invocation record, written only by the harness |
| `operator` | direct operator handoff — the operator supplies the record file to the evaluator invocation themselves |

The evaluator's caller presents each record **together with the recorder identity of the boundary it was obtained from**; §3.4 step 3 requires that identity to equal the record's `observer_id`. A record surfacing from any other path (e.g. handed over by the worker, found in app state, embedded in a relay body) has **no recorder boundary** and is `non_applicable(observer-mismatch)` — an agent cannot self-supply an E3. Ceiling per §0.5: this is convention/boundary provenance against confusion, not a cryptographic signature against forgery.

## §4 — The instrumented-negative posture (ratified §3/§10, realized)

The deny→zero-send family is **instrumented build/integration-test evidence** — a fake transport/executor counting **zero** invocations — or an **independent live observer**. These paths bypass the conductor by design, so:
- they are **never conductor-observed** and **never become E2 via an app statement written into a conductor record**;
- the §10 annex rows this contract realizes: **"denied provider request sends nothing"** (fake-transport counter = 0 on the denied path) · **"credentials attach only after authorize"** (secret resolver never invoked on the denied path — the §1.2 placement property) · **"E3 is properly bound"** (the §3.4 mutation behavior, incl. the F65 split);
- the evidence-status table, stated once: instrumented negatives = **E2 as instrumented-TEST evidence about the build, held app-side/externally** · the worker's report of any outcome = **E0** · the external observer's live turn = **E3** · conductor-leg relay evidence = **E1/E2, conductor-captured, m-7-bound** — four rungs in two custody worlds (test/observer-held vs conductor-held); none convertible into another by restatement, and the two E2s are different custody, never the same channel.

## §5 — Annex: normative examples (the re-review vector set)

**A. One example record per scope** (universal fields elided to `…` for brevity; presence/absence is the normative content; note `build` no longer carries `policy_digest` per §3.3):

```json
{"schema":"m3.e3_observation.v1","scope":"build","claim":"release passes conformance","observed_outcome":"pass","release_digest":"aa…","tool_catalog_digest":"bb…","observer_id":"integration-harness","observation_ts":"…"}
{"schema":"m3.e3_observation.v1","scope":"artifact","claim":"connector enforces deny order","observed_outcome":"pass","artifact_ref":"m8_connector","m8_build_digest":"dd…","observer_id":"integration-harness","observation_ts":"…"}
{"schema":"m3.e3_observation.v1","scope":"artifact","claim":"policy is the ratified one","observed_outcome":"pass","artifact_ref":"policy","policy_digest":"ca364710764c3fb5fa6ca0f2faa6795e6aa6f49d3d55877ca4def10d540a3030","observer_id":"operator","observation_ts":"…"}
{"schema":"m3.e3_observation.v1","scope":"run","claim":"run r1 served exactly the 8-name set","observed_outcome":"pass","run_id":"r1","run_manifest_digest":"ee…","tool_catalog_digest":"bb…","release_digest":"aa…","policy_digest":"cc…","provider_lane_id":"lane-codex-1","observer_id":"operator","observation_ts":"…"}
{"schema":"m3.e3_observation.v1","scope":"turn","claim":"turn t1 stayed within bounds","observed_outcome":"pass","run_id":"r1","turn_id":"t1","run_manifest_digest":"ee…","tool_catalog_digest":"bb…","release_digest":"aa…","policy_digest":"cc…","provider_lane_id":"lane-codex-1","observer_id":"operator","observation_ts":"…"}
{"schema":"m3.e3_observation.v1","scope":"attempt","claim":"attempt a1 = one live provider turn","observed_outcome":"pass","run_id":"r1","turn_id":"t1","attempt_id":"a1","run_manifest_digest":"ee…","tool_catalog_digest":"bb…","release_digest":"aa…","policy_digest":"cc…","provider_lane_id":"lane-codex-1","observer_id":"operator","observation_ts":"…"}
{"schema":"m3.e3_observation.v1","scope":"relay_record","claim":"native tool produced stamped relay rel-9","observed_outcome":"pass","relay_id":"rel-9","run_id":"r1","run_manifest_digest":"ee…","tool_catalog_digest":"bb…","release_digest":"aa…","observer_id":"integration-harness","observation_ts":"…"}
```

**B. Multi-failure policy vector:** wrong lane + non-allowlisted endpoint + `authorization` header ⇒ `lane-mismatch` (§1.3a position 4; evaluation stopped there).
**C. Canonical-policy equivalence vector:** the same logical policy with `denied_header_names` unsorted or duplicated is **invalid** (P0 `policy-unavailable`), not differently-hashed; the §1.5 normative example is the exact 255-byte canonical form and its digest.
**D. Mutation vectors:** rebuild m-8 only ⇒ every prior `run`/`turn`/`attempt`/`relay_record` record and the `artifact_ref=m8_connector`/`release` records go `non_applicable(mismatched: [m8_build_digest])` (or `[release_digest]` under the covering form); edit the policy ⇒ all `run`/`turn`/`attempt` records + `artifact_ref=policy` records go non-applicable (`build` and `relay_record` unaffected — neither binds `policy_digest`); mutate the conductor ⇒ **zero** records of this type change status (the relay-leg exit evidence, m-7's, goes stale instead).
**E. Mixed unavailable-plus-mismatch vector (R2-F3):** an `attempt`-scoped record evaluated against a context where `tool_catalog_digest` is unobtainable AND `policy_digest` differs ⇒ **`non_applicable(vector-unavailable)`** — acquisition (§3.4.5) fails before comparison (§3.4.6) is ever attempted; the mismatch is not reported.
**F. Endpoint accept/reject pairs:** the §1.6 vector list is normative.

## §6 — Consumer set + close conditions

| Consumer | Confirms |
|---|---|
| **m-8** | the §1 policy it enforces (predicates, §1.3a order, P0 floor incl. the membership invariant, digest verification) + the §1.6 endpoint grammar (as comparison implementer + catalog-row producer, incl. method authority per §1.7) + the §2 events its outcomes populate |
| **m-9** | the §2 SITREP carriage (heading, fenced block, redaction rule) + `phase=unknown` mirroring + the `phase=cancelled` emission realization (their E0-population half, post-token per the F73 ladder) |
| **m-10** | the manifest-digest producer seam (`policy_digest` + `run_manifest_digest` at run freeze, incl. the §1.7 `pinned_lane` freeze-time equality) + the `pending_app_events`/attempt-outcome rows |
| **m-7 (upstream)** | m-3 confirms the §3.5 scope boundary against m-7's authored F68 conductor-identity contract (reciprocal: their supplement mirrors this edge) |

**Close conditions (per the §7 return path):** m-3.implementer pair-review of the FINAL bytes → report-only SITREP to master naming the approved bytes + hash → consumer confirmations on master's direction → the m-7 F68 consume/confirm before any interface-lock readiness claim. **No self-declared lock; no PLAN, token, or code is authorized by this document.**

## §7 — Fold log
- r4 (2026-07-18): **the cancellation phase, accepted as asked** (master routing `…-033727`; requirement + rationale = m-9 as E0 populator, `step3-mvp-design-m8/RECONCILE-planner-20260717-212600`; name/shape mine): §2.2 `phase` enum gains **`cancelled`** — the determinate non-failure terminal mirroring m-9's `turn_terminal{turn_cancelled}` + m-10's terminal `CANCELLED` row (r27 `db199b0d…`). m-9's ruling was this schema's own §2.2 rationale applied: `failed` lies about a fault, `unknown` fabricates indeterminacy (`:144` reserves it for genuine indeterminacy) — the enum was NOT total over determinate outcomes and honesty required silence; the token closes that gap. m-9's interim no-terminal-phase posture blessed as correct and superseded. H-14 discharged in the bullet: emission = m-9's population; consumption = E0 readers + row mirror-fidelity; the E3 record/evaluator never consume `phase` (verified — no evaluator disposition owed). §6 m-9 consumer row extended with the emission realization (post-token, non-gating per the F73 ladder). Three touches + header/status; nothing else moved. r3 SHA-256 `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.
- r0 (2026-07-16): authored per dispatch `…-041700` + supplement `…-043510`, against ratified r7 `2f75f2a1…` (hash verified in-session). Carried floors from the held `step3-amend-m3-egress` lane. SHA-256 `6438c6300643de5de50e32c6eba4e86152c0afb008d35758777c9b4078dc09eb`.
- r1 (2026-07-16): F1–F4 folded from `DESIGN-REVIEW-implementer-20260716-050730` (must-revise): total scope matrix + structured identities + observer binding + record-digest home (F1); P3 → reserved-header-name exclusion (F2); total deny order + canonical endpoint + closed policy parsing (F3); F65 authority bounding (F4). SHA-256 `d98dd6ad2021bbcb6229f8bb8e51f825fefc89f4e887ceb1260ea1b9be5ab843`.
- r3 (2026-07-17): **F-m9-L5-1 disposed, branch (a)** (master routing `step3-mvp-confirm-m3/RECONCILE-…-024139`; m-9's Leg-5 CONFIRM carried the finding; m-9-preferred + master-recommended branch adopted as the owner's call): §2.2 `turn_epoch` re-pinned from JSON number to the **canonical-decimal-uint64 STRING** (m-10 §A.2 grammar) — one encoding on every surface the event object crosses (frame · store row · SITREP copy). The prior "legal because non-trust-bearing" reading (r2 confirm edge-4, ledger L6) was true but incomplete: the same bytes cross a surface whose rule textually claims them, and three surfaces × two encodings is exactly the confusion class this domain exists to kill — branch (b)'s exemption would have maintained the variance for zero benefit. Non-trust-bearing status unchanged; no other byte moved. r2 pair approval (`…-054800`) was byte-bound ⇒ this fold reopens pair review. r2 SHA-256 `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- r2 (2026-07-16): R2-F1..R2-F4 folded from `DESIGN-REVIEW-implementer-20260716-052637` (must-revise on r1). **R2-F1** — canonical JSON = **RFC 8785 (JCS)** (aligned with the m-10 stage-1 contract), NFC as an explicit validity precondition, stored-bytes = exact JCS output (no terminal LF), a normative exact-byte example policy + computed digest; §1.6 re-cut from "canonical form with transforms" to a **restricted unique-spelling grammar** (no IP literals / percent / dot-segments / trailing dot / leading-zero ports; byte equality, zero transforms) + accept/reject vectors; `method_by_lane` REMOVED (the lane catalog fact is the sole method authority, §1.7); the P3 membership invariant made a P0 validity rule (digest-bound coverage). **R2-F2** — §3.3b evaluation-context table (every scope names its claim context; no ambient "currently-running"; `artifact` = immutable exact-digest evidence); the observer registry re-homed from run-scoped to **contract-pinned + interface-locked** (§3.6; historical provenance cannot silently change); `policy_digest` removed from `build` scope (a run-bound artifact, not a build input). **R2-F3** — §3.4 re-ordered to **acquire-then-compare** (any unobtainable element ⇒ `vector-unavailable` before any comparison; mixed vector = annex E); `mismatched:[fields]` ordered by the pinned canonical field order. **R2-F4** — §3.5 restated: the common invariant is app/provider-vertical-only for ALL scopes; "provider-turn leg" reserved for `run`/`turn`/`attempt`; `relay_record` = the app-side half only; Master+VP composite + m-7 confirmation unchanged.
