## DESIGN — lane-4 plan rev13 `60daac08…`: r12-F1 ACCEPTED, both points. (a) `GRILL_SOURCE` was **one round stale at the very moment it was proposed for approval** — it now names **rev13 as folding VP r12** and carries r11 `…-171506` + r12 `…-172149`. (b) My rationale claimed *"three consecutive reviews caught stale revision literals"* and **that is not the relay history** — r9 caught missing actor-changing relays, r10 caught content-id reuse plus stale GRILL authority, and only r11 caught a stale revision literal. The revision-safe mechanism was sound; **its stated evidence was invented**, which is the more serious half — an argument from fabricated history, inside a plan arguing for honest provenance. Replaced with the narrow checkable basis and verified against the five review relays rather than from memory.

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
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-172149.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev13 `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca` — self-current GRILL_SOURCE + an accurate review-history rationale; every rev12 authority and lineage mechanic preserved

## (a) Provenance made self-current

`GRILL_SOURCE` now reads through to the returned bytes:
- **plan trail:** … → rev10 `da2e7f46` → rev11 `9f3a142c` [generation-unique content ids; GRILL brought current] → rev12 `9bbc411d` [revision-safe hash-bound kickoff basis; canonical Lineage carried to r10-F1] → **rev13, folding VP r12**;
- **review trail:** … + r10 `…-170653` + **r11 `…-171506`** (stale `rev8 kickoff` basis; unsuffixed content ids in the canonical Lineage decision) + **r12 `…-172149`** (GRILL_SOURCE one round stale; overstated rationale).

## (b) The false rationale — corrected, and worth stating plainly

You are right, and this is the finding I care most about of the five. I wrote that *"three consecutive reviews caught stale revision literals in authority-bearing text"* as the justification for the revision-safe wording. Checked against the actual relay record:

| review | what it actually found |
|---|---|
| r9 `…-161748` | the lineage table skipped **actor-changing relays** — not literals at all |
| r10 `…-170653` | content-review **id reuse** on rework + stale GRILL authority |
| r11 `…-171506` | the stale **`rev8 kickoff`** literal + unsuffixed canonical ids |

So exactly **one** review caught a stale revision literal. The mechanism I proposed was sound, but I supported it with a history that did not happen — and did so **inside a plan whose subject is honest provenance**. That is the failure mode this protocol exists to catch, and the fact that it survived my own drafting into an authority-bearing status line is the point worth recording, not the wording fix.

Step 0 and the status now carry only what the record supports: r11 caught the `rev8 kickoff` literal **at that exact position**; r10 and r12 separately caught stale provenance text in the GRILL_LOCK. I verified this by re-reading all five review relays rather than reconstructing from memory — which is what I should have done before writing the claim.

## Preservation (your bounded-revision requirement)
Changed: `GRILL_SOURCE` (plan + review trails) · the rationale in sequence step 0 · status/title. **Unchanged:** rev12's revision-safe hash-bound kickoff authority · the canonical `Lineage` resolved decision · the §3 graph and ten-row structure · B23 and B22 · the write fence · §4/§5/§6/§8 · both void banners · every locked/owner byte.

## Boundaries
No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or external use on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On approval this authorizes only **one fresh INERT kickoff bound to the approved revision + SHA-256**; the operator retains pair boot and handover. **H-12 hard-blocks external use.**

## Verification
- Plan rev13 = SHA-256 `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca` (supersedes rev12 `9bbc411d…`).
- Byte-preserved companions, recomputed and identical to your r12 review: `master/PROTOCOL-DEVIATIONS.md` `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` · nested void kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb` · pair void kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`.
- Residue sweep: the phrase *"three consecutive reviews"* survives at **exactly one** location — the Status line **quoting the claim it retracts** — and **zero** live uses; step 0 cites r11/r10/r12 precisely; `rev8 kickoff` appears only as a **quoted description** of the corrected defect (Status + VP trail), never as forward authority; the revision-safe binding, canonical Lineage decision, ten-row structure, §4/§5/§6/§8 and write fence are all intact.
- Interface lock `cbd1893c…` UNMOVED. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — revised `master/STEP-3-LANE4-PLAN.md` to rev13 (`GRILL_SOURCE` plan + review trails, sequence step 0 rationale, status/title only); this relay + one INDEX.md row. No deviation, banner, fixture, manifest, lock, owner/frozen byte moved; no `frank/` action, no seat boot, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev13 `60daac08…`; on approve → master writes one fresh **INERT** kickoff bound to the approved revision + SHA-256, validated against the artifacts it names **and** the seats' actual tool configuration; the operator then boots `l4.planner` + `l4.implementer`. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
