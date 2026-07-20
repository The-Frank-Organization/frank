## DESIGN-REVIEW - m-7 adversarial review of the Step-3 transport, broker, and conductor-identity contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r1
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings close owner-local interface mechanics under the ratified topology and claim ceilings; no product fork is requested
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260716-050007.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260716-050908.md
SUBJECT: must-revise - close the metadata seam, credential custody, epoch/retry linearization, restart recording, and live-binary proof

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed the final contract bytes at SHA-256 `bddf868ba050e201302ece0c63e9544b4f6402b0ad5302b49769d98f33818b2e`, the addressed DESIGN, both master dispatches, the ratified MVP and architecture amendments, the parallel m-1/m-10 owner bytes, the m-1 pair-review verdict, and the live transport/MCP implementation at `frank@502e06c`.

The F67 grill closes: own-process broker placement is supported by the ratified topology, keeps credential custody out of app-main, and makes the m-10 no-verb/no-bytes rail testable at process and interface grain. Complete-and-deliver is also a defensible disposition for an invocation already admitted before the epoch linearization point. Five contract gaps still block pair approval. Two are confirmed upstream mismatches in the now-must-revise m-1 contract; three would permit implementations that pass the listed fixtures while violating the stated seam, stale-generation, or exact-artifact claims.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Findings

### F1 - The three-verb caller strands the required `DescribeTools` transport path

Section 1.2 defines the shared caller as only `Call(ctx, name canonical-verb, args)` and says it is implemented by both the direct client and the worker-side broker path (`2026-07-16-step3-mvp-transport-broker.md:37-38`). Section 1.3 correctly keeps `DescribeTools` transport in m-7 while assigning interpretation to m-2 (`:42-48`). Section 2.8 then says the worker IPC carries exactly the three relay verbs plus push and nothing else, while also saying describe rides the same seam (`:120-122`). Those statements cannot all be implemented.

This is not optional metadata. The live MCP path calls `Client.DescribeTools` directly to obtain the current form and digest before m-2's `SchemaFromForm` interpretation (`frank/cmd/frank-mcp/mcp.go:209-225`; `internal/channel/server.go:481-490`). The native worker needs the same rediscovery/re-render path for parity. With the proposed interface it either bypasses the broker with a credential, loses dynamic form refresh, or adds an unreviewed side channel.

Required revision:

1. Add a typed m-7-owned description/metadata operation to the shared transport abstraction and worker-broker IPC. Its return is raw `DescriptionResponse`; all FieldSpec/form interpretation stays in m-2.
2. Gate description requests at the same current-epoch linearization point as the three relay calls, because the operation still exercises the authenticated logical-seat channel. Keep it absent from the m-10 control interface.
3. Replace "three verbs + push and NOTHING else" with the exact closed worker surface: three canonical relay calls, typed describe/rediscovery, and push receipt. Do not misclassify describe as a conductor relay verb.
4. Add parity fixtures for direct/MCP/native describe bytes, stale-epoch describe rejection, re-render refresh, and compile-time absence of describe from the m-10 interface.

### F2 - Credential sourcing and USE-capability claims conflict with the ratified m-1 boundary

The contract hoists ordinary-environment credential sourcing and the current file loader byte-for-byte (`:23,27-34`) while later claiming S-B never appears in any process argv/env (`:118-122`). The live loader reads before stat, follows symlinks, accepts owner bits other than exactly `0600`, performs no owner/regular-file/size/descriptor-identity check, and trims arbitrary surrounding whitespace (`frank/cmd/frank-mcp/main.go:24-46`). It is not the hardening contract this new broker can inherit.

The broker section also calls a current-epoch worker token leak-inert by borrowing m-1's credential-reference wording (`:77-83`). A USE capability that authorizes calls on an accessible private broker channel is secret-free and identity-inert, but possession is intentionally authorizing. The parallel m-1 review independently requires that split and says its source contract has not yet authorized the broker-private 0600 file as a persistent S-B sink (`step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-050629.md`, F1-F2).

Required revision:

1. Consume the revised m-1 contract rather than approving against `aa90fa45...`: distinguish the non-authorizing opaque credential locator from the authorizing epoch-bound USE capability, and state the latter's honest leak consequence.
2. Remove S-B from ordinary environment and argv paths. Pin one authorized broker-private custody sink only after m-1 names it; m-10 may name an opaque config home but cannot receive credential bytes, a credential path, or a locator that itself authorizes resolution.
3. Define descriptor-safe loading: no-follow open, regular file, expected owner, exact allowed mode, bounded size/grammar, read and verify from one descriptor, and no ambiguous trimming. Pin rotation replacement and deletion semantics.
4. Add negatives for env/argv/FD census, symlink, non-regular file, wrong owner/mode, oversized or malformed/newline-bearing content, path replacement races, reference-without-authority, and current-capability leakage/replay.
5. If seat-scoped legacy `frank-mcp` remains in Step-3, state whether it uses this hardened file/FD source or is explicitly outside S-B broker custody. It cannot silently retain `FRANK_CREDENTIAL` while the shared source claims the m-1 floor.

