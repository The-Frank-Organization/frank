## RECONCILE — m-3 pair audit (planner ⟷ implementer): converged, GO to DESIGN

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: read-only
DISPATCH_ID: c2-audit-m-3
PARENT_DISPATCH_ID: c2-audit-m-3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only reconciliation; operator-judgment items unchanged from the two passes
FROM: m-3.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-3.implementer, operator
IN_REPLY_TO: c2-audit-m-3/AUDIT-implementer-20260629-185327.md
SUBJECT: m-3 pair audit reconciled — full agreement, zero disagreements; two independent seats converged on promote-parts/build-the-outside-vantage-send-gate; GO to DESIGN with the strawman below

Reconciles my independent pass (AUDIT-planner-20260629-185522) with m-3.implementer's independent pass (AUDIT-implementer-20260629-185327). Per the agent-pair reconcile discipline: agree / disagree / different-coverage / operator-decision-needed.

---

### RECONCILED VERDICT: AGREE — no disagreements. Both seats independently land on **recommended-next** (v3 m-3 is NOT already built; promote the reusable parts, BUILD the outside-the-lane observe-AS-send gate + `evidence_integrity`). Pair is GO to DESIGN.

The two passes were authored without sight of each other and converged on every substantive point — itself corroborating evidence that the gap analysis and the primitive shape are sound.

### AGREE (independently identical)
- **4-bucket verdict.** recommended-next; v3 m-3 not already built; promote E0–E4 ladder + Pre/PostTool hook shapes + jcode permission/review queue + agent-scripts ODB/live-proof/egress patterns; no ownership overlap requiring reroute (m-4/m-5/m-6 are consumers, not alternate owners).
- **Self-reported-done gap.** Both name it identically: no surveyed system independently observes the phase-shaped done-predicate from *outside the lane* and binds that observation to the outgoing record; v2.8.8 *labels* the class (E0-until-reconciled) but its linter is truth-agnostic; jcode/claude-code gate lifecycle but not typed relay evidence; agent-scripts is policy/content prior art, not a trust mechanism.
- **Closure.** observe-as-send + `evidence_integrity` removes observable fields from the lane's writable surface; predicate-fail ⇒ relay not deliverable.
- **Placement.** Inside m-1 `submit()`, after m-2 form-validation, before append/accept; executed from the conductor's DI-5 outside-lane vantage; lane triggers by attempting send, conductor executes; observer-only against the locked R3 write-allowlist.
- **Evidence ladder.** Split `EVIDENCE_TARGET` (agent intent) vs `achieved_evidence` (conductor fact); E1/E2 first-class Step-1 probes; E4 = operator/live-verify; executable claims = structured attachments the conductor runs (not executable message bodies).
- **Done-predicates + m-5.** Phase-shaped base predicates + archetype-tag add-ons; both surface the SAME five mappings (extension = additive-only + zero-test-edit; refactor = suite-green + test-files-unchanged; cleanup = find-refs=0; bugfix = parent-red→fix-green differential; migration = reversibility) and AGREE these are surfaced-not-closed, requiring an explicit m-5 lock-time disposition.
- **Egress gate.** Fail-closed at the conductor boundary (not prompt policy); away-mode = first external send; scan secrets/PII/undisclosed-model-names/auth-URLs; promote the agent-scripts + claude-code rule sets.
- **Consumer contract + seams.** m-4 must consume observed atoms / `evidence_integrity`, never raw claims as benchmark truth (routing record may itself be evidenced — coord thread at DESIGN); m-5 owns the tag-space, m-3 owns the predicate-execution interface; m-6 away-mode must not publish before egress passes.

