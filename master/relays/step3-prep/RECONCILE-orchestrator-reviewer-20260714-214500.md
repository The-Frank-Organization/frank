## RECONCILE -- VP re-review of the folded Step-3 kickoff-of-record

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-prep
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- the operator's org/scope ratification remains valid; these are contract-consistency corrections inside that ratified direction
GRILL_REQUIRED: no -- this relay reviews the kickoff; the required m-8/m-9 design grills and GRILL_LOCK_IDs remain pre-build gates
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-prep/RECONCILE-orchestrator-planner-20260714-213000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator
SUBJECT: revise -- the six directional folds are accepted, but current locked bytes make the chosen_model amendment mandatory and require explicit provider-egress, trusted-config, full-T4-transport, and live-evidence gates

VERDICT: revise

The planner folded the prior six findings in good faith, and the ratified Step-3 shape remains sound. I do not co-sign the current bytes because four implementation-bearing statements still conflict with the locked contracts they claim to consume. This is a narrow correction pass: preserve the org expansion, ownership split, vertical sequence, four-axis lane identity, design-only boot, and full carry ledger.

## Findings

### 1. Blocker -- the provider-send trust boundary does not yet exist

Kickoff §1 says every model request crosses "the m-3 egress boundary BEFORE m-8 translation" and that credentials, endpoints, and payload mutation stay in "m-7 config" (`master/STEP-3-KICKOFF.md:13`). Neither referenced locked surface currently supports that claim:

- m-3's locked egress gate is specifically the conductor **local-outbox** scanner, dormant until the m-6 away-mode bridge (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:129-136`).
- The implementation confirms it is present-but-dormant, outbox-record-shaped, and unimported by production (`frank/internal/egress/egress.go:1-7,51-56`; `frank/test/fixtures/s5_egress_test.go:100-128`). Its default confidentiality classifier flags ordinary `gpt-`, `claude-`, `llama`, and `model_name=` bytes (`frank/internal/egress/rules.go:22-43`), so it cannot simply be placed in front of a normal provider request.
- m-7's locked config is a conductor-composed policy artifact with sections authored by m-2 through m-6; it defines no provider credential, endpoint, secret-source, rotation, or redaction contract (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:107-125`). Its only governed external-send mechanism is the dormant local outbox (`:130-132`).

"Before translation" is also not a sufficient final-wire guarantee: translation, compatibility handling, endpoint binding, and authentication occur afterward. The egress-approved semantic object could therefore differ from the bytes sent.

Required fold before m-8/m-9 design-lock or build:

1. Name an **m-3 owner-authored review-driven amendment** for a provider-request egress class, hosted by m-7 and consumed by m-8/m-9. Define whether authorization is a final-wire gate or a deliberately specified pre/post pair; in either case no adapter mutation may occur after the final authorization point.
2. Name an **m-7 owner-authored trusted-config/credential amendment**, with m-1 trust/secret-boundary review and m-8 consumer review. It must settle secret sourcing, endpoint allowlisting/binding, rotation, redaction, and the rule that credential values never enter catalog, snapshots, records, seat surfaces, or evidence artifacts.
3. Add fail-closed negatives: planted secret/PII or disallowed endpoint => zero provider send; translation cannot mutate after final authorization; adapter, endpoint, credential, and pinned lane must agree. Do not inherit away-email `egress_blocked` behavior or model-name confidentiality rules without an explicit owner decision.

### 2. Blocker -- the m-4/m-2 routing-record amendment is already mandatory

Kickoff §3 and §7 retain the amendment as conditional: **if** `chosen_model` cannot bind the four-axis lane (`master/STEP-3-KICKOFF.md:34,64`). The audit condition is already false on current bytes:

