## DESIGN-REVIEW — approve m-9 lane-2 D/E/C/B delta r5 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r5
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the evaluation-order repair is determinate and reopens no operator-ratified product choice
GRILL_REQUIRED: no — no product or cross-domain boundary choice remains open in this component
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260722-140000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-1.planner, m-2.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260722-150000.md
SUBJECT: APPROVE exact-byte m-9 lane-2 delta r5 at SHA-256 c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b — M9-DAG-R4-F1 closes; producer confirmations, joint joins, integrated re-lock, PLAN, T4, and code remain separate

DESIGN_REVIEW_VERDICT: approve

m-9.planner — **APPROVE** the complete m-9 lane-2 additive delta r5 at exact SHA-256 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b`.

This approval is byte-bound. Any change to `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, including metadata or revision history, voids it and requires fresh full-byte pair review.

## R4 blocker closure

**M9-DAG-R4-F1 closes.** The terminal-seal rule is now §1.4 step 0, a state guard evaluated on raw physical suffix bytes whenever the last accepted record is `segment_seal`. It therefore dominates self-consistency/parsing, duplicate conflict, sequence gap/regression, chain validation, record admission, and later marker selection.

The disposition is total and unambiguous:

- EOF is the only zero-byte case and ends the sealed segment normally.
- An immediately adjacent raw line byte-identical to the accepted seal is the sole collapse exception; reapplying step 0 handles any finite run of identical repeats without admitting one.
- Every non-empty non-identical suffix fails closed for the whole run before parsing. That includes a lone newline, trailing NUL bytes, a torn/unparseable fragment, same-`seq` conflict, gap/regression, wrong `prev_digest`, second seal, and perfect continuation. None can fall through to an ordinary clean-prefix branch.

The torn-repeat residual is the correct conservative choice under the stated confusion-first invariant. A torn benign seal retry and a stale writer's torn append are observationally indistinguishable in these bytes; collapsing the former would also silently admit the latter. Fail-closed preserves the exclusive-writer claim, with the false-fault cost named honestly.

The §10 battery is strong enough to prevent regression to r4's ordering: all six post-seal suffix classes require the same step-0 whole-run failure and assert branch attribution; cases 2–5 would instead hit r4's ordinary branches, so r4 fails the negative battery. The exact-repeat positive control, torn-repeat negative, repeated step-0 loop, and boundary-digest-agrees assertions separately prove the exception and prove that the digest is not being substituted for the terminal-state mechanism.

## Full-byte review result

The rest of the current artifact remains coherent with the already-accepted r4 surfaces:

- one total non-self-referential `boundary_record`, including the `segment_open` at `seq` `"0"` zero-round case, and a seal only after an honoured marker;
- distinct late-write mechanisms: sealed suffix via step 0, unsealed boundary movement via digest mismatch, and unmarked/torn unsealed suffix as untrusted without a false mismatch;
- complete topology plus clauses 8–11 for composition identity, envelope homogeneity, rotation progression, and contiguous non-reused generation runs;
- clause 12 binds every append/create choice to the current `assign.generation_id`, while cross-generation forward-link intent is not falsely compared and backward/digest links remain mandatory;
- the per-run writer fence, marker membership, provenance, D2 reconciliation/no-work gate, E/C/B surfaces, retention, and first-action table retain their accepted meanings;
- all five consumer/joint inputs in §9 remain explicitly parked rather than consumed as settled bytes.

No new finding survives the adversarial pass.

## Exact evidence

- Addressed relay exact-file lint: OK; routing is directly `TO: m-9.implementer` and parents review-r4.
- Current delta SHA-256: `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b`.
- Released rev2 dispatch SHA-256: `254950dd5e164d151739aff827efd1f8ea67887832ca9b969334370453003f25`.
- Ratified Stage-6 amendment rev12 SHA-256: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Frozen m-9 worker r7 SHA-256: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Frozen m-9 lifecycle r21 SHA-256: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Frozen m-10 seam r40 SHA-256: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.

## Authority boundary

This is m-9 component pair approval only. It authorizes the planner's requested report/confirmation/join sequencing within the released DESIGN dispatch; it does **not** itself settle or attach any parked m-2/m-10 producer bytes, complete either joint m-10/m-1 record, grant integrated Master+VP re-lock, or authorize PLAN, T4/code, source edits, credentials, provider calls, release binding, live E3, merge, or deploy.

ACTIONS_GIT_REF: docs-workspace action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design document, frozen artifact, `frank/` source, branch, commit, integrated lock, producer confirmation, joint record, PLAN, credential, provider call, release binding, merge, or deploy action
FINAL_GIT_STATUS_SHORT: unavailable — harness root is not a git repository; `frank/` clean at `c78da3815a34480590071295c1e09bb7d53c10b6`
RELAY_LINT: OK — exact-file and dispatch-root lint exit 0
Next requested action: m-9.planner files the byte-bound lane SITREP, obtains the named F73 producer confirmations, and opens the §D joint record with m-10 plus m-1 redaction review; master owns the later integrated re-lock sequence.
