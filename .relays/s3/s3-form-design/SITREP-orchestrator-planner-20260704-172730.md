## SITREP — de-provision supplement to the s3-form DESIGN dispatch: all three question threads RESOLVED (guide ×3 confirmed · m-7 ×3 answered · master scope ruling (a)-DEFER ratified with conditions)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-form-design
PARENT_DISPATCH_ID: s3-form-audit
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s3-form-design/DESIGN-orchestrator-planner-20260704-170903.md
FROM: s3.orchestrator-planner
TO: s3-form.planner
CC: s3-form.implementer, s3.orchestrator-reviewer, operator
SUBJECT: de-provision the three provisional DESIGN sections — every pending thread is answered from locked text with zero amendments; the twelve constraints stand unchanged; grill agenda narrows (resolved rows enter as resolved, not re-asked); one new S3-owned bounded choice + one ratified posture + one materialized owed item

The three threads your DESIGN dispatch marked provisional are all resolved (reply relays, all lint-clean, my runs — read them in full; this supplement is the routing digest, the relays are the record):
- guide: `../.relays/s3/s3-guide-q1/SITREP-planner-20260704-173000.md`
- m-7 consult: `../.relays/s3/s3-consult-m7/SITREP-planner-20260704-171546.md`
- master scope ruling: `../.relays/s3/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md`

### Fold into the design (de-provisions constraint 8's pendings + constraint 4's adjudication + constraints 7/9's simulation posture)

**Guide Q1 — the GRILL_REQUIRED row shape (constraint 8):** `layer: header · owner: agent_enum_pick · type: bool {yes,no} default no · fill_constraints: monotonic` (RAISE-toward-yes; effective = OR/MAX across raisers — orchestrator dispatch ∪ pair-Planner; no seat lowers a requested grill) · **`gate_referenceable: true`** (the deliberate departure from the §4 default — forced by the m-6 §5 :105 meeting-lane route atom; R2-safe, not model-identity) · `seat_scope`: no seat denied `yes` (not an authority gate) · `consumers: [form_renderer, form_validator, human_surface]` · `lineage_role: none`. The formal registry row lands in YOUR design under the implementer's review — the guide supplied the shape, not the row. **One bounded S3-owned choice (yours, argue it in the design either way):** wire `GRILL_LOCK_ID`/`GRILL_SOURCE` as dependent-required (`required_when field:GRILL_REQUIRED == true`)? If adopted: R2-safe, adds `lineage_engine` to consumers.

**Guide Q2 — the replay adjudication (constraint 4):** CONFIRMED — the step-gated fixtures land in **"caught"** under a reconstructed `layer_present:observe = true` evaluation context; no third bucket. **Two binding naming guardrails:** the label reads exactly "caught — surviving §10b rule, fires under the reconstructed `layer_present:observe` context", never "live-caught on a Step-1 form" (the rule is correctly step-gated OFF live, AC17(a)); and such rows are never counted toward a live-Step-N coverage tally. Design the harness's context parameter + report vocabulary to make the qualifier structural, not prose.

**Guide Q3 — the R2 fixture grain (constraint 8):** flag-based grain CONFIRMED S3-sufficient: negatives prove the grammar rejects any `field:<id>` naming a `gate_referenceable: false` column, over the S3-declared columns **including a test column that genuinely carries BOTH `gate_referenceable: false` AND the model-identity mark** (proving the model-identity ⇒ non-referenceable ⇒ rejected chain specifically — a generic false-flagged column is not enough). Concrete `chosen_model` + live bucket-proxy negatives defer to S4/m-4 (their per-column declarations). Carry the S4 boundary flag in the design's out-of-scope notes.

**m-7 Q1 — member shape (constraints 1/7):** **(a) one larger `fieldspec` member** under the existing single top-level digest. Attribution honesty rides: the member is s3-built to the m-2 locked shape and m-2-guide-reviewed — it is NOT a CQ-4b m-2-stamped section; never present it as one. The member's internal JSON may carry its own metadata block (Q2's version home) — self-description claims nothing about section composition.

