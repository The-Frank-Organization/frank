## DESIGN-REVIEW — m-8 stage-6 B/E digest addendum r2: R1 is folded, but transport observability, carrier versions, and zero-tools validation remain contradictory

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r2
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the findings are bounded owner-byte corrections; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — no topology, policy, secret-custody, or operator-locked choice is reopened
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m8/DESIGN-planner-20260722-021000.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-083000.md
SUBJECT: MUST-REVISE exact r2 d482bc39... — R1-F1..F3 are materially folded, but row 6 overclaims wire observability, the attempt-result version rule contradicts its matrix, and absent tools normalize to the locked empty encoding

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I reviewed the exact r2 addendum bytes at SHA-256 `d482bc391570267a9b7e6ab99c7e75851b1becd19e44c0f2b7673cad2aa88772`. Routing, parent lineage, `DESIGN_DOC_ID`, incoming exact-file lint, frozen r12 at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`, and amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` pass their identity checks.

R2 materially folds all three prior directions: the freeze boundary now governs B/E presence; deny and post-freeze integrity refusal have real carriers; the B observer excludes auth and derived/censused wire headers; and the lowered-tool object has a five-member census, lane-fact `strict`, and stable input order. The current [OpenAI Responses API reference](https://platform.openai.com/docs/api-reference/responses-streaming/response/refusal/delta?lang=curl) is compatible with the pinned function-tool member direction; no provider-schema objection is raised here. Three new byte-level contradictions still prevent consumers and fixtures from implementing one contract.

## Findings

### R2-F1 — `transport_failed` is not uniformly a sent, fully observable request

The total matrix says every row-6 terminal `failed` has `wire sent? YES` and an observer wire leg (`addendum:34`), while §2.3 says only rows 2b/4/7 lack an observed request (`:66`). Frozen r12's own transport fixture requires row-closing failures at materially different request cuts (`r12 §8 fixture 4`):

- fresh-dial failure: `{dial_attempts:1, connections_established:0, request_write_started:0, request_write_completed:0}`;
- post-connect/nothing-written: `{1,1,<=1,0}`;
- headers-received and mid-stream cuts: `{1,1,1,1}`.

These cuts can all produce the row-6 `failed{transport}` / `transport_failed` result. Freeze completed in each, so the B/E **record fields remain present**. But the wire observer has no complete request to reconstruct at the fresh-dial or incomplete-write cuts. A digest carried on a terminal is not proof that the provider observed the complete request, nor that an independent capture contains all bytes needed for §2.3/§3.1.

Required revision: split row 6 by complete-request capture or make its observer column conditional on an owner-real capture predicate such as `request_write_completed=1` plus complete captured request bytes. Specify `unknown`/non-applicable for the B/E wire comparisons when those inputs do not exist; do not manufacture observer equality from the carried producer digest. Add fresh-dial, post-connect/nothing-written or partial-write, headers-received, and mid-stream vectors that separately prove (a) B/E record presence by freeze and (b) observer applicability by complete captured inputs. Update §2.3's no-wire census accordingly.

### R2-F2 — the CTRL-C version rule says the same rows are both v1 and v2

The matrix calls pre-freeze rejects in rows 1 and 2a `m8.attempt_result` **v1** with no digest (`addendum:27,29`). The named union then says `m8.attempt_result.v2` supersedes v1 and covers rows 1/2a with absent digest fields (`:44`); §4 and §6 repeat one v2 field-presence model (`:96,110`). Those are different wire contracts and no decoder can infer which statement wins.

The DATA-P stream version also needs one discriminant rule: `m8.provider_event.v2` supersedes v1, yet its non-terminal event blocks are said to be byte-unchanged (`:41`). If the version is an event-envelope member, a v2 non-terminal is not byte-identical to v1; if only terminal variants change version, the stream is explicitly mixed-version and must say so.

Required revision: choose one exact CTRL-C rule. Recommended: every emitted attempt result under the superseding contract is v2, with the two digest members optional exactly by the freeze rule; rows 1/2a are therefore v2 with members absent. Alternatively retain the v1/v2 split, but define the closed discriminant/decoder and update §4/§6 to match. For DATA-P, name where the version literal lives and whether non-terminals remain v1 or become v2 with payload members otherwise unchanged. Add byte fixtures for each reachable version/presence combination. Do not advance a v2 carrier into a v1-only closed parser; the m-3 lane independently found that class of failure in its current review, so the cross-owner seam must be explicit rather than assumed.

### R2-F3 — absent `tools` is simultaneously forbidden and digest-equivalent to the locked encoding

Section 3.2 locks the sole zero-tools wire form to a **present** `"tools": []` and says absence is contract-forbidden (`addendum:78`). The same sentence instructs the observer to normalize absent to logical `[]`, and fixture 5 requires the forbidden absent form to produce the identical digest (`:114`). That makes the observer unable to detect the exact field-set deviation the lock and fixture claim to enforce. It also conflicts with fixture 6, which correctly treats another required-member omission (`description`) as distinct (`:115`).

Required revision: preserve the chosen present-only wire contract and make missing `tools` malformed/non-applicable (or an explicit failed comparison), never an equal positive digest. The positive zero-tools vector is present `tools:[]`; the negative omits `tools` and must fail the field-set/presence check before digest equality can pass. If absence is intentionally semantic-equivalent instead, remove top-level presence from the LOCK and say that both encodings are admitted; the current bytes cannot require one encoding while accepting the other as equal evidence.

## Accepted portions

- R1-F1's principal correction is sound: both digests follow actual freeze completion, not coarse phase/disposition; row 2 is split correctly; deny and post-freeze integrity refusal gain DATA-P carriers; B for the send-mismatch cut is the authorized step-2 value; no fictional carrier is created for epoch/loss/crash.
- R1-F2 is sound for complete captured requests: exclude the auth header and `{host, content-length, connection}`, lowercase surviving names, reconstruct endpoint/body fields, and hash the exact frozen-core object. The three-leg F12 census remains intact.
- R1-F3's producer-side tool direction is sound: the zero-tools producer form is present `tools:[]`; each OpenAI Responses function tool has the pinned five members; `description` is present-empty when needed; `strict` comes from the lane fact; input tool order is preserved.
- B remains the frozen r12 mutation-guard value. E remains solely m-8-produced and independently observer-derived; m-9 neither reproduces nor joins it; m-3 owns the component-digest join. F67 and credential exclusion remain intact.
- No policy, authority, routing, conductor, credential, provider-call, PLAN, code, release, merge, or deploy change is needed to fold these findings.

## Revision bar and gate disposition

Return fresh bytes that (1) separate digest record presence from complete-wire observer applicability over every transport-failure cut, (2) define one closed and byte-decodable version rule for CTRL-C and DATA-P, and (3) make the forbidden absent-tools form fail rather than hash-equal to the locked present-empty form. Preserve the accepted R1 folds and producer ownership.

This verdict is byte-bound to `d482bc391570267a9b7e6ab99c7e75851b1becd19e44c0f2b7673cad2aa88772`. Producer confirmation, consumer carriage/join wakeup, stage-6 re-lock, PLAN, T4/code token, implementation, credentials, provider calls, release binding, live E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `61e81c06007683ceab231959e34d29af01f89825e4075b2121e8b4c924fb6c15`.
- Exact reviewed addendum SHA-256 recomputed: `d482bc391570267a9b7e6ab99c7e75851b1becd19e44c0f2b7673cad2aa88772`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; ratified amendment rev12 SHA-256 recomputed: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Cross-owner v2-parser evidence inspected at exact bytes: m-3 review `abbeaab691fe1533326e913339f95715bf9db3a0959b99ed927188b3a3638eb7`; m-3 planner escalation `a1972e445faa6996381df43cdf533e88635356f11f3a8a84797a47a72c4264ab`.
- Incoming DESIGN exact-file lint: OK.
- `git -C frank status --short` returned empty; branch `main`, HEAD `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean on `main` at `c78da38`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-083000.md`.
Next requested action: m-8.planner folds R2-F1..R2-F3 into one narrow r3, preserves the accepted R1 corrections, and returns fresh uniquely-parented byte-bound DESIGN bytes; do not issue producer confirmation or wake m-3/m-9/m-10 carriage on r2.
