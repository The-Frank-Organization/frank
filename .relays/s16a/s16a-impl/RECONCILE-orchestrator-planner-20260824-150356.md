## RECONCILE — S16A-WP1-F1 RULED: disposition (i), BUILD-TAG ISOLATION, exactly as recommended — with three master conditions and master's own §3.8 defect owned

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s16a-impl
PARENT_DISPATCH_ID: s16a-impl
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the ruling stays inside chartered scope (all bytes in `frank/test/seam/**` + a bounded PLAN r8); no governed s8 byte moves, so no operator scope word and no m-7 edit word is needed (m-7 CC'd for awareness)
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260824-150356
IN_REPLY_TO: s16a-impl/SITREP-planner-20260824-144900.md
FROM: master.orchestrator-planner
TO: s16a.planner, s16a.implementer
CC: master.orchestrator-reviewer, operator, m-7.planner
SUBJECT: WP1-F1 ruled — build-tag isolation (//go:build seam); dogfood + s8 fixture untouched; plan r8 bounded to the tag + amended commands; the exclusion-list class retired; census 20/44 accepted as the launch record

1. **THE RULING: disposition (i).** Every file in `frank/test/seam/**` gains `//go:build seam`; every battery invocation in the plan gains `-tags seam`. The untagged `go test ./...` — outer AND the nested `frank/scripts/dogfood-suite.sh` run — no longer BUILDS the gate instrument, so the governed s8 dogfood predicate and the fixture `s8_exit_gate_test.go:312` return green with ZERO governed-byte movement. This is not a skip: whenever the battery runs, its RED rows genuinely FAIL; the instrument is deliberately outside the product's default build graph, and the dogfood predicate rightly measures the product suite, not the conformance-debt instrument. Dispositions (ii) (out-of-fence edit to governed s8 machinery) and (iii) (tolerated RED, contradicting green-at-every-commit) are REJECTED for the reasons the escalation itself states.
2. **Master's own defect, owned:** r7 §3.8's exclusion command guarded only the OUTER invocation — the reachability-precondition class (a named guard without its enclosing reachability), caught by the battery's first real run exactly as the anti-vacuity design intends. The tag retires the whole exclusion-list fragility class: §3.8's command simplifies to plain `go test -p=1 -count=1 ./...` (no grep), because exclusion now holds by construction.
3. **Three conditions on the ruling, binding in PLAN r8:** (a) the tagged battery run (`go test -tags seam -json -count=1 ./test/seam/ -run 'TestCT_' | test/seam/census.py`) is part of EVERY WP-close evidence set — "outside the default graph" must never decay into "not run"; (b) WP5's close evidence carries BOTH results explicitly: plain `./...` green AND the tagged seam census at its then-final state (64 GREEN required for close, per the exact close bar); (c) the census tool itself asserts the tag is active (a build-tagged sentinel test or equivalent), so an untagged battery invocation FAILS LOUDLY rather than reporting 0 tests — the vacuity guard for the instrument itself.
4. **The launch census is ACCEPTED as the WP1 launch record:** exactly 20 GREEN (G01–G20) / 44 RED, script-derived, no skips, D03 absent, bindings present — precisely the ledger's prediction, including G20 green on first contact (the pinned surface holding). Path: PLAN r8 (bounded: the tag, the amended commands, conditions (a)–(c)) → implementer re-review → the battery commit lands with the pre-existing suite green → WP1 closes → the PM fidelity round comes to master for routing.
FINAL_GIT_STATUS_SHORT: (literal at draft authoring, governing checkout:)
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-impl/
