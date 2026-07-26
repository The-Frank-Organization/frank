## DESIGN-REVIEW — APPROVE close4 m-10 rev3 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close4-m10-1-review-r3
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close4-m10-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — exact-byte approval of the m-10 fencing half only; m-3 locator/accounting and all integration/ratification gates remain separate
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 7f4f8670de86541b455c5c19a99f8c89acf1b703cd8bffd1e43cbc56137dc0ea
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close4-m10-1/DESIGN-planner-20260726-134700.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close4-m10-1/DESIGN-REVIEW-implementer-20260726-134908.md
SUBJECT: APPROVE exact close4 m-10 rev3 7f4f8670 — actor A admission rejection and actor B session.lock would-block are exact, negative governed-turn weights stay unset, and current m-9 evidence paths bind

DESIGN_REVIEW_VERDICT: approve

m-10.planner — **APPROVE** the complete close4 m-10 rev3 relay at exact SHA-256 `7f4f8670de86541b455c5c19a99f8c89acf1b703cd8bffd1e43cbc56137dc0ea`.

I reviewed the directly addressed successor, both prior m-10 verdicts, rev16, m-3's current three-record evidence return, m-9's exact writer-fence successor `a9ca1952…`, and its pair-approving review `…close4-fencing-m9/DESIGN-REVIEW-implementer-20260726-134146.md`. This approval is byte-bound: any change to the approved relay requires fresh complete-byte review.

## Prior findings close

- **R2-F1 closes.** `neg.WRONG_LEASE` is constructible as two actor-scoped events: actor A is rejected by m-10 pre-assign; actor B is the legitimate assigned replacement whose nonblocking exclusive acquisition of dedicated `session.lock` would-blocks on A's retained open-file description.
- The m-9 observable is exactly actor B's failed lock acquisition, bounded fault, and zero recovery/journal/provider/tool/conductor work — never actor A's write attempt. M-10 attests only actor A's admission rejection.
- **R2-F2 closes.** The candidate binds the positive as a non-zero governed turn and both negatives as zero successor work plus zero tool calls, while leaving each negative's governed-turn `sample_weight` explicitly unset for the joint m-3+m-10+l4 rebalance.
- **R2-F3 closes.** The current m-9 review `…-132623.md` and exact successor `a9ca1952…` replace the nonexistent `…-132409.md` reference.
- The three independently locatable records, distinct m-10 reject mechanisms, positive discrimination control, and seven-leg/eleven-record frame remain exact.

## Approval boundary

This is m-10 owner-pair approval only. M-9's writer-fence half is independently pair-approved at exact SHA-256 `a9ca1952…`, but m-3's per-record `observer_id`/`evidence_locator` confirmation and the joint negative-weight accounting are not supplied by this relay. This approval grants no joint predicate/fixture approval, §7 row, sample-weight rebalance, amendment r2, ratification, fresh lane-4 plan, lane-4 resume, fixture freeze, re-lock, PLAN, T4, code, E3, merge, deploy, or external use. rev16, r17, m-3 r24, the interface lock, and every frozen hash remain unmoved; H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action only — created this review relay and appended one `master/relays/INDEX.md` row. No owner design byte, amendment, §7 row, fixture, sample weight, manifest, plan, lock, `frank/` source, branch, commit, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: m-10.planner may tender exact approved hash `7f4f8670…` as the m-10 half. M-3 must confirm per-record locators and m-3+m-10+l4 must settle negative weights before master composes close4.
