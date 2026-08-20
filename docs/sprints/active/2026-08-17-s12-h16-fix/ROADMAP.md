# s12 — the H-16/H-26 fix-lane build pair (sprint doc root)

RUN_ID: `s12` · Commissioned by `master/relays/step3-h16-h26-lane/PLAN-orchestrator-planner-20260817-184653.md` under charter `CH-s12-h16-fix` (`master/subteams/s12-h16-fix/CHARTER.md` @ sha256 `e79c2cf03b66048bd0120c4b8d19edae6e2259b6bddd8d45edf50aac3e0d24d7`).

## Scope of record

The four-leg enumeration in `master/relays/step3-h16-h26-lane/RECONCILE-orchestrator-planner-20260817-184510.md` §3:
H-16 (rev20 realized within its Bounds) · H-26 (`-mint` onto the shared `AcquireRoot` helper) · H-16-REG (three fieldspec rows + `record_kind` tokens + marker bump) · R-INIT-UNLOCKED (`-init` under the same lock, three regression legs) — plus the rev20-bound cross-cutting obligations (presence-pinned validator, 48-executed-case T-R6F3 unit, 31-row shared-impact regression).

Spec of record (FROZEN; execute, never re-design): `master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` **rev21** @ sha256 `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05` (rev20 `e09fab09…` superseded 2026-08-17 via the void rail: the one-byte-run §4a `mint-predecessor-mismatch` naming, pair-approved `…202647`). Ruled realization authority for the registry surface: m-2's `master/relays/step3-h16-h26-lane/DESIGN-planner-m2-20260817-201520.md` §1.

## Declarations

- **Relay substrate:** the relay root is `frank/.relays/s12/` per charter CH-s12-h16-fix LAYOUT (a recorded divergence from this sprint root's own `.relays/`; the INDEX `root:` marker records it). TRACKED and banked on `main` by MASTER'S CHECKPOINT CADENCE (observed at workspace commits `2cb9a0c`/`414ee38` — master's act, not s12's; s12 itself commits nothing to `main`). The work branch `s12-h16-fix` carries CODE ONLY; merge is operator-only via the terminal MERGE-GATE.
- **Battery captures:** RED/GREEN battery outputs are FILE-captured, sequence-honest, under `frank/.relays/s12/batteries/` (s9/s11 precedent; declared here as the one additional directory), written by absolute path into the shared tree.
- **Branch/isolation:** all s12 code work on branch `s12-h16-fix` off `main`, hosted in an ISOLATED WORKTREE created by the Implementer at Task 0 via `superpowers:using-git-worktrees` (the shared workspace tree stays on `main`; the pair Planner creates no implementation branch).
- Standard directories are instantiated on first use; `plans/` opens with `PL-s12-build-plan-20260817.md`.

## Status

- 2026-08-17: sprint root + relay root opened by s12.planner; PLAN authored and routed to s12.implementer for adversarial plan-review (`s12-build-plan`).
- 2026-08-17: plan-review r1 = must-revise (five findings); r2 folded F1/F3/F5 and routed F2/F4 UP (`s12-build/SITREP-planner-20260817-194616.md`); the r1 structural-error waiver asked of the operator (`…194914`).
- 2026-08-17: both rulings landed (master `…202541`): the record_kind realization = m-2 `201520` §1 verbatim; the Task-10 byte = `mint-predecessor-mismatch`; the spec rebound to rev21 `cc8bcff3…` (pair-approved). Plan r3 folded and routed for re-review (`s12-build-plan-3`). Open gates before `DISPATCH IMPL`: the implementer's approve + the operator's r1 waiver disposition.
- 2026-08-17: plan-review r3 = must-revise on four accuracy/coverage findings (the ruled realization and rev21 bind PASSED); r4 folded them (governance-truth statement, the pinned relay-lint expected-output oracle, the one substrate model, Task 9's operator-only authority negatives) and routed for re-review (`s12-build-plan-4`). The rev21 owner/join tail (m-1→m-2 reruns + refreshed join) runs in parallel and does not gate the build tasks.
