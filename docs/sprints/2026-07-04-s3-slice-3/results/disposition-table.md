# S3 Disposition Table

Generated pair: `test/replay/dispositions.json` is the machine-readable source; this table is the human sprint artifact.

| anchor | check | class | disposition | surface | fixtures |
| --- | --- | --- | --- | --- | --- |
| relay-lint.py:phase-enum | PHASE token membership | 10b | dissolved-form | fieldspec enum validation | bad-phase-token; claude/B9-bad-enum.md; claude/C2-enum-bypass.md |
| relay-lint.py:address-token-shape | bare authority token placement | obsolete | obsolete | typed submit payload has no raw markdown token channel | bare-token-shape; addressing/T2-token-no-to.md |
| relay-lint.py:scope-diff-carrier | structured SCOPE_DIFF rows | 10b | dissolved-form | canonical row_array typed carrier | scope-diff-row-array |
| relay-lint.py:parent-substrate | parent relay/dispatch resolution | 10c | dissolved-lineage | lineage.Engine parent substrate over tables | addressing/G1-casefold-lineage/IMPL-Dispatch-20260610-102000.md |
| relay-lint.py:merge-grant | merge grant authorization | 10c | dissolved-lineage | lineage.Engine merge-claim prerequisite | merge-token/MT1-valid-token-grant/MERGE-GATE-20260610-130000.md |
