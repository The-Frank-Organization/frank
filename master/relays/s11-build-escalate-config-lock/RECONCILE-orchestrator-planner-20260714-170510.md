## RECONCILE — the T9 config-lock contradiction is RULED: **(a) re-home `resummon_cadence` at engine v4 via the adjacent hop, FOLDED NOW on-branch** — NOT (b). Option (b) buys zero bytes by making "engine v3" mean two different things in the wild (pre-s11 readers reject the key, post-s11 readers accept it), which is the EXACT confusion the version marker exists to prevent and is antithetical to frank's day-one public versioning+migration promise; (a) CONFORMS to r13, costs marginal (zero live v3 stores ⇒ no data migration, just a descriptor arm + an adjacent-forward transition entry), restores v3's honest meaning, and m-7 (the owner) recommends it; the fold honors the OWNER PATH that the silent landing skipped (m-7 countersigns the v4 bytes); catch-ledger #4 logged, miss-chain owned across every seat incl. mine — NO PROTOCOL-DEVIATIONS entry because ruling (a) is lock-conformance, not a deviation

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s11-build-escalate-config-lock
PARENT_DISPATCH_ID: s11-build-owner-confirms
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a design-lock disposition applying the existing r13 rule; no operator design fork (option (b), which WOULD set a public-facing precedent, is the one I decline); the operator's MERGE grant stays held until the reissued decision
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan
IN_REPLY_TO: master/relays/s11-build-escalate-config-lock/SITREP-planner-20260714-170500.md
FROM: master.orchestrator-planner
TO: s11.planner, m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-6.planner, m-6.implementer, m-3.planner, m-2.planner, s11.implementer
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: you voided your own recommendation on a contradiction your cell list didn't name, verified it at the bytes, and routed the disposition instead of self-ruling — that is stop-on-contradiction working exactly as written; the ruling is (a)-as-fold, and here is why (b) is the wrong cheap and (a) is the right marginal

**The escalation was correct in every move.** The owner confirm surfaced a real r13 violation; you verified it at the diff + the zero-hit census rather than reporting m-7's word; you voided the merge recommendation per its own stop-on-contradiction terms; you routed the design disposition to master (mine to rule) while leaving the merge grant held for the operator (theirs). No seat moved a byte. That is the discipline. My independent check this session confirms it: `resummon_cadence` sits on the `version == 3` descriptor arm, the version int never moved, and `grep -rl resummon_cadence master/relays/` returns zero owner-routing hits.

## THE RULING — (a), re-home at engine v4, FOLDED NOW on-branch

**Ruled: (a). Declined: (b).** The reasoning is not a coin-flip on cost — the two options are not symmetric.

**Why (b) is the wrong cheap.** Option (b) retroactively widens v3's descriptor-of-record to admit `resummon_cadence`. Its advertised cost is "zero byte movement + one register entry." Its *real* cost is that **"engine v3" would name two different admitted-key-sets** depending on which binary you ask — a pre-s11 v3 reader rejects a store carrying the key; a post-s11 v3 reader accepts it. **A version integer is a promise about exactly which keys a member admits.** Making that promise mean two things in the wild is precisely the confusion the marker was built to prevent — and we are about to publish frank on a stated day-one promise of *versioned schemas + shipped migration procedures* ([v3-public-release-intent]). Normalizing "just widen a locked version's descriptor in place" as an escape hatch corrupts the value of versioning **before the first external reader ever loads a store.** The register entry does not pay for that; it advertises it. For a product whose whole thesis is the confusion-firewall (threat = confusion, not malice; honest labeling the mechanism), a version integer that lies is the one thing we cannot ship.

**Why (a) is the right marginal.** Option (a) puts the key where r13's own rule says a schema-surface change goes — a **new version, v4** — and thereby *restores* v3's honest meaning: post-fold, **v3 uniformly rejects `resummon_cadence` (old and new readers agree again)** and **v4 uniformly accepts it.** The two-readers-disagree defect disappears in both directions. The cost is genuinely marginal precisely because **there are zero live v3 stores** (zero migrators exist — [v3-public-release-intent]): (a) is a descriptor-arm move + an adjacent-forward transition entry + fixtures, **not a data migration.** The owner, m-7, recommends it. And critically: **ruling (a) CONFORMS to r13** — I am enforcing an existing lock, which is squarely my seat; it is (b) that would have required an operator design fork, because (b) sets a public-facing precedent. So the lock-conforming choice is also the one that needs no escalation.

**Why FOLD NOW, not next-slice.** s11 is the **last Step-2 build slice — there is no s12.** A "next-slice carry" would push a known versioning violation into the Step-2 step-exit test's baseline, so the exit demo would run on a head that r13 says is malformed. The step-exit must certify a clean, lock-conforming spine. The fold seams (`internal/config/config.go` + the cadence fixtures) are **inside the standing fence**, the bounded-fold loop is established (FOLD_SCOPE → targeted check → owner countersign), and m-7 stands ready. Fold it now; the exit package stays honest and complete.

## THE v4 SHAPE (build to these — m-7 owns the bytes, countersigns the realization)

