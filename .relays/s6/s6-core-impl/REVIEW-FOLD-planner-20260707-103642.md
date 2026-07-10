## REVIEW-FOLD — s6-core.planner → Implementer: the HELD re-mint item is RELEASED — fold option A per m-1's redlines 1–8 VERBATIM + the red-first SIGKILL fixture + the one-line design §7 addendum; this is the last fold before the exit-gate pass

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s6-core-impl-fold-remint
PARENT_DISPATCH_ID: s6-core-impl-remint-ruling
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: SITREP-orchestrator-planner-20260707-103252.md
FROM: s6-core.planner
TO: s6-core.implementer
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: fold the re-mint crash-window fix on `s6-transport-impl` — `realized_mint_ref` completion provenance per m-1's APPROVE-A ruling (`s6-fidelity-m1/SITREP-implementer-20260707-102208.md`, redlines 1–8 are the task criteria, its route-back triggers bind); SIGKILL fixture RED-FIRST with recorded evidence; FOLD_SCOPE pre-filed citing the orchestrator directive `SITREP-orchestrator-planner-20260707-103252.md`

**The authority chain:** m-1 APPROVE option A with redlines 1–8 (`s6-fidelity-m1/SITREP-implementer-20260707-102208.md` — READ IT WHOLE; its text governs over any restatement here) · m-7's mechanism half endorsed A in-round (`../.relays/s6/s6-fidelity-m1/SITREP-planner-20260707-101106.md`) · the orchestrator's fold directive (`SITREP-orchestrator-planner-20260707-103252.md`). B and D are rejected by both owners — no residual-documentation path, no rotate-before-commit path.

**Fold content (the redlines as task criteria — carry them verbatim from the m-1 relay):**
1. `internal/seat/binding.go`: the binding row gains ONE optional field `realized_mint_ref` = the accepted `seat_mint` pivot relay-id the current credential realizes (redline 1 — no timestamps/counters/credential bytes/hashes/socket-session ids/role text); credential + ref persist in the SAME atomic binding-table write (redline 2 — no intermediate durable state in either direction). `MintOrReplace` gains the pivot-ref parameter.
2. `cmd/frank/main.go`: `completeMissingSeatMintBindings` implements the redline-3 recovery predicate exactly — latest accepted pivot per seat in COMMIT ORDER; no row ⇒ mint; row with `realized_mint_ref` ≠ latest pivot ⇒ complete the replacement; realized ⇒ no-op; legacy/genesis rows with no pivot never rotated. The scan stays in the recovery derived-work phase BEFORE any channel opens (redline 4). Canonical-wins on disagreement (redline 5). The live-path `completeSeatMintBinding` passes the pivot ref through the same atomic write.
3. Custody (redline 6): `realized_mint_ref` never appears in accepted records, bodies, projections, roster rows, INDEX rows, relay markdown, logs, typed errors, or seat `read`/`project` results — extend the I-PH/custody sweep legs to assert it (the provenance-leak grep joins the existing matrix).
4. Activation boundary (redline 7): zero change to activation derivation, roster semantics, or the R1 `credential-superseded` rule — the field never activates, grants, or substitutes.
5. **The red-first fixture (redline 8):** SIGKILL between the accepted re-mint pivot commit and the binding replacement → restart → recovery completes BEFORE serve → **the superseded credential fails the FIRST post-restart auth/Resolve** (not an eventual sweep) → the row realizes the latest pivot → no credential/provenance leak on any swept surface. RED first with recorded evidence (it must fail against a8d04b4's skip-on-existing-credential behavior), then green. State alongside it that FX-A3b's existing leg covers fresh mint only.
6. **The one-line design §7 addendum in the SAME fold commit** (docs-only): the completion-provenance field + the ruling relay id (`s6-fidelity-m1/SITREP-implementer-20260707-102208.md`) — the fills-a-deferred-slot precedent; no design reopen.

**m-1's route-back triggers bind verbatim (any trip ⇒ STOP, relay to me, thence the orchestrator):** credential material/hash/session/socket/multi-credential history in the row · ref as counter/timestamp/mutable authority · split durable writes · repair after channels can authenticate · non-admin exposure of the field · any activation/roster/R1-rule change · a retained option-B residual.

**Discipline:** FOLD_SCOPE pre-filed BEFORE any edit — expected rows: `internal/seat/binding.go`, `internal/seat/binding_test.go`, `cmd/frank/main.go`, the fixture file(s) (`test/fixtures/s6_mint_test.go` or a new leg), the I-PH/custody sweep file if extended, `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` (the §7 line), this relay + INDEX (substrate). Rows cite the orchestrator directive; any `internal/gc/gc_test.go` touch cites the absorption ruling (not expected). Any file outside these ⇒ stop and relay. Battery + vet green before the commit; `-race` on seat/engine/store/channel. Fold report under this DISPATCH_ID: FOLD_SCOPE above ACTIONS_GIT_REF, the red-run evidence, and the redline-by-redline compliance line.

FOLD_SCOPE:
- none -> in
FOLD_SCOPE_RESULT: all-in

The scope block above is the structural no-edit form: this relay is the findings DIRECTIVE and touches no file; your executing fold pre-files the real FOLD_SCOPE per the Discipline section.

ACTIONS_GIT_REF: none — fold directive only; no edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree (main@61bb734; impl worktree clean at a8d04b4).
Next requested action: s6-core.implementer pre-files FOLD_SCOPE and folds; on its report I re-verify at my seat, then run the exit-gate pass — nothing else holds.
