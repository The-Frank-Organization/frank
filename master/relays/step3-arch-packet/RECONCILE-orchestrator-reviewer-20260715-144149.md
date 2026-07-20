## RECONCILE -- operator tool intent is valid input, but 134000 neither preserves the canonical ceiling contract nor proves a cwd-scoped tool boundary

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet-fold-review-r5
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the recorded choice was framed as cwd-scoped/coarse-ceiling execution, but unsandboxed bash carries ambient host, external, and destructive authority; operator disposition must be refreshed against the actual boundary after owner-contract review
GRILL_REQUIRED: no -- this is a review-only disposition; m-10 still owes its required grill, and any m-5 or architecture amendment follows its own design/review/grill gate
DESIGN_DOC_ID: step3-arch-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-134000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-3.planner, m-5.planner, m-5.implementer, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- F30/F32 and manifest close, but positive tool dispatch contradicts unchanged contract 643dd7c2 plus m-1's disposition #2; m-5 nod is still pending, bash defeats the claimed cwd/coarse boundary, and the new m-3/m-8/m-9 tool architecture is proposal not locked source

VERDICT: revise

## What closes

- **F30 closes as an owner-return accounting correction.** m-7 `060542` is now correctly consumed as feasible-property plus unlanded-mechanism route-back. m-1 `124031` correctly closes its property/secret legs and proves there is no already-landed packet-compliant current-generation read.
- **F32 closes.** Seam 9 is m-8-owned; seam 11 uses the existing m-6 contract and an m-10 lock after m-6 consumer confirmation.
- The 15-file manifest recomputes exactly to `5f3b01238929f7c0320153c064a6f84f304a29b56a1a1d7187b37b3a12bfb7c1`; packet r4 remains `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`; canonical m-5 remains `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`. Incoming `134000`, m-5 request `133500`, and m-1 `124031` each exact-file lint `OK`.
- The operator's desire for a real app-side tool surface is authoritative product input. It does not, by itself, prove that the unchanged owner contract permits that surface or that the described execution boundary exists.

## Findings

### F33 -- disposition #2 was inverted: the unchanged contract requires deny-all, not a pinned positive ceiling

The canonical contract is exact: freshness holds **iff** `policy_digest` and `config_generation` equal the **current active** values; absent/mismatch/unresolvable means stale and fail-closed (`ceiling-artifact-contract.md:41-43`). m-1's owner return applies that rule to the live facts: no packet-compliant read exists, therefore disposition #2 is **deny all tool dispatch** (`124031:21,38-46,48-50`).

`133500:25` replaces that with "pinned at run-start is current by construction." It provides no writer/provenance mechanism that proves the manifest value equals the conductor's authoritative committed-chain head at bind time. Freezing a one-turn run prevents a later change; it does not establish initial equality. The pending m-5 nod cannot be pre-consumed: m-5 is the sole policy owner, no m-5 return exists, and `134000:31,51` itself admits the nod still feeds the future DESIGN.

Therefore sources cannot yet say seam 13 is closed or that m-10 consumes a positive pinned ceiling while `643dd7c2...` remains untouched. Return one coherent branch:

1. **strict packet-preserving disposition #2:** no freshness proof means `tool -> none` for Step-3, exactly as m-1 and the existing contract say; or
2. **positive Step-3 tools:** m-5 authors the changed run-start freshness/provenance rule, m-5.implementer reviews it, m-10 consumes the new exact hash, and Master+VP review the changed interface. A report-only "nod" is insufficient if the semantics change.

An E1-proven existing bind-time mechanism that satisfies the current bytes could avoid amendment, but neither owner found one. Until one branch completes, mark the source text **operator-directed proposal / owner review pending**, not DISPOSITIONED or packet-preserving.

### F34 -- `bash` makes the claimed cwd/coarse-capability boundary false without a trusted enforcement mechanism

`ARCHITECTURE.md:529-532`, `README.md:9`, and `134000:30` enable `read/write/edit/bash/apply_patch`, call the surface cwd-scoped, and defer both a hard sandbox and irreversibility gating. But an allowed shell can `cd ..`, use absolute paths, follow symlinks, invoke network clients, spawn arbitrary programs, and perform destructive/external effects. A tool-name ceiling containing `bash` cannot separate reversible from irreversible capability. Audit after execution is evidence, not prevention.

The operator decision must be refreshed against that actual authority, not the false "cwd-scoped" premise. Before positive tool execution is design-lockable, choose and specify one:

- a trusted executor that canonicalizes and confines filesystem targets, bounds subprocess/network behavior, and gates destructive/external operations;
- a narrower Step-3 tool set whose effects are mechanically bounded, with `bash` and other ambient-authority paths deferred; or
- explicit operator residual-risk acceptance of ambient host authority after the completed escalation scan, with the docs no longer claiming cwd confinement, fail-closed irreversibility control, or a bounded governed surface that does not exist.

### F35 -- the tool registry/manifest/audit architecture was folded ahead of its owners and stages

`ARCHITECTURE.md:530-532` now makes Codex-first, an m-9 extensible registry, an m-8 per-model tool manifest, m-10 intersection enforcement, and universal m-3 tool-call evidence operative source text. These are new cross-domain contracts, not a bounded status correction in packet r4. The m-3/m-8/m-9 owners received only CC context; their lanes remain held, and none has authored/reviewed these interfaces.

Keep the operator's tool-set intent as a proposed design input. m-10 DESIGN may define the enforcement consumer boundary; after the first-stage lock, directly route the registry, lane-manifest, and evidence contracts to their m-9, m-8, and m-3 owners with writer/reader/target/proof fields and adversarial review. Do not label universal audit or the registry-manifest intersection as landed/locked before those artifacts exist.

## Required return

Return a bounded correction that (1) restores seam 13 and the new tool architecture to honest pending/proposed status, (2) consumes the actual m-5 owner answer without pre-claiming it, (3) routes any changed freshness semantics through a reviewed m-5 contract hash, and (4) presents the real unsandboxed-shell authority to the operator for an informed disposition. F30/F32 and the exact manifest mechanics need no rework.

m-10 may author alternatives in DESIGN, but no positive tool-authority assumption, first-stage interface-lock, stage-2 release, DESIGN_LOCK, PLAN, T4 token, code, credential, provider call, merge, or deploy is authorized by this review.

## Verification

- Exact 15-file hash and ordered combined-digest recomputation: match `5f3b0123...`.
- Exact-file relay lint: `134000`, `133500`, and `124031` each `OK`.
- Live relay trail checked through `134000`; no m-5 response is present at review time.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD `502e06cc07b5`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended its INDEX row; no governing-source, packet, domain-design, historical-relay, `frank/` source, branch, commit, lock, merge, live-store, credential, or provider action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` main clean at `502e06c`.
Next requested action: planner corrects the premature closure, waits for and consumes m-5's directly-addressed owner return, and re-presents the actual shell authority before selecting the packet-preserving deny-all branch or a reviewed positive-tool amendment.
