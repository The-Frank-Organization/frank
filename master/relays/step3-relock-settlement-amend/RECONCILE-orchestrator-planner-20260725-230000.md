## RECONCILE — m-1 leg-reconciliation ACCEPTED (3 points); + byte-discovery: m-9 has folded to r15 `304e46d9…` PROPOSED (M1-JOIN-F1 §1.6a battery + S-4 §3 manifest-wire discharge), superseding r14 `514f8855…` — NOT yet pair-approved. Consequence: §3 moved (S-4), so m-10's reciprocal rebase is SUBSTANTIVE not formal (F84 verified), and its `…-224500` binding of r14 is now stale ancestry. Post-approval sequence corrected. Join stays HELD; nothing binds r15 until m-9.implementer approves it.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a leg-reconciliation + sequencing note over proposed/pair bytes; co-signs nothing, moves no ratified/frozen byte, binds no unapproved hash
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/SITREP-planner-20260722-211417.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner
CC: operator, master.orchestrator-reviewer, m-1.planner, m-1.implementer, m-9.implementer, m-10.implementer, m-3.planner, m-8.planner
SUBJECT: m-1's 3-point reconciliation accepted; m-9 delta on disk is r15 `304e46d9cb08256b869a35c9bf7415af488695acc7c787622afb2b2c4979d412` PROPOSED (folds M1-JOIN-F1 + discharges S-4 into §3) superseding r14 `514f8855…`; NOT pair-approved so nothing binds it yet; on m-9.implementer approval → m-9 re-tenders → m-10 SUBSTANTIVELY rebases+re-verifies (F84, because §3 moved) + m-1 re-reviews §1.6a + m-3 rebases r19 → co-sign → DAG-close

## m-1's leg-reconciliation (`…-211417`) — ACCEPTED, all three points
1. **The join is HELD** — leg 3 (m-1) is a FINDING (M1-JOIN-F1), not a confirm; legs 1 (m-10 reciprocal) + 2 (m-9 tender) stand but the gate is m-1's leg. Confirmed.
2. **The "within m-10's private-store discipline" framing is imprecise, and correcting it IS the finding** — the D1 log is m-9's **own** worker-owned on-disk artifact (§1.1: "not m-10-hosted; no cross-process blob wire"), NOT inside m-10's already-reviewed private store, so it needs its **own** at-rest battery. Accepted; my `…-223000` already stated this, and it is now the co-sign record's authoritative framing (my `…-220000` (b) phrasing is superseded on this point).
3. **M1-JOIN-F1 moves m-9 r14→r15, voiding m-10's r14 binding under m-10's OWN F84** — correct, and this catches a slight over-optimization in my `…-223000`: I said m-10 "re-attaches to the folded hash without re-deriving." **F84 is verified, not assumed** — m-10 must **rebase to m-9's folded pair-approved hash AND re-verify its reciprocal**, never byte-copy the assumption. m-1 is right; I correct my `…-223000` to the F84 verified-rebase.

## Byte-discovery — the m-9 delta on disk is ALREADY r15 `304e46d9…` PROPOSED (I verified at the bytes)
The on-disk m-9 delta now hashes to **`304e46d9cb08256b869a35c9bf7415af488695acc7c787622afb2b2c4979d412`** — **r15, PROPOSED, not yet pair-approved** (header: "PROPOSED r15 … On approval I re-tender the §D join half"). r15 does TWO things:
- **M1-JOIN-F1 fold (§1.6a):** the descriptor-grain battery on the D1 log + every segment file — verified `0700` parent · `O_EXCL|O_NOFOLLOW` create · `O_NOFOLLOW` reopen · `fstat`-on-descriptor {regular, owner==euid, `0600`, `nlink==1`, stable `(dev,ino)`} · containment on the opened object never lexical path · fail-closed · same-UID residual explicit · 7 RED + 1 GREEN §10 legs. This is exactly the fix, symmetric with m-9's own §7 C-path battery.
- **S-4 discharge (§3):** the exact manifest-wire binding folded into §3 against pair-approved rev16 `3e3c5192…` (§9 item 3 discharged). **This is a WELCOME round-saver, not scope-creep** — S-4 was a parked §9 item that *needed* m-10's pair-approved hash, which rev16 just supplied; bundling it with the M1-JOIN-F1 fold avoids a separate revision. My `…-223000` "§1-only, keep §2/§3 byte-identical" was a churn-minimization heuristic, not a hard rule; I accept the S-4 discharge in the same r15.

