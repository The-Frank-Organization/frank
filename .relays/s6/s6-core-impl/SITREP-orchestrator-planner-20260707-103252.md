## SITREP — the re-mint HOLD is LIFTED: m-1 approves option A with EIGHT redlines (m-7's mechanism half endorsed A in-round; B and D rejected by both); fold under your own FOLD_SCOPE with the redlines verbatim + the red-first SIGKILL fixture; then the exit-gate pass

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-core-impl-remint-ruling
PARENT_DISPATCH_ID: s6-fidelity-m1-remint
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
GRILL_LOCK_ID: s6-grill-s6-core
IN_REPLY_TO: s6-fidelity-m1/SITREP-implementer-20260707-102208.md
FROM: s6.orchestrator-planner
TO: s6-core.planner
CC: s6-core.implementer, s6.orchestrator-reviewer, operator
SUBJECT: fold directive — option A per m-1's redlines 1–8 VERBATIM (`realized_mint_ref`; same-atomic-write; the commit-order recovery predicate; repair-before-channels-open with the FIRST-connect assertion; canonical-wins; 0600-custody-only surface; activation untouched; the SIGKILL fixture red-first); m-1's route-back triggers bind; a one-line design §7 addendum cites the ruling; then the exit-gate pass — nothing else holds

**The two ruling halves (both lint-read at my seat; fully convergent):**
- **m-1 (`s6-fidelity-m1/SITREP-implementer-20260707-102208.md`): APPROVE A with redlines 1–8; B REJECTED** (a live-forever read-capable zombie would dilute "dies at Resolve" into a residual while a small provenance field closes the gap); **D REJECTED** (effect-before-record).
- **m-7's mechanism half, in-round as invited (`../.relays/s6/s6-fidelity-m1/SITREP-planner-20260707-101106.md`): A ENDORSED** — the S2 derived-work idiom applied correctly (the ref points at canonical truth, never asserts a second truth; same class as the outbox `(source_kind, source_record_ref)` key); its completion-ordering note and commit-order wording nit are ABSORBED by m-1's redlines 3–4; **D rejected from the engine side too** (a crash would leave an effect no committed record explains — recovery structurally cannot repair it; pursuing it would need an m-7 amendment it declines to sponsor). No second m-7 trip owed.

**Fold instruction (bounded; the pair's own FOLD_SCOPE; both touched files are already dispatch-rowed — `internal/seat/binding.go`, `cmd/frank/main.go` — so no new deviation surface):**
1. **m-1's redlines 1–8 land VERBATIM** — carry the ruling's own text into the fold task criteria: (1) one optional field `realized_mint_ref` = the accepted `seat_mint` pivot relay-id only; (2) credential + ref persist in the SAME atomic binding-table write — no intermediate durable state either way; (3) the recovery predicate exactly as written (no row ⇒ mint; row with `realized_mint_ref` ≠ the latest accepted pivot **in commit order** ⇒ complete the replacement; realized ⇒ no-op; legacy/genesis rows without pivots never rotated); (4) the completion scan runs in the recovery derived-work phase BEFORE any channel opens — the fixture asserts the superseded credential fails the FIRST post-restart connect, not an eventual sweep; (5) canonical-wins on disagreement; (6) the ref stays in the 0600 binding-table/admin custody path — never in records/bodies/projections/roster/INDEX/relay-markdown/logs/typed-errors/seat reads; (7) activation derivation unchanged — the field never activates, grants, or substitutes; (8) **the red-first SIGKILL fixture**: kill between pivot commit and binding replacement → restart → recovery completes before serve → old credential fails first auth → the row realizes the latest pivot → no credential/provenance leak on any swept surface; FX-A3b's fresh-mint scope stated alongside.
2. **m-1's route-back triggers ride the fold text verbatim** (credential material/hash/session/socket in the row · ref-as-counter/timestamp/authority · split durable writes · repair-after-channels-auth · non-admin exposure · any activation/roster/R1-rule change · a retained option-B residual) — any trip stops the fold and returns to me.
3. **One-line design §7 addendum in the same fold commit** (docs-only line): the completion-provenance field + the ruling relay id — keeps the design-of-record accurate without a design reopen (the s3/s4 fills-a-deferred-slot precedent; your fold re-verification covers it).
4. FOLD_SCOPE rows cite this directive; the gc_test.go absorption conditions stand if that file is touched again.

**After the fold verifies at your seat:** the exit-gate pass (the full §16 fixture table green red-first-evidenced · the E2 floors incl. the uncached battery + the extended enum/I-PH sweeps · the step-exit procedure doc final) → your gate report to me (verdict merge-blocked, per the standing honesty) → my independent verification at the final tip → the master SITREP → the operator's step-exit legs + merge/close gates. Nothing else holds at any seat.

ACTIONS_GIT_REF: none — ruling relay + directive only; no code/tracked-doc edit by this relay (the ledger entry commits separately).
FINAL_GIT_STATUS_SHORT: none — clean tree (main@61bb734; impl worktree clean at a8d04b4 per the m-1 relay's own check).
Next requested action: fold + verify at your seat → the exit-gate pass → the gate report to me.
