## SITREP — orchestrator (CTO) Step-1 RE-ENTRY boot (post-compaction anchor: where we are + the immediate task)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-orchestrator-planner
PARENT_DISPATCH_ID: master-boot-orchestrator-planner
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-planner
CC: operator
SUBJECT: Step-1 re-entry — design LOCKED+clean, product = frank, build repo = frank/ ; next task = decompose + dispatch the Step-1 PLAN (vertical-slice, VP-gated)

**Re-entry note to a freshly-compacted CTO.** Read the charter (`CLAUDE.md`) → this boot → `master/README.md` (live status) → `master/STEP-1-KICKOFF.md` (the Step-1 approach) → then act. Everything load-bearing is durable; this is the orientation.

### Where we are (Step-0 → Step-1 boundary, 2026-07-03)
- **Phase = Step-1 BUILD** (operator-opened; charter Cardinal #1/#2 transitioned; old charter archived `master/archive/CLAUDE-AUDIT-DESIGN-phase-20260703.md`).
- **Product is named `frank`** (postal *franking* = the trusted mark that authorizes a relay to pass; *frank* = candid, the claim-boundary ethos). The trusted-courier **core is still "the conductor"** (m-7 = frank's engine). Locked mechanism vocabulary (`{accepted, rejected, held}`, `submit/project/read`, FieldSpec) **unchanged** — brand layer only.
- **Build repo = `frank/`** (empty git repo, `main`, ready; formerly `pcode/` — historical relays/entries say `pcode`, read as `frank`).
- **Design-of-record is LOCKED + clean + VP-co-signed:** six-domain (c1–c3) + conductor-core (c4) + claim-sweep (c5) + re-review/cleanup (c6) + seam-hardening (c6.1/c6.1a) + the external-review §C4.3/I-PH amendment (`step1-prep`). A final 5-lane differential returned the **4 seam lanes CLEAN**. Spec = `master/ARCHITECTURE.md` + the 7 `master/domains/*/design/*.md`.

### The immediate task (do this next)
**Decompose + dispatch the Step-1 PLAN** per `master/STEP-1-KICKOFF.md`:
- **Vertical-slice-first** (external-review, GPT): Section 1 = the thinnest end-to-end relay through *all* layers (`mint→connect→submit→stamp→validate→lineage→append→project→deliver→gate-outbox`) with a **tiny MVP FieldSpec** + the interface guardrail + I-PH path hygiene — NOT m-1/m-2/m-7 as separate castles. Then thicken (§ proposed decomposition in the kickoff).
- **Hardened exit gate** (both reviews): adversarial + crash/replay fixtures (forged FROM, forbidden enum, bad parent, duplicate-sibling double-accept killed, `kill-9` mid-commit/mid-delivery + re-issued wake, corrupt-projection rebuild, replayed intake-id, dissolved-linter replay, I-PH, park/wake).
- **Owed-item projection promoted early** (scoped to *recorded* owed-items + the materialize-first rule — VP correction).
- **VP-gate the PLAN decomposition** before dispatch (loop the VP; it's a major decision). Pairs' Step-1 boots = their PLAN dispatch relays (cut them post-decompose). First real code lands in `frank/`.

### Owed carries to fold into the Step-1 PLAN (from the ARCHITECTURE §C4 ledger)
③ known-A / RAISE-ONLY NF · ⑤ ODB model-name egress · R2 `gate_referenceable`-per-column negatives · `GRILL_REQUIRED` FieldSpec row · (optional) `routing_escalation` §J2 member · the **code-layer interface-guardrail enforcement** (the one genuinely new build item) · **I-PH path-hygiene** fixture. Plus the framing: Step-1 = provenance + transport, **not** verified work (observe is Step-2); only the serialized-loop kill (+ R2) are operationally live in Step-1.

### Standing process lessons (playbook A.6 / Part D — apply in Step-1)
- **CTO cross-domain seam edits are the recurring weak point** (4 differential catches this session). Single-hand token convergence, run a **dedicated adversarial differential** at seam boundaries (presence-checks pass while formulas diverge), and require **owner-confirm** for any CTO edit to a pair-owned doc.
- **External fresh-eyes review at each step boundary** catches framing/claim/strategy issues the internal seam-differentials can't (it caught the tool-mediated-confusion gap + the vertical-slice strategy).
- **Owed items = typed records + a projection** (materialize-first); keep the claim boundary **tool-mediated-honest** (confusion-resistant ≠ malicious-proof; D5 covers shell-routed confusion).

### Team state
Seats = CTO (this seat) · VP (`master.orchestrator-reviewer`) · 7 pairs m-1…m-7 (planner+implementer). All operator-relayed sessions; operator hand-relays. Pairs stood down (to be re-engaged by the Step-1 PLAN dispatch). Relay trail + `master/relays/INDEX.md` = durable record.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/boot/master-boot-orchestrator-planner/SITREP-orchestrator-planner-20260703-042742.md` — (run below).
- Pointers exist: `master/STEP-1-KICKOFF.md`, `master/ARCHITECTURE.md` (§C4 ledger + §C4.3 I-PH), `master/README.md` (▶ STEP-1 BUILD OPENED), `master/RECONCILE.md` (cycle ledger), `master/CYCLE-PLAYBOOK.md` (A.6/Part D lessons), `frank/` (empty repo).
- `git status --short` — unavailable: cwd is not a git repo; `git -C frank status --short` clean (empty repo).

ACTIONS_GIT_REF: wrote this orchestrator Step-1 re-entry boot + appended `master/relays/INDEX.md`; no design/mechanism/code edit; no PLAN dispatched yet (awaits post-compaction decomposition).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; `frank/` is the code repo, empty).
Next requested action: after compaction, the CTO decomposes the Step-1 PLAN (vertical-slice per STEP-1-KICKOFF.md), VP-gates it, and dispatches per-section PLAN relays to the pairs; first code lands in `frank/`.
