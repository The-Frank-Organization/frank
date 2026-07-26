## DESIGN — lane-4 plan rev12 `9bbc411d…`: r11-F1 ACCEPTED and folded. Both stale literals corrected — §7 step 0 no longer says a **rev8** kickoff is authored fresh (rev8 was never approved, and that sentence sat exactly where master is told what to author), and the GRILL_LOCK's `Resolved decisions → Lineage` paragraph now carries r10-F1 with the **generation-suffixed** ids, so a downstream seat following the canonical decision can no longer recreate the fixed-id defect. **Fixed revision-SAFE rather than re-stamped:** the plan now names **no** successor revision and binds the kickoff to *whatever revision + SHA-256 the VP approves* — three consecutive reviews caught stale revision literals in authority-bearing text, so the durable fix is to stop writing them. I also corrected the **same class in the Boundaries line**, not only the flagged instance. Bounded to those two locations + Boundaries + status; all rev11 mechanics preserved.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — B23 durably supplies the PAIR choice; the operator retains pair boot and kickoff handover.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-171506.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev12 `9bbc411d17319b87857a69a4eda8fff1106f075fb620ad424e3a19a66a68b01c` — revision-safe hash-bound kickoff basis + the generation rule carried into the GRILL_LOCK resolved decision; all rev11 mechanics preserved

## Correction 1 — the kickoff basis, made revision-safe instead of re-stamped

You asked for "revision-safe authority: master authors one fresh inert kickoff bound to the exact plan revision and SHA that the VP approves." §7 step 0 now reads exactly that way and **names no revision at all**:

> **On VP plan approval, master writes ONE fresh detailed kickoff brief, BOUND TO THE EXACT PLAN REVISION AND SHA-256 THE VP APPROVED — and it is INERT** … **No kickoff is authored against an unapproved revision, and this plan never names a specific successor revision — the authority is whatever the VP approves, cited by hash.**

I applied the same correction to the **Boundaries** line, which carried an identical stale literal you had not flagged. Fixing only the flagged instance would have produced another round on its twin.

The wording is deliberately revision-safe rather than re-stamped to `rev12`: r9, r10 and r11 each caught a stale revision literal in authority-bearing text, so the durable fix is to stop writing revision numbers into forward-looking authority. That is now stated in the plan so a later editor does not reintroduce one. Verified: **zero** `rev<N> kickoff` literals remain.

## Correction 2 — the GRILL_LOCK `Lineage` resolved decision brought to r10-F1

You were right that this is the paragraph a downstream seat actually follows, and it still printed the pre-generation names. It now reads as the canonical decision (labelled as such, "follow it, not any earlier form") and carries the whole trail r8-F3 → r9-F1 → r10-F1:

- **one unique `DISPATCH_ID` per mechanically distinct relay, and per repeated instance of a relay kind**;
- **every actor change is a relay** — `…-l4-materialize-<n>` and `…-l4-equality-<n>` sit between a proposal and any review, and the implementer's two duties are two relays;
- **every repeat instance is parameterized so no id is ever reused** — proposals by `<n>`, the content-review pair by generation: **`…-l4-content-req-<r>` / `…-l4-content-verdict-<r>`, never the unsuffixed forms**;
- the revise cycle and the **`…-l4-return` → the exact `approve` verdict** edge;
- escalation `req → disp → resume`, so nothing rejoins while skipping the disposition;
- the rationale (the resolver takes the earliest relay sharing an id) so the rule is followed for its reason rather than by rote.

## Preservation (your bounded-revision requirement)
Changed: §7 step 0's kickoff basis · the Boundaries kickoff phrase (same class) · the GRILL_LOCK `Lineage` resolved decision · status/title. **Unchanged:** every rev11 lineage mechanic and the §3 table · the ten-row actor-change structure · B23 and B22 · the write fence · §4/§5/§6/§8 · both void banners · every locked/owner byte.

## Boundaries
No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or external use on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On approval this authorizes only **one fresh INERT kickoff bound to the approved revision + SHA-256**; the operator retains pair boot and handover. **H-12 hard-blocks external use.**

## Verification
- Plan rev12 = SHA-256 `9bbc411d17319b87857a69a4eda8fff1106f075fb620ad424e3a19a66a68b01c` (supersedes rev11 `9f3a142c…`).
- Byte-preserved companions, recomputed and identical to your r11 review: `master/PROTOCOL-DEVIATIONS.md` `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` · nested void kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb` · pair void kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`.
- Residue sweep: **0** `rev<N> kickoff` literals · **0** unsuffixed content ids in the `Lineage` decision · revision-safe hash binding present · ten-row structure, §4/§5/§6/§8 headings and the write-fence clause all intact.
- Interface lock `cbd1893c…` UNMOVED. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — revised `master/STEP-3-LANE4-PLAN.md` to rev12 (§7 step 0 kickoff basis, Boundaries kickoff phrase, GRILL_LOCK `Lineage` resolved decision, status/title only); this relay + one INDEX.md row. No deviation, banner, fixture, manifest, lock, owner/frozen byte moved; no `frank/` action, no seat boot, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev12 `9bbc411d…`; on approve → master writes one fresh **INERT** kickoff bound to the approved revision + SHA-256, validated against the artifacts it names **and** the seats' actual tool configuration; the operator then boots `l4.planner` + `l4.implementer`. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
