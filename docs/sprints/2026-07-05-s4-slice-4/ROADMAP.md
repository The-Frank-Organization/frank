# Sprint s4 — Slice-4: the WIRE-UP (live sessions on frank — the end of the operator-as-transport)

**RUN_ID:** `s4` · **Repo:** `frank/` (branch `main`, baseline tag `s3-close` = main@b5a2c95; code surface verified identical at HEAD 50290e1 — ledger-docs-only delta, my `git diff --stat s3-close..HEAD`) · **Ceremony:** medium · **Opened:** 2026-07-05 · **THE FIRST E3 SLICE** (E3 scoped to transport/provenance)

## Mandate (from master `s4-dispatch`)

**S4 ends the operator-as-transport.** A REAL agent session files a relay through `submit()` and a second REAL session receives it via `project()`/`read` — no human copy-paste anywhere in the loop. Build the **BRIDGE** against the **LOCKED contracts** — build against them, do **not** redefine them:

1. **The per-seat MCP shim** (`frank-mcp` or as named): a stdio MCP server a host session (Claude Code / Codex) launches; config = `{socket path, credential}`; performs `session/connect` on frank's socket; translates MCP `initialize`/`tools/list`/`tools/call` ↔ frank frames; surfaces nudges (MCP notification or documented poll-hint). **One shim per session, one credential per shim** — per-seat channel isolation (DI-2) preserved across the bridge.
2. **The submit tool's input schema IS the rendered form** through MCP — describe-grade per seat×phase×tier (S3 landed the server side; the shim presents it faithfully); re-render-on-drift bounce surfaced usably.
3. **Seat lifecycle hardening for live sessions:** reconnect (nudge flush + `project` catch-up — mailbox = truth); session restart; second-connect semantics — **pre-constrained [VP-W1]: one active channel per credential; reject active duplicates or recover proven-dead channels only** (live supersede/rotation/re-mint = locked-contract touch → escalate); credential-custody posture documented honestly (env-var over on-disk where the host allows; D5 note stated where not).
4. **The §7 config-change record** (discharges **OI-S3-CONFIG-CHANGE**): the commit-loop mutation class + recovery interaction; operator-authorized digest-change record per locked §7 (:109); registry/config evolution on an EXISTING store without re-genesis; the S2 crash-harness applicability map gains the class; m-1 fidelity on the `record_kind`. Dispositioned **through the live owed mechanism** on the real store. Inherits the s3-scope-q1 ruling conditions (m-7 guides · m-1 fidelity · crash-matrix class).
5. **Operational surface:** start/stop/status conventions; team-store + socket-path conventions (absolute paths); the minting workflow documented end-to-end.
6. **Usage-data posture (minimal):** the store IS the usage record; document how to read it (or a trivial read-only stat over `project`); **no analytics** — s5 consumes.

**OUT (escalate before any delegated dispatch that touches these):** consumer schema content (s5) · observe/evidence (Step-2) · routing execution (Step-3) · TUI/email UX (Step-4) · federation (horizon, zero pre-work) · external send/away-bridge · steer/interrupt beyond host-native · any replacement of the operator's *authority* (transport only).

Authorizing relay: `../.relays/s4/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md`
(read-only reference — the master governance trail lives in cwd-parent `master/`, not here).

## Exit gate (HARD acceptance — E3 arrives here, scoped to transport)

