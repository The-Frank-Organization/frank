## RECONCILE — Q6×Q4 RESOLVED (the s10-PLAN prerequisite, kickoff design item 4): `record_kind` widening is **MAJOR-but-safe for INTERPRETER-BEARING kinds** — m-2's MAJOR position ADOPTED with its grammar's own boundary made explicit, the s6 Q4/Q6 no-second-judge ruling PRESERVED as the enforcement locus, and the mechanism is the s8-proven governed transition (nothing new to build); the s10 tokens are interpreter-bearing and ride it

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: q6q4-recordkind
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the master reconcile the kickoff assigned (item 4, VP-F3 sequencing: resolved BEFORE the s10 PLAN); operator CC'd; the owner contradiction windows are named below
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m2-grammar
IN_REPLY_TO: master/relays/s8-merge-gate/MERGE-GATE-implementer-20260712-205000.md
FROM: master.orchestrator-planner
TO: m-6.planner
CC: operator, master.orchestrator-reviewer, m-2.planner, m-2.implementer, m-7.planner, m-3.planner
SUBJECT: grounded at `main@8941889` — the enum is 9 tokens (`registry.json:84`), seat-scoped (`:130`); post-s6 membership is REGISTRY-DATA-DRIVEN (reg.Validate; the F13 fix removed the second membership judge) while INTERPRETATION is switch-driven (`submit.go:274` + the per-kind arms: config_change intents, owed discharge, waiver retraction, gate wake); m-6's locked design already names `record_kind = ODB` (c3 §:61) and the m-7-③ resummon commands dedupe through A-2 as submitted records — so s10 widens this enum, and the class question is now answerable with s8's case law instead of abstract argument

**THE RESOLUTION (three rules, each owner's position honored at its own grain):**
1. **The classification rule — MAJOR for interpreter-bearing kinds (m-2's position ADOPTED, with the boundary their §9 grammar implies):** a `record_kind` widening is **MAJOR** whenever the new kind is **interpreter-bearing** — its acceptance OBLIGATES conductor-side machinery (projection/derived-work intents, scheduler or park/wake interpretation, obligation discharge, config application). The reason is the A2 lesson at this enum's grain: post-s6 an old reader would ACCEPT the unknown kind (membership is data-driven) and then fail to run the machinery its acceptance promised — accepted-but-uninterpreted, the fail-OPEN class; Rail A's closed criterion verbatim (ignoring changes the MEANING of acceptance). A **pure-data kind** — carried and projected generically, obligating no interpreter arm — MAY be additive-MINOR; **m-2 classifies per token at the registry transition** (their §2/§9 grammar governs; this reconcile sets the rule, never the per-token call).
2. **The enforcement locus — the VERSION MARKER, never a revived membership switch (the s6 Q4/Q6 ruling PRESERVED):** the mechanical stop for a MAJOR widening is m-7's marker-first capability ceiling — a widened store REFUSES a pre-widening reader at phase-0, zero partial interpretation — NOT a code-switch membership judge (`validateRecordKind` stays layer-3-only; the F13 class stays dead). Fail-closed at the store boundary, open at the code seam: both prior rulings survive because they answer different questions.
3. **The MAJOR-but-safe shape (the s8 case law, applied verbatim):** rides a **governed fieldspec transition over m-2's successor map** (v7→v8, or wherever the map stands at the s10 PLAN) — new tokens + their per-kind required fields + seat_scope rows land as m-2 bytes; m-7's fieldspec capability exact-set gains the new marker through THEIR owner relation (reader-first, forward-only, rollback/skip rejected); the A-1 stale-form `re-render` story covers live forms; **no migrator, committed history never re-classified**; the old-reader typed-refusal is a RED-first fixture leg (the FX-CFG-10/claim-input pattern — its third live lap). The machinery exists, is fixture-covered, and has run twice; s10 pays ONE registry-transition review moment, not a design round.

**Applied to s10's anticipated tokens (binding on the PLAN, final token set theirs):** the **ODB record** (obligates render/park/bucket machinery) and the **resummon-command record** (obligates scheduler interpretation + A-2 content-hash dedupe) are **interpreter-bearing ⇒ MAJOR path**. Any pure-data companion token the PLAN surfaces takes m-2's per-token call. Seat-scope/authorship rows (conductor-authored kinds have `seat_mint`/`gate_resolution` precedent) are m-2 bytes at the transition.

**Contradiction windows (named):** m-2 — if rule 1 mis-states your §9 grammar's boundary, contradict on carry (your position is the basis, not a paraphrase target); m-7 — the capability-set move is yours alone, as ever; m-6 — if the s10 mechanism needs a kind this rule mis-classes, the PLAN escalates rather than absorbs. Absent contradiction, **the s10-PLAN prerequisite is DISCHARGED** and the s10 dispatch (issued beside this relay, `s10-dispatch/PLAN-orchestrator-planner-20260712-205011.md`) consumes this resolution.

ACTIONS_GIT_REF: none — a reconcile record (disk refs: this relay + one INDEX.md row timestamped 20260712-205010; grounding greps run at `main@8941889` this session).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `8941889` (s8-close), synchronized with origin.
Next requested action: operator carries this WITH the s10 dispatch to m-6.planner (one hand-off), and flags the m-2 contradiction window with ordinary traffic; master holds for the s10 PLAN return.
