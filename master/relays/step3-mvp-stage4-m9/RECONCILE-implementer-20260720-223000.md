## RECONCILE — F92 CERTIFIED + P4 COUNTERSIGNED; P5 NOT COUNTERSIGNED — HUMAN DECISION REQUIRED because the proposed pin withdraws an operator-ratified L5=B product effect

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — P5 changes the operator-sourced L5=B digest-independence/rejected-A semantics recorded in the durable GRILL_LOCK; F92 and P4 are independently closed
GRILL_REQUIRED: no — this relay does not reopen the design itself; it routes the conflicting operator-ratified semantics back to the operator
DESIGN_DOC_ID: step3-mvp-design-m9-worker
IN_REPLY_TO: master/relays/step3-mvp-stage4-m9/RECONCILE-planner-20260720-221500.md
FROM: m-9.implementer
TO: operator
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-9.planner, m-10.planner, m-10.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-mvp-stage4-m9/RECONCILE-implementer-20260720-223000.md
SUBJECT: independent return on VP F92/F94/F95 — CERTIFY worker-r7's r19 citations revision-neutral against half r21; COUNTERSIGN P4 postbuild timing; DISSENT on pair-level P5 because complete-output inclusion withdraws the operator-ratified L5=B digest-independence and rejected-A effect, so operator disposition is required

Operator — I independently checked the exact artifacts and the ratified sources rather than adopting the planner's conclusion. The result is split: **F92 certifies; P4 is countersigned; P5 requires your decision and is not countersigned by this seat.**

## F92 — CERTIFIED: worker r7's r19 lifecycle citations are revision-neutral against approved half r21

I certify that worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` may remain byte-bound unchanged alongside lifecycle half r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; its header's r19 citation `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c` is revision-neutral at every mechanism it realizes.

Independent exact-byte basis:

- The historical r19 bytes remain available at `frank/master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md` and hash exactly to `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`; the live r21 bytes hash exactly to `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- An exact zero-context diff has only six hunks: header Status; owner-basis line; one added §2.2 `turn_open.admission_ref` consumer note; §5's reciprocal census; §7's consumed-hash/gate; and the two r20/r21 fold-log entries. No other line changes.
- Lifecycle §1 is byte-identical across r19/r21 at section SHA-256 `cd19b5632b0cd8b1436defbcf14dcfa6d876bd8b222e9a2731dca78b20c71bff`. Lifecycle §3 is byte-identical at section SHA-256 `23b6934bb3562df47cf7c5ddc8abbfafee70f561570f04692a1a8b5b5b09a25f`. §§4 and 6 are also byte-identical.
- In §2, r21 only adds the `admission_ref` presence/shape/epoch consumer note. Every pre-existing §2 byte cited by the worker remains unchanged. The added note explicitly assigns objective acquisition to worker §7.1/E16, which r7 already authors against m-10 r40; it neither replaces nor contradicts any worker realization.
- Worker citations reach lifecycle §1/§1.1/§1.2, §2.2/§2.3/§2.5/§2.6/§2a, §3/§3.2, and §4's deferral. They do not consume the r21-only §5 reciprocal bookkeeping or §7 review/gate bookkeeping. The one r21 addition inside §2.2 is complementary and points to the worker's own acquisition mechanism.

Therefore a metadata-only worker r8 is unnecessary for F92. This certification is my own adversarial act and may be cited by the corrected stage-6 packet; the planner's relay is not a substitute for it.

## F94 / P4 — COUNTERSIGNED

> **P4 — digests bind postbuild only; stage 6 binds ONLY the interface identity contract + the expected catalog vector. Actual build digests exist only at the postbuild RELEASE-BINDING.**

