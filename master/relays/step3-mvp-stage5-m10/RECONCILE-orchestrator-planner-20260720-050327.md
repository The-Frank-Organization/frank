## RECONCILE — the CORRECTED H-17 supplement to m-10 (VP F87; supersedes my `043341` as the operative census requirement): the canonical row schema = `master/H17-CENSUS-SCHEMA.md` v1 @ SHA-256 `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5` — use it verbatim; PLUS the VP's two concrete corrections to your current draft (@ `d47639…`): (1) COVERAGE — your seven rows omit authoritative transition families: run start/stop/recovery · process spawn/retire · turn admission · epoch publication · cancellation/control sends · provider-attempt transitions · app-event carriage; a generic STORE meta-row cannot replace per-effect authority/failure rows · (2) ROLES — recording or validating a worker/connector REPORT does not make you an independent OBSERVER of the underlying effect; name reporter/observer/validator/recorder separately per row (`none (self-reported)` is the honest cell where it's true); drafting + the grill CONTINUE, exits gated as with m-9

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a corrected additive statement requirement; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the stage-5 grill stands and continues
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-045522.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: same rule set as the parallel m-9 correction — every authoritative transition/effect family maps to a row or an explicit non-effect rationale; missing cells say `unknown`/`not specified`/`residual`, never invented, never a moved closed byte; no GRILL_LOCK / final-byte review / final SITREP until the complete inventory is folded; the grill (G-4+) and drafting continue uninterrupted

ACTIONS_GIT_REF: docs-workspace disk action — this supplement relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: fold + continue; the inventory arrives with your design SITREP when the exits unlock.
