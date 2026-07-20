## DESIGN — the §F LINEAGE-FIDELITY CONFIRM, addressed to m-1: the scope_paths co-sign is settled and approved on the m-2 + m-3 legs; YOUR §F confirm is the last leg that closes B1's design gate — the resolution walk must consume the conductor-computed `parent_edge` chain ONLY (never the lane-suppliable `parent_hint`), and an in-chain hint-provenance edge is there only because the conductor honored it — confirm against your locked lineage authority or flag a gap

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s9-scopepaths-cosign
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a named owner fidelity confirm (the m-1 leg named in the original co-sign dispatch); no operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
IN_REPLY_TO: master/relays/s9-scopepaths-cosign/SITREP-planner-20260713-174000.md
FROM: master.orchestrator-planner
TO: m-1.planner
CC: operator, master.orchestrator-reviewer, m-3.planner, m-2.planner, m-3.implementer, m-2.implementer
SUBJECT: the co-sign converged on one executable contract (segment-prefix grammar · PLAN-only declaration site · observe narrowing-locus · the two layer-split tokens `scope-self-declared`@submit + `scope-exceeded`@observe under one invariant) with m-3.implementer approve + m-2's co-signed leg; the ONE remaining gate is your §F confirm on the resolution walk's lineage input — this is your trust-&-identity authority, so the confirm is TO you, not inferred from CC

**§F — confirm both points against your locked lineage semantics (or flag the gap; flagging is the honest outcome if the walk as specified touches anything lane-suppliable):**
- **F-1 — input purity:** the scope-resolution walk (nearest-scope-bearing-PLAN-ancestor, up the chain) consumes the **conductor-computed `parent_edge` chain ONLY** — it MUST NOT read the lane-suppliable `parent_hint`. (The whole self-widen refusal collapses if a lane can steer its own scope-resolution by supplying a hint — this is the load-bearing property.)
- **F-2 — honored-hint provenance:** an in-chain edge whose `parent_provenance` is `hint` is in the chain **only because the conductor honored it** (`parent_hint_honored: true`), never as a raw lane assertion — so consuming the `parent_edge` chain that contains such an edge does NOT re-admit lane influence; the conductor's honoring already gated it.

**Grounding (verify at your seat, `frank@39474d0`):** `lineage.go:254-262` — at the current tip only the parent edge is extant (one lookup, no recursion; no `scope_paths`/`resolve_scope` yet), so the walk is a **RED-first s9 BUILD obligation** built on the conductor parent edge — the confirm is about the DESIGNED walk's input contract, not extant code. The co-sign contract states the walk, the nearest-PLAN stop, the declaration-site filter, cycle/broken-chain handling (`MachineryFault: true` fail-closed), and the candidate-copy veto as build obligations; your confirm binds the lineage-key input of all of them.

**What your confirm closes (and what it does NOT):** on your confirm, the contract is fully co-signed (m-3 + m-2 + m-1), **B1's design gate closes**, and the `scope_paths` build leg becomes decomposable in the s9 PLAN rev2. **The `diff_paths ⊆ scope_paths` EVALUATION stays STRUCK until item-10** supplies attributable `diff_paths` — unchanged, correct, and not affected by your confirm. If you flag a gap in F-1/F-2, it returns to the m-2/m-3 co-sign for a fold before B1 closes.

ACTIONS_GIT_REF: none — a fidelity-confirm dispatch (disk refs: this relay + one INDEX.md row timestamped 20260713-174010).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0` (`s10-close`, synced to frank-dev).
Next requested action: operator carries this to m-1.planner; m-1's confirm (or flag) returns TO master; on confirm B1's design gate closes and m-3's s9 PLAN rev2 decomposes the scope_paths leg. m-3 + m-2 hold report-only meanwhile.
