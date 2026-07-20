## PLAN — Slice-4 build dispatch (master → s4 slice-team; THE WIRE-UP: live sessions on frank — the end of the operator-as-transport)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s4-dispatch
PARENT_DISPATCH_ID: step1-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — operator elected the wire-up as s4 (2026-07-05, renumber recorded below); s4 runs on F2; the s4-close sign-off is the operator's, exercised separately
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below (F2, non-bootstrap)
GRILL_REQUIRED: no
IN_REPLY_TO: frank/.relays/s3/s3-exit-gate/SITREP-orchestrator-planner-20260704-224230.md
FROM: master.orchestrator-planner
TO: s4.orchestrator-planner, m-7.planner
CC: master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-7.implementer
SUBJECT: s4 = the MCP wire-up — the per-seat shim (live Claude Code/Codex sessions on frank's socket) + seat lifecycle hardening + the §7 config-change record (discharges OI-S3-CONFIG-CHANGE) + the FIRST LIVE E3; NEW slice-team; guide m-7; F2 conditioned delegation

**What this is.** The master dispatch for **Slice-4 = the wire-up** — the operator's elected fork (chartered "next after S3, operator's call" in the s3-dispatch; elected 2026-07-05). **VP pre-handoff review: APPROVE** (`s4-dispatch/RECONCILE-orchestrator-reviewer-20260705-001405`); its four watchpoints are folded inline below (marked **[VP-W1..W4]**). **Renumber of record:** the decomposition's old "Section 4" (consumer schema slices + §C4 fixtures) becomes **s5**; slice RUN_IDs count team instances. **Sequencing intent (operator):** s5 is then built *with s4 in use* — the s5 team runs over the wired conductor where practical, its registry additions landing as real §7 config-change records, generating the first usage data. Baseline: S1+S2+S3 CLOSED (`main`, tag `s3-close`; battery 20 green; the machine is whole on fresh stores and speaks the real protocol). **Riding in: `OI-S3-CONFIG-CHANGE`** — the only open owed item; it is discharged in this slice.

**The goal (one line):** **end the operator-as-transport** — a real agent session files a real relay through `submit()` and a second real session receives it via `project()`/`read`, with no human copy-paste anywhere in the loop. Step-1's stated goal, delivered live. **This slice produces frank's first E3 evidence.**

### To the s4 slice-team — your charter
- **NEW slice-team** (new sprint = new team). **Use `/orchestrator-planner`**; scaffold an **`s4`** sprint via `sprint-doc-setup` in `frank/`; relays live in `frank/`.
- **Onboard first — you built none of S1–S3.** Read the source + the three sprint ledgers; re-run the battery at `s3-close` yourself (the standing bar: fresh eyes have found real fragility in every slice so far). **m-7 (your guide) is the engine continuity** — attach/pipe lifecycle + the interface guardrail are m-7's owned domain, and the shim IS the attach surface.
- **Spec = read-only in cwd:** `ARCHITECTURE.md` §C4 (esp. C4.1 engine + C4.3 claim boundary/I-PH) + the m-7 domain doc (attach/pipe lifecycle, trusted config §7) + the m-1 doc (channel identity, credential/binding, §6 provenance) + the s3-scope-q1 ruling (`master/relays/s3-scope-q1/RECONCILE-…-171608` — the §7 conditions you inherit). Escalate spec problems; do not self-amend.
- **Build on `main`** (post-`s3-close`), on a branch; close-time integration is the operator's separate gate.

### Guide + contract boundaries
- **m-7 primary guide** — the shim is the attach/pipe-lifecycle surface (m-7-owned); the §7 config-change record is an engine mutation class (m-7-guided per the s3-scope-q1 ruling, condition 4).
- **m-1 consulted + fidelity** — channel identity, credential lifecycle/custody posture, the §7 `record_kind` (m-1 fidelity per the ruling), and second-connect semantics — **pre-constrained [VP-W1]:** *exactly one active channel per credential is the safe default.* s4 may **reject** a second active connect, or **recover a proven-dead/stale channel** as reconnect behavior — those are within the locked contracts. Any **live supersede, rotation, or re-mint-supersedes** behavior is a **locked-contract touch** → escalate through m-1/m-7 amendment review before delegated implementation authority. The slice does not choose beyond that fence.
- **m-2 light consult** — the describe-grade form + re-render bounce crossing the shim boundary intact.

### Slice-4 scope (IN)
1. **The per-seat MCP shim** (`frank-mcp` or as named): a stdio MCP server a host session (Claude Code / Codex) launches; config = `{socket path, credential}`; performs `session/connect` on frank's socket; translates MCP `initialize`/`tools/list`/`tools/call` ↔ frank frames; surfaces nudges (MCP notification or documented poll-hint). **One shim per session, one credential per shim** — per-seat channel isolation (DI-2) preserved across the bridge.
2. **The submit tool's input schema IS the rendered form** through MCP — describe-grade per seat×phase×tier (S3 landed the server side; the shim presents it faithfully), re-render-on-drift bounce surfaced usably.
3. **Seat lifecycle hardening for live sessions:** reconnect (nudge flush + `project` catch-up — mailbox = truth), session restart, second-connect semantics (above), credential-custody posture documented honestly (env-var over on-disk where the host allows; where not, the D5 note stated — theft ≠ confusion, out of scope, but *say it*).
4. **The §7 config-change record** (discharges **OI-S3-CONFIG-CHANGE**): the commit-loop mutation class + recovery interaction; **operator-authorized** digest-change record per locked §7 (:109); registry/config evolution on an EXISTING store without re-genesis; the S2 crash-harness applicability map gains the class; m-1 fidelity on the `record_kind`. The owed item is dispositioned **through the live owed mechanism** on the real store.
5. **Operational surface:** start/stop/status conventions, the team-store + socket-path conventions (absolute paths; store = the governance domain), the minting workflow documented end-to-end (the §2-of-record answer to "how do we actually run this").
6. **Usage-data posture (operator intent, minimal):** the store IS the usage record (append-only relay graph, bounce reasons, render→submit outcomes). Document how to read it (or a trivial read-only stat over `project`); **do not build analytics** — s5 consumes the data.

### Slice-4 scope (OUT — escalate before any delegated dispatch that touches these)
Consumer schema *content* (**s5**) · observe-as-send / evidence (**Step 2**) · routing execution (**Step 3**) · TUI / email-client UX (**Step 4**) · federation (**horizon — no pre-work sanctioned**) · external send / away-bridge (the outbox stays local; egress stays dormant) · steer/interrupt beyond what the host harness natively gives · any replacement of the operator's *authority* (gates still park for the human — **only the transport is replaced**).

### Slice-4 exit gate (HARD acceptance — E3 arrives here, scoped to transport)
- **The live relay (the centerpiece):** a REAL session on host harness A connects via the shim and files a relay; a REAL session on host harness B receives it (`project`/`read`) — `FROM` system-stamped, validated against the real registry, lineage-checked, committed crash-atomically, delivered with a nudge. **No hand-relay anywhere.** Ideally A = Claude Code and B = Codex (provider-agnosticism exercised, not assumed).
- **Adversarial (live):** connect without credential → rejected; bad credential → rejected; **second active connect on a live credential → rejected (or proven-dead recovery only) [VP-W1]**; the tool surface offers **no** FROM anywhere (probe it); forged/out-of-scope submit bounces; **I-PH across the shim boundary [VP-W3] — the shim/MCP surface classes enumerated in the test matrix, each one path-clean:** `tools/list` descriptions · tool input schemas · tool-call results · notifications/poll hints · reconnect errors · credential-failure errors · any MCP-returned shim diagnostics. Not just conductor-store projections — the *bridge's own* surfaces.
- **Crash/liveness (live):** kill frank mid-delivery with live clients attached → restart → wake re-issued, exactly-once effect, mailbox truth holds; kill the shim → host session reconnects → `project` catch-up complete; a queued nudge for an offline seat delivers on reconnect.
- **§7 round-trip on a real store [VP-W4]:** an operator-authorized registry change **evidenced as a store mutation on an EXISTING store — never simulated by re-genesis**; phase-0 accepts the new digest via the genesis chain; a superseded rendered form bounces "re-render" and the re-rendered form succeeds; crash legs for the new mutation class green; **OI-S3-CONFIG-CHANGE closed through the owed-item mechanism, open set empty**.
- **The E2 floors:** full battery green (S1+S2+S3 suites); zero regression; enum byte-exact; the guardrail surface still exactly `submit`/`project`/`read`.
- **Honesty [VP-W2]:** **every** s4 E3 claim surface says **"transport/provenance only"** — `record_integrity` and done-state remain **`self_reported`** until Step-2 observe lands; the credential-custody D5 posture stated wherever credentials are documented. The phrase stays visible at every claim surface, not just the gate record.

### Plan-gate (F2 — non-bootstrap; conditioned delegation)
Pair Implementer plan-review = the plan gate; `DISPATCH IMPL` delegated only under **{Implementer approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract or design-of-record amendment}**; any failure — including any OUT-item touch — **escalates to master (CTO + m-7 guide + VP)**. Second-connect semantics needing new *contract* surface = a locked-contract touch = escalate (amendment path, not improvisation).

### Deliverable format
The shim + hardened lifecycle + the §7 record on a branch; the exit-gate evidence (incl. the live-relay E3 record — transcript/store evidence of the real two-session exchange); your build relays in `frank/`; a SITREP back to master at the exit gate; close-time integration = the operator's gate.

### Operator-judgment items
- **The live-test seats are the operator's to designate** (which two real sessions play A and B — they hold minted credentials).
- **§7 changes are operator-authorized by design** — the live round-trip will need your explicit authorization record (that IS the mechanism working).
- **residual risk (accepted, restated):** D5 — credential theft by a shell-bearing co-resident process is out of scope (confusion-resistant, not theft-proof); custody posture documented, not over-claimed.

### Not authorized by this relay
No s4-close authority, no scope expansion, no locked-design amendment, no federation pre-work. `DISPATCH IMPL` only under the F2 conditions; failures escalate.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s4-dispatch` — run below (clean both modes).
- Executes the s3-dispatch's chartered fork (operator elected wire-up 2026-07-05) + the s3-scope-q1 ruling conditions (§7 → this slice). Renumber (old Section-4 → s5) recorded here + kickoff + dashboard; VP pre-handoff review covers it.
- Pointers: `ARCHITECTURE.md` §C4.1/§C4.3, m-7 §7 (:109), m-1 §5/§6, `master/relays/s3-scope-q1/RECONCILE-…-171608`, `frank/` (baseline `main`, tag `s3-close`).

ACTIONS_GIT_REF: wrote this s4-dispatch + the s4 boot + a VP review request + `INDEX.md` rows + kickoff/dashboard renumber notes; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no `frank/` write.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main`, clean, tag `s3-close`.
Next requested action: VP pre-handoff review (request relay alongside); on approve the operator relays the s4 boot then this dispatch to a fresh session; s4 onboards, scaffolds, plans, and dispatches under F2; SITREP back at the exit gate.
