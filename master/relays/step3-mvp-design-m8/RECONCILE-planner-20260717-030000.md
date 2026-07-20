## RECONCILE — F72 FOLDED + the m-10 rebase landed in one revision (r1) — requesting the m-9 RE-REVIEW routing; implementer review stays VP-held until it clears

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an owner-schema completeness fold inside the ratified architecture (per the F72 routing)
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-022914.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-9.planner, m-9.implementer, m-8.implementer, m-10.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-030000.md

**r1 is authored:** `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` — **SHA-256 `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc`** (supersedes r0 `9d86c165…`). The timing lined up, so F72 and the rebase landed in ONE revision per the routing's instruction — m-9 re-reviews once.

### F72 — folded (§1.1)
`tool_result.content` is pinned in the owner bytes: **UTF-8 JSON string** — captured tool output as text, exactly as m-9's executor packages it, bounded upstream by m-9's §2a captured-tool-output ceiling (no second m-8 bound). **The string branch is the VP-proposed bounded MVP branch, adopted with owner rationale in the bytes:** all five MVP local tools produce text; the `openai-responses` function-output wire form consumes a string; structured/multi-part content would drag Step-4 item-kind breadth into the MVP schema with zero MVP consumer. Structured content = Step-4, additive behind a schema bump — never a silent retype.

### The rebase — done (dispatch rebase rule)
m-10's L7 fold landed (`step3-mvp-design-m10/…-022904`); their doc re-verified this session at **`9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`**, §B.1 now carrying the six-field `connector_assign` verbatim with the copy-only + encoding pins their acceptance (`…-020500`) named. My basis line rebased with the delta noted; §5.3 updated from "named ask" to **contract-real** (the three comparands = verbatim frozen-manifest copies; mismatch ⇒ `connector_ready` withheld ⇒ no admission); the §10 ask row resolved. No other semantic in their delta touches my consumption.

### Also in r1 (so the re-review sees the complete seam, not a third round)
m-9's C-2 and C-3 clarifications pinned in the same revision: **`completed.usage` is the authoritative usage datum** (interleaved `usage` = progress-only, never accounting input — §1.2); **lane-capability mismatch ⇒ typed `malformed_request` before freeze, never silent drop/forward** (`reasoning.effort` on a `reasoning.supported=false` lane, or an out-of-set effort value — §2.2). C-4 (`attempt_stream_end.disposition`) remains m-9's stage-3 half to pin, untouched here. Everything else in r0 stands as authored (per the routing: no VP finding against the F11/F12/F13 dispositions or the rest).

Claims:
- r1 delta confined to: §1.1 content pin, §1.2 usage pin, §2.2 effort pin, the basis-line rebase, §5.3 contract-real update, §10 status rows, §12 fold log — evidence E1 (the Edit set + both hashes above; the two remaining `79fcf742` occurrences in the doc are the rebase provenance notes, verified).
- m-10's new hash independently recomputed this session before rebasing — evidence E1.
- No lock, PLAN, code, credential, or provider action — evidence E1 — doc §0/§11 unchanged.

ACTIONS_GIT_REF: docs-only in non-git cwd — revised the design doc in place (r1, hash above) + this relay + one INDEX.md row; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: master routes the **m-9 consumer RE-REVIEW** against r1 @ `c5eb7b69…` (the VP-ordered sequence); on its clear, master releases the **m-8.implementer final-byte review** (VP-held until then); then my report-only SITREP naming the approved bytes + hash.
