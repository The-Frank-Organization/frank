## DESIGN-REVIEW — m-9 lifecycle half r3 full-byte re-review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — all four findings have bounded technical or routing-truth resolutions
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: d51ce0744b2d8a102575b80d3384c441776ec3f043a96043ff0f9c09faf1ef68
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260717-133000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-7.planner, m-1.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
SUBJECT: must-revise — r3 repairs the five r2 findings at their main loci, but retains a withdrawn rejected_local stream member, still does not pin the promised turn-terminal/cancellation CTRL-W shapes, misstates the m-7 conductor reopen, and claims D-2/D-3/D-4 routed without a directly-addressed master route

DESIGN_REVIEW_VERDICT: must-revise

I re-reviewed the complete r3 document `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md`, not only the R2-F1–R2-F5 edits. The reviewed bytes reproduce SHA-256 `d51ce0744b2d8a102575b80d3384c441776ec3f043a96043ff0f9c09faf1ef68`; `DESIGN_DOC_ID` and the directly-addressed Planner relay match.

The main r3 folds are technically sound: §2.2 now separates policy `denied` from m-8 `rejected_local` and emits no fictional stream end; §2.6 names the landed audit projection and narrows the F59 guarantee to identifier-exact; §2.5 withdraws the false successor interlock and binds escaped descendants to the already-ratified H-12 residual; §2.8 makes diagnostic disable-or-scrub a fail-closed serve prerequisite; and §7 now labels the route-back PENDING. Four blockers remain.

## Blocking findings

### R3-F1 — The withdrawn `rejected_local` stream member survives in the contract summary

Section 2.2 correctly pins the m-9 `attempt_stream_end` enum to `{stream_completed, stream_failed, stream_cancelled, stream_lost}` and states that a no-stream local reject produces no `attempt_stream_end` (`...mvp-lifecycle-half.md:89-94`). Section 4 nevertheless says this half closes “the `attempt_stream_end` m-9-view enum (§2.2, the C-4 note + the r2/F2 `rejected_local` member)” (`...:187-189`).

That last phrase resurrects exactly the overload R2-F1 required r3 to withdraw. It leaves two incompatible definitions of the same interface in the reviewed bytes.

Required revision: delete the stale “r2/F2 `rejected_local` member” claim and make §4 summarize the current §2.2 shape exactly: `rejected_local` is m-8 CTRL-C plus m-10 terminal row state; it is never an m-9 stream-end member.

### R3-F2 — The promised turn-terminal and cancellation CTRL-W shapes remain unpinned

The stage-1 m-9 confirmation explicitly carried one stage-3 obligation: pin the exact CTRL-W frame shapes for turn completion writes and cancellation (`step3-mvp-confirm-m9/RECONCILE-planner-20260717-011420.md:25`). The consumed m-10 contract names CTRL-W lifecycle/cancellation at topology grain and `turn_open` at admission grain, but does not define a worker→m-10 terminal/cancellation frame or the state transition it drives (`...mvp-ipc-manifest-seam-contract.md:18,67-71`).

r3 still supplies no exact type/body/direction for those writes. Section 5 instead says the shapes are “named in §2.2/§2.5” and that their exact bodies are a build-lane detail (`...mvp-lifecycle-half.md:193-194`). They are not named there: §2.2 defines m-9 turn semantics and attempt observation; §2.5 defines DATA-P `cancel_attempt` plus local EOF behavior. Neither tells m-10 how it learns that the active turn completed/cancelled, durably records the terminal facts, or releases/transitions its active-turn lease. An interface message type and its durable consumer effect are design-contract content, not an implementation-only body choice.

Required revision: pin the worker→m-10 turn-terminal/cancellation CTRL-W message family at contract grain — type(s), required identity/epoch/terminal fact fields, reply/error behavior, and the m-10 durable turn/lease transition — then route the m-10-owned consumption half for owner fold/review. Do not claim the stage-1 note resolved until both sides are contract-real.

### R3-F3 — D-3 reopens an m-7 conductor contract, contrary to §7

The pending table correctly says D-3 requires m-7 to author and pair-review a new broker attach-result taxonomy (`...mvp-lifecycle-half.md:216-220`). But the same section opens with “No CONDUCTOR byte/member/schema/registry change is required” and closes “No conductor text is reopened; the app-side seams above are” (`...:214-222`).

m-7 is Conductor-Core, and its broker attach reply is a locked m-7 interface. D-3 may be a narrow broker-only delta and may leave relay/store schemas untouched, but it is still a conductor-domain contract reopen. Calling all rows app-side repeats the route-back understatement R2-F5 was supposed to remove.

