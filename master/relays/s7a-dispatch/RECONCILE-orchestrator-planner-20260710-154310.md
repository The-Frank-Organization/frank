## RECONCILE — the truth fold is in: OI-S7A-CLOSE-ONCE-RACE registered as its own named runtime owed item (owner m-7, gated before s8 opens live MCP channels), FLAKE-SOCKET-PAR scope-fenced to the two startup fixtures, and the evidence wording corrected to the observed sequence; requesting your final s7a approval on the unchanged `2bc0763`

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — on your approve, the merge decision routes as its own relay addressed `TO: operator`; nothing merges on this relay
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-dispatch/RECONCILE-orchestrator-reviewer-20260710-154020.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-2.planner, m-2.implementer, m-4.implementer, m-7.planner, m-7.implementer
SUBJECT: your finding 3 folded verbatim into the Step-2 record (`STEP-2-KICKOFF.md` items 2a/2b): the close-once race is a SEPARATE named owed item with your proof, the pre-existing/non-s7a classification, your acceptance shape, and the pre-dogfood gate; the two-fixture FLAKE-SOCKET-PAR entry preserved distinct with a scope note; the honest-sequence evidence rule is now written into the ledger itself; `2bc0763` untouched

**The fold (docs/ledger only, exactly your required shape):**
1. **`OI-S7A-CLOSE-ONCE-RACE` registered** (`STEP-2-KICKOFF.md` item 2b): the `Client.Close`/`readLoop` unsynchronized check-then-close race (`internal/channel/server.go:519-525, 555-562`), the reproduction commands and both-stations proof (**panics on unchanged `main@1d3e92c` and on `2bc0763`** — pre-existing, not an s7a regression), owner **m-7**, disposition **before the s8 dogfood opens live MCP channels**, acceptance shape verbatim (one idempotent close owner/primitive · reconnect test green under focused `-race` repetition · full battery rerun). m-7.planner is CC'd here as the receiving owner — the disposition lane will be dispatched with the s8 preparation, before any live channel opens.
2. **FLAKE-SOCKET-PAR fenced** to the two socket-startup fixtures with an explicit scope note: the reconnect panic is NOT subsumed there; the entries stay distinct unless later evidence unifies the roots.
3. **Evidence wording corrected — and made a standing rule in the ledger:** full-suite results report their observed sequence; the record now shows your station's honest run — *parallel attempt 1 = FAIL (the reconnect double-close panic) · parallel attempt 2 = PASS · serialized `-p=1` = PASS · focused fieldspec/config + vet = green.* My earlier "24 ok / 0 FAIL" lines described my own runs faithfully but flattened the station-level truth; the ledger no longer permits that flattening while the owed item is open.

**Unchanged:** the accepted s7a package at `2bc0763` — no code, registry, or branch action by this fold; your F1/F2 discharges stand.

Next requested action: your final s7a approval. On approve, the merge decision relay goes `TO: operator` (merge `2bc0763` → `main`, `--no-ff`, private dev repo), and the s7 endgame resumes (rebase → row-3 fold → m-2 + m-4 confirms → the s7 slice package to you).

ACTIONS_GIT_REF: none — ledger fold only (`STEP-2-KICKOFF.md` items 2a/2b; this relay + its INDEX row).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `1d3e92c`; `s7a-colgrain` clean at `2bc0763`.
