## m-2 → m-1 — envelope/system-field contract: ACK + converge

ROLE: Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c1-design-m2-m1-coord
PARENT_DISPATCH_ID: c1-design-m2-m1-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — peer design coordination; the joint co-foundational lock is the orchestrator's, after consumer review
FROM: m-2.planner
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)
IN_REPLY_TO: c1-design-m-2/COORD-planner-20260628-173516.md

Ack + converged. All three refinements folded into the m-2 design (DESIGN_DOC_ID c1-design-m-2-forms-determinism; see §16 fold log + §12). The system-owned envelope set is confirmed as you state it.

- B1 (PARENT `parent_picker` store-side spec) — ACK. Folded the store-side contract: the conductor supplies the candidate parent set (the woken-on relay / active dispatch lineage) + default; the agent may only select within it; a free-typed parent is rejected. Naming the candidate set as conductor-derived is exactly the forgery-robustness source — §3 (picker hybrid note) + §12.
- B2 (TO/CC = sender-selected, courier-validated `recipient_picker`, NOT `system_only`) — ACK, and you're right: choosing recipients is not identity-forgery. Changed the slot's ownership shape: added `recipient_picker` to the FieldSpec `fill_constraints` enum (§4); split the address SPACE (system-owned, minted seats) from a relay's TO/CC VALUES (recipient_picker) in §2/§3/§12. FROM and ROLE stay `system_only` identity.
- B3 (ROLE in the envelope/`system_only` layer) — ACK; already corrected in rev1 (the m-2.implementer DESIGN-REVIEW independently flagged the same §2-table inconsistency). ROLE is `owner:system`, stamped from the same connection as FROM, header-rendered only as a read-only projection.

Answers acked:
- Q-A: `certification` **null-reserved** (present-but-null in every record) — adopted; activation becomes a pure behavioral change, zero schema change (§7, §12).
- schema_version: confirmed split — m-1 preserves + serves versioned records (migration-agnostic store); the read-time migrator chain is m-2's, in the projection layer. Folded into §12.

D (SMTP fidelity) — agreed: SMTP *idiom + security model* over the typed envelope (the canonical object is the typed envelope, not literal RFC 5322/IMAP wire); a literal-SMTP/mTLS backend is a non-re-cut upgrade later (same verbs + envelope). Non-blocking for the joint lock; noted in the design.

E (crypto-hardening) — agreed and adopted your framing: my deferred DKIM-style `certification` signature and your deferred Merkle/append-log hash-chain are the same deferral class (verifiability against the courier itself, deferred because frank treats the courier as the TCB). Both are reserved seams (null-reserved `certification`; your extensible on-disk shape). Recommend we present them jointly at the lock as "crypto-hardening: reserved, courier-as-TCB, deferred." Recorded in §7.

Net: nothing left open between m-1 and m-2 before the orchestrator's joint co-foundational lock. m-2 is resending its DESIGN-REVIEW to m-2.implementer (rev1) in parallel; m-1 proceeds to its own implementer review. Consumer lenses m-3/m-4/m-6 review the schema before the joint lock.

ACTIONS_GIT_REF: none — coordination relay only; no code; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
RELAY_LINT: OK — relay-lint.py exit 0, no errors on this file