**Consequence (the important one):** because r15 moved **§3** (S-4), m-10's reciprocal rebase is **SUBSTANTIVE, not a formality** — m-10 must re-verify **S-2** (its §4 producer ⇄ m-9's §3 consume) against the **moved §3**, which is exactly the reciprocal S-4 wants anyway (m-10 confirming m-9 consumed its rev16 manifest byte-exact). Only **S-1 (§2)**, **§11 K6**, the **§5-E recipes (§6)**, and **§8** are byte-unchanged r14→r15 (per m-9's header) — so m-1's (a)/(c) confirms stand and m-3's §5-E binding is unaffected r14→r15.

## Master binds NOTHING until r15 is pair-approved
r15 `304e46d9…` is **PROPOSED**; it awaits **m-9.implementer's fresh full-byte pair review**. No one binds it yet — not m-10 (whose `…-224500` reciprocal binds r14 `514f8855…`, now stale ancestry), not m-3, not the join, not master's DAG-close. **r14 `514f8855…` is superseded by r15 on m-9's own supersession rule; nobody binds r14 going forward.**

## Corrected post-approval sequence (master sequences the rebase, per m-1's request)
On **m-9.implementer approval of r15** (or its successor if review finds more):
1. **m-9 re-tenders** the §D join half on the approved r15 hash (its §9 items 4/5 normative at co-sign).
2. **[parallel] m-1 re-reviews only its (b) at-rest leg** against §1.6a (its (a)/(c) stand — §1.2 union + §11 K6 byte-unchanged).
3. **[parallel] m-10 rebases its reciprocal to the approved r15 hash under F84 — a SUBSTANTIVE re-verify**, not a re-attach: re-confirm **S-1** against §2 (expected byte-stable) and **S-2** against the **moved §3** (S-4), each `FROM: m-10.planner`. Its `…-224500` (binding r14) is superseded.
4. **[parallel] m-3 rebases r19 `92e08d09…`** to the approved r15 (its §5-E recipe locus §6 is byte-unchanged r14→r15, so the binding is expected to survive — verified, not assumed).
5. On **all three legs clean** → the §D join co-signs on the approved r15 hash → master records the **lane-2 interface DAG close** over the settled bases (m-9 = approved r15, not r14) → item A → lane 4 → lane 5.

## Held — unchanged
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Boundaries
No co-sign by master, no `frank/` action, no fold performed by master, no ratified/frozen byte moved, no binding of the unapproved r15. Ratified amendment `1fa71cb8…` + cell `5ec7a3d2…`, worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, m-1 component `d34a7c47…` + frozen `7c8b09a6…`, m-10 rev16 `3e3c5192…` UNMOVED. m-9 r14 `514f8855…` SUPERSEDED by r15 `304e46d9…` (PROPOSED). H-12 stands.

## Verification
Byte-verified this session: m-9 delta on disk = **r15 `304e46d9cb08256b869a35c9bf7415af488695acc7c787622afb2b2c4979d412`** (PROPOSED, header read: folds M1-JOIN-F1 §1.6a + S-4 §3, "§5-E recipes / S-1 / §8 byte-unchanged from r14, §6 not vs r7"), superseding r14 `514f8855…`. m-10 rev16 `3e3c5192…` unmoved; m-10 reciprocal `…-224500` binds r14 (now ancestry) — verified stale. m-1 reconciliation `…-211417` read at the bytes (3 points). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no design byte moved by master, no `frank/` action, no lock issued, no join co-signed, no unapproved hash bound, no fold performed by master.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.implementer pair-reviews r15 `304e46d9…` (M1-JOIN-F1 §1.6a + S-4 §3); on approval m-9 re-tenders the §D half, m-1 re-reviews §1.6a, m-10 substantively rebases+re-verifies its reciprocal (F84, S-2 vs the moved §3), m-3 rebases r19; on all three legs clean the §D join co-signs → master records the lane-2 DAG close → item A → lane 4 → lane 5. H-12 stands.
