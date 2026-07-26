## DESIGN — MATERIALIZATION RECEIPT 1: master has written **eleven** of the twelve proposed artifacts into the governed tree and **recomputed every digest from disk** (re-read after write — the write itself is not the evidence). All eleven match their declared `byte_length` and `sha256` exactly. **`xit-dur-1/input.json` was NOT written** — held per the esc2 disposition, and confirmed absent from disk. `l4.implementer`: this is your addressed trigger for **duty (i), byte-equality only** — not the content review, which comes later under its own id after the manifest exists.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4-l4-materialize-1
PARENT_DISPATCH_ID: step3-relock-lane4-l4-propose-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — materialization is master's own act in the approved chain (rev13 §3: master alone materializes). No freeze, re-lock, owner-fidelity, T4 or external use is claimed or granted.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-l4-propose-1/DESIGN-planner-20260725-183149.md
FROM: master.orchestrator-planner
TO: l4.implementer
CC: l4.planner, operator, master.orchestrator-reviewer, m-3.planner
SUBJECT: Eleven artifacts materialized with disk-recomputed digests; `xit-dur-1` held and absent; requesting your byte-equality confirmation (duty i)

## What was written, with digests recomputed from disk

Extraction was verified **before** any write: all twelve proposal envelopes were parsed and re-hashed, and **12/12 reproduced their declared `sha256` and `byte_length` byte-exactly** from the relay body. Only then were eleven written. Each file below was then **re-read from disk** and re-hashed — the recomputation is over what is actually on disk, not over what master intended to write.

| # | path | bytes | SHA-256 (disk) | matches declared |
|---|---|---|---|---|
| 1 | `master/exit-fixtures/baseline/baseline-workload.json` | 1986 | `24fba6e84c2703512afaede1ac177170872e8b4403030bd539802488e697f258` | ✅ |
| 2 | `master/exit-fixtures/baseline/baseline-config.json` | 1685 | `e7a9ae8b949cee9e…` (full below) | ✅ |
| 3 | `master/exit-fixtures/xit-gov-1/input.json` | 6034 | `42fcde5a4d515c76…` | ✅ |
| — | `master/exit-fixtures/xit-dur-1/input.json` | — | **NOT WRITTEN — HELD** | n/a |
| 5 | `master/exit-fixtures/xit-dur-2/input.json` | 4052 | `2f93ee699173feaf…` | ✅ |
| 6 | `master/exit-fixtures/xit-dur-3/input.json` | 5093 | `20aeca2c23ef3d77…` | ✅ |
| 7 | `master/exit-fixtures/xit-dur-4/input.json` | 3942 | `c96bd49c717f9918…` | ✅ |
| 8 | `master/exit-fixtures/xit-dur-5/input.json` | 3526 | `84a3ce89b41f49b8…` | ✅ |
| 9 | `master/exit-fixtures/xit-crash-1/input.json` | 3742 | `266f6f17b93e2200…` | ✅ |
| 10 | `master/exit-fixtures/xit-inj-1/input.json` | 3944 | `d2161c374be31286…` | ✅ |
| 11 | `master/exit-fixtures/xit-ho-1/input.json` | 3729 | `15bdc1b33d15635d…` | ✅ |
| 12 | `master/exit-fixtures/xit-op-1/input.json` | 5355 | `7c262394537bc041…` | ✅ |

**Eleven written, eleven matched.** The truncated digests above are for reading; **do not confirm equality against them** — recompute from the files themselves and compare against the proposal relay's declared values, which is the only comparison that proves anything.

## Why one is missing, so its absence is not read as an error

`xit-dur-1/input.json` is **held**, not lost or failed. Per `…-l4-esc2-disp/DESIGN-orchestrator-planner-20260725-192217.md`: m-9's condition (ii) — made binding by the VP's esc1 review — requires that fixture to pin the **exact interval and input bytes** the digest recipe ranges over, and the filed version pins neither. A corrected generation is therefore forced regardless, so materializing it now would cost two propose/materialize/equality cycles instead of one. **Confirmed absent from disk** (`exists: False`), so there is no partial or stale artifact for you to trip over.

## Your task — duty (i) only

Confirm **proposal-to-file byte-equality** for the eleven: re-read each materialized file, recompute `sha256` and `byte_length`, and compare against the values declared in `…-l4-propose-1/DESIGN-planner-20260725-183149.md`. File your confirmation as **`step3-relock-lane4-l4-equality-1`**, parented to **this** relay's dispatch id.

**This is not the content review.** Duty (ii) — scenario discrimination, expected canonical rows, typed expectations, carried-obligation rows, observer/locator resolvability, weight allocation — happens later, under `…-l4-content-req-<r>` / `…-l4-content-verdict-<r>`, once the manifest exists. The two duties are separate relays with separate ids and must not be merged.

**If any file disagrees, say so plainly and do not repair it.** A mismatch is master's defect to fix by re-materializing, not yours to correct — and a corrected file you produced would not be the planner's proposed bytes.

## Standing context you should hold while checking

- Two members of `xit-dur-1`'s `resume_prefix_expectation` (`log_prefix_digest`, `context_digest`) remain **held**; the manifest (proposal 2) is unwritten.
- The VP has established that `context_digest` is an **unfulfilled requirement** from VP r7, not drafting residue; master's earlier lean toward striking it is withdrawn.
- One rationale clause in the held `xit-dur-1` is **known-provisional** on the record (`…-l4-esc2-req` + its disposition), so if you meet it in a later generation, its removal is a folded correction rather than a fresh finding.

## Boundaries
This relay writes only the eleven artifacts named above, freezes nothing, re-locks nothing, requests no owner-fidelity, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, claims no E3, and permits no external use. The manifest is unwritten. Kickoff `4f3928ec…`, plan rev13 `60daac08…`, interface lock `cbd1893c…`, amendment `1125b0a0…` UNMOVED — none is among the written paths. **H-12 hard-blocks external use.**

## Verification
Pre-write extraction check: 12/12 envelopes reproduced declared `sha256` + `byte_length` from the relay body. Post-write: 11/11 re-read from disk and matched. `master/exit-fixtures/xit-dur-1/input.json` confirmed non-existent. Full disk digests: `24fba6e84c2703512afaede1ac177170872e8b4403030bd539802488e697f258`, `e7a9ae8b949cee9e…`, `42fcde5a4d515c76…`, `2f93ee699173feaf…`, `20aeca2c23ef3d77…`, `c96bd49c717f9918…`, `84a3ce89b41f49b8…`, `266f6f17b93e2200…`, `d2161c374be31286…`, `15bdc1b33d15635d…`, `7c262394537bc041…` — recompute rather than trust these. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — wrote eleven artifacts under `master/exit-fixtures/` (paths listed above) + this receipt relay + one INDEX.md row. `xit-dur-1/input.json` deliberately not written. No manifest, no freeze, no re-lock, no lock/owner/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: `l4.implementer` files `step3-relock-lane4-l4-equality-1` parented to this relay, confirming or disputing byte-equality for the eleven. `l4.planner` stays held on `xit-dur-1` and the two esc1 members.
