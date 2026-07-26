## DESIGN-REVIEW - m-1 adversarial review r2 of the stage-6 lane-2 env/redaction additive delta

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m1-review-r2
PARENT_DISPATCH_ID: step3-relock-dag-m1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the two remaining defects have deterministic pair-local repairs within the released scope
GRILL_REQUIRED: no - no product fork remains open; this review narrows only exact environment and fixture semantics
DESIGN_DOC_ID: step3-relock-dag-m1
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m1/DESIGN-planner-20260721-235452.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-relock-dag-m1/DESIGN-REVIEW-implementer-20260722-000641.md
SUBJECT: must-revise - make the baseline environment byte-exact and bind sentinel expectations to their ingress routes

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner - I reviewed the revised delta at exact SHA-256 `734df728ccdd15eee675ef09c05b55023b1c27c298efcf3f2ea77de33cd3bc74`, the exact released rev2 dispatch at `9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`, amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, and frozen m-1 at `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.

The revision materially closes r1 F1-F5: the claim split, closed authority family, descriptor-grain at-rest boundary, empty class (ii), canonical digest preimage, and producer-first parking are all accepted. Two exactness defects remain. They do not reopen the architecture, but they block pair approval because the current bytes permit an internally false child environment and an ambiguous redaction oracle.

This review grants no stage-6 lock, producer confirmation, redaction co-sign, PLAN, T4/code token, credential, provider call, release binding, live E3, `frank/` edit, merge, or deploy.

## Findings

### R2-F1 - The baseline is not yet one exact, truthful, secret-free presented environment

The table binds `PWD` to the descriptor `cwd` (`2026-07-22-stage6-lane2-env-redaction.md:27`), while the consumed F103 descriptor contract defines `cwd` as a workspace-root-relative path and serializes the root as `"."` (`2026-07-22-relock-lane2-m9-delta.md:223`). Those are different facts. `PWD` describes the process working directory; a relative descriptor value can disagree with the physical execution cwd. The mismatch is executable: from this workspace, `env -i PWD=. /bin/bash --noprofile --norc -c 'printf "%s\n" "$PWD"; /bin/pwd -P'` reports `.` and `/Users/jack/Programming/harness`. A tool that trusts `PWD` and one that asks the OS therefore observe different directories, despite the table claiming a POSIX baseline and descriptor consistency.

The same table permits `TZ` to be either `UTC` or an unspecified "operator-fixed constant" (`:23`), but Section 1.3 then grounds digest safety in a fully system-sourced set with no operator-supplied values (`:33`). That optional value is neither one exact baseline nor consistent with the stated source proof. `PATH` is also a build/config constant, and generated `HOME`/`TMPDIR`, resolved paths, and the OS user name are not all "system constants"; therefore the sentence claiming non-UTF-8 "cannot fire in the MVP" is unsupported even though rejecting such a value is a valid rule.

Required revision: set `PWD` to the absolute, symlink-resolved execution cwd obtained by resolving the canonical descriptor `cwd` under the selected workspace root; keep the descriptor itself in its one relative F103 encoding. Pin `TZ=UTC` for this MVP, with any future operator-selected timezone routed through an owned surface. Describe the digest input as the closed, reviewed non-secret baseline rather than as having no operator/config provenance, and retain non-UTF-8 as a reachable pre-spawn typed reject instead of claiming it cannot occur. The digest remains over the exact values actually presented.

### R2-F2 - Sentinel expected results still depend on an unstated ingress route

Section 2 correctly says arbitrary open content has no byte-absence guarantee (`:41`), but Section 2.4 says sentinel S-A/S-B, replay, ticket, and capability values are merely "driven through a session" and must all be absent (`:49`); Section 5 repeats four undifferentiated NOT-list legs (`:60`). For replay and authority objects, absence follows only when the sentinel enters through the prohibited typed object route. For S-A/S-B, accidental-disclosure absence follows when the sentinel is planted in a governed secret source such as the credential holder, parent environment, or private file and ordinary execution does not explicitly read it. If the same bytes are deliberately supplied as `input_item.content` or returned as open `tool_result.content`, the contract explicitly permits them to be logged under the accepted same-UID/operator-content residual. The current fixture wording therefore admits both an expected-absent and an expected-present execution.

Required revision: make every sentinel leg route-specific. Pin S-A/S-B to one or more governed secret-source plant points and prove no accidental transition into allowed open content; pin replay/ticket/capability sentinels to their prohibited typed source-object routes and prove the writer rejects or omits those objects. Add the complementary open-content ceiling leg: the same marker deliberately placed in an allowed content string is outside the absence guarantee and must not be used as a redaction pass/fail oracle. Carry the same route labels into the parked m-9 criteria.

## Accepted portions

- **R1 F1 is substantively closed.** Structural exclusion is limited to typed routes; open content carries the honest confusion-not-malice and same-UID ceiling; scanning is not claimed; the absolute-guarantee mismatch route is named.
- **R1 F2 is closed.** F59 tickets, USE capabilities, broker/control handles, future authority-tagged types, credential references, and attribution identifiers are classified coherently with fail-closed type evolution.
- **R1 F3 is closed at design grain.** Exclusive no-follow create, no-follow reopen, opened-descriptor identity/mode/owner/link checks, verified-directory rotation, directory durability, resolved containment, and corresponding negative legs are all required.
- **R1 F4 is closed except for R2-F1's value exactness.** There are two positive classes plus deny-all; class (ii) is empty; no m-10 manifest surface is invented; HOME, TMPDIR, PATH, locale names, and control-variable absence are explicit.
- **R1 F5 is closed.** The oracle-by-shape claim is withdrawn; secret-free input is the safety condition; one JCS object recipe, duplicate rejection, and parent-environment exclusion are bound.
- **Lineage and staging remain correct.** The base is byte-identical, the delta remains additive, and all producer confirmations and the redaction join co-sign remain explicitly parked.

## Revision bar

Return fresh delta bytes and a fresh hash that:

1. Binds `PWD` to the actual absolute resolved execution cwd while preserving the relative descriptor encoding.
2. Pins `TZ` to one value and states the baseline's non-secret source argument without denying trusted config/generated/OS provenance.
3. Keeps non-UTF-8 as an honest reachable pre-spawn reject.
4. Gives each absent-sentinel fixture a prohibited or governed secret-source ingress route and adds the complementary allowed-open-content ceiling leg.
5. Preserves every accepted r1 fold, the frozen base, construct-from-empty, K6 exclusion, F57-narrow ceiling, descriptor-grain at-rest boundary, producer-first parking, and separate consumer/co-sign gates.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte review. Do not attach a producer confirmation or redaction co-sign before the relevant producer bytes are pair-approved and reviewed separately.

## Verification

- Reproduced SHA-256: delta `734df728ccdd15eee675ef09c05b55023b1c27c298efcf3f2ea77de33cd3bc74`; released rev2 `9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`; frozen m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Exact-file linted the addressed planner relay; verified `TO`, `IN_REPLY_TO`, `DESIGN_DOC_ID`, current lane ordering, no superseding m-1 file, and author-now/park-later scope.
- Re-ran the r1 F1-F5 bar against the revised bytes and consumed m-9 F103/env-digest contract; all original findings close except the narrower exactness defects above.
- Executed the relative-`PWD` behavior probe with an empty environment and compared shell `$PWD` to external `/bin/pwd -P`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source contract, historical relay, frozen byte, `frank/` source, branch, commit, lock, producer confirmation, redaction co-sign, PLAN, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `c78da3815a34`.
RELAY_LINT: OK - exact-file and lane-root relay-lint.py exit 0
Next requested action: m-1.planner folds only R2-F1 and R2-F2, re-hashes the delta, and sends fresh uniquely-parented bytes for pair review; producer confirmations and the redaction co-sign remain parked.
