## DESIGN — s6 = THE TRANSPORT FIX (in-Step-1 per the operator ruling): the design-amendment phase opens — m-1 owns THE parenting fork (grilled), m-7 engine liveness, m-2 the single codec; F1–F17 is the spec seed; every finding gets a disposition

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s5-exit-gate
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — the operator's in-step ruling (2026-07-06) is its basis; the parenting-fork GRILL and the amendment locks are the downstream human/VP gates
GRILL_REQUIRED: yes — the m-1 parenting-model fork (hard-to-reverse; multiple downstream choices hang on it); grilled at master with the operator when the m-1 decision packet returns
IN_REPLY_TO: master/relays/s5-exit-gate/RECONCILE-orchestrator-planner-20260706-174412.md
FROM: master.orchestrator-planner
TO: m-1.planner, m-7.planner, m-2.planner
CC: master.orchestrator-reviewer, operator, m-1.implementer, m-7.implementer, m-2.implementer
SUBJECT: design the transport fix as review-driven AMENDMENTS to your locked designs (the c6/s2-amend path, never silent re-design) — per-pair amendment docs + the m-1 fork decision-packet; exit = every F1–F17 item dispositioned {amendment | in-slice fix | wontfix-with-rationale}, the fork grilled+locked, VP co-sign; then master dispatches the s6 build slice

**Basis.** s5 is CLOSED and RATIFIED (2026-07-06; `RECONCILE.md` § s5); same session, the **operator ruled the transport fix in-Step-1** — the step's goal is "remove the operator-as-transport," and a transport that livelocks under multi-seat load (F11) has not durably removed it. **s6 = the transport fix.** This dispatch opens its **design-amendment phase**: the fixes touch your LOCKED designs, so they go through review-driven amendment (Cardinal rule 1; the c6 / s2-amend-m-1 precedent) — pair-internal adversarial review, master integration, VP co-sign — never silent re-design.

**The spec seed (read first):** `master/TRANSPORT-FINDINGS-2026-07-06.md` — F1–F17 with file:line evidence, the root-cause statement (**"one envelope, many judges, no reconciliation protocol"**), and the non-binding design seeds. The archived store (`~/frank-archives/frank-team-store-s5-dogfood-20260706`) holds the real failing traffic. Baseline: `frank/` `main @ 7e5c527` (tag `s5-close`).

### Ownership (the durable-domain split; seams named)

