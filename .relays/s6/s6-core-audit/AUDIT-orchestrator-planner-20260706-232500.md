## AUDIT — s6-core paired independent audits: the transport-fix surface at `s5-close` (spec-to-exit-gate map · 4-bucket verdicts · promote-don't-rebuild · fidelity surfaces · claim probes)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s6-core-audit
PARENT_DISPATCH_ID: s6-dispatch
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
FROM: s6.orchestrator-planner
TO: s6-core.planner, s6-core.implementer
CC: s6.orchestrator-reviewer, operator
SUBJECT: AUDIT (paired, independent, read-only) — map the co-signed s6 set onto the code at `s5-close`; verify every cited root cause yourself; bucket every IN cluster; inventory what exists to promote vs build; enumerate the m-1/m-2/m-7 fidelity surfaces; probe the claim boundaries; NO coordination before filing

**What this is.** The s6 AUDIT phase: both pair seats independently audit the transport-fix surface before any DESIGN. Do not coordinate or read each other's audit before filing — reconciliation is my job. Read-only: no source/test/tracked-doc edits, no branch. Baseline = `frank/` code surface `s5-close` (`7e5c527`; HEAD `main@e9ed6ab` is docs-only ahead — verify that yourself).

**Spec of record (read-only):** `master-docs/master/S6-AMENDMENT-SET-2026-07-06.md` (r3) + the m-1 §A–§F.1 / m-7 r5 / m-2 codec (+§11) amendment docs + `master-docs/master/GRILL-LOCK-parenting-fork-2026-07-06.md`; the story in `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md`; the fences + exit gate condensed in `docs/sprints/2026-07-06-s6-slice-6/ROADMAP.md`. Where older §C4 prose conflicts, the amendment docs win. A spec problem (contradiction, unbuildable clause, missing seam half) is an ESCALATION to me — never self-amended, never silently absorbed.

### Deliverables (each seat, one lint-clean AUDIT relay under this DISPATCH_ID)

1. **Spec-to-exit-gate map.** Every exit-gate line (ROADMAP §Exit gate: the GRILL_LOCK three · FX-A1a..FX-B1g (18) · m-1 §E set · m-2's set · the E2 floors · the step-exit legs · the honesty lines) grounds in named set/constituent text (file:line or §). Flag any gate line with no spec ground, and any spec obligation with no gate line — both are escalation-shaped findings, not silent fixes.
2. **Root-cause re-verification in code (the standing bar).** The amendment docs cite root causes at `main 67ee23e`/`7e5c527`. Re-verify each in the live code yourself — at minimum: the render/validate/gate/delivery four-judge divergence (F6/F7: `lineage.go` comma-splits vs `ParseTyped` vs `projections.go DeliveryRecipients` silent-drop vs `render.go` recipient_picker) · the F11/F4 candidate mechanics (`ActiveLineageCandidates`, `turnContextForSeat` — what feeds `WokenOn`, and whether rejects can anchor) · F9 (`journal.go Append` — dedupe vs replay; id minting) · F5 (`digestRenderedForm` + the per-field-vs-class `DigestExempt`) · F13 (`validateRecordKind`'s membership judging) · F17 (`checkReviewerVisibility` waiver semantics) · F14/F15 (lock absence; admin-time mint) · F16/F3 (shim reconnect; bounce-detail asymmetry). Cite YOUR line numbers. Anything the spec asserts about the code that you cannot reproduce is a finding.
3. **4-bucket verdicts** (still-open / already-closed / product-overlapped / recommended-next + PRIMARY_BUCKET), one per IN cluster: branch-A parenting · the codec · A-1 digest · A-2 intake · A-3 live mint · A-4+§D lock · §B projection/anchor · §C waivers+retraction · F13 three-layer · D-1 · D-2 · B-1 lifecycle+roster · B-2 boot form · B-3 derived activation · the registry pass.
4. **Duplicate/promote-don't-rebuild inventory.** Name what exists and its reuse posture — candidates I expect you to weigh (audit them, don't take my word): the S1/S2 crash harness + crashpoint registry (A-2/A-3/A-4 crash legs); the S3 replay harness + archived-store replay leg from s5 (the F11 dogfood-replay fixture's natural home; archive at `~/frank-archives/frank-team-store-s5-dogfood-20260706`, 41 records); the existing per-field `DigestExempt` mechanism (A-1 promotes it to a class); `ParseTyped`'s `address_list` decode (the codec's seed — §2.1 says unify, not rewrite); the s4 per-recipient `PushTo` delivery + redacted-read view pattern (B-1.3 roster + §B audit-scope precedents); `seat.Manager`/`mintSeat` (A-3 pivots its call site, custody unchanged); the derived-work-completion mechanism (A-3's endpoint leg); the owed/`resolves_gate` id_ref idiom (§C `retracts`). Rebuilding any of these is a red flag.
5. **Fidelity-surface enumeration.** Every store/lineage/waiver/lock/activation touch (→ m-1 fidelity packet at PLAN time), every codec/registry/boot-form/render/validate touch (→ m-2), every engine/loop/lifecycle/runtime touch (→ m-7 guide). This is the packet-routing input — completeness matters more than depth.
6. **Claim-boundary probes.** (a) the byte-exact `{accepted, rejected, held}` enum + three-verb surface — name every point the diff could threaten them; (b) I-PH over the NEW surfaces (roster, boot bounces, lock refusals, hint flags) — what fixture class proves path-free `Field:Class`; (c) the honesty ceiling — any set clause that could tempt an over-claim (activation ≠ identity upgrade; hint fallback ≠ validation; lock = confusion-resistant ops, D5 residual); (d) [VP-W3] — grep for any surviving activation-marker language anywhere in the build plan inputs; (e) the no-perf fence — name any set clause that could be misread as licensing optimization.
7. **Battery + status proof.** Your own uncached baseline run (boot runs citable if same code surface, state so) + `FINAL_GIT_STATUS_SHORT` from a fresh check.

### Hard criteria / gates
- Reject/narrow gate: if any IN cluster is already-closed or product-overlapped at the chartered grain, say so with evidence — do not design work that exists.
- Escalation-shaped findings (spec contradiction, locked-contract conflict, OUT-item entanglement, cross-slice collision) are named as such and routed to me — never resolved in-audit.
- Boundary contract: not applicable at AUDIT (read-only); the seam halves you enumerate in deliverable 5 feed the DESIGN-phase contracts.
- Operator-judgment items: flag any found; expected none at this phase.

**OUT (restated):** Step-2 observe · Step-3 routing execution · engine perf work of any kind · new seat verbs · federation · dogfood-in-slice · any governance-doc edit.

**Format:** lint-clean relay per seat under `.relays/s6/s6-core-audit/` (`AUDIT-planner-…` / `AUDIT-implementer-…`), stamps real wall-clock strictly > this relay's; INDEX rows at EOF; independence statement in-body.

ACTIONS_GIT_REF: none — read-only dispatch authored; no code/tracked-doc edit by this relay (relay files live under gitignored `.relays/s6/`; the RECONCILE ledger entry for the boot-ack reconciliation is a tracked docs-only edit recorded in its own right).
FINAL_GIT_STATUS_SHORT: docs/sprints/2026-07-06-s6-slice-6/RECONCILE.md modified (the boot-ack reconciliation ledger entry, this session, uncommitted at relay-write time); no code paths touched.
Next requested action: operator hand-relays this dispatch to both pair seats; audits return under this DISPATCH_ID; reconciliation at my seat; DESIGN dispatch follows.
