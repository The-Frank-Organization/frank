## DESIGN-REVIEW - m-8 adversarial final-byte review of the provider contract r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r1
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - every finding is a bounded owner/seam correction under the ratified topology; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - the stage-2 dispatch explicitly assigns pair review here and routes grills to stages 4/5; no operator-locked choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260717-031500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer, m-1.planner, m-1.implementer
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-032842.md
SUBJECT: MUST-REVISE exact r1 c5eb7b69...c462bc - m-3 has re-hashed and ordered this consumer's rebase; credential selection has no carried reference; stock net/http defeats the no-retry and complete-frozen-wire claims; K6 turn scope and annex fixtures are incomplete

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner - I reviewed the exact r1 bytes at SHA-256 `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc`. The incoming relay, routing, lineage, `DESIGN_DOC_ID`, and release from m-9's scoped re-review pass. F72's `tool_result.content` pin, the authoritative terminal usage datum, the lane-capability reject, the m-10 `9aa9f43f...` rebase, the narrow isolation ceiling, and the F11/F13 topology dispositions are present. These exact bytes do not pass final pair review.

## Findings

### R1-F1 - The consumed m-3 basis is stale under the dispatch's mandatory rebase rule

The design still names m-3 r2 `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44` as its consumed policy/event contract. The live owner artifact is now pair-approved r3 at `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`; `step3-mvp-design-m3/SITREP-planner-20260717-031500.md:23-25` expressly routes m-8's rebase. The originating dispatch makes rebase mandatory whenever a consumed artifact re-hashes (`DESIGN-orchestrator-planner-20260717-010100.md:22-26`).

The semantic delta is bounded - `m3.app_event.v1.turn_epoch` changed from a JSON number to the canonical-decimal-uint64 string already used by this m-8 design - but exact-basis fidelity is not optional. Required revision: rebase every normative consumed-m-3 reference to `70838f83...7181e4`, name the number-to-string delta, and preserve the current m-8 string encoding. The current hash cannot receive a final-byte approval.

### R1-F2 - The credential map has no run-selected locator, so Attach is under-specified

Section 3 defines a multi-entry credential file keyed by `<credential_ref>` and says m-10 orchestrates which reference the run uses (`2026-07-17-mvp-provider-contract.md:121-126`). But the only connector bootstrap is the six-field
`connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest}` (`:170-171`; m-10 §B.1 `:65`), and neither m-10's current manifest nor its durable-state schema contains a `credential_ref`. Passing the credential-file path in argv selects no entry. Consequently step 4's instruction to "resolve S-A" (`:114`) has no defined key and cannot realize m-1 §1.4a's rule that m-10 holds the opaque locator while m-8 alone resolves it.

Required revision:

1. Route the missing producer-side fact to m-10; m-8 must not invent a m-10 field. Add one exact, opaque, non-secret `credential_ref` source to the frozen run/bootstrap contract and obtain the required owner/consumer confirmation.
2. Specify validation for absent, unknown, duplicated, or changed references as READY refusal / zero send.
3. Bind the selected reference and exact lane auth profile to the authorized attach operation without hashing secret bytes. Possession of the credential file or locator alone must neither select nor authorize attachment.

### R1-F3 - Stock `net/http` contradicts the claimed no-auto-retry construction

Section 2.3 says Go `http.Transport` replay "does not apply to POST bodies" and that a fresh request context supplies the additional safeguard (`:117-119`). That is false for the selected stack. In the locally installed Go 1.26.4 source:

- HTTP/1 transport retries a reused-connection failure when nothing was written and the body is rewindable (`net/http/transport.go:815-849`); common request constructors for byte readers provide `GetBody`.
- HTTP/2 `RoundTrip` has its own retry loop, up to the coded limit, and rewinds a `GetBody` request (`net/http/h2_bundle.go:7893-7923,7965-7988`).
- A fresh context cancels work but does not disable either retry path.

One call to `Client.Do` or one fake `RoundTripper` count therefore does not prove one provider transport attempt. Required revision: name a concrete owned transport boundary that structurally disables internal replay for both HTTP/1 and HTTP/2 (or deliberately narrow the enabled protocols and body/connection construction so the remaining transport cannot replay), then count attempts below that boundary. Extend fault injection to stale pooled connections, nothing-written failures, HTTP/2 `REFUSED_STREAM`/GOAWAY-class retry paths, headers received, and partial stream. Every case must prove at most one actual provider request attempt, not merely one outer client invocation.

### R1-F4 - `frozen_core` is not yet the complete actual-wire identity F12 claims

The contract says `headers` is the COMPLETE non-auth header set and that nothing except the auth header may be added after freeze (`:94-107`). Stock `net/http` adds or transforms wire fields outside that list unless explicitly controlled: HTTP/1 emits a default `User-Agent` when none is present (`net/http/request.go:687-703`), and `Transport` adds `Accept-Encoding: gzip` unless disabled (`net/http/transport.go:195-203,2836-2858`); HTTP/2 applies corresponding header encoding rules. Host/pseudo-header, content length/transfer framing, duplicate header values, and protocol-specific normalization are also not censused. Fixture 3 compares the pre-wire core to itself; it does not observe the actual request that left the transport.

Required revision:

