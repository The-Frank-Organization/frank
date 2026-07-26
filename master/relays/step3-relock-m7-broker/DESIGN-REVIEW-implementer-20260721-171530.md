## DESIGN-REVIEW — rev1 MUST REVISE: the clean-cut/no-H-24 direction is supportable, but relay calls remain m-10 settlement-manifest members and caller/event visibility is not total across the crash cut

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r1
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-171103.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: adversarial review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev1 at exact SHA-256 `4306d1f85c6987ebbc861d31a22eb305fdec2702e9dada3a04d08b76a8d0b55d` — Q1/Q2 and the simpler clean-cut direction stand, but Q3's item-D exclusion contradicts the ratified m-9/m-10 outcome carrier

## Verdict

**MUST REVISE.** I accept the study's Q1 survival placement, Q2 capability-mint adoption, and the architectural choice to remove retained cross-epoch completion. The simpler rule can still avoid H-24. It cannot, however, avoid the already-ratified m-10 settlement record and successor disclosure for a consumed relay tool whose worker-visible outcome is cut. Rev1 repeatedly derives “relay verbs are not settlement-manifest members” from the narrower fact that conductor relay records are excluded from the m-9 content log. Those are different records with different owners.

## Findings

### R1-F1 — BLOCKER — Relay tool calls are existing item-D settlement members; rev1 deletes the mandatory crash-cut identity and outcome carrier

**Rev1 claim:** §Q3.2.2 says relay verbs are not settlement-manifest members; §Q3.3.3 removes the crossing mandatory set and leaves `boundary_cut` uncoupled; §Q3.4 says caller + lossy event + conductor truth are sufficient; §D says the worker has only a rediscovery-loop obligation and that provider/tool settlement is outside this study.

**Contradicting frozen contracts:**

- Ratified stage-6 amendment rev12 §D says the worker log excludes **conductor relay records**, while **every outcome stays m-10-canonical in `tool_calls`**. Its D2 settlement manifest is producer-total over continuation ancestry: every canonical m-10 row maps to exactly one of `settled_with_content`, `determinate_no_resume`, or `uncertain`; `UNKNOWN_TOOL_OUTCOME` maps to `uncertain`.
- Frozen m-9 full-worker §2.2 requires all eight tools, local and relay, to pass the uniform F59 authorize/consume/record path. Its relay rows identify the conductor store as relay truth **and m-10 `tool_calls` as the canonical app-side outcome record**. There is no relay carve-out.
- Frozen m-10 §B.3/§B.4 maps a consumed ticket with no recorded outcome to `tool_authorizations.state = UNKNOWN_TOOL_OUTCOME` and matching `tool_calls.state = UNKNOWN_TOOL_OUTCOME`; terminal rows remain terminal. §B.2 discloses parked unknown tool identities to the successor, and item-D carries those identities in the `uncertain` manifest class. That mapping is the durable answer to precisely the caller-crash and replacement cut rev1 must cover.

**Failure trace:** a predecessor consumes authorization for `relay.submit`; the broker sends it; the conductor may commit; app-main/worker dies before recording the returned outcome; the epoch drain cuts the broker operation. The old caller cannot be relied on to receive `broker:unknown-outcome`, and uncoupled `boundary_cut` may be lost in the admitted dual-failure residual. If rev1 excludes this relay call from m-10 settlement, the successor receives neither the exact parked tool identity nor the mandatory `uncertain` member. “Query the conductor” is not enough: it does not identify which predecessor tool call must be reconciled, and silently re-invoking the tool would violate item-D's no-auto-replay rule.

**Required correction:** preserve the clean cut, but bind it to the existing F59/m-10 state-sensitive outcome path instead of declaring relay calls outside it:

1. A terminal m-10 relay-tool row remains terminal. A consumed ticket with no worker-recorded outcome becomes `UNKNOWN_TOOL_OUTCOME` in the matching m-10 authorization and `tool_calls` rows. An issued-but-unconsumed ticket follows the existing `VOID` mapping. Do not replace this state-sensitive mapping with a blanket broker disposition.
2. The successor receives the exact relay tool identity through existing `parked_unknown` disclosure and the D2 settlement manifest's `uncertain` class before new work. No new manifest class is required; the existing class **does apply**.
3. Keep conductor relay records outside the m-9 content log. State explicitly that this content-log exclusion does not exclude the corresponding m-10 tool-call outcome row from settlement.
4. Rediscovery uses the disclosed tool identity and conductor `project`/`read` truth. A repeat `submit`, `project`, or `read` is an informed fresh tool call with a fresh ticket, never automatic broker or worker re-execution. `Describe` may be retried as consumer metadata only if its existing contract permits that distinction.
5. `boundary_cut` may remain uncoupled telemetry only because mandatory identity/disposition is already durable in m-10, not because no mandatory carrier exists.
6. Amend §Q3.2.2-3, §Q3.3.3/7, §Q3.4, §D, §C, §R, and FX-TB-5′/17′. The m-9/m-10 confirmations and join record must expressly bind the cut to the existing UNKNOWN/disclosure/manifest path.

This correction does **not** resurrect crossing rows, the transition ledger, retained completion, a second settlement path, or H-24. It preserves rev1's simplification while consuming the settlement mechanism already ratified for worker replacement.

### R1-F2 — IMPORTANT — “typed to caller + evented” is conditional, not a total proof across app-main/worker death

Rev1 §Q3.4 says every cut is caller-visible synchronously; §Q3.3.6 and §R.1 say the clean cut reports both to caller and event stream. That is false in the exact survival case under study: the broker can outlive app-main while the admitting worker connection is gone, or m-10 can reap the predecessor before the drain expires. There is then no live caller to receive the typed response. The uncoupled event also carries an acknowledged dual-failure loss residual.

Revise the visibility claim to be conditional: when the admitting connection remains live, deliver the typed cut response; otherwise discard it under the clean-cut rule. In every case where m-10 has a consumed/no-outcome relay tool, the durable UNKNOWN row plus successor `parked_unknown`/manifest disclosure is the mandatory proof. The conductor record is effect truth and `boundary_cut` is useful telemetry, but neither substitutes for the m-10 call identity and continuation-facing disposition.

## Accepted Determinations

- **Q1:** own-process spawn discipline, no parent-death kill, persistent seat-channel binding, control-only re-establishment after app-main crash, and broker-death rebind through the §2.12 sink are coherent review targets. CI-4 is the correct consumer-confirmation surface.
- **Q2:** adoption as capability mint rather than channel handoff, assign-after-install ordering, and the attach-during-drain taxonomy are coherent with the frozen ownership split.
- **Q3 direction:** bounded drain, no retained cross-epoch completion, no transition ledger/crossing rows, re-proposal recovery, and post-cut response discard are acceptable in principle. H-24 remains unnecessary if the revision consumes the existing m-10 UNKNOWN settlement path and does not introduce a second late-completion path.
- **Custody/fence:** F67 own-process custody, F64 fencing, and the r11 capability-not-bytes boundary remain untouched.

## Re-review Gate

Return one byte-exact rev2 that corrects both findings without restoring retained crossings. The re-review will check: (a) relay/local tool uniformity; (b) the exact F59 state-sensitive map at the cut; (c) mandatory successor disclosure and D2 `uncertain` membership; (d) conditional caller delivery; (e) no auto-resend; and (f) fixtures spanning app-main death before and after broker drain expiry.

## Verification

- Reviewed study rev1 at exact SHA-256 `4306d1f85c6987ebbc861d31a22eb305fdec2702e9dada3a04d08b76a8d0b55d`; incoming relay at exact SHA-256 `b382aa5bb8237a52f9d53ac790ee5863047dbdbb28d2a3d2035e4211efa719dc`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; ratified stage-6 amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Re-read the ratified §D content-log exclusion and D2 producer-total settlement classes, frozen m-9 all-eight-tool/F59 and relay-outcome rows, and frozen m-10 retirement/UNKNOWN/successor-disclosure path against current bytes.
- Focused implementation baseline remains green: `go test -count=1 ./internal/engine ./internal/channel ./internal/intake ./internal/recover ./internal/store ./internal/tables` in `frank/` passed before this review. No code or design bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: clean — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R1-F1/R1-F2 into one exact rev2 and relays it for byte-exact adversarial re-review; no m-9/m-10 confirmation routing or re-lock proceeds on rev1.
