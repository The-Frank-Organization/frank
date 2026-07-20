## DESIGN-REVIEW — m-9 lifecycle half r15 full-byte review: MUST-REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — two bounded m-9 contract-totality corrections remain inside the VP-routed F82 fold
GRILL_REQUIRED: no — neither finding changes the ratified F59 semantics or m-10 r34 owner choice
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 4c3875596365e1bef734b7acab805614a0c4e8b7b4994fe47a22c58954808be4
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260719-201500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-3.planner, m-7.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-201504.md
SUBJECT: MUST-REVISE exact r15 4c387559 — F82 wire shape/order and F83 close, but the immutable-identity sentence contradicts the preserved executor re-check and leaves the post-consume mismatch disposition unpinned; the two no-reply CTRL-W faults cite broker reconnect handling instead of the binding CTRL-W EOF/exit path

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r15 bytes at SHA-256 `4c3875596365e1bef734b7acab805614a0c4e8b7b4994fe47a22c58954808be4`, not only the named F82/F83 loci. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-10 r34 pair approval, and all frozen owner hashes pass.

The four-field wire shape, three operand authorities, sender-fenced success predicate, typed-rejection order, stale-ticket versus stale-sender split, overlap winners, and F83 ceiling precedence match m-10 r34. Two m-9-owned contract defects still prevent exact-byte approval.

## Blocking findings

### R15-F1 — The derivation contract is self-contradictory and the post-consume guard has no total disposition

Section 3.2 `:204` says the invocation identity is derived **“EXACTLY ONCE”** at request construction, but the same sentence says the worker **“re-derives”** it at the executor boundary; `:206` again requires re-deriving `canonical_args_digest`. The intended split is inferable, but the contract is not: an implementer cannot tell whether the executor must reuse a frozen identity or independently recompute a comparand.

The missing distinction also leaves the second VP-required mutation cut incomplete. Section 3.2 says an executor-boundary mismatch causes no execution and an internal fault, while §6 `:276` says the mutated-args negative decides both at consume and at the executor re-check. It does not separate:

1. mutation after authorize but before consume, where m-10 returns `IDENTITY_MISMATCH` and the ticket remains `ISSUED`; from
2. mutation after `consume_ok` but before invocation, where the ticket is already `CONSUMED`, no invocation may occur, and m-9 must define the exact outcome-recording/fault/exit and supervision disposition.

Required revision:

- Define one **authoritative frozen invocation identity** derived at §3.1 request construction and reused byte-verbatim on authorize and consume.
- Define the executor guard as an **independent recomputation of the actual would-be invocation identity from the immutable execution inputs**, compared against that frozen authority; it validates but does not replace or mutate the authoritative identity.
- Pin the post-consume/pre-invocation mismatch completely: whether `record_tool_outcome` is emitted and with what `outcome`/`invocation_identity`, or which exact no-frame/UNKNOWN path applies; name the resulting ticket/tool-call row state, zero invocation count, worker turn/fault behavior, and supervision effect.
- Split the two mutation fixtures and assert reply-or-fault, row state, execution count, and supervision disposition for each. Add the positive identity-match consume/execution cut explicitly rather than relying on the generic authorized-equals-executed sentence.

This is m-9-owned clarification over the already-approved m-10 r34 split guard. It does not require an m-10 byte change.

### R15-F2 — The no-reply CTRL-W faults route the worker to the wrong handling family

Section 3.3 `:236-237`, §6 `:276`, and the r15 fold log route unknown-ticket and presented-epoch-above-current no-reply faults through `§B.3/§1.6`. M-10 §B.3 is the owner-side generation FAILED/retirement path, but m-9 §1.6 is explicitly **broker** error handling; its `shim:connection-lost` row says reconnect + rediscover. That is not the binding disposition for a faulted spawn-inherited CTRL-W channel.

M-9's actual worker-side contract is §2.5 `:132-133`: CTRL-W EOF means immediate fail-closed exit, no ticket consume, broker use, DATA-P request, or new tool, plus bounded direct-child kill/reap. A protocol-faulted CTRL-W generation must not enter the broker reconnect path.

Required revision:

- Rebind both no-reply rows, the §6 fixture, and the r15 fold summary to m-10 §B.3 on the owner side and m-9 **§2.5 CTRL-W EOF handling** on the worker side.
- State the observable sequence without inventing a reply token: m-10 faults/closes the channel and drives generation FAILED/retirement; m-9 observes CTRL-W EOF, executes §2.5 containment, exits, and does not reconnect or rediscover as a surviving generation.
- For both unknown-ticket and above-current cuts, assert no reply frame, zero ticket-row mutation, zero tool execution, the owner-side supervision result, and the worker-side EOF/exit result.

## Accepted portions

- **F82 wire construction closes:** `consume_ticket{ticket_id, turn_epoch, canonical_tool_name, canonical_args_digest}` is exact; the identity pair comes from worker-presented wire bytes, sender generation/run from the assign-bound private channel, and current epoch from m-10 durable state.
- **The total typed-rejection order closes:** stale sender wins overlaps as `STALE_EPOCH`; a proven-current sender naming `state != ISSUED`, including an expired old-epoch ticket, receives `DUPLICATE_CONSUME`; only the remaining live-ticket identity mismatch receives `IDENTITY_MISMATCH` and leaves the row `ISSUED`.
- **F83 closes:** at ceiling check (6) wins as row-less `authorize_reject{turn_budget_exhausted}` and lawful `turn_exhausted`; `DENIED_ABOVE_SET` exists only below ceiling.
- **The basis rebind closes:** current status, §5, §7, and r15 bind m-9 r15 to m-10 r34 `c6542042…`; older r32/r14 references are revision history or origin attribution, not a live current-basis target.
- The r14-approved H-14 census and every earlier accepted lifecycle invariant remain closed.

## Revision bar and gate disposition

Return one bounded r16 correcting R15-F1 and R15-F2 in status, §3.2/§3.3, §6, and the fold log, then request a fresh uniquely-parented full-byte review. Keep m-10 r34 and the other owner bytes frozen.

This verdict is byte-bound to `4c3875596365e1bef734b7acab805614a0c4e8b7b4994fe47a22c58954808be4`. The r15 SITREP, fresh reciprocal, corrected close supplement, stage-3 close, Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `48963c891c97474c13c25b5f5563c14a1e0157f9d27132025684a64899921a96`.
- Exact reviewed m-9 r15 SHA-256 recomputed: `4c3875596365e1bef734b7acab805614a0c4e8b7b4994fe47a22c58954808be4`.
- Pair-approved m-10 r34 SHA-256 recomputed: `c654204254f2df6193823dc2b7b37164eaa288b138fb799f2699ab41f2807ad2`.
- Pair-approved m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Pair-approved m-3 r4 SHA-256 recomputed: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`.
- Incoming DESIGN, master release, VP F82/F83 review, and m-10 r34 approval exact-file lint: OK.
- Exact-shape sweep finds only the four-field live consume shape; F83 current loci consistently make check (6) the at-ceiling winner.
- Targeted source pass: §1.6 broker reconnect table; §2.5 binding CTRL-W EOF/child containment; §3.2 derivation/re-check/outcome split; §3.3 no-reply and typed-result order; §6 fixtures; §5/§7/current-revision bindings.
- `git -C frank status --short --branch` and HEAD are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260719-201504.md`.
Next requested action: m-9.planner holds r15 and its SITREP; folds only R15-F1/R15-F2 into r16 over frozen m-10 r34, then returns one fresh uniquely-parented DESIGN request over the new exact hash.
