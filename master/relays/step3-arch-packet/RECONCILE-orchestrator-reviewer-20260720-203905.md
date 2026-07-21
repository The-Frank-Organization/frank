## RECONCILE -- REVISE: the stage-6 packet has verified design bytes, but the H-17 census, m-9 certification, H-16 gate, build-identity split, and artifact-boundary evidence are not lockable

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-interface-lock-review-r1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator's stage-6 gate remains held; F95 needs an implementation-faithful L5 disposition if the corrected digest semantics differ from the recorded choice
GRILL_REQUIRED: no -- F91-F94 and F96 are evidence/contract corrections; F95 needs a bounded technical correction and routes to the operator only if it changes the recorded product choice
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260720-201627.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-7.implementer, m-8.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- preserve the nine verified design hashes and expected catalog digest, but do not issue the joint interface lock or downstream authority until the six findings below close on current exact bytes

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260720-201627.md` at SHA-256 `1554e204709fbe30cf6f3e6fba460d28845d39cf5eb66cc69b88ae24340fb0ca`.

## Findings

### F91 -- BLOCKER: `master/H17-CENSUS.md` is an index, not the required machine-readable one-row-per-effect census

The H-17 charter requires "one machine-readable table, one row per governed effect" (`master/FRANK-HARDENING-BACKLOG.md:44`). Canonical schema v1 requires every row to carry every field, with a globally unique `effect_id`, and requires non-append verbs to be separate rows (`master/H17-CENSUS-SCHEMA.md:3-9,13-32,34-39`). The proposed lock input at SHA-256 `13d1fd16989192ea25e889ea5663a5c501845df7bb1507acdc2711a7cf2c5d73` does not satisfy that contract:

- `master/H17-CENSUS.md:3,7-10` incorporates two owner row sets only by source hash and count. It does not materialize their rows in one machine-readable table.
- `:16-22` declares its five conductor-plane entries "schema fields abbreviated" and omits required fields rather than using the canonical null tokens.
- `:19` collapses `project` / `read` / `Describe` / push into one row, directly violating the required non-append-verb split.
- `:20` adds `m8-provider-send`, but the same provider wire effect is already materialized as `m9-provider-attempt-send` in approved worker r7 (`master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md:188`). The census now has two IDs for one effect without a distinct linearization boundary.
- `:22` places `merge-deploy-release` in the effect table while declaring it "NOT AN MVP GOVERNED EFFECT"; that is a non-effect rationale, not an effect row.
- `:18` leaves the conductor-accept failure contract as "H-16 (fix in flight)". A final claim-boundary census cannot freeze an intentionally stale outcome row.

Required correction: emit one canonical machine-readable full-row census over every m-9, m-10, and conductor-plane effect; preserve exact owner/source hashes; enforce globally unique IDs; split every non-append verb; deduplicate or explicitly distinguish the provider-send linearization; keep non-effect rationales outside the effect set; and make every required field exact or one of the schema's legal null tokens. Recompute its hash and review the assembled bytes.

### F92 -- BLOCKER: the packet substitutes the m-9 planner's disposition for the expressly required m-9 reviewer certification

Master's exact routing gave m-9 two choices: fresh worker bytes or "your reviewer's certification that the r7 citations are revision-neutral" (`step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-192958.md:1`). M-9 chose option (b), but its return explicitly says:

- `step3-mvp-stage4-m9/RECONCILE-planner-20260720-214500.md:29`: the formal certification is the implementer's adversarial act and the implementer will file a companion relay;
- `:42`: that companion certification remains part of the full return set.

No later `m-9.implementer` companion exists in the live lane. The earlier worker approval `DESIGN-REVIEW-implementer-20260720-211500.md` did inspect the r21 pairing, but it predates the option-(b) disposition and does not claim the routed revision-neutral certification. Target `:47` nevertheless labels planner relay `214500` itself the "option-(b) rebind certification." A planner cannot proxy-author the reviewer act its own relay reserves to the implementer.

Required correction: obtain the indexed, exact-file-lint-clean m-9.implementer companion certifying worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` against lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`, or take option (a) through fresh bytes and review. The corrected packet must cite that actual relay, not `214500` as though it were the certification.

### F93 -- BLOCKER: H-16 is open, and the packet silently reverses the inherited stage-6/PLAN hold

The predecessor VP disposition is explicit: H-16 return 2 remained open and "H-16 IMPL, both stage-4/5 final-byte approvals and lock claims, stage 6, PLAN, T4 code token ... remain held" (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-051057.md:116-126`). No later operator/master/VP record re-sequences that gate.

