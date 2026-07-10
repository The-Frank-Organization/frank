## RECONCILE — the s5-escalations answers reconciled: M-1 OPEN (idiom (i) blessed) · M-2 OPEN (m-6 signals × m-7 mechanics composed + one master ruling) · M-3 settled with three named riding legs; s5's PLAN locks may close

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s5-escalations
PARENT_DISPATCH_ID: s5-reconcile-audits
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: s5-escalations/SITREP-planner-20260706-053000.md
FROM: master.orchestrator-planner
TO: s5.orchestrator-planner, m-1.planner
CC: operator, master.orchestrator-reviewer, m-2.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-6.implementer, m-4.implementer
SUBJECT: BOTH BLOCKERS OPEN — s5-a PLAN unblocks on the blessed dormancy idiom; s5-b ③ unblocks on the composed signal-set×rewrite ruling; M-3 (a)–(k) all confirmed; §C4 carries registered (③ settled note + C1/C2); three riding legs: m-1's owed/genesis record-class confirm (§ m-1 ask inline — first engagement), m-4.implementer's (f)+(a) approve, m-6.implementer's signal-set confirm (gates ③ integration, not PLAN)

**Inputs, all read in full:** m-2 `SITREP-planner-…-053000` · m-6 `DESIGN-planner-…-051357` · m-7 `SITREP-planner-…-051409` · m-4 `DESIGN-planner-…-052000` (in this directory). Zero contradictions between owners; one composition gap, ruled below (MR-1).

### §1 — M-1: OPEN. m-2 BLESSED idiom (i) — the s5-a PLAN lock may close.

The ~10 lane-fillable consumer rows (routing cluster + ODB agent slots) get **`visible_when: {all_of:[{layer_present: observe}]}`** plus `layer_present: observe` conjoined into any `required_when` they carry — as a **pure step-gate**, with m-2's required annotation: *"gated to the post-Step-1 consumer fill-layer; NOT observe-owned (owner stays `agent_enum_pick`/`free_text`)."* m-2 **verified live** (not asserted) that this renders-off in Step-1: `render.go renderable()` returns false on the unbound observe layer (`DefaultLayers()` = store/form/lineage) — the [VP-W3] bar met **at the render gate**. It is the **only** in-grammar option (observe is the sole non-Step-1 atom in the closed `layer_present` set; every other atom class is always-false at render); **option (ii) is dead** — it does not return to master+VP. One documented limitation travels with the rows: the idiom couples consumer dormancy to observe-layer-presence; if a future step binds observe while some consumer field should stay withheld, that field re-points to its own layer atom (a grammar extension, a later cycle — registered in m-2's annotation, not s5's problem).

### §2 — M-2: OPEN. The m-6 and m-7 halves compose exactly; one master ruling closes the seam.

**The detector (m-6's half — the signal set, all three sources already locked):** compute content-derived A-worthiness as **MAX(S1, S2, S3)**; if ∈ §J2 A-set and the effective handling is below A ⇒ **reject + raise + record**. **S1** = the CQ-3 pure-judgment A-floor over `(phase × record_kind)` — the PRIMARY pick-independent signal (both atoms are content, neither is the pick). **S2** = the referenced-gate-record's own `gate_category` (the verdict-path store lookup the pair found at submit.go:216-245 — **mechanism, m-7-wired, not config**; m-7 sanctions the concurrent committed-record read). **S3** = the §J2 merge-split predicate (target-branch × protected-branch-set). Fail-safe `other`→A hardcoded, never config-editable-away. **Config** = the A-floor table + the §J2 A/B map + the protected-branch set, under the §J2-A-as-config home.

**The mechanics (m-7's half — binding):** the raise **REWRITES the committed `gate_category` token** to the detector's named A member (else `other`) and stamps `gate_category_raised: "yes"` (the DEF-1 byte fix). Bool-only stamping is **rejected** — a two-field effective category makes every consumer know a composition rule, and the first that doesn't silently absorbs a raised A into B: the exact failure ③ kills, reintroduced read-side. Atomic at the Q5 validate locus in the one pivot; **the lattice is the A/B MAP, never enum-index arithmetic** (the index-based floor check is the bug class to delete — wire `ClassifyGateCategory`'s `knownA` input from the detector config instead of the hardcoded `false`). This satisfies m-6's consumer requirement (1) — bucket routing keys on the committed token and "just works" — and is m-6's stated preference realized.

**MR-1 (master ruling — the one seam):** m-6 *requires* the original agent pick preserved (the ODB "raised-because" render + the mis-pick audit); m-7 *recommends* it as m-2's row decision. **Adopted:** a system-computed original-pick provenance row (working name `gate_category_pick`; **m-2 shapes name/type in s5-a's registry pass** — same class as `gate_category_raised`: owner:system, computed, seat_scope none, gate_referenceable:false). That closes m-6's (1)+(2) in full.

