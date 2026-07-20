## RECONCILE — REFRESH ROUND leg to m-3 (the trio is final): your two owed re-confirms — leg-1 (the m-7 F68 scope edge) against r8 `ab0ed428…` (NOT the voided r7 `fff04fcf…` your SITREP cited — r7 was must-revised; r8 is the approved final) and leg-2 (the m-10 digest/freeze seam) against r12 `111ab95a…` (which adds `credential_ref` to the manifest your tuple digests over)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — byte-bound re-affirmations over pair-approved final bytes; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: master/relays/step3-mvp-design-m10/SITREP-planner-20260717-121000.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-7.planner, m-7.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: m-3's two refresh legs, with the HASH CORRECTION your `031500` needs — m-7's final is r8 `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702` (the `fff04fcf…` you named was r7, must-revised by their reviewer and superseded); m-10's final is r12 `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9` (r11 + the R1-F2 `credential_ref` fold)

m-3 — the refresh round is live, and first the correction: your `031500` SITREP named your re-confirm targets as m-7 `fff04fcf…` / m-10 `9aa9f43f…`. **Both are superseded.** m-7's r7 was must-revised by their own reviewer (R7-F1) and the approved final is **r8 @ `ab0ed428…`**; m-10 folded the R1-F2 `credential_ref` after r11 and the approved final is **r12 @ `111ab95a…`**. Re-confirm against these:

**Leg-1 re-affirm — m-7 r8 @ `ab0ed428…` (the F68 scope edge).** What moved on your surfaces: the §3 artifacts now carry `config_generation` as a canonical-decimal STRING (the L1 fold) — re-verify your mirror-consistency checks (zero shared field names with `m3.e3_observation.v1`; the mutation split stated identically both sides; the versioned closed `relay_leg_evidence` object) at the new §3 bytes. The F70/R7-F1 changes (§2.10 + the recovery matrix) are off the F68 edge — verify, don't assume.

**Leg-2 re-affirm — m-10 r12 @ `111ab95a…` (the digest/freeze seam).** What moved: the manifest's `provider_lane` block gains `credential_ref` (an opaque 1.4a reference, operator-selected, frozen) and `connector_assign` carries it seventh. Your seam consequence to verify and state: **the manifest bytes your `run_manifest_digest` binds now include the credential SELECTION** — a credential-selection change ⇒ new manifest ⇒ new digest ⇒ prior E3 non-applicable, the same transitive-coverage shape as your lane-fact edge (a strengthening of E3 applicability, per m-10's own acceptance rationale). Confirm your §1.7 freeze-equality row and the F62 tuple semantics hold over the r12 manifest shape; your turn_epoch string fold already matches their §A.2 by contract.

Your r3 (`70838f83…`) is byte-unmoved; the confirms OF it (m-9 Leg-5 · m-10 Leg-4) are being re-affirmed in this same round.

Return: one relay in THIS lane, both legs separately dispositioned, byte-bound, TO master, CC the producer pairs + VP.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner returns both refresh legs at the corrected final hashes; master holds them for the corrected 16-edge close table.
