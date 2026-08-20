# s12 H-16/H-26 Fix-Lane Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans (single Implementer seat; implementation starts ONLY on a live `DISPATCH IMPL` per the commission's delegated-dispatch conditions). Steps use checkbox (`- [ ]`) syntax for tracking.

**PLAN_LOCK_ID:** `s12-h16-fix-plan` · **DESIGN_LOCK_ID:** `h16-outcome-split-design` (**rev21**, FROZEN @ sha256 `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05`) · **Owner:** s12 · **Date:** 2026-08-17 · **Rev:** r4 — the bounded R3-F1..F4 accuracy/coverage fold over r3 (superseded; review `s12-build/PLAN-REVIEW-implementer-20260817-205921.md`, must-revise; the ruled realization and rev21 bind PASSED and carry unchanged): F1 the governance-truth statement (the pair's dispatch gates vs the still-open rev21 owner-rerun/join tail — both stated, the false absolute removed); F2 the routed-lint oracle replaced by the exact expected-output rule (the two immutable r1 errors make an unqualified "clean" unattainable); F3 the substrate-ownership contradictions removed (one byte-consistent model: worktree = code, shared relay root = append-written trail, nothing s12-committed on the branch or `main`); F4 Task 9 gains the operator-only raw-ingress negative authority battery for both `attempt_resolution` shapes + the positive operator control. r3 = the RULINGS fold over r2 (master's consolidated relay `master/relays/step3-h16-h26-lane/RECONCILE-orchestrator-planner-20260817-202541.md`): the design rebinds to rev21 (`e09fab09…` superseded via the void rail — the one-byte-run §4a Ask-2 delta, m-7-authored, pair-APPROVED at 0 blocking `…-202647`); MUST-CHECK 1 RULED (m-2 `DESIGN-planner-m2-20260817-201520.md` §1 is the Task 4/9 authority, realized byte-exactly); MUST-CHECK 7 RULED (the Task 10 byte = `mint-predecessor-mismatch`, rev21 `:89`); the Task 4 and Task 10 gates CLEAR; Task 16's substrate commit is REMOVED (master's checkpoint cadence banks the substrate on `main` — observed at `2cb9a0c`/`414ee38`; the branch carries code only); BASE re-pinned to worktree-creation time. r2 = the bounded F1–F5 fold over r1 (superseded; review `s12-build/PLAN-REVIEW-implementer-20260817-192631.md`, must-revise): F1 the active-root lint shape (bare `DESIGN_LOCK_ID`, no in-root record-kind claim, the plan hash annotated on `PLAN_LOCK_ID`); F2 `type: string` restored literal + the `record_kind` cardinality ESCALATED to master (MUST-CHECK 1, held); F3 Task 4 gains the version-marker consumers (`config.go` reader marker + transition, `genesis.go` downgrade step, the asserting tests) and the pinned provenance shape; F4 the predecessor-mismatch rejection byte ESCALATED (MUST-CHECK 7; Task 10 held on the ruling); F5 the boundary contract added, the escalation scan carried on the PLAN relay, Task 0 moved to the mandatory isolated-worktree procedure with the substrate commit moved to Task 16. No task decomposition, battery binding, or scope-fence change beyond the named folds.

**Goal:** realize the four-leg lane scope — the H-16 `decision_state × post_commit_state` Outcome split with its full derived-work/credential-integrity machinery, `-mint` (H-26) and `-init` (R-INIT-UNLOCKED) converged onto the shared `AcquireRoot` phase −1 lock helper, and the H-16-REG fieldspec registry realization — at E2, branch-only, with every rev21 battery green.

**Spec (the ONLY design authority; this plan adds no mechanism):** `master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` — **rev21, FROZEN at sha256 `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05`** (rev20 `e09fab09…` superseded by the void-rail Ask-2 delta: the one-byte-run §4a token naming, m-7-authored, m-2 section-split-proved, pair-approved `…202647` at 0 blocking). **Re-verify this hash from disk before consuming a byte. Any mismatch, and any defect found in the design's own bytes: STOP the leg, relay UP to master (the exact-hash void rail). Never a silent local fix.** Cited below as `§n`. Companion AUTHORITIES for the ruled surfaces: m-2's `master/relays/step3-h16-h26-lane/DESIGN-planner-m2-20260817-201520.md` §1 (the `record_kind` realization — Tasks 4/9 realize it byte-exactly). Companion inputs (guidance, not authority): m-2's confirm chain `…/DESIGN-planner-m2-20260817-133642.md` + `…-140244.md`; the scope words `…/RECONCILE-orchestrator-planner-20260817-011530.md` + `…-184510.md` (§3 = the enumeration of record); the commission `…/PLAN-orchestrator-planner-20260817-184653.md`; the consolidated rulings relay `…/RECONCILE-orchestrator-planner-20260817-202541.md`.

