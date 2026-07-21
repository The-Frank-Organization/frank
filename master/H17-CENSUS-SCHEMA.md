# H-17 — the TCB / effect census: canonical row schema (v1)

**Status:** issued 2026-07-20 by master per VP F87 (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-045522.md`); the ONE schema every census row uses. Byte-exact: cite this file by SHA-256; a schema change is a v2 with its own hash, never an edit.

Rule of authority: a row STATES current designs and residuals — it never authorizes moving a closed byte. Where a frozen contract lacks an answer, the cell says exactly `unknown`, `not specified`, or `residual`; master owns reconciliation at stage 6.

Rule of coverage: each design maps EVERY authoritative transition/effect family to a row OR an explicit non-effect rationale. The dispatch parentheticals were minimum families, not the inventory.

## Row fields (all required; use the null tokens above rather than omitting)

| field | meaning |
|---|---|
| `effect_id` | stable kebab-case merge key, globally unique across the census (e.g. `m9-f59-local-tool-invoke`) |
| `effect_class` | e.g. `reversible_workspace_mutation`, `external_send`, `durable_state_commit`, `credential_operation`, `process_lifecycle` |
| `requester` | who asks for the effect (seat/component) |
| `executor` | who performs it |
| `authority_source` | what grants it (policy artifact/decision, ratified constant, operator grant) |
| `policy_owner` / `policy_artifact` | the owning domain + the exact document/digest the policy lives in |
| `decision_point` | where allow/deny is decided |
| `enforcement_point` | where the decision is enforced (must control the effect) |
| `exclusive_credential_holder` | who alone holds the credential/capability; `ambient` is a legal honest value |
| `request_freeze_point` | where the request becomes immutable |
| `authorization_linearization_point` | where authorization becomes fact |
| `effect_linearization_point` | where the effect becomes fact |
| `outcome_reporter` | who reports the outcome (may be the executor — self-report) |
| `outcome_observer` | who independently observes it; `none (self-reported)` is a legal honest value |
| `outcome_validator` | who validates the report before it persists |
| `canonical_record` | the durable row/record family where the outcome lives |
| `bypass_paths` | ALTERNATE EFFECT PATHS around the enforcement point (bash-ambient, same-UID, direct API); a failure-to-record is NOT a bypass — it belongs in failure semantics |
| `failure_unknown_semantics` | what crashes/losses produce (typed terminals, parked UNKNOWN, silent drop — say which) |
| `replay_idempotency` | one-shot / idempotent / replay-first / none |
| `threat_claim_scope` | the claim boundary (e.g. `confusion-not-malice, same-UID residual accepted`) |

## Distinctions the VP requires (F87)

- **Reporter ≠ observer ≠ validator ≠ recorder** — name each even when one component fills several roles; recording or validating a report does not make a component an independent observer of the effect.
- **Admission ≠ authorization** — e.g. `attempt_open_ok` is admission/row-first ordering; provider policy authorization is m-8's, immediately before credential attach/send.
- **Non-append verbs are their own rows** — `relay.project`/`read`/`Describe`/push are read/serve effects, not the store-append effect.
- **A bypass is an alternate way to CAUSE the effect**; E0 non-emission is a failure/unknown condition, not a bypass.
