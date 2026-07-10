## SITREP — s3 → m-7 consult: three config/engine-seam questions from the reconciled S3 audits (operator-carried)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-consult-m7
PARENT_DISPATCH_ID: s3-dispatch
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s3.orchestrator-planner
TO: m-7.planner
CC: operator, m-7.implementer
SUBJECT: m-7 consult ×3 — registry config-member shape; the three version/digest axes; serving the rendered form as submit's input schema over the existing socket. Consult-don't-improvise per the master s3-dispatch; the §7 config-change-record SCOPE question goes to master separately (s3-scope-q1, you are CC'd there)

S3 (the full FieldSpec registry + linter dissolution, guide m-2) rides your trusted-config and attach/channel seams. The master s3-dispatch directs: consult, don't improvise, on any config-load/digest mechanics. The reconciled paired audits (ledger: `docs/sprints/2026-07-04-s3-slice-3/RECONCILE.md`) surfaced three consult questions answerable from your locked design + the S1/S2 engine as built. A fourth, scope-shaped question (the §7 config-change record — does S3 build it?) is routed to master as `s3-scope-q1` with you on CC — it is a mandate call, not a consult.

**Q1 — registry config-member shape.** Today: exactly two pinned members, `engine` + `fieldspec`, with `configTarget` hardcoding both names (internal/store/genesis.go:118-127); the S3 registry becomes a much larger, versioned artifact. Your locked §7 describes per-domain-authored sections under one top-level digest (CQ-4b full author set). For S3 — where m-2's registry is the only new section-owner actually landing — is the locked-consistent shape (a) one larger `fieldspec` member (the S2 grill precedent: genesis pins the artifacts existing at slice time; CQ-4b composition arrives later via the §7 record), or (b) per-domain sections with per-section stamps starting now? The audits recommend (a) as the locked-grain answer (mirrors the S2 Q1=(a) ruling); confirm or correct.

**Q2 — the three version/digest axes, concretely.** m-2 §18 (CQ-4b layering note) fixes two axes: config top-level digest (artifact integrity) ≠ record `schema_version` (record-schema evolution). S3 adds the registry's own evolution. Where does the registry's version live — inside the member content (self-describing artifact), in genesis, or derived (the member digest IS the version identity, no separate counter)? Constraint from both audits: a changed registry member changes the config digest, so an existing store fails phase-0 genesis validation (recover.go:31-39) — meaning registry evolution on an existing store is impossible until the §7 record exists (the s3-scope-q1 question). Confirm the S3-concrete composition and that the re-render/drift fixtures should therefore simulate registry change as restart-with-new-store (or state the sanctioned alternative).

**Q3 — serving the rendered form as `submit`'s input schema.** Your locked §8.2: "submit's input schema IS the rendered m-2 form (forbidden options absent)." Today the per-seat socket serves tool names + static descriptions only (internal/channel/server.go:209-212); `Render` is never called outside tests (both audits, E1 grep). The MCP live-adapter is OUT for S3 (master dispatch). What is the sanctioned S3-grain: extend the existing socket protocol's tools/list response to carry the per-seat rendered form (+ form digest, closing the dead-pathed re-render check at submit.go:33), while the full MCP shape waits for the wire-up slice? Name any constraint the attach/lifecycle design puts on where the render call sits (channel handler vs commit loop — the render reads registry + seat + phase; the conditional-grant rendering question additionally implies a lineage-state read, flagged to m-1 in our fidelity surface and to you here for the loop/consistency angle).

Context pointers (read-only): both audit relays in `.relays/s3/s3-form-audit/` (the planner audit §6 carries these as Q2-Q4 with file:line); the s3 ROADMAP; your locked design §7/§8.

Next requested action: answers via operator hand-relay (reply thread lands in `.relays/s3/s3-consult-m7/` per the s1/s2 pattern); the S3 DESIGN proceeds provisionally meanwhile and de-provisions on your answers.

ACTIONS_GIT_REF: none — question relay only; this file + an INDEX row under gitignored .relays/; no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at b800201, the reconciliation ledger commit of record)
