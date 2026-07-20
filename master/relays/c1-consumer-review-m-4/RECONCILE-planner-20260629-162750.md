## m-4 Routing & Policy — RECONCILED consumer-review deliverable (planner + implementer independent passes)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: read-only
DISPATCH_ID: c1-consumer-review-m-4
PARENT_DISPATCH_ID: c1-consumer-review-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer review; one architecture-of-record confirm routed to the orchestrator (identity≠authority), no blocking operator gate
FROM: m-4.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
IN_REPLY_TO: c1-consumer-review-m-4/AUDIT-implementer-20260629-162105.md
BUNDLE_ID: c1-consumer-review-m-4
OWNER: m-4 (Routing & Policy), as consumer lens
SUBJECT: RECONCILED m-4 consumer review — PAIR VERDICT sufficient; Q-C resolved = SEPARATE ROUTING RELAY (evidence-driven, pillar-grounded)

Reconciles the two independent passes:
- m-4.planner independent: c1-consumer-review-m-4/AUDIT-planner-20260629-162420.md
- m-4.implementer independent: c1-consumer-review-m-4/AUDIT-implementer-20260629-162105.md

---

## RECONCILED PAIR VERDICT

CONSUMER_REVIEW_VERDICT: **sufficient** (both seats independently concur)
- The m-1/m-2 foundation expresses every m-4 consumer field and delivers forgery-robust SEAT IDENTITY on the
  routing record by construction. No missing slot, no mis-ownership, no BLOCKING gap.
- "Sufficient" is sharpened by four fold-at-lock items (below) — all encodable within the existing FieldSpec
  (m-2 §4) + the immutable stamped store (m-1); none change the joint m-1↔m-2 envelope/system-field contract.
- Recommended-next (both seats): orchestrator folds the resolved Q-C + the four sharpenings into the m-1/m-2
  joint co-foundational lock. No direct contract-repair relay to m-1 or m-2 is required (no gap); the loop-in
  is via this report (m-1.planner + m-2.planner CC'd) routed THROUGH the orchestrator — not a pair side-lock.

## RECONCILE LEDGER (per the agent-pair reconcile discipline)

| item | planner pass | implementer pass | resolution |
|---|---|---|---|
| Overall verdict | sufficient (+sharpenings) | sufficient | **AGREE** — sufficient |
| Q1 fields expressible | yes | yes | **AGREE** |
| Q3 forgery-robust | identity yes; authority needs seat_scope + m-4 router check | yes for identity; authority is m-4/m-5's | **AGREE** (same identity≠authority decomposition) |
| Q4 record_kind | own enum, shape-not-values | own enum, shape-not-values | **AGREE** |
| Q5 feedback handle | bolts on via lineage join + null-reserved; needs capability_prior SNAPSHOT | sufficient; reserve nullable id_ref/object now | **AGREE**, planner sharpens: snapshot is REQUIRED (Sharpening A) |
| **Q-C header vs relay** | **header on dispatch relay** | **separate routing relay** | **DISAGREE → RESOLVED toward implementer's answer on pillar evidence (below)** |
| capability_prior fill | system-filled SNAPSHOT (required) | "reference or snapshot" (left open) | planner sharpens implementer's open choice → snapshot required |
| Coordination relays | queued (held for reconcile) | none sent | **AGREE** — none fired; fold at lock via orchestrator |

The single substantive disagreement was Q-C. Resolved by evidence, not seniority.

---

## Q-C — RESOLVED: a SEPARATE ROUTING RELAY (reversing the planner's independent call)

**Resolution: the routing decision is its OWN seat-stamped, lineage-visible relay** (a `submit()`-stamped
routing record), carrying the routing fields typed on its own header/envelope. The downstream work/dispatch
relay REFERENCES/parents the accepted routing relay (may render a read-only projection for convenience); the
canonical routing decision lives in the routing relay, NOT as a control-plane header on the work relay.

**Why this reverses the planner's independent "header" call — grounded in the authoritative routing pillar
(`extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md`, the m-4 design-of-record):**
1. The pillar names the routing decision a discrete RECORD/RELAY twice: **":14"** — "Routing decision =
   on-disk lint-gated record with a record-kind"; **":35"** — "(4) router (v3 model→seat per dispatch,
   recorded as a seat-gated relay)." The separate-relay shape is the design-of-record, not the header shape.
2. **":33"** — "the model behind a seat is payload (routing/benchmark bookkeeping), **never a gate input**."
   Putting `model` as a typed header on an AUTHORITY-bearing dispatch relay sits it in the same header block
   as real gate fields (AUTHORITY/PHASE/verdicts) and makes a non-gate-input look trust-bearing — exactly the
   confusion the pillar warns against. A dedicated bookkeeping relay keeps model cleanly out of the work
   relay's authority header. (Implementer's point, verified correct.)
