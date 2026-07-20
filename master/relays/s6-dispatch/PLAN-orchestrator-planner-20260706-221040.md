## PLAN — Slice-6 build dispatch (master → the s6 slice-team): BUILD THE TRANSPORT FIX — the VP-co-signed amendment set implemented whole; exit = the Step-1 step-exit test ON the fixed conductor; Step-1 closes on this slice

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s6-dispatch
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — the operator gates already exercised (the in-step ruling; the fork grill; the boot addendum); the s6-close sign-off + the step-exit test's operator legs are exercised downstream
DELEGATED_DISPATCH_AUTHORITY: yes — F2 conditions below (non-bootstrap)
GRILL_REQUIRED: no — completed at design (`GRILL-LOCK-parenting-fork-2026-07-06`)
IN_REPLY_TO: master/relays/s6-design/RECONCILE-orchestrator-reviewer-20260706-220325.md
FROM: master.orchestrator-planner
TO: s6.orchestrator-planner, m-7.planner
CC: master.orchestrator-reviewer, operator, m-1.planner, m-2.planner, m-7.implementer, m-1.implementer, m-2.implementer
SUBJECT: s6 = implement `S6-AMENDMENT-SET-2026-07-06` (r3, VP co-signed `…-220325`) at `main` post-`s5-close` — branch-A parenting · the one codec · A-1..A-4 · §B/§C/§D · F13 · D-1/D-2 · the boot stage (B-1/B-2/B-3, derived-only) · the seven-row registry pass; 21+ fixtures red-first; exit = battery + fixtures + THE STEP-EXIT TEST on the fixed conductor; NEW slice-team; guide m-7; F2

