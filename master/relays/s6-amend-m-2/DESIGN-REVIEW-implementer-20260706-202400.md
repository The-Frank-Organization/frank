## DESIGN-REVIEW - m-2.implementer review of B-2 typed BOOT form addendum

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-amend-m-2-b2-review-r1
PARENT_DISPATCH_ID: s6-amend-m-2-b2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review only; master/VP/operator gates remain upstream
DESIGN_DOC_ID: s6-amend-m-2-transport-codec
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s6-amend-m-2/DESIGN-planner-20260706-201500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: approve - B-2 typed BOOT form is bounded; SITREP plus lifecycle-gated form composes with B-1/B-3

VERDICT: approve. I reviewed the B-2 addendum in §11 of `master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md` against the addendum dispatch and the current B-1/B-3 seam texts. The design is bounded: it formalizes the boot form and token decision without adding a new phase, record kind, gate category, seat verb, or authority gate.

## Basis

- Incoming B-2 design relay: `master/relays/s6-amend-m-2/DESIGN-planner-20260706-201500.md`.
- B-2 amendment text: `master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md:141-169`.
- Addendum dispatch: `master/relays/s6-design/DESIGN-orchestrator-planner-20260706-200259.md:21-44`.
- B-1 seam text: `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:88-124`.
- B-3 seam text: `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:91-103`.
- Locked parent premise: `master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:19-25`.
- Live FieldSpec anchors: `frank/internal/fieldspec/registry.json:88-97,125,134`; `frank/internal/fieldspec/render.go:55-87`; `frank/internal/fieldspec/validate.go:54-61`.

## Review checks

1. **Token ruling: `PHASE: SITREP` is the correct low-cost choice.**

   A new `BOOT` phase would be shared vocabulary, with follow-on work in phase/authority consistency, phase-scoped routing, and every phase conditional. B-2 correctly rejects that for a record whose behavior is report-only acknowledgment. It also correctly avoids a new `record_kind`: the current registry already treats `record_kind` as a seat-scoped authorization surface, and s5 locked `genesis` out of every public seat scope; using record_kind for boot would reopen that major surface. No `gate_category` or §J2 impact is needed because boot grants nothing and resolves no human gate.

2. **Identity is the channel stamp, not a seat-filled claim.**

   B-2's required set omits any seat-filled identity field, which is correct. `FROM`/`ROLE` are system-owned and system-only in the live FieldSpec, and B-3 defines activation as liveness bookkeeping over the already channel-stamped identity, not an identity-strength upgrade. The boot form may report that the seat loaded its charter and read/awaits dispatch; it must not let a seat assert who it is.

3. **Lifecycle-gated form selection is not a new §5 predicate atom.**

   The design routes "pre-active" form choice through B-1's lifecycle render mirror: a minted-but-not-active seat is shown only the B-2 boot form, while a hand-crafted non-boot submit gets a typed `boot-required` refusal. That is a conductor render-context branch owned by m-7, not a field-level `required_when` predicate inside the bounded §5 grammar. B-2 keeps the two boot fields required within the boot form class and does not introduce a self-marker gate.

   Carry-forward: if master integration or B-1 chooses to materialize a `boot_ack` marker for recovery/classification convenience, it must be system-derived, non-gate-referenceable, and not seat-filled. Otherwise the m-2-approved branch is simply SITREP plus the two boot fields as the form class.

4. **`charter_loaded` as self-reported is honest for Step-1.**

   B-2 marks `charter_loaded` as non-gate and self-reported. That is the right Step-1 posture: there is no observe layer to verify charter loading, and the boot record is not an authority record. The fixture correctly requires that no predicate keys on `charter_loaded`.

5. **The un-bounceable claim is scoped correctly.**

   The claim is not generic bounce-proofing. It is limited to a pre-active seat filling exactly the rendered boot form after A-1 stable digest and branch-A conductor-computed PARENT. Under that scope, the live failure classes named by the addendum - digest churn and parent staleness - are structurally absent. Any future bounce would be a real schema/config mismatch or hand-crafted non-boot pre-active submit, not the observed boot retry loop.

## Held gates and boundaries

This approval covers only B-2, the m-2 typed BOOT form addendum under `DESIGN_DOC_ID: s6-amend-m-2-transport-codec`. It does not approve B-1 lifecycle/roster mechanics, B-3 identity activation semantics, the integrated amendment package, the VP co-sign, a c1 fold, or any code work. It also does not authorize a new phase, record kind, gate category, §5 predicate atom, or seat verb.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended the matching `master/relays/INDEX.md` row; reviewed docs/code only; no `frank/` code changes; no c1 fold; no PLAN or IMPL.
FINAL_GIT_STATUS_SHORT: docs workspace root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git` / exit 128; `git -C frank status --short` empty / exit 0.
RELAY_LINT: OK - relay-lint.py exit 0 on `master/relays/s6-amend-m-2/DESIGN-REVIEW-implementer-20260706-202400.md`
