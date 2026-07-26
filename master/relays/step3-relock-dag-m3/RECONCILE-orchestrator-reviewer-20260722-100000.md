## RECONCILE -- HUMAN DECISION REQUIRED: D1 is a technically sound v2 direction but changes two operator-ratified Tier-HARD schema identities; D2 is not mechanically closed; D4 must consume the live m-8 lane rather than redispatch an already-selected carrier choice

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-vp-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator ratification of an exact reviewed amendment that changes the ratified m-3 schema-version literals
GRILL_REQUIRED: no -- the technical direction is bounded; the gate is ratification authority over already-ratified interface bytes
DESIGN_DOC_ID: step3-relock-dag-m3
IN_REPLY_TO: master/relays/step3-relock-dag-m3/RECONCILE-orchestrator-planner-20260722-073000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-10.planner
SUBJECT: HUMAN DECISION REQUIRED -- route a narrow exact-byte amendment covering both m-3 v1 literals to VP review then operator ratification; make predicates 2/5 truly non-gating or bind their exact gate; consume the current m-8 carrier result instead of reopening its choice

VERDICT: human-decision-required

Review target: `master/relays/step3-relock-dag-m3/RECONCILE-orchestrator-planner-20260722-073000.md` at SHA-256 `951e36ab59639f4f6322584358b0a093b6c3234272045c62f194261771be2dde`.

## Findings

### M3-VP-R1-F1 -- HUMAN GATE / BLOCKER: D1 moves two explicit ratified Tier-HARD schema identities, not one realization-only literal

The technical diagnosis is correct: frozen m-3 r4's E3 evaluator accepts only `m3.e3_observation.v1`, applies closed parsing, and rejects unknown fields (`r4:195-205`). Adding the B field to v1 would weaken version identity and the closed-parser absorb-refusal. A new version with per-version closed matrices and explicit dispatch is the coherent technical direction.

The authority classification is not correct. Ratified Stage-6 amendment rev12 labels B-E as **Tier-HARD** and explicitly fixes **both** carriers as `m3.app_event.v1` and `m3.e3_observation.v1` (`STEP-3-STAGE6-AMENDMENT.md:110-114`). Moving either literal changes the ratified interface identity. Adding evaluator version dispatch also changes the ratified consumer behavior. Master may recommend and author that amendment, but cannot convert an explicit operator-ratified literal into an unratified "realization erratum" by declaration.

D1 is also incomplete on its own terms. The must-revised r0 defines **both** `m3.app_event.v2` and `m3.e3_observation.v2` (`stage6-lane2-e0-e3-delta.md:11-15`), while target D1 disposes only the E3 literal and evaluator. E0 is also a closed schema in frozen r4 (`r4:124-151`) and is the carrier for both `frozen_core_digest` and `logical_surface_digest`; its v1-to-v2 identity and producer/consumer version handling cannot remain implicit.

Required route:

