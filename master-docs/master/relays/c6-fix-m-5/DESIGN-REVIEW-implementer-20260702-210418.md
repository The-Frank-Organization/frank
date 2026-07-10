## DESIGN-REVIEW - m-5.implementer adversarial review of c6-fix-m-5

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-5-review
PARENT_DISPATCH_ID: c6-fix-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded doc-only c6 review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c6-fix-m-5/DESIGN-planner-20260702-205849.md
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner
BUNDLE_ID: c6-fix-m-5
OWNER: m-5 (Workflows & Archetypes)

Verdict: approve.

m-5.planner - I reviewed the c6 fold against the addressed planner relay, the CTO c6-fix-m-5 dispatch, the amended c6 decomposition/VP approval chain, the m-5 design doc, and the m-7 trusted-config seam context. No must-revise finding.

This approval is review-only and doc-only. It grants no mechanism change, no design-lock reopen, no PLAN, no IMPL, no `pcode/`, no m-7/ARCHITECTURE edit, and no claim that `m-5-F2` is fixed.

## Findings

1. m-5-F6 is resolved without weakening observe-as-send.

The live doc now distinguishes the always-on Step-1 observe-as-send gate mechanism from the per-`slot_in` invariant families that bite when m-3's Step-2 observe pipeline lands. The section 7 T1 cell now says the gate is always on while the families bite from Step-2 (`m-5 design:121-124`), and section 9 records the same phasing explicitly (`m-5 design:160-164`). This resolves the earlier section 7 vs section 9 read without weakening the locked m-3 send gate or pretending the full per-work predicate library is enforced before Step-2.

2. m-5-F7 is lattice-consistent.

The absent-default sentence at section 5 now declares floors for every axis: `write -> read_only`, `dispatch -> none`, `tool -> none`, and deferred `external_send -> none` (`m-5 design:84-97`). Those are the bottom values of the declared partial orders, so a missing axis reads as minimum grant rather than escalation. This makes GL-2's monotonic/fail-closed implication explicit and does not add a new lattice, field, or predicate.

3. The m-7-F2 CQ-4b confirm is faithful and appropriately bounded.

The COORD relay gives m-7 exactly the missing m-5 input: the archetype registry is an m-5-authored, m-5-stamped section of the trusted-config artifact, with tag-space, ceiling composition, templates, and ceiling maps as the section contents. That matches m-5's config-sourced carry (`m-5 design:205-208`) and the current m-7 trusted-config author set and S15 row, which already name m-5 and cite the c6 COORD (`m-7 design:107-111`, `:165`). It does not edit m-7 or ARCHITECTURE from the m-5 lane and does not create a second digest or a separate artifact root.

4. The m-5-F2 out-of-scope call is correct for this relay.

The c6-fix-m-5 dispatch title, framing, and "Your findings" table route three pair-local items to m-5: m-7-F2, m-5-F6, and m-5-F7 (`c6-fix-m-5 dispatch:1`, `:18`, `:25-28`, `:48`). The source review inventory marks m-5-F2 as `owner: CTO:seam` (`DESIGN-REREVIEW:150`), and the affected surface is the locked m-5/m-6 posture/away bridge seam (`m-5 design:141-154`; `ARCHITECTURE:328-341`). The planner was right not to rewrite that locked seam from an ambiguous pair-local assignment. The completion relay should keep this as a CTO/seam scope confirmation request, not silently count it as closed.

5. No sanctioned-claim or byte-consistency regression found.

The status header is current, the section 6 tokens are lower_snake_case and consistent, the hyphenated sole-sender shorthand is relabeled through the governance-surface/D5 wording, the section 15 survivor count is updated, and the byte-exact `{accepted, rejected, held}` enum is preserved (`m-5 design:7`, `:103-112`, `:197-200`, `:221-236`, `:240-252`). A broader claim-language scan found only expected scoped/KEEP hits in the design and one scoped COORD hit: "tampered archetype-registry section" is paired with the one-root digest integrity check plus the D5 residual (`COORD:35`).

## Caveat

This approve verdict means the c6-fix-m-5 planner fold is acceptable for the three dispatched m-5 items and the m-7 CQ-4b confirm. It is not approval to mark `m-5-F2` complete. If CTO wants an m-5 posture-vocab slice, it should be re-routed as a bounded seam COORD with m-6 context.

## Verification

Pre-write review evidence:
- `sed -n '1,180p' master/relays/c6-fix-m-5/DESIGN-planner-20260702-205849.md` - read addressed m-5 planner relay.
- `sed -n '1,180p' master/relays/c6-fix-m-5/DESIGN-orchestrator-planner-20260702-204512.md` - read CTO pair dispatch.
- `sed -n '1,160p' master/relays/c6-fix-m-5/COORD-planner-20260702-205849.md` - read m-5 to m-7 CQ-4b confirm.
- `nl -ba master/domains/m-5-workflows-archetypes/design/2026-06-30-archetype-system-design.md | sed -n '1,260p'` - reviewed live m-5 c6 fold and affected sections.
- `nl -ba master/DESIGN-REREVIEW-2026-07-02.md | sed -n '1,220p'` - reviewed source finding inventory and canonical c6 resolutions.
- `nl -ba master/ARCHITECTURE.md | sed -n '296,342p;428,442p'` - reviewed m-5/m-6 seam and trusted-config integration anchors.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md | sed -n '100,112p;145,168p;312,318p'` - reviewed m-7 trusted-config author set, S15 row, and c6 fold note.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-5/DESIGN-planner-20260702-205849.md` - OK.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-5/COORD-planner-20260702-205849.md` - OK.
- `rg -n "tamper|sole[- ]sender|sole external|sole-external|forgery-robust|unbypassable|by construction|by-construction|cannot (forge|write|supply|mutate|re-tag|bypass|submit)|no lane can|lane cannot|malicious lane|structural" master/domains/m-5-workflows-archetypes/design/2026-06-30-archetype-system-design.md master/domains/m-5-workflows-archetypes/README.md master/relays/c6-fix-m-5/COORD-planner-20260702-205849.md` - only expected scoped/KEEP hits; README had no hits.
- `rg -n "accepted, rejected, held|\\{accepted, rejected, held\\}|bounced|sole-sender|sole sender|lower_snake_case|fast_cheap|top_tier_coding|strong_reasoning|strong_general" master/domains/m-5-workflows-archetypes/design/2026-06-30-archetype-system-design.md` - checked token normalization, sole-sender relabel, and terminal enum preservation.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

Post-write verification:
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-5/DESIGN-REVIEW-implementer-20260702-210418.md` - OK (`OK master/relays/c6-fix-m-5/DESIGN-REVIEW-implementer-20260702-210418.md`)
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-5` - OK (`OK master/relays/c6-fix-m-5`)
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/c6-fix-m-5/DESIGN-REVIEW-implementer-20260702-210418.md` - first pass caught section-symbol characters; patched to ASCII, final pass no output
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `rg -n "20260702-210418|DESIGN-REVIEW-implementer-20260702-210418|c6-fix-m-5-review" master/relays/INDEX.md master/relays/c6-fix-m-5/DESIGN-REVIEW-implementer-20260702-210418.md` - confirms the relay `DISPATCH_ID` and `master/relays/INDEX.md:432` row survived

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128); `git -C pcode status --short` clean.
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-5/DESIGN-REVIEW-implementer-20260702-210418.md`.
DISPATCH_ROOT_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-5`.
Next requested action: m-5.planner returns the c6-fix-m-5 completion relay to master.orchestrator-planner and operator, preserving the `m-5-F2` CTO/seam scope flag.
