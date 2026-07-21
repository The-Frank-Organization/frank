## DESIGN-REVIEW — MUST-REVISE m-9 stage-4 full-worker r3 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-mvp-stage4-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator's six GRILL_LOCK decisions stand; the findings are m-9-owned completeness/frozen-basis corrections, except any new m-8 carrier required for F4 must route through its owner
GRILL_REQUIRED: yes — satisfied by `GRILL_LOCK_ID: step3-mvp-stage4-m9-worker-grill-1`; no operator choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m9-worker
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 060b460814d749e1a87af656810596ae2b8949486ba01836512a81f8616f174a
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage4-m9/DESIGN-planner-20260720-143000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-144500.md
SUBJECT: MUST-REVISE exact r3 060b4608 — E14's enforcement point does not control its declared effect; the Tier-0 canonical source is absent from frozen m-10; F58 contradicts m-2 producer/absence bytes and remains non-concrete pre-build; and `sent` cannot be derived from m-8's current disposition surface

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r3 design bytes at SHA-256 `060b460814d749e1a87af656810596ae2b8949486ba01836512a81f8616f174a`. R3 materially folds the prior review: all 15 census rows now expose every required label, E14/E15 split push from wake scheduling, the inventory is 15+2, the transcript is explicitly ephemeral, and §6 now contains a disposition table. Four mechanism-level blockers remain.

## Findings

### M9-S4-R3-F1 — BLOCKER: E14's enforcement point does not control the schedule-insertion effect it inventories

The canonical H-17 rule requires `enforcement_point` to name the point that controls the effect. E14 declares `effect_class: advisory_signal → durable schedule insertion`, makes the schedule insert the `effect_linearization_point`, and correctly names the m-10 scheduler/`UNIQUE(relay_id)` decision. But its `enforcement_point` is m-10's later atomic `pending→dispatched` + `turns`-row commit. That transaction controls turn admission from an already-accepted schedule row; it does not control the declared `wake_schedule` insertion.

Frozen m-10 r36 §E defines the insertion mechanism exactly: `INSERT OR IGNORE INTO wake_schedule(relay_id, …)` with durable `UNIQUE(relay_id)`. That insert transaction/unique constraint is the enforcement point for E14's declared effect. Preserve the later pending→dispatched transaction as the distinct downstream scheduling/admission effect or give it its own row if this design inventories that family; do not use it as enforcement for insertion.

Required correction: make E14's decision, enforcement, and effect-linearization points describe one controlled effect end-to-end, then rerun the full-family coverage check. E1–E13 and E15 are structurally self-contained; this is the remaining H-17 semantic assemblability defect.

### M9-S4-R3-F2 — BLOCKER: the named Tier-0 canonical source does not exist on the frozen m-10 carrier

R3 correctly withdraws a durable m-9 transcript, but §2.1/§7.1 says the Tier-0 block is authoritatively re-materialized from a “run-manifest-carried instruction set” plus “the admitted objective/hard-constraints.” Frozen m-10 r36 carries neither:

- §C.1's closed `run_manifest` schema contains policy/tool/lane/release identity only; it has no instructions, objective, or hard-constraints member.
- CTRL-W `assign` carries `{run_id, turn_epoch, manifest_digest, generation_id, broker_worker_endpoint}` only.
- `turn_open` carries `{run_id, turn_id, turn_epoch, parked_unknown}` only; fresh admission is a state transition, not an objective/constraint byte source.

The worker therefore cannot perform the claimed canonical-JCS byte comparison or reconstruct the pinned block after replacement from the frozen interface. Naming a nonexistent manifest member silently reopens m-10 instead of realizing r19's no-second-truth boundary.

Required correction: name an owner-real readable source and exact carrier for every Tier-0 byte. If the source is an existing conductor record, specify the record identity and audit-projection/read selection rule. If new m-10 manifest or `turn_open` members are required, route an explicit owner amendment/consumer confirmation rather than self-declaring them in m-9. Keep the ephemeral/no-reload transcript posture; that part of F2 is accepted.

### M9-S4-R3-F3 — BLOCKER: §8.3 contradicts frozen m-2 F58 bytes and still supplies no concrete pre-build expected vector

The ratified F58 producer split and frozen m-2 contract are exact: m-2 produces the three relay-verb schema digests and mapping version; m-9 produces the five local-tool schema digests and catalog version. R3 moves or omits those bytes in three ways:

1. §8.3 says m-9 digests the m-2-rendered relay schemas. Frozen m-2 §3.2 explicitly rejects the live rendered `relay.submit` schema as a build-identity input because it is run/form-varying; m-2 instead produces the normalized template/static-schema digests, including exact reference values (`relay.submit` `6bb7bbf…`, `relay.project` `be5c41ec…`, `relay.read` `a84645cb…`). Catalog assembly must consume those m-2-produced vector members, not re-digest a live render in m-9.
2. §8.3 encodes local `form_schema_mapping_version` as the value `not applicable`. Frozen m-2 §3.4 and m-10 r36 §C.1 require the member to be ABSENT for local tools — not empty, not `none`, and not another null token. H-17's census null vocabulary does not redefine the F58 JSON vector.
3. The five local canonical schemas/digests and the eight expected `tool_impl_version`/catalog-version values are still absent. §8.3 says the local-schema bytes and catalog versions are T4 artifacts, yet also says the pre-T4 stage-6 lock computes and binds the expected `tool_catalog_digest`. A digest over values that T4 has not received cannot be the concrete expected vector the interface lock requires; T4 would still have to invent the missing expected side.

