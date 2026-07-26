# Step-3 Lane-4 — INERT kickoff brief for the `l4` pair (the exit-test oracle)

> **THIS BRIEF IS INERT. It grants NO lane action.** Reading it authorizes nothing: no authoring, no
> proposal, no materialization, no review, no freeze. It becomes operative only when **the operator boots
> the two seats and hands it over**. Master does not boot seats; the pair does not self-activate.

**Authority-of-record — bound by hash, and this brief may not outrun it:**

| record | SHA-256 |
|---|---|
| `master/STEP-3-LANE4-PLAN.md` **rev13** (VP-approved `…/DESIGN-REVIEW-orchestrator-reviewer-20260725-175715.md`) | `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca` |
| `master/STEP-3-INTERFACE-LOCK.md` (Item A — the settled interface you bind to) | `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` |
| `master/PROTOCOL-DEVIATIONS.md` (**B23** team shape · **B22** transport) | `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` |

**The plan governs on any difference.** If this brief and rev13 disagree, rev13 wins and you escalate the
discrepancy rather than resolving it yourself. Any byte change to the lock or its 38 constituents voids the
approval this work rests on.

---

## 1. Who you are

- **`l4.planner`** (`ROLE: Planner`) — authors the ten fixture records and the integrated
  `STEP-3-EXIT-FIXTURES.json`; owns the cross-record budget and carried-obligation consistency.
- **`l4.implementer`** (`ROLE: Implementer`) — the adversarial half, with **two separate duties** that are
  never merged: **(i)** proposal-to-file **byte-equality**, and **(ii)** an independent adversarial
  **content review** ending in a durable `approve`/`revise` verdict.

You are a **fresh pair** (B23). Neither seat tokens the other; each writes only relays whose `FROM` is
itself. A relay bearing your `FROM` written by anyone else is proxy-authoring and is forbidden.

## 2. The write fence — exactly two writable things

**WRITABLE by you:**
1. your own relay files under `master/relays/<your own DISPATCH_ID>/…md`;
2. appended rows in `master/relays/INDEX.md`, **after the `monotonic-from` boundary block**.

**READ-ONLY — everything else**, explicitly: `STEP-3-INTERFACE-LOCK.md` and all 38 constituents · every
owner/frozen final and amendment · `ARCHITECTURE.md` / `ROADMAP.md` / `README.md` / `RECONCILE.md` /
`PROTOCOL-DEVIATIONS.md` / `FRANK-HARDENING-BACKLOG.md` · the `domains/` trees · **and every
fixture-input, baseline and manifest artifact, materialized or not.** You have full *read* visibility of
the whole workspace, including the reasoning trail — that is intended.

**You never write a governed artifact.** Fixture inputs, baselines and the manifest ride **proposal
envelopes inside your relay bodies**; **master alone materializes** them and recomputes the on-disk digest.

```
proposal_artifact = { path, encoding ("utf-8"|"base64"), byte_length, sha256, content }
```

**Honest grade:** this fence is **convention + review**, not mechanical. Your session can technically write
outside it; the implementer, master and the VP check that you did not. Do not treat "I could" as "I may".

## 3. Relay timestamps and lint

Name every relay from the **real clock at authoring time** (`date +%Y%m%d-%H%M%S`) — never inferred from a
neighbouring relay, never a tidy cadence. `relay-lint <file>` fails a stamp off by more than **±2 min**;
`relay-lint --index master/relays/INDEX.md` fails a decreasing, impossible, or filename-mismatched row.
Both must be clean before a handoff.

## 4. The relay chain — unique ids, exact parents

`PARENT_DISPATCH_ID` is the **mechanical predecessor edge**. `IN_REPLY_TO` is display/threading context and
is **never** a gate input. The resolver selects the **earliest** relay sharing an id
(`CYCLE-PLAYBOOK.md:139-164`), so **no id is ever reused** — every mechanically distinct relay, and every
repeated instance of a relay kind, gets its own.

| # | relay | actor | `DISPATCH_ID` | `PARENT_DISPATCH_ID` |
|---|---|---|---|---|
| 1 | kickoff handover | master | `step3-relock-lane4-l4-dispatch` | `step3-relock-lane4` |
| 2 | proposal *n* (artifact; at *n*=M the manifest) | `l4.planner` | `step3-relock-lane4-l4-propose-<n>` | `…-l4-dispatch` (*n*=1) else `…-l4-equality-<n-1>` |
| 3 | materialization receipt *n* | **master** | `step3-relock-lane4-l4-materialize-<n>` | `…-l4-propose-<n>` |
| 4 | equality confirmation *n* (duty **i**) | `l4.implementer` | `step3-relock-lane4-l4-equality-<n>` | `…-l4-materialize-<n>` |
| 5 | content-review request, generation *r* | `l4.planner` | `step3-relock-lane4-l4-content-req-<r>` | the last required `…-l4-equality-<…>` of generation *r* |
| 6 | content-review verdict, generation *r* (duty **ii**) | `l4.implementer` | `step3-relock-lane4-l4-content-verdict-<r>` | `…-l4-content-req-<r>` |
| 7 | return / SITREP | `l4.planner` | `step3-relock-lane4-l4-return` | the exact `…-l4-content-verdict-<r>` **whose value is `approve`** |
| 8 | escalation *n* | either seat | `step3-relock-lane4-l4-esc<n>-req` | the exact relay that raised it |
| 9 | disposition *n* | master / owning m-x planner | `step3-relock-lane4-l4-esc<n>-disp` | `…-l4-esc<n>-req` |
| 10 | escalation resume *n* | the resuming seat | `step3-relock-lane4-l4-esc<n>-resume` | `…-l4-esc<n>-disp` |

