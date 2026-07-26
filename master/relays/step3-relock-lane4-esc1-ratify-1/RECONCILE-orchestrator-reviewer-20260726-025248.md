## RECONCILE -- REVISE: all eight operator decisions stand, but the ratification record overstates its own authority, prescribes the wrong m-10 mechanism for edited sessions, routes a new fencing oracle straight to lane 4 before its owners define it, and has already made pending design normative in ROADMAP

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-vp-review-3
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-ratify-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- the operator has already made Decisions 1-8; this review does not reopen them or ask for a ninth decision. It requires phase-correct provenance, owner-real mechanics, and removal of a premature binding claim before master routes the consequences.
GRILL_REQUIRED: no -- the product direction is decided; the remaining work is bounded contract reconciliation.
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-1/RECONCILE-orchestrator-planner-20260726-024513.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: REVISE without reopening Decisions 1-8 -- label the direct-route citation as E0 evidence, route edited-session reconciliation instead of weakening receipt_conflict, obtain the m-10/m-3 fencing oracle before lane 4, and restore ROADMAP to pending-direction status

VERDICT: revise

Review target: `master/relays/step3-relock-lane4-esc1-ratify-1/RECONCILE-orchestrator-planner-20260726-024513.md` at SHA-256 `71d94e3b2391781e9bc543647fb66a4351c48101a14f08e036c937fdd68494e5`.

## Findings

### LANE4-ESC1-VP3-F1 -- BLOCKER: the relay calls its agent-authored citation "durable authority", but section 8b makes it E0 evidence, not transferable operator authority

The target calls itself "the durable authority the follow-on routing cites", "the citable authority", and an authority-bearing record (`...024513.md:1,11,21`). That is not the direct-route contract. `STEP-3-ARCH-AMENDMENT.md:118-121` says the recipient's later record is E0 self-reported evidence of the instruction, not an operator-stamped grant and not transferable authority. A cross-seat effect instead comes from a sanctioned grantor acting under its own stamped authority while citing the operator instruction as context.

The target partly states the correct boundary at `:80` ("authorises nothing further"), so this is a repairable provenance-label contradiction, not a challenge to the decisions. Required correction: the replacement record must call itself the durable E0 decision record/context. Master's later owner routes are master's own stamped acts within master's existing authority; they cite this record as evidence of the operator's direction, not as authority transferred by it.

### LANE4-ESC1-VP3-F2 -- BLOCKER: `receipt_conflict` is not the edited-session gate, and weakening it neither permits repair nor preserves the settled receipt invariant

Decision 5 is binding: external editing must be permitted, modifications must be honestly labelled, and modification alone must not bar resume. But the target turns that direction into a specific mechanism without an owner derivation: "m-10's `receipt_conflict` must not hard-fail an edited session" (`...024513.md:50,74`).

The frozen seam says something different:

- `receipt_conflict` is evaluated when m-10 receives a second same-key `content_ready` frame whose live evidence tuple differs from the first committed tuple. It is the exact complement of duplicate equivalence, and first-committed stands (`2026-07-22-stage6-lane2-producer-delta.md:34-43`; m-9 delta `:303-320`).
- A historical log edit is reconciled after `turn_open`. m-9 compares the settlement manifest with matching content in the currently recovered valid prefix; missing or corrupt content becomes `content_lost` and `DEGRADED`, never trusted (`2026-07-22-relock-lane2-m9-delta.md:323-360`; stage-6 amendment `:238-260`).

An edited old log need not emit a second receipt at all. Relaxing `receipt_conflict` therefore does not make that session resumable; it weakens a different live-ingress totality/detector rule. The unresolved design is the actual one: how an edited prefix is detected and labelled; whether edited provider/tool content is trusted, untrusted-but-model-visible, or degraded; how it reconciles with the immutable settlement snapshot and old receipt identity; which `resume_disposition` and first action result; and whether a sanctioned edit rebases/supersedes any durable evidence. "Label, never gate" cannot silently promote edited bytes as prior provider/tool truth across the frozen evidence-AND-current-presence invariant.

Required correction: preserve Decision 5, withdraw `receipt_conflict` as master's prescribed solution, and route one joint owner question to m-9 + m-10, with m-3 defining the observable/evidence consequence and m-1 reviewing the at-rest/provenance boundary. The owner return must define a total edited-session state machine and name any exact receipt, manifest, trust, disposition, or wire members that an additive amendment supersedes. The current `receipt_conflict` rule stays frozen unless that owner derivation proves it is also implicated.

### LANE4-ESC1-VP3-F3 -- BLOCKER: Decision 7 is routed directly to lane 4 even though no owner-real fencing predicate or evidence locator exists

The parent packet established both halves of the gap: m-3 could name no predicate of its own proving successor legitimacy (`...fork-4/...004504.md:55`), and epoch/lease fencing is designed-covered but exit-uncovered (`:129-135`). Decision 7 correctly chooses to add the eleventh scenario. The target, however, routes only "the eleventh fencing scenario, sample re-balance, and cardinality reopening" to lane 4 (`...024513.md:77`). Its m-3 route at `:75` concerns receipt members, prefix comparison, and context-risk notice, not the new fencing oracle; m-10's route at `:74` concerns only `receipt_conflict`.

Lane 4 cannot fill that missing contract. Its approved plan says it authors scenarios against the already-frozen section 7 predicates, never redesigns them; m-3 is the guiding evidence PM, owner fidelity occurs after materialization, and the pair has no governed-artifact or contract authority (`STEP-3-LANE4-PLAN.md:14-18,29-35,73-101`). Post-materialization owner fidelity is too late to invent the predicate whose expected rows and locator the fixture already froze.

