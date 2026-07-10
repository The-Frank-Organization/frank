## SITREP — s3 → master: SCOPE question — does S3 build the §7 config-change record? (escalation-shaped; not self-answered)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-scope-q1
PARENT_DISPATCH_ID: s3-dispatch
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s3.orchestrator-planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-7.planner, m-2.planner, operator
SUBJECT: scope question — the m-7 §7 config-change record: the locked S2 design says verbatim "the §7 config-change record (S3)", but the s3-dispatch IN list does not name it; registry evolution on an existing store is impossible without it. Build in S3, or defer with fresh-store posture recorded? Scope is master's call — we do not self-expand

**The question.** m-7's locked design §7 (:109) makes a legitimate config change "itself a committed store record." S2 built config as digest-pinned at genesis for the store's life, and the locked S2 design states verbatim that the change mechanism is deferred: **"the §7 config-change record (S3)"** (s2 design §1 :18). The s3-dispatch IN list (full registry · dissolution+replay · schema_version/migrators · R2/GRILL carries · re-render/drift) does **not** name the §7 config-change record. The S3 registry is a pinned config member, so this bites S3 directly: a changed registry member changes the top-level config digest, and an existing store then fails phase-0 genesis validation (internal/recover/recover.go:31-39) — **registry evolution on an existing store is impossible until the §7 record exists.**

**Why we are not self-answering.** Building it = engine-owned mutation mechanism (a new commit-loop mutation class + recovery interaction) riding S3 on an inference from S2's forward-pointer — a scope expansion if master didn't intend it. Not building it = the S3 re-render/drift gate line is only testable as restart-with-new-store, and the "registry live end-to-end" claim carries a fresh-store-only qualifier — honest, but master should ratify that posture explicitly rather than inherit it silently. Both audits (paired, reconciled — ledger `docs/sprints/2026-07-04-s3-slice-3/RECONCILE.md`) flagged this as the one escalation-shaped item; zero other scope questions surfaced.

**The two dispositions we see (recommendation stated, decision master's):**
- **(a) DEFER (our recommendation):** S3 posture = fresh `store.Init` carries the S3 registry; the §7 config-change record is recorded as a **typed owed item through the live S2 owed mechanism** (materialize-first with teeth, the OI-S1-F11-SWEEP precedent), owner = the wire-up/next slice at master's discretion; the S3 re-render/drift fixtures simulate registry change as restart-with-new-store; every S3 claim-surface states the fresh-store qualifier plainly. Rationale: keeps S3 = m-2's form system exactly (the s3-dispatch boundary), no engine-mechanism scope rides in on inference; the record lands where m-7 guides it directly.
- **(b) BUILD IN S3:** master amends the S3 IN list explicitly; m-7 becomes the mechanism guide for that item (it is engine-owned per the locked one-line boundary — m-2 guides the form system, not commit-loop mutation classes); m-1 fidelity on the new record_kind; the S2 crash-harness applicability map gains the new mutation class.

**Interaction note for the ruling:** m-7 consult Q2 (`s3-consult-m7`, m-7.planner TO'd, same batch) asks the version/digest-axes question whose answer partly depends on this ruling; the S3 DESIGN proceeds provisionally under (a)-shaped assumptions and de-provisions on your ruling — nothing locks before it.

Next requested action: master's scope ruling via operator hand-relay (reply lands in `.relays/s3/s3-scope-q1/` per the s1/s2 pattern). If (b), the amended IN line + guide split; if (a), ratification of the fresh-store posture + the owed-item owner.

ACTIONS_GIT_REF: none — question relay only; this file + an INDEX row under gitignored .relays/; no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at b800201, the reconciliation ledger commit of record)
