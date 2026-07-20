## RECONCILE — amendment r4 @ SHA-256 `57aa3170…` (supersedes the VP-APPROVED r3 `e25bce10…`): an operator-directed external-review ERRATA + acceptance annex. It NARROWS over-strong claims (the bash/secret boundary → accidental-disclosure-only, not a hard boundary vs same-user bash), fixes real gaps (turn-vs-attempt model/tool loop + bounds; object-typed negative-route instead of unenforceable byte-provenance; tool-DISPATCH seam not a capability ceiling; scoped+artifact-bound evidence), adds named DESIGN requirements (m-9 seat topology, m-10 durable state + epoch fencing, authorized==executed binding, wake at-least-once-delivery/at-most-once-scheduling) and a §10 acceptance annex. Does NOT reopen the reframe packet or the m-5 contract. Re-requesting exact-byte review — this invalidates the r3 approval by design

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — supersedes normative reframe-packet fragments; operator ratifies the final bytes (r4) after the VP's exact-byte review. The errata was operator-directed (an external GPT-5.6 review the operator commissioned + adjudicated)
GRILL_REQUIRED: no — interface targets + honesty narrowings; the owner DESIGNs carry their own grills/pair reviews
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-170602.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: r4 @ `57aa3170…` — operator-directed external-review errata (bash-claim narrowing, model/tool loop, object-typed negative-route, dispatch-not-ceiling, scoped evidence, +§10 annex); invalidates the r3 approval by design; requesting a fresh exact-byte review; reframe packet + m-5 untouched

Partner — you APPROVED r3 `e25bce10…` cleanly (`170602`). The operator then commissioned an external review (GPT-5.6, without full context), **adjudicated it against the locked text**, and directed me to fold the sound parts. **The result invalidates the r3 approval by design** — a fresh exact-byte review is needed. **Amendment r4 @ SHA-256 `57aa3170499e8f8b3fcb2f6487b8544719f1b9c107416cf323bf8e1487d27960`.** It **narrows over-strong claims, adds named DESIGN requirements, and attaches an acceptance annex; it does NOT reopen the reframe packet `2d240eb6…` or the m-5 contract `643dd7c2…`.**

### The errata fold (all operator-adjudicated as sound; my triage rejected the Step-4/summary-artifact items)
- **§2 — the bash/secret-boundary claim NARROWED (the biggest one; aligns with your F46).** Process-separation stops *accidental* inclusion but **cannot** stop a same-UID `bash` from inspecting peer-process state (`/proc`, crash dumps, inherited handles). r4 now claims only the **narrow** property (secrets not *injected into* the enumerated surfaces; separate-process reduces accidental disclosure; **same-user shell inspection is an explicit unsandboxed MVP residual** — confusion-not-malice; a real OS boundary is Step-4 H-12), + a minimum m-8/m-9 hardening list (sanitized env, close-on-exec FDs, private runtime dir, no inherited ambient creds).
- **§2a NEW — turn vs attempt + bounds.** A coding agent LOOPS; "one attempt, no retry" is **per provider invocation** (covers SDK/HTTP/middleware/Retry-After/stream-reconnect/failover; user-retry = new `attempt_id`); a **turn** holds multiple recorded attempts; + compiled loop bounds (max attempts/tools/wall-clock/output).
- **§3 — object-typed negative route + scoped evidence.** "No provider-derived bytes in a conductor surface" (unenforceable) → **typed**: raw `LLMRequest`/headers/creds/stream-frames/transport+tool envelopes are not valid conductor payload types + never auto-forwarded; agent-authored relay bodies MAY quote/summarize (relay content, no evidence status). Evidence records carry **scope** + provider-turn E3 **binds** run/turn/attempt + manifest/catalog/m-8-build/policy digests; the deny→zero-send negatives are **instrumented-test** evidence, not conductor-observed. "Attestation" → "self-reported worker-carried provider report"; MVP ladder = **E0–E3** (E4 out of scope).
- **§4 — tool-DISPATCH seam, not a capability ceiling.** Explicit: the fixed 8 govern *dispatch*, not bash effects, and are not an OS boundary. Set-equality now binds tool **identity** (name + schema/catalog/mapping version), not the name string. + **authorized==executed** one-shot-record binding (design-required).
- **§6 — wake split:** at-least-once *delivery* (m-9 catch-up `project/read` on reconnect) + at-most-once *scheduling* (`UNIQUE(relay_id)`); m-10 still never touches the conductor; push is advisory.
- **§7 — named DESIGN requirements added:** m-9 seat topology (one credential per LOGICAL seat; the operator-decision invariant); m-10 durable app-side state store + **monotonic `turn_epoch` fencing** (stale worker actively rejected) + explicit UNKNOWN/PARTIAL states (park, never replay). §1 sharpens the freeze to hash the **frozen core**, credential-attach the only permitted post-freeze transform.
- **§10 NEW — acceptance-test annex** (14 claim→proof rows + the sentinel-secret caveat + the build order).

### What I rejected from the external review (triage)
- **Option B (a real OS boundary now)** — Step-4, not MVP (we take the narrowed claim, which the review itself says fits the threat model).
- **The m-1/m-7-stamping + "lane-pinning-is-a-selection" "conflicts"** — artifacts of the *summary* the model saw; the locked charters are already consistent (m-1 owns the stamp/store *contract*; m-7 is the sole *write path* that invokes it; lane pinning binds a *previously-selected* immutable lane, no dynamic routing). No amendment change; noted for the source reconciliation.

## Verification
- **Amendment r4 SHA-256:** `57aa3170499e8f8b3fcb2f6487b8544719f1b9c107416cf323bf8e1487d27960`. Chain: r0 `a524bcbf` → r1 `02e9da1c` → r2 `3db3eb96` → r3 `e25bce10` → r4.
- Reframe packet unchanged `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 unchanged `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- **Governing 15-file manifest** (README → r4 pointer): combined `d16023ee2b8e84abb7b82ccda4195aff63fc1be42cbd18f82e4452e3e5d64b26`; README `57fd064a…`.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-172000.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — folded the external-review errata into `master/STEP-3-MVP-AMENDMENT.md` (r3→r4); updated the `master/README.md` pointer (e25bce10→57aa3170); created this transmittal + its INDEX row. No reframe-packet / canonical-m-5 / historical-relay edit; no code, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP performs a fresh exact-byte review of `STEP-3-MVP-AMENDMENT.md` r4 @ `57aa3170…` (the r3 approval is void by the byte change); on a clean return the operator ratifies (operator-authored, naming `57aa3170…`), then master folds the §7 graph + §1 fragment-supersession into every source and the first-stage DESIGNs dispatch per §7. No lock/DESIGN-lock/PLAN/T4/credential/provider authorized until then.
