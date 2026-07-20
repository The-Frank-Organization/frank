## PLAN — FENCE EXPANSION GRANTED (r2): exactly the two test-support rows at the named-seam grain — WITH one binding tripwire (lock-pinned values are never licensed by a fence grant; a red on a lock-pin is an amendment question, not a fixture edit) and the phantom-seat disposition (the T1 review obligation lands on s8.planner; the seat-list correction rides the next plan touch)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s8-build-escalate-fence-r2-ruling
PARENT_DISPATCH_ID: s8-build-escalate-fence-r2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the fence ruling is the orchestrator's per the dispatch conditions (operator CC'd; the reviewer-seat fact below is consumed AS operator-confirmed, not inferred); the slice merge stays separately operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
FENCE_EXPANSION_AUTHORIZED: granted — exactly two test-support rows, named-seam grain, tripwire binding
IN_REPLY_TO: master/relays/s8-build-escalate-fence/SITREP-implementer-20260711-202000.md
FROM: master.orchestrator-planner
TO: s8.implementer
CC: operator, master.orchestrator-reviewer, s8.planner, m-3.planner, m-7.planner, m-2.planner
SUBJECT: both rows granted exactly as requested — the reasoning is honest on all three no-workaround legs (a v6 marker with a v5 assertion is red by construction; a legacy-store serve exemption would re-open the fail-open hole m-7's correction just closed; a test-env bypass is the test/runtime mismatch the protocol rejects) — m-7's fail-closed `store-not-adopted` correction is endorsed as the Rail-A-correct call (closed where ignoring changes acceptance MEANING); continue under the conditions below

**GRANT (named seams, nothing wider):**
1. `frank/internal/fieldspec/registry_test.go` — exactly the ONE existing registry-version assertion, `s7a-fieldspec-v5` → `s8-fieldspec-v6` (m-2's reviewed Site 1). No other assertion in the file is licensed.
2. `frank/test/fixtures/s2setup_test.go` — exactly the shared fresh/production fixture config SOURCES to the lawful s8 three-member shape (engine `version: 1` + `present_layers.observe: false` + the pinned catalog source), so the existing production-serve fixtures ride a lawful store behind m-7's fail-closed guard. The two-member legacy expectation lives ONLY in the already-IN `s8_adoption_test.go`, as you stated; no recovery fixture is repurposed.

**THE BINDING TRIPWIRE (rides this grant and every future one — named now because the v6 move sits one step from it):** a fence grant licenses FILES for NAMED seams; it never licenses moving a **lock-pinned value**. The r13 lock pins the genesis fieldspec bytes at the `s7a-fieldspec-v5` SHA (§5.1, FX-CFG-7's hard-coded hash), the capability table's exact marker sets (§2.5), the terminal enum, and the catalog censuses. The v5→v6 reconcile must flow through the locks' OWN mechanisms — m-2's owner-supplied forward relation + the §2.4 acceptance gate + the m-7-confirmed T1 grains (reader capability before record, as designed). If the honest build result is that a LOCK-PINNED text itself must change (the genesis pin, the supported-set literal, an FX pin), that is a **lock-amendment escalation to the owning pair + master — never a fixture edit that makes the battery green**. A red on a lock-pin assertion is information, not friction. State in your next SITREP which side of this line the v6 sequencing landed on (genesis-at-v5 + live transition, or a pinned-v5 source artifact, or an escalation).

**Phantom-seat disposition (consumed as operator-confirmed fact; the obligation does not vanish):** there is no `s8.reviewer` — accepted. Effects, in order: (i) the T1 review request withdrawn from the nonexistent address is NOT dropped — **the review obligation lands on s8.planner** (the pair floor is the standing review gate; planner review at the plan's cadence, T1 included retroactively); (ii) the plan-of-record's seat list takes the one-line correction at the next plan touch (the pair's trail or m-3's doc — either, visibly); (iii) all future CC lines drop the phantom — including mine, which propagated it in the two prior rulings; owned here for the trail.

**Standing, unchanged:** the four locks, SEQ-1 + the addendum steer, conditions (a)–(g), the prior two-row grant's seam limits, and the operator-only merge gate. Your stop discipline (no OUT edit, token-clean, Rail A/B applied in the request itself) is noted approvingly — second clean catch of the slice, this one mid-implementation.

ACTIONS_GIT_REF: none — a fence ruling; no `frank/` edit (disk refs: this relay + one INDEX.md row timestamped 20260711-202010; stamped after the replied-to filename per the skew convention — author wall clock ~1.5h behind the s8 session).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; the build branch state is the pair's to report (last reported: `s8-observe-spine` with T1 commit `d87dc21`, tree clean).
Next requested action: operator carries this to s8.implementer (CC s8.planner — the T1 review obligation lands there); the implementer folds m-7's T1 correction test-first, applies m-2's four sites + the one assertion byte, runs the full uncached battery, reports with the tripwire-line statement; the pair's cadence continues.
