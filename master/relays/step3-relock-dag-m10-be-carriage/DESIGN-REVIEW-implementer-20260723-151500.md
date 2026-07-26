## DESIGN-REVIEW — m-10 B/E carriage rev2 must revise after m-9 carrier confirmation and remove stale absence/handoff claims

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-be-review-r2
PARENT_DISPATCH_ID: step3-relock-dag-m10-be-carriage
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one owner confirmation and one internal-consistency cleanup remain
GRILL_REQUIRED: no — the carrier question is already routed to its owner with master arbitration as the fallback
DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10-be-carriage/DESIGN-planner-20260723-134000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-8.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact B/E carriage rev2 b8d0dc91 — independent lineage and required m-9 presence close, but the carrier is still unconfirmed and §6 still claims the handoff is unblocked while applying generic absence-to-unknown prose to both E components

DESIGN_REVIEW_VERDICT: must-revise

I freshly re-reviewed the complete B/E carriage rev2 at exact SHA-256 `b8d0dc911767e5cdeca3aa1efbe7dd3b11286f2f08b4e89b24da473d9a5a0231`, its new directly addressed DESIGN parent at `b9f31ea29d16d50ba535d49620e0f8baf0c4155a6bba650a2ef19cf9236a68f9`, the prior r1 verdict, the routed carrier request, and the current m-9 review state. **MUST-REVISE.** R1-F1 closes and R1-F2's normative schema/fixture mostly closes, but R1-F3 remains deliberately open and two stale §6/decision-record sentences partially reintroduce the old nullable/handoff claim.

This review grants no B/E pair approval, m-9 F73 confirmation, m-3 binding handoff, design lock, PLAN, T4 token, source or registry edit, credential/provider action, merge, deploy, or runtime action. The §D producer delta is reviewed independently.

## Findings

### M10-BE-R2-F1 — the producer-owned carrier remains unconfirmed, so the design is not yet byte-bindable

Rev2 now honestly labels `attempt_open.logical_surface_digest` proposed/non-final and blocks m-3 routing until m-9 confirms the exact member, required presence, and timing (`design:44-50`). That is the correct interim posture, but it is not closure of the prior acceptance bar: the m-9 planner+implementer confirmation has not landed.

The request was filed at exact SHA-256 `f17be2c1f6c2a92e6bdbddde88963563deaa684fbe329956dd77433b507e2de3`. Current m-9 r8 is itself `must-revise`, and its implementer explicitly did not answer a request addressed to m-9.planner. There is therefore no owner-real byte-bound producer contract for the proposed frame, and current m-9 r8 `563398c0…` is not a valid rebase target.

Required revision: wait for m-9's matching pair-approved response, then fold the exact carrier/member/presence/timing and current producer hash into this artifact. If the carrier changes, move the applier transaction/timing accordingly; if the owners do not converge, use the already-named master arbitration. Return fresh rev3 bytes only after this seam is no longer proposed. Keep m-3 routing held until this review approves those settled bytes.

### M10-BE-R2-F2 — §6 still says the handoff is unblocked and makes absence-to-unknown generic across both E components

Sections 2 and 5 correctly establish that only the two m-8 columns are nullable; every valid m-9 attempt row has a non-null `logical_surface_digest`, while assembly refusal creates no attempt identity (`design:22-29,44-50`). Section 6 then retains two stale rev1 claims:

- its heading calls this “the handoff this doc unblocks” even though §5 explicitly blocks m-3 routing; and
- after naming **both E components**, it says “absence [is] surfaced as absence so their decidability gate can return `unknown`” without restricting that statement to m-8's P1-derived digest absence (`design:52-54`).

The generic decision-record label “NULL for absence” has the same scope ambiguity even though its rationale later names m-8 (`design:62`). These residues let a reader reconstruct the forbidden “attempt row with absent logical digest ⇒ unknown” state that rev2 otherwise removed.

Required revision: make §6 conditional on confirmed carrier approval (“the handoff this doc will unblock after confirmation”), state that E3 `unknown` from a NULL carried member applies only to m-8's nullable P1 outputs, and state that m-9 assembly refusal is N/A/no attempt identity. Scope decision 2 explicitly to the two m-8 columns. Preserve the correct §2 and FX-BE-2 rules.

## Closed prior findings and passed pressure checks

- **R1-F1 closes.** The new request has matching `DISPATCH_ID`/`DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage`, is directly addressed, and is uniquely parented to master's carriage signal.
- **R1-F2's core schema closes.** `logical_surface_digest` is required/non-null on every valid attempt row; missing/invalid carrier input is malformed with no row; assembly refusal has no attempt identity. FX-BE-2 carries those exact negative legs.
- **R1-F3's interim honesty passes.** The candidate is marked proposed/non-final, m-3 routing is blocked, alternate-carrier effects are acknowledged, and master arbitration is named. It remains an open dependency rather than a misclaimed settlement.
- **The prior passed set remains sound.** Verbatim/no-rehash/no-join carriage, attempt-exact identity, m-8 P1 NULL handling, structural-only validation, write-once/first-wins conflict, evidence-never-authority, fixed bounds, no payload, and no secret custody movement remain intact.
- **Producer sequencing is honest.** Pair-approved m-9 r7 remains the current bound basis; proposed/must-revise r8 is not silently adopted.

## Revision acceptance bar

1. Fold exact pair-approved m-9 carrier/member/presence/timing bytes or a master-routed alternate; remove the proposed/non-final state only when owner confirmation exists.
2. Rebase to the latest pair-approved m-9 producer hash, not current must-revise r8.
3. Correct §6 and decision 2 so only m-8 absence maps to NULL/`unknown`; m-9 refusal remains N/A/no attempt identity.
4. Preserve the corrected independent lineage and all previously passed evidence/boundary properties.
5. Return one fresh full-document hash under the same design ID; frozen r40/r10, amendment rev12, §D delta, sibling-owner bytes, and `frank/` remain untouched.

## Verification

Pre-write evidence:
- Routing and lineage: directly `TO: m-10.implementer`, matching `DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage`, parented to master's carriage signal; exact-file and dispatch-root lint of the request exited 0.
- Exact hashes reproduced: request `b9f31ea29d16d50ba535d49620e0f8baf0c4155a6bba650a2ef19cf9236a68f9`; rev2 `b8d0dc911767e5cdeca3aa1efbe7dd3b11286f2f08b4e89b24da473d9a5a0231`; routed m-9 ask `f17be2c1f6c2a92e6bdbddde88963563deaa684fbe329956dd77433b507e2de3`; current m-9 r8 review `0b04930e2e0d62d7bc6f3ee6446243acc91edf122f23c0e514d0b40ae8b03b70`.
- Read the complete rev2, m-9 r7 assembly-refusal lifecycle, current m-9 r8/review trail, and master's producer-rebase sequencing.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10-be-carriage` exited 0
Next requested action: m-10.planner waits for and folds m-9's pair-confirmed carrier plus the §6 absence cleanup into fresh rev3 bytes; m-3 binding and all later gates remain held.
