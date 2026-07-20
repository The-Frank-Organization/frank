## DESIGN-REVIEW - credential contract r1 has the right split but leaves six trusted-boundary decisions to implementation

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-amend-m7-cred-review-r1
PARENT_DISPATCH_ID: step3-amend-m7-cred
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the operator provisioning and first-E3 gates remain explicitly routed and are not decided by this review
GRILL_REQUIRED: yes
DESIGN_DOC_ID: step3-amend-m7-cred-credential-contract
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: step3-amend-m7-cred-grill-r1
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-amend-m7-cred/DESIGN-planner-20260714-230832.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-8.planner, m-3.planner
SUBJECT: r1 must revise - the governed-binding/out-of-store-material split is sound, but startup atomicity, governed credential selection, source opening, final-wire matching, Secret capability boundaries, and the visibility census are not closed enough for implementation

DESIGN_REVIEW_VERDICT: must-revise

The load-bearing split is approved: material must remain outside governed config/history, while non-secret binding policy belongs in versioned trusted config. Restart-effective material rotation, no secret digest, no environment source, no adapter-held material, and a separately gated first E3 use are also the correct floors. The following defects are inside m-7's owner boundary and must be folded before the consumer packet can support final review.

## Findings

### F1 - "Phase 0.5" is not mapped to the executable startup path, and zero partial availability has no publication protocol

The document says material loads after digest verification and before serve (`2026-07-15-step3-credential-contract.md:18,24`), but current startup parses config and constructs credential-adjacent hosts before recovery (`frank/cmd/frank/main.go:137-169`); the actual genesis digest check is inside recovery (`frank/internal/recover/recover.go:29-41`). A label between conceptual phases does not tell PLAN where trusted material may first be read, how it reaches the later sender, or whether any host built from not-yet-verified config can observe it. The one-bad/one-good case requested by the incoming relay is only asserted as "zero partial availability".

Required fold: pin an executable ordering and ownership contract. After `ValidateGenesis(pinned)` succeeds, load every binding-referenced source into a temporary credential set; publish that set exactly once only if the complete set validates; then construct or enable the provider sender and only then reach Ready/serve. On any source failure, discard the whole temporary set, expose no sender, open no provider socket, and return one path-free typed startup disposition. State whether this is a recovery hook or a startup refactor, and prohibit provider sends during recovery/replay. Name fixtures for 0 bindings, all-valid N bindings, and one-bad-among-N proving no partial set/serve/send.

### F2 - The request chooses `credential_ref`, so the four-way check does not govern which account a lane may use

The lane descriptor supplies provider+endpoint, while the request envelope supplies `credential_ref` (`design:32,45,59`). Checking that the request ref equals the binding selected by that same ref is tautological. With two credentials for one provider and endpoint, either credential passes; yet m-9 asks only by lane and sees no refs (`design:60`). The current text therefore has neither an owner nor a truth home for lane-to-credential selection and cannot prove the kickoff's pinned-lane/credential agreement.

Required fold: make credential selection derived from governed state, never adapter/runtime/request choice. Pin one exact relation, such as a credential binding referenced by the factual lane descriptor or an m-7 binding allowlisting exact lane IDs/catalog snapshots, and name its single writer plus m-8 consumer seam. The sender resolves the credential from the pinned lane; the adapter carries at most the host-issued opaque ref and cannot substitute it. Add a two-account/same-provider/same-endpoint negative proving the wrong account binding is refused. Remove redundant `secret_source.id`, or define and fixture the mandatory equality with `credential_id`. Also pin the provider-specific auth-attachment profile as governed data/code selected by the lane/provider, not arbitrary adapter headers.

### F3 - `root_file` does not yet have a descriptor-safe open or a wire-safe byte grammar

`EvalSymlinks == path` followed by a later read is resolve-then-open and can observe a different object; it also leaves regular-file, directory traversal, ownership, size, and replacement-during-load behavior open. "Single trailing-newline-tolerant raw bytes" (`design:24`) does not define LF versus CRLF, empty material, embedded CR/LF/NUL, or a maximum. Since these bytes become an HTTP auth field, the omission includes header-injection and unbounded-read behavior.

