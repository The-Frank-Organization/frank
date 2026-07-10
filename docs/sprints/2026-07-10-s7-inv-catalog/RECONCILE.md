# s7 INV-CATALOG reconciliation ledger

## 2026-07-10 - intake checkpoint

- Live authority verified: direct master dispatch `PLAN-orchestrator-planner-20260710-032426.md` is addressed only to `m-7.implementer` and carries the bare implementation token.
- Supersession verified: the B10 dispatch supersedes the unbooted `s7-core.implementer` assignment and the prior m-7 stand-down; the effective contract is r2 plus the four r3 replacements plus VP approval.
- Exact base verified: branch `s7-inv-catalog` starts at `1d3e92cc1f9f810da538b0369be9440ffd878f0a`.
- Baseline E2 verified in the isolated worktree: `go test -count=1 ./...` exited 0 with 24 tested packages and 2 no-test-file packages; `go vet ./...` exited 0 with no output.
- Scope remains test-only. No production source, registry, or record-kind edit is authorized.
- Implementation has not started. The intake ACK is waiting on `m-7.planner` co-sign as required by the dispatch.

Next: planner co-sign, then the ten-row invariant package and versioned catalog.

## 2026-07-10 - pair co-sign folded

- `m-7.planner` approved the intake in `.relays/s7/s7-intake-ack/PLAN-REVIEW-planner-20260710-033918.md`.
- Row 4 realization is corrected to the locked B-1 grain: `minted` and `active` are record-derived in commit order; `bound_now` is live channel state and restarts empty. Neither activation nor bound gets a persisted marker.
- Row 6 realization has six named families. The sixth is the `seat_mint` accept-reply, whose only exemptions are fresh credential and endpoint on the operator channel; planted store/config/outbox leakage must still fail inside that family.
- The fixed-point sets for rows 1 and 2 remain literal test bytes. Row 6's AST sink-pattern set is itself versioned catalog content.
- These are realization conditions already present in the locked record, not contract amendments. No second co-sign round is required.

Next: implement the ten named checks inside `test/invariants/**`; production paths remain read-only.
