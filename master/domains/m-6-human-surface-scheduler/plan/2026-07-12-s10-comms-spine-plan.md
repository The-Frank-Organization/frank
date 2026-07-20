# s10 Comms-Spine Implementation Plan (the minimum A-gate path end-to-end)

> **For agentic workers (the s10 build pair on `frank`):** implement task-by-task against `main@8941889` (`s8-close`). Steps use checkbox (`- [ ]`) syntax. This plan is authored **off-frank** by the m-6 guide seat: it fixes the **contracts, task boundaries, acceptance criteria, and RED-first negatives**; the literal Go + exact test bodies are written by `s10.implementer` against the live tree, gated by `s10.planner`'s task review and the delegated-dispatch conditions. Where a step names a primitive (`completePark`, `gate_resolution`, the observe gate, the A-2 replay machinery) it is a **lift** of an already-built s5/s8 mechanism — extend, never re-invent.

**Goal:** Build the minimum A-gate human-surface path end-to-end — a conductor-generated ODB parks a lane, the operator replies on a validated channel, the lane re-observes locally and wakes exactly once, with deterministic crash-safe resummons — proving ROADMAP exit leg 3 (*a parked lane wakes on reply*) live on a fresh dogfood store.

**Architecture:** Lift the s5-built park/wake primitives (`completePark`, wake-on-operator-verdict, durable-inbox liveness, crash-safe re-issued wake) into the locked m-6 c3 mechanism — the FSM spine subset, the ODB render/capture contract, J1 throughout. Two new **interpreter-bearing** `record_kind` tokens (`odb`, `resummon_command`) ride ONE governed fieldspec transition (the Q6×Q4 MAJOR-but-safe path, m-2 executes). Bucket A only; B/C/D projections, the elaborate-more fork, and 8a freeze/re-issue are s11.

**Tech Stack:** Go (`frank/`); the s8 observed store + governed supply + the serialized `submit()` commit loop; the m-7-hosted scheduler/timer + A-2 content-hash replay; the m-3 observe gate.

## Global Constraints (every task's requirements implicitly include these)
- **Byte-exact terminal enum** `{accepted, rejected, held}` — parking is `accepted`; `held` is the fault/fail-closed lane; `rejected`→bucket D. Never re-spell.
- **R2 — the model is never a gate input.** No new gate keys on a model identity. The ⑤ model-name field is payload/render only.
- **Confusion-firewall Rails (build-time review criterion, every new surface):** **Rail A** — additive/open where ignoring an unknown loses only detail; closed/fail-closed where ignore-unknown changes the MEANING of acceptance. **Rail B** — cut by FUNCTION not flavor; drift/crash-safety/history-truth/provenance machinery stays, claims are confusion-graded, no home-grown security primitive. Each task states its Rail-A open/closed choice + a Rail-B pass inline.
- **I-PH** — no store/config/socket/outbox path is ever seat-visible on any new surface (ODB render, FSM/resummon outputs). Swept in T11.
- **Egress stays fixture-scoped.** s10 comms are local-only; the live chokepoint stays dormant. The ⑤ ODB model-name field is carried typed+exempt-marked but its egress fixture pair registration is **s9/s11**, not s10.
- **No migrator; committed history is never re-classified.** The v7→v8 widening is reader-first, forward-only; rollback/skip rejected.
- **Locks are consumed, not re-designed.** Any locked-contract change routes to the owning pair via the amendment path + master (delegation condition c); this plan re-designs nothing.
- **The ONE named schema transition** is T1. Any registry/schema change beyond it → master before work (delegation condition b).

---

## The two named questions — ANSWERED (owner-confirm legs named)

**SEQ-2 — which store the spine's dogfood rides → FRESH GENESIS at the widened v8 schema.**
Rationale: the confusion-firewall directive steers a fresh store for the thin-slice proof (kickoff §6/§9), and s8's SEQ-1 already set fresh-genesis-first as the ratified default. A store *born at v8* needs no store-level transition and makes PARK-ACROSS-V8 vacuous for the live dogfood (no pre-v8 parked gates can exist). **Fixture consequence:** the governed-transition machinery is still exercised — by T1's registry transition (m-2 executes v7→v8) and by T1's old-reader typed-refusal RED fixture (its third live lap) — but NOT by the dogfood store. Owner-confirm leg: m-2 (transition executor) + m-7 (capability marker) confirm the fresh-v8 genesis rides their §7-apply/bless pattern.

