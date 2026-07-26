## DESIGN — lane-4 plan rev8 `1fc035fa…`: B22 REVERT. Operator stood down frank-as-courier and selected the PAIR shape, so rev8 restores the team shape + transport of your APPROVED rev5 `faa23c7f…` and withdraws B21's nested-team-on-frank (rev6/rev7). A revert-plus-fold, not new design: §4 (exact §7 schema + fixed values), §5 (carried obligations + owner-fidelity matrix), §6 (Master+VP reservations) and §8 are preserved **by line-slice, byte-stable**. Withdrawn with the courier: the preflight, the nested ceiling/lineage apparatus, the frame-fit/chunk contract. RETAINED from r7: role-canonical addressing + one-gated-leg-one-dispatch-id. The frank dogfood closed **successful-and-complete** — four findings banked, nine records exported — and it also caught two MASTER drafting errors, now folded as standing rules.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — B22 (stand down frank-as-courier; pair not nested team) is an operator directive, and the operator alone boots the pair. No preflight or activation gate remains.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-110000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev8 `1fc035fa69d3027c181eb3408c285bb196d33698c024c99a75515f28ddd9042e` — B22 revert to the approved rev5 pair shape on hand-relayed file relays; Item A lock `cbd1893c…` preserved

## What changed vs rev7 `e7a333e9…` — and why

**B22 (operator, 2026-07-25): frank-as-courier is STOOD DOWN for lane 4, and the pair shape is selected.** Recorded in `master/PROTOCOL-DEVIATIONS.md`. The operator's reasoning, which I put to them and did not soften:
- the protocol's worth is already established by ~2,300 relays here plus ~1,250 in pdc across eleven closed slices and two step-exits, so lane 4 on MCP re-proved it at a far worse exchange rate;
- the preflight cost roughly a dozen operator prompts for five relays, and **that ratio cannot be improved at the MCP layer** — MCP gives a server no way to make a client originate a turn, so an MCP-hosted seat is operator-scheduled *by construction*;
- the deliverable is a **single tightly-coupled manifest**, which favours a pair on throughput — the operator's own decision rule from the B21 grill (*"if throughput we're not doing it thru frank"*).

So rev8 **restores what you already approved at r5** (pair + hand-relay) rather than proposing anything new.

**Withdrawn with the courier:** the zero-authority preflight; the nested ceiling/lineage apparatus and the `l4.orchestrator-*` / `l4.w<k>.*` seat set; the frank `max_frame_bytes` frame-fit rule and the chunk/archive contract it forced (a file relay carries the envelope whole).

**Preserved byte-stable (by line-slice, not retyped): §4, §5, §6, §8.** **Retained from r7 because they are transport-independent:** role-canonical `ROLE: Planner`/`ROLE: Implementer` addressing (r7-F1) and one-gated-leg-one-dispatch-id (r7-F2, the shared-id resolver defect).

## The dogfood closed successful-and-complete — and corrected me twice

Four findings no manual relay could surface, all in `FRANK-HARDENING-BACKLOG.md`, with nine committed records exported to `master/relays/step3-relock-lane4/preflight-export/`:
1. the `form_digest`/`CEREMONY_TIER` trap — the advertised digest is valid only for a tier the caller did not choose, and recovery depends on a `tools/list_changed` capability Codex lacks: a real cross-harness interop defect against the stated MVP goal of retaining MCP for foreign harnesses;
2. empirical confirmation that `provableParentHint` enforces confusion-resistance by **refusing lineage a seat cannot prove**;
3. a CC'd relay capturing a reply's lineage parent via `woken_on`;
4. no roster introspection — a seat cannot distinguish "peer not booted" from "peer idle", nor verify its own credential is bound.

**Two of the three seat-reported blockers were mine, not frank's** — `CEREMONY_TIER: tiny` in the boot charters, and a §7 map asking workers to cite a thread they never see. Both are folded into rev8's GRILL_LOCK as standing rules: **an authority-of-record naming wire fields must be validated against the live tool schema, not against its own prose**, and **a lineage map may only ask a seat to cite a parent it can prove it participated in**. The nested kickoff is banner-VOIDed with both defects named, so the history reads honestly.

## What I ask the VP to review
- Is the revert clean — does rev8 restore r5's shape without silently carrying nested-team residue?
- Are §4/§5/§6/§8 genuinely byte-stable (no substance moved under cover of a revert)?
- Is the withdrawal of the preflight, the nested ceiling/lineage apparatus, and the frame-fit contract correct now that the courier is gone?
- Anything reopening a closed r1–r7 decision or a locked byte?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or team stand-up on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. The l4 frank store is torn down (conductor stopped, MCP wiring and Codex seat profiles removed); the store directory remains on disk pending operator disposal, its records already exported. On VP approval → master writes the **INERT** rev8 kickoff; the operator boots the pair. **H-12 hard-blocks external use.**

## Verification
Plan rev8 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `1fc035fa69d3027c181eb3408c285bb196d33698c024c99a75515f28ddd9042e` (supersedes rev7 `e7a333e9…`). §4/§5/§6/§8 preserved by line-slice from the rev7 bytes. No live `l4.orchestrator-*` / `l4.w<k>.*` seat reference remains (grep 0; the 13 residual matches are historical/superseded citations in Status and the GRILL_LOCK rejected-alternatives, which is intended). `master/STEP-3-LANE4-KICKOFF-NESTED.md` banner-VOIDed. B22 recorded in `master/PROTOCOL-DEVIATIONS.md` as a build-phase bullet ahead of B21. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-LANE4-PLAN.md` to rev8 (§3/§7/§9/§10 + header/boundaries; §4/§5/§6/§8 preserved) + banner-VOIDed `master/STEP-3-LANE4-KICKOFF-NESTED.md` + appended B22 to `master/PROTOCOL-DEVIATIONS.md` + two battle reports in `master/FRANK-HARDENING-BACKLOG.md` + exported nine preflight records + this relay + one INDEX.md row; tore down the l4 frank courier (conductor stopped, `.mcp.json` and four Codex profiles removed). No fixtures/manifest/lock/owner/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev8 `1fc035fa…`; on approve → master writes the INERT rev8 kickoff (validated against the artifacts it names) and the operator boots `l4.planner` + `l4.implementer`. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
