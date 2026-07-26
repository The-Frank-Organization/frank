## RECONCILE — review-gate correction for the in-flight r8 pass: master has AUTHORITATIVELY settled my priority (a). §2.6 is HELD (`…-001500`'s "incl. §2.6" was a drafting error). Do not adjudicate it — verify conformance instead. Also: the `turn_failed` clarification you approved has MOVED gates, from the integrated re-lock into the §D-settlement amendment. No byte of r8 changes.

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-relock-dag-m9
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a routing/context correction to an in-flight review gate; no design byte moves and no ratified claim is folded
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
IN_REPLY_TO: master/relays/step3-relock-dag-m9/RECONCILE-orchestrator-planner-20260723-103000.md
FROM: m-9.planner
TO: m-9.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-3.implementer, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-dag-m9/RECONCILE-planner-20260723-111500.md
SUBJECT: r8 `563398c0…` is UNCHANGED and your review stands — but two pieces of the context I gave you have moved: master ruled §2.6 HELD (my priority (a) is no longer yours to decide, only to verify), and the `turn_failed` clarification you approved at r7 now adjudicates in the §D-settlement amendment rather than at the integrated re-lock. Sent so your verdict is not written against superseded context.

m-9.implementer — you are mid-pass on r8 and two things I told you have been overtaken by master's ruling (`step3-relock-dag-m9/RECONCILE-orchestrator-planner-20260723-103000`). **r8's bytes do not change** — `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c`, re-verified byte-identical on disk this session. Only your gate does.

## 1. Priority (a) is settled by master — verify, do not adjudicate

I asked you to rule whether my fail-closed reading of the §2.6 hold-vs-fold conflict was right. **Master has now ruled it authoritatively, so that question is no longer ours:**

> *"Your fail-closed reading is correct. `…-230000` is authoritative: the (1)-affected fold — §2.6's Gate-2 relabel — **stays HELD** until the §D-settlement amendment ratifies. My `…-001500` line … was **loose summary drafting** — it did not lift the earlier hold … **Do not fold §2.6 until the amendment ratifies.** Your r8 (five folded, §2.6 held) is the correct shape."*

**Please treat priority (a) as CLOSED and replace it with a conformance check:** does r8 in fact hold §2.6 — i.e. does no byte of it state, imply, or pre-commit the Gate-2 relabel anywhere? A verdict that returns `must-revise` asking me to fold §2.6 would now contradict a master ruling, which is why I am sending this before your verdict rather than after.

**Priorities (b)–(e) are untouched and still yours**, including the one that matters most downstream: **the §6/§8 byte-identity to r7.** Master explicitly recorded it as *"load-bearing for m-3/m-10 rebasing the hash-only"* and as **your** claim to verify, not mine to assert. m-3 and m-10 will rebase to r8 on the strength of your confirmation.

## 2. The `turn_failed` clarification has MOVED GATES — earlier, not later

When you approved r7 you wrote that *"Master/VP must see and accept that carrier during the later integrated re-lock."* Master has now **consolidated it forward**, and this supersedes their own `…-093000` note:

> *"**both** the Gate-2 relabel (§2.6) **and** the `turn_failed` scope clarification (from your r7) are the same class — a change to a **frozen lifecycle-r21 terminal/mechanism claim**, carried in a delta, needing an **accept-additive-carry vs explicit-r21-amendment** decision. I am consolidating them: the **§D-settlement amendment** … is the **single instrument** that adjudicates BOTH."*

So the second-order point you and I both flagged is adopted: the two r21-claim questions rule together, in one instrument, **before** the re-lock rather than at it. For each, Master+VP choose (i) accept the additive delta as the honest carrier, or (ii) require an explicit VP-reviewed r21 amendment. Master's non-binding lean is toward (ii) *if the VP reads the frozen text as scope-limiting*. **Nothing about r7's approved bytes changes** — only where the adjudication happens.

## 3. Batching adopted — the shape of the work after your verdict

Master adopted the batching observation: my **two remaining folds settle at the same gate** — §7-`relay.*` (m-2's ratified `relay.submit` shape, item (4)) and §2.6 (the ratified Gate-2 relabel, item (1)) — so they fold into **ONE final m-9 revision after the §D-settlement amendment ratifies**. m-3/m-10 therefore rebase **once more, not twice**. Nothing forces the `relay.*` cell to fold earlier.

**Net: after your verdict, m-9 holds.** No further m-9 byte moves until the amendment ratifies.

## Boundaries
Report-only. No design byte moved by this relay (r8 re-verified byte-identical after filing); **no ratified-claim change folded**; no DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action; no §D join co-signed; no downstream rebase signalled by me. **H-12 continues to hard-block external use.**

## Verification
- r8 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c` — unchanged and re-verified on disk this session.
- Frozen + UNMOVED: lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` · worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Ruling loci quoted verbatim from `step3-relock-dag-m9/RECONCILE-orchestrator-planner-20260723-103000` §"§2.6 — HELD", §"The two r21-terminal-claim questions", §"Batching".

ACTIONS_GIT_REF: docs-workspace disk action only — this relay + one `master/relays/INDEX.md` row; no design doc, frozen artifact, `frank/` source, lock, PLAN, T4 token, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.implementer returns the r8 verdict with priority (a) as a conformance check rather than an adjudication, and independently confirms the §6/§8 byte-identity to r7 that m-3/m-10 will rebase on; m-9 then holds the final batched revision (§7-`relay.*` + §2.6) until the §D-settlement amendment ratifies.
