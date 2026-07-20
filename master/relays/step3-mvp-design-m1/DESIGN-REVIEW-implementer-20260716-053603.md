## DESIGN-REVIEW - m-1 exact-byte r3 review of the Step-3 secret-boundary and seat-identity contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m1-review-r3
PARENT_DISPATCH_ID: step3-mvp-design-m1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the sole residual is internal normalization of the already-ratified F64 surface; no product fork is requested
GRILL_REQUIRED: no - stage-1 pair review consumes the ratified decisions; no unresolved product semantic remains
DESIGN_DOC_ID: step3-mvp-design-m1-secret-boundary-seat-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m1/DESIGN-planner-20260716-053040.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-053603.md
SUBJECT: must-revise - normalize every normative F64 capability-surface statement to include typed Describe

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner - I reviewed the revised contract bytes at SHA-256 `7baffe40412f585afe05c8e35f41cd7ce3168af94b998c99849baa9ef79a4af9`, the addressed r2 fold relay, the r2 revision bar, and the current m-7 consumer surface.

MR1 is closed. MR2 is closed. MR3 lands correctly in Section 1.4b and its stale-epoch negative, but the document's later normative summaries retain the old three-verb-only wording. Because Section 1.4b explicitly says `Describe` is not a verb, those summaries do not include it by implication. One narrow consistency fix remains before exact-byte approval.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Finding

### R3-F1 - The identity and consumer clauses still define an incomplete F64 surface

Section 1.4b now correctly defines the USE capability as authorizing the three canonical relay verbs, typed `Describe`, and push/rediscovery, all behind the same epoch fence (`2026-07-16-step3-mvp-secret-boundary-seat-identity.md:33-36`). The same contract then states:

- Section 2.3: the broker checks every "verb invocation and push delivery" (`:56`).
- Section 2.5: while current, the capability authorizes "the three verbs + push" (`:66`).
- Section 3: m-7 realizes the "per-verb/push epoch gate" (`:86`).

Those are normative identity and consumer-obligation clauses, not historical notes. The contract also explicitly classifies `Describe` as not a verb, so the later phrases cannot be read as shorthand that includes it. The current m-7 contract correctly uses the closed worker surface `three canonical relay calls + typed Describe + push` and fences all four request operations (`2026-07-16-step3-mvp-transport-broker.md:99-114,148-158`). Leaving m-1 internally split creates an avoidable consumer-confirmation ambiguity.

Required revision: normalize all three statements to the same closed surface already approved in Section 1.4b: the three canonical relay verbs plus typed `Describe`, with push fenced at forward time. In Section 2.5, make the authority summary point to the complete Section 1.4b surface or enumerate it completely. In Section 3, require m-7's per-operation fence over the three relay verbs and typed `Describe`, plus the per-push fence. Preserve that `Describe` is metadata, not a relay verb or dispatch-tool name, and carries no relay-acceptance authority.

## Accepted portions

- **MR1 is closed.** S-A resolution follows the m-3 provider-request policy gate at m-8; S-B resolution follows the broker's operator-provisioned startup/re-auth path; F59 and F64 are correctly excluded from credential-resolution authority (`:34`).
- **MR2 is closed.** Broker restart leaves `turn_epoch` unchanged while killing every old connection-scoped token; eligible reattachment requires the authenticated restart protocol and fresh capability material (`:36`, `:78`, `:95`).
- **MR3's intended surface and negative are correct where revised.** Section 1.4b includes typed `Describe` behind the same epoch fence and precisely preserves its non-verb/non-dispatch/non-identity classification; Section 5 requires stale-epoch refusal (`:35`, `:95`). R3-F1 asks only that the duplicate normative summaries stop contradicting it.
- **The previously accepted sink, reference/capability split, five-event matrix, overlap fixture, F57 residual, identity/accountability model, and no-conductor/store-change route-back remain accepted.**
- **No operator decision or new grill is needed.** This is a closed-surface consistency correction under the existing MR3 ruling.

## Revision bar

Return fresh contract bytes and a fresh hash that replace the incomplete F64 summaries at Sections 2.3, 2.5, and 3 with the complete three-relay-verb plus typed-`Describe` request surface and per-push fence. Preserve every accepted r1/r2 portion and issue a fresh uniquely-parented DESIGN for exact-byte re-review. Consumer confirmations remain held until m-1 pair approval.

## Verification

- `shasum -a 256 master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` -> `7baffe40412f585afe05c8e35f41cd7ce3168af94b998c99849baa9ef79a4af9`.
- Read and linted the exact addressed relay `master/relays/step3-mvp-design-m1/DESIGN-planner-20260716-053040.md`; `TO`, lineage, `DESIGN_DOC_ID`, and claimed hash match.
- Re-tested r2 MR1-MR3 against the exact contract and current m-3/m-7/m-10 authority boundaries; only R3-F1 remains.
- Checked current m-7's capability and fence clauses at `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:99-114,148-158`; its complete surface confirms the m-1 normalization target.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file and lane-root relay-lint.py exit 0
Next requested action: m-1.planner normalizes only the three incomplete F64 surface summaries, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; consumer confirmation remains held.
