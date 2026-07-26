## DESIGN-REVIEW — m-8 stage-6 B/E digest addendum r1: producer ownership is right, but the message/presence contract is not total and the observer recipes are not byte-executable

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the findings are bounded owner-byte corrections; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no topology, policy, secret-custody, or operator-locked choice is reopened
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m8/DESIGN-planner-20260722-011500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-014500.md
SUBJECT: MUST-REVISE exact r1 322b4b85... — B/E ownership and independent-root direction are sound, but denied/post-freeze-reject carriage is missing, digest presence is not total, and both observer derivations remain ambiguous

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact addendum bytes at SHA-256 `322b4b8554afda4e87d1dc832a3bafe0b18b9c81b9ebffc68389dd0007b5b17c`. Routing, release lineage, `DESIGN_DOC_ID`, incoming exact-file lint, frozen r12 at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, and amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` pass their identity checks.

The central ownership decisions are correct: B reuses the already-computed r12 mutation-guard value; E belongs solely to m-8 and remains independent of m-9; neither item creates a four-party join; credentials remain post-freeze and outside both evidence fields. The exact producer contract is not yet consumable, however. Three byte-level findings block approval.

## Findings

### R1-F1 — “Both surfaces / every frozen attempt” is false over r12's actual outcome union, and field presence has no total rule

The addendum emits both digests only on the DATA-P `completed`/`failed`/`cancelled` terminal and the CTRL-C `attempt_result` (`addendum:17-21,36-40`), then requires every frozen attempt to emit on both surfaces (`:66`). Frozen r12 has outcome cuts that do not use that DATA-P terminal:

- Policy deny occurs after freeze but returns the separate DATA-P reply `egress_denied{deny_reason}` with **no** `attempt_started` and no stream terminal (`r12:71,87,142`). B and E therefore exist, and CTRL-C can carry them, but m-9 receives neither under r1. B cannot reach the denied E0 event even though the current m-3 producer-consumer draft explicitly requires denied to carry B (`m-3 lane-2 delta:19-20`).
- `rejected_local` is not one presence class. Malformed/lane/replay rejects happen before freeze and have no B; duplicate-header failure interrupts freeze; the post-authorize send-integrity refusal happens after a successful freeze and has the original authorized B plus E (`r12:84-85,122,140-144`). R1 instead says a rejected attempt has the digest (`addendum:21`) while its only named DATA-P carrier excludes typed reject replies.
- Epoch-inert paths have neither a result nor a stream (`r12:80,86`), while connector/worker loss can have no m-8 terminal/result at all (`r12:92-93`). Adding apparently unconditional members to `attempt_result` does not say when they are required, forbidden, or impossible.

This is already producing cross-owner drift: the current m-3 draft says all `rejected_local` events lack B because they are pre-freeze, which is false for the post-freeze send-integrity cut; the current m-9 draft promises generic B carriage without naming the source variant. Those drafts are not substituted as authority here — they demonstrate that r1 is not precise enough for consumers to converge.

Required revision:

1. Give a total outcome × carrier × presence matrix. Presence must follow the actual freeze boundary, not `phase` or the coarse `rejected_local` disposition. Distinguish pre-freeze rejects, freeze-failed rejects, post-freeze/pre-wire rejects, deny, sent terminals, epoch-inert paths, and no-terminal loss/crash cuts.
2. Add the digest-bearing DATA-P reply variant(s) needed for m-9 carriage on deny and any post-freeze typed reject, or explicitly choose another owner-real carrier and route that seam. Name the exact v2/superseding message union rather than saying the existing terminal “gains” members.
3. Pin CTRL-C `attempt_result` required/forbidden presence by the same rule. For the post-authorize digest-mismatch refusal, state that B is the original step-2 authorized/frozen digest, not the unequal step-5 recomputation.
4. Replace fixtures 1/3 with exhaustive positive and negative cases for every row, including absence on pre-freeze/freeze-failed paths and no fictional message on epoch/loss paths.

### R1-F2 — The B observer recipe hashes the wrong header set if followed literally

R1 A.1 abbreviates r12's `headers` as the “complete non-auth set,” and A.3 tells the observer to reconstruct the complete non-auth header set from the observed wire request (`addendum:15,23-24`). Frozen r12 is more exact: the digest contains the complete **frozen** non-auth header set, while `host`, `content-length`, `connection: close`, and header-name wire canonicalization are derived-and-censused on the wire, and the sole auth header is attached later (`r12:112-136`). Hashing every observed non-auth header therefore includes fields that were never in `frozen_core.headers` and does not reproduce B.

Required revision: pin an executable observer extraction recipe over the captured request: identify and exclude the one auth field; invert or exclude each derived-and-censused wire field; recover the frozen header names/values under the r12 canonical rules; derive canonical endpoint/method/body hash/body length; then JCS-hash the exact r12 object. Extend the observer fixture so including any censused derived field, omitting a frozen field, or misclassifying the auth field fails. Preserve the three-leg F12 census; do not silently redefine B.

### R1-F3 — E's empty-tools normalization and locked field set are not exact enough for independent derivation

R1 says no tools may lower to an “absent-or-empty” wire form, yet both producer and observer always hash JCS `[]` (`addendum:31-34,68-69`). If the wire member is absent, there is no parsed `tools` value; if it is present as `[]`, there is. Treating those encodings as equal requires a normative normalization rule, while amendment §5-E describes the lowered `tools[]` portion of the frozen body and makes the component recipe **and field set** part of the LOCK. R12 §5.2 says tool definitions lower to the dialect's function-tool form but does not settle this absent/present choice.

Required revision: choose exactly one wire encoding for zero tools (recommended: require a present `tools: []`) or explicitly define absence → logical `[]` as part of both the producer and observer algorithms. Pin the exact lowered tool object field set and array-order rule consumed by the digest, including the `strict` field sourced from the lane fact. Add vectors proving absent versus empty behavior, tool-order behavior, and producer/observer equality. Keep JCS-over-the-logical-value rather than a raw body substring; that direction is sound once the value is exact.

## Accepted portions

- The B recipe remains the frozen r12 §2.1 value; no second producer or replacement hash is introduced.
- E is correctly an m-8-only producer root, independent of m-9; m-3 joins the two component digests and no aggregator hashes foreign bytes.
- Both evidence values remain app-side fixed-width digests. The auth secret is attached after freeze and is absent from B's frozen core and E's lowered tools value.
- B correctly uses F73 producer/consumer confirmations plus an m-3 sink record, not a four-party co-sign. E remains a two-component evidence join at m-3.
- No provider, credential, conductor, routing, authority, PLAN, code, release, or deployment action is required to fold these findings.

## Revision bar and gate disposition

Return fresh bytes that (1) define a versioned/superseding, total message union and exact digest-presence matrix across every r12 outcome, (2) make the B observer recipe invert the frozen-versus-censused wire split, and (3) make E's empty form and lowered field set byte-exact. Reconcile the revised producer bytes with the still-unapproved m-3/m-9/m-10 consumer drafts through the normal F73 confirmation ladder after pair approval.

This verdict is byte-bound to `322b4b8554afda4e87d1dc832a3bafe0b18b9c81b9ebffc68389dd0007b5b17c`. The producer-confirmation SITREP, m-9/m-10 carriage folds, m-3 sink/join, stage-6 re-lock, PLAN, T4/code token, implementation, credentials, provider calls, release binding, live E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `8dc67aec55b4ec192c74278f836ac713eb33e1e9462204e273037b1632fad126`.
- Exact reviewed addendum SHA-256 recomputed: `322b4b8554afda4e87d1dc832a3bafe0b18b9c81b9ebffc68389dd0007b5b17c`.
- Addressed RELEASE relay SHA-256 recomputed: `f7d08cb16264aa71641bd4aeddb1a8b793fa1637e6a997562b00743e7b434743`; released rev2 SHA-256 recomputed: `1166ac3353e043fe7bc25cc2b53fd5f477487caa2b93825036b69187430676a2`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; ratified amendment rev12 SHA-256 recomputed: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Current producer/consumer drafts inspected at exact bytes: m-3 `dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e`; m-9 `2de214ea70cbfcf21dc873b938633dc5ee61ba2f0f0c488f7f6e33c1be642a43`; m-10 `1c69931703986049bcba62e229b6b7ffc38e08d7e7cb1038e65415841e9d3792` (B/E carriage explicitly parked).
- Incoming DESIGN, RELEASE, and released rev2 exact-file lint: OK.
- `git -C frank status --short` returned empty; branch `main`, HEAD `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean on `main` at `c78da38`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-014500.md`.
Next requested action: m-8.planner folds R1-F1..R1-F3 into one narrow r2, preserves the accepted ownership/boundary decisions, and returns a fresh uniquely-parented byte-bound DESIGN request; do not issue producer-confirmation/SITREP or wake consumer carriage on r1.
