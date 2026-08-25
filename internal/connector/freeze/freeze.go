// Package freeze owns the immutable m-8 request boundary between provider
// translation and authorization.
package freeze

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jackli/frank/internal/connector/authorize"
	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/jcs"
	connectorrequest "github.com/jackli/frank/internal/connector/request"
	"github.com/jackli/frank/internal/connector/translate"
)

var ErrFrozenMutation = errors.New("freeze: frozen request mutation")

// Request retains unexported, copy-on-read state. Later pipeline stages receive
// copies and must pass them back through VerifyCandidate immediately pre-wire.
type Request struct {
	core               authorize.Core
	coreBytes          []byte
	coreDigest         string
	body               []byte
	loweredToolsDigest string
}

func Freeze(translated translate.Result, lane catalog.Lane, build string) (*Request, error) {
	if !validBuild(build) {
		return nil, fmt.Errorf("%w: invalid connector build identity", connectorrequest.InternalIntegrityFault)
	}
	return freezeWithHeaders(translated, lane, defaultHeaders(build))
}

func (request *Request) Core() authorize.Core {
	core := request.core
	core.Headers = append([]authorize.Header(nil), request.core.Headers...)
	return core
}

func (request *Request) CoreBytes() []byte {
	return append([]byte(nil), request.coreBytes...)
}

func (request *Request) CoreDigest() string { return request.coreDigest }

func (request *Request) Body() []byte { return append([]byte(nil), request.body...) }

func (request *Request) LoweredToolsDigest() string { return request.loweredToolsDigest }

// VerifyCandidate applies both parts of the step-5 mutation guard: the exact
// five-member core digest and the actual body bytes bound by that core.
func (request *Request) VerifyCandidate(core authorize.Core, body []byte) error {
	if request == nil || !core.BodyIsDesignated {
		return ErrFrozenMutation
	}
	encoded, err := encodeCore(core)
	if err != nil || digest(encoded) != request.coreDigest {
		return ErrFrozenMutation
	}
	bodyDigest := sha256.Sum256(body)
	if hex.EncodeToString(bodyDigest[:]) != core.BodySHA256 || strconv.FormatUint(uint64(len(body)), 10) != core.BodyLen {
		return ErrFrozenMutation
	}
	return nil
}

func freezeWithHeaders(translated translate.Result, lane catalog.Lane, headers []authorize.Header) (*Request, error) {
	ownedHeaders := append([]authorize.Header(nil), headers...)
	sort.Slice(ownedHeaders, func(left, right int) bool { return ownedHeaders[left].Name < ownedHeaders[right].Name })
	if err := validateHeaders(ownedHeaders); err != nil {
		return nil, fmt.Errorf("%w: %v", connectorrequest.InternalIntegrityFault, err)
	}
	body := append([]byte(nil), translated.Body...)
	bodyHash := sha256.Sum256(body)
	core := authorize.Core{
		Method:           lane.Method,
		Endpoint:         lane.Endpoint,
		Headers:          ownedHeaders,
		BodySHA256:       hex.EncodeToString(bodyHash[:]),
		BodyLen:          strconv.FormatUint(uint64(len(body)), 10),
		BodyIsDesignated: true,
	}
	coreBytes, err := encodeCore(core)
	if err != nil {
		return nil, fmt.Errorf("%w: encode core: %v", connectorrequest.InternalIntegrityFault, err)
	}
	canonicalTools, err := jcs.Canonicalize(translated.LoweredTools)
	if err != nil || !bytes.Equal(canonicalTools, translated.LoweredTools) || len(canonicalTools) == 0 || canonicalTools[0] != '[' {
		return nil, fmt.Errorf("%w: invalid lowered tools", connectorrequest.InternalIntegrityFault)
	}
	return &Request{
		core:               core,
		coreBytes:          coreBytes,
		coreDigest:         digest(coreBytes),
		body:               body,
		loweredToolsDigest: digest(canonicalTools),
	}, nil
}

func defaultHeaders(build string) []authorize.Header {
	return []authorize.Header{
		{Name: "content-type", Value: "application/json"},
		{Name: "user-agent", Value: "frank-connector/" + build},
	}
}

func encodeCore(core authorize.Core) ([]byte, error) {
	if err := validateHeaders(core.Headers); err != nil {
		return nil, err
	}
	headerPairs := make([][2]string, len(core.Headers))
	for index, header := range core.Headers {
		headerPairs[index] = [2]string{header.Name, header.Value}
	}
	wire := struct {
		BodyLen    string      `json:"body_len"`
		BodySHA256 string      `json:"body_sha256"`
		Endpoint   string      `json:"endpoint"`
		Headers    [][2]string `json:"headers"`
		Method     string      `json:"method"`
	}{
		BodyLen:    core.BodyLen,
		BodySHA256: core.BodySHA256,
		Endpoint:   core.Endpoint,
		Headers:    headerPairs,
		Method:     core.Method,
	}
	raw, err := json.Marshal(wire)
	if err != nil {
		return nil, err
	}
	return jcs.Canonicalize(raw)
}

func validateHeaders(headers []authorize.Header) error {
	previous := ""
	for index, header := range headers {
		if !validLowerHeaderName(header.Name) || !validHeaderValue(header.Value) {
			return fmt.Errorf("invalid frozen header")
		}
		if index > 0 && previous >= header.Name {
			return fmt.Errorf("frozen header names are not sorted and unique")
		}
		previous = header.Name
	}
	return nil
}

func validLowerHeaderName(value string) bool {
	if value == "" || value != strings.ToLower(value) {
		return false
	}
	for index := range value {
		b := value[index]
		if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || strings.ContainsRune("!#$%&'*+-.^_`|~", rune(b)) {
			continue
		}
		return false
	}
	return true
}

func validHeaderValue(value string) bool {
	for index := range value {
		if value[index] == '\r' || value[index] == '\n' || value[index] == 0 {
			return false
		}
	}
	return true
}

func validBuild(value string) bool {
	if value == "" {
		return false
	}
	for index := range value {
		b := value[index]
		if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' || strings.ContainsRune("._+-", rune(b)) {
			continue
		}
		return false
	}
	return true
}

func digest(value []byte) string {
	sum := sha256.Sum256(value)
	return hex.EncodeToString(sum[:])
}
