## AUDIT - s3-form.implementer independent audit: form schema gates mapped to current S1/S2 surface

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s3-form-audit-implementer
PARENT_DISPATCH_ID: s3-form-audit
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s3-form.implementer
TO: s3.orchestrator-planner
CC: operator
SUBJECT: AUDIT - S3 form schema surface: FieldSpec registry, linter dissolution, replay, migrators, R2, drift

Current artifact: `.relays/s3/s3-form-audit/AUDIT-orchestrator-planner-20260704-152800.md`
Scope: read-only audit only. No source, test, sprint-doc, locked-spec, PLAN, IMPL, merge, branch, commit, PR, or live implementation work performed.

## Summary verdict

PRIMARY_BUCKET: still-open
still-open: The S3 exit gate is not implemented yet. Current code has useful seeds: a small form registry, stale-render digest checks, schema_version in the record envelope, parent-existence lineage checks, owed-item projection plumbing, trusted config pinning, and S1 replay census scaffolding. Those seeds do not yet satisfy the locked S3 form-schema gate.
already-closed: S1/S2 closed enough infrastructure to promote instead of restart: form render/validate hooks, submit-time validation ordering, typed record envelopes, accepted/rejected/held delivery states, materialized owed projections, config digest pinning, and crash/replay test harnesses.
product-overlapped: The largest overlap is with m-7 trusted config and m-1 store/lineage fidelity. S3 should implement the m-2 registry and validation engine, but must consult those boundaries before changing config member shape, read-side migrators, parent/candidate derivation, or store query semantics.
recommended-next: Proceed to DESIGN/PLAN only after reconciling this audit with the planner audit. Treat the current registry as a seed, not the full FieldSpec model, and require m-1 fidelity plus m-7 config consultation before any implementation dispatch that changes store/query/config surfaces.

## Spec-to-exit-gate map

1. Full FieldSpec registry live end to end:
- S3 exit gate: the next slice must close "Full FieldSpec registry live end-to-end" and "fill-time authority from registry" (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:57-60`).
- Locked text: m-2 defines every submitted field as a typed `FieldSpec` with owner, source, type, required/visible predicates, enum source, seat scope, gate referenceability, fill constraints, consumers, and lineage role (`the m-2 forms/determinism design-of-record (2026-06-28) :45-69`). It requires render, fill, send, and accepted/rejected commit semantics from that registry (`the m-2 forms/determinism design-of-record (2026-06-28) :71-77`), and ACs require registry, fill-time authority, X-field overflow, consumer fields, ROLE system ownership, and pre-append form plus lineage validation (`the m-2 forms/determinism design-of-record (2026-06-28) :224-234`).
- Current code: `internal/fieldspec/fieldspec.go:19-26` has only six small enum-like specs; `internal/fieldspec/registry.json:2-47` has raw enum values, not full FieldSpec rows. `internal/fieldspec/fieldspec.go:54-105` renders and validates a fixed subset directly, so it is a useful seed but not an end-to-end registry engine.

2. Linter dissolution and disposition table:
- S3 exit gate: linter dissolution must be backed by "disposition table + FULL replay" (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:60-61`).
- Locked text: m-2 says no check is dropped without proof; prose-only checks dissolve only because strict form-only submission removes the surface, typed-form checks move to form validation, and cross-relay checks move to a lineage engine (`the m-2 forms/determinism design-of-record (2026-06-28) :131-180`).
- Current code: `internal/fieldspec/fieldspec.go:81-105` implements a small typed validator, and `internal/lineage/lineage.go:43-64` validates only parent existence/accepted state. There is no 62-row disposition table, no retained lineage engine for the §10c checks, and no generated no-silent-drop assertion.

3. Full replay of the upstream corpus:
- S3 exit gate: the upstream linter corpus path is explicit in the roadmap (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:25-26`) and the exit gate requires full replay (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:60-61`).
- Locked text: m-2 §10 is the replay classification authority (`the m-2 forms/determinism design-of-record (2026-06-28) :131-180`).
- Current code: `test/replay/classmap.go:22-44` and `test/replay/classmap.go:116-127` provide a heuristic classifier over the corpus and known fixture classes. `test/replay/replay_test.go:8-34` pins that every representative has a disposition and that the generated report matches. This is a census harness, not yet an execution replay against the final S3 validator.

