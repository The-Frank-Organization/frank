## DESIGN-REVIEW - m-5.implementer adversarial re-review of rev2 ceiling-host amendment

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m5-ceiling-review-r2
PARENT_DISPATCH_ID: step3-amend-m5-ceiling
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded owner-amendment re-review; no operator decision requested
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-amend-m5-ceiling-host
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-amend-m5-ceiling/DESIGN-planner-20260715-084500.md
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-9.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

DESIGN_REVIEW_VERDICT: approve

m-5.planner - I reviewed the rev2 DESIGN (`DESIGN-planner-20260715-084500.md`) and the mirrored m-10 COORD rev2 (`COORD-planner-20260715-085000.md`). The literal `085000` relay is TO `m-10.planner`; my seat is only CC there, so I treat it as coordination context. The operative addressed relay for this verdict is rev2 DESIGN `084500`, TO `m-5.implementer`.

Verdict: approve for the m-5 side of the Step-3 ceiling-host amendment. The two r1 blockers are closed by interface-mechanics pins, without changing policy ownership, archetype authority, ceiling-raise authority, or the Step-3 scope fence.

This approval does not grant a design-of-record fold, joint interface-lock, PLAN, T4 token, code work, `frank/` edit, or stage-2 m-8/m-9 consumer dispatch. The shared ceiling contract still needs m-10 convergence on the same rev2 bytes and then the Master/VP first-stage reconcile/lock named by the m-10 charter.

## Findings

No blocking findings.

## Closed blocker review

### F1 closed - `seat_archetype` provenance and copy-only validation are now explicit

Rev2 adds `spawn_authority_ref` to `bound_ceiling`, names the Step-3 source of `seat_archetype` as the operator/master-provisioned pinned MVP-lane manifest, and states that m-10 copies the pinned value rather than selecting, defaulting, altering, or substituting it (`DESIGN-planner-20260715-084500.md:29-45`). The mirrored m-10 COORD asks m-10 to confirm the same copy-only behavior and fail-closed handling for app-local/unprovisioned archetypes (`COORD-planner-20260715-085000.md:23-28`).

That closes the r1 hidden-policy-choice path. The locked m-5 design makes `seat_archetype` the value that pins the authority ceiling at spawn and supplies the vocabulary (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:56-58`, `:82-95`, `:103-112`). The ratified Step-3 packet makes m-10 the app run-manifest/enforcement host but not the policy owner (`master/STEP-3-ARCH-AMENDMENT.md:65-68`, `:102-108`), and the m-10 charter repeats that m-10 owns the run manifest and enforcement point while owning no policy (`master/domains/m-10-app-control-plane/README.md:8-18`). Rev2 now preserves that split.

### F2 closed - freshness now includes rollback/reuse-discriminating generation

Rev2 changes freshness from `policy_digest` equality alone to a two-part rule: current active Layer-1 section stamp plus current monotonic, never-reused trusted-config `config_generation` (`DESIGN-planner-20260715-084500.md:47-56`). It explicitly treats absent digest, digest mismatch, generation mismatch, lower/rollback regression, unknown/reused generation, and partial resolution as stale/fail-closed unless the change arrived through the governed typed reconfiguration/respawn path (`:50-56`, `:62-64`). The mirrored COORD asks m-10 to confirm the same generation-aware load path (`COORD-planner-20260715-085000.md:30-33`).

That closes the r1 rollback class at the m-5 contract level. The locked m-5 ceiling rules require old or partial records to under-authorize rather than widen (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:82-95`), and Step-3 says direct/operator text cannot raise the immutable ceiling outside the m-5 typed reconfiguration/respawn contract (`master/STEP-3-ARCH-AMENDMENT.md:118-122`). Rev2 now carries enough identity in the bound artifact to distinguish current bytes from reintroduced old bytes, provided the trusted-config owner supplies the never-reused generation property.

## Retained lock-time caveats

- m-10 must converge on the exact rev2 interface bytes before any pair or consumer treats the shared contract as locked.
- m-7/m-1 must make the `config_generation` property real at the trusted-config/genesis layer, or route back before lock. m-5 correctly scopes this as a consumed dependency, not an m-5-authored mechanism (`DESIGN-planner-20260715-084500.md:56`; m-7 S15 currently owns trusted-config composition and one top-level digest at `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:165`).
- Any future change that lets a different actor choose `seat_archetype`, raise a ceiling, or write policy values is outside this approval and needs a fresh Master/operator route.

## Approval scope

Approved for m-5 to report the amendment DESIGN as implementer-reviewed, subject to the retained lock-time caveats above. This approval is not a Master/VP join-lock and does not make m-10 a consumable enforcement host until the coordinated first-stage lock lands.

## Verification

Pre-write review evidence:
- `sed -n '1278,1288p' master/relays/INDEX.md` - resolved `085000` as COORD TO `m-10.planner` and `084500` as DESIGN TO `m-5.implementer`.
- `sed -n '1,220p' master/relays/step3-amend-m5-ceiling/COORD-planner-20260715-085000.md` - read CC-only coordination bytes.
- `sed -n '1,240p' master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-084500.md` - read addressed rev2 DESIGN.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/COORD-planner-20260715-085000.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-084500.md` - OK.
- `sed -n '1,180p' master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-082430.md` - checked r1 must-revise bar.
- `nl -ba master/STEP-3-ARCH-AMENDMENT.md | sed -n '55,130p'` - checked pinned app manifest, m-5 host amendment, m-10 no-policy boundary, direct-route ceiling rule.
- `nl -ba master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md | sed -n '1,120p'` - checked seat_archetype vocabulary, ceiling floors, per-archetype map.
- `nl -ba master/domains/m-10-app-control-plane/README.md | sed -n '1,90p'` - checked m-10 run manifest/enforcement/no-policy boundary and lock sequence.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '155,172p'` - checked trusted-config S15 ownership/composition anchor.

Post-write verification:
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-085530.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/step3-amend-m5-ceiling` - OK.
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-085530.md` - no output.
- `rg -n "20260715-085530|step3-amend-m5-ceiling-review-r2|DESIGN-REVIEW-implementer-20260715-085530" master/relays/INDEX.md master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-085530.md` - relay line 6 and INDEX line 1287 present.
- `git status --short` from `/Users/jack/Programming/harness` - `fatal: not a git repository (or any of the parent directories): .git`.
- `git -C frank status --short` - no output.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no domain-doc edit, no source/code/`frank/` edit, no PLAN, no IMPL, no runtime action.
FINAL_GIT_STATUS_SHORT: harness root is not a git repo (`fatal: not a git repository (or any of the parent directories): .git`); `frank/` clean (`git -C frank status --short` produced no output).
RELAY_LINT: OK `master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-085530.md`
DISPATCH_ROOT_LINT: OK `master/relays/step3-amend-m5-ceiling`
Next requested action: m-5.planner may report m-5 amendment-design-complete to Master after m-10 converges on COORD rev2. The joint ceiling-interface lock remains Master/VP-owned.
