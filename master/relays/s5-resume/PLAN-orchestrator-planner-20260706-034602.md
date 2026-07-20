## PLAN — s5 CHECKPOINT-RESUME: the frank dogfood is stood down (operator-directed); resume as a standard file-relay team to the s5 checkpoint; exit-gate adaptations inline

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s5-resume
PARENT_DISPATCH_ID: s5-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator's stop-the-line direction (2026-07-06, in-session) IS the exercised human gate for the mode change; the s5-close sign-off remains the operator's, separate
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md
FROM: master.orchestrator-planner
TO: s5.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-2.planner, m-1.planner, m-7.planner
SUBJECT: frank transport STOOD DOWN at checkpoint (F11 lineage livelock — stop-the-line; ledger = TRANSPORT-FINDINGS-2026-07-06); s5 resumes file-relay per s1–s4 precedent; same team, same scope minus the live-store §7 leg (adapted); fidelity answers settled alongside; then the transport-fix cycle

**The decision (operator, 2026-07-06).** The s5-through-frank dogfood is **stopped where it stands**. Your F11 finding — a seat that can construct **no acceptable parent** under concurrent traffic, blocking even report-only relays — escalated the lineage/parenting cluster from friction to a **hard liveness failure**, and the operator ruled stop-the-line: the conductor is decommissioned, the live store archived (`~/frank-archives/frank-team-store-s5-dogfood-20260706` — evidence, preserved), and the transport fix becomes its own build **after your checkpoint**. **The dogfood succeeded** — it was run to find where frank breaks under real load, and it did: the run's primary deliverable is the findings ledger **`master/TRANSPORT-FINDINGS-2026-07-06.md`** (F1–F17; your F1–F11 are its spine, F8 noted as your numbering gap). This was the experiment; file-relay is the proven mode s1–s4 shipped on. Reverting is not regression.

**Mode change — effective immediately:**
- **Transport = file relays**, v2.8.8 discipline, operator hand-relay — exactly the s1–s4 pattern. `sprint-doc-setup` in `frank/`; your relays under `frank/.relays/s5/…`; `relay-lint` before every handoff.
- **Same team, same seats:** your addresses (`s5.orchestrator-planner`, `s5-a.planner/.implementer`, `s5-b.planner/.implementer`, `s5.orchestrator-reviewer`) carry over as file-relay addresses. Your sessions keep their context; the frank MCP tools simply stop working — do not retry them.
- **The F7/waiver story ends with frank.** File relays have no typed-CC validator: CC your reviewer normally per protocol; the `ORCH_REVIEW_WAIVER` pattern does NOT carry into file mode.
- **Routing unchanged:** m-x questions still route through master (now as file relays via the operator).

**Immediate actions (you):**
1. **File your held F9/F10/F11 report as a file relay** for the record — the ledger cites it; the trail should hold the primary artifact.
2. **Re-issue your two AUDIT dispatches as file relays** (`frank/.relays/s5/…`) — the F7 blocker does not exist in file mode.
3. **Consume `master/relays/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md`** — every Q1–Q11 settled, no owner rejection: **owed-#3 EMPTY** (both legs; nothing to build) · **routing_escalation HOLD LIFTED** (§J2 landed; m-2's exact registry delta inline — land it in s5-a's single registry pass with the B.3 OI-S4-TOKEN-SCOPE fold) · **no Step-1 slot_in writer** · **`scope_paths` struck** (optional dormant home) · **Q4 artifact set final** (version-label bump + zero-loss replay + negotiation legs; NO envelope migrator) · **⑤ build shape ruled R-2** (real scanner code, real chokepoint, present-but-dormant, fixture-driven) · **no live `record_kind` widening** · **`attestation_source` folded into s5-a scope (O-2)** · Q8/Q9/Q10/Q11 confirmed with the clarifiers + the `authority_class`-is-a-`bool` correction.
4. Proceed: pair PLANs → plan-gates (F2 conditions unchanged) → build to the checkpoint.

**Exit-gate adaptations (master-authored amendment to the s5-dispatch gate; operator-directed basis above; VP CC'd for confirmatory review — flag any objection within one relay):**
- **[VP-W7] "§7 on the live store" → §7 fixture legs.** The live wired store no longer exists to receive them. Your registry additions still prove the §7 class against test stores: operator-authorized record shape, **old→new digest, no re-genesis**, phase-0 digest acceptance via the genesis chain, stale-form re-render + re-rendered success. The §7 mechanism itself is already live-proven (s4 E3 + this run's real config usage). **The live-store application leg moves to the transport-fix relaunch** — applying s5's registry via §7 on the fixed conductor's blessed store becomes that cycle's first live act (a better test than landing it on a livelock-prone transport).
- **"The dogfood record" exit item → CLOSED EARLY.** The usage data exists and exceeds what the gate asked: the findings ledger + the archived store ARE the dogfood record. Cite them; done.
- **Everything else stands unchanged:** registry complete + dormant with the [VP-W3] enumerated negative fixture · the owed fixtures (③ per Q5, ⑤ per R-2, I-PH per Q11, routing_escalation member+fixture; GRILL_REQUIRED is EMPTY) · versioned + replay per Q4 · E2 floors (battery green, zero regression, byte-exact enum, guardrail surface) · honesty [VP-W1] (consumer fields declared-not-observed; transport/provenance-only phrasing travels).

**Worktrees (operator, unchanged ask):** `s5-a-registry` and `s5-b-mechanisms` off `main @ 67ee23e`, suggested at `~/frank-s5-team/s5-a` and `~/frank-s5-team/s5-b`.

**Scope fence unchanged** (the s5-dispatch OUT list stands; [VP-W4] §C4 is the source of truth) — plus one addition: **no transport-fix pre-work by s5.** The lineage/parenting/codec fixes are the next master-run cycle (m-7/m-1/m-2, grilled); your findings ledger is its input, your checkpoint is its precondition. Do not patch the conductor.

**Not authorized by this relay:** no s5-close authority, no scope expansion beyond the named adaptations, no transport-fix work, no conductor relaunch.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s5-resume` — run below.
- Basis: operator stop-the-line direction (2026-07-06, in-session); `master/TRANSPORT-FINDINGS-2026-07-06.md`; the s5-fidelity RECONCILE (settled answers); the s5-dispatch (VP-approved charter this amends); `frank/` `main @ 67ee23e`.

ACTIONS_GIT_REF: none — plan/mode-change relay; no git action, no `frank/` edit. The conductor decommission + store archive are operator-side ops recorded in the ledger.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main @ 67ee23e` (post-`s4-close`), clean.
Next requested action: operator hand-relays this + the fidelity RECONCILE to the s5 session (and the boot-pointer equivalents to the pair/reviewer sessions as s5 directs); s5 files its held report, re-issues its AUDIT dispatches as files, and runs to the checkpoint; VP confirmatory pass on the gate adaptations; master then opens the transport-fix cycle.
