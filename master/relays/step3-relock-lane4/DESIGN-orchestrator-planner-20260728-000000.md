## DESIGN — lane-4 plan rev2 `cc19beb2…`: folds VP r1's four gates + the required GRILL_LOCK (`step3-lane4-staffing-grill-1`). F1 ten fixture records/six legs + frozen-oracle-not-runnable-red framing; F2 the three carried obligations (N910 · env_digest parity · r7-mirror stop/reopen) as a closed checklist; F3 conforms to B13 (pair on frank + owner-fidelity + preflight) with the grill-resolved read-only + write-fence + authority ceiling; F4 H-16/H-26 before T4 + the exact five-step freeze order. Approach only — no build/lock/stand-up.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the plan CONFORMS to B13 (pair + frank + fidelity + preflight); the operator's grill answers (read-only access; "lets try" frank transport) pin open Part-F items and do NOT deviate from B13, so no operator ratification gate. Team stand-up waits on the preflight + the operator's green-light (GRILL_LOCK "still operator-owned").
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260727-230000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev2 `cc19beb2d9acd39e0ed2e175412906bd8cb8dee3fa6a7573b9877468c2e0f35c` + GRILL_LOCK `step3-lane4-staffing-grill-1` — ten fixtures, carried obligations, B13-conformant read-only team, H-16/H-26 + freeze order; on approve → detailed kickoff brief + preflight; Item A lock `cbd1893c…` preserved

## What changed vs rev1 `d79c44c1…`
Your r1 passed the separate-team + Master+VP-keeps-lock split, and B20 fired correctly. Four gates + a required GRILL, all folded:

- **F1 (count + framing).** "six fixtures" → **ten fixture records across six legs** (enumerated: `xit-gov-1`, `xit-dur-1…5`, `xit-crash-1`, `xit-inj-1`, `xit-ho-1`, `xit-op-1`). Reframed: lane 4 authors + content-addresses the immutable input/baseline artifacts and freezes the **expected-outcome oracle** — a **frozen test spec, NOT runnable RED execution**; the executable fixtures + code are T4. No amendment sought (we keep the ratified §7/§11 split).
- **F2 (carried obligations).** Added a **closed `carried_obligations` checklist** (§5): N910 honest `UNKNOWN_PROVIDER_OUTCOME`→`uncertain` in the durability/operability oracle (no complete-coverage claim); the `env_digest` byte-exact JCS preimage + duplicate-name reject + reachable non-UTF-8 reject + m-9↔m-3 observer parity under `xit-gov-1`; the r7-mirror **mandatory m-3 2a/2b check with STOP+reopen-if-yes**. Each mapped to fixture record + expected rows + owner + stop/reopen.
- **F3 (staffing/transport/authority — GRILL_LOCK §10).** Conformed to B13 instead of re-asking: a **.planner/.implementer pair on frank**, out-of-pair owner-fidelity, the T4-token preflight. The grill resolved the genuinely-open items: **READ-ONLY** workspace access (operator "only read"), team writes nothing into `master/` and delivers as **frank-relay proposals** that **Master+VP materialize + freeze**; the **authority ceiling** = author/design only within its relay channel, escalate on any `DELEGATED_DISPATCH_AUTHORITY` trigger; **master** owns the hardening-backlog append. Transport = **frank-as-courier, gated on the preflight** (operator "lets try it").
- **F4 (gates + freeze order).** T4 opens only after the re-lock **AND H-16 AND H-26**; the exact five-step freeze order (inputs → complete manifest with no placeholder/mutable slot → team review + owner-fidelity → Master+VP freeze + one durable re-lock binding both `cbd1893c…` and the manifest SHA → then T4).

## What I ask the VP to review (approach only)
- Corrected framing (ten records / frozen oracle) right?
- Carried-obligations checklist complete + correctly mapped (incl. r7-mirror stop/reopen)?
- Team envelope conforms to B13 + the GRILL_LOCK resolves the open items (read-only, write fence, authority ceiling)?
- Freeze order + H-16/H-26 exactly the ratified sequence, no surviving placeholder?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, or team stand-up. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the detailed kickoff brief; the preflight runs; the team stands up on the operator's green-light. **H-12 hard-blocks external use.**

## Verification
Plan rev2 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `cc19beb2d9acd39e0ed2e175412906bd8cb8dee3fa6a7573b9877468c2e0f35c` (supersedes rev1 `d79c44c1…`). All ten §7 fixture IDs present; the three carried obligations + H-16/H-26 + the GRILL_LOCK present. B13/Part F re-read for the settled staffing/transport/fidelity/preflight points. Item-A lock `cbd1893c…` PRESERVED (unmoved). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-LANE4-PLAN.md` to rev2 (incl. the folded GRILL_LOCK) + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no team stand-up, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev2 `cc19beb2…` + GRILL_LOCK; on approve → master writes the detailed team kickoff brief and the B13 preflight runs; the lane-4 pair stands up on the operator's green-light. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
