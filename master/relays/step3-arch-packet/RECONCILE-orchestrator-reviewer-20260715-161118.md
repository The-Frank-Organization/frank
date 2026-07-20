## RECONCILE -- operator scope direction stands, but the re-cut is a material architecture amendment and is not build-ready

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r7
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- the operator's smaller-MVP direction is accepted; this return requires architecture, ownership, and dependency corrections within that direction
GRILL_REQUIRED: no -- this is review-only; the replacement architecture delta and the m-9/m-10 owner designs carry their own review and grill gates
DESIGN_DOC_ID: step3-arch-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-152000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-5.implementer, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- preserve the smaller coding-agent MVP direction, but issue an explicit architecture amendment, close the ownerless allow-list and missing provider vertical, and correct the shared-client, wake, and source-fold boundaries before any interface lock or T4 token

VERDICT: revise

## What closes

- The operator may stand down the Step-3 m-5 freshness amendment and defer the canonical config-derived/per-role ceiling `643dd7c2...` to Step-4. F36-F38 are therefore OBE against that stood-down amendment. The later m-5 report-only stand-down `161000` agrees; no m-5 MVP amendment remains live.
- The ordered 15-file manifest recomputes exactly to `5374ee4ac6176126cd092a8967e41e270fed08e0279e2be9ff22feab7d8277dd`; packet r4 remains byte-exact at `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 remains `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- The code findings are directionally real. `cmd/frank-mcp` is a frontend over `channel.Client`; reusable schema/payload/reconnect behavior is command-local. Accepted-relay nudges already exist at `cmd/frank/main.go:337-353`, active-seat delivery at `internal/channel/server.go:167-179`, reconnect recovery at `cmd/frank/main.go:358-364`, and client receive at `internal/channel/server.go:510-518`.
- Incoming `152000` and the m-5 stand-down `152000` each exact-file lint `OK`. Root lint still reports the pre-existing INDEX/lineage debt before ending with those exact-file `OK` results.

## Findings

### F39 -- "packet r4 untouched" is byte truth, not architecture truth

The re-cut changes r4's normative architecture: r4 makes m-10 enforce an m-5-authored ceiling (`STEP-3-ARCH-AMENDMENT.md:27-29,65-71`), requires the m-10+m-5 coordinated first stage (`:102-112`), puts m-8+m-9 after that lock, and defines the MVP as one m-9-to-m-8 provider turn with live E3 (`:65-75,112`). Incoming replaces the ceiling, removes m-5, moves m-9 into first stage, and still advances directly to T4. Those are architecture changes even if the r4 file bytes do not move.

Preserve r4 as the historical lock, but author a separately hashed additive amendment that explicitly supersedes the affected r4 clauses and pins the new boundary/dependency graph. Master+VP review the exact bytes; the operator ratifies the final delta. Until then, the seven-file fold is proposed design input, not a new architecture-of-record, and no first-stage interface-lock can issue.

### F40 -- the "empty" allow-list is a real authority policy with no owner

A static list deciding whether `read/write/edit/bash/apply_patch` may execute is authorization. m-10 loading and enforcing it is an enforcement mechanism. Calling that both "the ENTIRE permission/authority system deferred" and a "permissive default" obscures the positive authority source, while r4 and the m-10 charter say m-10 owns no policy (`STEP-3-ARCH-AMENDMENT.md:28,103`; `m-10.../README.md:6,13,17`). "Fail closed absent/malformed" does not answer who writes the valid list, who may change it, or what binds it to a run.

The amendment must pin the coarse MVP policy honestly: authority source/owner, manifest writer and trusted load path, immutable run binding, exact tool identifiers, mutation/restart behavior, and deny behavior. If the operator's fixed five-tool decision is the policy source, say so and serialize exactly that reviewed set; m-10 may host enforcement but may not silently choose or widen it. Also narrow the product claim: with ceiling, sandbox, and audit deferred, local tool effects are not frank-governed; only relay communication is governed in this MVP.

### F41 -- the reduced sequence has no model/provider vertical

A coding-agent turn still needs a model path. The operative boundary sends m-9 through m-8, where m-8 owns provider translation, credentials, wire enforcement, and telemetry, with m-1/m-3 contracts (`STEP-3-ARCH-AMENDMENT.md:27-29,65-75`). Yet incoming keeps m-8/m-3 held, authorizes no credential or provider call, and says m-10+m-9 lock goes straight to T4. A build team cannot invent those owner contracts, and the promised live E3 cannot occur from m-9+m-10 alone.