**m-7 Q2 — the three axes (constraints 7/9):** top-level config digest = integrity identity, in genesis, the ONLY integrity root · **registry version = change attribution, INSIDE the member content** (a `version` + minimal provenance block in the registry JSON; never in genesis; not derived-only) · record `schema_version` = the unrelated per-record axis. **Two fixture consequences (adopt verbatim):** (i) a **positive** fixture — old store + changed registry ⇒ phase-0 fail-closed *serving reads, summoning operator* (the disposition, not just the error; the F7/NF-S15 property working); (ii) drift fixtures simulate registry change as **restart-with-new-store** (seat holds a prior-generation render ⇒ stale digest ⇒ `re-render` bounce). Master's sharpening is the claim wording: under the locked no-hot-reload model, restart-with-new-store IS the true semantics of a registry change — state it as "live change does not exist in this config model; the fixture tests the mechanism that does," never as a testing limitation.

**m-7 Q3 — serving the form + render placement (constraints 5/9, F-P5, F-P2):** YES — extend the existing socket protocol with a tools/describe-grade response carrying the per-seat rendered form + form digest, replacing the static descriptions; the MCP-proper shape inherits it at the wire-up slice. Closing the `Validate(…, "")` dead path with the real digest is REQUIRED (the drift fixtures must bite). **Render sits channel-handler-side** (locked §2.1: render advisory, loop authoritative — never a loop turn), inputs = immutable registry + connection-fixed SeatMeta (never payload) + phase. **Bind the form digest to (config digest, seat, phase)** so re-render is exact across store generations. The conditional-grant render's lineage read touches **committed records only** (never FIFO/in-flight); TOCTOU closes by construction (stale-positive rejected by the loop's authoritative check at submit; stale-negative just re-renders). I-PH extends to the served schema payload: candidate sets fine; no store/config paths, no config values.

**Master ruling — (a) DEFER, ratified (the constraint-3/7/9 posture + a new claim rule):** S3 does NOT build the §7 config-change record. Five conditions bind us: (1) **fresh-store qualifier on every claim surface** — exit-gate line reads "registry live end-to-end (fresh-store)"; disposition table + README/claim text state "registry rides `store.Init`; evolution on an existing store awaits the §7 record"; (2) `OI-S3-CONFIG-CHANGE` is materialized — ledger record at `docs/sprints/2026-07-04-s3-slice-3/results/OI-S3-CONFIG-CHANGE.md` (typed payload verbatim from the ruling; the optional real-store submit stays the operator's channel); (3) disposition owner = the wire-up slice, hard backstop = before any store is declared persistent; (4) when built it carries m-7-guide/m-1-fidelity/crash-class conditions — recorded, not yours to build; (5) the s2-design "(S3)" forward-pointer is superseded by the ruling (fold rides the next m-7-guided touch — not yours). If you find a SECOND item of this shape, escalate again — the ruling covers the §7 record only.

### Standing state after this supplement

- The twelve hard constraints stand unchanged; the provisional markers come off.
- **Grill agenda narrows:** the resolved-by-guide/consult/master rows enter the grill as resolved (not re-asked); the remaining operator agenda floor = the typed-header/record representation, the registry artifact shape (within (a)), the disposition-table artifact form, on-disk commitments — plus the one S3-owned choice above if you want the operator's read on it.
- Nothing further blocks the design lock except the grill + your implementer's DESIGN-REVIEW per the dispatch.

Next requested action: fold, de-provision, proceed to the grill + design completion per the standing DESIGN dispatch. No new authority granted or needed by this supplement.

ACTIONS_GIT_REF: main@8d5a0e6 — the reconciliation commit carrying the two tracked files (results/OI-S3-CONFIG-CHANGE.md + the RECONCILE.md entry); this supplement + one INDEX row live under gitignored .relays/.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 8d5a0e6, `git status --short` empty after the commit)
