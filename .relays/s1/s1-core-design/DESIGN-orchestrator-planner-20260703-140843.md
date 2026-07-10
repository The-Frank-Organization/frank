## Team s1-core — Slice-1 thin end-to-end conductor relay: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only for Planner; read-only challenge/answers for Implementer
DISPATCH_ID: s1-core-design
PARENT_DISPATCH_ID: s1-core-audit
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for the design phase itself; the downstream S1 PLAN is m-7-guide + master-VP gated, m-1/m-2 fidelity-reviewed, and merge is a separate human gate
GRILL_REQUIRED: no
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
DESIGN_DOC_ID: s1-slice-1-design
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@225788c
TARGET_BRANCH: main

Phase scope — DESIGN.
Current scope: Planner uses Superpowers brainstorming to design the S1 *implementation* (the code shape — the locked contracts are not designable); Implementer answers design questions, challenges alternatives with evidence, and flags anything needing operator/guide judgment.
Not in current scope: source/test edits, implementation branches, commits (EXCEPT the design doc itself under docs/sprints/2026-07-03-s1-slice-1/designs/ — a tracked sprint doc, committed on main per the sprint substrate), PRs, scaffolding, or prototype implementation.
Implementation begins only after a current relay under the active run's RELAY_ROOT contains the exact literal token `DISPATCH IMPL` bare, unfenced, un-backticked, alone on its own line, and addressed to the Implementer in `TO`. For this run there is an ADDITIONAL hard gate above pair delegation: the m-7-guide + master-VP plan gate and the m-1/m-2 fidelity reviews must pass first.

Reconciled audit summary (full ledger: docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md, entry 2):
- Primary bucket: still-open — greenfield confirmed E2 by both seats; no contradictions between the paired audits.
- Already-built/overlap: the upstream relay-lint = DO-NOT-COPY (the m-2 §10 dissolve map deletes it); the upstream store layout = REUSE-AS-SPEC'D; the 244-fixture corpus (checker-verified all-PASS, E2) = read-only replay input; jcode/claude-code = prior-art-only.
- Evidence: exit-gate map complete on both sides, no spec gaps; frozen m-1/m-2 surfaces enumerated E1 (the fidelity-review object); the three chartered owed carries typed.
- Pending external input: `s1-guide-q1` to m-7.planner (Q-A recovery line; ⑤ ODB-egress classification). Design the affected sections PROVISIONAL — clearly marked, resolvable by the guide's answer without re-architecting.

Design questions to resolve (from the reconciled audits; ids from the planner audit §6):
1. Q-B — conductor runtime/language + MCP server stack + per-seat channel/adapter boundary. Locked topology is fixed (one process, single intake-writer, one commit thread, concurrent lock-free reads, per-seat stdio/bearer channels, pipe-write wake — m-7 §1/§2/§8); "riding existing runtimes" constrains seats, not the conductor.
2. Q-C — the concrete MVP FieldSpec set. Floor from the audits: the m-1 :138 system envelope + TO/CC + PARENT; header enums sufficient for seat-scope + the A2 forbidden-option fixture; one gate pair (`HUMAN_GATE_REQUIRED` monotonic + `gate_category` at the grain the ③ carry needs); `delivery_state` + `failing_edge`; the `X-` overflow namespace. Selection criterion: maximize covered m-2 §10a/§10b classes per unit of build (fixes the R1 replay subset).
3. Q-D — deterministic crash-fixture mechanics: injection points (post-intake-fsync, pre/post record-fsync, pre/post rename, pre/post dir-fsync, projection write, delivery write, replayed intake-id) + mechanism (internal test-only crash hooks keyed by config absent from seat surfaces, vs external syscall-boundary sweep). I-PH applies to fixture outputs too.
4. Q-E — S1 minima for mint/connect (attach-time mint, credential re-present on re-attach, persisted binding table), park/wake (W1 semantics), the operator-channel realization (must exist for W1/B4), and which seat runtime drives the E2 fixtures (a scripted test seat suffices; lane-qualification probes are NOT S1 gates).
5. Q-A — recovery line: design to the pair's shared reading (intake journal + §4 pivot + dumb-replay recovery: staging cleanup, projection rebuild, intake−outcomes re-enqueue in arrival order, binding-table restore, wake re-issue; NO genesis/quarantine/GC/segment-rotation/phase-0-4 machinery) — PROVISIONAL until the guide confirms via s1-guide-q1.

