## RECONCILE -- REVISE: H-16 is correctly PRE-T4 but not implementation-ready; H-17 inputs and the T4 exit bar require correction

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-external-audit-disposition-review-r1
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- this relay identifies bounded contract inconsistencies; the operator retains the H-16 merge grant and the later interface-lock/exit-test ratifications
GRILL_REQUIRED: no -- this is an adversarial review; the corrected H-16 API/recovery choice requires its own focused decision record before IMPL
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260720-043351.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-7.planner, m-7.implementer, m-8.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- concur with H-16 PRE-T4 timing and the additive-only H-17 authority boundary, but block H-16 IMPL until durability and caller compatibility are contract-real; correct the incomplete census schema and narrow the T4 proof claims to the ratified governance boundary

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260720-043351.md` at SHA-256 `1964b2c600326159a371add3ed2ed29ef8aaa109d845159c61afc85ea69955cd`, including:

- H-16 dispatch `h16-outcome-split/DESIGN-orchestrator-planner-20260720-043321.md` at `3d2a6430982f1b3c08a6eaaa81636ae27e57c8b721b5b1e0fc175e9052102413`;
- m-9 H-17 supplement `step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-043331.md` at `98cf53718f3edb139a15dae355fc3297b43c8f4fffc024f66071d88a854d0d07`;
- m-10 H-17 supplement `step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-043341.md` at `eb7c672606eed16fc8e3b9fd60d3eec0ea0735554d7c2857f9c92c7cff2514d3`;
- H-16..H-25 and the T4 exit shape now recorded in `master/FRANK-HARDENING-BACKLOG.md`.

The downstream bytes were also reviewed because the target explicitly asks for narrowing before work runs ahead: m-7 H-16 design r1 `eb7a4bd3973991666360b701429c94a13cfb34f0ded98ada4abb8a3552014969`, its pair MUST-REVISE return `0d18639897f6aeffe47de2a4bcbcc085f32d45d1553c845cbd6094ef6701066a`, current m-9 stage-4 draft `fdc6aaf0d0f4d6d897f34dbbb6994cdb0a45ba6d2fe82feee4d6c0ca7ff97345`, and current m-10 stage-5 draft `d47639973458d481fd17c36cab45667dea8db8eec5e815ae960adcac0e48af23`.

## Findings

### F85 -- BLOCKER before H-16 IMPL: the dispatched path/state/recovery contract is not total

The monotonic split is correct: a committed `{accepted,rejected,held}` decision must never be relabeled by post-commit work. The issued lane is nevertheless too narrow and internally inconsistent:

1. `loop.go:279-280` is not a `callHandler` path. `callHandler` is pre-commit at `:210-249`; `:279-280` is the separate post-commit `completeTurn` branch in `supersededCredentialOutcome`. The normal path has four post-commit sites at `:168-185`, making five returned-fault sites total.
2. The total census must also disposition `processQuarantine` (`:122-128`), which discards `completeTurn` failure, the panic path (`:130-134,376-379`), which serves the committed decision without any derived-work state, and startup/recovery paths. These need not all change, but none may remain semantically unclassified.
3. The dispatch defines `post_commit_state {complete,pending,failed}` while canonical H-16 requires `{complete,pending,failed,unknown}` (`FRANK-HARDENING-BACKLOG.md:43`). `unknown` is required for a crash after a possibly non-idempotent effect when completion cannot be proved.
4. The current design r1's first retryable fault exists only in an in-memory slot. A crash loses the slot and hook cursor; restart then derives `complete`. Its own T2 crash/restart test cannot pass under that mechanism. The m-7 pair independently caught and blocked this as R1-F1.
5. Existing intake recovery replays only unconsumed commands (`internal/recover/recover.go:71-84`). Once the canonical decision exists, duplicate intake and restart short-circuit to `outcomeFromRecord`; neither is presently a durable post-commit work driver.

The correction must define, before returning:

- one durable work identity keyed by decision/intake plus hook contract version;
- an atomically committed or deterministically derivable unresolved state;
- an exact ordered hook cursor with `not_started`, `running_or_unknown`, and terminal states where needed;
- the retry owner, pre-Ready recovery drain, retry ceiling, park rule, and terminal-resolution record;
- table reconstruction and duplicate/restart behavior;
- error, panic, and hard-crash cuts at every post-commit boundary, including the superseded and quarantine routes.

An in-memory slot may schedule work, but it cannot be the source of truth.

### F86 -- BLOCKER before H-16 IMPL: additive ignored-field compatibility fails open, and hook replay is not uniformly idempotent

The target asks whether a caller migration row is required. It is required.

If legacy `state` becomes the immutable decision and a new `post_commit_state` is merely additive, a legacy decoder ignores the new field and treats `state:"accepted"` as complete even when derived work is `pending`, `failed`, or `unknown`. That is the opposite of the dispatch's fail-closed requirement. The live host contains state-only success consumers:

- delivery nudge: `cmd/frank/main.go:337-352`;
- gate/approval prompters: `internal/engine/prompter.go:81-99`;
- resummon emission: `internal/engine/resummon.go:228-249`;
- MCP/native/shared-client JSON projections and external clients that consume the public response.

The correction must carry a consumer migration/compatibility table covering pre-commit failure, committed+complete, pending, failed, and unknown. It must choose either a versioned/negotiated projection where old clients cannot observe completed acceptance for non-complete work, or one atomic migration of every success consumer. "Old JSON still decodes" is not a fail-closed proof. An omitted field also cannot distinguish a legacy emitter from a new emitter unless the wire has an exact version/capability rule.

Replay needs a per-hook idempotency and result-durability table. `AfterAccepted` performs seat mint delivery; `MintOrReplace` creates a new credential on each invocation, while credential/endpoint extras are absent from `outcomeFromRecord`. Blind retry can rotate a credential no caller receives. Conversely, if an earlier hook fails before `AfterAccepted`, later generic retry cannot silently skip mint and report complete or run mint on an unrelated command. The m-7 pair independently blocked this as R1-F2/R1-F3.

Because versioning, caller-present delivery, durable recovery, and secret-result handling are hard-to-reverse API semantics immediately consumed by T4, the corrected H-16 design needs a focused decision record, fresh exact-byte pair review, and master/VP review before any IMPL branch opens.

### F87 -- BLOCKER before stage-4/5 final-byte review and stage-6 lock: the H-17 supplements are additive, but not assemblable or effect-complete

The authority posture is accepted: H-17 rows state current open designs and residuals; they do not authorize movement of a closed byte. Where a frozen contract lacks an answer, the row must say `unknown`, `not specified`, or `residual`, and master owns the later reconciliation.

The current supplements are incomplete against canonical H-17. Both omit:

- `effect_class`;
- `policy owner/artifact`;
- `decision point`.

Master would have to invent those cells during stage 6. A machine-readable table also needs a stable `effect_id` merge key. The common minimum schema must be issued once, byte-exact, and include the full backlog schema; it should distinguish requester/executor, authority source, outcome reporter/observer/validator, canonical recorder, and threat/claim scope where those roles differ.

The current m-9/m-10 folds expose concrete coverage defects:

- m-9 labels m-10 `attempt_open_ok` as provider authorization, but that is admission/row-first ordering; m-8 owns policy authorization immediately before credential attach/send.
- m-9 groups `relay.submit/project/read` under a store-append effect even though project/read do not append, omits `Describe` and push from the closed client surface, and does not explicitly route the compaction provider call through a fresh ordinary m-8 attempt.
- m-9 classifies E0 non-emission as a bypass; it is a failure/unknown condition. A bypass is an alternate effect path.
- m-10's seven rows omit authoritative transition families including run start/stop/recovery, process spawn/retire, turn admission, epoch publication, cancellation/control sends, provider-attempt transitions, and app-event carriage. A generic STORE meta-row cannot replace per-effect authority/failure rows.
- m-10 records or validates several worker/connector reports; that does not make it an independent observer of the underlying effect. Reporter, observer, validator, and canonical record must be named separately.

The parenthetical effect lists in the supplements are minimum, non-exhaustive coverage families. Each design must map every authoritative transition/effect family to a row or an explicit non-effect rationale. Current drafting and the already-open grills may continue, but neither design may declare `GRILL_LOCK`, request final-byte pair review, or issue its final SITREP until the corrected schema and coverage return are folded.

### F88 -- BLOCKER before stage-6 lock/T4 PLAN: the adopted T4 exit sentence overclaims the ratified MVP

`FRANK-HARDENING-BACKLOG.md:54` says "ONE real effect" and then names both one provider send and one local write. More importantly, it applies "exclusive chokepoint" and "no alternate path" across both even though the ratified MVP says:

- the designated provider attempt is governed;
- local tool effects, including bash-originated network/file effects, are not governed;
- F59 proves the exact structured call was authorized and invoked, not what ambient bash did to the host;
- same-UID credential/file inspection remains an accepted residual.

The exit proof must be split:

1. **Provider-send leg:** prove the designated credentialed route's freeze -> policy authorize -> credential attach -> send chain, scoped as "no alternate designed credentialed provider path within the declared confusion-not-malice threat model." Do not claim system-wide OS isolation.
2. **Local-write leg:** prove exact F59 request binding, one-shot consume, invocation, and crash-safe honest outcome. Do not claim an exclusive filesystem chokepoint, host containment, or no ambient bash path.
3. **Composed governed-turn leg:** bind both records to the same run/turn/release evidence without upgrading either leg's claim.

"Directly-attributable authority" also conflicts with deferring H-18. The current direct route authenticates live operator ingress, but its later agent-authored citation is explicitly E0 self-report, not operator provenance (`STEP-3-ARCH-AMENDMENT.md:114-123`). Acceptable correction:

- keep full H-18 in Step 4 and narrow T4 to an attributable requester/actor, named policy source, and exact request binding; or
- retain literal A-class directly-attributable authority and move a minimal H-18a authenticated exact-decision record pre-exit.

Signed evidence remains conditional on H-20. Until H-20 lands, the exit record says `not applicable/deferred`; it does not count signed evidence as a passed property.

## Answers and dispositions

1. **H-16 scope:** REVISE. A caller migration row is mandatory, the path census must be corrected, and durability/replay/idempotency must be contract-exact before IMPL. The m-7 pair's current MUST-REVISE is sound.
2. **H-16 timing:** CONCUR PRE-T4. New T4 consumers must not bind the defective/ambiguous API. This does not reopen Step 2 or INV-CATALOG.
3. **H-17 supplements:** CONCUR ADDITIVE-ONLY, but REVISE the schema and coverage now. No closed byte moves; current rows are inputs, not a stage-6 lock.
4. **Claim discipline:** CONCUR. Apply it to current outward/public claims while preserving append-only historical wording as history.
5. **H-18:** full mechanism remains Step 4 only if the T4 claim is narrowed; otherwise split H-18a pre-exit.
6. **H-19:** CONCUR Step-4 amendment-class. State explicitly that generation provenance is neither authority provenance nor attestation.
7. **H-20:** CONCUR Step-4+; signed evidence is conditional, not an MVP pass.
8. **H-21:** CONCUR Step 4 only under the narrowed local-write proof. Literal exclusive local-effect mediation would pull a minimum H-21/H-12 leg pre-exit.
9. **H-22:** CONCUR Step-4+; no single-effect MVP dependency.
10. **H-23:** not a T4 blocker, but assign owner/tier before lifecycle dashboards become authority-bearing projections.
11. **H-24:** permit a bounded, non-gating model of ticket x epoch x crash and authorize x send x restart; turn counterexamples into fixtures. Do not put complete formalization on the T4 critical path.
12. **H-25:** land stable raw IDs/counters during T4 where retrofit risk is real; keep topology analysis/routing use for Step 4.

## Required return and gates

1. Master issues an append-only H-16 narrowing supplement that incorporates F85/F86 and the m-7 R1-F1..F3 pair return. Historical relays remain byte-exact.
2. M-7 returns corrected design bytes, focused decision record, fresh uniquely-parented pair review, and then master/VP review. **No H-16 IMPL branch yet.**
3. Master issues one canonical H-17 row schema and corrected additive supplements. M-9/m-10 return a complete row/rationale inventory at current hashes before final-byte review.
4. Master corrects the H-18/T4 disposition and exit-test sentence. If literal A-class attribution or system-wide exclusivity is retained, route the required pre-exit mechanisms rather than overclaiming the current MVP.
5. Return the corrected packet for fresh VP review.

H-16 design correction may proceed. Stage-4/5 drafting and active grill discussion may proceed within the already-open authority. H-16 IMPL, both stage-4/5 final-byte reviews and lock claims, stage 6, PLAN, T4 code token, credentials, provider calls, release binding, E3, merge, and deploy remain held. Step 2 remains closed.

## Verification

- The exact target and all three planner dispatches reproduce the hashes above and end in exact-file `OK`.
- H-16 design r1 and the m-7 pair MUST-REVISE were read in full; the latter ends in exact-file `OK`.
- Current `frank/internal/engine/loop.go`, recovery, seat-mint, delivery-nudge, prompter, resummon, and MCP response paths were reread at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.
- Focused current tests passed: `go test -count=1 ./internal/engine ./internal/recover ./cmd/frank-mcp`.
- Current m-9/m-10 H-17 sections were reread at the exact hashes above.
- Harness cwd is not a git repository. `frank/` is clean on `main...origin/main` at `6e4d65791322`, with origin delta `+0/-0`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-045522.md` and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean on `main@6e4d65791322` at `+0/-0` versus `origin/main`.
RELAY_LINT: OK -- exact-file proof rerun after the INDEX append; root-wide historical/index noise is outside this artifact.
Next requested action: issue the bounded H-16 and H-17 corrections plus the honest T4/H-18 disposition, then return current bytes and hashes for fresh VP review before any held gate advances.
