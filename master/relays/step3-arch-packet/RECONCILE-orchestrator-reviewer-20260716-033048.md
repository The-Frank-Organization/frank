## RECONCILE -- r6 implements the two mechanisms, but its release set and seat supersession remain incomplete

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r14
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the replacement hash still requires operator-authored ratification after a clean exact-byte review
GRILL_REQUIRED: no -- the operator-locked topology and Option B remain accepted; F65/F66 are exact-scope and supersession repairs, not new choices
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-031629.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW -- r6 adds the release gate and broker fence, but omits the conductor from its every-binary claim and leaves the worker-as-seat packet fragment operative against the new logical-seat model

VERDICT: revise

Review target: `master/STEP-3-MVP-AMENDMENT.md` r6 at SHA-256 `5d66bf246d2c78df5fd895b92d4b291c515bffb2ca2a48668413ff67b66a6578`, planner transmittal `031629`, reviewer requirements F63/F64 in `031044`, the unchanged reframe packet, and the live conductor channel boundary.

## Findings

### F65 -- the release vector still omits the separate conductor while claiming every binary that executes the governed turn

r6 correctly binds the three T4-built app artifacts and mechanically verifies the shipped tool catalog (`STEP-3-MVP-AMENDMENT.md:48-49,58,86,117,123`). But the candidate says **"Every binary that executes the governed turn is bound"** while its vector contains only app-main/m-10, m-9, m-8, and conditionally the shared conductor-client (`:48`). The conductor is explicitly its own service (`:36`; packet `:15`), and the same exit test defines the turn as including a governed **relay exchange** (`STEP-3-MVP-AMENDMENT.md:94`). The conductor service, not merely the shared client linked into m-9, executes that `submit`/`project`/`read` leg (packet `:73`). Its bytes can change while every r6 release field remains equal.

The packet does describe the live E3 observation specifically as a provider turn (`STEP-3-ARCH-AMENDMENT.md:75`), so either scope is defensible; the r6 text currently mixes them.

Required correction: choose one exact claim. If the live E3/exit binding covers the whole governed turn, add the conductor service artifact (and the governing config identity on which the observed relay claim depends) to the release/E3 vector and mutation test. If E3 covers only the app/provider vertical, replace "every binary that executes the governed turn" with that narrower scope and name the separately-bound exact conductor evidence used for the relay-exchange leg of the exit test. In either case, the evidence tuple, release event, exit test, and annex must describe the same artifact set.

### F66 -- the logical-seat fix contradicts an operative packet fragment and the amendment's own process-topology line

The new vocabulary says the durable **logical m-9 seat** owns the credential through a broker, while each supervised worker generation receives only an epoch-bound use capability and never the bytes (`STEP-3-MVP-AMENDMENT.md:15`). That is the right F64 model. However, section 1's exhaustive supersession list still does not supersede packet `:27`, which defines **"m-9 Model Runtime -- the WORKER"** as the seat and gives that worker process its **own private seat channel**. r6 then repeats the old model in section 2b: **"m-9 = the supervised worker process (the only conductor seat)"** (`STEP-3-MVP-AMENDMENT.md:38`). Those statements cannot remain simultaneously operative with a durable broker-held channel outside replaceable worker generations.

Required correction: add the exact packet `:27` worker-as-principal/private-channel fragment to section 1's supersession list. Replace it with the pinned distinction: the logical m-9 identity is the sole app seat/principal and owns the private authenticated channel through the m-7 broker; a replaceable m-9 worker acts only through its epoch-bound capability and is not the credential/channel holder. Align section 2b line 38 and require any allowed broker placement to live outside the replaceable worker generation. Preserve the packet's sole-app-seat, genuine-relay-only, m-10-no-seat, and no-provider-through-conductor rails. The post-ratification F44 source fold must carry that same wording into the m-9 charter/table.

## Accepted Closures

- F63 closes at the mechanism level: pre-build interface identity and post-build release binding are separated; the three app artifacts and shipped registry are mechanically bound; the run/E3/annex/build-order references agree.
- F64 closes at the mechanism level: every broker verb and push handoff is epoch-checked; linearization and in-flight disposition are required; m-10 supplies lifecycle state without gaining credentials or conductor verbs; the acceptance row exercises stale-generation denial.
- The one-credential-per-logical-seat topology, broker custody, Option B ticket, all three grill decisions, and F57-F62 remain accepted and must not be reopened.
- r6 changes are confined to F63/F64 plus revision/provenance bytes; the exact transmittal is named.

## Required Return

1. Fold F65/F66 only, preserving the accepted r6 mechanisms and every earlier closure.
2. Return r7 at a fresh SHA-256, refresh the README pointer and ordered 15-file manifest, and name the exact r7 transmittal.
3. Request another exact-byte review. Any byte change voids review of the prior candidate.

No ratification recommendation, source fold, first-stage dispatch/interface-lock, release-binding execution, `DESIGN_LOCK_ID`, PLAN, T4 code token, implementation, credential provisioning, provider call, external send, merge, deployment, or live-store mutation is authorized by this review.

## Verification

- Amendment r6 SHA-256 independently recomputed: `5d66bf246d2c78df5fd895b92d4b291c515bffb2ca2a48668413ff67b66a6578`.
- Ordered 15-file governing manifest independently recomputed: `215bf39a796d15e5ed88228d0989269a702b852432872e5d04184893dbda4a09`; README exact at `9614728697807d6076dd5a8156f14b9cec5667b91e092ee381cdc96d3a9825e1`.
- Reframe packet remains exact `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 artifact remains exact `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- Incoming `031629` exact-file lint ends `OK`; root-mode historical/INDEX lineage debt remains separate.
- Targeted `TestSecondConnectSameCredentialRejected|TestProvenDeadRecovery` passes, confirming the live channel remains credential-scoped rather than worker-epoch-scoped.
- `frank/` remains clean on `main@502e06cc07b5` (`s11-close`).
- New reviewer relay exact-file lint ends `OK`; its INDEX row is present exactly once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-033048.md` and appended its `master/relays/INDEX.md` row; no amendment, governing source, historical relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main remains clean at `502e06c`.
Next requested action: master.orchestrator-planner folds F65/F66 only into r7, refreshes the exact hashes and manifest, and requests a fresh exact-byte review; all ratification and build authority remain held.
