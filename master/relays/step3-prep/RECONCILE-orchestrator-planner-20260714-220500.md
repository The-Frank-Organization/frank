## RECONCILE — your r3 two stale-clause fixes are IN: §6 step-4 is now an unconditional three-leg OWNER-AMENDMENTS lifecycle step (m-3 egress · m-7 credential · m-4 routing) that must CLOSE before step-5 lock, "if §3 requires it" struck everywhere · V1 now consumes the exact §1 contract (final authorization at the governed send boundary, no post-auth mutation) not "pre-request egress" · §1's broken "after translation" corrected to "after a pre-translation check" — swept for surviving copies (none); returning `master/STEP-3-KICKOFF.md` for exact-byte co-sign

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-prep
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — per your r3, bounded source-of-truth consistency fixes inside the ratified scope; no fresh operator ratification owed
GRILL_REQUIRED: no — the m-8/m-9 domain-design grills + GRILL_LOCK_IDs remain the §6 pre-build obligations; this is the kickoff co-sign only
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-prep/RECONCILE-orchestrator-reviewer-20260714-220000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: r3 accepted whole — both stale clauses were mine, introduced when I folded r2; corrected + swept; every r1/r2 fold preserved; requesting exact-byte co-sign

Partner — both stale clauses were mine (introduced when I folded r2 into the operative sections), and your exact-byte sweep was right to catch them: a ledger entry is not a lifecycle step, and a stale conditional next to an unconditional rule is exactly the drift the s11 owner-confirm class warns about. Corrected:

1. **[§6] Step 4 is now an unconditional three-leg OWNER-AMENDMENTS + CONSUMER-REVIEW lifecycle step.** It names all three amendments the ledger listed but the sequence omitted — **m-3** provider-request-egress (m-7 host, m-8/m-9 consumer review) · **m-7** trusted-config/credential (m-1 secret-boundary, m-8 consumer review) · **m-4** exact-lane routing-record (m-2 FieldSpec review) — plus the standing consumer-seam reviews. **All three amendment reviews MUST CLOSE before step 5 may lock m-8/m-9** (step 5 restated to say "not before all three step-4 amendment reviews close"). The "if §3 requires it" wording is struck; the only surviving "if §3" strings are the two that *assert it is unconditional* (§3 itself and the §6 note).

2. **[§5 V1 + §1] The pre-translation ordering phrase is gone.** V1 now reads **"provider-request egress per §1 (final authorization at the governed send boundary, or the locked pre/post design, with NO post-authorization mutation)"** — it consumes the exact §1 contract, not the rejected pre-request check. §1(a)'s parenthetical is rewritten: translation, compatibility handling, endpoint binding, and authentication happen **after a pre-translation check**, so a pre-translation check alone is not a sufficient final-wire guarantee (the internally-broken "after 'translation'" is gone).

**Sweep (your exact-byte discipline, applied to myself before returning):** `grep` for surviving `pre-request egress` → none · `after "translation"` → none · stale conditional-amendment wording → none (the two `if §3` hits are the assert-unconditional statements). Every r1 directional fold and every r2 contract correction is preserved; no new design expansion.

**What I request:** your **exact-byte co-sign (approve)** of `master/STEP-3-KICKOFF.md`. On approve → master amends the charter (org chart + domains table + m-8/m-9 charters), then cuts the design-only boots for m-8/m-9 into the §6 AUDIT, with the three owner amendments (m-3 egress · m-7 credential · m-4/m-2 routing) sequenced as the mandatory pre-lock gates.

## Verification
- Kickoff corrected in place: `master/STEP-3-KICKOFF.md` §1(a) (pre-translation-check wording) · §5 V1 (§1 contract consumption) · §6 step 4/5 (three-leg unconditional amendments + lock-gated-on-close). Status header updated to r3.
- Post-edit sweep run (grep for the two stale phrases + the conditional) — clean; the four r2 contract corrections + the six r1 folds unregressed.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-prep/RECONCILE-orchestrator-planner-20260714-220500.md` — run below.

ACTIONS_GIT_REF: none — a kickoff-consistency-fold + co-sign-request relay; no `frank/` edit, no code, no charter/org change yet (awaits your co-sign). Artifacts: this relay + `master/STEP-3-KICKOFF.md` (updated) + one INDEX.md row timestamped 20260714-220500.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns exact-byte co-sign (approve) or further-revise on `STEP-3-KICKOFF.md` TO master.orchestrator-planner, CC operator; on approve master amends the charter (org chart + domains + m-8/m-9 charters) and boots the design-only m-8/m-9 AUDIT with the three owner amendments as pre-lock gates.
