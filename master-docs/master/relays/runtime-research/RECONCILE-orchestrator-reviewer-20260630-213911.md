## RECONCILE -- master.orchestrator-reviewer / runtime research synthesis review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: runtime-research
PARENT_DISPATCH_ID: runtime-research
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review of research synthesis; operator decisions surfaced but not made here
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator

Verdict: approve with tightened spike gates.

I reviewed:
- `master/relays/runtime-research/RECONCILE-orchestrator-planner-20260630-213521.md`
- `master/RUNTIME-RESEARCH.md`
- `master/RUNTIME-RESEARCH-GPT.md`
- the incoming relay lint verdict for `runtime-research`

Finding 1 -- Section 14 verdicts are accepted, with one security-gate expansion.

Q1: accept the primary-source verification pass. The correction is important and credible:
- attach-mode may claim confusion resistance, not construction-grade confinement;
- `srt` wrapping may claim sandboxed defense-in-depth and a parent-owned broker egress path;
- "sole external sender by construction" remains spike-gated and is not an automatic consequence of wrapping.

The "sole egress path" vs "sole exfiltrator" split is fair. However, the five-property spike is necessary but
not sufficient for the stronger "destination + content controlled by construction" wording. Add an explicit broker
policy/evidence property before that claim can ever lock:

`all outbound broker requests must pass the conductor egress/content gate on canonicalized destination, method,
protocol, payload class, and seat/run evidence reference; generic CONNECT/raw proxy behavior and broad allowlisted
hosts are not enough.`

Without that property, a compromised in-jail seat could still use the one permitted path as an exfiltration channel
to an allowed endpoint. The claim can be "single mediated network path" after the current five properties; it cannot
be "sole external sender with destination and content control" until broker policy semantics are proven too.

Finding 2 -- the three-rung decision-1 ladder is the right framing.

Q2: concur. The ladder gives the operator a useful decision without pretending the later security milestone has
already been earned:
1. attach first: Step-1 viable, confusion-resistant;
2. `srt` wrap: defense-in-depth, still not final construction-grade control;
3. construction-grade external-send control: separate spike milestone.

The design consequence is also correct: choosing attach-first does not block later wrap work, and choosing wrap-first
does not authorize the construction-grade claim without the spike.

Finding 3 -- runtime-provided identity is correctly rejected.

Q3: concur. Codex `session_source` and client-supplied IDs are telemetry/product-routing inputs, not
forgery-robust identity. The simplification is not a gap as long as the m-1 follow-up explicitly states that
runtime identity fields are never accepted as `FROM` authority, and that conductor-owned per-seat channel/credential
isolation is the source of the stamp.

This should become a bounded m-1 design-completeness fix after the operator's runtime decision, not a runtime
research caveat left only in prose.

Finding 4 -- the external brief is weighted correctly.

Q4: concur. `master/RUNTIME-RESEARCH-GPT.md` is evidence, not a seat, not a relay authority, and not a design lock.
The GLM/Kimi discrepancy resolution is sound: the routable unit is not "model family" in isolation, but
`model x provider x serving x compat-mode`, with canary qualification before autonomy.

One additional guardrail: specific model-family tier claims are drift-prone and should not be locked as durable
architecture. The durable rule is the tuple/canary policy; current GLM/Kimi/Qwen/etc. status belongs in a refreshed
qualification record when a lane is actually brought up.

Finding 5 -- governance state remains correct.

I agree with the relay's negative authority claims:
- no design changed;
- no PLAN opened;
- no spike authorized;
- no pair re-engaged;
- no MUST readiness gate cleared;
- no runtime choice made for the operator.

Approved next actions:
- Preserve `master/RUNTIME-RESEARCH.md` as report-only research evidence.
- Bring decision 1 to the operator using the three-rung ladder.
- If the operator selects attach-first, record the honest Step-1 claim as confusion-resistant.
- If the operator selects wrap or construction-grade control, route the bounded m-1/m-3/m-4 follow-ups only after
  that operator decision.
- Add the broker-policy/content-gate property above to any future construction-grade egress spike.
- Treat model-lane autonomy as tuple-qualified and canary-backed at the time of lane bring-up, not locked from
  static research prose.

Not authorized:
- no PLAN phase opening;
- no implementation, source/pcode edit, branch, commit, PR, merge, or live verification;
- no pair re-engagement from this relay alone;
- no "sole external sender by construction" design claim until the expanded spike proves both path confinement and
  broker policy/content mediation;
- no durable lock of current model-family rankings without a fresh lane-qualification artifact.

Verification:
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/runtime-research/RECONCILE-orchestrator-reviewer-20260630-213911.md` -> OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/runtime-research` -> OK
- `git -C pcode status --short` -> clean, no output
- `git status --short` -> unavailable, `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
