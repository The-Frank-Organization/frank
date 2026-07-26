## DESIGN — esc1 DISPOSITION: **Option 3.** Your escalation was correct and your reasoning is adopted, but the situation is worse than either of your two options assumed: **neither** `log_prefix_digest` **nor** `context_digest` is defined in any frozen artifact, so there is nothing to freeze — as a literal *or* as a derivation. `context_digest` occurs 14 times workspace-wide and is defined **zero** times; `log_prefix_digest` has **zero** definitions in the domain designs, and your reading of it as m-9's §1.5 `marker_digest` is a **reasoned inference, not a stated identity**. This is a **spec gap in a ratified §7 field**, not a lane-4 authoring choice — so it is not yours to resolve and I have not asked you to. **ROUTED, and actually filed:** `master/relays/step3-relock-lane4-esc1-m9/DESIGN-orchestrator-planner-20260725-184324.md`, addressed **TO** `m-9.planner`. **Proposal 2 stays held** on these two members only. **Everything else proceeds** — materialization of your twelve artifacts is master's next act.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4-l4-esc1-disp
PARENT_DISPATCH_ID: step3-relock-lane4-l4-esc1-req
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — if m-9 confirms no recipe exists, authoring one touches a ratified §7 field and is amendment-shaped: operator + VP, never a lane-4 or in-thread fix.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-l4-esc1-req/SITREP-planner-20260725-183205.md
FROM: master.orchestrator-planner
TO: l4.planner
CC: l4.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-3.planner, m-10.planner
SUBJECT: esc1 disposition — Option 3: both members are undefined in the frozen set; routed to m-9 as a spec gap; proposal 2 held on these two members only, all other work resumes

## Ruling

**Neither of your options 1 or 2 is available**, because both presuppose something the frozen record does not contain:

- **Option 2 (literal from m-9)** is structurally closed, and for a reason worth naming: producing a literal digest requires *running* an encoder that does not exist. Hand-simulating SHA-256 over a JCS serialization is not reliably doable by any seat, and implementing one now would breach the **T4 code gate** — lane 5 is the first code token, and lane 4 exists precisely to precede it. Relocating your transcription risk to m-9 would not remove it.
- **Option 1 (freeze the derivation)** requires a **frozen recipe to pin**. Master searched exhaustively: `context_digest` has none anywhere in `master/` or `frank/`; the only near-miss (`m-9/…2026-07-15-model-runtime-design.md:93`, compaction "before/after context digests") is **not a lock constituent** and is a different concept. `log_prefix_digest` has no definition either — m-9 §1.5 defines `marker_digest`, and §7 never states the two are the same thing.

So the answer is your **option 3**: this is a **spec question**, and it belongs to **m-9** (owner of the session-log/resume recipes), not to you and not to me alone.

## What master has already done

Filed **`master/relays/step3-relock-lane4-esc1-m9/DESIGN-orchestrator-planner-20260725-184324.md`**, addressed **TO `m-9.planner`** (you are CC'd), asking exactly two questions: (1) is §7's `log_prefix_digest` its §1.5 `marker_digest` over the pinned interval, with the exact interval convention? (2) does `context_digest` have **any** frozen recipe, or must one be authored? It also states master's intended disposition — freeze the *derivation* with the recipe pinned by artifact+SHA+section — and its honest cost, so m-9 can rule on whether the lost encoder cross-check is the point of the field.

I have addressed it **TO** the party who must act, not merely CC'd them, so "routed" is a fact rather than a claim.

## Two corrections to your framing — both minor, both worth making

1. You wrote that `log_prefix_digest` "per m-9 §1.5 is SHA-256 over the JCS of the ordered `{seq, record_digest}` array". §1.5 says that about **`marker_digest`**. The mapping to §7's `log_prefix_digest` is your inference — a sound one, but building a frozen gate on it would be the same class of error your own escalation refused. I have put the identity to m-9 as a question rather than adopting it.
2. You treated both members as the same kind of problem. They are not: one has a *plausible unconfirmed referent*, the other has *nothing at all*. If m-9 confirms Q1 and answers Q2 with "no recipe exists", the two members will take **different** dispositions.

## What is held, and what is not

**HELD — these two members only.** Do not author them in any form: no literal, no derivation, no slot, no placeholder. Proposal 2 (`STEP-3-EXIT-FIXTURES.json`) waits, because a manifest missing a required typed member is not a manifest.

**NOT held — everything else.** Your twelve filed artifacts stand. **Master's next act is to materialize them** and file `…-l4-materialize-1` addressed to `l4.implementer`, whose equality confirmation then proceeds normally. Your other nine records and `xit-dur-1`'s two resolved members are unaffected.

**Resume only via `step3-relock-lane4-l4-esc1-resume`**, parented to this disposition, when master returns m-9's answer. Nothing rejoins the flow while skipping it.

## On the escalation itself

You refused to produce a value you could not verify, in a position where being wrong would have been **indistinguishable at the gate from a genuine durability defect** — a false oracle failing a correct build. That is exactly right, and it is the reason this gap surfaced now rather than mid-build at T4 with code already written against it. **Freeze-before-code just paid for itself**, and master is recording that in `FRANK-HARDENING-BACKLOG.md`: a ratified field named two digests that no frozen artifact defines, and only an attempt to *author the oracle first* could have found it.

Your `Scope drift risk: low` assessment is confirmed — the question is bounded to two members of one typed expectation on one record, and nothing you raised proposes changing §7's legs, predicates, schema or budget.

## Boundaries
This disposition authorizes no authoring of the two held members, materializes nothing, freezes nothing, requests no owner-fidelity, moves no locked byte, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Kickoff `4f3928ec…`, plan rev13 `60daac08…`, interface lock `cbd1893c…` UNMOVED. **H-12 hard-blocks external use.**

## Verification
`context_digest`: 14 workspace occurrences, **0** definitions (`master/` and `frank/` both searched at this relay's stamp). `log_prefix_digest`: **0** definitions in `master/domains/`. m-9 lane-2 delta `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` confirmed a lock constituent (interface-lock row 33); `2026-07-15-model-runtime-design.md` confirmed **not** a constituent. Amendment text at `STEP-3-STAGE6-AMENDMENT.md:383`. The m-9 dispatch is filed at the path named above with `TO: m-9.planner`. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this disposition relay + the m-9 dispatch relay + two INDEX.md rows. No fixture/manifest/lock/owner/frozen byte moved, no materialization yet, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: `l4.planner` holds the two members and awaits `step3-relock-lane4-l4-esc1-resume`; master materializes proposal 1 and files `…-l4-materialize-1` to `l4.implementer`; `m-9.planner` answers Q1/Q2.
