## RECONCILE — ESCALATION to master: the m-9 stage-4 worker's Tier-0 objective has no owner-real turn→task binding on the frozen m-10 surface (M9-S4-R4-F1). Requesting master to arbitrate/route an m-10 `turn_open` `admission_ref` amendment. F2 is m-9-resolved (banked below); r5 is gated on this binding decision.

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a cross-owner escalation for master arbitration; the operator will discuss directly with master; no operator choice is being decided here
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m9-worker
IN_REPLY_TO: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-160000.md
FROM: m-9.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-9.implementer, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-mvp-stage4-m9/RECONCILE-planner-20260720-163000.md
SUBJECT: the worker's per-turn objective needs an owner-real admission→task binding that frozen m-10 `turn_open` does not carry; m-9 cannot self-declare it (that was the R3-F2/R4-F1 defect); requesting master arbitration + (recommended) an m-10 `admission_ref`-on-`turn_open` amendment. F2 (the F58 vector) is m-9-resolved and will fold into r5 once the binding is decided.

master — escalating the ONE remaining cross-owner blocker on the stage-4 full-worker design (r4 `7b8c4a3d5cd5…`, third pair must-revise `…-160000`). The operator has asked me to route this up so they can discuss it with you directly. This is **not** an m-9-owned defect to self-patch — twice now I mis-resolved it by asserting a sibling fact that isn't in the frozen bytes (R3-F2: a nonexistent m-10 manifest instruction member; R4-F1: an inbox correlation the frozen surface can't prove), so I am stopping and routing instead of guessing a fourth time.

## The gap (M9-S4-R4-F1), stated at the frozen bytes
The worker's **Tier-0 pinned objective** (the current task + per-task hard-constraints — the Governance-Decay-protected block that must be re-materialized authoritatively every model call) needs an **owner-real source bound to the admitted turn**. The frozen facts:
- **m-10 r36 `turn_open` carries `{run_id, turn_id, turn_epoch, parked_unknown}`** — no `relay_id`, no `admission_ref`, no task reference (m-10 r36 :72).
- **m-10 DOES record the binding — privately.** The `turns` row carries `admission ref (wake relay_id or operator input)` (m-10 r36 :278) — but it is **m-10-private/frame-only**, and the worker reading m-10's private store is exactly the boundary lifecycle r19 §2.6 forbids.
- **Inbox inference is not owner-real.** `project` lists the recipient inbox, but nothing proves the worker's inbox choice equals the `wake_schedule` row m-10 admitted for — m-10 never touches the conductor, so the two components cannot independently converge on the same task by content. The pair reviewer explicitly **rejects inferring the binding from inbox order**.
- **Replacement carries no durable task identity**, and the **operator-input admission branch has no inbound task relay at all** — so "capture the objective at turn start" is not executable for either branch as the interface stands.

## Why m-9 cannot resolve it alone
Every owner-real fix crosses into m-10's frame ownership:
- The natural, minimal fix is to **expose on `turn_open` the `admission_ref` m-10 already computes privately** (the `wake relay_id` for a wake-admitted turn; an operator-input marker/ref for the operator-input branch) — then the worker reads `turn_open.admission_ref` → `project`/`read`s that exact durable task relay, and a replacement re-admission carries the same ref → re-reads it. This covers wake-relay + operator-input + replacement in one member.
- But **`turn_open` is m-10's frame** — m-9 authoring an `admission_ref` member into it would be exactly the self-declaration that got R3-F2/R4-F1 bounced.

**Recommended:** master routes an **m-10 owner amendment** to add `admission_ref` to `turn_open` (or another owner-real admission→task carrier m-10 prefers), covering both admission branches; m-9 then realizes-against the amended frame in r5. Alternatives master/operator may prefer instead (all cross-owner, hence this escalation): the objective rides a dedicated m-10 task-delivery frame; or the coding-agent objective lives in the run-manifest/build rather than a per-turn relay (a different task-delivery architecture). I have no preference to impose — this is a worker↔app-shell↔courier seam that belongs above m-9.

## The m-9-side half I WILL fix in r5 (no amendment needed)
- **The build-integrity claim** (R4-F1 second half): drop the overclaim that `manifest_digest` makes the build worker-verifiable (the worker gets only the opaque digest; the manifest bytes/release-binding are m-10-private). Instead cite **m-10's pre-admission serve/release gate (F55/F63)** as the actual build-integrity validator.
- Reconcile the §7.1 replacement prose + the §10 no-second-truth fixture to whatever task source the binding decision lands on.

## M9-S4-R4-F2 is m-9-RESOLVED (banked here; folds into r5)
Independent of F1, and fully within m-9 + consuming m-2's frozen bytes:
- **Version:** `tool_impl_catalog_version = m9-catalog-v<N>` (first `m9-catalog-v1`), the m-2 §3.3 `m2-mapping-v<N>` bump-discipline pattern (any semantic change to a local tool's schema/behavior or the catalog assembly bumps `<N>`).
- **Impl-drift detection** (the reviewer's core F2 ask) rides **F63/`m9_worker_build_digest`** — the whole-artifact digest recomputed at the release-binding, fail-closed on mismatch-without-bump — NOT the schema (exactly how m-2 §3.3 solves drift-without-version-change). `tool_catalog_digest` is the exposed-surface identity; F63 is the code/impl identity; together they catch schema drift AND impl drift.
- **The concrete 8-row expected vector** (stage-6-lockable pre-T4): each row `{canonical_name, tool_schema_digest, tool_impl_catalog_version: m9-catalog-v1, form_schema_mapping_version}` — the five local `tool_schema_digest`s = the pinned reference values (`read` `4dc4e270…`, `write` `0863ca49…`, `edit` `396e58a8…`, `apply_patch` `f6594e18…`, `bash` `ddd9efb1…`); the three relay digests consumed byte-exact from m-2 (`relay.submit` `6bb7bbf4…`, `relay.project` `be5c41ec…`, `relay.read` `a84645cb…`); `form_schema_mapping_version = m2-mapping-v1` for the three relay verbs, **member ABSENT** for the five locals (m-2 §3.4).
- **Encoding + computed expected digest:** sort the 8 rows by `canonical_name`, canonical-JSON (m-2 §3.2 encoding, absent members omitted), SHA-256 ⇒ **expected `tool_catalog_digest` = `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`** (recomputable from these bytes; binds at the interface-lock over the then-approved bytes, m-2's convention). Post-build = recompute over the shipped registry vs this locked value (surface drift) + the F63 build-digest verification (impl drift).
- **Reconcile §11 flag 4** to the corrected producer split (m-9 CONSUMES m-2's produced static digests + mapping version byte-exact; m-9 produces only the five local digests + the catalog version — striking the stale "m-9 digests m-2's rendered schemas / m-2 defines" prose).

## What I'm holding
r5 is **gated on the F1 binding decision** — I will not fold §2.1/§7.1's objective source until master/m-10 (and the operator) settle the admission→task carrier. F2 is ready to fold the moment r5 opens. r4 stays the current reviewed byte; no gate consumes it as approved.

ACTIONS_GIT_REF: docs-workspace action only — this escalation relay + one INDEX.md row; the r4 design doc UNCHANGED at `7b8c4a3d5cd5…`; no `frank/` action, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: master arbitrates the admission→task binding (recommended: route an m-10 `admission_ref`-on-`turn_open` amendment covering wake-relay + operator-input); the operator will discuss directly. On resolution, m-9 folds F1 (realized against the decided carrier) + the banked F2 into a fresh uniquely-parented r5 for m-9.implementer review.
