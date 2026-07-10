## SITREP — m-1 fidelity packet for the s4 design (five items: config_change record_kind + provenance · credential-hash active-index · shim custody · per-recipient wake identity use · config/-as-derived store-shape posture)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-fidelity-m1
PARENT_DISPATCH_ID: s4-wire-design-complete
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s4-wire-design/SITREP-planner-20260705-042500.md
FROM: s4.orchestrator-planner
TO: m-1.implementer
CC: s4.orchestrator-reviewer, operator, m-1.planner
SUBJECT: FIDELITY REVIEW REQUEST (m-1) — review object = the s4 design's m-1-touching surfaces at frank main@2ef9437 (`docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md`); five specific items; verdict gates the pair's delegated dispatch (the S2/S3 pattern)

**Review object:** `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md`
at main@2ef9437 (r3 — pair-approved `s4-wire-design-r3-review-implementer`, grill-locked
`s4-grill-s4-wire`, m-7 guide answers folded from
`.relays/s4/s4-guide-q1/SITREP-planner-20260705-014633.md`). Context: the s4-dispatch +
the s3-scope-q1 ruling condition 4 (m-1 fidelity on the §7 `record_kind`). Your verdict is a
standing condition of the pair's delegated dispatch — no store/identity-touching dispatch
before your approve is on record in `.relays/s4/s4-fidelity-m1/` (the S2/S3 pattern; a
must-revise folds bounded).

**The five items:**

1. **`config_change` record_kind + provenance (design §6.1 — the ruling's named condition).**
   Token: `record_kind: config_change`; headers `{member, new_digest}`; body = the FULL new
   member content (compound record, grilled G-1); **authorship = operator-channel `submit`**
   (grilled G-2; m-7 guide Q3 confirmed (a), rejecting the `system`-stamped alternative as a
   second provenance path for an authority-bearing act on a live store; genesis stays the
   sole system-stamped exception). Registry seat-scopes the record_kind value to the operator
   seat (non-operator submit carrying it bounces, fixtured). Confirm token, field homes
   (headers vs body per your S2 F-M1-1 prescriptions), and the provenance reading.
2. **Second-connect active-channel index (design §4).** `serverConn` retains the SHA-256 of
   the presented credential post-`Resolve` (never the raw value); `Server.active
   map[credHash]*serverConn` under the existing mutex; typed path-free `auth:channel-active`
   reject; recovery = kernel-close removal. **NO binding-table shape change** (a needed change
   = hard stop + escalate). Rule source of record = the s4-dispatch [VP-W1] (locked m-1 text
   is silent on concurrent connects — the dispatch pre-constrains). Confirm the hash
   retention + index sit within your locked channel-identity reading (DI-2/I2) and that
   reject/proven-dead-recovery stays inside the locked contracts.
3. **Shim custody posture (design §2, §7; m-7 guide Q6 guardrail 2).** The shim holds the
   seat credential and IS the channel (your binding grain): `FRANK_CREDENTIAL` env-var
   default, 0600 credential-file secondary, never a bare CLI flag; D5 note verbatim on every
   custody surface; the no-in-band-rotation truth stated plainly ("a compromised credential
   means stopping the conductor and admin-time surgery" — your §13.3 carry, honestly
   deferred, not implied away). Confirm the custody wording meets your D4/D5 claim
   discipline.
4. **Per-recipient wake (design §5).** The channel layer keys wake events off the auth
   metadata's seat identity (your stamp source) — `PushTo(seat, …)` for exactly the committed
   record's recipients; the cross-seat metadata leak (a non-recipient seeing another seat's
   pending state) is FIXED with a negative fixture (S4-NG4). No new authority, no new verb;
   identity is consumed, not extended. Confirm.
5. **`config/` demoted to a derived projection (design §6.2 — store-shape posture).** After
   the first `config_change` record, the `<root>/config/` members become projection-grade
   artifacts re-materialized from the record chain (record-pivot-then-derived-bytes, the m-7
   compound-record idiom; phase-0 walks genesis→config_change digests in commit order,
   committed records only, re-materializes on mismatch, fail-closed-serving-reads on a
   persisting mismatch). Your S2 F-M1-3 placed pinned config under the store root with
   `store.Init` materializing source bytes — confirm the demotion + re-materialization
   posture is consistent with your store-shape ownership (fresh stores degenerate to today's
   exact behavior; `store.Init` untouched).

Per your S3 pattern: if the gated PLAN implements your conditions verbatim, state whether a
separate narrow re-review is needed or the verdict stands on the plan's verbatim carry (your
call, as in S3). Route-back triggers you name are binding on the pair.

ACTIONS_GIT_REF: wrote this packet + an INDEX row (relay substrate, git-untracked); no frank/ or spec-tree edits.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 2ef9437; battery 20-ok uncached + vet clean at this surface, my runs)
Next requested action: operator hand-relays to m-1.implementer; verdict lands in .relays/s4/s4-fidelity-m1/; the pair's delegated dispatch waits on it (PLAN drafting proceeds in parallel).
