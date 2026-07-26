## DESIGN — ITEM A OPENED: the hashable Tier-HARD interface bundle. Master defines the extraction recipe + bundle schema + the Tier-HARD/Tier-SOFT boundary rule + `bundle_sha256` + the `bundle-soft-stability` negative fixture + `STEP-3-EXIT-FIXTURES.json` (carrying N910 · r7-mirror · env_digest-preimage-parity). Each owner NOMINATES its Tier-HARD `lock_payload` section over its settled base (bounded, pair-reviewed F73); master integrates → computes `bundle_sha256` → freezes the fixtures → VP item-A review → lane 4. Produces the bundle; does NOT lock it (that is lane 4). H-12 stands.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — opens a design lane-item + routes bounded owner nominations over already-settled bytes; it issues no DESIGN-lock / PLAN / T4 token, moves no ratified/frozen byte, and locks nothing (the lock is lane 4)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md
FROM: master.orchestrator-planner
TO: m-1.planner, m-2.planner, m-3.planner, m-8.planner, m-9.planner, m-10.planner
CC: operator, master.orchestrator-reviewer, m-1.implementer, m-2.implementer, m-3.implementer, m-8.implementer, m-9.implementer, m-10.implementer
SUBJECT: item A opened — each owner nominates its Tier-HARD `lock_payload` section over its settled base per the recipe below (bounded, pair-reviewed); master assembles `STEP-3-INTERFACE-BUNDLE.json` + `bundle_sha256` (stable under Tier-SOFT edits) + the soft-stability negative fixture + `STEP-3-EXIT-FIXTURES.json`, then VP item-A review → lane 4 re-lock

Lane 2 is closed (`step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000`), the nine settled bases are byte-verified and locked as the MVP interface. **Item A now builds the hashable bundle over them** — the §11 sequence's next step, ahead of lane 4 (the shorter re-lock) and lane 5 (T4).

## What item A produces (master-integrated deliverables)
1. **`master/STEP-3-INTERFACE-BUNDLE.json`** — the canonical Tier-HARD interface bundle over the nine settled bases + the normative join records + the carried obligations.
2. **`bundle_sha256`** = SHA-256 over the JCS-canonicalized bundle — the single hash lane 4 re-locks against. **Stable under Tier-SOFT edits by construction** (it hashes extracted Tier-HARD content, not raw file bytes).
3. **the `bundle-soft-stability` negative fixture** — proves a Tier-SOFT edit to any owner artifact leaves `bundle_sha256` unchanged (and, dually, that any Tier-HARD change moves it).
4. **`master/STEP-3-EXIT-FIXTURES.json`** — the frozen exit-test fixtures (the six-property legs) carrying the three lane-2-carried obligations as fixture legs.

## The extraction recipe — the Tier-HARD / Tier-SOFT boundary
Each owner's **`lock_payload`** section = its **normative, semantically load-bearing elements only** — the schema shapes, closed enums, digest formulas + preimages, presence/absence rules, wire-frame member sets, gate predicates, and the state/transition contracts that define *what the interface does*. **Tier-SOFT (excluded from the payload, hence from `bundle_sha256`):** prose, rationale, examples, section ordering, comments, formatting, and any wording that does not change a normative element. Each `lock_payload` is JCS-serialized into a canonical object; the bundle is:

```
STEP-3-INTERFACE-BUNDLE.json (JCS-canonical) =
{ bundle_version,
  owners:       { m-1, m-2, m-3, m-8, m-9, m-10 → each owner's Tier-HARD lock_payload },
  base_hashes:  { the nine settled-base SHA-256 (provenance anchors, NOT the hashed payload) },
  joins:        { §D two-sided (r17×rev16 legs) · §B sink (m-9⇄m-3) · item-E (R1@r24) · B-carriage (R2/R3) · item-C (m-9 §7 + rev16 §5 + m-1 C-confirm) },
  carried:      { N910, r7_mirror, env_digest_preimage_parity } }
bundle_sha256 = SHA-256( JCS( STEP-3-INTERFACE-BUNDLE.json ) )
```

`base_hashes` records provenance (which settled bytes each payload was extracted from) but the **payload content**, not the file hash, is what `bundle_sha256` covers — so a cosmetic edit that re-approves an owner artifact at a new file hash **without** changing its Tier-HARD elements leaves `bundle_sha256` stable (the soft-stability property).

## The three carried obligations — bound as fixture legs in `STEP-3-EXIT-FIXTURES.json`
1. **N910 = accepted documented MVP limit** — no sink record on the loss cut; loss operator-disclosed by m-10 `UNKNOWN_PROVIDER_OUTCOME` → `uncertain`. Fixture: the loss cut yields no `m3.b_sink.v1` record AND the operator-visible `uncertain` disposition is present (the disclosure is not silent).
2. **r7-mirror = v3-deferred** — the 2a/2b `unavailable`-DATA-P coverage gap, with the standing re-open caveat (an `xit-gov-1`-gating E3 predicate needing independent m-10-side 2a/2b resolution → m-3 surfaces it, master re-opens route-now).
3. **env_digest preimage-parity** — recipe at **m-1 §5 `:63`** (the m-1-owned JCS-over-`{name:value}` recipe + duplicate-name reject + non-UTF-8 pre-spawn reject), realized **byte-for-byte on two sides**: **m-9 §7 derivation** AND **m-3's E3 observer**. Fixture: the preimage-parity leg (m-9-bytes == observer-bytes for one logical env set) + the two construction-reject legs.

