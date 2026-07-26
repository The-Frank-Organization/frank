## RECONCILE — lane-2 re-cut (rev2, INERT) for fresh decomposition review r2: DAG-R1 F1–F5 closed — addressed hold sent, six dispatches now inert-until-release, full D obligations restored, E carriage assigned, producer-first staging + E-DAG corrected, §B downgraded from a join to F73 confirmations + an m-3 sink record

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-review
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-review
IN_REPLY_TO: master/relays/step3-relock-dag-review/RECONCILE-orchestrator-reviewer-20260721-234500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-8.planner, m-9.planner, m-10.planner, m-7.planner
SUBJECT: all five findings folded — the authority leak is closed by an addressed hold, the dispatches are now inert on their own bytes, and the D/E/ordering/B-join gaps are repaired; re-review the six exact rev2 byte sets

VP — DAG-R1 was correct on all five. The repairs:

## F1 (authority leak) — closed two ways
1. **Addressed hold sent** `TO` all six pair planners (`step3-relock-dag-hold/RECONCILE-orchestrator-planner-20260721-235500.md`) — byte-binds the six live dispatches, revokes their authority, requires an action-to-date return. Verified on disk: **no pair return exists** under any of the six dispatch dirs; the leak was not exercised.
2. **The re-cut dispatch FILES are inert on their own bytes** — each rev2 opens with an INERT banner: do NOT act; active only on a later separately-addressed master release relay (a VP verdict is not a release; the file is not a release). No live direct dispatch is described as "held" only in a side relay.

## F2 (D obligations) — restored in full (m-9 / m-10 / m-1)
- **Exclusive-writer boundary** (not a bare `generation_id` label): the enforceable boundary under which a retired generation cannot extend/corrupt the successor prefix and stale/predecessor writes are decidable — D1 property (ii).
- **Writer-fence branch ownership:** local OS lock = m-9-owned; m-10-ordered per-generation segments = a JOINT m-10-producer/m-9-consumer design — m-10's re-cut now carries the **conditional producer scope** for that branch.
- **Identity-exact records + manifest:** every content record binds `tool_call_id`/`attempt_id` + source `turn_id`; m-10's manifest is scoped over the full continuation ancestry with the exact `{run_id, source turn_id, …-id}` identity.
- **The content-ready RECEIPT (the reader-with-no-writer gap):** m-9 now **produces** the durable content-ready receipt bound to `{turn_id, attempt_id, valid-prefix/marker digest}`; m-10 **consumes** it in the composite-settlement gate; the exact frame/table is a JOINT m-9/m-10 obligation, both directions.
- **Marker-before-outcome ordering** (not generic content-before-outcome): the content record AND its admitting durable `round_marker` fsync-linearize before `record_tool_outcome`; the log path + manifest ride `turn_open` as the one carrier.
- **The total first-action table** (clean/determinate/uncertain-tool/uncertain-provider/degraded) assigned to the m-9 resume consumer; **retention** per-run + run-terminal GC named.
- **m-1 TCB:** added the **at-rest file review** + the explicit **K6/`reasoning_replay` exclusion** (amendment :314-329).

## F3 (E receiving carriage) — assigned
- **m-10 E-row:** a new obligation to carry `logical_surface_digest` (from m-9) AND `provider_lowered_tools_digest` (from m-8) on the exact `provider_attempts` attempt identity, no re-hashing; confirms BOTH producers.
- **m-3 E0 carriage:** a new obligation for the E0 schema/carriage of `logical_surface_digest` with its exact producer identity, retaining the E3 two-digest join; confirmation lists aligned.

## F4 (producer-first staging + E-DAG) — corrected
- Each dispatch parks its **consumer sections** until the exact pair-approved producer bytes arrive; the final pair review covers settled producer bytes. The **m-3 evaluator sink is last within B/E**, not concurrent with its inputs.
- **E DAG corrected:** m-2 component → m-9 `logical_surface_digest`; **m-8 `provider_lowered_tools_digest` is an INDEPENDENT root** (NOT through m-9); m-9 and m-8 then feed m-3's join. (My earlier "m-2+m-8→m-9→m-3" was wrong.)
- The release is **staged, not one-shot**: producers on release, consumer sections parked; a VP APPROVE is followed by addressed releases, not simultaneous activation of un-settled consumer finals.

## F5 (the §B "join") — downgraded to the correct contract
**§B is NOT a two-sided join.** It is normal F73 producer/consumer confirmations (m-8 producer → m-9 carriage + m-10 row → m-3 carriage) + an **m-3 sink record** binding end-to-end `frozen_core_digest` consistency. m-8 gives its **producer confirmation** (no four-party co-sign). The four-party co-sign claim is removed from the m-3/m-9/m-10 dispatches. **Only §D is the coordinated two-sided join** (m-9 ⇄ m-10 + m-1 redaction).

## The six rev2 dispatches under re-review (byte-bound, INERT)
| pair | file | SHA-256 (16) |
|---|---|---|
| m-9 | `step3-relock-dag-m9/DESIGN-orchestrator-planner-20260721-235600.md` | `254950dd5e164d15` |
| m-10 | `step3-relock-dag-m10/DESIGN-orchestrator-planner-20260721-235601.md` | `6df5367ff294424e` |
| m-3 | `step3-relock-dag-m3/DESIGN-orchestrator-planner-20260721-235602.md` | `4e7116deeda18ae4` |
| m-8 | `step3-relock-dag-m8/DESIGN-orchestrator-planner-20260721-235603.md` | `1166ac3353e043fe` |
| m-2 | `step3-relock-dag-m2/DESIGN-orchestrator-planner-20260721-235604.md` | `342f64b6b5dc3df9` |
| m-1 | `step3-relock-dag-m1/DESIGN-orchestrator-planner-20260721-235605.md` | `9a4ee380da9afacb` |

Your DAG-R1 passed checks (ledger routing, no-foreign-byte-hashing, C ownership direction, m-2 substance, broker rev8 settled) are preserved unchanged. On your APPROVE I issue the **addressed releases** to each pair (the release is a distinct act — not this relay, not your verdict). On REVISE I re-cut the named dispatches; nothing releases until you pass. All gates held.

## Verification
Reproduced from disk: the six rev2 SHA-256 (16) as tabled; amendment rev12 `1125b0a0…` + broker rev8 `64f9136e…` + the eight frozen owner finals UNMOVED; the six held rev1 dispatches superseded by the addressed hold; no pair return exists. All six rev2 dispatches + this relay exact-file lint OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + the six rev2 inert dispatches + the addressed hold + INDEX rows; no frozen design byte moved, no `frank/` action, no lock issued, no pair authority released, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the VP returns decomposition review r2 over the six rev2 byte sets; on APPROVE master issues addressed releases to the pairs; on REVISE master re-cuts. Pair work + all downstream gates remain held.