**PARK-ACROSS-V8 — the slice's own transition vs its own parked gates → AVOIDED BY SEQUENCING.**
The spine builds parking AND widens the schema, so a gate parked across the v8 bump is the 8a class whose FULL contract (freeze/re-issue) is **s11 scope — NOT built here** (dispatch explicit). s10 avoids it structurally: **T1 (transition) + T2 (fresh v8 store) land before any park fixture (T4+) opens**, so no gate is ever parked under v7 within the slice, and the store never re-transitions mid-slice. **Honest residual (recorded, not built):** if a gate were parked under v7 and a bump landed mid-slice — which the sequencing prevents — the co-signed 8a floor applies (migrate-then-validate; un-migratable → `held` + escalate with `stale_schema`, never dropped). That floor is the **s11** contract; s10 states it and does not build the freeze/re-issue branch.

---

## File / seam map (responsibilities; exact Go paths are `s10.implementer`'s against the tree)
- **`registry.json` + m-2's successor map (v7→v8):** the two new `record_kind` tokens + per-kind required fields + seat_scope rows. **m-2 executes** (T1).
- **m-7 fieldspec capability exact-set:** gains the v8 marker; the marker-first phase-0 refusal of a pre-v8 reader. **m-7 owns** (T1).
- **the ODB render surface (m-6):** the `odb` record shape + the `agent_enum_pick` choice-set render (T3).
- **the park/wake FSM (m-6, lifting s5 primitives):** `completePark` → `parked_waiting_human`; wake-on-operator-verdict → `replied_pending_validation` → `resumed`; crash-safe re-issued wake (T4/T5/T7).
- **the observe-gate re-fire seam (m-3 fidelity):** the s8 observe gate re-fired locally on resume (T6).
- **the scheduler/timer + A-2 replay (m-7 build-fidelity):** the `resummon_command` record, content-hash keyed, deduped through A-2 (T8).
- **the m-3 check-policy hook (sunsets):** the long-run disposition + the `side_effecting` gate now call the m-6 park+ODB+verdict path (T9/T10).

---

### Task 1: The v7→v8 governed fieldspec transition (the Q6×Q4 MAJOR-but-safe path)

