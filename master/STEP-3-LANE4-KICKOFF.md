# Step-3 Lane-4 — detailed kickoff brief — ⛔ VOID / SUPERSEDED (pair-shaped, pre-B21)

> **⛔ VOID — DO NOT USE.** This draft was authored against the **pair-shaped** plan rev5 (`faa23c7f…`). It was superseded by the operator's B21 nested-team directive; B21 has since been **fully superseded for lane 4** (**B22** stood down the courier, **B23** set the team shape back to **pair**). So the shape this draft assumed is once again the authorized one — but **these bytes are still VOID**: they are not the rev9 successor, they were written against rev5 rather than rev9's write fence and unique-id lineage, and they grant **no lane action**. The operative kickoff will be authored fresh on VP approval of plan rev9.

**STATUS: VOID (superseded pre-use; INERT).** This brief grants **no lane action** and never did — it is superseded before ever being operative.

**Authority of record:** the approved lane-4 plan `master/STEP-3-LANE4-PLAN.md` **rev5 `faa23c7f6486a19e6960d2bb5e3a72a3f1e9b08cf4c33dee22ce1ad171e3679b`** (VP-approved `step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-070000.md`) + `GRILL_LOCK step3-lane4-staffing-grill-1`. Where this brief and the plan differ, **the plan governs** (this brief adds concrete marching orders, it does not re-decide).

**Spec of record (the frozen exam):** `master/STEP-3-STAGE6-AMENDMENT.md` **§7** — the six legs, the machine predicates, the manifest schema, and the sample-weight budget are **ratified and frozen; author TO them, do not paraphrase or re-derive them.**

