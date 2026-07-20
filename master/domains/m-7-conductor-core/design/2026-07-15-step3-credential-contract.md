# Step-3 owner amendment — the provider credential / trusted-config contract (r3)

- **DESIGN_DOC_ID:** `step3-amend-m7-cred-credential-contract`
- **Status:** r3 — folds the pair review's F7–F10 (`step3-amend-m7-cred/DESIGN-REVIEW-implementer-20260714-232820`, must-revise; F1–F6 confirmed closed there) whole; technical re-review requested. NOT closeable until the m-8/m-9 consumer packet feeds the final review (kickoff §6-amendment B14) and Master+VP reconcile all three Step-3 owner amendments. No lock is claimed by this document.
- **Revision log:** r1 (authored, `DESIGN-planner-20260714-230832`) → r2 (F1–F6 folded, `DESIGN-planner-20260714-232405`) → r3 (F7 the secrets census pinned as a catalog s8-v2 FORBIDDEN-ONLY entry with staged, never-red activation · F8 component-wise descriptor confinement + nonblocking open + dot-ID rejects + bounded descriptor read · F9 derivation returns an opaque `Selected` capability that Attach consumes, profile wire bytes + reserved-header policy pinned · F10 the canonical engine-v5 descriptor delta + composition-validation table, no longer deferred to PLAN).
- **Lineage:** owner amendment extending the m-7 config-host design (`2026-07-11-s8-config-host.md`, lock `s8-design-m7-config-r13`) and touching the executor-host env posture (r5) — review-driven amendment per the standing terminal-lock discipline. Dispatched by `step3-amend-m7-cred/DESIGN-orchestrator-planner-20260715-005530` under kickoff §1(b)/§6.4(b).
- **Reviewers:** m-7.implementer (adversarial pair review) · m-1 (secret boundary) · m-8 (consumer) · reconciles with `step3-amend-m3-egress`.
- **GRILL_REQUIRED:** yes — `GRILL_LOCK_ID: step3-amend-m7-cred-grill-r3` (§11; supersedes `-r2`, `-r1`).

## §1 Threat model + claim ceiling

The contract's enemy is **confusion, not a same-uid attacker**. The claim, stated at the grain the mechanisms actually enforce (r2/F5): *credential material cannot flow **accidentally** through any enumerated interface into a governed byte surface — store records, projections, catalog, snapshots, evidence artifacts, logs, the local outbox, or any seat-reachable surface — and cannot authorize a send to an endpoint the operator did not bind.* **Deliberate in-package misuse, reflection, or process-memory inspection remain possible under the same-uid/same-code residual — review-visible and canary-detectable, not type-impossible.** No OS-secrecy claim; no forgery-robustness-by-construction claim (the shelved wrap milestone).

## §2 The load-bearing split — governed BINDING vs out-of-store MATERIAL

1. **The binding (governed, non-secret):** a new engine-member section `provider_bindings` — engine **v5** by the r13 version rule. Each row:
   - `credential_id` — stable handle; **also the material file name** (the r1 `secret_source.id` was redundant and is REMOVED — r2/F2; `secret_source` is now `{kind}` only, v1 sole kind `root_file`);
   - `provider_id` — must match the m-8 catalog's provider axis;
   - `endpoint_allowlist` — structured rules per §4;
   - `auth_profile` — **closed enum, v1 `{bearer, x_api_key}`** — the governed selection of how material attaches to the wire (header name + shape are profile data, never adapter-supplied headers — r2/F2);
   - `lane_ids` — optional exact-lane allowlist (absent ⇒ every lane of that provider);
   - `generation` — **an operator assertion only, not proof material rotated** (r2/F6); monotone non-decreasing, enforced at the §7 transition gate.
   Binding rows carry **zero secret bytes and zero secret digests**. Binding changes ride the existing machinery unchanged: §7 `config_change` on the engine member, per-version descriptor arm, adjacent v4→v5 hop, restart-effective, composite-digest-chained. No member-set change ⇒ no bless/adoption.
2. **The material (secret, never governed):** `<root>/secrets/<credential_id>` (opening + grammar per §3), outside the digest chain, loaded through the §2.3 transaction into the §6 vault, never entering any record body or projection. `<root>/secrets/` enters the census as a **forbidden-only entry** per §2.5 (r3/F7 — NOT a thirteenth exact-live-equality row).

