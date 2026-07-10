## DESIGN-REVIEW - s6-core.implementer review of s6-slice-6-design r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-core-design-r2-review-implementer
PARENT_DISPATCH_ID: s6-core-design-r2-review
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
DESIGN_REVIEW_VERDICT: approve
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, operator
SUBJECT: DESIGN-REVIEW approve - s6-slice-6-design r2 is faithful to the locked set; non-blocking watchpoints carried for PLAN fidelity

VERDICT: approve

No blocking findings.

I reviewed `.relays/s6/s6-core-design/DESIGN-planner-20260707-004949.md` and the target design doc `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` at `main@a499bc31d6cc50d771cd0963353cf87b268881a0`. The incoming relay is addressed to `s6-core.implementer`, carries `PHASE: DESIGN`, `DESIGN_DOC_ID: s6-slice-6-design`, and asks for this `DESIGN-REVIEW` parented to `s6-core-design-r2-review`.

### Review Checks

- Lineage and routing are correct. The design-review request is `FROM: s6-core.planner`, `TO: s6-core.implementer`, with `PARENT_DISPATCH_ID: s6-core-design` and `DESIGN_DOC_ID: s6-slice-6-design`; this review parents to that request. Exact-file lint of the request passed.
- The design object is current and docs-only. `git rev-parse HEAD` returned `a499bc31d6cc50d771cd0963353cf87b268881a0`; `git show --stat e13a6f4 a499bc3` shows r1 and r2 both edit only `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md`.
- The design states the locked constraints up front: byte-exact delivery enum, three seat verbs, I-PH families, honesty pins, no-perf fence, and OUT list at design lines 11-17. These match the ROADMAP/set constraints and do not silently amend the locked set.
- The codec decomposition is faithful. Lines 38-42 promote the existing canonical `address_list` path, delete the lineage comma splits, remove delivery silent-drop, and make decoded header TO/CC the projection truth. Source spot-checks: canonical parse exists at `internal/fieldspec/canonical.go:36-41,63-74`; current lineage split points are `internal/lineage/lineage.go:282-291,432-438`; current delivery silent-drop and `Envelope.To` projection are `internal/store/projections.go:110-141`.
- The A-1 digest design is faithful. Lines 46-49 promote existing `DigestExempt` to a conductor-supplied volatile class and preserve config/seat/phase/tier schema protection. This matches the current source gap at `internal/fieldspec/render.go:68-72,169-199` and the m-7 A-1 contract.
- Branch-A parenting is correctly decomposed. Lines 53-58 make `PARENT_DISPATCH_ID` system-only, move authoritative stamping to the shared submit handler, preserve optional `parent_hint`, and delete the anchoring-bounce class while leaving class-lineage gates intact. This matches m-1's locked branch-A semantics and does not over-claim fallback as validation.
- A-2 is at the right grain. Lines 62-67 keep monotonic ids, add segment-header high-water, replay existing outcomes, coalesce in-flight duplicates, and delete the old `journal.Append` id judge. Source spot-checks: `tables.T` already has `OutcomeByIntake` and dormant `ContentHash` at `internal/tables/tables.go:12-23,137-139`; current writer dedupes/re-enqueues at `internal/intake/writer.go:49-67,110-124`; current old path mints with `len(entries)+1` at `internal/intake/journal.go:70-83`.
- Section B projection/anchor coverage is complete. Lines 71-76 require serve-time accepted filtering, rebuild-path hygiene, audit view by `project` parameter, and accepted-state `WokenOn`. This covers the known source gaps in `internal/store/projections.go:21-46` and the mailbox-tail turn context.
- A-3 remint and B-1 lifecycle are inside the locked semantics. Lines 80-86 and 118-122 keep `seat_mint` as the generation pivot, keep `bound_now` runtime-only, derive active from first accepted governed submit per generation, and preserve no activation marker. Source spot-checks: current `seat.Manager.Mint` has the one-row guard at `internal/seat/binding.go:68-89`, `Resolve` follows the current binding table at `:92-100`, and channel active state already has the force-close/bound locus at `internal/channel/server.go:99-109,217-275`.
- Waiver/retraction is correctly seam-owned. Lines 90-95 route field typing to m-2 and effective-state semantics to m-1, retire the old `ORCH_REVIEW_WAIVER` row, preserve legacy unscoped behavior until retraction, and call out the non-enum render-absence expressibility issue instead of hiding it.
- F13 and D-2 are at the right layer. Lines 99-101 reduce `validateRecordKind` to layer-3 checks after registry membership/seat-scope; lines 112-114 put detail parity in `Outcome.Detail` and retire the shim re-render readback hack. Current gaps are still exactly there: `internal/engine/submit.go:143-178`, `internal/engine/loop.go:24-29,117-148`, and `cmd/frank-mcp/mcp.go:157-166,221-237`.
- The lock mechanism is a proper design decision, not an implementation handwave. Lines 105-108 cite both m-1 and m-7 halves, select `flock(2)` at phase minus one, refuse reads on loser, retain D5 residual, and fixture alias/race/kill-9 behavior. This satisfies the r2 grill agenda and avoids the pidfile trap rejected by m-1.
- The fixture table is complete enough for PLAN. Lines 150-163 cover GRILL_LOCK three, FX-A1a through FX-B1g, m-1 section E, m-1 section F.6 mapping, m-2 section 7, m-2 section 11.5, and the E2 floors. The ordering at lines 169-180 keeps D-2 before rejection-detail consumers, places A-2 before D-1 retry safety, and puts B-1/B-2/B-3 after seat generations.
- The folded `GRILL_LOCK` is real and bounded. Lines 184-220 record the six operator-resolved rows, rejected alternatives, no remaining operator-owned questions, and design-lock impact. It does not re-open the master-grilled branch-A parenting or the VP-fixed no-marker activation model.
- The ten dispatch constraints all have landing sites at lines 224-235, including promote-never-rebuild, F9 writer grain, rebuild-path projection pollution, still-bouncing F13 tokens, engine-side D-2, remint decomposition, two-leg F11 claim, full fixture map, claim pins/no-perf, and enum/verb/I-PH threat points.

