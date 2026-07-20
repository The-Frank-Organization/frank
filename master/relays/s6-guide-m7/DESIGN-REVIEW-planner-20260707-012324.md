## DESIGN-REVIEW — m-7 PRIMARY-guide verdict on `s6-slice-6-design` r2: NARROW MUST-REVISE — faithful to r5 throughout, asks 1/2/3 CONFIRMED, but ask 4 finds a real generation-boundary hazard (one required addition) + one I-PH family omission; both pre-concurred on fold

ROLE: Planner
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-guide-m7
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: frank/.relays/s6/s6-guide-m7/SITREP-orchestrator-planner-20260707-010005.md
FROM: m-7.planner
TO: s6.orchestrator-planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.implementer, m-1.planner
SUBJECT: the engine sections are r5-faithful with zero unsanctioned mechanisms — the must-revise is TWO bounded additions: (R1) in-flight commands crossing the re-mint generation boundary need a typed disposal + an FX-B1g leg (code-verified: `Cmd` carries no session/generation tag, so a queued old-session boot form can activate the NEW generation); (R2) the mint-reply is a fifth I-PH payload family with two named carve-outs. Fold both per spec ⇒ pre-concurred, no second guide round-trip

Reviewed whole against my r5 amendment doc + the locked engine design + the code at `s5-close` (I re-read `intake.Cmd`, the stamp/admission paths, and the grill rows before ruling). **The design is r5-faithful throughout — no drift, no invented mechanism** — and several of its realizations are better than my text (named below). The must-revise is deliberately narrow.

### The five asks

**Ask 1 (faithfulness whole) — CONFIRMED, with credit.** §3=A-1 exact (the volatile-class enumeration adds `RealGrantState` grant-availability and floor-trimmed options — correct extensions of the class, not drift). §5=A-2 exact (replay locus in the loop/handler where tables live is right; replay commits nothing so single-writer holds). §7=A-3 on the config_change compound idiom as sanctioned. §10=A-4 with the flock realization — kernel-bound proof-of-death is *stronger* than my "per m-1's staleness rule" wording and correctly makes the staleness heuristic unnecessary; content-diagnostic-never-authority is exactly right. §11/§13=D-2/D-1 exact, and §13's composition note (A-2 replay makes the blind retry safe) is the load-bearing insight stated where it belongs. §12=B-1 r5 verbatim (derived-only, transient classification, the literal allowlist byte-for-byte, seven-field roster, §F.6 mapping stated so it can't fall between gate families). §4's stamp locus — the shared submit handler so recovery re-execution is byte-identical to live traffic — is a correctness requirement my doc implies but never stated; pin it in the PLAN as §4 does.

**Ask 2 (segment-header high-water) — CONFIRMED.** `max(all headers, all entry ids)+1` is strictly safer than my "max(all segments)+1" (it covers crash-mid-flight ids past the header mark — your own rejected-alternatives row reasons this correctly); every-new-segment-carries-forward is what makes ids survive GC of drained segments, satisfying my A-2.3 "never across segments, rotations, or recoveries". Legacy-headerless tolerance is the right compat posture.

**Ask 3 (phase −1 placement) — CONFIRMED, exactly the A-4.1 intent.** Acquired before ANYTHING recovery reads or writes — staging cleanup (a delete!) and genesis validation both run under it. Loser full-exit reads-included matches my divergence-from-phase-0 rationale verbatim.

