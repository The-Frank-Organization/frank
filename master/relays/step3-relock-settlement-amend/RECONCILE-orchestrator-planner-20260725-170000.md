## RECONCILE — §D-settlement amendment OPERATOR-RATIFIED (recorded §8b, agent-authored + operator-cited): the exact packet `1fa71cb8…` + `5ec7a3d2…` is now architecture-of-record; the four bounded corrections are in force; the propagation matrix is OPENED — routed TO m-9 + m-10 for their fresh post-ratification pair-reviewed successors (m-2 unchanged). Downstream stays held; H-12 stands.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this RECORDS the operator's hash-bound ratification (the human gate itself); master does not self-ratify and authors no `FROM: operator` leg — the operator ratified in-session and this relay is the durable agent-authored, operator-cited record (§8b)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-163000.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner
CC: operator, master.orchestrator-reviewer, m-9.implementer, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: RATIFIED — operator hash-bound ratification of packet `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b` (amendment) + `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f` (bound m-2 cell), recorded §8b; the four §D-settlement corrections are in force; m-9 + m-10 now produce fresh post-ratification pair-reviewed successors per the amendment's owner fold matrix; m-2 unchanged; §D two-sided join follows the folds

## The ratification (recorded §8b — agent-authored, operator-cited; NOT a forged operator leg)
The operator ratified the §D-settlement amendment **in-session on 2026-07-25**. This relay is the durable record of that human-gate act (per §8b: master authors the record and cites the operator; master does not self-ratify and does not author a `FROM: operator` relay), exactly as the m-3 schema-version amendment was ratified on 2026-07-22.

**The ratified packet (byte-bound as ONE unit; both re-hashed byte-exact on disk at ratification):**
- `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` @ SHA-256 **`1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b`** (rev4)
- `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` @ SHA-256 **`5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`** (the bound `relay.submit` cell)

Both files are now **architecture-of-record and stay byte-exact at these hashes** — the ratified artifacts are frozen; their "ratified" status lives in this relay trail, not in edits to the files (an edit would break the ratified hash binding).

## Now in force — the four bounded corrections (additive to rev12 `1125b0a0…` / r40 `d2ce9831…`; r21 `4d3bd14e…` byte-frozen)
1. **D-4 Gate-2 claim honestly relabeled** — Gate 1 delivers the disclosure guarantee; Gate 2 = fail-closed validator + drift-detector over MVP-unreachable states (comparator bytes stay; the operator-visible *claim* is corrected).
2. **Run-wide parked-set restore + `MAX_PARKED_ROWS_PER_RUN = 512` + the new operator-visible terminal `parked_unknown_capacity_exceeded`** — the worker-independent disclosure guarantee is restored; the retirement transaction commits in full on overflow (truncation forbidden, every parked identity queryable) + commits `runs.state = FAILED` atomically; the two D-4 carrier shapes carry compile-time frame assertions (no legal frame attains `FRAME_MAX`; the proven bound is `FRAME_CONTENT_BOUND = 3,704,832 B`).
3. **§5-C `relay.submit` `canonical_resource` cell** — the form-schema-derived cell (`5ec7a3d2…`), total via a REQUIRED `form_digest`.
4. **`turn_failed` zero-attempt branch** — an explicit, bounded semantic supersession of r21 `:115`'s after-attempt restriction, for the single pre-attempt assembly-refusal branch; r21 stays byte-frozen.

## The propagation matrix — OPENED. m-9 + m-10 act now (per the amendment's owner fold matrix + timeless-fold rule)
Each **changed owner** produces **one fresh, pair-reviewed successor over its then-current owner artifact** (no pre-ratification working revision — incl. m-10's live rev14 `b96a1511…` — is a durable fold; it is ancestry only). Fresh full-byte pair review each; consumer confirmations + the §D join follow.

- **m-9 (`04422965…` is the current pair-approved base):** fold, in one batched F73 revision — Correction 1's **consumer-side** Gate-2 label; Correction 2's **full run-wide consumer + comparator restore** + the Correction-2 terminal's consumer posture; Correction 3's bound `relay.submit` cell (into §7-`relay.*`); and Correction 4 — **REPLACE the pair-approved §6 `:423-426` classification** ("after the bounded attempt(s) is DESCRIPTIVE … an owner clarification") with the amendment-controlled explicit semantic supersession. This IS the final batched m-9 revision that was held for this gate. Return the fresh full-byte hash on pair-approval.
- **m-10 (fold over the THEN-CURRENT owner artifact — NOT a relabel of rev14 as-is):** Correction 1's **producer-side** claim relabel; Correction 2's **run-wide producer** + the `parked_unknown_capacity_exceeded` terminal + the schema/presence rules + the **two compile-time frame assertions** + the fixtures (incl. the §2.3 boundary/multi-row/no-prefix predicate + the §2.4 production max-witness `≤ FRAME_CONTENT_BOUND` + reduced-table exact-fit). Fresh full-byte pair review.
- **m-2:** **nothing** — the bound cell `5ec7a3d2…` is byte-final; matrix row `None`; no new pair cycle unless its bytes move.

## After the folds — the §D two-sided join, then the sequence resumes
Once **both** m-9 + m-10 successors are pair-approved, the **§D two-sided join (m-9⇄m-10 + m-1 redaction) co-signs**; consumers (m-3, incl. its lane-2 basis r19 `92e08d09…`) rebase/confirm against the settled bytes. That **closes the lane-2 interface DAG.** Then: item A (extraction bundle + `bundle_sha256` + `STEP-3-EXIT-FIXTURES.json`) → lane 4 (the shorter stage-6 re-lock; its exit-completeness claim records N910 = documented MVP limit + r7-mirror deferred-v3, never "complete lane-2 coverage") → lane 5 (T4).

## Still held — ratification opened ONLY the propagation matrix
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy is licensed. **H-12 continues to hard-block external/untrusted/multi-tenant use.** The §D join, the integrated re-lock, item A, and lanes 4–5 remain ahead, each behind its own gate.

## Boundaries
No `frank/` action, no self-ratification, no forged operator leg, no fold performed by master (the pairs fold). The ratified amendment `1fa71cb8…` + cell `5ec7a3d2…` stay byte-exact. rev12 `1125b0a0…`, r40 `d2ce9831…`, r21 `4d3bd14e…`, m-9 r12 `04422965…`, m-10 rev14 `b96a1511…` UNMOVED. H-12 stands.

## Verification
Re-hashed on disk at ratification: amendment `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b` + bound cell `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f` — both MATCH the VP-approved packet (`…-160000`). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this ratification-record + propagation-dispatch relay + one INDEX.md row; the ratified amendment + bound cell stay byte-exact (NOT edited); no `frank/` action, no lock issued beyond recording the operator's ratification, no fold performed by master.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9 folds its final batched revision (Corrections 1/2/3/4 incl. the §6 `:423-426` replacement) and m-10 folds its half over the then-current artifact (Corrections 1/2 + terminal + frame assertions + fixtures), each to a fresh pair-approval; m-2 unchanged; then the §D two-sided join co-signs → the lane-2 DAG closes → item A (bundle) → lane 4 (re-lock) → lane 5 (T4). H-12 stands.
