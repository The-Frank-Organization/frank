## DESIGN-REVIEW -- VP exact-candidate review of the Step-3 architecture-amendment packet r2

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-review-r2
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- only a revised VP-approved exact candidate may proceed to operator hash-bound ratification
GRILL_REQUIRED: yes -- the existing step3-arch-reframe-grill lock remains satisfied for retained decisions; reopen only if the repair chooses a new cross-seat grantor or ceiling-override model
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-040000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-3.implementer, m-4.planner, m-5.planner, m-5.implementer, m-6.planner, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
SUBJECT: must-revise at packet sha256 2cd16311 -- retain G8 direct authority, but remove proxy/transitive grants; route the m-5 ceiling-host amendment; define the no-conductor-change E0 carrier; return a clean exact candidate

VERDICT: revise

Review target: `master/STEP-3-ARCH-AMENDMENT.md` r2 at SHA-256 `2cd16311387d2410d5cf375e9b40490865994abc0c706c59961a0fb8d1f3a200` plus transmittal `040000`.

The prior prerequisite and most of `030000` are now closed correctly. The m-7 same-seat correction exists and is consumed; roadmap and kickoff hashes remain frozen; the packet chooses the coherent no-routing Step-3 branch; the m-9-only seat map, connector process split, request ordering, no-retry rule, per-family state/recovery matrix, carry ledger, m-10 pair dependency, and G1-G10 grill shape are materially repaired.

The exact candidate is not ratifiable yet. The flagged G8 fold turns an agent-authored report of an operator instruction into a potentially transferable authority record despite the landed grantor grammar; the packet also relocates m-5's locked enforcement host without an m-5-owned amendment and leaves the E0 app-summary carrier undefined under the claimed no-conductor-change branch.

## Findings

### F11 -- G8's direct authority decision is valid, but the packet's recording rule creates proxy and transitive authority

I do **not** reopen the operator's product decision: the live direct operator channel may be authority-bearing for its one addressed recipient, and the operator need not author a governed relay merely to prove an instruction already given in that channel.

The packet goes further. Section 8b says **any receiving agent** authors a governed effect under its own `FROM`, gives `HUMAN_MERGE_AUTHORIZATION` as the example, permits an app-side action or override, and calls the resulting audit trail sufficient (`STEP-3-ARCH-AMENDMENT.md:106-114`). That is not the landed authority grammar:

- the typed `grant` field allows `dispatch-merge` only to `operator` and `*.orchestrator-planner`; a pair planner may grant only `dispatch-impl`, and other seats have no grant option (`frank/internal/fieldspec/registry.json:105,111`);
- the m-1 contract makes an agent-stamped record proof only of that agent's authorship. Operator authorship is separately stamped through the operator channel (`master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:145-155`);
- the direct route is explicitly one-recipient/non-transitive, but treating the recipient's citation as operator authority lets that citation authorize a different consumer;
- `record-never-fabricate` is a behavioral instruction, not a confusion-resistant mechanism. A confused recipient can misquote scope or conditions without impersonating the operator. The agent-stamped citation is therefore self-reported evidence of the instruction, not by-construction operator provenance.

Required fold, preserving G8:

1. A direct instruction authorizes only the **directly addressed recipient** and only an action that recipient may perform within its already-bound capability/authority ceiling. The recipient may record the instruction and resulting effect under its own `FROM`, but that record is an evidence/audit record, not an operator-stamped grant and not transferable authority.
2. A different seat acts only after either (a) the operator directly instructs that eventual actor, still without authoring a relay, or (b) a protocol-sanctioned grantor such as `master.orchestrator-planner` emits the existing typed grant under its own stamped authority while citing the direct operator instruction as context. Remove the arbitrary-agent `HUMAN_MERGE_AUTHORIZATION` example.
3. A direct instruction cannot silently or merely textually raise m-5's immutable spawn/run ceiling. Any ceiling change uses the m-5-owned typed reconfiguration/respawn contract and its applicable gate; otherwise the direct recipient remains bounded by the current ceiling.
4. Scope the phrase "by construction" to authentication of the **live ingress**. Do not apply it to the later agent-authored citation or audit record.
5. If the intended design is instead that any seat's citation becomes a conductor-recognized cross-seat grant, this is a new m-1/m-2/lineage and operating-protocol amendment. Reopen G8 for that stronger branch, record the deviation in `master/PROTOCOL-DEVIATIONS.md`, and reconcile `master/CYCLE-PLAYBOOK.md` Part F. The current candidate cannot claim that branch while also claiming no conductor change.

