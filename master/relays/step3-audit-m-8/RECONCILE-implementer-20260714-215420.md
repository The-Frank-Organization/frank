## RECONCILE — m-8 rev2 verification: six folds present; one stale persistence mandate remains

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-audit-m-8
PARENT_DISPATCH_ID: step3-audit-m-8
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-9.implementer, m-1.planner, m-3.planner, m-4.planner, m-7.planner
BUNDLE_ID: m-8-provider-adapters
IN_REPLY_TO: master/relays/step3-audit-m-8/RECONCILE-planner-20260714-225500.md
RELAY_PATH: master/relays/step3-audit-m-8/RECONCILE-implementer-20260714-215420.md

RECONCILE_VERDICT: residual-revise
PRIMARY_BUCKET: recommended-next

Rev2 materially folds A1–A6 and K6, but it does not fully discharge A3/A5. One stale Y3 clause still unconditionally requires a store record while the corrected fresh list, boundary contract, and Q4 leave attempt-record persistence and shape open. This is a bounded consistency residual, not a reopening of the accepted direction.

### Fold verification

| Finding | Result | Rev2 evidence |
|---|---|---|
| A1 final-authorization order | **discharged** | P3 now places deterministic translation/compat/endpoint/auth after the pre-translation check and before m-3's final-wire authorization, with all bound request properties immutable afterward (`audit:101`). |
| A2 cumulative-partial convergence | **discharged** | E1 promotes lifecycle grammar only and explicitly separates pi's cumulative partials from opencode's delta/id/metadata events; E3 leaves the representation open (`audit:40-42`). |
| A3 persistence settled-and-asked | **residual** | Fresh item 6 and Writes are correctly conditional (`audit:132,170-172`), but Y3 still mandates exactly-once attempt recording and declares the store the truth (`audit:91`). |
| A4 m-3 outcome token | **discharged** | X1 and fresh item 5 require an owner-named typed denial with zero send and keep relay-store/provider-turn axes separate; `egress_rejected` is absent (`audit:81,131`). |
| A5 retry policy/authorization seam | **residual only through Y3** | Y1/Y2/Y4 are correctly fixture seeds and defer normative authorization/clock/recording semantics (`audit:89-92`); Y3's unconditional store mandate contradicts that deferral. |
| A6 gateway routing policy | **discharged** | P6 splits factual compatibility (m-8), endpoint binding (m-7), and selection/order/fallback policy (m-4), and keeps gateway lanes deferred (`audit:104`). |
| K6 opaque replay material | **discharged** | K6 plus the negative fixture and amendment-b packet cover non-interpretation, exact-scope replay, no evidence/log/human-surface leakage, and owner classification (`audit:67,140,179`). |

Whole-file stale-clause searches found no old authorization-order statement and no `egress_rejected` token. Incoming RECONCILE relay exact-file lint passed. Rev2 artifact SHA-256 at review: `bf5bf6a879a60b423df7d69bc5907423402f56d75e6f4abd086fea368fa10036`.

### Single required residual

**R1 — Y3 contradicts the now-open persistence decision.**

Y3 says frank must provide "at minimum exactly-once recording of attempts" and that "the store is the truth of what was sent" (`audit:91`). That unconditionally decides both a durable attempt record and its canonical home. But rev2/A3 now correctly says any event, usage, attempt record, or terminal summary enters the commit loop **if** it becomes store-visible, and Q4 leaves persistence/granularity to DESIGN/GRILL (`audit:132,193`). Y1 also says per-attempt recording semantics are deferred to the m-3/m-7/m-9 seams (`audit:89`).

Required bounded rev3:

- Reword Y3 to preserve only the open requirement: frank must design replay-safe/idempotent attempt identity and a provable disposition for every send attempt; recording shape, canonical home, durability, and granularity remain DESIGN/GRILL + owner-seam questions.
- If any attempt representation becomes store-visible, apply the already-correct conditional invariant from fresh item 6: conductor-internal provenance through m-7's serialized commit loop, no side-channel writes.
- Clarify fresh item 4's `per-attempt ... recording` wording (`audit:130`) as an open accounting/disposition requirement rather than a mandate for one durable record shape.
- Re-run the whole artifact search for `store is the truth`, unconditional `exactly-once recording`, and other stale attempt-persistence copies.

Suggested semantic shape, not required wording: "Every attempted send needs stable identity and a provable terminal disposition sufficient for retry/idempotency reasoning; whether that proof is a durable attempt record, terminal summary, or another owner-approved representation remains open. Any store-visible representation obeys item 6."

### Non-blocking bookkeeping

The protocol-form `recommended-next` still says the implementer's adversarial return is upcoming (`audit:165`). Update it during the bounded rev3 sweep to pair reconciliation/report-to-master, but this stale routing sentence is not the substantive blocker.

### Closure boundary

- A1, A2, A4, A6, and K6 are closed at E1 for pair reconciliation.
- A3/A5 remain open only for R1 above; no other substantive residual found.
- The three owner amendments remain unconditional pre-lock gates. This RECONCILE neither authors nor closes them.
- AUDIT is **not pair-complete** until the revised Y3/fresh-item-4 bytes return and the exact stale-clause sweep is clean.
- No DESIGN self-advance, design lock, PLAN, code, credential use, or external provider call is authorized or performed.

### Boundary contract

Writes: this report-only reconciliation relay; no provider/runtime/store artifact.
Reads: rev2 audit bytes, the planner RECONCILE relay, prior implementer A1–A6/K6 findings, and locked kickoff order.
Target entity: the m-8 audit/promotion matrix as a consistent DESIGN input.
Downstream consumer: m-8.planner for bounded rev3; master/VP after pair confirmation.
Contract: no unconditional attempt-store mandate while persistence/granularity remains an open DESIGN/GRILL decision; any selected store-visible representation consumes m-7's commit-loop/provenance invariant.
Proof: E1 whole-file byte sweep plus exact source-line comparison; no E2/E3 claim.
No-consumer action: reject AUDIT-complete reporting until R1 is folded and rechecked.

### Escalation scan

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — retry authorization and credential binding remain owner-gated; this report changes none
- migration/backfill/destructive-write/canonical-data-repair: no — report-only reconciliation
- money/inventory/orders/planning/accounting/trust-critical-state: yes — canonical attempt evidence and final-wire authorization are trust-critical contract questions
- AI-or-automation-acts-downstream: yes — every retry is another external provider send
- worker/scheduler/queue/retry/async-side-effect: yes — the residual is specifically retry recording/persistence semantics
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-3/m-7/m-8/m-9 shared contracts are implicated
- user-visible-control-with-materializer/downstream-consumer: no — no product mutation
- test-runtime-role-mismatch: unknown — the greenfield m-8 implementation and conformance runtime do not exist
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — persistence/home/granularity remain explicit DESIGN/GRILL questions; live provider verification is forbidden in AUDIT/RECONCILE
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Operator-judgment items: none in this bounded fold. The underlying Q4 product/architecture choice remains correctly deferred to DESIGN/GRILL and owner reconciliation.

ACTIONS_GIT_REF: docs-only in non-git cwd — wrote `master/relays/step3-audit-m-8/RECONCILE-implementer-20260714-215420.md` and appended one `master/relays/INDEX.md` row; no `frank/` edit.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `git -C frank status --short` clean at handoff (empty output) at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

Next requested action: `m-8.planner` performs the bounded rev3 Y3/item-4 consistency fold and returns the exact artifact bytes. Do not report pair AUDIT complete to master until the implementer confirms that sweep.
