// Package credentials owns the m-8-private S-A credential file and the sole
// post-authorization resolution/attachment path.
package credentials

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"syscall"

	"github.com/jackli/frank/internal/connector/authorize"
	"github.com/jackli/frank/internal/connector/freeze"
	"github.com/jackli/frank/internal/connector/jcs"
)

const credentialFileMax = 1024 * 1024

var (
	ErrCredentialFileInvalid = errors.New("credential-file-invalid")
	ErrCredentialResolve     = errors.New("credential-resolve-failed")
	ErrNotAuthorized         = errors.New("credential-attach-not-authorized")
	ErrEnvelopeMismatch      = errors.New("authorized-envelope-mismatch")
)

type Store struct {
	entries map[string][]byte
}

type Resolver interface {
	resolve(ref string) ([]byte, error)
}

// WireRequest is the sole secret-bearing request representation. Its fields
// stay private and it deliberately has no JSON or text serializer.
type WireRequest struct {
	core authorize.Core
	auth authorize.Header
	body []byte
}

func (wire *WireRequest) String() string { return "credentials.WireRequest{redacted}" }

func Load(path string) (*Store, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, invalidFile("open")
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return nil, invalidFile("stat")
	}
	if err := validateFileInfo(info, os.Geteuid()); err != nil {
		return nil, err
	}
	raw, err := io.ReadAll(io.LimitReader(file, credentialFileMax+1))
	if err != nil || len(raw) == 0 || len(raw) > credentialFileMax {
		return nil, invalidFile("read")
	}
	// Canonicalization is used only as a duplicate-key and JSON validity
	// oracle. Credential files need not themselves be canonical.
	if _, err := jcs.Canonicalize(raw); err != nil {
		return nil, invalidFile("duplicate or malformed JSON")
	}
	type entryWire struct {
		Secret *string `json:"secret"`
	}
	var wire struct {
		Entries map[string]entryWire `json:"entries"`
		Schema  *string              `json:"schema"`
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&wire); err != nil {
		return nil, invalidFile("closed schema")
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, invalidFile("trailing data")
	}
	if wire.Schema == nil || *wire.Schema != "m8.credentials.v1" || len(wire.Entries) == 0 {
		return nil, invalidFile("schema or entries")
	}
	store := &Store{entries: make(map[string][]byte, len(wire.Entries))}
	for ref, entry := range wire.Entries {
		if !ValidReference(ref) || entry.Secret == nil || *entry.Secret == "" {
			return nil, invalidFile("entry")
		}
		store.entries[ref] = append([]byte(nil), []byte(*entry.Secret)...)
	}
	return store, nil
}

func (store *Store) Has(ref string) bool {
	if store == nil || !ValidReference(ref) {
		return false
	}
	_, ok := store.entries[ref]
	return ok
}

func (store *Store) resolve(ref string) ([]byte, error) {
	if store == nil {
		return nil, ErrCredentialResolve
	}
	secret, ok := store.entries[ref]
	if !ok {
		return nil, ErrCredentialResolve
	}
	return append([]byte(nil), secret...), nil
}

func ValidReference(value string) bool {
	if len(value) == 0 || len(value) > 64 || !lowerAlnum(value[0]) {
		return false
	}
	for index := 1; index < len(value); index++ {
		if !lowerAlnum(value[index]) && value[index] != '-' {
			return false
		}
	}
	return true
}

// Attach makes resolution lexically unreachable until the allowed verdict and
// its immutable authorized envelope have both been verified.
func Attach(frozen *freeze.Request, verdict authorize.Verdict, resolver Resolver) (*WireRequest, error) {
	if !verdict.Allowed {
		return nil, ErrNotAuthorized
	}
	envelope, ok := verdict.Authorized()
	if !ok || frozen == nil || envelope.FrozenCoreDigest != frozen.CoreDigest() {
		return nil, ErrEnvelopeMismatch
	}
	if !ValidReference(envelope.CredentialRef) || envelope.AuthProfile.HeaderName == "" || envelope.AuthProfile.Scheme != "bearer" {
		return nil, ErrEnvelopeMismatch
	}
	core := frozen.Core()
	for _, header := range core.Headers {
		if header.Name == envelope.AuthProfile.HeaderName {
			return nil, ErrEnvelopeMismatch
		}
	}
	if resolver == nil {
		return nil, ErrCredentialResolve
	}
	secret, err := resolver.resolve(envelope.CredentialRef)
	if err != nil || len(secret) == 0 {
		return nil, ErrCredentialResolve
	}
	auth := authorize.Header{
		Name:  envelope.AuthProfile.HeaderName,
		Value: "Bearer " + string(secret),
	}
	return &WireRequest{
		core: core,
		auth: auth,
		body: frozen.Body(),
	}, nil
}

// PrepareHTTPRequest applies the mandatory post-freeze mutation guard and
// constructs a single-use, non-rewindable HTTP request. The returned request
// is the only representation outside this package that can contain S-A.
func PrepareHTTPRequest(ctx context.Context, frozen *freeze.Request, wire *WireRequest) (*http.Request, error) {
	if wire == nil || frozen == nil {
		return nil, freeze.ErrFrozenMutation
	}
	if err := frozen.VerifyCandidate(wire.core, wire.body); err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, wire.core.Method, wire.core.Endpoint, &singleUseReader{value: wire.body})
	if err != nil {
		return nil, freeze.ErrFrozenMutation
	}
	bodyLen, err := strconv.ParseInt(wire.core.BodyLen, 10, 64)
	if err != nil || bodyLen != int64(len(wire.body)) {
		return nil, freeze.ErrFrozenMutation
	}
	request.ContentLength = bodyLen
	request.GetBody = nil
	request.Close = true
	for _, header := range wire.core.Headers {
		request.Header.Set(header.Name, header.Value)
	}
	request.Header.Set(wire.auth.Name, wire.auth.Value)
	return request, nil
}

func validateFileInfo(info fs.FileInfo, expectedUID int) error {
	if info == nil || !info.Mode().IsRegular() || info.Mode().Perm() != 0o600 || info.Mode()&(fs.ModeSetuid|fs.ModeSetgid|fs.ModeSticky) != 0 {
		return invalidFile("mode")
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != expectedUID {
		return invalidFile("owner")
	}
	return nil
}

func invalidFile(class string) error {
	return fmt.Errorf("%w: %s", ErrCredentialFileInvalid, class)
}

func lowerAlnum(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= '0' && value <= '9'
}

type singleUseReader struct {
	value  []byte
	offset int
}

func (reader *singleUseReader) Read(destination []byte) (int, error) {
	if reader.offset >= len(reader.value) {
		return 0, io.EOF
	}
	written := copy(destination, reader.value[reader.offset:])
	reader.offset += written
	return written, nil
}
