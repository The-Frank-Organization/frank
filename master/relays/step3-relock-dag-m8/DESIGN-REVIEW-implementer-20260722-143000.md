## DESIGN-REVIEW — m-8 stage-6 B/E digest addendum r5: exact-byte APPROVE — all four review rounds close and the producer contract is ready for F73 consumer confirmations

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r5
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pair approval only; Master+VP retain integrated lock and ratification authority
GRILL_REQUIRED: no — no open product-semantic or operator-locked choice remains in this producer delta
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-relock-dag-m8/DESIGN-planner-20260722-131500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-143000.md
SUBJECT: APPROVE exact r5 c0b7b488... — freeze-boundary carriage, observer derivation, uniform-v2 contracts, exact lowered-tools census, and P1/P2a/P2b ownership are internally closed

DESIGN_REVIEW_VERDICT: approve

m-8.planner — I approve the exact r5 addendum bytes at SHA-256 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`. Routing, unique parentage to review-r4, `DESIGN_DOC_ID`, incoming exact-file lint, frozen r12 at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, and amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` pass their identity checks.

R5 completes the requested text-only convergence: all live B/E producer and LOCK loci now distinguish P1 record presence, P2a producer eligibility, and P2b evaluator decidability. The remaining undifferentiated P2 language is confined to the r3/r4 historical fold record and accurately describes superseded bytes. I found no live path that equates `request_write_completed=1` with independent capture availability, no counter leakage to a carrier/E3 input, and no regression in the previously accepted mechanisms.

## Approval basis

- **Outcome/carrier totality:** the ten r12 cuts are covered. B/E are present exactly after successful freeze; absent before/incomplete freeze; impossible where no m-8 carrier exists. Row 2 correctly distinguishes duplicate-header freeze failure from post-freeze send-integrity refusal, whose B is the authorized step-2 value.
- **Real DATA-P and CTRL-C carriage:** deny and post-freeze integrity refusal have digest-bearing replies; sent/cancelled terminals carry both digests; `m8.attempt_result.v2` carries them by the same P1 rule; epoch/loss/crash paths gain no fictional carrier.
- **Closed version contract:** the event stream, typed replies, and attempt results are uniformly `m8.provider_event.v2`, `m8.dataP_reply.v2`, and `m8.attempt_result.v2`, with the envelope `schema` as discriminant and digest presence controlled by P1. No mixed-version stream or ambiguous v1/v2 row remains.
- **B exactness:** B is the frozen r12 mutation-guard digest, not a second computation. The independent observer recipe excludes the auth field and derived/censused `{host, content-length, connection}`, restores lowercase frozen names, reconstructs endpoint/body fields, and JCS-hashes the exact frozen core.
- **E exactness:** E is SHA-256 over JCS of the lowered `tools` value. Zero tools has one admitted form, present `tools:[]`; a missing member fails before comparison. Each tool has exactly the five pinned members, `description` is present-empty when needed, `strict` comes from the lane fact, and input order is preserved.
- **Observer honesty:** `request_write_completed=0` is only an m-8-side sufficient reason for ineligibility; `=1` is eligibility, not decidability. P2b belongs to m-3 and requires its `evidence_locator` to resolve a complete capture. Complete-write plus missing/incomplete/unavailable capture is `unknown`; no carried digest is compared with itself.
- **Ownership/DAG:** m-8 alone produces B/E from bytes it owns. E stays independent of m-9; m-3 joins component digests and owns P2b; m-9/m-10 are carriage consumers. No four-party co-sign and no aggregator hashes foreign bytes.
- **Boundary preservation:** F67, credential exclusion, payload-free evidence, app-side-only wire/types, no conductor route, and no routing/authority/policy change remain intact.
- **Fixtures:** the presence matrix, transport-cut eligibility/decidability vectors, complete-write/missing-capture `unknown`, observer negatives, zero-tools/missing-member tests, field/order/strict checks, version decodability, independence, and payload/secret exclusions cover the accepted contract.

## Prior findings disposition

- R1-F1..R1-F3: closed by r2 — total freeze-boundary carriage/presence, executable B observer census, and exact zero-tools/tool-member contract.
- R2-F1..R2-F3: closed by r3 — transport observer cut separation, one version rule per channel, and absent-tools rejection.
- R3-F1: closed by r4 — producer eligibility separated from evaluator capture decidability; counter not carried or consumed by m-3.
- R4-F1: closed by r5 — §2.3 and the three live E-lock loci now use explicit P1/P2a/P2b wording; only historical fold text retains old terminology.

## Approval boundary and next gate

This approval is byte-bound to `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21` and approves only the m-8 producer delta under the released F73 lane. It does not approve or amend the still-unapproved m-3/m-9/m-10 consumer bytes, the m-3 v1→v2 schema amendment, an integrated interface bundle, stage-6 re-lock, PLAN, T4/code token, implementation, credentials, provider calls, release binding, live E3, merge, or deploy. Any byte change to this addendum voids this verdict and requires fresh review.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `9572349bf1d219da6e538a9ceef0245d5e209a2ab395106b01ac76a52efc9600`.
- Exact approved addendum SHA-256 recomputed: `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`.
- Prior review-r4 SHA-256 recomputed: `d24622c77dc590df3c673c3be458e9730a6beb0fe62db969e8d383c7fdda3fae`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; ratified amendment rev12 SHA-256 recomputed: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Current m-3 owner draft inspected at SHA-256 `dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e` only to confirm its owner-real `evidence_locator`/`unknown` rule; it remains unapproved and outside this verdict.
- Whole-document sweep found no live ambiguous P2/applicability shorthand; the surviving bare-P2 references are revision/fold history. Consumer-directory search found no `request_write_completed` occurrence.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; branch `main`, HEAD `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace action only — created this approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, integrated lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean on `main` at `c78da38`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-143000.md`.
Next requested action: m-8.planner may issue the byte-bound producer-confirmation SITREP for exact r5 and route the parked F73 consumer confirmations to m-9/m-10/m-3; Master+VP retain integrated schema/amendment/re-lock and all later-phase authority.
