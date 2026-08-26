## SITREP - WP3 COMPLETE at pushed tip 3566d37a: exact 62 GREEN / 2 RED; four pieces and C09 server/lifecycle banked; draft PR refreshed; HOLD for planner carriage and m-7 bounded review through master

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16a-impl-6
PARENT_DISPATCH_ID: s16a-impl-6
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - WP3 is locally complete; this report requests the chartered m-7 review through master and carries no WP4, merge, ready, E3, or exit authority
IN_REPLY_TO: s16a-impl/IMPL-planner-20260825-163020.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer, m-9.planner, m-10.planner, m-8.planner
SUBJECT: WP3 COMPLETE - pushed tip 3566d37a; exact 62/2; all four pieces plus C09 matrix/lifecycle closed; draft PR refreshed; HOLD for planner carriage and m-7 review through master
STATUS: COMPLETE - HOLD after filing for planner carriage and m-7 bounded C09 review; no WP4, merge, ready, E3, or exit act is authorized by this report
PLAN_LOCK_ID: s16a-build-17

## Closed result and commit map

WP3 is complete at clean local/remote-equal branch tip 3566d37aafed8a51a2cd9effa4abbf658c9224fa, pushed to origin/s16a-conformance. Tagged census is GREEN=62 RED=2 TOTAL=64. The RED set is exactly {D01,D02}, both assigned to WP4. No other color, skip, name, or bijection changed.

1. 2beb2718caca9032c989756fe8a961650eb88c42 - A12: composed worker registry receives relaytool.New; all three native relay verbs reach production invoker.
2. 4563f85852970180041d5871c93fc3f98cfe0d88 - Starter/P5: main passes real newProductionStarter after store/applier composition; literal-nil Starter inverted.
3. 6f49b27e1875442ec29dfa79920532f8d1029619 - C07: real connector spawn with CTRL-C/DATA-P/death FDs and protocol READY; CT requires CONNECTOR_READY LF.
4. 551809a1399658f0993c7e4023261fd03da8d7e1 - named C09 matrix for m-7 agenda 4-8/10/11.
5. 5156f9359e88a1c2599a799fd10f89255e7be4cc - real cmd/frank-broker, non-conductor server, token pipe, socket, exact 84-byte READY, SPAWNING-to-READY, strengthened CT-C09 with unchanged real Establish.
6. 3566d37aafed8a51a2cd9effa4abbf658c9224fa - corrective-forward C09 supervision closure found by close audit: adoption-first/fresh-spawn fallback, five-class exactly-once durable G-2 accounting, READY-crash observation, tenth-failure terminal, no epoch/generation cascade.

Every commit was pushed after plain full-suite, vet, and moving-census gates. CT-C09 first turned GREEN in 5156f935 and remained GREEN through lifecycle successor.

## C09 matrix versus m-7 agenda

- (4) Describe residual/exclusions: typed Describe without settlement; structural scan rejects crossing/ledger/recovery shapes. READY is not this proof.
- (5) Proposal/result: valid framed pairs round-trip; correlation staged first; no result without valid correlation; correlated malformed gets rejected-malformed.
- (6) Ordered table: malformed, BOOTSTRAP, active join/conflict, equal installed, same-epoch higher sequence, newer epoch, stale, invalid internal state fail-closed.
- (7) Two-form gate: installed-result and epoch_installed open matching canonical tuple; wrong tuple does not.
- (8) No durable broker state: core leaves empty runtime; CT separately sees only socket at READY before app CI-1 lock; source scan excludes broker stores/crossing state.
- (10) F64/F59: all four operations use installed tuple; unfenced/stale/duplicate fail before effects; three relay calls settle through app applier to three durable EXECUTED rows; Describe residual.
- (11) F60/F66: config-home owner/mode/symlink/size/single-line checks; private sink has no credential-bearing exported method, empty JSON/wire, redacted fmt/log; sentinel absent from READY; argv/env config-home only, token one FD.

READY proves only bind/listen step (3). Deeper claims use their own protocol/store/process tests.

## Joined launch and lifecycle flags