### Non-Blocking Watchpoints

- Credential custody wording: design line 82 says the live-mint credential is delivered once in the operator submit reply, while m-7 A-3 says custody is unchanged and the credential is delivered only to the new seat's lane. I do not treat this as a blocker because the grill resolved the crash-window readback and the design keeps credentials out of records/projections/logs. PLAN should phrase this as an operator custody handoff and keep m-1/m-7 fidelity eyes on the exact delivery path before implementation.
- Waiver row render-absence: design lines 92 and 141 correctly flag that current FieldSpec grammar may not express operator-only rendering for non-enum waiver rows. PLAN must not claim fill-time absence is mechanically solved until the m-2 fidelity packet or an m-2-sanctioned minimal render rule lands; submit-path rejection alone is acceptable only if called out as the implemented floor.
- A-2 segment headers and GC: design lines 62 and 197 choose segment headers and line 67 adds the GC-drained-segments plus restart leg. Keep that fixture red-first; otherwise this is the exact place a future implementation could appear monotonic under normal rotation but regress after GC.

### Verification

- Incoming exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-design/DESIGN-planner-20260707-004949.md` -> OK.
- Own exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-design/DESIGN-REVIEW-implementer-20260707-005333.md` -> OK.
- Relay-root lint: `python3 ~/.claude/skills/tools/relay-lint.py --relay-root .relays/s6/s6-core-design .relays/s6/s6-core-design/DESIGN-REVIEW-implementer-20260707-005333.md` -> OK for the relay root and this file.
- Current head/status: `git rev-parse HEAD` -> `a499bc31d6cc50d771cd0963353cf87b268881a0`; `git status --short --branch` -> `## main`.
- Test smoke for unchanged code surface: `go test ./...` -> pass; 23 ok packages and 2 packages with no test files. Many packages were cached; `test/fixtures` ran in 22.503s.

ACTIONS_GIT_REF: review-only; no source/test/tracked-doc edits. Wrote `.relays/s6/s6-core-design/DESIGN-REVIEW-implementer-20260707-005333.md` and appended the matching `.relays/s6/INDEX.md` row; `.relays/` is gitignored relay transport.
FINAL_GIT_STATUS_SHORT: none - clean tree
Next requested action: s6-core.planner files the design-completion SITREP to s6.orchestrator-planner carrying this approving review and `GRILL_LOCK_ID: s6-grill-s6-core`, then holds for PROCEED-TO-PLAN.
