## DESIGN-REVIEW — rev2 MUST REVISE: R1 settlement/caller findings close, but app-main recovery still contradicts m-10's E+1 lifecycle and control loss makes the bounded cut non-total

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r2
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-172150.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact re-review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev2 at SHA-256 `dcfed9d0ab654b890b1a724008da267b1d4bf940ef0784f8bab41ae93ebe4a2d` — R1-F1/R1-F2 close, but two pre-existing cross-contract races block approval

## Verdict

**MUST REVISE.** Rev2 correctly folds both r1 findings: relay calls remain on the uniform F59/m-10 settlement path; the state-sensitive terminal/UNKNOWN/VOID map is exact; the existing `parked_unknown` + D2 `uncertain` carrier is mandatory; rediscovery is identity-informed with a fresh ticket; and caller delivery is conditional. No crossing row, transition ledger, retained completion, second outcome carrier, or H-24 is required by that fold.

Approval still cannot issue because the current bytes contradict the frozen m-10 app-main recovery sequence and contradict themselves about when the bounded drain cuts under control loss.

## Closed From R1

- **R1-F1 CLOSED:** §§Q3.2.2, Q3.4, D, C, R.5 and FX-TB-5′ now distinguish conductor relay records from m-10 outcome rows and consume the existing settlement path without a relay carve-out.
- **R1-F2 CLOSED at the normative §Q3.4/fixture surface:** typed caller delivery is conditional; dead-connection discard and mandatory m-10 successor disclosure are explicit. R2-F3 below is residual consistency cleanup, not a reopening of the corrected mechanism.

## Findings

### R2-F1 — BLOCKER — Q1/FX-TB-20 restores old-epoch authority after app-main death, contradicting frozen m-10 and Q2

**Current bytes:** §Q1.2 says same-epoch in-flights “complete and deliver normally” during the app-main outage. FX-TB-20 says adoption “restores authority with NO epoch change.” §Q2 simultaneously says worker replacement is exactly E→E+1.

**Frozen lifecycle:** m-10 r40 §B.3 says app-main EOF makes both worker and connector fail closed and exit. On restart, a durable `LEASED` generation with no committed retirement takes the one keyed retirement transaction, which parks consumed/no-outcome rows and mints E+1 exactly once. Its common suffix explicitly separates broker-control adoption from epoch installation and requires the broker's E+1 install before successor assign. m-10 §B.4 fixes `RETIREMENT ⇒ new epoch`; adoption itself is not an epoch trigger, but full recovery of a leased active generation necessarily contains the retirement/E+1 transition.

**Failure:** rev2 conflates “control adoption does not itself mint” with “app-main recovery restores old-E authority.” In the ordinary leased-generation crash, the old worker is dead, old E is revoked by retirement, and no same-E successor authority may be restored. A response racing after the worker connection dies also cannot “deliver normally”; the consumed/no-outcome identity is parked UNKNOWN by the keyed retirement path.

**Required correction:** distinguish broker-control adoption from generation recovery and make Q1 consume the lifecycle-state matrix:

1. During the outage, the broker stays bound but suspended; child EOF makes the worker/connector exit. Delivery of an old in-flight response is conditional on the admitting connection still being alive, never universal.
2. Replacement app-main adopts broker CONTROL without rebind, re-auth, credential movement, or an adoption-caused epoch mint.
3. For the ordinary `LEASED`/not-yet-retired case, the keyed retirement transaction runs exactly once, parks F59 rows state-sensitively, mints E+1, and the Q3 broker transition installs E+1 before successor assign/attach.
4. Preserve no-mint only for m-10's already-current branches (`RETIRED_PENDING_REAP`, pre-lease wash-out, initial E=1) exactly as their matrix defines; do not generalize it to app-main recovery.
5. Replace FX-TB-20 with branch-exact legs, including the common leased-generation E→E+1 case and the no-double-mint crash-before/after-retirement cut. Keep “no channel rebind/no re-auth” independent of epoch movement.

### R2-F2 — BLOCKER — “cut-at-reestablishment” contradicts the broker-local bounded deadline and leaves late-response disposition undefined

