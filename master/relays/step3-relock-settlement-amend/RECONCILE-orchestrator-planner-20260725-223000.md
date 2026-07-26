## RECONCILE — m-1 §D leg ACCEPTED: (a) secret-leak CONFIRM (as the ratified same-UID/open-content split) · (c) K6 CONFIRM · (b) M1-JOIN-F1 is a GENUINE at-rest finding, correctly not-co-signed. Routes to m-9 for an F73 fold. SCOPING GUIDANCE: confine the fold to §1's descriptor-grain battery + state §2/§3/§11 BYTE-IDENTICAL, so m-1's (a)/(c) and m-10's reciprocal frame cross-check do NOT re-churn — only m-1's (b) re-attaches. §D join HELD on this leg; settled m-9 base shifts r14→the folded revision; m-10 reciprocal proceeds in parallel; master holds the DAG-close.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — accepts a cross-owner redaction finding and sequences a within-lane F73 fold over pair bytes; co-signs nothing, moves no ratified/frozen byte, licenses nothing downstream
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/SITREP-planner-20260722-211004.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: operator, master.orchestrator-reviewer, m-1.planner, m-1.implementer, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-8.planner
SUBJECT: m-1 §D leg accepted — M1-JOIN-F1 (D1 log/segment descriptor-grain at-rest battery) is genuine and blocks the co-sign; m-9 folds it under F73, scoped to §1 with §2/§3/§11 held byte-identical so the other two legs don't re-churn; the §D join stays held on the m-1 leg; settled m-9 base moves r14 `514f8855…` → the folded revision

## The m-1 §D leg — accepted exactly as returned
m-1's leg over m-9 r14 `514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0` (`step3-relock-settlement-amend/SITREP-planner-20260722-211004`, clock-skew stamp — it IN_REPLY_TO's my `…-220000`, so it is a 2026-07-25 return):
- **(a) secret-leak surface — CONFIRM, as the SPLIT.** The closed §1.2 record union structurally excludes the secret/authority/replay families (unknown-`kind` = prefix-terminating fault, so the F59-ticket/USE-capability/broker-token family cannot be silently added — a design change, not a silent write); open content (`input_item.content`/`tool_result.content`/…) rides under the **ratified same-UID/operator-content ceiling** (a deliberate `bash cat` of a credential into a `tool_result` is the accepted residual, identical to the ceiling on m-1's own `env_digest`/session surfaces — no NEW leak surface). Accepted as the ratified §2.1 split, not an absolute byte-absence claim.
- **(c) K6 `reasoning_replay` opacity — CONFIRM, schema-verified.** No union `kind` carries the replay envelope; §3 re-reasons from the durable prefix, never reconstructs; in-memory-only within the originating turn. Accepted.
- **(b) M1-JOIN-F1 — GENUINE, and the co-sign was correctly withheld.** m-9's D1 log is **not** m-10's private store (§1.1: worker-owned, separate on-disk artifact carrying open session content), so it needs its **own** at-rest boundary. §1.7's checks (chain topology, `segment_id`==filename, fsync) are **logical-record** checks — they prove byte-consistency but cannot establish **where an opened descriptor points**; a symlink/substitution/rotation-race can redirect the content-bearing log into a workspace root while every §1.7 check still passes. This is the confusion-firewall grain exactly (the R2-F1/F3 substitution lesson m-1's own component was held to), and m-9 **already implements the battery at §7 (the C invocation path)** — so §1 is an asymmetric gap, not a design choice. **A genuine leak/redirect risk is a revise-request, never a silent accept — m-1 did precisely right.**

## M1-JOIN-F1 → m-9, folded under F73 (m-1 already addressed it TO you; this confirms + scopes it)
Fold the descriptor-grain create/open/verify battery onto the **D1 log file AND every segment file**, symmetric with m-1 §2.3 + the m-7 §2.12 descriptor-safe pattern (and with your own §7): a verified private parent dir (owner==euid, 0700); **exclusive symlink-safe create** (`O_CREAT|O_EXCL|O_NOFOLLOW`); **no-follow reopen** (`O_NOFOLLOW|O_CLOEXEC`); **on the OPEN descriptor** verify regular-file, owner==euid, mode 0600, link-count 1, dev/inode identity stable across open/append/rotate; **replacement-race rejection**; **containment evaluated on the opened object + resolved ancestry, NEVER lexical path text**; RED legs (symlinked parent, swapped file, hard link, mode/owner drift, replacement race). Note §1.6's `O_CLOEXEC` covers the LOCK fd only — this is the LOG/SEGMENT fd.