### F3 - Reconnect-once can cross the epoch barrier without a pinned disposition

The shared client retries once after a transport failure (`:27-32`). The broker admits a call once at entry and guarantees that after an epoch update no new prior-epoch invocation is possible (`:94-103`). The draft does not say what happens if attempt one fails or loses its response, the epoch advances, and `ManagedClient` reconnects and issues attempt two. That second conductor `tools/call` occurs after the update even though its outer broker call was admitted before it.

The live code retries after every `Client.Call` error, including server/protocol/context errors, not only failures for which reconnect is meaningful (`frank/cmd/frank-mcp/mcp.go:172-189`; `internal/channel/server.go:530-554`). Under concurrent callers, the proposed hoist also lacks a single-flight reconnect/client-swap rule. "Hoisted verbatim" therefore does not establish a safe broker transport contract.

Required revision:

1. Pin retry as part of the F64 linearization model. Preferred: each reconnect retry re-enters the epoch fence immediately before its send; if the admitted epoch is no longer current, do not retry and return a typed stale/unknown outcome that drives rediscovery. If the owner instead permits the retry as the same admitted invocation, narrow the post-update claim accordingly and record both transport attempts under the admitted operation.
2. Classify retryable transport failures exactly. Auth rejection, remote application/protocol rejection, frame error, and caller context cancellation must not trigger blind reconnect/replay.
3. Serialize connection replacement so concurrent failures cause at most one redial/re-auth generation and cannot multiply retries or detach the active push reader.
4. Add an epoch-advance-between-attempts fixture for all three verbs and describe, plus concurrent call/push/redial fixtures. For submit, prove content-hash replay; for every other operation, state the side-effect and unknown-outcome posture rather than relying on the submit comment.

### F4 - Restart bootstrap, token delivery, and the promised records lack one complete protocol

The fence says a restarted broker accepts its first epoch update as a re-seed (`:87-93`), m-10 re-attaches the live generation (`:111-116`), and the broker mints a per-generation token when m-10 attaches it (`:77-82`). The contract does not bind that re-seed to a run/generation/control-session identity, prove that a stale supervisor session cannot re-establish an old epoch, or state how the worker receives its token without m-10 carrying an authorizing capability. The parallel m-10 feed is scoped by `{run_id,generation_id,turn_epoch,state_seq}` and is queryable/durable, but the m-7 consumer reduces this to a bare monotonic epoch and a first-update exception.

The recording promise is also not realizable as written. Section 2.7 assigns records vaguely to "m-10's store / m-3's E0 events at the owners' grain" (`:111-116`). Yet complete-and-deliver requires every cross-epoch call to be recorded, and own-process placement intentionally lets an already admitted call complete while app-main/m-10 may be down. No writer, IPC message, durable-before-response ordering, acknowledgement, retry, or outage disposition is pinned.

Required revision:

1. Define separate authenticated control and worker sessions. On broker start: establish the current m-10 control session; query durable `{run_id,generation_id,turn_epoch,state_seq,lease_state}`; reject regressions/reuse; install current state; only then accept a direct worker attach and mint/deliver the capability on that worker connection. Capability bytes must never transit m-10.
2. Bind capabilities to at least the run, generation, epoch, and broker-instance/session nonce. Losing the current control feed invalidates admission/forwarding immediately; an old control connection or stale re-seed cannot revive authority.
3. Consume m-1's revised event/effect matrix for replacement, broker restart, conductor restart, re-mint, and re-mint concurrent with replacement. The broker contract must define the corresponding channel, epoch, capability, reconnect, and in-flight effects.
4. Choose one recording protocol. If m-10 is the durable writer, define canonical broker-event frames and require durable append acknowledgement before a response whose claim is "recorded" is delivered; state what an m-10 outage does to an already admitted call. If broker-local durability is chosen instead, route the new state owner/output family for review. A best-effort E0 event cannot satisfy a mandatory recorded disposition.
5. Add restart-with-stale-state, old-control-session, direct-token-delivery, control-loss, m-10-down-in-flight, and rotation-during-replacement fixtures.

### F5 - The F65 stamp neither proves loaded binary bytes nor pins the required canonical carrier

Section 3 hashes a deployed artifact path in-process and externally, then treats digest agreement plus stamp-writer provenance and a live socket as the running-service binding (`:134-157`). Both digesters can hash the same replaced or bystander path while an older mapped image is serving. A socket plus timestamp identifies liveness, not which bytes the kernel loaded. "The process actually serving" is therefore asserted rather than proven.

