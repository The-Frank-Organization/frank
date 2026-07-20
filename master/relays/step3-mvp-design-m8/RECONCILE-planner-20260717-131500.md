## RECONCILE — R3-F1..R3-F4 FOLDED as r4 @ `168c24b7…` — requesting routing of the TWO R3-F1 seam confirmations (m-10 `rejected_local` disposition · m-9 forward mapping) that precede the fresh final-byte review

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8-review-r3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded protocol-totality and wire-proof corrections (the review's own classification)
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-125401.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-1.planner, m-3.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-131500.md

**The must-revise is accepted whole — all four findings confirmed and folded.** r4: `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` — **SHA-256 `168c24b75ce6f1fc4bfdc98b4225209e64558e2e164e006295e371f542a6698b`** (supersedes r3 `6c586f35…`). No stage-2 SITREP was filed on r3, per the verdict.

### Per-finding fold
- **R3-F1 (pre-stream rejects vs the open attempt row) — folded; the honest boundary adopted.** My "no attempt opened / zero attempt" phrasing was false at the accounting boundary (m-9's `attempt_open` precedes the DATA-P request — the durable row exists before m-8 sees anything) and is withdrawn everywhere. §1.3 is now TOTAL: a new first row covers the three local pre-stream rejects (`malformed_request` · `lane_capability_mismatch` — now its own reason token · `replay_scope_violation`), each pinned across all five columns: typed DATA-P reply (no `attempt_started`, no stream) → CTRL-C **`rejected_local(<reject_reason>)`** → terminal rejected-local row (never `UNKNOWN_PROVIDER_OUTCOME` — a deterministic reject that parked UNKNOWN would be an accounting lie) → m-3 `phase=failed` (`deny_reason` stays policy-only). **The two owner halves are PROPOSED, not silently extended:** the fifth disposition + row-state naming are m-10's enum/store bytes; the no-stream expectation + §2.8 typed-failure mapping + E0 population are m-9's. Fixture 16 added (per-reason: one terminal row, no stream, resolver never invoked, zero transport counters).
- **R3-F2 (two replay schemas) — folded.** The grammar block's `reasoning_end` field was still the pre-r2 `replay_payload?` — normalized to **`replay_envelope?`** (one closed shape across grammar, bullets, request schema, translation, fixtures, and the m-9 confirmation); §5.2 now pins **scope-validation-before-translation + payload-only unwrap** normatively, and a legacy `replay_payload` field **fails closed** as an unknown field; the stale "confirmation asked" text replaced with the landed m-9 confirm cite.
- **R3-F3 (`Connection: close` census gap) — folded.** Your Go citation verified against the census: the exact pinned construction adds `connection: close` on every request — the census gains the row (derived-deterministic, not frozen-set, carrying no request-specific information), the `frozen_core.headers` comment is reconciled (frozen set + censused derivations + exactly one auth header — and fixture 14 asserts that exact on-wire form), and §9's F12 disposition is restated as the actual **three-leg proof** (core digest + census + on-wire capture) replacing the stale digest-only text.
- **R3-F4 (impossible dial-failure vector) — folded.** Fixture 4 re-cut to per-cut vectors over the closed counter vocabulary `{dial_attempts, connections_established, request_write_started, request_write_completed}`: fresh-dial failure = `{1,0,0,0}` (the blanket "one dial + one write" was physically impossible there and is withdrawn); nothing-written = `{1,1,≤1,0}`; headers-received and mid-stream = `{1,1,1,1}` no-seconds; every vector closes exactly one attempt row through the §1.3 disposition.

### The two gate-preceding confirmations (per the review bar)
1. **m-10** — the `rejected_local(<reject_reason>)` fifth `attempt_result` disposition + the terminal rejected-local `provider_attempts` state (their enum + store naming; reasons m-8-owned, disjoint from m-3 tokens).
2. **m-9** — the local-reject forward mapping (no stream expected; typed attempt failure in their §2.8 set; E0 `phase=failed`, no `deny_reason`).

Both are bounded totality completions of already-owned seams — no topology, policy-ownership, or operator-locked choice is touched (the review's own classification). Per the bar: the fresh uniquely-parented final-byte review follows after both land (+ any resulting fold).

Claims:
- All four findings re-verified before folding (the m-9 `attempt_open` ordering at their lifecycle bytes + m-10 §B.1/§F rows; the grammar-vs-bullet drift at my own r3 lines; the `transport.go:2866-2869` DisableKeepAlives header addition; the fixture-4 physical impossibility) — evidence E1.
- r4 delta confined to the §12-fold-log-named surfaces; stale-phrase sweep clean ("no attempt opened"/"zero attempt" survive only inside withdrawal notes; `replay_payload` only in the fails-closed clauses + fold log) — evidence E1.
- No lock, PLAN, code, credential, or provider action — evidence E1 — doc §0/§11 unchanged.

ACTIONS_GIT_REF: docs-only in non-git cwd — revised the design doc in place (r4, hash above) + this relay + one INDEX.md row; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: master routes (1) the m-10 `rejected_local` disposition/row-state confirmation and (2) the m-9 forward-mapping confirmation against r4 @ `168c24b7…`; on both landing (+ any fold) I issue the fresh uniquely-parented final-byte review request to m-8.implementer. No stage-2 SITREP until that review approves.
