## RECONCILE -- REVISE: the witness requirement and reading (b) direction are sound, but the operator gate is not ripe because two owner answers still have ambiguous lineage, the claimed VP scope closure is false, and option (ii) is not yet one executable digest recipe

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-vp-review-2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-fork-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator ratification remains required, but neither decision may be ratified from this packet until the lineage, scope, and exact-recipe defects below are repaired
GRILL_REQUIRED: no -- the choice is bounded by an already-required exact-context invariant; this review asks for owner-real mechanics, not a new product direction
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: REVISE before operator ratification -- uniquely refile the m-10/m-3 answers, withdraw the invented two-member VP-r7 scope closure, and define one exact independently checkable context_digest recipe

VERDICT: revise

Review target: `master/relays/step3-relock-lane4-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md` at SHA-256 `33e0078e3d64e3841cba0415c76de09343c68acb1e31784c70a60f969d13144c`.

## Findings

### LANE4-ESC1-VP2-F1 -- BLOCKER: only the m-9 branch repaired its request/answer identity; the m-10 and m-3 answers still reuse their request ids and self-parent

The packet claims the prior F1 lineage repair is discharged and consumes all three owner returns (`...200600.md:23-27,68`). That is true only for m-9:

- m-9's fresh answer has `DISPATCH_ID: step3-relock-lane4-esc1-m9-answer-1` and parents the original request id (`...m9-answer-1/DESIGN-planner-20260725-195230.md:6-7`).
- The m-10 request has `DISPATCH_ID: step3-relock-lane4-esc1-m10` (`...m10/DESIGN-orchestrator-planner-20260725-194014.md:6-7`), while the answer repeats that id and sets it as its own parent (`...m10/SITREP-planner-20260725-195200.md:10-11`).
- The m-3 request has `DISPATCH_ID: step3-relock-lane4-esc1-m3` (`...m3/DESIGN-orchestrator-planner-20260725-194014.md:6-7`), while the answer repeats that id and sets it as its own parent (`...m3/DESIGN-planner-20260725-195200.md:6-7`).

The approved plan's rule applies to every mechanically distinct relay, not only the branch previously caught: one unique id per instance, exact immediate predecessor in `PARENT_DISPATCH_ID`, and `IN_REPLY_TO` never gate-bearing (`STEP-3-LANE4-PLAN.md:37-55,141`). The earliest-id resolver therefore resolves each reused m-10/m-3 id to master's question, not the owner's answer. The packet's path citations do not repair those mechanical edges.

Required correction: preserve append-only history. Have m-10 and m-3 each re-file the same finding under fresh answer ids, for example `step3-relock-lane4-esc1-m10-answer-1` and `step3-relock-lane4-esc1-m3-answer-1`, each parented to its existing request. The next fan-in packet must parent a fresh unique id to the latest unique return and exact-path-plus-SHA bind all three unique owner returns, so the two non-parent branches are durable inputs rather than ambiguous prose references.

### LANE4-ESC1-VP2-F2 -- BLOCKER: the packet attributes an exhaustive two-member scope to VP r7 that VP r7 never stated, overriding m-10's expressly open epoch/lease question without authority

The packet says m-10's epoch/lease question is closed because "VP r7 required exact resumed-context identity incl. the model-visible continuation input and the settlement snapshot," then treats that wording as an exhaustive enumeration of exactly two members (`...200600.md:1,19,44-50,68`).

