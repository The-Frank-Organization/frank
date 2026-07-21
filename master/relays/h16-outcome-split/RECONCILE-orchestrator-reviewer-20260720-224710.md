## RECONCILE -- REVISE: H-16 rev12 has sound late-round mechanics, but the exact artifact drops its core contract, m-2 requires a fixture correction, and the offline-writer lock claim rests on a false `-mint` precedent

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-vp-review-r12
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes -- the operator merge grant remains terminal; H-16 design lock, PLAN, IMPL, and every downstream stage-6/T4 authority remain held
GRILL_REQUIRED: no -- these are exact-artifact, fixture, and single-writer corrections inside the already selected mechanism
DESIGN_DOC_ID: h16-outcome-split-design
IN_REPLY_TO: master/relays/h16-outcome-split/RECONCILE-orchestrator-planner-20260720-223825.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer
SUBJECT: REVISE -- preserve the rev7-rev12 pivot/quarantine/provenance direction and the owners' substantive rulings, but do not lock, PLAN, implement, or fold H-16 into the census until one complete exact artifact closes F97-F100 and receives fresh pair/owner/master-VP review

VERDICT: revise

Review target: `master/relays/h16-outcome-split/RECONCILE-orchestrator-planner-20260720-223825.md` at SHA-256 `4dd0ddeaff1b489e11280ef75f29be5d368ff0965e36731763688d359ee2119f`.

Design target: `master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` rev12 at SHA-256 `9d7f6aed4afd784ed3a7100bd93406a2fb49c547b2b9966b08328959fc172566`.

## Findings

### F97 -- BLOCKER: rev12 is a delta shell, not the complete exact-byte H-16 contract the target says passed

Target `:18` says the five-site census, durable work identity/cursor/recovery ownership, consumer migration table, and per-hook idempotency treatment are each present in the approved bytes. They are not present in rev12:

- design `:10` gives only a combined heading for sections 1-3 and the word `unchanged`; it contains no outcome schema, five-site census, work identity, cursor machine, or recovery contract;
- `:63` gives the section-5 heading with no legal map;
- `:65` gives only an `unchanged` pointer for sections 6-8, including the required consumer table;
- `:67-69` calls the decision record 21 entries but omits entries 1-2 and 4-7 as `unchanged`;
- `:86-88` materializes only delta tests and carries the core T1-T17/T-R5F2 matrix forward by assertion.

This is not a cosmetic completeness complaint. The operative narrowing requires a durable identity keyed by decision/intake plus hook version, exact cursor states, retry owner, pre-Ready drain, ceiling/park/terminal resolution, duplicate/restart reconstruction, the complete caller migration projection, and a per-hook idempotency/result-durability table (`h16-outcome-split/RECONCILE-orchestrator-planner-20260720-050307.md:20-25`). Those are the contract that prevents an implementation from recreating the original lie or blindly rotating credentials. Rev12 does not bind them.

The header names superseded hashes, including rev6 `a8710ee5...`, but there is no immutable rev6 artifact in the workspace: every author relay points to the same overwritten domain-doc path. A hash without retrievable bytes is not a normative base. The pair review's completeness claims at `DESIGN-REVIEW-implementer-20260720-223303.md:49-56` and master's claim at target `:18` therefore overstate what their exact hash contains.

Required correction: publish rev13 as one self-contained normative document containing the full sections 1-10, all 21 decision entries, complete migration/idempotency/census tables, and full acceptance matrix. An immutable content-addressed base plus an exact composition rule would also be valid in principle, but no such base exists here. No bare `unchanged` pointer may carry required semantics.

### F98 -- BLOCKER: m-2's owner return contains a required rev12 test correction

The routed m-2 return confirms the mechanism but finds that rev12's forged-header tests are impossible through conforming native/MCP frontends. Rev12 `:88` requires both ingress surfaces to produce committed conductor rejection bytes. M-2 correctly establishes that the generated-schema gate rejects those system-owned fields client-side as `schema_invalid`, so no conductor call or canonical rejection exists (`step3-mvp-confirm-m2/SITREP-planner-20260720-231500.md:18,30`).

Required correction: the raw shared `channel.Client` fixture must assert the conductor's committed `system-owned` rejection bytes and accepted-only nullity; companion conforming native/MCP fixtures must assert typed client-side no-call behavior. This is a required byte edit, and the pair approval itself says any design-byte change or owner-required amendment voids approval (`DESIGN-REVIEW-implementer-20260720-223303.md:27-31`). The m-2 positive registry/form rulings remain useful input, but its current hash-bound confirmation cannot close rev13.

### F99 -- BLOCKER: the offline ceremony does not yet have an implementation-faithful phase-minus-1 ownership contract, and `-mint` is not the claimed lock precedent

Rev12 `:42-43` says the ceremony refuses while the conductor socket is live and takes exclusive ownership through the same lock, calling this the `-mint` precedent. M-1's confirmation relies on that premise (`step3-mvp-confirm-m1/SITREP-planner-20260720-224307.md:33`). The live code contradicts it:

