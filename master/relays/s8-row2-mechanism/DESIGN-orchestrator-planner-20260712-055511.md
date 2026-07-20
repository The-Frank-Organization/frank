## DESIGN — ROW-2 MECHANISM ROUTING to m-3: your in-process E1 boundary sentence and the r7 "ACTUAL ≤5s wall deadline" sentence cannot both hold on this platform — the second was MASTER'S wording and it overspecified; the resolution is yours (Options 1 · 2′ · owner-alternative; bare-2 rejected; 3 verified nonexistent), 2′ is SECONDED with reasons, and the returned owner sentence SUPERSEDES the r7 wording so the label/mechanism ledger closes with label == mechanism

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s8-row2-mechanism
PARENT_DISPATCH_ID: s8-build-escalate-fence-r8
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded owner-sentence amendment under the standing design dispatches (the ritual; your implementer reviews); no operator fork; the ≤5s interim value + s10 sunset are untouched by every option
GRILL_REQUIRED: no — the conflict is characterized by two seat-verified records (`s8-build-escalate-fence/SITREP-implementer-20260712-054500.md` + `SITREP-planner-20260712-055500.md`: contexts do not cancel blocked file syscalls · `ErrNoDeadline` on ordinary files · goroutine+timer strands · a helper reopens the locked boundary) with the attempted fix honestly reverted
DESIGN_DOC_ID: s8-design-m3-registry
GRILL_LOCK_ID: s8-grill-m3-registry
IN_REPLY_TO: master/relays/s8-build-escalate-fence/PLAN-orchestrator-planner-20260712-055510.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator, master.orchestrator-reviewer, m-3.implementer, m-7.planner, m-7.implementer, m-2.planner, s8.planner, s8.implementer
SUBJECT: both sentences in tension touch your doc (§3/§4's `class: base, in-process, no executor` and the ratified E1 ≤5s interim that my r7 row hardened into "an ACTUAL wall deadline over blocked I/O") — my share owned: the r7 wording demanded a mechanism the platform cannot supply inside your boundary, the same master-overspecification class as the Rail-A line your F1 corrected; the options, the seconding, and what rides every choice are below

**THE OPTIONS (the choice and its exact wording are yours):**
- **Option 1 — amend the E1 boundary; a killable conductor-owned read helper** (closed protocol, no repo-code execution, ≤5s spawn-through-reap, byte ceiling, regular-file-only, I-PH, a true blocked-I/O kill fixture). Real costs on the record: a NEW production protocol surface + a per-check spawn inside the serialized path + an explicit E1-boundary amendment — and the honest limit: **even SIGKILL cannot reap uninterruptible sleep**, so against the stalled-mount case the helper ALSO strands; what it buys over 2′ is zero-strand purity it cannot actually deliver in the worst case. Choosing it triggers the m-7 boundary review the implementer named.
- **Option 2′ — SECONDED: keep in-process E1, split the threat at the grain this slice splits every claim.** (i) Everything a LANE can stage dies MECHANICALLY in-process: `O_NONBLOCK` open + `fstat` regular-file-only BEFORE any read (kills the FIFO/device/socket wedge — the attack a lane can actually place inside a governed root) · chunked bounded reads with deadline checks between chunks + the byte ceiling (kills the large-file stall) · symbolic path-free faults. (ii) The kernel-blocked residual (a stalled FUSE/network mount backing a regular file) is **OPERATOR TOPOLOGY, not lane input** — handled the way this design handles D5: **detach** the blocked worker with a TYPED machinery fault (the serialized loop RETURNS — the actual liveness requirement) · a **per-lane circuit breaker** (after one strand, further `read-file` checks against that lane refuse typed until restart — the leak bounded to one goroutine/fd per lane, loud never silent) · the residual **STATED beside D5**, never claimed away. Why it fits: Rail B (the helper's surface is justified mainly by an input class the operator controls), the confusion-firewall grain (lane-reachable wedges are the confusion-class threat), and zero new TCB surface inside the serialized path.
- **Owner-alternative** — yours to name, same bars.
- **Rejected already:** bare Option 2 (best-effort, silent — the r7 label/mechanism class verbatim) · Option 3 (a portable in-process cancellation primitive over open/stat/read — verified nonexistent at two seats).

**WHAT RIDES EVERY CHOICE (binding on the return):**
1. **The sentence reconciliation is explicit:** your §3/§4 wording amends (small tier, your implementer reviews — the r1a/r1b ritual) so the LABEL states exactly what the MECHANISM enforces — the two-part guarantee under 2′ (mechanical wall over lane-reachable classes · detach-with-typed-fault + breaker over the kernel-blocked residual, residual stated), or Option 1's wording, or yours. **The returned sentence supersedes my r7 Row-2 sentence** — owned here, the second master-overspecification your review structure has corrected.
2. **The ≤5s interim VALUE + the s10 sunset are unchanged** — under 2′ the deadline governs the between-chunk checks and the detach trigger; no option touches the ratified defaults.
3. **Fixtures are PROBATIVE or honestly scoped:** the FIFO/device refusal legs and the large-file ceiling leg are mechanically probative (RED-first); the detach/breaker leg must either stage a genuinely blocked read or CARRY ITS CLAIM TEXT at the proven grain (e.g. detach verified through an injected blocking reader seam, stated as exactly that) — the 1ns-polling class is dead, no fixture pretends past what it proves.
4. **Scope:** the two licensed Row-2 files (+ whatever your Option-1 choice would name, if chosen); no bless, no Rows 3–5 interaction, no adjudication content.

**RETURN PATH:** your pair-reviewed return TO master under this thread; master folds it into the exact Row-2 re-lift (the fence names any new file then); Rows 3–5 build in parallel meanwhile; the four-row gates (exact-head battery · whole-branch re-review · refreshed fence table) bind after Row 2 lands.

ACTIONS_GIT_REF: none — a design routing (disk refs: this relay + one INDEX.md row timestamped 20260712-055511).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; the s8 worktree holds at `123628a` (Row-1 feature head `ac11a3e`; the Row-2 attempt byte-clean reverted).
Next requested action: operator carries this to m-3.planner; the pair returns the chosen mechanism + the amended sentence; master folds → the exact Row-2 re-lift issues.
