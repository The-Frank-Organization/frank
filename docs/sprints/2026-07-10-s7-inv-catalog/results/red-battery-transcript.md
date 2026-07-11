# s7 INV-CATALOG red-battery transcript

Source implementation commit: `s7-inv-catalog@eaaf5f0`

Scratch branch/worktree:

- Branch: `s7-red-battery-demo`
- Worktree: `/Users/jack/.config/superpowers/worktrees/frank/s7-red-battery-demo`
- Deliberate weakening: row 1's literal expected set changed from `{"accepted", "rejected", "held"}` to `{"accepted", "rejected"}`.

## Red

Command:

```text
go test -count=1 ./test/invariants
```

Exit: 1

```text
--- FAIL: TestLawTerminalEnumByteExact (0.00s)
    terminal_surface_test.go:31: terminal enum bytes = ["accepted" "rejected" "held"], want literal ["accepted" "rejected"]
FAIL
FAIL    github.com/jackli/frank/test/invariants    1.402s
FAIL
```

The failure names the weakened law and reports the missing literal token.

## Discard

`git worktree remove --force /Users/jack/.config/superpowers/worktrees/frank/s7-red-battery-demo` exited 0.

`git branch -D s7-red-battery-demo` deleted the scratch branch at `eaaf5f0`.

No weakening was applied to `s7-inv-catalog`.

## Green

Command on the real branch:

```text
go test -count=1 ./test/invariants
```

Exit: 0

```text
ok      github.com/jackli/frank/test/invariants    1.105s
```
