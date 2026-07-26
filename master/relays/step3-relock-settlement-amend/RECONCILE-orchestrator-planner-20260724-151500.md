## RECONCILE — §D-settlement amendment AUTHORED and routed for VP exact-byte review: four bounded mechanism-corrections ((1) D-4 Gate-2 relabel · (2)/(3) run-wide restore + MAX_PARKED_ROWS cap + parked_unknown_capacity_exceeded terminal · (3)/(4) relay.submit cell bound by hash · (4) turn_failed zero-attempt clarification), additive to rev12/r40, r21 byte-frozen; two frozen-terminal-claim adjudications flagged for the VP; master does NOT self-ratify

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — a ratified-mechanism amendment (a relabeled operator-visible claim + a new operator-visible terminal + a §5-C cell + an r21 scope clarification); on VP approve it goes to operator hash-bound ratification (§8b); master does not self-ratify
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-dag-m10/SITREP-planner-20260723-150000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: VP exact-byte review requested on `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` SHA-256 `1f822e4711d5772ffafc68c4183ddb0faa33250e5b9d9372ead0e4128c34dbe7` + the bound m-2 cell `5ec7a3d2…` as one ratification packet; two corrections are frozen-terminal-claim adjudications I authored as additive clarifications — please rule the adjudication class

## The instrument
**`master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` @ SHA-256 `1f822e4711d5772ffafc68c4183ddb0faa33250e5b9d9372ead0e4128c34dbe7`.** It is **additive** — it moves no byte of rev12 `1125b0a0…`, r40 `d2ce9831…`, or r21 `4d3bd14e…`. It carries the four §D-settlement escalations I ruled at `…-230000` `922b796d…`, with correction 3's exact cell owner-authored + pair-approved by m-2.

**Ratification packet = the amendment file `1f822e47…` + the bound m-2 cell `5ec7a3d2…`** (one unit, exactly as the m-3-schema packet bound its contract).

## The four corrections
1. **D-4 Gate-2 honest relabel** — the ratified two-gate disclosure claim is corrected to "Gate 1 delivers the guarantee; Gate 2 = fail-closed validator + drift-detector over MVP-unreachable states." **Comparator bytes in r40 §D.4 + m-9 §2.6 are NOT deleted** — only the operator-visible *claim* changes. Both pairs' vacuity traces accepted.
2. **Run-wide parked-set restore + `MAX_PARKED_ROWS_PER_RUN = 512` + the new terminal `parked_unknown_capacity_exceeded`** — restores the ratified worker-independent guarantee (rejecting m-9's turn-scoped fallback), bounds the otherwise-unbounded run-wide set, commits the retirement transaction in FULL on overflow + a typed loud terminal, truncation forbidden, every parked identity queryable. Frame arithmetic statically bounded (512×640 < `ADMISSION_REF_ENC_MAX`).
3. **§5-C `relay.submit` `canonical_resource` cell** — bound by hash to m-2's pair-approved `5ec7a3d2…`: `"relay.submit:" || SHA-256(JCS{form_digest, dispatch_id?, to?, cc?|cc_unparsed?})`, `form_digest` REQUIRED (total). Realizes my `230000` (4) form-derived direction; **I decline m-2's offered `∅` re-rule** (∅ discards real invocation context) — but flag that m-2 will take a re-rule to `∅` if you prefer after reading its §1.
4. **`turn_failed` zero-attempt scope clarification** — r21 `:115`'s "after the bounded attempt(s)" is descriptive, not a precondition; `turn_failed` also fires pre-attempt (ADMITTED→ASSEMBLING), the `turn_denied` precedent already pre-attempt. **r21 left byte-exact.**

## Two adjudications I need you to rule (the frozen-terminal-claim class)
Corrections **1 and 4** both change what a *ratified/frozen* terminal-claim means. I authored **both as additive named clarifications** with the underlying bytes (r40 comparator, r21 `:115`) **left byte-exact** — my position being that the mechanism bytes are unchanged and only the disclosure *claim* / *scope reading* is corrected. **You may instead rule that either requires an explicit edit to the ratified/frozen text** (a heavier change). This is precisely the ratified-vs-realization line you have corrected me on before, so I route the class to you rather than presuming the light form. Corrections 2 and 3 are unambiguously new-mechanism / gap-fill and carry their own operator-ratification need.

## Scope discipline (what the amendment does NOT do)
It authorizes no m-10 schema change beyond correction 2's named terminal + restored carriage (owner-folded under F73 post-ratification); it does not touch the ratified §5-E five-member set / formula (A3/B1 settled separately at m-9 r12 `04422965…` under your F73 classification `4c254307…`); it licenses no §D join, re-lock, DESIGN-lock, PLAN, T4/code, credential, provider call, release binding, live E3, merge, or deploy. H-12 external-use block stands.

## On approve
Operator hash-bound ratification (§8b agent-authored + operator-cited; master does not self-ratify). Then: m-9 folds its final batched revision (§2.6 relabel + §7-`relay.*` on the bound cell) in ONE F73 revision; m-10 folds correction 2 + its FX-M10-D4 leg; the §D two-sided join co-signs after; consumers rebase once. The m-3 schema amendment (already operator-ratified rev3 `9e874df8…`) is the separate, independent instrument.

## Verification
Hashed on disk this session: the amendment `1f822e4711d5772ffafc68c4183ddb0faa33250e5b9d9372ead0e4128c34dbe7`; bound m-2 cell `5ec7a3d254da…`; my §D ruling `230000` `922b796d…`; rev12 `1125b0a0…`, r40 `d2ce9831…`, r21 `4d3bd14e…` UNMOVED. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — the amendment doc `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` + this relay + one INDEX.md row; no frozen/ratified byte moved, no `frank/` action, no lock issued, no self-ratification, no fold performed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP exact-byte review of the amendment `1f822e47…` + bound cell `5ec7a3d2…` — rule the two frozen-terminal-claim adjudications (additive-clarification vs explicit-edit) and the `∅`-vs-form-target question; on approve → operator ratification; on revise → master folds and re-routes.
