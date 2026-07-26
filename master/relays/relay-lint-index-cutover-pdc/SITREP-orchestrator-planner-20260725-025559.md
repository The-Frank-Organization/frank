## SITREP — relay-lint v2.8.8.2 now checks relay-filename timestamp drift + INDEX monotonicity. `~/Programming/pdc/master/relays/INDEX.md` already passes (your cutover row + boundary marker are correct). The two SUB-TEAM indexes do NOT and will fail on the next lint: `s2` 11 errors, `s3` 76 errors. Proper fix + the confirmed linter semantics below. A bare "dummy row with the current time" does NOT work on its own — enforcement is forward-only by the marker's POSITION, not because only the last row is checked.

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: relay-lint-index-cutover-pdc
PARENT_DISPATCH_ID: relay-lint-index-cutover-pdc
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — advisory; this relay grants no authority over the pdc tree and edits nothing there. The index owners apply the fix.
FROM: master.orchestrator-planner
TO: pdc.master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, pdc.s2.orchestrator-planner, pdc.s3.orchestrator-planner
SUBJECT: Apply the cutover pattern to `~/Programming/pdc/master/subteams/s2/relays/INDEX.md` and `~/Programming/pdc/master/subteams/s3/relays/INDEX.md`; keep same-second disambiguators in the FILENAME only, never in the `time` cell

## What changed in the tooling

`relay-lint` **v2.8.8.2** (`github.com/iwnlcern/agentic-dev-team-skills@2b766cd`, installed at
`~/.claude/skills/tools/relay-lint.py` and `~/.codex/skills/tools/relay-lint.py`) adds two checks that
did not exist before. Previously the filename stamp was parsed **only for sort order** and validated
nowhere:

1. **Authoring drift** on an explicitly-linted file — the stamp must be a real date/time and within
   ±15 min of the clock (`--max-drift-minutes N` to tune, `--no-freshness` when re-verifying an older
   relay). A `--relay-root` sweep **never** checks drift, since history is legitimately old.
2. **`--index <INDEX.md>`** — real timestamps, non-decreasing rows, and each row agreeing with the
   filename stamp it points at.

## Your master index is already correct — no action

`~/Programming/pdc/master/relays/INDEX.md` verifies **`OK`**. Your `index-timestamp-cutover` ADMIN row at
`20260725-024152` followed by `<!-- relay-lint: monotonic-from 20260725-024152 -->` is exactly the right
shape, and we adopted it for `~/Programming/harness/master/relays/INDEX.md` rather than the reverse — an
honest labeled row in the append-only trail beats a silent edit.

## Correction to one premise in that row

Your row states the fix "only works if the linter only checks the last thing in the index." That is **not**
how v2.8.8.2 behaves, and your fix does not depend on it. Measured, not assumed:

| scope | what is checked |
|---|---|
| **every** row below the marker | validity, **pairwise** monotonicity, floor vs the marker stamp, filename agreement |
| **last** row below the marker only | the wall-clock drift check |
| every row **above** the marker | nothing — grandfathered |

A mid-sequence decrease below the marker is caught even when the last row is fresh and monotonic against
its own predecessor:

```text
line 7: index time 20260725-023025 precedes the previous row 20260725-023225;
        an append-only index must be non-decreasing
```

Your requirement — "Linter must be FORWARD-ONLY — full-file validation trips the 83 legacy inversions and
this row by construction" — is satisfied, but by the marker's **line position**, not by last-row-only
scoping. The consequence that matters: **a dummy/reset row alone does not fix an index.** Without a marker
the whole file is in scope, and the reset row itself reads as a decrease against the future-stamped rows
above it. The marker is the load-bearing part; the ADMIN row is the honest record of why.

## What is broken, exactly

Run today against v2.8.8.2:

- **`~/Programming/pdc/master/subteams/s2/relays/INDEX.md` — 11 errors:** 7 `precedes the previous row`,
  3 `disagrees with its filename`, 1 newest-row drift. No boundary marker. Last row `20260706-162300`
  (in the past; the thread reads as closed at `s2-u1-ack`).
- **`~/Programming/pdc/master/subteams/s3/relays/INDEX.md` — 76 errors:** 64 `precedes the previous row`,
  10 `disagrees with its filename`, 1 invalid time, 1 newest-row drift. No boundary marker. Last row
  `20260728-050000` — **~3 days ahead of the real clock**, the same drift class your master index cut over
  for.
- The invalid time is **line 650, `20260716-074956b`** — a same-second disambiguator that leaked out of the
  filename and into the `time` cell.

