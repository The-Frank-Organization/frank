## DESIGN-REVIEW - s4-wire.implementer review of s4-slice-4-design r2/r3-queue

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s4-wire-design-r2-review-implementer
PARENT_DISPATCH_ID: s4-wire-design-r2-review
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-design/DESIGN-planner-20260705-021500.md
SUBJECT: DESIGN-REVIEW verdict - must-revise before design-lock consumption; direction approved, r3 fold required

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

This is not a rejection of the architecture. The r2 design preserves the one-line boundary, keeps the shim as affordance rather than authority, promotes existing conductor/client/render/store substrates, carries the GRILL_LOCK, and matches the paired audits on second-connect, nudge grain, MCP-composite, custody, and the section 7 record.

It is not yet safe to issue an approving DESIGN-REVIEW that a later PLAN can mechanically consume. The live artifact itself says r2 is still not lockable until guide deltas fold, and the guide answer has now landed with material requirements. One additional Form-to-JSON-Schema carrier rule is also under-specified at the design grain.

## Review Scope

- Parent review request: `.relays/s4/s4-wire-design/DESIGN-planner-20260705-021500.md`.
- Candidate design doc: `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md`.
- Live git state checked: `HEAD/main` = `059b4c5`; r2 commit = `c28dab7`; `git diff --stat c28dab7..HEAD` = one design-doc line, the r3 queue revision-log note.
- Guide answer read because r2 marks PG-Q1..Q6 provisional: `.relays/s4/s4-guide-q1/SITREP-planner-20260705-014633.md`.
- Source anchors spot-checked: `internal/channel/server.go`, `cmd/frank/main.go`, `internal/fieldspec/render.go`, `internal/fieldspec/validate.go`, `internal/fieldspec/canonical.go`, `internal/record/record.go`, `internal/seat/binding.go`, `internal/store/genesis.go`, `internal/store/projections.go`, `test/fixtures/f11_test.go`, S3 design carrier text.

## Blocking Finding F1 - r2 is explicitly not lockable, and the now-landed guide answer contains required r3 content

Evidence:
- The design header says the `[PG-Qn]` sections are provisional and guide deltas must fold before lock (`s4-slice-4-design.md:6`).
- The revision log says r2 is "Still NOT lockable" until PG-Q1..Q6 guide deltas fold as r3 (`s4-slice-4-design.md:215`).
- Live `HEAD` only adds an r3 queue line and does not fold the answer body into the design (`git diff --stat c28dab7..HEAD` = one insertion in the design doc).
- The m-7 guide answer says de-provision on that relay and adds concrete requirements:
  - no cross-seat metadata in any seat's frames plus a negative fixture (`SITREP-planner-20260705-014633.md:22-27`);
  - config-change record-pivot-then-derived-bytes, chain walk, and crash/applicability class confirmation (`:29-35`);
  - operator-channel submit as the config-change authorship, with no extra self-approval loop (`:37-39`);
  - frame ceiling typed refusal on both inbound and outbound paths, plus a disclosure carve-out for the ceiling value (`:41-47`);
  - written wedged-host remedy and kill-then-reconnect fixture (`:49-51`);
  - composite MCP confirmed, with no shim-side schema reshaping and shim custody/D5 guardrails (`:53-55`).

Why this blocks approval:
An approving DESIGN-REVIEW is a lineage artifact. A later PLAN can consume it as a design-review approval for `DESIGN_DOC_ID: s4-slice-4-design`. Approving r2 while the document still says "not lockable" and while the guide deltas are only queued would make the mechanical gate stronger than the actual reviewed content.

Required revision:
- Fold the m-7 guide answer into the design body, not only the revision log.
- Remove or resolve the `[PG-Q1]` through `[PG-Q6]` provisional markers.
- Update section 5 and the fixture plan with the no-cross-seat-metadata negative fixture.
- Update section 6.4 and the fixture plan so frame refusal is typed and path-free in both directions, including outbound `project`/`read` growth cases and the ceiling-value disclosure carve-out.
- Update section 4/7/9 with the written wedged-host remedy plus a kill-host, reconnect, recovery-leg fixture.
- Update section 3/8/9 with the guide guardrail that the shim serves the rendered form/input schema without re-adding withheld affordances.
- Fold the r3 queue's seat-occupancy-model paragraph into an appropriate design section or explicitly drop it. If it changes scope, route the narrow re-review as the planner relay already promised.

## Blocking Finding F2 - Form-to-JSON-Schema mapping does not yet lock the structured carrier rule

