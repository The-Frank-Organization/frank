## DESIGN-REVIEW - m-5.implementer review of Step-3 permissive-tools amendment r4

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m5-ceiling-review-r4
PARENT_DISPATCH_ID: step3-amend-m5-ceiling
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review-only; the operator's Branch-B and ambient-bash risk choices are treated as recorded
GRILL_REQUIRED: no -- this review consumes the planner's folded GRILL_LOCK text; it does not open a new grill
DESIGN_DOC_ID: step3-amend-m5-ceiling-host
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-amend-m5-ceiling/DESIGN-planner-20260715-152000.md
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner, m-1.planner, m-7.planner, m-9.planner
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5
RELAY_PATH: master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-152530.md
SUBJECT: MUST-REVISE -- r4 closes the prior no-change/grill/bash-honesty issues, but it still must pin VP F38: the new m-5 hash may permit stage-2 design only and must not independently permit positive runtime tool dispatch before mandatory audit and m-9/m-10 enforcement are reviewed, implemented, and bypass-proven.

Verdict: must-revise.

This is close. The planner's 152000 relay fixes the earlier invalid no-change nod, restores `GRILL_REQUIRED: yes`, folds a durable `GRILL_LOCK_ID`, and states the operator-accepted ambient `bash` residual honestly. I accept those parts. The remaining blocker is narrow but load-bearing: the text still lets a reader infer that finalizing the m-5 addendum hash plus m-10 consumption is enough to start positive runtime tool dispatch.

## Findings

### F1 -- Missing F38 pre-execution gate

VP F38 is explicit: the m-5 amendment may authorize the policy branch, but positive runtime dispatch remains non-consumable until the m-3 audit contract, the app-side writer/reader boundary, and the m-9/m-10 enforcement integration are owner-reviewed, implemented, and proven to emit one auditable event per attempted call with no unaudited bypass (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-150756.md:45-49`).

The 152000 design makes universal audit mandatory (`DESIGN-planner-20260715-152000.md:34-38`) and says audit is not preventive (`:38`). That part is correct. But the path fence still says that after review, durable grill, new hash, m-10 consumption, Master+VP review, and operator ratification, "deny-all is the floor" only "until that lands" (`:68-69`), and the requested next action says the MVP remains `tool->none` "until then" (`:79-81`). Together with the posture sentence that "positive tools dispatch" inside the pinned ceiling (`:44`), this still makes the new m-5 hash look like the final runtime consumption gate.

Required revision: add a hard pre-execution clause to the addendum semantics and the path/scope fence:

`The new m-5 Step-3 addendum hash authorizes only the policy branch and downstream stage-2 design consumption. It does not authorize positive runtime tool dispatch. Positive runtime tool dispatch remains non-consumable until the m-3-owned universal audit contract, its app-side writer/reader boundary, and the m-9/m-10 enforcement integration are owner-reviewed, implemented, and proven to emit one auditable event for every attempted tool call with no unaudited bypass. Exact schema, emitter, carrier, and integration proof remain with the stage-2 owners.`

Then adjust the posture and next-action wording so they say the permissive rule is the policy branch to be designed against, not an executable runtime grant after the m-5 hash lands.

## Accepted

- Addendum overlay form is acceptable: a Step-3 boxed overlay over `643dd7c2...`, with Step-4 sunset and the base freshness/capability target preserved.
- Prior r3 blocker is closed: unchanged `643dd7c2...` remains deny-all, and positive tools require a changed rule/new hash rather than a vacuous "pinned equals current-active" nod.
- GRILL requirement is restored and folded in the reviewed bytes: `GRILL_REQUIRED: yes` and `GRILL_LOCK_ID: step3-permissive-tools-grill` are present in the planner relay (`:11-14,49-63`).
- Ambient `bash` authority is honest: no cwd confinement, sandbox, or irreversibility gate is claimed for Step-3 (`:40-42`).
- The relaxed-versus-preserved split is acceptable once F1 is added: freshness proof and sandbox defer to Step-4; the coarse tool-class ceiling, structural fail-closed floor, and archetype provenance remain in the m-5 policy contract (`:34-42`).
- H-11 remains correctly non-operative for this amendment (`:65-66`).

## Revision bar

Return a revised DESIGN that folds F1 exactly or equivalently. Do not mint the addendum file/hash yet. Do not present m-5 approval as runtime positive tool authority. The next review can approve if the F38 gate is explicit and the surrounding "until that lands" language no longer implies that the m-5 hash alone permits runtime dispatch.

## Verification

- Read exact incoming relay `master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-152000.md`; verified `TO: m-5.implementer`, `DESIGN_DOC_ID: step3-amend-m5-ceiling-host`, `GRILL_REQUIRED: yes`, and `GRILL_LOCK_ID: step3-permissive-tools-grill`.
- Exact-file lint of incoming relay was checked before this review: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-planner-20260715-152000.md` returned `OK`.
- Read master amendment request `step3-amend-m5-ceiling/RECONCILE-orchestrator-planner-20260715-145500.md`, master fold `step3-arch-packet/RECONCILE-orchestrator-planner-20260715-150000.md`, VP fold review `step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-150756.md`, base m-5 contract `master/domains/m-5-workflows-archetypes/design/2026-07-15-ceiling-artifact-contract.md`, and prior m-5 implementer review `DESIGN-REVIEW-implementer-20260715-144650.md`.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-152530.md` - OK.
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-152530.md` - no output.
- `rg -n "20260715-152530|step3-amend-m5-ceiling-review-r4|DESIGN-REVIEW-implementer-20260715-152530" master/relays/INDEX.md master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-152530.md` - relay header lines 6 and 22 plus INDEX line 1318 present.
- `git status --short` from `/Users/jack/Programming/harness` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.
- `git -C frank status --short --branch` - `## main...origin/main`.
- `git -C frank rev-parse --short HEAD` - `502e06c`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no governing-source doc, contract file/hash, historical relay, `frank/` source, branch, commit, lock, PLAN, code, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK `master/relays/step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-152530.md`
Next requested action: m-5.planner folds F1, returns revised DESIGN bytes for exact-byte review, and keeps the runtime positive-tool path non-consumable until mandatory audit and m-9/m-10 enforcement are owner-reviewed, implemented, and bypass-proven.
