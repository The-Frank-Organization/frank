## SITREP -- master.orchestrator-reviewer / org decomposition review

ROLE: Orchestrator Reviewer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-org-decomp
PARENT_DISPATCH_ID: master-org-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator must lock product scope and org depth before staffing
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator

Verdict: human-decision-required.

Boot acknowledgement. I am loaded as `master.orchestrator-reviewer` for RUN_ID `master`, with report-only authority. I read the charter (`CLAUDE.md` / `AGENTS.md`), the boot relay, the dashboard, the relay index, and the addressed `master-org-decomp` relay. Relay root is reachable at `master/relays/`. The boot relay and `master-org-decomp` each lint clean as standalone files.

Finding 1 -- product-scope expansion is not locked. The charter's product statement says v3 is "the conductor": a thin trusted courier with locked routing/policy/seat-trust decisions, and points to the extracted v3 design state. The reviewed relay says v3 has expanded into "a model/provider-agnostic, governed, multi-agent dev-team harness shipped as a TUI app" with a Claude-Code-style runtime, Zellij-style multiplexer, and integrated email client. That may be the right destination, but it is a major product-boundary change from conductor-governance substrate to full runtime/product. This is an operator-judgment category: product semantics, broad scope expansion, and ownership. Do not staff Runtime/Product/TUI/Email domains as equal Cycle-1 design domains until the operator explicitly confirms whether Cycle c1 designs (A) conductor substrate only, (B) conductor plus runtime seams, or (C) the full TUI/runtime/email product.

Finding 2 -- flat-and-lean Cycle 1 is the right operating shape, conditional on that scope lock. I agree with rejecting the 4-tier recursive org as the current operating structure. Evidence: the charter says pairs are independent operator-relayed sessions, the current phase is AUDIT + DESIGN only, and v2.8.8 models one orchestrator pair over flat downstream pairs. Adding T2 division orchestrators and T4 feature teams now multiplies manual relay hops and invents nested-run conventions before v3 exists to automate them. Capture recursive/fractal org support as a v3 design requirement, not as this phase's manual operating topology.

Finding 3 -- the 6 + 2 + 4 domain split is useful as a map, not as immediate staffing. The six conductor domains are coherent and already durable in the charter. The Runtime and Product domains are under-specified relative to the locked conductor substrate: they consume m-1 store/identity, m-2 schema, m-3 observe gates, m-4 routing, and m-6 human surfaces, but their exact authority and deliverables are not yet bounded. Treat Runtime/Product as consumer lenses or future expansion slots until m-1/m-2 interface sketches exist. If the operator chooses full product scope for c1, add explicit domain charters for Runtime Core, Provider Adapters, TUI Shell, Email Client, UI/UX, and Platform Plumbing before staffing them.

Finding 4 -- "m-1 and m-2 lock first" is directionally right but must not freeze in isolation. m-1 store/seat-stamping and m-2 typed schema are load-bearing foundations, but m-3 observation, m-4 routing records, and m-6 human-gate fields are primary consumers. Required change: make m-1/m-2 produce interface sketches first, then route consumer reviews before design-lock. Otherwise the team risks a writer-with-no-reader contract or a schema that cannot express observe-as-send, routing justification, or governance-vs-collaboration surfaces.

Required edits before staffing:
1. Add an operator decision line to `master/README.md` or a relay: `C1_PRODUCT_SCOPE: conductor-only | conductor-plus-runtime-seams | full-runtime-product`, with the operator's choice.
2. Update the org decomposition to separate "Cycle c1 staffed domains" from "future supported org/product domains."
3. For m-1 and m-2, require design outputs to include consumer-facing interface sketches and a named review pass by m-3/m-4/m-6 before design-lock.
4. If recursive 4-tier org is deferred, record it as an m-5 / architecture-of-record requirement, not as a rejected product capability.

ACTIONS_GIT_REF: wrote reviewer relay at master/relays/master-org-decomp/SITREP-orchestrator-reviewer-20260628-031820.md; docs workspace git status unavailable because cwd is not a git repo; pcode is the future code repo and was not edited.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
