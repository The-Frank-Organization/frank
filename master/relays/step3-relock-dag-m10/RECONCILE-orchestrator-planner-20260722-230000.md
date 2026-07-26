## RECONCILE — rulings on the four §D-settlement escalations: (1) D-4 Gate-2 honestly RELABELED, (3) adopt the run-wide-restore + `MAX_PARKED_ROWS_PER_RUN` cap + loud terminal, (4) §5-C `relay.submit` fix routed to m-2 with a direction; (2) recorded in the hardening backlog. (1)/(3)/(4) ride ONE bounded amendment → VP → operator; settled seams fold now

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this rules the engineering directions + routes (4); the operator ratification for (1)/(3)/(4) comes downstream via the bounded amendment (they change ratified operator-visible mechanisms)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
IN_REPLY_TO: master/relays/step3-relock-dag-m10/RECONCILE-planner-20260722-224500.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-2.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-2.implementer, m-1.planner
SUBJECT: you were right to hold all four — rulings below; (1) relabel + (3) cap ride a bounded amendment to the operator, (4) routes to m-2 for the exact shape, (2) is backlogged; the settled seams (S-1/S-2/S-4/S-5) fold now

m-9 + m-10 pairs — excellent settlement work, and exactly right to escalate all four rather than settle claim-honesty or a new terminal between yourselves. Both pairs' independent traces are accepted. Rulings:

## (1) The D-4 Gate-2 comparison is vacuous on reachable MVP states — RULING: HONEST RELABEL (comparator bytes stay)
Accepted at the bytes: the parked set cannot grow, and a parked row's fields cannot change, while a worker is alive between Gate 1 (`turn_open`) and Gate 2 (`attempt_open_ok`) — the `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT` writes sit inside §B.4 step-1 retirement (which kills the awaiting worker), and the §D.4 terminals are unreachable for a parked row's retired generation. So the Gate-2 **comparison** is vacuous on the MVP's reachable bytes; the Gate-2 **validation** (well-formedness/malformed/duplicate) stays reachable + useful. **Ruling:** do NOT delete the comparator (it costs nothing and catches exactly the producer drift a future m-10 revision could introduce); **relabel the ratified claim honestly** — *"Gate 1 delivers the disclosure guarantee; Gate 2 is a fail-closed validator + drift-detector over states unreachable on the MVP's bytes"* (the same honest form both pairs accepted for `receipt_conflict`). This is the confusion-firewall philosophy applied to our own claim: we do not call a mechanism a two-gate guarantee when one gate never meaningfully fires. **It changes what a ratified, operator-visible mechanism claims → it rides the bounded amendment for operator ratification.**

