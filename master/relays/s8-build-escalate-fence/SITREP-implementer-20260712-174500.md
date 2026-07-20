## SITREP — exact-head battery red on two stale fixtures: Decision2 is enumerated IN the conditioned class lane; double-init is OUT because its correction changes fixture genesis composition, so one exact test-body grant is required

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-five-row-battery-red
PARENT_DISPATCH_ID: s8-build-five-row-landed
RUN_ID: s8
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one exact master fence ruling and s8.planner stale-oracle review required; live adoption and merge remain operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: five-row exact-head battery correction
IN_REPLY_TO: master/relays/s8-build-t10-close/SITREP-implementer-20260712-171500.md
FROM: s8.implementer
TO: s8.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-7.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact-head battery at `fb6e51d` is honestly red; whole-branch review finds no production blocker and identifies two stale fixtures; I enumerate the Decision2 seam under the active class lane and request one exact double-init fixture-body grant because class-lane condition 4 forbids genesis-composition edits

The required exact-head serialized battery was run as `go test -p=1 ./... -count=1` at `fb6e51d`, file-captured at `.relays/s8/s8-build-five-row-battery/go-test-all-20260712-172000.txt`, SHA-256 `8937be95a98b481f789f6f70ff491d6daf5104c33e249be21a3f40808f273f9d`, exit 1. Twenty-four packages report `ok`; `test/invariants` (all ten INV-CATALOG laws), recovery, replay, dogfood replay, and zeroloss are green. Two direct fixture reds remain; the production dogfood exit red is downstream because its governed suite includes those two.

The required whole-branch review over all 32 commits / 52 paths (`691d034..fb6e51d`) found no production correctness blocker and returned **With fixes** solely on the battery reds. It re-derived all 52 paths against the original block and later grants: no unmapped head path; `internal/observe/registry.go` is mapped distinctly at Row 2's breaker/launch seam and Row 3's absence-floor dispatch seam.

**PRE-COMMIT STALE-ORACLE ENUMERATION — IN:**

- `test/fixtures/s8_decision2_test.go` — ONLY `TestS8Decision2E1MachineryFaultsAndObservedAbsence`'s `read file IO error` subcase. It creates a directory and still expects `observe-machinery-fault`. Superseding Row-2 §4a requires a lane-creatable non-regular input to refuse typed pre-read. Retarget only this subcase to assert symbolic `not-regular-file` and the correct no-vantage authority partition; preserve the independent `git status timeout` machinery-fault leg and every other assertion. This satisfies the r4 class-lane conditions: expectation-only, no lock pin, enumerated before commit, no source/history/composition change, normal battery/review gates retained.

**OUT OF THE CLASS LANE — exact grant requested:**

- `test/fixtures/main_assembly_test.go` — ONLY `TestFrankInitTwiceRejectsExistingGenesis`, and only the two `exec.CommandContext` argument lists. `writeFixtureConfigSources` already supplies `sources["catalog"]`; add `-catalog`, `sources["catalog"]` to BOTH first and second init commands so the first creates the now-required lawful three-member production genesis and the second reaches the unchanged existing-genesis rejection oracle. Preserve all assertions and every other test/body byte.

The double-init red is a stale fixture caused by Row 1's locked production requirement, but it is NOT self-licensed by the stale-oracle class lane: r4 condition 4 puts fixture SOURCE/history composition and anything touching genesis composition OUT. Supplying the third genesis member changes this fixture's genesis composition even though it is the mechanically correct setup. The older `main_assembly_test.go` grants name different seams (`loadAssemblyRegistry` helper and a T9 digest-mutation oracle); file presence is not a wildcard.

No production change is requested. No lock pin, terminal, census, store behavior, history mechanism, or assertion changes. After the exact grant, both test-only corrections receive focused runs and a new exact-head serialized file capture; the downstream dogfood red must clear as consequence or becomes a new concrete finding.

Owner fidelity remains pending and is not bypassed. No merge, push, PR, tag, live adoption, or merge-ready claim is made.

## Verification

- Exact capture SHA reproduced: `8937be95a98b481f789f6f70ff491d6daf5104c33e249be21a3f40808f273f9d`.
- Whole review: `.superpowers/sdd/five-row-whole-branch-review.md` — no production blocker; two stale-fixture direct reds; 52-path map clean.
- Source readback: double-init omits the already-available catalog on both commands; Decision2 creates a directory while expecting machinery.
- Active class-lane source re-read: `PLAN-orchestrator-planner-20260711-205010.md`, conditions 1–5; condition 4 is the reason the double-init seam is routed rather than inferred.

ACTIONS_GIT_REF: s8-observe-spine@fb6e51d; no source/test correction committed by this relay; new battery capture only at `.relays/s8/s8-build-five-row-battery/go-test-all-20260712-172000.txt`
FINAL_GIT_STATUS_SHORT: ` M .relays/s8/INDEX.md`, `?? .relays/s8/s8-build-five-row-battery/`, and `?? .relays/s8/s8-build-t9-review/` — relay/battery artifacts only; committed source/test tree clean at `fb6e51d`
Next requested action: s8.planner reviews the enumerated Decision2 class-lane seam and routes the exact double-init test-body grant to master; s8.implementer may land the enumerated Decision2 expectation-only correction, but the double-init edit waits for the explicit grant; then both corrections receive the recaptured exact-head battery and review.
