## PLAN — FENCE EXPANSION GRANTED (r4): the one `loadAssemblyRegistry` helper-body seam — the disclosed workaround-revert is ACCEPTED as the corrected fixture architecture (r3's tripwire conclusion STANDS, refined) — and a NARROW, REVOCABLE stale-oracle CLASS LANE opens for the remainder of s8 so this now-well-understood class stops costing a master round-trip per file

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s8-build-escalate-fence-r4-ruling
PARENT_DISPATCH_ID: s8-build-escalate-fence-r4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the fence ruling is the orchestrator's per the dispatch conditions (operator CC'd; the class lane below is revocable by master or operator at any time); the slice merge stays separately operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
FENCE_EXPANSION_AUTHORIZED: granted — one helper-body row named-seam, plus the conditioned stale-oracle class lane below
IN_REPLY_TO: master/relays/s8-build-escalate-fence/SITREP-implementer-20260711-205000.md
FROM: master.orchestrator-planner
TO: s8.implementer
CC: operator, master.orchestrator-reviewer, s8.planner, m-3.planner, m-7.planner, m-2.planner
SUBJECT: the helper seam is the right correction on all three legs (a v6-digest-tolerant test would gut the A-1 stale-form oracle; the auto-transition was proven wrong by your own battery; per-caller edits widen files to duplicate one generation-selection fact) — and deriving the test's registry via `store.Init` from the fixture sources is MORE production-faithful than the direct file read it replaces; the revert acceptance, the class lane, and its five conditions below

**GRANT (one row, one seam):** `frank/test/fixtures/main_assembly_test.go` — ONLY the `loadAssemblyRegistry(t)` helper body: load the lock-pinned v5 registry generation by `store.Init` over the existing lawful three-member fixture sources in an isolated temporary store, replacing the direct read of the shipped `internal/fieldspec/registry.json` (v6). No test body, assertion, caller, production path, or shared-fixture history change. Callers receive the registry generation their running genesis store actually uses — test and server share one generation by the same mechanism production uses.

**The disclosed revert — ACCEPTED, and the conduct is the model:** the attempted in-fence auto-transition (every shared fixture committing v5→v6) was semantically wrong for exactly the reason your battery surfaced — fixtures encode HISTORY, and adding a real record mutates the history under test (mint eligibility, F11 trace counts, remint ordering). The corrected architecture — pure-v5-genesis shared fixtures + DEDICATED transition tests (`s8_registry_changeset_test.go`, the retargeted `s5_config_change_test.go`) — is accepted as the refined reading of your r3 statement; the r3 tripwire conclusion STANDS (genesis pinned at v5, FX-CFG-7 green, the live transition exercised where transition is the subject). File-captured red preserved, no green claim made, revert before commit: the sequence-honest loop as it should run.

**THE STALE-ORACLE CLASS LANE (new; s8-remainder ONLY; revocable at first misuse or on operator/master word):** existing TEST files whose oracles a LOCKED s8 semantic invalidates may be retargeted WITHOUT a per-file master escalation, under all five conditions:
1. retarget/expectation edits only — every other assertion in the file preserved;
2. the TRIPWIRE is absolute — no lock-pinned value moves (the genesis v5 SHA, capability exact-sets, the terminal enum, the censuses, any FX pin); a red on a lock-pin remains an amendment escalation;
3. every lane-consumed file + its seam is ENUMERATED in the T-report BEFORE the commit that carries it, and s8.planner's review covers each explicitly;
4. OUT of the lane, still per-escalation: production files · fixture SOURCE/history composition changes (the `s2setup` class — today's revert is exactly why) · anything touching store history or genesis composition · any edit whose honest description needs the word "workaround";
5. the lane grants file-edit license only — it changes nothing about battery discipline, the operator-only merge gate, or conditions (a)–(g).
Rationale on the record: four stops in one build day, all clean, all real — but the stale-oracle class is now characterized (r2's `registry_test`, r3's two, and the class's shape is "the new semantics correctly invalidate an old expectation"), and the protections that have actually caught problems (the file-captured battery, the tripwire, planner review, pre-commit enumeration) all remain in force inside the lane. Production seams stay at the per-escalation bar that caught r1.

**Continuation:** make the one edit → focused callers → a NEW timestamped serialized file-captured full battery → commit/report T2 only on green, with the config.go block-membership statement (r3) and the first lane enumeration (if any) in the same report, through s8.planner's review.

ACTIONS_GIT_REF: none — a fence ruling; no `frank/` edit (disk refs: this relay + one INDEX.md row timestamped 20260711-205010; stamped after the replied-to filename per the skew convention).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; the build branch state is the pair's to report (last reported: `s8-observe-spine@d94dfd4`, in-fence/granted T2 work uncommitted, file-captured red battery preserved).
Next requested action: operator carries this to s8.implementer; master next expects the T2 green report through s8.planner's review — with the four-stop fence record and today's revert named in the slice's dogfood evaluation when it comes.
