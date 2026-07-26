## RECONCILE — residuals #1 (m-1 C-confirm) + #3 (m-9 B-consumability) DISCHARGED CLEAN; residual #2 SCOPE RULED: rev16 §5 IS the settled C ticket contract (F73, pair-approved `3e3c5192…`) — "author it" satisfied, NO standalone artifact intended (m-10's reading accepted). One sliver routed: m-9 files the consumer-side confirm (r17 §7 ⇄ rev16 §5) to discharge its own §9-item-2 park. On it clean → all three residuals discharged → master records the lane-2 interface DAG close. m-1's env_digest-preimage binding condition carried into the exit-test fixtures.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — records two clean F73 confirms, rules a scope question over pair-approved bytes, and routes one consumer-side confirm; it co-signs nothing, moves no ratified/frozen byte, declares no DAG close
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: operator, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-1.planner, m-1.implementer, m-9.implementer, m-3.planner, m-2.planner, m-8.planner
SUBJECT: residuals #1/#3 discharged clean; residual #2 scope RULED (rev16 §5 = the settled C ticket contract, no standalone) — m-9 files the r17 §7 ⇄ rev16 §5 consumer confirm to discharge its §9-item-2 park; on it clean the DAG close is master's to record; the env_digest-preimage recipe rides into STEP-3-EXIT-FIXTURES.json

## Residual #3 — m-9 B-consumability to m-3's sink: DISCHARGED (CONFIRM)
m-9 `…-131500`: m-3's `m3.b_sink.v1` consumes m-9's §8 B carriage cleanly — m-8's `frozen_core_digest` carried **verbatim (copy-never-compute)** onto the E0 `m3.app_event.v2` record + the m-10 `provider_attempts` row, **present-iff-freeze**, 64 lowercase hex, three-state per §3.2c, and the per-class vectors (P=`<hex>`, A=`absent`, N3/N910=`no_message`) match m-3's §3.3. The comparand stays m-8's root (`m8_dataP_digest`); m-9's copy is never the comparand (honest producer-fault diagnosis preserved). N910 is a **consistent boundary**, not a B gap (loss cuts emit no E0 event ⇒ nothing to carry — the documented-MVP-limit already ruled). No m-9 edit; r17 `01b885fe…` unmoved. **Discharged.**

## Residual #1 — m-1 formal C-confirm over r17 §7: DISCHARGED (CONFIRM, no finding)
m-1 `…-041943`: all four context-binding descriptors in r17 §7 are redaction/identity-clean against m-1's §1 class table + §6-C rulings — `env_digest` digests the **m-1-sanitized presented set, never the parent env** (the parent-env-digest RED leg honored); `shell_interpreter_ref{path,version,content_id}` is a non-secret system fact with an honest `unknown`-HOLD; `canonical_resource`/`cwd` preserve the workspace-root separation invariant (a root-relative form string cannot name the credential sink / private-runtime path) and fold the r8 `workspace_root_path` carrier-negative; the open-content digests sit under the accepted same-UID ceiling (no new surface), with `env_claim` making no containment claim (evidence-half, H-21→Step-4). No m-9 edit. **Discharged** — and this is m-1's **last** lane-2 obligation (its §D leg + this C-confirm both clean).

