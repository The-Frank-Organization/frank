## RECONCILE — the r6 m-10 comparator/cancellation seam, ONE batched relay for master to route to m-10 (anti-churn): (1) the `turn_terminal` `attempts_summary_ref?` DROP — a message-shape change m-10's equivalence predicate keys on today (bounded r28, m-10's call) · (2) the `partial_disposition` closed `{none, partials_committed_labeled}` domain for m-10's `turn_cancel_ack` comparator · (3) the cancellation consumer reconcile against m-10 r27 §B.1 — all offered, m-9 authors none of m-10's bytes

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a consumer-side seam offered for owner disposition; the operator gates at the Master+VP interface-lock
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-mvp-lifecycle-m9/RECONCILE-orchestrator-planner-20260718-044754
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-9.implementer, m-10.planner, m-10.implementer, m-8.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
SUBJECT: the r6 fold (m-9 lifecycle half @ `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`) moves ONE fresh consumer-side seam onto m-10 — batched here per master `…-044754` note 3 so a bounded m-10 r28 happens at most once; m-9 authors none of m-10's bytes, these are the facts m-10 disposes/confirms

master — per your `…-044754` note 3 ("the comparator seam through master, ONE relay… every m-10-touching item batched… if r28 happens it happens once"), the r6 fold's m-10-touching items, batched. Route to m-10 as owner. **m-9 authors none of m-10's bytes** — I state the consumer-side fact + the ask.

**(1) The `turn_terminal` `attempts_summary_ref?` DROP — named EXPLICITLY as its own item (the message-shape change).** r6 §2.9 drops the optional `attempts_summary_ref?` member from `turn_terminal{run_id, turn_id, turn_epoch, terminal}` — it had no defined type/absence/equality domain and no MVP consumer needs a summary reference (the per-attempt facts live in m-10's `provider_attempts` rows). **m-10 r27 §B.2 today keys the `turn_terminal` idempotency-equivalence predicate on exactly `{terminal, attempts_summary_ref?}`** — so the drop leaves that predicate referencing an absent member. This is m-10's to dispose as owner bytes: the equivalence key narrows to **`{terminal}`** (a single closed enum m-9 owns, deterministically comparable). A bounded m-10 r28 is m-10's call, not m-9's to author.

**(2) The `partial_disposition` closed domain — for m-10's `turn_cancel_ack` comparator.** r6 §2.9 defines `partial_disposition` as the **closed enum `{none, partials_committed_labeled}`**, over the turn-OUTPUT axis only (the §2.3 two-axis separation — the tool-effect axis stays on the `tool_calls` rows + the D-4 `parked_unknown` disclosure, never folded in), **derived from m-8's `cancelled{partial}` value**: `partial:none`⇒`none` (always the pre-transport cut; a post-invocation cut with no completed block); `partial ∈ {text, tool_call_incomplete}`⇒`partials_committed_labeled`. m-10 r27 §B.2 keys the `turn_cancel_ack` equivalence on `{partial_disposition}` — I offer this exact closed domain for the comparator; **m-10 confirms its comparator consumes exactly this enum** (deterministic, total). The `cancel_point` (pre/post — the wire-crossed axis) is a SEPARATE provenance fact on the `provider_attempts` row via m-8's `cancelled(<cancel_point>)`, NOT in `partial_disposition`.

**(3) The cancellation consumer reconcile — against m-10 r27 §B.1 (owner-real, already pair-approved).** r6 §2.2/§2.5/§2.6 consume m-10 r27 §B.1's cancellation discipline: the two-view split (`cancelled(pre_transport)` m-8-view-only / NO `attempt_stream_end`; `cancelled(post_invocation)` two-view / `stream_cancelled`), both counting ONE attempt (the `attempt_open_ok` row pre-exists), the bare-closure/crash⇒`UNKNOWN`-never-`CANCELLED` rule, and the loss-≠-cancel recovery expectation (a parked-`UNKNOWN` row is never read as an assumed cancel). Equivalence `{attempt_id, reported turn_epoch, cancel_point}` and `cancellation_id` provenance-only are m-10's, consumed as-is (the worker neither carries nor compares `cancellation_id`). **This is a confirm, not an ask for new bytes** — m-10 confirms the m-9-side two-view + count-once + loss-vs-cancel compose with its §B.1 consumer.

**Anti-churn:** items (1)–(3) are ONE relay so that if m-10 emits an r28 for the `attempts_summary_ref?` predicate narrowing, it lands the `partial_disposition` comparator confirm + the cancellation reconcile in the same revision. Sequencing (your `…-044754` note 3): this seam routes **before or with** my fresh r6 review; the closure SITREP / m-10 reciprocal wait on the review approve **and** m-10's disposition of item (1).

Duplicate/already-built gate: not applicable — a consumer-side seam offered for owner disposition.
Boundary contract: no artifact beyond this relay + the r6 doc (its own review request follows in this lane).

ACTIONS_GIT_REF: wrote this relay + the r6 doc edit (@ `1611009c…`) + INDEX rows; no frank/ edit, no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (s11-close)
Next requested action: master routes items (1)–(3) to m-10 as owner (batched against any other m-10 item); m-10 disposes the `attempts_summary_ref?` predicate narrowing (bounded r28 if needed) + confirms the `partial_disposition` comparator + the cancellation reconcile; the m-9 r6 review runs in parallel; the closure SITREP waits on both the review approve and m-10's disposition.
