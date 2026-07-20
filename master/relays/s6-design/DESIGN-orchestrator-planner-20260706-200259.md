## DESIGN — s6 ADDENDUM leg (operator-requested): the FORMAL BOOT STAGE — seat lifecycle {minted → bound → active} + a typed BOOT form + identity-activation; one bounded pair round each (B-1 m-7 · B-2 m-2 · B-3 m-1), folds into the integration package; VP objection window on the scope addendum

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator requested this addition in-session (2026-07-06), which IS the exercised human gate for the scope addendum; the VP objection window below is the technical gate
GRILL_REQUIRED: no — mechanism formalization of an existing convention; no hard-to-reverse fork identified (each pair escalates if it finds one)
IN_REPLY_TO: master/relays/s6-design/RECONCILE-orchestrator-planner-20260706-195300.md
FROM: master.orchestrator-planner
TO: m-7.planner, m-1.planner, m-2.planner
CC: master.orchestrator-reviewer, operator, m-7.implementer, m-1.implementer, m-2.implementer
SUBJECT: add the boot stage to your s6 amendment docs as ONE bounded leg each — the conductor gains a seat lifecycle (B-1), the v2.8.8 boot convention becomes a typed first-class form (B-2), the boot record closes the mint→active identity loop A-3 opened (B-3); premise audited honestly (the observed boot BOUNCES are already fixed by A-1/branch-A — this leg buys lifecycle visibility + ordering, not bounce-proofing)

**Basis + the honest premise (design against this, not the folklore).** The operator requested a formal boot stage (2026-07-06, in-session). The archive audit says: **every observed boot-relay failure was `form_digest:re-render`** (intakes 000012×3, 000013, plus digest+parent compounds on the orchestrator's boot sends) — i.e. the F5/parent classes **your existing amendments already kill** (A-1 stable-schema digest; branch A computed PARENT). So this leg is NOT bounce-fixing. What is actually missing: **the conductor has no concept of a seat lifecycle.** Boot today is paste-convention — a minted seat is invisible until it volunteers an ordinary relay; nothing orders a seat's first actions; the roster (who is minted/wired/booted) exists only in the operator's head; and A-3's live-mint amendment creates the mint edge of a lifecycle whose remaining edges are undefined. Formalize it.

### The three legs (one bounded addendum to each existing amendment doc; same pair discipline)

- **B-1 (m-7 — the lifecycle + the roster):** seat states **`minted → bound → active`** — `minted` by the A-3 `seat_mint` pivot; `bound` on first authenticated channel connect; `active` on the seat's **first accepted relay**. You rule the recording shape (derived table vs records — note `seat_mint` is already a pivot; the transitions may be derivable rather than new record classes). Plus: **a roster view** — an operator/orchestrator-visible projection of seat states + liveness (a `project` parameter or operator-surface view, **never a fourth seat verb**), replacing "master learns a seat exists when it talks." And: **out-of-order semantics** — what happens when a not-yet-active seat submits non-boot work (refuse-with-typed-bounce vs allow-and-mark; rule it, confusion-resistant framing — this is affordance ordering, not adversarial containment).
- **B-2 (m-2 — the typed BOOT form):** the v2.8.8 boot convention (DISPATCH_ID `<run>-boot-<seat>`, report-only ack) becomes a **first-class minimal form**: the smallest honest required set (identity claim, skill/charter-loaded assertion, dispatch-read-or-awaiting), no exotic seat-scoped enums, no parent semantics (branch A), **un-bounceable by construction** post-A-1 for a seat that fills what the schema shows. You rule the token question (a new `BOOT` phase atom vs `SITREP` + a boot marker field) and state its enum/§J2 impact explicitly — a new phase token touches shared vocabulary, so say so if you choose it.
- **B-3 (m-1 — identity activation):** the accepted boot record is the **identity-activation edge** — the channel-stamped proof that credential, binding, and session are one live seat, closing the loop A-3's `seat_mint` opens. Rule: the mint→bound→active binding semantics (what each edge asserts, what it may never assert); **session-restart semantics** (a reconnect is a re-`bound`, NOT a re-boot — one activation per mint unless re-minted; or rule otherwise with grounds); and the D4/D5 honesty line (activation is confusion-resistant liveness bookkeeping, not an identity-strength upgrade — the channel stamp already carries identity).

### Constraints (unchanged, restated)
The seat surface stays exactly `submit`/`project`/`read`; byte-exact terminal enum; channel-stamped FROM; I-PH on every new surface (roster view included — no paths); the claim ceiling stays tool-mediated confusion-resistance; **the boot stage must not become an authority gate beyond its own ordering** (it sequences a seat's first actions; it grants nothing). No Step-2 observe pre-work. If any pair finds this leg non-bounded (a hidden fork, a lock conflict), **escalate — do not stretch**.

### Process
- One bounded round each: extend your existing s6 amendment doc (a `B-n` section), pair-grill it (planner designs, implementer grills), report the pair verdict to master. It **folds into the same integration package** the VP co-signs — no separate cycle.
- **[VP objection window]** — this relay ADDS scope to the VP-fenced s6-design dispatch on the operator's in-session request. VP: object within one relay if the addendum should not ride s6; silence past your next relay = the addendum stands for the co-sign, where you review its content anyway.
- Does not touch, hold, or reorder: the m-4 blocking confirm (Sharpening-D), the m-1 §A fold, or anything pair-locked.

### Not authorized by this relay
No code, no `frank/` edits, no lock mutation before the pair reviews + integration + VP co-sign, no build-slice work, no new seat verbs, no boot-prompt automation claims (the operator still pastes the boot text into the host session — that is the ride-existing-runtimes reality until the Step-3 runtime owns seat spawning; this leg makes the *conductor-side* state real, not the paste disappear).

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s6-design` — run below.
- The premise audit: archived store boot bounces enumerated this pass (all `form_digest:re-render` + two digest/parent compounds — the A-1/branch-A classes); the latency addendum in `TRANSPORT-FINDINGS-2026-07-06.md` (no perf work) stands.
- Chain: the operator request (in-session, 2026-07-06) · the s6-design dispatch + [VP-W1..W3] · `GRILL-LOCK-parenting-fork-2026-07-06` (branch A — why B-2 has no parent semantics) · m-7 A-3 (the mint edge this completes).

ACTIONS_GIT_REF: wrote this addendum relay + INDEX row; ran the read-only boot-bounce audit over the archive; no code, no `frank/` edit; cwd is a docs workspace (not a git repo).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `7e5c527` (tag `s5-close`), untouched.
Next requested action: operator hand-relays to the three pairs (any order; independent of the m-4 confirm + m-1 §A fold in flight); each pair extends its doc with its B-leg + pair-grills + reports; master folds the B-legs into the integration package → VP co-sign → the s6 build-slice dispatch.
