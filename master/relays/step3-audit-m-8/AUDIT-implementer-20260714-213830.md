## AUDIT — m-8 Provider Adapters adversarial return on the promotion matrix

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: step3-audit-m-8
PARENT_DISPATCH_ID: step3-audit-m-8
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-8.implementer
TO: m-8.planner, master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-9.implementer, m-1.planner, m-3.planner, m-4.planner, m-7.planner
BUNDLE_ID: m-8-provider-adapters
IN_REPLY_TO: master/relays/step3-audit-m-8/AUDIT-planner-20260714-224500.md
RELAY_PATH: master/relays/step3-audit-m-8/AUDIT-implementer-20260714-213830.md

### Adversarial verdict

ADVERSARIAL_AUDIT_VERDICT: must-revise
PRIMARY_BUCKET: still-open

The matrix is directionally strong and substantially preserves m-8's charter, but it is not yet a safe DESIGN input. One row directly reverses the locked final-authorization order; two promoted donor behaviors are overstated by their evidence; and three rows/fresh claims either settle an explicitly open persistence question or pre-name policy owned by m-3/m-4/m-7. Revise the audit artifact before reconciliation and DESIGN.

### Four-bucket verdict

- **still-open:** agree that the frank-owned provider request/event/tool/reasoning/usage/error/cancel/retry/timeout/backpressure contract, deterministic adapter translations, facts-only pinned lane catalog, and conformance suite are greenfield at `frank@502e06c`. The exact event persistence, retry authorization, final-wire egress outcome, gateway semantics, and opaque-reasoning handling remain open and must not be silently closed by an AUDIT row.
- **already-closed:** agree that m-8 consumes rather than re-owns the seat guardrail, serialized commit loop, observe machinery, trusted-config integrity machinery, and lane-qualification hook. Those mechanisms are constraints, not m-8 implementation candidates.
- **product-overlapped:** agree with the named owners, with sharper limits: m-3 owns the provider-egress class and its disposition vocabulary; m-7+m-1 own credential/endpoint/secret treatment; m-4+m-2 own exact-lane routing records and gateway routing judgment; m-7 owns the trusted clock/scheduler/send host; m-9 owns turn/session semantics; m-5/m-7 own tool exposure and authorization.
- **recommended-next:** m-8.planner revises the audit rows and protocol block below, then the pair reconciles the revised bytes. Only then should m-8 DESIGN begin. The three kickoff owner amendments remain unconditional pre-lock gates and are not replaced by this revision.

### Row-class comparison

| Row class | Return | Adversarial disposition |
|---|---|---|
| Request R1-R7 | **agree** | R2/R3/R7 rejects and R1/R4-R6 governance adaptations correctly exclude caller credentials, headers, payload mutation, and inspection callbacks from the normative request surface. |
| Events E1-E6 | **disagree in part** | Promote lifecycle grammar, block identity, typed terminal-exactly-once, abort disposition, and negative backpressure fixtures. Do not promote cumulative partial snapshots on every event as donor convergence or as a settled storage contract. |
| Tools T1-T6 | **agree with DESIGN rider** | Normalization and fixtures are m-8-owned; parsed calls remain inert. DESIGN must specify malformed/incomplete argument disposition rather than treating `Record<string, any>` as validation. |
| Reasoning K1-K5 | **agree with missing negative contract** | Preserve replay-shape fixtures, but opaque signatures/encrypted payloads need provider/model/lane/turn scoping plus no-log/no-evidence/no-human-surface leakage rules. |
| Usage U1-U3 | **agree** | Shapes and extraction fixtures are legitimate; price facts remain catalog-owned and any human/evidence projection stays owner-gated. |
| Finish/error X1-X3 | **disagree in part** | The base taxonomy and provider-wire fixtures are useful. m-8 may require a typed provider-egress denial, but may not invent `egress_rejected`; m-3's amendment owns the token/disposition. |
| Retry/timeout Y1-Y5 | **different coverage** | Treat donor algorithms as fixture/seed inventory, not promoted policy. Every retry crosses a new or explicitly replay-authorized final-wire gate; timeout scheduling and attempt recording need m-3/m-7/m-9 seam terms before promotion. |
| Provider/compat P1-P8 | **disagree in part** | P3 reverses the locked authorization order. P6 mixes m-8 gateway facts with m-4 upstream selection/order/fallback policy. Other rejects/adaptations are sound. |
| Catalog C1-C6 | **agree** | Shape import, pinning, facts-only rows, no live refresh/override rows, no policy overlay, and consume-not-rebuild config integrity are correct. |
| Cancellation N1 | **agree with owner seam** | Abort normalization is valid; who requests cancellation and cancel/redeliver sequencing remain m-9/m-3/m-7 concerns. |

