## DESIGN — s2-core design dispatch: the S2 engine-thickening code shape (recovery phase machine · intake-writer · GC/genesis · the generalized owed-item projection)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s2-core-design
PARENT_DISPATCH_ID: s2-core-audit
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: s2-core-audit/AUDIT-planner-20260704-003144.md, s2-core-audit/AUDIT-implementer-20260704-002839.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: DESIGN dispatch — design the S2 implementation against the reconciled audits; implementer is CC'd as design-challenger (formal DESIGN-REVIEW comes TO them per the lineage-gate discipline); Q1/Q2 provisional-pending-guide

**Basis:** the reconciled paired audits (RECONCILE.md entry 2026-07-04; full agreement, zero contradictions) + the locked spec (m-7 §2.1/§2.2/§5/§6/§10/§13; m-1 §5/§6; ARCH §C4) + the s2-dispatch charter. This designs the *code*, not the contracts — locked text is never amended pair-side; spec problems escalate via me.

### Scope

Design the S2 implementation of the four IN items on the S1 baseline (main@cc85049; code surface = s1-close, no diff — implementer-verified). Deliverable: `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md` with a DESIGN_DOC_ID, decision-per-section shape (the s1 design is the house style), acceptance-criteria draft keyed to a fixture namespace you propose, rejected-alternatives log, and the claim-boundary held in every sentence.

### Hard constraints (from the reconciled audits — each is a review blocker if violated)

