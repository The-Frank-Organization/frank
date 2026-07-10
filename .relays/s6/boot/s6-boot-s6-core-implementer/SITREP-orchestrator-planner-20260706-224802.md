## BOOT — initialize s6-core.implementer for RUN_ID s6 (Slice-6: the transport fix — adversarial reviewer/executor of the one build pair)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-boot-s6-core-implementer
PARENT_DISPATCH_ID: s6-dispatch
RUN_ID: s6
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s6.orchestrator-planner
TO: s6-core.implementer
CC: operator, s6.orchestrator-reviewer
SUBJECT: BOOT — initialize s6-core.implementer for RUN_ID s6 (the transport-fix pair; report-only; no work authority — the AUDIT dispatch follows your ack)

You are **s6-core.implementer** — the adversarial design-reviewer + eventual executor of the s6 slice-team's single build pair. Slice-6 = **THE TRANSPORT FIX**, the last Step-1 slice: implement the VP-co-signed amendment set `master-docs/master/S6-AMENDMENT-SET-2026-07-06.md` (r3), whole, against baseline `frank/` `main @ 7e5c527` (tag `s5-close`). Step-1 closes on this slice's exit gate.

**Come online:**
1. Load **`agent-pair-implementer`** (+ its `protocol.md`).
2. Read the team charter `master-docs/CLAUDE.md` (auto-loads if your cwd is in the harness tree).
3. **Read the story before the spec:** `master-docs/master/TRANSPORT-FINDINGS-2026-07-06.md` — WHY every change exists (evidence archive: `~/frank-archives/frank-team-store-s5-dogfood-20260706`, later the F11-replay fixture input). THEN the spec (**read-only**): the set (r3) + its four constituents — the m-1 + m-7 `2026-07-06-s6-transport-amendments.md` domain docs, the m-2 `2026-07-06-s6-transport-codec-amendment.md`, and `master-docs/master/GRILL-LOCK-parenting-fork-2026-07-06.md`. Where older ARCHITECTURE §C4 prose conflicts, the amendment docs win. Never edit governance docs; escalate spec problems through your pair Planner or to me.
4. Read the sprint frame: `docs/sprints/2026-07-06-s6-slice-6/ROADMAP.md` (IN/OUT fences + exit gate, condensed from the s6-dispatch).
5. **Onboard — you built none of s1–s5.** Read the source (`frank/`, Go) + the five sprint ledgers (`docs/sprints/*/RECONCILE.md`); **re-run the battery at `s5-close` yourself, uncached** (`go clean -testcache && go test ./...` — expect 23 packages ok). The standing bar: six fresh teams straight have found real fragility the builders missed. Look hardest at `internal/lineage` (the F11 mechanics being deleted), `internal/fieldspec` (render/validate — the codec's home), `internal/intake` (F9), `internal/channel`/`cmd/frank` (mint, lifecycle).

**Your seat's teeth:** your DESIGN-REVIEW is the pair design gate (lineage — your approve is what the gated PLAN parents); your PLAN-REVIEW is the F2 plan gate; you execute IMPL only on a live delegated `DISPATCH IMPL` addressed TO you with intact lineage + SCOPE_DIFF all-in. The pair record this cycle caught real defects at every station (B-1 classifier hole; m-2 r1 self-contradictions; m-1 §A sharpening) — grill for substance, not ceremony. Red-first fixtures are the slice's currency: every exit-gate fixture lands red before the production change.

**The one-line boundary:** the co-signed set, whole, nothing else. Seat surface stays exactly `submit`/`project`/`read` (roster + audit views are `project` parameters). m-7 guides; m-1 fidelity on store/lineage/waiver/lock/activation (§E/§F.1 carry-forwards are IN the locks); m-2 fidelity on codec/registry/boot-form/render/validate. The set rules; escalate to amend — never improvise a locked-contract change.

**OUT (name-and-escalate, never touch):** Step-2 observe · Step-3 routing execution · **engine performance work of ANY kind** (latency exonerated by measurement) · new seat verbs · federation · dogfood-in-slice (governance = file relays this slice).

**Carried context (binding):** [VP-W2] FX-B1g (re-mint ⇒ new generation starts `minted`; pre-re-mint accepteds do NOT activate; a fresh boot accept does) is in the exit gate; [VP-W3] the registry pass is EXACTLY the seven named rows — no activation-marker row; activation is DERIVED-ONLY (first accepted governed submit per mint-generation). GRILL_LOCK: unprovable `parent_hint` ⇒ fallback, never bounce (+ `parent_hint_honored: no`). Honesty: transport/provenance only; done-state + `record_integrity` stay `self_reported`; I-PH (`Field:Class`, path-free) over every NEW surface (roster, boot bounces, lock refusals, hint flags). The step-exit test's operator legs are the operator's — never simulate them.

**Relay discipline:** file-first under `.relays/s6/<DISPATCH_ID>/`; lint `python3 ~/.claude/skills/tools/relay-lint.py <file>` + `--relay-root .relays/s6`; INDEX rows append at end-of-file in write order; stamps = real wall-clock `date +%Y%m%d-%H%M%S`, strictly greater than the parent's (the s4 lesson).

**Acknowledge (file relay):** SITREP under `.relays/s6/boot/s6-boot-s6-core-implementer/`, FROM your seat, TO s6.orchestrator-planner — identity, loaded skill, reachable `frank/` + spec set, **your own uncached battery result at `s5-close`**, the one-line boundary, the OUT fence, the [VP-W2/W3] watchpoints, and the fallback-never-bounce semantics restated. No work authority until dispatched.

ACTIONS_GIT_REF: none — report-only boot; no code edit by this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `main@e9ed6ab`, docs-only ahead of tag `s5-close`; `.relays/` gitignored).
