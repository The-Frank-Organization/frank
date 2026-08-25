# PL-s13-build-plan — the s13 straight-through build plan (m-10 supervisor MODULE)

PLAN_DOC_ID: PL-s13-build-plan
STATUS: r5 — binds master's R6 successor ruling (`…-20260821-033642.md`): the ONE superseded byte — the go
  directive spelling `go 1.25` → the executable **`go 1.25.0`** (E-1 row updated; versions, floor, and every
  other ruled byte stand; the R6 §1(iv) divergent-byte tripwire carried). Every other r4 byte preserved. Also
  folds the must-revise r3 review `s13-build-plan/PLAN-REVIEW-implementer-20260821-031316.md`
  (S13-PR-R3-F1: the self-referential dispatch-carried manifest REPLACED — T0 now waits for a master-banked git
  checkpoint committed AFTER the approving PLAN-REVIEW + IMPL dispatch exist and imports both trees from that
  exact commit, a finite external attestor; F2: T0's stale r2 plan pin corrected — the imported plan artifact
  verifies against the PLAN_LOCK_ID digest carried in the successor PLAN relay the approving review parents to,
  an attestor outside the artifact). Also folds the r2 review `s13-build-plan/PLAN-REVIEW-implementer-20260821-024813.md`
  (S13-PR-R2-F1: the relay tier restored to the commissioned PRODUCTION-RISK — carried by the r3 PLAN relay and
  inherited by the future PLAN-REVIEW/IMPL dispatch; F2: the boundary contract made runtime-semantic below, the
  path fence retained as G-A scope evidence; F3: T0's byte provenance closed — the IMPL dispatch carries a
  per-file SHA-256 manifest of both trees, and T0 verifies every imported file against it, so untracked
  authority-chain bytes are attested independently of banking cadence). Also folds the r1 review
  `s13-build-plan/PLAN-REVIEW-implementer-20260821-021351.md`
  F2–F6 (T0 deterministic two-tree import · T15 hand-relay close per the 020247 suspension, no pair store act ·
  FX-M10-C1 owned at T11 + the surface no-mutating-verbs negative at T13 + the reduced-tag battery command named
  at T3/T11/T15 · fake cores owned at T2/T9 · the boundary contract + design-§12 out-of-scope carry); F1 (the
  dispatch-predicate/trigger disposition) is ESCALATED to master (`s13-build-plan/SITREP-planner-20260821-021917.md`)
  and its returned ruling binds in the r2 PLAN relay, which files ONLY after that return. LOCKS the approved design
DESIGN_LOCK_ID: DS-s13-m10-module @ SHA-256 3e74c4c1fe76dbaf07f3ca9547ffeb88a958cf23d0f6180afa9d49cd71a81361
  (r3; approved `s13-build-design/DESIGN-REVIEW-implementer-20260821-004530.md`, verdict approve, F1–F4 all closed;
  path `designs/DS-s13-m10-module-20260820.md` in this sprint tree)
BASE: main@b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9 (LAUNCH_BASE, charter-pinned)
BRANCH: s13-m10-module · TARGET: main (operator-only MERGE-GATE, serialized s13 → s14 → s15)

## Global constraints (bind every task)

- G-A **Fence:** every branch byte lives in `frank/cmd/frank-app/**` · `frank/internal/appctl/**` ·
  `frank/internal/appipc/**` · `frank/.relays/s13/**` · `frank/docs/sprints/active/2026-08-20-s13-m10-module/**`
  (the R1-granted addition). `frank/go.mod`/`go.sum` move ONLY under master's E-1 ruling
  (exact bytes per the ruling). Anything else = escalation by construction.
- G-B **Straight-through cadence:** task-by-task, NO per-task review relays; the implementer builds ALL source/test
  bytes and commits (charter corrigendum 1); the planner's ONE adversarial end-review at completion.
- G-C **Green-at-every-commit:** `go build ./...` + `go test ./...` (incl `test/invariants` — the ten INV-CATALOG
  laws) + `go vet ./...` green at every commit; RED-first for every negative (write the failing leg, watch it fail,
  then the mechanism).
- G-D **Frozen-spec conformance:** every wire/store/order semantic realized from the composed spec (design §0);
  byte-exact tokens everywhere the docs close a set (`{accepted, rejected, held}` untouched conductor-side; the
  closed reject/terminal/void/stop_reason domains verbatim). A needed deviation = escalation, never a local fix.
