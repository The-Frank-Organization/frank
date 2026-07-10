# s5-b "mechanisms & versioning" — DESIGN: the ③ known-A raise · DEF-2 submit guard · the ⑤ egress drain+scanner (dormant) · zero-loss replay · §7 s5-delta legs · I-PH extensions

**DESIGN_DOC_ID:** `s5-b-mechanisms-design`
**Owner:** s5-b pair — authored by `s5-b.planner`; design-challenger + formal DESIGN-REVIEW addressee = `s5-b.implementer` (per the protocol's lineage rules)
**Dispatch:** `.relays/s5/s5-design-s5-b/DESIGN-orchestrator-planner-20260706-045327.md` (base) + `…-052753.md` (③ HOLD lifted; DEF-2 wire-in-full)
**Specification of record for ③:** `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` §2 (composed m-6 detector × m-7 mechanics + MR-1); §C4 registrations cited, not restated
**Rev:** r3 (DESIGN-REVIEW folds — r2: `…-055134` blocker 1, the ⑤ drain provenance/render contract §3.1a; r3: `…-060027` blocker 1, the scanned item carries the rendered value — `Item.Field` is now a full `RenderedField {Name, Value, Origin}` so `Scan` inspects the actual outbound bytes; findings stay Field:Class-only) · **Date:** 2026-07-06 · **Tier:** medium · **Evidence:** E1 (all code cites verified in the s5-b audits at `main @ 67ee23e`) · **GRILL_REQUIRED:** no

---

## 0. Frame + claim boundaries

**The bundle.** s5-b owns the mechanism half of Slice-5 (consumer schemas): the ③ known-A/RAISE-ONLY enforcement + fixture, the DEF-2 submit guard, the ⑤ egress scanner wired present-but-dormant at the real drain chokepoint + its fixture legs, the Q4-final versioning artifact set (zero-loss replay + read-path refusal legs; NO envelope migrator), the §7 fixture legs over the real s5-a registry delta, and the I-PH sweep extension. registry.json and every row/column change is **s5-a's single pass** (R-s5-2); we consume row ids/enums by contract.

**③ Step-1 claim boundary (mandatory on every ③ surface — [VP-W1]; master §2 wording):** Step-1 detection claims **exactly (S1)+(S2)+(S3)+fail-safe** — NOT "catches every content mis-pick." A mis-pick invisible at S1's grain, outside the merge case, referencing no live A-gate, falls back to the agent's monotonic pick + `other`→A. This boundary is registered with the fixture (fixture-file doc comment + this doc); no code/doc/tool text we write may claim broader detection.

**⑤ claim boundary:** the scanner is **present-but-dormant** — no code/doc/tool text may claim scanning is live (sweep-enforced); activation = the away-mode bridge (explicit carry, OUT here). Content detection is **best-effort and encoding-evadable — not "structurally unbypassable"**; the outbox is the only egress *the governance system offers* (D5-qualified, m-3 §7 / m-7 §9 phrasing travels verbatim).

**Locked inputs consumed (never reopened):** master §2 (③ composed ruling + MR-1 + DEF-4 shape + sequencing), M-3(h) three grains (DEF-2), R-2 (⑤ shape), Q4-final/R-1 (versioning), [VP-W7] (§7 fixture legs), Q11 (I-PH), R-s5-1..7 (write surfaces, DEF-1 assignment, drain locus discretion, replay legs, sequencing).

---

## 1. D-③ — the known-A raise (full map)

### 1.1 Where it lives

The raise stays at the exact audited locus: inside `Registry.Validate`, post-stamp / pre-commit in the one pivot (`engine/submit.go:29` stamp → `:47` Validate → `loop.go:133` commit). Two structural changes to the current partial build:

1. **The raise moves out of the per-field loop.** Today it fires only when `gate_category` is present (`validate.go:61-66` sits under the `!present || raw == ""` skip). The ruled full map includes **B-absorb** (no pick at all over known-A content), so the raise becomes a **post-loop step** in `Validate`, still the same locus, running on every submission after per-field checks pass their shape/enum/scope gates.
2. **The detector is injected engine-side.** `RenderEnv` gains a `KnownA` member (precedent: `ParentCandidates` — engine-supplied evaluation context already threaded into Validate, `validate.go:68-74`):

```go
// fieldspec side (the contract)
type KnownADetector func(cand record.Record, fields map[string]string) (member string, hit bool)
// RenderEnv: + KnownA KnownADetector   (nil ⇒ no detector ⇒ Step-1 behavior = fail-safe only)
```

```go
// engine side (the composite — constructed with tables + pinned config)
func KnownADetector(reg *fieldspec.Registry, tab *tables.T, cfg DetectorConfig) fieldspec.KnownADetector
```

### 1.2 The detector — MAX(S1, S2, S3), master §2 verbatim

- **S1 — the CQ-3 A-floor table over `(phase × record_kind)`** — primary, pick-independent. Contract: `AFloor(phase, recordKind) (member, ok)`. The table is **config** (m-6-authored section under the §J2-A-as-config home). **Config SHAPE binds at IMPL** when m-6.implementer's signal-set confirm lands (master §5.3; gates ③ IMPL-integration at the orchestrator's gate, not this design or our PLAN).
- **S2 — the referenced gate record's own `gate_category`** at the verdict path: when `resolves_gate` is set (the row s5-a now declares IS this reference — master §3(j)), look up the referenced committed gate record in the tables snapshot (the read `classifyVerdict` already performs, `submit.go:216-245`; m-7-sanctioned concurrent committed-record read); if its committed category ∈ A-set → that member. Mechanism, not config.
- **S3 — the §J2 merge-split predicate** (target-branch × protected-branch set). Contract: `MergeSplit(fields) (member, ok)` returning `merge_to_protected`. The protected-branch set is config (same home). **Open input note:** no target-branch registry row exists at `67ee23e`; the atom S3 reads is named by the m-6 confirm — if it names a new header row, that row rides s5-a's pass. Tracked in §11.
- **Fail-safe:** `other`→A stays **hardcoded** (`ClassifyGateCategory`'s unknown/unlisted fallback, `validate.go:186`), never config-editable-away.
- **Composition:** any source hit ⇒ A-worthy. When more than one source names a member, the committed member is chosen by precedence **S2 → S3 → S1** (record-specific evidence over structural floor) — a deterministic tie-break inside the ruled semantics, flagged for design-review challenge (§12), overridable by the m-6 confirm without design rework (it changes one comparison, not the shape).

### 1.3 The mechanics — token REWRITE + stamp (m-7's half, binding)

In the post-loop raise step, with `pick = cand.Headers["gate_category"]`:

| case | action |
|---|---|
| detector hit ∧ effective handling below A (pick ∈ B, or pick absent) | **REWRITE** `gate_category` := detector's named A member (else `other`); stamp `gate_category_raised: "yes"`; preserve the original pick in `gate_category_pick` (MR-1 — row shaped by m-2 in s5-a's pass; we write the value only when an actual pick existed) |
| pick == `other` (the existing minimal case) | classification stays A-raised (fail-safe); stamp `"yes"` (DEF-1 byte fix folds here — `validate.go:64` currently writes `"true"` against the `["no","yes"]` enum) |
| pick ∈ A (any A member) | already ≥ A ⇒ no raise, no stamp — RAISE-only; there is **no A→B path on any code branch** |
| pick ∈ B ∧ no detector hit | commits as picked — this IS the claim boundary (§0) |

