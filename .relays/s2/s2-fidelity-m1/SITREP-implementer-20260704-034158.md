## SITREP - m-1 fidelity verdict on S2 store-touch proposal surface

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-fidelity-m1
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: SITREP-orchestrator-planner-20260704-023129.md
FROM: m-1.implementer
TO: s2.orchestrator-planner
CC: operator, s2.orchestrator-reviewer
SUBJECT: m-1 fidelity verdict - S2 store proposal surface needs bounded fold before dispatch

VERDICT: must-revise before any S2 implementation dispatch.

This is a store-shape fidelity finding, not a rejection of the S2 mechanics. The pair correctly isolated design section 4 as a proposal surface, and most items are acceptable once the canonical record homes are made exact. No S2 implementation dispatch should go live until the revisions below are folded into the design/plan and re-reviewed narrowly.

Basis read:
- Incoming fidelity request: `.relays/s2/s2-fidelity-m1/SITREP-orchestrator-planner-20260704-023129.md:22-36`.
- S2 design proposal surface: `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md:125-136`, with mechanics at `:47-68`, `:71-90`.
- S2 plan r2: `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md:63-77`, `:121-132`, `:137-142`, `:155-160`.
- Locked m-1 contract: `the m-1 trust/identity design-of-record (2026-06-28) :122-145`.
- Locked m-7 recovery/genesis/GC text: `the m-7 conductor-core design-of-record (2026-07-01) :89-104`, `:134-137`.
- Current code shape: `internal/record/record.go:16-32`; current internal derived records already use `From: "system", Role: "system"` at `internal/gate/derived.go:59-67`, `:105-114`.

## Required revisions

### F-M1-1 - Canonical field homes and internal provenance are under-specified

The plan currently puts `schema_version` inside genesis headers (`s2-slice-2-plan.md:66`), while m-1's locked system-field contract makes `schema_version` a conductor-stamped system field, and the current canonical record shape stores it in `Envelope.SchemaVersion` (`record.go:16-25`). The same ambiguity applies to the proposed internal records: genesis, incident, GC-marker, and derived outbox records all need exact envelope provenance, not empty or ad hoc envelope fields.

Fold this table into the design/plan before dispatch:

| surface | m-1-required home |
|---|---|
| `schema_version` | `Envelope.SchemaVersion` only. Do not duplicate as `Headers["schema_version"]`. |
| `record_kind` | `Headers["record_kind"]`. Tokens approved for this S2 surface: `owed_item`, `owed_disposition`, `genesis`, `incident`, `gc_marker`. |
| owed item fields | `Headers["owner"]`, `Headers["source"]`, `Headers["target_surface"]`, `Headers["disposition_path"]`; body may be empty or narrative, but the open-set projection and validation must use the headers. |
| disposition reference | `Headers["disposes_owed"]`. Duplicate disposition remains a typed reject. |
| quarantine incident reference | `Headers["quarantined_ref"]` and `Headers["failure_class"]`. |
| GC marker payload | `Body` JSON naming the collected journal segments; `Headers["record_kind"] = "gc_marker"`. |
| derived outbox payload | `Body` JSON. This is approved and is required for canonical sufficiency. |

Internal conductor-authored records must use a reserved internal provenance convention, not an unstamped envelope:

```
From: "system"
Role: "system"
DeliveryState: "accepted"  # genesis, gc_marker, normal derived records
DeliveryState: "held"      # quarantine incident
SchemaVersion: 1
```

`system` is not a lane address and must never be accepted from the public `submit` path. It is reserved for conductor-internal records produced by `store.Init`, recovery, the loop, and GC. Normal owed items and owed dispositions ride the ordinary submit path and are stamped from the submitting channel. The real `OI-S1-F11-SWEEP` owed item remains `FROM=operator` via the operator channel, as S2 already states.

Genesis shape, after fold:

```
Envelope.RelayID = "genesis"
Envelope.DispatchID = "genesis"
Envelope.From = "system"
Envelope.Role = "system"
Envelope.DeliveryState = "accepted"
Envelope.SchemaVersion = 1
Headers["record_kind"] = "genesis"
Headers["config_digest"] = <digest>
Headers["address_space_seed"] = <seed>
Headers["created_ts"] = <timestamp>
```

