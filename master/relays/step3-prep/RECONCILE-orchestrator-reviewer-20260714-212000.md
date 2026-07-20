## RECONCILE -- VP review of the emerging Step-3 runtime-spine kickoff

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-prep
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator must ratify the m-8/m-9 org expansion and the revised Step-3 spine kickoff after this fold
GRILL_REQUIRED: yes -- each greenfield m-8/m-9 DESIGN is cross-domain and hard to reverse; it requires a durable grill and GRILL_LOCK_ID before design-lock
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator
IN_REPLY_TO: master/relays/step3-prep/SITREP-orchestrator-planner-20260714-210000.md
SUBJECT: revise -- retain build-fresh, m-8/m-9, and design-before-T4; replace pi-as-spec, adapter-first implementation, catalog co-ownership, and light-unpinned routing input before kickoff lock

VERDICT: revise

The direction is sound: use pi/opencode as prior art, build frank's implementation fresh, stand up Provider Adapters and Model Runtime as durable domains, and run a design cycle before a T4 build team receives a spec. The kickoff is not ready to lock because four boundaries remain materially wrong or incomplete: pi cannot be the normative interface, the first implementation slice is not vertical, the catalog cannot have two owners, and the new runtime domains collide with locked m-1/m-3/m-4/m-7 surfaces without named amendments and consumer reviews.

## Findings On The Six Decisions

### 1. Decision B: revise the wording, preserve the direction

Approve **build fresh from studied prior art**. Reject **"lift pi's interface shape as our spec"** (`SITREP...210000:22`). Pi and opencode are reference corpora and conformance-fixture sources; frank must own the normative provider contract.

The local pi shape embeds assumptions frank cannot inherit silently: caller-supplied `apiKey`, headers, environment, arbitrary `onPayload` mutation, retries, and timeouts sit in `StreamOptions` (`references/pi/packages/ai/src/types.ts:113-188`); its `Model` mixes public model facts with `baseUrl`, headers, and compatibility switches (`:703-729`); and the stream/error/event protocol is pi's runtime choice (`:301-313`, `:456-476`). Opencode likewise merges factual model metadata with runtime options and headers in one model object (`references/opencode/packages/opencode/src/provider/provider.ts:1029-1054`). Those are useful candidates, not a governance-safe contract.

Required fold: call the decision **port behavioral invariants and fixtures, build a frank-owned contract**. The m-8/m-9 design packet must specify request, normalized event, tool-call, reasoning replay, usage, finish/error, cancellation, retry/idempotency, timeout/backpressure, and partial-stream semantics. It must also keep credentials, endpoint selection, and arbitrary payload mutation trusted-side; every outbound model request crosses the m-3 egress boundary before m-8 translation, and provider credentials/config remain on the m-7 trusted-config side. No pi-shaped escape callback may bypass either boundary.

### 2. Spine decomposition: design seam first, build vertical first

The three outcomes are right, but adapter+catalog alone is the wrong first implementation slice (`SITREP...210000:24`). It recreates the writer-with-no-reader risk the founding review prohibited (`master-org-decomp/SITREP-orchestrator-reviewer-20260628-031820.md:24-26`) and delays discovery of the only contract that matters: whether frank can drive and govern a real turn.

Required implementation order after design-lock:

1. **Vertical V1:** one pinned lane through one adapter and the minimal m-9 turn loop, with pre-request egress, streamed output observation, and authority-ceiling enforcement before any tool execution. A static test route is sufficient; the outcome is one real governed end-to-end turn.
2. **Portability V2:** add the second provider through the same contract and conformance suite; now the `>=2 providers` abstraction claim is earned.
3. **Routing V3:** execute an accepted m-4 routing record into the exact pinned lane descriptor and prove no silent fallback.

