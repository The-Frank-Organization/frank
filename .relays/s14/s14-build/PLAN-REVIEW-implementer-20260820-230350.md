## PLAN-REVIEW — s14 m-8 connector build plan r1: MUST-REVISE; separate RFC 8785 from artifact-specific NFC validity and make the fixture map honest at the s14 boundary

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-1
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the planner must revise; dependency or locked-contract arbitration routes through master if the revision cannot remain inside the commissioned fence
GRILL_REQUIRED: no — this review tests the plan against frozen contracts and introduces no design decision
FILED_AT_LOCAL: 20260820-230350
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260820-223912.md
PLAN_LOCK_ID: s14-build @ sha256 6442936b5b36ef6352152a3f74208a2a0f246565e3df14841187a55422077b57
PLAN_REVIEW_VERDICT: must-revise
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: must-revise — split JCS from NFC validity, resolve the stdlib-only NFC strategy, and classify cross-lane fixture legs instead of claiming local completion

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`. The fifteen-task ordering, write fence, freeze/authorize/attach/send sequence, deny-token order, transport pins, B/E presence rules, opacity rule, and implementation/review role split are materially aligned with the frozen bases. Two blocking defects prevent a literal implementation dispatch. No source/test byte, branch, worktree, dependency, stage, or commit was created or changed in this review.

## S14-PR1-F1 — BLOCKER: T1 conflates RFC 8785 with the separate NFC validity rule, while the stdlib-only pin leaves NFC unimplementable as planned

T1 specifies an `internal/connector/jcs` implementation as “RFC 8785: NFC strings” (`PLAN:42`). The frozen m-3 contract says the opposite at the mechanism boundary: JCS settles encoding, while NFC is an explicit P0 validity precondition and “JCS does not normalize” (`master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md:74-78`). A generic canonicalizer must therefore preserve JSON string values according to RFC 8785; it must not silently normalize them or globally reject strings that are valid on non-policy surfaces. NFC rejection belongs in the validators for artifacts whose frozen schemas require it, including the policy and catalog loci named by their contracts.

The same plan pins `stdlib-only`, `go.mod`/`go.sum` byte-untouched, and an in-fence JCS implementation (`PLAN:30-32,68`). Live E1 shows `go doc unicode` exposes rune-property/case APIs but no normalization API, while `go list golang.org/x/text/unicode/norm` fails with “no required module provides package”. GOROOT's vendor copy is not an importable module dependency for this repository. Thus the plan neither names a conformant NFC validator nor permits the ordinary `golang.org/x/text/unicode/norm` dependency that could supply one.

Required revision: define two distinct mechanisms and tests:

1. Pure RFC 8785 canonicalization that preserves strings, with canonical numbers, escaping, unique object keys, member ordering, and the frozen byte vectors.
2. Artifact-specific NFC validation at every schema locus that requires it, rejecting non-NFC without rewriting accepted bytes.

For item 2, either route a durable `go.mod` dependency arbitration through master and return its ruling, or specify a vetted, version-pinned, in-fence NFC-checking algorithm/table with exhaustive conformance vectors sufficient to justify the stdlib-only claim. Do not describe normalization as a JCS operation and do not silently normalize malformed input.

## S14-PR1-F2 — BLOCKER: acceptance claims locally complete fixtures whose frozen contracts explicitly assign legs to sibling owners

Acceptance requires every r12 §8 fixture 1–17b and every B/E addendum §6 fixture 1–11 to be “implemented” and “swept complete” inside the s14 fence (`PLAN:55-60`). The frozen contracts prohibit that unqualified claim:

- r12 fixture 12 distinguishes m-8-realized sentinel legs from m-9 prompt/tool-output legs, m-10 launch-path legs, and the cross-lane conductor-record harness; it explicitly says m-8 “does NOT claim them realized here” (`master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md:225`).
- r12 fixture 13b is owned by the m-7/m-2 conductor-entry harness, not locally realized by m-8 (`provider-contract.md:226`).
- r12 fixture 17b includes m-10 durable-row results, m-9 stream/turn/E0 carriage, `pending_app_events`, retirement/reap races, and worker-crash effects beyond s14's write fence (`provider-contract.md:230`).
- addendum §6.2, §6.4, §6.6, and §6.7 contain m-3 observer/evaluator derivations over independently captured requests; the addendum says P2b is m-3-owned, not m-8-owned (`master/domains/m-8-provider-adapters/design/2026-07-22-stage6-BE-digests-addendum.md:144-155`; owner split at `:25-27`).

Required revision: add a per-fixture boundary matrix that classifies every leg as one of:

- s14/m-8 production-realized E2;
- s14 producer output or fake-counterpart contract assertion;
- restack/integration obligation, naming the consuming slice and rerun gate; or
- sibling-owned m-2/m-3/m-7/m-9/m-10 proof that s14 must not claim locally.

Bind T12–T15 and §3 acceptance to that matrix. Local completion may assert all m-8-owned legs and all fake-counterpart outputs green; full cross-lane fixture closure remains a named restack/integration predicate. Preserve r12's explicit “no local fixture claims a cross-lane row” rule.

## Passed surfaces

- The plan hash reproduced exactly as `6442936b5b36ef6352152a3f74208a2a0f246565e3df14841187a55422077b57`; exact-file lint passes with historical freshness disabled.
- The cited frozen basis hashes match the commissioned charter, and no connector implementation exists at `LAUNCH_BASE main@b7f406b2` that would make this a duplicate-build plan.
- The nine-token P0→PS→P1→P2→P3→P4 precedence, byte-equality endpoint comparison, post-authorize credential attachment, owned no-retry HTTP transport, B/E carrier split, `refusal_stage` invariant, and opaque-item round trip are correctly decomposed at plan altitude.
- The in-fence frame codec plus named s13 `appipc` restack evaluation is acceptable as a plan, provided any post-restack shared-seam change is escalated rather than silently authored by s14.

## Boundary contract review

Writes: `cmd/frank-connector/**`, `internal/connector/**`, and pair-local `.relays/s14/**` only.
Reads: frozen m-8 provider/B-E contracts plus m-1, m-3, m-9, m-10, and join/carriage bases; fake counterpart CTRL-C/DATA-P frames.
Target entity: the m-8 connector process, catalog/policy validators, authorization/freeze/attach/send path, provider event normalization, and attempt reporting.
Downstream consumers: m-10 supervisor, m-9 worker, m-3 observer/evaluator, and the later restacked `appipc` seam.
Contract: exact framed carriers, freeze-boundary digests, credential non-injection, deny-to-zero-send, epoch fencing, and normalized event/outcome schemas.
Proof: revised task-level RED/GREEN E2 batteries plus a fixture-boundary matrix; fake counterparts prove s14 outputs, while sibling-real effects remain integration obligations.
No-consumer action: reject dispatch and do not claim cross-lane closure until the revised plan separates local producer proof from downstream-real proof.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — provider credential custody and the authorize-before-attach boundary
- migration/backfill/destructive-write/canonical-data-repair: no — no database/data migration or destructive repair is commissioned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — provider-send authority, attempt truth, and evidence digests are trust-critical state
- AI-or-automation-acts-downstream: yes — the connector sends model requests and normalizes provider output
- worker/scheduler/queue/retry/async-side-effect: yes — supervised process, bounded channels, deadlines, cancellation, streaming, and one-shot provider effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-8/m-9/m-10/m-3 framed contracts and the s13 `appipc` restack seam
- user-visible-control-with-materializer/downstream-consumer: no — this plan introduces no user-visible control
- test-runtime-role-mismatch: yes — local fake-counterpart proof must not be reported as sibling-runtime proof
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — unresolved NFC implementation strategy and overbroad cross-lane fixture closure
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; retain production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane-governance action only — this PLAN-REVIEW file plus its append-only s14 INDEX row; no source/test edit, branch, worktree, dependency, stage, commit, or implementation adoption
FINAL_GIT_STATUS_SHORT: (literal, frank/ cwd, captured after this relay and its INDEX row existed:)
 M .relays/s13/INDEX.md
 M .relays/s13/docs/designs/DS-s13-m10-module-20260820.md
 M .relays/s13/docs/plans/PL-s13-build-plan-20260820.md
 M .relays/s14/INDEX.md
 M .relays/s15/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-221508.md
?? .relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-223944.md
?? .relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-224152.md
?? .relays/s13/s13-build-design/DESIGN-planner-20260820-223212.md
?? .relays/s13/s13-build-design/SITREP-planner-20260820-223211.md
?? .relays/s14/s14-build/
?? .relays/s15/s15-build-2/
?? .relays/s15/s15-build-3/
?? .relays/s15/s15-build/PLAN-REVIEW-implementer-20260820-220110.md

Next requested action: `s14.planner` returns a successor PLAN closing S14-PR1-F1 and S14-PR1-F2. Only an approved successor PLAN-REVIEW may parent the literal direct implementation dispatch.
