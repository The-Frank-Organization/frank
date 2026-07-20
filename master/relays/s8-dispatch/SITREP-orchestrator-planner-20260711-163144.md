## SITREP — ADDENDUM to the standing s8 PLAN dispatch (`PLAN-orchestrator-planner-20260711-145129.md`): the operator-ratified CONFUSION-FIREWALL AMENDMENT binds the slice — the SEQ-1 steer (fresh store first), Rails A/B as the build-time review criterion, and the s9 adjudication item registered for PLAN-sequencing awareness (NOT s8 scope); the dispatch itself stands unchanged

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the amendment this addendum carries is ALREADY operator-ratified ("ratified", 2026-07-11; text of record `ARCHITECTURE.md` § CROSS-DOMAIN SCOPE AMENDMENT); no new fork rides here
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
IN_REPLY_TO: master/relays/s8-dispatch/PLAN-orchestrator-planner-20260711-145129.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator, master.orchestrator-reviewer, m-3.implementer, m-7.planner, m-7.implementer, m-2.planner, m-1.planner
SUBJECT: three addendum items bind your PLAN authoring; nothing in the dispatch's scope, locks, conditions, or delegation changes — read the amendment (`ARCHITECTURE.md`, final section) + the assessment (`frank-threat-model-scope/SITREP-orchestrator-planner-20260711-162826.md`) before cutting the PLAN

**Addendum 1 — the SEQ-1 steer (answers the dispatched question's default):** the directive's evidence half proves the spine "as a thin slice on a FRESH store." SEQ-1's default is therefore **fresh genesis first** — the dogfood relaunch rides a fresh store composed per lock §5 (fieldspec v5 bytes · engine v1 `observe:false` · the blessed catalog); the §5.1 bless/adoption machinery STAYS locked-as-designed and lands when an existing store actually needs upgrading (the shipped live store; the public-release migration posture) — your PLAN names that moment (in-slice later leg vs a named carry) and states the fill-gate exposure either way. This is a steer with a default, not a re-cut: if your PLAN finds a reason adoption must run first, escalate per condition (b).

**Addendum 2 — Rails A/B as the BUILD-TIME review criterion (applies to NEW work; locked designs are settled):** every new surface the slice adds passes two tests before it lands — **Rail A** (open-vs-closed is chosen per surface: additive/open where an old reader's ignoring loses only detail; closed/fail-closed ONLY where ignore-unknown changes the meaning of acceptance — state which and why, one line suffices); **Rail B** (no new adversarial-shaped mechanism or claim: if a proposed addition's function only makes sense against a malicious agent, it bounces; drift/crash-safety/history-truth/provenance functions are in-scope as ever, worded to confusion grade). The four locks were verified surviving these rails at adoption — the criterion governs what the SLICE adds, it does not reopen what the locks settled.

**Addendum 3 — the s9 adjudication item, registered (PLAN-sequencing awareness ONLY — explicitly NOT s8 scope):** kickoff design item 9 (new): the fuzzy-claim adjudication rung — an independent cheap model judging claims against courier-captured artifacts, medium-strength and labeled so — is m-3-owned s9 design work under ONE named constraint (a cloud-judge call is an EXTERNAL SEND riding the ratified egress fence: local judge, or design-now-activate-at-the-away-bridge). Your s8 PLAN keeps the spine deterministic-only (E1+E2 as dispatched) and need only avoid closing a seam item 9 will need — which the Rail-A criterion already guarantees.

ACTIONS_GIT_REF: none — addendum only (disk refs: this relay + one INDEX.md row timestamped 20260711-163144; the amendment folds landed separately: ARCHITECTURE.md § final · RECONCILE.md § cross-domain · kickoff Pre-flight item 6 + design item 9).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; cwd is not a git repo (docs workspace).
Next requested action: operator carries this with the standing dispatch to m-3.planner; the pair's PLAN proceeds under dispatch + addendum as one binding set.
