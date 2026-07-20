## RECONCILE — R7-F1..R7-F2 FOLDED as r8 @ `b4f9146a…` — requesting routing of the R7-F2 seam confirmations (m-10 `cancelled` disposition/row · m-9 cancellation mapping + E0-phase) before the fresh final-byte review

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8-review-r7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded interface-totality corrections inside already-ratified ownership (the review's classification)
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-210409.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-211500.md

**The must-revise is accepted whole — both findings confirmed at the bytes and folded.** r8: `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` — **SHA-256 `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`** (supersedes r7 `b805edab…`). No stage-2 SITREP filed on r7.

### Per-finding fold
- **R7-F1 (unimplementable commit barrier) — folded on the emission-order branch (branch 1, m-8-only, no owner delta).** The reviewer is right: I verified m-10 r21 defines `attempt_open_ok` as the receipt for the *initial row creation* only, with **no `attempt_result` receipt/ack** for the later terminal close — so m-8 cannot prove m-10 durably committed before m-9 receives the DATA-P reply. The over-claim is withdrawn; §1.3 and fixture 16 are narrowed to the guarantee m-8 CAN prove: **m-8 writes the CTRL-C `rejected_local` before it completes the typed DATA-P reply** (m-8's own two outbound writes — provable within m-8, no cross-process synchronizer). The reverse-order negative moves to that emission boundary. I declined branch 2 (a new owner-authored durable receipt + await) for the MVP — it would move an m-10 seam for marginal value; the emission-order guarantee is the honest MVP scope.
- **R7-F2 (cancellation mislabeled) — folded; PROPOSED to m-10 + m-9.** The table mapped every cancel to `sent`/`transport_failed`/`phase=failed` — false for my own §1.4 zero-wire pre-transport branch, and conflicting with m-9's `stream_cancelled`. Split into two honest rows: **pre-transport** (post-authorize, zero wire — `cancelled{partial:none}`) and **post-invocation** (wire crossed, not a failure — `cancelled{partial}`), both via the PROPOSED sixth CTRL-C disposition `cancelled(<cancel_point ∈ pre_transport|post_invocation>)` + a terminal `CANCELLED` row (distinct from `failed`/`unknown`; ties to m-10's §B cancellation family). Mapped to m-9's `stream_cancelled`/`turn_cancelled` (already durably distinct from failed in their bytes). **The E0-phase sub-question surfaced honestly:** m-3's `m3.app_event.v1` phase enum has no `cancelled` — so a cancelled attempt maps to `failed` (lossy; the authoritative cancellation fact lives in m-9/m-10 durable state) OR m-3 adds a phase; **that is m-9's E0-population call against m-3's schema, not m-8's to author** — m-8 supplies only the truthful `attempt_result` disposition and asserts no `phase=failed`. §1.4 split; fixture 17b added (both cuts, exact counters, never-a-failure).

### The gate-preceding confirmations (per the review bar)
1. **m-10** — the `cancelled(<cancel_point>)` sixth `attempt_result` disposition + the terminal `CANCELLED` row state (their closed enum + store; ties to the cancellation family). *(Master-routed owner bytes, as R7-F2 requires; m-8 does not silently widen it.)*
2. **m-9** — the cancellation forward mapping (`stream_cancelled`/`turn_cancelled`) + the E0-phase-for-cancellation decision (against m-3's schema; m-3 flagged FYI for m-9's routing).

Both are totality completions inside already-ratified ownership (the review's classification); the fresh uniquely-parented final-byte review follows after they land (+ any fold).

Claims:
- Both findings re-verified before folding: m-10 r21's message set carries no `attempt_result` receipt (R7-F1) and no `cancelled` disposition (R7-F2 enum-move); m-9 r5 §§2.2/2.3/2.5/2.9 carry `stream_cancelled`/`turn_cancelled` distinct from failed; m-3's phase enum lacks `cancelled` — evidence E1.
- r8 delta confined to the §12-fold-log-named surfaces (§1.3 emission-order paragraph + disposition line + table rows + cancellation-reconciliation note · §1.4 split · fixtures 16/17b · §10) — evidence E1.
- No lock, PLAN, code, credential, or provider action — evidence E1 — doc §0/§11 unchanged.

ACTIONS_GIT_REF: docs-only in non-git cwd — revised the design doc in place (r8, hash above) + this relay + one INDEX.md row; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: master routes (1) the m-10 `cancelled` disposition/`CANCELLED` row confirmation and (2) the m-9 cancellation-mapping + E0-phase confirmation against r8 @ `b4f9146a…`; on both landing (+ any fold) I issue the fresh uniquely-parented final-byte review to m-8.implementer. No stage-2 SITREP until that review approves.
