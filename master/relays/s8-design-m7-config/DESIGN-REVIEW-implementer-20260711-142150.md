## DESIGN-REVIEW - s8 config host r11 must consume the completed m-2 grammar ruling

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r11
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - no operator fork; the m-2 owner confirmation already exists and master has consumed it
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-141456.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r11 closes F6 technically, but it treats m-2 confirmation as pending and omits the MAJOR-fail-CLOSED classification master made binding from the already-consumed owner record

DESIGN_REVIEW_VERDICT: must-revise

r11 resolves r10/F6 completely: `new_digest` remains the one canonical header path, acceptance recomputes it, the body is exactly `{catalog, engine}` in order, malformed-set cases reject pre-pivot, and FX-CFG-13 proves chain advance. The remaining issue is not mechanism design; it is failure to consume the live owner/reconcile trail that landed before this relay.

## Finding

### F7 - The owner gate is already satisfied, but its binding version-class ruling is absent from r11

m-2's confirmation `s8-adoption-grammar-m2/SITREP-planner-20260711-135623.md` already made the same correction as F6: `new_digest` stays a header and the body carries members only. It confirmed `member: adoption`, reserved-token semantics, malformed-body grammar, and the two registry byte sites. It also ruled the enum widening **MAJOR on the record-schema axis, fail-CLOSED, forward-only, with no migrator** because old exhaustive readers reject the new discriminator at submit instead of ignoring or misapplying it.

Master then consumed that owner record in `SITREP-orchestrator-planner-20260711-140516.md` and made five items binding for the m-7 sweep. r11 carries the header/body correction, reserved-token effect through its exact name set, an explicit two-member count policy, and a consistent outer wrapper. It does not carry the MAJOR/fail-CLOSED/no-migrator classification or explicitly name both m-2 registry byte sites. Its relay instead says m-2 confirmation remains pending and sends a redundant superseding request.

Required fold: consume the existing m-2 confirmation and master's ruling directly. State in §5.1/GRILL_LOCK that `adoption` is a record-schema **MAJOR-but-safe** discriminator extension: pre-adoption readers typed-reject it at submit, committed history is not reclassified, the transition is forward-only, and no migrator is required. Name both scheduled m-2-owned byte sites: `config_member` enum set and `member.seat_scope.operator`.

Mark the owner gate SATISFIED by `…-135623` plus master's `…-140516` consumption. The later `…-141455` request is unnecessary for count policy: m-2 explicitly left member count to m-7, and r11's exact-two decision is within that authority. Preserve it as append-only history but do not leave lock effectiveness waiting on a duplicate confirmation.

Add or identify the fixture leg proving an old exhaustive reader rejects `member: adoption` at submit and cannot silently accept it; retain singular committed-record replay as the no-migrator/backward-history leg.

## Confirmed

- r10/F6 is fully closed by the header-only digest, exact two-member body, pre-pivot validation, and chain-walk fixture.
- The adoption interpreter, crash recovery, offline bless path, version-transition gate, and capability table remain technically approved.
- m-2 has confirmed the discriminator/body grammar; no spelling or product-semantics decision remains open.
- All descriptor, census, step-4.5, activation, and operator F5 folds remain accepted.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r11 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F7 as a trail-consistency r12, consumes the satisfied m-2/master ruling, and returns a DESIGN relay for final re-review; master holds reconcile-A completion meanwhile.
