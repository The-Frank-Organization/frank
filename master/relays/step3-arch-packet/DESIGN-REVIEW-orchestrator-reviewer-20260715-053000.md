## DESIGN-REVIEW -- VP exact-candidate review of the Step-3 architecture-amendment packet r3

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-review-r3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- only a revised VP-approved exact candidate may proceed to operator hash-bound ratification
GRILL_REQUIRED: yes -- step3-arch-reframe-grill remains required and satisfied; the bounded corrections below do not reopen an operator decision
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-050000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-3.implementer, m-4.planner, m-5.planner, m-5.implementer, m-6.planner, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
SUBJECT: must-revise at packet sha256 8a6154e3 -- pin the exact E0 carrier envelope, reconcile m-5 amendment ordering and propagation, and fold the non-transitive G8 rule into the resolved lock text

VERDICT: revise

Review target: `master/STEP-3-ARCH-AMENDMENT.md` r3 at SHA-256 `8a6154e38e43a9fc08945c69ef69d0f55344d65f2ffe1b7e53a7b621ca95a046` plus transmittal `050000`.

R3 preserves the accepted topology and materially implements the intended F11-F14 repairs. The direct route is now explicitly non-transitive and ceiling-bounded; m-5 has an owner/reviewer amendment with a pinned interface; the E0 event is separated from conductor evidence; and the hash/status/action references are honest.

Three exact-candidate gaps remain. They are bounded source reconciliation, not new product forks, and do not require another operator grill.

## Findings

### F15 -- "ordinary existing non-authority relay" is still a classification, not the exact carrier shape required by F13

Section 3a selects the body of an ordinary non-authority relay and names body non-gate-referenceability, no `grant`, no gate resolution, and the downstream readers (`STEP-3-ARCH-AMENDMENT.md:55-59`). That closes the provenance branch, but not the exact relay envelope requested by `043000` F13.

The landed authority classifier does not infer non-authority from the word "ordinary." It returns authority-bearing for any grant, human gate, A-category gate, `PLAN`/`IMPL`/`REVIEW-FOLD`/`MERGE-GATE`/`LIVE-VERIFY` phase, implementation/merge/live/fold authority, or non-SITREP/non-RECONCILE orchestrator-planner record (`frank/internal/lineage/lineage.go:39-58`). A plain pair/worker `SITREP` is the tested non-authority shape (`frank/internal/lineage/lineage_test.go:14-30`). Without pinning those headers, a downstream design can satisfy the prose while accidentally selecting an authority-bearing envelope.

Required fold:

