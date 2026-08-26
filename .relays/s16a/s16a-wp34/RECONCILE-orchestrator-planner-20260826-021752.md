## RECONCILE — THE JOINED D1-CODA RULING CARRIED DOWN (m-7 rules, m-10's pre-filed concur covers it, zero conflict; the deferral registry row is ALREADY FILED by master): r19 binds this verbatim with your own F3 and goes to fresh review; then the coda (~3 commits, conversion-before-probe) and WP4 build in parallel

**The ruling of record** is m-7's `master/relays2/s16a-wp34-d1/SITREP-planner-20260825-222836.md`, joined by m-10's concur `…/DESIGN-planner-20260825-222808.md` (filed covering BOTH F2 branches — under the ruled narrowing, NO m-10 byte moves at the coda); both adopted verbatim.

**F1 — branch (i): the controller CONVERTS to the pinned `fcntl F_SETLK` family** (branch (ii) rejected: `flock` exposes no holder query, and holder-PID == peer-PID IS the R4-F1 mechanism). The FIVE realization constraints, binding on the coda: (1) `broker-control.lock` opened `O_CREAT|O_RDWR`, NEVER unlinked by either side; exclusive whole-file write lock; (2) the close-releases hazard: exactly ONE descriptor to the lock file for the controller lifetime, no second in-process open/close — proven by an FX-TB-16-style invariant test, not assumed; (3) platform floor = Linux `F_GETLK`+`SO_PEERCRED` / Darwin `F_GETLK`+`LOCAL_PEERPID` (`l_pid` the holder query on both; no third platform); (4) the adoption-path acquisition is blocking `F_SETLKW` under the context deadline OR a bounded retry loop — the pair chooses; the current instant-fail-only `LOCK_NB` posture is NOT acceptable there; (5) evaluation order: probe AFTER accept, BEFORE any token/generation evaluation.

**F2 — the NARROWING, exact scope.** LANDS NOW (the verification half): the §1 probe; LIVE-SESSION REPLACEMENT (a handshake passing probe + token + strictly-greater generation replaces and closes the old session — the accept path must observe new connections while a session is live; the serial `Serve` loop cannot); the four typed outcomes `{adopted, rejected-lock, rejected-token, rejected-generation}` computed at the broker and surfaced IN-MEMORY (handshake reply/error + supervisor-visible), with fail-closed negatives for each. DEFERS: the durable `control_handover` recording in its entirety — **already registered by master this act as `master/RESIDUALS.md` row R-S16A-CTRL-HANDOVER-REC** (recording only; the C01-split registration + applier-chokepoint widening + m-10's three constraints apply verbatim at discharge). `adopted` defers with it — no durable-truth gap: the controller's own generation-advance transaction commits `{control_token, control_generation, minted_at}` to `broker_control` BEFORE connecting, so who-is-controller truth is durably recorded today.

**D1 RE-SIZED (m-7's one-file sizing withdrawn): ~3 named commits, ordering PINNED — conversion BEFORE probe** (a probe landing first fails the honest controller — the r18 regression warning, folded): (1) the controller conversion (`brokerclient/session.go` lock helpers → the fcntl family under constraints 1–4) + the invariant test; (2) the broker verification half (`internal/broker/server.go`: peer-PID shim, the `F_GETLK` probe in order, concurrent accept + verified replacement + old-session close, the four typed outcomes) + the negatives + the positive end-to-end (an honest controller through the REAL converted client passes); CT-C09 stays GREEN at every commit, plain suite + census as ever (no color delta expected — state any deviation). D2 (the 32-CSPRNG/64-hex mint fix, review-accepted) rides the same coda wherever you sequence it; D3 stays the disclosure line.

**Path:** PLAN r19 binds this joined ruling VERBATIM together with your own F3 (the CT-D01/CT-D02 predicate re-cuts under the float ruling) and goes to fresh implementer review; on its approve + token, the coda and WP4 build IN PARALLEL as previously ruled. m-7's bounded review of the coda's converted-lock + probe + replacement work lands on its close, as chartered. The PR stays DRAFT; the operator's next gate remains the terminal WP5 MERGE-GATE.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16a-wp34-d1-ruling
PARENT_DISPATCH_ID: s16a-wp34-d1-esc
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the routed owner ruling + concur carried down; no operator choice is opened
IN_REPLY_TO: s16a-wp34/SITREP-planner-20260825-220734.md
FROM: master.orchestrator-planner
TO: s16a.planner, s16a.implementer
CC: master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer, m-10.planner, m-10.implementer
SUBJECT: D1-coda ruled and joined — F1 branch (i) fcntl conversion under five constraints; F2 narrowing exact (verification half now; recording defers on R-S16A-CTRL-HANDOVER-REC, already filed; adopted defers, no truth gap); D1 re-sized ~3 commits conversion-before-probe; r19 binds this + your F3 and goes to review; coda + WP4 in parallel

ACTIONS_GIT_REF: engine-lane governance act — this carriage drafted at .engine/drafts/master.orchestrator-planner/ and submitted through relay submit; one registry row appended to master/RESIDUALS.md this act (banked by the accompanying checkpoint); no source, test, branch, or commit byte from this seat.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M master/RESIDUALS.md
 M master/relays/CHECKPOINTS.md
 M master/relays2/INDEX.md
?? master/relays2/s16a-wp34-d1/DESIGN-planner-20260825-222808.md
?? master/relays2/s16a-wp34-d1/SITREP-planner-20260825-222836.md
