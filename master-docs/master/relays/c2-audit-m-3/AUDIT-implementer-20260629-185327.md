## Team m-3 Implementer - Observation & Evidence AUDIT

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-audit-m-3
PARENT_DISPATCH_ID: c2-audit-m-3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-3-observation-evidence
OWNER: m-3 (Observation & Evidence), implementer audit lens

AUDIT_VERDICT: sufficient for DESIGN entry; frank m-3 is not already built.

PRIMARY_BUCKET: recommended-next
still-open: the exact frank observe-as-send primitive is absent. The upstream protocol has protocol/lint approximations, jcode and Claude Code have permission/hook gates, and agent-scripts has strong proof/egress policy, but none system-fills conductor-observed evidence onto the relay being delivered.
already-closed: promote the useful pieces: the upstream protocol's evidence ladder/action-ref discipline; Claude Code Stop/PreTool/PostTool hook shapes; jcode permission queue/review queue; agent-scripts Owner Decision Brief, live-proof discipline, Public Model Identifier Gate, and degradation-honest observation vocabulary.
product-overlapped: no ownership overlap requiring reroute. m-3 owns the observe/evidence primitive; m-4, m-5, and m-6 are consumers/seam partners, not alternate owners.
recommended-next: DESIGN a frank-owned observer hook/API against the locked c1 write-allowlist and DI-5 read vantage, then coordinate the m-3<->m-4 evidence-record seam and the m-5 archetype/tag parameterization before lock.

## Duplicate / Already-Built Gate

No reference provides the required trust primitive.

- The upstream protocol defines the E0-E4 evidence ladder and warns not to call claims fixed/done/safe/verified without evidence level (E1: the upstream protocol release corpus (not vendored), `agent-pair-implementer/protocol.md:192-202`), requires action/git proof and final status for read-only reports (E1: the upstream protocol release corpus (not vendored), `agent-pair-implementer/protocol.md:204-227`), and relay-lint enforces missing/substantive status/action fields structurally (E1: `<upstream relay-lint tools>/relay-lint.py:927-944`). But protocol also states incoming sitreps are E0 until reconciled and action claims without git evidence are E0 (E1: the upstream protocol release corpus (not vendored), `agent-pair-implementer/protocol.md:364`), and relay-lint is explicitly truth-agnostic (E1: the upstream protocol release corpus (not vendored), `agent-pair-implementer/protocol.md:395-405`).
- The upstream protocol's dispatch/review gates are useful gate precedent, not observed done evidence. Literal dispatch shape and delegated SCOPE_DIFF are linted before dispatch (E1: `<upstream relay-lint tools>/relay-lint.py:968-985`); orchestrator authority relays must address an orchestrator-reviewer in TO/CC unless operator-waived (E1: `<upstream relay-lint tools>/relay-lint.py:1075-1129`). These gates prove routing/lineage shape, not ground truth.
- jcode has a persistent safety queue and permission request types (E1: `references/jcode/crates/jcode-base/src/safety.rs:54-74`, `references/jcode/crates/jcode-base/src/safety.rs:150-199`), and its design makes external/human communication permission-required (E1: `references/jcode/docs/SAFETY_SYSTEM.md:17`, `references/jcode/docs/SAFETY_SYSTEM.md:102-113`). But `request_permission` is ambient-session-only and queues agent-authored action/rationale/context (E1: `references/jcode/crates/jcode-app-core/src/tool/ambient.rs:93-101`, `references/jcode/crates/jcode-app-core/src/tool/ambient.rs:640-683`).
- jcode `pre_tool` can block a tool and `post_tool` observes result metadata before logging lifecycle done (E1: `references/jcode/crates/jcode-app-core/src/tool/mod.rs:584-620`, `references/jcode/crates/jcode-app-core/src/tool/mod.rs:637-645`). That is prior art for lifecycle hooks, not typed relay evidence.
- Claude Code has a stronger generic hook system: PreToolUse can return permission decisions/updated input; PostToolUse can add context/update MCP output; PermissionRequest can allow/deny (E1: `references/claude-code/src/types/hooks.ts:70-130`, `references/claude-code/src/types/hooks.ts:250-270`). Tool execution runs pre-tool hooks before permission resolution and tool execution (E1: `references/claude-code/src/services/tools/toolExecution.ts:800-805`, `references/claude-code/src/services/tools/toolExecution.ts:916-995`), and post-tool hooks after tool output exists (E1: `references/claude-code/src/services/tools/toolExecution.ts:1397-1490`). Stop hooks can block the next turn by injecting blocking errors (E1: `references/claude-code/src/query.ts:1267-1289`, `references/claude-code/src/query/stopHooks.ts:175-265`). This is the best runtime-hook precedent, but not conductor-owned observed evidence on a typed relay.
- agent-scripts provides the best policy/content prior art: Owner Decision Brief requires refreshed state and a seven-part owner question (E1: `references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:53-69`); live proof is required before landing and cannot be inferred from merge/release confidence (E1: `references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:169-181`); Public Model Identifier Gate emits PASS/BLOCKED over diffs, logs, artifacts, and public proof (E1: `references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:183-194`); transcript tooling fail-closes on unresolved secrets/private keys/cookies/auth URLs (E1: `references/agent-scripts/skills/agent-transcript/SKILL.md:10-24`, `references/agent-scripts/skills/agent-transcript/scripts/agent-transcript:191-222`). These should shape schemas and scans, not become the trust mechanism.

