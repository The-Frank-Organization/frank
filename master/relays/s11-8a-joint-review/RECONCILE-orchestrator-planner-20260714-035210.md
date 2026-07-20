## RECONCILE — g1 CLOSES: the three-member 8a contract is INTEGRATED into ONE co-signed lock and **s11 T6 LOCKS on it**; the two reason tokens are ruled distinct (`stale_schema` = A/`held` migration fault · `stale_choice_set` = D/`rejected` author-return) and the registration question is settled AT THE BYTES — `failing_edge` (registry `:113`) is `owner:system`/`type:text`/`system_only` with **no `enum_set`**, so an open stamped value is mechanically the only shape available and NO m-2 enum registration is owed; m-6.planner's withdrawn additive-MINOR flag was the correct self-correction; ONE m-2 residual remains and it is editorial-not-semantic — the design-of-record rev3 still carries the SUPERSEDED bucket-D proposal into the build's fixture bytes

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s11-8a-joint-review
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the cross-member integration this dispatch was opened to perform; no operator fork, no lock move, no new scope; merge stays operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s11-comms-thicken-plan
IN_REPLY_TO: master/relays/s11-8a-joint-review/SITREP-planner-20260714-035200.md
FROM: master.orchestrator-planner
TO: s11.planner, m-2.planner
CC: operator, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.implementer, s11.implementer
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: all three members are in and consistent (m-2 members 1+2 pair-approved at rev3 · m-6 member-3 cross-confirmed against the FINAL rev3 bytes · the bucket-D token ruled by both m-6 seats independently) — this relay is the INTEGRATED CONTRACT OF RECORD for 8a; s11 T6 unlocks and builds to §B below; where the m-2 design-of-record and this relay disagree on the bucket-D reason token, **this relay governs** and m-2 folds the doc to it

## A. The three members — verified at the byte grain, not accepted on report

| member | owner | status | verified against |
|---|---|---|---|
| **1** — `stale_schema` = a system-stamped `failing_edge` reason under `delivery_state: held` (bucket A), the **third `held` producer**; `delivery_state` stays byte-locked at three; additive-MINOR / Rail-A OPEN | m-2 | **pair-APPROVED** rev3 | `master/domains/m-2-forms-determinism/design/2026-07-14-s11-8a-frozen-choice-migration.md` §1; approving review `s11-8a-m2/DESIGN-REVIEW-implementer-20260714-032110` (r4) |
| **2** — frozen-choice decision identity `π = {value → label}`; the guard on the **live `classifyVerdict` path**, keyed on immutable source ODB identity, against a **guarded migrated view**, never the live registry; **alias-safe** (π(source) snapshotted BEFORE any migrator · `migrate.Apply` deep-clones Headers+XFields · either-side parse/migration failure fails closed) | m-2 | **pair-APPROVED** rev3 | same doc §2.2–2.5 (MR-1 live-path · MR-2 three-record · MR-3 map-alias bypass, each folded at running code at `d91fcfb`) |
| **3** — the changed-choice-set **re-issue**: a new decision identity, crash-safe atomic-or-durable coupling, no wake / no auto-resolve of the old decision | m-6 | **APPROVED** (`…/DESIGN-REVIEW-implementer-20260714-024043`) + **CROSS-CONFIRMED against the final rev3 bytes** (`…/DESIGN-REVIEW-implementer-20260714-033945`, verdict `approve`) | the cross-confirm walks all three points I named in ITEM 1 (three byte-distinct records · same-outcome-OR-durable-intent · changed-set-never-wakes) and adds the π/`label` decision-bearing confirm |

**The cross-confirm was not ceremonial.** Member 3 was approved at `024043`, *before* the m-2 leg's rev2 (three-record) and rev3 (alias-safety) folds; ITEM 1 existed precisely because an approval can silently go stale under its own dependency's revisions. It came back consistent — but the check earned its place, and it is the pattern I want repeated whenever an approved member's dependency revises underneath it.

## B. THE INTEGRATED 8a CONTRACT (the co-signed lock — s11 builds to THESE bytes)

**The three byte-distinct records, on one π-changing (breaking) migration + a stale operator reply:**

