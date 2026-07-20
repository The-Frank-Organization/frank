## DESIGN dispatch — §7 stage-1 owned contract: the provider-EGRESS policy delta + the E0 APP-EVENT schema + the F62 E3 observation-record schema / APPLICABILITY EVALUATOR (F65-scoped) (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-1 owned interface contract under the ratified amendment; the operator gates return at the Master+VP interface-lock, not per-artifact
GRILL_REQUIRED: no — stage-1 owner contracts carry pair review + consumer confirmation; the grills ride the stage-4/5 build lanes (§7)
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: author the m-3 stage-1 contract: the provider-send egress-policy class + the E0 worker-carried app-event schema + the F62 external-E3 record schema and applicability evaluator (F65: app/provider-vertical vector; the conductor identity bound separately) — pair-reviewed final bytes, consumers confirm

m-3 — the Step-3 MVP amendment is **ratified + operative** (`master/STEP-3-MVP-AMENDMENT.md` r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`; operator-ratified 2026-07-16, VP approve `…-035505`). Your charter carries the delta. This dispatch opens your **§7 stage-1 owned contract** — the re-cut successor to the held `step3-amend-m3-egress` lane (its reframe-delta floors carry: a NEW egress class, NOT the away-email scanner; no model-name confidentiality inheritance; no fourth `delivery_state`).

### Author (you own these bytes; m-3.implementer pair-reviews the FINAL bytes)
1. **The provider-egress policy delta (§3):** the policy m-8 enforces at `freeze → authorize → attach → send` — deny ⇒ **zero send** (the secret resolver never invoked on the denied path); the policy bytes carry a **`policy_digest`** (you produce it; it binds into the E3 tuple). The honest boundary is ratified text: local tool effects INCLUDING bash-originated egress are NOT governed in the MVP — your policy governs the designated m-9→m-8 provider attempt only.
2. **The E0 app-event schema (§3):** the **self-reported, worker-carried provider report** ("attestation" is avoided) — carried in the m-9 worker's existing `SITREP` body, `event_evidence=E0`/`event_integrity=self_reported`, flooring at E0, never gate-satisfying, never promoting; **no conductor schema/member change** (the tuple below applies to the app-side/external artifacts ONLY).
3. **The F62 E3 observation record + applicability evaluator:** the EXTERNAL E3 record (written by the separate integration harness / operator observer = `observer_id`; stored outside the conductor store) carrying `scope ∈ {build, artifact, run, turn, attempt, relay_record}` + the binding tuple: `run_id`/`turn_id`/`attempt_id` + `run_manifest_digest` (m-10, at run freeze) + `tool_catalog_digest` (m-9's catalog build, mechanically verified at the F63 release-binding) + the **F63 release digests** (`app_main_build_digest`/`m-9_worker_build_digest`/`m-8_build_digest` or a covering `release_digest` — the build pipeline produces them at the post-build release-binding event) + `policy_digest` (yours) + `provider_lane_id`/`observer_id`/`observation_ts`. **The evaluator rule (yours):** an E3 applies to a claim only while ALL bound digests equal the currently-running vector; any mismatch ⇒ non-applicable (re-observe). **F65 scope:** the vector = the app/provider vertical ONLY; the conductor service identity (build digest + governing config identity) is bound SEPARATELY in the exit-test record for the relay-exchange leg, whose evidence is the conductor's own observe-as-send E1/E2 records — never laundered into the provider-turn E3.
4. **The instrumented-negative posture (§3/§10):** the deny→zero-send negatives are instrumented build/integration-test evidence (fake transport/executor counting zero) or an independent live observer — NOT conductor-observed; they never become E2 via an app statement in a conductor record.

### Boundaries
No conductor byte/member/schema change. The MVP ladder is E0–E3 (E4 out of scope). Consumers to confirm: m-8 (the policy it enforces + the E0 events it emits), m-9 (the report carrier), m-10 (the manifest-digest producer seam). No DESIGN-lock, PLAN, T4 token, or code is authorized by this dispatch.

### Return path (§7 stage 1)
m-3.planner authors the DESIGN parented to THIS dispatch → m-3.implementer DESIGN-REVIEW as a uniquely-parented child (fresh review on any byte revision) → report-only SITREP to master naming the approved bytes + hash → consumer confirmations route on master's direction. The Master+VP interface-lock is the gate; no self-declared lock.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner opens the DESIGN on this dispatch (grounding: the ratified amendment §3/§7/§10 + your charter deltas incl. the held egress-lane floors); pair review; report-only SITREP to master.
