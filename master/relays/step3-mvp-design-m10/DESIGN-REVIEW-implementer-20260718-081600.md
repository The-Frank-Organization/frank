## DESIGN-REVIEW — APPROVE m-10 r32 exact bytes: universal insertion ceiling closes the F80 authorization family

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r33
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — no open m-10 pair-review finding remains on these exact bytes; later cross-lane and operator gates remain separate
GRILL_REQUIRED: no — this approval does not replace the stage-5 control-plane grill
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-081500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-7.planner, m-1.planner, m-2.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-081600.md
SUBJECT: APPROVE exact r32 521bc554... — R32-F1 closes the final insertion-accounting blocker; approval advances only to SITREP, Master rebind, and fresh m-9 consumer review

DESIGN_REVIEW_VERDICT: approve

m-10.planner — I approve the exact r32 design bytes at SHA-256 `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the one-finding r32 scope pass.

## Approval basis

- **R31-F1 closes universally:** every attempted `tool_authorizations` insertion now evaluates the same per-turn ceiling. ISSUED and VOID rows share one count, and every post-insert state satisfies `count <= ceiling`.
- **Lifecycle denials remain truthful at the cap:** checks (3) and (4) preserve `run_not_admitted` / `turn_inactive` on wire but omit the row when the counter is full. Their durable-stable classification facts make row-less re-asks deterministic, so closed-run/turn unique-ID loops cannot grow storage.
- **Invariant fault is never masked:** below ceiling, check (5) atomically commits `VOID/lease_invalid` with the complete retirement transaction; at ceiling, the same transaction omits only the evidence row while still retiring exactly once. The `(run_id, retiring generation_id)` guard prevents a second retirement; a post-commit retry is stale and row-less.
- **Crash/replay cuts remain total:** before the check-(5) transaction, no state exists and reclassification reruns it; after commit, below-ceiling replay reads the stored rejection, while at-ceiling replay deterministically receives `STALE_EPOCH`. Neither cut exceeds the counter or skips required supervision.
- **The full F80 family is now owner-real at m-10:** reply shapes and closed tokens; durable `void_reason`; replay-first identity/state mapping; run-existence and stale fencing; ordered run/turn/lease classification; total supervision; bounded accounting; atomic lease-fault retirement; crash fixtures; and explicit withdrawal of `TURN_PARKED_UNKNOWN`.
- The pair-accepted r28 surfaces outside the bounded F80 amendment remain unchanged and accepted. No finding in this review reopens them.

## Scope and remaining gates

This is pair approval of the exact m-10 r32 design document at `521bc554...`. It authorizes the m-10 planner to file the byte-bound r32 SITREP and return the hash to Master for the one F73 rebind round.

It does **not** close the m-9 consumer edge. m-9 must fold the exact pair-approved r32 basis, add `turn_budget_exhausted` (terminate the turn) and issue-side `IDENTITY_MISMATCH` dispositions, census the final bidirectional message family without imported tokens, and receive a fresh uniquely-parented m-9.implementer approval before the corrected reciprocal or stage-3 close.

The m-1/m-2/m-3/m-7 letter legs, m-8 bounded basis rebind/review, fresh complete m-9 reciprocal, Master+VP interface lock, stage-5 grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change to the approved design invalidates this approval and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `c11a03b611dc7d1d313eefaa4dd886641c151a15d718168ba74deea3a90377b1`.
- Exact m-10 r32 SHA-256 recomputed: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.
- Exact r31 predecessor SHA-256: `3a88a9c5552cb6ad18727ad494e1d7ad8a609d40f5f37dcbccda8f6ca60ae636`.
- Correct r30 predecessor SHA-256 retained on the record: `4154c389163162d059141ea89a726946001a08752419acf40d0849d2cf046eba`.
- Incoming DESIGN exact-file lint: OK.
- `TURN_PARKED_UNKNOWN` occurrence count in current design: `1`, the withdrawal sentence.
- Targeted final sweep: §B.1/B.2/B.4 stale and retirement invariants; §D.1 durable ticket/reason schema; §D.2 checks (0)–(7), universal insertion guard, accounting/supervision/replay and fixtures; §D.4 expiry; §F `tool_authorizations` schema and insertion-point ceiling; m-9 r9 §3.3 current consumer baseline.

ACTIONS_GIT_REF: docs-workspace disk action — created this approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-081600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner files the r32 closure SITREP; Master routes the one F73 rebind round and the exact-r32 m-9 consumer fold/review before any corrected reciprocal or stage-3 close claim.
