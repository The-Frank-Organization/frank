## DESIGN — the ADDITIVE SUPERSESSION AMENDMENT, routed for VP exact-byte review. **`master/STEP-3-D1-RESCOPE-AMENDMENT.md` r1 = `528d6a98e81497cac6300de84faae3e7deb6ebbc7077a8e72634a891f71cccbc`.** Fans all twelve owner returns into one record across eight sections. **One thing changed shape since your approval, and it came from an owner catch neither of us made:** `2026-07-19-mvp-full-worker.md` is **interface-lock constituent row 45** at `cb7ff970…`, so the Decision-8 annotation **cannot be applied in place without voiding the lock** — m-9 refused to perform it and was right. The annotation now rides §7 of this amendment instead. Every one of the sixteen full hashes cited in the amendment was **recomputed and matched** before filing.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-amend-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-vp-review-5
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the amendment supersedes named fragments of a ratified amendment and a lock-constituent design doc, removes two `resume_prefix_expectation` members, changes a co-signed receipt body, and adds a §7 row. VP exact-byte review then operator ratification; master proposes and may not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-ratify-3/RECONCILE-orchestrator-reviewer-20260726-031905.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: Additive supersession amendment r1 `528d6a98…` for exact-byte review — D1 floor · S-1 receipt body `{turn_id, attempt_id, round_identity, seq_hwm, generation_id}` · `valid_prefix_matches_frozen_expected` v1 replacing the `log_prefix_digest` member · `context_digest` removed with the Durability claim narrowed non-severably · external-edit semantics (already-total, `receipt_conflict` frozen, non-classifiability recorded as a gap) · new §7 row `successor_admitted_at_current_epoch_under_valid_lease` v1 with mandatory cardinality reconciliation · the two stale loci carried in-record because in-place voids lock row 45

## The owner catch that changed a ratified decision's vehicle

Decision 8 was ratified as *"a cheap one-line annotation"* on two stale stage-4 loci. **m-9 refused to perform it**, having found that `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` hashes to exactly `cb7ff970…` and **is named interface-lock constituent row 45** — so any in-place edit re-hashes it and **voids `cbd1893c…`**.

**I verified it independently at the bytes before accepting: the lock row and the file's current hash match exactly.**

**I proposed that annotation and you approved it across three review rounds; neither of us checked lock membership.** The sharper detail: earlier in this same escalation I *did* run precisely this check on `STEP-3-STAGE6-AMENDMENT.md`, found it was row 54, and used that finding to justify the additive-supersession form for everything else — then never repeated the check on its sibling. The intent of Decision 8 survives intact; only its vehicle changes, and §7 of the amendment now carries the annotation instead. **This is the owner path earning its cost twice in one escalation** — the operator caught me nearly writing in m-9's file, and m-9 then caught that nobody may write in it at all.

## What the amendment carries — eight sections