### F-M1-2 - `read(relay_id)` needs two typed corruption states, not one

`quarantine/` as a store-root member is approved, and name-preserving eviction is approved. The typed read behavior needs one revision:

- Live read of a corrupt file still present at `records/<relay_id>.json`: return `checksum-mismatch`, include `relay_id`, remain path-free, and enqueue the internal quarantine-disposition command. The reader still does not mutate.
- Read after the record has been evicted to `quarantine/<relay_id>.json`: return `record-quarantined`, include `relay_id`, `incident_id` when present, and `failure_class=checksum-mismatch`, remain path-free. Do not return `checksum-mismatch` for the post-eviction state.
- `read(incident-<relay_id>)` serves the HELD incident record. `Records()` and projection rebuilds operate over clean canonical records only and must not re-include quarantined bytes.

This keeps `checksum-mismatch` as the detection event and `record-quarantined` as the stable post-disposition store state. It also avoids creating a fourth delivery outcome; both are API error classes, while the incident record's envelope remains `held`.

### F-M1-3 - Pinned config placement must be store-root, not externally adjacent

The proposal's "store-adjacent" wording is not precise enough for m-1. Because the S2 design and plan both bind recovery to store-only replay, the digest inputs that genesis pins must be conductor-owned store-root members, not external docs or caller paths.

Required shape:

```
config/fieldspec/registry.json
config/engine.json
config/manifest.json   # optional persisted manifest; digest algorithm may also compute this in memory
```

`store.Init` may accept operator-supplied source paths, but it must materialize the pinned bytes under the store root before writing genesis. Phase-0 validation then compares genesis against the store-root config members, never against an outside path. The config members are not canonical records, not projections, not seat-visible resources, and not domain-stamped sections in S2.

## Per-item verdicts

| item | verdict | m-1 answer |
|---|---|---|
| 1. `owed_item` | approve after F-M1-1 fold | `record_kind` and the four typed fields live in canonical headers. Normal submit stamping applies. |
| 2. `owed_disposition` | approve after F-M1-1 fold | `record_kind=owed_disposition`, `disposes_owed=<owed relay_id>` in headers. Duplicate disposition reject is correct. |
| 3. genesis record | must-revise | Fixed `relay_id=genesis` is approved, but exact envelope/provenance and `schema_version` home must be folded as above. |
| 4. incident record | approve after F-M1-1/F-M1-2 fold | HELD-class incident is right: `From/Role=system`, `DeliveryState=held`, headers `record_kind=incident`, `quarantined_ref`, `failure_class`. |
| 5. `quarantine/` | must-revise | Store-root `quarantine/` and name-preserving rename approved. Post-eviction `read` returns `record-quarantined`, not `checksum-mismatch`. |
| 6. GC-marker record | approve after F-M1-1 fold | `record_kind=gc_marker`; body JSON names collected segments; envelope `system/accepted`; canonical records are never GC'd. |
| 7. journal segmentation | approve | Layout approved as `journal/intake/000001.jsonl` and `journal/redo/000001.jsonl`, six-digit zero-padded sequence, highest segment active, rotation by pinned config bytes. |
| 8. derived-record body embedding | approve | Required for canonical sufficiency. The body is canonical record content; no envelope/layout objection after F-M1-1. |
| 9. `projections/owed/OPEN.md` | approve | Derived, rebuildable projection under `projections/owed/OPEN.md`; it carries no independent authority. Empty-at-genesis is valid. |
| 10. pinned config + `engine.json` | must-revise | Store-root `config/` members required. External source paths are init inputs only, not phase-0 recovery dependencies. |

## Dispatch condition

Do not dispatch S2 implementation until:
1. The design/plan fold F-M1-1 through F-M1-3.
2. The fold keeps the already-approved S2 mechanics otherwise unchanged.
3. A narrow re-review relay points to the folded lines and confirms the table above is implemented in the plan.

ACTIONS_GIT_REF: none - report-only fidelity relay; no code or design file changed by m-1.
FINAL_GIT_STATUS_SHORT: clean before relay write; expected git status remains clean because `.relays/` is gitignored.
