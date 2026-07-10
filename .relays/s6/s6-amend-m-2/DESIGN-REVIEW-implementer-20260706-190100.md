## DESIGN-REVIEW - m-2.implementer re-review of s6 transport-codec amendment rev1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-amend-m-2-review-r2
PARENT_DISPATCH_ID: s6-amend-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review only; master/VP/operator gates remain upstream
DESIGN_DOC_ID: s6-amend-m-2-transport-codec
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: .relays/s6/s6-amend-m-2/DESIGN-planner-20260706-190000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: approve - rev1 closes the three m-2 blockers; hold for master integration and VP co-sign

VERDICT: approve. I re-reviewed the revised DESIGN relay `.relays/s6/s6-amend-m-2/DESIGN-planner-20260706-190000.md` and the revised amendment doc `master-docs/master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md` for the same `DESIGN_DOC_ID`. The three blockers from `s6-amend-m-2-review-r1` are folded without widening the s6 boundary.

## Basis

- Revised design relay: `.relays/s6/s6-amend-m-2/DESIGN-planner-20260706-190000.md`.
- Revised m-2 amendment: `master-docs/master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md:50-62`, `:77-85`, `:111-116`, `:132-137`.
- Prior m-2 review blockers: `.relays/s6/s6-amend-m-2/DESIGN-REVIEW-implementer-20260706-184218.md`.
- S6 assignment: `.relays/s6/s6-design/DESIGN-orchestrator-planner-20260706-180315.md:25-44`.
- Live source anchors: `internal/fieldspec/registry.json:84,95-97,125,134`; `internal/engine/submit.go:47-64,143-178`; `internal/fieldspec/validate.go:54-61`; `internal/store/projections.go:110-141`; `internal/lineage/lineage.go:170-184,282-292,432-439`; `internal/fieldspec/render.go:55-87`; `internal/fieldspec/fieldspec.go:15-20`.
- S5 scope anchors: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md:23-28`; `.relays/s5/s5-b-merge-gate/RECONCILE-orchestrator-planner-20260706-153721.md:24-27`.
- m-1 F17 seam: `master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:44-52`.

## Re-review checks

1. **F13 record_kind authorization is now layered correctly.**

   The revised §2.6 distinguishes named enum membership, seat-scope/form authorization, and per-kind required-field validation. That matches the live order: `reg.Validate` runs before `validateRecordKind`, and `validate.go` enforces enum membership plus seat scope before lineage and per-kind checks. The revised fixture obligation now includes the required negative leg: `record_kind=genesis` remains known to the enum but rejected by seat scope because it is in no seat scope. This preserves the s5 lock while removing the second hardcoded membership judge that rejected authorized kinds like `disposition` and `diagnostics`.

2. **Recipient projection now uses the canonical header list everywhere it matters.**

   The revised §2.4 no longer treats `Envelope.To` as the recipient set. It makes decoded canonical Header `TO`/`CC` the source for rendered relay markdown, index display, delivery intents, and reviewer-visibility lineage. That closes the exact leak I flagged: current `projections.go` prints `Envelope.To` in render/index while delivery tries header recipients. The fixture in §7.1 now proves full multi-TO plus multi-CC preservation across all four surfaces and the no-silent-drop negative path.

3. **F12 is now bound to the m-1 F17 waiver record instead of an inexpressible header fallback.**

   The revised §4 drops the operator-scoped arbitrary free-text header branch. That is the right constraint: the current FieldSpec shape can seat-scope enum options, not arbitrary free-text header values. Binding the rationale to m-1's operator-only F17 waiver record preserves m-1 ownership of waiver class/scope/retraction while giving m-2 the row-typing work for `rationale`, `waiver_scope`, and `retracts`. The retired `"*"` header is no longer a competing carrier.

4. **The parent seam remains acceptable as a consumer contract.**

   §3 does not decide the m-1 fork. Under conductor-computed PARENT, the rendered field dissolves through the existing `renderable()` system/system_only guard; under a widened candidate branch, the render list and lineage acceptance list remain one source. That is enough for m-2 approval, with the fork still held for the separate master/operator grill.

## Held gates and boundaries

This approval is only the m-2 pair DESIGN-REVIEW approval for `s6-amend-m-2-transport-codec`. It does not fold c1, does not lock the m-1 parent fork, does not approve m-7 runtime mechanics, and does not authorize code. The design remains held for master integration of the three amendment docs plus seam text, VP co-sign, and the later s6 build-slice dispatch.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended the matching `master-docs/master/relays/INDEX.md` row; reviewed docs/code only; no `frank/` code changes; no c1 fold; no PLAN or IMPL.
FINAL_GIT_STATUS_SHORT: docs workspace root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git` / exit 128; `git -C frank status --short` empty / exit 0.
RELAY_LINT: OK - relay-lint.py exit 0 on `.relays/s6/s6-amend-m-2/DESIGN-REVIEW-implementer-20260706-190100.md`
