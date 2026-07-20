## DESIGN-REVIEW — m-9 lifecycle half r1 full-document review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the findings are contract defects with bounded technical resolutions; no operator product decision is required
GRILL_REQUIRED: no — this lifecycle-half review does not pre-empt the stage-4 full-worker grill
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: b4e08545a90267d3e1c646d0e1b13c9afc86218e83959ca42ad5fc56b494d29e
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260717-124600.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-7.planner, m-1.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
SUBJECT: must-revise — r1 has seven lifecycle contract defects: restart reader gap, pre-translate attempt gap, contradictory denial terminal, incomplete F70 attach seam, EOF tool-child residual, incomplete K6 surface census, and a stale m-10 body hash

DESIGN_REVIEW_VERDICT: must-revise

I reviewed the entire r1 document `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md`, not merely its refresh delta. The reviewed bytes reproduce SHA-256 `b4e08545a90267d3e1c646d0e1b13c9afc86218e83959ca42ad5fc56b494d29e`. All seven declared basis hashes were checked against the named files or their preserved revision provenance; the three refresh hashes reproduce exactly: m-7 r8 `ab0ed428…`, m-10 r12 `111ab95a…`, and m-3 r3 `70838f83…`.

The consume-before-execute core, duplicate/stale/mismatch rejection posture, one-active-turn distinction, counter-string rule, durable-rediscovery claim boundary, and exact-turn/exact-lane replay custody are directionally sound. Approval is blocked by the seven findings below. F1–F6 affect executable lifecycle semantics; F7 fails the requested no-stale-hash sweep.

## Blocking findings

### F1 — A replacement worker has no reader path for the durable state it must reconstruct

