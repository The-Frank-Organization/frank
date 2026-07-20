## DESIGN-REVIEW - m-1 exact-byte r2 review of the Step-3 secret-boundary and seat-identity contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m1-review-r2
PARENT_DISPATCH_ID: step3-mvp-design-m1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining findings reconcile already-ratified authority paths and current consumer contracts; no product fork is requested
GRILL_REQUIRED: no - stage-1 pair review consumes the ratified grill decisions; the remaining pressure questions are answered by current owner and consumer bytes
DESIGN_DOC_ID: step3-mvp-design-m1-secret-boundary-seat-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m1/DESIGN-planner-20260716-051843.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-052634.md
SUBJECT: must-revise - correct provider and broker authorization examples, kill old broker tokens on restart, and include typed Describe in the fenced capability surface

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner - I reviewed the revised contract bytes at SHA-256 `f3994fdcc20e44906b391caeb4860ef1d9afc35d2c60946f6ece47b3257a0192`, the addressed r1 fold relay, the ratified MVP amendment, and the current m-3, m-7, and m-10 consumer bytes.

The revision closes r1 F2 by naming the broker-private 0600 sink and closes the central r1 F3 gap with the five-event matrix and concurrent re-mint/replacement fixture. It also now separates references from USE capabilities. Three narrow inconsistencies remain inside that split. They are implementation-significant and already contradict the current consumer contracts, so pair approval remains blocked.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Findings

### MR1 - Section 1.4a binds independent authorization to the wrong mechanisms

Section 1.4a correctly says possession of a credential reference is not authorization, then gives the F59 one-shot ticket as the independent authorization for provider attempts and the F64 epoch gate as the independent authorization for broker channel use (`2026-07-16-step3-mvp-secret-boundary-seat-identity.md:32-35`). Neither example authorizes resolution of the referenced credential.

F59 authorizes one app-side tool call, bound to tool name and arguments; it is not the provider-request authorization. For S-A, the m-3 policy gate authorizes the frozen provider request at m-8, after which the m-8 secret resolver may attach the provider credential (`STEP-3-MVP-AMENDMENT.md:19,61`; `2026-07-16-step3-mvp-egress-e0-e3.md:30-31,56-58`; m-10 contract Section F59). For S-B, the broker resolves its fixed 0600 provisioning sink during startup or re-auth; F64 subsequently fences worker use of the bound seat channel. Calling either mechanism the credential-reference authorization conflates separate trust decisions.

Required revision: keep the generic non-authorizing-reference rule, but state the real paths separately. S-A resolution/attachment occurs only after m-3-owned provider-request policy authorizes the frozen core at m-8. S-B resolution occurs only inside the broker's operator-provisioned startup/re-auth path; F64 gates worker channel use after binding and does not authorize credential resolution. Preserve that neither reference is a bearer credential or a credential-derived verifier.

### MR2 - The broker-restart row lets an old connection-scoped capability survive

The new matrix says that on broker restart the capability may "survive iff the epoch is unchanged" and be re-established over new IPC (`2026-07-16-step3-mvp-secret-boundary-seat-identity.md:72-82`). Section 1.4's negative and the revised m-7 contract require the opposite token disposition: old connection-scoped capability material dies on broker restart or connection replacement. An unchanged `turn_epoch` can preserve the logical delegation, but not the old token.

Required revision: state that broker restart does not advance or reset `turn_epoch`, but every old connection-scoped capability/token becomes unusable. If the same current worker generation remains eligible, the broker may reattach it at the unchanged epoch only through the authenticated restart protocol and issue a fresh connection-scoped capability. Add or retain a negative proving replay of the old token fails after broker restart.

### MR3 - The USE-capability surface omits the typed Describe path required by the live transport

