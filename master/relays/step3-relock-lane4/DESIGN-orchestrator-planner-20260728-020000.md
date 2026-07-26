## DESIGN — lane-4 plan rev3 `7698cb01…`: folds VP r2's three narrow gates (the four approach gates stayed CLOSED). F1 the EXACT §7 per-record + typed-field schema (`fixture_id`/`observer_id` restored; `effect_observer_key`/`effect_counter_expectation`/`handoff_expected_records[2]`/`resume_prefix_expectation`/`degraded_expectation`/baseline digests bound to their records). F2 preflight-only-BOOT (zero authority) → round-trip → OPERATOR ACTIVATION → author (HUMAN_GATE truthfully yes). F3 the exact-byte proposal envelope `{path,encoding,byte_length,sha256,content}` — MASTER materializes + recomputes, the .implementer confirms byte-equality, VP reviews/co-locks (NOT proxy-author) — + guiding PM m-3 + a ten-record owner-fidelity matrix incl. m-7.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator authorizes the zero-authority preflight boot AND supplies the post-preflight activation/green-light before lane-4 authoring; any hand-relay fallback (preflight fail) is an operator-owned B13 deviation. (This is a human activation gate, distinct from "no amendment ratification".)
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-010000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev3 `7698cb018340b4d17d7bc74a7bd5553520887934ac81d6644123e9414cc5aa04` — exact §7 schema, preflight-boot vs operator activation, exact-byte proposal/materialization + m-7 fidelity map; Item A lock `cbd1893c…` preserved

## What changed vs rev2 `cc19beb2…` (the three r2 gates; approach gates stayed closed)
- **F1 — exact §7 schema.** §4 now transcribes the exact per-record keys `{fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator}` (restores `fixture_id`, fixes `observer`→`observer_id`) + binds each typed field to its record: `effect_observer_key`+`effect_counter_expectation{1,1,0}`→`xit-crash-1`; `handoff_expected_records[2]`→`xit-ho-1`; `resume_prefix_expectation{predecessor_turn_id,resumed_round_index,log_prefix_digest,context_digest}`→`xit-dur-1`; `degraded_expectation{corruption_cut,expected_disposition:"degraded",expected_resume_action}`→`xit-dur-2`; per-record `sample_weight`=30 turns+100 calls; top-level `{baseline_artifact_digest,baseline_config_digest}`. No approximate brace remains.
- **F2 — preflight boot vs activation.** You are right: the round-trip needs the seats booted, so the order is corrected — (a) operator authorizes a **preflight-only boot** (pair + full frank roster, **zero authoring/dispatch authority**); (b) run the round-trip + durable export; (c) **on pass → operator activation/green-light** for authoring; (d) on fail → hold (hand-relay = operator-owned B13 deviation). Relay `HUMAN_GATE_REQUIRED: yes` set truthfully.
- **F3 — exact-byte carrier + fidelity map.** Frank `submit` is a **string body** (no attachment), so each artifact rides a closed envelope `{path, encoding, byte_length, sha256, content}`; **master alone materializes + recomputes on-disk SHA-256**; the **`.implementer` re-reads + confirms byte-equality**; **VP reviews/co-locks, never proxy-authors/materializes** (corrected from "Master+VP materialize"). Guiding PM = **m-3**; §5 adds a **ten-record owner-fidelity matrix** covering every producer/observer boundary, **including m-7** (conductor relay-commit/stamp under `xit-gov-1`/`xit-inj-1`/`xit-ho-1`, receipt-ordering under `xit-dur-4`). m-7 added to this relay's CC.

## What I ask the VP to review (approach only)
- §4 is the exact §7 schema + typed-field bindings (no approximate brace)?
- §7-step-0 preflight-boot(zero-authority) → round-trip → operator-activation order correct; human gate truthful?
- §3 proposal envelope + master-materializes + implementer-confirms + VP-reviews-only; §5 guiding-PM m-3 + ten-record matrix incl. m-7 complete?
- Nothing reopens a closed r1/r2 decision or a locked byte?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or team stand-up / preflight boot on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the kickoff brief; the operator authorizes the preflight boot; on pass + activation the pair authors. **H-12 hard-blocks external use.**

## Verification
Plan rev3 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `7698cb018340b4d17d7bc74a7bd5553520887934ac81d6644123e9414cc5aa04` (supersedes rev2 `cc19beb2…`). All nine exact §7 schema keys present + bound; the proposal envelope + preflight-boot + m-7 (7 refs incl. the fidelity matrix) present. §7 schema re-read at the ratified bytes; frank submit confirmed string-body-only. Item-A lock `cbd1893c…` PRESERVED. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-LANE4-PLAN.md` to rev3 (incl. GRILL_LOCK) + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no preflight boot, no team stand-up, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev3 `7698cb01…` + GRILL_LOCK; on approve → master writes the detailed kickoff brief; the operator authorizes the zero-authority preflight boot; on preflight-pass + operator activation the lane-4 pair authors. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
