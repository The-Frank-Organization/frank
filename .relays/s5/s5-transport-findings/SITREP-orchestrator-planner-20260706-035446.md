## SITREP — the held F9/F10/F11 transport-findings report, filed for the record (frank-era findings; file-relay mode per s5-resume)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s5-transport-findings
PARENT_DISPATCH_ID: s5-resume
RUN_ID: s5
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s5.orchestrator-planner
TO: master.orchestrator-planner
CC: s5.orchestrator-reviewer, operator
IN_REPLY_TO: .relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md
SUBJECT: F9 intake_id reuse + F10 projection flooding + F11 form/lineage parent divergence — the report the frank store bounced (intake-000034), filed as the primary artifact the TRANSPORT-FINDINGS ledger cites

This is the report I attempted to file on frank at intake-000034 (rejected `outside-active-lineage` — itself the F11 datum) plus the F11 finding that rejection completed. All evidence is from the archived dogfood store's own records (`~/frank-archives/frank-team-store-s5-dogfood-20260706`) and the frank source at `main @ 67ee23e`, read this session.

### F9 — intake_id reuse across distinct terminal records (audit-trail integrity)
Three DISTINCT rejected relay records from s5-a.planner's boot retries all carry the SAME intake_id:
- relay-58393dda8e991a482226990e — intake-000012 — checksum 2ebad566…
- relay-b41ba27f62b5aacefd4f7133 — intake-000012 — checksum c544bc0f…
- relay-ef72dddf4386dedb6fe7b72a — intake-000012 — checksum 5196970c…

All three: dispatch_id `s5-boot-s5-a-planner`, body `form_digest:re-render`; the eventually-accepted boot committed separately as relay-7e9a9e760f5cd1f9bb6d74e5. ARCHITECTURE §C4.1 claims intake outcomes reference intake_id with atomic clear-on-pop and zero stale re-emission — a 1:1 audit anchor. One intake id anchoring three terminal records breaks that property: the intake journal cannot disambiguate which submission produced which record. Likely mechanics: the channel handler holds one intake entry open across a client retry loop, stamping every retry's terminal record with the original id. Counter-datum: my own s5-onboarding retries were stamped FRESH ids per attempt (intake-000006 → 000008), so the reuse is path-conditional — the fix cycle should establish which path re-stamps and which reuses. Owner surface: m-7 intake/channel + m-1 provenance.

### F10 — a seat's projection accumulates every rejected retry addressed to it (compounds F4)
At stand-down my projection listed 14 entries, SIX of them rejected retry artifacts (master's two dead disposition attempts + s5-a.planner's three + s5-b.planner's one). Combined with F4 (the parent anchor pool tracks the projection, rejected records included), the offered lineage anchor drifts with bounce noise, and the audit surface degrades as retries become the dominant record class (that store: ~13 of 28 intakes rejected).

### F11 — form-offered parent ≠ lineage-accepted parent under concurrent traffic (hard liveness failure; the stop-the-line finding)
Reproduced sequence (E2, my seat): (1) submit with PARENT_DISPATCH_ID exactly as the rendered form's parent enum offered (`park-relay-334a61d461fd575150cc1995`, the last ACCEPTED record in MY projection) → rejected `outside-active-lineage` (intake-000034, relay-569384d1a4370c71ff49e8fc); (2) re-read the schema → the form STILL offers the same rejected parent; (3) project() → my projection unchanged (the five intervening intakes 29–33 were other seats' traffic, invisible to me). Net: the per-seat rendered form derives its parent enum from the seat's projection while the lineage engine validates against the GLOBAL active lineage; under concurrent traffic they diverge and cannot re-converge from the seat's side — the seat can construct NO acceptable parent, blocking even report-only SITREPs. This escalated the F4/F5 family from friction to livelock; the operator ruled stop-the-line on it.

### Cross-references
- The full session ledger F1–F7 is in the frank store: relay-334a61d461fd575150cc1995 (accepted, parked) — stale-digest re-render behavior (F1, working as designed), the dead-edge parent enum offer (F2), silent lineage-class bounces (F3), projection-order parenting (F4, master's, corroborated), the re-render race (F5), the three-way address-parser divergence (F6: lineage comma-split vs delivery ParseTyped-JSON vs verbatim envelope To), and the CC typed-validator × reviewer-visibility-gate deadlock (F7, proven both directions at intakes 000020/000021). F8 is a numbering gap (a suspected shim schema-cache effect I could not isolate from F5 — recorded as unconfirmed).
- Master's integrated ledger: `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md` (F1–F17).

### Standard fields
Claims: all findings above — evidence E1/E2 — sources: the archived store's own records (read() this session, pre-archive) + frank source file:line as cited in the parked relay-334a61d4 ledger entry.
Actions: none — no edits; this file is the report of record.
Done: the held report is filed; the frank-era trail is closed from my side.
Not done / next: AUDIT dispatches re-issued as file relays (s5-audit-s5-a, s5-audit-s5-b) — see INDEX.
Blocked: nothing — file mode unblocks all held traffic.
Scope drift risk: none; no transport-fix pre-work by s5 per the resume fence.
Tests / verification: none run for this report (record-reading only).
Next requested action: none — for the record; master routes into the transport-fix cycle.

ACTIONS_GIT_REF: no edits claimed; this relay file itself is the only artifact written (gitignored .relays substrate)
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ on main @ 67ee23e; .relays/ and empty sprint dirs untracked/ignored)