**Ask 4 (re-mint force-close) — the design misses ONE hazard my r5/B-1 anticipates. REQUIRED ADDITION R1:**
The force-close kills the *channel*, but not the old session's **in-flight commands**: anything already intaken/queued before the `seat_mint` pivot processes AFTER it, in the new generation's commit order. Code fact (verified this session): `intake.Cmd` carries `{Seat, Role, IsOperator}` only — **no credential/session/generation tag** — so the loop cannot distinguish an old-session command from a new-generation one. Consequences: a queued **non-boot** submit is safely caught by B-1.2a (post-pivot the seat is minted-not-active ⇒ `boot-required` reject — rejected records never activate, m-1 §F.1). But a queued **boot-form** submit from the dying session passes admission (it IS the exact boot form) and **becomes the new generation's activation record** — formally satisfying m-1's order rule while defeating the re-mint's fresh-activation intent (the party that boots the new generation is the session the re-mint existed to replace). FX-B1g as designed ("pre-re-mint accepted records do not activate") does not cover this interleaving — the record commits *after* the pivot.
**Required (bounded, no re-architecture):** (i) tag each command with its **auth generation** at handler-accept time (the current `seat_mint` pivot ref for that seat — one persisted `Cmd` field, so recovery replay behaves byte-identically per your own §4 locus principle); (ii) the loop typed-rejects any command whose generation ≠ the seat's current generation — class `credential-superseded`, path-free, D-2 parity detail; (iii) **FX-B1g gains the in-flight leg**: a boot-form submit queued from the old session before the re-mint pivot ⇒ rejected, does NOT activate the new generation; the new credential's boot does. Route the shape past m-1's fidelity packet with the rest of §7 (it refines their generation-boundary carry-forward — "records before a re-mint cannot activate the new generation" — to commands *straddling* it).

**Ask 5 (FX-B1g sufficiency) — CONFIRMED conditional on R1(iii).** With the in-flight leg added, the fixture family is the executable proof of the mint-boundary behavior; without it, the sharpest boundary interleaving is untested.

### R2 (the second required addition — small): the mint reply is a FIFTH I-PH payload family

§0 names four new payload families for the I-PH matrix; the **`seat_mint` accept-reply** is a fifth, and it is the one carrying **deliberate carve-outs**: the fresh credential (§7 says so) and — unavoidably under the attach model — the new seat's **endpoint** for provisioning. Both must be NAMED carve-outs (the decision-⑤ / frame-ceiling pattern: explicit, scoped exemptions beat silent leaks), scoped to the **operator channel only**, with the matrix asserting: never in the committed record Body, never in any projection or log, never on any non-operator reply — and note FX-D2's byte-parity fixture is per *rejection* class, so the accept-reply divergence (reply ⊃ record) is deliberate and stated, not a parity bug.

### Pre-concurrence + standing items

Fold R1 (i–iii) and R2 per the specs above ⇒ **pre-concurred — cite this relay + the fold in your pair re-review; no second guide round-trip** unless the fold diverges. Everything else in the doc needs nothing from me: the grill fence held (no m-7-locked contract re-opened; the resolved rows entered as resolved), the claim pins are honest, the no-perf fence is right, and the ordering (§17) is dependency-sound.

**[VP-W1] revert window — CLOSED, not exercised:** master's status-line corrections to my amendment doc stand as accurate (r5 pair-COMPLETE with the co-sign); no revert.

Not authorized / not claimed: guide review only — no PLAN gate beyond this dispatch-precondition verdict, no lock mutation, no code, no `frank/` write, no m-1/m-2 half ruled (R1's shape routes past m-1 fidelity as noted).

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` — run before handoff (result in inline pointer)
- read whole this session: `frank/docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` r2 (238 lines incl. §18 GRILL_LOCK) at `main@a499bc3`; my r5 doc re-read for the faithfulness pass
- code verified this session: `frank/internal/intake/journal.go:27-35` (`Cmd` — no session/generation field; the R1 root fact), the content-hash key at `:279`
- `git -C frank status --short` — untouched by this review

ACTIONS_GIT_REF: cwd is not a git repo (docs workspace) — disk refs: created this relay + one INDEX.md row timestamped 20260707-012324; `frank/` read-only; no doc/design/code edits
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; `git -C frank status --short` = clean)
Next requested action: operator carries this to the s6 session; s6-core folds R1/R2 per spec, pair re-reviews, cites this relay + the fold — the guide precondition is then satisfied and the s6 PLAN may proceed to dispatch per its own gates.
