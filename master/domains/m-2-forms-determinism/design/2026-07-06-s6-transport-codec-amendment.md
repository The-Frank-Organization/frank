# v3 Forms & Determinism — s6 TRANSPORT-CODEC AMENDMENT (review-driven; amends c1 §2/§4/§12; no silent re-design)

**DESIGN_DOC_ID:** s6-amend-m-2-transport-codec
**Amends (does NOT reopen):** `c1-design-m-2-forms-determinism` — §2 (envelope), §4 (FieldSpec/recipient_picker), §12 (consumer contract). Folded into the locked doc only on VP co-sign (the c6 / s2-amend path).
**Basis:** the `s6-design` dispatch (`master/relays/s6-design/DESIGN-orchestrator-planner-20260706-180315.md`); `TRANSPORT-FINDINGS-2026-07-06.md` F2/F6/F7/F12/F13; the operator's in-Step-1 transport-fix ruling (2026-07-06). Baseline `frank/` `main @ 7e5c527` (tag `s5-close`); codec surfaces mapped at this seat (Explore pass).
**Owner:** m-2 (Forms & Determinism) · pair-Planner authored · for **m-2.implementer DESIGN-REVIEW** (the standing D5 recast: planner designs, implementer grills).
**Tier:** medium · **Evidence:** E1 (design against read code; no build in this phase).

---

## 0. Scope + boundary

**m-2's s6 charter (per dispatch § m-2):** the **single interpretation** — ONE canonical envelope codec consumed by render, typed-validate, the lineage gate, and delivery projection — killing the render↔validate↔gate↔deliver divergence **class** (not each instance). My five findings: **F6** (silent recipient drop), **F7** (CC deadlock), **F2** (dead-edge parent offer), **F13** (validator rejects registry-advertised tokens), **F12** (waiver-flag typing).

**This amendment does NOT touch** (named to hold the seam lines): F3 bounce-detail parity (m-7); F4/F5/F11 liveness — anchor model / digest lease / livelock (m-1 + m-7); **the parenting-model FORK decision** (m-1's grilled decision-packet — I design as a *consumer* of whichever branch wins, §3); F17 waiver **scoping/retraction record class** (m-1 — I seam my F12 typing to it, §4); F9 intake identity (m-7); F14 store-lock invariant (m-1) + runtime (m-7). No mechanism change to `{accepted, rejected, held}`, the `submit`/`project`/`read` surface, channel-stamped FROM, or I-PH. No Step-2 observe pre-work. **No c1 reopen** — a scoped, review-driven amendment.

---

## 1. Root cause in the m-2 layer

The global root — *"one envelope, many judges, no reconciliation protocol"* — lands in my domain as **four independent decoders of the recipient list**, mapped at `7e5c527`:

| Judge | Site | How it reads TO/CC | Form it demands |
|---|---|---|---|
| **Render** | `render.go:65-87,103-111` | `recipient_picker` is **unimplemented** → falls to `baseOptions` → `address_list` has no `enum_set` → **nil (0 candidates offered)** | — (offers nothing) |
| **Validator** | `canonical.go:19-61` `ParseTyped` | `json.Unmarshal` + exact `CanonicalMarshal` round-trip | **canonical JSON** `["a","b"]` |
| **Lineage gate** | `lineage.go:282-292,432-439` | `strings.Split(raw, ",")` | **raw comma-joined** `a,b` |
| **Delivery** | `projections.go:122-167` | `Envelope.To` verbatim **+** header `ParseTyped`; on any header parse-fail **returns only `Envelope.To`** (silent) | **canonical JSON**, else silent-drop |

Plus `Envelope.To` (`record.go:16-25`) is **one verbatim mailbox key**. So a protocol-legal multi-TO relay: canonical-JSON headers pass validate+delivery but the lineage gate comma-splits them into one non-matching token (**F7**); raw-comma headers pass the gate but fail validate; and a header parse-fail silently delivers to `Envelope.To` alone or to no one (**F6**). Advisory-render / authoritative-revalidate is *safe but not live*.

**Thesis (kill the class):** exactly **one codec** — one canonical wire form for a recipient list and **one decode function** that render, validate, the lineage gate, and delivery all call. With one decoder there is one judge; F2/F6/F7/F13 dissolve as instances, and F12 is a small typing fix.

---

## 2. The amendment — ONE canonical envelope codec

### 2.1 Canonical wire form (the single decode)
The recipient list's canonical on-wire form is the **typed `address_list`** already declared in §4 — a canonical-JSON array of minted addresses (`["m-2.planner","operator"]`), the form the validator + delivery already accept. A **single exported codec** `DecodeAddressList(raw) → ([]addr, error)` / `EncodeAddressList([]addr) → raw` becomes the *only* path any judge uses to read/write TO/CC. **The raw-comma-split is deleted** (`lineage.go:286,433`): the lineage gate's `addressedTo`/`addressedInHeader` decode via the codec, not `strings.Split(",")`. This is a **codec-unification**, not a grammar change — `address_list` (§4) and `recipient_picker` (§4) are already locked; the amendment makes their *interpretation* single-sourced.

### 2.2 All four judges read the codec
- **Validator** — already the canonical source; keep `ParseTyped` as the codec's decode entry.
- **Lineage gate** — `addressedTo`/`addressedInHeader`/`checkReviewerVisibility` (`lineage.go:170-184,282-292,432-439`) call `DecodeAddressList` on the header, membership-test the decoded list. **Kills F7**: a canonical-JSON CC now decodes to the same `[reviewer]` the validator sees, so **one encoding satisfies both** the typing gate and the reviewer-visibility gate — the deadlock class is gone.
- **Delivery** — `DeliveryRecipients` (`projections.go:122-142`) reads the decoded canonical list.
- **Render** — §2.5.

### 2.3 No silent drop; the validated-at-submit guarantee
A recipient list that reaches delivery has **already passed `ParseTyped` at submit** (the c1 pre-append form-validation, §4 step 4). So delivery must **never re-parse-and-silently-drop**: `DeliveryRecipients`' `if !ok { return recipients }` (`projections.go:135-136`) is removed — an unparseable TO/CC at delivery is an **engine invariant violation** (it should have bounced at submit), surfaced as an internal fault, never a silent loss. At the submit gate, an ill-formed recipient list bounces **`Field:Class`** (I-PH preserved), naming the offending field — never silently accepted-then-dropped. **Kills F6.**

### 2.4 the canonical Header list is the single recipient truth for EVERY projection (rev1 — implementer blocker 2)
`Envelope.To`'s single-verbatim-key role (`record.go:20`) is the second F6 encoding. **Resolution (tightened rev1):** the **canonical decoded Header `TO`/`CC` `address_list` is the single source of recipient truth for ALL projections** — the rendered relay markdown header (`projections.go:110-113`), the `INDEX.md` row, the lineage gate, **and** the mailbox delivery intents. None of these may read `Envelope.To` as the recipient set. `Envelope.To` **may** remain an *optional* primary-recipient compat/routing projection **iff m-7 needs it** (its stamping is m-7's commit loop — the **m-7 seam**; m-2 defines only the projection rule), but it is **never the sole display/index representation of recipient truth**. Rationale: the rev0 "`Envelope.To` = primary projection" left the render/index paths printing one recipient while delivery reached all — recreating the exact "one path sees less than another" divergence the codec exists to kill. The fixture (§7.1) proves the full multi-TO+CC set is preserved identically across markdown, index, mailbox intents, and the reviewer-visibility gate.

