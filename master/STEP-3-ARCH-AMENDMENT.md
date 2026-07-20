# STEP-3 ARCHITECTURE-AMENDMENT PACKET (r4) — the conductor is one service in a larger app shell

**Status: r4 CANDIDATE — folds VP r3-review F15–F17 (`step3-arch-packet/053000`), bounded source reconciliation (no product fork, no grill rerun); operator grill remains CLOSED (§9 `GRILL_LOCK: step3-arch-reframe-grill`, 2026-07-15); awaiting VP exact-candidate re-review, then operator hash-bound ratification.** The candidate SHA-256 of this finalized file is reported in the transmittal `step3-arch-packet/DESIGN-orchestrator-planner-20260715-060000.md`. Standalone packet **outside the locked kickoff bytes**.
Review lineage: `step3-arch-reframe/011000 → 013000 → 020000 → 023000 → 024000 → step3-arch-packet/030000(must-revise) → 033000(input-record approved 034000) → r2 040000 → 043000(must-revise F11–F14) → r3 050000 → 053000(must-revise F15–F17) → THIS r4`.
Hashes: locked kickoff `983508fcbec5a511e343c731e78399af98f264a7110f75965395f940ab3fbd43` (**historical lock, preserved**) · ROADMAP provisional non-operative baseline `3977c9f2e1a0fd57903b5381031c337e4a8d25fe7b1d26de5b06f2eaec58c7c3` (frozen) · packet r1 `818c3d87…`, r2 `2cd16311…`, r3 `8a6154e3…` (all must-revise, superseded by this r4) · **candidate SHA-256 recomputed on this finalized r4 file, reported in the transmittal.**
**Supersession of the reframed kickoff sections is effective ONLY on operator ratification of the VP-approved exact hash.** Historical relays are append-only (referenced by supersession, never edited). External hub-and-spoke input is **non-authoritative**.
Source-action honesty: the master.orchestrator-planner session authored the ROADMAP reframe fold + the four Step-2 currency fixes (disclosed, `step3-arch-reframe/033000`); the whole roadmap topology diff folds atomically at ratification, not line-by-line now.

---

## 0. Ratified inputs (operator, 2026-07-15) — fixed points, not reopened
- **R-Q1** conductor = governed **relay plane for stamped participants** (agent seats, orchestrator seats, the operator channel, reserved system-authored governance records) + own isolated store + sole governed writer; NOT app supervisor/run-DB/provider-client/turn-engine/tool-broker/terminal-multiplexer/general-IPC. Steps 1–2 unchanged.
- **R-Q2** Step-3 = the **MVP**: one honest governed turn end-to-end live (E3), **single pinned lane, no routing execution**. Step 4+ = ship.
- **R-Q3** new domain **m-10 App Control Plane / Supervisor** (hosts + sequences; owns **no** policy).
- **R-Q4** conductor is its own isolated service; app state stays out, **per-family owners/writers, no cross-domain writer**.
- **Operator direct path** retained as a **named non-governed route** (contract in §8b).
- **VP fork carries:** provider-send mechanism app-side; **policy stays m-3**; m-8/m-10 enforce last pre-wire; m-1 secrets; conductor records a decision/evidence summary, never request/stream/credential.

Legend: **[RATIFIED]** · **[VP-RULED]** fixed by 030000 findings · **[GRILL]** operator-owned, resolved in §9 (recommendation carried, not pre-closed).

---

