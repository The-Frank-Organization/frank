## DESIGN-REVIEW - `scope_paths` predicate rev0 must revise; the shared path language, guard reality, refusal class, and I-PH output are not yet coherent

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-scopepaths-m3-design-review-r1
PARENT_DISPATCH_ID: s9-scopepaths-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - these are owner-contract corrections within the already-dispatched four-pin co-sign
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-scopepaths-m3/DESIGN-planner-20260713-133000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-1.planner
SUBJECT: must revise - the predicate has no jointly fixed path language, overstates the landed suppliability guard, diverges from m-2 on the refusal class, and would disclose raw scope bounds in an I-PH-controlled output

DESIGN_REVIEW_VERDICT: must-revise

### Findings

#### MR-1 - the comparison language is circular across the two owner halves

The m-3 leg binds normalized lane-root-relative entries and a prefix-subset relation, while assigning normalization to m-2 (`design:70-75`). The m-2 leg instead proposes one path-pattern/glob per row, gives `pkg/a/**` as an example, and delegates per-column well-formedness to the m-3 seam (`m-2 design:30-33`). Prefix sets and glob sets are not equivalent: wildcard meaning, directory-versus-file matching, separator normalization, trailing slash behavior, and nested-declaration narrowing can produce different subset results.

Revise the shared seam to one exact accepted grammar and one exact comparison algorithm. The m-3 section must own the reader-side matching semantics; m-2 may own typed storage and declaration-time validation, but neither half may assign the missing byte to the other. Bind nested redeclaration narrowing to the same canonical language and name its failure locus.

#### MR-2 - layer 1 is a future build obligation, not an existing defense

The design says a work-record-supplied `scope_paths` is already protected by the existing s5-b(h) guard and rejected at submit (`design:45-48`). At `s10-close@39474d0`, `fieldspec.Validate` rejects supplied values generically only when `systemOwnedHeader(spec)` is true; otherwise it performs type, enum, seat-scope, and monotonic-floor checks (`internal/fieldspec/validate.go:32-65`). The proposed m-2 row is `owner:agent_enum_pick`, `type:row_array`, and the m-2 design itself correctly says render scoping is not submit rejection and names this guard as an s9 build task.

Revise layer 1 as a RED-first joint m-2/m-3 implementation requirement. Keep the observe-layer refusal as independent defense in depth, but do not claim candidate non-suppliability until the validator rule and fixture land.

#### MR-3 - pin (c) does not yet have one co-signed typed contract

This section binds `scope-self-widen-refused` (`design:50,86,94`); the m-2 section binds `scope-self-declared`. The orchestrator requested one named typed refusal, and the two documents call this a shared contract. The observe table also relies on the distinction between observed contradiction and machinery fault without pinning the complete `PredicateResult` signal expected by the existing gate.

Converge on one exact class, or explicitly define two layer-specific classes and their mapping. For every self-declaration/widen, ambiguous-source, and broken-chain branch, state the exact predicate value, failure class, `MachineryFault` boolean, terminal-by-authority result, and bounded fixture expectation. Ambiguous/broken resolution must set `MachineryFault:true`; a self-widen contradiction must remain observed-false rather than entering the no-vantage path.

#### MR-4 - the proposed refusal output violates the locked I-PH contract

The refusal says it names the offending record and the resolved ancestor bound (`design:50`). The locked s8 I-PH contract requires verdicts, bounces, rows, and `failing_detail` to remain path-free, bounded, symbolic, and path-redacted, with no raw paths or effective values (`s8 design:138-141`). A resolved scope bound is precisely a set of governed path values and cannot be exposed in the seat-deliverable detail.

Revise the output to a bounded symbolic class and, if needed, a non-path correlation identifier. The fixture must prove the raw candidate paths and resolved ancestor path set do not appear in the refusal, bounce, claim row, or failing detail.

### Passed Pressure Checks

- Reading only the conductor-stamped `parent_edge` and never `parent_hint` is the correct trust boundary. The current engine stamps a resolved accepted relay ID into `PARENT_DISPATCH_ID`; m-1 still owns the requested fidelity confirmation.
- Nearest-bearing-wins is coherent once the shared grammar enforces monotonic narrowing. The no-declaration E0 carve-out and the item-10 attributable-LHS gate are honest and must remain.
- The resolution table correctly separates no vantage, observed drift, and machinery failure at the conceptual level. MR-3 asks for executable signal bytes, not a change to `{accepted,rejected,held}`.
- The co-sign correctly does not un-strike live `diff_paths subset-of scope_paths` evaluation before item 10 lands.

### Revision Acceptance Bar

1. Both owner halves use one exact path grammar, normalization contract, subset algorithm, and nested-narrowing rule.
2. Candidate-supply rejection is stated and tested as future s9 work, not an extant v8 property.
3. The refusal class contract and complete observe signal are byte-consistent across m-2 and m-3.
4. Every seat-visible failure surface remains symbolic, bounded, and path-redacted under I-PH.
5. The parent-edge-only walk, no-silent-green floor, item-10 carve-out, and m-1 fidelity leg remain intact.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s9-scopepaths-cosign` m-3 rev0; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, PLAN, IMPL, or lock token
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner revises the predicate leg against MR-1 through MR-4, reconciles the shared bytes with m-2, and returns rev1 for DESIGN-REVIEW; the scope predicate remains blocked and struck.
