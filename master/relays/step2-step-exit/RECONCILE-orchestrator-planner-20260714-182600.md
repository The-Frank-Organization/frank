## RECONCILE — THE STEP-2 STEP-EXIT TEST PASSES on the merged spine `main@502e06c` (tag `s11-close`): all three roadmap legs green + live on the dogfood store · the INV-CATALOG red-battery demo fired (a planted fourth `delivery_state` turned the battery RED with `TestLawTerminalEnumByteExact` NAMING ITSELF + three interlocked laws, reverted to green, main never touched) · the uncached full battery green (26 ok / 0 FAIL / vet clean / ten laws). The Step-2 BUILD is complete (s7·s8·s10·s9·s11) and the master-owed executable exit is SATISFIED — the FULL LIVE RELAUNCH (real seats, honest E3/E4) remains the operator's SEPARATE checklist act, not a gate on this test. T5/T10 are acceptance-OPEN behind un-returned gates (g2/dc) — I RECOMMEND rescoping them as named carries, not holding Step-2 open, since they are additive surfaces beyond the minimal-comms floor the exit certifies. Requesting the operator's Step-2 close ratification + the VP's adversarial confirm.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step2-step-exit
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — this relay REQUESTS the operator's Step-2 CLOSE ratification (matching the Step-1 operator-granted close) + names the VP for the adversarial close-confirm; nothing flips the charter phase line to Step-3 until both land
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-merge-gate/MERGE-GATE-implementer-20260714-181240.md
FROM: master.orchestrator-planner
TO: operator, master.orchestrator-reviewer
CC: m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, s11.planner, s11.implementer
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the roadmap Step-2 exit test (:100-101) run live on the dogfood store — one FAILED observe-as-send (false "done" rejected pre-send, predicate named) + one PASSING send (conductor-captured evidence) + a parked lane waking exactly-once on the operator's validated reply — all green at `502e06c`; plus the INV-CATALOG red-battery demonstration and the uncached green battery; T5/T10 disposition + the Step-3 carry queue below; the close is the operator's to ratify

## A. VERDICT — the master-owed Step-2 step-exit test is SATISFIED
Run this session, uncached, on the merged spine `frank/ main@502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` (tag `s11-close`, obj `0462c24`, synced to private frank-dev). The three roadmap exit legs (`ROADMAP.md:100-101`), the INV-CATALOG red-battery demonstration, and the uncached green battery all pass. The Step-2 BUILD is complete across all five slices (s7 INV-CATALOG · s8 observe spine · s10 comms spine · s9 evidence thicken · s11 comms thicken).

## B. THE THREE LEGS — green, named, live on the dogfood store
Each leg is an executable proof on the current spine (run `-v`, uncached, this session; total 106s with the production leg at 104.87s on the real daemon socket):

1. **Leg 1 — a false "done" is rejected PRE-SEND, the predicate named.** `TestS8ExitGateFreshGenesisActivationAndDogfoodLegs` (exit-leg-1): the record commits `delivery_state: rejected`, `achieved_evidence: E0`, `record_integrity: observed`, the body names the unmet `done-predicate`, and **zero delivery intents** are emitted (nothing leaves the courier). Reinforced through the **real production daemon path** by `TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate` — a false-done relay over the live socket returns `rejected` and reads back the failing predicate.
2. **Leg 2 — a passing send carries conductor-captured evidence.** Same fresh-genesis test (exit-leg-2): `delivery_state: accepted`, `achieved_evidence: E1`, `target_gap_result: met`, `record_integrity: observed`, `attestation_source: conductor`, and the accepted record **projects to the recipient**. The production path adds the E2 suite leg (`achieved_evidence: E2` on a run-suite claim over the live daemon).
3. **Leg 3 — a parked lane wakes EXACTLY ONCE on the operator's validated reply.** `TestS10ExitLeg3FreshV8GateWakesExactlyOnceAfterLocalReobserve` on a fresh **v8 dogfood store**: gate → park (`GateParkedWaitingHuman`) → the operator's validated reply → local re-observe → **exactly-one** accepted wake (`resolves_gate` matched, `from: operator` → `to: seat-a`, mailbox count 1). The A-2 dedupe holds: a second wake does not fire.

## C. THE INV-CATALOG RED-BATTERY DEMONSTRATION — the constitution is a compile-time tripwire
In a **throwaway worktree** at `502e06c` (main never touched), I planted a single production violation — a fourth value `"deferred"` added to `delivery_state` in `internal/fieldspec/registry.json:71` — and ran the ten-law battery:
- **`TestLawTerminalEnumByteExact` went RED naming ITSELF**, printing the exact byte delta: `terminal enum bytes = ["accepted" "rejected" "held" "deferred"], want literal ["accepted" "rejected" "held"]`.
- **Three further laws caught the same perturbation** (`TestLawIntakeOutcomeOneToOne`, `TestLawPathHygiene`, `TestLawRebuildBeforeOpen`) — the catalog laws are interlocked; weakening the terminal surface trips more than its own guard.
- **Reverted** (worktree discarded) → the ten laws are green again on the clean spine; `main@502e06c` is byte-pristine (`git status` empty).

This is Cardinal rule 1 realized as landed behavior: **weakening a global law = a red battery that names the law.** The demonstration is honest end-to-end (planted → red-with-name → reverted → green), and no production or committed byte moved.

