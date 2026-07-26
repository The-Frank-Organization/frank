# STEP-3 D1 RE-SCOPE AMENDMENT — additive supersession, r2

**Status:** PROPOSED **r2** by `master.orchestrator-planner`, 2026-07-26. **Supersedes r1** (`528d6a98e81497cac6300de84faae3e7deb6ebbc7077a8e72634a891f71cccbc`, rejected at VP review `step3-relock-lane4-esc1-amend-vp-1` on nine findings, all folded here). **Awaiting VP exact-byte review, then operator ratification.** Ratifies nothing by existing.

**Form — additive, on the MVP-amendment precedent.** This record supersedes named fragments of the artifacts below while **every superseded file stays byte-exact.** No lock constituent is edited: `STEP-3-STAGE6-AMENDMENT.md` is interface-lock row 54 at `1125b0a0…`, and `2026-07-19-mvp-full-worker.md` is **row 45 at `cb7ff970…`** — an in-place edit to either voids the interface lock `cbd1893c…`.

**Authority basis.** Operator Decisions 1–8 (E0-class decision record `…-esc1-ratify-3/RECONCILE-orchestrator-planner-20260726-031526.md` @ `bda1c941…`, VP-approved) plus the operator's edit-surface ruling (edit surface = the m-9 session journal ONLY; m-10's settlement store effectively uneditable; recovery involving that store = post-product hardening carry). Evidence of direction, **not transferable authority**; this amendment is master's proposal and requires its own ratification.

---

## §0 — The binding set: pair-approved owner-final artifacts ONLY (VP1-F1 folded)

Every normative input below is **owner-authored, implementer-exact-byte-approved** — none is a planner-only return. The contract text lives in these artifacts; this amendment binds them by hash and does not re-transcribe them.

