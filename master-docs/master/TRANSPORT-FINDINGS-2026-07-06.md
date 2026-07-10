# frank transport findings — the s5-dogfood ledger (2026-07-06)

**Status:** the primary deliverable of the s5-through-frank run (the first slice-team to run its governance ON frank),
**stopped at checkpoint by operator direction 2026-07-06** when finding F11 escalated from friction to a hard liveness
failure. This ledger is the spec seed for the **transport-fix cycle** (next build after the s5 checkpoint).
**Evidence base:** the archived live store (`~/frank-archives/frank-team-store-s5-dogfood-20260706` — 41 records /
34 intakes at archive; s5's running count: ~13 of its first 28 intakes rejected), the cited relay/intake ids (all
persist in the archive), s5's findings relays (`relay-334a61d4…` F1–F7; F9–F11 operator-relayed from its held report),
the operator seat's waiver-submission report, and the master seat's own transcript. Code cites @ `main 67ee23e`.
**Honesty ceiling:** all store evidence is transport/provenance-grade (the Step-1 claim); content claims are the
reporting seats' own.

## Why this exists

s4 proved the transport works for **one relay between two live seats**. s5 was deliberately run *through* frank to
learn what a **busy multi-seat store** does. Answer: the safety layer held everywhere; the **liveness and
single-interpretation layers failed**. A single-seat test could never have surfaced F11 — this is the finding the
dogfood was designed to buy.

## Root cause (one statement)

**One envelope, many judges, no reconciliation protocol.** The rendered form (m-2 registry render), the typed
validator (`ParseTyped`), the lineage gate (raw-header comma-splits), and the delivery projector (JSON `addressList`)
each interpret the same envelope **independently** — so the form can offer what the validator rejects (F2, F12, F13),
demand an encoding the gate can't read (F6/F7), and pin the seat to a parent/digest snapshot that **global concurrent
traffic has already moved** (F4/F5 → F11). Add client-supplied fields the engine already knows (PARENT), a validation
context with no lease/retry token, and rejected records polluting the anchor pool (F10), and the honest retry path
**livelocks** under exactly the load frank exists to carry. Advisory-render / authoritative-revalidate is **safe but
not live**.

## The findings

*Numbering continues s5's series (F1–F7 filed by s5; F9–F11 s5 via operator; F8 reserved — an s5 numbering gap, never
received at master; F12–F17 assigned here by master, folding the operator-seat and master-seat observations.)*

