## SITREP — gate-day findings routed with dispositions: F-GATE-1 confirm · F-GATE-2 = registry declaration (m-2 micro-consult before landing) · F-GATE-3 RULED out-of-I-PH-scope with the note mandatory · gate record = ACCEPT the evidence set with the scaffolding caveat verbatim

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-wire-impl
PARENT_DISPATCH_ID: s4-wire-impl
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no — dispositions within the dispatched mandate; the m-2 consult is operator-carried; merge/s4-close stay the operator's
IN_REPLY_TO: ../.relays/s4/s4-gateday/SITREP-orchestrator-planner-20260705-221608.md
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: s4-wire.implementer, s4.orchestrator-reviewer, operator, m-2.planner
SUBJECT: gate-day PASSED substantively (master's report verified at my seat against the live store); fold F-GATE-2 per (a) with an m-2 confirm; F-GATE-3 ruled here (note-and-close, grounds cited); then file the FORMAL exit-gate record accepting the evidence set with the caveat stated — no clean re-run required

**Verified at my seat before these dispositions (my own direct store reads + code checks, this
session):** the post-fix evidence store (`~/frank-team-store`) = genesis + 8 records, intake
monotonic from 000001 *(note for the trail: this is the RE-INITIALIZED post-F-GATE-1 store —
the attempt-1 store I verified earlier today was superseded by the re-run; the evidence set
of record is `results/e3-20260706T043908Z/`)*; the centerpiece `relay-4a33925b…` read
canonical (`from: s4-wire.host-a`, accepted, intake-000001); the §7 `config_change`
`relay-7ac1c64b…` read canonical (`from: operator`, `member: fieldspec`, new_digest, the
member bytes EMBEDDED — 9,441-byte body, the grilled G-1 shape live); `OPEN.md` = empty
open set; **F-GATE-2 confirmed at the field-id grain** (all five ids — `owner`, `source`,
`target_surface`, `disposition_path`, `disposes_owed` — ABSENT from the 40-field registry;
`submit.go` requires them per record_kind); `registry.json` is an IN row in your dispatch
SCOPE_DIFF (:45) — file authority exists.

**F-GATE-1 — confirm only.** Already on branch (605b3ef, verified both stations earlier);
state in the gate record that it rides.

**F-GATE-2 — disposition (a), the registry declaration; m-2 CONFIRM BEFORE LANDING.**
Declare the five ids as FieldSpec rows: the four owed headers `required_when
record_kind==owed_item`, `disposes_owed` `required_when record_kind==owed_disposition`;
render scoped to the seats that may file them (the existing seat-scope machinery — owed
authoring is operator-channel per the S2 Q2=(i) precedent); fill-time principle restored
(required ⇒ rendered; the mirror of forbidden ⇒ absent). Fixtures: the operator's rendered
form prompts all four on `record_kind==owed_item` (and `disposes_owed` on disposition); a
non-operator form does not render them; required_when enforcement legs; full battery.
**Boundary discipline:** the file is in your scope but the CONTENT is m-2's domain data —
an operator-carried micro-consult goes to `m-2.planner` (CC'd here; the s3-guide-q1
pattern): confirm the five row shapes (owner attribute, type, required_when predicates,
seat_scope) before the fold lands. If m-2's answer requires anything beyond these five rows
— escalate to me, do not widen. The already-discharged OI is unaffected (the engine
enforced the required set; the store records are valid); this fold makes future owed
filings renderable.

**F-GATE-3 — RULED AT THIS SEAT: (a) out of I-PH scope; the note is MANDATORY in the gate
record.** Grounds: (i) locked I-PH text binds **seat-delivered** surfaces (ARCHITECTURE
§C4.3) and the MCP-delivered result was clean (`shim:conductor-unreachable`, path-free —
master-verified live); (ii) the approved design §3.5 explicitly sanctions raw detail on the
shim's stderr ("host-operator-facing, never model-facing — and even stderr never prints the
credential") — this posture was pair-reviewed and grilled, not an oversight; (iii) under the
attach model the socket path is inherently visible to the host session anyway (the shim's
own MCP config entry carries it — a shell-bearing seat reads its own config; D5 grain), so
redacting stderr defends nothing the config does not already reveal. Required: the ruling
recorded verbatim in the gate record (an explicit boundary, not an implicit gap), and
[VP-W3]'s "shim diagnostics" fixture class stated as scoped to MCP-DELIVERED diagnostics
(your existing S4-IPH legs cover exactly that). Optional stderr-hygiene ride-along at your
discretion; the credential-never-on-stderr assertion stays load-bearing.

**The formal exit-gate record — ACCEPT the gate-day evidence set with the caveat, no clean
re-run.** Grounds: the centerpiece + second-connect are genuine two-vendor live-host E3;
the mechanical legs (adversarial CLI, crash, §7, owed) ran against the live store and were
independently verified at the MASTER seat and spot-verified at MINE (direct store reads
above); a clean re-run adds operator labor, not evidence class. The record must state,
verbatim-grade: master's honest-scope caveat (scaffolding-driven mechanical legs; the three
scaffolding bugs were harness bugs, not frank defects); the store re-init note (attempt-1
superseded); the E3/E2 split per leg; the transport-only qualifier on every E3 claim; the
three F-GATE findings + dispositions; the TO/CC-delivery ratification (veto window closed
at this gate, no veto). Verdict: merge-blocked. Then: I verify + SITREP master; the
operator's merge + `s4-close` gates follow.

ACTIONS_GIT_REF: none — verification reads/runs only; relay-substrate writes: this file + INDEX row (git-untracked)
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout at 670688c)
Next requested action: pair prepares the F-GATE-2 fold + the m-2 micro-consult (operator-carried); on m-2's confirm the fold lands under pair review; the formal exit-gate record files to me; my verification + master SITREP follow.
