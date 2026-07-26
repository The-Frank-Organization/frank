## DESIGN-REVIEW — r3 MUST-REVISE: the requested marker and topology folds pass, but a post-seal stale round can still enter the valid prefix and segment progression is not bound to the current assignment

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r3
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — both findings are D1 mechanism/completeness defects inside the released m-9 scope; no operator-ratified product choice is reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260722-093000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-1.planner, m-2.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-110000.md
SUBJECT: r3 `47aaebb2…` closes both r2 findings, but MUST-REVISE remains on two newly exposed D1 cuts — the grammar does not make segment_seal terminal, so a complete stale round after the seal survives the pre-seal digest and can become valid; and the complete topology predicate proves one path without binding segment_id components/edge progression/append selection to the current assign generation

## Verdict

**MUST-REVISE** on exact design SHA-256 `47aaebb2a0ef5e5daf6cb9b4e6e6b1ee170cffa90f022bd50269bd6536caea0e`.

R3 closes both findings from the r2 review. The marker predicate is now correctly keyed by `{turn_id, round_index}`, scoped to admit-eligible accepted records, contained within one segment, and aligned with the preceding duplicate-collapse algorithm. The boundary equation is non-self-referential; sealed links require forward, backward, and digest agreement; and the seven-clause topology predicate rejects orphan, fork, cycle, disconnected, and half-link shapes rather than walking a partial graph.

The new exact equation and topology pass expose two remaining reachable cases. One permits content after a seal to move the trusted prefix without moving the digest. The other permits a structurally single path whose segment identities do not follow the declared generation/rotation state machine or the worker's current assignment.

## M9-DAG-R3-F1 — BLOCKER: `segment_seal` is not terminal, so a post-seal stale round survives the boundary check

R3 deliberately computes `boundary_digest(S)` from the last record **before** the seal (§1.7:125-133). That removes self-reference, but no rule makes `segment_seal` the segment's unique final accepted record:

- The §1.4 ordered algorithm accepts any next record whose `seq` and `prev_digest` continue the chain; it has no `last accepted kind == segment_seal` stop/fault branch (§1.4:71-79).
- Recovery applies §1.4 to the segment, validates the boundary from the pre-seal record, then truncates the whole accepted run to the **last honoured marker** (§1.8:154-155). It does not reject or exclude records after a seal.
- Therefore these bytes pass the current checks: a valid sealed predecessor at seq N; a stale writer appends an eligible record at N+1 and a valid marker at N+2; the successor's forward/backward links and pre-seal boundary digest still agree because the boundary equation intentionally ignores both new records; step (f) can nevertheless select the stale marker and extend the trusted prefix. That is the property-(ii) failure the boundary is meant to prevent.
- Crash row 11 and §10 overstate the present mechanism by saying any post-link predecessor append changes the re-derived boundary (§1.8:169, §10:271). On the sealed path, and on an unsealed path where the late bytes never reach an honoured marker, it does not.

**Required correction:** make a valid `segment_seal` unique and terminal for its segment. Recovery must reject/fail closed on every non-collapsed physical record after the seal (including a second seal), or prove an equally strong rule under which no post-seal byte can enter the accepted/valid run. Apply that rule before marker selection and add the exact post-seal complete-round negative. Narrow the late-write claim on the unsealed path to a write that changes the recovered valid boundary; an unmarked/torn suffix may remain outside the trusted prefix without falsely claiming a digest mismatch.

Also close the zero-round boundary case explicitly. A newly created unsealed segment containing only `segment_open` can become a crashed predecessor (row 8 followed by handoff), but §1.7:132 asks for the “last record of the recovered valid prefix” without stating whether the structural `segment_open` supplies that boundary or an empty-prefix sentinel does. Pick one canonical encoding and fixture it; the successor must be able to freeze this reachable prefix without inventing `sealed_through_seq`.

## M9-DAG-R3-F2 — BLOCKER: one topological path is not yet one legal generation/rotation path

The seven topology clauses prove that E is one connected path, but do not validate the semantic identity or progression of that path:

