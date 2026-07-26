## RECONCILE — §D-settlement amendment rev2 `7137b18a…`: SETTLE-VP-R1-F1…F4 all folded; passed directions + the bound m-2 cell `5ec7a3d2…` preserved byte-exact; routed for VP exact-byte re-review

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — only fresh VP-approved amendment bytes may proceed to operator hash-bound ratification; master does not self-ratify
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260724-160000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: rev2 `7137b18a87a64fe67110ad83460d523f3ab7baa3525ac786b0e47713e12fea91` closes F1 (full cap-terminal lifecycle), F2 (both structural compile-time frame bounds + threshold interpretation), F3 (turn_failed reframed as explicit supersession, not already-descriptive), F4 (explicit owner fold matrix); Corrections 1/3 unchanged, m-2 cell byte-preserved

All four blockers were correct; every one is folded. Fresh ratification candidate: **`master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` @ SHA-256 `7137b18a87a64fe67110ad83460d523f3ab7baa3525ac786b0e47713e12fea91`**, bound m-2 cell **`5ec7a3d2…` preserved byte-exact** (no m-2 redispatch).

## How each finding is closed
- **F1 (cap-terminal lifecycle).** Correction 2.2 now selects the complete durable outcome to the `resume_frame_overflow` precedent's bar, **self-contained in the amendment** (not resting on m-10's non-bindable `48062d18…`): (1) the r40 §B.4 retirement commits in full **and the same transaction** commits `runs.state = FAILED` + `stop_reason = parked_unknown_capacity_exceeded`; (2) the closed D-4/resume reason set becomes exactly **`{resume_frame_overflow, parked_unknown_capacity_exceeded}`** with present-IFF rules, and **`resume_action = operator_new_run` stays EXCLUSIVE to `resume_frame_overflow`** — the cap terminal carries none; (3) **no** successor/continuation/lease/snapshot/revival; the operator surface renders the terminal + the complete queryable parked set; next work is an ordinary new run; the Step-4 direct-operator clearing path stays non-MVP; (4) an **acceptance predicate in the amendment** observes the full batch + terminal in one transaction, no successor, and the durable operator projection.
- **F2 (frame bounds + threshold).** Correction 2.4 states **both** D-4 carrier bounds as **compile-time assertions** with the exact terms: `turn_open`: `ADMISSION_REF_ENC_MAX + M_MAX + PARKED_MAX + PATH_MAX_M10 + OVERHEAD_MAX = 2,564,096 + 1,232,896 + 327,680 + 4,096 + 65,536 = 4,194,304 ≤ FRAME_MAX` (exact boundary); `attempt_open_ok`: `ATTEMPT_ACK_MEMBERS_MAX + PARKED_MAX = 1,024 + 327,680 = 328,704 ≤ FRAME_MAX`. I reproduced both at the bytes (turn_open sits exactly at 4 MiB — the intended tripwire on any future term increase). Correction 2.3 fixes the **threshold interpretation** — max **nonterminal** parked count, not a storage cap: `511+1=512` continues, `512+1=513` commits full row + terminal atomically, multi-row retirement may overshoot with **every row retained/queryable**; the boundary/multi-row/no-prefix cases ride the acceptance predicate so it cannot be built as a 512-row truncation.
- **F3 (turn_failed).** Correction 4 is reframed exactly as you required: r21 `:115`'s "after" **is** scope-limiting, `turn_denied`'s pre-attempt branch proves only the family **can** operate pre-attempt — so this is an **explicit semantic supersession effective only on operator ratification**, of that one clause **for the single named pre-attempt assembly-refusal branch** (ADMITTED→ASSEMBLING validation failure; zero digest/attempt/DATA-P/provider-row; exactly one `turn_terminal{…, terminal: turn_failed}` + ordinary `turn_receipt{terminal_recorded}`; no auto-retry / second assembly). It **no longer claims** frozen r21 already meant this; r21 stays byte-exact historical source; any other zero-attempt use is a fresh amendment.
- **F4 (propagation).** The shorthand is replaced by an explicit **owner fold matrix**: **m-9** folds Correction 1's consumer label + Correction 2's full run-wide consumer/comparator restore + terminal consumer posture + Correction 3's bound cell (§7-`relay.*`) + Correction 4's exact branch (one batched F73 revision); **m-10** folds Correction 1's producer label + Correction 2's run-wide producer + the terminal + schema/presence + the **two compile-time frame assertions** + fixtures, over a **bindable successor to `48062d18…`** (its current bytes are non-bindable, so its written relabel is explicitly **not** a durable fold); **m-2** folds nothing (cell byte-final). Fresh pair review each; consumer confirmations + the §D join only afterward.

## Unchanged (VP-passed)
Correction 1 (Gate-2 relabel, additive claim amendment, comparator bytes historical), Correction 3 (the m-2 cell, form-derived total, effect-vs-authority separation, `null` rejected), the two-file packet, master-does-not-self-ratify, exact hash binding, the H-12 boundary, and all downstream holds — carried verbatim.

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. rev12 `1125b0a0…`, r40 `d2ce9831…`, r21 `4d3bd14e…`, the m-2 cell `5ec7a3d2…`, m-9 r12 `04422965…`, m-3 r15 `d004dbc7…` all UNMOVED. m-10 `48062d18…` non-bindable/must-revised. H-12 external-use block stands.

## Verification
Hashed on disk this session: rev2 amendment `7137b18a87a64fe67110ad83460d523f3ab7baa3525ac786b0e47713e12fea91`; bound m-2 cell `5ec7a3d254da…` (byte-identical to r1); VP r1 review target `342ec33d…` consumed. Frame arithmetic reproduced independently (turn_open = 4,194,304 = FRAME_MAX exact; attempt_open_ok = 328,704). rev12/r40/r21 frozen bytes re-confirmed unmoved. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — the rev2 amendment `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` + this relay + one INDEX.md row; no frozen/ratified byte moved, no `frank/` action, no lock issued, no self-ratification, no fold performed, no m-2 redispatch.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP exact-byte re-review of rev2 `7137b18a…` + the byte-bound cell `5ec7a3d2…`; on approve → operator hash-bound ratification → the propagation matrix executes; on any residual → master folds and re-routes uniquely parented.
