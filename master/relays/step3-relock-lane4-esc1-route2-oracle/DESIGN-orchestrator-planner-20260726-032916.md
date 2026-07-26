## DESIGN — ROUTE 2 of 5, JOINT to m-9 + m-3: define the **direct-prefix oracle**. The operator chose a direct whole-prefix comparison over m-3's proposed ordered-list form, which is a valid decision that left the evidence mechanics newly open — and I routed five things while routing this to nobody (VP4-F1). **The shape genuinely changes:** `log_prefix_digest` is today a **typed digest member** of `resume_prefix_expectation`, whereas a direct comparison is a **predicate over two bounded artifacts**. And the VP's sub-catch, which neither of us raised: frank **keeps appending after it resumes**, so "compare the log" is under-specified by construction — the return must prove the comparison is over the **frozen interval, not the post-resume append tail**.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-route2-oracle
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-vp-review-5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the return decides whether a ratified `resume_prefix_expectation` member is removed, renamed or replaced, which is amendment-shaped (Master+VP+operator). This relay asks; it authors no contract and moves no owner or locked byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-3.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-3.implementer, m-10.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: Route 2 JOINT — define the direct-prefix comparison contract: extraction boundary · canonical byte/record representation · independently content-addressed expected artifact · actual evidence locator · mismatch diagnostic · manifest/predicate shape · explicit disposition of `log_prefix_digest` (removed/renamed/replaced) · and PROOF the comparison binds the frozen interval rather than the append tail

m-9, m-3 — a joint act, both of you in `TO` because both halves are owed and CC creates no obligation.

## What the operator decided, and what it left open

**Decision 4:** `log_prefix_digest` collapses to **one whole-prefix comparison** — the produced log must equal the frozen expected — rather than an ordered per-record digest list. The reasoning was the operator's and it is sound: the harness authors the record contents (condition (iv), which m-9 confirmed), so it holds **both sides**; a digest is a compact stand-in for content you do not have, and diagnosis on failure comes from diffing two files already in hand.

**What that left open.** m-3's bound return proposed the ordered `{seq, record_digest}` / per-round list form and explicitly said it was naming a proposed replacement, not authoring the amendment. The operator chose otherwise. So the evidence mechanics are newly unowned, and I compounded it by not routing them.

## The contract owed, and the ownership split

**m-9 owns** valid-prefix extraction and the canonical journal byte/record representation. **m-3 owns** the E3 predicate, expected/actual independence, and the evidence locator. Please return one contract covering all of:

1. **The extraction boundary** — what exactly constitutes "the valid prefix" for comparison purposes, and how its end is determined.
2. **The canonical representation** — the exact byte or canonical-record form compared. If any member is non-deterministic or engine-supplied, name it. *(m-9 already flagged `ts_monotonic` as safe only because the prefix is read and replayed, never re-stamped — that reasoning should be restated for whichever representation you choose.)*
3. **The expected artifact's source and content address** — how the frozen expected side comes to exist, and how it is independently addressed. Condition (iv) freezes the **authored record contents** and the harness derives from them; the expected artifact must be a closed function of frozen bytes and **must not be harvested from a build run**, or the oracle is circular.
4. **The actual evidence locator** — where the observer reads the produced artifact from, in m-3's recorder terms.
5. **The mismatch diagnostic** — what the leg reports on failure. The operator's rationale for accepting the loss of per-position granularity was that a diff of two held files supplies it; please state where that diff is produced and by whom, so it is designed rather than assumed.
6. **The manifest/predicate field shape** — and say plainly whether `log_prefix_digest` is **removed, renamed, or replaced by a closed expectation object.** It is currently a typed digest member (`STEP-3-LANE4-PLAN.md:79`); a predicate over two artifacts is not a digest value, so leaving the member as-is would be a shape mismatch.

## The frozen-interval proof — the part I would not have caught

**frank keeps appending after it resumes.** So an unbounded "compare the log" would compare a file that has *legitimately grown* past the resume point, and would either fail spuriously or be silently re-scoped to whatever the reader happened to read. **The return must prove the comparison binds the intended frozen interval and not the post-resume append tail** — a bound derived from frozen fixture material, not from where the file happens to end at read time.

This is a circularity risk inside the decision that was meant to remove circularity, and it is load-bearing. Credit to the VP for finding it (LANE4-ESC1-VP4-F1).

## What must NOT come back

Not the ordered digest list in another costume — the operator rejected it, and reintroducing per-record digests as the comparison term would be re-deciding a made decision. If you conclude the direct comparison **cannot** be specified soundly without per-record structure, **say that explicitly as a blocker** and name why; that routes back to the operator as a re-decision rather than arriving as a quiet restoration.

Also not asked: the edited-session state machine (Route 3), the fencing predicate (Route 4), the member-set removal (Route 5), or any fixture authoring — lane 4 stays held and will bind a *fresh* plan revision after the amendment, so nothing you return here needs to accommodate the current stale plan text.

## Boundaries
This relay ratifies nothing, authors no contract or predicate, changes no fixture or manifest, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, lane-4 plan `60daac08…` (byte-unmoved; its stale loci await the successor revision), m-9 delta `01b885fe…`, m-3 r24 `651c9aec…`, m-10 rev16 `3e3c5192…`, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- Operator decisions + VP approval: `…-esc1-ratify-3/…-031526.md` (`bda1c94197251ff3f8848c3303c3f918dd0302195a5a9c7f28d3ef3aa4334bb8`); approval at `…-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md`; the F1 finding and its frozen-interval sub-catch at `…-esc1-ratify-2/RECONCILE-orchestrator-reviewer-20260726-030947.md`.
- m-3's proposed list form at its bound bytes `…-esc1-m3-scope-ans-1/DESIGN-planner-20260726-003000.md` (`2720d109…`); m-9's condition-(iv) acceptance and the authorable-vs-derived split at `…-esc1-m9-conditions-ans-1/DESIGN-planner-20260725-225330.md` (`6f4e4f15…`).
- `log_prefix_digest` as a typed member at `master/STEP-3-LANE4-PLAN.md:79`; plan hash `60daac08…` unmoved.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No contract authored, no member removed/renamed, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner + m-3.planner return one joint direct-prefix-oracle contract under a fresh unique DISPATCH_ID parented to this one, including the frozen-interval proof and the explicit disposition of `log_prefix_digest`. Master folds the exact return into the additive supersession record. Amendment ratification, fresh lane-4 plan, lane-4 resume, fixture freeze, re-lock, T4 and external use remain held.
