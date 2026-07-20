## DESIGN-REVIEW - s8 executable-claim declaration semantics r3 approved; master and m-7 capability reconciliation remains an external gate

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m3-review-r3
PARENT_DISPATCH_ID: s8-claim-input-m3
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review approved; master parent correction and m-7 capability-owner return remain required before v7 finalization or T9 execution
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s8-claim-input-m3/DESIGN-planner-20260712-015500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, s8.planner, s8.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve r3 - declaration semantics are deterministic and fail-closed; m-7 now explicitly owns the v7 capability transition and machinery-fault remains distinct from no-vantage

DESIGN_REVIEW_VERDICT: approve

R3 closes F3 and F4 without reopening c2 or broadening the declaration surface. The m-3 semantics leg is approved for return to master.

## Findings Closed

### F3 - CLOSED: v7 compatibility is now an explicit m-7 owner-relation change

Section 12 no longer claims the capability set is untouched or that seam (v) alone supplies the closed-compatibility proof. It now binds:
- a governed m-2 `v6 -> v7` fieldspec transition;
- an m-7-owned reader capability-table transition adding `s8-fieldspec-v7`;
- marker-first phase-0 refusal by a v6-only reader before partial interpretation;
- reader-first, forward-only sequencing with rollback/skip rejection;
- a v7-capable reader that starts, reads the present declaration, and executes it;
- the stale-v6-form `re-render` leg.

The tripwire is now correctly expressed as no lock-pinned value move **by proxy**: the capability set may move only through m-7's owner relation, while the genesis v5 SHA, FX-CFG-7 hash, and catalog pin remain untouched. The cross-owner handshake explicitly expands LEG m-7 beyond composition seam (v) (`...probe-design.md:231-235,250,254`).

### F4 - CLOSED: machinery fault is keyed by its explicit conductor signal

The aggregation matrix now distinguishes:
- observed false -> `rejected` for both authority classes;
- explicit conductor-classified machinery fault -> authority `held`, non-authority `rejected`/author-return;
- `skipped`/`unsafe` and blocked/degraded **without** the machinery-fault signal -> the record-integrity/no-vantage path;
- all pass -> `accepted` with maximum passing rung.

This matches the locked `MachineryFault`-before-`Blocked/Degraded` ordering and avoids the timeout-versus-no-vantage conflation. The fixture matrix now includes pass+unsafe, pass+genuine-fault, false+genuine-fault, and pass+genuine-no-vantage-blocked (`...probe-design.md:239-248`).

## Full Amendment Check

- Shape/cardinality: `(claim_ref, check_id, params)`, many-per-record.
- Identity: `claim_ref` non-empty, bounded, symbolic, unique; duplicate typed reject and observe-time refusal.
- Validation: registry-schema-aware fill-time validation plus authoritative pre-spawn normalization/refusal.
- R2: input declaration non-gate-referenceable; conductor-observed output drives disposition.
- Suppliability: seat-owned `executable_claims` input remains distinct from system-owned `executable_claim_results` output.
- Aggregation: canonical row order, run-all, one result row per declaration, deterministic precedence.
- Rail A: optional absence is open/no-vantage; present declaration is closed/fail-closed on compatibility.
- Rail B/I-PH: honest-agent utility retained; rendered/bounce/verdict surfaces remain symbolic, bounded, path-free, and effective-config-value-free.
- Scope: no s9 adjudication content, no new terminal, byte-exact `{accepted, rejected, held}`, interim defaults and s10 sunset unchanged.

## Remaining External Gates

This pair approval does not itself finalize `s8-fieldspec-v7`, modify the m-7 capability table, authorize implementation, or lift T9. Before those actions:
1. master must ratify the corrected Rail-A parent posture and expand LEG m-7;
2. m-7 must confirm its exact supported set, forward relation, implementation locus, and both phase-0 proof legs;
3. m-2 must finalize its reviewed bytes against this approved m-3 semantic return;
4. master must reconcile the three owner legs and issue the bounded T9 fold/grant.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no frank source/test edit, branch, commit, push, PR, or merge
FINAL_GIT_STATUS_SHORT: `git -C frank status --short` returned empty at `main@691d034`; active `s8-observe-spine@3cce8cd` contains existing build-lane changes not made by m-3.implementer
Next requested action: m-3.planner returns the approved §12 semantics to master; master reconciles the parent correction and expanded m-7 capability leg before m-2 v7 finalization and T9
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-claim-input-m3/DESIGN-REVIEW-implementer-20260712-016200.md`
