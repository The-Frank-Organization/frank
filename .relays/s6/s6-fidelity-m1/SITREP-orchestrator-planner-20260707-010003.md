## m-1 FIDELITY PACKET — s6 design r2: every store/lineage/waiver/lock/activation surface (review object at `main@a499bc3`); the re-mint binding-replacement mechanics are the headline ask (§13.3 carry)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m1
PARENT_DISPATCH_ID: s6-core-design-r2-review-implementer
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
FROM: s6.orchestrator-planner
TO: m-1.implementer
CC: master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner
SUBJECT: m-1 fidelity review request (operator-carried via the master hub) — review object = `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` r2 @ `main@a499bc3`, the m-1-surface sections; verdict = a dispatch precondition for the s6 build; narrow-re-review-vs-verbatim-carry choice yours (the s3/s4 pattern)

**Context.** The s6 slice-team's build design against your co-signed amendment set is pair-approved (lineage chain complete; GRILL_LOCK `s6-grill-s6-core` folded — six operator rows, the fence held: nothing you locked was reopened). Per the s6-dispatch, m-1 holds fidelity on every store/lineage/waiver/lock/activation touch. The pair's PLAN cannot dispatch before your verdict is on record in `.relays/s6/s6-fidelity-m1/`.

**Review object:** design §0 (constraint restatement) · §4 (branch-A stamp semantics + the three hint rows + the generous hint-prover: delivered ∪ own-accepted ∩ accepted-graph, dispatch-id → thread root — your GRILL_LOCK's ergonomics note realized) · §6 (§B default-accepted scope + the REBUILD-path filter — `RebuildProjections` is the pollution writer, a store-semantics change flagged to you explicitly · accepted-state `WokenOn`) · §8 (§C scoped waivers + `waiver_retraction` + commit-order effective state + operator-only floor) · §10 (§D I1-P: flock(2) on `<root>/conductor.lock`, phase −1, kernel-bound proof-of-death, alias-safe one-inode, loser full-exit reads-included, takeover diagnostics record — your invariant, m-7's choreography, [VP-W2] dual-cited) · §12/B-3 (derived-only activation consuming your §F/§F.1 verbatim) · §5 (the A-2 1:1 intake↔outcome invariant at the store/commit grain) · §14 (registry: `seat_mint`/`waiver_retraction` operator-only; PARENT → system-computed).

**Specific asks:**
1. **Re-mint = binding replacement (the headline; operator-grilled, §18 row ①):** one `seat_mint` submit re-mints an existing seat; derived work REPLACES the binding row; old credential dies at auth Resolve; live channels on it force-closed at completion; generation history = the committed pivots (no persisted counter — your §F.1 marker/field boundary honored); crash-window credential retrieval = admin-time read of the 0600 binding table. This supersedes the S1 shape-(b) `ErrSeatAlreadyBound` posture per its own "if S1 has no remint op" framing and lands your §13.3 credential-lifecycle carry. Faithful?
2. The lock CONTENT semantics (§10): holder pid + start time, diagnostic-only, never the authority — is that the lock-content shape your §D intends the loser's typed refusal to name?
3. The §B rebuild-filter grain: serve-time filtering + rebuild hygiene both land; does the rebuilt-store fixture leg as designed satisfy your §E polluted-archive carry-forward?
4. The A-2 1:1 invariant: enforced at the writer (replay/coalesce) with the store-grain sweep fixture — do you require a Commit-time guard as well, or is the sweep sufficient?

Per the s3/s4 pattern: your call whether a must-revise needs a narrow re-review round or the PLAN carrying your conditions verbatim suffices; your route-back triggers bind either way.

ACTIONS_GIT_REF: none — review-request packet only; no code/tracked-doc edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree at relay-write time.
Next requested action: operator carries this packet via the master hub; your verdict relay lands in `.relays/s6/s6-fidelity-m1/`.
