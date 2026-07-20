## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r1 must revise: three stale active clauses contradict the repaired mechanisms

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r2
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings are bounded exact-byte consistency defects under already-ratified and pair-approved inputs; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this stage-1 re-review does not reopen the operator-locked topology/F59/F60 decisions; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-081500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-073921.md
SUBJECT: must-revise exact 2a5fb6e1... - r1 repairs R1-R4 in substance but three stale implementation-directing clauses still contradict those repairs

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `2a5fb6e117c9f5067c29d95b3f5fb89bbddecd6381436b43af9584a16212d2d4`. Routing, lineage, `DESIGN_DOC_ID`, the r7 basis hash, and the requested byte binding pass. The revision closes the prior R1-R4 mechanisms in substance, but it is not pair-approvable yet: three active implementation-directing loci still contradict the repaired contract.

## Findings

### R2-F1 - The primary frame grammar still encodes trust-bearing counters as JSON `uint64`

The normative frame schema declares `seq: <uint64>`, `re: <uint64|absent>`, and `turn_epoch: <uint64|absent>` (`2026-07-16-mvp-ipc-manifest-seam-contract.md:27-34`). The immediately following repaired rule says those same fields, plus every other trust-bearing counter crossing JSON, are canonical decimal **strings**, never JSON numbers (`:35-36`). An implementer following the primary grammar would produce the representation R4 forbids.

Required return: rewrite the primary A.2 schema placeholders as canonical-decimal-uint64 strings (including the absent variants) and sweep every copied JSON schema/message example for numeric counter placeholders. Keep the pinned grammar, `< 2^64` bound, numeric decoded comparison, and malformed-frame disposition.

### R2-F2 - The m-10 durable transition ledger omits m-7's required `PREPARING` state

Section B.5 correctly receives the broker's PREPARING crossing-set barrier (`2026-07-16-mvp-ipc-manifest-seam-contract.md:87-95`), but the normative m-10 store census defines `epoch_transitions.state` as only `PROPOSED|CROSSERS_DURABLE|INSTALLED|ABORTED` (`:208-211`). The consumed pair-approved m-7 contract requires the full durable m-10 ledger `PROPOSED -> PREPARING -> CROSSERS_DURABLE -> INSTALLED | ABORTED` (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:129-138,312-313`). Omitting PREPARING makes m-10 unable to distinguish a received/frozen crossing set whose crossing-row transaction has not committed, which is exactly the recovery boundary the identified handshake exists to preserve.

Required return: add `PREPARING` to the `epoch_transitions` state machine and pin its durable write timing: after receipt/validation of the broker's frozen crossing set and before the one-transaction crossing-row commit advances the ledger to `CROSSERS_DURABLE`. Align the transition-ID recovery text with that durable state so crash recovery does not infer whether the set was received.

### R2-F3 - The fault disposition still directs a connector-only restart

The repaired DATA-P lifecycle correctly says death/replacement of either endpoint owner causes m-10 to replace **both** worker and connector under a new epoch with a fresh socketpair (`2026-07-16-mvp-ipc-manifest-seam-contract.md:23-25`). The active crash clause later says, on connector failure, merely “restart connector” (`:68-75`). Those are mutually exclusive lifecycle instructions; connector-only restart cannot acquire a DATA-P endpoint because m-10 retained none.

Required return: replace the connector-only restart instruction with the generation-paired co-restart rule: park the in-flight attempt, terminate and reap the surviving worker under the lease/fencing rules, mint/persist the new epoch, and launch both owners with a fresh DATA-P pair. Sweep the nearby “one connector, restartable” shorthand so it cannot be read as an independent DATA-P-preserving restart.

## What closes from review r1

- R1's m-7 integration is now present in substance: broker-owned dial-in control, CI-1 lock/token/generation adoption, expanded worker assignment, source-specific epoch authority, the identified barrier handshake, broker-event durability/dedup, and closed-family unknown-field behavior.
- R2's manifest now carries `policy_digest`, a distinct `tool_catalog_digest`, the closed build-digests-XOR-release-digest binding, freeze-time lane equality, and exact serve-time checks.
- R3's intended lifecycle is now generation-paired co-restart, and the payload-free CTRL-W/CTRL-C attempt/event ingress is named.
- R4's canonical decimal-string rule and full unsigned-64 bound are explicit. R2-F1 is the remaining stale schema locus, not a rejection of that repaired rule.

## Gate disposition

This verdict is byte-bound to `2a5fb6e117c9f5067c29d95b3f5fb89bbddecd6381436b43af9584a16212d2d4`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `2a5fb6e117c9f5067c29d95b3f5fb89bbddecd6381436b43af9584a16212d2d4`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R2-F1 through R2-F3 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
