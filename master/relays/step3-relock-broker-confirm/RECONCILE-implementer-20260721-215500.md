## RECONCILE — m-9 byte-bound F73 CONFIRMATION on broker-study rev8 `64f9136e…` + m-9 §D join-record co-sign; zero revise-requests, with the current-`parked_unknown` / later-D2-`uncertain` boundary explicit

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-relock-broker-confirm
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the exact-byte consumer scope is coherent; the only precision boundary is already assigned by the ratified stage-6 amendment to the affected m-9 final
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-broker-study
IN_REPLY_TO: master/relays/step3-relock-broker-confirm/DESIGN-planner-20260721-213000.md
FROM: m-9.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-9.planner, m-7.planner, m-7.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-broker-confirm/RECONCILE-implementer-20260721-215500.md
SUBJECT: independently confirmed the full m-9 §C scope on broker-study rev8 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; co-sign the m-9 consumer half of §D as a forward contract for the affected final, without claiming the D2 manifest consumer already exists in worker r7 / lifecycle r21

Master — I independently checked the planner assessment against the exact study, frozen r11, approved m-9 worker/lifecycle bytes, and ratified stage-6 amendment. **F73 CONFIRM: the rev8 m-9 §C scope is consumable as written; no revise-request.** I also co-sign the m-9 consumer half of the §D join record below, with one necessary temporal boundary made explicit: the existing `parked_unknown` two-gate consumer is already in r21/r7, while the separate D2 settlement-manifest `uncertain` consumer is a ratified obligation for the affected m-9 final and is not present in those frozen stage-3/stage-4 bytes.

## Independent m-9 §C confirmation

1. **Typed cut results — confirmed.** Rev8 §Q3.1 maps a live-connection cut `relay.submit` to `broker:unknown-outcome` and cut `project`/`read`/`Describe` to `broker:stale-epoch`. Lifecycle r21 §1.6 already disposes `unknown-outcome` as rediscover-never-resend and `stale-epoch` as fenced-generation fail-closed/no-local-retry. The study's “re-invoke under the successor” is not permission for the fenced generation to retry; it is a fresh-capability action by the successor. No ambiguity remains when generation scope is preserved.
2. **Conditional delivery — confirmed.** Rev8 correctly makes the typed response conditional on the admitting connection surviving. Loss of that response does not erase the mandatory relay-call disposition: m-10's canonical outcome row survives, and m-9 r21 §1.4 already treats conductor `project`/`read` truth as the rediscovery authority. Caller response and `boundary_cut` telemetry are not elevated into a second outcome carrier.
3. **Informed disclosure — confirmed with the boundary below.** The currently approved m-9 mechanism is r21 §2.6 D-4: `turn_open.parked_unknown` is surfaced before ASSEMBLING, then `attempt_open_ok.parked_unknown` is compared totally before DATA-P, with added/changed state forcing reassembly. That is sufficient for the already-designed state-only disclosure. Stage-6 amendment rev12 §D2 separately requires the continuation `turn_open` settlement manifest and its producer-total `uncertain` class; consuming/reconciling that manifest remains affected-final scope.
4. **Fresh F59 ticket, never automatic — confirmed.** Worker r7 §5 puts all eight model tools through the single authorize→consume→invoke→record settlement point. Lifecycle r21 §2.6/§3.4 narrows one-shot safety to identifier-exact `UNIQUE(run_id, turn_id, tool_call_id)`: a re-invocation is an informed new call with a new `tool_call_id` and therefore a fresh ticket, never replay of the cut identified call.
5. **Uniformity and `Describe` — confirmed.** Worker r7 §5.1 has no local/relay/read-only carve-out for model-issued tools. `Describe` is different because it is worker consumer metadata, not a model-issued tool: the r7 census classifies it as ticket-free/idempotent metadata service, consistent with rev8 §Q3.4's effect-free free-retry posture. “Free retry” remains successor/current-capability scoped; it does not weaken r21's fenced-generation rule.
6. **Attach taxonomy — confirmed unchanged.** Lifecycle r21 §1.2 already makes `broker:attach-suspended` total over the §2.4 suspension causes plus PREPARING, with bounded transient retry; `attach-tuple-mismatch` remains terminal for a fenced generation. Rev8 preserves FX-TB-19 and removes only the crossing machinery beyond the PREPARING barrier.
7. **Supersession sweep — confirmed.** Rev8 supersedes r11's `CROSSERS_DURABLE`/`ABORTED` machine, crossing identities/rows, crossing event classes, and recovery matrix. None is a worker-consumed surface. The worker-consumed attach taxonomy, forward-time fence, counter encoding, restart capability death/fresh acquisition, and read/quarantine posture survive explicitly. The `broker:record-unavailable` disposition remains a valid on-receipt worker row; rev8 narrows the producer circumstances without invalidating the consumer's total error-class table.

