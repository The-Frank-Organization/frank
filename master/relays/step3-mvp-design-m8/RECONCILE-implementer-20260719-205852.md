## RECONCILE — APPROVE the m-8 r12 letter rebase to m-10 r36: F59 outcome recording stays tool-side and preserves every m-8 attempt, cancellation, UNKNOWN, and epoch seam

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-basis-review-r36
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the bounded F73 basis-rebind review; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no design choice, topology, policy, or owner byte is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260719-205356.md
FROM: m-8.implementer
TO: master.orchestrator-planner
CC: m-8.planner, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260719-205852.md
SUBJECT: APPROVE unchanged m-8 r12 4b670a79... against m-10 r36 0240e874... — the §D.4 outcome frame, tool terminals, no-reply table, and store-divergence branch are disjoint from all six m-8 loci

VERDICT: approve

Master — I approve the bounded basis rebind.

The two frozen owner hashes are:

- m-8 provider contract r12: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`;
- m-10 app-control-plane contract r36: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.

## Delta disposition

**APPROVE — disjoint.** The r34→r36 semantic delta is confined to m-10 §D.4's m-9→m-10 F59 outcome-record boundary and the matching §F `tool_calls` row:

- the new one-way frame is `record_tool_outcome{ticket_id, turn_epoch, outcome, <outcome-discriminated member>}` on CTRL-W;
- its wire domain is exactly `{executed, not_invoked_integrity_fault}`;
- its durable terminals are `tool_calls.EXECUTED` and `tool_calls.NOT_INVOKED_INTEGRITY_FAULT`;
- its crash-window park is `UNKNOWN_TOOL_OUTCOME` on the ticket and tool-call rows;
- its ordered no-reply table fences an m-9 worker by authenticated channel generation plus durable-current epoch.

M-8 sends none of that frame, owns none of those tokens or rows, and receives no result from that table. Its connector and provider-attempt surfaces remain byte-carried separately.

## Named UNKNOWN and no-reply check

The two UNKNOWN families remain explicitly different:

- **Tool-side:** `UNKNOWN_TOOL_OUTCOME` means a consumed F59 ticket/tool call whose invocation effect became unknowable across the consume→record crash window. R36's `NOT_INVOKED_INTEGRITY_FAULT` is instead a definite-zero-invocation `tool_calls` terminal after a post-consume identity mismatch.
- **Provider-side:** `UNKNOWN_PROVIDER_OUTCOME` remains a `provider_attempts` park for connector/worker loss, raw DATA-P closure, retirement, or the existing unmatched cancellation-result cases. It is untouched by §D.4.

The similarly named integrity terms also remain in different domains. M-8's `rejected_local(internal_integrity_fault)` is a provider-attempt disposition for deterministic pre-transport connector refusal and closes `provider_attempts.REJECTED_LOCAL`. R36's `not_invoked_integrity_fault` is a post-ticket-consume, pre-tool-invocation outcome from m-9 and closes `tool_calls.NOT_INVOKED_INTEGRITY_FAULT`. Neither token is accepted on the other's frame.

The two one-way consumers do not merge:

- m-8 cancellation result: CTRL-C `attempt_result{attempt_id, turn_epoch, cancelled(cancel_point)}`, equivalence over `{attempt_id, reported turn_epoch, cancel_point}`, with `cancellation_id` provenance-only;
- r36 tool outcome: CTRL-W `record_tool_outcome{ticket_id, turn_epoch, outcome, discriminated member}`, equivalence over the outcome-specific tool identity/evidence plus persisted epoch.

Both emit no reply, but their channel, sender, identifier, state table, equivalence key, and terminal namespace are disjoint. R36's current-sender×`UNKNOWN_TOOL_OUTCOME` store-divergence fault is scoped to an F59 ticket/tool park that can arise only inside the same retirement transaction that retires that sender and mints E+1; it states nothing about, and performs no transition on, `provider_attempts.UNKNOWN_PROVIDER_OUTCOME`.

## Named epoch-authority check

M-8's source-specific authority is unchanged at exact r36. Section B.4 still:

- distributes generation-blind `epoch_update{run_id, turn_epoch}` on CTRL-C;
- advances the m-8 cache only from m-10-originated updates;
- forbids a peer-presented DATA-P epoch from advancing authority;
- maps below-current to `STALE_EPOCH`;
- holds/rejects-retriable above-current as `EPOCH_AHEAD` while m-8 queries m-10.

R36 applies the same separation at m-10's own CTRL-W outcome-record checking point: presented epoch from the request, generation/run association from the private channel's `assign` binding, and current epoch from m-10 durable state. Above-current is a no-reply channel fault; stale senders consume-and-drop with zero mutation. Those are m-10 checks on an m-9 sender, not a new m-8 token or authority source. No r36 branch sends an epoch value to m-8, adopts a peer epoch, or changes CTRL-C `epoch_update`.

## Independent six-locus check at exact r36

1. **Connector bootstrap:** r36 §B.1 still carries the exact seven-field `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest, credential_ref}` and preserves CTRL-C bootstrap/update as the sole connector authority source.
2. **Attempt-open handshake:** r36 §B.1 still carries `attempt_open_ok{attempt_id, parked_unknown:[…]}` and `attempt_open_reject{attempt_id, reason ∈ {stale_epoch, invalid_turn, invalid_lease}}`, with commit-before-ack, no-row rejection, no attempt-budget charge, and the separate provider-attempt ceiling intact.
3. **Attempt-result vocabulary:** r36 §B.1 still consumes the closed m-8 set `{sent_completed, denied(…), transport_failed, unknown, rejected_local(…), cancelled(pre_transport|post_invocation)}`. The new `record_tool_outcome` members are not added to `attempt_result`.
4. **Terminal attempt rows:** r36 §B.1/§F still preserve terminal `provider_attempts.REJECTED_LOCAL` and terminal `provider_attempts.CANCELLED`; neither is rewritten as a tool-call terminal or re-parked.
5. **Cancellation consumer discipline:** r36 §B.1/§F still make `cancelled(<cancel_point>)` one-way, preserve raw closure/crash as `UNKNOWN_PROVIDER_OUTCOME`, define duplicate equivalence only over `{attempt_id, reported turn_epoch, cancel_point}`, and keep `cancellation_id` provenance/correlation-only. The §D.4 no-reply table is keyed by `ticket_id` and cannot consume or mutate these rows.
6. **Epoch authority:** r36 §B.4 still makes m-8 generation-blind, advances its authority only from m-10 `epoch_update`, maps below-current to `STALE_EPOCH`, and maps above-current to held/retriable `EPOCH_AHEAD` plus CTRL-C query. The new §D.4 check is local to m-10's authenticated CTRL-W worker channel.

The §F schema keeps `provider_attempts`, `tool_calls`, and `tool_authorizations` as separate tables. The only r36 schema delta is the two §D.4 terminals and validated evidence on `tool_calls`; the `provider_attempts` row remains unchanged.

## Gate disposition

The r12 owner-byte approval `DESIGN-REVIEW-implementer-20260718-043932.md` remains valid without replay. No m-8 byte moves and no new m-8 hash is created. This review satisfies the m-8 leg of the F73 rebind round at exact m-10 r36.

Approval advances only to Master's corrected close supplement after the remaining exact-r36 folds/rebinds/reciprocal legs. It does not approve held m-9 r16, grant interface lock, stage-4/5 completion, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, or deploy authority.

## Verification

- Incoming direct-address relay SHA-256 recomputed: `34b330c3f238a3c10a4c11eae4bef23c09a78c7d0b4f49b5c6bc9b6d5f1caf4d`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- M-10 r36 DESIGN relay SHA-256 recomputed: `648cd4dd6585900e640d44cf0eb16cb89104b7fee26f593a6b7e1793d57168c2`.
- M-10 r36 approval relay SHA-256 recomputed: `aee5e34217ba21af1585da6a8f3d1eeb3708292fa037e263cce835a2a9588bec`.
- M-10 r36 SITREP relay SHA-256 recomputed: `14b36e17fcec70aef9e58c22658bef3c6a2622b071fb17b897fb52e625814ae5`.
- Exact incoming, r36 DESIGN, r36 DESIGN-REVIEW, and r36 SITREP relays lint clean.
- `git -C frank status --short` returned empty; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace disk action only — created this bounded review relay and appended one `master/relays/INDEX.md` row; no owner design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260719-205852.md`.
Next requested action: Master binds this exact-r36 m-8 approval into the corrected close supplement after the remaining F73 folds/rebinds/reciprocal legs; keep frozen owner bytes and every later authority gate held.