## Self-Reported-Done Gap

Named gap: **self-reported-done gap** - current systems can record that an agent claims, reports, or reaches a status that implies done, but the trusted courier has not independently observed the phase-shaped done predicate from outside the lane and bound that observation to the outgoing record.

- The upstream protocol labels this class but does not remove the surface: self-reported/action claims remain E0 until checked, and relay-lint checks shape, not truth (E1: the upstream protocol release corpus (not vendored), `agent-pair-implementer/protocol.md:364`, `agent-pair-implementer/protocol.md:397-405`).
- jcode `await_members` treats member status matching target status, or persisted final response under still-satisfied current status, as done enough for waiters (E1: `references/jcode/crates/jcode-app-core/src/server/comm_await.rs:54-64`, `references/jcode/crates/jcode-app-core/src/server/comm_await.rs:189-195`, `references/jcode/crates/jcode-app-core/src/server/comm_await.rs:272-280`). That is status/output driven, not independent task-evidence observation.
- Claude Code task prompts instruct the model not to mark partial work complete (E1: `references/claude-code/src/tools/TaskUpdateTool/prompt.ts:13-20`), and remote tool results are converted by shape while the adapter notes a parent ID is unreliable for distinguishing tool results from prompt echoes (E1: `references/claude-code/src/remote/sdkMessageAdapter.ts:176-185`). This leaves completion largely prompt/status/tool-result governed.
- c1 locked the closure seam: DI-5 says the conductor reads lane ground truth from outside the lane, and R3 says observed evidence carries `evidence_integrity {observed | self_reported}` with the m-3 hook observer-only against a positive write-allowlist (E1: `master/ARCHITECTURE.md:31-34`, `master/ARCHITECTURE.md:70-71`).
- The frank design intent closes the gap by making the conductor system-fill evidence fields from outside the lane at submit time, with `send` itself triggering observation so the passing observation atomically binds to the emitted relay (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:43-45).

Closure statement: observe-as-send + `evidence_integrity` closes the gap by removing observable evidence fields from the lane's writable surface. A lane may still provide judgment/prose and declare an `EVIDENCE_TARGET`, but `achieved_evidence`, probe outputs, target-gap result, executable-claim results, and egress scan result are conductor-observed or explicitly labeled `self_reported`. If the predicate fails, the relay is not deliverable.

## Design Recommendation

Minimal primitive:

```text
observe_gate(record_candidate, seat, phase, authority, evidence_target, artifact_refs, executable_claims, egress_targets, archetype_tags)
  -> observe_result {
       predicate_result: pass | fail | blocked | degraded
       veto: none | block_delivery
       evidence_integrity: observed | self_reported | mixed
       achieved_evidence: E0 | E1 | E2 | E3 | E4
       target_gap_result: met | target_gt_achieved | not_applicable
       observed_fields: closed m-3 system/computed field set
       executable_claim_results: per-claim pass | fail | skipped | unsafe
       egress_scan_result: pass | blocked | not_applicable
       degradation_notes: bounded-observation floor/truncation/backoff notes
     }
```

Placement:
- Run inside m-1 `submit()` after m-2 form-validation and before append/delivery/acceptance; the c1 architecture already carries the two-state submitted-to-accepted write path and `blocks before dispatch` lineage semantics (E1: `master/ARCHITECTURE.md:56-59`).
- Execute from the conductor's DI-5 outside-lane vantage, not from the lane. The lane triggers by attempting send; the conductor executes. This preserves forcing-function feedback while avoiding self-observation forgery (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:43-45).
- The hook is observer-only: it writes only the closed m-2-declared m-3 observed/computed field set plus veto, matching R3 (E1: `master/ARCHITECTURE.md:70-71`).

Evidence ladder:
- Keep E0-E4 as the vocabulary, but split `EVIDENCE_TARGET` (agent intent) from `achieved_evidence` (conductor fact). The adaptive-routing pillar already frames target vs achieved this way (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:43).
- E1/E2 are first-class Step-1 probes: source/file presence, git/worktree status, lint/test command output. E3/E4 remain progressively deeper runtime/live verification, with E4 often operator/live-verify gated (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:43).
- Executable claims are structured attachments, not executable message bodies. The agent supplies a runnable check; conductor runs it as part of observe-send and stamps the result (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:47).

