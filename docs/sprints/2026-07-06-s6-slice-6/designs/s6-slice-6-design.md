# s6 Slice-6 build design — the transport fix, decomposed against the co-signed amendment set

**DESIGN_DOC_ID:** `s6-slice-6-design` · **Rev:** r3 (EXTERNAL-VERDICT FOLD per the orchestrator directive `s6-core-design-verdicts` 20260707-012809 — m-7 R1(i–iii) + R2 folded verbatim from `master/relays/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` (must-revise, PRE-CONCURRED on faithful fold); m-1 F-S6-M1-4 commit-time guard folded from `s6-fidelity-m1/SITREP-implementer-20260707-012143.md`; r2 = grill folded §18)
**Owner:** s6-core.planner · for s6-core.implementer DESIGN-REVIEW (v2.8.5)
**Dispatch:** `.relays/s6/s6-core-design/DESIGN-orchestrator-planner-20260706-234822.md` (r2) · **GRILL_REQUIRED: yes**
**Spec of record (read-only, this doc re-decides NOTHING it locks):** `master/S6-AMENDMENT-SET-2026-07-06.md` (r3) + m-1 §A–§F.1 + m-7 r5 + m-2 codec (+§11) + `master/GRILL-LOCK-parenting-fork-2026-07-06.md`. Baseline: `frank/` code surface `s5-close` (`7e5c527`).
**Audit basis:** both s6-core audits (reconciled, zero contradictions) — all root causes reproduced at audited line numbers; cites below are from those audits.

## §0 Constraints, watchpoints, claim pins (binding on every section)

- Byte-exact `{accepted, rejected, held}`; seat surface exactly `submit`/`project`/`read` — roster + audit views are `project` parameters (the violation shape = a new `ToolSet` entry / `names()` member); `boot-required` and the lock refusal are a rejection class and a typed process exit, never delivery states; `held` stays off seat default projections.
- I-PH (`Field:Class`, path-free) over every surface incl. the FIVE new payload families (roster rows · boot-required per-field detail · lock-refusal diagnostic · hint flags · **the `seat_mint` accept-reply — r3/R2**); `bounce.Format` stays the single formatter. **The fifth family carries TWO NAMED CARVE-OUTS** (the decision-⑤/frame-ceiling pattern — explicit scoped exemptions, never silent leaks): the fresh credential + the new seat's endpoint, **operator-channel-only**; the matrix asserts both appear in NO committed record Body, NO projection or log, NO non-operator reply; FX-D2's byte-parity fixture is per REJECTION class, so the accept-reply divergence (reply ⊃ record) is deliberate and stated, not a parity bug.
- [VP-W2]: FX-B1g is in the exit gate. [VP-W3]: the transport registry pass = EXACTLY the seven named rows, NO activation-marker row, activation DERIVED-ONLY (see §14.1 for the boot-field reading, surfaced for explicit confirm).
- Fallback-never-bounce with the empirical why (zero true bad picks in recorded history) carried into fixtures. The step-exit operator legs are never simulated.
- Claim pins: "at-least-once intake, exactly-once EFFECT" (never unqualified exactly-once) · "un-bounceable boot" = structural absence of digest/parent/divergence classes only (never immunity from auth/admission rejection) · activation = liveness bookkeeping, never identity-strength · hint fallback ≠ hint validation (intent + outcome + flag is the honest triple) · the lock = confusion-resistant ops, D5 same-uid residual verbatim · `charter_loaded` = `self_reported`.
- **No-perf fence inside edited files:** the known hot-path rescans (per-submit `tables.Build` in `cmd/frank/main.go process()`; `gate.Complete` store scans) are NOT touched, even incidentally, while this slice edits those files. Engine performance work of any kind is OUT.
- OUT (name-and-escalate): Step-2 observe · Step-3 routing execution · perf · new seat verbs · federation · dogfood-in-slice · governance-doc edits · locked-contract changes (amendment path only).

## §1 Shape of the diff (module map)

| module | changes | spec |
|---|---|---|
| `internal/fieldspec` | exported address-list codec (§2); `recipient_picker` render (§2); digest volatile-class (§3); registry pass (§14); boot-form render gating consumes lifecycle (§12) | m-2 §2/§2.5, m-7 A-1, m-2 §11 |
| `internal/lineage` | comma-splits deleted → codec (§2); candidate mechanics deleted (§4); waiver gate reads scoped-waiver effective state (§8); hint-class validation host (§4/§14) | m-2 §2.2, m-1 §A/§C |
| `internal/engine` | commit-time PARENT stamp in the shared handler (§4); pre-active admission B-1.2a (§12); `validateRecordKind` → layer 3 only (§9); `seat_mint`/`waiver_retraction` classifiers (§7/§8); Outcome detail parity (§11); intake replay/coalescing consult (§5) | m-1 §A, m-7 B-1/A-2/D-2, m-2 §2.6 |
| `internal/intake` | durable counter (segment header high-water); dedupe→replay hand-off; in-flight coalescing; dead `journal.Append` id path deleted/unified (§5) | m-7 A-2 |
| `internal/store` | `Project` accepted-filter + view params (§6); rebuild filtering (§6); lock acquire/refuse (§10 half); mailbox/render/INDEX print full decoded recipient truth (§2) | m-1 §B/§D, m-2 §2.4 |
| `internal/recover` | phase −1 lock claim (§10); phase-3 lifecycle derivation (§12); rebuild hygiene (§6) | m-7 A-4/B-1.1 |
| `internal/seat` | re-mint choreography: binding replacement + old-credential death (§7, GRILL row 1) | m-1 §F/§F.1/§13.3 carry |
| `internal/channel` | `bound_now` exposure from the active index; old-credential channel close on re-mint (§7/§12) | m-7 B-1.3 |
| `internal/tables` | `ContentHash` wired; lifecycle table (generation, activation) (§5/§12); waiver/retraction effective-state (§8) | m-7 A-2/B-1, m-1 §C |
| `cmd/frank` | `turnContextForSeat` → accepted-state (§4/§6); `project` params (§6/§12); mint CLI → genesis-only (§7); roster view (§12) | m-1 §B, m-7 A-3/B-1.3 |
| `cmd/frank-mcp` | one transparent reconnect-retry (§13); re-render pattern-match hack retired → typed-field consumer (§11) | m-7 D-1/D-2 |
| `test/*` | the full fixture table (§16), red-first | all |

