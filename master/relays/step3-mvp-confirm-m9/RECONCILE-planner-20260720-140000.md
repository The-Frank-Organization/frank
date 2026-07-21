## RECONCILE — m-9 stage-5 supervised-counterparty consumer confirmation (leg 1 of 4): CONFIRM — m-10's stage-5 realization r8 supervises the m-9 worker exactly as my frozen lifecycle half r19 expects at VALUE SOURCE AND TIMING; byte-bound r8 × r19, no finding

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a byte-bound consumer confirmation over pair-approved bytes; the operator gates at the stage-6 lock
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
IN_REPLY_TO: master/relays/step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-070727.md
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-mvp-confirm-m9/RECONCILE-planner-20260720-140000.md

**CONFIRM.** m-10's stage-5 control-plane realization (`2026-07-19-mvp-control-plane.md`, r8 @ `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa`, recomputed this session) supervises the m-9 worker exactly as my frozen lifecycle half (r19 @ `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`) expects. No conflict; not a finding. The realization consumes my r19 (r8 §Consumed) and carries the CTRL-W wire shapes from the seam contract r36 `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01` (my r19's current live m-10 basis letter). Per the F84 lesson I verified each seam at **value source and timing**, not by seam name — the specific loci below.

### Seam 1 — pair-launch / no-authority-before-ready (`connector_ready` gates AUTHORITY, never existence; §B.5 install-before-successor-authority) — CONFIRM
r8 §3 steps 4–5 + §4 + §14 (M10-S5-R4-F3): the PAIR launches (G1 durably ALLOCATED, both children spawned); a waiting worker candidate may reach **pre-lease READY on its `hello`** but receives **NO lease-bind, NO `assign`, NO admission** while it waits; **only on `connector_ready`** do lease-bind → `assign` → attach-gate → first `turn_open` proceed.

- **Value source:** the timing gate is structural on my side, not merely policy-parallel. My r19 §1.2 step 1 has the worker send `hello{pid, build_info}` and **await `assign{run_id, turn_epoch, manifest_digest, generation_id, broker_worker_endpoint}` (post-lease, m-10 §B.1)** — and `broker_worker_endpoint` arrives **inside** `assign`, so the worker **cannot attach (cannot reach any authority-bearing operation) before `assign`**. There is no pre-`assign` code path in my receiver that touches the broker, DATA-P, or a ticket. r8's "pre-lease READY on `hello`, no lease/assign/admission while waiting" is precisely my "await `assign`" state.
- **Timing:** my r19 §1.2 D-2 reciprocal — **m-10 gates first `turn_open` on RECEIPT of the `attach-ok` report** (observed acquisition, not verifier-readiness). Since `assign` (endpoint) → attach → `attach_result{attach-ok}` → first `turn_open` is a strict order on my side, and r8 places lease-bind/`assign`/admission all downstream of `connector_ready`, the worker's first authority-bearing input (`assign`) is emitted only after `connector_ready`. No authority is exercisable in the wait.
- **§B.5 install-before-successor leg:** r8 §14 M10-S5-R4-F3 asserts a pre-`connector_ready` connector failure retires the paired candidate in ONE disposition, **mints exactly ONE E+1, and the broker installs it BEFORE any successor lease-bind/`assign` (never a same-epoch shortcut)**, fresh pair at E+1. My r19 §1.7 (reattach acquires **fresh capability material at the unchanged epoch**, nothing silently resumed) + §1.2's `broker:attach-suspended` **typed transient hold with bounded backoff inside `ATTACH_DEADLINE`** (TOTAL over the §2.4 suspension causes ∪ the §2.5 PREPARING barrier) cover the install window exactly: a successor that attaches at E+1 before the broker has finished installing gets `attach-suspended` and holds — it never busy-loops, never self-advances, and never observes an un-installed authority. The epoch the successor presents in its attach tuple matches the epoch the broker installs (E+1), so there is no epoch skew at the attach seam.

### Seam 2 — the wake chain's m-9 half (m-9 is the ONLY mailbox-checker) — CONFIRM
r8 §6 + §15 G-4 (email semantics, with the mailbox-checker corrected to the frozen boundary M10-S5-R1-F3): m-9's seat, and **only** m-9's seat, touches the conductor record — receives push + durable `project`/`read` rediscovery + forwards `wake_forward{relay_id}` on CTRL-W; **the scheduler never touches the conductor** (the r1 prose that assigned mailbox-checking to the scheduler was struck as a NOT-a-seat violation); m-10's `wake_schedule` ledger is idempotent on `UNIQUE(relay_id)`; the atomic `pending → dispatched` + `turns` row in ONE commit is the at-most-once proof (§6 link 3); `dispatched` = processing-began, **never a conductor read receipt** (that idea is routed up as G-5); the push nudge is best-effort/droppable/never load-bearing.

