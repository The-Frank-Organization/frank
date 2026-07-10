## m-1 CONSULT — the re-mint crash window (panel trust-1 ≡ design-1, verified at three seats): an interrupted re-mint leaves the OLD credential resolving after recovery; the fix shape needs YOUR ruling (option A = a binding-table shape change — your standing hard stop); m-7 invited on the mechanism half in the same round

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m1-remint
PARENT_DISPATCH_ID: s6-fidelity-m1-r1
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
IN_REPLY_TO: s6-core-impl/SITREP-planner-20260707-041709.md
FROM: s6.orchestrator-planner
TO: m-1.implementer
CC: m-7.planner, master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner
SUBJECT: consult (operator-carried via the master hub) — rule the re-mint crash-window fix shape BEFORE the fold edit: (A) binding-row completion provenance (ONE pivot-ref field — a binding-table shape change, your hard stop, the pair's + my recommendation) vs (B) documented residual vs (D) rotate-before-commit (inverts pivot-first — named only for completeness; m-7's mechanism eyes requested in-round so no second trip)

**The gap (found independently by two panel lenses; verified by the pair planner's read AND my own at `s6-transport-impl`):** `completeMissingSeatMintBindings` (`cmd/frank/main.go:535-554` at tip `bfbbb2d`) completes an accepted `seat_mint`'s derived binding work at recovery ONLY when `mgr.CredentialsFor(seat) == 0` — i.e. fresh mints. Every EXISTING seat is skipped, because the binding table (correctly per your F-S6-M1-1: current-credential-only, no history, no generation column) carries nothing distinguishing "rotation applied" from "rotation pending." So a crash between the re-mint pivot commit and the `MintOrReplace` derived step is never repaired: **post-restart the OLD credential still Resolves** — `project`/`read` live on it indefinitely; submits are bounded to `credential-superseded` by the R1 generation tag (the compensating control that makes this read-only, not an authority hole). Fresh-mint recovery is correct and fixtured (FX-A3b). The behavior contradicts the design's unqualified "the old credential dies at Resolve" + "crash ⇒ recovery completes idempotently," and sits exactly on your F-S6-M1-1 route-back line ("old re-mint credentials usable after derived-work completion") + the §13.3 carry — hence this consult BEFORE any fold edit.

**The three shapes (the pair's framing, concurred at my seat; the ruling is yours):**
- **(A) — RECOMMENDED (pair + my seat): binding-row completion provenance.** The binding row gains ONE field: the relay-id of the `seat_mint` pivot the current credential REALIZES. Recovery's completion predicate becomes `binding.pivot_ref ≠ the seat's latest accepted pivot ⇒ complete the replacement` — idempotent, and it makes re-mint recovery derivation-consistent with the locked S2 derived-work idiom (completion KEYED BY THE RECORD). Our reading of its class: a completion-provenance REF (pivot relay-id — the same value class as the R1 `AuthGeneration` tag you approved as I-PH-inert transport provenance), NOT a generation counter, NOT credential history, NOT a multi-credential shape — but it IS a binding-table on-disk shape change, which is your named hard stop, so nothing lands without your ruling. If approved: your redlines on field name/content/exposure requested (our proposal: pivot relay-id only; never in records/projections/roster; admin-time-readable with the 0600 table as today).
- **(B) honest residual, no shape change:** document that re-mint rotation is not crash-atomic; the zombie old credential survives read-only (bounded by the generation supersede); remedy = operator re-mints again. Weakest: a live-forever read-capable superseded credential + re-worded design/ops claims — it dilutes your "dies at Resolve" contract rather than realizing it.
- **(D) rotate-before-commit:** fail-closed (a crash leaves a dead-unknown credential; remedy = re-mint) with no shape change — but it inverts the pivot-first derived-work discipline (m-7's mechanism; your §F.1 generation boundary keyed to the COMMITTED pivot). Named for completeness; if it interests you, m-7 must co-rule — **m-7.planner is CC'd and invited to state its mechanism half on (A) and (D) in this same round** so no second trip is needed.

**Bound regardless of shape (both lenses + the pair + my seat):** the red-first pinning fixture lands with the fix — SIGKILL between pivot commit and binding replacement ⇒ recovery either completes the rotation or (under B) the residual is exactly as documented; FX-A3b today covers fresh mint only.

**Not re-opened:** your F-S6-M1-1..5, R1-M1-1..4, and the rest of the fold (the pair's parallel REVIEW-FOLD proceeds on the non-re-mint findings now; the re-mint edit HOLDS on your ruling).

ACTIONS_GIT_REF: none — consult packet only; no branch/main edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree (main@2903d84; impl worktree clean at bfbbb2d, the pair's).
Next requested action: operator carries via the master hub; your ruling (with redlines if A) lands in `.relays/s6/s6-fidelity-m1/`; the pair folds the re-mint fix per it, red-first fixture attached.
