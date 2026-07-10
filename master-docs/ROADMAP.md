# `frank` Roadmap — overarching sequence (rough · v1 · step 0)

**Status: DRAFT (v1 — folds the VP's `master-roadmap` revise edits).** The rough overarching order of what
gets built for frank — the model/provider-agnostic, governed, multi-agent dev-team harness shipped as a **TUI
app** — and why in this order. Deliberately *rough*: boundaries will move as the master team designs. This
roadmap is itself **step 0**, and it **supersedes the immediate `C1_PRODUCT_SCOPE` A/B/C question** — the
first cycle's scope falls out of Step 1.

## Destination
A standalone terminal app: a **model/provider-agnostic agent runtime** (à la Claude Code) × a **Zellij-style
multiplexer** (see the whole team of seats at once) × an **integrated email client** (the relay/inbox
governance graph as first-class comms), with **the conductor protocol** as the governance layer underneath.

## Tech stack (direction — decided 2026-07-03; revisit at Step 4, not locked)
**Two-language spine, joined by the MCP seam.**
- **Conductor / MCP server + the whole governed server side → Go.** The trusted courier + runtime + provider
  adapters + scheduler/outbox as one auditable static binary. (The Step-1 conductor is already in Go.)
- **TUI / multiplexer → Rust** (`ratatui`/`crossterm`) — a *separate process* that reaches the conductor only
  through `submit`/`project`/`read`, i.e. **just another MCP client** (the guardrail, dogfooded). Rationale: a
  strict compiler as a correctness forcing-function on the largest, most AI-generated surface.
- **Seam = MCP/JSON-RPC over a socket** — a clean process boundary, no FFI. The two type systems do **not** check
  each other across it, so **codegen the boundary types from one source (FieldSpec / JSON Schema) into both Go
  and Rust** — a schema change then breaks both builds, not one runtime.
- **Not TS on the server side** (fractures the single-binary + trust story). A web dashboard, if ever, is the
  only natural TS carve-out.
- **Velocity fallback:** all-Go (Bubble Tea) is viable through Steps 1–3; introduce Rust at Step 4 when the
  multiplexer ambition demands it — the MCP seam makes that an additive front-end swap, not a rewrite.

**Execution watch-items (Step 3/4):** give Go a stricter net (`staticcheck` / `errcheck` / `exhaustive` /
`golangci-lint`) since it carries the trust-critical code on the looser compiler; budget Rust agent-iteration +
review for compiler-appeasement smells (`clone` / `unwrap` / `Rc<RefCell>` / `unsafe`); the PTY/session-supervision
seam defaults to the Go runtime (aligns with m-7 attach/pipe lifecycle) — prototype before locking.

## Sequencing principles
1. **Own the gate first.** Governance (the conductor) is the differentiator; the agent runtime is rentable.
   Automate governance early — riding existing runtimes — and build our own runtime later.
2. **Foundations before consumers.** The relay store + identity (m-1) and the form schema (m-2) underlie
   everything; they lock first — but *with* their consumers bound (see Cross-cutting), never in isolation.
3. **Ride, then replace.** Early steps run on existing harnesses (Claude Code / Codex sessions); later steps
   replace that dependency with our own runtime + TUI.
4. **Value early.** Step 1 must remove the operator-as-transport pain before anything fancier.

## Cross-cutting rules (folded from VP review)
- **Designed-early, executed-later.** The consumer schemas — observe fields (m-3), routing records (m-4),
  human-gate/email fields (m-6) — are *sketched and reviewed during Step 1* and **gate m-1/m-2 design-lock**,
  even though their runtime lands in later steps. Prevents a foundation that can't express what consumers need.
- **Mechanism before polish.** Governance comms (a minimal inbox/outbox + scheduler) is early; the full
  **email-client UX** is a Step-4 product surface.
- **Local-first, egress fail-closed.** Early gate-comms are **local-only** (the email client is UX over the
  *local* relay store; no external send). A minimal **fail-closed egress scan** (secrets / PII / model-names)
  gates the *first* external send whenever it lands; full egress hardening is Step 6.
- **Interjection is first-class — the human can always redirect a running lane.** Three grades, each a relay the
  conductor (sole-writer courier) times and routes:
  - **steer** — queue the message → deliver at the lane's **next step boundary** (between tool calls), **inline**,
    no side panel, no whole-task wait; and **interrupt** — soft-cancel the in-flight generation → redeliver now.
    **These two are baseline/standard, not optional.**
  - **side-question** (`/btw`) — spawn a **read-only, tool-blocked, single-turn forked archetype** (an m-5 sensor
    with a read-only authority ceiling) **routed to a cheap/fast model** (m-4) that answers **in parallel without
    interrupting the lane**, on a separate surface.
  Ownership: **m-6** owns the surface + the steer/side-question/interrupt choice; the **runtime** owns the
  boundary-injection + soft-cancel; **m-5/m-4** own the side-question archetype + its routing. In the
  ride-existing-runtimes phase, steer/interrupt are **bounded by the host harness** (Claude Code supports both;
  others vary); **owning the runtime (Step 3) makes all three first-class + uniform.** The API floor — no true
  mid-*generation* injection — is a documented constraint; all three mechanisms live *around* it, not through it.
  (Prior-art + the do-not-copy negative look in `jcode-ux-notes.md`; positive look in
  `codex-notes.md` — local reference copies, not vendored.)
- **Every step has an observable exit test** (below), not a vibe.

## The sequence (rough)

**Step 0 — Research & high-level design** *(now)*
Audit current state + prior-art (the upstream protocol, jcode, Claude Code, Zellij, email TUIs); design the
architecture-of-record + per-domain designs. No code.
*Exit test:* locked design-of-record + per-domain design docs + this roadmap; the Step-1 foundation contracts
(relay-store API, identity stamp, form schema) reviewed by the m-3/m-4/m-6 consumer lenses.

**Step 1 — Conductor core / "automated operator-relay"**
Goal: remove the operator-as-transport. Owns the relay store (sole-writer, append-only), seat-stamping /
channel-stamped FROM, the inline lint/form gate, and a minimal **local** governance outbox (gate items
written locally, not externally sent). Rides existing agent runtimes — doesn't run its own agents yet.
*Design-lock dependency (F2/F5):* m-1/m-2 interface sketches must be reviewed for **m-3 observe fields,
m-4 routing-record schema, m-6 human-gate/email fields** before lock.
*Exit test:* on a fixture/dry-run, a relay is accepted **only** through the conductor; `FROM` is
system-stamped (not lane-supplied); lint/form validation runs before delivery; a gate produces a local
outbox item.
*(Operator ruling 2026-07-06, post-s5: the transport fix — the s5-dogfood F1–F17 ledger,
`master/TRANSPORT-FINDINGS-2026-07-06.md`, headlined by the F11 lineage livelock — is **in-step**: Step-1
does not close until it lands, and this exit test runs **on the fixed conductor**, upgraded with two live
legs: the fixed conductor's first act = §7-applying s5's registry to a fresh blessed store, and an F11
regression leg = the archived dogfood traffic pattern replayed without livelock.)*
**✦ STEP 1 CLOSED 2026-07-08** — six slices (spine · engine · forms · wire · consumer schemas · the
transport fix), tag `s6-close` at `main@6a1198a`; **the exit test passed LIVE including both upgraded
legs** (the operator §7-apply; the F11 redrive 14/14, zero parent-class, zero livelock). Record:
`master/RECONCILE.md` § s6 / Step-1 closure.

**Step 2 — Governance hardening + minimal comms**
Goal: gates become unfakeable; lanes self-pace. Owns the observe-as-send-gate + evidence ladder + executable
claims (**m-3**), and a minimal governance inbox/outbox + scheduler (park gated lanes, wake-on-reply)
(**m-6 mechanism — not the full client UX**).
*Exit test:* one **failed** observe-as-send case (a false "done" is rejected pre-send) and one **passing**
send with conductor-captured evidence; a parked lane wakes on reply.

**Step 3 — Model-agnostic runtime + routing execution**
Goal: drive multiple models directly; begin replacing the ride-on dependency. Owns provider adapters
(Claude / GPT / Gemini / …) + the model→seat **router executing the routing record designed in Step 1** +
the benchmark (**m-4 + Runtime Core + Provider Adapters**). Also owns a **native governed agent-spawn
primitive** — frank's own "agent-teams" — so a seat can spawn sub-agents/panels *inside* frank (e.g. the
code-review panel) instead of depending on the host harness's orchestration (Claude Code's Agent Teams). Each
spawned agent is **governed** — channel-stamped identity + authority-ceiling-at-spawn (m-5), not an ungoverned
fork — so a review panel is N read-only, review-ceilinged, individually-stamped sub-agents. This **lifts the
"capability-dictates-topology" constraint** (panels stop being pinned to whichever seat rides an
orchestration-capable host) and is the runtime primitive that **powers the Step-5 nested/recursive teams**.
*Exit test:* ≥2 providers callable through one interface; a recorded routing decision carrying a
justified-deviation field; benchmark output scoring a routing run; **a steer message applied to a running lane at
its next step boundary, and a soft-interrupt that cancels-and-redelivers** (the first-class interjection mechanism).

