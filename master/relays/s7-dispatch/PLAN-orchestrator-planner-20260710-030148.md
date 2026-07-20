## PLAN — s7 INV-CATALOG r3 (narrow): findings 3–6 folded — the kickoff's contradictory sentence fixed at source; row 6 rewritten exhaustive-with-census; row 9 restored to the locked at-most-one grain; fidelity scopes bound over every owner-touched row. Everything else stands as r2.

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — VP phase-opener plan re-review; no new operator decision
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-024232.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-2.planner, m-4.planner
SUBJECT: s7 plan r3 — a delta on r2 (`PLAN-…-023635`), which is incorporated unchanged except the four corrections below; BASE `frank/main@1d3e92c` unchanged; requesting your gate; the implementation dispatch waits on it

**This r3 is r2 with exactly your four folds.** The ten-row floor, named-test mapping, sole-Implementer execution model (one master dispatch `TO: m-7.implementer`), test-only fence, command-pinned red demo, separate full-battery leg, staged-governance statement, and the s8 §7-pinning carry all stand as written in r2.

**Fold 3 — the kickoff contradiction is fixed at source.** `STEP-2-KICKOFF.md:34` (the s7 strategy bullet) no longer grants s7 the governed-like-`registry.json` property; it now reads: the catalog ships as a *versioned artifact carrying the single-writer / owner-fidelity-on-change convention*, with the registered property completing at s8 via the mandatory §7-digest-pinning carry. It and `:56` now agree; `ARCHITECTURE.md:505` needs no amendment.

**Fold 4 — row 6 (`TestLawPathHygiene`) is exhaustive, not representative.** Replacement row contract: the check **enumerates and counts every current seat-delivered output family** — bounce/reason text · process/tool errors · tool descriptions/results · rendered projections · delivery payloads — and scans the **complete corpus** for the canonical store/config/outbox path families. It includes a **planted-leak negative** (a synthetic path seeded into a scanned surface must turn the check red — the scanner is proven to bite). The census is explicit in the check (named families, counted): **a future seat-visible family that is not in the census fails the check until the census is updated** — a new surface cannot silently ride outside the law. (This is `ARCHITECTURE.md:478`'s *any seat-delivered surface, every bounce/error included*, made executable.)

**Fold 5 — row 9 (`TestLawIntakeOutcomeOneToOne`) restored to the locked A-2 grain.** Replacement one-line claim: *at all times, at-most-one outcome per `intake_id`, unique non-empty `intake_id` per outcome; settled ⇒ exactly one.* Replacement test contract, all five clauses: (i) no `intake_id` ever has more than one outcome; (ii) every outcome carries a unique, non-empty `intake_id`; (iii) every settled/consumed intake has exactly one outcome; (iv) pending entries are re-enqueued at most once by recovery (`intake − outcomes`); (v) duplicate-content retries replay/coalesce to the same outcome. **A pending intake with zero outcomes is legitimate state, never a law violation.** (Per the m-7 s6 transport amendment `:35-39`.)

**Fold 6 — fidelity scopes bound over every owner-touched row.** Same three fidelity seats, review scope now explicit: **m-1 → rows 2, 4, 5, 6** (three-verb shape consumed from m-1 · derived-only · I1-P · I-PH's store half); **m-2 → rows 1, 3, 6** (terminal-state vocabulary · R2 registry face · I-PH's bounce/reason-text half); **m-4 → row 3** (R2 routing face). Acceptance item 4 of r2 is amended accordingly; no team or ceremony added.

**Acceptance (r3 = r2 with item 4 amended):** (1) `go test -count=1 ./test/invariants` green at the s7 tip, all ten named rows present; (2) the command-pinned red-battery demo (fail naming the law → discard → green); (3) the full uncached repository battery green at the s7 tip (24 ok + 2 no-test-files = 26 pkgs, vet clean); (4) fidelity confirms by relay per the Fold-6 scopes; (5) the diff is test-only, verified at integration review.

Next requested action: your verdict on r3. On approve, I issue the single implementation dispatch `TO: m-7.implementer` and s7 begins at `1d3e92c`.

ACTIONS_GIT_REF: none — no git action by this relay (this r3 + its INDEX row; the one-line kickoff `:34` correction).
FINAL_GIT_STATUS_SHORT: `frank/` = `## main...origin/main` clean at `1d3e92c`; cwd is not a git repo (docs workspace).
