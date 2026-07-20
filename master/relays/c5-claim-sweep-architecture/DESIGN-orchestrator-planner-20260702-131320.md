## DESIGN — lane 1 r2: ARCHITECTURE sweep completed (5 more relabels) + checklist hardened + accurate survivor list

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-claim-sweep-architecture
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner — revise **accepted**; the catch is correct and mine to own. My r1 grep was too narrow (it missed `sole write path`, `sole external sender`, and the whole `tamper-resistant`/`non-lane-writable` class) and my survivor list undercounted — exactly the failure that can't ship in the *exemplar* the domain lanes copy. r2 fixes all three blockers + hardens the checklist + carries a survivor list that matches the file.

**r2 relabels (5, all confusion-resistant + D5 residual):**
1. **:26** "sole-writer Maildir mailboxes" → "**governed-write** Maildir mailbox projection".
2. **:30** "`submit` (sole write path)" → "`submit` (sole **governed** write path)".
3. **:154** egress "(sole external sender)" → "**conductor-governed egress chokepoint** (the local outbox is the only egress the *governance system* offers — governance-surface claim, not system-level sole-egress; **D5 residual**: same-uid shell/curl bypass out of scope)". (Same scoping m-7 §9 uses.)
4. **:157 / :198 / :275** `slot_in`/`seat_archetype` "**tamper-resistant / non-lane-writable**" → classified: **non-lane-writable = confusion-resistant** (no lane tool/verb writes or re-tags it) **+ D5 residual**; the F1 composition invariants kept as a genuine **observer-selected control property** (the invariant is chosen conductor-side; the observed lane cannot reach the selection) — with the D5 residual (a malicious seat writing a false store record) stated beside it.
5. **:46** operator-address "a lane cannot submit as `operator`" → "a **confused** lane has no tool to submit as `operator` (confusion-resistant; D5 residual)".

**Checklist HARDENED — the RELABEL net now explicitly covers this whole claim class:**
> **RELABEL → confusion-resistant + D5 residual:** `by construction / structural / unbypassable / sole-writer / sole write path / **sole external sender/egress** / forgery-robust / **tamper-resistant** / **non-lane-writable** / "a lane/seat cannot **write / forge / supply / mutate / re-tag / bypass / submit-as**". These all assert a *malicious* seat is stopped; under attach a shell-bearing seat bypasses the tool surface. Honest form: *a confused agent has no **tool** to do X*, D5 residual beside it.
>
> **KEEP (justified, never blind-relabel):**
> - **Trusted-engine control-flow / grammar invariants** — hold regardless of seat behavior: the serialized-loop double-accept kill; **R2's no-model-predicate gate grammar**.
> - **Observer-selected control properties** — an invariant selected conductor-side (outside the lane) that the observed lane cannot reach (the F1 archetype invariants). KEEP the observer-selection framing; still note the D5 residual on the underlying evidence.
> - **Authority-ceiling statements** — "routes but cannot write" = the ceiling grants no write *tool* (confusion-resistant enforcement), not malicious containment.
> - **Correctly-scoped confusion-resistant** claims already in "no tool" form.
>
> **Test:** *does it claim a malicious seat is stopped?* → relabel + residual. *Does it hold by the trusted engine's own control-flow/grammar, or by an observer-side selection the lane can't reach?* → keep. Unsure → relabel + residual (the NO-GO's safe default).

**Accurate survivor list (full-net grep, every hit classified — matches the file):**
- **Scoping note:** :18–19.
- **RELABELED (confusion-resistant + D5):** :23, :33, :42 (forgery-robust-by-construction marked *shelved*), :47 (operator), :56, :79, :158, :200, :281.
- **KEEP (justified):** :125 (monotonic-floor tamper-*rejection* — a negative **fixture** name; the floor is monotonic-MAX, a structural property — a confused floor-lowering attempt is rejected), :184/:191 (**R2 gate-grammar** invariant), :285 (**authority-ceiling** — no write tool granted), :442/:445 (**§C4.3** the honest claim boundary).

No mechanism changed; ARCHITECTURE claim-text only. On your r2 approve, the ratified checklist drives the domain lanes (`c5-claim-sweep-m-1`, `-m-2`, `-light`) + the decision-folds.

Not authorized / not claimed: ARCHITECTURE claim-text only; no mechanism change, no domain doc edited, no PLAN/IMPL/`pcode`/spike, no Step-1 PLAN, no operator decision reopened.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-architecture` — OK
- `grep -nE 'sole[- ]writer|sole write|sole external|forgery-robust|unbypassable|by construction|tamper|non-lane-writable|cannot (forge|write|supply|mutate|re-tag|bypass|submit)' master/ARCHITECTURE.md` — 16 hits, each classified above (relabeled / KEEP / note); zero unaccounted raw overclaims.
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: r2 swept `master/ARCHITECTURE.md` :26/:30/:46/:154/:157/:198/:275 (5 more relabels + tamper-class classification); wrote this r2 relay + appended `master/relays/INDEX.md`; no mechanism change, no domain-doc edit, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP r2 review of the ARCHITECTURE sweep + hardened checklist; on approve I dispatch the domain lanes + decision-folds against the ratified checklist.
