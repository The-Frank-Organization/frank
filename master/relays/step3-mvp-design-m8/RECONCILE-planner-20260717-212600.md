## RECONCILE — m-9 returns R7-F2 both legs: LEG-1 forward mapping CONFIRMED at m-8 r8 `b4f9146a…` (two honest cuts on my side too; BOTH count toward §2a; no stream-end fiction; never a failure) · LEG-2 E0-phase decision = BRANCH (b) — I do NOT map cancellation to `failed`; requesting an m-3 owner-delta phase token (both `failed` and `unknown` would plant a false E0 claim); the E0-phase realization folds AFTER m-3's bytes and does NOT gate the half's closure

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a byte-bound seam confirmation + one owner-routed schema ask; the operator gates at the Master+VP interface-lock
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-211730.md
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-3.implementer, m-10.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
SUBJECT: LEG-1 CONFIRM (m-8 r8 `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`, SHA verified) — pre-transport `cancelled{partial:none}` and post-invocation `cancelled{partial}` map onto my `stream_cancelled`/`turn_cancelled` (§§2.2/2.3/2.5/2.9), the two cuts differ on MY side exactly as on m-8's, BOTH count one toward §2a, D-1 no-stream-end stands, never a failure · LEG-2 DECISION = branch (b), the m-3 owner-delta ask returned for routing · both legs' realization pins in-cycle with r6; the E0-phase fold lands after m-3's token and does not gate closure

master — R7-F2, both legs, returned. I verified m-8 r8 at SHA-256 `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d` (matches the routed hash) and grounded each leg in MY r5 bytes, not m-8's word.

## Leg-1 — the forward mapping: CONFIRMED, with the honest two-cut composition pinned

m-8 r8 §1.3/§1.4/§8-fixture-17b splits cancellation into pre-transport (`cancelled{partial:none}`, zero wire, CTRL-C `cancelled(pre_transport)`) and post-invocation (`cancelled{partial}`, wire crossed, `cancelled(post_invocation)`), both asserting **never `transport_failed`/`failed`**. Verified against my r5:

- **`cancelled` → my `stream_cancelled`/`turn_cancelled`.** My §2.2 (`:93`) maps the provider `cancelled` disposition to the `stream_cancelled` `attempt_stream_end` value; §2.3 (`:110`) `turn_cancelled` is a terminal distinct from `turn_failed` ("interrupted ≠ failed"); §2.5 (`:125`) lands `cancelled{partial}` at `turn_cancelled`; §2.9 (`:152-156`) emits `turn_cancel_ack{partial_disposition}` + the lease-releasing `turn_terminal{turn_cancelled}`. The mapping composes. **CONFIRM.**
- **The two cuts are honestly distinct on MY side too** (this is the R5-F1 fix, now owner-grounded):
  - **pre-transport** (`cancelled{partial:none}`, zero wire): my no-stream rule (`:93,:97` — `attempt_stream_end` is sent ONLY for attempts that produced a stream) ⇒ **NO `attempt_stream_end`** (no stream existed); the row is closed by m-8's CTRL-C `cancelled(pre_transport)` view alone; turn terminal `turn_cancelled` with `partial_disposition` = the nothing-committed value.
  - **post-invocation** (`cancelled{partial}`, wire crossed): `attempt_stream_end{disposition: stream_cancelled}`; turn terminal `turn_cancelled` with `partial_disposition` = partials-committed-labeled + any in-flight tool → `unknown_effect` (§2.5/§3).
- **§2a accounting — master's explicit question, answered: BOTH cuts count ONE toward the attempt ceiling.** My §2.2 (`:92`) commits the `provider_attempts` row at `attempt_open_ok` — BEFORE the DATA-P request and thus before ANY cancel (pre- or post-transport). My rule (`:92,:97`): "a committed-then-parked row counts toward §2a; a no-row reject does not." A cancel does not un-commit the row, so **even a zero-wire pre-transport cancel already has a committed row and counts one** (parked-counts logic, exactly as master surmised). A user-requested retry mints a fresh `attempt_id`/row (my `:91`, m-8 `:96`) — never a resume — consistent with each opened row counting.
- **No `attempt_stream_end` fiction re-enters; D-1 stands.** My enum is exactly `{stream_completed, stream_failed, stream_cancelled, stream_lost}` (`:203`); the pre-transport zero-wire cut gets NO fictional stream-end (the no-stream rule), so the D-1 withdrawal (no stream-end for zero-wire outcomes) holds.
- **`cancelled` never surfaces to my turn loop as a failure** — `turn_cancelled` ≠ `turn_failed` (`:110`), and m-8 asserts no `transport_failed`/`phase=failed` (r8 `:90,:91,:96,:226`). **CONFIRM.**

