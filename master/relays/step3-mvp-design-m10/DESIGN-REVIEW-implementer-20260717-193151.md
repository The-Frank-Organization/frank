## DESIGN-REVIEW — m-10 stage-1 IPC/manifest seam r16 must revise: D-4 has no authority-bearing disposition path, and three frame families remain non-exact

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r17
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the planner can close these as bounded interface corrections, including choosing m-9's state-only D-4 branch; retaining a new runtime operator→m-10 authority ingress would instead require Master/operator disposition
GRILL_REQUIRED: no — this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-193000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-193151.md
SUBJECT: MUST-REVISE exact 859cc7b6... — admission placement and cancellation ordering are repaired, but D-4 cites a forbidden/non-transitive authority path, D-5 has no exact replies, D-2's by-name values differ from m-7, and attempt-open invariant failures still have no named disposition

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r16 bytes at SHA-256 `859cc7b69c982e892c87a21b97cb04113558d0c89cf5aa2736e1c50725271e21`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, both ratified amendments, m-8 r6, m-9 r4, and the live m-7 r10 bytes pass their identity checks. r16 corrects the broad placement and state-machine errors from r15, but four interface-totality blockers remain.

## Findings

### R17-F1 — D-4's “operator disposition” has no legal authority path, exact record, or total recovery semantics

The run-scoped predicate is now at the right gates: `turn_open` and `attempt_open` are withheld while a parked tool-unknown row exists (`2026-07-16-mvp-ipc-manifest-seam-contract.md:72`), and §D.2 is correctly only defense in depth (`:195`). The clearing half is not contract-real.

r16 says a “durable, operator-authorized record” reaches m-10 “via the operator's direct-route instruction to the addressed recipient or the governed gate, per §8b.” But ratified §8b says the direct route is **operator-to-one-agent only**, never generic app IPC, and is non-transitive: the direct instruction authorizes only that addressed recipient; a later citation is E0, not transferable operator authority (`STEP-3-ARCH-AMENDMENT.md:114-124`). m-10 is explicitly not a seat and this contract defines no live operator endpoint. An m-9 worker cannot convert its direct instruction into an operator-authorized m-10 state transition by forwarding it over CTRL-W. The alternative “governed gate” is also unnamed: no sanctioned grant type, producer, consumer, or verification rule is specified.

The durable half is equally open:

- `RUN_PARKED_UNKNOWN{row refs}` uses an open placeholder rather than exact reference fields;
- §F has no operator-disposition record/table, authorization reference, terminal disposed state, or transition key (`m-10 :225-244`);
- no idempotency, conflict, stale-target, multi-row, or crash disposition is pinned; and
- the old recovery prose still says a tool unknown may recover by a new attempt/new ticket while the new predicate forbids that until disposition (`:81,208`).

Required revision — choose one:

1. Use m-9's option (a): an exact m-10 state-only parked-unknown frame/turn-open state that reaches the worker/model before successor work, avoiding a new operator-authority ingress; or
2. If retaining option (b), first route the architecture question through Master/operator, then pin the exact authorized ingress/principal, message/grant, verification rule, durable record schema, target-row key, closed disposition enum, atomic transition, idempotency/conflict/crash behavior, and m-9 reply. Do not treat a §8b citation as transferred authority.

In either branch, replace `row refs` with a closed shape and sweep §B.3/§D.4/§F so the recovery text and stored states agree with the chosen gate.

### R17-F2 — D-5 fixes the inputs but still has no exact success/error reply family

r16 now consumes m-9's exact two request shapes and correctly makes `turn_cancel_ack` record-only while only `turn_terminal{terminal: turn_cancelled}` releases the lease (`m-10 :71`; m-9 r4 `:144-148`). But it then says the two **incoming request** frames are “both `re`-correlated.” Under this contract, `re` belongs to the reply that correlates to a request; it is required on reply-class messages, not on the request itself (`m-10 :27-41`).

No successful ack type/fields are named for either request; stale-epoch/unknown-turn are only called “typed-rejected” without an error frame; “same reply” for duplicate terminal refers to no defined reply; and duplicate/lost-reply behavior for `turn_cancel_ack` is absent. This leaves m-9's required “ack durable receipt / consume the ack” half non-executable.

Required revision: define either two exact success replies or one closed discriminated receipt reply, each `re`-correlated to the request; define exact stale-epoch and unknown-turn rejection shapes; make duplicate `turn_terminal` and duplicate `turn_cancel_ack` idempotent with the same durable reply; and pin the cancellation ordering/crash cuts so a recorded cancel ack cannot silently substitute for, lose, or race the terminal transaction.

### R17-F3 — D-2 says “bound by name,” but its values are not the m-7 producer names

