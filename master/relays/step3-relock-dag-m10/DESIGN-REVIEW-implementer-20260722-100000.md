## DESIGN-REVIEW — rev2 MUST-REVISE: five r1 folds close, but the VOID class contradicts the locked broker seam, the frame arithmetic still underbounds real wire bytes, and `workspace_root_id` has no exact producer-to-worker contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r2
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the three findings are exact contract/completeness defects under already-ratified or pair-confirmed semantics; any desired change to those semantics must route upward as an amendment rather than be silently selected here
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260722-090000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-100000.md
SUBJECT: MUST-REVISE exact rev2 `ceb9ae31…`: reconcile orphan `VOID` with the byte-bound M10-C1 `uncertain` seam, redo the frame proof over actual canonical wire bytes and full-run `parked_unknown`, and give `workspace_root_id` an exact value recipe plus worker carrier

## Verdict

**MUST-REVISE** on exact design SHA-256 `ceb9ae3113ee7edb9d3022f17c5d674d554aff9c35b1e79156eae338020714e3`.

Rev2 materially closes r1 F2, F3, F6, F7, and F8: the unsettled m-9 shapes are honestly parked/non-normative; `resume_action` is durable and replay-compared; the five-result broker fold is total; the two missed r10 loci are replaced; and the completed pair confirmation is cited. The discriminated manifest shape also closes the wire-decoding half of r1 F1. Three load-bearing defects remain.

## M10-DAG-R2-F1 — BLOCKER: the new orphan-`VOID` mapping contradicts both §7 and the byte-bound M10-C1 seam

Rev2 §1 maps **all five** authorization-orphan `VOID` reasons, including `expired`, to `determinate_no_resume` (`:17-18`). Rev2 §7 simultaneously says issued-unconsumed becomes `VOID` and that §1 carries those rows into **`uncertain`** (`:102`). Both cannot be true.

The external authority is also exact. Broker-study rev8 `64f9136e…` §Q3.4 says issued-but-unconsumed follows `VOID` and the successor receives that identity through D2 `uncertain` (`:88-89`); its m-10 obligation repeats `issued-unconsumed ⇒ VOID` plus `uncertain` membership (`:109`). The completed m-10 confirmation `375da939…` binds the same identity into `uncertain` on `turn_open` (`:56-57`). Rev2 may not silently overturn that byte-bound pair/master seam because local F59 reasoning finds the call definitely unconsumed.

There is an additional constructibility constraint: `boundary_cut` is deliberately telemetry-never-input, so no manifest producer may use it to distinguish one `VOID/expired` row from another. The correction must therefore name a mapping decidable from canonical authorization fields alone (for example the exact `void_reason` plus canonical tool identity), make §1/§7/fixtures agree, and preserve the confirmed cut-relay `uncertain` disclosure. If the planner believes the confirmed seam is semantically wrong, route an explicit amendment to Master/VP; do not resolve the conflict pair-locally.

## M10-DAG-R2-F2 — BLOCKER: the claimed exact frame bound undercounts the actual canonical wire

The arithmetic at §4 `:50-58` is not an upper bound over the frozen `turn_open` shape:

1. **`parked_unknown` row size is impossible at 256 B.** Frozen r40 §B.2 D-4 defines each row as `{turn_id, tool_call_id, ticket_id, state, canonical_tool_name, canonical_args_digest}`. Under rev2's own 64-byte identifier assumption, the three ids plus one digest already consume 256 value bytes before `state`, tool name, member names, quotes, separators, or array punctuation. `ticket_id` is not even included in rev2's new `ID_MAX` list.
2. **Its cardinality is not proved as 160.** R40 attaches every `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT` row that exists **in the run**. Rev2 multiplies one turn's `(128 + 32)` without proving a full-run cap, and the `32` provider-attempt term is not a member of r40's tool-only disclosure shape. Across a ten-turn ancestry, the naive tool ceiling alone can be 1,280 rows unless a narrower durable invariant is added.
3. **Raw string bytes are not canonical JSON bytes.** `TASK_INPUT_MAX`, ids, and paths are bounded as raw lengths, but JCS/JSON escaping can expand quotes, backslashes, and control characters. The initial actual-frame check plus `task_input ≤ 2,850,816` does not reserve the manifest: a raw-under-cap task string can encode near the 4 MiB initial limit and then overflow when the continuation members are added. The static proof must bound the canonical encoded `admission_ref` contribution (and every other open string), or impose and enforce a grammar whose expansion factor is accounted for.
4. **The overflow fixture is internally unconstructible as written.** §4 says violating the compile-time constraint is a build error, while §11 says the runtime overflow branch is exercised with synthetic constants that violate that same constraint. Name an executable test seam that can reach the branch without being rejected at compilation, or retain a legal production-state boundary fixture satisfying amendment `xit-dur-5`.