**Current bytes:** §Q3.1 says the broker starts a compiled-constant drain at PROPOSED, terminates every unresolved operation at expiry, and then installs atomically. §Q3.3.1 likewise says install occurs at drain end. But FX-TB-17′ says control loss during PREPARING causes “cut-at-reestablishment,” and recovery starts an “ordinary barrier” over whatever old-E in-flights remain.

**Failure trace:** broker enters PREPARING and starts the deadline; control drops; the deadline expires while control remains absent; a conductor response arrives after expiry but before control re-establishment. The normative rule says the operation was already cut, while the fixture/recovery wording defers the cut and can treat the same response as an in-window completion or restart another full drain. The cut linearization, response delivery, and boundedness are therefore not total.

**Required correction:** pin one broker-local cut linearization independent of controller availability:

1. The deadline continues through control loss. At its expiry the broker irrevocably cuts every unresolved old-E operation; every later response is discarded even if E+1 install/event delivery must wait.
2. State explicitly whether an already-accepted proposal installs at that local expiry while control is absent, or whether installation remains suspended until verified control returns. Either choice must preserve the expiry cut and must not revive E or restart a full drain.
3. Re-proposal after control return is idempotent over the already-cut state; it cannot create a second cut, second disposition, or new drain for those identities.
4. Add the exact fixture: control lost before expiry, response arrives after expiry/before re-establishment, response is not delivered or buffered, m-10 holds exactly one state-sensitive disposition, and install/ack recovery converges.

This correction stays within the simpler rule set. It does not require retained completion, crossing rows, a transition ledger, or H-24.

### R2-F3 — IMPORTANT — Two stale caller/event proofs remain, and the mandatory-carrier claim overreaches `Describe`

- §Q3.3.6 still says the clean cut reports through “typed error to the caller + `boundary_cut` event,” despite §Q3.4 retracting unconditional caller visibility.
- §R.2 still rejects broker-local durable memory because “caller + event + conductor record” carry the cut, omitting the now-mandatory m-10 outcome/disclosure carrier and reusing the invalid proof set.
- §§Q3.1/Q3.3.3 say the durable m-10 outcome carrier covers every cut operation, while §Q3.4 correctly says `Describe` is outside the eight-tool/F59 path and has no relay tool row. The m-10 carrier therefore covers the three relay tool calls, not `Describe`.

Replace both stale proof sentences with the durable m-10 row/disclosure proof. Scope the m-10 mandatory-carrier statements to F59 relay tool calls. Then state `Describe`'s separate cut/recording posture explicitly: it is retryable metadata outside item-D settlement, and its `boundary_cut` event retains the already-admitted uncoupled/dual-failure telemetry residual. If “complete-or-reject, recorded” is still claimed universally for `Describe`, explain what durable record satisfies it; do not cite a nonexistent m-10 `tool_calls` row.

## Re-review Gate

Return one byte-exact rev3 that:

- preserves the now-correct R1 settlement and conditional-caller text;
- distinguishes control adoption from the lifecycle-state-dependent E+1 retirement path;
- makes the drain-expiry cut linearization total across control loss;
- removes the two stale caller/event proof sentences;
- scopes the mandatory m-10 carrier to F59 relay tool calls and states `Describe` separately; and
- adds branch/race fixtures for app-main death and deadline-before-control-return.

The accepted clean-cut/no-H-24 direction, F67 custody, F64 fence, capability-mint adoption, CI-3 shrink, and CI-4 survival placement remain open to approval once these exact contradictions close.

## Verification

- Reviewed rev2 at exact SHA-256 `dcfed9d0ab654b890b1a724008da267b1d4bf940ef0784f8bab41ae93ebe4a2d`; incoming relay at exact SHA-256 `5e8fad133987895547832711e227027b2a8ecd1b0c7e565213bdab312ce710bf`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 full worker `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; ratified amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`; root-mode output retains unrelated historical `INDEX.md`/lineage noise.
- No design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R2-F1/F2/F3 into one exact rev3 and relays it for byte-exact adversarial re-review; m-9/m-10 confirmation routing and re-lock remain held on rev2.
