## DESIGN-REVIEW -- VP review of the Step-3 architecture-amendment packet

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-review-r1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator must close the revised one-question-at-a-time grill and ratify the exact final packet hash
GRILL_REQUIRED: yes -- section 9 is presently an open agenda, not a durable GRILL_LOCK
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-023000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-3.implementer, m-4.planner, m-4.implementer, m-5.planner, m-6.planner, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
SUBJECT: must-revise -- choose one coherent no-routing/no-conductor-change MVP, correct stale m-7 state, close component identity/direct route/credential/recovery boundaries, then run the operator grill

VERDICT: revise

Review target: `master/STEP-3-ARCH-AMENDMENT.md` at SHA-256 `818c3d871bd90afd6815cf09b14f605042fdd0b69dd271340d3944959893b83c` plus its `023000` transmittal.

The packet is a useful topology draft, but it is not ready for grill closure or exact-byte operator ratification. It imports stale handoff state, treats open grill recommendations as operative sequence decisions, and simultaneously claims both branches of the central MVP fork: "single pinned route / no routing execution / no conductor changes" and "m-4 exact-lane routing decision + V3 + new conductor-carried records/events." Those cannot all be true.

## Findings

### F1 -- The packet bypasses the operative review and carries the known-stale m-7 state

The transmittal replies to `013000`; the packet lineage at `:8-9` and `:181-183` omits the operative `023000` and `024000` reviewer relays. No corrected m-7 hold return exists. Nevertheless, the transmittal calls all five handoffs clean and the disposition table says m-7 F7-F10 are unreviewed (`:148`).

The durable r3 reviewer relay `step3-amend-m7-cred/DESIGN-REVIEW-implementer-20260714-234854.md` is `must-revise`. It confirms F8 closed and accepts the F7/F9/F10 directions while opening F11-F13: catalog-v2 transition/drift, full immutable freeze binding, and schema-versus-composition validation/timing.

Required fold:

- obtain the same-seat corrected m-7 hold SITREP required by `024000`; do not proxy-author its status;
- carry full lineage `011000 -> 013000 -> 020000 -> 023000 -> 024000 -> packet`;
- state the honest m-7 disposition: no lock/PLAN/code/provisioning, but substantial r3 design and three review rounds exist; F1-F6 remain confirmed, F8 closed, F7/F9/F10 directions accepted, F11-F13 open, and all material is provisional input for a fresh owner/reviewer;
- remove "all clean" and "zero unwound work". The accurate narrow claim is zero code, credential, provider, live-store, lock, or PLAN action.

### F2 -- A draft packet cannot supersede the locked kickoff or silently normalize the early roadmap fold

Packet `:6` says the reframed kickoff sections are superseded by this packet while the packet is still a draft awaiting two gates. Change that to **proposed supersession effective only upon operator ratification of a hash-bound final packet**. Until then, the kickoff remains the historical lock and the five holds prevent unsafe use of its stale Step-3 framing.

The current roadmap baseline is SHA-256 `3977c9f2e1a0fd57903b5381031c337e4a8d25fe7b1d26de5b06f2eaec58c7c3`. Section 8 merely says the re-cut is already folded; it does not disclose the baseline/provenance or reconcile the contradictory old tech-stack, PTY/supervision, and conductor-timed interjection clauses identified in `023000`/`024000`.

Required fold:

- record that roadmap hash as a provisional, non-operative direction baseline and include its complete correction in the post-ratification fold;
- do not edit historical boot/design/review relays. Packet `:178-179` must replace "sweep" of the m-9 boot relay with append-only supersession references; current charters and architecture may change, historical relays remain the record;
- compute the final packet SHA-256 **before** operator ratification, have the reviewer confirm that exact candidate, and bind the operator decision to that hash;
- correct the transmittal action claim: creating `STEP-3-ARCH-AMENDMENT.md`, the relay, and INDEX row is a docs-workspace disk action, not `ACTIONS_GIT_REF: none`.

### F3 -- Step-3 cannot both defer routing execution and execute an m-4 routing decision

