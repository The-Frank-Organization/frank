## DESIGN-REVIEW - m-7 adversarial review of typed broker attach results r10

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r10
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the finding completes the existing broker suspension taxonomy without changing topology, custody, or policy
GRILL_REQUIRED: no
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260717-192414.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-193032.md
SUBJECT: must-revise - r9's three-cause defect is closed, but PREPARING is a fourth reachable suspension with no total attach result or legal event reason

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed the fresh r10 contract bytes at SHA-256 `da1ed8029cfa20999894ab49ad19204f343c1281114cce682928177604322162`, uniquely parented from the r9 review.

R9-F1 is correctly closed at every requested locus: the three named section 2.4 causes now precede tuple evaluation; equal/mismatch are restricted to live, actively-authorizing state; retained-state control loss is typed truthfully; FX-TB-19 now has a licensed same-worker suspended-then-success cut; and malformed attach frames remain distinct from malformed epoch updates. The typed split and no-custody-move result remain sound.

One adjacent suspension state makes the exact-three-result contract non-total. This review grants no owner approval, consumer confirmation/rebind, interface lock, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, release binding, merge, or deploy.

## Finding

### R10-F1 - PREPARING retains E while suspending all arrivals, but attach has neither a total result nor a recordable reason

Section 2.5 defines PREPARING as a reachable broker barrier after E+1 is proposed and before the crossing-set commit ack. During that interval:

- the broker has retained installed E state;
- the verified control session is live;
- E is no longer authorizing;
- "anything arriving is rejected typed `broker:suspended`"; and
- the broker may remain PREPARING-suspended across a lost ack/control outage (`2026-07-16-step3-mvp-transport-broker.md:129-142`).

An already-assigned E worker can lose its connection and re-present its otherwise-current tuple during this interval. R10's attach taxonomy says it returns exactly one result, but its first branch covers only the three section 2.4 causes `{no installed state, control lost, malformed update}` (`:214-216`). Its success and mismatch branches both require an actively-authorizing installed state (`:217-218`), which PREPARING intentionally lacks.

Therefore either:

- PREPARING falls outside the first branch and no result applies, violating the exact-one claim; or
- PREPARING is implicitly `attach-suspended`, but the closed attach event cannot encode it because the required suspended-reason enum is only `{no-installed-state, control-lost, malformed-update}` (`:245`).

The negative is operationally reachable under the licensed flow: the E worker was assigned before the transition; the broker enters PREPARING while retaining E; that worker reconnects before E+1 installs. This is not the impossible pre-assign 5b cut from r9.

Required revision:

1. Define `broker:attach-suspended` over the broker's full attach-blocking suspension predicate, including the section 2.5 PREPARING barrier, before tuple evaluation or capability mint.
2. Add an exact attach-event reason for that state, e.g. `preparing`, while preserving the existing three reasons.
3. Keep `attach-ok` and `broker:attach-tuple-mismatch` restricted to a live control session with no active suspension barrier and an actively-authorizing installed state.
4. Extend FX-TB-19 with the reachable cut: an assigned E worker re-presents during PREPARING and gets `attach-suspended{reason: preparing}` with no capability; after E+1 installs, another presentation of the old E tuple gets terminal `attach-tuple-mismatch`. This proves bounded transient handling can resolve into the terminal branch when authority advances.
5. Sweep any other explicit broker suspension/barrier state that can coexist with a retained installed tuple and either include it in the attach-suspended predicate/reason enum or prove it cannot receive worker attach.

## Accepted portions

- **R9-F1 closes.** All five requested repairs are present and internally consistent for no-state, control-loss, and malformed-update suspension.
- **The D-3 behavioral split remains correct.** Suspension licenses bounded hold/retry under supervision; tuple mismatch is terminal for the presenting generation and licenses no retry.
- **The event shape is correct in form.** `outcome=suspended` plus a required exact reason is the right closed-family representation; it needs the PREPARING member for totality.
- **No fourth behavioral result is required.** PREPARING has the same bounded-hold behavior as the existing suspended result; the missing item is predicate/reason coverage, not a new worker disposition.
- **No custody move or m-1 rebind is required.** The correction changes no credential, capability binding, lifecycle owner, section 2.7 custody surface, or conductor protocol/store shape.
- **Every r8-approved decision remains accepted.** Controller proof, transition recovery, F70/F73, transport, recording, F65/F68, m-1 binding, placement grill, and the no-conductor-protocol/store-change result are untouched.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Make the attach taxonomy total over PREPARING with suspension precedence and no mint.
2. Record PREPARING through a legal required suspended reason.
3. Prove the assigned-E reattach-during-PREPARING cut and its post-install transition to terminal mismatch.
4. Preserve every accepted r10/r9/r8 decision and the no-custody-move result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. m-9's rebase, consumer rebinds, and the Master+VP interface lock remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `da1ed8029cfa20999894ab49ad19204f343c1281114cce682928177604322162`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260717-192414.md` lints OK; routing, lineage, `DESIGN_DOC_ID`, scope, and superseded-r9 posture match.
- Re-read the D-3 routing, section 2.4's three-cause floor, section 2.5's PREPARING barrier and outage persistence, section 2.10's exact-three attach taxonomy, section 2.11's closed event reasons, and FX-TB-19.
- Current m-9 owner bytes remain `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`; current m-10 r16 bytes are `859cc7b69c982e892c87a21b97cb04113558d0c89cf5aa2736e1c50725271e21`. Their pending consumer/review work does not repair this broker-side totality hole.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0.
Next requested action: m-7.planner folds only R10-F1, preserves the repaired three-cause taxonomy and no-custody-move result, re-hashes, and sends fresh uniquely-parented bytes; m-9 rebase and rebinds remain held.
