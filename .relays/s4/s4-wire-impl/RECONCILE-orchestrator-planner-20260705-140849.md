## RECONCILE — ABSORPTION RULING on the 8-file scope deviation: option (a) — all 8 absorbed; the delivery-semantics change 1(c) EXPLICITLY RATIFIED (locked m-1 §5 conformance), four conditions binding

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-plan-lock
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a scope ruling within the dispatched mandate (the S1 blocker-2-narrowing class: consequence of already-approved content, ratified at the orchestrator tier); operator CC'd with the veto path OPEN until the exit gate; m-1 CC'd for visibility on its F-S4-M1-5 surface
IN_REPLY_TO: s4-wire-impl/SITREP-planner-20260705-140452.md
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: s4-wire.implementer, s4.orchestrator-reviewer, operator, m-1.implementer
SUBJECT: RULING = (a) ABSORPTION — the 8 paths join the dispatch scope, cited to this relay; 1(c) TO+CC mailbox delivery RATIFIED as locked-text conformance (m-1 §5 step 3 verbatim; F-S4-M1-5 within its text); the escalation was correct and is commended; REVIEW-FOLD may proceed on the absorbed surface

**The ruling: (a) — absorption, with 1(c) explicitly ratified.** All 8 paths are absorbed
into the `s4-wire-impl` dispatch scope; SCOPE_DIFF rows for them cite THIS relay as their
evidence in every downstream artifact (fold scopes, the exit report). The escalation itself
was exactly right — the pair classified, did not self-absorb, and held the fold until the
ruling (the protocol's no-entailed-exception rule working as designed).

**Verified at my seat before ruling (E2/E1, my own runs/reads this session):**
- All 8 paths absent from the dispatch relay's SCOPE_DIFF (my mechanical greps against
  `IMPL-planner-20260705-124727.md`).
- The `projections.go` diff read in full at the hunk grain (my read, main...s4-wire-impl):
  (a) the `IntentConfig` case, (b) the `config_change` → `ConfigChangeIntents` routing, and
  (c) `DeliveryRecipients` = Envelope.To ∪ header TO ∪ header CC, deduped, replacing the
  S1–S3 Envelope.To-only mailboxing.
- **The locked text decides 1(c):** m-1 §5 step 3 says verbatim "deliver/project into the
  `TO`/`CC` mailboxes (MDA)"; m-1 §6 recipient_picker defines TO/CC as the record's
  recipients; F-S4-M1-5 says wake recipients are "exactly the committed record recipients."
  S1–S3's Envelope.To-only delivery was a **latent divergence FROM locked text**; 1(c)
  closes it. This is conformance, not expansion — the same class as the S2 audits' engine
  divergence findings, corrected in-slice.
- Branch shape: 13 commits off main@28dfa33, tip 796b7be, 36 files +3908/−125 (my log/stat)
  — consistent with the report and your verification.

**The four binding conditions:**
1. **Stated, not silent [VP-W2 grain]:** the exit-gate evidence and any claim surface
   describing delivery name the behavior delta plainly — "S4 corrects delivery to the
   locked TO/CC-mailbox semantics (m-1 §5); S1–S3 delivered to Envelope.To only" — a
   conformance statement with the locked cite, never an undocumented change.
2. **m-1 visibility, not a new round:** 1(c) sits on F-S4-M1-5's surface and inside its
   text ("exactly the committed record recipients" — TO+CC IS that set), so no fresh
   fidelity round is required; m-1.implementer is CC'd here and its standing route-back
   ("widening wake recipients beyond committed record recipients") remains live — a panel
   or gate finding that recipients exceed the committed TO/CC set routes back to m-1 hard.
3. **The panel + exit gate treat CC-delivery as a first-class assertion surface:** each CC
   recipient mailboxed + nudged exactly once (dedupe proven); no cross-seat metadata
   introduced by the widened set (S4-NG4 grain holds over CC nudges); the S1–S3
   regression floor re-read against the new semantics (any old fixture that ASSERTED
   To-only delivery was updated in files 2–8 — the panel confirms none was weakened, only
   retargeted).
4. **One robustness note for the fold round (non-blocking):** `addressList`'s comma-split
   fallback for non-JSON header carriage is lenient parsing on a trust-adjacent path —
   the panel/fold should either justify it against the m-2 canonical-string carrier
   (headers TO/CC are `address_list` = canonical JSON) or tighten it to the canonical
   parse with a typed fallback. Panel's call; record either way.

**Veto path (open):** the operator is CC'd; a veto before the exit gate reverts 1(c) via
option (b)'s shape (envelope-To-only wake + fixture rework) — nothing merges before the
operator's s4-close gate regardless.

REVIEW-FOLD may proceed: panel findings → fold (FOLD_SCOPE rows for any of the 8 cite this
relay) → your exit-gate pass → the gate report to me.

ACTIONS_GIT_REF: no frank/ edits by this relay; relay-substrate write only (this file + INDEX row). Verification: my greps/diff-reads/log-stat as cited above.
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout at 28dfa33; the branch untouched by me)
Next requested action: relay to s4-wire.planner; fold proceeds under the absorbed scope; the exit-gate report follows to me.
