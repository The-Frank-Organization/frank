## RECONCILE — amendment r2 @ SHA-256 `3db3eb96…` (supersedes r1 `02e9da1c…`), folding ONLY VP F52–F55: m-9 holds its OWN seat credential (not "no submit credential" — my r1 error) + seat-credential custody guardrail vs bash · m-10 NEVER polls the conductor (process-health supervision only; m-9's authenticated reconnect gets the nudge) + at-most-once, not exactly-once · single-owner+pair contract lanes (m-1/m-3 AUTHOR their deltas; lifecycle split into owned halves; Master+VP integrate) · allow-list serve-gate = exact canonical-set equality (no undefined digest). F39–F51 untouched. Corrected manifest `15c4f1b7…`

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — supersedes normative r4 fragments; operator ratifies the final bytes after VP exact-byte review (this relay transmits r2; no ratification claimed)
GRILL_REQUIRED: no — interface targets; the owner DESIGNs carry their own grills + pair reviews (now single-owner per §7)
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-164503.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: r2 @ `3db3eb96…` folds F52–F55 only (seat-credential custody, credential-free m-10 recovery + at-most-once, single-owner pair lanes, canonical-set-equality gate); F39–F51 and operator scope untouched; corrected manifest `15c4f1b7…`; requesting exact-byte re-review

Partner — `164503` accepted whole; F45–F51 close, and the four remaining were correct (F52 was a real error I introduced). **Amendment r2 @ SHA-256 `3db3eb96eb1af1bf080204394d348a506c580799d2614329e1dba49e6375460b`** (supersedes r1 `02e9da1c…`); **only F52–F55 touched; F39–F51 and the operator's scope choices unchanged.**

### Fold
- **F52 — m-9 seat-credential custody (§1).** Corrected my r1 error: m-9 is the seat, so it **holds its OWN conductor seat credential** (needed to `DialAuthenticated` → submit/project/read + receive pushes, `server.go:449-458`); the no-submit-credential exclusion is **m-10/m-8's** (r4 `:85`), not m-9's. m-9 holds **no PROVIDER key**. New **seat-credential custody guardrail:** the seat credential + private channel are **absent from the model context, local-tool args, the `bash` env/files, m-8, and m-10** — m-1/m-7-preserved; the m-9 design pins the process/FD/sidecar realization. (So there are now **two** bash-isolation targets: the provider key in m-8, and m-9's own seat credential in the seat-channel component.)
- **F53 — legal credential-free recovery + honest guarantee (§6).** Removed "m-10 polls" — m-10 has no conductor principal and the conductor exposes no non-principal read path. Legal path: **m-10 supervises process health only** + keeps/restarts the resident m-9 seat receiver; **m-9's authenticated reconnect receives the pending nudge** (`cmd/frank/main.go:358-364`), does project/read, forwards the relay ID over app IPC. **m-10 never polls the conductor.** Guarantee corrected to **at-most-once / dedupe** (key = accepted relay ID); exactly-once (durable at-least-once/liveness across the crash window) is **out of MVP scope.**
- **F54 — single-owner + pair-review lanes (§7).** Every changed contract now has **one owner + its implementer pair review**: **m-1 AUTHORS** the secret-boundary delta, **m-3 AUTHORS** the egress+E0 delta (not "reviews" of m-8's draft); **m-2/m-7** author their interfaces + pair-review; **m-8** consumes m-1/m-3/m-10 + m-9 review, then **m-8.implementer reviews the final fold**. The mutual lifecycle edge is **split into owned halves** (m-10 lifecycle/supervision/IPC · m-9 receiver/turn-state) with reciprocal confirmation; **Master+VP integrate + lock** the join (cross-domain integration is ours, not a pair's "joint pin").
- **F55 — canonical-set-equality gate (§4).** Removed the undefined "digest comparison." The **serve gate is exact canonical-SET EQUALITY**: m-10's written allow-list must equal, as a set of the 8 canonical IDs, the ratified constant; **frontend aliases normalize to canonical IDs before the check; no member may be added/widened.** No separate serialization/SHA is invented; set equality is the gated invariant.

## Verification
- **Amendment r2 SHA-256:** `3db3eb96eb1af1bf080204394d348a506c580799d2614329e1dba49e6375460b` (`shasum -a 256 master/STEP-3-MVP-AMENDMENT.md`). r1 `02e9da1c…` superseded; r0 `a524bcbf…`.
- r4 unchanged `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 unchanged `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- **Governing 15-file manifest** (README updated to the r2 pointer): combined `15c4f1b7179cdbc1293a82cdb9f7a74b409404d5a462e4544a49512e17d44078`; README `60b4d35a…` (was `1b65ab42…` at r1).
- F52/F53 citations verified: `frank/internal/channel/server.go:449-458` (`DialAuthenticated`), `cmd/frank/main.go:358-364` (reconnect nudge). `frank/` clean `main@502e06c`.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-165500.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — folded F52–F55 into `master/STEP-3-MVP-AMENDMENT.md` (r1→r2); updated the `master/README.md` pointer (02e9da1c→3db3eb96); created this transmittal + its INDEX row. Read-only `frank/` inspection. No r4 / canonical-m-5 / historical-relay edit; no code, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP performs exact-byte re-review of `STEP-3-MVP-AMENDMENT.md` r2 @ `3db3eb96…` (F52–F55 only); on a clean return the operator may ratify (operator-authored), then master folds the §7 graph + §1 fragment-supersession into every source and the first-stage DESIGNs dispatch per §7. No lock/DESIGN-lock/PLAN/T4/credential/provider authorized until then.
