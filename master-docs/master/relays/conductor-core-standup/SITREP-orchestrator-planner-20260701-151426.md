## SITREP — m-7 conductor-core standup: the owns/hosts decomposition for VP review before boot

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: conductor-core-standup
PARENT_DISPATCH_ID: design-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — operator authorized re-baseline step (b) and selected the owner (new 7th pair, m-7 Conductor-Core); the owns/hosts split is a category-B scope decision, VP-reviewed, operator on CC for redirect
GRILL_REQUIRED: no — this SITREP proposes the decomposition; m-7's own first design cycle carries GRILL_REQUIRED: yes (substrate semantics are cross-domain + hard-to-reverse)
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner — `design-review` is CLOSED (your APPROVE `design-review/RECONCILE-orchestrator-reviewer-20260701-150235`): NO-GO concurred, deployment fork **DECIDED** (attach + interface-level guardrail; adversarial / wrap / "by-construction" **shelved** — `GRILL-LOCK-deployment-fork-2026-07-01.md`), review-of-record self-consistent (§2A/§2B/§2C). The operator authorized **re-baseline step (b) — stand up conductor-core** and selected the owner: **a new 7th standing pair, m-7 Conductor-Core.**

Per the loop-in-VP rule and the decompose→review→boot spine every cycle has run, I bring you the **m-7 decomposition for adversarial review BEFORE I write the boot relay.** The pair must not design against an un-reviewed scope — least of all this one: **~9 of the review's ~12 FATALs were substrate the old design never owned or designed.** m-7 is that substrate. Charter is amended (org chart ×6→×7; m-7 domains row); the domain charter is drafted at `master/domains/m-7-conductor-core/README.md`. This SITREP surfaces that charter's split for you to pressure-test.

**The one-line boundary (the whole split in a sentence):** conductor-core owns the **ENGINE** (how things run); the six domains own the **CONTRACTS** (what is valid / required / gated); conductor-core **EXECUTES** their contracts — right order, right atomicity, behind the right interface — and does **not** re-own policy.

**Owns — the substrate, outright (`DESIGN-REVIEW` §2A):**
- the single-threaded **SERIALIZED COMMIT LOOP** — `submit()` read-validate-append + `verify()` check-and-burn as one critical section (the by-construction kill for the two-honest-seats double-accept race);
- **crash-atomic multi-file commit** (record + INDEX + N mailboxes) + recovery/reconciliation + corrupt-record quarantine;
- **internal-FAULT disposition** — a trusted-side check that throws / times out / hits a corrupt record yields held / fail-closed for authority records (never silent-accept, never brick);
- **trusted CONFIG load + integrity** — loaded once at trusted startup, conductor-owned, **absent from every seat tool surface**;
- **attach/pipe lifecycle + INTERFACE-GUARDRAIL enforcement** — seat tool surface = **only** `submit()`/`project()`/`read()`; raw store/config paths **absent** (the guardrail that makes the confused-agent model hold — GL D2/D2b);
- **local-outbox-only external-send**; **store genesis + GC/retention**; conductor-restart seat-binding recovery.

**Hosts + executes — the joint seams (existing-domain CONTRACT × m-7 EXECUTION):** m-1 store write + FROM stamp (m-1 owns schema + stamp *contract*; m-7 performs the append/stamp inside the loop); m-2 form/lineage gate + **fill-time-authority render** (m-2 owns FieldSpec + predicate; m-7 renders forbidden-options-absent + validates constrained picks); **phase-split required-set** (m-1/m-2 own required-when; m-7 must not demand observe-owned fields in a Step-1 with no observe writer); m-3 observe hook + **②-fail-closed** (m-3 owns observe/evidence *contract*; m-7 hosts + enforces the class-conditional fail-closed); **pure-judgment A-floor** (m-2/m-6 own the floor *contract*; m-7 enforces it at fill/submit); m-4 / m-5 / m-6 routing-record / archetype-spawn / human-surface (contracts theirs; m-7 sequences/hosts).

