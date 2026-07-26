# master — CYCLE PLAYBOOK

**How the standing team actually ran, from ideation to the Cycle c1 lock — and the reusable cycle template
distilled from it.** This is the *process* record (the "how"), complementary to `CLAUDE.md` (the team
constitution — "who/what"), `ROADMAP.md` (sequencing — "what next"), and `ARCHITECTURE.md` (the design-of-record
— the "what we decided"). Built for future workflow building: Part A is the worked example (c1, from the real
relay trail); Part B distills it into a repeatable skeleton; Part C is the friction log — the lessons that cost
us something to learn.

Primary sources: `master/relays/INDEX.md` (the append-only event log) + the 80 relay files under
`master/relays/`. Everything below is reconstructed from that trail, not from memory.

---

## Part A — The c1 operation, as it actually ran

### A.0 Ideation → setup (2026-06-28 03:11 → 04:43)
The operator dropped a zipped frank design export into the cwd and framed the engagement: **a research +
high-level-design sprint** to design frank ("the conductor") — *using our own orchestrator protocol to design
itself*. Scope was fixed up front: **AUDIT + DESIGN only**, terminate at design-lock, docs in cwd never in
`pcode/`. We stood the team up as a **standing governing team** (not a transient sprint), modeled as a software
org: CTO (orchestrator-planner, me / Opus 4.8), VP-Eng (orchestrator-reviewer / GPT-5.5 high), and Sr-EM/
Tech-Fellow **pairs** `m-1..m-6` (each a `.planner` + adversarial `.implementer`, distinct models/lanes).