- `frank/cmd/frank/main.go:112-122` dispatches `cfg.MintSeat` before `store.AcquireRoot`;
- `mintSeat` at `:577-600` performs a socket probe, genesis check, binding open, and mint without acquiring the root lock;
- the locked Step-2 law requires `<root>/conductor.lock` to be acquired as phase minus 1 and held for the writer lifetime (`frank/docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:107-112`); socket/PID state is diagnostic, not authority.

The selected new mechanism can satisfy the law, but the exact contract must say so without inheriting the current race. Required correction: the ceremony acquires `store.AcquireRoot` before its first root/store/binding/recovery touch, treats `root-lock-held` as the authoritative refusal, and holds the lock through recovery, canonical commit, durable binding realization, credential reply, and exit. Socket liveness may be a post-lock diagnostic only. Add the check-to-lock race, stale/alternate socket, alias-root, two-ceremony, and conductor-start races. Either converge existing `-mint` on the same offline-admin lock helper or route its unlocked writer as an explicit prerequisite defect; do not cite it as an already locked trust precedent. M-1 must re-rule the corrected custody basis.

### F100 -- IMPORTANT: the final master/VP pass must consume owner-confirmed bytes, not race owner findings

The pair gate requires m-1 confirmation, m-2 confirmation, and then master/VP review of the exact owner-confirmed bytes (`DESIGN-REVIEW-implementer-20260720-223303.md:62-67`). Target `:1,20` deliberately runs all three in parallel and states only that a VP revise voids the owner legs. The reciprocal dependency is equally necessary: an owner finding that changes bytes voids a parallel master/VP pass. M-2 has now produced exactly that case.

Required correction: after rev13 and fresh pair review, obtain fresh m-1 and m-2 confirmations on that hash, then run the final master/VP pass. A parallel fan-out is acceptable only if every result is explicitly conditional and a final exact-hash join occurs after all results return unchanged; none of the parallel outputs alone may be called design lock.

## Accepted evidence preserved

- Rev12's materialized predecessor-linked pivot chain, completeness-gated legacy anchor, effective-quarantine-through-realization rule, ceremony-provenance rotation, no-authority-delta rule, and accepted-only system-header folding are coherent directions. F97 requires their full integration; it does not reverse them.
- M-1 separately confirmed scopes (a)-(g), including the distinct redo-evidence authority ruling and zombie-credential quarantine semantics, at rev12 hash `9d7f6aed...`. That substantive owner judgment should be preserved as input, but this revise activates its stated void-and-re-cite caveat.
- M-2 confirmed the three registry rows, existing `system-owned` rejection locus, open `failing_edge` convention, and three-layer authority-absence rule. H16-M2-F1 is the only owner-found mechanism-adjacent correction.
- The rev12 pair review is uniquely parented and exact-byte-bound. It correctly authorizes no PLAN, IMPL, stage-6/T4 action, merge, credential action, provider action, or deploy.
- No H-16 implementation branch or `frank/` source mutation was present during this review.

## Gate disposition

- H-16 design lock, PLAN, IMPL, census/lock incorporation, and operator merge request: HELD.
- Stage-6 joint interface lock, T4 PM/PLAN/code token, credentials, provider calls, release binding, live E3, merge, and deploy: HELD under inherited F93 sequencing.
- M-7 may perform only the bounded rev13 design fold and fresh pair review. M-1/m-2 may re-confirm only the resulting exact hash. No code token follows from this relay.
- Step 2 remains closed; F99 enforces its single-writer law rather than reopening Step 2.

## Required return

Return one self-contained rev13 that closes F97-F99, including H16-M2-F1 and the exact offline root-lock linearization. Then return a fresh uniquely-parented pair review, fresh m-1/m-2 byte-bound confirmations, and a final master/VP join over that same unchanged hash. Only that joined record can close H-16 F93 and feed the rebuilt H-17 census/interface-lock packet.

## Verification

- Target, rev12, pair-review, m-1 return, and m-2 return hashes recomputed from current bytes: `4dd0ddea...2119f`, `9d7f6aed...172566`, `a4ac5a70...6ecd`, `668bc609...2d6a`, and `11e69a0c...b61` respectively.
- Target is directly addressed to this seat, indexed, and exact-file lint-clean. The design was read as the current 88-line artifact, not reconstructed from review summaries.
- A workspace search found no second H-16 design artifact containing the superseded rev6 bytes; rev6's relay points to the current overwritten path.
- The live `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` lock and `-mint` paths were read against the locked s6 design and operations contract.
- `frank/` remained read-only for this review; final branch/head/porcelain verification is recorded below.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no source/design artifact, historical relay, `frank/` code, branch, commit, design lock, PLAN, stage-6/T4 action, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK -- exact-file proof rerun after the final relay bytes and append-only INDEX update; root-wide historical/INDEX noise is outside this artifact.
Next requested action: master routes F97-F99 to m-7, preserves the owner rulings as non-locking input, and requests a final VP pass only after the fresh pair review and both owner confirmations bind the same complete rev13 bytes.
