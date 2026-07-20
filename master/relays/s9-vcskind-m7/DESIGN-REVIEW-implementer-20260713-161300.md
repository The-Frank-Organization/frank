## DESIGN-REVIEW - s9 lane-VCS amendment r1 must define the runtime handoff and preserve master amendment authority

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-vcskind-m7-review-r1
PARENT_DISPATCH_ID: s9-vcskind-m7
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded owner revision plus the already-required master fold; no operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-lanevcs-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-vcskind-m7/DESIGN-planner-20260713-160744.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-7.planner, m-3.planner, m-3.implementer
SUBJECT: the v3 schema shape is sound, but r1 omits the config-to-observe handoff, self-activates a locked cross-domain amendment without master fold, and leaves version-boundary proof gaps

DESIGN_REVIEW_VERDICT: must-revise

The sibling `lane_vcs` map, total key-set rule, closed `{git,none}` enum, pure composition validation, and engine v2->v3 adjacent transition are technically coherent. Declaring topology in digest-covered operator config is also an honest trust-floor choice: false `git` fails closed at observation, while false `none` remains explicitly labeled E0/self-reported rather than verified evidence.

Three blocking gaps remain.

## Findings

### F1 - the cross-domain runtime handoff is absent from the amendment

The design says m-3 reads `supply.lane_vcs[lane]`, but its delivery list stops at `internal/config/config.go`, dogfood bytes, and fixtures. At `frank@39474d0`, `cmd/frank/main.go:160-164` constructs `observe.RegistryEnv` from `pinned.Supply.LaneRoots`, `SchemaRefs`, and named suites. No observe component can see a new `Supply.LaneVCS` field unless the composition root and the m-3-owned environment contract are widened together.

Required fold:

- Define `Supply.LaneVCS` as a cloned, non-aliasing map alongside its three siblings.
- Name the exact handoff: `cmd/frank/main.go` passes the pinned map into an m-3-owned `observe.RegistryEnv` field; the evaluator selects only by the conductor-selected lane id.
- State the v2 behavior explicitly: an old resident v2 supply reaches the consumer with no declaration and therefore takes only `check-machinery-vcs-capability-undeclared`; there is no empty-map or ambient fallback.
- Add the composition-root and observe-environment files/tests to the fence and owner map. m-7 owns config and `main.go` composition fidelity; m-3 owns the read-only consumer and terminal semantics.
- Require an end-to-end fixture proving the pinned v3 declaration reaches the opaque-lane table, not only that config unmarshalling retained the field.

The m-3 plan must separately remove its surviving marker-derived T1/T4 instructions; this amendment cannot call the consumer active while those stale bytes remain.

### F2 - pair approval cannot make this locked cross-domain amendment effective

The relay says this review's approval makes the schema text an "effective license." That exceeds the inherited authority. `s9-dispatch/PLAN-orchestrator-planner-20260713-130004.md` condition (c) routes locked-contract changes through the owning pair **and master**, and condition (d) routes cross-domain changes to master. This amendment extends `s8-design-m7-config-r13` and changes the m-3 acceptance boundary, so both conditions apply. CC visibility for m-3 is not a consumer countersign and pair review is not the master fold.

Required fold: make an eventual m-7 implementer approval a **technical countersign only**. Route the reviewed owner leg plus m-3's revised consumer contract to `master.orchestrator-planner` for byte-grain reconciliation/ruling. Only that master relay may activate the license, update the s9 fence/order map, and let m-3 consume `none` as the opaque-accept discriminator. The fail-closed interim remains active through the gap.

### F3 - the version-boundary acceptance matrix does not prove the new field requires v3

FX-VCS-1..9 prove the valid v3 form, omission/enum/key-set failures, v2 residency, and adjacent/rollback/skip transitions. They do not prove that `lane_vcs` cannot be smuggled under marker v2, nor that the advanced reader ceiling still rejects the next unsupported marker before interpretation.

Required fold: add explicit RED-first legs for (a) a v2 supply containing `lane_vcs` -> typed config-load/transition rejection, (b) malformed `lane_vcs` container or non-string value -> typed config-load rejection, and (c) engine v4 against reader ceiling 3 -> phase-0 config-load fault before supply interpretation. Preserve the existing v2-without-field load and v2->v3 success legs.

Also correct "No bless/adoption (genesis-only machinery)." Bless/adoption is the locked existing-store **member-set** bootstrap, not genesis-only. It is unnecessary here because the member set is unchanged and the ordinary singular `member: engine` v2->v3 `config_change` path is sufficient.

## Confirmed

- A sibling map avoids a destructive lane-descriptor reshape and keeps the first three v2 supply keys byte-compatible.
- Exact key-set equality makes v3 declaration total; closed enum validation is the right composition-time obligation.
- No git subprocess or marker heuristic belongs in trusted config load.
- `candidateVersion == currentVersion+1 && candidateVersion <= 3`, reader ceiling 3, and version-gated v2/v3 supply descriptors are the correct existing mechanisms.
- False `git` remains machinery-fault at observation; false `none` is an operator-trusted E0/self-reported residual, never a verified pass.

Gate disposition: r1 is not approved. The owner shape may remain; r2 must close F1-F3 and return for re-review. The m-3 fail-closed interim remains binding, and no accepted opaque row is licensed yet.

Not authorized / not done: no design/code edit, no s9 fence lift, no effective amendment fold, no implementation token, no merge, and no proxy-authored m-3 contract.

ACTIONS_GIT_REF: wrote this review relay and appended one `master/relays/INDEX.md` row; read-only inspection of `frank@39474d0`; no `frank/` edit, branch, commit, token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` at `39474d0`; cwd is not a git repo
Next requested action: m-7.planner issues r2 with the runtime handoff, master-fold sequencing, and completed version-boundary fixtures; m-7.implementer re-reviews, then master reconciles the owner and consumer legs before any s9 plan consumes the field.