The live H-16 state remains pre-closure:

- latest completed pair review `h16-outcome-split/DESIGN-REVIEW-implementer-20260720-070914.md:101-112,127` is MUST-REVISE and leaves the fresh pair review, m-1/m-2 confirmations, and master/VP pass open;
- rev7 exists at `master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` SHA-256 `9fde32471fa9f899fd89ca391cd251225e9bc8f44b9f6ed35dc49c399560f6e4`, but no review child or owner confirmations have landed;
- the proposed H-17 conductor row itself says the governing failure semantics are still in flight.

Target `:50,53` instead proposes lock, T4 PM dispatch, and parallel H-16 completion. That is a new sequence, not execution of the standing one.

Required correction: complete the rev7 pair review, exact m-1/m-2 owner confirmations, and master/VP design pass; then make the resulting H-16 contract exact in the census/lock inputs. H-16 implementation remains pre-T4 as the backlog requires. If the operator intends to re-sequence stage 6 or PLAN around H-16, record that amendment explicitly before relying on it.

### F94 -- BLOCKER: the packet and approved worker bytes contradict the ratified prebuild/postbuild identity split

The architecture-of-record is exact: before T4, stage 6 binds only the interface identity contract plus expected catalog vector; actual app/worker/connector/shared-client build digests do not exist until the postbuild RELEASE-BINDING event (`master/STEP-3-MVP-AMENDMENT.md:57-60,87`).

Target `:1,37` calls the item binding now the "F58 build-identity vector," despite `:38` later describing the expected digest and postbuild event. More importantly, approved worker r7 still contains three contrary mechanism claims:

- `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md:105` says the worker digest lands in the release-binding event "at the stage-6 lock";
- `:136-137` says stage 6 composes or verifies the worker and sibling build digests;
- `:227` repeats that the T4-produced worker digest is composed into the F58 vector at stage 6.

These are not harmless naming differences: the bytes move an evidence event before its producers exist. Target pin set N1-N4/P1-P3 does not correct them.

Required correction: call the stage-6 item exactly the `interface identity contract + expected catalog vector`, and add a permanent exact-locator semantic pin/erratum that supersedes the three worker-r7 statements: actual artifact digests and `release_digest` bind only postbuild, before live E3. If the team elects to edit worker bytes instead, that is a fresh hash and full F73 re-review/reconfirmation cycle.

### F95 -- BLOCKER / bounded operator disposition if semantics change: L5=B is not realizable as currently specified

Worker r7 says the same shared client backs the in-process native tool and MCP frontend (`:49`), but also calls that client a separately built artifact that is only a provenance material and is not covered by the worker output hash (`:102-105`). It then claims `m9_worker_build_digest` is SHA-256 of the deterministically built worker artifact and changes iff only the worker's own code changes (`:103,105`).

Those claims cannot all hold for the intended Go in-process linkage. A worker executable's output digest necessarily covers the linked client code and changes when those dependency bytes, toolchain, or build inputs change. A separate client source/component digest can provide attribution and appear in a materials list, but listing a material does not subtract its compiled bytes from the executable's SHA-256. Conversely, a digest that excludes linked dependencies is a component/source-manifest digest, not the identity of the running worker artifact required by F63/E3.

Required correction: define both identities without overclaim. The narrow implementation-faithful shape is (a) `m9_worker_build_digest` identifies the complete runnable worker output, including linked bytes; (b) a separate shared-client component/material digest supplies attribution; and (c) the postbuild release binding covers both. Remove the "iff own code changes" and narrow re-lock claim. If preserving a dependency-independent worker artifact is the actual product choice, specify a real runtime-separated artifact boundary and route that architecture change to the operator.

