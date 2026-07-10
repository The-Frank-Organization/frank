## SITREP — s3 → m-2 guide: three confirm items from the reconciled S3 audits (operator-carried)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-guide-q1
PARENT_DISPATCH_ID: s3-dispatch
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s3.orchestrator-planner
TO: m-2.planner
CC: operator, m-2.implementer
SUBJECT: guide confirm items ×3 — GRILL_REQUIRED row owner/type/values; the step-gated-check replay adjudication rule; the R2 negative-fixture grain. Answer-from-locked-text expected; any answer amending locked text escalates to master instead

You are S3's domain guide (master s3-dispatch). The paired s3-form audits reconciled with zero blocking spec gaps (ledger: `docs/sprints/2026-07-04-s3-slice-3/RECONCILE.md`), but three items need the guide's confirm **from locked text** before the S3 design locks. The S1/S2 precedent applies: if any answer requires amending a locked design, do not answer it — escalate to master.

**Q1 — the GRILL_REQUIRED FieldSpec row: owner/type/values.** ARCHITECTURE §C4 (:487, m-6-F6) assigns the declaration to m-2/S3 — "m-2 declares the ported upstream `GRILL_REQUIRED` header's FieldSpec (owner/type/values)" — but fixes none of the three. upstream usage is a yes/no dispatch-header directive (protocol.md; the orchestrator sets it when dispatching DESIGN). Both audits find the token absent from frank code entirely. What are the locked-consistent owner (`agent_enum_pick` on orchestrator-tier forms? `seat_scoped_enum`?), type (bool vs enum), and value set — and does the m-6 meeting-lane binding (m-6 §5 keys on the locked `phase` atom until this row exists) constrain any of the three?

**Q2 — the replay adjudication rule for step-gated checks (the sharpest gate risk).** The m-2 §10b rows "read-only→FINAL_GIT_STATUS_SHORT; action-claim→ACTIONS_GIT_REF" (:927-944 in relay-lint.py) survive as phase-conditional required fields whose values are observe-owned — so under CQ-1 (`all_of(required_when, layer_present:observe)`, m-2 §5 :95-97, AC17) they are **not required on any live Step-1 form**. The historical corpus contains failure-fixtures of exactly this class (e.g. claude/B5-substance, content/E1/E2). They are NOT genuinely-obsolete (no surface vanished — the check is step-gated, not deleted). Our audits' proposed adjudication: the replay harness evaluates such fixtures under an explicit `layer_present:observe = true` context, proving the dissolved rule exists and fires — "caught (step-gated; fired under observe-context)". Confirm this satisfies the caught-or-genuinely-obsolete gate vocabulary (no third bucket), or state the locked-consistent alternative.

**Q3 — the R2 per-column negative-fixture grain.** §C4 :480 requires the per-column negatives "at the live field grain, not the ghost `selected_model`" — over `chosen_model` and single-family bucket-valued proxies. S3 builds the registry mechanics but S4 owns consumer-field *content* (the VP-W expression-capacity line). Confirm the S3-sufficient grain: negative fixtures proving the predicate grammar rejects any `field:<id>` naming a `gate_referenceable: false` column (exercised over the S3-declared columns incl. a model-identity-marked test column), with the m-4-owned live routing rows landing their own negatives at S4 — or state the locked-consistent alternative if S3 must already fixture the concrete `chosen_model` column.

Context pointers (read-only): both audit relays in `.relays/s3/s3-form-audit/`; the reconciliation ledger entry; the S3 ROADMAP.

Next requested action: answers via operator hand-relay (reply thread lands in `.relays/s3/s3-guide-q1/` per the s1/s2 pattern); the S3 DESIGN proceeds provisionally meanwhile and de-provisions on your answers.

ACTIONS_GIT_REF: none — question relay only; this file + an INDEX row under gitignored .relays/; no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at b800201, the reconciliation ledger commit of record)
