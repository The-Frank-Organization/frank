## AUDIT — paired independent audits for Slice-4 (the wire-up): onboarding + spec-to-exit-gate map + 4-bucket verdicts + reuse gate + claim-boundary probes

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s4-wire-audit
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
FROM: s4.orchestrator-planner
TO: s4-wire.planner, s4-wire.implementer
CC: s4.orchestrator-reviewer, operator
SUBJECT: AUDIT (paired, independent, read-only) — S4 wire-up ground truth: the socket/channel surface, seat lifecycle, describe/render path, §7 mutation-class surface, shim reuse candidates; file independently, do NOT coordinate; reconciliation is mine

**Paired independent audits.** Each of you files your own AUDIT relay under this DISPATCH_ID.
Do not coordinate, do not read the other's audit before filing yours; reconciliation is my job.
Read-only: no code, no fixture, no doc edits anywhere in `frank/` or the spec tree.

### Scope (what your audit must cover)

1. **Onboarding proof (the standing bar):** you built none of S1–S3. Evidence your own reads of
   the three sprint ledgers + the source, and your OWN uncached battery run
   (`go test -count=1 ./...` + `go vet` at the s4 baseline; state the HEAD you ran at and its
   relation to tag `s3-close`). Fresh eyes have found real fragility in every slice so far
   (S2: 2 latent S1 races; S3: F-P1..F-P6); the wire-up rides the least-exercised packages —
   look hardest at `internal/channel` (the socket server IS the integration surface),
   `internal/seat` (binding/credential), `cmd/frank/main.go` (assembly, channelTools, Describe,
   turnContextForSeat, mint/operator-submit), and the describe/render path into
   `internal/fieldspec`.
2. **Spec-to-exit-gate map:** every s4 exit-gate line (ROADMAP `## Exit gate`) mapped to the
   locked text that grounds it — m-7 §8 (attach/pipe lifecycle, guardrail, wake), m-7 §7 :109
   (the config-change record), m-1 §4 (DI family/I2), §5 (API), §6 (system-field contract +
   conductor-internal provenance), §13.3 (credential-lifecycle carry), ARCHITECTURE §C4.1/§C4.3
   (I-PH, claim boundary), the s3-scope-q1 ruling conditions, and the s4-dispatch IN items 1–6.
   Flag any exit-gate line you cannot ground (spec-gap escalations to me — do not self-amend).
3. **4-bucket verdicts** (`still-open | already-closed | product-overlapped | recommended-next`
   + PRIMARY_BUCKET) for each s4 IN item: (1) the per-seat MCP shim, (2) the rendered-form
   schema through MCP, (3) seat lifecycle hardening (reconnect / restart / second-connect /
   custody posture), (4) the §7 config-change record, (5) the operational surface, (6) the
   usage-data posture.
4. **Duplicate/reuse gate — named candidates (promote, don't rebuild; verify, don't assume):**
   `internal/channel.Client` (Dial/DialAuthenticated/DescribeTools/Call/NextPush — is this the
   shim's conductor-side half already?); `test/seatproc/testseat.go` (the stub seat binary);
   `cmd/frank` `-operator-submit` + `-mint` (the existing socket-client + custody workflow);
   the S3 describe-grade render (server-side landed — what exactly does the shim still owe?);
   the S1/S2 crash harness (crashpoint registry + child-SIGKILL, the §7 crash-legs machinery);
   the live owed-item mechanism (obligation package — the OI-S3-CONFIG-CHANGE disposition path);
   `store.Init` config pinning + recovery phase-0 (the surface the §7 record mutates against).
5. **m-1/m-7 boundary surface enumeration:** every store/binding/config surface s4 work will
   touch (the m-1 fidelity packet input; the m-7 guide consult input). Name file:line. The §7
   record is an engine mutation class — enumerate what the commit loop, recovery phases, redo
   journal, and the S2 applicability map must gain for it.
6. **Claim-boundary probes (verify or refute each independently, file:line/E2 evidence):**
   (a) nudge delivery is broadcast to ALL connected clients (`server.Push`/`broadcast`/
   `QueuePush`) vs the locked per-recipient pipe write (m-7 §8.3) — and the `recovery-nudge`
   frame carries the all-seats pending list to every client; (b) the global `pending` slice is
   flushed to every newly-authed client and never cleared; (c) no second-connect constraint
   exists — any number of `session/connect` on one credential (VP-W1's target); (d) no
   rotation/revocation surface (m-1 §13.3 carry — enumerate what custody honesty requires);
   (e) credential custody today: `-mint` → stdout, `-operator-submit -credential` → CLI flag
   (ps-visible); (f) `bufio.Scanner` default frame ceiling on both socket sides vs
   describe-grade form payloads + large `project` results — measure a real rendered-form
   payload size; (g) the I-PH surface census across the bridge classes named in the exit gate
   (tools/list descriptions · input schemas · tool-call results · notifications/poll hints ·
   reconnect errors · credential-failure errors · shim diagnostics) — which are path-clean
   today (`safeError`, `auth:*` strings, outcome JSON), which have no fixture; (h) MCP fidelity:
   the m-7 design says "the conductor is an MCP server" — the S1-S3 socket protocol is a
   bespoke JSON-RPC-shaped frame dialect; map what a REAL MCP host (Claude Code / Codex)
   requires (initialize handshake, tools/list schema shape, notifications) vs what the shim
   must translate; flag anything that smells like a locked-contract touch rather than shim
   affordance.
7. **Second-connect ground truth for the VP-W1 fence:** what "proven-dead" can honestly mean on
   a unix socket today (connection close detection, kernel buffer states); what the reject leg
   needs; anything beyond reject/proven-dead-recovery = locked-contract touch = name it, do not
   design it.
8. **Honesty-line census:** where the E3-transport/provenance qualifier and the D5
   credential-custody posture must appear (README, shim docs, minting workflow, gate records) —
   the [VP-W2] every-claim-surface rule.

### OUT (name-and-escalate, never audit-into-design)
Consumer schema content (s5) · observe/evidence (Step-2) · routing execution (Step-3) ·
TUI/email UX (Step-4) · federation (zero pre-work) · external send/away-bridge · authority
replacement. An audit finding that any IN item requires an OUT touch or a locked-contract
amendment is an escalation to me, not a design proposal.

### Deliverable
One lint-clean AUDIT relay per seat under `.relays/s4/s4-wire-audit/`
(`AUDIT-<role>-<ts>.md`), INDEX row appended (end-of-file, write order). 4-bucket verdicts
with PRIMARY_BUCKET per IN item; every claim carried with evidence level (E1 file:line / E2
own-run); `FINAL_GIT_STATUS_SHORT` from a fresh check after all audit work (read-only proof).
Boot-ACK first if you have not yet ACKed (separate relay under your boot DISPATCH_ID).

Not authorized by this relay: DESIGN, PLAN, IMPL, REVIEW-FOLD, MERGE, LIVE-VERIFY, or any
edit. The DESIGN dispatch follows my reconciliation of the paired audits.

FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 13c6ce2, the stand-up ledger commit; scaffold 56a19ec; code surface = tag s3-close exactly)
