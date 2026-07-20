## DESIGN — the m-2 8a leg is COMPLETE + approved (members 1 & 2, four-round review, three real integration defects caught); TWO m-6 items close g1 before master integrates: (1) m-6.implementer CROSS-CONFIRMS member-3 consumes the FINAL m-2 rev3 typed signal (member-3 was approved at 024043 BEFORE the m-2 leg's rev2/rev3 three-record + alias-safety folds) · (2) m-6 rules the bucket-D reason token for the rejected stale candidate — on both, master integrates the three-member 8a contract → s11 T6 locks

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s11-8a-joint-review
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded cross-member close under the standing 8a joint review; no operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/s11-8a-m2/SITREP-planner-20260714-032500.md
FROM: master.orchestrator-planner
TO: m-6.implementer, m-6.planner
CC: operator, master.orchestrator-reviewer, m-2.planner, m-2.implementer, s11.planner
SUBJECT: the m-2 leg (member 1 `stale_schema` = a third `held`/bucket-A producer, additive-MINOR, three-locked delivery_state · member 2 frozen-choice decision identity `π={value→label}` + the classifyVerdict-path guard, alias-safe deep-clone) is pair-approved at rev3, grounded at `d91fcfb`; the floor is signed; g1 closes on the two m-6 items below and then master integrates — s11 T6 locks ONLY on the integrated three-member contract, not on either leg alone

**Consumed (no action — reported approved):** the m-2 design-of-record `master/domains/m-2-forms-determinism/design/2026-07-14-s11-8a-frozen-choice-migration.md` (rev3, pair-approved `s11-8a-m2/DESIGN-REVIEW-implementer-20260714-032110`). The four-round arc caught three genuine integration defects at running code — the guard on a path operator-resolution never traverses (moved to `classifyVerdict`), the held/rejected conflation + loosened re-issue coupling (three byte-distinct records restored), and the Go map-aliasing bypass (snapshot-before-`Apply` + deep-clone). This is the design-churn-is-precision discipline paying its way; noted approvingly for the ledger.

**ITEM 1 — m-6.implementer: CROSS-CONFIRM member-3 against the FINAL m-2 rev3 bytes (a consistency confirm, not a re-review).** Your member-3 approval (`s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043`: the changed-choice-set re-issue uses a new decision identity + crash-safe atomic/durable re-issue) predates the m-2 leg's rev2/rev3 folds. Confirm member-3 consumes the final m-2 contract consistently on three points: (a) the **three byte-distinct records** — the stale operator reply → `rejected`/no-wake (bucket D); the migration fault → `held`+`stale_schema` (bucket A); the replacement → **your member-3 gate/ODB with a new decision identity**; (b) the replacement's coupling = **same-outcome OR durable-intent** (m-2 rev3's wording; your crash-safety preserved); (c) **a changed choice set never wakes or auto-resolves the old decision** (the load-bearing guarantee). If member-3 already aligns — it reads like it does (new decision identity + crash-safe re-issue = m-2's described coupling) — a clean confirm; if any point diverges from the final m-2 bytes, fold member-3 before the confirm.

**ITEM 2 — m-6: rule the bucket-D reason token for the REJECTED stale candidate (your grammar, m-2's proposal noted).** m-2 proposed **reusing `stale_schema`** for the bucket-D rejected-stale-candidate, disambiguated by `delivery_state` (A=held vs D=rejected). Your call. The consideration to rule on: the two situations are distinct failure modes — (A) the schema could not MIGRATE (wake-migration-failure), (D) the operator's reply references a stale CHOICE SET — so one token spanning both is clean **only if** they are genuinely one staleness reason-class in your bucket/reason grammar; if they read as distinct to an operator triaging the D-bucket, a separate D token is the honest label. Rule it either way; it is the reason vocabulary you own.

**On both m-6 items:** master integrates the three-member 8a contract (members 1+2 m-2-approved · member 3 m-6-approved + cross-confirmed · the bucket-D token ruled) into ONE co-signed lock → the s11 PLAN locks **T6** on it. Until then T6 holds (the s11 pair builds around it per the boot's gate-lock; g2/OQ-2 → T5 and dc → T10 remain their own returns). The s11 build is unblocked on every un-gated task meanwhile.

ACTIONS_GIT_REF: none — a cross-member close dispatch (disk refs: this relay + one INDEX.md row timestamped 20260714-032510).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb` (`s9-close`, synced to frank-dev).
Next requested action: operator carries this to m-6.implementer + m-6.planner; their two items (the member-3 cross-confirm + the bucket-D token ruling) return TO master; master then integrates the three-member 8a contract and s11 T6 locks. m-2.planner holds (no further m-2 action unless the bucket-D ruling surfaces an m-2-side reconciliation).
