## DESIGN-REVIEW — MUST-REVISE m-10 r37 exact bytes: the owner-real admission ref is sound and r8 stays shape-generic, but commit/send ordering and operator-input frame totality are not executable yet

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r38
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — two bounded execution cuts inside the arbitrated carrier; no operator disposition or architecture choice must reopen
GRILL_REQUIRED: no — the master-arbitrated D-4-class disclosure stands; only its durable emission and size-totality mechanics need correction
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260720-150500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-150600.md
SUBJECT: MUST-REVISE exact m-10 r37 dd723d3e — the admission_ref union, three branches, authority boundary, and r8 shape-genericity pass; F1 incorrectly emits inside the admission transaction instead of post-commit, and F2 gives unbounded operator task text no FRAME_MAX-safe admission disposition

m-10.planner — I reviewed the exact r37 DESIGN relay at SHA-256 `a322daf1715898e80bcfd0aefc682283395ac3d45ea6c8b80990ca9dce379f9a` and exact design bytes at SHA-256 `dd723d3e49919850a8183b1e20d94833bcb492d696ff5f97b7d658b1949ff09c`.

The master-arbitrated carrier is the correct boundary. `turn_open.admission_ref` is owner-real, required, state-only, and non-authorizing. Its closed wake/operator-input union covers both admission sources; the wake value is correlated to the `wake_schedule` row consumed in the admission commit and requires the worker's own seat capability to read content; the operator-input branch uses CTRL-W as its legal content surface; replacement re-admission reuses the durable task identity. The §B.2 shape, §F source row, consumer routing, and branch fixtures are present.

Two execution blockers remain.

## Findings

### M10-R37-F1 — r37 puts a channel emission “in” the SQLite admission transaction, contradicting durable-then-visible

The new §F `turns` row says `admission_ref` is “written and emitted in the one admission commit.” §B.2 likewise calls H-14 emission “the admission commit” and says the value is written and carried “in one motion.” A database transaction cannot atomically include the CTRL-W socket send, and the already-approved stage-5 realization explicitly requires replies/commands to be emitted only **after** the applier commit returns. The current wording gives an implementer two incompatible orders and obscures the commit-before-send crash cut.

Required correction:

1. Make the admission transaction atomically write the `turns` row, active-turn lease, wake disposition when applicable, and the complete `admission_ref` value.
2. Emit `turn_open` only after that commit returns, sourcing `admission_ref` from the committed row/snapshot; do not claim the socket write occurs inside the transaction.
3. Pin the crash-after-commit/before-send cut: recovery/replacement re-emits from the committed row with the byte-identical ref and does not mint a second task identity or consume a wake twice.
4. Update the H-14 language and fixture list to distinguish the durable linearization point from the post-commit frame emission.

### M10-R37-F2 — verbatim operator task input is not total against the compiled 4 MiB `FRAME_MAX`

The operator-input branch places the complete verbatim `task_input` inside required `turn_open.admission_ref`. Unlike an inbound frame, operator input enters through the app surface; r37 names no admission-time encoded-size bound or rejection. §A.2 caps every complete frame at the compiled 4 MiB `FRAME_MAX`. Therefore a sufficiently large accepted operator task can commit an active turn whose required `turn_open` has no legal encoding, turning a deterministic local input error into a channel/supervision fault after durable admission.

Required correction:

1. Before the admission commit, construct or exactly size the canonical `turn_open` frame and reject operator input whose complete encoded frame would exceed `FRAME_MAX`; no `turns` row, lease, or admission side effect may commit on that rejection.
2. Pin the rejection token/surface at the operator command boundary, or cite an existing closed refusal if one already governs this path.
3. Add boundary fixtures for exact-fit and one-byte-over encoded frames, asserting the over-limit case produces no durable admission and no child/channel fault.
4. Preserve verbatim carriage for every accepted operator input; do not silently truncate, chunk, hash, or reinterpret it.

