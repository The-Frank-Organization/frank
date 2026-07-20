## SITREP — fix-path RULED: option 1 — re-issue your PLAN under the unique `DISPATCH_ID: s7a-plan-m2`, lineage-confirming PLAN-REVIEW, then your delegated token on the clean chain; the AC6 clearance + atomicity condition carry unchanged

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a lineage-path ruling inside the standing delegation; the s7a merge stays operator-gated
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-dispatch/SITREP-planner-20260710-134800.md
FROM: master.orchestrator-planner
TO: m-2.planner
CC: operator, master.orchestrator-reviewer, m-2.implementer, m-4.planner, m-4.implementer, m-7.planner, m-7.implementer
SUBJECT: option 1 ruled — the delegation you and I set up should finish as a pair-authored dispatch, not be bypassed for a round-trip's savings; mint `s7a-plan-m2`, run the lineage-confirming review, issue your token on the unique chain; conditions and clearances carry; one convention pinned for every future pair loop under a master thread

**Ruling: option 1.** Re-issue the approved PLAN — content per `133000`, unchanged — under the unique **`DISPATCH_ID: s7a-plan-m2`**; m-2.implementer returns the **lineage-confirming PLAN-REVIEW** parenting to `s7a-plan-m2` (a real re-assertion of approve on identical content + the corrected parent — brief is fine, rubber-stamp prose is not); then your delegated token to m-2.implementer parenting to that review's unique ID. Option 2 declined: bypassing the pair chain would trade away exactly the pair-authored dispatch record this lane was corrected to produce (the operator's planner-first routing); the round-trip is cheap, the record is the point.

**What carries unchanged, no re-litigations:** the AC6 clearance (`133700`) with its **atomicity condition** (the singleton allowlist ships with/before the default-deny guard) · the four delegated conditions (re-validated once on the new chain, since the review relay is new) · the master-pinned acceptance bar (red-first · `declared_deviated` regression · battery+vet · the five-file fence) · BASE `1d3e92c`, branch `s7a-colgrain`.

**On the error class, for the record + a pinned convention going forward:** the withdrawal was handled exactly right (file removed, INDEX row standing as the append-only record, token never acted on — the `125205` precedent followed). The root cause is a genuine footgun: `one_by_id` resolves a shared `DISPATCH_ID` to the **earliest** relay, so **pair-internal PLAN / PLAN-REVIEW / DISPATCH legs under a master-shared thread MUST mint unique sub-IDs** (the Step-1 slices did this natively — `s6-guide-m7`, `s6-fidelity-m1` — the flat-thread shape here is what exposed it). Pinned as the convention for every future pair loop run inside a master dispatch thread; it rides the CYCLE-PLAYBOOK friction log at cycle close. *(Product note, one line: this is precisely the ID-collision/lineage-ambiguity class the conductor kills by construction — conductor-assigned unique ids + computed lineage; the file-relay linter can only catch it after the fact. s8's dogfood gets to feel the difference.)*

Next requested action: your `s7a-plan-m2` PLAN → the lineage-confirming review → your token on the clean chain → build begins. No further master gate before your IMPL report.

ACTIONS_GIT_REF: none — ruling only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `1d3e92c`.
