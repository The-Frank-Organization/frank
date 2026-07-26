## DESIGN-REVIEW — APPROVE m-10 B/E carriage rev3: carrier confirmed at m-9 r12, R3 row existence is explicit, and the prior absence/handoff defects close

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-be-review-r3
PARENT_DISPATCH_ID: step3-relock-dag-m10-be-carriage
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — exact-byte approval of the bounded m-10 carriage row; m-3 must still reproduce and bind it
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10-be-carriage
DESIGN_DOC_SHA256: cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-dag-m10-be-carriage/DESIGN-planner-20260724-160000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-8.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: APPROVE exact B/E carriage rev3 cd17db32 — m-9 r12 pair-confirmation closes the carrier dependency, the five-member freeze makes attempt_open timing sound, R3 present|not_found is authoritative and digest-conditional, and the m-8-only NULL/unknown scope remains exact

DESIGN_REVIEW_VERDICT: approve

m-10.planner — **APPROVE** the complete B/E carriage rev3 at exact SHA-256 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`. I reviewed the directly addressed request at `7c9a91baf5ab638a3247701d7298be9f366f53a41382810890aab495eb6a0149`, the full rev3 bytes, the prior rev2 must-revise verdict, m-9 r12 §6 and its byte-bound approving review, master's R3 routing, and m-3's current row-state machine.

## Prior findings close

- **R2-F1 closes.** The exact producer dependency now exists: m-9 r12 `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35` is pair-approved. Its §6 confirms `attempt_open.logical_surface_digest`, REQUIRED/non-null, before the m-10 row commit, and its implementer explicitly lands the second confirmation half. Rev3 binds that hash rather than any superseded r5-r11 revision.
- **The carrier timing is sound.** R12 freezes all five digest members at first assembly; Gate-2 reassembly changes assembled input content, not a member, and a true member change creates a new attempt. The A3 sentinel-bearing `user_text{text}` carrier therefore does not stale the already-sent digest. Rev3's scoped m-10 confirmation remains only the two manifest-carried arrays; it does not re-own the other three recipes.
- **R2-F2 closes.** §6 now conditions the handoff on this pair approval and m-3 reproduction. NULL/absence-to-`unknown` is restricted to m-8's two P1-derived columns. `logical_surface_digest` is required on every committed attempt row; assembly refusal or malformed/missing carriage creates no attempt row and never a NULL logical digest.

## R3 row-existence review

The new §5a contract is coherent with master's routed requirement and m-3's sink:

- `m10_row_state` is the closed `present|not_found` result of an authoritative query against m-10's committed single-writer store, distinct from timeout, malformed, unavailable, or ambiguous acquisition.
- `present` means a durable `provider_attempts` row exists and makes its digest columns queryable; `not_found` means no row exists and the digest members are absent rather than NULL.
- A committed row never regresses to `not_found`; retirement parks rather than deletes it. The relevant N3 stale-rejected identity remains the legal no-row shape, while m-3 independently rejects `not_found` for P/A/N910.
- FX-BE-7 covers committed, stale-rejected, retirement, and never-NULL logical-digest legs. FX-BE-8 covers byte equality between the `attempt_open` digest and the surface ultimately sent across Gate-2 reassembly.

## Preserved boundaries

Verbatim carriage/no rehash/no join, attempt-exact identity, m-8 P1 NULL handling, write-once/first-wins conflict, structural validation, evidence-never-authority, fixed bounds, no-payload storage, and the F60/F67 secret boundary remain intact. The VP's delegated A3/B1 ruling adds no m-10 manifest or request schema member.

This approval is exact-byte and narrow. It makes rev3 available for m-3's independent reproduction and binding; it does not itself complete m-3 R3, the §D join, settlement amendment, integrated re-lock, DESIGN lock, PLAN, T4/code, credentials, provider calls, release binding, E3, merge, or deploy.

## Verification

- Incoming DESIGN relay: `7c9a91baf5ab638a3247701d7298be9f366f53a41382810890aab495eb6a0149`.
- Approved B/E rev3: `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`.
- Pair-approved m-9 r12: `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`; approving review `DESIGN-REVIEW-implementer-20260724-023000.md` read in full.
- Exact prior acceptance bars, current §5/§5a/§6/§7 bytes, m-9 r12 §6, master's R3, and m-3 RS0-RS3 were re-read directly.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10-be-carriage` exited 0
Next requested action: m-10.planner returns pair-approved rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` to Master for m-3's independent reproduction and R3 binding; all later gates remain separately held.
