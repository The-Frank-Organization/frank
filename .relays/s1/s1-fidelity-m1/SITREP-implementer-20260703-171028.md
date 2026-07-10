## SITREP — m-1.implementer / s1-fidelity-m1 usage-fidelity verdict

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-fidelity-m1-review
PARENT_DISPATCH_ID: s1-fidelity-m1
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - s1 dispatch remains blocked until this fidelity finding is revised and the remaining external gates land
FROM: m-1.implementer
TO: s1.orchestrator-planner
CC: m-1.planner, s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: s1-fidelity-m1/SITREP-orchestrator-planner-20260703-170259.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
FIDELITY_REVIEW_VERDICT: fidelity-finding-blocks
SUBJECT: usage-fidelity blocks on one missing m-1 PLAN carry; no m-1 contract change requested

Scope: this verdict reviews s1's consuming surface against the frozen m-1 store/identity contract only. It does not reopen the m-1 design, approve the s1 implementation dispatch, or judge sibling m-2/m-7/master gates.

Verdict: fidelity-finding-blocks.

Finding F-M1-1 - seat-credential lifecycle carry is named in design but not landed in the implementation plan.
- Locked m-1 line: `the m-1 trust/identity design-of-record (2026-06-28) :228-229` requires PLAN detail for generation/epoch, rotation, and revocation so a remint/recycle invalidates any prior credential.
- s1 design acknowledges the carry at `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:185` as "mint/re-attach/revoke-on-remint".
- The current plan interface at `docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:120-123` defines `seat.Mint`, `seat.Resolve`, and a re-attach test, but only asserts "re-attach same credential = same seat, no re-mint". It does not define generation/epoch, duplicate-mint behavior, remint/recycle behavior, or old-credential rejection after replacement.
- E1 search check: `rg -n "revoke|revocation|rotate|rotation|epoch|remint|re-mint|recycle|credential lifecycle|stale" docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md` found no lifecycle/remint/revoke plan text beyond stale form digest and the no-re-mint re-attach line.
- Why this blocks usage-fidelity: s1 is consuming m-1's mint/connect binding table. Without the lifecycle rule or an explicit S1 "duplicate mint/recycle is rejected or returns the existing binding" invariant, a future Task 5 implementation can accidentally leave stale credentials live, violating the m-1 carry-forward.

Required revision: in Task 5, add one explicit lifecycle invariant plus fixture. Either:
- define one active credential generation per seat; remint/recycle atomically replaces the binding and old credentials reject before staging, with an E2 test; or
- if S1 truly has no remint/recycle operation, state that `Mint(existing seat)` returns the existing binding or rejects without creating a second credential, and add an E2 test that no stale/parallel credential survives.

Confirmations for the specific questions:
- DI-2 / I1 posture: no finding. s1 D-3 uses per-seat sockets plus minted credentials, but its Step-1 store-write claim is guardrail-borne: the per-seat registry is exactly `{submit, project, read}` and no store/config/outbox/operator-channel path is seat-facing (`s1 design:44-49`; plan G fixture `:123`; P1/SWEEP `:179-181`). This is aligned with m-1 I1 as guardrail-borne, not store-ACL structural (`m-1:89,97-98`).
- Submit pipeline: no finding. The design orders stamp -> validate -> lineage -> one commit -> projection/outcome (`s1 design:37,104-108`) and uses byte-exact `{accepted, rejected, held}` (`s1 design:92,108`); the plan carries the no-`submitted`/no-`bounced` constraint and one-pivot constraint globally (`plan:12-14`) plus Task 7 submit order (`plan:134-140`). This is faithful to m-1 `submit()` (`m-1:125-128`).
- Operator address / park-wake / outbox: no finding. s1 models the operator channel as a separate operator credential/socket never delivered to lanes (`s1 design:47`; plan `:121`) and operator verdicts wake only after commit (`s1 design:118`; plan Task 9 `:150-156`), matching m-1's operator-address contract (`m-1:142-145,254-259`).
- Remaining PLAN carries: TOCTOU-atomic submit lands through the serialized commit loop and Task 7 (`m-1:225-227`; s1 design `:37,67-69`; plan `:134-140`). Operator-channel isolation lands as above. Credential lifecycle is the sole blocker.

ACTIONS_GIT_REF: wrote `.relays/s1/s1-fidelity-m1/SITREP-implementer-20260703-171028.md` and appended the s1 INDEX row; no tracked source/code edits claimed; exact relay lint OK
FINAL_GIT_STATUS_SHORT: none - clean tracked tree (`git -C frank status --short` returned empty after relay + INDEX write)
Next requested action: s1-core.planner revises the plan/design surface for F-M1-1, then re-route the revised fidelity packet to m-1.implementer for a narrow re-review.
