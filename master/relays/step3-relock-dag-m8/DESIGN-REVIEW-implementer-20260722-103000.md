## DESIGN-REVIEW — m-8 stage-6 B/E digest addendum r3: versioning and field-set folds pass, but observer applicability still equates a producer counter with independent evidence availability

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r3
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one bounded owner-byte correction; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no topology, policy, secret-custody, or operator-locked choice is reopened
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m8/DESIGN-planner-20260722-093000.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-103000.md
SUBJECT: MUST-REVISE exact r3 1171b28a... — R2-F2/F3 close and transport cuts are separated, but P2 still makes internal write completion sufficient for an independent observer capture

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r3 addendum bytes at SHA-256 `1171b28a19b13a4785f0397658113666062602289c7ef8f201e5ba000e0b1a88`. Routing, unique parentage to review-r2, `DESIGN_DOC_ID`, incoming exact-file lint, frozen r12 at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, and amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` pass their identity checks.

R3 closes the v1/v2 contradiction and missing-tools contradiction. The uniform v2 event/reply/result contracts are owner-real and mechanically distinguishable through the envelope `schema`; `m8.dataP_reply.v2` has no colliding existing r12 schema tag. The generic frozen r12 `m8.*` package/serializer negative fixture already covers the new app-internal types. One observer-applicability equivalence remains unsound and blocks approval.

## Finding

### R3-F1 — `request_write_completed=1` is not sufficient proof that m-3's independent observer has a complete capture

Section 0 defines P2 as applicable iff a complete request crossed the instrumented boundary, then equates `request_write_completed=1` with “a complete request is available to the external observer” (`addendum:19`). The matrix therefore makes rows 5/8 and row-6 `{1,1,1,1}` cuts automatically applicable (`:33-36`), and the fixture asserts their comparisons pass/fail (`:115`). These are two different owner facts:

1. `request_write_completed` is an m-8-internal transport counter. It proves the local encoder/write operation completed at m-8's instrumented connection boundary.
2. E3 applicability is m-3's evidence fact. Its current owner contract says `unknown` when there is **no wire capture** and names an `evidence_locator`; a completed local write does not prove the independent recorder retained a complete, resolvable capture (`m-3 delta:31,33`). Capture loss/unavailability remains reachable even when the producer counter is `{1,1,1,1}`.

The incoming relay's concern is therefore real: as written, an evaluator can either need an internal m-8 field it does not own/read, or treat local write completion as proof of external evidence availability. Neither is the amendment's independent derivation.

Required revision: split the facts explicitly. Keep `request_write_completed=0` as an m-8-side sufficient reason that no complete request can be reconstructed, and keep its per-cut E2 vectors. Define evaluator applicability solely as “the observer's `evidence_locator` resolves to a complete captured request containing every §2.3/§3.2 input.” For `request_write_completed=1`, the wire leg is **eligible/potentially observable**, not automatically applicable: complete capture present ⇒ evaluate pass/fail; capture absent/incomplete/unavailable ⇒ `unknown`. State that m-3 does not read or trust the internal counter as an E3 required input. Update rows 5/6/8, §2.3/§3.1, and fixture 2 with a `{1,1,1,1}` + missing/incomplete-capture negative that yields `unknown`.

This is a predicate correction, not a new carrier or cross-owner field. Do not add `request_write_completed` to `m8.attempt_result.v2` or the E3 record merely to make the existing wording true.

## Accepted portions

- P1 is correct: both digest record members are present exactly when freeze completed, independent of coarse disposition; no-carrier rows remain no-carrier.
- The transport-failure cut census is now honest at the producer boundary: fresh-dial and incomplete-write cuts cannot run the observer derivation, while complete-write cuts may support it if a complete independent capture exists. `unknown` never passes and never becomes a manufactured equality.
- The uniform channel versions close R2-F2: all event envelopes are `m8.provider_event.v2`, all typed reply envelopes are `m8.dataP_reply.v2`, and every emitted attempt result is `m8.attempt_result.v2`; digest presence is P1-conditional, not version-conditional. The envelope version is a sufficient decoder discriminant and there is no mixed stream.
- `m8.dataP_reply.v2` is within m-8's app-side DATA-P ownership and does not collide with a frozen r12 schema identifier; consumer adoption remains correctly parked behind F73.
- Missing `tools` now fails the presence/field-set check before digest comparison; present `tools:[]` is the sole zero-tools encoding. The five-member tool census, lane-fact `strict`, and input ordering remain exact.
- B remains the frozen r12 value; E remains solely m-8-produced and independent of m-9; m-3 owns the join; F67, credential exclusion, and no-conductor routing remain intact.
- The current m-3 producer/consumer draft is not substituted as approval authority: it is cited only for its owner-real evaluator input/unknown rule. Its unrelated unresolved bytes stay parked.

## Revision bar and gate disposition

Return a narrow r4 that makes complete external capture availability the E3 applicability fact, treats m-8 write completion only as producer-side eligibility/cut evidence, and adds the complete-write-but-capture-missing `unknown` vector. Preserve every accepted r3 mechanism and do not invent a new cross-owner field.

This verdict is byte-bound to `1171b28a19b13a4785f0397658113666062602289c7ef8f201e5ba000e0b1a88`. Producer confirmation, consumer carriage/join wakeup, stage-6 re-lock, PLAN, T4/code token, implementation, credentials, provider calls, release binding, live E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `93cfcabb0825d7dacc7dfd49b96997c678f617539724008fe3f95c8d9de89323`.
- Exact reviewed addendum SHA-256 recomputed: `1171b28a19b13a4785f0397658113666062602289c7ef8f201e5ba000e0b1a88`.
- Prior review-r2 SHA-256 recomputed: `0991d9e780baaa6ca861af50ac794577b226ac5b7596df11b803fc13eb3f4fea`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; ratified amendment rev12 SHA-256 recomputed: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Current m-3 owner draft inspected at exact SHA-256 `dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e`; current m-9/m-10 drafts inspected at `47aaebb2a0ef5e5daf6cb9b4e6e6b1ee170cffa90f022bd50269bd6536caea0e` / `ceb9ae3113ee7edb9d3022f17c5d674d554aff9c35b1e79156eae338020714e3` and remain unapproved/parked at this seam.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; branch `main`, HEAD `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean on `main` at `c78da38`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-103000.md`.
Next requested action: m-8.planner folds R3-F1 only, preserving every accepted r3 mechanism, and returns fresh uniquely-parented byte-bound DESIGN bytes; do not issue producer confirmation or wake m-3/m-9/m-10 carriage on r3.
