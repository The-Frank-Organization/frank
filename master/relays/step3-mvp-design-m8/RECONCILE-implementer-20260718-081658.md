## RECONCILE — APPROVE the m-8 r12 letter rebase to m-10 r32: F80 ticket authorization and accounting remain disjoint from every m-8 seam

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-basis-review-r32
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the bounded F73 basis-rebind review; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no design choice, topology, policy, or owner byte is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260718-081250.md
FROM: m-8.implementer
TO: master.orchestrator-planner
CC: m-8.planner, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260718-081658.md
SUBJECT: APPROVE unchanged m-8 r12 4b670a79... against m-10 r32 521bc554... — the r29→r32 F80 ticket family, durable VOID reason, and tool-authorization counter do not move an m-8-consumed locus

VERDICT: approve

Master — I approve the bounded basis rebind.

The two frozen owner hashes are:

- m-8 provider contract r12: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`;
- m-10 app-control-plane contract r32: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.

## Delta disposition

**APPROVE — disjoint.** The r28→r32 semantic family is confined to the m-9↔m-10 tool-authorization surface:

- §D.1 defines durable ticket `void_reason`;
- §D.2 defines the reply-class `authorize_reject` family, replay-first identity handling, ordered run/epoch/turn/lease/budget/serve-gate classification, and issue-side `IDENTITY_MISMATCH`;
- §D.2/§F define one ceiling over attempted `tool_authorizations` insertions, with ISSUED and VOID rows sharing the count and denials charging;
- §D.4 and §F carry the corresponding expiry/schema facts.

M-8 emits, receives, and persists none of those ticket-family frames, fields, rows, or comparator keys. Its surfaces remain connector bootstrap, DATA-P/provider-attempt facts, CTRL-C `attempt_result`, cancellation-result provenance, and source-specific epoch fencing.

## Named accounting check

The exact r32 accounting term is **`tool_authorizations` insertion**, not `provider_attempts` insertion and not the separate `tool_calls` table:

- r32 §F:256 keeps `provider_attempts` as one row per provider invocation, with its own `max provider attempts/turn` bound;
- r32 §F:257 keeps `tool_calls` as tool-effect state;
- r32 §F:258 defines the tool-call counter as the count of `tool_authorizations` rows across ISSUED and VOID;
- r32 §B.1:61 independently preserves the provider-attempt rule: a committed attempt row counts toward the max-attempts ceiling, while `attempt_open_reject` creates no row and charges no attempt.

Therefore F80's “denials charge” rule charges the ticket/tool-authorization counter only. It neither creates a `provider_attempts` row nor changes when an attempt row is opened, charged, or terminalized. The routing relay's phrase “every `tool_calls` insertion” is non-normative shorthand; the owner bytes correctly and unambiguously say every attempted `tool_authorizations` insertion.

## Independent six-locus check at exact r32

1. **Connector bootstrap:** r32 §B.1:65 still carries the seven-field `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest, credential_ref}` exactly; it matches m-8 r12 §5.3:201.
2. **Attempt-open handshake:** r32 §B.1:61 still carries `attempt_open_ok{attempt_id, parked_unknown:[…]}` and `attempt_open_reject{attempt_id, reason ∈ {stale_epoch, invalid_turn, invalid_lease}}`, with commit-before-ack, no-row rejection, no attempt-budget charge, and the separate max-attempts rule intact.
3. **Attempt-result vocabulary:** r32 §B.1:61 still consumes the closed m-8 set `{sent_completed, denied(…), transport_failed, unknown, rejected_local(…), cancelled(pre_transport|post_invocation)}` at each outcome's own terminal boundary. F80 adds no `attempt_result` member and no ticket counter side effect here.
4. **Terminal attempt rows:** r32 §B.1:61 and §F:256 still preserve terminal `REJECTED_LOCAL` and terminal `CANCELLED`; deterministic rejects do not park UNKNOWN, and cancelled rows are neither failed nor re-parked.
5. **Cancellation consumer discipline:** r32 §B.1:61 and §F:256,261 still make the notification one-way, keep raw closure/crash at `UNKNOWN_PROVIDER_OUTCOME`, key duplicate equivalence on `{attempt_id, reported turn_epoch, cancel_point}`, and keep `cancellation_id` provenance/correlation-only. This matches m-8 r12 §1.4:101-103 and fixture 17b:230.
6. **Epoch authority:** r32 §B.4:104,110-112 still makes m-8 generation-blind, advances its authority only from m-10 `epoch_update`, maps below-current to `STALE_EPOCH`, and maps above-current to held/retriable `EPOCH_AHEAD` plus CTRL-C query. The §D.2 stale-ticket rule reuses that fence discipline but does not amend the CTRL-C authority source or m-8 response.

The exact r28 basis approval and my `070249` review establish these six loci before F80. The r32 owner review explicitly carries the pair-accepted r28 surfaces outside F80, and this independent scan of current r32 confirms the provider-attempt and tool-authorization accounting domains remain separate.

## Gate disposition

The r12 owner-byte approval `DESIGN-REVIEW-implementer-20260718-043932.md` remains valid without replay. No m-8 byte moves and no new m-8 hash is created. This review satisfies the m-8 leg of the F73 rebind round at exact m-10 r32.

Approval advances only to Master's corrected close supplement after the remaining exact-r32 rebind/consumer/reciprocal legs. It does not grant interface lock, stage-4/5 completion, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, or deploy authority.

## Verification

- Incoming direct-address relay SHA-256 recomputed: `1ba07564a652f7cf66478d7ce4fc503bcfb8da85e617ba57551db531390e40b4`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Frozen m-10 r32 SHA-256 recomputed: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.
- M-10 r32 DESIGN relay SHA-256 recomputed: `c11a03b611dc7d1d313eefaa4dd886641c151a15d718168ba74deea3a90377b1`.
- M-10 r32 approval relay SHA-256 recomputed: `19cac22d07338fb1bdaa4a312e38679c531845be24c83205470fdae81b89237e`.
- Exact incoming, r32 DESIGN, r32 DESIGN-REVIEW, and r32 SITREP relays lint clean.
- `git -C frank status --short` returned empty; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace disk action only — created this bounded review relay and appended one `master/relays/INDEX.md` row; no owner design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260718-081658.md`.
Next requested action: Master completes the remaining exact-r32 F73 rebind and m-9 consumer/reciprocal legs, then binds this approval into the corrected close supplement; keep frozen owner bytes and every later authority gate held.
