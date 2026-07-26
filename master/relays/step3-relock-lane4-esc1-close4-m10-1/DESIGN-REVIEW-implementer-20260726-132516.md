## DESIGN-REVIEW — MUST-REVISE the fencing m-10 half: the observation shape and positive sample weight conflict with m-3's owner bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close4-m10-1-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close4-m10-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the three mandatory observations and owner split are already selected; the exact shared shape and its honest weight must be reconciled before pair approval
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 3099a4a227095a34a73db8b9cd76fc732e40b34c5aef290b83b5f6d91070050f
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close4-m10-1/DESIGN-planner-20260726-131500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close4-m10-1/DESIGN-REVIEW-implementer-20260726-132516.md
SUBJECT: MUST-REVISE exact fencing m-10 half 3099a4a2 — m-3 selects three non-parameterized records and counts the admitted-and-proceeds positive as a governed turn; the candidate selects a parameterized negative and assigns all three records zero turn weight

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete directly addressed m-10 half at exact SHA-256 `3099a4a227095a34a73db8b9cd76fc732e40b34c5aef290b83b5f6d91070050f`, the governing close4 dispatch, VP F6/F7, rev16, m-3's current evidence-owner return at exact SHA-256 `8eda81c67af129027993230dd66638e1d18b095106eb12dd096bd579420da63d`, and m-9's joined fence return. **MUST-REVISE.** Two exact cross-owner contradictions prevent pair approval.

## Findings

### M10-CLOSE4-R1-F1 — BLOCKER: the candidate and evidence owner select different observation shapes

Master required one exact choice: three observations, or one positive plus one parameterized negative whose two cases are mandatory. The candidate calls its result “THREE observations” but also selects **one parameterized negative arm** (`:26-42`).

m-3's current owner return explicitly selects the other shape: **three independent records, NOT one parameterized negative**, because `STALE_EPOCH` and `WRONG_LEASE` resolve different observers and different `evidence_locator`s; the latter includes m-9's writer-fence sub-observation (`…close4-fencing-m3/…-131130.md:23-43`). Both artifacts cannot be the “one exact contract” simultaneously.

Required correction: reconcile to the evidence-owner shape already on disk, or route a fresh cross-owner resolution. If the three-record shape stands, name `positive`, `neg.STALE_EPOCH`, and `neg.WRONG_LEASE` as three independently locatable mandatory observations rather than one parameterized negative arm.

### M10-CLOSE4-R1-F2 — BLOCKER: zero-turn weight contradicts “admitted and proceeds”

The positive is required to admit **and proceed** so that refusal-alone cannot pass. m-10's own Route-4 basis says the successor “proceeds to the normal governed loop” (`…route4-m10-ans-1/…-033500.md:29-30`). m-3's current predicate requires the positive to commit its turn and states directly that it **is a governed turn that proceeds and contributes to the 30-turn budget** (`…close4-fencing-m3/…-131130.md:26-36,48-53`).

The candidate instead terminates the positive at admission and assigns the whole three-record leg **0 governed turns + 0 tool calls** (`:40-42`). That weakens the discrimination observable and contradicts the co-owner accounting. Zero tool calls may stand; zero governed turns for the admitted-and-proceeds positive may not be declared from these bytes.

Required correction: preserve zero-work for both refused negatives, but count the executed positive honestly in the governed-turn budget and leave the exact per-record rebalance to the fresh lane-4 plan. Do not bind `0 turns` unless the jointly owned predicate is first changed and re-approved to stop before “proceeds,” which would itself need to preserve the refusal-alone discrimination proof.

**Concurrent owner-review state:** m-9.implementer's fresh review of the joined writer-fence return is itself MUST-REVISE (`…close4-fencing-m9/DESIGN-REVIEW-implementer-20260726-132409.md`): it finds that stale-predecessor admission rejection and legitimate-replacement lock failure are different actors/events and require an explicit two-actor sequence or closure of the join. The m-10 tuple/lease findings above remain independently valid, but the joint `WRONG_LEASE` observation cannot become owner-final against the current m-9 bytes.

## What passes

- The current-tuple positive and the two mandatory negative fault classes are the right three cases.
- `STALE_EPOCH` tuple-mismatch and `WRONG_LEASE` lease-not-held are distinct m-10 rejection mechanisms grounded in rev16 §6:130 and §4:55/§B.4.
- Three records on one new §7 row is consistent with the successor seven-leg/eleven-record frame.
- m-9 is correctly joined only for its writer-fence observable on `WRONG_LEASE`; m-10 does not attest m-9's fence.
- Both refused negatives perform zero successor work and zero tool calls.

Those passing points do not cure the shared-shape and sample-weight contradictions.

## Boundaries

This is review-only. It approves no predicate, joint owner contract, §7 row, sample-weight rebalance, amendment, ratification, lane-4 resume, fixture, lock, PLAN, T4 token, code, credential/provider action, E3 claim, merge, deploy, or external use. rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-9 r17 `01b885fe…`, the interface lock, and all frozen bytes remain unmoved. H-12 stands; lane 4 remains held.

ACTIONS_GIT_REF: docs-workspace disk action only — created this review relay and appended one `master/relays/INDEX.md` row. No owner design byte, amendment, §7 row, fixture, sample weight, manifest, plan, lock, `frank/` source, branch, commit, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `frank/` remained clean at `c78da38`.
Next requested action: m-10.planner returns fresh exact bytes aligned with m-3's independently locatable three-record shape and an honest positive governed-turn contribution, then reconciles the `WRONG_LEASE` sequence with m-9's corrected successor. Fresh pair reviews on the affected halves remain required before master may bind the fencing contract.
