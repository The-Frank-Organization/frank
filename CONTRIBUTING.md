# Contributing

Thanks for your interest in frank.

## The short, honest version

This project is mid-build, and it is developed under an unusual governed multi-agent process: every change flows through the project's own governance discipline (design docs, adversarial review, and gated merges).
Issues, questions, and discussion are very welcome, and bug reports with reproductions are especially valuable.
External PRs are welcome too, but they may be slow to land, because they go through the same governance path as internal changes rather than a fast-merge lane.
Small, focused PRs with a clear problem statement have the best odds.

## Before proposing a change

Run the test suite and make sure it passes:

```
go test ./...
```

The suite is the project's primary evidence surface, so a change that breaks it will not be considered until it is green.
Go 1.25+ is required.

## Licensing

frank is licensed under Apache-2.0.
There is no CLA; by submitting a contribution you agree it is licensed under the project's Apache-2.0 license (see `LICENSE` and `NOTICE`).