The important safety mechanism closes: the worker reports actual acquisition, first admission waits for that report, suspended is bounded/retriable, tuple mismatch is immediate terminal, and a missing result times out (`m-10 :70`).

The exact binding does not. m-10 declares `result ∈ {attach_ok, attach_suspended, attach_tuple_mismatch}` while the current repaired m-7 producer taxonomy is bytewise `{attach-ok, broker:attach-suspended, broker:attach-tuple-mismatch}` (`m-7 r10 :214-220`, SHA-256 `da1ed802…`). Underscored wrapper values are not “consumed by name.” The m-10 bytes also still call r9 “under revision” after r10 was filed, and r10 has now received MUST-REVISE because its taxonomy is not total over the PREPARING suspension (`step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-193032.md`).

Required revision: after m-7's repaired D-3 bytes receive pair approval, either carry the producer tokens byte-exact in `attach_result.result` or define an explicit total one-to-one mapping from every approved broker result to a distinct m-10 wrapper value. Bind the final m-7 hash and re-run the three result/timeout reachability legs; do not call a nonmatching, unapproved vocabulary a by-name binding.

### R17-F4 — `attempt_open_reject` names two invariant failures but still does not name their supervision dispositions

The no-row half is materially repaired: `attempt_open_reject` is an exact reply-class shape with a closed reason set, no row is committed, and m-9 emits no DATA-P and charges no attempt budget (`m-10 :61`). `stale_epoch` also maps to the established fence/retirement machinery.

For `invalid_turn` and `invalid_lease`, however, the contract says only “the named fail-closed supervision disposition.” No disposition is actually named: it does not say whether the worker generation becomes FAILED, whether the active turn parks, whether §B.4 retirement runs, or what m-9 does after consuming the reply. Naming fixtures does not supply those state transitions.

Required revision: pin the exact m-9 action and m-10 worker/turn/lease transition for each non-epoch reason (or merge them into one exact internal-invariant-fault disposition if they are intentionally identical). Reuse or explicitly map the existing rejection vocabulary, and make each fixture assert the resulting durable state plus zero DATA-P/zero budget — not only the returned token.

## What closes

- R16-F1's core safety fact closes: first admission depends on observed worker capability acquisition, not broker-feed readiness.
- R16-F2's placement closes: D-4 is run-scoped at `turn_open`/`attempt_open`, and ticket issue is defense in depth.
- R16-F3's state ordering closes: the exact m-9 request fields are restored and cancel ack never terminalizes/releases.
- R16-F4's no-row success/rejection split, `attempt_open_ok` commit ordering, DATA-P gate, and budget rule close.
- The r14 `rejected_local` semantics and the other previously approved surfaces remain intact.

## Gate disposition

MUST-REVISE is byte-bound to `859cc7b69c982e892c87a21b97cb04113558d0c89cf5aa2736e1c50725271e21`. The r16 SITREP, m-8 final-byte review, m-9 r5 rebase/re-review, consumer rebind round, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. If the planner retains an operator-authorized D-4 clearing path rather than choosing the already-offered state-only branch, that topology/authority change must route upward before new m-10 bytes claim it.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `90f20959e4a1b12850c60a584abf40fa40b1b0c8a5ea973a3fdd20f057a9be61`.
- Exact m-10 r16 SHA-256 recomputed: `859cc7b69c982e892c87a21b97cb04113558d0c89cf5aa2736e1c50725271e21`.
- Ratified MVP amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified architecture amendment SHA-256 recomputed: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Current m-8 r6 SHA-256 recomputed: `ab63f6eb94c93dd4d62d2067fd174e1feddff5e6bf1a9e54d647c52f2718bc83`.
- Current m-9 r4 SHA-256 recomputed: `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Current m-7 r10 SHA-256 recomputed: `da1ed8029cfa20999894ab49ad19204f343c1281114cce682928177604322162` (D-3 bytes; pair verdict MUST-REVISE at `…-193032`).
- Incoming DESIGN relay exact-file lint: OK.
- Targeted seam sweep: m-10 `:27-41,61,68-81,180-208,225-244`; m-9 `:35-38,128-130,144-149,189-192`; m-7 r10 `:214-220`; ratified §8b `:114-124`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file mode exit 0.
Next requested action: m-10.planner folds R17-F1..F4 into fresh bytes, choosing the state-only D-4 branch or routing any new operator-authority ingress upward, binds D-2 to pair-approved m-7 producer bytes, makes D-5 and `attempt_open_reject` total at the reply/state-transition grain, recomputes the SHA-256, and requests a fresh uniquely-parented m-10.implementer review; do not route the r16 SITREP/rebind/final-review round.
