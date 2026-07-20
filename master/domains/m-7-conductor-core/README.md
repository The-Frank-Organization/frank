# m-7 — Conductor-Core (the runtime substrate)

**Status: DESIGN-LOCKED — c4 CLOSED 2026-07-02.** The conductor-core design-of-record is locked:
`design/2026-07-01-v3-conductor-core-design.md` (DESIGN_LOCK_ID `c4-design-m-7-lock`, §22 lock block;
GRILL_LOCK `c4-grill-m-7`; pair-approved r3 design + r5 lock package; **VP co-sign
`master/relays/c4-design-m-7/RECONCILE-orchestrator-reviewer-20260702-040327`**; CQ gate satisfied —
all eight design-LOCK CQs closed, CQ-6 on the base with `re-mint-supersedes` a §2C build-carry).
c6 doc-only re-review folds applied per `master/DESIGN-REREVIEW-2026-07-02.md` (fold-log: design doc §21).
Step-1 PLAN remains a separate operator-opened gate.

*(Domain origin, kept for history: the 7th standing pair, created in the 2026-07-01 re-baseline after the
adversarial design review found the running program the six policy domains ride on was never designed or
owned. m-7 owns that substrate.)*

Seats: `m-7.planner` (design-lead) · `m-7.implementer` (adversarial design-reviewer). Domain dir layout per the
charter (`audit/`, `design/`). c4 (audit + design, CLOSED) delivered the substrate audit
(`audit/2026-07-01-substrate-audit.md`, `audit/2026-07-01-audit-pair-reconcile.md`) and the locked
design-of-record covering `DESIGN-REVIEW-2026-07-01.md` §2A.

## The one-line boundary
**Conductor-core owns the ENGINE (how things run); the six domains own the CONTRACTS (what is valid / required /
gated). Conductor-core EXECUTES their contracts — in the right order, with the right atomicity, behind the right
interface. It does not re-own policy.**

## Owns (the substrate — conductor-core's outright)
- **Process / concurrency model + the single-threaded SERIALIZED COMMIT LOOP** — `submit()` read-validate-append and
  `verify()` check-and-burn as one serialized critical section (this is what makes "replay-closed" and
  "authority-consume-once" true; kills the two-honest-seats-double-accept race by construction).
- **Crash-atomic multi-file commit + recovery/reconciliation** — record + INDEX + N mailboxes committed atomically or
  recovered; no lost governance approval, no bricked store; corrupt-record quarantine.
- **Internal-FAULT disposition** — a trusted-side check that throws / times out / hits a corrupt record yields a
  distinct **held / fail-closed** outcome for authority records (never silent-accept, never unrecoverable-brick).
- **Trusted CONFIG load + integrity** — loads the (m-2/m-3/m-4/m-5/m-6-authored — c6 x3-F2 full set) policy config artifact **once at trusted
  startup**, integrity-checked; config is conductor-owned and **not in any seat tool surface**.
- **Attach/pipe lifecycle + the INTERFACE-GUARDRAIL enforcement** — the persistent-seat pipe lifecycle; the seat tool
  surface exposes **only** `submit()`/`project()`/`read()`; raw store/config filesystem paths are **absent** from it
  (the guardrail that makes the confused-agent threat model hold — `GRILL-LOCK-deployment-fork-2026-07-01.md` D2/D2b).
- **Local-outbox-only external-send posture**; **store genesis + GC/retention**; conductor-restart seat-binding recovery.

## Hosts + executes (existing-domain CONTRACT × conductor-core EXECUTION — joint seams)
- **m-1** — store write + channel-stamped FROM stamp: m-1 owns the store schema + the stamp *contract*; m-7 is the
  process that performs the append and the stamp inside the serialized loop.
- **m-2** — form/lineage gate + **fill-time-authority form rendering**: m-2 owns the FieldSpec + predicate *contract*;
  m-7 renders (forbidden options absent) and validates constrained picks. Only m-2's by-construction *claim* collapses
  (`DESIGN-REVIEW` §2B); the mechanism runs here.
- **Phase-split required-set** (`§2A.5`): m-1/m-2 own which fields are required-when; m-7 must not demand
  observe-owned fields in a Step-1 that has no observe writer — a joint m-1/m-2/m-7 fix.
