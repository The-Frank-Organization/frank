# s5-b Implementation Plan — mechanisms & versioning (③ raise · DEF-2 guard · ⑤ dormant egress · zero-loss replay · §7 s5-delta legs · I-PH)

> **For the executing Implementer:** execute under Superpowers `executing-plans`, task-by-task, in plan order. Every task ends with the battery green (`go build ./...` + `go test -count=1 ./...` + `go vet ./...`) and a commit (`s5-b:` prefix). Implementation authority arrives ONLY via the delegated dispatch relay under the F2 conditions — not from this document.

**Goal:** land the s5-b bundle of Slice-5 (consumer schemas): the ③ known-A raise at the validate locus, the DEF-2 submit guard, the ⑤ egress drain+scanner present-but-dormant behind the Renderer contract, the Q4 zero-loss replay + read-path refusal legs, the §7 legs over the real s5-a registry delta, and the I-PH sweep extension — zero regression, byte-exact terminal enum, guardrail surface untouched.

**Locked design:** `docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md` **r3** (DESIGN_DOC_ID `s5-b-mechanisms-design`; approving review `DESIGN-REVIEW-implementer-20260706-060550.md`). §n references below are design sections. **The design rules; this plan sequences it.** ③ semantics of record: `master/relays/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` §2.

**Build surface:** branch `s5-b-mechanisms` off `main @ 67ee23e` (operator-cut worktree). Baseline: 21 packages green (audit E2).

## Global constraints (bind every task)

- Guardrail surface stays exactly `submit`/`project`/`read`; terminal enum byte-exact `{accepted, rejected, held}` — no new value anywhere.
- **registry.json / registry_test.go / render_test.go / validate_test.go registry-content fixtures are s5-a's — NEVER touched** (R-s5-2). The `gate_category_pick` + `resolves_gate` rows and the `routing_escalation` delta arrive via s5-a; tasks that consume them are the sequenced tail (T7).
- **③ claim boundary on every ③ surface** ([VP-W1], master §2 wording): detection claims exactly (S1)+(S2)+(S3)+fail-safe — NOT "catches every content mis-pick."
- **⑤ honesty:** no code/doc/tool text claims live scanning; present-but-dormant + D5-qualified phrasing travels verbatim (§0).
- I-PH on every new seat-deliverable byte: Field:Class only, no path/config/store-layout token, formatter untouched.
- Run-FAIL-first on every negative fixture leg (guard rejects, raise negatives, egress blocks, refusal legs).
- No transport-fix work (lineage/parenting/codec); no live egress activation; no live `record_kind` widening; archived store propose-only.
- **Live conductor wiring of the ③ detector config (cmd/frank/main.go) is NOT in this plan** — the config shape binds at ③ IMPL-integration at the orchestrator's gate when the m-6.implementer signal-set confirm lands. This plan builds the mechanics + engine API + fixture-injected detector; `cmd/` is untouched.

## File map (the SCOPE_DIFF universe — complete)

| surface | files |
|---|---|
| ③ + DEF-1 + DEF-2 mechanics | `internal/fieldspec/validate.go`, `internal/fieldspec/render.go` (RenderEnv type only), `internal/fieldspec/fieldspec_test.go`, `internal/fieldspec/validate_test.go` (mechanics-grain additions only — no registry-content fixtures), `internal/engine/submit.go` (env plumb if needed), NEW `internal/engine/detector.go` (+ `detector_test.go`) |
| ⑤ egress (new, dormant) | NEW `internal/egress/egress.go`, `internal/egress/render.go`, `internal/egress/rules.go` (+ package tests) |
| replay | NEW `test/replay/zeroloss/zeroloss.go`, `test/replay/zeroloss/zeroloss_test.go` |
| fixtures | NEW `test/fixtures/s5_submit_guard_test.go`, `s5_gate_raise_test.go`, `s5_egress_test.go`, `s5_config_change_test.go` (T7), `s5_iph_test.go` |
| **OUT (hard)** | `internal/fieldspec/registry.json`, `registry_test.go`, `render_test.go`/`validate_test.go` registry-content fixtures, `internal/bounce/formatter.go`, `internal/migrate/migrate.go`, `cmd/*`, `internal/lineage/*`, `internal/store/*` (the drain lives in `internal/egress`, not store) |

## Tasks