Required design alternatives (at minimum):
- Q-B: two candidate stacks compared (implementation language + MCP server library + process/IPC shape), tradeoff criteria: fixture-drivability (kill -9 determinism), fsync/rename control fidelity, dependency surface, team familiarity.
- Q-D: internal crash-point hooks vs external injection — tradeoff criteria: determinism, I-PH safety, S2 reusability.

Hard constraints the design must carry (the guide-gate checklist is the PLAN's rubric — design to it now; ROADMAP.md §guide-gate):
- Pivot shape from line one: canonical-record `rename()` = the single commit pivot; fsync-before-rename, dir-fsync-after; presence=committed; projections derived; outcome records reference `intake_id`; F11 one-pivot-per-mutation.
- Byte-exact `{accepted, rejected, held}`; `bounced`/`submitted` never appear as value tokens.
- The interface guardrail is a code-layer property: per-seat tool registry exactly `{submit, project, read}`; raw store/config/outbox/operator-channel paths + config values absent from every seat surface.
- The three chartered owed carries land materialize-first (typed records already in both audits — carry them into the design).
- Claim honesty per m-7 §16 sweep in everything seat- or user-facing (code comments, tool descriptions, docs): D5 residual beside every exclusivity-shaped claim; only the serialized-loop kill + constrained-grammar R2 presented as operationally live.
- Adopt the planner audit's fixture-id namespace (B1-B4, A1-A4, C1-C6, R1, P1, L1, W1) so the plan's fixture names trace to the exit gate 1:1.

Boundary contract to design around (locked, from the charter):
- Writes: append-only relay-store records (terminal byte-exact {accepted,rejected,held}) + a local outbox/ODB item
- Reads: the locked m-1 store API + m-2 FieldSpec envelope
- Target entity: the conductor's committed relay + its rebuilt projection
- Downstream consumer: seat inboxes (via project) + the local outbox
- Contract: submit/project/read; channel-stamped FROM; MVP FieldSpec; crash-atomic commit
- Proof: the S1-scoped hardened exit gate (E2 fixtures)

Out of scope:
- Everything on the ROADMAP scope-OUT list (S2/S3/S4/Step-2/Step-3); m-2 §9 migrator machinery (schema_version stamping only — reconciled disposition); organic master-trail failures as R1 gate inputs (fixture matrix is the corpus of record); re-designing any frozen contract; editing anything under ../master or ../extracted.
- If the design seems to need an OUT item: escalate to s1.orchestrator-planner; do not absorb.

Deliverable:
Design lock via Superpowers brainstorming recorded as DESIGN_DOC_ID `s1-slice-1-design` at docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md — selected options, rejected alternatives, boundary contract, acceptance-criteria draft keyed to the fixture-id namespace, provisional sections flagged, open questions. Design Q&A with your Implementer may stay inline; the lock is file-first. After writing the doc, send a Template-I DESIGN-REVIEW request `TO: s1-core.implementer` (a design TO the orchestrator with the Implementer on CC is NOT a review request). On `DESIGN_REVIEW_VERDICT: approve`, send the design-completion SITREP to s1.orchestrator-planner for PROCEED-TO-PLAN.

ACTIONS_GIT_REF: wrote this relay + an INDEX.md row under .relays/s1/ (gitignored substrate); audit-reconciliation entry committed on main@225788c
FINAL_GIT_STATUS_SHORT: none — clean tree
