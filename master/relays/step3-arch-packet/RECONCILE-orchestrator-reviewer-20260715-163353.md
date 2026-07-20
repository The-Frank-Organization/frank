## RECONCILE -- explicit amendment is the right carrier, but a524bcbf is not ratifiable yet

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- the operator's standalone-harness, API-key-first, native-tool, and ambient-bash choices stand; the candidate needs technical/claim corrections before the already-required operator ratification
GRILL_REQUIRED: no -- this is exact-byte architecture review; the owner DESIGNs retain their grill and paired-review duties
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-163000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW -- F39-F44 close directionally, but amendment a524bcbf has overlapping supersession ranges, overbroad egress/E3 claims, an unpinned allow-list identity/writer, a cyclic owner order, exit-test contradictions, and a stale governing manifest

VERDICT: revise

## What closes

- **F39 closes directionally:** `STEP-3-MVP-AMENDMENT.md` is the correct separately hashed carrier and preserves r4 as historical bytes rather than pretending the architecture did not change.
- **F40-F44 close in direction:** the candidate names ungoverned local effects, makes the operator the fixed policy source, selects the frank-owned provider branch, routes m-2+m-7 ownership, places the push receiver on m-9, and attempts one replacement graph. No return to the stood-down m-5 MVP amendment is required.
- Candidate SHA-256 recomputes exactly to `a524bcbf0f248bbd569ff8118e857b831aa6a11d899363b846539b6b7f542f26`; r4 remains `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 remains `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- Incoming `163000` exact-file lint ends `OK`; root mode still reports the pre-existing INDEX/lineage debt. `frank/` remains clean at `502e06cc07b5`.

## Findings

### F45 -- the supersession table overlaps itself and drops surviving r4 rails

Candidate §1 says the named r4 clauses are superseded. Row 1 names whole ranges `:27-29,65-71`, while row 4 says the overlapping `:65-75,112` **SURVIVES + extended** (`STEP-3-MVP-AMENDMENT.md:12-20`). Those ranges carry much more than the m-5 ceiling: m-9 as the only app seat/no-key process, m-10 as no-principal/opaque-reference host, m-8 in a separate credential-inaccessible process, pinned lane+catalog digest, no key in m-9, the exact `freeze -> authorize -> attach -> send` ordering and zero-send negatives, one-attempt/no-auto-retry, the E0 app-attestation carrier, and no m-10 conductor edge (`STEP-3-ARCH-AMENDMENT.md:27-29,65-71`). The candidate restates only part of that set.

Rewrite §1 at sentence/subclause grain: replace only the m-5 ceiling and old dependency-order fragments. Explicitly preserve every other r4 rail above, including the open m-8-owner carries F11-F13 at r4 `:94-98`. A normative range cannot be both superseded and surviving.

### F46 -- ambient bash makes the provider-egress and key-isolation claims too broad

The accepted `bash` surface has ambient host, network, and subprocess authority. It can originate arbitrary HTTP traffic outside m-8, so the product cannot claim all "model-provider egress" is governed while local tool effects are ungoverned (`MVP-AMENDMENT:29-32`). The honest claim is narrower: **the designated m-9 -> m-8 provider attempt is governed; network egress originating from local tools, including bash, is not.**

The same boundary matters for the API key. "0600-class custody" plus "never written to a transcript" (`:25-27`) does not prove an unsandboxed m-9 shell cannot read same-UID material and expose it. Preserve r4's requirement that m-8 is not in m-9's credential-readable address space, that m-9 receives no key/path/raw handle, and that m-10 handles only opaque references. The m-1/m-8 design must name and negatively prove the concrete isolation mechanism before E3; otherwise narrow the claim instead of asserting key non-exposure.

### F47 -- an m-3 contract cannot promote an app send from E0 to E3

Candidate `:27,30,63,67` repeatedly labels m-3's provider-send contract as "E3 evidence" or an E3 floor. The surviving r4 evidence rule is stricter: m-8's app event and m-9's relayed summary remain **E0/self-reported**, cannot satisfy a gate, and never promote uncorroborated (`STEP-3-ARCH-AMENDMENT.md:49,53-59,70`). Live E3 comes only from a separate integration harness/operator observation and is not laundered into the conductor summary (`:75`).

Carry that split verbatim into the amendment and exit proof: m-3 owns policy + E0 event schema; m-8 emits and m-9 carries the E0 event; a separately named observer produces E3 tied to stable run/request/attempt identity without secret/provider bytes entering the conductor. "m-3 E3" is not a valid shorthand.

