## PLAN-REVIEW — MUST-REVISE s12 build plan r1: the 17-task battery map is strong, but delegated dispatch is blocked by two routed-lint failures, two H-16-REG pin conflicts, omitted version-transition consumers, and one unruled rev20 rejection-class byte

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s12-build-plan-review
PARENT_DISPATCH_ID: s12-build-plan
RUN_ID: s12
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s12.planner can file the bounded successor and route the named owner/master escalations; the operator-only merge gate remains terminal and is not requested now
FILED_AT_LOCAL: 20260817-192631
PLAN_LOCK_ID: s12-h16-fix-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s12-build/PLAN-planner-20260817-191828.md
ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — credential mint, quarantine, recovery, and operator authority surfaces
- migration/backfill/destructive-write/canonical-data-repair: yes — legacy pivot upgrade anchors, redo completeness, and binding repair
- money/inventory/orders/planning/accounting/trust-critical-state: yes — canonical relay and credential truth are trust-critical state
- AI-or-automation-acts-downstream: yes — conductor recovery and retry machinery performs automated canonical/binding actions
- worker/scheduler/queue/retry/async-side-effect: yes — Class-G/Class-D retries, scheduler arming, and post-commit effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — fieldspec registry, Outcome wire shape, and auth refusal contract are shared interfaces
- user-visible-control-with-materializer/downstream-consumer: yes — offline ceremony and state-only consumers have materialized downstream effects
- test-runtime-role-mismatch: no — the plan uses real-process fixtures where role fidelity matters; no mismatch found in review
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — record-kind cardinality, one rejection byte, and routed lineage are unresolved
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk waiver is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-7.planner
SUBJECT: r1 must revise before delegated dispatch; exact bounded successor requirements and two route-ups

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`. The task decomposition, RED-first cadence, four-leg coverage, 48-case unit, 31-row regression, terminal merge boundary, and battery-to-task map are materially sound. The current artifact cannot enter delegated IMPL because the active-root structural gate is red and the scope/contract issues below would force either out-of-plan edits or locally invented authority-critical bytes.

## Blocking findings

### F1 — The active relay root is structurally red; prose about the external design does not satisfy the routed gate

The target declares `DESIGN_RECORD_KIND: design-doc` and a `PLAN_SHA256`/`PLAN_ARTIFACT` pair (`PLAN-planner-20260817-191828.md:14-18`), while explaining that no same-owner DESIGN/REVIEW exists in this root (`:30-36`). Exact-file lint with `--no-freshness` passes and the INDEX passes, but routed lint reports both:

1. no earlier same-owner `DESIGN` carrying `h16-outcome-split-design`; and
2. no `plans/PL-s12-build-plan-20260817.md` at any artifact probe root.

Those are dispatch-blocking lineage/integrity failures, not documentation nits. Required successor: make the external frozen-design adoption protocol-valid in this active root without s12 proxy-authoring or self-declaring a direct override, and place/link the plan artifact at a linter-resolvable `plans/` locus (or use another protocol-valid artifact layout). Run routed lint clean before any scope diff or dispatch.

### F2 — MUST-CHECK 1 and 2 exceed or contradict the exact H-16-REG commission

The commission binds **two** new `record_kind` tokens and the literal m-2-confirmed row shape `type: string` (`master/.../PLAN-orchestrator-planner-20260817-184653.md:29-33`). The owner confirmation repeats `type: string` (`DESIGN-planner-m2-20260817-133642.md:31-33`). Plan r1 instead adds **four** tokens and writes all three rows as `type: text` (`PL-...md:60-61,125-130`). The premise that every existing registry string carrier is `text` is false: live rows such as `slot_in`, `seat_archetype`, `decision_deadline`, and `model_name` use the supported `string` type (`registry.json:161-173`; `validate.go` and `canonical.go` accept both).

Required successor: do not locally convert two tokens to four or `string` to `text`. Route the token-cardinality/disposition-class interpretation UP to master with m-2 guidance, and retain literal `type: string` unless an owner-authorized successor says otherwise. Any resulting ruled file/scope delta must appear in the successor plan before dispatch.

### F3 — Task 4's marker bump omits mandatory live consumers and therefore cannot keep `-init` or `go test ./...` green

Task 4 permits only `internal/fieldspec/registry.json` and says nothing else moves (`PL-...md:123-130`). Live code accepts fieldspec reader markers only through `s10-fieldspec-v8` (`internal/config/config.go:317-328`), permits transitions only through `v7 -> v8` (`:425-436`), and downgrades fresh-genesis fieldspec bytes beginning at v8 (`internal/store/genesis.go:28-51`). Existing tests also assert the v8 marker and exact provenance map (`internal/fieldspec/registry_test.go:15-28`; `test/fixtures/s10_fieldspec_test.go:20-79`; the s8 historical inverse fixture is v8-shaped).

A registry-only v9 edit is therefore rejected by config preflight and is not covered by genesis compatibility; the current file fence omits at least `internal/config/config.go` and functional `internal/store/genesis.go` changes plus their asserting tests. Required successor: add the reader-supported marker, explicit `v8 -> v9` successor transition, v9-to-v8 genesis downgrade/preservation path, historical-fixture preservation, fresh-v9-init proof, and all affected exact-marker/provenance tests to the file structure, Task 4, batteries, and scope diff. Pin the provenance successor-relation shape rather than leaving an implementation-time conditional object shape.

### F4 — Rev20 requires a typed predecessor-mismatch rejection but supplies no byte; r1 silently leaves it for implementation

Rev20 says an implied `mint_predecessor` mismatch commits REJECTED "with a typed class" (`h16-outcome-split.md:89`), but its m-2 scope and confirmation enumerate only the four §4b classes plus `anchor-target-resolved` and `retry-authority-delta` (`DESIGN-planner-m2-20260817-133642.md:47-49`). Plan r1's exact-token list likewise contains no predecessor-mismatch class, while Task 10 still requires an unspecified typed class (`PL-...md:194-200`). Choosing that externally visible `failing_edge` byte during implementation would be a silent locked-contract decision.

Required successor: invoke the exact-hash void/locked-contract rail now—route this gap to master for m-7/m-2 disposition, mark Task 10 held until a byte is ruled, and bind the resulting token plus RED assertion before dispatch reaches that task. Do not improvise or reuse an unrelated class locally.

### F5 — The PLAN header/process contract is incomplete for this risk class

`HUMAN_GATE_REQUIRED: yes` contradicts its own explanation that no fresh human decision is needed for this transition (`PLAN...md:11`); the terminal merge gate is re-evaluated later, not latched. The PLAN also lacks the required completed escalation scan and a top-level boundary contract despite crossing canonical store, registry, engine, binding, channel, tables, and client consumers. Finally, Task 0 directly checks out a branch in the shared workspace (`PL-...md:69-77`) while the loaded Pair Implementer contract requires `superpowers:using-git-worktrees` before implementation.

Required successor: set the current human-gate predicate to `no` with the terminal gate preserved in prose; carry a trigger-present scan without claiming a downgrade; add the canonical Writes/Reads/Target entity/Downstream consumer/Contract/E2 Proof/No-consumer action block; and make Task 0 invoke the mandatory worktree/environment procedure, reconciling its result with the charter's exact `s12-h16-fix` branch before any checkout.

## Six MUST-CHECK dispositions

1. `record_kind`: **contest and escalate** — four tokens are not the commission's exact two-token wording.
2. row type: **contest** — owner-confirmed `string` is literal and already supported; `text` is not an equivalent the pair may substitute.
3. marker: **contest as incomplete** — `s12-fieldspec-v9` is a reasonable label, but the transition/downgrade consumers, exact provenance shape, and tests must be planned.
4. genesis predecessor: **confirm** — the live genesis relay ID is byte-exact `genesis` (`internal/store/genesis.go:125-140`).
5. ceremony verb: **confirm with the stated freedom** — exact CLI spelling may be selected locally while lock-first, no-socket, and surface absence remain tested.
6. exclusions: **confirm** — `-bless` and `operatorSubmit` remain census/report-only; any discovered writer defect routes to master and receives no local edit.

## Other review dimensions

Task ordering is coherent after F3's compatibility work is inserted with Task 4. Battery coverage is otherwise total at plan grain; F3 needs marker/transition/genesis batteries and F4 needs the ruled exact-class negative. The boundary contract is absent rather than contradicted. Out-of-scope fences, frozen exit fixtures, INV-CATALOG bytes, H-12, operator-only merge, and no live redeploy all remain correctly preserved.

BOUNDARY_CONTRACT: not applicable to this review-only artifact — it writes no runtime object; the target PLAN must add its cross-component boundary contract
ACTIONS_GIT_REF: docs-workspace disk action — this PLAN-REVIEW relay plus one append-only frank/.relays/s12/INDEX.md row; no source, test, plan, roadmap, branch, commit, PR, merge, or deploy action
FINAL_GIT_STATUS_SHORT:
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s12/
?? frank/docs/sprints/active/

## Verification

- Rev20 SHA-256 rechecked earlier in this seat: `e09fab09cdcc2a6025fd42f06b02e573a3e091542f9eb44cebe42a7fa5ff7de9`.
- Plan artifact SHA-256 rechecked: `e30a8af89555f28e7c5f3bd53afae652cb1ea2b0954a35ccdb2c8eff4b477d51`.
- Incoming relay exact-file lint: OK with `--no-freshness`; INDEX lint: OK.
- Incoming active-root lint: FAIL with the two F1 errors; delegated dispatch remains blocked.

Next requested action: s12.planner files the bounded r2 successor under the protocol's incremented authority-chain identity, routes F2/F4 UP as named, and returns it for re-review. No code byte moves; no dispatch token is authorized by this verdict.
