## DESIGN-REVIEW — m-9 lifecycle half r13 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the finding is a bounded closure-sequencing correction against the directly addressed master release
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 9d99c56339bc33a2bef742897cbf6d5d03635cb18c6a76e9de64cdbe24439010
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260718-092000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-041909.md
SUBJECT: MUST-REVISE exact r13 9d99c563... — H-14 is correctly consumed, but the live status and §7 require a fresh complete reciprocal where master authorized only the one-item m-10 delta-confirm with 093000 standing for every other item

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r13 bytes at SHA-256 `9d99c56339bc33a2bef742897cbf6d5d03635cb18c6a76e9de64cdbe24439010`, not only the H-14 fold loci. The directly addressed relay, `DESIGN_DOC_ID`, incoming exact-file lint, lineage, master release, reciprocal finding, and current owner hashes pass.

The substantive H-14 fold is correct. Issue-side `STALE_EPOCH` now appears in the §3.1 arrow and §3.3 with m-10 r32's exact §D.2 (2) properties: row-less, counter-neutral, no ticket, no §2a charge, fenced cease, no local retry, await replacement. The row covers both the ordinary stale-authorize race and §D.2 (5)'s at-ceiling crash-after-retirement reclassification. Section 5 enumerates `STALE_EPOCH` on both issue and consume sides and adds the genuinely consumed `consume_ok`; §6 carries both required stale cuts. The r12-approved replay, rejection, cancellation, executor, and lifecycle invariants survive.

A fresh whole-document pass nevertheless finds one live gate-sequencing contradiction. These exact bytes cannot receive final pair approval.

## Blocking finding

### R13-F1 — The live gate asks for a fresh complete reciprocal instead of the authorized one-item delta-confirm

The status block (`:6`) and §7 closure gate (`:299`) say approval advances through “the fresh complete reciprocal over m-9 r13 × m-10 r32.” The r13 fold log (`:317`) repeats that sequence and then also says m-10 files the one-item delta-confirm.

That is not the directly addressed master route. `RECONCILE-orchestrator-planner-20260718-091006:24` says:

- the complete reciprocal `confirm-m10/093000` stands for every other r12 locus;
- after r13 review and SITREP, m-10 files **one one-item delta-confirm** on r13; and
- the corrected close supplement composes those two evidence artifacts.

The incoming r13 DESIGN relay states the same sequence in its subject and review request. Requiring a new complete reciprocal both contradicts that route and obscures which evidence is reusable versus newly owed.

Required revision: make the live status, §7 gate, and r13 fold-log gate state the authorized composition precisely:

1. fresh r13 m-9.implementer approval;
2. r13 SITREP naming the approved hash;
3. m-10's one-item H-14 delta-confirm on r13, with `step3-mvp-confirm-m10/RECONCILE-planner-20260718-093000` incorporated for every other reciprocal item; then
4. master's corrected close supplement.

Do not request or imply a second full reciprocal. Historical descriptions of the already completed r12 × r32 reciprocal remain accurate.

## Accepted portions

- H-14 is closed substantively: issue-side `STALE_EPOCH` is a standalone typed reject outside the four-reason `authorize_reject` family and carries the exact row-less/counter-neutral fenced posture from m-10 r32.
- The §D.2 (5) at-ceiling crash-recovery cut is represented honestly: retirement has minted E+1, the re-ask reaches check (2), and the worker ceases without local retry or charge.
- `STALE_EPOCH` is now enumerated on both issue and consume sides at §5; `consume_ok` is present in the consumed list and already has its semantic gate at §3.2.
- Status, §5, and §7 bind the current bytes to m-9 r13 × m-10 r32; no stale r12 current-review target remains.
- R10-F1, R11-F1, the r12 approval basis, and every earlier lifecycle invariant remain closed.

## Revision bar and gate disposition

Return fresh bytes changing only the three live post-approval sequence statements so they require the one-item m-10 delta-confirm plus incorporation of the standing `093000` reciprocal, followed by master's corrected close supplement.

This verdict is byte-bound to `9d99c56339bc33a2bef742897cbf6d5d03635cb18c6a76e9de64cdbe24439010`. The r13 SITREP, m-10 delta-confirm, corrected close supplement, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `9e4e5d4c3d14ce57e86324b152e1d404385298fdc7db4856ff75ec0bce7808fc`.
- Exact reviewed m-9 r13 SHA-256 recomputed: `9d99c56339bc33a2bef742897cbf6d5d03635cb18c6a76e9de64cdbe24439010`.
- Pair-approved m-10 r32 SHA-256 recomputed: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, master release, m-10 `093000` reciprocal, and prior r12 approval exact-file lint: OK.
- Live sweep confirms every H-14 token/census/fixture companion and the sole contradictory sequence at status `:6`, §7 `:299`, and r13 fold log `:317`.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-041909.md`.
Next requested action: m-9.planner holds r13 and its SITREP; corrects only the live post-approval sequence to one-item m-10 delta-confirm + standing `093000` incorporation + master close supplement; then returns one fresh uniquely-parented DESIGN request over the new exact hash.
