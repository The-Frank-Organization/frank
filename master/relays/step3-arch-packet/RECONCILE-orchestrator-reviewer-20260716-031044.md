## RECONCILE -- r5 closes F57-F62 and the grill, but two exact-contract gaps block ratification of e47d514d

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r13
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the replacement hash still requires operator-authored ratification after a clean exact-byte review
GRILL_REQUIRED: no -- the three operator decisions are accepted and remain pinned; F63/F64 are contract-completeness repairs that must not reopen topology or Option B
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-030247.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW -- r5 closes F57-F62 and the three-decision grill, but E3 omits the app/worker build identities and the broker contract does not fence stale workers from conductor authority

VERDICT: revise-narrow

Review target: `master/STEP-3-MVP-AMENDMENT.md` r5 at SHA-256 `e47d514d3b87082dc9881e6b024aeb18a718aab3a671e59f2a7ce195fe426b72`, planner transmittal `030247`, grill records `023557` / `024350` / `025642`, the current m-1/m-7 ownership sources, and the live authenticated-channel implementation and lifecycle tests.

## Findings

### F63 -- the pre-build interface lock cannot bind the actual implementation, and E3 omits two binaries that execute the governed turn

The amendment now correctly distinguishes the eight-name policy identity from a build identity that does not yet exist (`STEP-3-MVP-AMENDMENT.md:56`). But it binds `{name, schema digest, implementation/catalog version, mapping version}` at the Master+VP **first-stage interface-lock** (`:58,86`) -- explicitly before the PM-to-T4 build. That event can lock the identity schema and expected interface/catalog contract; it cannot prove which implementation bytes T4 later builds. Nothing requires the implementation/catalog version to be mechanically derived from, or checked against, the shipped tool registry and executable.

The E3 tuple then binds only `m-8_build_digest` (`:48-49`). The observed governed turn also executes in the app-main/m-10 binary and the m-9 worker binary: those bytes create the manifest, epoch, authorization ticket, tool catalog, provider request, and evidence carrier. Either binary can change while every currently-bound digest remains equal. The annex nevertheless claims that changing "binary" invalidates the old E3 (`:117`), which the tuple cannot establish.

Required correction: make stage 6 bind the **interface identity contract / expected catalog vector**, then add a post-build release-binding event before the live E3 and exit gate. The run manifest and E3 must bind exact digests for at least the app-main/m-10 artifact, m-9 worker artifact, and m-8 connector artifact, or one canonical release/bundle digest that transitively and reproducibly covers them; include the conductor/shared-client artifact wherever the observed claim depends on it. Mechanically derive or verify `tool_catalog_digest` against the shipped registry so implementation drift without an identity/version change fails closed. Update the annex to mutate each bound artifact and prove the prior E3 becomes non-applicable.

### F64 -- broker-held seat custody does not yet fence a stale worker from the conductor channel

The process vocabulary is internally inconsistent: the preserved rail says the m-9 app seat "holds its OWN" credential (`STEP-3-MVP-AMENDMENT.md:15`), the topology defines m-9 as the supervised worker process (`:38`), and the resolved seat contract says the broker holds the credential and worker generations never receive its bytes (`:91`). The intended distinction appears to be **logical m-9 seat component** versus **m-9 worker process**, but the normative text must say so explicitly.

More importantly, `turn_epoch` is bound to provider attempts, tool authorization/execution, cancellation, and completion writes (`:92`), not to calls through the broker. The acceptance row checks old-epoch requests only at m-8 and the executor (`:112`). In the live conductor, the credential is authorized once when the channel connects and the resulting tool surface then serves `submit`, `project`, and `read` without an app epoch (`frank/internal/channel/server.go:277-336,391-417`). The one-active-channel lifecycle tests prove credential-channel exclusion and reconnect recovery, not worker-generation fencing. A resident broker can therefore preserve a valid conductor channel while an old worker keeps invoking its IPC unless the broker API independently revokes that generation.

Required correction: define the broker as credential holder for the **logical m-9 seat**, while the supervised m-9 worker receives only an epoch-bound, revocable use capability. The m-7 broker contract under m-1 identity semantics must check the current worker generation / `turn_epoch` on every `relay.submit`, `relay.project`, and `relay.read` operation and on push delivery/forwarding; specify the epoch-change linearization and in-flight-call disposition. m-10 may supply lifecycle/epoch state to that gate but must still receive no credential bytes and gain no conductor verb. Extend the stale-worker acceptance proof so an old epoch cannot invoke any of the three relay verbs or receive/forward a new push after replacement. This preserves the operator-locked one-credential-per-logical-seat choice; it does not re-grill it.

## Accepted Closures

- F57 closes: the packet address-space phrase is exactly superseded and the non-injection claim/test now match.
- F58 closes as to policy/build identity separation, named producers, applicability, and the later lock direction; F63 is the narrower missing post-build artifact binding.
- F59 closes: Option B, ticket tuple, atomic one-shot consume, crash disposition, split ownership, and actual-invocation proof are pinned.
- F60 closes as to the one-credential-per-logical-seat choice and m-1/m-7 ownership; F64 is the missing generation-fence contract at the broker boundary.
- F61 and F62 close: wake semantics now match the live best-effort path, and evidence scope/writers/applicability plus the automatic typed-route boundary are explicit.
- Grill records `023557`, `024350`, and `025642` are accepted as durable operator decisions. No further amendment-level grill is required for this return.

## Required Return

1. Fold F63 and F64 only, preserving all accepted r5 decisions and the F57-F62 closures.
2. Return r6 at a fresh SHA-256, update the README pointer, recompute the ordered 15-file manifest, and identify the exact r6 transmittal in the ratification clause.
3. Request another exact-byte review. Any byte change voids review of the prior candidate.

No ratification recommendation, source fold, first-stage dispatch/interface-lock, `DESIGN_LOCK_ID`, PLAN, T4 code token, implementation, credential provisioning, provider call, external send, merge, deployment, or live-store mutation is authorized by this review.

## Verification

- Amendment r5 SHA-256 independently recomputed: `e47d514d3b87082dc9881e6b024aeb18a718aab3a671e59f2a7ce195fe426b72`.
- Ordered 15-file governing manifest independently recomputed: `e4e1af2634e0ea79730f01c3d7389e5297f51c61833729449fa06130b451887f`; README exact at `2445011dcc9b95f6c712b77953426f91f53fb7da6b9c1112c3ff3fa4930a164f`.
- Reframe packet remains exact `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 artifact remains exact `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- Incoming `030247` and all three grill relays exact-file lint end `OK`; root-mode historical/INDEX lineage debt remains separate.
- Live channel inspection confirms connect-time credential authorization, one active connection per credential, and no per-call app epoch. Targeted `TestSecondConnectSameCredentialRejected|TestProvenDeadRecovery` passes; those tests do not prove worker-generation revocation.
- `frank/` remains clean on `main@502e06cc07b5` (`s11-close`).
- New reviewer relay exact-file lint ends `OK`; its INDEX row is present exactly once at live EOF.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-031044.md` and appended its `master/relays/INDEX.md` row; no amendment, governing source, historical relay, `frank/` source, branch, commit, lock, merge, live-store, credential, provider, external-send, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main remains clean at `502e06c`.
Next requested action: master.orchestrator-planner folds F63/F64 only into r6, refreshes the exact hashes and manifest, and requests a fresh exact-byte review; all ratification and build authority remain held.
