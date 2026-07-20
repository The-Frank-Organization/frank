## MERGE-GATE — the s9 MERGE-DECISION relay TO the operator: RECOMMEND GRANT — the slice is built, reviewed (code approve, zero blockers), folded, fence-reconciled under master's grant, and battery-green; TWO named preconditions remain open (the m-7 landing confirm · the m-3 owner-fidelity + AO-1 byte) — grant on their clean return; a contradiction STOPS; no merge is executed by this relay

ROLE: Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s9-merge-decision
PARENT_DISPATCH_ID: s9-build-slice-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — this relay REQUESTS the operator's HUMAN_MERGE_AUTHORIZATION (condition e); nothing merges without it
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev13
IN_REPLY_TO: master/relays/s9-build-escalate-fence/PLAN-orchestrator-planner-20260714-013011.md
FROM: s9.planner
TO: operator
CC: master.orchestrator-planner, master.orchestrator-reviewer, s9.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-2.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: what merges = `s9-evidence-thicken` `39474d0 → d91fcfb` (9 linear commits, fast-forward; +2,288/−308 incl. the committed evidence package); the two open preconditions are STOP-ON-CONTRADICTION and the AO-1 answer forks (carry → grant as-is; deliver → one bounded in-fence task + my re-verify first); push/publication stays your separate move

**RECOMMENDATION: grant `HUMAN_MERGE_AUTHORIZATION` for the s9 slice once the two named preconditions return clean.** Everything the pair and master control is closed; the two open items are owner returns already requested and CC'd to you.

### What merges

`frank/` branch `s9-evidence-thicken`, base `main@39474d0` (`s10-close`), head `d91fcfb340b029c39c8493084ce2f227409aa546` — nine linear commits (fast-forward-able): T1 `ba26b27` (shared FS worker) · T2 `5f6a7ec` (find-references) · T5 `8a2b73d` (attestation negative) · T3 `0f1aa42` (§6.1 binding pass on the rev13 guide byte) · T4 `e6a80d6` (m-3 semantics: `none`-only opaque-accept + total table + probe removal) · T4 `db9a166` (m-7 owner bytes E1–E10 verbatim, isolated) · T4 `1b87261` (FX-VCS + dogfood fixtures) · T9 `76179ec` (exit-fixture set + ⑤ ODB egress pair) · MF-1 `d91fcfb` (evidence-table fold, one file). T6 was verification-only by design (no commit).

### The evidence chain of record (each link a lint-clean relay or captured battery)

r1-approved pair PLAN (`s9-build-plan`, adoption sha-pinned) → the delegated token (`s9-build-token`, all-in fence, correct lineage, void-token excluded) → the T3 blocker correctly HELD uncommitted on a real plan-table gap → m-3's rev13 guide byte + m-3.implementer's bounded re-review APPROVE (`s9-baserefusal-m3/DESIGN-REVIEW-implementer-20260714-011500`) → T4 on verbatim owner bytes → my ONE B11 end-of-slice adversarial review (`s9-build-slice-review/REVIEW-FOLD-planner-20260714-013000`): **code approve, zero blockers/must-fix** across my own verification + two read-only reviewer lenses (confinement/TOCTOU, breaker, binding-pass bypass, redaction, opaque-accept reachability, incomplete-scan honesty — all sound; terminal enum byte-exact; marker sweep zero) → the MF-1 evidence fold (`d91fcfb`, one file, pre-edit scope artifact, post-fold full battery green 137.972s) → **master's fence-row grant** (`s9-build-escalate-fence/PLAN-orchestrator-planner-20260714-013011`, `FENCE_EXPANSION_AUTHORIZED: granted`). My own verification runs: full uncached `go test ./... -count=1` green at `76179ec` and (implementer-run, captured) post-fold at `d91fcfb`; owner-byte isolation diffed byte-faithful; `executor.go`/`fieldspec` zero-diff.

### The fence reconciliation (master's directed citation — the license record)

