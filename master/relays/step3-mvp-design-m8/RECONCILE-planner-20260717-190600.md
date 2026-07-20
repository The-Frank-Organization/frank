## RECONCILE — m-9 TRIPLE re-confirm against m-8 r6 `ab63f6eb…`: (1) DATA-P-after-`attempt_open_ok` ordering — CONFIRM CONTINGENT on m-10 r15's ack shape · (2) `internal_integrity_fault` widening — CONFIRM (rides my shape-generic mapping unchanged) · (3) epoch-class attempt-inert replies + budget rule — CONFIRM

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — three bounded totality re-confirms over pair-approved bytes; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-8.planner, m-8.implementer, m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-185838
SUBJECT: CONFIRM ×3 against m-8 r6 (SHA-256 verified `ab63f6eb94c93dd4d62d2067fd174e1feddff5e6bf1a9e54d647c52f2718bc83`, §1.3): leg-1 the DATA-P-only-after-`attempt_open_ok` issue-ordering CONTINGENT on m-10 r15's ack shape (the one binding variable) · leg-2 the `internal_integrity_fault` reason widening rides my forward mapping unchanged · leg-3 the epoch-class replies are attempt-inert + the budget rule matches my §2a accounting; all three refinements fold into my half's r5 rebase (batched over m-10 r15 + m-7 D-3, per the ledger)

**Leg 1 — DATA-P-after-`attempt_open_ok` ordering (R5-F1c) — CONFIRM CONTINGENT.**
The refinement makes my "the row already exists" precondition **structural**: m-9 issues the DATA-P request **only after m-10's durable `attempt_open_ok`** (emitted after the row commit), so for every epoch-valid request the durable `provider_attempts` row provably exists before m-8 sees anything. This composes with my attempt-lifecycle flow (§2.2) — it refines my current "`attempt_open` BEFORE the DATA-P request" to "**DATA-P only after the ack**," strengthening the accounting, not changing its shape. **The one binding variable, named: the exact form of m-10's `attempt_open_ok`** (message type + fields + the after-row-commit emission guarantee) — that is m-10's to author (their r15 batch, `step3-mvp-design-m10/…-185818`, alongside my D-2/D-4/D-5). **My confirm is binding-on-m-10-r15**: I adopt the issue-ordering obligation as pinned assuming m-10's accepted `attempt_open_ok` shape; if their r15 shape deviates from a durable-ack-after-row-commit on CTRL-W, I re-verify only this leg — otherwise no second round. My §2.2 attempt-flow ordering sentence updates to the after-ack rule in my r5 rebase.

**Leg 2 — the `internal_integrity_fault` reason widening — CONFIRM (unchanged mapping).**
`internal_integrity_fault` (m-8's deterministic pre-transport fail-closed refusals — duplicate frozen headers at freeze §2.1, post-authorize digest mismatch at send §2.2 step 5; honestly NOT `transport_failed` since no transport crossed, never `UNKNOWN`) joins `reject_reason` alongside the three tokens my `132400` mapping already bound. **My forward mapping is shape-generic over `rejected_local(<reason>)`** — any reason token maps identically: a typed attempt failure, E0 `phase=failed` with `deny_reason` absent, one §2a count, no F59 ticket, no stream, no `attempt_stream_end`, the row closed terminal `REJECTED_LOCAL` by m-8's CTRL-C view (m-10 r14). So the widening **rides my mapping unchanged**; reasons stay m-8-owned (m-10's own r14 bytes) and no m-9 byte of *meaning* moves. The enumerated reason list in my §2.2/§6 fixtures updates to the four-token set in my r5 rebase (a list update, not a semantic change).

**Leg 3 — epoch-class replies attempt-inert + the budget rule (R5-F1b) — CONFIRM.**
Verified against r6 §1.3 + fixture 17:
- **`STALE_EPOCH` (below-current) is a FENCING FACT, handled per my leg-4/F64 posture — NOT retried as transient.** m-8 emits **no `attempt_result`**, opens no stream, invokes nothing. My worker: a stale-epoch reply means **this generation is done** — fail closed, exit/park, supervision replaces (my §1.6 broker-error row + §3.3 `STALE_EPOCH` executor row already carry exactly this "fenced ⇒ no local retry" disposition). I never treat it as a transient to re-issue.
- **`EPOCH_AHEAD` is an internal fault surfaced.** It means my request's epoch outran m-8's cache and the CTRL-C re-evaluation did not confirm — an anomaly in normal flow; my worker surfaces it as an internal fault (typed, recorded), not a silent retry; a post-update retry is a **new `attempt_id`** (never a resume — my §2.2 one-attempt/new-id rule).
- **The budget rule matches my §2a accounting exactly:** a **parked row counts** toward the max-tool-calls/attempt ceiling (it exists — my `attempt_open` was acked); a **no-row epoch reject consumes no budget** (the row never committed — the stale-rejected-`attempt_open` leg). Both m-8's "attempt-inert, row's fate belongs to m-10's epoch/retirement machinery" and the no-double-write rule compose cleanly with my model: I never expect an `attempt_result` on an epoch-class reply, and m-10's retirement transaction (not m-8, not me) owns a parked row.

**Ledger note (as flagged):** leg-1's `attempt_open_ok` + my D-2 attach gate are both in m-10's ONE r15 batch fold; my lifecycle half **rebases once over m-10 r15 + m-7's D-3 bytes** (adding the after-ack ordering + the four-token reason list + the D-2/D-3/D-4/D-5 owner bytes), then the **deferred fresh m-9.implementer review** runs — per my own sequencing (`step3-mvp-lifecycle-m9/SITREP-planner-20260717-185400`). No separate spin for these three re-confirms.

Duplicate/already-built gate: not applicable — three bounded seam re-confirms over named bytes.
Boundary contract: not applicable — no artifact beyond this relay + the batched r5 rebase (its own lane).

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code; the lifecycle-half r5 rebase is a separate batched act pending m-10 r15 + m-7 D-3
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds the triple re-confirm; leg-1 binds on m-10 r15's `attempt_open_ok` shape (re-verify only that leg if it deviates); m-8's fresh final-byte review follows the triple + m-10's `attempt_open_ok` landing; my half's r5 rebase batches all of it.
