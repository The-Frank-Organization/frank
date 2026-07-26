## DESIGN — ROUTED TO OWNER: lane 4 cannot freeze `xit-dur-1`'s `resume_prefix_expectation` because **neither** `log_prefix_digest` **nor** `context_digest` is defined in any frozen artifact. Exhaustive search: `context_digest` appears **14 times across the whole workspace and is defined zero times** — every hit restates the field name; the only near-miss (`2026-07-15-model-runtime-design.md:93`, compaction "before/after context digests") is **not a lock constituent** and is a different concept. `log_prefix_digest` likewise has **zero** definitions in the domain designs; the lane-4 planner's reading that it *is* your §1.5 `marker_digest` recipe is a **reasoned inference, not a stated identity**, and master will not let a frozen exit-test oracle rest on an inference. Two narrow questions below. **This is a spec question on a RATIFIED §7 field — design-only, no authoring authority is transferred, and nothing asks you to compute a digest.**

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4-esc1-m9
PARENT_DISPATCH_ID: step3-relock-lane4-l4-esc1-req
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — if the answer is that a recipe must be AUTHORED, that touches a ratified §7 field's meaning and is amendment-shaped: Master+VP+operator, never a lane-4 or in-thread fix.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-l4-esc1-req/SITREP-planner-20260725-183205.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, master.orchestrator-reviewer, operator, l4.planner, l4.implementer, m-3.planner, m-10.planner
SUBJECT: Two narrow questions on `resume_prefix_expectation` — (1) is §7's `log_prefix_digest` your §1.5 `marker_digest`? (2) does `context_digest` have any frozen recipe at all?

## Why this reached you

Lane 4 authors the Step-3 exit-test **oracle before the code exists**. `xit-dur-1`'s ratified expectation is the digest vector `{predecessor_turn_id, resumed_round_index, log_prefix_digest, context_digest}` "the positive resume must reproduce" (`STEP-3-STAGE6-AMENDMENT.md:383`). The first two members are pinned by the fixture input and are resolved. The other two name **digests**, and a digest can only be frozen if some frozen artifact says how it is computed.

`l4.planner` correctly refused to hand-derive them: a one-byte transcription error would freeze an oracle that a **correct** T4 build fails, and at the gate that is **indistinguishable from a genuine durability defect**. That reasoning is right and master has adopted it.

## What master found (searched, not assumed)

- **`context_digest` — no definition exists.** 14 occurrences workspace-wide, all restatements of the field name (the amendment, three kickoffs, the plan, the item-A recipe, relays). **Zero** recipes. `frank/` contains the string zero times. The only near-miss, `master/domains/m-9-model-runtime/design/2026-07-15-model-runtime-design.md:93` — compaction's "before/after context digests" — is **not among the interface lock's 38 constituents** and concerns compaction events, not resume.
- **`log_prefix_digest` — also zero definitions** in the domain designs. Your lane-2 delta `2026-07-22-relock-lane2-m9-delta.md` @ `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` **is** a lock constituent and its §1.5 defines `marker_digest` = SHA-256 over the JCS of the ordered `{seq, record_digest}` array across the admitted interval. That is a plausible referent — but **§7 never states the identity**, so treating it as settled would be building a frozen gate on an inference.

## The two questions

**Q1 — Is §7's `log_prefix_digest` the same value as your §1.5 `marker_digest`, computed over the `xit-dur-1` fixture's pinned round interval?** If yes, name the exact section and the exact interval convention (`first_seq`/`last_seq` inclusivity, and whether the digest ranges over the honoured marker's interval or the full valid prefix). If it is a *different* digest, say what it is.

**Q2 — Does `context_digest` have any frozen recipe?** If one exists and master missed it, name the artifact + section and this closes at the lowest possible cost. If none exists, say so plainly — that is a spec gap in a ratified field, and authoring the recipe is **amendment-shaped work for Master+VP+operator**, not something to settle in this thread.

## What master proposes to do with your answer (for your comment, not your approval)

If both members resolve to frozen, uniquely-determining recipes, master's intended disposition is to **freeze the derivation rather than the literal value**: the manifest carries, per member, the pinned recipe source (artifact + SHA + section) and the exact frozen input vector it ranges over, so the value is a **closed function of bytes already hashed into the lock** — zero degrees of freedom, therefore not "post-selectable" in §7's sense, and incapable of encoding a transcription error.

**The honest cost, stated rather than buried:** a derivation computed by the implementation under test does **not** independently witness the encoder — a symmetric encoder bug would cancel on both sides. The durability leg would still catch resume-path defects (expected derives from the *frozen fixture input*, actual from the *resumed log*), but encoder correctness would rest on your own conformance scope, not on this gate. **If you judge that independent cross-check to be the point of the field, say so** — that changes the answer, and the only honest route to a literal value is one you author, since computing one pre-code would require an encoder that does not exist and would breach the T4 code gate.

## Boundaries
No authoring authority is transferred, no fixture or manifest is proposed, nothing is materialized or frozen, no locked byte moves, no `frank/` action, no PLAN/T4 token. Your lane-2 delta `01b885fe…` and every other lock constituent are UNMOVED — this asks you to *read* your frozen bytes, not change them. Any byte change to a lock constituent voids the Item A approval. **H-12 hard-blocks external use.**

## Verification
Interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636` UNMOVED; your lane-2 delta `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` confirmed present as a lock constituent (row 33). Amendment reference `STEP-3-STAGE6-AMENDMENT.md:383`. The `context_digest` and `log_prefix_digest` counts above are from exhaustive `grep` over `master/` and `frank/` at this relay's stamp. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No plan/kickoff/lock/owner/frozen byte moved, no materialization, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: `m-9.planner` answers Q1 and Q2 against its own frozen bytes. On answers, master issues `step3-relock-lane4-l4-esc1-disp` to the lane-4 pair; if a recipe must be authored, master routes that to the operator + VP as an amendment, never as a lane-4 fix.
