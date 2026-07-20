## DESIGN-REVIEW - m-1 implementer grill of s6 transport amendments

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review only; the parenting fork still returns to master for operator grill
GRILL_REQUIRED: no - this relay is the pair adversarial grill result; Section A does not lock here
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s6-design/DESIGN-planner-20260706-183326.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-2.planner
SUBJECT: m-1 s6 amendment grill approves packet forwarding and locks Sections B-D with fixture carry-forwards

VERDICT: approve. Section A is approved-to-forward as a decision packet only and does not lock here. Sections B, C, and D are approved and may be reported locked to master, subject to the fixture/carry-forward constraints below.

## Basis

- Incoming design relay: `master/relays/s6-design/DESIGN-planner-20260706-183326.md:18-69`.
- Amendment doc: `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:13-76`.
- S6 dispatch and VP handoff constraints: `master/relays/s6-design/DESIGN-orchestrator-planner-20260706-180315.md:23-44`; `master/relays/s6-design/RECONCILE-orchestrator-reviewer-20260706-182326.md:28-32`.
- Transport findings: `master/TRANSPORT-FINDINGS-2026-07-06.md:20-29`, `:36-54`, `:66-86`.
- Locked m-1 design anchors: `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md:99-150`, `:225-270`.
- S3 parent-picker carry-forward being superseded by the fork if Option A wins: `frank/.relays/s3/s3-fidelity-m1/SITREP-implementer-20260704-184437.md:45-57`.
- Adjacent S6 seam docs: `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:52-58`, `:68-86`; `master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md:61-80`, `:104-117`.
- Code/archive probes: `frank` HEAD `7e5c527`; `frank/internal/lineage/lineage.go:170-184`, `:346-371`, `:432-439`; `frank/cmd/frank/main.go:255-304`, `:389-402`; `frank/internal/store/store.go:60-96`, `:151-172`; `frank/internal/store/projections.go:93-142`; archived s5 store shows rejected relay ids present in mailboxes despite current code not adding new rejected default intents.

## Section A - parenting fork packet quality

VERDICT: approve-to-forward. No m-1 lock is granted by this review.

The packet honestly states the failed locked text: current `PARENT_DISPATCH_ID` remains a lane-selected `parent_picker` in the locked m-1/m-2 field contract, and live code still renders active-lineage candidates from `WokenOn`, `ActiveDispatch`, and same-dispatch accepted records. It also names the real supersession cost: Option A replaces locked Section 6 `parent_picker`, supersedes my S3 five-point active-lineage candidate derivation, and re-homes Sharpening-D from candidate-set membership into hint/bookkeeping semantics.

Option A is a defensible recommendation to send to master: it removes the render/submit snapshot race from anchoring by moving the anchor stamp to the same conductor-owned class as FROM/ROLE/relay_id/timestamp. The claim ladder is not overstated if the packet keeps the distinction it already names: anchoring becomes conductor-computed transport; class-lineage gates still bounce when authority lineage is missing or wrong.

No fourth option is required before forwarding. The only unresolved branch is already exposed inside Option A and must be explicit in the operator grill: invalid or unprovable `parent_hint` fallback versus hint-reject/hold for semantic precision. My recommendation for the grill packet is to ask that as the first operator decision, because it is the one cost that changes author experience. I would not block forwarding to split it into a formal Option D.

Section A forward conditions:
- The packet to master must preserve "packet only, no lock before operator grill."
- The grill must explicitly decide hint fallback versus hint rejection/hold, semantic precision loss on multi-target authoring, and Sharpening-D re-home.
- If Option A wins, the amendment fold must rewrite the old `parent_picker` text, not leave parallel parent authorities behind.
- If any branch keeps lane-selected parent values, it must also carry a no-livelock proof for the archived F11 traffic pattern; otherwise the branch fails the finding.

## Section B - F10 projection and anchor hygiene

VERDICT: approve - locks.

The split is correct. The anchor rule was already locked as an accepted-graph property; a rejected or held record must never be anchor-eligible. The new amendment needed here is the default projection rule: `project()` returns accepted records addressed to the caller by default, while rejected/held records remain terminal audit records reachable only through explicit audit/read paths.

Carry-forwards for the lock:
- Default `project()` must filter by canonical record delivery state, not only by future mailbox write behavior. The archived s5 store already has rejected relay ids in mailbox files, so the fixture must prove default projection filters an existing polluted mailbox/rebuilt archived store as well as future commits.
- Turn context must be derived from the same accepted-only default view or from accepted-graph state directly. It must not use the raw last mailbox line as `WokenOn`.
- The explicit audit projection is a parameter on `project`, not a fourth verb. Its eligibility and output shape must stay path-free and must not become the default seat inbox.
- `read()` of a rejected record by id is approved for the author and approved audit/operator surfaces. If the build keeps broad authenticated `read(relay_id)`, it must not silently recreate the default projection flood or a rejected-record anchor source.

## Section C - F17 waiver scoping and retraction

VERDICT: approve - locks.

The amendment closes the dangerous current semantics: live code accepts any prior accepted operator record with non-empty `ORCH_REVIEW_WAIVER`, run-wide and forever. Scoped waivers plus append-only `waiver_retraction` preserve immutability while restoring prospective governance.

Carry-forwards for the lock:
- Scope is required and first-class: target relay, dispatch, record-class by dispatch, or explicit run-wide. Unscoped legacy waivers may read as run-wide only until retracted; do not reinterpret historical accepted records.
- Retraction is a new accepted record with an id_ref to the waiver it retracts. Retraction never mutates the original waiver and never retroactively invalidates records accepted under a then-live waiver.
- Effective waiver state is evaluated in commit order over accepted records: accepted waivers minus accepted retractions at gate time.
- Optional bounded validity must be deterministic and store-derived. Do not use wall-clock or host-local TTL expiry as waiver semantics; a bound must replay the same way during recovery and on a rebuilt store.
- Waiver and retraction remain operator-only at the fill-time floor. Widening later is a registry/config amendment with m-1 route-back.

## Section D - F14 store-lock invariant

VERDICT: approve - locks.

The invariant is in the right domain and correctly split from m-7's runtime enforcement: m-1 owns "at most one live conductor has write authority over a store root"; m-7 owns acquisition, refusal, takeover choreography, and fixtures. The doc also keeps the D4/D5 claim honest: this is accidental/confused double-serve protection, not a malicious same-uid containment claim.

Carry-forwards for the lock:
- "Store root" must mean the canonical store-root identity, not a raw path string. The build must not allow two live claims through symlink, relative-path, or equivalent aliasing of the same root.
- The claim is acquired before phase-0 recovery and held for the serve lifetime. A loser refuses to serve; m-7's doc correctly strengthens that to reads included.
- Takeover requires proof that the prior holder cannot still write. A stealable live-writer claim or bare pidfile heuristic alone violates I1-P.
- The final integrated disposition must cite both halves per VP-W2: this m-1 invariant plus m-7's runtime fixture obligations.

## Required fixture obligations to carry forward

- F11 replay: archived dogfood traffic pattern completes without parent-anchor livelock under the grilled fork branch.
- Parent hint: honored and fallback/reject/hold behavior matches the operator-grilled branch, with audit-visible `parent_hint_honored` or equivalent.
- Projection: default `project()` filters rejected/held records from a polluted archived mailbox; explicit audit scope returns them; accepted records still deliver normally.
- Anchor: rejected and held records are never `WokenOn`, active-lineage candidates, or class-lineage parents in any interleaving.
- Waiver: scoped waiver passes only in scope; retraction re-arms prospectively; pre-retraction accepts stand; non-operator waiver/retraction is absent at fill-time and rejected on submit.
- Store lock: two-conductor race yields exactly one server; alias paths to the same root do not bypass the claim; kill -9 holder permits takeover only by the chosen proof-of-death mechanism; loser refusal is typed and path-free.

## Closeout

Sections B-D may be reported as m-1 pair-approved locks for master integration. Section A may be forwarded to master as the m-1 decision packet and remains unlocked until operator grill and the resulting GRILL_LOCK.

ACTIONS_GIT_REF: wrote `master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-184120.md` and appended `master/relays/INDEX.md`; no `frank/` edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short` is clean.
RELAY_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-184120.md`
