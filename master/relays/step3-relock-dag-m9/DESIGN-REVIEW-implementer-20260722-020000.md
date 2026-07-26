## DESIGN-REVIEW — r1 MUST-REVISE: the D2/C/E/B surfaces are disciplined, but the D1 writer fence does not exclude cross-generation writers and the segment/marker recovery grammar is not yet deterministic

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — all findings are mechanism/completeness defects within the released m-9 design scope; no operator-ratified product choice must be reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260722-012000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-1.planner, m-2.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-020000.md
SUBJECT: four blocking D1 findings on r1 `2de214ea…`: the per-segment lock is not a cross-generation exclusive-writer lock; rotation creates the two-active-segment state it declares fatal; segment identity/implicit-seal linkage is incomplete; marker/provenance/recovery validation cannot mechanically prove the claimed exact prefix

## Verdict

**MUST-REVISE** on exact design SHA-256 `2de214ea70cbfcf21dc873b938633dc5ee61ba2f0f0c488f7f6e33c1be642a43`.

The producer-first staging, D2 temporal split, `content_lost` classification, first-action prohibitions, §7.1 narrowing, E/C/B ownership, and parked-consumer discipline are sound on these bytes. The blocker is the D1 core: property (ii) is not enforced by the selected lock domain, and the segment/marker grammar admits reachable states for which the claimed single deterministic valid prefix is either rejected by its own rules or not mechanically decidable.

## M9-DAG-R1-F1 — BLOCKER: the selected lock is not mutually exclusive across generations

The amendment requires an **enforceable exclusive-writer boundary** under which a retired generation cannot extend or corrupt the successor's trusted prefix; a bare generation label is insufficient (`STEP-3-STAGE6-AMENDMENT.md:179-195`). R1 does not yet provide that boundary:

- §1.6:66 locks **the active segment file**. §1.6:67 gives every generation its **own** segment file. A predecessor and successor therefore lock different inodes and both locks can succeed; the lock does not exclude the cross-generation writer the property is about.
- §1.6:65 says successor assignment is always ordered after predecessor reap, but frozen m-10 r40 §B.2 permits lease replacement after the prior generation is **reaped OR the operator disposes of it explicitly** (`2026-07-16-mvp-ipc-manifest-seam-contract.md:67-69`). The unreaped-but-disposed branch is omitted. On that branch the predecessor may remain alive while the successor opens and locks a different segment.
- A late predecessor append into its own segment carries the same generation id as that segment's `generation_open`, so §1.6:68's mismatch test does not classify it stale. If the successor has already treated that segment as implicitly sealed (§1.7:75 / crash row 5), the predecessor can move the predecessor prefix after the successor linked past it — exactly the review priority's unclosed interleaving.

**Required correction:** either (a) make the m-9 lock a single stable per-run lock identity shared by every generation and held over recovery plus the complete append/rotate lifetime, with one selected OS-lock primitive and explicit close/fork/inheritance semantics; acquisition failure remains fail-closed, including the disposed-but-live predecessor cut; or (b) select and route the amendment's joint m-10-ordered-segments branch. Per-generation segment locks may remain local corruption guards, but cannot be cited as the cross-generation exclusive-writer proof.

## M9-DAG-R1-F2 — BLOCKER: the rotation order creates the exact two-active-segment state the recovery rule declares fatal

§1.7:76 orders rotation as: fsync the new segment's `generation_open` → fsync its directory entry → fsync the old segment's seal → append to the new segment. After steps (1) or (2), both old and new segments are durable, unsealed, and carry the **same current generation id**. But §1.7:75 says exactly zero or one unsealed segment may match the current `assign.generation_id`, and two means fail-closed. Crash row 6 nevertheless calls all cuts during steps (1)–(3) ordinary and claims one prefix. Both statements cannot be true.

**Required correction:** pin a state/commit rule that keeps the new segment non-active and unreachable until the old segment's durable link commits (or choose another crash-safe order), and make active selection derive from that committed state rather than “unsealed + current generation” alone. Expand the crash table across at least: before new-file creation; after file creation but before directory fsync; after directory fsync but before old seal/link fsync; after the link commit but before first new append; and after first new append/marker. Every cut must name the one selected chain and disposition without contradicting the two-segment fault rule.

## M9-DAG-R1-F3 — BLOCKER: segment identity and implicit-seal linkage are not closed schemas

`segment_id` is load-bearing in the first-record seed (§1.3:50), seal digest (§1.7:73), links (§1.7:74), filenames/ownership (§1.6:67), and active selection, but it is absent from both the common envelope and the `generation_open` payload (§1.2:32, §1.3:48). Saying it is “derived from generation_id” cannot identify multiple size-rotation segments within one generation. The bytes therefore cannot prove which file a `generation_open` names or distinguish rotated siblings.