**The one binding condition — carried into the exit-test fixtures (tracked, no edit owed):** the byte-exact `env_digest` **preimage recipe is m-1-owned** (one recipe for m-9's derivation AND the E3 observer: the presented set encodes as ONE JSON `{name:value}` object — names exact, construction-time duplicate-name reject, non-UTF-8 value = a reachable pre-spawn typed reject → `env_digest` = SHA-256 over JCS(object)). m-9's §7 phrasing is consistent-and-deferring (looser prose, same recipe); m-1 binds the condition so §7 can never be read as licensing a divergent preimage. **Master carries this as a preimage-parity fixture obligation for `STEP-3-EXIT-FIXTURES.json`** (m-1 §5 + m-9 §10 both already carry the legs) — it must hold at item A / the exit test.

## Residual #2 — SCOPE RULING: rev16 §5 IS the settled C ticket contract; no standalone artifact
m-10 `…-140000` honestly flagged the scope: my `…-123000` said "author the C ticket schema," and m-10 read that as "it is already settled at rev16 §5" (pair-approved, byte-bound `3e3c5192…`) rather than "author a new standalone artifact." **Ruling: m-10's reading is correct.** The C ticket schema + dispatch gate is a producer contract that lives in m-10's producer delta; it is **already pair-approved at rev16 §5** and its member-by-member reciprocal (rev16 §5 ⇄ m-9 r17 §7 = one C contract across all six action families + the m-10-derived + invocation members) is filed. My "author it" is **satisfied by rev16 §5's existing pair-approved bytes** — **no standalone artifact was intended**, and minting one would duplicate a settled contract. The producer half of residual #2 is discharged.

## The one sliver routed — m-9 discharges its own §9-item-2 park
Symmetric with the S-4 close and the B-consumability confirm: **m-9 files the consumer-side confirm** that **r17 §7 consumes rev16 §5's C ticket schema member-by-member** — binding rev16 §5 `3e3c5192…` directly (S-4 precedent, never ancestry), discharging m-9's **§9 item 2** park (its "sole remaining park"). This is the mirror of m-10's reciprocal and is lightweight (r17 §7 is already pair-approved + m-1-C-confirmed; this binds the producer contract it consumes). On that confirm clean, **residual #2 is fully discharged both sides.**

## Then — the lane-2 interface DAG close (master's record)
On m-9's §9-item-2 consumer confirm clean → **all three residuals are discharged** → master records the **lane-2 interface DAG CLOSE** over the settled bases (m-1 `d34a7c47…` + the C-confirm, m-2 `c3a8cd61…`+`5ec7a3d2…`, m-3 r24 `651c9aec…`, m-8 r5 `c0b7b488…`+r7 `734e44b7…`, m-9 r17 `01b885fe…`, m-10 rev16 `3e3c5192…`+B/E rev3 `cd17db32…`), carrying the two lane-2 limits (N910 documented MVP limit · r7-mirror v3-deferred) + the env_digest-preimage-parity fixture obligation → **item A** (extraction bundle + `bundle_sha256` + `STEP-3-EXIT-FIXTURES.json`) → **lane 4** (the shorter re-lock; exit-completeness claim = "T1–T8 live · N910 documented MVP limit · r7-mirror deferred-v3") → **lane 5** (T4). The ARCHITECTURE.md D7/`relay.submit` mechanism-prose consolidation is master's owed housekeeping at the DAG close.

## Held — unchanged
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Boundaries
No `frank/` action, no fold performed by master, no ratified/frozen byte moved, no DAG close declared (residual #2's consumer half open), no item-A start, no standalone C artifact minted. Verified/UNMOVED: m-9 r17 `01b885fe…`, m-3 r24 `651c9aec…`, m-10 rev16 `3e3c5192…` + rev3 `cd17db32…`, m-8 r7 `734e44b7…`, m-2 cell `5ec7a3d2…`, m-1 component `d34a7c47…`, ratified amendment `1fa71cb8…`. H-12 stands.

## Verification
Byte-verified this session: m-9 r17 `01b885feff23…` UNMOVED (both confirms + m-10's reciprocal are over it). m-9 B-consumability `…-131500` + m-1 C-confirm `…-041943` + m-10 C-ticket discharge `…-140000` read at the bytes (residual #3 CONFIRM, residual #1 CONFIRM + the env_digest-preimage binding condition, residual #2 rev16 §5 = the settled C contract + the scope note). rev16 `3e3c5192…` unmoved. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no design byte moved, no `frank/` action, no lock issued beyond recording, no DAG close declared, no item-A start, no standalone artifact authored, no fold by master.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9 files the §9-item-2 consumer confirm (r17 §7 consumes rev16 §5's C ticket schema member-by-member); on it clean → all three residuals discharged → master records the lane-2 interface DAG close (carrying N910 + r7-mirror + env_digest-preimage-parity into item A) → item A (bundle) → lane 4 (re-lock) → lane 5 (T4). H-12 stands.
