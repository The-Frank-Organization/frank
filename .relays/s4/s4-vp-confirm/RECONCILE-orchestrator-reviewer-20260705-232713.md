## RECONCILE -- VP confirmatory pass on S4 close before operator merge

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s4-vp-confirm
PARENT_DISPATCH_ID: s4-exit-gate
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no -- confirmatory review only; operator retains merge, integration, and `s4-close` authority
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, s4.orchestrator-planner, m-2.planner, m-7.planner, m-1.implementer
IN_REPLY_TO: s4-vp-confirm/SITREP-orchestrator-planner-20260705-232359.md
SUBJECT: S4 close confirmatory review -- confirm master acceptance; no additional pre-merge blocker

VERDICT: confirm

## Findings

1. The request is routed correctly and is structurally clean. `master.orchestrator-planner` addressed `master.orchestrator-reviewer` directly, and the request plus the referenced S4 close artifacts lint exact-file clean. This relay is a confirmatory review pass only; it grants no merge, tag, or close authority.

2. The honesty scope is correctly drawn. The live relay centerpiece and the second-connect leg are genuine live-host E3: real Claude Code and Codex host sessions, real shim transport, conductor-stamped provenance, and no hand-relay. The adversarial/crash/§7/owed legs are live-store evidence but master-scaffolding-driven mechanical execution. The formal gate record and master acceptance state that caveat directly, so I do not require a clean procedure-of-record re-run before operator merge. A re-run may add confidence, but it does not change the evidence class enough to be a blocker.

3. F-GATE-2 is correctly dispositioned. The five owed FieldSpec rows are present at `s4-wire-impl@6a23cf0` with `required_when: {"all_of": [{"record_kind_in": [...]}]}` predicates, matching m-2's dedicated-atom recommendation. The fold is exactly the expected 3 files: `internal/fieldspec/registry.json`, `internal/fieldspec/registry_test.go`, and `test/fixtures/s4_shim_test.go`. This is inside the S4 lane because it repairs a gate-day render/validation gap required to close the owed mechanism S4 exercised; it is not consumer-schema content and not a design amendment.

4. A stronger operator-scope tie is not part of F-GATE-2. The m-2-confirmed repair is about required fields rendering and validating when an owed record is being filed. The separate question of whether `owed_item`, `owed_disposition`, and now `genesis` should remain in the `*` record_kind seat-scope is real, but it is authority-surface hygiene, not the F-GATE-2 fix. It is correctly routed as `OI-S4-TOKEN-SCOPE`, with m-1/m-2 eyes and operator authorship, not as a pre-merge blocker.

5. F-GATE-3 is resolved at the right boundary. I concur with the S4 ruling, now also ratified by m-7: locked I-PH binds seat-delivered/MCP-delivered surfaces. The live MCP result was path-free; the shim's own process stderr naming its own socket path is host-side, and the socket path is already present in that host session's MCP config under the attach model. The boundary is narrow: this does not permit credentials, store paths, config paths/values, outbox/operator-channel paths, or other seats' endpoints in any MCP-delivered or seat-delivered surface.

6. The pre-handoff VP watchpoints are shown, not merely asserted. [VP-W1] is evidenced by the live `auth:channel-active` second-connect refusal. [VP-W2] is present in claim surfaces and keeps `record_integrity`/done-state `self_reported`. [VP-W3] is ruled and owner-ratified as above. [VP-W4] is evidenced by config-change and owed disposition on the existing live store, not by re-genesis.

7. `OI-S4-TOKEN-SCOPE` does not gate S4. A non-operator owed filing grants no authority and only creates a tracked obligation, so the present surface is hygiene rather than an integration hole. The follow-on should carry the m-7 sharpening: include `genesis` in scope along with `owed_item`/`owed_disposition`, because a once-at-store-birth system kind should not remain generally fillable even if stray submissions do not currently consume authority.

8. Later local history does not change this verdict. The `frank` main branch has advanced past the request's pre-integration base with S4 reconciliation ledger commits, and m-7 has filed an ACK that strengthens the F-GATE-3 and token-scope positions. I therefore ground this confirm on the stable branch tip `s4-wire-impl@6a23cf0` and the current relay trail, not on the stale main-base wording in the request.

## Confirmation

I confirm the master acceptance in `.relays/s4/s4-exit-gate/RECONCILE-orchestrator-planner-20260705-231116.md`.

No additional VP blocker before the operator's merge/tag gate.

## Verification

- Source request lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-vp-confirm/SITREP-orchestrator-planner-20260705-232359.md` -> OK.
- Source relay-root lint: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s4/s4-vp-confirm` -> OK.
- Referenced close relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-exit-gate/SITREP-orchestrator-planner-20260705-230525.md .relays/s4/s4-exit-gate/RECONCILE-orchestrator-planner-20260705-231116.md .relays/s4/s4-gateday/SITREP-orchestrator-planner-20260705-221608.md` -> all OK.
- Formal gate and fold trail lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-exit-gate/SITREP-planner-20260705-231210.md .relays/s4/s4-wire-impl/SITREP-planner-20260705-224500.md .relays/s4/s4-wire-impl/SITREP-planner-20260705-223638.md .relays/s4/s4-wire-impl/REVIEW-FOLD-implementer-20260705-224449.md .relays/s4/s4-exit-gate/SITREP-planner-20260705-230104.md` -> all OK.
- Branch pointer check: `git -C frank rev-parse --short s4-wire-impl` -> `6a23cf0`; `git -C frank rev-parse --short main` -> `a47381a`.
- F-GATE-2 row check: `git -C frank show s4-wire-impl:internal/fieldspec/registry.json` shows the five rows with `record_kind_in` predicates: `owner`, `source`, `target_surface`, `disposition_path` for `owed_item`, and `disposes_owed` for `owed_disposition`.
- F-GATE-2 fold stat: `git -C frank show --stat --oneline --no-renames 6a23cf0 --` -> exactly 3 files changed, 139 insertions.
- Token-scope check: branch registry row for `record_kind` has `operator` including `genesis`, `owed_item`, `owed_disposition`, `gate_resolution`, `disposition`, `diagnostics`, `config_change`; `*` includes `genesis`, `owed_item`, `owed_disposition`, `gate_resolution`, `disposition`, `diagnostics`.
- Later m-7 ACK read: `.relays/s4/s4-exit-gate/SITREP-planner-20260705-231210.md` ratifies F-GATE-3 and endorses `OI-S4-TOKEN-SCOPE` including `genesis`.
- `git -C frank status --short --branch` before filing this relay -> `## main`.
- Filed relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-vp-confirm/RECONCILE-orchestrator-reviewer-20260705-232713.md` -> OK.
- Filed relay-root lint: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s4/s4-vp-confirm` -> OK.
- INDEX row check after filing: `tail -n 6 master-docs/master/relays/INDEX.md` shows the `20260705-232713` confirm row at EOF.
- `git -C frank status --short --branch` after filing this relay -> `## main`.
- Harness root `git status --short` after filing -> unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master-docs/master/relays/INDEX.md`; no code edits and no `frank/` edits.

FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository: `fatal: not a git repository (or any of the parent directories): .git`; `frank/` status is `## main`.
