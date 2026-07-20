## DESIGN-REVIEW - m-3 r3 bounded turn_epoch fold approved at exact SHA

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m3-review-r4
PARENT_DISPATCH_ID: step3-mvp-design-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair review approves the routed branch-(a) fold; Master+VP still own the interface lock
GRILL_REQUIRED: no - unchanged from the stage-1 dispatch
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-mvp-design-m3/DESIGN-planner-20260717-030000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-10.implementer, m-8.planner, m-7.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve exact r3 SHA-256 70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4 - F-m9-L5-1 branch (a) faithfully aligns the E0 turn_epoch encoding with m-10 and changes no other approved contract semantics

DESIGN_REVIEW_VERDICT: approve

I reviewed the exact r3 bytes of `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` at SHA-256 `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4` against:

- the byte-bound r2 approval at `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`;
- master's branch-(a)/(b) routing in `step3-mvp-confirm-m3/RECONCILE-orchestrator-planner-20260717-024139.md`;
- m-9's source finding `step3-mvp-confirm-m9/RECONCILE-planner-20260717-023800.md`;
- m-10's current §A.2 contract at SHA-256 `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`.

Approve. Branch (a) is faithfully realized and does not disturb the r2 contract.

## Bounded-Fold Proof

The r3 delta is exactly:

1. title/status bookkeeping from r2 to r3;
2. the §2.2 example value `"turn_epoch": 0` to `"turn_epoch": "0"`;
3. one §2.2 normative bullet pinning the string grammar and E0 ceiling;
4. one r3 fold-log entry.

I mechanically reversed only those changes in memory. The reconstructed bytes hash to:

`51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`

That is the exact previously approved r2 hash. No predicate, deny token, policy byte rule, scope matrix row, applicability-evaluator step, annex vector, F65/F68 boundary sentence, or E3 field changed.

## Semantic Review

- The grammar is byte-identical in meaning to m-10 §A.2: `^(0|[1-9][0-9]*)$`, decoded value `< 2^64`, numeric comparison.
- One representation now crosses the CTRL-W `app_event` frame, the opaque `pending_app_events` event copy, and the worker SITREP copy.
- This event field remains descriptive E0/`self_reported`. Uniform encoding does not make it a fence, key, ordering authority, or conductor evidence.
- The event payload string does not alter m-10's separate durable representation for trust-bearing counter columns; no store-boundary ownership is absorbed.
- The E3 record still has no `turn_epoch` field, so the approved F62/F65 applicability schema is unchanged.
- Branch (b)'s permanent exception and two-encoding maintenance burden are correctly absent.

## Downstream Gates

Approval is byte-bound to r3 SHA `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`. Any further byte change requires fresh review.

m-3.planner may file the report-only SITREP naming this hash. The combined refresh round must rebind m-9, m-10, and m-8 references to the new m-3 bytes; m-7 F68 confirmation and the Master+VP composite join remain required.

No interface lock, PLAN, T4, code, credential, provider call, merge, or deploy authority is granted.

## Verification

- Incoming r3 relay exact-file lint: OK.
- r3 SHA-256 reproduced: `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.
- Reverse-delta reconstruction reproduced approved r2 SHA-256 `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- m-10 current design SHA-256 reproduced: `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`.
- `frank/` remained untouched at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design byte, frank source/test, branch, commit, push, PR, merge, credential, or provider action
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `git -C frank status --short` returned none - clean at `502e06cc07b5`
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260717-030100.md`
Next requested action: m-3.planner files the report-only SITREP naming exact approved r3 SHA 70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4 for master's combined refresh round
