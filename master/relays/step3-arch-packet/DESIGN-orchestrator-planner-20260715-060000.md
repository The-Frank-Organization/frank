## DESIGN — packet r4 CANDIDATE for exact re-review: SHA-256 `2d240eb6…`; folds VP F15–F17 as bounded source reconciliation (no product fork, grill NOT rerun) — exact SITREP E0-carrier envelope · m-5 coordinated-first-stage + staged propagation · canonical G8 lock line made non-transitive

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — on your clean re-review, the operator ratifies THIS exact candidate hash; this leg seeks adversarial re-review only
GRILL_REQUIRED: yes — required AND SATISFIED: `GRILL_LOCK: step3-arch-reframe-grill` remains closed; F15–F17 are bounded source reconciliation and reopen no operator decision
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-053000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-5.implementer, m-6.planner, m-7.planner, m-8.planner, m-9.planner
SUBJECT: packet r4 candidate `2d240eb6…` — exact SITREP E0-carrier envelope (F15) · m-5 coordinated-first-stage + staged propagation + charter-delta line (F16) · canonical GRILL_LOCK G8 line + typed-grant mandatory + no-lock qualifier (F17); F1–F14 untouched

Partner — `053000` accepted F11–F14 and raised three bounded source-reconciliation findings. r4 folds **F15–F17 only**, each a source reconciliation inside the accepted topology — no product fork, **grill not rerun**:

**`master/STEP-3-ARCH-AMENDMENT.md` r4 SHA-256 `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.**

### The three folds
- **F15 — exact E0 carrier envelope.** §3a now pins the carrier as the m-9 worker's existing **`PHASE: SITREP`, `AUTHORITY: report-only`, `HUMAN_GATE_REQUIRED: no`** relay — the landed *tested non-authority shape* (`lineage.go:39-58` returns authority-bearing for any grant/human-gate/A-gate/PLAN·IMPL·REVIEW-FOLD·MERGE-GATE·LIVE-VERIFY/impl·merge·live·fold/non-SITREP-non-RECONCILE-planner record; a plain worker SITREP is non-authority, `lineage_test.go:14-30`). **No `grant`, no `gate_category`, no gate-resolution/authority field**; routing **`TO: master.orchestrator-planner`, `CC: m-3.planner`**. Evidence separation pinned: top-level `EVIDENCE_TARGET`/`achieved_evidence` = **carriage only**; the **body event carries its own `event_evidence=E0` + `event_integrity=self_reported`** so a top-level E1/E2 relay observation **cannot upgrade the embedded event**.
- **F16 — m-5 sequencing + propagation.** The dependency graph now makes m-10 boundary design **and** the m-5 ceiling-host amendment a **coordinated first stage** that interface-locks the shared ceiling contract **before** any m-8/m-9 consumer lock; m-5 is **removed from the step-3 parallel authoring list** (it's a stage-1 item). m-5 is **added to the charter-delta line** (sole policy owner; enforcement host moves to m-10). Propagation is **staged**: the immediate ratification fold records the topology + the **pending, non-consumable m-5 amendment gate** (the locked m-5 design is **not** silently rewritten); the replacement flow then creates the m-5.planner/-implementer amendment and records its approved design-of-record fold **before** m-10/m-9 lock — both stages named.
- **F17 — canonical G8 wording + governed-action rule.** The **GRILL_LOCK Resolved-decisions G8 line** is replaced with the r3 non-transitive/ceiling/E0/typed-grant wording (a downstream reader following the lock now gets the bounded rule, not the pre-F11 broad copy). §8b item 3 sharpened: a direct instruction may authorize a directly-addressed **app-side** action within the current ceiling, but a **conductor-governed action requiring a typed grant/lineage edge goes through the sanctioned typed-grant branch — the direct message is context, not a substitute**. Item 7 clarified: direct content in a fresh submission stays **non-authority unless it independently satisfies the grant grammar**. The sequence "No lock" qualified as **no `DESIGN_LOCK_ID`/architecture ratification** (the grill lock exists).

### Untouched (your accepted set)
F1–F10 (043000) and F11–F14 (053000) — the topology, principals, connector/E3 ordering, no-retry, state/recovery, carry ledger, m-10 pair, credential owner, and the F11–F14 repairs — all unchanged.

**What I request:** your **exact-candidate re-review bound to `2d240eb6…`** — confirm the three source reconciliations. On a clean return, the operator ratifies this exact hash, then the atomic source fold + refreshed consumer audit + replacement dispatches (m-10 + the m-5 ceiling amendment as the coordinated first stage). Five holds remain; no lock/PLAN/T4 code token crosses that gate.

## Verification
- Candidate `master/STEP-3-ARCH-AMENDMENT.md` = SHA-256 `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69` (fresh `shasum`); r1/r2/r3 superseded. ROADMAP frozen `3977c9f2…`; kickoff `983508fc…`.
- F15 grounded at `lineage.go:39-58` + `lineage_test.go:14-30`; F16 reconciles §8 graph/charter/propagation internally; F17 reconciles §9 lock line vs §8b.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-060000.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — edited `master/STEP-3-ARCH-AMENDMENT.md` (r3→r4: §3a envelope, §8 graph/charter/propagation, §8b items 3/7, §9 GRILL_LOCK G8 line + no-lock qualifier, header) and created this relay + its INDEX row. The §9 durable `GRILL_LOCK` persists; **no `DESIGN_LOCK_ID`, no architecture ratification**; no `frank/`, code, source-of-truth fold (ROADMAP frozen/untouched), credential, provider, live-store, PLAN, or lane-resumption action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns an exact-candidate re-review bound to `2d240eb6…` (approve / must-revise) TO master.orchestrator-planner, CC operator; on approve, master routes the exact candidate to the operator for hash-bound ratification, then the atomic source fold + consumer audit + replacement dispatches (m-10 + m-5 first).
