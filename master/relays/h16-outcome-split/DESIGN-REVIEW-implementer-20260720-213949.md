## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev9

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r9
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-213611.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-213949.md
SUBJECT: must-revise - rev9 closes redo completeness and stale-byte scope, but quarantine clears before binding repair is guaranteed, the local administrative anchor has no pinned sole-writer transport/lifecycle, and operator-role eligibility is undefined when the disputed canonical pivots disagree

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev9 and its focused decision record at exact SHA-256 `921eb0587c9f0d1d35695a104cfd75e17c74b7e3e9942bd0f88a97b5731b9166`, parent relay SHA-256 `d89c5a37b9f894b27aae686a67655d41f31cbd102054c2bfbc933e12aebbce0c`, prior pair-review SHA-256 `b7589cf8d7b18afc20d6afb8f50718c57b9c45f58da51c5e8f04a27f5862e749`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`. The current VP `203905` F93 hold remains operative: H-16 still requires an approving pair review, both owner confirmations, and master/VP pass before stage-6/IMPL/T4 authority.

Rev9 closes R8-F1 in design direction: automatic legacy selection now requires exactly-once coverage of the canonical pivot set in raw redo, with anomaly/absence fail-closed, snapshot-once restart mechanics, and a distinct m-1 authority ruling. It closes R8-F3: the stale conflict/rejection/scope clauses are removed and the final confirmation scopes are explicit. Its typed quarantine and three-shape recovery split are also the right direction, but the new executable contract still has two runtime gaps and one unresolved canonical predicate.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, stage-6 lock, T4 action, or deploy.

## Findings

### R9-F1 - clearing quarantine at the resolving commit reopens the stale credential before binding repair is guaranteed

Rev9 says the accepted anchor/disposition updates derived chain status in the same serialized-loop step, repair then calls `MintOrReplace`, and quarantine "clears with that resolving state" (`2026-07-20-h16-outcome-split.md:35-37`). Canonical resolution and binding realization are not one atomic state transition on current bytes.

The loop commits first, updates/publishes tables through `completeTurn`/`AfterCommit`, and only afterward invokes `AfterAccepted` (`internal/engine/loop.go:156-186`; `cmd/frank/main.go:286-303`). The mint completion that atomically replaces the credential and `realized_mint_ref` happens inside that later callback (`main.go:603-621`). If quarantine is derived only from canonical chain status, the accepted anchor makes the chain non-conflicted before replacement succeeds. An auth attempt in that interval can match the old bound credential and observe a cleared quarantine. A repair error or crash after anchor commit preserves the same mismatch until recovery. This is the zombie window m-1's locked ruling forbids (`s6-fidelity-m1/.../102208.md:23-41,47`).

Required revision:

1. Define effective quarantine as unresolved/conflicted chain OR resolved canonical tip not yet realized by the binding row. Canonical resolution alone cannot clear it.
2. Pin the publication order and synchronization: the old credential must remain refused until `MintOrReplace` has durably replaced credential plus `realized_mint_ref`; auth must not interleave between binding replacement and quarantine-clear publication.
3. Keep the seat quarantined on repair error and across crash/restart until pre-serve recovery proves the binding realizes the selected tip. Do not report the anchor outcome as a usable recovery while derived completion failed.
4. Add a concurrent-auth leg plus crash/error cuts after anchor commit, after canonical fold publication, and before/after binding persistence. Every pre-realization attempt with the old credential must return `auth:seat-quarantined`; after realization it must return `auth:invalid-credential`.

### R9-F2 - the local administrative anchor names no realizable sole-writer transport or process lifecycle

Rev9 introduces "a new admin verb on the store host" that commits through the serialized loop and returns the new credential directly (`:38-42,67`). That does not yet choose a write path.

On current startup, recovery and pre-serve binding completion finish before the engine loop is constructed and started (`cmd/frank/main.go:259-330`). The existing `-mint` command is an offline/admin-time binding operation: it refuses while the conductor socket is live and has no canonical commit-loop path (`main.go:577-600`). Therefore:

- an offline anchor command has no running serialized loop or synchronous post-restart reply channel;
- an online command needs a distinct local administrative endpoint and an injection path into the one loop, because the ordinary channel requires the credential this case has quarantined; and
- direct `Store.Commit` from a second process while the conductor serves would violate the sole-governed-writer contract and race process-local store locking.

Required revision:

1. Elect one exact mechanism: an online local-admin IPC admitted outside seat auth and submitted into the existing engine loop, or an offline/bootstrap durable-intake ceremony with exclusive store ownership and a defined startup consumption/reply protocol.
2. Pin endpoint/path custody, OS access check, serving/stopped precondition, process ownership, command schema, loop/recovery injection point, idempotency key, and protected credential reply. State how the mechanism remains absent from seat `submit/project/read` surfaces.
3. Pin automatic pre-serve system-anchor insertion too: current `engine.Loop` starts after recovery, so say whether the migration runs in a serialized startup writer phase or startup ordering changes before claiming each anchor goes "through the serialized commit loop."
4. Add online/offline exclusion, competing-process, duplicate/replay, and crash cuts from admin intake through anchor commit, binding repair, and credential reply. Exactly one canonical anchor and one realized credential may result.

### R9-F3 - "operator-role AND quarantined" is undefined when the conflicted pivot chain carries disagreeing authority bits

The sole-operator admin path is accepted only when the target is operator-role and quarantined (`:41`). But `role` and `is_operator` are fields on every canonical `seat_mint` pivot (`internal/engine/submit.go:391-421`). A broken, forked, or unanchored multi-pivot chain can contain candidates that disagree on those fields. The whole reason for the anchor is that canonical bytes cannot yet identify the current pivot, so rev9 cannot use "the latest pivot's operator-role" before selection. The binding row's `SeatMeta.IsOperator` is derived mutable state and cannot decide canonical authority when accepted records disagree; accepted canonical records win over that row under the m-1 ruling.

The same ambiguity affects the claim that this is the "sole-operator" shape: rev9 does not name the source used to prove that no other non-quarantined operator credential exists. As written, the restriction can either trust stale binding state and grant an admin bypass, or fail closed and leave an authority-disagreement conflict permanently unrecoverable.

Required revision:

1. Define the canonical pre-anchor predicate for target operator eligibility and for "no other usable operator." Do not use the unresolved tip or mutable binding row as authority without an explicit m-1 ruling.
2. Define the recovery for candidate pivots that disagree on role/`is_operator`. It must remain reachable without turning the local admin verb into an unrestricted canonical-write bypass.
3. Carry this authority-disagreement case explicitly in m-1 scope (f), not only the statement that local store access equals operator authority.
4. Add role-flip fixtures for legacy and linked conflicts, with the binding row reflecting each competing candidate in turn. The same canonical history must produce the same eligibility/recovery result, and no case may become permanently unrecoverable.

## Accepted portions

- R8-F1 closes in direction: raw exactly-once coverage of canonical legacy pivots plus anomaly detection is the right safety predicate; every deficient/partial case fails closed; the one-time noncanonical evidence authority is explicitly reserved for m-1 confirmation.
- Snapshot-once migration, accepted-anchor idempotence, skip-on-restart, and redo-disposal gating are sufficient in direction once the exact startup writer locus is pinned.
- R8-F3 closes: `system-owned` is consistent, the stale bound-credential exception is gone, decision-record metadata is current, and final m-1/m-2 scopes are enumerated.
- The typed refusal leak rule is sound: only a matching bound credential can receive `auth:seat-quarantined`; probes remain `auth:invalid-credential`.
- The worker / second-operator / sole-operator recovery partition is complete as a case split, subject to R9-F2/F3 making the last path executable and canonically eligible.
- Rev8's accepted-only forged-header fold, canonical transition discipline, linked-chain tip predicate, header/version carrier, fresh-instance, Class-G, evidence, compatibility, projection, and consumer direction remain accepted.

## Gate disposition

MUST-REVISE is byte-bound to rev9 `921eb0587c9f0d1d35695a104cfd75e17c74b7e3e9942bd0f88a97b5731b9166`.

Before any H-16 IMPL branch or downstream stage-6/T4 release:

1. m-7 returns fresh design/decision-record bytes closing R9-F1..F3;
2. m-1 confirms the exact legacy-order evidence/anchor authority, effective-quarantine predicate, typed leak rule, role-disagreement eligibility, and local-admin authority basis, while m-2 confirms the exact headers/rejection bytes/classes, transition dispositions, anchor/disposition schemas, and local-recovery form;
3. a fresh exact-byte pair review passes; and
4. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy either owner confirmation or the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Keep the stale credential quarantined until the selected canonical tip is durably realized in the binding row.
2. Choose one concrete local-admin transport/lifecycle that preserves the one governed writer and has a crash/replay-safe credential reply.
3. Define operator eligibility from canonical truth even when disputed pivots disagree on authority fields.
4. Preserve rev9's completeness proof, partial-redo fail-closed rule, accepted-only forged-header fold, typed leak distinction, canonical transition discipline, linked-chain predicate, and final owner/master/VP sequence.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-213611.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `921eb0587c9f0d1d35695a104cfd75e17c74b7e3e9942bd0f88a97b5731b9166`; parent relay hash recomputed: `d89c5a37b9f894b27aae686a67655d41f31cbd102054c2bfbc933e12aebbce0c`; prior review hash recomputed: `b7589cf8d7b18afc20d6afb8f50718c57b9c45f58da51c5e8f04a27f5862e749`.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: loop commit/AfterCommit/AfterAccepted ordering, startup recovery/loop/server ordering, binding replacement, channel auth, genesis admin mint, and canonical `seat_mint` authority fields.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-213949.md`.
Next requested action: m-7.planner folds R9-F1..F3, obtains the exact m-1/m-2 confirmations after a passing pair review, and returns fresh design/decision-record bytes for a new uniquely-parented review; H-16 IMPL and downstream stage-6/T4 authority remain held.
