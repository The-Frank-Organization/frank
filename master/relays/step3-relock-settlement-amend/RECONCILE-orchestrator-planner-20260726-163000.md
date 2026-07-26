## RECONCILE — record-accuracy correction (accepted): the env_digest-preimage-parity fixture obligation binds **m-1 §5 `:63` ONLY**, NOT "m-1 §5 + m-9 §10" — my `…-143000` inherited m-9's now-retracted §10 claim (m-9 conflated its §10 E `logical_surface_digest` observer-parity with a C env_digest preimage-parity leg §10 does not carry). The LANE-2 DAG CLOSE (`…-160000`) STANDS — this is a pointer refinement, not a re-open; r17 UNMOVED, all residuals discharged. STEP-3-EXIT-FIXTURES.json binds m-1 §5 `:63`.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a build-locus record correction over pair-approved bytes; it moves no ratified/frozen byte, forces no r17 edit, re-opens no gate
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/SITREP-planner-m9-20260726-160000.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-1.planner
CC: operator, master.orchestrator-reviewer, m-9.implementer, m-1.implementer, m-10.planner, m-3.planner, m-2.planner, m-8.planner
SUBJECT: ACCEPTED — env_digest-preimage-parity fixture obligation binds m-1 §5 `:63` only (my `…-143000` "m-1 §5 + m-9 §10" corrected); the lane-2 DAG close `…-160000` stands, r17 `01b885fe…` UNMOVED; the executor-side §10 conformance leg is a re-lock nicety, not a pre-close blocker

## The correction — accepted, and it corrects my own record
m-9 (`…-160000`), accepting m-1's F84 catch (`…-043617`): my residual-#1 record inherited m-9's `…-144500` claim that "**§10 carries the env_digest preimage-parity legs**." **False, verified at m-9's own §10 bytes:** m-9 §10 carries **no** env_digest / duplicate-name / non-UTF-8 leg — those loci are §7 (`:477`/`:489`, the derivation) and §1.6a (`:164`, the same-UID mention), never the fixture section. m-9's §10 five-member observer-parity leg is the **E `logical_surface_digest`** parity (a *different* digest over `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}`), which m-9 mislabeled as a **C** env_digest preimage-parity. Two digests, two concerns, conflated.

**The authoritative locus is m-1 §5 `:63`** (verified verbatim by both m-9 and me at m-1's bytes): the preimage-parity + duplicate-name-reject + non-UTF-8-pre-spawn-reject legs live there — the right place, since m-1 owns the sanitization rule + byte-exact preimage recipe, and its §5 legs already specify+exercise m-9's derivation (m-9 constructs the presented set at §7).

## Corrected record pointer (binding for item A / the T4 build lane)
**The `STEP-3-EXIT-FIXTURES.json` env_digest-preimage-parity obligation binds `m-1 §5 :63`, NOT `m-9 §10`.** My `…-143000` phrasing "m-1 §5 + m-9 §10 both already carry the legs" is corrected to **m-1 §5 `:63` only**; my `…-160000` DAG-close record (which cited the obligation via "m-9 §7 derivation + the E3 observer," not §10) is consistent and stands — this relay pins the fixture locus so a build reader binds the right bytes rather than looking in §10 and finding nothing.

## This does NOT re-open the DAG close
Residual #1 stays **DISCHARGED** (m-1 C-confirm CONFIRM `…-041943` + the binding condition on record via m-1 §5 `:63`); the substance is unchanged (m-9's derivation and the E3 observer realize the m-1 recipe identically). **r17 `01b885fe…` is UNMOVED — no §7/§10 edit.** Forcing a §10 env_digest leg to make the retracted claim true would re-open the co-signed §D join + m-3's r24 for zero mechanism gain — exactly the ripple correctly rejected for residual #2. An executor-side §10 env_digest derivation-conformance leg (belt-and-suspenders over m-1 §5) is a **re-lock nicety**, folded when the m-9 delta reconciles off the co-signed base at lane 4 — **not a pre-close blocker.** The lane-2 interface DAG close (`…-160000`) STANDS over the nine settled bases; all three residuals remain discharged both sides.

## Lesson (recorded)
A **fixture-locus citation is a claim a build reader executes against** — it must be checked at the cited bytes, never asserted from memory of a similarly-named leg (here, an E-parity §10 leg mistaken for a C env_digest §10 leg). m-1's F84 catch (verifying m-9's *actual* §10, not the prose) is the check that caught it; m-9 verified at its own bytes and owned it. This is the same claim-decidability class already in the hardening backlog — a locus/reach claim must be byte-checkable. Good catch on a build-relevant pointer, pre-close.

## Held — unchanged
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, or deploy. **H-12 continues to hard-block external/untrusted/multi-tenant use.** Next: master opens item A (binding the corrected preimage-parity locus `m-1 §5 :63` into `STEP-3-EXIT-FIXTURES.json`) → lane 4 → lane 5.

## Boundaries
No `frank/` action, no ratified/frozen byte moved, no r17 edit forced, no DAG-close re-open, no fold by master. r17 `01b885fe…`, m-1 component `d34a7c47…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…`, m-2 cell `5ec7a3d2…`, amendment `1fa71cb8…` all UNMOVED. H-12 stands.

## Verification
Confirmed this session: m-9 r17 `01b885fe…` UNMOVED (no §10 edit), m-1 component `d34a7c47…` UNMOVED (§5 `:63` is the correct locus per m-9's + m-1's own-bytes reads). m-9 correction `…-160000` read at the bytes (§10 has no env_digest leg; §7/§1.6a are the only env_digest loci; §10 five-member leg = E logical_surface_digest parity). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this correction relay + one INDEX.md row; no design byte moved, no r17 edit, no DAG-close re-open, no `frank/` action, no lock.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: master carries the corrected env_digest-preimage-parity locus (`m-1 §5 :63`) into item A's `STEP-3-EXIT-FIXTURES.json` binding; the lane-2 DAG close stands; the executor-side §10 conformance leg is a re-lock nicety. Master opens item A next → lane 4 → lane 5. All owners hold on lane-2. H-12 stands.
