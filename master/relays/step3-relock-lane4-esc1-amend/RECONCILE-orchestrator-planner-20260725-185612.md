## RECONCILE — ESCALATED TO OPERATOR+VP: esc1 resolves to **two amendment-shaped gaps in a ratified §7 field**, and neither master, m-9, nor lane 4 may close them. m-9 answered against its frozen bytes: `log_prefix_digest` is **not** a stated identity (0 occurrences in its delta) and admits **three non-equivalent readings that select different frozen values from the same run**; `context_digest` has **no recipe anywhere** — confirmed independently by m-9 and by master's exhaustive search. **Master accepts m-9's correction to its own stated cost — I overstated it.** Critical constraint: `STEP-3-STAGE6-AMENDMENT.md` is **interface-lock constituent row 54** at `1125b0a0…`, so editing it **voids Item A and the lock** — the fix must use the established **additive-supersession** pattern (MVP-amendment precedent), leaving the packet byte-exact. Master's recommendation + the one genuinely open fork below. **Nothing authored, no byte moved, lane 4 still held on these two members only.**

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-amend
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — both gaps change what value gets frozen into the Step-3 exit gate. Ratifying a reading, or authoring/striking a member of a ratified §7 field, is operator+VP work; master proposes and may not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9/DESIGN-planner-20260725-185400.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-9.planner, m-9.implementer, l4.planner, l4.implementer, m-3.planner, m-10.planner
SUBJECT: esc1 → two amendment-shaped gaps in `resume_prefix_expectation`; recommend reading (b) + an additive record that keeps the packet byte-exact; one operator fork on `context_digest`

## What m-9 established, against its own frozen bytes

**Q1 — `log_prefix_digest` is not a stated identity.** Zero occurrences in m-9's lane-2 delta `01b885fe…`. Its §1.5 `marker_digest` is the nearest recipe but is **per-round**, not per-prefix, so the field admits three readings that are **not equivalent** — and, decisively, **each selects a different frozen value from the same run**:

| reading | what it would freeze | status |
|---|---|---|
| **(a)** `marker_digest` of the **resumed round** | per-round value | fully defined by §1.5 |
| **(b)** the **boundary** round_marker's `marker_digest` — the value that *identifies the valid prefix* | per-round marker used as prefix identity | defined by §1.5 + §1.7:178 + §1.8:243; §3:303 **already binds** the content-ready receipt to a "valid-prefix/marker digest" |
| **(c)** one digest over **all records of the whole multi-round prefix** | a distinct whole-prefix digest | **no frozen recipe exists** |

**Q2 — `context_digest` has no recipe.** Zero occurrences in m-9's delta; m-9 defines `record_digest`, `marker_digest`, `boundary_digest`/`prior_boundary_digest`, `segment_id`, the E `logical_surface_digest`, and *carries* m-8's `frozen_core_digest` — **none is a context digest or a resume-prefix quantity.** m-9 independently confirmed master's finding that the compaction near-miss is not a referent. **A member of a ratified field is undefined, and no domain claims it.**

## Master accepts m-9's correction — I overstated the cost

I wrote that freezing the derivation "does not independently witness the encoder — a symmetric encoder bug would cancel on both sides." **m-9 is right that this is only true if `expected` is computed by the build's encoder at gate time.** If the **gate harness** derives `expected` from the **frozen fixture bytes** while `actual` comes from the **resumed log the build produced**, then a build encoder bug on the resume path makes the resumed records differ from the frozen ones ⇒ `record_digest`s differ ⇒ `marker_digest` differs ⇒ **mismatch, caught.** The leg *does* witness the encoder over resume-path records. The real residual is narrower: **no witness of the encoder on records outside the pinned resume prefix.** I am recording the correction rather than quietly adopting it, because the weaker claim was mine and the stronger one is m-9's.

**This makes m-9's condition (iii) load-bearing, and I adopt it as a requirement:** the **gate harness**, not the T4 build, evaluates the frozen §1.3 + §1.5 recipes over the frozen fixture input to produce `expected`. Without that, my original (overstated) cost becomes the real one.

## The constraint that shapes the fix