| id | finding | evidence | class / severity |
|---|---|---|---|
| **F1** | Stale-digest `re-render` bounce fires + recovers correctly — the S3 drift-guard working as designed | intakes 000001/000004/000006/000022 | worked-as-designed (usage datum) |
| **F2** | The rendered form **offers** the dead-edge parent token `s5-dispatch`; lineage rejects it `parent-invalid-dead-edge` | intake-000007 | render↔validate divergence / SEV-2 |
| **F3** | Lineage-class bounces return **no violations detail** in the submit response — reason discoverable only by `read()` of the rejected record (re-render-class DOES carry detail) | intakes 000007/9/10/20/21 | bounce-surfacing gap / SEV-2 |
| **F4** | Parent validity is **projection-order-dependent**: candidate set = the seat's projection tail (`WokenOn`) + same-dispatch accepted set (`cmd/frank/main.go:379-392`, `internal/lineage/lineage.go:346-371`); the semantically-correct reply target and the thread root both bounce `outside-active-lineage`; **rejects sit in the anchor pool** | master intakes 000009/000010 | lineage model / SEV-1 (mechanism of F11) |
| **F5** | `form_digest` moves with **every projection change**; on a busy store a seat's submit is raced by any concurrent accept → re-render churn | s5 intake-000022; master reconnect rotations | liveness / SEV-1 (mechanism of F11) |
| **F6** | **Three envelope parsers disagree**: lineage gate comma-splits raw CC (`lineage.go:432-439`); delivery `addressList` JSON-parses TO/CC and **silently drops ALL header recipients on parse failure**; envelope `To` = one verbatim mailbox key — a comma-joined multi-TO relay (protocol-legal) is deliverable to **neither** recipient | code-verified; s5 adapted with per-seat single-TO copies | correctness — silent delivery loss / SEV-1 |
| **F7** | **CC deadlock** (F6's sharpest edge): canonical-JSON CC passes typing, fails the reviewer-visibility gate; raw CC passes the gate, fails typing (`CC:typed`). **No encoding satisfies both** → every broad-SET orchestrator relay blocked | intakes 000020 (`relay-c13dc32f`) / 000021 (`relay-1f99aadf`); `lineage.go:171-184` × `registry.json:85` | liveness / SEV-1 — mitigated by the operator waiver `relay-52552d61` (→ F17) |
| **F8** | *reserved — s5 numbering gap; no F8 received at master* | — | — |
| **F9** | **`intake_id` reuse**: three distinct rejected records all carry `intake-000012` (different checksums) — breaks §C4.1's 1:1 intake↔outcome anchor ("atomic clear-on-pop, zero stale re-emission"); **path-conditional** (other retries drew fresh ids) | s5 report (its pair's boot retries) | audit-trail correctness / SEV-1 |
| **F10** | **Projection flooding**: rejects are returned by `project()` and dominate a busy store (~13/28 intakes rejected at s5's count; master's inbox at one point 5-of-6 rejects) — floods the seat-facing inbox AND the F4 lineage-anchor pool | s5 report + master projection | UX + anchor pollution / SEV-2 (feeds F4/F11) |
| **F11** | **THE SHOW-STOPPER — lineage livelock under concurrent traffic**: five foreign intakes (029–033) moved the global active-lineage between a seat's render and its submit; the per-seat form still offers the stale local tail; the seat **cannot construct any acceptable parent** — even report-only SITREPs are blocked until new inbound traffic refreshes its tail | s5 report (its F9/F10 SITREP itself bounced this way) | **LIVENESS / SEV-1 — stop-the-line** |
| **F12** | `ORCH_REVIEW_WAIVER` is typed a literal single-value flag (`"*"`) — cannot carry the rationale text the waiver *design* intends; text had to ride SUBJECT/body | operator seat (waiver submission) | schema/design mismatch / SEV-3 (worked around) |
| **F13** | The live validator **rejects tokens the rendered schema advertises**: `gate_category: "other"` rejected; `record_kind: gate_resolution` → `record_kind:unknown` | operator seat (pre-waiver bounces, in archive) | render↔validate divergence (F2's family) / SEV-2 |
| **F14** | **No store lock** — two conductors can serve one store root (observed live: the s4 gate-day leftover + the blessed conductor, simultaneously); the single-writer invariant is unguarded | master ops (leftover pid 16331 killed) | ops / SEV-2 |
| **F15** | **No live seat enrollment** — mint is admin-time-only (`main.go:403-405`), so every mint = a fleet-wide conductor bounce: all channels drop + every in-flight form re-render-bounces (s5's B.4 prediction, confirmed twice this run) | master ops (two mint bounces) | ops / SEV-3 (compounds F5) |
| **F16** | Shim reconnect after a conductor restart needs **two calls**: first returns `shim:connection-lost`, retry re-establishes | master ops | shim UX / SEV-3 |
| **F17** | The reviewer-visibility **waiver is run-wide and irrevocable**: the gate passes on ANY accepted operator waiver record (`lineage.go:176-181`); the store is append-only with **no retraction/scoping record class** — using the escape permanently disables the check on that store | code-verified; waiver `relay-52552d61` live | governance gap / SEV-2 |

## What held (credit where due)

- **Crash-atomic store:** three conductor stop/starts this run — zero corruption, projections rebuilt, mailbox truth
  converged every time.
- **Channel-stamped FROM:** every record on the archive carries conductor-stamped identity; zero confusion incidents
  across 8 live seats.
- **I-PH:** every bounce surfaced `Field:Class` only — no path ever leaked through any seat-delivered surface.
- **The guards that fired, fired honestly:** F1 re-render, the fill-time seat-scoping, and the waiver escape all
  behaved exactly as designed. The failures above are *design gaps*, not broken code paths.

## Addendum (2026-07-06, post-grill): engine latency EXONERATED by transcript audit

Operator-perceived tool slowness was measured across all 117 frank MCP calls in every session transcript (s5 group + hostA + master): `project`/`read` **median 10ms**, `submit` **median 40ms / p90 100ms** — faster than Bash/Edit/Write in the same sessions, fsync-heavy commit path included. The only slow calls (209s, 197s, both gate-day hostA) were **Claude Code MCP permission prompts awaiting the human**, not frank. Perceived latency = permission prompts + the bounce-retry choreography (whose root causes the s6 amendments delete: A-1 digest churn, branch-A parent bounces). **The s6 build should do NO speculative engine-performance work.** Ops note for the relaunch: pre-allowlist `mcp__frank__*` in seat sessions so unattended seats never park on a prompt.

## Disposition

1. **s5 resumes as a standard file-relay team to its checkpoint** (`.relays/s5/s5-resume/` relay) — the
   consumer-schema work never needed the live transport.
2. **The transport-fix cycle is the next build** (after the s5 checkpoint; before the Step-1 exit test rides a live
   store again). Owners: **m-7** (engine: liveness, intake identity, enrollment) + **m-1** (store/lineage: anchor
   model, lock, retraction) + **m-2** (single interpretation: one canonical envelope codec for render, validate, gate,
   delivery). **GRILL_REQUIRED: yes** at its DESIGN step — the parenting model is a hard-to-reverse choice.
3. **Priority order:** the liveness cluster **{F11, F4, F5}** → the correctness pair **{F9, F6}** → the
   render↔validate family **{F2, F13, F12, F3}** → ops **{F14, F15, F16}** → governance **{F17}**.

## Non-binding design seeds (inputs to the fix cycle's AUDIT/DESIGN — not decisions)

- **Conductor-computed PARENT**: the engine already knows the seat's turn context — stamp lineage server-side exactly
  as FROM is stamped; the seat should never supply what the engine knows (kills F4/F11 at the root).
- **One canonical envelope codec** shared by render, typed-validate, lineage gate, and delivery (kills F6/F7/F13).
- **Digest lease or bounded-tolerance revalidate** so a submit races an accept without a bounce (kills F5 churn).
- **Idempotent client submit tokens** → fixes F9 and de-noises retry storms.
- **`project()` excludes rejects by default** (audit view separate); rejects leave the anchor pool (F10).
- **Store lockfile** (F14); **live enrollment or documented bounce choreography** (F15); **waiver scoping + a
  retraction record class** (F17); **bounce-detail parity** for lineage-class rejections (F3).
