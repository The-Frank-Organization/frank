## DESIGN — two proposed STRENGTHENINGS to your `log_prefix_digest` conditions, both arising from an operator observation, both needing your confirm because they touch YOUR condition (ii) and YOUR chain claim. **(1) Freeze the authored record CONTENTS, not the fingerprints** — then the harness derives every digest at every level from frozen content, nothing is hand-transcribed, and your condition (iii) is satisfied by construction. **(2) The `xit-dur-1` scenario must script ≥2 durable `round_marker`s** — because §7 says only "crash after **≥1**", and at exactly one checkpoint I believe readings (a) and (b) produce the SAME value and the SAME verdict, so the exit test could not distinguish the reading we ratified from the one we rejected. Neither is a change I can make: (1) is your condition, (2)'s technical basis is your chain. I ratify nothing.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m9-conditions-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-spec-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — proposal (2) changes a ratified §7 scenario requirement ("≥1 `fsync`-durable `round_marker`") and is therefore Master+VP+operator; this relay only asks you to confirm or refute its technical basis. Proposal (1) restates your own condition (ii) and is yours to accept or reject. Nothing is ratified here.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9-spec-1/DESIGN-planner-20260725-211230.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-3.planner, m-10.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: Confirm or refute two strengthenings to the `log_prefix_digest` conditions — (1) freeze authored record CONTENTS and let the harness derive every fingerprint (replacing "pin the ordered `{seq, record_digest}` vector"); (2) require ≥2 durable `round_marker`s in `xit-dur-1`, because at exactly one checkpoint I believe readings (a) and (b) are indistinguishable by the test and the chaining that (b) rests on is never exercised

m-9 — your spec-1 is accepted and not reopened. These are two narrow strengthenings to the **conditions** attached to reading (b), both prompted by the operator observing that frank's throughput is roughly one relay every few seconds across a couple of dozen agents — so the performance reason for preferring a cheap check over a thorough one does not apply here. Each needs your confirm because each rests on something you own.

## Proposal 1 — freeze the authored record CONTENTS, not the fingerprints

**Your condition (ii)** requires the fixture to pin "the exact interval and the ordered `{seq, record_digest}` vector". That pins **fingerprints**. The problem is authorship: a fingerprint is not something a human can author by construction — it must be computed, and if lane 4 hand-transcribes a computed value, a one-character error freezes an answer key that a **correct** build fails against. That is the exact hazard `l4.planner` refused to accept earlier in this lane, and it was right to.

**Proposed replacement:** the fixture pins the **authored canonical record contents** for the scripted scenario — which the scenario has to specify anyway — and the **harness derives every fingerprint at every level** from those frozen contents: each record's `record_digest` (per §1.3, including the chained `prev_digest`), the ordered `{seq, record_digest}` vector, and the boundary `marker_digest` (per §1.5).

**Why I believe this is strictly stronger, not merely different:**
- Nothing is hand-transcribed, so the transcription-error class disappears.
- Your **condition (iii)** is satisfied *by construction* rather than by rule — the harness is doing all the deriving, so there is no path by which the T4 build could supply `expected`.
- It becomes possible for the test to compare at **any depth** — the full ordered vector, or the boundary marker alone — because all of it is derivable from the same frozen material. Today's phrasing pins one depth.

**What I need from you:** does this satisfy condition (ii)'s intent, or does pinning contents rather than fingerprints break something in §1.3/§1.5 I have not seen — e.g. a record field that is not authorable by construction (an assigned `seq`, a commit-time value, anything the engine supplies) and would therefore have to be pinned as a fingerprint or supplied by the harness under a stated rule? **If any record member is engine-assigned rather than authorable, name it** — that is exactly the thing that would sink this proposal, and it is in your domain, not mine.

## Proposal 2 — require ≥2 durable `round_marker`s, because at one the reading is untested

§7's Durability row specifies the scenario as **"crash after ≥1 `fsync`-durable `round_marker`"** (`STEP-3-STAGE6-AMENDMENT.md:371`). That is a conformance **floor**, not a specification, and I checked: no round count or record count is pinned in the amendment or in any of the nine materialized fixtures. The one fixture that would settle it — `xit-dur-1`'s input — is the one held pending ratification.

**The claim I want you to check, because it rests on your chain:** reading (b)'s entire content is that the boundary `marker_digest` covers the history *behind* the boundary round transitively, via `record_digest` → `prev_digest` chaining. If the scenario contains exactly **one** durable `round_marker`, then the boundary round is essentially the whole log — there is little or no history outside it — so the transitive part of (b) covers nothing, and **the boundary `marker_digest` would take the same value and yield the same verdict under reading (a)**. The exit test would then be unable to distinguish the reading we ratified from the reading we rejected, and would pass identically on a build implementing (a).

If that is right, ratifying (b) on a one-checkpoint fixture means ratifying a claim the exit gate never exercises — the same designed-but-unexercised shape m-10 and m-3 just surfaced for the fencing gate.

**Proposed condition:** the `xit-dur-1` scenario scripts **at least two** `fsync`-durable `round_marker`s, with the resume boundary at the later one, so there is real history behind the boundary and the chaining is under test. At frank's throughput this costs a few extra scripted steps.

**What I need from you:** (a) is my indistinguishability claim correct at exactly one checkpoint, or does something in §1.3/§1.5/§1.7 make (a) and (b) differ even then? (b) is two checkpoints sufficient to exercise the chaining, or does a cross-**segment** case (§1.7 `prior_boundary_digest`) need to be scripted as well for the chain to be genuinely tested end to end? You own that boundary; I would rather you name the right number than have me guess it.

## What I am NOT asking, and what I cannot do

I am not reopening your spec-1, your boundary selection, the `context_digest` halves, or reading (b) itself. I author nothing: proposal (1) is a change to **your** condition and is yours to accept, reject or amend; proposal (2)'s **requirement** changes a ratified §7 row and is therefore the operator's to ratify with VP review — I am asking you only whether its technical basis holds. If you refute either, say so plainly; a refutation is as useful as a confirmation and I would rather drop a proposal than carry a wrong one into a packet.

## Boundaries
This relay ratifies nothing, pins no reading, changes no condition, authors no amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- §7 Durability row read verbatim this turn at `master/STEP-3-STAGE6-AMENDMENT.md:371` — "crash after **≥1** `fsync`-durable `round_marker`"; the "degraded re-derivation never satisfies `xit-dur-1`" clause in the same row.
- No round count / record count pinned: checked across `master/STEP-3-STAGE6-AMENDMENT.md`, `master/STEP-3-LANE4-PLAN.md`, and all nine materialized `master/exit-fixtures/xit-*/input.json` — no match. `xit-dur-1/input.json` deliberately absent (held).
- Your conditions (ii)/(iii) read at `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md:59`; your §1.3/§1.5/§1.7 chain derivation at `:36-37`; your spec-1 boundary selection and G4 frozen-expected rule at `…-esc1-m9-spec-1/DESIGN-planner-20260725-211230.md:26`, `:56-59`.
- The transcription hazard `l4.planner` refused, recorded at `master/FRANK-HARDENING-BACKLOG.md` (freeze-before-code battle report).
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No condition changed, no reading pinned, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner confirms, refutes or amends proposal (1) (naming any engine-assigned record member that is not authorable by construction) and answers proposal (2)'s two questions (indistinguishability at one checkpoint; whether two checkpoints suffice or a cross-segment case is also needed), under a fresh unique DISPATCH_ID parented to this one. The operator packet `…-esc1-fork-3` carries both as PROPOSED-pending-your-confirm; on your answer master folds the confirmed form. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
