# DESIGN-REREVIEW — adversarial re-review of the post-re-baseline frank design-of-record

**Status: REVIEW-OF-RECORD (2026-07-02). Verdict: CONDITIONAL-GO — clear the c6 punch-list (doc-only), then open Step-1 PLAN.**
**Supersedes-in-force:** the 2026-07-01 NO-GO is DISCHARGED at the structural level (no FATAL survives); its residue is tracked as c6 line-items below.
**Awaiting:** VP concurrence on this verdict + the §4 decomposition/routing before c6 fix-dispatch.

---

## 0. BLUF

- **No FATAL. The NO-GO is not reopened.** 90 findings confirmed, 1 refuted, 0 unverified across 10 coarse lanes. Severity: 0 FATAL / 28 MAJOR / 52 MINOR / 10 OBSERVATION.
- **The re-baseline's load-bearing achievements held under scrutiny:** the m-7 conductor-core substrate is sound; its one by-construction claim (the single-threaded serialized-loop two-honest-seats double-accept kill) survived; the attach+interface-guardrail *confusion-resistant* framing held; the D5 residual boundary held; no lane found a structural collapse.
- **But the design-of-record is not yet internally consistent enough to hand a Step-1 PLAN author.** Every MAJOR is bounded and doc/fold-level — no mechanism is wrong — and they nearly all trace to one root cause (§2).
- **Disposition:** one bounded, DOC-ONLY cleanup cycle (**c6**), routed to the owning m-1..m-7 pairs (domain-local) + CTO/VP (architecture-of-record, cross-domain seams, governance surfaces), VP-gated at both ends. No mechanism change, no re-design, no re-lock of sound decisions.

## 1. Method