**What this is.** The build dispatch for **Slice-6 = the transport fix** — the last Step-1 slice. The design is DONE and locked: **`master/S6-AMENDMENT-SET-2026-07-06.md` (r3, VP co-signed)** integrates the three pair-complete domain amendment docs; **the set and its constituents are your spec — this dispatch points, it does not restate.** Step-1 closes when your exit gate passes (the operator's in-step ruling, 2026-07-06). Baseline: `frank/` `main @ 7e5c527` (tag `s5-close`), battery 23 green.

### To the s6 slice-team — your charter
- **NEW slice-team** (new sprint = new team). **Use `/orchestrator-planner`**; `sprint-doc-setup` in `frank/`; relays under `frank/.relays/s6/…`; file-relay transport (the operator hand-relays; frank carries governance again only AFTER you fix it — no dogfooding this slice).
- **Onboard first — you built none of s1–s5.** Read the source + the five sprint ledgers + `TRANSPORT-FINDINGS-2026-07-06.md` (why every change exists) + the set + its four constituent docs (m-1/m-7 amendment docs, m-2 codec amendment, the GRILL_LOCK); **re-run the battery at `s5-close` yourself, uncached** (the standing bar: fresh eyes have found real fragility six runs straight).
- **Spec = read-only:** the set (r3) + the domain amendment docs are the authoritative deltas (the ARCHITECTURE §C4 pointer marks them design-of-record; older §C4 prose loses conflicts). Escalate spec problems; **do not self-amend** — locked-contract touches go through the amendment path.
- **Build on `main` @ `7e5c527`**, on a branch (suggest `s6-transport-impl`), worktree operator-set. Team granularity yours; the operator assists with worktrees as you decide.

### Guide + fidelity edges
**m-7 primary guide** (the engine is most of the diff: A-1..A-4, B-1, the parenting stamp locus, D-1/D-2). **m-1 fidelity** on every store/lineage/waiver/lock/activation touch (its §A–§F are the contracts; its §E/§F.1 carry-forwards are IN the locks). **m-2 fidelity** on the codec, the registry pass, the boot form, and every render/validate surface. All questions via master (file relay) — the m-x owners answer through the standing hub.

### Scope (IN — the set, whole; its §"Build-slice obligations" is the authoritative list)
Branch-A parenting (conductor-computed PARENT + `parent_hint` fallback semantics per the GRILL_LOCK) · the ONE canonical `address_list` codec (F6/F7 class deleted) · A-1 stable-schema digest · A-2 idempotent-replay intake + durable monotonic ids · A-3 live mint (`seat_mint` pivot; CLI mint retired to genesis-time) · A-4+§D the store lock (I1-P; proof-of-death takeover; refuse-reads loser) · §B `project()` default-accepted + the accepted-graph anchor fix · §C scoped waivers + `waiver_retraction` (commit-order effective state) · F13 three-layer record_kind authorization · D-1 shim transparent-reconnect · D-2 bounce/reply detail parity · **the boot stage**: B-1 lifecycle `minted→bound→active` + the literal admission allowlist + the roster `project` view · B-2 the boot form (`SITREP` + lifecycle-gated) · B-3 **derived-only activation** (first accepted governed submit per mint-generation — **[VP-W3 of `…-220325`]: the registry pass carries EXACTLY the seven named rows; NO activation-marker row may be reintroduced**) · the registry pass (seven rows + `waiver_retraction`/`seat_mint` record classes + the `ORCH_REVIEW_WAIVER` `"*"` header retired; MINOR, **no envelope migrator** — R-1).

### Scope (OUT — escalate before touching)
Step-2 observe (declared fields stay dormant) · routing execution (Step-3; the C1/C2 §C4 carries are the router builder's) · **engine performance work of any kind** (the ledger's latency addendum: exonerated by measurement — do not optimize) · new seat verbs · federation · any locked-contract change outside the co-signed set (amendment path only).

### Exit gate (HARD; Step-1 closes on it)
1. **Fixtures, red-first, all green:** the GRILL_LOCK three (the archived-dogfood **F11 replay without livelock** · hint honored/fallback · concurrent-accept-no-parent-bounce) · m-7 **FX-A1a..FX-B1g — all 18, FX-B1g explicitly [VP-W2]** (re-mint ⇒ new generation starts `minted`; pre-re-mint accepteds do NOT activate; a fresh boot accept does) · m-1 §E (polluted-archive projection filter · reject-never-anchor · waiver scope/retraction/re-arm · the lock race/alias/kill-9 legs) · m-2's (codec round-trip · full-set-in-every-projection · three-layer record_kind + the genesis negative · waiver-as-record · boot renders-pre-active + un-bounceable · shared-vocab negative).
2. **The E2 floors:** full battery green (s1–s5 suites + yours), uncached, zero regression; byte-exact enum; the three-verb surface; **I-PH swept over every NEW surface** (roster view, boot bounces, lock refusals, hint flags — path-free, `Field:Class` shape preserved).
3. **THE STEP-EXIT TEST (live, on your fixed conductor — Step-1's closing bell):** a fresh blessed store; live seats wired (ops: **pre-allowlist `mcp__frank__*`** per the ledger addendum); then (a) the ROADMAP:83-85 legs — a relay accepted ONLY through the conductor, FROM system-stamped, form validation pre-delivery, a gate → local outbox item; (b) **the first live act = the operator §7-applies s5's registry** (operator-authored — the adapted [VP-W7] discharged); (c) **the F11 regression leg: the archived dogfood store's traffic pattern replayed live against the fixed conductor — completes without livelock**; (d) a live boot: mint (A-3, no restart) → wire → the B-2 boot → `active` derived → the roster shows it.
4. **Honesty:** transport/provenance-only phrasing travels; done-state + `record_integrity` stay `self_reported`; the ③/⑤ dormant claims unchanged; no claim exceeds the §C4.3 ceiling.

### Plan-gate (F2 — non-bootstrap) + process
Pair-implementer plan-review = the plan gate; `DISPATCH IMPL` under the standard F2 conditions; escalations to master via file relay. **[VP-W1 of `…-220325`]: if m-7 exercises its one-relay revert window on the two master-applied stale-text corrections, master holds this dispatch and reconciles first — the operator relays this dispatch only after that window passes quiet.** Exit SITREP to master; master runs its independent battery + probes at your close tip (the standing discipline); integration = the operator's gate; **then the Step-1 close fold.**

### Not authorized
No s6-close authority, no scope beyond the set, no design amendment (route back), no Step-2 pre-work, no perf work, no dogfood-in-slice (frank carries its own governance again starting with the relaunch AFTER this slice proves the fix).

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s6-dispatch` — run below.
- Chain: the co-sign `s6-design/RECONCILE-orchestrator-reviewer-20260706-220325` ([VP-W1..W3] folded above) · the set r3 · the GRILL_LOCK · the in-step operator ruling (dashboard/RECONCILE/ROADMAP, 2026-07-06) · `frank/` `main @ 7e5c527`.

ACTIONS_GIT_REF: wrote this dispatch + the s6 boot + a VP pre-handoff request + the §C4 pointer + the dashboard bullet + INDEX rows; no code, no `frank/` write, no dispatch token; cwd is a docs workspace (not a git repo).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main @ 7e5c527` (tag `s5-close`), clean.
Next requested action: VP pre-handoff review (request alongside; package HELD); on approve — and the m-7 revert window quiet — the operator relays the boot then this dispatch to a fresh session; s6 onboards, plans, builds; exit SITREP at the gate; Step-1 closes on it.
