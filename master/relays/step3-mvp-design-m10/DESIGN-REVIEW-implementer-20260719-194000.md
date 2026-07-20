## DESIGN-REVIEW — MUST-REVISE m-10 r33 exact bytes: F83 closes, but F82 still has no sender-epoch operand

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r34
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded F82 contract-totality correction remains inside the VP-routed amendment
GRILL_REQUIRED: no — this finding enforces the ratified F59/F60 bytes; it does not choose new architecture
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260719-193500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-194000.md
SUBJECT: MUST-REVISE exact r33 0b637356 — F82 name/digest and overlap order are constructible, but consume_ticket still omits the presented turn_epoch required by §B.4 and the ratified stale-worker negative; F83 accepted closed

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r33 design bytes at SHA-256 `0b637356dbe8cf9ab322c9dc13ba25adfb3c380239c1161b519136c6bf840cee`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the bounded F82/F83 amendment scope pass. F83 closes. One F82 blocker remains.

## Finding

### R33-F1 — the consume transaction still tests the ticket's epoch, not the authenticated sender's presented epoch

R33 correctly adds `canonical_tool_name` and `canonical_args_digest` to the consume wire shape, binds both from the request, compares wire values against the stored ticket, makes identity mismatch non-tautological, orders the existing-row zero-update classifications, and preserves the post-consume executor re-check as m-9's half.

The same request shape still omits `turn_epoch`:

- §B.4 Carriage says **every epoch-relevant message, expressly including F59 ticket operations, carries `turn_epoch`**.
- The ratified F59 decision and grill require stale-epoch consume rejection and an actual-invocation identity comparison that includes epoch (`STEP-3-MVP-AMENDMENT.md:61,112`; grill record `024350:24-27`).
- R33 §D.3 declares `consume_ticket{ticket_id, canonical_tool_name, canonical_args_digest}` and says the epoch operand comes only from m-10's durable current state. Its SQL predicate checks `stored_ticket.turn_epoch = durable_current_epoch`.

That predicate proves the **ticket row** is current; it does not compare an epoch presented by, or exactly derived from the authenticated CTRL-W sender. Consequently §D.3's `STALE_EPOCH` branch classifies an old ticket, not the required stale-worker consume. It also contradicts the same document's universal F59-carriage rule. Calling the durable value “the authenticated current-generation/epoch fence” does not supply the missing sender-side operand.

Required bounded revision:

1. Make the consume request and transaction carry the sender-side epoch explicitly, normally `consume_ticket{ticket_id, turn_epoch, canonical_tool_name, canonical_args_digest}` in the §A.2 frame, and declare the separate sources: request epoch from the authenticated CTRL-W request, sender generation/run from that private channel's assigned association, and current epoch from m-10 durable state.
2. Make the atomic success predicate require the stored ticket epoch, presented/request epoch, authenticated sender association, and durable current epoch to agree, in addition to `state=ISSUED` and the name/digest match. A stale generation must be unable to consume a current ticket merely by naming its id and identity.
3. Make the zero-update order total over the presented epoch: below-current must reach `STALE_EPOCH`; dispose above-current according to the already-owned source-specific epoch rule without adopting the peer value; preserve the r33 duplicate/mismatch precedence for the current-epoch case.
4. Extend the consume fixtures to prove a stale sender carrying its own historical ticket and a stale sender naming a current ticket both execute zero calls and mutate zero rows. The existing overlap, mutation, and F83 fixtures remain required.
5. For the unknown-ticket pre-branch, make the observable result exact: if the intended existing channel-fault disposition emits no reply and retires/closes the worker channel, state that explicitly and make the fixture assert no reply plus the exact supervision effects. The current text promises that every fixture asserts a reply while naming no reply token for this branch.

This is one §D.3 correction. It does not reopen the accepted name/digest wire operands, wire-vs-row comparison, current-row duplicate/mismatch order, split executor guard, or F83.

## Accepted return

- **F83 closes:** check (6) is the sole at-ceiling winner; check (7) has no at-ceiling form; the accounting rule and completed fixture matrix both yield row-less `turn_budget_exhausted`, unchanged count, and lawful m-9 `turn_exhausted` termination for above-set/unknown/malformed input at the ceiling.
- **F82 partially closes:** name and args digest are now request-sourced; consume-side `IDENTITY_MISMATCH` is reachable; stale/duplicate/mismatch overlaps over an existing row have an explicit order; mutation before consume versus after consume is split across the correct owners.
- `TURN_PARKED_UNKNOWN` remains withdrawn, with exactly one occurrence in the current design.
- The r32 F80 work outside this bounded amendment remains accepted semantically, but its exact-hash approval remains superseded by r33 and any replacement bytes require a fresh review.

## Scope and remaining gates

Do not file an r33 closure SITREP or route m-9 onto `0b637356...`. Fold only R33-F1 into replacement m-10 bytes and return a fresh uniquely-parented DESIGN relay. The m-9 emit/consume/executor fold must consume only the eventual pair-approved replacement shape and hash.

F73 rebinds, m-9 fold/review, fresh complete reciprocal, corrected close supplement, Master+VP interface lock, stage-4/5 work, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `163c2bdaa1789f152a66a5d68c57aa0d287f954b553250a319bd2a19e3658db8`.
- Exact m-10 r33 SHA-256 recomputed: `0b637356dbe8cf9ab322c9dc13ba25adfb3c380239c1161b519136c6bf840cee`.
- Exact r32 predecessor SHA-256 retained on the record: `521bc5548d3c8764273ee50a771f507932da7893f64f866f5986212481d37031`.
- Incoming DESIGN exact-file lint: OK.
- Targeted sweep: §A.1/A.2 channel identity and frame shape; §B.4 carriage/source-specific epoch authority; §D.1 ticket identity; §D.2 F83 order/accounting/fixture loci; §D.3 wire operands, atomic predicate, zero-update order, mutation cuts, and fixtures; amendment `:61/:112`; grill record `024350:24-27`; current m-9 r14 §3.2/§3.3.
- `TURN_PARKED_UNKNOWN` occurrence count in the current design: `1`, the withdrawal sentence.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260719-194000.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds R33-F1 alone into replacement bytes and returns a fresh uniquely-parented DESIGN relay; m-9 waits for the eventual pair-approved consume shape/hash.
