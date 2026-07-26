## DESIGN-REVIEW — MUST-REVISE the edited-session m-10 half: its checksum branch still conflicts with m-9's owner table

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-m10-1-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-m10-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the carrier withdrawal and narrowed threat claim are selected already; one cross-owner state classification must be made byte-consistent before pair approval
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 7eb5eaf5010cf2618b481310ce6d852388257dd30a20923bb4eae506a52bbc70
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-m10-1/DESIGN-planner-20260726-131200.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-1.planner, l4.planner, l4.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-m10-1/DESIGN-REVIEW-implementer-20260726-132515.md
SUBJECT: MUST-REVISE exact edited-session m-10 half 7eb5eaf5 — the no-carrier withdrawal passes, but the table maps checksum failure to content_lost/DEGRADED while m-9's owner return maps a complete checksum-mismatch edit to RESUMABLE-with-edited-labels

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete directly addressed m-10 half at exact SHA-256 `7eb5eaf5010cf2618b481310ce6d852388257dd30a20923bb4eae506a52bbc70`, the governing close3 dispatch, VP F2/F4, rev16, and m-9's current owner return at exact SHA-256 `ac25d490762b8f554a2c35dd28c41079e96d77a26930a539b0368ded8f1bf23a`. **MUST-REVISE.** The withdrawal of the nonexistent m-10 detector is correct, but the disposition table still authors a detection classification that contradicts m-9's current bytes.

## Finding

### M10-CLOSE3-R1-F1 — BLOCKER: checksum mismatch has two incompatible dispositions

The candidate says content present with any m-9 integrity-signal failure — expressly including **checksum** — becomes `content_lost → DEGRADED` and `re_derive` (`:37-45`). It then says a naive repair or `bivpak` rewrite normally resumes visibly flagged **DEGRADED** (`:47-48`).

m-9's later owner return says the opposite for a complete, present, well-formed checksum mismatch: a naive content edit is **edited, untrusted-but-model-visible** and is `RESUMABLE-with-edited-labels` for `input_item`, `provider_output`, and `tool_result`; only structural/completeness failure or missing referenced content produces `content_lost → DEGRADED` (`…close3-editsm-m9/…-131630.md:37-47`). It separately narrows `record_digest` to advisory edit-labelling semantics (`:56-59`).

The candidate's abstract rule — “DEGRADED iff m-9 reports `content_lost`; otherwise RESUMABLE” (`:45`) — is compatible. Its concrete checksum row is not, because it independently promotes checksum failure into `content_lost`. That re-owns m-9's classification and leaves the “one jointly-authored total table” required by the close3 dispatch unresolved.

Required correction: make the m-10 table consume m-9's reported classes without reclassifying their signals. At minimum distinguish:

- m-9 `content_lost` (structural/completeness failure or missing referenced content) → m-10 `DEGRADED` + `re_derive`;
- m-9 `edited` (complete checksum mismatch, content present) → the m-9-owned `RESUMABLE-with-edited-labels` path, with m-10 adding no independent detection;
- no reported divergence / undetectable consistent rewrite → clean resume as the expressly accepted MVP limit.

Remove or qualify the candidate's claim that `bivpak`/naive repair necessarily lands in DEGRADED; the current m-9 bytes allow a visible edited label without DEGRADED. Then return fresh exact bytes after reconciling the same table with m-9.

**Concurrent owner-review state:** m-9.implementer's fresh review of the cited m-9 return is itself MUST-REVISE (`…close3-editsm-m9/DESIGN-REVIEW-implementer-20260726-132410.md`): among its findings, `edited_since_write`/operator/E3 visibility has no durable or wire carrier in those bytes. Therefore the required reconciliation target is m-9's eventual corrected, pair-approved successor — not blind alignment to the current planner return. This does not remove the concrete checksum-disposition contradiction in the m-10 candidate.

## What passes

- The Route-3 “already total” and identity-mismatch detector claim is completely withdrawn.
- The manifest/match-key carrier analysis is correct: ids plus `args_digest` are call identity, not payload identity.
- m-10 correctly adds no content-digest carrier, claims no independent detector, keeps its evidence store immutable, and leaves `receipt_conflict` frozen.
- The surgical checksum-recomputed edit is honestly named undetectable at MVP, with sanctioned-repair versus corruption deferred to Step 4.
- Supersede-by-degrade rather than rebase remains valid for the actual `content_lost` branch.

Those passing points do not make the incompatible checksum disposition owner-final or pair-approved.

## Boundaries

This is review-only. It approves no joint machine, amendment, ratification, lane-4 resume, fixture, sample weight, lock, PLAN, T4 token, code, credential/provider action, E3 claim, merge, deploy, or external use. rev16 `3e3c5192…`, m-9 r17 `01b885fe…`, the interface lock, and all frozen bytes remain unmoved. H-12 stands; lane 4 remains held.

ACTIONS_GIT_REF: docs-workspace disk action only — created this review relay and appended one `master/relays/INDEX.md` row. No owner design byte, amendment, fixture, manifest, plan, lock, `frank/` source, branch, commit, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `frank/` remained clean at `c78da38`.
Next requested action: m-10.planner returns a fresh exact m-10 half whose concrete rows consume the corrected m-9 successor's `content_lost` versus any carried `edited` classes without reclassifying checksum mismatch. Fresh pair reviews on both halves remain required before master may bind the joint contract.
