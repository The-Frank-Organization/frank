# S3 Disposition Table

Generated pair: `test/replay/dispositions.json` is the machine-readable source; this table is the human sprint artifact.

| anchor | check | class | disposition | surface | fixtures |
| --- | --- | --- | --- | --- | --- |
| relay-lint.py:accepted-control | expected-OK oracle entries | retained | retained | typed equivalent accept path | claude/A1-valid-audit.md; addressing/G1-casefold-lineage; design-review/DR1-valid-design-doc-chain |
| relay-lint.py:phase-enum | PHASE token membership | 10b | dissolved-form | fieldspec enum validation | claude/B9-bad-enum.md; claude/C2-enum-bypass.md |
| relay-lint.py:address-token-shape | bare authority token placement | obsolete | obsolete | typed submit payload has no raw markdown token channel | addressing/T2-token-no-to.md; addressing/T3-token-to-planner.md; addressing/T4-token-two-implementers.md |
| relay-lint.py:scope-diff-carrier | structured SCOPE_DIFF rows | 10b | dissolved-form | canonical row_array typed carrier | content/E6-out-row-allin.md; content/E9-unparseable-row.md; content/E13-scopediff-detached-row.md |
| relay-lint.py:fold-scope-carrier | structured FOLD_SCOPE rows and required fold scope | 10b | dissolved-form | canonical row_array typed carrier plus REVIEW-FOLD required_when | fold/FD1-fold-edit-no-foldscope.md; fold/FD3-out-row-with-edit.md; fold/FD10-detached-row-foldscope.md |
| relay-lint.py:observe-actions-substance | FINAL_GIT_STATUS_SHORT/ACTIONS_GIT_REF observe-context substance | 10b-observe | dissolved-form | observe-layer predicate context and path-hygiene bounce surface | content/E1-empty-final-git-status.md; content/E2-empty-actions-git-ref.md; p9/P9b-claim-after-scan-blank-line.md |
| relay-lint.py:header-grammar | relay header grammar and required-field checks | 10b | dissolved-form | registry required_when, enum, and typed scalar validation | claude/B2-why-before-scan.md; lint-test/bad1.md; probes/N4-scan-result-mismatch.md |
| relay-lint.py:parent-substrate | parent relay/dispatch resolution and addressee lineage | 10c | dissolved-lineage | lineage.Engine parent substrate over tables | lineage/LI3-no-plan-review-parent; lineage/LI4-non-addressee-impl-report; lineage/LI5-edgeless-delegated-cc-trap |
| relay-lint.py:merge-grant | merge grant authorization | 10c | dissolved-lineage | lineage.Engine merge-claim prerequisite | merge/M1-merge-claim-no-auth; merge/M4-canonical-claim-no-auth; merge-token/MT9-cross-dispatch-runroot |
| relay-lint.py:design-review-lineage | design-review parent chain | 10c | dissolved-lineage | lineage.Engine design-review lock checks | design-review/DR2-edge-less-no-review; design-review/DR10-review-parent-not-design; design-review/DR19-F-thread-mustrevise |
| relay-lint.py:row-truth-check | ROW_TRUTH_CHECK scope/fold evidence | 10c | dissolved-lineage | lineage.Engine scope/fold flip drift checks | rowtruth/RT2-missing-fold-evidence.md; rowtruth/RT4-missing-scope-evidence.md; rowtruth/RT5-two-relay-out-in-flip |
| relay-lint.py:orch-review-visibility | orchestrator-reviewer visibility gate | 10c | dissolved-lineage | lineage.Engine orchestrator-review visibility over tables | orch-review/OR7-audit-no-reviewer; orch-review/OR10-override-impl-no-reviewer; orch-review/A3-self-waiver |
| relay-lint.py:role-from-stamp | payload ROLE/FROM proxying | obsolete | obsolete | seat.Stamp owns FROM/ROLE as the single channel | identity/S4a-proxy-from.md |
