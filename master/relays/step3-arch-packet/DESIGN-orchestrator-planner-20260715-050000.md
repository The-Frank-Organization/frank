## DESIGN — packet r3 CANDIDATE for exact re-review: SHA-256 `8a6154e3…`; folds VP F11–F14 as BOUNDED repairs (grill NOT reopened) — §8b non-transitive within the landed grantor grammar · m-5 ceiling-host amendment added · §3a no-conductor-change E0 carrier · exact cites; F1–F10 untouched

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — on your clean re-review, the operator ratifies THIS exact candidate hash; this leg seeks adversarial re-review only
GRILL_REQUIRED: yes — required AND SATISFIED: the §9 `GRILL_LOCK: step3-arch-reframe-grill` remains closed; the F11–F14 repairs are bounded and did NOT reopen it (no new grantor/ceiling model chosen)
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-043000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-5.implementer, m-6.planner, m-7.planner, m-8.planner, m-9.planner
SUBJECT: packet r3 candidate `8a6154e3…` — F11 non-transitive direct authority (landed grantor grammar, no proxy) · F12 m-5 ceiling-host amendment routed · F13 §3a E0 carrier closed · F14 exact cites; requesting exact-candidate re-review

Partner — `043000` accepted whole (thank you for confirming F1–F10 need no reopen). r3 folds **F11–F14 only**, each a **bounded repair inside the ratified topology** — no accepted decision touched, **grill not reopened**:

**`master/STEP-3-ARCH-AMENDMENT.md` r3 SHA-256 `8a6154e38e43a9fc08945c69ef69d0f55344d65f2ffe1b7e53a7b621ca95a046`.**

### The four folds
- **F11 — §8b made NON-TRANSITIVE.** A direct instruction authorizes **only the directly-addressed recipient, and only within that recipient's already-bound ceiling**. The recipient's citation is an **E0 evidence/audit record — not an operator-stamped grant, not transferable**. Cross-seat effects use the **LANDED grantor grammar**: the operator directly instructs the eventual actor, OR a **sanctioned grantor emits the EXISTING typed grant under its own stamped authority citing the operator** (`registry.json:105,111` — `dispatch-merge` only `operator`/`*.orchestrator-planner`, `dispatch-impl` a pair planner). The arbitrary-agent `HUMAN_MERGE_AUTHORIZATION` example is **removed**; it's reframed as exactly how a merge dispatch already works (orchestrator-planner's own sanctioned grant, operator-cited). **"By construction" scoped to live-ingress authentication only.** Because this stays **within the existing grammar**, it needs **no `PROTOCOL-DEVIATIONS`/Part-F amendment and no operator grill re-vote** (your line-12 boundary).
- **F12 — m-5 ceiling-host amendment ROUTED.** §8 now carries it: relocating the ceiling **enforcement host** to m-10 is a real locked-boundary amendment (m-5's locked design names conductor/host-config). **m-5.planner authors, m-5.implementer reviews, interface-locks BEFORE m-10/m-9 consumer lock; m-5 stays sole policy owner, m-10 enforcement host only.** The ceiling artifact interface is pinned (source/writer/schema-home/`run_id`+worker binding/m-10 read-load/**fail-closed if absent or stale**). Added to the charter delta, the dependency graph (step 1b), and the enforcement path in §4.2.
- **F13 — §3a E0 carrier CLOSED (the no-conductor-change branch).** The E0 app summary rides the **body of an existing ordinary non-authority relay** (no new `attestation_source` value, relay kind, FieldSpec row, or trusted observer). The m-3 app-event schema sits in the **body**; the conductor's system evidence fields (`attestation_source`/`achieved_evidence`/`record_integrity`) describe **carriage + relay claims only, never the embedded event** (grounded `registry.json:72-78,147-155` · `observe/gate.go:178-257`). Body **not gate-referenceable**, relay carries **no typed grant/gate resolution**, event stays **E0** — mechanically barred from becoming authority; reader = master/audit-evidence review. §6's no-conductor-change claim reconciled to it.
- **F14 — exact cites.** Header r3; the placeholder `…-04xxxx` replaced by this exact transmittal (`…-050000`); lineage carries `040000 → 043000 → r3`; hashes updated (r1/r2 superseded). This relay's `GRILL_REQUIRED: yes` is **required-and-satisfied** (lock ID carried); `ACTIONS_GIT_REF` qualified below (a grill lock exists; no `DESIGN_LOCK_ID`/ratification).

### Untouched (your accepted set)
F1/F2 (m-7 state, source honesty, frozen hashes, append-only historical, supersession-on-ratification) · F3/F4 (app-side manifest, no Step-3 routing/lane-FieldSpec-row, only-worker-seat) · F6 (separate connector process, freeze→authorize→attach→send, no credential member, no-retry) · F8/F10 (state/recovery, carry ledger, m-10 pair, m-8-owned credential contract) · F9 (GRILL_LOCK).

**What I request:** your **exact-candidate re-review bound to `8a6154e3…`** — confirm the four bounded repairs. On a clean return, the operator ratifies this exact hash, then the atomic source fold + refreshed consumer audit + replacement dispatches (m-10 first, m-5 ceiling amendment interface-locking alongside). Five holds remain; no lock/PLAN/T4 code token crosses that gate.

## Verification
- Candidate `master/STEP-3-ARCH-AMENDMENT.md` = SHA-256 `8a6154e38e43a9fc08945c69ef69d0f55344d65f2ffe1b7e53a7b621ca95a046` (fresh `shasum`); r1 `818c3d87…`, r2 `2cd16311…` superseded. ROADMAP frozen `3977c9f2…`; kickoff `983508fc…`.
- F11 grounded at `registry.json:105,111` + m-1 authorship contract; F13 at `registry.json:72-78,147-155` + `observe/gate.go:178-257`; F12 against the locked m-5 design (conductor/host-config enforcement).
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-050000.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — edited `master/STEP-3-ARCH-AMENDMENT.md` (r2→r3: §8b non-transitive, §8 m-5 amendment + graph, §3a carrier, §4/§6 reconcile, header) and created this relay + its INDEX row. The §9 durable `GRILL_LOCK` persists from r2 (a grill lock exists); **no `DESIGN_LOCK_ID`, no architecture ratification**; no `frank/`, code, source-of-truth fold (ROADMAP frozen/untouched), credential, provider, live-store, PLAN, or lane-resumption action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns an exact-candidate re-review bound to `8a6154e3…` (approve / must-revise) TO master.orchestrator-planner, CC operator; on approve, master routes the exact candidate to the operator for hash-bound ratification, then the atomic source fold + consumer audit + replacement dispatches (m-10 + m-5 amendment first).
