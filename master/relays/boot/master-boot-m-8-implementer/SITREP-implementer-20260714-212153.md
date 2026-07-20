## SITREP — m-8.implementer / Provider Adapters adversarial AUDIT waiting on planner matrix

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-8-implementer-wait
PARENT_DISPATCH_ID: master-boot-m-8-implementer
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-8.implementer
TO: m-8.planner, master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-9.implementer, m-1.planner, m-3.planner, m-4.planner, m-7.planner
BUNDLE_ID: m-8-provider-adapters
SUBJECT: independent m-8 adversarial source audit complete; final AUDIT comparison is correctly withheld because m-8.planner's promotion matrix is not yet present in the relay trail

Phase:
Report-only wait SITREP inside the authorized DESIGN-only AUDIT boot. This is not the final adversarial AUDIT, does not advance phase, and carries no design-lock, PLAN, implementation, credential, external-call, or provider-send authority.

Current artifact:
- Addressed boot: `master/relays/boot/master-boot-m-8-implementer/SITREP-orchestrator-planner-20260714-223510.md`.
- Required upstream planner boot exists at `master/relays/boot/master-boot-m-8-planner/SITREP-orchestrator-planner-20260714-223500.md`.
- Required upstream output does not: repeated `rg --files master/relays` checks found no m-8 planner `AUDIT`/promotion-matrix relay, and the live `master/relays/INDEX.md` tail ends with the four Division-II boot rows.
- Independent source pass completed against `references/pi`, `references/opencode` plus `references/opencode-notes.md`, `frank/` at exact `s11-close@502e06c`, and the locked m-1/m-3/m-4/m-7 contracts.

Independent adversarial checkpoints already established (E1; provisional until compared with the planner matrix):
1. **pi has promotable fixtures, not a promotable interface.** Its normalized text/thinking/tool-call/usage/finish/error stream grammar is useful conformance input (`references/pi/packages/ai/src/types.ts:333-476`), and its compatibility fields expose concrete replay/stream differences worth testing (`types.ts:482-597`). But `StreamOptions` lets the caller supply `apiKey`, overriding headers, provider env, retry/timeout policy, and arbitrary `onPayload` mutation (`types.ts:113-191`); wholesale promotion would cross m-7 credential/config, m-3 final-wire authorization, and m-4 policy ownership.
2. **opencode's factual schema must be split before reuse.** The models.dev conversion contains useful factual candidates (provider/model identity, API family, limits, modalities, reasoning, cost) (`references/opencode/packages/opencode/src/provider/provider.ts:1205-1248`), but its public `Model`/`Info` shapes co-locate arbitrary `options`, headers, env names, and a credential `key` (`provider.ts:1029-1054`), and `toPublicInfo` serializes the provider wholesale (`provider.ts:1072-1085`). No such combined object may become frank's catalog, snapshot, record, seat surface, or evidence.
3. **opencode contains explicit governed-send escape shapes to reject.** SDK resolution merges config/env/API credentials and model headers (`provider.ts:1417-1538,1666-1718`), accepts an arbitrary custom fetch, then either calls it/global fetch or dynamically installs/imports a provider package (`provider.ts:1730-1794`). Those are useful negative fixtures, not frank architecture: the m-7-hosted governed send boundary must be the only frank-offered provider-send path, with final authorization after translation/compatibility/endpoint/auth binding and no mutation after that point.
4. **Normalized event semantics are worth adapting, not inheriting.** opencode demonstrates stateful text/reasoning/tool-input/tool-call/usage/finish normalization (`references/opencode/packages/opencode/src/session/llm/ai-sdk.ts:76-278`) and the `reasoning_content` replay transformation (`references/opencode/packages/opencode/src/provider/transform.ts:286-313`). Frank must own the event grammar, error taxonomy, partial-stream terminal rules, retry/idempotency, cancellation, and backpressure; reference events become conformance cases against that contract.
5. **Routing and authority behavior must be cut out of m-8.** pi exposes provider ordering/fallback controls (`references/pi/packages/ai/src/types.ts:599-615`), while opencode's native runtime can dispatch tool calls in-process (`references/opencode/packages/opencode/src/session/llm/native-runtime.ts:103-137`). The former is m-4 policy/no-silent-fallback territory; the latter is m-9 request plus m-5/m-7 authority territory. Neither belongs in an m-8 adapter/catalog promotion.
6. **All three owner amendments are proven necessary, not optional.** Frank's present egress package is dormant away-mode outbox logic and flags ordinary model names (`frank/internal/egress/egress.go:1-7`; `frank/internal/egress/rules.go:22-43`); trusted config has lane roots/schema refs/suites/lane VCS but no provider credential/endpoint contract (`frank/internal/config/config.go:35-63,710-772`); the live route row records `chosen_model` but no provider/serving/compat/lane identity (`frank/internal/fieldspec/registry.json:175`). Therefore the m-3 provider-request-egress, m-7 trusted credential/config, and m-4/m-2 exact-lane routing-record amendments must all close before m-8 lock.

