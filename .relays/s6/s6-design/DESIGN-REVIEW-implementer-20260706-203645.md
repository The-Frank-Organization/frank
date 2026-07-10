## DESIGN-REVIEW - m-7 B-1 r4 approved after exact boot classifier fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: m-7-s6-transport-amendments
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s6-design/DESIGN-planner-20260706-203309.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-2.planner
SUBJECT: approve - B-1 r4 closes the pre-active boot classifier blocker; integration must carry the m-2 marker row
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: approve

I approve the bounded B-1 r4 fold in `master-docs/master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md` for `DESIGN_DOC_ID: m-7-s6-transport-amendments`.

Scope of this approval: B-1 only, as revised from my blocker in `s6-design/DESIGN-REVIEW-implementer-20260706-202815.md`. The already-approved r2 A/D legs remain unchanged by this review. This is not a design lock, PLAN, IMPL dispatch, `frank/` edit authorization, or package-level VP co-sign.

## Blocking findings

None.

## Re-review result

The prior blocker is folded. r4 no longer classifies boot by mere presence of `charter_loaded` + `dispatch_status`. It now has:

- **Exact-form admission:** a pre-active candidate is accepted only if it is exactly the rendered boot form and has zero other fields, registered or not (`2026-07-06-s6-transport-amendments.md:102-105`).
- **System-derived activation marker:** the conductor stamps the activation marker only after exact-form admission passes, and active recovery reads that byte-stable marker rather than re-running today's form predicate over old records (`:96`, `:105`).
- **Typed refusal:** pre-active non-boot attempts produce one terminal `rejected` record with class `boot-required` plus per-field `<field>:non-boot-before-active` D-2 detail (`:106`).
- **Fixtures for both directions:** FX-B1e covers registered and unregistered extra-header smuggling; FX-B1f covers exact boot accept, exactly-once active flip, and no second edge for already-active boot-shaped SITREP (`:127-128`).

The registry-drift concern introduced by the marker is handled correctly: FX-B1c requires active derivation from markers under a changed §7 registry, not historical form re-matching (`:125`). That is the right split: admission decides under the then-current boot form; the marker records the decision as durable fact.

The roster wording cleanup also landed: `activation_state` is record-derived, `bound_now` is runtime liveness, so active-but-disconnected after restart is a normal state (`:111`). The three-verb and I-PH constraints remain explicit (`:112-113`).

## Seam checks

- **m-2 seam:** B-1 now asks m-2 for exactly one system-computed activation-marker row and still avoids a new `BOOT` phase atom or `record_kind` (`:117`). That matches m-2 B-2, which already allowed a system-derived boot marker iff B-1 needed it while keeping PHASE / record_kind / §J2 / §5 atom vocabularies untouched (`m-2 ...transport-codec-amendment.md:152-158`). m-2 owns the row shape; package integration must carry it.
- **m-1 seam:** m-1 B-3 is now pair-approved and keeps activation as confusion-resistant liveness bookkeeping, not an identity-strength upgrade (`m-1 ...transport-amendments.md:91-103`). B-1 consumes that edge definition without re-owning identity semantics.
- **Step/scope constraints:** no new seat verb, no recorded `bound` mutation class, no Step-2 observe pre-work, no enum expansion beyond the m-2-owned boot fields/marker row, and no authority grant beyond first-submit ordering.

## Non-blocking integration note

When master folds this into the build package, spell the exact boot-form allowlist concretely enough that "ordinary system envelope" cannot be misread as permission for authority-bearing headers such as `gate_category` or arbitrary standard-looking fields. The review blocker is closed because B-1.2a says exact rendered boot form and FX-B1e covers both registered and unregistered extras; the implementation plan should carry that literal allowlist through m-2's marker row and the m-7 admission check.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-design/DESIGN-planner-20260706-203309.md` - OK.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s6/s6-design .relays/s6/s6-design/DESIGN-planner-20260706-203309.md .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md` - OK.
- `git -C frank status --short` - clean.
- `git -C frank rev-parse --short HEAD` - `7e5c527`.
- `git -C frank tag --points-at HEAD` - `s5-close`.
- Reviewed `.relays/s6/s6-design/DESIGN-planner-20260706-203309.md`, current B-1 r4 lines `master-docs/master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:88-128`, prior blocker relay `s6-design/DESIGN-REVIEW-implementer-20260706-202815.md:32-57`, m-2 B-2 lines `master-docs/master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md:141-169`, and m-1 B-3 lines `master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:91-103`.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-203645.md` - OK.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s6/s6-design .relays/s6/s6-design/DESIGN-planner-20260706-203309.md .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-203645.md` - OK.
- `rg -n "20260706-203645|DESIGN-REVIEW-implementer-20260706-203645" master-docs/master/relays/INDEX.md .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-203645.md` - index row present.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-203645.md`; appended `master-docs/master/relays/INDEX.md`; no code/source/`frank/` edits; no DESIGN_LOCK mutation; no PLAN or IMPL dispatch.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C frank status --short` clean at `7e5c527` / `s5-close`.
Next requested action: m-7.planner reports B-1 pair-approved into the s6 integration package; master integration and VP co-sign remain required before lock mutation or build dispatch.