**Interface lock (byte-stable, do not touch):** `master/STEP-3-INTERFACE-LOCK.md` `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

---

## 1. The team
- **`l4.planner`** — the **proposal/content author**: designs the ten fixture-input scenarios + baselines + the manifest, emits them as proposal envelopes.
- **`l4.implementer`** — two distinct read-only duties: **(i) proposal-to-file byte-equality** for every materialized file incl. the manifest (+ any chunk/archive reassembly); **(ii) an independent adversarial CONTENT review** of the full materialized set + manifest, with a **durable approve/revise verdict**. Byte-equality ≠ content review; both required.
- Both seats are **READ-ONLY on the entire workspace** — full visibility, write nothing into `master/`.

## 2. Mission
Author the **frozen exit-test oracle** — the **ten fixture records across the six §7 property legs** — as a content-addressed test specification. This is the **answer key**, not runnable tests: you author the input scenarios, baselines, and the manifest carrying the expected canonical outcomes; the **runnable fixtures + frank code are built at T4** (not here).

## 3. Deliverable — ten records to the EXACT §7 schema
Author, for each of the ten records, its **immutable input-artifact scenario + baseline**, and assemble the draft `STEP-3-EXIT-FIXTURES.json`. **Do not invent the predicate — author the scenario that makes §7's predicate for that record decidable.**

**Per-record base keys (exact §7):** `{ fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator }`.

**Typed expectation fields (exact §7 shapes + fixed values), each bound to its record:**
- `xit-crash-1` → `effect_observer_key` + `effect_counter_expectation { counter_before_recovery: 1, counter_after_recovery: 1, invocations_after_recovery: 0 }`
- `xit-ho-1` → `handoff_expected_records[2]`
- `xit-dur-1` → `resume_prefix_expectation { predecessor_turn_id, resumed_round_index, log_prefix_digest, context_digest }`
- `xit-dur-2` → `degraded_expectation { corruption_cut, expected_disposition: "degraded", expected_resume_action }`

**Per-record `sample_weight`** whose values sum to **EXACTLY 30 governed turns and 100 tool calls** (aggregate totals, NOT per-record). **Top-level `{ baseline_artifact_digest, baseline_config_digest }`.**

**The ten records + their §7 leg + fidelity owners (author each TO its §7 predicate; owners per the plan §5 matrix):**

| record | §7 leg (author to its frozen predicate) | out-of-pair owner-fidelity |
|---|---|---|
| `xit-gov-1` | Governance-binding | m-8 · m-10 · m-9 · m-7 · m-1 · m-3 |
| `xit-dur-1` | Durability — resume | m-9 · m-10 · m-3 |
| `xit-dur-2` | Durability — degrade | m-9 · m-10 · m-3 |
| `xit-dur-3` | Durability — provider conjunction / N910 | m-8 · m-9 · m-10 · m-3 |
| `xit-dur-4` | Durability — receipt gate | m-9 · m-10 · m-7 (selected conductor-action obs; provider/tool branch → m-8/tool owner) · m-3 |
| `xit-dur-5` | Durability — frame overflow | m-9 · m-10 · m-3 |
| `xit-crash-1` | Crash-honesty | m-9 · m-10 · m-3 + fixture-owned external observer |
| `xit-inj-1` | Injection-visibility | m-10 · m-9 · m-7 · m-1 · m-3 |
| `xit-ho-1` | Governed handoff | m-7 · m-1 · m-2 · m-3 |
| `xit-op-1` | Operability | m-8 · m-10 · m-9 · m-3 |

Guiding PM = **m-3** (evidence / exit-gate owner). Consult it first on any §7-predicate question.

## 4. Carried obligations (closed checklist — realize each as fixture content, don't just cite)
1. **N910 (documented-MVP-limit):** the honest `UNKNOWN_PROVIDER_OUTCOME` → `uncertain` disposition must appear in the `xit-dur-3` / `xit-op-1` expected rows; **never** author expected rows that imply complete lane-2 coverage. (owner: m-3 / m-10 / m-8)
2. **`env_digest` preimage parity:** under `xit-gov-1`, the expected rows must exercise m-1's byte-exact JCS preimage recipe, duplicate-name rejection, reachable non-UTF-8 typed rejection, and **m-9↔m-3 observer parity**. (owner: m-1, realized by m-9 + m-3)
3. **r7-mirror (deferred-v3):** at E3-predicate authoring for `xit-gov-1`, perform the **mandatory m-3 check** whether it needs independent m-10-side 2a/2b resolution — **if YES: STOP and escalate to reopen (route-now)**; otherwise record the non-gating deferred disposition. (owner: m-3)

## 5. How you deliver (read-only → frank → master materializes)
- Every proposed file — each input artifact, each baseline, **and the complete final `STEP-3-EXIT-FIXTURES.json`** — rides a **proposal envelope** in the frank relay body: `{ path, encoding ("utf-8"|"base64"), byte_length, sha256, content }`. You write **nothing** into `master/`.
- **Master alone materializes** the exact bytes and recomputes the on-disk length + SHA-256; **`l4.implementer` confirms proposal-to-file byte-equality**; then the **`l4.implementer` content review** (approve/revise); then owner-fidelity + VP; then Master+VP freeze.
- **Frame fit:** each encoded envelope must fit frank's live `max_frame_bytes` (1 MiB default) with overhead. If an artifact won't fit, **do not truncate or hand-copy — escalate**; master defines a deterministic content-addressed chunk/archive contract (with a reassembled-byte proof the implementer verifies) before you author that artifact.
- **No placeholder, mutable slot, unresolved owner, or arithmetic-only weight** may reach the manifest you propose for freeze.

## 6. Escalation + gap-logging
- Escalate **up through master to the owning m-x planner** on any `DELEGATED_DISPATCH_AUTHORITY` trigger: a §7-spec mistake, a better way found while authoring, a locked-contract touch, a cross-domain collision, or the r7-mirror STOP.
- Every escalation caused by **"the artifacts didn't answer"** is a design-of-record gap. Raise it explicitly; **master** appends the battle report to `master/FRANK-HARDENING-BACKLOG.md` (master-owned; you surface, master appends).

## 7. Gates + sequence (you enter only at activation)
0. This brief is written INERT.
1. **Operator authorizes a preflight-only boot** of the pair + full frank roster — **zero authoring/dispatch authority**; run the real up-and-back round-trip (team → master → team) + a durable relay export.
2. **On pass → operator activation** turns this brief operative. On fail → hold (hand-relay fallback = operator-owned B13 deviation).
3. You **author** the ten input/baseline artifacts (proposal envelopes) → master materializes → `l4.implementer` byte-equality.
4. You assemble + propose the **complete manifest** → master materializes + recomputes → `l4.implementer` byte-equality → **`l4.implementer` content review (approve/revise)**.
5. On implementer content-approval → **out-of-pair owner-fidelity (§3 matrix) + VP review** on the materialized bytes.
6. **Master+VP freeze `STEP-3-EXIT-FIXTURES.json`** + one durable re-lock binding `cbd1893c…` + the frozen manifest SHA. **T4 opens only after that re-lock + H-16 + H-26.**

## 8. Definition of done (your half)
A complete, materialized, implementer-content-approved set of ten input/baseline artifacts + a manifest that: matches the exact §7 schema + fixed values; carries the three carried-obligation rows; resolves every `observer_id`/`evidence_locator`; and whose `sample_weight`s sum to exactly 30 governed turns + 100 tool calls — with no placeholder. Freeze + re-lock are **Master+VP**, not yours.

## 9. Boundaries
Read-only workspace; write nothing into `master/`. No freeze, no re-lock, no lock/contract/owner-byte write, no `frank/` code, no T4, no credential, no provider call, no live E3, no merge, no deploy. **H-12 hard-blocks external use.** The interface lock `cbd1893c…` and all §7/owner/frozen bytes are immovable; if your work seems to require moving one, **STOP and escalate.**