Ratified R-Q2 says one pinned route and no routing execution (`:22-24`). Sequence A instead has m-10 record an m-4 governed routing decision before the run and explicitly invokes V3 (`:93-96`). Section 8 then re-dispatches the m-4 amendment inside the MVP dependency graph (`:185-189`).

The live schema proves this is not free compatibility: `frank/internal/fieldspec/registry.json:175` has `routing_assignments` with `chosen_model` and no `lane_ref`, `catalog_snapshot_digest`, or four-axis lane tuple. The locked kickoff itself says adding exact-lane routing requires the m-4/m-2 amendment. Therefore packet `:126-138` cannot also claim no conductor member changes.

Recommended coherent MVP branch:

1. Step-3 V1 uses an **app-side pinned run manifest**, m-10-owned as run state and bound to an immutable m-8 lane ID/catalog digest. It is not an m-4 policy decision and performs no routing.
2. Lane identity in that manifest is never a conductor gate input. A later ordinary relay may carry an opaque, E0-labeled run summary, but Step-3 adds no lane-bearing FieldSpec row.
3. Defer the m-4/m-2 routing-record amendment and V3 to the named Step-4 routing-execution gate. m-4 may consumer-review the manifest boundary now; its held amendment does not re-dispatch as a Step-3 writer.
4. Remove `routing DECISION` from sections 2/4/6/8 for the MVP and remove G4's presumption that the old provisional Q1 shape lands now.

If the operator instead wants an accepted m-4 exact-lane decision in Step-3, R-Q2 must be reopened and the packet must admit routing execution plus an m-2/m-4 conductor-schema amendment. Do not describe that branch as "no routing execution/no conductor change."

### F4 -- Domain ownership labels are conflated with runtime conductor identities

Rows `:42-44` and compatibility text `:127-129` call m-10/m-9 runtime components stamped participants/seats without defining a new principal class. The landed m-1 contract mints **agent seats**, gives `operator` a special trusted channel, and reserves `system` for conductor-internal provenance. `m-8`, `m-9`, and `m-10` are domain owners; those labels do not automatically mint runtime identities.

For a no-conductor-change MVP, pin this identity map:

- the instantiated m-9 agent worker is the only app component that is also a conductor **seat**, with its own private existing seat channel;
- m-10 control plane and m-8 connector are trusted app components, not seats, not stamped participants, and hold no `submit` credential;
- m-10 supervises without authoring as the worker; m-8 returns app events/attestations over IPC but does not stamp a conductor record;
- any conductor relay about the run is submitted by the actual worker seat and is labeled as app-supplied/E0; component provenance is not upgraded by the worker's seat stamp;
- the scheduler bridge reuses the worker seat's existing delivery/project/read path. It is not a new conductor event or m-10 principal.

If component principals are required, route a separate m-1/m-2/m-7 identity/schema amendment and retract "no conductor byte changes."

### F5 -- The operator direct route remains only a slogan

Sections `:20-21`, `:45`, `:66-67`, and `:130-132` enumerate the route but do not close any of the mandatory boundaries from `023000` F3. Add an explicit direct-route contract and grill item covering:

- authenticated local operator ingress and the trusted component that identifies it;
- operator-to-one-agent scope only; no agent-to-agent, connector/provider, credential, or generic app IPC use;
- conversational versus authority-bearing effects;
- no `FROM: operator`, operator verdict, relay-store mutation, or operator-channel impersonation through this route;
- no implicit m-5 ceiling/tool-authority bypass; any override is separately typed and evidenced;
- app-side transcript/audit writer, retention, and labels that never call the route conductor-stamped/observed;
- no cross-seat proxying, unlabelled fallback after conductor failure, secret carriage, or provider-byte carriage.

Recommended MVP shape: an authenticated local m-10 surface may deliver direct operator text to one worker; it remains app-side input. Any conductor mutation still enters through the existing operator-relay channel, and any cross-seat message is a fresh worker submission through the conductor.

### F6 -- The connector/credential boundary contradicts the live E3 floor and the authorization ordering is unsafe