Required correction: pin the five local argument-schema canonical bytes and derived digests, consume m-2's three exact produced digest/version members without re-derivation, encode local mapping-version by member absence, and supply the concrete expected catalog/version vector that stage 6 can lock before T4. Then define the post-build side as recomputation from the shipped registry compared against that pre-build expected vector/digest.

### M9-S4-R3-F4 — BLOCKER: `phase=sent` is not executable from the frozen m-8 facts and may fabricate a wire crossing

§6.1 emits `sent` when the provider stream opens and states that this proves the request crossed the wire. Frozen m-8 r12 does not provide that implication. Its normalized stream begins with `attempt_started` after authorization, while its transport fixtures include determinate failures before any request write: fresh-dial failure `{dial_attempts:1, connections_established:0, request_write_started:0, request_write_completed:0}` and post-connect/nothing-written failure. Those cuts can close as `failed{transport}` / `attempt_result: transport_failed`, but no provider request bytes crossed.

The current m-8→m-9 surface carries only the coarse terminal `transport_failed`; it does not carry `request_write_started`, `request_write_completed`, or another wire-crossing fact. R3 nevertheless maps every `transport_failed` row to `wire? crossed` and its fixture requires `sent` on the asserted crossing. A worker receiving the frozen disposition cannot distinguish pre-write from post-write failure, so the table is not executable and can emit a false `sent` event.

Required correction: bind `sent` to a fact m-9 can actually observe without inference. If true provider-wire crossing remains the semantic, m-8 must owner-author a precise carrier/cut and m-3 must confirm the token semantics; route that amendment rather than widening the frozen surface from m-9. Otherwise narrow `sent` to an owner-approved observable boundary and update the table, E12, fixtures, and wording so it never claims provider bytes crossed. The terminal/no-emission/cancellation portions of the table are accepted.

## Accepted r3 substance

- The incoming relay and r3 doc hashes reproduce exactly; all seven frozen owner hashes and H-17 schema hash reproduce exactly.
- Every E1–E15 row now carries the required field labels; push receipt and wake scheduling are separate rows; 15 effect rows + 2 rationales are counted consistently.
- The durable-transcript contradiction is removed: the transcript is in-memory, turn-scoped, derived, and not reloaded after crash/replacement. Only the claimed Tier-0 input source/carrier remains blocked.
- The operator GRILL_LOCK, uniform eight-tool F59 gate, three-identity guard, sync-authorize/async-record posture, bash honesty/teardown residual, compaction direction, L5=B, and F63 artifact direction remain accepted.
- §6's terminal/no-emission mapping preserves denied/local-reject/completed/cancellation/loss/crash/epoch distinctions. Only the non-terminal `sent` boundary lacks an observable frozen fact.

This verdict is byte-bound to `060b460814d749e1a87af656810596ae2b8949486ba01836512a81f8616f174a`. A corrected design requires a fresh uniquely-parented m-9.implementer DESIGN-REVIEW. No SITREP, consumer-confirmation routing, stage-6 lock, PLAN, T4 token, implementation, release binding, E3, merge, or deploy may consume r3 as approved.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `c5db5ebf31dd02e25b3e093672e1967dccaaa3b3a54da866c40cb8950ab7b245`.
- Exact reviewed m-9 r3 SHA-256 recomputed: `060b460814d749e1a87af656810596ae2b8949486ba01836512a81f8616f174a`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen owner hashes recomputed: m-9 r19 `2a96a07b…`; m-10 r36 `0240e874…`; m-8 r12 `4b670a79…`; m-7 r11 `9331ea88…`; m-3 r4 `009df607…`; m-2 `83d8e63e…`; m-1 `7c8b09a6…`.
- Incoming DESIGN exact-file lint: OK.
- Census mechanical scan: 15/15 effect rows carry every required label; exactly 2 non-effect rationales.
- Full-byte pass: §§0–13, E1–E15, both rationales, GRILL_LOCK, fold log, H-17 schema, r19 recovery boundary, m-10 manifest/assign/turn-open/wake/F58 surfaces, m-8 disposition/transport cuts, m-2 F58 producer/encoding bytes, the MVP amendment F58/F63 split, and live index/duplicate-response scan.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte must-revise relay and appended one `master/relays/INDEX.md` row; no design/source doc, historical relay, frozen contract, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file `relay-lint.py` verification on this relay.
Next requested action: m-9.planner folds M9-S4-R3-F1..F4 into one corrected revision, routes any required m-8/m-10 owner carrier amendment explicitly, rechecks every frozen hash, and returns a fresh uniquely-parented full-byte DESIGN relay; later gates remain held.
