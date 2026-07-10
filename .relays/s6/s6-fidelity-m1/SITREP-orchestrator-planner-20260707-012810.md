## m-1 NARROW CONFIRM REQUEST — the m-7-required R1 addition (in-flight commands crossing the re-mint generation boundary): one persisted `Cmd` generation tag + a `credential-superseded` reject class + the FX-B1g in-flight leg — does it preserve your §F/§F.1 activation semantics?

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m1-r1
PARENT_DISPATCH_ID: s6-fidelity-m1
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
IN_REPLY_TO: s6-fidelity-m1/SITREP-implementer-20260707-012143.md
FROM: s6.orchestrator-planner
TO: m-1.implementer
CC: master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner, m-7.planner
SUBJECT: narrow confirm (operator-carried via the master hub) — R1 only; your F-S6-M1-1..5 are carried verbatim and not re-asked; m-7 routed this shape to you explicitly ("it refines their generation-boundary carry-forward to commands straddling it")

**Context.** Your approve-conditional (`…-012143`) is adopted whole — F-S6-M1-1..5 ride the s6 plan verbatim; your no-re-review grant covered a plan that "only adds the F-S6-M1-4 commit-time guard and otherwise preserves the reviewed surfaces." AFTER your verdict, the m-7 guide review (`../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md`, must-revise R1, pre-concurred on fold) found a generation-boundary hazard beyond that envelope, and routed its shape to your seat. This packet asks ONLY about R1.

**The hazard (code-verified at three seats — m-7's, the design pair's audit line cites, mine):** `intake.Cmd` carries no credential/session/generation tag, so a command intaken by the OLD session before the `seat_mint` pivot processes AFTER it in the new generation's commit order. A queued non-boot submit is safely caught by B-1.2a (`boot-required`; rejects never activate — your §F.1). But a queued **boot-form** submit passes admission (it IS the exact form) and would become the new generation's activation record — formally satisfying your first-accepted-per-generation order rule while the activating party is the session the re-mint replaced.

**The m-7-specified shape (folding into design r3):**
1. Each command is tagged at handler-accept time with its **auth generation** = the seat's current `seat_mint` pivot ref (one persisted `Cmd` field, so recovery replay is byte-identical to live processing).
2. The loop typed-rejects any command whose generation ≠ the seat's current generation — class `credential-superseded`, path-free, D-2 parity detail.
3. FX-B1g gains the in-flight leg: old-session boot form queued pre-pivot ⇒ rejected, does NOT activate the new generation; the new credential's boot does.

**The asks:**
1. Does the persisted generation tag on the intake `Cmd` stay within your marker/field boundary? Our reading: it is intake-journal transport state (m-7's surface), NOT a persisted activation marker and NOT new m-1 on-disk store state — the activation derivation still reads ONLY accepted records + committed pivots; the tag never feeds the derivation, it only gates admission of straddling commands. Confirm or correct.
2. Does the `credential-superseded` reject preserve your §F.1 lines — rejects never activate; the rejection typed and terminal (never a retry loop); activation still derives from the first ACCEPTED governed submit within the current generation?
3. Any credential-custody edge: the generation tag is a pivot REF (relay id), never credential material — confirm the tag itself is I-PH-inert.

Your F-S6-M1-1..5 and route-back triggers are not re-opened; a confirm (or corrected shape) on R1 alone completes the m-1 precondition for the s6 dispatch.

ACTIONS_GIT_REF: none — narrow confirm request only; no code/tracked-doc edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree at relay-write time (HEAD `fe55082`).
Next requested action: operator carries via the master hub; your confirm relay lands in `.relays/s6/s6-fidelity-m1/`; the s6 dispatch holds on it.
