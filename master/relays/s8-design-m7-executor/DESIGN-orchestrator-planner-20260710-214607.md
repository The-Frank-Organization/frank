## DESIGN — s8 pre-design dispatch to m-7.planner (sole author): the EXECUTOR HOST design — the unprivileged suite-executor boundary with the claim scoped per the VP's F4 (handle absence ≠ ambient process authority); the grill must resolve the eleven enumerated questions; your Implementer design-reviews before the lock feeds the s8 PLAN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s8-design-m7-executor
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — design dispatch; the OS-sandbox election the grill may surface is an operator-owned durable question
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-executor
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-201733.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-3.planner
SUBJECT: supersedes the withdrawn `step2-prep/SITREP-…-193209` (report-only relays cannot direct design; consultation context only) — you are the SOLE AUTHOR of `s8-design-m7-executor`; the m-3 registry/probe face is the SEPARATE artifact `s8-design-m3-registry`; the VP's F4 claim discipline binds verbatim

**Your artifact (the m-7-owned executor host, one author, your FROM):**
1. **The claim, scoped honestly (F4, binding):** repo-resident suite code run by the executor can inherit environment variables, file descriptors, filesystem/network/process access, toolchain caches, a writable working tree, and child processes — an absence-set of conductor handles proves none of those are *provided APIs*; it is **not OS containment** and cannot support an unqualified "never writes". The design claim is exactly: **no conductor-governed mutation handles, and no writes to canonical conductor state through the provided surface** — with the same-uid ambient-access residual stated (the D5 family), **unless the operator elects an OS sandbox as a separate, separately-grilled design addition** (a durable operator question if your grill reaches it).
2. **The grill must resolve, at minimum (the VP's list, verbatim):** snapshot/worktree identity · working directory + writable-temp/cache policy · inherited environment and descriptors · credential/signing-secret scrubbing · filesystem/store/config/outbox path exposure · network and subprocess policy · timeout plus descendant/process-group termination · output/resource ceilings · cleanup after timeout/crash · operator-gated side-effect semantics · the exact binding of candidate bytes + declared inputs to the typed verdict.
3. **The NF isolation fixture:** reports **handle-surface absence separately from the accepted same-uid ambient residual** — the latter is never labeled "by construction".
4. **Boundary contract:** name `s8-design-m3-registry` (which checks are suite-class, the probe intent, the timeout/operator-gate policy inputs — m-3's) where your host consumes it; do not author m-3's content.

**Process (binding):** the **grill** per the design-grill skill (one question at a time; codebase-answerable from the locked record — your §8.4 absence-set idiom, the c4 lock, the s7a claim-ceiling precedent; operator forks durable) → `GRILL_LOCK_ID` folded into the doc → **m-7.implementer returns `PHASE: DESIGN-REVIEW`** → the reviewed artifact to master. Master reconciles with `s8-design-m3-registry`; **the s8 PLAN consumes only the reconciled locks** (this design build-gates the spine's first E2 check).

Next requested action: your grill → the design doc → your Implementer's DESIGN-REVIEW → the reviewed artifact to master.

ACTIONS_GIT_REF: none — design dispatch only.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `2e1b4f0` (tag `s7-close`); cwd is not a git repo (docs workspace).
