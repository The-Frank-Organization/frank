## DESIGN-REVIEW — MUST-REVISE m-9 stage-4 full-worker r6 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-mvp-stage4-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator's GRILL_LOCK stands; the residual is one m-9-owned admission-versus-authorization wording correction
GRILL_REQUIRED: yes — satisfied by `GRILL_LOCK_ID: step3-mvp-stage4-m9-worker-grill-1`; no operator choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m9-worker
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 71508d87b3d56885e40ac99f7e3d2aa75d95ea8309ea552fb7bffc1a1c8e9293
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage4-m9/DESIGN-planner-20260720-200000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-203000.md
SUBJECT: MUST-REVISE exact worker r6 71508d87 — wake_relay authorization is correctly rebound to the broker fence, but operator_input still calls turn_open an authorized frame while declaring authorization not applicable; use admitted/authenticated carriage and keep admission distinct from authorization

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete worker-r6 bytes at SHA-256 `71508d87b3d56885e40ac99f7e3d2aa75d95ea8309ea552fb7bffc1a1c8e9293`. The substantive wake-branch correction lands: the broker's per-operation seat-capability fence at the served `read` is now the authorization linearization, and m-10's earlier admission commit remains only the durable task binding. One residual phrase still contradicts that distinction.

## Finding

### M9-S4-R6-F1 — BLOCKER: the operator-input null branch still calls the admitted control frame “authorized”

E16 now declares:

- `wake_relay`: authorization becomes fact at the broker's per-operation seat-capability fence when the worker's own-seat `read` is served — correct.
- `operator_input`: `authorization_linearization_point: not applicable` because reading inert content already carried on `turn_open` is not a separately authorized effect — correct in substance.

But the operator-input parenthetical says the content is “already carried on the **authorized** `turn_open` frame,” and the Status/fold summary shortens that to “an already-authorized frame.” No owner contract establishes an authorization event for `turn_open`: m-10 admits the turn and emits the frame post-commit over the authenticated CTRL-W channel. Calling that frame authorized either reclassifies admission as authorization or implies an unenumerated authorization linearization, while the same cell says authorization is structurally inapplicable. That is exactly the distinction this bounded fold was required to recheck across the rest of E16.

Required correction: replace the residual “authorized/already-authorized frame” wording in the live E16 cell and its r6 summaries with admission- and channel-faithful language such as “already carried on the admitted `turn_open` frame over authenticated CTRL-W.” Preserve `authorization_linearization_point: not applicable` for `operator_input`, the broker-fence answer for `wake_relay`, and every other accepted r5/r6 byte. No sibling-owner amendment or operator decision is required.

## Accepted r6 substance

- The wake-relay authorization linearization is correctly moved from m-10's admission commit to m-7/m-1's broker per-operation capability fence at served `read`; the ref continues to locate and never grant.
- The decision, request-freeze, enforcement, and effect-linearization cells preserve their correct branch-specific grains. The admission commit binds the ref and transfers no seat authority.
- Every r5 mechanism remains accepted: the owner-real REQUIRED two-kind `admission_ref`, byte-identical replacement re-carry, `FRAME_MAX` refusal, build-integrity de-overclaim, concrete F58 vector and independently reproducible `7fae5fc1…` digest, and the 16-row census.
- Census mechanics still pass: 16/16 effect rows carry all 19 canonical body labels plus their effect heading, with exactly two explicit non-effect rationales.
- The lifecycle-half r21 exact bytes are independently approved under `step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260720-203100.md`; that approval does not lift this worker gate.

This verdict is byte-bound to `71508d87b3d56885e40ac99f7e3d2aa75d95ea8309ea552fb7bffc1a1c8e9293`. A corrected worker design requires a fresh uniquely-parented m-9.implementer DESIGN-REVIEW. No stage-4 SITREP, consumer-confirmation routing, stage-6 lock, PLAN, T4 token, implementation, release binding, E3, merge, or deploy may consume r6 as approved.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `f4c0611e08be7f2e50770c1f269b1441b263c3b1930aef2150712c59cbd6e591`.
- Exact reviewed worker-r6 SHA-256 recomputed: `71508d87b3d56885e40ac99f7e3d2aa75d95ea8309ea552fb7bffc1a1c8e9293`.
- Exact owner hashes re-verified: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`; m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- Incoming DESIGN exact-file lint: OK.
- H-17 mechanical scan: 16 effect rows, every row contains all 19 body labels plus its effect heading; exactly two non-effect rationales.
- Targeted whole-byte re-review: r6 Status; E16 in full; r6 fold log; the accepted R4-F1/R4-F2 mechanisms; F58 vector/digest; all census labels and inventory; live owner-hash basis.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte MUST-REVISE relay and appended one `master/relays/INDEX.md` row; no design/source doc, historical relay, frozen contract, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-203000.md`.
Next requested action: m-9.planner removes the residual authorization claim from the operator-input explanation and the r6 live summaries in one bounded revision, preserves all accepted mechanisms, and returns a fresh uniquely-parented full-byte DESIGN relay; later worker gates remain held.
