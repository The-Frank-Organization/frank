## RECONCILE -- r4 has sound narrowings, but six exact-contract defects block ratification of 57aa3170

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r12
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the final replacement hash still requires operator ratification, and two amendment-level boundary choices below require an operator-visible grill disposition first
GRILL_REQUIRED: yes -- r4 introduces unresolved cross-domain seat-topology and authorize/execute placement choices; downstream owner grills cannot retroactively resolve an already-ratified ambiguous architecture contract
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-172000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- accept the external-review direction, but r4 preserves the hard credential rail it claims to narrow, ratifies tool identities without values, leaves authorize/execute and seat topology unresolved, overclaims wake delivery, and mismatches several acceptance proofs

VERDICT: revise

Review target: `master/STEP-3-MVP-AMENDMENT.md` r4 at SHA-256 `57aa3170499e8f8b3fcb2f6487b8544719f1b9c107416cf323bf8e1487d27960`, planner transmittal `172000`, the unchanged reframe packet and m-5 contract, current m-1/m-7/m-9/m-10 ownership sources, and the live conductor push/project implementation.

## Findings

### F57 -- the credential claim is not actually narrowed because the old hard rail remains explicitly preserved

Section 1 says only the ceiling and first-stage-order fragments are superseded, then explicitly preserves m-8 as **"NOT in m-9's credential-readable address space"** and says the m-9 credential is absent from `bash` files (`STEP-3-MVP-AMENDMENT.md:9-21`, especially `:14,16`). Section 2 correctly says the opposite security truth for a same-UID unsandboxed shell: peer-process state, debug interfaces, crash dumps, and runtime files may remain readable, and only non-injection / accidental-disclosure prevention is claimed (`:23-26`). The acceptance row then re-expands the claim to "bash inherits no control capabilities" although its proposed env/FD test proves only no **injected or inherited** handles (`:101`).

Required correction: add an exact supersession for the packet `:29` credential-readable-address-space phrase and align every preserved copy to the narrow non-injection claim. Preserve separate-process placement, but do not call it a hard unreadability boundary. Rename the acceptance claim to what the test proves: no secret/control handle is injected through argv, env, inherited FDs, or the enumerated accidental surfaces. Same-user discovery/use remains the explicit Step-4 residual.

### F58 -- the ratified tool constant contains names, while the enforced identity contains missing future values

Section 4 calls the tool set an operator-fixed ratified constant and says it is "PINNED + hashed in this amendment," but the candidate pins only eight names (`:43-47`). It then defines equality over `name + tool-schema digest + tool-implementation/catalog version + form-to-schema-mapping version` (`:48`) without supplying any digest/version values, canonical derivation, field applicability for local versus relay tools, or owner-produced identity vector. Ratifying this file cannot ratify values that do not exist yet. The same undefined provenance resurfaces in `run_manifest_digest`, `tool_catalog_digest`, and `policy_digest` evidence bindings (`:39,50,105`).

Required correction: separate **policy identity** from **build identity**. The operator may ratify the eight canonical names. The owner graph must define who produces each schema/mapping/implementation identity component, its canonical encoding and applicability, and the exact later lock event that binds the built catalog vector/digest. m-10 then compares the run manifest to that locked vector. Alternatively, include all exact identity values now. Do not claim the complete identity is already pinned in this amendment when it is not.

### F59 -- `authorized == executed` is an unresolved cross-domain architecture branch, not a design requirement with one owner

Section 4 leaves two materially different architectures open: m-10 directly invokes the executor, or m-10 emits a durable one-shot authorization record consumed elsewhere (`:49`). Option A moves execution into m-10 despite the current m-9 MVP source assigning local tools to m-9 (`master/domains/m-9-model-runtime/README.md:39`) and changes the control-plane boundary. Option B needs a writer, consumer/executor, canonical-args owner, atomic consume/replay rule, crash disposition, and epoch check. The section 7 graph names m-10's dispatch seam and m-9's turn state but does not assign that complete cross-domain contract (`STEP-3-MVP-AMENDMENT.md:69-76`). The acceptance row proves only that an authorization was bound to a digest; it does not compare the actual executor invocation to the authorized call (`:99`).

Required correction: resolve A versus B before ratification. **Recommended:** preserve the current boundary with a m-10-owned durable one-shot authorization/consume protocol and an m-9-owned executor, then make both halves explicit in the graph and reciprocal consumer review. If A is chosen instead, explicitly amend the m-9/m-10 ownership sources. In either case, the proof must capture the actual executor invocation's canonical tool/args/epoch identity and compare it to the authorization, including duplicate, stale-epoch, mutation, and crash-window negatives.

### F60 -- seat topology is a list of unanswered questions assigned to the wrong sole owner

The named m-9 requirement asks whether the broker is a process/thread/module, whether credentials are shared, what a logical seat is across worker instances/restarts, and how launch custody works (`STEP-3-MVP-AMENDMENT.md:78-80`). Those are architecture and identity/channel questions, not merely m-9 implementation details. m-1 owns the stamp/seat identity contract; m-7 owns authenticated channel lifecycle; m-10 owns worker replacement; m-9 consumes the seat. The graph currently routes only secret-boundary text to m-1 and transport/client text to m-7 (`:69`), while assigning the unresolved topology to m-9.

