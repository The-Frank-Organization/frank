## DESIGN — ROUTE 3 of 5, the largest: the **edited-session state machine**. The operator has made external session editing a REQUIREMENT (Decision 5) — a session that cannot be repaired is a data-loss event, and `bivpak` legitimately rewrites session content. **I do NOT prescribe a mechanism, because I already tried and the VP blocked it:** I named `receipt_conflict` as the gate to relax, which is the wrong seam — an edited *old* log emits no second receipt at all, so relaxing it would weaken a live-ingress rule for no benefit. **`receipt_conflict` stays frozen unless your derivation proves it implicated.** And the constraint that must govern the whole design: *"'label, never gate' cannot silently promote edited bytes as prior provider/tool truth across the frozen evidence-AND-current-presence invariant"* — my rule would have had frank present edited content as the provider's actual output, which is the fabrication class m-9's own survey faults deepagents for.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-route3-editsm
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-vp-review-5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the return defines new trust/disposition semantics over frozen contracts and will name members an additive amendment supersedes (Master+VP+operator). This relay asks; it prescribes no mechanism, authors no contract, and moves no owner or locked byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-3.planner, m-1.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.implementer, m-1.implementer, m-8.planner, l4.planner, l4.implementer
SUBJECT: Route 3 JOINT — define a TOTAL edited-session state machine: detection · labelling · trust classification of edited provider/tool content · reconciliation with the immutable settlement snapshot and prior receipt identity · resulting `resume_disposition` and first action · whether a sanctioned edit rebases or supersedes durable evidence; naming any exact receipt/manifest/trust/disposition/wire member an additive amendment must supersede. `receipt_conflict` stays frozen unless proven implicated

m-9, m-10, m-3, m-1 — four seats in `TO` because four halves are owed and CC creates no obligation. Roles below.

## The requirement (operator, Decision 5) — and why it is a product property

External session editing must be **permitted**, **honestly labelled**, and **not by itself a bar to resume**. Two drivers:
- **Repairability.** A session can become unresumable through no fault of the user — the precedent is malformed thinking blocks rendering Claude Code sessions unopenable, recovered by hand-editing the session file. **A session that cannot be repaired is a data-loss event**, and repairability must therefore degrade to "a text editor works", because the emergency case is exactly the one where frank cannot parse the file and so cannot record its own repair.
- **Third-party tooling.** `bivpak` packs a session, rewrites paths inside its history, and reopens it on another machine. That is a legitimate wholesale rewrite. *(It is also what killed hash-chaining: chaining treats change as corruption; the requirement needs change treated as history.)*

## What I got wrong, so you do not inherit it

I converted this direction into a prescribed mechanism — *"m-10's `receipt_conflict` must not hard-fail an edited session"* — without an owner derivation. The VP blocked it on the frozen seam: `receipt_conflict` is evaluated when m-10 receives a **second same-key live `content_ready` frame** whose tuple differs from the first committed one; it is the exact complement of duplicate equivalence and first-committed stands. **An edited historical log need emit no second receipt at all**, so relaxing that rule does not make the session resumable — it weakens a different live-ingress totality/detector rule. A historical edit is reconciled **after `turn_open`**, where missing or corrupt content becomes `content_lost`/`DEGRADED` and is never trusted.

**So `receipt_conflict` stays frozen** unless your derivation proves it genuinely implicated — and if it does, say so with the derivation rather than by analogy to my error.

**The governing constraint on the whole design** (VP, verbatim): *"'Label, never gate' cannot silently promote edited bytes as prior provider/tool truth across the frozen evidence-AND-current-presence invariant."* My formulation would have had frank present **edited content as the provider's or tool's actual output** — the fabrication class m-9's own §9 survey faults deepagents for (synthesising a tool result for a call it cannot account for). Whatever you design, **an edited byte must not silently inherit the trust status of the original.**

## The contract owed — TOTAL, not a patch