| artifact | owner-final SHA-256 | pair approval |
|---|---|---|
| m-9 S-1 receipt body doc `2026-07-26-s1-receipt-body-onefile.md` | `56e40261fc80d209…` (full hash in the doc's approval) | `…-close5-body-m9/DESIGN-REVIEW-implementer-20260726-135539.md` |
| m-9 edited-session doc `2026-07-26-edited-session-onefile.md` | `1f8ec7b6c99c63ca…` | same approval chain (close3 r3, `…-135539`) |
| m-9 close4 writer-fence observable | `d38cd3c3775ed6fc77e048292dafa6cc113b4fb5a8b16b756bf55e3fbeaeb668` | `…-close4-fencing-m9/DESIGN-REVIEW-implementer-20260726-134146.md` |
| m-10 close3 disposition half rev3 | `4d494778a16f7eaa9044f921375db8735df50a876a1a3fdea26486713ca7325a` | `…-close3-m10-1/DESIGN-REVIEW-implementer-20260726-134907.md` |
| m-10 close4 observation-shape rev3 | `7f4f8670de86541b455c5c19a99f8c89acf1b703cd8bffd1e43cbc56137dc0ea` | `…-close4-m10-1/DESIGN-REVIEW-implementer-20260726-134908.md` |
| m-10 close5 byte-confirm of `56e40261…` | `92c9b3a8534d1f4fedf53783a3daa3ed73f990f3e1648806e49824affe1ee6c1` | `…-close5-m10-1/DESIGN-REVIEW-implementer-20260726-142005.md` |
| m-1 boundary re-confirm rev4 (BR-INV′) | `909ba17b229c66f2740bcfce3934d00edb7fc27736f91985ee5a4d6cc6377d9c` | `…-close3-editsm-m1/DESIGN-REVIEW-implementer-20260726-135900.md` |
| m-3 consolidated closure bundle | `7d7b6dbebab6377da344f91b7f1b76bf6964ac8175003be7e6bb491306a6239b` | `…-close-m3-ans-r2/DESIGN-REVIEW-implementer-20260726-145429.md` @ `b73085e6…` |
| **§D join RE-SIGN, m-9 half** | `ef72c732fdf92217f0d1eeea2f53e33b168356ed0d0105882b3370045969581a` | co-filed act over `56e40261…` |
| **§D join RE-SIGN, m-10 half** | `9f8c290f3bf37af71f55419b489c9255610c3f7d683dff061e3e83efd2a9f8dc` | + consistency confirm `edd2351e…` |
| m-10 concurrence (12-record coupling + 0/0) | `f3e0c2ae99cd7e46d904eda199a7c73e0eb8fafc6e9a4bc9e050efd54ccc7e6c` | report of a closed joint half |
| l4 design-level concurrence (30/100 recomputed) | `234daf6f6dfb4e81e8e56bdbb8bdfcedcb829ff2e019d2099c53696daa9d3c24` | design concurrence; lane 4 stayed held |

Boundary/evidence returns from the route round (m-1 route-3, m-3 route-* halves, m-8/m-9/m-10 scope answers) remain **cited evidence**, not normative contract text — the normative successors above supersede them where they overlap.

**Governing hashes at authoring, all UNMOVED:** interface lock `cbd1893c…` · stage-6 amendment `1125b0a0…` · §D-settlement amendment `1fa71cb8…` · lane-4 plan `60daac08…` · worker r7 `cb7ff970…` · m-9 r17 `01b885fe…` · m-10 rev16 `3e3c5192…` · m-3 r24 `651c9aec…` · m-8 `4b670a79…`.

---

## §1 — The D1 floor (Decisions 1 + 2)

**RELEASED from the MVP:** `prev_digest` hash-chaining · size-triggered segment rotation · the cross-segment boundary equation · the terminal seal · the generation-as-chained-segment model, replaced by **ONE FILE PER RUN**.

**The GATING floor is `{last-record completeness · round checkpoint · per-run writer fence}`.** Completeness alone covers torn writes. The per-record checksum is **advisory and not ignored**: a bit-rot diagnostic that also fires the §4 edit classification. Every retained element ties to an admitted hazard; **none rests on tamper-evidence, and this amendment claims no tamper-evidence anywhere** (the claim is formally retired from frank's defensible core — PRIOR-ART.md §2d).

**Grounds for releasing chaining, all three recorded:** malice-only over checksum + `seq` contiguity (m-9, against its own §9 survey); internal self-binding is redundant where the exit leg has a frozen oracle — robust to scale (m-3); and it would reject the operator's own `bivpak` repack as a forgery (operator, decisive).

**§1a — Named re-entry test (operator, verbatim):** *"we can bring them back in the future if agents start tampering with past session data, somehow, but very unlikely."* **Runtime self-integrity is OUT OF SCOPE** by the same decision.

---

## §2 — The S-1 receipt body (Decision 3), with the §D join RE-SIGNED

**Superseded:** the co-signed receipt body at m-9 r17 §2:308, `{turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}` — the exact fragment both re-sign halves name.

**Successor (m-9-authored, m-10-byte-confirmed, join re-signed):** **`{turn_id, attempt_id, round_identity, seq_hwm, generation_id}`** per `56e40261…`:
- `segment_id` **REMOVED** — no sibling file to disambiguate; `run_id` names the file (m-3 concurs: not required by its locators).
- `marker_digest` → **`round_identity`**: renamed with a derivation-only change beneath (SHA-256 over JCS of the round's record content in the single file, no chaining), **stored**, carrying the four properties m-10 binds — stable per round · unique per round · byte-reproduced verbatim · equality-comparable. `receipt_conflict` stays total/decidable over the reduced evidence tuple.
- `seq_hwm` **RETAINED** — canonical-decimal-uint64, the round's `last_seq`/committed-end: m-3's locator bound AND §3's frozen-interval stop.
- `generation_id` **RETAINED** — m-10's validate-and-drop fencing operand, excluded from equivalence.

m-1's redaction leg is **not body-shape-sensitive** (stated and grounded in both re-sign halves, not assumed). S-2/S-4/S-5 unchanged.

---

## §3 — The direct-prefix oracle (Decision 4)

**Superseded:** `log_prefix_digest` as a typed digest member of `resume_prefix_expectation` (`STEP-3-STAGE6-AMENDMENT.md:383`; lane-4 plan `:79`). **REPLACED** by a closed expectation object under the typed predicate **`valid_prefix_matches_frozen_expected` v1** — a predicate over two bounded artifacts.

- **Extraction boundary (m-9):** records `[start … the round_marker for the fixture-pinned resumed_round_index]`, completeness-gated, no chaining.
- **Canonical representation (m-9):** the ordered valid-prefix records' canonical JSON content.
- **Frozen-interval stop:** the actual-side read halts at the pinned **`seq_hwm`** — the post-resume append tail (`seq > seq_hwm`) is excluded **by construction**, never by where the file ends at read time.
- **Locator (m-3):** `{run_id → the single run file} + [first_seq … seq_hwm]`.
- **Expected side:** a closed function of the **authored** record contents (condition (iv)); **never harvested from a build run.** Mismatch diagnosis = the harness diffs the two artifacts it holds.

The ordered per-record digest list is not restored in any form.

---

## §4 — External session editing (Decision 5 + the operator's edit-surface ruling)

**The boundary (operator-ruled, m-1-re-confirmed as BR-INV′, strictly stronger):** the external-edit surface is **m-9's session journal ONLY**. m-10's settlement store (receipt rows + resume snapshot) is **effectively uneditable**; recovery from an unresumable state involving that store is a **post-product hardening carry** (backlog), not MVP. The edit surface never reaches the m-1 governed store.

**§4a — The honest MVP detection claim (VP1-F2 folded — the "already total" claim is withdrawn and replaced).** There is **NO carrier** for a cross-store content-identity comparison at MVP: the settlement manifest carries ids/terminals/`args_digest`, not content identity; m-9 cannot read m-10's receipt table; m-10 does not read the journal; `turn_open` carries no round-identity operand, and **none is added**. A journal payload edit is detectable only by: **(a)** the advisory checksum (naive edit ⇒ mismatch), **(b)** completeness/parse (structural edit), **(c)** presence-at-consume (`content_lost`). **A checksum-recomputing, present, outcome-consistent edit is UNDETECTABLE at MVP** — an accepted, documented limit. **Frank does not claim to distinguish a sanctioned repair from corruption**; that classification arrives with the Step-4 versioning/provenance carry.

**§4b — The composed disposition machine (VP1-F4 folded; total, per `1f8ec7b6…` + `4d494778…`):** structural/completeness failure, missing referenced content, or an unresolvable reference ⇒ **`degraded` / `re_derive`**. A present, well-formed advisory-checksum-mismatch ⇒ **`resumable`** with a **local, in-memory** trust label (`untrusted-content` for provider/tool kinds, `checksum_mismatch` for input kinds) applied to model context only. An undetected edit resumes as if clean. m-10 consumes only the carried `{resumable, degraded}`; it dispositions, it does not detect. **There is NO universal "edited ⇒ degraded" rule.**

**§4c — Trust and honesty limits.** An edited byte **never** silently inherits the original's trust: edited provider/tool content is never presented as the provider's/tool's actual output (the fabrication class m-9's own survey rejects). The trust label is **in-memory only — `edited_since_write` is RETRACTED**: nothing is persisted, nothing rides any wire, no E3-visible edit state exists. Consequently (m-3, pair-approved): E3 may assert only `{resumable, degraded}` + direct-prefix equality/divergence, **never "edited"** — and **no exit claim may assert "frank labeled edited content untrusted"**, because no evaluator can observe the label.

**§4d — `receipt_conflict` (VP1-F9 exact wording):** this amendment **relaxes no `receipt_conflict` rule**; it stays frozen exactly as co-signed. A historical edit emits no second live receipt, so that rule was never the edited-session mechanism.

---

## §5 — `context_digest` dropped, claim narrowed (Decision 5) — NON-SEVERABLE

**Superseded:** `context_digest` as a member of `resume_prefix_expectation`. **REMOVED**, as a **risk acceptance, not a refuted analysis** — m-3's false-PASS finding stands; the operator overrules on cost-versus-likelihood, and zero of the eight surveyed harnesses verify this because none claims it.

**Bound non-severably: the Durability leg's claim NARROWS** to *the frozen record returned intact* — never *the model resumed the same conversation*. The §7 row wording and the final fixture oracle must carry the narrowed claim. Additionally (m-3): a Route-2 `fail` **at exit** is unambiguously a durability defect, because exit fixtures are never edited; the edit/defect conflation exists only in production runtime, which this section's narrowing covers.

---

## §6 — New §7 row: successor legitimacy (Decision 7) — SEVEN legs, TWELVE records (VP1-F6/F7 folded)

The successor exit suite is stated directly: **seven legs, twelve fixture records.** (The prior six-declared/six-listed state was consistent; this row is the seventh leg, and its arrival changes the counts — no historical defect is revived.)

- **Typed predicate:** `successor_admitted_at_current_epoch_under_valid_lease` v1 — **three independently-locatable observations**: `positive` + `neg.STALE_EPOCH` + `neg.WRONG_LEASE`, both negatives mandatory (they exercise distinct frozen reject mechanisms: assign-gate tuple-mismatch vs lease-not-held). `pass` iff the positive is admitted-and-proceeds AND both negatives refuse fail-closed with zero successor work; refusal alone never passes. `WRONG_LEASE` carries the joined **m-9 writer-fence observable** (`d38cd3c3…`): the replacement acquires the per-run `flock`; a disposed-but-live predecessor's would-block attempt is the observed refusal.
- **Records: the two negatives only.** The **positive arm rides the regenerated `xit-dur-1` fixture** under a **binary, testable precondition** (m-10-sharpened): the fixture must **expose the assign-gate-open-at-the-current-tuple observation** (rev16 §4:55/§6:130) — not merely "a resume happened." Precondition failure = **hard stop + a separately owner-reviewed 13-record successor** — never a silent cardinality change. An unclean `xit-dur-1` makes the positive arm **`unknown`**, never a spurious fencing `fail`.
- **Weights:** each refused record = `sample_weight {governed_turns: 0, tool_calls: 0}` (an admission-gate event; zero work by construction). l4 recomputed its ten filed weights from its own relay: the aggregate stays **exactly 30 governed turns / 100 tool calls** at twelve records. The formal re-balance lands in the fresh lane-4 plan.

---

## §7 — The two stale stage-4 loci (Decision 8) — carried HERE; §5's risk acceptance INHERITED (VP1-F8 folded)

Carried in-record because `2026-07-19-mvp-full-worker.md` is lock row 45 (m-9 caught this; in-place voids the lock). **Superseded and now named:** `:88` (crash semantics: *"nothing durable to reload … no reload boundary"*) and `:155` (the §10 no-second-truth fixture item) — superseded by r17 §5's narrowing as re-scoped by §1 of this amendment.

**What the retained checks prove — stated at §5's grain, no more:** last-record completeness and the per-run writer fence prove **record completeness and single-writer identity only.** The old absence-based stale-summary safety claim is **superseded, not re-secured**: under §5's risk acceptance, no MVP mechanism proves summary freshness or semantic context equality, and this amendment does not claim otherwise.

---

## §8 — What this amendment does NOT do, and what follows it

Changes no byte of any superseded file · edits no lock constituent · moves no lock hash · **relaxes no `receipt_conflict` rule** · adds no carrier, wire member, or E3-visible edit state · authors no fixture, manifest or expected value · claims no tamper-evidence · issues no PLAN/T4 token · touches no `frank/` path · permits no external use. **H-12 hard-blocks external/untrusted/multi-tenant use** independently of everything here.

**Re-lock:** the lane-4 re-lock exact-hash-binds **three inputs** — the interface lock, this amendment, and the frozen fixture manifest.

**Lane-4 propagation:** the plan (`60daac08…`) is stale by decision and is not patched here. After ratification master authors a **fresh plan revision** for VP review; the resume binds that approved hash. The revision must clear: the ten→twelve count and the "already-frozen / not up for redesign" framing (`:12`, `:22`) · the `xit-dur` inventory + `resume_prefix_expectation` schema (`:73-83`, `:79`) · the owner-fidelity matrix (`:88`) · the sequence/status text (`:111`, `:147`, `:160`, `:164`) **and the Status line `:3`** ("the ten-row structure … preserved") · the 30/100 re-balance with the weight-0 records · the `xit-dur-1` regeneration exposing the assign-gate observation (§6's precondition) · and the spec note that an edited-session fixture, if ever authored, asserts a **disposition, never a prefix match**.
