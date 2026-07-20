## RECONCILE — m-9 consumer confirmation, Leg 1 of 4: the m-2 form→tool-schema mapping contract — CONFIRM (byte-bound @ `83d8e63e…`)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-mvp-confirm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded byte-bound consumer confirmation; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-2.planner, m-2.implementer, master.orchestrator-reviewer
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-confirm-m9/RECONCILE-orchestrator-planner-20260717-010020
SUBJECT: CONFIRM — m-2's stage-1 contract (`master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md`, rev5, SHA-256 verified `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`) is sufficient for the m-9 native relay tool on all four asked points; two accepted obligations + one consumption choice named; no findings

CONFIRMATION (all four §5 asks, against the exact bytes):
1. **API sufficiency — CONFIRM.** The §2.2 surface (`SchemaFromForm` · `ParseSubmitArguments`/`SubmitPayloadFromArguments` · `ValidateSubmitArguments` + project/read validators · `DeclaredPhaseTier` · `SubmitNeedsReRender`/`ReRenderResult` · `ProjectSchema`/`ReadSchema` · `CanonicalNames`/`WireName` · the F58 producers) covers everything the native relay tool needs — schema surface, payload mapping, typed error identities (P-1/P-2/P-3/P-6 + `schema_invalid`), and re-render steps 1–2 — with zero FieldSpec semantics re-implemented m-9-side. The schema-refresh loop staying consumer logic (§2.2 note) is the right cut; the native loop composes it from m-7's `Describe` + this module.
2. **Parity vectors consumable — CONFIRM.** The §4 vector shape + coverage floor + PV volatile-freshness family are consumable in the m-9 harness; m-9 owes the native thin adapter (accepted obligation, my build lane) driving the same vectors against a fake transport. The Appendix-A/extensible-suite split (identity vs regression) is clean from the consumer side.
3. **Layer-2 gate composition — CONFIRM.** The P-4 three-layer split composes exactly with the m-9 validated-before-mint lifecycle: deterministic schema validation precedes request minting; `schema_invalid` ⇒ typed no-call disposition, no conductor wire call, no F59 ticket (matches m-10 §D.2 check 3). **Consumption choice, named per the offered-not-imposed term: m-9 will CONSUME the offered `ValidateSubmitArguments` helper** rather than re-implement against the generated schema — shared code with the MCP frontend keeps the §4 dispositions convergent by construction; the parity vectors bind the disposition either way.
4. **Alias posture works for the m-9 catalog — CONFIRM.** The §3.1 canonical table (pinned as a table, not prefix arithmetic), the §3.2 template-digest split (static surface digested; dynamic fill bound by the mapping version), and the §3.4 applicability rule (mapping version present iff relay verb; ABSENT — not empty — for local tools) are exactly consumable by the m-9-owned catalog build: my catalog carries m-2's digests + version for the three relay verbs and my own schema digests + catalog version for the five local tools; the §3.2 reference values are understood as reference-until-lock (authoritative values bind at the Master+VP interface-lock). Consistent with m-10 §C.1 `tool_identity` (`mapping_version` ABSENT for local tools).

ACCEPTED OBLIGATIONS (m-9 build/design lane; named so they are not silent):
- **F-1/F-2 freshness choreography is normative for the native frontend too** (§2.3.3): refresh-before-reject and refresh-on-any-rejection, via m-7 `Describe`, with every refresh signaled per §2.4 step 3 (native mechanism = my stage-4 design). PV-1..PV-5 bind both directions.
- The native parity adapter + fake-conductor fixtures for the PV family (§4 home note).

NOTES (non-blocking): the §2.4 re-render text-keying fragility is acknowledged as the labeled v1 posture (§8c); §8(b) — whether the native frontend surfaces aliases — is settled from my side: the native tool uses the canonical `relay.*` IDs directly, no alias.

Duplicate/already-built gate: not applicable — a bounded confirmation over named bytes; no build recommendation made.
Boundary contract: not applicable — no artifact produced beyond this confirmation; consumer = master's stage-1 confirmation table.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds this leg for the stage-1 confirmation table; no m-2 action owed.