- `ClassifyGateCategory`'s `knownA` input is wired **from the detector** — the hardcoded `false` at `validate.go:62` is retired.
- **The lattice is the A/B MAP, never enum-index arithmetic** (master §2): the index-based `belowMonotonicFloor` check class is **deleted for `gate_category`** — concretely, the `fill_constraints == "monotonic"` index check (`validate.go:58-60`, `:79-94`) is skipped for the row whose `enum_set == "gate_category"`; category monotonicity is enforced exclusively by the raise above. (The index floor stays for bool-enum rows — HUMAN_GATE_REQUIRED/GRILL_REQUIRED — where order is the lattice. Render-side `monotonicOptions` is untouched: no gate_category floor is ever supplied in Step-1, and the render affordance question is m-2/m-6 surface, not ours.)
- Atomicity: the rewrite+stamp mutate the candidate **before** the single `st.Commit` in the same loop turn — the raise is atomic with acceptance, per Q5. A raised A token makes `isGateCandidate` true (`submit.go:120-122`), so the A-path consequences (park via `completePark`, ODB/outbox item via `completeOutbox`) follow mechanically — DEF-4's consequence leg rides existing machinery.

### 1.4 The ③ fixture — DEF-4 shape (adopted), `test/fixtures/s5_gate_raise_test.go`

