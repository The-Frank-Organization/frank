# master — `frank` team dashboard

Living status board for the standing `master` governing team **building `frank`** (the conductor — our thin trusted courier). The stable constitution is the charter
(`CLAUDE.md` / `AGENTS.md`); the sequencing plan is `ROADMAP.md`; the integrated design-of-record is
`ARCHITECTURE.md`; **how the team runs a cycle** is `CYCLE-PLAYBOOK.md` (process record + reusable template);
this file tracks live status.

## Current phase
**▶ STEP-1 BUILD OPENED (2026-07-03) — the product is named `frank`; the charter transitioned to the build phase.** Step-0 design is COMPLETE + clean: re-baseline + c6 / c6.1 / c6.1a + the external-review `step1-prep` §C4.3/I-PH amendment — all VP-co-signed; a final 5-lane differential returned the **4 seam lanes CLEAN**. **Slice-1 is CLOSED — COMPLETE at E2 (2026-07-03).** The first `frank` code exists: the thin end-to-end conductor relay, built + closed in `frank/` (tag `s1-close`→`main@f0dcb85`, ledger `674c844`; `go test ./...` 15 packages green + `go vet` clean — CTO re-ran uncached; **operator-ratified + VP confirmatory pass**; full sign-off chain in `master/RECONCILE.md` § S1). The path: decomposed into **4 sequential vertical slices** (team-per-slice, m-x-guided) → VP-gated (`step1-plan`) → operator-ratified → `s1` dispatched (`s1-dispatch`) → guide(m-7)+VP plan-gate → built by the `s1-core` pair through the m-1/m-2 F3 fidelity edge → S1-scoped hardened exit gate → guide deviation ruling (C7 fixture caught a real bug **red-first**) → operator merge gate. **The design→build method is validated on its first slice.** **▶ S2 DISPATCHED (2026-07-03, `s2-dispatch`)** — thicken store/engine (full recovery phases 0–4, durable FIFO, GC/genesis) + the **owed-item-as-typed-record projection** (`OI-S1-F11-SWEEP` = its first customer); guide **m-7** (refines the decomposition's m-1; VP-concurred); F2 conditioned delegation; **a new slice-team** (new sprint = new team — operator; it onboards to the S1 code, m-7 the continuity). Dispatch `…-223913` (r2, supersedes `…-223146`; dispatch-root lint fix per VP F-S2-1). **Deferred (operator — no testbed yet):** the MCP live-adapter + fuller-FieldSpec "wire-it-up" slice — S1 is *built-but-unwired* (real socket daemon, E2-proven, but no live agent session has connected; speaks frank's own socket + a minimal FieldSpec dialect); it returns when there's a live testbed. The design-of-record is the spec; `frank/` is the build target; slice-team relays live in `frank/`.
The post-re-baseline design-of-record was re-reviewed (10 lanes; **90 findings, 0 FATAL, CONDITIONAL-GO** — the 2026-07-01 NO-GO
**structurally discharged**) and cleaned in **c6** (90 = 44 CTO + 45 pair [7/7 pair-approved] + 1 subsumed; all seams converged;
`c6-apply.diff` 18 files/+433/−177), then **seam-hardened in c6.1** — a scoped differential caught + corrected 5 cross-pair seam divergences + 3 dropped flags, all **owner-confirmed** (`c61-fix.diff` +35/−17). The design-of-record is now **verified-at-the-seams, not asserted**, and the confusion-resistance claim was sharpened to *tool-mediated* + a path-hygiene invariant (`I-PH`, external review). **Step-1 PLAN GATED (revise→approve `step1-plan`); Slice-1 (`s1`) DISPATCHED (`s1-dispatch`);** the owed carries (away-bridge, R2/altitude-B fixtures, ③/⑤/GRILL_REQUIRED/routing_escalation, I-PH, m-5-F2) all land at
their own **build step**, not at PLAN. *History / why:* the 2026-07-01 adversarial
pre-build review (`DESIGN-REVIEW-2026-07-01.md`; 16-lens × 3-verifier fleet, ~48 verified findings / ~12 FATAL)
**retracted the "MUST-gate satisfied" certification as unsound** and returned NO-GO. Root causes (~8): the design assumes a **WRAP**
deployment but the operator locked **ATTACH** (①) — inverting every "structural / by-construction / sole-writer /
unbypassable" claim; **conductor-core** (the running program: serialization / crash-atomicity / config-integrity /
recovery) is **nobody's domain and has no design doc** yet Step-1 IS "build the conductor core"; policy config is
**lane-writable under attach** (one Bash line = total governance bypass); altitude-B grain breaks the routing veto;
the Step-1/Step-2 split dead-ends the write path; R2 under-specified/untestable at schema grain (VP-corrected — not a confirmed `chosen_model` leak); pure-judgment
A-categories silently absorb; and the **five operator decisions were never folded into the docs** (② fail-closed vs
still-locked fail-OPEN m-3 text). **The bones are sound — it's a bounded re-baseline (weeks), not a rewrite.**
**Design-review dispatch CLOSED (2026-07-01, VP-approved `design-review/…-150235`)** — NO-GO concurred; deployment fork
**decided**; review-of-record self-consistent (corrected §2A/§2B/§2C bucketing). Re-baseline is the active phase:
- **(a) ✅ DONE — deployment fork DECIDED** (`GRILL-LOCK-deployment-fork-2026-07-01.md`): **attach + interface-level
  guardrail; adversarial containment / wrap / "by-construction" shelved.** Standing claim = *confusion-resistant;
  malicious code-executing agent explicitly out of scope.*
- **(b) ✅ COMPLETE — m-7 Conductor-Core substrate DESIGNED + LOCKED** (`c4` CLOSED 2026-07-02; VP design-lock co-sign
  `c4-design-m-7/…-040327`; `DESIGN_LOCK_ID c4-design-m-7-lock`; close declaration `…-040839`). The **engine** — one-writer
  durable-FIFO serialized commit loop + atomic clear-on-pop, rename-pivot crash-atomicity, phase-0–4 recovery, HELD
  internal-fault disposition, byte-exact `{accepted,rejected,held}`, MCP `{submit,project,read}` interface guardrail,
  per-domain-section single-digest trusted config, local-outbox-only send, seat-binding base — folded into
  `ARCHITECTURE.md` **§C4** + `RECONCILE.md` **c4**. **8 CQs closed full-pair** (6 pairs re-engaged; the gate/config
  cluster took an r1→r2 revise cycle; closure gate held throughout); grilled (`c4-grill-m-7`); **claim boundary held —
  confusion-resistant; m-7's headline by-construction claim = the §2.4 serialized-loop double-accept kill (one of the four licensed classes — §C4.3).** *The NO-GO's
  core defect — the running substrate was nobody's domain — is corrected.*
- **(c) ✅ COMPLETE — global claim-sweep + decision-folds** (`c5` CLOSED 2026-07-02; VP close-approve `c5-decomp/…-143205`).
  All 7 design-of-record docs (m-1..m-6 + `ARCHITECTURE`) swept to the honest line — adversarial-strength claims
  ("by-construction / sole-writer / forgery-robust / tamper-resistant") → **confusion-resistant + D5 residual**; the genuine
  trusted-engine invariants (R2 gate-grammar, observer-selected controls, authority-ceilings) **kept**; byte-consistent
  (`{accepted,rejected,held}`; `bounced` retired). Decisions **③** raise-only (+ §J), **⑤** ODB carve-out, **④** away-token
  **recorded as §2C build-carry** folded; 6 lanes owner+implementer-approved (3 must-revise→approve cycles). Folded into
  `RECONCILE.md` **c5**. **The docs now tell the honest story end-to-end.**
- **(d)** §2C at its build step (the `re-mint-supersedes` + decision-④ away-token carries + the routing-lane **R2 `gate_referenceable`-per-column** & **altitude-B per-row** carries land there);
  **(e)** THEN Step-1 PLAN.

**Step-1 PLAN remains the operator-opened gate (re-baseline step e). (a)/(b)/(c) done; (d)'s items are deferred
build-step carries — no PLAN / code / spike yet.**

