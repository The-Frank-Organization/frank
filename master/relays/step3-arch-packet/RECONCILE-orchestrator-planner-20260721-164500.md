## RECONCILE — OPERATOR RE-SCOPE GATE: stage-6 amendment rev12 is VP-APPROVED at decomposition grain; ratify the exact bytes `1125b0a0…` to re-scope the MVP (sandbox forgone + H-12 hard-blocked · durable-resume built · six property legs + overhead budget · deployment envelope)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this is YOUR re-scope ratification decision. Nothing downstream issues until you act; this relay does not self-satisfy the gate.
GRILL_REQUIRED: no — the §3 GRILL_LOCK is already recorded from your 2026-07-21 grill; this gate ratifies it + the §7 numbers + the §10 envelope, it does not re-open the grill
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-163600.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: the twelve-round decomposition loop has converged — VP APPROVE r12 over rev12 `1125b0a0…`; this is the exact byte set for your re-scope ratification, and what ratifying opens

VERDICT: routed — master routes the VP-approved amendment to you for the re-scope decision (RATIFY / REVISE)

## 1. What you are deciding
The stage-5.1 external review said the frozen kernel was real governance but not yet honestly a coding-agent MVP (ungoverned bash, no exact-effect binding, context lost on worker replacement, plumbing-only exit test). You chose to keep the "frank harness MVP" label and pull the scope up to match. `master/STEP-3-STAGE6-AMENDMENT.md` **rev12 (`1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`)** is the result, now **VP-APPROVED at decomposition grain** (r12, `RECONCILE-orchestrator-reviewer-20260721-163600.md`) after twelve review rounds. Ratifying these exact bytes:
- **supersedes** the held joint interface-lock record `b7e1f0ef` and its pending all-artifact gate (it does not issue — it is replaced by this re-scoped plan + a later, shorter re-lock);
- ratifies the **§3 GRILL_LOCK** (your 2026-07-21 decisions), the **§7 overhead budget numbers**, and the **§10 deployment envelope**;
- authorizes the **§11 sequence** to begin (m-7 broker study first, then the interface DAG legs, then the recipe/bundle, then the shorter re-lock, then T4) — each still separately gated.

It does **NOT** issue any DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy. Those stay held behind their own gates.

## 2. The re-scope, in plain terms (the §3 grill you are ratifying)
- **D1 — sandbox FORGONE for the MVP.** Bash stays ambient (you run `--dangerously-skip-permissions` anyway; the 8-harness survey showed sandboxing is commodity, not the moat). The gap is documented, and **H-12 is promoted to a HARD pre-external-use blocker**: no untrusted / external / security-sensitive / multi-tenant use until a real sandbox lands (§10). The moat stays the governed evidence plane.
- **D2 — the bash claim is narrowed** to binding the exact invocation *context* (not affected-resources or per-effect holds).
- **D3 — the exit GATE = six governance-property legs + an objective overhead budget**, not a benchmark score: Governance, Durability, Crash-honesty, Injection-**visibility** (not prevention), Handoff, Operability.
- **D4 — utility is DEMONSTRATED, not gated:** public dogfood (the CRM — informal ask, you retain all rights — + bivpak, as open auditable testaments) + an honestly-labeled agent-as-operator SWE-bench run with **no threshold gate** (avoids the benchmax incentive).
- **D5 — decoupling invariant:** real-work / dogfood start is ⊥ the MVP-exit gate. You can begin the CRM/bivpak work without waiting on the exit test.
- **D6 — the effect descriptor is context-binding** (`backend_id="ambient"`, no containment claim).
- **D7 — durable session state + resume is BUILT** (not deferred). This is the item I under-called as "just a JSON log" — the review and the VP progressively showed a sound resume is a genuine crash-safe subsystem, because a best-effort resume would fabricate settled effects, which is exactly the failure class frank exists to kill. It is a **worker-owned session-content log** (field-standard: Codex rollout+replay, Claude Code transcript+`--resume`, opencode store+snapshot, deepagents checkpointer) with positive settlement reconciliation and a continuation-turn lifecycle; **outcomes stay m-10-canonical**. The twelve rounds hardened this to: a producer-total three-class manifest, a **two-time-scoped trust invariant** (content is trusted only under settlement evidence AND its presence in the current recovered prefix — else `content_lost`/degrade, never fabricated), and a single fail-closed terminal on an un-emittable resume frame.

## 3. The structural amendment (items A–E) + re-sequence
- **A** — the interface lock is re-cut into a **hashable Tier-HARD bundle** (a mechanically-extracted canonical digest via a versioned recipe → one `bundle_sha256`, stable under Tier-SOFT edits), not prose labels.
- **B** — a `frozen_core_digest` join binds the frozen kernel.
- **C** — the effect descriptor as context-binding evidence (D6).
- **D** — the durable-resume subsystem (D7).
- **E** — `model_surface_digest` + typed E3 predicates.
- The cyclic pair-order is replaced by an **acyclic DAG** (§6), and the **m-7 broker study resolves first**.

## 4. The numbers you are ratifying (§7 overhead budget)
F59 authorize p95 ≤ 250 ms · relay round-trip ≤ 1 s · journal-commit ≤ 100 ms · per-turn added wall-clock **p50 ≤ 20% (PASS) / 20–100% (HOLD, operator-waivable) / > 100% (FAIL)**, hard-fail > 2×. These are operator-ratifiable — if any number is wrong for how you actually run frank, say so and I revise before you ratify.

## 5. What I do on your RATIFY
Per [[roadmap-maintenance-rule]] and the §8b operator-authorship rule (recorded agent-authored + operator-cited, never a forged `FROM: operator`):
1. Stamp the re-scope + the lock re-cut basis into `master/ARCHITECTURE.md` + `master/README.md` + `ROADMAP.md`.
2. Open the **§11 sequence** as separately-gated lanes: (a) the **m-7 broker study** (+ H-24 if cross-epoch completion survives); (b) the **interface DAG legs** under the F73 ladder with join records for the two-sided seams; (c) **author the extraction recipe + bundle** + the `bundle-soft-stability` negative fixture + freeze `STEP-3-EXIT-FIXTURES.json`; (d) the **shorter stage-6 re-lock** over `bundle_sha256` + the whole-file-hard owner contracts; (e) **T4** behind the re-lock + H-16/H-26.
3. Because of D5, you may start the CRM/bivpak dogfood in parallel whenever you want — it is not gated on the exit test.

## 6. Your decision
- **RATIFY** — reply approving the exact rev12 SHA-256 `1125b0a0…` (the byte-exact citation is what the §8b record binds). I then execute §5 above.
- **REVISE** — name what to change (a §7 number, an envelope line, a grill decision); I fold it, re-run the shorter VP check, and return a fresh gate. Any byte change voids the VP approval and needs a fresh review.

## Verification
Recomputed from disk: amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` (UNMOVED since VP APPROVE r12 `163600`); VP approval is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (the amendment carries PROPOSED owner-delta acceptance properties, no byte edit to any frozen doc). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no amendment/design/source byte moved, no `frank/` action, no lock issued, no gate self-satisfied (the operator gate is yours to act).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: operator returns RATIFY (approving exact `1125b0a0…`) or REVISE; on RATIFY master executes §5 and opens the §11 sequence; all other action stays held.
