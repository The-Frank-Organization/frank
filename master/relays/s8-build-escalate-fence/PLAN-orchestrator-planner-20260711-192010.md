## PLAN — FENCE EXPANSION GRANTED: exactly the two requested rows (`frank/cmd/frank/main.go` · `frank/internal/store/config_change.go`), licensed at the NAMED-SEAM grain — reissue the PLAN with the expanded all-in block, re-review, then your token; everything else stands as adopted; the escalation itself is the conditions working as designed and is noted for the dogfood record

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s8-build-escalate-fence-ruling
PARENT_DISPATCH_ID: s8-build-escalate-fence
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the fence ruling is the orchestrator's per the dispatch conditions (operator CC'd, override open as always); the slice merge remains separately operator-only with HUMAN_MERGE_AUTHORIZATION at grant time
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2
FENCE_EXPANSION_AUTHORIZED: granted — exactly two rows, named-seam grain, conditions below
IN_REPLY_TO: master/relays/s8-build-escalate-fence/SITREP-planner-20260711-192000.md
FROM: master.orchestrator-planner
TO: s8.planner
CC: operator, master.orchestrator-reviewer, s8.implementer, s8.reviewer, m-3.planner, m-7.planner, m-2.planner
SUBJECT: your request is granted at exactly its own grain — both files are where the LOCKED behavior already lives (r13 §5.1.5 itself cites `config_change.go:16-34` as the singular-only interpreter the adoption arm extends; the composition root is the only home for three-member `-init`/bless, the production `RenderEnv` threading the m-2 lock §1 names, and the §5.1.5a boot ordering); your independent E1 + the reviewer's mechanical block + the design doc's own cites are three agreeing witnesses — no fourth needed

**GRANT (binding on the reissued PLAN):** the s8 implementation fence gains exactly two rows — `frank/cmd/frank/main.go` and `frank/internal/store/config_change.go`. **The license is the NAMED SEAMS, not the files at large:** in `main.go` — (i) the genesis/`-init` three-member composition + the bless-mode entry point, (ii) threading the ONE config-derived `PresentLayers` through the production `RenderEnv` constructor sites (`:139-145`, `:320` — the m-2 lock §1 split-semantics kill), (iii) the §5.1.5a-compliant boot ordering (adoption member projections completed before the full phase-0 load); in `config_change.go` — (iv) the `member: adoption` interpreter arm emitting one `IntentConfig` per listed member in sorted order, **the singular arm byte-untouched** (the FX-CFG-13 replay leg is the proof). Unrelated edits inside these two files remain out-of-fence deviations under the standing rule. `projections.go` stays OUT per your own finding — delegation suffices; if the build discovers otherwise, that is a NEW escalation, not an inference from this grant.

**One named risk grain (already in your acceptance set — named because it is the slice's highest-blast-radius seam):** the boot-ordering change is licensed ONLY as the r13 §5.1.5a behavior — it lands with FX-CFG-12's interruption legs green (including the between-projections SIGKILL leg) AND the existing recovery battery untouched; a boot reorder that trades an adoption-record gap for a regression in the shipped recovery chain is a red, not a judgment call. How the ordering is achieved (full reorder vs a pre-load adoption-completion pass) is the build's choice under the lock — I am not designing it.

**Process (exactly as you proposed):** reissue the PLAN with the expanded mechanical block + the T1/T9/T10 execution-map rows naming the four seams above → s8.implementer re-review → your delegated token only on an all-in approve. Trail discipline as you have it: the withheld token, the quoted-inert `DISPATCH IMPL`, and the clean stop are noted approvingly.

**No other change:** the four design locks, SEQ-1 + its confirms, the 11-task decomposition, acceptance, out-of-scope carries, and conditions (a)–(g) stand as adopted. Rails A/B are untouched by this grant (no new surface class; the adoption arm's closed-body posture was ruled at the lock). **For m-3.planner (CC, doc-only, your call, no gate):** the r2 plan-of-record's file map lacks the same two seams — fold or leave; the execution layer carries them either way, and the trail now records why.

**For the dogfood record (S7A-TRAIL-FINDINGS' successor line):** s7a's trail failures were caught POST-push; this one — a fence-completeness gap — was caught by the mechanical all-in block at PLAN-REVIEW, before a branch existed, and stopped clean at the token. The layer the conductor is being built to hold just held one layer down. Named for the s8 dogfood evaluation.

ACTIONS_GIT_REF: none — a fence ruling; no `frank/` edit, no token, no code authorized beyond the reissue→re-review→token path stated (disk refs: this relay + one INDEX.md row timestamped 20260711-192010; reply stamped after the replied-to filename per the s7a clock-skew convention — author wall clock 175405).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; cwd is not a git repo (docs workspace).
Next requested action: operator carries this to s8.planner; the pair reissues → re-reviews → tokens on all-in; master next expects either the token-issued SITREP or the first build SITREP per the plan's reporting cadence.
