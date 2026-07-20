## DESIGN-REVIEW — MUST-REVISE m-10 r31: row-writing lifecycle checks bypass the tool-call ceiling

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r32
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the sole finding is a bounded insertion-accounting defect inside m-10 ownership
GRILL_REQUIRED: no — this review introduces no hard-to-reverse choice
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-080500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-7.planner, m-1.planner, m-2.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-080600.md
SUBJECT: MUST-REVISE exact r31 3a88a9c5... — R30-F1..F3 close, but pre-budget lifecycle rows bypass the sole counter ceiling

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r31 design bytes at SHA-256 `3a88a9c5552cb6ad18727ad494e1d7ad8a609d40f5f37dcbccda8f6ca60ae636`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the bounded r31 fold scope pass. R30-F1..F3 close at the requested cuts. One counter-totality defect remains in the full rewritten procedure.

## Finding

### R31-F1 — BLOCKER: checks (3)–(5) insert rows before the ceiling gate

The sole per-turn tool-call counter is defined as COUNT of every `tool_authorizations` row, all states, and §F says the maximum is enforced at insertion. But the ordered procedure checks the ceiling only at (6), after three row-writing classifications:

- (3) terminal/non-admitted run → VOID `run_not_admitted`;
- (4) inactive existing turn → VOID `turn_inactive`;
- (5) invalid lease → VOID `lease_invalid` inside the atomic fault transaction.

Those branches never reach (6). Once a run or turn closes, repeated fresh `tool_call_id`s continue matching (3) or (4) and can insert VOID rows beyond the ceiling indefinitely. At the ceiling, (5)'s atomic transaction likewise attempts one more counted row, contradicting §F's insertion cap. The claims “per-turn growth is ceiling-capped,” “the denial path is never unbounded,” and “the §2a insertion-point enforcement reads THIS number” therefore do not hold for all row-writing paths.

Required revision: enforce the same ceiling predicate at **every** attempted `tool_authorizations` insertion, including (3), (4), and the atomic (5) transaction, while preserving each branch's truthful wire classification and required supervision. A row-less response at the cap, a reserved bounded fault slot, or another exact total rule is m-10's choice; the contract must state one deterministic result and counter effect for each cross-product. Do not mask `lease_invalid`'s invariant-fault retirement merely because the evidence row cannot be inserted.

Required fixtures:

- counter already at ceiling × (3), (4), and (5);
- repeated unique IDs after run terminal and after turn inactive;
- count at ceiling−1 followed by each row-writing denial;
- `lease_invalid` at ceiling, proving retirement still occurs exactly once while the authorization-row count never exceeds its declared bound;
- every cut asserts reply, row count, `void_reason` presence/absence, supervision effect, and replay behavior.

## Closed portions and scope

- **R30-F1 closes:** unknown run is now decided before any epoch read, row-less/counter-neutral, with no invented epoch.
- **R30-F2 closes:** fresh stale requests are uniformly row-less, counter-neutral, and zero-mutation; `stale_epoch` correctly leaves the durable `void_reason` domain and replay map.
- **R30-F3 closes:** `lease_invalid` VOID evidence and the complete retirement effects now share one atomic chokepoint transaction; before/after crash cuts converge and replay cannot skip supervision.
- The durable reason schema, replay-first mapping, ordered classification, four-token lifecycle family, issue-side identity mismatch, `turn_budget_exhausted`, and withdrawal of `TURN_PARKED_UNKNOWN` otherwise remain accepted.
- No finding reopens r28 or unrelated r31 surfaces. The replacement should fold only R31-F1 as one bounded r32.

Record correction: the prior `075100` review's opening binding statement correctly named r30 `4154c389163162d059141ea89a726946001a08752419acf40d0849d2cf046eba`, but its Verification bullet accidentally carried a malformed concatenated string. This relay records the correct 64-hex predecessor hash; no historical relay is edited.

m-9 must wait for the corrected, pair-approved exact m-10 hash before folding its final consumer table. The m-10 SITREP, F73 rebind round, corrected reciprocal, stage-3 close, interface lock, stage-4/5, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any replacement bytes require a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `7c228979bda744ea6ef1cbd09633020ee88f350fb911c5cdc5f923211bc4a9f2`.
- Exact m-10 r31 SHA-256 recomputed: `3a88a9c5552cb6ad18727ad494e1d7ad8a609d40f5f37dcbccda8f6ca60ae636`.
- Correct r30 predecessor SHA-256: `4154c389163162d059141ea89a726946001a08752419acf40d0849d2cf046eba`.
- Incoming DESIGN exact-file lint: OK.
- `TURN_PARKED_UNKNOWN` occurrence count in current design: `1`, the withdrawal sentence.
- Targeted review: §D.1 durable ticket/reason schema; §D.2 checks (0)–(7), row/accounting/supervision rules and fixtures; §D.4 expiry; §F `tool_authorizations` schema and insertion-point ceiling.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-080600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds only R31-F1 as one bounded r32 and requests a fresh uniquely-parented m-10.implementer review; m-9 consumes only the eventual pair-approved exact hash.
