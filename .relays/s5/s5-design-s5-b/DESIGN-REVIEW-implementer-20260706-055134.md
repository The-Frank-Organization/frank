## DESIGN-REVIEW - s5-b.implementer review of s5-b-mechanisms-design

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-design-s5-b
PARENT_DISPATCH_ID: s5-design-s5-b
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: s5-b-mechanisms-design
DESIGN_REVIEW_VERDICT: must-revise
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-design-s5-b/DESIGN-planner-20260706-053613.md
SUBJECT: DESIGN-REVIEW must-revise - egress Drain lacks the provenance path needed to prove the ODB carve-out at the real call site

## Routing and authority

- This review is directly addressed work: the planner relay `DESIGN-planner-20260706-053613.md` is `TO: s5-b.implementer`.
- Authority is DESIGN-REVIEW only. I made no source, test, registry, sprint-doc, branch, commit, PR, merge, or implementation edits.
- Reviewed artifacts: the planner relay, the design doc `docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md`, the s5 escalation update, the master reconciliation, and the current submit/fieldspec/outbox/projection code.

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

The design is close and the ③ raise, DEF-2 guard, replay, §7, and I-PH shapes are implementable against the current code. I do not approve yet because the ⑤ egress design requires `egress.Drain` to prove the ODB carve-out at the real call site, but the specified `Drain(st, rules)` path has no defined way to obtain the trusted provenance and render fields that `Scan` needs.

## Blocking finding

### 1. `Drain(st, rules)` cannot currently supply `Dest`, `Field`, or `Origin.ConductorODB`

The design defines `Scan` over an `Item` that includes `Meta gate.OutboxItem`, `Source record.Record`, `Dest`, `Field`, and `Origin`, with `Origin.ConductorODB` set only by conductor-side construction (`docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md:105-113`). It also names `Drain(st *store.Store, rules Rules)` as "THE real call site", walking `outbox/<item>.json`, resolving source records, and scanning each item before render-for-egress (`.../s5-b-mechanisms-design.md:111-115`).

The live outbox item file does not carry those scan inputs. `gate.OutboxItem` has only `item_id`, `source_kind`, `source_record_ref`, `seat`, `gate_category`, `created_ts`, and `schema_version` (`internal/gate/derived.go:8-16`). The real obligation path emits the same shape before writing `outbox/<item>.json` (`internal/obligation/obligation.go:149-181`), and store projection just writes that payload (`internal/store/projections.go:76-81`). There is no `Dest`, no rendered `Field`, and no trusted conductor-origin bit for `Drain` to reconstruct from the current outbox file alone.

This matters because §3.2 requires the fixture to use outbox items produced via the real obligation path and then invoke `egress.Drain` for the pass and block legs (`.../s5-b-mechanisms-design.md:123-132`), while acceptance criterion 3 requires those legs to be "driven through `egress.Drain` at the real call site" (`.../s5-b-mechanisms-design.md:210-213`). As written, the pass leg `conductor-generated ODB model_name -> operator passes` can only be proven by directly constructing an `egress.Item` for `Scan`, or by a test-only side channel. That would not prove the real drain chokepoint or the non-lane-reachable origin claim.

Required revision: specify the provenance/render-context contract at the real call site. Acceptable shapes include changing `Drain` to receive a conductor/away-bridge resolver that derives `Dest`, `Field`, and trusted `Origin` from a source record, or moving the "real call site" boundary to a render adapter that already owns those fields and calls `Scan(Item, rules)`. The fixture must prove the ODB pass leg and the lane-supplied fake-exempt block leg through that same contract, not by direct `Scan` construction alone. Also state explicitly that lane content cannot set or persist `Origin.ConductorODB`.

## Compatible checks

- ③ raise: compatible. The post-loop raise, detector injection, S1/S2/S3 composition, token rewrite, `gate_category_raised: "yes"` byte, `gate_category_pick`, and deletion of enum-index floors for `gate_category` match the escalation ruling and the current submit/validate locus (`.../s5-b-mechanisms-design.md:27-75`).
- DEF-2: compatible. The pre-loop submit-path guard with `Class: "system-owned"` is the right fix for the current `ignorePayloadField` skip behavior, and the conductor/authorship asymmetry is stated clearly (`.../s5-b-mechanisms-design.md:79-89`).
- Replay/versioning: compatible. The mandatory constructed-store leg, optional `FRANK_S5_REPLAY_STORE` archive leg, `Reader.Read` refusal legs, and `migrate.Current == 1` discipline are scoped correctly (`.../s5-b-mechanisms-design.md:136-155`).
- §7 config-change legs: compatible. The five legs consume s5-a's registry after the orchestrator integration gate and extend the s4 harness rather than rebuilding it (`.../s5-b-mechanisms-design.md:159-169`).
- I-PH: compatible, with the ⑤ caveat above. The new strings are enumerated and the formatter valve remains untouched (`.../s5-b-mechanisms-design.md:173-183`).
- Boundary contract: compatible. The write/read surfaces are within s5-b ownership and keep s5-a registry content as an input contract (`.../s5-b-mechanisms-design.md:187-216`).

## Design-grill notes

The eight planner-highlighted decisions are code-answerable from the current tree. I found no blocker in the ③ precedence, `RenderEnv.KnownA` contract, `system-owned` reject class, post-loop raise placement, computed-result guard, replay subpackage, or `Origin.ConductorODB` key as a type. The single unresolved decision is the egress locus/provenance pairing: `internal/egress` is a fine package boundary, but the design must name the conductor-origin resolver or adapter boundary that feeds `Drain`/`Scan`.

## Actions and verification

ACTIONS_GIT_REF:
- Source/test/registry edits: none.
- Relay artifact written: `.relays/s5/s5-design-s5-b/DESIGN-REVIEW-implementer-20260706-055134.md`.
- Index row appended: `.relays/s5/INDEX.md`.
- Pre-write `git status --short`: `?? docs/sprints/2026-07-06-s5-slice-5/`.

FINAL_GIT_STATUS_SHORT: `?? docs/sprints/2026-07-06-s5-slice-5/` - unchanged from the pre-write status; `.relays/` is ignored.

VERIFICATION:
- Exact-file relay lint: `OK .relays/s5/s5-design-s5-b/DESIGN-REVIEW-implementer-20260706-055134.md`.
- Index tail readback: row `20260706-055134 | DESIGN-REVIEW | Implementer | s5-design-s5-b | ... | must-revise | s5-design-s5-b/DESIGN-REVIEW-implementer-20260706-055134.md` present.
- Final `git status --short`: `?? docs/sprints/2026-07-06-s5-slice-5/`.
