# 2026-06-28 — m-2 AUDIT (planner): form schema + field-ownership vs the stock protocol

**Report of record (lint-clean):** [`master/relays/c1-audit-m-2/AUDIT-planner-20260628-152808.md`](../../../relays/c1-audit-m-2/AUDIT-planner-20260628-152808.md). This is the durable domain home; the relay file carries the full report. Implementer audit + reconciliation pending.

## Verdict at a glance (AC1)
**PRIMARY_BUCKET: still-open** — a thin net-new layer over a large promote-and-formalize base.
- **still-open (net-new):** the explicit field-ownership model; fill-time authority (forbidden option absent); one canonical schema artifact tool/courier/linter all read; dynamic required-set as a declared PHASE×CEREMONY_TIER function.
- **already-closed (promote, DO NOT rebuild):** the ~50-field catalog, all enums, the required-when conditionals, the phase↔authority map, the lineage relationships — all already in relay-lint.py + protocol.md.
- **product-overlapped:** system-fill of FROM/PARENT (m-1), observe-fill of evidence (m-3), routing values (m-4), email projection (m-6) — m-2 declares the slot, siblings own the mechanism.
- **recommended-next:** promote the catalog/enums/required-when into one schema; add the ownership column + fill-time render; dissolve the ~32 prose defenses; port the lineage engine; design render+validate.

## Dissolve/survive map (AC2)
62 relay-lint.py checks → **~32 DISSOLVE** (prose/markdown/token-lexical layer) / **~17 SURVIVE as form-validation** (enum/required/phase-authority/scan-consistency) / **~13 SURVIVE as a cross-relay lineage engine** (design-review gate, pair-planner dispatch lineage, merge-claim lineage, OUT→IN drift, orchestrator-review visibility).
**GATE-1 named:** (1) the lineage engine is NOT preservable by a per-relay form — survives separately, and is *strengthened* to forgery-robust once the courier system-fills PARENT (m-1 dep); (2) `DISPATCH IMPL`/`DISPATCH MERGE` lexical detection dissolves but its **authority** must survive as a seat-scoped + phase-scoped form field.

## Design recommendation (AC3)
`field · owner · type · required-when · enum-set · consumer`, carried as **JSON-Schema core + `x-owner`/`x-seat-scope`/`x-consumer` extensions** (one file = the single source). Render = owner-determined affordance per seat (forbidden options omitted at render → fill-time authority; precedents: OpenAI Structured Outputs, Swarm handoff-as-tool). Validate = two send-time pre-flights (form-conformance + m-3 observe-gate). Hold canonical-iff-consumed strictly so free-text stays free (counter-evidence: the "Format Tax").

## Boundary contract (AC4)
Schema must express m-3 evidence fields (system/observe-owned), m-4 routing-record fields (record-kind reuses DESIGN_RECORD_KIND shape), m-6 HUMAN_GATE (system, monotonic-raise) + email projection + Owner-Decision-Brief sub-schema. All four consumer edges named; no field left unrouted.

## Open questions → design/operator
Q1 carrier (JSON-Schema+ext vs DSL); Q2 schema versioning/evolution (operator-judgment, m-1+m-6 cross-cut); Q3 sanctioned overflow channel; Q4 bounded required-when predicate; Q5 implementer reconciliation of the dissolve/survive split.