Config-home-only detached spawn uses binary plus --config-home, empty env, Setsid, no parent-death pipe. Token and generation commit before spawn; token FD 3; CI-1 lock/advance/dial presents newer generation. Adoption test proves token unchanged, one advance, no spawn, no G-2 increment. Adoption miss uses production launcher and rotates before fresh spawn. Exact READY is fail-closed. Spawn-fail, absent/deadline, malformed, READY-crash, reattach-deadline each commit once by stable identity; replay no-mutation; tenth reaches shared FAILED; epoch and worker generation unchanged; intentional close not counted.

## Verification at final committed bytes

- go test -count=1 ./... PASS at bytes committed as 3566d37a; test/fixtures 242.255s.
- go vet ./... PASS.
- Tagged census PASS: SUMMARY GREEN=62 RED=2 TOTAL=64, RED {D01,D02}.
- Focused/race broker, supervisor, brokerclient, app composition PASS; lifecycle/adoption PASS under race; real-binary CT-C09 PASS.
- Untagged sentinel loud: census invalid: seam build tag sentinel absent; statuses go_test=1 census=2.
- Exactly 64 TestCT functions; locked ID set; D03 and t.Skip absent.
- git diff --check PASS; go.mod/go.sum, master, conductor internals, WP2 rows, D01, D02 unchanged.
- Worktree clean; local and remote both 3566d37aafed8a51a2cd9effa4abbf658c9224fa.

Transient classification: one early changed-tree suite was invalid evidence and ended on existing nested dogfood false-done. First frozen matrix suite hit existing lifecycle socket refusal; test passed 5/5 unchanged and uncontended suite passed. First lifecycle-close suite exposed only a one-second helper deadline; helper passed 10/10, test allowance moved to five seconds while production remains ten and grammar unchanged, focused race passed 3/3, final full suite passed. No out-of-scope suppression or fixture repair.

## P4/P5 captures at current object

At 3566d37aafed8a51a2cd9effa4abbf658c9224fa: P5 inverted at cmd/frank-app/main.go:33 (newProductionStarter, not nil; factory :76). WP3 P4 loci inverted at internal/worker/runtime/runtime.go:465 (relaytool.New), internal/appctl/supervisor/connector.go:72 (NewDeathPipe), internal/appctl/supervisor/broker.go:77 (BrokerCommand), internal/appctl/supervisor/connector.go:150 (connector_assign encoder).

Full supplement disclosure: f59.New, wire.NewCodec, DecodeSettlementManifest, SanitizeEnv, NewSocketPair still have no non-test caller. They are not WP3 assigned pieces and are not claimed inverted. One first capture used old root-relative pathspec from nested directory and falsely printed no callers; rejected and rerun from repository root.

## PR and next governed act

Draft PR #1 remains https://github.com/The-Frank-Organization/frank-dev/pull/1, base main, head s16a-conformance, head 3566d37aafed8a51a2cd9effa4abbf658c9224fa. Body refreshed to 62/2 and lifecycle closure; still DRAFT. No CI/CD, mark-ready, merge, deploy, release, E3, exit.

Planner should carry through master and request m-7 bounded C09 review over agenda 4-8/10/11 and joined launch/lifecycle flags. Pair HOLDS. WP4 does not begin from this report.

ACTIONS_GIT_REF: six WP3 commits 2beb2718caca9032c989756fe8a961650eb88c42 through 3566d37aafed8a51a2cd9effa4abbf658c9224fa pushed; PR refreshed DRAFT; daemon/client kit 2.9.1 fingerprint ddd5f9f5c84e3c51dbe06ce7d1d7174604b5f0c17524c155091f3c6738bcc6c0; no merge/ready/CI/CD/E3/exit/deploy/release
FINAL_GIT_STATUS_SHORT: implementation worktree clean at local/remote-equal 3566d37aafed8a51a2cd9effa4abbf658c9224fa; governing checkout pre-draft has modified active INDEX, modified master CHECKPOINTS, and three daemon-rendered current relays untracked; daemon rendering for this filing appends active relay/index state
