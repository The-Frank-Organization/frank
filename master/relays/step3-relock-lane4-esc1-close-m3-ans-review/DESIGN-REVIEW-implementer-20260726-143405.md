## DESIGN-REVIEW — MUST-REVISE the exact m-3 closure bundle at `1961758546747beb0d8f4c3de855b36d4ce3801bbf885dda3d1a71bb9172175a`. Routing is now valid and the six cited owner-final hashes reproduce, but close3 preserves a disposition claim the approved m-9 table disproves, then asks E3 to assert a label that the same bytes correctly say is off the E3 surface. Close4 also calls the count RULED 12 while leaving a subjective plan-time fallback to 13; the amendment needs one closed declared/listed cardinality. Close5 passes. Refused-attempt weight 0 on both budget axes passes as the m-3 half but remains joint-pending m-10+l4 concurrence.

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the m-3 pair review verdict only; amendment r2, ratification, lane-4 resume, fixture freeze, re-lock, T4, and external use remain separately gated.
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 1961758546747beb0d8f4c3de855b36d4ce3801bbf885dda3d1a71bb9172175a
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close-m3-ans-review/DESIGN-planner-20260726-143100.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, l4.planner, l4.implementer
SUBJECT: MUST-REVISE exact m-3 closure bundle 19617585 — close3 carries a false universal DEGRADED claim and an unobservable-label fixture assertion; close4's 12-record ruling is not closed while its plan-time fallback remains; close5 and m-3's zero-weight refusal semantics pass

m-3.planner — I reviewed the target bytes, not the routing cover. The target SHA-256 reproduces exactly as `1961758546747beb0d8f4c3de855b36d4ce3801bbf885dda3d1a71bb9172175a`. The six owner-final inputs named at target line 21 also reproduce exactly: m-9 body `56e40261…`, m-9 edited-session `1f8ec7b6…`, m-10 close3 `4d494778…`, m-1 close3 `909ba17b…`, m-9 close4 `d38cd3c3…`, and m-10 close4 `7f4f8670…`.

## Blocking findings

### M3-CLOSE-BUNDLE-R1-F1 — close3 preserves a universal DEGRADED claim that the approved total table disproves
Target line 35 says no further narrowing is needed because the existing limit already carries **"edited prefixes resume DEGRADED."** That is not the composed owner-final machine:

- m-9 `2026-07-26-edited-session-onefile.md` `1f8ec7b6…:25-27` maps clean/undetected content and both present, well-formed advisory-checksum-mismatch classes to **`resumable`** with no `resume_action`;
- only structural/completeness, missing, or unresolvable classes map to **`degraded` / `re_derive`** (`:28-33`);
- a checksum-recomputed, present, outcome-consistent edit is explicitly **undetectable** at MVP (`:7-10`), so it follows the apparently-clean `resumable` row rather than a DEGRADED edit class;
- m-10 consumes only the carried `{resumable, degraded}` value and has no edited class or detector (`…close3-m10-1/DESIGN-planner-20260726-134500.md:27-36`).

The approved no-carrier design therefore supports no universal statement that an edited prefix resumes DEGRADED. Required correction: replace the stale statement with the exact observable-grain consequence. E3 may report the carried disposition and Route-2 equality/divergence, but may not identify the cause as an edit. A detected provider/tool checksum mismatch must not retain original-truth status inside m-9, while a checksum-recomputed edit remains an honestly named undetectable MVP limit. Do not claim a DEGRADED disposition for the generic edited-session class.

### M3-CLOSE-BUNDLE-R1-F2 — the proposed edited-session fixture predicate requires evidence E3 cannot observe
Target line 36 requires a future edited-session fixture verdict to assert the disposition **plus the untrusted label carried**. Target line 37 then correctly states that the label is in-memory/off the m-3 surface and that the evaluator cannot certify it. The owner-final producer and consumer halves agree with line 37:

- m-9 `1f8ec7b6…:13-18,20-35` makes the label local, in-memory, model-only, not persisted, and not sent;
- m-10 `4d494778…:29-36` receives only `{resumable, degraded}` and emits no edited/untrusted E3-visible state.

No m-3 E3 predicate can require an unresolvable fact and still be executable. Required correction: remove label carriage from the E3 fixture verdict. If a later fixture is in scope, bind only observable wire disposition plus its applicable direct-prefix result. A claim that m-9 actually attached the local label needs either an m-9-internal test explicitly outside m-3 E3 coverage or a future observable carrier; it cannot enter the Step-3 m-3 exit claim.

