package credentials

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jackli/frank/internal/connector/authorize"
	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/freeze"
	"github.com/jackli/frank/internal/connector/policy"
	"github.com/jackli/frank/internal/connector/translate"
)

const credentialJSON = `{"entries":{"provider-main":{"secret":"sentinel-provider-secret"}},"schema":"m8.credentials.v1"}`

func TestLoadAcceptsPrivateClosedCredentialFile(t *testing.T) {
	t.Parallel()

	store := loadFixture(t, credentialJSON, 0o600)
	if !store.Has("provider-main") || store.Has("missing") {
		t.Fatalf("credential membership mismatch")
	}
	if !ValidReference("provider-main") || ValidReference("Provider_Main") || ValidReference(strings.Repeat("a", 65)) {
		t.Fatal("credential reference grammar mismatch")
	}
}

func TestLoadRejectsInvalidCredentialFiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  string
		mode fs.FileMode
	}{
		{name: "wrong mode", raw: credentialJSON, mode: 0o640},
		{name: "unknown outer field", raw: `{"entries":{"provider-main":{"secret":"x"}},"extra":true,"schema":"m8.credentials.v1"}`, mode: 0o600},
		{name: "unknown entry field", raw: `{"entries":{"provider-main":{"extra":true,"secret":"x"}},"schema":"m8.credentials.v1"}`, mode: 0o600},
		{name: "duplicate ref", raw: `{"entries":{"provider-main":{"secret":"x"},"provider-main":{"secret":"y"}},"schema":"m8.credentials.v1"}`, mode: 0o600},
		{name: "invalid ref", raw: `{"entries":{"Provider_Main":{"secret":"x"}},"schema":"m8.credentials.v1"}`, mode: 0o600},
		{name: "empty secret", raw: `{"entries":{"provider-main":{"secret":""}},"schema":"m8.credentials.v1"}`, mode: 0o600},
		{name: "wrong schema", raw: `{"entries":{"provider-main":{"secret":"x"}},"schema":"m8.credentials.v2"}`, mode: 0o600},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			path := writeFixture(t, test.raw, test.mode)
			if _, err := Load(path); !errors.Is(err, ErrCredentialFileInvalid) {
				t.Fatalf("Load() error = %v, want ErrCredentialFileInvalid", err)
			}
		})
	}
}

func TestValidateFileInfoRejectsOwnerMismatch(t *testing.T) {
	t.Parallel()

	info := fakeFileInfo{mode: 0o600, sys: &syscall.Stat_t{Uid: uint32(os.Geteuid() + 1)}}
	if err := validateFileInfo(info, os.Geteuid()); !errors.Is(err, ErrCredentialFileInvalid) {
		t.Fatalf("validateFileInfo() error = %v, want ErrCredentialFileInvalid", err)
	}
}

func TestAttachResolvesExactlyOnceOnlyAfterAllowedVerdict(t *testing.T) {
	t.Parallel()

	frozen := frozenRequest(t)
	resolver := &countingResolver{secret: []byte("sentinel-provider-secret")}
	verdict := allowedVerdict(t, frozen)
	wire, err := Attach(frozen, verdict, resolver)
	if err != nil {
		t.Fatalf("Attach() error = %v", err)
	}
	if resolver.calls != 1 || resolver.lastRef != "provider-main" {
		t.Fatalf("resolver calls/ref = %d/%q", resolver.calls, resolver.lastRef)
	}
	if wire.core.Method != "POST" || wire.core.Endpoint != "https://api.openai.com/v1/responses" || string(wire.body) != `{}` {
		t.Fatalf("wire identity = %+v", wire)
	}
	authCount := 0
	for _, header := range append(wire.core.Headers, wire.auth) {
		if header.Name == "x-openai-auth" {
			authCount++
			if header.Value != "Bearer sentinel-provider-secret" {
				t.Fatal("auth header did not contain the resolved bearer value")
			}
		}
	}
	if authCount != 1 {
		t.Fatalf("auth header count = %d, want 1", authCount)
	}
}

func TestReferenceWithoutAuthorityNeverResolves(t *testing.T) {
	t.Parallel()

	frozen := frozenRequest(t)
	for _, reason := range []authorize.DenyReason{
		authorize.PolicyUnavailable,
		authorize.PolicyDigestMismatch,
		authorize.MalformedCore,
		authorize.LaneMismatch,
		authorize.LaneEndpointInvalid,
		authorize.EndpointMismatch,
		authorize.EndpointNotAllowlisted,
		authorize.MethodMismatch,
		authorize.ReservedAuthHeaderInCore,
	} {
		resolver := &countingResolver{secret: []byte("sentinel-provider-secret")}
		if _, err := Attach(frozen, authorize.Verdict{DenyReason: reason}, resolver); !errors.Is(err, ErrNotAuthorized) {
			t.Fatalf("Attach(%s) error = %v, want ErrNotAuthorized", reason, err)
		}
		if resolver.calls != 0 {
			t.Fatalf("Attach(%s) invoked resolver %d times", reason, resolver.calls)
		}
	}
}

