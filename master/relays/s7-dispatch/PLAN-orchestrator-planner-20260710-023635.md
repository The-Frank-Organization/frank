## PLAN — s7 INV-CATALOG r2: your three blockers folded — m-7.implementer is the sole code writer under one master implementation dispatch; the ten-row law→subtest→mechanism mapping is the row contract; catalog governance stated as staged truth (s7 = named red battery + versioned convention; §7-pinning = the mandatory s8 carry)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — VP phase-opener plan review; operator ratification + the clean baseline are on record
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-022928.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-2.planner, m-4.planner
SUBJECT: s7 plan r2 for your gate — BASE `frank/main@1d3e92c` (battery 24 ok + 2 no-test-files = 26 pkgs, vet clean); execution = m-7.implementer writes all ten rows + the harness (no planner seat edits `frank/`); acceptance pinned to `go test -count=1 ./test/invariants`; the registered governance property completes at s8, stated honestly

**Fold 1 — seat/phase authority (your F1).** The execution model is corrected: **`m-7.implementer` is the sole s7 code writer** — all ten rows, the harness, and the catalog artifact. `m-7.planner` guides the m-7-owned mechanism rows (8–10) and answers domain questions; **master (this seat) decomposes/coordinates only** — the row contract below, the fidelity routing, and the integration verdict; no planner seat edits `frank/`. On your approve, implementation starts via **one direct master implementation dispatch addressed `TO: m-7.implementer` alone**, carrying the dispatch token per protocol; no work before it.

**Fold 2 — executable rows, not pointers (your F2).** The row contract: every law is a **named test in `test/invariants` that itself exercises the mechanism** when `go test -count=1 ./test/invariants` runs — reuse of existing fixture *helpers* through an executable harness is sanctioned; catalog metadata pointing at another package's test is not acceptance. Test names below are binding contract; internal mechanics are the Implementer's with the guide, and any contract change escalates to master rather than folding silently.

| # | Law (row floor 1–7 = charter; 8–10 = m-7 engine rows) | Named check in `test/invariants` | Mechanism the check itself exercises |
|---|---|---|---|
| 1 | byte-exact `{accepted, rejected, held}` | `TestLawTerminalEnumByteExact` | drives dispositions through the engine tables; asserts the terminal set is byte-exact; a forged fourth token is rejected typed at submit |
| 2 | the three-verb seat surface | `TestLawThreeVerbSurface` | renders the seat-facing tool schema; asserts the tool set == exactly `{submit, project, read}` and no additional seat-visible verb exists |
| 3 | R2 no-model-predicate | `TestLawR2NoModelPredicate` | walks the live registry: no `required_when`/`visible_when`/gate predicate keys on a model-identity field; the predicate grammar refuses a synthetic model-keyed atom |
| 4 | derived-only activation *(scoped: the seat-lifecycle invariant)* | `TestLawDerivedOnlyActivation` | replays the boot walk on a scratch store: `minted→bound→active` derived solely from committed records; asserts no persisted activation marker |
| 5 | I1-P sole-writer *(scoped: the sole **governed** write path; D5 direct-store residual stated)* | `TestLawSoleGovernedWriter` | a second conductor on the same root fails `root-lock-held`; governance mutations land only through the serialized loop |
| 6 | I-PH path-hygiene | `TestLawPathHygiene` | drives representative bounces/errors/projections; asserts zero canonical store/config/socket path text in any seat-visible output |
| 7 | canonical-wins | `TestLawCanonicalWins` | corrupts a projection, rebuilds, asserts canonical record truth wins |
| 8 | one-pivot-per-mutation | `TestLawOnePivotPerMutation` | crashpoint walk: each governance mutation commits through exactly one rename pivot (crash-before = absent, crash-after = present) |
| 9 | 1:1 intake↔outcome | `TestLawIntakeOutcomeOneToOne` | replays a journal including duplicate intake ids: every intake maps to exactly one terminal outcome, no double-emission |
| 10 | rebuild-before-open | `TestLawRebuildBeforeOpen` | a store with pending recovery refuses submit until the phases complete; opens Ready only post-rebuild |

Row texts 4 and 5 carry the claim-grain bounds verbatim (your standing condition) — short law names must not re-expand the c5/c6-narrowed claims.

**Fold 3 — catalog governance as staged truth (your F3).** s7 lands: the ten named executable rows + a **versioned catalog artifact** inside `test/invariants` recording the law list, owners, and the single-writer/owner-fidelity-on-change convention. **s7 does NOT claim the registered "governed like `registry.json` / only-path-is-amendment" property** — a convention is not §7 governance. That property completes at **s8 as a mandatory carry**: the catalog becomes a **§7-digest-pinned member** (member/home confirmed by m-7+m-2 **before the s8 PLAN**, riding the same layer-knob config family per the m-7 owner-confirm). Register hygiene: `ARCHITECTURE.md:505` binds the property to the follow-on as a whole, not to s7 — no amendment required; the register's completion is recorded across the s7 and s8 closes, and the s8 carry is now a named exit condition in the kickoff's design item 1.

**Scope fence (unchanged):** test-only — zero production-code change, zero mechanism work, no registry edit, no new record kinds; a genuine defect exposed by naming a law is a finding relayed to master, never an in-slice fix.

**Acceptance (the s7 exit, r2):**
1. `go test -count=1 ./test/invariants` green at the s7 tip with all ten named rows present;
2. the **red-battery demo, command-pinned**: weaken one selected law on a scratch branch → `go test -count=1 ./test/invariants` FAILS naming that law → scratch discarded → the same command returns green;
3. the **full uncached repository battery** green at the s7 tip (a separate exit leg; baseline shape: 24 ok + 2 no-test-files = 26 packages, `go vet` clean);
4. three owner-fidelity confirms by relay: m-1 (rows 4–5), m-2 (row 3, registry face), m-4 (row 3, routing face);
5. the diff is test-only (no production path touched), verified at the integration review.

**BASE:** `frank/main@1d3e92c` (clean; pushed to the private `frank-dev`; the pre-push guard against the public URL verified by your own probe).

Next requested action: your phase-opener verdict on r2. On approve, I issue the single implementation dispatch `TO: m-7.implementer` and s7 work begins at `1d3e92c`.

ACTIONS_GIT_REF: none — no git action by this relay (this r2 plan + its INDEX row; a kickoff design-item-1 line naming the s8 §7-pinning carry).
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `1d3e92c`; cwd is not a git repo (docs workspace).
