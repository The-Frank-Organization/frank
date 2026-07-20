## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r0 must revise: broker CI contract, release vector, restartable DATA-P, and exact counter encoding

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r1
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings are contract-completeness defects under already-ratified and pair-approved inputs; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this stage-1 review does not reopen the operator-locked topology/F59/F60 decisions; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-070900.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-072454.md
SUBJECT: must-revise exact b971ea2f... - four load-bearing interface defects block pair approval; fresh bytes/hash and a uniquely-parented re-review are required

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact proposed bytes at SHA-256 `b971ea2f342f950c73808903ea12378c7f024099cdd9270d571cf0943bbdf880`. Routing, lineage, `DESIGN_DOC_ID`, the r7 basis hash, and the requested byte binding pass. The design is not pair-approvable yet: four interface defects remain.

## Findings

### R1 - The claimed F64 integration omits and contradicts the pair-approved m-7 CI-1/CI-2/CI-3 contract

The m-10 draft reduces the broker edge to an ordered/queryable `epoch_state` feed and says every checker may adopt any epoch above its cached value (`2026-07-16-mvp-ipc-manifest-seam-contract.md:73-81`). It also gives the worker only `assign{run_id, turn_epoch, manifest_digest}` (`:53`), applies socketpair-before-spawn to "each channel" (`:17-23`), ignores unknown fields in every known message (`:33-35`), and has no broker-control/event/transition rows in its store census (`:163-180`).

Those bytes do not consume the already pair-approved m-7 r6 contract at `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`:

- m-7 makes m-10's feed the ONLY epoch authority; the broker never derives, guesses, defaults, or adopts a worker-presented epoch (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:115-121`). The global "adopt-and-record above" rule therefore cannot govern DATA-P or broker requests: a replaceable worker must not be able to advance a checking point's authority cache merely by presenting a larger integer.
- Epoch change is a durable identified handshake, not publish-and-eventually-see: `PROPOSED -> PREPARING -> CROSSERS_DURABLE -> INSTALLED | ABORTED`, with a frozen crossing set, m-10 durable rows before ack, no return to E after E+1 publication, and transition-ID recovery (`:123-149`). The current draft supplies none of m-10's half.
- m-7 pins a broker-owned dial-in control listener, durable `broker_control` token/generation + controller lock/adoption lifecycle, expanded `assign` fields, broker event dedup, crossing rows, and the transition ledger as the three named m-10 deltas (`:181-230,309-313`). The draft contains none of CI-1/CI-3 and omits `generation_id` plus `broker_worker_endpoint` from `assign` (CI-2).
- m-7's broker-event schema is closed and requires m-10 to reject unknown fields (`:210-230`), which conflicts with the draft's channel-global ignore-unknown-field rule.

Required return: consume m-7 r6's CI-1/CI-2/CI-3 and barrier/ack ordering in full; make epoch authority source-specific (only m-10's authenticated control/feed may advance it); scope unknown-field evolution per message family so the closed broker-event family stays closed; add the exact durable rows and recovery transactions; then request a fresh review of new bytes. Until then m-10 cannot issue the required exact-byte consumer confirmation to m-7.

### R2 - The manifest cannot produce the pair-approved m-3/F63 applicability vector

The manifest schema carries `policy_source_ref`, one ambiguous `provider_lane.catalog_digest`, and `release_binding: {…}` (`2026-07-16-mvp-ipc-manifest-seam-contract.md:85-100`). It later says `tool_catalog_digest` is verified, but never gives that digest a manifest/release-binding field, and the serve gate says only ``catalog_digest`/`release_binding`` (`:101-115`). This leaves the m-8 factual lane-catalog digest and m-9 tool-catalog digest conflated. It also omits `policy_digest` entirely.

The pair-approved m-3 contract requires the run manifest to carry `policy_digest` at freeze (`master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md:38,82,114`), requires the release-binding record to supply build digests XOR `release_digest` plus the separate `tool_catalog_digest`, and sources the run/turn/attempt vector from the frozen manifest plus that exact binding (`:156-188`). Its consumer table explicitly assigns m-10 the `policy_digest` + `run_manifest_digest` freeze seam and pinned-lane equality (`:251-260`). Ratified r7 separately names `tool_catalog_digest` as the m-9-produced digest mechanically verified at F63 (`master/STEP-3-MVP-AMENDMENT.md:49-50,59`).

