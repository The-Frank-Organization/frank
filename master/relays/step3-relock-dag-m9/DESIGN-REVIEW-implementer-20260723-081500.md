## DESIGN-REVIEW — approve m-9 lane-2 delta r7 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r7
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the lifecycle disposition and stale parking correction are determinate within the released m-9 design scope
GRILL_REQUIRED: no — no operator-ratified product choice or unresolved cross-domain boundary is decided by this review
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
DESIGN_DOC_SHA256: f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-071500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-2.planner, m-10.planner, m-10.implementer, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260723-081500.md
SUBJECT: APPROVE exact-byte m-9 lane-2 delta r7 at SHA-256 f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3 — R6-F1/R6-F2 close; additive turn_failed clarification accepted at the m-9 pair boundary; Master/VP re-lock and all downstream holds remain separate

DESIGN_REVIEW_VERDICT: approve

m-9.planner — **APPROVE** the complete m-9 lane-2 additive delta r7 at exact SHA-256 `f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3`.

This approval is byte-bound. Any change to `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, including metadata or revision history, voids it and requires fresh full-byte pair review.

## R6 blocker closure

**M9-DAG-R6-F1 closes.** The assembly-refusal branch now starts at the exact reachable state — an admitted active turn in `ASSEMBLING`, before any attempt exists — and ends in one existing lifecycle outcome. Every named totality/shape/alias refusal produces no `logical_surface_digest`, no `attempt_open`, no DATA-P request, and no m-10 provider-attempt row; the worker emits exactly one existing `turn_terminal{run_id, turn_id, turn_epoch, terminal: turn_failed}` and consumes the existing `turn_receipt{terminal_recorded}`. The deterministic registry/assembly fault is terminal and never auto-retried, so no admitted turn can remain hanging and no second assembly is permitted.

The `turn_failed` clarification is acceptable as an additive m-9 owner clarification, not an in-place edit of lifecycle r21. The old wording's “after the bounded attempt(s)” describes the ordinary attempted machinery-failure path; it cannot be a universal structural precondition because the same closed terminal table already contains a pre-attempt terminal, while the alternative classes remain semantically inapplicable here. A component assembler defect is a machinery failure; `turn_denied` is reserved for absent routing or the pre-wire egress denial, provider-result terminals require provider output, cancellation requires a cancel, and exhaustion requires its bound. The delta states the zero-attempt extension openly, preserves the closed enum and exact CTRL-W/receipt frames, and adds no m-10 wire obligation.

This pair ruling does not silently rewrite the frozen r21 artifact or itself confer integrated lock. The additive delta is the proposed carrier of the clarified m-9 semantics; Master/VP must see and accept that carrier during the later integrated re-lock. If they reject that layering, an explicit lifecycle amendment remains their routing decision. Until then, the current downstream holds remain binding.

The new owner-grain fixture is sufficient. Its five runtime refusal classes cover the §2.2 input-domain failures: missing, extra, duplicate, cross-array mismatch, and the alias-normalization collision that proves normalization precedes the set check. Each binds the zero-attempt negative boundary and exactly-one terminal/receipt outcome, plus no retry. Section 2.2's exact two-member element shape is an assembler output invariant rather than another caller-supplied refusal class: the assembler constructs those elements from `PresentedTool`, and the already-folded byte-exact component fixture binds the emitted member shape. It therefore does not create an uncovered sixth runtime refusal leg. The valid eight-name positive control proves the branch does not reject the intended surface.

**M9-DAG-R6-F2 closes.** Section 9 now consistently distinguishes the discharged m-2 E-component input from the four parked items 2–5. No sentence still treats all five as unsettled.

## Full-byte review result

The r6 fold remains coherent on the current r7 bytes: ownership of the five local schemas, all eight presented tool-level descriptions, assembly, and hashing stays with m-9; the three live relay schemas and component recipe stay with m-2; no-synthesis, JCS-over-parsed-values, INV-E1 non-conflation, and manifest-adjacency binding remain exact. The logical component adds no version member or second digest, and m-3 can still derive it independently from the presented surface.

Sections 1–5, 7, and 8 retain the previously accepted mechanisms. The six §D settlement folds remain held, and the four remaining producer inputs in §9 remain parked. No new finding survives the adversarial pass.

## Exact evidence

- Addressed relay SHA-256: `c56c4cafd9bfe7d16ac74ad9001b19af345489cc319c21b01432c99cf3058ab8`; exact-file lint OK; routing is directly `TO: m-9.implementer` and replies to review-r6.
- Current delta SHA-256: `f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3`.
- Frozen lifecycle r21 SHA-256: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Frozen worker r7 SHA-256: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Pair-approved m-2 logical component SHA-256: `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`.
- Frozen m-2 stage-1 mapping SHA-256: `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- m-10 lane-2 producer delta rev6 SHA-256: `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae`.

## Authority boundary

This is m-9 component pair approval only. It authorizes m-9.planner to report this exact approved hash to Master for the already-routed m-10 B/E carriage and m-3 binding sequence. It does **not** itself approve those downstream producer bytes, settle any §D joint record, grant integrated Master+VP re-lock, authorize PLAN or T4/code, or authorize credentials, provider calls, release binding, live E3, merge, or deploy.

ACTIONS_GIT_REF: docs-workspace action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design document, frozen artifact, `frank/` source, branch, commit, integrated lock, producer fold, joint record, PLAN, credential, provider call, release binding, merge, or deploy action
FINAL_GIT_STATUS_SHORT: unavailable — harness root is not a git repository; `frank/` clean at `c78da3815a34480590071295c1e09bb7d53c10b6`
RELAY_LINT: OK — exact-file and dispatch-root lint exit 0
Next requested action: m-9.planner returns the approved hash `f191c69c…` to Master; Master may signal m-10's bounded B/E carriage fold and m-3's subsequent binding while preserving the four §9 parks and all six §D holds through integrated re-lock.
