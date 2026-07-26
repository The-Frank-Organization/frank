# `frank` Roadmap — overarching sequence (rough · v1 · step 0)

**Status: DRAFT (v1 — folds the VP's `master-roadmap` revise edits; + the 2026-07-15 Architecture reframe re-cut of Steps 3–4, RATIFIED + operative, see below).** The rough overarching order of what
gets built for frank — the model/provider-agnostic, governed, multi-agent dev-team harness shipped as a **TUI
app** — and why in this order. Deliberately *rough*: boundaries will move as the master team designs. This
roadmap is itself **step 0**, and it **supersedes the immediate `C1_PRODUCT_SCOPE` A/B/C question** — the
first cycle's scope falls out of Step 1.

> **Maintenance (operator directive 2026-07-15):** this file is updated **only by the master orchestrator**
> (`master.orchestrator-planner`), and **must be updated on every milestone** (slice/step close, phase
> open/close, ratified architecture decision). Other seats / research subagents may gather facts and propose
> deltas; only the master orchestrator applies the edits.

## Destination
A standalone terminal app: a **model/provider-agnostic agent runtime** (à la Claude Code) × a **Zellij-style
multiplexer** (see the whole team of seats at once) × an **integrated email client** (the relay/inbox
governance graph as first-class comms), with **the conductor protocol** as the governance layer underneath.

## Architecture reframe — the conductor is ONE service, not the app hub *(RATIFIED + operative 2026-07-15 — architecture-of-record = `master/STEP-3-ARCH-AMENDMENT.md`, operator-ratified at SHA-256 `2d240eb6…`, VP-approved `step3-arch-packet/063000`; this section is now OPERATIVE, not provisional)*
The end-state is a **bundle of local services — a modular monolith + supervised worker processes over local
IPC, NOT networked microservices** (no API gateway; per-family stores/writers). **The conductor is just one
service in that bundle, not the app's central hub.** This corrects a mid-Step-3 drift ("add HTTPS to the
conductor for provider APIs") that was a category error — the conductor never dials a provider.
- **The conductor = the governed relay plane for stamped participants** — agent seats, orchestrator seats, the
  **operator channel**, and reserved system-governance records — plus its **own isolated store + sole governed
  writer**. It is NOT the app supervisor, run DB, provider client, turn engine, tool broker, terminal
  multiplexer, or general IPC bus. **Steps 1–2 stand unchanged**; any conductor change is a separately-flagged
  amendment, never a silent reopen.
- **The app shell around it** (greenfield): **m-10 App Control Plane / Supervisor** (hosts + sequences worker
  lifecycle / run-state / scheduling; owns **no** policy) · **m-8 Provider Adapters / LLM connector** (holds
  provider credentials + does the provider wire — **app-side, never in the conductor**) · **m-9 Model Runtime**
  (the turn/session/context loop — an **app-side worker**, not conductor-internal) · human/terminal surfaces ·
  per-family persistence.
- **Provider traffic bypasses the conductor**: a worker reaches a provider **worker → connector** over app IPC,
  and uses the conductor only to relay a governed message to another seat. Provider-send **policy/evidence**
  stays **m-3-owned**; the **mechanism/enforcement** is app-side (m-8 / m-10 last pre-wire); **secrets stay
  m-1-governed**.
- **The operator keeps a trusted, out-of-band direct-interaction path to agents** that the conductor does not
  mediate — a *named non-governed route* (operator = the trusted relay + final authority; confusion-not-malice),
  not an accidental hole.
This reframe **re-cuts Steps 3–4** below (Step 3 shrinks to an MVP; the rest becomes the Step 4+ ship arc). The
detailed boundary/traffic/state matrices land in the architecture-amendment packet → `master/ARCHITECTURE.md`,
not here.

## Tech stack (direction — decided 2026-07-03; revisit at Step 4, not locked)
**Two-language spine, joined by the MCP seam.**
- **The governed server side → Go.** *(Reconciled to the 2026-07-15 reframe: NOT one fused binary. The app is a
  **modular monolith + supervised workers** — the **conductor is its own isolated governed-relay service** with its
  own store/writer; the **app shell (m-10 control plane · m-8 connector · m-9 runtime)** are distinct app-side
  components/processes over local IPC. Go remains the language for the governed server side; the physical
  conductor↔app-shell split is real, not a single static binary.)* (The Step-1 conductor is already in Go.)
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
seam defaults to the Go runtime — **(reframe 2026-07-15) app-side, owned by m-10, NOT m-7 attach/pipe (the conductor is one isolated service; PTY/supervision is app-shell)** — prototype before locking.

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
- **Interjection is first-class — the human can always redirect a running lane.** *(Reframe 2026-07-15: interjection is Step-4 work, and under the reframe steer/interrupt/redirect are **APP-SIDE runtime controls** — the m-10 scheduler + m-9 turn loop — **NOT conductor-timed**; the conductor times/routes governed **relay** traffic only, never runtime control-plane traffic. The description below is the pre-reframe framing, retained as design intent for the app-side mechanism.)* Three grades, each a relay the
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
  others vary); **owning the runtime (the Step-3 MVP unlocks it; first-class + uniform delivery is Step 4+ per the 2026-07-15 re-cut) makes all three first-class + uniform.** The API floor — no true
  mid-*generation* injection — is a documented constraint; all three mechanisms live *around* it, not through it.
  (Prior-art + the do-not-copy negative look in `jcode-ux-notes.md`; positive look in
  `codex-notes.md` — local reference copies, not vendored.)
