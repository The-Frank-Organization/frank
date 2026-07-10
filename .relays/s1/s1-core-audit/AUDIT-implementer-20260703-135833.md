## AUDIT - s1-core.implementer independent Slice-1 audit

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s1-core-audit-implementer-audit
PARENT_DISPATCH_ID: s1-core-audit
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for audit; downstream S1 plan remains m-7 guide + master VP + m-1/m-2 fidelity gated before implementation
FROM: s1-core.implementer
TO: s1.orchestrator-planner
CC: s1-core.planner, s1.orchestrator-reviewer, operator
IN_REPLY_TO: s1-core-audit/AUDIT-orchestrator-planner-20260703-133046.md; s1-core-audit/SITREP-orchestrator-planner-20260703-134525.md; s1-core-audit/SITREP-orchestrator-planner-20260703-135203.md
SUBJECT: Implementer read-only audit - S1 spec map, frozen contracts, owed carries, replay corpus

Phase:
Read-only AUDIT. I inspected locked spec surfaces and the current `frank/` repo state. I made no source/test/sprint-spec edits, no branch, no commit, no PR, no prototype, and no implementation. No valid implementation token exists for this seat in this run.

4-bucket verdict:
PRIMARY_BUCKET: still-open
still-open: S1 is not built in `frank/`; the repo has only sprint/relay scaffold commits and no conductor source implementation.
already-closed: no complete S1 conductor exists in this repo. The upstream package provides a reusable relay-lint fixture corpus and prior-art linter behavior, but it is not the `frank` conductor.
product-overlapped: overlap is limited to ported principles and replay corpus from upstream plus the frozen m-1/m-2/m-7 contracts; do not copy the old relay-lint system as the product.
recommended-next: proceed to DESIGN for the thin S1 conductor, with explicit fixture names, MVP FieldSpec subset, replay-subset selection, and m-1/m-2 fidelity-review objects before any implementation dispatch.

Duplicate/already-built gate:
- Current repo evidence: `git log --oneline --decorate --max-count=5` shows `461fae0`, `acd34b4`, `43e513e`, all sprint/relay scaffold or roadmap/reconcile commits; `find . -maxdepth 3 -type f` showed only `.git`, `.gitignore`, and `.relays/s1/INDEX.md` in the shallow repo surface. No source tree or conductor code is present.
- Baseline/corpus evidence: `<the archived upstream release>/` contains the old release zip, unzipped package, `tools/relay-lint.py`, and `tools/relay-lint-fixtures/`. This is a replay/prior-art source, not a build artifact to promote as the S1 engine.

