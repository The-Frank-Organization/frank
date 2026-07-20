## DESIGN-REVIEW - m-5.implementer adversarial review of step3 m-5 ceiling-host amendment

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m5-ceiling-review-r1
PARENT_DISPATCH_ID: step3-amend-m5-ceiling
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded owner-amendment review; no operator decision requested
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-amend-m5-ceiling-host
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-amend-m5-ceiling/DESIGN-planner-20260715-081500.md
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-9.planner, m-7.planner
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

DESIGN_REVIEW_VERDICT: must-revise

m-5.planner - I reviewed the owner amendment, the m-5-to-m-10 COORD, the master dispatch, the VP fold-review context, the ratified Step-3 architecture packet, the locked m-5 design, and the m-10 charter. Verdict: must-revise.

The host relocation is directionally correct: m-5 remains the policy owner, m-10 is only the app-side enforcement host, the conductor is not made an app supervisor, and the proposal correctly avoids a design-of-record fold, lock, PLAN, code, or joint interface-lock claim in this leg. The blockers are both inside the ceiling-artifact interface: the source of `seat_archetype` is not pinned tightly enough, and `policy_digest` freshness does not yet close rollback/stale-generation cases.

This review is read-only and scoped only to `step3-amend-m5-ceiling/DESIGN-planner-20260715-081500.md`. It grants no lock, no stage-2 consumer dispatch, no PLAN, no T4 code token, no `frank/` edit, and no Master/VP first-stage join lock.

## Findings

### F1 - `seat_archetype` provenance is under-specified, leaving a hidden policy choice at the m-10 binding point

The proposal says Layer 2 resolves from Layer 1 for the worker's `seat_archetype`, then says `seat_archetype` is a pinned run-manifest input and not an m-10 policy choice (`DESIGN-planner-20260715-081500.md:40-51`). It also makes m-10 the Layer-2 writer at worker spawn (`:56-58`) and the m-10 manifest the instance home (`:58`, `:91`). That is almost the right split, but it does not name the authoritative writer/source of the `seat_archetype` input or the validation m-10 performs before copying it.

That omission matters because the locked m-5 design makes `seat_archetype` the value that fixes the authority ceiling at spawn (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:18-20`, `:95`). The ratified Step-3 packet says m-10 owns the app run manifest and owns no policy (`master/STEP-3-ARCH-AMENDMENT.md:27-28`, `:47`, `:103`), while Step-3 uses a pinned manifest rather than an m-4 routing decision (`:66`, `:86`, `:112`). If the source of the worker `seat_archetype` is not pinned, m-10 can indirectly choose the ceiling by choosing the archetype while still "only" projecting m-5 policy.

Required revision: add an explicit provenance field and validation rule for the archetype input, for example `seat_archetype_source` / `spawn_authority_ref` / equivalent. The revised contract must say who is allowed to pin the worker `seat_archetype` in Step-3, what evidence m-10 checks, that m-10 only copies the already-pinned value into `bound_ceiling`, and that absent, ambiguous, mismatched, or app-local archetype input fails closed. If the answer is "operator/master-provided pinned run manifest for the single MVP lane", say that explicitly and carry it into the COORD to m-10.

### F2 - `policy_digest` equality alone does not close rollback or stale-generation freshness

The proposal validates freshness with `policy_digest == the currently-loaded Layer-1 stamp` plus active `run_id` / `worker_identity` key equality (`DESIGN-planner-20260715-081500.md:60-61`). It asks this review to probe digest rollback (`:106`), but the artifact currently carries only a digest, not a monotonic trusted-config generation, top-level config digest, or rollback disposition. A rollback to an older policy section can make an old digest look "current" again, or can loosen a future binding without passing through the typed reconfiguration/respawn gate.

That conflicts with the locked fail-closed principle that old or partial records under-authorize rather than widen (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:86-95`) and with the Step-3 direct-route ceiling rule: a ceiling raise must use the m-5 typed reconfiguration/respawn contract and gate, not a textual/direct shortcut (`master/STEP-3-ARCH-AMENDMENT.md:118-122`; `DESIGN-planner-20260715-081500.md:65-69`). Digest mismatch is not the whole stale class if rollback can restore a previous digest as the apparent current stamp.

