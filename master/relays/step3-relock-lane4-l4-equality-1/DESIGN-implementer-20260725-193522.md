## DESIGN — BYTE-EQUALITY CONFIRMATION 1: the eleven materialized artifacts are byte-for-byte identical to their proposal-envelope `content`; every disk byte length and SHA-256 equals the planner's declared value. The twelfth artifact, `xit-dur-1/input.json`, remains held and is absent as required. This is duty (i) only — no content-review verdict is made or implied.

ROLE: Implementer
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-l4-equality-1
PARENT_DISPATCH_ID: step3-relock-lane4-l4-materialize-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this relay records the independent byte-equality check requested by the addressed materialization receipt. It grants no authoring, materialization, freeze, re-lock, owner-fidelity, T4, or external-use authority.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-l4-materialize-1/DESIGN-orchestrator-planner-20260725-192902.md
FROM: l4.implementer
TO: master.orchestrator-planner
CC: l4.planner, operator, master.orchestrator-reviewer, m-3.planner
SUBJECT: Equality 1 confirmed for eleven materialized artifacts; direct binary comparison + declared length/SHA match; held xit-dur-1 absent; duty (i) only

BYTE_EQUALITY_VERDICT: confirm

## Independent method

I parsed all twelve `proposal_artifact` envelopes directly from
`master/relays/step3-relock-lane4-l4-propose-1/DESIGN-planner-20260725-183149.md`.
For each envelope I independently checked:

1. the envelope body's byte length and SHA-256 against its declared `byte_length` and `sha256`;
2. for each of the eleven materialized paths, the file re-read from disk against the declared length and SHA-256;
3. direct binary equality between the disk bytes and the proposal envelope's `content` bytes.

This does not trust the materialization receipt's digest table. The direct binary comparison is the equality proof; the independently recomputed length and SHA-256 are additional E2 checks.

## Results

| path | bytes | independently recomputed disk SHA-256 | direct proposal-body equality |
|---|---:|---|---|
| `master/exit-fixtures/baseline/baseline-workload.json` | 1986 | `24fba6e84c2703512afaede1ac177170872e8b4403030bd539802488e697f258` | PASS |
| `master/exit-fixtures/baseline/baseline-config.json` | 1685 | `e7a9ae8b949cee9e3a712c5b1384706649f8764039f68940acfd238ea96e1740` | PASS |
| `master/exit-fixtures/xit-gov-1/input.json` | 6034 | `42fcde5a4d515c761fa2eb942975e3ae884ea844307045314a19ee232c153a18` | PASS |
| `master/exit-fixtures/xit-dur-2/input.json` | 4052 | `2f93ee699173feaf95d06283e257b36dfeaddf91883ed85f002701d6ffb79d05` | PASS |
| `master/exit-fixtures/xit-dur-3/input.json` | 5093 | `20aeca2c23ef3d77cc957c811b1835e30f8da981bb9231f16304caf7122e847e` | PASS |
| `master/exit-fixtures/xit-dur-4/input.json` | 3942 | `c96bd49c717f991878e15ba3bfba03429fc1ac61017e63762d3bd9fbbab49d71` | PASS |
| `master/exit-fixtures/xit-dur-5/input.json` | 3526 | `84a3ce89b41f49b8163e295cd383213b82534021c3531410c555e4c679dba7bf` | PASS |
| `master/exit-fixtures/xit-crash-1/input.json` | 3742 | `266f6f17b93e22007d3d9ed937b3d1d12d36c2ca3ad39752652ed32a9dd18304` | PASS |
| `master/exit-fixtures/xit-inj-1/input.json` | 3944 | `d2161c374be31286591b8567ef9cc3cb3e902542c1681c25177e147d582ec354` | PASS |
| `master/exit-fixtures/xit-ho-1/input.json` | 3729 | `15bdc1b33d15635db940e86c5119b3356dcb053ab8467e48cf6b136362bb8dd2` | PASS |
| `master/exit-fixtures/xit-op-1/input.json` | 5355 | `7c262394537bc041b85db4feced39fee150f77e0da0ee1a702f06b09572686b2` | PASS |

Summary: `12/12` proposal bodies reproduce their declarations; `11/11` materialized files match their declarations and proposal bodies; `master/exit-fixtures/xit-dur-1/input.json` is absent; failures `0`.

## Held artifact

`master/exit-fixtures/xit-dur-1/input.json` is not part of this equality confirmation. I independently confirmed it does not exist. Its absence matches the addressed receipt and the esc2 disposition: it remains held for a corrected generation that pins the exact interval and frozen input bytes and retires the known-provisional clause.

## Duty separation

This relay decides proposal-to-file byte equality only. I did not review scenario discrimination, expected canonical rows, typed expectations, carried-obligation rows, observer/locator resolvability, sample weights, or owner fidelity. Those belong exclusively to duty (ii), after an addressed `…-l4-content-req-<r>` arrives and the manifest exists.

## Boundaries

No materialized artifact, proposal envelope, manifest, lock constituent, owner/frozen final, or `frank/` path was modified. Nothing is frozen or re-locked; no held member is resumed; no owner-fidelity request, PLAN/T4 token, E3 claim, external use, merge, or deploy is authorized. H-12 remains in force.

## Verification

Fresh E2 comparison command at this relay's authoring session returned:
`envelopes=12 materialized=11 direct_equal=11 held_absent=1 failures=0`.
Both the proposal relay and addressed materialization receipt independently pass exact-file `relay-lint --no-freshness`.

ACTIONS_GIT_REF: docs-workspace disk action — this `l4.implementer` equality relay plus one append-only `master/relays/INDEX.md` row. No governed artifact or `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd `/Users/jack/Programming/harness` is not a git repository (governance docs workspace).
Next requested action: master consumes this equality confirmation in the approved chain. I remain standing by for the next obligation directly addressed to `l4.implementer`; duty (ii) does not begin before an addressed `…-l4-content-req-<r>`.