Section 1.4b limits the USE capability to exactly three broker verbs plus push receipt/rediscovery (`2026-07-16-step3-mvp-secret-boundary-seat-identity.md:34`). The revised m-7 contract correctly gates typed `Describe` metadata on the same authenticated, current-epoch worker channel. The live native/MCP path also depends on `DescribeTools` (`frank/internal/channel/server.go:313-320,481-490`; `frank/cmd/frank-mcp/mcp.go:216`). Leaving it outside the capability statement strands metadata discovery or creates an unfenced fourth path.

Required revision: define the capability authority surface as exactly the three canonical relay verbs plus typed transport metadata `Describe` and push/rediscovery, all behind the same current-epoch fence. State that `Describe` is not a fourth relay verb, is not one of the fixed eight dispatch tool names, creates no principal or identity consequence, and carries no relay acceptance authority. Add a stale-epoch `Describe` refusal negative consistent with the m-7 contract.

## Accepted portions

- **r1 F2 is closed.** The contract now identifies one broker-private permission-checked 0600 file as the only app-side persistent S-B sink and carries it through census, overwrite, backup, dump, log, and deletion obligations (`:13-16`, `:57-60`).
- **The reference/capability type split is structurally correct.** A reference is non-authorizing and verifier-free; a current USE capability is intentionally authorizing, secret-free, identity-inert, and connection/epoch constrained (`:32-35`, `:62-67`). MR1 and MR3 correct the named authority paths and complete the surface; they do not reopen the split.
- **The two-counter matrix and overlap proof close the main r1 F3 requirement.** Worker replacement, conductor restart, re-mint, and concurrent replacement/re-mint preserve independent counter authority and pin accepted in-flight ownership (`:72-82`, `:94`). MR2 corrects only the broker-restart token wording.
- **F57 honesty and identity accountability remain correct.** The same-user residual is explicit; logical identity remains the committed mint-pivot chain; worker generation is app-side bookkeeping, not a conductor principal.
- **The route-back result remains clean.** No new conductor verb, record member, registry entry, or store change is required. The three fixes reconcile m-1 wording to already-current m-3/m-7/m-10 interfaces.
- **No operator decision is needed.** These are contract-consistency fixes under ratified F57/F59/F60/F64, not a product or authority fork.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Replaces the F59/F64 credential-reference examples with the actual S-A provider-policy authorization and S-B broker provisioning/re-auth path.
2. States that every old broker connection token dies on restart, while an eligible current worker may be reattached at unchanged epoch only with fresh capability material.
3. Includes typed `Describe` in the current-epoch capability fence without classifying it as a relay verb or dispatch tool.
4. Preserves the accepted sink, reference/capability split, two-counter matrix, overlap fixture, F57 residual, channel-stamped FROM, committed-pivot generation authority, and no-conductor/store-change result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. Consumer confirmations remain held until m-1 pair approval.

## Verification

- `shasum -a 256 master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` -> `f3994fdcc20e44906b391caeb4860ef1d9afc35d2c60946f6ece47b3257a0192`.
- Read the exact addressed relay `master/relays/step3-mvp-design-m1/DESIGN-planner-20260716-051843.md`, its r1 review parent, and the current m-1 contract; routing, lineage, and `DESIGN_DOC_ID` match.
- Checked ratified `master/STEP-3-MVP-AMENDMENT.md:19,61,92-93,108-112` for provider request authorization, F59 tool-call scope, and F64 worker-channel fencing.
- Checked current m-3 provider egress bytes `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md:30-31,56-58` and current m-10 F59 tool-authorization contract `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:121-150`.
- Checked current m-7 consumer bytes at SHA-256 `8cf86753eaa711715ab3f063fda7d76efd4910cd8184eb066ff1588654ce39cc`, including locator/capability split, restart invalidation, fenced typed `Describe`, and stale-epoch negatives.
- Checked the live typed metadata path at `frank@502e06c`: `internal/channel/server.go:313-320,481-490` and `cmd/frank-mcp/mcp.go:216`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file and lane-root relay-lint.py exit 0
Next requested action: m-1.planner revises only MR1-MR3, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; consumer confirmation remains held.
