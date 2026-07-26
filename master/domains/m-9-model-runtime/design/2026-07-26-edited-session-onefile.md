# m-9 edited-session recovery half — one-file-per-run successor (Route-3 / close3)

**Status:** PROPOSED — m-9-authored exact successor contract for the m-9 half of external-session-editing recovery under the operator-ratified one-file-per-run D1 re-scope (Decision 5: external editing is a REQUIREMENT; the edit surface is the m-9 journal ONLY; the m-10 settlement store is uneditable).
**This revision folds VP2-F4, resolved to option (b)** (the in-memory trust label had no executable lowering and is withdrawn; detected provider/tool mismatches route down the existing degrade path) — it **supersedes the r3 bytes `1f8ec7b6…`** on the F4 axis only.
Fresh design artifact (own `DESIGN_DOC_ID`, own SHA); binds no lock; r17 `01b885fe…` UNMOVED.
Awaiting m-9.implementer exact-byte review → m-10 close3 disposition reciprocal → m-3 evidence confirm → m-1 boundary confirm.

## 1. The honest detection limit (unchanged, correct)
At MVP no carrier can compare recovered payload **content** to m-10's frozen receipt identity: the receipt evidence tuple is `{run_id, turn_id, attempt_id, round_identity, seq_hwm}` and `round_identity` is matched for **equality only**, never re-derived over recovered content by m-10.
Therefore a well-formed edit — one that recomputes the advisory `record_digest`, keeps every record present, and stays outcome-consistent — is **undetectable** at MVP.
The m-10 settlement store stays outside the edit surface; canonical outcomes stay m-10-owned; `receipt_conflict` stays frozen; structural loss (a missing/short/unparseable record) still degrades via the existing path.

## 2. The no-carrier narrowing — the honest MVP contract, F4 resolved to (b) (M9-CLOSE3-R1-F2 · VP2-F4)
There is **no new carrier at MVP**, so m-9 makes **no durable-edit and no operator/E3-visible-edited claim.**

**F4 = (b) — the in-memory trust label is WITHDRAWN; detected provider/tool mismatches route down the existing degrade path.** VP2 found the r2 "local in-memory trust label" has **no executable lowering**: m-8's `input_item` enum is a **closed schema-law set** — `user_text{text}` · `assistant_text{text}` · `assistant_tool_call{…}` · `tool_result{tool_call_id, content:string}` · `reasoning_replay{envelope}`, with donor-hazard fields **deliberately absent** (`2026-07-17-mvp-provider-contract.md:44-47`). There is **no trust member on the wire shape**, so "untrusted-but-model-visible" cannot actually reach the model; "resumable with an invisible label" was **silent trust inheritance** for provider/tool content in practice. m-9 therefore resolves F4 to **(b)**: reuse the **existing `content_lost`/degrade path**, add no mechanism, change no wire, involve no m-8. Concretely:
- A **detected** advisory-checksum mismatch on **original-truth-bearing kinds** (`provider_output`, `tool_result`) is treated as **`content_lost`** → **`degraded` + `re_derive`**. Non-promotion is now **executable**: the suspect bytes are **re-derived, never presented** as the provider's/tool's actual output — the guarantee rests on a frozen path, not on a fictional label.
- A **detected** advisory-checksum mismatch on the model's own **transcript kinds** (`input_item`, `reasoning_replay`) — which make **no external-truth claim** — stays **`resumable`** (no label, no mechanism).
- An **undetected** well-formed edit (recomputed advisory `record_digest`, present, outcome-consistent) stays the **accepted documented limit** (§1) — the non-promotion claim is scoped to **detected** classes.
- m-10 receives **only the existing disposition** `report_resume_disposition.body = {turn_id, disposition, resume_action?}`, `disposition ∈ {resumable, degraded}`, `resume_action = re_derive` iff `degraded` (r17 §3:338-349). No new field, no `edited_since_write` event; m-10's manifest `uncertain` stays on its own separate axis (no m-9 wire class is invented).
- **Every Route-2 direct-prefix inequality remains `fail`.** An inequality becomes an authenticated-edit DEGRADED third outcome **only** when an authenticated edit record exists; absent that record, divergence stays `fail`.