4. schema_version and migrators:
- S3 exit gate: schema_version plus migrators are explicit S3 deliverables (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:61-62`).
- Locked text: m-2 requires a versioned form envelope and a projection/read-time migrator chain with zero migrators at birth, while m-1 persists and serves versioned records (`the m-2 forms/determinism design-of-record (2026-06-28) :122-129`; `the m-1 trust/identity design-of-record (2026-06-28) :129-135`).
- Current code: `internal/record/record.go:16-24` includes `schema_version`, `internal/engine/submit.go:29-31` stamps schema version 1, and `internal/store/store.go:77-79` defaults missing versions on commit. A code search found no migrator registry or read-side migration path.

5. R2 negatives and `GRILL_REQUIRED` row:
- S3 exit gate: R2 negative fixtures and `GRILL_REQUIRED` FieldSpec row are explicit (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:62`).
- Locked text: ARCH owes "R2 gate_referenceable per-column negative fixtures" and an "m-2 GRILL_REQUIRED FieldSpec row" (`../master/ARCHITECTURE.md:477-489`). m-2 defines `gate_referenceable` and forbids routing predicates over non-referenceable columns like observed model identity (`the m-2 forms/determinism design-of-record (2026-06-28) :57`, `the m-2 forms/determinism design-of-record (2026-06-28) :87-99`, `the m-2 forms/determinism design-of-record (2026-06-28) :383-389`).
- Current code: `internal/fieldspec/fieldspec.go:108-123` has gate_category classification and a known-A raising boundary, but no gate-referenceable registry property, no R2 predicate grammar, no per-column negative fixtures, and no `GRILL_REQUIRED` row.

6. Re-render and drift with no regression:
- S3 exit gate: no S1/S2 regression and no materialize-first regression are required (`docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md:64-65`).
- Locked text: m-2 requires a fill-time render, submit-time validation, accepted/rejected commit semantics, X-field overflow, and no persisted "submitted" state (`the m-2 forms/determinism design-of-record (2026-06-28) :71-77`, `the m-2 forms/determinism design-of-record (2026-06-28) :103-107`, `the m-2 forms/determinism design-of-record (2026-06-28) :235-239`).
- Current code: `internal/fieldspec/fieldspec.go:86-89` rejects stale form digests with class `re-render`, and `internal/fieldspec/fieldspec_test.go:36-58` proves that seed. It is not yet tied to versioned registry digest, config digest, migration-aware drift, or stale-render bounce coverage after registry changes.

Spec gaps: none found in the locked S3 exit-gate text. The open issues are implementation, fidelity, and consult surfaces, not missing roadmap authority.

## 62-check and corpus census

Linter source census: the upstream linter still has 62 assertion sites relevant to the S3 dissolution table. The line anchors I counted are `relay-lint.py:116`, `128`, `138`, `142`, `157`, `173`, `211`, `242`, `250`, `266`, `291`, `307`, `311`, `345`, `355`, `374`, `394`, `398`, `407`, `414`, `447`, `455`, `464`, `470`, `482`, `494`, `498`, `522`, `531`, `535`, `541`, `547`, `555`, `578`, `591`, `609`, `622`, `626`, `630`, `639`, `648`, `662`, `685`, `696`, `712`, `721`, `740`, `764`, `785`, `801`, `1006`, `1020`, `1034`, `1047`, `1051`, `1075`, `1091`, `1104`, `1132`, `1415`, and `1424` in `<relay-lint tools>/relay-lint.py`.

Fixture oracle census: `check-relay-lint-fixtures.py:17-164` defines 146 oracle cases. This audit run passed the checker with zero failures. Case split: 68 file cases and 78 root cases; 50 expected OK and 96 expected failure.

Raw corpus census: the fixture tree contains 243 Markdown fixture files excluding the root README, across 14 top-level categories: addressing 11, claude 20, content 17, design-review 73, fold 13, identity 1, lineage 17, lint-test 7, merge 31, merge-token 15, orch-review 26, p9 2, probes 4, and rowtruth 6.

Against m-2 §10:
- §10a prose-only dissolution covers legacy raw markdown, row grammar, bare tokens, fence boundaries, identity/proxy-from prose, and merge-token surfaces that strict form-only submission removes (`the m-2 forms/determinism design-of-record (2026-06-28) :142-149`).
- §10b typed form validation covers enum byte exactness, required fields, address grammar/cardinality, `FOLD_SCOPE`, `SCOPE_DIFF`, row-evidence fields, R2 grammar, and the gate_category A/B set (`the m-2 forms/determinism design-of-record (2026-06-28) :154-161`, `the m-2 forms/determinism design-of-record (2026-06-28) :270-283`).
- §10c cross-relay lineage covers accepted parent dispatch lineage, design-review lineage, dispatched-implementer lineage, orchestrator-review visibility, merge grant lineage, and OUT-to-IN drift (`the m-2 forms/determinism design-of-record (2026-06-28) :166-172`).

