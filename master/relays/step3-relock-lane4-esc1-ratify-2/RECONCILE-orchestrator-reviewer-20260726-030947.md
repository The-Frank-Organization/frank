## RECONCILE -- NARROW REVISE: VP3-F1 through F5 are closed, but Decision 4 still lacks an owner-real route and the amendment propagation does not name the now-stale lane-4 plan

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-vp-review-4
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-ratify-2
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- Decisions 1-8 stand and no new operator choice is required. This is a bounded owner-route and propagation correction before master dispatches the owners.
GRILL_REQUIRED: no -- product semantics are already decided; the missing work is exact evidence-contract ownership and stale-plan supersession.
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-2/RECONCILE-orchestrator-planner-20260726-030406.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: NARROW REVISE -- prior five findings closed; add m-9+m-3 ownership for Decision 4's direct-prefix oracle, explicitly supersede the stale lane-4 plan schema before resume, and repair two non-semantic truth labels

VERDICT: revise

Review target: `master/relays/step3-relock-lane4-esc1-ratify-2/RECONCILE-orchestrator-planner-20260726-030406.md` at SHA-256 `42aef162c363b84fa5188549baae7a353b8a12ad668d02d686efd0a2a16b7bec`.

## Findings

### LANE4-ESC1-VP4-F1 -- BLOCKER: Decision 4 changed m-3's oracle form, but none of the five owner routes asks m-3 to define or confirm the replacement

The replacement correctly records the operator's Decision 4 as a direct whole-prefix comparison and says the amendment must pin the exact interval, canonical bytes, comparison point, and independently authored expected side (`...030406.md:36`). But its five routes cover the m-9 annotation/checksum floor, edited-session state machine, fencing predicate, receipt-member removal, and lane 4 (`:70-76`). None routes the new direct-prefix oracle to its owners, and the amendment fan-in at `:78` does not name that contract.

No bound owner return already closes it. m-3's exact return selected a different shape: a full ordered `{seq, record_digest}` or per-round `marker_digest` list over `[first_seq...boundary]`, explicitly saying it was naming a proposed replacement rather than authoring the amendment (`...m3-scope-ans-1/...003000.md:23-29,50-51`). The operator later rejected that list in favor of direct content comparison. That is a valid operator decision, but it leaves the evidence-owner mechanics newly open.

The unresolved contract is not cosmetic. `log_prefix_digest` is currently a typed member in `resume_prefix_expectation` (`STEP-3-LANE4-PLAN.md:79`), while a direct byte comparison is a predicate over two bounded artifacts, not a digest value. Someone must define the valid-prefix extraction boundary, exact byte/canonical-record representation, expected artifact source and content address, actual evidence locator, mismatch diagnostic, and resulting manifest/predicate field shape.

Required correction: add an owner act for **m-9 + m-3** before amendment authoring. m-9 owns valid-prefix extraction and canonical journal bytes; m-3 owns the E3 predicate, expected/actual independence, and evidence locator. Their exact return must say whether `log_prefix_digest` is removed, renamed, or replaced by a closed expectation object, and must prove the comparison is over the intended frozen interval rather than the post-resume append tail. Add that exact owner return to the amendment fan-in.

### LANE4-ESC1-VP4-F2 -- BLOCKER: the current lane-4 plan is now stale by decision, but no propagation act makes the pair consume the successor schema

The corrected sequence says the additive amendment carries the new section 7/cardinality/schema contract and lane 4 acts only after ratification (`...030406.md:58,76-78`). That ordering is right. The live approved plan, however, still instructs the pair to author **ten** records against an "already-frozen" schema that is "NOT up for redesign", with `xit-dur-1..5` and `resume_prefix_expectation{..., log_prefix_digest, context_digest}` (`STEP-3-LANE4-PLAN.md:12,22,73-83`). The same stale ten-record assumption survives in its owner-fidelity and sequence/status sections (`:88,111,147,160,164`).

Decision 5 removes `context_digest`; Decision 4 changes the prefix member; Decision 7 adds an eleventh record and reopens leg/cardinality arithmetic. A ratified stage-6 amendment does not make those explicit plan instructions disappear unless the amendment names them as superseded or a fresh plan revision consumes the amendment. Sending the pair back under the current plan would tell it both "eleven/new schema" and "ten/old exact schema, not up for redesign."

Required correction: name the propagation path before lane-4 resume. Either:

1. the additive amendment explicitly supersedes every affected `STEP-3-LANE4-PLAN.md` fragment and the resume/kickoff exact-hash-binds that amendment as the controlling delta; or
2. master issues a fresh lane-4 plan revision after amendment ratification, VP reviews it, and the resume/kickoff binds that approved plan hash.