- G-E **Worktree discipline:** all branch work in a dedicated git worktree (`harness-s13-m10-module`); the shared
  checkout at `~/Programming/harness` is NEVER switched off main by this pair. Lane relay files + sprint docs are
  authored in the shared checkout (untracked, in-fence) and committed on the branch from the worktree.
- G-F **Commit discipline:** path-scoped commits, fence paths only; no push until master directs; the MERGE claim
  only after the operator's MERGE-GATE per the charter's MG-F1 rule.
- G-G **No secrets, no payload:** no credential bytes, no provider payload, no `workspace_root_path`/
  `session_log_path` on any conductor-bound surface (m-1 carrier negatives in the battery).

## Escalations (planner files; build routes around them)

- **E-1 — RULED (R2 Route A, the directive spelling as RE-RULED by R6
  `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-033642.md`):** T4 opens with the
  implementer's exact ruled edits on the branch — `go.mod`: directive **`go 1.25.0`** (the executable spelling;
  identical on every branch touching the directive) + `require modernc.org/sqlite v1.57.0`; `go.sum` exactly as
  `go mod tidy` generates; **if tidy emits a `toolchain` line or ANY directive byte other than `go 1.25.0`, that
  is a NEW divergent byte = the same arbitration stop, fire-before-commit** (R6 §1(iv)); NOTHING else in go.mod
  moves; the README floor line is master's (never this pair's). Any NFC need: `golang.org/x/text@v0.41.0`
  (R4 one-version seam, unchanged by R6).
- **E-0 — RULED (R1):** the sprint tree `frank/docs/sprints/active/2026-08-20-s13-m10-module/**` is s13 fence;
  the artifacts live there; T0's docs-bank commit includes it + the s13 relay lane.
- **E-2/E-3 (conditional):** R-M10-BURST-CANDS bite / frozen-base defect → up the ladder, per design §14.

## Tasks (dependency-ordered; ⊕ = parallelizable with its predecessor once its own deps are met)

- **T0 — Worktree + branch + docs bank (deterministic import; F2).** Create the worktree on `s13-m10-module`
  from LAUNCH_BASE `b7f406b2…` (record the exact worktree base SHA in the IMPL report). TWO in-fence populations
  bank in the first commit, imported from ONE master-banked git checkpoint (below):
  (1) `frank/.relays/s13/**` (relays + INDEX, the complete authority chain through the IMPL dispatch); (2)
  `frank/docs/sprints/active/2026-08-20-s13-m10-module/**` (the R1-granted tree: the approved design + the
  CURRENT plan revision — verify the design digest `3e74c4c1…` and verify the plan artifact against the exact
  `PLAN_LOCK_ID` digest carried in the PLAN relay that the approving PLAN-REVIEW parents to; both attestors live
  OUTSIDE the artifacts they attest). Import = byte-copy of exactly those two path trees. **Byte provenance (finite + non-self-referential; F3 as corrected by
  R3-F1):** T0 WAITS for a master-banked git checkpoint COMMITTED AFTER both the approving PLAN-REVIEW and the
  IMPL dispatch exist, whose tree contains the complete s13 authority chain (design/plan relays, every review,
  the dispatch itself) AND both exact in-fence populations; T0 records that exact commit SHA and imports BOTH
  trees FROM THAT COMMIT'S TREE (e.g. `git checkout <sha> -- <the two paths>` into the worktree), never from the
  live working tree — a commit SHA attests every byte it contains, no attestation artifact rides inside the
  imported population, and no self-reference exists. A missing authority-chain file or a hash mismatch at the
  two named verifications = STOP, escalate — never import. **No source/code work begins while T0 waits** (master's
  banking cadence is frequent; the wait is bounded by one checkpoint). A pre-commit census additionally asserts the staged path set == those two
  trees and NOTHING else (no foreign checkpoint path rides in). Files: the two trees only.
- **T1 — appipc core: `jcs.go`, `counter.go`, `frame.go`, `envelope.go`.** JCS encoder (sorted keys, `,`/`:`,
  UTF-8, no floats on trust paths) + the counter grammar (wire ^(0|[1-9][0-9]*)$ < 2^64; store 20-digit pad codec)
  + uint32_be framing with FRAME_MAX + envelope schema/family registry (closed vs additive; reply-class requires
  `re`). Battery: JCS vectors incl reproducing the limits-doc §9 builder integers (`|S_cont|=770`, step-by-1 lemma);
  counter RED legs (leading zero, sign, 2^64, non-digit); oversized-frame fault; unknown-type/unknown-field/no-`re`.
- **T2 ⊕ — appipc message families + the fake frame cores (F5).** Every CTRL-W/CTRL-C/broker message struct per
  design §1's inventory, with the closed-set validators (attach tokens, reject reasons, outcome domain,
  disposition domains) byte-exact; PLUS the testutil foundations `fakeworker` + `fakeconnector` (design §10 —
  scriptable peers speaking these exact frames over real socketpairs; consumed by T5+ batteries). Battery:
  per-family encode/decode round-trips + closed-set RED legs + fake round-trip smoke.
- **T3 ⊕ — limits tables + assertions.** `limits.go` (production constants) + `limits_reduced.go`
  (`//go:build frank_test_reduced_limits`) + the negative-array-size assertions (P₁ positive in production, FLIPPED
  under the tag; P₂ positive in both) + the reduced-table content-addressed test artifact + the witness builder.
  Battery: exact-fit W commits-shape / one-byte-over W⁺ classification at the sizing function level (full runtime
  legs land with T10/T11), run under **`go test -tags frank_test_reduced_limits ./internal/appipc/...`** (the
  reduced-tag command, named here at its first owner); FX-M10-ACK-BOUND's build-level negative via the design-§9
  executable harness
  (`limitscompileneg/` fixture + `TestLimitsCompileMatrix` shelling the three toolchain commands; the
  complementary `!frank_test_reduced_limits`/`frank_test_reduced_limits` constraints close the selector).
- **T4 — store (GATES ON E-1).** Driver wiring; open/genesis (`user_version=1`, one transaction);
  higher-refuse/lower-migrate/corrupt-refuse; full schema DDL per design §2 (all tables, CHECKs, UNIQUEs, padded
  counters); floor fixtures (WAL + synchronous=FULL PRAGMA read-back; private-file perms; forward-only migration;
  transaction semantics). Battery: boot family §14 (genesis/refuse/migrate/corrupt).
- **T5 — applier.** The single-writer event loop; typed events; commit-then-emit; snapshot reads; timers-as-events.
  Battery: durable-then-visible probe (no observable pre-commit state); single-writer import test.
- **T6 — manifest + serve gate.** The 8-name constant; workspace_root recipe (the ordered refusal table);
  C.1 construction + freeze-at-admission transaction; C.2 digest; C.3 serve gate incl F63 staging (null-until-lock
  semantics, fail-closed). Battery: FX-M10-WR-R (all rows incl crafted-cwd, symlink-to-root, both-defects-one-token);
  FX-M10-WR; manifest §14 legs (freeze immutability, frozen-row-only serve, F63 mismatch).
- **T7 — F59 host.** D.1 ticket + §5 effect descriptor (R/∅ table; operand source split; replay-identity extension);
  D.2 (0)–(7); D.3 consume; D.4 record + expiry. Battery: the COMPLETE §D.2/§D.3/§D.4 fixture lists verbatim
  (design §11) + FX-M10-C-1.
- **T8 — supervision + epoch.** Spawn (sanitized env, socketpairs, CoE, runtime dir), worker/connector machines,
  health, terminate ladder, THE retirement transaction (as amended — no ledger), G-2 counter+backoff, epoch mint/
  distribution/rejection (source-specific authority), co-restart pairing, EOF handling. Battery: lifecycle §14
  family COMPLETE (incl the 4-assertion pre-ready connector-failure fixture, 9th-vs-10th, backoff persistence,
  cancel-during-backoff, loud surfaces).
- **T9 — broker client.** OPENS with the `fakebroker` testutil (design §10 — the CI-1 listener fake: control
  handshake, the state_proposal disposition table, event emission incl `boundary_cut`/`epoch_installed`,
  adoption/loss; F5). Then CI-1 dial-in (fcntl lock → durable `control_generation` advance → connect; spawn-pipe
  token), the state_proposal receiver fold (total), the two-form assign gate, broker-event recording (amended
  family), CI-4 broker spawn discipline. Battery: FX-M10-C0 complete; FX-M10-C2 (broker survives app-main SIGKILL);
  the FX-TB-5′/17′/20/21 reciprocals; recovery re-proposal legs.
- **T10 — scheduler.** Admission transaction (gates ∧-ed, lease + wake flip + `turn_disclosures` snapshot in ONE
  commit), sizing gates (initial refusal / continuation terminal), `turn_open` emission (snapshot-derived;
  genesis `run_disposition` re-derivation), `attempt_open`/`attempt_open_ok` (live-derived disclosure; row-commit-
  before-ack), D-5 tables, the B.1 cancellation table, wake chain (§E). Battery: FX-M10-D4 complete; scheduler §14
  legs; §B.1 attempt-observation + cancellation cut lists; D-5 fixture list; wake at-most-once crash cuts.
- **T11 — settlement + carriage.** D2 producer; S-1 receipts (REDUCED tuple); D3 continuation (+ the PAIR,
  S-2 frames); the cap terminal; B/E carriage columns + `m10_row_state`. Battery: FX-M10-D2-1/1b/2/3; FX-M10-R
  (a)–(d)+stale+unknown; FX-M10-D3-1/2/3/4; FX-M10-CAP (i)–(iv); **FX-M10-C1** (a cut relay call yields exactly ONE
  `uncertain` manifest entry from the canonical row; a `boundary_cut` event with NO matching canonical row yields
  NO entry — telemetry-never-input asserted); FX-BE-1..8; the runtime overflow legs (exact-fit commits ·
  one-byte-over ⇒ the terminal triple · encoded-reserve) under the reduced tag — **run by the named command
  `go test -tags frank_test_reduced_limits ./internal/appipc/... ./internal/appctl/...` (the reduced-tag battery;
  ordinary `go test ./...` does not select it)**.
- **T12 ⊕ — E0 persist.** m-3-schema validation + `pending_app_events`. Battery: schema-conformance + no-silent-drop.
- **T13 — cmd/frank-app.** main; boot=recovery-always; the §10 terminal surface; loud failures. Battery: surface
  legs (payload-free views; non-zero exits; persistent alert rendering; **the executable no-mutating-verbs
  negative — every verb beyond `run start`/`run stop`/`run cancel` is absent/refused, and no surface path can
  clear a parked row or forge a disposition, asserted against the command registry AND at runtime**).
- **T14 — recovery + negatives sweep.** Matrix (a)–(d) end-to-end with re-proposal; §G negatives (not-a-seat import
  test; no-credential-bytes; no-provider-payload type boundary; stale-fenced; crash-honest; wake-no-dup); the m-1
  carrier negatives.
