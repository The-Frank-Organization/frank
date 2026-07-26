# m-8 — Provider Adapters

**Division II — Harness Runtime. Stood up Step-3 (2026-07-14). Greenfield: DESIGN-only until the domain locks.** Charter authorized by the VP-co-signed `STEP-3-KICKOFF.md` (`step3-prep/RECONCILE-orchestrator-reviewer-20260714-222000`, approve).

## Engagement
The provider-abstraction layer for frank's own model-agnostic runtime — the pair that lets frank **drive models directly** instead of riding on Claude Code / Codex. Design-lead `.planner` + adversarial design-reviewer `.implementer`, independent operator-relayed sessions.

## Scope (owns)
- **The frank-owned normative provider contract** — request · normalized event · tool-call · reasoning-replay · usage · finish/error · cancellation · retry/idempotency · timeout/backpressure · partial-stream semantics. pi/opencode are **prior art + conformance fixtures**, NOT the spec (kickoff §1).
- **Provider wire translation + normalized provider events** — per-provider adapters that speak each provider's API and normalize to frank's contract, including the `reasoning_content` thinking-token handling (RUNTIME-RESEARCH §6.2).
- **The factual lane catalog** (single writer) — the `{model_id, provider_id, serving_profile_id, compat_mode}`-keyed table + spec-sheet payload (context_window, max_output, reasoning{supported, effort_levels, emits_reasoning_content}, tool_use, modalities, streaming, cost, latency_tier, `source: seeded|generated`). Seeded now, models.dev-shaped for later auto-generation; pinned per kickoff §4 (schema/version · source+effective-time · canonical digest · immutable lane ID). **No secrets in catalog bytes.**
- **Provider conformance fixtures** — the deterministic E2 suite every adapter must pass.

## Boundaries (does NOT own)
- **Credentials · endpoint selection · egress policy · authority enforcement · routing judgment.** Those stay with m-7 (trusted config), m-3 (egress/qualification), m-5/m-7 (authority), m-4 (routing).
- The catalog is **facts only** — m-4 owns the **policy overlay** (capability priors, route selection) keyed TO m-8 lane IDs; m-4 never writes m-8 rows, and m-8 never writes routing policy (kickoff §3).

## Consumer-lock seams (interface-lock-first; Master+VP lock the joins)
- **m-8↔m-3** — the **provider-request-egress amendment (m-3-authored, m-7-hosted)** + qualification evidence over lanes. m-8 is a *consumer reviewer*, not the author.
- **m-8↔m-7** — the **trusted-config/credential amendment (m-7-authored)** + the egress *host* substrate. m-8 is a *consumer reviewer*, not the author.
- **m-8↔m-1** — the secret boundary (credentials never enter catalog/snapshots/records/seat surfaces/evidence).
- **m-8↔m-4** — the factual-catalog ↔ policy-overlay contract (m-4 keys to m-8 lane IDs).
- **m-8↔m-9** — the adapter the runtime calls.

