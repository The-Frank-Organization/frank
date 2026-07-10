## RECONCILE -- VP pressure-test: revise bucketing; collapse claims, not confused-agent interface mechanisms

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: design-review
PARENT_DISPATCH_ID: design-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner -- I concur with the operator's deployment-fork decision, but I do **not** approve the bucketing as currently
phrased. Verdict: **REVISE.**

The useful distinction is: adversarial **security claims** collapse; confused-agent **interface mechanisms** do not.

## What I Concur With

1. **D1-D5 stand.** The operator has made the threat-model/risk-appetite decision: Step-1 is attach, confused-not-
   adversarial, interface-level guardrail, and malicious code-executing agents are explicitly out of scope
   (`master/GRILL-LOCK-deployment-fork-2026-07-01.md:30-49`). I do not reopen wrap.

2. **The global claim sweep is correct.** The old structural/by-construction/sole-writer/unbypassable claims must be
   globally relabeled to the D4 claim-set. Under the selected threat model, those no longer need to be Step-1
   architecture guarantees against a malicious lane.

3. **The remaining Step-1 NO-GO items are still real.** Conductor-core, serialization/crash recovery, the phase-split
   dead-end, pure-judgment A-floor, and decision-② authority-class fail-closed all remain before Step-1 PLAN
   (`master/DESIGN-REVIEW-2026-07-01.md:22-24`, `master/READINESS-REGISTER.md:381-385`).

## Required Revisions

1. **Do not collapse the interface guardrail into accepted-risk prose.** D2b is not just a claim sweep: it says seats act
   only through `submit()`/`project()`/`read()`, raw store/config paths are not in the seat tool surface, and policy config
   is conductor-owned/loaded at trusted startup (`master/GRILL-LOCK-deployment-fork-2026-07-01.md:36-39`). That must be a
   conductor-core design requirement and acceptance fixture. If a Step-1 attached seat still has a general shell or file
   tool with access to store/config paths, then config-lane-writable, store-write, and sole-external-sender are **not**
   adversarial-only; a confused agent can follow a bad instruction around the gate.

2. **Do not bucket fill-time authority as adversarial-only.** Fill-time authority is the confused-agent product
   mechanism: forbidden choices are absent from the rendered form, and the courier validates constrained picks. The
   adversarial version of that claim collapses, but the form/interface mechanism remains Step-1 core. Downstream text
   should say "collapse the by-construction security claim, retain fill-time authority as the confusion-resistant
   interface contract."

3. **Resolve the contradiction inside `DESIGN-REVIEW-2026-07-01.md`.** The update says the attach-vs-wrap FATALs collapse
   to claim sweep + accepted-risks and "not architecture changes" (`:15-24`), but §2 still lists config integrity as a
   must-fix architecture mechanism (`:113-118`) and the re-baseline path still includes config-integrity mechanism
   (`:203-207`). I think the right resolution is: remove config-integrity as an adversarial isolation redesign, but keep
   **trusted config load + not-in-seat-tool-surface** inside conductor-core.

4. **R2 and altitude-B are not mere hygiene if Step-1 ships their schema.** I accept that they need not block the
   conductor-core-only slice under the confused-threat model, but m-2/m-4 schema work must not ship with R2
   under-specified or altitude-B record-grain wrong. Phrase as: "not a conductor-core Step-1 blocker unless the Step-1
   slice includes the routing FieldSpec / routing record; otherwise before the routing/schema build step." The current
   "cheap and correctness-improving" language underplays it.

5. **Keep away-token replay/restart binding as later-step MUST, not disappeared.** It is not primarily malicious-only:
   stale approval and sibling-token reuse can bite normal operator flows. It can be before park/wake / away bridge rather
   than before conductor-core Step-1, but it should remain a hard gate for that build step.

6. **Fix the grill-lock path spelling before downstream dispatch text.** The real file is
   `master/GRILL-LOCK-deployment-fork-2026-07-01.md`; the planner relay/register shorthand omits the hyphenated date in
   places. Not a blocker, but avoid ambiguous artifact references in the next dispatch.

## Safe Bucketing

- **Collapse to claim sweep + accepted-risk:** adversarial DI-1/DI-2 bypass guarantees, malicious direct store append,
  malicious operator-FROM forgery, malicious config edit, malicious external-send bypass, "by-construction" ladder.
- **Retain in conductor-core Step-1:** interface-only seat tool surface, raw store/config path exclusion from seat tools,
  trusted config load, fill-time authority/form rendering, local outbox-only external-send posture, serialized commit +
  crash recovery + internal-fault disposition, phase-split required-set, pure-judgment A-floor, authority-class
  fail-closed.
- **Retain before their build step:** R2 opt-in referenceability + row schema, altitude-B per-row deviation grain,
  away-token sibling burn/restart binding, away egress/ODB collisions.

Not authorized / not claimed: no Step-1 PLAN, no code/pcode/spike, no pair dispatch, no domain-design edit, and no
operator decision reopened. This is a correction to the re-baseline task text.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/design-review/RECONCILE-orchestrator-planner-20260701-143917.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/design-review/RECONCILE-orchestrator-reviewer-20260701-144217.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/design-review` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read-only review of `master/relays/design-review/RECONCILE-orchestrator-planner-20260701-143917.md`, `master/GRILL-LOCK-deployment-fork-2026-07-01.md`, `master/DESIGN-REVIEW-2026-07-01.md`, `master/READINESS-REGISTER.md`, `ROADMAP.md`, and the prior VP relay; wrote this reviewer relay and appended `master/relays/INDEX.md`; no domain-design edits, no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: pcode clean; cwd status unavailable (not a git repo)
Next requested action: planner revises the re-baseline task text so collapsed items are limited to adversarial security claims, while D2b interface/tool-surface constraints remain conductor-core Step-1 requirements.
