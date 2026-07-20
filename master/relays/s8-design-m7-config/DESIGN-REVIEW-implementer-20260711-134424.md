## DESIGN-REVIEW - s8 config host r9 must revise the compound adoption-record contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r9
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the offline bootstrap direction is technical; its record grammar requires owner confirmation, not a new operator fork
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-134100.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r9 closes the bootstrap ordering deadlock, but its two-member config_change is a new canonical record shape whose bytes, owner contract, intents, and recovery order are undefined

DESIGN_REVIEW_VERDICT: must-revise

The offline bless mode is the correct bootstrap locus. It closes r8/F4's old-reader/full-reader deadlock, defines the two initial transitions, preserves one canonical pivot, and supplies the required legacy/adopted crash fixture. One contract gap remains in the act that carries both changed members.

## Finding

### F5 - A multi-member `config_change` is not the inherited singular mechanism

Section 5.1 calls the adoption act the ordinary §7 compound `config_change`, carrying both engine and catalog bytes, and says no new machinery is introduced. The shipped mechanism is materially narrower:

- The canonical record has one `member` header and its body is the raw bytes of that one member (`frank/internal/store/config_change.go:16-34`).
- `member` is an m-2-owned enum whose operator scope is exactly `{fieldspec, engine}` (`internal/fieldspec/registry.json:85,136`).
- `ConfigChangeIntentsStrict` resolves one `configTarget` and emits one `IntentConfig`; recovery reuses that same singular interpreter (`internal/store/projections.go:205-208`).

There is therefore no canonical encoding from which recovery can derive two member files. PLAN would have to invent whether `member` becomes a set/sentinel, whether the body is a JSON wrapper or raw bytes, how member bytes are represented byte-exactly, which fields m-2 owns, and how old singular records remain replayable. The one-pivot property is inherited; the multi-member record and projection interpreter are not.

Required fold: specify the adoption record's closed canonical byte shape at design grain. Name the header discriminator, body schema and byte-preservation rule for both members, deterministic member ordering, recomputed digest field, and singular-record backward compatibility. Route any `member`/body/form grammar delta to m-2 for explicit confirmation; m-7 may own the trusted interpreter but cannot silently widen m-2's enum or field shape. Define how this record deterministically yields exactly two `IntentConfig` projections while ordinary `config_change` records continue to yield one.

Also bind recovery ordering: after a post-pivot interruption, `store.Open`/derived-work recovery must interpret and complete both member projections before the full s8 phase-0 `config.Load`; a bless re-run must detect the committed adoption record and recover/validate the three-member state rather than re-enter legacy validation. Extend FX-CFG-12 with a crash between the first and second member projection and assert restart converges before capability checks run.

Sweep the relay/design claims from "nothing new" and "ordinary config_change" to the honest boundary: one-pivot/derived-artifact mechanics are reused, while the closed multi-member adoption-record variant and interpreter are an explicit reviewed extension.

## Confirmed

- The offline, root-locked, no-serve bless mode is the correct solution to r8/F4.
- Legacy two-member validation, absent/0-to-initial transitions, distinct bless provenance, and post-bless full-reader sequencing are sound.
- One canonical record remains the right atomic truth; member files remain recoverable derived artifacts.
- All r7/r8 transition, capability, descriptor, census, step-4.5, and operator F5 folds remain accepted.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r9 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner coordinates the m-2 record-grammar confirmation, folds F5 into config-host r10, and returns a DESIGN relay for re-review; master holds reconcile-A completion meanwhile.
