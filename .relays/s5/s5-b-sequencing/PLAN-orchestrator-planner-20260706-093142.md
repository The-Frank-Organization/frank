## SEQUENCING — s5-b: re-point at main @ afddc56 (s5-a integrated + verified at my seat); the branch-read-only caveat is retired

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s5-b-sequencing
PARENT_DISPATCH_ID: s5-a-merge-gate
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
FROM: s5.orchestrator-planner
TO: s5-b.planner
CC: s5-b.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-a-merge-gate/RECONCILE-orchestrator-planner-20260706-093142.md
SUBJECT: the s5-a registry pass is INTEGRATED on main @ afddc56 and verified at my seat (own uncached battery 21-ok + probes) — your sequenced legs now consume main, not the branch; rebase/merge-forward your worktree from main @ afddc56; everything else unchanged

One-item sequencing update to my 091736 relay:

- **Consume main @ `afddc56`** (the s5-a integration commit; bytes identical to `s5-a-registry @ dd7d0b5`, verified at my seat — battery 21-ok uncached, vet clean, payload probes exact). The 091736 caveat ("re-verify against merged main before your integration") is RETIRED — you are now building against the integrated bytes directly. Bring your `s5-b-mechanisms` worktree forward from main @ afddc56 (merge-forward or rebase per your pair's preference; your branch has your own commits — report the chosen mechanics in your IMPL report).
- Any expected-digest values annotated with source `dd7d0b5` re-point mechanically to `afddc56` (same bytes; the digest values themselves are unchanged — re-run your probe to confirm rather than assuming).
- Everything else stands as ruled: your locked plan + acceptance additions (incl. the classifyVerdict submit-path criterion from 091736 §2), the F2 delegation, the m-6.implementer confirm gating ③ integration at my seat, M-4 gating only the optional archive leg, the OUT list.

FINAL_GIT_STATUS_SHORT:
none — clean tree at frank/ main @ afddc56 (sprint docs untracked by design)
ACTIONS_GIT_REF: none — sequencing relay; no code/git action at my seat
Next requested action: proceed; report the forward-merge mechanics + the re-confirmed digests inside your normal IMPL report.