### Superseded (history — the retracted certification)
**~~Step-0 COMPLETE + build-readiness pass COMPLETE — the full MUST-before-Step-1 gate is SATISFIED (2026-07-01)~~.**
[RETRACTED by the adversarial review above; preserved for the audit trail.]
After c3 locked all six domains, the team ran the **measure-twice pass**:
- **Build-readiness review** — 7 red-team lenses → `READINESS-REGISTER.md` (GO-WITH-FIXES; 8 clusters).
- **Runtime + model research** — how the conductor rides existing runtimes, **provider-agnostically**, with **persistent
  seats** → `RUNTIME-RESEARCH.md`; cross-checked vs an external brief, **primary-source-verified** (§14), **VP-reviewed**.
- **Five operator decisions ①–⑤ RECORDED** (`READINESS-REGISTER.md` → *Operator decisions*). **① = attach-first +
  "confusion-resistant"** (conductor-owned identity; "sole external sender **by construction**" is a **6-property
  spike-gated** milestone, never shipped on assumption). ② fail-closed authority-class · ③ raise-only A/B · ④ away-token
  rotate+re-observe · ⑤ narrow ODB model-name exemption.
- **The two MUST-reconcile fixes CLOSED (VP co-signed):** **Cluster 1** (m-1/m-2 `submit()` write-path seam — pre-append
  form+lineage → one atomic `accepted`; **Step-1 = store+form+lineage**, observe reserved Step-2; `readiness-fix-c1`
  VP co-sign `…-230335`) · **Cluster 4** (m-2 routing FieldSpec ← m-4 `:200-210`; R2 grammar-enforced; computed-field
  homes; posture rides `seat_archetype`; `readiness-fix-c4` VP co-sign `…-013613`).

**Tracked carries — route before their relevant PLAN, NOT gate blockers:** row-parity §15 Q-F (pre-Step-1-PLAN) ·
runtime-`away` posture-vocab split (pre-Step-2 away-bridge) · the ①-unblocked m-1 claim/identity fixes + the
decision-implied owner-fixes (5a/6a/6b/6c) · the m-6 `delivery_state` `bounced→rejected` ripple · the **Cluster-3
security spike** (the "by-construction" milestone — a separate operator-opened gate).
**Readiness relay trail:** `readiness-decomp` · `readiness-reconcile` · `runtime-research` · `readiness-fix-dispatch` ·
`readiness-fix-c1` · `readiness-fix-c4`.

### c3 lock (history — all six domains locked, `ARCHITECTURE.md` §1–§C3)
**Cycle c3 — CLOSED / LOCKED** (`c3-lock` 20260630-191315/191525; C3.6 integration capstone; m-5/m-6 stood down).

What c3 locked (detail: `ARCHITECTURE.md` §C3; the two design docs):
- **m-5** — an **archetype = one governed expansion-slot** binding {topology + gate-set + ceiling-at-spawn +
  observe-invariants + routing-prior}; `lower_snake_case` tag-space; observer-selected control-property invariants (F1);
  3-axis open-map ceiling vector + modular `external_send`; T1/T2/T3 + sensor-full; `actuator` = derived class.
- **m-6** — a promote-and-bind human surface: A/B/C/D buckets; ODB render+capture + the read-only elaborate-more fork;
  the 7-state park/wake machine; the opt-in egress-gated away-bridge; meeting-lane routing; the interjection host.
- **The seam-of-record** — human-mode = posture × `surface_intent` (four-class, conductor-derived, non-gate;
  gate-bearing binds off locked fields → **no m-2 micro-fold**).
- **Seam C = A** — m-1 owns the away-mode inbound-token mint/verify via its reserved `certification` seam (**additive**,
  no c1 reopen); m-6 owns the bridge. The first conditional-upstream-contract-check.
- **C3.6 capstone** — the six domains compose (writer-backed + acyclic; three seams closed; locked invariants intact;
  deferrals recorded as build-carries, not gaps).

**Cycle c2 — CLOSED / LOCKED.** m-3 Observation & Evidence + m-4 Routing & Policy locked as the Step-1
runtime-intelligence layer (`c2-lock` 20260630-043859 / -044308). Detail below + `ARCHITECTURE.md` §C2.