- **m-3** — observe gate (Step-2 hook) + **decision-② fail-closed** on authority-class `record_integrity ∈ {self_reported, mixed}` (c6-widened per the CQ-2 canonical; design doc §12 S7): m-3 owns the
  observe/evidence *contract*; m-7 hosts the hook and enforces the class-conditional fail-closed at the gate.
- **Pure-judgment A-floor** (`§2A.6`): m-2/m-6 own the `gate_category`/`HUMAN_GATE`/`CEREMONY_TIER` floor *contract*;
  m-7 enforces the mandatory floor at fill/submit.
- **m-4 / m-5 / m-6** — routing record, archetype spawn, human-surface delivery: contracts theirs; m-7 sequences/hosts.

## Consumes (collision edges)
All six policy contracts (it is the host). Foundational upstreams: m-1 store API + m-2 schema. **Collision arbitration
+ the substrate↔policy seam split stay with CTO+VP.** (The ownership split proposed here was VP-approved at the
conductor-core standup and is now realized in the locked design-of-record — the hosted-contract seam matrix,
design doc §12, is the authoritative per-seam statement.)

## Explicitly NOT conductor-core (stays policy-domain-owned)
Model→seat routing decisions (m-4), archetype/ceiling semantics (m-5), gate→email bucketing + ODB + park/wake policy
(m-6), the observe done-predicates + evidence ladder (m-3), the FieldSpec vocabulary (m-2), the identity/trust model
(m-1). m-7 runs these; it does not decide them.

## Deployment posture (inherited, locked)
Attach + interface-level guardrail; adversarial containment / wrap / "by-construction" **shelved**
(`GRILL-LOCK-deployment-fork-2026-07-01.md`). Standing claim: **confusion-resistant; malicious code-executing agent
explicitly out of scope.** m-7 builds the confused-agent guardrail as a real requirement + fixture, not the
adversarial isolation.

## Step-3 reframe delta (RATIFIED 2026-07-15 — `master/STEP-3-ARCH-AMENDMENT.md` @ `2d240eb6…`)
m-7 **retains its conductor-host scope UNCHANGED** (store · single-threaded serialized commit loop · crash-atomic recovery · trusted-config · attach/pipe · seat interface-guardrail). Under the reframe **the conductor is ONE isolated service, NOT the app hub** — m-7 hosts the six conductor policy domains (m-1..m-6) but does **NOT** host the app shell (m-8/m-9/m-10). m-7 **LOSES the provider-credential contract**: it re-owns connector-side (m-8-authored, m-1 secret boundary); **no engine-v5 `provider_bindings` credential member** is created. The r3 credential design (`step3-amend-m7-cred`) is **provisional audit input** (r3 `must-revise`, F11–F13 open) for the fresh owner — no r4/lock from m-7.

## Step-3 MVP amendment delta (RATIFIED 2026-07-16 — `master/STEP-3-MVP-AMENDMENT.md` r7 @ `2f75f2a1…`; the F44 fold)
m-7's conductor-host scope stays UNCHANGED (the amendment touches no conductor byte/member/output). **NEW stage-1 authorship (§5/§7 + grill #3/F64):** m-7 authors **(a)** the **shared transport/client boundary** — hoisting `Call`/reconnect/socket lifecycle out of `cmd/frank-mcp` into the shared conductor-client (`internal/channel`; the m-2-owned FieldSpec mapping must NOT be absorbed — it lands in an m-2 module; parity tests MCP↔native required) — and **(b)** the **authenticated channel/BROKER contract** under m-1's identity semantics: the broker holds the LOGICAL m-9 seat's credential (one credential per logical seat, never copied into worker generations), lives **OUTSIDE the replaceable worker generation** (placement beyond that — own process / protected thread or module in the app main — is this DESIGN's call), and **checks the current worker generation / `turn_epoch` on EVERY `relay.submit`, `relay.project`, `relay.read` AND on every push delivery/forwarding (F64)** — connect-time channel authorization alone is insufficient (`internal/channel/server.go:277-336,391-417` authorizes once at connect); the contract specifies the epoch-change linearization + in-flight-call disposition (complete-or-reject, recorded). m-10 supplies lifecycle/epoch state to the gate, receives no credential bytes, gains no conductor verb. **(F65)** the conductor service identity (build digest + governing config identity) is bound SEPARATELY in the exit-test record for the relay-exchange leg — never an app-release field. The stage-1 DESIGN dispatch is issued with this fold.