S3 implication: the replay gate must prove each of the 62 checks is either retained in typed form/lineage validation or dissolved only because the old surface cannot be submitted. The current `test/replay` classification is useful census scaffolding, but it still labels items as `uncovered-S3`, which is incompatible with closure.

## Six S3 bucket verdicts

### Full FieldSpec registry

PRIMARY_BUCKET: still-open
already-closed: A small registry, renderer, validator, seat-scoped grant hiding, X-field carriage, and stale digest seed exist (`internal/fieldspec/fieldspec.go:19-123`; `internal/fieldspec/fieldspec_test.go:11-74`; `internal/record/record.go:27-32`).
still-open: Full FieldSpec rows are absent: owner, source, type, required_when, visible_when, enum source, seat_scope, gate_referenceable, fill constraints, consumers, and lineage role. Dynamic required sets, form-owned/system-owned separation, and registry-driven consumer rules are not implemented.
product-overlapped: m-7 owns trusted config loading and restart semantics; m-1 owns envelope stamping and parent/recipient system fill. The registry implementation must not silently move those responsibilities into m-2.
recommended-next: Promote the current flat registry into a versioned FieldSpec data model loaded through trusted config, then make render and validate iterate over those rows instead of fixed fields.

### Linter dissolution plus disposition table

PRIMARY_BUCKET: still-open
already-closed: The locked m-2 §10 map exists, the current linter source can be counted, and the upstream oracle checker still passes under the archived tool.
still-open: There is no generated 62-row disposition table, no retained-check engine for the §10c lineage checks, and no proof that every legacy failure maps to "caught" or "obsolete because surface vanished." `internal/lineage/lineage.go:43-64` only checks parent existence and accepted state.
product-overlapped: The lineage half overlaps m-1 record graph and m-7 dispatch semantics, but the disposition ledger itself is m-2/S3 work.
recommended-next: Build the table from the actual upstream assertion inventory, then require each row to name one of: form validation, lineage validation, or vanished legacy surface.

### Full replay

PRIMARY_BUCKET: still-open
already-closed: `test/replay/classmap.go:22-44` walks the corpus and classifies known representatives; `test/replay/replay_test.go:8-34` pins report coverage and known dispositions. The raw corpus and oracle counts were reproduced in this audit.
still-open: The classifier is heuristic and does not execute the final S3 validator. It still permits `uncovered-S3`, which must be eliminated for the exit gate.
product-overlapped: None requiring reroute; replay should be owned by S3 but will exercise m-1 and m-7 seams through the submitted-form path.
recommended-next: Replace heuristic replay with execution against the registry validator plus lineage validator. Expected-OK fixtures should pass; expected-failure fixtures should be either caught by a named new invariant or obsolete by a named removed surface.

### schema_version and migrators

PRIMARY_BUCKET: still-open
already-closed: `schema_version` exists in the record envelope (`internal/record/record.go:16-24`), submit stamps version 1 (`internal/engine/submit.go:29-31`), store commit defaults empty versions (`internal/store/store.go:77-79`), and derived outbox bodies include schema_version (`internal/obligation/obligation.go:170-210`).
still-open: No migrator registry, read-side migration chain, unknown/future-version refusal, unversioned-record fixture, or projection-time migrated view exists.
product-overlapped: This is the strongest m-1 fidelity surface. m-2 can define the form schema and migrator contract, but m-1 owns read/project API behavior and store shape.
recommended-next: Add a read/projection migration layer above the migration-agnostic store. Do not change canonical persistence without m-1 fidelity approval.

### R2 negatives plus `GRILL_REQUIRED` row

PRIMARY_BUCKET: still-open
already-closed: `gate_category` has byte-exact A/B categories in current config and code (`internal/fieldspec/registry.json:33-39`; `internal/fieldspec/fieldspec.go:108-123`), and m-2 already specifies A/B/held delivery semantics (`the m-2 forms/determinism design-of-record (2026-06-28) :270-283`).
still-open: There is no `gate_referenceable` property in code, no predicate grammar, no negatives for observed model identity or other non-referenceable fields, and no `GRILL_REQUIRED` FieldSpec row.
product-overlapped: The §J2 A-set currently exists twice: `internal/fieldspec/registry.json:33-39` and `internal/lineage/lineage.go:66-77`. That duplication is a drift risk and should be collapsed under the trusted registry/config source.
recommended-next: Add `gate_referenceable` to the registry, reject predicates over non-referenceable fields, include per-column negative fixtures, and add the `GRILL_REQUIRED` row from ARCH.