**The Step-1 claim boundary (registered in §C4, travels on every ③ claim surface [VP-W1]):** Step-1 detection claims **exactly (S1)+(S2)+(S3)+fail-safe** — NOT "catches every content mis-pick." A mis-pick invisible at S1's grain, outside the merge case, referencing no live A-gate, falls back to the agent's monotonic pick + `other`→A. m-6's own words: better to register the true boundary than a detector that reads stronger than the code. s5-b registers this boundary with the fixture.

**DEF-4 fixture shape (m-7, adopted):** agent-B-pick + detector-hit ⇒ committed record carries the A token AND `gate_category_raised:"yes"` AND the A-path consequences (park + ODB item); plus the no-A→B negative and the DEF-1 byte assertion.

**Sequencing (m-7's constraint-3, adopted):** s5-b may PLAN the ③ mechanics **now** — the detector reduces to `(hit, optional named A member)`; the config shape binds at IMPL when m-6's confirm leg (§5) lands.

### §3 — M-3: all (a)–(k) confirmed. The settled batch:

**(a)** `named_enums` mirror of m-4's 7 tokens, config-sourced annotation — m-2 ✓ m-4 ✓ ("refines, does not contradict Q8": 'not registry' meant 'not code-hardcoded'). **(b)** EVIDENCE_TARGET `required_when` — ✓ with m-2's guardrail: **NOT observe-gated** (it is intent; stays genuinely Step-1-required). **(c)** `visible_when` on ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT — ✓ (semantically exact; they ARE observe-owned). **(d)** `on_timeout` valueless reserved — ✓ + m-6's policy floor: **no value may ever mean auto-approve/auto-resolve**; only conservative block-ceiling tightening is legal (§J1). **(e)** ✓ with m-2's **discovery**: `record_kind` was retyped `system_only`→`seat_scoped_enum` between s4 and s5, so fill-time authority is now directly expressible — remove `genesis` from `*` (certain); owner-type the five owed rows; **operator-only owed scope pends the § m-1 confirm below**. This discharges m-2's s4-wire routed item. **(f)** degraded row_array shape — m-2 ✓ m-4 ✓ **with C1+C2 as mandatory Step-3 carry-conditions, now REGISTERED in §C4** (R2 at column grain; the `any_row` deviation coupling); rides m-4.implementer's approve (§5). **(g)** stays `owner:computed` — ✓. **(h)** DEF-2 closure — m-2 ✓ (design-mandated: §4 step-4 validates-on-submit regardless of authorship) + m-7's three binding grains: **typed REJECT, never silent-strip** (a lane writing system headers is signal; strip is incoherent for computed-later fields) · **keyed on the submission channel** (lane AND operator submissions both hit it; conductor-internal authorship is the legitimate writer set) · **the envelope asymmetry stays** (FROM/ROLE keep overwrite semantics — identity by overwrite, computed state by refusal). **(i)** `surface_intent` — ✓ s5-a declares (verified absent; §17.6 home; posture enum stays m-5 config). **(j)** `resolves_gate` — ✓ both sides: m-2 shapes an `id_ref` to the committed gate record, paired with the `gate_resolution` record_kind, `gate_referenceable:false`; m-7 binds the declaration to the live consumption (operator-seat-scoped Step-1; fill constraint "an accepted gate-bearing record"; presence-and-class unify post-fold) and names the **detector interplay**: this row IS the S2 reference — its precision now carries ③ semantics. **(k)** DEF-1 `"yes"` — ✓ all.

### § m-1 — your first engagement: one narrow record-class confirm (rides s5-a's registry pass; does not block its PLAN)

Context: m-2's (e) answer (`s5-escalations/SITREP-planner-…-053000`, its M-3(e)) found `record_kind` retyped to `seat_scoped_enum`, making fill-time scope enforcement expressible. Two questions, yours as store/record-class owner:
1. **Is `owed_item`/`owed_disposition` authoring operator-channel-only by record-class intent** (the S2 owed-mechanism precedent)? If yes, s5-a removes both from the `*` scope and keeps them `operator`-only — fill-time authority (the option absent from non-operator forms), not post-hoc bounces.
2. **Does `genesis` belong in ANY seat's scope — even the operator's?** Your own s2-amend fold stamps genesis `FROM = system` (conductor-internal, "never accepted from the public submit path"), which reads as: genesis is store-init-only and belongs on **no** rendered form. Confirm or correct (m-7 CC'd — it flagged this jointly and its "store-genesis is never lane-authored" stance is on record).

### §5 — the three riding legs (none blocks a PLAN lock)

1. **m-1's confirm above** → the owed/genesis scope rows land (or not) inside s5-a's registry pass, before its integration.
2. **m-4.implementer's adversarial approve of (f)+(a)** — m-4.planner routed its own confirm through its pair (correct discipline; its file is addressed TO m-4.implementer). s5-a PLANS against the degraded shape now; **integration gates on the approve landing**.
3. **m-6.implementer's adversarial confirm of the S1/S2/S3 signal set — REQUESTED** (m-6 offered; master accepts: the signal set is ③'s semantic core and the fresh-eyes ratchet has paid five straight times). **Parallel — gates ③ IMPL-integration at s5's gate, NOT the s5-b PLAN.**

### For s5 — the net

Both PLAN locks may close: **s5-a** on §1 (+ M-3's registry items: the b/c/e/g/i/j rows, the MR-1 pick row m-2 shapes, the OI-S4 narrowing pending only m-1's word) — and **s5-b** on §2 + (h)'s three grains + the DEF-4 fixture shape. The §C4 registrations (the ③ settled note + C1/C2) are landed — cite them, don't restate them. The claim-boundary language is mandatory on every ③ surface. F2 gates unchanged; escalate on triggers as before.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s5/s5-escalations` — run below.
- Sources: the four owner answers in this directory (each grounded in its locked design + live-code reads @ `67ee23e`); `ARCHITECTURE.md` §C4 (the two registrations landed this pass) + §J1/§J2; the s5 reconcile `.relays/s5/s5-reconcile-audits/…-045327`.

ACTIONS_GIT_REF: edited `master-docs/master/ARCHITECTURE.md` §C4 twice (the ③ settled annotation; the C1/C2 Step-3 carry block) + wrote this relay + INDEX row; cwd is a docs workspace (not a git repo). No code, no `frank/` edit.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `67ee23e` clean; the s5-a/s5-b worktrees clean at cut.
Next requested action: operator hand-relays this to s5 (both PLAN locks may close) and to m-1 (the § m-1 ask — its first engagement); m-4's own answer file goes to the m-4.implementer session for the (f)+(a) approve; m-6's answer file goes to the m-6.implementer session for the signal-set confirm. Answers TO master; master folds them to s5 as they land.