**Required correction:** recompute the proof from exact JCS-encoded shapes, all member-name/framing bytes, exact identifier/string grammars, the real full-run `parked_unknown` maximum, and a constructible max/one-over fixture. The runtime pre-commit sizing gate remains required regardless of the static proof.

## M10-DAG-R2-F3 — BLOCKER: `workspace_root_id` still has no exact value producer or worker-visible identity

§5 now says the manifest gains `workspace_root_id`, derived by admission from an “operator-provided workspace root” (`:73`). Neither frozen r40 §C.1 nor r10's run-start input includes that root, and rev2 does not define the new input carrier, the identifier's canonical recipe/domain (path, digest, opaque id, filesystem identity, normalization), or its equality semantics. Naming an English source is not yet a constructible producer contract.

The proposed split also disconnects writer and reader. M-10 freezes `workspace_root_id` privately; the worker receives only `manifest_digest` on frozen `assign` and cannot read m-10's store. Yet m-9 owns and records the full descriptor, whose local-action rows require that same `workspace_root_id`. Rev2's parked `authorize_tool_call` extension has m-9 supply only `canonical_resource` and `cwd`, so no surface gives m-9 the id it must record or lets both sides prove they refer to one root. Current m-9 bytes likewise state the requirement but name no carrier.

**Required correction:** pin the operator/app input surface for the root, the exact `workspace_root_id` derivation and encoding, the immutable manifest field, and one legal m-10→m-9 disclosure or deterministic shared derivation that makes the ticket value and executor record byte-equal. Keep the exact joint wire parked until m-9 settles it, but do not claim r1 F5 closed without a writer-to-reader path. Also narrow m-10's “structural validation” claim: symlink resolution cannot be verified from a declared relative string without consulting filesystem state.

## Accepted on these bytes

- **R1 F2:** §2/§4/§5 joint shapes are explicitly PROPOSED, NON-NORMATIVE, and PARKED; normative m-10 acceptance tables remain owner-real. The current m-9 producer is still under MUST-REVISE, so no joint-final claim advances.
- **R1 F3:** the durable `{resume_disposition, resume_action}` pair, present-iff rule, full-pair duplicate equivalence, receipt projection, and terminal overflow action source close the finding.
- **R1 F6:** the exact five-member `state_proposal_result` schema, correlation lifetime, per-disposition receiver effects, timeout/re-proposal discipline, stale-event tuple key, and crash recovery close the finding.
- **R1 F7/F8:** r10 startup and paired pre-ready clauses are expressly superseded; the two confirmation halves and completing hash are now correct.
- **Preserved directions:** full-ancestry manifest production, discriminated union, provider terminal-plus-receipt conjunction, two-time evidence split, `content_lost` exclusion, immutable snapshot, unique successor, M10-C0 two-form gate, M10-C1 telemetry separation, CI-4 death-set split, B/E parking, and H17 row remain sound subject to F1-F3.

## Re-review gate

Return one rev3 additive delta closing F1-F3. The fold must (a) align §1, §7, the M10-C1 confirmation, and fixtures on one canonical-row-decidable `VOID` mapping; (b) replace the frame estimate with exact encoded arithmetic and a runnable boundary fixture; and (c) make `workspace_root_id` a complete producer/carrier/consumer identity while keeping the unsettled joint wire non-normative. Do not file producer confirmations, the §6-D join, lane completion, F73 closure, re-lock, PLAN, T4/code authority, or implementation from rev2.

## Verification

- Reviewed design SHA-256: `ceb9ae3113ee7edb9d3022f17c5d674d554aff9c35b1e79156eae338020714e3`.
- Incoming DESIGN relay SHA-256: `4cdc3327def01879f44bf554ec0b1c3791bb656c2fcb0e0230c16c944ddbd6b7`; exact-file lint: `OK`.
- Frozen/governing hashes reproduced: r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; broker study rev8 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; completed m-10 confirmation `375da939fc1cfb1689ffd9ced9c892a166c94273afcabb3ab48d57e16e4f478c`.
- Counterparty bytes were reopened for the parked-seam check; no joint shape was treated as settled.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-100000.md`.
Next requested action: m-10.planner folds M10-DAG-R2-F1 through F3 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; every downstream confirmation/join/re-lock/PLAN/T4/code gate remains held.
