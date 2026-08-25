package journal

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRecoveryCrashTable(t *testing.T) {
	genesis := genesisRecord(t, "generation-a")
	cleanRound, marker := completedToolRound(t, "generation-a")
	cases := []struct {
		name            string
		journal         []byte
		wantDisposition string
		wantBoundary    BoundaryKind
		wantNextSeq     uint64
	}{
		{name: "clean genesis", journal: journalBytes(t, genesis), wantDisposition: DispositionResumable, wantBoundary: BoundaryGenesis, wantNextSeq: 1},
		{name: "torn tail", journal: append(journalBytes(t, genesis), []byte(`{"seq":`)...), wantDisposition: DispositionDegraded, wantBoundary: BoundaryGenesis, wantNextSeq: 1},
		{name: "mid append complete unmarked tail", journal: journalBytes(t, append([]Record{genesis}, cleanRound...)...), wantDisposition: DispositionResumable, wantBoundary: BoundaryGenesis, wantNextSeq: 1},
		{name: "marker durable", journal: journalBytes(t, append(append([]Record{genesis}, cleanRound...), marker)...), wantDisposition: DispositionResumable, wantBoundary: BoundaryMarker, wantNextSeq: 4},
		{name: "post marker untrusted suffix", journal: journalBytes(t, append(append(append([]Record{genesis}, cleanRound...), marker), withRoundAndSeqFinalized(t, recordForKind(t, KindInputItem), "1", "4"))...), wantDisposition: DispositionResumable, wantBoundary: BoundaryMarker, wantNextSeq: 4},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Recover(testCase.journal, expectedIdentity())
			if result.Disposition != testCase.wantDisposition || result.Boundary.Kind != testCase.wantBoundary || result.NextSeq != testCase.wantNextSeq {
				t.Fatalf("Recover = %+v, want disposition=%s boundary=%s next=%d", result, testCase.wantDisposition, testCase.wantBoundary, testCase.wantNextSeq)
			}
			if result.Boundary.Offset > int64(len(testCase.journal)) {
				t.Fatalf("boundary offset %d exceeds journal length %d", result.Boundary.Offset, len(testCase.journal))
			}
		})
	}
}

func TestRecoveryFirstFaultWinsAndLaterMarkerCannotRescue(t *testing.T) {
	genesis := genesisRecord(t, "generation-a")
	round, marker := completedToolRound(t, "generation-a")
	bad := round[0]
	bad.Seq = "9"
	bad = mustFinalized(t, bad)
	journal := journalBytes(t, genesis, bad, round[1], marker)
	result := Recover(journal, expectedIdentity())
	if result.Disposition != DispositionDegraded || result.Boundary.Kind != BoundaryGenesis || result.FaultClass != FaultSequence {
		t.Fatalf("Recover = %+v", result)
	}
}

func TestRecoveryContentTrustClassifiesExternalAndModelBytesDifferently(t *testing.T) {
	genesis := genesisRecord(t, "generation-a")

	model := withSeqFinalized(t, recordForKind(t, KindToolCall), "1")
	model.Fields["args"] = json.RawMessage(`{"path":"changed-but-replayed-verbatim"}`)
	modelMarker, err := BuildRoundMarker("2", "generation-a", "turn-1", "0", "2", []Record{model})
	if err != nil {
		t.Fatalf("BuildRoundMarker(model): %v", err)
	}
	modelResult := Recover(journalBytes(t, genesis, model, modelMarker), expectedIdentity())
	if modelResult.Disposition != DispositionResumable || modelResult.Boundary.Kind != BoundaryMarker {
		t.Fatalf("model-transcript mismatch result = %+v", modelResult)
	}

	external := withSeqFinalized(t, recordForKind(t, KindProviderOutput), "1")
	external.Fields["content"] = rawString(t, "changed external bytes")
	externalMarker, err := BuildRoundMarker("2", "generation-a", "turn-1", "0", "2", []Record{external})
	if err != nil {
		t.Fatalf("BuildRoundMarker(external): %v", err)
	}
	externalResult := Recover(journalBytes(t, genesis, external, externalMarker), expectedIdentity())
	if externalResult.Disposition != DispositionDegraded || externalResult.Boundary.Kind != BoundaryGenesis || externalResult.FaultClass != FaultContentTrust {
		t.Fatalf("external-truth mismatch result = %+v", externalResult)
	}
}