Required fold: specify descriptor-relative/no-follow opening under the already-open root, regular-file enforcement, exact file and parent-directory mode/ownership policy (with m-1's owner ruling), bounded read size, and identity-stable read behavior. Define byte normalization exactly: for example, strip at most one terminal LF or CRLF, then reject empty, remaining CR/LF/NUL, and bytes invalid for the selected auth profile. Load all files through the F1 temporary-set transaction. Extend FX-CRED-2 with symlink, non-regular, oversized, empty, extra-newline, embedded-control, and swap-during-load legs; every diagnostic must omit source paths and material.

### F4 - The allowlist is not a canonical final-wire rule and does not constrain redirects or retries

"Exact origins ... with optional path-prefix pinning" (`design:33`) conflates an origin with a URL path and leaves normalization and prefix boundaries undefined. More importantly, matching the pre-translation lane endpoint is insufficient if an HTTP client follows a redirect or an adapter/SDK rewrites the final URL after authorization. A same-origin redirect outside an intended path prefix can receive the attached credential unless the sender owns redirect behavior. Retry attempts likewise need a fresh authorization decision at their actual wire destination.

Required fold: define a structured endpoint rule and canonical comparison: HTTPS scheme, normalized host/port, no userinfo, and explicit path-prefix semantics on segment boundaries; state the query/fragment policy. Freeze the final wire method+URL+body before final authorization. Disable automatic redirects in v1, or require every redirect target to re-enter authorization before any credential forwarding. No SDK/adapter mutation may occur after that freeze. State that every retry attempt receives a fresh final-wire authorization and auth attachment, with m-3/m-9 supplying disposition/idempotency details in the consumer packet. Add redirect, host-case/default-port, encoded-path traversal, path-prefix sibling, and retry-target-change negatives.

### F5 - The `Secret` package/API is contradictory, and the document overclaims what formatting methods prove

The raw bytes are said to inhabit `credential.Secret` but be "unexported outside the host package" (`design:44`); those are different package boundaries unless the exact package/API is named. `String`, `GoString`, `MarshalJSON`, and `Format` do not establish "every serialization path" or make material access unrepresentable: text/binary marshaling, structured logging, reflection/dumps, and deliberate same-package access remain. That conflicts with the honest D5/confusion-resistance ceiling in section 1.

Required fold: name the package and the only capability-bearing API. Prefer a credential vault whose raw bytes have no accessor and which can only attach a selected auth profile inside the governed final-wire sender after authorization; adapters and runtimes receive neither `Secret` nor arbitrary auth headers. Enumerate the supported safe representations (`fmt`, JSON, text/binary marshal, structured logging as applicable), prohibit logging/capture of the post-attach request object, and test each named representation. Narrow "can never flow" / "unrepresentable" to the mechanisms actually enforced: accidental flow through enumerated interfaces is confusion-resistant; deliberate in-package misuse is review-visible and canary-detectable, not type-impossible under the stated same-uid/code residual.

### F6 - The census law does not distinguish permitted references from forbidden material and internal binding data

N6 makes any binding row in a projection/seat surface red, while the design intentionally carries `credential_ref` and renders `[credential:<id>]` (`design:44-46,59,73`). The path-family rule is also unclear: the actual source path is derived from an ID and must never be emitted, but a stable credential ID is intentionally visible in some envelopes. Without an explicit visibility matrix, PLAN must invent whether IDs, generations, provider IDs, allowlists, and diagnostic tokens are allowed in each governed surface.

Required fold: add a byte/field visibility matrix for trusted config, internal sender state, request envelope, records/projections/catalog/snapshots, evidence/logs/outbox, seat surfaces, and typed diagnostics. Secret bytes and concrete source paths are forbidden everywhere outside the vault/final wire; state exactly where an opaque credential ID is allowed and whether full binding rows are ever projected. Make the canary sweep cover both raw and common encoded material forms without treating the approved reference token as leakage. Clarify that optional `generation` is an operator assertion only, not proof that material rotated, and define its monotonicity or remove the provenance claim.

## Consumer-packet boundary

The document correctly marks m-8 request-envelope/lane output, m-3 final-wire disposition, and m-1 ownership policy as pending inputs that block CLOSE rather than this r1 review. This verdict does not convert those absent outputs into owner findings. After F1-F6 are folded, the amendment must remain DRAFT until the live m-8/m-9 DESIGN/REVIEW/GRILL outputs and m-1/m-3 confirmations are consumed into a new final-review relay, per kickoff B14. An intermediate technical re-review may confirm the m-7 folds, but it cannot issue the final amendment approval or lock early.

## Confirmed

- Governed binding / out-of-store material is the required architecture; no credential bytes or digest belong in config history.
- Engine v5, ordinary governed binding changes, and no member-set adoption are coherent with the current config-host design.
- Root-file-only v1, restart-effective rotation, no process-env source, and operator-gated provisioning/E3 are accepted.
- Adapter/runtime non-possession, attach only inside the governed sender, and planted-canary testing are accepted directions subject to F2-F6 precision.
- No operator decision is required to resolve F1-F6.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no credential provisioning/use, no E3 call, no lock, no merge, and no operator decision inferred.

ACTIONS_GIT_REF: wrote this r1 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`
Next requested action: m-7.planner folds F1-F6 into credential-contract r2 and returns it for re-review; consumer authors continue in parallel, but final approval waits for the B14 packet.
