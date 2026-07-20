# s11 Comms-Thicken Implementation Plan (the last Step-2 build slice)

> **For agentic workers (the fresh s11 build pair on `frank`):** implement task-by-task against `s9-close@d91fcfb`. Steps use checkbox (`- [ ]`). This plan is authored **off-frank** by the m-6 guide seat: it fixes contracts, task boundaries, acceptance criteria, and RED-first negatives; literal Go + exact test bodies are written by `s11.implementer` against the live tree, gated by `s11.planner`'s review (the fresh pair runs its OWN same-owner PLAN→review→token loop consuming this plan-of-record — the s9 Option-A model; the guide does not token the pair). Named primitives (`completePark`, `gate_resolution`, the resummon dedupe keys, the §9 read-time migrator) are **lifts** of already-built s5/s8/s9/s10 mechanisms — extend, never re-invent.

**Goal:** Thicken the s10 A-gate spine into the full m-6 human-surface mechanism — the B/C/D bucket projections, the complete 7-state FSM surface, the elaborate-more fork, 8a hardening, the bucket/fixture matrix + the ③ known-A NF fixture, the 9-item refactor, and the G4 resummon-cadence config surface — completing Step-2's comms so only the step-exit test remains.

**Architecture:** Build on `s9-close@d91fcfb` (observed store + v8 schema + the s10 comms spine + the s9 evidence layer). Purely sequential (s9 closed; no m-2 collision edge). Consume the locked m-6 c3 + folds; two build tasks (the fork, 8a hardening) and one design-cell task **lock only when their named design gate returns** — decompose now, lock on gate.

**Tech Stack:** Go (`frank/`); the s10 FSM spine + resummon dedupe; the s9 evidence layer + §9 read-time migrator registry; the m-7-hosted scheduler/store seams.

## Global Constraints (every task's requirements implicitly include these)
- **Byte-exact terminal enum** `{accepted, rejected, held}` — parking is `accepted`; `held` is the fault/fail-closed lane; `rejected`→bucket D. Never re-spell.
- **R2 — the model is never a gate input.** The ⑤ model-name field stays payload/render only.
- **Confusion-firewall Rails (build-time review criterion, every new surface):** Rail A — additive/open where ignoring an unknown loses only detail; closed/fail-closed where ignore-unknown changes the MEANING of acceptance. Rail B — cut by FUNCTION not flavor; no home-grown security primitive; claims confusion-graded. Each task states its Rail-A choice + Rail-B pass inline.
- **I-PH** — no store/config/socket/outbox path seat-visible on any new surface. Swept in T11.
- **Egress stays fixture-scoped.** s11 comms are local-only; the away-bridge (and the live `egress_blocked` trigger) is the deferred step-(d) carry — its FSM state + local-park-resummon semantics are defined + fixtured here, the away-mirror send that triggers it is NOT built.
- **B11 straight-through cadence IN EFFECT** (once-through build, review churn accepted at the end — operator ruling 2026-07-13). **B12 DECLINED** — design churn is licensed: the design gates take the rounds they need.
- **Locks are consumed, not re-designed.** Locked-contract change → the owning pair via the amendment path + master (condition c).
- **Fence-union (s9 refinement):** the plan fence STARTS from m-6's domain seams ∪ the standing cross-cutting set (`config.go`, `registry_test.go`, executor + `main.go` composition roots); if any mid-build amendment activates, its test/fixture seams reconcile into the fence AT activation — an owner return naming un-fenced loci is an **escalation trigger**, not an in-fence assertion.
- **Institutionalized (day-one exit items):** both mechanical tables (consumption→supply · diff→license), the label→mechanism sweep, verify-the-summary-line, RED-first per task, named seams for same-file multi-task edits; the catch ledger continues as one series (condition g).

