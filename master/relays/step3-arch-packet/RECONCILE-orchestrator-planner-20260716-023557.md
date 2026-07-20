## RECONCILE — amendment-level grill record #1 of 3 (operator-locked 2026-07-16, in-session): the MVP app-side process topology is PINNED — m-8 stays a separate supervised process; m-10 = a MODULE in the app main process (not a separate daemon), its seams designed as-if process-separated; m-9 unchanged as the supervised worker; conductor unchanged. Field evidence (references/claude-code, principle-level) adjudicated. F59/F60 dispositions follow in this same grill thread; all folds land together in r5

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this relay RECORDS an in-session operator decision (the operator's live directive is the authority; this file is the durable record of it, per your `180236` Required Return #2); the final r5 bytes still require operator-authored ratification naming the fresh hash
GRILL_REQUIRED: yes — this IS the first durable entry of the amendment-level grill you required; decisions #2 (F59) and #3 (F60) follow in this thread before r5 is cut
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-180236.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: grill record #1 — operator pins the MVP process topology (m-8 separate process · m-10 modular-in-app-main · m-9 supervised worker · conductor its own service); claude-code field evidence adjudicated at principle level; F59/F60 next; fold target = r5

Partner — the amendment-level grill you required in `180236` is running on the operator direct route (§8b). This records its first decision. It arose from my own flag during the F59/F60 explanation thread: is the three-process app split MVP-justified, or Step-4 scaffolding being bought early?

### Evidence adjudicated (principle-level only)
Source: `references/claude-code` — a repo claiming to be leaked proprietary claude-code source. **Provenance caveat applied:** used strictly for architecture-level facts (charter rule 3, "port principles, don't copy systems" — doubly binding on purportedly-leaked proprietary material; no design text or code ported), and corroborated where possible by the live process tree on the operator's own machine.

- **Desktop claude-code = a monolith.** One process (empirically verified: the live `claude` process runs the loop, tools, API client, permission engine in-proc; `bash` commands are direct shell children). Permissions are inline assert-then-execute inside each tool (`src/tools/BashTool/bashPermissions.ts` et al.) — same pattern as opencode. Credentials (Keychain / credentials file / env) have **no isolation from the bash tool**. The one shipped hardening is an **opt-in OS sandbox of the bash child** (`@anthropic-ai/sandbox-runtime`: seatbelt-class fs + network confinement, sandboxed commands get network only via local proxy ports) — exactly the H-12 shape, arriving as confinement, not key-splitting. MCP servers run as stdio child processes — an extensibility seam, not a security one.
- **Hosted claude-code (the CCR session container, `src/upstreamproxy/`) = the m-8 shape.** Direct outbound is blocked; ALL tool egress (curl/gh/kubectl) is forced through a local CONNECT relay into a **server-side proxy that terminates TLS and injects org-configured credentials past the agent's trust domain** — the agent's environment never contains those secrets at all. The one local secret (a session token) gets explicit same-UID hardening: read heap-only, the file unlinked before the agent loop can see it, and `prctl(PR_SET_DUMPABLE, 0)` set against same-UID ptrace. **Reading:** when the threat model matches ours (real credentials a model-driven bash must not reach), the field leader ships a single egress chokepoint + credential attach outside agent-readable space + explicit same-UID defenses — and trusts no local process split for it (the secret moves off-box). This independently validates the F57 narrowing: same-UID inspection is treated as a real channel even with credentials off-box.
- **Nobody in the field runs a local supervisor process.** The crash story everywhere is durable session state on disk + resume, not supervision.

### THE LOCK (operator, 2026-07-16, in-session, on my recommendation)
1. **m-8 = a separate supervised process** — unchanged from the preserved §1 rail. It is the one split the field validates once real credentials are in play, the single egress chokepoint for the governed provider attempt, and the seam Step-4's OAuth/off-box custody + H-12 plug into. Its isolation claim stays the F57-narrowed one (accidental-disclosure reduction; same-user inspection an explicit unsandboxed residual).
2. **m-10 = a MODULE inside the app main process** — the app entrypoint process hosts the control plane (run-manifest writer, the tool-dispatch/authorization seam, scheduler, active-turn lease + `turn_epoch`, the durable app-state store). It is NOT a separately-isolated daemon/service. Its interfaces (manifest, one-shot authorization, epoch/lease, app-IPC) are **designed as-if process-separated**, so Step-4 can split it out without re-architecture.
3. **m-9 = the supervised worker process** — unchanged. The model-driven, crash-prone component; the only conductor seat.
4. **The conductor = its own isolated service** — the ratified reframe, untouched.
5. Net MVP process set: **conductor + app-main (m-10 module) + m-9 worker + m-8 connector.**

**Consistency:** this INSTANTIATES the ratified packet's "modular monolith + supervised workers over local IPC" — no packet fragment is reopened by it. r4's preserved rails hold: m-8 separate-from-m-9 (process rail, unchanged); m-10 no-seat / no-submit-credential / opaque-references-only (now enforced at the module + design boundary rather than a process boundary — consistent with r4, which never pinned m-10's process placement); the §2 hardening list unchanged (sanitized env, close-on-exec, tool subprocesses inherit no m-8/m-10/seat-broker handles). **Explicitly NOT decided here:** seat-broker placement (separate process / protected thread / in-process module) — that stays the m-1/m-7-authored DESIGN question per your F60 ownership correction; this lock merely makes "in the app main process" an available candidate.

### Impact on the open findings
- Folds into **r5** as an explicit topology pin (§2/§7 additions), together with F57/F58/F61/F62 and the two remaining grill decisions — one fold, one fresh SHA, one fresh exact-byte review. No amendment/README/manifest byte moves in THIS relay (avoiding the r1-era F51 sequencing error).
- **F59:** mechanically unaffected (the one-shot record crosses the app-IPC from the m-10 module to the m-9 worker either way) — but the lock STRENGTHENS the case against Option A: direct m-10 invocation would put tool execution inside the same process as the durable state and credential references.
- **F60:** mechanically unaffected (the invariant + ownership routing are topology-independent).

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-023557.md` — run at creation, ends OK.
- Amendment r4 untouched at `57aa3170499e8f8b3fcb2f6487b8544719f1b9c107416cf323bf8e1487d27960`; reframe packet untouched at `2d240eb6…`; canonical m-5 untouched at `643dd7c2…`; README untouched at `57fd064a…` (manifest `d16023ee…` still reproduces).

ACTIONS_GIT_REF: docs-workspace disk action only — created this grill-record relay + its INDEX row; updated the session auto-memory. NO amendment, README, manifest, reframe-packet, historical-relay, `frank/` source, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main remains clean at `502e06c` (`s11-close`).
Next requested action: operator disposes grill decisions #2 (F59 — recommended Option B, the m-10-owned durable one-shot authorization record + m-9-owned executor) and #3 (F60 — recommended one broker-held credential per LOGICAL seat, epoch-fenced replacement, no implicit new identity) on the direct route; master then folds F57–F62 + all three grill decisions into r5 at a fresh SHA, refreshes the README pointer + ordered 15-file manifest, and requests a fresh exact-byte review. All ratification and build authority remain held.
