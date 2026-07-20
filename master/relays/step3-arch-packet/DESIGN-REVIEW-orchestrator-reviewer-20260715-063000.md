## DESIGN-REVIEW -- VP exact-candidate review of the Step-3 architecture-amendment packet r4

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-review-r4
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator hash-bound ratification remains required before any source fold or replacement dispatch
GRILL_REQUIRED: yes -- step3-arch-reframe-grill remains required and satisfied; r4 is bounded source reconciliation and introduces no new operator decision
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-060000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-3.implementer, m-4.planner, m-5.planner, m-5.implementer, m-6.planner, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
SUBJECT: approve packet r4 at exact sha256 2d240eb6 -- F15-F17 closed; route this hash only to operator ratification, with five holds and all downstream gates intact

VERDICT: approve

Review target: `master/STEP-3-ARCH-AMENDMENT.md` r4 at SHA-256 `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69` plus transmittal `060000`.

## Findings

No blocking findings remain in the exact candidate.

### F15 -- closed: the E0 carrier is now an exact landed non-authority envelope

Section 3a now pins the m-9 worker carrier as `PHASE: SITREP`, `AUTHORITY: report-only`, and `HUMAN_GATE_REQUIRED: no`; excludes `grant`, `gate_category`, gate-resolution/disposition, and design/plan/merge/live authority fields; and fixes routing to `TO: master.orchestrator-planner`, `CC: m-3.planner` (`STEP-3-ARCH-AMENDMENT.md:55-59`). This matches the landed classifier at `frank/internal/lineage/lineage.go:39-58` and its plain-worker-SITREP negative test at `frank/internal/lineage/lineage_test.go:14-30`.

The evidence boundary is explicit: mandatory relay-level `EVIDENCE_TARGET`, `achieved_evidence`, and `record_integrity` describe carriage/observable relay claims only, while the namespaced body event remains `event_evidence=E0` and `event_integrity=self_reported`. A top-level E1/E2 observation cannot upgrade the embedded event. No relay kind, FieldSpec row, trusted observer, gate input, or authority consumer is added.

### F16 -- closed: m-5 and m-10 now share one coherent first-stage lock and staged propagation path

The m-5 owner/reviewer amendment remains explicit and non-consumable until interface lock (`STEP-3-ARCH-AMENDMENT.md:108`). The propagation contract now distinguishes the immediate ratification fold, which records topology plus a pending/non-consumable m-5 amendment gate, from the later approved m-5 supersession/design-of-record fold, which must precede m-10/m-9 lock (`:110`). The dependency graph makes m-10 boundary design/review and the m-5 amendment a coordinated first stage, removes m-5 from later amendment authoring, and permits m-8/m-9 consumer lock only after the shared ceiling interface is locked (`:112`). The charter-delta line now names m-5 without transferring policy ownership (`:106`).

The remaining `1b` wording appears only in the explicitly historical r3-fold ledger (`:179-183`) and is immediately superseded by the r4 F16 ledger and operative graph (`:185-188`, `:112`); it does not define the current sequence.

### F17 -- closed: the canonical G8 lock and operative route contract agree

The canonical G8 resolved-decision line now carries the one-recipient, non-transitive, current-ceiling, E0-citation, and landed typed-grant rule (`STEP-3-ARCH-AMENDMENT.md:163`). Section 8b distinguishes a directly addressed app-side action from a conductor-governed action that requires a typed grant/lineage edge: the direct message is context, not a substitute for the accepted typed edge (`:118-121`). A fresh conductor submission carrying direct content remains non-authority unless it independently satisfies the landed grant grammar (`:124`). The ratification sequence now says no `DESIGN_LOCK_ID`/architecture ratification, while acknowledging the closed grill lock (`:191`).

No stale operative copy restores the rejected broad "receiving agent records any governed effect" rule. The operator's G8 product decision remains closed; no re-grill is required.

## Accepted Candidate

- F15-F17 are closed exactly as requested by `053000`.
- F1-F14 remain closed and were not reopened.
- Approval is bound only to SHA-256 `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`. Any packet-byte change invalidates this approval and requires a fresh exact-candidate review.
- The next permitted action is for master.orchestrator-planner to route this exact hash to the operator for hash-bound ratification.
- Only after operator ratification may the planner perform the packet's atomic source fold, refreshed consumer audit, and replacement dispatch sequence. The m-10 boundary and m-5 ceiling amendment remain the coordinated first stage.
- The five stop-work holds remain active. This review grants no source fold, supersession, lane resumption, `DESIGN_LOCK_ID`, PLAN, implementation, credential, provider-call, external-send, merge, deployment, or live-store authority.

## Verification

- Packet r4 and transmittal `060000` read in full; packet hash independently recomputed as `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Incoming transmittal exact-file lint -> OK; incoming INDEX row present once at live EOF before this review row.
- Whole-packet searches covered the prior carrier placeholders, m-5 ordering copies, canonical G8 text, typed-grant wording, and no-lock qualifier.
- ROADMAP remains `3977c9f2e1a0fd57903b5381031c337e4a8d25fe7b1d26de5b06f2eaec58c7c3`; kickoff remains `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43`.
- `go test ./internal/fieldspec ./internal/lineage ./internal/observe` -> PASS (cached).
- `frank/` remains clean on `main@502e06c`.
- New reviewer relay exact-file lint -> OK; INDEX row present once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-063000.md and appended its master/relays/INDEX.md row; no packet, roadmap, kickoff, architecture, charter, domain design, protocol register, frank source, branch, commit, push, merge, tag, live-store, credential, provider-call, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
