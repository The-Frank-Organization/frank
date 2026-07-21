## RECONCILE — the CORRECTED H-17 supplement to m-9 (VP F87; supersedes my `043331` as the operative census requirement): the canonical row schema is now ONE byte-exact artifact — `master/H17-CENSUS-SCHEMA.md` v1 @ SHA-256 `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5` — use it verbatim; PLUS the VP's four concrete corrections to your current draft rows; drafting + the grill CONTINUE, but NO GRILL_LOCK, final-byte review, or SITREP until the complete row/rationale inventory is folded

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a corrected additive statement requirement; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the stage-4 grill stands and continues
DESIGN_DOC_ID: step3-mvp-design-m9-worker
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-045522.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, master.orchestrator-reviewer, operator
SUBJECT: the VP reviewed your current draft (@ `fdc6aaf0…`) and found the rows additive-safe but not assemblable — the missing schema fields (`effect_id` merge key · `effect_class` · `policy_owner`/`policy_artifact` · `decision_point` · the reporter/observer/validator/recorder split · `threat_claim_scope`) are now in the canonical schema; and four content corrections at your rows: (1) `attempt_open_ok` is ADMISSION/row-first ordering, NOT provider authorization — m-8 owns policy authorization immediately before credential attach/send; relabel · (2) `relay.project`/`read` do not append — they are their own read/serve rows, and `Describe` + push are missing from your closed client surface; the compaction provider call must route explicitly as a fresh ordinary m-8 attempt row · (3) E0 non-emission is a FAILURE/UNKNOWN condition, not a bypass — a bypass is an alternate way to CAUSE an effect · (4) coverage: every authoritative transition/effect family gets a row or an explicit non-effect rationale — the dispatch parenthetical was a minimum, not the inventory

m-9 — fold the schema + the four corrections into the working draft; the grill with the operator continues uninterrupted; the gate is only at the exits (no GRILL_LOCK / final-byte review / SITREP before the complete inventory). Where a frozen contract lacks a cell's answer, write `unknown`/`not specified`/`residual` — never invent, never move a closed byte.

ACTIONS_GIT_REF: docs-workspace disk action — this supplement relay + one INDEX.md row (the schema file was created this session, recorded in the VP-return cover); no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: fold + continue; the inventory arrives with your design SITREP when the exits unlock.
