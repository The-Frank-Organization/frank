## Team m-6 - Human Surface & Scheduler AUDIT - Implementer independent pass

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c3-audit-m-6
PARENT_DISPATCH_ID: c3-audit-m-6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- read-only audit; operator-judgment items surfaced below
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-5.planner, m-5.implementer, operator
BUNDLE_ID: m-6-human-surface-scheduler
OWNER: m-6 (Human Surface & Scheduler)

IN_REPLY_TO: master/relays/c3-audit-m-6/AUDIT-orchestrator-planner-20260630-051950.md

## Verdict

PRIMARY_BUCKET: still-open
still-open: m-6 has no prior design-of-record; c3 is explicitly the first full design for gate->email buckets, ODB surface, scheduler park/wake, away bridge, and the m-5 seam (master/domains/m-6-human-surface-scheduler/README.md:3-24). New design work is required, but it should be a thin projection over locked upstream records rather than a new gate system.
already-closed: strong promote/wire candidates exist: m-1 already gives the operator mailbox/projection target (master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:107-118); m-2 already declares bucket drivers and ODB slots (master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:250-276); m-3 already exposes observe evidence, record_integrity, and fail-closed egress (master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:41-63,116-122,146-151); m-4 already exposes routing failure/escalation without a new gate class (master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:144-158,211-213); agent-scripts already has the ODB content rule (references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:53-69); v2.8.8 already has TO/CC action/context semantics (extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/agent-pair-implementer/protocol.md:28-40).
product-overlapped: m-6 must not absorb sibling mechanisms. m-5 owns the concrete human-mode vocabulary, tag-space, sensor/actuator archetypes, and authority ceilings (master/domains/m-5-workflows-archetypes/README.md:7-23); m-3 owns egress/observe; m-4 owns routing; m-1 owns store/addressing/stamping. m-6 binds surfaces to those records and hosts the operator-facing projection.
recommended-next: DESIGN should lock a minimal local-first governance inbox + scheduler: A/B/C/D bucket projection, ODB rendering and reply grammar, parked-lane state machine, opt-in egress-gated away bridge, meeting/collaboration surfaces, and explicit m-5 seam placeholders. Defer all concrete per-archetype human-mode behavior until m-5 declares the vocabulary.

## Audit Basis

Read-only sources inspected:
- Dispatch and live status: master/relays/c3-audit-m-6/AUDIT-orchestrator-planner-20260630-051950.md; master/README.md.
- Locked contracts: master/ARCHITECTURE.md:79-102,131-164,180-220; m-1/m-2/m-3/m-4 design docs cited inline.
- m-6 and m-5 domain charters: master/domains/m-6-human-surface-scheduler/README.md; master/domains/m-5-workflows-archetypes/README.md.
- Prior art: references/agent-scripts, v2.8.8 release, references/jcode-ux-notes.md, references/codex-notes.md, references/claude-code.
- Web prior-art check, non-binding and used only as corroboration: Temporal HITL workflows document durable signal/timer/audit patterns; GitHub Copilot cloud-agent docs show background work with PR/session-log handoff; Claude Code hooks docs show notification/hook surfaces. Local file:line evidence remains the lock evidence.

## Gate To Email Bucket Taxonomy

The minimal design should use four buckets, but the "email" surface is a projection. Step 1/2 default is local-only over the relay store; external email appears only through opt-in away mode after m-3 egress.

1. Bucket A - human decision required.
   - Inputs: HUMAN_GATE_REQUIRED at or above the monotonic floor, gate_category in the A set, routing_unavailable/human_decision_required escalation, or any unclassified "other" fail-safe (master/ARCHITECTURE.md:80-98,210-213).
   - Surface behavior: top of operator inbox; always ODB-backed; lane parks; resummon cadence applies; never auto-approve.
   - Return path: operator selects one exact choice from the ODB choices row_array or rejects/asks for more information. The reply becomes a stamped FROM: operator record through the operator-relay channel, not a lane-authored payload (master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:107-110).

2. Bucket B - orchestrator absorbed / governance visible.
   - Inputs: routing, sequencing, and scope-within-bounds categories from the B set (master/ARCHITECTURE.md:94-98; master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:253-256).
   - Surface behavior: visible in the local governance inbox/digest for auditability, but not mirrored externally by default and not parked on the operator unless a monotonic raiser promotes it to A.
   - Return path: orchestrator action stays in the relay graph; operator gets optional digest visibility, not a blocking decision.

3. Bucket C - FYI / CC-only.
   - Inputs: CC-addressed, accepted, non-gating records and informational meeting/collaboration items. v2.8.8 defines CC as context-only with no action obligation (extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/agent-pair-implementer/protocol.md:28-40).
   - Surface behavior: low-priority local digest and searchable thread context. No lane park.
   - Return path: explicit human reply is a new relay, not an implicit decision on the original item.

