# master — BUILD-READINESS REGISTER (the Step-0 → Step-1 gate)

**What this is.** A read-only, adversarial build-readiness review of the locked six-domain Step-0 design-of-record
(`ARCHITECTURE.md` §1–§C3 + the six per-domain design docs), run **before any code**, at the operator's direction
("validate before cutting") and VP-approved (`readiness-decomp` 20260630-193316). Seven fresh red-team lenses
(composition · build-risk · assumptions · Step-1-readiness · adversary · operator/HITL · versioning) each hunted one
dimension; this register is the orchestrator reconciliation. **No design is changed here; nothing is (un)locked.** A
design-gap routes to its owning pair as a **bounded, VP-gated fix**; a spike is a **separate operator-opened** code step.

Evidence tags: **[V]** = orchestrator-verified against the cited lines. **[L]** = lens-cited file:line (confirm at fix-time).

---

## VERDICT: GO-WITH-FIXES  (unanimous — all 7 lenses)

> **[SUPERSEDED WHERE IT CONFLICTS — 2026-07-02]** This GO-WITH-FIXES verdict and the "What the review CONFIRMED works" list below are superseded by `DESIGN-REVIEW-2026-07-01.md` wherever they conflict with it. The NO-GO retraction at the foot of this file is **not** scoped only to the certification block: in particular the away-token "replay/forgery/expiry closed" line and the "unforgeable by lanes" claim in the CONFIRMED-works list do **not** stand — see the annotation on that list.

