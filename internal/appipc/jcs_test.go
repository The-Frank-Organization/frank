package appipc

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

func TestMarshalJCSSortsByUTF16AndUsesMinimalJSON(t *testing.T) {
	got, err := MarshalJCS(map[string]any{
		"\ue000":     2,
		"\U0001f600": 1,
		"ascii":      "<>&\u2028\b\t\n\f\r\"\\",
	})
	if err != nil {
		t.Fatalf("MarshalJCS: %v", err)
	}
	want := "{\"ascii\":\"<>&\u2028\\b\\t\\n\\f\\r\\\"\\\\\",\"😀\":1,\"\":2}"
	if string(got) != want {
		t.Fatalf("canonical JSON = %q, want %q", got, want)
	}
}

func TestMarshalJCSRejectsFloatsAndInvalidUTF8(t *testing.T) {
	for _, value := range []any{
		map[string]any{"nested": []any{float64(1)}},
		float32(1),
		string([]byte{0xff}),
		map[string]any{string([]byte{0xff}): "bad key"},
	} {
		if _, err := MarshalJCS(value); err == nil {
			t.Fatalf("MarshalJCS(%T) accepted a forbidden trust-path value", value)
		}
	}
}

func TestMarshalJCSReproducesLimitsWitness(t *testing.T) {
	empty := limitsWitnessFrame("")
	encoded, err := MarshalJCS(empty)
	if err != nil {
		t.Fatalf("MarshalJCS(empty witness): %v", err)
	}
	if got, want := len(encoded), 770; got != want {
		t.Fatalf("empty witness length = %d, want %d", got, want)
	}
	if got, want := digestHex(encoded), "60a40957f033255cdfe24b0d1b77a82805726fe727f45625576816c435bfe2f3"; got != want {
		t.Fatalf("empty witness digest = %s, want %s", got, want)
	}

	for _, n := range []int{0, 1, 2, 1000} {
		got, err := MarshalJCS(limitsWitnessFrame(strings.Repeat("a", n)))
		if err != nil {
			t.Fatalf("MarshalJCS(witness %d): %v", n, err)
		}
		if want := 770 + n; len(got) != want {
			t.Fatalf("witness length at n=%d = %d, want %d", n, len(got), want)
		}
	}

	exact, err := MarshalJCS(limitsWitnessFrame(strings.Repeat("a", 3326)))
	if err != nil {
		t.Fatalf("MarshalJCS(exact witness): %v", err)
	}
	if got, want := len(exact), 4096; got != want {
		t.Fatalf("exact witness length = %d, want %d", got, want)
	}
	if got, want := digestHex(exact), "740553dfa398b769f75f8a6358fbbf1c98c94699e431a6323ac444fa7dbc4806"; got != want {
		t.Fatalf("exact witness digest = %s, want %s", got, want)
	}

	over, err := MarshalJCS(limitsWitnessFrame(strings.Repeat("a", 3327)))
	if err != nil {
		t.Fatalf("MarshalJCS(over witness): %v", err)
	}
	if got, want := len(over), 4097; got != want {
		t.Fatalf("over witness length = %d, want %d", got, want)
	}
}

func digestHex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func limitsWitnessFrame(taskInput string) map[string]any {
	provider := map[string]any{
		"kind":       "provider",
		"class":      "uncertain",
		"run_id":     "run-000000000000000000000000",
		"turn_id":    "turn-00000000000000000000",
		"attempt_id": "att-0000000000000001",
		"terminal":   "completed",
	}
	settlement := map[string]any{
		"settlement_manifest_v": 1,
		"run_id":                "run-000000000000000000000000",
		"produced_for_turn_id":  "turn-00000000000000000001",
		"entries":               []any{provider},
	}
	return map[string]any{
		"v":          1,
		"chan":       "ctrl-w",
		"type":       "turn_open",
		"seq":        "0",
		"run_id":     "run-000000000000000000000000",
		"turn_epoch": "1",
		"body": map[string]any{
			"turn_id":             "turn-00000000000000000001",
			"admission_ref":       map[string]any{"kind": "operator_input", "task_input": taskInput},
			"parked_unknown":      []any{},
			"run_disposition":     "resume",
			"create_auth_id":      strings.Repeat("0", 32),
			"session_log_path":    "/var/run/frank/run-000000000000000000000000/session.log",
			"settlement_manifest": settlement,
			"predecessor_turn_id": "turn-00000000000000000000",
		},
	}
}
