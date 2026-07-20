## REVIEW-FOLD — MF-1 folded under explicit authority with behavioral RED→GREEN on both expiry decision paths; AO-1 comment landed; CT-1 labels corrected; one commit, captured full battery, clean head; no merge

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s10-build-slice-review
PARENT_DISPATCH_ID: s10-build-slice-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the pair-internal fold is complete; merge remains operator-only via HUMAN_MERGE_AUTHORIZATION
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-slice-review/REVIEW-FOLD-planner-20260713-031100.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the corrected-authority fold is complete in one commit `f481042` over reviewed head `9722744`; the pre-edit scope artifact preceded every source/test edit; the behavioral negative proved both replay lookup and live Apply returned unauthorized Extend at RED, then both fail closed at GREEN after the mirrored operator-FROM guards; AO-1's registered-carry comment landed; CT-1 is corrected below; full uncached battery is captured and a fresh post-commit battery is green; request the merge-decision relay TO the operator, not a merge from this seat

FOLD_SCOPE:
- frank/internal/engine/expiry.go -> in
- frank/test/fixtures/s10_timeout_test.go -> in
- frank/internal/observe/registry.go -> in
- frank/.relays/s10/ -> in
FOLD_SCOPE_RESULT: all-in

The required pre-edit artifact is `.relays/s10/fold-scope-20260713-031828.md`. It contains these exact rows and existed before the test or production edits. No OUT discovery occurred.

## Findings folded

- **MF-1 live path:** `ExpiryPrompter.Apply` now requires `resolution.Envelope.From == "operator"` before an accepted resolution can reach a pending expiry decision.
- **MF-1 replay path:** `ExpiryPrompter.existingDecision` now requires `rec.Envelope.From == "operator"` before an accepted stored resolution can supply `kill` or `extend`.
- **Behavioral negative:** `TestS10ExpiryPrompterRejectsNonOperatorExtendAcrossReplayAndLivePaths` drives the real prompter through two subtests. At RED, replay lookup returned `extend` and live Apply returned `extend`; after the two guards, both wait to the conservative context result and return `kill`.
- **AO-1:** the approved→`side-effecting-execution-refused` branch now states that approval deliberately lifts only the prompt gate and real side-effecting execution remains the registered carry. No behavior changed.
- **CT-1 corrected labels:** the **catalog** pin is `943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d`; the two **fieldspec-v5** pins are `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`. Both label families remain intact; the earlier report's abbreviation attached the catalog value to the wrong label.

## TDD and verification evidence

- RED capture: `.relays/s10/slice-review-fold-red-20260713-031947.txt` — the new test failed on both paths with `non-operator ... decision = "extend", want kill`.
- Targeted GREEN capture: `.relays/s10/slice-review-fold-targeted-green-20260713-032022.txt` — the new regression, the existing S10 expiry set, `internal/engine`, `internal/observe`, and `git diff --check` all exited 0.
- Required full uncached capture: `.relays/s10/slice-review-fold-full-green-20260713-032244.txt` — `go test ./... -count=1` exited 0; fixtures 122.337s, invariants 3.858s, and all replay suites green.
- Fresh post-commit verification at `f481042`: the exact regression passed in 0.715s; another `go test ./... -count=1` exited 0 with fixtures 123.806s, invariants 3.629s, replay 0.866s, dogfood 0.728s, and zeroloss 2.025s.
- `git diff --check HEAD^ HEAD` exited 0. The commit changes only the three named code/test files and four files under `.relays/s10/`.

## Boundary and authority

The fold does not change ODB choices, scheduler state, executor behavior, schema/config, the T10 rider, any owner bytes, or any named carry. It adds no new public surface and touches no file outside the authorized four-row scope.

No merge, push, tag, or cleanup was performed. This seat still has no merge authority.

ACTIONS_GIT_REF: `s10-comms-spine@f48104261841809eb21f2ec5f4b73631ae4ce341`; one fold commit `f481042` (`fix: require operator expiry decisions`) over reviewed head `97227447d8bf6960de73a2e706cf52e87788de78`; exact base remains `8941889a432a0a02a4183a16fb0015c230c9561b`
FINAL_GIT_STATUS_SHORT: none - clean tree

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/s10-build-slice-review/REVIEW-FOLD-implementer-20260713-032537.md` — run before handoff.
- `git -C /Users/jack/.config/superpowers/worktrees/frank/s10-comms-spine status --short` — no output after the post-commit battery.

Next requested action: s10.planner reviews this bounded fold and authors the merge-decision relay TO the operator. Merge remains blocked until a separate valid HUMAN_MERGE_AUTHORIZATION reaches the authorized merge seat.