Required revision: define freshness over the active trusted-config section plus a monotonic generation/epoch or top-level config digest, and state that rollback, unknown generation, partial resolution, or current-policy regression is stale/fail-closed unless it is itself the result of the governed typed reconfiguration/respawn path. The bound artifact should carry enough identity to distinguish "same bytes still current" from "old bytes reintroduced" at the m-10 load path.

## Accepted Portions

- The enforcement-host move is a real amendment to locked m-5 section 9 and is correctly staged rather than silently rewriting the locked design (`DESIGN-planner-20260715-081500.md:26-35`, `:87-94`; `master/STEP-3-ARCH-AMENDMENT.md:108-112`).
- The policy/host split is correct at the owner level: m-5 owns policy, resolution rule, schemas, lattices, and floors; m-10 hosts/invokes/loads/applies (`DESIGN-planner-20260715-081500.md:53-63`).
- The absent/mismatch/malformed/key-mismatch fail-closed floor is directionally right and maps to the locked m-5 per-axis floors (`DESIGN-planner-20260715-081500.md:60-61`; locked m-5 design `:86-95`).
- Immutable binding plus respawn rather than in-place mutation is the right mechanism shape (`DESIGN-planner-20260715-081500.md:65-69`; `master/STEP-3-ARCH-AMENDMENT.md:122`).
- The planner correctly treats the interface as proposed pending m-10 convergence and Master/VP reconciliation, not as already interface-locked (`DESIGN-planner-20260715-081500.md:24`, `:96-100`, `:118`), which avoids the VP F22 circular completion defect on the m-5 side.

## Revision bar

Return a revised m-5 DESIGN that:
1. Adds the authoritative source/provenance and m-10 validation rule for `seat_archetype`.
2. Adds rollback/stale-generation handling for `policy_digest` freshness.
3. Mirrors those exact revisions into the m-10 COORD so m-10 converges on the same interface bytes.
4. Preserves the current scope fence: no lock, no design-of-record fold, no PLAN, no code, and no joint interface-lock until m-10 convergence plus Master/VP reconcile.

No new operator product fork is required if the revision only pins these two interface mechanics. If the revision changes who may choose worker archetype, who may raise ceilings, or who writes policy values, route that as a fresh Master/operator decision.

## Verification

Pre-write review evidence:
- `sed -n '1,260p' master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-081500.md` - read addressed m-5 planner amendment.
- `sed -n '1,240p' master/relays/step3-amend-m5-ceiling/DESIGN-orchestrator-planner-20260715-073010.md` - read master m-5 amendment cue.
- `sed -n '1,240p' master/relays/step3-amend-m5-ceiling/COORD-planner-20260715-082000.md` - read m-5 to m-10 interface proposal.
- `sed -n '1,220p' master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-080000.md` - checked VP fold-review and F22 sequencing correction.
- `nl -ba master/STEP-3-ARCH-AMENDMENT.md | sed -n '1,180p'` - checked packet boundary, manifest, m-5 amendment, direct-route, and grill anchors.
- `nl -ba master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md | sed -n '80,180p'` - checked locked ceiling floors, per-archetype mapping, and old enforcement-host text.
- `nl -ba master/domains/m-10-app-control-plane/README.md | sed -n '1,120p'` - checked m-10 ownership and no-policy boundary.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-081500.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/COORD-planner-20260715-082000.md` - OK.

Post-write verification:
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-082430.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/step3-amend-m5-ceiling` - OK.
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-082430.md` - no output.
- `rg -n "20260715-082430|step3-amend-m5-ceiling-review-r1|DESIGN-REVIEW-implementer-20260715-082430" master/relays/INDEX.md master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-082430.md` - relay line 6 and INDEX line 1284 present.
- `git status --short` from `/Users/jack/Programming/harness` - `fatal: not a git repository (or any of the parent directories): .git`.
- `git -C frank status --short` - no output.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no domain-doc edit, no source/code/`frank/` edit, no PLAN, no IMPL, no runtime action.
FINAL_GIT_STATUS_SHORT: harness root is not a git repo (`fatal: not a git repository (or any of the parent directories): .git`); `frank/` clean (`git -C frank status --short` produced no output).
RELAY_LINT: OK `master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-082430.md`
DISPATCH_ROOT_LINT: OK `master/relays/step3-amend-m5-ceiling`
Next requested action: m-5.planner revises the amendment and COORD rows for F1/F2, then returns the revised DESIGN for another m-5.implementer review. m-10 convergence and Master/VP join-lock remain held.
