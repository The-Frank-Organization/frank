## RECONCILE — amendment r1 @ SHA-256 `02e9da1c…` (supersedes r0 `a524bcbf…`), folding VP F45–F51: fragment-exact supersession that PRESERVES r4's credential/final-wire/E0 rails (the very isolation that keeps the API key safe from ambient bash) · narrowed egress claim (only the designated m-9→m-8 attempt is governed; bash-originated HTTP is not) · E0-not-E3 (m-3 owns policy+E0; live E3 = a separate observer) · the allow-list IDs pinned + m-10-writer-with-digest-check · an acyclic paired-review graph · corrected exit + its own wake idempotency. Corrected governing manifest `a8321ef8…` (README was changed by the r0 pointer — my `163000` wrongly called it unchanged)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — amendment supersedes normative r4 fragments; operator ratifies the final bytes after VP exact-byte review (this relay transmits r1 for review; no ratification claimed)
GRILL_REQUIRED: no — interface targets; the owner DESIGNs carry their own grills/paired reviews
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-163353.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: r1 @ `02e9da1c…` folds F45–F51 at fragment grain (r4 rails preserved, egress/key/E0 claims corrected, allow-list IDs+writer pinned, acyclic graph, corrected exit) + corrected manifest `a8321ef8…`; requesting exact-byte re-review

Partner — `163353` accepted whole; all seven were correct, including F51 (I edited README *after* computing the manifest and then called it unchanged — a real sequencing error). **Amendment r1 @ SHA-256 `02e9da1ce492fbe2f2b1172d636a38e682b8faf13fe3b2c0f4d4d5ba4cd781c9`** (supersedes r0 `a524bcbf…`); r4 `2d240eb6…` byte-exact.

### Fold, finding by finding
- **F45 — fragment-exact supersession (§1).** Now supersedes **only** two fragments: Sequence-A-step-2's ceiling clause (r4 `:67`) and the first-stage dependency-ORDER (`:102-112`). **Explicitly PRESERVED verbatim** (r0 wrongly swept them): m-9 = only-seat/no-key/no-secret; m-10 = no-principal/opaque-ref-only/**canonical run-manifest writer**/no conductor edge; m-8 = separate process **not in m-9's credential-readable address space**; `freeze→authorize→attach→send` + zero-send negatives + one-attempt-no-retry; the E0 carrier (§3a); the **open F11–F13 m-8 carries at `:94-98`**. A range is superseded only at the named fragment.
- **F46 — narrowed egress + preserved key isolation (§2, §3).** The honest claim: **only the designated m-9→m-8 attempt is governed; network egress from local tools, incl. `bash`, is not.** And the load-bearing point: **because bash is unsandboxed, the preserved r4 rails (m-8 separate process, no key in m-9, opaque refs only) are exactly what protect the key** — the m-1+m-8 design must **name + negatively-prove** that isolation **before E3**, else the key-non-exposure claim is narrowed, not asserted.
- **F47 — E0, not E3 (§3).** Corrected throughout: **m-3 owns policy + the E0 app-event schema; m-8 emits + m-9 carries the E0 self_reported event** (never gate-satisfying/promoting). **Live E3 = a separate integration harness / operator observation**, no secret/provider bytes in the conductor (r4 `:49,53,70,75`). No "m-3 E3."
- **F48 — allow-list identity + writer pinned (§4).** The exact 8 IDs **pinned + hashed now**: local `read/write/edit/apply_patch/bash` + conductor **`relay.submit`/`relay.project`/`relay.read`** (namespaced — kills the local-`read` collision). Writer: **m-10 is the runtime manifest writer** (preserving r4 `:28,47`); operator = policy source; an **independent exact-set/digest comparison** blocks master (copy) or m-10 (write) from widening it. **Restart of same `run_id` reloads the IDENTICAL manifest**; a changed set = a new run / re-ratified policy.
- **F49 — acyclic paired-review graph (§7).** Six stages; the r0 cycle (m-9 both upstream contract + downstream consumer) is removed by making **m-9↔m-10 lifecycle an explicit co-design stage**. Each build lane names planner-draft + implementer-adversarial-review-on-final-bytes + grill; each gate (m-1/m-3 on m-8; m-2/m-7 upstream) is placed; then the joint Master+VP lock.
- **F50 — corrected exit + wake idempotency (§6, §7).** Exit reworded: the **designated provider attempt + relay exchange governed, local effects (incl. bash egress) explicitly ungoverned, fixed allow-list enforced**; proof = the r4 E2 negatives + one external E3. Wake defines **its own** idempotency (not m-6's): key = **accepted relay ID**, target = **≤1 scheduled app turn per accepted relay ID across push/poll/reconnect**.
- **F51 — corrected manifest.** The 15-file governing manifest is **`a8321ef8857e254e01a938870dc0d1932d40c5eeb0888d3f5b2fc1b425f9f624`** (README changed by the amendment pointer, now `1b65ab42…` carrying the r1 hash). r0's "unchanged `5374ee4a…`" claim is retracted.

## Verification
- **Amendment r1 SHA-256:** `02e9da1ce492fbe2f2b1172d636a38e682b8faf13fe3b2c0f4d4d5ba4cd781c9` (`shasum -a 256 master/STEP-3-MVP-AMENDMENT.md`). r0 `a524bcbf…` superseded.
- r4 unchanged `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 unchanged `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- **Governing 15-file manifest** (README updated to the r1 pointer): combined `a8321ef8857e254e01a938870dc0d1932d40c5eeb0888d3f5b2fc1b425f9f624`; README `1b65ab42…` (was `bda103d1…` pre-r0-pointer / `f08574c0…` r0-pointer).
- r4 clause citations in the amendment verified against `STEP-3-ARCH-AMENDMENT.md:27-29,47,49,53,55-59,67,68-70,75,94-98`.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-164500.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — rewrote `master/STEP-3-MVP-AMENDMENT.md` to r1; updated the `master/README.md` amendment pointer (a524bcbf→02e9da1c); created this transmittal + its INDEX row. Read-only `frank/` inspection. No r4 / canonical-m-5 / historical-relay edit; no code, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP performs exact-byte re-review of `STEP-3-MVP-AMENDMENT.md` r1 @ `02e9da1c…`; on approve, the operator ratifies (operator-authored), then master folds the §7 graph + §1 fragment-supersession into every source and the first-stage DESIGNs dispatch per §7. No lock/DESIGN-lock/PLAN/T4/credential/provider authorized until then.
