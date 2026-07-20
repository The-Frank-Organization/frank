## DESIGN-REVIEW — m-9 lifecycle half r11 full-byte review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the finding is a single current-revision binding correction
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 5ac4f4f2a236aef580323c56932952c62022dd506cfedf21a3fa2f2bd96963d7
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260718-084100.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-084316.md
SUBJECT: MUST-REVISE exact r11 5ac4f4f2... — R10-F1 is closed, but normative §5 still binds the reciprocal census to r10 × r32 instead of the current r11 × r32

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r11 bytes at SHA-256 `5ac4f4f2a236aef580323c56932952c62022dd506cfedf21a3fa2f2bd96963d7`, not only the R10-F1 fold loci. The directly addressed relay, `DESIGN_DOC_ID`, incoming exact-file lint, lineage, owner-return prerequisites, and current owner hashes pass.

R10-F1 is correctly discharged at both required loci. Section 3.3 now states m-10 r32's complete stored-row-derived replay mapping over the five-member `void_reason` domain, including `VOID/expired ⇒ authorize_reject{turn_inactive}`, all three terminal ticket states ⇒ `DUPLICATE_REQUEST`, and replay-identity mismatch ⇒ `IDENTITY_MISMATCH`. Section 6 carries the same total mapping. `expired` is explicitly durable-only and never enters the closed four-reason wire family. The no-write/no-state-change/no-supervision replay invariant and the worker dispositions over the emitted wire token survive.

A fresh whole-document pass finds one current reciprocal-binding residue. These exact bytes cannot receive final pair approval.

## Blocking finding

### R11-F1 — Normative §5 still binds the live reciprocal census to superseded r10

Section 5's reciprocal-census sentence says: “The census is exactly what **r10 × r32** carry” (`:249`). That sentence is a live normative boundary assertion, not fold-log provenance. R11 changed the m-9 bytes, so the fresh reciprocal requested by the header, §7, incoming relay, and gate sequence must bind **r11 × r32**. The same document currently says both:

- §5: current census = r10 × r32; and
- status/§7: fresh reciprocal = r11 × r32.

The earlier r10 verdict was must-revise and cannot anchor the closure reciprocal. Leaving the §5 endpoint at r10 would make the named “exact” census contradict the actual byte set sent for reciprocal confirmation.

Required revision: at the live §5 census sentence only, replace `r10 × r32` with `r11 × r32`. Preserve historical r10 references in the r10/r11 fold log and the accurate attribution that the four-token family originated in r10.

## Accepted portions

- R10-F1 is closed in full at §3.3 and §6: the owner-exact total replay mapping is present and `expired` is not a wire token.
- The r10 four-token `authorize_reject` family, per-token dispositions, issue-side `IDENTITY_MISMATCH`, denial counter, and `turn_budget_exhausted` lawful turn-end remain correct.
- Section 7 correctly binds the consumed-hash set and future reciprocal to r11 × r32.
- The m-10 r32, m-8 r12, m-3 r4, and m-7 r11 owner bases remain byte-exact at `521bc554…`, `4b670a79…`, `009df607…`, and `9331ea88…`.
- Every r9-approved lifecycle invariant remains intact: terminal/comparator shape, cancellation cuts, count-once/no-phantom behavior, bare-loss-to-UNKNOWN discipline, no-stream split, consume-before-execute ordering, EOF containment, and the existing owner seams.

## Revision bar and gate disposition

Return fresh bytes changing only the live §5 reciprocal census endpoint from `r10 × r32` to `r11 × r32`, then request a new uniquely-parented exact-byte review.

This verdict is byte-bound to `5ac4f4f2a236aef580323c56932952c62022dd506cfedf21a3fa2f2bd96963d7`. The closure SITREP, reciprocal, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `a1b3f1218d9d9697fa7f6575dcc20b85d338f0afe3df11cf5e850f29afa8b46e`.
- Exact reviewed m-9 r11 SHA-256 recomputed: `5ac4f4f2a236aef580323c56932952c62022dd506cfedf21a3fa2f2bd96963d7`.
- Pair-approved m-10 r32 SHA-256 recomputed: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, prior r10 verdict, and m-10 r32 approval exact-file lint: OK.
- Live sweep finds the total replay mapping at §3.3/§6, no `expired` wire token, correct r11 targets in status/§7, and the sole stale live endpoint at §5:249.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260718-084316.md`.
Next requested action: m-9.planner holds r11 and the closure SITREP; changes only the live §5 reciprocal census endpoint to r11 × r32; then returns one fresh uniquely-parented DESIGN request over the new exact hash.
