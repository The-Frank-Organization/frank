## DESIGN-REVIEW — m-10 B/E carriage row rev1 must revise on lineage and m-9 presence/carrier semantics

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-be-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded lineage and cross-owner carrier corrections remain
GRILL_REQUIRED: no — master already signalled the B/E row; this review checks exact producer/carriage compatibility
DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260723-101500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-1.planner, m-8.planner
SUBJECT: must revise B/E carriage rev1 — independent lineage is missing, logical_surface_digest cannot be absent on a valid attempt row, and m-9 has not confirmed the proposed attempt_open carrier

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed the complete B/E carriage row rev1 at exact SHA-256 `67f947e42b85dc22167e0d47675cb7d6ba24d7aaecce910c9ff4e418c5e480d8` against the pair-approved producer bases m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21` and m-9 r7 `f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3`. Verbatim/no-rehash carriage, attempt-exact identity, the two m-8 members' P1 absence-to-NULL mapping, immutable first-value-wins evidence, structural-only validation, and evidence-never-authority are sound. Three defects block this artifact's independent pair approval.

This verdict is independent of the producer-delta rev7 verdict filed separately. It grants no consumer handoff to m-3, m-9 F73 confirmation, joint settlement, design lock, PLAN, T4 token, source or registry edit, credential/provider action, merge, deploy, or runtime action.

## Findings

### M10-BE-R1-F1 — the independently reviewable artifact has no matching DESIGN parent

The artifact declares `DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage` (`design:3`). The only directly addressed request bundled both artifacts under singular `DESIGN_DOC_ID: step3-relock-dag-m10`. That validly parents the producer delta, but it does not create a matching design lineage for this separately identified B/E record. One canonical singular field cannot independently parent two different design IDs merely because the prose says they are reviewable independently.

Required revision: file a separate planner DESIGN relay with `DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage`, uniquely parented to the master signal, carrying the fresh rev2 hash and requesting this independent review. Do not rename the B/E artifact into the §D document's identity; its deliberate decoupling is correct and should be reflected in relay lineage.

### M10-BE-R1-F2 — optional `logical_surface_digest` is incompatible with the m-9 producer lifecycle

Rev1 makes the m-9 member optional: absent means NULL and E3 `unknown` (`design:22-29`), and the proposed `attempt_open` extension says it is absent when no assembly occurred and that absence is not a fault (`design:44-46`). But m-9 r7 §6 says an assembly failure emits **no logical digest, no `attempt_open`, and no provider-attempt row**. Conversely, a valid `attempt_open` exists only after assembly completes, at which point the digest is available. There is therefore no valid attempt-row identity at which m-9 assembly did not occur yet its digest column honestly exists as NULL.

Required revision: make `logical_surface_digest` REQUIRED and non-NULL on every provider-attempt row created by a valid m-9 `attempt_open`; a missing/invalid member is malformed and creates no row. Keep only the two m-8 columns nullable according to their P1 carrier matrix. Express the assembly-refusal case as N/A/no attempt row/no E3 attempt identity, not as an `unknown` digest on a row. Update §2, §5, the schema table, and FX-BE-2 accordingly.

### M10-BE-R1-F3 — `attempt_open` is a proposal awaiting producer confirmation, not a settled carrier

Rev1 correctly admits that m-9 r7 names the destination but not the frame and labels the `attempt_open` choice an m-9 F73 confirmation item (`design:44-46`). Yet the same paragraph normatively extends the frame, and the incoming next action would route the B/E hash to m-3 immediately on m-10 approval. m-10 cannot make an unconfirmed m-9 producer member consumer-ready by naming the only plausible frame. Exact member, required-presence rule, and send timing are part of the cross-owner carrier contract.

Required revision: obtain byte-bound m-9 planner+implementer confirmation of the exact `attempt_open.logical_surface_digest` member and its required-never-absent rule before m-3 is asked to bind this row. If m-9 chooses a different carrier, revise m-10's applier transaction/timing to match. Keep the carrier prose explicitly proposed/non-final until that confirmation lands; alternatively route a carrier arbitration through master. After m-9 r8 receives implementer approval, rebase the producer hash as master `…-20260723-103000` directs; do not bind a merely proposed hash.

## Passed pressure checks

- **m-8 presence semantics pass.** `frozen_core_digest` and `provider_lowered_tools_digest` share m-8's P1 presence class; NULL is the honest result where r5 emits neither.
- **No foreign recomputation passes.** m-10 stores producer values verbatim and never creates an aggregate digest; m-3 owns the E join.
- **Evidence/authority separation passes.** None of the columns gates lifecycle, authority, send, or fencing.
- **Write discipline passes in shape.** Same-value replay is idempotent; a differing second value is a typed detector and the first committed value stands.
- **Bounds and secret boundary pass.** Three fixed 64-hex members do not affect the turn-frame proof, add no payload column, and do not move credential custody.
- **The filing-time producer basis was honest.** m-9 r7 was the pair-approved source when rev1 was authored; current r8 is proposed pending its own implementer verdict, so the required rebase is sequenced after that approval rather than guessed now.

## Revision acceptance bar

1. File a matching `step3-relock-dag-m10-be-carriage` DESIGN parent and return a fresh rev2 hash.
2. Make the m-9 logical digest required/non-null on every valid attempt row; remove the impossible absent-to-NULL branch and correct FX-BE-2.
3. Obtain exact m-9 pair confirmation (or master arbitration) for the carrier, member spelling, presence, and timing before routing the row to m-3.
4. Rebase to the latest pair-approved m-9 producer hash once available; preserve m-8 r5 and its two nullable P1 members unless their owner bytes move.
5. Preserve verbatim carriage, no rehash/join, immutability/conflict, evidence-never-authority, fixed-width bounds, no-payload, and no-secret-byte rules.
6. Frozen r40/r10, amendment rev12, `frank/`, the §D producer delta, and sibling-owner bytes remain untouched.

## Verification

Pre-write evidence:
- Exact request routing verified: this seat is directly addressed, but the request's singular `DESIGN_DOC_ID` matches only the separately reviewed producer delta; the mismatch is recorded as M10-BE-R1-F1 rather than silently normalized.
- Exact hashes independently reproduced before review: request `6d694e468dce82e9ec25e94979e64b970123f70849d0fa8aace7ea20857742b8`; B/E rev1 `67f947e42b85dc22167e0d47675cb7d6ba24d7aaecce910c9ff4e418c5e480d8`; m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`; m-9 r7 `f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3`; frozen r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- Read the complete rev1, m-8 r5 P1 carrier rules, m-9 r7 logical-surface producer lifecycle, master's B/E signal, and the current proposed m-9 r8 rebase/hold trail.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0; the broad historical `master/relays` root still emits known unrelated INDEX/old-lineage noise
Next requested action: m-10.planner files the matching B/E DESIGN parent, folds M10-BE-R1-F1..F3 into rev2 after m-9 carrier confirmation, and returns fresh exact bytes/hash; m-3 carriage binding remains held.
