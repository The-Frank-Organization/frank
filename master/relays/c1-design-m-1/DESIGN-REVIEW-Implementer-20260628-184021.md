## Team m-1 - Trust & Identity: DESIGN-REVIEW

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c1-design-m-1-review
PARENT_DISPATCH_ID: c1-design-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - design review only; no implementation authority
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, operator
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c1-design-m-1/DESIGN-planner-20260628-180828.md

### Verdict

APPROVE. The design satisfies the m-1 DESIGN-REVIEW gate for Trust & Identity. It gives a concrete target
entity (the trusted conductor / sole-writer relay store), names the boundary contract with m-2/m-3/m-4/m-6,
states the operator-grilled transport decision, proves the strength claim conditionally on deployment
invariants, and preserves the non-re-cut path from Step-1 to standalone attestation.

This is review-only. I made no source or `pcode/` changes.

### Scope checked

Reviewed:
- `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md`
- m-1 DESIGN request `c1-design-m-1/DESIGN-planner-20260628-180828.md`
- c1 reconcile ledger and m-1/m-2 coordination thread (`COORD` 172553 -> 173516 -> `COORD-ACK` 174201)

Skill note: the `design-grill` wrapper loaded, but its referenced vendored file
`vendor/mattpocock/grill-me/SKILL.md` is absent in this installation. I therefore reviewed against the
wrapper's required grill-lock surfaces and the design's embedded `GRILL_LOCK_ID: c1-design-m-1-grill`;
this did not block the review because the design records the operator decisions, rejected alternatives, and
remaining non-blocking open questions explicitly.

### Findings

1. **Lineage and target entity pass.** `DESIGN_DOC_ID: c1-design-m-1-trust-identity` is present and matches the
   review request. The design identifies the target entity as the trusted courier/conductor that is sole writer
   of the relay store and stamps `FROM` from the inbound connection, not from lane-authored payload. This is the
   correct m-1 artifact to lock before consumer designs depend on it.

2. **I1/I2 proof is adequate and honestly conditional.** The design's strength claim rests on DI-1 through
   DI-4, then proves I1 store isolation and I2 channel isolation by enumerating payload lie, victim credential,
   unbound new connection, and replay/borrowed-handle attacks. The important point is not over-claimed: if DI-2
   cannot be realized, the claim degrades to operator-attested/confusion-resistant instead of
   by-construction. That conditional claim is acceptable for DESIGN.

3. **TOCTOU around `submit()` is not a design blocker, but must be preserved in PLAN.** The API sequence resolves
   the seat, fills system fields, runs the inline pre-send gate before append, then appends and delivers. That is
   sufficient as an architecture surface. PLAN should make this one conductor transaction: the resolved seat,
   validated envelope, m-3 observe result, append, index write, and mailbox projection must not re-enter
   lane-controlled state between validation and append.

4. **Credential reuse / remint is a PLAN acceptance detail.** The design closes credential theft/guessing/replay
   under DI-2 and DI-3, and `mint_seat` is conductor-only. PLAN should add generation or epoch semantics: a
   reminted seat credential must invalidate any prior credential or bind records to a single active credential
   generation so stale credentials cannot survive a seat recycle. This is a lifecycle detail, not a change to the
   four verbs or the design lock.

5. **Confused-deputy inline-gate placement is bounded correctly.** The design keeps m-1 ownership to identity +
   store and routes authority to m-4/m-5/orchestrator. The m-3 pre-send hook is positioned before append. PLAN
   should keep m-3 as an observer/validator that cannot author `FROM`, `ROLE`, `PARENT`, `relay_id`, or mailbox
   delivery effects. With that constraint, the inline gate does not become a second identity authority.

6. **Mailbox-write side channel is closed at the design boundary.** Mailboxes are conductor-written projections;
   lanes read through `project()`. That avoids reproducing the claude-code inbox gap where a lane can write
   another seat's inbox. The exact mailbox backing can remain a PLAN choice because both named options satisfy
   the m-1 invariant if the conductor is sole writer.

7. **`parent_picker` and `recipient_picker` are correctly split.** The m-1/m-2 coordination is folded into the
   design: `PARENT_DISPATCH_ID` is selected only from a conductor-derived candidate set; `TO`/`CC` are
   sender-selected but validated against the minted address space. This prevents parent forgery while avoiding
   the false requirement that the conductor invent recipients. Misdelivery policy and authority ceilings remain
   m-6/m-4/m-5 concerns, not m-1 identity.

8. **Reuse-vs-rebuild is balanced.** The design reuses v2.8.8 relay paths, append-only index discipline, and
   claude-code inbox projection as projections of a typed envelope, while removing lane-authored markdown from
   the trusted submit path. That ports the useful shape without copying the self-asserted identity failure.

9. **Non-re-cut path passes.** Step-1 minted credentials and later SO_PEERCRED/mTLS/SPIFFE swap only the
   `lookup(connection) -> seat` attestation backend. The public verbs, typed envelope contract, store shape, and
   stamp semantics stay invariant.

### Acceptance refinements for PLAN

Carry these into the eventual PLAN/E2 criteria without reopening DESIGN:
- DI-2 realization must be tested as an isolation property, not asserted: a sibling lane must fail to read or
  use another seat's credential/connection.
- `submit()` must be specified as an atomic conductor-owned commit path from seat resolution through append and
  mailbox projection.
- Seat credential lifecycle must define generation, rotation, and revocation semantics for remint/recycle.
- The m-3 hook must be unable to mutate system-owned identity/lineage fields or perform mailbox writes.

### Residual risk

The only material residual is the same one the design names: Step-1 DI-2 is a load-bearing infra choice. Until
the orchestrator chooses a concrete lane/credential isolation mechanism at PLAN, the by-construction claim is a
design proof obligation rather than an E2 fact. That is acceptable for this phase because the design makes the
claim conditional and gives the honest fallback label.

### Disposition

`DESIGN_REVIEW_VERDICT: approve`

This approval is for the m-1 pair review boundary. The joint m-1/m-2 co-foundational lock still needs the
orchestrator's integration and the m-3/m-4/m-6 consumer lenses named in the dispatch.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and updated master/relays/INDEX.md; no source or pcode edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (git status --short returns fatal at /mnt); pcode git status --short is clean.
RELAY_LINT: OK - relay-lint.py exit 0 on this file; relay-root mode still reports the standing INDEX.md header-field errors.