## 1. Boundary matrix
| Component | Owner | Is it a conductor SEAT? | Process boundary | API / IPC | Canonical state | Writer | Secrets | Authority / gates | Evidence |
|---|---|---|---|---|---|---|---|---|---|
| **Conductor — governed relay plane** | m-1..m-7 (m-7 hosts) | — (it IS the plane) | own isolated service [RATIFIED] | `submit`/`project`/`read`, stamped participants only | conductor relay store | m-7 sole governed writer | none | the locked relay gates (m-2 form/lint+lineage · m-3 observe-as-send · m-1 stamp · m-6 park/wake+ODB) | E0–E4 over **relay** traffic |
| **m-9 Model Runtime — the WORKER** | m-9 | **YES — the only app component that is a seat** [VP-RULED F4] | app-side worker process, supervised by m-10 | app IPC to m-10 + m-8; its **own private seat channel** to the conductor for genuine seat↔seat relays | m-9 turn/session/context state | m-9 | **none** | parses tool calls → INERT until the **app-side authority point (m-10 enforces m-5 ceiling)** authorizes | the worker seat is the **only submitter** of an E0-labeled run summary |
| **m-10 App Control Plane / Supervisor** (NEW) | m-10 | **NO — trusted app component, not a seat, no `submit` credential** [VP-RULED F4] | app-side process (the hub) | app IPC to workers + connector | app **run manifest** + supervisor + scheduler state + active-turn lease | m-10 | orchestrates only **opaque credential references**, never secret bytes [VP-RULED F6] | hosts the app-side authority enforcement point (m-5 ceiling); owns **no** policy; **supervises without authoring as the worker** | emits app events; not a conductor principal |
| **m-8 Provider Adapters / connector** | m-8 | **NO — trusted app component, not a seat** [VP-RULED F4] | **separate trusted process from the m-9 worker BEFORE the first E3** (same host OK; **same credential-readable address space NOT**) [VP-RULED F6] | receives `LLMRequest` over app IPC; provider HTTPS out via a frank-owned client | connector provider-attempt / stream telemetry | m-8 | **credential runtime holder + secret-store reader/writer under an m-1-authored boundary** | the **last pre-wire enforcement host**; the **`freeze → authorize → attach → send`** locus | the **attestation source**; returns app events over IPC, **does not stamp a conductor record** |
| **Human / terminal surface** | m-6 (governance semantics) + m-10 (PTY/TUI, Step-4+) | m-6 gates via conductor; terminal app-side | conductor-side gate→bucket/ODB; PTY app-side | m-6 surfaces via the relay plane; terminal via app IPC | m-6 HUMAN_GATE fields (relay store); PTY state app-side | m-7 / m-10 | none | m-6 HUMAN_GATE; **operator direct route = §8b** | conductor-side gate evidence |

---

## 2. Traffic matrix
**Governed relay traffic — transits the conductor (unchanged):** seat↔seat relays · operator channel (receiving) · reserved system-authored governance records · a run's **E0-labeled governance/evidence summary** submitted by the worker seat · gate/park/wake/ODB (m-6). **No routing DECISION record and no lane-bearing FieldSpec row is added in Step-3** [VP-RULED F3].

**Negative routes — NEVER the conductor:** provider request bytes + response stream · credentials/secrets · `LLMRequest` · tool-execution request+result payloads · PTY/terminal streams · model-turn traffic · **the app run manifest + run/session/turn/attempt state** · worker lifecycle/supervision. A violation is the category error this reframe corrects.

**Operator direct route** — a named non-governed route (contract §8b): its residual + non-guarantees appear beside every exclusivity-shaped claim.

---

