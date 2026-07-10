# DESIGN-REVIEW — adversarial pre-build review of the frank design-of-record

**Status: REVIEW-OF-RECORD. Verdict: NO-GO — re-baseline Step-1 before any build (bounded; not a rewrite).**
This is the "measure a third time" pass the operator commissioned before authorizing the Step-1 cut. It
**overturns** the `FULL MUST-BEFORE-STEP-1 GATE SATISFIED` certification (VP-co-signed `readiness-fix-c4/…-013613`,
2026-07-01) — that certification is **unsound** (see §6). Docs-only, in cwd. No design changed by this doc; it is
the evidence base that drives the re-baseline.

**VP-reviewed (`design-review/RECONCILE-orchestrator-reviewer-20260701-141436.md`): CONCUR with NO-GO + retraction.**
Both grounds independently confirmed against source: decision ② unfolded (m-3:63/109-111/128-130 vs register:340-344)
and conductor-core unowned (ARCHITECTURE:171/378-381). The re-baseline sequence (§5) is concurred. **One bounded
correction folded:** the R2 finding was overclaimed as a confirmed `chosen_model` leak → corrected to
"under-specified / untestable at schema grain" (§1.6, §3.6). This sharpens, does not weaken, the NO-GO.

**UPDATE 2026-07-01 — the deployment fork (re-baseline step a / must-fix #1) is now DECIDED via
`GRILL-LOCK-deployment-fork-2026-07-01.md`.** Operator grill outcome: Step-1 threat model = **confused-not-adversarial**;
Step-1 = **ATTACH** with an **interface-level** guardrail (seats act only through `submit()`; config is not a seat
tool); **adversarial containment / wrap / "by-construction" is SHELVED indefinitely** (research-gated — even our §14
found the sandbox foolable). **This collapses ~9 of the 12 FATALs** (the attach-vs-wrap inversions) to two cheap
actions — a **global claim sweep** ("by-construction/structural/unbypassable" → "confusion-resistant; malicious
code-executing agent explicitly out of scope") + **documented accepted-risks**. **VP correction (`…-144217`): what
collapses is the adversarial *security claim*, NOT the confused-agent *interface mechanism*.** The interface guardrail
(seats act only through `submit()`/`project()`/`read()`; raw store/config paths **excluded from every seat tool
surface**; config conductor-owned + trusted-load; local-outbox-only send) and **fill-time authority / form rendering**
remain **hard conductor-core requirements + fixtures** — because a confused seat with a shell/file tool could follow a
bad instruction *around* the gate. So **§2 must-fix #2 (config-integrity) is REVISED**: drop the adversarial *isolation
redesign*, keep **trusted config-load + not-in-seat-tool-surface** inside conductor-core. The **threat-independent
MUSTs remain**: stand up **conductor-core** (interface guardrail + fill-time-authority mechanism + single-threaded
serialized commit + crash-atomic commit + recovery + internal-fault disposition), fix the **Step-1/Step-2 phase-split**,
and the two **confused-agent guardrails** (pure-judgment A-floor + decision-② fail-closed). **R2, altitude-B, and the
away-token/restart binding remain MUSTs *before their build step*** (R2/altitude-B before the routing schema ships;
away-token before park/wake — stale-approval/sibling-reuse bite *normal* operator flows, not just malicious ones). Full
corrected bucketing: `GRILL-LOCK-deployment-fork-2026-07-01.md` (Design-lock impact) + VP `…-144217`. Re-baseline is
smaller than the raw count implied, but the confused-agent mechanisms are *built*, not waved away.

**Provenance.** A maximally-skeptical multi-agent adversarial review (`design-of-record-adversarial-review`
wf_e4a39ca9-99e + continuation wf_7a4d05c2-767): **16 hostile skeptic lenses** over the whole design-of-record →
**every finding cross-examined by 3 diverse verifiers** (refute / severity / reproduce; a finding survives only on
≥2-of-3 not-REFUTED) → a **gap-hunt critic panel** → **synthesis**. ~293 agent-runs, ~26M subagent tokens. The
first run's tail (2 lenses' verifiers + gap-hunt + synthesis) died on a session limit; a surgical continuation
recovered them without re-running the completed work. Raw data:
`tasks/wljxt3ugu.output` (37 pass-1 survivors), `scratchpad/survivors37.json`, `tasks/wblshda94.output`
(recovered + gap + synthesis).

**The scoreboard: ~48 verified surviving findings, ~12 FATAL.** The FATALs are **not independent bugs — they
collapse into ~8 root causes.** That collapse is the whole diagnosis: the fix is a small number of architectural
decisions + mechanical folds, not a bug-bash — which is why the recommendation is a **bounded re-baseline (weeks),
not a rewrite.**

---

## 0. BLUF

- **NO-GO on starting Step-1 build.** The three-or-four load-bearing claims the product is sold on —
  *forgery-robust-by-construction identity*, a *sole-writer append-only store*, a *structurally-unbypassable egress
  chokepoint*, and *"a decision that must reach the operator reaches the operator"* — are **FALSE on the
  deployment the operator actually locked (attach-first, decision ①)**, and the runtime substrate that would host
  all six domains (**conductor-core**) was never designed or assigned to anyone.
- **RE-BASELINE, not fix-then-build.** The flaw density is the diagnostic signature of a design written for a
  *different deployment* than the one chosen, plus a *missing foundational domain* — you cannot patch that during a
  build.
- **The bones are sound.** Domain decomposition, seam-reconciliation machinery, the SR-26-2 deviation-register
  concept, observe-as-send, owner-typed FieldSpecs, and the honest-fallback *pattern* are all good. Every FATAL
  identity/store/config finding is **"the honest-fallback pattern was not applied here"** — mechanical to fix, and
  the team already knows its shape.

---

## 1. Root causes (the ~8 themes the ~48 findings collapse into)

1. **ATTACH-mode inverts every "structural / by-construction" governance claim — the design was written for WRAP.**
   *~9 findings are one architectural mismatch expressed nine times.* I1 sole-writer, unforgeable operator-FROM,
   sole-external-sender, config-integrity, fill-time authority, and the "three-rung hardening ladder" were all
   proven conditional on the conductor **owning the process tree** (separate uid / 0700 / no lane file-handles).
   Under attach (same uid, lanes hold Bash/curl, native runtime comms un-neutered) they degrade from
   **structural → best-effort/cooperative**, and the docs a *builder* reads still assert the strong version. Must
   be resolved by **decision** (wrap-for-Step-1 vs honest downgrade), not patched.
2. **The runtime substrate is unowned: conductor-core + its serialization/crash-atomic primitive are named, not
   designed.** Step-1 *is* "build the conductor core," but the running program (process/concurrency, serialized
   commit loop, crash-atomic multi-file append, recovery, config load, fault posture, store genesis/GC) is **no
   pair's domain and has no design doc**. This is why ≥5 findings are literally `owner: none`. The C3.6 "build-ready"
   capstone only checked the inter-domain consume-graph; it never checked the substrate exists.
3. **Altitude-B fan-out breaks record-level fields (wrong grain).** The silent-deviation veto and the SR-26-2
   override-register — the staked novelty — are record-level scalars; `any()`-aggregation lets one honest
   deviation mask every silent deviation on sibling seats. Breaking schema-shape change.
4. **Pure-judgment A-categories are silently absorbed — the core promise is false for the judgment half.**
   `gate_category` and `CEREMONY_TIER` are agent-free-picks with no floor; a seat labels an A-worthy
   `authz_security`/`product_semantics` decision as B (or picks `ceremony=low` to shrink its own required-set) and
   it never reaches the operator. RAISE-ONLY (decision ③) only governs records that *already entered* a gate.
5. **The Step-1/Step-2 phase split was done without splitting the required-set — Step-1 as specced cannot run.**
   Step-1's form gate requires observe-owned fields (`ACTIONS_GIT_REF`/`FINAL_GIT_STATUS_SHORT`) whose only writer
   is the m-3 hook deferred to Step-2 → **every action-claim relay bounces uncompletably.** The first shipped
   conductor cannot accept its first IMPL report.
6. **R2 (model ≠ gate input) is under-specified / untestable at schema grain (VP-corrected).** *Not* a confirmed
   live leak — m-2 §17.3:292 labels `chosen_model` non-gate-referenceable in prose, and m-4:96-104 lets the observe
   layer read it as *payload* for bucket-binding while keeping model-derived predicates out of gates. The real
   pre-build defect: the predicate **grammar** examples exclude the **stale ghost name** `selected_model`/`model_*`
   (m-2:84/91/303/345) while the live row field is `chosen_model`, and row-array columns have **no enforceable
   per-column `gate_referenceable:bool` FieldSpec + negative fixtures**. So R2 is asserted in prose but not
   *mechanically* enforced or testable at the schema the tool reads — "structurally impossible to violate" is
   unproven, not disproven.
7. **Recorded operator decisions + reconciliations were never folded into the locked docs — yet the gate was
   certified satisfied.** The five operator decisions (①–⑤) + multiple renames were shipped as post-gate "tracked
   carries." **Decision ② (fail-closed on the authority path) is contradicted by still-universal fail-OPEN text in
   locked m-3** (`self_reported` still delivers) — a faithful builder ships the behavior ② closed. Plus a
   keying-on-a-never-written-value class (`record_kind=routing` vs `routing_decision`; `bounced` vs `rejected`)
   that silently no-ops seams. "Byte-consistency verified" is unsound.
8. **The away-mode bridge — what makes autonomous away-runs safe — is riddled with collisions + a forgery.** It
   fail-closes on its own verdict token, strands the secret-bearing high-stakes A-gates it exists for on an
   invisible local resummon, deadlocks under `unattended` posture, routes design/grill gates into a Step-3/4
   vacuum, and (FATAL) its per-choice token lets a still-valid `approve` overturn a rejected gate.

---

## 2. What must happen before Step-1 build — CORRECTED buckets (VP `…-144217`/`…-145656`)

Load-bearing distinction: **adversarial security CLAIMS collapse; confused-agent interface MECHANISMS become
conductor-core requirements.** This supersedes an earlier flat "10 must-fixes" list; the buckets below are the
authoritative re-baseline scope, mirroring `GRILL-LOCK-deployment-fork-2026-07-01.md`.

### 2A — CONDUCTOR-CORE Step-1 requirements + acceptance fixtures (MUST before Step-1 PLAN)
1. **Deployment fork — DECIDED** (`GRILL-LOCK-deployment-fork-2026-07-01.md`): attach + interface-guardrail; adversarial
   containment / wrap / "by-construction" shelved. Drives the 2B claim sweep.
2. **Interface guardrail — built + fixtured** (this is what makes the malicious findings adversarial-*only*): seats act
   ONLY through `submit()`/`project()`/`read()`; **raw store/config paths excluded from every seat tool surface** (a
   confused seat following a bad instruction has no *tool* to reach them); config is **conductor-owned, trusted-loaded
   once at startup, not in any seat tool surface**; **local-outbox-only external-send**. Fixture: a confused seat has no
   tool that writes store/config. *(NOT the old "lane write to a config path fails closed" adversarial-isolation
   fixture — that claim collapses to 2B.)*
3. **Fill-time authority / form rendering**: forbidden choices absent from the rendered form; courier validates
   constrained picks. The by-construction *claim* collapses (2B); the *mechanism* is the confusion-resistant interface
   contract and stays.
4. **Stand up conductor-core** (named owner): process/concurrency; a **single-threaded serialized commit loop**
   (`submit()` read-validate-append + `verify()` check-and-burn as one serialized critical section — kills the
   concurrency race by construction); **crash-atomic multi-file commit + recovery/reconciliation**; a uniform
   **internal-FAULT disposition** (held/fail-closed for authority records, never silent-accept, never brick,
   corrupt-record quarantine); store genesis/GC. Re-scope C3.6 to "inter-domain POLICY composition only."
5. **Resolve the Step-1/Step-2 phase dead-end**: step-gate the observe-owned `required_when` predicates on the observe
   layer's presence, OR a Step-1 conductor-side filler; strike the stale m-1 §13.2/§7 carry-forward.
6. **Pure-judgment A-floor**: mandatory HUMAN_GATE floor by (phase × record_kind) for A-capable work; `CEREMONY_TIER`
   monotonic-with-system-floor (a below-baseline pick auto-sets `gate_category=ceremony_downgrade`).
7. **Decision-② fail-closed**: class-conditional fail-closed on authority-class `self_reported` — fold into locked m-3
   §3.2/§8/§12.
8. **Fold the five operator decisions (①–⑤) + the renames into the locked docs** and re-verify byte-consistency before
   any doc is treated as build-ready.

### 2B — COLLAPSE to a global claim sweep + documented accepted-risks (adversarial security CLAIMS only — a doc change, NOT mechanism removal)
Relabel every "by-construction / structural / sole-writer / unbypassable / forgery-robust" → the D4 claim
("confusion-resistant; a malicious code-executing agent is explicitly out of scope"); record the D5 accepted-risks
(config / store-write / operator-FROM). Items that collapse here: adversarial DI-1/DI-2 bypass guarantees, malicious
direct store-append, malicious operator-FROM forgery, malicious config edit, malicious external-send bypass, the
three-rung ladder. **None of the 2A mechanisms collapse — only their malicious-lane claims.**

### 2C — MUST before their build step (not conductor-core-only-slice blockers UNLESS that slice ships the surface)
- **R2** opt-in `gate_referenceable:bool` + per-column `row_array` FieldSpec (must **not** ship with R2 under-specified);
  exclude single-family bucket-valued proxies.
- **Altitude-B per-row deviation grain**: move `deviated_observed`/`justified_deviation`/`deviation_reason_code` into
  the `routing_assignments` row; disposition per-row; block on any single (false,true).
- **Away-token sibling-burn + restart-binding**: decision-scoped resolution (burn all sibling nonces + a `decision_id
  unresolved` check) + a persisted seat-binding table — **stale approval / sibling-token reuse bite NORMAL operator
  flows**, so a hard gate before park/wake + the away-bridge, not disappeared. Plus away egress/ODB collisions.

---

## 3. The ranked FATAL findings (full detail in `tasks/*.output`)

*(Descriptive finding inventory — as originally found. The AUTHORITATIVE corrected re-baseline scope is §2/§5 +
`GRILL-LOCK-deployment-fork-2026-07-01.md`; where a finding below reads as a pre-Step-1 architecture MUST, defer to
the §2A/§2B/§2C bucketing.)*

1. **[FATAL] Policy config is lane-writable under attach — one Bash line is total, invisible governance bypass.**
   No integrity owner; defeats J1 + egress-fail-closed + the silent-deviation veto at once.
2. **[FATAL] Conductor-core is nobody's domain and has no design-of-record**, yet Step-1 IS "build the conductor
   core"; the capstone certified "build-ready" without checking the substrate exists.
3. **[FATAL] Attach invalidates I1 (sole-writer store) + unforgeable operator-FROM** — a lane appends a forged
   `AUTHORITY:merge-grant` record + INDEX row with an ordinary file write, never touching `submit()`.
4. **[FATAL] No serialization / crash-atomic invariant** — "replay-closed" and "authority-consume-once" false under
   concurrent MCP-server deployment; a restart strands every seat; dispositioned as a skippable SHOULD-spike.
5. **[FATAL] Step-1 form gate requires observe-owned fields whose only writer is Step-2** — every action-claim
   relay bounces uncompletably.
6. **[pre-build] R2 under-specified / untestable at schema grain (VP-corrected — NOT a confirmed leak)** — grammar
   keyed on the ghost name `selected_model`/`model_*`; live field `chosen_model` is prose-labeled non-referenceable
   but lacks an enforceable per-column `gate_referenceable` FieldSpec + fixtures. Fix by positive (opt-in)
   referenceability + row-column schemas + negative fixtures over `chosen_model` / model-identity / single-family
   bucket proxies.
7. **[FATAL] Pure-judgment A-categories + `CEREMONY_TIER` are floor-less agent-picks** — a seat self-absorbs an
   A-worthy decision (or shrinks its own required-set) and it never reaches the operator.
8. **[FATAL] Silent-deviation veto defeated at altitude-B** — record-level `deviated_observed` vs per-row
   `declared_deviated`; one honest deviation masks every silent one.
9. **[FATAL/systemic] Recorded decisions (esp. ② fail-closed) never folded into locked docs** — the build reads the
   fail-OPEN text; "MUST-gate satisfied / byte-consistency verified" unsound.
10. **[FATAL] Away-token replay forgery** — resolving one choice does not burn sibling tokens; a still-valid
    `approve` overturns a rejected/parked A-gate.
11. **[FATAL] Internal-fault disposition undefined** — the pipeline defines only {accepted, terminal-rejected}; a
    corrupt-record write under attach either bypasses the lineage gate (fail-open) or bricks the store (fail-closed);
    the design never picks.
12. **[SERIOUS→builds-on-FATAL] Observe registry runs repo-authored test code (conftest.py) inside the unsandboxed
    TCB** — "arbitrary lane code is rejected" and "thin trusted courier" both false; contradicts m-3's own safety claim.

---

## 4. What is genuinely sound (why re-baseline, not rewrite)

The skeleton is good; what failed is the **marketing of the guarantees, not the bones.** (1) The six-domain
decomposition + seam-reconciliation machinery genuinely work — they're what surfaced most of this as *tracked
carries* rather than production incidents. (2) The SR-11-7→SR-26-2 deviation-register is a defensible,
precisely-located novelty **once it moves to per-row grain** — nothing touched the concept. (3) observe-as-send +
observer-selects-the-invariant (DI-5) is the right anti-reward-hack shape; specific invariants are over-claimed
tamper-proof, the mechanism is not wrong. (4) The owner-typed FieldSpec / forbidden-option-absent-at-fill /
lineage-engine model is a strong primitive — it just can't be "structural" on a rented runtime. (5) **The
honest-fallback discipline itself (I2→confusion-resistant, I3→self_reported) is exactly right** — every FATAL is
"this pattern was not applied *here*." (6) Governance-first, model/provider-agnostic DNA is intact. **A re-baseline
away from build-ready, not a rewrite away.**

---

## 5. The re-baseline path (bounded — weeks, not restart)

**a. ✅ DONE — deployment fork DECIDED** (`GRILL-LOCK-deployment-fork-2026-07-01.md`): attach + interface-guardrail;
adversarial containment / wrap / "by-construction" shelved.
**b.** Stand up a **conductor-core design-of-record** with a named owner — owns the **§2A** set: the interface guardrail
+ fill-time-authority mechanisms (+ fixtures), the single-threaded serialized commit + crash-atomic commit + recovery +
internal-fault disposition + store genesis/GC, and the phase-split-aware required-set. Re-scope the C3.6 cert.
**c.** **Global claim sweep (§2B)** + **fold the five operator decisions + the two confused-agent guardrails**
(§2A.6 pure-judgment A-floor, §2A.7 decision-② fail-closed) into the locked docs; re-verify byte-consistency. *(No
config-integrity "isolation redesign" here — the confused-agent config fixture lives in conductor-core §2A.2; the
adversarial config claim collapses in §2B.)*
**d.** Fix the **§2C** items **at their build step** (R2 / altitude-B before the routing schema ships; away-token
before park/wake).
**e. Then** open Step-1 PLAN.

---

## 6. The process failure (owned)

The `FULL MUST-BEFORE-STEP-1 GATE SATISFIED` certification (CTO-driven, VP-co-signed) is **unsound** for two
concrete reasons: (1) the five operator decisions were **recorded but never folded into the domain docs a builder
reads** — so decision ②'s fail-closed is contradicted by still-locked fail-OPEN m-3 text; (2) the C3.6 capstone
**certified policy composition, not substrate readiness** — it could not have caught "conductor-core doesn't
exist" because it never looked. The measure-twice process was strong at catching *cross-domain drift* and blind to
*"is there a machine to run this on"* and *"do the docs actually contain the decisions."* This third pass existed
to close exactly that gap, and did. The certification is **RETRACTED** (see `READINESS-REGISTER.md`).

---

## 7. SHOULD-FIX (real defects, fix before their build step — not gate blockers)

- Give `self_reported` an **initial-release enforcing consumer** (a lineage rule that an authority parent must be
  `evidence_integrity=observed` or carry an operator waiver) — today its only discounter is the later-release benchmark.
- **Downgrade over-claimed tamper-resistance** (red→green proves a differential, not that the repro exercises the
  bug; `slot_in` classifier is adversary-gameable) — require repro vetting / strictest-invariant / fail-closed.
- **Relabel the egress content-scan** everywhere: destination control is by-construction (post-spike); CONTENT
  safety (regex/entropy) is best-effort and encoding-evadable — never "structurally unbypassable."
- **Away egress collisions as a set** (exempt the conductor token; render sensitive ODBs from evidence atoms or
  send a content-free "sealed gate awaits"; make `unattended` non-stranding).
- **Sweep the keying / row-parity / homing mismatches** (`routing` vs `routing_decision`; `bucket_binding_observed`
  + `rank1_recommended_bucket` have no m-2 home; posture axis has no writer post-`001537`; `bounced→rejected`
  never propagated to m-6 bucket-D → author-repair inbox dead-letters).
- **State + lint the `required_when ⇒ visible_when` coherence invariant.**
- **Assign the read-time migration APPLY step** (m-2 says migrate-at-read, m-1's projector "does not migrate").
- **Sandbox the observe registry checks** (unprivileged executor, no signing-key / store-handle).
- **Fix the `seat_archetype`/`authority_ceiling` confused-deputy** (planner-owned row vs F1 conductor-owned;
  add child ≤ spawner authority invariant).
- **ODB fallback for design/grill A-gates** (today route to a nonexistent Step-3/4 meeting lane); fold decision ④'s
  re-observe as a 6th away-token verify step.
- **State plainly that attach→srt-wrap is a RE-LAUNCH, not a bounded increment**, and move the wrap/launch decision
  into Step-1 (wake + store-isolation are needed at Step-1).

---

## Appendix — finding inventory & data

- Pass-1 (16 lenses): 66 findings raised → **37 survived** 3-verifier verification (11 FATAL, 25 SERIOUS, 1 MINOR;
  15 flagged pre-build-blocker). Full: `scratchpad/survivors37.json`, `tasks/wljxt3ugu.output`.
- Continuation (2 recovered lenses + gap-hunt): **3 + 8 survived.** Full: `tasks/wblshda94.output`.
- These are adversarial-agent findings, verified but not infallible; the load-bearing ones (config-lane-writable,
  conductor-core-unowned, R2 ghost-field, decisions-unfolded, phase-split dead-end) are concrete and
  CTO-spot-confirmed. The *convergence* (12 FATAL → ~8 root causes, multiple independent lenses hitting the same
  wrap-vs-attach inversion) is the strong signal.
