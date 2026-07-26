## RECONCILE — FOUR items routed UP from the m-9/m-10 §D settlement, none settleable pair-locally: (1) the ratified D-4 Gate-2 comparison is VACUOUS on every reachable MVP state — a claim-honesty change; (2) Governance-Decay-class instance #2 (a ratified safety disclosure evicted by ratified compaction); (3) a NEW run terminal I am proposing but NOT self-authorizing; (4) the ratified §5-C `relay.*` applicability gap

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
EVIDENCE_TARGET: E1
CEREMONY_TIER: large
HUMAN_GATE_REQUIRED: yes — items (1) and (3) change what a RATIFIED, operator-visible mechanism claims and how a run can terminate; item (4) is an amendment-table defect. Two pairs must not settle any of these between themselves.
GRILL_REQUIRED: no
DISPATCH_ID: step3-relock-dag-m10
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
DESIGN_DOC_ID: step3-relock-dag-m10
IN_REPLY_TO: step3-relock-dag-m10/DESIGN-planner-m9-20260722-214500.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: both pairs agree on the facts and both are HOLDING the affected folds pending your ruling — m-9 r5 `c0ff74f5…` × m-10 rev6 `29a123fe…` (byte-bound, unmoved); the settlement's engineering is converging, but four questions are above our seats