## The named design gates — decompose now, LOCK on return
- **(g1) 8a joint review** (`s11-8a-joint-review/DESIGN-orchestrator-planner-20260714-023001`): FLOOR co-signed (wake = migrate-then-validate via the §9 read-time migrator, then current-form validate; un-migratable → `held`/escalated, never dropped/auto-resolved). **Three open members gate T6's lock:** (m1) the `stale_schema` reason token (m-2 confirms into the bounce/reason vocab); (m2) the choice-set FROZEN at park + the migrator's frozen-choice interaction (m-2 confirms); (m3) the choice-set-changed ⇒ bounce-as-stale + RE-ISSUE branch (m-6.implementer reviews — its exactly-once story, its relation to the s10 resummon dedupe keys [new decision identity vs same], never-dropped-under-crash). **T6 locks only on the co-signed 8a contract returning to master.**
- **(g2) OQ-2 fork ceiling** (`s11-oq2-ceiling`, m-5 RULED, awaiting m-5.implementer review → completion to master): the fork is a `sensor`-class read-only spawn — ceiling `{write:read_only, dispatch:none, tool:read}`, external_send:none, posture `interactive`, emit `advisory` (never gate-bearing, §7.1 integrity split), non-interrupting, dies-on-verdict, parked lane byte-unmodified. **T5 (the fork mechanism = m-6's plumbing half) locks only on the g2 ceiling completing to master.** Nothing else in s11 waits on g2.
- **(dc) the s11 design cell** (m-3+m-6, J1-adjacent, its own ritual before T10 locks): non-`accepted`-resolution / blocked-prompter re-prompt semantics + the claimless authority-floor `held` edge at resolution. **T10 locks only on the m-3+m-6 design-cell ritual returning.**

## File / seam map (responsibilities; exact Go paths are `s11.implementer`'s)
- **the bucket projections (m-6):** the B/C/D saved-query-over-tags projections + terminal-token→bucket routing (T1–T3).
- **the FSM surface (m-6, lifting the s10 spine):** `bounced_repair` + `egress_blocked` states completing the 7-state machine (T4).
- **the elaborate-more fork (m-6 plumbing / m-5 ceiling):** the read-only context-fork spawn/feed/retire under the g2 ceiling (T5).
- **8a hardening (m-6 + m-2 + m-6.implementer via g1):** freeze-at-park + migrate-then-validate + re-issue (T6).
- **the fixture matrix + ③ NF (m-6 + m-3):** the A/B/C/D × terminal × failing_edge matrix + the ③ known-A NF fixture (T7).
- **the 9-item cleanup refactor (m-6, behavior-preserving):** the s10 cleanup card (T8).
- **the G4 config surface (m-6):** resummon cadence interim defaults → operator-config (T9).
- **the re-prompt / claimless-held edge (m-6 + m-3 via dc):** T10.
- **standing cross-cutting:** `config.go`, `registry_test.go`, executor + `main.go` composition roots.

---

### Task 1: Bucket B projection (orchestrator-absorbed; non-interrupting digest)
**Interfaces:** Consumes: the s10 A spine + `gate_category`. Produces: bucket B = saved query over `gate_category ∈ §J2 B-set`. **Rail A: closed** (a mis-absorbed A must promote, not stay B). **Rail B: pass.**
- [ ] **Step 1 (RED):** a B-category record lands in the **live local digest, non-interrupting** (G3), NOT the operator decision queue; a monotonic raiser (RAISE-ONLY ③) promotes B→A. Expected: FAIL.
- [ ] **Step 2:** Implement bucket B as a saved-query projection over tags (not an exclusive container); locked writer = `gate_category ∈ §J2 B-set`.
- [ ] **Step 3 (GREEN):** B-absorbed non-interrupting; promotion path works. **Step 4:** Commit.

### Task 2: Bucket C projection (CC-FYI)
**Interfaces:** Consumes: the addressing graph (operator on CC vs TO). Produces: bucket C = low-priority FYI feed. **Rail A: open** (informational; ignoring loses only detail). **Rail B: pass.**
- [ ] **Step 1 (RED):** an operator-**CC** (not TO) record → bucket C, **no decision obligation**; batchable. Expected: FAIL.
- [ ] **Step 2:** Implement bucket C as the CC-FYI saved query. **Step 3 (GREEN):** CC creates no obligation. **Step 4:** Commit.

### Task 3: Bucket D projection (observe-bounce, author-facing)
**Interfaces:** Consumes: the terminal-token→bucket map (`rejected`→D) + `failing_edge`. Produces: bucket D = author-return. **Rail A: closed** (a bounce must return to the author, not vanish). **Rail B: pass.**
- [ ] **Step 1 (RED):** a `delivery_state=rejected` + `failing_edge` relay → **bucket D author-return** (to the authoring seat), NOT the operator queue — UNLESS the failed edge is itself an A-gate (egress/D precedence, §2). Expected: FAIL.
- [ ] **Step 2:** Implement bucket D + the egress/D-vs-A precedence keyed on `failing_edge`+stage (s10 already routes A; this completes D). **Step 3 (GREEN):** author-facing; precedence holds. **Step 4:** Commit.

### Task 4: Complete the 7-state FSM surface (`bounced_repair` + `egress_blocked`)
**Interfaces:** Consumes: the s10 5-state A spine. Produces: the full 7-state machine. **Rail A: closed.** **Rail B: pass.**
- [ ] **Step 1 (RED):** `bounced_repair` — an acceptance-stage veto (`delivery_state=rejected`) drives the FSM label → bucket D author-return; `egress_blocked` — the state + its **local-park-resummon** semantics (park locally + resummon, **never auto-redact/send**) are defined and fixtured, the live away-mirror trigger fixture-scoped (away-bridge dormant, step-(d)). Expected: FAIL.
- [ ] **Step 2:** Add both states + transitions to the s10 FSM; `egress_blocked`'s away-mirror send stays UNBUILT (step-(d)) — only the state + local-park-resummon logic land. **Step 3 (GREEN):** 7-state surface complete; `bounced_repair` live, `egress_blocked` fixture-scoped. **Step 4:** Commit.

### Task 5: The elaborate-more fork (m-6 plumbing) — GATE-PENDING on g2
**Interfaces:** Consumes: the s10 ODB/park; the g2 ceiling. Produces: the read-only context-fork. **LOCKS only on the g2 OQ-2 ceiling completing to master. Rail A: closed** (an auto-resolving fork would change the meaning of "the operator resolves"). **Rail B: pass.**
- [ ] **Step 1 (RED):** on the operator choosing **elaborate-more** instead of a verdict button, the conductor forks the parked lane's context under the **g2 ceiling** — `write:read_only` (**no store append, no gate resolution, parked lane byte-unmodified**), `tool:read` (reads context + the ODB `id_ref` slots), `dispatch:none`, posture `interactive`; the fork emits **advisory only, to the operator's decision surface, never gate-bearing**; it **dies on the operator's real verdict**. Negatives: a fork store-mutation → rejected; a fork advisory treated as a verdict → rejected (only the operator verdict resolves); the fork surviving the verdict → rejected. Expected: FAIL.
- [ ] **Step 2:** Build the fork spawn/feed/retire (the m-6 mechanism half; m-5 owns the ceiling). Wire the **annotation-C ruling**: a design/grill gate lands as an **A-bucket gate eligible for the elaborate-more fork**, never silently force-compressed to a bare ODB — the §5(b) no-compress guarantee holds via the fork (the auto-routed meeting surface stays Step-3/4). The verdict rides the same bounded-choice → operator-FROM path; **m-3 re-observe on wake** before the woken lane proceeds.
- [ ] **Step 3 (GREEN):** fork is read-only, advisory-never-gate-bearing, dies-on-verdict; parked lane byte-unmodified; design/grill gate fork-eligible not compressed. **Step 4:** Commit. *(Do not merge T5 until g2 is confirmed landed.)*

### Task 6: 8a hardening (freeze-at-park + migrate-then-validate + re-issue) — GATE-PENDING on g1
**Interfaces:** Consumes: the §9 read-time migrator; the s10 resummon dedupe keys. Produces: the 8a parked-across-bump contract. **LOCKS only on the co-signed 8a contract returning (g1). Rail A: closed** (a silently-dropped parked gate changes acceptance meaning). **Rail B: pass.**
- [ ] **Step 1 (RED):** the co-signed FLOOR — a parked gate across a schema bump: **migrate-then-validate** at wake (via the §9 read-time migrator, then current-form validate); **un-migratable → `held` + escalated + `stale_schema`, never silently dropped, never auto-resolved**. Then the three g1 members: (m2) the bounded-choice set **frozen at park** (the migrator may not silently rewrite a frozen `agent_enum_pick` set); (m3) **choice-set-changed ⇒ bounce-as-stale + RE-ISSUE** — the re-issue is exactly-once, its resummon dedupe keys relate to the original per the g1 ruling (new decision identity vs same), never-dropped-under-crash. Expected: FAIL.
- [ ] **Step 2:** Build the floor + the three members **exactly as the g1 co-signed contract rules them** (do not pre-decide m1/m2/m3 ahead of g1). `stale_schema` (m1) enters the reason vocab as m-2 confirms it. **Step 3 (GREEN):** all floor + member fixtures pass. **Step 4:** Commit. *(Do not lock/merge T6 until g1's co-signed contract lands.)*

### Task 7: The bucket/fixture matrix + the ③ known-A NF fixture
**Interfaces:** Consumes: T1–T4 + the §J2 A-set (my s5-escalations ③ signal-set). Produces: the full matrix + the ③ NF. **Rail A: closed.** **Rail B: pass.**
- [ ] **Step 1 (RED):** the A/B/C/D × terminal-token × `failing_edge` matrix, one RED-first negative per cell; **the ③ known-A NF fixture** — a B-pick / B-absorb over a **known-A** (`gate_category ∈ §J2 A-set`) record ⇒ **raised to A + `gate_category` recorded**, never silently orchestrator-absorbed (the owed m-7-side ③ fixture, now registered). Expected: FAIL.
- [ ] **Step 2:** Land the matrix + the ③ NF (m-3 fidelity on the ③ edge). **Step 3 (GREEN):** every matrix cell + the ③ NF pass. **Step 4:** Commit.

### Task 8: The s10 9-item cleanup refactor (behavior-preserving)
**Interfaces:** Consumes: the s10 codebase. Produces: the refactored surface. **Rail A: n/a (refactor).** **Rail B: pass.** **Invariant: the existing battery stays green + the two mechanical tables reconcile.**
- [ ] **Step 1:** Refactor the 9 items — generic prompter · shared soft-expiry arbiter · one ODB builder · tables-snapshot prompter lookups · drop per-emit `tables.Build` · shared system→operator envelope builder (×5 call sites) · ContentHash/GateID prefix decoupling · `finalizeRun` preserve-flag ownership · genesis reverse-ladder growth. Each item: refactor + run the full battery green before the next.
- [ ] **Step 2:** Confirm the two mechanical tables (consumption→supply · diff→license) reconcile after each item; behavior byte-identical. **Step 3:** Commit per item (frequent commits).

### Task 9: The G4 resummon-cadence config surface
**Interfaces:** Consumes: the s10 resummon timers; the §J config surface. Produces: the operator-config cadence. **Rail A: closed** (a cadence value that meant auto-approve would change acceptance meaning). **Rail B: pass.**
- [ ] **Step 1 (RED):** the two G4 timers (no-response vs answered-but-stalled) read their cadence from **operator-config**, not MF-6's hardcoded 1h/class interim default; **no config value ever means auto-approve** — only channel escalation or the optional per-gate conservative **block**-ceiling (J1: escalate the channel, never the verdict; no hard deadline). Expected: FAIL.
- [ ] **Step 2:** Move MF-6's 1h/class interim defaults onto the §J operator-config surface. **Step 3 (GREEN):** cadence is operator-config; no auto-approve path exists. **Step 4:** Commit.

### Task 10: Re-prompt semantics + the claimless authority-floor `held` edge — GATE-PENDING on dc
**Interfaces:** Consumes: the s10 verdict path + the s9 evidence/authority-floor. Produces: the re-prompt + claimless-held edges. **LOCKS only on the m-3+m-6 design-cell ritual returning. Rail A: closed.** **Rail B: pass.**
- [ ] **Step 1 (RED):** a **non-`accepted` resolution / blocked prompter** re-prompts per the dc-ruled semantics (J1: refresh-before-resummon, never-auto-approve); a **claimless resolution hitting the authority-floor → `held`** (the J1-adjacent fault edge), operator-visible, non-suppressible, never auto-resolved. Expected: FAIL.
- [ ] **Step 2:** Build both edges **exactly as the m-3+m-6 design cell rules them** (design ritual first; do not pre-decide). **Step 3 (GREEN):** re-prompt + claimless-held fixtures pass. **Step 4:** Commit. *(Do not lock T10 until the dc ritual returns.)*

### Task 11: The s11 exit package + hardened battery (feeds the Step-2 step-exit)
**Interfaces:** Consumes: T1–T10. Produces: the s11 exit package.
- [ ] **Step 1:** the full surface demonstrated live on the dogfood store — B/C/D projections, the 7-state FSM, the elaborate-more fork (if g2 landed), 8a hardening (if g1 landed), the ③ NF, the G4 config surface.
- [ ] **Step 2 (institutionalized, day-one):** both mechanical tables green; the label→mechanism sweep + verify-the-summary-line in review; RED-first confirmed per task; named seams recorded for same-file multi-task edits; the catch ledger extends the s8+s9+s10 series.
- [ ] **Step 3 (I-PH sweep):** no store/config/socket/outbox path seat-visible on any new surface (buckets, FSM, fork, ODB).
- [ ] **Step 4:** the merge-decision relay TO the operator (merge operator-only; `HUMAN_MERGE_AUTHORIZATION` at grant). The exit package feeds the master-owed **Step-2 step-exit** (all three legs live + INV-CATALOG red-battery demo + uncached green battery).

---

## Boundary contract
- **Consumes (locked, not reopened):** m-6 c3 + folds (§2 buckets, §3 ODB/fork, §4 FSM/G4, §5(b) no-compress, §7 bind matrix); the s10 spine + resummon keys; the s9 evidence layer + §9 migrator; the co-signed 8a floor; the g1/g2/dc gate returns; the §J2 A-set (③); the m-5 OQ-2 ceiling.
- **Produces:** the B/C/D projections; the full 7-state FSM; the elaborate-more fork; 8a hardening; the fixture matrix + ③ NF; the 9-item refactor; the G4 config surface; the re-prompt/claimless-held edges; the s11 exit package.

## Out of scope (step-(d) / Step-3+ / later — do NOT build)
- the away-bridge / Seam-C token / decision-④ rotate+re-observe (step-(d)) — incl. the live `egress_blocked` away-mirror trigger (state + local semantics only here) · live-lane interjection/steer (Step-3, annotation B) · the auto-routed meeting surface + attach (Step-3/4) · E3/E4 · egress live-chokepoint activation · the TUI/email-client UX (Step-4) · side-effecting execution + the OS sandbox (shelved).

## Acceptance criteria
1. B/C/D projections live: B non-interrupting + promotable; C no-obligation; D author-facing with egress/D-vs-A precedence (T1–T3).
2. The full 7-state FSM: `bounced_repair` live, `egress_blocked` state+local-semantics fixture-scoped (T4).
3. The elaborate-more fork under the g2 ceiling: read-only, advisory-never-gate-bearing, dies-on-verdict, parked lane byte-unmodified; design/grill gate fork-eligible not compressed — **locked on g2** (T5).
4. 8a: migrate-then-validate floor + un-migratable→`held`/escalated/`stale_schema`/never-dropped + frozen-choice + re-issue — **locked on g1** (T6).
5. The bucket/fixture matrix + the ③ known-A NF fixture (T7).
6. The 9-item refactor behavior-preserving; battery stays green; mechanical tables reconcile (T8).
7. G4 cadence on operator-config; no auto-approve path (T9).
8. Re-prompt + claimless-`held` edges — **locked on dc** (T10).
9. Every task carries Rail-A/Rail-B; I-PH clean; both mechanical tables green (Global/T11).
10. The s11 exit package feeds the Step-2 step-exit (T11); byte-exact `{accepted, rejected, held}` + R2 preserved throughout.

## Self-review (against the s11 dispatch + kickoff r4 s11 bullet + g1/g2/dc)
- **Spec coverage:** kickoff s11 bullet → T1–T7; 9-item card + G4 → T8/T9; design cell → T10; exit → T11; g1 gates T6, g2 gates T5, dc gates T10 (decompose-now/lock-on-return); annotation-C ruling → T5 Step 2. No uncovered requirement.
- **Placeholder scan:** literal Go/test bodies delegated to `s11.implementer` (off-frank guide grain); T5/T6/T10 intentionally do NOT pre-decide their gated members (g2/g1/dc rule them) — that is a lock-discipline, not a placeholder.
- **Consistency:** FSM/bucket/terminal-token names match c3 §2/§4; the fork ceiling matches the g2 vector; the 8a floor matches the g1 co-sign; the ③ NF matches the s5-escalations §J2-A-set signal set.
