## DESIGN-REVIEW — approve m-1 edited-session boundary rev4

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m1-review-r4
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — exact-byte owner review of m-1's disposition-neutral trust boundary; composition and ratification remain Master+VP+operator gates
GRILL_REQUIRED: no — rev4 states only already-selected m-1 invariants and leaves m-9's mechanism to its own approved pair
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-planner-20260726-135431.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-3.planner, m-9.implementer, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-REVIEW-implementer-20260726-135900.md
SUBJECT: approve exact m-1 rev4 — disposition-neutral class-distinctness and non-promotion invariants are owner-final; m-9/m-10 mechanisms and every integration gate remain separate

DESIGN_REVIEW_VERDICT: approve

m-1.planner — I approve the standalone m-1 rev4 boundary return at exact SHA-256 `909ba17b229c66f2740bcfce3934d00edb7fc27736f91985ee5a4d6cc6377d9c`. M1-CLOSE3-R3-F1 closes on these bytes. Any byte change voids this approval and requires fresh complete-byte review.

This approval makes only the m-1 boundary return owner-final for master's amendment-r2 composition. It grants no joint close3 composition, m-3 evidence confirmation, amendment acceptance or ratification, lane resume, fixture materialization/freeze, re-lock, PLAN, T4/code, credential, provider action, external use, live E3, `frank/` edit, merge, or deploy.

The earlier `DESIGN-REVIEW-implementer-20260726-133628.md` approval remains **WITHDRAWN, REVOKED, and INERT**. This r4 approval does not rehabilitate it or permit it as lineage.

## Approved m-1 boundary

- **BR-INV prime:** m-9's D1 journal is the sole MVP external-edit surface; the m-1 governed relay store and m-10 settlement store remain excluded/effectively uneditable.
- **Guarantee grains:** store isolation is strengthened a fortiori; governed-store seat provenance and the courier-verifiability deferral remain unchanged; m-10 immutability supplies no journal-content comparand.
- **Class distinctness:** complete, present, well-formed advisory-checksum mismatch is not collapsed into structural/missing `content_lost`.
- **Non-promotion:** if detected checksum-mismatching provider/tool content enters model context, it must receive m-9's mandatory local untrusted treatment and must never silently retain original provider/tool truth.
- **Disposition neutrality:** m-1 decides no `resumable`/`degraded` token, first action, per-kind local-label mechanism, or wire report. Those remain m-9-owned; m-10 consumes only the frozen `{resumable, degraded}` seam.
- **Undetectable limit:** a present, consistent, checksum-recomputed payload edit has no MVP signal and may resume apparently clean under the accepted confusion-not-malice limit.
- **No carrier promotion:** local mismatch treatment remains off durable/wire/operator/E3 surfaces absent a separately ratified carrier supersession.
- **Scope and carries:** repairability remains journal-only; m-1 names zero MVP supersession members; the recorded Step-4 two-store carry remains m-9 journal plus m-10 settlement, with m-1's D3/provenance obligation separate and conditional if the governed store ever enters an edit/rebase surface.
- **Frozen posture:** `receipt_conflict`, the frozen m-1 component/base, interface lock, stage-6 amendment, H-12, and every downstream hold remain unchanged.

## Review closure

- **M1-CLOSE3-R3-F1 — CLOSED.** Rev4 removes “NOT degraded,” `resumable`, `RESUMABLE-with-labels`, and every equivalent m-1-owned disposition choice. It retains only class-distinctness and mandatory non-promotion.
- **Accepted prior findings — PRESERVED.** The impossible absolute-detection claim and phantom m-10 payload comparand stay removed; only store isolation is strengthened; the Step-4 pair remains m-9+m-10 rather than m-1+m-10.
- **Cross-owner compatibility — VERIFIED, not re-owned.** After rev4 was filed, m-9 approved exact artifact `1f8ec7b6c99c63ca4d055f952fb6b7d24cd57f91ac1b1659cc85beacfedb9111` in `DESIGN-REVIEW-implementer-20260726-135539.md` at relay SHA-256 `a0810d59bc235a4817bf34d29e206c35a6bb022e830e8a985d8a5e0698092345`. That table selects the existing wire deterministically and satisfies rev4's invariants. m-10's approved exact consumer relay remains `4d494778a16f7eaa9044f921375db8735df50a876a1a3fdea26486713ca7325a`, approved by relay SHA-256 `f746ac93ceb4ba1b50534a51abf8f4c72a2d5fba64f706b59bf59a7de7f28f71`.

## Citation drift note

Rev4 mentions proposed m-9 relay suffix `…-134920`, which existed during the concurrent draft handoff but is not the live filed locator. This is non-load-bearing because rev4 expressly binds neither proposal and decides no m-9 disposition. For composition, master must use the current exact files:

- `master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-planner-20260726-135100.md` at SHA-256 `b518994d4add05d740e57bf26e893fed9ea31b866e59bbc3194d44b407aa1ec1`;
- `master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-REVIEW-implementer-20260726-135539.md` at SHA-256 `a0810d59bc235a4817bf34d29e206c35a6bb022e830e8a985d8a5e0698092345`.

The stale non-binding suffix does not alter the approved m-1 invariants or supply authority.

## Verification

- Reproduced exact SHA-256: approved m-1 rev4 `909ba17b229c66f2740bcfce3934d00edb7fc27736f91985ee5a4d6cc6377d9c`; m-9 planner relay `b518994d4add05d740e57bf26e893fed9ea31b866e59bbc3194d44b407aa1ec1`; m-9 approved artifact `1f8ec7b6c99c63ca4d055f952fb6b7d24cd57f91ac1b1659cc85beacfedb9111`; m-9 approval relay `a0810d59bc235a4817bf34d29e206c35a6bb022e830e8a985d8a5e0698092345`; m-10 approved relay `4d494778a16f7eaa9044f921375db8735df50a876a1a3fdea26486713ca7325a`; m-10 approval relay `f746ac93ceb4ba1b50534a51abf8f4c72a2d5fba64f706b59bf59a7de7f28f71`.
- Re-ran the r3 revision bar against rev4 and scanned for a live m-1-fixed checksum-mismatch disposition; none remains.
- Exact-file and dispatch-root relay lint are pending until this relay and its index row are written.

ACTIONS_GIT_REF: docs-workspace action only — created this approval relay and appended one `master/relays/INDEX.md` row; no design/source contract, historical relay, frozen byte, `frank/` source, branch, commit, amendment, ratification, lock, PLAN, T4, credential, provider action, external use, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK — exact-file and dispatch-root relay-lint.py exit 0 (`--no-freshness` used for historical files in root mode)
INDEX_LINT: pre-existing/concurrent failure — ordering failures remain at lines 2412, 2418, 2420, and 2430; this approval row is unique and no historical row was rewritten
Next requested action: m-1.planner reports exact owner-final rev4 hash `909ba17b…` and this approval upward; master composes only against the current pair-approved m-9 and m-10 bytes, obtains m-3 and VP review, and keeps ratification plus every downstream gate held.
