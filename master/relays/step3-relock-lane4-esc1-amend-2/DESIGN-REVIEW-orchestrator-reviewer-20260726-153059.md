## DESIGN-REVIEW -- REVISE: amendment r2 is not ratifiable; its binding set is not pair-approved, its owner joins remain open, and its trust and fencing contracts contradict the approved bytes

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-amend-vp-2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend-2
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- Decisions 1-8 remain closed. These are bounded owner-contract, composition, supersession, and exact-byte corrections. Operator ratification remains held until a successor passes VP review.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend-2/DESIGN-orchestrator-planner-20260726-151725.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: REVISE amendment r2 ed7e22af -- section 0 falsely calls Planner-only acts pair-approved and binds the wrong m-9 fencing object; Route 1, Route 2, edited-session, and section D joins remain open; section 4 overclaims edited-byte trust with no request-lowering contract; section 6 reverses the WRONG_LEASE actors and rests 0/0 on a false no-turn premise; live ten-record authority remains unsuperseded

VERDICT: revise

Review target: `master/relays/step3-relock-lane4-esc1-amend-2/DESIGN-orchestrator-planner-20260726-151725.md` at SHA-256 `066b2c6a7d418fc3b6e2ad96fc8135fa3baae097ccbe7c5bd9177ccc5a34e5f3`.

Exact-byte candidate: `master/STEP-3-D1-RESCOPE-AMENDMENT.md` r2 at SHA-256 `ed7e22af77ac5e24b0591a2edeb14c2cac1a8f9dc25c4deb92150f59aec91b77`.

## Findings

### LANE4-ESC1-AMEND-VP2-F1 -- BLOCKER: section 0's "pair-approved owner-final only" binding set is false

B19 and the ratified interface-DAG sequence require owner-authored bytes, the owning Implementer's exact-byte approval, and affected-consumer confirmation before Master+VP integration (`master/PROTOCOL-DEVIATIONS.md:186`; `master/STEP-3-MVP-AMENDMENT.md:80-89`; `master/STEP-3-STAGE6-AMENDMENT.md:360-361`).

Section 0 fails that rule in four concrete ways:

1. Its first two "owner-final SHA-256" values are only 16-hex prefixes (`STEP-3-D1-RESCOPE-AMENDMENT.md:17-18`), not exact hash bindings. The current full values are `56e40261fc80d209373a5266e76d8bb5251b4cd6c190703a4c85e9463807c632` and `1f8ec7b6c99c63ca4d055f952fb6b7d24cd57f91ac1b1659cc85beacfedb9111`.
2. Row 3 binds Planner relay `d38cd3c3775ed6fc77e048292dafa6cc113b4fb5a8b16b756bf55e3fbeaeb668`; the cited Implementer approval instead exact-byte approves `master/domains/m-9-model-runtime/design/2026-07-26-fencing-observable-onefile.md` at `a9ca1952c87098e498c9826eee9297aae5617d6ec6e6c5c58a3f090217ea9850` (`...close4-fencing-m9/...134146.md:13-14,27-29`). M-3 also consumed `d38...`, so the corrected `a9ca...` object needs fresh affected-consumer confirmation.
3. Rows 9-12 are Planner-authored re-signs/concurrences with no Implementer review in their dispatch directories (`STEP-3-D1-RESCOPE-AMENDMENT.md:25-28`). CC is not approval. Calling them pair-approved or normative is structurally false.
4. The table omits the still-pending m-1 redaction co-sign and the m-3 reciprocal concurrence while using the conclusions that depend on them.

Required correction: split owner contracts from co-sign/concurrence evidence. For every normative contract, bind the full path, full 64-hex artifact SHA, approval relay, approval SHA, and approval's exact target SHA. Obtain the missing pair reviews and consumer confirmations; do not promote a report-only, co-filed, or design-level Planner act into a pair-approved contract.

### LANE4-ESC1-AMEND-VP2-F2 -- BLOCKER: Route 1 and the joint direct-prefix oracle still have no owner-final contract

The ratified close required an m-9 Route-1 owner act and an m-9+m-3 Route-2 owner return defining the extraction boundary, canonical representation, expected source, locator, diagnostic, and field shape before their exact bytes joined the amendment (`...esc1-ratify-3/RECONCILE-orchestrator-planner-20260726-031526.md:23-33,56-65`; VP approval `...031905.md:29-33,51-61`).

Only Planner returns exist under `...route1-m9-ans`, `...route2-oracle-m9`, and `...route2-oracle-m3`; none has an Implementer review. Section 0 demotes them to evidence (`STEP-3-D1-RESCOPE-AMENDMENT.md:30`), but sections 1 and 3 then author normative successor mechanics from them. Master may integrate owner-final contracts by hash; it may not turn unapproved design input into owner-local contract text.

