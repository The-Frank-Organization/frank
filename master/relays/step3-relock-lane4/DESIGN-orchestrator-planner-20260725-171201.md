## DESIGN — lane-4 plan rev11 `9f3a142c…`: both r10 gates ACCEPTED and folded. **F1** — the rework branch wasn't compatible with my own rule: a `revise` cycle needs a *second* content request/verdict, but rev10 gave those relays **fixed** ids, so the second review would either reuse them (reviving the earliest-relay resolver defect) or have no ids at all. Content review is now parameterized by **generation `<r>`**, exactly as proposals are by `<n>`, with the full revise cycle spelled out and `…-l4-return` parented to the verdict **whose value is `approve`**. **F2** — the GRILL_LOCK's authority block still stopped at rev8/r7, omitted B23 and r8–r10, said staffing was re-answered at B22 (**B22 left it open; B23 supplied it**), and carried the superseded `one gated leg = one dispatch id` shorthand; all corrected, since stale text in a GRILL_LOCK is authority-bearing, not history. Bounded exactly as directed — the ten-row structure, B23, B22, the write fence, §4/§5/§6/§8 and both void banners are byte-preserved.

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
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-170653.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev11 `9f3a142c2a26fdd21bf890f5c7da193cef936eec7aa231e72e7969a4ee1fdfc2` — generation-unique content-review ids + a current B23/rev11 GRILL_LOCK; all other rev10 decisions preserved

## F1 — the rework branch, corrected in your four steps

1. **Parameterized:** `…-l4-content-req-<r>` / `…-l4-content-verdict-<r>`.
2. **Request `<r>`** parents to the **last required equality confirmation of that generation** (`…-l4-equality-<M>` at *r*=1; after a revise, the last post-correction `…-l4-equality-<n'>`); **verdict `<r>`** parents to request `<r>`.
3. **On `revise`,** the first corrective `…-l4-propose-<n'>` parents to `…-l4-content-verdict-<r>`; the next request is `…-l4-content-req-<r+1>`, parented to the last required post-correction equality confirmation. The full cycle is written out in §3: `verdict-<r>` (revise) → `propose-<n'>` → `materialize-<n'>` → `equality-<n'>` → `content-req-<r+1>` → `content-verdict-<r+1>`.
4. **`…-l4-return`** parents to the exact `…-l4-content-verdict-<r>` **whose value is `approve`** — never a revised verdict of an earlier generation.

The ten-row table remains a table of relay **kinds**, as you allowed; the rule now stated above it is that **a kind may instantiate more than once, but every instance must be mechanically unique** — proposals by `<n>`, content-review attempts by generation `<r>`, and **nothing is ever re-filed under a used id**.

## F2 — the GRILL_LOCK brought current

- **Plan trail** carried through rev8 → rev9 → rev10 → rev11, each with its hash and what it did.
- **VP trail** carried through r8 (`…-160401`), r9 (`…-161748`) and r10 (`…-170653`) with their verdicts; the r7 APPROVE is marked superseded for structure and transport, with its F1/F2 findings retained and strengthened.
- **Staffing provenance corrected:** cite **B23 for staffing**, **B22 for transport only**, with B22's open item and rev8's wrong claim both named. `re-answered at B22` is gone (grep 0).
- **Design-lock impact** replaced: the `one gated leg = one dispatch id` shorthand is out, replaced by one unique id per mechanically distinct relay, every actor-changing kind enumerated, and every repeated instance parameterized so no id is reused.

You were right to treat this as authority-bearing rather than cosmetic: the resolved-decision body and the summary block disagreed, and the summary is what a downstream seat inherits.

## Note on the pattern, since this is the third lineage revision

r8-F3, r9-F1 and r10-F1 were the same error in three places: I described the lineage I had pictured rather than deriving it from the workflow, so each round the missing relays were exactly the ones I hadn't pictured — first the request/verdict split, then the actor changes (materialization, the two duties), now the repeat instances. rev11 states the **derivation rules** rather than a longer list: every actor change is a relay; every repeat instance is parameterized; no id is ever reused. If the graph is still incomplete, it should now fail against a rule rather than against an omission.

## Preservation (your bounded-revision requirement)
Changed: the content-review id/parent templates in §3 (+ the rework-generation clause) · §7 step 4 text · `GRILL_SOURCE` · the design-lock impact line · status/title. **Unchanged:** the ten-row actor-change structure · B23 and B22 · the §3 write fence · §4/§5/§6/§8 · both void banners · every locked/owner byte.

## Boundaries
No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or external use on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On approval this authorizes only a **fresh INERT kickoff**; the operator retains pair boot and handover. **H-12 hard-blocks external use.**

## Verification
- Plan rev11 = SHA-256 `9f3a142c2a26fdd21bf890f5c7da193cef936eec7aa231e72e7969a4ee1fdfc2` (supersedes rev10 `da2e7f46…`).
- Byte-preserved companions, recomputed and identical to your r10 review: `master/PROTOCOL-DEVIATIONS.md` `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` · nested void kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb` · pair void kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`.
- Residue sweep: **0** fixed `…-l4-content-req` / `…-l4-content-verdict` without a generation · **0** `re-answered at B22` · the only surviving `one gated leg = one dispatch id` occurrence is the Status line **quoting** the shorthand it removed · ten-row structure, §4/§5/§6/§8 headings and the write-fence clause all present.
- Interface lock `cbd1893c…` UNMOVED. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — revised `master/STEP-3-LANE4-PLAN.md` to rev11 (content-review id/parent templates + rework-generation clause, §7 step 4, GRILL_SOURCE, design-lock impact line, status/title only); this relay + one INDEX.md row. No deviation, banner, fixture, manifest, lock, owner/frozen byte moved; no `frank/` action, no seat boot, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev11 `9f3a142c…`; on approve → master writes the fresh **INERT** rev11 kickoff, validated against the artifacts it names **and** the seats' actual tool configuration; the operator then boots `l4.planner` + `l4.implementer`. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
