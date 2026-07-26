## DESIGN-REVIEW -- MUST-REVISE-NARROW: rev3 closes the literal schema and activation core, but the manifest byte path, kickoff order, and owner-fidelity rows are not yet complete

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r3
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator still authorizes the preflight-only boot and separately activates lane authoring after the round-trip passes
GRILL_REQUIRED: yes -- GRILL_LOCK step3-lane4-staffing-grill-1 exists and its core decisions pass, but its exact-byte delivery and activation sequence must fold the corrections below
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-020000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve rev3's exact schema and preflight activation fixes; revise only the manifest proposal/materialization path, kickoff ordering, and the owner-real fidelity matrix

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-020000.md` at SHA-256 `e0691b0b56a36d0934dfd9160bd4c056c81a057a31c096d70f91f97d2a38574d`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev3 at SHA-256 `7698cb018340b4d17d7bc74a7bd5553520887934ac81d6644123e9414cc5aa04`.

Upstream lock preserved: `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R3-F1 -- GATE: the read-only pair's final manifest bypasses the exact-byte proposal/materialization chain

Plan Section 3 correctly says each proposed artifact rides `{path, encoding, byte_length, sha256, content}`, master materializes, the lane implementer confirms equality, and only then owners plus VP review the materialized bytes. The numbered sequence does not apply that chain to the final lock target:

- step 2 sends the ten fixture-input and baseline artifacts;
- step 3 says master materializes, then the team "assembles the complete manifest";
- step 4 begins review; and
- step 5 freezes `STEP-3-EXIT-FIXTURES.json`.

The team is read-only, so "assembles" cannot create that governed-tree file. No later step has the team emit the complete manifest as an envelope, master materialize it, master recompute its on-disk length/hash, or the implementer confirm manifest proposal-to-file equality. That leaves the exact bytes Master+VP would freeze without the authorship/fidelity proof Section 3 requires.

Required correction: make the closed chain apply to **every proposed file, including the complete final `STEP-3-EXIT-FIXTURES.json`**:

1. pair emits its final manifest envelope after all final artifact digests, typed expectations, carried rows, observers/locators, and weights are resolved;
2. master alone materializes it and recomputes on-disk byte length plus SHA-256;
3. the lane implementer re-reads and confirms exact proposal-to-file equality;
4. only then do owner-fidelity and VP review/freeze occur.

Frank's configured frame ceiling also bounds a one-envelope-per-artifact carrier (`EngineConfig.FrameBytes()` defaults to 1 MiB). The detailed kickoff must either require every encoded proposal frame to fit the live `max_frame_bytes` with overhead, or define a deterministic content-addressed archive/chunk contract before authoring; an oversized artifact must HOLD rather than be truncated or hand-copied.

### LANE4-VP-R3-F2 -- GATE: the ten-record matrix is named but not owner-real at four tested boundaries

The matrix now names every record and includes m-7, but "every producer/observer boundary" requires the actual frozen owner, not merely a populated row.

- **`xit-gov-1`: m-9 is absent and m-10 is labeled for invocation.** Ratified Section 7 gates `local_invocation_matches_effect_descriptor`; m-3's predicate consumes the descriptor m-9 derives/records plus the executor-captured invocation, and m-9 Section 7 explicitly owns that half. The plan itself also requires m-9-to-m-3 `env_digest` parity.
- **`xit-inj-1`: m-9 is absent.** The induced relay action is parsed, derived, executed, and recorded by m-9 under the F59 ticket m-10 issues; m-7/m-1 cover the stamped commit, not the executor half.
- **`xit-ho-1`: m-2 is absent.** The predicate requires the response/ack's parent/reference lineage to record 1; m-2 owns the record schema and cross-relay lineage engine. m-7/m-1 stamping alone cannot fidelity-check that edge.
- **`xit-op-1`: m-8 is absent even though the plan binds N910 to this record.** The N910 scenario begins at m-8's DATA-P/connector-loss boundary, then survives as m-10 `UNKNOWN_PROVIDER_OUTCOME`/`uncertain` on the operator surface.

The `xit-dur-4` row also assigns "work-ordering vs receipt" to m-7. Frozen m-9 Section 3 says m-9 waits for the post-commit receipt before any provider attempt, tool effect, or conductor verb; m-10 owns the durable disposition/receipt. If the fixture selects the conductor branch, m-7 owns the selected conductor-action observation, not the receipt gate itself. If it instead selects the provider branch, m-8 becomes required.

Minimal owner-real rows:

```text
| `xit-gov-1` | m-8 (frozen-core/provider request) . m-10 (ticket-bound descriptor) . m-9 (descriptor derivation + captured invocation + `env_digest` parity) . m-7 (relay commit/stamp execution) . m-1 (`env_digest` recipe + stamp contract) . m-3 (E3 predicates) |
| `xit-dur-4` | m-9 (receipt-gated no-work ordering) . m-10 (durable disposition + post-commit receipt) . m-7 (selected conductor-action observation) . m-3 |
| `xit-inj-1` | m-10 (F59 ticket) . m-9 (parsed induced relay call + descriptor/invocation record) . m-7 (`FROM`-stamped commit) . m-1 (stamp contract) . m-3 (honest-outcome observer) |
| `xit-ho-1` | m-7 (channel-stamped `FROM` + conductor E1/E2) . m-1 (identity/stamp) . m-2 (record schema + parent/reference lineage gate) . m-3 |
| `xit-op-1` | m-8 (connector/provider-loss boundary for N910) . m-10 (operator surface + `uncertain`) . m-9 . m-3 |
```

