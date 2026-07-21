## DESIGN-REVIEW — MUST-REVISE m-10 r38 exact bytes: durable emission and size gating now work, but the operator refusal remains unnamed and the earlier r8 genericity certification is withdrawn

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r39
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded operator-boundary token/shape closure plus the master-offered stage-5 batch reissue; no product or architecture choice reopens
GRILL_REQUIRED: no — the arbitrated disclosure, durable ordering, and size policy stand
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260720-152000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-152100.md
SUBJECT: MUST-REVISE exact m-10 r38 0ebd5c1e — M10-R37-F1 closes and F2's pre-commit size mechanics pass, but the promised typed refusal has no exact token/shape or closed cited predecessor; r8 certification withdrawn because its census falsely says every crash leaves wake pending while r38's post-commit crash leaves it dispatched and re-emits

m-10.planner — I reviewed the exact r38 DESIGN relay at SHA-256 `6654aa7605dcced01df82c75f35a2acef3bd7b0b1d20ded3b803bfd9146caf73` and exact design bytes at SHA-256 `0ebd5c1edf06a0b3cdd38f8cb57bdb2f50401288b5a10ec0016174a9c3d308ba`.

M10-R37-F1 closes. The admission transaction now atomically writes the complete ref, active-turn lease, and wake disposition; `turn_open` emits only after commit from the committed row/snapshot. The post-commit/pre-send crash cut re-emits byte-identically without minting a task identity or consuming the wake twice. §B.2, H-14, the fixture, and §F distinguish durable linearization from frame emission.

M10-R37-F2's substantive size mechanics also pass: the complete canonical frame is constructed or exactly sized before commit, over-limit operator input commits nothing, accepted input remains verbatim, and the exact-fit/one-byte-over fixtures assert no post-admission child/channel fault. One closure requirement remains unmet.

## Finding

### M10-R38-F1 — “the same typed admission-refusal class” does not pin any operator-visible token or shape

r38 says an oversized operator input is rejected at the operator command boundary using “the same typed admission-refusal class as the structural manifest checks.” Neither r38 nor the cited stage-5 census names a closed refusal token, response shape, or concrete existing class. The stage-5 row only says “typed refusal”; the r38 contract contains no `run_start` rejection family to consume. An implementer still cannot know what to return, and a fixture cannot assert an exact result.

Required correction:

1. Name the exact machine-visible operator-boundary refusal token/shape for the over-limit case, including the stable reason member if the family is discriminated; alternatively cite a genuinely closed existing refusal definition by exact artifact/section and show the selected member.
2. Pin its emission point before any admission transaction and its operator-surface consumption (including the non-zero/typed CLI result if that is the chosen surface).
3. Make the one-byte-over fixture assert that exact token/shape in addition to zero durable admission and zero child/channel fault.
4. Preserve the accepted complete-frame sizing, exact-fit behavior, and verbatim/no-truncation rule unchanged.

This is a wire/surface totality fix inside M10-R37-F2, not a new product decision.

## Stage-5 r8 certification correction

I withdraw the `150600` certification that exact stage-5 r8 `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa` is fully shape-generic.

The earlier sweep correctly found no closed member enumeration, but it missed the `m10-turn-admission` census row's unqualified `failure_unknown_semantics`: “crash ⇒ the wake row stays `pending`.” r38 now correctly pins two distinct cuts:

- crash before the admission commit: wake remains `pending`;
- crash after commit but before `turn_open` send: wake is already `dispatched`, and recovery/replacement re-emits without re-consuming it.

The r8 census sentence is therefore false for the newly frozen post-commit cut, and r8 is still byte-bound to superseded r36. Master's `145111` arbitration explicitly offered the batch-reissue path if genericity was refuted. That path is now required: update the stage-5 admission/wake realization and census to distinguish both cuts and bind the eventual approved r39 hash, then obtain a fresh exact-byte stage-5 review. The prior r8 approval remains historical proof for r8 bytes; it cannot certify the amended r39 realization.

## Accepted basis

Everything else in r38 is accepted and need not be redesigned:

- the required closed two-kind `admission_ref`, all wake/operator/replacement branches, worker-owned wake read, and no-authority-transfer boundary;
- durable admission commit followed by post-commit `turn_open`, with byte-identical recovery re-emission and no double wake consumption;
- complete-frame pre-commit sizing, exact-fit admission, one-byte-over zero-side-effect rejection, and verbatim accepted input;
- the three-locus amendment scope and every previously approved r36 surface outside it.

## Scope and remaining gates

Do not file the r38 closure SITREP or route the m-9 consumer fold on `0ebd5c1e…`. Correct only M10-R38-F1 on fresh contract bytes. Also batch the required stage-5 reissue described above; its exact review can follow the contract approval in the owner-ordered sequence.

Fresh m-10 contract approval, stage-5 reissue/review, the amendment SITREP, m-9 consumer fold/review, reciprocal delta, letter rebinds, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming r38 DESIGN relay SHA-256 recomputed: `6654aa7605dcced01df82c75f35a2acef3bd7b0b1d20ded3b803bfd9146caf73`.
- Exact r38 design SHA-256 recomputed: `0ebd5c1edf06a0b3cdd38f8cb57bdb2f50401288b5a10ec0016174a9c3d308ba`.
- Prior r37 MUST-REVISE relay SHA-256 recomputed: `a7c0f42aa6826b76004cc0641a446b834bdc9fff98d9836a52152b22b16ae84c`.
- Exact stage-5 r8 design SHA-256 recomputed: `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`.
- Exact stage-5 r8 approving review SHA-256 recomputed: `4fc928cdfc67b86566bd5eb1c7e3bcdc8149441644338ae97624507ec7f0d259`.
- Incoming r38 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Ordering residue sweep: the prior “in one motion” and “written and emitted in the one admission commit” phrases have zero matches; §B.2 line 73 and §F line 279 now state commit then post-commit emission.
- Size gate proof: §A.2 line 29 pins 4 MiB `FRAME_MAX`; §B.2 line 73 sizes the complete frame before commit, preserves verbatim accepted input, and carries exact-fit/one-byte-over fixtures.
- Refusal-shape proof: the only cited predecessor is stage-5 line 106's generic “typed refusal”; no exact admission-refusal token, response shape, or closed operator rejection family exists in the searched r38/r8 bytes.
- Certification refutation: stage-5 r8 line 120 says crash leaves wake `pending`; r38 line 73 makes post-commit/pre-send crash leave the already-flipped wake disposition and re-emit from the committed row.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, stage-5 doc, historical relay, `frank/` source, branch, commit, lock, SITREP, consumer fold, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-152100.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner pins the exact over-limit operator refusal on fresh contract bytes, then runs the master-offered stage-5 batch reissue to totalize the pre-commit and post-commit crash cuts; all downstream gates wait.
