## DESIGN-REVIEW - m-1 adversarial review r3 of the stage-6 lane-2 env/redaction additive delta

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m1-review-r3
PARENT_DISPATCH_ID: step3-relock-dag-m1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining stale provenance clause has one deterministic pair-local correction
GRILL_REQUIRED: no - no product or boundary choice remains open
DESIGN_DOC_ID: step3-relock-dag-m1
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m1/DESIGN-planner-20260722-001938.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-relock-dag-m1/DESIGN-REVIEW-implementer-20260722-002830.md
SUBJECT: must-revise - remove the last false no-operator-provenance inference; all r2 mechanism folds pass

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner - I reviewed the revised delta at exact SHA-256 `2f09175192e4ee1d7da5a8db5f41f0b034c62bbb5cad788eb864bde28d122eee`, the exact released rev2 dispatch at `9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`, amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, and frozen m-1 at `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.

R2-F1 and R2-F2 are mechanically folded: `PWD`, `TZ`, reachable UTF-8 rejection, route-specific absent sentinels, the complementary expected-present open-content ceiling leg, and the parked m-9 criteria all pass. One stale parenthetical still contradicts the accepted provenance correction and the r2 revision bar. It is narrow, but it remains in the load-bearing digest-sensitivity paragraph and therefore blocks exact-byte approval.

This review grants no stage-6 lock, producer confirmation, redaction co-sign, PLAN, T4/code token, credential, provider call, release binding, live E3, `frank/` edit, merge, or deploy.

## Finding

### R3-F1 - The digest paragraph still falsely infers that an empty class (ii) removes operator/config provenance

Section 1.3 now correctly grounds safety in a `CLOSED, REVIEWED, NON-SECRET set` and explicitly names trusted-config, generated-directory, and OS-fact provenance (`2026-07-22-stage6-lane2-env-redaction.md:33`). But the same sentence retains the parenthetical `class (ii) empty ⇒ no operator-supplied values`. That implication does not follow. Class (ii) being empty proves only that no run-manifest environment member is admitted through that class. Class (i) still includes a build/config `PATH`, an absolute `PWD` derived from the selected workspace root, generated directory paths, and OS identity facts. Those values can carry trusted operator/configuration influence without becoming secret or unsafe.

This is the exact source-model overclaim the r2 bar required the revision to avoid: digest safety comes from reviewing each closed baseline source and value as non-secret, not from asserting that no operator/config provenance exists. The incoming relay says the false wording is gone, but the exact design bytes retain it.

Required revision: delete the parenthetical, or narrow it to `class (ii) empty ⇒ no run-manifest/operator-selected class-(ii) values`. Keep the operative proof as the closed, reviewed, non-secret baseline with trusted-config/generated-directory/OS-fact provenance. No mechanism, schema, route, fixture, producer criterion, or operator decision changes.

## Accepted portions

- **R2-F1's execution facts pass.** `PWD` is the absolute symlink-resolved execution cwd, the F103 descriptor remains workspace-root-relative, `TZ=UTC` is pinned, and the truth fixture compares child `$PWD`, `/bin/pwd -P`, and descriptor resolution.
- **R2-F1's UTF-8 correction passes.** Non-UTF-8 values are a reachable pre-spawn typed reject, including generated-path inputs; the digest remains over the exact presented set.
- **R2-F2 passes.** Governed secret-source plants have an expected-absent accidental-disclosure result; replay/ticket/capability objects use prohibited typed routes; allowed open content has the complementary expected-present ceiling leg and is not a redaction oracle.
- **All prior accepted boundaries remain intact.** Construct-from-empty, empty class (ii), closed authority types, K6 exclusion, F57-narrow ceiling, descriptor-grain at-rest handling, no scanning claim, canonical JCS preimage, and parent-environment exclusion remain coherent.
- **Lineage and staging pass.** The frozen base is byte-identical, the delta is additive, the addressed relay is current, and producer confirmations plus the redaction co-sign remain parked.

## Revision bar

Return fresh delta bytes and a fresh hash that remove or precisely narrow the stale `no operator-supplied values` parenthetical while preserving every other rev3 byte semantically. Issue a fresh uniquely-parented DESIGN for exact-byte review. Do not attach a producer confirmation or redaction co-sign before the relevant producer bytes are pair-approved and reviewed separately.

## Verification

- Reproduced SHA-256: delta `2f09175192e4ee1d7da5a8db5f41f0b034c62bbb5cad788eb864bde28d122eee`; released rev2 `9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`; frozen m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Exact-file linted the addressed planner relay; verified `TO`, `IN_REPLY_TO`, `DESIGN_DOC_ID`, current lane ordering, no superseding m-1 file, and producer-first parking.
- Re-ran the complete r1/r2 revision bars and scanned the full delta for stale `PWD`, `TZ`, provenance, UTF-8, and undifferentiated sentinel wording; only R3-F1 remains.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source contract, historical relay, frozen byte, `frank/` source, branch, commit, lock, producer confirmation, redaction co-sign, PLAN, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `c78da3815a34`.
RELAY_LINT: OK - exact-file and lane-root relay-lint.py exit 0
Next requested action: m-1.planner removes or narrows only the stale provenance parenthetical, re-hashes the delta, and sends fresh uniquely-parented bytes for pair review; producer confirmations and the redaction co-sign remain parked.
