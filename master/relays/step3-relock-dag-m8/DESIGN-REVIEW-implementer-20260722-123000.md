## DESIGN-REVIEW — m-8 stage-6 B/E digest addendum r4: the R3 observer split closes, but three live E-lock loci still reference the superseded undifferentiated P2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r4
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one bounded stale-locus correction; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no topology, policy, secret-custody, or operator-locked choice is reopened
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m8/DESIGN-planner-20260722-111500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-123000.md
SUBJECT: MUST-REVISE exact r4 23b36d42... — P2a/P2b closes R3-F1, but the E placement, derivation, and LOCK summaries still say undefined old P2 applicability

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r4 addendum bytes at SHA-256 `23b36d423951385c94809b8d8428e67ea90c2594c872bc60e26e8005fa2a3625`. Routing, unique parentage to review-r3, `DESIGN_DOC_ID`, incoming exact-file lint, frozen r12 at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, and amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` pass their identity checks.

The load-bearing R3-F1 fold is correct: P2a is m-8-internal write eligibility only; P2b is m-3-owned evaluator decidability from a resolvable complete capture; `{1,1,1,1}` plus missing/incomplete/unavailable capture yields `unknown`; the counter is neither carried nor an E3 input. The uniform-v2, presence, field-set, digest, ownership, and credential-boundary mechanisms remain intact. One stale-summary class remains inside the live E LOCK surface and blocks an exact-byte approval.

## Finding

### R4-F1 — §3.3–§3.5 still use the superseded, now-undefined “P2 applicability” contract

R4 replaces the single P2 with two non-interchangeable facts: P2a producer eligibility and P2b evaluator decidability (`addendum:14-24`). The B recipe and matrix use those names exactly (`:28-41,62-73`). Three live normative E loci did not receive the same fold:

- §3.3 says E has “same P2 applicability” (`addendum:90-91`).
- §3.4 says the observer derives E “when P2 applicable” (`:93-94`).
- §3.5 includes undifferentiated “applicability” in the E LOCK surface (`:96-97`).

There is no standalone P2 in r4. These are not harmless history lines: §3.5 expressly declares the producer half of the amendment's E LOCK, and §3.3/§3.4 are the live placement/derivation statements a consumer follows. A reader can still map “P2 applicable” to r3's rejected meaning (`request_write_completed=1` ⇒ applicable), contradicting the correct §3.1 rule that complete-write eligibility is insufficient without a resolvable complete capture.

Required revision: replace the three live loci with the exact split already established elsewhere. §3.3 should say same P1 presence, P2a producer eligibility, and P2b evaluator decidability. §3.4 should say the observer derives E only when P2b is decidable from its complete capture (with P2a remaining producer-side eligibility, not an evaluator input). §3.5 should name P1/P2a/P2b explicitly in the LOCK summary. Then sweep the live document for any non-historical bare `P2`/“applicability” shorthand; the r3 fold-log description may remain historical because it accurately records the superseded r3 bytes.

No mechanism change, new field, cross-owner amendment, or additional fixture is required. This is the same stale-summary failure class that exact-byte review must remove before a lock surface can be approved.

## Accepted portions

- R3-F1 closes at the operative definitions, matrix, recipes, consumer response, and fixture: write completion is eligibility only; complete independent capture controls decidability; missing capture yields `unknown`; m-3 never reads the internal counter.
- P1 remains exact across all ten outcome cuts; the row-2 freeze split, deny/post-freeze-reject carriers, and step-2 authorized B value remain sound.
- The three uniform v2 channels remain closed and byte-decodable; no mixed-version stream or v1/v2 row split reappears.
- Missing `tools` fails before comparison; present `tools:[]` is the sole zero-tools form; the five-member object census, lane-fact `strict`, and input ordering remain exact.
- B is still the frozen r12 mutation-guard value. E remains solely m-8-produced and independent of m-9; m-3 owns decidability and the component-digest join. No aggregator hashes foreign bytes.
- F67, credential exclusion, payload-free digests, no-conductor routing, and producer-first/F73 parking remain intact.
- The current m-3 bytes at `dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e` support the corrected P2b rule (`evidence_locator`; `unknown` for no wire capture or carried digest) but remain unapproved and are not substituted as pair-approval authority.

## Revision bar and gate disposition

Return fresh bytes that replace the three live E-lock references to old P2/applicability with explicit P1/P2a/P2b wording, leaving the correct r4 mechanism and all accepted contracts unchanged. This should be a text-only convergence pass with no new design choice.

This verdict is byte-bound to `23b36d423951385c94809b8d8428e67ea90c2594c872bc60e26e8005fa2a3625`. Producer confirmation, consumer carriage/join wakeup, stage-6 re-lock, PLAN, T4/code token, implementation, credentials, provider calls, release binding, live E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `b5a9dc7284513eaa76a404cb53dc1201af5846733c66534bc55fa5478de74c5f`.
- Exact reviewed addendum SHA-256 recomputed: `23b36d423951385c94809b8d8428e67ea90c2594c872bc60e26e8005fa2a3625`.
- Prior review-r3 SHA-256 recomputed: `93ee71b19964e93aad734af0ed4b9e9d0b8c667b2687fbd80b174d99335986e4`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; ratified amendment rev12 SHA-256 recomputed: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Whole-document counter search found `request_write_completed` only in the m-8 addendum/r12 owner context and no occurrence in the current m-3/m-9/m-10 consumer design directories; no carrier/E3 leakage found.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; branch `main`, HEAD `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean on `main` at `c78da38`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-123000.md`.
Next requested action: m-8.planner folds R4-F1 only, preserving the complete accepted r4 mechanism, and returns fresh uniquely-parented byte-bound DESIGN bytes; do not issue producer confirmation or wake m-3/m-9/m-10 carriage on r4.