**Why the split is forced:** §7 `config_change` bodies carry member bytes into the append-only store — a credential-in-member design writes secret values into permanent history on every rotation. The split makes that unrepresentable rather than forbidden-by-review.

### §2.3 The startup publication protocol (r2/F1 — pinned to the executable path)

At `502e06c` the startup order is: `config.Load` (`main.go:137-140`) → existing hosts built (`:157-164`) → **recovery** `frankrecover.RunWithProcessor` (`:259`, which performs `ValidateGenesis` + replay) → the Ready branch constructs the loop. The provider sender **does not exist today** (m-8 audit F-4: zero outbound HTTP); it is a NEW host, and the protocol places it exactly:

1. **Position:** credential material is first read **after** `RunWithProcessor` returns a valid `Ready` (genesis digest verified, replay complete) and **before** the provider sender is constructed — a **startup insertion in `run()`'s Ready branch, not a recovery hook and not a refactor of existing hosts** (no existing host ever receives credentials).
2. **The transaction:** load *every* binding-referenced source into a **temporary set** under §3's rules → validate the **complete** set → **publish exactly once** (a single assignment of an immutable vault) → construct the provider sender **from the published vault** → proceed to serve. The vault is immutable post-publication (rotation is restart-effective, §5).
3. **Failure = one typed, path-free startup disposition:** any source failure discards the entire temporary set; **no vault is published, no sender is constructed, no provider socket ever opens, and serve does not start** (`credential-material-unavailable`, carrying at most the `credential_id` per the §6.4 matrix — never a path). Rationale: a declared binding the runtime cannot honor is a trusted-config integrity failure, the same class as a missing member — phase-0 posture.
4. **Zero declared bindings is lawful:** a store with no `provider_bindings` serves normally with no sender constructed; any provider-send request refuses typed (`credential-selection-unavailable`). Pre-Step-3 stores are untouched.
5. **No provider sends during recovery/replay — structural:** the sender is constructed only after recovery completes and the vault publishes; there is no earlier object to send through.
6. **Fixtures (FX-CRED-1/2 extended):** zero bindings ⇒ serve, no sender · N valid ⇒ vault publishes once, sender live · **one-bad-among-N ⇒ no partial set, no sender, no serve, the typed disposition** — the review's exact leg.

### §2.4 The canonical engine-v5 descriptor delta + composition rules (r3/F10 — nothing left to PLAN)

**Descriptor delta (the per-version schema, the F1–F6/cadence pattern):** v5 allowed-keys = v4's set ∪ `{provider_bindings: array}`; ceiling `> 4` → `> 5`; adjacent hop `<= 4` → `<= 5`; reader ceiling 5. **`provider_bindings` is OPTIONAL at v5; a PRESENT-EMPTY array is REJECTED** (zero bindings has exactly one representation: absence — §2.3.4's lawful state). A v4-or-lower store carrying the key is the version-smuggle reject.

**Composition-validation table (every rule typed, all enforced at BOTH the descriptor layer — load and transition-candidate — and composition, the lane_vcs two-surface posture):**

| Node | Rule | Failure |
|---|---|---|
| `provider_bindings` | array of closed row objects, **rows sorted by `credential_id`** (the descriptor-rows precedent) | descriptor/composition reject |
| row keys | exactly `{credential_id, provider_id, endpoint_allowlist, auth_profile, generation}` + optional `lane_ids` — unknown/extra/missing key rejects | reject |
| `credential_id` | supply-ID charset ≤128 · **`.` and `..` explicitly rejected** (§3) · **unique across rows** | reject |
| `provider_id` | supply-ID charset, non-empty | reject |
| `endpoint_allowlist` | **non-empty** array of closed rule objects `{scheme, host, port?, path_prefix?}` per §4.2 · duplicate rules reject | reject |
| `auth_profile` | closed enum `{bearer, x_api_key}` | reject |
| `lane_ids` | if present: **non-empty** array of unique strings (absent ≠ empty; empty rejects) | reject |
| `generation` | non-negative integer; monotone non-decreasing per `credential_id` at the §7 transition gate | reject / `config-version-transition` |
| cross-member references | **whether `provider_id`/`lane_ids` resolve against the m-8 catalog is PROVISIONAL pending the B14 packet** — v1 default: unresolved references are **selection-time refusals** (`credential-selection-unavailable`), upgradeable to composition-time once m-8 pins the catalog's load-time surface; m-7's structural/uniqueness rules above are final regardless | (see left) |

