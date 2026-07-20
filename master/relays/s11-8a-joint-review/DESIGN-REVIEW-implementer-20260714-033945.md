## DESIGN-REVIEW - m-6.implementer cross-confirm for 8a member 3 and bucket-D reason token

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-8a-joint-review
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/s11-8a-joint-review/DESIGN-orchestrator-planner-20260714-032510.md
FROM: m-6.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-2.implementer, s11.planner
BUNDLE_ID: m-6-human-surface-scheduler
DESIGN_REVIEW_VERDICT: approve

Verdict: approve / confirm from `m-6.implementer`. Member 3 remains consistent with the final m-2 rev3 contract, and the bucket-D rejected stale candidate should use a distinct author-return reason token: `stale_choice_set`.

## Cross-Confirm

1. Member 3 consumes the final m-2 rev3 signal as three byte-distinct records:
   - The stale operator reply candidate is `delivery_state: rejected`, bucket D, with no `resumed` wake and no projection intents.
   - The migration-fault signal is `delivery_state: held` with `failing_edge: stale_schema`, bucket A.
   - The replacement is a fresh gate/ODB with a new decision identity and resummon keys `(same seat, NEW decision_id, restarted cadence_slot series)`.

2. The replacement coupling remains the m-6-approved rule: emit the stale rejection, held signal, and replacement in the same serialized outcome, or commit a durable re-issue intent whose recovery replays to the same replacement decision identity. A crash after rejecting the stale reply but before a durable replacement exists remains a silent drop and is not acceptable.

3. A changed choice set never wakes or auto-resolves the old decision. It rejects the stale submitted candidate, parks/escalates the incompatible old gate state through the held record, and asks the operator again through a new decision identity.

4. The m-2 rev3 alias-safety folds strengthen the upstream detector but do not change the member-3 disposition contract: the live `classifyVerdict` path produces one deterministic typed signal after snapshot-before-Apply and Apply deep-clone; member 3 consumes that signal using the three-record disposition above.

5. Confirm `pi = {value -> label}` as the ODB decision identity from the m-6 surface side. `label` is decision-bearing because it is the operator-visible meaning of the offered choice. A relabel is a re-meaning and must trip the stale-choice path rather than silently resolving the old decision.

## Bucket-D Reason Token Ruling

Rule: the bucket-D rejected stale candidate uses `failing_edge: stale_choice_set`, not `stale_schema`.

Reasoning:

- `stale_schema` remains the m-2/member-1 migration-fault reason under `delivery_state: held`, bucket A.
- The bucket-D author-return case is different: the operator submitted a reply against an obsolete frozen choice set. That is a stale choice-set veto, not a migration engine failure.
- Reusing `stale_schema` would force `delivery_state` to carry too much semantic disambiguation at the exact author-return surface where the reason text should tell the operator what happened.
- A distinct token keeps operator triage honest: `stale_choice_set` means "your submitted choice belonged to an obsolete offered set; a replacement decision was issued."

Shape:

- Rejected stale candidate: `delivery_state: rejected`, `failing_edge: stale_choice_set`, bucket D, no wake, no old-decision resolution.
- Migration fault record: `delivery_state: held`, `failing_edge: stale_schema`, bucket A.
- Replacement gate/ODB: new decision identity, same serialized outcome or durable re-issue intent.

This does not add a fourth `delivery_state`, does not require an exhaustive enum, and does not change m-2's store shape. It is an open system-stamped `failing_edge` reason value for the rejected candidate only.

## Scope

This closes the `m-6.implementer` cross-confirm requested in ITEM 1 and returns this seat's bucket-D grammar ruling for ITEM 2. It does not proxy-author `m-6.planner`; if master requires an owner-planner echo, `m-6.planner` should file that separately.

Not authorized / not claimed: no code or `frank/` edit, no PLAN or IMPL dispatch, no merge authority, no design-doc mutation, no full 8a lock closure by this seat alone.

## Verification

- Read addressed relay: `master/relays/s11-8a-joint-review/DESIGN-orchestrator-planner-20260714-032510.md`.
- Read final m-2 rev3 design: `master/domains/m-2-forms-determinism/design/2026-07-14-s11-8a-frozen-choice-migration.md`.
- Read prior m-6 member-3 approval: `master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043.md`.
- Read concurrent CC-only `m-6.planner` owner ruling after this relay was written: `master/relays/s11-8a-joint-review/DESIGN-planner-20260714-033932.md`; it agrees on distinct `stale_choice_set` and leaves ITEM 1 to `m-6.implementer`.
- `find master/relays/s11-8a-joint-review -maxdepth 1 -type f -name 'DESIGN-*' -print | sort` - latest direct relay resolved to `DESIGN-orchestrator-planner-20260714-032510.md`.
- `tail -n 8 master/relays/INDEX.md` - latest row before this relay was the 032510 orchestrator dispatch.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-033945.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s11-8a-joint-review master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-033945.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s11-8a-joint-review/DESIGN-planner-20260714-033932.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s11-8a-joint-review master/relays/s11-8a-joint-review/DESIGN-planner-20260714-033932.md` - OK.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.
- `git -C frank status --short` - clean output.

ACTIONS_GIT_REF: wrote `master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-033945.md` and appended `master/relays/INDEX.md`; no `frank/` edit, no code/source edit, no PLAN, no IMPL, no design-doc mutation.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace); `frank/` status clean at `d91fcfb`.
Next requested action: master integrates this with the m-2 pair-approved leg and any separately required `m-6.planner` echo, then locks the s11 T6 8a contract.