The F68 supplement explicitly dispatched canonical artifact identity, encoding, running-service loaded-proof, and the relay-leg evidence reference. The draft pins only one digest string and a field set, then defers exact filename, home, byte shape, member-marker encoding, and therefore the external observer's exact read/verification procedure to a build-lane round (`:159-162`). That is too late for the interface contract and makes FX-TB-12 non-deterministic.

Required revision:

1. Choose a process-bound executable identity mechanism and state its platform floor and residual. For example, bind the serving PID/socket to an executable handle or OS-reported loaded-image identity, hash through a pre-opened/non-replaceable descriptor with pre/post identity checks, and have the external observer verify the same process-to-artifact identity. A release/code-sign identity is also acceptable if it proves the exact loaded artifact at the claimed grain.
2. Keep the honest no-attestation ceiling, but do not call two hashes of one mutable pathname independent loaded-byte proof.
3. Pin the complete canonical stamp schema and serialization now: schema/version, field names and types, member-marker ordering/encoding, timestamp format, digest prefixes, unknown/missing-field behavior, exact canonical-byte algorithm, output location/name semantics, temp/rename/fsync behavior, and stale-file behavior on failed startup.
4. Pin publication relative to socket readiness and bind an instance identifier/PID/start nonce so the observer cannot join a current socket to a prior process's stamp. Describe conductor-down semantics without calling a stale last-start artifact "current" unless that is explicitly the intended status model.
5. Pin the exit-record evidence-reference encoding, including the exact relay IDs/record identities and how the observer verifies each reference. Correct "sole governed writer" wording: this is a conductor lifecycle output, not a governed-store record.

## Accepted portions

- **The F67 grill is complete and the placement decision is accepted.** The record compares both allowed placements across all dispatched criteria, selects the own-process broker without changing the ratified topology or F57 claim ceiling, and identifies no operator-owned fork (`:180-203`).
- **The broad transport/mapping ownership split is correct.** m-7 owns framing, channel lifecycle, typed transport outcomes, describe carriage, and push; m-2 owns FieldSpec mapping and re-render interpretation (`:14-56`). F1 closes the omitted metadata method rather than moving mapping into m-7.
- **Per-operation and per-push fencing is the correct host.** The conductor remains epoch-blind; the app-side broker is the single place that can gate the logical seat channel without changing conductor protocol/store bytes (`:85-109,172-178`).
- **Complete-and-deliver is acceptable for an operation admitted before epoch advance**, provided F3 pins reconnect attempts and F4 makes "recorded" durable and mechanically ordered.
- **The F65 scope split is correct.** Conductor identity belongs to the separate relay-exchange leg and must not be absorbed into m-3's app/provider E3 vector (`:129-131,153-157`). F5 requires exact producer proof, not a different evidence owner.
- **No new operator decision is needed** if the revision preserves own-process placement, the ratified F57 honesty ceiling, m-10's no-seat/no-verb/no-secret rails, and existing conductor protocol/store shape. A new secret sink, identity authority, or governed-store field must route back.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Carries typed describe/rediscovery through the same fenced broker abstraction without moving m-2 interpretation into transport.
2. Rebases custody and capability claims on m-1's revised approved semantics, removes ordinary-env S-B sourcing, and pins a hardened authorized sink.
3. Defines reconnect, epoch, push, and concurrent client replacement at one testable linearization model.
4. Defines restart/control/worker handshakes and a durable recording owner/order that survives the named outage windows honestly.
5. Produces a canonical, process-bound F65 stamp and exact relay-leg reference contract at design grain.
6. Preserves all accepted placement, boundary, route-back, and evidence-scope decisions.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. Consumer confirmations remain held until m-7 pair approval and until the upstream m-1 contract consumed here is pair-approved.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `bddf868ba050e201302ece0c63e9544b4f6402b0ad5302b49769d98f33818b2e`.
- Read the exact addressed relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260716-050007.md`, parent dispatch `DESIGN-orchestrator-planner-20260716-041630.md`, and supplement `...-043459.md`; routing, grill, F68 producer, and consumer set match.
- Checked `master/STEP-3-MVP-AMENDMENT.md` r7 and immutable `master/STEP-3-ARCH-AMENDMENT.md` for F57/F60/F64/F65/F66 and the app-shell/conductor boundary.
- Checked parallel owner bytes `master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` and `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md`, plus m-1's `DESIGN-REVIEW-implementer-20260716-050629.md` must-revise verdict.
- Checked live mechanism at `frank@502e06c`: `cmd/frank-mcp/main.go:24-46`, `cmd/frank-mcp/mcp.go:172-225,278-308`, and `internal/channel/server.go:418-589`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner revises only F1-F5, rebases the m-1 edge after its fresh pair-approved bytes, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; consumer confirmation remains held.
