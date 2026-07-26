## DESIGN-REVIEW — rev3 MUST REVISE: r2 recovery/deadline findings close, but identityless re-proposal is not joined to the `epoch_installed` assign gate

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r3
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-173535.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact re-review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev3 at SHA-256 `abaec50dae151684e27c0fcb714a80483164f495ac1cb25c8d455d7fc7382d22` — r2 lifecycle/cut corrections stand; the no-ledger recovery still lacks an exact install-proof join

## Verdict

**MUST REVISE.** R2-F1 and R2-F2 close: rev3 correctly separates broker-control adoption from generation retirement, consumes the keyed E+1 lifecycle, runs the broker-local deadline through control loss, installs locally at expiry while non-authorizing, and makes late-response discard total. R2-F3's stale caller/event proof sentences and mandatory-carrier scope are mostly corrected.

The remaining blocker is in the new simplified recovery mechanism itself: rev3 drops all durable transition identity while retaining an `epoch_transition_id`-bearing `epoch_installed` event as m-10's assign gate. It never defines how a replacement controller joins an old correlation ID to a new re-proposal, nor how duplicate old/new install evidence advances lifecycle exactly once.

## Closed From R2

- **R2-F1 CLOSED:** §§Q1.2/Q1.3 and FX-TB-20 now make the ordinary leased crash run one keyed retirement, one E+1 mint, state-sensitive parking, broker install, then successor assign. Adoption itself remains no-mint/no-rebind/no-re-auth.
- **R2-F2 CLOSED:** §§Q3.1/Q3.3.1, §R.2, and FX-TB-17′ now pin deadline expiry independently of control, irrevocable cut, post-expiry discard, local install, non-authorizing suspension, and no second drain for cut identities.
- **R2-F3 PARTIALLY CLOSED:** the two named stale caller/event proof sentences are corrected and `Describe` is separated from F59/item-D. R3-F3 identifies one residual universal claim.

## Findings

### R3-F1 — BLOCKER — No-ledger recovery has no exact join from an old install correlation to the replacement controller's assign gate

**Current bytes:** §Q3.3.1 says no transition identity survives an outage and `epoch_transition_id` is only an attempt-local event correlation token. The same paragraph promises `epoch_installed` re-delivery after control loss. §Q3.3.3 leaves the `epoch_installed{epoch_transition_id, generation_id, turn_epoch, state_seq, …}` schema unchanged. §C still names `epoch_installed` as m-10's assign gate, while CI-3 drops the transition ledger entirely.

**Failure trace:** m-10 durably publishes E+1 and proposes correlation T1; the broker enters PREPARING; m-10 dies; the broker reaches its local deadline, cuts, installs E+1, and queues `epoch_installed{T1,…}`. The replacement m-10 knows durable E+1 but, by design, has no durable T1. It re-proposes under T2. The broker equal-acks and may deliver the queued T1 event. Rev3 does not say whether m-10 may accept an event naming an unknown T1, whether the T2 ack is install proof, how T1 and any T2 install event deduplicate, or which exact durable tuple opens successor assign. Queue overflow or broker death during the outage makes the missing rule sharper: the old T1 event may never arrive even though the epoch value was installed.

**Required correction:** keep the ledger removed, but define one exact no-ledger install-proof protocol:

1. State the canonical lifecycle key used by m-10 to accept install proof independently of attempt correlation, at minimum the durable trusted tuple `{run_id, generation_id, turn_epoch, state_seq}`.
2. Define the equal-current-state re-proposal response and whether it itself is install proof or causes a fresh `epoch_installed` proof under the new correlation. Do not rely on an old event that may have been lost with the pending queue.
3. Treat T1/T2 as correlation-only: delayed T1 evidence and current T2 evidence for the same canonical tuple may create telemetry rows but advance the lifecycle/assign gate exactly once. A mismatched tuple must fail closed.
4. Cover surviving-broker queue loss, old-event-late arrival, controller crash before/after local install, and broker death after local install. Every leg must converge to one installed lifecycle fact and at most one successor assign without a durable transition row.
5. Add this exact rule to §Q3.3.1/3, §C's m-10 confirmation scope, §R, and FX-TB-17′. If the solution instead makes transition identity durable again, that reintroduces a ledger/identity owner and must be priced as a departure from the chosen simplification.

### R3-F2 — IMPORTANT — “equal epoch” is not sufficient idempotence for same-epoch generation/`state_seq` recovery

Rev3 recovery says an equal proposal against an already-installed epoch acks idempotently and only a genuinely newer epoch installs. Frozen m-10 also has no-mint branches where the epoch stays equal while the trusted broker snapshot changes: a pre-lease wash-out allocates G+2 under the same E+1 and increments `state_seq`; lease-state publication similarly changes the full tuple. Treating epoch equality as state equality can leave the broker on G+1 while m-10 later assigns G+2, producing the first-attach tuple mismatch Q2 claims impossible.

Pin the comparison over the full trusted `epoch_state`, not only `turn_epoch`:

- byte/current-state equality ⇒ idempotent installed proof;
- same epoch with a strictly newer valid `state_seq`/generation/lease tuple ⇒ ordinary trusted state install, no epoch drain;
- genuinely newer epoch ⇒ Q3 bounded-drain transition;
- stale/regressing state ⇒ reject fail-closed.

Add a pre-lease wash-out recovery fixture: broker has `{G+1,E+1,S}`, durable m-10 state is `{G+2,E+1,S+1}`, recovery installs the newer full tuple before assign, and G+2's first attach is `attach-ok`.

### R3-F3 — IMPORTANT — §Q3.2.4 still claims universal durable recording after `Describe` was scoped out

§Q3.4 now correctly says a cut `Describe` has no m-10 row and only conditional response plus uncoupled telemetry with the admitted dual-failure residual. But §Q3.2.4 still says F64 “complete-or-reject, recorded” is satisfied for every cut via §Q3.4, and the study header says F64 is unweakened. In frozen r11, “recorded” is reserved for durable commit and `Describe` participated in the mandatory crossing set.

Make the claim exact: the F64 **fence** remains unweakened for all four broker operations; the three effectful F59 relay calls retain mandatory durable settlement through m-10; `Describe`'s former mandatory crossing-record posture is intentionally superseded under the §8 simpler-rule authority and now carries only the named uncoupled telemetry residual. Remove any universal “recorded” claim that cites a nonexistent durable `Describe` carrier.

## Re-review Gate

Return one byte-exact rev4 that preserves every closed r1/r2 mechanism and adds only:

- the canonical no-ledger install-proof join and exact once-only assign gate;
- full-`epoch_state` comparison for same-epoch state advancement;
- crash/queue-loss/late-event fixtures for T1/T2 convergence;
- the pre-lease wash-out same-epoch/higher-state fixture; and
- honest F64 fence-versus-`Describe` recording wording.

The clean-cut/no-H-24 determination remains supportable. These corrections need no crossing rows or retained completion.

## Verification

- Reviewed rev3 at exact SHA-256 `abaec50dae151684e27c0fcb714a80483164f495ac1cb25c8d455d7fc7382d22`; incoming relay at exact SHA-256 `61c5e386629825fc8eb9658b9c8a0151a92a0703ff4f5999d5af1cf171cbe0ff`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`; root-mode output retains unrelated historical `INDEX.md`/lineage noise.
- No design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R3-F1/F2/F3 into one exact rev4 and relays it for byte-exact adversarial re-review; m-9/m-10 confirmations and re-lock remain held on rev3.
