## SITREP — m-1 fidelity verdict on s7 INV-CATALOG rows 2/4/5/6: rows 2/4/5 confirm; row 6 must revise its canonical store-path census

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-fidelity-m1
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — scoped fidelity return; VP integration and operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-fidelity-m1/SITREP-orchestrator-planner-20260710-112541.md
FROM: m-1.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-1.planner, m-7.planner, m-7.implementer
SUBJECT: rows 2/4/5 confirm at s7-inv-catalog@35aabb9; row 6 must revise because its forbidden-path set omits current canonical store families even though its seat-family census mechanics are sound

**VERDICT: must-revise row 6 only. Rows 2, 4, and 5 confirm.** The focused four-law battery and the whole `test/invariants` package pass at `35aabb9`; the blocker is contract coverage, not a current red test.

### Row verdicts

1. **Row 2 — `TestLawThreeVerbSurface`: CONFIRM.** The locked surface remains exactly `submit` / `project` / `read` (`master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:10`). The check populates all current `channel.ToolSet` fields, pins the literal field census `{Submit, Project, Read, Describe}`, asserts `ListTools` returns exactly the three verbs, and rejects `tools/call` for both `describe` and an arbitrary name (`test/invariants/terminal_surface_test.go:106-184`). The internal description callback does not become a fourth callable seat verb.

2. **Row 4 — `TestLawDerivedOnlyActivation`: CONFIRM.** The catalog claim is correctly scoped to the seat lifecycle (`catalog.v1.json:102-106`). The test derives `minted` from the committed `seat_mint`, derives `active` from the first accepted seat-authored submit, keeps the activation reference fixed under a later accept, derives `bound_now` from live channel state, and proves restart yields active-but-disconnected (`lifecycle_writer_test.go:25-198`). Its record/header/body/registry sweep rejects persisted activation/bound markers (`:299-327`). This matches the committed-pivot generation boundary, accepted-`FROM=<seat>` activation rule, restart-empty `bound`, and no-marker boundary (`master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:95-111`).

3. **Row 5 — `TestLawSoleGovernedWriter`: CONFIRM.** The claim text preserves the exact bounded grain — “sole governed write path” plus the D5 same-uid direct-store residual (`catalog.v1.json:109-113`) — and does not re-expand to structural sole-writer language. The check proves a second owner of the canonical root gets typed `root-lock-held` and exercises a governed mutation through the intake writer plus serialized engine loop (`lifecycle_writer_test.go:201-272`). This is faithful to I1-P's canonical-root, pre-phase-0, serve-lifetime lock and loser-refuses-service rule (`2026-07-06-s6-transport-amendments.md:87`) while retaining the locked D5 ceiling (`2026-06-28-v3-trust-identity-design.md:99,105-108`).

4. **Row 6 — `TestLawPathHygiene`: MUST REVISE (m-1 store half).** The seat-delivered-family side is sound: six families are named and counted; the `seat_mint` reply has exactly the two operator-only carve-outs; planted-leak, non-operator carve-out, unregistered-family, and unregistered-sink negatives all bite (`path_hygiene_test.go:47-169`). That satisfies the mechanical “future seat-visible family turns red” leg.

   The canonical-path-family side is incomplete. The locked law forbids canonical store/config/outbox paths in every seat-delivered surface (`master/ARCHITECTURE.md:478`), and the effective r3 contract requires the complete corpus to be scanned across those families (`master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-030148.md:23`). Production `store.Open` currently creates `records`, `staging`, `journal`, `projections/relays`, `mailboxes`, `outbox`, and `binding` under the canonical root (`internal/store/store.go:48-62`); production also has the lazy `quarantine` home and root `conductor.lock`. But the scanner's literal family set names only `/records/`, `/staging/`, `/outbox/`, `/binding/`, `registry.json`, and `engine.json` (`path_hygiene_test.go:98-113`). `journal`, `projections`, `mailboxes`, `quarantine`, and `conductor.lock` are not censused. The two captured root strings happen to cover some current samples, but they do not make those omitted canonical families a named, mechanically checked law: a path under another scratch/runtime root can evade the family scan unless it also contains one of the listed fragments.

### Required narrow fold for row 6

- Keep the six seat-delivered families, carve-outs, sink census, and existing negatives unchanged.
- Add an explicit complete canonical path-family census for the current store/config/outbox layout. At minimum it must cover `records`, `staging`, `journal`, `projections`, `mailboxes`, `outbox`, `binding`, `quarantine`, `conductor.lock`, `config/engine.json`, and `config/fieldspec/registry.json`.
- Mechanically tie that census to the production layout (or to a live initialized scratch root exercising the lazy homes), so a newly introduced canonical store/config/outbox home turns `TestLawPathHygiene` red until registered.
- Add a table-driven planted-path negative for every censused path family, using a root different from the positive captures, so each family is proven to bite independently of exact temp-root equality.
- Test-only fold. Any production edit or contract weakening routes back to master; no m-1 redesign is authorized.

### Verification

- Pinned worktree: `/Users/jack/.config/superpowers/worktrees/frank/s7-inv-catalog` at `35aabb9dfae324181196cc7dce0af624d25dd26a` on `s7-inv-catalog`.
- `go test -count=1 ./test/invariants -run 'TestLaw(ThreeVerbSurface|DerivedOnlyActivation|SoleGovernedWriter|PathHygiene)$' -v` — PASS; all four laws and all four row-6 negative subtests pass.
- `go test -count=1 ./test/invariants` — PASS (`ok github.com/jackli/frank/test/invariants`).
- Base-to-tip scope: no production-code changes; s7 code surface is `test/invariants/**` plus sprint docs/relay trail.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s7-fidelity-m1/SITREP-implementer-20260710-113129.md` — OK.

ACTIONS_GIT_REF: no `frank` code or git action by this report-only review; durable artifacts are `master/relays/s7-fidelity-m1/SITREP-implementer-20260710-113129.md` and its append-only `master/relays/INDEX.md` row.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable — `fatal: not a git repository (or any of the parent directories): .git`
- `frank/`: none — clean tree
- pinned worktree `s7-inv-catalog@35aabb9`:
  ` M .relays/s7/INDEX.md`
  `?? .relays/s7/s7-inv-catalog-impl/FOLD_SCOPE-implementer-20260710-111831.md`
  `?? .relays/s7/s7-inv-catalog-impl/IMPL-implementer-20260710-035937.md`
  `?? .relays/s7/s7-inv-catalog-impl/RECONCILE-planner-20260710-040806.md`
  `?? .relays/s7/s7-inv-catalog-impl/REVIEW-FOLD-implementer-20260710-112144.md`
Next requested action: fold the row-6 test-only correction, rerun the focused law and full invariant package, then return the same row for narrow m-1 fidelity confirmation before VP integration. Rows 2/4/5 need no re-review unless changed.
