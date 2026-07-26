## DESIGN-REVIEW — r1 MUST-REVISE: the producer direction is sound, but the manifest union is not wire-total, the continuation action is not durable, the C descriptor has no authoritative operands, and the broker fold is incomplete

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — every finding is a mechanical completeness or cross-owner contract defect inside the released design scope; no operator-ratified product choice must be reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260722-013000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-073000.md
SUBJECT: MUST-REVISE exact rev1 `1c699317…`: close the D2 wire/input union, settle the two joint m-9 seams, persist and compare `resume_action`, prove the exact frame bound, source the C descriptor operands, totalize the M10-C0 receiver, and complete the r10 sweep

## Verdict

**MUST-REVISE** on exact design SHA-256 `1c69931703986049bcba62e229b6b7ffc38e08d7e7cb1038e65415841e9d3792`.

The additive discipline, full-ancestry production point, provider terminal-plus-receipt conjunction, two-time evidence split, `content_lost` exclusion, conditional writer-fence routing, B/E parking, C applicability direction, M10-C1 telemetry separation, CI-4 death-set direction, and H17 row shape are good. Approval is blocked by owner-real constructibility and totality defects in the emitted manifest, the joint m-9 frames, the continuation disposition record, the C ticket, and the simplified broker fold.

## M10-DAG-R1-F1 — BLOCKER: the manifest input set and emitted union are not closed or decodable

§1 declares the input set to be every `tool_calls` and `provider_attempts` row, then maps an authorization-only `VOID` row even though `tool_authorizations` is outside that set. Frozen r40 also gives `VOID` five typed reasons; only `VOID/expired` can represent the issued-unconsumed expiry cut, while the other four are determinate issue denials. The unqualified rule “authorization `VOID` with no `tool_calls` terminal” therefore both exceeds the declared input domain and can misclassify definite no-effect denials as `uncertain`.

The per-entry schemas contain identity plus `terminal` but omit the closed class discriminant and a tool/provider kind discriminant. A top-level `entries[]` carrying those shapes cannot be decoded into the ratified three-class union or sorted/validated without inference from field presence.

**Required correction:** state one exact producer input domain, qualify any authorization-orphan cell by the exact durable state/reason that represents the ratified cut, prove every included row maps once, and emit a closed discriminated union whose canonical schema carries class and entry kind explicitly. Include fixtures for all five `VOID` reasons and for unknown/extra/missing discriminants.

## M10-DAG-R1-F2 — BLOCKER: the purported JOINT frames do not match m-9's current producer bytes

§2 proposes `content_ready{run_id, turn_id, attempt_id, turn_epoch, prefix_digest}` plus an acknowledgement. Current m-9 rev1 proposes the one-way frame `content_ready{run_id, turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id, turn_epoch}`. §4 proposes `report_resume_disposition{run_id, turn_id, turn_epoch, disposition, resume_action?}`; m-9 currently proposes `{turn_id, disposition, resume_action}`. These are not one joint seam: the members, digest identity, generation evidence, and reply semantics differ.

The m-9 pair review correctly leaves both shapes parked and notes the mismatch. This m-10 review cannot certify either proposal as JOINT-FINAL before the two owners settle one exact frame, sender/receiver table, durable row, receipt, duplicate-equivalence predicate, and crash fixtures. Keep the proposals explicitly non-normative until those exact shared bytes exist; the final pair review must cover the settled result.

## M10-DAG-R1-F3 — BLOCKER: `resume_action` is neither durable nor part of replay equivalence

§4 adds only `resume_disposition` to the durable turn/snapshot state; it makes inbound `resume_action` optional; and its duplicate/conflict branches compare only the disposition. A replay with the same disposition and a different action therefore receives the same success receipt, while no committed projection can supply the operator surface required by amendment §5-D. The same omission leaves the `resume_frame_overflow` terminal branch without a named durable manual action source.

**Required correction:** define the closed presence/domain rule for `resume_action`, persist it with every disposition or terminal-overflow record that requires it, compare the complete `{disposition, resume_action}` identity for equivalent duplicate versus conflict, carry the committed pair in the receipt/projection, and add conflicting-action and crash/replay fixtures. The worker's no-work gate must key on the committed receipt without allowing the action to drift.

## M10-DAG-R1-F4 — BLOCKER: the required exact frame-bound proof is only an estimate

§4 says an entry is “≈ ≤ 512 B” and that constants are chosen so the worst case is at most `FRAME_MAX/2`, but supplies neither the constants nor bounded lengths for identifiers, digests, paths, `admission_ref`, and framing overhead, nor the arithmetic. A runtime sizing gate prevents an un-emittable commit; it does not discharge the amendment's separate exact-bound design obligation.

**Required correction:** name every contributing maximum and compiled chain/depth constant, show canonical-encoding and frame overhead arithmetic against the 4 MiB limit, and bind the proof to exact-fit and one-byte-over fixtures. If any string is not intrinsically bounded, the design must supply its owner-real bound or remove the static claim and route the missing contract.

## M10-DAG-R1-F5 — BLOCKER: m-10 cannot construct the C descriptor from the wire and rows it owns

§5 says m-10 derives canonical paths/target sets, relay targets, cwd, workspace root, and implementation references at ISSUE. Frozen r40's `authorize_tool_call` gives m-10 only the tool identity and one-way `canonical_args_digest`; r40 explicitly stores and matches that digest and never interprets args. No `workspace_root_id` producer is named in the frozen manifest either. A digest cannot be inverted into the resource/cwd operands, so the proposed “unconstructible ⇒ refuse” gate has nothing authoritative to validate.

