## DESIGN-REVIEW -- MUST-REVISE lane-4 approach: separate-team authoring is sound, but the plan misstates the frozen deliverable, drops three carried obligations, conflicts with B13 staffing/transport, and skips pre-T4 gates

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- only if master proposes to deviate from B13's settled planner/implementer team and frank transport; no team stand-up before the durable grill and transport preflight
GRILL_REQUIRED: yes -- this is large, new, cross-domain work with explicit open staffing/transport/authority decisions and a hard-to-reverse frozen oracle
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260727-220000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve Item A cbd1893c and the separate-team/Master+VP split; revise the lane-4 plan to bind ten fixture records, three carried obligations, B13 pair+frank mechanics, owner fidelity, the exact freeze order, and H-16/H-26

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260727-220000.md` at SHA-256 `0b9532ba1d0cbe249e05abf426275ed2c00268bff40dfafe3e77d2b7fd32ed8d`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` at SHA-256 `d79c44c12c857879cf835bb0beddacd6c2f8cae5dd661679c1484d5af2da8f06`.

Upstream lock preserved: `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R1-F1 -- GATE: six property legs are not six fixtures, and lane 4 freezes an oracle rather than runnable red tests

The plan repeatedly says "six governance-property fixtures", "six fixtures' input scenarios", and "authors the six fixture inputs." The ratified §7 has **six legs but ten fixture IDs**:

`xit-gov-1`, `xit-dur-1`, `xit-dur-2`, `xit-dur-3`, `xit-dur-4`, `xit-dur-5`, `xit-crash-1`, `xit-inj-1`, `xit-ho-1`, `xit-op-1`.

That count error is deliverable-shaped: a team following the prose can under-produce while still believing it completed six fixtures.

The same sections call lane 4 a runnable "TDD red phase" and say it writes "failing tests." Ratified §7 and simplification-amendment §4 instead require lane 4 to author and content-address immutable **input/baseline artifacts**, then freeze a manifest containing the expected canonical oracle. The executable fixtures are **built at T4**. A frozen test oracle before implementation is valid test-first design, but it is not demonstrated RED execution. If lane 4 is intended to create runnable failing tests, that changes the §7/§11 split and needs an amendment through the existing gate.

Required correction: use "six legs / ten fixture records" throughout; enumerate all ten records in the deliverable and sequence; describe lane 4 as the content-addressed frozen oracle/test specification. Reserve RED execution and executable fixture code for T4 unless master explicitly routes an amendment.

### LANE4-VP-R1-F2 -- GATE: §7 alone is not the complete lane-4 spec; all three carried obligations are absent

Plan §4 says the already-frozen §7 spec is complete, but the later lane-2 close plus the ratified simplification amendment add three operative carries that lane 4 alone must realize as fixture records/expected rows:

1. **N910 documented-MVP-limit cut:** preserve the honest no-`b_sink`/`UNKNOWN_PROVIDER_OUTCOME` to `uncertain` disposition in the appropriate durability/operability oracle; never claim complete lane-2 coverage.
2. **`env_digest` preimage parity:** realize m-1's byte-exact JCS preimage recipe, duplicate-name rejection, reachable non-UTF-8 typed rejection, and m-9-versus-m-3 observer parity under `xit-gov-1`.
3. **r7-mirror deferred-v3 caveat:** at E3-predicate authoring, perform the mandatory m-3 check for whether `xit-gov-1` needs independent m-10-side 2a/2b resolution; if yes, stop and reopen route-now, otherwise record the non-gating deferred disposition.

The generic escalation paragraph does not discharge a known named checkpoint. Required correction: add a closed `carried_obligations` deliverable/checklist sourced from the lock's carried lineage and the lane-2 close; map each carry to its fixture record, expected rows, owner, and stop/reopen behavior.

### LANE4-VP-R1-F3 -- GATE: the proposed staffing/transport choices contradict B13, and the full-access team has no pinned authority/fidelity envelope

The plan presents "single orchestrator vs planner/implementer" and "frank vs operator hand-relay" as ordinary open choices. The adopted B13/`CYCLE-PLAYBOOK.md` Part F contract already fixes:

- a slice team as its own `sN.planner` + `sN.implementer`;
- relay transport through frank, not operator hand-relay;
- out-of-pair m-x owner-fidelity before close; and
- a T4-token preflight of full frank roster + escalation round-trip + durable export.

Master may propose a deviation, but it is then a durable operator decision/amendment, not an unmarked plan option. The current `HUMAN_GATE_REQUIRED: no` and `GRILL_REQUIRED: no` are therefore wrong for this large cross-domain design with open team/access/transport semantics. The plan also says "full workspace access" without separating environment visibility from write authority, while Part F explicitly leaves the first team's authority ceiling and nested lineage to be pinned.

Required correction:

- conform to a planner/implementer pair on frank, or route an explicit operator-owned B13 deviation before stand-up;
- set `GRILL_REQUIRED: yes` and produce a durable GRILL_LOCK resolving team/access/transport/authority, answering settled questions from B13 rather than re-asking them;
- require the frank preflight evidence before stand-up: full roster, one real up-and-back escalation round-trip, and durable relay export;
- define an exact write fence for the team's draft input artifacts, baselines, manifest, own relays, and gap reports; read-all is not write-all, and no lock constituent, ratified contract, or final lock authority is delegated;
- name affected m-x PM fidelity checks outside the lane-4 pair before Master+VP freeze; and
- state who appends artifact-gap battle reports to the master-owned hardening backlog.

### LANE4-VP-R1-F4 -- GATE: the sequence jumps from lane-4 lock to T4 without H-16/H-26 and under-specifies the freeze gate

Plan §7 step 5 says lane 5/T4 builds immediately after the re-lock. Ratified rev12 §9 and §11 require **both H-16 and H-26 before T4**, in addition to the lane-4 lock.

The plan also compresses the exact ratified freeze order. The revised sequence must state:

1. author and content-address all ten immutable fixture-input artifacts and baseline artifacts;
2. assemble the complete manifest with final non-placeholder input/baseline digests, all typed expectation fields, all carried-obligation rows, resolved observers/evidence locators, and concrete per-record weights totaling exactly 30 governed turns + 100 tool calls;
3. run independent team review plus the named out-of-pair owner-fidelity checks;
4. Master+VP freeze `STEP-3-EXIT-FIXTURES.json` and issue one durable re-lock binding both full interface-lock SHA `cbd1893c...` and the frozen manifest's full SHA; and
5. open T4 only after that re-lock **and** H-16/H-26. H-12 continues to block external use.

No placeholder, mutable slot, unresolved owner, or arithmetic-only weight promise may survive the freeze.

## Passed scope

- B20 correctly fired before any kickoff, stand-up, fixture authoring, or build action.
- Item A is correctly treated as closed at exact hash `cbd1893c...`; a fresh F73 pass still reports 38 distinct paths and zero mismatches.
- Separate-team authorship is a sound independence boundary, and retaining freeze/re-lock authority with Master+VP is correct.
- Gap logging is useful instrumentation once ownership of the backlog append is explicit.
- The plan correctly grants no fixture, manifest, lock, T4, implementation, merge, provider, credential, E3, deploy, or external-use authority.

## Gate disposition

- Do not write the detailed kickoff, stand up the team, author inputs, create the manifest, freeze, re-lock, or open T4 on this revision.
- Preserve interface lock `cbd1893c...` and all 38 named constituent files byte-for-byte.
- Return a revised high-level plan plus its GRILL_LOCK. If master conforms to B13 pair+frank, no new operator choice is needed for those settled points; any deviation requires a durable operator gate.
- Lane 4 remains at approach review. Lane 5/T4 remains held behind the corrected lane-4 plan, completed lane-4 lock, H-16, and H-26.

## Verification

- Recomputed SHA-256: incoming `0b9532ba...`; plan `d79c44c1...`; interface lock `cbd1893c...`.
- Exact-file lint is `OK` for the incoming relay.
- Ratified §7 fixture census returns ten distinct fixture IDs across six property legs.
- A bounded plan scan finds no `N910`, `env_digest`, `r7-mirror`, `H-16`, `H-26`, owner-fidelity, authority-ceiling, frank-roster, durable-export, or GRILL_LOCK term.
- B13/Part F directly fixes pair staffing, frank transport, owner-fidelity, and the roster/round-trip/export preflight.
- Fresh interface-lock manifest rehash: `rows=38 distinct=38 mismatches=0`.
- `git -C frank status --short --branch` is `## main...origin/main`, with empty porcelain and HEAD/origin both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no plan, deviation, GRILL_LOCK, fixture, manifest, lock, constituent, hardening backlog, `frank/` source, branch, commit, team stand-up, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260727-230000.md`.
Next requested action: revise the lane-4 high-level plan at the four gates above, produce the durable GRILL_LOCK, and return exact hashes for VP re-review. No kickoff, stand-up, fixture authoring, freeze, lock, or T4 action before approval.
