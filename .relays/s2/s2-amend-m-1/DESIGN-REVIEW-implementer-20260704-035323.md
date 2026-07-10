## DESIGN-REVIEW - s2-amend-m-1 conductor-internal provenance amendment

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s2-amend-m-1
PARENT_DISPATCH_ID: s2-amend-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: DESIGN-planner-20260704-035030.md
SUBJECT: approve the narrow m-1 conductor-internal provenance amendment

Findings: none.

Review scope was the narrow amendment requested in `.relays/s2/s2-amend-m-1/DESIGN-planner-20260704-035030.md`: whether the new m-1 design-of-record text faithfully records the `system`/`system` conductor-internal provenance convention from the S2 fidelity verdict, stays inside m-1's stamping lane, and preserves the byte-exact outcome enum.

Verdict: approve.

Evidence:
- The fold-log entry records the basis and scope at `master-docs/master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md:51-53`, including the S2 fidelity source, non-`submit` authorship path, m-1 stamping-only boundary, D4/D5 framing, and no c1 reopen.
- The §6 bullet captures the required convention at `master-docs/master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md:150`: `FROM = "system"`, `ROLE = "system"`, non-lane/non-seat reserved provenance, no public-submit acceptance, `schema_version` in the envelope/system-only home, and `DeliveryState` constrained to `{accepted, rejected, held}`.
- This matches the S2 fidelity requirement at `.relays/s2/s2-fidelity-m1/SITREP-implementer-20260704-034158.md:48-58` and the genesis envelope detail at `:60-72`.
- The amended §6 bullet explicitly keeps the internal-record catalog and on-disk shapes with m-7/S2, and keeps `record_kind` as an m-2 header slot. That preserves the lane boundary and does not claim the S2 item catalog as m-1 design-of-record.
- Preservation probes passed: `grep -oE '\{accepted, rejected, held\}' ... | wc -l` = 5; `grep -c 'observer-selected' ...` = 9; `grep -n 'unspoofable' ...` shows only the existing §0.d fold-log meta-reference at line 49.
- Incoming planner relay lint passed with `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-amend-m-1/DESIGN-planner-20260704-035030.md`.

Boundary note: this approval covers only the m-1 design-of-record hygiene amendment. It does not replace the outstanding S2 core fold/re-review gate from the S2 fidelity relay; S2 still needs to fold F-M1-1/F-M1-2/F-M1-3 into its own design/plan before any implementation dispatch.

ACTIONS_GIT_REF: wrote `.relays/s2/s2-amend-m-1/DESIGN-REVIEW-implementer-20260704-035323.md` and appended `master-docs/master/relays/INDEX.md`; no `frank/` code edits; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace); `git -C frank status --short` is clean.
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-amend-m-1/DESIGN-REVIEW-implementer-20260704-035323.md`