1. **B-pick over detector hit** ⇒ committed record carries the A token AND `gate_category_raised: "yes"` AND the A-path consequences (park record + outbox item present) AND `gate_category_pick` = the original B token.
2. **B-absorb** (no pick, detector hit) ⇒ same, `gate_category_pick` absent.
3. **No-A→B negative** — an A pick with any detector state never lowers, never stamps.
4. **DEF-1 byte assertion** — an `other` pick end-to-end: committed `gate_category_raised == "yes"` (not `"true"`), and the committed record **re-validates against its own registry** (the enum-conformance bug class pinned dead).
5. **Full-map extension** incl. the incoming `routing_escalation` member — consumes s5-a's landed delta; sequenced per R-s5-6 (§8).
6. **Claim-boundary negative** — a B mis-pick invisible to S1/S2/S3 commits as B (the boundary is registered in the fixture, with master §2 wording in the file doc comment).

Detector injection in fixtures: the composite with a **fixture-scoped config** (A-floor rows + protected set + the tables-backed S2) — the config SHAPE placeholder is isolated in one constructor so the m-6 confirm binds at IMPL without fixture rework.

---

## 2. D-DEF-2 — the submit guard (M-3(h) confirmed; wire in full)

One validator rule, three binding grains (m-7, verbatim):

1. **Typed REJECT, never silent-strip.** In `Registry.Validate`, before the per-field loop: any **non-empty header** whose spec is `owner ∈ {system, computed}` or `fill_constraints ∈ {system_only, computed_result}` ⇒ `Violation{Field: <id>, Class: "system-owned"}` (Field:Class, path-free, formatter-compatible). A lane writing system headers is signal; strip is incoherent for computed-later fields.
2. **Keyed on the submission channel.** The guard lives in the submit-path validator — *every* record entering via `submit()` hits it, lane AND operator alike. Conductor-internal authorship (genesis `store.Init`, obligation-derived records, `faultOutcome` records) never passes through `reg.Validate` — the legitimate-writer set is structural, not seat-conditional. The ③/computed stamps happen **after** the guard inside the same Validate call, so the conductor never trips its own rule.
3. **The envelope asymmetry stays.** FROM/ROLE keep `seat.Stamp` overwrite semantics (identity by overwrite — `submit.go:29`); `schema_version`/`relay_id` keep their engine overwrites (`:34-36`). The guard is **header-grain only**: computed state by refusal, identity by overwrite.

