## DESIGN — lane 1: ARCHITECTURE.md claim-sweep DONE (CTO); + the SEMANTIC CHECKLIST the domain lanes inherit

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

Partner — lane 1 done: I swept `master/ARCHITECTURE.md` (CTO-owned) to the honest claim line, per §2B. **Mechanisms unchanged; only malicious-lane CLAIMS relabeled.** This is the exemplar + the checklist the domain lanes (m-1/m-2/light) inherit; VP review requested.

**What I relabeled (adversarial-strength → confusion-resistant + D5 residual):**
- §1 header + model: "sole-writer stamping courier" → "**governed-write** stamping courier"; "sole writer of an append-only store" → "**sole writer through the governed `submit()` path**"; added a **§1 claim-boundary note** (this c1–c3 record predates the fork; all "sole-writer/forgery-robust/by-construction" below scoped confusion-resistant per §C4.3; D5 residual).
- §2 m-1: "**Identity by construction** … the sender cannot write the trusted field" → "**Identity is conductor-stamped (confusion-resistant)** … the seat has **no tool** to write it (D5: a same-uid code-executing seat can write store files directly — out of scope, §C4.3)"; **I1** "sole-writer/store-isolation" → "**governed-write** / store-isolation (not malicious-lane containment)"; **I2** "(forgery-robust `FROM`)" → "(**confusion-resistant** `FROM`)"; Transport "⇒ **forgery-robust by construction**" → "⇒ **confusion-resistant** identity; forgery-robust-by-construction is the **shelved** wrap milestone (GL D3)".
- §3 m-2: "One-FROM rule = **DMARC-by-construction**" → "One-FROM rule (closes the DMARC two-FROM confusion for a **confused** agent)".
- §4 seam: PARENT "confusion-robust → **forgery-robust by construction**" → "→ **confusion-*resistant*** (seat cannot supply `PARENT`; D5 residual)".
- §5: "m-1 owns who (**forgery-robust** identity + **sole-writer** store …)" → "(**confusion-resistant** identity + **governed-write** store …)".
- §C3 scheduler: "the **sole-writer** store as the checkpointer" → "the **governed-write** store".

**What I KEPT (legitimate — the checklist's other half):**
- **R2-preserving by construction** (§C2.3 :182/:189) — a **grammar invariant**: no gate predicate references the model field, so no seat (confused *or* malicious) can make a gate read the model. Holds structurally regardless of seat behavior — the same class as the licensed serialized-loop claim. **KEEP.**
- **§C4.3** (the m-7 claim boundary) — already the honest reference ("confusion-resistant; sole licensed by-construction = the serialized-loop kill; wrap/sole-writer-by-construction shelved"). **KEEP.**
- ceiling partial-order "routes but cannot write" (§C3 :274) — an *authority-ceiling* statement (no write tool granted), confusion-resistant; not a malicious-containment claim. **KEEP.**

---

## THE SEMANTIC CHECKLIST (domain lanes m-1/m-2/light apply this to their own docs)

**RELABEL → confusion-resistant + state the D5 residual** — any claim that a mechanism prevents a **malicious / code-executing / same-uid** seat from doing X "by construction / structurally / unbypassably / sole-writer / forgery-robust / cannot forge/write." Under attach, a shell-bearing seat bypasses the tool surface; the honest claim is *the confused agent has no **tool** to do X* (`GRILL-LOCK-deployment-fork-2026-07-01.md` D4), with the D5 accepted-risk (config / store-write / operator-FROM) stated beside it.

**KEEP (do not relabel):**
1. **Genuine control-flow / grammar invariants** that hold regardless of seat behavior — the serialized-loop double-accept kill; R2's no-model-predicate gate grammar. These are properties of the trusted engine's own structure, not malicious-containment claims.
2. **Confusion-resistant properties correctly scoped** — "the seat has no tool to X" (already honest).
3. **Claims about the conductor's own behavior** — not about containing a malicious seat.

**The test:** *does the claim assert a **malicious** seat is stopped?* → relabel. *Does it hold because the trusted engine's control-flow/grammar makes it true regardless of the seat?* → keep. When unsure, relabel + state the residual (the safe default the NO-GO earned).

Owning-implementer semantic review = confirm each relabel is (a) faithful to the mechanism (no mechanism change), (b) on the right side of the test, (c) carries the D5 residual where scoped. Batched review OK for low-hit domains.

---

Not authorized / not claimed: ARCHITECTURE claim-text only (no mechanism change, no design-lock reopen); no domain doc edited by me (those are owner-folded next); no PLAN/IMPL/`pcode`/spike, no Step-1 PLAN, no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-043749.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-architecture` — OK
- `grep -nE 'by construction|sole[- ]writer|forgery-robust' master/ARCHITECTURE.md` — post-sweep: remaining hits = the §1 scoping note, the scoped `submit()`-path sole-writer, the shelved-milestone relabel, R2 grammar (KEEP), §C4.3 (KEEP).
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: swept `master/ARCHITECTURE.md` §1/§2/§3/§4/§5/§C3 (8 claim relabels + a §1 scoping note; R2 grammar + §C4.3 kept); wrote this lane-1 relay + appended `master/relays/INDEX.md`; no mechanism change, no domain-doc edit, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP reviews the ARCHITECTURE sweep + ratifies the semantic checklist; I then dispatch the domain sweep lanes (`c5-claim-sweep-m-1`, `-m-2`, `-light`) with candidate relabels + the decision-fold lanes (③/⑤/④).
