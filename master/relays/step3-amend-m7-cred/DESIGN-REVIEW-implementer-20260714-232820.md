## DESIGN-REVIEW - credential contract r2 closes F1-F6 but exposes four new mechanism-grain gaps

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m7-cred-review-r2
PARENT_DISPATCH_ID: step3-amend-m7-cred
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the operator provisioning and first-E3 gates remain routed and are not decided by this review
GRILL_REQUIRED: yes
DESIGN_DOC_ID: step3-amend-m7-cred-credential-contract
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: step3-amend-m7-cred-grill-r2
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-amend-m7-cred/DESIGN-planner-20260714-232405.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-8.planner, m-3.planner
SUBJECT: r2 must revise - F1-F6 are substantively closed, but the secrets census has no descriptor row or activation semantics, the parent descriptor can still escape/block, Attach accepts a free-form ID without a pinned header profile, and engine-v5 closed-shape validation is deferred to PLAN

DESIGN_REVIEW_VERDICT: must-revise

R2 closes the six original findings at their requested grain. In particular, post-recovery all-or-nothing publication is executable against the current startup path; request-selected credential refs are gone; final-wire freezing, redirect-off, and per-attempt authorization are explicit; the vault claim is narrowed honestly; and the visibility matrix resolves the r1 ambiguity. Four second-order defects in those new mechanisms remain.

## Findings

### F7 - `secrets` is claimed as a canonical path family but the locked descriptor remains a closed 12-row census

R2 says `<root>/secrets/` joins `canonical_path_families` (`2026-07-15-step3-credential-contract.md:24,79,159`) without defining its row. The amended lock still inherits the r13 descriptor whose rows are closed `{id, relative, forbidden, directory}`, sorted by ID, and pinned to exactly twelve rows (`2026-07-11-s8-config-host.md:65-90`). The live oracle and catalog invariant likewise enumerate twelve and have no `secrets` home (`frank/test/invariants/path_hygiene_test.go:370-451`; `catalog_test.go:121-122`). PLAN would have to invent the row, cardinality, live-root creation rule, and config/census activation order.

Required fold: choose and pin one coherent mechanism. If `secrets` is a canonical root home, add the exact sorted row (expected shape: `{id:"secrets", relative:"secrets", forbidden:"/secrets/", directory:true}`), change the expected census to thirteen, and define whether the 0700 directory is always present even with zero bindings. Specify how an existing twelve-row store reaches the engine-v5/thirteen-row state without an invariant-red intermediate, including whether catalog+engine config changes require an atomic transition or a staged order with an explicit compatibility state. If the directory is conditional, it cannot inhabit the current exact-live-equality family unchanged; define a separate forbidden-only census field/law instead. Preserve r2's zero-binding and no-genesis-break claims under the selected path.

### F8 - Leaf-only `O_NOFOLLOW` does not bind the material to `<root>/secrets`, and a FIFO can block before rejection

Section 3 says to open the "already-flock'd root's secrets directory" and then applies `O_NOFOLLOW` only to the credential leaf (`design:42`). The current root lock actually flocks `<root>/conductor.lock`, not a root-directory descriptor (`frank/internal/store/lock.go:43-53`). Unless every component is descriptor-opened no-follow, `secrets` itself may be a symlink to an owner-matched 0700 directory outside the root and still pass all stated leaf checks. Also, `openat(..., O_RDONLY)` on a FIFO can block indefinitely before `fstat` reaches the promised non-regular-file reject. The landed rooted walker already demonstrates the needed component-wise `O_DIRECTORY|O_NOFOLLOW`, `O_NONBLOCK`, and `CLOEXEC` posture (`frank/internal/observe/fs_worker.go:209-274,341-350`).

Required fold: open a root directory FD, then open `secrets` relative to it with `O_DIRECTORY|O_NOFOLLOW|O_CLOEXEC`; `fstat` that directory before opening a leaf. Open the leaf with `O_NOFOLLOW|O_NONBLOCK|O_CLOEXEC`, then require regular-file identity. Reject `credential_id` values `.` and `..` explicitly rather than claiming the inherited charset excludes traversal when it admits both. Enforce the 4096-byte ceiling with a bounded descriptor read of at most 4097 bytes, not only a pre-read `st_size` check, so growth after `fstat` cannot make the read unbounded. Add parent-symlink, dot-ID, FIFO-no-writer/nonblocking, and grow-after-stat legs to FX-CRED-2. The exact parent mode/owner values may still await m-1; descriptor confinement and nonblocking behavior are m-7-local.

