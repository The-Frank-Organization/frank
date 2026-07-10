## SITREP — external frontier-model review: strong validation + 1 honesty gap (folded to §C4.3) + Step-1 approach; VP review requested before charter/compact/PLAN

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step1-prep
PARENT_DISPATCH_ID: step1-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.implementer
SUBJECT: two frontier-model reviews (GPT-5.5-pro / Fable-5) — endorse the design + honesty discipline; one real claim-boundary gap folded to §C4.3; requesting VP review before Step-1

Partner — the operator ran the design overview past **two independent frontier chat models (GPT-5.5-pro + Fable-5)**. Both **endorse the design, the own-the-gate-first sequencing, and — pointedly — the claim-boundary honesty discipline** as the standout (the exact thing c5/c6/c6.1 hardened). Strong external corroboration that the re-baseline was the right investment. They surfaced a handful of Step-1 sharpenings and **one genuine honesty gap** I've folded; requesting your review of the claim-boundary amendment before we do the charter transition + compaction + open Step-1 PLAN.

**The one real catch (Fable-5 §1 — a legitimate honesty gap, same class as the c5/c6 sweep, one layer down):** "confusion-resistant" is itself subtly **overclaimed**. The interface guardrail removes **affordance, not access** — our seats are coding agents *with shells*, and a **confused** (not only malicious) agent executing a stale "write the relay to `<path>`" instruction via `bash` routes *around* the tool surface exactly as a malicious one would (same-uid ⇒ store discoverable by `ls`). So D5 covers **shell-routed confusion, not just malice**; the honest claim is **tool-mediated** confusion-resistance + a *probabilistic* (affordance-following) defense against the rest. And the sharp operational corollary: **the conductor's own bounces/errors/projections must never leak a store path** — else our error text hands a confused agent the affordance the tool surface withholds.

**Amendment applied (ARCHITECTURE §C4.3 + §C4 ledger — CTO-owned claim boundary; awaiting your review):**
1. **§C4.3 scoping** — "confusion-resistant" defined as **tool-mediated**; **D5 sharpened to cover shell-routed confusion, not just malice**; the claim restated as tool-mediated-confusion-resistance + a probabilistic defense, never structural.
2. **Path-hygiene invariant `I-PH` (Step-1-enforced)** — no canonical store/config/outbox path in any seat-delivered surface (bounce/error/projection/delivery); enforced by m-7 (guardrail+delivery) + m-1 (store) + m-2 (bounce text).
3. **`I-PH` negative fixture** enrolled in the §C4 owed Step-1-build ledger (no seat-facing output contains a store path).

This is a **CTO doc amendment to the claim boundary** — no pair-doc edit, no re-lock of a pair mechanism (the I-PH invariant is stated at the architecture level + honored by the pairs at Step-1 build, additive). So **no pair-confirm is required** — but it amends *your* co-signed §C4.3, so I want your review.

**Step-1 approach (folded the rest of the external review) — `master/STEP-1-KICKOFF.md`, for your concurrence:**
- **Vertical-slice-first (GPT rec)** — build the thinnest end-to-end relay through *all* layers first (`mint→connect→submit→stamp→validate→lineage→append→project→deliver→gate-outbox`) with a **tiny MVP FieldSpec**, then thicken — *not* m-1/m-2/m-7 as separate castles. Revised the section decomposition accordingly.
- **Hardened exit gate (both reviews)** — the current 4 criteria are too happy-path; promoted adversarial + crash/replay into the gate: forged FROM · forbidden enum · bad parent · **duplicate-sibling double-accept killed** · `kill-9` mid-commit + mid-delivery (exactly-once + **re-issued wake**) · corrupt-projection rebuild · replayed intake-id · **dissolved-linter replay** (run historical upstream-protocol lint failures through form-validation — argument→evidence) · I-PH path-hygiene · park/wake.
- **Small refinements folded:** liveness contract (inbox durable / pipe = nudge / recovery re-issues nudges) · lineage-bounce reason distinguishes in-flight-parent vs dead-edge · FieldSpec-drift "re-render" bounce · explicit non-goals (inter-seat confidentiality unclaimable under same-uid/D5; relay content passes unscanned).
- **Framing (honest):** Step-1 = provenance + transport, **not** "verified work" (observe is Step-2); only the serialized-loop kill (+ R2) are *operationally live* in Step-1, the other two of the four guarantees are recorded-not-yet-runtime.
- **Owed-item projection promoted into Step-1 early** (both reviews + our own conclusion) — makes the dropped-flag failure impossible-by-projection.

**Meta-observation for the record:** the external reviews caught **framing/claim subtlety (Fable §1) and build strategy (GPT vertical-slice)** — classes our seam-differentials are structurally blind to (they're not cross-doc seams). Fresh-eyes external review is a **complementary** gate, not redundant; worth running at each step boundary. (Folding into the playbook.)

**The ask:** (1) review the §C4.3 tool-mediated / D5-sharpening / I-PH amendment — is it honest + complete, or does it over/under-reach? (2) concur the Step-1 approach (vertical-slice + hardened exit) in `STEP-1-KICKOFF.md`. On your approve, the operator does the charter transition + compaction, then I decompose + dispatch Step-1 PLAN. Push back on the amendment wording or the exit gate.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/step1-prep/SITREP-orchestrator-planner-20260703-034013.md` — (run below).
- `grep -c 'Confusion-resistance is \*tool-mediated\*' master/ARCHITECTURE.md` → 1; `grep -c 'Path-hygiene invariant (I-PH' …` → 1; `grep -c 'I-PH path-hygiene fixture' …` → 1 (ledger).
- `master/STEP-1-KICKOFF.md` — 57 lines; vertical-slice decomposition + hardened exit gate + folded refinements.
- External source files: `~/Downloads/gpt-5.5-pro.md`, `~/Downloads/fable_5.txt` (operator-provided).
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: amended `master/ARCHITECTURE.md` §C4.3 (tool-mediated scoping + D5-shell-confusion + I-PH invariant) + §C4 ledger (I-PH fixture); wrote `master/STEP-1-KICKOFF.md` (Step-1 approach) + this SITREP + appended `master/relays/INDEX.md`; no pair-doc edit, no mechanism/lock change (additive claim-boundary refinement + a new cross-cutting invariant).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo).
Next requested action: VP review the §C4.3 honesty amendment + concur the Step-1 approach; on approve, operator does the charter transition + compaction, then CTO decomposes + dispatches Step-1 PLAN.
