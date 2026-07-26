## DESIGN — lane-4 plan rev10 `da2e7f46…`: r9-F1 ACCEPTED and folded. You caught that my "exhaustive" table listed the relays I had thought of rather than the ones the workflow requires — it jumped planner-proposal → planner-review-request, **skipping master's materialization** (where the actor changes), and gave one generic implementer verdict although **byte-equality and content review are two distinct duties at two distinct points**, so `…-l4-return` could not mechanically prove it descended from the content approval; the escalation branch had no relay parenting to its disposition. §3's table now carries **all ten actor-changing rows** plus explicit rework edges, with the principle stated: **an unrecorded actor change is an unrecorded handoff — operator prompting and filesystem visibility are never the handoff.** Bounded exactly as directed: only the lineage table, matching §7 text, the GRILL_LOCK lineage line and the status changed. **B23, B22, the write fence, §4/§5/§6/§8 and both void banners are byte-preserved.**

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator's PAIR choice is durable at B23; the operator retains the pair-boot and kickoff-handover gate.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-161748.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev10 `da2e7f46c7fe303fac98778c3b5dcd556e83ff5d0be4f8111a3e03a6cb71fe4f` — the complete materialization / equality / content-review / escalation-resume chain; all other rev9 decisions preserved

## The corrected chain — all six required additions, in your order

| # | relay | actor | `DISPATCH_ID` | `PARENT_DISPATCH_ID` |
|---|---|---|---|---|
| 1 | kickoff handover | master | `…-l4-dispatch` | `step3-relock-lane4` |
| 2 | proposal *n* (artifact; at *n*=M the manifest) | `l4.planner` | `…-l4-propose-<n>` | `…-l4-dispatch` (*n*=1) else `…-l4-equality-<n-1>` |
| 3 | **materialization receipt *n*** — path, recomputed byte length + SHA-256, disk-action evidence *(your item 1)* | **master** | `…-l4-materialize-<n>` | `…-l4-propose-<n>` |
| 4 | **equality confirmation *n*** — duty **i** *(item 2)* | `l4.implementer` | `…-l4-equality-<n>` | `…-l4-materialize-<n>` |
| 5 | **final content-review request** — only once every fixture **and** the manifest carry an equality confirmation *(item 3)* | `l4.planner` | `…-l4-content-req` | `…-l4-equality-<M>` |
| 6 | **final content-review verdict** — duty **ii**, durable approve/revise *(item 4)* | `l4.implementer` | `…-l4-content-verdict` | `…-l4-content-req` |
| 7 | pair → master return / SITREP *(item 5)* | `l4.planner` | `…-l4-return` | `…-l4-content-verdict` — **the content approval, never an earlier or generic verdict** |
| 8 | escalation *n* | either seat | `…-l4-esc<n>-req` | the exact relay that raised it |
| 9 | escalation disposition *n* | master / owning m-x planner | `…-l4-esc<n>-disp` | `…-l4-esc<n>-req` |
| 10 | **escalation resume *n*** — names the stage it rejoins and the exact relay it continues *(item 6)* | the resuming seat | `…-l4-esc<n>-resume` | `…-l4-esc<n>-disp` |

**Rework edges named explicitly**, so no rework silently re-parents: a **failing** equality confirmation is the parent of the corrective `…-l4-propose-<n'>`; a content verdict of **revise** is the parent of the first corrective `…-l4-propose-<n'>`. Rework never rejoins by pointing at a superseded predecessor.

## The general lesson I took from this finding

My previous two lineage attempts were wrong in the same way for the same reason: I enumerated the relays I had in mind rather than deriving them from the workflow, so both times the missing rows were exactly the ones where **the actor changes** — first the unique request/verdict split, now materialization and the two-duty split. Recorded in the GRILL_LOCK as a derivation rule rather than another patch: **every actor change in the sequence is a relay with its own id, or it is an unrecorded handoff.** That also disposes of the "do not rely on operator prompting or filesystem visibility as an unrecorded handoff" instruction — it is now stated in the plan text, not just honoured by the table.

## Preservation (your bounded-revision requirement)
Changed: §3's lineage table · the matching §7 sequence steps 3–4 · the GRILL_LOCK lineage line · the status/title. **Unchanged:** B23 and B22 · the §3 write fence · §4 (exact §7 schema + fixed values) · §5 (carried obligations + owner-fidelity matrix) · §6 (Master+VP reservations) · §8 · both void banners (byte-identical, no rev pointer needed) · every locked/owner byte.

## What I ask the VP to review
- Does the chain now represent the actual workflow — is every actor change a relay, with no remaining implicit handoff?
- Are the rework and escalation-resume edges correct, and does `…-l4-return` provably descend from the content approval?
- Is the revision genuinely bounded — nothing else moved under cover of it?

## Boundaries
No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or external use on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On approval this authorizes only a **fresh INERT kickoff**; the operator retains pair boot and handover. **H-12 hard-blocks external use.**

## Verification
- Plan rev10 = SHA-256 `da2e7f46c7fe303fac98778c3b5dcd556e83ff5d0be4f8111a3e03a6cb71fe4f` (supersedes rev9 `a0009ef9…`).
- Byte-preserved companions, recomputed: `master/PROTOCOL-DEVIATIONS.md` `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` · nested void kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb` · pair void kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af` — all three identical to the values in your r9 review.
- Chain present: all ten ids appear in §3; §4/§5/§6/§8 headings intact; the write fence clause and the "`IN_REPLY_TO` never a gate input" rule both retained.
- Interface lock `cbd1893c…` UNMOVED. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — revised `master/STEP-3-LANE4-PLAN.md` to rev10 (§3 lineage table, §7 steps 3–4, GRILL_LOCK lineage line, status/title only); this relay + one INDEX.md row. No deviation, banner, fixture, manifest, lock, owner/frozen byte moved; no `frank/` action, no seat boot, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev10 `da2e7f46…`; on approve → master writes the fresh **INERT** rev10 kickoff, validated against the artifacts it names **and** the seats' actual tool configuration; the operator then boots `l4.planner` + `l4.implementer`. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
