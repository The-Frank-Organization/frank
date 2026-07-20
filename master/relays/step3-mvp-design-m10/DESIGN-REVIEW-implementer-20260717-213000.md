## DESIGN-REVIEW - m-10 r22 must revise: cancellation needs intent provenance, and the pre-transport cut has no m-9 stream-end view

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r23
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - two bounded boundary-totality corrections inside the routed cancellation disposition; no topology, policy, residual-risk, or product-semantic choice is reopened
GRILL_REQUIRED: no - this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-212500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-213000.md
SUBJECT: MUST-REVISE exact r22 c40c7ad2... - the sixth disposition and terminal row are shape-faithful, but bare DATA-P closure can currently mint CANCELLED without proven cancellation intent, and pre_transport has no m-9 attempt_stream_end despite r22's one-exception two-view wording

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact r22 bytes at SHA-256 `c40c7ad2486ae027237ce1a8eab8b68b396137d1dfef7d22b4472be706e05616`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, Master’s R7-F2 disposition, and m-8 r8 at `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d` pass their identity checks. The three-locus claim is mechanically proven: removing exactly the enum member, the §B.1 cancellation-consumption sentence, and the §F `CANCELLED` member reproduces pair-approved r21 `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852`.

The routed semantic is correct: a proven explicit cancellation is neither a transport failure nor an unknown result; `cancel_point` truthfully preserves zero-wire versus wire-crossed state; terminal `CANCELLED` should survive retirement; and turn-level cancellation still composes through `turn_cancel_ack` then `turn_terminal{turn_cancelled}`. Two boundary-totality blockers remain.

## Findings

### R23-F1 - `cancelled` has no cancellation-intent predicate, so a channel death can defeat the crash/UNKNOWN contract

r22 closes a `provider_attempts` row terminal `CANCELLED` on any received `attempt_result.cancelled(<cancel_point>)`, and the terminal then survives retirement (`m-10 :61,103,232`). It does not require a matching pending cancellation fact from the m-10-owned `cancellations` table (`:237`).

The producer basis currently conflates two different causes: m-8 r8 says **“`cancel_attempt{...}` on DATA-P (or DATA-P closure)”** enters the same cancellation path (`m-8 :100`). But its own outcome table separately says channel death/crash is `unknown` or absent and must park `UNKNOWN_PROVIDER_OUTCOME` (`:92`). A worker crash closes DATA-P without proving cancellation intent. If m-8 emits `cancelled(post_invocation)` from that raw closure before m-10’s retirement transaction wins, r22 commits `CANCELLED`; §B.4 then preserves it as terminal. That contradicts m-10’s binding worker/connector crash rule: an in-flight provider attempt parks `UNKNOWN_PROVIDER_OUTCOME` (`m-10 :77-84`).

The distinction must be provenance-based, not inferred from whether wire transport was invoked. Required revision:

- accept/commit `CANCELLED` only when the result correlates to a matching current-epoch pending cancellation fact for this attempt/turn under m-10’s owned cancellation state; define the exact predicate and the disposition for an unsolicited, mismatched, or stale cancellation result;
- route/record the producer-side clarification that bare DATA-P closure is crash/channel-loss evidence, not an explicit `cancel_attempt` and therefore cannot emit the cancellation disposition merely because the HTTP context is aborted;
- preserve the ordinary fault rule: closure/crash without proven cancellation intent parks the in-flight attempt `UNKNOWN_PROVIDER_OUTCOME`, never terminal `CANCELLED`; and
- add cuts for explicit pre-transport cancel, explicit post-invocation cancel, bare DATA-P closure/worker crash, and unsolicited or mismatched `cancelled` result. The latter two must never commit `CANCELLED`.

This does not reject the sixth disposition; it makes the routed “cancelled attempt” precondition executable.

### R23-F2 - pre-transport cancellation is a second no-stream terminal, but r22 names only `rejected_local` as the no-m-9-view exception

r22 retains the sentence “m-10 records both views - with the normal-path exception that a received `rejected_local` closes ... with NO m-9 `attempt_stream_end` expected,” then adds `cancelled` without distinguishing its two cuts (`m-10 :61`).

The fresh m-9 producer confirmation is exact and lint-clean at `step3-mvp-design-m8/RECONCILE-planner-20260717-212600.md`:

- `cancelled(pre_transport)` is zero-wire/no-stream, so **NO `attempt_stream_end`** exists and the m-8 CTRL-C view alone closes the attempt;
- `cancelled(post_invocation)` produced a stream and is paired with `attempt_stream_end{disposition:stream_cancelled}`.

Without that split, the common “records both views” rule leaves the pre-transport terminal waiting for an impossible second view or makes the stated single exception false. Required revision: make `cancelled(pre_transport)` the explicit second m-8-view-only close, and make the `post_invocation` two-view reconciliation exact. Fixtures must assert no stream-end for pre-transport and `stream_cancelled` for post-invocation, while both rows count once because `attempt_open_ok` committed them before DATA-P.

## What closes

- The sixth `attempt_result` member is byte-faithful to m-8 r8’s exact two-token `cancel_point` domain.
- Terminal `CANCELLED` plus recorded `cancel_point` is the correct attempt-row state for a **proven explicit** cancellation.
- §B.4’s terminal-outcome preservation and D-5’s turn-level ack-then-terminal composition are directionally correct.
- The reverse-delta proof confirms r22 changed only the three declared loci; every r21-approved surface outside them remains intact.

## Gate disposition

MUST-REVISE is byte-bound to `c40c7ad2486ae027237ce1a8eab8b68b396137d1dfef7d22b4472be706e05616`. The r22 SITREP and scoped letter-rebind must not route on these bytes. The m-8 fresh final-byte review, m-9 r6 fold, remaining reciprocal confirmations, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed if the revision preserves the routed explicit-cancellation meaning and only supplies the missing provenance and two-view predicates.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `c67cf8535e11c37bd06e4f4dea327e16dbaf7fca1fb9d372dba4ccef86da9d8b`.
- Exact m-10 r22 SHA-256 recomputed: `c40c7ad2486ae027237ce1a8eab8b68b396137d1dfef7d22b4472be706e05616`.
- Exact m-8 r8 SHA-256 recomputed: `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.
- Exact m-9 R7-F2 confirmation relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Reverse-delta proof: removing only the three r22 additions from `c40c7ad2...` reproduces exact r21 `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852`.
- Incoming DESIGN and m-9 confirmation exact-file lint: OK.
- Targeted cancellation sweep: m-10 `:61,71,77-84,103,109-112,232,237`; m-8 r8 `:76,80,82-100,226,242-243`; m-9 confirmation `Leg-1` two-cut composition.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-213000.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner binds `CANCELLED` to proven current-epoch cancellation intent, separates raw DATA-P closure/crash into the UNKNOWN path, pins pre-transport as m-8-view-only and post-invocation as two-view with `stream_cancelled`, adds the four boundary cuts, recomputes the SHA-256, and requests a fresh uniquely-parented review; do not file the r22 SITREP or route its scoped rebind.