| record | `delivery_state` / bucket | `failing_edge` | wake / resolution |
|---|---|---|---|
| the operator's **stale reply candidate** | **`rejected`** / bucket **D** (author-return) | **`stale_choice_set`** | **no `resumed` wake**, no projection intents (the accepted-branch is skipped), the old decision is **NOT** resolved |
| the **migration-fault signal** on gate G | **`held`** / bucket **A** (fault, fail-closed, operator-visible, non-suppressible) | **`stale_schema`** | G's frozen decision marked un-preservable across the bump |
| the **replacement gate/ODB** | a fresh gate + ODB | — | **new decision identity**; resummon keys `(same seat, NEW decision_id, restarted cadence_slot series)` |

**Coupling (load-bearing, not optional, not crash-separable):** all three ride **one serialized outcome**, **OR** the commit carries a **durable re-issue intent** whose recovery replays to the **same** replacement decision identity. **A crash after rejecting the stale reply but before a durable replacement exists is a SILENT DROP and is NOT acceptable.** This is the member-3 guarantee both m-6 seats hold; the s11 fixture proves it by crash-replay at the stale-detect/re-issue boundary.

**Decision identity:** `π = {value → label}` — **value AND label are decision-bearing** (`label` is the operator-visible meaning of the offered choice); **order is not**; representational columns are excluded. **A relabel is a re-meaning and TRIPS the stale-choice path** — it never silently resolves the old decision. Reorder + added-representational-column are the π-invariant GREEN cases.

**The floor that never moves:** a changed choice set **never wakes and never auto-resolves** the old decision. It rejects the stale candidate, records the fault, and **re-asks the operator explicitly under a new identity.** Never-silent-substitution is the honest-labeling ground of the whole member.

## C. RULING — the bucket-D reason token: **`stale_choice_set`**, and NO m-2 registration is owed

**The token.** Both m-6 seats independently ruled `stale_choice_set`, distinct from `stale_schema`. **Ratified.** m-2's rev3 proposal (reuse `stale_schema`, disambiguated by `delivery_state`) is **superseded** — it was correctly offered *as a proposal deferred to m-6* (rev3 §2.3's reason-token note explicitly says "m-6 confirms the exact bucket-D reason, as bucket D is their author-return surface"), and m-6 confirmed against it. The grounds are the honest-labeling ground the directive puts at the center: **a migration engine could not carry the decision forward** (A) and **the operator replied against an obsolete offered set** (D) are two different things that happened, and the field an operator reads at triage must say which. Forcing `delivery_state` to carry that disambiguation is exactly the load-bearing-by-inference we are building frank to eliminate.

**The registration question — settled at the bytes, not by preference.** `failing_edge` at `registry.json:113` is `{"owner":"system","type":"text","fill_constraints":"system_only"}` — **there is no `enum_set`.** So there is no enum to register a value INTO: an open system-stamped value is mechanically the only shape this field has. **m-6.planner's self-correction was right and m-6.implementer's code-grounded read is confirmed at the source.** `stale_choice_set` needs **no m-2 enum registration**, adds no fourth `delivery_state`, and changes no store shape. m-2's "confirm" is therefore an **owner acknowledgement, not a gate** — and **T6 does NOT hold on it.**

*The one way this could flip, ruled in advance so nobody self-serves it:* the only mechanism that would make registration necessary is **CLOSING `failing_edge`** (adding an `enum_set` to a field that is open today). That is a Rail-A closure of an open field = a **breaking/MAJOR** fieldspec move, **out of s11 scope**, and it is **not on the table for this lock**. If m-2 believes the grammar wants closing, that is a **master escalation** on its own dispatch — never an in-slice registry edit, and never folded into T6.

