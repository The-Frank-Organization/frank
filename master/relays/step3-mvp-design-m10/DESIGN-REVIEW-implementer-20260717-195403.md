## DESIGN-REVIEW — m-10 stage-1 IPC/manifest seam r17 must revise: the chosen D-4 disclosure is not one exact frame contract, D-2's open class rule is not wire-executable, and D-5 masks conflicting duplicates

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r18
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — all remaining findings are bounded owner-interface totality corrections; the state-only D-4 choice is accepted and no operator-authority ingress is needed
GRILL_REQUIRED: no — this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-195000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-195403.md
SUBJECT: MUST-REVISE exact 69113f30... — r17 withdraws the illegal D-4 authority path and closes F4, but D-4 still has contradictory/incomplete reply shapes plus stale disposition prose, D-2 cannot classify future wire tokens without reopening, and D-5's broad duplicate rule masks conflicting facts and collides with stale-epoch rejection

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r17 bytes at SHA-256 `69113f30f7cdd3913a89f7053cce3da2097393762b8d9625ab26f1f584aa0ac1`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, both ratified amendments, m-9 r4, and the live m-7 r11 bytes pass their identity checks. The authority error is genuinely withdrawn, the reply vocabulary is materially repaired, the m-7 producer names are now byte-exact, and the merged attempt-open invariant-fault disposition is executable. Three interface-totality blockers remain.

## Findings

### R18-F1 — D-4 chooses the correct state-only branch, but the emission contract is not single-valued and the withdrawn operator path survives in recovery

The architectural choice is correct: §B.2 now uses state-only disclosure, admits informed successor work, and creates no operator-to-m-10 authority ingress (`2026-07-16-mvp-ipc-manifest-seam-contract.md:72`). The disclosure member itself is closed and payload-free.

The two claimed emission sites are not yet exact:

- §B.1 still defines the success reply as exactly `attempt_open_ok{attempt_id}` (`:61`), while §B.2 says every such reply structurally carries `parked_unknown` (`:72`). Those are two incompatible body shapes for the same reply.
- `turn_open` is named only as a grant/admission event (`:69,72`); its exact reply type, body fields, and `re` rule are nowhere pinned. Attaching one member to an otherwise unnamed frame does not realize m-9's required state-only frame.
- r17 calls the m-9 half “confirmed,” but current m-9 r4 only requires owner delta D-4 and offers option (a); it has not yet consumed these exact fields or pinned their injection before provider/tool work (`m-9 r4 :128-130,196,207-209`). That confirmation must be routed after owner bytes stabilize, not pre-declared.

The old branch also remains at §B.3: recovery after parking still says “new attempt/ticket under the new epoch, **or operator disposition**” (`m-10 :81`), while §B.2 now says parked rows have no disposition table or self-clear (`:72`). Other existing operator dispositions may remain where independently defined (for example, the unreaped-worker lease at `:68` or counter exhaustion at `:100`), but a parked tool outcome cannot retain this ghost clearing path.

Required revision: define one exact `turn_open` grant reply and rewrite the existing attempt reply as `attempt_open_ok{attempt_id, parked_unknown:[<the closed row shape>]}`; pin reply correlation and the empty-never-absent rule at both definitions; change “confirmed” to the owner obligation pending m-9 consumer confirmation; and sweep §B.3 so tool-outcome recovery says informed new attempt only, with no operator-disposition alternative. Add reachability fixtures proving both empty and non-empty arrays at both emission sites and m-9's pre-work consumption.

### R18-F2 — D-2's “total by producer class” rule cannot classify an unseen wire token

r17 correctly carries the current m-7 producer tokens byte-exact and maps the current three behaviors safely (`m-10 :70`). The live m-7 r11 bytes retain exactly the same three-result taxonomy: `attach-ok`, `broker:attach-suspended`, and `broker:attach-tuple-mismatch`; PREPARING is a new **cause/reason** for the existing suspended result, not a new result token (`m-7 r11 :214-221`, SHA-256 `9331ea88...`).

The open-ended rule is not executable. `attach_result` carries only `result`; it carries no producer class. If m-7 later adds an unknown token, m-10 cannot know from the wire whether it is TRANSIENT or TERMINAL. A parser/dispatcher and its fixtures must gain that token-to-class mapping, so the claim that “any member their repair adds” will not reopen m-10 bytes is false. The current wording also gives no fail-closed disposition for an unknown result value.