Closes DEF-2 generally: `failing_edge`, `delivery_state`, `gate_category_raised`, and every new dormant system/computed row (s5-a's O-2/Q9 homes included) become non-lane-suppliable at the same stroke. `record_kind` (seat_scoped_enum) and free-text/X-* rows are untouched.

**Fixtures (per class), `test/fixtures/s5_submit_guard_test.go`:** lane-supplied `failing_edge` (owner:system) rejects `system-owned`; lane-supplied `gate_category_raised` (computed_result) rejects; lane-supplied `delivery_state` rejects; **operator-channel** submission of the same rejects (grain 2); positive leg — the conductor's own raise stamp still commits (shared with §1.4 case 1); bounce text enters the I-PH sweep (§6).

---

## 3. D-⑤ — the egress drain + scanner (R-2 shape; locus = new package `internal/egress`)

**Locus decision (R-s5-4, ours):** a new `internal/egress` package — not `store/drain.go`. Rationale: the chokepoint is m-7-hosted but the rule semantics are m-3 policy; keeping it out of `internal/store` keeps the TCB store package free of content-policy code, gives the away-bridge a single import point, and makes dormancy structurally visible (nothing imports `egress` in production). No import cycle: `egress` imports `store`/`gate`/`record`; nothing imports `egress`.

### 3.1 Shape

```go
package egress // PRESENT-BUT-DORMANT: no production caller in Step-1; activation = the
               // m-6 away-mode bridge (explicit carry). Detection is best-effort and
               // encoding-evadable; the outbox is the only egress the governance
               // system offers (D5 residual accepted) — see design §0 claim boundary.

type Class string // "confidentiality" | "safety"
type Verdict struct { Blocked bool; Findings []Finding } // Finding renders Field:Class only
type Origin struct { ConductorODB bool }   // runtime-only value — see §3.1a provenance rule
type Item struct {                         // (r3) the scanned unit = ONE rendered field
    Meta   gate.OutboxItem
    Source record.Record
    Dest   string
    Field  RenderedField                   // Name + Value + Origin — §3.1a
}

func Scan(item Item, rules Rules) Verdict  // rules classify item.Field.Value (the actual
    // outbound bytes — which for renderer-generated slots like the ODB model_name may
    // exist NOWHERE in source.Headers/Body); findings still render Field:Class ONLY
    // (item.Field.Name + class token — never the value) for I-PH
func Drain(st *store.Store, rules Rules, render Renderer) (Report, error) // THE real call
    // site: walks outbox/<item>.json, resolves each source record (st.Read on
    // source_record_ref), calls render() to obtain the egress-render fields (§3.1a),
    // and scans EVERY RenderedField (Value with its paired Name + Origin) pre-send;
    // Blocked ⇒ withheld (item never sent; reported as the non-terminal egress_blocked
    // disposition in Report — no new terminal token, no new record_kind)
```

### 3.1a The provenance/render contract at the real call site (r2 — review blocker 1)

The live outbox file carries only the seven `gate.OutboxItem` fields (`derived.go:8-16`; the obligation writer emits the same shape, `obligation.go:149-181`) — `Dest`, `Field`, and `Origin` are **not reconstructable from outbox bytes**, and must never be. They come from the **conductor-side egress renderer**, the component that turns an outbox item into the outbound payload:

```go
type RenderedField struct { Name, Value string; Origin Origin }
type Rendered struct { Dest string; Fields []RenderedField }
type Renderer func(meta gate.OutboxItem, source record.Record) (Rendered, error)
```

- **Origin is a property of the render pipeline, never of the bytes.** A renderer sets `Origin.ConductorODB` only on fields **it itself generated** (the m-6 conductor-rendered ODB slots); every field whose value derives from record content (headers/body — lane-authored bytes) is emitted lane-origin. `Origin` has **no serialized form**: it is never parsed from record headers, body, or outbox JSON, and never persists — so lane content cannot set it, and nothing plantable in a store can resurrect it. This is the structural statement behind fixture leg (c).
- **The shipped default:** `egress.DefaultRenderer` — real production code for the item kinds that exist today (gate/held): renders the source record's outbound view with **every field lane-origin** and `Dest` from the item's addressing. Under the default renderer **nothing is exempt** — fail-closed is the shipped posture.
- **The ODB renderer is the away-bridge carry** (m-6 §3 owns the ODB render; no live ODB exists in Step-1 — Q6). The ⑤ pass leg is therefore proven through the SAME `Drain(st, rules, render)` contract with a **conductor-shaped ODB renderer stand-in** (fixture-scoped, the exact analogue of Q6's fixture-scoped registry view): it generates the `model_name` slot itself and marks only that field `ConductorODB`. When the real m-6 renderer lands at the away-bridge, it replaces the stand-in behind the same `Renderer` contract — no `Drain`/`Scan` rework.

- **Two classes, ruled semantics:** *safety* (secrets / PII / auth-URLs) is **never exempted** anywhere; *confidentiality* (model-name tokens) has the **single typed carve-out**: `class==confidentiality ∧ item.Field.Name=="model_name" ∧ item.Field.Origin.ConductorODB ∧ item.Dest=="operator"`. Fail-closed: an unclassifiable finding ⇒ block (`other`→block).
- **Rules:** Step-1 compiled-in defaults (minimal honest patterns — enough to prove the chokepoint + carve-out semantics; breadth is NOT the Step-1 claim); the rule-set config surface is m-3/operator policy, documented dormant.
- **ODB identification without live enum widening (Q6):** the carve-out keys on `Origin.ConductorODB` — conductor-generated typed context (m-6's exempt-mark binding), supplied per §3.1a by the renderer, never parsed from bytes — never on a live `record_kind` member; `registry.json:72` gains no `ODB` token and `validateRecordKind` (`submit.go:124-157`) is untouched. Fixtures use a fixture-scoped registry VIEW only.
- **Disposition:** blocked items stay in the outbox and are reported `egress_blocked` (non-terminal — m-3 §7; the live park+resummon wiring is the away-bridge carry, OUT).

### 3.2 The ⑤ fixture, `test/fixtures/s5_egress_test.go` — m-3's acceptance pair at the real call site

Constructed store, outbox items produced via the **real** obligation path (`completeOutbox`), then `egress.Drain(st, rules, render)` invoked directly (the fixture-enabled drain leg — the test is the ONLY activator). **Every leg goes through the same `Drain`+`Renderer` contract of §3.1a — never by direct `Scan(Item)` construction** (r2; direct-construction legs would prove neither the chokepoint nor the origin claim):

- **pass leg:** the conductor-shaped ODB renderer stand-in (§3.1a) generates the `model_name` field, `Dest=="operator"` → the confidentiality scan **passes** exactly that field (the carve-out);
- **(a)** the same model-name in a lane-origin rendered field (any non-ODB item under `DefaultRenderer`) **blocks**;
- **(b)** a secret / PII / auth-URL rendered from an ODB item's record content still **blocks** (safety class scans every field regardless of Origin);
- **(c)** lane-supplied "exempt-marked" content still **blocks** — a source record whose lane-authored headers/body mimic an exempt mark renders lane-origin under the same renderer (Origin is pipeline-derived, never byte-parsed, never persisted — §3.1a);
- **fail-closed leg:** unclassifiable finding blocks; and under `DefaultRenderer` alone, NOTHING passes the confidentiality class (the shipped posture is exemption-free);
- **dormancy assertion:** grep-style test that no production package imports `internal/egress` (structural dormancy pinned).

---

## 4. D-replay — the zero-loss replay + read-path refusal legs (Q4 final)

**Layout:** new subpackage `test/replay/zeroloss/` — structural independence from `test/replay/harness.go` and its external oracle paths (`harness.go:24-25`); zero imports of the parity-replay code, no `dispositions.json`.

```go
func Replay(root string, reg *migrate.Registry) (Report, error)
// walks the store's redo entries + records/ through migrate.Reader — the same read
// surface the seat read tool uses (cmd/frank/main.go:302)
```

**Assertions (count / identity / canonical-wins):**
- **count** — every committed record is readable through the versioned read path; `Report.Read == Report.Total`; zero lost (quarantined records are accounted, not silently skipped);
- **identity** — for `SourceVersion == Current`, the view deep-equals the stored record (`Apply` no-op at current, `migrate.go:62`);
- **canonical-wins** — where the store holds a canonical record AND a derived projection (the outbox pair), the replay reads the record as authority (the `gate/derived_test.go:123-147` pattern generalized).

**Legs (R-s5-5):** the **constructed-store leg is MANDATORY** (deterministic CI): `store.Init` fresh + variety per the s4 pattern — accepted/rejected/held, a gate record (park + outbox derivation), owed_item/owed_disposition, a config_change chain. The **archived-copy leg is optional** pending M-4: parameterized by `FRANK_S5_REPLAY_STORE` (skip when unset); asserts the same invariants with counts derived from the copy's own redo log (never hard-coded); the archive itself is operator-held — propose-only, never touched in place.

**Refusal legs at the READ path** (extends `migrate_test.go:51-69`, which proves them at `Apply()` grain only): planted record files with `SchemaVersion` 0 (`ErrUnversioned`), `Current+1` (`ErrFutureVersion`), and a gap case (version 1 against a local `Registry{Current:2}` with no step) surfaced **through `Reader.Read` over a real store**. Refusal text joins the I-PH sweep (§6; the three bare error strings are already swept, `iph_test.go:43-45` — we extend, not rebuild).

**NO envelope migrator (R-1 discipline):** `migrate.Current` stays **1**; zero steps registered. The named false-trigger: treating m-2's compat LABEL (MINOR/additive — the enum delta + new dormant rows) or the registry version-label bump (`s5-fieldspec-v3`, s5-a's byte per R-s5-1) as a record-SHAPE change. Registry contents are config-member bytes governed by §7 digests; the envelope shape is untouched by every s5 deliverable.

---

## 5. D-§7 — the s5-delta legs (adapted [VP-W7]), `test/fixtures/s5_config_change_test.go`

The five legs against TEST stores with the **real s5-a registry** as the config_change body (the s4 suite proves the class with a synthetic marker mutation — we rebuild nothing, `s4_config_change_test.go` cited leg-by-leg in our audits):

1. operator-authorized record shape (s5-payload variant of `TestConfigChangeNonOperatorRejected`);
2. old→new digest over the real delta (`fixtureDigestWithMember` pattern, `:353-365`);
3. **no re-genesis** — with the implementer's explicit assertion: genesis record count == 1 after the chain (not just `ErrGenesisExists` at `Init`);
4. **phase-0 acceptance** — with the explicit forced walk: stop/start + a direct `st.ValidateGenesis(pinned)` call over the chained digest (`genesis.go:104-141`), not only the implicit restart success;
5. stale-form re-render + re-rendered success — where the re-rendered form **shows the new enum membership** (`routing_escalation` present in the fresh `gate_category` options).

Store construction: the s4 pattern exactly — `store.Init` fresh per-test root + prior chain moves (`newS4ShimHarness` consumed read-only). **Sequencing (R-s5-6):** these legs consume s5-a's landed registry — they integrate after s5-a's registry pass at the orchestrator's gate; everything else in this design parallelizes.

---

## 6. D-I-PH — sweep extension (Q11), `test/fixtures/s5_iph_test.go`

Every seat-deliverable surface this bundle adds, through the existing NF-S18/NF-S15 patterns (`assertNoPathFamilies`, `iph_test.go:106-115`; `assertNoS4IPHLeaks`, `s4_iph_test.go:234-246`):

1. ⑤ scan verdict/Finding strings (Field:Class only; class tokens embed no path/config value/store layout);
2. ⑤ drain diagnostics (reuse the `safeReason` class pattern, `loop.go:208-210`);
3. DEF-2 `system-owned` bounce text;
4. zero-loss/refusal wrapper text (any NEW wrapper joins the already-swept migrate error strings);
5. ③ raise — adds a header value, no new rendered string; any bounce leg emits Field:Class like every violation.

**The formatter valve is preserved untouched:** `bounce.Format` (`formatter.go:11-26`) stays Field:Class-only; no surface we add requires Reason in output; no formatter change. F-GATE-3 boundary discipline travels verbatim in any doc/comment text.

---

## 7. File/function targets (all within R-s5-2 surfaces)

| Target | Change |
|---|---|
| `internal/fieldspec/validate.go` | DEF-2 guard (pre-loop); the post-loop ③ raise step (rewrite+stamp+pick, DEF-1 byte); retire the hardcoded `knownA=false`; skip index-floor for `enum_set=="gate_category"` |
| `internal/fieldspec/render.go` | `RenderEnv` + `KnownA KnownADetector` (type only; render behavior untouched) |
| `internal/fieldspec/fieldspec_test.go` | classifier-grain additions (s5-b per R-s5-2) |
| `internal/engine/` (new `detector.go`) | the MAX(S1,S2,S3) composite + `DetectorConfig` (shape binds at IMPL on the m-6 confirm); wiring into the submit handler's env |
| NEW `internal/egress/` | `Scan`/`Drain`/rules — §3 |
| NEW `test/replay/zeroloss/` | §4 |
| NEW `test/fixtures/s5_gate_raise_test.go`, `s5_submit_guard_test.go`, `s5_egress_test.go`, `s5_config_change_test.go`, `s5_iph_test.go` | §§1-6 fixtures (filename-disjoint from s5-a's `s5_registry_*_test.go`) |
| **Untouched** | `registry.json`/`registry_test.go`/`render_test.go`/`validate_test.go` registry-content fixtures (s5-a); `internal/bounce/formatter.go`; `internal/migrate/migrate.go`; lineage/parenting/codec (transport-fix, OUT); the archived store |

## 8. Boundary contract

- **Writes:** the ③ raise (committed `gate_category` rewrite + `gate_category_raised`/`gate_category_pick` values) at the validate locus; the DEF-2 reject class; the dormant egress package; the zeroloss harness; the five fixture files.
- **Reads:** s5-a's row ids/enums by contract — `gate_category`/`gate_category_raised` (:92/:93), the `named_enums` delta (`routing_escalation`), the MR-1 `gate_category_pick` row (m-2-shaped in s5-a's pass), the `resolves_gate` row (M-3(j) — it IS the S2 reference), the landed registry bytes for §5.
- **Target entity:** the engine commit path + the test battery. **Downstream consumers:** the s5 exit gate; the transport-fix relaunch (the replay becomes its blessed-store proof); m-6's bucket routing (keys on the committed rewritten token — master §2 "just works"); the away-bridge (imports `egress` at activation).
- **Proof:** E2 at exit — battery green, zero regression, byte-exact `{accepted, rejected, held}` untouched, guardrail surface exactly `submit`/`project`/`read`.
- **No-consumer action:** reject speculative surfaces (nothing here ships without a named consumer above).

## 9. Acceptance criteria (PLAN inherits these)

1. ③: all six §1.4 fixture legs green; `ClassifyGateCategory` knownA wired from the detector; no A→B on any branch (negative leg); the index-floor check dead for gate_category; claim-boundary text present in the fixture file.
2. DEF-2: all §2 fixture legs green incl. the operator-channel leg; stamp-after-guard ordering proven by the ③ positive leg.
3. ⑤: pass + (a)/(b)/(c) + fail-closed + dormancy legs green, every leg driven through `egress.Drain(st, rules, render)` under the §3.1a Renderer contract at the real call site — no direct `Scan(Item)` construction in any acceptance leg; the Origin no-serialization property stated in the package doc.
4. Replay: constructed leg green with all three assertion families; refusal legs at `Reader.Read` grain; archive leg skips cleanly when unset; `migrate.Current == 1` asserted.
5. §7: five s5-delta legs green post-s5-a-integration, incl. the explicit genesis-count and forced-ValidateGenesis assertions.
6. I-PH: every §6 surface swept; formatter byte-untouched.
7. Battery: `go build ./...` + `go test ./...` all-green; zero regression vs the 21-ok baseline.

## 10. Out of scope + anti-half-fix guards

OUT (escalate before touching): registry.json + registry-content test files (s5-a); step-(d) away-bridge set; transport-fix pre-work; live egress activation; live `record_kind` widening; the archived store (propose-only); render-side gate_category affordance (m-2/m-6).
**Anti-half-fix:** the DEF-1 byte fix ships WITH the ③ fixture (never alone — the stamp path stays untested otherwise); DEF-2 rejects, never strips (a strip would half-fix into silent data loss); the ⑤ scanner ships WITH its dormancy assertion (a scanner without the no-production-importer pin is a live-claim risk); the replay ships WITH the refusal legs (zero-loss without refusal proves half the Q4 set).

## 11. IMPL-binding dependencies (none blocks the PLAN lock)

1. **m-6.implementer signal-set confirm** — binds the S1 table shape, the S3 target-branch atom, and (if it names a new header) one s5-a row; gates ③ IMPL-integration at the orchestrator's gate (master §5.3). The detector reduces to `(hit, member)` regardless.
2. **s5-a's registry pass** — `routing_escalation`, `gate_category_pick` (MR-1), `resolves_gate` rows; gates the §5 legs + fixture leg §1.4.5 (R-s5-6).
3. **M-4 (operator)** — the archived-store copy for the optional replay leg; the constructed leg does not wait.

## 12. Design decisions logged for review challenge (s5-b.implementer)

(i) S2→S3→S1 member precedence (§1.2); (ii) detector threading via `RenderEnv` vs a new `Validate` param; (iii) the `system-owned` violation-class token; (iv) `internal/egress` locus vs `store/drain.go`; (v) the raise as a post-loop step (required for B-absorb); (vi) DEF-2 covering `owner:computed` alongside `computed_result` (M-3(g) kept the row computed — the guard, not a row flip, closes suppliability); (vii) `test/replay/zeroloss` subpackage vs same-package file; (viii) `Origin.ConductorODB` as the carve-out key (vs a typed field-mark on the item schema).

**r2 fold-log:** review blocker 1 (`DESIGN-REVIEW-implementer-20260706-055134.md`) — `Drain(st, rules)` had no defined source for `Dest`/`Field`/`Origin` (the outbox file carries only the seven `OutboxItem` fields). Folded as §3.1a: the `Renderer` contract at the real call site (the reviewer's first acceptable shape — a conductor-side resolver deriving the render fields from the source record), the shipped exemption-free `DefaultRenderer`, the fixture-scoped ODB renderer stand-in (Q6-analogue), the Origin never-parsed/never-persisted rule, and the no-direct-`Scan`-construction requirement on every acceptance leg. Compatible checks (③/DEF-2/replay/§7/I-PH/boundary) untouched.

**r3 fold-log:** review blocker 1 (`DESIGN-REVIEW-implementer-20260706-060027.md`) — r2's `Item` kept `Field string` with no rendered value, so `Scan` had no bytes to classify (fatal for renderer-generated slots like the ODB `model_name`, whose value may exist nowhere in `source.Headers`/`Body`). Folded as the reviewer's suggested shape: `Item.Field` is a full `RenderedField {Name, Value, Origin}`; `Drain` scans each `RenderedField.Value` with its paired `Name` and `Origin`; findings render `Field:Class` only (name + class token, never the value) preserving I-PH. The r2 provenance fix and all six compatible areas untouched.
