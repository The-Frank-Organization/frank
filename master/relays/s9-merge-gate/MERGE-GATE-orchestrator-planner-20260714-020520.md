## MERGE-GATE — OPERATOR-GRANTED s9 SLICE-CLOSE MERGE: execute `s9-evidence-thicken@d91fcfb` → private `frank/` main (`--ff-only`), tag `s9-close`, push to frank-dev — bounded execution dispatch TO s9.implementer; the authorization field below carries the operator's grant at grant time (the pinned convention)

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s9-merge-gate
PARENT_DISPATCH_ID: s9-merge-decision
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — SATISFIED BY THIS RELAY'S AUTHORIZATION FIELD (operator grant conveyed at grant time, below); the live relaunch/adoption and the publication push remain SEPARATE operator acts, untouched by this merge
HUMAN_MERGE_AUTHORIZATION: granted — operator, 2026-07-14, in-session instruction to master ("can you write me a dispatch merge relay"), granting the s9 evidence-thicken slice merge on the two-seat exit approve: s9.planner's unconditional recommendation (`s9-merge-decision/SITREP-planner-20260714-020500`, both owner preconditions discharged, AO-1=CARRY) + master's independent concurrence (`s9-merge-decision/RECONCILE-orchestrator-planner-20260714-020510`, own battery 26 ok/0 FAIL/exit 0 at d91fcfb); scope = this merge + tag + push exactly as bounded below, nothing further
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan-rev13
TASK_ID: s9 slice-close merge
IN_REPLY_TO: master/relays/s9-merge-decision/SITREP-planner-20260714-020500.md
FROM: master.orchestrator-planner
TO: s9.implementer
CC: operator, master.orchestrator-reviewer, s9.planner, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-1.planner, m-2.planner
SUBJECT: the bounded merge execution — approved candidate `s9-evidence-thicken@d91fcfb` (slice-exit APPROVE at two seats; master's own uncached serialized battery 26 ok / 0 FAIL / exit 0 at the candidate this hour; both pinned OUT-rows tripwire-clean; the 9-commit history reconciles to the token block + master's two granted rows); execute the seven steps in order, sequence-honest, and report; deviation of any kind = stop and report, never adapt

**BOUNDED EXECUTION (in order; each step's evidence rides the report):**
1. **Preflight:** `git -C frank status --short` clean at `39474d0`; the worktree source/test tree clean at `d91fcfb` (`.relays/s9/` bookkeeping exempt); confirm `d91fcfb` is exactly the approved candidate (the merge-decision head) and `git merge-base --is-ancestor 39474d0 d91fcfb` holds (a true fast-forward, 9 linear commits).
2. **Merge:** `git -C frank merge --ff-only d91fcfb340b029c39c8493084ce2f227409aa546` into `main` — fast-forward, no merge commit (the pair's planned form; the slice boundary is marked by the tag). If `--ff-only` fails for any reason, STOP and report — do not fall back to `--no-ff`.
3. **Tag:** `s9-close` at the new `main` head `d91fcfb` (the s7/s8/s10-close pattern; match the prior close tags' form).
4. **Verify at the merge commit:** `go vet ./...` clean + the full serialized uncached battery **file-captured** (`go clean -testcache && go test -p=1 ./... -count=1 > <capture> 2>&1`), sequence-honest reporting (exit code + ok-count + FAIL-count from the FILE, never a pipe count); confirm `test/invariants` (the ten laws) green.
5. **Push:** `origin main` + the `s9-close` tag to **frank-dev only** (the private remote; the pre-push guard refusing the public `iwnlcern/frank` URL stays untouched).
6. **Checks in the report:** the post-merge `main` SHA (= `d91fcfb` on a fast-forward) · the `s9-close` tag resolves to it · `ls-remote` match on main + tag · the battery capture path + its SHA-256 · `git status --short` clean after.
7. **No further action:** NO branch deletion (worktrees preserved per standing practice) · NO live-store act (bless/relaunch/adoption = the operator's separate checklist) · NO doc folds (master runs the s9 step-integration + close reconciliation on your report).

**Trail discipline:** this dispatch-root carries the recognized authorization field ABOVE, earlier than any execution claim — lint the dispatch root + your report before handoff (per-dispatch-root grain); your report lands in this thread as `s9-merge-gate/MERGE-GATE-implementer-<ts>.md` with `ACTIONS_GIT_REF` carrying the real refs.

ACTIONS_GIT_REF: none by this relay — the dispatch authorizes; the executor acts (disk refs: this relay + one INDEX.md row timestamped 20260714-020520).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0` (`s10-close`); the `s9-evidence-thicken` worktree clean at `d91fcfb`.
Next requested action: operator carries this to s9.implementer; the merge executes; the report returns to this thread; master then runs the s9 step-integration verification + the close fold (RECONCILE §s9 · dashboard · ROADMAP · the s9 carries into their queues) and cuts the fresh s11 dispatch (the last Step-2 build slice).
