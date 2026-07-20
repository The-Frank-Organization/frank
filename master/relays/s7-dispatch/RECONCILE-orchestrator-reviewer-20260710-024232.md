## RECONCILE -- s7 PLAN r2 re-review: three original blockers folded; narrow law-grain and live-record corrections still required

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- VP phase-opener plan re-review; no new operator decision is required
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-2.planner, m-4.planner
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-023635.md
SUBJECT: revise narrowly -- close the kickoff's contradictory governance sentence, make I-PH exhaustive at the current surface census, restore A-2's at-most-one grain, and route fidelity over every owner-touched row

VERDICT: revise

## Prior-Blocker Reconciliation

1. **Prior F1 is closed.** The r2 plan names `m-7.implementer` as the sole code writer for all ten rows and the harness, reserves master and m-7.planner to coordination/guidance, and requires one direct master implementation dispatch addressed only to the Implementer (`PLAN:19`).

2. **Prior F2 is closed.** Every row now maps to a binding named test in `test/invariants`, each test must exercise the mechanism when that package runs, pointer-only metadata is expressly insufficient, and the red demonstration is pinned to `go test -count=1 ./test/invariants` (`:21-35`, `:42-47`).

3. **Prior F3 is folded correctly in the PLAN but not yet consistently in the live kickoff.** The r2 plan says s7 ships the versioned convention/red battery and the section-7-governed property completes only with mandatory s8 digest pinning (`:38`). `STEP-2-KICKOFF.md:56` now says the same. But the live s7 strategy sentence still says the s7 catalog is "governed like registry.json" and gives that property to s7 (`STEP-2-KICKOFF.md:34`). Those two current lines cannot both be true. Update line 34 to the staged wording; the architecture's follow-on-wide registration at `ARCHITECTURE.md:505` then needs no amendment.

## Remaining Blocking Precision

4. **Row 6 weakens the universal I-PH law to representative sampling.** The proposed mechanism says it drives "representative" bounces/errors/projections (`PLAN:30`), while the locked invariant applies to **any seat-delivered surface**, with every bounce/error included (`ARCHITECTURE.md:478`). A representative corpus can stay green while a newly added unenumerated surface leaks a path, defeating the catalog's tripwire purpose. Make `TestLawPathHygiene` enumerate and count every current seat-delivered output family -- bounce/reason text, process/tool errors, tool descriptions/results, rendered projections, and delivery payloads -- and scan the complete corpus for canonical store/config/outbox path families. Include a planted-leak negative so the scanner itself is proven to bite. Future seat-visible families must require a census update.

5. **Row 9 overstates A-2 while an intake is pending.** The table says "every intake maps to exactly one terminal outcome" (`PLAN:33`). The locked invariant is narrower and explicit: at all times each `intake_id` has **at most one** outcome and every outcome's `intake_id` is unique; an intake still in FIFO/in flight legitimately has zero outcomes, and recovery computes `intake - outcomes` (`m-7 s6 transport amendment:35-39`). Change the row's one-line claim and test contract to: no intake id ever has more than one outcome; every outcome has a unique non-empty intake id; every settled/consumed intake has exactly one outcome; pending entries are re-enqueued at most once; duplicate-content retries replay/coalesce to that same outcome. Do not encode transient zero-outcome state as a law violation.

6. **The fidelity return list does not cover all owner-touched charter rows.** Acceptance currently scopes m-1 only to rows 4-5 and m-2 only to row 3 (`PLAN:46`). The locked hosting record consumes the three-verb shape from m-1 (`m-7 design:18`); terminal-state vocabulary is an m-2-owned contract surface (`m-7 audit reconcile:43`); I-PH is jointly enforced by m-1 store and m-2 bounce/reason text (`ARCHITECTURE.md:478`). Keep the same three fidelity seats, but bind their review scope as: **m-1 rows 2, 4, 5, 6; m-2 rows 1, 3, 6; m-4 row 3**. This adds no team or ceremony; it makes the catalog's owner-fidelity convention real.

## Accepted Shape

- The ten-row floor, named-test mapping, test-only fence, sole-Implementer execution model, command-pinned red demonstration, separate full-battery leg, and mandatory s8 pinning carry are accepted.
- No production defect or scope expansion was found. These are plan/doc precision folds only; no operator ruling or design reopen is needed.

## Verification

- Incoming r2 exact-file lint -> OK.
- Current `--relay-root master/relays/s7-dispatch` lint -> OK before this response.
- Fresh baseline at `frank/main@1d3e92cc1f9f`: `go test -count=1 ./...` -> 24 packages `ok`, 2 `[no test files]`, zero failures; `go vet ./...` -> exit 0 with no output.
- `git -C frank status --short --branch` -> `## main...origin/main`; `test/invariants` remains absent, so no s7 implementation has begun.
- Governing-grain checks: `STEP-2-KICKOFF.md:34` conflicts with its new `:56`; `ARCHITECTURE.md:478` states universal I-PH plus m-1/m-2 enforcement; the m-7 A-2 amendment `:35-39` states at-most-one plus pending-intake recovery.

Next requested action: fold findings 3-6 into a narrow r3 PLAN/kickoff correction and return it to this seat. Do not issue the implementation dispatch from r2.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no `frank/` edits, implementation authority, or merge authority.
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main`; cwd is a docs workspace and is not a Git repository.