Done predicates:
- Base predicates are phase-shaped: AUDIT/DESIGN read-only report requires allowed relay artifact plus no unauthorized source actions; IMPL requires expected diff/artifacts, green required checks, and no scope drift; MERGE/LIVE-VERIFY requires target commit/deploy/live proof as applicable. The adaptive-routing pillar names phase-shaped done examples and defers the exact predicate list to m-3 build-time design (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:45).
- m-5 archetype tags parameterize predicate add-ons. The existing design notes already name mutation-tag gates: extension = additive-only diff + existing-suite-green with zero test edits; refactor = suite-green and test-files-unchanged; cleanup = find-references=0; bugfix = parent red / fix green differential; migration = reversibility (E1: the pre-build design-state export (not vendored), adaptive-routing pillar doc:53-58).

Egress/content-safety gate:
- Add an egress pre-publish stage for external targets: PR/public issue/comment, Slack/Discord/email, and away-mode email bridge. It scans outbound body plus attached/generated proof artifacts for secrets, PII/private user data, undisclosed model identifiers, auth URLs/tokens, and destination/audience mismatch.
- The c1 architecture already identifies away-mode external bridge as the first external send and requires fail-closed egress scanning for secrets, PII, and model names (E1: `master/ARCHITECTURE.md:84-87`).
- Export references explicitly call egress/content-safety a new missing gate class, with Owner Decision Brief and bounded-observation patterns worth taking (E1: the pre-build design-state export (not vendored), `EXTERNAL-REFERENCES.md:108-124`).
- jcode's `send_message` path sends the supplied string to a configured channel or all send-enabled channels without a visible prior scan in that tool path (E1: `references/jcode/crates/jcode-app-core/src/tool/ambient.rs:1044-1112`). That is a negative prior-art datum: frank should put egress at the conductor boundary, not only in prompt policy.

## Consumer Boundary Contract

m-3 exposes to m-4 Routing & Policy:
- `observed_evidence_ref`: immutable pointer to the observe atom for the accepted record.
- `achieved_evidence`, `target_gap_result`, and selected probe summaries as benchmark/routing-quality inputs.
- `evidence_integrity` so routing records and benchmark outcomes can distinguish observed proof from self-reported claims.
- Seam note: a routing record may itself be an evidenced record. The m-3<->m-4 DESIGN coord thread should decide whether benchmark observation atoms are the same type as relay observe atoms or a profile of them.

m-3 exposes to m-5 Workflows & Archetypes:
- A tag-parameterized predicate registry: base phase predicate plus archetype/tag predicate add-ons.
- Drift result: declared archetype/tags vs observed diff/behavior.
- Seam note: m-5 owns the tag-space and per-archetype observe invariants. m-3 should own the predicate execution interface and observed result shape. Lock requires explicit m-5 disposition for concrete tags.

m-3 exposes to m-6 Human Surface & Scheduler:
- Gate result classes and bounce text for operator/human surfaces.
- Egress scan result and redacted/safe outbound body metadata for away-mode bridge.
- Owner Decision Brief evidence bundle fields: completed proof, residual risk/missing evidence, recommendation, exact choices. Agent-scripts prior art gives the content shape (E1: `references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:53-69`), but the evidence atoms must be conductor-observed.

No-consumer action: none. Every proposed field/result above has a named consumer or is part of the m-1/m-2 locked write path.

## Seam Notes

m-3<->m-4:
- Do not let routing policy consume raw agent claims as benchmark truth. m-4 should consume m-3 observed atoms or records explicitly marked `self_reported`.
- Coordination question for DESIGN: is `routing_record.observed_evidence_ref` required for all routing decisions, only benchmark loop samples (a later release), or only deviations?

m-5:
- m-3 cannot lock predicate semantics alone. Predicate mechanics are shared, but tag semantics and archetype presets belong to m-5. The m-3 design should use opaque archetype/tag atoms until m-5 supplies the registry.

m-6:
- Away-mode email bridge must not publish before egress scan passes. A blocked egress scan should park the gate and ask locally, not auto-redact-and-send unless operator policy later permits a specific destination/content class.

## Operator Judgment Items

None blocking for AUDIT-to-DESIGN. Future DESIGN decisions likely need operator judgment on:
- Whether executable claims are allowed to run arbitrary project commands or only registry-approved check descriptors.
- Whether fail-closed egress may auto-redact low-risk content or must always block for human review on first release.

## Evidence Levels

- This audit's claims are E1: file:line citations only.
- No E2/E3/E4 verification was attempted; phase is read-only AUDIT.

ACTIONS_GIT_REF: wrote relay report `master/relays/c2-audit-m-3/AUDIT-implementer-20260629-185327.md` and updated `master/relays/INDEX.md`; no source/pcode/branch/commit/PR changes; cwd is not a git repo.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128: fatal: not a git repository (or any of the parent directories): .git)
