## SITREP — m-1 fidelity packet: the S2 store-touch proposal surface (10 items, design §4) for your shape verdicts before any S2 dispatch

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-fidelity-m1
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: ../.relays/s2/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md
FROM: s2.orchestrator-planner
TO: m-1.implementer
CC: operator, s2.orchestrator-reviewer
SUBJECT: S2 fidelity review request — the pair-APPROVED design's m-1 proposal surface (design §4, 10 items, at frank main@6e3b67f); per the s2-dispatch you are the fidelity reviewer for anything touching the locked m-1 store contract; NO s2 `DISPATCH IMPL` goes live before your verdict is on record

Per the s2-dispatch (m-1 keeps authority over the owed-item `record_kind`, the store layout, and store-API fidelity; m-1.implementer fidelity-reviews store-touches before their dispatch), this is the S2 fidelity packet. The s2 pair's design is pair-approved at r2; every store-layout/record-shape item in it is framed as a **PROPOSAL to you** — the pair fixes nothing.

**Review object:** `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md` **§4** (self-contained, 10 items) at `main@6e3b67f`, with mechanism context in §2 (D-3 genesis, D-5 quarantine, D-6 canonical sufficiency, D-7 projection, D-8 GC) and the grill/guide provenance in §8. Charter semantics (locked for this slice): `open = owed-record with no disposition-record`, materialize-first; m-7 guide answers on record (`.relays/s2/s2-guide-q1/SITREP-planner-20260704-004750.md`): genesis digest = pin-what-exists (Q1=(a)), owed-record operator-authored (Q2=(i)); the guide notes the genesis-record LAYOUT half is yours.

**The 10 proposal items (design §4 verbatim surface; your verdict per item):**
1. `record_kind: owed_item` header + payload `{owner, source, target_surface, disposition_path}`;
2. `record_kind: owed_disposition` + `disposes_owed: <relay_id>` (mirrors the S1 `parks_gate` idiom; one disposition per owed id, duplicate ⇒ typed reject);
3. genesis record — fixed `relay_id: genesis`, fields per m-7 :136, record #1 in `records/`;
4. incident record — HELD-class, references quarantined `relay_id` + failure class, compound single-pivot;
5. `quarantine/` store-root member; eviction = name-preserving rename of the corrupt file;
6. GC-marker record — body names collected segments;
7. journal segmentation — `journal/intake/<seq>.jsonl` + `journal/redo/<seq>.jsonl`, zero-padded monotonic `<seq>`, highest = active, size-based rotation;
8. derived-record body embedding (canonical-sufficiency) — outbox/derived records embed their item payload in body; **flags a real S1 latent gap the pair found: today the outbox record's payload lives only in its redo intent, so a collected redo segment would make the projection unrebuildable from canonical** — no envelope/layout change proposed, body usage only;
9. `projections/owed/OPEN.md` — derived, rebuildable open-set artifact (layout home yours);
10. pinned-config manifest + `engine.json` (store-adjacent conductor-owned members whose digest genesis pins; operator-ratified build config in the m-2-locked shape, no domain-author stamps — guide sharpening).

**Specific asks beyond per-item verdicts:** (a) the `record_kind` field home + token spellings (items 1–2) are wholly yours — correct freely; (b) item 5: what should `read(relay_id)` of a quarantined record return (typed error class — the pair proposes `checksum-mismatch`, I-PH-clean)? (c) item 7: segment naming/rotation layout fidelity vs your locked on-disk shape (m-1 §6 reuse-upstream-layout applies to relays/INDEX — journals are engine-side, but layout authority is yours); (d) item 10: placement of config members relative to the store root.

**Gate statement:** per the F2 conditions, NO s2 `DISPATCH IMPL` goes live before your verdict is on record in `.relays/s2/` (the s2 PLAN is drafted in parallel; a must-revise from you folds bounded + narrow re-review, the S1 F-M1-1 pattern). Deliverable: a lint-clean relay FROM your seat TO s2.orchestrator-planner with per-item verdicts + any required revisions, filed (operator-carried) into `.relays/s2/s2-fidelity-m1/`.

ACTIONS_GIT_REF: none — report-only fidelity request; this relay file + an INDEX row under gitignored .relays/.
FINAL_GIT_STATUS_SHORT: none — clean tree
