## RECONCILE — APPROVE the m-8 r12 letter rebase to m-10 r28: the r27→r28 D-5 amendment is disjoint from every m-8 seam

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-basis-review-r28
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the bounded F76 basis-addendum review; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no design choice, topology, policy, or owner byte is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260718-065744.md
FROM: m-8.implementer
TO: master.orchestrator-planner
CC: m-8.planner, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260718-070249.md
SUBJECT: APPROVE unchanged m-8 r12 4b670a79... against m-10 r28 4ffaa9ec... — the r28 turn-terminal comparator amendment is m-9-facing and does not move an m-8-consumed locus

VERDICT: approve

Master — I approve the planner's basis addendum at SHA-256 `daf909f3f876b29780773c32e140ef7472cb4b693305e02dc6d583452812817b`.

The two frozen owner hashes are:

- m-8 provider contract r12: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`;
- m-10 app-control-plane contract r28: `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`.

## Delta disposition

**APPROVE — disjoint.** The r27→r28 semantic amendment is confined to m-10 §B.2's m-9↔m-10 turn-level D-5 block (`mvp-ipc-manifest-seam-contract.md:71`):

- `turn_terminal` now consumes `{run_id, turn_id, turn_epoch, terminal}` after dropping the undefined `attempts_summary_ref?`;
- duplicate equivalence is narrowed to `{terminal}` alone and remains total over same-terminal resend versus different-terminal conflict;
- the neighboring confirmation pins `turn_cancel_ack` equivalence to `{partial_disposition}` over `{none, partials_committed_labeled}` and keeps `cancel_point` as separate attempt provenance.

M-8 does not emit, receive, compare, or persist `attempts_summary_ref?`, the D-5 `turn_terminal` duplicate key, or the turn-output `partial_disposition` comparator. Its only turn-terminal references describe the downstream semantic outcome `turn_cancelled`; they do not consume the m-9↔m-10 D-5 frame or equivalence predicate. The r28 delta therefore cannot change an m-8 wire, row, fence, or bootstrap obligation.

## Independent six-locus check at exact r28

1. **Connector bootstrap:** r28 §B.1:65 still carries the seven-field `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest, credential_ref}` exactly; it matches m-8 r12 §5.3:201.
2. **Attempt-open handshake:** r28 §B.1:61 still carries `attempt_open_ok{attempt_id, parked_unknown:[…]}` and `attempt_open_reject{attempt_id, reason ∈ {stale_epoch, invalid_turn, invalid_lease}}`, with commit-before-ack/no-row-reject semantics unchanged.
3. **Attempt-result vocabulary:** r28 §B.1:61 still consumes the closed m-8 set `{sent_completed, denied(…), transport_failed, unknown, rejected_local(…), cancelled(pre_transport|post_invocation)}` at each outcome's own terminal boundary.
4. **Terminal attempt rows:** r28 §B.1:61 and §F:232 still preserve terminal `REJECTED_LOCAL` and terminal `CANCELLED`; deterministic rejects do not park UNKNOWN, and cancelled rows are neither failed nor re-parked.
5. **Cancellation consumer discipline:** r28 §B.1:61 and §F:232,237 still make the notification one-way, keep raw closure/crash at `UNKNOWN_PROVIDER_OUTCOME`, key duplicate equivalence on `{attempt_id, reported turn_epoch, cancel_point}`, and keep `cancellation_id` provenance/correlation-only. This matches m-8 r12 §1.4:101-103 and fixture 17b:230.
6. **Epoch authority:** r28 §B.4:104,110-112 still makes m-8 generation-blind, advances its authority only from m-10 `epoch_update`, maps below-current to `STALE_EPOCH`, and maps above-current to held/retriable `EPOCH_AHEAD` plus CTRL-C query. This matches m-8 r12 §1/fixtures at :80,86,218,231,244.

The exact r27 approval at `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7` establishes those attempt/cancellation surfaces as the accepted pre-amendment basis. The r28 owner relay and pair review identify the bounded two-locus D-5 amendment, and the independent current-byte scan above confirms that its fields and comparator are outside all six m-8 loci.

## Gate disposition

The r12 owner-byte approval `DESIGN-REVIEW-implementer-20260718-043932.md` remains valid without replay. No m-8 byte moves and no new m-8 hash is created. This review satisfies the F76 implementer-review condition for the letter-level m-10 basis rebase.

Approval advances only to Master's corrected stage-1/2/3 close supplement and fresh VP close-confirm. It does not grant interface lock, stage-4/5 dispatch, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, or deploy authority.

## Verification

- Incoming direct-address relay SHA-256 recomputed: `efd0719b54a8cc82ecfd6afd8b25c2d1bbb8a293e67395e1f2a189cde7c8f6ec`.
- Planner basis addendum SHA-256 recomputed: `daf909f3f876b29780773c32e140ef7472cb4b693305e02dc6d583452812817b`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Frozen m-10 r28 SHA-256 recomputed: `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`.
- Exact incoming, addendum, m-10 r28 DESIGN, and m-10 r28 DESIGN-REVIEW relays lint clean.
- `git -C frank status --short` returned empty; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace disk action only — created this bounded review relay and appended one `master/relays/INDEX.md` row; no owner design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260718-070249.md`.
Next requested action: Master binds this approval into the corrected stage-1/2/3 close supplement and returns that supplement for fresh VP close-confirm; keep frozen owner bytes and every later authority gate held.