**Hard constraints:** F67 holds because m-9 still possesses only the epoch-bound USE capability and no S-A/S-B bytes. F64 holds because every relay call/`Describe` remains per-operation fenced and push stays forward-time fenced. F60/F66 hold because capabilities stay connection-scoped, die on broker/connection replacement, and are freshly acquired without credential propagation. Item-D honesty holds because a cut outcome is parked, never fabricated or auto-resubmitted.

## Precision boundary — accepted, not a revise-request

The phrases “existing `parked_unknown` disclosure” and “existing D2 `uncertain` class” refer to two different design strata:

- **Already realized in approved m-9 bytes:** r21 §2.6's D-4 state-only `parked_unknown` fields and two-gate pre-work consumer, carried by r7 §11.
- **Ratified but still owed in the affected final:** stage-6 amendment rev12 §D2/D3's three-class settlement manifest on continuation `turn_open`, including `uncertain`, plus m-9's log inspection/reconciliation and post-commit disposition-receipt no-work gate.

Therefore this confirmation does **not** claim that worker r7 `cb7ff970…` or lifecycle r21 `4d3bd14e…` already parses the D2 manifest. It confirms that rev8 introduces no new manifest class or second settlement path and that the mandated later m-9 final can consume the broker cut through the already-ratified D2 `uncertain` branch. Worker r7 and lifecycle r21 remain byte-unchanged on this confirmation path; master opens their affected-final amendment later in the §11 sequence.

## m-9 signed half of the §D join record

**Against broker-study rev8 SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`, m-9 co-signs the consumer reading:**

- When a relay tool call is cut after its F59 ticket was consumed but before the worker recorded an observed outcome, m-9 does not synthesize completion and does not re-drive that identified call. m-10's canonical rows park `UNKNOWN_TOOL_OUTCOME`; the conductor record remains effect truth.
- The successor is informed before work through the existing identity-exact `parked_unknown` disclosure and, in the affected continuation final mandated by stage-6 D2/D3, through the same identity's settlement-manifest `uncertain` entry on `turn_open`. These are coordinated disclosures over one m-10 outcome carrier, not two outcome authorities.
- m-9 reconciles the disclosed relay identity against conductor truth using its own seat capability and the durable `project`/`read` path. The affected D2 final also performs the amendment-required manifest/log reconciliation and waits for m-10's post-commit disposition receipt before any provider attempt, tool effect, or conductor verb.
- Any later invocation is an informed **fresh** model tool call with a fresh `tool_call_id` and fresh F59 ticket. Neither the broker nor the worker automatically resends the cut submit; identifier-exact one-shot remains true without claiming semantic-effect dedup across a new identity.
- The cut adds no broker-owned manifest class, receipt, or m-9 content-log outcome record. The D2 manifest class is m-10-produced under the ratified stage-6 contract; m-9 consumes it in the affected final.

Signed by the channel-stamped `FROM: m-9.implementer` on the exact rev8 bytes. This is the m-9 half only; m-10 signs its producer half under its own channel stamp.

## Verification

- Incoming m-9 planner relay SHA-256: `5a3c39a43357ce49f225896aa1d0c88586e549e0d2d0e8c084738e442756b15b`; exact-file lint: `OK`.
- Broker study rev8: `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; pair approval is byte-bound in `step3-relock-m7-broker/DESIGN-REVIEW-implementer-20260721-205236.md`.
- Frozen m-7 r11: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; approved m-9 worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; approved lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; ratified stage-6 amendment rev12: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- No governing design byte or `frank/` source byte was modified.

ACTIONS_GIT_REF: docs-workspace disk action only — created this m-9 F73 confirmation / join-record RECONCILE relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-broker-confirm/RECONCILE-implementer-20260721-215500.md`.
Next requested action: master pairs this m-9 confirmation/co-sign with m-10's independently pair-reviewed producer half, carries the NO-H-24 determination into §11, and opens the affected m-9/m-10 final-amendment lanes; all lock/PLAN/T4/code gates remain held.
