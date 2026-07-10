## BOOT — initialize s6.orchestrator-planner for RUN_ID s6 (Slice-6: THE TRANSPORT FIX — the last Step-1 slice)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-boot-orchestrator-planner
PARENT_DISPATCH_ID: s6-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: s6.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.planner
SUBJECT: BOOT — initialize s6.orchestrator-planner for RUN_ID s6 (the transport-fix slice; a NEW slice-team; Step-1 closes on your exit gate)

You are **s6.orchestrator-planner for RUN_ID s6** — the orchestrator-planner of a **NEW slice-team** standing up **Slice-6 = the transport fix**. New sprint = new team: you built none of s1–s5. Your job: stand up your team, plan s6, and drive it to the exit gate — **which is Step-1's closing bell** (operator ruling 2026-07-06: Step-1 does not close until the transport is fixed). Transport = **file relays** (the operator hand-relays) — frank does NOT carry governance this slice; your build is what makes it trustworthy again.

**Where the build is.** s1 (spine) · s2 (engine) · s3 (the full form system) · s4 (the wire-up — the operator-as-transport ended live) · s5 (consumer schemas, registry 47→83 rows) are **CLOSED at E2** — baseline `main @ 7e5c527`, tag **`s5-close`**, battery 23 packages green. **The mandatory owed set is EMPTY** — what rides in is the work itself: the s5 dogfood (the first team to run its governance ON frank) surfaced **17 transport findings** — headlined by **F11, a lineage livelock that blocked even status reports under multi-seat load** — the operator ruled stop-the-line, and a full design-amendment cycle produced your spec: **`master-docs/master/S6-AMENDMENT-SET-2026-07-06.md` (r3, VP co-signed `s6-design/RECONCILE-orchestrator-reviewer-20260706-220325`)**.

**Come online:**
1. Load **`/orchestrator-planner`** (+ `protocol.md`; brings `sprint-doc-setup`).
2. Read the team charter: **`CLAUDE.md`** / `AGENTS.md` (auto-loads — your cwd is inside the harness tree; if not, read `master-docs/CLAUDE.md` explicitly).
3. **Read the story before the spec:** `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md` — WHY every change exists (the archived dogfood store `~/frank-archives/frank-team-store-s5-dogfood-20260706` is the evidence and later your F11-replay input). THEN the spec (**read-only**): the amendment set (r3) + its four constituents — the **m-1** + **m-7** `2026-07-06-s6-transport-amendments.md` domain docs, the **m-2** `2026-07-06-s6-transport-codec-amendment.md`, and **`master-docs/master/GRILL-LOCK-parenting-fork-2026-07-06.md`** (the operator-grilled parenting fork: branch A, fallback-never-bounce, the empirical basis). The ARCHITECTURE §C4 pointer marks these design-of-record; where older §C4 prose conflicts, **the amendment docs win**. Do not edit governance docs; escalate spec problems.
4. Read your **work dispatch**: **`.relays/s6/s6-dispatch/PLAN-orchestrator-planner-20260706-221040.md`** — scope (the set, whole) · the OUT fence · the exit gate (fixtures + battery + THE STEP-EXIT TEST). Your authority comes from it, not this boot.
5. **ONBOARD — you built none of s1–s5.** Read the source (`frank/`, Go, `github.com/jackli/frank`) + all five sprint ledgers (`docs/sprints/…`); **re-run the battery at `s5-close` yourself, uncached.** The standing bar: every fresh team so far — six straight — found real fragility the builders missed. Look hardest at `internal/lineage` (the F11 mechanics you're deleting), `internal/fieldspec` (render/validate — the codec's home), `internal/intake` (F9), and `internal/channel`/`cmd/frank` (mint, lock, lifecycle).
6. **Scaffold an `s6` sprint** via `sprint-doc-setup` in `frank/`; your relays live there (`.relays/s6/…`); stand up your sub-seats (granularity = your call) with their own boots. Code work in worktrees the operator sets up off `main @ 7e5c527` on your branch call (suggest `s6-transport-impl`).

