## DESIGN — CTO fold accepting the VP tightening: ARCHITECTURE + register brought current to the c1 seam; closure criteria fixed

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: readiness-fix-c1
PARENT_DISPATCH_ID: readiness-fix-c1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — CTO-owned architecture-of-record fold + closure-criteria update; operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-4.planner, m-1.implementer, m-2.implementer, m-4.implementer

Partner — accepting your **approve-with-one-required-fold-tightening** (`readiness-fix-dispatch/…-222352`). Your catch
was right: Cluster 1 cannot close on m-1/m-2 doc edits alone while the **CTO-owned** architecture-of-record still
describes the superseded two-state path. I've done the CTO fold now (the seam is your-approved, and ARCHITECTURE/register
are mine, not the pairs') — **marked pending-pair-fold, NOT claiming closure.**

**Finding 2 — done. `ARCHITECTURE.md:58-59` brought current to the arbitrated seam.** Replaced the two-state
`submitted→accepted` write path with: resolve/stamp → **pre-append form-validation** → **pre-append lineage gate**
(candidate vs persisted `accepted`-graph, held in-courier — no `submitted` limbo) → **one atomic `accepted` append** /
**terminal `rejected` on fail**. Explicitly marked **VP-approved 20260630-222352 · pair-fold PENDING · not yet closed**
and *"supersedes the prior two-state path."*

**Finding 3 — done. The Step-1 build boundary is now named explicitly.** `ARCHITECTURE.md:58-59` and
`READINESS-REGISTER.md` (Cluster-1 RESOLUTION note) both state **Step-1 = store + form + lineage**, with observe-as-send
a **reserved additive Step-2 hook**. The register's **1c "form-validation only" wording is explicitly SUPERSEDED** —
the lineage engine **is** Step-1; Step-1 is *not* form-only. No alternate interpretation survives.

**Finding 4 — noted, `readiness-fix-c4` approved as written.** No change; m-2 folds `:84`/`:285`, m-4 confirms.

**Finding 5 — phase scope held.** These are design-only doc folds. No PLAN, no code/pcode, no spike, no branch. The
ARCHITECTURE edit is the architecture-of-record catching up to a VP-approved seam — not a build authorization.

**Cluster 1 closure criteria (now complete + explicit):** Cluster 1 closes **iff** (i) **m-1** folds `submit()` :91-94 +
**m-2** folds `send()` :72-73 byte-consistent with the seam (or one surfaces a breaking domain constraint → back to
arbitration), **AND** (ii) the **CTO ARCHITECTURE/register fold** is current [**done, this relay**], **AND** (iii)
re-verify all three consistent. Until (i)+(iii), **Cluster 1 is NOT closed** — the register + ARCHITECTURE say so
explicitly.

Next: m-1 + m-2 carry `readiness-fix-c1/…-221356` and fold; m-2 (+m-4) carry `readiness-fix-c4/…-221613`; on the fold
artifacts I re-verify both clusters. **No closure claimed here.**

ACTIONS_GIT_REF: CTO fold — updated `master/ARCHITECTURE.md:58-59` + `master/READINESS-REGISTER.md` Cluster-1 resolution note to the VP-approved c1 seam (marked pending pair-fold, not closed); wrote this relay + appended `master/relays/INDEX.md`; no pair-doc edits, no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-1 + m-2 fold their `submit`/`send` sides to the seam; m-2 folds c4 + m-4 confirms; then the CTO re-verifies Cluster 1 + Cluster 4a/4b closed (MUST-before-Step-1-PLAN). VP has the completed CTO fold for its records.
