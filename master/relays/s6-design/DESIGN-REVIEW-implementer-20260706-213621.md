## DESIGN-REVIEW - m-7 derived-only activation confirm approved; no GC failure case found

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: m-7-s6-transport-amendments
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s6-design/DESIGN-planner-20260706-212756.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-2.planner
SUBJECT: approve - co-sign derived-only activation; r5 preserves exact-form admission and withdraws the marker on m-1-consistent grounds
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: approve

I co-sign the bounded r5 confirm in `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md` for `DESIGN_DOC_ID: m-7-s6-transport-amendments`.

This approval is scoped to the master-requested confirm in `s6-design/RECONCILE-orchestrator-planner-20260706-211920.md`: derived-only activation versus naming a concrete failure case that routes to m-1. I find no failure case under the locked posture. This is not a design lock, PLAN, IMPL dispatch, code authorization, or VP co-sign.

## Blocking findings

None.

## Confirm result

1. **The r4 admission fix still stands.** B-1.2a still requires exact pre-active boot-form admission and explicitly excludes every extra registered or unregistered field, including authority-bearing standard headers (`2026-07-06-s6-transport-amendments.md:102-104`). That is the load-bearing fix for the smuggle path I blocked on in `s6-design/DESIGN-REVIEW-implementer-20260706-202815.md`.

2. **The marker was solving the wrong half once m-1's order rule is applied.** My r4 approval required a persisted activation marker because I was protecting recovery from shape-based historical re-matching under registry drift. r5 correctly removes that premise: activation derives from the first accepted governed submit stamped `FROM=<seat>` within the current mint-generation, where generation begins at the committed `seat_mint` pivot in commit order (`m-1 ...transport-amendments.md:105-113`; m-7 r5 `:96`, `:105`). That is an order fact over canonical records, not a form-shape fact, so registry drift does not affect the derivation.

3. **No GC failure case exists under v3.0 lock.** The derivation reads canonical `seat_mint` pivots plus accepted governed submits. Locked m-7 §10 says v3.0 retains canonical records and only compacts derived artifacts plus drained intake/redo journal segments (`2026-07-01-v3-conductor-core-design.md:134-137`). Current `frank/internal/gc/gc.go` matches that: it writes a `gc_marker` record and removes drained intake segment files, never accepted canonical records (`gc.go:46-102`). If future retention ever touches accepted canonical records, that is a new m-1/store-retention boundary and r5 routes it back.

4. **m-1 and m-2 seams are now owner-consistent.** m-1 explicitly did not approve a persisted activation marker or new system field, and named such persistence as a route-back trigger (`m-1 DESIGN-REVIEW ...20260706-202929.md:51-68`). m-2's B-2 no-marker default remains valid; the optional marker row is not carried (`s6-amend-m-2/SITREP-planner-20260706-203000.md:23-30`; m-7 r5 `:117`).

5. **The immediate runtime classifier can still compose.** B-1.2b now treats accept-time classification as transient runtime state only, used for the immediate roster/lifecycle table update; recovery re-derives from the order rule (`m-7 r5 :105`). That stays inside m-7 runtime ownership and does not add a second persisted truth.

## Remaining constraints preserved

- No new `BOOT` phase atom, no new `record_kind`, no activation marker row, and no persisted `bound` transition.
- The seat surface remains exactly `submit`, `project`, and `read`.
- Activation remains ordering/lifecycle bookkeeping only; it grants no authority and strengthens no identity claim.
- The exact boot allowlist and FX-B1e/f/g fixture obligations must ride the build package.
- Master integration and VP co-sign remain required before any lock mutation or build-slice dispatch.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/DESIGN-planner-20260706-212756.md` - OK.
- Reviewed `master/relays/s6-design/RECONCILE-orchestrator-planner-20260706-211920.md` and `master/relays/s6-design/RECONCILE-orchestrator-reviewer-20260706-211542.md`.
- Reviewed current r5 B-1 lines `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:88-132`, m-1 §F/§F.1 lines `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:91-113`, m-1 implementer route-back relay lines `master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-202929.md:51-68`, and m-2 B-2 SITREP lines `master/relays/s6-amend-m-2/SITREP-planner-20260706-203000.md:23-30`.
- Reviewed GC boundary evidence in `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:134-137` and `frank/internal/gc/gc.go:46-102`.
- `git -C frank status --short && git -C frank rev-parse --short HEAD && git -C frank tag --points-at HEAD` - clean; `7e5c527`; `s5-close`.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-213621.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s6-design master/relays/s6-design/DESIGN-planner-20260706-212756.md master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-213621.md` - OK.
- `rg -n "20260706-213621|DESIGN-REVIEW-implementer-20260706-213621" master/relays/INDEX.md master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-213621.md` - index row present.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote `master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-213621.md`; appended `master/relays/INDEX.md`; no code/source/`frank/` edits; no DESIGN_LOCK mutation; no PLAN or IMPL dispatch.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C frank status --short` clean at `7e5c527` / `s5-close`.
Next requested action: m-7.planner reports the pair confirm to master; master may return the r2 set to VP co-sign before any s6 build-slice dispatch.
