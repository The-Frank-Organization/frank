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

## 2026-07-03 — m-7 guide advisory read of the design reconciled: faithful, zero must-fix; two should-fix folded to the pair

Input: `master/relays/s1-guide-design-read/SITREP-planner-20260703-154742.md` (advisory, per the operator-suggested early read; locked cites re-verified by the guide that session).

**Verdicts on the six asks:** C7 fold faithful to §4/§6/F11 and the right idiom (instantiates the locked presence-based-derivation pattern); R-1..R-3/O-1..O-3/F9-whole all landed, no drift; B1 narrowing consistent with §16 modulo the held-visibility qualifier (should-fix ①); §J2 set **pre-verified byte-exact** by the guide against ARCHITECTURE :110-115 (formal byte-custody stays with the m-2 fidelity review); grant narrowing forecloses nothing m-7-side (authoritative in-loop check is what matters; A2 covers it); Go stack fine, D-2 faithful.

**Should-fix (fold as r3 one-paragraph edits or PLAN-carried lines — either satisfies the gate):**
- **① held-visibility posture:** the doc commits `held` records (fixture H) but never states how the operator learns one exists or what resolves it in S1. B1 needs the locked m-2 carve-out beside it (`held` is operator-visible by locked text — m-7 :100, m-2 :76/:376 — never consumed by downstream work authority). Pick and state one compliant shape: (i) guide-recommended — a `held` record derives an ODB/outbox item via the C7 derived-work mechanism (one more derived-intent class; §6 self-exclusion survives); or (ii) stated deferral — held terminal for the intake, author gets the typed outcome, operator sees it via project()/INDEX on the operator channel, resolution flow is S2, swept by SWEEP. Silence fails the gate.
- **② wake-push fallback invariant:** the D-3 MCP-fallback invariant list must add **server-initiated nudge push on the held per-seat connection** (m-7 §8.3); a fallback that drops push breaks L1/W1 (polling-only wake is exactly what the locked design rejects). The capability check must test for it explicitly.

**Advisory sharpenings (carry into doc/PLAN, no re-architecture):** state that C7 derived-work completion executes on the single-writer commit path (loop goroutine, or recovery single-threaded before it starts); PLAN notes S2's owed-item projection generalizes the C7 mechanism (C7's scan becomes an instance of it at S2, not a parallel mechanism). Positive notes recorded (crash-point registry reuse, D-11 correctly-scoped by-construction claim, outbox pivot, binding-table fold, D-10 replay posture) — no action.

Disposition: fold relay `s1-core-plan/SITREP-orchestrator-planner-20260703-155605.md` issued to the pair; ①/② join the plan-gate rubric; PLAN drafting continues unpaused.

## 2026-07-03 — plan-completion reconciled; README fence RULED IN; the four external-gate packets issued

Inputs: plan-completion SITREP (`s1-core-plan/SITREP-planner-20260703-163256.md`); approving PLAN-REVIEW (`s1-core-plan-lock-r2/PLAN-REVIEW-implementer-20260703-162820.md`, approve at main@a24bf57); r1 must-revise (`s1-core-plan-lock/PLAN-REVIEW-implementer-20260703-162202.md`, two blockers, both folded).

Verified on disk: plan r2 = main@a24bf57, design r4 = main@3882763 (r3 = 061882a — the guide should-fix folds: held-visibility shape (i), wake-push invariant, single-writer sharpening; r4 = two-source ODB envelope). Delegation lineage chain intact: plan-review approve → gated PLAN r2 (`s1-core-plan-lock-r2`, DESIGN_LOCK_ID + design-doc kind) → approving DESIGN-REVIEW (`s1-core-design-r2-review-implementer`). Ratification conditions 1–3 verified landed (narrowing stated in gated PLAN text; Task 12 README/SWEEP; guide-confirmed additive S3 render).

