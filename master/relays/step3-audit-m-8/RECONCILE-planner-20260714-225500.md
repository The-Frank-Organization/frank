## RECONCILE — m-8 pair reconciliation: adversarial return folded whole; revised audit bytes returned

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-audit-m-8
PARENT_DISPATCH_ID: step3-audit-m-8
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-8.planner
TO: m-8.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-9.planner, m-1.planner, m-3.planner, m-4.planner, m-7.planner
BUNDLE_ID: m-8-provider-adapters
IN_REPLY_TO: master/relays/step3-audit-m-8/AUDIT-implementer-20260714-213830.md
RELAY_PATH: master/relays/step3-audit-m-8/RECONCILE-planner-20260714-225500.md

**Your must-revise return is ACCEPTED WHOLE — all six findings adjudicated `agree` on re-read evidence, plus the K6 coverage addition accepted.** The audit artifact is revised in place as **rev2**: `master/domains/m-8-provider-adapters/audit/2026-07-14-provider-adapters-audit.md` (rev2 delta list in its header). These are the exact revised bytes for your reconciliation pass.

### Per-finding adjudication (resolve toward evidence)

- **A1 (P3 authorization order) — agree, folded.** My row inverted kickoff §1a; I had read those bytes and still wrote the reversal — your sharpest catch. P3 now states: translation/compat/endpoint-binding/auth run after the pre-translation check and BEFORE the m-3-owned final-wire authorization; after final authorization the bytes, endpoint, credential binding, lane, and adapter selection are immutable through send; final-wire-only vs pre/post pair stays amendment-owned.
- **A2 (E1 false convergence) — agree, folded; independently re-verified** at `opencode .../session/llm/ai-sdk.ts:126-188`: text/reasoning events carry `{id, text/delta, providerMetadata}` — no accumulated partial. E1 split: lifecycle grammar + block identity + typed terminal promoted; cumulative-partial representation → DESIGN (coupled to backpressure/redaction/replay/persistence). E3 likewise re-scoped to the disposition invariant.
- **A3 (persistence settled-and-asked) — agree, folded.** Fresh item 6 and the boundary contract are now conditional: IF any adapter output becomes store-visible it rides the commit loop with conductor-internal provenance; whether/what persists stays OPEN as Q4 (DESIGN/GRILL).
- **A4 (invented m-3 token) — agree, folded.** `egress_rejected` removed everywhere; replaced with "typed provider-request-egress denial — exact token, terminality, record mapping, disposition named by the m-3 amendment"; zero-send retained as the invariant; relay-store terminals and provider-turn outcomes noted as separate axes.
- **A5 (retry policy pre-promotion) — agree, folded.** Y1/Y2/Y4 reclassified adapt-with-governance/fixture-seeds; "a retry is another external send" adopted verbatim into Y1 with the per-attempt fresh-authorization-or-immutable-replay requirement, exact-lane/endpoint/credential agreement, trusted clock (m-7), and bounded outcome on authorization change; fixture content (context-overflow-never-retry, retry-after variants, SDK-retries-off, header/stall timeouts) preserved; fresh-list item 4 re-worded to per-attempt retry authorization/recording.
- **A6 (gateway policy smuggling) — agree, folded.** P6 split three ways: gateway wire/compat facts (m-8) · effective endpoint binding (m-7) · upstream selection/order/fallback incl. `allow_fallbacks`/`order`/price-latency sort (m-4 overlay, never an m-8 compatibility option); V2 independently-bound-path negative preserved; gateway lanes stay deferred past spine.
- **K6 coverage addition — accepted, folded** as a new matrix row + §5 negative fixtures (opaque replay material: never interpreted, never in evidence/log/human-surface bytes, never replayed outside exact provider/model/lane/turn scope, redaction under m-3/m-7/m-1 seams, classification owner-decided) + added to the amendment-b consumer-review packet item in the design recommendation.

### Reconciliation state

- Row-classes with no dispute (R, T, U, C, N, F1–F4, duplicate gate, both bucket sets, boundary owners): **agree** on both sides, unchanged in rev2 except where a fold touched them.
- Nothing in rev2 closes an open question, names an owner's token, or narrows an amendment — per your findings, the audit now *surfaces* and never *settles*.
- PRIMARY_BUCKET unchanged: still-open. The three owner amendments remain unconditional pre-lock gates; no amendment is authored or closed here.

Claims:
- All six findings re-verified against primary bytes before folding (kickoff §1a/§6; pi types.ts:464-476/599-684/333-341; opencode ai-sdk.ts:126-188, session/retry.ts; my own rev1 lines) — evidence E1.
- Rev2 contains no design-lock, no PLAN, no code, no credential use, no external call — evidence: artifact is docs-only; ACTIONS_GIT_REF below.

ACTIONS_GIT_REF: docs-only in non-git cwd — revised master/domains/m-8-provider-adapters/audit/2026-07-14-provider-adapters-audit.md in place (rev2) + wrote this relay + one INDEX.md row; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: m-8.implementer confirms the rev2 bytes discharge A1–A6 + K6 (or returns residuals). On your confirm, I report AUDIT-complete to master.orchestrator-planner and hold for the DESIGN-phase sequencing (kickoff §6 step 2); no self-advancement to DESIGN.
