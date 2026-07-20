## PLAN — s7 INV-CATALOG (the lean phase-opener): one `test/invariants` package naming the standing laws with one executable check each, catalog governed as a visible artifact; BASE = `1d3e92c` (battery 24-ok uncached, vet clean); requesting your phase-opener plan gate

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator ratified the s7-lean model and directed the baseline (landed); your plan gate is the gate this relay requests
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step2-plan/PLAN-orchestrator-planner-20260710-021150.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-2.planner, m-4.planner
SUBJECT: the s7 plan for your gate — scope: test-only consolidation of the standing global laws into `test/invariants` (the 7 charter laws as the mandatory row floor + m-7's 3 engine rows), one named executable check per law, the red-battery tripwire proven on a scratch branch; BASE `frank/main@1d3e92c`; fence: zero production-code change, no mechanism work

**BASE (the F4 gate, satisfied):** the operator directed the baseline; the 1105-entry publication sweep landed as one commit — **`frank/main@1d3e92c`** — with the full uncached battery at it (24 pkgs ok / 0 FAIL, `go vet` clean, status clean). Remote system per the operator's release-separation ruling: `frank/` → the private `iwnlcern/frank-dev` (main + close tags pushed; a pre-push guard refuses the public `iwnlcern/frank` URL); releases ship only from the separate release tree.

**Goal (one line):** every standing global law becomes a NAMED row with ONE executable check in a single `test/invariants` package, so weakening a law = a red battery **naming the law** — Cardinal rule 1 as a compile-time tripwire, landed before any observe code exists.

**The row set:**
- The **7 charter laws** (the mandatory floor, ARCHITECTURE §C4): byte-exact `{accepted, rejected, held}` · the three-verb seat surface · R2 no-model-predicate · derived-only activation · I1-P sole-writer · I-PH path-hygiene · canonical-wins.
- **m-7's 3 engine rows** (its claimed half, from `step2-prep/SITREP-planner-20260710-010806` item 5): one-pivot-per-mutation · 1:1 intake↔outcome · rebuild-before-open. (m-7 also hosts the harness.)
- **Claim-grain watchpoint (your co-sign condition, carried into the check text):** "derived-only activation" scopes to the **seat-lifecycle invariant** (minted→bound→active derived from committed records, no persisted marker); "I1-P" states the **sole *governed* write path with the D5 direct-store residual**. Short law names must not re-expand the claims narrowed in c5/c6 — each row carries a one-line claim statement with its honest bound, and the check enforces the mechanism, not the overclaim.

**Method — consolidate, don't duplicate:** each row wraps (or points at) the strongest existing fixture for that law where one exists (the s1–s6 batteries are full of them), and adds a check only where no single fixture currently *names* the law. Net-new test code is expected to be small; the value is the naming, the single package, and the tripwire semantics.

**Catalog governance:** the law list ships as a versioned catalog artifact inside `test/invariants` (single-writer, owner-fidelity-on-change discipline recorded in the file header); the m-7-proposed §7-pinning of the catalog as a config member is **deferred to s8's layer-knob config work** (same mechanism family; bounded; lands there if m-7+m-2 confirm) — s7 does not touch engine config.

**Scope fence:** test-only — zero production-code change, zero mechanism work, no registry edit, no new record kinds. If naming a law exposes a genuine defect, that is a finding relayed to master, not an in-slice fix.

**Acceptance (the s7 exit):** (i) all 10 rows named + one check each, green at BASE; (ii) the **red-battery demo**: deliberately weaken one law on a scratch branch → the battery goes red naming the law (the step-exit tripwire leg, pre-proven; scratch branch discarded); (iii) the full uncached battery green at the s7 tip; (iv) fidelity confirms from m-1 (I1-P/derived-only rows), m-2 (R2 row), m-4 (R2's routing face) by relay.

**Execution model (operator-ratified):** lean — no slice team; m-7 executes its claimed half (engine rows + harness), I execute the coordination + the cross-domain rows against the fidelity confirms; both under this plan once you gate it.

Next requested action: your phase-opener plan verdict (approve / must-revise). On approve I cut the m-7 execution dispatch and s7 work begins at `1d3e92c`.

ACTIONS_GIT_REF: the baseline commit `1d3e92c` on `frank/main` + the push of main/tags to the private `frank-dev` remote (operator-directed, recorded above); no s7 work begun.
FINAL_GIT_STATUS_SHORT: `frank/` clean at `main@1d3e92c` (verified post-commit and post-battery); cwd is not a git repo (docs workspace).