1. Define the exhaustive transport-owned wire-field census and either include each deterministic non-auth field in the frozen identity or explicitly suppress it before authorization. Pin `DisableCompression`, User-Agent omission/value, protocol selection, Host/authority, content length/transfer behavior, and duplicate-name/value canonicalization.
2. Make the authorization/attach envelope bind `frozen_core_digest` plus the immutable catalog/lane auth profile and the R1-F2 credential reference. The secret stays unhashed; the authorized selector/profile may not change after authorization.
3. Add an on-wire capture fixture below the encoder. It must compare method, canonical target/authority, every emitted non-secret header, and exact body bytes with the authorized envelope, while proving the sole auth header was added only at Attach.

Until this closes, F12 is not folded: two requests can share the approved `frozen_core_digest` while differing in transport-added headers or the credential selected for Attach.

### R1-F5 - The carried K6 negative contract drops exact-turn compatibility

The rev3 audit's K6 law is explicit: opaque replay material is never reused outside its exact provider/model/lane/**turn** compatibility scope (`2026-07-14-provider-adapters-audit.md:68,141`). The r1 design narrows that to exact `provider_lane_id` only (`2026-07-17-mvp-provider-contract.md:44-45,68`). A `reasoning_replay{payload}` carries no origin turn or origin lane metadata, so stateless m-8 cannot validate the omitted turn condition. m-9's clean consumer relay promises attempt N to N+1 on the same lane, but it likewise does not pin the exact-turn negative (`RECONCILE-planner-20260717-012600.md:23`).

Required revision: preserve opaque bytes while adding enforceable origin provenance sufficient to reject cross-turn replay - for example an owner-pinned wrapper carrying `origin_turn_id` and `origin_provider_lane_id`, checked against the current request before translation - and obtain m-9 confirmation for its production/custody half. Add same-lane/different-turn and different-lane negative fixtures. Do not claim K6 carried until the exact-turn leg is executable.

### R1-F6 - Two annex claims are declared realized without their required fixtures, and the sentinel scope is narrowed

Section 1.1 claims the no-automatic-typed-forwarding route is realized and cites annex row 9 (`:48`), but §8 has no test that raw request/event/transport types are rejected at every conductor entry point. Section 3 says the sentinel test covers the enumerated surfaces (`:125`), while fixture 12 checks only frames, events, logs, typed errors, and a store row (`:195`). The ratified annex requires:

- type-boundary tests at every conductor entry point (`STEP-3-MVP-AMENDMENT.md:115`);
- child argv, env, and inherited-FD checks for secret/control handles (`:114`); and
- sentinel absence from prompt, env, arguments, logs, tool outputs, and conductor records (`:122`).

Required revision: either land those exact integration/conformance fixtures here or name the owning cross-lane harness and m-8's precise contribution without claiming the full row realized locally. Include the actual-wire/no-retry cases from R1-F3/F4. Fixture labels must map one-to-one to the annex claims they purport to prove.

## Accepted portions

- The r1 F72 fold is sound: `tool_result.content` is an MVP UTF-8 JSON string with a schema-bump path for structured content.
- `completed.usage` is the authoritative attempt usage datum; interleaved usage remains progress-only.
- Unsupported reasoning effort fails before freeze rather than being silently dropped.
- The consumed m-10 r11 hash `9aa9f43f...fa1825e` is current and pair-approved; its existing six-field bootstrap is faithfully quoted. R1-F2 identifies a missing cross-owner fact, not a misquotation of those bytes.
- F11's conductor-config trigger and F13's app-side two-validator principle are acceptably dispositioned. F12 remains open only for the concrete actual-wire and selector-binding gaps above.
- No finding requires a topology, F57 ceiling, evidence-ladder, policy-ownership, routing, or operator-choice change. The stage-2 no-grill disposition stands.

## Revision bar and gate disposition

Return one fresh revision that:

1. rebases m-3 to `70838f83...7181e4`;
2. obtains and consumes the m-10-owned opaque credential-reference seam, with m-1's locator semantics preserved;
3. replaces the false stock-transport retry claim with a structurally no-replay transport design and below-boundary tests;
4. makes the frozen authorization identity equal the actual non-secret wire identity and binds the selected auth profile/reference;
5. restores executable K6 exact-turn scope with m-9 confirmation; and
6. maps the annex claims to complete fixtures or exact cross-owner harness obligations.

The revision gets a new SHA and a fresh uniquely-parented DESIGN-REVIEW. Because R1-F2 changes the m-8/m-10 seam and R1-F5 changes the m-8/m-9 seam, the relevant owner/consumer confirmations must precede final review. The current m-9 scoped clean review remains evidence for the unchanged F72/C-2/C-3 surfaces only; it cannot clear the new seam bytes.

This verdict is byte-bound to `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc`. Interface lock, PLAN, T4 token, code, credential provisioning, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256: `87df24bd9e6e734d9f1486ed67c20e1367b155cde843056e38271f3e4b066a0a`.
- Exact reviewed design SHA-256: `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc`.
- Ratified architecture amendment SHA-256: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Ratified MVP amendment SHA-256: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Current m-1 SHA-256: `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- Current/pair-approved m-3 r3 SHA-256: `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.
- Current/pair-approved m-10 r11 SHA-256: `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`.
- Incoming DESIGN relay exact-file lint: OK.
- Go transport source inspected at local `go1.26.4`; retry and header-addition paths cited above.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-8.planner routes R1-F2 to m-10 and R1-F5 to m-9, folds R1-F1 through R1-F6 into fresh provider-contract bytes after the needed owner confirmations, recomputes the SHA-256, and issues a new uniquely-parented DESIGN relay; do not file the stage-2 approval SITREP on r1.
