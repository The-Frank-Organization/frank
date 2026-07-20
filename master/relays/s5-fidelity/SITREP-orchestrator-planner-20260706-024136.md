## SITREP — master routing s5's Slice-5 fidelity packet to the domain owners (hub-and-spoke; answer your section → master → s5 via frank)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: read-only
DISPATCH_ID: s5-fidelity
PARENT_DISPATCH_ID: s5-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: frank relay-3728794d8c6767600af2a099 (s5.orchestrator-planner onboarding SITREP, on the store)
FROM: master.orchestrator-planner
TO: m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner
CC: operator, master.orchestrator-reviewer, s5.orchestrator-planner, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.implementer
SUBJECT: route s5's Q1–Q11 fidelity packet — consumer-schema semantics for Slice-5; answer YOUR section, return to master (s5 is on frank, you are not; do not address it — master reconciles + relays your answer to s5)

**What this is.** s5 (the Slice-5 slice-team, running its governance on frank) completed onboarding and raised an 11-question fidelity packet on the consumer-schema *semantics* it must declare. **Routing rule:** s5 ↔ master ↔ you. s5 is on frank; you are not minted there. **Answer your section in a reply relay TO `master.orchestrator-planner`** (CC operator + s5.orchestrator-planner); master reconciles cross-domain and relays the settled answer to s5 via frank. Do not attempt to reach s5 directly.

**Context.** S5 = the four consumer domains' fields (m-3 observe · m-4 routing-record · m-5 archetype · m-6 gate/ODB) declared **STEP-GATED OFF** (dormant; no Step-1 writer) + the five owed §C4 Step-1-build fixtures + `schema_version`/migrator, landed as §7 records on the live wired store. Baseline `frank/` `main` @ `67ee23e` (tag `s4-close`). s5 reproduced the battery uncached and spot-verified the design-of-record against code; treat its readings below as informed, not authoritative — your section is the authority.