### 2.5 `recipient_picker` render (the render judge must offer only deliverable recipients)
Implement `recipient_picker` in `render.go` symmetrically to `parent_picker`: offer candidates = the **minted address space** (m-1's addressing graph, the §12 `recipient_picker` contract), and mark the field **`DigestExempt: true`** as `parent_picker` already is (`render.go:65-87`) — an address-candidate refresh must not churn the form digest (the m-2 render contribution to the F5 churn class; the digest *lease* itself is m-7's). This closes the render judge's "0 candidates" gap so the form offers exactly the recipients validate/gate/delivery will accept. **Kills the render half of F6** (and removes an F5 churn contributor).

### 2.6 the THREE authorization layers for record_kind (F13) — rev1 (implementer blocker 1)
rev0 conflated *membership* with *authorization* and would have accepted `record_kind=genesis` from any seat — which **contradicts my own s5 ruling** (s5-escalations M-3(e): `genesis` **removed from every `seat_scope`** — store-init machinery, kept in the named enum for compat only, `registry.json:84` vs `:125`). Corrected: record_kind acceptance is **three distinct layers**, each single-sourced, no judge duplicating another:
1. **Named-enum membership** — the token is known to the registry `enum_set` (`validate.go:54-61`, already single-sourced ✓).
2. **Authorization (seat_scope / form offer)** — the token is offered to *this submitting seat* (`reg.Validate` seat-scope; the fill-time-authority layer — a forbidden kind is absent from the seat's form). **`genesis` is in no seat_scope ⇒ a public submit of it is rejected here**, correctly.
3. **Per-kind required-field validation** — extra checks for *authorized* kinds (owed_item→owner/source/…, owed_disposition→disposes_owed, §s4-wire).

**Amendment:** `submit.go:143 validateRecordKind` **stops being a second membership judge** — membership (layer 1) + seat-scope authorization (layer 2) are enforced by `reg.Validate` *first* (`submit.go:47-64`); `validateRecordKind` runs only *after* and keeps **only** layer 3 (per-kind required fields). Its `default → "unknown record_kind"` (`:177`) — which today rejects authorized-and-offered kinds like `gate_resolution`/`disposition`/`diagnostics`/`config_change` — is removed for kinds `reg.Validate` already authorized (they pass with no extra requirements), while `genesis` still fails at layer 2. Same three-layer rule resolves the `gate_category: other` side (validate.go already accepts `other`; no second judge rejects it) **without widening `record_kind`**. **Kills F13**: every token the schema advertises **and offers to the seat** is accepted; unauthorized/unoffered kinds (`genesis`) are still rejected — by seat-scope, not by a hardcoded switch. *(Consistent with my s5 Q4/Q6 ruling: a validation-source unification, not an enum widening — no `:126` bump.)*

---

## 3. The three-way parent seam (F2) — consuming m-1's fork

**F2** (render offers dead-edge parent `s5-dispatch`; lineage rejects `parent-invalid-dead-edge`) is a render↔validate divergence on **PARENT_DISPATCH_ID**, and it is downstream of **m-1's parenting FORK** (the grilled decision-packet). I design my half **composable with either branch** (the way m-7 built its M-2 mechanics composable with any m-6 signal-set):

- **If the fork = conductor-computed PARENT** (the leading seed — engine stamps lineage server-side as it stamps FROM): `PARENT_DISPATCH_ID` becomes `owner: system` / `system_only`, so `render.go renderable()` returns false → **the parent field is not rendered at all** → the render-offers-a-dead-edge divergence **dissolves entirely** (exactly as ROLE/FROM dissolved `role_from_consistency_error`, c1 §2/§10a: no seat-fillable field → nothing to diverge; confusion-resistant, D5 residual). This is the strongest form and I recommend consuming it.
- **If the fork = widened/validated candidate set:** the render judge's `parent_picker` candidate source (`env.ParentCandidates`) must be the **same computation** the lineage gate validates against — one candidate source, so render can never offer a parent the gate rejects. m-2's codec makes render and gate read one parent-candidate source; m-1 owns what that set contains.

Either way the m-2 rule is invariant: **render offers a parent only if the lineage gate will accept it** (single candidate source), reaching its limit (no rendered parent) under conductor-computed PARENT. The FORK decision is m-1's; this section is the consumer contract, locked to neither branch.

---

## 4. F12 — `ORCH_REVIEW_WAIVER` typing (seam with m-1's F17)

At `7e5c527` `ORCH_REVIEW_WAIVER` is `owner: seat_scoped_enum`, `type: text`, `seat_scope: {operator: ["*"]}` — a **bare flag** whose only operator option is `"*"`; the lineage gate checks presence only (`lineage.go:179`, `!= ""`), so the waiver's **rationale had to ride SUBJECT/body** (F12). The waiver *design intent* is a **justified** escape, not a bare flag.

**Amendment (rev1 — implementer blocker 3; the header-fallback branch dropped as not-expressible).** rev0 offered a fallback "operator-scoped `free_text` header value" — but that is **exactly the free-text-cannot-be-seat-scoped limit I myself established** (s4-wire: `seat_scope` is enum-only; s5-escalations M-1: the render context cannot seat-gate a free-text field). The current FieldSpec model (`fieldspec.go:15-20`, `registry.json:134`) has **no way** to express "operator-only arbitrary `free_text` header with submit-time enforcement" without new machinery (a grammar extension — out of s6 scope). So F12 **binds to m-1's F17 waiver-record carrier** (single branch):
- F17 makes the waiver an **operator-only record class** with scope + retraction (m-1, `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:43-51`); operator-only comes from **record-authorship authority** (channel-stamped FROM), *not* from free-text-field seat-scoping — sidestepping the limit entirely.
- **m-2 defines the exact field rows on that record:** a `rationale` (`free_text`) row carrying the justification, plus the typed rows F17's scope/retraction semantics need (e.g. a `waiver_scope` and a `retracts` `id_ref` — shaped to m-1's record-class contract). The bare `ORCH_REVIEW_WAIVER: "*"` header shape is **retired**; the lineage reviewer-visibility gate reads the waiver *record* (its scope + not-retracted state), not a presence-only header flag.