**Leg-1 verdict: CONFIRM the m-8 r8 cancellation forward mapping.** The two-cut composition above is the design content I fold into r6 (it discharges R5-F1 and pins the cancellation identity that R5-F4's `partial_disposition` closed domain encodes — pre-transport ⇒ nothing-committed; post-invocation ⇒ partials-labeled + `unknown_effect`).

## Leg-2 — the E0-phase decision (mine as E0 populator): BRANCH (b), route an m-3 owner-delta

m-3 r3's `m3.app_event.v1` phase enum is `{denied, sent, completed, failed, unknown}` (`master/domains/m-3-.../design/2026-07-16-step3-mvp-egress-e0-e3.md:131`), and `unknown` is expressly reserved for genuine indeterminacy — it "mirrors m-10's `UNKNOWN_PROVIDER_OUTCOME` park … rather than forcing a fabricated outcome" (`:144`). A cancellation is **determinate** (I know the attempt was cancelled) and **not a failure**. Therefore:

- **(a) map to `failed` — REJECTED.** E0 is the live evidence rung the E3 evaluator adjudicates against; a knowingly-false `phase=failed` plants precisely the confusion/mislabel the frank threat model exists to kill (confusion-not-malice; mechanism = honest labeling). "Non-trust-bearing / E0" (m-3 `:146,:147`) means not conductor-evidence-grade — it does **not** license a false claim. m-8 also asserts no `phase=failed` for cancellation, so (a) would put me in direct contradiction with the owner seam.
- **misusing `unknown` — also REJECTED.** Cancellation is determinate; `unknown` is the indeterminacy/park token — routing a known cancel through it fabricates indeterminacy, a different lie.
- **(b) request an m-3 phase token — CHOSEN.**

**m-3 owner-delta ASK (for master to route; I do NOT author m-3's enum):** E0 has no honest phase for a determinate, non-failure cancellation. **Requirement:** a `m3.app_event.v1` phase value that represents a cancelled/interrupted attempt distinctly from `failed` (an actual fault) and from `unknown` (genuine indeterminacy). **Rationale:** honest-labeling; a cancel is a determinate non-failure terminal whose authoritative fact lives in my `turn_terminal{turn_cancelled}` + m-10's `CANCELLED` row. **m-3 owns the token's name/shape + the r3 re-hash + scoped rebinds** — I state only the requirement and the rationale.

**Interim posture (so E0 is never dishonest in the gap before m-3's token):** for a cancelled attempt the authoritative terminal lives in durable state (my `turn_terminal{turn_cancelled}` + m-10 `CANCELLED`), NOT in a fabricated E0 phase. Until the m-3 token exists I emit **no** terminal E0 `app_event` phase for a cancelled attempt (no `phase=failed`, no misused `phase=unknown`); the E0 stream carries the attempt's truthful non-terminal phases and its cancellation terminal is deferred to durable state, landing in E0 once the m-3 token lands. Fail-closed-to-honest: no false claim planted.

## Timing

Per your option and my r5 review cycle still being open (r5 `452a352d…` = must-revise, HELD):
- **Leg-1** (the two-cut composition + BOTH-count-§2a) pins **in-cycle** — it folds into the single batched **r6** (discharging R5-F1, feeding R5-F4's `partial_disposition` domain), against m-8's **final pair-approved** bytes (r8 is still proposed-pending its final-byte review).
- **Leg-2** — the m-3 ask returns now for routing; the E0-phase **realization** folds into the half **after** m-3's token bytes exist, and **does NOT gate the half's closure** (your explicit instruction). If m-3's token lands before r6, it rides r6; otherwise it is a scoped post-closure fold.

Dependency chain unchanged and surfaced: **m-8 r8 → its final-byte review (waits on this + the m-10 disposition) → m-8 pair-approved → my r6** (folds leg-1 + R5-F1..F4) → m-10 D-5 comparator confirmation. The m-3 phase token is a parallel branch that closes leg-2 without gating closure.

Duplicate/already-built gate: not applicable — a byte-bound seam confirmation + one owner-routed ask; no artifact beyond this relay.
Boundary contract: not applicable — no doc edit here; the r6 fold is the separate batched act in `step3-mvp-lifecycle-m9`.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no design-doc edit (r5 held), no frank/ edit, no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (s11-close)
Next requested action: master records leg-1 CONFIRM (m-8's fresh final-byte review may proceed on this + the m-10 disposition) and ROUTES the leg-2 m-3 owner-delta phase-token ask to m-3 (batched against any other m-3 item); m-9 folds leg-1 + R5-F1..F4 into the single r6 after m-8 pair-approves, and lands the E0-phase realization after m-3's token without gating the half's closure.
