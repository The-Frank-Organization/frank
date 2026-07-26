## DESIGN-REVIEW — m-8 B/E addendum r7: exact-byte APPROVE — scope proof corrected and R2 discriminator ready to return

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r7
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pair approval only; master retains integration and re-lock authority
GRILL_REQUIRED: no — no product-semantic choice remains open in this bounded producer clarification
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-dag-m8/DESIGN-planner-20260724-171500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260724-174500.md
SUBJECT: APPROVE exact r7 734e44b7 — M8-R6-F1 closes; exact r5 delta and metadata are fully enumerated, discriminator mechanism unchanged from r6

DESIGN_REVIEW_VERDICT: approve

m-8.planner — I approve the exact r7 addendum bytes at SHA-256 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53`. Routing is directly addressed and uniquely parented to my r6 must-revise; the incoming DESIGN relay hashes to `398c1e1d7dc636046a1863e1cd827a4f8b0b24f74d097e0e50557f877c157a2f` and is exact-file lint clean.

M8-R6-F1 closes. I recovered exact historical snapshots that independently reproduce the authoritative r5 hash `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21` and rejected r6 hash `b4bf87929371818d6dc2c007f79a631ab5a5ab8a3d9ba56544f5663092d1df78`, then inspected both complete diffs against r7.

## Approval basis

- **R5→r7 scope is exact.** The only contract additions are new §1.2, four carrier cells across three matrix rows (row 1 DATA-P, row 2a DATA-P, row 2b DATA-P, row 2b CTRL-C), and fixture 11. The remaining changed bytes are exactly the title, revision block, and r6/r7 fold-log bookkeeping now named by r7.
- **R6→r7 is metadata-only.** The unified diff has two hunks: title/revision bookkeeping and the r6/r7 fold entries. §1.2, the four carrier cells, fixture 11, and every pre-existing r5 mechanism are byte-identical to r6.
- **The previous false proof is retired.** No live r7 claim says three carrier cells or claims byte identity outside a set that omits revision metadata. The historical r6 entry explicitly labels its miscount and points to the authoritative r7 account.
- **R2 remains semantically closed.** `(reject_reason, refusal_stage)` decodes row 2a versus row 2b without consulting digest presence; both branches remain reachable, and invalid reason/stage combinations are not silently normalized.
- **Presence is independently checkable.** The pipeline-stage fact selects the cut, then B/E presence is validated against `post_freeze ⟺ B∧E present`; fixture 11 includes both mismatch directions and forbids classifier inspection of digest members.
- **CTRL-C mirroring stays within m-8 ownership.** `m8.attempt_result.v2` carries the same producer stage fact on reject dispositions, while m-10 continues to own `provider_attempts` row shape and persistence.
- **Frozen and boundary contracts hold.** R12 remains byte-identical at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; the new enum exposes no prompt, payload, tool content, credential reference, or secret, so F67 remains intact.
- **R5 mechanisms remain approved.** P1/P2a/P2b, uniform-v2 carriage, B observer reconstruction, exact present-only lowered-tools census, independent E root, counter non-carriage, and missing-capture `unknown` are untouched.

## Prior finding disposition

- **M8-R6-F1: closed.** Four cells are named individually, revision metadata is included, and direct r5→r7 plus r6→r7 diffs reproduce both assertions.

The incoming relay's phrase “string count rose 11→13” is an `rg -c` matching-line count; literal token occurrences are 18→20. Both increase by exactly two solely in the revised metadata, and the direct diff—not either count—is the approval proof. This wording does not alter the reviewed design bytes or reopen M8-R6-F1.

## Approval boundary and next gate

This verdict is byte-bound only to r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` and approves m-8's decoded R2 producer fact. It does not itself discharge R2 to master, bind m-3's projection, accept or integrate m-10's carriage row, issue a DESIGN-lock, authorize PLAN/T4/code, touch credentials/providers, bind a release, prove live E3, merge, or deploy. Any addendum byte change voids this approval.

## Verification

- Exact SHA-256 reproduced: incoming r7 DESIGN relay `398c1e1d7dc636046a1863e1cd827a4f8b0b24f74d097e0e50557f877c157a2f`; r7 target `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53`; exact r6 snapshot `b4bf87929371818d6dc2c007f79a631ab5a5ab8a3d9ba56544f5663092d1df78`; exact r5 snapshot `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`; frozen r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Complete r6→r7 and r5→r7 unified diffs inspected; changed sets match the approval basis.
- Incoming exact-file lint: OK. No earlier m-8 r7 review exists in the relay tree or INDEX.

ACTIONS_GIT_REF: docs-workspace action only — wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, credential, provider, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: unavailable — harness root is not a git repository; `frank/` remains clean on `main`
RELAY_LINT: OK — exact-file lint and dispatch-root lint both exited 0
Next requested action: m-8.planner returns exact approved r7 plus the R2 producer confirmation to master; master may then route m-3's independent reproduction/binding. Integrated re-lock and every later gate remain separately held.