**Implementer concurrence.** The ratified `master/STEP-3-MVP-AMENDMENT.md` `2f75f2a1…` is unambiguous at §4 line 59 and §7 stage 6 line 87: stage 6 precedes T4 and binds only the interface identity contract plus expected catalog vector; exact app-main/m-10, worker, connector, and shared-client artifact digests bind after T4 at RELEASE-BINDING before E3. Worker r7's `:105`, `:136-137`, and `:227` incorrectly place actual T4-produced digests at stage 6. P4 corrects an impossible timing/conformance statement; it leaves the expected `tool_catalog_digest = 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4` stage-6-lockable and preserves the two-layer postbuild verification mechanism.

Although r7 `:227` labels the mistaken timing “source: operator,” it conflicts with the already-ratified architecture's causal order: a pre-T4 gate cannot bind not-yet-produced T4 output bytes. I therefore countersign P4 as architecture conformance, not a new product choice.

## F95 / P5 — NOT COUNTERSIGNED; operator decision required

The proposed pin is:

> **P5 — `m9_worker_build_digest` = the complete runnable worker output INCLUDING linked bytes; the shared conductor-client gets a SEPARATE component/material digest for attribution; the release binding covers BOTH; the "changes iff the worker's own code changes" claim is WITHDRAWN.**

The implementation fact behind P5 is correct: for the intended in-process Go linkage, the complete runnable worker output includes the linked client bytes, so its SHA-256 changes when those bytes or other build inputs change. A materials list supplies attribution; it cannot subtract linked bytes from the output digest.

But this is not revision-neutral to the operator-ratified L5=B choice recorded in worker r7's durable GRILL_LOCK:

- §8.1 `:102-103` defines B as “separate artifact, referenced as a material,” says the worker digest changes **iff the worker's own code changes**, and promises that a shared-client patch does not drag the worker through re-review.
- GRILL_LOCK `:226` marks L5=B as `source: operator`.
- Rejected alternative `:235` expressly rejects L5 option A — folding client bytes into the worker's digest — because the worker identity would change for reasons outside the worker.
- P5 makes the complete worker digest cover those linked bytes and explicitly withdraws the `iff` property. A separate client component digest preserves attribution and avoids “single digest only,” but it does not preserve the ratified dependency-independent worker-identity effect.

The planner's narrower claim — that B bought attribution only and digest independence was merely a false corollary — is contradicted by the GRILL_LOCK's explicit end-app effect and rejected-alternative rationale. This seat therefore cannot countersign P5 as a pair-only semantic pin.

## Operator decision requested

Please choose the product semantics:

1. **Accept implementation-faithful P5:** amend L5=B to mean the complete runnable worker digest covers linked bytes, while a separate shared-client component/material digest supplies attribution and the postbuild release binding covers both. The dependency-independent `iff` claim is withdrawn; a client change necessarily moves the runnable worker artifact digest, though review scope may still be narrowed using the separate component evidence.
2. **Preserve the original dependency-independent L5=B effect:** require a real artifact boundary or distinguish a dependency-independent worker-component/source identity from the F63 complete-runnable-output digest. This needs owner design bytes rather than the proposed P5 pin alone.

Until that decision is recorded, P5, the corrected stage-6 packet, and downstream PLAN/T4 remain held. F92 and P4 need no further m-9 work. Worker r7 and lifecycle r21 bytes remain unchanged and approved; this relay moves no design or code byte.

## Verification

- Direct handoff relay SHA-256: `20a884e7c25dfb8554c7de5a7985274d121cb2075fdf65b0525ddccdfc33b1d3`; exact-file lint: OK.
- Worker r7 SHA-256: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Lifecycle r19 archived SHA-256: `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`; lifecycle r21 live SHA-256: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Ratified MVP amendment SHA-256: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Exact r19→r21 diff hunk locations: old/new `-6/+6`, `-9/+9`, added `+91`, `-261/+262`, `-312/+313`, and added `+339,+340`.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace disk action only — created this F92/P4/P5 companion RECONCILE relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage4-m9/RECONCILE-implementer-20260720-223000.md`.
Next requested action: operator chooses whether to accept P5's implementation-faithful complete-output semantics or preserve the original dependency-independent L5=B effect; master then folds the decision together with this completed F92 certification and P4 countersign into a fresh lock packet.
