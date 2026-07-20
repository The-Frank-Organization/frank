## DESIGN-REVIEW - s8 config/atom grammar must revise before lock

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r1
PARENT_DISPATCH_ID: s8-m2-grammar-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - activation authorization remains operator-owned and must resolve before the design lock feeds PLAN
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260710-220000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner
SUBJECT: must revise s8-design-m2-grammar - activation direction and axis split are sound, but required/rendered semantics, stale-form precedence, version carrier/fail-open wording, and the operator fork are not lockable yet

DESIGN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - the design conflates seat-rendered, system-observed, and computed required fields.** The fixture and consequence text say the same Block-A fields become “required/rendered” when observe activates (`design:19,46`). Current code makes that impossible as one uniform state:
   - `renderable` suppresses every `system`, `computed`, `system_only`, and `computed_result` row (`render.go:88-98`), so the listed system/computed Block-A fields cannot become seat-rendered merely by changing `PresentLayers`.
   - `Validate` skips requiredness entirely for `system` and `system_only` rows via `ignorePayloadField` (`validate.go:31-42,137-143`), while a `computed`/`computed_result` row can reach requiredness despite being non-renderable. Without an explicitly ordered conductor fill/compute stage, activation can therefore either do nothing for system rows or make a computed row impossible for the seat to satisfy.
   - The live registry demonstrates the split: `ACTIONS_GIT_REF` and `FINAL_GIT_STATUS_SHORT` are seat-owned/renderable; `achieved_evidence`, `evidence_integrity`, `executable_claim_results`, and most Block-A rows are system-owned; `authority_class` is computed.

   **Required revision:** replace the blanket required/rendered claim with an ownership/stage matrix. Pin which rows are seat-rendered and seat-required, which are conductor-produced and validated before commit, and which are computed. Name the pre-validation enrichment boundary that makes system/computed requiredness executable, or explicitly leave those rows outside seat `Render`/`Validate`. The two-state fixture must assert each class at its real consumer grain, not require every Block-A row to appear on a seat form.

2. **BLOCKER - “config-derived PresentLayers” and one-boundary A-1 need a single context plus exact stale-form precedence.** The direction away from `DefaultLayers()` is correct, but three independent evaluation sites currently instantiate the code default: `Render` (`render.go:54`), `Validate` (`validate.go:29`), and grant digest shape (`render.go:254-260`). The design must require one immutable layer context derived from the same pinned config and threaded through all predicate/digest paths; changing only the obvious render/validate calls would leave split semantics.

   The “re-render once, not a bounce-storm” claim also lacks the decisive rule. `Validate` records stale `form_digest` at `validate.go:22-24` and then continues into required-set evaluation (`:27-65`). After activation, a pre-activation submission can therefore receive both `form_digest:re-render` and newly activated required violations based on fields its stale form never showed. **Required revision:** specify whether stale-digest rejection short-circuits required-set validation. Recommended invariant: stale form gets only the typed re-render disposition; after re-render, a fresh digest is evaluated against the activated required set. Add a fixture for old digest -> one re-render response -> fresh form/digest -> activated requirements, with no retry loop.

3. **BLOCKER - the member grammar has no committed version carrier and calls catalog fail-open a safe default.** Section 2 calls the knob and catalog “two new config members,” but section 1 allows the knob to be a field inside existing `EngineConfig`. `EngineConfig` has no version member (`config.go:24-28`), while unknown JSON fields are merely ignored by `json.Unmarshal` (`config.go:52-55`). “Additive-MINOR” is therefore a classification with nowhere to encode/read it if m-7 chooses an engine field. Choose one boundary contract: either a dedicated versioned activation member, or an engine-member version field with exact bump semantics. The m-7 home cannot remain unconstrained while m-2 claims a mechanically detectable MINOR.

   Also split old-reader defaults. Unknown activation state -> observe off is fail-closed for activation. Unknown catalog -> no law enforcement is explicitly fail-open governance, not a safe default (`design:35,40`). Preserve the honest weaker-legacy statement and prohibit old-reader participation in the governed s8 runtime if catalog enforcement is required; do not group both under `ignore-unknown-safe-default`.

4. **HUMAN DECISION - activation authority cannot remain “Still operator-owned” in an otherwise approved lock.** The GRILL_LOCK correctly identifies WHO/WHEN as non-inferable (`design:90-95`), and this decision changes who may activate a new required-set epoch. It must resolve before approval, not merely before a later knob-flip plan. The sibling m-7 dispatch already proposes `operator-authored config_change`; either cite that as an orchestrator-bound input and fold the exact operator gate, or route the open question to the operator and return after the answer. The recommendation is technically coherent: operator-authored, digest-pinned `config_change`, restart-effective; ordinary lanes have no activation surface.

### Accepted Direction

- `layer_present:observe` and CQ-1 remain unchanged; only their runtime input becomes config-derived.
- Knob value changes belong to the config-digest axis, not record `schema_version`.
- Adding catalog laws is additive; weakening/removal requires owner-fidelity governance and a loud compatibility class.
- Forward-only history is sound: store `Read` verifies the committed record (`store.go:161-175`), `Project` reads committed records (`:178-195`), and migration applies record-schema steps (`migrate.go:80-90`); none invokes current fieldspec required-set validation. Keep the proposed post-activation read/project fixture.
- The m-2/m-7 ownership split is correct in shape; reconciliation must bind the exact member home/version carrier, runtime layer context, activation gate, and genesis state before PLAN.

### Re-review Bar

Return one revised design under a new unique pair sub-ID that:

- replaces “all Block A required/rendered” with the executable ownership/stage matrix and fixtures;
- threads one config-derived layer context through every render/validate/digest predicate site;
- defines stale-form short-circuit/precedence and the two-step re-render fixture;
- chooses or constrains a version-bearing knob home and labels catalog old-reader behavior fail-open;
- folds the operator activation-authorization decision into the GRILL_LOCK and design-lock impact.

No c1 atom redesign is requested. No code, PLAN, or IMPL action is authorized by this review.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s8-design-m2-grammar` against `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for both the live directory and this relay

Next requested action: m-2.planner folds the four findings, obtains the operator activation-authority answer or binding citation, and returns a revised DESIGN request for re-review before master reconciliation.
