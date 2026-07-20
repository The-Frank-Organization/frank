## Team m-1 — Trust & Identity: PROCEED TO DESIGN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c1-design-m-1
PARENT_DISPATCH_ID: c1-reconcile
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — operator owns Step-1 transport strength (grill it); design surfaces operator-judgment items
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-1-trust-identity
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)

Phase scope — DESIGN. Planner leads via Superpowers brainstorming + the design-grill step; Implementer answers and challenges with evidence and flags product-semantics decisions. Not in scope: source/test edits, branches, commits, PRs, scaffolding, prototype code. This is the v3 research+design phase; design-lock is the terminal — no implementation.

Basis: your reconciled c1-audit-m-1 (see master/RECONCILE.md) is APPROVED by the VP (c1-reconcile verdict: approve). Design the primitive the audit recommended — but per the VP's load-bearing guardrail, treat the converged audit points as HYPOTHESES to PROVE in design, not as proven facts.

Co-foundational with m-2 (hard boundary): the m-1 store/stamper FILLS the system fields m-2's schema DECLARES (FROM, ROLE, PARENT, timestamps, evidence); system-filled PARENT is exactly where the lineage engine strengthens. Neither domain design-locks in isolation — coordinate the shared system-field contract with m-2.planner; the orchestrator runs the joint lock after consumer review.

Design questions to resolve (grill the operator on the operator-owned ones):
1. OPERATOR — Step-1 transport strength: operator-attested manual relay, vs minted-token-over-isolated-per-seat-connection, vs OS-peer / mTLS / SPIFFE. This choice DECIDES the strength claim (see proof requirement). VP guidance: to claim "by construction," choose minted per-seat credentials inaccessible to other lanes; if operator-attested, name the result "operator-attested / confusion-resistant," NOT by-construction.
2. OPERATOR — Merkle / tamper-evidence deferral: confirm we accept "trust the courier as the TCB" for v3.0 and defer Merkle/CT to a later cycle.
3. identity vs authority boundary: confirm m-1 owns identity + store only; m-4/m-5 own what a stamped seat may do.
4. attest(): confirm folding attestation into mint_seat / connection-setup (no public attest()); submit() fails on an unbound connection.

Required design alternatives (lay out, with tradeoffs + a recommendation): for Step-1 transport, at least Option A (minted-token-over-isolated-connection — strength claim "by construction") vs Option B (operator-attested — simpler, weaker claim). Tradeoff axis: licensed strength claim vs Step-1 simplicity.

Hard proof requirements (VP carry-forwards):
- PROVE I1 (sole-writer / store-isolation) and I2 (channel-isolation) for the SELECTED Step-1 transport, and STATE the exact strength claim that transport licenses ("forgery-robust by construction" only if I2 holds against a malicious lane that can reach the conductor directly).
- Specify the submit / project / read / mint_seat API surface, the on-disk shape (a projection of a conductor-stamped record, reusing the v2.8.8 layout + claude-code projection pattern), and the isolation boundary — including the Step-1→standalone non-re-cut path (same API; swap only the attestation backend).

Boundary contract — name the consumer fields BEFORE lock (consumer review by m-3/m-4/m-6 is a lock prerequisite):
- m-2: the system-filled field set + the sole-writer store the form-validation/lineage engine reads.
- m-3: the store-isolation boundary (probe-from-outside-the-lane) + the submit() pre-send gate hook point.
- m-4: the stamped store as the write target for routing records (record FROM = the router seat's stamped identity).
- m-6: the addressing graph (TO/CC) + seat-address space (inbox = projection of the addressing graph).

Out of scope: m-2 schema internals, the observe / routing / human-surface mechanisms, any code.

Relay hygiene (VP note): keep the pair-thread DISPATCH_ID c1-design-m-1; address the design-review request TO m-1.implementer (Template I), not TO the orchestrator.

Deliverable: a design doc (Superpowers brainstorming + design-grill), recorded as DESIGN_DOC_ID c1-design-m-1-trust-identity, containing — the selected transport with proven I1/I2 and the licensed strength claim; the submit/project/read/mint_seat interface sketch; the on-disk + isolation model and non-re-cut path; rejected alternatives; operator decisions/defaults from the grill folded into a GRILL_LOCK; the named consumer boundary contract; open questions. Then send the design-review request TO m-1.implementer (Template I), and report design-complete to the orchestrator for the joint co-foundational lock (after m-3/m-4/m-6 consumer review). Design Q&A inline; the design lock is file-first.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
