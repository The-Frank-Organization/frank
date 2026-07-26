# Index timestamp cutover — one-time backward step

**Effected 2026-07-25 by the master orchestrator-planner, operator-directed, alongside relay-lint v2.8.8.2.**

## What changed

`master/relays/INDEX.md` carries a positional boundary at its end:

```text
<!-- relay-lint: monotonic-from <stamp> -->
```

Rows **below** it are held to the rule — real timestamps, never decreasing, at/after the boundary, and
matching each relay's own filename stamp. Rows **above** it are grandfathered history, inspected with
`relay-lint --index master/relays/INDEX.md --index-audit`.

Immediately above the boundary sits an `ADMIN` row (`index-timestamp-cutover`) recording the step in the
append-only trail itself, so the discontinuity is a labeled record rather than a silent edit.

## Why — the evidence at cutover

The `time` column had been functioning as a sequence key wearing a timestamp: stamps were inferred from
neighbouring relays and advanced to keep a tidy order, not read off a clock. Measured across 2293 rows:

- **110 rows stamped in the future**, the newest by **~3.4 days** (`2026-07-28 12:00:00` against a real
  clock of `2026-07-25 02:32`);
- **297 non-monotonic** adjacent pairs;
- **2 impossible times** (`20260712-016200` — minute 62), one of which is a real relay *file*
  (`s8-claim-input-m3/DESIGN-REVIEW-implementer-20260712-016200.md`);
- **3** rows disagreeing with the filename they point at;
- a synthetic exactly-hourly cadence across the most recent block.

## Why the boundary is positional, not value-based

A value-based cutoff ("rows after time T must be monotonic") would be defeated by the pre-existing
future-stamped rows: the next honest row, stamped at the real clock, sorts *below* them, so the rule would
force every later row to keep matching the fabrication (≥ `2026-07-28`) instead of the real clock. The
boundary is therefore a **line position**, which makes enforcement forward-only by construction.

## Nothing was rewritten

No relay renamed, no historical stamp altered, no citation touched. The append-only trail is intact; the
cutover is additive. Repairing history was rejected deliberately — the stamps are the only record of
authoring order that exists, wrong as they are, and rewriting them would invalidate every relay citation
that names a filename stamp.

## The rule from here

Read the clock. Never infer a stamp from a sibling relay, continue a numbering pattern, or round to a tidy
cadence. Append new index rows **after** the boundary block, in `YYYYMMDD-HHMMSS` form matching the
relay filename they point at.

Verification (both must be clean before a handoff):

```sh
python3 <skills>/tools/relay-lint.py <relay file>              # authoring-drift check
python3 <skills>/tools/relay-lint.py --index master/relays/INDEX.md
```

`--relay-root` sweeps never check drift (history is legitimately old); `--no-freshness` is for
re-verifying an older relay.

## Linter behaviour this relies on (confirmed empirically, v2.8.8.2)

- **All** post-boundary rows: validity, pairwise monotonicity, floor vs the boundary stamp, filename
  agreement. A decrease is caught **anywhere** in the post-boundary block, not only in the last row.
- **Last** post-boundary row only: the wall-clock drift check.
- Rows above the boundary: not checked.

Enforcement is forward-only via the boundary's position — not because only the newest row is examined.
