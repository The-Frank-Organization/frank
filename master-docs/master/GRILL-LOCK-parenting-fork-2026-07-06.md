# GRILL_LOCK — the s6 parenting/anchor fork (F11/F4): conductor-computed PARENT, fallback hints, m-4-gated Sharpening-D

GRILL_LOCK_ID: GRILL-LOCK-parenting-fork-2026-07-06
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relay read: the m-1 §A decision packet (`master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md` §A) · the m-1 pair grill (`s6-design/DESIGN-REVIEW-implementer-20260706-184120`, §A approve-to-forward + forward conditions) · `TRANSPORT-FINDINGS-2026-07-06.md` F4/F11 + root cause
- code/docs inspected: `lineage.go:346-371` + `main.go:379-392` (the moving-window candidate mechanics — `WokenOn` = the projection tail); the s6-design dispatch + [VP-W3]
- questions answered from codebase (not asked): **the empirical bad-pick audit** — archived dogfood store: **9 parent bounces, classified: 5 = real ACCEPTED relay-ids in that very store** (the window rejected valid references), **3 = system-fault/shape** (1 form-offered dead-edge token; 2 real ids from the adjacent file namespace), **0 = author error** (no typos, no hallucinated ids). Manual file trail: **913 relays, 892 parents resolve exactly (97.7%)**; the 21 "dangling" = two boot **family-name** references (`master-boot` ×16, `s1-boot` ×5), harmless, uncaught for four days, zero consequences. Also established: hint provability (read-history ∩ accepted-graph) is **monotone** — a valid reference can never become invalid; failures are always author-side shape/namespace, never timing.
- questions asked operator: 4, one at a time (branch · hint-fallback · precision cost · Sharpening-D process), with email's threading model (`Message-ID`/`In-Reply-To`/`References`; sender-asserted, never validated, never bounces, reader-reconstructed) walked as the reference architecture.

## Resolved decisions

- **Branch = A — conductor-computed PARENT with validated hint.** The engine stamps PARENT at commit from the seat's turn context (woken-on relay → active-dispatch lineage → dispatch root, first defined) exactly as it stamps FROM; the anchoring-bounce class ceases to exist; class-lineage authority gates unchanged. — *Why:* kills F11/F4 at the root; the seat never supplies an engine-known authoritative field; the operator's own monotone-validation alternative (see rejected B′) converges on A once defaults + error ergonomics are added. — *Source: operator (recommendation concurred by both m-1 seats, m-7, m-2, the ledger seeds).*
- **Unprovable `parent_hint` → FALLBACK, never bounce.** Engine stamps the computed default; the submit response carries `parent_hint_honored: no` (informed same-round-trip, free retry); the verbatim hint is preserved on the record permanently (intent + outcome + flag, all three). — *Why:* the empirical audit found **zero** true bad picks — every observed failure was a good-faith reference in an unexpected shape; reject would punish exactly the observed-good class and reopen the bounce-retry-storm channel (F10); email's one universal law (threading never blocks delivery), upgraded with verification email never had. — *Source: operator + the data.*
- **The precision cost → ACCEPTED as designed** (no monitoring owed-row). The hint's provable set is wider than the F4-narrow status quo; `IN_REPLY_TO` carries reply semantics untouched; the honored-rate is store-observable, so any future widening trigger surfaces from evidence. — *Source: operator.*
- **Sharpening-D re-home (the dispatch→routing-record citation moves to hint-class) → direction approved; the LOCK IS CONDITIONED on an explicit, BLOCKING m-4 confirm — "m-4 first."** m-4 was not on the fork packet's CC and this is its consumer surface; per the certify-your-own-surface discipline it rules before the clause folds. Everything else in this lock is unconditional. — *Source: operator.*

## Rejected alternatives

- **B (widened candidate set + grace window)** — narrows the race, doesn't remove it; fails a SEV-1 stop-the-line bar; erodes the S3 authority narrowing.
- **B′ (operator-raised, honestly weighed): keep seat-owned PARENT, fix the rule to monotone validation (accepted ∩ read-history).** Correct that records never go stale and the window rule was the bug — but B′ still needs an engine default (most relays don't care) and an error disposition (bounce = retry storms, fallback = A), and it leaves the *authoritative* graph edge lane-supplied. Pushed to its conclusion it *is* A with the seat's input renamed.
- **C (lineage lease/token)** — freezes the race in a validity window; renewal storms; heavy token/expiry machinery in the trusted core.
- **Hint-REJECT** — killed by the data: in our entire recorded history it would only ever have bounced correct references.
- **Hint-HOLD** — `held` is reserved for authority/integrity questions; spends operator attention on bookkeeping.
- **Q3 objection-window / silent-confirm** — operator chose the blocking m-4 confirm.

## Still operator-owned

- None residual — all four grill questions resolved. (The pending m-4 confirm is m-4's domain ruling, not an operator item.)

## Design-lock impact

- **m-1 folds §A as the amendment** (the packet's binding forward condition holds): supersede locked §6 `parent_picker` + the S3 five-point derivation — **rewritten, no parallel parent authorities**; PARENT becomes `system_only` computed; `parent_hint` / `parent_hint_honored` / `parent_provenance` rows added (m-2 shapes); **the Sharpening-D clause folds ONLY on m-4's confirm.**
- **m-7:** per-seat turn-context maintenance + the commit-time stamp locus (its A-1..A-4 already composable, pair-complete).
- **m-2:** the parent field leaves every form (its §3 codec amendment binds to branch A); the three hint rows land in its row-shaping.
- **Fixture obligations (build slice, red-first):** the F11 regression leg — the archived dogfood traffic replayed without livelock · hint-honored + hint-fallback fixtures · a concurrent-accept-during-submit fixture proving no parent-class bounce exists.
- **Non-binding build-slice note (from the audit):** observed hint failures are *shape variants of real references* — a generous hint-prover (dispatch-id → thread-root; known prefixes) cheaply drives the honored-rate toward 100%. Ergonomics, not design; fallback semantics don't depend on it.

*Referenced from: the s6-design integration package; the m-1 amendment fold (its DESIGN_LOCK carries this GRILL_LOCK_ID). Precedent artifact: `GRILL-LOCK-deployment-fork-2026-07-01.md`.*
