## DESIGN-REVIEW — MUST-REVISE m-9 stage-4 full-worker r4 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-mvp-stage4-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator's GRILL_LOCK decisions stand; the remaining defects require explicit owner-real correlation/version bytes, not a new product choice
GRILL_REQUIRED: yes — satisfied by `GRILL_LOCK_ID: step3-mvp-stage4-m9-worker-grill-1`; no operator choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m9-worker
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 7b8c4a3d5cd5098ba342313336a654041de0f87d65714902ab7906f75a1dacae
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage4-m9/DESIGN-planner-20260720-153000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-160000.md
SUBJECT: MUST-REVISE exact r4 7b8c4a3d — the Tier-0 inbox source has no turn/admission correlation or replacement-readable identity and its build anchor is opaque; the claimed concrete F58 vector still lacks actual version values/derivation/catalog digest and retains contradictory stale m-2 consumption prose

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete r4 design bytes at SHA-256 `7b8c4a3d5cd5098ba342313336a654041de0f87d65714902ab7906f75a1dacae`. R4 closes R3-F1 and R3-F4: E14 now controls the wake insert it inventories, and `sent` is an honest proposal over observed `attempt_started`, not a fabricated provider-wire claim. The m-2 relay digests, five local schema bytes/digests, and local-member absence encoding are also corrected. Two assemblability blockers remain.

## Findings

### M9-S4-R4-F1 — BLOCKER: the Tier-0 inbox relay is readable but not correlated to the admitted turn

§2.1/§7.1 now name an owner-real storage surface — an inbound task relay in the worker's recipient inbox — but do not provide an owner-real selection/join from the admitted turn to that relay:

- Frozen m-10 r36 `turn_open` carries only `{run_id, turn_id, turn_epoch, parked_unknown}`. It carries no `relay_id` or `admission_ref`.
- m-10's private `turns` row does contain `admission ref (wake relay_id or operator input)`, but the worker cannot read m-10's private store. That exact private-reader boundary is the reason r19 §2.6 rejected recovery from m-10 state.
- `project` can list the recipient inbox, but the inbox may contain multiple records and no rule in r4 proves that an inbox ordering/choice equals the `wake_schedule` row m-10 selected. m-10 never touches the conductor, so the two components cannot independently infer the same task by content.
- R4 says replacement re-reads “the same inbox relay,” yet no durable relay identity is carried to the replacement. §7.1 replacement semantics and the §10 no-second-truth fixture still say it rebuilds only its own authored audit-projection records, contradicting the new recovery claim.

Consequently, “captured into turn state at turn start” is not executable: the worker can read tasks but cannot know which task owns this `turn_id`, and the operator-input admission branch has no inbound task relay at all.

The static-build half is also overclaimed as “integrity-anchored by `manifest_digest`.” `assign` gives the worker only the opaque manifest digest; manifest bytes/release binding are m-10-private, so the worker cannot resolve that digest to an expected `m9_worker_build_digest` or compare its own artifact. If the intended guarantee is m-10's pre-admission release/serve gate, cite and consume that owner enforcement rather than presenting the digest alone as a worker-verifiable build anchor.

Required correction: specify one stable, replacement-readable admission→task binding. If that requires `admission_ref`/`relay_id` on `turn_open` or another m-10 frame, route an m-10 owner amendment; do not infer it from inbox order. Cover both wake-relay and operator-input admission. Reconcile §7.1 derivation/replacement prose and §10 fixtures with the actual task source. For the build half, name the real validator/comparison path and do not claim an opaque digest is independently resolvable by m-9.

### M9-S4-R4-F2 — BLOCKER: the F58 expected vector is described as concrete but cannot be serialized or hashed

§8.3 materially fixes the schema half: all five local schema digests recompute exactly, the three m-2-produced relay digests are consumed correctly, and local mapping-version is absent. But the required version members remain undefined:

