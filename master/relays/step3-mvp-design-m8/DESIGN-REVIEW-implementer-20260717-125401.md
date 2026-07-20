## DESIGN-REVIEW - m-8 adversarial final-byte review of the provider contract r3

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-review-r3
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings are bounded protocol-totality and wire-proof corrections under the ratified topology; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this review does not reopen an operator-locked choice; the stage-2 dispatch routes grills to stages 4/5
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m8/DESIGN-planner-20260717-124500.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-1.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-125401.md
SUBJECT: MUST-REVISE exact r3 6c586f35...fb5617 - the seam confirmations landed, but pre-stream rejects cannot close the already-open attempt row; replay field names drift; DisableKeepAlives adds an uncensused Connection header; the dial-failure fixture requires an impossible request write

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner - I reviewed the exact r3 bytes at SHA-256 `6c586f3576bc036890c75707eb66619ac7ea62e68536c86c0cd2b718f9fb5617`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, master release, m-10 r12 owner approval, and m-9's replay-custody confirmation all pass. The r3 provenance flip is faithful: `credential_ref` is contract-real in m-10 at `111ab95a...`, and the m-9 custody half is confirmed and now normative in its current lifecycle-half bytes.

The original R1-F1..F6 repairs are materially present, but a fresh whole-document pass finds four active contradictions. These exact bytes cannot receive final pair approval.

## Findings

### R3-F1 - Pre-stream typed rejects have no total attempt disposition after `attempt_open`

m-8 says malformed requests, lane-capability mismatches, and replay-scope violations are rejected before freeze with “no attempt opened” / “zero attempt” (`2026-07-17-mvp-provider-contract.md:46-47,125,213`). The consumed lifecycle says the opposite at the durable accounting boundary: m-9 sends `attempt_open` on CTRL-W **before** the DATA-P request, creating one m-10 `provider_attempts` row for the invocation (`m-9 ...mvp-lifecycle-half.md:75-87`; m-10 ...ipc-manifest-seam-contract.md:61,228-230).

The §1.3 outcome table calls itself total but omits every pre-stream reject. The only m-8 CTRL-C dispositions are `{sent_completed, denied(<m-3 token>), transport_failed, unknown}`; `denied` is reserved for m-3 policy tokens, while `transport_failed` is false before transport. m-9's current §2.8 says `replay_scope_violation` becomes a typed attempt failure “in the existing disposition set,” but no such no-stream failure disposition exists. With the current bytes, m-10 sees `attempt_open` plus no recognized terminal view and parks `UNKNOWN_PROVIDER_OUTCOME` for a deterministic local reject.

Required revision:

1. Make §1.3 total over every pre-stream reject (`malformed_request`, replay-scope violation, lane-capability mismatch, and any other pre-freeze typed reject). Pin the exact DATA-P reply, m-8 CTRL-C result, m-9 stream/no-stream behavior, m-10 terminal row state, m-3 phase, and turn-terminal mapping.
2. Route any new CTRL-C disposition/store state to m-10 and the worker mapping to m-9; m-8 must not silently extend their enums.
3. Replace “no attempt opened” / “zero attempt” with the honest boundary: the durable attempt row already exists, while provider transport invocations remain zero.
4. Add fixtures proving each local reject leaves exactly one terminal, non-UNKNOWN attempt row, emits no provider stream, and invokes no resolver or transport.

No operator choice is required; this is protocol totality across already-owned seams.

### R3-F2 - The replay-envelope fold left two incompatible event schemas

The request side and normative event bullet use the new wrapper:

- `reasoning_replay{envelope}` and `replay_envelope{origin_provider_lane_id, origin_turn_id, payload}` (`m-8 :44-46`);
- `reasoning_end.replay_envelope` (`:69`);
- m-9's confirmed/normative custody row consumes that exact envelope (`m-9 :115-120`).

But the closed lifecycle grammar still emits `reasoning_end{block_id, replay_payload?}` (`m-8 :55-64`), the pre-r2 field. A conforming producer and consumer can therefore disagree on the field name and whether the origin tuple exists. Section 5.2 also says only `reasoning_replay → opaque reasoning item, verbatim` (`:182-183`) without normatively stating that scope validation precedes translation and only `envelope.payload` is unwrapped into the provider wire item.

Required revision: choose one closed field shape and use it in the grammar, bullets, request schema, translation profile, fixture corpus, and m-9 confirmation. The expected repair is `reasoning_end{..., replay_envelope?}`, pre-translate origin validation, then payload-only wire re-emission. Unknown legacy `replay_payload` must fail closed rather than bypass provenance. Update the stale “confirmation asked” text at `:46`; the confirmation has landed and is normative in current m-9 bytes.

### R3-F3 - The “exhaustive” wire census omits `Connection: close`, which the pinned transport adds

The no-retry construction pins `DisableKeepAlives=true` (`m-8 :133,179-180`). In the selected Go 1.26.4 `net/http` implementation, that exact setting adds an extra header below the frozen request:

`transport.go:2866-2869`:

```
if pc.t.DisableKeepAlives && !req.wantsClose() ... {
    req.extraHeaders().Set("Connection", "close")
}
```

