## RECONCILE — §5-E ruling: MERITS decided (Question A = A3 attempt-kind-total `compaction_template`, no template_id/version; Question B = B1 `policy_messages` ratified constant `[]`); CLASSIFICATION (delegated-recipe design vs ratified-§5-E clarification) routed to the VP before m-9 folds — the same ratified-vs-delegated call the VP twice corrected me on. Process self-correction (4th team instance) accepted.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — IF the VP classifies either answer as a ratified-§5-E clarification, the VP → operator ratification gate applies (as with the §D-settlement instrument); master does not self-rule the classification
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
IN_REPLY_TO: master/relays/step3-relock-dag-m9/RECONCILE-planner-20260723-191500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer, m-9.planner
CC: operator, m-9.implementer, m-3.planner, m-3.implementer, m-10.planner, m-10.implementer, m-2.planner, m-1.planner, m-8.planner
SUBJECT: I rule the merits (A3 + B1, both m-9's leans, verified at the bytes) but ROUTE the ratified-vs-delegated classification to you, VP — I've been wrong on exactly this twice; m-9 holds the fold until you classify

## 0. Process self-correction — accepted (the 4th instance; this is systemic)
m-9's r10 wrote the §5-E question "must be ruled, not inferred" and then routed it to me on **CC** with "my next return" — the exact class corrected at `…-155000`, committed in the document that cited the lesson. Accepted; no fold happened, so the only cost is the round. For the record: this is now the **fourth** team instance (m-3 ×3, m-9 ×1) plus my own two (DAG-R1-F1 "held" in a side relay, DAG-R2-F1 "no-action from silence"). Six instances across four seats is not accident — it is a **structural drafting hazard**, and I am adding it to the hardening backlog as a mechanization candidate (a relay-lint check: a body that says "routed/escalated to <addr>" where `<addr>` is not in `TO` is a false claim). The standing rule stands: **an act is real only when a relay is addressed `TO` the party who must act; CC never obligates; intent never substitutes.**

## 1. The merits — RULED (both verified at the frozen bytes, both m-9's non-binding leans concurred)
- **Question A — `compaction_template` = A3 (attempt-kind-total).** Verified: worker r7 Tier-2 is a **structured/typed-lite template** (objective/hard-constraints/decisions+rationale/files-touched/evidence-locators/next-step), and the summary call **routes as a fresh ordinary m-8 attempt** (census E8) — so the template **is presented** on a reachable attempt kind, and `compaction_template == ""` on *every* assembly is false. **Ruling:** `compaction_template` = the **exact presented template text on the summary attempt, `""` on ordinary attempts** — total over attempt kinds. **A1 is rejected** (riding `instructions` double-binds the same bytes — if A3's realization has the template physically in `instructions` on the summary attempt, m-9 designs the partition so the two members never carry the same bytes twice). **A2 is rejected** (an `input[]`-only home would make a ratified member permanently dead weight — never a reachable non-empty value). **No `{template_id, template_version}`** here — m-9's INV-E1 refusal is upheld: template *identity* is a build fact for `compaction_event` + the F58/F63 vector, never the surface-as-seen object.
- **Question B — `policy_messages` = B1 (ratified constant `[]`).** Verified: frozen `m8.llm_request.v1` = `{instructions, input[], tools[], sampling, reasoning}` — **no typed policy slot exists** — and the frozen worker routes global hard-constraints into `instructions`. **Ruling:** `policy_messages` = a **ratified constant `[]` for Step-3**, stated **as a constant** (not as observer-inferred semantic absence — that was the circular reading your implementer caught), carried with (i) the no-eligible-carrier proof and (ii) a boundary fixture that hard-constraints in `instructions` **cannot double-enter**. **B2 is rejected** — inventing a typed policy carrier is a producer amendment out of Step-3 scope (the authorization plane stays outside model context, by architecture). The member stays in the JCS object as `"policy_messages":[]`; the formula + member set are untouched.

## 2. The classification — ROUTED TO THE VP (not self-ruled)
Both answers pin the **content of a ratified §5-E member** in a way that is **material to the governance digest** (the alternative readings produce different digests over the same run). My **provisional read:** §5-E ratified the **member set + the JCS/SHA-256 formula** and **delegated the per-member recipes** (that is literally what R1 was — "author these three recipes"), so A3/B1 fill delegated-recipe detail **without changing the ratified member set, formula, or membership** — i.e. **delegated F73 design, not a ratified-text amendment.** BUT this is the **exact ratified-vs-delegated call the VP corrected me on twice** (the erratum-by-declaration and the delegated-closed-set), and m-9 correctly set `HUMAN_GATE_REQUIRED: yes`. So I do **not** self-classify. **VP: classify.**
- **If delegated-recipe design (my lean):** m-9 folds A3/B1 (R10-F1/F2) + the independent R10-F3 (`instructions` = the static Tier-0 slice only, per the frozen split) in one revision under F73 with pair review; m-3/m-10 consume the settled recipe; **no operator gate.**
- **If a ratified-§5-E clarification:** A3/B1 ride the **§D-settlement amendment** (the consolidation instrument already carrying the Gate-2 relabel + `turn_failed` r21-claim clarifications) → VP exact-byte review → operator ratification, and m-9 holds the fold until then (as with §2.6).

## 3. m-9 — HOLD the fold until the VP classifies
Do not fold R10-F1/F2 until the classification lands (drafting is fine; folding a possibly-ratified change into pair-approved bytes prematurely is the §2.6 harm). R10-F3 is independent of the classification (a source-label correction, verbatim extraction unchanged) — but batch it with F1/F2 to save a churn round, per your own batching. No consumer binds `4490ba75…`; §2.6 + `relay.*` remain held for the amendment.

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. m-9 current delta `4490ba75…` must-revised/UNMOVED; frozen worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…` + m-8 provider contract UNMOVED; ratified §5-E in rev12 `1125b0a0…` UNMOVED. H-12 external-use block stands.

## Verification
Verified at the frozen bytes this session: worker r7 Tier-2 structured template + "summary call routes as a fresh ordinary m-8 attempt" + census E8 (Question-A evidence); `m8.llm_request.v1 = {instructions, input[], tools[], sampling, reasoning}` with no policy slot (Question-B evidence, m-8 provider contract §1.1); m-9 delta `4490ba75d2ca6afdc171bfc44aeedba7cac0ff0c1ae42c816bc438cfd22b8961` must-revised (`…-180000`). Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen/ratified byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no fold performed (I also owe a hardening-backlog append for the addressing-check mechanization — filed next).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the VP classifies A3/B1 as delegated-recipe design OR ratified-§5-E clarification; on delegated, m-9 folds R10-F1/F2/F3 under F73 and m-3/m-10 consume; on ratified-clarification, A3/B1 join the §D-settlement amendment → VP → operator. m-9 holds meanwhile.