The matrix marks a co-located connector as settled for the MVP and says physical split occurs when a real key lands (`:44`), but Step-3's E3 call necessarily requires a real provider credential. If co-located means one address space with the worker, "keys never enter any seat" is false. If it means same host but separate process, say that exactly.

Required architecture:

- before the first E3 call, m-8 is a separate trusted connector process from the m-9 worker; same host is acceptable, same credential-readable address space is not;
- m-8 is the credential runtime holder and secret-store writer/reader under an m-1-authored boundary; m-10 handles only opaque binding/reference orchestration and never secret bytes;
- no m-7 `provider_bindings`/engine-v5 credential member is created;
- use one immutable authorization identity over method, canonical endpoint, non-auth headers, and body. The current sequence `credential attach -> policy -> final authorization` (`:101-103`) must be replaced by an exact freeze/authorize/attach/send contract that neither exposes secrets to m-3 nor permits mutable request substitution. The fresh owner must explicitly consume m-7 r3 F12/F13 as findings, not as accepted design.

For the MVP, default to **one provider attempt and no automatic retry**. Packet `:106` assigns retries to m-8 but defines no durable attempt state, provider idempotency support, reauthorization rule, or crash ambiguity. A retry-enabled branch needs its own owner contract and negative proofs; m-9's C8 no-replay identity does not make a provider send exactly once.

### F7 -- The packet preserves the E0 label but leaves writer, reader, and proof semantics open

Section 3 correctly states that connector outcome is E0/self-reported absent corroboration. Sequence A then says "m-3 produces" the conductor summary (`:104-105`), but m-3 is a policy owner, not a runtime process or conductor principal. No writer, schema, consumer, or submitter is named. G3 remains open.

Close the planes separately:

- **app governance mechanism:** m-3 owns the policy/attestation schema; m-8 enforces and emits the app event; m-9 is the only seat that may submit an explicitly E0-labeled relay summary in the no-conductor-change MVP;
- **conductor authority:** that summary is neither authoritative app state nor conductor observation and cannot satisfy a gate or promote evidence;
- **MVP verification:** deterministic E2 negatives prove policy denial -> zero socket send, no post-authorization mutation, no secret leakage, above-ceiling tool -> zero execution, and no provider bytes in conductor surfaces. A separate integration harness/operator observation proves the real provider turn at E3. Do not launder that E3 result into the conductor summary unless a future trusted observer contract earns it.

Define "honest governed turn" as the app enforcing the locked owner policies with correctly labeled proof, not as every app event becoming conductor-observed.

### F8 -- The state/recovery and scheduler matrices name owners but do not define recovery

Rows `:75-79` say only "m-10-owned", "m-9-owned", or "m-8-owned" and leave m-9's writer/home open. That does not satisfy the required canonical-store, writer, crash/recovery, and authority matrix. G5 is still open. The scheduler bridge at `:120-122` also implies a new typed conductor event while section 6 claims no conductor change.

Minimum architecture-level closure for the MVP:

- m-10 owns canonical run/supervisor state and the active-turn lease;
- m-9 owns/writes canonical turn/session state;
- m-8 owns/writes canonical provider-attempt/telemetry state;
- all three bind stable `run_id`, `turn_id`, `request_id`, and `attempt_id`; no cross-store atomicity is claimed;
- on crash or disagreement, recovery is fail-closed to interrupted/held diagnosis, never automatic provider resend; m-10 may start a replacement only after the prior worker/attempt is proven terminal or an explicit operator disposition;
- the m-9 local one-active-turn invariant and m-10 one-active-worker/lease invariant are distinct and both named;
- a governed human gate is opened/read by the worker seat through existing conductor verbs; m-10 receives app IPC from that worker. No new m-10 conductor address or wake API is implied.

Domain designs may choose concrete storage engines later, but these ownership, ambiguity, and recovery dispositions belong in the architecture packet.

### F9 -- Section 9 is an open question list, not a GRILL_LOCK

The required grill artifact must record sources, questions answered from code, questions asked operator one at a time, resolved decisions, rejected alternatives, still-operator-owned items, and design-lock impact. Section 9 has none of that shape. It marks G2 `SETTLED-IN-GRILL` in section 1 while G2 remains open in section 9; it adopts G4's recommendation in Sequence A while G4 remains open; and it omits the direct-route, runtime-principal, retry, and recovery forks.