### F9 - The attach API reintroduces free-form credential choice inside the host and the auth profiles have no exact wire mapping

R2 correctly removes credential choice from adapters and requests, but the sole capability is `Vault.Attach(req, credentialID)` (`design:75`). A free string lets the call site accidentally substitute any valid credential after governed derivation; it does not mechanically bind attachment to the unique result of section 4a. The enum `{bearer, x_api_key}` is named (`design:20,66`), yet neither exact case-insensitive reserved header set nor exact header/value mapping is pinned. "Headers-sans-auth" therefore has no executable definition: PLAN must decide whether a preexisting `Authorization`/`X-Api-Key` is rejected, overwritten, or coexists, and what bytes each profile attaches.

Required fold: make derivation return an opaque selected-binding capability and make attachment consume that capability, or combine derive+attach in one vault/sender operation; do not cross the final boundary with a caller-selected credential-ID string. Pin v1 profile bytes exactly (header name, value prefix/shape, and whether the key is replaceable), reserve auth header names case-insensitively, and reject any pre-freeze adapter/request occurrence or duplicate rather than silently overwrite it. State whether host-level `Request.Host` must be empty/equal to the canonical URL authority. Add wrong-handle/ID-substitution, mixed-case reserved-header, duplicate-header, and cross-profile-header negatives.

### F10 - The engine-v5 trusted-config shape and composition rules are still deferred to build-time invention

The current transition gate validates every candidate against a per-version closed schema before accepting a successor (`frank/internal/config/config.go:263-284`; config-host lock `2026-07-11-s8-config-host.md:21`). R2 lists row fields, but section 12 leaves "the exact v5 descriptor bytes" to a build-time owner spec (`design:170`) without pinning the representation and composition rules needed to write that descriptor. Open questions include whether `provider_bindings` is required-empty or optional at v5, exact row key closure and array element kinds, unique `credential_id`, empty/duplicate allowlists and lane IDs, absent versus empty `lane_ids`, row/list ordering, non-negative generation, and whether unknown provider/lane references fail composition or only later selection.

Required fold: add the canonical engine-v5 descriptor delta and a separate composition-validation table. Pin required/optional fields, exact node kinds and closed key sets; uniqueness and ordering semantics; all value constraints; and typed failure timing. Keep m-8-owned provider/lane factual validation explicitly provisional where its consumer packet is required, but do not leave m-7's own structural and uniqueness rules to PLAN. Extend the transition fixtures with unknown/extra/missing key, wrong node kind, duplicate credential ID, malformed empty-set cases, and lawful zero-binding v5.

## Confirmed

- R1/F1 is closed: publication is post-valid-Ready, pre-sender, complete-set atomic, with zero partial availability and no replay sends.
- R1/F2 is closed: adapter/runtime/request credential choice is removed; unique governed derivation and ambiguity refusal are the correct account boundary.
- R1/F3 is closed as to byte grammar, descriptor-identity intent, mode/owner handoff, and all-or-nothing loading; F8 is the remaining parent/open/read mechanics defect.
- R1/F4 is closed: canonical endpoint rules, final-wire freeze, redirect-off, and per-attempt authorization are pinned; F9 only closes the auth-header attachment bytes/capability.
- R1/F5-F6 are closed: claim ceiling, vault visibility, enumerated representations, post-attach capture ban, visibility matrix, canary encodings, and generation semantics are acceptable.
- The B14 consumer packet still blocks final approval/close, not this technical re-review. F7-F10 are owner-local and require no operator decision.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no credential provisioning/use, no E3 call, no lock, no merge, and no operator decision inferred.

ACTIONS_GIT_REF: wrote this r2 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`
Next requested action: m-7.planner folds F7-F10 into credential-contract r3 and returns it for technical re-review; final amendment approval remains held for the B14 consumer packet.
