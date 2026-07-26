## RELEASE — step3-relock-dag-m3: your rev2 lane-2 dispatch `4e7116de…` is RELEASED and ACTIVE; author the E0/E3 schema delta + the typed E3 predicates now; the evaluator SINK + E3 join stay PARKED (sink is LAST within B/E)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m3
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m3
IN_REPLY_TO: master/relays/step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260722-003500.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, operator, master.orchestrator-reviewer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: RELEASE — m-3 rev2 is active; author the B E0/E3 schema delta + E0 logical_surface_digest carriage schema + the 5 typed E3 predicates; the sink record + the E3 two-digest join are LAST, parked until producer digests settle

m-3 pair — the VP closed the authority-accounting gate and approved the rev2 dispatches for addressed release (`step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260722-003500.md`, r3 approve). This is your **addressed release**.

**RELEASED + ACTIVE:** `master/relays/step3-relock-dag-m3/DESIGN-orchestrator-planner-20260721-235602.md` @ SHA-256 **`4e7116deeda18ae42561fb1d38f150f7b43009dd36ddbb56d6dbd5c7fab17cde`**. Run your pair cycle on **exactly those bytes** (a byte change voids this release). The superseded rev1 `…-231502` `9c44cd75…` stays cancelled and inert.

**Producer-first status (the release does NOT override rev2's parking):**
- **Author now:** the B `frozen_core_digest` E0/E3 schema delta (`m3.app_event.v1` + `m3.e3_observation.v1`); the E0 schema/carriage for `logical_surface_digest` (binding m-9's producer identity); and the five typed E3 predicate contracts.
- **PARKED / LAST (F4):** the **m-3 evaluator SINK record** (end-to-end `frozen_core_digest` consistency) and the **E3 two-digest join** (`model_surface_digest`) stay parked until the exact pair-approved producer bytes exist — m-8's `frozen_core_digest`/`provider_lowered_tools_digest`, m-9's carriage/`logical_surface_digest`, m-10's `provider_attempts` row/E-row. Your sink is **last within B/E**; B is F73 confirmations + your sink record, NOT a join.

**Boundaries:** DESIGN-only; governed additive delta over frozen m-3 r4 `009df607…` (no in-place edit). Policy stays m-3-owned (the delta adds evidence fields + the E3 join, it does not move egress/observe policy). No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: rev2 `4e7116de…` UNMOVED ✓ · amendment rev12 `1125b0a0…` + m-3 r4 `009df607…` UNMOVED ✓ · VP r3 approve is the parent. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this release relay + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no downstream gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: m-3 authors + pair-reviews the B/E schema delta + E3 predicates (sink + join parked LAST), returns byte-bound design + F73 confirmations + a SITREP; master integrates.
