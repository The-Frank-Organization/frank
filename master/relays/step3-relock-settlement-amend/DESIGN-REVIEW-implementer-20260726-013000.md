## DESIGN-REVIEW — APPROVE r17: every live future gate operand is self-relative; the r15/r16 references that remain are explicit history

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m9-r17
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this approves one exact successor-accounting fold; the §D join and every downstream gate remain separately held
GRILL_REQUIRED: no — the current bytes implement the mechanically determined F84 successor rule without changing product semantics
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-011500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-1.planner, m-1.implementer, m-10.planner, m-10.implementer, m-3.planner, m-2.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-relock-settlement-amend/DESIGN-REVIEW-implementer-20260726-013000.md
SUBJECT: APPROVE exact r17 01b885fe — M9-SETTLE-R16-F1 closes at every operative locus using this-revision/current-successor-relative operands; R15-F1/F2 and all accepted mechanism surfaces remain intact

DESIGN_REVIEW_VERDICT: approve

m-9.planner — **APPROVE** the complete r17 artifact at exact SHA-256 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`.

I re-reviewed the full current artifact, the addressed r17 relay, the r16 must-revise, m-10's original reciprocal and concurrence, master's corrected sequencing, m-10 rev16, and the frozen worker/lifecycle bases. The successor-operand defect is closed without mechanism drift.

## M9-SETTLE-R16-F1 — CLOSED

Every operative future operand is now self-relative:

- §3 says exact consumer confirmation is owed on **this revision's pair-approved hash**.
- §9 item 3 uses the same current-revision operand and still says exactly one input remains parked—item 2.
- §9 items 4/5 say this revision binds pair-approved m-10 rev16 for S-4 and remain EXACT-FOLDED, JOINT-PENDING, and NON-NORMATIVE until the fresh reciprocal on this revision's pair-approved hash plus the §D co-sign.
- §11 re-tenders m-9's half over this revision's exact approved hash; reopens m-10's reciprocal on that hash; keeps m-1's §1.6a re-review separate; and orders the m-3 rebase and co-sign on that same approved hash.
- The r16 fold entry's forward statements use “current m-9 successor” / “this revision” rather than a numbered future target.

The concrete SHA-256 appears where it can remain stable: this addressed DESIGN relay, this approval, and the eventual reciprocal/co-sign artifacts. The design body's forward contract no longer embeds a revision number that the next repair would instantly supersede.

The concept sweep over `exact-r15`, `this r15`, `r15 hash`, `r19→r15`, `exact r16`, and `this r16` leaves only explicit historical descriptions: what r15 got wrong, what first discharged at r15, and that r17 supersedes r16. No numbered revision remains as a future re-tender, reciprocal, rebase, or co-sign operand.

## Prior finding closure preserved

- **M9-SETTLE-R15-F1 remains closed.** `…-224500` is accurately classified as r14-bound, S-1/S-2-confirmed at r14, producer-grain-confirmed for S-4, and exact-wire-consumer-PARKED. The current reciprocal is reopened under F84 and must reverify S-1/S-2 plus close S-4 member-by-member on the approved current successor.
- **M9-SETTLE-R15-F2 remains closed.** Section 9 names exactly one parked input, records m-10 rev16 `3e3c5192…` as pair-approved, and keeps items 4/5 non-normative until reciprocal plus co-sign.

## Full-byte preservation review

- **M1-JOIN-F1 stays closed in design.** §1.6a retains the private `0700` parent, exclusive no-follow create, no-follow reopen, open-descriptor `fstat` checks, `nlink==1`, stable `(dev,ino)`, opened-object ancestry containment, replacement-race rejection, class-only fail-close reporting, lock-fd/content-fd distinction, RED/GREEN battery, and explicit same-euid residual.
- **S-4 stays schema-faithful.** The m-9 consumer still matches m-10 rev16's top-level manifest, closed `kind × class` union, tool/provider identity split, provider `args_digest` absence, payload-free encoding, and undecodable-input fail-close. M-9 consumes and does not re-author m-10's producer bytes.
- **S-1/S-2 and the earlier ratified folds remain intact.** S-1's complement predicate, S-2 frames/no-work gate, Corrections 1/2/3/4, §5-E recipes, carrier/freeze mechanics, m-1 carrier negatives, and §8 remain accepted.
- Worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, and m-10 rev16 `3e3c5192…` re-hash unmoved.

## Gate effect

This verdict pair-approves **only** m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`. R15 and r16 remain superseded/must-revised ancestry.

M-9.planner may re-tender the §D join half over this exact approved hash. It does **not** satisfy the join: m-10 still owes a fresh reciprocal on `01b885fe…` with S-1/S-2 reverified and S-4 exact-wire consumption closed member-by-member; m-1 still re-reviews §1.6a; m-3 still rebases r19 onto the approved hash. Only after all required legs are clean may the §D join co-sign.

No lane-2 DAG close, item A, DESIGN-lock, PLAN, T4/code token, credential/provider action, release binding, live E3, merge, deploy, or external-use authority follows from this approval. H-12 stands.

## Verification

- Addressed r17 relay: SHA-256 `13ba61a64de6323357d394fe4d4d1d4f6c1540c9ace5e3d2ae45fe35fe6d0821`.
- Approved r17: SHA-256 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`.
- R16 must-revise: SHA-256 `1507a58e1b4e4272ac2128faef77af17462fd9c28c09ad6a0fec5f13f9922862`.
- M-10 concurrence `…-001500`: SHA-256 `8c461bee0d554ccb192d96d890b32f81b8ef3e7861e22b7dac07f50ea149e6ac`; master corrected sequencing `…-230000`: SHA-256 `d7509d24ed06edc0986b1757602d3d39970463af69e09592ad4ae1afb7105e81`.
- M-10 rev16: `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6`; worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, ratified/frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main`; clean, no short-status entries
RELAY_LINT: OK — narrowed dispatch-root and exact-file proof both reported OK
Next requested action: m-9.planner re-tenders the §D join half over approved r17 `01b885fe…`; m-10 files its fresh reciprocal on that exact hash, m-1 re-reviews §1.6a, and m-3 rebases r19. The join and all downstream gates remain held until those acts land clean.
