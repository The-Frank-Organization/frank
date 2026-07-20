# DESIGN (m-2 leg) — 8a joint review: the `stale_schema` reason token + the frozen-choice/migration interaction

**DESIGN_DOC_ID:** `s11-8a-joint-review` (m-2 leg — members 1 & 2; m-6.implementer owns member 3, the re-issue branch; one joint co-signed return)
**Owner:** m-2.planner (Forms & Determinism) · **Co-review:** m-6.implementer (member 3) · **Consumes:** the co-signed 8a FLOOR
**Base:** `s9-close@d91fcfb`, registry `s10-fieldspec-v8`.

> **rev2 fold-log (against m-2.implementer `s11-8a-m2/DESIGN-REVIEW-implementer-20260714-030153`):** Member 1 **APPROVED**; §2.1 projection **PASSED** — both unchanged. Two integration blockers folded:
> - **MR-1 (guard was off the live path):** rev1 put the guard in `migrate.Reader.Read`, which operator resolution never traverses (`submitHandlerWithObservation` builds a raw un-migrated `tables.Build` table → `classifyVerdict` reads the ODB directly, `submit.go:43-66,146-153,527-552`; `tables.Build` has no migrator, `tables.go:110-125`). **rev2 moves the guard onto the live verdict path — `classifyVerdict` obtains the ODB through a guarded migration step** (§2.2), keyed on the **immutable source** record being `record_kind==odb` (so a kind-changing migrator can't bypass), fail-closed on either-side projection parse, validating the pick against the **migrated** view. End-to-end `SubmitHandler` wake fixtures bound.
> - **MR-2 (held/rejected conflated; re-issue coupling loosened):** rev1 called `held`+`stale_schema` itself the "bounce-as-stale" and made re-issue optional/downstream. **rev2 separates the byte-distinct records** (§2.3): the operator's stale reply candidate = **`rejected`/no-wake (bucket D)**; the migration-fault signal = **`held`+`stale_schema` (bucket A)**; the fresh replacement gate/ODB = **member 3**, emitted in the **same serialized outcome or a durable re-issue intent** (m-6's approved atomic coupling, not optional, not crash-separable).

