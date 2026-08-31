package freeze

import (
	"bytes"
	"errors"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/connector/authorize"
	"github.com/The-Frank-Organization/frank/internal/connector/catalog"
	connectorrequest "github.com/The-Frank-Organization/frank/internal/connector/request"
	"github.com/The-Frank-Organization/frank/internal/connector/translate"
)

func TestFreezeProducesExactFiveMemberCoreAndDigests(t *testing.T) {
	t.Parallel()

	frozen, err := Freeze(translatedBody(), lane(), "build-1")
	if err != nil {
		t.Fatalf("Freeze() error = %v", err)
	}
	wantCore := `{"body_len":"2","body_sha256":"44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a","endpoint":"https://api.openai.com/v1/responses","headers":[["content-type","application/json"],["user-agent","frank-connector/build-1"]],"method":"POST"}`
	if string(frozen.CoreBytes()) != wantCore {
		t.Fatalf("CoreBytes()\n got: %s\nwant: %s", frozen.CoreBytes(), wantCore)
	}
	if frozen.CoreDigest() != "0642d1d5db3bc3b73b46af7600b93a6f7389b107b75c090c86e77fd23455f369" {
		t.Fatalf("CoreDigest() = %q", frozen.CoreDigest())
	}
	if frozen.LoweredToolsDigest() != "4f53cda18c2baa0c0354bb5f9a3ecbe5ed12ab4d8e11ba873c2f11161202b945" {
		t.Fatalf("LoweredToolsDigest() = %q", frozen.LoweredToolsDigest())
	}
	core := frozen.Core()
	if core.BodyLen != "2" || !core.BodyIsDesignated {
		t.Fatalf("Core() = %+v", core)
	}
}

func TestFreezeRefusesDuplicateLowercaseHeaderNames(t *testing.T) {
	t.Parallel()

	_, err := freezeWithHeaders(translatedBody(), lane(), []authorize.Header{
		{Name: "content-type", Value: "application/json"},
		{Name: "content-type", Value: "application/other"},
	})
	if !errors.Is(err, connectorrequest.InternalIntegrityFault) {
		t.Fatalf("freezeWithHeaders() error = %v, want internal_integrity_fault", err)
	}
}

func TestFreezeCopiesInputsAndReturnsCopies(t *testing.T) {
	t.Parallel()

	translated := translatedBody()
	frozen, err := Freeze(translated, lane(), "build-1")
	if err != nil {
		t.Fatalf("Freeze() error = %v", err)
	}
	translated.Body[0] = '['
	body := frozen.Body()
	body[0] = '['
	core := frozen.Core()
	core.Headers[0].Value = "mutated"

	if string(frozen.Body()) != `{}` || frozen.Core().Headers[0].Value != "application/json" {
		t.Fatal("frozen state changed through an input or copy-on-read mutation")
	}
}

func TestFrozenCoreDigestBindsMethodBodyAndNonAuthHeaders(t *testing.T) {
	t.Parallel()

	baseline := mustFreeze(t, translatedBody(), lane(), defaultHeaders("build-1"))

	differentMethod := lane()
	differentMethod.Method = "PUT"
	method := mustFreeze(t, translatedBody(), differentMethod, defaultHeaders("build-1"))

	differentBody := translatedBody()
	differentBody.Body = []byte(`{"changed":true}`)
	body := mustFreeze(t, differentBody, lane(), defaultHeaders("build-1"))

	headers := append(defaultHeaders("build-1"), authorize.Header{Name: "x-trace", Value: "one"})
	header := mustFreeze(t, translatedBody(), lane(), headers)

	for name, digest := range map[string]string{
		"method": method.CoreDigest(),
		"body":   body.CoreDigest(),
		"header": header.CoreDigest(),
	} {
		if digest == baseline.CoreDigest() {
			t.Fatalf("%s mutation did not change frozen core digest", name)
		}
	}
}

func TestVerifyCandidateDetectsEveryPostFreezeMutation(t *testing.T) {
	t.Parallel()

	frozen, err := Freeze(translatedBody(), lane(), "build-1")
	if err != nil {
		t.Fatalf("Freeze() error = %v", err)
	}
	if err := frozen.VerifyCandidate(frozen.Core(), frozen.Body()); err != nil {
		t.Fatalf("VerifyCandidate(valid) error = %v", err)
	}

	tests := []struct {
		name   string
		mutate func(*authorize.Core, *[]byte)
	}{
		{name: "method", mutate: func(core *authorize.Core, _ *[]byte) { core.Method = "PUT" }},
		{name: "endpoint", mutate: func(core *authorize.Core, _ *[]byte) { core.Endpoint = "https://other.example/v1/responses" }},
		{name: "header", mutate: func(core *authorize.Core, _ *[]byte) { core.Headers[0].Value = "text/plain" }},
		{name: "body", mutate: func(_ *authorize.Core, body *[]byte) { *body = []byte(`[]`) }},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			core := frozen.Core()
			body := frozen.Body()
			test.mutate(&core, &body)
			if err := frozen.VerifyCandidate(core, body); !errors.Is(err, ErrFrozenMutation) {
				t.Fatalf("VerifyCandidate() error = %v, want ErrFrozenMutation", err)
			}
		})
	}
}

func TestDefaultFrozenHeadersContainOnlyPinnedProfileFields(t *testing.T) {
	t.Parallel()

	headers := defaultHeaders("build-1")
	want := []authorize.Header{
		{Name: "content-type", Value: "application/json"},
		{Name: "user-agent", Value: "frank-connector/build-1"},
	}
	if len(headers) != len(want) {
		t.Fatalf("defaultHeaders() = %+v", headers)
	}
	for index := range want {
		if headers[index] != want[index] {
			t.Fatalf("defaultHeaders()[%d] = %+v, want %+v", index, headers[index], want[index])
		}
	}
	for _, forbidden := range []string{"authorization", "accept-encoding", "content-length", "connection", "host"} {
		for _, header := range headers {
			if header.Name == forbidden {
				t.Fatalf("derived, suppressed, or auth header %q entered frozen set", forbidden)
			}
		}
	}
}

func translatedBody() translate.Result {
	return translate.Result{
		Body:           []byte(`{}`),
		LoweredTools:   []byte(`[]`),
		ProfileVersion: translate.OpenAIResponsesProfileVersion,
	}
}

func lane() catalog.Lane {
	return catalog.Lane{Method: "POST", Endpoint: "https://api.openai.com/v1/responses"}
}

func mustFreeze(t *testing.T, translated translate.Result, lane catalog.Lane, headers []authorize.Header) *Request {
	t.Helper()
	frozen, err := freezeWithHeaders(translated, lane, headers)
	if err != nil {
		t.Fatalf("freezeWithHeaders() error = %v", err)
	}
	return frozen
}

func TestFrozenBodyDoesNotAliasCoreBytes(t *testing.T) {
	t.Parallel()

	frozen, err := Freeze(translatedBody(), lane(), "build-1")
	if err != nil {
		t.Fatalf("Freeze() error = %v", err)
	}
	if bytes.Equal(frozen.Body(), frozen.CoreBytes()) {
		t.Fatal("body unexpectedly aliases the core representation")
	}
}