**Owner:** m-2 executes the registry bytes (the s8 T2 owner-fidelity pattern); m-7 executes the capability-set marker move (theirs alone). m-6 guides; consumes the tokens.
**Files:** Modify `registry.json` (+ m-2's successor map) — new `record_kind` members + their per-kind required fields + seat_scope rows; Modify m-7's fieldspec capability exact-set; Test: the old-reader typed-refusal fixture (FX-CFG-10/claim-input pattern, third lap).
**Interfaces:**
- Produces: `record_kind` members **`odb`** (conductor-authored; obligates render/park/bucket machinery) and **`resummon_command`** (conductor-authored; obligates scheduler interpretation + A-2 dedupe) — *proposed bytes; m-2 makes the final per-token byte + required-field call at the transition.* Seat-scope: conductor-authored (the `seat_mint`/`gate_resolution` precedent).
- Consumes: nothing (slice-opening task).
- **Rail A: CLOSED** — these are interpreter-bearing kinds; an old reader accepting-but-not-interpreting is the fail-OPEN class, so the widening is MAJOR and gated marker-first. **Rail B: pass** — no security primitive; the marker is history-truth/capability machinery.

- [ ] **Step 1 (RED):** Write the old-reader refusal fixture — a pre-v8 (v7-capability) reader opening a v8 store REFUSES at phase-0, **zero partial interpretation**. Expected: FAIL (marker move not yet made).
- [ ] **Step 2:** m-2 lands the two tokens + per-kind required fields + seat_scope rows as registry bytes over the successor map (reader-first, forward-only; rollback/skip rejected). No migrator; committed history untouched.
- [ ] **Step 3:** m-7 adds the v8 marker to the fieldspec capability exact-set (the marker-first phase-0 ceiling).
- [ ] **Step 4 (GREEN):** the old-reader refusal fixture passes; a v8 reader accepts the v8 store; the A-1 stale-form `re-render` story covers any live form. Confirm no other registry/schema change rode along (delegation condition b).
- [ ] **Step 5:** Commit. Fidelity: m-2 confirms the transition bytes; m-7 confirms the capability move.

### Task 2: Fresh-genesis dogfood store at v8 (SEQ-2)

**Files:** the slice's dogfood store genesis (the §7-apply/bless pattern, a fresh blessed store).
**Interfaces:** Consumes: T1 (the v8 schema). Produces: the fresh v8 dogfood store the spine rides. **Rail A: open** (genesis is additive). **Rail B: pass.**

- [ ] **Step 1:** Genesis a fresh store at the widened v8 schema (fresh-store steer; the store is born at v8 → PARK-ACROSS-V8 vacuous for the dogfood).
- [ ] **Step 2:** Confirm no pre-v8 records exist in the store (the sequencing invariant that keeps 8a out of s10).
- [ ] **Step 3:** Commit. The dogfood evaluation extends the s8 ledger (delegation condition g — one series).

### Task 3: Minimal ODB render (the `odb` record)

**Files:** the ODB render surface (m-6); Test: the choice-set validation fixture.
**Interfaces:**
- Consumes: T1 (`odb` token), T2 (store); the promoted m-3 7-field bundle (`completed_proof`/`record_integrity` shown) + the envelope/provenance fields.
- Produces: a conductor-generated `record_kind=odb` record carrying (a) the operator-facing render fields and (b) the **bounded choice-set as `agent_enum_pick` buttons** (the operator can pick only a legal verdict). The choice-set is **frozen at emit** (the schema-facing companion to J1 "never re-ask a stale gate").
- **Rail A: CLOSED** on the choice-set (an out-of-enum choice must be rejected, not ignored — ignoring changes the meaning of "the operator picked a legal verdict"). **Rail B: pass.**

- [ ] **Step 1 (RED):** an ODB whose captured choice is outside the frozen `agent_enum_pick` set → REJECTED. Expected: FAIL.
- [ ] **Step 2:** Render the minimal ODB: the 7-field bundle + `agent_enum_pick` choice buttons; carry the ⑤ model-name field **typed + exempt-marked** (payload only; egress dormant; R2 untouched — it is never a gate input).
- [ ] **Step 3 (GREEN):** legal choice accepted; out-of-enum rejected. I-PH: no store/config path in the render.
- [ ] **Step 4:** Commit.

### Task 4: Park the A-gate (lift `completePark`)

**Interfaces:** Consumes: T3 (ODB). Produces: FSM `active → parked_waiting_human` on an accepted A-gate. **Rail A: closed** (only A-gates park; a mis-classified non-A must not silently park). **Rail B: pass.**

- [ ] **Step 1 (RED):** a non-A record (no A-set `gate_category`, no `HUMAN_GATE_REQUIRED`) does NOT park (B/C/D are s11). Expected: FAIL.
- [ ] **Step 2:** On an `accepted` A-gate, emit the ODB (T3) and call the s5 `completePark` primitive → `parked_waiting_human`. Lift, do not re-invent.
- [ ] **Step 3 (GREEN):** A-gate parks + emits ODB; non-A does not park.
- [ ] **Step 4:** Commit.

### Task 5: Validated operator reply + capture (operator-FROM verdict)

**Interfaces:** Consumes: T4 (parked gate). Produces: FSM `parked_waiting_human → replied_pending_validation`; a validated operator-FROM verdict. **Rail A: closed** (a non-operator or out-of-set reply must be rejected). **Rail B: pass.**

- [ ] **Step 1 (RED):** three negatives — (a) a non-operator-FROM reply → rejected; (b) a picked choice ∉ the frozen set → rejected; (c) a second resolution of an already-resolved gate → rejected by the single-resolution guard. Expected: FAIL.
- [ ] **Step 2:** Lift the s5 wake-on-operator-verdict (operator-FROM only, single-resolution guard). Validate: operator-FROM channel + form/lineage + choice ∈ frozen set → `replied_pending_validation`.
- [ ] **Step 3 (GREEN):** valid operator verdict advances the FSM; all three negatives reject.
- [ ] **Step 4:** Commit.

### Task 6: Local re-observe on wake (m-3 fidelity; annotation A)

**Interfaces:** Consumes: T5 (validated reply). Produces: FSM `replied_pending_validation → resumed` (reply appended + delivered + observe re-fired). **Rail A: closed** (a stale-done wake must not proceed as-if-done). **Rail B: pass.**

- [ ] **Step 1 (RED):** on wake, if the re-fired observe gate finds the ground-truth done-state now FAILS, the verdict does NOT silently apply — the lane re-summons / holds (J1 never-auto-approve). Expected: FAIL.
- [ ] **Step 2:** Re-fire the **s8 observe gate locally** on the woken record's resume (newest-authoritative; refresh-before-resummon). **No away-token machinery** (annotation A; decision-④'s token re-observe stays step-(d)). On observe-pass → `resumed`.
- [ ] **Step 3 (GREEN):** clean re-observe resumes; stale-done re-summons. **m-3 fidelity** confirms the J1/refresh contract against the landed mechanism (their named edge).
- [ ] **Step 4:** Commit. *(Choice-set staleness across a schema bump — the 8a case — is out of scope here; T6 is the observe-gate re-fire only, per PARK-ACROSS-V8.)*

