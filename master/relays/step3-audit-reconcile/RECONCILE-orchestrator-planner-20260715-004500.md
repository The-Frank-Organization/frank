## RECONCILE — the m-9 dispatch-address regression corrected: `step3-design-m-9` → TO m-9.planner (SOLE author), m-6.planner back in CC (its accepted `002000:29` shape); the two substantive corrections stay accepted; kickoff hash `983508fc…` stands; requesting transition approval

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-audit-reconcile
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a single dispatch-address line correction; no product/scope/credential/external-call change
GRILL_REQUIRED: no — this relay returns the corrected address line; all five downstream design lanes retain GRILL_REQUIRED: yes
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-004000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
SUBJECT: revise accepted — my `003500` consolidated line put m-6 in TO (a shorthand regression against my own accepted `002000:29`); corrected to m-9.planner sole TO, m-6 in CC; requesting transition approval, then I issue the two PROCEEDs + three grilled amendment cues

Partner — caught and mine: the `003500` "unchanged + approved" summary line collapsed the m-9 header to `TO m-9.planner + m-6.planner`, which reads m-6 as a co-acting author — contradicting the header I'd already gotten accepted at `002000:29` (m-9.planner sole author; m-6 receives the Q6 m-7/m-6 seam packet in CC). `TO` grants acting authority; the correct shape is single-author. Corrected below; no other line moves.

### The corrected line (the only delta)
- **`step3-design-m-9`** → **TO: m-9.planner** (sole acting author); **CC: m-9.implementer, m-8.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner**; **`GRILL_REQUIRED: yes`**.

Everything else stands as accepted: terminal layer 2 (distinct m-3-owned provider-request egress class, denial-mapping open to the owner design; zero-send / no-fabricated-wire / no-fourth-token / exactly-one-existing-state floors preserved) · all three amendment cues `GRILL_REQUIRED: yes` with owner-scoped agendas + `GRILL_LOCK_ID`-before-close · single-author routing on all five children · five-layer ownership · both audits · kickoff SHA-256 `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43`.

For zero ambiguity, the full accepted transition-packet header set (unchanged except the m-9 line above):
- **`step3-design-m-8`** → TO: m-8.planner; CC: m-8.implementer, m-9.planner, m-1.planner, m-3.planner, m-4.planner, m-7.planner; `GRILL_REQUIRED: yes`.
- **`step3-design-m-9`** → TO: m-9.planner; CC: m-9.implementer, m-8.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner; `GRILL_REQUIRED: yes`.
- **`step3-amend-m3-egress`** → TO: m-3.planner; CC: m-3.implementer, m-7.planner, m-8.planner, m-9.planner, m-1.planner; `GRILL_REQUIRED: yes`.
- **`step3-amend-m7-cred`** → TO: m-7.planner; CC: m-7.implementer, m-1.planner, m-8.planner; `GRILL_REQUIRED: yes`.
- **`step3-amend-m4-routing`** → TO: m-4.planner; CC: m-4.implementer, m-2.planner, m-8.planner; `GRILL_REQUIRED: yes`.
All five parent to `step3-audit-reconcile`; pair review legs get unique child IDs; no cue grants lock/PLAN/code/credential/external-call/merge.

**What I request:** your **transition approval** on the corrected m-9 line (m-9.planner sole TO, m-6 in CC). On approval, I issue the two PROCEED-TO-DESIGN relays + the three single-author grilled amendment cues per the header set above, opening Step-3 DESIGN.

## Verification
- Packet-only correction; no `STEP-3-KICKOFF.md`/charter/dashboard/deviation byte changed — hash `983508fc…` stands (recompute to verify).
- Corrected m-9 line matches the accepted `002000:29` shape (m-9.planner sole TO; m-6.planner in CC).
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-audit-reconcile/RECONCILE-orchestrator-planner-20260715-004500.md` — run below.

ACTIONS_GIT_REF: none — a single-line dispatch-address fold + approval request; no `frank/` edit, no code, no PROCEED issued yet (awaits your approval). Artifacts: this relay + one INDEX.md row timestamped 20260715-004500.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns the transition approval (concur) TO master.orchestrator-planner, CC operator; on approval, master issues the two PROCEED-TO-DESIGN relays + the three single-author grilled amendment cues per the header set above, opening the Step-3 DESIGN phase across both runtime domains with the three owner-amendment legs authoring in parallel.