Duplicate/already-built gate:
- `frank/internal` and `frank/cmd` contain no provider-adapter/model-runtime package or frank-owned normalized provider contract. The landed overlap is boundary substrate only: the seat tool surface (`frank/internal/channel/server.go:24-31,76-85`), serialized submit/observe host (`frank/internal/engine/submit.go:51-125`), dormant away-egress scanner, trusted config host, and model-valued routing carrier.
- Result: the m-8 domain is **still open**, not already built. Existing substrate must be consumed through owner amendments rather than rebuilt or silently widened by m-8.

Done:
- Routing and phase authority verified.
- Role skill + protocol applied.
- Locked kickoff/domain/playbook anchors read.
- Independent adversarial pass completed with byte evidence.
- Incoming boot exact-file lint passed; root-mode noise was limited to pre-existing INDEX/lineage findings while the target file returned `OK`.

Not done:
- No final four-bucket AUDIT verdict issued.
- No row-by-row adversarial comparison against m-8.planner's promotion matrix performed, because that artifact is absent.
- No design authored or reviewed; no design lock, plan, code, credential access, provider call, external send, branch, commit, or PR action performed.

Blocked:
- The comparative half of the boot assignment requires m-8.planner's actual promotion matrix. Inventing its rows would violate the evidence-over-seniority and file-first rules.

Scope drift risk:
- High if a final verdict is issued before the planner bytes exist: it could miss a promoted escape callback, catalog secret/policy leak, or inherited send path. Low while this SITREP holds the lane.

Boundary contract status:
- Writes: final m-8 adversarial AUDIT relay only after the planner matrix arrives.
- Reads: planner promotion matrix + the four audited source sets.
- Target entity: master reconciliation input for the m-8 pre-design AUDIT.
- Downstream consumer: m-8.planner and master/VP AUDIT reconciliation.
- Contract: every promoted/adapted/rejected row must name its governance seam and preserve frank ownership; all three owner-amendment gates must be explicit.
- Proof: E1 file/line comparison in the eventual AUDIT.
- No-consumer action: defer the final verdict; do not speculate.

Tests / verification:
- E1: exact boot relay and all required local anchors inspected.
- E1: reference and frank byte anchors above inspected at exact `frank` head `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` (`s11-close`).
- E2: `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-m-8-implementer/SITREP-orchestrator-planner-20260714-223510.md` returned target-file `OK` (root reported pre-existing INDEX/lineage noise).
- E2: `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/boot/master-boot-m-8-implementer/SITREP-implementer-20260714-212153.md` returned `OK` on the report-of-record.
- E2: final `git -C frank status --short` returned no output at exact head `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`; the appended routing row survived at live `master/relays/INDEX.md` EOF.

ACTIONS_GIT_REF: wrote only `master/relays/boot/master-boot-m-8-implementer/SITREP-implementer-20260714-212153.md` and appended its routing row to `master/relays/INDEX.md`; no `frank/` edit; final `frank` head/status proof = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` with empty `git status --short` output.
FINAL_GIT_STATUS_SHORT: unavailable — `git status --short` at `/Users/jack/Programming/harness` exits 128 (`fatal: not a git repository`); `git -C frank status --short` returned no output at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`

Next requested action:
- m-8.planner authors and routes the lint-clean AUDIT/promotion matrix to `m-8.implementer`; the operator then re-presents that exact path. I will reopen the current bytes, compare every promote/adapt/reject row against the independent checkpoints above, and return the final four-bucket adversarial AUDIT to m-8.planner + master.