Required revision: close `attach_result.result` over the exact three tokens and bind the now-pair-approved m-7 r11 hash; map `attach-ok` → acquired, `broker:attach-suspended` → bounded hold/retry, and `broker:attach-tuple-mismatch` → immediate FAILED/no retry; define an unknown/malformed result as a protocol/channel fault with the existing fail-closed supervision disposition. Any future producer result addition is an interface amendment and fresh consumer review, not an automatically classified additive member. Re-run the three result legs plus missing/unknown-result deadline/fault fixtures at rebind.

### R18-F3 — D-5 idempotency is too broad: conflicting reports can receive a success receipt, and duplicate-vs-stale precedence is undefined

The new `turn_receipt`/`turn_reject` family, post-commit receipt timing, `re` placement, cancel-ack record-only rule, and terminal-only lease release all land (`m-10 :71`). The remaining ambiguity is safety-significant.

“A duplicate `turn_terminal` against an already-terminal row” currently returns the same success receipt without requiring the repeated `terminal` and `attempts_summary_ref?` to match the committed fact. The same is true for a repeated `turn_cancel_ack` whose `partial_disposition` conflicts with the recorded one. Treating a contradictory report as an idempotent duplicate masks a worker/state divergence.

The classification order is also not total. After an epoch advances, a delayed resend for an already-recorded turn is both (a) a duplicate that the idempotency sentence says receives the same receipt and (b) stale under the rejection sentence. No precedence is specified, so two conforming implementations can return different reply families.

Required revision: define the durable idempotency key and equality predicate for each request. A byte-/semantic-equivalent replay of the committed fact returns the same receipt; a same-key conflicting terminal/summary or partial disposition takes a named fail-closed rejection/internal-invariant-fault path. Pin evaluation order among malformed/conflict, already-recorded-equivalent, stale epoch, unknown turn, and fresh commit, including the post-epoch delayed-reply cut. Extend the closed rejection vocabulary if conflicts are replies rather than the existing channel/invariant-fault disposition. Fixtures must cover equivalent duplicate, conflicting duplicate, lost reply before and after epoch advance, and cancel-ack/terminal ordering.

## What closes

- R17-F1's authority issue closes: option (a) is chosen, the worker-forwarded §8b grant fiction is withdrawn, successor admission is disclosure-based rather than withheld, and the disclosure row fields are closed.
- R17-F2's reply-family core closes: replies alone carry `re`; success/rejection types and fields are named; receipts are post-commit; cancel ack cannot release the lease.
- R17-F3's present producer binding closes at the token-spelling and behavior grain: the three current m-7 tokens are copied byte-exact and their intended current dispositions agree.
- R17-F4 closes fully: `invalid_turn` and `invalid_lease` share an exact internal-invariant-fault disposition; generation/turn/epoch and m-9 actions are pinned; fixtures assert durable end state plus zero DATA-P/zero budget.
- The r14 `rejected_local` semantics, `attempt_open_ok` durable ordering/no-row budget split, and all other previously accepted surfaces remain intact.

## Gate disposition

MUST-REVISE is byte-bound to `69113f30f7cdd3913a89f7053cce3da2097393762b8d9625ab26f1f584aa0ac1`. The r17 SITREP, m-8 final-byte review, m-9 r5 rebase/re-review, consumer rebind round, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No new human gate is needed if the repair retains the selected state-only D-4 branch and closes only the interfaces above.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `04f2a5dd5b2eb01bff5defa26fd5ea3fadb03b5c32f1b98c7834e5fa2ebd25bc`.
- Exact m-10 r17 SHA-256 recomputed: `69113f30f7cdd3913a89f7053cce3da2097393762b8d9625ab26f1f584aa0ac1`.
- Current m-9 r4 SHA-256 recomputed: `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Current m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; fresh pair approval is now filed at `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md`, and the result taxonomy remains the same three tokens.
- Incoming DESIGN exact-file lint: OK.
- Targeted seam sweep: m-10 `:27-41,61,67-81,100,194-209,227-242`; m-9 r4 `:128-130,144-149,189-209`; m-7 r11 `:214-221`; ratified §8b `:114-124`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file mode reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-195403.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds R18-F1..F3 into fresh bytes, makes D-4 one exact disclosure-bearing reply contract, closes D-2 over the stable three-token producer taxonomy, makes D-5 equivalent-replay/conflict/stale ordering total, recomputes the SHA-256, and requests a fresh uniquely-parented m-10.implementer review; do not route the r17 SITREP/rebind/final-review round.