Do not ask the VP to invent and close G1-G7 in one review. First fold the source-resolvable corrections above. Then conduct the operator grill one dependent question at a time and produce a real `GRILL_LOCK_ID: step3-arch-reframe-grill` containing:

- G1 conductor boundary confirmation;
- G2 exact process/address-space split before E3;
- G3 attestation schema boundary and E0/E3 proof split;
- G4 Step-3 pinned-manifest versus reopened m-4 routing branch;
- G5 state writers/lease/recovery;
- G6 existing-seat scheduler bridge versus conductor amendment;
- G7 no conductor credential config and Step-3/Step-4 lane-ref disposition;
- G8 direct-operator route;
- G9 runtime component/principal map;
- G10 retry policy and crash ambiguity.

Only after every item is resolved may the document call that ID a lock.

### F10 -- The replacement graph and carry ledger remain incomplete

"m-10 charter lands first" is not enough for consumers to design against. Name the m-10 planner/reviewer seats and require its boundary design + adversarial review before, or interface-locked ahead of, m-8/m-9 consumer lock. The credential re-owner must have exactly one writer: recommended m-8 planner owns the connector credential design, m-1 reviews the secret boundary, and m-10/m-3 consume it.

The packet also lacks the complete carry table required by `023000` F5. Add an explicit disposition for V2/two-provider portability, V3/routing execution, benchmark, native spawn, steer/interrupt, Step-2 T5/T10, the soft-expiry cell, T4 team mechanics/relaunch, and each old consumer-lock seam. Every item must be retained in MVP, replaced, deferred to a named Step-4+ gate/owner, or retired by operator decision. Preserve Steps 4, 5, and 6 as distinct observable gates.

## Required Revision Sequence

1. Obtain and consume the corrected m-7 hold return.
2. Fold F1-F8 and the complete carry ledger into packet r2; remove premature supersession and freeze the roadmap baseline.
3. Reconcile the packet internally with a whole-document search for stale routing/V3/conductor-principal/no-change/credential/retry copies.
4. Run the expanded one-question-at-a-time operator grill and write the durable lock artifact.
5. Return the folded packet with its candidate SHA-256 for adversarial re-review.
6. Only an approved exact candidate goes to the operator for hash-bound ratification; only then may source propagation, refreshed consumer audit, or replacement design dispatch occur.

The five holds remain in force. No architecture lock, source fold, lane resumption, PLAN, code, credential, provider call, external send, merge, or deployment is authorized by this review.

## Verification

- Packet and transmittal read in full; packet review bound to SHA-256 `818c3d871bd90afd6815cf09b14f605042fdd0b69dd271340d3944959893b83c`.
- Incoming transmittal exact-file relay lint -> OK; INDEX row present once.
- Operative `023000`/`024000` reviewer requirements reread; no corrected m-7 hold return exists.
- m-7 r3 `DESIGN-REVIEW` reread: `must-revise`, F11-F13 open.
- Live FieldSpec checked at `frank/internal/fieldspec/registry.json:175`: `chosen_model` exists inside `routing_assignments`; exact lane/digest fields do not.
- Conductor tool surfaces checked at `frank/internal/channel/server.go:391-415` and `frank/cmd/frank-mcp/mcp.go:129-155`: exactly `submit`/`project`/`read`.
- m-1 operator/seat/system identity contract reread at its design `:145-156`; no m-8/m-10 service-principal class exists.
- Roadmap SHA-256 remains `3977c9f2e1a0fd57903b5381031c337e4a8d25fe7b1d26de5b06f2eaec58c7c3`; kickoff remains `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43`.
- `frank/` was not modified; clean on `main@502e06c`.
- New reviewer relay exact-file lint -> OK; INDEX row survival check -> reviewer row present once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-030000.md and appended its master/relays/INDEX.md row; no packet, roadmap, kickoff, architecture, charter, dashboard, domain artifact, frank source, branch, commit, push, merge, tag, live-store, credential, provider-call, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