- “`tool_impl_version` for each local tool + the catalog version are content-derived” gives no derivation function, input byte set, token grammar, or actual value for any of the eight vector rows. It also does not state the concrete m-2 mapping-version value in the expected rows.
- Binding local implementation versions only to schema bytes is insufficient: implementation behavior may change while the argument schema stays byte-identical. Such a build would keep the declared version and `tool_catalog_digest` unchanged, contradicting F58/F63's implementation-drift-without-identity-change fail-closed purpose.
- The document provides no enumerated eight-row expected vector and no computed expected `tool_catalog_digest`. Therefore stage 6 still cannot serialize the claimed vector, recompute the value, or hand T4 a concrete expected comparand.
- §11 flag 4 retains pre-r4 contradictory text: it says the catalog consumes m-2 “rendered schemas” and “m-9 digests, m-2 defines.” Frozen m-2 §3.2 and corrected §8.3 say the opposite — m-2 produces the normalized static/template digests and m-9 consumes them byte-exact without re-digesting a render.

Required correction: define the exact version field(s), their canonical derivation/grammar and concrete expected values; enumerate the eight expected vector rows including `m2-mapping-v1` where applicable and member absence where not; compute and state the resulting expected `tool_catalog_digest`; and define a post-build comparison that detects implementation/catalog drift rather than only schema drift. Reconcile §11 flag 4 to the corrected producer split.

## Accepted r4 substance

- Incoming relay, r4 document, H-17 schema, and all seven frozen owner hashes reproduce exactly; incoming exact-file relay lint is OK.
- R3-F1 is closed: E14's decision/enforcement/effect-linearization now consistently describe `INSERT OR IGNORE` + `UNIQUE(relay_id)` → committed `wake_schedule`; downstream turn admission remains m-10-owned.
- R3-F4 is closed on m-9's side: `sent` is emitted only on observed `attempt_started`, explicitly carries no provider-wire claim, and the table/fixture/no-emission cuts agree. m-3 token-semantics and m-8 observability confirmations remain mandatory before stage 6, as r4 correctly states.
- The 15 H-17 rows remain self-contained and mechanically complete; the 15+2 inventory is consistent.
- The five local schema reference digests independently recompute exactly, the three m-2 relay digest values match frozen m-2, and local mapping-version absence is correct.
- The ephemeral/no-reload transcript, GRILL_LOCK, F59, bash/teardown honesty, compaction direction, L5=B, and the release-binding direction remain accepted.

This verdict is byte-bound to `7b8c4a3d5cd5098ba342313336a654041de0f87d65714902ab7906f75a1dacae`. A corrected design requires a fresh uniquely-parented m-9.implementer DESIGN-REVIEW. No SITREP, consumer-confirmation routing, stage-6 lock, PLAN, T4 token, implementation, release binding, E3, merge, or deploy may consume r4 as approved.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `9a16f052984751558d9d16460a8ce4633e9dd4dbe04284834384127fab4389f4`.
- Exact reviewed m-9 r4 SHA-256 recomputed: `7b8c4a3d5cd5098ba342313336a654041de0f87d65714902ab7906f75a1dacae`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen owner hashes recomputed: m-9 r19 `2a96a07b…`; m-10 r36 `0240e874…`; m-8 r12 `4b670a79…`; m-7 r11 `9331ea88…`; m-3 r4 `009df607…`; m-2 `83d8e63e…`; m-1 `7c8b09a6…`.
- Local schema digests independently recomputed: `read` `4dc4e270…`; `write` `0863ca49…`; `edit` `396e58a8…`; `apply_patch` `f6594e18…`; `bash` `ddd9efb1…` — all exact.
- Census mechanical scan: 15/15 effect rows carry every required label; exactly 2 non-effect rationales.
- Full-byte pass: §§0–13, E1–E15, both rationales, GRILL_LOCK, fold log, r19 recovery boundary, m-10 `turn_open`/private admission-ref/wake/manifest surfaces, m-8 `attempt_started` + disposition cuts, m-2 F58 producer/version/absence bytes, amendment F58/F63 staging, and live index/duplicate-response scan.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte must-revise relay and appended one `master/relays/INDEX.md` row; no design/source doc, historical relay, frozen contract, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file `relay-lint.py` verification on this relay.
Next requested action: m-9.planner resolves M9-S4-R4-F1..F2 in one corrected revision, routes any needed m-10 admission-ref carrier amendment explicitly, makes the F58 expected vector/digest actually concrete, rechecks every frozen hash, and returns a fresh uniquely-parented full-byte DESIGN relay; later gates remain held.
