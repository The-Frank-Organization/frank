## RECONCILE -- REVISE NARROW: Gate-2 relabel and the bound m-2 cell pass, but the cap terminal/frame proof and turn_failed supersession are not ratification-complete

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-vp-review-r1
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- only fresh VP-approved amendment bytes may proceed to operator hash-bound ratification
GRILL_REQUIRED: no -- the selected directions are settled; this return closes exact lifecycle, frame-totality, supersession, and propagation semantics
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260724-151500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: REVISE exact amendment 1f822e47 -- preserve the Gate-2 claim relabel and m-2 cell 5ec7a3d2; make the cap terminal and both frame bounds exact, state turn_failed as an additive semantic supersession, and route every correction to every affected owner

VERDICT: revise

Review target: `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260724-151500.md` at SHA-256 `342ec33decea7969a1b9bdaa8f22165d07eec487e6653e213fcb836d25a53938`.

Exact ratification candidate reviewed:
- `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` at SHA-256 `1f822e4711d5772ffafc68c4183ddb0faa33250e5b9d9372ead0e4128c34dbe7`;
- bound m-2 cell `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` at pair-approved SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.

## Findings

### SETTLE-VP-R1-F1 -- BLOCKER: the new cap terminal does not select its complete durable lifecycle outcome

Amendment `:32-37` says the full retirement commits and the same transaction commits a typed run terminal with `stop_reason = parked_unknown_capacity_exceeded`. It does not name the run state, the closed `stop_reason` domain after this addition, presence rules, successor/revival posture, or operator next action. Those are lifecycle semantics, not storage encoding details.

The prior Stage-6 overflow amendment was not ratification-ready until it selected exactly one outcome: `runs.state = FAILED`, a closed reason, no successor/lease/snapshot/same-run revival, and an operator-visible manual next action. The same bar applies here. The m-10 working delta contains the missing detail, but it is not bound into this packet and remains under a live must-revise verdict at `48062d18...`; operator ratification cannot depend on non-bindable working bytes.

Required amendment fold:

1. The ordinary r40 B.4 retirement transaction commits in full, including every parked row and all normal fence/retirement effects; that same transaction commits `runs.state = FAILED` and `stop_reason = parked_unknown_capacity_exceeded`.
2. The closed amendment-visible reason set is `{resume_frame_overflow, parked_unknown_capacity_exceeded}`, with exact present-IFF rules. `resume_action = operator_new_run` remains exclusive to `resume_frame_overflow`; the cap terminal carries no `resume_action`.
3. The cap terminal admits no successor generation, same-run continuation, or revival. The operator surface renders the terminal and the complete queryable parked set; any next work is an ordinary new run. The deferred Step-4 direct-operator clearing path remains non-MVP.
4. Add an acceptance predicate that observes the full retirement batch and terminal in one transaction, no successor, and the durable operator projection. Do not leave terminality or revival to the later F73 fold.

### SETTLE-VP-R1-F2 -- BLOCKER: the claimed static frame bound omits one carrier and does not define the terminal overshoot boundary

Amendment `:35` gives only `512 * 640 = 327,680` and `ADMISSION_REF_ENC_MAX = 2,564,096`, then claims every term is statically bounded. Target `:27` further summarizes the proof as `512*640 < ADMISSION_REF_ENC_MAX`, which compares unrelated budgets and does not prove either complete frame.

Both D-4 carriers must be structural:

```text
PARKED_MAX = MAX_PARKED_ROWS_PER_RUN * PARKED_ROW_MAX = 512 * 640 = 327,680

ADMISSION_REF_ENC_MAX + M_MAX + PARKED_MAX + PATH_MAX_M10 + OVERHEAD_MAX
  <= FRAME_MAX

ATTEMPT_ACK_MEMBERS_MAX + PARKED_MAX
  <= FRAME_MAX
```

At the accepted constants those are `2,564,096 + 1,232,896 + 327,680 + 4,096 + 65,536 = 4,194,304` for `turn_open`, and `1,024 + 327,680 = 328,704 <= 4,194,304` for `attempt_open_ok`. The fresh amendment must require both compile-time assertions; prose arithmetic is not the gate.

The threshold also needs an exact interpretation. It is the maximum parked count on a run that may remain nonterminal and emit another D-4 frame, not a hard storage truncation limit: `511 + 1 = 512` continues; `512 + 1 = 513` commits the full row and terminal atomically; a multi-row retirement may legitimately overshoot farther, with every row retained/queryable. No later D-4 frame is emitted because the run is terminal. Put those boundary and multi-row/no-prefix cases in the amendment acceptance predicate so `MAX_PARKED_ROWS_PER_RUN` cannot be implemented as a 512-row storage cap.

### SETTLE-VP-R1-F3 -- BLOCKER: correction 4 is a semantic extension, not a reading already supplied by frozen r21

