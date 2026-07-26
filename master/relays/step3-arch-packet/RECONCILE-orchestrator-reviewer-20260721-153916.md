## RECONCILE -- REVISE: the grain boundary is accepted, but provider content settlement and the persisted resume snapshot still have impossible cross-owner ordering/state claims

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev8 is not ready until the two remaining master-level resume seams are coherent
GRILL_REQUIRED: no -- D7/build-it-properly and the decomposition-versus-pair-design boundary are settled; the return is limited to cross-owner ordering, identity, persistence, and transition semantics
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-074500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- grain correction accepted and no operator arbitration needed; fix provider settlement as terminal-plus-durable-content evidence and persist actual manifest bytes with a pending-to-final disposition receipt gate

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-074500.md` at SHA-256 `b53c0359f8256b1d80e6fb1e0ce028dc7da61f9e042b760500aef8985ce58667`.

Proposed amendment rev8: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `9d5e8a34249e56f350d796dc8d4ad0aba24c73c0e341ba9fbbc3986ffabe57e3`.

## Grain disposition -- ACCEPTED

The level correction is right. Master should ratify the load-bearing decisions, acceptance properties, ownership, and F73 obligations; m-9/m-10 should detail record payloads, canonical encoding, sequence rules, locking/segment mechanics, rotation, and the exhaustive crash table in their own DESIGNs with pair review. No operator grain arbitration is needed.

Rev8's D1 acceptance properties are sufficient at decomposition grain: fsync is the durability point; the writer boundary must be enforceable; recovery must select one deterministic valid prefix; identity is fail-closed; content must precede outcome settlement; and the log path has one carrier. The detailed mechanism remains correctly held behind the owner-design/F73 ladder. One ownership wording cleanup is required but not a separate blocker: if m-9 chooses a local OS lock, the mechanism can be m-9-owned; if it chooses **m-10-ordered per-generation segments**, that branch is an m-10-producer/m-9-consumer seam and must be jointly designed/reviewed, consistent with §6 and `:231-233`, not called an m-9-only obligation.

## Findings

### F105-D2-R8 -- BLOCKER: provider settlement is not ordered behind worker durability, and content-record fsync alone does not place content in the valid prefix

Rev8 `:182-200` correctly removes the ownerless round index, adds positive provider identity, scopes reconciliation over the ancestry, and chooses durable-content-before-outcome. The tool half is almost sound, but the stated ordering is one step short: D1's resumable prefix ends at a valid `round_marker`, while D2 orders only the `tool_result` content-record fsync before `record_tool_outcome`. A crash after content fsync and m-10 outcome commit but before the enclosing marker is durable leaves m-10 `EXECUTED` while recovery truncates the content outside the valid prefix. To prove `settled => content in the durable valid prefix`, the **content plus its enclosing marker** must linearize durably before the settlement-producing step, or D1 must define a different prefix rule. This is an acceptance-order decision, not record-format detail.

The provider half has a larger producer-order contradiction. Frozen m-8 emits `attempt_result{...sent_completed|transport_failed|...}` at the attempt's own terminal boundary, independently of m-9 (`m-8 provider contract :76,80-93`), and m-10 terminalizes `provider_attempts` from that CTRL-C fact. No current m-8/m-9/m-10 message lets the worker's provider-output/marker fsync occur before that terminal commit. Therefore raw `provider_attempts=completed|failed` cannot by itself enter `settled_providers` while preserving rev8's invariant. The phrase "settled tool-round for provider output" names no owner-real step.

The decomposition must define **resume settlement** for provider content as a composite fact, for example: canonical provider terminal **AND** a durable m-9 content-ready reference/receipt bound to `{turn_id, attempt_id, log-prefix-or-marker digest}`. m-10 may keep its existing canonical attempt terminal unchanged, but it must not emit a `settled_providers` entry until the content-ready half committed. Missing content-ready evidence maps to `uncertain` or `content_lost/degraded`, never settled. The exact frame/table implementation belongs to the pairs; the conjunction, owners, order, and receipt dependency belong here because they alter the m-9/m-10 seam and DAG.

Two schema-totality corrections travel with that decision:

- Full-ancestry entries need source `turn_id`. Frozen tool-call identity is scoped by `UNIQUE(run_id, turn_id, tool_call_id)` (`m-10 seam :186-192`); `tool_call_id + args_digest` alone is not identity-exact across continuation turns. Provider entries should carry the same source-turn scope.
- The provider state partition must be total over m-10's canonical states, not only `completed|failed`: the frozen family also includes determinate `denied`, `REJECTED_LOCAL`, and `CANCELLED`, plus `UNKNOWN_PROVIDER_OUTCOME`/`PARTIAL_STREAM` (`m-8 :82-93`; m-10 seam `:279-280`). Pin which states mean settled content, definite no-content/discard, and uncertain/partial so no state reaches absent-means-not-happened accidentally.

Required correction: make the durable valid-prefix marker precede tool settlement; define provider resume-settlement as terminal-plus-durable-content evidence with an owner-real receipt dependency; add source-turn identity; and make the provider state partition total. Pair DESIGN remains free to choose the concrete record/frame/table mechanics that prove those properties.

### F105-D3-R8 -- BLOCKER: `resume_snapshot` stores a digest but promises bytes, and the disposition is circular at admission

Rev8 `:202-223` correctly chooses one `turn_open` carrier, a durable operator-visible disposition, and a total first-action table. But the persisted shape cannot perform the claimed byte-identical re-emission. `resume_snapshot` stores the settlement-manifest **digest**, log path, and initial disposition, while `turn_open` carries the manifest itself. A digest cannot reconstruct the manifest bytes, and rev8 expressly says m-10 never recomputes them. Persist the canonical manifest bytes in the turn row or an immutable snapshot row/blob addressed by the digest; then derive every initial/replayed `turn_open` from that committed object. The digest may verify the bytes but cannot replace them.

The disposition timeline is also circular. The continuation-admission transaction stores an "initial `resume_disposition`", but only m-9 can inspect its private log after receiving `turn_open`, and `:217-220` says m-9 subsequently reports whether it is `resumable` or `degraded`. At admission m-10 cannot truthfully choose either value. The durable state therefore needs an explicit pre-report state such as `PENDING`, followed by one idempotent transition to `RESUMABLE` or `DEGRADED`; the immutable manifest/path snapshot and mutable disposition should not be conflated if byte-identical admission replay is required.

Finally, "reports before any governed effect" does not prove the report committed. The typed report has no durable receipt, yet no-work-before-disposition depends on m-10's committed state. The decomposition must require m-9 to wait for an m-10 post-commit receipt (same receipt on equivalent replay; conflict/stale handling delegated but required) before any provider attempt, tool effect, or conductor verb in the continuation. Crash-before-report, crash-after-report-before-commit, and crash-after-commit-before-receipt must all recover to the same durable state without permitting work.

Required correction: persist canonical manifest bytes behind the digest; split immutable admission snapshot from `PENDING -> RESUMABLE|DEGRADED` status; and make a post-commit disposition receipt the no-work gate. The pairs own the exact frame/replay implementation under F73 after these seam invariants are fixed.

## Closed in rev8

- **F105-D1 at decomposition grain:** CLOSED, with internal mechanism/detail deferred to the owning pair DESIGNs and the cross-owner branch clarified above.
- **F105-D3 first-action table:** CLOSED. Clean positive continuation may open the next normal attempt; uncertain provider state cannot auto-resend and requires the frozen user-requested path; uncertain tool/degraded branches remain visible and non-silent.
- **F106:** CLOSED. `xit-dur-2` is a required Durability sub-fixture, the gate consistently remains six legs, and the positive/degraded expectations are structured for the later hashed fixture artifact.
- F101/F102/F103/F104, F106 crash counter and handoff, K6 custody, source-specific UNKNOWN, G-2 bound, H-12, and every prior accepted closure remain closed.

## Gate disposition

- Proposed stage-6 amendment rev8 `9d5e8a34...`: REVISE on the two master-level seams above; not ready for operator re-scope ratification.
- The grain boundary itself is accepted; no operator arbitration is requested on it.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r9 over new amendment bytes that: (1) order the valid-prefix marker and provider content-ready evidence before resume settlement, with source-turn identity and a total provider-state partition; (2) persist the actual canonical settlement-manifest snapshot, introduce the pre-report disposition state, and gate all continuation work on a durable disposition receipt; and (3) align D1 mechanism ownership with whichever local-versus-cross-owner branch the pair later selects. Do not author the delegated record/lock/segment/rotation internals in the amendment. Preserve every closure above.

## Verification

- Target SHA-256: `b53c0359f8256b1d80e6fb1e0ce028dc7da61f9e042b760500aef8985ce58667`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1886.
- Amendment rev8 SHA-256: `9d5e8a34249e56f350d796dc8d4ad0aba24c73c0e341ba9fbbc3986ffabe57e3`; r7 reviewer parent SHA-256: `7c5de47165f529e9bac1327a87f61c7b046fb7e3a0d4518171d5750dbe74b237`.
- Relevant frozen bases recompute unchanged: m-8 provider contract `4b670a79...`; m-9 lifecycle `4d3bd14e...`; m-9 worker `cb7ff970...`; m-10 seam `d2ce9831...`; m-10 control plane `6fd1d655...`; MVP amendment `2f75f2a1...`.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds only the remaining cross-owner provider-settlement and resume-snapshot/disposition seams, then returns amendment rev9 for decomposition review r9; operator re-scope ratification remains held.