- **Every step has an observable exit test** (below), not a vibe.

## The sequence (rough)

**Step 0 — Research & high-level design** *(COMPLETE — current milestone is Step-3 DESIGN, the app-shell MVP; see the Architecture reframe)*
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
**▸ IN PROGRESS (opened by the operator 2026-07-10; kickoff = `master/STEP-2-KICKOFF.md`; slices s7–s11, order s7→s8→s10→s9∥s11).** **s7 (INV-CATALOG) CLOSED 2026-07-10** — tag `s7-close` at `main@2e1b4f0`: the standing global laws as ten named executable checks in every battery (`test/invariants`); found + fixed F-S7-R2-COLGRAIN (the s7a column-grain guard, its own gated lane) and surfaced `OI-S7A-CLOSE-ONCE-RACE` — since resolved by the s7b micro-lane (CLOSED 2026-07-11, `main@691d034`, with `FLAKE-SOCKET-PAR`/`CRASHPOINT-KILL-RETURN`; the s8 live-channel gate lifted). Record: `master/RECONCILE.md` §s7 + §s7b. **s8 (THE OBSERVE SPINE) CLOSED 2026-07-12** — tag `s8-close` at `main@8941889`: observe-as-send live in the atomic submit through governed supply only; false-done rejected typed at the real socket; exit legs 1+2 of the step-exit LIVE. Two ritual amendments (claim-input v6→v7; supply-set engine v1→2) + five-row hardening rode the build; eight failure classes caught in-slice. Record: `master/RECONCILE.md` §s8. **s10 (THE COMMS SPINE) CLOSED 2026-07-13** — tag `s10-close` at `main@39474d0` (pushed frank-dev 2026-07-13): **EXIT LEG 3 PROVEN LIVE** (a parked lane wakes exactly once on the operator's validated reply, on the fresh v8 dogfood store) — **all three step-exit legs now exist**; both operator sunsets demonstrated gone. Record: `master/RECONCILE.md` §s10. **s9 (EVIDENCE THICKEN) CLOSED 2026-07-14** — tag `s9-close` at `main@d91fcfb` (pushed frank-dev): the observed evidence layer thickened (find-references E1 · the §6.1 verdict pass, CheckVerdict untouched · the master-activated `lane_vcs:none` opaque-accept branch, branch-only, Option-2 E0 floor preserved · FX-VCS matrix · the ⑤ egress pair). Record: `master/RECONCILE.md` §s9. **s11 (THE COMMS THICKEN) CLOSED 2026-07-14 — the LAST Step-2 build slice** — tag `s11-close` at `main@502e06c` (pushed frank-dev): the B/C/D bucket projections + the complete 7-state FSM + the full g1 §B 8a hardening (both reason tokens `stale_schema`/`stale_choice_set` byte-exact, frozen-π guard, crash-replay) + the ③ known-A NF + T8 cleanup (eight of nine — item 2 soft-expiry arbiter deferred to a post-Step-2 m-7+m-3 design cell) + the G4 cadence re-homed at **engine v4** (the owner-caught r13 contradiction, ruled (a) and countersigned). g1 CLOSED (T6 built); **T5/T10 HELD OPEN behind g2/dc, acceptance-OPEN per FINDING-4** — master's step-exit disposition. Catch-ledger completed the cycle-datum pair: the end review catches what a green battery cannot (s10); the owner confirm catches what the review panel cannot (s11). Record: `master/RECONCILE.md` §s11. **★ THE STEP-2 BUILD IS COMPLETE — all five slices closed (s7 · s8 · s10 · s9 · s11).** The ONLY remaining Step-2 act is the master-owed **step-exit test** (all three legs live on the dogfood store + the INV-CATALOG red-battery demo + uncached green) → on its pass, Step-2 closes. **Frank-dev is synced through s11-close; the public flip stays a separate remote + decision.** **✦ STEP 2 CLOSED 2026-07-14 (operator-ratified in-session; VP adversarial close-confirm filed `approve` 2026-07-14, `step2-step-exit/RECONCILE-orchestrator-reviewer-20260714-211500`) — the step-exit test PASSED at `main@502e06c`:** all three legs green live on the dogfood store (false-done rejected pre-send + predicate named, via fresh-genesis AND the real production daemon socket · passing send with conductor-captured E1/E2 evidence · a parked lane waking exactly-once on the operator's validated reply on the fresh v8 store) · the INV-CATALOG red-battery fired (a planted fourth `delivery_state` turned the battery RED with `TestLawTerminalEnumByteExact` naming itself + three interlocked laws, reverted to green, `main` byte-pristine) · uncached battery exit 0 / 26 ok / 0 FAIL / vet clean. Master-owed executable exit SATISFIED; the FULL LIVE RELAUNCH (real seats, honest E3/E4) stays a SEPARATE operator act, not a gate. **T5/T10 acceptance-OPEN behind un-returned gates g2/dc — master recommends rescope as named Step-3 carries, not hold-open.** Close relay `step2-step-exit/RECONCILE-…-182600` TO operator + VP. Both returns received (operator ratify + VP `approve`, both 2026-07-14); the charter phase line flipped to Step-3.

**Step 3 — MVP: a barely-enough coding agent on the governed courier** *(scope CONVERGED 2026-07-15; re-cut per the Architecture reframe above)*
Goal: a **minimal coding agent that does real work and collaborates with other agents through the governed
conductor.** The smallest vertical slice — a **thin m-10 control plane** (run manifest + worker supervision +
the **empty permission seam** + the wake loop) + **one m-8 connector lane** (Codex; holds provider creds, does
the wire) + the **m-9 coding-agent worker** with **local tools `read/write/edit/bash/apply_patch`** + the
**conductor as a NATIVE relay tool** (`submit/project/read` over `internal/channel`, agent-to-agent comms; the
**MCP server retained** for foreign harnesses; a **shared conductor-client refactor** de-dupes both frontends) —
driven by frank's **own** runtime, not a ridden host harness.
**Authority is a built-but-EMPTY seam:** NO config-derived permission ceiling in the MVP — authorization is the
**fixed tool-DISPATCH seam** (the operator-ratified **8 canonical tool names** as the policy identity; build
identity binds at the interface-lock + a post-build release-binding event; **authorized == executed via a
one-shot epoch-fenced ticket**); the whole permission/authority
system (m-5 ceiling + `config_generation` freshness + per-role permissions) is **Step-4**. `bash` = ambient host
authority, **operator-accepted** (no sandbox; a trusted executor is Step-4 hardening H-12). **Stretch goal:**
push-based **wake-on-relay** (an agent wakes when a relay lands in its inbox — the `internal/channel`
`PushTo`/`NextPush` plumbing already exists; best-effort push + durable rediscovery + at-most-once scheduling).
*Exit test:* a **live** coding agent (Codex, honest **E3** floor) runs a governed turn, uses its local tools to
do real work AND uses the native conductor tool to exchange a relay with another agent — governed on the
conductor relay plane, provider bytes **never transiting the conductor**, the permission seam present
(permissive); the stretch adds a second agent waking exactly-once on the incoming relay.
*Design history:* DESIGN opened 2026-07-14 (Division II — m-8/m-9 greenfield; kickoff LOCKED after 4 rounds) →
**HELD 2026-07-15** on the Architecture reframe (`step3-hold-*`, bytes preserved) → re-scoped to the app-shell
MVP → **MVP scope CONVERGED 2026-07-15** (coding agent + native conductor tool + empty permission seam + wake
stretch; the whole authority system deferred to Step-4 — this **dissolved the seam-13 / freshness /
m-5-amendment knot**, and the m-5 MVP amendment stood down). The team runs ON frank as courier (carriage+observe
live at ~/frank-live) → **the explicit MVP architecture amendment RATIFIED 2026-07-16**:
`master/STEP-3-MVP-AMENDMENT.md` **r7 @ SHA-256 `2f75f2a1…`** (operator-ratified after the VP's byte-bound
approve; 7 revisions under exact-byte review, F39–F66; a 3-decision operator grill — the process topology pin
(m-10 = a module in the app main process; m-8 a separate connector process), the F59 one-shot authorization
ticket, and the F60 one-credential-per-logical-seat broker model). It amends the reframe packet at four named
fragments; **the §7 acyclic graph is the first-stage plan** (5 stage-1 owner contracts → m-8 → the lifecycle
halves → the m-9/m-10 designs → the Master+VP interface-lock → T4 build → the post-build release-binding →
live E3). Record: `master/RECONCILE.md` §Step-3-MVP-amendment + `master/ARCHITECTURE.md` "Step-3 MVP" § +
`master/relays/step3-arch-packet/` (ratification `…-040405`).
**2026-07-19 — §7 stages 1–3 CLOSED (VP close-confirm `step3-arch-packet/…-224500`, close-review r6).**
All seven interface contracts pair-approved byte-bound and cross-confirmed (16-edge census + the field-grain
m-9↔m-10 reciprocal): m-1 `7c8b09a6…` · m-2 `83d8e63e…` · m-3 r4 `009df607…` · m-7 r11 `9331ea88…` ·
m-10 r36 `0240e874…` · m-8 r12 `4b670a79…` · m-9 r19 `2a96a07b…`.
Six close-review rounds (F70–F84, every finding folded owner-real); the F59 authorize→consume→execute guard
closed at the three-identity/two-derivation-point form with both ratified negatives constructible.
**2026-07-21 — §7 stages 4–5 DESIGNED, the stage-6 interface-lock RE-SCOPED + operator-RATIFIED.** The stage-4
(m-9 full worker `cb7ff970…`) + stage-5 (m-10 control plane `6fd1d655…`) designs completed and both stage-6
lock halves closed on exact bytes (master r4 `5b36c64c…` + VP APPROVE `2c1b1437…`) — then, **before ratifying**,
a **stage-5.1 third-party review** (`master/STAGE-5.1-EXTERNAL-REVIEW-2026-07-21.md` `b4e79f3b…`) faulted the
milestone CLAIM + lock SCOPE (strong governance kernel, not yet honestly a coding-agent MVP: ungoverned bash /
no exact-effect binding / context lost on worker replacement / plumbing-only exit test). The operator chose to
keep the "frank harness MVP" label and pull scope up. Master HELD the all-artifact lock and authored a **bounded
stage-6 re-scope amendment** — `master/STEP-3-STAGE6-AMENDMENT.md`, driven through **twelve VP decomposition-review
rounds** to **rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`**, **VP-APPROVED** (r12,
`step3-arch-packet/…-163600`) and **operator-RATIFIED in-session 2026-07-21** (`…-165500`, agent-authored +
operator-cited per §8b). The re-scope: **sandbox FORGONE** for the MVP (ambient bash; **H-12 promoted to a HARD
pre-external-use blocker** — no untrusted/external/security-sensitive/multi-tenant use until a real sandbox
lands) · bash claim narrowed to invocation-context · a **six governance-property exit gate** (Governance ·
Durability · Crash-honesty · Injection-**visibility** · Handoff · Operability) + an **objective overhead budget**
(F59 p95≤250ms · relay≤1s · journal-commit≤100ms · per-turn wall-clock p50≤20% PASS/20–100% HOLD/>100% FAIL)
replacing any benchmark score · **utility DEMONSTRATED not gated** (public dogfood CRM+bivpak + honestly-labeled
agent-as-operator SWE-bench, no threshold) · **durable session-state + resume BUILT** (a worker-owned crash-safe
session-content log with a two-time-scoped trust invariant — content trusted only under settlement evidence AND
current-prefix presence, else `content_lost`/degrade, never fabricated; outcomes stay m-10-canonical) · the
interface lock re-cut into a **hashable Tier-HARD bundle** (`bundle_sha256`, stable under soft edits — **later
superseded 2026-07-27 by the plain byte-bound interface-lock RECORD `master/STEP-3-INTERFACE-LOCK.md` `cbd1893c…`**,
per the operator's MVP-minimality call; see the 2026-07-27 milestone below) + the pair
order made an **acyclic DAG**. The held joint lock `b7e1f0ef…` is SUPERSEDED (replaced by a later shorter
re-lock). **Decoupling (D5): real-work / dogfood start is ⊥ the exit gate** — CRM/bivpak may begin in parallel.
Record: `master/STEP-3-STAGE6-AMENDMENT.md` + `master/relays/step3-arch-packet/` (twelve r1–r12 rounds; ratify
`…-165500`) + `master/RECONCILE.md` §stages-1–3-close.
**Now open — the ratified §11 sequence (separately-gated lanes):** (1) the **m-7 broker study FIRST** (+H-24 if
cross-epoch completion survives) → (2) the **interface DAG legs** under the F73 ladder + join records for the
two-sided seams → (3) **author the interface-lock RECORD** `master/STEP-3-INTERFACE-LOCK.md` (**item A** — the
bundle mechanism of step 3 was superseded 2026-07-27; the exit-fixtures freeze moved to lane 4) → (4) **lane 4** =
the **Master+VP re-lock over the record's external SHA** + author/content-address the exit-fixture inputs → freeze
`STEP-3-EXIT-FIXTURES.json` → lock → (5) **T4** (behind the re-lock + H-16/H-26). Ratification issued NO lock/PLAN/T4/cred/provider/
release/E3/merge/deploy.
**2026-07-25 — §11 lane status.** **Lane 1 (m-7 broker study) DONE 2026-07-21**: pair-approved rev8
`64f9136e…` additive over m-7 r11 `9331ea88…`, **DETERMINATION: simpler rule set — cross-epoch completion
NOT retained ⇒ NO H-24**; m-9/m-10 F73 confirmations + a co-signed two-sided §D join record landed
(`master/relays/step3-relock-broker-confirm/`). **Lane 2 (interface DAG legs) CONVERGED this session** — the
producer wave across every owner is pair-approved byte-bound: m-1 env/redaction `d34a7c47…` · m-2 §5-E
component `c3a8cd61…` + `relay.submit` cell `5ec7a3d2…` · m-3 lane-2 E0/E3 delta **r19 `92e08d09…`** (an
HONEST PARTIAL — T1–T8 live; **N910** — the per-attempt DATA-P non-emission loss cut — ruled a **documented
MVP limit**, disclosed via m-10's `UNKNOWN_PROVIDER_OUTCOME` → `uncertain` surface, not a defect; the r7-mirror
question **deferred to v3**, both dispositions accepted by m-3 as consistent with r19 with no byte change
owed, `master/relays/step3-relock-dag-m3/SITREP-planner-20260725-113000.md`) · m-8 B/E producer r5 `c0b7b488…`
+ 2a/2b discriminator **r7 `734e44b7…`** · m-9 lane-2 delta **r12 `04422965…`** (the three §5-E-delegated
recipes — `compaction_template`=A3 attempt-kind-total, `policy_messages`=B1 Step-3 constant `[]`, both
DELEGATED to m-9 under VP classification F73 `4c254307…`, no operator gate — `master/relays/step3-relock-dag-m9/RECONCILE-planner-20260724-033000.md`) ·
m-10 B/E carriage **rev3 `cd17db32…`** pair-approved (a further producer delta at rev14 `b96a1511…` is live
but NOT yet pair-approved). **Item A / Lane 4 / Lane 5 NOT STARTED** *(as of 2026-07-25 — SUPERSEDED by the 2026-07-27 milestone below: lane 2
closed, item A ratified + authored as the record-lock, in VP+F73 review).*
**The current live gate — the §D-settlement amendment**: `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md`
**rev4 `1fa71cb8…`**, bound as one packet with the pair-approved m-2 `relay.submit` resource cell
`5ec7a3d2…`. VP-APPROVED zero findings 2026-07-25 (`master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-160000.md`)
and **OPERATOR-RATIFIED 2026-07-25 (recorded §8b at `step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-170000`; four corrections in force, propagation matrix open)** (no downstream action
precedes this gate). Four corrections carried: (1) the D-4 Gate-2 honest relabel; (2) a run-wide parked-set
restore + `MAX_PARKED_ROWS_PER_RUN=512` + a new terminal `parked_unknown_capacity_exceeded`; (3) the
`relay.submit` `canonical_resource` cell; (4) the `turn_failed` zero-attempt explicit supersession. On
ratification: the propagation matrix (fresh pair-reviewed m-9 + m-10 successors; m-2 unchanged) → the §D
two-sided join → then Item A → Lane 4 → Lane 5 (T4).
**H-12 stands throughout as a HARD pre-external-use blocker** (the MVP forgoes the sandbox): external /
untrusted / multi-tenant use PROHIBITED until a real sandbox lands, independent of where the §11 sequence sits.

**2026-07-27 — Lane 2 CLOSED; item A RATIFIED + AUTHORED.** The interface DAG (lane 2) **closed** over 9 settled
owner bases + 5 joins (close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`);
the §D-settlement amendment ratified. **Item A's mechanism was simplified**: after 10 VP REVISE-NARROW rounds, the
machine-checkable interface **bundle** (`bundle_sha256` / extractor / markers / soft-stability fixture) was replaced —
on the operator's MVP-minimality call — by a plain byte-bound **interface-lock record**. The item-A simplification
amendment `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` **rev7 `3443f73d…`** is **VP-APPROVED + OPERATOR-RATIFIED
2026-07-27** (ratify `step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-130000.md`); the r3 recipe
`06e6956e…` is WITHDRAWN; owners RELEASED. Master then **authored item A** — `master/STEP-3-INTERFACE-LOCK.md`
(external SHA `cbd1893c…`): the closed byte-bound manifest of 38 files / 42 semantic rows (owner bases · frozen finals · governing
amendments+ratify relays · joins · carried-source lineage) + 5 typed precedence edges + whole-file invalidation +
external (no-self-hash) binding. **VP + F73 accepted the record at record/F73 grain** (`step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-170000.md`, over the corrected-record transmittal `…/DESIGN-orchestrator-planner-20260727-160000.md`); the source-fold correction is `…/RECONCILE-orchestrator-planner-20260727-180000.md`. **Item A is CLOSED** at interface-lock `cbd1893c…` (VP + F73 final approve `step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-210000.md`; fresh rehash 38/38 distinct, zero mismatch). **Next: lane 4** (Master+VP re-lock over `cbd1893c…` + author fixture inputs → freeze
`STEP-3-EXIT-FIXTURES.json` → lock) → **lane 5 (T4).** H-12 unchanged.
**Deferred to Step 4+** (moved out of MVP scope): the **permission/authority system** (m-5 config-derived ceiling
+ `config_generation` freshness + per-role permissions) · a **hard sandbox + irreversibility gating** (H-12) ·
the **per-model tool manifest / extensible registry / model carousel** · second-provider portability + **routing
execution** + justified-deviation · benchmark · **native governed agent-spawn** · first-class steer/interrupt.

**Session versioning — git-like, own implementation** *(added 2026-07-26, operator-originated)*. The Step-3
re-scope removed hash-chaining from the durable session log, and the reason it had to go is the reason this
carry exists: chaining treats change as **corruption**, version control treats change as **history** — the
same cryptographic property (one value naming a whole history state) with the opposite stance toward
editing, and frank needs the second, because editing sessions is a requirement rather than an attack. Two
drivers, both operator-supplied: **repairability** — a session can become unresumable through no fault of
the user (precedent: malformed thinking blocks rendering Claude Code sessions unopenable, recovered by
hand-editing the file), and a session that cannot be repaired is a data-loss event; and **third-party
rewriting** — `bivpak` packs a session, rewrites paths inside its history and reopens it on another machine,
which under chaining is indistinguishable from forgery. The access pattern is genuinely a version-control
workload: append-dominant, with occasional rollback, fork, sparse repair edits and tail deletion for
recovery. **Direction:** a content-addressed object store plus a commit chain built for this purpose — no
index, no worktree, no checkout — composed with in-log **supersession records** for sanctioned tool-driven
edits (never mutate history, file a supersession: the discipline the relay store already lives by), and
honest labelling underneath both so hand-repair always works, since a file frank cannot parse is a file it
cannot append its own repair record to. Repo-per-session is rejected for operational sprawl; one shared
store is rejected because git-style write contention funnels every commit through a single writer and one
slow session stalls the rest. **git is explicitly NOT a dependency:** either clean-room against git's
documented on-disk format (loose objects, zlib framing, blob/tree/commit) or fork/vendor a minimal subset
with attribution — noting git is **GPLv2**, so copying its source would make frank GPLv2 and collide with
releasing publicly under a license of our choosing, whereas implementing the documented format does not;
loose objects plus a commit chain is bounded work against a stable spec, and **packfiles** are where the
cost lives and where session-sized data very likely never needs to go. Pre-reads before any design lock on
this carry (2026-07-26 sweep, PRIOR-ART.md §2d): codex `rollout_lineage` (`05f000263b` — forks over a frozen
source-history prefix, child-owned records) + its per-thread single-writer locks (`5c94796dc9`), and
OpenRath (arXiv 2606.19409, session as a branchable/replayable value) — the field shipped this direction
within days of the operator choosing it, which validates the direction and supplies the port sources.
**Status — PENDING, not binding.** Decision 5 (operator, 2026-07-26) requires that external session edits be
permitted, honestly labelled, and not by themselves a bar to resume — a session that cannot be repaired is a
data-loss event. The **exact MVP trust/reconciliation mechanism is undesigned**: how an edited prefix is
detected and labelled; whether edited provider/tool content is trusted, untrusted-but-model-visible, or
degraded; how it reconciles with the immutable settlement snapshot and prior receipt identity; which
disposition and first action result; and whether a sanctioned edit rebases or supersedes durable evidence.
That is owed from m-9 + m-10 (joined by m-3 for the observable consequence, reviewed at the m-1
at-rest/provenance boundary), then VP review, operator ratification, and an additive supersession — **no
Step-3 contract moves before then.** In particular this carry does NOT assert that any specific comparison
becomes non-gating, and does NOT touch `receipt_conflict`, which stays frozen unless the owner derivation
proves it implicated (VP LANE4-ESC1-VP3-F2/F4: 'label, never gate' cannot silently promote edited bytes as
prior provider/tool truth across the frozen evidence-AND-current-presence invariant). The
forward-compatibility observation stands on its own: an append-only line-oriented journal imports into a
content-addressed store cleanly, which the removed chained log would not have.

**Step 4+ — Ship the harness (the product arc, Steps 4–6)** *(re-cut 2026-07-15 — "MVP at Step 3, then ship at Step 4+")*
Goal: build the proven MVP spine out into a shippable standalone app. **Step 4 opens the shipping arc** and
**absorbs the runtime-maturation work that moved out of the Step-3 re-cut**: **routing execution** (≥2 providers
through one interface, the routing record + justified-deviation, benchmark), **native governed agent-spawn**
(channel-stamped identity + authority-ceiling-at-spawn (m-5), not an ungoverned fork — the primitive that powers
the Step-5 nested/recursive teams; lifts the "capability-dictates-topology" constraint), and **first-class
steer / interrupt / side-question**. Plus the standalone **Zellij-style multiplexer** (tabs/panes/layouts showing
the team), **terminal/tmux integration + worker supervision at scale**, provider failover, the polished
**email-client UX** (over the Step-2 comms mechanism), UI/UX, and Platform Plumbing.
*Exit test:* a single TUI process shows live sessions + inbox state with **no external multiplexer dependency**;
≥2 providers callable through one interface with a recorded justified-deviation routing decision + a benchmark
scoring a routing run; **a steer applied at a running lane's next step boundary + a soft-interrupt that
cancels-and-redelivers**.

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
3. **Routing earliness** — resolved via Cross-cutting: routing-record contract designed in Step 1; **router
   *execution* moved to Step 4+ by the 2026-07-15 re-cut** (Step 3 runs a single pinned route only — one
   governed turn, no routing execution). The governed routing *decision* record may live as a conductor relay
   (m-4-governed, m-2-shaped); connector/worker selection + invocation state are app-side.

## How this maps to cycles
Each step = one or more design→build cycles. **We are in Step-3 DESIGN — the app-shell MVP (see the Architecture reframe above); as of 2026-07-25 the §11 stage-6 re-lock is running — Lane 1 (broker study) CLOSED, Lane 2 (interface DAG) CONVERGED, and the §D-settlement amendment RATIFIED 2026-07-25 — its propagation matrix (m-9 + m-10 folds) now open (see the 2026-07-25 §11 lane-status entry above).** *(The rest of this sentence is Step-0 origin history:)* Designing Step 1's foundations
*with its consumers bound* (the Cross-cutting "designed-early" rule) **is** `C1_PRODUCT_SCOPE = B` (conductor
+ runtime seams) — so the roadmap lands there, now with the consumer-review dependency made explicit.
