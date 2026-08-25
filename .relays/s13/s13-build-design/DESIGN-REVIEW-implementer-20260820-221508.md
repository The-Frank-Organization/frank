## DESIGN-REVIEW — DS-s13-m10-module r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-design-review
PARENT_DISPATCH_ID: s13-build-design
RUN_ID: s13
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s13.planner owns the required r2 and routes the named write-fence/toolchain arbitrations to master
FILED_AT_LOCAL: 20260820-221508
DESIGN_DOC_ID: DS-s13-m10-module
DESIGN_DOC_SHA256: 713c583f5e4f7004043fd6a995a3e7da772a33d902f88092e44c24ecb9ffa29c
DESIGN_REVIEW_VERDICT: must-revise
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: must-revise DS-s13-m10-module r1 — relay-root collision, cap-terminal successor contradiction, Go-version incompatibility, and incomplete reduced-limit selector

## Verdict

`must-revise`. The effective-spec composition, removed M10-C0 ledger tables, reduced receipt tuple, m-1 carrier
negatives, and broad battery census are directionally faithful to the frozen bytes. Four implementation-blocking
defects remain in the realization design. No PLAN may lock r1 and no build dispatch may issue from this review.

## Findings

### S13-DR-R1-F1 — the chosen sprint-doc location makes the active relay root mechanically invalid

The design places the design and plan under `frank/.relays/s13/docs/` (`DS` lines 333–336), inside the active
`RELAY_ROOT=.relays/s13`. The current linter recursively treats those Markdown files as relay payloads. The exact
command

`python3 .../relay-lint.py --relay-root .relays/s13 --no-freshness`

returns required-header and action-reference errors for BOTH `docs/designs/DS-s13-m10-module-20260820.md` and
`docs/plans/PL-s13-build-plan-20260820.md`. Exact-file lint of the routed DESIGN relay and INDEX lint pass; root
lint does not. This is not inherited noise: the two r1 artifacts create it.

Required r2 correction: route the write-fence/location conflict to master before moving bytes. Choose a durable
artifact location that is both inside an explicitly authorized fence and outside the recursive relay-payload
surface, or obtain an explicit compliant relay-root mechanism. Prove the resulting active root clean with the
full root command. Do not silence the error, add a monotonic marker, or make design docs masquerade as relays.

### S13-DR-R1-F2 — the retirement algorithm commits a successor on a frozen no-successor terminal

`DS` lines 168–171 define THE retirement transaction unconditionally as `E+1 mint -> G+1 prealloc`, with no
parked-cap branch. The governing producer delta lines 80–81 requires the same retirement transaction to commit
the complete parking batch and `runs.state=FAILED/parked_unknown_capacity_exceeded` when the post-commit count is
greater than 512, while committing NO successor generation, same-run continuation, lease, snapshot, or revival.
The design names the stop-reason enum and the FX-M10-CAP battery but never realizes this branch; a test name cannot
repair an algorithm that preallocates G+1 on every retirement.

Required r2 correction: make the single transaction total over both post-parking outcomes. At `count <= 512`,
ordinary successor preallocation may proceed. At `count > 512`, commit fence + full parking + epoch effects + the
single FAILED terminal while committing no successor row/lease/snapshot/continuation. State the no-truncation and
no-revival proof at the mechanism locus, then bind FX-M10-CAP (i)–(iv) to it.

### S13-DR-R1-F3 — the proposed SQLite release silently raises frank's supported Go floor

`DS` lines 119–134 selects `modernc.org/sqlite` and cites v1.57.0 as the authoring candidate while claiming build
hermeticity. Live module metadata reports `modernc.org/sqlite@v1.57.0` has `GoVersion: 1.25.0`. The repository
declares `go 1.22` (`go.mod:3`) and advertises `Requires Go 1.22+` (`README.md:46`). Therefore the candidate cannot
land as merely a go.mod/go.sum dependency update without changing the public/toolchain baseline; that additional
surface is outside this pair's fence and absent from E-1.

Required r2 correction: bind the driver decision to an exact Go-1.22-compatible release, or explicitly escalate
and obtain master authorization for the larger toolchain/support-floor change before plan approval. The E-1
request must name the exact version and resulting `go`/`toolchain` behavior. Evidence command:
`go list -m -json modernc.org/sqlite@v1.57.0` (E2; reports Go 1.25.0).

### S13-DR-R1-F4 — reduced-limit selection is not closed at the Go file boundary

The binding limits artifact requires tag-guarded substitution: production P1 positive, reduced P1 polarity-flipped,
P2 positive in both, and no production selector reaching the reduced table. `DS` lines 241–249 gives only the
reduced file a `//go:build frank_test_reduced_limits` constraint while placing production constants/assertions in
`limits.go`. Unless the production declarations are explicitly excluded under the reduced tag, both files enter
the package and collide; the design does not state the complementary selector. It also promises a build-level
negative without naming an executable command/fixture that demonstrates the intentionally violating table fails
compilation.

Required r2 correction: specify complementary build constraints (for example production `!frank_test_reduced_limits`
and reduced `frank_test_reduced_limits`), preserve P2 in each selected file, and name the executable compile-negative
harness. The reduced-tag build must compile with flipped P1; an independently violating P2 configuration must fail
compilation; the ordinary build must contain no reduced selector.

## Verified non-findings

- The five m-10 basis files recompute to the charter-pinned hashes.
- r1 correctly composes the M10-C0 replacement: `epoch_transitions`/`crossing_ops` absent, always-propose plus the
  two-form gate present, and the amended broker event family used.
- The burst-fold receipt is reduced to `round_identity` + `seq_hwm`, and the genesis `run_disposition` carve-out is
  carried explicitly.
- The reviewed DESIGN relay is exact-file lint-clean; `.relays/s13/INDEX.md` was lint-clean before this return.

## Required r2 acceptance

1. Resolve F1 through master and make full relay-root lint clean.
2. Repair the cap-terminal transaction at the mechanism locus and preserve the full no-successor predicate.
3. Resolve the driver/toolchain compatibility in the design and exact-version E-1 route.
4. Close the build-tag selector and executable compile-negative proof.
5. Reissue DESIGN with a new digest; retain `DESIGN_DOC_ID: DS-s13-m10-module`; route a fresh DESIGN-REVIEW request.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay + one live-EOF s13 INDEX row; no source/test/branch/plan/design/store/token byte changed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s15/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-221508.md
?? frank/.relays/s15/s15-build/PLAN-REVIEW-implementer-20260820-220110.md