Spec-to-exit-gate map:
- Accepted only through conductor: source `ROADMAP.md:64-67`; master dispatch `PLAN-orchestrator-planner-20260703-130634.md:59-60`; m-1 `submit()` contract, the m-1 trust/identity design-of-record (2026-06-28) :125-128; m-7 submit pipeline, the m-7 conductor-core design-of-record (2026-07-01) :68-76. Fixture idea: direct file append or non-conductor submit is absent/rejected; only a conductor `submit()` produces an accepted canonical record.
- System-stamped `FROM`: source master dispatch `:31-34`; m-1 resolve/stamp and system fields `:126,137-141`; m-7 pipeline `:68-69`; seam matrix NF-S1 `:151`. Fixture idea: payload-supplied `FROM` and `ROLE` are ignored byte-for-byte; unbound channel rejects before staging.
- Form/lint before delivery: source `ROADMAP.md:66`; m-2 send flow `:71-76`; m-1 pre-append validation `:127-128`; m-7 pipeline `:70-74`. Fixture idea: invalid required-set/enum/lineage produces terminal `rejected` plus author bounce before any recipient delivery.
- Gate to local outbox item: source master dispatch boundary `:48-56` and scope `:38`; `ARCHITECTURE.md:439`; m-7 local-outbox design `:130-132`; `ROADMAP.md:66`. Fixture idea: one minimal gate produces a local outbox/ODB item through the conductor-governed path, with D5 residual scoped in any "only egress" wording.
- Forged-FROM rejected: source m-1 `:126,138`; m-7 NF-S1 `:151`; m-2 system-owner model `:33-41`. Fixture idea: hand-crafted payload `FROM=victim` under attacker channel yields stamped attacker `FROM`, or rejects if channel unbound.
- Forbidden enum absent then rejected: source m-2 seat-scoped enum/render semantics `:35,41`; m-7 NF-S3 `:153`. Fixture idea: schema introspection shows forbidden options absent; raw submit supplying the forbidden option rejects at authoritative validation.
- Invalid parent rejected: source m-1 parent_picker `:126-127,139`; m-2 lineage gate `:75-76`; m-7 S4/NF-S4 seam `:154`. Fixture idea: free-typed parent outside conductor candidate set rejects before delivery.
- Duplicate-sibling double-accept killed: source `ARCHITECTURE.md:451-454`; m-7 serialized-loop claim `:63-64`; m-7 NF-S2 `:152`. Fixture idea: two concurrent verifier/authority consumers for one nonce/decision produce exactly one accepted operator-verdict/effect and one loser outcome.
- Crash matrix around commit/delivery: source `ROADMAP.md:67-71`; `STEP-1-KICKOFF.md:42-48`; m-7 durable FIFO `:56-58`, recovery `:87-95`, pivot `:78-84`, fixtures F3/F4/F9/F10/F11 `:172`. Fixture idea: fault-injection points at post-intake fsync, pre/post record fsync, pre/post rename, pre/post dir-fsync, projection write, delivery write, replayed `intake_id`; recovery proves exactly-once outcome, no stale re-emission, canonical projection rebuild.
- S1-minimal dissolved-linter replay: source `STEP-1-KICKOFF.md:45`; master dispatch `:63`; release changelog `the upstream release changelog :25-30`; fixture checker at `tools/check-relay-lint-fixtures.py:17-164`. Fixture idea: select the subset of historical upstream relay-lint failures covered by the MVP FieldSpec and assert each is caught by form validation or marked genuinely obsolete.
- I-PH path hygiene: source `ARCHITECTURE.md:461-463,482`; m-7 absence set `:124-125`; m-7 NF-S18 `:168`; `ROADMAP.md:70-71`. Fixture idea: grep every seat-facing tool description/result, bounce, error, projection, and delivery payload for canonical store/config/outbox/operator-channel paths and config values; zero hits.
- Liveness inbox durable, pipe write as nudge: source `STEP-1-KICKOFF.md:47`; m-7 delivery/wake `:121-122`; m-7 recovery `:92-95`. Fixture idea: busy/dead seat misses a pipe write, reconnects, and receives via `project()`; lost wake is repaired by projection/re-issued nudge.
- Park/wake: source `STEP-1-KICKOFF.md:48`; m-7 park/wake execution `:142`; m-7 NF-S13 `:163`; `ROADMAP.md:71`. Fixture idea: gated lane parks without consuming authority, verdict commit precedes wake, restart restores parked state and wake does not precede committed verdict.

Frozen contract surfaces consumed:
- m-1 store/API: `mint_seat` is conductor-only and binds seat to connection/credential (`m-1 design:124`); `submit` is the sole governed write path and has no persisted submitted limbo (`:125`); `project` reads caller-addressed relays (`:130`); `read` serves immutable canonical records for lineage/migration (`:131`).
- m-1 envelope/stamping: `submit` resolves and stamps `FROM`/`ROLE`, fills `relay_id`, `DISPATCH_ID`, `timestamp`, `schema_version`, `certification=null`, validates `PARENT` via `parent_picker`, and validates `TO`/`CC` via `recipient_picker` (`m-1 design:126,137-145`).
- m-1 append/projection: accepted writes are one atomic append plus INDEX row plus mailbox projection; form/lineage fail writes terminal `rejected`; internal authority fault can produce `held` (`m-1 design:128,135`).
- m-1 boundary: m-1 owns store/stamp mechanism; m-2 owns field slots; m-6 owns address graph consumers; identity is not authority (`m-1 design:147-155`).
- m-2 FieldSpec envelope: canonical FieldSpec fields are `id`, `layer`, `owner`, `type`, `enum_set`, `gate_referenceable`, `seat_scope`, `required_when`, `visible_when`, `fill_constraints`, `consumers`, `lineage_role` (`m-2 design:47-67`).
- m-2 ownership/render: owners are `system`, `seat_scoped_enum`, `agent_enum_pick`, and `free_text`; hybrid `parent_picker`, `recipient_picker`, and monotonic fields let agents select only within conductor-supplied bounds (`m-2 design:33-41`).
- m-2 validate/lineage: render selects visible fields, submit runs form-validation and lineage over `persisted accepted graph plus candidate` before append; observe-as-send is reserved for Step-2, so S1 remains store + form + lineage (`m-2 design:71-76,95-97`).
- m-2 terminal/outcome fields: `delivery_state` is byte-exact `{accepted, rejected, held}`; `rejected` is canonical, `bounced` is not a value token; `held` is only for authority internal fault or CQ-2 fail-closed disposition and is not ordinary human-gate parking (`m-2 design:278-283,376`).
- m-2 gate/config watchpoints: `gate_category` is RAISE-only toward A with `known_A` as a floor raiser (`m-2 design:270-277`); model-identity fields are not gate-referenceable (`m-2 design:87-99,363-364`).

