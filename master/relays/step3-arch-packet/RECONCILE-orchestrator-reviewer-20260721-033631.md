## RECONCILE -- REVISE: the re-scope direction is right, but the proposed lock split is not mechanically bindable, the authority host is moved to m-9, and the claimed acyclic order is cyclic

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- Option-1 settles sandbox inclusion and the milestone label, but the backend matrix, executable `bash` policy, and product-evaluation bar still require explicit operator choices before re-scope ratification
GRILL_REQUIRED: yes -- the amendment introduces hard-to-reverse cross-domain wire, durable-state, authority, containment, and evidence contracts; the earlier Option-1 grill did not settle the remaining choices named in F106
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-031500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- preserve the hold, Option-1 direction, and B-E interface intent; make the tier lock hashable, restore policy/authority ownership, close the `bash` effect claim, replace the pair-total order with an interface DAG, define the journal truth model, and surface the remaining operator choices

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-031500.md` at SHA-256 `9c96b02f987d0753dee8d154044e3d79f800de34cf869610175f008236aee4f2`.

Proposed amendment: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `0634b6e42e324ab3fd8858bfcf0c58105be12f8662b8f24e03a9b0a7d0dd6985`.

## Findings

### F101 -- BLOCKER: Tier-HARD and Tier-SOFT have no hashable artifact boundary

Amendment `:59-76` classifies clauses, while `:180-181` says the shorter lock will bind changed Tier-HARD interfaces and carry unchanged design hashes forward. The actual realization documents mix hard wire/schema material and soft wording/ergonomics in the same files. A soft edit changes the same whole-file digest as a hard edit. The proposed lock therefore has no deterministic answer to either question: what exact bytes are locked, and what exact byte change invokes F73?

Required correction: make Tier-HARD a manifest of dedicated, exact-path/full-SHA interface artifacts or a mechanically generated canonical interface bundle with a versioned extraction recipe and digest. Whole-file hashes may carry only when the whole file is hard. Do not use prose section labels or line ranges as the lock boundary. For `model_surface_digest`, lock the canonical recipe, field set, producer ownership, and carriage schema; the per-attempt value and revisionable prompt/description bytes are not themselves a constitutional constant.

Also split the current broad SOFT classes. Cosmetic surface wording may be soft; required UNKNOWN/hold/approval visibility and action semantics are safety contracts and cannot ride an unconstrained `user-surface behavior` exemption. Likewise, only tool UX semantics may be soft; canonicalization, confinement, process teardown, and effect identity are hard under F103.

### F102 -- BLOCKER: the fixed policy silently moves the ratified authority-enforcement boundary into m-9 and omits its policy/surface owners

Amendment `:121-138` adds a real fixed authority policy and assigns its enforcement point to m-9. That conflicts with the standing ownership table: `CLAUDE.md:62-67` makes m-9 a requester/consumer whose parsed call remains inert and m-10 the app-side authority enforcement point; `master/ARCHITECTURE.md:541-554,565-568` says m-10 hosts F59 while m-9 consumes and executes, m-5's ceiling is stood down for the old MVP, and m-6 owns human-gate surface semantics.

The new policy is not just the existing operator-fixed 8-name constant. It introduces path/network/secret ceilings plus dynamic destructive-operation holds and protected-path approvals. That necessarily amends the old `m-5 owns no MVP policy` / empty-seam disposition even if the policy remains fixed rather than config-derived.

Required correction:

- state the exact old architecture clauses this amendment supersedes;
- name the operator-ratified fixed policy as a Tier-HARD policy artifact and route m-5 for ceiling semantics, m-3 for local egress/content-safety, m-1 for secret/path isolation, and m-6 for hold/approval surface and durable human-gate semantics;
- keep m-10 as the F59 policy decision/authorization host; m-9 plus the sandbox backend enforce the authorized descriptor at execution as actuator and defense in depth, not as the sole authority source;
- add m-5/m-6 to the owner/consumer and confirmation graph. If master intentionally wants operator-owned policy with no m-5 policy role, say so as an explicit ownership amendment and still preserve m-10 hosting plus m-6's human-gate boundary.

### F103 -- BLOCKER: the descriptor and fixed policy cannot make the stated exact-effect claim for arbitrary `bash`

Amendment `:89-95` binds a command/context descriptor, but `:121-145` puts path resolution, symlink behavior, shell identity, fixed cwd/env, resource limits, and background-process behavior in Tier-SOFT. Those values determine both the authorized effect and the containment guarantee. A canonical path string is not stable under symlink/rename races, a `backend_id` is not an executable/image identity, and a command string cannot enumerate the filesystem/network/process effects an arbitrary shell will perform.

The policy also says destructive filesystem operations require a hold and protected-path writes require approval, but an unrestricted shell can perform both unless the design supplies a mediation mechanism. A static workspace/network sandbox alone cannot infer destructive intent from arbitrary shell semantics.

Required correction: define the hard local-effect contract before the operator gate. It must bind the canonicalization algorithm and race-resistant resource handle/snapshot, executable or sandbox-image identity, sandbox-policy/profile digest, cwd/env application, network policy, process/resource/daemon teardown, and outcome identity. Then make an explicit product choice for `bash`: for example read-only by default, always-held, or confined to a disposable copy-on-write workspace whose changes are separately reviewed/applied. If `bash` remains arbitrary workspace-write authority, narrow the claim to exact invocation-context binding and do not claim exact affected resources or protected/destructive per-effect holds.

### F104 -- BLOCKER: the `m-1+m-8 -> m-10 -> m-3 -> m-9 -> m-2` order is not acyclic

The owner facts in amendment `:83-138` contradict the total order at `:171-179`:

- B's E0 event is authored/carried by m-9, but B omits m-9 from the chain.
- C requires m-10's ticket schema before m-9's executor derivation.
- D requires m-9's journal content/source and resume contract before m-10 can define faithful persistence, while m-9 also consumes m-10's commit/ack semantics. This is a two-party contract, not a one-way m-10-before-m-9 edge.
- `model_surface_digest` has m-9 and m-2 producers, yet m-3 is ordered before both while owning the consuming E3 binding; m-2 is placed last despite being a foundational schema producer.
- The m-7 study can change Tier-HARD epoch/attach semantics consumed by F59, journal recovery, and UNKNOWN handling. Its design decision cannot finish in parallel after those consumers or after the re-lock.

Required correction: replace the pair-total order with a per-interface dependency DAG and explicit join records. At minimum: m-8 frozen-core producer -> m-9 carrier -> m-10 attempt row -> m-3 E0/E3; m-10 descriptor/ticket contract -> m-9 executor consumer; m-9 journal source/schema <-> m-10 persistence contract as a coordinated two-sided seam; m-2+m-9 surface component producers -> canonical aggregator -> m-3 E3 consumer. Resolve the m-7 simplify/retain decision before affected m-10/m-9 finals and before the re-lock. If cross-epoch completion survives, run H-24 before re-lock rather than locking a protocol that the pre-T4 model may immediately invalidate.

### F105 -- BLOCKER: the run journal is called a projection without naming a canonical source capable of rebuilding it

Amendment `:97-104` includes provider-visible output items, compaction events, objective/constraint refs, settled tool results, and immutable content references. Several are currently ephemeral. If the journal is the first durable copy, it is canonical context state, not a projection; if it is truly a projection, the amendment must name the canonical event/blob records from which every field is reproducible.

Required correction: provide a field-level source map and define the content/blob store, writer, reader, append/checkpoint linearization, atomicity with canonical tool/attempt records, idempotency, crash cuts, retention/GC, size bounds, integrity digests, and secret/redaction/access boundary. Route m-1 over journal content and at-rest references. Distinguish the hard durable record schema and resume predicate from the soft model-rendering/compaction presentation. Preserve no-second-outcome-truth by construction, not by label.

### F106 -- BLOCKER: the earlier Option-1 decision does not settle the remaining product choices, so `GRILL_REQUIRED: no` is premature

Target `:12` treats the product choice as settled, but amendment `:107-125,146-155` independently chooses three shipped backend modes, weaker-native-macOS claim semantics, a dynamic hold/approval policy for `bash`, and a one-`representative task` utility gate. Those are new cross-domain and user-visible decisions. Option-1 as recorded settles only `sandbox in` plus retaining the Frank-harness-MVP label.

Required correction: run and durably record a bounded operator grill before re-scope ratification for:

1. backend scope: Linux-strong only for Step-3, Linux-strong plus VM parity, or also a shipped native macOS weaker tier;
2. `bash` authority: read-only, always-held, copy-on-write/reviewed-apply, or explicitly broad workspace authority with a narrowed claim;
3. product proof: predeclare the task corpus/hash, baseline/model/config parity, pass threshold, prompt-injection fixture, and governed handoff requirement. Item 6 may remain a step-close gate rather than an interface-lock artifact, but `representative` cannot be selected post hoc.

## Accepted direction preserved

- Holding and superseding the unratified joint-lock proposal `b7e1f0ef5116ca8c7bbfb8fde152498c630571fe683fa13340c2fd984a56bb5a` is correct. No prior lock issued.
- The external-review artifact, proposed amendment, backlog update, and target relay hashes reproduce exactly; the target is directly addressed, indexed, and exact-file lint-clean.
- B's `frozen_core_digest` join, C's move from raw-args identity to execution-context identity, D's settled-round durable recovery contract, and E's sandbox/surface/predicate direction belong in the Tier-HARD interface work once the findings above are corrected.
- The item-4 partial framing is sound: a separate secret-holding broker and worker-generation epoch fence remain justified even if app-main survival/adoption/cross-epoch completion is removed. The study's placement, not that framing, must change.
- Product evaluation belongs at step close, not as proof available before a build. Its acceptance contract must be fixed before T4 execution so the gate is not self-selected afterward.
- The nine prior design finals, H-16 rev16, and census remain byte-identical to the r4-approved set. The r4 approval remains true about those bytes; it grants no approval to this new amendment.

## Gate disposition

- Proposed stage-6 re-scope amendment `0634b6e4...`: REVISE; not ready for operator ratification.
- Pending all-artifact joint-lock proposal `b7e1f0ef...`: correctly HELD and must not issue.
- Stage-6 re-lock: NOT REQUESTABLE until the corrected hard-interface artifacts complete their owner reviews and consumer joins.
- PLAN, T4/code token, H-16/H-26 implementation outside their own gates, credentials, provider calls, release binding, live E3, merge, deploy, and Step-3 close remain HELD.
- Step 2 remains closed.

## Required return

Return decomposition review r2 only after F101-F106 are answered in current amendment bytes: a mechanically hashable hard-lock artifact set; explicit authority/policy/surface ownership and superseded clauses; an executable `bash`/effect contract; a per-interface DAG with the broker decision before re-lock; a canonical run-journal source/commit model; and a durable operator grill record for the remaining product choices. Preserve the accepted direction and all unchanged prior hashes.

## Verification

- Target SHA-256: `9c96b02f987d0753dee8d154044e3d79f800de34cf869610175f008236aee4f2`; exact-file lint: OK.
- Amendment SHA-256: `0634b6e42e324ab3fd8858bfcf0c58105be12f8662b8f24e03a9b0a7d0dd6985`.
- External review SHA-256: `b4e79f3bfdc4ca0b2f5023b441da1e0006c7064a5d4b6b06bb4e9e2a7b5bf7d1`.
- Held joint-lock proposal SHA-256: `b7e1f0ef5116ca8c7bbfb8fde152498c630571fe683fa13340c2fd984a56bb5a`.
- Prior nine design finals, H-16 rev16, and census recompute to the r4-approved hashes.
- `frank/` is current at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken by this review.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master corrects the decomposition and returns r2; only a passing r2 routes the revised amendment and bounded grill record to the operator re-scope gate.
