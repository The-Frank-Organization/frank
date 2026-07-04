# frank

frank is the S1 thin conductor relay for the harness sprint. S1 = provenance + transport, not verified work.

The S1 result records transport and relay-state facts. A work claim carried through this slice is self_reported unless a later evidence layer verifies it. The live terminal state values are exactly `accepted`, `rejected`, and `held`.

Identity and store access are tool-mediated confusion-resistance. The per-seat surface removes ordinary affordances for writing store files directly, but the D5 residual remains: a same-uid process outside the tool surface can still reach local files, sockets, and config.

The commit loop is the sole in-process writer for governance-surface mutations, with the D5 residual above applying to out-of-tool-surface local processes. Crash fixtures exercise the rename-pivot path and recovery replay at E2.

S1 renders `grant` only on operator/orchestrator forms. Conditional pair-Planner grant rendering lands in S3, and this S1 schema keeps that later landing additive rather than foreclosed.
