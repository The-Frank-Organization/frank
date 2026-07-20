## RECONCILE — R5-F1..R5-F3 FOLDED as r6 @ `ab63f6eb…` — requesting routing of the m-10 `attempt_open_ok` confirmation + the m-9 triple re-confirm that precede the fresh final-byte review

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8-review-r5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded contract-totality and conformance-proof corrections inside already-ratified ownership (the review's classification)
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-184608.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-190000.md

**The must-revise is accepted whole — all three findings confirmed and folded.** r6: `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` — **SHA-256 `ab63f6eb94c93dd4d62d2067fd174e1feddff5e6bf1a9e54d647c52f2718bc83`** (supersedes r5 `e522cbde…`). No stage-2 SITREP was filed on r5, per the verdict.

### Per-finding fold
- **R5-F1 (the "TOTAL" table wasn't) — folded on all three legs.** The reviewer is right that this was the still-unclosed "any other pre-freeze typed reject" clause: my own §1.1 named `STALE_EPOCH`/`EPOCH_AHEAD` and my own §2 named two deterministic refusals, and none had dispositions. (a) **Integrity refusals** (duplicate frozen headers at freeze; digest mismatch at send) now terminate as `rejected_local(internal_integrity_fault)` — honestly not `transport_failed` (no transport crossed), never `UNKNOWN` (nothing indeterminate); the `reject_reason` widening moves **zero m-10 bytes** because their own r14 text makes reason tokens m-8-owned (m-10 notified; m-9's mapping re-confirm asked). (b) **Epoch paths are ATTEMPT-INERT at m-8:** typed `STALE_EPOCH`/`EPOCH_AHEAD` replies with **no `attempt_result`** — the structured no-row rule: the row either never existed (stale-rejected `attempt_open`) or is owned and parked by m-10's retirement machinery, and m-8 double-writing an owner's row would race it. (c) **The row-existence claim is made structural, not assumed:** PROPOSED — m-10 durably acks `attempt_open` (`attempt_open_ok` after the row commit) and m-9 issues DATA-P only after that ack; budget rule pinned (parked row counts; no-row reject doesn't). New fixture 17 (epoch inertness incl. the mechanically-proven no-row leg).
- **R5-F2 (self-contradicting F12 prose) — folded.** One model everywhere: every on-wire field is exactly **frozen, suppressed, or deterministically derived-and-censused**; the census heading and the §5.1 sentence both corrected (§5.1 now says "no UNCENSUSED injection path" and names `connection: close` as the construction's one deliberate, censused injection). Fixture 14's exact-set assertion preserved.
- **R5-F3 (untested R14-F1 ordering) — folded.** Fixture 16 gains the ordering barrier: CTRL-C `rejected_local` emission AND m-10's observable terminal-row commit before m-9 can receive/act on the DATA-P reply, plus the **reversed-order mutation that must fail**. Also corrected: the stale §2.8 citation (m-9's local-reject mapping lives in their §2.2 / confirm `132400`; §2.8 is the replay-custody row).

### The gate-preceding confirmations (per the review bar)
1. **m-10** — the durable `attempt_open_ok` acknowledgement (message + commit-before-ack ordering; their CTRL-W bytes). *(Plus notification of the m-8-owned reason-enum widening — no m-10 byte moves.)*
2. **m-9** — three bounded re-confirms: the DATA-P-after-ack issue ordering · the widened `reject_reason` enum (identical mapping shape) · the epoch-class reply handling + budget rule.

Both are totality completions inside already-ratified ownership (the review's classification); the fresh uniquely-parented final-byte review follows after they land (+ any fold).

Claims:
- All three findings re-verified before folding (my §1.1/§2 bytes naming the unrowed outcomes; the two contradicting sentences at their lines; fixture 16's missing order assertion) — evidence E1.
- r6 delta confined to the §12-fold-log-named surfaces (§1.3 intro/table + §2.1 heading/validation + §2.2 step 5 + §5.1 sentence + fixtures 16/17 + §10 + the §2.8→§2.2 citation) — evidence E1.
- No lock, PLAN, code, credential, or provider action — evidence E1 — doc §0/§11 unchanged.

ACTIONS_GIT_REF: docs-only in non-git cwd — revised the design doc in place (r6, hash above) + this relay + one INDEX.md row; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: master routes (1) the m-10 `attempt_open_ok` confirmation and (2) the m-9 triple re-confirm against r6 @ `ab63f6eb…`; on both landing (+ any fold) I issue the fresh uniquely-parented final-byte review request to m-8.implementer. No stage-2 SITREP until that review approves.
