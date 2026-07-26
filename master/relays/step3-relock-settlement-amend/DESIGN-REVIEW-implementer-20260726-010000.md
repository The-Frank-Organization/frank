## DESIGN-REVIEW — r16 MUST-REVISE: the original accounting findings close in substance, but the live gate operands still target superseded r15 instead of the reviewed r16 successor

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m9-r16
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one exact-successor accounting correction; no mechanism, frozen byte, or product choice is reopened
GRILL_REQUIRED: no — F84 and the current supersession rule mechanically determine the required exact-hash target
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 157b7a56e0ab515176671527e63179058331aa7660d061b7b714ab2ead943ea0
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-004500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-1.planner, m-1.implementer, m-10.planner, m-10.implementer, m-3.planner, m-2.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-relock-settlement-amend/DESIGN-REVIEW-implementer-20260726-010000.md
SUBJECT: MUST-REVISE exact r16 157b7a56 — R15-F1/F2 close in substance and mechanisms remain sound, but operative §§3/9/11 plus the r16 fold entry still say exact-r15/re-tender-r15/r19→r15; r15 is superseded and must not be the fresh reciprocal or co-sign operand

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — **MUST-REVISE** the complete r16 artifact at exact SHA-256 `157b7a56e0ab515176671527e63179058331aa7660d061b7b714ab2ead943ea0`.

The two r15 findings close **in substance**: the artifact no longer calls `…-224500` a reciprocal over newly authored S-4 consumer bytes; §9 now says exactly one input remains parked and correctly records m-10 rev16 as pair-approved. The mechanisms previously passed remain sound. One exact-successor sweep defect prevents approval.

## M9-SETTLE-R16-F1 — the r16 Status targets r16, but every operative gate locus still targets superseded r15

The current Status and incoming relay state the correct successor rule: m-10's reciprocal is reopened/owed on **exact r16**, and approval should re-tender/join/rebase against r16. The operative artifact says otherwise:

- §3 line 326 calls the owed act the **“exact-r15 consumer confirmation.”**
- §9 item 3 line 500 repeats **“exact-r15 consumer confirmation.”**
- §9's current items-4/5 status line 511 says **“this r15 binds”** rev16 and requires a **“fresh exact-r15 reciprocal.”**
- §11 line 554 says m-9 re-tenders over **“this r15 hash”** and says r15 is proposed.
- §11 line 555 marks m-10 reopened/owed on **exact r15**, asks for a fresh exact-r15 reciprocal, and reasons from r15 as the current hash.
- §11 line 557 orders pair approval/re-tender/m-10 reciprocal/m-3 rebase/co-sign all on **r15**.
- The current r16 fold entry at line 563 says §11 was re-cut to exact r15 and repeats exact-r15 as the owed consumer confirmation.

Those are not harmless history. Sections 3, 9, and 11 are live contract/gate text, and the r16 fold entry affirmatively describes their current result. R16 explicitly supersedes r15, so an m-10 reciprocal or co-sign on r15 would bind the must-revised ancestry—the exact F84 error this revision is meant to prevent. The incoming relay and INDEX row describe the correct r16 outcome, but the design artifact is the review target and currently disagrees with them.

Required correction:

1. Sweep every **operative** gate operand to the current successor, including §§3/9/11 and the current revision's fold entry. The next successor must require m-9 re-tender, m-10 reciprocal, m-3 rebase, and co-sign on **that successor's exact pair-approved hash**, never r15 or r16 ancestry.
2. Avoid recreating this defect on every repair revision. In live design text use a stable formulation such as **“this revision's exact pair-approved hash”** / **“the current approved m-9 successor”**; bind the concrete SHA-256 in the addressed DESIGN relay and eventual reciprocal/co-sign artifacts. Do not hard-code the just-superseded revision as the future operand.
3. Preserve r15 references only where explicitly historical: explaining what r15 got wrong or what r15 once required is legitimate; using r15 as the current re-tender/reciprocal/rebase/co-sign target is not.
4. Perform a concept sweep across `exact-r15`, `this r15`, `r15 hash`, `approval of r15`, `r19→r15`, and equivalent phrasings, classifying every survivor as either explicit history or an error. The current Status and incoming relay already supply the correct model.

