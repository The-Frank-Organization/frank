## RECONCILE — §D-settlement amendment VP-APPROVED (rev4 `1fa71cb8…`, zero findings): the exact two-file ratification packet is presented to the operator for hash-bound ratification. Master does NOT self-ratify (§8b) — this relay requests the operator's ratification, it does not perform it.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — operator hash-bound ratification of the exact two-file packet; the VP approval is not ratification, and master does not self-ratify or author a `FROM: operator` leg
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-160000.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: operator ratification requested — the exact packet `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b` (amendment) + `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f` (bound m-2 cell); VP approved zero-findings at `…-160000`; on your ratification the propagation matrix executes (m-9 + m-10 fresh successors; m-2 unchanged) then the §D join

Operator — the §D-settlement amendment is VP-approved with **zero findings** (`step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-160000`, VP-review-r4). It is ready for your hash-bound ratification. I present it; I do not ratify it (§8b — master does not self-ratify, and I do not author a `FROM: operator` leg; your ratification is recorded agent-authored + operator-cited when you give it).

## The exact ratification packet (byte-bound as ONE unit; any byte change to either file voids the VP approval)
- **`master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md`** @ SHA-256 **`1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b`** (rev4)
- **`master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md`** @ pair-approved SHA-256 **`5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`** (the bound `relay.submit` cell)

Both re-hashed byte-exact on disk this session; both match the VP's approved hashes.

## What ratification enacts — the four bounded corrections (additive to rev12 `1125b0a0…` / r40 `d2ce9831…`; r21 `4d3bd14e…` byte-frozen)
1. **D-4 Gate-2 claim honestly relabeled** — Gate 1 delivers the disclosure guarantee; Gate 2 is a fail-closed validator + drift-detector over MVP-unreachable states (comparator bytes stay; only the operator-visible *claim* changes).
2. **Run-wide parked-set restore + `MAX_PARKED_ROWS_PER_RUN = 512` + a new operator-visible terminal `parked_unknown_capacity_exceeded`** — restores the ratified worker-independent disclosure guarantee; the retirement transaction commits in full on overflow (truncation forbidden, every parked identity queryable) and commits `runs.state = FAILED` atomically; the two D-4 carrier shapes carry compile-time frame assertions (no legal frame attains `FRAME_MAX`; the proven bound is `FRAME_CONTENT_BOUND = 3,704,832 B`).
3. **§5-C `relay.submit` `canonical_resource` cell** — the owner-authored, pair-approved, form-schema-derived cell (`5ec7a3d2…`), total via a REQUIRED `form_digest` (not `∅`).
4. **`turn_failed` zero-attempt branch** — an explicit, bounded semantic supersession (effective only on your ratification) of r21 `:115`'s after-attempt restriction, for the single pre-attempt assembly-refusal branch; r21 stays byte-frozen.

## What ratification does NOT authorize (all still held)
Ratification enacts the four corrections and opens the propagation matrix — **nothing further.** It does **not** license any DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or external use. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## On ratification — the sequence master runs next
1. **The propagation matrix executes (post-ratification, per the amendment's timeless-fold rule):** **m-9** and **m-10** — the only changed owners — each produce a **fresh, pair-reviewed successor over their then-current owner artifact** (no pre-ratification revision, incl. m-10's live rev14 `b96a1511…`, is a durable fold). **m-2 is unchanged** — its bound cell is a byte-final packet component, matrix row `None`.
2. **The §D two-sided join co-signs** after those folds.
3. Then the settled lane bases (incl. m-3 r19 `92e08d09…`, fully dispositioned) carry into the **integrated re-lock** → item A (bundle) → lane 4 → lane 5 T4.

## The ratification act (yours; recorded §8b)
To ratify, confirm the exact two-file packet (`1fa71cb8…` + `5ec7a3d2…`). I will then record your ratification as an agent-authored, operator-cited RECONCILE (never a forged `FROM: operator`), byte-binding both hashes, exactly as the m-3 schema amendment was ratified on 2026-07-22. If you want any change first, it voids the VP approval and I re-route.

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, `frank/` action, self-ratification, or forged operator leg. Packet hashes `1fa71cb8…` + `5ec7a3d2…` byte-exact; rev12 `1125b0a0…`, r40 `d2ce9831…`, r21 `4d3bd14e…`, m-9 r12 `04422965…`, m-10 rev14 `b96a1511…` (live, not pair-approved) all UNMOVED. H-12 external-use block stands.

## Verification
Hashed on disk this session: amendment rev4 `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b` + bound cell `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f` — both byte-exact vs the VP's approved packet. VP approval `…-160000` read at the bytes (verdict `approve`, zero findings, "route only this packet to operator"). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen/ratified byte moved, no amendment byte changed (rev4 stands byte-exact for ratification), no `frank/` action, no lock issued, no self-ratification, no operator leg authored, no fold performed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: operator ratifies the exact packet `1fa71cb8…` + `5ec7a3d2…` (or requests a change, voiding the VP approval); on ratification master records it §8b, then routes the propagation matrix to m-9 + m-10 for their fresh post-ratification pair-reviewed successors (m-2 unchanged) → the §D join → the integrated re-lock.
