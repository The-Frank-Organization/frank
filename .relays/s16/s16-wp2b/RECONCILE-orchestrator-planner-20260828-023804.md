## RECONCILE — WP2b OPENED: the bounded fix + release-pipeline lane — the pair runs ONE plan round covering the THREE ruled work items (the catalog-key serialization fix · the two-derivation-point F58 wiring fix · the canonical release-pipeline realization); the V32 law governs (plan → review → all-in SCOPE_DIFF → master's fresh direct dispatch); no code byte before the dispatch

**The governing basis (all banked on the master root under `s16-wp2-owners`/`s16-wp2-disp`; cite by relay path):** m-9's Q-CATALOG ruling (`SITREP-planner-20260828-022110.md` — the comparand is `7fae5fc1…`; live `form_digest` contract-forbidden in F58; the lock's row keys govern; the §1(d) fix sequence) · m-2's keys-only proof (`DESIGN-planner-20260828-022111.md`) · m-10's tautology finding + two-derivations law + structural gate acceptance (`SITREP-planner-20260828-022057.md`) · m-8's ratified F63 proposal (`SITREP-planner-20260828-022252.md`, ratified at `s16-wp2-disp` conditional on m-9's co-sign) · master's dispositions (`s16-wp2-disp/…`).

**The three work items for plan-6 (mechanics yours; contracts the owners'):**
1. **The catalog serialization fix:** `internal/worker/catalog` serializes the F58 identity rows under the LOCKED row member keys (`{canonical_name, tool_schema_digest, tool_impl_catalog_version, form_schema_mapping_version?}`, absence-encoded, sorted, JCS); the builder consumes m-2's three produced digests as CONSTANTS (never a render, never a store read — they already are; keep it that way provably); `ExpectedDigest` recomputed from the corrected builder. **Expected outcome: equality with `7fae5fc1…`. If the corrected recompute differs, that is a STOP + finding to m-9 — the constant is NEVER edited to whatever the build emits.**
2. **The F58 wiring fix (the tautology):** `cmd/frank-app` derives the manifest's `tool_catalog_digest` member and the gate's comparand from TWO INDEPENDENT POINTS — the member from the composition's own catalog derivation, the comparand injected from the shipped pin (`catalog.ExpectedDigest`) — so `gate.go:57` binds something again; the LIVE presented-surface digest stays ONLY on `logical_surface_digest` (attempt grain, INV-E1), out of the manifest member and out of the gate. m-10's structural acceptance stands (pin-equality against an injected comparand; only the WIRING moves, plus a rename if "Shipped" needs one); the composed battery and the WP1 evidence assertions may need their expectations updated to the STATIC member value — that is in-scope, honestly labeled.
3. **The canonical release pipeline (the ratified m-8 §2/§3 shape):** the canonical build step — `go build -trimpath`, the PINNED Go toolchain recorded, per `GOOS/GOARCH` — emitting the FOUR ratified F63 members {frank-app · frank-worker · frank-connector · frank-broker} into `dist/<goos>-<goarch>/` with the RELEASE-MANIFEST (toolchain version · build flags · per-binary lowercase-hex SHA-256). NO binding claim — the pipeline PRODUCES what the separate Master+VP binding act will bind; `frank`/`frank-mcp` are NOT members (F65 / excluded). If the pipeline surface (a build script, `dist/` layout) falls outside the standing plan-5 §2.4 fence, the plan says so and the SCOPE_DIFF lists it — the dispatch will carry the fence extension explicitly.
- **Rails unchanged:** the regression floor at every commit; the CT-G03/frozen-member fences; findings UP never absorbed; the draft PR refreshed (V30); WP1/WP2's banked claims stand at their filed grain — this lane is forward-only. **Sequencing note:** the plan round may run NOW; master's dispatch will not issue before m-9's Q-F63 co-sign lands (the one open ratification leg).

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16-wp2b-open
PARENT_DISPATCH_ID: s16-wp2-close
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a chartered plan-round opening; the token remains master's fresh direct dispatch; the operator's gates remain the WP5 MERGE-GATE and the Step-3 ratification
IN_REPLY_TO: s16-wp2/SITREP-planner-20260828-015935.md
FROM: master.orchestrator-planner
TO: s16.planner, s16.implementer
CC: master.orchestrator-reviewer, operator, m-9.planner, m-2.planner, m-10.planner, m-8.planner
SUBJECT: WP2b opened - one plan round over the three ruled items (catalog locked-key serialization with the 7fae5fc1 expected outcome and the STOP-if-differs law; the two-derivation-point gate wiring fix; the canonical -trimpath four-member release pipeline into dist/ + RELEASE-MANIFEST); V32 governs; plan round may run now, the dispatch waits on m-9's co-sign

ACTIONS_GIT_REF: engine-lane governance act — drafted at .engine/drafts/master.orchestrator-planner/ on the s16 root and submitted through the v2.9.1 client; the governing rulings verified at the master root this act (six returns linted; the tautology re-verified at bytes; the instrument hash exact); no source, test, branch, commit, push, PR, build, or merge byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M master/relays2/INDEX.md
?? master/domains/m-7-conductor-core/design/2026-08-28-f65-conductor-identity-grain.md
?? master/relays2/s16-wp2-owners/DESIGN-planner-20260828-022111.md
?? master/relays2/s16-wp2-owners/DESIGN-planner-20260828-022209.md
?? master/relays2/s16-wp2-owners/SITREP-planner-20260828-022057.md
?? master/relays2/s16-wp2-owners/SITREP-planner-20260828-022110.md
?? master/relays2/s16-wp2-owners/SITREP-planner-20260828-022143.md
?? master/relays2/s16-wp2-owners/SITREP-planner-20260828-022252.md