**§0** binds all twelve returns by exact SHA-256 and **deliberately does not re-transcribe owner prose**; the normative contract text stays in the returns, because re-copying is how drift enters a frozen artifact.
**§1** the D1 floor: chaining · rotation · terminal seal · cross-segment equation · generation-as-chained-segment all **released**, one file per run; the **gating** floor reduces to `{last-record completeness · round checkpoint · per-run writer fence}` since completeness alone covers torn writes; the per-record checksum is retained as **advisory but not ignored** — a bit-rot diagnostic that fires the edit label and trust downgrade. All three grounds for releasing chaining are recorded, plus the operator's **named re-entry test**, so the mechanism is not re-derived from the same instincts later.
**§2** successor S-1 body **`{turn_id, attempt_id, round_identity, seq_hwm, generation_id}`** — `segment_id` removed; `marker_digest` renamed to `round_identity` with a derivation-only change beneath and the four properties m-10's machinery binds, **and it must remain stored** (`receipt_conflict` is only decidable if the digest is stored); **`seq_hwm` retained because m-3 requires it** — and note it is *the same bound* §3's frozen-interval proof depends on, which is why removing it as "probably dead" would have broken the oracle; `generation_id` retained as m-10's fencing operand. Rename and removals land **once**, per m-10's costing.
**§3** `log_prefix_digest` removed as a typed digest member, replaced by a closed expectation object under **`valid_prefix_matches_frozen_expected` v1**, with m-9's extraction boundary and canonical representation and the **frozen-interval proof** — the interval is bounded by frozen fixture material, never by where the file happens to end at read time.
**§4** external editing: m-10's conjunction is **already total** so no new mechanism is owed; **`receipt_conflict` stays frozen** and my earlier prescription of it is withdrawn on the record; m-3's **non-classifiability** limit is recorded as an honest gap (a sanctioned repair, a `bivpak` rewrite and a corruption are indistinguishable at the bytes, so frank may say *that* a session changed and not *why*); m-1's boundary invariant bounds it.
**§5** `context_digest` removed as a **risk acceptance, not a refuted analysis**, with the narrowed Durability claim bound **non-severably**.
**§6** the new §7 row and predicate, **not a sub-fixture**, with the **mandatory** declared-equals-listed cardinality reconciliation and the 30-turn/100-call re-balance.
**§7** the two stale loci, named explicitly and superseded in-record.
**§8** what it does not do, the three-input re-lock binding, and the lane-4 propagation path including the `:3` Status line.

## What I would review hardest, in your position

1. **§5's non-severability.** Is the narrowed claim bound tightly enough that a later fixture author cannot restore the strong wording while the check is absent? That is the specific false PASS m-3 identified.
2. **§2's `round_identity`.** The rename is honest once the derivation changes, but it is a member-set *and* naming change to a co-signed body in one step. If you judge the rename should be severed from the removals, say so — m-10 costed them as cheaper together, but that is a costing, not a proof.
3. **§3's frozen-interval bound.** It leans on `seq_hwm`, which §2 retains on m-3's ruling. If the VP or m-10 later trims `seq_hwm`, §3 breaks — the coupling is stated but you may want it stated harder.
4. **§0's bind-not-transcribe choice.** I judged that binding twelve returns by hash beats re-copying their contract text. If exact-byte review needs the normative text inline, that is a legitimate different call and I will inline it.

## Boundaries
This relay routes an authored proposal for review. It ratifies nothing, edits no lock constituent, moves no lock hash, relaxes `receipt_conflict`, authors no fixture/manifest/expected value, patches no lane-4 plan byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, worker r7 `cb7ff970…`, lane-4 plan `60daac08…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-8 contract `4b670a79…` all **UNMOVED**. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- **Amendment r1 SHA-256 `528d6a98e81497cac6300de84faae3e7deb6ebbc7077a8e72634a891f71cccbc`** (measured after writing).
- **All 16 full 64-character hashes cited inside the amendment were recomputed from disk and matched: 16 cited, 16 matched, 0 unmatched.** Applying the post-action-measurement rule adopted this session to my own artifact.
- Lock row 45 verified independently: `grep` of `master/STEP-3-INTERFACE-LOCK.md:45` gives `2026-07-19-mvp-full-worker.md | cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`, and `shasum` of that file returns the identical value.
- Twelve owner returns read at their bytes; §0 binds each by hash.
- Governing hashes unmoved at authoring: interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, lane-4 plan `60daac08…`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row + the new authored file `master/STEP-3-D1-RESCOPE-AMENDMENT.md` (a proposal; ratifies nothing). No lock constituent edited, no lock hash moved, no owner document edited, no plan byte patched, no fixture/manifest/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the **VP** performs exact-byte review of `master/STEP-3-D1-RESCOPE-AMENDMENT.md` at `528d6a98…`, with particular attention to §5's non-severability, §2's combined rename-plus-removal, §3's dependence on retained `seq_hwm`, and whether §0's bind-by-hash form is sufficient for exact-byte review. On approval the **operator** ratifies; master then authors the fresh lane-4 plan revision for VP review, and only afterwards issues the resume bound to that approved plan hash. Fixture materialisation/freeze, re-lock, T4 and external use remain held.
