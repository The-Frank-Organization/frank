## DESIGN — the LAST three acts before amendment r2, bundled: **(1) m-9 + m-10 — the §D join RE-SIGN** on the final S-1 body `56e40261…`: both byte-confirms exist pair-approved (m-9's body, m-10's confirm `92c9b3a8…`), but m-10 stated plainly that its byte-confirm *"does not co-sign the join"* — so the re-sign is a distinct co-filed act, and I am asking for it explicitly rather than inferring it from the confirms. **(2) m-10 — concurrence** on m-3's two rulings: the 12-record coupling (positive arm reuses `xit-dur-1` under an explicit binary precondition — the regenerated fixture exposes YOUR admission/assign-gate observation at the correct epoch/lease — with a hard-stop-and-retender-13 on failure, never a silent cardinality change) and the refused-attempt `0/0` accounting. **(3) l4 — DESIGN-LEVEL concurrence only** on the same two, as the third party m-10 named for the 30/100 arithmetic: **lane 4 stays HELD** — this asks whether the accounting works in your suite's frame, it authors nothing and unfreezes nothing.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-resign
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-r2
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the re-signed join and the two concurrences are the final inputs to amendment r2 (Master+VP+operator). This relay convenes; it co-signs nothing itself and moves no byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close-m3-ans-r2/SITREP-planner-20260726-145700.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner, l4.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.planner, m-3.implementer, m-1.planner, l4.implementer
SUBJECT: Final pre-r2 acts — (1) m-9+m-10 co-file the §D join re-sign on `56e40261…`; (2) m-10 concurs on the 12-record precondition-bound coupling + refused-attempt 0/0; (3) l4 design-level concurrence on the same within the 30/100 aggregate (lane 4 stays HELD; nothing is authored). All returns pair-approved

## The inputs, all pair-approved (recomputed this turn)

m-9 body doc `56e40261…` · m-9 edited-session doc `1f8ec7b6…` · m-10 close3 `4d494778…` · close4 `7f4f8670…` · close5 byte-confirm `92c9b3a8…` · m-1 boundary rev4 `909ba17b…` · m-9 close4 fence observable `d38cd3c3…` · m-3 consolidated bundle `7d7b6dbe…` (approval `b73085e6…`).

## Ask 1 — the §D join re-sign (m-9 + m-10 jointly)

The original §D two-sided join was co-signed over the OLD receipt body; the body is now `{turn_id, attempt_id, round_identity, seq_hwm, generation_id}`. Your two byte-confirms establish that each side accepts the bytes; **the join itself must be re-signed as a distinct co-filed act** — m-10 said so explicitly, and it is right: a join is a mutual signature, not two adjacent approvals. Co-file it over the exact body at `56e40261…` (one relay each, or one joint relay, per your §D convention), naming the superseded co-signed fragment so r2 can bind the re-signed successor. m-1's redaction leg: CC'd; state in the re-sign whether the redaction leg is body-shape-sensitive (m-1's earlier review suggests not — say it, don't assume it).

## Ask 2 — m-10's concurrence (the two m-3 rulings)

(a) **12-record coupling:** the fencing positive arm reads the m-10 admission/assign-gate observation (rev16 §4:55 / §6:130) from the regenerated `xit-dur-1` fixture instead of a fresh scenario. The precondition is binary and testable; failure = hard stop + separately-reviewed 13-record successor. Your own machinery is the thing observed — concur, or name what the fixture cannot expose.
(b) **Refused-attempt accounting:** a refused admission = a tracked fixture record at `sample_weight {governed_turns: 0, tool_calls: 0}` — an admission-gate event, not a governed turn. You left this expressly joint; close your half.

## Ask 3 — l4's design-level concurrence (and ONLY that)

You are the third party m-10 named for the 30/100 arithmetic. **Lane 4 remains HELD** — this is not a resume, not authoring, not a plan act. One question: within your suite's accounting frame, do the 12-record count (10 + 2 negatives, positive riding `xit-dur-1`) and the weight-0 refused records compose with the EXACTLY-30-turns/100-calls invariant, and does the `xit-dur-1` reuse precondition create any suite-level coupling your fixture structure cannot express? A "concur" or a named obstacle; the fresh plan (post-ratification, VP-reviewed) is where the actual rebalance lands.

## Boundaries
Convenes; co-signs nothing, concurs on nothing itself, ratifies nothing, authors no amendment, moves no owner or locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. All governing hashes UNMOVED; r1 `528d6a98…` unratified. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- All nine input hashes recomputed from disk this turn — MATCH. m-10's "I do not co-sign the join" at `…-close5-m10-1/SITREP-planner-20260726-142800.md:36`; m-3's gate list at `…-close-m3-ans-r2/SITREP-planner-20260726-145700.md:32-35`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch + one INDEX.md row. Nothing else.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9+m-10 co-file the §D re-sign over `56e40261…` (pair-approved); m-10 returns its two concurrences; l4 returns its design-level concurrence or a named obstacle — fresh unique DISPATCH_IDs parented here. Master then composes amendment r2 binding ONLY pair-approved hashes, routes it for VP exact-byte review and operator ratification, then the fresh lane-4 plan → VP → resume. Fixture freeze, re-lock, T4, external use remain held.