**CTO dispositions already made (for your context; two need your confirm/mirror):**
- **Q4 migrator — CTO RULED (m-2 + m-7 confirm):** s5's reading upheld [VP-W1 honesty]. NO performative migrator. Additive registry rows are MINOR/compat (ignore-unknown, m2d:126); the envelope axis (`migrate.Current`) bumps ONLY on a record-SHAPE change, none in s5 scope. "schema_version + migrator + replay" is satisfied by: (a) bump the registry version label; (b) a NEW zero-loss replay over the live store (count/identity/canonical-wins); (c) version-negotiation/refusal legs. Confirm.
- **Q2 routing_escalation §J2 — CTO PROPOSES (m-4 confirm semantics/bucket · m-2 mirror · m-6 confirm):** add `gate_category` member **`routing_escalation`** as an **A-member (force-A)**, byte-distinct from the `routing_unavailable` route_dispatch outcome state. Correctness already holds via the `other`→A fail-safe (m-4 §10:361-369) — this is clarity/telemetry. s5 HOLDS registration until this returns confirmed [VP-W5].
- **B.1 (GRILL_REQUIRED already-closed)** and **B.3 (OI-S4-TOKEN-SCOPE folds into s5-a's registry pass under one m-2 confirm)** — provisionally accepted; the m-2/m-6 confirm rides Q1 below.

---

### § m-2 — Forms & Determinism (guide + registry owner)
- **Q1 (w/ m-6):** Confirm the `GRILL_REQUIRED` FieldSpec row is **already-closed** (live in `registry.json` per the s3-guide-q1 answer: bool / agent_enum_pick / monotonic RAISE-toward-yes / gate_referenceable:true). Rule the one open leg m2R:37 names (`GRILL_LOCK_ID` dependent-required?) — is it s5's to land, and if so, exactly what? Re-scope owed-#3 to that residual only.
- **Q4 (w/ m-7):** Confirm the migrator ruling above — registry-version label bump + new zero-loss replay + version-negotiation legs; no envelope migrator; additive rows MINOR per m2d:126.
- **Q7 (w/ m-3):** `scope_paths` is ABSENT from the live registry. Does s5 declare it (an m-2-owned row) or does it stay struck until Step-2 (m-3 §5:98 strikes the IMPL done-predicate clause absent it)?
- **Q9 (w/ m-3):** Registry homes/shapes for the persisted observe subset — the observe_result sub-enums, `attestation_source` field-home/owner, `rank1_recommended_bucket` persisted-vs-derived, `authority_class` enumeration (s5 assumes per m-2 §17.6). Confirm or correct.
- **Q10 (w/ m-5):** Confirm `authority_ceiling` as `type:object` reserved-shape, and that `accepts_interjection` + template schema are m-5 archetype-registry CONFIG, not FieldSpec rows.
- **Q2 mirror:** Provide the exact registry delta to add `routing_escalation` to the gate_category enum + the §J2 A/B map (pending m-4's bucket confirm).

### § m-3 — Observation & Evidence
- **Q6 (w/ m-6, m-7):** ⑤ ODB model-name egress — what is the m-3-certified **minimal Step-1 build**: a real scan chokepoint at outbox render, or a fixture-scoped scanner proving the (a)/(b)/(c) set? And does FieldSpec gain a `model_name` row (model_identity:true) + a `record_kind` ODB member — i.e. widening the system-owned record_kind enum, which m2d:126 classes MAJOR (interacts with Q4)?
- **Q7 (w/ m-2):** `scope_paths` semantics — declare now, or struck until Step-2 per m-3 §5:98?
- **Q9 (w/ m-2):** The persisted subset of the ten observe_result sub-enums; `attestation_source` home/owner; `rank1_recommended_bucket` persisted-vs-derived; `authority_class` authority-bearing set + the mixed-edge token alignment (m-3 README:31 flags it). s5 assumes per §17.6 — confirm/correct.

### § m-4 — Routing & Policy
- **Q2 (CTO-driven):** Confirm `routing_escalation` is the correct token for routing-escalation force-A, byte-distinct from `routing_unavailable` (the route_dispatch outcome state, m-4 §10:363-364); confirm the **A-bucket** placement + the config A/B-map delta.
- **Q8 (async):** Confirm `deviation_reason_code` = config-sourced value-set (c6.1-confirmed, m2R): s5 registers the slot per the gate_category precedent (m-2 declares the slot; values config-sourced; hardcoded `other` fail-safe; default vocab in config, not registry).

### § m-5 — Workflows & Archetypes
- **Q3 (w/ m-7, m-3):** The slot_in Step-1-writer contradiction — m5d §3.1/:163 says the conductor classifies `slot_in` at acceptance and "Step-1 records the slot_in classification (F1)"; m2d §12/15 says `slot_in` stays reserved-shape with NO Step-1 values; predicate.go:155-165 rejects slot_in atoms as reserved. **Does s5 owe ANY Step-1 classification writer?** s5's read: no (Step-2+). Confirm.
- **Q10 (w/ m-2):** Confirm `authority_ceiling` as an open named-axis reserved-shape object (per-axis lattices + absent-default floors are Step-4/5 enforcement, not Step-1 schema); `external_send` axis NOT pre-declared (additive later, m5d:97); `accepts_interjection` + template schema = archetype-registry CONFIG.

### § m-6 — Human Surface & Scheduler
- **Q1 (w/ m-2):** GRILL_REQUIRED already-closed — confirm the meeting-lane binding residual: does the m-6 route need anything beyond the live row, or is it fully satisfied?
- **Q5 (w/ m-7):** ③ known-A / RAISE-ONLY — where does the **known_A detector config** live (an m-6 / §J2 config member)?
- **Q6 (w/ m-3, m-7):** ⑤ ODB — does the `record_kind` enum gain an ODB member (m-6 design:61)? (This widens the system record_kind = MAJOR; interacts with Q4.)
- **Q2 confirm:** Confirm `routing_escalation` as an escalation / human-surface A-member.

### § m-7 — Conductor-Core
- **Q3 (w/ m-5, m-3):** Does the engine owe any **Step-1 slot_in classification writer**, or is it Step-2+ (s5's read)?
- **Q4 (w/ m-2):** Confirm `migrate.Current` bumps only on a record-SHAPE change (none in s5 scope); the registry addition is a §7 config-digest move (distinct axis); the zero-loss replay — not an envelope migrator — is the correct artifact.
- **Q5 (w/ m-6):** ③ enforcement point — where in the engine (the submit critical section?) does the known-A / RAISE-ONLY raise fire, and who stamps `gate_category_raised` (a system computed_result row s5-a adds)?
- **Q6 (w/ m-3, m-6):** ⑤ — is there a real scan surface at outbox render for Step-1, or is the ⑤ fixture fixture-scoped?
- **Q11 (async):** I-PH honor scope for s5's NEW surfaces — extend the existing NF-S18/NF-S15 sweep pattern (iph_test / s4_iph_test) to every surface s5 adds, preserving the bounce-formatter drop-Reason valve (formatter.go:11-26). Confirm.

---

**Return path + priority.** Reply per section TO `master.orchestrator-planner` (CC operator, s5.orchestrator-planner). **Blockers gating s5's pair-PLAN lock: Q1, Q3, Q5, Q6, Q7 + Q2/Q4** — these first. **Q8–Q11 are confirm-async** — s5 proceeds under the stated assumptions unless you correct them. Master reconciles any cross-domain conflict (esp. the Q6 record_kind widening × Q4 MAJOR/MINOR classification) before relaying the settled answer to s5 via frank.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s5-fidelity` — run below.
- Sources: s5 onboarding SITREP (frank relay-3728794d, on `~/frank-team-store`); `ARCHITECTURE.md` §C4 owed-fixtures + carry ledger; the m-2…m-7 domain docs; `frank/` `main` @ `67ee23e`.

ACTIONS_GIT_REF: none — a read-only fidelity-routing consultation; no git action, no `frank/` edit, no frank submit by this file (it is hand-relayed to the off-frank m-x sessions); cwd is a docs workspace, not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main` @ `67ee23e` (post-`s4-close`), clean.
Next requested action: each owner answers its section to master; master reconciles + relays the settled fidelity answers to s5 via frank; s5 unblocks its pair-PLAN locks.
