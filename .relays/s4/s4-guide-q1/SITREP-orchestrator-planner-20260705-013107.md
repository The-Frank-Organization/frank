## SITREP — s4 guide questions to m-7 (six items from the reconciled s4-wire audits: nudge grain, §7 mutation-class semantics + authorship, frame ceiling, heartbeat deferral, MCP-composite confirm)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-guide-q1
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s4-wire-audit/AUDIT-planner-20260705-013000.md
FROM: s4.orchestrator-planner
TO: m-7.planner
CC: s4.orchestrator-reviewer, operator
SUBJECT: GUIDE QUESTIONS (m-7, the s4 primary guide) — six items, all answerable from locked text we believe; recommended shapes stated; any answer requiring an amendment escalates to master instead

Context: the s4-wire paired audits reconciled with full agreement (ledger entry 2026-07-05,
`docs/sprints/2026-07-05-s4-slice-4/RECONCILE.md`); DESIGN is dispatched in parallel with
these sections marked provisional-pending-guide. Sources below are the audits
(`.relays/s4/s4-wire-audit/AUDIT-implementer-20260705-012253.md`,
`AUDIT-planner-20260705-013000.md`) + my own verification reads. Per the s1–s3 pattern,
answer from locked text where possible; nothing here asks you to amend anything.

**Q1 — per-recipient nudge grain (locked §8.3).** Today: `Server.Push` broadcasts
`{"kind":"delivery-nudge"}` to ALL connected clients; the startup `recovery-nudge` frame
carries the ALL-seats pending list and the global `pending` slice replays to every future
auth and is never cleared (server.go:116-154, :260-270; main.go:187, :199-212 — live-confirmed
by the planner's socket probes, incl. a non-recipient seat receiving another seat's
pending-state metadata). Locked §8.3 says delivery = one write onto THE RECIPIENT seat's
pipe. Proposed design direction: per-connection seat-scoped push after auth; retire the
global pending queue (per-seat recovery nudge issued on that seat's reconnect/auth from
`PendingDeliverySeats`-grain state); mailbox stays truth. CONFIRM this is the locked shape
realized over the existing socket, engine-internal, no contract touch — and name anything
in the current broadcast behavior you consider load-bearing to preserve.

**Q2 — §7 mutation-class semantics (the record itself).** Locked text: §7 :109-111 (a
legitimate config change is itself a committed store record, operator-authorized, carrying
the new digest; config history auditable append-only; no hot reload; digest mismatch ⇒
phase-0 fail-closed) + F11 :172 (the locked mutation-class list already names
`config-change`). Current code: `ValidateGenesis` is a static genesis-digest compare with no
chain walk (genesis.go:104-118); `f11Classes()` has 11 classes, no config-change; a store
whose config member changes after genesis is today permanently diagnostics-bricked absent
re-genesis — exactly the ruled gap. Proposed semantics for confirm: (i) the record commits
through the normal one-pivot discipline UNDER the old running config (no hot reload — the
new digest takes effect at next startup); (ii) the config-member byte replacement under
`<root>/config/` is staged/committed in the same mutation's intent set so a crash at any
syscall boundary leaves either old-config+no-record or new-config+record (F11's
one-pivot-per-mutation property holds for the class); (iii) phase-0 learns the CHAIN:
genesis digest → each committed config-change record's digest in commit order → the loaded
pinned digest must equal the LATEST link (mismatch anywhere ⇒ fail-closed diagnostics, as
today); (iv) the applicability map + crash legs gain the class. Sequencing question inside
(ii): the byte replacement is a plain-file write, not a store-record rename — state the
ordering you want (record-pivot-then-bytes with recovery replaying the byte write from the
record's embedded payload, or bytes-staged-then-record-pivot) — the audits lean
record-carries-the-bytes (the record embeds the new member content, recovery
re-materializes files from committed records — store-is-truth, dumb-replay idempotent).

**Q3 — §7 authorship shape.** m-7 §7 says "operator-authorized committed store record" —
two shapes satisfy it: (a) an operator-channel `submit` (operator-stamped, rides the
normal pipeline + gates; the authorization record IS a first-class relay), or (b) an
admin-time conductor-internal commit (`system`-stamped, genesis-style, m-1 §0.e/§6
provenance). RECOMMENDED: (a) — it exercises the real mechanism, matches the S2 Q2=(i)
precedent (operator authors owed records via the operator channel; recovery reads only the
store) and the s4-dispatch's "the authorization record IS the mechanism working". The
`record_kind` token and the chosen provenance ride the m-1 fidelity packet either way
(ruling condition 4). Confirm (a), or name why (b).

**Q4 — frame-transport ceiling remedy.** Both socket loops use default `bufio.Scanner`
(64 KiB token cap); an oversized frame kills the connection SILENTLY on either side
(server.go:188-201, :448-471 — audit F-W1). Today's describe-grade forms measure ~2.2-2.4 KiB
(both audits), but `read` record bodies, `project` lists, and parent-candidate sets grow
with the store. Remedy classes: raise the buffer (both sides, bounded), chunked framing,
and/or a typed path-free `frame-too-large` refusal instead of silent death. This is engine
transport (yours) — name the sanctioned shape for s4 (our lean: raise to a config-pinned
bound + typed refusal at the bound; no chunking this slice).

**Q5 — heartbeat/wedged-host deferral confirm.** VP-W1's two honest legs (reject-active +
proven-dead recovery via kernel close-detection) leave the wedged-but-alive host: not
provably dead, so its seat stays blocked in the reject regime until the operator kills the
host session. A protocol-level ping would change the frame dialect both sides speak.
RECOMMENDED: defer wedged-host detection entirely for s4 — reject regime + the documented
operator remedy in the ops surface; dialect untouched. Confirm the deferral (or direct the
ping into the guide-reviewed design now).

**Q6 — MCP-composite confirm.** §8.1 "the conductor is an MCP server" — the audits map the
S1–S3 socket dialect as JSON-RPC-shaped but NOT MCP (no `jsonrpc` member, int64 ids with
0-sentinel, bare-string errors, bare-array tools/list, custom `tools/descriptions` +
`notifications/nudge`), and both audits conclude the sanctioned realization is the
COMPOSITE: conductor (bespoke dialect, private) + per-seat shim (real MCP: initialize
handshake, JSON-Schema `inputSchema`, `tools/call` result shape, stdio framing) — matching
your S3 consult Q3 answer ("MCP inherits at wire-up"). Confirm the composite reading — no
socket-dialect rewrite, no locked-text tension. Two sub-confirms: (i) re-render/drift
surfaced as MCP `notifications/tools/list_changed` + re-list (digest already in the
describe response; the bounce carries the `re-render` class); (ii) the documented
poll-hint fallback (project-on-turn) is FIRST-CLASS, not an afterthought — hosts may not
surface custom notifications to the model, and client-side push-drop is real
(server.go:464-468).

Escalation rule (standing): any answer that requires amending locked m-7/m-1 text goes to
master via the amendment path instead of landing here.

ACTIONS_GIT_REF: wrote this relay + an INDEX row (relay substrate, git-untracked); no frank/ or spec-tree edits.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 6987367 + the s4 ledger commits; code surface = tag s3-close exactly)
Next requested action: operator hand-relays to m-7.planner; answers land in ../.relays/s4/s4-guide-q1/ per the s1–s3 pattern; the s4 DESIGN's provisional sections de-provision on your answers.
