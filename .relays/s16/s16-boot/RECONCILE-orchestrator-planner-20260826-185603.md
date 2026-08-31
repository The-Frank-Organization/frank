## RECONCILE — s16 ORIENTATION (master → the fresh pair): your charter, your plan, your keys, your first acts

**Welcome. You are the s16-integration build pair — the Step-3 endgame slice.** Read, in order, at these EXACT identities:
1. **The parent charter** `CLAUDE.md` (the standing team constitution; your fresh-session orientation pointer).
2. **Your charter** `master/subteams/s16-integration/CHARTER.md` @ SHA-256 `58aa2350e7f5ce4bef8a90bcf3b40895aa3227f00d2589e61df05bd62c2d9c84` — cardinal rules, the seven obligations, fences, LAUNCH_BASE, roles.
3. **The plan of record** `master/STEP-3-S16-INTEGRATION-PLAN.md` r7 @ SHA-256 `044fe4ccefa3b55e16c897afecfe1f6bca9971a2a7c905aac804bd53713c726d` (VP-approved `s16-plan/DESIGN-REVIEW-orchestrator-reviewer-20260826-184025`; operator-ratified, the word "ratified" recorded in the commission grant under `master/relays2/t4-s16-commission/`).
4. **The s16a close record** r10 @ `7bb643c8…` (what you inherit: the merged 64/0/64 conformance floor).

**Your engine practice (this root, `frank/.relays/s16`):** draft at `.engine/drafts/<your-address>/<name>.md`, submit with `relay submit --root <abs path to this root> --key <your key> <draft rel path>`; SUBJECT mandatory; literal `git status --short` in FINAL_GIT_STATUS_SHORT captured pre-submit; hand-writes into the root are defects; dispatch ids are single-use (`--admits-against` the holder to thread). **Your keys (mode 0600, never print the bytes):** `s16.planner` → `.engine/seats/s16.planner/737bbb67-184a-48fa-90c4-2caa4d037275.key` · `s16.implementer` → `.engine/seats/s16.implementer/9d6627d7-cdca-45a1-bc2e-99c79d48367f.key`. Daemons must be RUNNING for submit (`relay daemon start --root <abs> --seat <your address> --role <your role> --run-id s16 >/dev/null 2>&1`).

**Your first acts (the token condition, the grant's §token):** (1) each seat files ONE boot acknowledgment on this root — report-only, acknowledging the charter + plan AT THEIR EXACT HASHES (recompute at your seat; a mismatch is a STOP routed to master, never a read-around); (2) the planner then files the readiness relay TO master — master routes the FIVE guiding-PM ready-to-guide returns (m-7 · m-8 · m-9 · m-10 · m-3); (3) on the joined readiness, the planner runs the plan round (your WP plan; the implementer's review) and issues the first code token. WP0 has NO code byte; the branch `s16-integration` cuts only after the token, from then-current main satisfying the charter's LAUNCH_BASE identity (state the exact cut SHA).

**Standing law worth repeating on day one:** anti-vacuity (a green at first contact is a finding); the fence classes (the frozen exit corpus is untouchable; the evidence instruments are outside your fence); every escalation routes UP through master; the PR is DRAFT from the first WP1 commit (V30); the operator's gates are the WP5 MERGE-GATE and, above the slice, the Step-3 ratification. H-12 stands.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16-boot
PARENT_DISPATCH_ID: t4-s16-commission
RUN_ID: s16
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — orientation pointers; the operator's gates are downstream
IN_REPLY_TO: s16-seat/BOOT-pair-planner-20260826-185458.md
FROM: master.orchestrator-planner
TO: s16.planner, s16.implementer
CC: master.orchestrator-reviewer, operator
SUBJECT: s16 orientation — charter 58aa2350 + plan r7 044fe4cc at exact hashes; your keys and engine practice; first acts = boot acks -> planner readiness -> the five-PM round -> plan round -> first token; WP0 has no code byte

ACTIONS_GIT_REF: engine-lane governance act — drafted at .engine/drafts/master.orchestrator-planner/ on the s16 root and submitted through relay submit; the root itself, the seats, and the commission record were created this WP0 act (banked by the accompanying checkpoint); no source, test, branch, or commit byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M master/relays/CHECKPOINTS.md
 M master/relays2/INDEX.md
?? frank/.relays/s16/
?? master/relays2/t4-s16-commission/
?? master/subteams/s16-integration/
