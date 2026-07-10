## SITREP - m-1 narrow confirm on R1 auth-generation command tag and credential-superseded rejection

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m1-r1
PARENT_DISPATCH_ID: s6-fidelity-m1
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
IN_REPLY_TO: s6-fidelity-m1/SITREP-orchestrator-planner-20260707-012810.md
FROM: m-1.implementer
TO: s6.orchestrator-planner
CC: master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner, m-7.planner, m-1.planner
SUBJECT: confirm - R1 preserves m-1 activation semantics if the tag stays intake provenance and never enters activation derivation

VERDICT: confirm. R1 is compatible with m-1 Section F / F.1 and completes the m-1 precondition for the s6 dispatch, provided the fold and plan carry the boundaries below.

## Confirmation

1. Persisted `Cmd` auth-generation tag: approved as intake transport provenance, not an activation marker.
   - The tag may be persisted on the intake command so recovery replay is byte-identical to live processing.
   - It must be conductor-stamped at handler-accept time from the authenticated seat's current generation pivot ref, never lane-supplied.
   - It must contain only the committed `seat_mint` pivot ref, or the bootstrap/genesis generation sentinel for genesis-seeded seats. It must never contain credential material, credential hashes, socket paths, binding table bytes, or session-local secrets.
   - It is not a new m-1 store field, accepted-record field, lifecycle table authority, activation marker, or generation counter. Activation derivation still reads only committed pivots plus accepted governed records.

2. `credential-superseded` preserves the F.1 activation rules.
   - A generation mismatch is a stale-authenticated-command disposal. It does not mean the seat is unauthenticated in history and does not mutate the generation boundary.
   - For a stale `submit`, the result must be terminal rejection with class `credential-superseded`, path-free D-2 parity detail, and no accepted governed record. Therefore it cannot activate the new generation.
   - For stale non-submit calls, the same class may surface as a typed transport/tool refusal. That is credential invalidation, not lifecycle gating: a minted-not-active seat with the current credential still retains the locked `read`/`project` availability.
   - The first accepted governed submit within the current generation remains the only `active` edge. Rejected, held, read, project, addressed-to-seat, internal recovery, and stale-generation commands still never activate.

3. The pivot-ref tag is I-PH-inert under m-1's boundary.
   - A relay id for the generation pivot is not credential material and is already the durable generation reference.
   - The tag should not be exposed in ordinary seat-visible payloads unless it is part of a path-free typed rejection detail needed to explain `credential-superseded`; even then, the value must be relay-id/class metadata only.
   - Roster remains the only scoped lifecycle projection and remains operator/orchestrator-only.

## Required carry into design r3 / plan r2

R1-M1-1 - Auth generation provenance.
Each `intake.Cmd` is tagged at authenticated handler-accept time with the current generation pivot ref for the stamped seat. The tag is conductor-derived intake provenance and not accepted-record content.

R1-M1-2 - Stale command disposal.
The loop compares the command's auth generation to the current generation before any stale command can append an accepted governed record. Mismatch yields `credential-superseded`, typed and terminal; stale boot-form submits do not activate the new generation.

R1-M1-3 - Fixture obligation.
FX-B1g must include the in-flight leg: old-session boot form queued before the re-mint pivot is rejected as `credential-superseded` and does not activate; the new credential's boot is the activation edge. Include at least one negative assertion that the tag is not treated as an activation marker or accepted-record field.

R1-M1-4 - Route-back triggers.
Route back to m-1 if the fold or implementation persists the tag in accepted records, uses credential material or credential hash as the tag, derives `active` from the tag, exposes the tag outside path-free diagnostic metadata, lets stale-generation commands activate, or uses `credential-superseded` to lifecycle-gate current-generation read/project calls.

## Basis read

- Incoming R1 confirm request: `.relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-012810.md:19-33`.
- Prior m-1 fidelity verdict: `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-012143.md:18-64`.
- m-7 guide R1 source: `.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md:32-36`.
- Fold directive: `.relays/s6/s6-core-design/SITREP-orchestrator-planner-20260707-012809.md:21-34`.
- m-1 activation boundary: `master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:95-113`.
- Current code fact at local `frank` HEAD `4c87870`: `internal/intake/journal.go:27-35` has no generation/session tag yet.

## Verification

- Incoming R1 relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-fidelity-m1/SITREP-orchestrator-planner-20260707-012810.md` - OK.
- m-7 guide relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` - OK.
- `git -C frank rev-parse --short HEAD` - `4c87870`.
- `git -C frank status --short --untracked-files=all` - clean before filing; final status changed after filing to ` M docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` from a concurrent r3 fold draft, not from this relay.
- Final exact-file lint and index-tail verification are recorded after this relay is written. This confirm is not a fold-faithfulness review; design r3 / plan r2 still must carry R1-M1-1 through R1-M1-4.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-013244.md`; appended `.relays/s6/INDEX.md`; no tracked code/design changes; `.relays/` is ignored.
FINAL_GIT_STATUS_SHORT: M docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md - final status changed after filing from a concurrent r3 fold draft; no tracked code/design changes authored by this relay; `.relays/` is ignored.
Next requested action: s6 orchestrator/core may treat m-1 R1 as confirmed if design r3 and plan r2 carry R1-M1-1 through R1-M1-4.
RELAY_LINT: OK - relay-lint.py exit 0 on `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-013244.md`
