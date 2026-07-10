## Team s1-core — Slice-1 thin end-to-end conductor relay (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s1-core-audit
PARENT_DISPATCH_ID: s1-dispatch
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for the audit; the downstream S1 plan is m-7-guide + master-VP gated, and merge is a separate human gate at S1 close
FROM: s1.orchestrator-planner
TO: s1-core.planner, s1-core.implementer
CC: s1.orchestrator-reviewer, operator
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@43e513e
TARGET_BRANCH: main

Implementer phase scope — AUDIT.
Current scope: read-only inspection of the locked spec + this repo, safe read-only commands, independent paired audit, questions, and findings.
Not in current scope: source/test edits, implementation branches, commits, PRs, scaffolding, or prototype implementation.
Implementation begins only after a current relay under the active run's RELAY_ROOT contains the exact literal token `DISPATCH IMPL` bare, unfenced, un-backticked, alone on its own line, and addressed to the Implementer in `TO`, or a direct message to that single Implementer contains the same bare own-line token. Urgency is not dispatch; inline, quoted, fenced, CC-only, cross-read, or non-addressee mentions are inert. For this run there is an ADDITIONAL hard gate above pair delegation: no `DISPATCH IMPL` is live until the S1 plan passes the m-7-guide + master-VP gate and the m-1/m-2 fidelity reviews.

Pair roles:
- Planner audits via Superpowers brainstorming and surfaces design questions.
- Implementer runs an independent audit and answers questions.
- Planner does not implement or spawn the Implementer.

Context:
Master dispatch `s1-dispatch` (../.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md) charters this team to build Slice-1: `mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox` in this greenfield repo. Sprint mandate, team table, exit gate, and scope-OUT list: docs/sprints/2026-07-03-s1-slice-1/ROADMAP.md.

Bundle:
The whole slice — engine skeleton (single-threaded serialized commit loop + crash-atomic rename commit), channel/identity (connect binds the channel; FROM system-stamped), interface guardrail (seats reach only `submit`/`project`/`read`; + I-PH path hygiene), tiny MVP FieldSpec, minimal validate + lineage, terminal append `{accepted | rejected | held}` (byte-exact), projection rebuild + read, deliver via project, one gate → local outbox/ODB item, S1-minimal dissolved-linter replay.

Owned surface:
The entire `frank/` codebase-to-be (currently: sprint scaffold only). This is ownership, not a prescriptive edit list.

Audit targets (read-only; all master paths are LOCKED spec — never edit, escalate problems via s1.orchestrator-planner):
- ../master/ARCHITECTURE.md §C4 (the engine) + §C4.3/I-PH (path hygiene) — extract what §C4 concretely requires of the S1 slice, incl. the owed §C4 carries (code-layer guardrail enforcement; I-PH fixture; the ③ known-A / RAISE-ONLY NF fixture, guardrail-adjacent portion).
- ../master/domains/ m-1 (FROZEN store API: `submit`/`project`/`read`, append-only, sole-writer), m-2 (FROZEN FieldSpec envelope), m-7 (conductor-core, our guide domain).
- ../master/STEP-1-KICKOFF.md + ../.relays/s1/s1-dispatch/… (our charter; the exit-gate list).
- The upstream shipped baseline + historical linter-failure corpus under ../extracted/ — locate the concrete failure set the S1-minimal dissolved-linter replay must run against (prior art: port principles, don't copy systems).
- This repo's actual state (expect: scaffold commit only).

Design questions to resolve (surface in audit; decide in DESIGN):
1. Runtime/language + process shape for the engine — what do the locked docs prescribe vs leave to us ("riding existing runtimes")?
2. Concrete MVP FieldSpec field set — the minimum this slice needs, per the m-2 envelope.
3. Crash-fixture mechanics — how to `kill -9` mid-commit deterministically (fault-injection points).
4. What `mint` / `connect` / park-wake minimally mean in S1, per m-1/m-7 docs.

Hard acceptance criteria (for the audit):
1. Every exit-gate line item (ROADMAP.md §exit gate) mapped to its spec source (file + section) and to a testable fixture idea, or flagged as a spec gap to escalate.
2. The frozen m-1/m-2 contract surfaces we will consume are enumerated verbatim (E1: file:line), so the later fidelity review has a definite object.
3. The owed §C4 carries written up as typed owed-item records `{owner, source, target surface, disposition path}`.
4. Historical replay corpus located (or its absence escalated).
5. Scope-OUT list re-checked against findings: anything that seems to force an OUT item is escalated, not absorbed.

Boundary contract (from the master dispatch — locked):
- Writes: append-only relay-store records (terminal {accepted,rejected,held}) + a local outbox/ODB item
- Reads: the locked m-1 store API + m-2 FieldSpec envelope
- Target entity: the conductor's committed relay + its rebuilt projection
- Downstream consumer: seat inboxes (via project) + the local outbox
- Contract: submit/project/read; channel-stamped FROM; byte-exact {accepted,rejected,held}; MVP FieldSpec; crash-atomic commit
- Proof: the S1-scoped hardened exit gate (E2 fixtures)
- No-consumer action: n/a — S1's consumer is the relay flow itself

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: repo is greenfield, but check the upstream baseline and the extracted export for anything the spec says to PORT (principles) vs anything we might wrongly rebuild or wrongly copy.
1. If any exit-gate item cannot be mapped to locked spec, flag it as a spec gap → escalate to s1.orchestrator-planner (who escalates to master); do not invent spec.
2. If the frozen m-1/m-2 contracts appear to conflict with the slice path, report the conflict; do not re-design a locked contract.

Anti-half-fix guards:
- No dead controls; target-entity semantics (the committed relay + rebuilt projection is the thing that must change); speculative-build reject if a component has no S1 consumer.
- Honesty framing binds the audit too: S1 = provenance + transport, not verified work; do not describe recorded-not-enforced properties as live.

Out of scope:
- Any edit anywhere (read-only phase). Any write to ../master or ../extracted ever.
- Designing the solution (that is the DESIGN phase; capture questions, not decisions).
- Everything on the ROADMAP.md scope-OUT list (S2/S3/S4/Step-2/Step-3 items).

Deliverable:
Independent audits from Planner and Implementer, each: 4-bucket verdict (PRIMARY_BUCKET expected `still-open` — greenfield — but run the duplicate gate honestly); spec-to-exit-gate map; frozen-contract surface enumeration (E1); typed owed-item records; replay-corpus location; design questions; boundary-contract assessment; evidence levels per claim; ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT (read-only phase — `ACTIONS: none` requires exact `git status --short` output). File relays under .relays/s1/s1-core-audit/, reply `TO: s1.orchestrator-planner`. No source changes, no PR.

ACTIONS_GIT_REF: wrote this audit-dispatch relay + an INDEX.md row under .relays/s1/ (gitignored relay substrate; no tracked-file change); tracked tree unchanged at main@43e513e
FINAL_GIT_STATUS_SHORT: none — clean tree