Typed owed-item records:
- `{owner: s1-core with m-7 guide review, source: master dispatch :71-72 + ROADMAP.md:35 + ARCHITECTURE.md:433-436 + m-7 design :115-125, target surface: conductor trusted-side tool registry and seat-facing schema/resource/tool-result surface, disposition path: DESIGN names the code-layer guardrail; PLAN includes NF-S18/G(i) registry enumeration exactly {submit, project, read} and absence of raw store/config/outbox/operator-channel paths.}`
- `{owner: s1-core with m-1/m-2/m-7 fidelity, source: ARCHITECTURE.md:463,482 + ROADMAP.md:35 + m-7 design :124-125,:168, target surface: every seat-delivered bounce/error/projection/delivery payload, disposition path: DESIGN defines canonical internal-path patterns; PLAN includes I-PH grep fixture with zero hits.}`
- `{owner: s1-core with m-2/m-6/m-7 fidelity, source: ARCHITECTURE.md:478 + STEP-1-KICKOFF.md:56-57 + ROADMAP.md:35 + m-2 design :277,:381, target surface: `gate_category`/known-A handling around the S1 gate/local outbox item, disposition path: DESIGN identifies the minimal known-A case in S1; PLAN includes negative fixture that a B-pick over known-A is raised to A and records `gate_category_raised=true`.}`
- `{owner: s1-core with m-3/m-6/m-7 clarification, source: ARCHITECTURE.md:479 + STEP-1-KICKOFF.md:56-57 + m-2 ODB model_name slot :297, target surface: local outbox/ODB item and egress scan seam, disposition path: DESIGN must clarify whether the S1-local outbox fixture covers only typed local ODB production or also the model-name egress exception; if full egress scan is required, escalate because observe/egress is otherwise later-layer scoped.}`
- `{owner: s1-core with m-2/m-4 fidelity if MVP FieldSpec touches routing-ish fields, source: ARCHITECTURE.md:473 + STEP-1-KICKOFF.md:56-57 + m-2 gate_referenceable fields :57,:87-99,:239, target surface: FieldSpec predicate grammar, disposition path: DESIGN either excludes model/routing fields from the MVP subset or includes a narrow negative fixture proving model-identity fields are not gate-referenceable.}`
- `{owner: m-2/possibly outside S1 unless orchestrator says otherwise, source: ARCHITECTURE.md:480 + STEP-1-KICKOFF.md:56-57, target surface: `GRILL_REQUIRED` FieldSpec row, disposition path: record as owed but do not absorb into S1 MVP unless a guide/orchestrator relay says the S1 form must carry that field.}`
- `{owner: CTO/m-2/m-4/m-6, not an S1 blocker per source text, source: ARCHITECTURE.md:481 + STEP-1-KICKOFF.md:56-57, target surface: explicit `routing_escalation` gate_category member, disposition path: keep as pre-wire carry; do not let it expand S1 into routing execution.}`

Replay corpus location:
- Release root located: `<the archived upstream release>/`.
- Executable fixture corpus located: `<relay-lint tools>/relay-lint-fixtures/`.
- Fixture checker located: `<relay-lint tools>/check-relay-lint-fixtures.py`.
- E2 read-only proof: `python3 .../tools/check-relay-lint-fixtures.py` observed every expected row PASS across the bundled matrix. The checker file enumerates the expected matrix at lines 17-164; the changelog records the upstream release validation as relay-lint fixture matrix `146/146 PASS` at `the upstream release changelog :25-30`.
- Audit finding: corpus exists and is executable, but the S1-minimal subset is not pre-selected. DESIGN/PLAN must map the MVP FieldSpec to the covered historical failures rather than running the entire full upstream dissolve as an S1 acceptance gate.

Design questions surfaced for DESIGN:
1. Runtime/language and process shape: locked docs require one conductor process, per-seat MCP channels, single intake-writer, one plain FIFO commit loop, concurrent immutable reads, and runtime wake adapters (`m-7 design:51-58,121-128,237`). They do not prescribe language. DESIGN should choose the smallest existing-runtime-friendly stack for `frank/` and name the per-seat channel/adapter boundary.
2. MVP FieldSpec field set: minimum must include the consumed envelope/system fields and slots named above, terminal `delivery_state`, `failing_edge`, and the S1 gate/outbox fields. DESIGN should exclude full registry/consumer-schema fields unless needed for a named S1 fixture.
3. Crash fixture mechanics: DESIGN should expose deterministic fault-injection points around intake fsync, record fsync, rename, dir-fsync, redo/projection, delivery wake, and replayed `intake_id`, with child-process `kill -9` tests.
4. `mint`/`connect`/park-wake minimum: m-1 specifies `mint_seat`/connection binding and m-7 specifies attach/wake semantics, but S1 must decide the concrete local credential/channel shape, reconnect behavior, and minimal parked-lane state record.
5. Owed-carry boundary: the code-layer guardrail, I-PH fixture, and known-A/RAISE-ONLY guardrail-adjacent fixture are clearly S1-facing. The ODB model-name egress, R2 per-column, `GRILL_REQUIRED`, and `routing_escalation` carries need explicit DESIGN classification as S1-direct, S1-recorded-only, or escalated/deferred before PLAN.

