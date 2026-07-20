# s8 OBSERVE-SPINE Implementation Plan (Step-2 phase-opener slice)

> **For agentic workers:** REQUIRED SUB-SKILL: use `superpowers:executing-plans` (or `superpowers:subagent-driven-development`) to implement this plan task-by-task. Steps use checkbox (`- [ ]`) tracking. This is a **governance slice PLAN**: it maps four EFFECTIVE design locks to reviewer-gated build tasks; it re-derives no design. Every task cites the lock/fixture it discharges. Slice-internal build relays live in `frank/` per `sprint-doc-setup`; this plan + its plan-review live under `master/relays/s8-plan-m3/` for master visibility.

**Goal:** Ship the thinnest end-to-end **observed** relay — the conductor observes done-ness at send time inside the atomic `submit()`, captures E1/E2 evidence, and stamps it conductor-side, so a false "done" is rejected pre-send and a true "done" accepts with observed stamps.

**Architecture:** Rides the Step-1 conductor + existing runtimes. A config-derived `observe` layer (off at genesis, flipped by an operator `config_change`) activates the Block-A observe field-set; an `observe_gate()` hook in the serialized commit path (post-form/lineage + post-`slot_in`, pre-append) runs phase-shaped done-predicates via a conductor-owned check registry (E1 in-process; E2 suite-class through an unprivileged m-7 executor), then stamps the conductor-computed observe fields and drives the terminal token `{accepted, rejected, held}`.

**Tech Stack:** Go (frank), the existing `fieldspec`/`engine`/`store`/`egress`/`obligation` packages; new `observe` + `executor` packages; the `test/fixtures` + `test/invariants` batteries.

**DESIGN_LOCK_ID (consumed, EFFECTIVE — this plan builds against them, changes none):**
- `s8-design-m3-registry` r1+r1a+r1b (m-3 — check registry, E1 {read-file, git-status} + E2 run-suite, §6.1 I-PH contract, the isolation-probe, the verdict binding, the §3.3 disposition, the ratified-interim defaults w/ the **mandatory s10 sunset**).
- `s8-design-m7-executor` r4 (executor host — `spawn()`→`CheckVerdict`, RunResult host-internal, FX-EXE-1..6, per-run workdir, §2.7 kill state machine, v1 handle-surface ceiling).
- `s8-design-m7-config` r13 (config host — the `present_layers` knob + `version` carrier, genesis §5, bless/adoption §5.1, singular-catalog §5.2/step-4.5, FX-CFG-1..15).
- `s8-design-m2-grammar` (approved — Option B config-derived `PresentLayers`; the s8 registry changeset).

**DESIGN_RECORD_KIND:** design-doc (four locks) · **RUN_ID:** master · **branch off:** `frank/ main@691d034`.

**Revisions:** **r2 (2026-07-11)** — folded m-3.implementer PLAN-REVIEW F1 (`s8-plan-m3/PLAN-REVIEW-implementer-20260711-155230`, must-revise): added **Task 6.5** for the m-2-locked staged pipeline (step-3 `authority_class` compute + m-7-hosted step-4.5 profile-aware completeness + Option-B `surface_intent` derivation) with the owner split; + the s8-dispatch addendum (`s8-dispatch/SITREP-orchestrator-planner-20260711-163144`): the SEQ-1 fresh-genesis-first operator steer, **Rails A/B** as the build-time review criterion, and the **s9 fuzzy-claim adjudication** item registered as NOT-s8-scope (seam kept open). SEQ-1 m-7 half CONFIRMED (`…-151621`); m-2 half pending — the pre-token gate holds.

---

## Global Constraints (every task inherits these — verbatim from the locks)

