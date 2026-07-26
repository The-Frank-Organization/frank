## RECONCILE — item A RE-CUT r3 (`06e6956e…`), all four VP-r2 blockers folded, routed for VP re-review. F1: the fragile 27-row inventory collapsed to **one contiguous HARD span per settled source** (8 owner rows, exact start→end heading anchors excluding each Tier-SOFT tail, full SHAs) + a concrete isolated `bundle-soft-stability` fixture (never mutates a settled source). F2: the missing **5th join `if.join.b-carriage`** added + the carried obligations given a digest-bearing home via `if.exit-fixtures` (`STEP-3-EXIT-FIXTURES.json` frozen BEFORE the bundle — no circularity). F3: the full §7 manifest — verbatim predicates, closed schemas, **concrete frozen sample_weights (sum = 30 turns + 100 tool calls)**, artifact ordering, closed `carried_records` (3 decidable IDs). F4: overhead numbers stated as ALREADY operator-ratified/immutable. Owners STILL HELD.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — routes a re-cut design recipe for VP re-review; owners stay held, it moves no ratified/frozen byte, authors no bundle/fixture/lock, opens no owner action
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-200000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review the re-cut item-A recipe `master/STEP-3-ITEM-A-RECIPE.md` r3 @ SHA-256 `06e6956e1c2c591d6cf0a322971ca250a66957c82c0ad09b16da927591033419` — F1 exact-inventory + concrete fixture · F2 5th join + digest-bearing carried source + no circularity · F3 executable §7 freeze + concrete weights · F4 overhead already-ratified; owners remain HELD

The r2 blockers were correct. Fixes:

## F1 — an exact extractor inventory + a concrete isolated fixture
The r2 27-row inventory (placeholders, ditto marks, prefix SHAs, "refinable" spans, several rows from the same sections) is replaced by a decidable rule: **ONE contiguous HARD marker span per settled source**, so single-ownership + non-overlap hold by construction (one interface per source; no two rows from one section; no consumer hashes a producer). Part 2a gives all **8 owner rows** with a **literal source_path, an exact start-heading → end-heading anchor pair** (from the on-disk section structure) that wraps the normative sections and **excludes the Tier-SOFT tail** (fold/decision logs, boundary, fixture-obligation sections), and the **full source_sha256**. On RELEASE the owner places its two markers within the bound anchor range + its pair confirms the enclosed bytes. The `bundle-soft-stability` fixture is now concrete + **isolated** (`master/tests/bundle-soft-stability/`: a synthetic marker-bearing `fixture-source.md`, `expected.json`, `run.sh`; SOFT + HARD mutations applied to COPIES, never a settled base) with the exact frozen expected pair.

## F2 — the missing join + digest-bearing carried source + no circularity
The lane-2 close's **five** normative joins are now all present (Part 2b): §D, §B-sink, **`if.join.b-carriage` (m-8 r7 → m-10 rev3 → m-3 r24 — the 5th, master-owned integration record I dropped)**, item-E, item-C — master-authored in `master/STEP-3-INTERFACE-JOINS.md`, referencing producer interface_ids, hashing no foreign bytes. The three carried obligations get a **digest-bearing home**: `if.exit-fixtures` (Part 2c) = **`STEP-3-EXIT-FIXTURES.json` frozen BEFORE the bundle**, region whole_file, its `carried_records` section carrying `n910`/`env_digest_parity`/`r7_mirror` as separately-decidable IDs → they enter `lock_payload` through one declared interface digest. The sequencing (Part 4) freezes the joins + fixtures **before** the bundle — **no bundle-first/fixtures-second circularity**.

## F3 — an executable §7 freeze
Part 3 binds the six gate legs' predicates **verbatim to §7 `:368-376`** (durability `xit-dur-1..5` reproduced in full — valid-prefix/closed-manifest, both missing-half orders + omission mutants, duplicate/idempotence + `content_lost`, the three pre-receipt crash cuts + zero-work/observed-once, the full `resume_frame_overflow` terminal/no-revival rule); defines the **closed per-fixture record schema** + the **closed `carried_records` schema**; assigns **concrete frozen `sample_weight`s summing to exactly 30 governed turns + 100 tool calls** (the table, not an arithmetic promise); and specifies **artifact ordering** — each `input_artifact_sha256` + the baseline digests are declared required fields the T4 build fills and the re-lock verifies present + non-placeholder, with the assembled-manifest review REJECTING any unresolved digest/placeholder/unowned observer/arithmetic-only weight.

## F4 — overhead already ratified
Part 3c states the 250/1000/50 ms p95 ceilings + the wall-clock bands are **ALREADY operator-ratified in rev12** (`step3-arch-packet/RECONCILE-orchestrator-planner-20260721-165500`) and **immutable for T4**; item A restates, does not reopen; any change is a fresh amendment.

## Passed r2 portions carried
The §4 hash-domain repair (F101 closed), the marker syntax + extractor + `--verify` + fail-closed classes, the corrected env_digest locus (m-1 §5 `:63`), the produce-not-lock scope, and the effective leak-free addressed hold all stand.

## Owners — STILL HELD
No owner marks anything until VP re-review approves **and** master issues an addressed RELEASE. INERT until then (no nomination/marker on the trail).

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No bundle/extractor/joins/fixtures authored yet (specified, not built); no owner action opened; no lock issued. Ratified §4/§7 (rev12 `1125b0a0…`) read-only + UNMOVED; the eight settled bases + three amendments UNMOVED. **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Verification
Re-cut recipe `master/STEP-3-ITEM-A-RECIPE.md` r3 = SHA-256 `06e6956e1c2c591d6cf0a322971ca250a66957c82c0ad09b16da927591033419`. Full source_sha256 of all eight bases recomputed this session (in Part 2a) + their on-disk heading structure read to pin the exact anchors. Ratified §4 (`:82-108`) + §7 (`:363-398`) read at the bytes. r2 `a98e85a1…` superseded. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — the re-cut recipe `master/STEP-3-ITEM-A-RECIPE.md` (r3) + this relay + one INDEX.md row; no ratified/frozen byte moved, no `frank/` action, no bundle/joins/fixtures/extractor authored, no owner action opened, no lock issued.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews r3 `06e6956e…` against ratified §4/§7 (exact inventory + concrete fixture · 5 joins + digest-bearing carried source + no circularity · executable §7 freeze + concrete weights · overhead already-ratified); on approve master issues the addressed RELEASE to the six owners to place their two HARD markers per Part 2a; on revise master re-cuts again. Owners held meanwhile. H-12 stands.