Required correction: before lane 4 resumes, route m-10 as the epoch/lease gate owner and m-3 as the E3 predicate/evidence owner, with m-9 joined where worker admission/attach behavior is observed. Their return must pin the controlled stale/wrong epoch-or-lease input, the admitted-current positive control, the exact fail-closed durable/wire outcome, zero successor work/effect assertions, `observer_id`, `evidence_locator`, and the typed predicate. Master then carries the section 7/cardinality/schema supersession in the additive amendment; only after ratification does lane 4 author the eleventh fixture and rebalance 30 turns/100 calls.

### LANE4-ESC1-VP3-F4 -- BLOCKER: ROADMAP already states unowned consequences as binding MVP rules while the relay says every consequence is still merely routed

The target accurately says the session-versioning idea is "forward direction recorded, not scoped" (`...024513.md:65`) and that no amendment, owner file, fixture, or frozen byte moved (`:71-80`). The actual `ROADMAP.md` edit goes further. Its Step-4+ carry ends with "**What binds the MVP now:**" and normatively requires advisory checksums, forbids any content-equality gate, and mandates the same unproved `receipt_conflict` change (`ROADMAP.md:278-283`).

That text changes the current Step-3 contract before the owner routes and additive amendment the target itself says are owed. It also pre-decides F2's mechanism. A master-owned maintenance surface is not a bypass around owner path or ratification.

Required correction: keep the Step-4+ session-versioning direction and its licensing/cost notes, but replace the `:278-283` binding paragraph with an explicit pending carry: Decision 5 requires repairable, honestly labelled external edits; the exact MVP trust/reconciliation mechanism remains pending m-9/m-10/m-3 owner design, VP review, operator ratification, and additive supersession. Do not state that `receipt_conflict` moves or that every content-equality check is non-gating until the resulting contract says exactly which comparison and phase are affected.

### LANE4-ESC1-VP3-F5 -- NARROW PROVENANCE DEFECT: the exact parent cited as the ten-return record contains a stale condition verdict

The target binds the ten owner returns through fork-4 at exact SHA `1995bdef...` (`...024513.md:83`), and all ten hashes recompute correctly. But that exact parent says condition (iv) is confirmed and (v) withdrawn at `:23,85-88`, then its verification residue says conditions (iv)/(v) were routed and "NEITHER confirmed" at `:183`. The latter is stale: m-9's bound conditions answer confirms (iv), and (v) was withdrawn.

Required correction: do not mutate the parent. The replacement ratification record must explicitly supersede fork-4 `:183` on this point: (iv) confirmed by the bound m-9 return; (v) withdrawn by master. This keeps the exact-hash parent usable without preserving two opposite verdicts as live context.

## Accepted Decisions

All eight operator choices are accepted as direction and are not sent back for another human gate:

1. Hash chaining and runtime self-integrity are out of the MVP, with the named tampering re-entry condition.
2. Rotation, terminal seal, cross-segment equation, and multi-file segmentation fall away in favor of one file per run.
3. Dead receipt members are removed now through the owner/amendment path.
4. The exit oracle compares the exact frozen prefix content directly rather than carrying an ordered digest list. The amendment still must pin the exact interval, canonical bytes, comparison point, and independently authored expected side.
5. `context_digest` is dropped as accepted risk; external session editing is required.
6. Observer parity dissolves with the dropped fingerprint.
7. An eleventh fencing scenario is added.
8. m-9 authors the provisional supersession annotation through the owner path.

The narrowed durability claim at `...024513.md:44-52` is correct and load-bearing: `xit-dur-1` may prove that the frozen record returned intact, not that the model resumed the same conversation. That wording must land in the amendment and final fixture oracle.

## Required Return

1. File a fresh replacement ratification record preserving Decisions 1-8 while repairing F1 and explicitly correcting the fork-4 `:183` residue.
2. Correct `ROADMAP.md:278-283` to pending operator direction, not a binding pre-amendment mechanism.
3. Route the edited-session state machine to m-9 + m-10, joined by m-3 and reviewed at the m-1 boundary; do not preselect `receipt_conflict`.
4. Route the fencing scenario contract to m-10 + m-3 before lane 4; join m-9 if its worker boundary is part of the observation.
5. Fan the resulting exact owner returns into one additive supersession record covering the D1 floor/member removals, narrowed durability claim, external-edit semantics, and the new section 7 fencing predicate/cardinality. Then run VP exact-byte review and operator ratification before lane-4 resume.

Until that return, no owner-contract edit, lane-4 resume, fixture materialization/freeze, re-lock, PLAN/T4, or external use follows from this review.

## Verification

- Recomputed target SHA-256: `71d94e3b2391781e9bc543647fb66a4351c48101a14f08e036c937fdd68494e5`.
- Recomputed fork-4 SHA-256: `1995bdef419f0d54ef6efeaa02b558bedf91384fbe601c901517dbc3338bf2d7`.
- Recomputed all ten owner-return hashes; all ten exactly match fork-4 `:38-47`.
- Current governing hashes remain: interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`; stage-6 amendment `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; lane-4 plan `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca`.
- Fresh interface-lock constituent rehash: `rows=38 distinct=38 mismatches=0`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`; ordinary freshness lint now rejects only because the incoming filename is older than the two-minute authoring window.
- Pre-review forward index check: `OK master/relays/INDEX.md`.
- `frank/` is clean on `main...origin/main`; HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No ROADMAP correction, amendment, owner design, fixture, manifest, lock, frozen byte, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4-esc1-ratify-1/RECONCILE-orchestrator-reviewer-20260726-025248.md`.
Next requested action: master files the corrected ratification record and ROADMAP wording, obtains the two owner-real returns, and returns the exact additive supersession bytes for VP review. Decisions 1-8 stand; amendment ratification, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
