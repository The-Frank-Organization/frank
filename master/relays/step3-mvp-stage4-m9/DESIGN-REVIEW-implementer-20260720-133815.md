## DESIGN-REVIEW — MUST-REVISE m-9 stage-4 full-worker r2 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-mvp-stage4-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator's six GRILL_LOCK decisions stand; the findings are m-9-owned completeness and frozen-basis conformance corrections
GRILL_REQUIRED: yes — satisfied by `GRILL_LOCK_ID: step3-mvp-stage4-m9-worker-grill-1`; no operator choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m9-worker
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: e666025792a5019514751cc178cd8cccd0387380fddf8b6b0c1ab3603f49c138
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage4-m9/DESIGN-planner-20260720-061500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-133815.md
SUBJECT: MUST-REVISE exact r2 e6660257 — H-17 rows are not canonical or effect-complete; transcript durability contradicts the frozen no-second-truth boundary; the m-9-owned F58 catalog build is absent; and the E0 phase population is not a total executable mapping

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r2 design bytes at SHA-256 `e666025792a5019514751cc178cd8cccd0387380fddf8b6b0c1ab3603f49c138`. The frozen bases, operator GRILL_LOCK, uniform F59 gate, three-identity guard, sync-authorize/async-record discipline, bash honesty, precise process-group teardown, governed compaction direction, L5=B, and F63 worker-artifact direction pass. Four m-9-owned blockers prevent final-byte approval.

## Findings

### M9-S4-R2-F1 — BLOCKER: the H-17 census is neither canonical-row assemblable nor effect-complete

The canonical v1 schema at `ea173abc…` requires every effect row to carry the same 20 fields, with requester/executor, decision/enforcement, authorization/effect linearization, and reporter/observer/validator/record kept distinct. §11.5 instead contains shorthand and collapsed cells:

- E2 says “As E1 except” and replaces the required cells with `authorization/effect/record`; E3 collapses requester/executor/authority/decision/enforcement/freeze/linearization/reporter/observer/validator/record into one inherited clause. Neither is a standalone canonical row.
- E7 omits `policy_owner/policy_artifact` and collapses `decision_point/enforcement_point` plus `authorization/effect linearization`.
- E9–E14 repeatedly collapse required fields (`requester/executor`, `decision_point/enforcement_point`, or `authorization/effect linearization`) and omit required cells. Concrete omissions include policy owner/artifact in E9–E13, the E11 credential-holder and authorization-linearization cells, and E13's request-freeze and authorization-linearization cells. The schema permits exact null tokens; it does not permit absent cells or inherited prose.
- E14 combines two authoritative families with different boundaries: conductor→seat push receipt/forwarding and worker→m-10 `wake_forward` schedule insertion. Its requester, enforcement/effect points, and loss semantics describe the schedule half while merely mentioning the inbound push. F87 required push as its own read/serve effect; split the families into canonical rows or provide one explicit non-effect rationale and one complete effect row without erasing either boundary.
- The live inventory is E1–E14 plus handshake and compaction rationales, but §13 still says “13 effect rows + the compaction non-effect rationale.” That contradicts the status, incoming relay, and live §11.5 inventory.

Required correction: emit every effect as a self-contained canonical v1 row with all 20 labels (using only `unknown`, `not specified`, or `residual` where an owner contract is silent), separate the push/wake families honestly, and make all live inventory counts agree. Re-run the full-family coverage scan after expansion; master must not have to infer cells at stage 6.

### M9-S4-R2-F2 — BLOCKER: transcript/compaction state creates an unnamed second durable truth

§2.1 sources `input[]` from “the durable turn transcript,” and §7.1 says the worker “holds the durable turn transcript” and reloads it each iteration. Frozen lifecycle r19 §2.6 states the opposite boundary: the worker owns no durable canonical session store; canonical facts live in its conductor records and m-10's private frame-only rows, and stage 4 may add only a derived-with-digest transcript without introducing a second canonical truth. A worker cannot reconstruct m-10's private rows; r19 names only its own conductor audit projection plus fresh admission as its readable recovery surface.

R2 names no canonical source set, writer, storage owner, derivation digest, reload rule, or crash/replacement behavior for this “durable” transcript. The §11.5 compaction rationale then calls Tier-1/Tier-2 purely in-memory transforms with no durable row, leaving unclear whether a reload discards the summary, repeats a nondeterministic provider summary, or reads a second store. Tier-0 is said to reload “authoritatively,” but its exact source and deterministic equality representation are also absent.

Required correction: state the canonical readable sources, derived-only representation and digest, writer/storage locus or explicitly ephemeral reconstruction rule, replacement/crash semantics, and the exact Tier-0 comparison input. Preserve r19's no-second-truth rule. If any durable transcript/summary effect remains, give it a canonical H-17 row; otherwise make the non-effect rationale exact about derivation and reload.

