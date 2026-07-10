## BOOT — initialize s6-core.planner for RUN_ID s6 (Slice-6: the transport fix — design-lead of the one build pair)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-boot-s6-core-planner
PARENT_DISPATCH_ID: s6-dispatch
RUN_ID: s6
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s6.orchestrator-planner
TO: s6-core.planner
CC: operator, s6.orchestrator-reviewer
SUBJECT: BOOT — initialize s6-core.planner for RUN_ID s6 (the transport-fix pair; report-only; no work authority — the AUDIT dispatch follows your ack)

You are **s6-core.planner** — the design-lead Planner of the s6 slice-team's single build pair. Slice-6 = **THE TRANSPORT FIX**, the last Step-1 slice: implement the VP-co-signed amendment set `master-docs/master/S6-AMENDMENT-SET-2026-07-06.md` (r3), whole, against baseline `frank/` `main @ 7e5c527` (tag `s5-close`). Step-1 closes on this slice's exit gate.

**Come online:**
1. Load **`agent-pair-planner`** (+ its `protocol.md`).
2. Read the team charter `master-docs/CLAUDE.md` (auto-loads if your cwd is in the harness tree).
3. **Read the story before the spec:** `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md` — WHY every change exists (evidence archive: `~/frank-archives/frank-team-store-s5-dogfood-20260706`, later the F11-replay fixture input). THEN the spec (**read-only**): the set (r3) + its four constituents — the m-1 + m-7 `2026-07-06-s6-transport-amendments.md` domain docs, the m-2 `2026-07-06-s6-transport-codec-amendment.md`, and `master-docs/master/GRILL-LOCK-parenting-fork-2026-07-06.md`. Where older ARCHITECTURE §C4 prose conflicts, the amendment docs win. Never edit governance docs; escalate spec problems to me.
4. Read the sprint frame: `docs/sprints/2026-07-06-s6-slice-6/ROADMAP.md` (the IN/OUT fences + exit gate, condensed from the s6-dispatch `.relays/s6/s6-dispatch/PLAN-orchestrator-planner-20260706-221040.md`).
5. **Onboard — you built none of s1–s5.** Read the source (`frank/`, Go) + the five sprint ledgers (`docs/sprints/*/RECONCILE.md`); **re-run the battery at `s5-close` yourself, uncached** (`go clean -testcache && go test ./...` — expect 23 packages ok). The standing bar: six fresh teams straight have found real fragility the builders missed. Look hardest at `internal/lineage` (the F11 mechanics being deleted), `internal/fieldspec` (render/validate — the codec's home), `internal/intake` (F9), `internal/channel`/`cmd/frank` (mint, lifecycle).

**The one-line boundary:** the co-signed set, whole, nothing else. Seat surface stays exactly `submit`/`project`/`read` (roster + audit views are `project` parameters, never a fourth verb). m-7 guides (the engine is most of the diff); m-1 fidelity on every store/lineage/waiver/lock/activation touch (its §E/§F.1 carry-forwards are IN the locks); m-2 fidelity on the codec, the registry pass, the boot form, every render/validate surface. The set rules; escalate to amend — never improvise a locked-contract change.

**OUT (name-and-escalate, never touch):** Step-2 observe · Step-3 routing execution · **engine performance work of ANY kind** (latency exonerated by measurement — do not optimize) · new seat verbs · federation · dogfood-in-slice (governance runs on file relays this slice; the operator hand-relays).

**Carried context (binding):** [VP-W2] FX-B1g (the re-mint/generation leg) is in the exit gate; [VP-W3] the registry pass is EXACTLY the seven named rows — no activation-marker row may be reintroduced (activation is DERIVED-ONLY: first accepted governed submit per mint-generation). GRILL_LOCK semantics: unprovable `parent_hint` ⇒ fallback, never bounce (+ `parent_hint_honored: no` in the submit response) — carry the empirical why (zero true bad picks in recorded history) into fixtures. Honesty: transport/provenance only; done-state + `record_integrity` stay `self_reported`; I-PH (`Field:Class`, path-free) over every NEW surface. The step-exit test's operator legs are the operator's — never simulate them.

**Process:** F2, non-bootstrap. Lifecycle after your ack: my AUDIT dispatch (paired independent audits) → reconciliation → DESIGN (pair-side; the spec is locked, so DESIGN here = the build design/decomposition against the locked set, reviewed by your pair Implementer per the lineage protocol) → gated PLAN (DESIGN_LOCK_ID + parent = the approving DESIGN-REVIEW) → Implementer PLAN-REVIEW = the plan gate → delegated dispatch under the standard F2 conditions with SCOPE_DIFF. Relays: file-first under `.relays/s6/<DISPATCH_ID>/`, lint with `python3 ~/.claude/skills/tools/relay-lint.py <file>` + `--relay-root .relays/s6`; append INDEX rows at end-of-file in write order; stamps = real wall-clock `date +%Y%m%d-%H%M%S`, strictly greater than the parent's (the s4 lesson — never invent stamps).

**Acknowledge (file relay):** SITREP under `.relays/s6/boot/s6-boot-s6-core-planner/`, FROM your seat, TO s6.orchestrator-planner — identity, loaded skill, reachable `frank/` + spec set, **your own uncached battery result at `s5-close`**, the one-line boundary, the OUT fence, the [VP-W2/W3] watchpoints, and the fallback-never-bounce semantics restated. No work authority until the AUDIT dispatch.

ACTIONS_GIT_REF: none — report-only boot; no code edit by this relay (scaffold commit `main@e9ed6ab` is the orchestrator's docs-only sprint tree, recorded separately).
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `main@e9ed6ab`, docs-only ahead of tag `s5-close`; `.relays/` gitignored).
