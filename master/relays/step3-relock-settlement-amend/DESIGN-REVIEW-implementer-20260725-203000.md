## DESIGN-REVIEW — APPROVE r14: both Correction-3 consumption defects close at the live loci; the string-array predicate and `cc="[1]"` negative now match m-2 exactly

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m9-r14
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this approves only the exact m-9 post-ratification owner fold; the operator-ratified packet stays byte-exact and the separate §D join remains gated
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_DOC_SHA256: 514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-3.implementer, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-20260725-193000.md
RELAY_PATH: master/relays/step3-relock-settlement-amend/DESIGN-REVIEW-implementer-20260725-203000.md
SUBJECT: APPROVE exact r14 514f8855 — M9-SETTLE-R13-F1/F2 close; live §7 has one no-longer-held state and the CC decode/fixture contract is byte-faithful to m-2's JSON-string-array rule

DESIGN_REVIEW_VERDICT: approve

m-9.planner — **APPROVE** the complete r14 artifact at exact SHA-256 `514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0`. I re-reviewed the full current artifact, the incoming correction relay, the ratified amendment, the byte-bound m-2 cell, frozen worker r7, and frozen lifecycle r21. Both r13 findings close without widening m-9 ownership or changing any ratified/frozen producer byte.

## Finding closure

### M9-SETTLE-R13-F1 — CLOSED

Live §7 now has one current state:

> `ALL SIX action families settled; the relay.submit cell CLOSED by Correction 3 ... no longer held`

The contradictory live `relay.* HELD` qualifier is gone. A current-section sweep over §§0–11 finds no other live hold for this cell; §7's body agrees that Correction 3 closed it.

**Non-blocking lineage precision:** the incoming relay says the historical hold survives “only in §12.” A literal artifact sweep also finds it in the top Status paragraph explicitly marked `r8 — retained for lineage`. Both occurrences are unmistakably historical and neither competes with the live §7 state, so this does not affect the design verdict. Future status summaries should say “explicit historical lineage” rather than “only §12.”

### M9-SETTLE-R13-F2 — CLOSED

The live derivation now matches m-2 §2.1 exactly:

- non-empty CC that decodes as a **JSON array of strings** selects `cc`, preserving presented order with no dedup/trim/sort;
- non-empty CC that does not so decode selects the presented string in `cc_unparsed`;
- empty/absent CC omits both;
- the two branches remain mutually exclusive and the closed target stays total through required `form_digest`.

The concrete divergence witness is now disposed correctly: schema-valid `cc="[1]"` parses as an array but not an array of strings, so §7 selects `cc_unparsed:"[1]"`, never `cc:[1]`. §10 carries the matching three-leg fixture — string-array → `cc`, non-string-array `"[1]"` → `cc_unparsed`, empty/absent → both omitted — and requires distinct JCS inputs. M-9 derives the value from the presented members and does not re-author m-2's recipe.

## Full-byte preservation review

- **Corrections 1, 2, and 4 remain accepted and unchanged in substance.** Gate 1 owns the disclosure guarantee while Gate 2 remains the frozen validator/drift detector; the run-wide restore and `parked_unknown_capacity_exceeded` no-resume consumer posture remain exact; and the zero-attempt `turn_failed` branch remains an explicit, ratified, non-generalizing supersession rather than an “already descriptive” claim.
- **Correction 3 is now exact end-to-end.** The bound cell hash, required `form_digest`, optional-by-omission coordinates, CC mutual exclusion, parsed-value JCS discipline, target-as-named boundary, and no-authority-from-CC rule all survive.
- **Previously accepted mechanisms remain intact.** The §5-E five-member recipes and sentinel boundary, S-1's complement predicate, carrier/freeze mechanics, m-1 carrier-scoped negatives, and §8 copy-not-compute carriage are unchanged on the reviewed current bytes.
- **The joint gate remains honest.** §9 items 4/5 are still EXACT-FOLDED and JOINT-PENDING, not normative before matching pair approvals and the §D co-sign. During this review m-10 rev16 became separately pair-approved at exact hash `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6` (`step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260725-204500`). R14 was authored before that arrival, so §9's “current producer revision is proposed” status sentence is now superseded by the relay trail; its binding rule remains correct and executable — the §D join must rebase to and name this exact approved successor, never an ancestry/proposed hash.
- Worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, the m-2 cell `5ec7a3d2…`, and the ratified amendment `1fa71cb8…` re-hash byte-exact and unmoved.

## Gate effect

This verdict pair-approves **only** m-9 r14 `514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0` as the completed m-9 propagation-matrix successor. M-9.planner may return that exact hash to Master.

It does **not** co-sign or close the §D two-sided join. Both owner successors are now separately pair-approved — m-9 r14 `514f8855…` by this verdict and m-10 rev16 `3e3c5192…` by `…-204500` — but the join still requires m-9's exact-hash rebase/cross-confirmation and the m-1 redaction/at-rest leg. It grants no integrated re-lock, DESIGN-lock, PLAN, T4/code token, credential/provider action, release binding, live E3, merge, deploy, or external-use authority. H-12 stands.

## Verification

- Incoming relay: `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260725-193000.md` @ SHA-256 `af3df7b6b8bf2eb824235d1143039ef18ca3439b81f801662afe911c871f2b8c`.
- Approved r14: `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md` @ SHA-256 `514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0`.
- Ratified amendment: `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` @ SHA-256 `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b`.
- Bound m-2 cell: `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` @ SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.
- Frozen worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; frozen lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Separately approved m-10 successor: rev16 `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6`; approval relay SHA-256 `35b0b3ca9fc8fa79c20c6c71cad548e99e7097c84f4893d546257414695881ed`.
- Exact-file lint and narrowed dispatch-root lint run after filing; INDEX row appended at live EOF and re-read.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, ratified/frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-settlement-amend/DESIGN-REVIEW-implementer-20260725-203000.md`; narrowed dispatch-root proof reported `OK master/relays/step3-relock-settlement-amend`
Next requested action: m-9.planner returns pair-approved r14 `514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0` to Master. Both owner successors now exist pair-approved; the §D join next binds m-9 r14 `514f8855…` and m-10 rev16 `3e3c5192…` with the m-1 redaction/at-rest leg before any downstream gate moves.
