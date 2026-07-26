## DESIGN-REVIEW - m-1 adversarial review of the stage-6 lane-2 env/redaction additive delta

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m1-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings can be resolved within the released TCB scope; any decision to add a manifest environment surface must instead route to master
GRILL_REQUIRED: no - the release fixes the product semantics; the pressure questions were answered from current owner and producer bytes with no operator-owned fork remaining
DESIGN_DOC_ID: step3-relock-dag-m1
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-relock-dag-m1/DESIGN-planner-20260721-233847.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-relock-dag-m1/DESIGN-REVIEW-implementer-20260721-234730.md
SUBJECT: must-revise - make the open-content redaction claim honest, close authority-material and at-rest routes, and define an owner-real sanitized environment

DESIGN_REVIEW_VERDICT: must-revise

m-1.planner - I reviewed the final delta bytes at SHA-256 `41385dfcdb28d7f49813bfa4881e9b3791d7973facfaa40bd36b674a99fd09c8`, the exact released rev2 dispatch at `9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`, amendment rev12 at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, frozen m-1 at `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`, and the current m-9/m-10 lane-2 producer bytes as non-attached evidence.

The release and additive-document lineage are valid, the frozen base is unchanged, and the producer-attaching halves are correctly parked. Five TCB statements still overclaim or lack an owner-real enforcement path. They block pair approval and any later redaction co-sign.

This review grants no stage-6 lock, producer confirmation, join co-sign, PLAN, T4/code token, credential, provider call, release binding, live E3, `frank/` edit, merge, or deploy.

## Findings

### F1 - A record-class gate cannot prove arbitrary content strings contain no secret bytes

Section 2 says every loggable class is already secret-free upstream and therefore "no frame class carrying secret-bearing material exists on the log path" (`2026-07-22-stage6-lane2-env-redaction.md:31-34`). That premise is broader than the frozen contract. Frozen m-1 excludes S-A/S-B from model context, local-tool arguments, bash env/files/argv, inherited FDs, and logs; it does not prove arbitrary local-tool output secret-free, and it explicitly retains same-UID bash inspection as an unsandboxed residual (`2026-07-16-step3-mvp-secret-boundary-seat-identity.md:20-23`).

The current m-9 producer union makes the mismatch concrete: `input_item.content`, `tool_call.args`, `tool_result.content`, and `provider_output.content` are open content values (`2026-07-22-relock-lane2-m9-delta.md:27-44`). A closed `kind` enum prevents a credential object or replay envelope from being routed directly, but it cannot make the bytes inside an allowed string type secret-free. A confused or deliberately same-UID bash/read path can place arbitrary bytes in a tool result; content scanning cannot repair that reliably.

Required revision: split the claim. Pin the strong structural claim to typed object routes: secret-holder outputs, credential objects, replay envelopes, and authority objects have no writer input variant and cannot be serialized as such. For open user/provider/tool content, state the honest confusion-not-malice ceiling: ordinary governed paths inject no S-A/S-B, accidental-disclosure absence is sentinel-tested, but operator/user-supplied content and deliberate same-UID extraction remain outside a byte-absence guarantee until H-12. If the ratified "any S-A/S-B secret bytes" row is intended as an absolute guarantee even against open-content acquisition, route the spec mismatch to master because the current threat ceiling and record union cannot realize it.

### F2 - The closed NOT-list omits the F59 ticket and does not define the authority-object family completely

Section 2.2 excludes USE/control capability tokens but not F59 one-shot tickets (`:35-40`). F59 tickets are non-secret under frozen m-1, but they are authorizing and the ratified worker invariant says the durable log carries "no credential, no ticket, no epoch grant" (`2026-07-22-relock-lane2-m9-delta.md:146-150`). Non-secret is not equivalent to safe for durable content persistence.

Required revision: exclude the complete authority-object family by type, explicitly including F59 one-shot tickets, m-7 worker USE capabilities, broker control tokens, connector/control-channel tokens or handles, and any future object tagged as authority-bearing. Keep opaque credential references separately classified as non-secret/non-authorizing but non-loggable. Define the writer gate over source object types so a newly added authority type fails closed rather than silently becoming loggable. Preserve `turn_epoch`/generation identifiers as attribution facts, not grants.

### F3 - The at-rest review states modes and location but not a confusion-resistant create/open boundary

Section 2.3 requires a 0700 directory and 0600 file and says workspace-relative paths cannot name the log (`:41-46`). A path outside a workspace plus mode prose does not establish where an opened descriptor points. The current m-9 producer receives the path from `turn_open` and verifies in-file identity, but its present bytes do not pin no-follow creation/open, regular-file/owner/mode/link checks, or descriptor identity across open/read/append (`2026-07-22-relock-lane2-m9-delta.md:23-25`). A symlink/substitution path can redirect the durable content into a workspace or another file while all logical record checks still pass.

Required revision: make the at-rest condition descriptor-grain. Require a verified owner/mode private parent directory; create new files exclusively without following links; reopen with no-follow/close-on-exec; verify regular-file, owner, exact mode, link count, and opened-descriptor identity before use; reject replacement races; keep rotation/segments in the verified directory with directory durability. State that workspace/private-runtime separation is evaluated on the opened object and resolved ancestry, not lexical path text. Carry matching RED legs to the later m-9 review. This remains confusion resistance; the accepted same-UID bypass residual stays explicit.

### F4 - The environment rule invents an ownerless manifest-variable class and leaves baseline paths unsafe or non-enumerated

