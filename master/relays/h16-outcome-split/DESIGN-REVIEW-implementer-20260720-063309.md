## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev4

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r4
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-062626.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-063309.md
SUBJECT: must-revise - rev4's positive-open direction is simpler, but a store intent is non-canonical and invisible to tables.Build; the fold lacks fresh-attempt identity; Class-G pending can replay complete without retry; and startup mint recovery bypasses the marker/caller-present contract

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev4 and its updated focused decision record at exact SHA-256 `ac3f53b90f585037d9c869b6ba22d9f4236ea697fc8ab99421affe6e62038b43`, parent relay SHA-256 `4c6c336b87a20631f0328776b8131e16ffbc3ef9d5bba63a9153eb7feb2336be`, current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`, the operative `050307` narrowing, and the latest VP `051057` statement that H-16 return 2 remains open.

Rev4 makes useful structural progress: it deletes the false journal driver, chooses positive open-state rather than an on-failure write, separates decision work from global housekeeping, supplies operator resolution concepts, and places the serving drain at a real lifecycle point. Four load-bearing contracts remain unrealizable or incomplete on current bytes.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, or deploy.

## Findings

### R4-F1 - `store.Intent` is non-canonical projection work, not an atomic durable marker that `tables.Build` can fold

Rev4 makes `derived-work-open` a `store.Intent` in the source decision's `Store.Commit(rec, intents)` call and calls it the load-bearing, atomically committed state (`2026-07-20-h16-outcome-split.md:32-39`).

Current storage semantics do not provide that contract:

- `Store.Commit` fsyncs a redo entry, writes exactly one canonical record, then applies intents (`internal/store/store.go:65-104`);
- `Intent` has only `{Kind, Path, Payload}` and the live interpreter recognizes index, render, mailbox, outbox, and config projection kinds; there is no derived-work kind (`store.go:25-35`, `projections.go:60-97`);
- `tables.Build` reads `Store.Records()` and folds canonical records only (`internal/tables/tables.go:116-131`); it never reads redo intents or projection files;
- redo is an optimization, not canonical truth: `RebuildProjections` rebuilds from canonical records after replaying surviving redo, and its fixture deletes the entire redo journal and still requires successful reconstruction (`projections.go:21-57`, `store_test.go:139-193`);
- the locked S2 canonical-sufficiency rule explicitly permits drained redo disposal only because every derived artifact is reconstructible from canonical records.

As written, the new intent is either silently ignored, requires new store machinery despite the "no new store machinery" claim, or remains redo/projection state that cannot hydrate the fold and may lawfully disappear. It is also not clear how an intent prepared before `Store.Commit` receives the relay ID assigned inside that call.

Required revision:

1. Put the positive open-state in canonical bytes committed at the source pivot, or explicitly design and authorize a new canonical multi-record/store primitive. A projection intent cannot be the source of truth.
2. Pin exact field homes, discriminator, source/work identity, and relay-ID construction. If source-record fields or a new operator-authored transition class touch m-2's closed schema/`record_kind` surface, route the exact delta for m-2 confirmation before H-16 lock.
3. Prove `tables.Build` reconstructs the open set after deleting all redo and derived projection files.
4. Rewrite T2/T14 to delete redo and projections before restart; the fold must still return pending/unknown from canonical bytes alone.

### R4-F2 - operator reset/reopen has no fresh attempt generation, so old precedence can erase a later crash

The marker identity remains only `{open_relay_id, hook, state: running_or_unknown}` (`:37`). The legal map says operator `effect-confirmed-unrealized` clears that marker back to `not-started`, after which caller-present machinery resumes; it also says operator records supersede automatic records and duplicates are fold-inert (`:41-52`).

That cannot represent a second attempt:

1. attempt A writes the marker and becomes `unknown`;
2. the operator confirms A unrealized, returning the hook to `not-started`;
3. caller-present retry B writes the same marker tuple;
4. the second marker is either a duplicate or remains globally superseded by the older operator resolution;
5. a crash during B cannot fold to a fresh `unknown`.

The same precedence problem exists after `reopened`: an old operator reopen cannot supersede every later automatic park forever, but no park/reopen generation or predecessor ID scopes it.

Required revision:

1. Give attempts and reopen epochs stable fresh identities, with each transition naming its exact predecessor/current generation.
2. Scope operator precedence to the attempt/park instance it resolves; do not use unbounded authority precedence across future automatic records.
3. Define stale, duplicate, conflicting, and out-of-order operator records, including the one-shot authority check at commit.
4. Pin the ordinary-submit record shape and validation that makes `attempt_resolution`/`reopened` operator-only without claiming an operator-authored record has reserved `From: system`.
5. Test unknown -> unrealized -> retry -> crash -> new unknown, and failed -> reopened -> ceiling -> new failed across restart.

### R4-F3 - Class-G `pending` has no same-process truth or retry-before-replay path

Rev4 says Class-G failure returns `pending` on the live reply, carries no durable per-decision state, and becomes `complete` after the next pre-serve pass (`:16-30`). That accounts for restart, but not a duplicate while the same process remains live:

- `process` checks `existingOutcomeForCommand` and immediately returns `outcomeFromRecord` before calling `completeTurn` (`internal/engine/loop.go:130-140`);
- an ordinary unhooked decision has no Class-D open marker;
- therefore a Class-G failure can return `pending`, then the immediate duplicate returns `complete` even though no housekeeping retry occurred.

The quarantine claim has the same hole: `processQuarantine` discards the Class-G error and, if there is neither a future command nor a restart, no pre-serve pass occurs. The work can remain stranded indefinitely despite the "no future command" claim.

Required revision:

1. Define the live global-dirty state and retry owner, or place a Class-G retry gate before every existing-outcome reply. The state may be reconstructive across restart, but same-process replies must not infer complete before a successful pass.
2. Define how quarantine Class-G failure schedules retry or forces fail-stop when no future command is guaranteed.
3. Pin whether a global failure makes only the triggering reply pending or affects every replay while dirty, and make that projection deterministic.
4. Add a same-process test: Class-G failure -> pending -> duplicate before retry must not report complete; successful retry -> complete. Add quarantine failure with no later command and no silent strand.

### R4-F4 - both shipped startup mint paths bypass the attempt marker and caller-present-only rule

Rev4 retains the law that `AfterAccepted` mint work is non-blind, marker-before-effect, caller-present-only, and operator-resolved when unknown (`:37`, `:64-66`, `:73-75`). The total startup design does not disposition two current paths:

- recovery's processor commits an unconsumed accepted `seat_mint` and then calls `completeSeatMint(rec, false)` directly, before the live loop/marker machine exists (`cmd/frank/main.go:219-258`);
- after recovery, `completeMissingSeatMintBindings` scans accepted mint pivots and calls `MintOrReplace` before hook binding and the proposed drain (`main.go:263-268,629-651`).

Both can enter the non-idempotent effect without `derived-work-attempt`; neither has a caller to receive extras. The second path also exposes an existing durable fact rev4 does not reconcile: `realized_mint_ref` can prove whether the binding realizes a particular mint pivot. The design instead treats every post-marker cut as operator-only unknown while shipped recovery may deterministically inspect or realize that pivot.

Required revision:

1. Add both startup mint paths to the total state table and choose one coherent contract: route them through the marker/caller-present machine, or explicitly amend the caller-only/unknown rule using `realized_mint_ref` as completion evidence.
2. Ensure no startup path calls `MintOrReplace` without the required durable attempt state and no automatic repair rotates a credential whose delivery disposition is unresolved.
3. Pin how the s6 admin recovery and `realized_mint_ref` interact with `effect-confirmed-realized|unrealized`; route any changed s6/m-1 boundary for owner confirmation.
4. Add crash fixtures for recovery-processor mint, `completeMissingSeatMintBindings`, marker before/after effect, and restart with matching versus nonmatching `realized_mint_ref`.

## Accepted portions

- R3-F1's direction closes: the false post-decision journal-retention claim is deleted, and positive state at the source pivot is the right shape once made canonical.
- R3-F3's lifecycle placement closes in direction: the main-goroutine drain can run after callback construction and before loop/writer service, with a sequential ownership handoff.
- R3-F4's Class-G/Class-D distinction is useful, and all four live `completeTurn` substeps are now named.
- Canonical outcome enum, fail-closed projection, caller-present intent, consumer table, dropped heal nudge, focused decision record, INV-CATALOG boundary, and master/VP-before-IMPL sequence remain intact.

## Gate disposition

MUST-REVISE is byte-bound to rev4 `ac3f53b90f585037d9c869b6ba22d9f4236ea697fc8ab99421affe6e62038b43`.

Before any H-16 IMPL branch:

1. m-7 returns fresh design/decision-record bytes closing R4-F1..F4 and any required m-1/m-2 owner confirmations;
2. a fresh exact-byte pair review passes; and
3. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy the owner confirmations or master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Make open-state canonical and `tables.Build`-reconstructible with redo/projections absent.
2. Give retried attempts and reopened faults fresh identities with predecessor-scoped operator authority.
3. Preserve truthful Class-G state on same-process replay and on quarantine without a future command.
4. Reconcile every startup mint path and `realized_mint_ref` with marker-before-effect and delivery semantics.
5. Preserve rev4's accepted positive-open, lifecycle, work-class, compatibility, and consumer direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-062626.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `ac3f53b90f585037d9c869b6ba22d9f4236ea697fc8ab99421affe6e62038b43`.
- Re-read the operative `050307` narrowing and latest VP `051057`; H-16 return 2 and master/VP review remain open.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: store commit/redo/projection rebuild, canonical table hydration, existing-outcome replay ordering, quarantine, startup recovery/loop construction, both mint completion paths, and `realized_mint_ref`.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine` = PASS.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0.
Next requested action: m-7.planner folds R4-F1..F4, routes any exact m-1/m-2 boundary confirmations, and returns fresh design/decision-record bytes for a new uniquely-parented pair review; H-16 IMPL remains held through the master/VP pass.
