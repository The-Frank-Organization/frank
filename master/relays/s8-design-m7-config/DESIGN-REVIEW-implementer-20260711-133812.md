## DESIGN-REVIEW - s8 config host r8 must revise the existing-store adoption path

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r8
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the F5 operator selection stands; this is a technical bootstrap finding
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-132922.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r8 transition predicate approved, but existing two-member stores have no executable path to the required three-member s8 state

DESIGN_REVIEW_VERDICT: must-revise

The r8 restructure resolves all three r7 findings. Per-version schema descriptors detect same-path type/container changes without treating values or indices as schema; the implication-based rule is internally consistent with semantics-only bumps; and owner-supplied forward relations close rollback/skip between supported string markers. One pre-lock bootstrap defect remains.

## Finding

### F4 - "Reader first, record second" deadlocks adoption of an existing two-member store

Section 2.4 requires upgrading the reader before accepting a new version, while section 5 says an already-initialized store adopts the catalog and engine-version state only through an accepted `config_change`. Neither side of that sequence can perform the first transition:

- The shipped reader accepts only `fieldspec` or `engine` in `classifyConfigChange` (`frank/internal/engine/submit.go:340-343`), and `configTarget` likewise has no catalog target (`internal/store/genesis.go:211-219`). It cannot append/materialize the catalog transition.
- A full s8 reader adds `catalog` to `StoreRootConfigPaths`; phase 0 loads every supplied path before serving (`internal/config/config.go:43-49`, `cmd/frank/main.go:101-104`). Against the existing two-member store, the catalog path is absent and the expected digest is still two-member, so the new reader cannot start to accept the record.
- The transition validator also defines only `current -> V`; it does not define the initial absent-member -> `s8-v1` catalog transition or pre-s8 absent engine version 0 -> 1 under the same governed act.

The design's genesis path is sound for new stores, but the dogfood spine is already initialized. A raw file copy cannot fill this gap because it would bypass the sole-governed-writer/config-history claim.

Required fold: define one concrete, crash-atomic existing-store bootstrap path. It must run under the trusted root/serialized authority, validate the current two-member expected digest, validate the candidate engine/catalog bytes against their initial descriptors, define the lawful absent/0 -> initial-version transitions, append the operator-authorized accepted history, materialize both required members, and advance the composite digest without an intermediate unbootable state. This may be a narrowly scoped offline migration/blessing mode or an explicitly specified transitional reader, but it cannot rely on the old or full-s8 serve path as currently ordered.

Add an executable adoption fixture that starts from the shipped `main@691d034` two-member store, performs the governed bootstrap, restarts the full s8 reader, proves the three-member digest and version/capability checks pass, and proves interrupted bootstrap recovers to either the complete old state or complete new state. Sweep section 5's "only through config_change," section 2.4's reader-first sequencing, and section 4's blessing wording to the chosen single mechanism.

## Confirmed

- r7 F1 is closed by version-specific schema descriptors and the type/container/value fixture legs.
- r7 F2 is closed by the one-way implication rule; semantics-only bumps are consistently lawful.
- r7 F3 is closed by owner-supplied forward-transition relations and the multi-marker rollback/skip fixture.
- The r7 capability table, acceptance-gate locus, recovery rule, descriptor/census, step-4.5, and operator F5 folds remain accepted.
- No new operator decision is required unless the planner presents materially different bootstrap product semantics.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r8 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F4 into config-host r9 and returns a DESIGN relay for re-review; master holds reconcile-A completion meanwhile.
