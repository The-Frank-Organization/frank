# Step-3 Lane-4 — nested-team INERT kickoff brief — ⛔ VOID / SUPERSEDED (nested-on-frank; retained as history)

> **VOID as of 2026-07-25.** frank-as-courier was stood down for lane 4 (**B22**) and the team shape was set to
> **pair** (**B23** — the durable record; B22 itself left the shape open). The nested team this brief describes no
> longer exists. Superseded by the rev9 kickoff (to be authored on VP approval of plan rev9). **Grants no action.** Two defects found by the live
> preflight are recorded here so the history reads honestly: §7 names `PARENT_DISPATCH_ID`, a field that
> **does not exist on the frank wire** (lineage is `headers.parent_hint` + `headers.IN_REPLY_TO`); and its
> worker-leg parent map asks a worker to cite a thread it never sees, which the conductor correctly refuses
> as unprovable lineage. Both were master drafting errors, not frank defects — see
> `FRANK-HARDENING-BACKLOG.md`.


**STATUS: INERT.** This brief grants **no lane action** — no seat mint/boot, no preflight, no authoring, no dispatch authority. It becomes operative **only** at the operator's post-preflight activation (§9). It **supersedes** the voided pair-shaped draft `master/STEP-3-LANE4-KICKOFF.md` (kept as history).

