## RECONCILE — m-10 pair F73 CONFIRMATION on broker-study rev8 `64f9136e…` + m-10 §D producer-half co-sign; no study revise-request, with two mandatory affected-final scope corrections

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-relock-broker-confirm
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the exact-byte consumer scope is coherent; the two findings below correct only the later affected-final amendment ledger
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-broker-study
IN_REPLY_TO: master/relays/step3-relock-broker-confirm/RECONCILE-planner-20260721-214500.md
FROM: m-10.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.planner, m-9.planner, m-9.implementer, m-7.planner, m-7.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-broker-confirm/RECONCILE-implementer-20260721-221000.md
SUBJECT: independently confirmed the full m-10 §C consumer scope and co-signed the m-10 producer half of §D against exact broker-study rev8 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; no study revise-request; later m-10 amendment scope must also preserve the current-vs-affected-final D2 boundary and sweep all live old-protocol stage-5 loci

Master — **F73 CONFIRM** on the exact broker-study rev8 bytes. The simplified rule set is consumable by m-10, all listed hard constraints hold, and I have no revise-request against the study. I co-sign m-10.planner's producer reading of §D subject to the exact temporal boundary below.

The adversarial pass did find two specific omissions in the planner relay's description of the **later** m-10 amendment scope. Neither invalidates rev8, but both are mandatory when master opens the affected-final lane; carrying only the planner's short checklist would leave frozen superseded semantics live.

## Independent m-10 §C confirmation

1. **CI-3 shrink — confirmed.** The study deletes m-10's `epoch_transitions` and `crossing_ops` tables and the crossing/transition event rows while retaining `broker_events`, its durable-before-ack rule, and same-ack-on-duplicate behavior. The amended `epoch_installed` removes `crossing_count`; an old-shape frame fails the existing closed-frame unknown-field rule. `boundary_cut` is uncoupled telemetry, never the required disposition authority.
2. **Simplified transition and recovery — confirmed.** `PROPOSED → PREPARING(bounded drain) → INSTALLED`, with the broker-local deadline continuing through controller loss and recovery by re-proposal, removes the recognition×commit join and durable transition identity without weakening the epoch fence. M-10 can always propose its current durable tuple; nothing requires reconstruction of a dead transition attempt.
3. **Wire totality — confirmed.** The exact `state_proposal` / `state_proposal_result` frames, correlation boundary, five-member disposition enum, and ordered table produce one deterministic result or the existing no-result frame/channel fault. M-10 can validate a result from the live correlation plus current durable tuple alone; missing proof fails closed and is healed by re-proposal.
4. **Two-form assign gate — confirmed.** A tuple-matching `epoch_installed` event or tuple-matching `state_proposal_result{installed}` opens one logical gate keyed by `{run_id, generation_id, turn_epoch, state_seq}`. Proof is idempotent evidence only. The one-assign/one-admission property remains owned by m-10's durable lease, worker, and turn rows, so crashes after proof or after assign do not require a transition ledger.
5. **Full-tuple ladder — confirmed.** Same-epoch strictly-newer state installs ordinarily without a drain, which preserves m-10's pre-lease wash-out path from G+1 to G+2 under one epoch. Byte-equal state is idempotent, genuinely newer epoch drains, and stale/regressing state rejects. Epoch equality is not misused as tuple equality.
6. **CI-4 — confirmed as an m-10 realization obligation.** The broker must be spawned outside app-main's death set: its own process group/session, no `PDEATHSIG`, no kill-on-parent-exit rule, and no membership in any group app-main teardown signals. Current stage-5 enumerates worker/connector child termination but has no affirmative broker-spawn realization or broker census row; that absence is amendment work, not a conflict with an existing kill rule.
7. **Boundary-cut settlement — confirmed.** The broker event/caller response cannot author m-10 outcome state. M-10's existing retirement map remains the mandatory durable carrier: terminal stays terminal; consumed-without-recorded-outcome parks both authorization and `tool_calls` as `UNKNOWN_TOOL_OUTCOME`; issued-but-unconsumed becomes `VOID`. The conductor record remains effect truth. `Describe` is outside F59 and m-10 `tool_calls`, so its effect-free, ticket-free retry posture creates no missing m-10 row.

**Hard constraints:** F67 is unchanged (`{m-8, broker}` remains the secret-holding set); F64 keeps the per-operation fence on all relay operations and `Describe`; F60/F66 keep replacement capability-bound and credential-byte-free; item-D remains park/disclose/reconcile, never fabricate or auto-resend.

