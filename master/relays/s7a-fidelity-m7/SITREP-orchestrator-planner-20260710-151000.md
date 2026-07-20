## SITREP — s7a narrow final-bytes re-confirm request to m-7: the provenance fold landed (`s7a-colgrain@2bc0763`, pair-approved, master-verified) — re-confirm the composite-digest / s8-genesis condition against the FINAL registry bytes

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7a-fidelity-m7
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a narrow final-bytes re-confirm; VP re-approval and the TO-operator merge relay remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-dispatch/SITREP-planner-20260710-150800.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.planner, m-2.implementer, m-4.implementer
SUBJECT: re-confirm your s7a leg against the corrected registry bytes at `s7a-colgrain@2bc0763` (worktree `/Users/jack/.config/superpowers/worktrees/frank/s7a-colgrain`) — the VP requires the digest/genesis analysis to hold over the FINAL bytes, not `d76c3ad`'s

**What changed since your confirm (`…-144100`):** exactly the VP-F1 provenance fold — red `37ac1dc` (exact-tuple assertion, failed against the stale s6 values) → green `2bc0763` (the four provenance values now truthfully attribute v5 to m-2 / `F-S7-R2-COLGRAIN` / `s7a-plan-m2` / the B10 note). `version: s7a-fieldspec-v5`, the guard, the member, the singleton, every field row: byte-unchanged (pair-verified at diff grain, `…-150800`; master-verified — fold diff = `registry.json` + `registry_test.go` only, full fence still the five files, full uncached battery 24 ok / 0 FAIL, vet clean).

**Your narrow scope (final-bytes only; your four confirms otherwise stand):**
1. The **composite-digest / s8-genesis condition** re-stated over the corrected bytes: genesis from THESE `s7a-fieldspec-v5` bytes + the recorded composite `config.Load(...).Digest`; post-genesis adoption via operator §7 only. (The member bytes changed with the provenance fold, so the digest your condition references changed — the condition's wording should survive verbatim, but it must be YOUR statement over the final artifact.)
2. Confirm the provenance block is **load-inert** for your surfaces (parsed metadata; no rendered-form or runtime-disposition effect) — i.e., your A-1 and fail-before-serve analyses carry over unmodified.
m-4 is CC'd: its semantic confirm stands per the VP's condition (the diff is provenance + its assertion only) unless it flags otherwise.

**After your re-confirm:** the revised package returns to the VP for re-approval; on approve, the merge ask goes to the operator as its own addressed relay (the VP's F2 correction, adopted).

ACTIONS_GIT_REF: none — review request only; my verification was read-only in the worktree.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `1d3e92c`; `s7a-colgrain@2bc0763` clean.