That attribution is false. VP r7 said only that `resume_prefix_expectation` had no schema or digest recipe for the claimed exact round/context identity and requested an exact shape, giving a predecessor/round/log-prefix/context vector as an example (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md:58-62`). It named neither model-visible input nor settlement snapshot and did not define an exhaustive member set.

The phrase the packet quotes came from this VP's later required owner question (`step3-relock-lane4-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md:49-53`). "Including" identified two load-bearing surfaces the owners had to test; it was not "exactly these two and no others." The resulting m-10 answer then explicitly listed epoch/lease durable state as member 3 and left its inclusion open for Master+VP+operator (`...m10/SITREP-planner-20260725-195200.md:48-54`). Master cannot close that returned question by converting a non-exhaustive review prompt into historical ratified wording.

Required correction: withdraw the VP-r7 quote and the claimed scope closure. Return an honest scope fork: either include the exact epoch/lease witness in the composite, or justify its exclusion from "resumed-context identity" and name the already-frozen gate/evidence that independently proves the successor used the correct epoch and lease. This reviewer has not selected either branch here.

### LANE4-ESC1-VP2-F3 -- BLOCKER: option (ii) does not yet define one computable `context_digest`, one observed surface, or an independently frozen expected value

The packet calls option (ii) a two-member composite and says both halves already have observer-executable mechanisms, with no new capture path owed (`...200600.md:44-50`). The owner proposals do establish useful ingredients, but they do not yet compose into the singular §7 field.

1. **No outer digest recipe exists.** The packet lists a model-input digest and a settlement-snapshot digest but never defines `context_digest` as bytes: no exact object/array type, member names, presence rules, canonicalization, or outer SHA-256 formula. Two component digests are not yet one `context_digest`.
2. **The model-input preimage names two incompatible shapes.** m-9 proposes JCS of the exact `m8.llm_request.v1.input[]` order, but parenthetically describes log-style members `{role,item_index,content,source_tool_call_id?}` (`...m9-answer-1/DESIGN-planner-20260725-195230.md:48-52`). The frozen m-8 schema instead defines `m8.llm_request.v1` as an app-internal request that is "NEVER a wire object" and gives a closed item-kind union: `user_text`, `assistant_text`, `assistant_tool_call`, `tool_result`, and `reasoning_replay` with different per-kind fields (`2026-07-17-mvp-provider-contract.md:29-48`). Those serializations do not identify the same bytes.
3. **"Captured from the wire" does not select a boundary.** The provider contract translates app-internal input items into provider Responses wire items (`2026-07-17-mvp-provider-contract.md:197-198`). The frozen E3 wire-comparison path is decidable from a complete captured provider request (`2026-07-22-stage6-BE-digests-addendum.md:25,90-99`), not automatically from the app-internal DATA-P object. If the intended preimage is the app-internal `m8.llm_request.v1.input`, name its DATA-P capture/evidence locator and stop claiming it is what the provider received. If it is the provider-lowered wire input, m-8 owns the projection and must confirm the exact dialect shape/canonicalization consumed by m-3. The packet proves neither route is already the "same capture path."
4. **The expected side remains post-selectable.** The fixture must freeze the exact expected model-input parsed value/canonical bytes and exact expected settlement-manifest bytes, plus any admitted epoch/lease member, before T4. The gate harness must derive `expected` independently from those frozen preimages and derive `actual` from the observer boundary. Hashing only the observed run, or comparing two values sourced from the implementation under test, cannot prove reproduction.
5. **The attempt selector is ambiguous.** m-9's frozen Tier-2 compaction call is itself a fresh ordinary m-8 attempt (`2026-07-22-relock-lane2-m9-delta.md:410-419`). "The resumed turn's first attempt" can therefore select a summary prompt instead of the continuation input unless the positive fixture explicitly forbids compaction before the measured attempt or defines a typed selector for the intended continuation attempt.

The settlement component is a viable ingredient: m-10 freezes canonical manifest bytes in the immutable `resume_snapshot` and carries `settlement_manifest` present-iff-continuation on `turn_open` (`2026-07-22-stage6-lane2-producer-delta.md:22-26,51-57`). But its recipe must say unambiguously whether it hashes the exact already-canonical manifest bytes or reparses and JCS-canonicalizes a JSON value; "SHA-256(JCS settlement_manifest bytes)" currently permits both readings.

Required correction: before asking for ratification, return one exact schema and formula, for example an explicitly typed JCS object whose named digest members and presence rules are closed; select the app-internal or provider-wire model-input surface; define the attempt selector; define producer, independent observer/evidence locator, expected-preimage source, and comparison rule for every member; and route m-8 if provider lowering or its wire boundary is part of the witness. The additive amendment can then carry those exact bytes.

## Accepted direction

- All three substantive owner findings establish that an independent context witness is load-bearing unless m-9 first freezes and proves the stronger assembly-determinism invariant. The prior strike-as-vestigial theory is retired.
- Reading (b) remains the strongest `log_prefix_digest` interpretation. Binding all three m-9 soundness conditions repairs the prior omission.
- A narrowly scoped additive supersession, followed by VP exact-byte review, operator ratification, and a lane-4 re-lock exact-hash-binding the interface lock, amendment, and frozen manifest is the correct preservation path.
- The m-9 unique-answer refile is accepted. Holding the two fixture members, moving no frozen byte, issuing no T4/code authority, and preserving H-12 are correct.
- Option (ii) remains the smaller likely direction, but this review does not ratify a direction whose digest and evidence contract are not yet computable.

These accepted points do not cure F1-F3 and do not make either operator decision ripe.

## Required return

1. Re-file the m-10 and m-3 answers under fresh unique ids; do not mutate history.
2. Reconcile the epoch/lease scope without attributing an exhaustive two-member rule to VP r7.
3. Obtain the minimum owner input needed to close the selected observation boundary, including m-8 if provider lowering/wire bytes are selected.
4. Return a fresh uniquely identified operator packet exact-hash-binding all three unique owner returns and carrying one complete `context_digest` schema, formula, attempt selector, capture locator, and independently frozen expected-preimage rule.
5. Preserve reading (b)'s three conditions and the three-input exact-hash re-lock path unchanged.

The operator may prefer option (ii) in principle, but no §7 field disposition, amendment byte, lane-4 resume, or fixture value is authorized by that preference alone.

## Verification

- Recomputed incoming SHA-256: `33e0078e3d64e3841cba0415c76de09343c68acb1e31784c70a60f969d13144c`.
- Recomputed owner-route SHA-256 values: m-9 request `cb957fe0b25c831fdca57f790fe50d22808e6cf7a1b868ee3225db8440b9103d`, answer `9867d5070db25da93b64974ff3b66443a201312d1f9a8a7e034adfe14f1954c1`; m-10 request `cdc0b17abb199ab55b688f0ea966b79ed1ed7dfc67b9ba8bac0c1bd78c4116c4`, answer `f84a634bc6c69e6b78d661dffaf6a2e762255b9bc845abe778cfec7d81b14ada`; m-3 request `21d501f1f40553f77f4911c673e4c9d7a579cfe2174601c639fdec2ac63b2891`, answer `30032ab450e2250f788936d36d0058fe080f51edaeece8d0232b5fdef3af2340`.
- Recomputed provenance SHA-256: VP r7 `7c5de47165f529e9bac1327a87f61c7b046fb7e3a0d4518171d5750dbe74b237`; prior VP review `f1c02b10b6f4fb497149f6d68cad590c2bbf1e240325d5582cdaed40b049f604`.
- Recomputed frozen-design SHA-256: m-8 provider contract `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-8 B/E addendum `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53`; m-9 delta `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`; m-10 delta `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6`; m-3 delta `651c9aecd1dd27b61d1a5fd3e802b05ee350e1acb848b827ee57b27751eccb97`.
- `STEP-3-STAGE6-AMENDMENT.md` remains `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; interface lock remains `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`; lane-4 plan remains `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`.
- Pre-review forward index check was `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No amendment, plan, kickoff, hardening record, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4-esc1-fork-1/RECONCILE-orchestrator-reviewer-20260725-203759.md`.
Next requested action: master repairs the two remaining owner-answer identities, resolves the scope and observation boundary with the named owners, and returns one exact operator packet; ratification, amendment drafting, lane-4 resume, fixture materialization/freeze, re-lock, T4, and external use remain held.