- The locked m-4 Step-3 promise says the Step-1 API/record does not change, and the assignment row records one model-valued `chosen_model` (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:164-166,205-218`).
- Bucket-binding and replay semantics are explicitly expressed as `chosen_model in members(declared_bucket)` and model membership, not provider/serving/compat lane identity (`:41,93-106,128-136`).
- The live FieldSpec row exposes `chosen_model` but no provider, serving profile, compatibility mode, or canonical lane reference (`frank/internal/fieldspec/registry.json:175`).

An exact `{model_id, provider_id, serving_profile_id, compat_mode}` execution choice therefore cannot be replay-complete without either silently redefining `chosen_model`/bucket membership or making provider and serving selection after the recorded decision. Both violate the kickoff's own no-silent-widen/no-silent-fallback rule.

Required fold: replace the conditional audit with a **mandatory m-4 owner-authored review-driven amendment and m-2 FieldSpec review before m-8/m-9 lock or runtime build**. The owners choose the representation -- canonical lane reference or explicit lane tuple -- but the accepted route record must bind the exact executable lane and snapshot while preserving R2's non-gate-referenceability.

### 3. High -- the T4 transport gate must include the owning m-x PM

Kickoff §0 says the operating model is T4 on frank while "the m-x pairs stay on manual relay for now" (`master/STEP-3-KICKOFF.md:7`), and §8 only names unlaunched slice-team seats (`:67`). That is acceptable during the design-only pre-build phase, but it is not an executable statement of the activated T4 model.

Part F requires the complete escalation path `slice -> master -> owning m-x planner -> back`, says those escalations use `submit`/`project`/`read` with no operator hand-relay, and says everything flows through frank (`master/CYCLE-PLAYBOOK.md:376-385`). B13 makes the same transport change normative (`master/PROTOCOL-DEVIATIONS.md:170`). A slice-only relaunch would stop at the exact PM boundary the new tier exists to exercise.

Required fold: explicitly limit manual m-x relaying to m-8/m-9's pre-build AUDIT/DESIGN phase. Before the first T4 code token, the relaunch roster must include the T4 pair, master router/arbitrator seat, and owning m-x PM seat on frank. The shakedown must prove one full escalation and response round trip `T4 -> master -> m-x -> master -> T4`, with stamped provenance and live-store-to-durable-trail export, without operator transport.

### 4. High -- the spine exit needs a live evidence floor and provider-path criterion

V1 currently promises "one real governed end-to-end turn" and V2 a "second provider" (`master/STEP-3-KICKOFF.md:45-46`), but the exit does not say whether mocks, aliases, or one compatibility endpoint behind two names satisfy those claims. The result could pass fixtures while never exercising the credential, endpoint, wire, streaming, cancellation, and egress boundaries whose abstraction the spine claims to prove.

Required fold:

- V1 and V2 each require an E3 live call against a real external provider endpoint, in addition to deterministic E2 conformance/failure fixtures.
- V2 must exercise two independently bound provider paths. Two catalog aliases routed through the same effective endpoint/protocol path do not earn the `>=2 providers through one interface` claim.
- Name the negative exit legs: egress reject => zero network send; above-ceiling tool call => zero execution; absent/invalid exact lane at V3 => typed `routing_unavailable` or `human_decision_required`, never fallback.

## Prior Findings Closed

The following portions of the fold are accepted and should not be reworked: pi/opencode as prior art rather than spec; frank-owned normative contract; m-8 Provider Adapters plus m-9 Model Runtime; factual-catalog versus policy-overlay single writers; the four-axis lane key and pinning; V1 vertical -> V2 portability -> V3 routing; design-only boot followed by audit/design/grill/consumer-review/reconcile/lock; non-terminal spine wording; and the full Step-3 carry ledger.

No fresh operator ratification is required if the planner makes the bounded contract-consistency folds above. If the planner instead changes the ratified product scope or rejects a named owner amendment, route that delta to the operator. This relay grants no charter amendment, boot, design-lock, PLAN, implementation, merge, or deployment authority.

## Verification

- Incoming planner relay and `master/STEP-3-KICKOFF.md` read in full; incoming exact-file relay lint -> OK.
- Locked/current surfaces checked at the cited bytes: m-3 egress design and dormant implementation, m-7 trusted config/outbox design, m-4 routing record, live m-2 FieldSpec registry, Part F, and B13.
- `frank/` remains clean on `main` at `502e06c`; no source byte was changed.
- New relay exact-file lint -> OK; INDEX EOF survival check -> reviewer row present after filing.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-prep/RECONCILE-orchestrator-reviewer-20260714-214500.md and appended its master/relays/INDEX.md row; no kickoff, charter, domain, frank source, branch, commit, push, merge, tag, live-store, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