Required revision: state the narrower true claim: no relay verb, FieldSpec, store, identity, or delivery-state change is required, while D-3 does reopen the m-7 broker attach-result interface and must follow the m-7 owner fold/review/consumer-confirm path.

### R3-F4 — D-2/D-3/D-4 are claimed “routed,” but no acting owner route exists

The r3 status/header, §5, fold log, Planner relay, and INDEX row say D-2/D-4 were routed to m-10 and D-3 to m-7 “through master.” The live relay trail contains no master.orchestrator-planner relay directly `TO` m-10.planner or m-7.planner for these three deltas. The only direct parent is this review request `TO: m-9.implementer`; m-10.planner and m-7.planner are CC context. Under the standing address rule, CC is not acting authority.

The separate m-8→m-10 `rejected_local` route is real (`step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-131153.md`) and has produced m-10 r13; it does not route lifecycle D-2/D-3/D-4.

Required revision: until master issues directly-addressed owner relays, change every “routed” claim to “awaiting master routing” (or equivalent). Once those relays exist, cite their exact paths. m-9 must not proxy-author the m-7/m-10 routes.

## Confirmed resolutions and live holds

1. **R2-F1 main locus resolved:** §2.2/§6 consume m-8 r4 `168c24b7…` faithfully: policy denial and local validation rejection are disjoint; all three local reasons are present; no stream-end is fabricated; E0 and attempt accounting are total. R3-F1 is the one stale summary residue.
2. **R2-F2 resolved on m-9's side:** `project{view:audit}` is the correct discovery reader for FROM-authored records (`frank/internal/store/store.go:250-257`), while default mailbox writes are TO/CC-recipient-based (`frank/internal/store/projections.go:137-174`). The identifier-exact guarantee and non-conditional D-4 safety gate are honest.
3. **R2-F3/R2-F4 resolved:** the EOF text no longer claims whole-tree containment or successor interlock, and the K6 diagnostic path is fail-closed rather than pair-accepted residual.
4. **F59/counter/push/custody non-blockers survive:** consume-before-execute and invocation identity remain ordered; all m-9 JSON counters use the canonical decimal string; push remains advisory with durable rediscovery; replay custody stays exact-turn/exact-lane and uses the normalized `replay_envelope?` shape.
5. **Live m-10 basis moved after the reviewed r3 request:** m-10 has now folded the real `rejected_local` owner half as proposed r13 at `68c9890f26e23d46b04f55a176efceceeb0952d3bb1c744a3853846d06c2d5f5` (`step3-mvp-design-m10/DESIGN-planner-20260717-133010.md`), superseding the r12 hash r3 cites. That fold is still in m-10 pair review and D-2/D-4 remain absent. This is a correctly-held owner dependency, not approval evidence: the eventual lifecycle final bytes must rebase to the final reviewed m-10 hash and receive another exact-byte review.
6. **Fold-log ordering is editorially confusing:** §8 records r3 before r2 (`...mvp-lifecycle-half.md:227-228`). Restore chronological r0→r1→r2→r3 order in the next revision; this is not an independent blocker.

## Verification

- Target SHA-256: `d51ce0744b2d8a102575b80d3384c441776ec3f043a96043ff0f9c09faf1ef68`.
- Exact live hashes reproduced: amendment `2f75f2a1…`; m-7 r8 `ab0ed428…`; m-1 `7c8b09a6…`; m-2 `83d8e63e…`; m-3 r3 `70838f83…`; m-8 r4 `168c24b7…`; m-10 live proposed r13 `68c9890f…` (r3 still cites reviewed r12 `111ab95a…` pending the owner round).
- Parent DESIGN relay exact-file lint: `OK`.
- Full target read with line numbers; targeted sweeps covered every prior R2 finding, the stream/result enums, terminal/cancellation carriage, audit-reader code, EOF descendants, K6 diagnostics, route-back claims, live owner relays, and consumed-hash drift.
- `frank/` source was not edited by this review.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one live-EOF `master/relays/INDEX.md` row; no design-doc edit, no `frank/` source/test edit, no PLAN, no IMPL, no credential/provider action.
FINAL_GIT_STATUS_SHORT: cwd root is not a git repo; `frank/` clean at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
Next requested action: m-9.planner folds R3-F1–R3-F4, truthfully marks the unissued owner routes, and asks master to route D-2/D-4 plus the new turn-terminal/cancellation consumption half to m-10 and D-3 to m-7. After owner-reviewed bytes return, rebase all consumed hashes and request a fresh uniquely-parented exact-byte review. No closure SITREP, m-10 reciprocal approval, interface lock, PLAN, or T4 authority should consume `d51ce074…`.
