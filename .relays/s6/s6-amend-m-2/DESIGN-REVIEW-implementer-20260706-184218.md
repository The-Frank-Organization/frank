## DESIGN-REVIEW - m-2.implementer review of s6 transport-codec amendment

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-amend-m-2-review-r1
PARENT_DISPATCH_ID: s6-amend-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review only; master/VP/operator gates remain upstream
DESIGN_DOC_ID: s6-amend-m-2-transport-codec
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: .relays/s6/s6-amend-m-2/DESIGN-planner-20260706-184500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: must-revise - codec direction is right, but record_kind scope and recipient projection must be tightened before lock

**Verdict: `must-revise`.** The amendment's central move is correct: there should be one canonical `address_list` decoder used by validation, lineage, render/fill, and delivery. That kills the F6/F7 class instead of patching one call site. I cannot approve the amendment as written because two acceptance/visibility rules still overrun the locked s5 contract, and the F12 fallback branch is not yet implementable in the current FieldSpec shape.

**Blocking findings.**

1. **F13 currently conflates named enum membership with submit authorization.**

   Amendment §2.6 says `validateRecordKind` should accept "every registry-advertised `record_kind`" and §7.3 requires acceptance for every `record_kind` token the registry offers, including `genesis`. That is too broad against the s5 lock. The live registry intentionally keeps `genesis` in the named enum for compatibility (`internal/fieldspec/registry.json:84`) while removing it from every `seat_scope` (`internal/fieldspec/registry.json:125`). The s5 reconciliations also locked that distinction: `genesis` is store-init machinery, not a public rendered/submittable record kind (`.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md:25-26`; `.relays/s5/s5-b-merge-gate/RECONCILE-orchestrator-planner-20260706-153721.md:26`).

   Required revision: split the contract into three layers:
   - named enum membership: the token is known to the registry;
   - seat_scope/form offer: the token is authorized for this submitting seat;
   - per-kind required-field validation: extra checks for authorized kinds.

   `validateRecordKind` may stop being a second independent membership judge, but only after `reg.Validate` has enforced enum membership and seat scope (`internal/engine/submit.go:47-64`; `internal/fieldspec/validate.go:54-61`). The fixture plan must include a negative leg proving public `record_kind=genesis` is rejected by seat scope, while authorized operator-offered kinds such as `gate_resolution`, `disposition`, `diagnostics`, `config_change`, and owed kinds remain accepted subject to their required fields. This also fixes the `gate_category: other` side of F13 without widening `record_kind`.

2. **`Envelope.To` as "primary recipient projection" is not enough for derived relay visibility.**

   Amendment §2.4 says `Envelope.To` becomes a display/index projection and delivery/lineage read the decoded Header list. In the current code, rendered markdown and `INDEX.md` rows still print `rec.Envelope.To` directly (`internal/store/projections.go:110-113`), while mailbox intents use `DeliveryRecipients` (`internal/store/projections.go:116-141`). If `Envelope.To` becomes only the primary recipient, a multi-TO/multi-CC relay can be delivered to all intended mailboxes but rendered/indexed with incomplete recipient visibility. That recreates the same "one path sees less than another path" failure mode the amendment is meant to remove.

   Required revision: make canonical decoded Header `TO`/`CC` the source for rendered relay headers, index display, lineage, and delivery. `Envelope.To` can remain a primary-recipient compatibility/routing projection if m-7 needs it, but it must not be the only display/index representation of recipient truth. Add a fixture that submits multi-TO plus multi-CC canonical `address_list` values and proves the projected relay markdown, index row, mailbox intents, and reviewer-visibility gate all preserve the same full recipient set.

3. **F12's header fallback is under-specified for an operator-scoped arbitrary rationale.**

   The amendment correctly rejects the current `ORCH_REVIEW_WAIVER: *` presence-only shape, but the fallback "header carries free_text rationale" branch is not yet a complete m-2 schema contract. The current FieldSpec model has a seat-scoped enum path for constrained values (`internal/fieldspec/registry.json:134`), and generic field type/options metadata (`internal/fieldspec/fieldspec.go:15-20`); it does not by itself express "operator-scoped arbitrary free_text header value" with submit-time enforcement. m-1's F17 direction is an operator-only waiver record class with scope/retraction semantics (`master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:43-51`).

   Required revision: either bind F12 to the m-1 waiver-record carrier and have m-2 define the exact rationale/scope/retraction rows for that record, or specify the exact header fallback validator that enforces operator-only arbitrary rationale text without weakening seat-scope semantics. As written, the header fallback names the desired type but not the enforcing mechanism.

**Non-blocking confirmations.**

- The F6/F7 codec thesis is directionally sound. The live divergence is real: validation already requires canonical `address_list`, lineage still raw-splits comma-delimited headers, and delivery silently falls back to `Envelope.To` on canonical parse failure (`internal/fieldspec/canonical.go:19-41`; `internal/lineage/lineage.go:282-292,432-439`; `internal/store/projections.go:122-141`).
- The parent seam is directionally composable. If m-1 locks conductor-computed `PARENT`, current render behavior already hides `owner: system` / `system_only` fields (`internal/fieldspec/render.go:55-63`). If the widened-candidate branch locks instead, m-2 still needs the same source-of-truth rule: the render candidate list and lineage acceptance list must be one list, not parallel lists.
- The amendment does not need to reopen c1 if revised under this DESIGN_DOC_ID and folded only through the integrated VP co-sign path.

**Revision request.** Please send a revised DESIGN relay for `s6-amend-m-2-transport-codec` that tightens §2.4, §2.6, §4, and §7 fixtures as above. I will re-review the revised amendment in DESIGN-REVIEW scope only.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended the matching `master-docs/master/relays/INDEX.md` row; reviewed docs/code only; no `frank/` code changes; no c1 fold; no PLAN or IMPL.
FINAL_GIT_STATUS_SHORT: docs workspace root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git` / exit 128; `git -C frank status --short` empty / exit 0.
RELAY_LINT: OK - relay-lint.py exit 0 on `.relays/s6/s6-amend-m-2/DESIGN-REVIEW-implementer-20260706-184218.md`
