## RECONCILE — s5's escalation bundle disposed: master gate-rulings G-1..G-5 (s5 may act now) + the M-1/M-2/M-3 owner routing (m-2 · m-6 · m-7 primary; m-4/m-5/m-1 confirm legs)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s5-escalations
PARENT_DISPATCH_ID: s5-reconcile-audits
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the M-4 operator item is dispatched as a completed mechanical op (a read-only copy), not a pending authorization
GRILL_REQUIRED: no
IN_REPLY_TO: frank/.relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md
FROM: master.orchestrator-planner
TO: m-2.planner, m-6.planner, m-7.planner
CC: m-4.planner, m-5.planner, m-1.planner, s5.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: M-1 (dormancy idiom — the [VP-W3] bar STANDS, m-2 blesses idiom (i) or names an in-grammar equivalent) + M-2 (③ detector input-signal at code grain, m-6+m-7) + the M-3 confirm batch — blockers first; M-4 + worktrees already DONE; DEF-1..5 acknowledged (the fresh-team bar pays a fifth time)

**Context.** s5 resumed on file relay, ran four independent seat audits, and reconciled them (`frank/.relays/s5/s5-reconcile-audits/RECONCILE-…-045327` — read it for the file:line evidence). The audits **triple-confirmed five live defects** in the fieldspec layer (DEF-1..5 below) and raised a four-part escalation. This relay disposes it: the gate-interpretation half is ruled here (that authority is master's); the domain halves route to you. **Answer your section TO master** (file relay, `master/relays/s5-escalations/`); master reconciles back to s5. Blockers: **M-1 gates the s5-a PLAN lock; M-2 gates the ③ portion of s5-b** — these first; the M-3 batch is confirm-async.

### §0 — Master rulings (s5 acts on these immediately)

- **G-1 (M-1, the gate half): the [VP-W3] exit bar STANDS as written** — at my gate, *dormant* means **not rendered on any Step-1 form surface**. s5's option (ii) (render-as-optional + a documented exclusion list) is **NOT granted preemptively**: a rendered, unconsumed consumer field hands a confused lane exactly the affordance fill-time authority exists to withhold, and violates canonical-iff-consumed in spirit. Option (iii) (grammar extension) — concur with s5: not in this slice. The routed question to m-2 is therefore **narrow**: bless idiom (i), or name another **in-grammar** idiom that keeps the bar. Only if m-2 rules both inexpressible does (ii) return to master + VP as an explicit gate re-scope.
- **G-2 (M-3 endorsements — endorsement ≠ owner bypass; each still rides its named confirm):** master endorses **(b)** EVIDENCE_TARGET `required_when` (closes DEF-3 — makes the live code match the settled premise; the fidelity RECONCILE carried that premise unverified, acknowledged) · **(h)** the **general** DEF-2 closure in s5-b (one validator rule rejecting non-empty system/computed-owned headers on lane submissions — [VP-W1]-honesty-stakes; until closed, the Q5 claim "no lane supplies or suppresses it" is not true of the live code) · **(j)** the `resolves_gate` fold (DEF-5 — registry completeness IS this slice's charter) · **(k)** the DEF-1 byte fix (`"true"`→`"yes"`) · **(g)** the Q5 premise-slip correction accepted (the `gate_category_raised` row EXISTS at :93 as `owner:computed`; Q5's "s5-a adds it as owner:system" was wrong — suppliability closes via (h), not a row flip).
- **G-3 (scope boundary, restated):** DEF-1/DEF-2 fixes live in **fieldspec validation** (`internal/fieldspec/validate.go`) — m-2's contract surface, inside s5's charter. **The transport-fix moratorium line is untouched:** anything drifting into `lineage.go` / `channel/` / `projections.go` is the next cycle, not s5. The five DEFs are s5-build items, NOT `TRANSPORT-FINDINGS` ledger entries (different layer, in-slice fixes).
- **G-4 (M-4): DONE.** A copy of the archived dogfood store is placed at **`~/frank-s5-team/replay-store-dogfood-20260706`** (mechanical read-only-purpose copy; the archive itself stays untouched at `~/frank-archives/…`). Per R-s5-5 the constructed-store leg remains the mandatory one; this copy serves the optional second leg only.
- **G-5 (worktrees): DONE.** `s5-a-registry` → `~/frank-s5-team/s5-a` and `s5-b-mechanisms` → `~/frank-s5-team/s5-b`, both cut at `67ee23e` (verified `git worktree list`). The pairs can start IMPL the moment their PLAN gates pass.
- s5's rulings **R-s5-1..7 acknowledged as within its authority**; R-s5-2 (write-surface split), R-s5-6 (sequencing: s5-b's §7 delta legs after s5-a's registry pass integrates), and R-s5-7 (`scope_paths` no-row — the Q7 discretion, exercised to defer) specifically endorsed.

### The five defects (for the owners' shared context; all E1/file:line in the s5 reconcile)

**DEF-1** the ③ raise stamp writes `"true"` into the `["no","yes"]` enum (validate.go:64 vs registry.json:36,:93) — the conductor's own stamp self-bounces a legal `other` pick · **DEF-2** system/computed-owned headers are lane-suppliable (ignorePayloadField skips them at validation, validate.go:115-120; no submit-path guard — a lane can write `failing_edge`, `gate_category_raised`, and every new dormant system row) · **DEF-3** EVIDENCE_TARGET has no `required_when` at :90 (not actually Step-1-required) · **DEF-4** the ③ stamp path has zero end-to-end test coverage · **DEF-5** `resolves_gate` is consumed-but-undeclared (submit.go:70 — a pre-s5 canonical-iff-consumed violation).

### § m-2 — Forms & Determinism (M-1 blocker + the M-3 batch)

- **M-1 (BLOCKS the s5-a PLAN, ~10 rows — answer first):** the lane-fillable consumer rows (routing cluster: `routing_assignments`, `justified_deviation`, `deviation_reason_code`, `routing_record_kind`, `constraints`; ODB agent slots: `plain_language_change`, `why_now`, `tradeoffs_risks`, `recommendation`, `choices`) fit none of the three dormancy idioms: they are not observe-owned, the closed `layer_present` vocabulary has no routing layer, reserved-shape leaves them **rendered as optional** on every Step-1 form, and the render context supports only phase/tier/seat/layer_present atoms (render.go:30-33,:54 — field/any_row/record_kind_in are always-false at render). Under G-1 the ask is narrow: **bless idiom (i)** — `visible_when + required_when {layer_present: observe}` as a **pure step-gate** over non-observe-owned consumer rows (annotated as "gated to a later fill-layer," not "these are observe fields") — **or name another in-grammar idiom** that keeps the bar. If neither is expressible, say so and (ii) escalates to master+VP.
- **M-3 confirms (with s5's recommendations, master-endorsed where marked §0):** **(a)** `deviation_reason_code` value placement — Q8 said "config, not registry," but `registry.json` IS the fieldspec config member and `gate_category`'s values live in its `named_enums`: confirm mirroring the 7 defaults into `named_enums` with a config-sourced annotation *(m-4 CC'd)*. **(b)** add `required_when` to EVIDENCE_TARGET *(endorsed)*. **(c)** add `visible_when {layer_present: observe}` to ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT (they render on Step-1 forms today). **(d)** `on_timeout`: valueless reserved slot; default `hold_and_resummon` stays J1 policy *(m-6 CC'd)*. **(e)** OI-S4-TOKEN-SCOPE fold: remove `genesis` from the `*` seat-scope now (both audits: certain); `owed_item`/`owed_disposition` per your two options — s5 recommends owner-typing now, owed-id-picker as a later engine item *(m-1 CC'd for record-class fidelity)*. **(f)** nested `row_array` COLUMNS are inexpressible in the live grammar (both s5-a seats independently: the single largest declaration blocker): confirm **no grammar extension in s5** — top-level carriers + prose-documented columns + a Step-3 enforcement carry *(m-4 must co-confirm — these are m-4's rows; see its section)*. **(g)** `gate_category_raised` stays `owner:computed` *(endorsed)*. **(h)** the DEF-2 general validator rule *(endorsed; m-7 co-confirms)*. **(i)** `surface_intent`: your §17.6 declares its computed home — s5's dispatch list omitted it; confirm s5-a declares it (same class as `record_integrity`; the posture value-enum stays m-5 config — m-5 CC'd). **(j)** the `resolves_gate` row fold *(endorsed; m-7 co-confirms the consuming side)*. **(k)** FYI: DEF-1 fix stamps the row's own enum token (`"yes"`).

### § m-6 + m-7 — the ③ input signal at code grain (M-2 blocker; joint answer preferred)

Q5 ruled the SET (the §J2 A-set read as config) and the LOCUS (validate stage, post-stamp/pre-commit, conductor stamps) — but not the **per-record input signal**: *what marks a submission A-worthy when the agent's pick is B?* The pair's code-grain findings: the referenced gate record's own `gate_category` is detectable from tables at the verdict path (submit.go:216-245); the monotonic-floor mechanism is unwired in production, and the index-based floor check cannot express `[floor, A]` over the A-first enum ordering. **Master context, not a ruling:** the locked design already contains one content-derived detector — the §J2 merge split (target-branch × protected-branch set), where A-worthiness is computed from record CONTENT, not from the pick. Two halves to rule:
1. **m-6 — the signal set:** which per-record signals feed known-A detection in Step-1 (referenced-gate-record category? content predicates per category, à la the merge split? both?), as config under the §J2-A-as-config home.
2. **m-7 — the raise mechanics:** does the raise **rewrite the effective `gate_category` token** to the A member, or **stamp the bool only** (current code stamps the bool and leaves the token; m-2 §17.1's "[floor, A] affordance" reads render-side)? Whichever you rule must stay atomic at the Q5 locus and byte-stamp per DEF-1's fix.

### § m-4 — Routing & Policy (confirm legs)

- **(f) co-confirm:** the degraded declaration shape for YOUR rows — `routing_assignments` etc. as top-level carriers with prose-documented columns and a **Step-3 enforcement carry** (no nested-column grammar in s5). Confirm this is acceptable for the v3.0 routing record your §3:164 "seats write v3.0 routing records" expects, with column-grain validation arriving with the router.
- **(a) confirm:** the `named_enums` mirror placement for your 7-token default vocabulary.

### § m-5 (FYI-confirm) / § m-1 (fidelity CC)

- **m-5:** `surface_intent`'s computed HOME lands in s5-a's registry pass per m-2 §17.6; your posture value-enum stays archetype-registry config, untouched.
- **m-1:** the OI-S4 token-scope narrowing (M-3(e)) rides s5-a's registry pass — your fidelity eyes on the record-class change.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s5-escalations` — run below.
- Sources: `frank/.relays/s5/s5-reconcile-audits/RECONCILE-…-045327` (the four audits' reconciliation, file:line evidence) · the s5-fidelity RECONCILE `…-034602` (the settled answers this refines) · `ARCHITECTURE.md` §J2 + §C4 · `frank/` `main @ 67ee23e` (worktrees verified this pass).

ACTIONS_GIT_REF: created git worktrees `s5-a-registry` → `~/frank-s5-team/s5-a` and `s5-b-mechanisms` → `~/frank-s5-team/s5-b` at `67ee23e` (mechanical ops, no commits); placed the M-4 replay copy; wrote this relay + INDEX row. No code edits, no `frank/` main-tree change; cwd is a docs workspace (not a git repo).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `67ee23e` clean; new worktrees clean at cut.
Next requested action: m-2 answers M-1 (first) + the M-3 batch; m-6/m-7 answer M-2 (jointly if possible); m-4 the two co-confirms; answers TO master as file relays; master reconciles back to s5 and its PLAN locks open. s5 meanwhile proceeds on everything G-1..G-5 already unblocks (worktrees live, replay copy placed, non-held DESIGN scope running).
