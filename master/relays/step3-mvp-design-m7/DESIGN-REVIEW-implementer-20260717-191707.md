## DESIGN-REVIEW - m-7 adversarial review of typed broker attach results r9

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r9
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the finding is a bounded attach-predicate correction inside the routed D-3 owner fold
GRILL_REQUIRED: no
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260717-190616.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-191707.md
SUBJECT: must-revise - attach-suspended must cover the full §2.4 suspended state, and the claimed 5b same-worker reachability contradicts assign-after-install

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed the fresh r9 contract bytes at SHA-256 `ed66e03892015e44e6f9ebb3d6eb514520b8c9a2f3533db5f17a0d9e5e948c69`, uniquely parented from the D-3 owner routing. The r8 approval is correctly void at these bytes.

The transient-versus-fenced split is the right interface direction, tuple mismatch is correctly terminal for the presenting generation, malformed attach frames correctly remain protocol errors, and the custody surfaces do not move. One predicate error makes the taxonomy incomplete at the exact fail-closed cut it must type.

This review grants no owner approval, consumer confirmation/rebind, interface lock, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, release binding, merge, or deploy.

## Finding

### R9-F1 - `attach-ok` can mint during control-session suspension, while FX-TB-19's transient leg is not worker-reachable

Section 2.4 defines the broker's suspended floor as the OR of three conditions: no installed state, control session lost, or malformed epoch update (`2026-07-16-step3-mvp-transport-broker.md:114-121`). The lifecycle matrix explicitly preserves the broker and its epoch state across app-main/control loss while making every capability useless in effect (`:161-168`).

R9 instead defines:

- `attach-ok` whenever the tuple equals the installed state; and
- `broker:attach-suspended` only when **no installed state exists**, while parenthetically listing control loss as though it necessarily erased that state (`:214-218`).

Control loss does not clear the installed state in the approved recovery model; it makes that retained state non-authorizing until a verified controller re-establishes the feed. A current tuple presented during that window therefore satisfies r9's `attach-ok` predicate and can mint a new capability even though §2.4 requires suspension. This is an authority regression, not merely an imprecise error label.

The claimed H-14 reachability leg is also contradictory. FX-TB-19 says a worker reaches the 5b no-state floor and later the same worker retries successfully. But the same section pins `publish -> broker installs -> assign`, and the worker learns `{generation_id, turn_epoch, broker_worker_endpoint}` from `assign` only after install. A legitimate successor cannot present during 5b; an old-generation worker can present, but after E+1 installs its old tuple is terminal mismatch rather than same-worker success. The fixture therefore does not prove the transient result is reachable by the licensed worker flow.

Required revision:

1. Give the broker's **global §2.4 suspended state precedence** over tuple evaluation. If the broker is suspended for any reason - no installed state, control-session loss, or malformed epoch update - attach returns `broker:attach-suspended`, mints no capability, and does not assign fencing meaning to the tuple.
2. Evaluate tuple equality only when a verified control session is live and an install-eligible current state is actively authorizing. Equal then returns `attach-ok`; unequal returns terminal `broker:attach-tuple-mismatch`.
3. Make the closed attach event truthful for all suspended causes. Replace `suspended-no-state` with a generic `suspended` outcome, or add an exact required reason enum such as `{no-installed-state, control-lost, malformed-update}`; do not record control-loss suspension as no-state.
4. Replace or expand FX-TB-19 with a genuinely worker-reachable transient cut: retain an installed state, suspend via control-session loss or malformed update, present the otherwise-current assigned tuple and prove `attach-suspended`/no capability; restore verified state and prove the same tuple can then return `attach-ok`. Retain the stale-generation terminal/no-hammering leg.
5. Keep malformed **attach frames** in the protocol-error class and distinguish them from a previously malformed epoch update that has put the broker itself on the suspended floor.

## Accepted portions

- **The D-3 split is necessary.** Transient broker suspension and installed-state tuple mismatch require opposite worker behavior and must be distinct typed results.
- **Tuple mismatch semantics pass.** Once the broker is actively authorized by a live control session and current installed state, inequality means the presenting generation is fenced; no retry is licensed and repeated presentation changes no state.
- **No additional fencing member is needed after suspension precedence is fixed.** Recognition mismatch remains a control-surface concern; malformed attach input remains a protocol error.
- **Recording in the existing attach event is the right ownership shape.** The event remains uncoupled and inside the closed m-10-written family; only its suspended outcome/reason needs truthful completion.
- **All r8 decisions remain accepted and untouched:** controller proof, recognition-by-commit recovery, F70/F73, custody/capability, transport, recording, F65/F68, m-1 binding, placement grill, and the no-conductor-protocol/store-change result.
- **No m-1 custody rebind is required if the revision stays within these predicates/event fields.** No credential, capability binding, lifecycle custody, or §2.7 surface needs to move.
- **No operator decision is required.** This restores the already-ratified §2.4 fail-closed rule and supplies a reachable fixture; topology and claim boundaries remain unchanged.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Return `attach-suspended` for every §2.4 suspended cause before any tuple equality check or capability mint.
2. Reserve `attach-ok`/`attach-tuple-mismatch` for the live-control, actively installed state.
3. Make the attach event outcome/reason truthful for no-state, control-loss, and malformed-update suspension.
4. Prove a worker-reachable suspended-then-success cut plus the terminal stale-generation cut in FX-TB-19.
5. Preserve every accepted r9/r8 decision and the no-custody-move result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. m-9's rebase, all rebinds, and the Master+VP interface lock remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `ed66e03892015e44e6f9ebb3d6eb514520b8c9a2f3533db5f17a0d9e5e948c69`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260717-190616.md` lints OK; routing, lineage, `DESIGN_DOC_ID`, D-3 scope, and voided-r8 posture match.
- Re-read D-3 routing `step3-mvp-design-m7/DESIGN-orchestrator-planner-20260717-185828.md`, every active attach locus, §2.4's full suspension floor, §2.7's surviving-broker control-loss row, and m-9's current attach consumption at SHA-256 `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Current m-10 owner bytes are `fe65062f5226706694ad491871c91d35e9da7c5d8fbd26f49f505cfb7c29e3e4`; their D-2 fold remains in separate fresh review and does not repair this broker-side predicate.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner folds only R9-F1, preserves the typed split and no-custody-move result, re-hashes, and sends fresh uniquely-parented bytes; m-9 rebase and rebinds remain held.
