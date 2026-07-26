## RECONCILE — OPERATOR RATIFICATION RECORD for the esc1 re-scope + witness re-cut, decisions 1–8, **agent-authored and operator-cited per §8b** (no operator `FROM` is forged). This is the durable authority the follow-on routing cites; without it every downstream act would cite a conversation. **Headline: the durable session log trims to a floor, the context fingerprint is DROPPED as an accepted risk with the leg's claim narrowed to match, an eleventh fencing scenario is ADDED, and dead receipt members are removed now rather than carried.** Two reversals of my own recommendations are recorded as such. Five routing acts follow; none is performed here.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-ratify-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-fork-4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this RECORDS decisions the operator has already made in-session; it is the citable authority for them. It performs no trim, authors no amendment, edits no owner's document, and moves no byte. Every consequence below is routed to its owner, never executed here.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-fork-4/RECONCILE-orchestrator-planner-20260726-004504.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, l4.planner, l4.implementer
SUBJECT: Operator ratification of esc1 decisions 1–8 — D1 floor adopted (chaining/rotation/seal/cross-segment out, one-file-per-run in) · runtime self-integrity OUT OF SCOPE with a named re-entry test · dead receipt members removed NOW · `log_prefix_digest` collapses to one whole-prefix comparison · `context_digest` DROPPED as accepted risk + leg claim narrowed · parity dissolved · eleventh fencing scenario ADDED · provisional supersession routed to its OWNER per the binding fence-row-vs-owner-path rule

**Provenance discipline.** Every decision below was made by the operator in-session on 2026-07-25/26. This relay is **agent-authored and operator-cited** under §8b; no operator-authored relay is fabricated. The VP is `TO` because adversarial review of master's decomposition is its seat; the operator is `CC` as the deciding authority.

## The eight decisions, as ratified

| # | decision | authority |
|---|---|---|
| **1** | **Hash-chaining OUT** of the durable session log; **runtime self-integrity OUT OF SCOPE** for the MVP. | operator |
| **2** | **The whole releasable set goes:** size rotation · terminal seal · cross-segment boundary equation · **plus one-file-per-run**, without which cross-segment is not actually dead. | operator |
| **3** | **Remove the dead receipt members NOW**, not later — *"we haven't started building yet, now is a good time to keep schema bloat away."* | operator |
| **4** | `log_prefix_digest` collapses to **one whole-prefix comparison** (the produced log must equal the frozen expected), not an ordered per-record digest list. | operator, on master's recommendation |
| **5** | **`context_digest` DROPPED**, and the requirement that **external editing of sessions be permitted** is ADDED. | operator |
| **6** | **Dissolved** — parity was "should an observer re-derive the fingerprint"; there is no fingerprint. | consequence of 5 |
| **7** | **ADD the eleventh fencing scenario.** | operator, on master's recommendation |
| **8** | **Provisional supersession annotation approved** — and **routed to its owner**, not executed by master. | operator |

## Decision 1 — the re-entry test is the durable part

Chaining is released on three independent grounds: m-9's own assessment against its own §9 survey (malice-only over per-record checksum + `seq` contiguity); m-3's evidence-side reasoning (chaining is *internal self-binding* for a no-ground-truth regime, and the exit leg has a frozen oracle — a reason robust to scale, unlike the throughput argument master offered); and the operator's decisive ground — **`bivpak` performs exactly the wholesale internally-consistent rewrite chaining is built to reject**, legitimately.

**Named re-entry test, operator's words:** *"we can bring them back in the future if agents start tampering with past session data, somehow, but very unlikely."* Recorded so a future reader sees **why the mechanism is absent and what would restore it**, rather than finding a gap and re-deriving the apparatus from the same instincts that produced it the first time.

**Runtime self-integrity is out of scope by the same decision**, closing m-3's scoped caveat: one cannot want production resume to reject rewritten logs while building a tool whose purpose is rewriting them.

## Decision 5 — recorded as a RISK ACCEPTANCE, not as a refuted analysis

m-3's finding stands: without a context witness, a resume that rebuilds a *different* prompt from a *correct* log scores PASS. The operator overrules on cost-versus-likelihood — *"overengineering for an edge case that almost never happens"* — which is a legitimate call, and the record must say so honestly rather than imply the analysis was wrong.

**Bound to it, non-severably: the Durability leg's claim NARROWS.** It now proves *the record came back intact*, **not** *the model resumed the same conversation*. Keeping the stronger language while the check is gone is precisely the false PASS m-3 identified — claim without evidence. This is the sandbox pattern: drop the mechanism, narrow the claim, write the gap down. **The narrowed wording is owed at amendment time and is not optional.**

**The added requirement (external editing) is a PRODUCT property, not a test concern.** Drivers: a session can become unresumable through no user fault (precedent: malformed thinking blocks rendering Claude Code sessions unopenable, recovered by hand-editing), and **a session that cannot be repaired is a data-loss event**; plus third-party tooling (`bivpak`). The governing rule is **label, never gate** — on a content mismatch frank says *"this session was modified after it was written"* and resumes. Consequences routed, not executed: per-record checksums become advisory; **m-10's `receipt_conflict` must not hard-fail an edited session** (as frozen it is fail-closed with first-committed-stands, turning a legitimate repair into a hard error). Repairability itself is owed as a stated design property — it currently exists only in conversation.

