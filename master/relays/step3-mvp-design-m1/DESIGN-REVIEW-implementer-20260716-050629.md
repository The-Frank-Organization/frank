## DESIGN-REVIEW - m-1 adversarial review of the Step-3 secret-boundary and seat-identity contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m1-review-r1
PARENT_DISPATCH_ID: step3-mvp-design-m1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings narrow interface mechanics under ratified F57/F60; no product fork is requested
GRILL_REQUIRED: no - stage-1 pair review consumes the ratified grill decisions; the pressure questions below were answered from current owner/consumer bytes
DESIGN_DOC_ID: step3-mvp-design-m1-secret-boundary-seat-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m1/DESIGN-planner-20260716-045225.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-050629.md
SUBJECT: must-revise - split credential references from USE capabilities, pin the broker provisioning sink, and close the rotation/replacement overlap proof

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner - I reviewed the final contract bytes at SHA-256 `aa90fa45b9d8f37d5a820927e93eb3a30680ced9ec5111911c5881ec820a315d`, the addressed DESIGN, its master parent, the ratified MVP amendment, the locked m-1 lifecycle substrate, the live re-mint mechanism, and the current parallel m-7/m-10 consumer bytes.

The F57 claim boundary is honest and the F60 identity model is directionally correct: identity remains the committed mint-pivot chain, worker replacement remains app-side supervision, the conductor cannot attribute a relay to a worker generation, and no conductor protocol/store change is required. Three interface ambiguities still block pair approval. The first two are already producing divergent readings in the parallel m-7 contract, so they must be fixed in the m-1 source before consumer confirmation.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Findings

### F1 - The contract collapses a non-authorizing credential reference and an authorizing USE capability

Section 1.4 says an opaque credential reference "confers nothing," is identity-inert and secret-inert if leaked, and is resolved only inside m-8 or the broker (`2026-07-16-step3-mvp-secret-boundary-seat-identity.md:31-32`). Section 2.5 separately defines the worker's USE capability as a revocable delegation that does confer use of the seat channel through the broker's epoch gate (`:57-62`). Those are different security objects and need different claims.

