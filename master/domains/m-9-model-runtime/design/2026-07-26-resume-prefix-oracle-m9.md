# m-9 half of the resume-prefix oracle — one-file-per-run (Route-2)

**Status:** PROPOSED — m-9-authored owner-final contract for the **m-9 half** of the direct-prefix oracle: the extraction boundary, the canonical compared representation, and the frozen-interval proof. m-3 owns the E3 predicate, expected/actual independence, and the evidence locator (§3–4 below name m-3's members with the m-9 constraint on them).
Brings the planner-only Route-2 answer (`…-route2-oracle-m9/DESIGN-planner-20260726-034100`) to owner-final. Fresh design artifact (own `DESIGN_DOC_ID`, own SHA); binds no lock; r17 `01b885fe…` UNMOVED. Awaiting m-9.implementer exact-byte review → m-3 consumer-confirm.

## 1. The extraction boundary — what "the valid prefix" is, and how its end is fixed
Under the trimmed floor (one file per run · completeness-gated · settled-tool-round checkpoints; no chaining/rotation/boundary-equation):
> **The valid prefix = the ordered records from file start up to and including the `round_marker` whose `round_index` == the fixture's pinned `resumed_round_index`.**

Its end is fixed by that **pinned round checkpoint**, not by where the file ends. A round is in the prefix only if it is **complete** (all its records present and each record complete — the gating-floor completeness test, `2026-07-26-gating-floor-onefile.md` §1) and closed by its `round_marker`; an incomplete final record (torn write) or an in-flight partial round is **not** in the prefix. Completeness + the pinned round checkpoint determine the boundary totally — no chain/boundary-equation is needed.

## 2. The canonical compared representation
> **The ordered sequence of the valid-prefix records, each in its §1.3 canonical JSON form** (keys sorted lexically, no insignificant whitespace, UTF-8, one record per line). The comparison is over the record **contents** (authored bytes the harness holds on both sides), **not** over per-record digests.
- **Included:** all authored content members — `seq`, `kind`, `turn_id?`, `round_index?`, `ts_monotonic`, and the per-kind payload.
- **`ts_monotonic` is included** because the valid prefix is **read and replayed verbatim on resume, never re-stamped** — the actual prefix reproduces the frozen `ts_monotonic` byte-for-byte. (A future scenario that re-authors a prefix record would need a harness-supplied `ts_monotonic` rule; not the case for `xit-dur-1`.)
- **Excluded:** the per-record content checksum (`record_digest`) — advisory under Route-1, not a member of the compared content; `prev_digest` is gone (no chaining); `segment_id`/`round_identity`/`seq_hwm`/`generation_id` are S-1 **receipt** members (Route-5), not journal-record members, and are not part of this comparison.

The compared object is a **closed, ordered, canonical record-content sequence** — a bounded artifact, both sides in hand.

## 3–4. Expected artifact + actual locator (m-3's members; the m-9 constraint)
These are m-3's (independence + evidence locator). m-3's exact members, which the m-9 half is built to satisfy:
- **`frozen_prefix_ref`** — a **SHA-256 content address** of the frozen expected valid-prefix record sequence. **m-9 constraint:** the expected artifact is the fixture's **authored** valid-prefix record contents (a closed function of frozen bytes, **never harvested from a build run** — else the oracle is circular).
- **`boundary_seq`** — the `seq` of the pinned boundary. **m-9 reconciliation:** `boundary_seq` == the `last_seq` of `round_marker(resumed_round_index)`, and the S-1 receipt's **`seq_hwm == boundary_seq`** for the resumed round (the one exact mapping from m-9's round-marker boundary to m-3's locator).
- **the predicate** — "actual valid prefix `[start … round_marker(resumed_round_index)]` equals the frozen expected, record-for-record in canonical form (§2)." **`log_prefix_digest` is REPLACED** (not renamed, not kept as a digest) by this closed **`resume_prefix_match`** expectation object — a predicate over two bounded artifacts, not a digest member (`STEP-3-LANE4-PLAN.md:79` is the superseded member). Final field names are m-3's E3 call; m-9 confirms the reconciliation above holds against them.

Independence: expected is harness-derived from frozen contents; actual is the observed resumed journal — the two are compared, neither derives from the other.

## 5. The frozen-interval proof (load-bearing — VP4-F1)
frank keeps appending after it resumes, so an unbounded "compare the log" would compare a legitimately-grown file. The comparison binds the **frozen interval by construction:**
> The compared interval's upper bound is `round_marker(resumed_round_index)` — a value **frozen in the fixture input** (`predecessor_turn_id` + `resumed_round_index`), **not read from the file's current end**. The extractor walks from start and **stops at the pinned round_marker**; it does not read "to end of file." Records after it (rounds with `round_index` > `resumed_round_index` — post-resume appends) are **outside the interval by definition** and excluded **before they are ever read**.

So the interval is `[start, round_marker(resumed_round_index)]`, derived entirely from frozen fixture material; a legitimately-grown file cannot spuriously fail or be silently re-scoped. This is **not** the ordered `{seq, record_digest}` list in another costume — the compared term is record **contents**, with no per-record digest as a comparison operand.

## 6. Mismatch diagnostic
On failure the leg reports a **diff of the two ordered canonical record sequences** (frozen expected vs resumed actual) — both bounded, both in hand — which supplies per-record/position granularity for free (the operator's rationale for accepting loss-of-digest-granularity). Where/by whom the diff is produced is m-3's recorder/E3 call; the m-9 side supplies the extracted actual prefix (§1) in canonical form (§2).

## 7. Boundaries
Design-only m-9 half. No contract authored on m-3's axis (I state the disposition; the amendment is master's), no locked byte moved. r17 `01b885fe…`, lane-4 plan `60daac08…`, interface lock `cbd1893c…` UNMOVED. No `frank/`, no PLAN/T4. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.
