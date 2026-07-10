## DESIGN-REVIEW — s2-slice-2-design r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s2-core-design-r1-review-implementer
PARENT_DISPATCH_ID: s2-core-design-r1-review
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: DESIGN-REVIEW verdict — must-revise for one narrow phase/capability contradiction before DESIGN_LOCK consumption

Reviewed:
- Parent request: `.relays/s2/s2-core-design/DESIGN-planner-20260704-014127.md`.
- Design doc: `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md` at `main@594259a`.
- Binding dispatch inputs: r2 `.relays/s2/s2-core-design/DESIGN-orchestrator-planner-20260704-005310.md`, de-provision supplement `.relays/s2/s2-core-design/SITREP-orchestrator-planner-20260704-005315.md`, r1 hard constraints `.relays/s2/s2-core-design/DESIGN-orchestrator-planner-20260704-004400.md`.
- Locked text/code anchors spot-checked: m-7 `:52`, `:56-58`, `:89-95`, `:109-111`, `:136-137`; m-1 `:122-145`; ARCH §C4.1/C4.3; `cmd/frank/main.go`, `internal/channel/server.go`, `internal/intake/journal.go`, `internal/store/store.go`, `internal/gate/derived.go`, `internal/recover/recover.go`, `test/fixtures/f11_test.go`.

## Verdict

`DESIGN_REVIEW_VERDICT: must-revise`

This is not a rejection of the design direction. The bulk of r1 satisfies the dispatch: it promotes the S1 pieces rather than rebuilding them; closes the single-intake-writer and live-quarantine audit findings; keeps the m-1 surfaces as proposals; folds Q1/Q2 and GRILL_LOCK; and supplies fixture-keyed acceptance criteria G1-G6. One structural contradiction must be fixed before an approving design-review can be consumed by PLAN.

## Blocking finding

### F1 — `Ready` currently over-gates the read-only/diagnostics surface

The locked recovery contract allows read-only service before open: phase 0 digest/genesis failure must "serve read-only diagnostics, accept nothing" and phases 1-3 may serve reads while consuming zero authority (`the m-7 conductor-core design-of-record (2026-07-01) :89-95`; guide supplement sharpened the same disposition in `.relays/s2/s2-core-design/SITREP-orchestrator-planner-20260704-005315.md`). The r1 design adopts that requirement in D-4 and S2-V3: missing/mismatched genesis keeps the conductor up, renders `project`/`read` only, omits `submit`, emits path-free diagnostics, and mutates no store files (`docs/.../s2-slice-2-design.md:55-57`, `:103-105`).

But D-1 also says `channel.ServeAuthenticated` requires the phase-4 `Ready` value and that only phase 4 mints `Ready` (`docs/.../s2-slice-2-design.md:34-36`). As written, the authenticated channel cannot serve the read-only registry or operator diagnostics until phase 4, while D-4 requires exactly those surfaces when phase 4 is not reached. Current assembly shows why this matters: today `frankrecover.RunWithProcessor` completes before `channel.ServeAuthenticated` is wired (`cmd/frank/main.go:72-88`), and `ServeAuthenticated` is the surface that installs the per-seat toolset (`internal/channel/server.go:62-69`, `:156-185`).

Required revision: split the capability boundary so mutation/authority consumption is structurally gated by phase 4 without suppressing read-only diagnostics. Acceptable shapes include a separate read-only/diagnostics capability, a mode-typed recovered state whose toolset excludes `submit`, or `Submit` requiring `Ready` while `project`/`read` require only a read-only recovery proof. The revised text must preserve all three properties together:
- phase-0 failure and phases 1-3 can expose only path-free `read`/`project`/diagnostics where allowed;
- `submit`, the intake writer, and the commit loop remain unconstructible until phase 4;
- S2-V3 and S2-PM2 fixture wording names this capability split, so implementation cannot satisfy one half by deleting the other.

## Required clarification in the same revision

### F2 — D-7's obligation-source definition is narrower than its own instances

D-7 defines an obligation class as a "source predicate over committed records" and `open(class) = source-record with no completion-record` (`docs/.../s2-slice-2-design.md:71-79`). Two r1 mechanisms need durable non-record facts too: quarantine incident healing uses the quarantined file as durable intent (`:61-63`, table row `quarantine/ member` at `:78`), and GC completion uses a marker record plus still-present segment files (`:87-90`). If implemented literally as "committed records only," quarantine crash-window healing has no source record after evict-before-incident.

Required revision: define the generalized obligation source as a durable store fact, with committed records as the normal case and explicitly listed file-backed intents for quarantine/GC. This is a wording/shape clarification, not a request for a second mechanism.

## PLAN-time carry-forward notes

- Preserve explicit F11 applicability rows for owed-item and owed-disposition commits even if both use the generic submit-accept mutation path. Dispatch r1 named owed-item and disposition among the new mutation-class arms; r1 says they are "ordinary submit-accept instances" (`docs/.../s2-slice-2-design.md:94-97`). That can be fine, but the reviewed fixture artifact must show the mapping instead of hiding them behind a broad class name.
- Keep the m-1 fidelity packet gate hard. Section 4 is extractable and correctly marked PROPOSAL (`docs/.../s2-slice-2-design.md:125-136`); no store-shape-touching implementation dispatch should issue until the m-1 approve path has accepted or revised those shapes.

## Non-blocking approvals

- DESIGN_DOC_ID lineage is clean: parent DESIGN request carries `DESIGN_DOC_ID: s2-slice-2-design`; this review carries the same ID and parents to `s2-core-design-r1-review`.
- GRILL_LOCK is present and folded: `s2-grill-s2-core` is in §8 with guide-resolved Q1/Q2 and operator GC rows (`docs/.../s2-slice-2-design.md:155-187`).
- Canonical-sufficiency is a real S2 design requirement, not scope creep: current outbox records commit only headers while the item payload exists in the redo intent (`internal/gate/derived.go:105-114`), and current projection rebuild is redo-driven (`internal/store/projections.go:18-35`). D-6's body embedding is therefore necessary before redo-segment GC can be honest.
- The single-intake-writer finding is correctly targeted: current authenticated per-connection handlers can call `journal.Append` concurrently (`cmd/frank/main.go:87-95`, `internal/channel/server.go:72-88`), and `Append` currently computes the next id by `ReadAll()+len` without a writer lock (`internal/intake/journal.go:41-72`).
- Claim-boundary language is mostly clean: r1 pins exactly-once EFFECT / materialize-first / D5 / I-PH in the opening boundary and claim-sweep note (`docs/.../s2-slice-2-design.md:10`, `:189-191`).

## Verification

- `git show --stat --oneline --decorate HEAD` -> `594259a (HEAD -> main) s2 DESIGN: s2-slice-2-design r1 ...`; one tracked design doc added.
- `git status --short` before relay write -> no output.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md` -> `OK .relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md`.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2 .relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md` -> expected routing-table noise for `.relays/s2/INDEX.md` missing relay headers, then `OK .relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md`.
- `.relays/s2/INDEX.md` EOF row verified: `20260704-014646 ... DESIGN-REVIEW ... must-revise ... s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md`.

ACTIONS_GIT_REF: no source/test edits; wrote gitignored relay `.relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md` plus `.relays/s2/INDEX.md` row; current tracked HEAD `main@594259a`.
FINAL_GIT_STATUS_SHORT: none — clean tree
