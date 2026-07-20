## RECONCILE -- revise both pre-s8 design requests: report-only SITREPs cannot authorize amendments; replace them with unique medium-tier DESIGN lineages and grill the config-integrity and executor-ambient-authority boundaries

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step2-prep
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- routing/claim corrections first; any operator decision exposed by the grills is routed after the durable questions exist
GRILL_REQUIRED: yes
IN_REPLY_TO: master/relays/step2-prep/SITREP-orchestrator-planner-20260710-193209.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer
SUBJECT: revise `SITREP-...-193204` and `SITREP-...-193209` before either pair acts -- both bodies direct design amendments under report-only authority; both are cross-domain/hard-to-reverse and require DESIGN lineage, unique IDs, one author per artifact, Implementer design review, and durable GRILL_LOCK decisions

VERDICT: revise

## Findings

### F1 -- BLOCKER: both relays issue design work from `SITREP` / `report-only`

Each relay declares `PHASE: SITREP`, `AUTHORITY: report-only`, then directs two pair Planners to produce and lock a design-doc amendment before s8 PLAN (`...-193204.md:3-19`; `...-193209.md:3-19`). The earlier pair intakes explicitly stopped at consultation: m-7 claimed no Step-2 design lock and waited for design dispatch; m-3 said no PLAN/design amendment and held report-only until dispatched (`step2-prep/SITREP-planner-20260710-010806.md:30-32`; `...-022000.md:62,68-70`). Exact/root lint passing does not supply missing phase authority.

Replace these with actual `PHASE: DESIGN`, `AUTHORITY: design-only` dispatches carrying stable `DESIGN_DOC_ID` and `DESIGN_RECORD_KIND: design-doc`. A real pair-Planner design artifact must receive its same-owner Implementer's `PHASE: DESIGN-REVIEW` before the design lock feeds s8 PLAN.

### F2 -- BLOCKER: shared `step2-prep` IDs and joint authorship recreate the trail-collision class s7 just recorded

Both concurrent lanes reuse `DISPATCH_ID: step2-prep`, and each asks two Planners for "one joint design note" while leaving file ownership/split to them. That permits duplicate IDs, competing writers, proxy/co-authorship, and ambiguous immediate parents -- the same `one_by_id`/ID-collision class preserved in S7A-TRAIL-FINDINGS.

Use unique sub-dispatch and design-doc IDs per owned artifact. Each Planner authors only its domain section/artifact under its own `FROM`, receives its own Implementer review, and master reconciles the reviewed m-7/m-2 and m-7/m-3 artifacts. Shared content is a boundary contract, not a shared writer.

### F3 -- BLOCKER: both designs trigger the grill rule and medium ceremony

The relays mark `CEREMONY_TIER: small`, `GRILL_REQUIRED: no`, but both are cross-domain contracts and hard-to-reverse trust surfaces that gate several downstream implementation choices. The design-grill rule therefore requires `GRILL_REQUIRED: yes`, a durable `GRILL_LOCK_ID`, and medium ceremony before lock. "Shape agreed" narrows the search space; it does not settle the member/home/integrity mechanism or the executor's ambient authority.

### F4 -- executor claim boundary must distinguish handle absence from ambient process authority

`...-193209` says the executor receives only candidate bytes/declared inputs, "never writes," and has zero conductor authority, while simultaneously limiting the mechanism to process/context-grade isolation and retaining D5. Repo-resident suite code can otherwise inherit environment variables, file descriptors, filesystem/network/process access, toolchain caches, a writable working tree, and child processes. An absence-set of conductor handles proves none of those are provided APIs; it is not OS containment and cannot support an unqualified "never writes."

The executor GRILL_LOCK must resolve at least: snapshot/worktree identity; working directory and writable-temp/cache policy; inherited environment and descriptors; credential/signing-secret scrubbing; filesystem/store/config/outbox path exposure; network and subprocess policy; timeout plus descendant/process-group termination; output/resource ceilings; cleanup after timeout/crash; operator-gated side-effect semantics; and exact binding of candidate bytes/declared inputs to the typed verdict. Scope the claim to **no conductor-governed mutation handles and no writes to canonical conductor state through the provided surface** unless the operator elects an OS sandbox as a separate design addition.

The adversarial fixture must report handle-surface absence separately from the accepted same-uid ambient-access residual. Do not label the latter "by construction."

### F5 -- config digest integrity does not prove owner fidelity, and the governed census set is incomplete as stated

`...-193204` asks for catalog "single-writer/owner-fidelity enforcement at load." Section-7 digest/genesis verification proves the loaded bytes match an operator-authorized config history; per-section metadata can state an owner. Without owner-signed attestations, load cannot prove the named domain owners actually reviewed the change. The design must either keep owner fidelity as a relay/design-review gate and claim only digest/member/provenance-shape enforcement at load, or explicitly design an attestation mechanism (a larger grilled addition).

The governed discovery descriptor must cover the final s7 mechanism, not only "walk roots + two-file boundary": scan roots/exclusions, registered boundary files, tree-wide portable/contextual egress idioms, and the exact current site/family census all affect whether the law bites. Lock the canonical runtime/deployment home and byte-copy path for the source catalog; `test/invariants/catalog.v1.json` is not by itself a guaranteed installed runtime-config home.

## Required Replacement Shape

1. Issue unique medium-tier DESIGN dispatches for the m-7 config host, m-2 config/atom grammar, m-7 executor host, and m-3 registry/probe semantics (or an equivalently collision-free one-author-per-artifact split).
2. Each design carries `DESIGN_DOC_ID`, `DESIGN_RECORD_KIND: design-doc`, `GRILL_REQUIRED: yes`, and its durable `GRILL_LOCK_ID`; each receives same-owner Implementer DESIGN-REVIEW.
3. Master reconciles the reviewed cross-domain contracts into the s8 design input. Only then may the s8 PLAN consume them.
4. No pair should treat the two superseded SITREPs as design authority.

## Verification

- Both incoming exact-file lints: OK; `--relay-root master/relays/step2-prep` plus both files: OK.
- Live kickoff design items 1-2 say bounded amendments via the ritual and bind at/before s8 PLAN; they do not convert a report-only relay into DESIGN authority (`STEP-2-KICKOFF.md:54-58`).
- Locked config design proves top-level digest/genesis integrity and operator-authorized config history, not owner participation (`m-7 ... conductor-core-design.md:107-111`).
- Locked m-3 record calls executor isolation a non-locking build carry and preserves D5; it does not settle process containment (`m-3 ... observe-evidence-design.md:88,221`).
- `frank/main` is clean at `2e1b4f0`, so these are governance/design corrections only.

Next requested action: withdraw/supersede `...-193204` and `...-193209`, issue the corrected DESIGN lineages, and hold s8 PLAN until their reviewed reconciled locks exist.

ACTIONS_GIT_REF: wrote this reviewer relay and appended its row to `master/relays/INDEX.md`; no governance design doc, `frank/` source/test, branch, commit, merge, tag, or remote action.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable -- not a git repository (docs workspace)
- `frank/`: `## main...origin/main`

