## DESIGN-REVIEW - s8 executable-claim declaration amendment must revise: Rail-A classification contradicts acceptance behavior; many-row aggregation is undefined

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m3-review-r1
PARENT_DISPATCH_ID: s8-claim-input-m3
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded design-amendment review; master reconciliation is required because F1 also corrects the parent amendment's prescribed Rail-A posture
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-claim-input-m3/DESIGN-planner-20260712-011500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, s8.planner, s8.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise - distinguish optional absence from compatibility handling and define deterministic many-claim aggregation before m-2 finalizes v7 bytes

DESIGN_REVIEW_VERDICT: must-revise

The declaration surface closes the real reader-with-no-writer gap, and the input/output ownership split, pre-spawn revalidation, R2 posture, I-PH constraint, transition tripwire, and no-c2-reopen fence are directionally correct. Two acceptance-bearing semantics remain unresolved.

## F1 - BLOCKER: `executable_claims` is not Rail-A open when an ignored present row can change the terminal

Section 12(e) says an old reader may ignore a present `executable_claims` row without changing acceptance meaning, then equates that case with the valid absence/no-vantage behavior. Those are different cases.

An aware reader runs the selected check. A failed check produces an observed-false predicate and rejects before delivery. A reader that ignores the present declaration behaves as `Evaluate:nil`; for a non-authority record the no-vantage disposition can accept. Ignoring the new field therefore can turn a would-be rejection into acceptance. That is exactly the Rail-A closed/fail-closed criterion from the standing confusion-firewall rule.

Evidence:
- Design §12(c) states the gate acts on conductor-observed results selected by the declaration (`...probe-design.md:225`).
- Design §12(e) nevertheless says ignoring the declaration changes no acceptance meaning (`...:229`).
- The implementation calls the evaluator when present and maps `Predicate: fail` to `rejected` (`internal/observe/gate.go:100-135` in `s8-observe-spine@3cce8cd`).
- `Evaluate:nil` yields the no-vantage path, whose non-authority branch accepts (`gate.go:99-102,147-156`).
- The preserved production test requires the selected false-done path to reject (`test/fixtures/s8_exit_gate_test.go:104-190`).

Required fold:
1. Separate **optional absence in a v7-aware reader** from **a present declaration encountered by an incompatible/ignoring reader**.
2. Keep absence optional if that is the intended no-vantage semantic, but classify compatibility handling for a present declaration as **closed/fail-closed**. Name the concrete compatibility enforcement, likely the governed v6-to-v7 digest/capability transition plus refusal by readers that cannot interpret the row.
3. Add a fixture/acceptance statement proving a present declaration cannot be silently treated as absent by an incompatible reader.
4. Route the Rail-A correction through master because the parent amendment at `s8-claim-input-amendment/...-004511:21` itself prescribed additive/open.

## F2 - BLOCKER: many-per-record introduces terminal aggregation semantics that §12 does not define

Section 12(a) allows several declarations and says `achieved_evidence` is the maximum passing rung, but it does not define the rest of the aggregate contract. The runtime must still choose one terminal predicate/disposition from multiple outcomes. Without a lock, implementations may differ on first-fail versus run-all, row ordering, duplicate `claim_ref`, mixed pass/fail/skipped/unsafe, and machinery-fault precedence.

This is observable behavior, not an implementation detail:
- `PredicateResult` carries a verdict slice, and stamp computation already iterates it (`internal/observe/gate.go:28-34,187-215`).
- The current registry evaluator emits exactly one verdict (`internal/observe/registry.go:128-145`), so the new many-row composition must deliberately define aggregation.
- `evidence_integrity` is keyed by `claim_ref` (`gate.go:200-205`); duplicate refs would overwrite each other and make typed bounce/result identity ambiguous.

Required fold:
1. Pin `claim_ref` as non-empty, bounded, symbolic, and unique within the record; duplicate refs receive a named typed fill-time reject and authoritative observe-time refusal.
2. Pin deterministic declaration/result ordering.
3. Define whether all checks run or evaluation short-circuits, and define overall predicate, machinery-fault, and terminal precedence for mixed outcomes.
4. Preserve one conductor-computed result row per declaration and the already-stated maximum-passing-rung rule.
5. Add acceptance fixtures for at least duplicate refs and a mixed multi-claim result containing one pass plus one fail; include unsafe/machinery-fault precedence in the locked matrix or an existing cited matrix.

## Boundary And Scope Check

Writes: a seat-declared `executable_claims` input row and conductor-computed per-declaration result rows.
Reads: v7 FieldSpec, check registry schemas, lane/schema/suite registries, and the locked executor boundary.
Target entity: the candidate relay evaluated inside serialized submit before delivery.
Downstream consumer: observe gate/disposition, evidence stamps, and recipient projection.
Contract: cannot be approved until F1 compatibility behavior and F2 multi-row aggregation are explicit.
Proof target: focused v7 transition/compatibility fixtures plus duplicate/mixed-outcome fixtures, then the existing T9 production RED turns green.
No-consumer action: reject the amendment rather than land a declaration writer without deterministic evaluation.

No finding reopens c2, the byte-exact terminal enum, the interim defaults/s10 sunset, the catalog/genesis pins, or s9 adjudication scope. m-2 should not finalize v7 bytes and T9 should remain held until the revised m-3 semantics are re-reviewed and master reconciles the Rail-A correction.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no frank source/test edit, branch, commit, push, PR, or merge
FINAL_GIT_STATUS_SHORT: `git -C frank status --short` returned empty at `main@691d034`; active `s8-observe-spine@3cce8cd` contains existing build-lane changes not made by m-3.implementer
Next requested action: m-3.planner folds F1/F2 into §12 and reissues for DESIGN-REVIEW; master corrects the parent Rail-A instruction before m-2 finalizes v7
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-claim-input-m3/DESIGN-REVIEW-implementer-20260712-012200.md`