### Must-revise findings

#### A1 — P3 reverses the locked authorization order

The artifact says translation happens "after final authorization" (`audit:99`). The locked kickoff says translation, compatibility handling, endpoint binding, and authentication happen after a **pre-translation check**, while the final authorization point is later and no adapter mutation may occur after it (`STEP-3-KICKOFF.md:14-16`). These statements are opposites.

Required revision: P3 must say deterministic adapter translation occurs after the optional/required pre-translation check and **before** the m-3-owned final-wire authorization point. After that final authorization, the bytes, endpoint, credential binding, lane, and adapter selection are immutable through send. Whether m-3 chooses a final-wire-only gate or a pre/post pair remains amendment-owned.

#### A2 — E1 falsely attributes cumulative-partial convergence to opencode

Pi attaches an accumulated `partial` to every nonterminal event (`references/pi/packages/ai/src/types.ts:464-476`). Opencode's text/reasoning lifecycle events carry block id, delta text, and provider metadata, not accumulated partial snapshots (`references/opencode/packages/opencode/src/session/llm/ai-sdk.ts:126-188`). The two donors converge on lifecycle grammar, not on the cumulative-partial field.

Required revision: split E1. Promote/adapt the lifecycle grammar, stable block identity, and typed terminal. Keep cumulative partial snapshots and E3's partial-output representation as an m-8/m-9 DESIGN question coupled to backpressure, redaction, replay, and the open event-persistence decision.

#### A3 — the artifact both settles and asks the event-persistence question

Fresh item 6 says normalized events/usage/attempt records enter the store (`audit:130`), and the boundary contract repeats that write (`audit:168-170`). Question 4 explicitly leaves store-records versus ephemeral events with terminal summary open (`audit:191`). AUDIT cannot simultaneously close and grill the same choice.

Required revision: state only the inherited invariant: **if** any adapter output becomes store-visible, it uses conductor-internal provenance and m-7's serialized commit loop; no side-channel writes. Leave transport-event persistence, summary granularity, and attempt-record shape open for DESIGN/GRILL.

#### A4 — X1 invents m-3's outcome token

`egress_rejected` appears as a required new taxonomy token in X1 and fresh item 5 (`audit:79,129`), but the kickoff specifies a provider-request egress class, a zero-network-send negative, and owner authorship by m-3; it does not assign that token (`STEP-3-KICKOFF.md:13-16,63-66`). The relay-store terminal vocabulary and provider-turn outcome vocabulary are also separate axes.

Required revision: replace the token with "typed provider-request-egress denial outcome; exact token, terminality, record mapping, and disposition are named by the m-3 amendment." Retain zero-send as the invariant. Do not collapse it into `{accepted,rejected,held}` without the owners' mapping.

#### A5 — Y1/Y2/Y4 promote policy before the retry authorization seam exists

Opencode's retry source proves header parsing and a backoff algorithm (`references/opencode/packages/opencode/src/session/retry.ts:35-75`), not a frank policy. A retry is another external send. "Each attempt is observable" is insufficient: each attempt needs a fresh final-wire authorization or an owner-approved immutable replay proof; it also needs exact-lane/endpoint/credential agreement, idempotency semantics, a trusted clock/scheduler, and a bounded outcome if authorization changes between attempts.

Required revision: classify Y1/Y2/Y4 as **adapt-with-governance / fixture seeds**. Preserve context-overflow-never-retry, retry-after variants, SDK-retries-disabled, and header/stall timeout cases as deterministic fixtures. Defer normative budgets, clock behavior, per-attempt authorization/recording, and idempotent replay to the m-3/m-7/m-9 seam and the owner amendments.

#### A6 — P6 smuggles gateway routing policy into an m-8 row

The donor gateway object includes `allow_fallbacks`, `order`, `only`, `ignore`, price/latency sorting, and quantization constraints (`references/pi/packages/ai/src/types.ts:599-684`). Those are routing judgments, not merely wire facts. Even the proposed named-provider form selects an upstream path. m-8 may represent a gateway lane and its factual supported options; m-4 owns upstream selection/order/fallback and exact-lane policy; m-7 binds the effective endpoint.

Required revision: split P6 into factual gateway compatibility/catalog facts (m-8), endpoint binding (m-7), and selection/order/fallback policy (m-4). Keep gateway lanes deferred past the spine and preserve the V2 independently-bound-path negative.

### Required coverage addition