- name the carrier as the m-9 worker seat's existing `PHASE: SITREP`, `AUTHORITY: report-only`, `HUMAN_GATE_REQUIRED: no` relay, with no `grant`, no `gate_category`, no gate-resolution/disposition field, and no design/plan/merge/live authority field;
- name exact routing: recommended `TO: master.orchestrator-planner`, `CC: m-3.planner` plus the eventual audit reviewer if separately seated;
- state that mandatory top-level `EVIDENCE_TARGET` and conductor-produced `achieved_evidence` concern carriage/relay claims only. The namespaced body event separately carries explicit `event_evidence=E0` and `event_integrity=self_reported` (exact body-schema names remain m-3's downstream design), so an E1/E2 top-level relay observation cannot upgrade the event;
- retain the no-new-FieldSpec/no-new-record-kind choice and the prohibition on gate/authority consumption.

That is the minimum exact carrier contract. It does not require a conductor change.

### F16 -- the m-5 amendment is simultaneously interface-locked at step 1b and first authored at step 3, and its design-of-record propagation is still absent

Section 8 correctly says the m-5 planner authors, the implementer reviews, and the interface locks before m-10/m-9 consumer lock (`STEP-3-ARCH-AMENDMENT.md:108`). The dependency graph then says the amendment **interface-locks at 1b**, but lists the same amendment among work that **authors at step 3** (`:112`). Both cannot be the operative sequence. The graph also says m-10 boundary design/review lands first even though m-10 consumes the m-5 interface that must precede its consumer lock.

The source propagation row still lists the m-5 charter but not the review-driven m-5 design amendment/supersession required by `043000` F12 (`:110`). The `Charter deltas` line itself omits m-5 despite the transmittal claiming it was added (`:106`). That would leave the ratification fold and locked m-5 design disagreeing without an explicit staged lineage.

Required fold:

1. Make m-10 boundary design and the m-5 ceiling-host amendment a coordinated first stage: author/review both, interface-lock their shared ceiling contract, then permit m-10 and m-9 consumer locks. Remove m-5 from the later step-3 authoring list, or change step 1b to "begins" and place its one actual lock point before consumers.
2. Add the m-5 delta explicitly to the charter-delta line.
3. Make propagation timing honest: the immediate ratification fold may record the new topology plus the **pending/non-consumable m-5 amendment gate**, but it must not silently rewrite the locked m-5 design. The replacement flow then creates the m-5 planner/implementer amendment and records its approved supersession/design-of-record fold before m-10/m-9 lock. Name both stages in the propagation set.

No operator decision is needed; this is sequencing and design-lock lineage for the already-selected host move.

### F17 -- the durable G8 resolved-decision line still preserves the pre-F11 broad wording

Section 8b now has the correct bounded rule (`STEP-3-ARCH-AMENDMENT.md:114-123`), and the post-grill r3 fold records it at `:178-182`. But the `GRILL_LOCK`'s canonical **Resolved decisions** entry still says a receiving agent records "any governed effect" under its own `FROM`, with none of F11's non-transitivity, ceiling, E0, or typed-grant qualifications (`:162`). A downstream reader following the lock rather than its later appendix can recover the rejected broad interpretation.

Required fold:

- replace the G8 resolved-decision entry with the r3 non-transitive wording; this is a bounded clarification of the existing operator answer, not a re-vote;
- sharpen section 8b item 3: direct instruction alone may authorize a directly addressed app-side action within the current ceiling, but when a conductor-governed action requires a typed grant/lineage edge, the sanctioned typed-grant branch is mandatory. A direct message is context, not a substitute for the required accepted grant;
- clarify item 7 that a fresh conductor submission can carry direct content but remains non-authority unless it independently satisfies the landed grant grammar;
- qualify the final line's "No lock" as no `DESIGN_LOCK_ID`/architecture ratification, since the grill lock exists.

This preserves G8's authority-bearing direct channel and no-forced-operator-relay decision while keeping conductor governance mechanically honest.

## Accepted R3 Folds

- F11: one-recipient/non-transitive authority, E0 agent citation, current-ceiling bound, landed grantor set, and live-ingress-only "by construction" scope are accepted subject only to F17's stale lock copy and typed-edge clarification.
- F12: m-5 sole policy ownership, m-10 enforcement hosting, owner/reviewer amendment, immutable run/worker binding, and fail-closed absent/stale behavior are accepted subject only to F16 sequencing/propagation.
- F13: body-carried app schema, separate relay-versus-event evidence semantics, named readers, and no typed provenance amendment are accepted subject only to F15's exact envelope.
- F14: exact transmittal, required-and-satisfied grill status, candidate hash, and qualified action record are closed.
- F1-F10 remain closed as recorded in `043000`; do not reopen them.

## Required Revision Sequence

1. Fold F15-F17 only and reconcile all repeated copies, especially the canonical G8 resolved line and dependency graph.
2. Recompute the packet SHA-256 after the exact envelope, graph, propagation, and final status text are stable.
3. Return the new exact candidate for VP review. No operator grill rerun is required unless a new grantor, ceiling semantic, or typed conductor member is introduced.

The five holds remain in force. No operator ratification, supersession, source fold, replacement dispatch, lane resumption, design lock, PLAN, code, credential, provider call, external send, merge, or deployment is authorized by this review.

## Verification

- Packet r3 and transmittal `050000` read in full; review bound to SHA-256 `8a6154e38e43a9fc08945c69ef69d0f55344d65f2ffe1b7e53a7b621ca95a046`.
- Incoming transmittal exact-file lint -> OK; INDEX row present once.
- ROADMAP remains `3977c9f2e1a0fd57903b5381031c337e4a8d25fe7b1d26de5b06f2eaec58c7c3`; kickoff remains `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43`.
- Landed carrier authority behavior checked at `lineage.go:39-58`; plain pair/worker SITREP non-authority behavior checked at `lineage_test.go:14-30`.
- `go test ./internal/fieldspec ./internal/lineage ./internal/observe` -> PASS (cached).
- `frank/` remains clean on `main@502e06c`.
- New reviewer relay exact-file lint -> OK; INDEX row survival check -> reviewer row present once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-053000.md and appended its master/relays/INDEX.md row; no packet, roadmap, kickoff, architecture, charter, domain design, protocol register, frank source, branch, commit, push, merge, tag, live-store, credential, provider-call, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
