## DESIGN — D-3 disposition to m-7 (from the m-9 half r4 @ `1cb4ab57…`): the broker ATTACH-RESULT TAXONOMY — a NARROW, broker-only reopen of your attach-result interface per your own owner path: distinguish `suspended-no-state` (transient; the F70 floor; retry-under-supervision) from `tuple-mismatch` (this generation is fenced; terminal) as TYPED results; an ACCEPT folds as r9 + fresh review

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m7
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded, broker-only interface completion inside your ratified ownership; no topology/claim-boundary change (the m-9 review's R3-F3 correction names this honestly as a REOPEN of your attach-result interface, routed to you rather than proxy-authored)
GRILL_REQUIRED: no — the placement GRILL_LOCK is untouched; this is a result-taxonomy completion, not a placement/custody choice
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
IN_REPLY_TO: master/relays/step3-mvp-lifecycle-m9/SITREP-planner-20260717-185400.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, m-9.planner, m-9.implementer, m-10.planner, m-1.planner, master.orchestrator-reviewer, operator
SUBJECT: the D-3 disposition — m-9's half consumed your F70 attach-refusal and found ONE refusal shape carrying TWO different meanings: suspended-no-installed-state (transient hold; resolves when a state installs) vs presented-tuple-mismatch (this generation is fenced; permanent for it) — their worker must not retry a fenced generation nor exit on a transient floor; the split is YOUR bytes (typed attach results), their consumption is pinned r4-side

m-7 — m-9's lifecycle half r4 (`1cb4ab570813d3b2423576ea…`) consumed your F70 two-arm branch faithfully and surfaced a real taxonomy gap AT your attach surface (their R3-F3 leg honestly re-classified it as a narrow reopen of YOUR interface — the owner path, not proxy authorship):

**The gap:** your r8 attach refusal is one typed refusal, but its two causes demand OPPOSITE worker behavior — **`suspended-no-state`** (the F70 withheld-arm floor: transient; attach becomes possible when an installed state exists; the worker holds under supervision visibility with bounded retry) vs **`tuple-mismatch`** (the presented `{run_id, generation_id, turn_epoch}` fails verification: this generation is FENCED; retrying is a stale-generation hammering the fence; the honest disposition is terminal for the generation). m-9's r4 pins their consuming disposition split (typed transient hold ≠ fencing) and marks your half AWAITING.

**Dispose accept / refine / reject.** An ACCEPT = the **typed attach-result taxonomy in your bytes** (r9 on your approved r8 `ab0ed428…`): the two named results (plus any third your matrix requires — e.g. the recognition-mismatch row from R7-F1, if it surfaces at attach), each with its meaning + the worker-behavior contract it licenses; the F64 fence and the §2.10 branch semantics unchanged (this TYPES the existing refusal, it does not weaken it); the H-14 reachability lens on both new result values. Fresh uniquely-parented m-7.implementer review → SITREP. The m-9 half rebases on your reviewed bytes (their deferred re-review runs once, over stable owner bytes); m-1's row-6 confirm gets a scoped rebind if your §2.7/custody surfaces move (name it in the SITREP if so).

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-7.planner disposes; on ACCEPT, r9 + fresh review + SITREP naming whether the custody surfaces moved; master routes the rebinds with the r15 round.
