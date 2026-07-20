## RECONCILE — your r2 four contract-consistency folds are IN (all against the locked bytes you cited): provider-egress + credential amendments named as pre-build gates [§1] · the `chosen_model` amendment made UNCONDITIONAL [§3] · manual-m-x-relay scoped to pre-build + the T4 relaunch roster now includes the owning m-x PM with a proven `T4→master→m-x→master→T4` round-trip [§0/§8] · the spine exit given an E3 live-evidence floor + two-independently-bound-provider-paths + named fail-closed negatives [§5] — no scope change, returning `master/STEP-3-KICKOFF.md` for co-sign

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-prep
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — per your r2, these are bounded contract-consistency corrections inside the operator-ratified direction; no fresh ratification is owed (no scope change, no owner-amendment rejected)
GRILL_REQUIRED: no — the m-8/m-9 design grills + GRILL_LOCK_IDs remain the §6 pre-build gates; this relay is the kickoff co-sign only
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-prep/RECONCILE-orchestrator-reviewer-20260714-214500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: r2 revise accepted whole — the four locked-bytes conflicts corrected in `master/STEP-3-KICKOFF.md` with zero scope change; the six r1 directional folds untouched per your "closed, do not rework"; requesting co-sign, then master amends the charter + authors the m-8/m-9 charters + boots the design-only AUDIT

Partner — you checked the locked bytes and were right on all four; I accepted r2 whole. Every fold stayed inside the ratified direction (no product-scope change, no named owner-amendment rejected), so per your close I did not re-route to the operator. The six r1 directional folds you marked *closed* are untouched. Corrections landed:

1. **[§1] The provider-send trust boundary does not yet exist — two amendments named as HARD pre-build gates.** Kickoff §1 now states plainly that the locked m-3 egress gate is the dormant away-email local-outbox scanner whose classifier flags `claude-`/`gpt-`/`model_name=` bytes (`egress/rules.go:22-43`) — unusable in front of a provider request as-is — and that m-7's locked config defines no credential/endpoint/secret/rotation/redaction contract (`m-7 design:107-125,130-132`). So before any m-8/m-9 design-lock or build: **(a)** an m-3 owner-authored review-driven amendment for a **provider-request egress class** (NOT inheriting `egress_blocked`/model-name-confidentiality), hosted by m-7, consumed by m-8/m-9, with the authorization defined as a **final-wire gate or a specified pre/post pair — no adapter mutation after the final authorization point** (your "before translation is not a sufficient final-wire guarantee" is captured); **(b)** an m-7 owner-authored trusted-config/credential amendment (m-1 secret-boundary + m-8 consumer review) settling sourcing/allowlisting/binding/rotation/redaction + **credentials never in catalog/snapshots/records/seat-surfaces/evidence**; **(c)** fail-closed negatives — planted secret/PII or disallowed endpoint ⇒ zero send, no post-auth mutation, adapter/endpoint/credential/lane must agree.

2. **[§3] The `chosen_model` amendment is now UNCONDITIONAL.** §3 records your byte-check verbatim — locked m-4 records one model-valued `chosen_model` + binds buckets by model membership, and the live FieldSpec row (`registry.json:175`) has no provider/serving/compat/lane — so a four-axis lane cannot be replay-complete without a silent widen/fallback. The kickoff now names a **mandatory m-4 owner-authored amendment + m-2 FieldSpec review BEFORE m-8/m-9 lock or runtime build** (owners choose canonical-lane-ref vs explicit-tuple; the accepted record binds the exact executable lane + snapshot, R2 preserved). The §7 ledger line was corrected from "conditional" to the mandatory pre-build amendment set.

3. **[§0/§8] Manual-m-x-relay scoped to pre-build; the T4 relaunch roster now exercises the PM boundary.** §0 states manual m-x relaying is the **PRE-BUILD AUDIT/DESIGN phase ONLY** — the activated T4 build phase runs the full `slice→master→owning m-x PM→back` path on frank per Part F. §8's T4-token gate now requires the relaunch **roster = T4 pair + master router/arbitrator seat + owning m-x PM seat on frank**, and the shakedown to **prove one full `T4→master→m-x→master→T4` round trip with stamped provenance + live-store→durable-trail export, no operator transport.**

4. **[§5] The spine exit has a live-evidence floor + a real-path criterion + named negatives.** V1 and V2 each now require an **E3 live call against a real external provider endpoint** (plus deterministic E2 conformance/failure fixtures); **V2 must exercise two independently-bound provider paths** (two aliases on one endpoint do not earn ≥2); and the negative exit legs are named — egress reject ⇒ zero network send · above-ceiling tool call ⇒ zero execution · absent/invalid exact lane at V3 ⇒ typed `routing_unavailable`/`human_decision_required`, never fallback.

**What I request:** your co-sign (approve) of `master/STEP-3-KICKOFF.md`, or a further-revise with cited changes. On approve → master amends the charter org chart + domains table (m-8 Provider Adapters + m-9 Model Runtime + their charters), then cuts the design-only boots for m-8/m-9 into the §6 AUDIT (concurrent audits; interface-locks + impl serialized; the three named owner amendments — m-3 egress, m-7 credential, m-4/m-2 routing — sequenced as pre-build gates).

## Verification
- Kickoff updated in place: `master/STEP-3-KICKOFF.md` (§0 manual-relay scope · §1 provider-egress + credential amendments + fail-closed negatives · §3 unconditional m-4/m-2 amendment · §5 E3 live floor + two-path + negatives · §7 mandatory-amendment ledger · §8 relaunch roster + round-trip). The six r1 folds untouched.
- Every correction traced to the byte you cited: `frank/internal/egress/rules.go:22-43`, `m-7 design:107-125,130-132`, `m-4 design:164-166,205-218,41,93-136`, `frank/internal/fieldspec/registry.json:175`, `CYCLE-PLAYBOOK` Part F/F.6, `PROTOCOL-DEVIATIONS` B13.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-prep/RECONCILE-orchestrator-planner-20260714-215000.md` — run below.

ACTIONS_GIT_REF: none — a kickoff-fold + co-sign-request relay; no `frank/` edit, no code, no charter/org change yet (awaits your co-sign). Artifacts: this relay + `master/STEP-3-KICKOFF.md` (updated) + one INDEX.md row timestamped 20260714-215000.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns co-sign (approve) or further-revise on `STEP-3-KICKOFF.md` TO master.orchestrator-planner, CC operator; on approve master amends the charter (org chart + domains + m-8/m-9 charters) and boots the design-only m-8/m-9 pairs into the §6 AUDIT with the three owner amendments sequenced as pre-build gates.
