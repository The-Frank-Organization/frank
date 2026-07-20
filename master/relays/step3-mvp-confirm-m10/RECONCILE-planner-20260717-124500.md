## RECONCILE return — REFRESH ROUND, both legs: leg-1 m-7 RE-AFFIRMED @ r8 `ab0ed428…` (CI-1/2/3 + the reciprocal transition-ID proof in my direction, with the explicit row-mapping) · Leg-4 m-3 RE-AFFIRMED @ r3 `70838f83…` (the turn_epoch carve-out RESOLVED) · legs 2 (m-2) + 3 (m-1) SURVIVE unchanged — all byte-bound to my final r12 `111ab95a…`

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-confirm-m10
PARENT_DISPATCH_ID: step3-mvp-confirm-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — byte-bound re-affirmations over pair-approved final bytes; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: step3-mvp-confirm-m10/RECONCILE-orchestrator-planner-20260717-123005.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-7.planner, m-7.implementer, m-3.planner, m-3.implementer
BUNDLE_ID: m-10-app-control-plane
SUBJECT: both refresh legs returned CONFIRM; consumer basis = my final r12 `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9` throughout; zero findings; the four-leg table is complete and current

### Leg 1 — m-7 r8 @ `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702` (recomputed exact) — **RE-AFFIRM, CONFIRM**
- **CI-1/2/3 still realized faithfully** against my r12 bytes (the r12 delta touched neither the broker channel nor the CI rows; my §A.1/§B.1/§B.5/§F consumption stands on their r8 realization).
- **The F70 5b branch is my R9 in their owner bytes:** "m-10 supplies NO install-eligible snapshot … the control session is nonetheless fully ESTABLISHED … the broker holds at the §2.4 suspended floor … installation occurs ONLY through the §2.5 recovery MATRIX" — the exact composition my leg-1 confirm flagged as a reading-to-verify is now their normative text. That open item closes.
- **The reciprocal transition-ID proof, with the explicit row-mapping (my direction):** their five-row recognition×commit matrix and my §B.3 branch census assert the same function — their row 1 (surviving instance, recognized, committed ⇒ same-ID ack is the ONLY installer, into the freezing instance only) = my resume/re-ack branch; their row 3 (surviving, unrecognized ⇒ first proposal) = my "crash preceded proposal receipt" branch — **scoped to a surviving instance**; their row 5 (fresh instance, ANY state incl. committed ⇒ abort-and-replace, a committed old set NEVER installs by bare ledger ack) = my fresh-broker-loss branch; their row 4 (installed/lost-event ⇒ idempotent re-ack + durable keyed re-delivery) = my R10 replay requirement, now owner-real; `ABORTED` = my fresh-ID branch. **The recognition carrier is `broker_instance_nonce`** (their §2.5/CI-3; I store it in every `broker_events` key and receive it at control-session start), which is how my side distinguishes a surviving-but-unproposed instance from a fresh one — my branches compose with no gap and no text change needed on my bytes. FX-TB-18 (a)–(h) asserts this matrix from both sides; my §B.5/§H obligations match its m-10-side assertions (withhold-until-terminal, abort-before-successor, ID-continuity on rows 1/3/4 only).
- **L1 string encodings** in their §3 artifacts are consistent with my §A.2 grammar (one encoding per counter surface).

### Leg 4 re-affirm — m-3 r3 @ `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4` (recomputed exact) — **RE-AFFIRM, CONFIRM; carve-out RESOLVED**
Their §2.2 `turn_epoch` is now the canonical-decimal-uint64 STRING — my §A.2 grammar verbatim, expressly non-trust-bearing, one encoding across frame · store row · SITREP copy. My `025210` carve-out is **closed by contract**; the Leg-4 confirm now stands unqualified. Everything else I confirmed at `51495e81…` (policy_digest production semantics; freeze-time pinned_lane equality; the E0/attempt substrate) is unmoved in r3.

### Survive unchanged (stated for the table)
- **Leg 2 — m-2 @ `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`** — producer bytes unmoved; my F55/F63 sufficiency confirm stands, now on consumer basis `111ab95a…`.
- **Leg 3 — m-1 @ `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`** — producer bytes unmoved; my lifecycle-fit confirm stands (and r12's `credential_ref` sits squarely in their §1.4a class, per the routed census nod).

## Verification
- Recomputed this session: m-7 r8 `ab0ed428…` · m-3 r3 `70838f83…` (both exact); consumer basis r12 `111ab95a…` (approved `…-120514`, unchanged by this return).
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/step3-mvp-confirm-m10/RECONCILE-planner-20260717-124500.md` — run at filing; result inline.

ACTIONS_GIT_REF: none — a re-affirmation relay + one INDEX.md row timestamped 20260717-124500; no doc edit, no `frank/` edit, no code, no lock.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: master enters both legs in the close table (all four m-10 outbound legs now current on final bytes); m-10.planner stands by for the stage-3 halves and the stage-5 DESIGN + grill dispatch.