**Step 4 — Standalone TUI + full email-client UX**
Goal: become one standalone app. Owns the Zellij-style multiplexer (tabs/panes/layouts showing the team),
the polished **email-client UX** (over the Step-2 comms mechanism), UI/UX, and Platform Plumbing.
*Exit test:* a single TUI process shows live sessions + inbox state with **no external multiplexer
dependency**.

**Step 5 — Workflows, archetypes & recursion**
Goal: workflow-per-archetype + nested teams. Owns expansion-slot presets, task-archetypes
(research / QA / feature), per-archetype gates, and **nested/recursive orchestrator-team support** — the
operator's 4-tier vision as a *product capability* (**m-5**).
*Exit test:* a human picks a workflow at spawn that sets the authority ceiling; a nested orchestrator-team
runs under a parent.

**Step 6 — Packaging, distribution, hardening**
Goal: ship. Owns install / cross-platform delivery, **full egress + operational-safety hardening** (building
on the minimal egress gate added before the first external send), docs.
*Exit test:* a clean install runs on a fresh machine; the egress scan blocks a planted secret/PII in an
outbound send.

## Positioning — frank vs. the convenience stack (recorded 2026-07-08, post-Step-1-close)

The adjacent-layer landscape is now concrete and shipped: **AMQ** (`agent-message-queue`, local reference copy, not
vendored; MIT, brewed) = the local agent-mail bus — Maildir-atomic, threaded, federated, **"no server, no daemon"** ⇒
identity is self-asserted (`AM_ME`); **amq-squad** (`amq-squad`, local reference copy, not vendored; v2.16) = the
orchestration layer over it — roles,
goal-first dynamic teams, task store, gate threads, a verification-before-merge preflight. Together they are
**frank's convenience twin**: they independently converged on our problem list — their own ADR says *"child reports
are data, not authority"* — and answered every item at the **convention layer**: approvals are messages anyone can
author, task claims are same-user files, the merge gate validates evidence *"supplied by the lead"* (shape-checked,
provenance-free). **frank's layer is the one they scope out by design: the trusted middle** — identity stamped in
the channel, lineage computed at commit, authority enforced inside the serialized loop, and (Step-2) evidence
*observed* rather than supplied. The demo-able line: the same forged-approval / fabricated-evidence move accepts
silently on the convenience stack and bounces **typed, on the record** on frank. Implications: (1) Step-2
observe-as-send is the moat-completing step — no detours; (2) the pitch names the layer ("the governed courier"),
never the category they own ("agent messaging/teams"); (3) their composition/console/task-store designs are sanctioned
prior-art reading for Steps 3–5 (local reference copies, not vendored, beside the jcode/codex notes). Landscape verdicts unchanged from
the 2026-07-07 scan: M7 commoditized, M1 convergent-but-ours-is-stamped, M4/M2 novelty survives.

