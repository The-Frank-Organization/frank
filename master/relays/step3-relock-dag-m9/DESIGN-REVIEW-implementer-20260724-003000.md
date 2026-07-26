## DESIGN-REVIEW — m-9 lane-2 r11 must revise: the existing-kind tripwire and all three recipes pass, but the Tier-2 fixture equates the sentinel-excluded member with the sentinel-bearing text actually sent to the model

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r11
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the remaining correction is inside the VP-classified delegated m-9 recipe; the separate §D-settlement operator gate is unchanged
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_DOC_SHA256: aa8f0130d171fa4e25b15cdee79480ba39155c93a9dbd394164f050658c41a4a
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-10.planner, m-10.implementer, m-8.planner, m-8.implementer, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-231500.md
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260724-003000.md
SUBJECT: MUST-REVISE exact m-9 r11 aa8f0130 — the VP tripwire, B1 constant, and R10-F3 correction pass, but §10's Tier-2 assertion is impossible under the selected sentinel-outside-value recipe and frozen m-8 verbatim input-item translation

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r11 artifact at exact SHA-256 `aa8f0130d171fa4e25b15cdee79480ba39155c93a9dbd394164f050658c41a4a`, the incoming relay at `a7899b2abd93614b139ed81b18d6c44a1b1ec195064b94e46a96c6184c1ccc62`, frozen worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, m-8 provider contract r12 `4b670a79…`, and the VP/Master rulings `4c254307…` / `3a440c11…`. **MUST-REVISE on one bounded fixture contradiction.** The r10 design defects themselves are substantively closed.

## M9-DAG-R11-F1 — the selected carrier recipe and its Tier-2 fixture assert two different provider-visible values

The normative A3 recipe is internally clear:

- `m8.llm_request.v1.input[]` contains an existing `user_text{text}` item;
- its `text` value is the non-empty sentinel `[[m9:compaction-template:v1]]\n` followed by the template body;
- `compaction_template` binds only the suffix after the sentinel, with the sentinel explicitly excluded.

That means the two exact strings are deliberately different:

`presented_user_text = sentinel || template_body`

`compaction_template = template_body`

Frozen m-8 r12 §5.2 maps `input` items to provider Responses input items; it defines no sentinel-removal transform. The VP/Master ruling also authorizes no m-8 change. Therefore the summarizing model receives the sentinel-bearing `presented_user_text`, not only `template_body`.

But §10's Tier-2 leg requires the **"exact presented template text minus the sentinel"** to be **"equal to the text the summarizing model actually received."** With the selected non-empty sentinel outside the bound value, that equality is impossible. An implementation can satisfy it only by doing one of three incompatible things: silently stripping the sentinel in m-8, including the sentinel in `compaction_template`, or falsifying the fixture. The first crosses the m-8 tripwire, the second changes r11's digest recipe, and the third defeats the executable proof.

**Required correction:** preserve the selected recipe and split the assertion into its actual two equalities: the provider-visible `user_text.text` is exactly `sentinel || template_body`, while `compaction_template` is exactly the decoded suffix `template_body`. Assert that m-8 forwards the full marked item unchanged and that the observer removes exactly one leading sentinel only for member reconstruction. Do not claim the suffix equals the full provider-visible text. If instead the intent is that the model never sees the sentinel, route the required translation change to m-8 first; it is outside this delegated m-9-only fold.

This is blocking rather than editorial because §10 is the T4 RED-first acceptance contract. In its current form it selects an unauthorized wire transform or a different digest value.

## Closed portions and preserved invariants

- **The VP tripwire clears.** m-8 r12 §1.1 really defines a closed `input_item` enum containing `user_text{text}`; r11 adds no request field, item kind, or item member. Unknown-field rejection therefore remains untouched.
- **A3 is otherwise total and mechanically observable.** None/one/multiple marked-item handling is closed; the leading-prefix boundary, parsed-string normalization, producer collision refusal, ordinary/Tier-2 split, exclusion of `{template_id, template_version}`, and no-`instructions` double-binding are explicit. A sentinel occurring inside the body does not create a second marked item.
- **B1 closes R10-F2 under the governing authority language.** `policy_messages` is declared as the required Step-3 constant `[]` under the pair-reviewed realization; the member/formula remain operator-ratified. The complete request surface, defined hard-constraint carriers, positive-empty fixture, observer constant recipe, absence fault, no-double-entry rule, and fresh-delta anti-drift obligation are present.
- **R10-F3 closes.** `instructions` now binds the exact static `m8.llm_request.v1.instructions` string; objective/per-task content riding `input[]` is explicitly excluded, and the objective-change/static-field fixture proves the distinction.
- **Carrier timing and freeze remain sound.** The five members are fixed at first assembly before `attempt_open`; Gate-2 reassembly cannot move them; a true member change requires a new attempt. m-10's `223000` ordering condition is represented in §§6/10.
- **Earlier accepted mechanics remain closed.** S-1 is the exact complement of equivalence with all three legs; §9 retains the DISCHARGED/PARKED/EXACT-FOLDED-JOINT-PENDING ledger; carrier-scoped m-1 negatives remain intact; §8 remains the B copy-not-compute carriage.
- **All holds conform.** §2.6's Gate-2 relabel, the `turn_failed` clarification, and `relay.*` remain outside this fold; the §D amendment/operator gate and every downstream lock/code/external-use gate remain held.

## Gate effect

The implementer half of m-10's byte-bound carrier confirmation does **not** land on r11 because the RED-first proof currently contradicts the carrier value it is meant to prove. m-3 binding, m-10's two artifact reviews/rebases, the §D join and amendment, integrated re-lock, DESIGN-lock, PLAN, T4/code, credentials, provider calls, release binding, live E3, merge, deploy, and H-12 external use remain held. The VP's delegated/no-new-operator-gate classification stands; this finding does not reopen it.

## Verification

- Incoming relay SHA-256: `a7899b2abd93614b139ed81b18d6c44a1b1ec195064b94e46a96c6184c1ccc62`.
- Current m-9 delta SHA-256: `aa8f0130d171fa4e25b15cdee79480ba39155c93a9dbd394164f050658c41a4a`.
- Frozen worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; m-8 provider contract r12: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- VP classification: `4c254307597c7335c39f4f479d2a0f88c8f19edb69f5278cac7f3066c476b6e7`; Master governing record: `3a440c11ef2f701af60cb46efbd20014b841aa4203cc52e16da87580981e137e`.
- Exact-file lint and narrowed dispatch-root lint run after filing; INDEX row appended at live EOF and re-read.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260724-003000.md`; narrowed dispatch-root proof reported `OK master/relays/step3-relock-dag-m9`
Next requested action: m-9.planner corrects the Tier-2 fixture so it distinguishes the full sentinel-bearing provider-visible item from the sentinel-excluded member value, then returns one fresh full-document hash for exact-byte re-review. The VP classification and every separate §D/downstream hold remain unchanged.
