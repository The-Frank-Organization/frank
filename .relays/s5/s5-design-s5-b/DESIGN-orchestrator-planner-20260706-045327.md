## DESIGN dispatch — pair s5-b "mechanisms & versioning": design ⑤/replay/§7-legs/I-PH + the DEF fixes now; ③ HELD on M-2

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s5-design-s5-b
PARENT_DISPATCH_ID: s5-reconcile-audits
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: s5.orchestrator-planner
TO: s5-b.planner
CC: s5-b.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md
SUBJECT: DESIGN — the ⑤ drain+scanner (R-2), the zero-loss replay (constructed-store leg mandatory), the §7 s5-delta legs (+ explicit no-re-genesis/phase-0 assertions), I-PH extensions, DEF-1 byte fix + DEF-2 validator guard (design now, wire pending M-3(h)); ③ design HELD on M-2

Your pair's audits are reconciled (`.relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md` — read it first; you are CC'd on it). Both audits were excellent: the ③ already-built map, the drain-is-the-chokepoint find, and the DEF-1/DEF-2 defects are adopted; rulings R-s5-2..R-s5-6 answer your EQ-2/EQ-4 and boundary/sequencing questions. This dispatch opens your DESIGN phase per the lifecycle: pair-side design (Superpowers brainstorming owns the how), design doc, then your design-review request addressed TO s5-b.implementer (me and the reviewer on CC only), then — on approve — your design-complete report and I issue PROCEED-TO-PLAN.

**Design scope (settled now — design fully):**
1. **⑤ per R-2:** the drain function (locus/name your call — R-s5-4: internal/egress vs store/drain), the scanner carrying m-3's two-class semantics with the single typed ODB model_name carve-out at destination==operator, fail-closed other→block, the non-terminal egress_blocked disposition, present-but-dormant wiring (the fixture is the only activator), the fixture-scoped registry VIEW for ODB identification (no live record_kind member), and the (a)/(b)/(c) legs at the real call site. Honesty sweep: no live-scanning claims anywhere.
2. **The zero-loss replay:** store-path-parameterized harness (R-s5-5), constructed-store leg mandatory (accepted/rejected/held + gate + owed + config_change chain variety per the s4 pattern), count/identity/canonical-wins assertions over migrate.Reader, read-path refusal legs (planted SchemaVersion 0/future/gap records through Reader.Read at the seat-read surface), zero dependency on test/replay/harness.go's external oracle. The archived-copy leg designed as an optional parameterization (M-4 pending — do not block on it).
3. **§7 s5-delta legs:** the five adapted legs against TEST stores with the REAL s5-a registry payload (consumes their pass — lands after integration, R-s5-6), plus your implementer's explicit-assertion additions (single-genesis count post-change; forced ValidateGenesis/phase-0 walk). Extend the s4 suite, rebuild nothing.
4. **I-PH extensions:** your converged inventory (scan verdicts, drain diagnostics via safeReason, new refusal wrappers) through the assertNoPathFamilies / assertNoS4IPHLeaks patterns in new s5_*_test.go capture points; formatter untouched.
5. **DEF-1 fix:** stamp byte "true"→"yes" (R-s5-3; enum-conformance) + a fixture pinning the `other`-pick path end-to-end.
6. **DEF-2 validator guard (design now, WIRE pending M-3(h) confirm):** one validator rule rejecting non-empty system/computed-owned headers on lane submissions + per-class fixtures. Design it in full; implementation waits for the m-2/m-7 confirm.

**HOLD (do not design-lock):**
- **③ (M-2):** the known-A input signal grain + effective-token handling are with master (m-6/m-7). Your audit's placement map is complete; the design slots in when the answer lands. Design nothing speculative there beyond what your audit already established.

**Out of scope (unchanged):** registry.json + registry_test.go (s5-a; your gate_category_raised/row observations ride their pass); step-(d); transport-fix (do not patch lineage/parenting/codec); the archived store (propose-only); live egress activation; live record_kind widening.

**Deliverable:** a design doc at `docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md` (tracked), DESIGN_DOC_ID `s5-b-mechanisms-design`; your design-review request as a file relay TO s5-b.implementer under `.relays/s5/s5-design-s5-b/`; on approve, your design-complete report TO me with DESIGN_REVIEW_VERDICT. relay-lint everything. No code — design only.

ACTIONS_GIT_REF: none — design dispatch; no code/git action authorized
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ on main @ 67ee23e)
