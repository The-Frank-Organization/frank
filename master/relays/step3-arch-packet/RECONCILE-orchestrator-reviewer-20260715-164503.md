## RECONCILE -- r1 closes F45-F51, but authenticated-seat custody, wake recovery, and owner review remain underspecified

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- the operator-decided product scope remains accepted; these are technical identity, recovery, and owner-review corrections before the pending operator ratification
GRILL_REQUIRED: no -- exact-byte architecture review only; the corrected owner graph must preserve each lane's later grill and adversarial pair review
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-164500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW -- r1 closes F45-F51 and hashes verify; correct m-9 seat-credential custody, make absent-worker recovery credential-free for m-10, restore single-owner paired contract lanes, and pin the allow-list digest mechanism before ratification

VERDICT: revise

## What closes

- **F45-F47 close.** Supersession is now fragment-exact; the credential/final-wire/E0 rails and open m-8 F11-F13 carries survive. Only the designated m-9-to-m-8 attempt is governed; bash-originated network traffic is explicitly ungoverned. m-8/m-9 evidence remains E0 and external observation supplies E3.
- **F48 closes on canonical IDs and writer authority.** The eight semantic IDs are pinned without the `read` collision; the operator is policy source and m-10 is canonical run-manifest writer. One digest-mechanism detail remains as F55.
- **F50 closes on product/exit claim scope.** The exit now distinguishes governed provider+relay surfaces from ungoverned local effects and names an app-specific wake-dedupe target rather than inheriting m-6's scheduler guarantee. Recovery wording still needs F53.
- **F51 closes.** The ordered 15-file governing manifest recomputes exactly to `a8321ef8857e254e01a938870dc0d1932d40c5eeb0888d3f5b2fc1b425f9f624`; README is `1b65ab42ec843d11cd03b689ccef18a334ef0755b59f3ad83dc268bedaaf3828`.
- Amendment r1 recomputes exactly to `02e9da1ce492fbe2f2b1172d636a38e682b8faf13fe3b2c0f4d4d5ba4cd781c9`; incoming `164500` exact-file lint ends `OK`; `frank/` remains clean at `502e06cc07b5`.

## Findings

### F52 -- m-9 cannot be the authenticated seat and hold no submit credential

Section 1 says m-9 is the only app seat **and "holds no submit credential"** (`STEP-3-MVP-AMENDMENT.md:13-15`). That is not what r4 says. r4 gives m-9 its own private seat channel (`STEP-3-ARCH-AMENDMENT.md:27`) and excludes submit credentials only from m-10/m-8 (`:85`). The existing client must call `DialAuthenticated(..., credential)` before it can call conductor tools or receive seat pushes (`frank/internal/channel/server.go:449-458`).

Correct the boundary: m-9's **seat-channel component** holds only m-9's own conductor credential; m-9 holds **no provider credential/key**. The seat credential/private channel must be absent from the model context, local-tool arguments, bash environment/files, m-8, and m-10. Name that channel-side custody as an m-1/m-7-preserved guardrail and let the m-9 design pin the process/FD/sidecar realization. Without it, the native relay tool and push receiver cannot authenticate.

### F53 -- "m-10 polls" has no legal conductor target, and at-most-once is not exactly-once

Section 6 correctly says m-10 has no submit credential, then says that for an absent worker "m-10 polls" (`MVP-AMENDMENT:52-55`). The conductor exposes no non-principal project/read path. `PushTo` reaches only an authenticated m-9 connection; the existing recovery path emits a nudge **when m-9 authenticates again** (`cmd/frank/main.go:358-364`). m-10 therefore cannot poll relay state without violating the no-principal invariant.

Pin the legal recovery path: m-10 polls/supervises **process health only**, restarts the resident m-9 seat receiver, and m-9's authenticated reconnect receives the pending-delivery nudge, performs project/read, then forwards the relay ID over app IPC. Or keep the seat receiver continuously resident while only the model-turn worker sleeps. State explicitly that m-10 never polls the conductor.

Also call the current target what it is: **dedupe/at-most-one scheduled turn per accepted relay ID**. Exactly-once would additionally require an at-least-once/liveness proof across the crash window. Either add that durable scheduling/recovery contract or remove the sentence claiming "exactly-once" against an at-most-one target.

### F54 -- the graph still replaces owner contracts with reviews and leaves joint bytes ownerless

The graph is staged better, but it does not yet preserve the standing single-author/pair-review model. m-3 **owns** the new provider egress policy + E0 schema and m-1 owns the extended secret boundary; they cannot merely "review" m-8's draft (`MVP-AMENDMENT:58-65`). m-2 and m-7 author new shared-client contracts, but §7 guarantees final-byte implementer review/grill only for the three labeled build lanes. m-8's own final-byte pair review is left to the general sentence rather than placed after m-1/m-3/m-9 consumption. Finally, m-9 and m-10 "jointly pin" lifecycle bytes with no canonical writer, despite m-10 owning lifecycle/supervision/IPC and cross-domain integration belonging to Master+VP.

Give every changed contract one owner and pair chain: m-2 and m-7 each author their owned interface and receive implementer review; m-8 authors provider/credential/wire bytes, consumes m-1/m-3 owner contracts plus m-9 review, then its implementer reviews the final fold; m-1 authors/reviews the connector secret-boundary delta; m-3 authors/reviews the egress+E0 delta. Split the mutual lifecycle edge into owned halves (m-10 lifecycle/supervision/IPC, m-9 receiver/turn-state), require reciprocal consumer confirmation, then let Master+VP integrate and lock the join. "Jointly pin" is coordination, not provenance.

### F55 -- the runtime digest check has no canonical digest input

Section 4 says the set is "PINNED + hashed" and that a ratified-constant digest is compared with the manifest (`MVP-AMENDMENT:33-41`), but it defines no canonical serialization, digest value, or policy-reference field. The document SHA binds the whole amendment, not an executable encoding of the eight-element set. PLAN would still have to invent what m-10 compares.

Choose one exact mechanism in the amendment: either pin canonical set equality as the gate and remove the separate digest claim, or define the ordered/canonical encoding, SHA-256 value, manifest field, and expected-source home. Frontend aliases must normalize to the canonical IDs before this check and cannot add members.

## Required return

Return r2 at a fresh hash with only F52-F55 folded, plus the refreshed README pointer and 15-file manifest. Do not reopen F39-F51 or the operator's scope choices. After a clean exact-byte review, the operator may ratify; the source reconciliation and first-stage dispatches remain after that ratification.

No ratification recommendation, source reconciliation, first-stage dispatch/interface-lock, DESIGN_LOCK, PLAN, T4 code token, credential provisioning, provider call, implementation, merge, or deploy is authorized by this review.

## Verification

- Amendment r1 SHA-256: exact match `02e9da1c...`; r4 and canonical m-5 remain byte-exact.
- Ordered 15-file governing manifest: exact match `a8321ef8...`.
- Exact-file relay lint: incoming `164500` ends `OK`; historical root debt separately observed.
- Live INDEX read through incoming row 1325; `git -C frank status --short --branch` = `## main...origin/main`; HEAD `502e06cc07b5`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended its INDEX row; no amendment, governing-source, historical-relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main clean at `502e06c`.
Next requested action: planner folds F52-F55 into amendment r2 at a fresh hash and returns exact bytes plus the refreshed governing manifest; all ratification and build authority remain held.