Required return: replace `{…}` with an exact closed `release_binding` schema and validity rule; carry `tool_catalog_digest` distinctly from the m-8 lane-catalog digest; add the m-3-produced `policy_digest` and freeze-time pinned-lane equality; pin the build-digest XOR covering-release form (including the shared-client applicability rule); and define which exact bytes/records C.3 compares. `policy_source_ref` may remain provenance but cannot substitute for `policy_digest`.

### R3 - The direct DATA-P channel cannot survive the independent restart behavior the same document promises

The draft creates DATA-P as a socketpair whose two endpoints are inherited only when the children spawn, then requires m-10 to retain no endpoint (`2026-07-16-mvp-ipc-manifest-seam-contract.md:17-23,186-193`). It separately says the connector is independently restartable and, on connector failure, m-10 restarts it without automatically retrying the parked attempt (`:47-70`). The symmetric worker-generation replacement is also required (`:47-81`).

After either child dies, the surviving child owns only its endpoint of the dead pair; m-10 owns no DATA-P endpoint to hand to the replacement. The framed-byte-only as-if-process-separated rule defines no descriptor-passing or rendezvous mechanism. Therefore "restart connector" and "replace worker generation" cannot re-establish DATA-P without an unstated co-restart.

Required return: pin one complete lifecycle that preserves the no-hot-path/no-payload-at-m-10 boundary - for example, co-restart both children under the new epoch and park the prior attempt, or define a private listener/SCM_RIGHTS-style endpoint handoff with exact ownership, close-on-exec, stale-endpoint fencing, and crash recovery. Also name the payload-free attempt-start/outcome control messages that let the sole m-10 writer truthfully maintain `provider_attempts` and `pending_app_events`; direct DATA-P alone gives m-10 no observation path for those rows.

### R4 - JCS frames cannot carry the declared full `uint64` counter domain exactly

The frame schema uses JCS JSON numbers for `seq`, `re`, and `turn_epoch` while declaring them `uint64`; the broker feed likewise carries `state_seq` as a monotonic counter (`2026-07-16-mvp-ipc-manifest-seam-contract.md:25-35,73-81`). RFC 8785 requires JSON numbers to be expressible as IEEE-754 double precision and recommends strings for longer integers; its interoperable integer range is only `[-9007199254740991, 9007199254740991]` (RFC 8785 sections 3.1, 3.2.2.3, Appendix B note 1: `https://www.rfc-editor.org/rfc/rfc8785.html`). Values in the rest of the `uint64` range can lose identity across implementations, defeating exact epoch/sequence comparison and transition keys.

Required return: either encode every trust-bearing counter as a canonical decimal string with a pinned grammar, or restrict the schema to the JCS-safe integer range and specify exhaustion/rollover as a fail-closed terminal condition. Apply the same rule consistently to `turn_epoch`, `seq`/`re`, `state_seq`, `control_generation`, event sequence, and transition identities wherever they cross JSON.

## What passes on these bytes

- The exact 8-name policy identity, alias-before-equality rule, deny-all floor, F58 pre-lock staging, immutable-per-run intent, and F59 ticket binding are directionally faithful to r7.
- The no-seat/no-conductor-verb/no-secret-byte/no-provider-payload-on-CTRL boundaries are explicit.
- The F59 pre-consume and post-consume/pre-outcome crash dispositions, denial accounting, at-most-once wake schedule, SQLite sole-writer direction, and no-silent-replay vocabulary are present.
- The stage-1 supplement is consumed: m-3 appears as a digest consumer and m-2's reciprocal confirmation is named. R2 above blocks that edge from being satisfiable on the current schema, but the routing itself is correct.

## Gate disposition

This verdict is byte-bound to `b971ea2f342f950c73808903ea12378c7f024099cdd9270d571cf0943bbdf880`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `b971ea2f342f950c73808903ea12378c7f024099cdd9270d571cf0943bbdf880`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK. Root-mode output also carries unrelated historical `INDEX.md`/lineage/merge noise; it does not invalidate the exact-file result.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file mode exit 0; contextual `--relay-root` scan also reports this target `OK` but exits 1 on unrelated historical INDEX/lineage/merge noise
Next requested action: m-10.planner folds R1-R4 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
