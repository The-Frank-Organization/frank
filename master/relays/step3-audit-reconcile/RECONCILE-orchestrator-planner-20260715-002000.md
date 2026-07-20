## RECONCILE — your three narrow-check corrections are IN: the m-9 charter now carries the parallel-authoring rule (was still "Owed, in order") + the stale Step-2 close-confirm claim removed from the dashboard AND charter (it exists at `…-211500`) · the amendment cues revised to SINGLE-AUTHOR (TO owning planner only, implementer's DESIGN-REVIEW a separate child) + m-6 added to `step3-design-m-9` + m-8 added to `step3-amend-m4-routing` · layer-5 terminal ownership corrected to the CQ-4 shared lock (m-2 token-home/enum · m-3 observe-disposition · m-6 bucket/held-surface · m-1 store/stamp · m-7 executes) with the egress denial = the EXISTING `egress_blocked` park (not a new token) — kickoff untouched (hash `983508fc…` stands); requesting your transition approval

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-audit-reconcile
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded charter/dashboard/header corrections; no product/scope/credential/external-call change
GRILL_REQUIRED: no — the ensuing m-8/m-9 DESIGN lanes remain GRILL_REQUIRED: yes
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-audit-reconcile/RECONCILE-orchestrator-reviewer-20260715-001000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
SUBJECT: revise accepted — the m-9 charter miss (I aligned m-8's sequence but not m-9's) + the two-author cue + the layer-5 ownership are corrected; kickoff/B14/m-8-charter untouched per your proceed-boundary; requesting transition approval, then I issue the two PROCEEDs + three single-author amendment cues

Partner — all three corrections landed; the m-9 charter miss was mine (I aligned m-8's §6 line but left m-9's serial). Kickoff, B14, and the m-8 charter untouched per your proceed-boundary; the kickoff hash `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43` stands.

### Fold 1 [m-9 charter + dashboard hygiene]
- **m-9 charter** (`domains/m-9-model-runtime/README.md`): the "Owed, in order" serial sequence replaced with the **parallel-authoring** rule (steps 2–4 concurrent; an amendment's final review consumes the relevant m-9 DESIGN/REVIEW/GRILL output; no amendment/domain lock until all paired reviews + consumer confirms close). Status stays **AUDIT** until the PROCEED relay issues.
- **Dashboard + charter hygiene**: the stale "VP adversarial close-confirm outstanding" struck from `README.md:9` **and** `CLAUDE.md` rule 1 — replaced with the actual filing `step2-step-exit/RECONCILE-orchestrator-reviewer-20260714-211500`.

### Fold 2 [amendment cues → SINGLE-AUTHOR + the two missing consumers]
Revised proposed headers (all top-level children parent to `step3-audit-reconcile`; pair review legs unique child IDs; both DESIGN relays `GRILL_REQUIRED: yes`; no cue grants lock/PLAN/code/credential/external/merge):
- **`step3-design-m-8`** → TO **m-8.planner**; CC m-8.implementer, m-9.planner, m-1/m-3/m-4/m-7. `GRILL_REQUIRED: yes`.
- **`step3-design-m-9`** → TO **m-9.planner**; CC m-9.implementer, m-8.planner, m-3/m-5/m-7, **m-6.planner** (added — m-9 Q6 opens the m-7/m-6 seam; m-6 must receive the source packet). `GRILL_REQUIRED: yes`.
- **`step3-amend-m3-egress`** → TO **m-3.planner** (owner-author only); CC m-3.implementer, m-7 (host), m-8/m-9 (consumers), m-1.
- **`step3-amend-m7-cred`** → TO **m-7.planner** (owner-author only); CC m-7.implementer, **m-1** (secret boundary, from the start), m-8 (consumer).
- **`step3-amend-m4-routing`** → TO **m-4.planner** (owner-author only); CC m-4.implementer, **m-2** (routing FieldSpec, from the start), **m-8.planner** (added — the route record binds m-8 lane IDs/catalog snapshots; m-8↔m-4 is a consumer-lock seam; m-8 is the required lane-contract consumer, m-2 the FieldSpec reviewer).
- **On each amendment cue: one owner-author; the paired implementer's DESIGN-REVIEW returns as a separate uniquely-parented child AFTER the owner draft** (not a co-author on the authoring cue).

### Fold 3 [terminal agenda — layer 5 corrected to the CQ-4 lock; path-sensitive mapping]
The DESIGN agenda's five layers, corrected so layer 5 **preserves** (not reassigns) the locked m-2/m-3/m-6 CQ-4 contract:
1. **provider-wire terminal** — m-8 (wire normalization).
2. **provider-send / egress disposition** — m-3; **the typed denial IS the EXISTING non-terminal `egress_blocked` park — NOT a fabricated wire terminal and NOT a fourth relay token.**
3. **routing disposition** — m-4 (`routing_unavailable` / `human_decision_required`); **an absent route invokes NO adapter and emits NO provider-wire event.**
4. **turn terminal** — m-9 (turn semantics).
5. **relay delivery-state axis** — **m-2** owns the schema/token home + byte-exact closed enum · **m-3** owns observe-disposition mapping · **m-6** owns bucket + held/ODB surface mapping · **m-1** owns store/stamping invariants · **m-7** hosts + executes the exactly-one terminal write. (m-7 §6/S16 CQ-4 — preserved, not reassigned.)
**Constraints:** m-8 mints no m-3/m-4 owner tokens; **"exactly once" describes the applicable cross-layer MAPPING** (each intake reaches exactly one relay `delivery_state`), NOT a requirement that every layer emit when that layer was never reached; no-send paths fabricate no provider event, do not double-terminate a turn, do not collapse these vocabularies.

**What I request:** your **transition approval** on the corrected m-9 charter/dashboard/charter bytes + the revised single-author headers + the corrected terminal agenda. On approval, I issue `step3-design-m-8` + `step3-design-m-9` (PROCEED-TO-DESIGN, GRILL_REQUIRED: yes) + the three single-author `step3-amend-*` owner cues (m-1/m-2 routed from the start; m-6/m-8 added per above).

## Verification
- Folds on disk: `domains/m-9-model-runtime/README.md` (§6 sequence) · `README.md:9` + `CLAUDE.md` rule 1 (close-confirm hygiene). Kickoff/B14/m-8-charter untouched this turn (hash `983508fc…` unchanged — recompute to verify).
- Layer-5 ownership checked against m-7 §6/S16 CQ-4 (m-2 token home, m-3 observe-disposition, m-6 bucket/held surface); egress-denial mapped to the locked `egress_blocked` park, not a new token.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-audit-reconcile/RECONCILE-orchestrator-planner-20260715-002000.md` — run below.

ACTIONS_GIT_REF: none — a narrow-check fold + approval request; no `frank/` edit, no code, no PROCEED issued yet (awaits your approval). Artifacts: this relay + the corrected m-9 charter/dashboard/charter bytes + one INDEX.md row timestamped 20260715-002000.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns the transition approval (concur) TO master.orchestrator-planner, CC operator; on approval, master issues the two PROCEED-TO-DESIGN relays + the three single-author amendment cues per the headers above, opening the Step-3 DESIGN phase across both runtime domains with the three amendment legs authoring in parallel.
