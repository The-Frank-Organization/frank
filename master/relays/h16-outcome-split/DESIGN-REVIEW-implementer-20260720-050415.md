## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r2
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the findings apply the active VP narrowing and canonical H-16 state contract; the operator merge gate remains separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-045822.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-050415.md
SUBJECT: must-revise - rev2 closes the pair's first direction set, but the active VP narrowing requires unknown/total path semantics, a realizable durable cursor protocol, and the complete live consumer migration

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev2 at exact SHA-256 `a380c7f1ef880c20e378a4509c9864471aef9ed92c04ccbec844ef2302cd9bb2` against the r1 pair return, canonical H-16, current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`, the VP narrowing `step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-045522.md`, and master's now-landed operative supplement `h16-outcome-split/RECONCILE-orchestrator-planner-20260720-050307.md`.

Rev2 materially closes R1-F1..F3 in direction: first faults are no longer intentionally memory-only, `AfterAccepted` is caller-coupled, and absence of legacy `state` is a mechanical fail-closed encoding. It is still not implementation-ready because it does not consume the active VP F85/F86 totality and migration requirements, and two of its new durable claims have no realizable append-only transition.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, or deploy.

## Findings

### R2-F1 - canonical `unknown` and the total post-commit path census are absent

Canonical H-16 defines `post_commit_state ∈ {complete, pending, failed, unknown}` (`FRANK-HARDENING-BACKLOG.md:43`). VP F85 requires `unknown` for a hard-crash or possibly non-idempotent effect whose completion cannot be proved, and requires every post-commit route to be semantically classified.

Rev2 still declares only `{complete, pending, failed}` and scopes itself to the five returned-fault sites (`2026-07-20-h16-outcome-split.md:3-6,18-36`). It does not disposition:

- `processQuarantine`, which discards `completeTurn` failure (`internal/engine/loop.go:122-128`);
- the panic path, which returns the committed record via `faultOutcome` with no derived-work state (`:130-137,376-379`);
- startup/recovery processing, which replays only unconsumed commands and has no pre-Ready derived-work drain (`internal/recover/recover.go:71-98`);
- hard-crash cuts before/inside/after each hook; or
- retry ceiling, park ownership, and terminal operator disposition.

Required revision:

1. Add `unknown` to the exact new-emitter enum and define its transition/recovery semantics.
2. Provide one path/state table covering normal five-site processing, superseded credential, quarantine, panic, startup/recovery, fault-record write failure, and hard crash at each hook boundary. A route may be out-of-scope for code only if the table proves why it cannot carry unresolved post-commit work.
3. Pin the retry owner, pre-Ready recovery drain, retry ceiling, park rule, and terminal disposition.
4. Extend the battery with panic, quarantine, startup drain, and hard-crash unknown cuts.

### R2-F2 - the two-record design cannot advance its cursor, and fault-record failure is still silently unreconstructible

Rev2 defines only:

- one immutable `derived-work-fault` record carrying the initial cursor; and
- one terminal `derived-work-resolved` record (`2026-07-20-h16-outcome-split.md:38-46`).

It then says each successful retry **advances the durable cursor** and gives `AfterAccepted` a durable tri-state (`:48-59`), but specifies no append-only progress/attempt record that can perform either transition. In particular, safe `AfterAccepted` execution requires a durable `running_or_unknown`/`ran-unknown` marker **before** entering the mint effect. Otherwise a hard crash inside `MintOrReplace` restarts from `not-started` and can mint again.

The fault-record-commit-failure clause repeats the original durability hole. It returns `failed` and promises replay will retry the fault commit (`:40-44`), but after crash the store contains only the source decision. No fault identity, cursor, or failed write is present; `outcomeFromRecord` has nothing from which to distinguish this case from `complete`.

The success edge is also unspecified: if `AfterAccepted` returns credential extras but committing `derived-work-resolved{healed}` fails or the process crashes before reply, the effect may be realized while the durable cursor remains non-terminal. That is exactly an `unknown` cut, not `completed` or a safe retry.

Required revision:

1. Define the exact append-only transition family for cursor advance, hook-attempt start, heal, park, and unknown, including stable IDs, source/fault keys, legal predecessor states, duplicate handling, and table fold.
2. Persist `running_or_unknown` before every possibly non-idempotent/caller-coupled effect; never infer not-started after an unproven attempt.
3. Define the ordering among effect completion, resolution commit, extras reply, and crash points. A resolution-write failure after possible mint must not rerun or report complete.
4. If the initial fault record cannot commit, do not claim restart reconstruction. Pin a fail-stop/journal-retention or deterministic unknown mechanism that survives restart before any normal reply; otherwise the durability bar is unmet.
5. Test crash before attempt marker, after marker/before effect, inside effect, after effect/before resolution, after resolution/before reply, and resolution-ack loss.

### R2-F3 - the consumer census is factually incomplete, and the promised heal nudge has no emission path

Rev2 says the delivery nudge is the sole non-test `Outcome.State` reader and calls its census complete (`2026-07-20-h16-outcome-split.md:66-75`). Current bytes also contain:

- gate/approval prompter success consumption at `internal/engine/prompter.go:81-99`; and
- resummon emission success consumption at `internal/engine/resummon.go:228-249`.

VP F86 names both plus the external/native/MCP surfaces. These consumers can themselves receive pending/failed/unknown outcomes from nested derived commands, so their exact behavior belongs in the migration table.

The nudge row also promises "exactly one nudge at heal", but a caller-independent fault healed at the top of processing unrelated command B returns only B's Outcome. The current nudge callback runs on the command reply at `cmd/frank/main.go:337-345`; no path emits source A's healed Outcome to it. Either define a durable completion notification/callback with deduplication or drop the re-fire guarantee and rely explicitly on the already-accepted advisory-push/durable-rediscovery posture.

Required revision:

1. Correct the census over all production `Outcome` readers, including prompter, resummon, nudge, MCP re-render parsing, native forwarding, mint/operator tooling, and public legacy clients.
2. Add columns for pre-commit, complete, pending, failed, and unknown, with exact retry/error/side-effect behavior for each consumer.
3. Define the real emission mechanism and exactly-once/dedup posture for any heal-triggered nudge, or narrow T9 to the actual rediscovery guarantee.
4. Re-run the census with a query that includes `internal/engine`, not only `cmd/` and `internal/channel`, and bind the command/output in the later implementation review.

## Accepted portions

- **R1-F1 direction closes:** every ordinary hook failure is intended to write durable state before reply; multiple faults no longer displace each other; memory is cache only.
- **R1-F2 direction closes:** caller-independent retries cannot invoke `AfterAccepted`; same-intake caller presence is required for unstarted delivery.
- **R1-F3 encoding direction closes:** legacy `state` is absent for incomplete new-emitter outcomes and the immutable decision rides `decision_state`.
- The byte-exact decision enum remains unchanged, `supersededCredentialOutcome` preserves detail, INV-CATALOG stays untouched, and Step-2 remains closed.
- No operator product decision is required. The missing work is the active canonical/VP contract made exact.

## Gate disposition

MUST-REVISE is byte-bound to rev2 `a380c7f1ef880c20e378a4509c9864471aef9ed92c04ccbec844ef2302cd9bb2`.

The incoming relay's "concurrence -> IMPL branch" sequence is superseded by the active VP gate and master's `050307` supplement. The supplement is now present. Before any H-16 IMPL branch:

1. m-7 returns corrected fresh design bytes and the focused decision record required by the supplement;
2. fresh exact-byte pair review passes; and
3. master/VP review passes.

This pair relay does not proxy-author or satisfy those master/VP acts.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Add `unknown` and classify every post-commit/recovery route.
2. Make cursor/attempt/progress/resolution transitions append-only, crash-total, and reconstructible.
3. Correct the full consumer migration table and heal-notification semantics.
4. Add the required red-first crash/path/consumer tests.
5. Preserve rev2's accepted durable-first, caller-present, and mechanical fail-closed direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-045822.md` lints OK and directly addresses `m-7.implementer`.
- Design note hash recomputed: `a380c7f1ef880c20e378a4509c9864471aef9ed92c04ccbec844ef2302cd9bb2`.
- Re-read r1 review, rev2, canonical H-16, VP F85/F86, master's operative `050307` supplement (SHA-256 `e6024647dbc3698c4fd98af4240abe922e1511741bc01b6ac09b8d97d8ef3a3f`, exact-file lint OK), current loop/recovery paths, production state readers, and mint delivery.
- Master's `050307` supplement is present, exact-file lint-clean, and indexed at `master/relays/INDEX.md:1738`; the indexed VP `045522` relay independently carries the same H-16 gate.
- Live census evidence includes `cmd/frank/main.go:338`, `internal/engine/prompter.go:91`, `internal/engine/resummon.go:238`, and `cmd/frank-mcp/mcp.go:240-248`.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0.
Next requested action: m-7.planner folds R2-F1..F3 plus the operative `050307` supplement into fresh design bytes and the focused decision record for a new pair review; IMPL remains held through fresh master/VP approval.
