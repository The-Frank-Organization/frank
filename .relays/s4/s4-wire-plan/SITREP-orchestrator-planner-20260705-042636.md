## SITREP — m-1 verdict relayed as BINDING PLAN content: approve-conditional, F-S4-M1-1..6 verbatim; the redaction rule is the new sharp edge; no separate m-1 re-review if the plan carries them verbatim

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-wire-plan
PARENT_DISPATCH_ID: s4-fidelity-m1
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s4-fidelity-m1/SITREP-implementer-20260705-042308.md
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: s4-wire.implementer, s4.orchestrator-reviewer, operator
SUBJECT: m-1 fidelity supplement — dispatch condition 2 resolves from "await verdict" to "the gated PLAN carries F-S4-M1-1..6 VERBATIM with the redaction rule executable in fixtures"; your r1 plan (main@276d2b7, filed after the verdict) is checked against this by your implementer's PLAN-REVIEW; route-back triggers are binding

**The verdict** (`s4-fidelity-m1/SITREP-implementer-20260705-042308.md`, lint-clean, my run):
**approve-conditional for PLAN** — all five packet items approved (items 2/3/4 outright; 1 and 5
conditional), six conditions **F-S4-M1-1..6**, and m-1's own dispatch-condition paragraph: the
fidelity gate is satisfied when the gated PLAN carries the six conditions verbatim AND makes
the read/projection redaction rule executable in acceptance fixtures — **no separate m-1
narrow re-review in that case**. The named route-back triggers re-engage m-1 hard.

**The conditions, binding on the plan (fold verbatim; summaries here, m-1's text governs):**
- **F-S4-M1-1** `config_change` = the ONLY new record_kind; operator-channel provenance
  (never `system`); non-operator submit bounces typed + path-free; the registry home for the
  token moves from `system_only` posture to an explicitly operator-scoped rule — **if that
  cannot be expressed without broadening ordinary lane authority, route to m-2 and back to
  m-1 BEFORE implementation** (the S3 seat-scoped-enum machinery should express it; if the
  plan finds otherwise, that consult comes through me).
- **F-S4-M1-2 (the sharp edge — new, not in the design's r3 text):** the full-body record is
  approved for canonical recovery, BUT non-operator seat-facing surfaces (`read`, projection,
  nudge, tool result, schema, prompt, error) must NOT expose effective config member bytes.
  The broad authenticated `read(relay_id)` facade must either serve a REDACTED view for
  `config_change` (visible at minimum: relay_id, envelope source/version, record_kind,
  member, new_digest, a typed redaction marker) or reject non-operator reads of it. Redaction
  is a channel/view rule ABOVE raw store bytes — canonical records and checksums untouched.
  Also: `new_digest` is CONDUCTOR-RECOMPUTED from the replacement member + committed members;
  a payload-carried digest is a claim to verify, never authority. This condition must appear
  in the plan as executable acceptance fixtures (non-operator read of a config_change ⇒
  redacted/rejected; operator/admin path sees the bytes; digest recompute asserted).
- **F-S4-M1-3** the active-channel index = an in-memory admission cache only: SHA-256 hash on
  serverConn, `active` map under the existing lock, typed `auth:channel-active`, kernel-close
  cleanup ONLY; never persisted, never in tool output, never logged as a credential
  surrogate; binding table + raw-credential posture unchanged.
- **F-S4-M1-4** custody: env default, 0600-file secondary, no credential CLI flag; socket
  path + credential absent from every MCP-visible byte; the D5 residual stated honestly; one
  seat = one credential; no rotation claims.
- **F-S4-M1-5** per-recipient wake: recipients = exactly the committed record's recipients
  from stamped auth metadata; nudges carry only path-free cross-seat-safe metadata (relay id
  + generic marker); no new verb/authority.
- **F-S4-M1-6** `config/` demotion: fresh `store.Init` untouched; one canonical pivot under
  the old config; derived files re-materialized from committed records only; fail-closed for
  serving reads on unpersistable materialization; no second checksum root / side journal /
  hidden migration state.

**Route-back triggers (m-1's list, binding):** exposing full config_change bodies to
non-operator seats · `system` provenance for operator config changes · binding-table shape
change · persisting the active index · heartbeat/takeover/supersede/rotation · widening wake
recipients · changing fresh `store.Init` · any config authority outside canonical records.

**Sequencing note:** your r1 gated PLAN (main@276d2b7) was filed after the verdict landed;
whether the six conditions are already in it verbatim is exactly what your implementer's
PLAN-REVIEW now checks — this supplement makes any gap a mechanical must-revise, not a drift
risk (the S3 pattern). The F-S4-M1-2 redaction rule is design-adjacent but sits inside the
slot the design explicitly deferred to this packet (§6.1 "token + provenance ride the m-1
fidelity packet"); it lands in the PLAN with fixtures — no design r4 required. If folding it
surfaces a contradiction with locked design text, stop and route to me (never self-amend).

ACTIONS_GIT_REF: wrote this supplement + an INDEX row (relay substrate, git-untracked); no frank/ edits.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 276d2b7)
Next requested action: fold F-S4-M1-1..6 verbatim into the gated PLAN (or verify already carried); your implementer's PLAN-REVIEW checks the carry; the fence ruling on README rides a parallel relay under your fence-ask thread.
