## DESIGN — step3-relock-dag-m3: the schema-version amendment is RATIFIED — your parked lane-2 r1 is UNBLOCKED; fold the cut-matrix + verdict machines + §B sink + the `logical_surface_digest` recipe-binding against the settled m-8 r5 `c0b7b488…` + m-9 r5 `c0ff74f5…`; `model_surface_digest`/E-join stays deferred to the future v3 delta

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m3
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m3
IN_REPLY_TO: master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-234500.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, operator, master.orchestrator-reviewer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: fold your parked lane-2 r1 now — schema v2 ratified (`9e874df8…` + contract `6e2abe40…`), settled producer bases handed you; the model_surface_digest E-join is NOT in this fold (deferred to v3)

m-3 pair — the schema-version amendment is **operator-RATIFIED** (`step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-234500.md`; amendment rev3 `9e874df8…` + your bound contract `6e2abe40…`). Your parked lane-2 r1 is now **unblocked** on real bytes. Run your normal pair cycle (planner authors, implementer adversarially reviews).

### The settled bases handed to you
- **m-8 producer r5 `c0b7b488…`** — pair-approved (m-8 review `…-143000`). This is the **cut basis** for your exhaustive cut-matrix: the deny/post-freeze-reject carrier matrix (which m-8 disposition carries B/E on which versioned carrier, which cuts are message-absent) is now settled bytes, not assumed. Author your cut-matrix against it.
- **m-9 producer r5 `c0ff74f5…`** — pair-approved. This is the settled source for the **`logical_surface_digest` recipe-binding** (the E0 carriage field is ratified at schema grain in your v2 contract; its recipe-binding confirmation was parked in D3 and now unparks against m-9's settled producer).

### What you fold in r1 (against the ratified v2 schema + the settled bases)
- The **exhaustive cut-matrix** — one row per settled m-8 r5 disposition, each naming freeze-reached / the exact carrier field / E0 required-or-forbidden / E3 required-or-forbidden / the predicate-1 result where no wire exists (the parametric per-cut requiredness rule from your ratified contract, now instantiated over the settled cut-LIST).
- The **five verdict machines** (ordered, mutually-exclusive, total over every schema-valid input incl. missing/contradictory; governed-read domain `committed|not_found|unavailable`).
- The **§B sink record** — the end-to-end `frozen_core_digest` consistency + independent observer-derivability across m-8 r5 / m-9 r5 / m-10 rev6 carriage (F73 confirmations + your m-3 sink record; §B is NOT a two-sided join).
- The **`logical_surface_digest` recipe-binding** confirmation against m-9 r5.

### Explicitly NOT in this fold (deferred by the ratified amendment)
- **`model_surface_digest` / the E-join** — the ratified amendment puts it OUT of v2, deferred to a **later governed v3 E-join delta**. Do NOT fold it here. When the v3 delta is commissioned it is its own governed cycle.

### Boundaries
DESIGN-only; governed additive delta over frozen r4 `009df607…` (v1 byte-frozen) realizing the ratified v2 contract `6e2abe40…`. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. H-12 hard-blocks external use. Escalate spec mistakes / cross-owner conflicts UP through master (the escalation discipline that served the §D settlement well).

### Where this sits
Your r1 fold is the last m-3 piece of §11 lane 2. Once it + the other pairs' folds + the §D join co-sign land, master integrates → item A (the extraction recipe + `bundle_sha256`) → lane 4 (the shorter re-lock) → lane 5 (T4).

## Verification
Reproduced: amendment rev3 `9e874df8…` + bound contract `6e2abe40…` ratified/UNMOVED; frozen r4 `009df607…` UNMOVED (v1 frozen); m-8 r5 `c0b7b488…` + m-9 r5 `c0ff74f5…` approved. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-3 pair folds + pair-reviews its lane-2 r1 (cut-matrix + verdict machines + §B sink + logical_surface_digest binding) against the settled m-8 r5 + m-9 r5, and returns the byte-bound design + F73 confirmations + a SITREP; master integrates toward item A + the re-lock.