The crash-without-seal path is also underbound. Row 5 treats the predecessor as implicitly sealed at its recovered prefix, while the successor's link carries optional `prior_seal_digest`; it does not say what durable digest freezes the implicit boundary or how that digest is validated when no predecessor `segment_seal` exists. This becomes safety-critical in F1's late-write cut.

**Required correction:** add one canonical unique segment identity (including the within-generation rotation component) to the first-record/file identity and every link/seed that consumes it; bind it to the file selected on disk; define uniqueness/order. For an unsealed predecessor, persist and validate a successor-side link digest over the exact recovered predecessor boundary (or reject the chain) so later bytes cannot silently move the linked prefix.

## M9-DAG-R1-F4 — BLOCKER: round admission, provenance, and duplicate recovery are not mechanically total

Three related grammar gaps prevent §1.8's “single deterministic pass” and §1.11's property-(iii)/(v) claims:

1. §1.5 says `admits[]` lists **exactly** the records of a round, but admit-eligible records carry no `round_index`, and the honour rule checks only that every listed pair exists. It does not reject an omitted eligible record, an extra/cross-round record, a duplicate member, or wrong order. The system therefore cannot mechanically prove set equality between a marker and the round it claims to admit. A `tool_result` can precede a valid marker yet be absent from `admits[]`, leaving “in the chained prefix” versus “admitted to the valid prefix” contradictory at the settlement proof.
2. The common envelope requires `turn_id` on **every** record, but `generation_open` can precede any turn and `segment_seal` can span turns; no value/absence rule is defined. Conversely, the ratified identity-exact requirement says content binds its effect id plus source turn, while `input_item{role: tool_result}` has no `tool_call_id`, assistant-derived content has no `attempt_id`, and the design does not state which content kinds are exempt or how their source provenance is recovered.
3. §1.4 calls a later byte-identical duplicate `seq` a benign collapsed re-append, while the same replayed line's `prev_digest` normally differs from the immediately preceding record and §1.4 also says that mismatch terminates the prefix. The evaluation order and legal duplicate scope are unspecified, so the same bytes admit both “collapse and continue” and “stop here.”

**Required correction:** make round membership mechanically reconstructible (for example, a round id/index on every admitted record plus exact ordered-set validation, or an exact contiguous interval rule); pin per-kind common-envelope field presence and source-effect provenance; and define one ordered duplicate/chain-validation algorithm with fixtures for a single duplicate, a replayed suffix, a conflicting duplicate, and a duplicate after later sequence values.

## Accepted on these bytes

- **Additivity/authority:** the released rev2 `254950dd…`, amendment rev12 `1125b0a0…`, worker r7 `cb7ff970…`, and lifecycle r21 `4d3bd14e…` hashes reproduce; r7/r21 did not move. This remains DESIGN-only.
- **D2:** the settlement-time evidence property and resume-time evidence-plus-presence property remain separate; `content_lost` is only the post-inspection result; the disposition-receipt no-work gate is explicit.
- **First-action table:** all five ratified branches and prohibitions are present; the broker-cut relay identity routes through `uncertain` and any re-issue is fresh, never automatic.
- **§7.1 / authority:** the narrowing is correctly outcome-specific; the log contains content, not durable outcome truth or authority.
- **E/C/B:** `logical_surface_digest` is pre-lowering; m-8's lowered digest stays an independent root; `content_id`-unobtainable holds the E3 leg at `unknown`; nonexistent cwd rejects before spawn; `frozen_core_digest` is carried, not recomputed.
- **Parking:** m-2's component, m-10's C/D wires and disposition receipt, and the joint content-ready receipt contract are explicitly parked rather than treated as settled. The current m-9/m-10 receipt proposals differ, but both mark the shape JOINT-FINAL/parked, so that difference is not a blocker in this producer-only review.

## Re-review gate

Return one revised additive delta that closes F1–F4 together. Re-review must bind the new full-document hash and re-run the append/handoff/rotation battery against the exact revised state machine. The existing D2/C/E/B sections need change only if the D1 repair alters their consumed digest/path semantics. Do not file the lane SITREP, F73 confirmations, or §D join record from r1.

## Verification

- Reviewed design SHA-256: `2de214ea70cbfcf21dc873b938633dc5ee61ba2f0f0c488f7f6e33c1be642a43`.
- Incoming DESIGN relay SHA-256: `1c06e589e00c2c1b8553cb90fe987c0ccfa9c799c65cb890d29baad038322604`; exact-file lint: `OK`.
- Released rev2 dispatch SHA-256: `254950dd5e164d151739aff827efd1f8ea67887832ca9b969334370453003f25`.
- Frozen bases reproduced: worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-020000.md`.
Next requested action: m-9.planner folds M9-DAG-R1-F1 through F4 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; all downstream lane-2 gates remain held.
