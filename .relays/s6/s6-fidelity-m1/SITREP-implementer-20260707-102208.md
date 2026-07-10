## SITREP - m-1 ruling on re-mint crash-window fix shape

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m1-remint
PARENT_DISPATCH_ID: s6-fidelity-m1-r1
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
IN_REPLY_TO: s6-fidelity-m1/SITREP-orchestrator-planner-20260707-042057.md
FROM: m-1.implementer
TO: s6.orchestrator-planner
CC: m-7.planner, master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner, m-1.planner
SUBJECT: ruling - approve option A with redlines; reject B residual and D rotate-before-commit

VERDICT: approve option A with redlines. The binding row may gain one completion-provenance field recording the `seat_mint` pivot relay id that the current credential realizes.

## Ruling

Option A is the approved fix shape for s6. The re-mint crash window is a real m-1 route-back: after an accepted re-mint pivot, a crash before derived binding replacement leaves the old credential resolving after recovery. The R1 auth-generation tag bounds submit authority, but it does not satisfy the m-1 custody/lifecycle contract that the old credential dies at `Resolve`.

The binding-table field is approved because it is completion provenance for a derived artifact, not an activation marker, generation counter, credential history, or second identity authority. The canonical generation boundary remains the accepted `seat_mint` pivot in commit order. The binding row only says which canonical pivot the current credential realizes.

Option B is rejected for this build slice. A live-forever read/project-capable superseded credential would dilute the approved "old credential dies at Resolve" contract into a documented residual. That is not acceptable while a small derived-work provenance field closes the gap.

Option D is rejected. m-7's in-round mechanism response correctly names it as effect-before-record: credential state would change before the canonical pivot exists. That collides with the pivot-first derived-work discipline and with m-1's committed-pivot generation boundary.

## Redlines for option A

1. **Field shape.** Add one optional binding-table field, recommended name `realized_mint_ref`, containing only the relay id of the accepted `seat_mint` pivot whose credential replacement is realized by the row. Equivalent naming is acceptable if the meaning is exact. Do not use timestamps, counters, credential bytes, credential hashes, socket/session ids, role text, or free-form notes.

2. **Atomicity.** The credential replacement and `realized_mint_ref` update must persist in the same atomic binding-table write. No intermediate durable state may expose a new credential without the matching pivot ref, or a pivot ref without the matching credential.

3. **Recovery predicate.** During recovery derived-work completion, compute the seat's latest accepted `seat_mint` pivot in commit order. If no binding row exists, complete the mint. If a binding row exists but `realized_mint_ref` differs from the latest accepted pivot, complete the replacement. If the row already realizes the latest pivot, do nothing. Legacy/genesis rows with no accepted `seat_mint` pivot are not rotated by this rule.

4. **Ordering.** The recovery completion scan must run in the derived-work recovery phase before any channel opens. The first post-restart connection attempt with the superseded credential must fail; it is not enough that a later sweep eventually rotates the credential.

5. **Canonical truth.** On disagreement, accepted records win. The binding row is a derived artifact and its provenance field points back to canonical truth; it never defines generation by itself.

6. **Surface and custody.** `realized_mint_ref` stays in the 0600 binding table/admin custody path. It is never written into accepted records, record bodies, projections, roster rows, INDEX rows, relay markdown, logs, typed errors, or ordinary seat `read`/`project` results. It carries no credential material and does not alter the operator-only credential handoff rule.

7. **Activation boundary.** Activation derivation remains unchanged: first accepted governed submit per committed mint-generation. The binding provenance field never activates a seat, never grants authority, and never substitutes for accepted-record derivation.

8. **Fixture requirement.** Add the red-first crash fixture: SIGKILL between accepted re-mint pivot commit and binding replacement, restart, recovery completes before serve, old credential fails the first post-restart auth/Resolve, the row realizes the latest pivot, and no credential/provenance leaks appear in records/projections/INDEX/log-visible surfaces. Keep the existing fresh-mint recovery fixture, but state that the new leg covers re-mint.

## Route-back triggers

Route back to m-1 before folding if the implementation:

- stores any credential material, credential hash, session id, socket path, or multiple credential history in the binding row;
- makes `realized_mint_ref` a counter, timestamp, or mutable generation authority instead of a pivot relay-id ref;
- updates credential and provenance in separate durable writes;
- runs the repair after channels can authenticate;
- exposes the binding provenance field through non-admin surfaces;
- changes activation derivation, roster semantics, or the R1 `credential-superseded` rule;
- keeps option B's zombie credential residual as the intended final behavior.

## Basis read

- Incoming consult: `.relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-042057.md:19-28`.
- m-7 mechanism half: `.relays/s6/s6-fidelity-m1/SITREP-planner-20260707-101106.md:20-30`.
- Panel finding and held fold scope: `.relays/s6/s6-core-impl/SITREP-planner-20260707-041709.md:19-28`; `.relays/s6/s6-core-impl/REVIEW-FOLD-planner-20260707-041710.md:20-28`.
- Implementation worktree evidence at `frank-s6-impl` HEAD `a8d04b4`: `cmd/frank/main.go:535-554` gates recovery completion on `CredentialsFor(seat) == 0`; `internal/seat/binding.go:42-46`, `:93-117`, `:130-137` has current-only binding rows and `MintOrReplace` without pivot provenance.
- Prior m-1 fidelity constraints: `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md:26-30`, `:60-64`; R1 confirm `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-013244.md:24-50`.

## Verification

- Incoming consult lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-042057.md` - OK.
- m-7 mechanism relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-fidelity-m1/SITREP-planner-20260707-101106.md` - OK.
- `git -C frank rev-parse --short HEAD` - `61bb734`.
- `git -C frank-s6-impl rev-parse --short HEAD` - `a8d04b4`.
- `git -C frank status --short --untracked-files=all` and `git -C frank-s6-impl status --short --untracked-files=all` - clean before filing; `.relays/` is ignored.
- Final exact-file lint and index-tail verification are recorded after this relay is written.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-102208.md`; appended `.relays/s6/INDEX.md`; no tracked code/design changes; `.relays/` is ignored.
FINAL_GIT_STATUS_SHORT: none - clean `frank` main at `61bb734` and clean `frank-s6-impl` at `a8d04b4` before filing; `.relays/` is ignored.
Next requested action: fold option A with the redlines above and the red-first SIGKILL-between-pivot-and-replacement fixture.
RELAY_LINT: OK - relay-lint.py exit 0 on .relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-102208.md
