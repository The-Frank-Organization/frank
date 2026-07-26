# m-9 S-1 content-ready receipt body — one-file-per-run successor (Route-5 / close5)

**Status:** PROPOSED — m-9-authored exact final body for the S-1 content-ready receipt under the operator-ratified one-file-per-run D1 re-scope (Decisions 1–3). Supersedes the body at `2026-07-22-relock-lane2-m9-delta.md:308` (`{turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}`). Member set fixed by operator Decision 3 (drop `segment_id`; keep `seq_hwm` + `generation_id`); the **name + derivation** are m-9-authored (this doc), the round-identity naming a joint m-9↔m-10 call. Awaiting m-9.implementer exact-byte review → m-10 byte-confirm → m-3 locator confirm → §D join re-sign. Binds no lock; r17 `01b885fe…` UNMOVED (this is a successor body for amendment r2, not an edit of r17).

## 1. The exact final body
The S-1 receipt is emitted by m-9 over CTRL-W after the round's records + its round checkpoint have `fsync`-linearized (unchanged trigger). Envelope frozen §A.2, unchanged:
```
envelope: { v: 1, chan: "ctrl-w", type: "content_ready", seq, run_id, turn_epoch }   # `re` ABSENT
body:     { turn_id, attempt_id, round_identity, seq_hwm, generation_id }
```
**Removed:** `segment_id` (dead — one-file-per-run has no segments to name; unconstructible).
**Renamed:** `marker_digest` → **`round_identity`** (naming rationale §3; joint call, m-9 proposes, m-10 concurs).
**Retained:** `turn_id`, `attempt_id` (identity keys); `seq_hwm` (m-3 locator + Route-2 frozen-interval bound); `generation_id` (m-10 fencing operand).

## 2. Per-member encoding + derivation
| member | type / encoding | derivation |
|---|---|---|
| `turn_id` | canonical string | identity key (unchanged) |
| `attempt_id` | canonical string | identity key (unchanged) |
| `round_identity` | 64 lowercase hex | **SHA-256 over the JCS of the ordered array `[ canonical_content(r) for r in the round's records first_seq…last_seq ]`**, where `canonical_content(r)` is r's §1.3 canonical JSON **excluding the advisory `record_digest`** (advisory per close3/Route-1 Ask-3) and with **no `prev_digest`** (chaining removed, Decision 1) and **no `segment_id`** term (Decision 3). Over the round's own records in the single per-run file; **no cross-segment/chain term.** |
| `seq_hwm` | canonical-decimal-uint64 string | the round's highest `seq` (= `last_seq`); the committed-end bound |
| `generation_id` | canonical string (opaque) | the writer-generation identity m-10 assigns; carried for m-10's fencing — a **validate-and-drop fence operand, NOT receipt evidence** (§5) |

## 3. The naming call — `marker_digest` → `round_identity` (bundled into the one amendment)
The derivation changes regardless (over the round's content in the one file, no chaining). The old name `marker_digest` implied a chained-marker-digest structure that no longer exists; **`round_identity` honestly names what the member now IS** — an opaque identity of the round that m-10 binds by **properties**, not structure. Since Decision 3 already changes the body (removing `segment_id`), **bundling the rename into the SAME amendment is one body change, not two** (m-10's own affordability point) — cheaper than a keep-now-rename-later. m-9 proposes `round_identity`; m-10 concurs or keeps `marker_digest` — either way the four properties + stored-ness below are the contract, and the name rides one amendment.

## 4. The four properties m-10 binds (preserved verbatim; derivation-only change is the free path)
`round_identity` is **stable per round · unique per round · byte-reproduced verbatim · equality-comparable**, and is **STORED in the receipt** (rev16 §2:39 — `receipt_conflict` is decidable only if the identity is stored). m-10 matches it **for equality only** and never reads its structure (rev16 §2:41 — "does NOT rest on that derivation"), so the derivation change costs m-10 nothing; only the member-set change (segment_id removal) + the optional rename are shape changes.

## 5. Equivalence, fencing, and `receipt_conflict` — the frozen validate-and-drop split preserved (M9-CLOSE5-R1-F1)
`generation_id` (and envelope `turn_epoch`) are **fence operands, NOT receipt evidence.** They are **validate-and-dropped** — the receiver validates the sender/epoch, then drops both from the equivalence comparison; neither ever contributes to duplicate equivalence or to `receipt_conflict`. This is the exact frozen predicate on both sides: r17 §2 excludes `{generation_id, turn_epoch}` from duplicate equivalence because they are fenced separately (`2026-07-22-relock-lane2-m9-delta.md:310-317`); rev16 persists the key + round-identity operand + locators but validate-and-drops `{generation_id, turn_epoch}`, and its ordered receiver compares the evidence tuple **before** stale-sender fencing (`2026-07-22-stage6-lane2-producer-delta.md:39/:41`).

The **reduced evidence tuple** is therefore exactly:
```
{ run_id, turn_id, attempt_id, round_identity, seq_hwm }
```
(envelope `run_id`; body `turn_id, attempt_id, round_identity, seq_hwm`). The frozen **first-match ordering** is preserved:
1. **equivalent duplicate** — a same-key `{run_id, turn_id, attempt_id}` input whose evidence tuple is byte-equal ⇒ idempotent, no second row;
2. **sender/epoch fencing** — evaluated on `generation_id` / `turn_epoch` (validate-and-drop); a generation/epoch difference is fenced here and **is not a receipt-evidence conflict**;
3. **`receipt_conflict`** — only a same-key input whose **evidence** tuple (`round_identity` or `seq_hwm`) differs ⇒ `receipt_conflict`, first-committed stands.

**Removing `segment_id` does not weaken decidability** — `round_identity` (over the round's content) already functionally determines the round, and `seq_hwm` is its committed-end; the old `segment_id` was a locator made redundant by one-file-per-run. `generation_id` stays a REQUIRED body member (§1) — required on the body, validated, and **NOT persisted** (validate-and-drop for fencing); only `round_identity` and `seq_hwm` persist as evidence. It is never evidence. m-10 byte-confirms this holds against its stored/equality/fencing semantics (`...close5-m10-1/...131600.md:28-34`).

## 6. Boundaries
No lock moved; r17 `01b885fe…` UNMOVED (successor body for amendment r2, authored fresh, not an r17 edit). `receipt_conflict` mechanism frozen (only the tuple shape changes, operator-decided). The §D join re-signs on this exact body once both pairs approve; only the re-signed artifact enters r2. H-12 hard-blocks external use.