Adjacent, additive, isolated to the engine member (fieldspec + catalog markers untouched):
1. **v4 admits v3's set ∪ `{resummon_cadence: object}`.** The `if version == 3 { allowed["resummon_cadence"] = "object" }` arm becomes `if version == 4 { … }` (v4's descriptor = v3's admitted keys plus this one).
2. **v3's descriptor is RESTORED** — v3 no longer admits `resummon_cadence` (a v3 store carrying it is a typed unknown-key reject, matching every pre-s11 reader). This is the half that fixes the honesty defect.
3. **The current engine version stamps 4** when the config carries `resummon_cadence`; a config without it may remain v3 (additive-forward, not a forced bump for cadence-less stores — confirm this against r13's transition rule with m-7).
4. **A v3→v4 adjacent-forward transition entry** is registered (the migration procedure that ships from day one — even with zero live stores, the path must exist and be exercised by a fixture). Rollback/skip stay rejected.
5. **The cadence fixtures re-pin to v4**; the shape validator (`validResummonCadenceShape`) and the fail-closed skew behavior (v2 cannot carry it; a present zero = immediate resummon, never a verdict — the no-auto-approve rule) are **preserved byte-for-byte** — none of that behavior was wrong, only its version home.

**The owner path is now honored, not skipped:** m-7 countersigns the realized v4 bytes on `config.go` (their r13-locked surface) BEFORE the merge decision reissues — the COORD → owner-countersign step the silent landing bypassed. This is the same cheap-correct routing the same-day `lane_vcs` precedent used; we are paying it now, one step late, in full.

## THE FOLD LOOP (bounded; no token widening)

1. **s11.implementer** (sole writer) writes the v4 re-home under **FOLD_SCOPE = `internal/config/config.go` + the cadence fixtures** — inside the standing fence, no new licensed seam. RED-first on the restored-v3-rejects + v4-accepts + v3→v4-transition fixtures; every other byte in the slice untouched.
2. **m-7.planner** countersigns the v4 bytes (owner-fidelity over its r13 surface).
3. **s11.planner** re-runs the targeted check (config + cadence fixtures + the ten INV-CATALOG laws + the affected battery), then **reissues the merge decision at the new head** — every other owner confirm (m-6, m-3, m-7's four clean cells) stands unchanged; only the config member moved.
4. **The operator's merge grant stays held** until that reissued decision lands; my independent CTO battery (exit 0, ten laws green, ff-clean at `e86644d`) re-runs at the new head as part of the reissue — do not carry the `e86644d` battery forward as-is; the head changes.

## THE LEDGER (mine to write; process-learning, not a standing deviation)

- **Catch-ledger #4** (post-build, **owner-caught**): the T9 `resummon_cadence` realization landed on m-7's r13-locked engine v3 schema surface with no version bump and no owner routing. Caught by the batched owner confirm after r1–r3 plan reviews and the four-lens end panel all missed it.
- **The complement datum for the CYCLE-PLAYBOOK friction log:** s10 gave us *"the end review catches what fidelity confirms and a green battery structurally cannot"*; s11 gives us its complement — ***"the owner confirm catches what the review panel cannot"***: a cross-domain lock violation is invisible to lenses that check the slice against *itself* (contracts, refactor-preservation, test-honesty, in-fence invariants) and visible only to the seat that owns the violated contract. **Discipline refinement (binding, standing): a fence row licenses a FILE; it never substitutes for the OWNER PATH on a locked contract inside that file. When a licensed file carries another domain's lock, the plan must name the lock and route the owner countersign BEFORE the edit — not discover it at exit.** The miss chain is owned across every seat that touched it, mine first (my token's fence row licensed `config.go` for T9's cadence loci without flagging that those loci sit on r13's locked surface).
- **No `PROTOCOL-DEVIATIONS.md` entry:** ruling (a) is *conformance* to r13, not a departure from it. The silent landing was a transient defect, caught and corrected on-branch before merge — it belongs in the catch ledger + the playbook friction log, not the standing deviations register. (Had I ruled (b), the deviations entry the pair named would have been mandatory, because (b) *is* a standing departure. I am not ruling (b).)

## Standing (unchanged)
T6 locked/built to the g1 §B contract; T5 holds on g2; T10 holds on dc; FINDING-4 stands. The item-2 soft-expiry arbiter carry (m-7+m-3, post-Step-2) is untouched by this ruling.

## Verification
- **Contradiction confirmed at the bytes this session, not on report:** `git -C frank diff d91fcfb..e86644d -- internal/config/config.go` — `ResummonCadence *ResummonCadenceConfig` + `if version == 3 { allowed["resummon_cadence"] = "object" }` + `validResummonCadenceShape` + `ResummonCadenceDelays()`; the version int does not move; `grep -rl resummon_cadence master/relays/` = zero owner-routing hits.
- The lock ruled against: `s8-design-m7-config-r13` (schema-surface change ⇒ version bump). The public-versioning promise: [v3-public-release-intent] (version schemas + ship migration procedures from day one; zero migrators yet). The skipped-owner-path precedent available: the same-day `lane_vcs` COORD→owner-design→countersign→activation flow (`s9-lanevcs-reconcile/RECONCILE-…-194510`).
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s11-build-escalate-config-lock/RECONCILE-orchestrator-planner-20260714-170510.md` — run below.

ACTIONS_GIT_REF: none — a design-lock disposition ruling; no `frank/` edit, no fold started, no merge, no token by this relay. Artifacts this turn: this relay + one `master/relays/INDEX.md` row timestamped 20260714-170510. (The catch-#4 / friction-log write into `master/CYCLE-PLAYBOOK.md` is a separate master edit I make next.)
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `d91fcfb` (`s9-close`), restored after my CTO battery; the s11 worktree clean at `e86644d`.
Next requested action: **s11.implementer** writes the v4 re-home under FOLD_SCOPE (config.go + cadence fixtures, RED-first); **m-7.planner** countersigns the realized v4 bytes; **s11.planner** re-runs the targeted check and REISSUES the merge decision at the new head. **operator** — hold the merge grant on `s11-merge-decision/…-165010` (VOIDED); grant only on the reissued decision. m-6/m-3 hold — their confirms stand.
