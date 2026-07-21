## RECONCILE — m-8 stage-5 supervision-seam consumer confirmation (leg 2 of 4): CONFIRM — m-10's stage-5 realization r8 supervises m-8 exactly as frozen r12 expects; byte-bound r8 × r12, no finding

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a byte-bound consumer confirmation over pair-approved bytes; the operator gates at the stage-6 lock
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260720-070737.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260720-090000.md

**CONFIRM.** m-10's stage-5 control-plane realization (`2026-07-19-mvp-control-plane.md`, r8 @ `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`, recomputed this session) supervises the m-8 connector exactly as my frozen provider contract (r12 @ `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`) expects. No conflict; not a finding. The realization consumes m-8 r12 byte-bound (r8 §Consumed:6) and carries the wire/handoff shapes from the seam contract r36 `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01` (verified this session — the seven-field `connector_assign` is byte-identical at r36, and it is my r12 basis chain's current letter).

### Locus 1 — connector-first launch order (`connector_ready` gates AUTHORITY) — CONFIRM
r8 §3 steps 4–5 + §5: the PAIR launches (G1 allocated, both children spawned); a waiting worker candidate may reach pre-lease READY on `hello` but receives **NO lease, NO `assign`, NO admission** until the connector runs `hello → connector_assign → connector_ready`; **only on `connector_ready`** do lease-bind / `assign` / attach-gate / first `turn_open` proceed; the scheduler refuses turn admission while the incarnation is not READY. This is exactly my r12 §5.3 ordering gate — *"no worker admission, no DATA-P acceptance, and no provider-send path exists for this incarnation before the [connector_ready] ack."* The realization enforces my gate; `connector_ready` gates authority, never the pair's existence.

### Locus 2 — seven-field `connector_assign` + opaque `credential_ref` orchestration — CONFIRM
r8 §5 + the §11a `m10-connector-assign-credential` census row: `connector_assign` carries the **seven fields byte-verbatim from the frozen manifest** (r36 §B.1 — m-10 derives/authors nothing); `credential_ref` is operator-selected, m-10-written verbatim, **m-10 never resolves/opens/logs/validates beyond presence/grammar** (m-1 §1.4a), resolution is m-8's inside the authorized attach only, and a mismatch ⇒ `connector_ready` withheld ⇒ zero send, no admission. This matches my r12 §3 (the opaque 1.4a reference; m-10 orchestrates the reference never the bytes) and §5.3 (m-8 verifies its loads against the seven fields; mismatch ⇒ READY withheld ⇒ no admission). The seven-field shape is byte-identical at r36 (verified). `exclusive_credential_holder: m-8 at runtime (secret bytes); m-10 holds the opaque reference only` — my secret-boundary consumption is honored.

### Locus 3 — paired pre-ready-failure disposition (G-2/F4 fold) — CONFIRM
r8 §4 crash-loop policy + §5 co-restart: a connector failure BEFORE readiness rides the frozen paired lifecycle — **no connector-only path** (the supervisor has no "restart connector only" verb, by design); death of either DATA-P owner replaces BOTH under the new epoch (§A.1); the §B.4 canonical transition retires the paired worker generation and **mints exactly ONE E+1** in **ONE retirement disposition** covering the paired owners; then the **§B.5 distribution/install runs BEFORE any successor authority**; both owners are reaped; the retry spawns the fresh PAIR at E+1 under the frozen order (connector bootstrap → `connector_ready` → lease-bind → `assign`). This is consistent with my r12 §1.4/§3/§5 crash/retirement expectations — *"a connector-only restart does not exist; the co-restart pairs both DATA-P owners under the new epoch; the parked attempt is never retried automatically (one attempt per invocation); m-10 retains no DATA-P endpoint to re-issue."* The realization's E+1-then-install-then-fresh-pair sequence honors my UNKNOWN-park + co-restart contract; nothing lands at an m-8 seam that my contract does not already expect.

**Supervision policy above my contract (noted, no conflict):** the G-2 10-try exponential-backoff terminal (10th consecutive failure ⇒ run FAILED, counter over BOTH supervised children incl. paired pre-`connector_ready` connector failures) is m-10's operator-ratified supervision decision. It bounds retries and never touches an m-8 wire/row/fence/bootstrap obligation — my contract asks only for the paired co-restart + E+1 + UNKNOWN park, all preserved.

### Verification
- m-10 stage-5 realization r8 SHA-256 recomputed: `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa` (the confirm target).
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` (unchanged, unedited).
- Seam-contract r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`; the seven-field `connector_assign` byte-identical there (the already-confirmed wire-shape basis).
- Confirmed loci: r8 §3:45-46 · §5:65-67 · §4:61 · §11a:128 · §12 G-2:152 · §14:166.

Claims:
- All three confirmation loci verified verbatim against the r8 bytes and my frozen r12; the seven-field wire shape verified byte-identical at r36 — evidence E1.
- No conflict found; my r12 stays frozen and unedited at `4b670a79…` — evidence E1.
- No lock, PLAN, T4 token, code, credential use, provider call, or deploy performed or claimed — evidence E1.

ACTIONS_GIT_REF: docs-only in non-git cwd — this confirmation relay + one INDEX.md row; the approved design doc UNCHANGED at `4b670a79…`; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 6e4d657
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree (fresh status, empty output)

Next requested action: master carries this leg-2 supervision-seam confirmation into the stage-6 lock packet; m-8.planner holds for the Master+VP interface-lock.
