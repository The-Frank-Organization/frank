# Sprint s1 — reconciliation ledger

Append-only. Each entry: date, what was reconciled, agreement/disagreement/coverage-delta, evidence level, disposition.

## 2026-07-03 — boot ACKs reconciled (all seats online)

Reconciled the three boot-ACK SITREPs against disk (relay files read, INDEX rows present, `git status --short` clean):

- `s1.orchestrator-reviewer` — ACK `boot/s1-boot-orchestrator-reviewer/SITREP-orchestrator-reviewer-20260703-134923.md`; skill loaded from `~/.codex/skills` (distinct lane, as intended); standing by for CC'd broad-SET relays. E1/E2 claims verified.
- `s1-core.planner` — ACK `boot/s1-boot-s1-core-planner/SITREP-planner-20260703-134904.md`; agent-pair-planner + protocol v2.8.8 loaded; sees the s1-core-audit dispatch but holds under a direct operator instruction ("only boot"). E1/E2 claims verified.
- `s1-core.implementer` — ACK `boot/s1-boot-s1-core-implementer/SITREP-implementer-20260703-134911.md`; agent-pair-implementer loaded (`~/.codex` lane); same operator hold noted. Used a fresh DISPATCH_ID (`…-ack`, parented to the boot id) vs the others reusing the boot id — both shapes acceptable; no action.

Agreement: full — all three assumed the intended identities, claimed no work authority, and left the tree clean. Coverage delta: none. Disposition: team online; the only pending item is the **operator's release of the audit hold** (orchestrator-side, the s1-core-audit AUDIT dispatch is live and unmodified).

## 2026-07-03 — paired s1-core audits reconciled (s1-core-audit)

Inputs: `AUDIT-planner-20260703-140226.md` + `AUDIT-implementer-20260703-135833.md` (both lint-clean, both `FINAL_GIT_STATUS_SHORT: none — clean tree`). Reviewer RECONCILE on the dispatch itself: approve (`RECONCILE-orchestrator-reviewer-20260703-135325.md`).

**Agreement (full, no contradictions found):**
- `PRIMARY_BUCKET: still-open` — greenfield confirmed independently (E2 both sides); recommended-next = DESIGN.
- Spec-to-exit-gate map complete on both sides; **no spec gaps** — every exit-gate line maps to locked spec (file:line cited both sides, consistent).
- Duplicate gate: nothing to promote in `frank/`; v2.8.8 relay-lint = DO-NOT-COPY (m-2 §10 dissolves it); v2.8.8 store layout = REUSE-AS-SPEC'D; fixture corpus = read-only replay input; jcode/claude-code = prior-art-only.
- Frozen m-1/m-2 surfaces enumerated consistently (same lines cited: m-1 :124-131/:135-145, m-2 :21-97/:118-126/:278-283).
- Replay corpus located at the same path; implementer added E2 execution proof (`check-relay-lint-fixtures.py` all-PASS; changelog 146/146). S1-minimal subset selection = DESIGN/PLAN task.
- Q-A scope boundary (S1-minimal recovery vs S2 full recovery): both read it the same way — S1 needs intake journal + §4 pivot + dumb-replay recovery sufficient for the crash matrix; genesis/quarantine/GC/segment-rotation stay S2-OUT. **Needs guide confirmation** → relayed to m-7.planner as `s1-guide-q1`.

**Coverage deltas (complementary, not conflicting):**
- Implementer typed 7 owed-carry records vs the 3 the charter assigns S1 (charter :72 names exactly three). Disposition: the three chartered carries (code-layer guardrail, I-PH fixture, ③ known-A guardrail-adjacent portion) are S1 obligations; the extra four (⑤ ODB model-name egress, R2 per-column, `GRILL_REQUIRED` row, `routing_escalation`) are **recorded-as-context, not S1 obligations** — S1 must not contradict them, and the ODB-egress classification question is folded into the guide relay as a confirm.
- Implementer ran the fixture checker (E2); planner counted/located only — implementer's proof adopted.

**Planner's reconcile questions → dispositions (orchestrator):**
1. Q-A boundary agreement — implementer's scope-OUT re-check concurs; guide confirm pending (`s1-guide-q1`).
2. Additional organic corpus (master-trail failed relays) — the v2.8.8 fixture matrix is the S1 corpus of record; organic master-trail failures MAY enrich DESIGN thinking but are NOT R1 gate inputs (keeps the S1-minimal replay bounded; full dissolve validation is S3).
3. m-2 §9 migrator-registry — concur out of S1: `schema_version` stamping only; no migrator machinery.

**Reviewer watch item adopted:** downstream relays cite the current `main` head they actually inspected (now `461fae0`+), not a stale BASE.

Disposition: **proceed to DESIGN** — dispatch `s1-core-design` issued to `s1-core.planner` (implementer in CC as design-challenger), carrying Q-B..Q-E, the fixture-id namespace from the planner's map (B1-B4/A1-A4/C1-C6/R1/P1/L1/W1), and Q-A/ODB-egress marked provisional-pending-guide. Evidence: E1/E2 throughout as cited above.

## 2026-07-03 — guide answers reconciled (s1-guide-q1): both readings CONFIRMED; design de-provisioned

Input: `master/relays/s1-guide-q1/SITREP-planner-20260703-141628.md` (m-7 guide; answered entirely from locked text — no amendment, no master escalation). Reviewer RECONCILE on the DESIGN dispatch: approve, no blockers (`s1-core-design/RECONCILE-orchestrator-reviewer-20260703-141544.md`), one watch item (Implementer is CC-context during DESIGN until the Template-I review request — expected shape, no action).

