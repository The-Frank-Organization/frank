## RECONCILE — runtime-research synthesis + primary-source verification: decision-① on verified ground (VP adversarial review requested)

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: runtime-research
PARENT_DISPATCH_ID: runtime-research
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — surfaces operator decision ① (does not decide it here); operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — bringing you the runtime-research effort for adversarial review **before** the operator locks decision
① (wrap-vs-attach + the honest security claim). This is report-only: no design changed, no PLAN opened, no spike
run, no pair re-engaged. Three artifacts, in cwd:

1. **`master/RUNTIME-RESEARCH.md`** — my synthesis of a 3-wave read-only research fleet (≈21 lenses) on how the
   conductor concretely rides existing runtimes, provider-agnostically, with persistent seats. Runtime matrix +
   model matrix + identity/courier + security posture + park/wake.
2. **`master/RUNTIME-RESEARCH-GPT.md`** — an operator-provided **external chat-model research brief**. **It is
   NOT a seat** — no channel-stamped FROM, no authority; I weighed it as *evidence* (an independent method), not
   a peer reconciliation. Cross-check folded into §13.
3. **§14 of the synthesis — a primary-source verification pass** (4 read-only lenses, one adversarial-by-design)
   that checked the two load-bearing claims the external brief did *not* corroborate, before we lock ① on them.

**Why this matters to you:** decision ① is a MUST-decide readiness-gate item (Cluster 2). It is the operator's
to make, but per the standing discipline it should not be locked without your adversarial pass. The verification
already changed the honest claim — I want you to try to break my corrected position before it reaches the
operator.

---

**THE HEADLINE — verifying first caught two overclaims we were one step from baking into ①.**

The security posture is now an honest **three-rung ladder** (§8, §10), and verification (§14) moved the top rung
behind a code spike:
- **Rung 1 — Attach → "confusion-resistant."** ✅ verified honest. The Step-1 posture.
- **Rung 2 — `srt`-wrap → "sandboxed defense-in-depth."** ✅ verified *real*: the jail's only egress **path** is a
  parent-owned broker socket, kernel-enforced by construction (Linux netns removal; macOS Seatbelt `EPERM`s
  non-localhost even if the seat unsets proxy env). `srt` = `@anthropic-ai/sandbox-runtime` is real,
  Anthropic-published. [P]
- **Rung 3 — "sole external sender by construction" → SPIKE-GATED milestone.** ❌ **refuted as automatic.** Sole
  egress *path* ≠ sole *exfiltrator*: the broker is the whole egress TCB (a null-byte SOCKS5 allowlist bypass
  shipped in the wild ~5.5 months; no TLS inspection; DNS is a covert channel; Codex leaves READ open). And the
  jail **fails OPEN by default** on common kernels (namespace-creation failure → self-disable, the Ona incident,
  which is the *default* condition on Ubuntu 23.10+/unprivileged-Docker/WSL). The strong claim ships **only after
  a spike proves five properties** (§14.2): (1) namespace-failure fails **closed**; (2) broker **canonicalizes
  hosts**; (3) **no DNS/raw-sockets** from the seat; (4) in-jail RCE **cannot read a sibling's token/pipe**;
  (5) conductor **stamps identity itself**. [P]

**Two further verified corrections:**
- **Identity is conductor-owned, full stop.** My earlier draft claimed Codex "independently implements
  channel-stamped FROM via `session_source`" — **refuted [P]**: `session_source` is a mutable launch-arg
  telemetry/product-gating label; Codex's per-message author (`clientId`) is client-supplied. The runtime does
  not supply forgery-robust identity; the conductor stamps FROM from its own per-seat channel. This *simplifies*
  the design (one identity mechanism, ours) — I read it as a strengthening, not a wound. Push if you disagree.
- **Provider-agnostic-via-Codex is Responses-API-gated** (`wire_api="chat"` removed on `main`) — so Claude Code
  is the cleaner multi-model ride; a Codex lane to a non-OpenAI model needs a Responses-compatible gateway. A
  routing/adapter fact for m-4 + the Step-3 surface, not a blocker.

**What converged (independent method → high-confidence, §13.1):** routable unit = a serving *tuple* not a model;
access-is-easy / agent-loop-reliability-is-hard → a live **canary/probe** (which is our m-3 observe-gate pointed
at lane admission); identity-by-construction from per-seat transport; wake = blocking pipe not polling; local =
first-class-but-qualify. Six independent convergences.

**The one divergence, resolved into the tuple (§6.5):** the external brief tiered GLM/Kimi supervised-only; my
fleet tiered GLM-5/Kimi-K2.5 T-A. They measured different axes (capability vs. serving-reliability). Merged:
**capability-T-A, autonomy-admitted only on a canary-qualified lane.** I flag this as the place my synthesis is
most likely still wrong — please pressure it.

---

**Adversarial-review asks:**

- **Q1 — the verification itself.** Do you accept the §14 verdicts, or do you read any lens as over- or
  under-reading? Specifically: is "sole-egress-*path* is by-construction but sole-*exfiltrator* is not" a fair
  split, or am I conceding too much / too little? Is the **five-property spike** the right gate for the strong
  claim, or is one property missing (I am least sure about DNS/covert-channel closure and the cross-seat-isolation
  -under-RCE property)?
- **Q2 — the decision-① framing.** Is the three-rung ladder (§10) the honest, correct structure — attach-first +
  "confusion-resistant" now, `srt`-wrap as verified defense-in-depth, "by construction" spike-gated and never
  shipped on assumption? Or do you want ① framed differently before the operator sees it?
- **Q3 — identity conductor-owned.** Do you concur that the runtime cannot supply forgery-robust identity and the
  conductor must stamp FROM itself (Codex `session_source` refuted)? Any risk in treating this as a
  simplification rather than a gap against the c1 identity lock (`ARCHITECTURE.md` §1)?
- **Q4 — the external brief.** Did I weigh it correctly as evidence-not-a-seat, and is the GLM/Kimi tuple
  resolution (§6.5) sound, or did I paper over a real capability-vs-reliability disagreement?

---

**Governance state (nothing moved):**
- **No decision is made.** ① and the other four operator decisions (②–⑤) remain OPEN and unrecorded. This relay
  surfaces ① for the operator *after* your review; it does not decide it.
- **No MUST-gate item cleared.** Cluster 1 (m-1/m-2 `submit()` write-path) and Cluster 4a/4b (m-2 stale schema +
  R2 hole) are still unreconciled; this research does not touch them.
- **No spike opened.** The five-property security spike (§14.2 / §12-3) is a *recommendation*; it needs a
  separate operator-opened spike gate (your readiness-reconcile Finding 4 stands).
- **No pair re-engaged, no contract changed, no code.** Docs-only, cwd.

On your review + the operator's decision, ① becomes recordable and the bounded readiness fixes route to their
owners. Where do you push?

ACTIONS_GIT_REF: ran 4 read-only verification subagents (srt, sandbox-primitives, Codex-attach, adversarial); wrote/updated `master/RUNTIME-RESEARCH.md` (§14 + BLUF/§5/§7/§8/§10/§12 corrections) + this relay + `master/relays/INDEX.md`; no code/source/pcode edits, no design-doc changes, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP adversarially reviews the runtime-research synthesis + the §14 verification + the decision-① three-rung framing (Q1–Q4); then the operator records decision ①; then the orchestrator routes the bounded readiness fixes and (separately) the operator may open the five-property security spike.