**ASK-1 ruling — root README.md IN-FENCE** (orchestrator's call per protocol). Rationale: it is the S1 honesty/claim surface serving ratification condition 2, SWEEP-covered, m-7 §16-swept. The PROCEED-TO-PLAN scope fence is amended by ruling relay `s1-core-plan/SITREP-orchestrator-planner-20260703-170259.md`; the pair cites it as SCOPE_DIFF row evidence. Fallback (move under docs/) not needed.

**ASK-2 — external-gate packets issued** (three relays, operator-carried):
- `s1-plan-gate/SITREP-orchestrator-planner-20260703-170259.md` → m-7.planner + master.orchestrator-reviewer (the formal co-gate; carries the design rev ledger r2→r3→r4, the checklist heads-up, and the narrowing statement).
- `s1-fidelity-m1/SITREP-orchestrator-planner-20260703-170259.md` → m-1.implementer (store-API usage; DI-2 realization flagged; four specific fidelity questions).
- `s1-fidelity-m2/SITREP-orchestrator-planner-20260703-170259.md` → m-2.implementer (FieldSpec usage; **§J2 byte-custody formally requested in-verdict**; five specific fidelity questions).

Standing state: pair HOLDS. `DISPATCH IMPL` becomes issuable by the pair Planner only when all four approve relays exist in .relays/s1/ + SCOPE_DIFF all-in (README row citing the ruling) + the remaining delegation conditions. Merge remains a separate human gate at S1 close.

## 2026-07-03 — external-gate verdicts reconciled: 3 APPROVE + 1 BLOCK (m-1 F-M1-1); bounded plan revision routed

All four verdicts returned same-day:
- **m-7 guide — APPROVE** (`s1-plan-gate/PLAN-REVIEW-planner-20260703-171032.md`, copy filed from master trail): checklist 7/7 PASS; both should-fixes verified folded; lineage chain verified on disk; "no code before the gate" verified held. One **non-blocking nit**: design D-7 :112 still keys crash-window text by `gate_record_ref` while r4 :115 makes `(source_kind, source_record_ref)` THE key — add a one-line parenthetical (gate-sourced item: `gate_record_ref` IS `source_record_ref` with `source_kind=gate`). Two standing IMPL watch-surfaces (heads-up, already in plan): F11 asserts EXACTLY-one pivot (not at-least-one); P1 captures push frames + tool descriptions, not just bounce/error text.
- **master VP — APPROVE** (`s1-plan-gate/RECONCILE-orchestrator-reviewer-20260703-170942.md`): lineage + fence + coverage + claim honesty confirmed; watchpoints — SCOPE_DIFF README row must cite the fence ruling; scoped-root lint blind spot on cross-dispatch design lineage noted (exact-file lint = proof of record).
- **m-2 fidelity — APPROVE** (`s1-fidelity-m2/SITREP-implementer-20260703-171027.md`): all five questions faithful; **§J2 byte-custody formally confirmed** (8 A + 4 B + other→A + `routing_escalation` non-member, cited to ARCHITECTURE :105-115); grant narrowing, held usage, R1 classification, MVP instantiation all faithful.
- **m-1 fidelity — BLOCKS** (`s1-fidelity-m1/SITREP-implementer-20260703-171028.md`): single finding **F-M1-1** — the m-1 PLAN carry :228-229 (credential generation/rotation/revocation) is acknowledged in design :185 ("mint/re-attach/revoke-on-remint") but plan Task 5 (:120-123) only asserts re-attach-same-credential; no generation/epoch, duplicate-mint, remint/recycle, or old-credential-rejection invariant or fixture (E1 rg-verified). All four specific questions otherwise no-finding (DI-2, submit pipeline, operator address, TOCTOU carry). Required revision: ONE lifecycle invariant + E2 fixture in Task 5, either shape — (a) one active generation per seat; remint atomically replaces binding, old credential rejects before staging; or (b) if S1 has no remint op: `Mint(existing)` returns existing binding or rejects; test that no stale/parallel credential survives.

Disposition: revision relay `s1-core-plan/SITREP-orchestrator-planner-20260703-171643.md` to the pair — bounded plan-r3 fold of F-M1-1 (+ the guide nit riding the same pass) + pair-Implementer narrow re-review; then I re-route the revised surface to m-1.implementer for narrow re-review. Guide/VP/m-2 approves do NOT need refreshing (m-1's own scoping; the revision touches only the m-1 surface + a design parenthetical). Dispatch remains blocked until m-1's approve lands.

## 2026-07-03 — F-M1-1 fold reconciled; m-1 narrow re-review routed

Input: fold-completion SITREP (`s1-core-plan/SITREP-planner-20260703-173107.md`); pair narrow re-review approve (`s1-core-plan-lock-r3/PLAN-REVIEW-implementer-20260703-172725.md`, parented to gated PLAN r3, which still parents to the approving DESIGN-REVIEW — delegation lineage intact).

Verified on disk: fold commit main@50b2b94 (design r5 D-3 single-credential-generation invariant, shape (b) — duplicate `mint_seat` ⇒ typed reject `ErrSeatAlreadyBound`, no second credential, binding table unchanged; plan r3 Task 5 interface + E2 lifecycle fixture; guide key-scheme parenthetical in D-7); cleanup commit main@d09278a (stale `revoke-on-remint` shorthand in §6.4 aligned to the no-remint wording — contract unchanged). No-broadening boundary held; F11/P1 watch-surfaces and OUT lines untouched per the narrow reviewer.

Disposition: narrow re-review packet `s1-fidelity-m1/SITREP-orchestrator-planner-20260703-173251.md` → m-1.implementer (revised surface only at main@d09278a; the four no-finding confirmations not re-asked). On m-1's approve: all four external conditions green → pair runs SCOPE_DIFF (README row cites the fence amendment) → delegated `DISPATCH IMPL` per the standing conditions. Merge remains a separate human gate at S1 close.

## 2026-07-03 — ALL FOUR EXTERNAL GATES GREEN; pair released to SCOPE_DIFF → delegated dispatch

Input: m-1 narrow re-review **APPROVE** (`s1-fidelity-m1/SITREP-implementer-20260703-174333.md`) — F-M1-1 closed on the revised surface (design D-3/:185/:219 + plan Task 5 :121/:123, all E1-cited; rg-verified no stale shorthand); prior no-finding confirmations unchanged.

**Final gate state (all verdict relays in .relays/s1/):**
| gate | verdict | relay |
|---|---|---|
| m-7 guide | APPROVE (7/7) | s1-plan-gate/PLAN-REVIEW-planner-20260703-171032.md |
| master VP | APPROVE | s1-plan-gate/RECONCILE-orchestrator-reviewer-20260703-170942.md |
| m-1 fidelity | APPROVE (r2, narrow) | s1-fidelity-m1/SITREP-implementer-20260703-174333.md |
| m-2 fidelity | APPROVE | s1-fidelity-m2/SITREP-implementer-20260703-171027.md |

Disposition: green-light status relay `s1-core-plan/SITREP-orchestrator-planner-20260703-175058.md` to the pair. The pair Planner now runs the mechanical SCOPE_DIFF (README.md row citing the fence amendment `s1-core-plan/SITREP-orchestrator-planner-20260703-170259.md`) and, only on `all-in` with the full standing condition set, issues the delegated `DISPATCH IMPL` (parent = the approving PLAN-REVIEW r3 `s1-core-plan-lock-r3-implementer-review`). Locks: PLAN `s1-slice-1-plan` r3 / DESIGN `s1-slice-1-design` r5 at main@d09278a. Remaining gates AFTER impl: the S1-scoped hardened exit gate (E2 fixtures; guide watch-surfaces F11-exactly-one + P1-push-frames), SITREP to master, human merge gate at S1 close.

## 2026-07-03 — S1 EXIT-GATE report reconciled; battery re-verified by orchestrator (E2); two deviations provisionally accepted pending guide concurrence; master SITREP issued

Input: exit-gate report `s1-core-impl/SITREP-planner-20260703-195800.md`. Build trail: delegated IMPL (`s1-core-impl/IMPL-planner-20260703-175416.md`, conditions verified) → 17 impl commits (139aaa9..9c1839e, task-per-commit) → 5-lens adversarial panel → REVIEW-FOLD round 1 (B1-B7, M1-M7; fold @7f66057, FOLD_SCOPE pre-filed) → round 2 (RB1/RB2/RM1; fold @80c5df5) → planner-verified residue folds → plan file-list absorption @964b120.

**Orchestrator verification (my own runs this session, E2):** `go test -count=1 ./...` — 15 packages ok (uncached); `go test -race -count=1` on fixtures/engine/recover — ok; `go vet ./...` — clean. The crash matrix is real child-process SIGKILL (planner-read f11_test.go:65-149: per-class kills at named crash-points, wait-status-verified, checksum + at-most-one-rename assertions; F9 whole: 5/2/3 across a real crash).

**Deviation rulings (provisional — routed to the m-7 guide for concurrence, since deviation 1 sits on the guide's declared F11 watch-surface):**
1. F11 breadth (plan says crash at EVERY crashpoint name × class; shipped = 7 representative class×point cases + registry-names-live static pin): PROVISIONALLY ACCEPT as S1-sufficient — every exit-gate-NAMED window is crash-covered (the charter gate, not the plan wording, is the acceptance authority; the full cross-product is S2-reusable machinery already built). Guide may bounce for the full sweep.
2. C7 re-crash-during-Complete covered by composition (crash-before-completion + idempotence double-run) rather than one literal mid-Complete fixture: PROVISIONALLY ACCEPT — the composed pair proves the same property; guide may require the literal fixture.
Neither was self-waived by the pair — correctly surfaced. Honesty checks: E3/E4 explicitly stated out-of-S1, merge explicitly not requested, verdict merge-blocked per protocol.

Disposition: S1 exit-gate SITREP to master (`s1-exit-gate/SITREP-orchestrator-planner-20260703-200108.md`, TO master.orchestrator-planner + m-7.planner) — the charter deliverable — carrying the gate evidence, the two deviations + provisional rulings for guide concurrence, and the merge question routed to the operator. Gate CLOSES on guide concurrence (or the required follow-up fold); merge remains the operator's human gate.
