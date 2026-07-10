## DESIGN-REVIEW - m-7 implementer approve on s6 transport amendments r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: m-7-s6-transport-amendments
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s6-design/DESIGN-planner-20260706-183120.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-2.planner
SUBJECT: approve - m-7 s6 transport amendments are pair-review clean; A-1 root-cause wording has one non-blocking code-cite sharpening
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: approve

I approve `master-docs/master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md` r1 as the m-7 pair design-review answer for `s6-design/DESIGN-planner-20260706-183120.md`.

This is a review-only approval. It does not lock the amendment, authorize code, authorize `frank/` writes, reopen Step-1 beyond the s6 slice, or bypass the master integration + VP co-sign gate. The m-1 parent-semantic fork still has its own operator-grill path; this review only confirms that m-7 can consume any branch without re-owning the m-1/m-2 policy surfaces.

## Blocking findings

None.

## Review checks

1. Scope and disposition are complete. The amendment accounts for every m-7-assigned finding from the orchestrator relay: F5 via A-1, F9 via A-2, F15 via A-3, F14 runtime half via A-4, F16 via D-1, and F3 via D-2. It does not silently grab F10/F17/F12 or the parent-meaning branch.

2. A-1 is lockable as a conductor-core contract. The intended surface is correct: digest is stable over volatile rendered affordances, while authoritative current-state validation still rejects genuinely stale or impossible submissions. The fixture class `N foreign accepts => zero form_digest re-render bounces` is the right regression target. Non-blocking sharpening: current code already marks the `parent_picker` field digest-exempt and strips its options/default from the digest path (`internal/fieldspec/render.go:65-73`, `:169-199`; `internal/fieldspec/render_test.go:90-119`). The final integration wording should not imply that parent-candidate contents are presently digest-covered in all cases. Say "volatile rendered affordances/dynamic fields" rather than "parent-candidate set currently rotates digest." The amendment's class-based rule and fixtures remain acceptable.

3. A-2 closes the actual intake duplicate gap. Current intake hashing dedupes to an existing intake id, but the writer path can still hand a duplicate job back for execution (`internal/intake/journal.go:70-83`; `internal/intake/writer.go:42-67`, `:110-117`). The amendment's replay/coalesce/no-new-execution outcome and durable monotonic counter semantics are the right conductor-owned fix.

4. A-3 correctly moves live seat minting into the serialized loop without expanding the seat tool surface. The current CLI mint path is admin-time and refuses while serving (`cmd/frank/main.go:405-415`). The amendment's `seat_mint` pivot, derived endpoint/binding recovery, and no-restart/no-fleet-bounce requirement fit the existing canonical-pivot model from the locked m-7 design (`master-docs/master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:80-83`, `:116-128`). m-2 row shape and m-1 credential fidelity remain downstream seam checks, not blockers on this m-7 review.

5. A-4 respects the F14 ownership split. The companion m-1 amendment owns invariant/staleness/takeover semantics; m-7 owns runtime choreography: claim before phase 0, loser refuses reads and writes, stale takeover follows m-1 policy, and takeover is auditable. That is the split required by the VP correction, and it does not re-home the invariant to m-7.

6. D-1 and D-2 are correctly left as in-slice build fixes. One transparent reconnect+reauth retry on conductor-loss and submit reply/detail parity for every recorded rejection class do not require architecture-of-record mutation beyond the s6 transport amendment.

7. The three-way parent seam is branch-composable. If m-1 Option A wins, m-7's responsibility is the conductor-maintained turn context plus commit-time stamp locus; the row/form shape runs through m-2. If another branch wins, m-7 still has a bounded integration seam and no reason to hold this review on the operator grill.

## Remaining constraints preserved

- The byte-exact submit outcome enum remains `{accepted,rejected,held}`; richer rejection detail is body metadata, not a fourth outcome.
- The seat-facing interface remains `submit()`, `project()`, and `read()` only.
- No Step-2 observe-as-send-gate or benchmark loop work is authorized by this amendment.
- The D5/same-uid residual and "policy runs here, m-7 does not own it" ceiling remain in force.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-design/DESIGN-planner-20260706-183120.md` - OK.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s6/s6-design .relays/s6/s6-design/DESIGN-planner-20260706-183120.md` - OK.
- `git -C frank status --short && git -C frank rev-parse --short HEAD && git -C frank tag --points-at HEAD` - clean; `7e5c527`; `s5-close`.
- Reviewed `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md`, `.relays/s6/s6-design/DESIGN-orchestrator-planner-20260706-180315.md`, `master-docs/master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md`, `master-docs/master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md`, and companion m-1/m-2 relays for seam consistency.
- Reviewed code evidence in `internal/fieldspec/render.go`, `internal/fieldspec/render_test.go`, `internal/intake/journal.go`, `internal/intake/writer.go`, `internal/engine/submit.go`, `internal/lineage/lineage.go`, `internal/store/projections.go`, `cmd/frank/main.go`, and `cmd/frank-mcp/mcp.go`.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-184027.md` - OK.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s6/s6-design .relays/s6/s6-design/DESIGN-planner-20260706-183120.md .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-184027.md` - OK.
- `rg -n "20260706-184027|DESIGN-REVIEW-implementer-20260706-184027" master-docs/master/relays/INDEX.md .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-184027.md` - index row present.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-184027.md`; appended `master-docs/master/relays/INDEX.md`; no code/source/`frank/` edits; no DESIGN_LOCK mutation; no PLAN or IMPL dispatch.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C frank status --short` clean at `7e5c527` / `s5-close`.
Next requested action: m-7.planner may carry this approval into the s6 integration package; master integration and VP co-sign remain required before any lock mutation or build dispatch.
