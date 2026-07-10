## SITREP - m-1 narrow fidelity re-review confirm for S2 folded store shapes

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-fidelity-m1
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: SITREP-orchestrator-planner-20260704-043100.md
FROM: m-1.implementer
TO: s2.orchestrator-planner
CC: operator, s2.orchestrator-reviewer
SUBJECT: m-1 narrow confirm - F-M1-1 through F-M1-3 folded; homes table implemented in plan

FIDELITY_VERDICT: approve.

This is a narrow confirmation of the folded lines requested in `SITREP-orchestrator-planner-20260704-043100.md`. I did not re-open the already-approved S2 mechanics or the approved items 7/8/9 from the prior fidelity verdict. The three blocking m-1 findings from `SITREP-implementer-20260704-034158.md` are folded in the live design/plan surfaces, and the plan implements the homes table.

## Findings

None blocking.

## Confirmation

### F-M1-1 - canonical field homes + internal provenance

Confirmed.

- Design D-3 now defines the genesis record with `Envelope{RelayID: "genesis", DispatchID: "genesis", From: "system", Role: "system", DeliveryState: "accepted", SchemaVersion: 1}` and headers `{record_kind: "genesis", config_digest, address_space_seed, created_ts}`; it explicitly keeps `schema_version` envelope-only at `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md:49`.
- Design D-3 reserves `system` as a non-lane address, rejects `seat.Mint("system")`, and keeps owed items/dispositions on the ordinary submit path at `:50`.
- Design section 4 carries the m-1 homes table through the itemized surface: owed item headers at `:128`, disposition headers at `:129`, genesis at `:130`, incident at `:131`, GC marker body/header/envelope at `:133`, and pinned config at `:137`.
- Plan Task 2 implements the genesis envelope/header split, explicitly says no `schema_version` header, and adds the `ErrReservedSeatName` leg at `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md:67` and `:69`.
- Plan Task 9 keeps owed fields in headers and the projection reads headers at `:125`.
- Plan Task 10 keeps the GC-marker shape as `Headers{record_kind: gc_marker}`, `system/system/accepted`, and body JSON at `:133`.

### F-M1-2 - two typed corruption read states

Confirmed.

- Design D-5 now separates `checksum-mismatch` as the detection event for a corrupt file still in `records/` from `record-quarantined` as the stable post-eviction state, with `incident_id`, `failure_class`, readable incident records, clean-only `Records()`, and no fourth delivery outcome at `s2-slice-2-design.md:64`.
- Plan Task 3 implements `store.ErrChecksum{RelayID string}` and `store.ErrQuarantined{RelayID, IncidentID, FailureClass string}`, keeps incidents readable, and keeps `Records()`/projection rebuilds clean-only at `s2-slice-2-plan.md:75` and `:77`.
- Plan Task 11 carries the same split through the read-handler/channel frames at `:141` and the live fixture leg at `:143`.

### F-M1-3 - pinned config placement is store-root

Confirmed.

- Design D-3 places pinned members under `config/fieldspec/registry.json` and `config/engine.json`, with optional `config/manifest.json`, and requires `store.Init` to materialize source bytes under the store root before genesis; phase-0 validation never depends on an outside path at `s2-slice-2-design.md:51`.
- Plan Task 1 loads store-root config members at serve/recovery time and treats outside paths as init sources only at `s2-slice-2-plan.md:58`.
- Plan Task 2 materializes sources into `<root>/config/...` before the genesis commit and validates only against store-root members at `:67` and `:69`.
- Plan Task 11 keeps config flags as init-time sources only; the serve path ignores outside config paths and loads from store-root members at `:141`.

## Verification notes

- Incoming re-review relay exact-file lint: `OK .relays/s2/s2-fidelity-m1/SITREP-orchestrator-planner-20260704-043100.md`.
- Current `frank` HEAD observed as `62afa8e`, which adds the S2 reconcile routing note after `ca23a44`; `git -C frank status --short` was clean during review.
- Pair r6 plan review is approved at `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-041850.md:15` and confirms the residual labels folded at `:33-44`; its verification block includes exact-file lint and `go test -count=1 ./...` at `:52-60`.
- Targeted probe over the design/plan found no live `m-1-approved` / `m-1-APPROVED` / `m-1-fixed` labels in the folded surfaces; remaining approval-state text is either conditional pending-confirm language or history.
- Targeted probe found no `schema_version` header home in the folded genesis plan; the live references either say envelope-only or assert "NO `schema_version` header".

Non-blocking hygiene note: the design status/rev-history line still says the pair narrow re-review and m-1 narrow re-review are pending (`s2-slice-2-design.md:5`, with a similar historical note at `:204`). The r6 pair approval is on record and this relay closes the m-1 fidelity confirm, so this is status prose cleanup for the S2 owner, not a remaining m-1 store-shape blocker.

## Gate result

The m-1 fidelity gate for F-M1-1 through F-M1-3 is satisfied. Remaining sequencing is the S2 pair/orchestrator process already named in the plan: scope-diff/all-in check, then any implementation-dispatch authority must parent the operative r6 plan-review approve.

ACTIONS_GIT_REF: wrote `.relays/s2/s2-fidelity-m1/SITREP-implementer-20260704-043529.md` and appended `.relays/s2/INDEX.md`; no source/test/sprint-spec edits by m-1.
FINAL_GIT_STATUS_SHORT: none - clean tree
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-fidelity-m1/SITREP-implementer-20260704-043529.md`