K1/K2 expose opaque encrypted/redacted replay material: pi explicitly stores an opaque encrypted payload in `thinkingSignature` (`references/pi/packages/ai/src/types.ts:333-340`). Add a negative contract/fixture that such bytes are opaque, never interpreted, never exposed in evidence/log/human surfaces, never reused outside their exact provider/model/lane/turn compatibility scope, and are redacted under the m-3/m-7/m-1 policy seams. This is not a credential classification claim; the owners must classify it, but m-8 must not normalize it into an ordinary printable string.

### Agreements worth preserving unchanged

- F1-F4 and the duplicate gate: the provider wire/catalog surface is greenfield; governance substrate is already built.
- Reject caller credentials/headers/env, `onPayload`, plugin mutation, custom fetch, dynamic SDK loading, live catalog refresh, override-bearing catalog rows, silent fallback adapters, routing rankings, and identity spoofing.
- Preserve deterministic tool/reasoning/usage/error wire-reality fixtures, pinned `compat_mode`, compiled/versioned adapters, facts-only catalog rows, and exact-lane/no-fallback negatives.
- Preserve the inert-tool rider and the three unconditional owner-amendment gates.

### Owner-amendment assessment

- **m-3 provider-request egress:** necessity is sharply evidenced, but P3/X1 currently pre-decide its order/token incorrectly. Revise as A1/A4.
- **m-7 credential/trusted config with m-1 review:** necessity and redaction negatives are sharply evidenced. Add opaque-replay-material handling to the consumer-review packet; do not classify it unilaterally.
- **m-4/m-2 exact-lane routing record:** necessity is sharply evidenced. Revise P6 so gateway upstream selection cannot become an m-8 compatibility option.

All three remain hard pre-lock gates under kickoff §6 step 4. No amendment is authored or closed by this AUDIT.

### Duplicate/already-built gate and boundary contract

Duplicate result: **pass**. Fresh inspection remains consistent with no provider HTTP/config implementation in `frank@502e06c`. m-8 must consume the existing guardrail, commit loop, trusted-config integrity mechanism, observe hooks, and qualification hook.

Boundary after revision:

- m-8 writes factual lane-catalog bytes and versioned deterministic adapter/event/fixture specifications.
- m-8 emits normalized transport outputs through m-7-hosted execution; it does not independently write the store. Any selected store-visible projection uses conductor-internal provenance and the serialized commit loop.
- m-8 reads an m-4-selected exact lane, m-7-bound endpoint/credential reference without secret exposure, and m-3 provider-egress verdict; the final-wire authorization binds the exact post-translation request and no mutation follows it.
- m-9 consumes the provider-turn contract; m-4 consumes factual lane IDs; m-3 consumes qualification evidence. No consumer is missing.

### Escalation scan

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — credential/endpoint binding, secret redaction, and opaque replay material are load-bearing owner seams; this AUDIT changes none
- migration/backfill/destructive-write/canonical-data-repair: no — report-only audit; no data or code mutation
- money/inventory/orders/planning/accounting/trust-critical-state: yes — final-wire authorization and the exact-lane binding are trust-critical state
- AI-or-automation-acts-downstream: yes — provider sends and later model-requested tool calls are downstream automation boundaries
- worker/scheduler/queue/retry/async-side-effect: yes — retry, timeout, cancellation, backpressure, and external sends are explicitly audited
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the m-8/m-9/m-3/m-4/m-7 contracts and normalized event surface are shared interfaces
- user-visible-control-with-materializer/downstream-consumer: no — no product mutation; downstream consumers are named but unchanged
- test-runtime-role-mismatch: unknown — the greenfield m-8 implementation and conformance runtime do not yet exist
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — event persistence, retry authorization, gateway policy, and final-wire outcome vocabulary remain owner/GRILL questions; live provider verification is forbidden in AUDIT
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk acceptance is requested by this read-only AUDIT
- decision_scope: none
- residual_risk_accepted: none

The existing medium-tier paired audit plus mandatory owner amendments, consumer reviews, and GRILL path remains required.

### Evidence and actions

Evidence level is E1: exact local source and locked-contract lines were re-read for every must-revise claim. No external call, credential use, code edit, DESIGN authorship, DESIGN lock, PLAN, or implementation occurred.

ACTIONS_GIT_REF: docs-only relay in non-git cwd — wrote `master/relays/step3-audit-m-8/AUDIT-implementer-20260714-213830.md` and appended one `master/relays/INDEX.md` row; `frank/` untouched at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `git -C frank status --short` clean at handoff (empty output) at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

Next requested action: `m-8.planner` revises the audit artifact and returns the exact revised bytes for pair reconciliation. Master/VP should not treat the current matrix as the DESIGN skeleton until A1-A6 are corrected.