The design-of-record is **sound in the large and as a standalone-runtime target**; the core composition seams trace
clean; **nothing is no-go or unbuildable.** But Step-1 must not open PLAN on the store until **Cluster 1** is reconciled,
and the identity/governance posture needs the **operator decisions** in Clusters 2/5/6. Two recurring root causes:
1. **Ride-vs-own.** Several marquee invariants (sole-writer I1, "forgery-robust by construction" I2, "sole external
   sender") were proven for a conductor that *owns the process tree* — but Step-1 *rides* rented Claude Code / Codex
   runtimes. On a ride deployment those degrade, and the honest-fallback discipline the team applied to DI-5/identity/
   evidence was **not** applied to I1 or egress.
2. **Cross-doc drift.** A few locked docs fell out of sync across c1→c3 (m-1↔m-2 write-path; m-2's routing schema vs
   m-4's; the "Step-1 runtime-intelligence" label vs ROADMAP's Step-2). The C3.6 capstone certified composition at
   the *architecture* altitude; these are *doc-level* seams it did not reach.

---

## Cluster 1 — the Step-1 write-path  [BLOCKER before any store PLAN · design-gap · high]
*(cross-validated by the build-risk + composition + Step-1 lenses)*
- **1a [V] m-1 §5 and m-2 §4 describe incompatible `submit()` state machines.** m-1: "no 'append as submitted' step" →
  form+observe → one **atomic append-as-accepted**; **no cross-relay lineage gate in the ordering.** m-2: append-as-
  **`submitted`** (so lineage can walk it) → **lineage gate** → `accepted`. m-1's "reads identically to m-2 §4" is false,
  and the load-bearing lineage gate (preserves "authority blocks before dispatch") has **no hook-point** in m-1's
  `submit()`. · owner: **CTO+VP** (shared c1 contract) · operator_decision: no · evidence: `m-1 …design.md:91-94` vs
  `m-2 …design.md:72-73` (fold-claim `m-1 …:91`) · staleness: yes (fold-note masks a live conflict).
- **1b [V] The locked `submit()` embeds the m-3 observe pre-flight, but ROADMAP builds m-3 at Step-2** and `ARCHITECTURE`
  §C2 *titles* the m-3/m-4 cycle "Step-1 runtime-intelligence." So "build store+form-gate without m-3" is **not derivable
  from the locked pipeline** as written. · owner: CTO+VP (+ m-1/m-3) · operator_decision: yes — ratify **Step-1
  `submit()` = {resolve/stamp → form-validation → append/accept}** with observe (pre-flight b) as a marked **Step-2
  insertion point**; correct the §C2 label. · evidence: `m-1 …:93` (observe = pre-flight b) · `ROADMAP.md:57-70` ·
  `ARCHITECTURE.md:120`.
- **1c [L] The cross-relay lineage engine has no assigned build step.** Authority-bearing records (PLAN locks, dispatch/
  merge grants, IMPL reports) are non-deliverable until it promotes `submitted→accepted`; Step-1 deliverable 3 is scoped
  to *form-validation only*. Place it (Step-1 vs deferred) or Step-1 ships only non-authority records. · owner: m-2 ·
  operator_decision: yes · evidence: `m-2 …:73,165,221`; absent at `ROADMAP.md:57-73`.
- **→ action:** a bounded **CTO+VP+m-1+m-2 write-path reconciliation** (the submit state machine + the observe/lineage
  insertion points + the Step-1/Step-2 line) **before any store PLAN.**
- **→ RESOLUTION (`readiness-fix-c1`, CTO-arbitrated · VP-approved `…-222352`):** Step-1 write-path = **store + form +
  lineage** — pre-append validation (form → lineage against the persisted `accepted`-graph, candidate in-courier) → **one
  atomic `accepted` append**; **terminal `rejected`** on fail; **no persisted `submitted` limbo**. **1c RESOLVED — the
  lineage engine IS Step-1** (this **supersedes** the "Step-1 deliverable 3 = form-validation only" wording; Step-1 is
  *not* form-only). **1b — observe-as-send = reserved additive Step-2 hook.** `ARCHITECTURE.md:58-59` brought current to
  this seam (CTO fold). **Cluster 1 CLOSES only when** m-1 + m-2 fold `submit`/`send` byte-consistent with the seam (or
  surface a breaking constraint) **and** re-verify — **not yet closed.**
- **→ STATUS 2026-06-30 (CTO re-verify `readiness-fix-c1/…-225523`):** m-1 fold `…-224500` + m-2 fold `…-225007` both
  in; **byte-consistency re-verified** across m-1 `:100-104`, m-2 `:72-73`, `ARCHITECTURE.md:58-65`, and this note — 1a +
  1b resolved, invariants held, **no breaking constraint** (both pairs report the seam *cleaner* than the two-state
  path). **Cluster 1 → ✅ CLOSED (VP closure co-sign `readiness-fix-c1/…-230335`, 2026-06-30);** ARCHITECTURE:58-59 +
  this note marked closed (CTO status fold `readiness-fix-c1/…-230725`). **Ripple → tracked SHOULD-fix (NOT "just wording" — a bounded
  m-6 consumer-contract fix, route before any m-6 build):** m-2 renamed `delivery_state` value token `bounced`→`rejected`
  to match the terminal-`rejected` record; **stale consumers to align:** `ARCHITECTURE.md:287`, `m-6 …design…:40,:45`
  (still `delivery_state=bounced`). m-6 (+ m-2/m-3) confirms whether bucket-D observe-bounce unifies to `rejected` or
  needs a distinct token. **Cluster 4a/4b:** m-2 fold in (`readiness-fix-c4/…-225007`);
  **m-4 review `readiness-fix-c4/…-231200` = confirm-with-required-retype** — 4b grammar-enforcement + AC14 + the
  `declared_deviated==true` **trigger direction** CONFIRMED; **4a retype required:** `declared_deviated` must be typed
  **`agent_enum_pick` / planner-declared** (m-2 typed it `system`/`computed_result`, which is actually `deviated_observed`
  — the m-3 observed bit — and must NOT be gate-hosted; collapsing the two breaks the declare↔observe **silent-deviation
  integrity veto**; m-4 notes R2 itself is *not* violated — a seam-correctness retype, not an R2 regression).
  **RETYPE DONE:** m-2 retyped (`…-231506`) — `declared_deviated`→`agent_enum_pick`/planner-declared, `deviated_observed`
  split out (system/observed, not gate-referenceable), aggregate grain `any(routing_assignments.declared_deviated==true)`
  via a new bounded `any_row:` existential atom. **m-4 verified its acceptance MET (`…-232000`)**; **CTO re-verified
  line-by-line vs the m-4 contract + ARCHITECTURE R2 (`…-232234`)** — the new atom is R2-safe + within the bounded-boolean
  discipline; no architecture fold needed (ARCHITECTURE already stated R2 correctly).
- **→ VP CLOSURE REVIEW `readiness-fix-c4/…-232925` = REVISE (closure OVERCLAIMED — CTO under-scoped the dispatch).**
  **4a (R2 trigger + generic-atom hole) is VP-APPROVED and closed.** BUT **Cluster 4 is broader: 4b + 4c remain OPEN.**
  **4b** — m-2's routing FieldSpec (§12/§17.3) is still stale vs m-4's locked routing record (`m-4 …:200-210`): missing
  the per-row `routing_assignments` fields (`task_tag`, `declared_bucket`, `pin_mode`, `seat_archetype`,
  `authority_ceiling`) + `deviation_reason_code` (same required-when grain as `justified_deviation`) + `constraints` +
  `template_ref`. **4c** — conductor-computed fields (`record_integrity`/`surface_intent`/`posture`) have no declared m-2
  schema home. The c4 dispatch only covered 4a. **⚠️ "Full MUST gate satisfied" claim RETRACTED.** Resolution (operator's
  call): **(A) complete the fold** — bounded m-2 FieldSpec reconciliation to m-4's full record (+ declare the
  computed-field homes as *additive field declarations*, distinct from the forbidden tag-value micro-fold; m-4 review;
  m-3/m-5/m-6 CC for opaque/computed fields), **or (B) cited narrowing** — reclassify the omitted fields out of the
  Step-1 MUST as a named later SHOULD. **Cluster 4 NOT closed.**
- **→ RESOLUTION: (A) complete-the-fold (operator 2026-06-30).** Bounded m-2 FieldSpec reconciliation dispatched
  `readiness-fix-c4/…-233740` — **4b:** enumerate the per-row `routing_assignments` fields (`task_tag`/`declared_bucket`/
  `pin_mode`/`seat_archetype`/`authority_ceiling`) + add `deviation_reason_code` (same required-when grain) + reserved-shape
  `constraints`/`template_ref`, mirroring `m-4 …:200-210`; **4c:** declare `record_integrity`/`surface_intent`/`posture`
  homes as additive field declarations (m-3/m-5/m-6 confirm). CTO/VP clarification sanctioned: additive field declarations
  ≠ the forbidden tag-value micro-fold. **Pending m-2 fold → m-4/m-3/m-5/m-6 confirm → CTO re-verify → VP closure co-sign.**