**Fixtures (FX-CRED-5 extended):** unknown/extra/missing key · wrong node kind per field · duplicate `credential_id` · empty allowlist · present-empty `provider_bindings` · empty `lane_ids` · unsorted rows · lawful zero-binding v5 · the full transition matrix (smuggle/rollback/skip/ceiling).

### §2.5 The secrets census — a FORBIDDEN-ONLY entry, not a thirteenth row (r3/F7)

The r13-locked descriptor's `canonical_path_families.rows` is a **closed twelve-row exact-live-equality census** (rows `{id, relative, forbidden, directory}` sorted by ID; the live oracle enumerates exactly the `StoreRootConfigPaths` homes). `<root>/secrets/` cannot inhabit that family: it is **conditional by design** (zero-binding stores lawfully lack it, §2.3.4) and it is not a config home — what the census must guarantee about it is only the **forbidden token** (no `/secrets/` byte in any output family), never existence. So:

1. **The catalog member gains an optional `forbidden_only` list** — entries `{id, forbidden}`, v1 content exactly `[{id: "secrets", forbidden: "/secrets/"}]` — a catalog SCHEMA change, honestly versioned: **marker `s8-v1` → `s8-v2`**, riding the existing catalog machinery (`validateCatalogSchema` gains the v2 arm; `validateStringMarkerTransition` gains the `s8-v1`→`s8-v2` adjacent hop; the r13 §5.2 singular catalog arm carries the change; the same bytes land in the source artifact per the §4 drift law).
2. **The twelve-row exact-equality family is UNTOUCHED** — count stays twelve; the leak-scan law consumes rows' forbidden tokens **∪ `forbidden_only`** tokens.
3. **Staged, never-red activation:** the catalog `s8-v2` change and the engine `v5` change are **independent singular §7 records in either order with no invariant-red intermediate** — the forbidden-only law is satisfiable with or without bindings or the directory (it forbids bytes in outputs; it asserts no live state), and engine v5 without the catalog entry merely lacks the extra scan token until the second record lands (defense-in-depth deepens; nothing reds). No atomic two-member transition machinery is needed or introduced.
4. **The directory is operator-created at provisioning** (0700, owner-matched — §3/§7): genesis is UNTOUCHED; zero-binding stores never have it; the §2.3 loader requires it only when bindings are declared. r2's zero-binding and no-genesis-break claims survive verbatim.

## §3 Secret source — descriptor-safe open + wire-safe byte grammar (r2/F3; r3/F8 the confinement + nonblocking mechanics)

- **v1 sole source `root_file`:** `<root>/secrets/<credential_id>`.
- **Opening (component-wise descriptor confinement — r3/F8; the landed `fs_worker.go:209-274` rooted-walker posture, and NOT leaning on the flock, which locks `conductor.lock`, not a directory):**
  1. open the root directory itself as an FD (`O_DIRECTORY|O_CLOEXEC`);
  2. `openat(rootfd, "secrets", O_DIRECTORY|O_NOFOLLOW|O_CLOEXEC)` — **the parent is descriptor-opened no-follow too**, so a symlinked `secrets` pointing outside the root fails here, not after passing leaf checks; `fstat` THIS directory descriptor: directory · mode 0700 · owner uid == process uid (**the exact mode/owner pair pending the m-1 ruling; this is the r3 default**);
  3. `openat(dirfd, credential_id, O_RDONLY|O_NOFOLLOW|O_NONBLOCK|O_CLOEXEC)` — **`O_NONBLOCK` so a FIFO cannot block the open before the non-regular reject is reachable**; then `fstat` the leaf descriptor: **regular file** (a FIFO/device/dir rejects here, having never blocked) · mode 0600 · owner-matched;
  4. **`credential_id` values `.` and `..` are EXPLICITLY rejected** (the inherited supply-ID charset admits both — the r2 "traversal unrepresentable" claim was wrong at exactly those two values and is retracted); the remaining charset (`[a-zA-Z0-9._:-]`, ≤128, no `/`) confines every other name to the directory;
  5. **bounded DESCRIPTOR read: read at most 4097 bytes from the open FD** — >4096 bytes read ⇒ reject — so growth after `fstat` cannot make the read unbounded (the `st_size` check is advisory only);
  6. identity-stable throughout: every check is on the descriptor the bytes are read from; a name-swap after open changes nothing the loader sees.