**Note: this does NOT disturb Decision 4.** The exit comparison runs inside a scripted scenario where nothing is edited. External editing breaks *runtime* gating, not the test.

## Two reversals of master's own recommendations, recorded as such

- **Decision 4.** Master carried m-3's ordered-list form forward as the natural replacement for the chain, without asking whether *either* was needed once record contents are frozen. The operator's challenge was correct: the harness holds both sides, so it compares content; a digest is a compact stand-in for content you do not have. Diagnosis is not lost — on failure the harness diffs two files it already holds. Same blind spot as the apparatus itself: **replacing a mechanism instead of asking what it was for.**
- **Decision 5.** Master first recommended a composite field to keep the ratified vector byte-exact, then recommended pinning the assembly invariant instead. The operator's ruling on 3 dissolved the first, and the second was itself doing too much. Both superseded.

## Consequences and what falls away

**Falls away:** m-9's P2 refinement (cross-segment case + chaining-discrimination negative legs) · master's strengthening (v) (≥2 durable checkpoints) · the objective-assembly pin (nothing left to condition) · the composite formula, its version envelope, its explicit-null aliasing rule and its both-halves-in-evidence requirement · Decision 6 entirely.

**Costs accepted with Decision 7:** the frozen `sample_weight` accounting must re-balance to exactly 30 governed turns + 100 tool calls, and the leg-cardinality problem the reviewer already flagged (six legs declared, seven rows present) is reopened.

**Forward direction recorded, not scoped:** git-like session versioning, own implementation, **git explicitly NOT a dependency** — now inline in `ROADMAP.md`'s Step-4+ block with the GPLv2 flag on the vendor route and the packfile cost boundary. The MVP floor imports into it cleanly, which the chained log would not have.

## Decision 8 — routed to its OWNER, per a binding rule master nearly broke

The two stale loci (`2026-07-19-mvp-full-worker.md:88` crash-semantics and `:155` fixture item) live in a **GRILL-ratified m-9 document**. Master annotating it directly would violate `CYCLE-PLAYBOOK.md:222`: *"a fence row licenses a FILE; it never substitutes for the OWNER PATH on a locked contract living inside that file… route the owner countersign **before** the edit."* **The operator caught this before I filed it.** The playbook also names the cause — *"the cheap correct routing was available and skipped because hand-relaying it felt expensive"* — which is the same transport tax, from the same direction, and it is in master's own standing memory as binding. **So: the operator ratifies; m-9 authors the annotation in its own document.**

## The five routing acts owed (none performed here)

1. **m-9** — the provisional annotation (Decision 8, m-9 to author) · advisory-not-gating checksums · does last-record completeness cover torn writes alone (which would shrink the floor again).
2. **m-10** — `receipt_conflict` must not hard-fail an edited session; a change to a fail-closed rule m-10 owns.
3. **m-3** — the exact member kill-list (`segment_id` is dead under one-file-per-run; **`seq_hwm` probably is not** — a high-water mark still means something in a single file) · confirm a whole-prefix comparison suffices without the per-record list · formal notice that the context check is dropped as accepted risk, not refuted analysis.
4. **m-9 + m-10** — the member-set removal amendment (Decision 3), which re-opens the co-signed receipt body and is m-3-joined.
5. **lane 4** — the eleventh fencing scenario, the sample re-balance, and the cardinality reopening.

## Boundaries
This relay records decisions and authorises nothing further. It performs no trim, authors no amendment, edits no owner's document, annotates no locked file, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, §D-settlement amendment `1fa71cb8…`, m-9 delta `01b885fe…`, m-9 stage-4 `2026-07-19-mvp-full-worker.md`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- All ten owner returns of this escalation read at their bytes and SHA-256-bound in `…-esc1-fork-4/RECONCILE-orchestrator-planner-20260726-004504.md` (`1995bdef419f0d54ef6efeaa02b558bedf91384fbe601c901517dbc3338bf2d7`).
- Baseline verified against source, not recollection: across all eight harnesses in `references/`, every chaining/tamper identifier match was unrelated to session logs; **zero session-log hash chains**. Prompt-hashing exists in exactly three and never for resume verification — claude-code `hashSystemPrompt` (telemetry span attribute, system prompt only), jcode `messages_hash` (TUI re-render/cache invalidation), kimi-code (test fixtures + one `systemPromptHash` assertion). Codex writes `trace.jsonl` via `append_with_context_best_effort` (failures `warn!`-ed and swallowed) with **zero** fsync/sync_all/checksum/crc in the module.
- The binding owner-path rule quoted verbatim from `master/CYCLE-PLAYBOOK.md:222`; the two stale loci at `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md:88`/`:155`; the recorded root supersession at `2026-07-22-relock-lane2-m9-delta.md:373` (§5).
- Forward direction inlined in `ROADMAP.md` (Step-4+ block), rewrapped to the file's 108-char convention; the standalone draft doc was removed at operator instruction.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this ratification relay + one INDEX.md row + the earlier `ROADMAP.md` inline edit (master-owned per the standing maintenance rule). No trim performed, no amendment authored, no owner document edited or annotated, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the **VP** reviews this ratification record against the ten bound returns and the eight decisions — in particular whether Decision 5's narrowed Durability claim is bound tightly enough to prevent a claim-without-evidence, and whether any consequence has been recorded as ratified that the operator did not actually decide. Master then issues the five routing acts. Amendment drafting, lane-4 resume, fixture materialisation/freeze, re-lock, T4, and external use remain held.
