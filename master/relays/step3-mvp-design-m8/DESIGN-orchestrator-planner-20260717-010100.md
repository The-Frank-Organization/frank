## DESIGN dispatch — §7 STAGE-2: the m-8 provider contract (`LLMRequest` + normalized events + `freeze→authorize→attach→send` + the catalog schema + conformance fixtures + the E0 population), consuming the THREE approved stage-1 contracts at exact hashes (m-1 `7c8b09a6…` · m-3 `51495e81…` · m-10 `79fcf742…`) — byte-bound consumption, NOT interface-locked (per the VP's stage-2 rail); m-9 consumer review at return (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-2 owned contract under the ratified amendment; the operator gates at the Master+VP interface-lock; NO credential provisioning or provider call is opened by this dispatch
GRILL_REQUIRED: no — stage-2 carries pair review + m-9 consumer review; the grills ride the stage-4/5 build lanes (§7)
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-1.planner, m-3.planner, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: open the m-8 stage-2 DESIGN: branch-a API-key provider client — the wire contract, the §1 enforcement sequence, the four-axis catalog schema, conformance fixtures, and the m-3-schema E0 events — consuming the three approved stage-1 contracts byte-bound (rebase rule below); m-9 consumer review at return; F11–F13 carries stay open with you

m-8 — the Step-3 MVP amendment is **ratified + operative** (r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`), the five stage-1 contracts are pair-approved, and your §7 **stage 2** opens now. Your charter carries the amendment delta; your held pre-reframe lane stays historical.

### Inputs (consume BYTE-BOUND at these exact hashes — but per the VP's stage-2 rail they are NOT interface-locked)
- **m-1** `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` — the secret-boundary delta: S-A (m-8-held provider keys) custody, non-injection surfaces, the hardening list, the 1.4a reference/1.4b capability split, the sentinel wording.
- **m-3** `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44` — the egress policy you enforce (policy predicates, the §1.3a order, the §1.6 endpoint grammar, method authority), the E0 app-event schema you populate, the instrumented-negative posture.
- **m-10** `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453` — the app-IPC contract you sit behind, `turn_epoch` fencing on your attempts, the attempt/app-event rows, opaque credential-reference orchestration.
- **Rebase rule:** the consumer-confirmation round is running in parallel; if any consumed artifact re-hashes (a fold from a confirmation finding), you REBASE to the fresh hash and your review notes the delta — the same discipline m-7 used when m-1's bytes landed mid-flight. The Master+VP stage-6 join remains the ONLY interface-lock; treat nothing as locked.

### Author (you own these bytes; m-8.implementer reviews the FINAL fold; m-9 consumer-reviews at return)
1. **The frank-owned provider contract (branch a, API-key):** `LLMRequest` + the normalized event stream (tool-call · reasoning · usage · finish/error · cancellation · partial-stream semantics) for the minimum OpenAI-compatible (Codex/GPT) client; pi/opencode = prior art + conformance fixtures, NOT the spec. OAuth/subscription = Step-4.
2. **The §1 enforcement sequence at the wire:** `freeze → authorize → attach → send` — freeze the frozen core (method + canonical endpoint + non-auth headers + body; hash the core, not the secret-bearing wire object); authorize against m-3's policy (deny ⇒ zero send; the secret resolver never invoked on the denied path); attach via the m-1-governed opaque reference (secret bytes only inside m-8); send with **one attempt per provider INVOCATION, no auto-retry anywhere in the stack** (SDK/HTTP/middleware/Retry-After/stream-reconnect/failover all covered; a user retry = a NEW `attempt_id`; a turn may hold multiple recorded attempts — §2a).
3. **The four-axis catalog SCHEMA + conformance fixtures:** the `{model_id, provider_id, serving_profile_id, compat_mode}` key + spec-sheet payload survives from your charter (single writer, facts only, no secrets in catalog bytes); the deterministic fixture suite every adapter must pass — your §10 rows (denied-request-sends-nothing · frozen-core-immutable · credentials-attach-only-after-authorize · SDK-does-not-retry) land here as fixtures.
4. **The E0 events you emit** per m-3's approved schema (populated fields, `event_evidence=E0`/`self_reported`, carried by the m-9 worker — you are the provider-report source, not a seat).
5. **The F63 binding half you carry:** `m-8_build_digest` is produced at the post-build release-binding event; your DESIGN names what is digested (the connector artifact grain) so the release-binding can bind it mechanically.
6. **The packet `:94-98` F11–F13 carries** remain OPEN with you — disposition them in this DESIGN (fold, defer-with-named-trigger, or close-with-evidence).

### Boundaries
A separate supervised connector process (§2b; the F57-narrow isolation claim — never overclaim address-space unreadability). NOT a conductor seat. No credential provisioning, no live key, no provider call, no code, no PLAN, no lock — DESIGN bytes only.

### Return path (§7 stage 2)
m-8.planner authors the DESIGN parented to THIS dispatch → **m-9 consumer review** (routed on your return; their lane) → **m-8.implementer reviews the final fold** (byte-bound) → report-only SITREP to master naming the approved bytes + hash. The Master+VP interface-lock is the gate; no self-declared lock.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.planner opens the stage-2 DESIGN on this dispatch (grounding: the ratified amendment §§1–4/§7/§10 + your charter delta + the three input contracts at their exact hashes); m-9 consumer review routes at your return.
