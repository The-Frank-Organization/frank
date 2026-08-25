## SITREP — WP2 HOLD at CT-B10: the battery requires invalid UTF-8 bytes literally inside the persisted canonical-JSON journal, an impossible representation; 20 authorized row commits are green and clean through B09, no B10 byte was changed

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16a-wp2-b10-hold
PARENT_DISPATCH_ID: s16a-impl-5
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this requires the chartered master-to-m-9 contract arbitration, not an operator choice
IN_REPLY_TO: s16a-impl/IMPL-planner-20260825-041042.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner
SUBJECT: WP2 deterministic HOLD at B10 — invalid UTF-8 cannot occur literally inside the canonical-JSON session.log; route the test/contract contradiction UP before any B10 byte
STATUS: HOLD

The approved r14/token command contract questions and a test that appears wrong against the contract to route UP, never be silently re-read or predicate-weakened. That condition is met at CT-B10.

## Exact blocker

`TestCT_B10` injects `[]byte{0xff, 0xfe, 0x00, 'x'}` through the live worker provider-item path and requires `bytes.Contains(result.PersistedTranscript, want)`. `result.PersistedTranscript` is the byte-for-byte `session.log`. The journal is a newline stream of canonical JSON records: `journal.FinalizeRecord`/`MarshalRecord` canonicalize each record, and recovery requires every record to decode as canonical JSON before the runtime returns the transcript.

Those two requirements cannot both hold. Bytes `0xff` and `0xfe` are not legal UTF-8 JSON string bytes, and JSON encoding replaces or escapes them; byte `0x00` must also be escaped inside a JSON string. Therefore the four-byte test needle cannot occur literally in a valid canonical-JSON journal. The current runtime honestly preserves the provider item at the provider boundary, then JSON-encodes it for the durable record; the focused test fails only at its literal raw-byte containment assertion:

```text
S16A_SEAM_TAG_ACTIVE
--- FAIL: TestCT_B10 (0.03s)
    provider_worker_test.go:170: CT_B10 contract not yet satisfied: worker opaque-item carriage lost bytes: preserved=false err=<nil>
FAIL
```

The pinned m-9 r11 design requires a replayable canonical journal and opaque carriage of supported provider envelopes; it does not define a binary capsule or a lossless encoding for non-UTF-8 provider items. The design also explicitly excludes an opaque provider capsule from the MVP. The earlier WP1 fidelity return already identified CT-B10 as unfixable before the fold; the successor fold moved the probe into `Runner.Run`, but did not supply a representable journal encoding or decode-roundtrip assertion.

## Unsafe local choices deliberately not taken

1. I did not weaken or rewrite the test to accept replacement characters or an encoded value; test-contract changes require the owning PM's word through master.
2. I did not append raw bytes after the canonical journal or redefine `PersistedTranscript` as a synthetic buffer; either would make the returned evidence differ from the persisted file and violate replay truth.
3. I did not change the journal to a binary container or add a sidecar; that is an unplanned format/contract change beyond this row and this seat.

## Requested ruling

Route this contradiction through master to m-9. Recommended bounded correction: replace literal containment with a defined lossless JSON-safe carrier plus an encode/decode round-trip assertion (for example a closed base64-tagged byte carrier), while retaining the live `Runner.Run` path and the byte-equality check after decoding. If m-9 instead requires raw arbitrary bytes in `session.log`, master must authorize and plan the journal-format amendment and its recovery migration explicitly. On a successor ruling, I resume at B10 from clean HEAD; no unrelated WP2 row needs reopening.

## Banked implementation state

Branch `s16a-conformance` is clean at `89177405a912a3ebbebe1fc0b56d18b92b3cb802`, descended from launch head `f70406668d3db96d882f879d90c6569c7be858b3`. Twenty row/cluster commits are complete:

- `6c080d86` A05; `3bdf3aa0` A10; `b707eefb` A02; `fe94f18a` A03+A04; `4c1b5996` A01.
- `7f954039` A06; `476e6eb7` A07; `9094ecc3` A09; `7a31977a` A11; `99e7721d` A18; `a7d4e4e3` A19.
- `d1719ec0` A13 W+CP; `a1fb7bba` A14 W+CP; `64b9b091` A16 W+CP.
- `817428ad` B01+B02+B06; `08df0ea6` B03; `517dbae9` B04; `1f9bd3db` B05; `8b9ca4c3` B07; `89177405` B09.

At HEAD, the valid tagged bijection is `GREEN=44 RED=20 TOTAL=64`; the only delta from the prior commit is B09 RED-to-GREEN. The B09 per-commit plain `go test -p=1 -count=1 ./...` completed green including `test/fixtures` in 231.854s, and `go vet ./...` is green. The failed focused B10 command made no byte. The carried earlier transient dispositions remain unchanged: the s8 nested-dogfood false-done transient and the unrelated `cmd/frank-mcp` `auth:channel-active` transient both passed on immediate focused rerun and on the respective fresh full-suite rerun, with no out-of-scope edit.

ACTIONS_GIT_REF: implementation commits through `89177405a912a3ebbebe1fc0b56d18b92b3cb802`; this HOLD relay is drafted under `.engine/drafts/s16a.implementer/` and submitted through the daemon; no B10 source/test byte, no push, no PR, no merge, no E3 claim
FINAL_GIT_STATUS_SHORT: implementation worktree clean at `89177405a912a3ebbebe1fc0b56d18b92b3cb802`; governing checkout retains the operator-owned pre-existing relay/INDEX/CHECKPOINTS dirt and is modified only by daemon rendering for this filing
