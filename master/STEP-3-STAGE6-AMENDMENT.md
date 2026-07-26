# Step-3 Stage-6 Amendment — the bounded coding-agent-MVP correction (rev12: time-scoped trust invariant; one terminal overflow state; observed-action fixtures)

**Status: PROPOSED rev12 — pending VP decomposition review r12 → operator re-scope ratification.** Supersedes rev11
(`61fe014c0fe66c3096a750d9da3ca08c3ae6030f3c4a891b62749a0ee20da0dd`, held by VP decomposition-review r11
`RECONCILE-orchestrator-reviewer-20260721-162721.md` — the producer-total three-class manifest, the post-inspection
`content_lost` result, the source-real schemas, and the no-successor overflow direction CLOSED). rev12 folds r11's
three exact-byte consistency defects: **D2** — replace the timeless `settled_with_content ⇒ content-present`
implication (which forbade the accepted `content_lost` branch) with **TWO time-scoped properties** — settlement-time
*evidence* (content+marker/receipt durably linearized then) + resume-time *evidence-AND-current-presence* trust
(evidence without current content ⇒ `content_lost`); **D3** — replace the `run-terminal/parked` token with **ONE
terminal lifecycle state**: the run commits to terminal **`FAILED`** with the closed `resume_frame_overflow` reason,
no same-run successor/revival, operator manual `resume_action`; **F106** — `xit-dur-3` requires exact `uncertain`
carriage on **both** missing-half orders, `xit-dur-4` requires an **observed** exactly-once post-receipt action (not
mere reachability), `xit-dur-5` asserts the exact terminal `FAILED`/`resume_frame_overflow` state. Subsystem internals
stay delegated under F73. Additive to `STEP-3-MVP-AMENDMENT.md` (r7, `2f75f2a1…`) and the reframe packet
(`2d240eb6…`). It **HOLDS** the joint interface-lock record (`…-023000.md`, `b7e1f0ef…`); that lock does not issue,
its pending operator gate is superseded by the re-scope gate this amendment requests. Nothing here reopens the
ratified architecture direction (§2). Until ratified: proposed design input only — no DESIGN-lock, PLAN, T4 token,
credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use (§10) is authorized.