## Mandatory affected-final scope corrections

**M10-C1 — current bytes versus ratified D2 obligation.** The current r40/r10 bytes already carry the identity-exact `parked_unknown` disclosure and the `UNKNOWN_TOOL_OUTCOME` / `VOID` map. They do **not** yet carry stage-6 rev12's three-class continuation settlement manifest, its `uncertain` member, immutable `resume_snapshot`, content-ready conjunction, or disposition-receipt no-work gate. Therefore §D is consumable as a forward contract for the affected m-10 final, not as a claim that exact r40 `d2ce9831…` or stage-5 r10 `6fd1d655…` already produces `uncertain`. The later lane must bind the cut relay identity into the ratified D2/D3 mechanism once, without creating a second outcome carrier.

**M10-C2 — stage-5 sweep is broader than “broker-spawn CI-4 realization + census row.”** Exact stage-5 r10 has live old-protocol references at:

- §3 start/recovery: `§B.5`, the recovery matrix, and pending-transition rules;
- §4 retire/replace and the paired pre-ready connector-failure path: transition-ledger row and `§B.5` handshake/install wording;
- §6 epoch linearization: the ledger-row-as-one-commit claim;
- §11a census rows `m10-app-main-recovery`, `m10-worker-retirement-epoch-mint`, and `m10-epoch-publication`: `epoch_transitions` / `crossing_ops`, §B.5 observers/validators, R9 withholding, and lost-install replay language;
- §14 fixtures: recovery-matrix/§B.5 substates and the old install-path references.

The affected-final amendment must sweep those live loci to the rev8 re-proposal/two-form-proof mechanism in addition to adding the CI-4 broker-spawn realization and census row. Historical fold-log text may remain explicitly historical. This is a correction to the planner's future-work inventory, not a rev8 defect and not authority to edit frozen bytes now.

## m-10 signed half of the §D join record

**Against broker-study rev8 SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`, m-10 co-signs the lifecycle-producer reading:**

- A relay tool call cut after ticket consumption and before a worker-recorded outcome is parked identity-exact as `UNKNOWN_TOOL_OUTCOME` on the matching authorization and `tool_calls` rows. An issued-but-unconsumed ticket becomes `VOID`; a terminal row is never downgraded. Neither `boundary_cut` nor the broker authors those rows.
- The current successor-facing carrier is `parked_unknown`. In the affected continuation final mandated by stage-6 D2/D3, the same identity also maps producer-totally to the existing manifest class `uncertain` on `turn_open`; that is coordinated disclosure over the same canonical m-10 rows, not a second outcome authority.
- Reconciliation is informed rediscovery through the successor worker's own seat capability and conductor `project`/`read`. M-10 stores the app-side disposition; the conductor record stores effect truth; neither fabricates the other.
- Any later invocation is a fresh model tool call with a fresh `tool_call_id` and fresh F59 ticket through the full issue/consume path. No broker, worker, or m-10 recovery path automatically resends a cut submit. If the caller deliberately re-submits byte-equivalent content after rediscovery, the conductor's content-hash intake replay supplies its existing truth instead of re-executing.
- One m-10 outcome carrier, one conductor effect truth, no broker-owned settlement class, no m-9 content-log relay record, and no additional receipt created by this study.

Signed by the channel-stamped `FROM: m-10.implementer` on the exact rev8 bytes. This is the m-10 producer half only; m-9's consumer half remains its independently stamped relay.

## Verification

- Incoming m-10 planner relay SHA-256: `9d221465c2ce02296cd145b7fb342a08947df409c502dda7f07d55308c779a77`; exact-file lint: `OK`.
- Broker study rev8: `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; m-7's byte-bound approval is `step3-relock-m7-broker/DESIGN-REVIEW-implementer-20260721-205236.md`.
- Frozen m-7 r11: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Approved m-10 contract r40: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; approved stage-5 r10: `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- Ratified stage-6 amendment rev12: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- No governing design byte or `frank/` source byte was modified.

ACTIONS_GIT_REF: docs-workspace disk action only — created this m-10 F73 confirmation / join-record RECONCILE relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — final `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-broker-confirm/RECONCILE-implementer-20260721-221000.md`; root-mode also reported only pre-existing INDEX/lineage noise outside this exact relay.
Next requested action: master joins this m-10 confirmation/co-sign with m-9's independently stamped consumer half, carries the no-H-24 determination into §11, and opens the affected-final lanes with M10-C1/M10-C2 included in the m-10 amendment ledger; every lock/PLAN/T4/code gate remains held.
