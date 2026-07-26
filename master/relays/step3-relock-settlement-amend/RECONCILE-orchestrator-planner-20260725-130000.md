## RECONCILE — §D-settlement amendment rev3 `ab10e6ef…`: SETTLE-VP-R2-F1/F2/F3 folded; original F1–F4 closures + both frame constants/assertions + the m-2 cell `5ec7a3d2…` preserved byte-exact; routed for VP exact-byte re-review

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — only fresh VP-approved amendment bytes may proceed to operator hash-bound ratification; master does not self-ratify
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-120000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: rev3 `ab10e6ef9987e6535510bfee12aadd618f1aa5e68570d21fd4b9d8a0b4f1befb` closes R2-F1 (envelope-saturation vs unattainable frame + two carrier shapes / one growth site), R2-F2 (timeless-fold rule replaces the stale m-10 working-state snapshot), R2-F3 (m-9 matrix now names the §6 `:423-426` classification for replacement); everything else preserved byte-exact

All three R2 findings were correct; each is folded, nothing else touched. Fresh ratification candidate: **`master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` @ SHA-256 `ab10e6ef9987e6535510bfee12aadd618f1aa5e68570d21fd4b9d8a0b4f1befb`**, bound m-2 cell **`5ec7a3d2…` preserved byte-exact** (no m-2 redispatch).

## How each finding is closed
- **R2-F1 (frame claim).** §2.4 is corrected and I do not repeat the overclaim here. The two compile-time assertions + every constant are preserved. The `4,194,304` figure is now stated as the **production limits-table / envelope sum over conservative ceilings that cannot all be attained simultaneously** — it **saturates `FRAME_MAX` by construction** (because `ADMISSION_REF_ENC_MAX` is *defined* as the residual `FRAME_MAX −` the other four ceilings), and the amendment **no longer claims any legal production frame attains 4 MiB.** The proven bound on an actual legal frame is `FRAME_CONTENT_BOUND = 3,704,832 B` (m-10's owner artifact): the production max-witness fixture asserts measured size **≤ `FRAME_CONTENT_BOUND`**, and **exact-fit / one-byte-over fixtures exist only under the build-tagged, test-only reduced limits table** (no production-equality witness — JCS has no padding). The two assertions are now labeled as covering **two carrier shapes** (`turn_open`, `attempt_open_ok`), while the parked set has **one B.4 growth site** (r40 `:81`). `resume_frame_overflow` stays statically unreachable under the production assertion + the fail-closed backstop. This matches m-10 rev14 `b96a1511…` `:66-79` at the bytes.
- **R2-F2 (stale m-10 working-state).** Removed every "current m-10 working delta `48062d18…`" / "bindable successor to it" snapshot from the packet. Replaced with a **timeless-fold rule**: no pre-ratification working artifact, in-flight revision, or relabel of any owner is a durable post-ratification fold, regardless of how far its content has advanced; after ratification each owner produces a **fresh pair-reviewed successor over the then-current owner artifact**, and any pre-folded revision is historical ancestry, never a substitute. The m-10 matrix row now folds "over a successor to the **then-current** m-10 owner artifact," not over a frozen hash. (`48062d18…` no longer appears anywhere in the packet.)
- **R2-F3 (m-9 fold locus).** The m-9 matrix row now directs m-9 to **REPLACE the pair-approved Section 6 classification — m-9 delta `04422965…` `:423-426`, the "after the bounded attempt(s) is DESCRIPTIVE … an owner clarification" text — with the amendment-controlled explicit semantic supersession (Correction 4)**, batchable with the §2.6 Gate-2 label and §7-`relay.*` folds but named as the §6 `:423-426` replacement, so the contradictory pair-approved text cannot survive. (I verified that locus at the m-9 bytes.)

## Live status, reported from current bytes (not a frozen snapshot — per R2-F2's own discipline)
As of this relay: m-10's live owner artifact is **rev14 `b96a1511…`** (routed `20260724-202500`, **not yet pair-approved**) — it has pre-folded the Correction-2 content, but per the timeless-fold rule it is **ancestry, not the durable fold**; the durable fold is a fresh post-ratification pair-reviewed successor over whatever is then-current. m-9 delta `04422965…` is pair-approved (its §6 `:423-426` is the classification the matrix now supersedes). The amendment binds **neither** as a working-state snapshot — it binds the rule and the named locus. *(Even this status line is a point-in-time report, not a packet binding.)*

## Preserved byte-exact (VP-passed, not reopened)
The original SETTLE-VP-R1-F1…F4 closures — the full cap-terminal lifecycle (§2.2), the nonterminal-threshold interpretation + 511/512/513 + multi-row-overshoot + no-truncation (§2.3), both compile-time assertions + constants (§2.4 mechanism), Correction 4's honest supersession framing (§ Correction 4 body), and the owner-set coverage — plus Corrections 1 and 3, the m-2 hash binding, the two-file packet, master-does-not-self-ratify, the exact-byte human gate, the H-12 boundary, and all downstream holds. No m-2 redispatch.

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. rev12 `1125b0a0…`, r40 `d2ce9831…`, r21 `4d3bd14e…`, the m-2 cell `5ec7a3d2…`, m-9 r12 `04422965…`, m-10 rev14 `b96a1511…` (live, not pair-approved) all UNMOVED by this relay. H-12 external-use block stands.

## Verification
Hashed on disk this session: rev3 amendment `ab10e6ef9987e6535510bfee12aadd618f1aa5e68570d21fd4b9d8a0b4f1befb`; bound m-2 cell `5ec7a3d254da…` (byte-identical to rev2); m-10 rev14 `b96a1511d015…`; m-9 r12 `044229654f25…`. Frame-bound framing reproduced from m-10 rev14 `:66-79` (production envelope saturates FRAME_MAX; `FRAME_CONTENT_BOUND = 3,704,832` legal upper bound; exact-fit only under the reduced table; single §B.4 growth site). m-9 §6 `:423-426` classification loci read at the bytes. VP r2 review target `be64e8af…` consumed. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — the rev3 amendment `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` + this relay + one INDEX.md row; no frozen/ratified byte moved, no `frank/` action, no lock issued, no self-ratification, no fold performed, no m-2 redispatch.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP exact-byte re-review of rev3 `ab10e6ef…` + the byte-bound cell `5ec7a3d2…`; on approve → operator hash-bound ratification → the propagation matrix executes (fresh post-ratification pair-reviewed successors per owner) + the §D two-sided join; on any residual → master folds and re-routes uniquely parented.
