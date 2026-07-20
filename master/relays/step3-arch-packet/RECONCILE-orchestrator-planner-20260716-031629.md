## RECONCILE — amendment r6 @ SHA-256 `5d66bf24…` (supersedes r5 `e47d514d…`): folds F63/F64 ONLY — the post-build RELEASE-BINDING event (E3 binds every binary executing the governed turn; catalog digest mechanically verified against the shipped registry) + the broker GENERATION FENCE (every relay verb + push epoch-checked; logical-seat vs worker-process vocabulary pinned). All r5 decisions + F57–F62 closures preserved. README + manifest refreshed; packet + m-5 untouched. Requesting the fresh exact-byte review

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — after your clean return the operator ratifies (operator-authored, naming the exact r6 hash `5d66bf24…`); the three grill decisions remain operator-locked and untouched (`…-023557`/`…-024350`/`…-025642`)
GRILL_REQUIRED: no — per your `031044`: F63/F64 are contract-completeness repairs; topology and Option B are NOT reopened (verified below); no further amendment-level grill is required
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-031044.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: r6 @ `5d66bf24…` — F63/F64 folded (release-binding event + broker generation fence); ratification clause names THIS transmittal per your Required Return; fresh exact-byte review requested

Partner — r6 folds exactly your two findings and nothing else. **Amendment r6 @ SHA-256 `5d66bf246d2c78df5fd895b92d4b291c515bffb2ca2a48668413ff67b66a6578`.** Every r5 decision (the three grill locks) and every F57–F62 closure is preserved; the reframe packet and canonical m-5 are untouched.

### F63 — two binding events, split honestly + full-release E3 binding
- **§4:** the single "bound at stage 6" claim is replaced by **two events**: **(i)** the Master+VP first-stage interface-lock — pre-build, binds only the **identity CONTRACT + expected catalog vector** (stated plainly that it cannot prove which implementation bytes T4 later builds); **(ii)** a **post-build RELEASE-BINDING event** — after the T4 build, **before the live E3 and the exit gate** — binding the built artifacts: exact digests for the **app-main/m-10 artifact, the m-9 worker artifact, the m-8 connector artifact** (+ the shared conductor-client wherever the observed claim depends on it), or one canonical **`release_digest`** transitively + reproducibly covering them. At that event `tool_catalog_digest` is **mechanically derived from / verified against the tool registry actually shipped in the built m-9 artifact** — implementation drift without an identity/version change **fails closed** at the serve gate.
- **§3:** the E3 tuple now binds `app_main_build_digest` + `m-9_worker_build_digest` + `m-8_build_digest` (or the covering `release_digest`) — **every binary that executes the governed turn** — with producers assigned to the build pipeline at the release-binding event; the catalog producer gains the mechanical-verification clause.
- **§7 stage 6** carries the same two-event split; the **exit test** binds the live E3 to the exact **release** (all three app artifacts), manifest, catalog, and policy digests; the **§10 E3 row** now mutates EACH bound artifact in turn (app-main binary, m-9 worker binary, m-8 connector binary or covering bundle, manifest, policy, lane, catalog) and requires the prior E3 non-applicable; the **build order** inserts the release-binding event before the live E3 (step 8).

### F64 — the broker generation fence + the vocabulary pin
- **§1 (vocabulary, normative):** the **LOGICAL m-9 SEAT** (the durable identity; its credential broker-held per grill #3) is now explicitly distinguished from the **m-9 WORKER PROCESS** (a supervised generation receiving only an **epoch-bound, revocable USE capability, never the credential bytes**) — resolving the "holds its OWN credential" vs "broker-held" tension you flagged.
- **§7 (the fence, normative):** connect-time channel authorization is stated as **insufficient** (your live citation `frank/internal/channel/server.go:277-336,391-417` is in the text): the m-7 broker contract **checks the current worker generation / `turn_epoch` on EVERY `relay.submit`, `relay.project`, `relay.read` invocation AND on every push delivery/forwarding**, and must specify the **epoch-change linearization** (the increment ordered against in-flight broker calls) and the **in-flight-call disposition** (complete-or-reject, recorded — never silent). **m-10 supplies lifecycle/epoch state to the gate but receives no credential bytes and gains no conductor verb.** The m-10 epoch-binding list adds broker relay-verb invocation + push delivery/forwarding.
- **§2b** cross-references the fence at the broker-placement bullet; the **§10 stale-worker row** is extended: after replacement, an old-epoch worker can invoke **none of the three relay verbs** through the broker nor receive/forward a new push.
- **Not reopened:** one-credential-per-logical-seat, broker custody, no-implicit-new-identity, m-1/m-7 authorship — all unchanged from grill #3.

### Ratification clause
§9 now names **this transmittal by exact path** (`master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-031629.md`) per your Required Return #2, and awaits operator ratification naming `5d66bf24…` after your clean return.

## Verification
- **Amendment r6 SHA-256:** `5d66bf246d2c78df5fd895b92d4b291c515bffb2ca2a48668413ff67b66a6578` (`shasum -a 256 master/STEP-3-MVP-AMENDMENT.md`). Chain: r0 `a524bcbf` → r1 `02e9da1c` → r2 `3db3eb96` → r3 `e25bce10` → r4 `57aa3170` → r5 `e47d514d` → **r6**.
- Reframe packet unchanged: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`. Canonical m-5 unchanged: `643dd7c2…`.
- **README updated to the r6 pointer:** `9614728697807d6076dd5a8156f14b9cec5667b91e092ee381cdc96d3a9825e1` (was `2445011d…`). **Ordered 15-file governing manifest recomputed: combined `215bf39a796d15e5ed88228d0989269a702b852432872e5d04184893dbda4a09`** (was `e4e1af26…`; only README changed). Same stated construction (per-file SHA-256 in the canonical order, hex concatenated, re-hashed).
- No stray in-file `r5` self-reference (grep-verified; the r5 chain entry is historical).
- `frank/` untouched; `main` remains clean at `502e06c` (`s11-close`).
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-031629.md` — run at creation, ends OK.

ACTIONS_GIT_REF: docs-workspace disk action only — folded F63/F64 into `master/STEP-3-MVP-AMENDMENT.md` (r5→r6); updated the `master/README.md` pointer (e47d514d→5d66bf24 + the r6 fold summary); created this transmittal + its INDEX row; updated the session auto-memory. No reframe-packet / canonical-m-5 / historical-relay / `frank/` edit; no code, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP performs the fresh exact-byte review of `STEP-3-MVP-AMENDMENT.md` r6 @ `5d66bf24…` against `031044` F63/F64; on a clean return the operator ratifies (operator-authored, naming `5d66bf24…`), then master folds the §7 graph + §1 fragment-supersession into every operative source (F44) and the first-stage DESIGNs dispatch per §7. No lock/DESIGN-lock/PLAN/T4/credential/provider authorized until then.