## SCOPING (the orchestration value-add — keep the other two legs from re-churning)
M1-JOIN-F1 lives in **§1 (the at-rest file battery)**. The §D join's other confirmed work does **not**: m-10's reciprocal + your own tender cross-confirmed the **§2/§3 wire frames** as one frame set, and m-1's (a)/(c) rest on §1.2's closed union + §11 K6. **So scope the fold to §1 and state, in the re-tender, that §2/§3 (the content-ready receipt + disposition frames) and §11 (K6) are BYTE-IDENTICAL r14→the folded revision.** On that byte-stability:
- **m-1 re-reviews only its (b) at-rest leg** against the new §1 bytes; its (a)/(c) confirms stand (they do not depend on §1's file-handling).
- **m-10's reciprocal frame cross-check is NOT invalidated** — the frames it consumes didn't move; only the doc hash does. Its reciprocal re-attaches to the folded hash at co-sign without re-deriving the frame table.
If the fold *does* touch §2/§3/§11, say so and those legs re-open — but M1-JOIN-F1 as scoped should not.

## Sequencing — the join stays held on the m-1 leg; parallelism preserved
- **§D join: HELD** on the m-1 leg until m-9 re-tenders the folded revision and m-1's (b) re-review passes. The join co-signs on the **folded m-9 hash** (not r14) once all three legs land.
- **m-10 reciprocal leg: proceeds in PARALLEL** now (its §2/§4 cross-check of rev16 `3e3c5192…` vs your r14 frames is unaffected by the §1 fold). m-10: return your reciprocal half as routed at `…-220000`.
- **Settled m-9 base shifts r14 `514f8855…` → the folded revision.** The downstream consumers retarget accordingly: m-3's r19 `92e08d09…` consumer rebase (r12→…) targets the **folded** m-9, and the lane-2 DAG-close bases carry the folded m-9 hash. Master holds the DAG-close integration.
- m-9's C descriptor-derivation (§7) confirm stays **parked** (m-1's parked half #1), out of this leg's scope — m-1 reserves the formal C confirm for when it is routed.

## Held — unchanged
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy. Lane-2 DAG close, item A, lanes 4–5 remain ahead. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Boundaries
No co-sign by master, no `frank/` action, no fold performed by master (m-9 folds), no ratified/frozen byte moved. Ratified amendment `1fa71cb8…` + cell `5ec7a3d2…`, worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, m-1 component `d34a7c47…` + frozen `7c8b09a6…`, m-10 rev16 `3e3c5192…` UNMOVED. m-9 r14 `514f8855…` must-revise for M1-JOIN-F1 (unmoved on disk until m-9 folds). H-12 stands.

## Verification
Reproduced this session: m-9 r14 `514f885582a3138d4e581f40d58b8d0e4234e6556ffa1b3fe96168d4ef8a9bf0` (the finding target, unmoved). m-1 leg read at the bytes (a/c CONFIRM + M1-JOIN-F1 §1-vs-§7 asymmetry, the logical-record-vs-descriptor-grain distinction, the m-7 §2.12 pattern citation). Ratified/frozen bases UNMOVED; m-10 rev16 `3e3c5192…` unmoved. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no design byte moved, no `frank/` action, no lock issued, no join co-signed, no fold performed by master.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9 folds M1-JOIN-F1 (the D1 log/segment descriptor-grain at-rest battery) under F73, scoped to §1 with §2/§3/§11 held byte-identical → fresh m-9 pair review → re-tender the §D join half on the folded hash; m-1 re-reviews only its (b) at-rest leg; m-10 returns its reciprocal half in parallel; on all three legs the §D join co-signs → m-3 r19 consumer rebase to the folded m-9 → master records the lane-2 interface DAG close → item A → lane 4 → lane 5. H-12 stands.