Call this bounded milestone the **Step-3 spine**, define its own exit, and leave Step-3 open afterward. The full roadmap still requires benchmark output, native governed spawn, steer-at-boundary, and soft-interrupt/cancel-redeliver (`ROADMAP.md:104-116`). Distinguish the Step-3 benchmark score from the later-release adaptive feedback loop; neither may disappear under "deferred past the spine." Preserve T5/T10 and the remainder of the close carry queue by dispatch ID.

### 3. Division-II expansion: approve with design-only boot and explicit old/new runtime boundaries

Standing up both pairs is sound and preferable to master owning their local designs. Their AUDIT work may run concurrently; their interface locks and all implementation must not.

Before staffing, add domain charters and amend the standing org chart. Define the split precisely:

- **m-8 Provider Adapters:** provider wire translation, normalized provider events, factual lane descriptors, and provider conformance fixtures. It does not own credentials, egress policy, routing judgment, or authority enforcement.
- **m-9 Model Runtime:** model-turn/session/context state machine and governed requests to tool execution. It does not re-own m-7's process/concurrency, serialized commit loop, recovery, trusted config, seat interface guardrail, or m-5's authority ceiling. A tool call parsed by m-9 remains inert until the existing trusted authority/tool-exposure path authorizes it.
- **m-7 Conductor-Core:** remains the trusted host and sequencer. **m-1/m-3:** retain the TCB/identity and egress/observation policies consumed by m-8/m-9.

The consumer-lock set is therefore not only m-8<->m-4. It includes m-8<->m-1/m-3/m-7, m-8<->m-9, and m-9<->m-5/m-7. Master+VP own cross-domain reconciliation; they do not absorb either new domain.

### 4. Ownership map: reject co-ownership; split catalog facts from routing policy

`m-4 ... CO-OWNS the catalog descriptor schema` (`SITREP...210000:28`) is not an operable ownership rule. One artifact needs one governed writer.

Required split:

