## DESIGN-REVIEW — MUST-REVISE m-10 r35 exact bytes: outcome members are honest, but the one-way transition is not total

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r36
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — two bounded §D.4 totality corrections remain inside the Master-routed F59 outcome-record amendment
GRILL_REQUIRED: no — these findings make the routed owner contract executable; they introduce no architecture choice
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260719-204500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-204600.md
SUBJECT: MUST-REVISE exact r35 dbfcb1bf — the two wire outcomes and definite-no-effect terminal are honest, but §D.4 omits UNKNOWN/VOID/stale-sender branches and does not validate the expected-vs-observed mismatch evidence

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r35 design bytes at SHA-256 `dbfcb1bf9d0a2980ba73cc2291f2fd916d90f4bd8c7be44f2c70bcad4d77655b`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the anti-churn scope pass. The owner-real wire members and storage states answer the core m-9 R16-F1 requirement, but the claimed one-way totality has two blockers.

## Findings

### R35-F1 — the record transition omits reachable ticket states and sender fencing

The ticket state domain is `ISSUED | CONSUMED | OUTCOME_RECORDED | VOID | UNKNOWN_TOOL_OUTCOME`. R35's table handles:

- `CONSUMED` as the fresh commit;
- `OUTCOME_RECORDED` as equivalent/conflicting duplicate;
- `ISSUED` and unknown ticket as channel faults.

It does not classify `UNKNOWN_TOOL_OUTCOME` or `VOID`. `UNKNOWN_TOOL_OUTCOME` is not theoretical: §B.3/§B.4/§D.4 atomically parks a consumed-no-record ticket there during retirement/recovery. A `record_tool_outcome` frame racing that retirement can therefore arrive after the UNKNOWN commit. The contract currently gives that reachable frame no result, despite claiming totality.

The same completion write is epoch-relevant under §B.4 and rides an assign-bound CTRL-W channel, but the record transaction is conditioned only on `state=CONSUMED`. It gives no order for committed-equivalent duplicate versus stale sender, current `CONSUMED` versus stale sender, or a delayed stale frame against an already-parked UNKNOWN row.

Required revision:

1. State the exact frame/envelope carriage (`turn_epoch`) and separate authenticated sender association from durable current epoch, as at consume.
2. Give one ordered, no-reply transition table over malformed/unknown ticket and every ticket state. Preserve committed terminal facts; a delayed record must not silently rewrite `UNKNOWN_TOOL_OUTCOME` or `VOID`.
3. Pin stale-sender behavior: no stale/superseded sender may move `CONSUMED`, `UNKNOWN_TOOL_OUTCOME`, `VOID`, or a terminal row. State whether an equivalent already-committed duplicate precedes stale, consistent with the existing durable-fact replay discipline.
4. Add the record-versus-retirement crash cuts on both sides of the chokepoint commit, plus delayed stale records against UNKNOWN/VOID and equivalent/conflicting duplicates after epoch advance. Each fixture must assert no reply, ticket state, `tool_calls` state, epoch/generation effects, and terminal non-downgrade.

### R35-F2 — the no-invocation evidence pair is stored without a validation predicate

R35 truthfully removes `invocation_identity` on the zero-invocation branch and introduces labeled `integrity_evidence{expected_identity, observed_identity}`. But the transaction is conditioned only on `state=CONSUMED`; it never requires:

- `expected_identity` to equal the consumed ticket's stored identity;
- `observed_identity` to be schema-valid and unequal to expected;
- the outcome-specific fields to be exact.

Without those predicates, a confused sender can persist an arbitrary or equal pair as `NOT_INVOKED_INTEGRITY_FAULT`; the durable row would claim a mismatch that the owner never checked.

Required revision:

1. Close the two discriminated frame shapes: `executed` requires `invocation_identity` and forbids `integrity_evidence`; `not_invoked_integrity_fault` forbids `invocation_identity` and requires the complete evidence pair.
2. For `executed`, retain actual identity = stored ticket. For `not_invoked_integrity_fault`, require expected = stored ticket, observed schema-valid, and observed ≠ expected before the atomic terminal commit.
3. Missing, forbidden, malformed, false-expected, or equal expected/observed members must take one exact no-reply fault disposition with zero ticket/tool-row mutation.
4. Define the equivalent-duplicate key outcome-by-outcome: executed compares its actual identity; no-invocation compares the full labeled pair. Add fixtures for false expected, equal pair, and forbidden/missing conditional members, alongside the valid positive and valid mismatch paths.

## Accepted return

- The wire `outcome` domain is correctly separated from durable states and closed to `executed | not_invoked_integrity_fault`.
- `executed` truthfully represents an invocation that returned; tool success/failure remains result content rather than lifecycle taxonomy.
- `not_invoked_integrity_fault` truthfully represents consumed-but-never-invoked, definite no-effect and is distinct from crash-window uncertainty.
- `invocation_identity` is absent on the zero-invocation branch; the expected/observed labels preserve the defect rather than fabricating an actual invocation.
- `EXECUTED` and `NOT_INVOKED_INTEGRITY_FAULT` are appropriate `tool_calls` terminals; the latter is never parked UNKNOWN after a valid record commit.
- Ticket and tool terminal updates are atomic at the §F chokepoint; the frame is one-way.
- The r34 F82/F83 and earlier F80 surfaces remain semantically accepted and are not reopened by these findings. Their exact r34 approval is superseded by the r35 byte change.

## Scope and remaining gates

Do not file an r35 closure SITREP or route m-9 onto `dbfcb1bf...`. Fold R35-F1/F2 only in the §D.4 outcome-record block and the directly corresponding §F row/schema wording, then return fresh uniquely-parented replacement bytes.

The m-9 fold/review, affected F73 rebinds, field-grain reciprocal, corrected close supplement, Master+VP interface lock, stage-4/5 work, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `e15561f409db1b1096bdcbf8cbf3386730dcd81f58dc55a8c9bc7ccb4e24d739`.
- Exact m-10 r35 SHA-256 recomputed: `dbfcb1bf9d0a2980ba73cc2291f2fd916d90f4bd8c7be44f2c70bcad4d77655b`.
- Exact r34 predecessor SHA-256 retained: `c654204254f2df6193823dc2b7b37164eaa288b138fb799f2699ab41f2807ad2`.
- Exact r34 approval relay SHA-256 retained: `47b4ea9524110776934baab60bbd1b929f600b3097b02373a58f6b56ab3ac6b4`.
- Incoming DESIGN exact-file lint: OK.
- Targeted sweep: §A.1/A.2 one-way channel/frame semantics; §B.3/§B.4 retirement, completion carriage, sender fencing, and UNKNOWN parking; §D.1 full ticket-state domain; §D.3 split guard; §D.4 outcome shapes, transition order, identity predicates, duplicates, crash windows, and fixtures; §F `tool_calls`/`tool_authorizations`; Master route `203401`; m-9 R16-F1 review/SITREP.
- `TURN_PARKED_UNKNOWN` occurrence count in the current m-10 design: `1`, the withdrawal sentence.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-204600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds R35-F1/F2 only and returns a fresh uniquely-parented DESIGN relay; m-9 waits for the eventual pair-approved outcome-record shape/hash.