## 3. State-and-recovery matrix (closes VP F3/F8)
| State family | Owner | Canonical store | Writer | Crash / disagreement | Conductor authority over it |
|---|---|---|---|---|---|
| Conductor relay records + projections | m-1/m-7 | conductor store | m-7 sole governed writer | existing crash-atomic commit + recovery | **authoritative** (the record IS truth) |
| App **run manifest** + supervisor + **active-turn lease** | m-10 | m-10 store | m-10 | **fail-closed to interrupted/held; NEVER automatic provider resend**; m-10 starts a replacement only after the prior worker/attempt is proven terminal or on explicit operator disposition | evidence, not authoritative |
| m-9 turn / session / context | m-9 | m-9 store | m-9 | m-9 **one-active-turn** invariant (distinct from m-10's lease) | evidence about a turn |
| m-8 provider-attempt / telemetry | m-8 | m-8 store | m-8 | attempt state bound to `attempt_id`; no auto-resend | **attestation** source, E0 unless corroborated |
| Credential references + secret material | m-1 boundary / m-8 runtime holder | out-of-band secret store; **never** conductor store | m-8 under m-1 boundary | rotation without genesis break | never in the conductor store |
| Terminal / PTY (Step-4+) | m-10 | app-side | m-10 | app-side | evidence surface only |

**Stable identity across the three app stores:** `run_id`, `turn_id`, `request_id`, `attempt_id`; **no cross-store atomicity is claimed.** **Attestation-authority rule (m-3):** a conductor relay recording an app-side send is **evidence, attestation-graded — never authoritative app state, never the payload**; an app-side send report floors at **E0 `self_reported`** and **never promotes** uncorroborated (grounded: `frank/internal/record/record.go:11-13` `{accepted,rejected,held}`).

### 3a. The E0 app-attestation CARRIER (VP F13 — the no-conductor-change branch, chosen)
The landed schema does not express an app/connector attestation provenance: `attestation_source` is a **system-owned enum of only `{conductor, operator}`** and the observe gate stamps `attestation_source=conductor` + computes `achieved_evidence`/`record_integrity` for the **relay observation**, not for an embedded connector event (`frank/internal/fieldspec/registry.json:72-78,147-155` · `frank/internal/observe/gate.go:178-257`). So the E0 app summary **must not** ride a new `attestation_source` value — that would be a conductor member change (the rejected typed-provenance branch). Chosen carrier, **no conductor change:**
- **Exact carrier envelope (VP F15):** the m-9 worker seat's existing **`PHASE: SITREP`, `AUTHORITY: report-only`, `HUMAN_GATE_REQUIRED: no`** relay — the landed *tested non-authority shape*. The classifier returns authority-bearing for any `grant` / human gate / A-category gate / `PLAN`·`IMPL`·`REVIEW-FOLD`·`MERGE-GATE`·`LIVE-VERIFY` phase / impl·merge·live·fold authority / non-SITREP-non-RECONCILE orchestrator-planner record (`frank/internal/lineage/lineage.go:39-58`); a plain worker `SITREP` is non-authority (`frank/internal/lineage/lineage_test.go:14-30`). The carrier therefore has **no `grant`, no `gate_category`, no gate-resolution/disposition field, and no design/plan/merge/live authority field.** Routing: **`TO: master.orchestrator-planner`, `CC: m-3.planner`** (+ the audit reviewer if separately seated). No new relay kind, no new FieldSpec row, no trusted-observer input.
- **Evidence separation — top-level vs body (VP F15):** the relay's mandatory top-level `EVIDENCE_TARGET` and the conductor-produced `achieved_evidence`/`record_integrity` concern **carriage + observable relay claims ONLY**. The **namespaced body event separately carries explicit `event_evidence=E0` + `event_integrity=self_reported`** (exact body-schema field names are m-3's downstream design). So an **E1/E2 top-level relay observation cannot upgrade the embedded event** — the app event stays E0 regardless of the relay's own carriage evidence. The m-3-owned app-event schema is serialized in the body.
- **Mechanically prevented from becoming authority:** the body is **not gate-referenceable** (no gate reads relay-body app-event fields), the relay carries **no typed `grant` and no gate resolution**, and the embedded event stays **E0/self-reported**. **Reader/consumer:** master + the audit/evidence review read it as E0 app evidence; **no gate, promotion, or authority derives from it.**

---

## 4. End-to-end sequences

**Sequence A — the app-side governed turn (MVP; no routing, no conductor change):**
1. m-10 writes the **app-side pinned run manifest** (run state) binding **one immutable m-8 lane ID + catalog digest**. **This is not an m-4 routing decision and performs no routing** [VP-RULED F3]. Lane identity in the manifest is **never a conductor gate input**.
2. m-9 (the worker seat) assembles the turn; a parsed tool call is **INERT** until the **app-side authority point authorizes it — m-10 enforces the m-5-authored ceiling artifact bound to `run_id` + worker identity, failing closed if absent/stale** (§8 m-5 amendment). An above-ceiling call has **zero execution** (E2 negative).
3. m-9 → m-8: `LLMRequest` over app IPC. **No keys in m-9.**
4. m-8 performs the exact **`freeze → authorize → attach → send`** contract [VP-RULED F6]: freeze one **immutable authorization identity** over method + canonical endpoint + non-auth headers + body → authorize (apply m-3 policy at the last pre-wire point; **denial ⇒ zero socket send, no post-authorization mutation, no secret exposed to m-3**) → attach the m-1-governed credential (opaque reference from m-10; secret bytes only inside m-8) → send via the frank-owned client. **Zero conductor transit.** **MVP default: one attempt, no automatic retry** (a retry branch is a separate owner contract; C8 no-replay ≠ exactly-once send).
5. m-8 returns the normalized stream + an **attestation**. **m-3 owns the app-event/attestation schema; m-8 emits the app event; the m-9 worker seat carries it up** — via the **E0 carrier of §3a** (an existing ordinary non-authority relay, the m-3 app-event serialized in its **body**), which **cannot satisfy a gate or promote evidence.**
6. Recovery/cancellation app-side (§3); a governed human gate is opened/read **by the worker seat through the existing conductor verbs** — no new m-10 conductor address or wake API.