func TestAttachRejectsEnvelopeMismatchBeforeResolution(t *testing.T) {
	t.Parallel()

	frozen := frozenRequest(t)
	verdict := allowedVerdict(t, frozen)
	verdict.Envelope.FrozenCoreDigest = strings.Repeat("0", 64)
	resolver := &countingResolver{secret: []byte("sentinel-provider-secret")}
	if _, err := Attach(frozen, verdict, resolver); !errors.Is(err, ErrEnvelopeMismatch) {
		t.Fatalf("Attach() error = %v, want ErrEnvelopeMismatch", err)
	}
	if resolver.calls != 0 {
		t.Fatalf("resolver calls = %d, want 0", resolver.calls)
	}
}

func TestCredentialErrorsAndPreAttachStateExcludeSecret(t *testing.T) {
	t.Parallel()

	frozen := frozenRequest(t)
	resolver := &countingResolver{err: errors.New("resolver failed")}
	_, err := Attach(frozen, allowedVerdict(t, frozen), resolver)
	if err == nil || strings.Contains(err.Error(), "sentinel-provider-secret") {
		t.Fatalf("Attach() error leaked sentinel or was nil: %v", err)
	}
	if strings.Contains(string(frozen.CoreBytes()), "sentinel-provider-secret") || strings.Contains(frozen.CoreDigest(), "sentinel-provider-secret") {
		t.Fatal("secret appeared before attach")
	}
}

func TestPrepareHTTPRequestRechecksFrozenCandidateAndDisablesReplay(t *testing.T) {
	t.Parallel()

	frozen := frozenRequest(t)
	wire, err := Attach(frozen, allowedVerdict(t, frozen), &countingResolver{secret: []byte("sentinel-provider-secret")})
	if err != nil {
		t.Fatalf("Attach() error = %v", err)
	}
	request, err := PrepareHTTPRequest(context.Background(), frozen, wire)
	if err != nil {
		t.Fatalf("PrepareHTTPRequest() error = %v", err)
	}
	if request.GetBody != nil {
		t.Fatal("request.GetBody must be nil")
	}
	if request.ContentLength != int64(len(wire.body)) || !request.Close {
		t.Fatalf("request length/close = %d/%v", request.ContentLength, request.Close)
	}

	wire.body = []byte(`[]`)
	if _, err := PrepareHTTPRequest(context.Background(), frozen, wire); !errors.Is(err, freeze.ErrFrozenMutation) {
		t.Fatalf("PrepareHTTPRequest(mutated) error = %v, want ErrFrozenMutation", err)
	}
}

func loadFixture(t *testing.T, raw string, mode fs.FileMode) *Store {
	t.Helper()
	store, err := Load(writeFixture(t, raw, mode))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	return store
}

func writeFixture(t *testing.T, raw string, mode fs.FileMode) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "credentials.json")
	if err := os.WriteFile(path, []byte(raw), mode); err != nil {
		t.Fatalf("write credential fixture: %v", err)
	}
	if err := os.Chmod(path, mode); err != nil {
		t.Fatalf("chmod credential fixture: %v", err)
	}
	return path
}

func frozenRequest(t *testing.T) *freeze.Request {
	t.Helper()
	lane := catalog.Lane{Method: "POST", Endpoint: "https://api.openai.com/v1/responses"}
	result, err := freeze.Freeze(translate.Result{Body: []byte(`{}`), LoweredTools: []byte(`[]`)}, lane, "build-1")
	if err != nil {
		t.Fatalf("Freeze() error = %v", err)
	}
	return result
}

func allowedVerdict(t *testing.T, frozen *freeze.Request) authorize.Verdict {
	t.Helper()
	lane := catalog.Lane{
		LaneID:   "lane-codex-1",
		Method:   "POST",
		Endpoint: "https://api.openai.com/v1/responses",
		Auth:     catalog.Auth{HeaderName: "x-openai-auth", Scheme: "bearer"},
	}
	const rawPolicy = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`
	loaded, err := policy.Load([]byte(rawPolicy), lane)
	if err != nil {
		t.Fatalf("policy.Load() error = %v", err)
	}
	verdict := authorize.Evaluate(authorize.Input{
		Policy:           loaded,
		PolicyDigest:     loaded.Digest,
		ProviderLaneID:   lane.LaneID,
		PinnedLaneID:     lane.LaneID,
		FrozenCoreDigest: frozen.CoreDigest(),
		CredentialRef:    "provider-main",
		Lane:             lane,
		Core:             frozen.Core(),
	})
	if _, ok := verdict.Authorized(); !ok {
		t.Fatalf("authorize.Evaluate() = %+v", verdict)
	}
	return verdict
}

type countingResolver struct {
	calls   int
	lastRef string
	secret  []byte
	err     error
}

func (resolver *countingResolver) resolve(ref string) ([]byte, error) {
	resolver.calls++
	resolver.lastRef = ref
	return append([]byte(nil), resolver.secret...), resolver.err
}

type fakeFileInfo struct {
	mode fs.FileMode
	sys  any
}

func (info fakeFileInfo) Name() string       { return "credentials.json" }
func (info fakeFileInfo) Size() int64        { return 0 }
func (info fakeFileInfo) Mode() fs.FileMode  { return info.mode }
func (info fakeFileInfo) ModTime() time.Time { return time.Time{} }
func (info fakeFileInfo) IsDir() bool        { return false }
func (info fakeFileInfo) Sys() any           { return info.sys }