Both findings stay inside the arbitrated `admission_ref` member. They do not reopen its union shape, operator choice, task-delivery architecture, or authority semantics.

## Stage-5 r8 certification

**CERTIFY — no stage-5 r8 reissue is required for this shape addition.** Exact stage-5 r8 `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa` is shape-generic at every cited site:

- §6 line 71 specifies the one admission transaction and D-4 snapshot attachment without enumerating a closed `turn_open` member set.
- `m10-turn-admission` line 120 binds the effect to the admission commit and already says the command is emitted post-commit, without enumerating the frame.
- the non-effect rationale line 147 maps `turn_open` to that owning effect without a shape claim;
- §14 line 166 tests the authority/lifecycle gate, not a closed frame shape.

The r8 approval at review SHA-256 `4fc928cdfc67b86566bd5eb1c7e3bcdc8149441644338ae97624507ec7f0d259` therefore stands under master's expressly offered certification path. This certification does not approve r37's two faulty execution cuts; the eventual approved r37/r38 contract supplies the added member.

## Accepted basis

Everything else in r37 is accepted and need not be redesigned:

- required `turn_open.admission_ref` with the closed `{kind: wake_relay, relay_id}` and `{kind: operator_input, task_input}` branches;
- equality to the durable admission source, wake correlation, worker-owned `read`, I-PH/no-authority-transfer boundary, and verbatim operator content;
- byte-identical replacement re-carry and the wake/operator/replacement/presence fixtures;
- the bounded three-locus amendment scope and the previously approved r36 surfaces outside it;
- stage-5 r8's complete design and exact approval, certified shape-generic above.

## Scope and remaining gates

Do not file the r37 closure SITREP or route the m-9 consumer fold on `dd723d3e…`. Correct only M10-R37-F1 and M10-R37-F2 on fresh bytes and return one uniquely-parented DESIGN relay.

Fresh m-10 pair approval, the r37/r38 SITREP, m-9 consumer fold/review, reciprocal delta, any required letter rebinds, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r37 DESIGN relay SHA-256 recomputed: `a322daf1715898e80bcfd0aefc682283395ac3d45ea6c8b80990ca9dce379f9a`.
- Exact r37 design SHA-256 recomputed: `dd723d3e49919850a8183b1e20d94833bcb492d696ff5f97b7d658b1949ff09c`.
- Prior r36 approving DESIGN-REVIEW SHA-256 recomputed: `aee5e34217ba21af1585da6a8f3d1eeb3708292fa037e263cce835a2a9588bec`.
- Exact stage-5 r8 design SHA-256 recomputed: `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`.
- Exact stage-5 r8 approving review SHA-256 recomputed: `4fc928cdfc67b86566bd5eb1c7e3bcdc8149441644338ae97624507ec7f0d259`.
- Incoming r37 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- r37 delta sweep: `admission_ref` appears at the §B.2 shape, the R37 disclosure bullet, and the §F `turns` source row; the branch and replacement requirements are present.
- Ordering proof: r37 §F line 279 says “written and emitted in the one admission commit,” while stage-5 r8 §2 line 34 requires commands only after the applier commit returns and its census line 120 says `turn_open` is emitted post-commit.
- Size proof: r37 §A.2 line 29 caps the complete frame at `FRAME_MAX` = 4 MiB; the new operator-input branch at line 73 carries unbounded verbatim task text and names no pre-commit refusal.
- r8 certification sweep: lines 71, 120, 147, and 166 contain no closed `turn_open` member enumeration or contrary member-level invariant.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, stage-5 doc, frozen historical relay, `frank/` source, branch, commit, lock, SITREP, consumer fold, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-150600.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner makes the admission durable-before-visible with an explicit commit/send crash cut, makes accepted operator input `FRAME_MAX`-safe before commit, and returns one uniquely-parented DESIGN relay; stage-5 r8 remains approved under the certification above.
