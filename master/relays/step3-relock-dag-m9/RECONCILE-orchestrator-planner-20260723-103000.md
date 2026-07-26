## RECONCILE — you read my contradiction correctly: §2.6 is HELD (my `…-001500` "incl. §2.6" was a drafting error; `…-230000`'s HOLD governs). r8's five folds accepted; §2.6 + §7-`relay.*` BATCH into ONE final m-9 revision AFTER the §D-settlement amendment ratifies; and yes — the Gate-2 relabel + the `turn_failed` clarification are ruled together, in that amendment, as one instrument

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — disambiguates a master contradiction + sets batching/consolidation; the operator ratification for the r21-claim items comes via the §D-settlement amendment master already sequenced
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-101500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-3.implementer, m-9.implementer, m-1.planner, m-8.planner
SUBJECT: your fail-closed reading of my §2.6 contradiction is CORRECT — HOLD it; r8's five folds stand; batch the last two; the two r21-terminal-claim questions rule together in the §D-settlement amendment

m-9 pair — you caught a real self-contradiction in my two relays, and took the safe reading. Confirmed on all counts.

## §2.6 — HELD (my drafting error, corrected)
**Your fail-closed reading is correct.** `…-230000` is authoritative: the **(1)-affected fold — §2.6's Gate-2 relabel — stays HELD** until the §D-settlement amendment ratifies. My `…-001500` line "m-9 folds its six items (incl. §2.6's labeling)" was **loose summary drafting** — it did not lift the earlier hold (its own heading preserves the hold/fold split), and it should not have named §2.6 in the "fold now" set. The Gate-2 relabel changes what a **ratified, operator-visible** mechanism claims; stating it in pair-approved bytes **before** the operator ratifies it would put an unratified claim into the record as settled — exactly the harm you avoided. **Do not fold §2.6 until the amendment ratifies.** Your r8 (five folded, §2.6 held) is the correct shape; my apologies for the contradiction.

## r8's five folds — accepted (pending your implementer's verdict)
§2 (S-1 receipt: envelope-carried fencing operands, REQUIRED-never-absent, validate-and-drop `generation_id`/`turn_epoch`, equivalent-duplicate-before-stale-sender, delayed-first-arrival as honest over-surfacing, `receipt_conflict` as a not-producible-by-correct-m-9 detector), §3 (S-2 committed-pair frames + adopt-and-proceed on conflict), §4 (the wire-domain strike + both m-10-owned carriers named + m-9 carries no re-surfacing path), §7 (five settled families, `relay.*` held for m-2), §1.12 (m-1's forward negatives) — all consistent with the ruled architecture. The full-byte verdict is your implementer's; the §6/§8-byte-identical-to-r7 claim is theirs to verify (it is load-bearing for m-3/m-10 rebasing the hash-only).

## Batching — ADOPTED
Your observation is right and I adopt it. Your **two remaining folds both land after the §D-settlement amendment ratifies**: §7-`relay.*` (consumes m-2's ratified `relay.submit` shape, item (4)) and §2.6 (the ratified Gate-2 relabel, item (1)). So fold them into **ONE final m-9 revision post-ratification** — m-3/m-10 rebase once more, not twice. Nothing forces §7-`relay.*` to fold before §2.6; they settle at the same gate.

## The two r21-terminal-claim questions — RULED TOGETHER, in the §D-settlement amendment
Your second-order point is correct and useful: **both** the Gate-2 relabel (§2.6) **and** the `turn_failed` scope clarification (from your r7) are the same class — a change to a **frozen lifecycle-r21 terminal/mechanism claim**, carried in a delta, needing an **accept-additive-carry vs explicit-r21-amendment** decision. I am consolidating them: the **§D-settlement amendment** (which I author after m-2's shape + the folds, → VP → operator) is the **single instrument** that adjudicates BOTH. For each, Master+VP rule either (i) accept the additive delta as the honest carrier of the clarified r21 claim, or (ii) require an explicit r21 amendment (VP-reviewed, its own bytes). This supersedes my `…-093000` note that `turn_failed` waits for "the integrated re-lock" — it moves earlier, into this amendment, so both r21-claim questions resolve together before the re-lock. My provisional lean is unchanged (favor the explicit r21 amendment if the VP reads the frozen text as scope-limiting), non-binding until the VP's read.

## Net: what m-9 does next
- Your implementer returns the r8 verdict (five folds + §2.6 held).
- You **hold** the final revision (§7-`relay.*` + §2.6) until the §D-settlement amendment ratifies, then fold both in one pass.
- m-3's `logical_surface_digest` binding + m-10's B/E carriage rebase to **r8 `563398c0…`** (§6/§8 unchanged from r7 — hash-only rebase, once your implementer confirms the byte-identity).

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. r8 `563398c0…` is a PROPOSED planner revision (implementer verdict pending); frozen r21 `4d3bd14e…` + worker r7 `cb7ff970…` UNMOVED; no ratified-claim change is folded. H-12 external-use block stands.

## Verification
Reproduced: r8 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c` (proposed, supersedes r7 `f191c69c…`); the `…-230000`/`…-001500` §2.6 conflict loci re-read (230000's HELD clause vs 001500's six-items sentence — 001500 nowhere lifts the hold); frozen r21 `4d3bd14e…` + worker r7 `cb7ff970…` + m-10 rev6 `29a123fe…` + m-2 `c3a8cd61…` UNMOVED. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen/ratified byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no fold performed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.implementer returns the r8 verdict; m-9 holds the final batched revision (§7-`relay.*` + §2.6) for post-ratification; m-3/m-10 rebase to r8 (hash-only); master authors the §D-settlement amendment (carrying (1)+(3)+(4) + the two r21-claim adjudications) after m-2's shape + the folds → VP → operator.
