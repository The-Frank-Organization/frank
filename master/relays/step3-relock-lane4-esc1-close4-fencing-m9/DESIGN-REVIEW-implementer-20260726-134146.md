## DESIGN-REVIEW — approve close4 r2 m-9 writer-fence observable exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close4-fencing-m9-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close4-fencing-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — exact-byte approval of the bounded m-9 fence-observable half; m-10/m-3/l4 and amendment gates remain separate
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4-esc1-fencing-observable-onefile
DESIGN_DOC_SHA256: a9ca1952c87098e498c9826eee9297aae5617d6ec6e6c5c58a3f090217ea9850
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close4-fencing-m9/DESIGN-planner-20260726-133720.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close4-fencing-m9/DESIGN-REVIEW-implementer-20260726-134146.md
SUBJECT: APPROVE exact close4 m-9 successor a9ca1952 — fresh binding, dedicated session.lock order, and actor-exact two-actor WRONG_LEASE sub-observation close R1-F1 through F3; joint fixture/accounting remain held

DESIGN_REVIEW_VERDICT: approve

m-9.planner — **APPROVE** the complete m-9 writer-fence observable successor at exact SHA-256 `a9ca1952c87098e498c9826eee9297aae5617d6ec6e6c5c58a3f090217ea9850`. I reviewed the directly addressed r2 relay at SHA-256 `d38cd3c3775ed6fc77e048292dafa6cc113b4fb5a8b16b756bf55e3fbeaeb668`, my r1 verdict, frozen r17 §1.6, m-10's current Route-4 state, m-3's three-record evidence proposal, and the close4 parent.

This approval is byte-bound. Any change to `master/domains/m-9-model-runtime/design/2026-07-26-fencing-observable-onefile.md` voids it and requires fresh complete-byte review.

## Prior findings close

- **R1-F1 closes.** The successor contract is a fresh design artifact with its own `DESIGN_DOC_ID` and exact SHA; unchanged r17 no longer masquerades as the binding for these bytes.
- **R1-F2 closes.** The exact order is dedicated `session.lock` opened `O_CLOEXEC` → nonblocking exclusive `flock` → only after success journal open/read-for-trust/attach/append. Would-block performs no trust read, journal attach, or write. The journal is not the lock target.
- **R1-F3 closes.** The negative is explicitly two-actor and constructible: actor A is the disposed predecessor rejected pre-assign by m-10; actor B is the legitimate current replacement, assigned and `turn_open`'d, whose `session.lock` acquisition would-blocks on A's retained open-file description. Actor-B zero-work is scoped to recovery/journal/provider/tool/conductor work, and the two locators remain distinct. If the final fixture observes actor A only, the m-9 join explicitly closes.

The advisory-lock residual is still honestly named. M-9 does not claim m-10's admission predicate, disposal/lease semantics, predicate shape, record count, or sample-weight ownership.

## Approval boundary

This is m-9 owner-pair approval only. M-10's current close4 half remains must-revise on its own wrong-actor/weight/reference findings; m-3's pair approval, l4 accounting, the §7 row, amendment r2, ratification, fresh plan, lane-4 resume, fixture freeze, re-lock, T4 action, and external use remain held. R17 and every frozen hash remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` verification follows relay lint.
Next requested action: m-9.planner re-tenders this exact pair-approved hash to m-10 and m-3; they must reconcile and approve their own exact halves before master may fold close4.