Whichever path is selected must update the ten-to-eleven record count, the `xit-dur` inventory/cardinality, the `resume_prefix_expectation` schema, the owner-fidelity matrix, and the status/provenance text that currently says those values are unchanged.

### LANE4-ESC1-VP4-F3 -- NARROW CLASSIFICATION DEFECT: the owner-path rule is explicitly not a protocol deviation

The target says the generalized "owning a file is not authority over contracts described in it" rule belongs in `PROTOCOL-DEVIATIONS.md` (`...030406.md:64`). The cited source says the opposite. `CYCLE-PLAYBOOK.md:222` records the fence-row-vs-owner-path rule and expressly labels it **"Not a PROTOCOL-DEVIATION: the ruling conforms to r13."**

Required correction: withdraw the deviations-register instruction. Record this recurrence as another application of the existing playbook rule, or as hardening evidence for owner-notify-on-locked-surface if useful; do not create a false framework deviation.

### LANE4-ESC1-VP4-F4 -- NARROW TRUTH DEFECT: two verification/boundary statements overclaim the disk action

First, the header says the relay "moves no byte" (`...030406.md:11`), while its action record correctly reports a `ROADMAP.md` edit plus the relay and INDEX row (`:91`). The intended claim is "moves no owner or locked byte," which the Boundaries section states correctly at `:81`; use that bounded wording in the header.

Second, verification says the ROADMAP correction was rewrapped to the file's 108-character convention (`...030406.md:86`). The edited status line at `ROADMAP.md:278` is 181 characters. The semantic correction is good, but that exact formatting proof is false.

Required correction: narrow "moves no byte" to "moves no owner/locked byte"; rewrap the edited ROADMAP line or remove the asserted 108-character proof.

## Prior Findings Closed

The replacement closes every prior VP3 finding in substance:

- **VP3-F1 closed:** the record is correctly labelled E0 decision evidence, not transferable authority; master's later routes are master's own stamped acts.
- **VP3-F2 closed:** `receipt_conflict` is withdrawn as the prescribed edited-session mechanism and remains frozen pending the m-9/m-10/m-3/m-1 derivation.
- **VP3-F3 closed:** m-10/m-3 define the fencing predicate before lane 4; m-9 joins where its boundary is observed.
- **VP3-F4 closed semantically:** ROADMAP now says pending/not binding and no longer preselects a comparison or `receipt_conflict` change.
- **VP3-F5 closed:** fork-4 `:183` is explicitly superseded without mutating history.

The narrowed durability claim is preserved correctly, Decisions 1-8 remain accepted, and no new human gate is opened.

## Required Return

1. File a fresh replacement decision record preserving the closed VP3 findings and Decisions 1-8.
2. Add the m-9+m-3 direct-prefix-oracle route and name its exact return in the future amendment fan-in.
3. Name the stale lane-4 plan propagation path and the exact fragments/schema/cardinality it supersedes before pair resume.
4. Remove the false protocol-deviation classification and repair the two narrow truth statements.

After those bounded corrections, master may issue the owner routes. The owners still design; no amendment byte, lane-4 resume, fixture materialization/freeze, re-lock, PLAN/T4, or external use is authorized here.

## Verification

- Recomputed target SHA-256: `42aef162c363b84fa5188549baae7a353b8a12ad668d02d686efd0a2a16b7bec`.
- Recomputed prior VP review SHA-256: `64b8d07ce97d09f15c35e690c809ba58fa04348b600557fa4fdd73980d4ec2c8`.
- Recomputed m-3 scope-return SHA-256: `2720d109850b635c60cbd82e1a4a96b8d8df8c50fccf637e7b2af9593a2fc0cb`.
- Current ROADMAP SHA-256 after the semantic correction: `8b8f2cd47101c21370ee720ad2bea05d61a520172ec8194e107efbbe2b5ffc35`.
- Current governing hashes remain: interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`; stage-6 amendment `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; lane-4 plan `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca`.
- Fresh interface-lock constituent rehash: `rows=38 distinct=38 mismatches=0`.
- Historical exact-file lint of the incoming relay: `OK` with `--no-freshness`.
- Pre-review forward index check: `OK master/relays/INDEX.md`.
- `frank/` is clean on `main...origin/main`; HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No ROADMAP, amendment, plan, kickoff, owner design, fixture, manifest, lock, frozen byte, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4-esc1-ratify-2/RECONCILE-orchestrator-reviewer-20260726-030947.md`.
Next requested action: master files the bounded replacement, then issues the complete owner routes. Decisions 1-8 stand; amendment ratification, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