If `xit-dur-4` selects provider or tool instead of conductor, replace the selected-action owner accordingly; do not retain m-7 as a generic ordering owner.

### LANE4-VP-R3-F3 -- GATE: the kickoff has two incompatible positions around preflight and activation

Section 7 step 0 runs preflight boot, round-trip, and operator activation; numbered step 1 then writes the detailed kickoff. The plan boundary and incoming relay instead say: on VP approval, master writes the kickoff, then the operator authorizes preflight boot.

Either can be coherent, but both cannot be the operative sequence. Required correction: choose one literal order in the sequence, boundary, incoming summary, and GRILL_LOCK. If master writes the kickoff before preflight, state that it is inert and grants no lane action until the operator's post-pass activation. If kickoff authoring waits until preflight passes, place it before the separate operator activation. In either form, preflight-only boot remains zero-authority and no pair authoring starts before activation.

### LANE4-VP-R3-F4 -- RECORD CORRECTION: the plan has the exact schema, but the incoming summary still uses approximate/count-wrong claims

Plan Section 4 correctly carries the six base keys, all typed extension names and members, per-record weights summing to the two totals, and both top-level baseline digests. The incoming relay's change summary abbreviates the crash expectation as `{1,1,0}`, writes `sample_weight=30 turns+100 calls` rather than per-record values summing to those totals, and then says no approximate brace remains. Its verification also says "all nine exact Section 7 schema keys" although the enumerated base, typed, weight, and top-level field names exceed nine.

This does not reopen the plan's now-correct schema. The rev4 transmittal must drop the numeric key-count claim and summarize the exact named shapes without substituting tuple shorthand or a per-record total.

## Closed findings and passed scope

- **R2-F1 CLOSED in the plan:** the literal Section 7 manifest schema and fixture-specific/top-level bindings are correct.
- **R2-F2 core CLOSED:** operator-authorized preflight-only boot -> real round-trip plus durable export -> separate operator activation is correct, and `HUMAN_GATE_REQUIRED: yes` is truthful. Only the kickoff's position must be normalized.
- **R2-F3 role split CLOSED except the manifest gap:** the pair authors proposals; master alone materializes/recomputes; the implementer confirms; VP reviews/co-locks and does not proxy-author.
- Guiding PM m-3, planner/implementer pair staffing, frank transport, read-only workspace, authority ceiling, escalation route, and master-owned hardening-backlog append remain closed.
- Six legs / ten records, frozen-oracle-not-runnable-RED framing, N910/`env_digest`/r7-mirror carries, exact freeze/re-lock authority, H-16/H-26 before T4, and H-12 remain closed.
- Item A remains byte-stable at `cbd1893c...`; fresh F73 is 38 distinct paths with zero mismatch.
- Exact-file lint is `OK` for the incoming relay; `frank/` remains clean at local/origin `c78da38`.

## Gate disposition

- Preserve rev3's closed decisions. Return one bounded rev4 correcting only F1-F4 above.
- Do not write the detailed kickoff, authorize/boot the preflight roster, activate the pair, send proposal artifacts, materialize files, run fidelity, freeze the manifest, re-lock, or open T4 on rev3.
- Plan approval, when it lands, authorizes master to proceed to the normalized kickoff/preflight sequence only; the operator retains both human actions, and preflight boot carries zero authoring/downstream-dispatch authority.
- No fixture, manifest, owner/frozen contract, interface lock, `frank/` code, credential, provider call, E3 evidence, merge, deploy, or external-use action is authorized by this review.

## Verification

- Recomputed SHA-256: incoming `e0691b0b...`; plan rev3 `7698cb01...`; interface lock `cbd1893c...`.
- Exact-file lint is `OK` for the incoming relay.
- Ratified Section 7 and plan Section 4 compare clean at schema grain; the remaining schema defect is confined to the incoming relay's shorthand/count claim.
- The read-only delivery sequence was traced file-by-file: no manifest envelope/materialization/equality step exists before review/freeze.
- Frozen ownership was checked at m-3 predicate Section 2.3, m-9 lane-2 Sections 3/7, m-2 lineage ownership, and m-8 provider-contract loss row 9.
- Frank's default frame limit is `1 << 20` bytes; one string-body proposal therefore needs an explicit fit/hold or deterministic multi-part rule.
- Fresh interface-lock manifest rehash: `rows=38 distinct=38 mismatches=0`.
- `git -C frank status --short --branch` is `## main...origin/main`, with empty porcelain and HEAD/origin both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, GRILL_LOCK, kickoff, preflight, activation, proposal, fixture, manifest, lock, owner/frozen artifact, hardening backlog, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-030000.md`.
Next requested action: issue bounded plan rev4 folding the manifest exact-byte chain, one kickoff/preflight order, the owner-real matrix, and an exact transmittal summary; return exact hashes for VP re-review. No kickoff, boot, activation, proposal, fixture, freeze, lock, or T4 action before approval.
