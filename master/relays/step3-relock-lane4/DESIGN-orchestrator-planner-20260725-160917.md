## DESIGN — lane-4 plan rev9 `a0009ef9…`: all three r8 findings ACCEPTED, none contested. **F1** — you were right that the durable record did not authorize the pair: the operator chose in-session but master never recorded it, so rev8 asserted an authorization B22 explicitly left open. Now closed by **B23** (agent-authored + operator-cited, §8b), superseding B21's team-shape half; **standing lesson folded: a decision is authorized when it is DURABLE, not when it is spoken.** **F2** — the read-only/authorship contradiction is real and was created by removing `relay.submit` without naming a replacement; §3 now states an exact write fence (own relays + INDEX rows ONLY; every governed artifact read-only; master alone materializes; no proxy-authored `FROM`). **F3** — accepted: shared request/verdict ids do NOT preserve the r7 guard, since the resolver takes the earliest relay sharing an id; §3 now gives **every mechanically distinct relay its own `DISPATCH_ID`** with a concrete table. §4/§5/§6/§8 remain byte-stable.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator's PAIR choice is now durable at B23 (cited, not forged); the operator alone boots the pair and hands over. No preflight or activation gate remains.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-160401.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev9 `a0009ef930a3a8f3bc4e2edf80e601c9b6c686c8ab24a3535ced8ef1a2206a52` — r8-F1/F2/F3 folded; B23 records the PAIR durably; Item A lock `cbd1893c…` preserved

## F1 — accepted; the failure was mine and it is worth naming precisely

You are right that **routing a plan does not select a team shape**, and that a planner-authored claim cannot stand in for a still-open durable record. What happened: the operator *did* choose the pair in-session — answering the exact question *"does lane 4 keep B21's nested team shape on file relays, or revert to the pair shape of VP-approved rev5?"* with **"pair is good"** — and I folded that straight into rev8 while B22's durable text still read *"Open, operator-owned."* So the plan carried an authorization its own governing record did not.

**Disposition, in your required order:**
1. the operator's exact choice is **PAIR on file relays**;
2. recorded durably at **`master/PROTOCOL-DEVIATIONS.md` → B23** — **agent-authored + operator-cited per §8b, never a forged `FROM: operator`** — which **supersedes B21's team-shape half** (B21's courier half was already superseded by B22, so B21 is now fully superseded for lane 4), and B22's open item is closed by explicit cross-reference so the two records cannot contradict;
3. plan, both void banners and the GRILL_LOCK now all cite **B23** for shape and **B22** for transport.

Folded as a standing lesson in the GRILL_LOCK: **a decision is authorized when it is durable, not when it is spoken — record it before building on it.**

## F2 — accepted; the contradiction was created by removing `relay.submit`

Exactly as you diagnosed: frank let a read-only seat author through `relay.submit` without a filesystem write surface, and rev8 removed that without naming a replacement. Under rev8's own words the pair could not create its own relay, and master writing one bearing `FROM: l4.*` would be proxy-authoring.

§3 now draws the fence at **authorship**, not at the filesystem:
- **Writable by the pair — exactly two paths:** its own relay files under `master/relays/<its own DISPATCH_ID>/`, each seat writing only relays whose `FROM` is itself; and appended `INDEX.md` rows after the `monotonic-from` boundary. This is the ordinary mechanism every prior slice pair used.
- **Read-only — everything else, enumerated:** the interface lock + all 38 constituents · every owner/frozen final and amendment · `ARCHITECTURE`/`ROADMAP`/`README`/`RECONCILE`/`PROTOCOL-DEVIATIONS`/`FRANK-HARDENING-BACKLOG` · the `domains/` trees · **and every fixture, baseline and manifest artifact, materialized or not.**
- **The pair never writes a governed artifact** — those still ride proposal envelopes; **master alone materializes** and recomputes. **The freeze discipline is unchanged**; only the authorship path is named.
- The kickoff must be validated against the seats' **actual tool configuration** before handover (F2.4), applying this plan's own live-artifact-not-prose rule.

The operator's "only read" is preserved in intent — it meant *no governed-artifact authority*, which this states exactly.

## F3 — accepted; I will not describe a shared id as satisfying the guard

Your mechanical argument holds and I verified the same behaviour in the preflight: `acceptedDispatchRoot` returns the **first** accepted relay for a dispatch id, so a gate expecting a verdict resolves to the request. rev8's `…-l4` / `…-l4-review` / `…-l4-esc<n>` shared-thread ids were a regression of the r7 guard, and rev8 compounded it by claiming they *avoided* the shared-id defect. Withdrawn.

§3 now carries a concrete table giving **every mechanically distinct relay its own `DISPATCH_ID`** — `…-l4-dispatch`, `…-l4-propose-<n>`, `…-l4-review-req-<n>`, `…-l4-review-verdict-<n>`, `…-l4-esc<n>-req`, `…-l4-esc<n>-disp`, `…-l4-return` — with `PARENT_DISPATCH_ID` set to the exact unique id of the immediate predecessor, a shared **namespace prefix** carrying grouping, and **`IN_REPLY_TO` explicitly never a gate input.**

## What I ask the VP to review
- F1 closed — is the staffing authorization now durable and consistent across B22/B23/plan/banners/GRILL_LOCK?
- F2 closed — is the write fence exact and correct, with the freeze discipline demonstrably unchanged?
- F3 closed — does every mechanically distinct relay carry a unique id with an exact predecessor edge?
- Are §4/§5/§6/§8 still byte-stable, and does anything reopen a closed r1–r7 decision or a locked byte?

## Boundaries
No kickoff, pair boot, proposal, fixture, manifest, materialization, owner-fidelity request, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or external use on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On approval this authorizes only a **fresh INERT kickoff**; the operator retains the seat-boot/handover gate. **H-12 hard-blocks external use.**

## Verification
- Plan rev9 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `a0009ef930a3a8f3bc4e2edf80e601c9b6c686c8ab24a3535ced8ef1a2206a52` (supersedes rev8 `1fc035fa…`).
- `master/PROTOCOL-DEVIATIONS.md` = `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b` (B23 added ahead of B22; B22's open item closed by cross-reference).
- Void banners re-stamped: `master/STEP-3-LANE4-KICKOFF-NESTED.md` = `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`; `master/STEP-3-LANE4-KICKOFF.md` = `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af` (its assumed shape is authorized again, but the bytes stay VOID — written against rev5, not rev9's fence/lineage).
- Residue sweep: 0 `rev8` title refs · 0 "writes nothing into the governed tree" · 0 "ONE gated leg = ONE dispatch id" · B23 cited 5×.
- Interface lock `cbd1893c…` UNMOVED. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — added B23 to `master/PROTOCOL-DEVIATIONS.md` + closed B22's open item; revised `master/STEP-3-LANE4-PLAN.md` to rev9 (§3 write fence + unique-id lineage table, header/§7/§9/§10/boundaries; §4/§5/§6/§8 byte-stable); re-stamped both void kickoff banners; this relay + one INDEX.md row. No fixtures/manifest/lock/owner/frozen byte moved, no `frank/` action, no seat boot, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev9 `a0009ef9…`; on approve → master writes the fresh **INERT** rev9 kickoff, validated against the artifacts it names **and** the seats' actual tool configuration; the operator then boots `l4.planner` + `l4.implementer`. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
