## m-6.implementer consumer review of c2 m-3/m-4 designs

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-consumer-review-m-6
PARENT_DISPATCH_ID: c2-consumer-review-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- read-only consumer review
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-4.planner, operator
BUNDLE_ID: c2-consumer-review-m-6
OWNER: m-6 (Human Surface & Scheduler), consumer lens

PRIMARY_BUCKET: already-closed
still-open: none blocking the c2 lock from the m-6 consumer lens.
already-closed: m-3 and m-4 expose the required reader-facing writers for observe-veto bucket projection, away-mode egress gating, Owner Decision Brief evidence content, routing gate categorization, and rare routing A-escalation choices.
product-overlapped: none; m-3 owns observe/egress/evidence atoms, m-4 owns routing record/policy content, and m-6 remains the downstream surface/scheduler owner for email buckets, ODB rendering, away-mode inbox, and scheduler policy.
recommended-next: m-6.planner can reconcile as sufficient for c2 lock, carrying the noted future m-6 design obligations without reopening m-3 or m-4.

VERDICT: sufficient

## Evidence Basis

- Dispatch scope: `c2-consumer-review-m-6` directly addresses `m-6.implementer` for read-only AUDIT consumer review and asks m-6 to verify m-3/m-4 writer surfaces before c2 lock in master/relays/c2-consumer-review-m-6/AUDIT-orchestrator-planner-20260629-212435.md:1-35.
- Locked m-6 forward contract: away-mode external bridge mirrors A-bucket gates to the operator's real inbox and is gated by fail-closed egress scanning before external send in master/ARCHITECTURE.md:79-87.
- Locked m-6 forward contract: `gate_category`, A/B mapping, protected-branch set, `other` to A fail-safe, and protected-branch merge split are operator-configurable m-6/config surfaces in master/ARCHITECTURE.md:89-102.
- m-3 observe surface: `observe_gate()` returns `predicate_result`, `veto`, observed fields, achieved evidence, executable claim results, `egress_scan_result`, and degradation notes in master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:41-58.
- m-3 write/veto boundary: the observer-only allowlist writes `egress_scan_result`, routing-profile observed set, and veto; veto causes include false done-predicate, declared-vs-observed integrity mismatch, and failed egress scan in master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:61-64.
- m-3 egress gate: the conductor is sole external sender, scans outbound body and proof artifacts for secrets/PII/model-name/auth-URL classes, and blocks plus holds/resummons on any flag in master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:111-118.
- m-3 ODB content: m-3 explicitly writes to m-6 the observe-veto bucket projection, away-mode egress gate, and Owner Decision Brief evidence bundle with completed-proof, residual risk, recommendation, exact choices, conductor-observed evidence atoms, and `record_integrity` label in master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:141-147.
- m-4 routing surface: `route_dispatch()` fails closed by emitting `human_decision_required` or `routing_unavailable`, and must not silently fall back to a default model in master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:141-158.
- m-4 routing FieldSpec: routing assignments include seat, role, task tag, declared bucket, chosen model, declared deviation, and pin mode; consumers include `human_surface`; model values stay payload, not predicates in master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:189-210.
- m-4 m-6 contract: `routing` is category B by default, escalates to A only on `human_decision_required` or `routing_unavailable`, and exposes routing `gate_category` plus ODB recommendation and enumerated model choices for rare A escalation in master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:308-323.
- m-4 config pattern: bucket membership and recommendations ship default-seeded and operator-configurable, matching the locked m-6 `gate_category` policy pattern in master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:162-185.

## Surface Verdicts

1. Observe-gate veto to gate email bucket: sufficient. m-3 exposes the mechanical writer surface m-6 needs: `veto`, `predicate_result`, failing claim/results, `target_gap_result`, `egress_scan_result`, and bounce text commitments. m-6 can derive a bucket from record context rather than asking m-3 for an agent-fillable email bucket. For example, a failed evidence/claim observe-bounce stays a conductor bounce/local correction path; a failed egress scan or human decision floor can project to A because it holds/resummons the operator.

2. Egress/content-safety gate for away-mode external bridge: sufficient. m-3 does not implement the away-mode bridge itself, but it exposes the fail-closed pre-external-send gate at the conductor chokepoint and names the result field. That is the writer m-6 needs for the first external send; m-6 later owns the external inbox bridge, notification cadence, and reply ingestion.

3. Owner Decision Brief evidence-summary content from m-3: sufficient. The design carries the prior seven-part ODB pattern and the m-6-critical content: `completed_proof` as evidence-backed proof, residual risk, recommendation, exact choices, evidence atoms, and `record_integrity`. This is enough for m-6 to render an operator-facing ODB without treating prose claims as observed proof.

4. Routing `gate_category`: sufficient. m-4 keeps ordinary `routing` in category B and escalates to A only for `human_decision_required` or `routing_unavailable`. That maps cleanly to locked m-6 buckets: B for orchestrator-absorbed routing; A for no safe route or a human-only routing decision. The fail-closed no-default rule prevents an unbucketed silent fallback.

5. Rare routing A-escalation ODB content: sufficient with ownership boundary. m-4 should expose routing-specific content, not the whole ODB renderer. Its recommendation plus enumerated model/bucket choices are the routing-specific choices m-6 needs; the surrounding ODB fields can be composed by m-6 from m-4's routing record (`routing_assignments`, `capability_prior_snapshot`, `justified_deviation`, `deviation_reason_code`) plus m-3's evidence/record-integrity layer. No m-4 contract gap blocks lock.

6. Section J forward requirements: not foreclosed. The operator-configurable A/B map and protected-branch set remain m-6/config-owned in the locked c1 contract, while m-4 follows the same config-sourced pattern for routing priors. The away-mode external-inbox bridge remains m-6-owned and is not blocked by m-3 because the egress gate is dormant until first external send and then becomes the required gate.

## Gaps

No blocking m-3 or m-4 contract gaps found.

Non-blocking carry-forward for m-6 design: m-6 should specify the concrete derivation table from observed veto/egress/routing outcomes to A/B/C/D email buckets, including whether pure observe-bounces are local correction only, lint-bounce style D, or A when the bounce reason itself requires operator judgment.

## Operator-Judgment Items

None blocking for c2 lock.

Future m-6 design should ask the operator for notification intensity and escalation cadence for away-mode inbox mirroring, plus the editable default `gate_category` map and protected-branch set. Those are already reserved to m-6/config and do not require m-3/m-4 changes now.

## Coordination

No direct gap relay sent to m-3 or m-4. This is the implementer-side challenge input; `m-6.planner` still owns the final reconciled m-6 consumer-review relay if the pair follows the dispatch's lead/reconcile method.

## Relay Lint

RELAY_LINT: OK -- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c2-consumer-review-m-6/AUDIT-implementer-20260629-213803.md`

RELAY_ROOT_LINT: non-clean due to `INDEX.md` only -- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c2-consumer-review-m-6/AUDIT-implementer-20260629-213803.md` reports OK for this relay file, then errors on lint-exempt `master/relays/INDEX.md` missing relay header fields.

## Actions

ACTIONS_GIT_REF: relay artifact created at master/relays/c2-consumer-review-m-6/AUDIT-implementer-20260629-213803.md; git status unavailable because cwd is not a git repository.

FINAL_GIT_STATUS_SHORT: unavailable -- fatal: not a git repository (or any of the parent directories): .git
