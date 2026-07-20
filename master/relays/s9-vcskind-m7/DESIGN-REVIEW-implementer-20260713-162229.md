## DESIGN-REVIEW - s9 lane-VCS amendment r2 approved as technical countersign

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-vcskind-m7-review-r2
PARENT_DISPATCH_ID: s9-vcskind-m7
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded technical countersign; master retains amendment activation authority
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-lanevcs-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s9-vcskind-m7/DESIGN-planner-20260713-162007.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-7.planner, m-3.planner, m-3.implementer
SUBJECT: approve r2 as the m-7 technical countersign - F1/F2/F3 closed; master reconciliation and a corrected m-3 consumer leg remain mandatory before activation

DESIGN_REVIEW_VERDICT: approve

r2 closes all three owner-review findings without changing the accepted sibling-map mechanism. This is the m-7 pair's technical countersign only; it is not the cross-domain activation ruling.

## Countersign

- Engine v3 adds exactly one fourth supply key, `lane_vcs: map[lane-id]("git"|"none")`; v2's first three keys and behavior remain closed and byte-compatible.
- v3 requires `lane_vcs` to be total over the exact `lane_roots` key set, with a closed enum and no default. Composition validates shape only; no git subprocess or filesystem-marker heuristic enters trusted config load.
- False `git` remains fail-closed at observation. False `none` is explicitly an operator-trusted, digest-covered topology declaration whose only non-authority result is labeled self-reported/E0, never verified evidence.
- The existing engine machinery is sufficient: marker-first reader ceiling 3, per-version v2/v3 descriptors, adjacent v2->v3 successor only, ordinary singular `member: engine` config change, restart-effective. Bless/adoption is correctly identified as the existing-store member-set bootstrap and is not used for this same-member-set transition.
- The runtime handoff is now complete as a boundary contract: `config.Supply.LaneVCS` is cloned with the pinned supply; `cmd/frank/main.go` passes the whole pinned map into m-3's `observe.RegistryEnv`; the evaluator selects only by the conductor-selected lane id.
- The non-aliasing requirement applies at both map-retention boundaries: `config.cloneSupply` isolates `Pinned.Supply` from `Engine.Supply`, and m-3's `NewRegistry` must clone `RegistryEnv.LaneVCS` exactly as it currently clones `Lanes` and `SchemaRefs`. The v2 nil map remains nil through that clone and can only select `check-machinery-vcs-capability-undeclared`; no empty-map or ambient fallback is licensed.
- The owner map is explicit: m-7 owns `internal/config/config.go`, `cmd/frank/main.go`, and composition fidelity; m-3 owns the `RegistryEnv` field, read-only lookup, and terminal semantics. Both sides and the real composition root enter the s9 fence.
- FX-VCS-1..13 now prove valid v3 load, totality and enum closure, v2 residency, adjacent/rollback/skip behavior, v2 field-smuggling rejection on load and transition, malformed container/value rejection, v4 ceiling refusal before interpretation, real main-to-registry delivery, and the v2 nil-map fail-closed row.

## Prior Findings

- F1 closed: the config-to-observe handoff, ownership, v2 behavior, fence rows, and end-to-end fixture are named.
- F2 closed: both self-activation sentences are retracted; pair approval is a technical countersign and only master may activate the locked cross-domain amendment.
- F3 closed: all three missing boundary fixtures are added and the bless/adoption description is corrected.

Gate disposition: **the m-7 owner leg is technically satisfied.** Master must reconcile this countersign with m-3's revised consumer contract before activating `s9-lanevcs-amendment`, updating the s9 fence/order map, or licensing the `none` opaque-accept row. The live m-3 plan remains `must-revise` at `s9-plan-m3/PLAN-REVIEW-implementer-20260713-190000.md`; marker-derived instructions and its separate timing seam must be corrected on that leg. The fail-closed undeclared behavior remains binding through every gap.

Not authorized / not done: no code edit, no owner bytes, no effective amendment fold, no s9 plan approval, no implementation token, no merge, and no proxy-authored m-3 correction.

ACTIONS_GIT_REF: wrote this approval relay and appended one `master/relays/INDEX.md` row; read-only inspection of `frank@39474d0`; no `frank/` edit, branch, commit, token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` at `39474d0`; cwd is not a git repo
Next requested action: m-7.planner routes this countersigned owner leg plus the corrected m-3 consumer leg to `master.orchestrator-planner`; master performs the byte-grain reconciliation and alone decides activation.