> **rev3 fold-log (against m-2.implementer `s11-8a-m2/DESIGN-REVIEW-implementer-20260714-031120`):** MR-1 + MR-2 **PASSED**. **MR-3 (Go map-aliasing bypass)** folded — the guard's `π(source)==π(migrated)` was unsafe because `record.Record.Headers`/`XFields` are maps (`record.go:29,31`) shared by `Apply`'s shallow `out := rec` (`migrate.go:64`), and the canonical migrator mutates `Headers` in place (`migrate_test.go:31-32`): an in-place choice migrator would mutate the *shared* map, so both "source" and "migrated" expose the changed choices (`π` equal ⇒ bypass) **and** the live `t.Records[odb]` view is corrupted. **rev3 (§2.2):** (1) **snapshot `π(source)` BEFORE any migrator runs**; (2) **`migrate.Apply` deep-clones** (Headers+XFields) before each migrator step — the ownership locus (m-2's read-side seam, §9), defense-in-depth for all callers; (3) compare the **pre-Apply snapshot** vs `π(migrated)`, either-side parse fail-closed; (4) an **in-place mutating-migrator fixture** proves the guard fires + source/table byte-equivalent in memory AND on disk.

> **rev4 fold-log (EDITORIAL — no semantics move; against master RECONCILE `s11-8a-joint-review/RECONCILE-orchestrator-planner-20260714-035210` §C/§D, after the 8a lock integrated + T6 locked):** the bucket-D reason token is ruled **`stale_choice_set`** (distinct from the bucket-A `stale_schema`); my superseded rev3 proposal (reuse `stale_schema`) is folded out of the three loci the build reads — §2.3 record-table row 1 (names `stale_choice_set`), §2.3 reason-token note (RULED: two distinct tokens, no m-2 registration owed since `failing_edge` has no `enum_set`), §2.5 `E2E-stale-decision-reject` (asserts BOTH tokens byte-exactly). No round reopened, no re-approval owed — spec hygiene so the build fixtures prove the A/D distinction (the s9 stale-locus lesson).

> **The co-signed 8a FLOOR (both pairs — not re-opened):** a parked gate waking across a schema bump runs **migrate-then-validate**; **un-migratable → `held`/escalated, NEVER silently dropped, never auto-resolved.**

---

## 0. Frame — the v1 honesty rail

At `d91fcfb` the migrator seam is landed but **zero migrators are registered** (`migrate.go:11` `Current=1`, empty `steps`): every record is v1, `Apply` returns unchanged, no bump exists. **Neither member fires at v1** — both are the contract the first breaking bump activates, RED-first-testable against a synthetic v1→v2 step, not present behavior. Same rail as §9.

---

## 1. Member 1 — the `stale_schema` reason token (CONFIRM — pair-APPROVED, unchanged)

A **system-stamped `failing_edge` reason value** (registry `:113` — `owner:system`/`type:text`/`system_only`; no new `enum_set`) — the **third `held` producer** (wake-migration-failure), joining §17.1's two. **Disposition** = `delivery_state: held` (bucket A per the m-6 terminal map — fault/fail-closed, operator-visible, non-suppressible); `delivery_state` stays byte-locked at three (`:71`) — `stale_schema` is a reason UNDER `held`, not a fourth outcome. **Version class = additive-MINOR / Rail-A OPEN** (no exhaustive consumer on the reason; held/escalate keys on `delivery_state==held`; unknown reason still routes to generic held-escalate — never a silent accept). *(Pair-verified: no exhaustive `failing_edge` consumer in the tree.)*

## 2. Member 2 — the frozen-choice / migration interaction (CONFIRM — enforceable on the live path, rev2)

**The subject.** A gate parked with a bounded choice set FROZEN at park — the ODB `choices` row_array (registry `:183`; each row exactly `{value, label}`, value-unique, per the decoder `odb.go:125-129`) on an immutable committed record. The migrator seam permits an arbitrary transform (`migrate.go:19`, `Apply` advances version only `:64-75`), so the invariant needs an **executable, live-path** boundary.

### 2.1 The canonical decision projection `π` (PASSED — unchanged)

> **`π(choices) := { value → label }`** — the value-keyed map from the parsed rows.
- **Decision-bearing = `value` AND `label`** (value = pick identity `ValidateODBChoice` matches on, `odb.go:133`; label = offered meaning — a relabel is a re-meaning).
- **Order NOT decision-bearing** (π is a map; `ValidateODBChoice` matches by value order-agnostically) ⇒ **reorder is a legal structural transform.**
- **Representational columns excluded by construction** (π reads only value+label; a new column is invisible ⇒ representational-by-default).
- Decision identity: `π(source) == π(migrated)` as maps.

### 2.2 The enforcement locus — on the LIVE operator-resolution path (MR-1 fix)

The guard runs where resolution actually happens: **`classifyVerdict`** (`submit.go:527-552`), at the point it selects the `odb-<gateRef>` record from the raw table. The executable order is **alias-safe** (MR-3): the source projection is captured before any untrusted migrator runs, and the migrator receives a deep-cloned record.

1. **Key on immutable source identity.** The guard applies when the **stored source** `odb-<gateRef>` record has `record_kind == odb` — read from the immutable committed record, **not** the migrated `record_kind` (so a migrator that changes the kind cannot bypass the choice guard).
2. **Snapshot `π(source)` BEFORE any migrator runs (MR-3).** Parse the immutable stored `choices` and capture `π(source)` as a **value** (a copy of the `{value→label}` pairs, never a reference into the record's `Headers` map) **before** calling `Apply`. This is the invariant's left operand; it cannot be mutated by a later migrator.
3. **Deep-clone at the migrator boundary — the alias-safety locus (MR-3).** `record.Record.Headers`/`XFields` are maps (`record.go:29,31`); `Apply`'s `out := rec` is a shallow copy sharing those maps (`migrate.go:64`), and a migrator may mutate `Headers` in place (`migrate_test.go:31-32`). So **`migrate.Apply` deep-clones the record (Headers + XFields) before invoking each migrator step** — the chosen ownership locus (m-2's read-side migration seam, §9), which protects **every** caller (the `read` tool, replay, and this verdict path) and prevents any migrator from aliasing the caller's source record or the live `t.Records[odb]` table view. *(Named s11 hardening of `migrate.Apply`; not extant at `d91fcfb`.)*
4. **Obtain the guarded migrated view.** `classifyVerdict` resolves the ODB through this guarded step — `migrate.Apply(reg, source)` returning a deep-cloned migrated view; the raw `t.Records[odb]` read is replaced by it (the migrate registry threaded into the verdict path alongside the existing `migrate.Current` use at `submit.go:56` — a named s11 integration). *(Equivalent alternative: table construction carries a guarded source+view pair — either satisfies the "no raw `tables.T` bypass" bar.)*
5. **π-invariance guard + fail-closed parse.** Require **`π(source-snapshot) == π(migrated)`** (the left operand is the step-2 pre-Apply snapshot, not a re-parse of a possibly-aliased record). **Projection parse failure on EITHER the source snapshot or the migrated view fails closed** into the typed incompatibility path (never silently proceeds).
6. **Validate the pick against the MIGRATED view.** On guard pass, `ValidateODBChoice(migrated, pick)` — so the GREEN structural-column case validates against the v2 shape (wiring "resolution proceeds after the structural migration," which rev1 wrongly attached to `Reader.Read`).
7. Any of {`Apply` fails · guard fails · either-side parse fails · migrated view fails current-form `Validate`} → the deterministic typed incompatibility signal (§2.3).

**Named s11 build obligation, RED-first.** m-2 owns π + the pre-Apply snapshot + the `Apply` deep-clone + the guard + the typed signal; the disposition of that signal (§2.3) is the m-6 seam.

### 2.3 The byte-distinct records + m-6's atomic coupling (MR-2 fix)

The typed staleness signal is deterministic; its disposition is **three byte-distinct records**, not one — matching the m-6 terminal-bucket map (`rejected`→D author-return; `held`→A fault-lane; never synonyms) and m-6.impl's approved branch (`s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043`):

| record | delivery_state / bucket | who | trigger |
|---|---|---|---|
| **the operator's stale resolution candidate** | **`rejected`** / bucket **D** (author-return), **no `resumed` wake** | the reply the operator submitted | `classifyVerdict` returns it rejected when the guard detects staleness — **`failing_edge: stale_choice_set`** (m-6-ruled, `s11-8a-joint-review/RECONCILE-orchestrator-planner-20260714-035210` §C; distinct from the bucket-A `stale_schema`); no projection intents emitted (the `submit.go:147` accepted-branch is skipped) |
| **the migration-fault signal** | **`held` + `failing_edge: stale_schema`** / bucket **A** (fault/fail-closed, operator-visible, non-suppressible) | the gate G | marks G's frozen decision un-preservable across the bump (member 1) |
| **the fresh replacement gate/ODB** | new gate/ODB, **new decision identity** | **member 3 (m-6)** | the re-issue — new `(seat, NEW decision_id, restarted cadence-slot series)` resummon keys (m-6.impl pt 2) |

**Atomic coupling (m-6.impl pt 3, preserved — NOT loosened).** The typed migration signal is **consumed in the same serialized outcome** that rejects the stale candidate and emits the held record + the fresh replacement — **or** the commit carries a **durable re-issue intent** recovery replays to the same replacement decision identity. A crash after rejecting the stale reply but before a durable replacement exists is a silent drop and is **not acceptable.** rev1's "downstream authorized act" is **retracted** — re-issue is not optional or crash-separable.

**Decision integrity is still preserved** (the concern rev1 over-corrected): the operator's stale reply is **rejected** (never silently accepted against a shifted decision), and the replacement is a **new-identity** gate (the operator is re-asked explicitly, never auto-resolved). Reject-the-stale-reply + atomically-re-issue-a-new-identity-gate reconciles both m-6's durability and the never-silent-substitution floor.

*(Reason-token note — RULED (rev4 fold, `s11-8a-joint-review/RECONCILE-orchestrator-planner-20260714-035210` §C): the two reasons are **distinct tokens**. The bucket-A migration-fault `held` record carries **`stale_schema`** (member 1 — a migration engine could not carry the decision forward). The bucket-D `rejected` stale candidate carries **`stale_choice_set`** (the operator replied against an obsolete offered set). Both m-6 seats independently ruled `stale_choice_set`; my earlier proposal to reuse `stale_schema` disambiguated by `delivery_state` is **superseded** — forcing `delivery_state` to carry that disambiguation is the load-bearing-by-inference frank exists to eliminate; the field an operator reads at triage must say which thing happened. **No m-2 enum registration is owed:** `failing_edge` (`registry.json:113`, `owner:system`/`type:text`/`system_only`) has **no `enum_set`**, so an open stamped value is mechanically the only shape — `stale_choice_set` adds no enum, no fourth `delivery_state`, no store-shape change. Closing `failing_edge` (adding an `enum_set`) would be a breaking/MAJOR move, a master escalation on its own dispatch, out of s11 scope — not an in-slice edit.)*

### 2.4 The operator's resolution validates the FROZEN offered set (CONFIRM — pair-verified at v1)

`classifyVerdict` finds the committed `odb-<gate>` record and validates against **its** `choices` (`submit.go:527-552`), never re-deriving from the live registry. **Additive widen ≠ retroactive widen** — nothing in the resolution path widens a committed ODB's choices; new options appear only in new gates. The common (additive) bump leaves parked gates untouched; only a π-changing (breaking) migration trips §2.3.

### 2.5 Fixtures (helper-level RED/GREEN + end-to-end wake — MR-1 item)

**End-to-end `SubmitHandler` wake fixtures (production verdict path — required):**
- **`E2E-stale-decision-reject`** — a parked v1 gate+ODB, a v2 with a π-changing migrator, an operator reply from the old set → the candidate is **`rejected` + `failing_edge: stale_choice_set`** / no-wake; a **distinct `held` + `failing_edge: stale_schema`** fault record + a fresh replacement gate/ODB (new decision identity) are emitted in one serialized outcome (or durable intent); crash-replay at the stale-detect/re-issue boundary reaches the same replacement identity. **The fixture asserts BOTH tokens byte-exactly** (`rejected`+`stale_choice_set` on the candidate AND `held`+`stale_schema` on the fault) — proving the A/D distinction this lock exists to make.
- **`E2E-structural-migration-resolves`** — a v2 whose choice-row form permits an added representational column, a migrator that fills it preserving every `{value,label}` → the reply **resolves against the migrated view** (π invariant), gate wakes normally.

**Helper-level RED/GREEN (projection guard unit):** RED → typed-incompatible: add / remove / rename-value / relabel / wholesale-replace, plus `NF-migrate-gap` (`ErrMigrationGap`). GREEN → π invariant: reorder + structural-column. Each asserts **source stored bytes immutable** + **resolution uses the preserved frozen π.**

**Alias-safety fixture (MR-3 — required):** `NF-migrate-choice-inplace-mutate` — an **in-place mutating** migrator (matching the `migrate_test.go:31-32` style) that replaces `Headers["choices"]` through the shared map. Assert: (a) the guard **fires** (`π(source-snapshot) ≠ π(migrated)` — the pre-Apply snapshot is unaffected by the in-place mutation) ⇒ the typed incompatibility signal; (b) the **raw source/table record is byte-equivalent in memory** (the `t.Records[odb]` view is untouched — proving the `Apply` deep-clone isolates the migrator) **and on disk** (the immutable stored record is unchanged).

## 3. Seam to member 3 (m-6.implementer — grounded, not owned here)

My §2.2 boundary produces the **single deterministic typed staleness signal** on the live verdict path; m-6's member 3 owns its **disposition** — the three §2.3 records and the atomic-or-durable coupling (m-6.impl pts 1–3, approved). Grounded: a changed offered set = a **new decision identity** = a new gate; its resummon keys are `(same seat, NEW decision_id, restarted cadence-slot series)` (m-6.impl pt 2, `resummon.go:281-288/120-127`), so the old `(seat, old decision_id, slot)` keys cannot suppress the replacement. m-6 owns whether the coupling is one serialized commit or a durable re-issue intent. My leg fixes only the upstream detection + typed signal.

## 4. Out of scope / anti-half-fix guards

- **v1 honesty rail (§0):** the guard + migrator are s11 RED-first build obligations, not extant.
- `delivery_state` byte-locked at three; `stale_schema` is a `failing_edge` reason.
- The guard is on the **live verdict path**, keyed on **immutable source ODB identity**, **fail-closed** on either-side parse — not a `Reader.Read`-only check.
- **Alias-safety (MR-3):** `π(source)` is snapshotted **before** any migrator runs, and `migrate.Apply` **deep-clones** (Headers+XFields) before each untrusted migrator step — so no migrator can alias the caller's source record or the live `t.Records` view. The guard compares the pre-Apply snapshot, not a re-parse of a possibly-mutated record.
- `held` (bucket A, migration fault) and `rejected` (bucket D, stale reply) are **byte-distinct** — never synonyms.
- Re-issue is **not optional/crash-separable** — atomic-or-durable per m-6.impl pt 3.
- No lock moves: §9, §17.1, the `{accepted, rejected, held}` byte-lock, the immutable store, the `{value,label}` ODB row schema, and the m-6 terminal-bucket map are unchanged.

## 5. Open items for the joint co-sign / pair review

1. **m-6.implementer:** member 3 consuming the §2.3 signal — the three-record disposition + atomic-or-durable coupling (your pts 1–3, already approved); confirm the guard firing at `classifyVerdict`/the wake path (you own the ODB surface + scheduler), and the bucket-D reason for the rejected stale candidate (§2.3 note).
2. **m-6 (ODB surface):** confirm `π={value,label}` as the decision-identity — in particular **`label` decision-bearing** (a relabel = re-meaning); flag if held display-only.
3. **m-2.implementer (bounded re-review):** MR-3 closure — (a) `π(source)` is snapshotted before any migrator runs; (b) `migrate.Apply` deep-clones (Headers+XFields) before each migrator step so no migrator can alias the source record or the `t.Records` view; (c) the guard compares the pre-Apply snapshot, either-side parse fail-closed; (d) the `NF-migrate-choice-inplace-mutate` fixture proves the guard fires + source/table byte-equivalence in memory and on disk. *(MR-1 live-path, MR-2 three-record/atomic-coupling, the projection, member 1, and the v1 honesty rail passed rev2 — not reopened.)*