- **10 coarse lanes** (operator-chosen granularity, coarser than the 07-01 pass's 16 micro-lenses): 7 per-domain (m-1..m-7) + 3 cross-cutting (x1 original-FATAL-resolution audit, x2 two-sided claim-honesty sweep, x3 seam + byte-consistency integrity).
- **Fable 5 on max effort** per reviewer; **single adversarial verifier** per finding (told to REFUTE; re-litigating a locked decision or re-raising the accepted D5 residual ⇒ auto-REFUTED). Pipeline review→verify.
- Every lane was briefed on the **locked boundary** (attach+guardrail = confusion-resistant; D5 accepted; the four sanctioned by-construction claims) so it re-reviewed the CURRENT docs rather than re-running the 07-01 NO-GO.
- **Honesty note on the run:** a session limit killed 25 verifiers (the three x-lanes + one m-7) mid-run; they were resumed from cache once the limit cleared (unchanged agents replayed free; only the 25 re-ran). Final state: 101/101 agents, 0 errors, 0 unverified. The 1 refutation (x3-F10, verifier caught a misread of NF-S14's exception) is the signal the verify layer was not rubber-stamping.

## 2. Root cause — scope-boundary leakage

Almost every MAJOR traces to one fact: **the c4 locks and the c5 claim-sweep were scoped to the seven design-of-record docs (m-1..m-6 + ARCHITECTURE).** The re-review found the leakage at exactly that scope boundary:

- **Retired vocabulary survives in the docs a builder actually boots on** — `CLAUDE.md` (loaded every session), the m-1/m-3 domain READMEs (domain front doors), the live dashboard's "Decisions (c1 — LOCKED)" section, and the `RECONCILE` ledger still assert *sole-writer / forgery-robust by construction / tamper-resistant / unforgeable*. The c5 close certified the sweep "COMPLETE / the docs now tell the honest story end-to-end" — a certification the very dashboard making it falsifies.
- **Decisions folded into one doc but not their enforcing twin** — decision ③ landed in m-6 + §J but cites an m-7 fixture that does not exist and has no home in the m-2 schema; the §2C carry-ledger was narrowed during the c5 close (a CTO regression) and dropped the routing-lane carries from every current (d) status list.
- **Genuine gaps the coarse lanes surfaced in the newly-locked substrate + the seams** — m-7 multi-record linearization, the m-3↔m-4 deviated_observed formula fork, R2's ghost field name, the CQ-2 `mixed` edge certified closed while admittedly open, and the one original FATAL (conftest.py / observe-runs-lane-code) that landed in no ledger.

## 3. The findings, clustered (the 28 MAJORs into 8 themes)

1. **Surviving overclaims outside the swept 7 docs** — CLAUDE.md, m-1/m-3 READMEs, dashboard, RECONCILE, + one in-scope miss (m-3 §2 "unforgeable"). *Fix: extend the sweep to these surfaces + put them under the standing byte-consistency guard.* (m-1-F1/F2, m-2-F3, m-3-F3, x1-F2, x2-F2/F3/F4, + the MINOR/CARRY tail)
2. **DI-1 never dispositioned** — the sweep relabeled I1's claim but left m-1's requirement/realization/test layer demanding the shelved wrap ("no lane write access / separate uid / E2: a direct lane write fails"). *Fix: disposition DI-1 = D3-shelved; Step-1 I1 rides the m-7 guardrail; re-cut AC#1 to the D4 fixture.* (m-1-F3)
3. **The conftest.py FATAL (original #12) survives homeless** — observe runs repo-authored code inside the TCB; no owner, no ledger, no relabel. Not D5. *Fix: record a named non-locking carry (m-3 §13, m-7-hosted) gating the Step-2 observe build + an m-7 executor-isolation fixture.* (m-3-F1, x1-F5)
4. **Decision ③ folded asymmetrically** — cites a nonexistent m-7 NF fixture; m-2 still types the pick as a free full-set enum. *Fix: register the m-7 fixture as an owed artifact; fold the RAISE-ONLY / known-A typing into m-2; reword m-6:44 until it lands.* (m-6-F1, m-2-F1)
5. **§2C carry-ledger narrowed — CTO c5 regression** — routing-lane carries (altitude-B per-row, R2 per-column) dropped from every (d) status list. *Fix: restore to all three ledgers + deferral markers in m-2/m-4.* (m-4-F1, x1-F1)
6. **CQ-2 certified closed over an open edge** — decision-② keys on `== self_reported`, so a `mixed` authority record delivers, against m-3's own §6 pessimistic rule. *Fix: broaden the key to `∈ {self_reported, mixed} ⇒ held` (or record a deliberate mixed-delivers decision) + fixtures.* (m-2-F2, m-3-F2, x3-F3, x1-F3)
7. **The claim-registry contradicts itself** — ARCHITECTURE §C4.3 licenses ONE by-construction claim; the ratified set is FOUR classes. *Fix: reword §C4.3 to enumerate the four exactly as RECONCILE c5 records them.* (x2-F1, m-1-F8, m-7-F8)
8. **m-7 substrate gaps in the just-locked core** — multi-record outcomes have no linearization point (crash-atomicity covers only single-record commits); a 3-way disposition mismatch (NF-S6 `held` vs §6 `rejected` vs m-3 `accepted`) is not jointly satisfiable. *Fix: designate one pivot per mutation; qualify NF-S6 to authority-class + run an m-3/m-7 COORD.* (m-7-F1, m-7-F3)

**Adjacent seams (mostly build-step carries):** deviated_observed formula fork (m-4-F2), R2 ghost field name `chosen_model` (m-4-F3), template_ref "null in Step-1" unsatisfiable (m-5-F3), archetype/authority_ceiling confused-deputy (m-5-F1), trusted-config author-set omits m-5 (m-7-F2/x3-F2), routing-record authoring model (m-4-F4), deviation_reason_code value set (m-4-F5).

**The MINOR/OBSERVATION tail (62):** predominantly one mechanical cluster — stale line-anchors and stale pre-lock STATUS headers the c4/c5 edits shifted — plus the low-severity tail of the scope-hole. One hygiene sweep clears them.

## 4. Disposition & routing (what the VP signs off)

c6 routes domain-local findings to the owning pairs and architecture-of-record / cross-domain / governance-surface findings to CTO/VP. Tags: **BLOCKER** = clear before Step-1 PLAN opens; **HYGIENE** = mechanical sweep; **m** = minor/carry (fix in-cycle or record).

| Owner | Findings | Load (B/H/m) | Blocker themes |
|---|---|---|---|
| **m-1 pair** | 12 | 4B / 2H / 6m | README+CLAUDE overclaim; DI-1 disposition |
| **m-2 pair** | 5 | 0B / 0H / 5m | (seam-driven; see CTO) |
| **m-3 pair** | 14 | 5B / 1H / 8m | conftest.py carry; ②-mixed; unforgeable |
| **m-4 pair** | 9 | 4B / 2H / 3m | §2C carries; deviation_reason_code; R2 names |
| **m-5 pair** | 6 | 0B / 2H / 4m | (seam-driven; see CTO) |
| **m-6 pair** | 6 | 1B / 0H / 5m | `held` definition |
| **m-7 pair** | 8 | 2B / 1H / 5m | linearization pivot; config author-set |
| **CTO/VP · cross-domain seams** | 13 | 9B / 1H / 3m | ③, deviated_observed, template_ref, archetype, R2, disposition, config |
| **CTO/VP · governance surfaces** | 17 | 4B / 4H / 9m | §C4.3 one-vs-four; dashboard/RECONCILE/CLAUDE overclaim |

## 5. Proposed canonical resolutions for the coordinated (cross-file) clusters

CTO-proposed, for VP ratification. Each is doc-consistent (follows the finding's own recommendation or the doc's own stated rule). **◆ = design-substantive** (changes a disposition/semantics — needs explicit VP ratification); **○ = mechanical** (relabel / restore / align). Each carries the canonical text so the owning pairs apply consistent per-file slices.

- ◆ **CQ-2 `mixed` edge (theme 6):** broaden the fail-closed key for authority-class records to `record_integrity ∈ {self_reported, mixed} ⇒ held`, per m-3 §6's own pessimistic-reduction rule; extend NF-S7 + m-2 AC16 with an `authority ∧ mixed` fixture. (Alt: record a deliberate mixed-delivers decision — not recommended; contradicts §6.)
- ◆ **m-7 multi-record linearization (theme 8):** one pivot per mutation — embed the candidate in the held-disposition record; merge burn+verdict into one operator-verdict record whose presence implies the decision-scoped burn; add a crash-between-canonical-renames fixture.
- ◆ **m-7 disposition boundary (theme 8) — VP-amended 2026-07-02 (two-axis; canonical wording, verbatim):** *Qualify NF-S6 by authority class AND by whether the trusted check could run. **No-vantage-at-start / m-3 unobservable:** non-authority records deliver as `accepted` with the self-reported/mixed label; authority-class records with `record_integrity ∈ {self_reported, mixed}` are `held` and escalated. **Trusted-side machinery-ran-and-broke / check-could-not-run:** authority-bearing records are `held`; non-authority records are `rejected` or author-returned with the fault edge named. m-3/m-7 COORD owns the wording split; m-2 consumes the `mixed` fail-closed key; m-6 keeps `held` on bucket A and `rejected` on bucket D.* — Rationale: m-7 §6 locks authority-bearing internal-fault → `held` (distinct from `rejected`); the earlier "internal-fault → rejected" shorthand erased the authority split and would break m-6's surface. Fix = split NF-S6 by authority class, not invert authority faults to `rejected`.
- ◆ **Decision ③ registration (theme 4):** register the m-7 known-A/RAISE-ONLY NF fixture as an owed artifact (ledger entry parallel to §2C, gated to Step-1 build); fold the RAISE-ONLY typing into m-2 (monotonic-toward-A hybrid: system floor = known-A membership, pick constrained within [floor, A]); reword m-6 §2:44 to "flagged to m-7, fixture registration owed" until it lands.
- ○ **§2C ledger restore (theme 5 — CTO regression):** restore R2 (gate_referenceable per-column + negative fixtures over `chosen_model` & single-family bucket proxies) and altitude-B per-row grain to all three (d) ledgers (RECONCILE c5 line, README both lines, ARCHITECTURE §C4 — retitle from "away-bridge" to the full §2C set) + non-locking deferral markers in m-2 §17.3 / m-4.
- ○ **deviated_observed fork (seam):** fold the GL-1 canonical (bucket-vs-bucket on `declared_bucket` + the auxiliary `bucket_binding_observed` atom) into m-3 §9; delete / precisely-condition the "equivalent fallback" bracket in m-4 §9.
- ○ **DI-1 disposition (theme 2):** DI-1 = D3-shelved wrap invariant; Step-1 I1 rides the m-7 interface guardrail (store path absent from every seat tool surface); re-cut m-1 §12 AC#1 to the D4 guardrail fixture + an explicitly D3-scoped wrap-spike criterion; fix ARCHITECTURE:39-40.
- ○ **claim-registry §C4.3 (theme 7):** reword to enumerate the four licensed by-construction classes exactly as RECONCILE c5 records them (serialized-loop kill; R2 gate-grammar; observer-selected control; authority-ceilings-at-spawn).
- ○ **conftest.py sandbox carry (theme 3):** named non-locking carry in m-3 §13 (owner m-3 + m-7-hosted execution) gating the Step-2 observe build — suite-class registry entries run in an unprivileged executor with no store/config/outbox handle and no signing key — + an m-7 fixture (NF-S6-analogue) proving the executor's privilege separation; relabel m-3 §4 honestly.
- ○ **trusted-config author-set (seam):** add m-5's archetype-registry section + m-2's declared section to the artifact author set (m-7 §7/S15, ARCHITECTURE C4.1); obtain the missing m-5 CQ-4b confirm (bounded COORD).
- ○ **archetype/ceiling + template_ref (seam):** split propose-vs-stamp (planner proposes role/archetype from a fill-time-pruned candidate set; conductor stamps the resolved `seat_archetype` + `authority_ceiling` per-column); fix m-2 mirror "null in Step-1" → "set when spawned from a template; null otherwise."

## 6. What held (why c6 is cleanup, not re-baseline)

- **The conductor-core substrate** (m-7, locked 2026-07-02): the serialized commit loop, crash-atomic single-record commit + recovery, internal-fault disposition, interface-guardrail enforcement — no lane broke the core. The two-honest-seats double-accept kill remains a genuine control-flow invariant.
- **The deployment fork** (attach + guardrail = confusion-resistant; D5 accepted): held. No lane re-derived the wrap-inversion that sank the 07-01 design; the honest claim-set is the right claim-set.
- **The sanctioned by-construction claims** (serialized-loop kill; R2 gate-grammar; observer-selected control; authority-ceilings): survived the two-sided honesty audit. The one defect (§C4.3) is that the registry mis-states their *count*, not that any is unearned.
- **The DNA** — observe-as-send, owner-typed FieldSpecs, honest-fallback labeling — intact.

## 7. The c6 gate plan

1. **VP sign-off (this relay)** — concurrence on the verdict + §4 decomposition + §5 proposed resolutions.
2. **Dispatch** — per-domain fix relays to m-1..m-7 pairs; CTO takes the seam + gov surfaces; coordinated clusters carry the §5 canonical resolution so cross-file edits agree.
3. **Verify** — re-run the claim-language grep net, byte-consistency token check, stale-anchor + cross-file agreement checks; residuals folded.
4. **c6 close, VP co-sign** — then, and only then, (e) Step-1 PLAN opens.

Standing rule reinforced by this review: `CLAUDE.md`, the domain READMEs, `master/README.md`, and `RECONCILE.md` join the design docs under the standing claim-sweep + byte-consistency guard. The sweep's scope was the bug.

## Appendix — full confirmed-finding inventory (90)

| id | sev | category | owner | tag | location |
|---|---|---|---|---|---|
| m-1-F1 | MAJ | surviving-overclaim | m-1 | B | master/domains/m-1-trust-identity/README.md:7-11 (§Owns) |
| m-1-F2 | MAJ | surviving-overclaim | m-1 | B | CLAUDE.md:26 (product paragraph) and CLAUDE.md:44 (m-1 org-chart r… |
| m-1-F3 | MAJ | step1-readiness-blocker | m-1 | B | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-2-F1 | MAJ | cross-doc-contradiction | CTO:seam | B | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-2-F2 | MAJ | gap-or-untestable | CTO:seam | B | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-2-F3 | MAJ | surviving-overclaim | CTO:gov | B | master/README.md :142 ('Decisions (c1 — LOCKED)' section) + master… |
| m-3-F1 | MAJ | original-fatal-not-resolved | m-3 | B | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F2 | MAJ | gap-or-untestable | m-3 | B | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F3 | MAJ | surviving-overclaim | m-3 | B | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-4-F1 | MAJ | cross-doc-contradiction | m-4 | B | master/RECONCILE.md:337 (c5 close, step-(d) description) |
| m-4-F2 | MAJ | cross-doc-contradiction | CTO:seam | B | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-4-F3 | MAJ | gap-or-untestable | CTO:seam | B | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-4-F4 | MAJ | gap-or-untestable | m-4 | B | m-4 design §7 (:245-263) + §0 GL-3/GL-4 |
| m-4-F5 | MAJ | gap-or-untestable | m-4 | B | m-4 design §5 :206 (+ §6 Stage-2 :222) |
| m-5-F1 | MAJ | cross-doc-contradiction | CTO:seam | B | master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-pol… |
| m-5-F3 | MAJ | cross-doc-contradiction | CTO:seam | B | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-6-F1 | MAJ | step1-readiness-blocker | CTO:seam | B | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-6-F2 | MAJ | cross-doc-contradiction | m-6 | B | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-7-F1 | MAJ | gap-or-untestable | m-7 | B | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| m-7-F2 | MAJ | cross-doc-contradiction | CTO:seam | B | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| m-7-F3 | MAJ | cross-doc-contradiction | CTO:seam | B | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| x1-fatal-resolution-F1 | MAJ | cross-doc-contradiction | m-4 | B | master/RECONCILE.md:337 (c5 close, step (d)) |
| x1-fatal-resolution-F2 | MAJ | surviving-overclaim | CTO:gov | B | master/README.md:139 and :142 (section 'Decisions (c1 — LOCKED)', … |
| x2-claim-honesty-F1 | MAJ | cross-doc-contradiction | CTO:gov | B | master/ARCHITECTURE.md §C4.3 :449-450 (vs its own §1 note :18-21 a… |
| x2-claim-honesty-F2 | MAJ | surviving-overclaim | m-1 | B | master/domains/m-1-trust-identity/README.md :7-11 (Owns section) |
| x2-claim-honesty-F3 | MAJ | surviving-overclaim | CTO:gov | B | master/README.md :139 (Decisions (c1 — LOCKED)) |
| x3-seam-byte-integrity-F2 | MAJ | cross-doc-contradiction | m-7 | B | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| x3-seam-byte-integrity-F3 | MAJ | gap-or-untestable | m-3 | B | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-1-F10 | MIN | byte-inconsistency | m-1 | H | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-1-F4 | MIN | gap-or-untestable | m-1 | M | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-1-F5 | MIN | new-regression | m-1 | M | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-1-F6 | MIN | cross-doc-contradiction | m-1 | M | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-1-F7 | MIN | byte-inconsistency | m-1 | H | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-1-F8 | MIN | cross-doc-contradiction | CTO:gov | M | master/ARCHITECTURE.md:449-450 (§C4.3) |
| m-1-F9 | MIN | surviving-overclaim | m-1 | M | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-2-F10 | MIN | gap-or-untestable | m-2 | M | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-2-F4 | MIN | byte-inconsistency | CTO:seam | H | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-2-F5 | MIN | gap-or-untestable | m-2 | M | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-2-F6 | MIN | new-regression | CTO:seam | M | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-2-F7 | MIN | cross-doc-contradiction | m-2 | M | master/domains/m-2-forms-determinism/README.md rows :24, :31, :34 |
| m-2-F8 | MIN | surviving-overclaim | m-2 | M | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-3-F10 | MIN | new-regression | m-3 | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F4 | MIN | cross-doc-contradiction | m-3 | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F5 | MIN | gap-or-untestable | m-3 | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F6 | MIN | gap-or-untestable | m-3 | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F7 | MIN | gap-or-untestable | m-3 | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-3-F8 | MIN | surviving-overclaim | m-3 | M | master/domains/m-3-observation-evidence/README.md:35 |
| m-3-F9 | MIN | byte-inconsistency | m-3 | H | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| m-4-F6 | MIN | byte-inconsistency | m-4 | H | m-4 design doc :13-14 (status header), :138-141 (§2 box), :318 (§9… |
| m-4-F7 | MIN | byte-inconsistency | m-4 | H | m-4 design doc :338-350 (§10), :435-439 (§15 M4-1) |
| m-4-F8 | MIN | surviving-overclaim | m-4 | M | m-4 design doc :127-129 (§2 snapshot-provenance) |
| m-4-F9 | MIN | cross-doc-contradiction | m-4 | M | master/ARCHITECTURE.md:99-102 (§J1 forward note) |
| m-5-F2 | MIN | cross-doc-contradiction | CTO:seam | M | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| m-5-F4 | MIN | byte-inconsistency | m-5 | H | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| m-5-F5 | MIN | byte-inconsistency | m-5 | H | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| m-5-F6 | MIN | cross-doc-contradiction | m-5 | M | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| m-6-F3 | MIN | cross-doc-contradiction | m-6 | M | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-6-F4 | MIN | surviving-overclaim | CTO:gov | M | master/ARCHITECTURE.md:346 (§C3.4, the Seam-C away-token row) |
| m-6-F5 | MIN | cross-doc-contradiction | CTO:gov | M | master/ARCHITECTURE.md:99-102 (§J1 forward bullet) — grep confirms… |
| m-6-F6 | MIN | gap-or-untestable | m-6 | M | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-6-F7 | MIN | cross-doc-contradiction | m-6 | M | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-6-F8 | MIN | new-regression | m-6 | M | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-7-F4 | MIN | byte-inconsistency | m-7 | H | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| m-7-F5 | MIN | gap-or-untestable | m-7 | M | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| m-7-F6 | MIN | gap-or-untestable | m-7 | M | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| m-7-F7 | MIN | cross-doc-contradiction | m-7 | M | master/domains/m-7-conductor-core/README.md L3 (and L47-49 "under … |
| m-7-F8 | MIN | cross-doc-contradiction | CTO:gov | M | master/ARCHITECTURE.md C4.3 (L449) |
| x1-fatal-resolution-F3 | MIN | gap-or-untestable | CTO:seam | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| x1-fatal-resolution-F4 | MIN | byte-inconsistency | CTO:gov | H | master/READINESS-REGISTER.md:14 (VERDICT header), :251-252 ('What … |
| x1-fatal-resolution-F5 | MIN | original-fatal-not-resolved | m-3 | B | master/DESIGN-REVIEW-2026-07-01.md:197-198 (§3.12) and :261 (§7 sa… |
| x2-claim-honesty-F4 | MIN | unearned-by-construction | m-3 | M | master/domains/m-3-observation-evidence/design/2026-06-29-v3-obser… |
| x2-claim-honesty-F5 | MIN | surviving-overclaim | m-5 | M | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| x2-claim-honesty-F7 | MIN | surviving-overclaim | CTO:gov | M | master/RECONCILE.md :24 (also :21, :44-45, :59-60, :85, :113) |
| x3-seam-byte-integrity-F1 | MIN | cross-doc-contradiction | m-4 | M | master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-pol… |
| x3-seam-byte-integrity-F4 | MIN | byte-inconsistency | CTO:gov | H | master/domains/m-6-human-surface-scheduler/design/…:113 |
| x3-seam-byte-integrity-F5 | MIN | byte-inconsistency | CTO:gov | H | m-1 design :22 and :112 ('ARCHITECTURE.md:58-65') |
| x3-seam-byte-integrity-F6 | MIN | surviving-overclaim | m-5 | M | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| x3-seam-byte-integrity-F7 | MIN | surviving-overclaim | m-1 | M | master/README.md:76 |
| x3-seam-byte-integrity-F8 | MIN | cross-doc-contradiction | CTO:gov | M | m-3 design :7 and :9 |
| x3-seam-byte-integrity-F9 | MIN | cross-doc-contradiction | CTO:gov | M | m-6 design :46 (§2) and :84 (§4) |
| m-1-F11 | OBS | other | CTO:gov | C | master/RECONCILE.md:21-24, :85, :113 (c1 sections) |
| m-1-F12 | OBS | cross-doc-contradiction | m-1 | C | master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-ident… |
| m-2-F9 | OBS | cross-doc-contradiction | m-2 | C | master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-sch… |
| m-4-F10 | OBS | byte-inconsistency | CTO:gov | H | master/README.md:17-18 |
| m-5-F7 | OBS | gap-or-untestable | m-5 | C | master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-arche… |
| m-6-F9 | OBS | other | m-6 | C | master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-hu… |
| m-7-F9 | OBS | gap-or-untestable | m-7 | C | master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-c… |
| x2-claim-honesty-F6 | OBS | surviving-overclaim | m-3 | C | master/domains/m-3-observation-evidence/README.md :35 |
| x2-claim-honesty-F8 | OBS | cross-doc-contradiction | m-7 | C | master/domains/m-7-conductor-core/README.md :3 (also :48-49) |
| x3-seam-byte-integrity-F11 | OBS | gap-or-untestable | CTO:gov | C | master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-015800.md:2… |

**Refuted (1):** x3-seam-byte-integrity-F10 — 'NF-S14 unconditional vs m-6 bucket-D exception' — verifier confirmed NF-S14 carries the exception; misread. WITHDRAWN.