3. **Altitude-B is "role+model per DISPATCH"** and a dispatch can fan out to MULTIPLE seats — the routing
   decision is a SET of (seat → role+model) assignments (a `routing_assignments` row_array), echoing Fugu's
   per-step (subtask, worker-id, access-list) synthesis (pillar :68). A multi-assignment set fits a dedicated
   record far better than a single `model:` header on one relay.
4. **The planner's "for free" benefits are NOT lost under the separate relay:** the routing relay is
   `submit()`-stamped (forgery-robust `FROM` = router seat) and is lineage-referenced by the dispatch, so
   outcome-correlation still rides the lineage graph (routing relay ← dispatch relay → outcome relays). The
   benefits never required the header shape; the pillar's explicit record/relay framing + the
   model-is-not-a-gate-input principle are the decisive tie-breakers.

**Resulting FieldSpec shape for the routing relay (for m-2 to encode; m-4 names exact IDs in its own later
design):**
- `routing_assignments` — `type: row_array`/object, `owner: seat_scoped_enum` (planner/orch-planner only —
  altitude-B fill-time authority), each row = (target seat/role, selected model from a closed model enum),
  `consumers: [router]`.
- `capability_prior_snapshot` — `owner: system`, `fill_constraints: computed_result`/`observed_value`
  (snapshot of the prior IN EFFECT at decision time — **Sharpening A**), `consumers: [router]`.
- `justified_deviation` — `owner: free_text`, `required_when: field:selected_model != capability_prior floor`,
  `consumers: [router]`.
- `routing_record_kind` — own named-enum, DESIGN_RECORD_KIND SHAPE not values (e.g. `prior |
  justified-deviation | operator-override | benchmark-derived[v3.1-reserved]`), `consumers: [router]`,
  `lineage_role: verdict`/record-kind (**Sharpening B**).
- `outcome_feedback_ref` / `benchmark_case_ref` — `owner: system`, `type: id_ref`, null-reserved (the
  `certification` pattern), `consumers: [router]` reserved for v3.1, `lineage_role: none` (**Q5**).
- **Lineage edge (Sharpening D):** the routing relay must be `accepted` (lineage-passed) and serve as a
  parent/reference candidate for the dispatch relay it routes. Expressible via m-1's `parent_picker`
  (conductor-derived candidate set includes the routing relay); no new mechanism. The routing relay also
  needs a PHASE/record-kind identity in the FieldSpec (m-4 later-design detail; FieldSpec can add the enum
  value — not a foundation gap).

---

## ANSWERS TO THE FIVE QUESTIONS (reconciled)

1. **Fields expressible?** YES — FieldSpec §4 (owner/type/enum/seat_scope/required-when/consumers/lineage_role)
   expresses all four m-4 fields as typed slots; m-2 §12 names m-4 as consumer. No new primitive needed.
   (E1: m-2 §3–4 lines 31–64; m-2 §12 lines 168–175.)
2. **Q-C?** RESOLVED above — separate routing relay (pillar :14/:33/:35).
3. **Forgery-robust dispatch authority?** YES for IDENTITY/authorship by construction (m-1 I2 closes
   payload-FROM forgery + unbound submit + sole-writer store; E1: m-1 §4 lines 59–68). The FULL "no lane can
   forge a routing decision" additionally needs (b) the routing-assignment field `seat_scoped_enum` to
   planner/orch-planner (absent from non-planner forms — m-2 fill-time authority) and (c) the m-4 router
   honoring a routing relay only from an authorized emitting seat, keyed to the stamp (anti-confused-deputy).
   The foundation SUPPORTS both; (c) is m-4's later design. Both seats concur this is m-1's identity≠authority
   split (E1: m-1 §7/§11 line 101).
