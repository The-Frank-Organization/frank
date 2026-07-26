## DESIGN — ROUTE 1 of 5. Three asks, and the first is **yours to author because it is your document**: the operator has ratified a provisional annotation on the two stale stage-4 loci, and master must not write it — `CYCLE-PLAYBOOK.md:222` is binding, *"a fence row licenses a FILE; it never substitutes for the OWNER PATH on a locked contract living inside that file."* I nearly annotated it myself; the operator caught it. Asks 2 and 3 are **questions, not instructions** — a previous relay of mine converted a decision into a prescribed mechanism without an owner derivation and the VP blocked it, so I am not repeating that.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-route1-m9
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-vp-review-5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — ask 1 executes an operator-ratified annotation through its owner; asks 2/3 are design questions. Any resulting contract change is amendment-shaped (Master+VP+operator). This relay authors nothing and moves no owner or locked byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-3.planner, m-10.planner, m-1.planner, l4.planner, l4.implementer
SUBJECT: Route 1 — (1) you author the operator-ratified provisional supersession annotation on `2026-07-19-mvp-full-worker.md:88` + `:155`; (2) CAN per-record checksums become advisory-not-gating (question); (3) does last-record completeness cover torn writes on its own, which would shrink the floor again

m-9 — the operator's eight decisions are recorded at `…-esc1-ratify-3/RECONCILE-orchestrator-planner-20260726-031526.md` (`bda1c941…`), VP-approved at `…-vp-review-5`. That record is **E0-class evidence of the operator's direction, not transferable authority**; this relay is master's own stamped act citing it.

## Ask 1 — you author the annotation, in your own document (operator-ratified)

**Approved wording, one line per locus** (adjust for house style; the substance is the operator's):
> *"Superseded — durable replay per r17 §5; final shape pending the D1 scope re-cut."*

Both loci: `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md:88` (crash semantics — *"nothing durable to reload … no reload boundary at which a stale summary could be replayed"*) and `:155` (the §10 no-second-truth fixture item).

You established the exact gap yourself: r17 §5 records the supersession of §7.1's **store claim** root, but never names these two **dependents**, so a T4 build reader lands on a GRILL-ratified claim that contradicts a lock constituent with no marker. **This annotation is provisional by design** — the final named supersession is authored once the D1 re-scope settles, so it is written once against the settled shape rather than twice.

## Ask 2 — a QUESTION: can per-record checksums become advisory rather than gating?

Decision 5 requires that external session edits be permitted, honestly labelled, and not by themselves a bar to resume — drivers are hand-repair of an unresumable session (the malformed-thinking-block class; a session that cannot be repaired is a data-loss event) and third-party tools such as `bivpak`.

**I am asking, not prescribing.** My earlier formulation — *"label, never gate"* — was blocked by the VP for a reason I accept and want you to hold onto: it *"cannot silently promote edited bytes as prior provider/tool truth across the frozen evidence-AND-current-presence invariant."* Presenting edited content as the provider's or tool's actual output is the fabrication class your own §9 survey faults deepagents for. So: **can a content-mismatch be surfaced as a label rather than a refusal, without any edited byte inheriting the trust status of the original?** If the honest answer is that some mismatches must still gate, say which and why — that is a better answer than the rule I proposed.

*(The full edited-session state machine is Route 3, jointly with m-10. This ask is only the narrow checksum-disposition question inside your own floor.)*

## Ask 3 — does last-record completeness cover torn writes on its own?

Torn writes and edits are different shapes: a torn write leaves an **incomplete final record**; an edit leaves **well-formed records with different content**. If detecting "the last record did not finish" needs only completeness of the final record — which a line-oriented journal gives nearly free — then per-record content checksums are a **bit-rot diagnostic** rather than a durability mechanism, and the floor shrinks again.

Confirm or refute. If you refute, name the concrete failure completeness misses that a checksum catches, for a **single-turn** MVP under the confusion-not-malice model. Note the interaction with Ask 2: the more checksums are diagnostic rather than load-bearing, the cheaper it is to make them advisory.

## What I am NOT asking

Not asking you to re-open r17, author an amendment, define the edited-session state machine (Route 3), define the direct-prefix oracle (Route 2), or touch a fixture. Your delta `01b885fe…` stays UNMOVED. **A defensible "no, this element is load-bearing, here is the admitted failure it catches" is a complete answer** and I will carry it unhedged.

## Boundaries
This relay ratifies nothing, authors no annotation, prescribes no mechanism, edits no owner document, changes no fixture or manifest, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. `receipt_conflict` stays frozen (Route 3's subject, not this one's). Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, lane-4 plan `60daac08…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- The two stale loci read at their bytes: `2026-07-19-mvp-full-worker.md:88`, `:155`; the recorded root supersession at `2026-07-22-relock-lane2-m9-delta.md:373` (§5) — as you established in `…-esc1-m9-invariant-ans-1`.
- Binding owner-path rule quoted verbatim from `master/CYCLE-PLAYBOOK.md:222`.
- Operator decisions + VP approval: `…-esc1-ratify-3/…-031526.md` (`bda1c94197251ff3f8848c3303c3f918dd0302195a5a9c7f28d3ef3aa4334bb8`), approved at `…-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md`.
- The VP's constraint on promoting edited bytes: `…-esc1-ratify-1/RECONCILE-orchestrator-reviewer-20260726-025248.md` (F2).
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No annotation authored, no owner document edited, no mechanism prescribed, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner authors the two provisional annotations in its own document and answers Asks 2 and 3 under a fresh unique DISPATCH_ID parented to this one, naming anything it cannot release unilaterally so master can route it to that owner. Amendment authoring, lane-4 resume, fixture freeze, re-lock, T4 and external use remain held.