All 29 changed paths reconcile to the token block plus master's two granted rows: **`test/invariants/store_recovery_test.go`** — granted scope-bound to the E10 descriptor lines only (tripwire satisfied: no `TestLaw*`/catalog/census byte moved; the ten INV-CATALOG laws green) — and **`internal/config/lane_vcs_test.go`** — the FX-VCS legs granted in place. The committed evidence table records both as-found (`OUT — escalated`), honestly; this relay + master's ruling are the grant record. Nothing else in the diff is outside the block.

### The two OPEN preconditions (grant on clean return; a contradiction STOPS everything)

1. **m-7 landing confirm** (`s9-build-fidelity-m7/SITREP-planner-20260714-013002`, out): E1–E10 + FX-VCS landed faithfully.
2. **m-3 owner-fidelity + the AO-1 byte** (`s9-build-fidelity-m3/SITREP-planner-20260714-013003`, out). **The AO-1 fork:** if m-3 answers **carry**, grant as-is; if **deliver** (surface the bounded find-references count/saturation), ONE bounded in-fence task + my re-verification lands BEFORE your grant — I will relay the updated head.

### Dispositions at close (the dogfood-evaluation section, condition g — one series)

- **T7/T8:** never opened — their m-2/m-1 co-signs did not arrive in-slice; they exit as **report-and-hold carries** to master's queue, not s9 failures; no code shadow exists.
- **Blocked ledger intact:** B1 (scope_paths — the §F m-1 leg STILL pending) · B-diff · B3 · B2 · B4 — all governance-only, none stubbed.
- **The catch ledger, s9's entries:** (1) the mis-seated staffing token — linted clean, caught only by the operator's read (the standing "gates don't check staffing semantics" limit); (2) the T3 totality gap — caught TYPED and PRE-COMMIT by the fail-closed tuple collapse + the wide regression battery (the machinery catching its own plan's false claim); (3) **the fence anatomy (master's ruling §3, shared three ways):** the mid-build `lane_vcs` activation introduced test seams no fence had named — activation ruling didn't enumerate them, my routing asserted "in-fence" without the row check, the build's table licensed at task grain — **caught by the mechanical diff→block table pre-merge, exactly as designed**. Master's discipline refinement is on the record: *a mid-build amendment activation must reconcile its test/fixture seams into the fence, and an owner return naming un-fenced loci is an escalation trigger, not an in-fence assertion.*
- **Named carries + stated residuals:** C-1 (the pre-existing `git status` exec on the serialized loop — the one un-detached path post-T1; routed to the m-3 ledger, m-7 CC'd) · R-1 (hardlink confinement residual, D5-class, label locus = m-3's byte) · AO-1 (per the fork above).

### What executes on your grant — and what does not

On your `HUMAN_MERGE_AUTHORIZATION` (a relay under this DISPATCH_ID, or your direct instruction): fast-forward `frank/ main` to the slice head and plant tag `s9-close`. **NOT executed by this relay: no merge, no tag, no push.** Publication (push) stays your separate move under the standing publication plan, as at s10-close.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/s9-merge-decision/MERGE-GATE-planner-20260714-014500.md` — run before handoff.
- This turn's consumption checks: master's grant read in full (tripwire re-checked against the E10 diff — descriptor lines only); the fold commit `d91fcfb` verified one-file with the 29/29 path-total table and both OUT rows preserved as-found.

ACTIONS_GIT_REF: none — a merge-decision recommendation; no merge, tag, push, or `frank/` edit is executed by this relay. Artifacts this turn: this relay + one `master/relays/INDEX.md` row timestamped 20260714-014500.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `frank/` main clean at `39474d0` (`s10-close`); the `s9-evidence-thicken` worktree clean at `d91fcfb`.
Next requested action: operator holds the grant until the m-7 and m-3 returns land clean (both CC you); on AO-1 "deliver" I relay the updated head after the bounded task + re-verify; on your grant the fast-forward + `s9-close` tag execute and master runs its step-integration verification. No merge without `HUMAN_MERGE_AUTHORIZATION`.
