## DESIGN-REVIEW — APPROVE m-10 r34 exact bytes: sender-bound consume closes F82/F83 on the owner side

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r35
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — no open m-10 pair-review finding remains on these exact bytes; cross-owner and later operator gates remain separate
GRILL_REQUIRED: no — this approval does not replace the stage-5 control-plane grill
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260719-195500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-195600.md
SUBJECT: APPROVE exact r34 c6542042 — R33-F1 closes; m-10's F82/F83 owner amendment is pair-approved, advancing only to SITREP + Master-routed m-9 fold/rebind/reciprocal sequence

DESIGN_REVIEW_VERDICT: approve

m-10.planner — I approve the exact r34 design bytes at SHA-256 `c654204254f2df6193823dc2b7b37164eaa288b138fb799f2699ab41f2807ad2`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the one-block R33-F1 scope pass.

## Approval basis

- **The consume request is constructible and §B.4-conformant:** `consume_ticket{ticket_id, turn_epoch, canonical_tool_name, canonical_args_digest}` carries the presented epoch plus the invocation identity. The sender generation/run comes independently from its private assign-bound CTRL-W channel, while durable current epoch remains m-10-owned authority.
- **The success predicate fences the sender, not merely the row:** consume requires stored epoch = presented epoch = durable current epoch, the channel-authenticated generation to hold the run's current lease, `state=ISSUED`, and exact name/digest equality. A stale generation naming a current ticket cannot update it.
- **Zero-update classification is total:** unknown ticket and above-current presentation are explicit no-reply channel faults with generation FAILED/supervision effects; stale channel association or below-current presentation returns `STALE_EPOCH`; a proven-current sender naming a non-ISSUED ticket receives `DUPLICATE_CONSUME`; only the remaining name/digest mismatch receives `IDENTITY_MISMATCH`.
- **Stale sender and stale ticket are distinct:** epoch mint atomically converts an old ISSUED ticket to `VOID/expired`, so a current sender naming that row lands on `DUPLICATE_CONSUME`, while the ratified stale-worker case lands on `STALE_EPOCH`. The ordered overlap table is consistent with that construction.
- **The fault paths are observable:** both fault-only branches state that no reply frame exists, invoke the existing channel-fault/§B.3 supervision path, mutate no ticket row, and have reply-or-fault fixtures asserting supervision and execution count.
- **The mutation guard remains correctly split:** pre-consume identity mutation is rejected by m-10's wire-vs-row match; post-consume/pre-invocation mutation remains m-9's executor re-derivation and actual-invocation-capture half.
- **F83 remains closed:** check (6) is the sole at-ceiling winner; check (7) has no at-ceiling form; rule and fixture matrix agree on row-less `turn_budget_exhausted`, unchanged count, and lawful turn termination.
- The previously accepted F80/r32 surfaces outside this bounded F82/F83 amendment remain semantically intact. `TURN_PARKED_UNKNOWN` remains expressly withdrawn.

## Scope and remaining gates

This is pair approval of the exact m-10 r34 owner document at `c6542042...`. It authorizes the m-10 planner to file the byte-bound r34 SITREP and return this hash to Master.

It does **not** close the m-9 edge. The current m-9 document still emits the superseded one-field `consume_ticket{ticket_id}` shape. m-9 must consume the exact approved four-field shape, bind the presented epoch to its assign value, preserve its executor re-check and no-reply channel-fault handling, rebase to `c6542042...`, and receive a fresh uniquely-parented m-9.implementer approval.

The affected F73 rebinds (including the bounded m-8 basis review), fresh complete m-9/m-10 reciprocal, corrected close supplement, Master+VP interface lock, stage-4/5 work, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change to the approved m-10 design invalidates this approval and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `76eb4d7240629bd3ccaa907503a7b8ee54fa12850db7453bd3104e05e426a75b`.
- Exact m-10 r34 SHA-256 recomputed: `c654204254f2df6193823dc2b7b37164eaa288b138fb799f2699ab41f2807ad2`.
- Exact r33 predecessor SHA-256 retained: `0b637356dbe8cf9ab322c9dc13ba25adfb3c380239c1161b519136c6bf840cee`.
- Exact R33-F1 review relay SHA-256 retained: `4d8e19dac96610d8ba19e27576449bad51a9517e82a21d7eae3aecc76a76287c`.
- Incoming DESIGN exact-file lint: OK.
- The old three-field m-10 consume shape has zero matches; the four-field shape has one normative match. The current m-9 document still has one one-field match, correctly leaving the consumer fold open.
- `TURN_PARKED_UNKNOWN` occurrence count in the current m-10 design: `1`, the withdrawal sentence.
- Targeted final sweep: §A.1/A.2 channel identity and frame semantics; §B.3 fault effects; §B.4 carriage, sender association, epoch authority, and ticket expiry-at-mint; §D.1 identity; §D.2 retained F83 rule/fixtures; §D.3 success predicate, all zero-update branches and overlaps, fault observability, mutation cuts, and fixtures; §D.4 crash windows; amendment `:61/:112`; grill record `024350:24-27`; current m-9 r14 §3.2/§3.3.

ACTIONS_GIT_REF: docs-workspace disk action — created this approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-195600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner files the r34 closure SITREP; Master routes the exact-r34 m-9 consumer fold/review, affected F73 rebinds, fresh reciprocal, and corrected close supplement before any stage-3 close claim.
