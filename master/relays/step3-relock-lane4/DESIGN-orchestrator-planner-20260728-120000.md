## DESIGN — NESTED KICKOFF AUTHORED (INERT): `master/STEP-3-LANE4-KICKOFF-NESTED.md` `8bb3a0ac…`, to the VP-approved plan rev7 `e7a333e9…` + the r7-F2 instantiation guard. Concrete per-leg dispatch-ids (`…-l4-boot`/`-l4-decomp`/`-l4-w<k>`/`-l4-integrate`/`-l4-esc<n>`), immediate-predecessor PARENT + IN_REPLY_TO, root id never reused. Supersedes the VOID pair-shaped `STEP-3-LANE4-KICKOFF.md`. Conformance-checked (1 defect — a duplicate `-l4-decomp` id — caught + fixed). GRANTS NO LANE ACTION. Next = the operator + master set up the team together (mint/boot seats → zero-authority three-tier preflight → activation).

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the INERT kickoff grants no action; the operator alone mints + boots the seats, wires frank creds, and authorizes the zero-authority preflight, then supplies the post-preflight activation. Master does not mint/boot/activate.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-110000.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Nested-team INERT kickoff `8bb3a0ac66a60312bce52248a228b89989de546e8f6fa01426401239a196934f4` authored to rev7 `e7a333e9…` + the instantiation guard; ready for operator-led seat mint/boot + zero-authority preflight; Item A lock `cbd1893c…` preserved

## What this delivers
Per the VP approval (`step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-110000.md`: "master authors one fresh nested-team **INERT** kickoff … consistent with those bytes and the instantiation guard"), master has authored `master/STEP-3-LANE4-KICKOFF-NESTED.md` = `8bb3a0ac66a60312bce52248a228b89989de546e8f6fa01426401239a196934f4`. It:
- binds the approved plan rev7 `e7a333e9…` + GRILL_LOCK + B21 + §7 + the lock `cbd1893c…` as authority-of-record (plan governs on any difference);
- carries the **canonical seats** (`l4.orchestrator-planner`/`l4.orchestrator-reviewer` + `l4.w<k>.planner`/`.implementer`), the honest operator-mints/boots spawn model, the decomposition-review gate, the read-only proposal chain, the owner-fidelity matrix (guiding PM m-3), and the concrete **instantiation-guard** dispatch-id map (distinct per-leg ids, immediate-predecessor PARENT + IN_REPLY_TO, root id never reused);
- folds the PRIOR-ART §2c prior-art provenance (codex/paperclip/omp — independent arrival for the nested ceiling);
- is **INERT** and supersedes the voided pair-shaped `master/STEP-3-LANE4-KICKOFF.md`.

**Conformance-checked before this record:** one defect (a duplicate `…-l4-decomp` dispatch-id across the decomposition request/verdict rows — the exact shared-id resolver shape) caught and fixed → one leg, one id, within-thread ordering by IN_REPLY_TO. All else (plan fidelity, §7 schema, the lock SHA, the fidelity matrix, internal consistency) clean.

## Next — operator-led setup (jointly with master)
1. **Operator** mints (`seat_mint`) + boots as independent read-only sessions, with **zero authoring/dispatch authority**: `l4.orchestrator-planner`, `l4.orchestrator-reviewer`, and ≥1 probe worker (`l4.w1.planner`/`.implementer`); wires each seat's frank credential.
2. **Zero-authority preflight** (transport proof only): a real three-tier exchange exercising both orchestrator seats + ≥1 worker — accepted `submit`/`read` + correct immediate-predecessor PARENT/IN_REPLY_TO chaining per the §7 map (using the concrete `…-l4-boot`/`-l4-decomp`/`-l4-w1` ids) + a durable relay export. The probe carries **no** fixture-authoring authority.
3. **On pass → operator activation** turns the kickoff operative → the decomposition gate begins. **On fail → hold** (hand-relay fallback = operator-owned deviation; the gaps become battle reports).

## Boundaries
No seat minted/booted, no preflight, no activation, no decomposition, no worker dispatch, no fixtures/manifest, no freeze/re-lock, no PLAN-lock, no T4/code token, no credential wired, no provider call, no `frank/` action by this record. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. **H-12 hard-blocks external use.**

## Verification
`master/STEP-3-LANE4-KICKOFF-NESTED.md` = SHA-256 `8bb3a0ac66a60312bce52248a228b89989de546e8f6fa01426401239a196934f4`. Approved plan rev7 `e7a333e9…` + lock `cbd1893c…` cited, both immovable. Dispatch-id map: 5 legs, no duplicate concrete id, root `step3-relock-lane4` not reused as a leg id. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — authored `master/STEP-3-LANE4-KICKOFF-NESTED.md` (INERT) + this relay + one INDEX.md row; no seat mint/boot, no preflight, no activation, no fixtures/manifest/lock/`frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the operator (with master) mints/boots the l4 seats + runs the zero-authority three-tier preflight; on pass the operator activates and the decomposition gate opens. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
