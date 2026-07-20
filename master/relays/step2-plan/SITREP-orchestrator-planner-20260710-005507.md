## SITREP — STEP-2 OPENED by the operator: the kickoff package is cut (`master/STEP-2-KICKOFF.md`); requesting your adversarial review of the scope fence, the s7–s10 decomposition, and the hardened exit gate before `step2-plan` formalizes

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step2-plan
PARENT_DISPATCH_ID: s6-exit-gate
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — three operator ratifications are requested (kickoff + naming + the egress-stays-fixture-scoped ruling) and one operator pre-flight (the `frank/` baseline commit); none block your review
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s6-exit-gate/RECONCILE-orchestrator-planner-20260708-031330.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: Step-2 kickoff (operator-opened 2026-07-10) — PLAN-approach at `master/STEP-2-KICKOFF.md`: INV-CATALOG first (s7), then the vertical observe spine (s8) → evidence thicken (s9) → the m-6 comms mechanism (s10) → step-exit; away-bridge fenced OUT (step-(d)); `step2-prep` = 8 bounded design items; charter + dashboard flipped

**What happened:** the operator opened Step-2 ("governance hardening + minimal comms", ROADMAP:96-101). I re-grounded against the design-of-record (m-3 `c2-lock` + folds; m-6 `c3-lock` + folds; m-7 `c4-design-m-7-lock`), the s6 close state (zero F1–F17 deferred; two typed OIs riding; INV-CATALOG queued first), and the in-tree dormant scaffold (registry Block A behind `layer_present:observe`; the mapped commit-path slots; the built park/wake primitives; the dormant egress scanner with the decision-⑤ carve-out already coded). The kickoff package is `master/STEP-2-KICKOFF.md`; the charter rule-1 phase line and the dashboard banner are flipped.

**The decomposition proposal (for your teeth):** s7 = **INV-CATALOG** (the close record's queue-first item — 7 standing laws, one named executable check each, catalog governed like `registry.json`; lands the tripwire *before* observe code exists). s8 = the **observe spine**, vertical-slice-first: governed layer activation → `observe_gate()` in the serialized commit path (post-form/lineage, pre-append) → one E1 + one E2 check → conductor-stamped `achieved_evidence`/`record_integrity` → false-done typed bounce / true-done observed accept → decision-② both legs → the s5-b (h) suppliability guard. s9 = **evidence thicken** (check-registry + isolated executor + `scope_paths` + the two riding OI-S6s + the ⑤ fixture pair). s10 = the **m-6 mechanism** (7-state FSM, resummon cadence, re-observe-on-wake, wake-on-reply, A/B/C/D buckets, ODB render+capture, Q6×Q4 widening reconcile, 8a). Step-exit = the ROADMAP three legs live on the dogfood store + the INV-CATALOG red-battery demo.

**Rulings I am proposing (operator ratifies; flag if you disagree):** (1) naming — slices continue s7…, cycle dispatches `step2-prep`/`step2-plan`; (2) **egress stays a fixture-scoped scanner through Step-2** — the live chokepoint activates with the first external send (the away-bridge, step-(d), post-Step-2), consistent with ROADMAP:49-51 local-only comms and both domain surveys; (3) the away-bridge (Seam-C token + decision-④ rotate/re-observe) is OUT of Step-2 scope.

**Where I most want adversarial pressure:** (a) the s8/s9 boundary — is the spine honest-thin, or does executor isolation belong wholly in s8 (it build-gates E2, and a spine whose only E2 check skips isolation could ship a false floor)? (b) the s10 single-slice m-6 load — FSM + buckets + ODB may be two slices; (c) the `step2-prep` list — 8 bounded items, none re-design; tell me which are PLAN-blocking vs parallel; (d) the layer-activation knob — I've required "governed, visible act, never a silent code default" but left the mechanism to step2-prep; if you think it needs an operator decision (§7-record vs engine-config), say so now; (e) anything in the exit gate that is still happy-path.

**Operator asks (CC'd):** the `frank/` baseline commit (42 uncommitted publication-sweep paths on `main` at `s6-close` — one commit before s7 opens) · ratify the kickoff + the three rulings above · the dogfood relaunch rides s8's first live store per the close record.

Next requested action: your review verdict on the kickoff package (approve / must-revise with findings). On approve + operator ratification, I cut `step2-prep` dispatches (the 8 bounded items to their owning pairs) and the s7 INV-CATALOG dispatch.

ACTIONS_GIT_REF: none — no git action by this relay (docs-workspace edits only: `STEP-2-KICKOFF.md` new; `CLAUDE.md` rule-1/rule-2 phase lines; `master/README.md` banner + org-status entry; this relay + its INDEX row).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` untouched by this relay, still `main@a1bc6d4` (post-`s6-close` fold) with the 42 publication-sweep paths uncommitted, tag `s6-close` intact.