Before a T4 code token, choose and review the actual provider boundary. If the MVP uses m-8, complete the minimum m-8 connector plus m-1 secret-boundary and m-3 pre-wire/evidence contracts and place them in the dependency graph. If it delegates the model turn to an existing Codex runtime, explicitly amend the m-8/m-9 boundary, credential custody, wire/evidence claim, and E3 proof. Do not call the MVP vertical or first-stage-to-T4 path complete while this branch is unstated.

### F42 -- the shared-client extraction crosses m-7 and m-2 ownership

`channel.Client.Call`/reconnect/socket behavior is m-7's conductor interface-guardrail surface. But `SchemaFromForm` and `SubmitPayloadFromArguments` encode m-2-owned FieldSpec/form semantics (`cmd/frank-mcp/schema.go:11-47,90-129`; charter m-2 row). m-9 owns the native model-facing tool consumer, not either upstream contract. Incoming directly routes neither owner and omits m-2 even from CC.

Route owner work before consumer lock: m-7 authors/reviews the shared conductor transport/client boundary; m-2 authors or explicitly approves the form-to-tool-schema and submit-payload mapping; m-9 consumes them for the native tool. Pin the package split so `internal/channel` does not absorb frontend/FieldSpec policy by convenience, and require parity tests proving the retained MCP frontend and native frontend produce equivalent conductor calls and re-render behavior.

### F43 -- existing push cannot directly wake an absent worker or m-10

The conductor pushes only to a currently authenticated connection whose stamped seat matches the relay recipient (`server.go:167-179`). That principal is m-9; r4 forbids m-10 from holding a submit credential (`STEP-3-ARCH-AMENDMENT.md:27-28,84-86`). The reconnect nudge is supplied only when m-9 later authenticates. Therefore existing push can notify a resident m-9 seat receiver, but it cannot itself start an absent m-9 process, and m-10 cannot call `NextPush` as m-9.

The optional wake design must choose the lifecycle: a resident m-9 seat-side receiver forwards a nudge over app IPC to m-10, which schedules a turn, or m-10 uses an explicitly non-principal fallback such as polling. Pin reconnect, coalescing, duplicate-nudge, and read-after-nudge behavior. Keep polling as the MVP fallback; do not claim "m-10 turns the conductor push into a turn" until this boundary is owner-reviewed.

### F44 -- the source fold contains mutually operative sequences

The new dashboard/register text says m-10+m-9 form first stage, but the same live sources still say m-10+m-5 interface-lock first and m-9 re-dispatches at stage 2: `m-10.../README.md:23,29-35`; `m-9.../README.md:37,39`; `m-5.../README.md:49-51`; and `master/README.md:9,150`. In particular, the newly appended m-9 paragraph itself ends with "stage 2," contradicting incoming `152000:41-42`.

After F39-F43 resolve, fold one explicit current dependency graph into every operative source and mark the exact old clauses superseded. Preserve historical relays and r4 bytes; do not leave contradictory current charter instructions for the owner seats.

## Required return

Return a bounded amendment candidate and corrected source fold, not another claim that r4 is architecturally unchanged. It must: name the MVP's honest governance boundary; assign the static allow-list policy/provenance; choose the provider vertical and owner gates; route m-7+m-2 ownership for the shared client; define the m-9-to-m-10 wake handoff; and present one non-contradictory dependency graph. Refresh the ordered manifest and request exact-byte VP review plus operator ratification.

No first-stage interface-lock, m-9/m-10 lock, T4 code token, DESIGN_LOCK, PLAN, implementation, credential, provider call, merge, or deploy is authorized by this review.

## Verification

- Exact ordered 15-file manifest and combined digest recomputation: match `5374ee4a...`.
- Exact-file relay lint: incoming packet `152000` and m-5 stand-down `152000` each `OK`; root-mode historical debt separately observed.
- Live trail read through INDEX row containing m-5 stand-down SITREP `161000`.
- `git -C frank status --short` empty; HEAD `502e06cc07b5`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended its INDEX row; no governing-source, packet, domain-design, historical-relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main clean at `502e06c`.
Next requested action: planner authors the explicit architecture amendment and owner-correct dependency graph, then returns its exact bytes and refreshed source manifest for VP review; all design locks and build authority remain held.
