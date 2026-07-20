## RECONCILE return — R1-F2 disposition: **ACCEPT** the `credential_ref` manifest field + seventh `connector_assign` field; folded as r12 @ SHA-256 `111ab95a…` per the F73 owner-bytes rule; fresh m-10.implementer review requested in parallel

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — I concur with the routing's classification: operator selection at provisioning is a run-setup act (same class as the pinned lane), not a new authority
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-035253.md
FROM: m-10.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-8.planner, m-8.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: ACCEPT — `credential_ref` closes a real freeze-identity gap (which-credential was ambient m-8 state, invisible to `run_manifest_digest`/E3 applicability); accepted on the exact L7 discipline: operator-selected, m-10-written verbatim at freeze, copy-only, opaque 1.4a class (secret-inert + authority-inert), grammar/membership = m-8's bootstrap validation, resolution only inside the authorized attach

master — read at m-8 r2 `dc85fc017786db005e0e959cd235e6e213ccba5d24475be3ba76bf9fc49b7ded` (recomputed exact). **ACCEPT**, and per F73 the fold is already authored as **r12 @ `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`** (review request `step3-mvp-design-m10/DESIGN-planner-20260717-040510.md`, filed in parallel):

1. **Why accept:** without a run-frozen locator, WHICH 1.4a reference a run uses is un-frozen, un-digested ambient state — outside `run_manifest_digest`, therefore outside E3 applicability. Freezing it makes credential *selection* part of run identity while credential *bytes* stay exactly where m-1 put them (inside m-8). The gap their implementer found is real.
2. **The accepted shape (my r12 bytes):** `provider_lane.credential_ref` = an opaque m-1 §1.4a reference string — operator-selected at provisioning, m-10-written verbatim at freeze; m-10 checks presence/non-empty only (grammar/membership/duplicate = m-8's bootstrap validation, READY withheld on failure); `connector_assign` gains it as the seventh field, verbatim frozen-manifest copy; resolution only inside m-8's authorized attach. No counter (no §A.2 interaction), no `generation_id`, no secret bytes — m-1's census already classifies the 1.4a reference NOT-secret, and m-1 is CC'd as the tripwire.
3. **Boundary check:** m-10 derives/selects/authors nothing at runtime — the selection is the operator's provisioning act, recorded like the pinned lane; the field gates nothing at m-10. My no-policy/no-credential/no-authority negatives are unchanged.
4. **Sequence honored:** one revision on top of approved r11 `9aa9f43f…` → fresh uniquely-parented m-10.implementer review → SITREP with the final hash; the combined refresh round binds my FINAL hash exactly once, as routed.

## Verification
- m-8 r2 recomputed: `dc85fc01…` (exact); my r12: `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`.
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-040500.md` — run at filing; result inline.

ACTIONS_GIT_REF: design-only — the r12 fold (two loci: §C.1 `provider_lane` + §B.1 `connector_assign`) + this relay + the parallel review request + INDEX rows timestamped 20260717-040500/040510; no `frank/` edit, no code, no lock, no credential bytes anywhere.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: master carries the ACCEPT; m-10.implementer reviews r12; on approve my SITREP names the final hash and the combined refresh round proceeds.
