# S6-AMENDMENT-SET-2026-07-06 — the integrated transport-fix design (assembled for VP co-sign)

**Status:** **r3 — VP CO-SIGNED (`s6-design/RECONCILE-orchestrator-reviewer-20260706-220325`): this set IS the
design-of-record amendment.** The §C4 pointer is landed (`ARCHITECTURE.md`); the s6 build slice dispatches against
this set (`s6-dispatch`, VP pre-handoff approved `…-221835`). Assembled by master from three **pair-complete** domain
halves + the grilled fork + m-4's surface certification. Lineage: r1 → **r2** (the persisted activation-marker
REMOVED per the VP's cross-domain catch `…-211542`; m-7's bounded confirm landed **derived-only on the merits** —
planner `…-212756` + implementer co-sign `…-213621` withdrawing its own prior integration condition; the GC probe
verified twice, locked §10 + live `gc.go`, no failure case; nothing routed to m-1; the m-7 doc at r5) → **r3**
(byte-fidelity per `…-215052`: m-7 cited at r5/current, FX-B1g in the roll-up @18, the marker-row carriage deleted).
The co-sign's build watchpoints [VP-W1..W3] are folded in the build dispatch. Baseline: `frank/` `main @ 7e5c527`
(tag `s5-close`). *(Status block updated post-co-sign per the VP's non-blocking note in `…-221835` — next-touch cleanup, handoff not reopened.)*

## The set (the substance lives in these; this doc integrates, it does not restate)

| artifact | content |
|---|---|
| `domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md` §A–§F.1 **+ the c1-doc folds** (§0.f fold-log; §6 rewritten — branch-A PARENT + Sharpening-D hint-class + `routing_ref_honored`; §5/§10/§12 consistency; zero `HELD` markers survive) | the fork (folded) · F10 split · F17 waivers+retraction · F14 invariant · B-3 activation |
| `domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md` **r5 (current)** | A-1 stable-schema digest · A-2 idempotent-replay intake · A-3 live mint · A-4 lock runtime · D-1/D-2 in-slice · B-1 lifecycle/roster (r5 = the marker withdrawal, pair-co-signed) |
| `domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md` **+ §11 (B-2)** | the ONE canonical envelope codec · three-layer record_kind authorization · waiver typing · the boot form |
| `master/GRILL-LOCK-parenting-fork-2026-07-06.md` | branch A · fallback-never-bounce (empirical basis) · precision accepted · the m-4 gate — **now discharged** (`s6-design/DESIGN-planner-…-201500` confirm → m-1 fold `…-210259`) |
| m-4's certification `s6-design/DESIGN-planner-20260706-201500` | Q-C intact · R2 **strengthened** · the `routing_ref_honored` condition (folded verbatim) |

## The complete disposition map (the design-phase exit bar: no silent drops)

| finding | disposition | owner |
|---|---|---|
| F1 | worked-as-designed — no change (credit) | — |
| F2 | codec amendment (render offers only gate-accepted); largely **dissolves under branch A** (no parent field) | m-2 |
| F3 | in-slice D-2 — bounce/reply detail parity, one reply-shaping path | m-7 |
| F4 | **dead at the root** — no candidate set exists under branch A | m-1 (§A) |
| F5 | A-1 stable-schema digest (volatile affordances digest-exempt by class) + `recipient_picker` DigestExempt | m-7 + m-2 |
| F6 | the single decode — raw comma-split + silent-drop **deleted** | m-2 |
| F7 | the reviewer-visibility gate decodes via the codec — one encoding satisfies typing + visibility | m-2 |
| F8 | reserved — the recorded numbering gap | — |
| F9 | A-2 — content-hash dedupe = full outcome REPLAY; durable monotonic ids; the §C4.1 1:1 anchor restored | m-7 |
| F10 | §B split — in-slice accepted-graph anchor fix + the `project()` default-accepted amendment (audit view = a parameter, never a fourth verb) | m-1 |
| F11 | **dead at the root** (branch A) + the dogfood-replay regression fixture | m-1 (§A) |
| F12 | waiver typing rows (`rationale`/`waiver_scope`/`retracts`; the `"*"` header retired) bound to §C | m-2 ↔ m-1 |
| F13 | three-layer record_kind authorization (membership / seat-scope / per-kind requireds) | m-2 |
| F14 | §D **I1-P invariant** (canonical root identity; claim-before-phase-0; proof-of-death takeover) + A-4 runtime enforcement — dual-cited per [VP-W1/W2] | m-1 + m-7 |
| F15 | A-3 — mint becomes an operator-channel loop mutation (`seat_mint` pivot); no restart, no fleet bounce | m-7 |
| F16 | in-slice D-1 — shim one-transparent-reconnect-then-retry | m-7 |
| F17 | §C — scoped waivers + the `waiver_retraction` record class; commit-order effective state; operator-only floor | m-1 |
| latency | **exonerated by audit** (ledger addendum) — NO engine perf work in the build | — |
| B-1 | seat lifecycle `minted→bound→active` from existing records + runtime `bound_now`; the literal pre-active admission allowlist (B-1.2a, the pair-caught classifier fix); **activation classification DERIVED-ONLY** — B-1.2b's persisted marker REMOVED per m-1's §F boundary (the VP catch, `211542`): accept-time classification via the admission allowlist, recovery re-derives by m-1's first-accepted-per-mint-generation rule; *m-7 pair CONFIRMED derived-only, `…-213935` (doc r5)*; the roster view on `project` | m-7 |
| B-2 | the typed boot form: `PHASE: SITREP` + lifecycle-gated — **no new phase/record_kind/§J2/atom/verb**; required set `charter_loaded` + `dispatch_status`; identity = the channel stamp | m-2 |
| B-3 | the accepted boot = the identity-activation edge; session restart = re-`bound`, never re-boot; liveness bookkeeping, not an identity upgrade | m-1 |

## The named seams (integration-checked; each doc states its half)

1. **Parenting (three-way):** m-1 defines the stamp semantics · m-7 maintains per-seat turn context + the commit-time stamp locus · m-2 drops the parent field from every form + shapes `parent_hint`/`parent_hint_honored`/`parent_provenance`.
2. **F14:** m-1 invariant / m-7 enforcement — the [VP-W2] dual citation present in both docs.
3. **F12↔F17:** m-1 record class + gate semantics / m-2 row typings — neither locks alone; both now pair-locked.
4. **Envelope.To + digest:** m-2 projection rule (header list = recipient truth) + DigestExempt render / m-7 the commit-loop stamp + A-1.
5. **Sharpening-D:** m-1 §6 carriage = m-4 §5:216 `lineage_role: routing_ref` — **the mapping confirmed identically from both sides**; the unverified-citation flag folded verbatim.
6. **Boot (ONE owner-consistent activation model — r2, per the VP review):** **m-1 §F OWNS activation semantics** — `active` is **derived**: the first accepted governed submit per mint-generation; **no persisted activation marker, no new system field, no new m-1 on-disk state** (m-1's boundary + its implementer's route-back trigger, both binding). Under B-1's own ordering rules a pre-active seat can only land a boot-form accept, so the first accepted record IS the boot by construction — activation is a pure fold over the store (canonical wins; no duplicated truth). **m-7 B-1** classifies at accept time via the admission allowlist (B-1.2a) and re-derives lifecycle at recovery by m-1's rule — B-1.2b restates as *transient runtime classification* — **CONFIRMED by the m-7 pair on the merits** (`…-212756` + `…-213621`): the GC probe found no failure case (GC collects `gc_marker` + drained intake segments only, never accepted canonical records — verified at locked §10 AND live `gc.go:46-102`); the one residual boundary (future retention changes touching accepted records) is recorded in B-1.2b under m-1's own verbatim trigger. **m-2 B-2** takes its own stated default shape: `SITREP` + the two boot fields as the form class, **no marker row**.

## Build-slice obligations (inherited red-first)

- **Registry pass (one, m-2-guided):** seven rows — `parent_hint` · `parent_hint_honored` · `parent_provenance` · `routing_ref_honored` · `rationale` · `waiver_scope` · `retracts` (**no activation-marker row** — r2); record classes `waiver_retraction` + `seat_mint`; the `ORCH_REVIEW_WAIVER` `"*"` header retired. **MINOR compat class throughout (m-2-ruled); no record-SHAPE change ⇒ no envelope migrator (R-1 held).**
- **Fixtures:** the GRILL_LOCK three (the archived-dogfood F11 replay without livelock · hint honored/fallback · concurrent-accept-no-parent-bounce) + m-7's **FX-A1a..FX-B1g (18** — r5 adds **FX-B1g, the re-mint/generation leg**: a re-minted seat shows `minted` not `active` for its new generation, pre-re-mint accepteds do NOT activate it, a fresh boot accept does — the executable proof of the derived-only model's mint-boundary behavior**)** + m-1's §E set (polluted-archive projection filter · reject-never-anchor · waiver scope/retraction/re-arm · the store-lock race/alias/kill-9 legs) + m-2's (codec round-trip · full-set-in-every-projection · three-layer record_kind incl. the genesis negative · waiver-as-record · boot renders-pre-active + un-bounceable · shared-vocab-untouched negative).
- **The step-exit test (operator-ruled, in-step):** the ROADMAP:83-85 legs on the fixed conductor + §7-applying s5's registry as the first live act + the F11 replay leg.
- **Ops (non-fixture):** pre-allowlist `mcp__frank__*` in seat sessions at relaunch.

## Constraints held (restated in every constituent doc)

Byte-exact `{accepted, rejected, held}` · the three-verb seat surface (roster + audit views are `project` parameters) · channel-stamped FROM · I-PH on every surface including the new ones · the tool-mediated confusion-resistance ceiling (no by-construction creep; D5 residual stands) · no Step-2 observe pre-work · the credit list (crash-atomicity, FROM-stamping, I-PH) untouched · **boot sequences, it never grants.**

## Process record

The [VP-W1..W3] handoff watchpoints held throughout (F14 never collapsed; the dual-cite present; the fork grilled before any m-1 lock). The boot addendum rode the operator's in-session request; the VP objection window passed without objection. Pair discipline caught real defects at every station this phase (the B-1 classifier hole; m-2's r1 self-contradictions; the m-1 §A packet sharpening) — six pair rounds, two must-revises, zero silent folds. **And the co-sign gate itself caught the r1 set's cross-domain conflict** (the activation-marker persistence contradicting m-1's derived-only boundary + undischarged route-back trigger — `211542`): the r2 revision restores the one-truth model. *(The VP also noted stale status prose in the m-7/m-1 amendment docs — non-blocking, later relays close those items; the pairs tidy at next touch.)*
