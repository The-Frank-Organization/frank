#!/usr/bin/env python3
"""Derive the s16a CT row census from `go test -json` on stdin."""

import json
import re
import sys


EXPECTED = (
    [f"G{i:02d}" for i in range(1, 21)]
    + [f"A{i:02d}" for i in range(1, 20)]
    + [f"B{i:02d}" for i in range(1, 12)]
    + [f"C{i:02d}" for i in range(1, 11)]
    + ["D01", "D02", "D04", "D05"]
)
TEST = re.compile(r"^TestCT_([GABCD][0-9]{2})$")
TAG_SENTINEL = "S16A_SEAM_TAG_ACTIVE"


def main() -> int:
    results: dict[str, str] = {}
    tag_active = False
    for line in sys.stdin:
        try:
            event = json.loads(line)
        except json.JSONDecodeError:
            continue
        if event.get("Action") == "output" and event.get("Output", "").strip() == TAG_SENTINEL:
            tag_active = True
        match = TEST.match(event.get("Test", ""))
        if match and event.get("Action") in {"pass", "fail"}:
            results[match.group(1)] = "GREEN" if event["Action"] == "pass" else "RED"

    if not tag_active:
        print("census invalid: seam build tag sentinel absent", file=sys.stderr)
        return 2

    missing = [row for row in EXPECTED if row not in results]
    extra = sorted(set(results) - set(EXPECTED))
    if missing or extra:
        print(f"census invalid: missing={missing} extra={extra}", file=sys.stderr)
        return 2

    for row in EXPECTED:
        print(f"{row} {results[row]}")
    green = sum(result == "GREEN" for result in results.values())
    red = sum(result == "RED" for result in results.values())
    print(f"SUMMARY GREEN={green} RED={red} TOTAL={len(results)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
