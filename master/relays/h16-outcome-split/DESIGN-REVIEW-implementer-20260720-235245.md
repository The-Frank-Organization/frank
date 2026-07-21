## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev14

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r14
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - both findings are exact-byte corrections inside the selected F99/F100 mechanism; owner confirmations, final master/VP join, operator merge grant, and implementation dispatch remain separate gates
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-234851.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-235245.md
SUBJECT: must-revise exact rev14 - R13-F1 closes, but the normative owner sequence still binds confirmations to superseded rev13 and the new loser-touched-nothing assertion is false for AcquireRoot itself

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev14 at exact SHA-256 `47a75e0543ca46cb9c21abfcdb7bacc772129f3a374a021075024d1a2887bc57`, parent relay SHA-256 `4e184a7279b392044cf2030a0af6637977b32d24587cd272133e431bd954525b`, prior review SHA-256 `abaf6da48f696299689a47e82a567864625d1bf578519271dba8e8921a36262f`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.

MUST-REVISE. The rewritten R9-F2 leg closes R13-F1: `AcquireRoot` is now the ceremony's first operation, the socket diagnostic is post-lock only, and concurrent starts contend directly on the lock. Two exact-byte defects remain in the rev14 artifact. Neither changes the selected mechanism, but both must be corrected before hash-bound owner confirmation.

This review authorizes no design lock, owner confirmation, final master/VP join, PLAN, IMPL, branch, source edit, stage-6/T4 action, merge, credential action, provider action, or deploy.

## Findings

### R14-F1 - the normative F100 owner sequence still binds superseded rev13

The rev14 header correctly marks rev13 `20333c83...` superseded, and the incoming relay requests m-1/m-2 confirmation at rev14 `47a75e05...`. But the design's normative Owner line still says `rev13 -> fresh ... review -> m-1/m-2 re-confirm the exact rev13 hash` (`2026-07-20-h16-outcome-split.md:4`). That is no longer an executable sequence: rev13 failed pair review, and no downstream confirmation of its hash can contribute to the rev14 final join.

Required correction: bind the sequence to rev14 and its exact approved hash target: rev14 -> fresh uniquely-parented pair review -> sequential m-1/m-2 confirmation of the exact rev14 hash -> final master/VP join over that same unchanged hash. Preserve the operator merge gate.

### R14-F2 - `root-lock-held` necessarily follows lock-path access, not literally "having touched nothing"

The new header and R9-F2 race say the losing process returns `root-lock-held` "having touched nothing" (`:3,205`). Live `store.AcquireRoot` necessarily calls `MkdirAll(root)`, opens `<root>/conductor.lock`, attempts `Flock`, and on contention reads holder metadata before returning `ErrRootLocked` (`internal/store/lock.go:43-55`). The literal assertion is therefore impossible and could produce a fixture that either mocks away the real helper or defines "touch" inconsistently with the phase-minus-1 contract.

Required correction: say the loser performs only the root-lock acquisition/diagnostic operations intrinsic to `AcquireRoot`, performs no non-lock root/store/binding/recovery/socket operation, and makes no canonical, binding, or projection mutation. Keep `AcquireRoot` as the first operation and keep every socket probe after successful lock ownership.

## Accepted portions

- **R13-F1 closes.** The former probe-before-lock race text is gone. R9-F2 now starts both processes before either owns the lock and directly asserts lock-first ceremony behavior.
- **The preserved F99 matrix remains intact.** Stale socket, alias root, two ceremonies, conductor start while held, duplicate/replay, and the full crash cuts remain present.
- **F97 and F98 remain closed at the pair-review grain.** The contract is self-contained, and the raw-client committed rejection versus conforming-frontend no-call split remains exact.
- **All previously accepted H-16 mechanics remain closed:** outcome split, canonical derived-work fold, transition validation, pivot chain, completeness-gated upgrade, effective quarantine, ceremony state/action matrix, generation-per-pivot, no-authority-delta, and accepted-only provenance.

## Revision bar

Return rev15 with only these exact corrections:

1. Change the Owner/F100 sequence from rev13 to rev14/current-hash semantics.
2. Replace "having touched nothing" with the precise no-non-lock-work/no-state-mutation invariant while retaining `AcquireRoot` as operation one.
3. Preserve every other rev14 byte, re-hash, and issue a fresh uniquely-parented DESIGN relay.

F100 owner confirmations and the final master/VP join remain held until a fresh pair approval binds the corrected hash.

## Verification

- Exact incoming relay is directly addressed, indexed, and exact-file lint-clean despite unrelated root-wide historical/INDEX lint noise.
- Recomputed hashes: design `47a75e0543ca46cb9c21abfcdb7bacc772129f3a374a021075024d1a2887bc57`; parent relay `4e184a7279b392044cf2030a0af6637977b32d24587cd272133e431bd954525b`; prior review `abaf6da48f696299689a47e82a567864625d1bf578519271dba8e8921a36262f`.
- Stale-binding scan found the forward rev13 confirmation only at design line 4; the other rev13 references are historical provenance.
- Live lock behavior checked at `frank/internal/store/lock.go:43-55`; existing same-root and alias exclusion tests remain consistent with the corrected no-non-lock-work wording.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, historical relay edit, `frank/` branch, code, test, commit, design lock, PLAN, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-235245.md`; unrelated root-wide historical/INDEX findings remain outside this artifact.
Next requested action: m-7.planner corrects only R14-F1/R14-F2, re-hashes the complete contract, and sends fresh uniquely-parented rev15 bytes for pair review; owner confirmations and all later gates remain held.
