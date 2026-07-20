## RECONCILE -- REVISE-NARROW: stage-1 dispatches need one grill, one F65 owner edge, and three consumer-confirmation corrections

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage1-dispatch-review-r1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- no new operator choice is required unless the m-7 grill would change a ratified topology or claim boundary
GRILL_REQUIRED: yes -- the m-7 dispatch leaves broker process placement open at large ceremony with no later m-7 grill
DESIGN_DOC_ID: step3-mvp-stage1-dispatch-bundle
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW the five stage-1 dispatches -- m-1 is sound; supplement m-2/m-3/m-7/m-10 before final-byte closure or interface lock

VERDICT: revise

Review target: the F44 source-fold transmittal `041610` and the latest five planner-authored relays at review start:

1. `master/relays/step3-mvp-design-m2/DESIGN-orchestrator-planner-20260716-041620.md`
2. `master/relays/step3-mvp-design-m7/DESIGN-orchestrator-planner-20260716-041630.md`
3. `master/relays/step3-mvp-design-m10/DESIGN-orchestrator-planner-20260716-041640.md`
4. `master/relays/step3-mvp-design-m1/DESIGN-orchestrator-planner-20260716-041650.md`
5. `master/relays/step3-mvp-design-m3/DESIGN-orchestrator-planner-20260716-041700.md`

## Findings

### F67 -- BLOCKER: m-7's unresolved broker placement requires a durable grill before final DESIGN review

The m-7 dispatch says `GRILL_REQUIRED: no` (`041630:12`) while leaving **own process vs protected thread/module in the app-main process** to the owner DESIGN (`:26`). That is new, still-open work at large ceremony; it is a cross-domain and hard-to-reverse boundary for credential custody, crash isolation, broker restart ownership, m-10 separation, epoch fencing, push routing, and IPC/backpressure. The amendment-level topology grill constrained the broker only to live outside the replaceable worker generation; it did not choose this placement. No later m-7 design stage or grill exists.

Required correction: supplement `step3-mvp-design-m7` with `GRILL_REQUIRED: yes`. Before m-7.implementer reviews final bytes, fold a durable `GRILL_LOCK` comparing at least the two named placements against the ratified F57 claim boundary, m-1 credential semantics, m-10's no-conductor-verb rail, process lifecycle/recovery, epoch-change linearization, in-flight disposition, push delivery, and failure isolation. This remains an m-7 design choice unless the selected answer would alter a ratified topology or claim boundary; only that latter case returns to the operator.

### F68 -- BLOCKER: F65's conductor build/config identity has no dispatched producer contract

The m-7 dispatch merely says the F65 note "stands" (`041630:31`), and the m-3 dispatch repeats that conductor identity is bound separately (`041700:25`) while assigning m-3 only the app/provider E3 schema and evaluator. Neither dispatch assigns the canonical bytes that make **conductor service build digest + governing config identity** comparable to the running relay-exchange leg. The ratified exit test requires named values in the same exit-test record; prose that they are "separate" does not define their producer, canonical encoding, current-process binding, or evidence reference.

Required correction: assign m-7, as conductor lifecycle/config host, the canonical conductor-identity producer contract: exact build-artifact identity, governing-config identity, canonical encoding, how the running service proves those values are loaded, and the relay-leg evidence reference carried into the exit-test record. m-3 must consume and confirm the scope boundary so its app/provider evaluator neither absorbs nor omits the separate relay-leg binding. Master+VP retain ownership of the final composite exit-record join. Add m-3 as an explicit m-7 consumer/CC and m-7 as an explicit m-3 upstream/CC. This realizes r7 F65; it does not reopen it or add a conductor protocol/store field.

### F69 -- BLOCKER: three direct producer-consumer confirmations are omitted from the body contracts

The dispatch headers expose some peers in `CC`, but each body narrows the consumers required to confirm, so visibility is not confirmation authority:

- **m-2 -> m-10:** m-2 produces the relay schema digests and mapping version used in the locked tool-identity vector (`041620:24`); m-10 hosts exact-set identity verification and emits the manifest digest (`041640:24`). Add m-10 to the m-2 consumer-confirmation set (`041620:28`).
- **m-7 -> m-2:** m-7 owns the shared transport half beside the m-2 mapping module and the parity boundary, but omits m-2 from its consumer set (`041630:31`). Add m-2 confirmation so the transport contract cannot absorb or strand the mapping layer.
- **m-10 -> m-3:** m-10 produces `run_manifest_digest` (`041640:24`), which m-3 consumes in the external-E3 tuple (`041700:25`). Add m-3 to the m-10 `CC` and consumer-confirmation set (`041640:17,30`). The reciprocal direction is already present because the m-3 dispatch names m-10 as a consumer (`041700:29`).

These are supplemental routing corrections. They do not change the ratified eight-name policy, tool identity shape, evidence scope, topology, or domain ownership.

## Disposition

- **m-1 dispatch:** approve as issued. Its secret-boundary and logical-seat semantics are owner-correct, require final-byte review by m-1.implementer, and route the resulting contract to m-7/m-8/m-9/m-10 consumers.
- **m-2, m-3, m-7, m-10 dispatches:** revise-narrow through append-only supplemental planner relays; do not edit the historical dispatch files.
- The F44 governing-source fold is internally consistent with exact r7 and its ordered manifest reproduces. No source-fold rollback, amendment rewrite, fresh amendment hash, or operator re-ratification is required if F67-F69 stay inside the boundaries above.
- Existing DESIGN-only authoring may continue, but no affected owner may claim final-byte pair closure, complete consumer confirmation, `DESIGN_LOCK_ID`, or readiness for the Master+VP interface lock until its correction is consumed. In particular, m-7 final review waits for the grill lock.

## Required Return

1. Issue direct supplemental planner relays to m-2, m-3, m-7, and m-10 carrying F67-F69 without rewriting the five historical relays.
2. Require the m-7 `GRILL_LOCK` in the final m-7 DESIGN record.
3. At stage-1 close, return a producer -> consumer confirmation table that includes m-2 -> m-10, m-7 -> m-2, m-10 -> m-3, and the m-7 -> m-3 F65 identity edge, alongside the already-routed confirmations.
4. Keep stage 2 and later consumers from treating an uncorrected stage-1 artifact as interface-locked. The Master+VP stage-6 join remains the only interface-lock event.

No PLAN, T4 code token, implementation, credential provisioning, provider call, release-binding execution, external send, merge, deployment, or live-store mutation is authorized by this review.

## Verification

- Exact five-relay file lints each end `OK`; root-mode output still carries the known historical `INDEX.md` and lineage debt, separate from these files.
- Ordered post-fold 15-file governing manifest independently recomputed: `d072e85e13554bc5739d8aaeaa8b3885971166ec70b39adcada35651485a0ffd`.
- Operative MVP amendment remains exact `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`; reframe packet remains exact `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 Step-4 basis remains exact `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- `frank/` was not changed by this review; porcelain-v2 is clean on `main@502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, tracking `origin/main` at `+0/-0`.
- New reviewer relay exact-file lint ends `OK`; its INDEX row is present exactly once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-043205.md` and appended its `master/relays/INDEX.md` row; no governing source, historical relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean on `main@502e06cc07b5` at `+0/-0` versus `origin/main`.
Next requested action: issue append-only supplemental dispatches closing F67-F69, then route the final owner bytes through pair review and the corrected consumer-confirmation graph before the Master+VP interface lock.
