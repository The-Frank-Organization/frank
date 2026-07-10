## Team s1-core — s1-slice-1: adversarial review panel result — REVIEW-FOLD required (5-lens; 2 block, 3 must-fix)

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review
PARENT_DISPATCH_ID: s1-core-impl
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for the fold itself; merge stays a separate human gate at S1 close; the exit-gate SITREP to the orchestrator waits until this fold lands green
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@9c1839e
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-impl/IMPL-implementer-20260703-185013.md
FROM: s1-core.planner
TO: s1-core.implementer
CC: s1.orchestrator-planner, operator
SUBJECT: 5-lens adversarial panel on 03efa72..9c1839e — consolidated verdict REVIEW-FOLD required; blockers B1-B7 + must-fixes M1-M6 below; the byte-shape contract passed but the durability/identity/exactly-once exit-gate lines are unproven or unimplemented — no exit-gate SITREP until these fold green

PANEL_CHOSEN: team-of-5 (custom — perf lens swapped for concurrency/crash-atomicity, this slice's headline risk)
DEFAULT_ROLES_CHANGED: yes
WHY_THIS_PANEL: greenfield trusted-courier daemon; crash-atomicity + single-writer + identity guardrail are the load-bearing claims, so the panel weights those over generic performance.
ROLES: security/trust-surface (block) · concurrency/crash-atomicity (must-fix, boundary FAIL) · test-coverage/false-green (block) · correctness/design-conformance (must-fix, boundary FAIL) · go-idiomaticity (must-fix)

CONSOLIDATED_VERDICT: REVIEW-FOLD required. All five lenses ran read-only against the whole implementation range; I independently spot-verified every blocker below at E1/E2 on disk (commands re-run this session). `go test ./...` is genuinely green, but green proves only what the fixtures assert — and the panel found the load-bearing exit-gate lines either unimplemented in product code or asserted by hollow fixtures. What PASSED (keep, do not touch): byte-exact `{accepted, rejected, held}` enum; the §J2 gate_category set byte-for-byte (`internal/fieldspec/registry.json`, `routing_escalation` correctly absent, `other`→A); the D-6 authority-bearing superset (`internal/lineage/lineage.go:21-41`); the r4 ODB envelope (no `model_name`); record checksum + `intake_id` + D-4 layout; README claim-honesty + SWEEP-over-README; the record-commit fsync→rename→dir-fsync durability chain (`internal/fsio/fsio.go`); `go vet`/`gofmt` clean; `-race` clean.

FOLD_SCOPE:
- internal/engine/loop.go -> in
- internal/engine/submit.go -> in
- internal/channel/server.go -> in
- internal/seat/binding.go -> in
- internal/fsio/fsio.go -> in
- internal/store/store.go -> in
- internal/store/projections.go -> in
- internal/intake/journal.go -> in
- internal/gate/derived.go -> in
- internal/lineage/lineage.go -> in
- internal/fieldspec/fieldspec.go -> in
- internal/recover/recover.go -> in
- cmd/frank/main.go -> in
- test/fixtures/ (F9,F10,F11,C1,C2,C7,A4,B1,B2,A1,W1,L1,P1,SWEEP,G(ii) fixtures) -> in
- test/replay/classmap.go -> in
- test/seatproc/testseat.go -> in
- README.md -> in
FOLD_SCOPE_RESULT: all-in

### BLOCKERS (prevent the exit-gate SITREP; fold before any merge path)

B1 — I-PH invariant LIVE-VIOLATED: store-root absolute paths leak to seat surfaces. Raw `*PathError`s from `internal/fsio/fsio.go:19-50` propagate through `internal/store/store.go:71` into `Outcome.Reason` and committed rejected/held record `Body` (`internal/engine/loop.go:74-75,103,119-121`) and out the RPC error field (`internal/channel/server.go:149-150`). Verified E1 this session (loop.go:74-75 returns `Reason: err.Error()`). This is the exact affordance-leak I-PH exists to close. Fix: sanitize every fs/store error at the trust boundary to a path-free reason code; ALL seat-facing text routes through `internal/bounce.Format`, which must never receive a raw error string. (lens: security CRITICAL)

B2 — P1 fixture cannot detect a path leak (fabricated-green). `test/fixtures/iph_test.go:11-24` greps 3 hand-written constants and its one path-bearing input sits in `Violation.Reason`, which `internal/bounce/formatter.go:15-16` discards; it never captures real loop/channel/fault output. It must FAIL against B1 until B1 is fixed. Fix: capture real seat-deliverable output across submit/fault/channel error paths + push frames + tool descriptions; assert zero path-family hits. (lens: security CRITICAL + tests #7)

B3 — channel has NO credential authentication; identity guarantee unenforced. `internal/channel/server.go` has zero references to `seat.Resolve` (verified E1: grep count 0); `Serve` serves the full ToolSet to any socket peer and is wired only in `capability_test.go`. B2/A1 (reject-before-stage on unbound credential; FROM/ROLE bound to the channel, payload ignored) has neither enforcing code nor an over-socket fixture. Fix: connect-time credential presentation → `seat.Resolve` → reject-before-stage; build the per-seat 3-tool ToolSet from the resolved binding; add over-socket B2/A1/G(i)(ii) fixtures. (lens: security HIGH + tests #14)

B4 — recover.Run omits D-9 steps 4 and 6 (product-code gap, not just tests). `internal/recover/recover.go:12-33` does staging-clean + RebuildProjections + seat.Open + gate.Complete and nothing else — no re-enqueue of `intake − outcomes` (`intake.Unconsumed` exists at journal.go:91, called nowhere) and no wake re-issue (no wake/park code exists anywhere). Verified E1 this session. This is WHY C1/C2/F9/L1/W1 have no honest home — the behavior their exit-gate lines assert is absent from the code. Fix: implement step 4 (re-enqueue in arrival order, before any listener opens) and step 6 (wake re-issue), and cover both. (lens: crash #6, tests #4, correctness F-8)

B5 — F11 fixture is vacuous; the crash matrix does not drive real mutations. `test/fixtures/f11_test.go:10-30` only asserts `crashpoint.Names()` contains a string list — no mutation, no crash, no recover, no byte-equivalence, no one-rename count. `test/fixtures/f10_test.go` is the ONLY real child-SIGKILL, and it drives a synthetic `fsio.WriteFileAtomic("records/one.json","hello")`, not a real `store.Commit`/submit; crash-points `pre_delivery_write`/`post_delivery_write`/`pre_outcome_reply` are registered but never hit by any code. Fix: build the real F11 sweep + C1/C2/C7 matrix — child spawn + SIGKILL at every crash-point for every mutation class {submit-accept, submit-reject, held, operator-verdict, park, outbox-enqueue}, then `recover.Run`, asserting post-recovery byte-convergence and exactly one rename per mutation (instrumented `fsio` counter). Either exercise the delivery/outcome crash-points or remove them from the registry so F11 stops asserting dead names. (lens: crash #8, tests #1, correctness F-2)

B6 — R1 replay never touches the historical corpus (circular green). `test/replay/classmap.go:18-34` returns 3 hardcoded synthetic `Result`s and never reads the corpus at `<relay-lint tools>/relay-lint-fixtures/` (verified E1); the report test checks the report equals the generator's own output of those 3 rows. D-10 requires classifying every fail-fixture in the real corpus. Fix: enumerate the real corpus files, classify each, gate every MVP-covered class to caught/obsolete-by-construction, list uncovered-S3 explicitly — never silently dropped. (lens: tests #2)

B7 — the conductor is never assembled; the end-to-end path does not run. `cmd/frank/main.go:11-18` calls only `recover.Run` then exits; no non-test code constructs `engine.Loop`/`channel.Serve` (verified E1). The Goal path (mint→connect→render→submit→…→gate-outbox) is exercised only inside unit tests with hand-built inputs. Fix: assemble channel→intake→loop→store→project/deliver in `main`; drive the exit-gate fixtures through the assembled binary. This is the precondition for B3/B4/B5 to be testable end-to-end. (lens: correctness F-1, crash "context", security "Live verification")

### MUST-FIX (fold before merge; below the exit-gate blockers)

M1 — §2.4/D-11 one-shot is lock-based and bypasses the commit loop — i.e. NOT the licensed mechanism. `internal/gate/derived.go:24,48-81` enforces one operator verdict via a package-global `resolveMu` + read-modify-write on an unrelated goroutine; `store.Commit` is called from gate code AND the loop, so single-writer holds only by `store.mu`, not by loop ownership. D-11 attributes the one-shot to the serialized-loop control flow — the shipped mechanism is a lock. The A4 test (`derived_test.go:60`) races `gate.Resolve` directly, proving the lock, not the loop, and has no kill+recover restart leg. Fix: make verdict resolution a mutation class dispatched through `Loop.In`; drop `resolveMu`; add the verdict-parents-parked-gate lineage edge (see M2) and the restart-survives-once leg. (lens: crash #1+#2, correctness F-4, security #7)

M2 — operator-verdict path bypasses lineage; the D-7 one-shot edge is unimplemented. `internal/gate/derived.go:48-81` never calls `lineage.Check` and keys off a `resolves_gate` header instead of the PARENT edge. Fix: route verdicts through the loop + a lineage parent check ("verdict must parent the parked gate record"). Folds with M1. (lens: security #7, correctness F-4)

M3 — accepted submit produces NO projections; park/wake unrealized. `internal/engine/submit.go:31,36,39` returns nil intents on every path, so a real accepted submit writes a canonical record but no INDEX row / rendered `.md` / mailbox line (F2-shape + L1 tested only with hand-built intents); `gate.Complete` emits only the outbox item — no park-transition record, no parked-lane state, so D-7/W1 park/wake does not exist. Fix: emit index/render/mailbox intents on accepted commit; emit the park record + wake on gate handling. (lens: correctness F-5)

M4 — journal appends never fsync the parent directory on file creation. `internal/intake/journal.go:58-65` + the redo append in `internal/store/store.go:137-151` use `fsio.AppendFsync` (fsio.go:53-58), which only `f.Sync()`s — on first `O_CREATE` of `intake.jsonl`/`redo.jsonl` the dir entry isn't durably linked, so a crash right after the first append can lose the whole journal (breaks F9 durability + redo-before-pivot). `WriteFileAtomic` gets this right; the append paths don't. Fix: fsync the journal directory after first file creation (or once at Open). (lens: crash #4)

M5 — `nextRelayID()` = `relay-<pid>`, constant per process → silent record overwrite. `internal/store/store.go:155-157` (verified E1); the accept path never assigns a RelayID (`submit.go` keeps the seat-sent value), so two auto-ID accepted submits collide on `records/relay-<pid>.json` and the second overwrites the first in an append-only/immutable store. Masked by every test presetting RelayID. Fix: server-assigned collision-free id (monotonic under `store.mu`, or random), independent of PID and client input. (lens: idiom #1, crash #3, security #8 — 3-lens convergence)

M6 — fault path classifies authority from PAYLOAD role. `internal/engine/loop.go:85-88` (verified E1): `faultOutcome` unmarshals `cand` from raw `cmd.Payload` and sets `meta.Role = cand.Envelope.Role` before `AuthorityBearing`, letting payload bytes steer the held-vs-rejected trust branch — contradicts "payload ignored byte-for-byte." Mitigated (Name from `cmd.Seat`; pessimistic direction) but wrong in principle. Fix: classify from the channel-stamped `SeatMeta`, never payload Role. (lens: security #5)

M7 — grant render/validate branches only on `IsOperator`, excluding orchestrator roles. `internal/fieldspec/fieldspec.go:70-76,102-104` gates `grant` on `seat.IsOperator` only; `ROLE=orchestrator-planner/orchestrator-reviewer` is never consulted, so an orchestrator seat gets no `grant` field and is rejected in-loop if it supplies one — but D-5 / the ratified narrowing say "operator/orchestrator forms only," and the orchestrator is the seat that issues delegated dispatch. A2 tests only operator-vs-pair, masking it. Fix: gate `grant` on operator ∨ orchestrator roles. (lens: correctness F-6)

### OPTIONAL (Implementer discretion unless folding anyway)
- O1 errcheck: 8 unchecked `Close()` fail `golangci-lint` (`internal/store/store.go:120,141`, `internal/store/projections.go:83`, `internal/fsio/fsio.go:65`, `capability_test.go` ×4). The durability-critical record Close IS already checked (fsio.go:36); these are lint-hygiene — but clear them so the lint gate is green. (idiom #3)
- O2 `appendUnique` (`internal/store/projections.go:68-85`) dedups via `strings.Contains` over a full re-read — O(n²) + substring false-dedup; line-anchored membership before mailbox projections land. (idiom #4, crash #7)
- O3 credential compare uses `==` in `internal/seat/binding.go:82-91` (TCB seat-trust) — use `crypto/subtle.ConstantTimeCompare`. (idiom #5)
- O4 `binding/seats.json` written 0644 world-readable (broader than D5 same-uid residual) — write 0600 / dir 0700. (security #6)
- O5 dead branch `internal/gate/derived.go:97-99` (assigns "" to "") — delete. (idiom #6)
- O6 `Push` (`internal/channel/server.go:82-87`) returns on first client write error, skipping remaining nudges — accumulate + continue. (idiom #8)
- O7 `Loop.Timeout` (`internal/engine/loop.go:42`) defined, never read; reply send `job.ReplyCh <- ...` (loop.go:52) unguarded — implement the per-job timeout + `select`-on-ctx send, or delete the dead field. Latent deadlock once real callers wire in (relates to B7 assembly). (idiom #2, crash #5)

### SCOPE NOTE (not a blocker; flag to orchestrator)
`internal/engine/submit.go` is a new source file NOT in the plan's file-structure block (which named only `internal/engine/loop.go` for that package) — a benign split of the submit pipeline within the in-fence `internal/engine/` package. Recorded for the record; not a fence breach (the package is in-fence), but the plan's file list should absorb it at the next revision. (lens: correctness scope-fence)

### Verdict + next
The byte-shape half of the contract is solid and should not be re-touched. The durability/atomicity/identity/exactly-once half — the reason this slice exists — is not yet real: B1-B7 must fold, then M1-M7. Because the fold implements missing behavior (B4 recovery steps, B7 assembly, M1-M3) and rewrites the fixtures that prove it, this is a genuine REVIEW-FOLD, not a nit pass. After fold-in I will run a targeted re-check on the blocker set + any file whose blast radius changed (not the full panel unless B7 assembly shifts the design surface). No exit-gate SITREP to the orchestrator and no merge path until the blocker set is green under a crash matrix that drives real mutations.

ACTIONS_GIT_REF: no edits this relay (fold-in-only authority begins with the Implementer's fixes); panel ran read-only; my spot-verifications were reads + `git log`/`grep`/`sed`, no tracked-file change; tree clean at main@9c1839e
FINAL_GIT_STATUS_SHORT: none — clean tree (`git status --short` empty at 20260703-190326; only subsequent writes are this relay + its INDEX row, both gitignored)
Next requested action: s1-core.implementer folds B1-B7 then M1-M7 (O1-O7 at discretion, O1 recommended for the lint gate), reissues an IMPL report; I run the targeted re-check. The scope note goes to s1.orchestrator-planner for the plan file-list absorption. This relay grants no merge authority and issues no dispatch token.
