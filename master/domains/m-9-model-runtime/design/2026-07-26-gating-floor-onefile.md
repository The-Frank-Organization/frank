# m-9 D1 gating floor + checksum-advisory posture — one-file-per-run (Route-1)

**Status:** PROPOSED — m-9-authored owner-final contract for what GATES resume vs what is ADVISORY in the one-file-per-run D1 journal, under the operator-ratified re-scope (Decisions 1–3) and the confusion-not-malice / MVP-minimality rails.
Brings the planner-only Route-1 answer (`…-route1-m9-ans/DESIGN-planner-20260726-033600`) to owner-final. Fresh design artifact (own `DESIGN_DOC_ID`, own SHA); binds no lock; r17 `01b885fe…` UNMOVED. Awaiting m-9.implementer exact-byte review.

## 1. The gating floor (what a resume is allowed to depend on)
Under the trimmed floor (one file per run · no chaining/rotation/boundary-equation/terminal-seal — Decisions 1–3), the durability elements that **GATE** resume reduce to exactly three:
1. **last-record completeness** — the final journal record parses as a complete record and (line-oriented, one record per line, `\n`-terminated, §1.3) carries its terminating `\n`; a torn/short/unparseable final record is caught here;
2. **the settled-tool-round checkpoint** (`round_marker{round_index, first_seq, last_seq}`, §1.5) — the last-good boundary a resume rewinds to;
3. **the per-run writer fence** (dedicated `session.lock`, acquired `O_CLOEXEC` + `flock(LOCK_EX|LOCK_NB)` before any journal open/read/attach — the close4 fencing contract `2026-07-26-fencing-observable-onefile.md`).

Nothing else gates. No `prev_digest` chain, no cross-segment boundary equation, no terminal seal is required for a correct resume under the floor.

## 2. Per-record content checksum is ADVISORY, not gating (Route-1 Ask 2/3)
The per-record content checksum (`record_digest`) drops from a gating operand to **advisory**:
- **Torn write is caught by completeness, not by the checksum.** A torn write = an incomplete final record (missing `\n` / unparseable); completeness (§1, floor element 1) detects it for free. The content checksum is therefore **not** the torn-write/durability mechanism.
- **What the checksum still does (its residual, advisory job):** it is a **bit-rot diagnostic** — it catches a *complete, well-formed, terminated but silently byte-flipped* record that completeness cannot. Under confusion-not-malice + single-turn + local filesystem, bit-rot is a rare hardware event, not an admitted adversary; a real harness on a plain file carries the identical residual (an engine-backed store gets per-page checksums for free). So the checksum's residual is a diagnostic, not a durability gate.
- **The checksum's edit-detection role routes to the degrade path, not a label (F4=(b)).** A detected checksum mismatch on **complete, well-formed** content is an **edit**. Per the F4=(b) resolution in the edited-session contract (`2026-07-26-edited-session-onefile.md` §2/§2.1): a detected mismatch on `provider_output`/`tool_result` → `content_lost` → `degraded`+`re_derive` (suspect bytes re-derived, never presented as original truth); a detected mismatch on `input_item`/`reasoning_replay` → `resumable` (transcript, no external-truth claim). **There is no in-memory trust label** (it had no executable lowering — VP2-F4). So "advisory" means: the checksum fires a bit-rot diagnostic and, on a complete-record mismatch, hands the class to the existing degrade path — it never gates resume by itself and never silently promotes edited bytes.

## 3. The Ask-1 interface-lock catch (recorded outcome, no unilateral edit)
The Route-1 dispatch asked m-9 to annotate `2026-07-19-mvp-full-worker.md:88/:155` in place. m-9 **declined**: that file hashes to frozen worker r7 `cb7ff970…` and is a named interface-lock constituent (`STEP-3-INTERFACE-LOCK.md:45`); an in-place edit re-hashes it and **voids the interface lock `cbd1893c…`** (owner-path rule, `CYCLE-PLAYBOOK.md:222` / [[fence-row-vs-owner-path]]). The provisional supersession of the worker's fresh-start loci is recorded via **governed carrier** (the r17 §5 precedent) and rides the **D1 re-scope re-lock wave** — worker r7 is re-hashed once against the settled shape and the interface lock re-binds in the same act, never a standalone lock-void. This doc records that outcome; it edits no locked byte.

## 4. Boundaries
Design-only owner contract. No owner/locked byte edited — worker r7 `cb7ff970…` and interface lock `cbd1893c…` UNMOVED. No amendment, no fixture. r17 `01b885fe…`, m-10 rev16 `3e3c5192…`, §D amendment `1fa71cb8…` UNMOVED. `receipt_conflict` frozen. No `frank/`, no PLAN/T4. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.