1. **Detection** — how an edited prefix is detected at all, and at which phase.
2. **Labelling** — how the edit is surfaced, and to whom (operator surface, worker, evidence).
3. **Trust classification** — is edited provider/tool content **trusted**, **untrusted-but-model-visible**, or **degraded**? This is the heart of it, and it must be total over the content kinds, not left to a default.
4. **Reconciliation** — with the **immutable settlement snapshot** and with **prior receipt identity**, both of which describe a prefix that no longer exists byte-for-byte.
5. **Disposition** — which `resume_disposition` results, and what the first action is.
6. **Rebase-or-supersede** — whether a **sanctioned** edit (a tool acting deliberately, as opposed to an emergency hand-repair) rebases or supersedes any durable evidence, and what records that act.
7. **The supersession list** — name any **exact** receipt, manifest, trust, disposition or wire members an additive amendment must supersede. Master carries them; I will not infer them.

## Ownership split

- **m-9 + m-10 — joint authors.** m-9 owns the session-content journal, resume assembly and the `content_lost`/degrade path; m-10 owns the settlement snapshot, receipt identity and admission/disposition. Items 1–6 are yours jointly; neither half is decidable alone.
- **m-3 — the observable/evidence consequence.** What an edited session means for the evidence ladder and the E3 predicates: what is observable, what is assertable, and whether any exit-gate claim narrows as a result. You have already narrowed one claim in this escalation; say if another follows.
- **m-1 — the at-rest / provenance boundary review.** Whether permitting external edits disturbs store isolation, the seat-stamped provenance model, or the deferred content-addressing seam. You deferred verifiability-against-the-courier; this requirement leans on that deferral, and I would rather you confirm that reading than have me assume it.

## Two forward constraints to design against, not to solve

- **Any digest over rewritten bytes breaks**, not just a chain. Per-record checksums go stale per edited record; m-10's stored round identity mismatches. The **direction of record** (`ROADMAP.md`, Step-4+ carry, explicitly pending) is either path-independent canonicalisation or an explicit authorised rebase across both stores. Nothing there is decided — do not treat it as a constraint you must satisfy, only as the shape the future wants.
- **Git-like session versioning is the recorded Step-4 direction** (own implementation, git not a dependency). If your design makes that materially harder or easier, say so.

## Boundaries
This relay ratifies nothing, prescribes no mechanism, authors no contract, changes no fixture or manifest, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. **`receipt_conflict` and every current lock constituent stay frozen.** Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, §D-settlement amendment `1fa71cb8…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-1 design, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- Operator Decision 5 + VP approval of the corrected routing: `…-esc1-ratify-3/…-031526.md` (`bda1c94197251ff3f8848c3303c3f918dd0302195a5a9c7f28d3ef3aa4334bb8`), approved at `…-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md`.
- The blocked prescription and its reasoning, read at the VP's bytes: `…-esc1-ratify-1/RECONCILE-orchestrator-reviewer-20260726-025248.md` (F2), incl. the `receipt_conflict` seam description and the promote-edited-bytes constraint.
- The deepagents fabrication precedent in m-9's own survey: `2026-07-19-mvp-full-worker.md` §9 (`patch_tool_calls` synthesising a cancelled tool result).
- The pending Step-4 carry + the explicit non-binding status: `ROADMAP.md` (`315c3e7c411e9efb2c82b5ed5bdd53de02d44388590c96df6b56977f8c475209`).
- m-1's deferral of verifiability-against-the-courier: `2026-06-28-v3-trust-identity-design.md:191`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No mechanism prescribed, no contract authored, no frozen rule relaxed, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner + m-10.planner return the joint total state machine (items 1–7) with m-3's observable/evidence consequence and m-1's boundary review, under fresh unique DISPATCH_IDs parented to this one, naming any exact member an additive amendment must supersede. Master folds the exact returns into the additive supersession record. Amendment ratification, fresh lane-4 plan, lane-4 resume, fixture freeze, re-lock, T4 and external use remain held.
