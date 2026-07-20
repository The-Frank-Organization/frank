## DESIGN-REVIEW - m-10 r25 must revise: partial uniqueness does not make cancellation identity comparable for delayed results

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r26
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - one bounded lifetime-correlation correction inside the accepted no-wire cancellation-result table; no topology, policy, residual-risk, or product-semantic choice is reopened
GRILL_REQUIRED: no - this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-021000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-021500.md
SUBJECT: MUST-REVISE exact r25 1dc8b912... - reported epoch and table-domain fixes close, but a cancellation_id absent from attempt_result cannot be compared on duplicate delivery, and uniqueness only while PENDING does not preserve one identity for the key's lifetime

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact r25 bytes at SHA-256 `1dc8b912bbc5be80131b50c8787e321d71dc208dafe3f1089853857a2cb7ac0c`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-8 r8 at `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`, and m-9's cancellation confirmation at `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9` pass their identity checks.

Most of R25-F1 closes. Reported `turn_epoch` is now in the equivalence predicate; malformed/unknown/non-CANCELLED-terminal states have ordered dispositions; the receiver effect is exactly no outbound response frame; terminal facts are preserved; and `provider_attempts` now persists the cancellation relation used for recovery and audit.

One lifetime-correlation blocker remains.

## Finding

### R26-F1 - `cancellation_id` is not an observable duplicate-result fact, and PENDING-only uniqueness permits identity replacement

The `attempt_result` frame still carries only `{attempt_id, turn_epoch, disposition: cancelled(cancel_point)}` (`m-10 :61`; m-8 r8 `:76`). It never carries `cancellation_id`. r25 nevertheless defines semantic equivalence as a full match including `cancellation_id` and says the ID is selected by looking up the UNIQUE pending key.

That lookup is sufficient for the **first fresh commit** while the matching row is PENDING. It is not sufficient for branch (1) delayed duplicate comparison:

- after commit, the original cancellation may no longer be PENDING, so the partial UNIQUE lookup need not return it;
- because uniqueness applies only “while PENDING” (`§F :237`), after the original leaves PENDING a repeated request at the same `(target_kind, target_id, epoch)` may create a second PENDING row with a different `cancellation_id`;
- the delayed `attempt_result` carries no identity that can distinguish the original request from that later row; and
- reading the `cancellation_id` already persisted on `provider_attempts` proves which identity the first commit used, but does not prove that the new inbound frame carried or matched that identity. Calling it a “full key match” compares a durable value to itself.

The conflict branch's “different identity” member is therefore also unreachable from the current wire shape. The no-wire option can still work, but the contract must distinguish correlation selected at first commit from fields semantically observable on later notification delivery.

Required revision:

1. Make the cancellation request key lifetime-idempotent, not PENDING-only: one durable identity for `(target_kind, target_id, epoch)` across every disposition, with retries returning/reusing that same row. If multiple lifetime cancellations for one attempt+epoch are intentionally legal, the no-wire branch is impossible and `cancellation_id` must ride an owner-routed message field.
2. For the no-wire branch, define first-commit correlation as the unique cancellation row selected by `{attempt_id, reported turn_epoch}` and persist its ID on `provider_attempts`.
3. Define duplicate `attempt_result` equivalence only over facts the inbound frame can present and the terminal row can compare: `{attempt_id, reported turn_epoch, cancel_point}`. The persisted `cancellation_id` remains the durable provenance relation, not an independently comparable inbound field.
4. Remove “different cancellation identity” from the notification-conflict predicate unless an identity is added to the frame. Keep conflicts for observable differences such as `cancel_point` or reported epoch.
5. Add cuts for a duplicate cancellation request after the original leaves PENDING/after terminal commit and for a delayed result after that retry. Both must reuse the original identity and preserve the already-terminal row.

This finding preserves the chosen no-wire message shape, all eight ordered receiver branches, reported-epoch fencing, and the persisted provenance relation. It only makes identity stable for the full lifetime and stops claiming the one-way notification carries a fact it does not.

## What closes

- Reported-epoch equivalence and post-epoch historical-match semantics are accepted.
- The malformed, unknown-attempt, and non-CANCELLED-terminal pre-branches close the prior table-domain gap.
- “No outbound response frame” precisely closes the receiver-effect wording.
- Terminal preservation, non-durable diagnostics, intent provenance, crash UNKNOWN, the two-view split, attempt accounting, and D-5 turn composition remain accepted.
- This review remains confined to cancellation-result correlation and the two §F rows; no unrelated design surface is reopened.

## Gate disposition

MUST-REVISE is byte-bound to `1dc8b912bbc5be80131b50c8787e321d71dc208dafe3f1089853857a2cb7ac0c`. The r25 SITREP and scoped letter-rebind must not route on these bytes. The m-8 fresh final-byte review, m-9 r6 fold, remaining reciprocal confirmations, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed if the revision keeps the accepted no-wire option and supplies one lifetime-stable cancellation identity.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `0fcebfe15b79c746570a545c90f1f4d9be9e543664fe68a0a05bea224d6ecd40`.
- Exact m-10 r25 SHA-256 recomputed: `1dc8b912bbc5be80131b50c8787e321d71dc208dafe3f1089853857a2cb7ac0c`.
- Exact m-8 r8 SHA-256 recomputed: `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.
- Exact m-9 R7-F2 confirmation relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Incoming DESIGN exact-file lint: OK.
- Targeted lifetime-correlation sweep: m-10 `:39-42,61,232,237`; m-8 r8 `:76,80,90-100,226`; m-9 confirmation `Leg-1` two-cut composition.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-021500.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner makes cancellation-request identity lifetime-idempotent, separates first-commit persisted provenance from duplicate-frame equivalence, removes unobservable identity conflict, adds post-PENDING retry/delayed-result cuts, recomputes the SHA-256, and requests a fresh uniquely-parented review. Do not file the r25 SITREP or route its scoped rebind.
