# Sprint s6 — Slice-6: THE TRANSPORT FIX (the last Step-1 slice)

**RUN_ID:** s6 · **Baseline:** `main @ 7e5c527` (tag `s5-close`), battery 23-ok · **Branch (proposed):** `s6-transport-impl`
**Work dispatch of record:** `../../../../master/relays/s6-dispatch/PLAN-orchestrator-planner-20260706-221040.md`
**Spec (read-only, design-of-record):** `master/S6-AMENDMENT-SET-2026-07-06.md` (r3, VP co-signed `s6-design/RECONCILE-orchestrator-reviewer-20260706-220325`) + its four constituents — m-1 `2026-07-06-s6-transport-amendments.md` (§A–§F.1), m-7 `2026-07-06-s6-transport-amendments.md` (r5), m-2 `2026-07-06-s6-transport-codec-amendment.md` (+§11), `master/GRILL-LOCK-parenting-fork-2026-07-06.md`.
**The story:** `master/TRANSPORT-FINDINGS-2026-07-06.md` (F1–F17; evidence archive `~/frank-archives/frank-team-store-s5-dogfood-20260706`).

## The one-line boundary
Build the co-signed amendment set — whole, nothing else — against the LOCKED amendments. Seat surface stays exactly `submit`/`project`/`read` (roster + audit views are `project` parameters). m-7 guides; m-1 fidelity on store/lineage/waiver/lock/activation; m-2 fidelity on codec/registry/boot-form/render/validate. The set rules; escalate to amend — never improvise a locked-contract change.

## IN (the set, whole — the set's §"Build-slice obligations" is authoritative)
Branch-A parenting (conductor-computed PARENT + `parent_hint` fallback-never-bounce per the GRILL_LOCK) · the ONE canonical `address_list` codec · A-1 stable-schema digest · A-2 idempotent-replay intake + durable monotonic ids · A-3 live mint (`seat_mint` pivot) · A-4+§D store lock (I1-P, proof-of-death takeover, refuse-reads loser) · §B `project()` default-accepted + accepted-graph anchor fix · §C scoped waivers + `waiver_retraction` · F13 three-layer record_kind authorization · D-1 shim transparent-reconnect · D-2 bounce/reply detail parity · B-1 lifecycle `minted→bound→active` + literal admission allowlist + roster view · B-2 boot form (SITREP + lifecycle-gated) · B-3 derived-only activation ([VP-W3]: registry pass = EXACTLY the seven named rows, NO activation-marker row) · the registry pass (seven rows + `waiver_retraction`/`seat_mint` classes + `ORCH_REVIEW_WAIVER` `"*"` header retired; MINOR, no envelope migrator — R-1).

## OUT (escalate before touching)
Step-2 observe · routing execution (Step-3) · **engine performance work of ANY kind** (latency exonerated by measurement) · new seat verbs · federation · any locked-contract change outside the co-signed set.

## Exit gate (HARD; Step-1 closes on it)
1. Fixtures red-first all green: the GRILL_LOCK three (archived-dogfood F11 replay without livelock · hint honored/fallback · concurrent-accept-no-parent-bounce) · m-7 FX-A1a..FX-B1g (all 18; FX-B1g = [VP-W2]) · m-1 §E set · m-2's set.
2. E2 floors: full battery green (s1–s5 suites + s6's), uncached, zero regression; byte-exact enum; three-verb surface; I-PH over every NEW surface (roster, boot bounces, lock refusals, hint flags).
3. THE STEP-EXIT TEST (live, on the fixed conductor): fresh blessed store; live seats (ops: pre-allowlist `mcp__frank__*`); (a) ROADMAP:83-85 legs; (b) first live act = operator §7-applies s5's registry (operator-authored); (c) F11 regression leg — the archived dogfood traffic replayed live without livelock; (d) live boot: mint (A-3, no restart) → wire → B-2 boot → `active` derived → roster shows it.
4. Honesty: transport/provenance-only phrasing; done-state + `record_integrity` stay `self_reported`; ③/⑤ dormant claims unchanged; no claim over the §C4.3 ceiling.

## Gate + escalation
F2 (non-bootstrap): pair-Implementer plan-review = the plan gate; delegated `DISPATCH IMPL` under {approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract amendment}. Any failure or OUT-touch escalates to master (CTO + m-7 guide + VP). All m-x questions route through master as file relays via the operator. [VP-W1]: if m-7 exercises its one-relay revert window on the two master-applied stale-text corrections, master holds and reconciles first.
