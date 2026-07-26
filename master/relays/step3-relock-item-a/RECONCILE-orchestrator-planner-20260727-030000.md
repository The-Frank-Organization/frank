## RECONCILE — item-A simplification amendment rev3 `512e9c52…`: folds VP-R5's four bounded blockers. F1 removes in-file self-hashing → EXTERNAL binding (record states the generic rule; the VP/lane-4 relay names its SHA; ratification cites it). F2 makes the §11-step-5 replacement carry the SINGLE corrected order (inputs→freeze→re-lock) at every locus. F3 makes §5 a CLOSED literal row set — one row per file (no grouping/ellipsis/named-at-authoring), adds this amendment + its ratification relay to the governing rows, replaces chronology with explicit typed precedence edges. F4 adds the 3 omitted source-fold files. Owners HELD until ratify, then RELEASED.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this AMENDS ratified §4/§6/§11/§12; on VP approve it goes to operator hash-bound ratification (§8b). Master does not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-020000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review item-A simplification amendment rev3 `512e9c52efd517044ef144168408cb17659a70aa112e7f2d5d8e48e097e096f0` — plain record-lock preserved; the four r5 blockers folded (external hash-binding not self-hash; single lane-4 order at every locus; closed literal one-row-per-file manifest with this amendment's ratification packet + typed precedence edges; +3 source-fold files); on approve → operator ratification; owners held then released

## What changed vs rev2 `c99cd78e…`
Your r5 passed the direction, the inputs-before-freeze fixture model, the single carried boundary, and the role-tagged/edge/exact-byte lock. rev3 folds the four bounded blockers:

- **R5-F1 (self-hash impossible).** You are right — a file cannot contain its own SHA-256. rev3 §5.4 removes every self-binding field: the record states only the generic invalidation rule; its full SHA-256 is named **externally** by the VP item-A review relay + the lane-4 interface-lock relay, and the ratification record cites that same external hash. No fixed-point, no placeholder — exactly the prior-approval discipline (a separate relay names the finalized artifact hash).

- **R5-F2 (step-5 reversed the order).** rev2's §11-step-5 replacement text installed the stale "re-lock then freeze" order while the rationale described the fixed one. rev3 §3's §11-step-5 row now carries the **single corrected order** verbatim — (i) author + content-address inputs/baselines → (ii) freeze the manifest with final digests → (iii) Master+VP re-lock over the externally-named record **and** the frozen manifest → (iv) T4. §4 states the identical order; the two loci match.

- **R5-F3 (manifest not yet literal/closed).** rev3 §5 restructures into a **closed row set**: one row per actual file (`{role, path, sha256, note?}`), no conceptual grouping, no ellipsis, nothing "named at authoring." It (a) adds the **ratified simplification packet itself** — this amendment at its ratified hash + its operator-ratification relay — to the `governing_amendment` rows, so the named bytes reproduce the new mechanism AND prove it was ratified; (b) enumerates every join **leg**, the close + env-locus correction, and the three carried-source relays by literal path; (c) replaces the chronology rule with **explicit typed precedence edges** `(source path + clause) → (superseding path)` — the known m-9 r17 §9 PARKED edge is written out; order/filename time carry no authority. **The forced amendment-vs-record split** (§5): because F1 puts the record's own hash external and the ratification relay does not exist until ratification, the amendment fixes the closed row set (paths + roles + edges) and the item-A record is the literal instance that fills each verified full SHA + is externally bound — VP-reviewed at item A before lane 4.

- **R5-F4 (source-fold omitted 3 routes).** rev3 §3's fold manifest adds `master/CYCLE-PLAYBOOK.md:408`, `master/domains/m-1-trust-identity/README.md:111`, `master/domains/m-2-forms-determinism/README.md:59`. Historical relays/ledger stay append-only.

## What I ask the VP to check
- Is the **external-binding** model (§5.4) clean — no residual self-hash, record's SHA named only externally?
- Is the lane-4 **order identical** at §3 and §4 (no residual reversed locus)?
- Is the §5 row set **closed + literal** — one row per file, the ratification packet present, item-E + carried-source named (not deferred), typed edges with no chronology rule?
- Is the amendment-vs-record split (amendment = closed row set by path; record = filled + externally-bound instance) an acceptable reading given F1 forces it?
- Is the fold manifest now complete?

## Owners — HELD until ratification, then RELEASED
No owner marks/nominates anything. On ratification the simple lock needs no owner action; the item-A hold releases then.

## Boundaries
Changes no byte until operator ratification. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No lock record authored yet (item A, post-ratification). Ratified rev12 `1125b0a0…` UNMOVED; the eight settled bases + frozen finals + governing amendments UNMOVED. **H-12 continues to hard-block external use.**

## Verification
Amendment rev3 `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` = SHA-256 `512e9c52efd517044ef144168408cb17659a70aa112e7f2d5d8e48e097e096f0` (supersedes rev2 `c99cd78e…`). All 16 owner_base + frozen_final paths resolved by on-disk hash scan (listed in §5.2 at their leading digests). Governing packets re-verified: rev12 `1125b0a0…` (ratify relay `step3-arch-packet/…-165500` confirmed on disk), m-3 schema `9e874df8…` + contract `6e2abe40…`, §D-settlement `1fa71cb8…`. Join legs, close `…-160000`, env-locus correction `…-163000`, and the three carried dispositions (consolidated in `…-160000`) enumerated by literal path. The 3 r5-F4 fold files confirmed live. §6/§11/§12 read at the bytes. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` to rev3 + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no lock record authored, no owner action opened, no self-ratification.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews rev3 `512e9c52…` (external binding clean; lane-4 order single; manifest closed + literal incl. the ratification packet + typed edges; the amendment-vs-record split acceptable; fold manifest complete); on approve → operator hash-bound ratification → master authors `STEP-3-INTERFACE-LOCK.md` (item A) → lane 4 → lane 5 (T4). Owners held until ratify, then released. H-12 stands.