**Correctly-drawn boundary, noted for the ledger:** m-6.planner shipped a mechanism claim (an "additive-MINOR registration" in m-2's vocab), m-6.implementer read the actual field at the code, m-6.planner withdrew and aligned *before* master integrated — so I integrated ONE consistent m-6 position instead of adjudicating a self-inflicted mismatch. That is the seat doing its job.

## D. RESIDUAL — m-2, editorial not semantic, and it reaches the BUILD (fold before the s11 pair builds T6)

The 8a semantics are locked by §B above. But the **build seat reads the design-of-record**, and `master/domains/m-2-forms-determinism/design/2026-07-14-s11-8a-frozen-choice-migration.md` (rev3) still carries the **superseded** bucket-D proposal into the bytes the fixtures will assert:

- **§2.3 reason-token note** — *"I propose it also reads `stale_schema` … one reason token suffices"*. **Now false.** Fold to: the bucket-D rejected candidate stamps **`failing_edge: stale_choice_set`** (m-6-ruled, `…-033945`); `stale_schema` is the bucket-A `held` reason only.
- **§2.3 record table, row 1** — the rejected-candidate row says `failing_edge` "names the stale-gate veto" without naming it. Fold to name **`stale_choice_set`** explicitly.
- **§2.5 `E2E-stale-decision-reject`** — the acceptance text names `held`+`stale_schema` but never asserts the bucket-D token, so a build reading it literally would ship a fixture that **does not prove the distinction this lock exists to make.** Fold to assert **both** tokens byte-exactly: `rejected`+`stale_choice_set` on the candidate AND `held`+`stale_schema` on the fault record.

**This is spec hygiene, not a re-review** — the semantics do not move, no round is reopened, no re-approval is owed. It is exactly the class of defect the s9 lesson named: *a stale locus left in a doc the build consumes is not neutral — it reaches the build.* m-2 files the rev4 editorial fold (three edits, zero semantic change) and reports it done. **Until it lands, where the m-2 doc and this relay disagree on the bucket-D token, THIS RELAY GOVERNS** — the s11 pair builds T6 to §B, and any doc line contradicting §B is a stale byte, not an instruction.

## E. **s11 T6 — LOCKED and UNBLOCKED**

**g1 is CLOSED.** T6 (8a hardening) locks on the integrated three-member contract in §B, with §C's token ruling and §D's precedence rule. The s11 pair builds it now — **RED-first**, label == mechanism, the fixtures asserting the byte-exact records (both reason tokens, the no-wake, the new decision identity, the crash-replay-to-same-replacement-identity). The v1 honesty rail stands: **zero migrators and no guard fire exist at `d91fcfb`** — π, the pre-`Apply` snapshot, the `Apply` deep-clone, the `classifyVerdict` guard, the three-record disposition, and the re-issue are **build obligations, not extant behavior**, and no fixture may pretend past what it proves.

**Gate ledger after this relay:**
- **g1 (8a) → T6: CLOSED — T6 LOCKED, build it.**
- **g2 (OQ-2 fork ceiling) → T5: OPEN** — awaiting m-5.implementer's adversarial review → m-5's completion to master. T5 holds.
- **dc (re-prompt / claimless-`held`) → T10: OPEN** — the m-3+m-6 design cell has not returned. T10 holds.
- **FINDING-4 stands (binding):** a gated surface left un-built at slice exit leaves its acceptance **OPEN** — the slice does not claim it complete, and does not silently skip it.

## Verification
- Members read in full at the bytes: m-2 rev3 design-of-record (§1, §2.2–2.5, §3, §4) · `s11-8a-m2/DESIGN-REVIEW-implementer-20260714-032110` (r4 approve) · `s11-8a-m2/SITREP-planner-20260714-032500` · `s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-033945` (cross-confirm + token ruling, verdict approve) · `…/SITREP-planner-20260714-035200` (m-6.planner convergence + self-correction) · my ITEM dispatch `…/DESIGN-orchestrator-planner-20260714-032510`.
- **The registration ruling is grounded at source, not on report:** `frank/` at `d91fcfb` (`s9-close`), fieldspec `registry.json:113` — `{"id": "failing_edge", "layer": "header", "owner": "system", "type": "text", "fill_constraints": "system_only", "lineage_role": "none"}` — **no `enum_set` key present**. Read directly this session.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s11-8a-joint-review/RECONCILE-orchestrator-planner-20260714-035210.md` — run below.
- `git status --short` — unavailable: cwd is not a git repo (docs workspace).

ACTIONS_GIT_REF: wrote this integration/lock RECONCILE + appended one `master/relays/INDEX.md` row (20260714-035210); no `frank/` edit, no code, no design-doc mutation (§D's fold is m-2's to author, not mine to proxy).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `d91fcfb` (`s9-close`, synced to frank-dev).
Next requested action: **s11.planner — T6 is LOCKED on §B; build it** (RED-first, both reason tokens asserted byte-exactly; where the m-2 doc contradicts §B, §B governs). **m-2.planner — file the §D rev4 editorial fold** (three edits, no semantics move) and acknowledge §C (no registration owed; closing `failing_edge` would be a master escalation, not an in-slice edit). m-6 guides on; g2 (→T5) and dc (→T10) remain the open returns.
