## DESIGN-REVIEW — m-8 provider contract r7 final-byte review: owner rebases land, but the local-reject commit barrier is not implementable and cancellation is recorded as transport failure

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r7
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — both findings are bounded interface-totality corrections inside already-ratified ownership; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — this review does not reopen topology, policy, secret custody, or an operator-locked choice
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260717-205500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-210409.md
SUBJECT: MUST-REVISE exact r7 b805edab... — r21 makes attempt-open durable before DATA-P, but supplies no result-commit receipt for fixture 16's later cross-channel barrier; cancellation also maps a deterministic zero-wire outcome to transport_failed/failed

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r7 bytes at SHA-256 `b805edab019400d1bd6505dd17beddf1e9b092a05bd1a1b2fbe549cefc721083`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-10 r21 at `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852`, and the bounded m-9 confirmations pass their identity checks.

The r6→r7 rebase is valid at its claimed surface: m-10's `attempt_open_ok` is reply-class, is emitted only after the attempt-row transaction commits, and m-9 gates DATA-P issuance on that reply. The no-row rejection half and the m-8 epoch backstop compose without a double-write. The four local-reject reason tokens, the frozen/suppressed/derived-and-censused model, and the reverse-order fixture text are all present. A fresh whole-document pass nevertheless finds two active defects. These exact bytes cannot receive final pair approval.

## Findings

### R7-F1 — Fixture 16 requires a durable cross-channel commit barrier that no owner-real protocol message implements

Fixture 16 requires both the CTRL-C `rejected_local` emission **and m-10's observable terminal-row commit** to occur before m-9 can receive or act on the typed DATA-P reply, and requires a DATA-P-first mutation to fail (`2026-07-17-mvp-provider-contract.md:222`). Section 1.3 repeats the same stronger claim by saying m-10 closes `REJECTED_LOCAL` before the typed DATA-P return completes (`:80,84-85`).

M-10 r21 does not provide that barrier. Its owner-real protocol defines:

- `attempt_open_ok` as the durable receipt for the **initial row creation**, before DATA-P issuance;
- m-8's later one-way `attempt_result{...rejected_local(...)}` emission on CTRL-C; and
- m-10's terminal close when that result is received.

It defines no `attempt_result` receipt/ack after the terminal-row commit (`m-10 r21 §B.1:61`). Therefore m-8 can guarantee only **CTRL-C emission before DATA-P reply completion**. Sending on CTRL-C first cannot prove that a separate process received and durably committed it before m-9 receives bytes on DATA-P. A fixture that merely interposes and withholds DATA-P until it observes the database would add a test-only synchronizer that the production contract lacks; the reversed-order mutation would test the harness, not the stated runtime guarantee.

Required revision — choose one honest branch and route any owner delta:

1. **Emission-order branch:** narrow §1.3 and fixture 16 to the owner-real R14-F1 guarantee that m-8 emits `rejected_local` before completing the typed DATA-P reply. Do not claim the m-10 terminal commit is already observable before m-9 receives the reply. Preserve the reverse-order negative at the emission boundary.
2. **Durable-commit branch:** have m-10 author a reply-class receipt emitted only after the terminal-row transaction commits; require m-8 to await that receipt before releasing the DATA-P reply; obtain the corresponding m-9/m-10 confirmations and add lost-receipt/crash cuts. M-8 must not invent that receipt in its own document.

This is a failed fold of R5-F3, not a rejection of `attempt_open_ok`: r21 establishes the row **before the attempt starts**, but it does not acknowledge the later `rejected_local` terminal close.

### R7-F2 — Cancellation can cross zero transport yet is unconditionally recorded as `transport_failed` and `phase=failed`

Section 1.4 expressly permits cancellation after authorize but before transport invocation: the attempt closes `cancelled{partial:none}` with **zero wire send** when transport has not yet been invoked (`m-8 :97`). The total table nevertheless labels every cancelled outcome “sent,” emits CTRL-C `transport_failed`, and maps m-9's E0 phase to `failed` (`:90`). That is false for the named zero-wire branch and conflicts with the consuming lifecycle:

- m-9 maps provider `cancelled` to `stream_cancelled` (`m-9 r5 §2.2:93`);
- cancellation lands the logical turn at `turn_cancelled`, with interrupted distinct from failed (`§2.3/§2.5:110,125`); and
- the m-9/m-10 cancellation family records cancellation as a distinct durable fact (`§2.9:151-156`; m-10 r21 §B.1:71).

The problem also exists for a wire-crossed cancellation: user/operator cancellation is not a transport failure merely because the HTTP context was aborted. R7's own honesty rule reserves `transport_failed` for an actual transport failure and introduced `rejected_local` precisely to avoid a false zero-wire label (`m-8 :76,93`).

Required revision:

1. Add an honest cancellation disposition to the m-8 CTRL-C attempt-result vocabulary, or route another exact owner-real representation that distinguishes cancellation from transport failure.
2. Split the table as needed to preserve the per-attempt wire fact: pre-transport cancellation is zero-wire; post-invocation cancellation records that transport was invoked without relabeling the cancellation as a failure.
3. Reconcile the new/selected disposition with m-10's row state and m-9's `stream_cancelled`/E0/turn-terminal mapping. This moves the m-10 closed enum and therefore requires Master-routed m-10 owner bytes plus fresh consumer confirmation; m-8 cannot silently widen it.
4. Add fixture cuts for cancellation before transport invocation and after request-write start, asserting the exact transport counters, CTRL-C result, m-10 terminal row, m-9 stream view, and turn disposition.

## Accepted portions

- The r7 m-10 rebase is byte-faithful for the `attempt_open_ok`/`attempt_open_reject` surface. The durable row exists before normal DATA-P issuance; rejected opens create no row and no budget charge.
- R5-F1's epoch and deterministic-integrity totality repairs survive: epoch-class DATA-P replies are attempt-inert at m-8, deterministic integrity refusals use `rejected_local(internal_integrity_fault)`, and fixture 17 covers both row-fate branches.
- R5-F2 closes: the active F12 prose consistently uses the frozen/suppressed/derived-and-censused model and prohibits only uncensused injection.
- M-9's bounded `internal_integrity_fault` and epoch/budget confirmations compose with r7. The current full m-9 r5 lane is separate and does not cure or create either finding here.
- No finding changes the selected HTTP client, provider dialect, credential custody, egress-policy ownership, lane model, or topology.

## Revision bar and gate disposition

Return fresh bytes that:

1. make fixture 16 assert an implementable owner-real ordering — emission-before-reply, or a newly owner-authored durable result receipt — without conflating the initial `attempt_open_ok` with the later terminal close; and
2. give cancellation a truthful attempt-result/row/E0 mapping across both zero-wire and wire-crossed cuts, with m-9/m-10 owner returns for every moved seam.

The new SHA requires a fresh uniquely-parented DESIGN-REVIEW. This verdict is byte-bound to `b805edab019400d1bd6505dd17beddf1e9b092a05bd1a1b2fbe549cefc721083`. The stage-2 approval SITREP, m-9 closure consumption of an approved m-8 hash, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `46d2d694f7a0cc83714080577913b0f28db4a3937096846143ab909954e27a6f`.
- Exact reviewed m-8 r7 SHA-256 recomputed: `b805edab019400d1bd6505dd17beddf1e9b092a05bd1a1b2fbe549cefc721083`.
- Pair-approved m-10 r21 SHA-256 recomputed: `f4012ec5b723ec98f9570a0d9470c5c41b3cd56b1e7599be615962917d770852`; no `attempt_result` receipt/ack exists in its message set.
- Current m-9 lifecycle r5 SHA-256 recomputed: `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`; its cancellation and stream mappings were checked at §§2.2, 2.3, 2.5, and 2.9.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-210409.md`.
Next requested action: m-8.planner folds R7-F1 on one explicit ordering branch, routes R7-F2's cancellation disposition/row mapping to m-10 and the consumer mapping to m-9 through Master, then returns a fresh uniquely-parented byte-bound DESIGN request; do not file the stage-2 approval SITREP on r7.
