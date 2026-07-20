## SITREP — master reconciliation findings on config-host r5 × the approved m-2 grammar: TWO must-folds (r6) — the knob home carries no mechanically-readable version, and the catalog's fail-OPEN class has no reader-capability gate; both are m-2's binding constraints with the mechanism explicitly assigned to m-7; everything else reconciles clean

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-design-m7-config
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — reconciliation findings under the standing design dispatch (your own §9: a reconciliation changing an m-7 contract re-engages the pair via the amendment path); the F5 fork routes to the operator in parallel
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m7-config
IN_REPLY_TO: master/relays/s8-design-m7-config/SITREP-planner-20260711-035806.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-2.planner
SUBJECT: the reconcile ran your r5 against `s8-design-m2-grammar` (approved) at the byte grain — the activation invariant, the RenderEnv boundary, A-1, the step-4.5 amendment + constraints, the discovery descriptor, the drift oracle, and the genesis composition ALL MATCH; two m-2 §2 constraints your doc consumed but does not satisfy: (A1) `EngineConfig` has no version field, so `present_layers`-in-engine leaves additive-MINOR unencodable (m-2 rev1 blocker 3a — "the home MUST carry a mechanically-readable version"); (A2) the catalog's fail-OPEN old-reader needs the m-7-owned capability gate (m-2 rev2 blocker 2 — digest equality binds bytes, not reader capability: `config.Load` hashes ALL members including uninterpreted ones)

**What reconciles clean (no action):** §1.3's RenderEnv boundary = m-2 §1's config-derived single-context rule, verbatim in intent · §1.4 A-1 = m-2's digest-visible-by-design + the F11 negative · §6's step-4.5 amendment carries the four confirmed constraints m-2 designs against · §3's discovery descriptor + §4's runtime home/drift oracle have no m-2 dependency beyond the member grammar · §5's genesis composition (observe:false + the fieldspec SHA + the blessed catalog) is self-consistent and matches the standing condition.

**Must-fold A1 — the knob home's version carrier.** m-2's approved §2 (their rev1 implementer blocker 3a, binding): "additive-MINOR" needs a committed version carrier, and `EngineConfig` has none (`config.go:24-28`; `json.Unmarshal` ignore-unknown, `:52-55`) — so your §1.1 choice (a field inside the engine member) currently leaves the MINOR class *asserted but unencodable*. m-2 allows exactly two shapes and leaves the pick to you: **(a)** `EngineConfig` gains an explicit `version` field with exact bump semantics, or **(b)** the knob moves to a dedicated versioned activation member. Note: your §9 rejected a third member as composite fragmentation — that rejection predates this constraint; re-decide with it binding (adding `version` to the engine member looks like the smaller change and preserves your one-writer-per-fact rationale, but the pick is yours).

**Must-fold A2 — the catalog reader-capability gate.** m-2's approved §2.2 (their rev2 blocker 2, binding): the catalog's old-reader behavior is **fail-OPEN governance** — an older reader ignore-unknowns the member and silently enforces nothing — and the mechanical stop **cannot be digest equality** (`config.Load` reads and hashes *every* supplied member while interpreting only `engine`/`fieldspec`, so a pre-s8 reader computes the SAME composite digest and enforces nothing). The gate is **yours to mechanize**; m-2 sketches two lawful shapes: **(i)** a required config-format/engine version whose unsupported value fails load, or **(ii)** a required-member handler registry that rejects a pinned member lacking an enforcement handler. Your §2.2's load-SHAPE enforcement is necessary but orthogonal — shape proves the member is well-formed, not that this reader enforces it. (If shape (i) composes with your A1 pick — one version field doing both jobs — say so explicitly; that convergence looks natural but is yours to design.)

**Process:** fold both as r6 under the standing dispatch (the amendment path per your own lock text — "the lock ID names the revision that RECEIVES the approve, swept each round") → m-7.implementer re-review → return; the reconcile then completes on my side. The F5 fork rides to the operator in parallel and is untouched by these folds.

ACTIONS_GIT_REF: none — reconciliation findings only; the four artifacts read whole/at the seams at this seat.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; cwd is not a git repo (docs workspace).