**Q-A CONFIRMED** — the pair's recovery line IS the locked line ("recovery = dumb replay" is GRILL_LOCK verbatim, m-7 :186). S1 builds: §2.2 intake journal with FULL semantics (fsync-before-FIFO, ack-on-outcome, `intake_id` on every entry, clear-on-pop atomic with outcome commit, content-hash dedupe) + §4 pivot + dumb replay (staging cleanup :91, projection rebuild :92, binding-table restore :93, re-enqueue intake−outcomes in arrival order :94, wake re-issue). S1 does NOT build: genesis, quarantine disposition, GC/segment rotation, the reified phase-0→4 machine (:95). Resolution of the apparent IN/OUT tension: journal *semantics* in S1, journal *lifecycle machinery* in S2. **F9 runs whole in S1** — name the fixture F9 itself, not an F9-minimal variant; S2 re-runs it under added machinery.

**Three shape-right-from-S1 recovery constraints (hard, same class as the pivot):**
1. Canonical record = self-contained **checksummed** record file from the very first commit (:81) — checksum field is S1 format; quarantine disposition consuming a mismatch is S2.
2. Outcome records reference `intake_id` from the very first record — clear-on-pop is record-shape, not a recovery feature.
3. **Rebuild-before-open** — no `submit` accepted until staging cleanup + projection rebuild + re-enqueue complete (the phase-ordering discipline without the reified machine).

**Q-2 CONFIRMED** — B4 = typed local ODB-item production only; egress scan dormant is the locked posture (§9 :132, activation is step-(d)); ⑤ is S4-bound (Step-1-wide ledger item, not S1's to record — only to not foreclose).

**Three ⑤-shaped outbox constraints (hard):**
1. Outbox enqueue = a loop mutation committing a store-visible queue item through the one-pivot discipline (F11 lists it); no side path may produce an outbox item.
2. Produce-only, claim-honest: no drain/external send; no claim the scan is live; "only egress" wording carries the §9 qualifier + D5 residual verbatim.
3. Open envelope: ODB field set is m-6/m-3-owned; S1 types only the minimal envelope, does NOT close the schema, does NOT pre-build `model_name`.

Disposition: de-provision relay `s1-core-design/SITREP-orchestrator-planner-20260703-142800.md` issued to the pair; the six constraints join the DESIGN hard-constraint set; the plan cites ⑤ (if at all) as the S4-bound §C4 carry.

## 2026-07-03 — design-completion reconciled; blocker-2 narrowing RATIFIED (orchestrator); PROCEED-TO-PLAN issued

Inputs: design-completion SITREP (`s1-core-design/SITREP-planner-20260703-152941.md`), approving DESIGN-REVIEW (`s1-core-design-r2-review/DESIGN-REVIEW-implementer-20260703-152445.md`, verdict approve at main@5622516), r1 must-revise (`…-151318.md`, five blockers, all folded + fold-verified in r2).

Verified against disk: design doc exists at designs/s1-slice-1-design.md, r2 = main@5622516 (r1 = e8faeed); lineage shape correct for the later gated PLAN (approve parents to the r2 Template-I request, same DESIGN_DOC_ID); clean tree claims consistent.

**Operator decision recorded during design:** D-1 conductor stack = **Go** (operator, 2026-07-03, inline; Python/TS/Rust rejected, comparison logged in the doc).

**Blocker-2 narrowing — RATIFIED at orchestrator tier** (requested by the completion SITREP; implementer approve explicitly excluded it):
- The narrowing: S1 renders the `grant` field (`{dispatch-impl, dispatch-merge}`) on operator/orchestrator forms ONLY (the frozen m-2 :177 rule verbatim); conditional pair-Planner delegated-dispatch rendering + the m-2 :167 lineage walk land with the full lineage engine (S3). Until S3, delegated dispatch through the conductor's rendered form is unavailable to pair Planners; it remains available via operator/orchestrator forms.
- Basis: this is a *consequence of already-ratified sequencing*, not a new product decision — the operator-ratified charter puts the full FieldSpec registry and lineage engine in S3; the design simply refuses to smuggle a slice of S3's lineage machinery into S1. It is contract-faithful to frozen m-2 text (implementer-verified :166-177), and S1 governs no live delegated-dispatch traffic (S1's own build governance runs on file relays, not the conductor; merge is human-gated at S1 close).
- Ratification conditions (bind the PLAN): (i) the narrowing is stated explicitly in the S1 PLAN so the m-7-guide + master-VP gate reviews it — their gate is the authoritative check above me; (ii) claim honesty — any S1 doc/surface describing grants states the S3 landing plainly; (iii) no S1 schema/format decision may foreclose the S3 conditional-render landing.
- Veto path: operator may override this ratification at any time before the plan gate; the CC trail carries it.

Disposition: **PROCEED-TO-PLAN** issued (`s1-core-plan/PLAN-orchestrator-planner-20260703-153356.md`) — sequencing only; the gated PLAN lock stays in the pair-Planner seat (`DESIGN_LOCK_ID: s1-slice-1-design`, parent = the approving r2 DESIGN-REVIEW). Delegated dispatch is granted CONDITIONALLY: implementer plan-review approve + SCOPE_DIFF all-in + **the external m-7-guide + master-VP plan gate and both m-1/m-2 fidelity approves relayed into .relays/s1/** — no `DISPATCH IMPL` is live before all of those.
