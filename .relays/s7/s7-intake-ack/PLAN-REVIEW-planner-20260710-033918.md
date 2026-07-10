## PLAN-REVIEW — pair-Planner CO-SIGN of the s7 intake (B10 checkpoint): APPROVE with TWO binding redlines folded — row 4 must not assert `bound` as record-derived (B-1 locks it as runtime state, deliberately NOT a record), and the row-6 census is SIX families, not five (the `seat_mint` accept-reply with its two NAMED carve-outs). Code may start with both incorporated.

ROLE: Planner
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s7-intake-ack
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: s7
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this IS the pair co-sign checkpoint the dispatch names (step 3); external gates (m-1/m-2/m-4 fidelity · VP integration · operator merge) unchanged
GRILL_REQUIRED: no
DESIGN_LOCK_ID: c4-design-m-7-conductor-core
PLAN_REVIEW_VERDICT: approve
IN_REPLY_TO: .relays/s7/s7-intake-ack/SITREP-implementer-20260710-033425.md
FROM: m-7.planner
TO: m-7.implementer
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: co-sign granted — the mechanism map is contract-faithful on 8 of 10 rows as written; the two redlines are ENFORCEMENT of claim grains the contract already binds verbatim (B-1 r5 for row 4; the R2 fifth-family fold for row 6), not contract changes — fold them and start; dispute either ⇒ no code, route to master per the dispatch

**CO-SIGN: APPROVE, with the two redlines below folded as binding realization conditions.** Both enforce locked texts the effective plan already binds (r2: "row texts 4 and 5 carry the claim-grain bounds verbatim"; r3 fold 4 + VP finding 4 for the census), so no master trip is required — this is the guide half doing exactly what B10 put it in the pair to do. If you read either as a contract-level change, do NOT fold silently and do NOT start: route it to master per the dispatch's blocked-or-ambiguous line.

### Redline 1 — row 4 (`TestLawDerivedOnlyActivation`): the `bound` leg is NOT record-derivable, and the law test must not pretend it is

Your map says "assert minted -> bound-now -> active from canonical records." Two-thirds of that is the law; one-third would violate the design being tested. B-1 r5 (`master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:95`) locks **`bound` = runtime channel state, deliberately NOT a record** — "recovery derives `minted`/`active` from records (store-is-truth) and opens with `bound = ∅`"; recording binds-as-records was considered and **rejected**. A law fixture that derives `bound_now` from canonical records would enshrine the rejected alternative in the invariant battery. Correct shape:
- `minted` derives from the committed `seat_mint` pivot; `active` derives ONLY from the first accepted governed submit per mint-generation, **in commit order** (m-1's order rule, `:96`) — these two are the "derived solely from committed records" claim;
- `bound_now` is produced by live runtime state in the walk (open a channel, observe the roster bit flip; restart ⇒ `bound = ∅` while `activation_state` survives — the active-but-disconnected state at `:111` is the sharp assertion worth one extra leg);
- the no-persisted-activation-marker assertion stays exactly as you have it, and gains its mirror: **no record, registry field, or marker encodes `bound` either**.

### Redline 2 — row 6 (`TestLawPathHygiene`): the census is SIX families — the `seat_mint` accept-reply, with its two NAMED carve-outs, is a census ROW

Your five families are the r3 list; the R2 fold (`master/relays/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md`, folded into the s6 design) established the **`seat_mint` accept-reply as its own seat-delivered I-PH family**, carrying the only two sanctioned carve-outs: the fresh credential + the new seat's **endpoint**, operator-channel-only. This is load-bearing for your scanner, not bookkeeping: your scan corpus includes **socket-path families, and the mint reply legitimately carries a socket endpoint** — without the named, scoped exemption the scanner either false-positives on sanctioned bytes or gets weakened to pass. Requirements: (i) the family is a named census row with the exemption stated (that family · operator channel · those two payload elements · nothing else); (ii) the planted-leak negative still bites INSIDE the mint reply — a planted canonical store/config/outbox path in the reply outside the two carve-outs must turn the check red; (iii) your AST sink census should discover the mint-reply sink mechanically — if it does not, that is a census-coverage finding, not a reason to hand-wave the family in.

### Notes (non-blocking, carried to the pair review)

- **Fixed-point discipline (rows 1/2):** the expected sets — `{accepted, rejected, held}` and `{submit, project, read}` — must be **literal bytes in the test file**, never read from the surface under test and compared to itself. Your row 1 has two distinct legs (the literal byte-exact comparison = the enum law; the seat-supplied fourth token = the system-owned fence — your typed class is right); keep their assertion messages distinct so a future edit cannot drop one silently.
- **Row 6 census mechanics:** the AST sink-pattern set is itself census content — name and count the patterns in the catalog artifact, so a future delivery mechanism that bypasses the pattern set is visible at review. Your synthetic-unregistered-sink negative satisfies VP finding 4's mechanical-red requirement; this note just makes the pattern set auditable.

### Confirmed as written (no change requested)

Supersession reading — matches mine exactly. File fence — exact per the dispatch (`test/invariants/**` · `docs/sprints/2026-07-10-s7-inv-catalog/**` · `.relays/s7/**`; nothing under `cmd/**`/`internal/**`). Row 5 — grain verbatim (sole *governed* path, D5 residual stated, typed `root-lock-held`). Row 8 — the child-process crashpoint walk is the sanctioned §4 idiom (crash-before absent / crash-after present, exactly one rename). Row 9 — all five r3 clauses including legitimate pending-zero and duplicate-content replay: this is my A-2 grain verbatim. Row 10 — the R-3 discipline realized correctly (no seat surface before Ready). Catalog artifact contents — correct, including the explicit no-§7-claim (s8 carry). Findings-not-fixes fence — correct.

**Start condition:** fold Redlines 1 and 2 into the mechanism map (a one-paragraph delta in your first build relay or IMPL report is sufficient — no second co-sign round) and begin. My adversarial pair review of the IMPL report will check against this co-sign + the pinned checklist in `master/relays/s7-dispatch/SITREP-planner-20260710-033057.md`.

Not authorized / not claimed: pair co-sign only — no code (I never write `frank/` code), no contract amendment (both redlines enforce already-bound texts), no fidelity verdict (m-1/m-2/m-4's), no integration or merge authority.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root .relays/s7 .relays/s7/s7-intake-ack/PLAN-REVIEW-planner-20260710-033918.md` — run before handoff (result in inline pointer)
- locked texts re-read for the redlines: B-1 `:95` (bound = runtime, NOT a record; rejected alternative recorded), `:96` (active order rule, commit order), `:111` (roster two-part shape, active-but-disconnected legibility); the R2 fold (five payload families + the two named carve-outs)
- the implementer's intake ACK read whole, incl. its verification block (base `1d3e92c`, battery 24 ok + 2 no-test-files, vet clean — matches my own probe at `20260710-033057`)

ACTIONS_GIT_REF: none — no git action by this relay (disk refs: this relay + one slice INDEX row, both under gitignored `.relays/s7/`; no code, no tracked-file, no doc/design edits)
FINAL_GIT_STATUS_SHORT: worktree `s7-inv-catalog` at `1d3e92c` — untracked `.relays/` + `docs/sprints/` scaffold only (per your ACK), plus this relay; no tracked-file change
Next requested action: operator carries this to the m-7.implementer session; the implementer folds the two redlines (one-paragraph delta, no re-round) and begins the build at `1d3e92c`; next pair touchpoint = the IMPL report to me for the adversarial pair review.
