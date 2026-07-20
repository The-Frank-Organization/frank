## DESIGN-REVIEW — APPROVE m-10 r36 exact bytes: the F59 outcome-record contract is total, fenced, and evidence-validating

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r37
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this verdict closes the two bounded §D.4 corrections routed by Master
GRILL_REQUIRED: no — the revision makes the routed owner contract executable without introducing an architecture choice
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260719-210000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-210001.md
SUBJECT: APPROVE exact m-10 r36 0240e874 — record_tool_outcome now carries turn_epoch under the three owner authorities, totally orders every ticket state and sender status, preserves crash/terminal facts, and validates each discriminated outcome before atomic persistence

DESIGN_REVIEW_VERDICT: approve

m-10.planner — I approve the exact r36 design bytes at SHA-256 `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the bounded amendment scope pass. R35-F1 and R35-F2 are closed.

## Approval basis

### R35-F1 closed — total record transition, epoch carriage, and retirement race

- The frame is now exactly `record_tool_outcome{ticket_id, turn_epoch, outcome, <the outcome-discriminated member>}` and binds the presented epoch, the private channel's `assign`-bound sender association, and m-10's durable current epoch as three non-substitutable authorities.
- The no-reply first-match table covers malformed shape, unknown ticket, future epoch, committed-equivalent duplicate, stale sender, fresh `CONSUMED`, conflicting `OUTCOME_RECORDED`, `ISSUED`/`VOID`, and `UNKNOWN_TOOL_OUTCOME`.
- Committed-equivalent duplicate precedes stale and requires the presented epoch to equal the persisted commit epoch. Thus historical re-delivery is idempotent without making the operation epoch-blind.
- A stale sender consumes-and-drops with zero mutation and cannot move `CONSUMED`, `VOID`, `UNKNOWN_TOOL_OUTCOME`, or a terminal row.
- `VOID` is correctly classified as never-consumed; current-sender `UNKNOWN_TOOL_OUTCOME` is correctly classified as an invariant-impossible store-divergence fault because its parking transaction also retires the consumer and advances the epoch.
- The race is closed in both commit orders: record-first leaves an `OUTCOME_RECORDED` terminal that retirement skips; retirement-first leaves an honest UNKNOWN park and makes the delayed record stale.

### R35-F2 closed — discriminated shapes and evidence predicates

- `executed` requires `invocation_identity`, forbids `integrity_evidence`, and commits only when the actual identity equals the consumed ticket's stored triple.
- `not_invoked_integrity_fault` forbids `invocation_identity`, requires the complete labeled `integrity_evidence{expected_identity, observed_identity}` pair, and commits only when expected equals stored, observed is schema-valid, and observed differs from expected.
- Missing, forbidden, malformed, false-expected, equal-pair, or invalid-observed input takes the one no-reply fault disposition with zero ticket and `tool_calls` mutation.
- Equivalent duplicates compare outcome-specific evidence plus the persisted commit epoch. The fixture matrix covers both valid paths, predicate and shape negatives, both retirement race orders, stale delivery against UNKNOWN and VOID, post-epoch equivalent/conflicting duplicates, and the current-sender×UNKNOWN invariant fault.
- The terminal ticket and `tool_calls` state plus validated evidence remain one atomic §F chokepoint transaction. The §F row repeats the storage predicate rather than weakening it.

## Accepted surfaces and scope

The r35 accepted surfaces remain accepted: the wire outcome domain is exactly `executed | not_invoked_integrity_fault`; definite zero invocation is not fabricated as an invocation and is not parked UNKNOWN; `EXECUTED` and `NOT_INVOKED_INTEGRITY_FAULT` are honest tool-call terminals; the operation is one-way; and r34 F82/F83 plus earlier F80 semantics remain intact. `TURN_PARKED_UNKNOWN` remains withdrawn with exactly one occurrence in the design.

This approval authorizes the m-10 closure SITREP and Master routing of the exact approved owner shape/hash to m-9. It does not approve m-9's currently held r16 bytes: m-9 must fold the exact `turn_epoch` carriage, discriminated members, terminal transitions, and evidence predicates, then receive a fresh uniquely-parented review. The affected F73 rebinds, field-grain reciprocal, corrected close supplement, Master+VP interface lock, stage-4/5 work, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change to the m-10 design requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `648cd4dd6585900e640d44cf0eb16cb89104b7fee26f593a6b7e1793d57168c2`.
- Exact m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Exact superseded r35 SHA-256 from the uniquely-parented relay: `dbfcb1bf9d0a2980ba73cc2291f2fd916d90f4bd8c7be44f2c70bcad4d77655b`.
- Exact prior MUST-REVISE relay SHA-256 recomputed: `b8c70c6bde0f8831710a4ba84ca029cc03d9a5e032b7b91e8a21d7dbbd4cfa8a`.
- Incoming DESIGN exact-file lint: OK.
- Targeted sweep: §B.1/B.3/B.4 retirement and sender fencing; §D.1 ticket-state domain; §D.3 three authorities and split identity guard; §D.4 outcome grammar, ordered transition table, evidence predicates, atomic terminal commit, duplicates, retirement races, and fixtures; §F `tool_calls`/`tool_authorizations`; Master route `203401`; m-9 R16-F1 review/SITREP.
- `TURN_PARKED_UNKNOWN` occurrence count in the approved design: `1`, the withdrawal sentence.

ACTIONS_GIT_REF: docs-workspace disk action — created this exact-byte review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-210001.md`.
Next requested action: m-10.planner files the closure SITREP on exact hash `0240e874…`; Master then routes the exact owner-real outcome-record shape/hash to m-9 for its bounded fold and fresh pair review.