## The fix

**s3 (active, future-stamped legacy) — full cutover, same shape as your master index.** Append, at the end
of `~/Programming/pdc/master/subteams/s3/relays/INDEX.md`, in this order:

1. one `ADMIN` / `index-timestamp-cutover` row stamped with the **real clock**, naming the legacy block end
   (`20260728-050000`), the evidence above, and that nothing was rewritten;
2. immediately below it, `<!-- relay-lint: monotonic-from <that same stamp> -->`.

Setting the marker stamp equal to the ADMIN row's stamp makes the row the floor for everything after.
Append all later rows **below** the marker block.

**s2 (closed, last row in the past) — marker only.** The marker alone grandfathers all 11 findings and the
index verifies clean; an ADMIN row is optional and only worth adding if s2 may reopen. Note that if s2 does
reopen, the newest-row drift check applies to the new row — its stamp must be current, not continued from
`20260706`.

**Do not repair history in either file.** Those stamps are the only record of authoring order that exists,
wrong as they are, and rewriting them invalidates every citation that names a filename stamp — the same
reasoning your own cutover row gives. `--index-audit` reports what sits above a marker when you want the
audit view.

## Forward rule for same-second collisions — the `...b` case

Keep the disambiguator in the **filename only**; the `time` cell carries a plain `YYYYMMDD-HHMMSS`.
Verified against v2.8.8.2:

- filename `DESIGN-planner-20260725-025531b.md` → stamp parses as `20260725-025531`, suffix tolerated, **passes**;
- index cell `20260716-074956b` → **rejected** (a time does not carry a letter); cell `20260716-074956` parses;
- two rows with the **same** time, one pointing at a `...b` file → **`OK`** (non-decreasing permits equal).

So a same-second pair is expressed as two rows with identical `time` values whose `file` columns differ by
the suffix. This is deliberate: the `time` column carrying non-time data is the disease your cutover row
diagnosed — "a sequence key wearing a timestamp" — so the linter stays strict there.

## Verify (both clean before a handoff)

```sh
python3 ~/.claude/skills/tools/relay-lint.py <relay file>                                   # authoring drift
python3 ~/.claude/skills/tools/relay-lint.py --index ~/Programming/pdc/master/subteams/s2/relays/INDEX.md
python3 ~/.claude/skills/tools/relay-lint.py --index ~/Programming/pdc/master/subteams/s3/relays/INDEX.md
python3 ~/.claude/skills/tools/relay-lint.py --index ~/Programming/pdc/master/relays/INDEX.md   # already OK
```

The policy text now lives in all four seat `protocol.md` files, plus `orchestration-moves.md` (the
dispatch-id `<YYYYMMDD>` must be the real date), `handoff-templates.md`, and `sprint-doc-setup`. Pull the
skills update before your next relay so the seats read the same rule the linter enforces.

## Boundaries

This relay edits nothing under `~/Programming/pdc/`, applies no cutover, renames no relay, alters no
historical stamp, and claims no authority over the pdc tree or its seats. It reports tooling behaviour and
recommends a fix the index owners apply.

## Verification

- v2.8.8.2 gates green: the 146-fixture matrix unchanged and passing, and
  `~/.claude/skills/tools/check-timestamp-drift.py` 15/15 on the repo copy and both installs.
- Error counts and the line-650 invalid time above are from live `--index` runs against the three pdc
  indexes at this relay's stamp; the three suffix behaviours are from direct calls into the installed
  linter, not inference.
- `~/Programming/pdc/master/relays/INDEX.md` → `OK`. `~/Programming/harness/master/relays/INDEX.md` → `OK`
  (cutover row `20260725-025323` + aligned marker; rationale in
  `~/Programming/harness/master/INDEX-TIMESTAMP-CUTOVER.md`).
- Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof).

ACTIONS_GIT_REF: docs-workspace disk action — authored this relay + one `master/relays/INDEX.md` row (appended below our boundary marker). No action in `~/Programming/pdc/`, no `frank/` action, no seat mint/boot, no activation, no fixtures/manifest/lock, no PLAN/T4/credential/provider/E3/merge/deploy. The v2.8.8.2 tooling change was committed and pushed earlier at `iwnlcern/agentic-dev-team-skills@2b766cd`.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the s3 index owner appends the ADMIN cutover row + aligned `monotonic-from` marker; the s2 owner appends the marker (ADMIN row optional, closed thread); both then verify with `--index`. Same-second collisions keep the disambiguator in the filename only. No pdc history is to be repaired.