The §2.1 census lists user-agent, compression, host, content-length, protocol/framing, and canonicalization, then declares any other field a defect (`m-8 :110-119`). It does not list `connection`. Thus the passing request defined by §5.1 necessarily violates fixture 14's “nothing else” assertion (`:212`) and differs from the claimed complete actual-wire identity.

Required revision:

1. Census `Connection: close` explicitly as a deterministic consequence of the no-reuse construction, or pin an alternate construction that emits no such uncensused field.
2. Reconcile the formal `frozen_core.headers` comment (“COMPLETE non-auth header set; nothing may be added later except auth,” `:101`) with the separately derived transport fields. Name whether the field is frozen or transitively derived, then make fixture 14 assert that exact form.
3. Replace the stale F12 disposition at §9 (`:217`), which still describes the old digest-only proof, with the actual three-leg proof: core digest + authorized envelope + exhaustive on-wire census/capture. The fold log says this was restated, but §9's active disposition was not.

The HTTP/1.1-only, no-reuse, `GetBody=nil` no-replay construction otherwise remains a viable bounded design; this finding does not reject that choice.

### R3-F4 - Fixture 4's dial-failure expectation is mechanically impossible

Fixture 4 applies one assertion to every injected cut, including fresh-dial failure: “exactly ONE dial + ONE request-write” (`m-8 :202`). If the dial fails, no connection exists and no request can be written; the truthful vector is one dial attempt and zero request-write starts. The current assertion either cannot pass or forces the counter to label an outer intention as an actual below-encoder write, defeating the proof boundary R1-F3 required.

Required revision: define the counter vocabulary and expected vector per cut, not one blanket result. At minimum distinguish `dial_attempts`, `connections_established`, `request_write_started`, and `request_write_completed` (or an equally unambiguous set):

- fresh-dial failure: one dial attempt, zero established connections, zero request writes;
- post-connect/nothing-written fault: one dial, one established connection, at most one write start, zero completed request;
- headers-received and mid-stream cuts: one dial, one request write, no second dial/write;
- pool-absence and redirect cases: no second dial/write.

Every vector must still close exactly one attempt row through the R3-F1 disposition and prove no automatic replay.

## Accepted portions

- **R1-F1 closes.** The m-3 basis is current at pair-approved r3 `70838f83...`; `turn_epoch` is the same canonical-decimal string across the relevant surfaces.
- **R1-F2's owner seam closes.** m-10 r12 `111ab95a...` owns the run-frozen `provider_lane.credential_ref` plus the seventh verbatim `connector_assign` field; its fresh pair review is approve. m-8 retains grammar/membership/duplicate validation, READY withholding, and post-authorize resolution.
- **R1-F5's custody prerequisite closes.** m-9 confirmed verbatim, exact-turn/exact-lane, in-memory-only custody; its current lifecycle-half §2.8 carries the row normatively. R3-F2 is only the remaining producer-schema consistency defect.
- **R1-F6 closes in ownership shape.** Local fixtures 13a/13b and the full sentinel owner map no longer claim that m-8 locally realizes every cross-lane annex row.
- F72, authoritative terminal usage, lane-capability rejection, claim ceiling, F11, and F13 remain accepted.
- No finding changes topology, policy ownership, secret custody, evidence altitude, or an operator-locked design choice.

## Revision bar and gate disposition

Return fresh bytes that:

1. make all pre-stream rejects terminal and non-UNKNOWN across m-8/m-9/m-10, with the necessary owner confirmations;
2. normalize the replay-envelope field and payload-only translation path;
3. include every transport-added field, especially `Connection: close`, in the F12 wire identity and active §9 disposition; and
4. give each no-retry fault cut a physically possible below-encoder counter vector.

Also remove the now-stale “confirmation asked” / “pending revision” status language where the current owner bytes have landed. The new SHA requires a fresh uniquely-parented DESIGN-REVIEW; any m-9/m-10 owner-byte change must land and be reviewed before that request.

This verdict is byte-bound to `6c586f3576bc036890c75707eb66619ac7ea62e68536c86c0cd2b718f9fb5617`. The stage-2 approval SITREP, Master+VP interface lock, PLAN, T4 token, code, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256: `84294a26587d3d8051bb3170095b0e2801f9bd52f71375bbb2d2b173e24c3c07`.
- Exact reviewed m-8 r3 SHA-256: `6c586f3576bc036890c75707eb66619ac7ea62e68536c86c0cd2b718f9fb5617`.
- Current/pair-approved m-10 r12 SHA-256: `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`.
- Current m-9 lifecycle-half SHA-256: `b4e08545a90267d3e1c646d0e1b13c9afc86218e83959ca42ad5fc56b494d29e`; §2.8 contains the normative custody row. Its own fresh pair review is a separate lane and is not claimed complete here.
- Current/pair-approved m-3 r3 SHA-256: `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.
- Current/pair-approved m-1 SHA-256: `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- Incoming DESIGN relay exact-file lint: OK.
- Go transport source inspected at local `go1.26.4`; automatic `Connection: close` addition reproduced at `net/http/transport.go:2866-2869`.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-8.planner routes the R3-F1 disposition gap to m-10/m-9, folds R3-F2 through R3-F4 after the needed confirmations, recomputes the design SHA-256, and issues a fresh uniquely-parented DESIGN request; do not file the stage-2 approval SITREP on r3.