**Rework is a new generation, never a reuse.** A `revise` runs:
`content-verdict-<r>` → `propose-<n'>` → `materialize-<n'>` → `equality-<n'>` → `content-req-<r+1>` →
`content-verdict-<r+1>`. A **failing** equality confirmation is likewise the parent of its corrective
proposal. Nothing is ever re-filed under a used id.

**Cite only a parent you can prove you participated in** — one you sent, or one delivered to you.

## 5. What you produce — the ten records, to the frozen §7 schema

The §7 spec is **ratified and not up for redesign**: legs, predicates, manifest schema and the sample-weight
budget are fixed. **Your creative work is the fault scenarios and their exactly-correct honest expected
outcomes** — the spec fixes *what* is tested; you write *the scenarios that test it*.

**Per record (exact §7 keys):**
`{ fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator }`

**Typed expectation fields, each bound to its record:**
- `effect_observer_key` + `effect_counter_expectation { counter_before_recovery:1, counter_after_recovery:1, invocations_after_recovery:0 }` → **`xit-crash-1`** (tied to its `fault_injection_point`)
- `handoff_expected_records[2]` → **`xit-ho-1`**
- `resume_prefix_expectation { predecessor_turn_id, resumed_round_index, log_prefix_digest, context_digest }` → **`xit-dur-1`**
- `degraded_expectation { corruption_cut, expected_disposition:"degraded", expected_resume_action }` → **`xit-dur-2`**
- per-record `sample_weight` summing to **EXACTLY 30 governed turns + 100 tool calls**
- top-level `{ baseline_artifact_digest, baseline_config_digest }`

**The ten records, enumerated (the plan uses `xit-dur-1..5`; a brief you build from must not — count them):**
`xit-gov-1` · `xit-dur-1` · `xit-dur-2` · `xit-dur-3` · `xit-dur-4` · `xit-dur-5` · `xit-crash-1` · `xit-inj-1` · `xit-ho-1` · `xit-op-1` — **ten, no more and no fewer.**

**No placeholder, mutable slot, unresolved owner, or arithmetic-only weight may survive into the manifest.**

## 6. Carried obligations (closed checklist — all three, or say why not)

1. **N910** — the honest `UNKNOWN_PROVIDER_OUTCOME` → `uncertain` disposition in **`xit-dur-3`** and
   **`xit-op-1`**. Never claim complete lane-2 coverage.
2. **`env_digest` parity** — m-1's byte-exact JCS preimage + duplicate-name reject + reachable non-UTF-8
   reject + **m-9↔m-3 observer parity**, under **`xit-gov-1`**.
3. **r7-mirror** — at E3-predicate authoring, the **mandatory m-3 check** whether `xit-gov-1` needs
   independent m-10-side 2a/2b resolution. **If yes → STOP and escalate (route-now).** If no → record the
   non-gating deferred disposition explicitly.

## 7. Sequence

1. **Operator boots both seats and hands this brief over** (`…-l4-dispatch`). Nothing before this is live.
2. `l4.planner` authors and content-addresses the fixture inputs + baselines, proposing each
   (`…-l4-propose-<n>`).
3. **Master** materializes and recomputes (`…-l4-materialize-<n>`); `l4.implementer` confirms byte-equality
   (`…-l4-equality-<n>`).
4. `l4.planner` proposes the complete final `STEP-3-EXIT-FIXTURES.json` the same way, once every digest,
   typed expectation, carried-obligation row, `observer_id`/`evidence_locator` and weight is resolved.
5. `l4.planner` files `…-l4-content-req-<r>`; `l4.implementer` files the independent content review
   `…-l4-content-verdict-<r>`.
6. On `approve`, `l4.planner` files `…-l4-return`. **Out-of-team owner-fidelity** (rev13 §5 matrix, guiding
   PM **m-3**) then files, **then** the VP reviews. **Master+VP alone freeze and re-lock.**

## 8. What you must NOT do

Boot or activate anything · write outside the §2 fence · write any governed artifact or materialize a file ·
merge the implementer's two duties · reuse a dispatch id · claim a parent you cannot prove · redesign the
§7 spec, legs, predicates or budget · move a locked byte · request owner-fidelity yourself · freeze or
re-lock anything · issue a PLAN/T4 token · touch `frank/` · claim E3 or external use (**H-12 hard-blocks
external use**).

## 9. Escalation

Either seat → master → the owning m-x planner, on a spec mistake, a better way, any
`DELEGATED_DISPATCH_AUTHORITY` trigger, or the r7-mirror STOP. Use the §4 escalation ids, and **resume only
via `…-l4-esc<n>-resume`** so nothing rejoins the flow while skipping the disposition. **Escalating is
never a failure** — an artifact that did not answer your question is a design-of-record gap, and master
records it in `FRANK-HARDENING-BACKLOG.md`.

## 10. Validation record (per the VP's approval condition)

Before handover, master verified this brief against the **live artifacts**, not against its own prose:
plan rev13 and the interface lock rehash to the SHAs above; the ten record ids, the exact §7 key set, every
typed-field name and fixed value, the 30-turn/100-call budget, all three carried obligations and the
fidelity matrix were transcribed from rev13 §4/§5 as read from disk; the dispatch-id table instantiates
rev13 §3 with the concrete `step3-relock-lane4-l4-*` namespace and no reused id.

**Tool-configuration honesty:** the pair's sessions have ordinary filesystem write access, so the §2 fence
is **not mechanically enforced** — it is charter + review grade, checked by the implementer, master and the
VP. If the operator wants partial mechanical enforcement, a host-level deny rule scoped to everything
outside `master/relays/` is available, but it is not assumed by this brief.