What's locked (detail: `ARCHITECTURE.md` §C2; the rev2 domain docs):
- **m-3** observe-as-send gate (conductor observes done from outside the lane, observer-only against R3; evidence
  ladder; fail-closed egress chokepoint).
- **m-4** routing governance-record (routing = a first-class recorded justifiable decision; two-layer bucket prior;
  fail-closed `route_dispatch()`; **model = payload, never a gate input — R2**); GL-4 routing-templates (record
  mechanism; pane-spawn rides existing multiplexer infra).
- **The R2-preserving seam** — silent-deviation block via m-3's generic integrity-veto; no model-derived predicate
  in any gate; bucket-vs-bucket; snapshot-provenance (holds for opaque lanes).
- **Two opaque archetype atoms** — `slot_in` work-archetype (conductor-classified at acceptance) + `seat_archetype`
  per-seat (per-assignment record home); concrete tag-space/semantics **reserved to m-5 (c3)**.
- **M4-1** — routing B→A escalation rides the c1 monotonic HUMAN_GATE routing-raise (no new gate class, R2-safe).

**Cycle c1 — CLOSED / LOCKED.** Step-1 foundations jointly locked: **m-1 Trust & Identity + m-2 Forms &
Determinism**. VP close-confirm `c1-joint-lock` 20260629-180934; operator §J ratified; close 181713.

**Cycle c1 — CLOSED / LOCKED.** Both Step-1 foundations jointly locked as the frank design-of-record:
**m-1 Trust & Identity + m-2 Forms & Determinism**. VP `approve` close-confirm (`c1-joint-lock`,
20260629-180934); operator §J ratified; close declaration `c1-joint-lock` 20260629-181713; pairs stood down.

## Org status
- `master.orchestrator-planner` (CTO) — Claude Opus 4.8; online.
- `master.orchestrator-reviewer` (VP) — GPT-5.5 (high); idle — **co-signed the c4 conductor-core DESIGN-LOCK** (`VP_DESIGN_LOCK_CO_SIGN: approve`, `c4-design-m-7/…-040327`; 5 checks passed, no blocking finding). Caught 4 routing/certification errors across the CQ cycle before they propagated.
- **m-1 Trust & Identity** — c1 **LOCKED**; stood down (+ the additive Seam-C inbound-token mint/verify carry, C3.7 — does not reopen c1).
- **m-2 Forms & Determinism** — c1 **LOCKED**; stood down.
- **m-3 Observation & Evidence** — **c2 LOCKED**; stood down.
- **m-4 Routing & Policy** — **c2 LOCKED**; stood down.
- **m-5 Workflows & Archetypes** — **c3 LOCKED**; stood down. The archetype system (the governed expansion-slot: tag-space,
  ceiling vector, T1/T2/T3, sensor + derived-actuator, the human-mode vocabulary).
- **m-6 Human Surface & Scheduler** — **c3 LOCKED**; stood down. The promote-and-bind human surface (buckets / ODB /
  park-wake / away-bridge [Seam C = A] / meeting / interjection host).
- **m-7 Conductor-Core** — **c4 LOCKED; stood down** (design-of-record locked 2026-07-02, VP co-sign `c4-design-m-7/…-040327`;
  re-baseline step (b) COMPLETE). The runtime substrate the six policy domains ride on — owns the
  ENGINE (serialized commit loop / crash-atomicity + recovery / internal-fault disposition / trusted config load +
  integrity / attach + interface-guardrail enforcement / local-outbox-only send / store genesis+GC), hosts+executes the
  six CONTRACTS, does not re-own policy. Owns `DESIGN-REVIEW-2026-07-01.md` §2A. Audit reconciled (still-open; VP co-sign
  `…-162319`; ledger `RECONCILE.md` c4). Design-of-record `design/2026-07-01-conductor-core-design.md` — grill-locked
  `c4-grill-m-7`, seam matrix biting, F8 semantic claim-sweep clean. **CQ gate SATISFIED** — all 8 CQs
  (CQ-1/2/3/4/4b/5/6/8) closed with approved full-pair triads via 3 COORD clusters (`c4-cq-gateconfig` / `c4-cq-m1` /
  `c4-cq-slotin`, 6 pairs re-engaged; gate/config took an r1 must-revise→fold→r2 approve cycle). CTO fold
  `…-031533` + CQ-6 correction `…-032227`, **certification VP-approved `c4-cq-coord/…-032843`** (CQ-6 base-closed;
  **re-mint-supersedes carried as a §2C away-bridge build-carry**, not lock-bearing; + 3 lock-assembly items: m-4
  per-section stamp · byte-exact `{accepted,rejected,held}` · exactly-one-outcome check). **DESIGN_LOCK `c4-design-m-7-lock`
  EFFECTIVE** (design r3 + lock r5 pair-approved; CTO-certified `…-040011`; **VP co-sign `…-040327`**; close declaration
  `…-040839`). Folded into `ARCHITECTURE.md` §C4 + `RECONCILE.md` c4. **c4 CLOSED; re-baseline step (b) COMPLETE** → (c)
  global claim-sweep + (d) §2C-at-build-step remain → **(e) Step-1 PLAN** operator-openable.