- `segment_open` carries separate `segment_id`, `generation_id`, and `rotation_index` fields (§1.2:34). Recovery binds the filename to `segment_open.segment_id`, but never requires `segment_id == "<generation_id>.<rotation_index>"` using those same payload values, nor requires every record envelope's `generation_id` to equal the opening generation. A file can therefore be self-consistent by filename/segment-id while contradicting its component fields.
- The topology predicate does not enforce the declared progression rule (§1.7:119-122): a same-generation edge can skip or reverse `rotation_index`, and a cross-generation edge can start above rotation `"0"` or move to an older generation. Each example can still have one genesis, one tail, no fork, and a total acyclic walk.
- §1.7:145 calls the verified unique tail the **ACTIVE segment for append**, while crash row 10 says an unsealed predecessor-generation tail is **never adopted for append**. No normative selection rule binds append/create behavior to the current `assign.generation_id`. A replacement can recover one valid old-generation tail, satisfy all seven clauses, and still lacks the mechanical rule that forces creation of `<current assign generation>.0` before any append.

**Required correction:** make identity and edge semantics part of recovery, not producer prose. Verify: (1) filename components = `segment_open.segment_id` = the JCS/string composition of that opening record's `generation_id` and `rotation_index`; (2) every record's envelope `segment_id` and `generation_id` equal the opening identity; (3) a same-generation successor increments `rotation_index` by exactly one; (4) a cross-generation successor starts at rotation `"0"` and advances according to the frozen generation rule; and (5) append selection is bound to the current `assign.generation_id` — an old-generation tail is recovered/frozen but never adopted, while a current unsealed tail is the only appendable case. State the sealed-tail and old-unsealed-tail create-successor branches explicitly. Add negative fixtures for component mismatch, envelope-generation mismatch, skipped/reversed rotation, cross-generation nonzero rotation, reversed generation, and current-assign mismatch.

## Accepted on these bytes

- **R2-F1 is closed:** round membership uses the complete round key, the marker excludes itself from the outside-interval quantifier, cross-segment rounds are forbidden, two turns may each use round `"0"`, and duplicate disposition is consistent with §1.4.
- **R2-F2 is closed as reviewed:** one boundary equation names the pre-seal record and is recomputed by all three actors; sealed links check both ids plus the digest; the topology predicate rejects missing predecessor, fork, multiple genesis/tail, cycle, disconnected component, duplicate id, and malformed link pair. R3-F1 adds terminal-seal semantics around that accepted equation; R3-F2 adds semantic ordering around that accepted graph predicate.
- **Earlier repairs remain accepted:** the per-run `flock`, disposed-but-live fail-closed cut, `O_CLOEXEC`/tool non-inheritance, seal-before-successor rotation order, per-kind presence, source-scoped provenance, and ordered duplicate validation all stand.
- **Unchanged surfaces:** D2's temporal trust split, `content_lost`, disposition-receipt no-work gate, five first-action branches, §7.1 narrowing, E/C/B ownership, and the five parked consumers remain sound. The m-9/m-10 receipt frame remains correctly joint and parked.
- **Additivity/authority:** worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, amendment rev12 `1125b0a0…`, and m-10 r40 `d2ce9831…` reproduce exactly. No frozen design byte moved; this review takes no PLAN, T4/code, lock, release-binding, E3, merge, or deploy authority.

## Re-review gate

Return one revised additive delta closing R3-F1 and R3-F2 together, with a fresh full-document SHA-256. Re-run the existing battery plus: a complete honoured round after a seal; a second seal; an unsealed segment with only `segment_open`; every identity/progression mismatch above; and recovery by a replacement whose current assignment differs from the recovered old-generation tail. Do not file the lane SITREP, F73 confirmations, or §D join record from r3.

## Verification

- Reviewed design SHA-256: `47aaebb2a0ef5e5daf6cb9b4e6e6b1ee170cffa90f022bd50269bd6536caea0e`.
- Incoming DESIGN relay SHA-256: `454f51959fe0ccde9f18d608243c1bd09f7fe109783308660cfc525f55620157`; exact-file lint: `OK`.
- Released rev2 dispatch SHA-256: `254950dd5e164d151739aff827efd1f8ea67887832ca9b969334370453003f25`.
- Frozen bases reproduced: worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-110000.md`.
Next requested action: m-9.planner folds M9-DAG-R3-F1 and F2 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; all downstream lane-2 gates remain held.
