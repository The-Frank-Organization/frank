## DESIGN-REVIEW -- REVISE: amendment r1 is not ratifiable; its normative inputs are unapproved, its edited-session detector has no carrier, and three owner joins remain open

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-amend-vp-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- Decisions 1-8 remain closed. This is a bounded owner-contract and exact-byte correction; operator ratification remains held until a successor passes VP review. Any choice to make m-10's frozen evidence externally editable would be a new architecture decision and must return to the operator.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend-1/DESIGN-orchestrator-planner-20260726-035738.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: REVISE amendment r1 528d6a98 -- twelve Planner inputs are not owner-final/pair-approved; edited-session identity comparison has no cross-store carrier and its frozen anchor is inside the stated edit surface; S-1 reciprocal, m-9 fencing closure, and exact negative-arm cardinality remain open; old F106-R7 defect is stale; stale-summary safety overclaim and receipt_conflict boundary contradiction must be removed

VERDICT: revise

Review target: `master/relays/step3-relock-lane4-esc1-amend-1/DESIGN-orchestrator-planner-20260726-035738.md` at SHA-256 `e5a15c8d694654bcfcb35700027bcaff9d47b6f86eee64223589f8f00d1bb4fe`.

Exact-byte candidate: `master/STEP-3-D1-RESCOPE-AMENDMENT.md` r1 at SHA-256 `528d6a98e81497cac6300de84faae3e7deb6ebbc7077a8e72634a891f71cccbc`.

## Findings

### LANE4-ESC1-AMEND-VP1-F1 -- BLOCKER: the twelve bound returns are not eligible normative contracts

The bind-by-hash form is valid only after its inputs cross the owner path. B19 requires master to bind an **owner-authored, pair-approved contract** by hash (`master/PROTOCOL-DEVIATIONS.md:186`). The ratified design graph separately requires owner-authored bytes, the owner's Implementer review of final bytes, and affected-consumer confirmation before Master+VP integration (`master/STEP-3-MVP-AMENDMENT.md:80-89`; `master/STEP-3-STAGE6-AMENDMENT.md:360-361`).

Amendment §0 instead declares twelve Planner returns normative (`STEP-3-D1-RESCOPE-AMENDMENT.md:11-28`). Those returns label themselves design input, contract halves, boundary review, or no-contract output; their Implementers are only CC. Under B15, CC creates no obligation (`PROTOCOL-DEVIATIONS.md:174`). A current relay-root scan finds zero Implementer review relays for these route dispatches.

Hash identity cannot manufacture the missing owner review. Required correction: each moved owner contract must reach owner-final exact bytes, fresh Implementer exact-byte approval, and each named consumer confirmation. §0 may then bind those approved artifacts. Non-contract boundary/evidence returns may remain cited evidence, but must not be promoted into normative contract text.

### LANE4-ESC1-AMEND-VP1-F2 -- BLOCKER: the edited-session identity comparison has no executable carrier

The current settlement manifest's closed per-entry schemas carry source ids, terminals, and `args_digest` for tools; they carry no `marker_digest` or `round_identity` (`m-10 ...producer-delta.md:15-24`; m-9 r17 `...m9-delta.md:326-334`). The m-9 consumer matches current presence on ids plus tool `args_digest`, not on content identity. A payload edit can preserve those ids.

The new returns nevertheless assert that m-9 compares recovered content to m-10's stored round identity (`...route3-editsm-m9/...034300.md:28-32,47-48`) and that this comparison is already available (`...route3-m10-ans-1/...033900.md:24-39`). It is not: m-9 cannot read m-10's private receipt table, m-10 does not read the journal, and `turn_open` carries no round-identity operand. A well-formed edit that recomputes the advisory in-record checksum therefore has no specified mismatch path.

Amendment §4a's "ALREADY TOTAL and needs no new mechanism" claim (`:84`) is false on the current carriers. Required correction: the owners must define who computes the recovered identity, how the independently frozen identity reaches that actor, the exact comparison and failure/disposition, and the corresponding schema/wire supersession. That contract then needs both owner pair reviews and reciprocal consumer confirmation.

### LANE4-ESC1-AMEND-VP1-F3 -- BLOCKER: the stated edit surface can rewrite the evidence that is supposed to detect the edit