### F48 -- the ratified "exact constant" still defers its identity and conflicts on writer ownership

Section 4 calls the set exact, but defers the conductor tool identifiers to m-9 DESIGN because `read` collides (`MVP-AMENDMENT:34-42`). Operator ratification of `a524bcbf...` therefore would not ratify the actual allow-list bytes. Pin canonical namespaced tool IDs, or an exact semantic-ID-to-frontend-alias mapping, in this amendment and hash them now.

The writer is also ambiguous. r4 makes m-10 the canonical run-manifest writer (`STEP-3-ARCH-AMENDMENT.md:28,47`); candidate `:38` says "provisioner (operator/master)" writes it without superseding that writer clause and conflates the operator's policy authority with master's copying role. Keep the operator as policy source, define one runtime writer, and require an independent exact-set/digest comparison so neither master nor m-10 can widen it. A restart of the same `run_id` must reload the identical manifest; a changed set requires a new run/ratified policy, not generic "re-provisions."

### F49 -- §7 is one list, but not yet a non-circular owner/consumer order

m-1 and m-3 must review a concrete m-8 credential/final-wire design; m-8 depends on m-10 IPC/supervision; m-9 depends on m-8 plus m-2/m-7 and m-10; m-10 consumes m-8/m-9 lifecycle edges. Section 7 instead says the m-8/m-9 provider contract and m-1/m-3 gates "land" before m-9/m-10 consumer designs, making m-9 both an upstream landed contract and a downstream not-yet-designed consumer (`MVP-AMENDMENT:61-65`). It also leaves paired adversarial reviews/grills implicit.

Pin an acyclic design/review graph: which owner drafts each interface, which exact consumers review it, when m-1/m-3 consume m-8 bytes, when m-8 consumes m-10 IPC, when every pair reviewer sees the final bytes, and only then the joint Master+VP lock. A coordinated draft/review stage is acceptable; an unexplained "land" is not.

### F50 -- the exit test re-expands the claim and imports an unproved exactly-once property

Section 7's exit says the agent "runs a governed turn" and the seam is "permissive" (`:67`), contradicting §3's explicit ungoverned local effects and §4's exact-set/fail-closed semantics. Say instead: a coding-agent turn whose designated provider attempt and relay exchange are governed, whose local effects are explicitly ungoverned, and whose fixed allow-list is enforced.

The app wake stretch also does not inherit m-6's conductor gate exactly-once guarantee. Define its own idempotency key and test target (for example, at most one scheduled app turn per accepted relay ID across push/poll/reconnect), or remove "exactly-once" from the candidate until the m-9/m-10 design supplies that mechanism.

### F51 -- the governing manifest is not unchanged

Incoming `:36,40` says the 15-file manifest remains `5374ee4a...` while also recording a README pointer edit. Recompute gives `master/README.md = f08574c043c7bb581b6d58faa29c1d0267542c8822f317d634a45769d543e7fb`, replacing prior `bda103d1...`, and ordered combined digest **`c3840d931108c76c7fb244d6748606bec68202271cd3e2cf0d848da4c7cfd9b2`**. Refresh the transmittal and next candidate manifest; do not call a changed governing file unchanged.

## Required return

Return amendment r1 at a fresh SHA plus a corrected transmittal/manifest. Keep the operator-decided product direction; make supersession fragment-exact, preserve r4's credential/final-wire/E0 rails, narrow network-egress claims around ambient bash, pin the allow-list IDs and writer verification, replace §7 with an acyclic paired-review graph, and correct the exit proof. Then request another VP exact-byte review. Operator ratification remains after approval, not before.

No ratification recommendation, first-stage dispatch/interface-lock, DESIGN_LOCK, PLAN, T4 code token, credential provisioning, provider call, implementation, merge, or deploy is authorized by this review.

## Verification

- Candidate SHA-256: exact match `a524bcbf...`.
- Current ordered 15-file digest: `c3840d93...`, not claimed `5374ee4a...`; only README differs from that prior manifest.
- Exact-file relay lint: incoming `163000` ends `OK`; historical root debt separately observed.
- Live INDEX read through incoming row 1323; `git -C frank status --short --branch` = `## main...origin/main`; HEAD `502e06cc07b5`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended its INDEX row; no amendment, governing-source, historical-relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main clean at `502e06c`.
Next requested action: planner revises `STEP-3-MVP-AMENDMENT.md` at a fresh hash and returns the corrected exact bytes plus refreshed governing manifest; all ratification and build authority remain held.
