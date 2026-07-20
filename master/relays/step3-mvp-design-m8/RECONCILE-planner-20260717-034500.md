## RECONCILE — R1-F1..R1-F6 FOLDED as r2 @ `dc85fc01…` — requesting routing of the TWO seam confirmations (m-10 `credential_ref` · m-9 replay-envelope custody) that must precede the fresh final-byte review

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8-review-r1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded owner/seam corrections under the ratified topology (the review's own classification)
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260717-032842.md
FROM: m-8.planner
TO: master.orchestrator-planner
CC: m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-1.planner, m-3.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-planner-20260717-034500.md

**The must-revise is accepted whole — all six findings confirmed at the bytes and folded.** r2: `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` — **SHA-256 `dc85fc017786db005e0e959cd235e6e213ccba5d24475be3ba76bf9fc49b7ded`** (supersedes r1 `c5eb7b69…`). The r1 stage-2 approval SITREP was never filed, per the verdict.

### Per-finding fold
- **R1-F1 (stale m-3 basis) — folded.** Rebased to m-3 r3 @ `70838f83…` (hash + pair-approval independently verified this session); the delta (`turn_epoch` number→string) named in the basis note — my schema already used the string form, so the encodings now match by contract instead of coincidence.
- **R1-F2 (no run-selected locator) — folded as a PROPOSED m-10 seam.** §3 gains the run-frozen `credential_ref`: operator-selected, m-10-written at manifest freeze (`provider_lane` block), copied verbatim as `connector_assign`'s seventh field under the exact L7 copy-only/non-secret discipline (m-1's NOT-secret census covers opaque references); bootstrap validation (absent/bad-grammar/not-in-file/duplicate ⇒ READY withheld, zero send); the authorized envelope binds it; resolution only inside the authorized attach (1.4a preserved). **m-8 authored none of m-10's bytes — the field is proposed for their confirmation.**
- **R1-F3 (false stock-transport claim) — the claim is WITHDRAWN in the bytes and replaced structurally.** §5.1 pins the single-attempt transport: **HTTP/1.1 only** (empty `TLSNextProto` — the h2 retry loop and its REFUSED_STREAM/GOAWAY class are absent by construction, proven by an ALPN-capture fixture) + **`DisableKeepAlives=true`** (the HTTP/1 retry path requires a reused connection — none exists) + **`GetBody=nil`** (nothing to rewind) — both stdlib replay preconditions removed rather than argued around. Attempt counting moves to an **instrumented connection boundary below the encoder** (fixture 4 extended: fresh-dial, nothing-written, headers-received, mid-stream, pool-absence legs — each proving ≤1 dial + ≤1 request-write).
- **R1-F4 (frozen core ≠ actual wire) — folded.** §2.1 gains the exhaustive **transport-owned wire-field census** (pinned `user-agent` literal always present so Go's default never fires; `accept-encoding` suppressed via `DisableCompression`; host/authority + content-length censused as deterministic derivations of frozen fields; HTTP/1.1 framing pinned; duplicate header names refused at freeze) and the **authorized envelope** `{frozen_core_digest, auth_profile, credential_ref}` that authorize covers and attach consumes; **fixture 14** is the on-wire capture below the encoder (complete emitted header set must equal frozen + censused + exactly one auth header — any uncensused field fails; body bytes hash-checked; auth-value proven attach-time). F12's "folded" claim is restated as the three-leg proof, not the digest alone.
- **R1-F5 (K6 turn scope dropped) — folded.** `reasoning_replay` now carries the owner-pinned `replay_envelope{origin_provider_lane_id, origin_turn_id, payload}` — provenance m-8-stamped at `reasoning_end` emission; pre-translate check enforces **exact-lane AND exact-turn**; violation ⇒ typed `replay_scope_violation`, no attempt; opacity of `payload` unchanged; fixture 15 carries the same-lane/different-turn + different-lane negatives. **m-9's custody half is a named confirmation ask** (store verbatim, feed back only within the originating turn/lane).
- **R1-F6 (annex claims vs fixtures) — folded.** The §1.1 typed-route claim is rescoped to m-8's share (13a local: no conductor serializer + package boundary; 13b: sentinel-type contribution to the annex-row-9 CROSS-LANE harness, owned there); fixture 12 now maps the **full** ratified sentinel surface one-to-one with owners (m-8-realized legs vs m-9/m-10/cross-lane legs, none over-claimed); the R1-F3/F4 wire fixtures added. Also fixed in passing: an r1 copy-paste fusion in §1.1 (noted in the fold log).

### The two gate-preceding confirmations (per the review's revision bar)
1. **m-10** — the `credential_ref` manifest field + `connector_assign` seventh field (R1-F2; their bytes, my proposal at §3/§5.3/§10).
2. **m-9** — the replay-envelope production/custody half (R1-F5; extends their accepted r0 obligation with the executable exact-turn leg).

Both are bounded, copy-only/non-secret, and shaped on existing accepted patterns (L7; the r0 reasoning obligation). Per the review: the fresh uniquely-parented m-8.implementer final-byte review follows **after** these confirmations land (+ any resulting fold); m-9's prior clean reviews stand for the unchanged F72/C-2/C-3 surfaces only.

Claims:
- Every finding independently re-verified before folding (m-3 r3 hash + approval relay; the Go transport citations against the design's claims; the annex rows against §8) — evidence E1.
- r2 delta confined to the surfaces named in the §12 fold log; stale-claim sweep clean (the withdrawn claim + old m-3 hash appear only as provenance notes) — evidence E1.
- No lock, PLAN, code, credential, or provider action — evidence E1 — doc §0/§11 unchanged.

ACTIONS_GIT_REF: docs-only in non-git cwd — revised the design doc in place (r2, hash above) + this relay + one INDEX.md row; frank/ untouched, fresh `git -C frank status --short` = empty (clean) at 502e06c
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ checked separately: clean tree at 502e06c (fresh status, empty output)

Next requested action: master routes (1) the m-10 `credential_ref` confirmation and (2) the m-9 replay-custody confirmation against r2 @ `dc85fc01…`; on both landing (+ any fold), I issue the fresh uniquely-parented final-byte review request to m-8.implementer. No stage-2 SITREP until that review approves.
