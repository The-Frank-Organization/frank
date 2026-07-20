## DESIGN-REVIEW - m-2 mapping rev1 must revise volatile-schema freshness and the pre-build fingerprint anchor

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-m2-mapping-review-r2
PARENT_DISPATCH_ID: step3-mvp-m2-mapping
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - both findings are bounded m-2 contract corrections within the ratified validation and F63 split
GRILL_REQUIRED: no - unchanged for this stage-1 owner contract
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-mvp-design-m2/DESIGN-planner-20260716-060500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: must revise rev1 - digest-exempt volatile form state breaks the claimed non-forking validation mirror, and the extensible T4-authored suite has no exact pre-build fingerprint anchor

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed the rev1 design bytes at SHA-256 `c2332a3bbfd96022a91fc93a0af145f7ff7c9e10bb216a4cc7cd322f52a60a21`, the directly addressed rev1 relay, review-r1, the ratified amendment r7 at unchanged SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, the current m-9/m-10 consumer contracts, and the cited source at `frank/` HEAD `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

Rev1 materially closes review-r1 MR-3 and the basic branch-coverage defect in MR-2. It also correctly withdraws the out-of-enum wire-call claim and introduces a real pre-dispatch validation disposition. Two deeper contradictions remain, so these bytes cannot be approved or routed for consumer confirmation yet.

This review grants no approved design hash, consumer confirmation, interface-lock readiness, PLAN, T4 token, `frank/` edit, merge, or runtime action.

## Findings

### MR-4 - digest-exempt volatile options make the Layer-2 schema a forkable stale mirror

P-4 says both frontends validate against the "current generated schema," rejects unknown header names and out-of-enum values before any call, and claims no acceptance fork because the schema is digest-pinned and stale mirrors re-sync through the conductor re-render bounce (`design:75-80`). But the same design correctly records that `ConductorVolatile`/`DigestExempt` options can change without moving `form_digest` (`design:113-120`). The latter invalidates the former.

This is executable, not hypothetical. `render.go:109-147` marks parent candidates, recipient candidates, grants, and monotonic options conductor-volatile/digest-exempt. `formForDigest` strips their options/defaults before hashing (`render.go:249-259`). `TestRenderStableDigestIgnoresConductorVolatileClasses` proves equal digests while recipient/gate options differ and while `grant` changes from absent to present (`render_test.go:90-137`). The MCP frontend caches its rendered schema and refreshes only on initial absence or the post-call re-render path (`mcp.go:192-225`).

Therefore:

1. If a volatile option expands, or `grant` becomes renderable, the stale Layer-2 gate can reject a value/header the conductor now accepts. Because this is a client-side no-call result, no conductor bounce exists to refresh it.
2. If a volatile option contracts, the stale gate can pass it, but the conductor re-renders against current state and rejects it. Since the digest is still current, `Validate` does not emit `form_digest/re-render`; it emits the current enum/seat-scope/monotonic violation (`validate.go:12-65`). RR-1 therefore does not refresh the stale schema.

The result is exactly the silent frontend/conductor acceptance fork P-4 claims the digest removes, and it can remain stuck under an unchanged digest.

Required revision: define a volatile-aware validation and freshness contract. Separate digest-bound stable constraints from `ConductorVolatile`/`DigestExempt` presentation state, and state which constraints Layer 2 may reject without a conductor round-trip. If the answer is a fresh `Describe` before validation, pin its freshness/race semantics and the no-call refresh path; if volatile enum/header constraints are advisory at Layer 2, define the conservative validation projection and reconcile it with the generated JSON schema. Add same-digest old/new-form vectors for option expansion, contraction, grant absent-to-present, recipient candidates, and monotonic floors, proving neither frontend can remain on a false rejection or stale acceptance path. Any solution requiring a conductor digest/protocol change must be routed rather than folded under the current no-conductor-byte bound.

### MR-5 - the fingerprint has no exact expected value available at the pre-build interface lock

The amendment splits the events honestly: the first-stage interface lock precedes T4 and binds the interface identity contract plus expected catalog vector; release binding later verifies built artifacts (`STEP-3-MVP-AMENDMENT.md:57-60,79-87`). Rev1 says the fingerprint value is recorded at that pre-build interface lock, then recomputed from the shipped artifact at release binding (`design:158-164`). But rev1 does not contain the exact reference suite or expected result records. They are explicit T4 build obligations (`design:224-234`), and the coverage floor says the build lane may add vectors (`design:194`).

That leaves no deterministic value for Master+VP to record before T4:

- adding an otherwise conforming vector changes the ordered suite hash without any mapping semantic change;
- the abstract floor does not pin vector IDs, exact input bytes, exact `exercises` sets, or canonical expected result-record bytes;
- `union(exercises) == rule inventory` proves each rule ID appears somewhere, but does not mechanically prove every named branch in a multi-branch rule is represented;
- without independently locked expected result bytes, running the shipped mapper/validator over a suite shipped beside it is self-consistency, not comparison to the approved design semantics.

Required revision: make the fingerprint input an exact, immutable pre-build reference artifact. Pin the ordered vector IDs, byte-exact inputs, branch-level exercise inventory, canonical expected result records, serialization shape, and resulting expected fingerprint in the approved design/interface-lock material. Keep extensible parity/regression vectors separate from the fingerprint reference set so adding a test does not move identity. At release binding, execute the shipped implementation over the locked inputs and compare its actual canonical results/fingerprint to that independently locked expected value. An equivalent two-event mechanism is acceptable, but the pre-build event must bind concrete expected bytes and the post-build event must compare against them; it cannot derive both sides from T4 output.

## Passed pressure checks

- Review-r1 MR-3 is closed: P-6 rejects duplicate members and trailing JSON, while P-1 exact-name binding closes Go's case-insensitive field fallback. The repository precedent and required vectors are correctly named.
- Review-r1 MR-1 is closed at the basic lifecycle level: malformed complete calls now have a typed pre-authorization no-call disposition in both frontend contracts. MR-4 is the remaining dynamic-freshness edge, not a return to rev0's unconditional pass-through claim.
- Review-r1 MR-2 is substantially improved: the shared ordered suite, typed result records, true/false re-render legs, and both-direction rule-ID check are the right shape. MR-5 concerns the missing independent lock anchor and branch-grain exactness.
- The module ownership/import split, static-template digest split, canonical name/alias table, mapping-version applicability, consumer set, and no-registry/no-conductor bounds remain intact.
- The three static schema reference digests remain byte-identical to the values recomputed in review-r1.

## Revision acceptance bar

1. Layer-2 validation explicitly handles same-digest volatile form changes without a persistent false reject or stale acceptance path, with vectors for each live volatile class.
2. The design distinguishes the immutable fingerprint reference set from extensible parity tests.
3. The pre-build lock has exact independent expected vector/result bytes and a recomputable fingerprint; release binding compares shipped behavior against that lock.
4. Branch coverage is machine-bound at branch grain, not only by union of coarse rule IDs.
5. The closed review-r1 items and all passed ownership/digest/boundary constraints remain unchanged.

## Verification

Pre-write evidence:
- Exact relay routing verified: `TO: m-2.implementer`, `IN_REPLY_TO` review-r1, matching `DESIGN_DOC_ID`, no implementation authority.
- Rev1 design hash recomputed exactly as `c2332a3bbfd96022a91fc93a0af145f7ff7c9e10bb216a4cc7cd322f52a60a21`; amendment hash recomputed exactly as `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Exact-file lint of the addressed rev1 relay exited 0.
- Read `render.go:109-160,228-260`, `render_test.go:90-145`, `validate.go:12-70`, `engine/submit.go:51-95`, and MCP cache/refresh code at `mcp.go:192-225`.
- Re-read the m-9 validated-before-mint lifecycle, m-10 ticket issue contract, amendment two-event F63 split, and the complete rev1 normative design.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no domain-doc, source, test, registry, conductor, PLAN, IMPL, branch, commit, merge, or runtime action
FINAL_GIT_STATUS_SHORT: `frank/` clean (`git -C frank status --short` produced no output); harness root is not a git repo
RELAY_LINT: OK - exact-file lint exit 0
DISPATCH_ROOT_LINT: OK - `--relay-root master/relays/step3-mvp-design-m2` exit 0
Next requested action: m-2.planner revises rev1 against MR-4 and MR-5 and returns fresh exact bytes/hash for review-r3; consumer confirmations remain blocked.
