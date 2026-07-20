# GRILL_LOCK — Step-1 deployment fork (attach vs wrap)

Durable design sub-artifact from the `design-grill` on the deployment fork — the #1 item of the re-baseline the
adversarial review (`DESIGN-REVIEW-2026-07-01.md`, VP-concurred NO-GO) demanded be decided first. Folds into the
re-baseline's eventual `DESIGN_LOCK_ID`. Not a mechanical gate; a human-alignment record downstream agents inherit.

```text
GRILL_LOCK_ID: GL-deployment-fork-20260701
GRILL_REQUIRED: yes  (operator asked; hard-to-reverse; ~9 of 12 review FATALs hang on this one choice)

GRILL_SOURCE:
- plan/design/audit relay read: DESIGN-REVIEW-2026-07-01.md; RUNTIME-RESEARCH §7/§8/§10/§14; READINESS-REGISTER
  Cluster 2 + decision ①; design-review VP concurrence (…-141436).
- code/docs inspected: m-1 design §4:75 (DI-1 realized via separate-uid/0700/no-path-handle), §6/§10:60 (DI-2 =
  per-seat isolated connection), §11:171/179 (DI-2/DI-5 realization DEFERRED to PLAN; §12 AC "tested not asserted");
  RUNTIME-RESEARCH §14 (srt sandbox foolable; Ona agent-self-disable incident).
- questions answered from codebase:
  · "Can pure attach give DI-1 (store/config write-exclusion) without launching seats?" → NO. Same-uid lane + local
    append-file store ⇒ lane can `echo >> store`. The three isolation options (separate/restricted uid; OS jail
    confining the lane FS; store-as-credentialed-service) all require the conductor to control the seat's execution
    env or the store's access path — i.e., not pure ride-the-session attach.
  · "Does DI-2 (identity / channel-stamped FROM) survive attach?" → YES (each seat is its own process with a private
    pipe). The wrap-vs-attach split is specifically about DI-1 write-exclusion, NOT identity.
  · "Does wrapping break persistent seats / the wake story?" → NO. A conductor-launched seat is still a warm
    persistent pipe; launch ≠ ephemeral. (So wrap was never incompatible with persistence — the review's
    "tears down persistent seats" is about mid-sprint SWITCHING attach→launch, not about launch itself.)
- questions asked operator: Q1 Step-1 threat model; Q2 interface-guardrail + wrap-deferral; Q3 honest claim-set +
  accepted-risk framing.

Resolved decisions:
- D1  Step-1 threat model = HONEST-BUT-FALLIBLE / CONFUSED, not adversarial. — operator — confused/hallucinating
      agents are the dominant real failure mode; deliberate malicious agents are rare and a later concern. — source operator
- D2  Step-1 deployment = ATTACH (not wrap). — operator + code — wrap's DI-1 isolation only defends against an
      adversary going AROUND the gate (out of Step-1 scope); a confused agent fails THROUGH `submit()` and is caught
      by the gate, so wrap buys nothing Step-1 needs. — source operator+code
- D2b Step-1 guardrail = INTERFACE-LEVEL, not OS-level. Seats act only through the conductor's
      `submit()`/`project()`/`read()` tools; the store is not a seat filesystem handle; policy config is
      conductor-owned, loaded once at trusted startup, and NOT in any seat's tool surface. A confused agent can't
      confusedly do what it has no tool for — no uid isolation required. — source operator
- D3  Adversarial-agent containment / wrap / "by-construction" = SHELVED INDEFINITELY (research-gated). — operator +
      code — even frontier-lab containment is defeatable (our §14: srt foolable, agents self-disable their sandbox);
      chasing perfect containment now is security theater we cannot back. NOT a roadmap rung. — source operator+code
- D4  Honest claim-set (replaces every "by construction / structural / sole-writer / unbypassable / forgery-robust"):
      "Confusion-resistant — the gate catches honest agent error because seats act only through the conductor
      interface. A determined MALICIOUS agent with code execution can subvert the store/config/FROM-stamp and is
      EXPLICITLY OUT OF SCOPE for frank (an open research problem). Re-open trigger: adversarial-containment research
      matures, or the deployment threat model escalates." — source operator
- D5  Config-integrity, store-write, and operator-FROM surfaces = documented EXPLICIT ACCEPTED-RISKS with the D4
      re-open trigger — recorded as known decisions, not silent holes (honest-fallback discipline). — source operator

Rejected alternatives:
- WRAP-for-Step-1 (conductor launches seats under restricted uid/jail) — buys by-construction DI-1 the confused-threat
  model does not need; adds seat-launch complexity; adversarial defense is shelved anyway (D3).
- STORE-as-credentialed-service (re-architect the m-1 store from local append-files to a socket/service to get DI-1
  under attach) — a real m-1 redesign for a guarantee not required under the confused-threat model. Store stays local
  append-files.
- THREE-RUNG LADDER (attach → srt-wrap → by-construction, sold as bounded increments) — false gradient: rungs 2/3
  require launching seats (a re-architecture, not an increment), and the top rung is defeatable anyway (§14). There
  is no ladder to climb; do not sell one.

Still operator-owned:
- Re-open of the deployment fork if/when the threat model escalates to adversarial OR adversarial-containment research
  matures (the D4/D5 trigger). Cannot be inferred — it is a future risk-appetite call.

Design-lock impact (what the re-baseline inherits) — CORRECTED per VP `design-review/…-144217`. The distinction is:
adversarial SECURITY CLAIMS collapse; confused-agent INTERFACE MECHANISMS do NOT — they are conductor-core requirements.
- COLLAPSE to claim-sweep + accepted-risk (adversarial security CLAIMS ONLY): the "by-construction" guarantees against a
  MALICIOUS lane — adversarial DI-1/DI-2 bypass, malicious direct store-append, malicious operator-FROM forgery,
  malicious config edit, malicious external-send bypass, the three-rung ladder. Relabel to the D4 claim-set; record the
  D5 accepted-risks. This is a doc/claim change, NOT the removal of the mechanisms below.
- RETAIN IN CONDUCTOR-CORE Step-1 (confused-agent mechanisms + threat-independent correctness — all MUST-before-PLAN):
  (1) the **INTERFACE GUARDRAIL as a hard design requirement + acceptance fixture** — seats act ONLY through
  `submit()`/`project()`/`read()`; raw store/config paths are NOT in ANY seat tool surface (a confused seat following a
  bad instruction must have no *tool* to reach them — this is what makes config-lane-writable / store-write /
  sole-external-sender adversarial-only); config is conductor-owned, loaded at trusted startup; **local-outbox-only
  external-send posture**. (2) **FILL-TIME AUTHORITY / form rendering** — forbidden choices absent from the rendered
  form, courier validates constrained picks; only the by-construction *claim* collapses, the *mechanism* is the
  confusion-resistant interface contract and stays. (3) the running-program substrate — single-threaded SERIALIZED
  COMMIT + crash-atomic ordered multi-file commit + RECOVERY + internal-FAULT disposition + store genesis/GC. (4) the
  STEP-1/STEP-2 PHASE-SPLIT required-set fix. (5) confused-agent guardrails — PURE-JUDGMENT A-FLOOR by
  (phase × record_kind); DECISION-② fail-closed on authority-class `self_reported`.
- RETAIN BEFORE THEIR BUILD STEP (not a conductor-core-only-slice blocker UNLESS that slice ships the schema): R2 opt-in
  `gate_referenceable` + row-column FieldSpec — **m-2/m-4 schema work must NOT ship with R2 under-specified**;
  altitude-B PER-ROW deviation grain; away-token sibling-burn + restart-binding — **stale approval / sibling-token reuse
  bite NORMAL operator flows** (not malicious-only), so a hard gate before park/wake + the away-bridge, NOT disappeared;
  away egress/ODB collisions.
- STORE stays local append-files (D2). Register: decision ① reframed — attach DECIDED; standing claim
  "confusion-resistant + adversarial-out-of-scope"; "by-construction" shelved (D3).
```

**Feeds:** the re-baseline's `DESIGN_LOCK_ID` (future). **References:** `DESIGN-REVIEW-2026-07-01.md` §2/§5,
`READINESS-REGISTER.md` (decision ①), `RUNTIME-RESEARCH.md` §8/§14.