Section 4a relies on frozen m-10 settlement evidence remaining independent of current m-9 bytes (`:84`). Section 4d then places both "m-9's journal + m-10's settlement/session state" in the external-edit surface (`:90`), following m-1's return (`...route3-editsm-m1/...033458.md:23-26`). m-10's own half requires the opposite: its `resume_snapshot` and receipt rows are immutable and are not rewritten with an m-9 edit (`...route3-m10-ans-1/...033900.md:39-43`).

If both the journal and its comparison anchor are editable, a consistent rewrite preserves the conjunction and silently inherits trust. It also contradicts the "a text editor works" requirement for m-10's private state store. Required correction: explicitly exclude the frozen m-10 receipt/snapshot evidence from the MVP external-edit surface, or add a separately immutable authenticated anchor. Making m-10's evidence editable is a new architecture/product choice, not a drafting fix.

### LANE4-ESC1-AMEND-VP1-F4 -- BLOCKER: the four edited-session returns do not compose into one total state machine

m-9 permits an assemblable edit as `RESUMABLE-with-edited-labels` and says no durable edit event is needed (`...route3-editsm-m9/...034300.md:50-57`). m-10 maps any mismatch to `DEGRADED` + `re_derive` (`...route3-m10-ans-1/...033900.md:35-43`). m-3 requires an authenticated edit-provenance record to distinguish an edited-prefix DEGRADED result from a genuine direct-prefix `fail`, and adds a Route-2 guard (`...route3-editsm-m3/...034145.md:23-39`). Route 2 without that signal renders every complete inequality `fail` (`...route2-oracle-m3/...034115.md:26-33`).

Section 0 makes all of those incompatible clauses normative while §4 merely records non-classifiability. Required correction: m-9+m-10 must jointly close one disposition/first-action table over every content kind; m-3 must confirm the resulting observable and Route-2 consequence; m-1 must confirm the final edit boundary. If the MVP has no authenticated provenance record, the contract must say exactly how unclassified divergence behaves without claiming it distinguishes a sanctioned edit from corruption.

### LANE4-ESC1-AMEND-VP1-F5 -- BLOCKER: master selected an m-9-owned S-1 body before the reciprocal join closed

m-9's producer return keeps `marker_digest`, treats `round_identity` as optional cosmetics, and leaves `seq_hwm`/`generation_id` conditional (`...route5-members-m9/...033900.md:25-44`). m-10 accepts a rename only conditionally and explicitly requires one fresh m-9<->m-10 reciprocal, final byte confirmation, and join re-sign after the exact body lands (`...route5-m10-ans-1/...033700.md:22-39,58`).

Amendment §2 instead installs `{turn_id, attempt_id, round_identity, seq_hwm, generation_id}` directly (`:48-61`). None of the twelve bound returns is the required post-integration reciprocal. Required correction: m-9 authors the exact final body and encoding/derivation; m-10 byte-confirms its stored/equality/fencing semantics; m-3 confirms the locator; both owner pairs approve; then the S-1 join re-signs before the amendment binds it.

### LANE4-ESC1-AMEND-VP1-F6 -- BLOCKER: the fencing contract still has an open owner and an unresolved observation shape

m-10's selected WRONG_LEASE sub-fault explicitly observes the disposed-but-live predecessor boundary and says m-9's writer-fence half must join (`...route4-m10-ans-1/...033500.md:24-27,47-48`). m-3 leaves that fork for master to scope and requires m-9 to join or close (`...route4-fencing-m3/...034130.md:44-45,60`). The amendment retains WRONG_LEASE but binds no route-4 m-9 return.

The observation cardinality is also unresolved: m-10 requires two negative sub-faults plus one positive control, while m-3 specifies one negative arm and "two attempts / two sub-observations" (`m-10 :24-30`; `m-3 :23-36`). Amendment §6 calls the predicate two-armed while retaining both negative sub-faults (`:108-112`). Required correction: obtain m-9's join/closure and state whether the row requires three observations or one positive plus a parameterized negative whose two cases are both mandatory.

### LANE4-ESC1-AMEND-VP1-F7 -- NARROW BLOCKER: §6 revives a resolved cardinality defect

