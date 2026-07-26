## RECONCILE -- REVISE: r9 closes the r8 seams, but terminal outcomes still fall out of the manifest and resume admission is not frame-total

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev9 is not ready until the manifest and its sole carrier are total
GRILL_REQUIRED: no -- D7, the grain boundary, and no-auto-retry are settled; this return is limited to state-carrier, first-action, frame-size, and proof totality
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-154500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- the r8 marker/composite-settlement and durable-snapshot/disposition seams close; add a closed determinate-terminal carrier/action and make the full resume frame admissible by construction

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-154500.md` at SHA-256 `dd596904d28efe881a127d4f3975de950059720969307c6cb2b58d6f5ff31059`.

Proposed amendment rev9: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `4e2e37506d99f69cdf1d4513734fd705e12901b3255c6bd3ab7ad9e3630b46a7`.

## r8 requested closures -- ACCEPTED

Rev9 closes the two seams returned in r8 at the correct grain:

- The tool content record **and its admitting durable `round_marker`** now linearize before `record_tool_outcome`; the prior settled-but-truncated cut is gone (`STEP-3-STAGE6-AMENDMENT.md:193-198`).
- Provider resume settlement is now an owner-real composite: m-10 canonical terminal **AND** committed m-9 content-ready evidence, with missing evidence prohibited from `settled_providers` (`:199-206`). This removes the terminal-before-worker-content contradiction without changing m-8's canonical terminal timing.
- Full-ancestry source-turn identity is required, and the D1 local-lock versus jointly-owned ordered-segment wording now follows the mechanism branch (`:179-183`, `:186-191`).
- `resume_snapshot` now persists canonical manifest bytes; disposition is separate and starts `PENDING`; m-9 waits for m-10's post-commit disposition receipt before provider, tool, or conductor work; the three report/commit/receipt crash cuts converge (`:220-250`).

Those r8 findings do not need another redesign. The grain boundary remains accepted, and no operator grain arbitration is requested.

## Findings

### F105-D2-R9 -- BLOCKER: the claimed total partition has no carrier or first action for determinate terminal/no-resume outcomes

Rev9 partitions provider rows into content-bearing settled, determinate discard, and uncertain (`:207-210`), but the reconciliation contract only names `settled_providers` and `uncertain`, then declares "absent from both -> not-happened" (`:201-216`). There is no manifest member for the middle class. The D3 "total first-action table" likewise has clean-positive, uncertain-tool, uncertain-provider, and degraded branches only (`:229-236`). A canonical denied/rejected/failed/cancelled row can therefore either disappear into `not-happened` or be treated as clean-positive and permit the normal next attempt. Both readings lose the terminal fact; the latter can silently advance after a cancellation or failure despite the frozen user-requested retry discipline.

The middle-class label is also factually too strong. Rev9 calls all `transport_failed` and `CANCELLED` rows "definite no-content" (`:207-210`), while the frozen m-8 contract permits provider events before terminal `failed` and explicitly defines `cancelled(post_invocation)` as wire-crossed with partial content (`m-8 provider contract :58-67, :88-92, :101`). Discarding those partial bytes from resumed context is conservative and valid; claiming they never existed is not. The truthful class is determinate terminal/no-resume with content intentionally discarded or untrusted, split from the genuinely zero-content deny/local-reject/pre-transport-cancel cases.

The same closure must cover tools. Frozen m-10 has terminal `NOT_INVOKED_INTEGRITY_FAULT`: definite zero invocation, no `invocation_identity`, and no frozen guarantee of a `tool_result` content record (`m-10 seam :243-257`). Rev9's universal `settled => content` invariant cannot silently put that terminal in `settled_tools`; omission would again turn a known integrity terminal into `not-happened`.

Required correction at decomposition grain:

1. Define a **closed manifest union** in which every canonical provider and tool terminal/parked state maps to exactly one carried class. At minimum distinguish content-bearing settled, determinate terminal/no-resume, uncertain/partial, and content-lost/degraded. Absence may mean not-happened only for a source row that truly does not exist, never for an omitted known terminal.
2. Give the determinate terminal/no-resume class a row-identity-exact carrier (including source turn and provider terminal/cancel point or tool integrity terminal as applicable) and a first-action branch. It must surface/terminalize without automatically opening a replacement attempt; any permitted retry remains an explicit fresh user-requested action.
3. For `NOT_INVOKED_INTEGRITY_FAULT`, either carry it as determinate no-effect with `turn_failed` as the action, or require a durable provider-visible error content record plus marker before recording the terminal. Do not call it content-settled without that content fact.
4. Split the per-entry schemas by owner-real source. Tool entries can carry the frozen canonical args digest; provider entries must not require an `args_digest` unless the amendment names its canonical producer and row source.

The pairs still own the exact frame/table encoding. The closed categories, exhaustive mapping, and first actions are master-level acceptance properties because they decide whether canonical outcomes are erased or replayed.

### F105-D3-R9 -- BLOCKER: a committed full-ancestry snapshot can exceed its only legal carrier

Rev9 makes the full-ancestry canonical manifest bytes required on `turn_open` and commits those bytes plus the active-turn lease in the continuation-admission transaction (`:186-191`, `:220-229`). The frozen IPC contract caps every frame at `FRAME_MAX = 4 MiB`; an oversized frame is a channel fault and supervision disposition (`m-10 seam :27-42`). Its existing pre-commit `turn_open` sizing rule was built for `admission_ref` and has only the closed `task_input_frame_overflow` refusal at that boundary (`m-10 seam :72-73`). Rev9 adds an unbounded-by-contract generated member but adds no maximum-size proof, pre-commit branch, or typed failure disposition for it.

That permits a bad durable state: admission, snapshot, lease, and `PENDING` commit successfully, but the sole carrier cannot be emitted. Recovery then derives the same oversized frame from the immutable bytes, repeatedly taking the channel-fault/supervision path instead of reaching disposition inspection. Byte-identical persistence does not help if the bytes have no legal frame.

Required correction: make resume admission frame-total. Choose and pin one of these decomposition-level outcomes while leaving exact encoding/calculation to the pairs:

- Prove from the G-2 ancestry bound, per-turn attempt/tool bounds, entry-size bounds, and all other required `turn_open` members that the maximum canonical frame is always `<= FRAME_MAX`; enforce those bounds before admission.
- Or size the complete candidate `turn_open` before the continuation transaction and take a closed, typed, operator-visible fail-closed/degraded branch with no un-emittable continuation/lease committed. The one-carrier decision rules out silently chunking the manifest unless master explicitly amends that decision.

The exact-fit and one-byte-over proof must include the settlement manifest, `admission_ref`, `parked_unknown`, path, and framing overhead, not only operator task input.

### F106-R9 -- BLOCKER: the new load-bearing conjunction and receipt gates are not yet made mutation-resistant

R8's six-leg count remains closed, but rev9 introduces mechanisms that the current Durability predicate does not discriminate. `xit-dur-1` checks a matching settled `tool_result`; it does not require provider content-ready conjunction ordering, the disposition post-commit receipt gate, or the newly required full resume-frame sizing (`:303-310`). An implementation that emits `settled_providers` from the terminal alone can still satisfy the written tool-only positive predicate.

Required correction: keep six legs, but add structured sub-cuts/expectations under Durability or bind equivalent Tier-HARD owner fixtures into that leg:

- provider terminal committed first, content-ready absent -> not settled;
- content-ready committed first, terminal absent -> not settled;
- both committed in either order -> exactly one settled entry; duplicate/equivalent reports remain idempotent;
- crash before disposition report, after report-before-commit, and after commit-before-receipt -> zero provider/tool/conductor work until the post-commit receipt;
- maximum legal resume frame passes and one byte over takes the selected pre-commit typed branch, never a committed-frame channel fault.

These are proof obligations for rev9's new seams, not a request to author the pairs' internal record grammar in the amendment.

## Gate disposition

- Proposed stage-6 amendment rev9 `4e2e3750...`: REVISE on the manifest/action and sole-carrier totality defects above; not ready for operator re-scope ratification.
- The r8 marker-before-outcome, provider composite-settlement, source-turn identity, immutable snapshot bytes, `PENDING` disposition, post-commit receipt gate, and D1 ownership correction are accepted.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r10 over new amendment bytes that: (1) make the settlement manifest a closed exhaustive carrier, truthfully distinguish zero-content from determinate-discarded partial content, and add the missing first actions; (2) make the complete persisted `turn_open` frame admissible by construction or take a typed pre-commit failure branch; and (3) bind mutation-resistant proof cuts for the new provider conjunction, disposition receipt, and frame boundary. Preserve every accepted r9 closure and keep record/lock/segment/rotation internals delegated under F73.

## Verification

- Target SHA-256: `dd596904d28efe881a127d4f3975de950059720969307c6cb2b58d6f5ff31059`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1888.
- Amendment rev9 SHA-256: `4e2e37506d99f69cdf1d4513734fd705e12901b3255c6bd3ab7ad9e3630b46a7`; VP r8 parent SHA-256: `2916d7232e6e20a35227142264a082b93ed31861cde82161415b9668ff688c30`.
- Relevant frozen bases recompute unchanged: m-8 provider `4b670a79...`; m-9 lifecycle `4d3bd14e...`; m-9 worker `cb7ff970...`; m-10 seam `d2ce9831...`; m-10 control plane `6fd1d655...`; MVP amendment `2f75f2a1...`.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update.
Next requested action: master folds only the closed-manifest/first-action, complete-frame totality, and proof-cut deltas, then returns amendment rev10 for decomposition review r10; operator re-scope ratification remains held.