m-2 owns the rationale/scope/retraction **field typings**; m-1 owns the waiver **record class + retraction semantics**. Neither locks in isolation — this is the F12↔F17 seam. *(Observation for the F17 pass: the retired header's `lineage_role: routing_ref` was mis-typed for a waiver; it disappears with the header.)*

---

## 5. Per-finding disposition

| Finding | Sev | Disposition | Mechanism (this doc) |
|---|---|---|---|
| **F6** silent recipient drop | 1 | **design amendment** | §2.1–2.4 — single codec; delete raw-comma-split + silent-drop; `Envelope.To` = projection |
| **F7** CC deadlock | 1 | **design amendment** | §2.2 — lineage gate decodes via the codec ⇒ one encoding satisfies typing + visibility |
| **F2** dead-edge parent offer | 2 | **design amendment (m-1 fork seam)** | §3 — render offers a parent only if the gate accepts; dissolves under conductor-computed PARENT |
| **F13** validator rejects advertised tokens | 2 | **design amendment** | §2.6 (rev1) — three layers: membership (`enum_set`) / authorization (`seat_scope`) / per-kind required-fields; `validateRecordKind` keeps only layer 3; `genesis` stays rejected by seat-scope |
| **F12** waiver-flag typing | 3 | **design amendment (bound to m-1 F17 record)** | §4 (rev1) — waiver becomes m-1's operator-only F17 record; m-2 types the `rationale`/scope/`retracts` rows; the `"*"` header is retired (free-text-header branch was inexpressible) |

No wontfix; no silent drop. The render-side `recipient_picker` implementation (§2.5) is the one **build-slice fix inside the amendment** (it realizes a locked §4 fill_constraint that was never implemented — not a new contract).

---

## 6. Constraints preserved (restated per dispatch)

Byte-exact `{accepted, rejected, held}` — untouched (the codec changes *interpretation*, not the outcome enum). Seat surface exactly `submit`/`project`/`read` — untouched (the codec is internal to the conductor's judges). Channel-stamped FROM — untouched (and the parent seam *extends* the same server-stamp discipline to PARENT under the conductor-computed branch). **I-PH on every bounce** — the codec's parse failures surface `Field:Class` only, never a path (§2.3). Claim ceiling = **tool-mediated confusion-resistance** — no wrap / by-construction creep (the codec is a reconciliation of judges, not a new trust claim); D5 residual accepted. Crash-atomicity, FROM-stamping, I-PH — not reopened.

---

## 7. Fixture obligations (owed to the s6 build slice)

1. **Codec round-trip + full-set-in-every-projection (F6/F7; rev1 blocker 2):** a protocol-legal **multi-TO + multi-CC** relay (canonical `address_list`) preserves the **identical full recipient set across ALL projections** — the rendered relay markdown header, the `INDEX.md` row, the mailbox delivery intents, **and** the lineage reviewer-visibility gate (which matches every CC member). None reads `Envelope.To` as the recipient set. Negative: a header parse-fail **bounces at submit** (`Field:Class`), and **no** code path delivers-then-silently-drops.
2. **One-encoding-satisfies-both (F7):** the exact archive relays `relay-c13dc32f` / `relay-1f99aadf` (intakes 000020/000021) now pass both the typing gate and the visibility gate on the single canonical encoding.
3. **Three-layer record_kind authorization (F13; rev1 blocker 1):** every `record_kind`/`gate_category` token that is both in the registry `enum_set` **and offered to the submitting seat** is accepted (incl. `gate_resolution`, `disposition`, `diagnostics`, `config_change`, owed kinds, `other`) subject to its per-kind required fields. **Negative leg (required):** a public seat submitting `record_kind=genesis` is **rejected at the seat-scope layer** (`genesis` in the enum for compat, in no `seat_scope`); a genuinely-unknown token bounces at the membership layer.
4. **Parent no-divergence (F2):** render offers no parent the lineage gate rejects; under the conductor-computed branch, **`PARENT_DISPATCH_ID` is absent from every rendered form** (dissolution fixture).
5. **Waiver-as-F17-record (F12; rev1 blocker 3):** the waiver is m-1's operator-only F17 record; m-2's `rationale` (`free_text`)/scope/`retracts` field rows carry the justification first-class (not SUBJECT/body); the reviewer-visibility gate reads the waiver *record* (scope + not-retracted), not the retired `"*"` header. Fixture composes with m-1's F17 record-class + retraction legs.
6. **`recipient_picker` render + digest-exempt (§2.5):** TO/CC render minted-address candidates and are `DigestExempt` (no digest churn on candidate refresh).

---

## 8. Boundary — what m-2 does NOT claim

The FORK decision (m-1); F17 scoping/retraction record class (m-1); `Envelope.To` **stamping** (m-7 commit loop); the digest **lease** (m-7, F5); F4/F11 anchor model + livelock (m-1/m-7); F3 bounce-detail parity (m-7); F9/F14/F15/F16 (m-7/m-1). Where my codec touches these (the `Envelope.To` projection rule, the parent-candidate source, the waiver rationale carrier), I define the **m-2 half** and name the seat that owns the other half. **[VP-W2 analog]** the parent seam cites both halves (render/validate = m-2; stamp/anchor = m-7/m-1).

---

## 9. Review routing

This doc → **m-2.implementer** for `PHASE: DESIGN-REVIEW` (the adversarial grill), same `DESIGN_DOC_ID`, verdict `{approve | must-revise | reject-narrow | human-decision-required}`. On approve → SITREP to master.orchestrator-planner (design complete, held for integration); master integrates the three amendment docs + the seam, VP co-signs, then the s6 build slice dispatches. **The m-1 fork does not lock without the operator grill** (dispatch [VP-W3]); my §3 is written to consume whichever branch that grill locks — this doc can approve and hold without waiting on the fork.

---

## 10. rev1 fold-log (m-2.implementer DESIGN-REVIEW `184218`, verdict must-revise; all three blockers verified correct + folded)

- **Blocker 1 (F13 over-broad) → §2.6/§5/§7.3 rewritten.** rev0 would have accepted `record_kind=genesis` from any seat — contradicting my **own** s5-escalations M-3(e) ruling (`genesis` removed from every `seat_scope`, kept in the enum for compat only). Folded the three-layer split — membership (`enum_set`) / authorization (`seat_scope`) / per-kind required-fields — with `validateRecordKind` reduced to layer 3 after `reg.Validate` enforces layers 1–2; added the negative fixture (public `genesis` rejected by seat-scope). This *is* my fill-time-authority thesis; rev0 conflated membership with authorization.
- **Blocker 2 (`Envelope.To` projection insufficient) → §2.4/§7.1 rewritten.** rev0 left render markdown + `INDEX.md` printing `Envelope.To` while delivery used the full list — the same "one path sees less" divergence. Folded: the canonical decoded Header `TO`/`CC` is the recipient truth for **every** projection (markdown, index, lineage, delivery); `Envelope.To` is an optional m-7 compat projection only; fixture proves the full set across all four surfaces.
- **Blocker 3 (F12 header fallback inexpressible) → §4/§5/§7.5 rewritten.** rev0's "operator-scoped free_text header" is the free-text-cannot-be-seat-scoped limit **I myself established** (s4-wire, s5-M-1). Folded: F12 binds solely to m-1's F17 operator-only waiver **record** (operator-only via record-authorship, not field seat-scoping); m-2 types the `rationale`/scope/`retracts` rows; the `"*"` header is retired.
- **Non-blocking confirmations accepted as-is:** F6/F7 codec thesis directionally sound; parent seam composable with either fork branch (`render.go:55-63` hides `system_only`); no c1 reopen under this DESIGN_DOC_ID. **Scope unchanged:** still interpretation-only; byte-exact `{accepted,rejected,held}`, the 3-verb surface, I-PH, and the confusion-resistance ceiling untouched; no c1 fold until VP co-sign.

---

## 11. B-2 addendum — the typed BOOT form (s6-design addendum `200259`; one bounded leg, folds into the same integration package)

**Basis + honest premise (per the addendum's own audit):** the observed boot **bounces** were all `form_digest:re-render` (F5) + digest/parent compounds — **already killed** by A-1 (stable-schema digest) + branch A (conductor-computed PARENT). So B-2 does **not** buy bounce-proofing; it makes the v2.8.8 boot convention (`DISPATCH_ID <run>-boot-<seat>`, report-only ack) a **first-class typed form** with lifecycle visibility. m-2 owns the FORM + the token decision; **B-1 (m-7)** owns the lifecycle state that gates it; **B-3 (m-1)** owns the identity-activation semantics of the accepted boot record. Bounded — no hidden fork, no lock conflict; designed, not escalated.

### 11.1 The minimal honest required set
The boot record's identity is its **channel-stamped `FROM`/`ROLE`** (system-owned, one-FROM rule, c1 §2/§3) — **not** a seat-filled "identity claim" (a seat cannot claim an identity it does not hold; the stamp *is* the claim, and B-3 makes the accepted boot the activation edge). So the seat-filled required set is the **smallest honest two fields**:
- `charter_loaded` — `bool` (`enum_set: bool`), `owner: agent_enum_pick`, `gate_referenceable: false` — the seat asserts it loaded its role charter/skill. **Honestly `self_reported`** (Step-1 has no observe layer to verify it — CQ-1/AC17; labeled, never trusted-as-observed). Not a gate input.
- `dispatch_status` — `enum {read, awaiting}`, `owner: agent_enum_pick`, `gate_referenceable: false` — did the seat read its boot dispatch or await one.

Plus the system envelope (FROM/ROLE/DISPATCH_ID/timestamp stamped as for any record). **No** exotic seat-scoped enums; **no** parent semantics (branch A — no PARENT field renders); **no** `gate_category` (boot grants nothing).

### 11.2 The token question — RULED: `PHASE: SITREP`, lifecycle-gated; NO new `BOOT` phase atom
Boot is semantically a **report-only acknowledgment** = the existing `SITREP` phase. **Ruling: reuse `PHASE: SITREP`.** Explicit vocabulary impact (as the dispatch requires):
- **A new `BOOT` phase token would be a shared-vocabulary change** — every phase consumer must then handle it: the lineage phase-conditionals, `authority_consistent` (phase↔authority — a new BOOT→report-only mapping), and m-6's meeting-lane `phase_in` routing. For a record that *is* a report-only SITREP, that buys nothing. **Rejected.**
- **`SITREP` reuse** touches no shared vocabulary — inherits the existing SITREP↔report-only mapping; the PHASE enum is unchanged. **Chosen.**
- **No `record_kind` boot member** — record_kind is a system-owned enum switched exhaustively (`validateRecordKind`); adding a member is **MAJOR** per my `:126` contract (the Q6/F13 line). Boot must not ride record_kind.
- **No §J2 / `gate_category` impact** — boot is not a gate.
- **Gated by lifecycle state, not a self-marker:** the renderer serves the **BOOT form** (SITREP + the two boot fields) as a **distinct rendered form when the seat is pre-active** (`bound` — the m-7 B-1 lifecycle state supplied to the render context, the seam) — so the fields are simply *required within the boot form*, needing **no new §5 predicate atom** and no self-asserted marker. Once `active`, ordinary forms render. m-7's active-transition detects "the seat's first accepted SITREP carrying the boot required-set" (B-1 rules whether that derivation suffices or adds a **system-derived** `boot_ack` marker — system-stamped/confusion-resistant, never seat-asserted, and still not a new phase/record_kind). **Net vocabulary impact: additive-MINOR at most** (the two new fields + one `dispatch_status` named-enum; one optional marker field iff B-1 needs it) — the shared PHASE / record_kind / §J2 / §5-atom vocabularies are **untouched**.

### 11.3 Un-bounceable by construction (post-A-1, branch A)
A seat filling exactly what the boot form shows submits **accepted** without bounce: **A-1** stable-schema digest ⇒ no `form_digest:re-render`; **branch A** computed PARENT ⇒ no seat-parent to bounce (`parent-invalid-dead-edge` cannot arise); **the codec** (§2/§2.6) ⇒ the two boot fields validate against the single registry `enum_set` (no render↔validate divergence). The observed boot bounces (all F5/parent class) are structurally absent — the required-set *is* the rendered set.

### 11.4 Constraints + seams
No new seat verb (the roster is m-7's `project`-view / operator surface — B-1). The boot form **grants nothing** — it sequences a seat's first action; it is not an authority gate (the addendum constraint). I-PH on any boot bounce (`Field:Class`). `charter_loaded` is `self_reported` (no Step-2 observe pre-work; honestly labeled). **Seams:** B-1 (m-7) — the `minted→bound→active` lifecycle state (the render gate) + roster + out-of-order semantics; B-3 (m-1) — the accepted boot record as the identity-activation edge, session-restart = re-`bound`-not-re-boot, and the D4/D5 line (activation is confusion-resistant liveness bookkeeping, **not** an identity-strength upgrade — the channel stamp already carries identity). B-2 owns only the form schema + the token decision.

### 11.5 B-2 fixture obligations
- The BOOT form renders (SITREP + `charter_loaded` + `dispatch_status` required) exactly when the seat is pre-active; a pre-active seat filling the shown set submits **accepted** with **no** re-render / parent / enum bounce (the un-bounceable-by-construction fixture, post-A-1 / branch-A).
- **Shared-vocab-untouched negative:** no new PHASE / record_kind / gate_category / §5-atom token appears in the registry as a result of B-2.
- `charter_loaded` is labeled `self_reported` (not observed) and is not a gate input (no predicate keys on it).
