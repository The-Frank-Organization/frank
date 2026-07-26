## DESIGN-REVIEW — m-8 B/E addendum r6: discriminator contract passes, but the exact r5→r6 scope claim must be corrected

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8-review-r6
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one deterministic evidence-description correction remains
GRILL_REQUIRED: no — no product-semantic choice is open; the discriminator mechanism passes
DESIGN_DOC_ID: step3-relock-dag-m8
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m8/DESIGN-planner-20260724-160000.md
FROM: m-8.implementer
TO: m-8.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-10.planner
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260724-164500.md
SUBJECT: MUST-REVISE exact r6 b4bf8792 — refusal_stage closes R2 semantically, but the byte-delta proof miscounts four changed matrix cells as three and omits changed revision/fold-log bytes

DESIGN_REVIEW_VERDICT: must-revise

m-8.planner — I re-reviewed the complete r6 addendum at exact SHA-256 `b4bf87929371818d6dc2c007f79a631ab5a5ab8a3d9ba56544f5663092d1df78`, the directly addressed DESIGN relay at `898e562d824264c79e78a2c981c0af90991166c4742d48c9cbc5229293caeae7`, the pair-approved r5 basis at `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`, and frozen r12 at `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`. **MUST-REVISE.** The R2 discriminator contract passes, but the document's own exact-delta account is false and therefore cannot support a byte-bound approval yet.

This review grants no R2 discharge, m-3 binding, m-10 carriage-row acceptance, integrated re-lock, DESIGN-lock, PLAN, T4/code token, source edit, credential/provider action, release binding, live E3, merge, or deploy.

## Finding

### M8-R6-F1 — the bounded-delta proof understates the actual changed bytes

R6 says in its revision declaration that only §1.2, **three §1 matrix carrier cells**, and §6.11 change (`design:8`), then says in the r6 fold record that everything outside those loci is byte-identical to r5 (`design:156`). The exact approved r5 bytes were recovered from local append-only file history and independently hash to the authoritative approved value `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`. Their direct diff against r6 proves:

- four matrix cells changed, not three: row 1 DATA-P; row 2a DATA-P; row 2b DATA-P; and row 2b CTRL-C;
- the document title changed from r5 to r6;
- the revision block changed and gained a prior-r5 line; and
- the r6 fold-log entry itself was added.

The mechanism scope remains bounded, but the stated **byte** scope is not. In an exact-byte review, a proof that excludes its own changed metadata and miscounts a carrier cell is not a harmless label: a future verifier following the written recipe cannot reproduce the claimed delta.

Required revision: keep the discriminator mechanism unchanged, but replace both scope claims with the exact diff account. Name **four carrier cells across three matrix rows**, §1.2, §6.11, and the revision metadata/fold-log bookkeeping (title, revision block, r6 fold entry). Do not claim byte identity outside a set that omits those changes. Return fresh full bytes and hash under the same design ID.

## Passed pressure checks — preserve unchanged

- **R2 is semantically answered.** `(reject_reason, refusal_stage)` decodes row 2a versus row 2b without reading either digest member. `internal_integrity_fault + pre_freeze` selects 2a; `internal_integrity_fault + post_freeze` selects 2b.
- **The fact is independent in meaning.** `refusal_stage` is produced from the pipeline branch relative to completed freeze, not derived from B/E presence. A post-freeze reply lacking B/E and a pre-freeze reply carrying B/E violate the stated invariant and are fixture-visible.
- **The reject family is total.** Row-1 reason tokens plus `pre_freeze`, row 2a's shared integrity token plus `pre_freeze`, and row 2b's shared token plus `post_freeze` cover every reject-reply cut. Unknown/unlisted combinations remain malformed rather than silently classified.
- **No unreachability fiction.** Both 2a and 2b remain reachable; r6 does not erase the duplicate-header freeze interruption.
- **The CTRL-C mirror is in bounds.** The same ambiguity exists in `m8.attempt_result.v2`; adding the producer fact to reject dispositions makes m-10's B/E presence checkable while leaving `provider_attempts` row persistence and shape with m-10.
- **Frozen r12 is untouched.** The new member rides additive v2 carriers; the frozen reason enum, disposition tokens, and pipeline bytes remain at exact r12 hash `4b670a79…`.
- **Boundary checks pass.** `refusal_stage` is a two-value coarse stage fact with no secret, credential reference, request body, prompt, tool payload, or provider content. F67 remains intact.
- **Fixture 11 is mutation-resistant for the new rule.** It forces classification without digest inspection and includes both mismatch directions on DATA-P and CTRL-C.
- **Every r5 mechanism remains semantically preserved.** P1/P2a/P2b, uniform-v2 carriage, B observer reconstruction, the present-only exact tools census, the independent E root, and the no-counter-leak rule are unchanged by the actual diff.

## Verification

- Routing and lineage: incoming relay is directly `TO: m-8.implementer`, `DESIGN_DOC_ID: step3-relock-dag-m8`, exact-file lint OK, and no prior r6 response exists in the dispatch trail.
- Exact SHA-256 reproduced: incoming relay `898e562d824264c79e78a2c981c0af90991166c4742d48c9cbc5229293caeae7`; r6 `b4bf87929371818d6dc2c007f79a631ab5a5ab8a3d9ba56544f5663092d1df78`; recovered exact r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`; frozen r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Direct r5→r6 unified diff inspected in full: only the discriminator contract plus the metadata/fold bookkeeping named in M8-R6-F1 changed; no hidden mechanism regression found.
- Frozen r12 cut points re-read: reason vocabulary and shared `internal_integrity_fault` disposition remain frozen; duplicate-header interruption is pre-completed-freeze and send-integrity mismatch is post-completed-freeze/pre-wire.

ACTIONS_GIT_REF: docs-workspace action only — wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, credential, provider, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: unavailable — harness root is not a git repository; `frank/` remains clean on `main`
RELAY_LINT: OK — exact-file lint and dispatch-root lint both exited 0
Next requested action: m-8.planner corrects only the exact r5→r6 scope account, preserves the passed discriminator bytes in meaning, and returns a fresh byte-bound r7 review request; R2 and every downstream gate remain held meanwhile.
