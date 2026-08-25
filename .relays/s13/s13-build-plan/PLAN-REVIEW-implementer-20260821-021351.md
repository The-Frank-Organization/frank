## PLAN-REVIEW — PL-s13-build-plan r1

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-plan-review
PARENT_DISPATCH_ID: s13-build-plan
RUN_ID: s13
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the blocking authority question routes first to master under the standing escalation ladder; no fresh operator decision is requested by this review
FILED_AT_LOCAL: 20260821-021351
IN_REPLY_TO: s13-build-plan/PLAN-planner-20260821-012742.md
DESIGN_LOCK_ID: DS-s13-m10-module
PLAN_LOCK_ID: s13-build-plan @ sha256 56563ee8914503f8fda32d061650d43e5ec7b47b93d2e3b569b7027f29d75127
PLAN_REVIEW_VERDICT: must-revise
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: plan r1 must revise — delegated-dispatch contradiction, stale store step, T0 scope conflict, and unmapped design obligations

## Verdict

`must-revise` for `PL-s13-build-plan` r1 at exact SHA-256
`56563ee8914503f8fda32d061650d43e5ec7b47b93d2e3b569b7027f29d75127`.

The plan correctly locks approved design `DS-s13-m10-module` r3 at
`3e74c4c1fe76dbaf07f3ca9547ffeb88a958cf23d0f6180afa9d49cd71a81361`; both hashes match disk, the parent is the
approving design review, exact PLAN relay lint passes, recursive s13 relay-root lint passes, and the s13 INDEX
passes its own index mode. The following blockers prevent an approving plan review and therefore prevent delegated
implementation dispatch.

## Findings

### S13-PR-R1-F1 — the delegated-dispatch predicate is false on the plan's own headers (BLOCKER)

The relay's `ESCALATION_SCAN` records seven `yes` rows and `ESCALATION_SCAN_RESULT: trigger-present`, while §2 says
the planner issues the token only on approval with “no scope deviation and no hard trigger.” The standing B4
deviation separately names a hard trigger as an escalate-to-master condition. An approving review cannot turn the
recorded triggers into “no hard trigger.”

Required successor: route this exact trigger set to `master.orchestrator-planner` and bind the returned disposition
before any token. The r2 PLAN must state a mechanically consistent dispatch predicate. If master rules that some
rows are already dispositioned by the commissioned baseline rather than fresh dispatch blockers, cite that ruling
row-by-row; do not silently rewrite factually true `yes` rows to `no`. No implementation dispatch may precede that
disposition.

### S13-PR-R1-F2 — T0's action and file set contradict each other (BLOCKER)

T0 promises a first commit containing “relays + INDEX + design + plan docs,” but its next line says
`Files: frank/.relays/s13/** only`. The design and plan now live under the R1-granted sprint tree, and current main
has already banked those bytes while the branch must still start at the older pinned LAUNCH_BASE.

Required successor: enumerate both in-fence populations in T0 and state the deterministic import/materialization
step that puts the exact then-current s13 relay bytes and sprint-doc bytes into the LAUNCH_BASE worktree without
bringing any foreign checkpoint paths. Record the source checkpoint/hash and verify the staged path population
before the first commit.

### S13-PR-R1-F3 — the plan is stale against the operative hand-relay suspension (BLOCKER)

T15 still assigns a “store-export at slice close.” The later directly addressed operating-model notice
`master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-020247.md` makes the store dormant for T4,
stops inbox polling/submission, and assigns master the export only through the suspension point.

Required successor: replace future store acts with file-first hand-relay steps, remove any pair-owned close-time
store export/poll/submit obligation during dormancy, and cite master's bounded suspension-point export ownership.
Historical store records remain citable; no new store twin is authored for r2.

### S13-PR-R1-F4 — the claimed full battery-to-task map has unmapped required legs (BLOCKER)

Against design §11:

- `FX-M10-C1` (one `uncertain` cut-relay entry; telemetry never input) has no task owner. T7 maps `FX-M10-C-1`,
  which is a different identifier; T9 maps C0/C2; T11 omits C1.
- The surface family requires the executable “no mutating verbs” leg; T13 lists payload-free views, non-zero exits,
  and persistent alerts but does not own that negative.
- The reduced-tag runtime overflow legs are named in T11, but neither T11 nor T15 names the command that actually
  selects and runs them. Ordinary `go test ./...` does not prove a build-tag-only battery.

Required successor: assign C1 and the surface negative to exact tasks, and name the reduced-tag test command at the
owning task and final battery (with the ordinary build/test/vet commands retained).

### S13-PR-R1-F5 — design §10's fake-counterpart deliverables have no construction owner (MUST-HAVE)

The design requires `fakeworker`, `fakeconnector`, and `fakebroker` speaking real frames over real socketpairs.
Multiple tasks consume those fakes, but no task explicitly builds/sequences their shared testutil foundations.

Required successor: assign each fake to a task after its appipc dependencies and before its first consuming
battery. The plan may distribute them across existing tasks; it need not add a ceremonial task.

### S13-PR-R1-F6 — boundary and exclusion contracts are not carried into the executable plan (MUST-HAVE)

This cross-component/shared-API plan has no explicit Writes/Reads/Target entity/Downstream consumer/Contract/Proof/
No-consumer-action block, and it does not carry design §12's exclusions. A generic frozen-spec sentence does not
make task scope auditable where in-fence code could still build toward deferred process split, multi-run scheduler,
wake ALO, forensic journal, ceiling policy, or conductor-side handshake work.

Required successor: add the boundary contract and an explicit out-of-scope list carrying design §12, including the
untouchable conductor/oracle surfaces and H-12. Keep the m-5 seam empty: enforcement socket only, no policy.

## Preserved strengths

- The design identity/hash, approved-review parent, LAUNCH_BASE, branch, target, and R1/R2/R4 rulings are carried
  accurately.
- The main T1–T14 dependency spine is coherent, and the retirement-cap, complementary limits selector, SQLite
  floor, restack/rerun/re-review, no-secret/no-payload, and operator-only merge constraints remain intact.
- The current relay contains no implementation token. This review creates none.

Next requested action: s13.planner files the F1 master escalation, then issues an r2 PLAN folding F2–F6 and the
returned F1 disposition for fresh PLAN-REVIEW. No source, test, branch, worktree, commit, push, PR, merge, deployment,
or store action is authorized.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only PLAN-REVIEW relay + one live-EOF s13 INDEX row; no source/test/design/plan/branch/worktree/token/store byte changed by s13.implementer
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s14/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-021351.md
?? frank/.relays/s14/s14-build/IMPL-planner-20260821-020810.md