## What each OWNER nominates (bounded, pair-reviewed F73) — the six routings
Each of you nominates its **Tier-HARD `lock_payload` section** over its own settled base — a concise enumeration of your normative elements (no new design; an extraction over already-approved bytes), pair-reviewed, returned `FROM` your seat. Scope per owner:
- **m-1** (`d34a7c47…`): the env class-table + sanitization rule, the §6-C descriptor rulings, the descriptor-grain at-rest battery (§2.3), the carrier-negatives, and the env_digest preimage recipe (§5 `:63`).
- **m-2** (`c3a8cd61…` + cell `5ec7a3d2…`): the §5-E logical component (the two arrays' member set + ordering + refusal) and the `relay.submit` `canonical_resource` cell formula.
- **m-3** (r24 `651c9aec…`): the E0/E3 v2 schemas + cut-matrix + the five E3 verdict machines + `m3.b_sink.v1` classifier/row-state + the `logical_surface_digest` recipe-binding — with the **N910 + r7-mirror limits marked** in the payload (they are part of the honest contract).
- **m-8** (r5 `c0b7b488…` + r7 `734e44b7…`): the `frozen_core_digest` formula + present-iff-freeze rule + the `refusal_stage` 2a/2b discriminator + the provider-attempt reply/result schemas.
- **m-9** (r17 `01b885fe…`): the five §5-E member recipes, the item-D resume-seam frames (S-1/S-2/S-4 + the manifest consume + no-work gate), the item-C §7 executor descriptor, the item-B §8 carriage, and the item-E `logical_surface_digest` — the frames that went **normative at the §D co-sign** (§9 items 4/5).
- **m-10** (rev16 `3e3c5192…` + B/E rev3 `cd17db32…`): the run-wide D-4 carriage + the `parked_unknown_capacity_exceeded` terminal + the two frame assertions + the settlement-manifest three-class + the C ticket schema (§5) + the B/E carriage row (`m10_row_state` + digests).

**Nominate the Tier-HARD elements + explicitly mark what you deem Tier-SOFT** (so the soft-stability boundary is owner-drawn, not master-guessed). If you judge an element ambiguous, name it and I rule.

## The integration path (master, after the six nominations land)
Master assembles the six `lock_payload`s + the join records + the carried obligations into `STEP-3-INTERFACE-BUNDLE.json` → computes `bundle_sha256` → authors the `bundle-soft-stability` negative fixture (a Tier-SOFT edit leaves the hash; a Tier-HARD change moves it) → freezes `STEP-3-EXIT-FIXTURES.json` → routes the assembled bundle + `bundle_sha256` to the **VP for the item-A review** → on approve, **lane 4** (the shorter re-lock over `bundle_sha256` + the whole-file-hard owner contracts + the Master+VP interface-lock). Also folded here: the owed **ARCHITECTURE.md D7/`relay.submit` mechanism-prose consolidation**.

## Dependencies / ordering
The six nominations are **independent and parallel** (each over its own settled base; no cross-owner dependency — the joins are already settled/normative). Master's assembly is the barrier. The nominations bind the **settled bases only** (never ancestry: not m-9 r14/r15/r16, m-3 r19, m-10 rev14).

## Boundaries
Item A **produces** the bundle; it does **not** lock it — the interface-lock is lane 4. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy is issued or implied. This dispatch moves no ratified/frozen byte and binds only the nine byte-verified settled bases. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Verification
The nine settled-base anchors re-verified on disk this session (all MATCH): m-1 `d34a7c47…`, m-2 `c3a8cd61…` + cell `5ec7a3d2…`, m-3 r24 `651c9aec…`, m-8 r5 `c0b7b488…` + r7 `734e44b7…`, m-9 r17 `01b885fe…`, m-10 rev16 `3e3c5192…` + B/E rev3 `cd17db32…`. Underpinning amendments UNMOVED: rev12 `1125b0a0…`, m-3 schema rev3 `9e874df8…` + contract `6e2abe40…`, §D-settlement rev4 `1fa71cb8…`. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this item-A opening dispatch + one INDEX.md row; no design byte moved, no `frank/` action, no lock/PLAN/T4 issued, no bundle authored yet (assembled after the nominations).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: each of m-1/m-2/m-3/m-8/m-9/m-10 nominates its Tier-HARD `lock_payload` section (+ its Tier-SOFT boundary) over its settled base, pair-reviewed, `FROM` its seat; master assembles `STEP-3-INTERFACE-BUNDLE.json` + `bundle_sha256` + the soft-stability fixture + `STEP-3-EXIT-FIXTURES.json`, then routes to the VP for the item-A review → lane 4 (re-lock) → lane 5 (T4). H-12 stands.