### T1 — DEF-2 submit guard (§2)
- [ ] `validate.go`: pre-loop guard — non-empty header whose spec is `owner ∈ {system, computed}` or `fill_constraints ∈ {system_only, computed_result}` ⇒ `Violation{Field: id, Class: "system-owned"}`; typed reject, never strip; headers only (envelope overwrite semantics untouched).
- [ ] `s5_submit_guard_test.go`: lane-supplied `failing_edge` / `gate_category_raised` / `delivery_state` reject; **operator-channel** submission rejects identically; conductor-internal writers (obligation/fault/genesis paths) unaffected (they bypass Validate — assert by driving one derived-record path). FAIL-first each leg.
- [ ] Battery green; commit.

### T2 — ③ raise mechanics + DEF-1 (§1.1–§1.3)
- [ ] `render.go`: `KnownADetector` type + `RenderEnv.KnownA` member (render behavior untouched).
- [ ] `validate.go`: move the raise out of the per-field loop into a post-loop step (B-absorb needs absent-pick handling): detector hit ∧ effective-below-A ⇒ REWRITE `gate_category` to the named A member (else `other`); stamp `gate_category_raised: "yes"` (**DEF-1: the `"true"` byte dies here**); preserve an actual original pick in `gate_category_pick` (value write only — the row is s5-a's); wire `ClassifyGateCategory`'s `knownA` from the detector (retire the hardcoded `false`); skip the index-based `belowMonotonicFloor` for the `enum_set=="gate_category"` row (the A/B map is the lattice — §1.3). No A→B branch exists.
- [ ] NEW `internal/engine/detector.go`: the MAX(S1,S2,S3) composite behind `fieldspec.KnownADetector` — S1 `AFloor(phase, record_kind)` + S3 `MergeSplit(fields)` against a `DetectorConfig` struct (shape ISOLATED in one constructor — binds at IMPL-integration on the m-6 confirm; fixtures supply fixture-scoped config); S2 = referenced-gate lookup over the tables snapshot (the `classifyVerdict` read pattern, `submit.go:216-245`); member precedence S2→S3→S1 (§1.2, review-cleared).
- [ ] `fieldspec_test.go` + `detector_test.go`: classifier-grain + composite-grain units (each source alone, MAX composition, precedence, nil-detector = fail-safe-only).
- [ ] Battery green; commit.

### T3 — ③ DEF-4 fixture (§1.4, legs 1–4 + 6)
- [ ] `s5_gate_raise_test.go` (claim-boundary doc comment, master §2 wording): (1) B-pick + detector-hit ⇒ committed A token + `"yes"` + park record + outbox item + `gate_category_pick` = original; (2) B-absorb ⇒ same, pick absent; (3) no-A→B negative; (4) `other`-pick end-to-end ⇒ `"yes"` (not `"true"`) and the committed record re-validates against its own registry; (6) claim-boundary negative — a mis-pick invisible to S1/S2/S3 commits as B. FAIL-first on 3/4/6.
- [ ] Leg 5 (full §J2 map incl. `routing_escalation`) is **T7** — not here.
- [ ] Battery green; commit.

### T4 — ⑤ egress package (§3.1 + §3.1a)
- [ ] NEW `internal/egress`: `Class`/`Finding`/`Verdict`/`Origin` (runtime-only — never serialized/parsed; package doc states it), `RenderedField{Name, Value, Origin}`/`Rendered{Dest, Fields}`/`Renderer`, `Item{Meta, Source, Dest, Field RenderedField}`; `Scan(item, rules)` classifying `item.Field.Value`; `Drain(st, rules, render)` walking `outbox/*.json` → `st.Read(source_record_ref)` → `render()` → scan every RenderedField pre-send; blocked ⇒ withheld + `egress_blocked` in Report (non-terminal; no new record_kind).
- [ ] `rules.go`: compiled-in Step-1 defaults — safety (secrets/PII/auth-URLs, never exempted) + confidentiality (model-name tokens; the single carve-out `item.Field.Name=="model_name" ∧ item.Field.Origin.ConductorODB ∧ item.Dest=="operator"`); fail-closed other→block.
- [ ] `render.go`: `DefaultRenderer` — every field lane-origin, **destination derivation named explicitly in code+tests from current source/item fields** (the design-review PLAN carry), exemption-free posture.
- [ ] Package doc: present-but-dormant + D5-qualified + claim-boundary phrasing (§0). Findings render Field:Class only.
- [ ] Battery green; commit.

### T5 — ⑤ acceptance fixture (§3.2)
- [ ] `s5_egress_test.go`: constructed store, outbox items via the REAL obligation path; every leg through `Drain(st, rules, render)` — **no direct `Scan(Item)` construction**: pass leg (ODB renderer stand-in generates `model_name`, `Dest=="operator"` ⇒ passes); (a) same model-name lane-origin blocks; (b) secret/PII/auth-URL from ODB record content blocks; (c) lane-mimicked exempt-mark renders lane-origin ⇒ blocks; fail-closed leg + DefaultRenderer-alone-exempts-nothing; dormancy assertion (no production package imports `internal/egress`). FAIL-first on (a)/(b)/(c)/fail-closed.
- [ ] Battery green; commit.

### T6 — zero-loss replay + refusal legs (§4)
- [ ] NEW `test/replay/zeroloss/`: `Replay(root, reg)` walking redo+records through `migrate.Reader` — zero imports of the parity harness/oracle.
- [ ] Constructed-store leg (mandatory): `store.Init` fresh + variety (accepted/rejected/held, gate → park+outbox, owed pair, config_change chain); assert count (all readable, zero lost, quarantine accounted) / identity (SourceVersion==Current ⇒ deep-equal) / canonical-wins (record beats projection).
- [ ] Refusal legs THROUGH `Reader.Read` over the real store: planted SchemaVersion 0 / future / gap files ⇒ `ErrUnversioned`/`ErrFutureVersion`/`ErrMigrationGap`. FAIL-first.
- [ ] Optional archive leg: `FRANK_S5_REPLAY_STORE` param, skip-if-unset, self-derived counts (M-4 pending; never touch the archive in place).
- [ ] Assert `migrate.Current == 1`; no migrator registered (R-1 discipline).
- [ ] Battery green; commit.

### T2b — classifyVerdict submit-path operator-seat enforcement *(added by `s5-b-sequencing/PLAN-orchestrator-planner-20260706-091736.md` §2 — acceptance sharpening, not new scope)*
- [ ] `internal/engine/submit.go`: `classifyVerdict` (:216-245) gains operator-seat enforcement for gate-resolution submissions per the settled `resolves_gate` shape (operator-seat-scoped, paired with `gate_resolution` record_kind), consistent with the (h) grains — typed reject, keyed on the submission channel, envelope asymmetry untouched.
- [ ] Negative fixture: a non-operator gate-resolution submission bounces AT THE SUBMIT PATH (not only un-rendered at fill time). FAIL-first.
- [ ] Battery green; commit.

### T7 — SEQUENCED TAIL: §7 s5-delta legs + ③ full-map leg (§5; R-s5-6)
**UNBLOCKED by `s5-b-sequencing/PLAN-orchestrator-planner-20260706-091736.md` §1:** s5-a is accepted at the orchestrator's gate at `s5-a-registry @ dd7d0b5` — consume that branch **READ-ONLY** (registry bytes + expected digests) now. Constraints: any hard-coded digest notes its source commit (`dd7d0b5`) so the re-point is mechanical; the consuming assertions **re-verify against merged main** before this pair's integration at the orchestrator's gate (byte-identical fast-forward expected ⇒ re-run, not rework; any pre-merge byte delta arrives via the orchestrator).
- [ ] `s5_config_change_test.go`: the five adapted legs with the REAL landed registry as config_change body — operator-shape · old→new digest · no-re-genesis **with genesis-count==1 assertion** · phase-0 **with a direct `st.ValidateGenesis` walk** · stale-form re-render where the fresh form shows `routing_escalation`.
- [ ] `s5_gate_raise_test.go` leg 5: full-§J2-map raise incl. `routing_escalation` as a named A member.
- [ ] Battery green; commit.

### T8 — I-PH extension (§6)
- [ ] `s5_iph_test.go`: capture points for every new surface — ⑤ scan Finding strings, drain diagnostics, DEF-2 `system-owned` bounce text, replay/refusal wrapper text — through the `assertNoPathFamilies` / `assertNoS4IPHLeaks` patterns. Formatter byte-untouched (assert no diff).
- [ ] Battery green; commit.

### T10 — ③ LIVE WIRING *(added per `s5-b-wire3/PLAN-orchestrator-planner-20260706-154556.md`; ungated by the m-6.implementer signal-set APPROVE 052907 + master GO 153721 — those three relays are the semantics record)*
**Branch: `s5-b-wire3` off `main @ b30df4d` (fresh, orchestrator-set).**
- [ ] `internal/engine/detector.go`: `DetectorConfigFromPinned(*config.Pinned) DetectorConfig` — parses an optional `"detector"` section from the pinned ENGINE member bytes (`Pinned.Members["engine"]`; unknown JSON keys are already ignored by `EngineConfig`, so `internal/config` stays untouched): `{a_floor: [{phase, record_kind, member}], target_branch_field, protected_branches}`. Load-once at startup (the §7 no-hot-reload discipline: a config_change to the engine member + restart re-binds). Every `a_floor.member` must ∈ the registry's `gate_category_A` set — load-fail loudly otherwise (the §J2-A-set-as-config read).
- [ ] `cmd/frank/main.go`: construct the detector **per submit** with the fresh tables snapshot (the panel's activation note — mirror the existing handler pattern at `main.go:111-116`): `env.KnownA = engine.KnownADetector(reg, tab, detCfg)` inside the handler closure. No other `cmd/*` change; the seat tool surface is untouched.
- [ ] **Named default 1 (S1):** shipped `a_floor` default = EMPTY table (the CQ-3 table is the documented config vocabulary — the map is m-6's surface, "the values the operator's", GRILL_LOCK G5); fixtures configure fixture-scoped rows to prove the live path.
- [ ] **Named default 2 (S3):** shipped `target_branch_field` default = UNSET ⇒ S3 inert (`s3MergeSplit` already fail-safes on empty field). Rationale: NO target-branch atom exists in the 83-row registry — config-pointing S3 at an undeclared header would recreate the DEF-5 canonical-iff-consumed violation class. S3 activates when its atom is declared in a future registry pass + the operator configures the field/set. **The claim boundary travels updated:** live detection = S1(as configured) + S2 + hardcoded `other`→A; S3 mechanism wired, input-atom-pending — this wording (plus the master §2 boundary) goes on every claim surface this wiring adds (code comment, config doc, fixture doc).
- [ ] NEW `test/fixtures/s5_wire3_test.go`: end-to-end live-path fixture — a store whose engine member carries a fixture-scoped detector section; submits through the REAL production handler construction (the `cmd` wiring pattern): (a) S1-configured hit raises end-to-end (committed A token + `"yes"` + consequences); (b) S2 hit via a live referenced gate raises; (c) detector-absent config ⇒ fail-safe-only behavior (today's semantics, unchanged); (d) a_floor member outside the A-set fails the config load loudly. FAIL-first on (a)/(d).
- [ ] Battery green (branch tree; branch is off current main tip — re-run combined if main moves); commit.

### T11 — M-4 archive-replay leg *(accepted into this hop by the same relay)*
- [ ] Run the zeroloss archive leg against the operator-placed copy: `FRANK_S5_REPLAY_STORE=~/frank-s5-team/replay-store-dogfood-20260706 go test -count=1 ./test/replay/zeroloss/` — **READ-ONLY on the copy** (the harness walks journal+records without writes by construction); counts derived from the copy's own journal (already the design — no hard-coding).
- [ ] Any unexpected incompatibility with the real-traffic store (41 records incl. rejects/bounces/waiver/config history) is a **FINDING to report, never a fix-in-place** on the copy; harness-side fixes are in-surface if needed.
- [ ] Record the leg's result (record count, verdicts) in the implementation report.

### T9 — closeout
- [ ] Full battery + `go vet`; zero regression vs the 21-green baseline; grep-sweep for forbidden claims ("live scan", detector over-claim) across every added file.
- [ ] Implementation report relay (branch@sha, per-task commits, fixture inventory, E2 evidence) TO s5-b.planner, CC orchestrator + reviewer.

## Acceptance criteria (design §9, restated as the gate)
1–7 of design §9 verbatim, minus the T7 items when integration has not yet been instructed (they gate s5-b CLOSE, not the T1–T6/T8 report). ③ IMPL-integration (live detector config + cmd wiring) is the ORCHESTRATOR's gate on the m-6 confirm — out of this plan.
8. *(added per `s5-b-sequencing` §2)* `classifyVerdict` enforces operator-seat at the submit path for gate-resolution submissions; the non-operator negative fixture green (T2b).

## Anti-half-fix guards (design §10)
DEF-1 ships only inside T2+T3 (never alone); DEF-2 rejects, never strips; the scanner ships with its dormancy assertion in the same task pair (T4+T5); the replay ships with the refusal legs (T6 is one task).

## Out of scope
Everything in the file-map OUT row; step-(d); transport-fix; live activation; `cmd/*`; render-side gate_category affordance; the archived store.

## Operator/orchestrator judgment items
None open at plan time. Escalation triggers: any need to touch an OUT file (esp. registry.json or cmd/) ⇒ stop, relay to s5.orchestrator-planner; s5-a integration timing (T7) and ③ IMPL-integration are orchestrator-gated.