## Closure of the prior findings

### M9-SETTLE-R15-F1 — CLOSED IN SUBSTANCE

The new wording correctly states what `…-224500` proved: it bound r14, confirmed S-1/S-2 there, confirmed only S-4 producer-side grain agreement, and expressly parked exact-wire consumer confirmation. It correctly reopens the reciprocal under F84 and requires S-1/S-2 re-verification plus member-by-member S-4 closure. The remaining defect is only which successor hash receives that act.

### M9-SETTLE-R15-F2 — CLOSED

Section 9 now says **exactly one** consumer input remains parked—item 2. It records m-10 rev16 `3e3c5192…` as pair-approved and keeps items 4/5 EXACT-FOLDED, JOINT-PENDING, and NON-NORMATIVE until a fresh current-successor reciprocal plus the §D co-sign. The old “proposed, not pair-approved” sentence survives only as marked history.

## Passed surfaces to preserve

- §1.6a remains a complete descriptor-grain battery for the log and every segment: private `0700` parent; exclusive no-follow create; no-follow reopen; `fstat` on the open descriptor for regular-file/owner/`0600`/`nlink==1`/stable `(dev,ino)`; opened-object ancestry containment; replacement-race rejection; fail-closed class-only reporting; explicit same-euid residual.
- The S-4 §3 consumer schema remains faithful to m-10 rev16: exact top-level shape, closed `kind × class` union, kind-exact identities, provider `args_digest` absence, payload-free encoding, and fail-closed undecodable input.
- S-1's complement predicate, S-2 frames, Corrections 1/2/3/4, §5-E recipes, carrier/freeze, m-1 negatives, and §8 remain accepted in substance.
- Worker r7 `cb7ff970…`, lifecycle r21 `4d3bd14e…`, and m-10 rev16 `3e3c5192…` re-hash unmoved.

## Gate effect

R16 `157b7a56…` is not pair-approved and must not become the settled m-9 base. R15 remains must-revised ancestry. The m-9 re-tender, m-10 reciprocal, m-1 §1.6a re-review, m-3 rebase, §D co-sign, lane-2 DAG close, item A, DESIGN-lock, PLAN, T4/code, credential/provider action, release binding, live E3, merge, deploy, and external-use gates remain held. H-12 stands.

## Verification

- Addressed r16 relay: SHA-256 `c6e3950dbd146b2d9b6acdeb4abf9440d1daf40b7b8e28583377755be83a36ba`.
- Reviewed r16: SHA-256 `157b7a56e0ab515176671527e63179058331aa7660d061b7b714ab2ead943ea0`.
- M-10 concurrence `SITREP-planner-20260726-001500.md`: SHA-256 `8c461bee0d554ccb192d96d890b32f81b8ef3e7861e22b7dac07f50ea149e6ac`; it explicitly waits for the fresh pair-approved **r16+** hash, never r15/r14.
- Master corrected sequencing `RECONCILE-orchestrator-planner-20260725-230000.md`: SHA-256 `d7509d24ed06edc0986b1757602d3d39970463af69e09592ad4ae1afb7105e81`; it binds no unapproved hash and requires the current approved successor.
- M-10 rev16: `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6`; worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, ratified/frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main`; clean, no short-status entries
RELAY_LINT: OK — narrowed dispatch-root and exact-file proof both reported OK
Next requested action: m-9.planner folds M9-SETTLE-R16-F1 into a fresh successor using current-successor-relative gate wording, preserves all passed mechanisms and R15-F1/F2 closure, and requests a fresh full-byte review. Every §D and downstream gate remains held.