- **The live relay (centerpiece):** real session on host A connects via the shim, files a relay; real session on host B receives it (`project`/`read`) — FROM system-stamped, registry-validated, lineage-checked, crash-atomic commit, nudge-delivered. No hand-relay. Ideally A = Claude Code, B = Codex.
- **Adversarial (live):** no-credential connect rejected; bad credential rejected; second active connect on a live credential rejected (or proven-dead recovery only) [VP-W1]; tool surface offers no FROM anywhere; forged/out-of-scope submit bounces; **I-PH across the shim boundary [VP-W3]** — every shim/MCP surface class path-clean: tools/list descriptions · tool input schemas · tool-call results · notifications/poll hints · reconnect errors · credential-failure errors · shim diagnostics.
- **Crash/liveness (live):** kill frank mid-delivery with live clients → restart → wake re-issued, exactly-once effect, mailbox truth; kill the shim → host reconnects → project catch-up complete; queued nudge for an offline seat delivers on reconnect.
- **§7 round-trip on a real store [VP-W4]:** operator-authorized registry change as a store mutation on an EXISTING store (never re-genesis); phase-0 accepts the new digest via the genesis chain; superseded rendered form bounces "re-render", re-rendered form succeeds; crash legs green; **OI-S3-CONFIG-CHANGE closed through the owed mechanism, open set empty**.
- **E2 floors:** full battery green (S1+S2+S3 suites); zero regression; enum byte-exact; guardrail surface still exactly `submit`/`project`/`read`.
- **Honesty [VP-W2]:** every s4 E3 claim surface says **"transport/provenance only"** — `record_integrity` and done-state remain `self_reported` until Step-2 observe; credential-custody D5 posture stated wherever credentials are documented.

## Spec (read-only references — ABSOLUTE paths; never edit; escalate spec problems via s4.orchestrator-planner to master)

- Charter dispatch (scope/gate/exit of record): `.relays/s4/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md`
- Engine spec (PRIMARY — the shim IS the attach surface): `master-docs/master/ARCHITECTURE.md` §C4.1 (engine + interface guardrail) + §C4.3 (claim boundary — tool-mediated confusion-resistance, I-PH, D5)
- **m-7 conductor-core design-of-record (GUIDE — attach/pipe lifecycle §8, trusted config §7 :109)** — `master-docs/master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md`
- m-1 store/identity contract (consulted + fidelity — channel identity §4/§5, credential lifecycle §13.3, conductor-internal provenance §6): `master-docs/master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md`
- The s3-scope-q1 ruling (the §7 conditions this slice inherits): `.relays/s3/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md`
- m-2 (light consult — describe-grade form + re-render bounce crossing the shim intact): `master-docs/master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md`
- S1+S2+S3 continuity (the code this slice bridges): `docs/sprints/2026-07-03-s1-slice-1/` + `docs/sprints/2026-07-03-s2-slice-2/` + `docs/sprints/2026-07-04-s3-slice-3/` + the source at tag `s3-close`
- Sequencing: `master-docs/ROADMAP.md` (Step-1)

**Guide:** m-7 (`m-7.planner`, via operator hand-relay) — the engine continuity; the shim is the attach/pipe-lifecycle surface (m-7-owned); the §7 record is an engine mutation class (m-7-guided per the s3-scope-q1 ruling, condition 4).

## Riding in

- **`OI-S3-CONFIG-CHANGE`** — the only open owed item (ledger record: `../2026-07-04-s3-slice-3/results/OI-S3-CONFIG-CHANGE.md`); discharged at this slice's exit gate through the live owed mechanism on the real store.

## Plan-gate (F2 — non-bootstrap; conditioned delegation)

Pair Implementer plan-review = the plan gate; `DISPATCH IMPL` delegated only under {Implementer approve · no scope/boundary deviation · no hard trigger · no cross-slice collision · no locked-contract or design-of-record amendment}; any failure — including any OUT-item touch — escalates to master (CTO + m-7 guide + VP). Second-connect semantics needing new *contract* surface = a locked-contract touch = escalate.

## Operator-judgment items

- Live-test seats are the operator's to designate (which two real sessions play A and B — they hold minted credentials).
- §7 changes are operator-authorized by design — the live round-trip needs the operator's explicit authorization record (that IS the mechanism working).
- Residual risk (accepted, restated): D5 — credential theft by a shell-bearing co-resident process is out of scope (confusion-resistant, not theft-proof); custody posture documented, not over-claimed.