- **T15 — completion.** Full battery — BOTH commands named: ordinary `go build ./... && go test ./... && go vet
  ./...` AND the reduced-tag `go test -tags frank_test_reduced_limits ./internal/appipc/... ./internal/appctl/...`
  + the T3 compile matrix — INV-CATALOG green; the implementer's IMPL report (literal
  `git status --short` + branch SHA + battery outputs); THEN the planner's ONE adversarial end-review; fold
  blockers (REVIEW-FOLD); merge-ready SITREP to master (undischarged queue rows enumerated per F.7.3), FILE-FIRST
  hand-relay transport per the operative suspension (`master/relays/step3-t4-preflight/RECONCILE-orchestrator-
  planner-20260821-020247.md`: the store is DORMANT — no pair store submit/poll/export act exists in this plan;
  the bounded store-records export through the suspension point is MASTER's, cited not performed); restack onto
  then-current main + battery RERUN + re-review of changed reviewed bytes BEFORE the operator MERGE-GATE
  (charter contract).

## Boundary contract (F6)

- **Runtime writes (the canonical objects/events this module authors):** the private m-10 SQLite state — every
  §F table row (`runs` incl. manifest bytes + digest, `workers`, `turns` incl. `resume_snapshot`/dispositions,
  `provider_attempts` incl. the B/E digest columns, `tool_calls`, `tool_authorizations` incl. the effect
  descriptor, `epochs`/`leases`/`state_seq`, `pending_app_events`, `cancellations`, `wake_schedule`,
  `broker_control`, `broker_events`, `content_ready_receipts`, `turn_disclosures`) and its committed snapshots ·
  the emitted CTRL-W frame family (`assign`, `turn_open`, `attempt_open_ok`/`attempt_open_reject`,
  `ticket_granted`/`authorize_reject`/typed rejects, `consume_ok`, `turn_receipt`/`turn_reject`, receipts/
  `disposition_conflict`, `ping`, `shutdown`, `admission_refused`) · the emitted CTRL-C family
  (`connector_assign`, `epoch_update`, `ping`, `shutdown`) · the broker-channel client sends (`state_proposal`,
  event acks) · the settlement manifest + `m10_row_state` facts · terminal-surface renderings (state-only).
- **Runtime reads (validated against the frozen schemas):** operator commands + run inputs (`run start` config,
  `workspace_root`, `credential_ref` opaque) · the received CTRL-W family (`hello`, `attach_result`,
  `attempt_open`, `attempt_stream_end`, `app_event`, `wake_forward`, `authorize_tool_call`, `consume_ticket`,
  `record_tool_outcome`, `turn_terminal`, `turn_cancel_ack`, `content_ready`, `report_resume_disposition`) ·
  the received CTRL-C family (`hello`, `connector_ready`, `attempt_result`) · broker results/events
  (`state_proposal_result`, the amended closed event family incl. `boundary_cut`/`epoch_installed`).
- **Writer→reader ownership:** m-10 = sole writer of its store + sole author of the frames above; the s15 worker
  writes the CTRL-W requests m-10 reads and reads m-10's CTRL-W commands/replies; the s14 connector writes
  `connector_ready`/`attempt_result` and reads `connector_assign`/`epoch_update`; the m-7 broker writes
  proposal results/events and reads `state_proposal`/acks + the `epoch_state` feed; the terminal surface reads
  committed snapshots only; the conductor's E0 consumer reads `pending_app_events` ONLY via the s15 worker's
  seat carriage (m-10 never submits); s16 consumes the composed app. A frame with a missing reader or writer at
  build time = escalation through master (serialized seam arbitration), never a local invention.
- **Fence (scope evidence; the G-A path law):** writes land only in `frank/cmd/frank-app/**` ·
  `frank/internal/appctl/**` · `frank/internal/appipc/**` · `frank/.relays/s13/**` ·
  `frank/docs/sprints/active/2026-08-20-s13-m10-module/**`; `frank/go.mod`/`go.sum` exactly the R2-ruled bytes;
  reads: the frozen basis set, the consumed contracts, the conductor tree + `frank/internal/seatclient/**`
  consume-only.
- **Target entity:** the m-10 supervisor MODULE (app control plane) + the appipc shared types package.
- **Downstream consumers:** s14 (CTRL-C/connector frames) and s15 (CTRL-W/worker frames) consume appipc — their
  change-requests arbitrate through master (never direct edits, either direction); s16 consumes the composed app.
- **Contract:** the composed frozen spec (design §0) + the interface lock `ab3dfa86…`; byte-exact closed token sets.
- **Proof (E2 — executable, both sides of every seam):** the design-§11 battery map (RED-first) establishes
  writer AND reader behavior per family — the §D.2/§D.3/§D.4 fixture sets (F59 both halves via fakeworker), the
  §B.1/§B.2 attempt/attach/D-5/cancellation tables, the receipt/disposition tables (S-1/S-2), the connector
  bootstrap legs (fakeconnector), the broker fold FX-M10-C0 + shared reciprocals (fakebroker), the store
  floor/boot family, the surface legs incl. the no-mutating-verbs negative — plus the closed-token conformance
  sweep, INV-CATALOG green, the two named battery commands, and the compile matrix. A leg that cannot be
  established at E2 = escalation, never a silent downgrade.
- **No-consumer-action:** nothing in this plan asks s14/s15/m-x to change bytes; any needed counterparty change is
  an escalation through master (serialized seam arbitration).

## Out of scope (carries design §12 verbatim; in-fence code MUST NOT build toward these)

R-M10-PROCESS-SPLIT (module stays in-process; as-if seams only) · R-M10-SCHEDULER (one run/one worker/one turn) ·
R-M10-WAKE-ALO (at-most-once/advisory only) · R-M10-FORENSIC-JOURNAL (chokepoint built, journal absent) · the m-5
ceiling POLICY (the seam is built EMPTY — enforcement socket only) · any conductor-side change incl. the version
handshake (design §13 = a note only) · the frozen exit oracle + RLBS-1 + everything under `master/exit-fixtures/`
(UNTOUCHABLE) · H-12 (nothing relaxes the external/untrusted/multi-tenant block) · the README Go-floor line
(master's edit) · any store submit/poll/export during the courier suspension (master's 020247).

## Exit criteria (merge-ready)

1. Every battery row in design §11 exists, ran RED-first where negative, and is green at the final SHA.
2. `go build/test/vet ./...` green incl `test/invariants` (the ten laws untouched and green).
3. Zero out-of-fence bytes in `git diff LAUNCH_BASE..HEAD --name-only` (except the E-1-ruled go.mod/go.sum bytes).
4. The R-M10-SQLITE-DRIVER floor fixtures green; the discharge citation ready for the merge brief.
5. The composed-spec conformance sweep: every closed token set grep-verified byte-exact against the frozen docs.
6. The end-review verdict filed; every finding folded or escalated; the merge brief enumerates undischarged queue
   rows; restack + rerun + re-review complete at the gate.
