## DESIGN-REVIEW — m-1 adversarial review r1 of the lane-4 edited-session boundary re-confirm

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m1-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the three corrections are owner-local claim narrowing against already-decided operator and master boundaries
GRILL_REQUIRED: no — no product choice remains open in the m-1 return
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-planner-20260726-130845.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-3.planner, m-9.implementer, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-REVIEW-implementer-20260726-132323.md
SUBJECT: must-revise — keep BR-INV prime but remove the impossible undetectable-yet-flagged claim, classify the three answers honestly, and restore the actual m-9 plus m-10 Step-4 store pair

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner — I reviewed the exact return at SHA-256 `38070c9d7a7935974a4751b31abf66f1dc414578cf7a3f1439974641d50dd35b`, the closure dispatch, the current m-9 and m-10 owner halves, the prior wider m-1 return, frozen m-1 component `d34a7c475c5988f12872a5aaf599b201fe56970b766b935bdb42b1d2851ab1ef`, frozen m-1 base `7c8b09a6…`, interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`, and stage-6 amendment `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

BR-INV′ is the correct operator-decided MVP boundary: only m-9's D1 journal is externally editable; the m-1 governed relay store and m-10 settlement store are excluded. The at-rest applicability, m-10 repairability carry, zero m-1 MVP supersession list, and lane holds also pass. Three statements overreach the decided boundary or the current owner bytes and block owner-final approval.

This review grants no owner-final fold, amendment-r2 acceptance, lane resume, fixture freeze, re-lock, PLAN, T4/code token, external use, credential, provider action, live E3, `frank/` edit, merge, or deploy.

## Findings

### M1-CLOSE3-R1-F1 — An undetectable edit cannot be visibly flagged, and m-10 has no payload-identity comparand

The return correctly admits that a payload edit which recomputes the advisory checksum has no specified mismatch path (`DESIGN-planner-20260726-130845.md:31-32`), then claims the same edit resumes visibly flagged/degraded, never silently clean, and never promoted. Those statements cannot all be true.

The current owner bytes make the missing edge exact. m-10's immutable entries carry call identity, `args_digest`, and terminal facts, not a content digest; a payload edit leaves that match key intact (`…close3-m10-1/DESIGN-planner-20260726-131200.md:26-35`). m-10 degrades exactly when m-9 reports `content_lost`, otherwise it resumes (`:37-48`). m-9's table therefore says a checksum-recomputed, present, outcome-consistent edit is unclassifiable and resumes as if clean at MVP (`…close3-editsm-m9/DESIGN-planner-20260726-131630.md:29-47`). The closure dispatch explicitly permits this narrowed honest answer (`…close3-editsm/DESIGN-orchestrator-planner-20260726-044200.md:29-33`).

The same error appears in answer (b): an immutable m-10 receipt is called the comparand against which edited journal content is checked (`…-130845.md:28,32`). Immutability prevents mutual mutation of settlement/call facts, but it does not manufacture the absent payload-identity member or carrier. Consequently the absolute non-promotion claim holds only for edits detected by m-9's actual integrity signals; the undetected class can inherit apparent clean trust under the accepted confusion-not-malice MVP limit.

Required revision: keep the governed-store seat-stamped `FROM` boundary unchanged, but narrow the session-content claim to the actual detection split. Structural/absent/checksum-mismatched/file-identity edits can be labelled or degraded through m-9; a well-formed checksum-recomputed, present, consistent payload edit is undetectable and may resume as clean at MVP. State that m-10's uneditable store preserves settlement/call facts but carries no payload identity and is not a current content comparand. Preserve the Step-4 provenance/versioning carry and the inability to distinguish sanctioned repair from corruption.

### M1-CLOSE3-R1-F2 — “Each strictly strengthened” contradicts the three answers

The heading says all three answers are strictly strengthened (`:26`), while answer (c) correctly says the courier-verifiability deferral is unchanged (`:29`). Answer (b)'s governed-store seat provenance is likewise unchanged because session content was never seat-stamped; m-10 immutability removes one edit surface but, per F1, adds no content witness. Only answer (a) is strictly stronger in the direct a-fortiori sense: fewer externally editable artifacts can disturb store isolation.

Required revision: classify the answers separately. State (a) strengthened by the narrower surface; (b) governed relay provenance unchanged, with m-10 settlement immutability as a helpful but non-content-identifying condition; and (c) unchanged. Remove the header and body claims that all three are strictly strengthened.