Required correction: run the amendment-level grill and pin the topology before ratification. **Recommended MVP default:** one broker-held credential per logical m-9 seat, never copied into worker generations; replacement workers are fenced by `turn_epoch` but do not mint an implicit new identity. m-1 authors/reviews identity and credential-lifecycle semantics, m-7 authors/reviews the authenticated channel/broker contract, and m-9/m-10 consume and confirm. If the operator chooses worker-per-seat instead, name that explicitly and update the ownership graph. `GRILL_REQUIRED: no` is not defensible while this and F59 remain open.

### F61 -- at-least-once wake delivery is unsupported and contradicts the accepted lost-wake stretch boundary

Section 6 claims notification delivery is at-least-once/retryable, then says push is advisory and a lost wake is acceptable (`STEP-3-MVP-AMENDMENT.md:61-64`). The live conductor ignores `PushTo` failure (`frank/cmd/frank/main.go:345-352`). On authentication it emits only a recovery nudge when the durable mailbox is nonempty (`:358-364`). `project()` can replay all accepted mailbox IDs and the mailbox remains durable (`frank/internal/store/store.go:202-247,273-281`), so reconnect/startup rediscovery is real, but there is no acknowledgment or periodic reconciliation that recovers a lost push while the connection stays up. That is not general at-least-once notification delivery.

Required correction: either (recommended for the non-gating stretch) state **best-effort push + durable rediscovery on startup/reconnect + at-most-once scheduling by unique relay ID**, with lost wake explicitly allowed, or add an acknowledgment/retry/periodic-reconciliation contract that actually proves at-least-once delivery. Do not claim both at-least-once delivery and acceptable loss.

### F62 -- the evidence/acceptance annex broadens unchanged conductor evidence and several rows do not prove their claims

"Every evidence record" is required to carry a new scope/provenance tuple including `relay_record` (`STEP-3-MVP-AMENDMENT.md:38-40`), but the amendment also preserves no conductor schema/member change and says m-3 owns only the app E0 schema. The candidate does not identify the external E3 artifact, writer, canonical digest producers, or applicability evaluator. As written, the phrase can be read as changing existing conductor evidence records.

The annex also mismatches mechanisms: the authorized/executed row observes only authorization (`:99`); the bash row proves inherited env/FD absence, not absence of same-UID control capability (`:101`); and "sentinel payloads at every conductor entry point" (`:102`) conflicts with the expressly legal case where an agent-authored relay quotes provider/tool content (`:37`) unless the test is constrained to **automatic typed forwarding**, not natural-language bytes.

Required correction: scope the new provenance tuple to a named app/external evidence artifact; assign m-3's schema/evaluator ownership and the observer writer; define canonical digest producers and applicability evaluation. Rewrite each annex row so its observable proves exactly its claim, especially actual executor equality, no injected/inherited handle, and no automatic typed-envelope route while preserving legal agent-authored relay content.

## Accepted Direction

- The r3 approval is correctly void after the byte change; `57aa3170...` receives no inherited approval.
- The exact candidate hash, unchanged reframe-packet hash, unchanged m-5 hash, README hash, and ordered 15-file manifest all reproduce.
- The threat-claim narrowing, explicit turn-versus-attempt model, no-SDK-retry requirement, full frozen-core identity, typed automatic-route boundary, durable app state, epoch fencing, honest unknown/partial states, and acceptance-annex approach are sound directions. They may remain after the contradictions and ownership gaps above are corrected.
- F39-F56 remain historical inputs; this review opens F57-F62 only against r4.

## Required Return

1. Resolve F57/F58/F61/F62 in the amendment text and annex without weakening the accepted narrow threat boundary.
2. Route F59/F60 through one amendment-level grill, record the operator-visible decisions durably, and fold the resulting single-owner graph. Owner-lane grills may refine implementation but cannot choose a different architecture after ratification.
3. Return r5 at a fresh SHA-256 with the README pointer and ordered 15-file manifest refreshed. Any further byte change gets another exact-byte review.

No ratification recommendation, source fold, first-stage dispatch/interface-lock, `DESIGN_LOCK_ID`, PLAN, T4 code token, implementation, credential provisioning, provider call, external send, merge, deployment, or live-store mutation is authorized by this review.

## Verification

- Amendment r4 SHA-256 independently recomputed: `57aa3170499e8f8b3fcb2f6487b8544719f1b9c107416cf323bf8e1487d27960`.
- Ordered 15-file governing manifest independently recomputed: `d16023ee2b8e84abb7b82ccda4195aff63fc1be42cbd18f82e4452e3e5d64b26`; README exact at `57fd064a793a3fcea8c6239e9b392ba4eaa211e12d9ef0a165c5c94c8cce4369`.
- Reframe packet remains exact `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 artifact remains exact `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- Incoming `172000` exact-file relay lint ends `OK`; root-mode historical/INDEX lineage debt remains separate.
- Live channel/store read verified push error handling, reconnect nudge, durable mailbox projection, and absence of a delivery acknowledgment in the cited current path.
- `frank/` remains clean on `main@502e06cc07b5` (`s11-close`).
- New reviewer relay exact-file lint ends `OK`; its INDEX row is present exactly once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-180236.md` and appended its `master/relays/INDEX.md` row; no amendment, governing source, historical relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main remains clean at `502e06c`.
Next requested action: master.orchestrator-planner routes the two amendment-level choices through a durable grill, folds F57-F62 only into r5, refreshes the hashes/manifest, and requests a fresh exact-byte review; all ratification and build authority remain held.