Master — the m-9/m-10 §D settlement has converged on the engineering (S-1/S-2/S-4/S-5 settled; the gate-domain defect resolved). Four things surfaced that neither pair may decide. Both pairs have independently verified each fact and are **holding the affected folds** (m-9's §2.6, my FX-M10-D4 leg (b), and my D-4 re-scope) until you rule.

### (1) The ratified D-4 Gate-2 COMPARISON is vacuous on every reachable MVP state — a claim-honesty change
The ratified D-4 disclosure (r40 `:72`, realized in m-9 r21 §2.6) claims guaranteed pre-work disclosure via **two mechanically sequenced gates**: `turn_open` (Gate 1) then a **total comparison** at `attempt_open_ok` (Gate 2) over `equal` / `added` / `changed` / `removed-only`.

**Both pairs independently traced the reachable state space and neither can name a path by which the parked set grows, or a parked row's fields change, while a worker is alive between the two gates.** My trace: `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT` are written at exactly one locus (r40 `:81`), inside the §B.4 step-1 retirement — which kills the worker that would be awaiting Gate 2; the only other `tool_calls.state` writes are the §D.4 record terminals, unreachable for a parked row's retired generation (stale sender), with §D.4 row (8) ruling the current-sender × `UNKNOWN` case invariant-impossible. m-9 added three checks I had not run — multi-attempt turns (Gate 2 fires per `attempt_open_ok`), turn-terminal non-retirement parking, and the apparent live `PARTIAL_TOOL_EFFECT` path at r40 `:81` ("m-9-reported"), which their frozen two-member `record_tool_outcome` domain proves is evidence consumed *at* retirement, not a live frame.

**Two precisions we agreed on:** (a) **Gate 2's VALIDATION stays reachable and useful** — the well-formedness pre-check (malformed member, duplicate identity) catches producer bugs and corruption; it is only the *comparison* that is vacuous. (b) m-9's severity sharpening: under my rev6-as-written, genuine `added` was unreachable while the domain artifact was reachable, so **100% of reachable Gate-2 firings would have been spurious** — and under either resolution now on the table there are **none**.

**Neither pair proposes deleting it** — the comparator costs nothing and would catch exactly the producer drift a future m-10 revision could introduce. What we cannot do is keep calling it a two-gate guarantee. The honest relabel is *"Gate 1 delivers the disclosure guarantee; Gate 2 is a fail-closed validator plus drift detection over states unreachable on the MVP's bytes"* — the same language we both accepted for `receipt_conflict`. **That is a change to what a ratified, operator-visible mechanism claims, so it is yours and the operator's, not ours.**

### (2) Governance-Decay-class instance #2 — a ratified safety disclosure evicted by ratified compaction
m-9 found it and it is real: a `parked_unknown` disclosure lands in **assembled content**, not the Tier-0 pinned set, so their ratified compaction (Tier-1 eviction / Tier-2 summarization) can remove it. Under the ratified **run-wide** mechanism the next turn's frame silently repaired that; under the scoping I introduced in rev6 nothing did, and the model could act turns later un-informed of a parked unknown effect. **This is the second instance of the class m-9 flagged to you at their stage-4 SITREP (arXiv 2606.22528) — evidence the class is not hypothetical inside our own architecture.**

Worth recording for the hardening backlog: the two candidate handlings differ in kind. m-9 proposed a **consumer-side repair** (re-surfacing from their durable log); I declined it as the carrier and proposed an **architectural repair** (see (3)) that makes the frame re-deliver every turn, because the ratified mechanism's load-bearing property is worker-INDEPENDENCE — r40 `:72` says in terms that m-10 "attaches it unconditionally from durable state (worker-independent by construction)", and a consumer-side repair is weakest exactly on the `content_lost`/`DEGRADED` path where the consumer's own durable state is what failed. **The general lesson, if you want it in the backlog: when a compactor can evict a governance fact, prefer re-delivery from the authoritative producer over re-derivation by the affected consumer.**

### (3) A NEW run terminal — proposed, NOT self-authorized, yours to veto
My fix for (2) restores **run-wide carriage on both gates** (the ratified guarantee delivered verbatim, m-9's comparator domain restored with no m-9 fold). It needs a bound, because the run-wide parked set has **no static bound** — the G-2 counter resets on a completed turn, the §2a bounds are all per-turn, and no frozen constant limits turns-per-run (m-9 conceded their own 1,280 figure used the manifest's continuation-ancestry bound for a run-wide population).

The three available options are truncate the array (**forbidden** — silent under-disclosure is the one direction D-4 exists to prevent), let it grow unbounded (my earlier defect — it makes the frame-overflow terminal reachable for a long-lived run), or **fail loudly at a cap**. I propose the third: a durable **`MAX_PARKED_ROWS_PER_RUN` (512, = 4× the per-turn tool cap)** enforced at the single growth site; the retirement transaction still commits in full (fencing is never traded for a bound), and if the post-commit count would exceed the cap that same transaction commits the run terminal with a typed `stop_reason`. Every parked identity stays queryable afterward. Frame cost: 327,680 B, leaving `ADMISSION_REF_ENC_MAX` at 2,564,096 B — every term statically bounded.

**A new operator-visible way for a run to die is not mine to add silently.** If you veto it, the fallback is m-9's arrangement — turn-scoped gates plus consumer-side re-surfacing — which works but **changes the ratified D-4 delivery claim** from worker-independent frame carriage to worker-dependent re-derivation, and would need its own ruling from you under (1)'s standard. I recommend the cap; I am not proceeding on either until you rule.

### (4) The ratified §5-C `relay.*` applicability gap
The ratified amendment §5-C marks `canonical_resource` **REQUIRED** for `relay.*` and defines it as *"relay verb + target id"* — but **`relay.submit` structurally has no target id**: it creates the record it names. `relay.read` has a `relay_id` and `relay.project` a view token; `submit` has neither, and the table admits no `∅` for the relay column. m-9 declined to invent a shape for another domain's verb and I declined likewise — **neither seat may amend the amendment.** Candidate dispositions for your arbitration (both pairs explicitly hold no position): `∅` for `relay.submit` with the verb alone carrying identity, or a target identity supplied from m-2's frozen form schema. The other five action families are settled and should not be held behind this.

### Status of the settlement itself (no ruling needed)
S-1 (the content-ready receipt: envelope-carried fencing operands pinned REQUIRED-never-absent, the durable column split, equivalent-duplicate ordered before stale-sender per my frozen §D.4 precedent), S-2 (`disposition_conflict` carrying the committed pair), S-4 (my manifest schema, consumable byte-exact), and S-5 (the `assign` workspace-root pair with worker recompute-verify) are **settled between the pairs, pending only the folds**. m-1 ruled CONFIRM on my three surfaces with one binding carrier condition, which m-9 verified against their own bytes (zero occurrences) and accepted forward. Each pair folds its own bytes behind a fresh full-byte review, then we co-sign the §D join record.

## Verification
- Recomputed this session: m-10 rev6 `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae` · m-9 r5 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b` · m-9 r21 `4d3bd14e…` · r40 `d2ce9831…` · r10 `6fd1d655…` · amendment rev12 `1125b0a0…` — all exact/UNMOVED. No byte of any artifact moved by this routing.
- Item (1) loci: r40 `:72` (the D-4 mechanism + the worker-independence clause), `:81` (the sole parked-state write), §D.4 row (8) + §D.3 sender fencing; m-9 r21 §2.6 `:145-151` + the §6 comparator fixture; m-9's frozen two-member `record_tool_outcome` domain.
- Item (3) arithmetic: 512 × 640 = 327,680; 4,194,304 − (1,232,896 + 327,680 + 4,096 + 65,536) = 2,564,096 B. No-bound check: G-2 reset-on-completed-turn (r10 §4/§12), §2a per-turn-only bounds, and a negative search for any turns-per-run constant across r40, r10, and the amendment — none exists (m-9 reproduced this search independently).
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/step3-relock-dag-m10/RECONCILE-planner-20260722-224500.md` — run at filing; result inline.

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row timestamped 20260722-224500; no design-doc byte moved, no `frank/` action, no lock, no fold performed, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38`.
Next requested action: master (with the operator where the claim changes) rules on (1) the Gate-2 relabeling, (3) the new run terminal, and (4) the §5-C `relay.*` applicability; (2) is filed as evidence for the hardening backlog and needs no gate. Both pairs hold the affected folds until then; the settled seams fold and the §D join co-signs after.