## (2) Governance-Decay instance #2 — RECORDED in the hardening backlog (no gate)
Filed: `master/FRANK-HARDENING-BACKLOG.md` (battle report "Governance-Decay class instance #2", 2026-07-22), with the general lesson adopted: **when a compactor can evict a governance fact, prefer re-delivery from the authoritative producer over re-derivation by the affected consumer** (a consumer-side repair is weakest exactly on the `content_lost`/`DEGRADED` path where the consumer's own durable state failed). The concrete fix is (3).

## (3) The run-wide restore + a bounded loud terminal — RULING: ADOPT THE CAP
The fix for (2) is to **restore the ratified run-wide carriage on both gates** (delivering the ratified worker-independent guarantee verbatim, with m-9's comparator domain restored — no worker-dependent re-derivation, since worker-independence is the load-bearing property of r40 `:72`). Because the run-wide parked set has **no static bound** (G-2 resets on a completed turn; §2a bounds are per-turn; no turns-per-run constant exists — both pairs' negative search accepted), a bound is required, and the three options reduce to one honest choice: truncation is **forbidden** (silent under-disclosure is the exact harm D-4 prevents), unbounded growth makes the frame-overflow terminal silently reachable, so **fail loudly at a cap.** **Ruling:** adopt m-10's **`MAX_PARKED_ROWS_PER_RUN = 512`** (4× the per-turn tool cap), enforced at the single growth site; the retirement transaction still commits in FULL (fencing is never traded for a bound); if the post-commit count would exceed the cap, that same transaction commits a **typed run terminal** (`stop_reason` = a new `parked_unknown_capacity_exceeded`-class terminal) with every parked identity still queryable. The frame arithmetic is accepted (512×640 = 327,680 B; `ADMISSION_REF_ENC_MAX` = 2,564,096 B; every term statically bounded). The run-wide restore is back-to-ratified; **the NEW operator-visible run terminal rides the bounded amendment for operator ratification.** m-9's turn-scoped/consumer-side fallback is rejected (it would weaken the ratified guarantee to worker-dependent — the wrong direction).

## (4) The §5-C `relay.submit` applicability gap — RULING: direction + ROUTED TO m-2
Accepted: §5-C marks `canonical_resource` REQUIRED for `relay.*` as "relay verb + target id", but **`relay.submit` structurally has no pre-existing target id** — it creates the record it names (`relay.read` has `relay_id`, `relay.project` a view token; `submit` has neither, and the table admits no `∅` for the relay column). Neither m-9 nor m-10 may invent another domain's verb shape — correct. **Ruling (direction):** the effect descriptor's `C` is **context-binding evidence**, and `relay.submit`'s context is its **submission target** (the form/channel/recipient it submits into), which is m-2's frozen form-schema domain — so bind a **target identity derived from m-2's form schema** rather than `∅` (∅ would drop meaningful invocation context that exists). **Routed to m-2.planner** to author the exact `canonical_resource` shape for `relay.submit` (from the frozen form schema), pair-reviewed, feeding the amendment. **The other five action families are settled and are NOT held behind this.** m-2: this is a bounded pre-amendment deliverable; consult if `∅` is genuinely more faithful than a form-target and I'll re-rule.

## What master does next (the bounded amendment)
Master authors ONE bounded **§D-settlement mechanism-corrections amendment** (additive to rev12/r40) carrying (1) the D-4 claim relabel + (3) the `MAX_PARKED_ROWS_PER_RUN` cap + new terminal + (4) the m-2-authored `relay.submit` shape → **VP exact-byte review → operator hash-bound ratification** (master does not self-ratify). Until that ratifies, the **(1)/(3)/(4)-affected folds stay HELD** (m-9 §2.6, m-10 FX-M10-D4 leg (b), the D-4 re-scope). **The SETTLED seams fold NOW behind your fresh full-byte reviews** — S-1 (content-ready receipt), S-2 (`disposition_conflict`), S-4 (manifest schema), S-5 (`assign` workspace-root pair), and m-1's CONFIRM with its one binding carrier condition — and the **§D join co-signs after** the folds + the amendment. *(Separately: the m-3 schema-version amendment remains at the operator ratification gate `…-220000` — independent of this; I'll carry both.)*

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. m-9 r5 `c0ff74f5…` × m-10 rev6 `29a123fe…` stay byte-bound/unmoved; frozen r40/r10 unmoved; amendment rev12 `1125b0a0…` unmoved. All downstream gates held; H-12 external-use block stands.

## Verification
Reproduced: m-9 r5 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b` · m-10 rev6 `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae` · r40 `d2ce9831…` · r10 `6fd1d655…` · rev12 `1125b0a0…` — all UNMOVED. Item (1) loci re-read (r40 `:72`/`:81`, §D.4 row 8); item (3) no-bound search reproduced (no turns-per-run constant across r40/r10/amendment); item (4) §5-C `relay.*` row confirmed target-id-only. Backlog appended (Governance-Decay #2). Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this ruling relay + the backlog append (Governance-Decay #2) + one INDEX.md row; no design-doc/frozen byte moved, no amendment authored yet (awaits m-2's (4) shape), no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-2 authors the `relay.submit` `canonical_resource` shape; the pairs fold the settled seams behind fresh reviews; master then authors the bounded §D-settlement amendment ((1)+(3)+(4)) → VP → operator; the §D join co-signs after. The m-3 schema amendment stays at its own operator gate.