Setup ran through the same relay protocol it would later use for design:
- `master-org-decomp` — I proposed the org decomposition; **looped in the VP**; folded the review. (The
  operator pushed back on an early 4-tier recursive org; we landed flat + lean — *"you are the transport
  layer"*, since the operator hand-relays between independent sessions to run the stock protocol and design frank.)
- `master-roadmap` — drafted the v0→v1 roadmap, VP-revised across two rounds to "own-the-gate-first,"
  `C1_PRODUCT_SCOPE = B`. Persisted as the top-level `ROADMAP.md` (Step 0).
- Wrote the **charter** (`CLAUDE.md`, symlinked to `AGENTS.md`) so the out-of-skill team architecture
  **survives context compaction** and loads into every session in this cwd — a hard operator requirement.
- **Boot relays** to the two foundational pairs (`m-1`, `m-2`) — the TCB and the envelope lock first;
  consumers design against them.

### A.1 The c1 lifecycle (2026-06-28 12:59 → 2026-06-29 18:17)
One **cycle** = one initiative run across the domains. c1 = "audit the current state + design the frank Step-1
foundations." It ran in these phases (DISPATCH_IDs in `code`):

| # | Phase | Dispatches | What happened |
|---|---|---|---|
| 1 | **AUDIT** | `c1-audit-m-1`, `c1-audit-m-2` | Each pair audited its domain against the stock protocol + the cloned references (jcode, claude-code, agent-scripts, zellij). Planner + adversarial implementer each returned. |
| 2 | **AUDIT-RECONCILE** | `c1-reconcile` | I reconciled both audits, **looped in the VP** → joint **PROCEED-TO-DESIGN** (co-foundational; `GRILL_REQUIRED: yes`). |
| 3 | **DESIGN** | `c1-design-m-1`, `c1-design-m-2`, `c1-design-m2-m1-coord` | Pairs designed under the grill. Intra-pair design-review iterated each doc rev0→rev1. The cross-pair **coordination sub-thread** is where the load-bearing seam was found: system-filled `PARENT` (m-1) strengthens the m-2 lineage engine confusion-robust → forgery-robust. |
| 4 | **CONSUMER-REVIEW** | boot `m-3/m-4/m-6` then `c1-consumer-review-m-3/m-4/m-6` | Before locking, the *consuming* domains reviewed the interface sketches — catching writer-with-no-reader / reader-with-no-writer mismatches while they were still cheap to fix. m-3 opened a coordination relay to m-1. |
| 5 | **RECONCILE ×2** | `c1-design-reconcile`, `c1-consumer-reconcile` | Two VP-co-reviewed reconciliations (each multi-round) folded the design + consumer findings into a coherent contract. |
| 6 | **REFINE (rev2)** | `c1-refine-m-1`, `c1-refine-m-2` | Pairs folded consumer findings into rev2, **re-affirmed the shared contract** to each other (`COORD-REAFFIRM` / `COORD-CONCUR`), and ran a rev2 intra-pair design-review. (m-1 stalled; required an orchestrator **poke**.) |
| 7 | **JOINT-LOCK** | `c1-joint-lock` | I proposed locking m-1 + m-2 **together**; VP approved **pending operator ratification of two §J judgment items**. |
| 8 | **OPERATOR §J** | `c1-joint-lock` (cont.) | Operator ratified J1 (`hold_and_resummon`) + J2 (`gate_category`, operator-configurable), with two refinements; I folded them into `ARCHITECTURE.md`; VP returned the **close-confirm `approve`**. |
| 9 | **CLOSE** | `c1-joint-lock` (terminal) | Close declaration + pair stand-down; dashboard, reconcile ledger, and architecture spine sealed. |

### A.2 c1 by the numbers
- **~80 relays**, ~16 dispatch threads, **5 pairs** engaged (m-1/m-2 full design; m-3/m-4/m-6 consumer review)
  = 10 pair-seats + CTO + VP = **12 active seats**.
- **2 foundations locked** (m-1 Trust & Identity, m-2 Forms & Determinism), each iterated **rev0→rev1→rev2**.
- Wall-clock ~39 h (with an overnight gap); VP co-review at **every** authority-bearing gate.
- Output: 2 rev2 design docs + 1 m-2 audit doc + the integrated `ARCHITECTURE.md` + the locked §J policy.

---

## Part B — The reusable cycle template

Distilled skeleton for any future cycle. A cycle is **scoped to one phase-band** (c1 was AUDIT+DESIGN; a later
cycle might be DESIGN-only, or PLAN, or IMPL — set the band explicitly and hold it).

```
SETUP (once per team, or when adding domains)
  └─ boot VP → org-decomp (VP-reviewed) → roadmap (VP-reviewed) → charter (persist!) → boot the cycle's pairs

CYCLE c<n>
  1. AUDIT            orchestrator → pairs;  pairs (planner + adversarial implementer) return
  2. AUDIT-RECONCILE  orchestrator reconciles → VP review → PROCEED-TO-<next phase> (or revise)
  3. DESIGN           orchestrator → pairs (GRILL_REQUIRED); intra-pair review iterates; cross-pair COORD for seams
  4. CONSUMER-REVIEW  boot + dispatch the consuming domains to review interface sketches BEFORE lock
  5. RECONCILE        orchestrator folds design + consumer findings → VP co-review (multi-round OK)
  6. REFINE (rev2)    pairs fold findings, RE-AFFIRM the shared contract to each other, re-review
  7. JOINT-LOCK       orchestrator proposes lock → VP approves (may gate on operator judgment items)
  8. OPERATOR GATE    operator ratifies the judgment calls → orchestrator folds → VP close-confirm
  9. CLOSE            close declaration + pair stand-down + seal dashboard / ledger / architecture spine
```

**Invariant rules that made it work (carry these forward):**
1. **Loop in the VP before executing any authority-bearing orchestrator decision.** Every broad SET
   (AUDIT / DESIGN / REVIEW-FOLD / MERGE-GATE / delegated-PLAN / override-IMPL) and reconciliation **CCs the
   VP**. The operator corrected me the one time I jumped ahead — the discipline is non-negotiable.
2. **Foundations lock first; consumers design against them.** Lock the interfaces the whole graph depends on
   (m-1 store API, m-2 schema) before the consumers commit. Co-foundational domains (a load-bearing seam
   between them) **lock together**, not sequentially.
3. **Consumer review precedes design-lock.** Cheapest place to catch a writer with no reader (or vice versa).
4. **Phase discipline is explicit and enforced.** State the cycle's phase-band; reject self-advancement
   (I corrected pairs' "READY FOR PROCEED-TO-PLAN" framing — there was no PLAN in c1).
5. **The orchestrator actively polls.** Independent sessions stall silently; **poke** a quiet seat rather than
   waiting (m-1 needed one in refine).
6. **Adjudicate intra-pair splits on evidence, not seniority of role.** Where `.planner` and `.implementer`
   disagreed (m-3, m-6), the decision went to whichever did the rigorous analysis.
7. **Every substantive handoff is a lint-clean file relay.** No proxy-authoring another seat's `FROM`; the
   operator is the trusted transport between sessions.

### B.1 Relay mechanics (the substrate)
- **Headers (every relay):** `ROLE` · `PHASE` · `AUTHORITY` · `DISPATCH_ID` · `PARENT_DISPATCH_ID` · `RUN_ID`
  (`master`) · `CEREMONY_TIER` · `EVIDENCE_TARGET` · `HUMAN_GATE_REQUIRED` · `FROM` · `TO` · `CC`. Body, then
  `ACTIONS_GIT_REF` + `FINAL_GIT_STATUS_SHORT` (here: `unavailable — cwd is not a git repo`).
- **DISPATCH_ID:** boot = `master-boot-<seat>`; work = `<cycle>-<phase>-<pair>` (e.g. `c1-design-m-1`);
  cross-pair seams get their own coord thread (`c1-design-m2-m1-coord`).
- **File naming:** `<PHASE>-<ROLE>-<YYYYMMDD-HHMMSS>.md` under `master/relays/<DISPATCH_ID>/`.
- **Lint every substantive relay before handoff:**
  `python3 /Users/jack/.claude/skills/tools/relay-lint.py <file>` (exit 0 = OK). Scoped lineage check:
  `relay-lint.py --relay-root master/relays/<DISPATCH_ID>`.
- **INDEX.md** (`master/relays/INDEX.md`) — append-only routing log, **lint-exempt**. Columns:
  `time | phase | role | dispatch | parent | from | to | cc | status | file`.
- **Role skills** (installed): `orchestrator-planner`, `orchestrator-reviewer`, `agent-pair-planner`,
  `agent-pair-implementer`, `design-grill`. Each carries `protocol.md` (the upstream relay/authority/evidence rules).

---

## Part C — Friction log (lessons that cost us something)

Carry these into future cycles and into the eventual **m-5 workflow/archetype** design — they are exactly the
edges a workflow engine should automate away.

- **`relay-lint` tripwires.** `merge = A` was parsed as a `merge=<sha>` git claim; a bare "backfill" read as a
  false action-claim demanding `ACTIONS_GIT_REF`. **Fixes:** use underscores (`merge_to_protected`,
  `merge_decision`); avoid `<word> = <token>` phrasings near reserved verbs; always set
  `ACTIONS_GIT_REF` / `FINAL_GIT_STATUS_SHORT` honestly (`none` / `unavailable — cwd is not a git repo`).
- **INDEX.md is volatile.** Pairs append asynchronously; it tripped "file modified since read" repeatedly.
  **Fix:** re-read the *tail* immediately before appending; append after the current last row. It is
  **lint-exempt** — filter it out of `--relay-root` runs (`grep -v INDEX.md`).
- **Addressing deviation.** `m-2.implementer` once addressed `TO: master.orchestrator-planner` instead of the
  pair partner (DISPATCH_ID `c1-audit-m-2-implementer`). Non-blocking — `PARENT` was preserved, reconciled
  transparently. **Lesson:** the lineage held because `PARENT_DISPATCH_ID` was correct; that field is the
  load-bearing one.
- **Sessions stall silently.** No "I'm done" signal ≠ done. The orchestrator must **poke** (m-1 in refine).
- **Don't over-format for the operator.** `AskUserQuestion` was rejected; the operator prefers **free-form
  prose** with recommendations + alternatives + rejections. Present judgment calls as prose, not menus.
- **Toolchain path moved.** The operator changed workstations (Linux → macOS) mid-engagement; the linter is at
  `/Users/jack/.claude/skills/tools/relay-lint.py` (with a `.codex` mirror). Don't hardcode a home path in docs.
- **"Explore, then decide" on operator judgment items.** For the §J calls the operator wanted **options with
  recommendations** first, then ratified — don't pre-bake a single answer into a lock; leave a clean gate.
- **`relay-lint` lineage resolver breaks on shared dispatch ids (BUG, shelved 2026-07-11; reported by an
  agent-pair Planner seat on protocol v2.8.8 + `tools/relay-lint.py`, via the operator — a recurring
  deterministic false positive that blocks delegated dispatch).** `PARENT_DISPATCH_ID` is resolved by
  `one_by_id()` (`relay-lint.py:1157-1161`), which returns the **earliest** relay bearing the id
  (`sorted(items)[0]`). But in real runs a dispatch id is a *thread* id shared by several legitimate relays
  (the orchestrator's PROCEED-TO-PLAN, the pair-Planner PLAN + its revisions, scope-diff/deviation/amendment
  relays) — and the earliest is almost always the orchestrator's PROCEED. The delegated-`DISPATCH IMPL` lineage
  walk then validates the pair checks against an orchestrator-authored relay and fires two false errors
  ("PLAN-REVIEW must review the pair Planner's PLAN, not a CC'd orchestrator dispatch"; "pair-Planner PLAN must
  address the Implementer in TO for review") even when a fully conforming pair-Planner PLAN exists in the same
  thread and is named by the review's own `IN_REPLY_TO`. **Observed:** twice in one sprint, once per pair, each
  at the delegated-dispatch moment — one contributed to a waived lineage-dirt set; the other fully blocked a
  dispatch and forced a second operator waiver. Any team using the thread-id convention hits it on **every**
  delegated dispatch. **Why it matters:** the gate exists to make dispatch authority mechanically checkable; a
  deterministic false positive at exactly that moment converts a satisfied gate into a mandatory human
  interrupt and **normalizes waivers**, eroding the gate for the true positives it was built to catch.
  **Suggested fixes (any one suffices; first is smallest):** (1) resolve by **required shape, not age** — among
  relays sharing the parent id, prefer the candidate matching the check being applied (`PHASE: PLAN`,
  `FROM <owner>.planner`, Implementer in `TO`), erroring only if *no* candidate matches; (2) use the child's
  `IN_REPLY_TO` file path as a disambiguating hint when present (the approving review already points at the
  exact PLAN file); (3) require unique dispatch ids per relay + a separate thread-id field — a breaking change
  to existing trees, so (1)/(2) preferred. **Repro sketch:** a thread with id X holding, in timestamp order, an
  orchestrator PROCEED (`FROM: orchestrator`) → a pair-Planner PLAN (`FROM: a.planner`, `TO: a.implementer`) →
  an approving PLAN-REVIEW (`PARENT_DISPATCH_ID: X`) → a dispatch relay parented to the review's unique id;
  `relay-lint.py --relay-root <dir>` fires both errors despite the conforming PLAN. *(Interim mitigation
  already conventionalized in-run: pair loops under a master thread mint unique sub-`DISPATCH_ID`s per leg.)*

---

## Part D — What this implies for the product (forward)
The conductor is meant to *be* this protocol as software. This playbook is therefore also a **requirements
trace** for it: the cycle skeleton (Part B) is the **m-5 workflow/archetype** expansion-slot preset; the
loop-in-the-VP / consumer-review-before-lock / poke-the-stall rules are **gate-set + scheduler** behavior
(m-6 + m-3); and every Part C friction is a place the engine should remove a human/lint round-trip. When m-5
designs archetypes, start here.

**Owed-item tracking is a first-class store capability, not a checklist (c6 forward input — build in Step-1).** c6 surfaced a
failure class the relay/INDEX substrate does not catch: an actionable item **flagged *inside* a delivered relay body** (a
cross-domain mirror, an owed fixture) can silently drop at integration **even though the relay was delivered AND read** — because
tracking is at the *relay* grain, not the *item* grain (the dropped m-4 routing-mirror is the motivating case). The fix is **not a
mutable checklist** (that fights the append-only grain and re-introduces the same drift — someone forgets to tick the box). It is the
**m-7 `held`/burn pattern generalized:** an owed item is a **typed record appended when raised**, **closed by a *subsequent* record
referencing it** (never a mutation), and its status is a **projection** — `open = owed-record with no disposition-record`. A
cycle-close gate then *projects* "open owed items" and refuses to close while any remain. So the checklist is a **query, not a
document you maintain**. Step-1 should build this as a first-class store capability (an owed-item `record_kind` + a `project()` over
it), making silent-drop impossible **for a *recorded* owed-item** (the projection guards recorded records — `open = owed-record with no disposition-record`; *materializing* the record stays an intake/triage step, not replaced by the projection) — exactly as the store already does for `held`. Deliberately
deferred as a Step-1 build, **not** a manual bandaid: the design-phase fan-out that exposed this (7 parallel pairs) does not recur at
Step-1's serial one-section-at-a-time implementation grain, so the orchestrator's own working context is a sufficient checklist until
the real capability ships — which is then dogfooded immediately.

---

## Part A.2 — Cycle c2, as it ran (Observation & Evidence + Routing & Policy)

The second cycle ran the same skeleton (Part B) and confirmed it, with three **new patterns** worth keeping:

1. **Early-seeded COORD thread.** The orchestrator seeded the m-3↔m-4 seam thread *with an agenda* at design
   dispatch, so both designs could cite its current state. The seam reconciled both sides early and held through
   the lock — the load-bearing R2-preserving "block the dishonesty, never the deviation" resolution came out of
   it. **Seed cross-domain seams early, with the questions written down.**
2. **Narrow mid-cycle engagement (m-5).** A scope addition (GL-4 templates) + a forward feature (the interjection
   side-question) pulled a *future* domain (m-5) into c2 — but **bounded**: a hard "must not become a full
   archetype-system design; the lock RESERVES your full ownership to c3" guardrail, written into both the boot
   and the dispatch and re-stated by the VP. It satisfied the lock prerequisite AND delivered initial-release surface
   without pulling c3 forward. **A domain can join a cycle for a reserved slice — bound it explicitly and
   reserve the rest.**
3. **Fold-confirm round (rev2-lite).** Consumer-lens findings were routed back to the design-complete pairs as a
   *bounded additive fold* (not a redesign), implementer-re-approved, with a "anything stronger → flagged
   micro-fold to the foundation owner" tripwire. Mirrors c1's refine→lock at lower weight.

The lock itself took a **revise → de-lock → co-sign** loop: the VP caught a real ambiguity — the design-of-record
*named* candidate archetype values right next to the "stays opaque / no concrete values" claim, so a reader could
treat examples as the locked tag-space. Fixed by making "named values are non-locking examples; the lock binds
only the opaque atoms + the c3 reservation" the operative rule. By the numbers: 5 pairs/seats engaged across c1+c2;
m-3+m-4 each iterated audit→design→fold; the whole c2 ran VP-gated end to end.

### New friction (folds into Part C)

- **First INTERTEAM relay (2026-07-13, operator-directed; precedent recorded).** A fork of the frank master seat, driven by the operator, authored `pdc:/master/relays/interteam-frank-advisory/SITREP-orchestrator-planner-20260713-022100.md` — `FROM: frank.orchestrator-planner` (deliberately non-masquerading: the cross-team spelling can never read as pdc's own master), TO pdc's `master.orchestrator-planner`, `PARENT_DISPATCH_ID: none` with the operator named as directing transport, ZERO authority claimed in the receiving protocol (advisory only; adoption routes through THEIR deviation register). Payload: the measured hand-relay survey (928 pdc relays, 81% cross-seat; 677 design-side vs 32 implementation = 48:1) + four design-churn suggestions; the tier-flattening suggestion recorded as CONSIDERED-AND-DROPPED by the operator (pdc's subteam tier validated as the s1–s6 frank shape). Convention seeds if it recurs: non-masquerading FROM · zero-authority advisory grade · filed in the receiving team's root only (no sender-side copy) · absolute-path citations readable from the receiver's cwd. Product note: this is the federation seam (multi-team frank) exercised by hand, once.


- **Per-task review legs vs hand-relay latency (s10, 2026-07-13 — the B11 cadence ruling).** Three slices of evidence: s8 ran per-task reviews with heavy master/owner round-trips (caught real classes, but every leg cost the operator two hand-carries); s10's first day spent more wall-clock on review-relay transport than on code. Operator repriced: straight-through build + one end review + batched owner confirms (stop-on-contradiction). The friction is the TRANSPORT, not the review content — the exact function (typed relay carriage, wake-on-reply, no human pump) the product's park/wake + observe layers exist to automate. When frank carries its own team's relays, the per-task cadence becomes nearly free and this ruling should be revisited — the cleanest possible dogfood metric: does the conductor make its own build reviews affordable again?

- **The owner confirm catches what the review panel cannot (s11, 2026-07-14 — catch-ledger #4, the complement to the s10 datum).** T9's `resummon_cadence` config key landed on m-7's r13-locked engine-v3 schema surface with **no version bump and no owner routing** — an r13 violation (schema-surface change ⇒ version bump) and a condition-(c) violation (locked-contract change routes to owner + master, never silent). It survived **three plan-review rounds AND a four-lens end panel** (contracts · refactor-preservation · test-honesty · in-fence invariants) and was caught only by the **batched owner-fidelity confirm** — because every review lens checks the slice against *itself*, and a cross-domain lock violation is structurally invisible to a self-check; only the seat that *owns the violated contract* can see it. s10 taught us the end review catches what a green battery cannot; s11 adds: the owner confirm catches what the review panel cannot. **Binding discipline refinement:** a fence row licenses a FILE; it never substitutes for the OWNER PATH on a locked contract living inside that file. When a licensed file carries another domain's lock, the PLAN must name the lock and route the owner countersign *before* the edit — discovering it at exit costs a voided merge recommendation + a fold loop. (Ruled (a) re-home at v4, folded on-branch — `s11-build-escalate-config-lock/RECONCILE-…-170510`. Not a PROTOCOL-DEVIATION: the ruling *conforms* to r13; the silent landing was a transient defect caught pre-merge.) The miss is the same TRANSPORT tax as the B11 entry from the other side — the cheap correct routing (COORD → owner countersign) that `lane_vcs` used the same week was available and skipped because hand-relaying it felt expensive; the product's owner-notify-on-locked-surface is exactly this automated.

- **Pairs may skip the intra-pair reconcile.** In c2's consumer-lens round, m-5 and m-6 each filed two
  independent passes but *no* pair-reconcile relay; the orchestrator assessed convergence instead. The VP flagged
  it as "not ideal." **Fix:** the dispatch should make the pair-reconcile relay an explicit deliverable, or the
  orchestrator reconciles and says so.
- **Lock-text can both reserve and name.** Illustrative value lists next to a "nothing is locked" claim create a
  design-lock ambiguity. **Fix:** when a cycle reserves a space to a later owner, mark every example value as
  non-locking candidate vocabulary, and state the operative lock rule (the opaque atom + the reservation) once.
- **Operator directives via chat aren't channel-stamped.** GL-4 / m-5 scope came through the session, not a
  `FROM:operator` relay. The sanctioned resolution: cite "operator-directed by current session context" (or the
  operator drops a stamped relay). The team applied its own forgery-robust-identity principle to itself.
- **Orchestrator-applied edit to a pair's doc = operator-directed + self-attributed only.** One trivial lock-text
  de-lock was applied by the orchestrator under explicit operator direction, transparently self-attributed, and
  VP-accepted **for that closure only — not a precedent.** Substantive changes to a domain's semantics still go
  through the owning pair's relay/review path.
- **INDEX async-append gaps.** Pair sessions (esp. the `.codex` lane on a different lint path) sometimes don't log
  their own rows, or log them out of order. **Fix:** the orchestrator runs the completeness + dangling check at
  each reconcile and backfills.

---

## Part A.3 — Cycle c3, as it ran (Workflows & Archetypes + Human Surface & Scheduler)  ✦ Step-0 complete

The final Step-0 cycle locked the last two domains and completed the six-domain design-of-record. Same skeleton;
**two new patterns** that only a *terminal* cycle needs, plus a clean execution of a mechanism the roadmap reserved:

1. **No downstream consumer-lens → conditional-upstream-check + a blocking capstone.** c3's domains ARE the last
   design layer (the only downstream, m-7..m-12, doesn't exist as seats). Instead of booting a premature reviewer,
   the substitute was: (a) re-engage a *specific* locked upstream owner only when a real question arises, and (b) a
   blocking pre-close **integration-completeness capstone** (consume-graph acyclic + writer-backed; seams close;
   locked invariants intact; deferrals recorded) folded into the lock co-sign. **A terminal cycle certifies the whole
   composition instead of reviewing downstream.**
2. **The conditional-upstream-contract-check, executed (Seam C).** m-6's away-mode bridge raised a real question
   against m-1's locked TCB. The orchestrator routed **one bounded question** (A: m-1-owns-mint vs B: m-6-owns-bridge)
   to compacted/stood-down m-1, which re-oriented from its locked design doc + boot relay and answered **A** —
   *proving* it forced by two locked invariants (DI-1 sole-writer store; DI-2 key custody) and **additive** via an
   already-reserved extension seam (`certification`, present-but-null). The away-cell was **lock-blocking but not
   dispatch-blocking** (everything else locked; only that cell held). **A locked domain can answer a new cross-domain
   question without reopening — if it reserved the extension point.**

The lock took a **revise → correct → co-sign** loop (like c2): the VP caught a narrow **stale-status contradiction**
— m-6's doc said the away-cell was both "locked" and "held/pending" after the Seam-C fold. m-6 fixed it (a full-doc
sweep found + cleared two more of the same class) and its implementer re-approved; the orchestrator re-emitted the
narrow lock and the VP co-signed the lock **and** the capstone. **No orchestrator proxy-edit** — ownership held.

### New friction (folds into Part C)
- **The F4 pair-reconcile gap, enforced this time.** m-5 filed two independent audit passes but no pair-reconcile
  (the recurring `.codex`-lane stop-at-two-passes). The VP's F4 bar (two artifacts **plus** reconciliation, OR one
  reconciled artifact) required it, so the orchestrator **held the audit-reconcile gate** and bounced it back — m-5
  filed two convergent reconciles. **Fix (applied):** enforce F4 explicitly at audit-reconcile; two-passes ≠ reconciled.
- **Don't write a "reconciled/locked" status ahead of the gate.** The orchestrator refreshed the m-5 charter to
  "RECONCILED (via orchestrator synthesis)" *before* m-5 actually reconciled — contradicting its own nudge. The VP
  caught it (revise). **Fix:** the durable status must match the live decision; never front-run a gate in a status doc.
- **Declare-before-bind requires reading the current COORD state first.** m-5 wrote a "FINAL" human-mode proposal
  without reading m-6's binding-confirm, crossing the seam; caught by m-5.implementer, resolved by m-5 conforming +
  retracting. **Fix:** in a declare-before-bind COORD, cite the thread's current state before declaring "final."
- **When folding a resolved cell, sweep the whole doc for its stale status.** m-6 folded Seam C in the primary
  sections but left held/pending language in §8/§10/§11/§12. **Fix:** a resolved-cell fold ends with a grep-sweep for
  the cell's old status strings across the whole doc (m-6 did exactly this on the re-fix — adopt it as the default).
- **Reserve extension seams from day one (product, Part D).** Seam C stayed additive *only because* m-1 had reserved
  a present-but-null `certification` field for the channel-stamp-unavailable case. Reserve versioned, present-but-null
  extension points so future cross-domain activations are additive, not reopens.

---

## Part A.4–A.6 — Beyond Step-0: conductor-core, honesty, re-review ✦ Step-0.5 (re-baseline)

*(Bridge — c4/c5 ran between A.3 and c6; captured compactly, then c6 in full as the freshest reusable pattern. Trigger: a 2026-07-01 adversarial pre-build review returned **NO-GO** — the design assumed a WRAP deployment while the operator had locked ATTACH, and the runtime substrate was nobody's domain. The re-baseline is bounded repair, not restart.)*

- **c4 (conductor-core, m-7)** — stood up the runtime-substrate domain the NO-GO exposed as nobody's: audit → design → **8-CQ gate** (full-pair) → design-lock → close. A **new seat added mid-flight** (m-7); the substrate **hosts-but-doesn't-re-own** the six policy contracts. Its one licensed by-construction claim = the serialized-loop two-honest-seats double-accept kill.
- **c5 (global claim-sweep)** — relabeled every adversarial-strength claim (sole-writer / forgery-robust / unforgeable / …) to **confusion-resistant + D5 residual** across the 7 design docs; folded operator decisions ③/④/⑤. Certified "honest end-to-end" — a certification c6 found **scoped too narrowly** (see friction).

### A.6 Cycle c6 — adversarial RE-REVIEW of the design-of-record + doc-only cleanup (2026-07-02)

**Shape (a new reusable pattern — a *review-and-repair* cycle, distinct from the build cycles A.1/B):**
1. **Re-review** — a Workflow: coarse lanes (10 = 7 domain + 3 cross-cutting), single adversarial verifier per finding, every lane briefed on the *locked boundary* so it reviews the CURRENT design instead of re-running the prior verdict. → **90 findings, 0 FATAL, CONDITIONAL-GO**; the NO-GO **structurally discharged**.
2. **Decompose + VP sign-off** — route findings by owner (domain-local → pairs; seams + governance surfaces → CTO); the VP concurs the verdict + routing **before any fix** (here it caught an over-broad token-semantics resolution and amended it).
3. **Hybrid apply** — CTO single-hands the seams + gov surfaces + token-convergence + mechanical sweeps via **file-partitioned agents (one writer per file)**; each agent **declines-don't-guess** anything needing domain judgment (8 declined → re-routed to pairs). **Snapshot the tree first** (no VCS).
4. **VP reviews the CTO apply half** — adversarial, on a ground-truth diff; caught a cross-doc convergence miss + a corrupted evidence artifact.
5. **Pair fan-out** — 46 findings to 7 pairs, each relay carrying its findings + the canonical resolution + "CTO-already-applied, don't redo" notes + seam-partner CCs; pairs run planner-fix → implementer-approve; seam pairs exchange **COORD** relays.
6. **CTO integration** — close the owed cross-surface tails the pairs flag to CTO surfaces (here: 4 §C4 edits); **enroll owed items in a build-carry ledger**.
7. **Verify sweep + close (VP co-sign)** — the orchestrator's *own* claim/token/seam sweep over the live state, then VP co-sign (revised once on close-record accounting).

**By the numbers:** 90 findings = **44 CTO + 45 pair (7/7 approved) + 1 subsumed**; 18 docs, +433/−177; re-review ≈101 agents, apply ≈14 agents; **3 VP revise→approve loops** (decomp amendment, apply-half defects, close-accounting).

**New friction (folds into Part C):**
- **The sweep's scope was the bug.** c5's claim-sweep was scoped to the 7 design docs and certified "honest end-to-end" — but the boot-path docs a builder reads *first* (`CLAUDE.md`, domain READMEs, `master/README.md`, `RECONCILE.md`) still carried the retired claims. **Fix:** the standing claim-sweep + byte-consistency net now includes the charter/navigational docs, not just the design-of-record.
- **Evidence artifacts must be self-consistent + reproducible.** The review diff came out corrupted (`diff` shell-aliased to `--color` → ANSI in headers; written *inside* the diffed tree → self-inclusion), and the close record's headline count (90 ≠ 52+38+4) and diff stats (+398 vs actual +433) didn't match the files. **Fix:** generate diffs with `command diff` over an *explicit file list* (never a tree walk that includes the artifact); a close record's numbers must reproduce from the cited artifact with a stated command.
- **Cross-doc token semantics need a single hand + a consuming-side check.** The CQ-2 `mixed` leg had to land *identically* in 4 docs; the apply pass fixed the producers (m-2/m-3) but m-7's fixture/ledger stayed `self_reported`-only — the author's own verify missed it; the VP's independent cross-check caught it. **Fix:** single-hand any cross-domain token convergence, and verify every *consuming* doc, not just the producers.
- **Decline-don't-guess is a load-bearing mechanism.** Telling file-apply agents to DECLINE (not guess) anything needing a design decision turned 8 mis-classified "mechanical" findings into clean pair-dispatches instead of bad auto-edits. **Fix:** bake "decline + report" into any automated apply pass over judgment-bearing artifacts.
- **Snapshot before a batch apply on a no-VCS workspace.** The docs tree has no git; a `cp` snapshot before the fan-out made every misedit recoverable. **Fix:** always snapshot before an automated multi-file edit pass here.
- **Owed items get a ledger home, not a prose flag.** The ③/⑤ owed fixtures + m-5-F2 were "flagged owed" in pair docs but only became inherit-able once **enrolled in the §C4 build-carry ledger**. **Fix:** an owed/deferred item is closed only when it sits in a ledger a downstream builder reads — not when a pair doc says "owed."
- **Seam convergence needs a dedicated adversarial differential — presence-checks and sampling pass while the *formulas* diverge (c6.1).** The c6 close certified "all seams converged" on a presence-check verify sweep (atoms present in both docs) + VP sampling; a scoped post-c6 differential then found **5 seams that had diverged on the actual formula/token** (m-3 folded `chosen_bucket ≠ declared_bucket` where the lock says `declared_bucket ≠ rank-1`; m-3 mapped egress→terminal `held`; m-5 said "every send observes" vs Step-1's no-observe boundary; two m-2 mirrors lagged m-4/m-7) **plus 3 dropped cross-domain flags** buried in delivered+read relays. **Fix:** when parallel pairs edit *toward* a shared seam, run a cheap independent differential before the close — each seam = one adversarial lane trying to find a **byte-level disagreement**, not "is the atom present." It caught what pair-review + orchestrator-sweep + reviewer-sampling all missed. (The item-level owed-ledger in Part D is the standing form of the flag-census half of this.)

---

## Part E — External-input folds (adopted lessons from outside the team)

### E.1 "How Agents Quietly Break Architecture" (video; transcript at `agents-quietly-break-architecture-transcript.txt`, local reference copy, not vendored; folded 2026-07-07 mid-s6, operator-directed)

The failure class (agents preserve *local* contracts while deforming *global* semantics; prose rules don't compile
and go quietly stale) matched our own worst defects (F6/F7, the s6 activation-marker conflict, DEF-2). **The
substance lives where it's enforced — this entry is the pointer, per the very policy adopted:**
- **`PROTOCOL-DEVIATIONS.md` B9** — the three adopted rules: the named **"green-by-erasure"** review lens · **every
  locked law names its executable fixture** (a design-phase exit criterion) · **point-not-restate** as standing policy.
- **`ARCHITECTURE.md` §C4 INV-CATALOG** — the registered carry: the named-invariant `test/invariants` battery gate
  (Cardinal rule 1 as a compile-time tripwire; lands at s6-close, NOT s6 scope).
Process note for future folds: the first draft of this very entry restated B9 nearly verbatim — three copies of one
truth. Second copies diverge; external-input folds get ONE substance home + pointers.

## Part A.7 — STEP-1, the build phase, as it ran (s1–s6, 2026-07-03 → 2026-07-08) ✦ Step-1 CLOSED

Six slices, each a fresh nested orchestrator-team (B1), each finding real fragility the builders missed — **six for
six on the fresh-eyes bar**. The arc: spine (s1) → engine (s2) → forms (s3) → the wire-up + the first live
cross-vendor relay (s4) → consumer schemas, begun AS the first team governed ON frank (s5) → the dogfood found the
transport's floor in one day (F1–F17; the F11 livelock; operator stop-the-line; `TRANSPORT-FINDINGS-2026-07-06.md`)
→ the in-step ruling → the s6 design-amendment phase (three pairs, the operator-grilled parenting fork, two pair +
three VP must-revises — every one a real defect) → the s6 build → **the step-exit test passed live on the fixed
conductor, including replaying the exact traffic that broke s5** (14/14, zero parent-class, zero livelock). Close:
`main@6a1198a`, tag `s6-close`, ten verification chains.

**What Step-1 taught the process (the distillates already live elsewhere; pointers per Part E discipline):**
- Dogfooding the product AS the process found in one day what fixtures never would (B8; the ledger = the fix's spec).
- The gates caught real defects at EVERY altitude — pair grills, fidelity edges, VP co-sign, the operator's own
  pre-submit holds on gate day (the digest-canonicalization refusal; the bad-picks challenge that re-rooted the
  fork's empirics). Authority layering (grant → review → execute → verify) never slipped once in six merges.
- Executable truth beats restated prose (B9; the s6 co-sign conflicts all lived in paraphrase; INV-CATALOG is the
  standing fix, first in the post-close queue).
- Honest claim ceilings survived six closes unbroken: transport/provenance only; every "what held" credited; every
  residual named. The step closed with its two newest findings riding OUT as typed OIs — the fence at the finish line.

*Maintenance: append a new Part A-style section per cycle as it closes; keep Parts B/C/D as the living distilled
template — and Part E for adopted external lessons. This doc is the team's process memory — update it at every cycle
close.*

## Part F — The Step-3 operating model: the slice-orchestrator tier, on frank (operator-decided 2026-07-14)

Part B is the *design-cycle* template (AUDIT→DESIGN→…→lock). This is its build-side complement for **Step-3 and forward**: how a build slice runs once the team **runs ON frank as courier** and adds a third tier. It is the activation of the **deferred T4 tier** (`master/relays/master-org-decomp/…-031111`, 2026-06-28) — shelved then on the explicit trigger *"until v3 automates the relaying"*; frank is now that courier, so the trigger is pulled.

### F.1 The three tiers
- **T1 — master apex:** `master.orchestrator-planner` (CTO) + `master.orchestrator-reviewer` (VP). Hold the cross-cutting contracts, the architecture-of-record, and arbitration; answer up-escalations; hold the merge-decision (operator grants).
- **T3 — domain pairs (`m-1…m-7`) = the PMs.** Each authors the **spec-of-record** (the plan/design doc) for a slice, owns its domain's locked contracts, guides, and rules escalations in its domain. `m-x.planner` = the PM who owns the spec; `m-x.implementer` = the domain's adversarial reviewer.
- **T4 — the slice orchestrator team (NEW, per slice or bundle):** its own `sN.planner` + `sN.implementer` (+ reviewer as the work warrants). Owns **local detail-design + plan + impl** against the m-x spec-of-record. Acts independently; escalates up on the triggers below.
- **T2 (division orchestrators) stays skipped** — Step-3 is essentially one division (Harness Runtime); revisit when Steps 4–6 parallelize across divisions.

### F.2 The authority boundary — what's local, what escalates
- **LOCAL (decide + do, no ask):** everything that *fills in the details of the m-x spec* — implementation-level design, the build, local plan-review + end-review, task sequencing **within** the spec and the fence.
- **ESCALATE (stop + route up):**
  - **(a)** a **mistake** the team finds in the spec doc;
  - **(b)** a **better way** it discovers while filling details;
  - **+ the standing `DELEGATED_DISPATCH_AUTHORITY` triggers:** touching a locked contract · scope/boundary (fence) deviation · cross-slice/bundle collision · any design-of-record amendment.
- **The clean line (why this is tighter than it sounds):** the overall design is the m-x PM's, never the T4 team's — so **any departure from the spec (contradiction, improvement, or a locked-surface touch) is a trigger by construction.** The team never silently re-designs. This structurally prevents the s11 catch-#4 class (a fence row licensing a silent design change; see Part C / [[fence-row-vs-owner-path]]).

### F.3 The escalation ladder (the relay flow)
`slice team → master.orchestrator-planner (route + arbitrate) → the owning m-x.planner (rules: amend the spec · accept the improvement · decline-with-reason) → back down.`
- **Master is router + arbiter, not a bottleneck:** cross-domain arbitration + the architecture-of-record stay master's; domain-specific answers route to the owning m-x PM.
- **On frank:** escalations are relays through `submit`/`project`/`read` to the higher-tier seat's inbox — no operator hand-relay. frank stamps identity, validates the form, observes done-ness, parks/wakes — the governance the hand-relay used to do by hand.

### F.4 The per-slice lifecycle (Step-3 shape)
1. **master** decomposes → assigns a slice/bundle to a **T4 team**, names the guiding **m-x PM** (authors/owns the spec-of-record), and grants `DELEGATED_DISPATCH_AUTHORITY` with the F.2 trigger set.
2. **T4 team:** adopt the spec (execute, don't re-litigate) → local detail-design + plan → own plan-review → build (RED-first, fence-disciplined, the ten INV-CATALOG laws green per commit) → local end review → **escalate on any F.2 trigger**.
3. **m-x PM:** answers up-escalations; files **owner-fidelity from OUTSIDE the T4 pair** — never tokens its own reviewer (the [[slice-team-staffing]] rule; a mis-seated token lints clean, so staffing is a human check).
4. **master:** routes/arbitrates escalations; issues the **merge-decision → the operator grants** (recognized `HUMAN_MERGE_AUTHORIZATION` field at grant time); runs the slice close fold + the step-exit. Everything flows through frank.

### F.5 Proven vs. new vs. still-owed
- **Proven (Step-1 + most of Step-2):** m-x authors the spec-of-record → the slice team builds against it → escalates on the `DELEGATED_DISPATCH_AUTHORITY` triggers (s9/s10/s11 ran exactly this).
- **New for Step-3:** (a) the T4 team owns **local design** too (not just impl), so the escalate boundary now includes spec-mistake (a) + better-way (b); (b) **three tiers, not two**; (c) transport **on frank**, not operator hand-relay.
- **Still owed (the convention debt — VP's 2026-06-28 caution, the half frank doesn't solve):** **nested-run lineage across three tiers** · the T4 team's **authority-ceiling-at-spawn** · how master's arbitration reaches down. m-5's archetype layer (`orchestrator_lead` seat + authority-ceiling-at-spawn, the deferred conductor/N-pair template) is the eventual carrier. This Part is the operating spec; the lineage/ceiling mechanics get pinned as the first slice actually runs on frank (or as a bounded design cell). The **frank live relaunch + seat roster + shakedown** is the separate execution artifact that stands this up.

### F.6 Working directory = `frank/`, not the harness root (operator, 2026-07-14)
The Step-3 team **launches with its working directory in `frank/`**, not the cwd this charter loads from. Reason: the team both *builds* frank and *runs ON* frank — the store, the `frank-mcp` seat surface, and the build code all live in `frank/`. This reinterprets charter rule 2 (which assumed the team's cwd = the harness root); the rule-2 split (governance docs at the root, code in `frank/`) **still holds**, but the team now reaches the governance side *upward* instead of *locally*:
- **The charter must be reachable from `frank/`.** `frank/` is a subdirectory of the harness root, so `../CLAUDE.md`/`../AGENTS.md` is a parent of the launch cwd — ensure it loads on boot (a parent-directory walk, or a `frank/AGENTS.md` pointer as belt-and-suspenders). **No governance CLAUDE.md is committed into `frank/`** — rule 2's "don't put governance docs in `frank/`" stays intact; `frank/` stays the clean product repo.
- **Durable governance stays at the harness root** (`../master/`): the charter, `ARCHITECTURE.md`/`RECONCILE.md`/`ROADMAP.md`, and the **master design-of-record relay trail** (`../master/relays/`). The team reads/writes these by relative path; they do NOT migrate into `frank/`.
- **Live courier relays flow through frank's STORE** (the runtime, inside `frank/`) — that is the point of running on frank. The clean split: **frank's store = the live transport carrier · `../master/relays/` = the durable design-of-record trail.** How live-store relays export/archive into the durable trail at each slice close is pinned by the relaunch plan (the open "where does the trail live" question — the store is the live carrier; the durable design-of-record is not abandoned).
- **relay-lint + paths:** governance relays lint against `../master/relays`; slice-team working paths are relative to `frank/`. Absolute paths in cross-tier relays must resolve from either cwd.

### Part F addendum — reconciled to the ratified architecture reframe (2026-07-15)
The Step-3 operating model above (the T4 slice-orchestrator tier, on frank) stands, with its **target re-cut by the ratified reframe** (`master/STEP-3-ARCH-AMENDMENT.md`, SHA-256 `2d240eb6…`): the T4 tier now builds the **app shell around the conductor** (m-10 App Control Plane/Supervisor · m-8 connector · m-9 worker), **not** a conductor-hosted runtime. Consequences for the model:
- **The PMs shift with the domains.** m-10/m-8/m-9 (Division II) are the spec-of-record PMs for the app-shell slices; m-1..m-7 remain the conductor-domain PMs. The `T4 → master → m-x → master → T4` escalation round-trip is unchanged.
- **Build order is the reframe dependency graph:** the **coordinated first stage** (m-10 boundary design + the m-5 ceiling-host amendment, interface-locking the shared ceiling contract) precedes any m-8/m-9 consumer lock; then m-8/m-9; then the remaining amendments; then the MVP vertical (one governed turn, single pinned lane, live E3).
- **The T4-token gate is unchanged** (full frank roster + the round-trip + durable export) — the reframe changes *what* is built, not the tier mechanics or the design→lock→PLAN→T4 gates.
- The "first Step-3 acts = operating-model spec + frank live relaunch + shakedown" bridge is complete; the reframe + its source fold is the operative Step-3-DESIGN entry point.

## Part A.8 — Stage-6 re-lock, Lane 2 (the interface DAG) + the §D-settlement amendment (2026-07-21 → 2026-07-25, cycle IN PROGRESS at time of writing)

Where Part F is the *build-side* operating model, this section is a **design-cycle worked example** at the finest grain yet run: not a domain-pair audit→design→lock, but an **interface-DAG reconciliation** across six pairs (m-2/m-3/m-8/m-9/m-10 + m-7's own lane-1 study) converging one hashed re-lock, immediately followed by a **standalone corrective amendment** (§D-settlement) that iterated four full VP rounds on its own. Both threads run under the Stage-6 re-scope addendum (`master/STEP-3-STAGE6-AMENDMENT.md` rev12 `1125b0a0…`, operator-ratified 2026-07-21) and the §11 sequence it opened: lane 1 (m-7 broker study) → lane 2 (interface DAG legs) → item A + fixture freeze → the shorter re-lock → T4. (Item A originally named a machine-checkable extraction-bundle mechanism; that mechanism is now superseded/withdrawn per the item-A simplification amendment `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev7 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373` (operator-ratified 2026-07-27), replaced by the plain byte-bound interface-lock record `master/STEP-3-INTERFACE-LOCK.md` (external SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`); as of this writing, lane 2 is CLOSED and item A is RATIFIED + AUTHORED and in VP + F73 review, ahead of lane 4's Master+VP re-lock over the record + the `STEP-3-EXIT-FIXTURES.json` freeze.)

### A.8.1 — Lane 1 closes first (2026-07-21), unblocking Lane 2
The m-7 broker study (`step3-relock-m7-broker`) ran DESIGN → SITREP → dual F73 consumer-confirm (m-9 `…-215500`, m-10 `…-221000`) → a two-sided §D join → **NO-H-24** (the simpler broker rule set holds; no TLA+/Alloy spike needed) → a VP integration-confirm that caught a real ledger gap (`F73-M10-R40-LEDGER`: the frozen r40 IPC/seam contract still carried broker fields a stage-5 sweep would have silently frozen over) → a corrected four-item ledger → VP APPROVE (`step3-relock-broker-confirm-review-r2`, 20260721-225500). **Lane 1 complete in ~5.5 hours, same day** — the fastest full pair-confirm-join-VP loop run yet, because the scope was pre-narrowed (one study, two confirmations) rather than fanned across pairs from the start.

### A.8.2 — Lane 2: a producer-first interface DAG, not a flat fan-out
Immediately on Lane 1's close, master dispatched **six parallel DESIGN legs in one batch** (20260721-231500 through -231504, dispatch ids `step3-relock-dag-m9`/`-m10`/`-m3`/`-m8`/`-m2`, all parented to `step3-relock-broker-confirm`) — but explicitly **ordered by data dependency, not by pair number**: m-8 was named the **"producer-FIRST root"** (item B `frozen_core_digest` + item E `provider_lowered_tools_digest` — m-8 alone lowers, nobody reproduces its bytes), with m-9/m-10 as **carriage consumers** and m-3 as the **join/sink** binding both digests into its E0/E3 evidence schema, and m-2 as the **narrowest producer** (just the `logical_tool_schemas[]` contribution to m-9's `logical_surface_digest`). The **first attempt at this batch was revoked and re-cut same night** (VP `step3-relock-dag-review-r1`, 20260721-234500 → master `step3-relock-dag-hold`, 20260721-235500): the original six dispatches under-specified the writer-fence/source-id/content-ready-receipt obligations, misassigned the E-digest rows, and didn't stage the producer-first dependency explicitly. The re-cut fixed exactly those four defects and re-released the batch — **the DAG shape itself was a design-review finding**, not a given.

Once producer-first bytes existed, subsystem *internals* the DAG legs surfaced (m-9's crash-safe resume log framing, m-10's F59 ticket-cap mechanics) were **classified DELEGATED under F73** — governed additive deltas the owning pair designs and self-reviews against its own frozen base, rather than a master-authored spec each pair confirms. This is a distinct, cheaper category from a master-dispatched DESIGN: F73 says *the owning pair is close enough to the mechanism that master doesn't need to pre-specify it*, so the deliverable is "owner bytes + fresh pair review + F73 confirmation from the consuming legs," not "master routes a spec down."

**m-8's leg closed first and fastest** (`step3-relock-dag-m8`, SITREP 20260722-151500): pair-approved byte-bound at r5, convergence **3→3→1→1→0 findings across five rounds**, then explicitly reported "producer-first satisfied → consumer halves unblocked" — the DAG discipline paying off exactly as ordered: nothing downstream could bind before the root did.

**m-3's leg (the join/sink) took by far the longest and the most rounds** (`step3-relock-dag-m3`, 21 revisions from r1 through r20, spanning 2026-07-22 through 2026-07-25) — because it sat at the confluence of every producer's bytes AND had to derive a five-way classification (`{P, A, N3, N910}` over the tuple `{data_p_reply_kind, stream_terminal, ctrl_c_disposition}`) purely from what other pairs had already frozen, with no writer of its own. Notable finding-classes along the way: a totality claim withdrawn when the planner caught its own doc asserting a projection was "total, pure" in the same section that proved an input underivable (r6, "the FOURTH INSTANCE, found where I asked the reviewer to look" — the planner had flagged the review posture and then caught the class itself); three consecutive rounds (r16→r19) chasing a single circular-authority bug where a `not_emitted` classification kept trying to corroborate itself off `m10_row_state`, which the design's own review caught reintroduced through a stray parenthetical even after two prior "removed" claims (r18, r19 — the r19 fold explicitly names this "the r10/r11 spelling-miss family, now in my own check": a zero-grep-count proves absence of one *spelling*, not of the *authorizing pattern*). The leg closed **honestly partial** (`step3-relock-dag-m3` SITREP, 20260725-033000): T1–T8 (every "present DATA-P acquisition" cut) settled and bound to three producer hashes it reproduced itself (m-9 r12 `04422965…`, m-8 r7 `734e44b7…`, m-10 rev3 `cd17db32…`); T9/N910 (the two "loss cut" rows with no DATA-P reply at all) **explicitly refused a record** rather than manufacture one — routing two decisions up to master instead of self-resolving them.

### A.8.3 — The N910 loss-cut: confirm-or-object as a convergence instrument
The N910 question — *is there any existing surface that can authoritatively prove a per-attempt DATA-P reply was never emitted (as opposed to "we didn't observe it")?* — is the worked example of the **confirm-or-object pattern**: master ruled a **lean disposition** (documented MVP limit, not a build-and-fix) and routed it as two narrow yes/no questions to the two owners who could falsify it, rather than commissioning a full design (`step3-relock-dag-m3` RECONCILE, 20260725-093000). m-9 answered first (20260725-101500): all three of its loss surfaces (`stream_lost`, a race-conditional E0, silent worker crash) *describe an outcome*, none *proves bounded absence* — and doing so would require exactly the bounded-terminal/timeout authority the stage-4 GRILL_LOCK's sync-authorize/async-record posture deliberately withheld from m-9's frozen design. m-10 answered next (20260725-101500): it holds **no DATA-P endpoint at all** post-spawn (§G.3), so it isn't merely under-built here — it is structurally the wrong party to ask; but it confirmed the loss is disclosed via its own `UNKNOWN_PROVIDER_OUTCOME` terminal, which the D2 settlement manifest maps to `uncertain` (never `settled_with_content`). Both no-objects converged the same way: **no cheap carrier exists + disclosure is already honest** ⇒ documented MVP limit, ruled by master (20260725-110000), accepted by m-3 with **zero bytes changed** on its own already-approved r19 (20260725-113000). Master's own framing named the general shape: *"the R2 parallel is INVERTED — R2 was a positive discriminator on a reply that IS emitted; N910 needs proof of a bounded ABSENCE, which is a different and harder claim."* Confirm-or-object is cheap exactly when the lean is right and the question is narrow enough that a no-object costs one relay each; it would have been the wrong tool for a question where either owner might plausibly want to reopen the mechanism.

### A.8.4 — The §D-settlement amendment: four rounds, monotonically shrinking findings
Once the N910/mirror decisions landed, master authored a standalone correction (`master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md`) fixing four items the Lane-2 legs had surfaced but not yet folded into the operative claim set: (1) an honest relabel of the D-4 Gate-2 claim (comparator bytes unchanged, only the operator-visible framing), (2) the run-wide `MAX_PARKED_ROWS_PER_RUN=512` restore + a new `parked_unknown_capacity_exceeded` terminal, (3) the `relay.submit` `canonical_resource` cell bound by hash to m-2's `5ec7a3d2`, (4) an explicit `turn_failed` zero-attempt scope clarification. This amendment ran **four full authorship/VP rounds**, each one narrowing:
- **rev1 `1f822e47`** (20260724-151500) → VP r1 REVISE (20260724-160000): four blockers — the cap terminal lacked exact FAILED/closed-reason/no-successor/operator semantics, the frame proof omitted the boundary assertion, the frozen r21 text needed an *explicit* supersession rather than a reinterpretation, and the owner matrix dropped two corrections.
- **rev2 `7137b18a`** (20260724-163000) → VP r2 REVISE (20260725-120000): all four rev1 findings closed; three fresh ones — a compile-time frame constant was mislabeled as an *attainable* value rather than a saturating upper bound, a citation to a historical hash needed removing now that live successors existed, and one correction's routing table conflicted with pair-approved text elsewhere in the packet.
- **rev3 `ab10e6ef`** (20260725-130000) → VP r3 REVISE-ONE-PHRASE (20260725-140000): all three rev2 findings closed; one residual — the new "timeless-fold" rule's quantifier ("any owner produces a fresh successor") contradicted the packet's own byte-final `m-2: None` row.
- **rev4 `1fa71cb8`** (20260725-150000) → VP r4 **APPROVE, zero findings** (20260725-160000) → presented to operator for hash-bound ratification (20260725-163000).

Findings per round: **4 → 3 → 1 → 0** — strictly monotonic, the same shape master named as the standing "honest convergence signal" in prior cycles (c6.1's seam differential, m-3's own Lane-2 sink leg r10→r19 — ten review rounds per m-3's SITREP). The amendment sat at the **operator ratification gate as of 2026-07-25 16:30**, unresolved at time of writing; ratification licenses only the four corrections plus opening the propagation matrix — no DESIGN-lock, PLAN, T4 token, or downstream authority follows automatically.

### New friction (folds into Part C)
- **CC ≠ obligation.** An act is real only when addressed `TO` the actor; being CC'd is visibility, not authority or instruction. First surfaced hard when a **reviewer-CC relay tried to revoke authority master had delivered direct to six pairs** — VP's `DAG-R1-F1` finding held that a CC'd review cannot itself revoke a direct dispatch, forcing an explicit `HOLD/STOP` re-cut addressed `TO` the six pairs (`step3-relock-dag-hold`, 20260721-235500) before the revocation was real. m-3 caught the identical class in itself two days later and self-corrected without being told (`step3-relock-dag-m3` RECONCILE, 20260723-155000: "same class as master's DAG-R1-F1/R2-F1 — act must be TO the actor, intent != addressed relay"). Per master's running tally this is now a **7-instance/4-seat recurring class** (confirmed by master 2026-07-25); it is now flagged as a **relay-lint mechanization candidate** — a lint rule that flags any relay narrating an obligation, revocation, or instruction toward an address that appears only in `CC`, never `TO`.
- **Reference owner state by rule, not snapshot — the timeless-fold rule.** The §D-settlement rev2→rev3 VP finding (`SETTLE-VP-R2-F2`, 20260725-120000) caught the amendment binding a *current-working-delta* snapshot hash (`48062d18…`) as if it were a durable fact, when the live artifact had already moved past it. The fix generalizes: **never freeze a mutable working-state pointer into a ratification's bytes** — cite either a frozen/pair-approved hash (timeless, safe to bind) or an explicit rule for deriving the current one at fold time (e.g. "each changed owner produces a fresh successor over the then-current artifact"), never a point-in-time snapshot of something still moving. The rule itself then needed a second round (`SETTLE-VP-R3-F1`, 20260725-140000) to fix its own quantifier — a universal "any owner" contradicted the one owner (m-2) whose row was already byte-final and had nothing to fold — **a generalization needs its own exemption list checked against every row it will be applied to before it ships.**
- **Producer-approved ≠ carriage-settled.** m-8's DAG leg closing pair-approved did not mean the digests it produces were live anywhere downstream — m-3's leg spent multiple rounds re-verifying it was still binding m-8's *reproduced* bytes rather than a stale citation, and the r18 fold explicitly re-stated the sequencing discipline as "producer → carriage → consumer," each stage independently checkable, none inferred from the stage before it having merely *closed*.
- **Confirm-or-object is a fast, narrow convergence instrument — not a substitute for design.** It worked for N910 because master's lean was falsifiable in one sentence per owner and neither owner had standing to want a different mechanism. The tell that it was the right tool: both replies were **structural "wrong vantage" answers** (m-9: "I'd have to synthesize a proof I don't have"; m-10: "I don't hold the endpoint at all") rather than "I could build this but chose not to" — the latter would have been a live design decision, not a confirm-or-object.
- **Owner-authored, ratifiable closed sets get bound by hash, not authored by the binder.** Recorded generally in the m-8 lane-2 dispatch framing (m-8 alone lowers its tool schema; m-9 does not reproduce it) and in m-2's §5-E logical-component design ("no owner hashes bytes it cannot see") — master/consuming pairs bind to a producer's hash; they never re-derive or re-assert the producer's bytes themselves, even when convenient.
- **VP findings converging monotonically to zero is the honest signal a document is actually settling**, not stalling — the §D-settlement amendment's 4→3→1→0 arc and m-3's Lane-2 leg's own round-over-round narrowing (each catching a *residue* of the prior fix, never a fresh unrelated class in the closing rounds) are both instances of the same pattern already named at c6.1: an amendment or design in real trouble re-opens *new* classes of finding late; one converging to a strictly shrinking residual is closing for real.

*(Continues Part F's "Working directory = `frank/`" convention: this section's relay trail lives under `master/relays/step3-relock-dag-*` and `master/relays/step3-relock-settlement-amend/`, addressed and read from the harness root, per the standard master/frank split.)*
