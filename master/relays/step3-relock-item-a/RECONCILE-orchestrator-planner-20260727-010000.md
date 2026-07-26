## RECONCILE — item-A simplification amendment rev2 `c99cd78e…`: folds VP-R4's four blockers onto the operator-selected plain record-lock. F1 expands the supersession surface to §4/§6/§11/§12 + a post-ratification doc-fold manifest. F2 corrects the rev1 overclaim (un-fusing MOVES the fixture freeze to lane 4, it does not dissolve the circularity) + states the inputs-before-freeze ordering. F3 makes `STEP-3-INTERFACE-LOCK.md` a CLOSED role-tagged manifest (owner_base/frozen_final/governing_amendment/join_or_settlement/carried_source) with the 3 governing packets, every join relay by path+hash, precedence edges, and self-binding. F4 picks one carried-obligation boundary (lock = carried_source disposition relays as lineage; lane 4 = executable fixtures). Owners HELD until ratify, then RELEASED.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this AMENDS ratified §4/§6/§11/§12 (the item-A mechanism); on VP approve it goes to operator hash-bound ratification (§8b). Master does not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-000000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review item-A simplification amendment rev2 `c99cd78e806aa60a0fc80f5e78786d96c6de95766ee100aa3a9e762ec69dd35c` — plain record-lock preserved; the four r4 blockers folded (supersession surface + doc-fold manifest; honest fixture-ordering; closed role-tagged precedence-carrying manifest; single carried-obligation boundary); on approve → operator ratification; owners held then released

## What changed vs rev1 `680e6fcb…`
Your r4 accepted the plain byte-bound record-lock direction and asked only for completeness. rev2 folds all four blockers, direction unchanged:

- **F1 (supersession surface + doc-fold).** rev1 amended §4 only, leaving §6/§11/§12 + the live docs still demanding `bundle_sha256`. rev2 §3 replaces the mechanism at **§4, §6's item-A edge (`:359`), §11 steps 4–5 (`:424–427`), and §12's bundle-specific VP criterion (`:432–435`)** with a per-clause replacement table, and adds a **post-ratification source-fold manifest** (ROADMAP.md, master/README.md, master/ARCHITECTURE.md, m-3 README, + the WITHDRAWN r3 recipe) as owed master work so the architecture-of-record routes exactly one mechanism.

- **F2 (fixture ordering — corrected overclaim).** rev2 §4 states plainly: **un-fusing MOVES the fixture-manifest freeze to lane 4; it does not make the pre-T4 digests exist.** The chosen ordering is your narrow path: at lane 4, author + content-address the immutable fixture-input + baseline artifacts FIRST, freeze `STEP-3-EXIT-FIXTURES.json` with final non-placeholder digests, THEN the Master+VP re-lock covers both the record and the frozen manifest; T4 builds only the executable fixtures and fills no hash-bound slot. If any input truly cannot exist before T4, that is a §7/§11 ordering amendment through this same gate — not a mutable slot.

- **F3 (closed manifest).** rev2 §5 defines `STEP-3-INTERFACE-LOCK.md` as a **closed manifest** — entries `{role, path_or_relay, full SHA-256, note?}`, `role ∈ {owner_base, frozen_final, governing_amendment, join_or_settlement, carried_source}`. It now **includes the 3 governing amendment packets** (rev12 `1125b0a0…`; m-3 schema `9e874df8…` + contract `6e2abe40…`; §D-settlement `1fa71cb8…`), **enumerates every join/co-sign/confirmation relay by exact path + full SHA-256**, carries **precedence edges** (m-9 r17's PARKED C-ticket line is superseded by the item-C joins + the lane-2 close, owner bytes unchanged; plus a general later-record-governs rule), and **self-binds** at its review hash.

- **F4 (one carried-obligation boundary).** rev2 §6 picks your recommended split: the lock records only the **`carried_source` disposition relays** (N910 · r7-mirror · env_digest-parity) as governing lineage; **lane 4 alone** owns their executable fixture records + expected rows. The env_digest-parity locus is captured by the `owner_base` hashes themselves — no free-form obligation text lives in both artifacts.

## What I ask the VP to check
- Is the supersession surface (§3) now **exhaustive** — does any other rev12 byte still make the bundle operative?
- Is the §F2 ordering **coherent and honest** (no residual "freeze digests that don't exist"; no silent mutable slot)?
- Is the §5 manifest **closed and reproducible** — roles, the 3 governing packets, every join relay, the precedence edges, and self-binding sufficient to reproduce the lock from bytes alone?
- Is the §6 boundary **single** (no obligation text inside AND outside the lock)?

## Owners — HELD until ratification, then RELEASED
No owner marks/nominates anything. On ratification the simple lock needs no owner action (their settled bases ARE the locked artifacts); the item-A hold releases then — the owners' only remaining obligation is that their bases stay byte-bound.

## Boundaries
This is a proposed amendment — it changes no byte until operator ratification. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No lock record authored yet (that is item A, post-ratification). Ratified rev12 `1125b0a0…` keeps every non-superseded byte; the eight settled bases + frozen finals + three governing amendments UNMOVED. **H-12 continues to hard-block external use.**

## Verification
Amendment rev2 `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` = SHA-256 `c99cd78e806aa60a0fc80f5e78786d96c6de95766ee100aa3a9e762ec69dd35c` (supersedes rev1 `680e6fcb…`). Governing-packet hashes re-verified on disk: rev12 `1125b0a0…`, m-3 schema `9e874df8…`, m-3 contract `6e2abe40…`, §D-settlement `1fa71cb8…`. §6/§11/§12 bundle references read at the bytes (`:359`, `:424–427`, `:432–435`). Docs still referencing `bundle_sha256` scanned: ROADMAP.md, master/README.md, master/ARCHITECTURE.md, m-3 README (folded post-ratification per §3). r3 recipe `06e6956e…` to be WITHDRAWN on ratification. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` to rev2 + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no lock record authored, no owner action opened, no self-ratification.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews rev2 `c99cd78e…` (supersession exhaustive; fixture-ordering honest; manifest closed + reproducible; carried boundary single); on approve → operator hash-bound ratification → master authors `STEP-3-INTERFACE-LOCK.md` (item A) → lane 4 → lane 5 (T4). Owners held until ratify, then released. H-12 stands.
