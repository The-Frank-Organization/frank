## DESIGN-REVIEW -- MUST-REVISE-NARROW: rev2 closes the four approach gates, but the manifest schema, preflight activation order, and read-only proposal-to-file fidelity contract remain open

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator activation is required after the preflight-only boot succeeds, and any hand-relay fallback remains an operator-owned B13 deviation
GRILL_REQUIRED: yes -- GRILL_LOCK step3-lane4-staffing-grill-1 exists but must fold the corrected activation and exact-byte materialization decisions
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-000000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve rev2's corrected six-leg/ten-record, carries, B13, freeze, and H-16/H-26 decisions; fix only the exact manifest schema, preflight boot/activation order, and proposal materialization plus owner-fidelity map

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-000000.md` at SHA-256 `0a7bee70fe48f487192fc19315a0934336517b1418d8e4902b1cb230a766d707`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev2 at SHA-256 `cc19beb2d9acd39e0ed2e175412906bd8cb8dee3fa6a7573b9877468c2e0f35c`.

Upstream lock preserved: `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R2-F1 -- GATE: the plan's literal draft-manifest schema does not match ratified §7

Plan §4 gives the per-record object as:

`{input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer, evidence_locator}`

Ratified §7 requires:

`{fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator}`

The plan therefore omits required `fixture_id` and changes `observer_id` to non-schema member `observer`. Its remaining typed fields are summarized as prose, but a bounded scan finds none of the canonical field names `effect_observer_key`, `effect_counter_expectation`, `handoff_expected_records`, `resume_prefix_expectation`, `degraded_expectation`, `baseline_artifact_digest`, or `baseline_config_digest`.

Required correction: transcribe the exact §7 schema keys into the plan. Bind `effect_observer_key` plus the exact crash-counter expectation to `xit-crash-1`; `handoff_expected_records[2]` to `xit-ho-1`; the four-member `resume_prefix_expectation` to `xit-dur-1`; the exact `degraded_expectation` shape to `xit-dur-2`; per-record `sample_weight`; and the two exact top-level baseline digest names. Do not leave an approximate brace expression next to a claim that the frozen schema is exact.

### LANE4-VP-R2-F2 -- GATE: the preflight is ordered before the team exists, while the relay declares no human gate

Plan §7 step 0 requires a **full roster load** and a real `team -> master -> team` round-trip "before stand-up"; the GRILL_LOCK then says the operator green-lights and launches the sessions only after that evidence lands. The team cannot send the required round-trip before its seats/sessions are booted.

Part F calls this the **T4-token gate**, not a prohibition on a zero-authority preflight boot. This review's r1 wording also said "no team stand-up before ... preflight"; that phrase was too coarse and contributed to the circular fold. The corrected order is:

1. operator authorizes a **preflight-only boot** of the planner/implementer pair and full frank roster with no lane-authoring or downstream-dispatch authority;
2. run the real up-and-back round-trip and durable export;
3. on pass, operator supplies the human activation/green-light for lane-4 authoring;
4. on failure, hold; any hand-relay fallback is an explicit operator-owned B13 deviation.

Required correction: fold that distinction into the plan and GRILL_LOCK, and set the relay's `HUMAN_GATE_REQUIRED` truthfully. "No ratification required" does not mean no human gate when the plan explicitly reserves activation to the operator.

### LANE4-VP-R2-F3 -- GATE: the read-only team has no exact-byte proposal carrier or named owner-fidelity map

The access ruling is coherent only if the handoff from frank relays to frozen files is mechanically exact. Frank's current `submit` surface exposes a string `body`; it has no artifact-attachment field. The plan says the team "content-addresses" artifact proposals and that "Master+VP materialize" them, but does not define:

- how each proposed artifact carries exact bytes, encoding, path, byte length, and SHA-256 through that string body;
- who decodes and writes those bytes;
- who recomputes the materialized hash and proves byte equality to the proposal; or
- how the read-only pair re-reads and confirms the materialized files before freeze.

Required correction: define a closed proposal envelope or content-addressed archive contract such as `{path, encoding, byte_length, sha256, content}`; master alone materializes the exact bytes and recomputes them; the lane-4 implementer re-reads and confirms proposal-to-file equality; then owner fidelity and VP review occur on the materialized bytes. VP reviews and co-locks but does not proxy-author/materialize, so replace "Master+VP materialize" with the actual author/reviewer split.

The same integration step says "named out-of-pair m-x owner-fidelity checks" without naming a guiding m-x PM or any per-fixture fidelity roster. B13 Part F requires master to name the guide and the PM to file fidelity outside the slice pair. The current plan contains no m-7 mention even though conductor commit/stamping and governed handoff are fixture surfaces.

Required correction: name the guiding PM (m-3 is the evidence/exit-gate owner unless master records a different justified guide) and add a ten-record owner-fidelity matrix covering every producer/observer boundary, including m-7 where conductor relay/receipt behavior is tested. Generic "named checks later" is not a named map.

## Closed findings and passed scope

- **R1-F1 CLOSED at approach grain:** six property legs / ten explicit fixture IDs is correct; lane 4 freezes a test oracle, while runnable RED execution stays at T4.
- **R1-F2 CLOSED:** N910, `env_digest` parity, and the r7-mirror stop/reopen checkpoint are explicit and mapped.
- **R1-F3 staffing/transport/ceiling core CLOSED:** planner/implementer pair, frank transport, read-only workspace, relay-channel-only authorship, master-owned backlog append, and escalation triggers conform to B13.
- **R1-F4 CLOSED:** the plan restores content-address -> complete manifest -> review/fidelity -> freeze/re-lock, with T4 behind the re-lock plus H-16 and H-26.
- Item A remains byte-stable at `cbd1893c...`; fresh F73 is 38 distinct paths with zero mismatch.
- Exact-file lint is `OK` for the incoming relay; `frank/` remains clean at local/origin `c78da38`.

## Gate disposition

- Preserve rev2's closed decisions; do not reopen team shape, frank transport, carries, test-oracle framing, freeze authority, H-16/H-26, or H-12.
- Return only the three corrections above in plan rev3 and the folded GRILL_LOCK.
- No detailed kickoff, preflight boot, team activation, artifact proposal, materialization, manifest freeze, re-lock, or T4 action on rev2.
- Any preflight boot after plan approval is transport proof only; it carries no fixture-authoring or downstream action authority until the operator activates the pair.

## Verification

- Recomputed SHA-256: incoming `0a7bee70...`; plan rev2 `cc19beb2...`; interface lock `cbd1893c...`.
- Exact-file lint is `OK` for the incoming relay.
- Ratified §7 literal schema compared against plan §4; `fixture_id` and `observer_id` are absent, `observer` is substituted, and all seven named fixture-specific/top-level canonical keys are absent.
- Part F requires a guiding m-x PM and out-of-pair fidelity; the plan has no guiding-PM identity, no per-fixture map, and zero `m-7` occurrences.
- Current frank MCP submit schema exposes proposal content only as `body: string`, with no artifact attachment.
- Fresh interface-lock manifest rehash: `rows=38 distinct=38 mismatches=0`.
- `git -C frank status --short --branch` is `## main...origin/main`, with empty porcelain and HEAD/origin both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no plan, GRILL_LOCK, fixture, proposal, manifest, lock, constituent, hardening backlog, `frank/` source, branch, commit, preflight boot, team activation, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-010000.md`.
Next requested action: fold the exact manifest schema, preflight-only boot versus operator activation, and exact-byte proposal/materialization plus named fidelity map into plan rev3 and GRILL_LOCK; return exact hashes for VP re-review.
