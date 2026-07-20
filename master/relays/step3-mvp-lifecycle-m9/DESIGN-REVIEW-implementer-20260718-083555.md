## DESIGN-REVIEW — m-9 lifecycle half r10 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the finding is a bounded consumer correction against pair-approved m-10 r32 bytes
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: aadf4a277ccc56252d50e78f66f6c6b93e5934ad1f795d53e6f3502711110e34
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260718-083000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-083555.md
SUBJECT: MUST-REVISE exact r10 aadf4a27... — the replay rule collapses durable VOID/expired into a nonexistent same-reason wire replay instead of m-10 r32's required authorize_reject turn_inactive translation

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r10 bytes at SHA-256 `aadf4a277ccc56252d50e78f66f6c6b93e5934ad1f795d53e6f3502711110e34`, not only the r32 consumer-fold loci. The directly addressed relay, `DESIGN_DOC_ID`, incoming exact-file lint, lineage, owner-return prerequisites, and current hashes pass. I re-verified the pair-approved owner bases: m-10 r32 `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`, m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, and m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.

The r10 fold correctly adds and separately disposes the four reply-class rejection tokens plus issue-side `IDENTITY_MISMATCH`; makes `turn_budget_exhausted` the lawful `turn_exhausted` wire carrier; keeps `turn_inactive`, `lease_invalid`, and `run_not_admitted` semantically distinct; charges ISSUED and VOID rows against one ceiling; withdraws the imported `TURN_PARKED_UNKNOWN`; and rebases the current m-10 basis to r32. The prior r9 findings remain closed. One replay mapping contradiction prevents approval of these exact bytes.

## Blocking finding

### R10-F1 — The replay rule does not preserve m-10 r32's durable-reason-to-wire translation

Section 3.3 says that a stored `VOID` row replays “the same `authorize_reject{reason}`/`DENIED_ABOVE_SET`” (`:209`). The §6 fixture repeats `VOID⇒same authorize_reject` (`:262`). Those statements are not total over the owner contract and make the document's exact-wire census ambiguous.

Pair-approved m-10 r32 §D.1 has the durable five-member `void_reason` domain:

`{run_not_admitted, turn_inactive, lease_invalid, denied_above_set, expired}`.

Its §D.2 replay rule has the closed four-member `authorize_reject` wire-reason family and explicitly translates:

`VOID/expired` ⇒ `authorize_reject{turn_inactive}`.

Therefore `expired` is never an `authorize_reject` reason, and a ticket that was originally granted does not replay the same original wire reply after its row is durably expired. The reply is derived from current stored row state through the owner's total mapping. R10's shorthand erases exactly that translation while §5 claims the census imports no token.

Required revision: replace the §3.3 and §6 shorthand with the total r32 mapping:

- `ISSUED` ⇒ the same `ticket_granted{ticket_id}`;
- `VOID/{run_not_admitted, turn_inactive, lease_invalid}` ⇒ the matching `authorize_reject{reason}`;
- `VOID/denied_above_set` ⇒ `DENIED_ABOVE_SET`;
- `VOID/expired` ⇒ `authorize_reject{turn_inactive}`;
- `CONSUMED | OUTCOME_RECORDED | UNKNOWN_TOOL_OUTCOME` ⇒ `DUPLICATE_REQUEST`; and
- same-key replay-identity mismatch ⇒ `IDENTITY_MISMATCH`.

Keep the already-correct no-write, no-state-change, no-supervision replay invariant and the worker dispositions over the actual emitted wire token. Do not add `expired` to the wire family.

## Accepted portions

- The four `authorize_reject` tokens are named as a closed family and disposed individually. In particular, `turn_budget_exhausted` is lawful turn exhaustion, `turn_inactive` is an ordinary terminal race, `lease_invalid` is an invariant-fault path with owner-side atomic retirement, and `run_not_admitted` is an external run-end disposition.
- Issue-side `IDENTITY_MISMATCH`, `DENIED_ABOVE_SET`, and `DUPLICATE_REQUEST` remain separate typed results; the issue-side identity-mismatch row correctly says no execution, row mutation, or supervision.
- Section 2.4 correctly binds the denial budget to m-10's ISSUED-plus-VOID row count and terminates on the row-less at-ceiling token.
- The §3.1 arrow, §5/§7 census, basis rebase to m-10 r32, and removal of `TURN_PARKED_UNKNOWN` are otherwise correct.
- Every r9-approved lifecycle invariant survives this fold: terminal/comparator shape, cancellation cuts, count-once/no-phantom behavior, bare-loss-to-UNKNOWN discipline, no-stream split, F59 consume-before-execute ordering, EOF containment, and the existing m-8/m-3/m-7 seams.

## Revision bar and gate disposition

Return fresh bytes that state the complete stored-row-derived replay mapping at both the normative §3.3 locus and the §6 fixture locus, including the explicit `VOID/expired` to `authorize_reject{turn_inactive}` translation and all three terminal ticket states to `DUPLICATE_REQUEST`.

This verdict is byte-bound to `aadf4a277ccc56252d50e78f66f6c6b93e5934ad1f795d53e6f3502711110e34`. The closure SITREP, fresh m-10 reciprocal over m-9 r10×m-10 r32, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `bef07bc029f410f077cc98a03f2e37ba3c3ec6aeb38b1727334f730fc69695e9`.
- Exact reviewed m-9 r10 SHA-256 recomputed: `aadf4a277ccc56252d50e78f66f6c6b93e5934ad1f795d53e6f3502711110e34`.
- Pair-approved m-10 r32 SHA-256 recomputed: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, m-10 r32 approval, Master r10 release, and prior r9 verdict exact-file lint: OK.
- Live normative sweep confirms the new four-token family and issue-side identity token are present, `TURN_PARKED_UNKNOWN` is absent from the live contract, and only the incomplete replay shorthand contradicts the owner mapping.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-083555.md`.
Next requested action: m-9.planner holds r10 and the closure SITREP; makes only the bounded replay-mapping correction above; then returns one fresh uniquely-parented DESIGN request over the new exact hash.
