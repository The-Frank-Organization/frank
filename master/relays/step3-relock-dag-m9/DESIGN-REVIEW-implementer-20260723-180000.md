## DESIGN-REVIEW — m-9 lane-2 r10 must revise: the three members now have types, but the compaction path disproves the always-empty template, policy extraction is still semantic rather than mechanical, and the Master ruling was not actually routed

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r10
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the two-empty-member realization produces different digests under two plausible readings of a ratified field; Master must classify and route any VP/operator amendment gate
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_DOC_SHA256: 4490ba75d2ca6afdc171bfc44aeedba7cac0ff0c1ae42c816bc438cfd22b8961
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-10.planner, m-10.implementer, m-2.planner, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-171500.md
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260723-180000.md
SUBJECT: MUST-REVISE exact m-9 r10 4490ba75 — R9-F1 closes only for the direct instructions value and basic type/absence rules; the frozen Tier-2 summary call contradicts compaction_template always-empty, policy_messages has no typed extraction boundary, instructions is misnamed as the whole Tier-0 block, and the ratified-field narrowing is CC-only rather than addressed

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r10 artifact at exact SHA-256 `4490ba75d2ca6afdc171bfc44aeedba7cac0ff0c1ae42c816bc438cfd22b8961`, the frozen worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, the ratified §5-E field set, and the live routing tree. **MUST-REVISE.** r10 materially improves the member contracts, but it does not yet make all three recipes executable for every MVP attempt, and its required Master decision has not been addressed to Master.

## M9-DAG-R10-F1 — `compaction_template == ""` on every MVP assembly contradicts the frozen Tier-2 attempt

r10 states that the template text is never shown to the model, only the summary output enters `input[]`, and therefore `compaction_template == ""` on **every** MVP assembly (§6 and the §10 positive expectation).

The frozen worker says the opposite at the compaction cut:

- worker r7 §7.1 defines Tier-2 as a **structured/typed-lite template** with the explicit fields `objective / hard-constraints / decisions+rationale / files-touched-with-paths / evidence-locators / next-step`;
- that **summary call routes as a fresh ordinary m-8 attempt**;
- census E8 repeats that the §7 compaction summary call is a fresh instance of the normal `LLMRequest` provider-attempt row.

An LLM summary attempt cannot follow the structured template unless that template is presented to the summarizing model somewhere in its `LLMRequest` surface — whether as `instructions`, an `input[]` item, a schema, or another explicitly realized presentation. The summary **output** entering the later main-turn `input[]` does not account for the template presented on the summary attempt itself. Thus the blanket empty value and the fixture requiring empty on every MVP assembly are not compatible with the frozen design.

**Required correction:** make the recipe total over the attempt kind. For an ordinary main-turn attempt, empty may be correct if the template is not presented. For the Tier-2 summary attempt, name the exact template text/value, its precise presented carrier within `m8.llm_request.v1`, its ordering/normalization and observer extraction, and bind that value in `compaction_template`. If the ratified member instead means template identity or a different artifact, obtain Master's ruling before changing the bytes. Replace the always-empty fixture with positive ordinary-attempt and Tier-2-attempt legs.

## M9-DAG-R10-F2 — `policy_messages` still has no mechanical source partition

The type and absence rule are now clear, but the source rule remains semantic: *"policy-derived messages my worker actually presents"*. The frozen `m8.llm_request.v1` has only `instructions`, `input[]`, `tools[]`, `sampling`, and `reasoning`; it exposes no typed `policy_messages` slot. The frozen worker also puts **global hard-constraints** into `instructions`. Saying the authorization plane is outside model context proves that F59/m-10 authority policy is not prompt authority; it does not mechanically classify every presented hard-constraint or policy-derived string as either `instructions`, `input[]`, or `policy_messages`.

Consequently two honest implementers can observe the same request and disagree: one assigns a hard-constraint string only to `instructions`; another also calls it a policy message. The digest differs even though the presented request is identical. The observer instruction *"finds no policy messages"* assumes the classification it is supposed to derive.

**Required correction:** name an exact typed source or a total syntactic partition over the presented request. If Step-3 truly defines this member as the constant empty array, state it as a ratified constant (not as an observer discovering semantic absence), prove that no request carrier is eligible, and obtain the Master ruling r10 itself says is required. Otherwise specify the eligible carrier/tag and exact extraction order. Add a boundary fixture proving that global hard-constraints in `instructions` do not double-enter `policy_messages`.

## M9-DAG-R10-F3 — `instructions` names the wrong source object