## §2 The codec — one decode for every judge (m-2 §2; kills F6/F7, the render half of F6)

- `internal/fieldspec`: export `DecodeAddressList(raw string) ([]string, error)` / `EncodeAddressList(list []string) (string, error)` — thin exported entries over the existing `ParseTyped` `address_list` path + `CanonicalMarshal` (promote, not rewrite; `canonical.go:36-40,62-75`).
- Call sites converted: `lineage.addressedTo` (:282-292), `addressedInHeader` (:432-439) — the raw `strings.Split(",")` is DELETED; `store.DeliveryRecipients`/`addressListHeaders` (projections.go:122-142); the rendered relay markdown header + `INDEX.md` row (projections.go:110-113) print the full decoded Header `TO`/`CC` list — the canonical decoded header list is the single recipient truth for EVERY projection (m-2 §2.4 rev1); `Envelope.To` remains an optional m-7 compat/routing projection only.
- **No silent drop:** `DeliveryRecipients`'s `if !ok { return recipients }` (:135-136) is removed. A recipient list that reaches delivery already passed `ParseTyped` at submit; a decode failure at delivery is an engine invariant violation → surfaced through the loop's internal-fault disposition (held/rejected per `AuthorityBearing`, the existing `faultOutcome` path), never a silent loss. At the submit gate an ill-formed list bounces `Field:Class` (already live via the typed check).
- `recipient_picker` implemented in `render.go renderField` (new case, symmetric with `parent_picker`): candidates = the minted address space (binding-table seat names via the render env), `DigestExempt`/volatile-class marked (§3). The form offers exactly what validate/gate/delivery accept.
- Fixtures: m-2 §7.1 (multi-TO+CC identical across markdown/INDEX/mailbox intents/visibility gate; parse-fail bounces at submit; no deliver-then-drop path) · §7.2 (the archived relays `relay-c13dc32f`/`relay-1f99aadf` pass both gates on one encoding) · §7.6 (recipient_picker renders + digest-exempt).

## §3 A-1 — digest = stable schema surface; volatility exempt by class (m-7 A-1; kills F5)

- Digest input: config digest + seat pattern + phase + tier + field SHAPES (ids, types, constraints, static enums). Fields whose options are conductor-supplied volatile are marked by CLASS at render (`conductor_supplied_volatile`) and stripped structurally in `formForDigest` — promote the existing per-field `DigestExempt` mechanism (render.go:68-72, :190-200) to the class rule.
- The volatile class at s5-close+this-slice: `recipient_picker` candidates (TO/CC), `parent_picker` (dies with §4 anyway), projection-derived `grant` availability (`RealGrantState`), monotonic floor-trimmed options. Static enum options stay digest-covered (they are config-shape, already pinned by the config digest).
- The bounded re-render contract: a submit bounces `re-render` iff the effective schema shape changed (config/registry §7 record; seat/phase/tier mismatch) — NEVER because unrelated traffic committed. Digest-lease/TTL stays rejected (recorded, not deferred — m-7 A-1.4).
- Fixtures: FX-A1a (N foreign accepts ⇒ zero re-render bounces) · FX-A1b (§7 config change ⇒ the held form bounces — the stale-SCHEMA guard still bites).

## §4 Branch-A parenting — conductor-computed PARENT + validated hint (m-1 §A + GRILL_LOCK; kills F11/F4 at the root)