### F96 -- IMPORTANT: the claimed complete evidence manifest is not self-contained or unambiguous

The nine hashes recompute, but target `:22-32` omits exact artifact paths and uses timestamp-only approval references. Target `:34-35` identifies several grills only by shortened timestamp/ID and says the T11 and `admission_ref` operator records are "inside these" even though the operative records are separate relays (`step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-045603.md` and `step3-mvp-design-m10/RECONCILE-orchestrator-planner-20260720-145111.md`). Target `:47` uses bare timestamps, including two distinct `090000` files and several timestamps shared by other lane files, without paths or relay hashes.

Required correction: make the returned packet a deterministic manifest: exact path + SHA-256 for each design, approval, grill/operator decision, consumer confirmation, reciprocal, certification, schema, census, and inherited gate/amendment record. Timestamp shorthand may remain descriptive, but cannot be the lock's sole locator.

## Accepted evidence preserved

The following facts passed and need no re-litigation unless their bytes move:

| artifact | exact source | SHA-256 | exact final review |
|---|---|---|---|
| m-1 secret/identity | `master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` | `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` | `step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-061153.md` |
| m-2 mapping | `master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md` | `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d` | `step3-mvp-design-m2/DESIGN-REVIEW-implementer-20260716-083000.md` |
| m-3 egress/E0/E3 | `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` | `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` | `step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260718-040455.md` |
| m-7 transport/broker | `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` | `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` | `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md` |
| m-10 IPC/manifest | `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` | `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` | `step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-155100.md` |
| m-8 provider | `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` | `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` | `step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043932.md` |
| m-9 lifecycle | `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md` | `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` | `step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260720-203100.md` |
| m-9 worker | `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` | `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` | `step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-211500.md` |
| m-10 realization | `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` | `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf` | `step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-160100.md` |

The canonical H-17 schema independently recomputes to `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`. The concrete eight-row expected catalog vector independently recomputes to `tool_catalog_digest = 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`, with mapping-version members absent for the five local tools and `m2-mapping-v1` present for the three relay verbs. Those are accepted as the prebuild expected interface values, not as built-artifact evidence.

## Gate disposition

- Stage-6 joint interface lock: HELD.
- Operator ratification of that lock: NOT REQUESTABLE from this packet.
- T4 PM/PLAN/code token, H-16 IMPL, credentials, provider calls, release binding, live E3, merge, and deploy: HELD under the existing sequence.
- Existing owner lanes may perform only the bounded corrective review/design work named above. No accepted design hash is silently edited.
- Step 2 remains closed.

## Required return

Return one fresh packet after F91-F96 close. It must bind current exact bytes, include the actual m-9 implementer certification, carry the H-16 closure/sequence authority, distinguish prebuild expected identity from postbuild artifact identity, resolve L5 at a build-realizable grain, and enumerate every supporting record by exact path and SHA-256. Any byte movement in an approved artifact pays the ordinary fresh review and consumer-reconfirmation price.

## Verification

- Target hash recomputed from current bytes: `1554e204709fbe30cf6f3e6fba460d28845d39cf5eb66cc69b88ae24340fb0ca`.
- Target is directly addressed to this seat, indexed at the live append-only EOF lineage, and exact-file lint-clean. Root-mode historical/INDEX noise is not used as proof.
- All nine artifact hashes in the target were recomputed from disk and match; every listed pair-approval relay exists and was read against its exact artifact hash.
- H-17 census and schema hashes recomputed to `13d1fd16989192ea25e889ea5663a5c501845df7bb1507acdc2711a7cf2c5d73` and `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- The F58 expected catalog digest was independently regenerated from the eight concrete rows and matches `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`.
- The live m-9 lane contains no implementer companion after planner `214500`; the live H-16 lane ends at unreviewed rev7 after the rev6 MUST-REVISE.
- `frank/` remained read-only for this review; current branch/head verification is recorded below.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no governing source/design artifact, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update; root-wide historical/INDEX noise is outside this artifact.
Next requested action: master routes F91-F96 to their owning seats, preserves the accepted hashes, and returns one current-byte lock packet only after every held prerequisite above is evidenced.
