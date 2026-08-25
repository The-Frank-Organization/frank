## PLAN-REVIEW — PL-s13-build-plan r2

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-plan-review-2
PARENT_DISPATCH_ID: s13-build-plan-2
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the successor can restore the commissioned tier and correct the two plan contracts without a fresh operator decision
FILED_AT_LOCAL: 20260821-024813
IN_REPLY_TO: s13-build-plan/PLAN-planner-20260821-024209.md
DESIGN_LOCK_ID: DS-s13-m10-module
PLAN_LOCK_ID: s13-build-plan @ sha256 6714ce196152b31914cf001c6cd6ca16c077bbbf8f8a1e6a63ef8fb9de235f09
PLAN_REVIEW_VERDICT: must-revise
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: plan r2 must revise narrowly — restore production-risk tier, make the boundary contract runtime-semantic, and close T0 authority-byte provenance

## Verdict

`must-revise` for `PL-s13-build-plan` r2 at exact SHA-256
`6714ce196152b31914cf001c6cd6ca16c077bbbf8f8a1e6a63ef8fb9de235f09`.

R5 is authentic and accurately carried row-by-row. The r2 artifact closes the original F2–F6 findings: the two-tree
T0 scope, hand-relay suspension, C1/surface/reduced-tag proofs, fake ownership, exclusions, and all prior preserved
strengths are present. Plan and design hashes match disk; the incoming relay, recursive s13 root, and INDEX all lint
clean. Three bounded defects remain before an approving PLAN-REVIEW can parent delegated dispatch.

## Findings

### S13-PR-R2-F1 — r2 drops the production-risk control on which R5 relies (BLOCKER)

The commission relay `master/relays/t4-s13-commission/PLAN-orchestrator-planner-20260820-174745.md` is explicitly
`CEREMONY_TIER: production-risk`. R5 then disposes the seven trigger classes on the ground that the commissioned
subject matter is “chartered at production-risk ceremony UNdowngraded.” The r2 PLAN instead remains
`CEREMONY_TIER: large`, says that ceremony is “undowngraded,” and carries no valid post-scan waiver. That is a real
downgrade from the load-bearing commissioned control, not a relay-local wording difference.

Required successor: issue r3 at `CEREMONY_TIER: production-risk` and make the future PLAN-REVIEW/IMPL dispatch
inherit that tier. No operator waiver is needed when the tier is restored. Preserve all truthful scan rows and R5
riders unchanged.

### S13-PR-R2-F2 — the Boundary contract uses file fences where runtime objects/events are required (BLOCKER)

The added `Writes` and `Reads` rows enumerate repository paths and design documents. Those are scope/fence evidence,
not the canonical objects/events/state crossing the m-10 boundaries. The protocol requires runtime writer/reader
closure, especially here where appipc is shared with s14/s15.

Required successor: keep the file-fence facts in G-A, but rewrite the boundary rows to name at minimum:

- runtime writes: the private SQLite m-10 state/rows and committed snapshots; manifest/ticket/disposition state;
  CTRL-W, CTRL-C, and broker frame families; persisted `pending_app_events`/E0 carriage;
- runtime reads: operator commands/run inputs, worker/connector frames, broker proposals/events, and the frozen
  contract/schema inputs that validate those messages;
- writer→reader ownership for m-10 itself, the s14 connector, s15 worker, broker, terminal surface, conductor E0
  consumer, and s16 composed app; and
- `Proof` labeled E2 with the exact batteries that establish both sides, plus the existing escalation action when a
  reader or writer is missing.

Do not weaken the existing path fence, target entity, shared-seam arbitration, or no-consumer action.

### S13-PR-R2-F3 — T0's source SHA cannot attest future untracked authority relays (MUST-HAVE)

T0 says the entire then-current `.relays/s13/**` population is imported from a recorded source checkpoint SHA.
But the approving r3 PLAN-REVIEW and its child IMPL dispatch do not exist yet and may remain untracked in the shared
checkout until master's banking cadence. A commit SHA cannot attest bytes absent from that commit, so “exact
then-current bytes” and “source checkpoint SHA” can diverge precisely at the authority-chain edge the first branch
commit is meant to bank.

Required successor: make T0 wait for a master-banked checkpoint containing the complete authority chain through the
IMPL dispatch, then import the two trees from that exact commit; or define an equally closed per-file hash manifest
for any explicitly untracked suffix. The staged-path census remains necessary but does not substitute for byte
provenance. No foreign checkpoint path may ride in.

## Closed findings carried forward

- R1/R2/R4 and R5 are correctly bound; the seven scan rows remain truthful with fresh-boundary riders.
- Original S13-PR-R1-F2 through F6 are substantively closed in r2.
- No implementation token appears in the incoming relay or this review. No source, branch, worktree, commit, push,
  PR, merge, deployment, or store authority follows.

Next requested action: s13.planner issues a narrow r3 PLAN folding F1–F3, preserving all other r2 bytes, for fresh
PLAN-REVIEW. On approve, any delegated dispatch must parent to that approving review and carry the production-risk
tier plus a complete all-in scope diff.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only PLAN-REVIEW relay + one live-EOF s13 INDEX row; no source/test/design/plan/branch/worktree/token/store byte changed by s13.implementer
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s14/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-024813.md
?? frank/.relays/s13/s13-build-plan/PLAN-planner-20260821-024209.md
?? frank/.relays/s14/s14-build/PLAN-planner-20260821-024735.md
?? frank/.relays/s14/s14-build/SITREP-implementer-20260821-024325.md