- **Byte grammar (these bytes become an HTTP auth field):** strip **at most one** terminal `LF` or `CRLF`; then REJECT: empty material · any remaining `CR`/`LF`/`NUL` · any byte outside the selected `auth_profile`'s charset (v1 profiles: printable ASCII `0x21–0x7E`, i.e. RFC 7230 field-value-safe with no spaces) — header injection and unbounded reads are rejected at load, not discovered on the wire.
- **Every diagnostic is path-free and material-free** (§6.4 matrix).
- **FX-CRED-2 extended legs:** leaf symlink · **parent (`secrets`) symlink** (r3) · non-regular · **FIFO-no-writer nonblocking open** (r3) · **dot-ID (`.`/`..`) rejects** (r3) · wrong mode · wrong owner · oversized · **grow-after-stat bounded read** (r3) · empty · double-trailing-newline · embedded CR/LF/NUL · swap-during-load (descriptor identity holds).
- Rejected/deferred source kinds unchanged from r1 (env rejected: heritable + no provenance artifact; keychain deferred behind a version bump).

## §4 Endpoint binding — canonical final-wire rule (r2/F4; kills the pi F-2 fusion)

1. **Direction of authority (unchanged):** the LANE names the endpoint; a credential can only AUTHORIZE it, never supply it.
2. **The structured rule:** `{scheme, host, port, path_prefix?}` — `scheme` MUST be `https` (composition reject otherwise) · `host` lowercase ASCII (non-ASCII/IDN rejected at composition — no normalization games) · `port` explicit, with 443 default-normalized on both sides · **no userinfo** (a rule or request URL carrying userinfo is rejected) · `path_prefix` optional with **segment-boundary semantics**: it matches iff the request path `== prefix` or begins `prefix + "/"` (`/v1` never matches `/v1evil`) · **query/fragment policy:** rules never carry them; request URLs may carry a query (it is part of the frozen wire bytes, not part of the match), never a fragment.
3. **Canonical comparison, canonical-form-only URLs:** the sender rejects request URLs containing dot-segments or percent-encoded `/` (`%2F`)/`.` sequences outright — comparison happens on canonical bytes or not at all (encoded-path traversal is a refusal, not a normalization exercise).
4. **The final-wire freeze:** the complete wire request — **method + canonical URL + headers-sans-auth + body — is FROZEN before final authorization** (m-3's evaluation point sees the frozen object); the **sole** post-freeze change is the host's auth-attach (§6). No SDK/adapter mutation after the freeze — frank owns the HTTP client construction (no pi-style custom-fetch/onPayload seam exists to violate it).
5. **Redirects: automatic following DISABLED in v1.** The client is constructed redirect-off; a 3xx is a terminal provider response surfaced through m-8's normalized-event contract, never auto-followed — so "a redirect target receives the credential" is unrepresentable, not merely checked. (Re-entrant authorization for redirect targets is a possible later amendment if a real provider requires it; it does not exist in v1.)
6. **Retries: every attempt is a fresh final-wire act** — re-frozen, re-authorized at its actual wire destination, re-attached. A retry whose target/bytes differ from the authorized freeze is a new authorization or no send. Retry *policy* (backoff, idempotency keys, which failures retry) is the m-9/m-3 consumer packet's lane — the invariant owned here is **no attempt without its own authorization**.
7. **The four-way agreement (restated with F2's direction):** sender refuses unless `lane.provider_id == binding.provider_id` ∧ frozen endpoint ∈ `binding.endpoint_allowlist` ∧ the m-3 authorization is present for exactly this freeze ∧ the binding was **derived by §4a** — else typed refusal, zero network send.
8. **Negatives (FX-CRED-3 extended):** redirect-not-followed · host-case + default-port equivalence · encoded-path traversal refusal · path-prefix sibling (`/v1` vs `/v1evil`) · retry-target-change re-authorization.

### §4a Credential selection is a GOVERNED DERIVATION (r2/F2 — the tautology removed)

r1 let the request envelope carry `credential_ref`, making the check "ref == binding(ref)" — circular, and with two same-provider credentials either would pass. r2 removes request choice entirely:

1. **No adapter-, runtime-, or request-supplied credential reference exists.** The m-8 request envelope carries the **lane** (and the frozen wire bytes); m-9 asks by lane; neither can name a credential.
2. **The sender DERIVES the binding from governed state alone:** the unique `provider_bindings` row with `binding.provider_id == lane.provider_id` ∧ (`binding.lane_ids` absent ∨ `lane.id ∈ binding.lane_ids`). **Zero matches ⇒ `credential-selection-unavailable`; more than one match ⇒ `credential-selection-ambiguous`** — both typed refusals, zero send. Disambiguation of two accounts on one provider is a governed act (the operator scopes `lane_ids` in the §7-transitioned binding), never a runtime choice.
3. **Single writer:** the relation lives in m-7's binding rows keyed by m-8's lane identity (m-8 owns lane facts; m-7 owns which credential serves which lane — the operator authors both through their governed paths). The m-8 consumer seam consumes lane-ID stability; it never writes or selects bindings.
4. **The auth-attachment shape is the binding's `auth_profile`** — governed data, not adapter headers (§2.1); the exact wire bytes per profile are pinned in §6.1a.
5. **Negative (the review's leg):** two credentials, same provider, same endpoint, no `lane_ids` scoping ⇒ ambiguity refusal proves the wrong-account path is closed; scoped `lane_ids` ⇒ exactly the scoped binding attaches.
6. **Derivation returns an opaque capability, not an identifier (r3/F9):** `Derive(lane, frozenEndpoint)` returns a **`vault.Selected`** — an unexported-field value constructible only inside the vault package, internally carrying the derived binding **and the frozen endpoint it was derived for**. `Attach` consumes `Selected` (§6.1), so **no credential-ID string ever crosses the final boundary** — a call site cannot substitute a different valid credential after derivation (the r2 `Attach(req, credentialID)` free-string parameter is REMOVED), and an attach against bytes other than the freeze the derivation saw refuses mechanically (`Selected` binds the freeze).

## §5 Rotation (unchanged from r1, with the F6 generation correction)

Material rotation = atomic file replace + restart; **zero governed bytes move**. Binding rotation = ordinary §7 v5 `config_change`. Revocation = binding-row removal (§7) and/or material deletion. `generation` is an **operator assertion** recorded for visibility — monotone non-decreasing at the transition gate, **carrying no claim that material actually rotated** (nothing governed can prove that; saying otherwise would be a false-provenance claim).

## §6 The vault, redaction, and non-flow (r2/F5 — capability API named, claims narrowed)

1. **The vault (`internal/provider/vault`, host-internal):** material loads into `vault.Vault` — immutable after the §2.3 publication; the material field is unexported; **no exported accessor returns raw bytes**. The capability pair (r3/F9): **`Vault.Derive(lane, frozenEndpoint) (Selected, error)`** (the §4a governed derivation, pre-authorization — it feeds the four-way check) and **`Vault.Attach(req, sel Selected) error`** (post-authorization, sole call site: the governed sender) — `Selected` is opaque, package-constructed only, and freeze-bound, so attachment is mechanically tied to the unique derivation result; no free-form credential identifier exists at the boundary. Adapters and runtimes receive **neither `Selected`, nor `Secret` values, nor arbitrary auth headers**.

### §6.1a The v1 auth-profile wire bytes + reserved-header policy (r3/F9 — exact, executable)

| Profile | Attached bytes (exactly one header) |
|---|---|
| `bearer` | `Authorization: Bearer <material>` |
| `x_api_key` | `X-Api-Key: <material>` |

- **Reserved header set, matched case-insensitively:** `{Authorization, X-Api-Key, Proxy-Authorization}`. **Any occurrence of any reserved name — any case, any profile, any count — in the pre-freeze adapter/request headers ⇒ typed refusal** (`credential-reserved-header-present`), never overwrite, never coexist, and cross-profile occurrences refuse identically (an `Authorization` header on an `x_api_key` lane is just as dead).
- Attach adds **exactly one** header to the frozen request; a duplicate is therefore unrepresentable post-check.
- **`Request.Host` must be empty or byte-equal to the canonical frozen URL authority** — else typed refusal (a Host-header/URL split is a mis-pointed send in disguise).
- **Negatives (FX-CRED-7 extended):** ID/handle substitution (a `Selected` derived for freeze A refuses request B) · mixed-case reserved header · duplicate reserved header · cross-profile reserved header · Host/authority mismatch.
2. **Enumerated safe representations (each one tested, none assumed):** `fmt` verbs (`%s`, `%v`, `%q`, `%#v`) yield `[credential:<id>]` · `encoding/json` marshals to the same token · the type implements **neither** `encoding.TextMarshaler` nor `encoding.BinaryMarshaler` (asserted by interface-check test) · a structured-logging leg proves the token, not material. **Outside the enumeration** (reflection, unsafe, deliberate in-package access): not claimed — review-visible and canary-detectable per §1.
3. **Attach-after-authorization (unchanged):** m-3 evaluates the frozen, pre-attach request; the attach is the sole named post-freeze change (reconciled with the m-3 amendment's grill — if their lock lands final-wire-only, the attach point IS the final wire step and no second authority exists).
4. **The post-attach request object is never logged, captured, or evidenced** — a named rule, enforced by the canary sweep.
5. **The census law (extended per F6):** the planted-canary fixture sweeps store bytes + logs + outbox + evidence for the canary in **raw, base64, and hex** forms and for the `secrets` path family ⇒ zero hits; the `[credential:<id>]` token is the approved reference form and is NOT leakage.

### §6.4 The visibility matrix (r2/F6 — what may appear where; PLAN invents nothing)

| Item \ Surface | trusted config (member bytes) | config_change history | host-internal sender state | request envelope (pre-freeze, m-8) | frozen wire bytes (post-attach) | records / projections / catalog / snapshots | evidence / logs / outbox | seat surfaces | typed diagnostics |
|---|---|---|---|---|---|---|---|---|---|
| material bytes | **never** | **never** | vault only | **never** | **wire only** (the attach) | **never** | **never** (canary: raw+b64+hex) | **never** | **never** |
| concrete source path | never (kind only) | never | loader-internal | never | never | never | never | never | **never** (path-free rule) |
| `credential_id` | yes | yes | yes | **no** (no ref field exists — §4a) | no (material attaches, ids don't) | no | only as the `[credential:<id>]` token | no | yes |
| binding facts (provider_id, allowlist, auth_profile, lane_ids) | yes | yes | yes | no | no | no | no | no | class tokens only |
| whole binding row | yes | yes (member bytes in the §7 record — non-secret by design) | yes | no | no | **never** (N6, restated at this grain) | no | **never** | no |
| `generation` | yes | yes | yes | no | no | no | no | no | no |
| post-attach request object | — | — | transient wire buffer only | — | the wire itself | never | **never** (named rule) | never | never |

## §7 The operator provisioning decision — ROUTED, not resolved (unchanged from r1)

Operator-owned: (i) which providers/accounts; (ii) each provisioning act (`<root>/secrets/<id>` write + binding ratification `{credential_id, provider_id, allowlist, auth_profile, lane_ids?}`); (iii) the first-E3-use go before V1/V2's live legs. Form: a HUMAN_GATE relay per credential; the §7 `config_change` then records the ratified binding in governed history. Nothing here resolves at design close.

## §8 Consumer seams (r2 deltas marked)

- **m-1:** as r1, plus the §3 mode/ownership pair (file 0600 + parent 0700, owner-matched) explicitly awaiting their ruling.
- **m-8:** the envelope carries the **lane and the freeze inputs — no credential field of any kind** (r2/F2 tightened); their DESIGN feeds my final review (B14).
- **m-9:** asks by lane; additionally consumes §4.6 — every retry attempt re-enters authorization (their Q4 packet supplies the policy; my invariant is per-attempt authorization).
- **m-3:** evaluates the **frozen** object (§4.4); the attach is the sole named post-freeze change; joint fixtures FX-CRED-8 + the planted-secret zero-send/zero-attach leg.
- **m-7.executor:** deny-by-default child env unchanged; material never exported into any spawn env.

## §9 Fail-closed negatives

| # | Input | Result |
|---|---|---|
| N1 | a declared binding's material absent/malformed/wrong-mode/oversized | typed `credential-material-unavailable` at §2.3 — no vault, no sender, no serve |
| N2 | frozen endpoint ∉ allowlist / provider mismatch / non-canonical URL | typed refusal, zero network send |
| N3 | planted secret in any request field | m-3 class fires on the frozen pre-attach object ⇒ zero send, zero attach (joint) |
| N4 | adapter/runtime attempts material access | no API exists to hand it (§6.1); in-package breach = review-visible + canary-detectable |
| N5 | v4 store carrying `provider_bindings` | typed version-smuggle reject (the FX-VCS-10 pattern at v5) |
| N6 | a binding row in any projection/seat/evidence surface | census red (per the §6.4 grain — config member bytes + §7 history are the lawful homes) |
| N7 | two matching bindings for one lane (§4a) | `credential-selection-ambiguous`, zero send |
| N8 | provider 3xx | surfaced terminal, never auto-followed with credentials (client redirect-off) |
| N9 | a retry attempt without its own authorization for its actual bytes | zero send |
| N10 | one-bad-among-N at startup | no partial vault, no sender, no serve |
| N11 | a reserved auth header (any case/count/profile) pre-freeze, or Host ≠ frozen authority | typed refusal, zero send (§6.1a) |
| N12 | attach with a `Selected` derived for a different freeze | mechanical refusal (§4a.6) |

## §10 Fixture families (RED-first at build)

`FX-CRED-1` §2.3 protocol: zero-bindings serve · N-valid single publication · one-bad-among-N (N10) · `-2` §3 source legs (symlink · non-regular · mode · owner · oversized · empty · newline grammar · embedded control · swap-during-load) · `-3` §4 endpoint legs (mismatch · canonical equivalence · encoded traversal · prefix sibling · redirect-off · retry re-authorization) · `-4` the canary census (raw+b64+hex; token approved) · `-5` v4→v5 transition matrix (smuggle/rollback/skip/ceiling) · `-6` rotation (material swap + restart ⇒ new material live, zero governed-byte diff; generation monotonicity at the gate) · `-7` §6.2 representation enumeration (fmt/json/interface-absence/logging) + the §6.1a attach negatives (handle substitution · reserved-header case/duplicate/cross-profile · Host mismatch) · `-8` (joint m-3) freeze→authorize→attach ordering + zero-send/zero-attach · `-9` §4a selection (derivation · unavailable · ambiguous · lane_ids scoping · two-account negative) · `-10` the §2.5 census legs (forbidden-only token in the leak scan · staged catalog-s8-v2/engine-v5 activation in both orders, never red).

## §11 GRILL_LOCK

```text
GRILL_LOCK_ID: step3-amend-m7-cred-grill-r3
GRILL_REQUIRED: yes
SUPERSEDES: step3-amend-m7-cred-grill-r2, -r1 (all prior resolutions carried except the r2 traversal-unrepresentable claim, corrected by F8; four added by the F7–F10 folds)
GRILL_SOURCE:
- plan/design/audit relay read: STEP-3-KICKOFF.md §1/§4/§6 · the cue (…-005530) · the sibling m-3 cue (…-005520) · step3-audit-m-8/AUDIT-planner-…-224500 (F-2, F-4) · the pair review DESIGN-REVIEW-implementer-20260714-231621 (F1–F6)
- code/docs inspected: frank/cmd/frank/main.go:137-169,259-278 · internal/recover/recover.go:29-41 · internal/egress/rules.go:15-50 · the r13 config machinery at 502e06c · s8-executor-host r5 (env posture) · pi auth/types.ts:8-12 + models.ts:248 (via the m-8 audit)
- questions answered from codebase: secret source + the §2 split (config_change bodies enter the store) · the §2.3 position (sender is greenfield; recovery does ValidateGenesis) · endpoint direction (the F-2 fusion) · rotation (material outside the digest chain) · redaction layers
- questions asked operator: NONE resolved in-grill; ONE routed (below)

Resolved decisions:
- binding/material split — engine-v5 section vs out-of-store 0600 root file — source: code
- secret source v1 = root_file only; env rejected; keychain deferred — source: code + design
- endpoint authority: lane names, credential authorizes; canonical structured rules; canonical-form-only URLs — source: code (m-8 F-2) + design
- credential selection = GOVERNED DERIVATION at the sender (provider match + optional lane_ids scope; ambiguity refuses); no request/adapter/runtime credential reference EXISTS — source: review F2 + design (r2)
- the final-wire freeze + attach-as-sole-post-freeze-change; automatic redirects DISABLED v1; every retry attempt re-authorized — source: review F4 + design (r2)
- the startup publication transaction: post-recovery, pre-sender, all-or-nothing, one path-free typed disposition; zero-bindings lawful — source: code (main.go/recover.go) + review F1 (r2)
- rotation = material swap + restart, zero governed bytes; generation = operator assertion, monotone, no rotation-proof claim — source: code + review F6
- redaction = vault capability API + enumerated-and-tested representations + attach-after-authorization + canary census (raw/b64/hex); claims narrowed to enforced mechanisms — source: review F5 + design (r2)
- the secrets census = a catalog s8-v2 FORBIDDEN-ONLY entry (the 12-row exact-equality family untouched); staged never-red activation (catalog s8-v2 ∥ engine v5, either order); directory operator-created at provisioning, genesis untouched — source: review F7 + the r13 census bytes (r3)
- source opening = component-wise descriptor confinement (root FD → secrets dir O_DIRECTORY|O_NOFOLLOW → leaf O_NOFOLLOW|O_NONBLOCK), dot-IDs explicitly rejected, bounded 4097-byte descriptor read — source: review F8 + fs_worker.go posture (r3)
- attachment = Derive→Selected (opaque, freeze-bound)→Attach; exact profile wire bytes; case-insensitive reserved-header refusal (never overwrite); Host==frozen authority — source: review F9 (r3)
- the engine-v5 descriptor delta + composition table pinned in-design (optional provider_bindings, present-empty rejected, sorted rows, closed keys, uniqueness, typed timing; cross-member references provisional pending B14) — source: review F10 + the lane_vcs/cadence patterns (r3)
- claim ceiling = confusion-resistance, D5 same-uid + same-code residuals stated — source: standing lock posture

Rejected alternatives:
- credentials as config-member bytes — secrets would enter append-only history on every rotation
- secret digests in governed bytes — offline-guessing surface
- caller/adapter-supplied apiKey/headers/baseUrl (pi posture) — the F-2 fusion; structurally absent
- request-envelope credential_ref (r1's own shape) — tautological check, no governed truth home for selection (review F2)
- free-string `Attach(req, credentialID)` (r2's own shape) — reintroduced in-host substitution after governed derivation (review F9)
- a thirteenth exact-live-equality census row for `secrets` — a conditional directory cannot inhabit an exact-equality family; forbidden-only is the honest census grain (review F7)
- new top-level config MEMBER — forces member-set adoption machinery for no gain
- live credential reload · automatic redirect following · IDN/normalization tolerance — each a later deliberate amendment, not an implied behavior

Still operator-owned:
- the provisioning decision (§7): providers/accounts · provisioning acts · binding ratifications · the first-E3-use go — routed as a HUMAN_GATE surface gating m-8 Q5 / the V1-V2 E3 legs, NOT this design's close

Design-lock impact:
- extends s8-design-m7-config-r13 by amendment: engine v5 `provider_bindings` (§2.4 descriptor delta + composition table) + the §2.3 startup publication step + the catalog s8-v2 `forbidden_only` census entry (§2.5) + the §6.4 visibility matrix as census grain
- adds conductor-core: the vault package (Derive→Selected→Attach capability pair, §6.1/§6.1a) · the governed sender with freeze/derive/authorize/attach ordering · redirect-off client construction · the §3 component-wise confined loader
- names consumer contracts: m-8 lane-only envelope · m-9 per-attempt authorization · m-1 mode/ownership ruling · joint m-3 ordering + zero-send fixtures
```

## §12 Open items pending the consumer packet (block CLOSE, not re-review)

1. m-8's request-envelope/freeze-input shape — their DESIGN feeds my final review (B14).
2. m-3's disposition grammar at the freeze point — their amendment's grill (final-wire-only vs pre/post).
3. m-1's ruling on the §3 mode/ownership pair.
4. m-9's Q4 retry/idempotency policy (my per-attempt-authorization invariant is fixed; their policy composes).
5. ~~the exact v5 descriptor bytes~~ — **pinned in-design at §2.4 (r3/F10)**; only the mechanical OLD→NEW owner edits remain for build time (the E1–E10 pattern), inventing nothing.
6. The cross-member reference-validation upgrade (composition-time vs the v1 selection-time default) — decided when m-8's catalog load-time surface lands (B14, §2.4's provisional row).
