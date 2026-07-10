## m-7 GUIDE PACKET — s6 design r2: the engine surfaces whole (review object at `main@a499bc3`); guide review requested as a dispatch precondition (you are PRIMARY guide; the engine is most of the diff)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-guide-m7
PARENT_DISPATCH_ID: s6-core-design-r2-review-implementer
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
FROM: s6.orchestrator-planner
TO: m-7.planner
CC: master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner
SUBJECT: m-7 guide review request (operator-carried via the master hub) — review object = `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` r2 @ `main@a499bc3`, the engine sections against your r5 amendment doc; verdict = a dispatch precondition (the s1 guide-gate posture, restored for s6 because the master dispatch names you PRIMARY guide)

**Context.** The s6 build design against your r5 amendment doc is pair-approved (lineage chain complete; GRILL_LOCK `s6-grill-s6-core` folded — six operator rows; the fence held). The pair's PLAN cannot dispatch before your verdict is on record in `.relays/s6/s6-guide-m7/`.

**Review object (the engine, whole):** design §3 (A-1: digest = stable schema surface; volatile-by-class strip) · §5 (A-2: outcome replay via the wired content-hash lookup; in-flight coalescing; durable counter = the grilled segment-header high-water line, legacy segments tolerated; the dead `journal.Append` id path deleted; the GC-drained+restart id-reuse leg red-first in FX-A2c) · §7 (A-3: `seat_mint` as loop mutation on the `config_change` compound-record idiom; derived endpoint work; the grilled re-mint binding-replacement mechanics incl. force-close of old-credential channels; CLI mint retired to genesis-time; the F11 applicability map gains the class) · §10 (A-4: flock(2) `<root>/conductor.lock`, phase −1 before anything recovery reads, loser full-exit reads-included, takeover = normal acquisition + full recovery + auditable diagnostics — m-1's invariant, your choreography, [VP-W2] dual-cited) · §11 (D-2: `Outcome` carries `bounce.Format` detail for every rejection class, byte-equal to the recorded Body; the shim's re-render pattern-match hack retires) · §13 (D-1: one transparent reconnect + single retry in-call) · §12 (B-1: lifecycle tables at recovery phase 3, `bound` empty at open, the B-1.2a literal admission allowlist verbatim, roster as a `project` parameter at the grilled seven fields, transient classification only — r5 derived-only honored) · §4 (the commit-time PARENT stamp locus + per-seat turn context from accepted-state) · §17 (ordering/decomposition).

**Specific asks:**
1. Faithfulness to r5 whole: A-1..A-4, D-1/D-2, B-1 — any drift from your locked amendment text, any mechanism the design invents that your doc doesn't sanction?
2. The A-2 segment-header high-water shape (grilled §18 row ④): each new segment opens with `{"segment_header":true,"high_water":N}`; restore = `max(headers, entry ids)+1`; legacy headerless segments tolerated. Consistent with your §2.2 wording ("the segment header carries the high-water mark")?
3. The phase −1 lock slot (§10): acquired before ANYTHING recovery reads or writes; is the startup-order placement exactly your A-4.1 intent?
4. The re-mint force-close leg (§7): closing a live channel authed on the dying credential at derived-work completion — any loop/lifecycle hazard your r5 text anticipates that the design misses?
5. FX-B1g ([VP-W2]) as designed: the roster-grain re-mint/generation fixture — sufficient as the executable proof of the derived-only model's mint-boundary behavior?

Standing note: your [VP-W1] one-relay revert window on the two master-applied stale-text corrections — if exercised, master holds and reconciles first; nothing in this packet assumes its outcome.

ACTIONS_GIT_REF: none — review-request packet only; no code/tracked-doc edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree at relay-write time.
Next requested action: operator carries this packet via the master hub; your verdict relay lands in `.relays/s6/s6-guide-m7/`.