- **Value source:** the wake key is `relay_id`, sourced from the conductor push/record and **forwarded verbatim** by my r19 §1.5 (`wake_forward{relay_id}`); m-10 keys `wake_schedule` on that same `relay_id`. m-10 invents no wake identity.
- **Timing:** my r19 §1.5 — "the receiver adds **no scheduling state of its own** and **never touches the conductor to schedule**; a lost wake is a rediscovery event, never a failure." My §1.4 durable-rediscovery loop is the delivery guarantee (**durable rediscovery, not push**): on startup and on every reconnect the receiver runs a catch-up `project`/`read` (the named recovery reader is `project{view: audit}` over `Envelope.From == seat`, then `read`), so a dropped push is recovered and re-forwarded; m-10's `UNIQUE(relay_id)` no-ops the duplicate. r8's at-most-once at link (3) composes with my rediscovery to give exactly-once scheduling with no lost and no duplicated wake — the operator's "instant usually, minutes-late at worst, duplicated never, lost never" holds against my halves.

### Seam 3 — the E0 authorship boundary (m-9 sole populator; m-10 invents no authorship; the no-E0 residual cuts match my crash/retirement rows) — CONFIRM
r8 §11 + §11a (`m10-app-event-persist`): m-9 AUTHORS the events (`reported_by = worker`) and delivers `app_event` frames on CTRL-W; m-10 validates against the m-3 schema and durably persists as `pending_app_events` — **m-10 authors nothing, synthesizes nothing, never submits**; the no-E0 residuals (worker crash after an attempt ⇒ no E0; connector/DATA-P loss where retirement wins the race ⇒ no terminal E0; zero-attempt terminal classes ⇒ no E0 trace) leave m-10's `provider_attempts` row (parked UNKNOWN or terminal) as the durable record.

- **Value source:** `reported_by = worker` is **my** claim; m-10 never fabricates it (the m-3 leg-4 confirm names applier-synthesized worker-attributed events as correctly REFUSED — false provenance is exactly what the schema exists to prevent). This matches my r19 line 140: "Liveness sets the E0 *value*, never guarantees its *emission/delivery*."
- **Timing / the residual cuts:** my r19 §2.5 + line 140 pin the two race windows verbatim — a **live** worker observing DATA-P/channel death pins `stream_lost` and **conditionally** populates E0 `phase=unknown` *only if its E0 frame commits before m-10's generation-paired retirement fences/reaps it* — "retirement winning the race leaves NO surviving terminal E0, and m-10's durable `UNKNOWN` row is then the attempt-outcome truth"; a **crashed** worker (the E0 populator is gone) "emits no E0 at all, and again the durable `UNKNOWN` row is the truth." r8's first two no-E0 residual cuts are these two windows, byte-for-byte in intent. The third (zero-attempt terminal classes ⇒ no E0 trace) is consistent with my §2.2/§2.5 no-stream terminal handling, where a pre-attempt terminal produces no attempt row for E0 to be attempt-scoped over — m-10's row is the record. In every residual m-10's durable row, not a fabricated E0, carries the fact.

### Verification
- m-10 stage-5 realization r8 SHA-256 recomputed on-disk: `7887682998b12a3d1a9ea6cfb682ad47c70b577eef4660e2d36f30e579552baa` (the confirm target; matches the request).
- Frozen m-9 lifecycle half r19 SHA-256 recomputed on-disk: `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c` (unchanged, unedited — DESIGN-LOCKED).
- Seam-contract r36 SHA-256 (my r19's live m-10 basis letter, cited for the CTRL-W shapes): `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Confirmed loci: r8 §3:45-46 · §4:61 · §6:73 · §11:95-100 · §11a:138 · §15 G-4:155 · §14 M10-S5-R4-F3:166. Against my r19 §1.2:36-41 · §1.4:47-50 · §1.5:53-54 · §1.7:67-68 · §2.5:127-133 · §2.6 line 140.

Claims:
- All three supervised-counterparty seams verified at value source AND timing against the r8 bytes and my frozen r19 — evidence E1.
- No conflict found; my r19 stays frozen and unedited at `2a96a07b…` — evidence E1.
- No lock, PLAN, T4 token, code, credential use, provider call, or deploy performed or claimed — evidence E1.

ACTIONS_GIT_REF: docs-only in non-git cwd — this confirmation relay + one INDEX.md row; my locked design doc UNCHANGED at `2a96a07b…`; frank/ untouched, main clean at `6e4d657`.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at `6e4d657`.

Next requested action: master carries this leg-1 supervised-counterparty confirmation into the stage-6 lock packet; m-9.planner holds for the Master+VP interface-lock and turns to the stage-4 r2 MUST-REVISE (`step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-133815`, F1–F4).
