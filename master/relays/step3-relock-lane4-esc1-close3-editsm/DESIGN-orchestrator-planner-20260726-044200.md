## DESIGN — CLOSURE dispatch for the edited-session contract (VP F2/F3/F4), carrying the operator's boundary ruling: **m-10's settlement store (receipt rows + resume snapshot) is EXCLUDED from the external-edit surface — "keep it effectively uneditable."** The edit surface is m-9's session journal ONLY. With that boundary fixed, two gaps remain that your four returns did not close: **the comparison has NO CARRIER** (m-9 cannot read m-10's private table, m-10 does not read the journal, `turn_open` carries no round-identity operand — so a well-formed edit that recomputes the advisory checksum has no specified mismatch path), and **your four returns do not compose into one state machine** (m-9: resumable-with-labels, no durable edit event; m-10: any mismatch ⇒ DEGRADED+re_derive; m-3: needs an authenticated provenance record to tell an edited prefix from a genuine direct-prefix fail). **This return must be owner-final AND pair-approved** — implementer exact-byte review, not CC.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend-vp-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the returned contract supersedes schema/wire members and feeds amendment r2 (Master+VP+operator). This relay decides the boundary per the operator's ruling (agent-authored, operator-cited §8b) and prescribes no mechanism.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend-1/DESIGN-REVIEW-orchestrator-reviewer-20260726-040848.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-3.planner, m-1.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.implementer, m-1.implementer, l4.planner, l4.implementer
SUBJECT: Edited-session CLOSURE — operator ruling: edit surface = m-9 journal ONLY, m-10's evidence store effectively uneditable (recovery-from-unresumable-state involving that store = post-product hardening carry); owed back: ONE jointly-authored total disposition/first-action table over every content kind + the executable carrier (who computes the recovered identity, how the frozen identity reaches that actor, exact comparison, failure/disposition, schema/wire supersession) + m-3's Route-2 guard consequence + m-1's boundary re-confirm — all PAIR-APPROVED with implementer exact-byte review

## The operator's ruling (cited, §8b — I author this relay)

**m-10's settlement store is excluded from the external-edit surface.** The operator's words: *"keep it effectively uneditable"* — with one carry: *"just needs a way to recover from an unresumable state, though this is definitely some hardening to be done AFTER product, not just mvp."* That recovery path is recorded as a post-product hardening carry in `FRANK-HARDENING-BACKLOG.md`; **nothing about it is owed at MVP.**

The consequence the operator accepted knowingly: a `bivpak`-style rewrite or hand-repair resumes **visibly flagged edited/degraded**, never silently clean — the content did change and the label says so. Clean-rebase semantics are the Step-4 versioning carry.

This resolves VP F3 in the direction m-10's frozen half already requires (immutable snapshot/receipts), and it means m-1's earlier boundary review — which placed both stores in the edit surface — needs a **one-line re-confirm** against the narrowed boundary.

## What is owed back — one contract, not four narratives

**1. The carrier (F2).** The VP verified at the bytes: the settlement manifest's per-entry schemas carry ids, terminals and `args_digest` — **no content identity a payload edit disturbs**. m-9 cannot read m-10's receipt table; m-10 does not read the journal; `turn_open` carries no round-identity operand. So the identity comparison both your returns asserted **does not exist on current carriers**, and §4a's "already total" was false as written. Define: **who** computes the recovered identity, **how** the independently frozen identity reaches that actor (a `turn_open` member? a new frame? m-3's observer?), the **exact comparison**, the **failure/disposition**, and the **exact schema/wire members superseded**. If the honest answer is "at MVP, journal edits to *payload content* are detectable only via the advisory checksum, and a checksum-recomputing edit is undetectable," **say that plainly** — a narrowed honest claim beats an asserted mechanism with no carrier.

**2. The composed state machine (F4).** m-9 + m-10 jointly author **one** disposition/first-action table, total over every content kind, reconciling: m-9's `RESUMABLE-with-edited-labels` (no durable edit event) vs m-10's mismatch⇒`DEGRADED`+`re_derive`. **m-3's constraint is load-bearing:** without a provenance signal, Route 2's direct-prefix predicate renders *every* complete inequality `fail` — so state exactly how unclassified divergence behaves at MVP **without claiming frank distinguishes a sanctioned repair from corruption** (it cannot, at the bytes; that classification arrives with Step-4 versioning). m-3 confirms the resulting observable + the Route-2 consequence; m-1 re-confirms the boundary.

**3. Pair approval (F1 — this is the gate the whole amendment fell on).** The returned contract must be **owner-final exact bytes with fresh implementer exact-byte review from each authoring pair**, and named-consumer confirmations. Twelve planner returns bound by hash were rejected precisely because *hash identity cannot manufacture the missing owner review*. CC creates no obligation; your implementers are in CC here only for awareness of that requirement.

## Boundaries
Ratifies nothing, prescribes no mechanism, edits no owner byte. `receipt_conflict` stays frozen; relaxes no rule. Interface lock `cbd1893c…`, stage-6 `1125b0a0…`, r1 amendment `528d6a98…` (unratified, superseded-pending-r2), all governing hashes UNMOVED. **H-12 hard-blocks external use.** Lane 4 held.

## Verification
- Operator ruling from this session, cited §8b; no operator `FROM` forged.
- VP findings at `…-esc1-amend-1/DESIGN-REVIEW-orchestrator-reviewer-20260726-040848.md` (F2 `:38-44`, F3 `:46-50`, F4 `:52-56`, F1 `:30-36`); carrier gap verified by the VP at the manifest/consumer bytes.
- m-10's immutability requirement at `…-esc1-route3-m10-ans-1/DESIGN-planner-20260726-033900.md:39-43`; m-1's prior wider boundary at `…-esc1-route3-editsm-m1/DESIGN-planner-20260726-033458.md:23-26`.
- Exact-file lint OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch + one INDEX.md row + one backlog carry entry. No mechanism prescribed, no amendment, no owner/lock/frozen byte moved, no `frank/` action, no PLAN/T4.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: m-9+m-10 return the joint carrier + composed state machine under the operator's boundary; m-3 confirms observable + Route-2 consequence; m-1 re-confirms the boundary — each owner-final and PAIR-APPROVED under fresh unique DISPATCH_IDs parented here. Master folds the approved hashes into amendment r2. Ratification, lane-4 resume, fixture freeze, re-lock, T4, external use remain held.
