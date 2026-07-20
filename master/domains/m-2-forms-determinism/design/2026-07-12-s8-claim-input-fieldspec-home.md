# s8 CLAIM-INPUT — the executable-claim DECLARATION FieldSpec home (m-2 leg of `s8-claim-input-amendment`)

**DESIGN_DOC_ID:** s8-claim-input-m2-home
**Owner:** m-2 (Forms & Determinism). One leg of the three-owner amendment `s8-claim-input-amendment` — **m-3 owns the semantics (first-among-equals); m-7 confirms the composition seam; my bytes FINALIZE against m-3's returned semantics.**
**Basis:** the amendment DESIGN dispatch (`s8-claim-input-amendment/DESIGN-orchestrator-planner-20260712-004511`); the T9 production RED (reader-with-no-writer: `observe.Env.Evaluate` nil, no seat-declared claim row in the activation matrix); c1 §5/§9; the T2-built `s8-fieldspec-v6` registry.
**Tier:** medium · **Evidence:** E1 (design against read registry + the locked matrix) · **GRILL_REQUIRED:** no (the gap is executable-RED-characterized; the decision record is the amendment thread).

This defines the **m-2 grammar home only** — the declaration row's shape, owner class, gating, R2 posture, and the `v6→v7` transition. It **consumes** m-3's declaration semantics (the triple's fields, cardinality, `EVIDENCE_TARGET` coupling, param-validation locus) at the named finalization points (§6); it does not author m-3 content.

---

## 1. The gap + the m-2 leg
Observe can **read** executable claims (`observe.Env.Evaluate`) but no FieldSpec row lets a lane **declare** them — production `Evaluate` is nil, T9's `TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate` is RED on the real path. m-2 supplies the **declaration row's home**; it lands as the **second governed fieldspec transition `v6→v7`** over my successor map.

---

## 2. The home (m-2-owned grammar — stable across m-3's returns)
Grounded in the live registry (`s8-fieldspec-v6` base):
- **Field `executable_claims`** — the natural seat-declared counterpart to the system-owned `executable_claim_results` (`registry.json:151`). (Exact id finalizes vs m-3, but the *home* is this shape.)
- **`layer: header`, `type: row_array`** — precedent: `routing_assignments:174`, `choices:182`, `executable_claim_results:151`.
- **`owner`: a seat-authored CARRIER — `agent_enum_pick` / `free_text` row_array** (as `choices:182` / `SCOPE_DIFF:123` actually are), meaning *seat-authored row data*. **Correction (rev1 — implementer Blocker 1, verified):** a top-level `seat_scoped_enum`/`seat_allowed_values` does **NOT** constrain the nested `check_id` — `row_array` columns have **no grammar enforcement** today (`ParseTyped(row_array)` proves only canonical `[]map[string]string` shape, `canonical.go:20-30`; the enum-token check fires only for `type == "enum"` and the seat-scope check runs on the whole field, never columns, `validate.go:54-62`; consistent with my own s7a R2-COLGRAIN finding that columns are not value-validated). So nested `check_id`/`params` validation is **DELEGATED to the m-3-confirmed fill/observe validator seam** — never implied by the carrier's owner/fill. The exact owner/fill pair **finalizes with m-3's validation locus** (leg b), NOT independently fixed by the routing precedent. *(A real nested-column grammar is REJECTED as the fix: that is speculative machinery, against the confusion-firewall "build the seam, not the mechanism" rail — delegate to the validator seam.)*
- **`gate_referenceable: false` — R2 (leg c):** the WHOLE row is non-gate-referenceable; **NO `gate_referenceable_columns`** (stricter than `routing_assignments`' one gate-ref column). Selection is lane **INTENT**; the gate-relevant value is the system-computed `executable_claim_results`, never the declaration. (R2 holds regardless of the owner/fill choice above.)
- **Gating: `visible_when {all_of:[{layer_present: observe}]}`** (render-eligible when observe active; mirrors `choices`/`routing_assignments`). **NO unconditional `required_when`** — the field is **optional**: a record that declares nothing gets the honest **no-vantage degrade** (`Evaluate: nil` preserved, r5 point 3), never a bounce. **But optionality is NOT the same as additive-open compatibility** — see the §3 correction: a *present* declaration is mechanically consumed and its ignore-by-an-incompatible-reader changes acceptance. A **conditional** `required_when` coupled to `EVIDENCE_TARGET` **finalizes vs m-3 leg (a)** — and any requiredness add re-runs §9 (breaking/MAJOR, `v3-form-schema-design.md:126`).
- **Columns (draft; finalize vs m-3 leg a):** `claim_ref`, `check_id`, `params` — enumerated in the annotation like `routing_assignments`. Cardinality = row_array (several claims per record); m-3 confirms.
- **The (h)-guard pairing (leg d):** INPUT `executable_claims` **seat-declared** / OUTPUT `executable_claim_results` (`:151`) **system-owned** — opposite sides of the s5-b (h) channel-keyed suppliability guard, by design.

---

## 3. The `v6→v7` transition — version class HELD for master reconciliation (rev1 — implementer Blocker 2 + m-3 F1)
**Correction: the rev0 "additive-MINOR / Rail-A-open" lock was FALSE.** Old-reader ignore-unknown is NOT semantically inert here — the row is **mechanically consumed at acceptance**:
- **Optional ABSENCE** (a `v7`-compatible reader sees no row) → the honest no-vantage degrade (`Evaluate: nil`). Fine — genuinely inert.
- **A PRESENT declaration encountered by an INCOMPATIBLE (`v6`) reader** → the `v6` reader ignores the row, falls to `Evaluate: nil`, and can **ACCEPT** a record a `v7` reader would have **REJECTED** (the declared check's observed-false result). Ignoring the field therefore **turns a would-be rejection into acceptance** — **Rail A CLOSED / fail-closed**, not additive-open.
- **Refusal proof (named):** a present `executable_claims` row **cannot be silently treated as absent by a reader that cannot interpret/execute it** — the reader-capability gate (m-7's A2 mechanism) must **refuse** before submit semantics can diverge.
- **The exact `v6→v7` compatibility / version class is HELD for master reconciliation** with the revised m-3 return (`s8-claim-input-m3` F1 — reached this independently) + m-7's reader-capability gate. m-2 does **NOT** finalize `v7` as MINOR against a parent "Rail-A additive" instruction now contradicted by both owner reviews. The marker stays `s8-fieldspec-v6 → s8-fieldspec-v7` (my successor map); only the **CLASS** (not MINOR) awaits reconciliation.
- **§9 re-run hook:** if the `EVIDENCE_TARGET` coupling later adds `required_when`, requiredness changes are independently breaking/MAJOR (`v3-form-schema-design.md:126`) — re-classify then.
- **Marker + machinery:** the T2-built §2.4 acceptance + FX-CFG-10 legs still run a second live lap (spine-proving). **A-1 stale-form** (a `v6`-rendered form bounces `re-render` post-`v7`) + **forward-only** (pre-`v7` records never re-validated) ride the standing machinery unchanged.

---

## 4. Byte-sites (finalized in the return, byte-exact against m-3)
1. **version marker** (`registry.json:2`): `s8-fieldspec-v6` → `s8-fieldspec-v7`.
2. **new field row** `executable_claims` inserted by the observe cluster (near `executable_claim_results:151`).
3. **[finalize vs m-3 leg b]** nested `check_id`/`params` validation lands at the **m-3-confirmed fill-time + authoritative observe-time validator seam**, keyed to the **check registry/schema source** (m-3's §6.1 catalog) — **NOT** a top-level FieldSpec `enum_set` (an `enum_set` types the whole `row_array`, never a nested column; the mechanism rev1 rejected). No registry byte-site claims nested-column validation.

*(The downstream test byte — `registry_test.go` expected version → `s8-fieldspec-v7` — rides the build handoff, as in T2. The successor SHA + transition `new_digest` are derived post-apply, not owner-supplied.)*

---

## 5. Constraints (binding this leg)
- **Tripwire — NO lock-pinned value moves:** not the genesis v5 SHA, not FX-CFG-7, not the capability exact-sets, **not the catalog byte-pin `943f07bb…`** (the s8 pair's pending-review artifact, explicitly not this thread's). Only the `v6→v7` marker + the new field row move.
- **Rails:** A = **CLOSED / fail-closed per this surface** (rev1 — a present declaration ignored by an incompatible reader changes acceptance; §3), the version class held for master reconciliation with m-7's reader-capability gate; B = **pass** — declaration is honest evidence-labeling / drift-prevention, claims stay confusion-graded (no adversarial machinery).
- **I-PH:** unknown `check_id` / non-conforming `params` / stale form ⇒ **typed rejects, no pivot** (m-3 states I-PH; the home's validation conforms — one terminal typed reject through the single pivot).
- **Scope:** the DECLARATION surface ONLY. **No adjudication-rung content** (that is design item 9, s9, separate).

---

## 6. m-3 finalization points (my bytes finalize against these — not guessed here)
- **leg (a):** exact columns + cardinality + the `EVIDENCE_TARGET` coupling → whether/how `executable_claims` becomes **conditionally** required.
- **leg (b):** the `check_id`/`params` validation locus (fill-time + authoritative observe-time) keyed to the **check registry/schema source** (m-3's §6.1 catalog — NOT a FieldSpec `enum_set`) + the typed reject for unknown `check_id` / non-conforming `params`.
- **legs (c)/(d):** R2 non-gate-referenceable (I assert), the (h)-guard INPUT/OUTPUT pairing (I assert, grounded `:151`) — m-3 confirms alignment.
- **leg (e) — the exact Rail-A split (rev2):** m-2 asserts only that **optional ABSENCE may degrade** (`Evaluate:nil`); a **PRESENT declaration is CLOSED/fail-closed** and its `v6→v7` version/capability class is **HELD for master reconciliation** (with m-7's reader-capability gate), NOT an m-2 additive/open assertion.

---

## 7. Review routing
→ m-2.implementer `PHASE: DESIGN-REVIEW` (same leg doc `s8-claim-input-m2-home`). On approve, the reviewed home returns TO master under sub-ID `s8-claim-input-m2`, **finalizing byte-exact against m-3's returned semantics**; the registry byte handoff to the s8 build follows the T2 owner-fidelity pattern. No c1 reopen; the amendment rides the ritual, not a re-design.

---

## 8. rev1 fold-log (m-2.implementer DESIGN-REVIEW `005100`, must-revise — both blockers verified correct + folded)
- **Blocker 1 (false nested-enforcement claim) → §2 owner corrected.** rev0 claimed `seat_scoped_enum`/`seat_allowed_values` constrains the nested `check_id`; the grammar provides no such thing — verified `canonical.go:20-30` (`ParseTyped(row_array)` proves only canonical shape) + `validate.go:54-62` (enum-token check is `type=="enum"` only; seat-scope runs on the whole field). Folded: the top-level owner is a **seat-authored carrier** (`agent_enum_pick`/`free_text`), nested validation **delegated to m-3's fill/observe validator seam**; the owner/fill finalizes with m-3's locus, not the routing precedent. Rejected the "invent a nested-column grammar" alternative as speculative machinery (confusion-firewall seam rail). R2 unchanged (top-level `gate_referenceable:false`).
- **Blocker 2 (false additive-open version class) → §3 reclassified + HELD.** rev0's Rail-A-additive/MINOR was wrong: a **present** declaration ignored by a `v6` reader falls to `Evaluate:nil` and can **accept** what a `v7` reader would **reject** — ignoring it changes acceptance ⇒ **Rail A CLOSED/fail-closed**, not open. Separated optional-absence (may degrade) from present-declaration-on-incompatible-reader (must fail closed via m-7's reader-capability gate); named the refusal proof. The `v6→v7` version class is **HELD for master reconciliation** with the revised m-3 return (`s8-claim-input-m3` F1, independently concurrent) + m-7's gate — m-2 does not finalize MINOR against a parent instruction now under owner review. §9 re-run hook retained for any later `required_when`.
- **Accepted unchanged (implementer-confirmed):** the `executable_claims`↔`executable_claim_results` (h)-guard pairing, `gate_referenceable:false`/R2, `visible_when {layer_present:observe}`, stale-form re-render, forward-only, Rail B, I-PH, the byte-site/tripwire scope, and the `s8-fieldspec-v6` build base.
- **Honest note:** both blockers are one over-claim from two angles — `row_array` columns have no grammar enforcement (B1), and a present declaration is mechanically consumed so ignoring it is not inert (B2). I asserted enforcement + benign additivity the code/acceptance-path don't provide — the recurring assert-before-verify. Verified both at source this pass (`canonical.go`, `validate.go`); the pair and m-3 independently caught them.
- **Return posture:** Blocker 1 is m-2-owned and folded (re-reviewable now); the Blocker-2 version class is routed to **master** for reconciliation (the parent "Rail-A additive" line is contradicted by both m-2 and m-3 reviews). The finalized m-2 home does not return to master until that reconciliation + m-3's semantics land.

**rev2 fold-log (m-2.implementer DESIGN-REVIEW `005700`, must-revise — CONSISTENCY-ONLY: both substantive folds accepted; two residual echoes swept).** The primary sections were corrected in rev1 but two summaries still carried the rejected model — my recurring echo-gap.
- **Echo 1 → §4 byte-site 3 rewritten.** It still listed a possible top-level "`check_id` enum_set (if `seat_scoped_enum` against a static check catalog)" — the exact mechanism rev1 rejected. Replaced: nested validation lands at the m-3-confirmed fill-time + observe-time validator seam keyed to the check registry/schema source; **no registry byte-site claims nested-column validation**.
- **Echo 2 → §6 leg (e) split + leg (b) de-enum'd.** leg (e) said "the Rail-A degrade (I assert)" (under-read the accepted split); rewritten to the exact two-case rule — optional ABSENCE may degrade / PRESENT declaration is CLOSED-fail-closed, version/capability class master-held. leg (b)'s "enum_set source (static vs dynamic)" also implied a nested enum → changed to "check registry/schema source (NOT a FieldSpec `enum_set`)".
- **Whole-doc grep sweep (done this pass, not deferred):** `enum_set`/`seat_scoped_enum`/`additive`/`MINOR`/`Rail-A additive` — every remaining occurrence is the correction, an explicit negation, historical fold-log, or the parent's contradicted instruction; none asserts the rejected mechanisms. No owner-confirmed decision changed; the version class stays master-held.

---

## 9. FINALIZED bytes (all three gates satisfied — grammar-approved home realized against m-3 §12 + m-7 capability)
Gate 1 (Rail-A ratified ABSENCE-OPEN/PRESENT-CLOSED, `…-020010` §1) · gate 2 (m-7 capability `{v5,v6,v7}` + marker preflight, `s8-claim-input-m7/…-232825`) · gate 3 (m-3 §12 r3-approved) all satisfied; the byte-exact delta is the finalization COORD `master/relays/s8-claim-input-amendment/COORD-planner-20260712-021000.md`. **Two sites** against the `s8-fieldspec-v6` build base: (1) version `s8-fieldspec-v6 → s8-fieldspec-v7`; (2) INSERT the `executable_claims` row — `owner: agent_enum_pick`, `type: row_array`, `gate_referenceable:false`, `fill_constraints: free_text`, `visible_when {layer_present:observe}`, **no `required_when`** (ABSENCE-OPEN); columns `(claim_ref[unique], check_id, params)` in the annotation, nested validation delegated to the m-3 registry-schema-aware fill/observe seam (no FieldSpec enum_set). PRESENT-CLOSED is carried by the `v6→v7` marker + m-7's capability ceiling, NOT a registry byte. Returned to master for the three-leg byte-grain fold + the bounded T9 grant; the s8 build applies on the grant (T2 pattern).