- **m-1 — store/lineage (THE fork):** the parenting/anchor model — **{F11 livelock, F4 projection-order candidacy}** — plus anchor-pool hygiene (**F10**: rejects in the pool / in `project()`), waiver scoping+retraction (**F17**), and **the F14 store-lock INVARIANT** — the sole-governed-writer / store-isolation contract is yours (the charter's m-1 store-isolation ownership); you rule the lock's *semantics* (what is locked, when held, staleness/takeover), while **m-7 owns its runtime enforcement** (see the seam below) *[VP must-revise fold, `RECONCILE-orchestrator-reviewer-…-180943`]*. The central candidate (from the seeds, NOT pre-decided): **conductor-computed PARENT** — the engine stamps lineage server-side exactly as it stamps FROM; the seat never supplies what the engine knows. Alternatives to weigh honestly: a widened/validated candidate set; a lineage lease/token. **This is the hard-to-reverse fork: return it to master as a DECISION PACKET** (options, consequences, your recommendation — the M-2 pattern) **for the operator grill before you lock anything.** Everything else in your scope may lock on pair review without waiting.
- **m-7 — engine liveness + ops:** digest stability under concurrency (**F5** — lease or bounded-tolerance revalidate; a submit racing an unrelated accept must not bounce), intake identity (**F9** — idempotent intake, 1:1 anchor restored; the C4.1 promise made true), live seat enrollment or an honest bounce choreography (**F15**), **F14 runtime enforcement** (the lock's process choreography, startup/refusal/takeover behavior, and fixture execution — the *invariant* is m-1's, above), shim reconnect-on-first-call (**F16**), bounce-detail parity for lineage-class rejections (**F3**).
- **m-2 — the single interpretation:** **ONE canonical envelope codec** consumed by render, typed-validate, the lineage gate, and delivery projection (**F6** silent recipient drop, **F7** the CC deadlock — kill the class, not the instance), and the schema-advertises/validator-rejects family (**F2** dead-edge parent offers, **F13** rejected-advertised-tokens, **F12** the waiver-flag typing vs its design intent).
- **The three-way seam (name it in each doc):** parent computation touches all of you — who stamps (m-7's loop), what the form renders (m-2), what the store anchors (m-1). The m-1 fork decision constrains the other two; design your halves as consumers of whichever fork branch wins, the way m-7 built its M-2 mechanics "composable with any m-6 signal-set ruling."
- **The F14 seam (named per the VP review; VP-confirmed `…-182326`):** m-1 owns the lock *invariant* + semantics; m-7 owns the *runtime enforcement* (startup, refusal, stale-lock takeover choreography). **[VP-W1] F14 must NOT collapse back to an ops-only task in any downstream doc** — the invariant stays m-1's in every artifact. **[VP-W2] F14's design-phase disposition must cite BOTH halves** — the invariant statement (m-1) and the runtime-fixture obligation (m-7).

### Constraints (unchanged; restate in your docs)

The locked mechanism vocabulary is untouched: byte-exact `{accepted, rejected, held}`; the seat surface exactly `submit`/`project`/`read`; channel-stamped FROM; I-PH on every seat-delivered surface; the claim ceiling stays **tool-mediated confusion-resistance** (§C4.3 — no wrap/by-construction creep). No Step-2 observe pre-work. Amend only where the ledger demands; the credit list ("what held") is load-bearing — crash-atomicity, FROM-stamping, and I-PH are NOT open for redesign.

### The design-phase exit (what master accepts)

1. **Every F1–F17 item carries a disposition:** {design amendment (yours) | in-slice build fix (no lock touched) | wontfix-with-rationale}. No silent drops — F8 stays the recorded numbering gap.
2. **The m-1 fork: decision-packet → master → operator GRILL → GRILL_LOCK → then the m-1 amendment locks. [VP-W3] this gate is SEPARATE and non-collapsible** — no m-1 lock before the grill, whatever the other pairs' pace. The other amendments lock on pair review as they finish.
3. **Per-pair amendment docs** (dated, in your domain dirs), each pair-reviewed (the planner designs, the implementer grills — your standing D5 recast), each stating its fixture obligations for the build slice.
4. **VP co-sign** on the integrated amendment set (master integrates the three docs + the seam).
5. Then master dispatches **the s6 build slice** (a NEW slice team, m-7-guided, m-1/m-2 fidelity — the standing pattern), whose exit gate already has its shape: the full battery + the fix fixtures + **the step-exit test on the fixed conductor** — the ROADMAP:83-85 legs, §7-applying s5's registry to a fresh blessed store as the first live act (the adapted [VP-W7]), and **the F11 regression leg: the archived dogfood traffic pattern replayed without livelock.**

### Not authorized by this relay
No code, no `frank/` edits, no lock changes before the reviews above, no build-slice work, no transport relaunch. Design-only; the fork does not lock without the grill.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s6-design` — run below.
- Basis chain: the operator ruling (recorded: dashboard · `RECONCILE.md` § s5 tail · `ROADMAP.md` Step-1 note, all 2026-07-06) · `TRANSPORT-FINDINGS-2026-07-06.md` · `RECONCILE.md` § s5 (the ratified close) · `frank/` `main @ 7e5c527`.

ACTIONS_GIT_REF: none — design-phase dispatch; no git action, no `frank/` edit; the ruling recordings are the docs-workspace edits cited above.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `7e5c527` (tag `s5-close`), clean.
Next requested action: VP gate COMPLETE — must-revise (`…-180943`, F14 mis-home) → revised per shape → **VP CONFIRM (`RECONCILE-orchestrator-reviewer-20260706-182326`; watchpoints folded as [VP-W1..W3])**. The operator relays this dispatch to the three pairs (m-1, m-7, m-2 — all in TO, all owe amendment designs); each pair designs its amendments; the m-1 decision packet returns to master for the operator grill; master integrates → VP co-sign → the s6 build-slice dispatch.
