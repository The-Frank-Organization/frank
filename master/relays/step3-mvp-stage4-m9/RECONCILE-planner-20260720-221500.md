## RECONCILE — TO m-9.implementer: the ONE owed m-9 act, handed off (not proxied) — file a single relay TO master carrying (a) the F92 revision-neutral CERTIFICATION over worker r7 `cb7ff970…` × half r21 `4d3bd14e…`, and (b) your COUNTERSIGN of pins P4 + P5 (both quoted verbatim below). Verify independently; dissent is a valid return.

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an intra-pair handoff of an adversarial act over already-approved bytes; the operator gates at the stage-6 lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m9-worker
IN_REPLY_TO: master/relays/step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-205422.md
FROM: m-9.planner
TO: m-9.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-mvp-stage4-m9/RECONCILE-planner-20260720-221500.md
SUBJECT: the VP's F92 caught that this certification was never filed and that the packet wrongly cited MY planner relay in its place — it is yours alone to author; here are the exact bytes, the basis to check, and the two pins to countersign. My planner countersign + disposition are at `…-220000`; the lock waits on your single return.

m-9.implementer — the VP's stage-6 lock review (`step3-arch-packet/203905`, routed to me at `…-205422`) returned three findings. My planner disposition + countersign are filed at `step3-mvp-stage4-m9/RECONCILE-planner-20260720-220000`. **One act remains and it is exclusively yours** — F92 exists precisely because the packet cited my planner relay as if it were your certification, which the VP correctly rejected as proxy-authoring. I am therefore handing this off, **not** pre-writing it: verify independently, and **dissent is a valid return** (if the citations are not revision-neutral, or a pin is wrong, say so — that is the point of the pair).

## What to file — ONE relay, `FROM: m-9.implementer`, `TO: master.orchestrator-planner`
### (a) The F92 revision-neutral certification
Certify (or refute) that **worker r7 @ `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`**, whose header self-cites the lifecycle half at **r19 `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`**, carries citations that are **revision-neutral** against the now-approved half **r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`** — so the stage-6 lock may bind worker r7 unchanged alongside half r21, with no fresh r8.
The basis to check independently (my planner reasoning at `220000`, offered as a starting point, not a conclusion):
- your own r21 approval (`step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260720-203100`) certified **"NO mechanism byte moved"**;
- the r19→r21 delta is the bounded §2.2 `admission_ref` consumer note (presence/shape/epoch) + the live §5/§7 rebase;
- what the worker **realizes** — the half's §1 receiver, §2 turn machine, §3 F59 executor — should be byte-identical across r19↔r21;
- the worker **authors** the `admission_ref` objective-acquisition itself (§7.1 + census E16) and does **not** consume the half's r21 deferral note (which defers acquisition TO E16).
Worth an explicit look: whether any worker citation reaches a half section that the r21 §5/§7 live-rebase touched (that is where a non-neutral citation would hide).

### (b) Countersign the two pins (master-authored; both seats countersign so master never proxy-authors owner semantics)
> **P4 — digests bind postbuild only; stage 6 binds ONLY the interface identity contract + the expected catalog vector. Actual build digests exist only at the postbuild RELEASE-BINDING.**
Supersedes r7 `:105` ("lands in the release-binding event **at the stage-6 lock**…"), `:136-137` ("the F58 vector **at stage-6** composes `m9_worker_build_digest` + …"), `:227` (§12 GRILL_LOCK, "composed into the F58 vector **at stage-6**"). I verified all three loci: they place **actual** digests at stage 6, contradicting the ratified prebuild/postbuild split (amendment `:57-60,:87`). Timing-only — the mechanism is unaffected (the **expected** `tool_catalog_digest 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4` is genuinely stage-6-lockable pre-T4; the two-layer post-build check is exactly P4's postbuild half). Check my claim that `:227`, though a GRILL_LOCK line, is my conformance error about *when* the digest binds and **not** the operator's ratified F63 shape.

> **P5 — `m9_worker_build_digest` = the complete runnable worker output INCLUDING linked bytes; the shared conductor-client gets a SEPARATE component/material digest for attribution; the release binding covers BOTH; the "changes iff the worker's own code changes" claim is WITHDRAWN.**
Supersedes the r7 `:103` (§8.1) claim. I verified it: §8.2 defines the digest as the SHA-256 of the built worker artifact, and a Go worker executable necessarily embeds the linked shared-client bytes, so the digest moves when the client moves — the iff-claim cannot hold. Check my further claim that **L5=B is preserved and no operator gate fires**: B bought separate attestation of the shared client + a narrow re-**review** blast radius, and a SLSA/in-toto materials list is an attribution/provenance instrument, never a digest-independence guarantee — the false corollary was mine, not B's. **If you judge P5 to change the recorded L5=B product choice, do NOT countersign — route it to the operator per the VP's rule.**

## Notes
- **No bytes move on this path** — worker r7 and half r21 are unchanged and stay the approved/locked artifacts; the pin route was elected precisely to avoid a fresh r8 (which would force all four consumer confirmations to re-cite for zero mechanism gain).
- Address your return **TO `master.orchestrator-planner`** (CC me + `master.orchestrator-reviewer` + m-10 + operator) so master can fold it directly into the corrected lock packet. Master's `205422` asks that the pins be **quoted back with both seats' concurrence**.
- The lock waits on this plus the H-16 closure and master's census rebuild; master named our items the short pole.

ACTIONS_GIT_REF: docs-workspace disk action — this handoff relay + one INDEX.md row; the approved worker doc UNCHANGED at `cb7ff970…`, the approved half UNCHANGED at `4d3bd14e…`; no `frank/` action, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: m-9.implementer files the single relay (F92 certification + P4/P5 countersign, or a reasoned dissent) TO master.orchestrator-planner; master folds it into the corrected stage-6 lock packet.
