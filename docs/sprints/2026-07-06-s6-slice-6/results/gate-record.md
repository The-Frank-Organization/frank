# s6 Gate Record

This is the pair gate record for the s6 step-exit run. It records transport/provenance evidence only. Work done-state and `record_integrity` remain self-reported until Step-2 observe.

## Evidence Set

Run date: 2026-07-08.

Store: `frank-s6-step`.

Conductor: `s6-transport-impl@1f6cd08`.

Evidence directory: `~/frank-s6-gate/results/`.

Procedure of record: `docs/sprints/2026-07-06-s6-slice-6/results/step-exit-procedure.md`.

Master evidence set and caveats, carried verbatim from `master/relays/s6-gateday/SITREP-orchestrator-planner-20260708-020233.md`:

> **The run (2026-07-08, store `frank-s6-step`, conductor from `s6-transport-impl@1f6cd08`, procedure of record followed; evidence files in `~/frank-s6-gate/results/`, division per the standing rule — operator-authorship legs operator-run, mechanical legs master-driven with that caveat stated):**
>
> - **(b) The first live act — the operator §7 apply: ACCEPTED** (`relay-4ba7f4c0…`, `config_change`/`fieldspec`, `new_digest` = the composite chain digest, master-verified byte-identical to the genesis pin; body byte-identical to the pinned member; **the adapted [VP-W7] is discharged**). Two operator-seat HOLDS preceded it — both correct procedure (the digest-canonicalization question; the missing-context question) — see findings 2 and 1.
> - **(a) ROADMAP:83-85: PASS ×3** — accepted-only-through-the-conductor (submit→project→read round-trip) · FROM system-stamped (a deliberate `FROM: s6-forged.seat` payload forgery committed with envelope `from: operator` — the forgery non-authoritative; see finding 3) · validation-pre-delivery (`SUBJECT:required` typed reject, nothing delivered) + **a gate produced a local outbox item** (`outbox/gate-relay-0175…`, verified store-side). **A-1 held throughout: four submits + a landed §7 record, zero digest rotation.**
> - **(c) The F11 redrive — the closing bell: PASS.** All 14 archived accepted s5-dogfood records re-driven in intake order through a live minted probe seat: **14/14 accepted · parent-class bounces 0 · same-context re-renders 0** · 3 context-switch refreshes (the designed one-bounce-per-new-(phase,tier) dance, each resolved by one shim refresh + retry) · 9 archived seat-picked parents carried as `parent_hint` — including four `park-relay-…` cross-namespace ids straight from the bad-picks audit — zero bounces. Normalization stated in the evidence (system-stamped surfaces + recipients to the live store; retired `ORCH_REVIEW_WAIVER` dropped; archived parents → hints). **The traffic that livelocked s5 lands whole on the fixed conductor.**
> - **(d) The live boot walk: PASS.** Operator-channel `seat_mint` (`relay-272654cf…`) on the SERVING conductor — **no restart, no channel drops (A-3 live)** → a real fresh session wired → the pre-active schema rendered the boot form → boot accepted (`relay-834fecb4…`) → **roster: `minted → bound_now → active`, activation ref = the boot relay — DERIVED-ONLY, no marker anywhere.** Incidental live confirmations en route: an active seat's boot-shaped resubmit judged as ordinary traffic with full-form requireds in force (FX-B1f's second half) + D-2 inline reply detail (`"detail": "EVIDENCE_TARGET:required"`).
>
> **Stated caveats (the s4 honesty pattern):** legs (a)/(c) and the setup were master-scaffolding-driven (the operator authored the §7, both mints, and all designations); redrive run-1's 12 rejects sit on the store as evidence of the pre-dance client (the run-2 JSON is the artifact of record); the redrive's honored/fallback split is a crude client-side classification — the full replies are in the JSON for your precise classification; the walk seat's `project → null` is an EMPTY INBOX (active seat, nothing addressed to it), not a pre-active artifact. All claims transport/provenance-only; done-state + `record_integrity` remain `self_reported`.

## Finding Dispositions

F-GATE-s6-1: bounce-class UX is materialized as `OI-S6-BOUNCE-CLASS-UX.md`. Requiring `PHASE` / `CEREMONY_TIER`, or folding absent context into typed-required, is outside the co-signed set and routes to m-2 before any bounded build touch.

F-GATE-s6-2: digest wording is folded into the step-exit procedure. `new_digest` means the composite pinned-config chain digest over all members, not the standalone member hash. The operator-refusal-until-shown episode is credited as custody discipline working.

F-GATE-s6-3: envelope-key hygiene is materialized as `OI-S6-ENVELOPE-KEY-HYGIENE.md`. Envelope-layer `system_only` keys persisted as inert strings, while identity stamping held; reject-or-scrub shape routes to the next m-1 / m-2 seam touch.

F-GATE-s6-4: hosted-seat caveat is folded into `docs/ops.md` and the step-exit procedure. Hosted seats do not consume `tools/list_changed`; after any re-render bounce, re-read the schema instead of trusting a cached constant.

## Redrive Classification

Precise redrive classification, carried verbatim from `s6-gateday/REVIEW-FOLD-planner-20260708-022023.md`:

> Reply-level classification of `step-exit-dogfood-redrive.json` (run 2, the artifact of record), read from the committed records at the step store: **accepted 14/14 · parent-class bounces 0 · hint-carrying 9 → honored 0, fallback 9 · no-hint 5 → engine default**. Master's client-side `hint_honored: 9` was the crude split its own caveat flagged: every hint-carrying record actually carries `parent_hint_honored: no` with the verbatim hint preserved (intent + outcome + flag all present on each record — the GRILL_LOCK triple). **Zero honored is correct engine behavior under the redrive's normalization, not a prover deficiency:** the hints were ARCHIVED s5 relay-ids, and the re-driven records carry fresh ids on the fresh step store, so no archived id can exist in that store's accepted graph — every hint is honestly unprovable, and fallback-never-bounce carried all 14 to acceptance (including the four cross-namespace `park-relay-…` hints from the bad-picks audit). On a live store where hints cite local ids the prover resolves them (the G-2 hint-honored fixture is the executable proof). Stamped provenance: 13× `woken_on`; 1 record (re-drive of `relay-3728794d…`) committed with NO parent edge and no provenance — the probe seat's turn context was empty at that instant, the firstDefined chain exhausted, and anchoring stayed live (accepted, no bounce; class gates demanded nothing of a SITREP) — the transport-is-always-live property at its boundary. **The F11 claim does not depend on the split: 0 parent-class is the claim-bearing number.**

## Honesty Rail

This gate record claims transport and provenance behavior only. It does not claim semantic completion of the work described inside relays. Done-state and `record_integrity` remain self-reported until Step-2 observe.

Credentials, markers, and custody-only binding details are not copied into this record. The live-store zero-hit sweep for marker / activation / `realized_mint_ref` / credential-material patterns is recorded in the orchestrator disposition relay.