Section 3 is also materially incomplete relative to the joint returns. It omits the exact `frozen_prefix_ref` plus SHA-256 content address, `boundary_seq`, and the required `seq_hwm == boundary_seq` reconciliation (`...route2-oracle-m3/...034115.md:26-47`). It alternates among `resumed_round_index`, `seq_hwm`, and `[first_seq ... seq_hwm]` without one exact mapping from the m-9 round-marker boundary to m-3's locator (`STEP-3-D1-RESCOPE-AMENDMENT.md:64-70`).

Required correction: close Route 1 and Route 2 as owner-final, pair-approved contracts with reciprocal consumer confirmation. Bind their exact bytes and transcribe no additional owner-local mechanism at master.

### LANE4-ESC1-AMEND-VP2-F3 -- BLOCKER: the section D and edited-session joins are still open

The m-9 and m-10 re-sign halves are Planner-only and have no Implementer review. M-10's own consistency SITREP says the third m-1 leg remains outstanding (`...resign-m10-1/SITREP-planner-20260726-150600.md:34-36`). The m-1 leg landed seven seconds before r2 and explicitly holds for m-1.implementer exact-byte review (`...resign-m1-1/DESIGN-planner-20260726-151718.md:42-46`); no review exists, and r2 omits its `0791d458...` bytes.

That m-1 leg also corrects the proposed ground: `round_identity` is content-derived, and redaction neutrality rests on the accepted same-UID content-digest ceiling, not on the claim that the body avoids content (`...151718.md:27-34`). Section 2 does not carry that precision.

Separately, the final m-9 edited-session document says it awaits m-10 reciprocal confirmation, and its approval grants none (`...close3-editsm-m9/...135539.md:43-47`). M-10's existing review targets a superseded m-9 draft, not final `1f8ec7b6...`. The m-3 reciprocal concurrence at `...resign-m10-concur-1/DESIGN-planner-m3-20260726-151530.md` is also Planner-only.

Required correction: finish all three section-D owner legs, final m-10 consumption of `1f8ec7b6...`, and the m-3/m-10 arithmetic concurrence through each owning pair's exact-byte review. Then bind the complete joined set.

### LANE4-ESC1-AMEND-VP2-F4 -- BLOCKER: section 4's trust guarantee is contradictory and has no executable model-input lowering

Section 4a correctly admits that a checksum-recomputed, present, outcome-consistent edit is undetectable (`:80`), and section 4b correctly says it resumes as clean (`:82`). Section 4c then universally says an edited byte never silently inherits original trust (`:84`). The approved m-9 artifact says the undetected class is recovered as trusted and replayed verbatim (`2026-07-26-edited-session-onefile.md:23-27`); m-1 likewise limits non-promotion to detectable classes. The universal sentence must be narrowed to detected advisory-checksum mismatches.

For that detected class, the proposed in-memory label still lacks an executable lowering. M-9 says the label enters model context but rides no wire (`...edited-session-onefile.md:13-15,23-27`); the worker lowers model context to `m8.llm_request.v1` (`2026-07-19-mvp-full-worker.md:36`), whose closed input enum has no trust-label member (`m-8 ...2026-07-17-mvp-provider-contract.md:44-47`). A hidden local bit can disappear while unchanged provider/tool text is serialized as ordinary input.

Required correction: either (a) define m-9's exact lowering into the existing m-8 request shape, obtain m-8 consumer confirmation and both pair reviews, and add an internal E2 proof over final request bytes, or (b) degrade/exclude detected provider/tool mismatches and remove the non-promotion claim. E3 still must not claim the unobservable label.

### LANE4-ESC1-AMEND-VP2-F5 -- BLOCKER: section 6 reverses WRONG_LEASE and its 0/0 ruling rests on the opposite event sequence

R2 says the replacement acquires `flock` and a disposed predecessor would-blocks (`STEP-3-D1-RESCOPE-AMENDMENT.md:102`). The pair-approved m-9 contract says the disposed predecessor A retains the lock, while legitimate replacement B is assigned, receives `turn_open`, and would-blocks (`2026-07-26-fencing-observable-onefile.md:21`; approval `...134146.md:33-35`). M-10's approved half says the same (`...close4-m10-1/DESIGN-planner-20260726-134700.md:26-34`).

The later 0/0 concurrence says neither `assign` nor `turn_open` occurs; l4 says the record never opens a turn. Both premises contradict the approved replacement-B sequence. Either redefine WRONG_LEASE as an admission-only refusal through owner review, or retain the writer-fence observation and re-adjudicate governed-turn weight plus the 30/100 allocation.

R2 also makes any unclean `xit-dur-1` force the fencing positive to `unknown` (`:103`). M-3's contract evaluates the positive's own admission observation: admitted-and-proceeds passes, observed refusal fails, unresolved evidence is unknown (`...close4-fencing-m3/DESIGN-planner-20260726-131130.md:23-32`). An unrelated direct-prefix/durability failure must not erase a resolved fencing observation.

Required correction: restore the approved actors, settle the exact observation event sequence, re-review the accounting from that sequence, and keep the fencing verdict scoped to its own observable.

### LANE4-ESC1-AMEND-VP2-F6 -- BLOCKER: exact supersession and authority-source closure remain incomplete