### Task 7: Exactly-once wake (crash-safe; lift the s5 re-issued wake)

**Interfaces:** Consumes: T4–T6. Produces: the exactly-once wake guarantee (EXIT LEG 3's crash companion). **Rail A: closed** (crash-safety is meaning-bearing). **Rail B: pass** (crash-safety machinery stays — Rail B keeps it).

- [ ] **Step 1 (RED):** crash mid-park and crash mid-wake → on recovery, **exactly one** wake, no stranded lane, no double-delivery. Expected: FAIL.
- [ ] **Step 2:** Lift the s5 durable-inbox liveness + crash-safe re-issued wake; extend the liveness contract to cover the resummon timers (T8).
- [ ] **Step 3 (GREEN):** both crash legs yield exactly-once wake.
- [ ] **Step 4:** Commit.

### Task 8: Deterministic resummon commands (m-7 build-fidelity; note ③)

**Interfaces:** Consumes: T4 (parked gate), T7 (wake machinery), T1 (`resummon_command` token). Produces: timer-fired resummon commands, content-hash keyed, A-2 deduped. **Rail A: closed** (a double-summon on crash-refire is meaning-bearing). **Rail B: pass.**

- [ ] **Step 1 (RED):** a crash-refire of a resummon with the SAME content-hash key (seat + decision + cadence-slot) → **deduped to one** command through the A-2 replay machinery, not double-summoned. Expected: FAIL.
- [ ] **Step 2:** Emit `resummon_command` records (interpreter-bearing) on the two G4 timers (no-response vs answered-but-stalled), content-hash keyed. **Escalate the SUMMON CHANNEL, never the verdict; no hard deadline** (J1/G4). Dedupe crash-refires through A-2.
- [ ] **Step 3 (GREEN):** timers fire resummons; crash-refire dedupes. **m-7 build-fidelity** confirms the scheduler/timer + A-2 dedupe + store seams.
- [ ] **Step 4:** Commit.

### Task 9: SUNSET 1 — timeout auto-kill → park + ALERT (joint m-6 + m-3)

**Binding:** operator sunset ruling (kickoff §37) — the silent 120s/5s auto-kill is testing-phase only and **MUST BE REMOVED by s10**.
**Interfaces:** Consumes: T3 (ODB), T4/T5 (park+verdict). **Rail A: closed** (a silent kill is a meaning-changing default). **Rail B: pass.**

- [ ] **Step 1 (RED):** a long-running check does NOT silently auto-kill — it **parks + ALERTS** the operator (channel escalation) with an ODB carrying `{kill, extend}` choices. Expected: FAIL (silent auto-kill still present).
- [ ] **Step 2:** m-3 re-points the long-run disposition hook to the m-6 park+ODB(kill/extend)+verdict path. **J1-composition (named now):** the hard ceiling stays as a **block-only fail-safe backstop behind the prompt** — it may only KILL (conservative/block), **never auto-extend/approve** (J1 never-auto-approve; the monotonic conservative block-ceiling of G4). The empty-allowlist floor is NOT sunset — only how approvals arrive.
- [ ] **Step 3 (GREEN):** long-run check parks+alerts; operator can extend; the backstop kills only on a safety bound with no operator response. The silent auto-kill is **gone**.
- [ ] **Step 4:** Commit. m-3 fidelity confirms the hook.

### Task 10: SUNSET 2 — static `side_effecting` gate → live prompt (m-6; m-3 hook)

**Binding:** operator sunset ruling (kickoff §37) — the static-config `side_effecting` gate MUST BE REMOVED by s10.
**Interfaces:** Consumes: T3/T4/T5 (the ODB+park+verdict path). **Rail A: closed.** **Rail B: pass.**

- [ ] **Step 1 (RED):** a `side_effecting` op with no static approval raises the **live prompt** (an ODB → operator verdict), NOT a static-config allow/deny. Expected: FAIL (static gate still present).
- [ ] **Step 2:** Replace the static `side_effecting` gate with the live-prompt ODB path; the operator verdict gates the op. The floor itself is unchanged — only the approval channel.
- [ ] **Step 3 (GREEN):** the static-only gate is **gone**; the live prompt gates side-effecting ops.
- [ ] **Step 4:** Commit.

### Task 11: EXIT LEG 3 + the hardened exit battery

**Interfaces:** Consumes: T1–T10. Produces: the step-exit leg-3 demonstration + the institutionalized battery.

- [ ] **Step 1:** End-to-end live on the fresh v8 dogfood store: A-gate → ODB → park → operator reply → local re-observe → **exactly-once wake** = **EXIT LEG 3 (a parked lane wakes on reply)**.
- [ ] **Step 2 (institutionalized s8 exit items, day-one):** the **two mechanical tables** — (i) consumption→supply reconciliation, (ii) diff→license fence-row reconciliation — both green; the **label→mechanism sweep + verify-the-summary-line** run in the review; **RED-first negatives** confirmed present per row/task; **sequence-honest FILE-captured** battery; named seams recorded for any same-file multi-task edits.
- [ ] **Step 3:** the two **sunset demonstrations**: the silent auto-kill is gone (T9); the static-only gate is gone (T10). *"s10 does not close while the silent auto-kill or the static-only gate survives."*
- [ ] **Step 4 (I-PH sweep):** no store/config/socket/outbox path is seat-visible on any new surface (ODB render, FSM/resummon outputs).
- [ ] **Step 5:** the merge-decision relay TO the operator (merge is operator-only; the executor acts only on a grant carrying `HUMAN_MERGE_AUTHORIZATION` at grant time — delegation condition e).

---

## Boundary contract (consumed / produced)
- **Consumes (locked, not reopened):** m-6 c3 design-of-record + folds; the Q6×Q4 resolution (interpreter-bearing tokens ride ONE transition, m-2 executes); the s5 park/wake primitives; the s8 observe gate + A-2 replay + `main@8941889` base; m-2's successor map + m-7's capability set; m-3's check-policy hook (sunsets).
- **Produces:** the `odb` + `resummon_command` `record_kind` tokens (m-2 bytes); the minimal ODB render/capture; the FSM spine subset with crash-safe exactly-once wake; the deterministic resummon commands; the two sunsets on live machinery; EXIT LEG 3.

## Out of scope (s11 / later — do NOT build)
- B/C/D bucket projections · the elaborate-more fork (OQ-2 posture) · the full 8a freeze/re-issue branch (only the sequencing-avoidance + the recorded co-signed floor here).
- the away-bridge / Seam-C token / decision-④ rotate+re-observe (step-(d)) · E3/E4 · egress live-chokepoint activation + the ⑤ fixture pair registration (s9/s11) · the TUI/email-client UX (Step-4).

## Acceptance criteria
1. EXIT LEG 3 passes live on the fresh v8 dogfood store (T11).
2. The v7→v8 widening is MAJOR-but-safe: old-reader refuses at phase-0 (RED fixture green), no migrator, history un-reclassified (T1).
3. Only A-gates park; parking is `accepted`, `held` is fault-only, `rejected`→D — byte-exact (T4/Global).
4. Operator reply validated on operator-FROM channel + choice ∈ frozen set + single-resolution guard (T5).
5. Local re-observe on wake (no away-token); stale-done does not auto-apply — m-3 fidelity confirms J1/refresh (T6).
6. Exactly-once crash-safe wake, both crash legs (T7).
7. Resummon commands deterministic + A-2 crash-refire deduped; escalate channel never verdict; no hard deadline — m-7 build-fidelity confirms (T8).
8. Both sunsets demonstrated gone; the block-only backstop never auto-approves (T9/T10).
9. Every task carries its Rail-A open/closed choice + a Rail-B pass; I-PH clean; the two mechanical tables green (Global/T11).
10. SEQ-2 (fresh v8 genesis) + PARK-ACROSS-V8 (sequencing-avoidance) executed as stated; the s11 8a branch not built.

## Self-review (run against the dispatch + kickoff r4 + Q6×Q4)
- **Spec coverage:** dispatch scope bullet → T3–T8; EXIT LEG 3 → T11; SEQ-2 → T2 + the named-question block; PARK-ACROSS-V8 → sequencing (T1/T2 before T4) + the named-question block; Q6×Q4 MAJOR path → T1; sunsets → T9/T10 (+ exit T11); institutionalized s8 items → Global + T11; team + delegation → the PLAN relay. No uncovered requirement.
- **Placeholder scan:** the only deferred literal is Go/test bodies — explicitly delegated to `s10.implementer` at the stated grain (off-frank guide plan); the exact `record_kind` bytes are m-2's per-token call at T1 (proposed `odb`/`resummon_command`). No TBD acceptance/contract.
- **Consistency:** FSM state names match m-6 c3 §4 (`parked_waiting_human`/`replied_pending_validation`/`resumed`); terminal tokens byte-exact; the resummon-command dedupe keys (seat+decision+cadence-slot) match m-7 note ③; re-observe = local observe-gate re-fire (annotation A), not decision-④'s token re-observe.
