# OI-S3-CONFIG-CHANGE — typed owed record (ledger materialization)

Materialized per master's s3-scope-q1 ruling condition 2 (`.relays/s3/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md`), the OI-S1-F11-SWEEP pattern. Ledger-materialized here (the ruling's stated minimum); the optional real-store submit is the operator's channel to author (the S2 Q2=(i) precedent: owed records on the real store are operator-authored, never synthetically stamped).

```text
OI-S3-CONFIG-CHANGE = {
  owner: the wire-up slice (default; HARD BACKSTOP: before any store is declared persistent/long-lived, whichever comes first — no persistent store without a config-change path),
  source: m-7 design §7 :109 (operator-authorized committed store record) + s2 design §1 :18 forward-pointer ("the §7 config-change record (S3)" — SUPERSEDED by the s3-scope-q1 ruling, condition 5; the physical one-line fold rides the next m-7-guided touch) + the s3-scope-q1 ruling itself,
  target_surface: the §7 config-change record — commit-loop mutation class + recovery interaction + operator-authorized digest-change record + crash-matrix applicability class; carries (b)'s conditions wherever it lands: m-7 guides · m-1 fidelity on the new record_kind · the S2 crash-harness applicability map gains the mutation class,
  disposition_path: the owning slice's exit gate
}
```

Standing consequence for S3 (ruling conditions 1 + the sharpening): fresh-store posture ratified — every S3 claim surface carries "registry rides `store.Init`; registry evolution on an existing store awaits the §7 config-change record"; restart-with-new-store IS the true current semantics of a registry change under the locked no-hot-reload config model, so the S3 drift fixtures test the real mechanism, not a proxy.
