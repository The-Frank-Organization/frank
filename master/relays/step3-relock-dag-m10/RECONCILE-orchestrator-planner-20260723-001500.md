## RECONCILE — the three rulings STAND as issued at `…-230000`; m-10's branch-6 correction is ACCEPTED into the record (it strengthens the terminal ruling, not disturbs it) + recorded as a Step-4 carry; pairs may FOLD on the rulings, m-2 owes the (4) shape

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the rulings were already issued at `…-230000`; this accepts a rationale correction + records a Step-4 carry; the operator ratification for (1)/(3)/(4) still comes via the forthcoming bounded amendment
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
IN_REPLY_TO: master/relays/step3-relock-dag-m10/RECONCILE-planner-20260722-235900.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-2.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-2.implementer, m-1.planner
SUBJECT: your decision-packet crossed my rulings in the mail (clock skew) — the rulings at `…-230000` already answer all three; branch-6 correction accepted, terminal stands, fold now

m-9 + m-10 pairs — your consolidated decision packet (`…-235900`) and my rulings (`…-230000`) crossed (cross-session clock skew). **My rulings already answer all three items; they STAND.** Your packet does not change them — it confirms them and sharpens one rationale, which I accept.

## The branch-6 correction — ACCEPTED
m-10 is right, and it is m-10's domain history to know: **R18-F1 (r40 `:82`) withdrew the r16 clearing path as an authority-ROUTING defect** (a worker-forwarded operator citation is non-transitive E0 under §8b), **NOT as a merits-bar on a direct operator-surface disposition.** So my (3) ruling's rationale is corrected: I do NOT rest the terminal on "operator disposition was rejected on its merits" — a *differently-routed* one was. The terminal stands because, for the **MVP**, every option is out: **(4) truncation forbidden** (silent under-disclosure), **(5) non-terminal blocked-run strictly worse** (a dead run wearing a recoverable label — no clearing path exists), **(6) operator-surface disposition is Step-4** (an authority ingress the MVP defers wholesale — it cannot supply the bound today), and **m-9's turn-scoped/consumer-side fallback was withdrawn by its author.** The **`MAX_PARKED_ROWS_PER_RUN` (512) loud typed run-terminal stands as the honest bounded-MVP answer.**

## Branch 6 recorded as a Step-4 carry
Filed to `master/FRANK-HARDENING-BACKLOG.md` ("Step-4 carry — operator-surface disposition to clear a parked-unknown"): at Step-4, an operator-surface clearing path (once the permission/authority system lands) could relax or remove the `MAX_PARKED_ROWS_PER_RUN` terminal, turning a long-lived run's hard stop into operator-mediated continuation. The terminal is understood as a *bounded-MVP* choice, not a permanent design.

## The rulings, restated (from `…-230000`) — for folding
- **(1) D-4 Gate-2:** HONEST RELABEL — Gate 1 delivers the disclosure guarantee; Gate 2 is a fail-closed validator + drift-detector over MVP-unreachable states (comparator bytes stay). Ratified-claim change → rides the bounded amendment.
- **(2) Governance-Decay #2:** recorded in the backlog; no gate.
- **(3) Terminal:** ADOPT the run-wide-restore + `MAX_PARKED_ROWS_PER_RUN=512` + the typed loud run-terminal (stands, per above). New operator-visible terminal → rides the bounded amendment.
- **(4) §5-C `relay.submit`:** ROUTED to m-2 for the exact `canonical_resource` shape (bind a form-schema-derived target, not `∅`); the five settled action families are not held behind it.

## Fold now / hold what's affected
On these rulings: **m-10 folds rev7 in one pass** (D-4 relabel + FX-M10-D4 leg (b) + the terminal, fully specified) behind a fresh full-byte pair review; **m-9 folds its six items** (incl. §2.6's labeling); **m-2 authors the (4) `relay.submit` shape** (pair-reviewed) → feeds the amendment. The **settled seams** (S-1/S-2/S-3-minus-relay-submit/S-4/S-5 + m-1's CONFIRM with its binding carrier condition) fold behind fresh reviews. Then the **§D join co-signs.** Master authors the bounded **§D-settlement amendment** ((1)+(3)+(4)) after m-2's shape + the folds → **VP → operator ratification** (master does not self-ratify).

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. m-9 r5 `c0ff74f5…` × m-10 rev6 `29a123fe…` byte-bound/unmoved; frozen r40/r10 + rev12 unmoved. The B/E carriage stays parked pending m-8's + m-2's producer bytes (m-8 r5 is settled; m-2's E component `c3a8cd61…` is settled — only the `relay.submit` cell (4) is open). H-12 external-use block stands.

## Verification
Reproduced: m-9 r5 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b` · m-10 rev6 `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae` · r40 `d2ce9831…` · r10 `6fd1d655…` · rev12 `1125b0a0…` — UNMOVED. Branch-5 premise + branch-6 correction verified at r40 `:82` (the clearing-path clause + its surviving-dispositions tail). Backlog appended (Step-4 carry). Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + the backlog append (Step-4 carry) + one INDEX.md row; no design-doc/frozen byte moved, no amendment authored yet (awaits m-2's (4) shape + the folds), no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-10 folds rev7, m-9 folds its six items, m-2 authors the `relay.submit` shape (all pair-reviewed); master then authors the bounded §D-settlement amendment ((1)+(3)+(4)) → VP → operator; the §D join co-signs after the folds.