1. **Promote, don't rebuild.** The audits' already-closed inventories are binding: staging cleanup, RebuildProjections, intake.Unconsumed, binding restore, gate.Complete-before-open, record.Verify, the crashpoint/child-SIGKILL harness, WriteFileAtomic, content-hash dedupe. Each existing piece lands IN its phase/mechanism slot; a re-implementation of any of them is a duplicate-gate violation.
2. **Reified phase machine (G1).** Phases 0–4 as enforced structure (m-7 §5 :89-95 strict order; no authority consumption until open — as structure, not call-order luck), promoting the S1 pieces into their slots; **phase-boundary crashpoints** (new registry names + Hit sites in the recovery path — today there are zero) so G1's crash-at-each-boundary is fixturable; the names-live-in-source pin extends.
3. **Quarantine replaces fail-stop (F-1, convergent finding).** Checksum-mismatch on a committed record ⇒ quarantine + HELD-class incident record (m-7 §5 :91), never fail-stop, never silent-skip — and the disposition must sweep the LIVE-path callers of `store.Records()` (lineage scan, gate.Complete, projections), not just recovery: today one corrupt record bricks both (store.go:114-117 → recover.go/lineage.go/derived.go). "The store never bricks" (§6 :102) is the property under test.
4. **Single intake-writer task (F-2, convergent finding).** m-7 §2.1 :52 verbatim: handlers hand commands over a bounded channel to ONE intake-writer; handlers never touch any file. Today per-connection goroutines call journal.Append directly (main.go:90) with no lock and a ReadAll+len id race. Design the writer task + the ordering fixture under REAL concurrent multi-seat submits across a crash (S1's F9 is single-writer-scripted; both audits concur).
5. **Genesis + GC to the locked posture, claim-pinned.** Genesis = first canonical record {schema_version, config digest, address-space seed, creation stamp} (§10 :136), phase-0 validation fail-closed to read-only diagnostics; state + fixture the idempotence property (re-init against an existing genesis writes nothing new — the duplicate-relay-id rejection is the natural discriminator). GC target set from locked text ONLY: {old rendered projections; drained intake/redo journal segments whose entries all have outcomes} — canonical records are NEVER GC'd (§10 :137); the gate fixture asserts records/ byte-untouched AND no outcome-less entry in any collected segment. Segment rotation designed jointly (size-based, operator-config).
6. **The owed-item projection generalizes derived.go — one mechanism, not two.** The guide advisory is binding design input (s1 design §9.4): gate/held derived-work becomes an INSTANCE of the generalized projection. Semantics from the charter verbatim: `open = owed-record with no disposition-record`, materialize-first. The owed-item `record_kind` + disposition-record shape are a PROPOSAL routed to m-1 fidelity — never fixed pair-side (F2 condition). No ODB/consumer-schema closure (S4-OUT). OI-S1-F11-SWEEP is the first customer: surfaces as open, dispositioned at the S2 exit gate.
7. **Claim honesty (blockers-by-construction for any S2 doc/fixture/string):** "exactly-once EFFECT" never unqualified "exactly-once" (m-7 §2.2 :58 wording verbatim); GC stated as the locked target set (never "keeps live records" — understates it); materialize-first beside every projection claim ("guards recorded owed-items only"); D5 residual beside every exclusivity-shaped claim; enum byte-exact; I-PH holds for every new surface (quarantine incident text, GC output, projection artifacts, phase diagnostics).
8. **Harness extension, not replacement (OI-S1-F11-SWEEP).** The full class×point sweep rides the S1 harness with three designed deltas (planner audit §5): a per-class crashpoint applicability answer (expected-reachable map or not-hit⇒clean-completion leg); recovery-phase crashpoint names; new mutation-class arms (genesis, GC-marker, owed-item, disposition). F9/F11 re-run whole under the new machinery.
9. **Shape-aware performance findings (design considerations, not gold-plating):** F-3 — gate.Complete's three full-store scans run per submit (main.go:98, derived.go); the generalization should make the open-set computed at recovery + maintained incrementally on the loop, making the O(N²) hot path disappear as a *consequence* of constraint 6. F-5 — appendUnique full-file reread sits adjacent to rotation/GC. F-4 — delivery nudge is a broadcast today (main.go:102), not per-recipient; design-awareness under G1's "no lost/double delivery," not a rebuild.

### Provisional-pending-guide (do NOT block on these; mark the sections provisional, as S1 did with Q-A)

- **Q1 genesis-digest scope** — design to our recommended shape (a): digest over the artifacts that exist now; CQ-4b composition lands later as a config-change record. If the guide corrects, the section re-cuts before lock.
- **Q2 OI-S1-F11-SWEEP authorship** — design to shape (i): operator-channel submit during IMPL. Same fold rule.
(Routed `s2-guide-q1/SITREP-orchestrator-planner-20260704-004330.md`, operator-carried.)

### Out of scope (escalate, never absorb)

The ROADMAP OUT list verbatim (S3 registry/linter · MCP live-adapter · observe Step-2 · routing Step-3 · consumer schemas S4) · any m-1 contract amendment (record_kind is a proposal TO m-1) · outbox drain/external send · any ../master or ../extracted edit · code (this is DESIGN — no source edits).

### Lineage + gates (design-review discipline)

You author the design (DESIGN relay carrying DESIGN_DOC_ID, FROM your seat). Your design-review request goes **TO s2-core.implementer** (me/operator CC only); the implementer's DESIGN-REVIEW verdict parents to your DESIGN relay on the same DESIGN_DOC_ID. On approve: design-completion SITREP to me → I reconcile → PROCEED-TO-PLAN (sequencing only; the gated PLAN lock stays in your seat). The m-1 fidelity packet (the audits' store-touch union — 10 surfaces) goes out at PLAN time on the S1 pattern; nothing store-shape-touching dispatches before m-1's approve.

### Acceptance criteria (for this DESIGN)

- Every hard constraint above has a design section that satisfies it or a named rejected-alternative explaining why not.
- Every S2 exit-gate line maps to a proposed fixture id; the fixture namespace covers G1–G6.
- The m-1 proposal surface (record_kind, disposition, genesis, quarantine, GC/rotation layout) is a self-contained section extractable as the fidelity packet.
- Claim-sweep clean (§16 classes) over the doc itself.
- No code, no tracked-file edits outside `docs/sprints/2026-07-03-s2-slice-2/designs/` (+ your relays/INDEX rows).

Operator-judgment items: none new (D5 restated; MCP deferral stands; Q2 carries operator visibility via the guide relay CC).

ACTIONS_GIT_REF: none — design dispatch authored as this relay file + an INDEX row under gitignored .relays/; no tracked-file edit in this action (the reconcile ledger entry rides its own commit, cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree
