## DESIGN-REVIEW -- m-1.implementer adversarial review of c6-fix-m-1 planner fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-1-review-r1
PARENT_DISPATCH_ID: c6-fix-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- adversarial review of doc-only c6 cleanup; no operator decision surfaced
GRILL_REQUIRED: no -- no new mechanism or product decision
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
BUNDLE_ID: c6-fix-m-1
OWNER: m-1 (Trust & Identity)
IN_REPLY_TO: c6-fix-m-1/DESIGN-planner-20260702-210000.md

m-1.planner -- I reviewed the c6-fix fold against the orchestrator dispatch, the c6 re-review source of record, the current m-1 design doc, and the m-7/ARCHITECTURE disposition anchors. Verdict: **must-revise**.

Most of the fold is on target: the fold-log exists, F12's `held` mapping is surfaced, F4's acceptance/carry-forward tests are re-cut to the D4 tool-surface grain, F6's Step-2 observe hook annotation landed, the byte-exact `{accepted, rejected, held}` enum is present, and the observer-selected sanctioned claim remains intact. Two stale sentences still undercut the requested c6 cleanup.

## Must Revise

1. **m-1-F3 is not fully cleared: §4 I1 still presents the D3 wrap mechanism as the current "Mechanism."**

The c6 source says DI-1 was still demanding the shelved wrap via "no lane write access / separate uid / E2: a direct lane write fails" and resolves it as "DI-1 = D3-shelved wrap invariant; Step-1 I1 rides the m-7 guardrail; re-cut AC#1 to the D4 fixture" (`master/DESIGN-REREVIEW-2026-07-02.md:34`, `:72`). The planner fold correctly fixes the DI-1 bullet, Step-1 row, proof, and AC#1 (`m-1 design:87`, `:96`, `:161`, `:208`). But the I1 subsection still says, unqualified, `Mechanism: the store lives on a path owned by the conductor with no write permission for lane processes (separate uid / 0700 / no path handle)` (`m-1 design:95`).

That sentence is the exact wrap-grade realization c6 is trying to shelve. Because it is labeled "Mechanism" inside the load-bearing I1 proof section, it can still be read as the Step-1 mechanism rather than as a D3 future criterion. Revise that sentence so the Step-1 mechanism is the m-7 interface guardrail / absent store path in seat tool surfaces, and move the separate-uid/store-ACL/no-path-handle wording into an explicitly D3-shelved wrap clause.

2. **m-1-F5 is incomplete in §10: the rejected-alternative rationale still uses the old by-construction contrast.**

The dispatch asks for the transport rationale to be rebuilt in `§3/§10`: A beats B because minted credentials are the non-re-cut foundation for D3 wrap / SO_PEERCRED/mTLS/SPIFFE, not because A earns a stronger Step-1 by-construction claim (`c6-fix-m-1/DESIGN-orchestrator-planner-20260702-204500.md:29`). The §3 paragraph now states this correctly (`m-1 design:80`). But §10 still says `Option B (operator-attested) -- only confusion-resistant, not by-construction; forfeits the v3 differentiator` (`m-1 design:186`).

That is the stale rationale F5 was meant to remove. Under the c5/c6 sweep, both A and B are only Step-1 confusion-resistant; the differentiator is A's upgrade-bearing non-re-cut substrate. Revise the §10 rejected-alternative bullet to say B is rejected because it is not upgrade-bearing / would require a re-cut to reach structural strength, while A is the non-re-cut minted-credential foundation.

## Checks Passed

- **F12 (`held`) is correctly scoped.** The §5 `submit()` atomic commit paragraph now distinguishes form/lineage failure from trusted machinery ran-and-broke / check-could-not-run, maps authority-bearing records to `held`, non-authority records to `rejected` or author-returned, and says the byte-exact persisted-outcome set is `{accepted, rejected, held}` (`m-1 design:126`). This matches the VP-amended two-axis wording (`DESIGN-REREVIEW-2026-07-02.md:68`) and m-7's C4.1/NF-S6 anchors (`ARCHITECTURE.md:429-431`; `m-7 design:100`, `:156`, `:166`).
- **F4 is substantially corrected.** §12 #3, §13 #1, and §13.rev2 #1 now use the D4 tool-surface negative fixture and explicitly label the wrap-grade sibling-read probe as D3-shelved (`m-1 design:210`, `:222`, `:247-251`).
- **F6 landed.** The m-3 consumer line and §13.2 now mark observe as `[Step-2+, reserved hook 2c]`, not Step-1 (`m-1 design:149`, `:223`).
- **Preservation checks are mostly clean.** Exact enum grep returns 3 hits for `{accepted, rejected, held}`; `unspoofable` is absent from the m-1 design; `observer-selected` remains present on DI-5 and `parent_picker` surfaces (9 hits).

No mechanism/verb/field/on-disk change is requested by this review. This is a narrow doc consistency revision only; no c1 reopen, no PLAN, no IMPL, no `pcode/`, no spike.

## Verification

- `sed -n '1,360p' master/relays/c6-fix-m-1/DESIGN-planner-20260702-210000.md` -- reviewed full planner fold relay.
- `sed -n '1,320p' master/relays/c6-fix-m-1/DESIGN-orchestrator-planner-20260702-204500.md` -- reviewed dispatch and per-finding scope.
- `nl -ba master/DESIGN-REREVIEW-2026-07-02.md | sed -n '1,140p'` and `sed -n '300,430p'` -- reviewed c6 source-of-record clusters and canonical resolutions.
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md | sed -n '1,290p'` plus targeted `sed -n '112,132p;206,258p'` -- reviewed folded m-1 design text.
- `rg -n -C 3 'Option B|forfeits|A . B|operator-attested|store lives on a path|no write permission|separate uid|0700|no path handle' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md master/DESIGN-REREVIEW-2026-07-02.md master/relays/c6-fix-m-1/DESIGN-orchestrator-planner-20260702-204500.md` -- found the two blockers above.
- `grep -oE '\{accepted, rejected, held\}' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md | wc -l` -- 3.
- `rg -n 'observer-selected' master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md` -- 9 hits.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-1/DESIGN-planner-20260702-210000.md` -- OK.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; no design-doc edit, no code/source/`pcode` edit, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ checked clean)
RELAY_LINT: OK -- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-1/DESIGN-REVIEW-implementer-20260702-210016.md`
DISPATCH_ROOT_LINT: OK -- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-1`
