## RECONCILE return — m-10's three confirmation legs: m-7 CI-1/2/3 + R9/R10 composition CONFIRM @ `f072bd99…` · m-2 F58 sufficiency CONFIRM @ `83d8e63e…` · m-1 lifecycle fit CONFIRM @ `7c8b09a6…` — zero findings; byte-bound; no lock claimed

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-confirm-m10
PARENT_DISPATCH_ID: step3-mvp-confirm-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded byte-bound confirmations; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — per the routing relay
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: step3-mvp-confirm-m10/RECONCILE-orchestrator-planner-20260717-010000.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-7.planner, m-7.implementer, m-2.planner, m-2.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: all three legs CONFIRM, each at the named exact bytes, against my approved `79fcf742…`; one composition reading named on leg 1 (not a finding); the two-counter/launch-custody/no-secret-store obligations of m-1 §3 accepted verbatim

Confirming seat + basis: m-10.planner; my side of every comparison = the pair-approved `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` @ `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453` (unchanged — this relay edits nothing).

### Leg 1 — m-7 @ `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e` (recomputed exact) — **CONFIRM**
- **CI-1** (dial-in `broker-control.sock` listener; m-10 as CLIENT; fcntl `broker-control.lock` → durable `control_generation` advance → connect ordering; spawn-pipe token handoff, rotation at broker spawn only): consumed verbatim in my §A.1 + §F `broker_control` row. Confirmed compatible — no socketpair-inheritance assumption survives for the broker channel on my side.
- **CI-2** (`assign` gains `generation_id` + `broker_worker_endpoint`): present in my §B.1 handshake, with my post-lease-bind assign timing layered on top (m-7's attach tuple-equality check is unchanged by when I send assign). Confirmed.
- **CI-3** (`broker_events` UNIQUE(nonce, event_seq) + same-ack-on-duplicate; crossing rows keyed {transition, operation} with the five-state disposition; the queryable transition ledger, one-transaction abort, at-most-one-non-terminal-per-run): consumed verbatim in my §B.5 + §F rows. Confirmed.
- **R9/R10 composition (the reading, named so the reconcile can check it):** my R9 pending-transition order needs NO m-7 byte change — it composes from m-7's own §2.4 fail-closed floor ("no installed state ⇒ suspended") + §2.5 PREPARING-barrier persistence + §2.7 app-main-crash row ("barrier PERSISTS; resume on installed/reconciled state") + m-10 being the sole state supplier: I withhold the install-eligible snapshot while a transition is pending, and their broker sits at its own suspended floor until my `CROSSERS_DURABLE` ack. My R10 lost-install replay rides their §2.5 committed⇒resume (re-ack the exact set, idempotent install) + §2.11 duplicate-key-same-ack. If m-7 reads their §2.10 adoption/bootstrap "installs the snapshot" as UNCONDITIONAL even when the controller withholds one, that would be a finding — my reading is that absence-of-snapshot ⇒ their §2.4 suspended floor, which their bytes state. Flagged as a reading to verify at the reconcile, not a dispute.
- Also noted, no action: m-7/m-1 fence `Describe` alongside the three verbs; my epoch-state supply is operation-agnostic, so this changes nothing on my side.

### Leg 2 — m-2 @ `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d` (recomputed exact) — **CONFIRM (sufficient for F55 + F63)**
- §3.1 byte-exact canonical names + alias normalization; digests = 64 lowercase hex over the pinned static TEMPLATE (run-varying slots normalized out — correctly immune to `form_digest`/volatile churn, so the member survives run-immutability); `m2-mapping-v<N>` grammar monotonic; **§3.4 absence rule matches my §C.1 exactly** (ABSENT member for local tools; presence on a local tool / absence on a relay verb = a SHAPE violation at my gate, not a value comparison). My F55 gate compares all present members byte-exact after normalization — nothing in their encoding requires my gate to parse form semantics, which was my §C.4 blocking-mismatch criterion. Clean.
- **F63 sufficiency:** their two-event split (lock records their doc hash + expected fingerprint bound to the locked version; the F63 event executes the SHIPPED artifact over the locked Appendix-A inputs and compares against lock-recorded expected bytes — never both sides from T4) supplies exactly the expected-vs-shipped comparands my §C.3 release check consumes. My serve gate compares the resulting vectors; the fingerprint execution is the F63 event's job, not my gate's. Confirmed sufficient.

### Leg 3 — m-1 @ `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` (recomputed exact) — **CONFIRM (lifecycle fit + my §3 obligations accepted)**
- **Two-counter law (§2.2):** fits my §B.4 exactly — `turn_epoch` advances automatically at retirement (mine), mint-generation only by operator act (theirs); nothing in my contract causes/requires/implies a `seat_mint`; m-10 holds no mint authority anywhere in my bytes (assign/connector_assign carry no credential; broker handover rides the third counter `control_generation`, coupled to neither).
- **§2.7 matrix vs my supervision, row-for-row:** worker replacement ⇒ epoch +1 exactly (my retirement mints exactly once; a pre-lease wash-out is not a replacement of an active worker and mints none — consistent with "+1 exactly"); broker restart ⇒ epoch unchanged, no reset/reuse (mine: transition rows `unknown-outcome` + durable abort; the already-published epoch stands; fresh ID for the SAME E+1); conductor restart ⇒ untouched on my side; re-mint ⇒ epoch unchanged (conductor-invisible to my lifecycle); overlap ⇒ each counter advances exactly once (independent transactions, no coupling in my bytes). No cell contradicted.
- **My §3 obligations accepted verbatim:** opaque-reference-only custody orchestration (my §G.2 — references are secret-inert AND authority-inert per their §1.4a, matching my design's assumption); launch-path hardening (my §B.1 sanitized-env/no-secret-argv/close-on-exec); epoch state supplied to the broker gate with no credential bytes and no conductor verb; **no secret-bearing column exists in my store by schema** (my §F/§G.2 — and their "NOT secret" census confirms `turn_epoch`/manifests/tickets/references are safe app-side, which is everything my store holds).

### Disposition
All three legs: **CONFIRM**, zero findings. One leg-1 composition reading flagged for the reconcile to verify with m-7 (above). My approved bytes are unchanged; any change on any side re-opens that pair's review per the routing rule. Stage-3 lifecycle halves may route on the leg-1 confirm whenever master directs; my stage-5 DESIGN (+ grill) still waits on its stage-3 inputs.

## Verification
- Hashes recomputed this session: m-7 `f072bd99…` · m-2 `83d8e63e…` · m-1 `7c8b09a6…` (all exact); my doc `79fcf742…` (unchanged).
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/step3-mvp-confirm-m10/RECONCILE-planner-20260717-013000.md` — run at filing; result reported inline.

ACTIONS_GIT_REF: none — a report-only confirmation relay + one INDEX.md row timestamped 20260717-013000; no design-doc edit, no `frank/` edit, no code, no lock, no credential, no provider call.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: master records the three confirms in the stage-1 confirmation table, verifies the leg-1 composition reading with m-7, and routes the stage-3 lifecycle halves; m-10.planner stands by.