**Sequence B — seat → conductor → recipient seat (governed relay, UNCHANGED from Steps 1–2):** `submit()` → m-7 stamps FROM → m-2 form/lint + m-3 observe-as-send + lineage + single-writer commit → recipient `project()`/`read()`.

**"Honest governed turn" = the app enforcing the locked owner policies with correctly-labeled proof — NOT every app event becoming conductor-observed.** Proof set: deterministic **E2 negatives** (policy-deny→zero send · no post-auth mutation · no secret leak · above-ceiling tool→zero execution · no provider bytes in any conductor surface) **+ one E3 live provider turn** via a separate integration harness / operator observation. **E3 is not laundered into the conductor summary** absent a future trusted-observer contract.

---

## 5. Scheduler split (VP F8)
Conductor governance-gate scheduler (m-6: park/wake exactly-once + ODB) — unchanged. App scheduler (m-10: worker scheduling, provider-await, cancellation, backpressure) — new, app-side. **The bridge reuses the worker seat's existing `submit`/`project`/`read` path for any governed gate; it is NOT a new conductor event or m-10 principal** [VP-RULED F8].

---

## 6. Compatibility proof — Steps 1–2 valid, no conductor byte changes
- `submit`/`project`/`read` unchanged; only the **m-9 worker is a seat**; m-10/m-8 hold **no `submit` credential**; the I-PH guardrail holds (verified: `frank/internal/channel/server.go:391-415`, `cmd/frank-mcp/mcp.go:129-155` expose exactly the three verbs).
- **No new relay kind / FieldSpec row / trusted-observer input** is required for the MVP: the pinned lane is an app-side manifest, not an m-4 record (live `routing_assignments` keeps only `chosen_model`, `frank/internal/fieldspec/registry.json:175`); and the **E0 app summary rides an existing ordinary relay body (§3a), not a new `attestation_source` value or carrier**. Therefore **"no conductor member changes" is true** — and if any future path needs one (e.g., typed app-attestation provenance), it is a **separately-flagged owner amendment**, not free compatibility.
- Operator verdicts + system records unchanged; the direct route is **additive** (§8b), altering no governed mechanism. The ten INV-CATALOG laws, `{accepted,rejected,held}`, exactly-once park/wake operate on the untouched relay store.

---

