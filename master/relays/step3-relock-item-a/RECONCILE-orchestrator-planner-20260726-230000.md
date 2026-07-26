## RECONCILE — OPERATOR-DIRECTED SIMPLIFICATION: replace §4's soft-edit-stable interface BUNDLE with a plain interface-lock RECORD (the same byte-bound-hash lock used for every approval this project). The three VP rounds correctly proved the bundle apparatus over-costly for the MVP; the operator directed "do the dead simple way." Amendment `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` `680e6fcb…` routed for VP review → operator ratification. Un-fuses the exit-fixtures to lane 4 (dissolves the F1 circularity). Owners held until ratify, then RELEASED (no owner item-A work).

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this AMENDS ratified §4 (the item-A mechanism); on VP review it goes to operator hash-bound ratification (§8b). Master does not self-ratify; the operator directed the simplification, this makes it a proper record.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-220000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP review the item-A simplification amendment `680e6fcb930a1fc0f2f6c04dd02a0dd8c76d98710927c4d3ef4ee27b2b8c9476` — replace §4's bundle mechanism with a plain byte-bound interface-lock record (WITHDRAWS the r3 bundle apparatus); un-fuse the exit-fixtures to §7/lane-4; on approve → operator ratification; owners held meanwhile then released

## What this is + why
Your three item-A reviews were correct and, taken together, made the real point: the ratified §4 bundle mechanism (dedicated per-interface artifacts, an extractor tool + `--verify`, closed discriminated-union schemas, a soft-stability fixture, fixture-input pre-freeze ordering) is **far heavier than the MVP needs** — the r3-F1 circularity, r3-F2 post-marker-hash problem, and r3-F3 foreign-bytes/stale-state problem are all symptoms of forcing a soft-edit-stable extraction over already-frozen mixed design docs. The **operator directed the dead-simple path.** This amendment executes it.

**The change (§4 mechanism):** replace the bundle with `master/STEP-3-INTERFACE-LOCK.md` — a plain record that **byte-binds** the settled owner bases + the five join records + the frozen finals + the two carried limits + the env_digest-parity locus, under "**named at this hash; any change to a named byte voids the lock → re-lock.**" This is the **same lock the whole project has used** (every settled base + amendment + join was byte-bound this way; and the earlier joint interface-lock `b7e1f0ef`). No extractor, no markers, no `bundle_sha256`, no dedicated artifacts.

**Why it is sound (not scope-dodging):** the settled bases are **already frozen + pair-approved + byte-bound** — frozen design-of-record is not cosmetically edited, so soft-edit-stability (F101) solves a problem that does not arise. It aligns with the ratified confusion-firewall + MVP-minimality philosophy (cut ceremony that does not earn its cost).

## The un-fusing (§7 clarified — dissolves the r3-F1 circularity you flagged)
The exit-fixtures manifest + the carried obligations were **never part of the interface lock**; §7 already says they are **hashed at the re-lock (lane 4)**. The amendment confirms **item A = the lock record only; the exit-fixtures freeze is lane-4 work.** My r1 recipe wrongly fused them into the bundle, which created the impossible "freeze fixtures whose input digests don't exist until T4 into the item-A lock" dependency — the amendment removes the fusion, not the fixtures.

## What I ask the VP to check
- Does the plain interface-lock record **satisfy item A's purpose** — a stable, byte-exact reference lane 4 re-locks against and T4 builds to — given the bases are already byte-bound?
- Is anything in **lane 4 / the exit test** genuinely dependent on `bundle_sha256` or the extractor such that dropping them breaks a downstream gate? (My read: no — lane 4's "shorter re-lock" becomes the Master+VP ratification of the record; the exit gate rides §7 unchanged.)
- Is the un-fusing (item A = lock record; exit-fixtures = lane-4/§7) a **faithful reading of §4/§7/§11**? (My read: yes — §7 says "hashed at re-lock," §11 says "A: … bundle authored last, over settled B–E.")
- The withdrawal of the r3 recipe `06e6956e…` + the bundle apparatus.

## Owners — HELD until ratification, then RELEASED
No owner marks/nominates anything. On ratification the simple lock needs **no owner action** (their settled bases ARE the locked artifacts), so the item-A hold is **released** at that point — the owners' only remaining lane-2/item-A obligation is that their bases stay byte-bound.

## Boundaries
This is a proposed amendment to ratified §4 — it changes no byte until operator ratification. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No lock record authored yet (that is item A, post-ratification). Ratified rev12 `1125b0a0…` UNMOVED; the eight settled bases + frozen finals + three prior amendments UNMOVED. **H-12 continues to hard-block external use.**

## Verification
Amendment `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` = SHA-256 `680e6fcb930a1fc0f2f6c04dd02a0dd8c76d98710927c4d3ef4ee27b2b8c9476`. The eight settled-base files re-verified on disk this session (m-1 `d34a7c47…`, m-2 `c3a8cd61…`/`5ec7a3d2…`, m-3 `651c9aec…`, m-8 `734e44b7…`, m-9 `01b885fe…`, m-10 `3e3c5192…`/`cd17db32…`). Ratified §4 (`:82-108`) + §7 (`:363-398`) read at the bytes. r3 recipe `06e6956e…` to be withdrawn on ratification. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — the simplification amendment `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no lock record authored, no owner action opened, no self-ratification.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP reviews the simplification amendment `680e6fcb…` (does the record-lock satisfy item A + not break lane 4 / the exit gate; is the un-fusing faithful to §4/§7/§11); on approve → operator hash-bound ratification → master authors `STEP-3-INTERFACE-LOCK.md` (item A) → lane 4 (Master+VP lock + exit-fixtures freeze) → lane 5 (T4). Owners held until ratify, then released. H-12 stands.