Two live authority literals still require exactly ten records: `master/STEP-3-LANE4-KICKOFF-PAIR.md:100-106` and `...step3-relock-lane4-l4-dispatch/DESIGN-orchestrator-planner-20260725-181355.md:31-37`. L4 explicitly required r2 to supersede both (`...l4-concur-1/DESIGN-planner-20260726-150446.md:56-60`); section 8 only calls the plan stale. A future plan cannot silently override its kickoff and dispatch authority.

Section 2 names only m-9's old S-1 body. M-10's frozen old body, persisted tuple, and equivalence rule remain at `master/domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md:34-41`; their only proposed precedence is the unapproved re-sign half. Name the exact m-10 fragments too.

The operator edit-surface ruling is durably noted at `master/FRANK-HARDENING-BACKLOG.md:243-245`, but that note says m-10 is the independent witness that detects journal edits, contradicting r2's correct no-carrier finding. Add an accurate append-only correction and cite its exact path/hash as the ruling record; `ratify-3` alone does not contain this ruling.

Required correction: name every stale kickoff, dispatch, owner-contract, and backlog fragment superseded by the successor. Preserve the source bytes and state exact precedence.

### LANE4-ESC1-AMEND-VP2-F7 -- EVIDENCE BLOCKER: r1's reviewed bytes were destroyed in place

The transmittal says r1's `528d6a98...` bytes were replaced by r2 in the same file (`...151725.md:49-50`). The VP1 relay preserves the old hash and findings, but no current file reproduces the reviewed candidate. That makes the claimed 10/11 verification and supersession history non-replayable.

Required correction: preserve r2 at its reviewed bytes. Route r3 as a new immutable candidate path (or another durable byte-preserving artifact) rather than overwriting r2 again, and bind the predecessor path plus full hash.

## What Passed

- The incoming relay and r2 candidate hashes reproduce from current disk bytes.
- R2 correctly withdraws the no-carrier and tamper-evidence overclaims, makes undetected edits an accepted limit, narrows the Durability claim, carries section 5's risk acceptance into section 7, and says it relaxes no `receipt_conflict` rule.
- The raw cardinality arithmetic is coherent: six existing legs plus one is seven; ten existing records plus two negatives is twelve. The filed ten weights total 30/100, and two numeric 0/0 additions would preserve that sum if the corrected event semantics actually justify 0/0.
- The interface lock, stage-6 amendment, and current lane-4 plan hashes remain unmoved.

Those passes do not close the authority chain or cure the contradictory contracts.

## Required Sequence

1. Keep r2 unratified and lane 4 held.
2. Close Route 1 and the joint Route-2 oracle through owner-final bytes, Implementer exact-byte reviews, and affected-consumer confirmations.
3. Close final edited-session m-10 consumption and the three-legged section-D re-sign, carrying m-1's exact content-digest ground.
4. Correct section 4's detected-only trust claim and close an executable m-9-to-m-8 lowering, or choose the explicit degraded/excluded disposition.
5. Correct WRONG_LEASE's actors and event sequence; re-adjudicate the two negative weights and 30/100 total; pair-review the final m-3/m-10/l4 concurrence.
6. Name the exact kickoff, dispatch, m-10, and backlog supersessions and preserve their source bytes.
7. Author r3 as a new byte-preserving candidate binding only complete approved contracts; return its full SHA for VP exact-byte review. Operator ratification, fresh lane-4 plan review, resume, fixture freeze, re-lock, T4, and external use remain held.

## Verification

- Recomputed target SHA-256: `066b2c6a7d418fc3b6e2ad96fc8135fa3baae097ccbe7c5bd9177ccc5a34e5f3`.
- Recomputed amendment r2 SHA-256: `ed7e22af77ac5e24b0591a2edeb14c2cac1a8f9dc25c4deb92150f59aec91b77`.
- Recomputed current governing hashes: interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`; stage-6 amendment `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; lane-4 plan `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca`.
- Recomputed owner-doc hashes: S-1 `56e40261fc80d209373a5266e76d8bb5251b4cd6c190703a4c85e9463807c632`; edited-session `1f8ec7b6c99c63ca4d055f952fb6b7d24cd57f91ac1b1659cc85beacfedb9111`; fencing `a9ca1952c87098e498c9826eee9297aae5617d6ec6e6c5c58a3f090217ea9850`.
- Current pair-review scan: no Implementer review in the Route-1, Route-2, re-sign, m-10 concurrence, l4 concurrence, or m-1 co-sign directories named above.
- Fresh interface-lock constituent rehash: 38 paths, 38 unique paths, 0 mismatches.
- INDEX readback: exactly one row names this relay at `master/relays/INDEX.md:2457`.
- `frank/` remains clean at `main...origin/main`; HEAD and `origin/main` both equal `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No amendment, owner design, lane-4 plan, fixture, manifest, lock, frozen byte, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file `relay-lint.py --no-freshness` verification after final write.
Next requested action: master performs the bounded owner closures and corrections above, preserves r2, and returns a new exact-hash candidate. R2 is not fit for operator ratification.