### Re-render and drift

PRIMARY_BUCKET: still-open
already-closed: `internal/fieldspec/fieldspec.go:86-89` rejects stale `formDigest` as `re-render`, and `internal/fieldspec/fieldspec_test.go:36-58` covers stale digest rejection.
still-open: Digesting the current fixed render is insufficient for full S3. The drift proof needs registry version/digest, member/config digest relation, submit-time recheck under current trusted config, and a fixture proving a stale rendered form bounces after registry change.
product-overlapped: m-7 owns trusted config digest and no-hot-reload semantics (`the m-7 conductor-core design-of-record (2026-07-01) :107-111`), so drift semantics must align with startup-only config changes.
recommended-next: Seed the render digest from the versioned FieldSpec registry and trusted config member digest, then make stale-form bounces explicit in tests and replay output.

## S1 grant-narrowing carry disposition

Disposition: still-open and in S3 scope. The current README explicitly says the conditional pair-Planner delegated-dispatch grant is an S3 continuation (`README.md:11`). Current code renders grant options only when `canGrant` returns true for operator/orchestrator seats (`internal/fieldspec/fieldspec.go:54-78`, `internal/fieldspec/fieldspec.go:99-109`; `internal/fieldspec/fieldspec_test.go:11-34`). m-2 §10 requires the dispatched-implementer lineage walk to survive as a lineage-engine check (`the m-2 forms/determinism design-of-record (2026-06-28) :166-168`) and converts dispatch/merge authority to seat-scoped enum options (`the m-2 forms/determinism design-of-record (2026-06-28) :176-177`). This is not a spec gap; it is a carry-forward implementation requirement.

## m-1 lineage and store-touch enumeration

1. Parent and candidate derivation: m-1 owns `PARENT_DISPATCH_ID` as a `parent_picker`, recipient pickers, and routing reference derivation (`the m-1 trust/identity design-of-record (2026-06-28) :141-150`). Current lineage only resolves an already-submitted parent string (`internal/lineage/lineage.go:43-64`). Any S3 move to derive, restrict, or certify parent candidates needs m-1 fidelity.

2. Store query/read behavior: a full lineage engine will need accepted-record graph queries by dispatch, role, phase, design artifact, and routing relation. Current code mostly scans all records (`internal/store/store.go:105-131`; `internal/engine/submit.go:138-172`; `internal/obligation/obligation.go:77-143`). Adding indexed or migrated read behavior touches m-1's submit/project/read contract.

3. Delivery-state and accepted-only graph: lineage checks must distinguish accepted, rejected, and held records. Current delivery states exist in `internal/record/record.go:10-14`, but any new graph API or rejection/held visibility rule is m-1-adjacent.

4. schema_version migration: m-2 requires read/projection-time migrators, while m-1 persists and serves versioned records (`the m-2 forms/determinism design-of-record (2026-06-28) :122-129`; `the m-1 trust/identity design-of-record (2026-06-28) :129-135`). The store should remain migration-agnostic unless m-1 approves a different split.

5. System-filled fields: m-1 owns stamping for ROLE/FROM, parent, recipients, operator address, and internal provenance (`the m-1 trust/identity design-of-record (2026-06-28) :141-150`). S3 should validate these after system fill, not move trust decisions into user-submitted form fields.

## m-7 config-seam questions

1. Should the S3 FieldSpec registry remain a single trusted config member named `fieldspec`, or should it become a per-domain section with m-4-style section stamps? m-7 requires per-domain sections, one top-level digest, and no hot reload (`the m-7 conductor-core design-of-record (2026-07-01) :107-111`), while current config has only `engine` and `fieldspec` members (`internal/config/config.go:17-22`; `internal/store/genesis.go:86-90`).

2. How should registry version, record `schema_version`, and trusted config digest be related without conflating them? m-2 and m-7 describe distinct axes: record schema migration versus trusted startup config digest (`the m-2 forms/determinism design-of-record (2026-06-28) :122-129`; `the m-7 conductor-core design-of-record (2026-07-01) :107-111`).

3. Is stale render drift always a restart-boundary issue because no hot reload exists, or should S3 simulate registry change in tests by restarting with a new trusted config member? The answer affects how `formDigest` fixtures are written.