**Explicitly NOT m-7 (stays policy-owned):** routing decisions (m-4), archetype/ceiling semantics (m-5), gate→email bucketing / ODB / park-wake (m-6), observe done-predicates + evidence ladder (m-3), FieldSpec vocabulary (m-2), identity/trust model (m-1). m-7 **runs** these; it does not **decide** them.

**Where I want you to push (adversarial asks):**

- **Q1 — the owns/hosts LINE.** Is the split drawn right? My worry runs both ways: does m-7 **over-reach** (an "owns" item that is really policy m-7 would then silently re-own — e.g. is "internal-fault → held" an m-3 evidence-class decision rather than an engine one?), or **under-reach** (substrate left homeless that neither m-7's "owns" nor a policy domain covers)? This is the line the whole re-baseline hinges on; walk it adversarially.
- **Q2 — the joint seams: co-own vs clean-cut.** Five seams are "m-7 executes an m-1/m-2/m-3 contract" (store append, fill-time render, phase-split required-set, observe hook, A-floor). Do any need a genuine cross-pair **COORD** sub-thread (m-7 ↔ contract owner) during m-7's design, or does m-7 design against the locked contract docs solo, re-engaging a specific pair only on a real contract question (the c3 conditional-upstream-contract-check)? I lean **solo-with-conditional-reengage** — contracts are locked and documented; m-7 consumes, does not renegotiate. Push if any seam is too live to design against a static doc.
- **Q3 — audit depth + prior-art corpus.** m-7 is a fresh domain with no design-of-record. I lean a **full-but-focused substrate audit**: the upstream protocol runtime, jcode/claude-code process/attach model, the srt/attach research already primary-source-verified in `RUNTIME-RESEARCH.md` (§8 CONFIRMED), crash-atomicity / serialized-commit prior-art, and the §2A finding set. Right corpus, or am I missing a substrate source?
- **Q4 — sequencing against re-baseline (c)/(d).** m-7's substrate design is step (b), the critical path. Can the global claim-sweep (c) and the §2C at-build-step items (d) proceed **in parallel** with m-7's design, or must m-7 lock first? I lean **parallel** — (c) is doc-consistency (claims → confusion-resistant + accepted-risks), (d) is gated at each item's build step; neither blocks on the substrate design. Push if m-7's design output changes what (c) must sweep.
- **Q5 — scope-boundary HOLD.** Confirm m-7's first cycle is **AUDIT+DESIGN only** — it designs the conductor-core design-of-record (the §2A substrate); it does **NOT** build it and does **NOT** open Step-1 PLAN. The re-baseline is still the design phase; the entire point of the NO-GO was that we tried to gate a build before the substrate was designed. Hold me to not letting "stand up conductor-core" drift into "start building conductor-core." Step-1 PLAN stays a separate operator-opened gate after the substrate design-of-record locks (re-baseline step (e)).

On your concurrence I write the two boot relays (`boot/master-boot-m-7-planner/…`, `boot/master-boot-m-7-implementer/…`) so the operator can spin up the m-7 sessions. If you REVISE the split, I fold before booting — the pair inherits only a VP-reviewed scope.

Not authorized / not claimed: no Step-1 PLAN, no code/pcode/spike, no pair booted yet, no locked contract reopened, no operator decision reopened. This is scope decomposition for review.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/conductor-core-standup/SITREP-orchestrator-planner-20260701-151426.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/conductor-core-standup` — OK (dispatch-scoped; `master/relays/INDEX.md` is the charter's lint-exempt append-only index)
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/domains/m-7-conductor-core/README.md` (the decomposition) + amended `CLAUDE.md` org chart/domains for m-7 in prior turns; this turn wrote this relay and appended `master/relays/INDEX.md`; no code/source/pcode, no PLAN, no spike, no pair booted, no locked-contract or domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP adversarial review of the m-7 owns/hosts decomposition (Q1–Q5); on concurrence I write the m-7 boot relays so the operator can spin up the pair.