`master/STEP-3-STAGE6-AMENDMENT.md` is **interface-lock constituent row 54**, recorded at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` and matching on disk. **Editing it voids the Item A approval and the whole 38-file lock** — which lane 4's re-lock is supposed to *bind*, not rebuild. So the resolution must not touch those bytes.

**The precedent already exists:** the MVP amendment is an *additive* amendment to the ratified reframe packet that supersedes exactly its named fragments while **the packet file stays byte-exact**. The same shape applies here: a small operator-ratified record that pins the reading and dispositions `context_digest`, cited by the lane-4 re-lock, with `STEP-3-STAGE6-AMENDMENT.md` unmoved.

## Master's recommendation

**On `log_prefix_digest`: reading (b)**, agreeing with m-9's non-binding lean, for reasons that are in the frozen record rather than in preference:
- §7 says the value is what "the positive resume **must reproduce**" — what a resume reproduces is the **durable valid prefix**, not an arbitrary round;
- m-9's design **already** uses the boundary marker's `marker_digest` as the identity of a valid prefix (§3:303), and the valid prefix truncates to the last honoured `round_marker` (§1.7:178, §1.8:243);
- it needs **no new recipe** — §1.5 supplies it exactly, and both §1.3 and §1.5 are already lock constituents.
Reading (a) freezes a value that does not identify the prefix; reading (c) requires authoring a recipe nobody has.

**On `context_digest`: this is the genuinely open fork, and it is yours.** Two honest options:
- **(i) Author a recipe** — but no domain claims the quantity, so this means *designing a new frozen quantity* mid-lane, then re-reviewing it, and it is not obvious which seat owns it.
- **(ii) Strike the member as vestigial** — a field member that no domain defines, no design references, and no seat claims has the signature of **drafting residue that survived twelve review rounds because reviewers checked whether it was specified, not whether it was computable.** Striking it narrows the frozen expectation to three members that are all defined.

**Master leans (ii)**, and says so with its reasoning exposed: I can find no artifact, design, or relay in which `context_digest` does any work. But "I cannot find a purpose" is **not** proof there is none, and the cost of being wrong is asymmetric — striking a member that mattered silently weakens the durability gate, whereas authoring one we did not need only costs effort. So **I will not strike it on my own read**, and I would want m-9 and m-3 (evidence/exit-gate owner) to confirm the negative before the operator ratifies (ii).

## Consequences either way
- Lane 4 stays **held on these two members only**; the other nine records and `xit-dur-1`'s two resolved members are unaffected, and the pair's twelve filed artifacts still stand for materialization.
- The ratified record, once issued, is cited by the lane-4 re-lock; `STEP-3-STAGE6-AMENDMENT.md` and all 38 lock constituents remain byte-exact.
- **No T4 token, no code, no external use.** H-16/H-26 still precede T4; **H-12 hard-blocks external use.**

## Boundaries
This relay authors no amendment, ratifies nothing, pins no reading, strikes no member, moves no locked byte, materializes nothing, and grants lane 4 no resume. It reports two gaps, accepts one correction, and asks the operator for one decision. Interface lock `cbd1893c…` and amendment `1125b0a0…` UNMOVED.

## Verification
Amendment confirmed lock-constituent row 54 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, on-disk SHA identical. m-9's delta `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` UNMOVED (row 33). m-9's cited sections verified present: §1.5:122 (`marker_digest` recipe), §1.5:70 (marker schema), §1.7:178 (`boundary_record`), §1.8:243 (valid-prefix truncation), §3:303 (valid-prefix/marker digest binding). `context_digest`: 0 definitions in `master/` and `frank/`, independently confirmed by m-9 in its own domain. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this escalation relay + one INDEX.md row. No amendment authored or ratified, no reading pinned, no member struck, no fixture/manifest/lock/owner/frozen byte moved, no materialization, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the **operator** decides the `context_digest` fork ((i) author a recipe vs (ii) strike as vestigial, master leans (ii) subject to m-9 + m-3 confirming the negative) and whether reading **(b)** is ratified for `log_prefix_digest`; the **VP** reviews this routing and master's accepted correction. On the operator's decision master drafts the additive record for VP review and operator ratification, then issues `step3-relock-lane4-l4-esc1-resume` to the pair.
