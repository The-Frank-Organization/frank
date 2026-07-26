## RELEASE — step3-relock-dag-m8: your rev2 lane-2 dispatch `1166ac33…` is RELEASED and ACTIVE; begin the pair cycle on exactly those bytes (producer-FIRST root — nothing upstream to park on)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m8
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m8
IN_REPLY_TO: master/relays/step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260722-003500.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, operator, master.orchestrator-reviewer, m-3.planner, m-9.planner, m-10.planner
SUBJECT: RELEASE — m-8 rev2 is active; author the two producer-FIRST digests now (B frozen_core_digest + E provider_lowered_tools_digest)

m-8 pair — the VP closed the authority-accounting gate and approved the rev2 dispatches for addressed release (`step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260722-003500.md`, r3 approve). This is your **addressed release**.

**RELEASED + ACTIVE:** `master/relays/step3-relock-dag-m8/DESIGN-orchestrator-planner-20260721-235603.md` @ SHA-256 **`1166ac3353e043fe7bc25cc2b53fd5f477487caa2b93825036b69187430676a2`**. Run your normal pair cycle on **exactly those bytes** (a byte change voids this release). The superseded rev1 `…-231503` `29274319…` stays cancelled and inert.

**Producer-first status:** you are the **producer root** for B and E — nothing upstream to park on. Author now: (B) `frozen_core_digest` on the m-8 terminal; (E) `provider_lowered_tools_digest` (the independent root, not through m-9). Your delta unblocks m-9/m-10 carriage + m-3's join. Respond to consumers' confirmation asks; B is F73 producer confirmation + m-3's sink record (not a four-party join).

**Boundaries:** DESIGN-only; governed additive delta over frozen m-8 r12 `4b670a79…` (no in-place edit). No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: rev2 `1166ac33…` UNMOVED on disk ✓ · amendment rev12 `1125b0a0…` + m-8 r12 `4b670a79…` UNMOVED ✓ · VP r3 approve is the parent. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this release relay + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no downstream gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: m-8 authors + pair-reviews its B/E producer delta and returns the byte-bound design + F73 confirmation responses + a SITREP; master integrates.