1. Author one narrow amendment/erratum artifact against exact Stage-6 rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` that covers **both** schema families: `m3.app_event.v1 -> v2` and `m3.e3_observation.v1 -> v2`.
2. Specify exact per-version closed field matrices and dispatch: v1 stays closed and unchanged; v2 carries only the authorized additive fields; unknown versions and cross-version field mixtures fail closed. Name the E0 producer/consumer version contract and the E3 evaluator's version dispatch.
3. Route those exact amendment bytes for VP review, then obtain an operator-authored ratification naming the approved hash. Master does not self-ratify.
4. Only after ratification may m-3 fold the v2 choice into r1. The recommended v2 direction need not be re-debated unless the operator rejects it.

### M3-VP-R1-F2 -- BLOCKER: "required to exist and pass" is gating, even when labeled non-gating

D2 correctly retracts rev2's false claim that all five typed predicates feed the Section 7 legs. Ratified `xit-gov-1` names predicates 1, 3, and 4 only (`STEP-3-STAGE6-AMENDMENT.md:363-375`); the pair must not invent another leg.

The proposed replacement is still not decidable. It calls predicates 2 and 5 "NON-GATING" while also saying their records are **REQUIRED evidence that must exist and pass**. If a missing, `fail`, or `unknown` record prevents any lock, release, or Step-3 exit, it is a gate regardless of its label. If none of those outcomes affects a gate, "must pass" is false and the records are diagnostic/supporting evidence only.

Required repair: name the exact consumer and consequence for each predicate.

- **Strict non-gating branch:** both contracts still exist because Section 5-E fixes the five-id set, but their verdicts are recorded/reported only; `fail` or `unknown` does not fail or hold any Section 7 leg or Step-3 exit condition. Remove "must pass."
- **Required-proof branch:** bind each predicate to an exact existing leg/composite condition and fixture-manifest consequence. That changes the ratified six-leg gate/proof contract and must ride the same reviewed, operator-ratified amendment as F1. Predicate 2 has an antecedent in the MVP deny-zero-send proof set; predicate 5 still needs an explicit ratified consumer.

Do not leave a mandatory seventh condition hidden outside the six-leg table.

### M3-VP-R1-F3 -- REVISE / STALE ROUTE: D4's producer choice already exists in the live m-8 lane

D4 identifies the correct owner and the original carrier defect, but the proposed next action is stale against current bytes. m-8 r2 `d482bc391570267a9b7e6ab99c7e75851b1becd19e44c0f2b7673cad2aa88772` has already selected and specified the carrier result:

- pre-freeze rejects carry no B/E;
- post-freeze `internal_integrity_fault` carries B/E on `internal_integrity_reject.v2` and `m8.attempt_result.v2`;
- policy deny carries B/E on `egress_denied.v2` and the attempt result;
- no-carrier epoch/loss/crash cuts remain message-absent.

The later m-8 implementer review at SHA-256 `0991d9e780baaa6ca861af50ac794577b226ac5b7596df11b803fc13eb3f4fea` **accepts that principal freeze-boundary/carrier correction** (`083000:59-64`) while returning must-revise on transport-observer applicability, byte-decodable carrier versions, and absent-tools validation. Producer confirmation and consumer wakeup remain held.

Therefore do not issue a new open binary asking m-8 to choose "carry versus exclude." Replace D4 with a dependency disposition: m-8 folds its current three review findings while preserving the accepted deny/post-freeze-reject carrier matrix; after exact r3 pair approval, master routes that final producer hash to m-3. m-3 then authors its exhaustive cut matrix against those settled bytes. The m-8 review's carrier-version finding is directly relevant to F1 and must be closed before m-3 claims version-compatible carriage.

## Passed disposition

### D3 -- APPROVE

Branch (a) is the correct master clarification of the rev2/release conflict: m-3 may author the E0 schema grain now, while recipe/binding confirmation remains parked until exact pair-approved m-9/m-8 producer bytes exist. This does not unpark the B sink or E two-digest join. Carry D3 unchanged into the corrected disposition; no renewed review of D3 is required absent byte movement.

The target also correctly preserves r0's must-revise status, frozen r4, the parked sink/join, and all downstream gates.

## Gate disposition

- Target D1-D4 package: **HUMAN DECISION REQUIRED + REVISE BEFORE RELEASE**.
- Do not issue the target's D1-D3 disposition to m-3 and do not issue a duplicate D4 choice to m-8.
- Master next authors the exact two-schema amendment, closes D2's consumer/consequence ambiguity, and updates D4 to consume the live m-8 review path. Route the amendment to VP review and then to the operator for hash-bound ratification.
- m-3 r0 `dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e` remains must-revised; m-3 r1 stays held.
- m-8 r2 remains must-revised under `083000`; no producer confirmation or m-3/m-9/m-10 consumer wakeup may advance on it.
- DESIGN-lock, PLAN, T4/code, credential, provider, release-binding, E3, merge, and deploy gates remain held.

## Verification

- Reproduced from current disk: target `951e36ab...`; m-3 r0 `dc3b6eb3...`; pair review `abbeaab6...`; escalation `a1972e44...`; Stage-6 rev12 `1125b0a0...`; frozen m-3 r4 `009df607...`; m-8 r2 `d482bc39...`; m-8 review `0991d9e7...`.
- Target, m-3 pair review/escalation/rev2/release, and current m-8 review exact-file lint: OK.
- Live INDEX state checked beyond the target: m-8's `083000` must-revise is current boundary evidence and makes D4's proposed open-choice dispatch stale.
- `frank/` remains clean on `main` at `c78da3815a34480590071295c1e09bb7d53c10b6`, equal to `origin/main`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, disposition release, pair dispatch, design byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update.
Next requested action: master authors the narrow two-schema amendment + exact D2 consequence + current-state D4 dependency, returns the amendment for VP exact-byte review, and then routes the approved hash to the operator for ratification. D3 carries unchanged; all pair and downstream action remains held meanwhile.
