## DESIGN-REVIEW - s4-wire.implementer narrow re-review of s4-slice-4-design r3

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s4-wire-design-r3-review-implementer
PARENT_DISPATCH_ID: s4-wire-design-r3-review
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-design/DESIGN-planner-20260705-041500.md
SUBJECT: DESIGN-REVIEW verdict - approve r3; prior F1/F2 closed; boundary and OUT fence preserved

## Verdict

DESIGN_REVIEW_VERDICT: approve

This approval is scoped to `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md` at `main@2ef9437` as reviewed through parent relay `.relays/s4/s4-wire-design/DESIGN-planner-20260705-041500.md`.

This review does not authorize PLAN, IMPL, code, spec edits, or a design-doc edit. A later PLAN may consume this approval only if it locks the same `DESIGN_DOC_ID: s4-slice-4-design` lineage, preserves `GRILL_LOCK_ID: s4-grill-s4-wire`, and references an explicit DESIGN_LOCK that includes this approving review.

## Review Scope

- Parent review request: `.relays/s4/s4-wire-design/DESIGN-planner-20260705-041500.md`.
- Candidate design doc: `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md`.
- Prior must-revise review: `.relays/s4/s4-wire-design/DESIGN-REVIEW-implementer-20260705-035327.md`.
- Reviewed git state: `HEAD/main` = `2ef9437`.
- Source anchors spot-checked for the carrier and boundary claims: `internal/fieldspec/registry.json`, `internal/fieldspec/canonical.go`, `internal/record/record.go`, and the S3 locked carrier text.

## Findings

No blocking findings.

## Closure Checks

### F1 - Guide-answer fold completeness closed

The r3 header records that the guide answers were folded and all six guide confirmations were resolved in place (`s4-slice-4-design.md:6`). The active design body now carries the material answers instead of relying on provisional markers:

- Q1 no cross-seat metadata and poll-first frame shape: section 5, especially `s4-slice-4-design.md:96-102`.
- Q2/Q3 config-change authorship, record pivot, derived bytes, chain walk, and crash class: `s4-slice-4-design.md:108-115`.
- Q4 typed frame ceiling refusal, both inbound and outbound, plus the ceiling-value disclosure carve-out: `s4-slice-4-design.md:121-127`.
- Q5 wedged-host remedy and kill-then-reconnect fixture: `s4-slice-4-design.md:91` and fixture plan `:151`.
- Q6 composite MCP confirmation, no shim-side schema reshaping, credential custody, and OUT-fence preservation: `s4-slice-4-design.md:70-82`, `:102`, and `:156`.

The remaining `[PG-Qn]` strings are confined to header/history/GRILL_LOCK/revision-log record material rather than unresolved active design sections. The new seat-occupancy paragraph is claim-honest, not scope-expanding: it says seats are durable identities occupied by sessions at launch, keeps assignment outside s4 automation, and explicitly fences conductor-spawned sessions out as m-5/m-4 plus Step-3 work (`s4-slice-4-design.md:133`).

### F2 - Structured carrier rule closed

The Form-to-JSON-Schema rule is now closed at design grain. Section 3.2 says every current `FieldSpec` type maps to a string-typed JSON-Schema property, and the shim passes the exact host-written string into `Headers map[string]string` without parsing, canonicalizing, re-encoding, or restructuring (`s4-slice-4-design.md:53-68`).

The structured live-set path is explicit: `row_array`, `object`, and `address_list` appear to the host as strings whose descriptions state the canonical JSON string shape; `ParseTyped` remains the only parser/validator; non-canonical encodings surface as normal conductor typed violations, never shim-side errors (`s4-slice-4-design.md:63`). Native JSON exposure and shim canonicalization are rejected alternatives (`:65`). The fixture plan includes the structured-carrier proof as S4-SCH2 (`:151`).

Supporting source checks match the design:

- Current registry type census: `2 address_list`, `4 bool`, `12 enum`, `7 id_ref`, `3 row_array`, `12 text`.
- `internal/record/record.go:27-31` keeps `Headers map[string]string`.
- `internal/fieldspec/canonical.go:19-75` keeps parsing and canonical-encoding enforcement in `ParseTyped`.
- The S3 locked carrier continues to define structured values as canonical JSON strings inside headers.

### Boundary contract and OUT fence preserved

The r3 design still makes s4 a tool-mediated confusion-resistance surface, not a new authority layer, and keeps D5 credential custody residual plus E3 transport/provenance boundaries visible (`s4-slice-4-design.md:15-17`). Boundary contracts remain enumerated without adding a new routing or authorization surface: host shim, conductor private dialect, config-change store path, lifecycle/channel (`s4-slice-4-design.md:137-147`).

The OUT fence is preserved. Consumer schema content, observe/evidence Step-2, routing Step-3, TUI/email UX, federation, external send, authority replacement, in-band credential rotation/supersede, and socket-dialect rewrite remain outside this design (`s4-slice-4-design.md:156`). I found no new verb, new authority, or design-doc side effect that would require a human-decision review.

## Non-Blocking Carry-Forward

- The preserved GRILL_LOCK block still contains historical `[PG-Q1]` through `[PG-Q6]` text (`s4-slice-4-design.md:224`). That is acceptable because r3 folds the guide answers into the active body and records the resolution in the header/revision log.
- PLAN review should verify that the plan locks this r3 design lineage and includes fixture coverage for S4-SCH2, S4-SC3, S4-NG4, S4-FR2, the I-PH frame-ceiling carve-out, and the config-change chain legs.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-design/DESIGN-planner-20260705-041500.md` -> OK.
- `git rev-parse --short HEAD` -> `2ef9437`.
- `git rev-parse --short main` -> `2ef9437`.
- `git show --stat --oneline 2ef9437` -> one docs-only design commit for r3.
- `jq -r '.fields[].type' internal/fieldspec/registry.json | sort | uniq -c` -> `2 address_list`, `4 bool`, `12 enum`, `7 id_ref`, `3 row_array`, `12 text`.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-design/DESIGN-REVIEW-implementer-20260705-040925.md` -> OK.

ACTIONS_GIT_REF: no source/test/design-doc/code edits; wrote relay `.relays/s4/s4-wire-design/DESIGN-REVIEW-implementer-20260705-040925.md` plus `.relays/s4/INDEX.md` row; tracked HEAD `main@2ef9437`.
FINAL_GIT_STATUS_SHORT: none - clean tree.