**Basis:** the stage-5.1 review (`b4e79f3b…`); VP decomposition-reviews r1 (`033631`) + r2 (`042157`) + r3
(`043904`) + r4 (`051503`) + r5 (`062953`) + r6 (`071202`), all accepted; the operator grill (§3 GRILL_LOCK
`step3-stage6-rescope-grill-1`, incl. D7 build-resume + the operator's 2026-07-21 build-it-properly confirmation) + VP
r7 (`073500`) + r8 (`153916`, grain accepted) + r9 (`160120`) + r10 (`161600`) + r11 (`162721`). rev12 folds r11's three
exact-byte consistency defects (time-scoped trust invariant · one terminal overflow state · observed-action fixtures);
no new product choice.

---

## §0 — Why an amendment (unchanged)
Both stage-6 halves are complete on exact bytes; the review faults the milestone claim + lock scope, not the frozen
bytes. This converts the pending all-artifact lock into a bounded interface amendment + a hashable lock re-cut,
routes non-interface items to T4 / the exit test, and re-runs a shorter lock over only the changed interfaces.

## §1 — Disposition (updated to the r11 closures)
CLOSED: **F101/F102/F103/F104/F106-handoff/crash-predicate**, the worker-owned placement (D7), the **grain boundary**
(no operator arbitration), **F105-D1** at decomposition grain, the **D3 first-action table**, the **r8 seams**
(marker-before-outcome ordering, provider composite-settlement, source-turn identity, immutable snapshot bytes, the
`PENDING` disposition + post-commit receipt gate, the D1 ownership branch), the **r10-accepted r9 folds**
(determinate-terminal semantics: genuine zero-content split from determinate-discarded partial + `NOT_INVOKED_INTEGRITY_
FAULT`→`turn_failed`; the no-auto-advance first-action table; complete-frame pre-sizing inputs + one-carrier/no-chunking;
the `xit-dur-3/4/5` sub-fixture structure under Durability), and the **r11-accepted r10 folds** (the producer-total
three-class manifest evidence with completed-without-receipt ⇒ `uncertain`; `content_lost` as the m-9 post-inspection
reconciliation result; source-real per-entry schemas with no provider `args_digest`; the no-successor overflow
direction). rev12 folds r11's three exact-byte consistency defects:
**D2 (F105-D2-R11)** — the trust invariant is now **two time-scoped properties**, not one timeless implication (which
would have forbidden the accepted `content_lost` branch): (1) at settlement/content-ready commit, `settled_with_content`
is durable *evidence* the content + admitting marker/receipt fsync-linearized **then**; (2) at resume, m-9 trusts a
result **only** under matching positive evidence **AND** presence in the **current recovered valid prefix** —
evidence-without-current-content ⇒ `content_lost` → `DEGRADED`, never trusted/reconstructed; **D3 (F105-D3-R11)** — the
oversized-`turn_open` outcome is now **ONE terminal lifecycle state**: the run commits to terminal **`FAILED`** with the
closed `resume_frame_overflow` reason (not a `run-terminal/parked` ambiguity), **no same-run successor or revival**, no
lease/snapshot, operator manual `resume_action`; the terminal-vs-parked + no-successor choices are master-level (exact
encoding is pair DESIGN); **F106 (F106-R11)** — `xit-dur-3` requires exact `uncertain` carriage on **both**
missing-half orders (terminal-first/receipt-absent AND receipt-first/terminal-absent, omission FAILS either way),
`xit-dur-4` requires the selected first action's durable/wire effect **OBSERVED exactly once after the receipt with zero
before** (mere reachability insufficient — a forever-held worker FAILS), `xit-dur-5` asserts the exact terminal
`FAILED`/`resume_frame_overflow` state — still six legs.

## §2 — The ratified architecture direction is NOT reopened (unchanged)
Conductor one isolated service; provider bypasses it; **m-10 owns lifecycle not policy + remains F59 host; m-9
requester/executor, inert-until-authorized**; m-8 holds provider creds; stores separated by state family; UNKNOWN
first-class; one lane + one turn; fixed 8-name catalog; native+MCP parity; H-17 census in the lock; **m-5 stood
down, m-6 out**. Adds interface obligations + a hashable re-tier; withdraws no approved mechanism; moves no bound byte.

## §3 — GRILL_LOCK (operator, 2026-07-21) — unchanged
`GRILL_LOCK_ID: step3-stage6-rescope-grill-1`. **D1** sandbox forgone (ambient bash; gap documented; H-12 promoted,
§10). **D2** bash claim narrowed (exact invocation-CONTEXT binding; no affected-resource/per-effect-hold claim;
fixed-policy holds dropped). **D3** exit GATE = six property legs + objective overhead budget (§7). **D4** utility
DEMONSTRATION (non-gated) = public dogfood (CRM — informal ask, operator retains all rights — + bivpak) + honestly-
labeled agent-as-operator SWE-bench, no threshold gate. **D5** dogfood/real-work ⊥ exit. **D6** effect descriptor =
context-binding (`backend_id="ambient"`). **D7 (2026-07-21): BUILD durable session state + resume** — grounded in
the field (Codex rollout+replay, Claude Code transcript+`--resume`, opencode message-store + git-snapshot, deepagents
LangGraph checkpointer all persist + resume; resume is uniform, unlike the split on sandboxing). Re-cut to the
field-standard **worker-owned session-content log** (not an m-10-hosted store): the worker writes its own durable JSON
log + a replacement reads its own file — **outcomes stay m-10-canonical**, the opaque replay envelope stays in-memory
(K6). Reverses the ratified "no m-9-owned durable session store" invariant (narrowed to "no m-9-owned durable
*outcome* store"). *Rejected:* the m-10-hosted journal (forced a blob wire + illegal store read-back); full
transcript-replay resume through an `INTERRUPTED→RESUMING` transition (heavier than a fresh continuation turn). Still
operator-owned (non-gated): dogfood N/timing; SWE-bench slice + governed-two-seat.

## §4 — Item A: the Tier-HARD bundle, soft-edit-stable (F101)
**Top-level artifact** `master/STEP-3-INTERFACE-BUNDLE.json`, schema `step3-interface-bundle.v1`:
`{ schema, recipe_version, recipe_sha256, bundle_sha256, lock_payload, provenance }`.
- **`lock_payload`** (the ONLY input to the constitutional digest) = `{ recipe_version, recipe_sha256,
  interfaces: [ {interface_id, extracted_sha256} … sorted ascending by interface_id ] }`. **`bundle_sha256` =
  SHA-256 over JCS(`lock_payload`).** It contains recipe identity + ordered interface ids + extracted-HARD-region
  digests, and **excludes mixed-document full-file source SHAs** — so a Tier-SOFT edit (bytes outside any HARD region)
  does NOT change `bundle_sha256`.
- **`provenance`** (integrity only, NOT fed to `bundle_sha256`) = `{ sources: [ {interface_id, source_path,
  source_sha256, region ∈ "marker"|"whole_file"} … ] }`. For a `whole_file` source the whole file is hard, so its
  `extracted_sha256` IS the full-file SHA and correctly rides `lock_payload`; a `marker` source's full-file
  `source_sha256` rides provenance only.
- **Hard-region markers:** literal `<!-- HARD-BEGIN interface_id=<id> recipe=<v> -->` … `<!-- HARD-END
  interface_id=<id> -->`; the extractor takes the enclosed bytes verbatim, `extracted_sha256` over them.
- **Fail-closed** (`--verify` exits nonzero; extraction refuses, no bundle, re-lock blocked): declared marker span
  absent · duplicate `interface_id` · `source_sha256` on-disk mismatch · `extraction_recipe_version` mismatch ·
  ill-formed span · **an UNDECLARED `HARD-BEGIN` marker present in any declared source but absent from the manifest**
  (a full-inventory marker scan, not just the declared set).
- **Owner/artifacts:** generator/verifier = master; extractor+recipe = `master/tools/extract-interface-bundle.py`
  (`recipe_sha256` = its own digest, at the named top-level field); `python3 …/extract-interface-bundle.py --verify`
  recomputes every `source_sha256`, re-extracts, checks `extracted_sha256`, runs the undeclared-marker scan, and
  recomputes `bundle_sha256`. **Negative fixture `bundle-soft-stability`** (shipped with the extractor): editing bytes
  OUTSIDE any HARD region leaves `bundle_sha256` unchanged while changing `provenance.source_sha256` — the executable
  proof the re-cut works. Bundle authored LAST (§11 step 4), F73-reviewed; a HARD byte change moves `bundle_sha256`
  → F73; a SOFT edit does not.
- **Soft-design ledger:** Tier-SOFT changes → `master/SOFT-DESIGN-LEDGER.md` (or the owning domain history), NOT
  `PROTOCOL-DEVIATIONS.md`.

## §5 — The interface items (B–E), executable contracts (Tier-HARD)

**B — `frozen_core_digest` join** (unchanged from rev3): the m-8-computed digest as a field on m-8 terminal · m-10
`provider_attempts` · m-3 `m3.app_event.v1` E0 · m-3 `m3.e3_observation.v1` attempt vector · the composite exit
proof. Observer derives it independently; no prompt/response bytes enter the conductor.

**C — the effect descriptor (context-binding evidence; F103 resolved).** The ticket binds the descriptor; per-action
applicability (`R` required / `∅` inapplicable):

| field | read/write/edit | apply_patch | bash | relay.* |
|---|---|---|---|---|
| `action` | R | R | R | R |
| `canonical_resource` | R = one workspace-root-resolved canonical path | R = **ordered target-set digest** (SHA-256 over the JCS-sorted resolved paths the patch names) | ∅ | R = relay verb + target id |
| `workspace_root_id` | R | R | R | ∅ |
| `cwd` | R | R | R | ∅ |
| `env_digest` | ∅ | ∅ | R | ∅ |
| `shell_interpreter_ref` | ∅ | ∅ | R = `{path, version, content_id}` resolved at invocation (see below) | ∅ |
| `backend_id` | R (`"ambient"`) | R (`"ambient"`) | R (`"ambient"`) | R (`"conductor-client"`) |
| `network_policy_id` | R (`"none"`) | R (`"none"`) | R (`"none"`) | R (`"none"`) |
| `tool_impl_ref` | R = WRAPPER (worker/backend, F58) | R | R = WRAPPER only (the interpreter is `shell_interpreter_ref`) | R = shared-client component |
| `canonical_args_digest` | R | R | R | R |
| `one_shot` | R (true) | R (true) | R (true) | R (true) |

**HARD semantics (bundle):** the schema + this table; the **single cwd encoding (F103):** exactly ONE canonical form
— the **workspace-root-relative POSIX path** (symlink-resolved · NFC-normalized · no trailing slash), paired with
`workspace_root_id`; the root itself serializes as `"."`; a **nonexistent cwd is a pre-spawn typed reject**
(`rejected_local`-class), never a descriptor value (bash requires an existing cwd) — no "absolute-and-relative"
ambiguity. **`env_digest`** = SHA-256 over the JCS-sorted COMPLETE env set ACTUALLY presented to the child — the
**m-1-sanitized environment** (m-1 hardening #1/#7, already frozen; no behavior delta — we digest what is presented,
we do not newly clear it); m-1 owns the sanitization rule + reviews secret-sensitivity. **`shell_interpreter_ref`
(F103, mandatory-or-unknown):** `{path, version, content_id}` — `path` = the symlink-resolved absolute interpreter
path; `version` from the interpreter's version command (e.g. `bash --version` first line); **`content_id` is
REQUIRED** = the SHA-256 of the resolved interpreter binary OR an equally-immutable OS/package-artifact digest — a
mutable binary keeping the same path+version is NOT exact identity, so a bare path+version is insufficient. If
`content_id` is **unobtainable, `local_invocation_matches_effect_descriptor` returns `verdict=unknown` (HOLDS the
leg)** — never a weaker pass. **`canonical_resource` per action** incl. the apply_patch ordered target-set digest;
teardown + UNKNOWN-visibility semantics (frozen worker §2.5). **SOFT:** descriptions/wording/presentation + tunable
numeric caps (names unremovable, values T4-tunable). No containment/affected-resource/per-effect-hold claim (D2). C is
the EVIDENCE half; **H-21** (backlog) is the Step-4 effect-descriptor AUTHORIZATION successor. *Owners:* m-10 (ticket
schema + gate), m-9 (executor derivation + record), m-1 (env/no-leak review). Authority boundary unchanged.

**D — durable session state + resume: the worker-owned session-content log (F105, re-cut per D7).** The
field-standard shape: the worker owns its OWN durable JSON session log (a plain file in the **per-run** runtime dir),
and a replacement generation reads *its own file* — like Codex's rollout / Claude Code's transcript. This is NOT
m-10-hosted; the worker never reads m-10's private store back, so there is **no cross-process blob wire** and the
one-way `record_tool_outcome` is **unchanged** (the rev5 `commit_round`/`round_committed` frames are DROPPED).

**Invariant supersession (named; needs an m-9 owner delta + confirmation).** The ratified worker rule *"no m-9-owned
durable session store; recovery is re-derive"* (worker §7.1/:85-88) is **narrowed to "no m-9-owned durable OUTCOME
store."** Content may persist in the log; **every outcome/decision stays m-10-canonical** (tool_calls, provider_attempts,
tickets) — so no-second-*outcome*-truth holds by construction; the log is a context cache, never a rival record of
what happened. The worker still holds **no durable authority** (the log carries no credential/ticket/epoch).

**What the log holds (CONTENT only) vs what it does not:**

| in the log (worker-owned content) | NOT in the log |
|---|---|
| assembled `input[]` items (user/assistant/tool_result), tool-call identities + args | effect OUTCOMES (stay m-10 canonical) |
| settled tool-result content (bounded §2a max-captured-tool-output) | the opaque `reasoning_replay` payload (stays in-memory, K6 §2.8 UNCHANGED) |
| provider-visible output items; compaction events + template/version | any S-A/S-B secret bytes (m-1 redaction review) |
| workspace snapshot id; objective/constraint refs; per-round index | conductor relay records (already readable via `project`/`read`) |

**Grain (this amendment vs the pairs' DESIGN).** D fixes the **decomposition**: the load-bearing *decisions*, the
seam *ownership*, and the *acceptance properties* — plus the named **Tier-HARD DESIGN obligations** the owning pairs
discharge post-ratification under F73 with their own adversarial pair review. The subsystem *internals* (exact record
grammar, the writer-fence mechanism, the segment/rotation state machine) are m-9/m-10 DESIGN deliverables, NOT
authored here — authoring them in the amendment would usurp the pairs' design ownership (CLAUDE.md: the m-x pairs own
local detail-design) and never converge.

**D1 — the crash-safe append log: the acceptance properties + the m-9 DESIGN obligation (F105-D1).** A worker-owned
line-framed append-only log in the per-run runtime dir (field standard — Codex rollout / Claude Code transcript).
Master fixes these **acceptance properties**: (i) **durable-append = the record's `fsync` return** (the §7
`session-log append` metric measures the fsync, not a buffer write); (ii) **an ENFORCEABLE exclusive-writer boundary**
— the property is *"a retired generation cannot corrupt or extend the successor's trusted prefix, and the durable
predecessor/stale-write boundary is decidable"*; a bare `generation_id` label does NOT satisfy it; (iii)
**valid-prefix recovery returns ONE deterministic prefix per crash cut**; (iv) **file identity** bound to
`{run_id, run_manifest_digest}`, fail-closed; (v) **content-before-outcome durability ordering** (D2 relies on it);
(vi) the path rides **`turn_open`** (one carrier). **m-9 DESIGN obligation (Tier-HARD, F73):** the closed record union
+ per-`kind` payload schemas + canonical encoding + `seq` grammar/contiguity/duplicate rule + `round_marker`
membership/digest + the exclusive-writer *mechanism* + the segment
seal/link/active-selection + rotation `fsync` order + the full append/handoff/rotation **crash table** yielding the
deterministic valid-prefix. **Ownership of the mechanism follows the branch m-9 selects:** a local **OS exclusive
lock** (acquired ordered-after-predecessor-termination) is **m-9-owned**; **m-10-ordered per-generation segments** is
a **joint m-10-producer/m-9-consumer seam** (designed + pair-reviewed together per §6-D), not an m-9-only obligation —
whichever branch, it must PROVE property (ii). Every content record MUST bind its `tool_call_id`/`attempt_id` +
**source `turn_id`** (so D2 is identity-exact across the continuation chain). The effect-UNKNOWN totality stays
m-10's, UNCHANGED.