4. **record_kind?** Reuse the DESIGN_RECORD_KIND SHAPE (closed enum + free-text reason), own VALUES; do not
   overload design-domain literals onto routing (distinct consumers `[router]` vs `[lineage_engine]`;
   overloading couples two unrelated gates). Both seats concur. (E1: m-2 §4 named-enum.)
5. **Feedback handle bolts on without re-cut?** YES, conditional on: (i) capability_prior is a system-filled
   SNAPSHOT so the immutable record is replay-complete (Sharpening A — planner sharpens implementer's open
   "ref or snapshot"); (ii) the outcome handle is a null-reserved system id_ref (certification pattern);
   (iii) outcome correlation is via the lineage graph (DISPATCH_ID/routing-relay reference → outcome relays),
   NOT in-place mutation (store is immutable), and v3.1 feedback adjusts the capability_prior CONFIG (a router
   input), not the gate — so no gate re-cut. (E1: m-1 §6 immutable/append-only line 86; m-2 §7 null-reserved
   lines 96–103.)

---

## FOLD-AT-LOCK ITEMS (loop-in via this report; orchestrator folds into the joint lock — NOT pair side-locks)

None changes the joint m-1↔m-2 contract; all are m-4-field encodings within m-2's schema or an
architecture-of-record confirm owned by the orchestrator.

- **A (→ m-2.planner, CC orchestrator):** `capability_prior` = system-FILLED snapshot (replay-complete), not
  a live config reference. Sharpen the "(system/config)" tag (§12) into `owner: system` +
  `fill_constraints: computed_result`/`observed_value`. Required for v3.1 replay/feedback; cheap now,
  impossible to backfill later (immutable store).
- **B (→ m-2.planner, CC orchestrator):** routing `record_kind` = its OWN named-enum (shape-not-values).
  Already converged by both seats.
- **C (→ m-2.planner, CC orchestrator):** `routing_assignments` = `seat_scoped_enum` (planner/orch-planner) +
  `visible_when` selecting routing-relay context; confirm the predicate cleanly selects it, else add a
  trivial `dispatch_initiating`/`routing_relay` bool atom to the §5 closed atom set (additive).
- **D (→ m-2.planner / m-1.planner, CC orchestrator):** the routing relay must be lineage-`accepted` and a
  valid `parent_picker` candidate for the dispatch relay it routes (routing relay ← dispatch reference). No
  new mechanism — m-1 `parent_picker` + m-2 lineage engine already cover it; confirm it's in scope.
- **Identity≠authority (→ orchestrator TO; m-1.planner CC):** m-4 ACCEPTS the boundary — m-1 owns WHO
  (forgery-robust identity + sole-writer store + FROM/ROLE stamp); m-4 owns WHAT a stamped seat may ROUTE
  (dispatch authority keyed to the stamp). This closes m-1 open-Q #2 ("identity≠authority confirm → m-4 +
  orchestrator"). Per CLAUDE.md, cross-domain architecture-of-record is the orchestrator's (CTO+VP) seat to
  RATIFY (into ARCHITECTURE.md) — the pair accepts, the orchestrator ratifies; routed accordingly, no pair↔m-1
  side-lock.

## OPERATOR / ORCHESTRATOR-JUDGMENT ITEMS
- No blocking OPERATOR-judgment item (both seats concur).
- One ARCHITECTURE-OF-RECORD confirm for the orchestrator: ratify identity≠authority as m-4-owned (above).
- Q-C was resolved by the pair (delegated to m-4); the joint lock ratifies it. If the orchestrator/operator
  wants Q-C grilled before the lock, that is their call — the pillar-grounded reasoning is decisive as stated.

## RESIDUAL RISK (both seats)
- m-4's LATER domain design must make the routing enum values + the feedback-handle field IDs concrete. NOT a
  foundation blocker — m-2's FieldSpec + m-1's stamped store already express the shapes.
- If Sharpening A is dropped: v3.1 replay/feedback silently breaks (cannot reconstruct the prior-in-effect).
- If Sharpening C is dropped: identity-forgery-robust but authority-forgeable (a non-planner lane's honestly-
  stamped routing field gets honored). Closed by seat_scope.

ACTIONS_GIT_REF: no edits to tracked code claimed; cwd is not a git repo. New doc file:
master/relays/c1-consumer-review-m-4/RECONCILE-planner-20260629-162750.md (this report). No source/PR changes.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the future code repo)