Section 2.6 says a new generation rebuilds its cache from committed conductor records plus m-10's durable `turns` / `provider_attempts` / `tool_calls` / ticket rows and "`m-10's assign`" (`...mvp-lifecycle-half.md:109-110`). Section 3.4 then requires the replacement to recognize a pre-existing `UNKNOWN_TOOL_OUTCOME` as terminal-for-that-call (`...:160-163`).

The consumed m-10 r12 seam makes its SQLite file private and exposes state only through frames (`...mvp-ipc-manifest-seam-contract.md:218-220`), while CTRL-W `assign` contains only `{run_id, turn_epoch, manifest_digest, generation_id, broker_worker_endpoint}` (`...:60`). No restore/snapshot/query frame carries the parked turn, attempt, tool-call, or ticket dispositions. Conductor `project` / `read` cannot recover m-10-private rows. Therefore the stated rebuild and the required `UNKNOWN_TOOL_OUTCOME` handling have no reader.

Required revision: pin a state-only restart handoff on an owner-approved m-10 seam (for example, a bounded post-assign restore frame/query containing the exact non-payload dispositions and identities), or narrow the recovery model so the replacement never claims to inspect m-10 rows and prove how every required parked disposition reaches its consumer. Route any new CTRL-W type to m-10; m-9 must not invent it unilaterally.

### F2 — `replay_scope_violation` and other pre-translate rejects have no total attempt disposition

m-9 emits `attempt_open` before the DATA-P request and says every invocation becomes an m-10 `provider_attempts` row (`...mvp-lifecycle-half.md:70-87`). The new §2.8 says `replay_scope_violation` becomes a typed attempt failure in the existing disposition set (`...:115-120`).

But m-8 r2/r3 defines `replay_scope_violation` and `malformed_request` as DATA-P rejects with **no attempt opened and no freeze** (`...mvp-provider-contract.md:44-47,125`), while its claimed-total CTRL-C `attempt_result` table has no local/pre-translate-reject member (`...:74-84`). Under m-10 r12, an m-9-only `stream_failed` view with the m-8 view absent parks `UNKNOWN_PROVIDER_OUTCOME` (`...mvp-ipc-manifest-seam-contract.md:61`). A deterministic zero-wire local rejection would therefore masquerade as an unknown provider outcome.

Required revision: reconcile the zero-attempt reject path across m-8/m-9/m-10. Pin whether m-9 refrains from opening an m-10 attempt until DATA-P acceptance, or extend the owner contracts with a total local-reject disposition and matching reconciliation. Add a fixture proving replay-scope and malformed-request rejects yield zero wire, no false UNKNOWN, and one unambiguous turn disposition.

### F3 — `turn_denied` contradicts the executor's denial behavior and its wire history

The terminal table says above-ceiling/absent authorization ends the turn "before a wire send" with zero/partial-zero wire (`...mvp-lifecycle-half.md:89-100`). That authorization request exists only after a provider attempt has already returned a complete `tool_call_end`, so provider wire activity has already occurred. Section 3.3 also says `DENIED_ABOVE_SET` returns a typed tool-error result to the model and merely counts toward the tool-call bound (`...:151-158`); §2.4 says repeated denials terminate only when the bound exhausts (`...:102-103`).

Required revision: make the state machine total and non-contradictory. Separate pre-wire egress denial from post-wire tool-authorization denial, state whether a single tool denial continues the same turn or terminates it, and reserve `turn_exhausted` for the bounded-loop terminal if that remains the rule. Every terminal must state provider-wire and tool-effect facts independently; "partial-zero wire" is not an honest disposition.

### F4 — The F70 attach behavior is not contract-real at either side of admission

Sections 1.2/1.6 rely on a typed distinction between "attach refused while suspended" and "stale-tuple refusal" and promise the hold is supervision-visible (`...mvp-lifecycle-half.md:35-38,52-60`). m-7 r8 does require refusal until installed state and tuple equality on attach, but its owner bytes do not pin an attach reply/error shape that distinguishes these two cases (`...step3-mvp-transport-broker.md:203-216`); the named `broker:suspended` / `broker:stale-epoch` taxonomy is pinned for post-capability operations, not for attach.

The other half is also absent: m-10 r12 sends `assign` after lease bind and says attach precedes first admission (`...mvp-ipc-manifest-seam-contract.md:60,92,104`), but its CTRL-W set contains no attach-ready/attach-held acknowledgment. `hello` already makes the candidate READY, so m-10 cannot observe that the worker actually acquired the capability before issuing `turn_open`, nor distinguish bounded F70 hold from a wedged attach.

Required revision: route a reciprocal seam correction. m-7 must pin the attach result taxonomy (including suspended/no-installed-state versus tuple mismatch), and m-10/m-9 must pin how successful/held attach becomes observable and gates first `turn_open`, including timeout/fault disposition. Add suspended→installed and stale-tuple negatives; neither arm may self-advance or busy-loop.

### F5 — CTRL-W EOF does not contain an already-running tool subprocess

The design explicitly launches tool subprocesses (§1.3), then claims CTRL-W EOF makes the worker immediately exit and "executes no tool" (`...mvp-lifecycle-half.md:40-41,105-107,173-178`). Exiting a POSIX parent does not by itself terminate or reap an already-running child. An authorized `bash` or other tool process can therefore continue effects after the worker has exited and after m-10 has fenced/replaced its generation. The current fixture observes only the worker verbs; it does not prove the effecting process is gone.

Required revision: define the executor's in-flight-child EOF discipline at contract grain — process-group/job ownership, parent-death/kill behavior where available, bounded termination and reap, and the honest `UNKNOWN_TOOL_OUTCOME` / `PARTIAL_TOOL_EFFECT` fallback if containment cannot be proven. The replacement rule must state whether a continuing effect blocks overlapping successor execution. Extend the EOF fixture with an effecting child that would leave a delayed sentinel unless containment succeeds.

### F6 — The K6 "no durable source / every m-9 surface" claim omits crash and diagnostic persistence

Section 2.8 correctly excludes CTRL-W, m-10 rows, E0, SITREPs, and ordinary logs, but it also claims no durable source carries the opaque replay payload and that every m-9 surface is covered (`...mvp-lifecycle-half.md:115-120`). An in-memory envelope can still be persisted through a core/crash dump, heap profile, trace, panic diagnostic, or debug snapshot. None is excluded or scrubbed, and the fixture list has no replay-payload canary census. This matters precisely on the crash/park exits the new row says destroy custody.

Required revision: either disable/scrub those diagnostic surfaces while replay payloads are resident or narrow the claim and record the residual. Add a planted replay-payload canary test over logs/errors/traces/profiles/core-or-crash dumps/debug snapshots plus the already-named frames/store/events/relays, and prove cache destruction at every terminal, park, cancellation, stale-epoch exit, and channel-fault exit.

### F7 — The normative reciprocal section still cites the superseded m-10 hash

The r1 header correctly rebases m-10 to r12 `111ab95a…` (`...mvp-lifecycle-half.md:9`), but §5 still says "This half consumes m-10's `79fcf742…` lifecycle half" (`...:170-171`). That directly fails the dispatch's requested "no stale-hash citation survives in the body" check.

Required revision: replace the normative §5 citation with the exact r12 hash `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`. The historical "rebase from" occurrence in the header may remain explicitly historical.

## Confirmed non-blockers

1. The declared target hash and the m-7/m-10/m-3 refresh hashes reproduce. m-1 and m-2 also reproduce. The m-8 r2 seam semantics are preserved by current r3's documented provenance-only rebase; this review does not demand a semantic redesign merely because the live m-8 file now has the r3 hash.
2. F59's normal path is correctly ordered: complete validated call → observe → digest → authorize → consume → re-digest at the executor boundary → execute → record actual invocation identity. Duplicate consume, stale epoch, and identity mismatch all fail closed.
3. The §2.7 canonical-decimal-string rule covers the trust-bearing counter surface currently emitted by m-9. No uncovered m-9-owned counter was found.
4. The push claim matches F61's narrow guarantee: best-effort advisory notification, durable truth, recovery at the next startup/reconnect/poll. No periodic or at-least-once guarantee is silently asserted.
5. The replay envelope is otherwise correctly limited to exact turn + exact lane and copied as a unit. F2 and F6 are the remaining disposition/custody holes, not a rejection of the envelope design.

## Verification

- `sha256sum`/`shasum -a 256` over the lifecycle half — exact `b4e08545a90267d3e1c646d0e1b13c9afc86218e83959ca42ad5fc56b494d29e`.
- Reproduced bases: MVP amendment `2f75f2a1…`; m-10 r12 `111ab95a…`; m-7 r8 `ab0ed428…`; m-1 `7c8b09a6…`; m-2 `83d8e63e…`; m-3 r3 `70838f83…`; current m-8 r3 `6c586f35…` with its r2 `dc85fc01…` provenance explicit in the live document.
- Read the complete target with line numbers and cross-checked m-10 §§A.2/B.1/B.2/B.3/B.4/D/F, m-7 §§2.4/2.10, m-8 §§1.1–1.4, and amendment F59/F61.
- Targeted sweeps covered stale hashes; attempt/replay dispositions; restart/restore readers; EOF/subprocess residuals; crash/debug persistence; terminal-denial wording; and all counter-emission surfaces.
- `git -C frank rev-parse HEAD` — `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`; `git -C frank status --short` — clean before this review relay.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no `frank/` source/test edit, no PLAN, no IMPL, no credential/provider action.
FINAL_GIT_STATUS_SHORT: cwd root is not a git repo; `frank/` clean at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
Next requested action: m-9.planner folds F1–F7 in one bounded r2 revision, routes the F2/F4 owner deltas to m-8/m-10/m-7 as needed, then requests a fresh uniquely-parented full-byte m-9.implementer re-review. No closure SITREP or m-10 reciprocal approval should issue on `b4e08545…`.
