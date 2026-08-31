package journal_test

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jackli/frank/internal/worker/journal"
)

// voidPendingRegenerationDigest is the SHA-256 of the edit-base substrate
// whose freeze was declared VOID at the member grain by the S16-WP3-F1
// composed disposition (s16-wp3-f1-rule2/RECONCILE-orchestrator-planner-
// 20260830-180414: four kinds carry non-contract payload members). The
// witness skips while that exact substrate is still on disk and arms itself
// the moment the Master+VP-re-frozen regenerated substrate lands.
const voidPendingRegenerationDigest = "feb1bf6cd25ce65469cc116551bce111214f74f4260eef1a7e01dde7b6f7d6db"

// TestFrozenSubstrateConformance is the S16-WP3-F1 conformance witness: the
// production reader must recover the UNTOUCHED frozen edit-base substrate
// (authored against the reader contract's §2.7 envelope+payload grammar)
// with disposition resumable. The substrate bytes are read in place from the
// frozen corpus and are never copied, translated, or repaired.
func TestFrozenSubstrateConformance(t *testing.T) {
	data := readFrozenSubstrate(t)
	digest := sha256.Sum256(data)
	if hex.EncodeToString(digest[:]) == voidPendingRegenerationDigest {
		t.Skip("edit-base substrate freeze is VOID-PENDING-REGENERATION at the member grain (s16-wp3-f1-rule2 disposition); the witness arms when the regenerated substrate lands")
	}

	var genesis struct {
		Kind    string `json:"kind"`
		Payload struct {
			RunID             string `json:"run_id"`
			RunManifestDigest string `json:"run_manifest_digest"`
			CreateAuthID      string `json:"create_auth_id"`
		} `json:"payload"`
	}
	firstLine := data
	if index := bytes.IndexByte(data, '\n'); index >= 0 {
		firstLine = data[:index]
	}
	if err := json.Unmarshal(firstLine, &genesis); err != nil {
		t.Fatalf("decode substrate genesis line: %v", err)
	}
	if genesis.Kind != journal.KindRunOpen {
		t.Fatalf("substrate does not open with run_open: %q", genesis.Kind)
	}

	identity := journal.Identity{
		RunID:             genesis.Payload.RunID,
		RunManifestDigest: genesis.Payload.RunManifestDigest,
		CreateAuthID:      genesis.Payload.CreateAuthID,
	}
	recovery := journal.Recover(data, identity)
	if recovery.Disposition != journal.DispositionResumable {
		t.Fatalf("frozen substrate did not recover resumable: disposition=%q fault=%q genesis_fault=%v",
			recovery.Disposition, recovery.FaultClass, recovery.GenesisFault)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	lines := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		if _, err := journal.DecodeRecord(line); err != nil {
			t.Fatalf("substrate line %d does not decode under the production reader: %v", lines+1, err)
		}
		lines++
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	if lines == 0 {
		t.Fatal("substrate is empty")
	}
}

func readFrozenSubstrate(t *testing.T) []byte {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate test source")
	}
	path := filepath.Clean(filepath.Join(filepath.Dir(file),
		"../../../..", "master", "exit-fixtures", "common", "edit-base-journal.jsonl"))
	data, err := os.ReadFile(path)
	if err != nil {
		t.Skipf("frozen substrate not present at %s (standalone frank checkout): %v", path, err)
	}
	return data
}