**D2 — positive settlement reconciliation, ordered behind worker durability (F105-D2; RESOLVED).** The continuation
carries a **settlement manifest** (m-10→worker, produced from m-10's canonical rows), scoped over the **full
continuation ancestry within the run**; **every entry carries source identity `{run_id, turn_id, …-id}`** —
source-`turn_id` is required because frozen tool identity is `UNIQUE(run, turn, tool_call_id)`, so a bare id is not
identity-exact across continuation turns (the per-entry `args_digest` split is below). `last_settled_round_index` is
**removed** (no m-10 source; the worker derives the resumable round from its own `round_marker`s ∩ the settled set).

- **Tool settlement — marker-before-outcome ordering (the tool half, corrected).** The valid prefix ends at a durable
  `round_marker`; so fsyncing only the `tool_result` record before `record_tool_outcome` is one step short (a crash
  after outcome-commit but before the enclosing marker is durable leaves m-10 `EXECUTED` with content outside the
  prefix). Master's acceptance order: **the content record AND the marker that admits it to the valid prefix
  fsync-linearize BEFORE the settlement-producing step (`record_tool_outcome`).** Therefore **`settled_tools ⇒ its
  content is in the durable valid prefix**; the residual (marker not yet durable, m-10 not yet settled) is `uncertain`.
- **Provider settlement — a COMPOSITE fact (the provider half, corrected).** Frozen m-8 emits `attempt_result` at the
  attempt's own terminal boundary INDEPENDENTLY of m-9, and m-10 terminalizes `provider_attempts` from that CTRL-C
  fact — so the worker's provider-output fsync cannot be ordered before that terminal. **A `settled_providers` entry
  therefore requires BOTH (a) m-10's canonical provider terminal AND (b) a durable m-9 content-ready receipt bound to
  `{turn_id, attempt_id, valid-prefix/marker digest}`; m-10 does NOT emit the entry until both have committed.**
  m-10's own canonical attempt terminal is unchanged; the receipt is the new cross-owner conjunction (the exact
  frame/table is the pairs' under §6/F73). Missing content-ready evidence → **`uncertain`** (producer-decidable),
  never settled.

**Two stages — the m-10 pre-admission MANIFEST EVIDENCE vs the m-9 post-inspection RESULT (F105-D2-R10).** The
manifest is produced by m-10 from its canonical rows + the receipt-presence predicate **before** admission; m-10
CANNOT inspect m-9's private log, so it can only carry evidence it owns. The manifest is a **closed union of THREE
producer-total classes** — every canonical row maps to **exactly one**, no omission, no producer-time "or":
- **`settled_with_content` (positive evidence)** — tool `EXECUTED` (the marker-before-outcome ordering guarantees the
  content+marker were durable before m-10 committed the terminal); provider `completed` **AND** the content-ready
  receipt committed. (Evidence that content *should* be in the prefix — m-9 verifies at stage 2.)
- **`determinate_no_resume`** — a known terminal with **no resumable content**, carried row-identity-exact + first
  action, never auto-retried: tool **`NOT_INVOKED_INTEGRITY_FAULT`** (definite zero-effect; no `invocation_identity`,
  no content — action `turn_failed`, per §2.9/D-5); provider **`denied`/`REJECTED_LOCAL`/pre-transport `cancelled`**
  (genuinely zero-content) AND **`transport_failed`/`cancelled(post_invocation)`** (determinate terminal whose partial
  pre-terminal bytes — m-8 permits provider events before `failed`, and `cancelled(post_invocation)` is wire-crossed
  with partial content — are **discarded/untrusted for resume**, NOT claimed never-to-have-existed; carries the
  terminal + `cancel_point`).
- **`uncertain`** — tool `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT`; provider `UNKNOWN_PROVIDER_OUTCOME`/
  `PARTIAL_STREAM`; **and provider `completed`-WITHOUT-a-committed-receipt** (producer-decidable: no receipt ⇒
  uncertain). Surfaced; no auto-resend.

**`content_lost` is the m-9 reconciliation RESULT, not a manifest class:** after `turn_open`, m-9 inspects its log
and maps a `settled_with_content` entry that is **missing or corrupt in the valid prefix** to `content_lost` → the
`DEGRADED` disposition (via the §D3 receipt-gated report). It is never a pre-admission manifest member.

**Per-entry schema by owner-real source:** **tool** entries carry `{run_id, turn_id, tool_call_id, args_digest,
terminal}` (the frozen canonical args digest); **provider** entries carry `{run_id, turn_id, attempt_id, terminal,
cancel_point?}` and **carry NO `args_digest`** (m-10 has no canonical provider args-digest producer).

**The trust invariant is TWO time-scoped properties (F105-D2-R11), not one timeless implication** — a single
`settled_with_content ⇒ content-present` implication would forbid the very `content_lost` branch above (a
`settled_with_content` entry whose referenced content is missing/corrupt at resume). The two true properties are
temporal + conjunctive:
1. **At settlement (content-ready commit) — evidence property:** m-10 emits `settled_with_content` **only** because
   the content record AND its admitting durable `round_marker` (tool), or the committed content-ready receipt
   (provider), had **already fsync-linearized at that earlier point**. The class is a durable record that content
   *was* made durable then — nothing more.
2. **At resume (continuation trust) — evidence-AND-presence property:** m-9 trusts a `tool_result`/`provider_output`
   **only** when the manifest carries a matching `settled_with_content` entry **AND** m-9 finds the matching content in
   the **current recovered valid prefix**. Positive-evidence-but-content-absent/corrupt yields **`content_lost`**
   (→ `DEGRADED`), **never** trusted or reconstructed content.
Status is never treated as reconstructed content; reconciliation is positive + identity-exact over the closed union;
the opaque replay envelope is absent (K6), so on `content_lost`/degrade the worker re-reasons from the durable prefix
(valid — `reasoning_replay?` optional, stateless).

**D3 — the continuation-turn lifecycle, durably snapshotted (F105-D3; RESOLVED).** `turns` gains
`predecessor_turn_id` + an **immutable `resume_snapshot`** holding the **canonical settlement-manifest BYTES** (not
merely its digest — a digest cannot reconstruct the bytes, and m-10 never recomputes) **+ the log path**, persisted
in the continuation-admission transaction; every initial and replayed `turn_open` is **derived from those committed
bytes** (the digest verifies, never replaces them). One carrier: the manifest + path ride **`turn_open`** only. The
**mutable `resume_disposition`** is a SEPARATE column (below), not conflated with the immutable snapshot. **Trigger:**
on replacement of
a turn parked `INTERRUPTED`, m-10 admits a continuation with **`admission_ref` inherited byte-identically**, in one
transaction (turns row + `predecessor` + `resume_snapshot` + active-turn lease, current epoch); `UNIQUE(run_id,
predecessor_turn_id)` (one successor), the chain **bounded by the ratified G-2 counter**.
**Frame-totality (F105-D3-R9) — the sole carrier must be emittable by construction.** The complete candidate
`turn_open` (the settlement-manifest bytes + `admission_ref` + `parked_unknown` + log path + framing overhead) is
**sized BEFORE the continuation transaction**; **≤ `FRAME_MAX` (4 MiB)** → commit + emit. **> `FRAME_MAX` → ONE
durable m-10-owned fail-closed outcome (F105-D3-R10/R11):** **no successor turn, no active-turn lease, no snapshot** is
committed; instead m-10 commits **the run to a single TERMINAL `FAILED` state** carrying the closed
**`resume_frame_overflow`** reason (NOT a parked/nonterminal record — there is exactly one lifecycle result: run
terminal, run GC/restart governed accordingly, **no same-run successor or revival**), with an operator-surface
`stop_reason` + a **manual** `resume_action` (operator-initiated new run only). The terminal-vs-parked choice and the
no-same-run-successor rule are master-level; the exact table column/message encoding stays pair DESIGN. There is NO
half-admitted "degraded continuation"
(a continuation would need a durable row + a legal `turn_open` the frame cannot supply, and m-10 cannot pre-choose
`DEGRADED` since only post-`turn_open` m-9 can), and NO auto re-derive without the manifest (the worker would miss the
determinate/uncertain rows that prevent replay). The **no-successor invariant** is master-level; the exact
record/message encoding is the pairs' under F73. One-carrier stands: the manifest is never silently chunked. The exact bound proof (from the G-2 ancestry depth + per-turn attempt/tool + per-entry size bounds) and the
exact-fit/one-byte-over sizing are the pairs' under F73; the pre-commit branch + no-un-emittable-commit are the
master-level acceptance property. **The total first-action
table** (resolves the retry contradiction — "no automatic provider RESEND" = the interrupted attempt is never
re-sent; the normal loop opening a *new* attempt is not a resend):
- **clean positive resume** (all prior effects `settled_with_content`) → the model loop continues; opening the next
  `attempt_id` (full input) is the normal governed loop, NOT a resend;
- **determinate terminal / no-resume** → the terminal fact is **surfaced + terminalized**, never advanced as
  clean-positive: a tool `NOT_INVOKED_INTEGRITY_FAULT` → `turn_failed`; a provider `denied`/`REJECTED_LOCAL`/
  `transport_failed`/`cancelled` → surfaced with its terminal + `cancel_point`, **no automatic replacement attempt**;
  any retry is an explicit fresh user-requested action;
- **uncertain tool** → surfaced (parked disclosure); no silent re-execute; an informed re-issue is a fresh ticket;
- **uncertain provider** → surfaced; **no auto-resend**; a new `attempt_id` only via the frozen user-requested path;
- **degraded / content_lost** → `resume_action` = re-derive (or operator abandon), operator-visible; no silent resume.
**Disposition — a durable PENDING→RESUMABLE|DEGRADED status with a no-work receipt gate (correcting rev7).** At
admission m-10 CANNOT truthfully choose `resumable`/`degraded` (only m-9 can inspect its private log, after
`turn_open`), so the durable `resume_disposition` column is committed **`PENDING`** at admission and makes **exactly
one idempotent transition** to `RESUMABLE` or `DEGRADED` when m-9, having inspected its log, sends a typed
`report_resume_disposition{turn_id, disposition, resume_action}`; m-10 commits it → the operator-surface
committed-snapshot projection reads it (an undurable report could not survive app-main failure). **no-work gate:**
m-9 MUST wait for m-10's **post-commit disposition receipt** (same receipt on equivalent replay; conflict/stale
handling is the pairs' under F73) **before any provider attempt, tool effect, or conductor verb** in the
continuation; inability to establish even the degraded path → fail-closed run-terminal. **Crash cuts recover to the
same durable state without permitting work:** before admission commit → predecessor stays `INTERRUPTED`, retried
idempotently; after admission before send → re-emit from the immutable snapshot, disposition still `PENDING`;
**crash-before-report / after-report-before-commit / after-commit-before-receipt** → the worker re-derives the report
and re-awaits the receipt (idempotent transition), no work until received; crash during a continuation → the next
continuation links to it.

**Retention/redaction/HARD-SOFT.** Retained per-run, GC'd on run-terminal; m-1 reviews log content + at-rest file for
secret-leak + the K6 exclusion. **HARD (decomposition)** = D1 acceptance properties (i)–(vi) + source-turn identity +
the DESIGN-obligation existence · the D2 **two-stage model** — the m-10 pre-admission **3-class producer-total manifest
evidence** (settled_with_content / determinate_no_resume / uncertain; completed-without-receipt ⇒ uncertain; every
canonical row mapped once; source-turn-scoped; tool-only args_digest) and **`content_lost` as the m-9 post-inspection
reconciliation RESULT** — + the **content-AND-marker-before-outcome ordering** + the **composite provider settlement**
(canonical terminal AND a durable content-ready receipt) + no-auto-resend + the **determinate-terminal first action** ·
the D3 `predecessor` + immutable **`resume_snapshot` (canonical manifest BYTES)** + byte-identical re-emission +
inherited-admission_ref + UNIQUE/chain/G-2 + the total first-action table + **frame-totality with ONE durable
`resume_frame_overflow` fail-closed outcome + the no-successor invariant** + the durable **`PENDING→RESUMABLE|DEGRADED`**
disposition + the **post-commit disposition-receipt no-work gate**. **PAIR-DESIGN (Tier-HARD, F73, not here)** = D1's
record/segment/rotation grammar + the fence mechanism; the exact provider content-ready receipt + disposition-receipt
frame/table/replay mechanics. **SOFT** = model-rendering. *Owners:* **m-9** (D1 log + D2-consume/reconcile + the
content-ready receipt + resume + §7.1 supersession) ⇄ **m-10** (D2 manifest-produce + the composite-settlement gate +
D3 lifecycle + durable snapshot/disposition + the receipt) — a coordinated two-sided seam, joint join record; **m-1**
(redaction).

**E — the model-surface digests, owned where the bytes exist (F104-E).** Two component digests, no owner hashes
bytes it cannot see:
- **`logical_surface_digest`** *(owner m-9)* = SHA-256 over JCS `{instructions, logical_tool_schemas[],
  tool_descriptions[], compaction_template, policy_messages}` — the pre-lowering surface m-9 assembles (its
  `logical_tool_schemas`/`tool_descriptions` component supplied by **m-2**). Rides m-9 → m-10 attempt row / E0.
- **`provider_lowered_tools_digest`** *(owner m-8)* = SHA-256 over the lowered `tools[]` portion of the frozen-core
  body (m-8 alone performs provider lowering, and has the bytes at freeze). Rides the m-8 terminal / attempt record,
  alongside `frozen_core_digest`. m-9 does NOT reproduce m-8's translation.
- **`model_surface_digest`** = SHA-256 over `{logical_surface_digest, provider_lowered_tools_digest}` — a join of the
  two component DIGESTS (not bytes), assembled by **m-3** at the E3 binding. The observer derives each component
  independently (logical from the worker's presented surface, lowered from the observed wire request). LOCK = the two
  component recipes + field sets + producer ownership + carriage + the join recipe.

**Typed E3 predicate ids** (unchanged): `provider_request_matches_frozen_core` · `provider_deny_caused_zero_transport`
· `local_invocation_matches_effect_descriptor` · `relay_record_committed_with_stamped_sender` ·
`no_alternate_credentialed_provider_route_observed` — each `{predicate_id, version, required_inputs[],
observed_facts[], evidence_locator, verdict ∈ pass|fail|unknown, observer_id, exact_digest}`. *Owner:* m-3.

## §6 — The per-interface dependency DAG (F104)
- **B:** m-3 (E0/E3 schema delta) + m-8 (terminal digest) FIRST → m-9-carriage ∥ m-10-row (siblings) → m-3 evaluator
  join. *(design edges, distinct from runtime dataflow.)*
- **C:** m-10 descriptor/ticket → m-9 executor consumer; m-1 env/no-leak review.
- **D:** a coordinated two-sided seam — **m-9** (the D1 crash-safe log + the D2 reconciliation-consume + resume + the
  §7.1 "no-durable-outcome-store" supersession — needs m-9 owner confirmation) ⇄ **m-10** (the D2 settlement-manifest
  producer + the D3 continuation-turn/predecessor/inherited-admission_ref lifecycle) + **m-1** (redaction) — with a
  joint join record. No blob wire, no m-10 store read-back, `record_tool_outcome` unchanged, K6 replay unchanged.
- **E:** m-2 (schema/description component) → m-9 `logical_surface_digest`; m-8 `provider_lowered_tools_digest`
  (independent); → m-3 joins the two digests in E3. No aggregator hashes foreign bytes.
- **A:** extraction recipe + bundle authored last, over settled B–E.
- **m-7 broker study (§8) resolves FIRST**; H-24 before re-lock if cross-epoch completion survives. Each leg:
  byte-exact delta → fresh implementer review → affected-consumer confirmations (F73) → join records for two-sided seams.

## §7 — The Step-3 exit gate: decidable predicates (F106)
Six legs (this supersedes any "four" wording). **Each leg passes only when the bound E3 record is `applicable` AND its
typed predicate `verdict=pass` AND (where the record carries one) `observed_outcome=pass`; any `fail` FAILS the leg;
any `unavailable`/`unknown` HOLDS (unknown).**

| leg | fixture | machine predicate (over structured fields) |
|---|---|---|
| Governance-binding | `xit-gov-1` | `provider_request_matches_frozen_core` ∧ `local_invocation_matches_effect_descriptor` ∧ `relay_record_committed_with_stamped_sender` each applicable+`verdict=pass` |
| Durability | **required sub-fixtures, ALL must pass** (mutation-resistant proof cuts for the new seams, F106-R9). `xit-dur-1` (POSITIVE): crash after ≥1 `fsync`-durable `round_marker`; resume from the **exact last valid prefix**, reconcile positively against the closed manifest union (trust a `tool_result` only under a matching `settled_with_content` present in the prefix), reproduce `resume_prefix_expectation`. `xit-dur-2` (DEGRADED): the fixed `corruption_cut` → degrade to re-derive + durable `degraded` disposition/`resume_action` on the operator surface. `xit-dur-3` (PROVIDER CONJUNCTION + no omission, BOTH missing-half orders): terminal-first/receipt-absent → **exactly one `uncertain`** entry for the completed row (an OMISSION mutant FAILS); **receipt-first/terminal-absent → the extant canonically-UNKNOWN/PARTIAL provider row must ALSO appear as exactly one `uncertain`** entry (never settled, and an OMISSION mutant here FAILS too); both (either order) → **exactly one** `settled_with_content`, duplicate reports idempotent; **separately**, positive-evidence-then-missing-prefix → m-9 `content_lost` + durable `DEGRADED`. `xit-dur-4` (RECEIPT GATE, both directions, OBSERVED action): crash before-report / after-report-before-commit / after-commit-before-receipt → **zero provider/tool/conductor work** until the post-commit receipt; AND, binding **one selected first-action branch**, its corresponding durable/wire action is **OBSERVED exactly once AFTER the receipt with zero such observations before it** (mere reachability is insufficient — a worker that exposes a runnable transition but never executes it, or holds forever, FAILS). `xit-dur-5` (FRAME BOUNDARY): the max legal frame passes; **one byte over** → the single `resume_frame_overflow` outcome — assert the run committed to **TERMINAL `FAILED`** with the closed `resume_frame_overflow` reason, **no successor turn / no lease / no snapshot / no same-run revival**, and the operator manual-`resume_action` projection (never a committed un-emittable frame, never a parked/nonterminal record). A degraded re-derivation never satisfies `xit-dur-1`. FAIL on any fabricated/duplicate/omitted settled effect, silent resume across an uncertain/determinate terminal, a permanent post-receipt hold, or a committed un-emittable frame |
| Crash-honesty | `xit-crash-1` | the fixture crashes at a fixed fault point (after the first effect is externally visible, before outcome commit). **(a)** the crashed call's `tool_calls` row parks `UNKNOWN_TOOL_OUTCOME` and STAYS parked (never `EXECUTED`); **(b)** the fixture-owned external observer records `{counter_before_recovery, counter_after_recovery, invocations_after_recovery}` and the leg passes only when **`counter_after_recovery == counter_before_recovery` AND `invocations_after_recovery == 0`** (recovery caused no second invocation and no second effect; for this fault point both counters == 1). **The MVP does NOT claim F59 semantic-effect dedup** (a fresh ticket/`tool_call_id` may be a semantic duplicate; ambient bash re-runs); any explicit *informed* retry runs in a **separate** fixture/result under a fresh ticket. The expected counter relation is in the frozen fixture manifest |
| Injection-visibility | `xit-inj-1` | the induced action has an F59 ticket + a `FROM`-stamped record + an honest recorded outcome (visibility, NOT prevention) |
| Governed handoff | `xit-ho-1` | **two correctly-stamped records + lineage:** (1) the handoff relay committed with channel-stamped `FROM` = originating seat; (2) the second seat's response/ack committed with `FROM` = second seat AND a parent/reference to record (1)'s `relay_id`. Channel-stamping puts each authoring seat on its own record — the origin `FROM` is never forged onto a second-seat record. The composite proof binds the conductor E1/E2 evidence for both |
| Operability | `xit-op-1` | the operator surface exposes `{last_event, stop_reason, unknown_effects[], resume_action}` all present + non-null |

**Frozen fixture manifest** `master/STEP-3-EXIT-FIXTURES.json` (hashed at re-lock; fixtures BUILT at T4 to this frozen
spec — not post-selectable): per fixture `{fixture_id, input_artifact_sha256, fault_injection_point,
expected_canonical_rows, observer_id, evidence_locator}` **plus** `effect_observer_key` **and its bound**
`effect_counter_expectation{counter_before_recovery:1, counter_after_recovery:1, invocations_after_recovery:0}` (for
`xit-crash-1`, tied to its `fault_injection_point` — the expected relation lives in the hashed manifest, not prose),
`handoff_expected_records[2]` (the two stamped-lineage records for `xit-ho-1`), `resume_prefix_expectation`
(`xit-dur-1` — the digest vector `{predecessor_turn_id, resumed_round_index, log_prefix_digest, context_digest}` the
positive resume must reproduce) and `degraded_expectation` (`xit-dur-2` — `{corruption_cut, expected_disposition:
"degraded", expected_resume_action}`, bound to its `fault_injection_point`), and per-fixture
`sample_weight` values that **sum to EXACTLY 30 governed turns + 100 tool calls**; a top-level
`{baseline_artifact_digest, baseline_config_digest}` names the frozen baseline. **Objective overhead budget** (numbers
master-proposed, operator-ratifiable at the gate; immutable for T4 once ratified): per-metric p95 ceilings — F59
authorize→consume **≤ 250 ms**, relay round-trip **≤ 1000 ms**, session-log append **≤ 50 ms** (worker-local file, each
pass/fail); the **total added-wall-clock rule** over baseline = **p50 ≤ 20 % PASS · 20 % < p50 ≤ 100 % HOLD · p50 >
100 % FAIL** (no undefined interval); the **HOLD band is cleared only by a durable operator `HUMAN_GATE` relay citing
the measured p50 + the threshold** (never a chat-only bypass). Evidence per leg `{fixture_id, run_id,
evidence_locators[], verdict, measured_metrics{}}`.

**Gate vs demonstration (D3/D4/D5):** the six legs + the overhead budget GATE Step-3 close; utility is demonstrated
OFF the gate (public dogfood + the honestly-labeled agent-as-operator SWE-bench). **The gate proves governance
correct + cheap, NOT real-work useful** — usefulness rides on the non-gated public artifacts. dogfood/real-work start
⊥ exit. Metrics per H-25.

## §8 — The m-7 broker study + the H-24 conditional gate (unchanged)
Focused m-7 study resolves before the affected m-10/m-9 finals + the re-lock: broker-survives-app-main-crash +
adoption + cross-epoch completion vs the simpler kill-all-children/rebind/no-op-crosses-epoch/bounded-drain/
unresolved→UNKNOWN/rediscovery+dedup rule. **Keep the separate secret-holding process (F67).** Any retained
cross-epoch completion ⇒ **H-24 (bounded TLA+/Alloy) before the re-lock.**

## §9 — H-16 + H-26 before T4 (unchanged)
H-16 PLAN → RED-first IMPL → merge grant; H-26 converges under explicit operator plan scope; both before T4.

## §10 — H-12 promotion + deployment envelope + prohibited-use classifier (unchanged from rev3)
H-12 PROMOTED (this amendment's precedence, recorded here + in `FRANK-HARDENING-BACKLOG.md`) from Step-4 residual to
a **HARD pre-external-use gate**. **Allowed envelope (MVP):** operator-driven · local host · single-tenant ·
operator-trusted input. **Prohibited until H-12:** untrusted input (third-party repos / external tickets / un-vetted
web content), network-reachable/service invocation, multi-tenant use. **Fail-closed classifier** gates any
out-of-envelope run. The public dogfood does NOT bypass this — building CRM/bivpak locally is trusted-local use;
publishing the OUTPUT artifacts is not frank exposure (the boundary is on frank's INPUT/exposure, never on publishing
outputs). **Gate owner = operator; artifact = a release-binding precondition asserting `h12_satisfied`** before any
out-of-envelope deployment. None is authorized by this amendment.

## §11 — Revised sequence
1. This amendment → **VP decomposition review r12** → **operator re-scope ratification** (supersedes `b7e1f0ef`;
   ratifies the §3 grill + the §7 overhead numbers + the §10 envelope).
2. **m-7 broker study first** (+ H-24 if cross-epoch completion survives).
3. **Interface DAG legs (§6)** via the F73 ladder + join records for the two-sided seams.
4. **Author the extraction recipe + bundle (§4)** over settled interfaces; ship the `bundle-soft-stability` negative
   fixture; F73-review it. Freeze `STEP-3-EXIT-FIXTURES.json`.
5. **Re-run the shorter stage-6 lock** over `bundle_sha256` + the whole-file-hard owner contracts (unchanged design
   hashes carry forward where untouched).
6. **T4** (behind the re-lock + H-16/H-26). Once frank is usable: dogfood/real-work + the property-leg/overhead exit
   test + the SWE-bench credibility run in parallel (D5); step close on the gate; out-of-envelope use stays behind §10.

## §12 — Gate this amendment needs
- **VP:** decomposition review r12 — does `bundle_sha256` exclude source provenance + stay stable under soft edits;
  does §5-C bind the actual env/interpreter + fix apply_patch; is §5-E executable where the bytes exist; is §5-D
  crash-total with no silent-resume gap; are the §7 legs decidable (applicability AND verdict) with frozen fixtures +
  a total overhead rule; does anything reopen an approved mechanism or move a bound byte?
- **Operator:** the re-scope ratification — the §3 grill, the §7 overhead numbers (ratify/adjust), the §10 envelope,
  the bounded interface amendment; the all-artifact lock held + superseded. Recorded agent-authored + operator-cited
  per §8b; never a forged `FROM: operator`.
DOES NOT issue: PLAN, T4 token, credentials, provider calls, release binding, live E3, merge, deploy, or
out-of-envelope use — all held.