**Retracted:** (i) the r2 mandatory in-memory trust label + its model-visible surfacing (no executable lowering — VP2-F4); (ii) the earlier `edited_since_write` recorded-mechanism + `RESUMABLE-with-edited-labels` claim. A durable/authenticated edit-provenance carrier + its schema/wire supersession (routed through m-8+m-10+m-3+m-1) is **Step-4** — only then can an edited provider/tool byte be **honored-with-provenance** instead of re-derived, and only then can a DEGRADED "edited" third outcome be claimed.

## 2.1 Total m-9 disposition / first-action table (M9-CLOSE3-R2-F1 · F4=(b))
Total over every recovered-content class on the existing `report_resume_disposition` wire (`disposition ∈ {resumable, degraded}`; `resume_action = re_derive` iff `degraded`). No in-memory trust label exists (F4=(b)); the non-promotion guarantee for provider/tool content is carried by the **degrade → re_derive** path, not by a label.

| recovered-content class | report pair (disposition / resume_action) | first action |
|---|---|---|
| clean, or undetected present+well-formed+consistent | `resumable` / — | replay verbatim into the assembled input (undetectable-edit limit §1) |
| advisory-checksum mismatch — `provider_output` / `tool_result` (present, well-formed) | `degraded` / `re_derive` | treat as `content_lost`; re-derive — the suspect bytes are **never presented** as original truth |
| advisory-checksum mismatch — `input_item` / `reasoning_replay` (present, well-formed) | `resumable` / — | replay verbatim (transcript kinds make no external-truth claim) |
| structural / completeness failure (torn, short, unparseable last record) | `degraded` / `re_derive` | re-derive from the last-good round checkpoint |
| referenced content missing | `degraded` / `re_derive` | re-derive |
| `objective_ref` re-resolves from immutable `admission_ref` (§3) | `resumable` / — | restore `objective_ref` from the byte-identical admission reference |
| `objective_ref` unresolvable (admission reference not byte-identical) | `degraded` / `re_derive` | re-derive |
| `workspace_snapshot` resolvable/matching (§3) | `resumable` / — | restore `workspace_snapshot` |
| `workspace_snapshot` mismatch / unresolvable (no MVP source, §3) | `degraded` / `re_derive` | re-derive |

The table is owner-final and deterministic: no row is optional, no class falls through, no row uses a disposition outside `{resumable, degraded}`, and no row depends on a label mechanism.

## 3. Recovery-source split (M9-CLOSE3-R1-F3)
The two references recovered at resume have **different authoritative sources** and must not share a row.

| reference | authoritative source | rule |
|---|---|---|
| `objective_ref` `{objective_locator, constraint_refs[]}` | the **immutable `turn_open.admission_ref`**, which carries objective/task identity (wake relay / operator input) | re-resolved from the byte-identical admission reference; a match restores `objective_ref` |
| `workspace_snapshot` `{workspace_root_id, snapshot_id}` | **NOT `admission_ref`** — the admission reference does not carry or derive workspace identity. At MVP m-9 has **no independent authoritative current source** for the snapshot. | a checksum-mismatching or unresolvable `workspace_snapshot` reference is classified **DEGRADED** (`disposition = degraded`, `resume_action = re_derive`, r17 §3). Do NOT infer workspace identity from the objective carrier. |

An independent authoritative workspace-snapshot source (so a mismatch could re-resolve rather than degrade) is Step-4 scope.

## 4. Boundaries
No lock moved; r17 `01b885fe…`, rev16 `3e3c5192…`, §D amendment, interface lock `cbd1893c…` UNMOVED. The m-10 settlement store stays uneditable; the edit surface is the m-9 journal only. H-12 hard-blocks external use.
