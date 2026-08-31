package recovery

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/supervisor"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestAppControlIsNotAConductorSeat(t *testing.T) {
	root := moduleRoot(t)
	command := exec.Command("go", "list", "-deps", "./internal/appctl/...")
	command.Dir = root
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("go list: %v\n%s", err, output)
	}
	for _, forbidden := range []string{
		"github.com/The-Frank-Organization/frank/internal/engine",
		"github.com/The-Frank-Organization/frank/internal/seat",
		"github.com/The-Frank-Organization/frank/internal/store",
		"github.com/The-Frank-Organization/frank/internal/intake",
	} {
		if linePresent(string(output), forbidden) {
			t.Fatalf("appctl imports conductor authority package %s", forbidden)
		}
	}
}

func TestControlAndBrokerTypesCarryNoProviderPayload(t *testing.T) {
	root := moduleRoot(t)
	for _, name := range []string{"msgs_ctrlw.go", "msgs_ctrlc.go", "msgs_broker.go"} {
		contents, err := os.ReadFile(filepath.Join(root, "internal", "appipc", name))
		if err != nil {
			t.Fatal(err)
		}
		lower := strings.ToLower(string(contents))
		for _, forbidden := range []string{"provider_payload", "llmrequest", "normalized_provider_event", "secret_bytes", "credential_bytes"} {
			if strings.Contains(lower, forbidden) {
				t.Fatalf("%s contains forbidden payload/secret type %q", name, forbidden)
			}
		}
	}
}

func TestStoreHasNoCredentialOrProviderPayloadColumn(t *testing.T) {
	root := moduleRoot(t)
	contents, err := os.ReadFile(filepath.Join(root, "internal", "appctl", "store", "schema.go"))
	if err != nil {
		t.Fatal(err)
	}
	lower := strings.ToLower(string(contents))
	for _, forbidden := range []string{"credential_bytes", "seat_credential", "provider_payload", "secret_bytes"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("store schema contains forbidden column family %q", forbidden)
		}
	}
}

func TestSourceSpecificEpochFence(t *testing.T) {
	for _, test := range []struct {
		presented, current string
		want               supervisor.EpochRelation
	}{
		{presented: "4", current: "5", want: supervisor.EpochStale},
		{presented: "5", current: "5", want: supervisor.EpochCurrent},
		{presented: "6", current: "5", want: supervisor.EpochFuture},
	} {
		got, err := supervisor.ClassifyEpoch(test.presented, test.current)
		if err != nil || got != test.want {
			t.Fatalf("ClassifyEpoch(%s,%s)=%s err=%v want=%s", test.presented, test.current, got, err, test.want)
		}
	}
}

func TestCredentialReferenceValidationDoesNotEchoValue(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	runID, epoch := "run", "1"
	_, err = registry.Encode(appipc.Envelope{
		V: 1, Channel: appipc.ChannelCtrlC, Type: "connector_assign", Seq: "0", RunID: &runID, TurnEpoch: &epoch,
		Body: appipc.ConnectorAssignBody{
			RunID: runID, TurnEpoch: epoch, RunManifestDigest: "bad", PolicyDigest: strings.Repeat("0", 64),
			ProviderLaneID: "lane", LaneCatalogDigest: strings.Repeat("0", 64), CredentialRef: "credential-secret-value",
		},
	})
	if err == nil || strings.Contains(err.Error(), "credential-secret-value") {
		t.Fatalf("validation error=%q", err)
	}
}

func moduleRoot(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime caller unavailable")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", "..", ".."))
}

func linePresent(output, want string) bool {
	for _, line := range strings.Split(output, "\n") {
		if line == want {
			return true
		}
	}
	return false
}