Evidence:
- S4 section 3.2 says a form field becomes a JSON-Schema property and "field `type` maps to JSON type" (`s4-slice-4-design.md:53-57`).
- Current rendered forms expose FieldSpec type strings only: `Form.Fields` maps to `Field{Type, Options, Default, DigestExempt}` (`internal/fieldspec/fieldspec.go:11-20`).
- The live registry includes protocol-specific types such as `address_list`, `row_array`, `object`, and `id_ref` (`internal/fieldspec/registry.json:82-109`).
- The record carrier remains `Headers map[string]string` (`internal/record/record.go:27-31`).
- The S3 locked carrier says structured types `row_array`, `object`, and `address_list` are canonical JSON strings inside `Headers map[string]string`, with validation parsing by row type (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:52-58`, `:75-77`).
- The implementation confirms that parse contract: `ParseTyped` parses `row_array`, `object`, and `address_list` from raw strings and rejects non-canonical encodings (`internal/fieldspec/canonical.go:19-75`).

Why this blocks approval:
An implementer cannot safely infer whether the MCP `inputSchema` should expose structured JSON values and have the shim canonicalize them into string headers before `submit`, or expose canonical JSON strings directly to match the conductor payload. Either choice affects user ergonomics, digest honesty, stale-form bounces, path-clean errors, and the "no shim-side reshaping" guardrail in the m-7 answer. This is design-level behavior, not plan trivia.

Required revision:
- Add a short type-mapping table for every current FieldSpec type or a closed rule that covers them.
- State explicitly whether structured fields (`row_array`, `object`, `address_list`) appear to the host model as JSON structures or as canonical JSON strings.
- If structured values are exposed, specify the one shim-side canonicalization step into `Headers map[string]string`, including where canonical-encoding violations surface and how `form_digest` remains proof of the exact host-visible schema.
- If canonical strings are exposed, say so plainly and fixture that the schema and help text keep the string carrier understandable without claiming native JSON structure.
- Preserve the m-7 guide guardrail: the shim must not re-add an affordance the renderer withheld.

## Non-Blocking Approvals

- DESIGN_DOC_ID lineage is clean: parent request carries `DESIGN_DOC_ID: s4-slice-4-design`; this review carries the same ID and parents to `s4-wire-design-r2-review`.
- GRILL_REQUIRED was correctly triggered and satisfied before this review. The design contains `GRILL_LOCK_ID: s4-grill-s4-wire` and resolved G-1 through G-6 rows (`s4-slice-4-design.md:145-205`).
- Boundary contract is directionally buildable: host shim, conductor private dialect, config-change store path, and lifecycle/channel surfaces are all enumerated (`s4-slice-4-design.md:120-130`).
- OUT fence is intact: consumer schema content, observe/evidence Step-2, routing Step-3, TUI/email UX, federation, external send, authority replacement, in-band rotation/supersede, and socket-dialect rewrite remain out (`s4-slice-4-design.md:137-139`).
- Second-connect design preserves the dispatch source of record and does not invent live supersede, rotation, re-mint, or binding-table shape changes (`s4-slice-4-design.md:73-81`; `internal/seat/binding.go:68-100` confirms current single-credential lookup shape).
- Current source evidence matches the design's baseline: current `Server.Push`/`QueuePush`/`flushPending` are broadcast/global (`internal/channel/server.go:116-153`, `:260-270`); active duplicate auth is not constrained (`:203-219`); scanner loops still use default `bufio.Scanner` (`:188`, `:448`); current `cmd/frank` submit wake is global and startup recovery sends all pending seats (`cmd/frank/main.go:184-212`).
- Section 7 direction matches the guide: `ValidateGenesis` is currently a static compare that must learn the chain (`internal/store/genesis.go:104-118`), and `f11Classes()` still lacks `config-change` (`test/fixtures/f11_test.go:303-317`), so the design targets the right missing surface.

## Required Next Review Shape

After the planner folds r3, return a narrow DESIGN-REVIEW request to this Implementer. The narrow re-review can focus on:
- guide-answer fold completeness;
- the structured Form-to-JSON-Schema carrier rule;
- preservation of the current boundary contract and OUT fence.

Do not consume this `must-revise` review as a PLAN design lock.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-design/DESIGN-planner-20260705-021500.md` -> OK.
- `git rev-parse --short HEAD` -> `059b4c5`.
- `git rev-parse --short main` -> `059b4c5`.
- `git show --stat --oneline c28dab7` -> one docs-only design commit.
- `git diff --stat c28dab7..HEAD` -> one design-doc insertion, the r3 queue line.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-design/DESIGN-REVIEW-implementer-20260705-035327.md` -> pending until after write.

ACTIONS_GIT_REF: no source/test/design-doc/code edits; wrote gitignored relay `.relays/s4/s4-wire-design/DESIGN-REVIEW-implementer-20260705-035327.md` plus `.relays/s4/INDEX.md` row; tracked HEAD `main@059b4c5`.
FINAL_GIT_STATUS_SHORT: none - clean tree
