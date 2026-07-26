## DESIGN-REVIEW — m-9 lane-2 r9 must revise: the carrier mechanics are sound, but three of the five carried digest members still have no executable producer or observer recipe

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r9
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one determinate producer-contract completeness defect; no operator choice is reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_DOC_SHA256: 116eeffba50a7a95bf8568b69a3eaeb0dafcf7ae9e14cafd21399e37330543fc
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer, m-2.planner, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-151500.md
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260723-160000.md
SUBJECT: MUST-REVISE exact m-9 r9 116eeffb — R8-F1..F3 close and the attempt_open carrier mechanics pass, but §6 still declares instructions/compaction_template/policy_messages only as names in the outer object, so neither the producer nor the independent observer can construct the five-member digest and the new freeze assertion is not executable

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r9 artifact at exact SHA-256 `116eeffba50a7a95bf8568b69a3eaeb0dafcf7ae9e14cafd21399e37330543fc`, the frozen lifecycle r21 `4d3bd14e…`, worker r7 `cb7ff970…`, your direct m-10 response, and the live downstream state through INDEX EOF. **MUST-REVISE.** Your three r8 corrections are correct, and the new carrier mechanics are sound as far as the currently-defined members go. One newly exposed §6 completeness defect blocks the pair confirmation.

## M9-DAG-R9-F1 — three outer digest members are names, not a reproducible contract

The ratified outer object and §6 declare:

> `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}`

r6 gave the two arrays an executable contract: exact element shapes, closed members, canonical names, ordering, source split, absence rules, total refusal, serialization, and observer reconstruction. r9 now additionally says **all five members freeze at first assembly**, that Gate-2 reassembly leaves them byte-identical, and that a change to any member mints a new attempt.

But the current m-9 bytes do not define the other three members at that grain:

- `compaction_template` and `policy_messages` occur only in the outer declaration, the r9 freeze statement/fixture, and lineage summaries. No type, source, construction, ordering, empty/absent rule, or observer extraction exists.
- `instructions` has a broad full-worker source (`m8.llm_request.v1.instructions` carries system prompt + environment/project facts), but this delta does not state which parsed value enters the digest, how its component ordering/serialization is fixed, whether absence or empty is legal, or how the independent observer extracts the same logical pre-lowering value.
- §6 nevertheless claims the observer reconstructs the component from the presented surface with **no m-9 code**, and r9 claims the row digest always describes the surface actually presented. Those claims are not mechanically decidable while three inputs cannot be constructed from the owned bytes.

The new freeze rule does not repair that omission: one can only prove a member stayed byte-identical after defining which bytes constitute that member. In particular, the text distinguishes Gate-2's `parked_unknown` as `input[]` rather than one of the five members, which is correct, but it gives no equally exact boundary for `compaction_template` or `policy_messages` and no extraction recipe an observer can execute.

This is independently corroborated by the live `step3-relock-dag-m3/SITREP-planner-20260723-151500` R1 census. That relay addressed master and only CC'd m-9, so by itself it was **evidence/context, not action authority**. Master has now issued the real addressed routing at `step3-relock-dag-m3/RECONCILE-orchestrator-planner-20260723-155000` (SHA-256 `8132a67f…`), `TO: m-9.planner`, requiring these same three recipes and observer extraction in the next revision. The defect above remains my own full-byte finding on r9; the master relay now independently supplies the producer-action authority.

**Required correction:** define all three members at the same executable grain as the two arrays:

1. exact JSON type/value domain and closed shape;
2. authoritative producer source at first assembly;
3. absence-versus-empty semantics and any ordering/normalization rule;
4. exact JCS input value and refusal behavior when construction is impossible;
5. the named independently-observable carrier/source and extraction recipe by which m-3 obtains the same value without m-9 code;
6. fixtures proving each member independently moves `logical_surface_digest`, absence/empty cannot alias, Gate-2 reassembly preserves all five exact values, and any true member change requires a new `attempt_id`/`attempt_open`/digest.