Section 1 says there are "exactly three allow classes," but class (iii) is "nothing else," so there are two positive classes plus a deny-all remainder (`:10-20`). More importantly, class (ii) depends on run-scoped workspace variables "named by the FROZEN MANIFEST." The current frozen m-10 manifest and its lane-2 producer delta have no environment-name/value member; the latter explicitly leaves `env_digest` invocation-resolved at m-9 and authors no environment source. This delta therefore assigns a new producer surface to m-10 without a released owner contract, while calling the result "no behavior delta."

The class-(i) list is also not fixed as claimed: `LC_*` is a wildcard, and `HOME`, `TMPDIR`, and `PATH` are accepted without pinning whether they expose the real user home, ambient credential discovery paths, the private runtime directory, or unapproved executable directories. That can reintroduce confused ambient credential discovery even when credential-valued variables are absent.

Required revision: define exactly two positive classes plus deny-all. Either make class (ii) empty for this MVP or route a specific frozen-manifest schema/carrier addition through master to m-10 before relying on it. Enumerate every locale variable by exact name; no wildcard. For each baseline variable pin its source and value constraints. In particular, use an empty/dedicated tool HOME and TMPDIR, a fixed approved toolchain PATH, and values that do not name the credential sinks, control/private runtime directory, or ambient user credential/config roots. If a real HOME or user-derived lookup remains necessary, state the resulting default-credential-discovery residual rather than calling ambient credentials structurally absent.

### F5 - The whole-set digest is still a confirmation oracle when one value is unknown

Section 1 says a whole-set `env_digest` "offers no per-variable confirmation oracle" (`:22-25`). A deterministic digest of the full set is a confirmation oracle whenever an observer knows all other members and can guess one candidate value. Secret-free input is the condition that makes the digest safe; whole-set hashing itself does not remove the oracle.

Required revision: withdraw the no-oracle-by-shape claim. State that any secret-bearing class-(ii) value makes the digest a secret-derived confirmation artifact and is a noncompliant operator provisioning error that the mechanism cannot detect by content. Pin or delegate one exact canonical preimage schema for the complete presented environment (object versus ordered name/value array, exact string encoding and duplicate-name rejection) so m-9 and the E3 observer cannot hash different logical sets. Keep the RED leg proving the pre-sanitization parent environment is never the input.

## Accepted portions

- **Release and F73 discipline are correct.** The addressed rev2 release is active at the reproduced hash; the frozen m-1 base remains byte-identical; the delta is additive rather than an in-place edit.
- **Construct-from-empty is the correct direction.** Inherit-then-filter would fail open. F4 asks for an owner-real, exact positive set, not a return to deny-listing.
- **The descriptor path direction is sound.** Digest the actually presented child environment, not the parent; `shell_interpreter_ref` is non-secret when its resolved object is outside secret/control roots; workspace-relative `canonical_resource`/`cwd` is acceptable when opened-object containment enforces the separation.
- **K6 exclusion, per-run retention, no backup/telemetry copy, and the frozen sentinel evidence ceiling are correct.** The revision must preserve that sentinel tests prove accidental-disclosure absence only.
- **Producer-first parking is correct.** Section 4 states judgment criteria and explicitly defers m-9/m-10 confirmation and the D join co-sign until exact pair-approved producer bytes exist; it does not act on those working bytes.
- **No operator decision is required** if the revision uses a static MVP environment and narrows the redaction claim to the already-ratified confusion-not-malice ceiling. A new manifest environment surface or an absolute open-content byte-absence guarantee must route to master.

## Revision bar

Return fresh delta bytes and a fresh hash that:

1. Separates structural object-route exclusion from the honest residual for open user/provider/tool content.
2. Excludes the complete authority-object family, including F59 tickets, with fail-closed type evolution.
3. Defines descriptor-grain secure create/open/containment requirements for the at-rest log and matching negative fixtures.
4. Replaces the ownerless/wildcard environment rules with an exact positive set whose source and path constraints are implementable on current released seams, or routes any new manifest field to master.
5. Corrects the digest-oracle claim and binds one canonical complete-environment preimage recipe.
6. Preserves the frozen base, K6 exclusion, F57-narrow ceiling, construct-from-empty direction, producer-first parking, and separate consumer/join gates.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte review. Do not attach a producer confirmation or D join co-sign before the relevant producer bytes are pair-approved and reviewed separately.

## Verification

- Reproduced SHA-256: released rev2 `9a4ee380da9afacbbafb74fd854a97c2cbb814b057bc0619e7e31f8b1815b3a3`; delta `41385dfcdb28d7f49813bfa4881e9b3791d7973facfaa40bd36b674a99fd09c8`; frozen m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Read and exact-file linted the addressed planner relay and the released rev2 dispatch; `TO`, `IN_REPLY_TO`, `DESIGN_DOC_ID`, and author-now/park-later scope match.
- Checked amendment Sections 5-C/5-D and the frozen m-1 non-injection, same-UID residual, authority-object, and sentinel ceilings.
- Checked the current m-9 record union, at-rest path/file-identity rules, no-durable-ticket invariant, and env-digest producer half as non-attached evidence; checked current m-10 manifest/ticket producer bytes for the absence of a manifest environment-variable source.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source contract, historical relay, frozen byte, `frank/` source, branch, commit, lock, producer confirmation, join co-sign, PLAN, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `c78da38`.
RELAY_LINT: OK - exact-file and lane-root relay-lint.py exit 0
Next requested action: m-1.planner revises only F1-F5 or routes the two named seam/spec choices to master, re-hashes the delta, and sends fresh uniquely-parented bytes for pair review; producer confirmations and the D join co-sign remain parked.
