## RECONCILE — c2 audit cross-pair reconciliation (orchestrator → VP): both domains GO, propose PROCEED-TO-DESIGN

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c2-reconcile
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — proposes PROCEED-TO-DESIGN; operator-judgment items are DESIGN-time, none blocking the AUDIT→DESIGN transition
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — both c2 pairs returned. I read all four independent passes (`c2-audit-m-3/AUDIT-{planner,implementer}`, `c2-audit-m-4/AUDIT-{planner,implementer}`) + both pair reconciles, and verified INDEX completeness (0 missing / 0 dangling; fixed one misplaced row + one dup). Both pairs genuinely converged — no hidden planner/implementer split: m-3 = AGREE, zero disagreements (two blind passes landed identically); m-4 = reconcile ledger with zero genuine disagreements, differences complementary and merged toward evidence. Both recommend PROCEED-TO-DESIGN with the two seam conditions we set. Below is the cross-pair integration view (my job, beyond the per-pair reconciles) + the proposed disposition.

**F1 — Both domains GO; verdicts E1-sound.**
- m-3 = **recommended-next**: PROMOTE the E0–E4 ladder vocabulary + the agent-scripts/claude-code egress rule-engine (fail-closed redact-then-omit) + the jcode/claude-code hook-point pattern + the ODB evidence template; **BUILD** the outside-the-lane observe-AS-send gate + `evidence_integrity` stamp + archetype done-predicates + the egress chokepoint placement. Not a rebuild — a promote-and-wire-the-parts, build-the-integration.
- m-4 = **still-open primitive / recommended-next**: PROMOTE jcode `MultiProvider` runtime + published priors (Step-3 wire-not-rebuild); **BUILD** the governance-record layer now — the `routing_decision` record (m-2 FieldSpec consumer), the declared+versioned+snapshotted capability-prior table, the 3-staged policy (prior floor → justified deviation → [a later release] outcome feedback), and the thin **fail-closed** router API (`route_dispatch()` → `human_decision_required`/`routing_unavailable`, never a silent default fallback).

**F2 — The m-3↔m-4 seam is mutually named: writer matches reader** (the c1 anti-pattern — writer-with-no-reader — is avoided). m-4's `outcome_feedback_ref` (reader) binds to m-3's `observed_evidence_ref` / `achieved_evidence` / `evidence_integrity` (writer). Both independently agree: the routing record is itself an evidenced record; the m-3 hook is observer-only (locked R3 allowlist); **m-4 declares which routing fields are observed, m-3 owns how**; and `self_reported` evidence is not a clean benchmark signal. The concrete first shared surface is the **`deviated`-vs-`capability_prior_snapshot` observation** (m-4's R2-CRITICAL declare/observe construction is, on m-3's side, an observe-predicate over an m-4 record). Two questions are correctly deferred to the single `c2-*-coord` COORD thread: (a) is a routing-record observe-atom the same type as a relay observe-atom, or a profile of it; (b) is `observed_evidence_ref` required on all routing decisions, only benchmark samples, or only deviations. Clean seam — resolve in COORD before any c2 lock.

**F3 — The m-5 seam is ONE shared disposition, not two.** Both domains reach m-5 through the *same* locked `slot_in` reserved opaque atom (m-2): m-3 → archetype tags parameterize the done-predicate (5 candidate mappings; 2 are verifier-tamper-resistant — refactor no-test-edits, bugfix red→green); m-4 → archetype tags parameterize the capability-prior lookup, and authority-ceiling-at-spawn caps what the router may assign (identity≠authority realized at the routing layer — the router refuses to staff a seat the ceiling forbids). In both, m-5 owns the tag-space and m-3/m-4 own the consuming mechanism. So the lock-time m-5 disposition you required (Finding 4) is a **single coordinated act covering both domains' tag-consumption** — satisfiable by one narrow m-5 review on both draft designs, or one reconcile disposition reserving the seam in both. I recommend we plan it as one disposition at the c2 lock.

**F4 — Locked c1 contract preserved; no foundation reopen.** m-4: R2 preserved *by construction* — the `required_when` predicate atom is the agent-declared `deviated` boolean, never a `model_*` comparison, so no model predicate enters the schema gate; identity≠authority is realized as the authority-ceiling cap; routing stays a category-B gate (model stays payload). m-3: the R3 observer-only allowlist is respected. The single flag is m-3's record-level `evidence_integrity: mixed` rollup — explicitly **not** an R3 reopen (the per-field tag stays two-value {observed | self_reported}); a record-level rollup over per-field integrity is an m-3-DESIGN-internal detail to settle at DESIGN. I concur it does not reopen the lock.

**F5 — Honest qualifications carried into DESIGN (no overclaim).** m-4 brought two: (a) Routesplain / Arch-Router already make routing rationale interpretable → the differentiator is the seat-stamped, persisted **deviation-against-a-declared-floor** audit artifact, not "interpretable routing"; (b) non-gradient bandit outcome-feedback already exists → the later-release novelty is the persisted, auditable decision+update artifact, not the non-gradient mechanism (the Fugu-reward-analog framing must be hedged). m-3 brought the reward-hacking-benchmark evidence (long-horizon agent work has no integrity instrumentation; an in-reach verifier is gameable) — the empirical justification that **outside-the-lane** observation is load-bearing, not cosmetic. These carry into both DESIGN docs as stated qualifications.

**F6 — Step-1 ride / Step-3 execute split is clean in both; no execution dependency blocks the DESIGN lock.** m-3 observe-as-send rides Step-1 wherever the conductor holds an outside-the-lane read handle; a fully-opaque remote lane degrades honestly to `self_reported` (locked DI-5 fallback); arbitrary-lane outside-vantage is the standalone-runtime carry-forward. m-4's routing record is writable on existing runtimes now; router execution is Step-3 (provider adapters). Both consistent with `ROADMAP.md`.

**Operator-judgment items (carried forward; none blocking AUDIT→DESIGN).** m-3: executable-claim execution surface (registry-approved descriptors vs arbitrary commands; reserved seam, ships without for now); egress fail-closed policy (auto-redact-low-risk vs always-block-on-first-release; config); Step-1 read-vantage honest-fallback floor. m-4: capability-prior seed values (ship a default seed, operator confirms/customizes); record scope = human-launched lanes now, execution Step-3. Most mirror the already-ratified §J config-customizability pattern. Recommend they ride to DESIGN as operator items, decided at/with the c2 design-lock — not now.

**Proposed disposition — PROCEED-TO-DESIGN for m-3 + m-4 (co-design), GRILL_REQUIRED: yes, with:**
1. one shared `c2-*-coord` COORD thread owning the m-3↔m-4 seam (the `deviated`-vs-snapshot surface + the two deferred questions), reconciled before any c2 lock;
2. a single m-5 lock-time seam disposition covering both domains (surface, do not close concrete archetype/tag semantics);
3. the F5 honest qualifications carried into both DESIGN docs (no overclaim);
4. operator items ride forward; none blocking.

No PROCEED-TO-PLAN / IMPL / merge / live-verify is sought. Requested: review this reconciliation and confirm PROCEED-TO-DESIGN (or revise), with cited changes. On your approve I dispatch the c2 DESIGN to both pairs (grilled) and open the `c2-*-coord` thread.

ACTIONS_GIT_REF: read the four c2 audit passes + both pair reconciles; verified + repaired INDEX (moved one misplaced row, removed one duplicate); wrote this relay. Docs-workspace artifacts only; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP returns PROCEED-TO-DESIGN confirm (or revise); on approve I dispatch c2 DESIGN + open the COORD thread.
