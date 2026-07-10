## DESIGN — Slice-4 design dispatch (the wire-up): the shim + lifecycle hardening + the §7 record, against the LOCKED contracts; GRILL_REQUIRED

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s4-wire-design
PARENT_DISPATCH_ID: s4-wire-audit
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the grill (operator) is required pre-lock per GRILL_REQUIRED below; close gates stay the operator's downstream
GRILL_REQUIRED: yes
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: s4-wire.implementer, s4.orchestrator-reviewer, operator
SUBJECT: DESIGN dispatch — one design doc for the s4 slice (shim · lifecycle · §7 record · ops/usage surface), twelve reconciled-audit constraints binding; grill with the operator pre-lock; provisional-pending-guide sections marked; DESIGN-REVIEW goes TO your implementer per the lineage-gate discipline

**What this is.** The DESIGN phase for s4, following the reconciled paired audits (ledger
entry 2026-07-05 — full agreement, zero contradictions). You author the design doc at
`docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md`; your implementer is the
design-challenger (CC now) and the formal DESIGN-REVIEW addressee (design-review lineage: your
review request goes TO `s4-wire.implementer`; the approving DESIGN-REVIEW parents your
DESIGN relay; the later gated PLAN locks `DESIGN_LOCK_ID` and parents the approving review).
Design against the LOCKED contracts — escalate spec problems to me; never self-amend.

### Hard constraints (from the reconciled audits — binding design inputs)

