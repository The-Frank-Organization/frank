## SITREP — master → s4: GATE-DAY RESULTS + three findings for the pair to fold before the FORMAL s4 exit-gate record — a live cross-vendor relay flew (genuine E3), every leg verified at the master seat; F-GATE-1 closed, F-GATE-2 (owed fill-time gap) + F-GATE-3 (shim-stderr I-PH borderline) are yours to disposition

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-gateday
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no — a results-and-findings handoff into the slice's gate-day lane; the pair folds the findings + files the FORMAL exit-gate record; the operator's merge + s4-close remain their own gates
IN_REPLY_TO: frank/.relays/s4/s4-exit-gate/SITREP-orchestrator-planner-20260705-145109.md
FROM: master.orchestrator-planner
TO: s4.orchestrator-planner
CC: operator, s4-wire.planner, s4-wire.implementer, s4.orchestrator-reviewer, m-7.planner, m-2.planner, m-1.implementer
SUBJECT: operator-run gate-day (master-supported) PASSED substantively — live Claude↔Codex relay + every adversarial/crash/§7/owed leg verified at the store; F-GATE-1 fixed+on-branch; fold F-GATE-2/F-GATE-3 + file the formal gate record; this is NOT itself that record

**What this is.** The master seat's independent report of the operator-run gate day (I supported diagnosis; the operator drove; you were on standby). **It PASSED substantively.** This relay is **not** the formal s4 exit-gate record — that is yours to file (procedure of record, clean, with these findings folded), after which the s4 orchestrator verifies and SITREPs master and the operator's merge/close gates run. I report results + three findings.

### What was proven (each independently verified by me against the persistent store `$HOME/frank-team-store`)
- **The live-host E3 centerpiece — REAL:** a governed relay filed by a live **Claude Code** session (`s4-wire.host-a`, implementer) via `submit`, received by a live **Codex** session (`s4-wire.host-b`, planner) via `project`/`read`, **no human transport** — `relay-4a33925bca720b0cd0f1e180`, conductor-stamped `from: s4-wire.host-a` (not agent-supplied), checksum identical store↔read. Two vendors: provider-agnosticism **exercised**, not assumed. This satisfies the live-host designation.
- **Adversarial (live/CLI):** no-cred + bad-cred → typed `shim:auth-failed` (path-free MCP result); **second-connect refused on a live channel** → `auth:channel-active` (the live Claude session held host-a); **forge** (submit as A claiming `from=operator`) committed `from: s4-wire.host-a` — channel beats claim.
- **Crash/liveness:** SIGKILL the conductor mid-run → tool call fails typed (`shim:conductor-unreachable`) → restart over the store → **byte-identical recovery** (checksum `1c41e0d1…` held) → **exactly-once** (record count stable, no re-emission).
- **§7 config-change round-trip on the EXISTING store (not re-genesis):** operator-authored `config_change` accepted (`relay-7ac1c64b807781db73f8fb9f`) → restart, **phase-0 accepted the genesis→config_change chain** → the form re-rendered (digest **`11ecf52b…` → `7b14f9c1…`**) → stale-digest submit **bounced** → fresh-digest submit **accepted**.
- **Owed lifecycle:** `owed_item` materialized (`relay-b38753d455bbb4ba448a1a65`, all four headers) → `owed_disposition` (`relay-7187783162df2d5410efaffa`, `disposes_owed` + cites the config_change) → **`OPEN.md` empty**. `OI-S3-CONFIG-CHANGE` — Step-1's last owed item — is discharged through the live projection.
- **Ledger integrity:** every rejection is a *correct* rejection (stale digest; missing `owner`), every acceptance legitimate, intake sequence monotonic 1–7.

### Findings to fold before the formal record

**F-GATE-1 — CLOSED (report only).** The shim's `initialize` omitted MCP-required `serverInfo.version`; live Claude Code rejected the handshake in 13ms. Fixed on-branch (`cmd/frank-mcp/mcp.go`, const `"0.4.0"`, + a red-first envelope fixture) at the gate commit `605b3ef`; verified by the pair-planner and re-confirmed reconnecting live. Confirm it rides the gate record.

**F-GATE-2 — OPEN (real; yours to disposition). Fill-time-authority gap on `owed_item`.** `internal/engine/submit.go:128` requires `owner`, `source`, `target_surface`, `disposition_path` for `record_kind=owed_item` — but **none of the four is a declared FieldSpec field** (zero occurrences in `internal/fieldspec/registry.json`; `disposes_owed` likewise for `owed_disposition`). So the operator's **rendered form never prompts them**, and a real operator agent filing an owed item hits `owner:required` with **no in-form guidance** — the mirror of S3's fill-time principle ("forbidden absent from the form"): here, *required, but not rendered*. Fixtures pass only because they hardcode the headers. **Disposition options:** (a) declare the four (+`disposes_owed`) in the registry with a `required_when: record_kind==owed_item` predicate so fill-time authority renders + enforces them (an m-2/registry fold — this looks like the right fix, and small); or (b) accept as a known gap and record it as its own owed item to the owning slice. My read: (a), and it is arguably in-fence for s4 if light, else an m-2 micro-fold — the pair + m-2 decide.

**F-GATE-3 — OPEN (borderline; pair's ruling). Shim stderr names the socket path.** During a conductor-down dial, `frank-mcp`'s own **process stderr** prints `dial unix /tmp/frank-s4.sock: connection refused`. The **MCP-delivered** result is clean and path-free (`shim:conductor-unreachable`) — so I-PH (seat-*delivered* surfaces) is **not** violated at the agent boundary. But the [VP-W3] enumerated surfaces included "host-visible shim diagnostics." **Rule it:** either (a) I-PH scope is the delivered MCP surface only → this is out of scope, note-and-close; or (b) the shim's stderr is in scope → redact the socket path from shim-side error text. Low-severity either way; make the call explicit so it doesn't sit as an implicit gap.

### Honest scope (for your record, so it isn't over-read)
The centerpiece (live relay + second-connect) is genuine **E3 live-host**. The other legs were operator/CLI-driven against the live store — real crashes, real §7, real owed cycle, but **E2-mechanical execution**, driven by master-supplied scaffolding whose **three own bugs** (a grep-escape misread of an accepted result; a missing-owed-header payload; a wrong `OPEN.md` match) were caught and fixed live and are **not** frank defects. The formal gate record should rest on **your** clean procedure-of-record run (or an explicit acceptance of this evidence set with the scaffolding caveat stated), not on this master report.

### Requested of the pair
1. Confirm F-GATE-1 rides the branch/gate record.
2. **Fold F-GATE-2** (recommend the registry declaration of the owed headers) or record its disposition explicitly.
3. **Rule F-GATE-3** (in/out of I-PH scope) and act or note accordingly.
4. Run the procedure of record clean (or accept this evidence set with the caveat), file the **formal s4 exit-gate SITREP**; the s4 orchestrator verifies + SITREPs master.
5. Then the operator's merge + `s4-close`; master folds S4 into `RECONCILE.md`; s5 (consumer schemas) builds over the wired conductor.

## Verification
- Independent store reconciliation this seat: the 8-record ledger (intake 1–7; two correct rejections); the live relay stamp + checksum; the re-render digest change; `OPEN.md` empty; I-PH sweep (the only path hit is the shim-stderr of F-GATE-3, not a delivered surface).
- Code anchors: `internal/engine/submit.go:128` (owed required set), `:153-171` (config_change operator-scope), `internal/fieldspec/registry.json` (owed headers absent); gate binaries built at `frank` `605b3ef`.
- Evidence dir (operator's machine): `frank/docs/sprints/2026-07-05-s4-slice-4/results/e3-20260706T043908Z/transcripts/` (10-live-relay-note + 11–13 relay + 20–24 adversarial + 30–32 crash + 40–43 config-change + 50–51 owed).

ACTIONS_GIT_REF: none — master-seat report + this relay + an `INDEX.md` row; no `frank/` code edit (gate-day artifacts are the operator's, on their machine); cwd is not a git repo (docs workspace).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` at gate commit `605b3ef` (branch `s4-wire-impl`), the fix on branch.
Next requested action: operator hand-relays this to the s4 session; the pair folds F-GATE-2/-3 + confirms F-GATE-1, files the formal exit-gate record; s4 orchestrator verifies + SITREPs master; then the operator's merge/close and master's S4 fold.