## 7. Disposition table — the five held lanes
| Lane | Disposition | Artifacts |
|---|---|---|
| `step3-design-m-8` | **RE-DISPATCH** (app-side connector) | rev3 audit **SALVAGE** (matrix, ~40 fixtures, 4-axis **schema**, owned-HTTP posture); discard conductor-side placement. |
| `step3-design-m-9` | **RE-DISPATCH** (app-side worker) | draft r0 **SALVAGE** (§1 machine/typed-terminals/no-fourth-token law, C7/C8/C9, two-point observe, compaction, Q4); **§2/§3/§6.4/§7 host-bindings re-point off m-7**; the "runs ON m-7" sentence carried by **append-only supersession refs** (m-9 boot relay, README:13, kickoff §2, audit rows) — **not edited** [VP-RULED F2]. |
| `step3-amend-m3-egress` | **RE-DISPATCH** (m-3 keeps policy; app-side enforcement + attestation contract) | nothing authored; do-not-inherit `egress/rules.go:22-43`. |
| `step3-amend-m7-cred` | **RE-OWNER** (credential contract → connector-side owner, m-1 boundary) | r3 doc = **reviewed-but-still-`must-revise` provisional audit input** (per m-7's corrected handoff `step3-hold-m7/…-020423`): **F1–F6 confirmed, F8 closed, F7/F9/F10 directions accepted, F11–F13 OPEN**; no r4/lock. **F12** (freeze binds endpoint-only, not the full authorized freeze) and **F11** (catalog-v2 activation vs drift law) transfer as **OPEN defects** the fresh owner consumes as findings, not accepted design. m-7 conductor-host scope untouched. |
| `step3-amend-m4-routing` | **DEFER to the named Step-4 routing-execution gate** [VP-RULED F3] — **not a Step-3 writer.** m-4 may **consumer-review** the manifest boundary now. | Q1 `lane_ref`/digest/tuple + fingerprint = **provisional old-dispatch grill state**, salvage only, not locked representation. |

---

## 8. Domain/charter delta · propagation · dependency graph
**m-10 App Control Plane / Supervisor (NEW).** Owns: worker lifecycle+supervision · app IPC+backpressure · run/session-state persistence orchestration · the app scheduler + active-turn lease · the app-side authority enforcement point (enforces m-5's ceiling, does not re-own it) · connector supervision + **opaque** credential-reference orchestration. **Owns NO policy.**
**[VP F10 — seats + owner, carried whole]:** **m-10 stands up a `.planner` + adversarial `.implementer` pair**; **m-10's component-boundary design + adversarial review must complete (or be interface-locked ahead) BEFORE any m-8/m-9 consumer lock.** **Exactly one connector-credential-contract owner: m-8.planner authors it, m-1 reviews the secret boundary, m-10 + m-3 consume** (unless the §9 grill resolves otherwise).

**Charter deltas:** m-8 → app-side connector (creds + owned wire + last-pre-wire `freeze→authorize→attach→send`) · m-9 → app-side worker (strike "runs ON m-7's substrate", via supersession ref) · m-3 → provider-send policy/evidence + attestation contract · m-4/m-2 → Step-4 (Step-3 = app-side manifest) · **m-5 → stays SOLE policy owner; the ceiling *enforcement host* moves to m-10 via the owner amendment below (m-5 does NOT lose policy ownership)** · m-7 → retains conductor-host, **loses** the provider-credential contract, **no engine-v5 `provider_bindings` credential member** · m-1 → secret/provenance boundary extends to connector attach · m-6 → gate semantics stay conductor-side.

**[VP F12] m-5 ceiling-host amendment — required (m-5's locked design names conductor/host-config enforcement, not m-10).** Relocating the ceiling **enforcement host** to m-10 is a real locked-boundary amendment, not a charter text substitution. **m-5.planner authors it, m-5.implementer adversarially reviews it, and it interface-locks BEFORE any m-10/m-9 consumer lock.** It preserves **m-5 as sole policy owner; m-10 as enforcement host only**, and pins the **ceiling artifact interface**: source, writer, schema/config home, **immutable binding to `run_id` + worker identity**, the m-10 read/load path, and **fail-closed behavior when the artifact is absent or stale** (no unbounded execution). Until this amendment interface-locks, m-10 cannot be a consumable enforcement host.

**Propagation list (fold ONLY on ratification, atomically, master-authored) — STAGED for m-5 (VP F16):** `ROADMAP.md` (whole topology diff incl. stale tech-stack/PTY/interjection clauses — frozen `3977c9f2…`) · `CLAUDE.md`(+`AGENTS.md`) · `master/ARCHITECTURE.md` (the matrices land here) · `master/README.md` · `master/STEP-3-KICKOFF.md` (reframed sections marked superseded, old hash `983508fc…` preserved) · `master/CYCLE-PLAYBOOK.md` Part F · `master/RECONCILE.md` · the m-3/m-4/m-6/m-7/m-8/m-9 charters **+ new m-10 charter** · append-only supersession refs for the "m-9 runs ON m-7" surfaces. **m-5 is STAGED, NOT a silent rewrite:** the immediate ratification fold records the new topology **plus the PENDING/non-consumable m-5 ceiling-host amendment gate** (it does **not** rewrite the locked m-5 design, which still names conductor/host-config enforcement); the replacement flow then **creates the m-5.planner/-implementer amendment and records its approved supersession / design-of-record fold BEFORE m-10/m-9 lock**. Both stages named.

**Replacement dependency graph (VP F16 — coordinated first stage):** (1) **COORDINATED FIRST STAGE:** the **m-10 charter + boundary design + adversarial review** AND the **m-5 ceiling-host amendment** (m-5.planner authors, m-5.implementer reviews) proceed together and **interface-lock their shared ceiling contract** — both before any m-8/m-9 consumer lock (m-10 consumes the m-5 ceiling interface, so neither locks a consumer ahead of it) → (2) **m-8 connector + m-9 worker** design against the m-10 + m-5 ceiling interfaces + the conductor dependency → (3) the **remaining** amendments (m-3 egress+attestation · connector-credential [m-8-owned] · — m-4 deferred; **m-5 already interface-locked in stage 1**) author in parallel (B14), consume m-8/m-9/m-10 boundaries before final review → (4) Master+VP reconcile+lock → (5) the **MVP vertical**: one governed turn, single pinned lane, live E3.

## 8b. Direct-operator-route contract (VP 023000 F3 / 030000 F5, **REFINED by the operator grill G8, 2026-07-15**) — a trusted authority-bearing channel by construction, not a D5 accident
**Governing principle (operator-directed — flagged to the VP as a relaxation of the drafted item 4):** under **confusion-not-malice**, the live interactive operator channel is a trusted authenticator **by construction** — a confused, non-adversarial agent cannot fabricate an interactive human presence; operator impersonation is *malice*, out of frank's threat scope. Forcing the operator to author a governed relay to "prove" authority is **adversarial-shaped ceremony and is CUT**. The route is therefore **authority-bearing** (carries operator decisions AND authorizations), and the operator is **never** required to personally author a governed relay.
1. **Endpoints/non-transitivity:** operator-to-ONE-agent direct interaction only; **never** agent↔agent carriage, generic app IPC, connector/provider transport, credential injection, or a fallback when conductor delivery fails.
2. **Ingress authenticity by construction:** the live interactive channel itself identifies the operator — no heavyweight auth mechanism required (confusion-not-malice); a seat/worker cannot fabricate an interactive operator presence.
3. **Authority + recording — NON-TRANSITIVE (VP F11):** a direct instruction authorizes **only the directly-addressed recipient**, and only an action **within that recipient's already-bound capability/authority ceiling**. The recipient may record the instruction + effect under its **own `FROM`**, but that record is an **evidence/audit record — NOT an operator-stamped grant, NOT transferable authority** (see item 4).
   - **Direct app-side action vs conductor-governed action (VP F17):** a direct instruction alone MAY authorize a **directly-addressed app-side action within the recipient's current ceiling**. But when a **conductor-governed action requires a typed `grant` / lineage edge** (merge, impl dispatch, any authority-bearing record), the **sanctioned typed-grant branch is MANDATORY** — the direct message is *context*, never a substitute for the required accepted grant.
   - **Cross-seat effects use the LANDED grantor grammar, not a citation:** a *different* seat acts only after either **(a)** the operator directly instructs that eventual actor (still authoring no relay), or **(b)** a **protocol-sanctioned grantor emits the EXISTING typed grant under its OWN stamped authority, citing the operator instruction as context** — per the registry, `dispatch-merge` is grantable only by `operator` / `*.orchestrator-planner`, `dispatch-impl` by a pair planner, others none (`frank/internal/fieldspec/registry.json:105,111`). **No arbitrary agent mints authority by citing the operator.** (This is exactly how a merge dispatch already works: the orchestrator-planner emits the grant under its own sanctioned authority, citing the operator's instruction — not a fabricated operator grant.)
4. **Integrity floor (what stays hard):** (a) **"by construction" is scoped to LIVE-INGRESS authentication ONLY** — it does NOT extend to the later agent-authored citation; `record-never-fabricate` is a behavioral norm, and a confused recipient could misquote scope/conditions without impersonating the operator, so the citation is **E0 self-reported evidence of the instruction, not by-construction operator provenance**; (b) the out-of-band message mints **no forged `FROM: operator`** and does not silently mutate the store — any governed effect is the acting seat's **own stamped grant/relay within its own authority**; (c) a **live-runtime credentialed leg** against the RUNNING conductor (§7 config-change / seat-mint) stays **operator-performed**.
5. **Ceiling rule (VP F11.3):** a direct instruction **cannot** silently or textually raise m-5's immutable spawn/run ceiling; **any ceiling change uses the m-5-owned typed reconfiguration/respawn contract + its gate** (§8, the m-5 ceiling-host amendment) — otherwise the addressed recipient remains bounded by its current ceiling.
6. **Evidence/labels:** app-side transcript/writer/retention; direct-route content labeled **direct/operator/app-side**; a governed relay authored on the operator's instruction carries the **agent's FROM + the operator-authorization citation**, never a forged operator stamp.
7. **Negative proofs:** a worker cannot present its own input as direct-operator input; direct content becomes a cross-seat message only via a **fresh conductor submission that itself remains non-authority unless it independently satisfies the landed grant grammar** (VP F17 — direct content carried into the store does not inherit authority); route failure does not fall back to an unlabelled channel; no secret/provider bytes enter conductor artifacts via the route.

## 8c. Carry ledger (VP 023000 F5 / 030000 F10) — every old commitment mapped
| Commitment | Disposition |
|---|---|
| V1 one governed turn | **RETAINED — the Step-3 MVP** |
| V2 two-provider portability | **DEFER → Step-4 gate** |
| V3 routing execution | **DEFER → Step-4 routing-execution gate** |
| m-4/m-2 routing-record amendment | **DEFER → Step-4** (Step-3 uses the app-side manifest; m-4 consumer-reviews only) |
| Benchmark | **DEFER → Step-4** |
| Native governed agent-spawn | **DEFER → Step-4+** (powers Step-5 recursion) |
| Steer / interrupt / side-question | **DEFER → Step-4** (Step-3 unlocks the runtime; first-class delivery Step-4) |
| Provider credentials/config | **REPLACED** → connector-side contract (m-8-owned, m-1 boundary); no conductor credential member |
| Step-2 T5/T10 | **RETAINED as carries** (acceptance-OPEN behind g2/dc) — unaffected by the reframe |
| item-2 soft-expiry arbiter cell (m-7+m-3) | **RETAINED as a carry** |
| T4 team mechanics / live relaunch | **RETAINED** — precondition to the first T4 code token, unchanged |
| old consumer-lock seams (m-8↔m-1/m-3/m-7 · m-8↔m-9 · m-8↔m-4 · m-9↔m-5/m-7) | **RE-CUT** against the app-shell boundaries here; m-8↔m-4 becomes a Step-4 seam |
| Steps 4 / 5 / 6 | **PRESERVED as distinct observable gates** (not collapsed into one shipping phase) |

---

## 9. GRILL_LOCK (CLOSED — operator grill run one-question-at-a-time 2026-07-15)
```
GRILL_LOCK_ID: step3-arch-reframe-grill
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relays read: the review chain 011000 → 013000 → 020000 → 023000 → 024000 → packet-030000 → 033000 → 034000; the five held-lane handoffs; m-7's corrected handoff 020423
- code/docs inspected: record.go:11-13 · egress/rules.go:22-43 · fieldspec/registry.json:175 · channel/server.go:391-415 · cmd/frank-mcp/mcp.go:129-155 · the m-1 operator/seat/system identity contract
- questions answered from codebase / VP rulings: G1, G2, G3, G5, G6, G9, G10-retry
- questions asked operator (one at a time): G4, G8 (detailed); G7, E3-gate (confirmations)

Resolved decisions:
- G1 conductor boundary — confirmed: stamped-participant governed relay plane (R-Q1) — source operator/VP
- G2 connector process split before E3 — separate trusted process from the worker; same host OK, NOT same credential-readable address space — source VP-ruled F6
- G3 attestation + proof split — connector send report = E0/self_reported app summary; the real turn proven at E3 via a separate harness/operator observation, never laundered into a conductor summary — source VP-ruled F7
- G4 Step-3 lane representation — **app-side pinned run manifest (m-10 run state, immutable m-8 lane ID + catalog digest); the governed m-4 routing-record + its conductor amendment DEFER to Step-4** — source operator (detailed grill)
- G5 state writers/lease/recovery — per §3: per-family writers, stable run/turn/request/attempt IDs, no cross-store atomicity, fail-closed to interrupted/held, two distinct lease invariants — source VP-ruled F8
- G6 scheduler bridge — reuses the worker seat's existing submit/project/read verbs; NO new conductor event or m-10 principal — source VP-ruled F8
- G7 credential/config residue — **NO conductor provider-credential member (engine-v5 `provider_bindings` dropped); the lane `lane_ref` is m-4 payload, deferred to Step-4** — source operator-confirmed
- G8 direct-operator route — **AUTHORITY-BEARING for its ONE directly-addressed recipient, NON-TRANSITIVE, and ceiling-bounded (confusion-not-malice → "by construction" scopes to LIVE-INGRESS authentication only). The operator is NEVER forced to author a governed relay. A direct instruction alone may authorize a directly-addressed app-side action within the current ceiling; a conductor-governed action requiring a typed grant/lineage edge goes through the LANDED grantor grammar (a sanctioned grantor emits the existing typed grant under its own authority, citing the operator — registry.json:105,111), never an arbitrary-agent citation. The recipient's citation is E0 self-reported evidence, NOT a transferable operator grant. Integrity floor: record-never-fabricate; no forged FROM:operator; no silent store mutation; no ceiling raise except via the m-5 typed contract; live-runtime credentialed legs stay operator-performed** — source operator (detailed grill); §8b (r3 non-transitive, VP F11/F17)
- G9 runtime principal map — only the instantiated m-9 worker is a conductor seat; m-10 + m-8 are trusted app components, not seats, hold no submit credential — source VP-ruled F4
- G10 retry — one provider attempt, no automatic retry for the MVP — source VP-ruled F6

Rejected alternatives:
- G4 Option B (m-4 governed routing decision inside Step-3) — REJECTED: over-built for one pinned lane; forces an m-2/m-4 conductor-schema amendment (breaks the no-conductor-change MVP); the governed lane-choice machinery sits idle with a single lane; its real value (governed declaration provenance) lands in Step-4 where routing exists
- G8 forced operator relay-authoring / operator-channel-only authority — REJECTED: adversarial-shaped ceremony that defends only against operator impersonation (malice, out of scope)

Still operator-owned (downstream, NOT packet-ratification inputs):
- E3 / credential provisioning: provisioning a real key + making the first live E3 call — a separate downstream operator gate BEFORE the first E3
- direct-route authority classes 2/3 beyond the recorded model — addable later as a separately-typed, evidence-bearing contract if a concrete need appears

Design-lock impact:
- §8b refined to the authority-bearing / no-forced-authoring model (operator-directed relaxation of the drafted item 4)
- §4/§7/§8c reflect G4 (app-side manifest; m-4 deferred to Step-4); no other packet content changed by the grill

r3 folds (VP r2-review F11–F14, post-grill — BOUNDED repairs, grill NOT reopened):
- F11: §8b made NON-TRANSITIVE — a direct instruction authorizes only the addressed recipient within its ceiling; the citation is E0 evidence not a transferable grant; cross-seat effects use the EXISTING typed grantor grammar (registry.json:105,111), no arbitrary-agent grant; "by construction" scoped to live-ingress only. (Bounded to the landed grammar → no PROTOCOL-DEVIATIONS/Part-F amendment, no grill re-vote.)
- F12: added the m-5 ceiling-host amendment (m-5 sole policy owner, m-10 enforcement host only; pinned ceiling artifact interface + fail-closed) to §8 charter delta + dependency graph (1b).
- F13: chose the §3a no-conductor-change E0 carrier (existing ordinary relay body, m-3 app-event schema, conductor evidence describes carriage only, body not gate-referenceable).
- F14: header/status/lineage/hash cites made exact (r3).

r4 folds (VP r3-review F15–F17, bounded source reconciliation — no product fork, grill NOT rerun):
- F15: §3a pins the EXACT E0 carrier envelope — the m-9 worker's `PHASE: SITREP` / `AUTHORITY: report-only` / `HUMAN_GATE_REQUIRED: no` relay (landed non-authority shape, lineage.go:39-58 + lineage_test.go:14-30), no grant/gate/authority fields, routing TO orchestrator-planner CC m-3.planner; top-level EVIDENCE_TARGET/achieved_evidence = carriage only, the body event carries its own event_evidence=E0/event_integrity=self_reported so a top-level E1/E2 cannot upgrade it.
- F16: m-5 amendment sequencing fixed to a COORDINATED FIRST STAGE (m-10 boundary + m-5 ceiling-host interface-lock together before consumer locks; removed from the step-3 parallel list); m-5 added to the charter-delta line; propagation STAGED (immediate fold records the pending non-consumable m-5 gate, replacement flow creates the m-5 amendment + records its design-of-record fold before m-10/m-9 lock — locked m-5 design not silently rewritten).
- F17: the canonical GRILL_LOCK G8 resolved-decision line replaced with the non-transitive/ceiling/E0/typed-grant wording; §8b item 3 sharpened (typed-grant mandatory for a conductor-governed action); item 7 clarified (direct content in a fresh submission stays non-authority unless it satisfies the grant grammar); the sequence "No lock" qualified as no DESIGN_LOCK_ID/architecture ratification.
```

**Sequence to ratification:** GRILL_LOCK closed → **candidate SHA-256 computed on this finalized packet, reported in the transmittal** → VP exact-candidate re-review → operator hash-bound ratification → source fold (§8, atomic, master-authored) → refreshed consumer audit → replacement dispatches (m-10 + m-5 amendment first). **No `DESIGN_LOCK_ID` / architecture ratification, PLAN, or T4 code token crosses that gate** (the §9 grill lock exists and is satisfied — this is not an architecture lock); the five holds remain in force.
