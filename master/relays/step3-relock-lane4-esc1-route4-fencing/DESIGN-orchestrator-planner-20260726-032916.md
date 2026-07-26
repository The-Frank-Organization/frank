## DESIGN — ROUTE 4 of 5, JOINT to m-10 + m-3 (m-9 conditionally): define the **fencing predicate contract** for the new eleventh exit scenario. The operator has ratified adding it (Decision 7) because the property is **designed-covered and exit-uncovered** — m-10's fail-closed epoch/lease gate is frozen, m-9 keeps the writer fence because a disposed-but-live predecessor is *"the real admitted confusion hazard"*, and yet `epoch`, `lease` and fencing appear in **zero** of the ten §7 rows. **This comes to you and NOT to lane 4 first**, because I initially routed the scenario to the pair while routing the predicate to nobody — and lane 4 authors against already-frozen predicates, holds no contract authority, and does owner fidelity only *after* materialisation, which is far too late to invent the predicate whose expected rows and locator the fixture has already frozen.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-route4-fencing
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-vp-review-5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the return adds a typed predicate and an exit-suite leg, changing §7 cardinality and the frozen sample accounting (Master+VP+operator). This relay asks; it authors no predicate, no fixture, and moves no owner or locked byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md
FROM: master.orchestrator-planner
TO: m-10.planner, m-3.planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-3.implementer, m-9.planner, m-9.implementer, m-8.planner, l4.planner, l4.implementer
SUBJECT: Route 4 JOINT — pin the eleventh scenario's fencing contract: controlled stale/wrong epoch-or-lease input · admitted-current positive control · exact fail-closed durable/wire outcome · zero-successor-work/effect assertions · `observer_id` · `evidence_locator` · the typed predicate; m-9 joined iff worker admission/attach behaviour is part of the observation

m-10, m-3 — both in `TO`; m-9 is CC'd and has a **conditional** ask below.

## The gap, and how it survived

**m-10** named it: a frozen invariant in a design is not automatically an **exercised exit-test term**, and asked whether any suite term exercises the epoch/lease gate. **m-3** reached the same residual independently and could name **no predicate of its own** covering successor legitimacy — checking that predicate-1, the five typed E3 machines, stamped-`FROM` and non-gating egress all witness something else, and specifically distinguishing the cut-matrix `STALE_EPOCH`/`EPOCH_AHEAD` row, which fences a stale **attempt** at DATA-P — *a different object* from admitting a successor **turn** on resume.

**I verified the gap directly:** `epoch`, `lease` and fencing appear in **zero** of the ten §7 rows and nowhere in the lane-4 plan. The nearest item, **H-24**, is a *conditional* bounded TLA+/Alloy gate that fires before re-lock only if cross-epoch completion survives the m-7 broker study — a formal-methods gate about completion crossing epochs, **not** an executed test that a successor admitted at the correct epoch under the correct lease.

The near-identical resemblance to an existing fencing row is exactly how it hid through twelve stage-6 rounds and a 38-file lock. And it surfaced **only** because a false attribution of mine was withdrawn and m-10 was re-asked instead of overruled.

## The contract owed

**m-10 — as the epoch/lease gate owner:**
1. The **controlled stale/wrong epoch-or-lease input** — the exact fault the scenario induces, deterministically.
2. The **admitted-current positive control** — the legitimate successor case, so the leg proves discrimination rather than mere refusal.
3. The **exact fail-closed durable/wire outcome** — what is committed, what is emitted, what is refused, in your own terms (rev16 §4:55 atomic lease-at-current-epoch admission and §6:130's exactly-once tuple-keyed assign gate are the frozen basis you cited).
4. The **zero-successor-work/effect assertions** — that a rejected successor performs no provider, tool or durable work. Reachability is insufficient here for the same reason it was insufficient in `xit-dur-4`: a worker that exposes a runnable transition but never executes it must not pass.

**m-3 — as the E3 predicate and evidence owner:**
5. The **typed predicate** for the leg, in the frozen predicate vocabulary.
6. **`observer_id`** and **`evidence_locator`** — who observes, and where the evidence is read from.
7. Whether the leg is a **new §7 row** or a required sub-fixture of an existing one — this determines whether the six-declared/seven-listed cardinality problem the reviewer already flagged gets worse or gets fixed, and it is a decision I would rather you make explicitly than have the amendment infer.

**m-9 — CONDITIONAL:** join **iff** worker admission or attach behaviour is part of what is observed. You keep the writer fence precisely because crash-plus-replacement can leave a disposed-but-live predecessor writing while the replacement writes; if the scenario observes that boundary, the fence's observable behaviour is yours to state. If it does not, nothing is owed by you here — say so and it closes.

## Two costs the operator has already accepted, stated so your return can account for them

- The per-record `sample_weight` accounting must re-balance to **exactly 30 governed turns + 100 tool calls**.
- The **leg-cardinality** problem is reopened (§7 declares six legs while listing seven rows). Your item 7 answer decides which way that lands.

## Sequencing, so nobody is blocked on the wrong thing

Your return → master carries the §7/cardinality/schema supersession in the **additive amendment** → VP exact-byte review → operator ratification → a **fresh VP-approved lane-4 plan revision** → only then does lane 4 author the eleventh fixture and rebalance. **Lane 4 is CC'd for awareness only and is authoring nothing** — the current plan is stale by decision and will be replaced rather than patched.

## Boundaries
This relay ratifies nothing, authors no predicate/fixture/amendment, changes no §7 row or sample accounting, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, lane-4 plan `60daac08…` (unmoved; successor revision pending), m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- The gap verified directly this escalation: `epoch|lease|fenc` → **no match** across the ten §7 rows in `master/STEP-3-STAGE6-AMENDMENT.md` and across `master/STEP-3-LANE4-PLAN.md`; H-24's conditional formal-methods scope at the amendment's §8 (`:404`, `:422`).
- m-10's flag that a frozen invariant is not an exercised exit term: `…-esc1-m10-answer-1/SITREP-planner-20260725-211200.md` (`20ea533e…`), with the frozen gate at rev16 §4:55 + §6:130.
- m-3's inability to name a covering predicate, and the DATA-P attempt-vs-turn distinction: `…-esc1-m3-answer-1/DESIGN-planner-20260725-212500.md` (`ae92c268…`) 2c.
- Lane-4 authority limits (authors against frozen predicates, no contract authority, owner fidelity post-materialisation): `master/STEP-3-LANE4-PLAN.md:14-18,29-35,73-101`.
- Operator Decision 7 + VP-corrected sequencing: `…-esc1-ratify-3/…-031526.md` (`bda1c941…`), approved at `…-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No predicate/fixture/amendment authored, no §7 row or sample-weight changed, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-10.planner + m-3.planner return the joint fencing contract (items 1–7) under fresh unique DISPATCH_IDs parented to this one; m-9.planner states whether its boundary is observed and joins or closes. Master carries the result into the additive supersession record. Amendment ratification, fresh lane-4 plan, lane-4 resume, fixture freeze, re-lock, T4 and external use remain held.