## Status + the pre-build design sequence (kickoff §6)
**DESIGN (design-only) — current** (AUDIT discharged `step3-audit-m-8`; PROCEED-TO-DESIGN issued `step3-design-m-8/DESIGN-orchestrator-planner-20260715-005500`, GRILL_REQUIRED: yes, at VP transition approval `…-005000`). Stages 2–4 author in PARALLEL per the kickoff §6 amendment (B14) — parallel authoring ≠ parallel locking; no m-8 lock until all named reviews + consumer confirms close:
1. **AUDIT** — audit/promotion matrix against pi, opencode, the landed frank interfaces, and the locked m-x contracts (may run concurrent with m-9's audit).
2. **DESIGN** — `.planner` authors; `.implementer` adversarial DESIGN-REVIEW.
3. **GRILL** — durable grill + `GRILL_LOCK_ID` on each hard-to-reverse cross-domain design (the live operator conversation is decision-input, NOT a substitute).
4. **OWNER AMENDMENTS + CONSUMER REVIEW** — m-8 consumer-reviews the **m-3-authored provider-request-egress amendment (m-7-hosted)** + the **m-7 credential amendment** + the **m-4/m-2 routing-record amendment**; all three must **close** before lock (parallel authoring per kickoff §6 amendment; parallel authoring ≠ parallel locking).
5. **RECONCILE + LOCK** — Master+VP reconcile + lock this charter's spec-of-record.
6. Only then does m-8 act as **PM** to a T4 build team (vertical-first per §5).

**No design-lock, PLAN, T4 code token, implementation, credential use, or external provider call is authorized by this charter** — those remain where the kickoff + standing protocol place them.

## Step-3 reframe delta (RATIFIED 2026-07-15 — `master/STEP-3-ARCH-AMENDMENT.md` @ `2d240eb6…`)
m-8 is the **app-side connector** — a **separate trusted process from the m-9 worker BEFORE the first E3** (same host OK, not same credential-readable address space); **NOT a conductor seat**. It **holds the provider credentials** at runtime under m-1's boundary, does the provider wire via a frank-owned client, and is the **last pre-wire enforcement host** (`freeze → authorize → attach → send`; MVP default **one attempt, no auto-retry**). m-8 **OWNS the connector-credential contract** (authors it; m-1 reviews the secret boundary; m-10/m-3 consume). The four-axis catalog **schema** + conformance fixtures survive; `step3-design-m-8` re-dispatches app-side. Provider egress **bypasses the conductor**.

**Supersedes (explicit, VP F20):** this delta SUPERSEDES the pre-reframe **§Boundaries** line ("Credentials · endpoint selection · egress policy · authority enforcement … stay with m-7 …") and the **§Consumer-lock-seams** lines that make the egress/credential amendment **m-3-authored/m-7-hosted** and the **credential amendment m-7-authored** — under the reframe credentials are **m-8-owned, app-side, m-1-governed**, and provider egress is **app-side** (mechanism), with **policy staying m-3**. The **§Status "DESIGN — current"** lane is **HELD** (`step3-hold-m8`); m-8 **re-dispatches app-side at STAGE 2**, only after the Master+VP first-stage interface-lock (m-10 + m-5). The old sections above are retained as record, **not operative**.

## Current status (as of 2026-07-25) — stage-6 Lane-2 (interface DAG): m-8's producer obligations COMPLETE
**Lane-2 producer delta pair-approved byte-bound at r5 `c0b7b488…`** (integrated into the frozen stage-2 provider contract; `step3-relock-dag-m8/` thread) **+ the decoded 2a/2b discriminator pair-approved byte-bound at r7 `734e44b7…`** — `design/2026-07-22-stage6-BE-digests-addendum.md` (`step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260724-174500.md`, zero open findings; supersession chain r5→r6→r7). r7 publishes **`refusal_stage ∈ {pre_freeze, post_freeze}`** on every `m8.dataP_reply.v2` reject reply (mirrored on `m8.attempt_result.v2`), letting consumers decode the row-2a/row-2b `internal_integrity_fault` ambiguity (`refusal_stage = post_freeze ⟺ B∧E present`) **without consulting digest presence** — closing the circularity m-3's sink needed resolved, and offered as the producer fact for m-10's carriage-row B-presence fold. **Nothing further owed from m-8 on lane-2 unless a consumer surfaces a new producer gap** (`step3-relock-dag-m8/SITREP-planner-20260724-181500.md`).

**Frozen provider contract: r12 `4b670a79…`** (`design/2026-07-17-mvp-provider-contract.md`) — UNMOVED throughout the lane-2 fold; the B/E delta and the discriminator are bounded additives over it, touching no frozen r12 byte.

**Consumers bound to m-8's lane-2 bytes:** m-3's E0/E3 sink (r19 `92e08d09…`) and m-10's B/E carriage row (rev3 `cd17db32…`) both key off r5/r7. m-9's §5-E `input_item` tripwire was checked clear against r12 (no field/item-kind/wire-transform change triggered).

**m-8 holds report-only** pending the §D-settlement amendment's operator ratification and the integrated re-lock; no DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action claimed. H-12 (no sandbox) continues to hard-block external use.

## Step-3 MVP amendment delta (RATIFIED 2026-07-16 — `master/STEP-3-MVP-AMENDMENT.md` r7 @ `2f75f2a1…`; the F44 fold)
The reframe delta above stands, with these amendments now operative: **(F57)** the "not same credential-readable address space" phrase is superseded — separation = **accidental-disclosure reduction + non-injection into the enumerated surfaces**, with same-user inspection of peer-process state an explicit unsandboxed MVP residual (no hard unreadability claim; the OS boundary is Step-4 H-12) + the §2 minimum hardening list (sanitized env, close-on-exec, private runtime dir, no inherited ambient creds). **(§0/§2)** branch (a): m-8 = the minimum OpenAI-compatible **API-key** client (OAuth/subscription → Step-4); a **separate supervised connector process** (grill #1) — the single egress chokepoint for the governed provider attempt. **(§2a)** "one attempt, no auto-retry" = **per provider INVOCATION** (covers SDK/HTTP/middleware/Retry-After/stream-reconnect/failover; a user retry = a NEW `attempt_id`); a turn may hold multiple recorded attempts. **(§1)** `freeze → authorize → attach → send` sharpened: hash the **frozen core**, credential-attach + deterministic wire encoding the only post-freeze transforms. **(§3)** m-8 is the **provider-report source** — self-reported, worker-carried E0 ("attestation" avoided); the deny→zero-send negatives are instrumented-test evidence; live E3 = a separate observer, bound to `m-8_build_digest` among the F63 release digests (produced at the post-build RELEASE-BINDING event). **(§7)** m-8 = **stage 2**: authors the provider/credential/wire bytes consuming m-1's secret-boundary + m-3's egress/E0 + m-10's IPC contracts, with m-9's consumer review + m-8.implementer final-fold review; the packet `:94-98` F11–F13 carries remain OPEN with the owner. Dispatch follows stage-1 artifacts — not yet issued at this fold.