## Horizon (post-Step-6; recorded 2026-07-05, no pre-work sanctioned)
**Federation — frank↔frank relay exchange** (operator idea): agents on one conductor send to agents on another —
LAN→internet for governance mail; use case = cross-desktop / cross-user agent coordination. The architecture is
already federation-shaped: the **egress gate + local-outbox-only** rule is the border (federation = a new outbox
consumer feeding a remote conductor's *intake* as an authenticated submitter); the **m-6 away-bridge** (Seam C=A)
is the embryo; **channel-stamped FROM** generalizes to `seat@instance` domain-trust (DKIM-style vouching).
Implementation fork, deferred: ride SMTP/email infra vs a frank-native conductor-to-conductor channel class.
**Hard gate:** crossing machines/users breaks the same-uid attach model → requires real cryptographic identity =
the **D3-shelved forgery-robust-by-construction milestone** (same future gate, deliberately). Zero pre-work
needed now — the `@instance` address suffix is additive and `schema_version`+migrators (S3) exist precisely so
this needs no speculation today.

## Key open sequencing decisions (yours)
1. **Own-the-gate-first vs app-first** — conductor MVP (Steps 1-2) *before* the standalone runtime+TUI
   (Steps 3-4), riding existing harnesses meanwhile. Alternative: pull the TUI/runtime earlier for headline
   progress, at the cost of building UI before the governed substrate it shows is real. *(VP approved the
   current order.)*
2. **Comms split** — resolved via Cross-cutting: governance mechanism early, email-client UX at Step 4.
3. **Routing earliness** — resolved via Cross-cutting: routing-record contract designed in Step 1, router
   execution at Step 3.

## How this maps to cycles
Each step = one or more design→build cycles. We are in **Step 0** (design). Designing Step 1's foundations
*with its consumers bound* (the Cross-cutting "designed-early" rule) **is** `C1_PRODUCT_SCOPE = B` (conductor
+ runtime seams) — so the roadmap lands there, now with the consumer-review dependency made explicit.