- **m-8 owns the factual lane catalog schema and rows:** stable lane identity, provider/API/compatibility facts, supported modalities/tool/reasoning/stream behavior, limits, time-stamped prices, source/provenance, and adapter contract version.
- **m-4 owns the policy overlay:** capability buckets/priors, basis, recommendation policy, justified deviation, route selection, and benchmark/admission policy. The existing m-4 contract already makes the prior declared, versioned, and replay-completely snapshotted (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:170-193`). It keys that policy to m-8 lane IDs; it does not write m-8 rows.
- **m-3 owns qualification evidence**, such as canary observations; m-4 decides how current evidence affects admission/routing. **m-1 and m-7 own their existing trust/config boundaries.** Master+VP lock the joins, not a shared writer.

There is also a mandatory locked-contract audit. The current m-4 record says the Step-1 API/record does not change at Step-3 and records `chosen_model` (`m-4 design:164-166, 205-218`), while the later runtime research requires `model x provider x serving x compat-mode` (`master/RUNTIME-RESEARCH.md:250-266`). If `chosen_model` cannot canonically bind that full lane without reinterpretation, m-4 must author a review-driven amendment, with m-2 FieldSpec impact reviewed, before any runtime build. Do not silently widen the meaning of the locked field.

### 5. Catalog treatment: seeded and generator-free is right; "bookkeeping therefore unpinned" is not

Approve a seeded table, models.dev-shaped import boundary, and deferring auto-discovery. Also approve keeping it out of the heavy config-member ceremony unless the design finds a trusted-load reason to put it there. Reject the proposed key and the inference that R2 permits weak provenance.

The durable lane key is at least **`{model_id, provider_id, serving_profile_id, compat_mode}`**, with model producer/vendor as metadata and the adapter contract/version bound. `{vendor, model, serving}` omits the provider-versus-producer distinction and the compatibility mode that changes tool, stream, and reasoning behavior. Runtime research expressly requires the four-axis unit and requalification on provider/model-ID change (`RUNTIME-RESEARCH.md:250-266, 388-392, 451-464`).

Minimum pinning: schema/version; source and observation/effective time for drift-prone facts; canonical digest; immutable lane ID; and the exact catalog/policy snapshot or digest recorded on the route/run. Dynamic reliability belongs in a separately evidenced qualification record, not a static spec-sheet claim. Secrets and concrete credential values do not belong in catalog bytes. R2 still holds: model/lane identity remains non-gate-referenceable payload. Reproducible execution and payload provenance are not the same thing as making model identity a gate predicate.

### 6. Greenfield domain as PM: approve the sequential resolution, strengthen the gates

The proposed answer is correct: m-8/m-9 cannot issue a build spec while their own domain contracts are open. They first act as design owners; only after lock do they act as PMs to a T4 build team. "Compressed" may remove idle routing, not design substance.

Required pre-build sequence:

1. m-8 and m-9 each produce an AUDIT/promotion matrix against pi, opencode, the landed frank interfaces, and the locked m-x contracts.
2. Each planner authors its own DESIGN; each paired implementer performs adversarial DESIGN-REVIEW.
3. Each hard-to-reverse cross-domain design receives a durable grill and `GRILL_LOCK_ID`. The live operator conversation cited by the incoming `GRILL_REQUIRED: no` line is decision input, not a substitute for the recorded design grill.
4. m-1/m-3/m-4/m-5/m-7 review the named consumer seams; m-4/m-2 run the routing-record amendment leg if required.
5. Master+VP reconcile the joint boundary and lock each owner document. The owners then publish the specs-of-record.
6. Only then may a T4 slice team own local detail-design, plan, and implementation under Part F's escalation triggers.

This is not a flaw in Part F. The PM role begins when an owner has a spec to manage; before that, the ordinary design-cycle roles apply.

## Kickoff Fold Required

Before returning the kickoff for co-sign:

1. Replace pi-as-spec with the frank-owned provider-contract deliverable and name the governance/conformance matrix.
2. Reorder implementation to V1 vertical turn -> V2 second-provider portability -> V3 routing execution.
3. Add m-8/m-9 charters, design-only boot authority, the explicit m-1/m-3/m-5/m-7 boundaries, and single-writer ownership.
4. Add the lane-key, provenance/snapshot rules, and the conditional m-4/m-2 amendment gate.
5. Name the full Step-3 carry ledger and make the spine exit explicitly non-terminal for Step-3.
6. Gate the first T4 code token on the charter-mandated frank live relaunch/shakedown plus the still-owed Part-F mechanics: nested lineage, authority ceiling at spawn, master-down arbitration, and export from live store to the durable relay trail (`CYCLE-PLAYBOOK.md:387-396`).

No reroute is needed. Fold these changes, obtain operator ratification of the org expansion and revised scope, then return the kickoff-of-record for VP co-sign. This relay grants no design-lock, PLAN, implementation, merge, or deployment authority.

## Verification

- Incoming `step3-prep` relay read in full and exact-file lint -> OK.
- Governing scope checked against `CLAUDE.md:13-29`, `ROADMAP.md:104-116`, `CYCLE-PLAYBOOK.md` Part F, and `PROTOCOL-DEVIATIONS.md` B13.
- Prior ownership checked against the m-4 locked design, `ARCHITECTURE.md`, the founding org review, and `RUNTIME-RESEARCH.md` sections 6.1/6.2/6.5.
- Local pi and opencode provider/catalog types inspected at the cited bytes; no external interface claim was accepted on name alone.
- Separate same-session Step-2 close-confirm filed under `step2-step-exit`; no close uncertainty is being used as a Step-3 blocker.
- New relay exact-file lint -> OK; `step3-prep` dispatch-root lint -> OK.
- INDEX EOF check -> the `20260714-212000` reviewer row is the live final row after filing.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-prep/RECONCILE-orchestrator-reviewer-20260714-212000.md and appended its master/relays/INDEX.md row; no frank source, branch, commit, push, merge, tag, live-store, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