**Authority of record (do not touch):**
- Approved plan `master/STEP-3-LANE4-PLAN.md` **rev7 `e7a333e9c4c5e34cb62dffa29c0b37f03d48022a233636a0d0c34b28006994d2`** (VP-approved `step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-110000.md`) + `GRILL_LOCK step3-lane4-staffing-grill-1` + deviation **B21**. **Where this brief and the plan differ, the plan governs.**
- Frozen exam: `master/STEP-3-STAGE6-AMENDMENT.md` **§7** — six legs, ten fixture records, the machine predicates, the exact manifest schema + fixed values, the sample-weight budget. **Author TO §7; do NOT paraphrase or re-derive it.**
- Interface lock (byte-stable, immovable): `master/STEP-3-INTERFACE-LOCK.md` `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

---

## 1. The team (canonical role-stamped addresses; the `l4.*` namespace is the operator's)
- **`l4.orchestrator-planner`** (`ROLE: Orchestrator Planner`) — decomposes the ten records, **dispatches** the (operator-booted) worker seats, **integrates** their proposals into the single manifest, owns the 30-turn/100-call budget + carried-obligation consistency.
- **`l4.orchestrator-reviewer`** (`ROLE: Orchestrator Reviewer`) — the durable **decomposition-review gate** (§8) + the two read-only verification duties (byte-equality; content review) with approve/revise verdicts.
- **`l4.w<k>.planner`** (`ROLE: Planner`) / **`l4.w<k>.implementer`** (`ROLE: Implementer`) — worker seats authoring individual fixture scenarios/baselines; a single-seat worker is `ROLE: Planner`. (The `.implementer` suffix stamps `ROLE: Implementer` — the relay-lint tripwire.)

**How the seats exist (B21 / VP-r7-F4 — honest current-generation mechanism):** native governed spawn is Step-4-deferred. So the **operator mints (`seat_mint`) + boots each seat as an independent session** (not an l4 spawn, not a subagent) and wires its frank credential; `l4.orchestrator-planner` **dispatches** work to the already-booted seats **by frank relay**. **frank = courier + seat-identity carrier, NOT a spawn engine.**

## 2. Mission
Author the **frozen exit-test oracle** — the **ten fixture records across the six §7 property legs** — as a content-addressed test spec (the answer key). Runnable fixtures + frank code are **T4**, not here.

## 3. Deliverable — ten records to the EXACT §7 schema
Author each record's **immutable input scenario + baseline**, and the draft `STEP-3-EXIT-FIXTURES.json`, **to §7's frozen predicate + schema**. Per-record base keys `{ fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator }`; the typed expectation fields + **fixed values** bound to their records — `effect_observer_key`+`effect_counter_expectation{counter_before_recovery:1,counter_after_recovery:1,invocations_after_recovery:0}`→`xit-crash-1`; `handoff_expected_records[2]`→`xit-ho-1`; `resume_prefix_expectation{predecessor_turn_id,resumed_round_index,log_prefix_digest,context_digest}`→`xit-dur-1`; `degraded_expectation{corruption_cut,expected_disposition:"degraded",expected_resume_action}`→`xit-dur-2`; per-record `sample_weight` summing to **exactly 30 governed turns + 100 tool calls**; top-level `{baseline_artifact_digest,baseline_config_digest}`. Records: `xit-gov-1`, `xit-dur-1..5`, `xit-crash-1`, `xit-inj-1`, `xit-ho-1`, `xit-op-1`.

## 4. Carried obligations (realize as fixture content, don't just cite — plan §5)
1. **N910:** honest `UNKNOWN_PROVIDER_OUTCOME`→`uncertain` in `xit-dur-3`/`xit-op-1`; never author rows implying complete lane-2 coverage.
2. **`env_digest` parity:** under `xit-gov-1` — m-1's byte-exact JCS preimage + duplicate-name reject + reachable non-UTF-8 reject + **m-9↔m-3 observer parity**.
3. **r7-mirror:** at `xit-gov-1` E3-predicate authoring, the **mandatory m-3 check** whether independent m-10-side 2a/2b resolution is needed — **if YES: STOP + escalate to reopen (route-now)**; else record the non-gating deferred disposition.

## 5. Owner-fidelity (guiding PM = m-3; the plan §5 ten-record matrix, out-of-team, before Master+VP freeze)
Each record's producer/observer boundaries file fidelity outside the l4 team: `xit-gov-1` (m-8·m-10·m-9·m-7·m-1·m-3); durabilities (m-9·m-10·m-3, +m-8 on dur-3, +m-7 selected-conductor-action on dur-4); `xit-crash-1` (m-9·m-10·m-3 + fixture external observer); `xit-inj-1` (m-10·m-9·m-7·m-1·m-3); `xit-ho-1` (m-7·m-1·m-2·m-3); `xit-op-1` (m-8·m-10·m-9·m-3). Exact rows: plan §5.

## 6. Delivery mechanics (read-only → frank → master materializes — plan §3)
- **READ-ONLY** on the whole workspace; **no seat writes into `master/`.** Every proposed file — each input artifact, each baseline, **and the complete final `STEP-3-EXIT-FIXTURES.json`** — rides a proposal envelope in the frank relay body: `{ path, encoding ("utf-8"|"base64"), byte_length, sha256, content }`.
- **Master alone materializes** + recomputes on-disk length + SHA-256; **`l4.orchestrator-reviewer` confirms byte-equality (duty i)**; **then its content review (duty ii, approve/revise)**; then out-of-team owner-fidelity; then VP; then **Master+VP freeze/re-lock** (NOT the team's).
- **Frame fit:** each envelope must fit frank's `max_frame_bytes` (1 MiB default). Oversized → **escalate** (master defines a deterministic chunk/archive contract + reassembly proof); never truncate or hand-copy.
- **No** placeholder/mutable slot/unresolved owner/arithmetic-only weight reaches the manifest proposed for freeze.

## 7. Nested lineage — concrete dispatch-id map (r7-F2 instantiation guard: distinct concrete ids per gated leg; `PARENT_DISPATCH_ID` = the immediate predecessor; NEVER reuse the root `step3-relock-lane4`)

| gated leg | DISPATCH_ID (concrete) | PARENT_DISPATCH_ID (immediate predecessor) |
|---|---|---|
| master → l4 boot/preflight assignment | `step3-relock-lane4-l4-boot` | the master kickoff/assignment relay that opens it |
| l4 decomposition → l4-reviewer verdict | `step3-relock-lane4-l4-decomp` | `step3-relock-lane4-l4-boot` (thread parent) |
| worker `w<k>` dispatch → return | `step3-relock-lane4-l4-w<k>` (e.g. `-l4-w1`, `-l4-w2`, …) | `step3-relock-lane4-l4-decomp` (the approving decomposition thread) |
| l4 integration request → final content verdict | `step3-relock-lane4-l4-integrate` | `step3-relock-lane4-l4-w<k>` (the worker returns it integrates) |
| escalation `n` → disposition → return | `step3-relock-lane4-l4-esc<n>` (e.g. `-l4-esc1`) | the thread it escalates from |

**One gated leg = one dispatch id.** Within a leg's thread (decomposition→verdict, dispatch→return, request→verdict, escalation→disposition), the **response relay orders after its request by `IN_REPLY_TO` = the exact predecessor relay file** — NOT a second dispatch id (a request and its verdict sharing a dispatch id is the normal thread shape; the resolver defect below arises only when something resolves a shared id expecting uniqueness — the unique-per-*leg* ids here prevent that). `PARENT_DISPATCH_ID` names the immediate-predecessor **thread**; `IN_REPLY_TO` names the immediate-predecessor **relay**. Tier ancestry lives in the id namespace (`step3-relock-lane4` → `…-l4-*` → `…-l4-w<k>`), a **unique sub-dispatch id per gated leg** (avoids the shared-dispatch-id resolver defect, `CYCLE-PLAYBOOK.md:139-164`); the root `step3-relock-lane4` is never reused as a concrete leg id.

## 8. The decomposition-review gate (r7-F3 — before any worker authors)
On activation: (a) **`l4.orchestrator-planner`** files the **decomposition** (`…-l4-decomp`): worker topology, per-worker fence, artifact ownership, cross-record budget allocation, carried-obligation allocation, escalation rules; (b) **`l4.orchestrator-reviewer`** files a **durable approve/revise verdict**; (c) **only an approve permits addressed worker dispatch** (a revise returns to (a)). Byte-equality + content-review duties (§6) are separate/later.

## 9. Gates + sequence (you enter only at activation)
0. This brief written INERT.
1. **Operator** mints + boots `l4.orchestrator-planner`, `l4.orchestrator-reviewer`, ≥1 probe worker as independent read-only sessions (**zero authoring/dispatch authority**) + wires frank creds; run a real **three-tier** exchange (both orchestrator seats + ≥1 worker; accepted `submit`/`read` + correct immediate-predecessor `PARENT` chaining per §7 + durable export). The probe carries **no** fixture-authoring authority.
2. **On pass → operator activation** turns this brief operative (on fail → hold; hand-relay = operator-owned deviation).
3. **Decomposition gate (§8):** planner decomposition → reviewer approve → worker dispatch.
4. Workers **author + content-address** the ten input/baseline artifacts (proposal envelopes); `l4.orchestrator-planner` **integrates**; master materializes; `l4.orchestrator-reviewer` byte-equality.
5. `l4.orchestrator-planner` proposes the **complete manifest** → master materializes + recomputes → reviewer byte-equality → **reviewer CONTENT review (approve/revise)**.
6. On content-approve → **out-of-team owner-fidelity (§5)** files → **VP review**.
7. **Master+VP freeze `STEP-3-EXIT-FIXTURES.json`** + one durable re-lock binding `cbd1893c…` + the frozen manifest SHA. **T4 opens only after the re-lock + H-16 + H-26.**

## 10. Escalation + gap-logging (three-tier)
Escalate `worker → l4.orchestrator-planner/reviewer → master → owning m-x planner` on any `DELEGATED_DISPATCH_AUTHORITY` trigger (§7 spec mistake · better way · locked-contract touch · cross-domain collision · the r7-mirror STOP). Each escalation relay is a `…-l4-esc<n>` leg; the disposition parents to it. Every "the artifacts didn't answer" escalation is a design-of-record gap — **master** appends the battle report to `master/FRANK-HARDENING-BACKLOG.md` (you surface; master appends).

## 11. Prior-art provenance (independent arrival — PRIOR-ART §2c; per the build-reference mandate §4)
The nested parent→child authority model + the **monotone-non-increasing ceiling** were designed before these were read; three shipped benchmarks now corroborate the shape: **codex** role-scoped spawn (`multi_agent_v2` Stable — agent roles control the child type) · **paperclip** audited parent-only one-hop exception (non-transitive) · **omp** workspace-root/AGENTS.md inheritance + the **quiescence barrier** (a worker's final yield with pending jobs = a PAUSE, not completion — a lifecycle trap once async tool execution exists; noted for the escalation rules). Independent arrival is evidence the shape is right; frank's ceiling stays stricter (author-only, no cross-hop).

## 12. Boundaries + done
READ-ONLY; write nothing into `master/`. No freeze, re-lock, lock/contract/owner-byte write, `frank/` code, T4, credential, provider call, live E3, merge, deploy. **H-12 hard-blocks external use.** The interface lock `cbd1893c…` + all §7/owner/frozen bytes are immovable — if your work seems to require moving one, **STOP + escalate.** **Done (your half):** a materialized, reviewer-content-approved set of ten input/baseline artifacts + a manifest matching §7 exactly (schema + fixed values + carried rows + resolved observers/locators + weights summing to 30 turns/100 calls, no placeholder). Freeze + re-lock are **Master+VP**.