### M1-CLOSE3-R1-F3 — The named Step-4 “both stores” carry is m-9 plus m-10, not m-1 plus m-10

The return says Step-4's authorised rebase across both stores brings the m-1 governed relay store and m-10 settlement store into the rebase surface (`:40-41`). That substitutes the wrong first store. Master's route assigns the two editable-session halves to m-9's session-content journal and m-10's settlement snapshot/receipt identity (`…route3-editsm/DESIGN-orchestrator-planner-20260726-032916.md:47-55`). The forward constraint arises because journal record digests and m-10's stored identity must be re-anchored together; the source reconciliation says the same thing (`…esc1-fork-4/RECONCILE-orchestrator-planner-20260726-004504.md:75-77`). The governed relay store is not one of that named pair.

Required revision: restore the actual Step-4 carry as an authorised rebase across m-9's journal and m-10's settlement store. m-1 may state a separate conditional obligation: if any future design ever brings the m-1 governed relay store into an edit/rebase surface, D3 courier verifiability and seat-provenance re-establishment must be designed before that crossing. Do not present that conditional m-1 concern as the already-recorded two-app-store carry.

## Accepted portions

- **BR-INV′ passes.** The sole MVP external-edit surface is m-9's D1 journal; m-1's governed relay store and m-10's settlement store are excluded.
- **The at-rest split passes.** The descriptor-grain battery governs the journal file, not a SQLite column; live replacement races fail closed, while between-session edits are outside the held-descriptor window and enter m-9's resume handling.
- **The repairability scope note passes.** MVP text repair applies to the journal only; recovery from an unresumable m-10 store is the named post-product hardening carry.
- **The m-1 MVP supersession list passes.** No m-1 redaction, at-rest, secret-boundary, seat-identity, `FROM`, receipt, manifest, disposition, or wire member needs MVP supersession for this requirement.
- **The frozen-byte and phase boundaries pass.** The lane-2 component/base, locks, `receipt_conflict`, H-12, lane hold, and no-code/no-PLAN/no-T4 posture remain intact.

## Revision bar

Return fresh exact bytes that:

1. Preserve BR-INV′ and the accepted portions above.
2. Replace the absolute flagged/non-promotion assertion with the exact detectable-versus-undetectable MVP split.
3. Remove any claim that m-10's immutable rows currently provide payload identity or a content comparison carrier.
4. Classify (a) as strengthened and (b)/(c) as unchanged at their owned guarantee grain, while noting m-10 immutability's narrower supporting benefit.
5. Restore the Step-4 rebase pair to m-9 journal plus m-10 settlement store; keep D3 retirement only as a conditional if the governed relay store is ever brought into such a surface.
6. Preserve zero m-1 MVP supersession members, repairability scope, frozen hashes, `receipt_conflict`, H-12, and all downstream holds.

Issue a fresh uniquely-parented DESIGN for exact-byte review. Do not call the m-1 return owner-final or fold it into amendment r2 before an approving pair review.

## Verification

- Reproduced incoming SHA-256 `38070c9d7a7935974a4751b31abf66f1dc414578cf7a3f1439974641d50dd35b`, component `d34a7c475c5988f12872a5aaf599b201fe56970b766b935bdb42b1d2851ab1ef`, interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`, and stage-6 amendment `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Exact-file lint of the addressed planner relay passes with `--no-freshness`; `TO`, `FROM`, `IN_REPLY_TO`, `DESIGN_DOC_ID`, and lane location are valid.
- Re-read the closure dispatch's no-carrier option, current m-9 and m-10 owner halves, the route's ownership split and both-store forward constraint, the prior m-1 return, and the m-10 column-versus-file ruling.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no design/source contract, historical relay, frozen byte, `frank/` source, branch, commit, owner-final fold, amendment, lock, PLAN, T4, credential, provider action, external use, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK — exact-file and dispatch-root relay-lint.py exit 0 (`--no-freshness` used for the historical planner relay in root mode)
INDEX_LINT: pre-existing failure — `master/relays/INDEX.md:2405` has `20260726-131600` after `20260726-131630`; this review's row is unique at live EOF and was not rewritten around the historical ordering defect
Next requested action: m-1.planner folds M1-CLOSE3-R1-F1 through F3 only, re-hashes the standalone return, and sends fresh exact bytes for pair review; master keeps the m-1 owner-final fold and all lane-4 downstream gates held.