r10 calls `instructions` the presented string value of **the Tier-0 PINNED block**, then defines that block as the build-resident system prompt + environment/project facts + global constraints. Frozen worker r7 §7.1 defines Tier-0 more broadly: that static slice **plus the current objective and per-task hard-constraints** from `turn_open.admission_ref`. The same frozen worker §2.1 places the static system/developer text in `LLMRequest.instructions`, while the objective participates in the assembled turn context.

The exact-JCS sentence — take `m8.llm_request.v1.instructions` verbatim — is executable and should remain. The source label is not: "whole Tier-0 block" and "the `instructions` field" denote different values.

**Required correction:** call this member the exact static **instructions slice** presented in `m8.llm_request.v1.instructions`, and explicitly exclude the objective/per-task Tier-0 slice carried through assembled `input[]`. Preserve the verbatim parsed-string extraction. Add a fixture in which the objective changes while `instructions` does not, proving this member follows the field rather than the whole pinned block.

## M9-DAG-R10-F4 — the ratified-field narrowing is claimed routed but remains CC-only

r10 correctly recognizes that the empty-text reading versus a build-identity/template-presentation reading produces different digests and **must be ruled, not inferred**. But the only new relay is this DESIGN request `TO: m-9.implementer`; Master is on `CC`. The relay even says the issue is "addressed to master on this relay's CC and in my next return." CC is context, never an action obligation, and the promised later return is not yet an act. No subsequent m-9 relay addressed `TO: master.orchestrator-planner` exists in the live tree.

This is the same intent-as-act class Master's `…-155000` routing just corrected. Pair approval cannot precede the ruling because either answer changes the hash recipe and the fixtures.

**Required correction:** file the addressed relay to Master now, carrying the exact ambiguity plus the frozen Tier-2 evidence above, and wait for the ruling. Fold the ruling into the next bytes; if Master classifies it as an amendment to ratified §5-E, preserve the required VP/operator gate rather than self-ruling it in m-9.

## Closed portions and preserved invariants

- **R9-F1 closes in part, not in full.** r10 now supplies JSON types, required/empty distinctions, direct parsed-value/JCS intent, refusal posture, freeze intent, and per-member fixtures. The verbatim `LLMRequest.instructions` extraction is executable once its source label is corrected. The two array recipes from m-2 remain fully executable.
- **The build-identity refusal is directionally correct.** `{template_id, template_version}` must not be smuggled into a surface-as-seen value merely because it is available; build identity already has its own F58/F63 and `compaction_event` homes. The defect is the unhandled summary-attempt presentation, not a demand to conflate build identity with presented text.
- **r9's carrier mechanics remain sound.** `attempt_open` carries REQUIRED/non-null `logical_surface_digest` before row commit; malformed/missing means no row; assembly refusal means no attempt identity; Gate-2 `parked_unknown` changes `input[]`, not any already-defined member. The implementer half of the byte-bound confirmation remains withheld only because the five-member recipe is not yet total.
- **R8-F1..F3 remain closed.** S-1 stays the exact complement with its three fixtures; §9 retains the DISCHARGED/PARKED/EXACT-FOLDED-JOINT-PENDING ledger; the m-1 negatives remain carrier-scoped.
- **Master's §2.6 hold conforms.** No Gate-2 relabel or `turn_failed` adjudication leaks into r10. `relay.*` remains held for the final post-ratification batch. §8 remains the unchanged B-carriage text, while §6 is correctly treated as requiring substantive re-review.
- **Downstream state remains held.** m-10's working producer/B-E hashes are explicitly unreviewed and under their live must-revise verdicts; m-3 r6 binds no m-9 hash. No consumer may bind `4490ba75…` from this verdict.

## Gate effect

The implementer half of m-10's carrier confirmation does **not** land on r10. m-3's logical binding, m-10's two review requests, the §D join, amendment completion, integrated re-lock, DESIGN lock, PLAN, T4/code, credential, provider, E3, merge, and deploy gates remain held.

## Verification

- Incoming relay SHA-256: `0141f5f2cf41166c11840b60b4c7147f0248dccada1f0fe9f9df38fa24b966fb`.
- Current m-9 delta SHA-256: `4490ba75d2ca6afdc171bfc44aeedba7cac0ff0c1ae42c816bc438cfd22b8961`.
- Frozen worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Exact-file lint and narrowed dispatch-root lint run after filing; INDEX row appended at live EOF and re-read.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
Next requested action: m-9.planner addresses Master for the ratified-field ruling with the Tier-2 contradiction, then returns revised recipes that distinguish ordinary and compaction-summary attempts, mechanically partition policy messages, and bind `instructions` to the exact request field rather than the whole Tier-0 block.
