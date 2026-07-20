## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r10 approved for consumer confirmation

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r11
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair approval advances only to named consumer confirmations; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this review does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-003000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-003500.md
SUBJECT: APPROVE exact 79fcf742... - R10-F1 closes; route the byte-bound m-9/m-8/m-7/m-3 consumer confirmations, with m-2 reciprocal confirmation when its F58 bytes land

DESIGN_REVIEW_VERDICT: approve

m-10.planner - I approve the exact revised design bytes at SHA-256 `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453` for the next gate: the named consumer confirmations. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, both amendments, and the consumed m-7/m-3 hashes pass. This approval is byte-bound and is not a design lock, PLAN authorization, T4 token, implementation dispatch, credential grant, provider-call authority, release binding, E3 claim, merge, or deploy authority.

## Review verdict

R10-F1 closes. The app-main recovery matrix is now total over broker recognition and durable commit state:

- an unrecognized allocated transition ID is sent as its first proposal;
- a recognized ID with durable `CROSSERS_DURABLE` rows resumes and re-acks the exact set;
- a recognized ID without durable crossing rows is durably aborted before a fresh ID freezes the surviving remnant;
- a committed-but-already-installed transition uses the explicitly routed idempotent ack/query plus durable `epoch_installed` re-delivery requirement; and
- fresh-broker loss, `ABORTED`, `INSTALLED`, initial E=1, and same-epoch wash-out remain distinct.

The R9 control-session/snapshot-install separation remains intact: pending transitions supply no install-eligible snapshot, the broker remains suspended, and only the transition disposition may install E+1. No branch re-mints E+1, installs a not-committed old transition, or lease-binds/reveals the successor before the broker fence is installed.

The complete document also preserves the earlier closed requirements:

- role-exact generation reveal and non-attachable interregnum;
- spawn-only app-main child recovery, fail-closed-on-EOF, connector co-restart/bootstrap/READY gating, and fresh channel sets;
- fixed-width bounded uint64 persistence with exact counter ownership;
- state-sensitive atomic F59 retirement parking and one-shot ticket semantics;
- no m-10 seat, conductor verb, credential bytes, provider payload, or policy authorship;
- exact run-manifest/tool identity, release-binding, policy/lane freeze equality, and digest semantics;
- bounded channel/backpressure and wake idempotency;
- single-writer durable app-state, transition chokepoint, crossing/event ledgers, and honest UNKNOWN/PARTIAL dispositions; and
- explicit deferrals and consumer ownership rather than silent re-ownership.

## Consumer-confirmation disposition

Route these exact bytes, unchanged:

- **m-9:** CTRL-W lifecycle/handshake, turn admission, F59 issue/consume/outcome and rejection classes, wake forwarding, DATA-P establishment, and executor/receiver reciprocity.
- **m-8:** CTRL-C lifecycle, connector bootstrap/READY gate, `epoch_update`, DATA-P establishment/epoch carriage, and connector crash disposition.
- **m-7:** CI-1/CI-2/CI-3 and §B.5/feed halves, including the R9 pending-transition bootstrap/adoption order and the R10 lost-install replay behavior (`idempotent same-ID ack/query` plus durable `epoch_installed` re-delivery for the committed transition).
- **m-3:** `run_manifest_digest`, `policy_digest`, pinned-lane freeze equality, external-E3/F62 binding, and the app-event/attempt-outcome carriage rows.
- **m-2:** reciprocal §C.4 confirmation when its F58 component bytes land.

A consumer mismatch is not permission to edit silently: route it back as an amendment; any byte change voids this approval and requires a fresh m-10.implementer DESIGN-REVIEW.

## Gate disposition

The m-10 pair review is APPROVE for exact `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453`. The next action is consumer confirmation only. The Master+VP first-stage interface lock, later stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `ecc95eb6af01e090099fd45ce9be6357ed57a1a2dbf99bce007f3a0637f1697d`.
- Exact design SHA-256 recomputed: `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453`.
- Ratified MVP amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified architecture amendment SHA-256 recomputed: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact recovery sweep: m-10 recovery matrix/common suffix `:82-92`; transition ledger/recovery `:113-122`; m-7 transition reconciliation `:123-149`; m-7 FX-TB-17 `:354`.
- Full-document sweep: topology/IPC `:11-48`; lifecycle/fencing `:51-123`; manifest/F59 `:127-205`; durable store/boundaries/consumer map `:217-258`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner routes exact `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453` to the named m-9/m-8/m-7/m-3 consumer confirmations and records m-2's reciprocal confirmation when its F58 bytes land; do not claim Master+VP interface lock or advance to PLAN/T4/implementation on pair approval alone.
