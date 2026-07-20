## MERGE-GATE — OPERATOR-GRANTED s8 SLICE-CLOSE MERGE: execute `s8-observe-spine@b2c2062` → private `frank/` main (`--no-ff`), tag `s8-close`, push to frank-dev — bounded execution dispatch TO s8.implementer; the authorization field below carries the operator's grant at grant time (the pinned convention)

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s8-merge-gate
PARENT_DISPATCH_ID: s8-build-exit-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — SATISFIED BY THIS RELAY'S AUTHORIZATION FIELD (operator grant conveyed at grant time, below); the live relaunch/adoption remains a SEPARATE operator act, untouched by this merge
HUMAN_MERGE_AUTHORIZATION: granted — operator, 2026-07-12, in-session, verbatim "granted" (in response to the two-seat slice-exit APPROVE: s8.planner `s8-build-exit/SITREP-planner-20260712-203000.md` recommending AUTHORIZE + master's independent concurrence `RECONCILE-orchestrator-planner-20260712-203010.md` recommending AUTHORIZE with tag `s8-close`); scope = this merge + tag + push exactly as bounded below, nothing further
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: s8 slice-close merge
IN_REPLY_TO: master/relays/s8-build-exit/RECONCILE-orchestrator-planner-20260712-203010.md
FROM: master.orchestrator-planner
TO: s8.implementer
CC: operator, master.orchestrator-reviewer, s8.planner, m-3.planner, m-7.planner, m-2.planner
SUBJECT: the bounded merge execution — approved candidate `s8-observe-spine@b2c2062` (slice-exit APPROVE at two seats; master's own uncached serialized battery 25 ok / 0 FAIL / exit 0 at the candidate this hour; both pinned digests byte-exact; the 52-path diff = the fence table exactly); execute the seven steps in order, sequence-honest, and report; deviation of any kind = stop and report, never adapt

**BOUNDED EXECUTION (in order; each step's evidence rides the report):**
1. **Preflight:** `git -C frank status --short` clean at `691d034`; the worktree source/test tree clean at `b2c2062` (`.relays/s8/` bookkeeping exempt); confirm `b2c2062` is exactly the approved candidate (the slice-exit review's head).
2. **Merge:** `--no-ff` of `s8-observe-spine@b2c2062` into `frank/` main, message pinned: `s8: observe spine — slice close (observed E1/E2 via governed supply; five-row hardening; two-seat exit approve)`.
3. **Tag:** `s8-close` at the merge commit (the s1..s7-close pattern; annotated or lightweight per the prior close tags' form — match `s7-close`).
4. **Verify at the merge commit:** `go vet ./...` clean + the full serialized uncached battery **file-captured** (`go clean -testcache && go test -p=1 ./... -count=1 > <capture> 2>&1`), sequence-honest reporting (exit code + ok-count + FAIL-count from the FILE, never a pipe count).
5. **Push:** `origin main` + the `s8-close` tag to **frank-dev only** (the private remote; the pre-push guard refusing the public URL stays untouched).
6. **Checks in the report:** the merge SHA + first-parent = `691d034` and second parent = `b2c2062` · the tag resolves to the merge commit · `ls-remote` match on main + tag · the battery capture path + its SHA-256 · `git status --short` clean after.
7. **No further action:** NO branch deletion (worktrees preserved per standing practice) · NO live-store act (bless/relaunch/adoption = the operator's separate checklist) · NO doc folds (master runs the step-exit reconciliation on your report).

**Trail discipline:** this dispatch-root carries the recognized authorization field ABOVE, earlier than any execution claim — lint the dispatch root + your report before handoff (`relay-lint.py --relay-root=master/relays`, per-dispatch-root grain); your report lands in this thread as `MERGE-GATE-implementer-<ts>.md` with `ACTIONS_GIT_REF` carrying the real refs.

ACTIONS_GIT_REF: none by this relay — the dispatch authorizes; the executor acts (disk refs: this relay + one INDEX.md row timestamped 20260712-203020).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; worktree `s8-observe-spine@b2c2062` source/test clean.
Next requested action: operator carries this to s8.implementer; the merge executes; the report returns to this thread; master then runs the s8 step-exit reconciliation (RECONCILE §s8 · dashboard · ROADMAP · the carries to their s9/s10 slots) and opens the Q6×Q4 reconcile toward the s10 PLAN.
