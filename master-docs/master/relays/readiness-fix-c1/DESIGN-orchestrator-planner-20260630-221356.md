## DESIGN — bounded MUST-fix (Cluster 1): reconcile the m-1/m-2 `submit()` write-path — CTO-arbitrated seam for co-sign + fold

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: readiness-fix-c1
PARENT_DISPATCH_ID: readiness-fix-c1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded design reconciliation of a locked c1 contract; operator on CC
GRILL_REQUIRED: no — re-lock is VP co-sign; VP may call a grill at fold if warranted
FROM: master.orchestrator-planner
TO: m-1.planner, m-2.planner
CC: master.orchestrator-reviewer, m-1.implementer, m-2.implementer, operator

m-1, m-2 — a **bounded** MUST-before-Step-1 fix from the build-readiness review (`READINESS-REGISTER.md` Cluster 1),
verified line-by-line. This is a **shared c1 contract collision**, so per the charter it is **CTO+VP-arbitrated**: I
propose the reconciled seam below; the VP co-signs (or pushes); you two fold it into your respective docs (or surface a
domain constraint that breaks it). **Scope is exactly this write-path seam — nothing else in c1 reopens.**

**The contradiction (verified):**
- **1a — the write-path ordering is NOT identical, and the lineage gate has no hook-point in m-1's `submit()`.**
  - **m-1** (`…trust-identity/design…:91-94`): "*there is no 'append as submitted' step (ordering reads identically to
    m-2 §4)*" → pre-flights **(a) form-validation (b) m-3 observe** run **before any append**, then **one atomic
    append-as-`accepted`** + deliver. **No `submitted` state; no lineage gate in the ordering.**
  - **m-2** (`…forms-determinism/design…:72-73`): on form+observe pass the courier "**appends it to the store as
    `submitted`/attempted … so the lineage engine can walk the graph including this record**," *then* the **lineage gate**
    runs, *then* it is marked `accepted`; a lineage failure leaves a `submitted`-but-non-deliverable record + bounce.
  - So m-1's "reads identically" is **false**: m-2 has a persisted `submitted` state **and** a lineage gate; m-1 has
    neither. The lineage gate that preserves "**authority blocks before dispatch**" (protocol.md:397/407) has nowhere to
    live in m-1's `submit()`.
- **1b — `submit()` embeds the m-3 observe pre-flight, but m-3 builds at Step-2.** m-1:93 step (2b) and m-2:72 step (4b)
  both put the m-3 observe gate **inside** the Step-1 write-path, yet `ROADMAP.md:57-70` builds m-3 at **Step-2** and
  `ARCHITECTURE.md:120` titles it Step-2. "Store + form-gate without m-3" is **not derivable** from the locked pipeline.

**Proposed CTO-arbitrated seam (for VP co-sign + your fold — push if a domain constraint breaks it):**

> **Canonical Step-1 write-path — one atomic accept, gates as PRE-APPEND validation, observe reserved for Step-2:**
> 1. **resolve + stamp** (m-1) — FROM/ROLE from the binding; system envelope fields; reject-unbound.
> 2. **pre-append validation — nothing persisted yet, candidate held in-courier:**
>    a. **form-validation** (m-2 — required-set, enums, seat-scope);
>    b. **lineage gate** (m-2 — the lineage engine validates the candidate's edges against the **persisted
>       `accepted`-graph**, with the candidate supplied **in-courier**; it does **not** require the candidate to be
>       appended-as-`submitted` first — walking `persisted ∪ {candidate}` is equivalent and needs no limbo state);
>    c. **[RESERVED — lands Step-2] observe-as-send** (m-3) — an **additive** hook-point named in the ordering but **not
>       required for Step-1** (satisfies 1b + the ROADMAP Step-1/Step-2 split; additive, non-destructive).
> 3. **atomic commit** (m-1) — on pass: **one append of an `accepted` record** + INDEX row + deliver, single transaction
>    (TOCTOU closed). On form/lineage fail: **a distinct terminal `rejected` evidenced record** + bounce with the failing
>    field/edge — **not** a persisted `submitted` limbo (preserves m-2's audit-of-attempts as a *terminal* record while
>    honoring m-1's "no un-gated deliverable state").

**Why this resolves both:** m-1 keeps "no persisted `submitted` state · one atomic accept · authority-blocks-before-
dispatch"; m-2 keeps its **lineage gate before delivery** (now as pre-append validation) **and** an auditable attempt
trail (now a terminal `rejected` record); the m-3 observe gate becomes the reserved Step-2 hook, so Step-1 = **store +
form + lineage** only.

**Fold split (each owns its side):**
- **m-1** — rewrite `submit()` :91-94 to the pre-append-validation + atomic-accept ordering; drop "reads identically to
  m-2 §4"; add the lineage-gate step (2b) and the reserved Step-2 observe hook (2c); define the terminal `rejected`
  record on fail.
- **m-2** — rewrite `send()` :72-73 to remove the persisted `submitted` append; state the lineage gate as **pre-append
  validation against the `accepted`-graph** with the candidate in-courier; recast the failure record as terminal
  `rejected`.

**Invariants that must NOT change (preserve):** sole-writer append-only store (I1); channel-stamped FROM; TOCTOU-closed
single accept transaction; "authority blocks before dispatch"; consumers act only on `accepted`.

**Acceptance ("reconciled") =** m-1 and m-2 `submit`/`send` orderings are **byte-consistent** on the seam above; the
lineage gate has an explicit hook-point present in **both** docs; Step-1 requires **no** m-3; VP co-signs the reconciled
seam.

Not authorized by this relay: no PLAN, no code/spike, no other c1 change, no re-scoping beyond this write-path seam.

ACTIONS_GIT_REF: wrote this bounded fix-dispatch relay + appended `master/relays/INDEX.md`; no code/source/pcode edits, no design-doc edits (the fold is the pairs'), no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP co-signs or pushes the arbitrated seam; m-1 + m-2 fold their sides to byte-consistency; then re-verify Cluster 1 closed (MUST-before-Step-1-PLAN).
