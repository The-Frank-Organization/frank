## SITREP — de-provision supplement: guide CONFIRMED Q1=(a) and Q2=(i); eight sharpenings + the recovery-reads-only-the-store constraint join the DESIGN hard-constraint set

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-core-design
PARENT_DISPATCH_ID: s2-core-audit
RUN_ID: s2
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: ../.relays/s2/s2-guide-q1/SITREP-planner-20260704-004750.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: Both provisional design sections de-provisioned on the guide's relay — design to the confirmed shapes; the grill folds Q1/Q2 as resolved-by-guide rows (do not re-ask the operator); sharpenings below are review blockers

The m-7 guide answered s2-guide-q1 entirely from locked text (no lock amendment): `../.relays/s2/s2-guide-q1/SITREP-planner-20260704-004750.md`. Both r2-provisional sections are now **confirmed** — de-provision them and design to these shapes. Per dispatch r2 grill-rule 3, Q1/Q2 enter the GRILL_LOCK as **resolved-by-guide rows**; the grill's remaining operator agenda = the m-1 proposal boundaries + the on-disk layout commitments + whatever your design tree surfaces.

**Q1 CONFIRMED = (a):** genesis pins ONE top-level digest over the config artifacts that exist at S2; the CQ-4b composed shape arrives later via §7's config-change record (a committed store record carrying the new digest). Sharpenings (each a design/review blocker):
1. **Honest attribution:** the pinned artifact is operator-ratified build config *in the m-2-locked shape* — NO domain-author stamp on any section a domain didn't author (no "m-2-attributed" fabrication).
2. **Phase-0 disposition, not error-exit:** mismatch/missing genesis ⇒ serve read-only diagnostics, accept nothing, summon operator (m-7 :90). A non-zero exit fails the locked line; quarantine-never-brick (:91) is the neighbor rule.
3. **Deterministic digest input:** canonical serialization of the pinned artifact set (stable ordering, named members) so the S3 change-record extends the member set unambiguously.
4. **Claim scope:** load-once / digest-verified / Phase-0-fail-closed claimed for the pinned artifacts ONLY; state that CQ-4b composition + per-section stamps land with the consumer sections (S3/S4). SWEEP the wording.

**Q2 CONFIRMED = (i):** the operator authors the OI-S1-F11-SWEEP owed-record through the operator channel. The guide's decisive engine-side reason is itself a design constraint: **recovery may read ONLY the store** — no docs-file read inside the trusted recovery/genesis path, ever; the engine never mints obligation records from out-of-store text (dumb-replay idempotence + store-is-truth). Sharpenings (blockers):
1. **Channel + stamp:** `FROM=operator` on the operator-relay channel (locked m-1 §6 first-class-operator model) — not a synthetic system stamp.
2. **Payload = the typed record:** `{owner, source, target surface, disposition path}`; `source` cites the s1 ledger entry (s1 RECONCILE.md :160-161) AND the guide's deviation-1 ruling (`../.relays/s1/s1-exit-gate/SITREP-planner-20260703-200827.md`); `target surface` = the F11 full class×point sweep on the existing harness; `disposition path` = the S2 exit gate.
3. **Projection semantics verbatim:** `open = owed-record with no disposition-record`; silent drop made impossible for a RECORDED item only — materializing is an intake/triage act by a record-authoring principal. Fixture proof: surfaces as open → closes at the S2 exit gate with a disposition record.
4. **Sequencing:** the operator submit happens during S2 IMPL once the operator channel is up; the projection must NOT gate on the record pre-existing genesis — an empty owed-projection at genesis is correct, not a bug.

Nothing else changes: dispatch r2 (`DESIGN-orchestrator-planner-20260704-005310.md`) remains the dispatch of record — nine hard constraints, grill requirement + fence, OUT list, the lineage-gate discipline all stand. This supplement only converts the two provisional sections to confirmed + adds the sharpenings above to the hard-constraint set.

ACTIONS_GIT_REF: none — report-only supplement; this relay file + an INDEX row under gitignored .relays/ (the ledger entry rides its own commit, cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree
