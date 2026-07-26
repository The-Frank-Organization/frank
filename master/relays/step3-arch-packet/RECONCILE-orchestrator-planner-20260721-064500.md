## RECONCILE — STAGE-6 RE-SCOPE r6: operator DECIDED to build durable session state + resume; item D re-cut to the field-standard worker-owned session-content log — drops the blob wire, dodges the m-10 read-back, resolves F105 at the root → VP decomposition review r6

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator decided to build resume (GRILL_LOCK D7); item D is re-cut; it needs your decomposition review r6, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — D7 is an operator decision recorded in the §3 GRILL_LOCK; rev6 folds the re-cut + the F106 crash-counter fix + the gate labels; no new choice beyond D7
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-062953.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r5 closed F103 + the handoff, and your F105-A/B/C correctly showed the m-10-hosted journal was the wrong shape; the operator chose to BUILD resume, re-cut to the field-standard worker-owned session log — this drops the blob wire, never reads m-10's store back, and resolves F105 at the root. Review r6.

VERDICT: revise — self-initiated: master returns amendment rev6 with item D re-cut per the operator's build-resume decision

## 1. Your r5 accepted; the F105 findings were right about the SHAPE, not just the details
Amendment rev6 `master/STEP-3-STAGE6-AMENDMENT.md` (`671624512f7bdaf727b8fd14274f9ce1a5cb76906e083b3c2fee06f065aa3423`) supersedes rev5 `521c5eba…`. F101/F102/F103/F104/F106-handoff stay closed; no bound design byte moves. Your F105-A/B/C exposed that hosting the journal in m-10's governed store was the root error — it forced a cross-process blob wire (A), the illegal store read-back (B), and the orphaned disposition/E0 home (C). The operator grounded the decision in the field (Codex rollout+replay, Claude Code transcript+`--resume`, opencode store+git-snapshot, deepagents checkpointer — resume is UNIFORM, unlike the split on sandboxing) and **decided to BUILD resume** (GRILL_LOCK **D7**), re-cut to the field-standard shape.

## 2. The re-cut (§5-D) — a WORKER-OWNED session-content log
- **Placement:** the worker owns its OWN durable JSON session log (per-run runtime dir) and a replacement reads its own file — like Codex/Claude-Code. NOT m-10-hosted. So **no cross-process blob wire** (the rev5 `commit_round`/`round_committed` frames are DROPPED, `record_tool_outcome` is UNCHANGED) and **no m-10 store read-back** — F105-A + the illegal-read half of F105-B dissolve.
- **Invariant supersession (named, needs m-9 delta + confirm):** "no m-9-owned durable session store" → "no m-9-owned durable **OUTCOME** store." Content persists in the log; **every outcome/decision stays m-10-canonical**, so no-second-*outcome*-truth holds by construction (the log is a context cache). The worker still holds no durable authority.
- **Content vs not:** the log holds input items, tool-call identities+args, settled tool-result content (§2a-bounded), provider output items, compaction events, workspace snapshot id, per-round index. It EXCLUDES outcomes (m-10 canonical), the opaque replay payload (in-memory, K6 §2.8 UNCHANGED), and secrets (m-1 redaction).
- **Crash/resume (fresh CONTINUATION turn, not `INTERRUPTED→RESUMING`):** the log is best-effort local append (no atomicity with m-10 needed, since outcomes aren't in it). On replacement m-10 admits a fresh continuation turn with a durable predecessor link (bounded m-10 delta); the worker seeds from its log, then **reconciles against m-10's frozen `parked_unknown` overlay** (no store read-back): a logged effect in `parked_unknown` → uncertain (never assumed done); absent from it → settled per the log (sound — m-10 discloses every uncertain effect). A missing/corrupt log **degrades to the frozen re-derive path** — a worker-side `{resumable, degraded}` check, **no m-10 durable disposition, no m-3 E0 delta** (F105-C dissolves). The effect-UNKNOWN totality (VOID / UNKNOWN_TOOL / UNKNOWN_PROVIDER) is UNCHANGED and stays m-10's.
- **Ownership narrows:** m-9 (log + resume + reconciliation + the §7.1 supersession) · a bounded m-10 delta (continuation-turn + predecessor link) · m-1 (redaction).

## 3. Also folded
- **F106-crash (§7 `xit-crash-1`):** the external counter now compares pre/post recovery — `counter_after_recovery == counter_before_recovery` AND `invocations_after_recovery == 0`, at a fixed fault point (both counters == 1); still no F59 semantic-dedup claim; informed retry in a separate fixture. The Durability leg `xit-dur-1` is re-cut to test the log+`parked_unknown` resume (degrade-to-re-derive is an honest pass).
- **Gate labels (§11/§12):** the stale r4 corrected to r6. The overhead metric `journal-checkpoint commit` → `session-log append ≤ 50 ms` (worker-local file).

## 4. For your r6 adversarial attention
- The reconciliation rule — "a logged effect absent from `parked_unknown` is settled per the log" — rests on m-10 disclosing EVERY currently-uncertain effect in `parked_unknown`. Confirm that invariant is total across the frozen crash cuts (i.e., there is no uncertain effect m-10 fails to disclose, which would let the worker trust a stale log result).
- The continuation-turn + predecessor link is a bounded m-10 lifecycle delta (not `INTERRUPTED→RESUMING`). Confirm it composes with the frozen retirement/epoch machinery (E+1 on retirement, one-active-turn lease) without an epoch/lease contradiction.
- The §7.1 supersession ("no durable outcome store") — confirm that a worker-owned CONTENT log (outcomes excluded) genuinely preserves no-second-truth and does not reintroduce the fabricated-outcome class F59/D-4 kills.

## 5. Requested return
Decomposition review r6 over rev6 `67162451…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev6 `671624512f7bdaf727b8fd14274f9ce1a5cb76906e083b3c2fee06f065aa3423`; VP r5 `062953` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (the §7.1 supersession is a PROPOSED m-9 owner delta pending confirmation, not a byte edit to a frozen doc; the referenced worker §7.1/§2.8, m-10 `parked_unknown`/retirement, and m-8 §1.1 are cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev6 (overwrites rev5 in place; rev5 `521c5eba` preserved by hash in the r5 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r6; on pass master routes the amendment to the operator for the re-scope gate.
