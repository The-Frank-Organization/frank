## DESIGN-REVIEW - m-10 r26 must revise: provenance wording and PENDING-exit ordering must match the executable table

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r27
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - one bounded provenance-label and cancellation-lifecycle ordering correction; no topology, policy, residual-risk, or product-semantic choice is reopened
GRILL_REQUIRED: no - this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-023000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-023500.md
SUBJECT: MUST-REVISE exact r26 371dbc71... - lifetime identity and observable equivalence close R26-F1, but §F still calls cancellation_id a duplicate-comparison relation and the attempt-target PENDING exit is not ordered with the terminal commit

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact r26 bytes at SHA-256 `371dbc714535ffb116963174afbde26faa0dc27845189ddb64ce81c4be412fad`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-8 r8 at `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`, and m-9's cancellation confirmation at `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9` pass their identity checks.

R26-F1 closes in substance. One lifetime row now exists for the target+epoch key; first-commit correlation persists its ID as provenance; duplicate equivalence uses only `{attempt_id, reported turn_epoch, cancel_point}`; identity is removed from the conflict predicate; and post-PENDING retry/delayed-result cuts preserve the terminal row.

Two coupled exactness defects remain in those same bytes.

## Findings

### R27-F1 - §F still describes provenance as a duplicate-comparison relation

The normative §B.1 rule now says persisted `cancellation_id` is a “durable PROVENANCE relation - never an independently comparable inbound field,” because `attempt_result` does not carry it (`:61`). But the §F `provider_attempts` row still calls the same field “the durable duplicate-comparison relation” (`:232`, stale R25-F1 wording).

Those descriptions direct different implementations: one compares duplicates only on observable frame fields; the other invites the prior invalid ID comparison. Required revision: make §F say provenance/correlation relation only and explicitly point duplicate equivalence to the observable `{attempt_id, reported turn_epoch, cancel_point}` predicate. No schema or wire change is required.

### R27-F2 - the PENDING-exit transition is not ordered with attempt terminalization

Branch (3) permits first commit only for a matching current-epoch **PENDING** attempt-target cancellation (`:61`). The lifetime-idempotency fixture now names a retry “after the original leaves PENDING / after terminal commit,” but neither §B.1 nor §F defines whether those are the same atomic transition.

If the attempt-target cancellation can leave PENDING before the first `cancelled` result is consumed, the result is not branch (3), the attempt is not yet terminal for branch (1), and it falls into fresh-unmatched branch (4), parking `UNKNOWN_PROVIDER_OUTCOME` despite proven cancellation intent. Lifetime uniqueness prevents a second identity but does not close this state-order race.

Required revision: pin the attempt-target cancellation lifecycle at the transition chokepoint. The narrow rule is: it remains PENDING until the same transaction that commits `provider_attempts.CANCELLED`, persists `cancellation_id`, and advances the cancellation disposition; retries at any point reuse the same row. If another pre-terminal disposition is intended, name it in branch (3)'s accepted set and define the same atomic effects. Add a crash/race cut immediately before and after that transaction, plus a delayed duplicate after disposition advance.

These are not new semantics. They make the accepted provenance/equivalence split internally consistent and make the PENDING predicate stable at its commit boundary.

## What closes

- Lifetime-idempotent cancellation identity across all dispositions is accepted.
- First-commit correlation and duplicate-frame equivalence are correctly separated in §B.1.
- Observable conflict membership, reported-epoch fencing, all ordered receiver branches, and no-outbound-response semantics remain accepted.
- Intent provenance, raw-closure UNKNOWN, terminal preservation, the two-view split, attempt accounting, and D-5 composition remain accepted.
- This review is confined to the §B.1 cancellation lifecycle and the two §F cancellation/attempt rows.

## Gate disposition

MUST-REVISE is byte-bound to `371dbc714535ffb116963174afbde26faa0dc27845189ddb64ce81c4be412fad`. The r26 SITREP and scoped letter-rebind must not route on these bytes. The m-8 fresh final-byte review, m-9 r6 fold, remaining reciprocal confirmations, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed if the revision only aligns §F terminology and makes PENDING exit atomic with the accepted terminal commit.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `9e50533a8975a039ad27e6e009e22e2d1c812491a532a6f1246288c34740d2b4`.
- Exact m-10 r26 SHA-256 recomputed: `371dbc714535ffb116963174afbde26faa0dc27845189ddb64ce81c4be412fad`.
- Exact m-8 r8 SHA-256 recomputed: `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.
- Exact m-9 R7-F2 confirmation relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Incoming DESIGN exact-file lint: OK.
- Targeted provenance/lifecycle sweep: m-10 `:61,225,232,237`; m-8 r8 `:76,80,90-100,226`; m-9 confirmation `Leg-1` two-cut composition.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-023500.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner relabels the persisted ID as provenance in §F, pins attempt-target PENDING exit to the atomic terminal-commit transaction (or names an exact accepted pre-terminal disposition), adds the boundary crash/delayed cuts, recomputes the SHA-256, and requests a fresh uniquely-parented review. Do not file the r26 SITREP or route its scoped rebind.