**THE ONE-LINE BOUNDARY:** you build the **TRANSPORT FIX** — the co-signed amendment set, whole, nothing else — against the **LOCKED amendments**: the seat surface stays exactly `submit`/`project`/`read` (roster + audit views are `project` parameters, never a fourth verb); **m-7 guides** (the engine is most of the diff); **m-1 fidelity** on every store/lineage/waiver/lock/activation touch (its §E/§F.1 carry-forwards are IN the locks); **m-2 fidelity** on the codec, the registry pass, the boot form, every render/validate surface. **The set rules; escalate to amend — never improvise a locked-contract change.**

**Gate + escalation (F2 — non-bootstrap):** pair Implementer plan-review + conditioned delegated dispatch — `DISPATCH IMPL` only under {Implementer approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract/design-of-record amendment}. Any failure — including any **OUT**-item touch (**Step-2 observe** · **Step-3 routing execution** · **engine performance work of ANY kind** — the ledger's latency addendum exonerated the engine by measurement, do not optimize · new seat verbs · federation · dogfood-in-slice) — **escalates back to master (CTO + m-7 guide + VP)**. All m-x questions route through master as file relays via the operator.

**Carried context:** the co-sign's build watchpoints ride your slice — **[VP-W2] FX-B1g** (the re-mint/generation leg) is explicitly in your exit gate; **[VP-W3] the registry pass is EXACTLY the seven named rows — no activation-marker row may be reintroduced** (activation is DERIVED-ONLY: first accepted governed submit per mint-generation). The GRILL_LOCK's semantics are binding: unprovable `parent_hint` ⇒ **fallback, never bounce** (+ `parent_hint_honored: no` in the submit response) — the empirical basis (zero true bad picks in the entire recorded history) is in the lock; carry the *why* into your fixtures. Honesty rides into code: **transport/provenance only** — done-state + `record_integrity` stay `self_reported`; the ③/⑤ dormant claims unchanged; I-PH over every NEW surface (roster, boot bounces, lock refusals, hint flags — path-free, `Field:Class`). **The step-exit test's operator legs are the operator's** (the §7 application of s5's registry; live-seat designation) — that is the mechanism working, never yours to simulate (the s4 honesty line). Ops note for the live legs: pre-allowlist `mcp__frank__*` in seat sessions.

Relay root: your `frank/` sprint tree. relay-lint: `~/.claude/skills/tools/relay-lint.py` — lint **exact-file AND root-mode**; keep dispatch dirs clean (superseded-file residue has cost waivers before).
Current authority: **report-only onboarding.** This boot grants no PLAN/IMPL/REVIEW authority; that comes from the s6-dispatch + your own pair-lineage under F2.
Acknowledge (file relay to master): identity (`s6.orchestrator-planner`, RUN_ID s6), loaded skill, reachable `frank/` + relay setup, **your own uncached battery result at `s5-close`**, the one-line boundary, the guide(m-7)/fidelity(m-1, m-2) split, the F2 gate + OUT-escalation (incl. the no-perf fence), the [VP-W2/W3] watchpoints, the fallback-never-bounce semantics, and the operator-legs honesty line; plus your proposed team shape + worktree ask; then proceed per the dispatch — onboard, scaffold, plan.

ACTIONS_GIT_REF: none — report-only boot onboarding; no code/`frank/` edit ("commit / merge / lock / mint" references denote engine concepts + build history, not git action claims by this relay). *(Boot r2: pattern-conformance fold to the s3/s4 boot shape — build-state stanza, scaffold step, boundary line, F2/OUT + carried-context blocks, lint discipline, full ack checklist; all content restated from the VP-approved dispatch, no new substance.)*
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main @ 7e5c527`, clean, tag `s5-close`.
