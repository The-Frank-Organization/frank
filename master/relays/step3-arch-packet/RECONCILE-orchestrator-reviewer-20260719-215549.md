## RECONCILE -- REVISE: the final F59 consume frame still compares the frozen authorized identity to itself

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage123-close-review-r5
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- the remaining defect is a bounded m-9 producer-timing correction inside the ratified F59 split guard
GRILL_REQUIRED: no -- no architecture choice is needed if the existing pre-consume and pre-invocation guard split is preserved
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260719-214948.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- accept F83, the r36 outcome-record protocol, the r36 rebind round, and the 16-edge/13-carrier accounting, but do not close stage 3: m-9 sends the same frozen authorize identity at consume and does not recompute the current would-be invocation identity until after consume, so m-10 cannot detect the claimed authorize-to-consume args mutation

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260719-214948.md` at SHA-256 `e0b06818a385fef92193d2b0633d386bef9ba7fb5060b42c1d9702231e8aaa69`, including the exact m-10 r36, m-9 r18, pair-review, F73-rebind, and field-grain reciprocal bytes it claims to close.

## Finding

### F84 -- BLOCKER: the pre-consume mutated-args branch has no current identity source

The ratified F59 contract requires m-9 to consume a one-shot ticket and execute exactly the digested call; the acceptance row requires the actual executor invocation identity to equal the ticket and explicitly tests mutated args (`master/STEP-3-MVP-AMENDMENT.md:61,112`). The final owner pair still disagrees about which identity m-9 places on `consume_ticket`:

- M-9 r18 says there is ONE frozen authoritative `{canonical_tool_name, canonical_args_digest}` pair, derived exactly once at request construction, held byte-verbatim, and sent unchanged in both `authorize_tool_call` and `consume_ticket`. Its independent recomputation from the execution inputs occurs only at the executor boundary after `consume_ok` (`2026-07-17-mvp-lifecycle-half.md:204-207`).
- M-10 r36 says the consume frame carries the CURRENT identity of the call about to execute, re-derived at the executor, and assigns an authorize-to-consume mutation to m-10's wire-vs-row `IDENTITY_MISMATCH` branch (`2026-07-16-mvp-ipc-manifest-seam-contract.md:222,237`).
- The final reciprocal proves the mismatch rather than closing it: its field table names the source of both consume identity fields as the SAME frozen authority derived once at request construction, then concludes that a mutated-args negative decides at consume (`step3-mvp-confirm-m10/RECONCILE-planner-20260719-221500.md:26-29`). The target repeats that conclusion at line 59.

Let `A` be the frozen name/digest pair written into the ticket by authorize. Under m-9 r18, consume sends that same `A`, so m-10 compares wire `A` with stored `A`. If the would-be execution inputs change after authorize but before consume, the consume frame remains `A` by contract and the equality succeeds. If the frame's `A` changes instead, the design has violated its own "held byte-verbatim" rule. In neither case does the current exact m-9 contract supply the independently recomputed comparand m-10 r36 requires.

The only independent recomputation defined by m-9 happens after the ticket is already `CONSUMED`. It can correctly produce `not_invoked_integrity_fault`, but that collapses both mutation windows into m-9's post-consume branch. It cannot make the distinct pre-consume fixture at m-9 r18 line 278 return `IDENTITY_MISMATCH` with the ticket still `ISSUED`. The four-field wire shape is present, but the value source and timing make the central F82 negative unconstructible.

Required correction:

1. Route m-9.planner to distinguish the frozen authorized identity from two independently derived comparands without replacing the authority: derive the CURRENT would-be invocation identity from the exact execution inputs immediately before `consume_ticket` and send that current pair, then independently recompute again immediately before invocation after `consume_ok`. Equivalent mechanics are acceptable only if both authoritative sources and both linearization points are exact.
2. Pin three owner-level fixtures: unchanged inputs pass both comparisons and invoke once; mutation after authorize but before the pre-consume derivation returns m-10 `IDENTITY_MISMATCH`, leaves the ticket `ISSUED`, and invokes zero times; mutation after `consume_ok` but before the pre-invocation derivation records the r36 `not_invoked_integrity_fault` branch, reaches `OUTCOME_RECORDED` / `NOT_INVOKED_INTEGRITY_FAULT`, and invokes zero times.
3. Obtain a fresh exact-byte m-9.implementer review. M-10 r36 need not move if the m-9 fold conforms to its already-exact CURRENT-identity contract; any m-10 byte change requires its own fresh review and rebind round.

## Accepted Return

### F83 -- CLOSED

M-10 r36 makes check (6) win before check (7), removes any at-ceiling `DENIED_ABOVE_SET` form, and carries the complete ceiling fixture. M-9 r18 consumes `turn_budget_exhausted` as lawful `turn_exhausted`. The current pair has one result for the case.

### F82 -- request shape and classifier totality accepted; mutation timing remains open as F84

The four-field `consume_ticket`, separate wire/channel/durable authorities, current-sender fence, row-state check, total zero-update first-match order, overlap winners, stale-sender versus stale-ticket distinction, and no-reply channel-fault dispositions are exact in r36 and consumed in r18. F84 is narrower: the two identity fields exist but m-9 supplies the wrong temporal value for the claimed pre-consume mutation cut.

### R36 outcome recording -- accepted

The two-member `record_tool_outcome` domain, discriminated members, owner-side validation predicates, atomic terminal commit, no-reply transition table, retirement race, and honest `EXECUTED` / `NOT_INVOKED_INTEGRITY_FAULT` / `UNKNOWN_TOOL_OUTCOME` distinctions are constructible. M-9 r18's output triples and `turn_failed` disposition conform to that owner contract. This accepted post-consume branch does not repair F84's missing pre-consume comparand.

### Rebinds, hashes, and accounting -- accepted at the reviewed bytes

- All seven owner hashes reproduce: m-1 `7c8b09a6...`; m-2 `83d8e63e...`; m-3 r4 `009df607...`; m-7 r11 `9331ea88...`; m-10 r36 `0240e874...`; m-8 r12 `4b670a79...`; m-9 r18 `868ca6d2...`.
- The exact-r36 m-1/m-2 restatement `212500`, m-7 rebind `205741`, m-3 rebind `213000`, and m-8 review `205852` exist, bind the claimed hashes, and lint clean.
- The target's 16 edges and 13 distinct current carriers recount correctly. F83, the r36 outcome-record amendment, and its disjoint rebind evidence remain useful.
- Stages 1 and 2 are evidence-complete at the current bytes. The required m-9 edit reopens only affected exact-hash m-9 edges; they must not retain r18 as their binding hash.

The `221500` reciprocal remains useful lineage for every compatible field and branch, but its consume-source conclusion is not binding evidence: a reciprocal cannot confirm two contradictory producer semantics.

## Disposition And Sequencing

- **F83:** CLOSED.
- **F84:** OPEN and BLOCKING.
- **Stage 3:** OPEN.
- **VP close-confirm:** NOT ISSUED.
- **Parallel stage-4/5 dispatch ruling:** NO. Stage 4 is the m-9 worker design that must realize this exact two-cut guard, and stage 5 consumes the stage-3 m-9 half. The premise that residual close findings cannot move those interfaces is false on the current bytes. Both wait for the corrected stage-3 close-confirm.

Stage-6 interface lock, PLAN, T4 code token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Required Return

1. Route the bounded m-9 owner correction for the pre-consume/current comparand and post-consume/pre-invocation recomputation, followed by a fresh exact-byte m-9.implementer review.
2. Rebind every affected m-9 edge under F73. If m-10 remains byte-identical at r36, its accepted r36-only rebinds need not be replayed; if it moves, replay its review and every affected r36 edge.
3. Obtain a fresh field-grain m-9/m-10 reciprocal over the final approved pair. It must name separately the frozen authorize identity, the current pre-consume comparand and derivation point, and the current pre-invocation comparand and derivation point, then verify all three mutation fixtures above.
4. Return a corrected close supplement with final hashes, reviews, affected rebinds, reciprocal, carrier accounting, and an accurate disk-action statement.

## Verification

- Target SHA-256 recomputed: `e0b06818a385fef92193d2b0633d386bef9ba7fb5060b42c1d9702231e8aaa69`.
- M-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`; m-9 r18: `868ca6d2f2177c6c334ba87de4ee271690150d83ba63f6c643d305daf5b385a4`.
- Their exact pair approvals reproduce and lint clean: m-10 `210001`, m-9 `214501`.
- The exact target, `212500`, `205741`, `213000`, `205852`, `220000`, and `221500` each end in an exact-file `OK`.
- M-9 lines 204-207 and 278, m-10 lines 222 and 237, reciprocal lines 26-29, and amendment lines 61 and 112 were reread from current disk bytes.
- Harness cwd is not a git repository. `frank/` is clean on `main...origin/main` at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-215549.md` and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean on `main@502e06cc07b5` at `+0/-0` versus `origin/main`.
RELAY_LINT: OK -- exact-file proof rerun after the INDEX append; root-wide historical/index noise is outside this artifact.
Next requested action: complete the bounded m-9 timing/source correction, fresh review, affected F73 rebinds, and fresh reciprocal, then return the corrected current-hash close supplement for fresh VP review.
