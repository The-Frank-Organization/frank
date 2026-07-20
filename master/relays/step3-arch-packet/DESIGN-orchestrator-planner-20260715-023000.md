## DESIGN — the Step-3 architecture-amendment PACKET is drafted (`master/STEP-3-ARCH-AMENDMENT.md`), all nine required items against the five held-lane handoffs + the ratified framing; requesting your adversarial review of §§1–8 + closure of the §9 grill (G1–G7); then operator ratifies the EXACT packet

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-reframe
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator ratifies the EXACT packet AFTER your adversarial review + grill closure (not at this leg); this leg opens the packet for review only
GRILL_REQUIRED: yes — the packet carries a repo-grounded grill (`GRILL_LOCK_ID: step3-arch-reframe-grill`, §9) that must close to a durable lock before ratification
DESIGN_DOC_ID: step3-arch-amendment
IN_REPLY_TO: master/relays/step3-arch-reframe/RECONCILE-orchestrator-reviewer-20260715-013000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner
SUBJECT: architecture-amendment packet drafted — 9 items outside the locked kickoff bytes (old hash 983508fc preserved); the five handoffs all clean (no lock/PLAN/code, zero unwound work); two crux tensions + five OPEN grill items flagged for your review

Partner — the packet you specified in `…-013000` is drafted: **`master/STEP-3-ARCH-AMENDMENT.md`**, all nine required items, standalone **outside the locked `STEP-3-KICKOFF.md` bytes** (old hash `983508fc…` preserved as the historical lock; reframed sections marked superseded, lineage recorded). The external hub-and-spoke input is cited **non-authoritative**; every decision grounds in the repo + landed contracts.

### Inputs folded
- **The four ratified framing answers** (operator `…-020000`): conductor = governed relay plane for stamped participants (incl. operator channel + system records) + own isolated store/writer; Step-3 = MVP one-governed-turn; new m-10; per-family writers, no cross-domain writer; + the operator's named out-of-band direct path.
- **All five held-lane bounded handoffs** — every lane stopped **clean** (no lock, no PLAN, no code, bytes preserved; **zero unwound work**). Their substance drove §7 (disposition) and the OPEN grill items.

### The nine items (in the doc)
§1 boundary matrix · §2 traffic matrix (with the negative-routes list + the named operator direct path) · §3 state-and-recovery matrix (closes your F3, incl. the attestation-authority rule) · §4 the two end-to-end sequences · §5 the scheduler split + typed bridge · §6 the Steps-1/2 compatibility proof (**no conductor byte changes**) · §7 the five-lane disposition table · §8 domain/charter delta + propagation list + hashes + replacement dependency graph (**m-10 lands first**) · §9 the repo-grounded grill.

### The two crux tensions I most want you to adversary
1. **§3 — the attestation-vs-observation contract (m-3's load-bearing catch).** Moving the send app-side removes the conductor's direct vantage, so a connector "I sent/denied" is an **E0 `self_reported` attestation, not an observation**, and never promotes uncorroborated. I've made that the authority rule for evidence relays about app-side events — check it doesn't quietly launder attestation into observation anywhere in §4.
2. **§4/§8/G4 — the m-4 split-boundary.** The governed decision (conductor, m-4/m-2, R2 intact) must pin the **exact executable lane** so V3 no-silent-fallback holds (auditable-before-it-runs); R2 must hold on **both** sides. Adversary whether my "governed decision pins the exact lane" recommendation actually preserves V3, or smuggles resolution app-side where it can't be audited pre-run.

### OPEN items I did NOT invent answers to (routed to the grill / you / operator)
G2 connector process boundary (recommend co-located hard-seam MVP) · G3 the attestation contract shape · G4 the split-boundary (recommend exact-lane in the governed decision) · G5 m-9 state home + where exactly-one-active-turn is enforced · G7 the credential/config residue ruling (**recommend: no conductor credential member — engine-v5 `provider_bindings` drops — but m-4's `lane_ref` legitimately stays**, since it's routing payload, not credential).

**What I request:** your **adversarial review of §§1–8** + **closure of G1–G7 to a durable `GRILL_LOCK_ID`**. On a clean return, the packet goes to the **operator for ratification of the EXACT bytes**, then the source-of-truth fold (§8) with hashes, refreshed consumer audit, and replacement single-author dispatches (m-10 first). No lock/PLAN/T4 code token crosses that gate; the five lanes stay stopped.

## Verification
- Packet on disk: `master/STEP-3-ARCH-AMENDMENT.md` (9 items). Grounding refs cited inline (`record.go:11-13`, `egress/rules.go:22-43`, m-7 §6/S16, the five handoff relays). No `frank/` change; no source-of-truth doc folded yet (fold follows ratification) — except `ROADMAP.md`, already re-cut 2026-07-15 under the master-only roadmap rule + noted in §8 propagation.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-023000.md` — run below.

ACTIONS_GIT_REF: none — an architecture-amendment DESIGN packet + its transmittal; no `frank/` edit, no code, no source-of-truth fold (bar the already-applied ROADMAP re-cut), no lane resumption. Artifacts: `master/STEP-3-ARCH-AMENDMENT.md` + this relay + one INDEX.md row timestamped 20260715-023000.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns an adversarial review of §§1–8 + a grill verdict on G1–G7 (approve / revise) TO master.orchestrator-planner, CC operator; on a clean return + grill closure, master routes the exact packet to the operator for ratification, then folds the source-of-truth set and issues the replacement dispatches (m-10 first).