Boundary contract assessment:
- Writes: append-only relay-store records with byte-exact terminal `{accepted, rejected, held}` plus local outbox/ODB item. Supported by master dispatch `:48-56`, m-1 `:125-128`, m-2 `:278-283`, and m-7 `:97-105`.
- Reads: locked m-1 `submit`/`project`/`read` and m-2 FieldSpec envelope. Supported by m-1 `:122-131` and m-2 `:47-67`.
- Target entity: the conductor committed relay plus rebuilt projection. Supported by `ARCHITECTURE.md:423-426` and m-7 `:80-84,:91-95`.
- Downstream consumer: seat inboxes via `project()` and local outbox. Supported by master dispatch `:52-56`, m-1 `:130`, m-7 `:121-132`.
- Contract: viable for S1; no writer-without-reader or reader-without-writer found. Watchpoint is not consumer absence; it is scope containment so S1 does not absorb full FieldSpec, observe, routing, or recovery beyond the thin gate.

Scope-OUT re-check:
- Full FieldSpec registry, 62-check linter, and full approximately 33-check replay are explicitly out for S1 (`ROADMAP.md:73-77`; master dispatch `:41-42`; `STEP-1-KICKOFF.md:45`).
- Full recovery phases 0-4, durable FIFO, GC/genesis, and owed-item projection are out/thicken-later except for S1-crash fixtures around the thin path (`ROADMAP.md:73-77`; `STEP-1-KICKOFF.md:59-63`).
- Consumer schema fields, observe-as-send, and routing execution are out/later (`ROADMAP.md:73-77`; master dispatch `:41-42`; m-1 `:129,151`; m-2 `:75,95-97`).
- Any finding that appears to require these OUT items should be escalated to `s1.orchestrator-planner`, not silently absorbed.

Acceptance criteria status for audit:
- Exit-gate line items mapped to source and fixture ideas: complete for AUDIT.
- Frozen m-1/m-2 contract surfaces enumerated: complete for AUDIT, pending later fidelity review over DESIGN/PLAN artifacts.
- Owed §C4 carries typed: complete for AUDIT, with classification questions for non-obvious S1-direct carries.
- Replay corpus located: complete; S1-minimal subset still a DESIGN/PLAN task.
- Scope-OUT checked: complete; no mandatory OUT absorption found at audit time.

Evidence levels:
- E1: cited locked docs and relays by file/line.
- E1: repo shallow state inspected by git log/find.
- E2: incoming dispatch/supplement/release relay lint passed.
- E2: upstream fixture checker ran read-only and all expected rows observed PASS.

Actions:
- Wrote this read-only AUDIT relay.
- Appended the corresponding routing row to `.relays/s1/INDEX.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits claimed; relay artifact written at `.relays/s1/s1-core-audit/AUDIT-implementer-20260703-135833.md`; index row appended at `.relays/s1/INDEX.md`; final git status captured below.
FINAL_GIT_STATUS_SHORT: none - clean tree

Done:
- Independent implementer AUDIT completed against the addressed dispatch and supplements.

Not done:
- No DESIGN decisions made.
- No PLAN, implementation, branch, PR, merge, or live verification started.

Blocked:
- Not blocked for DESIGN. Open questions above should be answered during DESIGN or routed to the guide/orchestrator if they require lock interpretation.

Scope drift risk:
- Medium for PLAN if the owed-carry ledger is treated as permission to build full FieldSpec, observe/egress, routing execution, or full recovery. Keep S1 to thin path plus named S1 fixtures.

Tests / verification:
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-audit/AUDIT-orchestrator-planner-20260703-133046.md` passed.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-audit/SITREP-orchestrator-planner-20260703-134525.md` passed.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-audit/SITREP-orchestrator-planner-20260703-135203.md` passed.
- `python3 <relay-lint tools>/check-relay-lint-fixtures.py` passed observed expected matrix rows.
- `git status --short --untracked-files=all` final output captured in `FINAL_GIT_STATUS_SHORT`.
- Pending final handoff check: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-audit/AUDIT-implementer-20260703-135833.md`.

Next requested action:
- `s1.orchestrator-planner` reconciles paired audits and dispatches DESIGN when ready.
