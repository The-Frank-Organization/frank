## MERGE-GATE — OPERATOR-GRANTED s11 SLICE-CLOSE MERGE (the LAST Step-2 build slice): execute `s11-comms-thicken@502e06c` → private `frank/` main (`--ff-only`), tag `s11-close`, push to frank-dev — bounded execution dispatch TO s11.implementer; the authorization field below carries the operator's grant at grant time (the pinned convention)

ROLE: Orchestrator Planner
PHASE: MERGE-GATE
AUTHORITY: merge-gated
DISPATCH_ID: s11-merge-gate
PARENT_DISPATCH_ID: s11-merge-decision
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — SATISFIED BY THIS RELAY'S AUTHORIZATION FIELD (operator grant conveyed at grant time, below); the public repo flip, the live relaunch/adoption, and any publication push beyond private frank-dev remain SEPARATE operator acts, untouched by this merge
HUMAN_MERGE_AUTHORIZATION: granted — operator, 2026-07-14, in-session instruction to master ("write me a dispatch merge"), granting the s11 comms-thicken slice merge on the unanimous exit package: s11.planner's reissued unconditional recommendation at the new head (`s11-merge-decision/MERGE-GATE-planner-20260714-175200`, owner-confirm set unanimously clean — m-6/m-3 CONFIRM + m-7 all cells with the r13 contradiction DISCHARGED by the ruled v4 fold + hunk-by-hunk owner countersign) + master's independent concurrence (own uncached serialized battery exit 0 / 0 FAIL / vet clean / ten INV-CATALOG laws green at `502e06c` this session; the v4 fold verified spec-exact hunk-by-hunk against the (a)-ruling `s11-build-escalate-config-lock/RECONCILE-…-170510` — v3 descriptor RESTORED, v4 admits `resummon_cadence` optionally, v3→v4 the sole new adjacent-forward hop with rollback rejected); scope = this merge + tag + push exactly as bounded below, nothing further
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
TASK_ID: s11 slice-close merge
IN_REPLY_TO: master/relays/s11-merge-decision/MERGE-GATE-planner-20260714-175200.md
FROM: master.orchestrator-planner
TO: s11.implementer
CC: operator, master.orchestrator-reviewer, s11.planner, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-7.implementer, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the bounded merge execution — approved candidate `s11-comms-thicken@502e06c` (reissued slice-exit recommendation at the new head; owner-confirm set unanimously clean incl. the discharged r13 contradiction; master's own uncached serialized battery exit 0 / 0 FAIL at the candidate this session; base `main@d91fcfb` is a true fast-forward, 18 linear commits; tag + remote-tag both absent; tree clean); execute the seven steps in order, sequence-honest, and report; deviation of any kind = stop and report, never adapt

**BOUNDED EXECUTION (in order; each step's evidence rides the report):**
1. **Preflight:** `git -C frank status --short` clean at `d91fcfb`; the `s11-comms-thicken` worktree source/test tree clean at `502e06c` (`.relays/s11/` bookkeeping exempt); confirm `502e06c` is exactly the approved candidate (the reissued merge-decision head) and `git merge-base --is-ancestor d91fcfb 502e06c` holds (a true fast-forward, 18 linear commits); confirm local `s11-close` and remote `s11-close` are both absent.
2. **Merge:** `git -C frank merge --ff-only 502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` into `main` — fast-forward, no merge commit (the slice boundary is marked by the tag). If `--ff-only` fails for any reason, STOP and report — do not fall back to `--no-ff`.
3. **Tag:** annotated `s11-close` at the new `main` head `502e06c` (match the s7/s8/s9/s10-close form), message naming the slice — e.g. `git -C frank tag -a s11-close 502e06c -m "s11: comms-thicken slice close (B/C/D buckets + 7-state FSM + g1 8a hardening; T8 8-of-9; cadence re-homed at engine v4; the last Step-2 build slice)"`.
4. **Verify at the merge commit:** `go vet ./...` clean + the full serialized uncached battery **file-captured** (`go clean -testcache && go test -p=1 ./... -count=1 > <capture> 2>&1`), sequence-honest reporting (exit code + ok-count + FAIL-count read from the FILE, never a pipe count); confirm `test/invariants` (the ten laws) green.
5. **Push:** `origin main` + the `s11-close` tag to **frank-dev only** (the private remote `https://github.com/iwnlcern/frank-dev.git`; the pre-push guard refusing the public `iwnlcern/frank` URL stays UNTOUCHED — the public flip is the operator's separate act).
6. **Checks in the report:** the post-merge `main` SHA (= `502e06c` on a fast-forward) · the `s11-close` tag resolves to it · `ls-remote` match on main + tag · the battery capture path + its SHA-256 · `git status --short` clean after.
7. **No further action:** NO branch deletion (worktrees preserved per standing practice) · NO live-store act (bless/relaunch/adoption = the operator's separate checklist) · NO public-remote push · NO doc folds (master runs the s11 step-integration + the Step-2 step-exit test on your report).

**Trail discipline:** this dispatch-root carries the recognized authorization field ABOVE, earlier than any execution claim — lint the dispatch root + your report before handoff (per-dispatch-root grain); your report lands in this thread as `s11-merge-gate/MERGE-GATE-implementer-<ts>.md` with `ACTIONS_GIT_REF` carrying the real refs.

ACTIONS_GIT_REF: none — this relay authorizes only; the executor acts (disk refs: this relay + one INDEX.md row timestamped 20260714-175210).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb` (`s9-close`); the `s11-comms-thicken` worktree clean at `502e06c`.
Next requested action: operator carries this to s11.implementer; the merge executes; the report returns to this thread; master then runs the **Step-2 step-exit test** on the merged spine (all three legs live on the dogfood store + the INV-CATALOG red-battery demo + the uncached green battery) and the s11 close fold (RECONCILE §s11 · dashboard · ROADMAP · the T5/T10 gate-bound disposition + the s9/s11 carry queues). NOTE: if the operator prefers master execute directly (I hold the frank worktree this session), retarget `TO: master.orchestrator-planner` on your word — same authorization, one fewer hand-relay.