Frozen r21 `:115` says `turn_failed` is a machinery failure **after** bounded attempts. That is scope-limiting text. The fact that the distinct `turn_denied` enum member has one pre-attempt branch proves only that the terminal frame family can operate before an attempt; it does not make `turn_failed`'s own `after` clause descriptive.

Adjudication: an additive amendment is the correct instrument, and r21 should remain byte-exact as historical source. No in-place edit is required. But amendment `:61-68` must call this what it is: an explicit semantic supersession, effective only upon operator ratification, of r21's after-attempt restriction for the one named assembly-refusal branch. It must not claim that frozen r21 already meant this.

Keep the extension bounded to the pair-reviewed branch: an ADMITTED turn fails totality/shape/alias validation in ASSEMBLING; zero `logical_surface_digest`, `attempt_open`, DATA-P request, and provider-attempt row; exactly one existing `turn_terminal{..., terminal: turn_failed}` followed by the ordinary `turn_receipt{terminal_recorded}`; no auto-retry or second assembly. Any other new zero-attempt use remains a fresh lifecycle amendment.

### SETTLE-VP-R1-F4 -- BLOCKER: the specific downstream fold list drops affected owner halves

Amendment `:6` says owners fold all corrections, but the more specific `:83` routes m-9 only the Gate-2 label plus `relay.*`, and m-10 only Correction 2 plus FX-M10-D4. That omits:

- m-9's Correction 2 run-wide/comparator restoration and Correction 4 terminal extension;
- m-10's Correction 1 claim relabel;
- the fact that current m-10 `48062d18...` is non-bindable, so its already-written relabel cannot be treated as a durable fold.

Replace the shorthand with an explicit owner matrix. m-9 folds Correction 1's label, Correction 2's full run-wide consumer/comparator semantics, Correction 3's bound `relay.submit` cell, and Correction 4's exact branch. m-10 folds Correction 1's producer-side label and Correction 2's run-wide producer, terminal, schema/presence rules, two frame assertions, and fixtures. Each changed owner artifact receives fresh pair review; affected consumer confirmations and the two-sided D join occur only afterward. The m-2 cell stays byte-bound and needs no new pair cycle unless its bytes move.

## Passed portions

- **Correction 1 direction and instrument pass.** The two-gate guarantee claim is false on reachable MVP bytes; relabeling Gate 1 as the guarantee and Gate 2 as reachable validation plus an unreachable-state drift detector is honest. Because the amendment explicitly changes the operator-visible claim and requires operator ratification, an additive claim amendment is sufficient; frozen comparator bytes may remain historical and unchanged.
- **Correction 2 direction passes.** Restore full run-wide, worker-independent carriage on both frames; reject the turn-scoped/consumer-derived fallback; never truncate; fail loudly at a bounded live-run threshold.
- **Correction 3 passes.** The m-2 cell is owner-authored, pair-approved, total through required `form_digest`, deterministic over the named optional destination coordinates, and correctly keeps effect-target binding separate from TO/CC authority. `canonical_resource = null` would discard available invocation context and is rejected.
- The two-file packet construction, master-does-not-self-ratify rule, exact hash binding, H-12 boundary, and all downstream holds are correct.

## Gate disposition

- Do not route amendment `1f822e47...` to the operator.
- Preserve the bound m-2 cell exactly at `5ec7a3d2...` and its pair approval.
- Master returns fresh amendment bytes and a uniquely parented review request closing F1-F4. No m-2 redispatch is needed unless the bound cell changes.
- Operator ratification, owner folds, the D join, integrated re-lock, DESIGN-lock, PLAN, T4/code, credentials, provider calls, release binding, live E3, merge, deploy, and external use remain held.

## Verification

- Recomputed exact hashes: target `342ec33d...`; amendment `1f822e47...`; bound m-2 cell `5ec7a3d2...`; master ruling `922b796d...`; rev12 `1125b0a0...`; frozen r40 `d2ce9831...`; frozen r21 `4d3bd14e...`; pair-approved m-9 r12 `04422965...`; non-bindable m-10 working delta `48062d18...`.
- Reproduced the frame arithmetic above and checked the 511/512/513 plus multi-row overshoot semantics against the current m-10 owner bytes and its adversarial review trail.
- Read the current r40 B.2/B.4/D.4 contracts, r21 terminal table and D-4 consumer, m-9 r12 held folds, m-10 cap fixtures, the prior Stage-6 terminal precedent, and the live INDEX through the target. The target exact-file lint is OK.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain; HEAD and `origin/main` both equal `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, bound cell, owner design, historical relay, `frank/` source, branch, commit, lock, ratification, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260724-160000.md`; the command's nonzero root result is the disclosed historical INDEX/lineage noise, not an error in this relay.
Next requested action: master folds F1-F4 into a fresh settlement-amendment hash, preserving m-2 `5ec7a3d2...`, and returns the exact bytes for VP re-review; only a clean successor may proceed to operator ratification.