4. Where should the §J2 gate_category A-set live? Current code duplicates it in config and lineage (`internal/fieldspec/registry.json:33-39`; `internal/lineage/lineage.go:66-77`), which is exactly the kind of drift the S3 registry is meant to remove.

## Claim-boundary probes

Linter dissolution claim: "obsolete" is only valid when the old input surface cannot be submitted. Strict form-only submission removes raw markdown, tables, bare tokens, fence scans, and user-authored ROLE/FROM mismatch surfaces (`the m-2 forms/determinism design-of-record (2026-06-28) :118-120`, `the m-2 forms/determinism design-of-record (2026-06-28) :142-149`). Anything still semantically required must be caught by form or lineage validation.

Fill-time authority claim: this is a tool-mediated confusion-resistance claim, not a cryptographic impossibility claim. Forbidden options are absent from the rendered form and rechecked at submit (`the m-2 forms/determinism design-of-record (2026-06-28) :31-41`, `the m-7 conductor-core design-of-record (2026-07-01) :119`). ARCH still carries D5 shell-bypass residual risk (`../master/ARCHITECTURE.md:457-470`).

Step-gate claim: CQ-1 only defers observe-owned required fields until the observe layer is present. System-owned, form-owned, lineage-owned, and always-required fields remain required in Step 1 (`the m-2 forms/determinism design-of-record (2026-06-28) :95-97`, `the m-2 forms/determinism design-of-record (2026-06-28) :242-244`). S3 must model owner and layer predicates, not weaken all required validation.

I-PH claim: validation and bounce text must not expose canonical store paths, config member paths, or outbox details. m-7 explicitly keeps internal absence reasons off seat surfaces (`the m-7 conductor-core design-of-record (2026-07-01) :124-128`), and ARCH keeps conductor-internal provenance guarded (`../master/ARCHITECTURE.md:430-448`).

## Reuse and duplicate findings

Reuse/promote:
- `internal/fieldspec/fieldspec.go:54-105` for render/validate hook placement, seat-scoped grant hiding, and stale digest class.
- `internal/record/record.go:16-32`, `internal/engine/submit.go:29-48`, and `internal/store/store.go:60-97` for schema_version envelope plus pre-append validation ordering.
- `internal/lineage/lineage.go:21-64` for authority-bearing and parent accepted-state seed, while replacing its narrow check set with §10c lineage rules.
- `internal/config/config.go:29-83` and `internal/store/genesis.go:31-117` for trusted config loading, digesting, materialization, and genesis validation.
- `internal/store/projections.go:130-177` and `internal/obligation/owed.go:8-30` for materialize-first owed open-set projection behavior.
- `test/replay/classmap.go:22-44` and `test/replay/replay_test.go:8-34` for corpus walking and report pinning.

Do not promote as final:
- The flat `internal/fieldspec/registry.json:2-47` shape is not the full FieldSpec schema.
- The heuristic `test/replay/classmap.go:116-127` classifier is not proof of final validation behavior.
- The hardcoded gate-category A-set in `internal/lineage/lineage.go:66-77` duplicates registry config and should not become a second source of truth.
- The archived `relay-lint.py` is evidence and replay input, not an implementation source to copy.

## Spec gaps and escalations

Spec gaps: none found.

Escalations for reconcile:
- m-1 fidelity required before implementation changes store read/query semantics, read-side migrators, parent picker derivation, recipient picker derivation, or accepted-record graph APIs.
- m-7 consult required before changing trusted config member shape, registry digest semantics, hot-reload assumptions, or the gate_category source of truth.
- No design-grill required by this implementer audit beyond the already-owed `GRILL_REQUIRED` FieldSpec row; that row is itself an S3 implementation deliverable.

## Evidence and verification

E1 evidence:
- Exact dispatch read: `.relays/s3/s3-form-audit/AUDIT-orchestrator-planner-20260704-152800.md`.
- Locked roadmap/specs read and cited: S3 roadmap, m-2 design, m-1 design, m-7 design, and ARCHITECTURE.
- Current code surfaces read and cited: `internal/fieldspec`, `internal/lineage`, `internal/record`, `internal/engine`, `internal/config`, `internal/store`, `internal/obligation`, and `test/replay`.
- Archived upstream linter and fixture checker read from `<relay-lint tools>/`.

E2 evidence:
- `go test -count=1 ./...` passed.
- `go vet ./...` passed.
- `check-relay-lint-fixtures.py` passed all 146 oracle cases.

ACTIONS_GIT_REF: `1ea7f25`
FINAL_GIT_STATUS_SHORT: clean worktree; final `git status --short` after this audit relay write and index append produced no output.