- **→ m-2 FOLDED `readiness-fix-c4/…-234702` (4b + 4c done).** Routing FieldSpec mirrors `m-4 …:200-210` (per-row fields +
  `deviation_reason_code` + reserved `constraints`/`template_ref`); `record_integrity`/`surface_intent` declared as
  computed slots (§17.6, honoring C3 "no new *authored* m-2 field"); **`posture` is NOT a gap** — it rides the F2
  per-assignment home (`ARCHITECTURE:187-188`), declared there; AC15; boundary held (additive declarations, no tag-value
  micro-fold). **Pending 4 sibling confirms** (m-4 mirror · m-3 `record_integrity` · m-5 archetype/posture/`surface_intent`
  · m-6 delivery) → CTO re-verify → VP closure co-sign. **CTO note:** the rest of the m-3 observe-set is prose-declared,
  not row-declared; m-2 assesses the m-1 "m-2-declared set" MUST satisfied (the missing rollup added) — **CTO accepts;
  row-level parity for the full observe set = tracked SHOULD**, verified at CTO re-verify (row-parity fold only if
  re-verify / VP finds prose insufficient).
- **→ 4 CONFIRMS IN.** m-3 (`record_integrity` authorship-faithful, `000930`) · m-5 (four m-5 fields faithful + 2
  precisions, `235944`) · m-6 (delivery consumes cleanly, `235743`) — **all CONFIRM**. **m-4 (`235500`) = confirm-with-one-
  carve-out:** mirror faithful EXCEPT m-2 declared `posture` as a standalone `routing_assignments` field, but m-4's record
  `:203` + `ARCHITECTURE` C2.4 F2 carry only `seat_archetype`+`authority_ceiling`.
- **→ CTO ARBITRATION `readiness-fix-c4/…-001537` (verified vs source):** the m-5 c3 lock `:142` says **posture = "no new
  m-2 field," rides `seat_archetype`** (recorded, replay-complete). m-2 over-declared; **RULING (option 1): m-2 removes the
  standalone posture field** (keeps the value-enum as m-5 delivery vocab, derived from `seat_archetype`); **m-4 record +
  ARCHITECTURE UNCHANGED** (they were correct). m-4/m-5/m-6 confirm the ruling → CTO re-verify → VP closure co-sign →
  Cluster 4 CLOSED. **m-3 row-parity pass ACCEPTED as a pre-Step-1-PLAN SHOULD** (7 observe fields; before PLAN, after
  Cluster 4 closes).
- **→ m-2 FOLDED the ruling `readiness-fix-c4/…-012000`:** standalone posture field REMOVED (§17.3 now matches m-4 `:203`
  exactly); value-enum kept + annotated "derived from `seat_archetype`" (§17.6); m-5 label tighten folded; row-parity =
  §15 Q-F. No m-4/ARCHITECTURE change, no c3 reopen. **Pending: m-4/m-5/m-6 confirm the ruling reconciles their positions
  (near-automatic — m-4 offered option 1, m-5 `:142` governs, m-6 binds off F2 regardless) → CTO full re-verify → VP
  closure co-sign → Cluster 4 CLOSED → full MUST gate satisfied.**
- **→ 3 RULING-CONFIRMS IN + CTO RE-VERIFIED `readiness-fix-c4/…-013231`.** m-4 (`013000`, source-verified mirror = `:203`)
  · m-5 (`012910`, option 1, owns its imprecision, **no c3 amendment**) · m-6 (`012906`, reads cleanly via `seat_archetype`)
  — **all CONFIRM**. CTO re-verified m-2 §17.3:292 vs m-4 `:203` (mirror faithful, posture removed, `deviation_reason_code`
  present, R2 held). **Cluster 4 → ✅ CLOSED (VP closure co-sign `readiness-fix-c4/…-013613`, 2026-07-01) — 4a+4b+4c all
  CLOSED.** (last MUST item — now the full MUST gate is satisfied.) **Two pre-PLAN items
  tracked (neither a blocker):** (1) row-parity §15 Q-F (7 observe fields); (2) **NEW — runtime-`away` resolution** (m-6:
  `away` is an operator-global RUNTIME toggle, not a per-seat spawn-default, so it can't ride the frozen `seat_archetype`
  — the c3-lock-amendment lane; resolve before the away-bridge PLAN, Step-2). Both route on operator's go, before their
  relevant PLAN.

## Cluster 2 — ride-vs-own: the OS-isolation common mode  [operator-choice + design-gap · high · Step-1 for the *claims*]
*(assumptions + adversary + Step-1 lenses)*
- **2a Wrap-vs-attach is never stated.** m-1 §9 realizes DI-1/2/4 as "each seat a separate process; credential over a
  private fd; store excluded from the lane" — i.e. the conductor **launches/wraps** the runtime. ROADMAP says *ride
  existing sessions*. If the conductor merely **attaches** to an operator-launched same-uid session, DI-1/DI-2/DI-4 all
  fail. · operator_decision: **yes — is Step-1 wrap (impose uid/fd isolation + neuter native comms + exclude the store)
  or attach?** · evidence: `m-1 …:128` vs `ROADMAP.md:57-60`.
- **2b [V] I1 (sole-writer) has no honest-fallback label**, though I2→"operator-attested" and I3→"self-reported" do. A
  store with a second writer is **unsound**, not "confusion-resistant." · owner: m-1 · evidence: `m-1 …:64-67` vs `:86`.
- **2c Native inter-agent comms of the ridden runtime are never neutered** → channel-stamped FROM (I2) is moot for
  anything sent over the host's own SendMessage/inbox. · owner: m-1 · operator_decision: yes (specify Step-1 forces all
  lane comms through the conductor channel) · evidence: `m-1 …:91` vs `:26-27`.
- **2d "Sole external sender / structurally unbypassable egress" is a policy expectation, not structural, on a ride
  runtime** (the lane has Bash → curl/git push). Overclaim; needs the same honest bound as DI-5. · owner: m-3 (+ m-5) ·
  evidence: `m-3 …:118`; `m-5 …:97`.