4. Bucket D - bounce / repair required.
   - Inputs: delivery_state=bounced and failing_edge, including form, lineage, observe, or egress failure (master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:256-258; master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:61-63).
   - Surface behavior: sent to the author/owner repair queue with the failing edge named; operator only sees it if the failed edge itself is an A-category governance item or the resummon policy says to surface stalled repairs.
   - Return path: author resubmits a corrected relay; the original stays as a non-deliverable/bounced attempt.

Design guardrail: no bucket may exist without a writer. The four above all have upstream writers in m-1..m-4. Anything beyond them should be rejected, deferred, or routed to m-5/config.

## Owner Decision Brief Design

Promote agent-scripts' ODB rule, do not invent a new brief. Its useful content invariant is: do not ask the owner for approval/access/land-delete with only a URL or status; refresh state first; include change, why now, proof, risks, recommendation, and exact choices (references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:53-69).

The v3 ODB should render these fields from the locked record:
- Envelope/provenance: subject_ref, FROM, TO, PARENT, DISPATCH_ID, and thread context from the stamped relay record (master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:271-276).
- Human-readable body: plain_language_change, why_now, tradeoffs_risks, recommendation.
- Evidence: completed_proof as m-3 evidence_ref, achieved_evidence, target_gap_result, evidence_integrity, and record_integrity, never agent free-text proof (master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:265-274; master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:107-112).
- Decision controls: choices as an enumerated row_array; gate_category; decision_deadline null by default; on_timeout default hold_and_resummon (master/ARCHITECTURE.md:80-83).
- Egress/away status: local-only by default; if away-mode mirroring is active, show egress_scan_result before any external send (master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:116-122).

Decision-capture rule: the reply parser accepts only one of the enumerated choices or an explicit "need more info / reject with reason" path. Free-form prose may accompany a choice, but it must not create a new hidden choice or lower a gate. Timeouts only park/resummon or block more conservatively.

## Scheduler Park/Wake And Away Bridge

Minimal state machine:
- active: lane can proceed.
- parked_waiting_human: Bucket A gate accepted and ODB emitted; lane is paused.
- resummon_due: no valid reply by cadence point; local prompt/digest escalates, never resolves the gate.
- replied_pending_validation: operator reply captured; form/lineage validation checks the reply against the ODB choices and operator-FROM channel.
- resumed: valid reply wakes the lane by appending the operator decision relay and delivering it to the parked lane.
- bounced_repair: form/lineage/observe failure; routes to Bucket D.
- egress_blocked: away-mode outbound scan blocked; park locally and resummon, never auto-redact/send by default.

Step fit:
- Step 1: local outbox item exists; no external send (ROADMAP.md:57-65).
- Step 2: local inbox/outbox + scheduler parks lanes and wakes on reply (ROADMAP.md:67-72).
- Step 4: full email-client UX can replace the minimal projection without changing the record contracts (ROADMAP.md:82-86).

Away bridge:
- Opt-in only; default remains local in-app inbox (master/ARCHITECTURE.md:84-87).
- Mirrors only Bucket A by default. Bucket B/C can be local digests unless the operator config promotes them.
- First external send goes through m-3 fail-closed egress; any blocked scan creates egress_blocked and a local resummon (master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:116-122).
- Reply ingestion must produce a stamped FROM: operator record through the operator channel; an email reply is evidence/input to that channel, not a lane credential (master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:107-110).

## Email Governance And Meeting Collaboration Surface

Design target: a calm local governance inbox over the relay store, not a noisy general GUI. jcode is the negative look: the notes call its GUI noisy/filler and its /btw side-panel deferral ugly/late (references/jcode-ux-notes.md:6-21,58-77). Codex is the positive look: stable scrollback, mutable tail, URL-aware wrapping, bottom composer, and queue-while-busy input (references/codex-notes.md:34-55,62-79).

Email-governance surface:
- show buckets as operator work queues, with A as blocking, B as absorbed/governance-visible, C as FYI, D as repair;
- render ODBs as compact decision cards backed by the record, not nested prose summaries;
- keep accepted/bounced state and failing_edge visible;
- preserve thread/provenance as first-class navigation, because m-1's inbox is a projection of the addressing graph (master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:118).

Meeting-collaboration surface:
- design/open-question traffic stays conversational and local unless it raises a HUMAN_GATE;
- verdict gates should follow agent-scripts' "drive to decision boundary, then ask once" rule for owner decisions (references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:41-69);
- scheduler should support "park until meeting answer" for explicit human questions, but should not convert every discussion message into a governance gate.

## m-5 Seam A/B

