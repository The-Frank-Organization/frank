## DESIGN-REVIEW - s1-core.implementer review of s1-slice-1-design r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s1-core-design-r1-review-implementer
PARENT_DISPATCH_ID: s1-core-design-r1-review
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
IN_REPLY_TO: s1-core-design-r1-review/DESIGN-planner-20260703-150440.md
DESIGN_DOC_ID: s1-slice-1-design
DESIGN_REVIEW_VERDICT: must-revise
SUBJECT: read-only design review - S1 r1 needs bounded revisions before lock

Phase:
Read-only DESIGN-REVIEW. I reviewed `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md` at main@e8faeed against the addressed review relay. I made no source/test/sprint-spec edits, no implementation branch, no commit, no PR, and no prototype.

Verdict:
must-revise.

Review scope:
- Incoming review request: `.relays/s1/s1-core-design-r1-review/DESIGN-planner-20260703-150440.md`.
- Candidate design: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md`.
- Guide/reconcile context: `s1-guide-q1/SITREP-orchestrator-planner-20260703-140843.md`, `s1-core-design/SITREP-orchestrator-planner-20260703-142800.md`, and `docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md`.
- Frozen contract surfaces cited by the design and dispatch: m-1 trust/identity store API, m-2 FieldSpec/lineage/gate semantics, and m-7 conductor core.

Findings:

1. Blocker - Gate, park, and outbox have a crash window that can violate B4 while every named S1 recovery step still passes.

Evidence:
- Design D-7 says the gate-bearing candidate first commits as accepted, then the loop commits separate park-transition and outbox-enqueue mutations: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:112`.
- AC B4 claims one gate-bearing accepted record yields exactly one outbox item through one pivot: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:150`.
- D-9 recovery rebuilds projections, re-enqueues intake outcomes, restores bindings, and re-issues wake nudges, but it does not say how a crash after the accepted gate record and before the park/outbox follow-up records is completed idempotently: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:128`.
- The guide-confirmed ODB answer keeps S1 to typed local ODB production through the conductor-governed pivot, so the design cannot rely on later egress machinery to heal this gap.

Why this blocks lock:
An implementer can build the listed recovery steps and still lose the required local outbox item if the process dies between the accepted gate commit and the derived park/outbox commits. That makes the AC row weaker than the exit gate.

Required revision:
Specify one of these buildable mechanisms: either the accepted gate, park transition, and outbox enqueue are one committed mutation/pivot, or the accepted gate record creates a durable pending derived-work item and startup completes park/outbox idempotently before channels open. Add a crash fixture for kill after gate-accept and before park/outbox completion.

2. Blocker - The MVP `grant` rendering appears to make delegated pair-Planner implementation dispatch impossible.

Evidence:
- D-5 hides the whole `grant` field from pair-seat forms: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:91`.
- AC A2 requires pair-seat rendered schema to omit both `merge-gated` and the entire `grant` field, with raw submissions rejected: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:152`.
- The active protocol requires implementation to start only from a valid bare implementation grant addressed to the Implementer, and m-2 treats delegated dispatch/merge grants as authority-bearing (`m-2 design:345`).

Why this blocks lock:
The review request asks for the implementer path to be buildable, but the proposed rendering forecloses the normal pair lifecycle where the Planner later sends a delegated implementation dispatch after gates clear. If S1 intentionally reserves implementation grants to orchestrator/operator only, that is a product/protocol decision that must be explicit and routed.

Required revision:
Split the rendering rule. At minimum, keep `dispatch-merge` absent from pair seats and phase-bound to MERGE-GATE, but either allow `dispatch-impl` only on authorized Planner dispatch forms after PLAN/external gates, or state explicitly that S1 does not support pair-Planner implementation dispatch and route that decision to the orchestrator/operator before lock.

3. Blocker - The S1 authority-bearing rule is narrower than the frozen m-2 authority surface.

Evidence:
- D-6 defines authority-bearing as `grant != absent` or gate-bearing only, and defers full `record_kind` taxonomy to S3: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:105`.
- Frozen m-2 classifies more than grant/gate as authority-bearing, including design-doc PLAN locks, delegated dispatch/merge grants, substantive IMPL reports, merge claims, orchestrator-authority relays, and A-category gates (`m-2 design:345`).
- D-6 then uses the authority-bearing test to decide lineage blocking and authority-fault held behavior: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:106-108`.

Why this blocks lock:
Deferring taxonomy to S3 is fine for full classification, but S1 still runs on real relays from the pair/orchestrator lifecycle. With the current rule, an authority-bearing orchestration or design-lock record that lacks `grant` and `gate_category` can bypass the S1 authority path.

Required revision:
Define the S1-minimal authority classes now using existing MVP-visible fields, even if the full taxonomy remains S3. The minimum needs to cover delegated implementation grants, merge grants, design/PLAN lock records, orchestrator-authority relays, substantive implementation reports, merge claims, and A-category gates, or else state which of those are deliberately out of S1 with guide/orchestrator approval.

4. Blocker - B1 overclaims direct file injection protection relative to the D5 and S2 checksum boundary.

Evidence:
- The design scope explicitly excludes genesis/quarantine/GC and segment-rotation machinery: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:18`.
- D-9 excludes canonical-checksum quarantine disposition from S1 recovery: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:128`.
- AC B1 says a hand-written file dropped into `records/` is not served as a record, while its parenthetical narrows the claim to "no-tool-path" and leaves detection to S2: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:147`.
- D5 acknowledges a same-uid shell-bearing seat can reach the store, sockets, and config outside the tool surface: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:10`.

Why this blocks lock:
The no-tool-path claim is S1-shaped and buildable. The stronger "drop a hand-written file into records and it is not served" claim is not guaranteed unless S1 validates and rejects/quarantines direct record files, which the design otherwise moves out. A same-uid shell can write a syntactically valid record file unless the design adds a stronger integrity discriminator.

Required revision:
Change B1 to assert only that no seat-facing tool/API path can create or serve a record except conductor `submit`, and move direct record-file injection detection/quarantine to S2. If S1 keeps a direct-file negative fixture, limit it to malformed/torn/staging files that S1 actually detects, and state the D5 residual plainly.

5. Blocker - The exact S1 `gate_category` token subset is still left to PLAN, but it drives validation, A/B classification, park/outbox, and known-A behavior.

Evidence:
- D-5 says the S1 subset of J2 tokens will be fixed at PLAN with guide visibility: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:89`.
- Open item 6.2 repeats that the exact S1 `gate_category` subset is a PLAN item: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:179`.
- D-7 gate/park/outbox and D-6 authority-bearing both depend on this field: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:105,112-117`.
- Frozen m-2 already supplies byte-exact J2 semantics and the monotonic known-A/RAISE-only rule (`m-2 design:270-277`, `m-2 design:381`).

Why this blocks lock:
An implementer cannot build the S1 validator, known-A floor, gate-bearing detector, or B4/A4 fixtures without inventing the actual enum set. This is not a PLAN sequencing detail because PLAN would be choosing a contract that the design claims to lock.

Required revision:
Either lock the exact S1 token subset in DESIGN or choose the full frozen J2 enum set for S1 while constraining fixtures to the minimal cases. Keep PLAN to naming tests and code placement, not inventing accepted enum tokens.

Non-blocking watch items:
- The binding table restore line is directionally correct, but PLAN should name the atomic write or static-setup assumption for `binding/seats.json` because `mint_seat` persists credentials and restart restores bindings (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:45,64,128`).
- The MCP Go SDK fallback is acceptable as an open implementation choice if the fallback preserves per-seat channels, stamped identity, and the exact `{submit, project, read}` seat registry (`docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:44-49,178`).

What is sound enough to keep:
- Go runtime selection and one-process goroutine topology are aligned with the operator decision and the m-7 single intake writer / serialized commit-loop shape.
- The store decomposition, staging/intake/outcome/projection split, and Package-A pivot direction are compatible with the guide.
- The R1 replay classification approach is appropriately scoped if PLAN maps each selected fixture to the MVP FieldSpec and treats uncovered S3 rows as explicit non-S1 coverage, not failures to port the old linter.
- A4 one-shot double-accept is the right minimal instance, assuming the authority-bearing rule above is corrected.

Required before approval:
1. Close the five blockers above in the design doc.
2. Add or revise the acceptance criteria so each blocker has a fixture that would fail against the current r1 ambiguity.
3. Keep OUT items out: do not absorb full observe/egress scan, full FieldSpec registry, genesis/quarantine/GC, full phase-0-to-4 recovery, or model_name routing into S1 while closing these blockers.

Actions:
- Wrote this DESIGN-REVIEW relay.
- Appended the corresponding routing row to `.relays/s1/INDEX.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits claimed; relay artifact written at `.relays/s1/s1-core-design-r1-review/DESIGN-REVIEW-implementer-20260703-151318.md`; index row appended at `.relays/s1/INDEX.md`; final git status captured below.
FINAL_GIT_STATUS_SHORT: none - clean tree

Tests / verification:
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-design-r1-review/DESIGN-planner-20260703-150440.md` passed.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-design-r1-review/DESIGN-REVIEW-implementer-20260703-151318.md` passed.
- `git status --short --untracked-files=all` final output captured in `FINAL_GIT_STATUS_SHORT`.

Next requested action:
- `s1-core.planner` revises `DESIGN_DOC_ID: s1-slice-1-design` or routes any product/protocol decision to `s1.orchestrator-planner` / operator before seeking approval again.