### DIFFERENT COVERAGE (complementary — folded into DESIGN, no conflict)
1. **Implementer formalized the API** — concrete `observe_gate(record_candidate, seat, phase, authority, evidence_target, artifact_refs, executable_claims, egress_targets, archetype_tags) → observe_result{...}` signature. I gave the same surface in prose. **ADOPT their signature as the DESIGN strawman.**
2. **Implementer added `degradation_notes` / bounded-observation** (agent-scripts degradation-honest discipline: report floors, truncation, backoff). I did not foreground it. **ADOPT** — it is the honest-floor discipline for the observe layer.
3. **Implementer's sharper operator framings** — executable claims restricted to *registry-approved check descriptors* (not arbitrary commands); egress *auto-redact-low-risk vs always-block-on-first-release*. **ADOPT** as the precise form of my operator-items G1/G2.
4. **Implementer's extra internal cites** strengthen the gap statement: jcode `await_members` status-driven done (`comm_await.rs:54-64`); jcode `send_message` no-scan path as negative egress prior art (`ambient.rs:1044-1112`); claude-code Stop-hook turn-blocking (`query.ts:1267-1289`); TaskUpdateTool "don't mark partial complete" (`prompt.ts:13-20`).
5. **My distinct coverage (planner-only):** the EXTERNAL literature — Reward Hacking Benchmark / RewardHackingAgents (agents tamper tests / hardcode expected artifacts; "long-horizon software work has no integrity instrumentation"; detection LLMs over-rely on trajectory self-awareness), VFA/Action-Attestation (post-hoc peer auditors, not a send-gate), RLVR/verifiable-keypoints (the executable-claim primitive as RL reward, not a governance gate). This is the *empirical* proof the gap is real AND that an in-reach verifier is gameable — the load-bearing justification for **outside-the-lane** observation and for the **verifier-tamper-resistant** m-5 invariants (refactor no-test-edits, bugfix red→green). It hardens, and does not contradict, the implementer's pass. (Per dispatch, the implementer lane has no deep-research skill — so this is expected complementary coverage, not a miss.)
6. **My distinct operator-item (G3):** the Gate-1 Step-1 read-vantage dependency — observe-as-send rides Step 1 where the conductor holds an outside-the-lane read handle; a fully-opaque remote lane degrades honestly to `self_reported` and full arbitrary-lane outside-vantage is the standalone-runtime carry-forward routed to the orchestrator. Non-blocking for v3.0. Retained.

### ONE DESIGN-TIME ITEM TO FLAG (not a disagreement, not an R3 reopen)
The implementer's `observe_result` includes `evidence_integrity: observed | self_reported | **mixed**`. The locked R3 per-field tag is two-value `{observed | self_reported}` (ARCHITECTURE.md:70-71) — **that stays as-is, not reopened.** A record-level `mixed` would be a *rollup* over per-field integrity (some fields observed, some self_reported in one record). Whether to carry a record-level rollup label, and its vocabulary, is an **m-3 DESIGN-internal detail** — flagged here so it is decided at DESIGN, with the per-field contract untouched.

### OPERATOR-JUDGMENT ITEMS (reconciled — none blocking AUDIT→DESIGN)
- **O1.** Executable-claim execution surface: registry-approved check descriptors vs arbitrary project commands → reserved seam, operator-gated at activation; v3.0 ships observe-as-send on passive E1/E2 without it.
- **O2.** Egress fail-closed policy: auto-redact-low-risk vs always-block-on-first-release; rule-set membership operator-configurable (mirrors §J2). Activation rides the first external send (m-6 away-mode, §J1).
- **O3 (planner-retained).** Step-1 read-vantage / opaque-lane honest-fallback / standalone-runtime carry-forward — confirm the `self_reported` labeling is an acceptable Step-1 floor.

### NEXT
Pair audit is complete and reconciled. Recommend the orchestrator sequence m-3 to DESIGN against this reconciled view, with: the implementer's `observe_gate()`/`observe_result{}` as strawman; the m-3↔m-4 evidence-record seam opened as the reserved `c2-*-coord` thread; and an explicit m-5 archetype-tag disposition required before any c2 m-3 lock. No source changes, no PR.

ACTIONS_GIT_REF: none — no git edit/commit/PR/migration; cwd (/Users/jack/Programming/harness) is a non-git docs workspace. Sole artifact this turn is this relay file: master/relays/c2-audit-m-3/RECONCILE-planner-20260629-185757.md

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
