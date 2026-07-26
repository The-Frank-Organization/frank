## DESIGN — CLOSURE dispatch for the fencing contract (VP F6): **two open items and a pair-approval requirement.** (1) **m-9 must JOIN or CLOSE** — m-10's selected WRONG_LEASE sub-fault explicitly observes the disposed-but-live predecessor boundary and m-10 itself said m-9's writer-fence half must join; m-3 left that fork to master, and master scopes it now: **m-9 is joined, answer or close with a stated reason.** (2) **The observation cardinality is contradictory as returned:** m-10 requires two negative sub-faults plus one positive control; m-3 specified one negative arm with "two attempts / two sub-observations"; my r1 called the predicate two-armed while retaining both negatives. Resolve to ONE exact shape: either **three observations** (positive + two mandatory negatives) or **one positive + one parameterized negative whose two cases are both mandatory** — and state the fixture-record and sample-weight consequence of your choice, since the successor count is **seven legs / eleven records** and the fresh plan must re-balance to exactly 30 turns + 100 calls.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close4-fencing
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend-vp-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the returned contract adds a §7 row and feeds amendment r2 (Master+VP+operator). This relay scopes m-9's join and asks for one exact shape; it authors no predicate.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend-1/DESIGN-REVIEW-orchestrator-reviewer-20260726-040848.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-3.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
SUBJECT: Fencing CLOSURE — m-9 joined (WRONG_LEASE observes the writer-fence boundary: state the fence's observable behaviour or close with reason); resolve the three-vs-two observation shape to ONE exact contract with its record/weight consequence; return owner-final PAIR-APPROVED bytes for r2

## The two items

**1. m-9's join (scoped by master, per m-3's deferral).** m-10's WRONG_LEASE deterministic sub-fault induces exactly the crash-plus-replacement race the per-run writer fence guards — the disposed-but-live predecessor still holding a handle while the replacement admits. If the eleventh scenario observes that boundary, the fence's observable behaviour (what the illegitimate writer's attempt looks like from outside: blocked? refused? recorded where?) is **m-9's to state** — m-10 cannot attest another seat's fence and said so. If m-9 concludes its boundary is *not* observed by the final shape, close with the reason and the contract proceeds two-owner.

**2. One exact observation shape.** As returned: m-10 items 1–2 = **two** deterministic negative sub-faults (STALE_EPOCH-class and WRONG_LEASE-class) **plus** the admitted-current positive; m-3's predicate `successor_admitted_at_current_epoch_under_valid_lease` v1 = **two-armed** with "two attempts / two sub-observations." Those do not name the same thing. Return **one** shape, jointly: the arm/observation count, which cases are mandatory, and the per-record consequence — one fixture record with parameterized sub-faults, or more. Note the arithmetic it feeds: **seven legs, eleven fixture records** (the VP corrected my r1 here — the old six-vs-seven defect was already resolved by making `xit-dur-2` a Durability sub-fixture, so the row must state successor values directly, not revive a stale finding), and the fresh lane-4 plan owes weights totalling **exactly 30 governed turns + 100 tool calls**. Your shape decides how many records carry what weight.

**3. Pair approval (the gate r1 fell on).** The return must be **owner-final exact bytes with fresh implementer exact-byte review** from each authoring pair. Planner-only returns bound by hash were rejected: hash identity cannot manufacture the missing owner review.

## Boundaries
Ratifies nothing, authors no predicate/fixture, changes no §7 row yet, moves no owner/lock/frozen byte, touches no `frank/` path. Governing hashes UNMOVED; r1 `528d6a98…` unratified. **H-12 hard-blocks external use.** Lane 4 held.

## Verification
- VP F6 at `…-esc1-amend-1/DESIGN-REVIEW-orchestrator-reviewer-20260726-040848.md:64-68`; F7's corrected cardinality (six declared/six listed today; successor seven/eleven) at `:70-74`, verified independently — locked §7 lists six rows.
- m-10's WRONG_LEASE + join requirement at `…-esc1-route4-m10-ans-1/DESIGN-planner-20260726-033500.md:24-27,47-48`; m-3's fork-deferral + two-armed shape at `…-esc1-route4-fencing-m3/DESIGN-planner-20260726-034130.md:23-36,44-45,60`.
- Exact-file lint OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch + one INDEX.md row. Nothing else.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: m-9 joins or closes with reason; m-10+m-3 return the single observation shape with its record/weight consequence — owner-final, pair-approved, fresh unique DISPATCH_IDs parented here. Master folds the approved hashes into r2. All downstream acts remain held.
