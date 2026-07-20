## DESIGN-REVIEW - supply-set m-7 r2 must correct schema-ref mismatch disposition

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-supply-set-m7-review-r2
PARENT_DISPATCH_ID: s8-supply-set-m7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one bounded semantic correction; no operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-supply-set-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-supply-set-m7/DESIGN-planner-20260712-005936.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-7.planner, m-3.planner, m-2.planner, s8.planner, s8.implementer
SUBJECT: F1/F2 mechanics closed, but FX-SUP-6 wrongly turns a known schema-ref content mismatch from observed false into a typed input refusal

DESIGN_REVIEW_VERDICT: must-revise

r2 closes both prior findings at the requested grain: schema-ref values are canonical SHA-256 digests; roots/commands resolve and validate at composition; static args are composition-checked against the executor I-PH contract; and the dogfood descriptor remains exact. One result-disposition error was introduced in the positive fixture wording.

## Finding

### F3 - A governed schema digest mismatch is observed false, not a typed refusal

FX-SUP-6 says the E1 `schema_ref` check "passes on match / typed-refuses on mismatch." The shipped reader distinguishes two cases:

- unknown/absent schema-ref id is an invalid selection and should be the new typed `schema-ref-unknown` refusal;
- a known id resolves to its governed expected digest, the file is read and hashed, and unequal bytes produce the ordinary observed-false `read-file-mismatch` verdict (`internal/observe/checks_base.go:17-48`).

That distinction is m-3 evidence semantics, not m-7 supply policy. Converting a real content mismatch into refusal would erase a valid negative observation and change downstream predicate/disposition behavior.

Required fold: rewrite FX-SUP-6 as three legs: known id + matching digest -> E1 pass; known id + mismatching file digest -> observed `fail` / `read-file-mismatch`; absent id -> typed `schema-ref-unknown` refusal before file comparison. Keep malformed configured ids/digests as startup composition faults. Sweep any "mismatch = refusal" echo; no mechanism or token beyond the already-proposed unknown-id refusal changes.

## Confirmed

- Canonical lowercase 64-hex SHA-256 is the correct `schema_refs` value grammar.
- Root/command canonical resolution and startup validation close F1.
- The composition-time static-argument policy and planted-token fixture close F2.
- Engine v2 home/ceiling, v1->2 transition, dogfood descriptor, timeout behavior, governed RegistryEnv supply, and I-PH posture remain accepted.
- m-3's governed-root/enum-face semantics and m-2's no-FieldSpec-impact confirmation remain compatible.

Gate disposition: the primary leg remains uncountersigned only for F3. No other finding remains open.

Not authorized / not done: no design/code edit, no T9 lift, no m-2/m-3 proxy content, no merge, and no effective amendment fold.

ACTIONS_GIT_REF: wrote this review relay and appended one `master/relays/INDEX.md` row; read-only inspection of the s8 worktree at `3cce8cd`; no `frank/` or s8-worktree edit by this seat
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main`; s8 worktree at `3cce8cd` retains the pair's in-flight T9/T10 changes, untouched by this seat
Next requested action: m-7.planner corrects FX-SUP-6/disposition wording and returns r3; the implementer re-reviews directly to master. The three-leg fold and T9 re-lift remain held meanwhile.