### F12 -- m-5's locked ceiling enforcement moved to m-10 without an owner amendment or consumable interface

The packet makes m-10 the app-side authority-enforcement point and requires the above-ceiling-tool E2 negative (`STEP-3-ARCH-AMENDMENT.md:27-29,61,69,97,100`). Its replacement graph, however, stands up m-10, then m-8/m-9, then m-3 and credential work; m-5 has no owner/reviewer leg (`:104`).

The locked m-5 design currently says ceilings are recorded per assignment, initially enforced by host config, and later enforced uniformly by the conductor; it explicitly routes the enforcement dependency to the orchestrator (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:158-174`). Reassigning the host to m-10 is a real locked-boundary amendment, not a charter-only text substitution. The packet also does not say what immutable m-5-authored artifact m-10 reads, who writes it, how it binds to the run/worker, or how m-10 fails closed when it is missing or stale.

Required fold:

- add an m-5 planner-authored, adversarially reviewed amendment/consumer interface before m-10/m-9 consumer lock (or interface-lock it ahead of those locks);
- pin the ceiling artifact's source, writer, schema/config home, immutable binding to `run_id` plus worker identity, m-10 read/load path, and fail-closed absent/stale behavior;
- preserve m-5 as sole policy owner and m-10 as enforcement host only;
- add that owner leg to the replacement dependency graph and propagation set. Do not silently rewrite only the m-5 charter while its locked design still names conductor/host-config enforcement.

This repair follows the ratified topology and needs no new operator choice unless it changes the ceiling semantics.

### F13 -- the E0 app-attestation summary still has no exact no-conductor-change carrier

The packet says m-3 owns the app policy/attestation schema, m-8 emits the event, and the m-9 worker submits one E0-labeled conductor summary that cannot satisfy a gate or promote evidence (`STEP-3-ARCH-AMENDMENT.md:35,53,64,69,79-80`). It simultaneously says no new relay kind, FieldSpec row, or trusted-observer input is needed (`:78-81`).

The landed schema does not yet express the asserted provenance boundary. `attestation_source` has only `{conductor, operator}` and is system-owned (`frank/internal/fieldspec/registry.json:72-78,147-155`); the live observe gate stamps `attestation_source=conductor` and computes `achieved_evidence`/`record_integrity` for the **relay observation**, not for an embedded connector event (`frank/internal/observe/gate.go:178-257`). Without an exact carrier, a reader can mistake the relay's conductor-produced evidence stamp for corroboration of the app event, which is the laundering F7 prohibited. No reader/consumer or non-gate carrier class is named either.

Choose and state one coherent branch in the packet:

1. **No-conductor-change branch (recommended):** name the exact existing ordinary non-authority relay shape used by the worker, the m-3-owned app-event schema serialized in its body, the reader/consumer, and the rule that the conductor's system evidence fields describe only carriage/observable relay claims, never the embedded connector event. The embedded event remains E0/self-reported, body data is not gate-referenceable, and the relay carries no typed grant/gate resolution. State how this is mechanically prevented from becoming authority.
2. **Typed conductor provenance branch:** add an m-2/m-3 FieldSpec/observe amendment for app/connector attestation provenance and retract the packet's no-conductor-member-change claim. This branch needs its own owner reviews before lock.

The downstream m-3 amendment may own the detailed app schema, but this architecture packet must close the carrier branch because it is what makes the no-conductor-change claim true or false.

### F14 -- the exact candidate and transmittal still contain final-byte/status contradictions

- Packet line 3 cites a placeholder transmittal `step3-arch-packet/...-04xxxx`; an exact ratification candidate must cite `040000` or its eventual replacement exactly.
- The transmittal says `GRILL_REQUIRED: no` while the design record correctly carries `GRILL_REQUIRED: yes` and a closed `GRILL_LOCK`. Required is not synonymous with pending; report it as required-and-satisfied and carry the lock ID.
- `ACTIONS_GIT_REF` says the leg wrote section 9's durable `GRILL_LOCK` and later says "no lock". Qualify the latter as no `DESIGN_LOCK_ID`/architecture ratification; a grill lock was created.
- Any F11 process deviation retained after repair must add `master/PROTOCOL-DEVIATIONS.md` and the affected Part-F text to the atomic propagation list. The recommended non-transitive repair stays within the existing grantor grammar and avoids that extra amendment.

These edits change the packet bytes. Recompute and return a new exact candidate hash; r2 at `2cd16311...` remains `must-revise` and must not be ratified.

## Accepted Folds That Need Not Be Reopened

- F1/F2: corrected m-7 state, honest source action, frozen hashes, append-only historical treatment, and supersession-on-ratification.
- F3/F4: app-side pinned manifest, no Step-3 routing/V3 or lane FieldSpec row, and only the instantiated m-9 worker as a conductor seat.
- F6: separate m-8 process before E3, immutable `freeze -> authorize -> attach -> send`, no conductor credential member, and one attempt/no automatic retry.
- F8/F10: per-family state writers, stable IDs, no cross-store atomicity, fail-closed recovery, complete carry ledger, m-10 planner/reviewer dependency, and one m-8-owned credential contract with m-1 review.
- F9: G1-G10 are recorded in a durable grill shape. G8 needs the mechanism correction above, not a product-decision re-vote.

## Required Revision Sequence

1. Fold F11-F14 without changing the accepted decisions above.
2. Whole-document reconcile the direct-route authority/non-transitivity/ceiling language and the E0 carrier/no-conductor-change claims.
3. Add the m-5 owner/reviewer dependency and exact ceiling interface requirement.
4. Recompute the packet SHA-256 only after all placeholder/status text is final.
5. Return the new exact candidate for VP review. Reopen the operator grill only if the planner chooses a stronger grantor/ceiling model than the bounded repairs above.

The five holds remain in force. No operator ratification, supersession, source fold, refreshed consumer dispatch, lane resumption, design lock, PLAN, code, credential, provider call, external send, merge, or deployment is authorized by this review.

## Verification

- Packet r2 and transmittal `040000` read in full; review bound to SHA-256 `2cd16311387d2410d5cf375e9b40490865994abc0c706c59961a0fb8d1f3a200`.
- Incoming transmittal exact-file lint -> OK; INDEX row present once.
- m-7 corrected same-seat handoff `step3-hold-m7/SITREP-planner-20260715-020423.md` exists and carries r3 `must-revise`, F11-F13 open, and reviewed-but-provisional status.
- ROADMAP remains `3977c9f2e1a0fd57903b5381031c337e4a8d25fe7b1d26de5b06f2eaec58c7c3`; kickoff remains `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43`.
- Landed authority/evidence behavior checked at the registry, lineage, observe, m-1, and m-5 sources cited above.
- `go test ./internal/fieldspec ./internal/lineage ./internal/observe` -> PASS.
- `frank/` remains clean on `main@502e06c`.
- New reviewer relay exact-file lint -> OK; INDEX row survival check -> reviewer row present once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-043000.md and appended its master/relays/INDEX.md row; no packet, roadmap, kickoff, architecture, charter, domain design, protocol register, frank source, branch, commit, push, merge, tag, live-store, credential, provider-call, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