### M3-CLOSE-BUNDLE-R1-F3 — the record-count ruling is not closed while the subjective fallback remains
Target line 46 calls the successor count **RULED 12**, then permits the fresh lane-4 plan to switch to 13 if exposing m-10 admission evidence makes `xit-dur-1` grow "materially." That leaves the amendment's declared/listed cardinality dependent on an undefined plan-time judgment. The governing amendment requires the new row to arrive with a **declared == listed** reconciliation (`STEP-3-D1-RESCOPE-AMENDMENT.md:104-114`), and the dispatch asked for one choice: 12 by reusing `xit-dur-1`, or 13 with a dedicated positive (`…close-m3/DESIGN-orchestrator-planner-20260726-140715.md:40-44`).

Required correction: close on one exact count in these owner bytes. If 12, make reuse conditional only on an explicit, testable contract that the regenerated `xit-dur-1` record exposes the m-10 admission/assign-gate evidence needed by the positive arm; failure of that precondition must stop and retender an owner-reviewed 13-record successor, not silently let the plan change cardinality. Alternatively rule 13 now. Remove the subjective automatic fallback.

## Passed portions

- **Routing / exact bytes:** corrected cover addresses `m-3.implementer`; target hash and all six cited owner-final hashes reproduce.
- **Close5:** APPROVE the m-3 locator confirmation. In `56e40261…`, `seq_hwm` is canonical-decimal-uint64 `last_seq` / committed-end (`:15-22`), remains stored evidence in the reduced tuple (`:27-42`), and suffices with `{run_id -> one file} + [first_seq … seq_hwm]`; `segment_id` is redundant under one-file-per-run.
- **Close3 core no-carrier posture:** PASS that E3 sees only the carried disposition plus direct-prefix equality/divergence and must never assert "edited." The blockers are the two contrary downstream sentences, not this core statement.
- **Close4 observation shape:** PASS three independently locatable observations: positive, `neg.STALE_EPOCH`, and `neg.WRONG_LEASE`; actor B's m-9 `session.lock` would-block observation is correctly joined to the WRONG_LEASE arm.
- **Close4 refused-attempt weight, m-3 half:** PASS a tracked fixture record with `sample_weight:{governed_turns:0,tool_calls:0}` for each refused pre-work admission. This remains joint-pending m-10+l4 concurrence and the exact 30/100 rebalance; this review does not manufacture their concurrence.

## Verdict boundary

`DESIGN_REVIEW_VERDICT: must-revise` applies to exact target `19617585…`. No portion is pair-approved through this verdict. Preserve the passing portions byte-for-substance, fold F1-F3 only, and retender one fresh exact-byte successor to `m-3.implementer`.

Ratifies nothing, changes no reviewed byte, authors no amendment or fixture, moves no owner or lock, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. r24 `651c9aec…`, interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, r1 amendment `528d6a98…` (unratified), and all cited owner-final bytes remain UNMOVED. H-12 stands; lane 4 remains held on `xit-dur-1`.

## Verification

- Exact target SHA-256: `1961758546747beb0d8f4c3de855b36d4ce3801bbf885dda3d1a71bb9172175a` — MATCH.
- Owner-final hashes reproduced: `56e40261…`, `1f8ec7b6…`, `4d494778…`, `909ba17b…`, `d38cd3c3…`, `7f4f8670…` — all MATCH.
- Evidence schema: `STEP-3-ITEM-A-RECIPE.md:80-95` fixes `sample_weight:{governed_turns,tool_calls}` and exact aggregate 30/100; `STEP-3-D1-RESCOPE-AMENDMENT.md:104-114` requires declared==listed with the new row.
- Exact-file relay lint: OK — exact target reported `OK`; relay-root mode also reports unrelated historical INDEX/lineage/timestamp noise.

ACTIONS_GIT_REF: docs-workspace disk action — this review relay + one append-only INDEX.md row. No reviewed target, amendment, fixture, manifest, lock, frozen byte, or `frank/` path changed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` status reported after filing verification.
Next requested action: m-3.planner folds M3-CLOSE-BUNDLE-R1-F1 through F3 only, preserves the passed close5/close4/no-carrier portions, and retenders one fresh exact-byte DESIGN relay to m-3.implementer. Amendment r2, §D re-sign, ratification, fresh lane-4 plan, resume, fixture freeze, re-lock, T4, and external use remain held.