func TestRecoveryRejectsSuspectMarkerToPriorBoundary(t *testing.T) {
	genesis := genesisRecord(t, "generation-a")
	round, marker := completedToolRound(t, "generation-a")
	marker.Fields["marker_digest"] = rawString(t, strings.Repeat("0", 64))
	marker = mustFinalized(t, marker)
	result := Recover(journalBytes(t, append(append([]Record{genesis}, round...), marker)...), expectedIdentity())
	if result.Disposition != DispositionDegraded || result.Boundary.Kind != BoundaryGenesis || result.FaultClass != FaultMarker {
		t.Fatalf("Recover = %+v", result)
	}
}

func TestRecoveryGenesisFaultHasEmptyTrustedPrefix(t *testing.T) {
	genesis := genesisRecord(t, "generation-a")
	genesis.Fields["run_manifest_digest"] = rawString(t, strings.Repeat("f", 64))
	genesis = mustFinalized(t, genesis)
	result := Recover(journalBytes(t, genesis), expectedIdentity())
	if result.Disposition != DispositionDegraded || result.ResumeAction != ResumeActionReDerive || result.Boundary.Kind != BoundaryNone || result.Boundary.Offset != 0 || !result.GenesisFault {
		t.Fatalf("Recover = %+v", result)
	}
}

func TestAdjacentByteIdenticalDuplicateCollapses(t *testing.T) {
	genesis := genesisRecord(t, "generation-a")
	line := append(mustMarshal(t, genesis), '\n')
	journal := append(append([]byte(nil), line...), line...)
	result := Recover(journal, expectedIdentity())
	if result.Disposition != DispositionResumable || result.Boundary.Kind != BoundaryGenesis || result.NextSeq != 1 {
		t.Fatalf("Recover = %+v", result)
	}
}

func expectedIdentity() Identity {
	return Identity{RunID: "run-1", RunManifestDigest: strings.Repeat("a", 64), CreateAuthID: strings.Repeat("b", 32)}
}

func genesisRecord(t *testing.T, generationID string) Record {
	t.Helper()
	record := recordForKind(t, KindRunOpen)
	record.GenerationID = generationID
	return mustFinalized(t, record)
}

func completedToolRound(t *testing.T, generationID string) ([]Record, Record) {
	t.Helper()
	call := withCoordinates(recordForKind(t, KindToolCall), "1", generationID, "turn-1", "0")
	result := withCoordinates(recordForKind(t, KindToolResult), "2", generationID, "turn-1", "0")
	members := []Record{mustFinalized(t, call), mustFinalized(t, result)}
	marker, err := BuildRoundMarker("3", generationID, "turn-1", "0", "3", members)
	if err != nil {
		t.Fatalf("BuildRoundMarker: %v", err)
	}
	return members, marker
}

func withSeqFinalized(t *testing.T, record Record, seq string) Record {
	t.Helper()
	record.Seq = seq
	record.TSMonotonic = seq
	return mustFinalized(t, record)
}

func withRoundAndSeqFinalized(t *testing.T, record Record, roundIndex, seq string) Record {
	t.Helper()
	record.RoundIndex = roundIndex
	return withSeqFinalized(t, record, seq)
}

func journalBytes(t *testing.T, records ...Record) []byte {
	t.Helper()
	var output bytes.Buffer
	for _, record := range records {
		output.Write(mustMarshal(t, record))
		output.WriteByte('\n')
	}
	return output.Bytes()
}