Seam A - consume m-5 human-mode vocabulary:
- m-6 should define a binding table shape, not concrete values yet: human_mode -> default bucket intensity, allowed notification channels, resummon cadence class, interjection affordances, and whether meeting-style collaboration is enabled.
- Concrete human_mode names and per-archetype defaults are m-5-owned and must arrive before m-6 binds behavior (master/domains/m-5-workflows-archetypes/README.md:7-23; master/domains/m-6-human-surface-scheduler/README.md:22-30).

Seam B - host interjection surface:
- m-6 owns the operator affordance: steer / side-question / interrupt.
- m-5 owns the side-question sensor archetype; m-4 routes it; runtime owns boundary injection and soft-cancel (ROADMAP.md:32-46; master/ARCHITECTURE.md:216-220).
- Adopt the Claude Code/Codex positive pattern: ordinary steer queues to the next safe boundary; side-question is a read-only, tool-blocked, one-turn fork; interrupt cancels and redelivers. Avoid jcode's side-panel/whole-task deferral (references/jcode-ux-notes.md:37-77; references/codex-notes.md:77-79).

Blocking seam risk for DESIGN: if m-5 has not declared at least the abstract human-mode vocabulary shape, m-6 can lock the neutral bucket/scheduler substrate but must leave archetype-specific behavior as a named binding placeholder.

## Findings

1. Still-open: no m-6 design-of-record exists. Evidence: master/domains/m-6-human-surface-scheduler/README.md:3-24. Impact: c3 design must lock the surface, scheduler, ODB rendering, and m-5 seam before Step 0 close.

2. Already-closed/promote: ODB field content and schema are sufficient to promote. Evidence: references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:53-69 and master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:260-276. Impact: do not design a new decision-brief schema; design the renderer and reply path.

3. Already-closed/promote: A/B/C/D bucket drivers exist. Evidence: master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:250-258 and master/ARCHITECTURE.md:89-102. Impact: m-6 owns projection/config, not field ownership.

4. Still-open: scheduler state machine and cadence are not specified. Evidence: ROADMAP.md:67-72 and master/ARCHITECTURE.md:80-87 only state park/wake, hold/resummon, and away-mode defaults. Impact: DESIGN must lock states, reply validation, cadence classes, and failure transitions.

5. Product-overlapped: concrete per-archetype human-mode behavior belongs to m-5. Evidence: master/domains/m-5-workflows-archetypes/README.md:7-23 and master/domains/m-6-human-surface-scheduler/README.md:22-30. Impact: m-6 must not pre-bind behavior before m-5 declares vocabulary.

6. Still-open: external away-mode bridge is a first-send risk, but upstream egress is adequate. Evidence: master/ARCHITECTURE.md:84-87 and master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:116-122. Impact: DESIGN should make away-mode opt-in, A-only by default, and egress-blocked -> local resummon.

## Operator-Judgment Items

- Away-mode destination policy: which inboxes are allowed, whether Gmail is the first supported bridge, and whether B/C digests may ever mirror externally by default.
- Resummon cadence: quiet/local vs push/loud escalation classes, and any maximum resummon count before block.
- Reply grammar: whether operator replies must use exact choice IDs, short aliases, or structured buttons only.
- Bucket B visibility: whether B items are hidden from the operator by default, daily digest, or live local digest.
- Meeting-collaboration boundary: which human questions are conversational vs governance gates.

## Coordination

No gap relay to m-1..m-4 is needed from this audit: their writer surfaces are sufficient for m-6. A seam coordination thread with m-5 is required in DESIGN before m-6 binds concrete human-mode behavior.

## Verification

Subagents used:
- Local prior-art scout for agent-scripts/v2.8.8/jcode/claude-code/codex reference evidence.
- Locked-contract scout for m-1..m-5 and c1/c2 boundary evidence.

Websearches used:
- Temporal human-in-the-loop workflow docs: durable waiting/signal/timer/audit prior art.
- GitHub Copilot cloud-agent docs: background agent sessions with PR/session-log handoff prior art.
- Claude Code hooks docs: hook/notification automation prior art.

RELAY_LINT: OK -- python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c3-audit-m-6/AUDIT-implementer-20260630-053053.md
RELAY_ROOT_LINT: non-clean due to INDEX.md only -- python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/c3-audit-m-6/AUDIT-implementer-20260630-053053.md reports OK for this relay file, then errors on lint-exempt master/relays/INDEX.md missing relay header fields.

ACTIONS_GIT_REF: relay artifact created at master/relays/c3-audit-m-6/AUDIT-implementer-20260630-053053.md and master/relays/INDEX.md row appended; git status unavailable at /Users/jack/Programming/harness because it is not a git repository; pcode/ status clean.

FINAL_GIT_STATUS_SHORT: unavailable -- fatal: not a git repository (or any of the parent directories): .git