1. **The guardrail surface stays exactly `submit`/`project`/`read`.** The shim is
   affordance, never authority: MCP front (initialize handshake, JSON-Schema `inputSchema`,
   `tools/call` result shape, stdio JSON-RPC 2.0 framing) ↔ `channel.Client` back
   (PROMOTE — it is the shim's conductor-side half; server.go:305-471). The conductor's
   bespoke dialect stays private behind the shim. No new seat-facing verb anywhere.
2. **Second-connect [VP-W1, source of record = the s4-dispatch]:** a per-credential
   active-channel index, engine-internal, exactly two legs — reject-active-duplicate
   (typed, path-free, e.g. `auth:channel-active`) and proven-dead recovery (prior conn
   left `s.clients` on kernel close). The serverConn retains (a hash of) its resolved
   credential for the index; NO binding-table shape change (a needed shape change = hard
   stop + escalate, the S2 sanction precedent); NO new seat-facing surface. Wedged-but-alive
   host = reject regime + documented operator remedy. Live supersede / rotation /
   re-mint-supersedes = locked-contract touch = escalate. Any text claiming second-connect
   already works is false (both audits).
3. **Per-recipient nudge grain (locked m-7 §8.3) — provisional-pending-guide Q1:** design
   the per-seat wake (per-connection seat-scoped push after auth); retire the global
   `pending` queue (per-seat recovery nudge on that seat's reconnect); mailbox stays truth;
   poll-hint-first posture is FIRST-CLASS (client push-drop at 16 frames is real —
   server.go:464-468; hosts may hide custom notifications from the model).
4. **The §7 config-change record — provisional-pending-guide Q2/Q3:** a first-class
   mutation class through the one-pivot discipline on an EXISTING store (never re-genesis);
   phase-0 learns the genesis→config-change digest CHAIN (today `ValidateGenesis` is a
   static compare — genesis.go:104-118); `f11Classes()` + the applicability map + crash
   legs gain the class (the LOCKED F11 list :172 already names it); the `record_kind`
   token + chosen provenance ride the m-1 fidelity packet (s3-scope-q1 condition 4);
   authorship recommended = operator-channel submit (guide Q3 confirms); discharge
   `OI-S3-CONFIG-CHANGE` through the live owed mechanism on the real store; the re-render
   bounce ("superseded rendered form bounces re-render; re-rendered form succeeds") is the
   drift leg of the same round-trip.
5. **I-PH extends to the shim's OWN surfaces [VP-W3]:** the shim knows the socket path +
   credential by config, and Go's default dial error carries the socket path verbatim —
   the shim scrubs/types every self-generated error (dial failure, auth failure, reconnect
   notices, diagnostics). The exit-gate class list is the fixture matrix: tools/list
   descriptions · input schemas · tool-call results · notifications/poll hints · reconnect
   errors · credential-failure errors · shim diagnostics. Conductor-side surfaces measured
   clean (live scan, zero hits over 12.7 KiB) — keep them that way.
6. **Frame-transport ceiling (audit F-W1) — provisional-pending-guide Q4:** the silent
   >64 KiB connection death (default scanner both sides) must become typed or bounded;
   lean = raise to a config-pinned bound + typed path-free refusal; no chunking this slice
   unless the guide directs.
7. **Custody posture, stated honestly (D5):** shim credential via env-var where the host
   allows; where not, the D5 note verbatim. Document plainly: mint prints to stdout;
   `-operator-submit -credential` is ps-visible; the binding table is plaintext-at-rest
   (0600/0700); **no in-band rotation/revocation exists — a compromised credential means
   admin-time surgery** (m-1 §13.3 carry; in-band rotation = escalate). Confusion-resistant,
   not theft-proof — the honest claim, everywhere credentials are documented.
8. **Ops surface:** start/stop/status conventions; team-store + socket-path conventions
   with the SHORT-SOCKET-PATH rule (darwin AF_UNIX ≈104-byte cap — the conductor's own
   bind fails on deep paths, audit F-W3; pre-flight length check = cheap hardening
   candidate); the minting workflow end-to-end. Usage posture: the store IS the usage
   record; document the read path (at most a trivial read-only stat over `project`); any
   aggregation/analytics = s5 OUT-touch.
9. **Honesty [VP-W2]:** every s4 E3 claim surface — design doc, shim README/usage doc,
   ops doc, the exit-gate evidence record, MCP-visible shim self-descriptions — carries
   "transport/provenance only; done-state and `record_integrity` remain `self_reported`
   until Step-2 observe". Shim tool descriptions must not over-claim ("files a governance
   relay", never "verifies work"). The root README fresh-store sentence goes STALE when
   the §7 record lands — flag the in-fence update as a PLAN-time ASK (S1 ASK-1 precedent).
10. **E2 floors:** full battery green, zero regression, enum byte-exact
    `{accepted, rejected, held}`, registry enumeration still exactly three tools.
11. **Reuse, never rebuild:** the audits' promote inventory is binding — channel.Client,
    the S3 describe render (server side whole; the shim owes Form→JSON-Schema presentation
    + drift surfacing only), the crash harness, the owed mechanism, store.Init/phase-0 as
    the §7 mutation target. `test/seatproc` is a stub, not substrate.
12. **Fresh-eyes findings F-W1..F-W5 dispositioned in the design:** F-W1/F-W3 as
    constraints 6/8; F-W2 as constraint 3's poll-first posture; F-W4 — the per-recipient
    redesign RETIRES the global queue rather than patching it; F-W5 — the zero-ID dialect
    ambiguity pinned in the shim's dialect handling.

### Grill (GRILL_REQUIRED: yes — run pre-lock, operator required)
Trigger profile (real on its own, the S2 lesson): still-open work at medium tier;
cross-domain boundary contracts (m-1 identity/credential/custody, the m-7 engine seam,
host-runtime integration); hard-to-reverse decisions (the §7 on-store mutation shape +
`record_kind` + phase-0 chain semantics; the second-connect regime; the shim dialect +
custody posture); multiple downstream choices hanging on the guide thread. Grill agenda
floor: the §7 record shape (payload embeds config bytes? chain-walk semantics; authorship),
the second-connect reject/recover regime + operator remedy, the shim custody posture +
host-config shape, the frame-ceiling bound. **Grill fence:** no re-opening of c1–c6
operator-grilled locks; guide-thread answers (s4-guide-q1, routed in parallel) enter as
resolved rows, never re-asked; an answer that would amend locked text escalates to master.
GRILL_LOCK_ID folds into DESIGN_LOCK_ID pre-lock; no design lock, no
design-review-consumed-toward-PLAN, no PROCEED-TO-PLAN until the GRILL_LOCK exists and
guide-answer deltas are folded. Drafting proceeds provisionally meanwhile.

### OUT (escalate, never design-into)
Consumer schema content (s5) · observe/evidence (Step-2) · routing execution (Step-3) ·
TUI/email UX (Step-4) · federation (zero pre-work) · external send/away-bridge ·
steer/interrupt beyond host-native · authority replacement (transport only) · in-band
credential rotation/supersede · socket-dialect rewrite toward MCP (the composite is the
sanctioned realization, pending guide Q6).

### Deliverable
The design doc (dated, in `designs/`), revision-logged; design-review chain per the lineage-gate discipline
(request TO your implementer; approving review parents your DESIGN relay); the GRILL_LOCK
at a numbered section; completion SITREP to me when the approving review + grill both
exist. Boundary contract fields per surface (the shim's host boundary; the §7 record's
store boundary; the lifecycle changes' channel boundary). Not authorized: PLAN, IMPL, any
code, any spec edit; PROCEED-TO-PLAN comes from me after reconciliation.

ACTIONS_GIT_REF: wrote this dispatch + an INDEX row (relay substrate, git-untracked); no frank/ code or spec edits.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ ledger commits current; code surface = tag s3-close exactly)
Next requested action: operator hand-relays to s4-wire.planner (and the s4-guide-q1 thread to m-7.planner); design + grill proceed; completion SITREP returns to me.