## D. THE UNCACHED GREEN BATTERY
`go clean -testcache && go test ./...` at `502e06c` (my own run this session, not the executor's report): **exit 0, 26 packages ok, 0 FAIL, `go vet ./...` clean, `test/invariants` (the ten laws) green.** Sequence-honest, file-captured; capture SHA-256 `ed78db5104534a498e4f3cc87cabc9f2cb9d420ccdb3904aeb010adc93575f4d`. This is the third independent green at the merged head (executor · s11.planner · master).

## E. SCOPE HONESTY — what this test IS and IS NOT
The master-owed **executable** step-exit is SATISFIED: the roadmap legs run live on the dogfood store (fresh-genesis + the real production daemon socket), the red-battery fires, the battery is green uncached. **The FULL LIVE RELAUNCH — a real team ON frank (Claude Code / Codex seats, the honest E3/E4 ladder, the lane-root pin · catalog bytes `943f07b…` · genesis from the `s7a-fieldspec-v5` bytes) — remains the operator's SEPARATE checklist act** (kickoff §2 dogfood relaunch / §3), and it is NOT a gate on this step-exit test. I flag it so the close is drawn at the right line: the step's *mechanism* is proven; the *live team dogfood* is a distinct operator election.

## F. T5 / T10 — acceptance-OPEN, and my disposition RECOMMENDATION
Both stayed gate-bound through s11 and their gates never returned:
- **T5 (elaborate-more fork)** holds on **g2** — m-5.implementer's OQ-2 `sensor`-ceiling adversarial review never completed to master.
- **T10 (re-prompt / claimless-`held`)** holds on **dc** — the m-3+m-6 design cell never returned.

Per **FINDING-4** they are acceptance-OPEN; s11 correctly built neither and claimed neither. **They are ADDITIVE comms surfaces beyond the minimal-comms floor the roadmap exit criteria certify** — the three legs (observe-as-send + evidence + park/wake) do not depend on the elaborate-more fork or the claimless-`held` re-prompt. **RECOMMENDATION: rescope T5/T10 as named carries** (a bounded gate-bound completion leg when g2/dc return, folded into early Step-3 or a small Step-2.x follow-on) **rather than hold Step-2 open** — the step goal (governance hardening + minimal comms) is delivered and demonstrated. The alternative — hold Step-2 open pending g2/dc — buys nothing the exit test needs and stalls the phase for two additive surfaces. **Operator's call; VP's to pressure-test.**

## G. THE STEP-3 CARRY QUEUE (named at close, not Step-2 failures)
- **T5 / T10** — gate-bound (g2/dc), per §F.
- **The item-2 soft-expiry arbiter design cell** — m-7 + m-3, post-Step-2; the executor-vs-fs_worker drain/termination divergence (join-with-SIGKILL vs breaker-trip-with-detached-worker) is its first question (`s11-build-escalate-fence/RECONCILE-…-143010`).
- **The s9 queue** — T7/T8 · C-1 (the one un-detached `git status` on the serialized loop) · R-1 (hardlink confinement, D5-class) · AO-1 (bounded saturation field when a real consumer exists) · **B1 scope_paths §F — the m-1 leg is STILL pending** (the sole governance-blocked item that predates this close).

## H. WHAT I REQUEST
1. **Operator:** ratify the **Step-2 CLOSE** on this step-exit PASS (the Step-1 precedent: the step closed on the operator's grant with the exit test passing). Rule the **T5/T10 disposition** (§F — recommend rescope).
2. **VP (`master.orchestrator-reviewer`):** the adversarial close-confirm — pressure-test the step-exit evidence, the scope line in §E (mechanism-proven vs live-relaunch-deferred), and the T5/T10 rescope. A REVISE returns it to me; a concurrence + the operator ratification CLOSES Step-2 and flips the charter phase line to Step-3.

On both, I run the Step-2 close fold (charter rule-1 phase line → Step-3; the carry queue seeded into Step-3 planning; the dashboard/RECONCILE/ROADMAP already carry the step-exit PASS, marked pending-ratification until then).

## Verification
- Legs: `go test ./test/fixtures/ -run 'TestS8ExitGateFreshGenesisActivationAndDogfoodLegs|TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate|TestS10ExitLeg3FreshV8GateWakesExactlyOnceAfterLocalReobserve' -v` — 3/3 PASS at `502e06c` this session (production leg 104.87s live).
- Red-battery: throwaway `git worktree` at `502e06c`; planted `delivery_state` fourth value; `go test ./test/invariants -run TestLaw -v` → `TestLawTerminalEnumByteExact` FAIL naming itself + the byte delta (+3 interlocked); worktree removed; re-run green; `main` `git status` empty.
- Battery: `go clean -testcache && go test ./...` → exit 0 / 26 ok / 0 FAIL / vet clean at `502e06c`; capture SHA-256 `ed78db51…f4d`.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step2-step-exit/RECONCILE-orchestrator-planner-20260714-182600.md` — run below.

ACTIONS_GIT_REF: none — no `frank/` code edit, no commit, no push by this relay; test executions were read-only, and the red-battery demo used a throwaway `git worktree` (added at `502e06c`, removed after) leaving `main` byte-pristine. Governance docs updated in cwd (not a git repo): RECONCILE §s11 + this step-exit record · ROADMAP Step-2 marker · README dashboard. Disk artifacts: this relay + one INDEX.md row timestamped 20260714-182600.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`, synced to frank-dev); the throwaway red-battery worktree removed, `git worktree list` clean of it.
Next requested action: the operator ratifies the Step-2 close + rules the T5/T10 disposition; the VP returns the adversarial close-confirm. On both, master runs the Step-2 close fold (charter phase line → Step-3) and seeds the carry queue into Step-3 planning. The live team-on-frank relaunch (E3/E4) stays the operator's separate act, un-gated by this close.
