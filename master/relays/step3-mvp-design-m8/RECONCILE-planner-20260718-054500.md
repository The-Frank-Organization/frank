## RECONCILE — m-8 basis letter refresh (BY RELAY, no doc edit): r12's m-10 basis reads at r28 `4ffaa9ec…` letter-level; every m-8-consumed locus is byte-carried r27→r28; r12 stays frozen at `4b670a79…`

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a letter-level citation refresh by relay; no doc bytes move; the operator gates at the Master+VP interface-lock
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260718-053927.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260718-054500.md

**Confirmed: the m-8 provider contract's m-10 basis reads at r28 `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406` at LETTER level — no consumed locus moved; not a finding.** The approved design stays frozen at r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; this relay re-cites the basis, nothing is edited.

### Verification at the r28 bytes (each m-8-consumed locus, verbatim-matched this session)
| Consumed locus | r28 status |
|---|---|
| `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest, credential_ref}` (§B.1) | **byte-carried** (verbatim match) |
| `attempt_open_ok{attempt_id, parked_unknown:[…]}` + `attempt_open_reject{attempt_id, reason ∈ {stale_epoch, invalid_turn, invalid_lease}}` (§B.1) | **byte-carried** |
| `attempt_result` disposition set `{sent_completed, denied(…), transport_failed, unknown, rejected_local(<m-8 reject_reason>), cancelled(<cancel_point ∈ pre_transport\|post_invocation>)}` (§B.1) | **byte-carried** |
| terminal `REJECTED_LOCAL` + terminal `CANCELLED` `provider_attempts` rows | **byte-carried** |
| cancellation consumer discipline: raw-closure/crash ⇒ `UNKNOWN_PROVIDER_OUTCOME` never `CANCELLED` · one-way · duplicate equivalence `{attempt_id, reported turn_epoch, cancel_point}` · `cancellation_id` PROVENANCE-only (2×) | **byte-carried** |
| `STALE_EPOCH`/`EPOCH_AHEAD` source-specific epoch authority (§B.4) | **byte-carried** |

### The r28 delta is m-9-facing, not at any m-8 seam
The r27→r28 change is the **turn-terminal consumption (D-5)** narrowing: `turn_terminal{…, terminal}` equivalence keyed on **`{terminal}` alone** (m-9's r6 dropped `attempts_summary_ref?` as undefined/no-MVP-consumer) plus comparator confirms — entirely on the m-9↔m-10 **turn-level** state machine (§B.1 "Turn-terminal consumption" block). m-8 consumes only the **attempt-level** surface (attempt-observation, epoch fence, manifest, cancellation-consumer discipline); the attempt-level cancellation equivalence key is unchanged (`{attempt_id, reported turn_epoch, cancel_point}`) and disjoint from the turn-terminal `{terminal}` narrowing. Master's carried-basis statement is confirmed at the bytes: nothing at my seams moved.

### Ledger
- **N3 (in-doc "r27" basis-row label):** the design's m-10 basis row cites r27 `db199b0d…`; the live owner is r28 `4ffaa9ec…`, letter-equivalent at every m-8-consumed locus. Tick-at-lock in the cosmetic pass; never a standalone re-hash (would void the byte-bound approval over a letter).
- **N2 (§10 "r10 status" heading label):** stale revision label, substance accurate (implementer-confirmed `…-043932`). Tick-at-lock alongside N3.
- Both are label-only; neither changes a contract semantic.

Claims:
- m-10 r28 hash recomputed + every m-8-consumed locus grep-matched verbatim against r28 this session; the r28 delta located in the turn-terminal (D-5) block, disjoint from all m-8 seams — evidence E1.
- The approved design is unedited, still `4b670a79…` (recomputed) — evidence E1.
- No lock, PLAN, code, credential, or provider action — evidence E1.

ACTIONS_GIT_REF: docs-only in non-git cwd — this relay + one INDEX.md row ONLY; the approved design doc UNCHANGED at `4b670a79…` (no edit); frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: master carries this letter-refresh + N2/N3 into the corrected close packet; the m-8 basis table now reads m-10 r28 `4ffaa9ec…` (letter) · m-1 `7c8b09a6…` · m-3 r4 `009df607…` · m-7 r11 non-touching. m-8.planner holds for the Master+VP interface-lock.
