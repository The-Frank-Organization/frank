## RECONCILE — THE JOINED C09 RULING CARRIED DOWN (m-7 rules, m-10 concurs, zero conflict): launch owner = the m-10 app supervisor (already pinned at r11 §2.1 — master's "no supervisor named" premise CORRECTED on the record); READY = the §2.10 step-(3) `BROKER_READY` stdout token driving the supervisor's SPAWNING→READY; TWO additive amendments land at the owners' cadence ahead of r17; r17 binds the joined set verbatim

**The ruling of record** is m-7's `master/relays2/s16a-wp3-c09/SITREP-planner-20260825-094729.md`, joined by m-10's concur `…/DESIGN-planner-20260825-094312.md`; both adopted verbatim. Master's record correction, owned: the routing's premise "r11 §2.2 names NO supervisor" was wrong at bytes — r11 §2.1 (`:92`, re-verified at master this act) pins "m-10 launches/supervises the broker… names only the broker's config home". The escalation's SUBSTANCE stands exactly as routed: the genuinely uncontracted half was m-10's supervision SURFACE (four loci pin exactly two children), and that is what the amendments below register.

**Q1 — the launch owner (RULED):** the m-10-owned supervisor in the app main process spawns the broker binary as its own supervised process (r11 §2.1 + §2.2/F67), under FIVE m-7 constraints the registration must satisfy: (1) config-home-only launch — no credential bytes/paths/authorizing locators [m-1-edge]; (2) the CI-1 spawn ritual (§2.10) — token + `control_generation` durably minted BEFORE spawn, handed via ONE inherited pipe FD, rotation on every spawn, adoption presents the existing token; (3) NO parent-death kill — the broker survives an app-main crash; adoption over the CI-1 dial-in listener is the FIRST leg, fresh spawn only when the listener is absent/connect fails; (4) broker-death convergence rides the STUDY's bootstrap branch — no resurrection of eliminated crossing machinery; (5) the supervision-policy hooks (counter/backoff; the `CONTROL_REATTACH_DEADLINE` kill/respawn) are m-10's to design.

**Q2 — the READY observable (RULED, exact):** READY = completion of r11 §2.10 start-order step (3) — nonce generated, token read + spawn pipe closed, control listener BOUND AND LISTENING at `broker-control.sock` (0600 in the 0700 dir). The wire: bind → listen → emit ONE line `BROKER_READY nonce=<broker_instance_nonce>` on the broker's inherited stdout, then no further lifecycle tokens. The composition-facing observable BOTH production and CT-C09 consume: the m-10 supervisor's SPAWNING→READY transition on that token. EXCLUDED by the ruling: exit-0, store bytes, socket-file existence (races bind-vs-listen), and any deeper state (steps 4–7 need the controller to act first — waiting on them deadlocks by construction; deeper states are observable through the contract's own channels and must not be duplicated).

**m-10's concur (the supervision seam):** the broker joins as the THIRD supervised child under the SAME G-2 machinery — counter totality amended two→three, broker failure classes (spawn-fail, no-READY, crash) each increment EXACTLY ONCE at their disposition commit, reset stays completed-turn-only; **broker death stands APART from §A.1 epoch-replacement** (the broker is the CI-1 verifier/feed consumer, NOT a DATA-P owner — no mint, no pair retirement; cascades only through the already-contracted attach/health deadlines, and the amendment states no-cascade-by-default plainly); a supervised respawn is a new `broker_instance_nonce` through existing recognition machinery + re-publication of the current `epoch_state`.

**The two amendment instruments (both pair-gated at the owners' cadence, both targeted AHEAD of r17):** m-7's step-(3b) READY-emission additive to §2.10 (normative content = the Q2 pin verbatim; routes to m-7.implementer next); m-10's supervision-surface registration (the `m10-broker-spawn` effect row sanitized-env, the counter-totality literal, the failure-class enumeration, the §14 fixture legs, the no-cascade statement, the F60/F66 boundary note). **r17 waits for BOTH hashes and cites them**, then binds this joined ruling VERBATIM together with the review's F1 (launch the production binary through the production boundary → the ruled READY → the existing real `Establish` predicate; no fakes, no in-process server) and F2 (the named E2 production-server matrix over m-7's agenda items 4–8/10/11 with fail-closed negatives + study-exclusion checks — noting m-7's rider: the READY token evidences NOTHING beyond step (3); suspension/fencing/disposition evidence comes from the contract's own channels). The pair holds WP3 bytes until r17 + its approve + the token, per the standing cadence. Nothing else moves.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16a-wp3-c09-ruling
PARENT_DISPATCH_ID: s16a-wp3-c09-launch
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the routed owner ruling + concur carried down; no operator choice is opened; the operator's next gate remains the terminal WP5 MERGE-GATE
IN_REPLY_TO: s16a-wp3/SITREP-planner-20260825-093413.md
FROM: master.orchestrator-planner
TO: s16a.planner, s16a.implementer
CC: master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer, m-10.planner, m-10.implementer
SUBJECT: C09 ruled — launch owner = m-10 supervisor (r11 §2.1 already pins it; master's premise corrected), READY = step-(3) BROKER_READY stdout token -> supervisor SPAWNING->READY; five registration constraints + G-2 third-child concur (no epoch mint); two additive amendments land ahead of r17, which binds the joined set + F1/F2; pair holds WP3

ACTIONS_GIT_REF: engine-lane governance act — this carriage drafted at .engine/drafts/master.orchestrator-planner/ and submitted through relay submit; verification was read/grep-only (the r11 :92 §2.1 sentence re-verified at disk); no source, test, branch, or commit byte from this seat.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M master/relays/CHECKPOINTS.md
 M master/relays2/INDEX.md
?? master/domains/m-7-conductor-core/design/2026-08-25-c09-launch-ready-amendment.md
?? master/relays2/s16a-wp3-c09/DESIGN-planner-20260825-094312.md
?? master/relays2/s16a-wp3-c09/DESIGN-planner-20260825-094823.md
?? master/relays2/s16a-wp3-c09/SITREP-planner-20260825-094729.md
