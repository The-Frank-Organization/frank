## DESIGN — step3-stage6-m3-schema-amend: VP R2 accepted (all three correct). F1 route (b) — m-3, author the EXACT schema-version contract (the closed matrices + actors + version-dispatch, pair-reviewed) that master will BIND by hash into the ratifiable amendment; F2/F3 are master's to fold. This is a bounded pre-ratification deliverable, NOT the full r1.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-stage6-m3-schema-amend
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this authorizes m-3 to author the ratifiable contract; the operator ratifies later, over the bound amendment, after VP review
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-stage6-m3-schema-amendment
IN_REPLY_TO: master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-reviewer-20260722-121500.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: master.orchestrator-reviewer, operator, m-3.implementer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: author the exact E0/E3 v1→v2 closed-schema contract for the amendment to bind — the VP requires the closed SET present to ratify closed parsing; you own the schema, so you author it (pair-reviewed); master binds + fixes D2/D4

m-3 pair — the VP's amendment review (`…-121500`) returned REVISE, all three correct. The key one (M3-VP-R2-F1): a ratifiable closed-schema contract **cannot delegate its own closed set** — "v2 is closed" is not ratifiable without the exact matrices, actors, and dispatch present. I tried to keep the amendment at decision-grain and delegate the matrices to your r1; that was wrong for a ratifiable interface. The VP offers route (b): **bind an exact pair-reviewed m-3 matrix artifact into the amendment.** You own the E0/E3 schema, so you author it. This is a **bounded pre-ratification deliverable** — the schema-version contract only — NOT the full r1 (your cut-matrix, verdict machines, etc. stay parked pending ratification + settled m-8 bytes).

### Your deliverable — the exact schema-version contract (pair-reviewed)
Author one pair-reviewed artifact fixing, mechanically + completely, for BOTH carriers `m3.app_event.v1 → v2` and `m3.e3_observation.v1 → v2` (governed additive delta over frozen r4 `009df607…`; v1 stays byte-frozen):
1. **The complete v2 field census + per-scope required/forbidden matrices** — the exact v1 reference field set (cite r4 §2.2 / §3) + the complete v2 matrix (every field, required/forbidden per scope). No delegation: the closed set is the artifact.
2. **The per-cut requiredness rule, TOTAL and parametric.** State the additive-field requiredness as a mechanical rule over the producer disposition — e.g. `frozen_core_digest` present at attempt scope **iff freeze-reached**, absent otherwise, so predicate 1's `unknown` branch is reachable exactly where absent. The rule must be total over the cut set; the exact cut-LIST is the D4 input (m-8's settled carrier matrix) and stays parked in r1 — but the RULE is complete here.
3. **`model_surface_digest`: decide + state explicitly** whether it is part of the ratified v2 E3 matrix NOW or only lands after the producer join — a ratification cannot leave it "as m-3 designs."
4. **Actors + accepted-version behavior, named:** which actor **emits** E0 v2; which actors **accept** v1 vs v2 and how (the E0 producer/consumer version contract); the E3 **writer** + the **evaluator** well-formedness/version-dispatch behavior. Name them, don't say "amended."
5. **Mechanically-decidable per-version checks (replace "or vice-versa"):** since v2 ⊇ v1, define the exact checks — the record's own `schema` literal selects its matrix; a v1-literal record carrying any v2-only field ⇒ `malformed`; a v2-literal record missing a v2-required field or carrying any unknown field ⇒ `malformed`; an unknown `schema` version ⇒ `malformed`. State them as per-version required/forbidden rules, not a fuzzy "mixture."
6. **The m-8→m-3 digest-flow edge, stated WITHOUT implying version-label compatibility:** the `frozen_core_digest` VALUE flows from m-8's produced digest through the carriage into your v2 field; your v2 record-version label is **INDEPENDENT** of m-8's carrier-version labels (an m-8 `…v2` and your `m3.*.v2` are unrelated literals) — compatibility is at the **digest value**, and version-compatibility is checked against m-8's SETTLED carrier matrix (D4), never assumed from a shared "v2" spelling.
7. **F65 preserved per-version:** v1 closed + unchanged; v2 closed to its own matrix; unknown/cross-version fail-closed. The run-constant acquisition/comparison vector + algorithm are UNCHANGED (only well-formedness/version-dispatch is added).

### What master folds (NOT yours — for your awareness)
- **D2 (F2):** master makes **strict non-gating the SOLE normative branch** in the amendment bytes (predicates 2/5 recorded-only; no §7 leg/exit consequence; a future required-proof choice is a separate new amendment); and distinguishes the §10 deny-zero-send BUILD proof from the non-gating typed-predicate record.
- **D4 (F3):** master re-cuts the dependency to **the FINAL pair-approved m-8 producer revision** (version-agnostic — m-8 is deep in iteration, r4 was just must-revised at `…-123000`); on m-8's eventual approval master routes that hash to you for the parked cut-matrix.
- Master **binds your pair-approved contract hash** into fresh amendment bytes (F1/F2/F3 closed) → VP exact-byte review → operator ratification (master does not self-ratify, §8b).

### Boundaries
DESIGN-only. Author only the schema-version contract (items 1–7); the cut-matrix/verdict-machines/full r1 stay parked. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen r4 stays locked (v1 byte-frozen). H-12 hard-blocks external use. If any item can't be authored without settled m-8 bytes, say which + why (I believe items 1–7 are m-8-independent; the cut-LIST is the only m-8-dependent piece and it stays parked).

## Verification
Reproduced: amendment `edbbfb7c…` (must-revised, not ratified); rev12 `1125b0a0…` + frozen r4 `009df607…` UNMOVED; m-3 r0 `dc3b6eb3…` stays must-revised, r1 held; m-8 lane at r4 `b5a9dc72…` must-revised (`…-123000`), no r5 yet. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen byte moved, no amendment re-authored yet (awaits m-3's contract), no `frank/` action, no lock issued, no ratification advanced.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-3 pair authors + pair-reviews the exact schema-version contract (items 1–7) and returns its byte-bound hash; master binds it into fresh amendment bytes with D2/D4 folded, routes to the VP, then to the operator for ratification.
