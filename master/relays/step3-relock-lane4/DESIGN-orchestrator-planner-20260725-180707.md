## DESIGN — INERT KICKOFF AUTHORED: `master/STEP-3-LANE4-KICKOFF-PAIR.md` = `4f3928ec…`, bound by hash to VP-approved plan rev13 `60daac08…`. New filename so it cannot be confused with the two VOID drafts. Instantiates rev13 §3's rules as concrete `step3-relock-lane4-l4-*` ids (ten rows, none reused, root used only as row 1's PARENT), transcribes §4/§5 from the live plan (exact §7 key set, all typed fields + fixed values, 30-turn/100-call budget, all three carried obligations), and states the write fence with its **honest grade** — convention + review, NOT mechanically enforced, because the seats' sessions have ordinary filesystem write. **GRANTS NO LANE ACTION.** Conformance-checked against the live artifacts before this record; one improvement folded (the ten records are enumerated literally rather than as `xit-dur-1..5` — a brief the pair counts from must not carry a range).

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the kickoff is INERT and grants nothing; the operator alone decides whether to boot `l4.planner` + `l4.implementer` and hand it over. Master does not boot seats.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-175715.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: INERT kickoff `4f3928ecf7076cef489281f49866a22790e8e907f2d8b0c3c1e7b062ce39f1ef` authored to approved plan rev13 `60daac08…`; ready for the operator's boot/handover decision

## What this delivers

Per the VP's final approval (`…/DESIGN-REVIEW-orchestrator-reviewer-20260725-175715.md`: *"master may now author one fresh inert kickoff brief"*), master has authored **`master/STEP-3-LANE4-KICKOFF-PAIR.md`** = SHA-256 `4f3928ecf7076cef489281f49866a22790e8e907f2d8b0c3c1e7b062ce39f1ef`.

It meets the three approval conditions exactly:
1. **cites plan rev13 at full SHA-256** `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca`, alongside the interface lock `cbd1893c…` and deviations `e117d059…`, with *"the plan governs on any difference"* stated;
2. **instantiates the approved unique dispatch-id and exact-parent rules** — ten concrete `step3-relock-lane4-l4-*` rows, no id reused, generation-suffixed content ids, the rework cycle and escalation `req → disp → resume`, and the root appearing only as row 1's `PARENT_DISPATCH_ID`;
3. **validated against every live artifact it names and the seats' tool configuration** (below).

**A third filename** (`…-KICKOFF-PAIR.md`) so it cannot be mistaken for the two VOID drafts, both of which remain byte-unchanged.

## Validation record — against live artifacts, not my own prose

The rule that earned this discipline is now applied to the brief itself:
- the three cited SHAs **recompute** from disk to the values printed in the brief;
- the ten record ids, the exact §7 key set, every typed-field name and fixed value (`{counter_before_recovery:1, counter_after_recovery:1, invocations_after_recovery:0}`, `handoff_expected_records[2]`, the four `resume_prefix_expectation` members, `expected_disposition:"degraded"`), the 30-turn/100-call budget, the top-level baseline digests, all three carried obligations (N910 · `env_digest` parity · the r7-mirror **STOP-and-escalate**) and the guiding-PM/fidelity boundary were **transcribed from rev13 §4/§5 as read from disk**;
- the dispatch-id table checks **unique**, with no root reuse.

**One improvement folded after the check:** rev13 §4 writes the records as `xit-gov-1; xit-dur-1..5; …`. A *plan* may use a range; a *brief the pair counts ten records from* must not, so the kickoff enumerates all ten literally and says "ten, no more and no fewer". This instantiates rev13 rather than departing from it.

## The honest limit I put in writing

**The §2 write fence is convention + review grade, not mechanical.** The pair's sessions have ordinary filesystem write access, so nothing *prevents* a write outside `master/relays/`; the implementer, master and the VP check that none happened. The brief says so plainly — *"do not treat 'I could' as 'I may'"* — rather than implying an enforcement that does not exist. If you want partial mechanical enforcement, a host-level deny rule scoped outside `master/relays/` is available; the brief does not assume it.

## Your decision — nothing is running

The kickoff is **inert**: no seat is booted, nothing is authored, and reading it authorizes nothing. When you choose to proceed:
1. boot `l4.planner` and `l4.implementer` as independent sessions (the implementer separately, never tokened by the planner);
2. hand this brief over as the `…-l4-dispatch` relay;
3. the pair authors → master materializes → implementer proves equality → implementer content-reviews → out-of-team owner-fidelity (guiding PM **m-3**) → VP → **Master+VP freeze + re-lock**.

Hand-relayed file relays throughout (B22); no frank, no MCP, no preflight.

## Boundaries
No seat booted, no handover, no activation, no proposal, no materialization, no equality or content review, no owner-fidelity request, no fixture, no manifest, no freeze, no re-lock, no PLAN-lock, no T4/code token, no credential, no provider call, no `frank/` action by this record. Item A lock `cbd1893c…`, plan rev13 `60daac08…`, both void banners and all owner/frozen/amendment bytes UNMOVED. **H-12 hard-blocks external use.**

## Verification
- Kickoff `master/STEP-3-LANE4-KICKOFF-PAIR.md` = SHA-256 `4f3928ecf7076cef489281f49866a22790e8e907f2d8b0c3c1e7b062ce39f1ef`.
- Approved plan rev13 `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca` · interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` · deviations `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` — all recomputed, all unmoved.
- Void banners unmoved: nested `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb` · pair-shaped `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`.
- Conformance pass: all ten record ids literal · seven exact schema/typed-value fragments present in **both** plan and kickoff · ten unique dispatch ids · no unsuffixed content id · three carried obligations · INERT/fence/H-12/two-duties/±2-min clauses all present.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — authored `master/STEP-3-LANE4-KICKOFF-PAIR.md` (INERT) + this relay + one INDEX.md row. No plan/deviation/banner/lock/owner/frozen byte moved, no `frank/` action, no seat boot, no handover, no activation, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the operator decides whether to boot `l4.planner` + `l4.implementer` and hand over the kickoff. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
