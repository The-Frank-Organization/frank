# Stage-6 §5-C — the `relay.submit` `canonical_resource` shape (the m-2-authored cell for the §D-settlement amendment)

**DESIGN_DOC_ID:** `step3-relock-c-m2-submit-resource` · **rev1** — the current revision; history in §7 (authoritative; this marker moves with every revision).
**Routed by:** master's ruling (4), `step3-relock-dag-m10/RECONCILE-orchestrator-planner-20260722-230000.md` §(4), restated `…-20260723-001500.md` — *"bind a target identity derived from m-2's form schema rather than `∅`"*, pair-reviewed, feeding the bounded §D-settlement amendment.
**Basis:** the ratified stage-6 amendment `master/STEP-3-STAGE6-AMENDMENT.md` §5-C (rev12, SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` — verified this session) × my frozen stage-1 form-schema contract (rev5, `83d8e63e…`, byte-exact/unmoved).
**Scope:** the `relay.submit` cell ONLY. `relay.read` (`relay_id`) and `relay.project` (view token) are unchanged and not mine to restate; no other §5-C field moves; no sibling-owner byte moves; no FieldSpec registry byte; additive — no frozen m-2 byte is edited.

---

## §1 — The gap, exactly

§5-C requires `canonical_resource` for `relay.*` as "relay verb + target id". `relay.read` and `relay.project` both name a **pre-existing** target (`relay_id`; a view token). **`relay.submit` has none — it CREATES the record it names**, and the table admits no `∅` in the relay column. The cell is therefore inapplicable as written, and neither m-9 nor m-10 could fill it without inventing another domain's verb shape (correctly escalated).

**Why `∅` is NOT the faithful answer** (master invited the counter-argument; I do not take it): `∅` means *no such context exists*. For `relay.submit` a meaningful target context **does** exist and is **structurally guaranteed present** — the submission is made *against a specific rendered form* whose identity (`form_digest`) is in the frozen R-1 **required** member set, and it is *addressed into* a lane/recipient the submission carries. Encoding present context as absent would drop invocation evidence the descriptor exists to bind — the opposite of C's purpose (`local_invocation_matches_effect_descriptor` would have strictly less to match). `∅` is honest only where nothing is there (`bash`'s `canonical_resource`, the relay column's `workspace_root_id`/`cwd`).

## §2 — The shape (normative)

### §2.1 Definition

```
canonical_resource(relay.submit) := "relay.submit:" || submission_target_digest

submission_target_digest := SHA-256 (lowercase hex) over the JCS (RFC 8785) encoding of
the CLOSED submission-target object:

  { "form_digest":  <the submission's form_digest member>,   // string — REQUIRED, always present
    "dispatch_id":  <the submission's dispatch_id member>,    // string — present iff non-empty
    "to":           <the submission's to member>,             // string — present iff non-empty
    "cc":           [<address>, …],                           // ARRAY — see the cc rule below
    "cc_unparsed":  <the submission's cc member> }            // string — see the cc rule below
```

**The `cc` rule (added at review-r1 F1; total by construction).** The top-level `cc` member is a *canonical-JSON-string carrier* (frozen R-2: an array of address strings encoded as a string), so it is bound at the **value** grain, never as serialization bytes (the approved §5-E §2.3 discipline):
- `cc` **empty/absent** ⇒ **both** members omitted;
- `cc` non-empty **and decodes as a JSON array of strings** ⇒ **`cc`** carries that decoded array, elements **in presented order, no dedup, no trimming, no sorting** (normalization at delivery is the store's business, §2.3a);
- `cc` non-empty **and does NOT so decode** ⇒ **`cc_unparsed`** carries the presented string.
The two are **mutually exclusive** — never both, never neither-when-non-empty. Distinct member names (rather than one member of two possible types) keep the digest input unambiguous: an array target-set can never collide with a string. The unparsed branch exists because it is genuinely reachable at binding time: the generated schema types `cc` as a plain string, so an undecodable value passes Layer-2 validation and reaches authorization even though the conductor will later reject it at typed parse — and the descriptor must be derivable for every schema-valid invocation, acceptance never a prerequisite.

- **Verb prefix:** the literal canonical verb id + `":"` — satisfying the table's "relay verb + target id" form and keeping the value self-describing at a glance (parallel to `relay.read`'s verb+`relay_id`).
- **Totality:** `form_digest` is in the frozen R-1 required set, so the target object is **never empty** and the value is **always derivable** — no branch of the MVP can produce an unfillable cell (the `cc_unparsed` branch keeps that true for undecodable CC carriers).
- **Absence by member omission** (never `null`, never `""`, never a sentinel): the frozen §3.4 discipline. An empty-string member is treated as absent (the mapping's P-0 keeps absent optionals absent; an empty `to`/`dispatch_id`/`cc` names no target).
- **Closed object:** exactly these member names, no others — an additional member would silently bind bytes no contract names (Rail A).
- **Encoding:** JCS over the parsed member **values** — identical discipline to the approved §5-E component (`c3a8cd61…` §2.3): original serialization bytes are never inputs; equivalent serializations converge.

### §2.2 What each member contributes (and why this set)

| member | what it identifies | why it belongs |
|---|---|---|
| `form_digest` | **the governed surface submitted into** — transitively the rendered form × seat pattern × phase × tier × config digest (the frozen render-digest inputs) | master's direction verbatim: the target derived from m-2's form schema. It is the closest analogue to a file path: the *thing the action acts upon*. |
| `dispatch_id` | the **lane** the record enters | destination coordinate; carried on the submission, observable |
| `to` | the **primary recipient** addressed | destination coordinate; master's framing names the recipient |
| `cc` / `cc_unparsed` | the **additional delivery recipients** named (REVISED at review-r1 F1) | CC addressees receive real mailbox deliveries — **verified at source** (`frank/internal/store/projections.go:149-174`: `DeliveryRecipients` adds every CC address after TO, deduped; `:137-145`: one `IntentMailbox` write per resulting recipient). Two otherwise-identical submissions differing only in CC therefore have **different delivery target sets**, so omitting CC violated this doc's own rule that a target change moves the resource — and cut against the `apply_patch` precedent, which binds the COMPLETE target set, not one primary target. |

**The authority/effect distinction, stated explicitly (review-r1 F1's required clarification):** binding CC as a delivery target grants a CC addressee **no relay authority whatsoever** — CC remains context-only: a CC'd seat is not addressed, owes no action, and may not act on a relay naming it only in CC (the standing protocol rule, unchanged). Rev0's error was conflating those two: **authority** semantics (CC = context-only) with **effect-target** semantics (CC = a real mailbox delivery). `canonical_resource` is context-binding EVIDENCE of the effect's targets, so effect-target semantics govern here.

**Excluded, with reasons stated (not silence):** `body`, `headers` — content, not target (they ride `canonical_args_digest`); anything not present on the submission. *(The `headers["CC"]` fold my mapping performs is not a second CC source: it is the same value re-homed by P-3, and binding the top-level member at the invocation grain binds it once.)*

### §2.3a Target-as-NAMED, not the store's post-projection delivery set (the boundary this cell does not cross)

The object binds the targets **as the invocation names them**. It deliberately does NOT re-derive the store's delivery-set computation (TO/CC union, dedup, trimming, the address-list-header fallback — `projections.go:149-174`): that projection is another owner's domain, the descriptor binds the invocation rather than the projection outcome, and re-deriving it here would couple this cell to a rule that can change under me. **Consequence, stated honestly:** two invocations naming the same address differently (e.g. once in TO and again in CC) may yield the same eventual delivery set but different `canonical_resource` values — they are different invocations, and §2.4's invocation-not-acceptance labeling already governs exactly this class.

### §2.3 The projection property (why this is not a duplicate of `canonical_args_digest`)

`canonical_resource` here is a **projection of the args that names the target** — exactly the relationship `apply_patch`'s ordered target-set digest already has to its args in the same ratified table (that cell is likewise derivable from the patch bytes and is nonetheless required separately, because it names the resources). Consequences, both intended:
- two submissions **to the same target with different bodies** share `canonical_resource` and differ in `canonical_args_digest` — the same normal relationship two `write`s to one path have;
- a change of form, lane, or recipient **moves** `canonical_resource` — the target genuinely changed.

### §2.4 Honest labeling (what this value does and does not assert)

The descriptor binds the invocation **as presented at authorization**, before execution. Therefore: `form_digest` here is the value the submission carried, **not** an assertion that the conductor accepted it. A stale `form_digest` is a conductor-side rejection (re-render; no record created) — the descriptor neither predicts nor needs to predict that outcome. This is evidence of *what was invoked*, which is precisely C's stated role ("C is the EVIDENCE half"), never an authorization or acceptance claim.

## §3 — Observer derivability (m-3's predicate)

`local_invocation_matches_effect_descriptor` recomputes the value from the observed invocation alone: read the target members off the observed submission arguments (`form_digest`, `dispatch_id`, `to`, `cc`) → §2.1 omission + CC decode rules → JCS → SHA-256 → prefix. Every input is a member of the submission the observer already sees; **no m-2 module code and no conductor state are required** (the recipe is the contract — the same independence property the approved §5-E component carries). A mismatch between the recomputed value and the ticket-bound descriptor is a **fail** for the predicate; an unobservable invocation is `unknown` per the existing §5-C rung discipline — I claim no new verdict vocabulary.

## §4 — Reference vectors (recomputable by anyone from the §2.1 rule)

| id | submission members | target object (canonical) | `canonical_resource` |
|---|---|---|---|
| RV-1 | `form_digest="ref-digest-1"`, `dispatch_id="d-1"`, `to="m-2.implementer"`, no cc | `{"dispatch_id":"d-1","form_digest":"ref-digest-1","to":"m-2.implementer"}` | `relay.submit:7f6479f5c8b30ebe2e92ce40c56ba18865cd0ced9ccff04fc26c2c609f395995` |
| RV-2 | `form_digest="ref-digest-1"` only (all optionals absent) | `{"form_digest":"ref-digest-1"}` | `relay.submit:b98691f9093ee214362ca81afef4063ddc99f1ca5c35199ff0f07350124ac3ee` |
| RV-3 | as RV-1 but `to="m-9.planner"` | `{"dispatch_id":"d-1","form_digest":"ref-digest-1","to":"m-9.planner"}` | `relay.submit:01c6d57c3ac77e5f7b69bfa4d8cefe575789a5a9fcea73711433dfb66efd4e32` |
| **RV-4** | as RV-1 **plus** `cc="[\"m-7.planner\"]"` | `{"cc":["m-7.planner"],"dispatch_id":"d-1","form_digest":"ref-digest-1","to":"m-2.implementer"}` | `relay.submit:cca45c4ff0e2399e0df6936fa1981dd93d7ea387d061bea41fd677166ce35ec0` |
| **RV-5** | as RV-4 but `cc="[\"m-7.planner\",\"operator\"]"` | `{"cc":["m-7.planner","operator"],"dispatch_id":"d-1","form_digest":"ref-digest-1","to":"m-2.implementer"}` | `relay.submit:24793638571ee3d522763dbc0f2bbd3e6b4fa794aff64ee2fdd8e8ebeeb47a5d` |
| **RV-6** | as RV-1 plus an **undecodable** `cc="m-7.planner"` (a bare address, not a JSON array) | `{"cc_unparsed":"m-7.planner","dispatch_id":"d-1","form_digest":"ref-digest-1","to":"m-2.implementer"}` | `relay.submit:e84184b1b0ffe56114f1bb436e0f6ab76c2e5b0361385a3c80adb207829459cc` |

(`ref-digest-1` is the same sentinel as the frozen Appendix-A fixture form, so the vectors compose with the existing m-2 fixture material.) **What each pair proves:** RV-1↔RV-3 the recipient participates · RV-1↔RV-2 omission is not the empty string · **RV-1↔RV-4 the review-r1 F1 requirement — form/lane/`to` held constant, CC changed, the resource MOVES** · RV-4↔RV-5 the CC target-set contents participate elementwise · RV-4↔RV-6 the decoded and unparsed branches are distinct values (no type collision). **RV-1/2/3 are byte-unchanged from rev0** — an added optional member cannot move a vector that lacks it, which is itself the omission rule under test.

## §5 — The amendment cell (the exact bytes master's amendment carries)

Replacement text for the §5-C table's `canonical_resource` × `relay.*` cell:

> `R = relay verb + target id — for `relay.read` the `relay_id`, for `relay.project` the view token, and for `relay.submit` (which creates the record it names, so has no pre-existing target) the **submission-target digest**: `"relay.submit:" || SHA-256(JCS{form_digest, dispatch_id?, to?, cc?|cc_unparsed?})` over the submission's own members — optionals omitted when absent; the CC delivery-target set bound decoded (`cc`) or, when the presented carrier does not decode as a JSON string-array, as the presented string (`cc_unparsed`), the two mutually exclusive; binding CC as a delivery target confers no relay authority (m-2-authored, §5-C-submit).`

No other cell, field, or row of §5-C moves.

## §6 — Build obligations (RED-first; nothing here exists at `frank@c78da38`)

1. The derivation helper + its omission rule (empty ⇒ absent), the **`cc` decode-else-`cc_unparsed` branch with mutual exclusion**, closed-object enforcement, JCS encoding.
2. The RV-1…RV-6 fixture (byte-exact values recomputed from the rule).
3. Target-sharing/target-moving legs: same complete target + different bodies ⇒ equal `canonical_resource`, differing `canonical_args_digest`; **a change of EACH member in turn — `form_digest`, `dispatch_id`, `to`, and `cc` (contents and cardinality) — moves the value** (the CC leg is review-r1 F1's required obligation).
4. Observer-parity leg: an m-3-side recomputation from the observed invocation equals the ticket-bound value, **including both CC branches**.
5. Totality leg: a minimal legal submission (required members only) still yields a value; **and a schema-valid-but-conductor-rejectable submission (undecodable `cc`) also yields one** — no branch produces an unfillable cell, acceptance never a prerequisite.
6. **Value-grain leg:** two invocations whose `cc` carriers are equivalent JSON serializations of the same address array ⇒ the same `canonical_resource` (serialization bytes are never inputs — the §5-E discipline, applied here).

## §7 — Revision log

- **rev1** (2026-07-23, m-2.planner): the single `step3-relock-c-m2-submit-resource-review-r1` blocker folded — **M2-C-R1-F1**, verified at frank source before folding (`internal/store/projections.go:137-145,149-174` at `c78da38`: `DeliveryRecipients` adds every CC address after TO with dedup, and each resulting recipient gets its own `IntentMailbox` write). The finding is correct and the error was mine in a specific way worth naming: I applied the **authority** rule (CC is context-only, owes no action — true, and unchanged) to an **effect-target** question (CC is a real delivery destination). Rev0's own "a target change moves the resource" rule, and the `apply_patch` precedent of binding the COMPLETE target set, both contradicted the exclusion. Fold: `cc` joins the closed projection with a **total** two-branch rule (decoded array when the carrier decodes as a JSON string-array; `cc_unparsed` string otherwise — mutually exclusive, distinct member names so an array can never collide with a string; the unparsed branch is genuinely reachable because the generated schema types `cc` as a plain string, so an undecodable value passes Layer-2 and reaches authorization); §2.3a added to bound the claim honestly (targets **as named by the invocation** — the store's dedup/union/fallback projection is deliberately NOT re-derived here); the authority/effect distinction stated explicitly; member table, exclusion text, amendment-cell bytes, vectors (RV-4/5/6 added — RV-1/2/3 byte-unchanged), and obligations 1–6 updated.
- **rev0** (2026-07-23, m-2.planner): authored on master's ruling (4). §5-C read at the ratified rev12 (`1125b0a0…`, hash verified); frozen m-2 rev5 (`83d8e63e…`) consulted read-only and unmoved; the approved §5-E component (`c3a8cd61…`) reused for the JCS/value-level and absence disciplines. `∅` considered and rejected with reasons (§1). Reference vectors computed from the rule, not asserted.
