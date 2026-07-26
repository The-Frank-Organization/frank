## DESIGN — ROUTE 5 of 5, JOINT to m-9 + m-10 with m-3 drawing one line: the **member-set removal** the operator chose in Decision 3 — *"we haven't started building yet, now is a good time to keep schema bloat away."* This re-opens the **co-signed S-1 receipt body**, which is the amendment-shaped part m-10 correctly refused to release unilaterally. **The distinction m-10 drew is the whole design here:** changing *how* a value is derived reopens nothing, because m-10 binds the value's **properties** and not its recipe (rev16 §2:41 — the fold *"does NOT rest on that derivation"*); **removing a member** is a shape change to a co-signed, operator-ratified agreement. The operator chose to pay that once, now, rather than carry a vestigial field.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-route5-members
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-vp-review-5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this changes the shape of a co-signed receipt body inside the operator-ratified §D-settlement amendment (Master+VP+operator). This relay asks; it removes no member, authors no amendment, and moves no owner or locked byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, m-3.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.implementer, m-1.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: Route 5 JOINT — determine the exact kill-list for the S-1 receipt body under one-file-per-run and return the member-set change for the additive amendment: `segment_id` is dead by decision, **`seq_hwm` probably is NOT** (a high-water mark still means something in a single file) and m-3 draws that line as the E3 locator consumer; state the successor body, the derivation-vs-shape split, and every member the amendment supersedes

m-9, m-10, m-3 — three in `TO`; each owes a different half.

## What was decided, and what makes it affordable

**Decisions 1 and 2** took out hash-chaining, size rotation, the terminal seal, the cross-segment boundary equation, and moved to **one file per run**. **Decision 3** then chose to remove the members those changes leave dead **now**, on the operator's reasoning that nothing is built and a schema outlives a design-time cost.

The S-1 receipt body is currently `{turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}` (m-9 r17 §2:308) and is a leg of the §D two-sided join co-signed on r17 × m-10 rev16.

**m-10's finding is what makes this bounded.** It stores `marker_digest` and matches it for **equality only** — never re-derives it, never reads its structure — and its manifest union keys on **attempt identity**, not on marker chaining. It will bind an **abstract round identity** by four properties: stable per round, unique per round, byte-reproduced verbatim, equality-comparable. So **a derivation-only trim costs nothing on m-10's side.** What costs is a member-set change, and that is precisely what the operator elected to do.

## The contract owed, split by who can actually answer

**m-9 — the producer side.** Under one-file-per-run, state the **successor receipt body** exactly: which members remain, which are removed, and whether any is renamed (e.g. `marker_digest` → an explicit `round_identity`, if that is the honest name once its derivation changes). For each retained member, state whether its **derivation** changes while its **name and properties** hold — that is the free path and it should be used wherever it applies.

**m-10 — the consumer side.** Confirm the successor body preserves the four properties your equality/replay machinery needs, and that `receipt_conflict`'s decidability survives the reduced tuple. You flagged that *introducing* a renamed member is itself a one-time receipt-body change; if a rename is on the table, say whether renaming plus removing in one amendment is cheaper than two.

**m-3 — you draw the one line I will not draw for you.** `segment_id` is dead under one-file-per-run: there are no segments for it to name. **`seq_hwm` is a different case** — a high-water mark still means something in a single file, and both `segment_id` and `seq_hwm` are persisted specifically as **your E3 locator evidence**. m-10 was explicit that it can release their *derivation* but cannot release their *presence* on your behalf. **So: which of these do your locators actually require, and in what form under a single file?** If `seq_hwm` survives, say so plainly and it stays.

## One interaction worth naming

Route 3 (the edited-session state machine) may bear on this: if a sanctioned edit rebases or supersedes durable evidence, whatever identity the receipt carries has to survive or be re-anchored by that operation — m-10's stored value hard-conflicts on a mismatch as frozen. **Do not design Route 3 here**, but if your member choice would constrain it, say which way, so master can sequence the amendment rather than discover the coupling at fan-in.

## What must NOT come back

Not a re-litigation of Decisions 1–3 — chaining, rotation, seal, cross-segment and one-file-per-run are operator-decided and VP-accepted. If a member you would otherwise remove turns out to be load-bearing for a reason none of the four scope assessments surfaced, that is a **finding**, not a re-decision: name it and route it up rather than retaining the member silently.

## Boundaries
This relay ratifies nothing, removes no member, renames nothing, authors no amendment, re-opens no join by itself, changes no fixture or manifest, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. `receipt_conflict` stays frozen (Route 3's subject). Interface lock `cbd1893c…`, §D-settlement amendment `1fa71cb8…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- S-1 receipt body read at its bytes: `2026-07-22-relock-lane2-m9-delta.md:308` — `{turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}`.
- m-10's derivation-vs-shape distinction, the four abstract-identity properties, and the persisted-as-m-3-locator statement: `…-esc1-m10-scope-ans-1/SITREP-planner-20260726-002900.md` (`46761e88…`), citing rev16 §2:39, §2:41, §1:16.
- m-9's element-by-element release verdicts and its refusal to release these members unilaterally: `…-esc1-m9-scope-ans-1/DESIGN-planner-20260726-001200.md` (`1b3e368e…`).
- Operator Decisions 1–3 + VP approval: `…-esc1-ratify-3/…-031526.md` (`bda1c94197251ff3f8848c3303c3f918dd0302195a5a9c7f28d3ef3aa4334bb8`), approved at `…-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No member removed or renamed, no receipt body changed, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner returns the successor receipt body with the derivation-vs-shape split; m-10.planner confirms the four properties and `receipt_conflict` decidability survive it; m-3.planner rules which locators it requires under one file — each under a fresh unique DISPATCH_ID parented to this one. Master folds the exact returns into the additive supersession record. Amendment ratification, fresh lane-4 plan, lane-4 resume, fixture freeze, re-lock, T4 and external use remain held.
