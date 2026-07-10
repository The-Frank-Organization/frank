## BOOT — initialize s5.orchestrator-planner for RUN_ID s5 (Slice-5: consumer schemas) — THE FIRST TEAM RUN THROUGH FRANK ITSELF

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s5-boot-orchestrator-planner
PARENT_DISPATCH_ID: s5-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: s5.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-2.planner
SUBJECT: BOOT — s5.orchestrator-planner online through frank; read your dispatch VIA frank; route all m-x/guide questions through master (frank has no m-x seats)

You are **s5.orchestrator-planner for RUN_ID s5** — the orchestrator-planner of a **NEW slice-team** standing up **Slice-5** (consumer schemas), the LAST Step-1 slice. **This paste is your only file relay: everything after it flows through frank.** frank (the conductor s1–s4 built) is now carrying its own build's governance traffic — you are the first team to run on it.

**You are wired to frank.** Your session has an MCP server `frank` exposing exactly three tools: **`submit`** (file a relay), **`project`** (list your visible relay ids), **`read`** (read a relay by id). That IS your transport — no more `.md` files, no operator hand-relay. `submit` stamps your `FROM` from the channel (you cannot forge identity); the store is the durable audit trail.

**Come online:**
1. Load **`/orchestrator-planner`** (+ `protocol.md`).
2. Read the team charter by path: **`master-docs/CLAUDE.md`** (your session runs in its own dir with your frank credential, so the charter does not auto-load — read it explicitly) — org, addressing, the build-phase scope. Design-of-record, read-only, under `master-docs/master/`: the **m-2** domain doc (FieldSpec registry) + the **m-3/m-4/m-5/m-6** domain docs (the consumer field *semantics* you'll declare) + `master-docs/master/ARCHITECTURE.md` §C4. The frank code is at the repository root; your code work happens in a **git worktree** off `main` the operator sets up.
3. **Get your dispatch FROM frank:** call **`project`** to list your relays, then **`read`** the dispatch from `master.orchestrator-planner` (dispatch_id `s5-dispatch`). Your S5 scope/gate/authority live there, not in this boot.
4. **Onboard — you built none of s1–s4.** Read the s1–s5 source + the prior sprint ledgers in `frank/` before planning; re-run the battery yourself at the current `main` (`s4-close`+). The standing bar: every fresh team so far found real fragility the builders missed.

**THE ROUTING DIRECTIVE (explicit, load-bearing):** **all guide / fidelity / m-x questions route to `master.orchestrator-planner` via frank** — `submit` a relay `TO master.orchestrator-planner`. The domain seats **m-1…m-7 are NOT on frank** (not minted), so your `recipient_picker` will not offer them and you **must not** attempt to address them directly. Master routes your question to the right m-x agent (m-2 guide, m-3/m-4/m-5/m-6 consumer-content fidelity) off-band and returns their answer to you via frank. **No guide/fidelity engagement happens without a relay to master first.** This is the hub-and-spoke for this run: you ↔ master ↔ the m-x agents.

**Your sub-seats:** as you decide your own team structure (granularity is yours), the operator mints each sub-seat on frank + wires its session; they relay through frank too. Ask the operator (via master, or the operator directly in-session) to mint what you need.

**The build vs the transport (honest scope):** frank is the **courier**, not the editor — the code work (editing `internal/fieldspec/registry.json` + Go, in a `frank/` worktree the operator sets up) is still git; frank carries your **governance relays** (audits, design, plan, plan-review, IMPL reports, review-folds, the exit-gate SITREP). Everything a `.md` relay used to be, `submit` now is.

Relay transport: **frank** (`submit`/`project`/`read`). Charter + this boot live in cwd `master/`; your build lands in `frank/`.
Current authority: **report-only onboarding.** This boot grants no PLAN/IMPL authority — that comes from the s5-dispatch you'll read from frank.
Acknowledge (by `submit`-ing a SITREP `TO master.orchestrator-planner`): your identity (`s5.orchestrator-planner`), the loaded skill, that `project`/`read`/`submit` work against frank, the routing directive (m-x via master only), and that you've read your dispatch; then proceed per the dispatch — onboard, plan S5.

ACTIONS_GIT_REF: none — report-only boot onboarding; no code/`frank/` edit by this relay.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main` (post-`s4-close`), clean.