### M9-S4-R2-F3 — BLOCKER: the stage-4-owned F58 tool catalog build is missing

The ratified amendment §4/F58 requires the stage-6 interface lock to bind an expected per-tool vector `{canonical name, tool-schema digest, tool-implementation/catalog version, form→schema-mapping version}`. It assigns local-tool schema digests and the catalog build/version to m-9, requires the owner designs to define canonical encoding and field applicability, makes local mapping-version absent, and later verifies `tool_catalog_digest` mechanically against the registry shipped in the m-9 artifact. Lifecycle r19 §4 explicitly deferred the m-9 catalog build to this stage-4 full-worker design.

R2 defines the eight names and later mentions `tool_catalog_digest` only as an F63 output. It does not define the five local-tool schemas/digests, catalog-version producer, canonical identity-vector serialization/sort/digest, relay-versus-local field applicability, expected pre-build catalog vector, or the shipped-registry derivation/verification rule. Stage 6 therefore has no m-9 interface identity contract to lock and T4 would have to invent it.

Required correction: add the m-9-owned catalog-build contract at design grain, consuming m-2's relay-tool vectors without re-owning them, and bind its expected vector/digest and post-build registry verification into F63 consistently.

### M9-S4-R2-F4 — BLOCKER: E0 population is asserted total but is not an executable total mapping

§6.1 says “per attempt outcome” and claims a total mapping over m-8 dispositions, but supplies only the enum plus the special cancellation rule. It never fixes the complete disposition/event-boundary table and gives no emission point for `phase=sent`, which is not a terminal attempt outcome. The frozen m-8 r12 table already distinguishes policy deny, deterministic local reject, completed, observed failure, both cancellation cuts, live-observer loss with conditional E0 delivery, worker crash with no E0, and attempt-inert epoch rejects with no E0 from that path. Those distinctions must survive producer realization; a generic “mapped total” sentence is not buildable and risks fabricating an event on the no-emission cuts.

Required correction: pin the exact event triggers and value mapping, including when `sent` is emitted, `deny_reason` presence, no-E0 epoch-inert and worker-crash cuts, and the live-loss delivery race. Reflect the same semantics in E12 and the fixtures, then route the complete table for m-3/m-8 confirmation.

## Accepted basis retained

- All seven consumed owner hashes and H-17 schema hash reproduce exactly.
- The operator GRILL_LOCK is durable and complete: all six decisions are reflected; none requires a human re-open.
- The eight-tool uniform F59 gate, identities #1/#2/#3 at points A/B, exact outcome branches, and crash-gap `UNKNOWN` posture conform to r19 × r36.
- Bash is labeled honestly: authorization/invocation is recorded, host effect is not claimed contained; `setpgid`/`killpg`, bounded direct-child reap, the `setsid` escape residual, and Step-4 H-12 deferral match frozen r19.
- The compaction direction, summary call through a fresh ordinary m-8 attempt, no opaque provider capsule, L5=B materials reference, and F63 worker-own-artifact direction are accepted in substance. F2/F3 require the missing executable state and identity contracts, not a new operator choice.

This verdict is byte-bound to `e666025792a5019514751cc178cd8cccd0387380fddf8b6b0c1ab3603f49c138`. A corrected design requires a fresh uniquely-parented m-9.implementer DESIGN-REVIEW. No SITREP, consumer-confirmation routing, stage-6 lock, PLAN, T4 token, implementation, release binding, E3, merge, or deploy may consume r2 as approved.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `0bbccd95060f0d71c602f8df457998ccf265c4c32a882f72e278658fbfcc7500`.
- Exact reviewed m-9 r2 SHA-256 recomputed: `e666025792a5019514751cc178cd8cccd0387380fddf8b6b0c1ab3603f49c138`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen owner hashes recomputed: m-9 r19 `2a96a07b…`; m-10 r36 `0240e874…`; m-8 r12 `4b670a79…`; m-7 r11 `9331ea88…`; m-3 r4 `009df607…`; m-2 `83d8e63e…`; m-1 `7c8b09a6…`.
- Incoming DESIGN exact-file lint: OK.
- Full-byte pass: §§0–13, all 14 census rows, both rationales, GRILL_LOCK, fold log, dispatch, corrected H-17 supplement, frozen r19/r36/m-8/m-3 seams, F58/F63 amendment, and live index/duplicate-response scan.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte must-revise relay and appended one `master/relays/INDEX.md` row; no design/source doc, historical relay, frozen contract, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file `relay-lint.py` verification on this relay.
Next requested action: m-9.planner folds M9-S4-R2-F1..F4 into one corrected revision, rechecks every frozen hash, and returns a fresh uniquely-parented full-byte DESIGN relay; later gates remain held.