The ambiguity is observable now. The parallel m-7 contract calls its per-generation worker token the m-1 Section 1.4 form, while also saying possession moves calls through the broker gate (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:77-83`). A current-epoch USE token on the private broker IPC surface is intentionally authorizing; it is secret-free and not an identity, but it is not leak-inert. Conversely, a credential reference given to m-10 must not become bearer authorization merely because its bytes are opaque.

Required revision:

1. Name and separate the two types: an opaque credential locator/reference used for custody orchestration, and the epoch-bound worker USE capability used for channel delegation.
2. State that possession of a credential reference alone neither authorizes resolution nor selects/attaches a credential. Resolution occurs only inside the holder after an independently authenticated and authorized operation; the reference contains no credential-derived verifier.
3. State honestly that a current USE capability authorizes only the three broker verbs plus push/rediscovery on its private IPC connection, subject to the current-epoch gate. It carries no secret bytes, creates no principal, and has no offline credential meaning, but leaking it onto an accessible current-epoch broker channel is not inert.
4. Add consumer negatives for reference-without-authority refusal and for capability replay after epoch advance/broker restart. Do not let m-7 or m-8 satisfy the credential-reference requirement merely by pointing to a bearer capability token.

### F2 - The visibility-floor wording does not identify the broker's allowed persistent provisioning sink

The contract says a new S-B credential appears only in the operator submit reply or 0600 binding-table admin path (`:12-15`) and repeats that the visibility floor is "operator submit reply / 0600 table only" (`:54-55`). It then requires the broker to obtain the new credential through operator/admin provisioning. The current m-7 realization chooses an operator-authored, permission-checked 0600 broker credential file and re-reads that file on rotation/restart (`2026-07-16-step3-mvp-transport-broker.md:60-66`, `:111-120`). That is an additional persistent secret-bearing custody sink unless the m-1 contract explicitly classifies it.

The current wording therefore admits two incompatible readings: either the broker reads the conductor binding table directly, which the m-7 contract rejects, or an operator copies the credential into a broker-private file, which the m-1 text currently fails to name while claiming "only" the original two surfaces.

Required revision: pin the handoff class. If the intended mechanism is the current m-7 choice, state that the operator submit reply/admin path is the source and that one broker-private, permission-checked 0600 credential file is the authorized runtime custody sink. Include it in the S-B file/census, rotation overwrite, crash-dump, log/error, backup, and deletion obligations. Preserve: no m-10 store/frame/path, no worker/m-8 surface, no conductor record/projection/error, and no app-side relay of credential bytes. If a different sink is intended, route the m-7 contract back before either pair approves.

### F3 - The two-counter law lacks the event matrix and overlap fixture requested by its own pressure question

Sections 2.2 through 2.4 correctly say replacement advances only `turn_epoch`, re-mint advances only mint-generation, broker channel cycling only re-binds, and neither counter implies the other (`:42-55`). Section 5 tests the two sequential independence legs (`:76-77`). It does not pin the full restart/overlap matrix, especially re-mint concurrent with worker replacement, even though the addressed relay explicitly asks whether rotation-during-replacement and broker re-auth remain airtight.

The parallel m-7 bytes make the missing cross-product concrete: epoch updates linearize at the broker, while conductor re-mint independently force-closes and re-authenticates the seat channel (`2026-07-16-step3-mvp-transport-broker.md:94-103`, `:111-116`). A build could implement each isolated leg correctly while leaving an overlap window or accidentally advancing/reissuing the other counter during recovery.

Required revision: add a compact event/effect matrix covering worker replacement, broker restart, conductor restart, re-mint, and re-mint concurrent with replacement. Pin for each event: mint-generation effect, `turn_epoch` effect, bind/re-bind effect, capability disposition, and accepted in-flight disposition owner. Add one overlap fixture proving that either event order yields exactly one mint advance plus exactly one epoch advance, no implicit `seat_mint`, no epoch reset/reuse, old credential rejection, old capability rejection, and successful use only after both the current credential channel and current-epoch capability are established. The m-7 contract may own the mechanism, but m-1 must own this identity-semantic expected result.

## Accepted portions

- **F57 honesty is correct.** The contract limits its claim to non-injection and accidental-disclosure reduction, explicitly retaining same-user peer inspection as the unsandboxed MVP residual (`:17-35`). The sentinel wording does not overclaim OS isolation.
- **The identity/accountability boundary is correct.** Logical seat identity remains `(seat address, committed mint-generation)`; worker generation is not a principal; conductor records remain stamped to the seat; app-side generation attribution is E0 bookkeeping, not identity (`:39-65`).
- **The main two-counter direction is correct.** Automatic supervision cannot invoke `seat_mint`, m-10 has no mint authority, and replacement does not create a new identity (`:42-52`). F3 asks for closure of the event cross-product, not a different model.
- **The s6 mapping is faithful.** `minted`, `bound`, and `active` retain their locked meanings; the broker becomes the binding party; active remains first accepted relay per committed mint-generation (`:48-55`; locked s6 contract `2026-07-06-s6-transport-amendments.md:91-113`).
- **The route-back claim is clean.** `turn_epoch`, references, capabilities, and app records are app-side; mint-generation and `realized_mint_ref` reuse the landed pivot/binding machinery. No conductor verb, record/member, registry, or store change is required (`:73-77`; live `frank@502e06c` binding and force-close paths checked).
- **No new operator decision is needed** if the revision adopts the already-authored m-7 provisioning choice and only disambiguates the existing ratified F57/F60 mechanics. A different secret sink or a new identity/counter authority would require route-back.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Separates opaque credential references from epoch-bound USE capabilities and states each object's authority/leak claim precisely.
2. Pins the broker credential provisioning/custody sink and updates the S-B non-injection/census obligations without weakening F-S6-M1-1.
3. Adds the five-event counter/lifecycle matrix and one rotation-during-replacement overlap fixture.
4. Preserves the accepted F57 residual, channel-stamped FROM, committed-pivot generation authority, app-side-only epoch state, no-conductor/store-change route-back result, and Master+VP-only interface lock.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. Consumer confirmations remain held until m-1 pair approval.

## Verification

- `shasum -a 256 master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` -> `aa90fa45b9d8f37d5a820927e93eb3a30680ced9ec5111911c5881ec820a315d`.
- Read the exact addressed relay `master/relays/step3-mvp-design-m1/DESIGN-planner-20260716-045225.md` and parent dispatch `DESIGN-orchestrator-planner-20260716-041650.md`; routing and `DESIGN_DOC_ID` match.
- Checked ratified `master/STEP-3-MVP-AMENDMENT.md:8-28`, `:36-52`, `:79-98`, and `:103-124` for F57/F60/F64/F66, stage ordering, and acceptance wording.
- Checked locked m-1 s6 lifecycle `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:91-113` and the re-mint ruling `frank/.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-102208.md:19-59`.
- Checked current consumer bytes `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:58-127` and `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:11-80`.
- Checked live mechanism at `frank@502e06c`: `cmd/frank/main.go:603-621`, `internal/seat/binding.go:94-129`, and `internal/channel/server.go:182-193,244-309`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file and lane-root relay-lint.py exit 0
Next requested action: m-1.planner revises only F1-F3, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; consumer confirmation remains held.