- Registry (§14): `PARENT_DISPATCH_ID` → `owner: system` / `fill_constraints: system_only` ⇒ `renderable()` false ⇒ **the parent field leaves every form** (F2 dissolves; m-2 §3 strongest form). New rows: `parent_hint` (lane-optional id_ref), `parent_hint_honored` (system-computed), `parent_provenance` (system-computed enum `{woken_on, active_dispatch, dispatch_root, hint}`).
- **Stamp locus = the shared submit handler** (`engine.SubmitHandlerWithRender`), not the loop body — so recovery re-execution through `RunWithProcessor`'s `process` closure behaves byte-identically to live traffic (the same handler serves both; this is a correctness requirement, stated here so the PLAN pins it). Stamp: PARENT = first-defined(seat's woken-on ACCEPTED relay → active-dispatch lineage → dispatch root), exactly the GRILL_LOCK derivation; payload-supplied PARENT ignored byte-for-byte (the NF-S1 idiom, `seat.Stamp` precedent).
- Hint honored iff engine-provable: hinted id resolves in the seat-provable set — (the seat's delivered mailbox history ∪ the seat's own accepted records) ∩ the accepted graph — with the generous prover (dispatch-id → that dispatch's accepted thread root; `relay-` id direct lookup) from the GRILL_LOCK's non-binding note. Unprovable ⇒ stamp the computed default, `parent_hint_honored: no` in the submit response, verbatim hint preserved on the record. **Never a bounce.**
- The anchoring-bounce class ceases: `validate.go:69-76` (`active-lineage-empty`/`outside-active-lineage`) and `lineage.ActiveLineageCandidates` (:346-372) are DELETED. Class-lineage authority gates (design-review chain, pair-dispatch grant, IMPL addressee, merge claim) are UNTOUCHED and still bounce — now with §11 detail parity. Stated interplay (honesty): where a class gate keys on PARENT (e.g. delegated `DISPATCH IMPL` must parent the approving PLAN-REVIEW), the seat steers via `parent_hint`; a fallback-stamped parent may then bounce at the CLASS gate with actionable detail — that is the authority layer working, not an anchoring bounce.
- Turn context: maintained from accepted-state only (§6's `WokenOn` fix feeds this by construction).
- Fixtures: the GRILL_LOCK three (§16 rows G-1..G-3) · m-2 §7.4 (PARENT absent from every rendered form) · hint honored/fallback with audit-visible flags.

## §5 A-2 — intake identity: replay, coalesce, durable ids (m-7 A-2; kills F9)

- **Durable counter (GRILLED — §18, LOCKED):** each journal segment file begins with a header line `{"segment_header":true,"high_water":N}` (N = counter at segment creation); recovery/writer-construction restores `next = max(all headers, all entry ids)+1`. Ids survive GC of drained segments (every live segment's header carries the mark forward) and rotation and crash. Legacy headerless segments tolerated (max of entries). Monotonicity retained as locked — ordering is load-bearing (arrival-order re-enqueue; trail legibility).
- **Replay:** on content-hash hit with an existing OUTCOME (the wired `tables.ContentHash` → `OutcomeByIntake` lookup — both maps exist; `ContentHash` is dormant and gets populated at commit + recovery scan): the seat receives the original typed outcome (same relay_id, state, detail) — **nothing re-executes, nothing new commits.** The replay decision lives in the loop/handler (where tables are), not the writer.
- **In-flight coalescing:** content-hash hit with NO outcome yet ⇒ same intake id, no re-enqueue; the loop holds the additional reply channel(s) and answers all on the one execution (today's `writer.go prepare()` re-enqueues — that path dies).
- The dead `journal.Append` `len(entries)+1` id path is deleted (unified into the writer path); no second id judge survives.
- Restored invariant (fixture-stated): every `intake_id` has AT MOST ONE outcome record and every outcome's `intake_id` is unique — the §C4.1 1:1 anchor at the id grain.
- **The commit-time last-writer guard (r3 — F-S6-M1-4 verbatim):** before appending ANY outcome record with a non-empty `intake_id`, the loop checks the store-derived `OutcomeByIntake`; an existing outcome ⇒ replay-or-fault, **never a second outcome record** — the invariant is enforced at the writer boundary, not only observed. `TestOneOutcomePerIntakeSweep` is evidence the guard held over a store; it is NOT the guard.
- Fixtures: FX-A2a (byte-identical replay, no new record) · FX-A2b (in-flight coalesce, single execution, both callers answered) · FX-A2c (id uniqueness across rotation + crash + **the GC-drained-segments + restart id-reuse leg** + the 1:1 sweep) · the F9 duplicate-retry-across-crash leg (m-7 A-2.5).

## §6 §B — projection default-accepted + accepted-graph anchors (m-1 §B; kills F10's flood + the reject-anchor pool)

- `project()` default scope = **accepted records addressed to the caller**, filtered by canonical delivery state at serve time (`store.Project` maps mailbox lines → drops non-accepted via the tables/record state) — serve-time filtering is required because the archived s5 store already holds rejected ids in mailbox files (m-1 §E carry-forward verbatim; pollution writer = the rebuild path, `projections.go:38-46`, per the reconciled audits).
- Rebuild hygiene: `RebuildProjections` applies mailbox/render/INDEX intents for canonical records **by delivery state** — mailbox delivery intents for accepted only; rejected/held keep their canonical records + INDEX visibility (operator-grade), never seat mailboxes.
- `project` gains a view parameter (payload, not verb): `default` (as above) · `audit` (the caller's own submit attempts incl. rejects — the audit-of-attempts view) · `roster` (§12, operator/orchestrator-scoped). Author's bounce discovery: §11 puts the detail in the submit reply; `read(own rejected relay_id)` still works (unchanged).
- `held` records surface on the operator/escalation path only (unchanged C4.1 posture).
- `turnContextForSeat` (main.go:389-403) derives from the accepted-filtered view — a reject can never be `WokenOn` or any anchor in any interleaving.
- Fixtures: m-1 §E family — default projection filters a polluted archived mailbox (rebuilt-store leg) · audit scope returns them · accepted delivery unaffected · rejected/held never `WokenOn`/candidate/class-parent · author `read()` of own reject succeeds · enum untouched.

## §7 A-3 — live mint + re-mint as loop mutation (m-7 A-3 + m-1 §F/§F.1; kills F15)

- `record_kind: seat_mint` (operator-only seat_scope). **Parameters ride Body JSON** `{seat, role, is_operator}` (the config_change compound-record idiom, `submit.go:185-210` template; the F-M1-1 payloads-in-Body home) — no new field rows beyond the seven ([VP-W3]). Layer-3 check (§9): Body parses + members present + reserved-name (`system`) rejected.
- Accepted `seat_mint` = the one pivot; **derived work** (the S2 derived-work-completion mechanism, keyed by the record): binding-table update via `seat.Manager` + wake plumbing. Crash between pivot and completion ⇒ recovery completes idempotently (FX-A3b). No restart, no fleet bounce, zero digest rotation (A-1 makes a new seat projection-state, not schema-shape — FX-A3a).
- Credential custody unchanged: minted conductor-side; delivered ONCE in the operator's submit reply (never in any record, projection, or log — I-PH hard line); crash-window remedy = documented admin-time read of the 0600 binding table (the existing custody posture), never a new verb. — GRILL row 1 confirms.
- **Re-mint (the ratified audit item; GRILLED — §18 row 1, LOCKED):** a `seat_mint` for an EXISTING seat = re-mint. Mechanics: derived work REPLACES the binding row (fresh credential); the old credential dies at `Resolve` (binding replaced) and any live channel authed on it is force-closed at completion (the per-credential `active` index locates it); generation identity = the committed pivot in commit order (m-1 §F.1 verbatim — the records are the generation history; the binding table stays current-credential-only). `Mint`'s `ErrSeatAlreadyBound` guard moves from "reject remint" to "reject duplicate CONCURRENT mint" semantics inside the derived-work path. Ordinary session/model switching is untouched (same credential ⇒ re-bound, never re-boot). Routed to the m-1 fidelity packet (§13.3 credential-lifecycle carry).
- The admin `-mint` CLI retires to genesis-time only (empty store, pre-channel); on a live store it refuses with the documented choreography (m-7 A-3.5).
- `seat-mint` joins the F11 crash-class applicability map (one pivot, crash at every boundary ⇒ committed-or-not; the s4 `config-change` precedent).
- **R1 — the generation boundary for IN-FLIGHT commands (r3; m-7 must-revise folded verbatim; the code fact: `intake.Cmd` carries no session/generation tag):** force-close kills the channel but not commands already intaken/queued before the pivot — a queued old-session BOOT form would pass B-1.2a after the pivot and become the NEW generation's activation record, defeating the re-mint's intent while formally satisfying the order rule. Fold: **(i)** every command is tagged with its **auth generation** at handler-accept time — the seat's current `seat_mint` pivot ref, ONE persisted `Cmd` field, so recovery replay is byte-identical (the §4 stamp-locus principle); **(ii)** the loop typed-rejects any command whose generation ≠ the seat's current generation — class **`credential-superseded`**, path-free, §11 detail parity; **(iii)** FX-B1g gains the IN-FLIGHT leg (a boot form queued from the old session pre-pivot ⇒ rejected, does NOT activate; the new credential's boot does). A queued non-boot old-session submit was already safe (post-pivot minted-not-active ⇒ `boot-required`; rejected records never activate — m-1 §F.1). The R1 shape refines m-1's generation-boundary carry-forward ("records before a re-mint cannot activate the new generation") to commands STRADDLING it — **routed to m-1 for the narrow confirm** (the orchestrator's parallel packet; a dispatch precondition).
- **R2 — the mint accept-reply is the fifth I-PH payload family** with the two named operator-channel-only carve-outs (§0 r3 statement; matrix legs in §16).
- Fixtures: FX-A3a · FX-A3b · FX-B1g (the generation leg, [VP-W2]) · reserved-name/duplicate legs at the record grain.

## §8 §C — scoped waivers + retraction (m-1 §C + m-2 §4; kills F17/F12)

- Registry (§14): `rationale` (text) · `waiver_scope` (typed object row, canonical JSON: `{"kind":"run|dispatch|relay|record_class_dispatch", ...referent members}` — m-2 shapes the exact grammar) · `retracts` (id_ref). The `ORCH_REVIEW_WAIVER` `"*"` header row is RETIRED (row removed; accepted legacy records unchanged on disk — MINOR, no migrator, R-1 held; its mis-typed `lineage_role: routing_ref` dies with it).
- `record_kind: waiver_retraction` (operator-only seat_scope; layer-3: `retracts` present + resolves to an accepted waiver + not already retracted — the `disposes_owed` idiom verbatim, `submit.go:154-171` template).
- A WAIVER = an accepted record with `waiver_scope` present AND channel-stamped `FROM = operator` (record-authorship authority, not free-text seat-scoping — m-2 §4's resolution of its own expressibility limit). Non-operator records carrying waiver rows: submit-path rejected (the `classifyConfigChange` operator-check idiom); the render-absence leg ("fill-time absent") for the non-enum waiver rows is an **m-2 fidelity question flagged explicitly** — if current FieldSpec grammar cannot express operator-only rendering of non-enum rows, the design lands submit-path rejection + operator-form-only rendering via the minimal m-2-sanctioned rule, and m-2 rules the shape.
- Gate lookup (`checkReviewerVisibility` rewrite): effective waiver state = accepted waivers (scope matching the candidate) − accepted retractions, evaluated in commit order at gate time; retraction re-arms prospectively; pre-retraction accepts stand (immutability); unscoped LEGACY `ORCH_REVIEW_WAIVER` records read as run-wide until retracted (no retroactive reinterpretation). Bounded validity, if declared, is store-derived/deterministic — never wall-clock (m-1 §E carry-forward).
- `tables.Waivers` extends to carry retraction state.
- Fixtures: m-1 §E family — scoped waiver passes only in scope · retraction re-arms (post-retraction identical submit bounces) · pre-retraction accepts stand · non-operator waiver/retraction fill-time-absent + submit-rejected · legacy-unscoped-until-retracted leg · m-2 §7.5 (rationale first-class; gate reads the record, not the retired header).

## §9 F13 — three-layer record_kind (m-2 §2.6)

- Layers 1–2 stay where they already live (`reg.Validate`: enum membership `validate.go:54-58`; seat-scope :59-61 — verified in both audits). `validateRecordKind` (`submit.go:143-179`) is REDUCED to layer 3 only: per-kind required checks for `owed_item`/`owed_disposition` (unchanged) + `config_change` (its classifier, unchanged) + NEW `seat_mint` (§7) + `waiver_retraction` (§8). The membership case-list and the `default → "unknown record_kind"` DIE — an authorized-and-offered kind with no extra requirements passes with none.
- Live still-bouncing tokens named for the fixture: `disposition` (operator-offered) + `diagnostics` (offered to `*`) — the class target; `gate_resolution`'s case landed post-`67ee23e` (the instance, not the class, was fixed).
- Fixtures: m-2 §7.3 — every enum+offered token accepted subject to layer 3 (incl. `disposition`, `diagnostics`, `other` on gate_category) · the `genesis` negative (rejected at seat-scope — in the enum for compat, in NO scope) · a genuinely-unknown token bounces at membership.

## §10 A-4 + §D — the store lock (m-1 invariant · m-7 choreography; [VP-W2] dual-cite; kills F14)

- **[VP-W2] both halves cited:** m-1 §D owns I1-P (at most one live conductor per store root; claim-before-phase-0; loser refuses to serve INCLUDING reads; takeover only on proof-of-death; canonical root identity, not path strings). m-7 A-4 owns the choreography below.
- Mechanism (GRILLED — §18 row 3, LOCKED): **`flock(2)` on `<root>/conductor.lock`**, acquired `LOCK_EX|LOCK_NB` as **phase −1** (before anything phase 0 could later write); held for the serve lifetime; kernel releases on ANY process death (kill -9 included) ⇒ proof-of-death is kernel-bound by construction, no staleness heuristic, no stealable-live-writer window; alias-safety by construction (two paths to one root reach one inode — the flock target — satisfying canonical-root identity; the alias fixture proves it). Lock-file CONTENT is diagnostic only (holder pid + start time), written after acquisition — never the authority (the flock is). Loser: typed, path-free refusal naming the holder identity + the documented operator remedy, then full exit (no read serving). Takeover after holder death = normal acquisition + full recovery; the takeover logs a store-visible diagnostics record (auditable).
- D4/D5 honesty: stops accidental double-serve and confused ops; a malicious same-uid process can delete/bypass any lock — the accepted residual, stated wherever the lock is described.
- Fixtures: FX-A4a (race ⇒ one serves; loser typed + path-free) · FX-A4b (kill -9 ⇒ takeover ⇒ recovery converges; auditable record) · m-1 §E legs (alias paths don't bypass · no two live claims in any interleaving · the s4-gate-day leftover scenario replayed ⇒ the leftover refuses).

## §11 D-2 — bounce-detail parity, engine-side (m-7 D-2; kills F3) · D-1 rides §13

- The typed `Outcome` carries the same `bounce.Format` output for EVERY rejection class: `Outcome.Detail` (new field) = the rejected record's Body, byte-equal (re-render, field classes, lineage classes, layer-3, admission `boot-required`, internal-fault — one reply-shaping path in the loop/handler; no per-class branches anywhere).
- The shim's re-render record-read pattern-match hack (`mcp.go:221-241` + `reRenderResult`) RETIRES: the shim consumes the typed detail field; the `tools/list_changed` schema-refresh trigger keys on the typed `form_digest:re-render` class in the detail — a display consumer, never a second judge.
- Fixture: FX-D2 — per rejection class, reply detail ≡ recorded detail byte-for-byte (path-freedom inherited from `bounce.Format`, already fixture-proven).

## §12 B-1/B-2/B-3 — lifecycle, boot form, derived activation (m-7 B-1 r5 · m-2 §11 · m-1 §F/§F.1)

- **Lifecycle derivation (no persisted marker — [VP-W3]):** `tables` gains per-seat lifecycle: generation = the seat's latest committed `seat_mint` pivot (genesis-seeded seats = generation 0 from the binding table); `active` ⟺ an accepted governed submit stamped `FROM=<seat>` exists within the current generation (m-1's order rule — an ORDER fact, shape-independent, registry-drift-proof). Derived in `OnCommit`/`Build`; recovery re-derives at phase 3 (the existing records scan). `bound_now` = the channel server's per-credential active index (runtime only; empty at open — truthful post-restart state). Operator seat always-active by construction (no fixture ever demands an operator boot).
- **B-1.2a admission (the literal allowlist, verbatim from the set):** pre-active seat's submit accepted iff header set ⊆ {conductor-stamped envelope: FROM/ROLE/relay_id/DISPATCH_ID/intake_id/timestamp/schema_version/delivery_state} ∪ {SUBJECT} ∪ {`charter_loaded`, `dispatch_status`}, all boot-requireds present+valid, ZERO other fields registered or not; NO authority-bearing or gate header in the set. Fails ⇒ ONE terminal `rejected`, class `boot-required`, per-field detail `<field>:non-boot-before-active` (§11 parity) — typed and terminal, never a retry loop (m-1 §F.1 liveness constraint cited). Accept-time activation classification is TRANSIENT runtime state only (B-1.2b r5); recovery re-derives by the order rule. An already-active seat's boot-shaped submit = ordinary accept, no second edge (first-accepted idempotent by construction).
- **B-2 boot form:** pre-active render serves SITREP + `charter_loaded` (bool, `self_reported`, not a gate input) + `dispatch_status` (enum `{read, awaiting}`) as the rendered/required set; once active, ordinary forms. Un-bounceable by construction post-A-1/branch-A (the required set IS the rendered set; no parent field; stable digest). No new PHASE/record_kind/§J2/atom/verb (the shared-vocab negative fixture).
- **Roster (GRILLED — §18, LOCKED at the seven fields):** `project(view=roster)`, seat-scoped to operator + orchestrator through the existing scoping machinery; rows = {seat · `activation_state` ∈ {minted, active} · `bound_now` · role · minted-at · activation record ref · last-accepted-at} — state-not-traffic (no pending-mail counts), path-free, no credential material, no socket paths. Non-privileged seat ⇒ typed refusal. `project`/`read` NEVER lifecycle-gated.
- Fixtures: FX-B1a..FX-B1g (all seven, incl. the B1e smuggle negatives and the B1g generation leg) · m-2 §11.5 three (boot renders-pre-active + un-bounceable · shared-vocab negative · self_reported label) · **m-1 §F.6's four activation fixtures land via FX-B1c/B1f/B1g — stated here so §F.6 cannot fall between gate families** (the reconciled-audit mapping note).

## §13 D-1 — shim transparent reconnect (m-7 D-1; kills F16)

On a failed call over a previously-live connection: one transparent re-dial + re-auth with the held credential, then retry THE SAME call once; only a second failure surfaces the typed class (`shim:conductor-unreachable`). Composition stated: §5's idempotent replay makes the submit retry safe (a duplicate submit replays the outcome; never a second effect). Ratified F-GATE-3 boundary unchanged (shim holds credential + socket path for its session lifetime). Fixture: FX-D1 (conductor restart ⇒ the seat's next single call succeeds).

## §14 The registry pass (one, m-2-guided; MINOR; no migrator — R-1 held)

1. **The seven transport rows** ([VP-W3] EXACT): `parent_hint` · `parent_hint_honored` · `parent_provenance` · `routing_ref_honored` · `rationale` · `waiver_scope` · `retracts`. NO activation-marker row (grep-verified absent in both audits; the negative stays fixture-guarded).
2. Record classes: `record_kind` enum += `waiver_retraction`, `seat_mint` (both operator-only seat_scope). `ORCH_REVIEW_WAIVER` row retired. `PARENT_DISPATCH_ID` → system-computed (§4).
3. **B-2 boot fields** (`charter_loaded`, `dispatch_status` + its named enum): the set's own separately-named additive-MINOR vocabulary impact (m-2 §11.2 verbatim). **Reading GRILLED and CONFIRMED (§18):** [VP-W3]'s "exactly seven rows" fences the TRANSPORT row set against the activation-marker's reintroduction; the B-2 boot fields are a distinct named obligation of the same co-signed set; the guarded negative stays absolute (no activation-marker row). Restated in the design-completion relay for orchestrator + m-2 fidelity visibility — never silently resolved.
4. `routing_ref_honored` semantics: the generic hint-class validation (§4's one path, second consumer) over `lineage_role: routing_ref` references — validated against the accepted graph, honored-flag stamped, never a bounce; carriage per m-1 §6 (rewritten) = m-4 §5:216, the m-4-certified mapping. Dormant until a routing_ref-class field is live post-retirement; the row + mechanics land now (the m-4 condition folded verbatim).

## §15 Seam boundary contracts (each half stated; consume, never re-own)

| seam | Writes | Reads | Contract | Proof |
|---|---|---|---|---|
| Parenting 3-way | m-7 loop stamps PARENT/hint flags at the §4 locus | seat turn context (accepted-state, §6); `parent_hint` lane field | m-1 defines stamp semantics; m-2 un-renders PARENT + shapes the three hint rows; m-7 maintains turn context | GRILL three + §7.4 dissolution fixture |
| F14 dual | m-7 writes the lock artifact + takeover diagnostics record | m-1's I1-P invariant + lock-content semantics | [VP-W2]: invariant m-1 / choreography m-7 — cited in §10 both ways | FX-A4a/b + §E lock legs |
| F12↔F17 | m-1 record class + gate effective-state; engine rejects non-operator waiver rows | m-2's row typings (`rationale`/`waiver_scope`/`retracts`) | neither locks alone; the render-absence expressibility question is flagged TO m-2 (§8) | §E waiver legs + §7.5 |
| Envelope.To + digest | m-7 commit loop stamps `Envelope.To` (optional compat) | m-2's projection rule: decoded header list = recipient truth for EVERY projection | §2.4 rev1 verbatim; A-1 keeps volatile candidates out of the digest | §7.1 full-set fixture + FX-A1a |
| Boot | m-7 B-1 transient classification + roster; m-2 B-2 form; m-1 §F owns activation semantics | m-1 §F.1 derives-only list consumed verbatim | derived-only, no marker, no new vocab; boot sequences, never grants | FX-B1a..g + §11.5 |
| Re-mint (NEW, ratified) | derived work replaces binding row; closes old-credential channel | committed `seat_mint` pivots (generation order) | within m-1 §F/§F.1 + §13.3 carry; PROPOSED §7, grilled, m-1 fidelity rules | FX-B1g + §7 legs |

## §16 The fixture table (complete obligation map; red-first is the build discipline)

| id/family | asserts | module home | spec |
|---|---|---|---|
| G-1 | archived dogfood traffic (41 records) re-driven ⇒ completes, zero anchor livelock — **the recorded-pattern leg** | `test/replay` (+ the env-gated zeroloss leg composes, never duplicates) | GRILL_LOCK · m-1 §E |
| G-2 | hint honored + hint fallback, `parent_hint_honored` audit-visible, verbatim hint preserved | fixtures/engine | GRILL_LOCK |
| G-3 | concurrent accept during submit ⇒ NO parent-class bounce exists — **the race-class leg** | fixtures/engine | GRILL_LOCK |
| FX-A1a/b | zero re-render on foreign accepts / §7 change still bounces | fixtures | m-7 A-1 |
| FX-A2a/b/c | replay byte-identical / in-flight coalesce / id uniqueness incl. GC+restart leg + 1:1 sweep | intake + fixtures | m-7 A-2 |
| FX-A3a/b | live mint zero-drop zero-rotation / crash-window derived completion | fixtures | m-7 A-3 |
| FX-A4a/b | race one-server + typed loser / kill-9 takeover + auditable record | fixtures | m-7 A-4 |
| FX-D1, FX-D2 | one-call reconnect / per-class reply≡record detail | mcp + fixtures | m-7 D-1/D-2 |
| FX-B1a..g | lifecycle/roster/admission/derivation/generation (B1g = [VP-W2]; **r3: B1g gains the IN-FLIGHT leg** — old-session boot form queued pre-pivot ⇒ `credential-superseded`, does NOT activate; the new credential's boot does) | fixtures + tables | m-7 B-1.5 + R1(iii) |
| R1 reject (r3) | any command tagged with a superseded auth generation ⇒ typed `credential-superseded`, path-free, detail parity; recovery replay of tagged commands byte-identical | engine/intake fixtures | m-7 R1(i–ii) |
| R2 matrix (r3) | the `seat_mint` accept-reply family: credential + endpoint appear ONLY there (operator channel); never in record Body/projection/log/non-operator reply; the accept-reply⊃record divergence stated deliberate | s6_iph fixtures | m-7 R2 |
| M1-4 guard (r3) | duplicate-outcome append attempt at the commit boundary ⇒ replay-or-fault, zero second outcome records (the guard, distinct from the sweep) | engine fixtures | m-1 F-S6-M1-4 |
| m-1 §E | polluted-archive filter (rebuilt store) · audit scope · reject-never-anchor · waiver scope/retraction/re-arm/non-operator/legacy · lock race/alias/kill-9/typed-refusal | store/recover/fixtures | m-1 §E |
| m-1 §F.6 | → lands via FX-B1c/B1f/B1g (stated mapping) | — | m-1 §F.6 |
| m-2 §7.1-7.6 | codec full-set · archive one-encoding · three-layer + genesis negative · parent dissolution · waiver-as-record · recipient_picker exempt | fieldspec/store/fixtures | m-2 §7 |
| m-2 §11.5 | boot renders-pre-active + un-bounceable · shared-vocab negative · self_reported label | fieldspec/fixtures | m-2 §11.5 |
| floors | full uncached battery zero-regression · enum grep (`sweep_test` extended over new outputs) · I-PH grep-matrix over the four new families · three-verb surface | battery/sweep | ROADMAP gate 2 |

Step-exit (gate 3) is the operator's, on the fixed conductor — prepared-for (ops: pre-allowlist `mcp__frank__*`), never simulated.

## §17 Ordering / decomposition (the PLAN locks this)

1. **Codec + F13 reduction** (§2/§9) — the single-decode foundation; smallest diff that kills a SEV-1 pair (F6/F7).
2. **Registry pass** (§14) — one m-2-guided edit; rows land dormant-tolerant for later stages.
3. **A-1 digest class** (§3) — prerequisite for every liveness fixture.
4. **D-2 parity** (§11) — early, because §4/§12 rejection fixtures assert its detail shape.
5. **Branch-A parenting + §B anchors/projection** (§4/§6) — the core liveness cluster (F11/F4/F10), on 1–4.
6. **A-2 intake** (§5) — replay/coalesce/counter (enables D-1's retry-safety composition).
7. **§C waivers + retraction** (§8).
8. **A-4 lock** (§10) — independent; parallel-safe anytime after 2.
9. **A-3 live mint + re-mint** (§7) — after 5 (stamp/turn-context stable) and the grill.
10. **B-1/B-2/B-3** (§12) — needs §7 generations.
11. **D-1 shim** (§13) — after 6.
12. **G-1 archived-traffic replay + step-exit prep** — last; the whole-set integration proof.

## §18 GRILL_LOCK — the s6-core build-decomposition grill (folded r2)

```text
GRILL_LOCK_ID: s6-grill-s6-core
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relay read: the r2 DESIGN dispatch (agenda floor rows 1–4); both reconciled s6-core audits (the two ratified Named items); the co-signed set + constituents (resolved-row fence below)
- code/docs inspected: seat/binding.go Mint (ErrSeatAlreadyBound — no remint path); intake writer.go/journal.go (in-memory counter; GC-unsafe derivation); gc.go Pass (segment removal); channel server.go active index (force-close locus); store/recover (no lock machinery; phase 0–4 order); syscall.Flock availability darwin+linux
- questions answered from codebase (not asked): the hint-provable set + generous prover (delivered ∪ own-accepted ∩ accepted-graph; dispatch-id → thread root) — ergonomics per the GRILL_LOCK-parenting note, fallback semantics independent; D-2 Outcome detail carriage shape (designer's call, §11); waiver-row render-absence expressibility — an m-2 grammar question, ROUTED to the m-2 fidelity packet (§8), not operator-grade
- questions asked operator: 6, one at a time, recommendation-first (2026-07-07 in-session)

Resolved decisions:
- Re-mint mechanics — A: binding replacement — one seat_mint submit re-mints an existing seat; derived work REPLACES the binding row (fresh credential); old credential dies at auth Resolve; any live channel authed on it is force-closed at completion; binding table stays one-row-per-seat; generation history = the committed seat_mint pivots (never a persisted counter); crash-window credential retrieval = documented admin-time read of the 0600 binding table. Ordinary model/session switching is UNTOUCHED: same credential ⇒ reconnect = re-bound, no re-boot, no generation change. — source: operator (recommendation adopted)
- F11 archived-replay grain — sequence re-drive, two-legs-one-claim: the archived leg re-drives the recorded submit sequence deterministically and claims exactly "the recorded dogfood pattern completes with zero parent/digest bounces"; the race class is proven by the synthetic concurrent-accept fixture (G-3); neither leg overstates. — source: operator
- Store-lock artifact — flock(2) on <root>/conductor.lock, LOCK_EX|LOCK_NB, acquired phase −1, held for serve lifetime; proof-of-death kernel-bound by construction (auto-release on any death incl. kill -9); alias-safe by construction (one inode); file content = diagnostic only (holder pid + start time), never the authority; loser = typed path-free refusal + full exit (reads included); takeover = normal acquisition + full recovery + auditable diagnostics record. — source: operator
- A-2 durable-counter shape — segment header line: each new journal segment begins with {"segment_header":true,"high_water":N}; restore next = max(all headers, all entry ids)+1; survives GC because every new segment carries the mark forward; legacy headerless segments tolerated. Monotonicity retained as locked (ordering is load-bearing: arrival-order re-enqueue + trail legibility). — source: operator
- Roster payload — the B-1.3 seven fields exactly: seat · activation_state{minted,active} · bound_now · role · minted-at (pivot ref) · activation record ref · last-accepted-at; state-not-traffic (no pending-mail counts); path-free, credential-free. — source: operator
- [VP-W3] reading — seven = the TRANSPORT row set; the B-2 boot fields (charter_loaded, dispatch_status + its enum) are the set's separately-named additive-MINOR obligation; the guarded negative stays absolute: NO activation-marker row. Restated in the design-completion relay for orchestrator + m-2 fidelity visibility. — source: operator

Rejected alternatives:
- Revoke-then-mint two-record dance — two pivots with a credential-less crash window between; wants a third record class ([VP-W3] allows exactly two)
- Persisted generation column in the binding table — duplicates derivable truth (the pivots ARE the generation order); the persisted-copy-can-disagree defect class; m-1 §F.1 marker/field boundary forbids it without route-back
- Timing-faithful archived replay — archive timestamps cannot reconstruct interleaving; flaky by construction; the race is covered structurally by G-3
- Archived leg as the only F11 proof — a sequence re-drive cannot prove race extinction; the claim would overstate
- Pidfile + liveness probe — named insufficient by m-1 §D.3 verbatim (pid reuse; probe-then-take races a live writer)
- Lock via store record — circular (claim must precede phase-0) and append-only cannot release on crash
- Sidecar counter file — second durable artifact with its own torn-write story; diverges from the locked segment-header wording
- Records-derived id max — under-counts exactly the crash-mid-flight ids
- Random/content-derived intake ids — breaks the locked "durable monotonic counter" (ordering is load-bearing); would be a contract amendment for no gain
- Roster + pending-mail counts — state-not-traffic (B-1.3's F10 wisdom); derivable per-seat when wanted
- Minimal three-field roster — under-delivers the locked B-1.3 shape (drops the audit anchor + staleness)

Still operator-owned:
- none — all six grill rows resolved; the two flagged domain questions route to fidelity packets by domain ownership (waiver-row render-absence → m-2; re-mint binding semantics → m-1), which is packet routing, not an unresolved operator item

Design-lock impact:
- §7 re-mint mechanics locked as binding-replacement (fidelity: m-1 §13.3 carry); §10 mechanism locked as flock/phase −1; §5 counter shape locked as segment header; §12 roster locked at seven fields; §16 G-1/G-3 wording locked two-legs-one-claim; §14.3 reading confirmed (restate at completion). DESIGN_LOCK_ID for the s6 PLAN references GRILL_LOCK_ID: s6-grill-s6-core.
- Grill fence held: no master-grilled or pair-locked contract re-opened; resolved rows entered as resolved (branch A · fallback-never-bounce · precision accepted · derived-only activation, no marker · the seven transport rows · the codec thesis · the A/B/C/D + B-1/B-2/B-3 pair-locked contracts).
```

## §19 The ten dispatch constraints — compliance map

| # | lands at |
|---|---|
| 1 promote-never-rebuild | §2 (ParseTyped) · §3 (DigestExempt) · §5 (OutcomeByIntake/ContentHash/writer.hashes) · §7 (seat.Manager, config_change idiom, derived-work) · §8 (disposes_owed idiom, tables.Waivers) · §12 (active index, redacted-view pattern) · §16 (crash/replay harnesses, archive leg) · §4 (seat.Stamp idiom) · §11 (bounce.Format single formatter) |
| 2 F9 at writer grain | §5 (all four legs incl. GC+restart in FX-A2c; dead path deleted) |
| 3 §B rebuild path | §6 (serve-time filter + rebuild hygiene + accepted-state WokenOn; rebuilt-store fixture leg) |
| 4 F13 live tokens | §9 (disposition/diagnostics named; class not instance; genesis negative) |
| 5 D-2 engine-side | §11 (Outcome.Detail every class; shim hack retired) |
| 6 re-mint decomposed | §7 + §15 re-mint seam (PROPOSED → grill row 1 → m-1 packet) |
| 7 F11 two legs one claim | §16 G-1/G-3 wording; zeroloss composition stated |
| 8 fixture table complete | §16 (incl. the §F.6 mapping statement) |
| 9 claim pins + no-perf | §0 (pins beside surfaces in §3/§4/§5/§7/§10/§12; no-perf fence named with the two hot-path sites) |
| 10 threat points by name | §0 + §2.3 (invariant-fault not silent) + §6 (held off default) + §12 (roster = param; typed refusal) + §16 floors row |

*Rev log: r1 provisional pre-grill → r2 GRILL FOLDED (§18 `s6-grill-s6-core`; six operator rows resolved 2026-07-07) → r3 EXTERNAL-VERDICT FOLD (m-7 R1 §7/§16 + R2 §0/§16 verbatim per the pre-concurred must-revise `s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324`; m-1 F-S6-M1-4 §5; directed by `s6-core-design-verdicts` 012809; fresh pair DESIGN-REVIEW required per the content-changing-fold rule).*
