## RECONCILE — STAGE-6 RE-SCOPE: the stage-5.1 review lands a strong critique; master HOLDS the all-artifact lock and proposes a bounded five-item amendment + a lock-granularity re-cut → VP adversarial review of the decomposition, then the operator's re-scope gate

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this is a milestone re-scope; it needs the VP's decomposition review then the operator's explicit re-scope ratification. The pending joint lock record `b7e1f0ef` is HELD and superseded by this path; it does not issue.
GRILL_REQUIRED: no — the operator already grilled Option-1 (sandbox in, keep the "Frank harness MVP" label) in-session 2026-07-21; this relay routes the decomposition, it does not re-open a settled product choice
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-022100.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r4 APPROVE stands on the bytes, but the operator commissioned a stage-5.1 external review that faults the milestone CLAIM + lock SCOPE (not the frozen bytes' consistency); master concurs on the load-bearing subset and proposes a bounded amendment — review the decomposition before the operator's re-scope gate

VERDICT: revise — self-initiated: master withdraws the all-artifact lock from the operator gate and requests the VP's adversarial review of this bounded amendment's decomposition

## 1. What changed since your r4 APPROVE
The operator commissioned a stage-5.1 third-party review (chat model + web-research harness) against `frank-dev main@c78da38`, stored verbatim at `master/STAGE-5.1-EXTERNAL-REVIEW-2026-07-21.md` (`b4e79f3bfdc4ca0b2f5023b441da1e0006c7064a5d4b6b06bb4e9e2a7b5bf7d1`). It does NOT fault the frozen bytes' internal consistency (your r4 checks stand); it faults the **milestone claim + lock scope**: strong governance kernel, not yet honestly a coding-agent MVP — containment (ungoverned `bash`), exact-effect binding, working-state durability, and task-level proof are deferred. The operator chose **Option-1**: keep the "Frank harness MVP" label, pull the sandbox IN. Master's full disposition (accept / partial / push-back per finding) is in the amendment §1.

## 2. Master's proposal — a bounded amendment, not a ratify and not a full re-open
`master/STEP-3-STAGE6-AMENDMENT.md` (`0634b6e42e324ab3fd8858bfcf0c58105be12f8662b8f24e03a9b0a7d0dd6985`), PROPOSED. It:
- **HOLDS** the joint lock record `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-023000.md` (`b7e1f0efa…` — see manifest) and supersedes its pending operator gate;
- does **not** reopen the ratified architecture direction (amendment §2 — conductor isolation, m-10-no-policy, m-8-creds, replaceable m-9, separated stores, first-class UNKNOWN, one lane/turn, fixed catalog, native+MCP parity, H-17-in-lock all STAND);
- **re-cuts the lock into Tier-HARD (wire/ids/encodings/effect-schemas/census/m-1·m-2·m-3 contracts — full-F73 re-open) vs Tier-SOFT (prompt wording, tool descriptions, compaction tuning, local-tool ergonomics, retry constants, user surface — revisionable)** (item A, §3);
- adds four Tier-HARD interface obligations (§4): **B** `frozen_core_digest` joined m-8→m-10→E0/E3→exit-proof; **C** F59 promoted from `canonical_args_digest` to an **effect descriptor** incl execution context; **D** a typed model-visible **run journal** (a projection, no second truth; settled-round resume); **E** the **sandbox** three-backend seam (Linux-strong reference / macOS-Seatbelt honest-weaker / Linux-in-VM parity; containment proof binds to Linux; `backend_id`+`network_policy_id` in the descriptor + one census row per backend) + one fixed local policy + `model_surface_digest` + typed E3 predicate ids;
- routes the non-interface items to T4/exit (§5 — detailed tool semantics as Tier-SOFT build-with-tests; the revised five-proof exit test) and the broker question to an m-7 study with a conditional H-24 pre-T4 gate (§6).

## 3. Where master did NOT simply defer to the review (for your adversarial attention)
- **Item 4 (broker) is a PARTIAL accept, not wholesale.** Keep the separate secret-holding process (F67). The review implies the epoch machinery mostly collapses if the broker dies with app-main — master's push-back: worker replacement advances the epoch on **every crash-loop retry**, so the F64 fence + linearization stay regardless; only *survive-app-main-crash + adoption + cross-epoch completion* is the questionable part. Confirm or break that framing.
- **Item 6 (product eval) is step-close, not a design-lock blocker** — repo tasks can't run before the T4 build exists. Routed to the exit test, not the interface freeze. Confirm.
- **Item 5's detailed tool semantics are Tier-SOFT**, not Tier-HARD. Confirm the tier assignment doesn't leak a real interface (e.g. does any tool-behavior detail bind a durable record shape? if so it's Tier-HARD).

## 4. The proposed re-scope sequence (§8 of the amendment) for your review
amendment → **your decomposition review** → **operator re-scope ratification** (supersedes `b7e1f0ef`) → owning-pair Tier-HARD deltas in the acyclic order `m-1+m-8 → m-10 → m-3 → m-9 → m-2` (∥ the m-7 broker study) each via the F73 ladder → **re-run the shorter stage-6 lock** over only the changed Tier-HARD interfaces (unchanged design hashes carry forward) → T4 (behind the re-lock + H-16/H-26) → the five-proof exit test → step close.

## 5. Requested return
Your adversarial review of THIS decomposition: (a) are B–E correctly scoped Tier-HARD and A's tier split sound; (b) is the §8 acyclic order correct against the frozen seams; (c) is the item-4 partial-simplification framing right; (d) does anything here silently reopen an approved mechanism or move a bound design byte? On your pass, the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, or deploy is requested.

## Verification
The review, amendment, and backlog battle report recomputed from disk: review `b4e79f3b…`, amendment `0634b6e4…`, backlog `5b0958f8ac6eb31e635842ba674e0ebef843da3ab754b8401b26d217058371de`. The nine design finals + H-16 rev16 + census `959b1928…` are UNMOVED by this amendment (it adds obligations + re-tiers the lock; it withdraws no approved mechanism). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof); `frank/` untouched by this authoring (the c78da38 vendoring is the reviewer's basis, unchanged).

ACTIONS_GIT_REF: docs-workspace disk action — this relay + the amendment doc + the stored review + one backlog battle report + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657` locally / `c78da38` on origin (the vendored reviewer basis).
Next requested action: the VP returns the decomposition review; on pass master routes the amendment to the operator for the re-scope gate.
