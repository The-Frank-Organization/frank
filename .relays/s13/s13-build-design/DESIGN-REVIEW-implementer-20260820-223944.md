## DESIGN-REVIEW — DS-s13-m10-module r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-design-review-2
PARENT_DISPATCH_ID: s13-build-design
RUN_ID: s13
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — master must rule E-0 artifact location/write fence and E-1 SQLite version/support floor before this design can receive an approve verdict
FILED_AT_LOCAL: 20260820-223944
DESIGN_DOC_ID: DS-s13-m10-module
DESIGN_DOC_SHA256: 84444951f56ff9ace316205e23245435d4fb1bc3e48368425e5fae7d1e50d06f
DESIGN_REVIEW_VERDICT: human-decision-required
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: human-decision-required on DS-s13-m10-module r2 — F2/F4 close; E-0/E-1 rulings still gate a final design approval

## Verdict

`human-decision-required`. The r2 byte digest matches the routed claim. The local mechanism corrections close
S13-DR-R1-F2 and S13-DR-R1-F4 at the correct design loci. S13-DR-R1-F1 and S13-DR-R1-F3 are now honestly filed
as E-0/E-1 master escalations, but escalation is not disposition: the artifact location remains mechanically
invalid under the active relay root, and the implementation design still carries two mutually exclusive SQLite
version/support-floor routes. A scalar `approve` would falsely lock unresolved bytes and ownership.

No implementation authority follows from this review. No PLAN may lock r2 as an approved design until master
rules both items and the planner returns the ruling-bound successor for review.

## Prior-finding disposition

### S13-DR-R1-F1 — open, correctly escalated as E-0; master decision required

r2 withdraws the false fence-conformance claim and the filed
`s13-build-design/SITREP-planner-20260820-223211.md` routes the exact conflict TO
`master.orchestrator-planner`. That is the correct ownership path. The underlying predicate is not yet closed:

`python3 .../relay-lint.py --relay-root .relays/s13 --no-freshness`

still exits 1 because both `.relays/s13/docs/designs/DS-s13-m10-module-20260820.md` and
`.relays/s13/docs/plans/PL-s13-build-plan-20260820.md` are recursively parsed as relay payloads and lack relay
headers/action evidence. Exact-file lint of the r2 DESIGN relay and INDEX lint both pass. Master must authorize
the recommended standard sprint-tree fence addition, or authorize a different lane-doc mechanism. After the
ruling, move through the ruled owner/fence and prove the full active root clean.

### S13-DR-R1-F2 — closed

Design lines 191–211 now make the retirement/disposition transaction total over the post-parking outcome. The
common prefix commits the full parking batch and E+1 fencing; only the ordinary `count <= 512` branch allocates
G+1. The `count > 512` branch commits the `FAILED/parked_unknown_capacity_exceeded` terminal with no successor
generation, same-run continuation, lease, snapshot, or revival, and binds FX-M10-CAP (i)–(iv), including the
multi-row/no-truncation case. This matches the frozen producer delta's complete cap-terminal predicate.

### S13-DR-R1-F3 — open, correctly escalated as E-1; master decision required

r2 now states the real compatibility boundary and files both exact routes. Fresh authoritative module-proxy
reads reproduce the claimed floors: `v1.36.1` declares `go 1.21`; `v1.36.2`, `v1.36.3`, and `v1.37.0` declare
`go 1.23.0`; every checked release `v1.50.0` through `v1.57.0` declares `go 1.25.0`. Frank still declares
`go 1.22` and advertises Go 1.22+. Therefore Route A (`v1.57.0` plus a Go 1.25 support-floor change and README
edit) and Route B (`v1.36.1`, floor-preserving) are materially different designs. Master must choose; then r3
must bind the exact pin and the ruled owner/scope rather than retain both branches.

### S13-DR-R1-F4 — closed

Design lines 281–300 now specify complementary production/reduced constraints, P1 with the required polarity,
P2 in both selected files, and an independently tagged compile-negative fixture. `TestLimitsCompileMatrix` names
the ordinary build, reduced-tag positive build, and P2-violating negative build, including the diagnostic
predicate. This is an executable implementation contract rather than a prose-only assertion.

## Required ruling-bound successor

1. Carry master's E-0 ruling into the artifact location/fence and demonstrate full `.relays/s13` root lint clean.
2. Carry master's E-1 ruling into one exact SQLite version and one exact Go/README ownership outcome.
3. Reissue the same `DESIGN_DOC_ID` with a new digest and request a fresh DESIGN-REVIEW. Preserve the closed F2/F4
   mechanism bytes unless the ruling itself requires a named delta.

## Verification

- Design SHA-256: `84444951f56ff9ace316205e23245435d4fb1bc3e48368425e5fae7d1e50d06f` — matches relay.
- Plan-draft SHA-256: `ff3bc8d2e8d67adff445c1340befa802d16b5df53f9f2df8adb9170788fb63f9` — matches relay; remains unlocked.
- r2 DESIGN relay exact-file lint: pass.
- s13 INDEX lint before this return: pass.
- Full s13 relay-root lint: fail only on the two E-0 sprint-doc payload collisions named above.
- No bare implementation token exists in the addressed r2 relay; no source/test/branch/commit action is authorized.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay + one live-EOF s13 INDEX row; no source/test/branch/plan/design/store/token byte changed by s13.implementer
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s13/docs/designs/DS-s13-m10-module-20260820.md
 M frank/.relays/s13/docs/plans/PL-s13-build-plan-20260820.md
 M frank/.relays/s14/INDEX.md
 M frank/.relays/s15/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-221508.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-223944.md
?? frank/.relays/s13/s13-build-design/DESIGN-planner-20260820-223212.md
?? frank/.relays/s13/s13-build-design/SITREP-planner-20260820-223211.md
?? frank/.relays/s14/s14-build/
?? frank/.relays/s15/s15-build-2/
?? frank/.relays/s15/s15-build/PLAN-REVIEW-implementer-20260820-220110.md
