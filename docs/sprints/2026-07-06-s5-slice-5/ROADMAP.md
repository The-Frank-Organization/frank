# s5 — Slice-5: consumer schemas (the LAST Step-1 slice)

RUN_ID `s5` · charter: the s5-dispatch (master, VP-approved) as amended by `master/relays/s5-resume/PLAN-orchestrator-planner-20260706-034602.md` (file-relay mode; adapted exit gate) · semantics of record: `master/relays/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md` (Q1–Q11) + `master/relays/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` (M-1/M-2/M-3 + MR-1..3) + `…-053113.md` (riding legs closed).

## The goal (one line)
Make frank's typed-envelope registry complete + versioned: declare the four consumer domains' fields (m-3 observe · m-4 routing · m-5 archetype · m-6 gate/ODB) step-gated OFF, land the owed §C4 Step-1 fixtures, ship the registry version + zero-loss replay — so Step-2 turns fields on rather than inventing them.

## Team
`s5.orchestrator-planner` (this seat) + `s5.orchestrator-reviewer` + pair `s5-a` "registry & rows" (sole registry.json writer) + pair `s5-b` "mechanisms & versioning" (engine/bounce/migrate/test-replay). First chartered as the first team to run its own governance ON frank (the dogfood); stood down to file relays at checkpoint after F11 (see `master/TRANSPORT-FINDINGS-2026-07-06.md` — the dogfood's deliverable).

## Task map → outcomes (all landed on main)
| Lane | Work | Landed |
|---|---|---|
| s5-a | Registry pass: 47→83 rows, 14→24 enums, `s5-fieldspec-v3`, routing_escalation delta, OI-S4/053113 record_kind scope (genesis in NO scope; owed operator-only), EVIDENCE_TARGET required (DEF-3), MR-1 `gate_category_pick`, [VP-W3] dormancy fixture (full seat×phase×tier sweep), 8 legacy files under the s5-a-impl-grant fence | `dd8189d`+`dd7d0b5` → merge `afddc56` |
| s5-b | Mechanisms: ③ raise (token REWRITE + `"yes"` stamp + full case table), DEF-2 typed-reject guard, ⑤ dormant egress scanner at the drain leg (R-2; (a)/(b)/(c)), I-PH extensions, zero-loss replay + negotiation legs (no envelope migrator, R-1), §7 s5-delta legs (tree-invariant after MF-5), T2b classifyVerdict operator enforcement | 13 commits → `82524f7` → merge `b30df4d` |
| s5-b wire3 | ③ live detector config: S1+S2 live, S3 input-atom-pending (claim boundary in code), fail-before-serving startup proof; M-4 archive-replay leg (real 41-record store, three-seat green) | `518a88f` → merge `f31d43a` |

## Gates exercised
Two must-revise design rounds per pair · plan-gate re-thread (per-hop DISPATCH_IDs after the shared-id lineage defect) · one bounded scope grant (legacy battery, class-fenced) · adversarial panels (s5-a team-of-4: 2 must-fix folded; s5-b team-of-4: 3 must-fixes + MF-1b/1c/4/5) · the run's first integration-gate catch (MF-5: combined-tree red on branch-green+main-green) · written merge-authorization chain (grant → reviewer approve → execution → orchestrator verification) ×3 merges.

## Registered follow-ons (not s5's)
S3 activation (target-branch atom + operator config, a future registry pass) · away-bridge carries (egress activation semantics, scan caps, Drain all-or-nothing) · step-(d) set per §C4 · the transport-fix cycle (lineage/parenting/codec — the F1–F17 ledger is its input) · O-3/O-4 (m-5's, with master).
