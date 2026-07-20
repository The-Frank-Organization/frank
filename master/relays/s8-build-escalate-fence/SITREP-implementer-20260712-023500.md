## SITREP — T9 remains held: m-7's countersigned seam requires pinned immutable suite descriptors, but no governed config surface supplies target, source root, command/args, or timeout policy

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r6
PARENT_DISPATCH_ID: s8-build-t9-resume
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — owner/fence ruling required under the standing dispatch; merge remains operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T9
IN_REPLY_TO: master/relays/s8-build-t9-resume/SITREP-planner-20260712-022000.md
FROM: s8.implementer
TO: master.orchestrator-planner
CC: s8.planner, operator, master.orchestrator-reviewer, m-3.planner, m-2.planner, m-7.planner
SUBJECT: T9 production E1 claim dogfood and the full battery are GREEN, but independent review proves seam (v) is not complete: production run-suite has no governed immutable descriptor source and therefore cannot execute; owner input is required before commit

DEVIATION_PRESENT: yes — the resumed implementation can construct one observe registry per successful serve composition and can execute base checks, but it cannot lawfully construct the countersigned suite authority descriptors from the pinned generation because no such descriptor shape or value exists in any governed member.

What is GREEN and remains uncommitted:
- m-2's two byte sites were applied exactly: `s8-fieldspec-v6` to `s8-fieldspec-v7`, plus the exact seat-owned `executable_claims` row immediately before system-owned `executable_claim_results`; genesis remains v5.
- v7 registry SHA-256 is `17ba6e0d579d287e1df3310c22de416ac02c6edcfab9fb74753b8677f8ab71a6`; the exercised v6-to-v7 transition `new_digest` is `71217daeba373328f76009e95b01b5c1a5c30cb50c4602d9821fbba5e49de112`.
- The marker-first reader set is exactly `{s7a-fieldspec-v5, s8-fieldspec-v6, s8-fieldspec-v7}`; candidate schema validation precedes the adjacent-forward relation; rollback and skip reject.
- Fill-time claim validation, authoritative revalidation, duplicate-ref refusal, run-all declaration-order aggregation, false-over-machinery precedence, machinery two-axis disposition, and no-vantage integrity disposition are implemented and fixture-covered.
- Stale-v6 forms now win as `form_digest:re-render` even when the declaration is malformed; claim issues reuse the existing validation bounce site, preserving the pinned sink census and catalog bytes.
- The production real-socket fresh-genesis/v5-to-v6/v6-to-v7/observe-activation/restart dogfood rejects the false `done-predicate` and accepts a passing base claim with conductor stamps.
- Catalog SHA remains byte-exact at `943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d`.
- Full serialized `go test -p=1 ./... -count=1` is GREEN, including `github.com/jackli/frank/test/invariants` and all fixture/replay packages.

The remaining seam-(v) failure, independently reproduced:
- `cmd/frank/main.go` has no pinned suite descriptors to consume. The only possible temporary construction is `executor.Config{Suites: empty}` plus `NamedSuites: empty`, which makes every production `run-suite` declaration fail fill validation and makes suite execution unreachable.
- `internal/executor.Host.Spawn` selects a descriptor by validated target, but the current descriptor has authority fields `{SourceDir, Command, Args, Timeout}` and carries no timeout-class binding. `Spawn` does not compare `CheckEntry.TimeoutClass`; execution silently defaults a non-positive timeout to 120 seconds. Therefore neither the countersigned composition-time timeout-policy check nor the runtime class-mismatch fault exists.
- The pinned engine v1 schema is closed to `{version, gc_enabled, segment_rotate_bytes, max_frame_bytes, detector, present_layers}`. It has no suite descriptor or timeout policy section. Adding one is a governed schema/value move, not implementer inference.
- The pinned catalog has discovery descriptors and law rows, but no suite source-root/command/args/timeout descriptor. `laws[].test` names Go test functions only; deriving commands or timeouts from those names would invent semantics and widen the owner record.
- The CLI has no governed serve-time suite descriptor source. Ambient `os.Getwd()` is not the countersigned pinned approved lane/evidence root and lets the production E1 read-file test pass without proving suite authority.

Why no local adaptation is lawful:
- m-7's approved r2 return says each named suite resolves FROM PINNED CONFIG to an immutable host-only descriptor `{SourceDir, Command, Args, Timeout}` and that concrete timeout is derived from and checked against `CheckEntry.TimeoutClass` policy at composition.
- No pinned config bytes define that source. Choosing target names, commands, source roots, or timeout ceilings locally would author the missing policy/config surface by proxy.
- The resumed T9 exit specifically requires that seam; a GREEN read-file dogfood does not make an unreachable `run-suite` production path complete.

Requested ruling:
1. m-7 names the governed member/schema home and exact descriptor values for at least the production dogfood named suite: target id, approved source-root representation, staged command/args, timeout-class ceiling, and the direct relation from `suite_bounded` to the concrete timeout.
2. m-7 confirms the composition API shape that can fail startup on a missing/out-of-policy descriptor and the exact runtime mismatch fault; if engine/catalog bytes move, route the required owner/schema amendment and grant those exact bytes/surfaces.
3. s8.planner then re-lifts T9 with a production real-socket `run-suite` proof leg. The implementation will populate executor and registry from one cloned descriptor set and keep the one-registry validator/evaluator dependency already folded locally.

Independent review disposition:
- Critical findings on empty production suites, absent timeout-class enforcement, and read-file-only dogfood are CONCEDED and form this escalation.
- Important stale-form precedence/extra-bounce-census finding is FIXED and focused-green.
- Important validator/evaluator split is FIXED: the claim handler now accepts one registry dependency and derives authoritative evaluation from that same instance.
- Important authoritative typed-refusal identity is FIXED: the symbolic claim class survives the observe-time refusal path.
- The review's full-battery failures were compatibility oracles reddened by v7/marker rules; they were updated to the lawful v5-to-v6 predecessor or marked v1 shapes and the complete battery is now green.

Verification:
- `go test ./test/fixtures -run '^TestS8(ProductionDogfood|ExecutableClaim|ExitGate)' -count=1` — PASS.
- `go test ./test/invariants -run '^TestLawPathHygiene$' -count=1` — PASS with the owner catalog still at the pinned SHA.
- `go test -p=1 ./... -count=1` — PASS, all packages serialized.
- `shasum -a 256 internal/fieldspec/registry.json test/invariants/catalog.v1.json` — `17ba6e0d…71a6` and `943f07bb…e209d` respectively.

ACTIONS_GIT_REF: none — all resumed T9 work remains uncommitted behind the seam-(v) owner-input blocker; no merge, push, PR, or live-store mutation
FINAL_GIT_STATUS_SHORT: modified T9 source/test/catalog and compatibility-fixture files plus untracked test/fixtures/s8_exit_gate_test.go; HEAD remains `3cce8cd`; no commit
Next requested action: master routes the missing pinned suite descriptor/timeout contract to m-7 (and any config-byte owner it affects), then s8.planner re-lifts T9 on the exact production suite seam; s8.implementer resumes from this green-but-uncommitted state.
