package endpoint

import "testing"

func TestValidateAcceptsOnlyThePinnedUniqueSpellingGrammar(t *testing.T) {
	t.Parallel()

	for _, value := range []string{
		"https://api.openai.com/v1/responses",
		"https://gateway.internal:8443/v1/chat",
		"https://xn--bcher-kva.example/v1",
		"https://example.com/",
	} {
		if err := Validate(value); err != nil {
			t.Errorf("Validate(%q) error = %v", value, err)
		}
	}

	for _, value := range []string{
		"http://api.openai.com/v1",
		"https://API.openai.com/v1",
		"https://api.openai.com.:443/v1",
		"https://api.openai.com:443/v1",
		"https://[2001:db8::1]/v1",
		"https://192.168.0.1/v1",
		"https://api.openai.com/v1//responses",
		"https://api.openai.com/v1/%2e%2e/admin",
		"https://api.openai.com/v1/",
		"https://user@api.openai.com/v1",
		"https://api.openai.com/v1?x=1",
		"https://api.openai.com:0443/v1",
		"https://api.openai.com:0/v1",
		"https://api.openai.com:65536/v1",
		"https://api_openai.com/v1",
	} {
		if err := Validate(value); err == nil {
			t.Errorf("Validate(%q) accepted a forbidden endpoint spelling", value)
		}
	}
}
