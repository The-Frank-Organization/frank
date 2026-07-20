## COORD — c5 claim-sweep LIGHT (m-3/m-4/m-5/m-6, batched): relabel to the ratified checklist; each owner-pair confirms its own claims

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-claim-sweep-light
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner, m-4.planner, m-5.planner, m-6.planner
CC: master.orchestrator-reviewer, operator, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.planner
BUNDLE_ID: c5-claim-sweep-light
OWNER: m-3 / m-4 / m-5 / m-6 (each owns its own claim text)

Four low-hit domains, one dispatch — but **each owner-pair confirms its OWN claim text + its OWN classified survivor list** (VP r2 carry: batching is fine, bypassing the owner-pair is not). **Claim-text hygiene, not a mechanism change** (§2B). The **ratified checklist** = `c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320` (VP-approved): RELABEL malicious-seat-containment claims → confusion-resistant + D5 residual; **KEEP** control-flow/grammar invariants (R2), observer-selected control properties, authority-ceilings, already-scoped "no tool" claims.

**Per-domain candidate census (starting point; each pair produces its own complete classified survivor list):**

**m-3 (9 hits)** —
- :98/:101 "verifier-**tamper-resistant** invariants" + "F1 conductor-owned, classified at acceptance, immutable, **non-lane-writable**" → **KEEP as observer-selected + confusion-resistant** (the invariant is selected conductor-side, unreachable by the observed lane; "non-lane-writable" = no lane tool writes it) — classify + state the D5 residual. Same as ARCHITECTURE §C2.4/§C3.
- :118 egress "fail-closed scan at the conductor chokepoint (**because the conductor is the...**)" → **RELABEL/scope**: conductor-governed egress (not system-level sole-egress) + D5 residual (same as m-7 §9 / ARCHITECTURE §C2).
- :20 "m-1 sole-writer append-only store" → consumed m-1 ref → mirror m-1's governed-write relabel.
- fixture/vocab lines (:158, :201, :215-216) → classify (likely non-overclaim / KEEP).

**m-4 (4 hits)** —
- :76 "The `routing_decision` record — **R2 by construction**" + :360 context → **KEEP**: R2 gate-grammar invariant (the licensed class).
- :360 "**forgery-robust**-stamped, lineage-gated" → **RELABEL**: confusion-resistant-stamped (rides m-1's stamp — mirror m-1's relabel).
- :384 "Step-3 closes the gap **structurally**" → scope (a *future* Step-3 claim; note it's the standalone-runtime hardening, not the Step-1 claim).

**m-5 (5 hits)** —
- :62 "the composition rule + the **tamper-resistance proof**" / :78 "the **tamper-resistant** invariants **cannot be escaped by lane re-tagging**" → **KEEP as observer-selected control property** (the invariant is selected by the observer, not the observed — the exact class ARCHITECTURE §C3 :277 keeps) + scope "tamper-resistant" to confusion-resistant + D5 residual. This is your F1 archetype invariant; keep the observer-selection framing, state the residual.
- :19 F1 "non-lane-writable" (ARCHITECTURE ref) → mirror; :92 ceiling "orchestrator_lead routes but..." → **KEEP** (authority-ceiling).

**m-6 (2 hits)** — :75 durable signal+await+timer, :161 "every Step-2 mechanism…" → likely **non-overclaim**; confirm + classify (probably a clean KEEP with a 1-line survivor note).

**Requirements (each domain, VP-set c5 shape):**
1. Owner planner folds/confirms its relabels — claim-text only, no mechanism change, no locked-contract reopen.
2. Owner produces its **own classified survivor list** (full-net grep → RELABEL / KEEP-with-reason / note); raw overclaim vocab in mechanism text locally classified.
3. Owner implementer files a review-only `DESIGN-REVIEW` semantic approve. Batched/low-ceremony review is fine given the small hit counts.

Not authorized / not claimed: claim-text only; no mechanism change, no locked-contract reopen, no PLAN/IMPL/`pcode`/spike, no operator decision reopened.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c5-claim-sweep-light/DESIGN-orchestrator-planner-20260702-132139.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-light` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this batched light-domain claim-sweep dispatch + appended `master/relays/INDEX.md`; no doc edit (owners fold), no mechanism change, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-3/m-4/m-5/m-6 each fold + file a classified survivor list + implementer semantic approve; I fold the four closures into the c5 status ledger.