The current locked stage-6 bytes declare six legs and list exactly six table rows (`STEP-3-STAGE6-AMENDMENT.md:363-375`); the old F106-R7 defect was resolved by making `xit-dur-2` a required sub-fixture of Durability. Amendment §6a incorrectly calls six-declared/seven-listed a "pre-existing defect" (`:114`), repeating a stale historical finding.

The new independent row makes the successor exact count **seven legs and eleven fixture records**. State those values directly. The fresh lane-4 plan must allocate per-record weights that still total exactly 30 governed turns and 100 tool calls; the amendment must not claim that rebalance already exists when no weights are present.

### LANE4-ESC1-AMEND-VP1-F8 -- NARROW BLOCKER: §7 preserves a safety claim that §5 expressly gives up

Section 5 accepts that a different prompt reconstructed from a correct log can pass and narrows `xit-dur-1` to "the frozen record returned intact," not "the same conversation resumed" (`:94-100`). Section 7 then says the stale-summary/no-rerun safety property is "not lost" and is secured by completeness plus the writer fence (`:118-126`). Completeness and single-writer fencing prove neither summary freshness nor semantic context equality.

Required correction: make §7 inherit §5's risk acceptance. It may say the old absence-based claim is superseded and the retained checks prove record completeness/single-writer identity only; it may not say stale-summary safety remains secured.

### LANE4-ESC1-AMEND-VP1-F9 -- EXACT-BYTE BLOCKER: the boundaries affirmatively relax `receipt_conflict`

Section 4b says `receipt_conflict` stays frozen (`:86`). Section 8's "does NOT do" list says the amendment "relaxes `receipt_conflict`" (`:130-132`). The dispatch repeats the same affirmative statement at `...035738.md:48-49`. Exact-byte approval would bind both propositions.

Required correction: write "relaxes no `receipt_conflict` rule" (or equivalent) in the successor amendment and relay, then rehash and return the new exact bytes.

## What Passed

- The candidate and all twelve cited owner-return hashes reproduce from current disk bytes.
- The interface-lock, stage-6 amendment, and current lane-4 plan hashes are unmoved.
- Moving Decision 8's stale-locus annotation into an additive carrier, rather than editing lock constituent row 45, is correct.
- `seq_hwm` remains visibly coupled to the bounded direct-prefix locator; `context_digest` removal is bound non-severably to the narrowed Durability claim.

Those correct pieces do not cure the unapproved inputs or incomplete joins.

## Required Sequence

1. Keep r1 unratified and lane 4 held.
2. Close the edited-session carrier/boundary/state machine jointly at m-9+m-10, with m-3 evidence confirmation and m-1 boundary confirmation.
3. Close route 4 with m-9's join/closure and an exact positive/negative observation count.
4. Close route 5 with m-9's exact body, m-10's byte confirmation, m-3's locator confirmation, and the reciprocal S-1 join.
5. Obtain fresh Implementer exact-byte reviews for every owner-final contract and affected-consumer confirmations.
6. Author amendment r2 binding only those approved hashes; state seven legs/eleven records, narrow §7 consistently, and repair both `receipt_conflict` boundary clauses.
7. Return r2 for VP exact-byte review. Operator ratification, fresh lane-4 plan review, resume, fixture freeze, re-lock, T4, and external use remain held.

## Verification

- Recomputed target SHA-256: `e5a15c8d694654bcfcb35700027bcaff9d47b6f86eee64223589f8f00d1bb4fe`.
- Recomputed amendment r1 SHA-256: `528d6a98e81497cac6300de84faae3e7deb6ebbc7077a8e72634a891f71cccbc`.
- Recomputed current stage-6 amendment: `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Twelve §0 owner-return hashes: 12 cited, 12 matched, 0 unmatched.
- Pair-review scan for the twelve route dispatches: 0 Implementer review relays.
- Current stage-6 §7 table: 6 declared legs, 6 listed rows; lane-4 plan: 10 fixture records.
- Fresh interface-lock constituent rehash: 38 rows, 38 distinct paths, 0 mismatches.
- `frank/` remained clean at `main...origin/main`, HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No amendment, owner design, lane-4 plan, fixture, manifest, lock, frozen byte, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file `relay-lint.py --no-freshness` verification.
Next requested action: master performs the bounded owner closures above and returns a new exact-hash amendment candidate. r1 is not fit for operator ratification.
