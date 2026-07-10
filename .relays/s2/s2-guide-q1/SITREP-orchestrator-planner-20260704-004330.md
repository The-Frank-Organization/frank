## SITREP — two audit-surfaced questions for the m-7 guide (Q1 genesis-digest scope · Q2 OI-S1-F11-SWEEP owed-record authorship)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-guide-q1
PARENT_DISPATCH_ID: s2-core-audit
RUN_ID: s2
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s2.orchestrator-planner
TO: m-7.planner
CC: operator, s2.orchestrator-reviewer
SUBJECT: Guide questions from the reconciled s2-core audits — genesis config-digest scope at S2, and who authors the OI-S1-F11-SWEEP owed-record materialization; answers fold into the S2 DESIGN (proceeding provisional meanwhile)

Both s2-core audits are in and reconciled (full agreement, zero contradictions; ledger entry in docs/sprints/2026-07-03-s2-slice-2/RECONCILE.md). Two questions need the guide's read of the locked text; DESIGN is dispatched and proceeds with both marked provisional-pending-guide, exactly as S1 handled Q-A.

**Q1 — genesis config-digest scope at S2 (joint guide + m-1 flavor; answer from locked text).**
m-7 §10 :136 has genesis carry "the config digest," and §7 :109 defines that digest over the CQ-4b composed artifact (per-domain sections, m-2/m-3/m-4/m-5/m-6-authored, single top-level digest). Those consumer sections are S3/S4/OUT for this slice; the only config artifact that exists at S2 is `internal/fieldspec/registry.json` (S1 explicitly deferred digest-pinning to genesis: s1 design D-5 :73). What may S2's genesis pin without building OUT machinery:
(a) a single digest over the artifacts that exist now (registry.json, + any S2-added trusted config), with the CQ-4b composition landing later as a config-change record (§7's "a legitimate config change is itself a committed store record" gives the upgrade path); or
(b) a composed-sections stub (one section, m-2-attributed) so the CQ-4b shape exists from genesis?
Our read: (a) — it pins what exists, claims nothing about sections that don't, and the config-change-record path makes it additive; (b) fabricates a composition whose authors haven't authored anything. Confirm or correct; m-1 gets the layout half in the fidelity packet.

**Q2 — who authors the OI-S1-F11-SWEEP owed-record materialization (operator visibility requested).**
Materialize-first means the projection surfaces only committed records; the owed item currently exists as ledger text (s1 RECONCILE.md :160-161, the typed record you conditioned deviation-1 on). Someone must commit it as the first owed-record. Candidate shapes: (i) an operator-channel submit during S2 IMPL (the operator authors the obligation record; cleanest authority story); (ii) a system-stamped commit at the S2 store's first recovery/genesis (the engine transcribes the ledger; no human in the loop). Our read: (i) — an owed item is an obligation somebody holds; the operator (or an orchestrator seat) authoring it through the governed interface is the honest provenance, and (ii) would have the engine minting obligations from out-of-store text. Confirm or correct; the fixture then asserts it surfaces as open and closes at the S2 exit gate.

Context for both: the reconciled audits (`.relays/s2/s2-core-audit/AUDIT-planner-20260704-003144.md` §2.3/§6, `AUDIT-implementer-20260704-002839.md`) — carried by the operator alongside this relay if useful.

Next requested action: operator hand-carries this to m-7.planner; guide answers from locked text (escalating to master only if either answer would amend a lock); answers fold into the S2 DESIGN via s2.orchestrator-planner.

ACTIONS_GIT_REF: none — report-only question relay; this file + an INDEX row under gitignored .relays/ only.
FINAL_GIT_STATUS_SHORT: none — clean tree