- **No c2 mechanism change.** The terminal enum is byte-exact `{accepted, rejected, held}`; the §3.1 observer-only write-allowlist stands; the §3.3 two-axis disposition stands; **R2** — no model-derived predicate enters any gate.
- **Ratified-interim defaults (RATIFIED 2026-07-11, `step2-plan/…-025643`) — built AS INTERIM, labeled, with a MANDATORY s10 sunset:** E1 hard-timeout ≤ 5s / E2 ~120s (silent hard-kill = testing-phase only; s10 adds operator-in-the-loop kill/extend) · the `side_effecting` auto-run allowlist is **empty** (fail-closed floor, not sunset) · the operator gate is a **static committed-config allowlist** (replaced by the live m-6 prompt at s10). The s10 sunset is a named s10 exit condition — s8 does not weaken it, s8 only labels it interim.
- **Activation is a governed, visible act:** `observe` flips on/off ONLY by an operator-authored §7 `config_change`, restart-effective, no hot reload, no code default (`s8-design-m7-config` r13; genesis is `observe:false`).
- **v1 executor ceiling = handle-surface absence + STATED same-uid D5 residual** — never "by construction"; NO OS-sandbox work in this slice (shelved future feature).
- **§6.1 I-PH contract (m-3-owned, conductor-enforced):** check params lane-scoped/normalized (absolute + `..` refused); `run-suite.target` a registry enum, not command text; `schema_ref` a registry id, not a path; verdicts/`failing_detail`/`executable_claim_results` path-free + effective-config-value-free + bounded/symbolic. I-PH swept over every new seat-visible surface (`STEP-2-KICKOFF.md:52`).
- **Battery/trail discipline:** serialized battery, FILE-captured, sequence-honest reporting; the ten INV-CATALOG laws stay green (census/descriptor moves per the config lock §3 are amendments, not drift); the dogfood evaluation names which S7A-TRAIL-FINDINGS failure classes the conductor demonstrably kills.
- **Merge is operator-only:** the merge-decision relay goes TO the operator; the executor acts only on a grant relay carrying `HUMAN_MERGE_AUTHORIZATION` at grant time.
- **Rails A/B — the build-time review criterion for every NEW surface (r2, `…-163144` addendum 2; locked designs are settled, this governs what the slice ADDS):** **Rail A** — open-vs-closed chosen per surface, stated in one line: additive/open where an old reader ignoring it loses only detail; closed/fail-closed ONLY where ignore-unknown changes the *meaning of acceptance* (state which + why). **Rail B** — no new adversarial-shaped mechanism or claim: if a proposed addition only makes sense against a *malicious* agent, it bounces; drift/crash-safety/history-truth/provenance functions stay in-scope, worded to **confusion grade** (never "tamper-proof/by-construction"). Every task's new-surface acceptance includes its one-line Rail-A choice + a Rail-B pass.
- **s9 seam kept open (r2, `…-163144` addendum 3 — NOT s8 scope):** the s8 spine stays **deterministic-only (E1+E2 as dispatched)**; the s9 fuzzy-claim adjudication rung (an independent cheap model judging claims, egress-fenced) is m-3-owned s9 work. The check-registry is **additive/open (Rail A)** so a future adjudication `CheckEntry` can be added without reopening the spine — do not close that seam.

---

## SEQ-1 — the named sequencing answer (order of record + dogfood store choice)

**The forcing constraint:** genesis §5 hard-pins `fieldspec = s7a-fieldspec-v5` (SHA-256 `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`; FX-CFG-7 hard-codes it), but the s8 registry changeset (adoption + catalog tokens at `registry.json:85`/`:136`, `surface_intent` `required_when` removal) mutates `registry.json` → a **successor** fieldspec. Genesis cannot both pin v5 AND embed the changeset. Therefore:

**Order of record (proposed; routed to m-7/m-2 for confirm via `s8-plan-m3-seq1/COORD`):**
1. **GENESIS FIRST** — the fresh store composes the three members at their genesis-pinned bytes: `fieldspec = s7a-fieldspec-v5` (`1ef6abab…`, FX-CFG-7 unchanged), `engine` = current + `present_layers.observe:false`, `catalog` = blessed catalog. Genesis writes members **directly** (no `submit`, no fill-validation), so v5-without-the-new-tokens is correct at genesis.
2. **REGISTRY CHANGESET** as a governed **fieldspec transition `v5 → successor`** (the ONE m-2-reviewed moment): adds `adoption` + `catalog` to the `config_member` enum (`registry.json:85`) and `member.seat_scope.operator` (`:136`); removes `surface_intent`'s static `required_when {layer_present:observe}` + `visible_when` (Option B). Lands on the fresh store as the forward-transition (FX-CFG-10 lawful-transition leg) and updates the source `registry.json` as the new shipped state. This makes the tokens **live** before any adoption/catalog submission and puts Option B in place before `observe` requires the Block-A set.
3. **OBSERVE-ACTIVATION `config_change`** — operator-authored §7 `config_change` flipping `present_layers.observe:false→true`, restart-effective; phase-0 re-walks the composite digest; `PresentLayers` gains `observe`; the Block-A required-set activates (FX-CFG-1 end-to-end).
4. **BLESS/ADOPTION path (§5.1)** — built + fixtured (FX-CFG-12/13/14/15) as the **existing-store migration**, NOT the dogfood path.

**Which store the dogfood relaunch uses: the FRESH-GENESIS spine store** (steps 1–3), not the blessed shipped store. Rationale: fresh genesis sidesteps the §5.1 bootstrap deadlock (`submit.go:340-343`/`genesis.go:211-219` have no catalog/adoption arm) entirely; the bless path is validated as the migration route but the dogfood rides the clean fresh store.

**Fill-gate exposure — STATED, not assumed (routed to m-7/m-2):** the **offline bless mode** (legacy stores) is a pre-serve bootstrap that writes the `adoption` record + materializes the `{catalog, engine}` members **directly** (provenance `channel: bless`), **NOT** through the fill-validated `submit` path — so it is **not** gated by the `registry.json:136` `seat_scope.operator` token (which the legacy store's v5 registry lacks). **After** bless (or on the fresh-genesis store post-changeset), a **LIVE** operator catalog/adoption submission IS fill-validated and requires the `:136` token — present post-step-2. m-7 owns the bless-mode exposure statement; m-2 owns the `:136` token; both confirm on the SEQ-1 COORD before T2/T10 lock.

**Operator steer (r2, `s8-dispatch/…-163144` addendum 1) + confirm status:** fresh-genesis-first is now the **operator-ratified SEQ-1 default** (the directive proves the spine as a thin slice on a fresh store). The **§5.1 bless/adoption machinery stays locked-as-designed**; T10 **builds + fixtures the bless mechanism in-slice** (FX-CFG-12/13/14/15), but **ACTIVATING it against the actual shipped live store is the public-release-migration moment — a NAMED CARRY, not s8** (no existing store needs upgrading during the s8 dogfood). **m-7 half CONFIRMED** (`s8-plan-m3-seq1/SITREP-planner-20260711-151621`: genesis pins v5 exactly / FX-CFG-7 stays valid / the changeset is the post-genesis `v5→successor` transition; bless bypasses the *seat* fill-gate offline while every *trusted* §5.1 validation still runs — no correction). **m-2 half pending.** Pre-token gate stands: **no delegated token until both SEQ-1 confirms are in hand; a correction reissues the affected tasks.**

---

## File map (what gets created / modified)

- **Create** `frank/internal/observe/gate.go` — `observe_gate()`, the §3.1 observer-only write-allowlist, the §3.3 disposition→token map, verdict binding.
- **Create** `frank/internal/observe/registry.go` — the `CheckEntry` table + selection + §6.1 param validation (conductor-side, pre-spawn).
- **Create** `frank/internal/observe/checks_base.go` — E1 `read-file`, `git-status` (in-process, no executor).
- **Create** `frank/internal/executor/executor.go` — `spawn()`, per-run workdir, §2.7 kill state machine, verdict redaction (m-7 build-fidelity).
- **Modify** `frank/internal/engine/submit.go` — insert the observe hook post-lineage(`:67`)/post-`slot_in`, pre-intents(`:98`); add the `config_change` `catalog`/`adoption` arms (`:340-343`); **wire step-3 `authority_class` compute (post-`slot_in`/gate-category, pre-observe) + the step-4.5 completeness stage (post-observe, pre-append)** — T6.5.
- **Modify** `frank/internal/engine/loop.go` — the commit point (`:142-159`) the handler-side observe hook precedes (per plan-review confirmation).
- **Create** `frank/internal/engine/completeness.go` — the **m-7-hosted step-4.5 profile-aware completeness stage** + Option-B `surface_intent` derivation (T6.5).
- **Modify** `frank/internal/config/config.go` (`Load`/`Digest` `:36-49`) + the `engine` member — add `present_layers` map + `version` int.
- **Modify** `frank/internal/fieldspec/predicate.go` (`DefaultLayers` `:17-23`) → config-derived `PresentLayers`; `validate.go:29` + `render.go:54/:259` consume the threaded context.
- **Modify** `frank/internal/fieldspec/registry.json` — `:85` enum_set (+`adoption`,`catalog`), `:136` seat_scope.operator, the `surface_intent` row (Option B).
- **Modify** `frank/internal/store/genesis.go` (`StoreRootConfigPaths` `:86-91`; genesis composition `:211-219`) — three-member genesis + the offline bless mode.
- **Create/extend** `frank/test/fixtures/s8_*_test.go` — FX-EXE-1..6, FX-CFG-1..15, the m-3 §8 negative set; `frank/test/invariants/` stays green.

---

## Milestone A — config foundation + the activation seam

### Task 1: config `present_layers` knob + `PresentLayers` threading + three-member genesis
**Discharges:** `s8-design-m7-config` r13 §1/§5 (activation knob, version carrier, genesis), `s8-design-m2-grammar` §1.2/§3 (config-derived PresentLayers, one immutable context). **m-7 build-fidelity rides here.**
**Files:** Modify `frank/internal/config/config.go` (+ engine member struct), `frank/internal/fieldspec/predicate.go:17-23`, `validate.go:29`, `render.go:54/:259`, `frank/internal/store/genesis.go:86-91,211-219`. Test `frank/test/fixtures/s8_config_activation_test.go`.
**Interfaces — Produces:** `EngineConfig.PresentLayers map[string]bool` (+ `Version int`, initial `1`); `config.PresentLayers(loaded) map[string]bool` = `{store,form,lineage}` always + `observe` iff the knob is set; the render/validate/grant-digest sites consume **one immutable** `PresentLayers` value derived at load. Genesis composes `{fieldspec(1ef6abab…), engine(observe:false), catalog}`.
- [ ] Write failing FX-CFG-7 (genesis composition — three members present; recorded composite digest == `config.Load(...).Digest`; fieldspec bytes hash == `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`).
- [ ] Write failing FX-CFG-1 (knob flip end-to-end — genesis `observe:false` → `config_change` → restart → phase-0 passes → post-restart render requires the observe-owned fields).
- [ ] Write failing FX-CFG-2 (the A-1 negative — N foreign accepts across the flip's config generation cause zero digest bounces within a generation).
- [ ] Implement `present_layers`+`version` on the engine member; derive+thread the immutable `PresentLayers`; compose genesis.
- [ ] Run FX-CFG-1/2/7 → PASS; the ten INV-CATALOG laws green; commit.

### Task 2: the s8 registry changeset (fieldspec `v5 → successor`) — the ONE m-2-reviewed moment
**Discharges:** `s8-design-m2-grammar` Option B + the registry changeset; `s8-design-m7-config` r13 §5.2 (catalog token). **m-2.planner executes inside the slice (sole-writer on `registry.json`); SEQ-1 confirmed first.**
**Files:** Modify `frank/internal/fieldspec/registry.json` (`:85` enum_set, `:136` seat_scope.operator, the `surface_intent` row), `frank/internal/engine/submit.go:340-343` (config_change member arms). Test `frank/test/fixtures/s8_registry_changeset_test.go`.
**Interfaces — Produces:** `config_member` enum gains `adoption` (reserved, never loadable) + `catalog`; `member.seat_scope.operator` gains both; `surface_intent` loses its static `required_when {layer_present:observe}` + `visible_when` (derived at step-4.5, `progress ⇐ otherwise`, non-gate only).
- [ ] Write failing FX-CFG-10 (fieldspec transition relation — added-key/type-change/rollback/skip ⇒ typed `config-version-transition` reject; value-only + lawful `v5→successor` ⇒ accept).
- [ ] Write failing token-fill-gate test (a LIVE operator `member: catalog`/`adoption` submission passes fill-validation ONLY with both `:85`+`:136` tokens present; absent either ⇒ fill-validation bounce).
- [ ] Apply the changeset bytes; land as the governed `v5→successor` transition.
- [ ] Run the changeset tests → PASS; INV-CATALOG green (the census move is an amendment, not drift); commit.

## Milestone B — the observe spine (vertical slice, thinnest end-to-end)

### Task 3: `observe_gate()` in the serialized commit path
**Discharges:** `s8-design-m3-registry` §2/§3.1/§3.3 (observe-as-send, observer-only allowlist, disposition→token). **The vertical-slice keystone.**
**Files:** Create `frank/internal/observe/gate.go`; Modify `frank/internal/engine/submit.go` (insert post-lineage `:67` / post-`slot_in`, pre-intents `:98`, inside the serialized loop). Test `frank/test/fixtures/s8_observe_spine_test.go`.
**Interfaces — Produces:** `observe.Gate(cand, seat, phase, authority, env) (ObserveResult, terminal)`; writes ONLY the §3.1 allowlisted fields; performs NO delivery effect; the hook reads (never writes) `slot_in`/`authority_class`; runs the phase-shaped done-predicate; on pass → `accepted`, on observably-false predicate → `rejected` naming the failing field. **Consumes:** T1 `PresentLayers` (gate active iff `observe` present).
- [ ] Write failing spine test: an observe-active submit with a passing read-only predicate ⇒ `accepted` + observed stamps; a false action-claim (no matching git ref) ⇒ `rejected` naming the predicate, candidate NOT delivered, a terminal evidenced record committed.
- [ ] Write failing allowlist test: the hook writing any envelope/identity field (FROM/ROLE/PARENT/relay_id) is a build-time/impl violation (positive allowlist enforced).
- [ ] Implement the hook inside the atomic commit path (TOCTOU-closed: observe runs inside `submit()`); wire §3.3 disposition→token.
- [ ] Run the spine tests → PASS; commit.

### Task 4: the E1 read-check family + §6.1 I-PH param hygiene
**Discharges:** `s8-design-m3-registry` §3 (E1 {read-file, git-status}) + §6.1 (I-PH).
**Files:** Create `frank/internal/observe/registry.go` (`CheckEntry` + selection + param validation), `frank/internal/observe/checks_base.go`. Test `frank/test/fixtures/s8_check_registry_e1_test.go`.
**Interfaces — Produces:** `CheckEntry{id,rung,class,executor_required,param_schema,produces,timeout_class}`; `read-file{path(lane-scoped,normalized;abs+`..` refused),expect:line|hash|schema_ref(registry id)}` → E1; `git-status{lane_ref}` → E1 (clean/dirty from outside, DI-5). **Consumes:** T3 gate.
- [ ] Write failing I-PH tests: `read-file` with an absolute or `..` path ⇒ **refused at validation** (pre-spawn); `run-suite.target` non-enum ⇒ refused; a verdict/`failing_detail` carrying a store/config/outbox/socket path or config value ⇒ redacted/refused.
- [ ] Write failing E1 tests: `read-file` reads a cited `file:line`; `git-status` observes clean vs dirty; a clean-tree-claimed/dirty-observed ⇒ veto.
- [ ] Implement the registry + the two base checks + conductor-side param normalization/refusal.
- [ ] Run E1 + I-PH tests → PASS; I-PH sweep over the new surface; commit.

### Task 5: the executor host + E2 `run-suite` (m-7 executor lock)
**Discharges:** `s8-design-m7-executor` r4 (spawn/CheckVerdict split, workdir, §2.7 kill, FX-EXE-1..6); `s8-design-m3-registry` §3/§5/§7. **m-7 build-fidelity co-owns the executor package.**
**Files:** Create `frank/internal/executor/executor.go`; wire from `frank/internal/observe/registry.go` (suite-class dispatch). Test `frank/test/fixtures/s8_executor_test.go`.
**Interfaces — Produces:** `spawn(check_id, params, lane_ref, timeout) → CheckVerdict{check_id,claim_ref,outcome,rung_reached,predicate,timing,failing_detail}` (closed shape — NO `RunResult`/output-ref/path text crosses); per-run `os.MkdirTemp` workdir (the only writable surface); env = the §2.3 allowlist; §2.7 kill: SIGKILL `-pgid` → reap → verify `kill(-pgid,0)==ESRCH` (bounded) → cleanup; survivor ⇒ typed machinery fault, **workdir PRESERVED**; cleanup exit-confirmed-only. **Consumes:** T4 registry (only `class:suite`, `executor_required` entries reach spawn).
- [ ] Write failing FX-EXE-1 (NF isolation probe — zero store/config/outbox/socket handles or path text, zero credentials/signing, env == §2.3 allowlist; the same-uid ambient residual reported separately, never "by construction").
- [ ] Write failing FX-EXE-2 (kill state machine — group SIGKILL · reap · bounded ESRCH · then cleanup; survivor ⇒ typed machinery fault, workdir preserved).
- [ ] Write failing FX-EXE-3/4/5/6 (exactly-once verdict; output ceiling truncate+mark; no-verdict ⇒ NF-S6 disposition; boundary byte-shape == closed `CheckVerdict` + `side_effecting`→typed host refusal).
- [ ] Implement the executor + `run-suite{target(registry enum),expect_green}` E2 check wired to the gate.
- [ ] Run FX-EXE-1..6 → PASS; commit.

### Task 6: verdict binding — conductor-computed observe fields
**Discharges:** `s8-design-m3-registry` §6 (verdict binding) + §6.1 (verdict redaction).
**Files:** Modify `frank/internal/observe/gate.go`. Test `frank/test/fixtures/s8_verdict_binding_test.go`.
**Interfaces — Produces:** the gate consumes `CheckVerdict` and computes `executable_claim_results` (row/claim), `achieved_evidence` (max rung reached by a passing check), `target_gap_result` (`EVIDENCE_TARGET` vs `achieved_evidence` → met|target_gt_achieved|not_applicable), `record_integrity` (rollup of per-field `evidence_integrity`). The executor writes nothing (no allowlist entry ⇒ forged write inert). **Consumes:** T5 `CheckVerdict`.
- [ ] Write failing tests: a passing E2 verdict ⇒ `achieved_evidence:E2` + a claim row; `EVIDENCE_TARGET:E2` met/gap → `target_gap_result`; a forged executor verdict asserting a field-write is **inert**.
- [ ] Implement the binding + the §6.1 conductor-side redaction pass on the returned verdict.
- [ ] Run → PASS; commit.

### Task 6.5: the m-2-locked staged pipeline — step-3 `authority_class` compute + step-4.5 profile-aware completeness + Option-B `surface_intent`
**Discharges:** `s8-design-m2-grammar` `:64-77`/`:91-94`/`:138-140` (the staged pipeline: step-3 `authority_class`, step-4 observe, step-4.5 completeness + `surface_intent` derivation). **The missing lock-bearing stage (plan-review F1).** Without it, decision-② (T7) reads an absent/lane-supplied `authority_class` and Option-B `surface_intent` is merely un-required rather than produced.
**Owner split (named):** **m-2** owns the `authority_class` formula (`record_kind ∈` authority-set ∨ `gate_category ∈ A`; never model) + the registry-row amendment (rides T2); **m-5/m-6** own the `surface_intent` producer/profile contract — **already locked** in the m-2 grammar §8:198-199 (`{progress, review_checkpoint, advisory, result}`, conductor-derived, TOTAL over non-gate-bearing, `progress ⇐ otherwise`; gate-bearing carry NONE), consumed not authored here; **m-7** hosts the step-4.5 pipeline stage; **m-3** consumes `authority_class` for disposition (T7) + observe-output completeness.
**Files:** Modify `frank/internal/engine/submit.go` (step-3 compute, after `slot_in`/gate-category classify, before the observe hook); Create `frank/internal/engine/completeness.go` (step-4.5 stage, m-7-hosted). Test `frank/test/fixtures/s8_pipeline_stages_test.go`.
**Interfaces — Produces:** `authority_class` computed at step-3 (system/computed, pre-disposition) — read by T3/T7, written by neither m-3 nor any lane; `surface_intent` derived at step-4.5 for **non-gate-bearing** records only (default `progress`), **absent** on gate-bearing records; the step-4.5 stage validates producer manifests before append. **Consumes:** T3–T6 observe outputs (step-4.5 runs after the observe manifest); the T2 registry amendment (Option B row). **Rail A:** `surface_intent` is **additive/open** (an old reader ignoring it loses only presentation detail); `authority_class` gating is **closed/fail-closed** (ignore-unknown would change the meaning of the decision-② acceptance). **Rail B:** pass — drift/history-truth/provenance-grade, no adversarial-only mechanism.
- [ ] Write failing test: `authority_class` is computed at step-3 from `record_kind`/`gate_category` (not model, not lane-supplied); T7 reads the computed value.
- [ ] Write failing test: step-4.5 derives `surface_intent` for a non-gate-bearing record (default `progress`), and it is **absent** on a gate-bearing record; a producer writing outside its manifest ⇒ rejected.
- [ ] Write failing **oracle test (F1 item 4):** m-3 **cannot write** `authority_class` or `surface_intent` (channel-keyed guard, T8 class); the committed behavior matches the m-2 fixture oracle.
- [ ] Implement the step-3 compute + the step-4.5 completeness stage + the Option-B derivation.
- [ ] Run → PASS; commit.

## Milestone C — dispositions + guard + activation

### Task 7: decision-② + the timeout machinery-fault edge
**Discharges:** `s8-design-m3-registry` §3.2(d)/§3.3/§4 (class-conditional fail-closed; timeout = machinery-fault, never `accepted`+`self_reported`); `s8-design-m7-executor` NF-S6.
**Files:** Modify `frank/internal/observe/gate.go`. Test `frank/test/fixtures/s8_decision2_test.go`.
**Consumes:** the **T6.5** step-3 computed `authority_class` (this task no longer assumes it exists). **Interfaces — Produces:** keyed on computed `authority_class` (m-2 step-3 pre-disposition; `record_kind ∈` authority-set ∨ `gate_category ∈ A`; never model): authority ∧ `record_integrity ∈ {self_reported, mixed}` ⇒ terminal `held`+escalate; non-authority + unobservable ⇒ `accepted`+`self_reported`+`degradation_notes`; observed-fail ⇒ `rejected` both classes; timeout/machinery-fault ⇒ authority `held` / non-authority `rejected`/author-return (**never** `accepted`+label — that is the separate no-vantage row).
- [ ] Write failing decision-② family: authority+unobservable → `held`+escalate (incl. the `mixed` leg explicitly); non-authority+unobservable → `accepted`+`self_reported`+`degradation_notes`; observed-fail → `rejected` both.
- [ ] Write failing timeout-≠-no-vantage pair: a killed check → authority `held` / non-authority `rejected`+fault-edge; paired with a distinct no-vantage → `accepted`+`self_reported`.
- [ ] Implement the class-conditional disposition + the timeout fault edge.
- [ ] Run → PASS; commit.

### Task 8: the s5-b (h) channel-keyed suppliability guard
**Discharges:** `s8-design-m3-registry` §3.1/§5.1 (non-lane-writability; forged observed-field REJECT).
**Files:** Modify `frank/internal/engine/submit.go` (fill-validation) + `frank/internal/observe/gate.go`. Test `frank/test/fixtures/s8_suppliability_guard_test.go`.
**Interfaces — Produces:** a lane supplying any Block-A observe field (`achieved_evidence`/`executable_claim_results`/…) is typed-REJECTED, channel-keyed; only the conductor fills them.
- [ ] Write failing test: a lane-submitted `achieved_evidence`/`executable_claim_results` ⇒ typed REJECT (not silently dropped, not accepted).
- [ ] Implement the channel-keyed guard.
- [ ] Run → PASS; commit.

### Task 9: observe-activation `config_change` + fresh-genesis dogfood store + exit-test legs 1+2
**Discharges:** SEQ-1 steps 1–3; the kickoff exit-gate baseline legs 1+2; the dogfood relaunch.
**Files:** the operator `config_change` fixture path; `frank/internal/store/genesis.go`. Test `frank/test/fixtures/s8_exit_gate_test.go`; ops note in the slice trail.
**Interfaces — Consumes:** T1–T8. **Produces:** a fresh-genesis store, activated to `observe:true` by an operator `config_change` (restart-effective); the dogfood relaunch rides this store.
- [ ] Write **exit-test leg 1** (a false "done" observe-as-send is **rejected pre-send**, naming the failing predicate — candidate never reaches a recipient).
- [ ] Write **exit-test leg 2** (a passing send with **conductor-captured** evidence — `achieved_evidence`/`target_gap_result`/`record_integrity` stamped, observed).
- [ ] Genesis the dogfood store; land the activation `config_change`; restart; verify phase-0 + activation.
- [ ] Run legs 1+2 → PASS; start the dogfood relaunch on this store; name which S7A-TRAIL-FINDINGS classes the conductor kills; commit.

## Milestone D — migration path + hardening

### Task 10: the bless/adoption path (§5.1) — existing-store migration (built + fixtured, NOT the dogfood path)
**Discharges:** `s8-design-m7-config` r13 §5.1/§5.2 (offline bless, adoption record, singular-catalog) + FX-CFG-12/13/14/15. **m-7 build-fidelity + m-2 tokens; SEQ-1 fill-gate exposure confirmed.**
**Files:** Modify `frank/internal/store/genesis.go` (offline bless mode), `frank/internal/engine/submit.go` (post-adoption singular `member: catalog` arm). Test `frank/test/fixtures/s8_adoption_test.go`.
**Interfaces — Produces:** the offline bless bootstrap writes an `adoption` record (body `members[].name ∈ {catalog, engine}`, base64, name-sorted, provenance `channel: bless`) directly (pre-serve, NOT fill-validated); post-adoption a singular `config_change` gains a state-aware `member: catalog` arm.
- [ ] Write failing FX-CFG-12 (two-member store → offline bless → restart full reader → three-member chain digest + capability checks pass).
- [ ] Write failing FX-CFG-13/14/15 (adoption `bytes_b64` == materialized members; exact-set/order/cardinality rejects; PRE-adoption `member: adoption`/`catalog` ⇒ typed reject at submit; post-adoption singular catalog accepts).
- [ ] Implement the bless mode + the singular-catalog arm.
- [ ] Run FX-CFG-12..15 → PASS; commit.

### Task 11: the adversarial fixture set + I-PH sweep + INV-CATALOG green
**Discharges:** `s8-design-m3-registry` §8 (full negative set); the kickoff exit-gate promoted families; `STEP-2-KICKOFF.md:52` (I-PH).
**Files:** `frank/test/fixtures/s8_adversarial_test.go`; `frank/test/invariants/` (unchanged-green). 
**Interfaces — Consumes:** T3–T10.
- [ ] Fabricated-done family: forged observed-field REJECT (T8); `EVIDENCE_TARGET:E2` failing suite → `rejected` naming the predicate; clean-tree/dirty veto; `slot_in` immutability / re-tag-to-escape leg.
- [ ] Executor-isolation NF (FX-EXE-1, m-7-hosted) + the I-PH negative fixture (§6.1) + the isolation-probe (provided-surface absence, residual stated).
- [ ] Run the full serialized battery FILE-captured, sequence-honest; the ten INV-CATALOG laws green, uncached; I-PH swept over every new seat-visible surface; commit.

---

## Acceptance criteria (slice exit)

1. **Exit-test legs 1+2 live + green** (T9): a false "done" rejected pre-send; a passing send with conductor-captured evidence.
2. **All FX-EXE-1..6 + the used FX-CFG-1/2/7/10/12..15 + the m-3 §8 negative set green**, uncached, serialized, FILE-captured.
3. **The dogfood relaunch runs on the fresh-genesis spine store** and names the S7A-TRAIL-FINDINGS classes the conductor kills.
4. **The ten INV-CATALOG laws stay green** across every changeset (amendments, not drift).
5. **No locked contract altered** — the four locks are consumed byte-exact; any needed change routes through the owning pair's amendment path + master (delegation condition (c)).
6. **The interim defaults are labeled interim in-code + in-trail** with the s10 sunset noted (Global Constraints).
7. **The staged pipeline is live + oracle-proven (T6.5):** `authority_class` is computed at step-3 (decision-② reads a computed value, not absent/lane-supplied); step-4.5 derives Option-B `surface_intent` (non-gate-only, `progress ⇐ otherwise`); m-3 **cannot write** either field, matching the m-2 fixture oracle.
8. **Every new surface carries its Rail-A open/closed choice + a Rail-B pass** (build-time review criterion, `…-163144`).

## Out of scope (OUT — do not build in s8)
The away-bridge / live egress-chokepoint activation (first external send / Step-6) · E3/E4 · the OS sandbox (v1 handle-surface ceiling; shelved) · the m-6 comms mechanism / 7-state FSM / wake-on-reply / ODB render / the live operator gate prompt (s10) · m-5 concrete tag-space + invariant selection + ceilings (c3-reserved — s8 ships base predicates + the opaque `slot_in` mechanism only) · interjection/steer (Step-3) · benchmark loop · engine perf (s6 ledger exonerated it) · **the s9 fuzzy-claim adjudication rung** (independent cheap-model judging, egress-fenced — m-3-owned s9; `…-163144` addendum 3; s8 keeps the check-registry seam open per Rail A but adds no adjudication entry) · **activating the bless/adoption path against the actual shipped live store** (the public-release-migration moment — a NAMED CARRY; s8 builds + fixtures the mechanism at T10 but does not migrate the live store).

## Boundary contract (cross-domain, this slice)
- **m-7** rides as build-fidelity on its two locks at the integration points named (T1 config, T5 executor, T10 bless) — by relay, sole-writer per seat.
- **m-2** executes the registry changeset (T2) + owns the `:85`/`:136` tokens + the fieldspec forward-transition, inside the slice.
- **m-3** guides the observe slices (T3/T4/T6/T7/T8/T9/T11).
- Cross-domain shape changes (m-1 store semantics, m-6 surfaces) → escalate to master (delegation condition (d)); B10 (domain pair as build pair) is NOT in effect for s8.

## Team shape (ratified staffing — the PLAN names the seats; the operator boots the sessions)
A fresh slice team on the s1..s6 slice-team pattern, sole-writer discipline per seat:
- **`s8.lead`** (m-3-guided) — the observe-spine seats (gate/registry/checks/verdict/disposition/guard/exit-gate/adversarial: T3,4,6,7,8,9,11).
- **`s8.config`** (m-7 build-fidelity) — config knob + genesis + executor + bless (T1,5,10).
- **`s8.registry`** (m-2) — the registry changeset (T2) + the tokens/transition.
- **`s8.reviewer`** — the slice's adversarial build-reviewer (per-task gate).
Each pair-internal leg uses a unique sub-DISPATCH_ID under the `s8-plan-m3` chain; the operator boots these sessions.

## Delegation conditions I operate under (from the dispatch, binding on the token I issue)
(a) unique sub-DISPATCH_IDs per pair-internal leg; (b) registry/schema changes beyond the scheduled T2 changeset → escalate to master BEFORE work; (c) any change to a locked design contract → the owning pair via the amendment path + master, never a silent edit; (d) cross-domain shape changes (m-1/m-6) → master; (e) merge is operator-only — the merge-decision relay goes TO the operator, the executor acts only on a `HUMAN_MERGE_AUTHORIZATION` grant; (f) a blocked/model-stopped lane reports and holds — no pushing through a red; (g) the dogfood evaluation inherits S7A-TRAIL-FINDINGS (name the killed classes).

## Self-review (writing-plans checklist, run by me)
- **Lock coverage:** every EFFECTIVE lock maps to ≥1 task — m3-registry (T3,4,6,7,8,11), m7-executor (T5,11), m7-config (T1,2,9,10), m2-grammar (T1,2,**T6.5** — the step-3 `authority_class` + step-4.5 completeness + Option-B `surface_intent` staged pipeline, r2/F1). SEQ-1 answered explicitly (steps 1–4 + dogfood store), operator-steered fresh-genesis-first, m-7 half confirmed / m-2 pending. Rails A/B applied per new surface; s9 seam kept open. ✓
- **Placeholder scan:** no "TBD"/"handle edge cases" — every task cites exact files (`file:line`), fixture ids, and interface signatures. ✓
- **Type/name consistency:** `CheckEntry`/`CheckVerdict`/`observe.Gate`/`present_layers`/`PresentLayers`/`authority_class` used identically across tasks; the `CheckVerdict` shape in T5 matches the T6 consumer + the m-7 lock's closed shape. ✓
- **Open dependency:** SEQ-1's fill-gate exposure + order-of-record await the m-7/m-2 COORD confirm (a named plan input, not a silent assumption) — flagged as the one pre-T2/T10 gate.