**Architecture (rev21's, restated for orientation only):** the committed decision (`decision_state`) never relabels; post-commit hook truth rides a second dimension (`post_commit_state ∈ {complete, pending, failed, unknown}`); legacy `state` is a fail-closed projection (present iff `complete`). Class-G global housekeeping heals via a dirty flag + pre-serve pass; Class-D decision-scoped work is derived from canonical bytes with durable attempt markers. Credential integrity rides the `mint_predecessor` pivot chain, the upgrade completeness proof, the effective-quarantine predicate, and the offline recovery ceremony under the phase −1 root lock. Three system-stamped headers (`hook_contract`, `mint_predecessor`, `admin_provenance`) join the registry with an executable presence-pinned rejection contract.

**Tech stack:** Go (existing module `frank/`), stdlib; `go test ./...` is the E2 umbrella. No new dependencies.

**Basis pins (verified by s12.planner 2026-08-17):** rev21 re-hashed from disk this act → `cc8bcff3…` exact. The defect loci were verified at `HEAD:frank == 80dfb152…` (the pin m-2's stage-2 confirm verified): `internal/fieldspec/validate.go:35` still carries `raw != ""`; `internal/engine/loop.go:160-187` still relabels on all four hook-failure sites; `cmd/frank/main.go:89-117` runs `-init`/`-mint` before `AcquireRoot` (`:118`). Workspace `main` moves under master's checkpoint cadence (the moves bank docs/relays, no `frank/` source byte) — **BASE = workspace `main` at worktree-creation time; Task 0 records the exact sha and re-verifies all loci at that branch point**; any source-loci drift → report to s12.planner before proceeding.

## Boundary contract (the plan-level contract; per-task contracts inherit it)

```text
Writes: canonical store records (Outcome-split emissions, derived-work attempt/transition/resolution records, mint_predecessor-linked seat_mint pivots, mint-chain-anchor records, admin_provenance-stamped ceremony commits); the fieldspec registry rows/tokens/marker (v9); the widened seat.Resolve refusal contract; the -init/-mint/ceremony lock acquisitions
Reads: rev21 @ cc8bcff3… (the sole design authority) + m-2's 201520 §1 (the ruled record_kind realization); the live registry bytes; canonical records + the RAW pre-dedup redo stream (§4a proof); the binding table; the intake journal
Target entity: the frank conductor binary + its canonical store semantics on branch s12-h16-fix
Downstream consumer: the engine loop's reply surface (MCP/native/shared-client forwarding); the state-only consumers (§8: nudge, prompters, resummon, MCP re-render) — asserted fail-closed, NOT migrated; seat auth via channel.Server; the s12 end review → master+VP lane review → the operator MERGE-GATE
Contract: rev21 §2 wire shape ({state?, decision_state, post_commit_state}); the §4 header-rejection contract (presence-pinned, path-exact classes); the §4c typed auth refusal (auth:seat-quarantined match-gated); the H-16-REG registry shape (m-2-confirmed literal)
Proof: E2 — go test ./... + the named batteries, RED-first, FILE-captured (incl. the F4 operator-only authority negatives at raw ingress); relay-lint per the pinned expected-output oracle (Global constraints)
No-consumer action: defer — consumers stay on fail-closed legacy reads (two-dimension migration is expressly T4-lane, out of scope)
```

## Global constraints (every task inherits; violations are review blockers)

- Branch `s12-h16-fix` off workspace `main`; edits confined to `frank/` paths + the relay root `frank/.relays/s12/`. Merge is OPERATOR-ONLY via the terminal MERGE-GATE — no s12 relay ever claims a merge.
- `go test ./...` green at every commit; the ten INV-CATALOG laws green at every commit (T15 is continuous, not terminal); commit messages prefixed `s12 IMPL:`.
- Terminal enum byte-exact `{accepted, rejected, held}` everywhere; no INV-CATALOG law touched; no refactor beyond the seam.
- RED-first per task: write the failing battery, run it, FILE-capture the RED output under `frank/.relays/s12/batteries/` (sequence-honest: capture files carry real timestamps in RED→GREEN order), then implement, then capture GREEN.
- Exact byte tokens (rev21's): rejection classes `system-owned` (reason `spec.ID + " is system-owned"`, `Field: spec.ID`), `lane-supplied-system-field`, `non-boot-before-active`, `duplicate-resolution`, `conflicting-resolution`, `stale-resolution`, `unknown-target`, `anchor-target-resolved`, `retry-authority-delta`, **`mint-predecessor-mismatch`** (rev21 `:89`, the ruled chain-tip byte); auth tokens `auth:seat-quarantined`, `auth:invalid-credential`, `root-lock-held`; states `complete|pending|failed|unknown`, terminals `realized-undelivered`, marker state `running_or_unknown`, dispositions `effect-confirmed-realized|effect-confirmed-unrealized`, park reason `retry-ceiling`; headers `hook_contract` (`"1"`), `mint_predecessor`, `admin_provenance` (`"ceremony"`).
- `DERIVED_RETRY_CEILING` = compiled constant, MVP value 8 (§6).
- Rejected records PRESERVE supplied headers as evidence (`rejectAtEdge` behavior retained); fold eligibility requires ACCEPTED (§4).
- UNTOUCHABLE: the frozen exit oracle (`d4580c52…`), everything under `master/exit-fixtures/`, m-2's frozen stage-1 contract, the live courier at `~/frank-live`. H-12 stands: nothing here relaxes the external/untrusted/multi-tenant block.
- Escalation rails (commission §2): (c) any registry/schema change beyond the MUST-CHECK register below → master BEFORE work; (d) any locked-contract change → owning pair + master; (e) any rev21 byte defect → master (void rail; the build HOLDS on that leg and continues on independent legs).
- **The relay-lint oracle, pinned as EXPECTED OUTPUT (F2 — an unqualified "routed clean" is unattainable: the superseded r1 relay's two structural errors are immutable):** (i) exact-file lint returns OK for every live/consumed relay; (ii) a routed invocation's error set contains ONLY the two operator-granted r1 errors (`191828`: the design-gate miss + the unresolvable artifact pair) and prints `OK` for the named target — any OTHER error is a defect; (iii) the resulting nonzero exit is accepted ONLY under the operator's exact `194914` waiver; if the operator denies or changes that disposition, dispatch and the end review HOLD until the ruled disposition is realized.
- Existing S8 semantics UNTOUCHED: the observe guard's precedence and the 31 rows' `lane-supplied-system-field` behavior on non-empty observe-enabled submissions change on NO path (§4; any precedence reorder is a master-routed cross-domain change).

## File structure (the decomposition of record; SCOPE_DIFF at dispatch enumerates exactly these)

```
cmd/frank/main.go                     # -init/-mint lock acquisition; ceremony verb; startup order §6; hook wiring (Tasks 1,2,6,9,11,12,13)
internal/store/lock.go                # UNCHANGED helper (AcquireRoot is already shared-shaped); consumers converge onto it
internal/store/genesis.go             # the v9 fieldspecGenesisDowngrades step (Task 4, F3); store.Init body itself unchanged (locking lives at the callers per Task 1)
internal/store/store.go               # raw redo stream access for the §4a completeness proof (Task 11)
internal/engine/loop.go               # Outcome split; five sites; Class-G dirty; drain; panic path (Tasks 5,6,7,8,9)
internal/engine/outcome.go            # NEW (only if loop.go grows unwieldy): Outcome projection + fold types (Task 5/7; benign split, record in commit msg)
internal/engine/submit.go             # §4b commit-time transition validation; seat_mint chain validation (Tasks 9,10)
internal/engine/derived.go            # NEW: derived-work record family, cursor fold, marker law (Tasks 7,8,9)
internal/fieldspec/validate.go        # presence-pinned predicate (Task 3)
internal/fieldspec/registry.json      # 3 rows + ruled record_kind tokens + version marker (Task 4)
internal/config/config.go             # v9 reader marker + v8→v9 transition allowlist (Task 4, F3)
internal/fieldspec/registry_test.go   # version + exact provenance assertions updated (Task 4, F3)
test/fixtures/s10_fieldspec_test.go   # historical v8 assertions preserved against pinned fixture bytes (Task 4, F3)
internal/seat/binding.go              # typed Resolve refusal; realization proof surface (Task 12)
internal/channel/server.go            # auth:seat-quarantined mapping (Task 12)
internal/store/quarantine.go          # effective-quarantine derivation if it fits the existing seam (Task 12)
internal/tables/tables.go             # open-set reconstruction from canonical bytes; dirty diagnostic row (Tasks 6,7)
internal/recover/recover.go           # recovery-processor stamping + callerless scan re-grounding (Tasks 7,10,11)
cmd/frank/ceremony.go                 # NEW: the offline one-shot ceremony verb (Task 13,14)
test/fixtures/…_test.go               # every battery below; one file per battery family, named h16_*/h26_*/init_lock_*/ceremony_*
frank/.relays/s12/batteries/*.txt|.md # RED/GREEN captures — append-written into the SHARED tree's relay root; NEVER committed by s12 (master's checkpoint cadence banks them on main)
docs/sprints/active/2026-08-17-s12-h16-fix/…  # this plan + ROADMAP (pair-Planner-authored)
```

Exact file boundaries inside `internal/engine/` are the Implementer's (follow existing package idiom; new files recorded per-commit); the PACKAGE set above is the scope fence. Any file outside `frank/` + the relay root is OUT — stop and escalate.

## MUST-CHECK register (interpretation pins; the plan-review confirms or contests each)

These are realization choices this plan pins where rev21/the commission leave the registry/code spelling to the build (pins 1 and 7 are now owner-RULED, not pins of mine). If the Implementer's review — or build-time evidence — shows any pin exceeding the ruled set or contradicting rev21's bytes, STOP that leg and escalate to master per rail (c)/(e).

1. **`record_kind` realization — RULED; the gate is CLEAR** (m-2's ruling `master/relays/step3-h16-h26-lane/DESIGN-planner-m2-20260817-201520.md` §1, m-7 spec-concurred `…202652`, adopted by master `…202541` §1 — **cite m-2's relay as the authority; realize its 1a–1d BYTE-EXACTLY**): (1a) the enum gains `mint-chain-anchor` · `attempt_resolution` · `derived-work-attempt` · `derived-work-transition`, spelled symbol-for-symbol as the design spells them; (1b) `seat_scope.operator` += exactly `{mint-chain-anchor, attempt_resolution}`; the `*` scope untouched; the two `derived-work-*` classes carry NO seat entry (writer-authored, genesis precedent); (1c) the `reopened` carrier = `attempt_resolution`-class records, the §4b commit-time validator typing the verb BY THE RESOLVED TARGET'S TYPE — `resolves` naming a `parked` transition ⇒ the reopen verb (body carries `resolves` ONLY); `resolves` naming a `derived-work-attempt` marker ⇒ the disposition verb (body REQUIRES `disposition` + `evidence_ref`); missing/extra required members reject through the existing commit-time parser discipline (the `ParseSeatMintBody` precedent); an unresolvable `resolves` is §4b `unknown-target`; NO new §4b class, NO new externally visible byte; (1d) the marker string + downgrade step are this pair's to stamp (MUST-CHECK 3 stands). Tasks 4 and 7–15 are UNBLOCKED against this ruling.
2. **Registry `type` token — `string`, literal (R2-corrected):** the live registry supports and uses `type: string` on system-stamped header rows (`slot_in`, `seat_archetype`, `decision_deadline`, `model_name` — `registry.json:161-173`). The three new rows carry EXACTLY the m-2-confirmed shape: `layer: header · owner: system · type: string · fill_constraints: system_only · lineage_role: none`, NO `enum_set`. r1's `text` substitution is retracted.
3. **Version marker + its LIVE CONSUMERS (R2-expanded per F3):** bump `"version"` to `"s12-fieldspec-v9"` AND land every consumer that gates it: (a) `internal/config/config.go` `preflightMemberMarkers` — add `"s12-fieldspec-v9"` to the `ValidateFieldspecReaderMarker` accepted list (`:317-328`); (b) the version-transition allowlist — add the `s10-fieldspec-v8 → s12-fieldspec-v9` pair (`:425-436`); (c) `internal/store/genesis.go` `fieldspecGenesisDowngrades` — a NEW v9 step (marker `"version": "s12-fieldspec-v9"`, exact-match strings including the post-ruling `record_kind` enum line, replacements to the v8 predecessor state, `preserveSource` semantics mirroring the v8 step) so fresh-genesis compatibility holds; (d) the asserting tests — `internal/fieldspec/registry_test.go` (version + the EXACT provenance map) and the `test/fixtures/s10_fieldspec_test.go`/s8-historical fixtures (historical v8 assertions PRESERVED against pinned fixture bytes, never deleted); (e) a fresh-v9-init battery (a genesis from the v9 registry succeeds and pins the right digest). Provenance pinned EXACTLY (flat string map, the registry's only provenance idiom): `{"owner": "m-2", "design_doc_id": "h16-outcome-split-design", "plan_lock_id": "s12-h16-fix-plan", "supersedes": "s10-fieldspec-v8", "note": "s12 H-16-REG realization under the operator-opened step3-h16-h26-lane; three system-stamped header rows + the ruled record_kind tokens"}`.
4. **Genesis predecessor token (§4a):** the first linked pivot's `mint_predecessor` value = `genesis` (the genesis record's fixed RelayID — an existing fixed token; rev21 says "the fixed genesis token" without spelling it).
5. **Ceremony verb spelling:** a new flag on the same binary (e.g. `-recover-seat <seat>` + `-retry-delivery <seat>` invocations of one ceremony mode), NO socket, structurally absent from the seat `submit/project/read` surfaces (§4c). Exact flag names are the Implementer's; the structural-absence assertion is a battery leg.
6. **Named exclusions (report, don't fix):** `-bless`/`store.BlessS8` (`main.go:106-110`) and `operatorSubmit` (`main.go:115-117`) are NOT in the four-leg scope. Task 0 runs a census: if either is a root/config/store WRITER running unlocked, REPORT it in the Task 0 capture as a candidate residual for master — no code change here.
7. **The `mint_predecessor` mismatch rejection byte — RULED; the gate is CLEAR**: the token is **`mint-predecessor-mismatch`**, carried by the design itself at rev21 `:89` (m-7-ruled and rev21-authored under the void rail; m-2 concurred as `failing_edge` class-text owner with the token verified at exactly two occurrences and NO enum registration following, per their standing 8a ruling; the rev21 pair review APPROVED at 0 blocking, `…DESIGN-REVIEW-implementer-20260817-202647.md`). Task 10 is UNBLOCKED; its RED leg asserts this exact token as a fixed writer-stamped `failing_edge` value.

---

### Task 0: Isolated worktree, baseline, census

**Files:** no source/test/registry file modified. Creates: the worktree (code-only), capture files under the shared tree's `frank/.relays/s12/batteries/`.

- [ ] **Invoke `superpowers:using-git-worktrees` (mandatory before implementation)** and reconcile its result with the charter's exact branch name: the worktree HOSTS branch `s12-h16-fix` created off workspace `main` (e.g. `git worktree add <worktree-path> -b s12-h16-fix main`). Record the worktree path AND the exact `main` sha at creation (the BASE of record) in the Task 0 capture. The ONE substrate model, byte-consistent everywhere (F3): ALL code edits, test runs, and commits happen in the worktree on `s12-h16-fix`; the shared workspace tree is never checked out and never SOURCE-mutated by this build, but its relay root `frank/.relays/s12/` IS append-written (relays, INDEX rows, battery captures — by absolute path); s12 commits NO relay/capture copy anywhere (not on the branch, not on `main`) — master's checkpoint cadence banks the shared-tree trail.
- [ ] Re-verify the spec hash from disk: `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` == `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05` (rev21). Mismatch → STOP, escalate (the void rail re-enters and master re-relays).
- [ ] Baseline: `go test ./...` in the worktree's `frank/` — expect GREEN; capture output → `.relays/s12/batteries/t0-baseline-green-<stamp>.txt`. Not green → STOP, report (the branch point must be green).
- [ ] Loci census at branch bytes: confirm `validate.go:35` (`raw != ""`), `loop.go:160-187` (four relabel sites), `loop.go:279-281`-region (`supersededCredentialOutcome`), `main.go:89-117` (`-init`/`-mint` pre-lock), `main.go:118` (`AcquireRoot` for serve). Run the MUST-CHECK-6 census on `-bless`/`operatorSubmit`. Record all in the capture file.
- [ ] Record the two ruling relays in the capture (both LANDED — MUST-CHECK 1: m-2 `201520` §1 via master `202541`; MUST-CHECK 7: `mint-predecessor-mismatch` at rev21 `:89`, pair-approved `202647`). All task gates are clear at dispatch.

### Task 1: R-INIT-UNLOCKED — `-init` under the root lock (leg 4)

**Files:** Modify `cmd/frank/main.go` (the `cfg.Init` branch, `:89-105`). Test: `test/fixtures/init_lock_test.go` (new).
**Interfaces:** consumes `store.AcquireRoot(root) (*store.RootLock, error)` unchanged; produces the locked `-init` path: lock acquisition is the branch's FIRST operation, held through exit.

- [ ] **RED:** three failing legs (spawn real second processes via the existing lock_test.go helper-process pattern, or `go run`/test-binary re-exec):
  1. *concurrent double-init:* two processes race `frank -init` on one root ⇒ exactly one succeeds; the loser exits non-zero with `root-lock-held` (or `ErrGenesisExists` only if it started after the winner's release) and has performed NO config/store write interleave — assert the surviving `<root>` config bytes + genesis `config_digest` are wholly the winner's.
  2. *differing-config:* the two racers carry DIFFERENT config sources ⇒ the surviving store is internally consistent: on-disk config bytes re-digest to exactly the committed genesis `config_digest`; no mixed member set.
  3. *alias-root:* second init addresses the same root via a symlink alias ⇒ still excluded (root identity, not path string).
  Run: `go test ./test/fixtures/ -run TestInitLock -v` ⇒ FAIL today (unlocked `store.Init` check-then-act). Capture → `batteries/t1-init-red-<stamp>.txt`.
- [ ] **Implement:** in `run()`, the `cfg.Init` branch acquires FIRST — before the source loads and the observe check:
  ```go
  if cfg.Init {
      rootLock, err := store.AcquireRoot(cfg.Root)
      if err != nil { return err }
      defer func() { _ = rootLock.Release() }()
      // …existing engine-config/catalog checks, Load, observe check, store.Init(cfg.Root, sources)
  }
  ```
  `store.Init`'s own body is unchanged (the genesis-absence check becomes safe under the lock; the s6 acquire-before-touch law lives at the consumer, matching the serve path's shape at `:118`).
- [ ] **GREEN:** rerun ⇒ PASS; whole suite green; capture → `batteries/t1-init-green-<stamp>.txt`. Commit `s12 IMPL: t1 -init under AcquireRoot (R-INIT-UNLOCKED)`.

### Task 2: H-26 — `-mint` under the root lock (leg 2)

**Files:** Modify `cmd/frank/main.go` (`mintSeat`, `:577-601`; the `cfg.MintSeat` branch stays at `:112`). Test: `test/fixtures/mint_lock_test.go` (new).
**Interfaces:** the F99 shape — `AcquireRoot` is `mintSeat`'s FIRST operation (phase −1), held through exit; the socket-liveness probe becomes a POST-LOCK DIAGNOSTIC.

- [ ] **RED:** failing legs:
  1. *lock-excludes-mint:* while another process holds the root lock, `frank -mint …` fails `root-lock-held` and writes NO binding byte (assert the binding table's bytes unchanged).
  2. *lock-before-probe (ordering proof):* lock held by another process AND a live decoy socket present ⇒ the error is `root-lock-held`, NOT "conductor is serving" — proving the probe did not run pre-lock.
  3. *stale socket:* lock free + dead socket file ⇒ mint proceeds (probe is diagnostic, not exclusion).
  4. *mint-vs-mint:* two concurrent mints ⇒ one wins, the loser `root-lock-held` cleanly, exactly one credential minted.
  Capture → `batteries/t2-mint-red-<stamp>.txt`.
- [ ] **Implement:** reorder `mintSeat`: `AcquireRoot` → defer Release → `socketIsLive` diagnostic ("conductor is serving" retained as a post-lock message) → `requireGenesisTimeMint` → `seat.Open` → `Mint`.
- [ ] **GREEN** + full suite; capture `batteries/t2-mint-green-<stamp>.txt`. Commit `s12 IMPL: t2 -mint under AcquireRoot (H-26)`.

### Task 3: The presence-pinned validator + the shared-impact regression (cross-cutting)

**Files:** Modify `internal/fieldspec/validate.go:35`. Test: `test/fixtures/h16_presence_pinned_test.go` (new).

- [ ] **RED:** (a) a candidate whose header map carries a system-owned key with the EMPTY value (use an existing 31-row member, e.g. `delivery_state`) reaches `Registry.Validate` ⇒ expect a `system-owned` violation — fails today (the `raw != ""` bypass). (b) the SHARED-IMPACT regression: enumerate the population FROM THE REGISTRY BYTES at test time — `layer == "header" && (owner ∈ {system, computed} || fill_constraints ∈ {system_only, computed_result})` — assert the count (31 at this task; 34 after Task 4, parameterized not hardcoded), assert EVERY member rejects present-empty at the validator, and assert the envelope-layer `FROM`/`ROLE`/`relay_id` rows are OUTSIDE the population (`systemOwnedHeader` false on `Layer != "header"`). (c) the PRODUCER-CONTROL leg (R17-F1): the conductor's own post-validation observation output carrying `degradation_notes: ""` (`internal/observe/gate.go:246-257`) still commits and RETAINS the empty member. Capture RED → `batteries/t3-validator-red-<stamp>.txt`.
- [ ] **Implement:** `validate.go:35` → `if present && r.systemOwnedHeader(spec) {` — rejection bytes unchanged (`Field: spec.ID`, class `system-owned`, reason `spec.ID + " is system-owned"`).
- [ ] **GREEN** + full suite (existing S8 fixtures must stay green — the S8-control assertion); capture. Commit `s12 IMPL: t3 presence-pinned system-owned predicate + 31-row regression`.

### Task 4: H-16-REG — the registry realization + its live version-gate consumers (leg 3)

**The MUST-CHECK-1 gate is CLEAR: realize m-2's `201520` §1 (1a–1d) byte-exactly and cite that relay in the changeset.**

**Files:** Modify `internal/fieldspec/registry.json`, `internal/config/config.go` (`preflightMemberMarkers` `:317-328`, the transition allowlist `:425-436`), `internal/store/genesis.go` (`fieldspecGenesisDowngrades` — the new v9 step), `internal/fieldspec/registry_test.go`. Preserve-and-pin: `test/fixtures/s10_fieldspec_test.go` + the s8 historical inverse fixtures (v8 assertions kept against pinned fixture bytes). Test: extend `h16_presence_pinned_test.go` + `test/fixtures/h16_registry_rows_test.go` (new).
**Interfaces produced:** the three header rows every later task's stamping/rejection machinery consumes; the RULED `record_kind` tokens Tasks 7–14 commit under; a v9 registry every live version gate accepts.

- [ ] **RED:** assert rows `hook_contract`/`mint_predecessor`/`admin_provenance` exist with EXACTLY `layer: header · owner: system · type: string · fill_constraints: system_only · lineage_role: none` and NO `enum_set`; assert the `record_kind` enum + `seat_scope` state matches m-2's `201520` §1a/§1b exactly (four members spelled symbol-for-symbol; `seat_scope.operator` += the two; `*` untouched; no seat entry for `derived-work-*` — cite the ruling relay in the test comment); assert version == `s12-fieldspec-v9` + the MUST-CHECK-3 exact provenance map; assert `ValidateFieldspecReaderMarker` ACCEPTS v9 and the transition check ACCEPTS `v8 → v9` (and still rejects `v7 → v9`); assert the genesis downgrade path handles a v9 registry (fresh v9 `-init` succeeds end-to-end and pins the right digest — the fresh-v9-init proof); assert the shared-impact population count is now 34 and each NEW row rejects present-empty at the validator; assert the three new rows joined `LaneSuppliedSystemField`'s registry-driven population (observe-enabled non-empty ⇒ `lane-supplied-system-field` — the S1/S3 wiring, proven end-to-end in Task 15). Capture RED (the reader-marker/transition legs fail against live v8-only gates — that IS the F3 finding made executable).
- [ ] **Implement:** the three row objects; the RULED enum/seat_scope additions; the marker bump + exact provenance map; the `config.go` reader-marker + transition additions; the `genesis.go` v9 downgrade step (exact strings including the post-ruling `record_kind` enum line, replacements to the v8 state, `preserveSource` mirrored from the v8 step); `registry_test.go` updated to the v9 version + exact new provenance; historical v8 fixtures preserved against their own pinned bytes.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t4 H-16-REG rows + ruled record_kind tokens + s12-fieldspec-v9 + version-gate consumers`.

### Task 5: The Outcome split core (§1, §2)

**Files:** Modify `internal/engine/loop.go` (the `Outcome` struct + the five sites + `outcomeFromRecord`); `internal/engine/submit.go` if `Outcome` literals live there. Test: `test/fixtures/h16_outcome_split_test.go` (new).
**Interfaces produced (every later task consumes):**
```go
type Outcome struct {
    State           string `json:"state,omitempty"`             // PRESENT (∈ {accepted,rejected,held}) IFF PostCommitState == "complete"
    DecisionState   string `json:"decision_state,omitempty"`    // fixed at commit, NEVER relabeled
    PostCommitState string `json:"post_commit_state,omitempty"` // {complete,pending,failed,unknown}
    // RelayID, IntakeID, Reason, Detail, Credential, Endpoint — unchanged
}
```
Pre-commit rejections keep one-dimensional truth (`State: rejected`, `PostCommitState: complete`, `DecisionState: rejected` on new emitters — §2: both new members ALWAYS present on new emitters; the mechanical fail-closed rule governs `State` alone).

- [ ] **RED — T1 (five-site relabel battery):** for EACH of the five sites (`completeTurn` fail; `AfterGateResolution` fail; `AfterApprovalResolution` fail; `AfterAccepted` fail; `supersededCredentialOutcome`'s completeTurn fail) inject the failure ⇒ assert the reply carries `decision_state == the committed state` (accepted stays accepted), `post_commit_state == "pending"`, `state` ABSENT, and a truthful typed reason — T6: the `credential-superseded` detail PRESERVED (never swapped for `obligation-error`). Today each returns `State: rejected` ⇒ RED. Capture.
- [ ] **RED — T7 (fail-closed decode):** run the ACTUAL legacy decode (state-only enum match) over a pending-accepted outcome ⇒ error path taken (state absent matches no member); a new two-dimension decode reads both members.
- [ ] **Implement:** the struct + the five sites return `{DecisionState: <committed>, PostCommitState: pending, RelayID, IntakeID, Reason: <truthful>}`; the projection rule (`State` set only when `PostCommitState == "complete"`) lives in ONE constructor (e.g. `outcomeFor(decision, postCommit string, …)`) so no site can hand-roll it; `outcomeFromRecord`/`existingOutcomeForCommand` emit through it. (`post_commit_state: pending` becomes fully honest in Task 6 — at this task the failure paths genuinely have unresolved work, so `pending` is already true; NO reply may claim `complete` after a failed hook.)
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t5 Outcome split — five sites stop relabeling (T1,T6,T7)`.

### Task 6: Class-G truth machinery (§3, §6-part, §7-part)

**Files:** Modify `internal/engine/loop.go` (dirty flag + retry-before-replay + drain), `cmd/frank/main.go` (serve gating + nudge), `internal/tables/tables.go` (dirty diagnostic row). Test: `test/fixtures/h16_classg_test.go` (new).

- [ ] **RED — T17:** Class-G failure ⇒ reply `pending`; immediate duplicate command while dirty ⇒ STILL `pending` (never a silent `complete`); successful retry pass ⇒ `complete`.
- [ ] **RED — T11 (quarantine proof):** failures injected into ALL FOUR Class-G substeps (`CompleteAuto`, GC, tables build/publication, scheduler arming) AND `processQuarantine` ⇒ nothing strands with NO future command: dirty set, a tables/diagnostic row visible, the next event OR the pre-serve pass completes it — no silent strand.
- [ ] **RED — T12:** the real host cannot accept a connection before the drain returns; drain-panic ⇒ no serve.
- [ ] **RED — T9 (nudge):** no nudge on any non-complete reply (state absent — mechanical); NO synthetic nudge on heal; ordinary complete nudge unchanged.
- [ ] **Implement:** `classGDirty` in the loop (same-process truth); EVERY reply path — including `existingOutcomeForCommand` replays — first attempts one Class-G pass while dirty (success clears; failure leaves dirty + reply `pending`; the deterministic projection: ALL replies carry `pending` while dirty); `processQuarantine` failure sets dirty + retries at every loop event + exposes the tables/diagnostic row; the §6 pre-serve synchronous main-goroutine drain (one Class-G pass; channels open gated on drain return; drain-panic ⇒ no serve). Dirty is deliberately NOT durable (cross-restart truth = the pre-serve pass).
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t6 Class-G dirty/retry-before-replay/pre-serve drain (T9,T11,T12,T17)`.

### Task 7: `hook_contract` + the derived open-state + the blind-replay fold (§3a-core, §4)

**Files:** Create `internal/engine/derived.go`; modify `loop.go` (stamping + hydration), `internal/recover/recover.go` (recovery-processor stamping), `internal/tables/tables.go` (open-set reconstruction). Test: `test/fixtures/h16_derived_fold_test.go`.
**Interfaces produced:** `derived.Cursor(rec) []string` (pure function of canonical bytes: `resolves_gate`-accepted ⇒ `{gate, approval}`; accepted `seat_mint` ⇒ `{mint}`); `derived.Fold(records) map[relayID]WorkStatus` with `WorkStatus{Cursor []string; Status string /* pending|failed|unknown|"" terminal-complete */}` consuming ACCEPTED records only, order-free; the marker/transition/resolution record kinds committed as typed system records under the Task 4 tokens.

- [ ] **RED — T2:** restart reconstruction FROM CANONICAL BYTES ALONE — redo + projections DELETED; the fold returns pending/unknown correctly; per-record `hook_contract` classification grandfathers unstamped records (absent ⇒ legacy `complete`); the ADVERSARIAL-ORDER leg: relay-ID lexical order OPPOSING commit order, classification correct from the per-record header alone. Version fail-closed: an ACCEPTED record carrying `hook_contract: "2"` folds `unknown`-class (never v1, never legacy).
- [ ] **RED — T3:** blind heal at drain/top-of-process/replay ⇒ accepted `cursor_advance` records ⇒ empty-cursor terminal; idempotency (double-run == single-run).
- [ ] **RED — T5:** two concurrent open items durable + reconstructible, neither erased.
- [ ] **RED — T14:** the executable `intake.Unconsumed` assertion (the decision consumes its own journal entry; the DERIVED open-state alone drives recovery).
- [ ] **RED — T-R5F2:** first-H16-startup recovery — an unconsumed gate resolution AND an unconsumed seat mint processed by the recovery processor with failure/crash after the decision commit and before hook completion ⇒ next restart projects `pending`/`failed`/`unknown` under the fold, NEVER grandfathered `complete` (the recovery processor stamps its commits — same binary, no activation window).
- [ ] **Implement:** writer-side stamping (loop AND recovery processor stamp ACCEPTED hooked-class records exactly `"1"`; every other ACCEPTED record carries NO `hook_contract`); `derived-work-attempt`/`derived-work-transition{cursor_advance{completed_hooks}}` commits through the ordinary serialized writer; the fold rule (§3a: cursor = hook list − ∪(accepted advances); status precedence failed-park > unknown-marker > pending; terminal on empty cursor); the loop's in-memory set = a scheduling cache hydrated from the fold; blind-replayable retry set (completeTurn/AfterGateResolution/AfterApprovalResolution) driven at drain/top-of-process/replay.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t7 hook_contract + derived open-state + blind fold (T2,T3,T5,T14,T-R5F2)`.

### Task 8: The marker law + `AfterAccepted` delivery + §4a evidence (§3a-delivery, §4a-part)

**Files:** Modify `internal/engine/derived.go`, `loop.go`, `cmd/frank/main.go` (`completeSeatMintBinding` call sites). Test: `test/fixtures/h16_marker_law_test.go`.

- [ ] **RED — T4 (marker-law cuts):** before marker (`pending`, re-runnable) / marker-commit-fails (effect NOT entered, `pending`, next caller-present replay retries) / after marker before effect (mint: EVIDENCE-driven — `realized_mint_ref` absent ⇒ `not-started`, caller-present resumes) / inside-or-after effect before advance (`realized_mint_ref == pivot relayID` ⇒ `realized-undelivered`, failed-class, NO re-run, NO extras) / after advance before reply (replay serves `complete`, extras ABSENT — extras exist ONLY on the reply whose terminal advance committed).
- [ ] **RED — T8:** earlier hook fails before an accepted `seat_mint`'s delivery ⇒ `pending`, NO mint; unrelated commands NEVER mint (negative — `AfterAccepted` executes only inside `process` for a command resolving by intake ID/content hash to the source decision); same-intake replay: marker → effect (`MintOrReplace`) → terminal advance → extras on THAT reply.
- [ ] **Implement:** the delivery ordering (marker durable BEFORE any non-blind effect; fresh-instance markers `{source_relay_id, hook, state: running_or_unknown, predecessor}`; a new marker legal only when the prior instance is resolved, naming that resolution); the §4a evidence fold for the mint (consult `realized_mint_ref` FIRST); `not-started` never inferred after an unproven attempt; caller-present-only delivery.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t8 marker law + delivery path + §4a evidence (T4,T8)`.

### Task 9: §4b transitions, §5 operator records, ceiling/parks, panic path

**Files:** Modify `internal/engine/submit.go` (commit-time validation), `derived.go`, `loop.go` (panic recover + ceiling). Test: `test/fixtures/h16_transitions_test.go`.

- [ ] **RED — T13:** ceiling (8th failure) ⇒ durable `parked{retry-ceiling}` ⇒ `failed`; operator `reopened` resumes; reopened → ceiling ⇒ a NEW park (instance-scoped); operator resolutions ONE-SHOT — duplicate/conflicting/stale/unknown-target each commit REJECTED with its exact typed class; redo/projections deleted + REPEATED rebuilds from SHUFFLED record slices ⇒ identical fold + anomaly set.
- [ ] **RED — T16:** fresh-instance identity: unknown(A) → `effect-confirmed-unrealized`(A) → attempt B → crash ⇒ fresh unknown(B) across restart; conflicting resolutions with relay-ID order OPPOSING commit order — the LATER commits REJECTED with its typed class (commit-time, not fold-time).
- [ ] **RED — T10:** post-commit panic returns `{decision_state, post_commit_state per the durable fold}` — never `faultOutcome`'s shape, never a relabel (the recover closure distinguishes pre/post-commit via the committed relayID in scope at the defer); pre-commit panic unchanged.
- [ ] **RED — the operator-only AUTHORITY negatives (F4; rev21 §5's "operator-ONLY, honestly enforced" made executable at raw ingress):** a NON-operator seat (a worker; plus one role-spoof-resistant variant — e.g. a worker credential submitting with a claimed operator role, which the channel-stamped FROM/role machinery must override) raw-submits BOTH `attempt_resolution` shapes — the reopen shape (`resolves` → a park) AND the disposition shape (`resolves` → a marker, with `disposition` + `evidence_ref`) — each otherwise valid ⇒ each commits REJECTED through the EXISTING ruled/registry authority path (the `record_kind` seat_scope machinery from Task 4's ruled realization), asserting the exact EXISTING violation byte the live path emits (read it at source before writing the leg; if NO existing class can express the refusal, STOP and route that locked-byte gap to master — never invent one), supplied evidence preserved per the rejected-record discipline, and ZERO fold/state effect (no reopen, no disposition, no instance transition). The POSITIVE control: the operator seat submitting the same two otherwise-valid shapes commits ACCEPTED with the expected fold effect.
- [ ] **Implement:** §4b commit-time validation against the LIVE fold inside the serialized writer (valid ⇒ ACCEPTED; else REJECTED with the fixed class per case — the disposition IS canonical bytes); operator records arrive as ordinary governed submits from an operator-role seat (channel-stamped FROM + role, NEVER `From: system`) under the RULED m-2 `201520` §1c carrier: BOTH §5 operator verbs commit as `attempt_resolution`-class records, the validator typing the verb BY THE RESOLVED TARGET'S TYPE (`resolves` → a `parked` transition ⇒ reopen, body carries `resolves` ONLY; `resolves` → a `derived-work-attempt` marker ⇒ disposition, body REQUIRES `disposition` + `evidence_ref`; missing/extra required members reject via the existing commit-time parser discipline, the `ParseSeatMintBody` precedent; unresolvable `resolves` ⇒ `unknown-target`; NO new §4b class); authority + instance state validated together; the belt-and-suspenders CONFLICTED fold (two accepted resolutions on one instance ⇒ `unknown`-class); `DERIVED_RETRY_CEILING = 8` in-memory per process lifetime.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t9 §4b commit-time transitions + §5 operator records + panic (T10,T13,T16)`.

### Task 10: The `mint_predecessor` pivot chain (§4a-chain)

**The MUST-CHECK-7 gate is CLEAR: the mismatch RED leg asserts the ruled byte `mint-predecessor-mismatch` exactly (rev21 `:89`) — a fixed writer-stamped `failing_edge` value; NO enum registration follows (m-2's standing 8a ruling).**

**Files:** Modify `internal/engine/submit.go` (chain validation at commit), `cmd/frank/main.go` + `internal/recover/recover.go` (callerless scans re-grounded on the tip predicate). Test: `test/fixtures/h16_pivot_chain_test.go`.

- [ ] **RED — T18 (head leg):** pivots A then B with B's relay ID sorting BEFORE A's; redo/projections deleted; restart ⇒ B ALONE is current (the chain-tip set predicate: the unique accepted pivot no other accepted pivot names as predecessor); an unresolved marker on B blocks BOTH callerless paths; authorized repair realizes B; A fails the first post-restart `Resolve`; no older pivot ever re-minted. PLUS the recovery-processor/`completeMissingSeatMintBindings` crash fixtures: rescan repairs unrealized marker-less pivots; realized pivots untouched; matching vs non-matching `realized_mint_ref`; the disjointness boundary (no callerless repair while a live marker is unresolved).
- [ ] **Implement:** writer stamps every accepted `seat_mint` with `mint_predecessor` = the seat's prior accepted pivot relayID (or `genesis`, MUST-CHECK-4) from live chain state at commit; commit-time validation (a candidate whose implied predecessor is not the current tip ⇒ REJECTED, typed class); `latestSeatMintPivots` re-grounded on the tip predicate; fail-closed conflict at rebuild (shared predecessor/cycle/broken link ⇒ CONFLICTED ⇒ no callerless repair — quarantine governs in Task 12); both callerless paths scan canonical records immediately before acting and SKIP unresolved-marker pivots.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t10 mint_predecessor chain + tip predicate (T18)`.

### Task 11: The upgrade fold + completeness proof (§4a-upgrade)

**Files:** Modify `internal/store/store.go` (RAW pre-dedup redo stream access), `cmd/frank/main.go` (the startup anchor pass in the §6 pinned order), `internal/engine/derived.go` (mixed-history rebuild). Test: `test/fixtures/h16_upgrade_fold_test.go`.

- [ ] **RED — T-R7F1:** zero/one/two legacy pivots × {binding realizes the true latest; binding realizes an older pivot while a newer accepted pivot is unrealized} × opposing relay-ID order × the REDO EVIDENCE MATRIX {proven-complete; absent entirely; malformed; syntactically valid PARTIAL omitting the TRUE LATEST pivot; partial omitting an OLDER pivot; a duplicated pivot entry; a segment-numbering gap} ⇒ ONLY proven-complete commits the system anchor (to the true tip); EVERY other case commits NO automatic anchor and quarantines; PLUS crash-mid-pass restart (anchored seats skipped via the canonical fold, the rest recomputed from a fresh snapshot, never double-anchored, redo disposal GATED throughout).
- [ ] **Implement:** the per-seat completeness proof EXACTLY as §4a pins it (over the RAW pre-dedup entry stream: every member of the canonical legacy pivot set P appears EXACTLY ONCE; segment sequence intact — gap/unreadable/torn-NON-final ⇒ INCOMPLETE; entries outside P ignored); legacy tip = the LAST member of P in redo order, iff complete; zero legacy ⇒ genesis token; exactly one ⇒ canonically unambiguous; the pass runs in the SERIALIZED STARTUP WRITER PHASE in the pinned order (recovery → upgrade anchors → pre-serve binding completion/repair → loop start → channels open), snapshot-once per pass, idempotent across crashes; redo GC gated while any multi-legacy seat lacks an anchor; two accepted anchors for one seat ⇒ CONFLICTED fail-closed.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t11 upgrade completeness proof + anchor pass (T-R7F1)`.

### Task 12: Effective quarantine + the typed `Resolve` contract (§4c-auth)

**Files:** Modify `internal/seat/binding.go` (`Resolve` widening), `internal/channel/server.go` (token mapping), `cmd/frank/main.go` (pre-serve derivation + publication), `internal/store/quarantine.go` if the derivation fits there. Test: `test/fixtures/h16_quarantine_test.go`.

- [ ] **RED — T-R6F2b:** chain-conflict fixtures ⇒ CONFLICTED seat ⇒ the already-bound superseded credential FAILS the first post-restart `Resolve` with typed `auth:seat-quarantined`; NO affected channel opens before canonical resolution; other seats serve; operator anchor/disposition ⇒ repair ⇒ quarantine lifts.
- [ ] **RED — R9-F1 legs:** a CONCURRENT auth attempt racing the resolving commit + crash/error cuts AFTER the anchor commit / AFTER canonical fold publication / BEFORE and AFTER binding persistence — every pre-realization attempt with the old credential ⇒ `auth:seat-quarantined`; repair error or crash retains quarantine across restart; the pre-serve pass must PROVE realization before any clear. The LEAK RULE: a NON-matching credential probe gets `auth:invalid-credential`, never the quarantine class.
- [ ] **Implement:** effective-quarantine = derived state over (canonical chain-status fold × durable binding row): CONFLICTED/unresolved OR resolved-tip-unrealized; computed BEFORE channels open, republished by the serialized writer; `Resolve` widens from `(SeatMeta, bool)` to a typed refusal carrier (`auth:seat-quarantined` ONLY on bound-credential match + effective quarantine); clearing ONLY after `MintOrReplace` durably replaced credential + `realized_mint_ref` — one atomic swap of published auth state, post-realization the old credential draws `auth:invalid-credential`; the §9-entry-12 carve-out (quarantine blocks the SEAT's auth, never the conductor).
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t12 effective quarantine + typed Resolve (T-R6F2b,R9-F1)`.

### Task 13: The offline recovery ceremony (§4c-ceremony)

**Files:** Create `cmd/frank/ceremony.go`; modify `cmd/frank/main.go` (verb dispatch — AFTER the phase −1 lock the verb itself takes). Test: `test/fixtures/ceremony_test.go`.

- [ ] **RED — T-R8F2 (the three conflict shapes):** ordinary conflicted worker (live operator anchors in-band); conflicted operator WITH a second live operator credential (the other operator anchors in-band); the no-in-band-path case ⇒ the OFFLINE ceremony: refused for a non-quarantined target; accepted for the effectively-quarantined seat; canonical `admin_provenance: "ceremony"` record committed; repair realizes; credential returned in the ceremony's own synchronous reply — each shape proves the old credential NEVER authenticates before DURABLE REALIZATION.
- [ ] **RED — R9-F2 legs (lock-first throughout):** NO ceremony root/store/binding/recovery operation AND NO socket diagnostic before the successful `AcquireRoot`; the STARTUP RACE cut (ceremony vs conductor started concurrently ⇒ exactly one wins; the loser returns `root-lock-held` having performed ONLY the operations intrinsic to `AcquireRoot` itself — root-dir create, lock-file open, flock attempt, holder-metadata read — NO canonical/binding/projection mutation); STALE-SOCKET (post-lock diagnostic warns, lock governs); ALIAS-ROOT (lock excludes via root identity); TWO CONCURRENT CEREMONIES (one wins, one `root-lock-held` cleanly); CONDUCTOR-START while ceremony holds the lock (blocks/refuses at the lock); duplicate/replay invocations; crash cuts from intake through anchor commit, binding repair, credential reply ⇒ EXACTLY ONE canonical anchor + ONE realized credential in every interleaving.
- [ ] **RED — R10-F2 matrix + R9-F3:** existing-anchor/unrealized ⇒ repair only, NO second anchor; unique-tip/unrealized ⇒ repair only, NO selector; existing-anchor/realized ⇒ delivery rule only; conflict-with-no-anchor ⇒ exactly one anchor then repair; an anchor against EACH resolved variant ⇒ REJECTED `anchor-target-resolved` with the anomaly row asserted. Role-flip fixtures (R9-F3): candidate pivots DISAGREEING on `role`/`is_operator`, binding row reflecting each in turn ⇒ SAME eligibility/recovery result every variant (custody basis, never role bits pre-selection); no authority-disagreement case unrecoverable.
- [ ] **Implement:** the one-shot verb (MUST-CHECK-5): phase −1 `AcquireRoot` held through exit (the helper's ceremony consumer — the F99 design driver); post-lock socket diagnostic; the serialized startup writer phase run by the ceremony itself (recovery → validation); the anchor/disposition committed as an ordinary canonical record with writer-stamped `admin_provenance: "ceremony"`; repair to durable realization; synchronous credential reply; exit. NO socket surface; structurally absent from `submit/project/read` (battery-asserted). Eligibility = the §4c state/action matrix over (canonical chain state × durable realization) — never an aggregate boolean, never role bits.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t13 offline ceremony + state/action matrix (T-R8F2,R9-F2,R9-F3,R10-F2)`.

### Task 14: The delivery-retry rotation (§4c-retry)

**Files:** Modify `cmd/frank/ceremony.go`, `internal/engine/submit.go` (`retry-authority-delta` commit-time check). Test: `test/fixtures/ceremony_retry_test.go`.

- [ ] **RED — R10-F1:** crash BEFORE any reply byte / after a PARTIAL reply / after a FULL reply without acknowledgement + REPEATED operator retries ⇒ every credential value ever created maps to exactly ONE canonical pivot (each retry commits its own ceremony-provenance rotation `seat_mint` pivot through the ordinary `mint_predecessor` machinery — never two values under one `realized_mint_ref`); exactly ONE current credential valid; every invalidated predecessor fails `Resolve`. Retry-scope legs: a healthy in-band-delivered seat REFUSED; the `realized-undelivered` callerless-repair state and both ceremony-provenance states ACCEPTED.
- [ ] **RED — R11-F1 (authority legs):** retries over worker, planner, and operator-role tips ⇒ every retry pivot's `seat`/`role`/`is_operator` byte-identical to its predecessor (copied by ceremony CODE from the canonical tip — the form carries target seat + `retry_reason` ONLY, rejecting authority-field input outright); attempted role/`is_operator` deltas on a ceremony-provenance pivot ⇒ REJECTED `retry-authority-delta`.
- [ ] **Implement:** the retry predicate (§4c: tip is ceremony-provenance accepted `seat_mint`, OR ceremony-anchor-realized, OR canonical delivery disposition `realized-undelivered`); the rotation pivot (ordinary machinery, anchor count untouched); the three-layer no-override proof; `retry_reason` a plain non-authority body member.
- [ ] **GREEN** + full suite. Commit `s12 IMPL: t14 delivery-retry rotation + retry-authority-delta (R10-F1,R11-F1)`.

### Task 15: T-R6F3 — the 48-executed-case forged-header unit + provenance legs

**Files:** Test-only: `test/fixtures/h16_forged_headers_test.go` (new; the raw `channel.Client` drive), frontend no-call legs beside the existing client-gate fixtures. Any code deltas discovered here are DEFECTS in Tasks 3–14's work — fix at the owning seam, never with matrix-local special cases.

- [ ] **The 24-cell × 2-shape unit, EXECUTED (R19-F1 — no under-run):** each of the three headers (`hook_contract`, `mint_predecessor`, `admin_provenance`) × {non-empty, present-EMPTY} × the four §4 states, EACH cell executed with BOTH candidate shapes (hooked AND unhooked) = **48 executed candidate cases** through the RAW shared `channel.Client`, asserting: the path-exact class (S1 non-empty ⇒ `lane-supplied-system-field`, S1 empty ⇒ `non-boot-before-active`; S2 both ⇒ `non-boot-before-active`; S3 non-empty ⇒ `lane-supplied-system-field`, S3 empty ⇒ `system-owned`; S4 both ⇒ `system-owned`); rejected-record header PRESERVATION where the path uses `rejectAtEdge`; canonical NULLITY (no membership, pivot edge, ceremony provenance, or unknown derived work from any rejected record); NO case commits ACCEPTED. Matrix preconditions per §4 (healthy machinery, valid digest, otherwise-valid candidate).
- [ ] **The conforming-frontend NO-CALL legs (F98):** the same forged fields through each conforming native/MCP frontend ⇒ typed `schema_invalid`, ZERO conductor calls; where a frontend's input representation CANNOT express a present-empty member, that impossibility is ASSERTED, not skipped.
- [ ] **The S8-CONTROL legs:** the existing observation fields' locked non-empty `lane-supplied-system-field` semantics asserted UNCHANGED; the producer-control leg (Task 3) re-asserted in-matrix.
- [ ] **R11-F2 (provenance legs):** forged `admin_provenance` raw ⇒ committed rejection with the state's exact matrix class + nullity; conforming ⇒ `schema_invalid` NO-CALL; an ordinary ACCEPTED `seat_mint` (header absent) NEVER satisfies the retry predicate while an offline-ceremony pivot and a ceremony anchor-repair realization DO; an accepted future provenance value folds fail-closed `unknown`. Plus the `hook_contract` rollback leg: forged `{absent, "1", "2"}` raw-client legs; accepted hooked records carry exactly `"1"`; ACCEPTED `"2"` folds `unknown`.
- [ ] **GREEN** + full suite; capture the full matrix run (48 case names visible in `-v` output) → `batteries/t15-r6f3-green-<stamp>.txt`. Commit `s12 IMPL: t15 T-R6F3 48-case unit + provenance legs (R11-F2)`.

### Task 16: Consumer table, replay coherence, final sweep

**Files:** Verify-only over `cmd/` + `internal/` (fix-at-owner for any delta): `main.go:337-352` nudge, `prompter.go:81-99`, `resummon.go:228-249`, `mcp.go:240-248`, MCP/native/shared-client forwarding. Test: `test/fixtures/h16_consumers_test.go`.

- [ ] **§8 columns asserted per consumer:** state-only consumers fail CLOSED on non-complete (nudge silent; prompters `fail()`→fallback; resummon error→retry-later idempotent via content-hash dedup; MCP re-render: no premature normalization, raw two-dimension outcome forwards); forwarding byte-transparent with both dimensions; NO consumer is migrated to two-dimension reads (that is T4-lane work — assert current fail-closed behavior, change nothing).
- [ ] **§7:** healing replays report `complete`; `not-started` delivery runs caller-present on THIS reply; the heal nudge stays DROPPED (T9 re-run).
- [ ] **§5a census re-walk:** each route's disposition matches the table (incl. the OUT-of-scope `processQuarantine` proof row — covered by T11).
- [ ] **T15:** the ten INV-CATALOG laws green (`go test ./test/fixtures/ -run INV -v` or the catalog's actual harness name — confirm at Task 0); full `go test ./...` green; captures → `batteries/t16-final-green-<stamp>.txt`.
- [ ] **Trail check (the r3 replacement for the r2 substrate commit — master's checkpoint cadence banks the substrate on `main`, observed at `2cb9a0c`/`414ee38`, so the branch carries CODE ONLY and no substrate copy is committed):** run relay-lint per the pinned expected-output oracle (Global constraints) and capture the literal output: exact-file OK for every live relay; the routed error set exactly the two operator-granted r1 errors and no other; the nonzero exit noted against the operator's `194914` waiver.
- [ ] Commit `s12 IMPL: t16 consumer table + coherence sweep (T7-full,T9,T15)`. File the IMPL report relay (per the commission: literal `git status --short` + branch@sha + battery index), TO s12.planner, CC master.orchestrator-planner, master.orchestrator-reviewer, operator.

---

## Battery → task map (every rev21 §10 battery bound exactly once)

T1→5 · T2→7 · T3→7 · T4→8 · T5→7 · T6→5 · T7→5,16 · T8→8 · T9→6,16 · T10→9 · T11→6 · T12→6 · T13→9 · T14→7 · T15→16(continuous) · T16→9 · T17→6 · T18→10 · T-R5F2→7 · T-R6F3→15 · T-R6F2b→12 · T-R7F1→11 · T-R8F2→13 · R9-F1→12 · R9-F2→13 · R9-F3→13 · R10-F1→14 · R10-F2→13 · R11-F1→14 · R11-F2→15. Lane legs: init×3→1 · mint×4→2 · validator/31-row/producer→3,4 · marker/transition/downgrade/fresh-v9-init (F3)→4 · the `mint-predecessor-mismatch` negative (rev21 `:89`)→10 · the operator-only authority negatives + positive control (R3-F4)→9.

## Exit / acceptance criteria (the end-review checks these; nothing else exits the lane)

1. All four legs realized within rev21's Bounds; every battery above green with sequence-honest RED→GREEN captures on file.
2. `go test ./...` green; ten INV-CATALOG laws green; focused packages named in the end-review relay (commission §2(h)).
3. Byte-exact tokens per Global constraints; rejected-record header preservation; ACCEPTED-only fold eligibility.
4. Zero edits outside `frank/` + `frank/.relays/s12/`; zero exit-oracle/exit-fixture bytes moved; zero INV law bytes moved; S8 precedence unchanged.
5. Every MUST-CHECK pin either confirmed in review or ruled before its task ran (pins 1 and 7 are RULED — Task 4/9 realize m-2 `201520` §1 byte-exactly; Task 10 asserts `mint-predecessor-mismatch`).
6. Relay-lint per the pinned expected-output oracle (Global constraints), at dispatch time and again at the end review; the branch carries code only (the substrate banks on `main` under master's checkpoints).
6a. Governance truth at lane exit (F1): the pair's OWN dispatch gates are the plan approve + the operator's `194914` disposition — AND, independently, the rev21 owner/join tail (the m-1 then m-2 exact-hash delta reruns + the refreshed Master+VP join, pre-sequenced by master `202541` §3) remains OPEN under master's sequence; it does not gate this pair's build tasks, but the lane's exit (the master+VP end review → MERGE-GATE) observes it, and the pair's merge-decision return must state its then-current status honestly.
7. The pair's merge-decision return TO master carries branch@sha + the literal test evidence; NO merge claim (operator MERGE-GATE terminal).

## Out of scope (expressly)

Everything Step-4 (permission/authority system, m-5 ceiling, routing execution, consumer two-dimension migration); any fixture/manifest byte of the frozen exit oracle; any fold; the live courier redeploy; `-bless` and `operatorSubmit` (census-report only, MUST-CHECK-6); any INV-CATALOG change; any m-2 stage-1-contract change; any observe-guard precedence change.