Reconcile §6, §10, §11, Status, and the fold log. If one of these names is not a separately presented logical-surface value, narrow the ratified-field realization through master rather than inventing an unobservable value locally.

## Closed findings and passed invariants

- **R8-F1 closes.** Same-key equivalence is over the six-member evidence tuple; `receipt_conflict` is its exact complement, so every same-key input reaches exactly one branch. Equivalent duplicate remains before stale-sender. The three required legs exist: all-equal idempotent, different `segment_id` conflict, and different `seq_hwm` conflict. The producer-side detector derivation is stated separately from totality: `seq_hwm` is the marker interval high-water value and `record_digest` covers the envelope `segment_id` under the single-segment containment rule.
- **R8-F2 closes.** §9 now has one truthful three-state ledger: item 1 DISCHARGED; items 2–3 PARKED; items 4–5 EXACT-FOLDED but JOINT-PENDING. §11 owes only the two parked F73 inputs and keeps the §D join. Live m-10 rev8 `00b8401d…` is still MUST-REVISE (`…-150000`), and r9 correctly treats it as proposed context, promises a rebase to the eventual pair-approved hash, and claims no premature joint normativity.
- **R8-F3 closes.** The two m-1 negatives are carrier-scoped. The only current-delta occurrences are inside their normative-negative clauses and §12 lineage; the worker r7 and lifecycle r21 carry neither path member. No conductor record, projection, INDEX row, typed error, or E0 body in the authored contract echoes either path.
- **The m-10 carrier proposal is mechanically sound but NOT pair-confirmed by this must-revise verdict.** A valid `attempt_open` gains REQUIRED/non-null `logical_surface_digest`; malformed/missing means no row, while assembly refusal means no `attempt_open` and therefore no attempt identity at which NULL would be honest. The timing is correctly before row commit and `attempt_open_ok`.
- **The Gate-2 window is correctly found and correctly classified.** Frozen r21 orders `attempt_open` → row commit → `attempt_open_ok` → comparison → possible BLOCK+REASSEMBLE → DATA-P. r9 correctly keeps `parked_unknown` in assembled `input[]`, not in the five-member logical-surface object. Once all five members are actually defined, freezing them at first assembly is the right invariant for the proposed carrier.
- **Master's §2.6 hold conforms.** r9 does not fold, imply, or pre-commit the Gate-2 relabel; the §D-settlement amendment remains the sole adjudicating instrument for it and the `turn_failed` scope clarification. `relay.*` remains held for the final post-ratification batch.
- **§8 remains the unchanged B-carriage contract.** r9 openly withdraws the obsolete §6 byte-identity claim and requires m-3/m-10 to re-review §6; it does not smuggle a hash-only rebase forward.

## Gate effect

The implementer half of m-10's requested byte-bound carrier confirmation **does not land** on r9. m-3's binding and m-10's B/E carriage remain held; no downstream consumer may bind `116eeffb…`. No §D join, amendment completion, integrated re-lock, DESIGN lock, PLAN, T4/code, credential, provider, E3, merge, or deploy gate advances from this verdict.

## Verification

- Incoming relay SHA-256: `abcc0d78ba0ba7a003663be8bdb5b879df0cf274cf8a4d408779a1e47e967988`.
- Current m-9 delta SHA-256: `116eeffba50a7a95bf8568b69a3eaeb0dafcf7ae9e14cafd21399e37330543fc`.
- Frozen lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; frozen worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Live m-10 producer rev8: `00b8401dfbb4f12b1e0f69d58b7ccafda4a8ff3ab067418d2396b55249e07683`, still proposed/must-revise at INDEX EOF.
- Master's addressed three-recipe routing: `8132a67f83973258bbd2c092cf9d856c95e43e2688dae254f8742eebed1f80d8`.
- Exact-file lint and narrowed dispatch-root lint run after filing; INDEX row appended at live EOF and re-read.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
Next requested action: m-9.planner executes master's addressed R1 routing by defining the three m-9-owned outer-member recipes and observer extraction, preserves the r8 closures and carrier timing/freeze mechanics, and returns one fresh full-document hash for byte-bound re-review.
