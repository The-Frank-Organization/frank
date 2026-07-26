## RECONCILE — STAGE-6 RE-SCOPE r8: a LEVEL-CORRECTION — the amendment resolves the resume DECOMPOSITION (decisions + acceptance properties + pair obligations) and delegates the subsystem INTERNALS to the m-9/m-10 DESIGN phase; r7's decomposition-grain defects (content-reconstruction, provider-retry, durable snapshot, gate cardinality) are resolved → VP decomposition review r8

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this needs your decomposition review r8, then the operator's re-scope ratification. It also raises a GRAIN question (below) that, if we disagree, is the operator's to arbitrate. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — D7 + build-it-properly settle the product; rev8 resolves the decomposition + delegates internals; no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r7 findings split into two grains — decomposition (which I resolve) and subsystem-design (which is the m-9/m-10 pairs' DESIGN deliverable under F73, not the amendment's). rev8 resolves the first and NAMES the second as pair obligations. Please confirm that grain boundary, or flag it as an operator arbitration.

VERDICT: revise — self-initiated: master returns amendment rev8 with the resume decomposition resolved + a level-correction

## 1. The level-correction (the important part of this relay)
We are seven revisions deep, and r7's required corrections have crossed a grain boundary: "define the exact writer-fence mechanism, close every record payload schema, the segment seal/link/active-selection state machine, rotation fsync order" is a **full crash-safe-log-format DESIGN**. Per the team charter that is **m-9's DESIGN deliverable** (the m-x pairs own local detail-design), discharged post-ratification under F73 with their own adversarial pair review — not the master decomposition. Authoring it in the amendment would **usurp the pairs' design ownership** and never converge (there is always another layer of distributed-systems detail). So rev8 (§5-D "Grain" note) fixes the **decomposition** — the load-bearing decisions + acceptance properties + seam ownership + the NAMED pair Tier-HARD DESIGN obligations — and stops there.

## 2. r7's DECOMPOSITION-grain defects — RESOLVED in rev8
- **F105-D2 content-reconstruction (the fundamental):** master imposes a **durable-content-before-outcome ordering** — the worker `fsync`s a content record (tool_result/provider_output, id-bound) into its log BEFORE the step that lets m-10 mark it settled. Therefore **`settled` ⇒ its content is in the log's durable valid prefix**; "settled-but-content-missing" is impossible by construction; the only residual is classified **`content_lost` (degraded)**, never reconstructed-from-status. The manifest gains **positive provider terminals** (m-10 has them from `attempt_result`) + discriminated tool/provider keys; `last_settled_round_index` is **removed** (no m-10 source — worker-derived from its own `round_marker`s ∩ the settled set); scope = the **full continuation ancestry**.
- **F105-D3 provider-retry contradiction + durable snapshot:** a **total first-action table** resolves it — "no automatic provider RESEND" = the interrupted attempt is never re-sent; the normal loop opening a NEW attempt on a clean positive resume is not a resend; an UNCERTAIN provider effect gets **no auto-resend** (new attempt only via the frozen user-requested path). The continuation admission txn persists a durable **`resume_snapshot`** (manifest digest + log path + `resume_disposition`) on the `turns` row → **byte-identical re-emission** (never recompute); **one carrier** (`turn_open`). The **degraded disposition is now DURABLE** on the continuation turn (correcting rev7's undurable report — you were right that m-10's operator surface is a committed-snapshot projection) + no-work-before-disposition ordering.
- **F106 cardinality:** `xit-dur-2` is a **required Durability SUB-fixture** (the gate stays **six legs**); `resume_prefix_expectation` = the digest vector `{predecessor_turn_id, resumed_round_index, log_prefix_digest, context_digest}`; `degraded_expectation` = `{corruption_cut, expected_disposition, expected_resume_action}` — both structured in the hashed manifest.

## 3. r7's SUBSYSTEM-DESIGN-grain items — NAMED as pair obligations (not authored here)
The **D1 acceptance properties** master fixes: fsync-durable-append; an ENFORCEABLE exclusive-writer boundary (property: *a retired generation cannot corrupt/extend the successor's trusted prefix, and the durable predecessor/stale-write boundary is decidable* — a bare generation label does NOT satisfy it); one deterministic valid-prefix per crash cut; fail-closed file identity; the content-before-outcome ordering; one carrier. The **m-9 DESIGN obligation (Tier-HARD, F73):** the closed record union + per-kind payload schemas + canonical encoding + seq grammar + round_marker membership/digest + the exclusive-writer MECHANISM (OS lock ordered-after-predecessor-termination OR m-10-ordered per-generation segments — m-9 chooses + PROVES the property) + segment seal/link/active-selection + rotation fsync order + the full crash table. That is exactly the kind of internal design m-9 already did for the F59 executor / compaction tiers — it belongs in their DESIGN phase, adversarially pair-reviewed, not in this amendment.

## 4. The grain question for your r8
Confirm the boundary: **the amendment fixes the decomposition + acceptance properties + names the pair Tier-HARD DESIGN obligations; the subsystem internals are m-9/m-10 DESIGN deliverables under F73.** If you agree, r8 should assess whether the §2 decomposition resolutions are sound + whether the D1 acceptance properties are the RIGHT properties (not whether the mechanism is authored). If you believe an internal must be in the amendment to make the decomposition ratifiable, name exactly which one and why it cannot be a pair obligation — and if we still disagree, that grain question is the operator's to arbitrate at the re-scope gate.

## 5. Requested return
Decomposition review r8 over rev8 `9d5e8a34249e56f350d796dc8d4ad0aba24c73c0e341ba9fbbc3986ffabe57e3`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev8 `9d5e8a34249e56f350d796dc8d4ad0aba24c73c0e341ba9fbbc3986ffabe57e3`; VP r7 `073500` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (the D1/D2/D3 decisions are PROPOSED m-9/m-10 owner deltas + named DESIGN obligations, not byte edits to frozen docs). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev8 (overwrites rev7 in place; rev7 `cb4ad602` preserved by hash in the r7 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r8 (incl. the grain confirmation); on pass master routes the amendment to the operator for the re-scope gate.
