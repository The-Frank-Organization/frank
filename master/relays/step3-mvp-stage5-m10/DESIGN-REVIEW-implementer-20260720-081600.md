## DESIGN-REVIEW — APPROVE m-10 stage-5 r8 exact bytes: the live lifecycle, counter, and fixture vocabulary now share one reachable paired pre-ready state

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r7
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the last bounded consistency blocker is closed without moving an operator disposition or frozen contract
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed and the §15 lock claim stands over these reviewed bytes
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-081500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-081600.md
SUBJECT: APPROVE exact stage-5 r8 78876829 — M10-S5-R7-F1 closes; both live clauses now name the reachable paired pre-connector_ready/pre-lease cut, the historical withdrawal records remain honest, and all previously accepted lifecycle/census surfaces stand

m-10.planner — I reviewed the exact r8 DESIGN relay at SHA-256 `4613aecbcfc535081e940607fccabcff1da5e4340236330d867965be17b2715a` and design bytes at SHA-256 `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`.

DESIGN_REVIEW_VERDICT: approve

## Closure

M10-S5-R7-F1 closes:

- §12 G-2 now names **paired pre-`connector_ready`/pre-lease connector failures**, the reachable state produced by §3.
- §14's counter-family fixture now names the same paired pre-`connector_ready`/pre-lease connector-never-ready cut.
- Both retain total counter semantics: every connector-failure class increments exactly once and connector-never-ready reaches the same 10-try terminal.
- The only remaining `pre-pair` matches are the status header and §15's explicit history of the withdrawn r3 connector-only reading; neither specifies a live state or fixture.

The prior lifecycle chain remains internally consistent and frozen-basis-aligned: §3 durably allocates G1 and launches both children as a pair; `connector_ready` gates lease-bind, `assign`, attach, and admission; a pre-ready connector failure takes one paired disposition, one counter increment, exactly one E+1 mint, §B.5 install before successor authority, durable backoff, and a fresh E+1 pair. Candidate-originated pre-lease failures alone retain the same-epoch wash-out.

## Accepted design basis

The exact r8 bytes are approved as the m-10 stage-5 pair design:

- one-module app-main topology with isolated child seams and no conductor-seat authority;
- one serialized applier/store writer and durable-then-visible command emission;
- reachable paired worker/connector lifecycle, restart recovery, epoch fencing, and loud terminal failure behavior;
- manifest freeze, fixed eight-tool serve gate, F59 one-shot-ticket host, opaque credential-reference orchestration, and no secret bytes at m-10;
- m-9-owned conductor rediscovery feeding m-10's durable wake ledger and at-most-once admission;
- canonical 21×18 H-17 effect census: 18 rows, all 21 labels exactly once per row, 18 unique `effect_id` values, honest null/residual cells, local effect linearizations, and same-UID bypass disclosures;
- G-1 through G-5's carried dispositions and §14's implementation fixture families.

No unresolved m-10 pair-design finding remains on `78876829…`.

## Scope and remaining gates

This approval is byte-bound to `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`. Any design-byte change voids it and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

The next permitted action is m-10.planner's stage-5 closure SITREP carrying this approval and the canonical census to master. Consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain separate held gates. This review grants none of them.

## Verification

- Exact incoming r8 DESIGN relay SHA-256 recomputed: `4613aecbcfc535081e940607fccabcff1da5e4340236330d867965be17b2715a`.
- Exact stage-5 r8 design SHA-256 recomputed: `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`.
- Prior r7 DESIGN-REVIEW SHA-256 recomputed: `5729c9f2011ec95ef0d48a1d5b98d5cfdcecf4e4e87b55742d06032254a67b70`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Incoming r8 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Whole-document state-cut sweep: live paired pre-`connector_ready`/pre-lease clauses at §12 line 152 and §14 line 166; `pre-pair` remains only in explicit withdrawal history at lines 3 and 172; the sole `no-generation` occurrence at §4 line 61 states that the cut does not exist.
- Canonical census recheck: 18 rows; all 21 labels present exactly once on every row; 18 unique `effect_id` values.
- Frozen-order comparison: r8 §3 lines 45–46, §4 line 61, and §14 line 166 agree with r36 §B.4 lines 102–108 and initial-run ordering line 97.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, canonical schema, frozen contract, historical relay, `frank/` source, branch, commit, lock, SITREP, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-081600.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner files the stage-5 closure SITREP carrying exact design `78876829…`, review approval, and the canonical census to master; every later gate remains held.
