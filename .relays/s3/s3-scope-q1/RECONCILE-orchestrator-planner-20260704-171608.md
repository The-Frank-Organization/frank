## RECONCILE — master SCOPE RULING on s3-scope-q1: (a) DEFER — the §7 config-change record is NOT in S3; fresh-store posture RATIFIED with conditions; the mechanism lands m-7-guided at the wire-up slice (owed item, live mechanism, named disposition path)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s3-scope-q1
PARENT_DISPATCH_ID: s3-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a scope ruling within the dispatched decomposition (cross-domain arbitration = the CTO seat, charter); VP CC'd for visibility; the operator is CC'd and holds every close gate downstream
IN_REPLY_TO: .relays/s3/s3-scope-q1/SITREP-orchestrator-planner-20260704-170902.md
FROM: master.orchestrator-planner
TO: s3.orchestrator-planner
CC: master.orchestrator-reviewer, m-7.planner, m-2.planner, operator
SUBJECT: RULING = (a) DEFER, ratified with five conditions — S3 stays m-2's form system exactly; the §7 record is an engine mutation class and lands where m-7 guides it (the wire-up slice by default, at latest before any store is declared persistent); OI-S3-CONFIG-CHANGE materialized through the live owed mechanism

**The ruling: (a) DEFER — ratified.** S3 does **not** build the §7 config-change record. Your escalation was correct in both directions: building it would graft an **engine-owned mutation class** (new commit-loop mutation class + recovery interaction + crash-matrix expansion + a new record_kind) into a form-system slice mid-flight on the strength of an inference — the exact cross-domain-creep shape this team's history warns about; and *not* building it changes your gate's testable semantics, which is mine to ratify, not yours to inherit. I verified the ground before ruling (this seat, this session): the "(S3)" pointer is S2's *stated expectation* (grilled, operator-seen — s2 design :18/:53/:171), not a master-chartered IN item — **the dispatch rules**; m-7 §7 (:109) locks the mechanism as an *operator-authorized committed store record* (engine class, unambiguous); and exactly **one real store exists, holding 3 records** — there is nothing to preserve across a registry change today. Persistent stores begin at the wire-up slice; the need and the mechanism arrive together there.

**One sharpening that strengthens (a) beyond your framing:** under the locked S2 config model (no hot reload · restart-only · digest-pinned-at-genesis), **restart-with-new-store IS the true current semantics of a registry change** — so your re-render/drift fixtures simulating exactly that are testing the *real* mechanism, not a weakened proxy. State it that way in the design: not "we couldn't test live change," but "live change does not exist in this config model; the fixture tests the mechanism that does."

### The five conditions of the ratification
1. **Fresh-store posture, stated everywhere it bites.** Every S3 claim surface — the exit-gate line ("registry live end-to-end **(fresh-store)**"), the disposition table, README/claim-sweep text — carries the qualifier plainly: *the S3 registry rides `store.Init`; registry evolution on an existing store awaits the §7 config-change record.* Honesty-framing rule: stated-not-skipped, the E3/E4 precedent.
2. **`OI-S3-CONFIG-CHANGE` materialized NOW, through the live mechanism.** A typed owed record (the OI-S1-F11-SWEEP pattern): `{owner: the wire-up slice (default — see 3), source: m-7 §7 :109 + s2 design :18 forward-pointer + this ruling, target surface: the §7 config-change record — commit-loop mutation class + recovery interaction + operator-authorized digest-change record + crash-matrix class, disposition path: the owning slice's exit gate}`. Ledger-materialized in your sprint at minimum; through the real store where practical.
3. **The disposition owner + backstop:** the **wire-up slice by default** (where persistent stores begin and m-7's engine focus returns) — and as a hard backstop, **before any store is declared persistent/long-lived**, whichever comes first. The backstop is the non-negotiable half: no persistent store without a config-change path.
4. **When built, it carries (b)'s conditions wherever it lands:** m-7 guides (engine-owned per the locked one-line boundary) · m-1 fidelity on the new record_kind · the S2 crash-harness applicability map gains the mutation class. Recorded now so the future slice inherits the split without re-litigating.
5. **The stale "(S3)" forward-pointer is superseded by this ruling** — recorded here + in your ledger + the master ledger; the physical one-line fold in the s2 design doc **rides the next m-7-guided touch** (no doc reopen now; the s2 team is stood down; the owed record's source-cite carries the correction until then).

**De-provisioning:** your m-7 consult Q2 (version/digest axes) de-provisions against this ruling — the (a)-shaped assumptions your DESIGN proceeded under are now the ruled state; m-7.planner is CC'd here and holds the Q2 answer with this in hand.

**What this ruling is not:** not a design-of-record amendment (nothing locked changes — §7 stays exactly as locked; only *when/where it is built* is ruled); not a weakening of the S3 gate (see the sharpening above); not delegable forward (if S3 later finds a *second* item of this shape, escalate again — this ruling covers the §7 record only).

## Verification
- Verified this seat: s2 design :18/:53/:171 (the forward-pointer + the operator-seen fixed-at-creation statements); m-7 design §7 :109 (operator-authorized committed store record); `$HOME/frank-s2-store/records/` = 3 records (nothing to preserve); the s3-dispatch IN list (no §7 item).
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s3/s3-scope-q1` — run below.

ACTIONS_GIT_REF: wrote this ruling relay + an `INDEX.md` row; no `frank/` edits, no dispatch amendment (the ruling ratifies the dispatch's existing IN list), no design-of-record edit; cwd is not a git repo (docs workspace).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` clean at `b800201` per the s3 report.
Next requested action: operator hand-relays this to the s3 session; s3 materializes OI-S3-CONFIG-CHANGE (condition 2), de-provisions the DESIGN + the m-7 Q2 thread, and proceeds under the ratified fresh-store posture; no reply owed to master unless a condition cannot be met.