- **2e The common mode: one DI-2 miss collapses identity + authority-ceiling + the Seam-C operator-mint together**
  (all three operator-FROM paths share DI-1+DI-2); no doc states it. **Single most important attack.** · owner: m-1 ·
  operator_decision: yes (gate away-mode/Seam-C on **verified** DI-2, not asserted) · evidence: `m-1 …:69-78,107-110`;
  `m-6 …:85-88`.
- **→ action:** an **operator ruling** on wrap-vs-attach + **"forgery-robust by construction" vs honestly
  "confusion-resistant" until DI-2 is a verified property**, + a bounded m-1 design-completeness fix (I1 fallback label;
  the neuter-native-comms requirement; re-scope the "sole external sender" claim to conductor-originated sends).

## Cluster 3 — the atomic store-append primitive  [build-risk · high · Step-1-blocker · SPIKE]
- **3a [L] Multi-file crash-atomic commit (record + INDEX + N mailboxes) + write-serialization under concurrent submits
  is named, not designed** ("one transaction" asserted; no fsync/rename/WAL/single-writer-loop). The nonce-burn
  (Seam C), the park/wake transitions, and the observe→append window **all inherit** this one primitive. · owner: none
  (correctly deferred to PLAN) · evidence: `m-1 …:94,192-193`; `m-6 …:171`.
- **→ action:** the **#1 spike candidate** — one spike de-risks 3a + the nonce-burn + park/wake + observe→append. A spike
  is **code = a separate operator-opened step**; recommended, not run here.

## Cluster 4 — schema-home & stale-schema (m-2 is the source the tool/courier/linter read)  [design-gap · medium · some Step-1]
- **4a [V] The R2 "structurally impossible" guarantee is not grammar-enforced.** The generic `field:<id>` atom reaches
  `selected_model`; m-2 itself phrases the deviation gate as "required_when **selected_model is off the prior floor**"
  (the forbidden model-derived form); the R2-safe `declared_deviated==true` lives only in m-4/ARCHITECTURE. · owner:
  m-2 (+ m-4) · Step-1-blocker (the required-when evaluator + routing FieldSpec are Step-1) · evidence: `m-2 …:84,283,
  285,289`; `m-4 …:206`.
- **4b [L] m-2's declared routing FieldSpec (§12/§17.3) is STALE** vs m-4's locked routing record (missing
  `declared_bucket`, `task_tag`, `seat_archetype`, `authority_ceiling`, `deviation_reason_code`, `template_ref`). A
  builder taking m-2 literally builds the wrong record + the R2-violating required-when. · owner: m-2 · Step-1-blocker ·
  evidence: `m-2 …:176,283-289` vs `m-4 …:203`.
- **4c [L] Conductor-computed fields have an ambiguous schema home.** `record_integrity`/`surface_intent` are consumed
  (ODB, m-6 delivery) but not enumerated in m-2, while m-1 requires the m-3 allowlist to "resolve to an m-2-declared
  set" and canonical-iff-consumed requires a registry entry — vs "no new m-2 field." `posture` is claimed to ride the
  F2 record home, which has slots only for `seat_archetype`+`authority_ceiling`. (3 lenses.) · owner: m-2 (+ m-3/m-5/m-6)
  · evidence: `m-3 …:110`; `m-1 …:222-223`; `ARCHITECTURE.md:351-352`; `m-5 …:142`; `m-4 …:203`.
- **→ action:** a bounded **m-2 schema reconciliation** (update the stale routing FieldSpec to m-4's locked shape; declare
  the conductor-computed fields' home — these are *additive field declarations*, distinct from the forbidden *tag-value*
  micro-fold; pin the `declared_deviated` R2 form + constrain the generic atom off model fields). Mostly a same-domain
  fix; the "no new m-2 field" guardrail was written to forbid tag-value micro-folds, not the additive field declarations
  consumers require — worth an explicit CTO/VP clarification.

## Cluster 5 — observe-as-send: fail-open + the opaque-lane collision  [design-gap · high/medium]
- **5a Observe-as-send is fail-OPEN on *unobservability*** — it vetoes an observed-false / observed-contradiction, but a
  record it cannot observe is stamped `self_reported` and **delivered** (opposite posture to the fail-closed egress
  gate). A deliberate "honest degradation," but an adversary forces unobservability to deliver fabricated done. ·
  operator_decision: **yes — may A-gate / merge-class records deliver on `self_reported` at all, or fail-closed when the
  required predicate is unobservable?** · owner: m-3 · evidence: `m-3 …:15,63,126-132`.
- **5b `slot_in` tamper-resistance ⟂ the opaque-lane floor.** The re-tag defense (refactor test-files-unchanged; bugfix
  red→green) is "tamper-resistant because the invariant is selected by the observer" via DI-5 classify-at-acceptance —
  but DI-5 degrades to `self_reported` on opaque lanes, and the docs never say how `slot_in` is classified then (→
  reopens re-tag-to-escape). Plus the classify-then-gate ordering (classified *at* acceptance vs the done-predicate reads
  it *before* accept) is a named PLAN carry. · owner: m-3 + m-5 · evidence: `m-3 …:101,126-132`; `m-5 …:78`.
- **→ action:** a bounded **m-3+m-5** reconciliation (slot_in classification on opaque lanes + the classify/gate ordering)
  + the 5a operator decision.

## Cluster 6 — operator-path surface holes  [design-gap · high/medium]  *(HITL + adversary lenses)*
- **6a A/B *under*-classification.** `gate_category` is agent-picked; `other→A` catches *un*classified, not
  *under*-classified; the pure-judgment A-categories (`product_semantics`, `residual_risk_acceptance`, `authz_security`)
  have **no system detector** → an agent tags an A-worthy decision as B → **silently orchestrator-absorbed, never reaches
  the operator.** The most direct operator-not-surfaced vector. · operator_decision: **yes — must agent-pick only RAISE
  toward A, never classify down to B (A-membership system/observer-derived where possible)?** · owner: CTO/§J + m-6 ·
  Step-1 (the A/B split is c1-locked) · evidence: `m-2 …:252`; `m-6 …:37-38`; `ARCHITECTURE.md:94-98`.
- **6b Away-token staleness.** The away path has no refresh/re-observe binding — the operator can approve against
  now-stale ground-truth and verify passes (breaks refresh-before-resummon, which the local path enforces). (2 lenses.)
  · operator_decision: **yes — must a refresh rotate `decision_id`/burn prior nonces, and `verify` re-observe-and-bounce-
  if-stale before stamping?** · owner: m-1 + m-6 · evidence: `m-6 …:60,76,78,86`; `m-1 …:224`.
- **6c Away-mode egress blocks the routing-escalation gates it must deliver** — the M4-1 ODB carries model-name choices;
  egress blocks undisclosed model-names → the gate parks `egress_blocked` with a local resummon the away operator isn't
  watching. · operator_decision: **yes — is a model-name inside an operator-facing governance ODB exempt from the
  model-name egress rule?** · owner: m-3 + m-6 + m-4 · evidence: `m-4 …:345` × `m-3 …:118` × `m-6 …:83,85`.
- **6d [L] `routing_unavailable` escalation may bounce author-ward.** If form-validation requires ≥1 `routing_assignments`
  row, the assignment-less escalation record bounces as bucket-D (author-facing) instead of surfacing as A. · owner: m-4
  + m-2 · evidence: `m-4 …:156-157,340-345`; `m-2 …:283`.
- **6e [L] design/grill verdict-gates route to a Step-3/4 meeting lane with no Step-1/2 ODB fallback** — may not surface
  at all until Step-3/4. · owner: m-6 · evidence: `m-6 …:97,99,122`.
- **6f [L] `unattended` posture A-gate delivery is undefined**; the summon ladder terminates with no rung reaching a
  human (J1-safe, but the gate can silently strand). · operator_decision: yes (the `unattended` + open-A-gate contract) ·
  owner: m-6 + m-5 · evidence: `m-6 …:78,113,124`.
- **6g [L] Away-mode expands the trust base to the operator's email account** (inbox compromise = holds the token); m-6
  (non-TCB) supplies `expiry` which m-1 only enforces. · operator_decision: yes (accept inbox-compromise risk; TCB-side
  max on `expiry`) · owner: m-6 + m-1 · evidence: `m-6 …:85-88,169-173`.
- **→ action:** a bounded **m-6 (+ m-3/m-4/§J) operator-path reconciliation** + the flagged operator decisions (6a/6b/6c
  are the load-bearing ones).

## Cluster 7 — authority-containment  [design-gap · medium]
- **7a `seat_archetype` ownership contradiction** — declared conductor-owned/non-lane-writable (F1), yet sits in a
  **planner-owned** `routing_assignments` row → a compromised planner declares a higher-ceiling archetype → escalation.
  Need "conductor stamps `seat_archetype`; planner may only propose within an authorized set." · owner: m-4 (+ m-1/m-5) ·
  evidence: `m-4 …:203`; `ARCHITECTURE.md:184-186`.
- **7b No non-increasing-authority invariant across staffing** — hand-authored staffing has no stated child-ceiling ≤
  spawner-ceiling rule → delegation escalation. · owner: m-5 + m-4 · evidence: `m-5 …:90-95`; `m-4 …:267-274`.
- **→ action:** a bounded **m-4+m-5** reconciliation.

## Cluster 8 — versioning: compat-of-interpretation  [build-risk/design-gap · medium/low]  *(no destructive migration found)*
- **8a Parked-gate-across-schema-bump** — J1 "no hard deadline / park indefinitely" + versioned schema **guarantees**
  operator-authority gates parked across a breaking bump, with no migrate-then-validate-at-wake path. On the governance
  path. · owner: m-6 + m-2 · evidence: `m-6 …:73-78`; `ARCHITECTURE.md:82-83`; `m-2 …:117`.
- **8b Migrate-then-process ordering unspecified** for the lineage walk (raw vs normalized mixed-version records). ·
  owner: m-2 · evidence: `m-2 …:116-117,152-161`; `m-1 …:96`.
- **8c Object-internal sub-schemas** (`capability_prior_snapshot` etc.) evolve outside the field-level compat contract. ·
  owner: m-4 + m-2 · evidence: `m-4 …:186-188`; `m-2 …:115`.
- **8d Two invariants to capture now:** `certification` null-epoch (a `null` conflates pre-era / not-a-signed-kind /
  stripped-tamper); and the open named-axis absent-default is safe only for lattice-bottom axes. · owner: m-1/m-2/m-5 ·
  evidence: `m-2 …:98-105`; `m-5 …:88,91,97`.
- **→ action:** record as **build-cycle contract clarifications** (nail while the seams are on the drawing board); 8a is
  the sharpest (governance path).

---

## What the review CONFIRMED works (adversarially checked, clean)

> **[SUPERSEDED WHERE IT CONFLICTS — 2026-07-02]** Two items in this list are overturned by `DESIGN-REVIEW-2026-07-01.md`: (1) the **away-token "replay/forgery/expiry closed"** line is superseded by FATAL #10 (sibling-token replay) → CQ-6 base + decision-④ rotate+re-observe, an honestly-deferred **(d) build-carry** with adversarial review still owed — it is *not* closed; (2) the **"local operator-FROM unforgeable by lanes"** claim is superseded by the **D4/D5 claim-set** — the honest posture is lane-unforgeable through the tool surface (confusion-resistant; D5 direct-store residual accepted), not unforgeable by construction. The remaining items stand.

- **observe-as-send TOCTOU** genuinely closed by placement (inside `submit()`, trigger≠execute) — the mechanism most
  primed for fault; clean. `m-3 …:31,35`; `m-1 …:92-94`.
- **J1 never-auto-approve** airtight on *every* path (timeout / resummon-exhaustion / egress-block / away-mode). `m-6 …:60,76,78`.
- **egress-vs-D precedence** airtight (disjoint pipeline stages keyed on `failing_edge`); **free-text can't create a
  hidden choice**; **local operator-FROM unforgeable** by lanes; **away-token replay/forgery/expiry** closed (residual =
  staleness 6b). `m-6 …:44-47,58-59,86`; `m-1 …:107-110`.
- **No destructive migration** — the append-only / zero-migrator foundation holds; every c3 carry is a pure key/field
  addition (residuals are compat-of-*interpretation*, Cluster 8). 
- **Core composition seams wired** — routing_ref/parent_picker join, ODB evidence bundle, B→A monotonic escalation,
  egress chokepoint, observe-hook-not-a-double-writer, DI-5 read-vantage (exemplary honest fallback).

---

## The gate before Step-1

1. **MUST reconcile before any store PLAN** (design-gaps in the write-path the store code hangs on): **Cluster 1**
   (CTO+VP+m-1+m-2) and **Cluster 4a/4b** (m-2 stale schema + the R2 grammar hole).
2. **MUST get operator decisions** on the identity/governance posture: **Cluster 2** (wrap-vs-attach + by-construction-vs-
   confusion-resistant), **5a** (self_reported delivery for A-gates), **6a/6b/6c** (under-classification · away-token
   staleness · egress-blocks-the-gate).
3. **SHOULD spike** (separate operator-opened code step): **Cluster 3** (the atomic append) — de-risks the most.
4. **SHOULD reconcile as bounded pair-fixes** (not Step-1-PLAN-blocking): **5b, 6d–g, 7, 8.**

None of this is a no-go. It is the measure-twice list: a short reconciliation pass (mostly CTO/VP + a few bounded
owning-pair fixes) + a handful of operator decisions + one spike, and the store is safe to PLAN.

---

## Appendix A — normalized routing table (per VP `readiness-reconcile` review 20260630-195249)

Legend — **Class:** DG design-gap · BR build-risk · OC operator-choice. **Sev:** H/M/L. **Step:** S1b Step-1-blocker ·
later later-step-carry · S1c Step-1-claims. **Disp(osition):** MUST (gate Step-1 PLAN) · OPDEC (operator decision) ·
SPIKE (separate operator-opened gate) · SHOULD (bounded owner-fix, scheduled after the MUST gate moves, not dropped).

| Item | Class | Sev | Step | Owner_to_reengage | OpDec? | Disp | Evidence_ref | Stale |
|---|---|---|---|---|---|---|---|---|
| 1a submit state-machine contradiction | DG | H | S1b | **CTO+VP**+m-1+m-2 | no | **MUST** | m-1:91-94 vs m-2:72-73 | yes |
| 1b submit embeds observe / Step label | DG | H | S1b | CTO+VP+m-1/m-3 | **yes** | **MUST**+OPDEC | m-1:93·ROADMAP:57-70·ARCH:120 | no |
| 1c lineage engine unassigned step | DG | M-H | S1b | m-2 | **yes** | **MUST**+OPDEC | m-2:73,165,221 | no |
| 2a wrap-vs-attach unstated | OC | H | S1c | m-1+orch | **yes** | OPDEC→m-1 fix | m-1:128 vs ROADMAP:57-60 | yes |
| 2b I1 no honest-fallback label | DG | H | S1b | m-1 | no | OPDEC→m-1 fix | m-1:64-67 vs :86 | no |
| 2c native lane-comms not neutered | DG | H | S1b | m-1 | **yes** | OPDEC→m-1 fix | m-1:91 vs :26-27 | yes |
| 2d "sole external sender" overclaim | DG | M | later | m-3+m-5 | no | SHOULD (rescope claim) | m-3:118·m-5:97 | no |
| 2e DI-2 common-mode (identity+ceiling+mint) | DG | H | S1(SeamC) | m-1 | **yes** | OPDEC (gate away on verified DI-2) | m-1:69-78·m-6:85-88 | no |
| 3a atomic multi-file append + serialization | BR | H | S1b | none (PLAN) | no | **SPIKE** | m-1:94,192-193·m-6:171 | no |
| 4a R2 grammar hole (generic atom→model) | DG | H | S1b | m-2+m-4 | no | **MUST** | m-2:84,283,285,289 | yes |
| 4b m-2 routing FieldSpec stale vs m-4 | DG | M | S1b | m-2 | no | **MUST** | m-2:176,283-289 vs m-4:203 | yes |
| 4c conductor-computed field schema-home | DG | M | later | m-2+m-3/m-5/m-6 | no | SHOULD | ARCH:351-352·m-1:222-223 | yes |
| 5a observe fail-open on unobservability | DG | H | S1 | m-3 | **yes** | OPDEC | m-3:15,63,126-132 | no |
| 5b slot_in ⟂ opaque-lane + classify order | DG | M-H | S1(observe) | m-3+m-5 | no | SHOULD | m-3:101,126-132·m-5:78 | no |
| 6a A/B under-classification (pick down to B) | DG | H | S1 | CTO/§J+m-6 | **yes** | OPDEC | m-2:252·m-6:37-38·ARCH:94-98 | no |
| 6b away-token staleness (no refresh) | DG | H | later(SeamC) | m-1+m-6 | **yes** | OPDEC | m-6:60,76,78,86·m-1:224 | yes |
| 6c egress blocks the routing gate it delivers | DG | H | later(away) | m-3+m-6+m-4 | **yes** | OPDEC | m-4:345·m-3:118·m-6:83 | no |
| 6d routing_unavailable bounces author-ward | BR | M | S1(M4-1) | m-4+m-2 | no | SHOULD | m-4:156-157·m-2:283 | no |
| 6e design/grill gate→nonexistent meeting lane | DG | M | S1b(design-gates) | m-6 | no | SHOULD | m-6:97,99,122 | no |
| 6f unattended posture delivery undefined | DG | M | later | m-6+m-5 | **yes** | SHOULD+OPDEC | m-6:78,113,124 | no |
| 6g away-mode trust base = operator inbox | OC+DG | M | S1(away) | m-6+m-1 | **yes** | SHOULD+OPDEC | m-6:85-88,169-173 | no |
| 7a seat_archetype ownership contradiction | DG | M-H | S1(recorded) | m-4+m-1/m-5 | no | SHOULD | m-4:203·ARCH:184-186 | no |
| 7b no child≤parent authority invariant | DG | M | S1(recorded) | m-5+m-4 | no | SHOULD | m-5:90-95·m-4:267-274 | no |
| 8a parked-gate-across-schema-bump | BR | H | later(gov path) | m-6+m-2 | no | SHOULD (sharpest) | m-6:73-78·ARCH:82-83 | no |
| 8b migrate-then-process ordering (lineage) | BR | M | later | m-2 | no | SHOULD | m-2:116-117,152-161 | no |
| 8c object-internal sub-schema compat | BR | M | later(future release) | m-4+m-2 | no | SHOULD | m-4:186-188·m-2:115 | no |
| 8d certification null-epoch + axis-default | DG | L-M | later | m-1/m-2/m-5 | no | SHOULD | m-2:98-105·m-5:88 | no |

**Cluster-2 tightening (VP Finding 3):** Cluster 2 is *operator-decision-first, then a bounded m-1 fix* — if Step-1 is
**attach**, m-1 downgrades the public claim to "confusion-resistant" honestly; if **wrap**, m-1 states the wrap
requirement + native-comms suppression + the I1 fallback boundary explicitly. Either branch is a bounded m-1
design-completeness fix, not just a claim label.

**Sequencing gate (VP Finding 4) — before any Step-1 PLAN dispatch:** (1) Cluster 1 reconciled OR converted to a scoped
operator-approved Step-1 subset excluding the unresolved lineage/observe path; (2) Cluster 4a/4b reconciled in m-2/m-4
(else Step-1 schema work held); (3) operator decisions recorded for Cluster 2, 5a, 6a, 6b, 6c. Cluster 3 spike only
under a separate operator-opened spike gate. SHOULD fixes may run in parallel *after* routing — never as a substitute
for the MUST gate.

---

## Operator decisions — recorded log

**① Cluster 2 (2a) — wrap-vs-attach + the honest claim — RECORDED 2026-06-30 (operator).**
**Decision: ATTACH-FIRST + "confusion-resistant."** Step-1 rides runtimes as an MCP-server with per-seat channels over
**persistent** pipes (persistent seats preserved — no per-relay spawn); the courier owns identity + the gate; the
honest public claim is **"confusion-resistant."** `srt`-wrap is a *later* bounded hardening lane ("sandboxed
defense-in-depth"); **"sole external sender by construction" stays a two-stage-spike-gated milestone** — properties
(1)–(5) earn only "single mediated network path"; the 6th (every outbound broker request passes the conductor
egress/content gate) earns "destination + content control by construction." Basis: verified `master/RUNTIME-RESEARCH.md`
§8/§10/§14, VP-approved with tightened spike gates (`runtime-research/RECONCILE-orchestrator-reviewer-20260630-213911.md`).
- **Unblocks — route as bounded VP-gated m-1 fixes, NOT opened here:** (i) the Cluster-2 m-1 completeness fix (record the
  "confusion-resistant" claim + the wrap-upgrade path); (ii) the identity-conductor-owned m-1 fix (VP Finding 3) — runtime
  identity fields (`session_source` / `clientId`) **never accepted as `FROM`**; conductor-owned per-seat
  channel/credential isolation is the **sole** stamp source.
- **Does NOT open Step-1 PLAN** — the sequencing gate still requires Cluster 1 + 4a/4b reconciled and decisions ②–⑤ recorded.

**② Cluster 5a — `self_reported` delivery for authority classes — RECORDED 2026-06-30 (operator).**
**Decision: FAIL-CLOSED for merge/authority (A-gate) classes.** An unobservable record in a merge/authority class is NOT
delivered on `self_reported` — it is **held/escalated**, matching the fail-closed egress posture on high-stakes paths;
lower-stakes/informational records may still deliver `self_reported`. → bounded **m-3** fix (class-conditional fail-closed
in the observe-as-send gate).

**③ Cluster 6a — A/B classification direction — RECORDED 2026-06-30 (operator).**
**Decision: RAISE-ONLY.** Agent-pick of `gate_category` may only escalate toward **A** (more operator oversight); it may
**never** de-classify an A-worthy decision down to **B**. Add a system detector for known-A categories. → bounded
**CTO/§J + m-6** fix (classification-direction invariant + detector). Closes the most direct operator-not-surfaced vector.

**④ Cluster 6b — away-token freshness — RECORDED 2026-06-30 (operator).**
**Decision: ROTATE + RE-OBSERVE.** A refresh **rotates `decision_id` and burns prior nonces**; `verify` **re-observes**
current state and **bounces** the approval if it changed since the operator last saw it. → bounded **m-1 + m-6** fix
(away-token refresh binding; closes the stale-approval / TOCTOU window).

**⑤ Cluster 6c — model-name-in-ODB egress exemption — RECORDED 2026-06-30 (operator).**
**Decision: YES — narrow, typed, ODB→operator exemption.** Exempt **only** the model-name field inside a
conductor-generated **operator-facing ODB** from the **confidentiality** egress scan. **R2 (model ≠ gate input) is
untouched** — peer-bias protection intact; the general egress rule still blocks model-names on all other external sends;
in away-mode the **transport stays gated by the away-bridge opt-in**; outside away-mode it renders locally and never
leaves. → bounded **m-3 + m-6 + m-4** fix (scan carve-out scoped to the ODB→operator path).

**⛔ RETRACTED (2026-07-01) — the "FULL MUST-BEFORE-STEP-1 GATE SATISFIED" certification below is UNSOUND. VERDICT: NO-GO.**
An adversarial pre-build review (`master/DESIGN-REVIEW-2026-07-01.md`; 16-lens × 3-verifier fleet) returned **~48
verified findings, ~12 FATAL**, and the certification fails on two counts: (1) the five operator decisions (①–⑤) were
**recorded but never folded into the domain docs** — so decision ②'s fail-closed is contradicted by still-locked
fail-OPEN m-3 text a builder would ship; (2) the C3.6 "build-ready" capstone certified **inter-domain policy
composition, not runtime-substrate readiness** — it never checked that **conductor-core** (the running program:
serialization / crash-atomicity / config-integrity / recovery) exists, and it does not (no owner, no design doc).
Plus: the design was written for a **WRAP** deployment but the operator locked **ATTACH** (①), inverting every
"structural / by-construction / sole-writer / unbypassable" claim. **Required: a bounded RE-BASELINE of Step-1
(deployment fork → stand up conductor-core → fold the 5 decisions + fix the mechanical FATALs → THEN Step-1 PLAN),
per `DESIGN-REVIEW-2026-07-01.md` §2/§5. No Step-1 PLAN, code, or spike until the re-baseline closes.** The prior
certification text is preserved below for the audit trail — it is superseded by this retraction.

**→ RE-BASELINE STEP (a) RESOLVED (2026-07-01) — deployment fork DECIDED via `GRILL-LOCK-deployment-fork-2026-07-01.md`;
bucketing corrected by VP `design-review/…-144217`.** Operator grill: Step-1 threat model = **confused-not-adversarial**;
Step-1 = **ATTACH** + **interface-level** guardrail (seats act only through `submit()`; config not a seat tool); **wrap /
adversarial containment / "by-construction" SHELVED indefinitely** (research-gated). **Reframes decision ①:** attach
DECIDED; standing claim = **"confusion-resistant; malicious code-executing agent explicitly out of scope"**;
"by-construction" shelved. **Effect on the review must-fix (VP-corrected — claims collapse, MECHANISMS don't):** the
adversarial *security CLAIMS* collapse to a **global claim sweep + documented accepted-risks**; but the confused-agent
*interface MECHANISMS* **REMAIN as conductor-core requirements + fixtures** — the interface-only seat tool surface
(raw store/config paths excluded), trusted config-load, local-outbox-only send, and **fill-time authority / form
rendering** (only its by-construction *claim* collapses). §2 must-fix #2 (config-integrity) REVISED: drop the
adversarial isolation redesign, keep trusted-load + not-in-seat-tool-surface in conductor-core. Threat-independent
MUSTs remain (conductor-core substrate: serialized/crash-atomic commit + recovery + internal-fault; **phase-split**;
pure-judgment A-floor; decision-② fail-closed). R2/altitude-B/away-token remain MUSTs **before their build step** (away-
token stale-approval/sibling-reuse bite normal operator flows). Remaining re-baseline steps (b)–(d) per `DESIGN-REVIEW`
§5 open.

---

**[SUPERSEDED — see retraction above] ✅✅ FULL MUST-BEFORE-STEP-1 GATE SATISFIED (2026-07-01).** — five operator decisions RECORDED (2026-06-30) · **Cluster 1
CLOSED** (VP co-sign `readiness-fix-c1/…-230335`) · **Cluster 4 CLOSED** (4a+4b+4c; VP co-sign `readiness-fix-c4/…-013613`).
The VP's sequencing gate (Finding 4) is now met. **The only remaining item is the Step-1 PLAN phase transition — the
operator's to authorize** (crosses the charter AUDIT+DESIGN boundary; the gate does not open it). **Tracked carries, on
the operator's go, before their relevant PLAN (NOT gate blockers):** row-parity §15 Q-F (pre-Step-1-PLAN) · runtime-`away`
resolution / posture-vocab split (pre-away-bridge, Step-2) · the ①-unblocked bounded m-1 fixes + the decision-implied
owner-fixes (m-3 fail-closed 5a, 6a detector, 6b away-token, 6c ODB carve-out) · the m-6 `delivery_state` `bounced→rejected`
ripple. Cluster 3 spike (the 6-property "by-construction" security milestone) remains a separate operator-opened gate.