- **s1 Slice-1 team** — **S1 CLOSED / COMPLETE at E2** (tag `s1-close`→`main@f0dcb85`; operator-ratified 2026-07-03 + VP confirmatory pass). The thin end-to-end conductor relay is built in `frank/` (15 pkgs green, `go vet` clean — CTO re-verified uncached). Built by the `s1-core` pair, m-7-guided, through the m-1/m-2 F3 fidelity edge (F-M1-1 block→fold→approve). Slice-1 CLOSED; the team continues into S2.
- **s2 Slice-2 team** — **DISPATCHED 2026-07-03** (`s2-dispatch` r2 `…-223913`); a **new** slice-team (new sprint = new team; onboards to the S1 code, m-7 the continuity), guided by **m-7** (m-1 keeps authority over the owed-item `record_kind`/store-layout/API + fidelity — VP watchpoint). **▸ S2 CLOSED / COMPLETE at E2 (2026-07-04)** — merged `main@b322b6d`, tag **`s2-close`** (operator-authorized, token-lineage clean; battery 18 pkgs green at main — CTO re-ran uncached; `master/RECONCILE.md` § S2). Engine thickened (recovery 0–4 · durable FIFO · GC/genesis · **the owed-item projection**); **`OI-S1-F11-SWEEP` closed THROUGH frank** (operator-stamped owed record → executed sweep → disposition → open set empty on a real store — **frank's first real governance transaction**); 2 latent S1 races found+fixed by the fresh team; **zero owed items ride out**. s2 team stands down. **▸ S3 CLOSED / COMPLETE at E2 (2026-07-04)** — `main@b5a2c95` (`--no-ff`), tag **`s3-close`**; VP pre-integration confirm + operator's four decisions on record; **battery 20 green at main — CTO re-ran uncached** (`master/RECONCILE.md` § S3). **frank speaks the real protocol on fresh stores** (MVP dialect deleted; the 62-check dissolution EXECUTED — 115 rows, zero uncovered, [VP-W] adjudication checked row-by-row); five independent verification chains; mid-slice s3-scope-q1 ruled DEFER cleanly. **Riding out: `OI-S3-CONFIG-CHANGE` alone** (→ the wire-up slice). s3 team stands down. **▶ FORK ELECTED (operator, 2026-07-05): s4 = THE WIRE-UP** (the per-seat MCP shim — live Claude Code/Codex sessions on frank; ends the operator-as-transport; discharges `OI-S3-CONFIG-CHANGE` via the §7 config-change record; **first live E3**, scoped to transport). **Renumber of record: old Section-4 (consumer schemas) → s5**, built *over* the wired conductor (its rows = the first real §7 changes; its relays = the first usage data). **s4 BUILD DONE — E2 floor green at 3 independent stations (`s4-wire-impl@7dc5f92`, 14 commits; 21 test-bearing pkgs, race clean, zero regression; S2 store untouched). ▶ LIVE GATE-DAY PENDS THE OPERATOR** (`s4-exit-gate …-145109`, TO operator): designate seats A/B (procedure assumes Claude Code ↔ Codex) → run `results/e3-gate-procedure.md` (init THE persistent team store → mint → wire → **the first no-hand-relay relay** → live adversarial/crash legs → **the §7 round-trip authored by the operator** → OI-S3-CONFIG-CHANGE disposition) → evidence returns via the pair → s4-orchestrator verifies → master SITREP → merge/`s4-close`. [VP-W1..W4] all folded pre-handoff. Step-1 remaining: s4 → s5 → the step-exit test. — forms/lineage: full FieldSpec registry (frank speaks the REAL protocol envelope) + the 62-check linter dissolution with the **FULL replay** as the gate centerpiece + `schema_version`/migrators + the R2-negatives/GRILL_REQUIRED owed carries. Guide **m-2** (m-7 consulted on config seams; m-1 fidelity on store touches); NEW slice-team; F2 conditions. **The wire-up slice is explicitly next-after-S3** (S3 is its prerequisite — real relays can't validate through frank until the registry lands). Scope: full recovery (phases 0–4) + durable FIFO + GC/genesis + the **owed-item projection** (`OI-S1-F11-SWEEP` = first customer). **F2** conditioned delegation (no bootstrap guide+VP plan-gate; escalate to master only on a trigger). Exit gate re-runs F9/F11 under the new recovery machinery + discharges the `OI-S1-F11-SWEEP` sweep. **Deferred:** the MCP live-adapter/wire-it-up slice (no testbed yet).
- **▸ s4 WIRE-UP — CLOSED / MERGED (2026-07-05): the operator-as-transport is OVER.** Merged `main@fb61fda` (`--no-ff`, parents `a47381a`+`6a23cf0`), tag **`s4-close`**; operator-directed grant (`s4-merge-gate/MERGE-GATE …-233708`) → `s4-wire.implementer` executed → **battery 21 green at the merge commit, CTO re-ran uncached**; VP-confirmed; `master/RECONCILE.md` § s4. **A governed relay flew Claude Code → Codex, conductor-stamped, no human transport** (`relay-4a33925b…`). §7 config-change + owed cycle exercised live; **`OI-S3-CONFIG-CHANGE` discharged → Step-1's mandatory owed set is EMPTY.** 3 gate-day findings dispositioned (F-GATE-1 fixed; **F-GATE-2 owed-fill-time fixed = 5 registry `required_when` rows**; F-GATE-3 ruled out-of-I-PH-scope). Discretionary follow-on `OI-S4-TOKEN-SCOPE` (incl. `genesis`, VP+m-7). **▶ Step-1 remaining: s5 (consumer schemas, over the wired conductor) → step-exit test → Step-1 closes.**
- **▸ s5 DOGFOOD (the first team ON frank) — STOOD DOWN AT CHECKPOINT (2026-07-06, operator-directed stop-the-line).** s5 ran its live governance through frank (dispatch `relay-45e263fa` · boot-ack + full onboarding via `submit`/`project`/`read` · 5 sub-seats + an operator seat minted · hub-and-spoke m-x routing worked · operator-authored waiver record live). The run delivered its real objective early: **`master/TRANSPORT-FINDINGS-2026-07-06.md` (F1–F17)** — headline **F11: lineage livelock under concurrent traffic** (a seat can construct NO acceptable parent; even report-only relays blocked), plus the render↔validate divergence family (F2/F6/F7/F13), the F9 `intake_id`-reuse audit break, and the F17 irrevocable-waiver gap. What held: crash-atomic store (3 restarts, zero corruption), channel-stamped FROM, I-PH clean. **Conductor decommissioned; store archived** (`~/frank-archives/frank-team-store-s5-dogfood-20260706`). **s5 RESUMES as a standard file-relay team to its checkpoint** (`s5-resume/PLAN …-034602`; [VP-W7] §7-live-store leg adapted → rides the transport-fix relaunch; dogfood-record exit item closed early by the ledger). **m-x fidelity packet Q1–Q11 fully reconciled** (`s5-fidelity/RECONCILE …-034602` — all six owners answered, zero rejections): owed-#3 **EMPTY** · **`routing_escalation` §J2 LANDED + the [VP-W5] hold LIFTED** (m-2's exact delta → s5-a's registry pass) · no Step-1 `slot_in` writer · `scope_paths` struck · Q4 artifact set final (no performative migrator) · ⑤ shape ruled (real code, real chokepoint, dormant, fixture-driven) · no live `record_kind` widening · `attestation_source` folded into s5-a (O-2). **▶ Next: s5 checkpoint (file-relay) → the TRANSPORT-FIX cycle (m-7/m-1/m-2, grilled — {F11,F4,F5} liveness first) → step-exit test → Step-1 closes.**
- **▸ s5 CONSUMER SCHEMAS — CLOSED (2026-07-06): STEP-1'S BUILD QUEUE IS EMPTY.** Close tip `main@f31d43a`, close record `7e5c527`, tag **`s5-close`**; three integrations under the WRITTEN grant chain; operator close-gate exercised in-session; **battery 23 green uncached — CTO re-ran at the close tip** + probes (`master/RECONCILE.md` § s5). Registry **47→83 rows @ `s5-fieldspec-v3`**, consumer fields declared + proven dormant (334-subtest sweep [VP-W3]); ③ live-wired (S1+S2 live, S3 inert per the confirmed claim boundary — in the code, `detector.go:25`) + ⑤ dormant at the real drain + `routing_escalation` + I-PH; zero-loss replay **incl. the real-store leg vs the archived dogfood store (three seats)**; NO envelope migrator (`Current=1`); **`OI-S4-TOKEN-SCOPE` closed by the scope narrowings** (`genesis` in NO scope; owed/`gate_resolution`/`disposition` operator-only). Fresh-team bar paid a **6th** time (DEF-1..5 fixed red-first, incl. DEF-2's lane-suppliable-system-headers reject rule). Q1–Q11 + M-1..M-4 settled by all six owners, three adversarial implementer approves, zero unresolved contradictions. s5 team → standby on the master fold. **s5 close RATIFIED (operator, 2026-07-06).** **▸ OPERATOR RULING (2026-07-06): Step-1 does NOT close until the transport is fixed** — the honest reading of the step's goal ("remove the operator-as-transport"; a transport that livelocks under multi-seat load hasn't). **▶ Step-1 remaining: s6 = THE TRANSPORT FIX (in-step; design-amendment phase first — m-1 lineage/parenting THE fork, m-7 engine liveness, m-2 one-codec; F1–F17 the spec seed; GRILL on the parenting fork) → the step-exit test ON THE FIXED CONDUCTOR (the ROADMAP:83-85 legs + §7-applies-s5's-registry as the first live act + the F11 regression leg = the dogfood traffic pattern replayed without livelock) → Step-1 CLOSES.**
- **▸ s6 DESIGN PHASE — CLOSED / VP CO-SIGNED (2026-07-06, `…-220325`).** Three pair-complete domain halves (m-1: the fork+§B/§C/§D+B-3 · m-7: A-1..A-4+D-1/D-2+B-1 · m-2: the one-codec+B-2) integrated as **`S6-AMENDMENT-SET-2026-07-06.md` (r3)**; the parenting fork **GRILLED with the operator** (`GRILL-LOCK-parenting-fork-2026-07-06`: branch A conductor-computed PARENT · fallback-never-bounce settled by the zero-true-bad-picks audit · m-4-gated Sharpening-D, discharged); the boot stage added on operator request (derived-only activation — no persisted marker). **The gates caught real defects at every altitude:** two pair must-revises (the B-1 pre-active classifier hole; m-2's r1 self-contradictions) + two VP must-revises (the F14 mis-home; the activation-marker cross-domain lock conflict + byte-fidelity) — zero silent folds. §C4 pointer LANDED (the amendment docs are the authoritative deltas pending the post-build fold). **▶ the s6 BUILD slice dispatching** (package held for VP pre-handoff; [VP-W1..W3] of the co-sign folded: hold-on-m-7-revert · FX-B1g explicit · seven registry rows, no marker row).
- **▸ s6 TRANSPORT FIX — CLOSED (2026-07-08) ★ STEP-1 CLOSED ★** Merge `main@6a1198a` (`--no-ff`, parents `73116e9`+`58f2233`), tag **`s6-close`**; operator-granted, VP-confirmed, **battery 24-ok uncached at the merge commit — CTO re-ran** (`master/RECONCILE.md` § s6). The co-signed amendment set implemented whole; **THE STEP-EXIT TEST PASSED LIVE**: the operator §7-apply as the first live act · ROADMAP:83-85 + gate→outbox · **the F11 redrive 14/14 with ZERO parent-class and ZERO same-context re-render bounces** · the live boot walk (`minted→bound→active`, derived-only, no restart). Ten verification chains; the fence held (2 typed OIs ride out). **Step-1's goal — remove the operator-as-transport — is DELIVERED on a transport that survives its own governance load.** **▶ NEXT: Step-2 planning (governance hardening + minimal comms — observe-as-send, evidence ladder, park/wake) with the INV-CATALOG follow-on first in the queue and the dogfood relaunch riding Step-2's first live store.**
- *(superseded)* s4 exit gate master-accepted + VP-confirmed — `s4-wire-impl@6a23cf0` (16 commits; **21 pkgs green, CTO re-ran uncached**; `s4-exit-gate/RECONCILE …-231116`). **THE OPERATOR-AS-TRANSPORT IS OVER, demonstrated live:** a governed relay flew from a live **Claude Code** seat to a live **Codex** seat, conductor-stamped `from: s4-wire.host-a`, no human transport (`relay-4a33925b…`, verified at the store). §7 config-change round-trip on an existing store + the owed cycle exercised live; **`OI-S3-CONFIG-CHANGE` discharged → Step-1's mandatory owed set is EMPTY.** Three gate-day findings all dispositioned: F-GATE-1 (handshake `serverInfo.version`) fixed+class-pinned; **F-GATE-2 (owed fill-time-authority gap — my finding) fixed as 5 registry `required_when` rows, red-first**; F-GATE-3 (shim-stderr) ruled out-of-I-PH-scope with grounds. One discretionary follow-on proposed: `OI-S4-TOKEN-SCOPE` (the first post-close owed item on the live store). **Remaining: (optional VP pass) → operator merges + tags `s4-close` → master folds S4 → s5** (consumer schemas, built OVER the wired conductor — first usage data) **→ Step-1 exit test.**
- Runtime / Product domains (m-8..m-12) — future cycles.

## Decisions (c1 — LOCKED)
- **Org / Roadmap / both c1 reconciles** — LOCKED (VP-approved); own-the-gate-first; `C1_PRODUCT_SCOPE = B`.
- **Transport** — Option A minted per-seat credential = confusion-resistant (GRILL-LOCK D4; forgery-robust-by-construction = the D3-shelved wrap milestone, not the Step-1 claim).
- **Schema carrier** — bespoke FieldSpec registry; strict form-only submit; `X-` overflow; zero-migrator versioning.
- **Shared seam** — R1 operator/special-address; R2 routing = separate seat-stamped relay (model never a gate input);
  R3 observe-integrity (`evidence_integrity`). PARENT (m-1) strengthens lineage (m-2) → more strongly confusion-resistant (D5 residual).
- **identity ≠ authority** — ratified (`ARCHITECTURE.md` §5).
- **§J operator-judgment** — RATIFIED (`ARCHITECTURE.md` §J): J1 `on_timeout = hold_and_resummon` (never auto-approve);
  J2 `gate_category` default set, **operator-configurable** (forward m-6/config); merge split by target branch +
  protected-branch set.

## Forward requirements (recorded; do not reopen the lock)
- Customizable `gate_category` membership / A·B map / protected-branch set — m-6 + config (`ARCHITECTURE.md` §J).
- Opt-in away-mode external-inbox bridge, egress-gated — m-6 scheduler + m-3.
- DI-2/DI-5 independent-isolation realization (the fork-2 infra call) — build cycle.
- m-3 / m-4 / m-6 full domain designs — later cycles.
- Full PLAN carry-forwards in `ARCHITECTURE.md` §6.

## Relay log
See [`master/relays/INDEX.md`](relays/INDEX.md). c1 terminal: close declaration
`c1-joint-lock/RECONCILE-orchestrator-planner-20260629-181713.md`.