Current m-9 lane-2 bytes place actual descriptor derivation with m-9, reinforcing that this is an unsettled owner seam rather than an m-10-local derivation.

**Required correction:** pin the exact ISSUE request extension and source split: which descriptor members m-9 supplies from the parsed call, which m-10 derives from named immutable manifest/constants, how m-10 validates each without re-owning arguments, and which complete members participate in replay identity and consume comparison. Preserve the ratified R/∅ table and invocation-stage split; do not widen authority.

## M10-DAG-R1-F6 — BLOCKER: the M10-C0 consumer is not total over the broker's ordered result table

§6 reproduces the exact `state_proposal` request but abbreviates `state_proposal_result` to validation of `installed` plus timeout/re-proposal. It does not reproduce the broker's five dispositions or define m-10's control action for each. In particular, `transition-started`, `rejected-stale`, `rejected-transition-active`, and `rejected-malformed` need bounded sender-side effects: whether the assign gate stays closed, what event or deadline permits re-proposal, and how retry avoids a busy loop. A mismatched old `epoch_installed` must never open the gate while another transition is active.

**Required correction:** carry the exact closed response schema and ordered five-member table from the pair-approved study, then give m-10's total receiver fold, correlation lifetime, duplicate rule, timeout, retry trigger, and gate state for every member and crash cut.

## M10-DAG-R1-F7 — BLOCKER: the r10 supersession sweep misses two live old-protocol loci

§8 replaces r10's recovery line, retire/replace line, epoch-linearization line, census rows, and §14. It does not supersede r10 §3's startup step 3, which still publishes `epoch_state` “per §B.4/§B.5 rules,” or r10 §4's paired pre-ready connector-failure paragraph, which still says the “§B.5 distribution/install sequence” runs before retry. Both are executable lifecycle statements and both retain the removed protocol vocabulary after the claimed sweep.

**Required correction:** quote and replace those exact clauses, then repeat the search over all r10 bytes for transition-ledger, §B.5 handshake/distribution, pending-transition, crossing-row, and lost-install terminology. Preserve the paired-owner E+1 behavior; only replace the retired broker mechanism.

## M10-DAG-R1-F8 — LINEAGE CORRECTION: cite the pair confirmation, not only the planner half

The basis and §7 call `step3-relock-broker-confirm/RECONCILE-planner-20260721-214500` the pair-approved/co-signed confirmation. That file is the planner half. The pair confirmation is completed by `RECONCILE-implementer-20260721-221000.md` at SHA-256 `375da939fc1cfb1689ffd9ced9c892a166c94273afcabb3ab48d57e16e4f478c`. Cite both halves or the completed implementer confirmation so the F73 lineage is mechanically true.

## Accepted on these bytes

- **Authority and additivity:** released rev2 `6df5367f…`, release `8dc0a0cd…`, amendment rev12 `1125b0a0…`, study rev8 `64f9136e…`, frozen r40 `d2ce9831…`, and frozen r10 `6fd1d655…` reproduce; the frozen bases did not move. This remains DESIGN-only.
- **D2/D3 direction:** production inside continuation admission from canonical owner rows, full ancestry, deterministic ordering, terminal-plus-receipt conjunction, settlement-time versus resume-time evidence, `content_lost` exclusion, immutable snapshot, unique successor, and two distinct overflow outcomes are sound subject to F1–F4.
- **C:** the ratified per-action R/∅ applicability table, single cwd encoding, nonexistent-cwd rejection, and invocation-resolved env/interpreter split are sound subject to F5's source contract.
- **M10-C0/C1/C2:** always-propose recovery, the two-form install-proof direction, CI-3 shrink, canonical-row-only manifest derivation, telemetry-only `boundary_cut`, and CI-4 broker death-set separation are sound subject to F6–F8.
- **Parking and conditionality:** B/E carriage remains properly parked. The m-10 segment branch remains inert because m-9 selected the worker-owned OS-lock branch; that selection is itself under m-9 MUST-REVISE and creates no m-10 producer obligation on these bytes.
- **H17:** `m10-broker-spawn` contains all 21 census labels exactly once and preserves the secret-holder boundary.

## Re-review gate

Return one revised additive delta closing F1–F8. Before requesting re-review, settle the exact §2/§4 joint frames with m-9 or keep them explicitly parked and non-normative; route the C ISSUE operands to the owning producer; quote every r10 replacement locus; and bind the new full-document hash. Do not file producer confirmations, a lane SITREP declaring completion, the §6-D join record, F73 closure, re-lock, PLAN, T4/code authority, or implementation from r1.

## Verification

- Reviewed design SHA-256: `1c69931703986049bcba62e229b6b7ffc38e08d7e7cb1038e65415841e9d3792`.
- Incoming DESIGN relay SHA-256: `760ac9b05e33916a1216b02c84048cafdb1bb812252baf946e8b5ed1d48fa2d6`; exact-file lint: `OK`.
- Released rev2 dispatch SHA-256: `6df5367ff294424e06e9f09e6e078330d85d16c47452018f12baf5e64e72a10d`; release relay SHA-256: `8dc0a0cdce0826edf00e744f9b269e7a009bb833164df2e43a9254c873320c64`.
- Governing/frozen hashes reproduced: amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; study rev8 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- Counterparty state checked: m-9 rev1 `2de214ea70cbfcf21dc873b938633dc5ee61ba2f0f0c488f7f6e33c1be642a43` remains MUST-REVISE; its review explicitly keeps the joint C/D wires parked.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-073000.md`.
Next requested action: m-10.planner folds M10-DAG-R1-F1 through F8 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; every downstream confirmation/join/re-lock/PLAN/T4/code gate remains held.
