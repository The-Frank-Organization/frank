package relaytool

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/seatclient/conduct"
	"github.com/The-Frank-Organization/frank/internal/worker/executor"
)

type fakeTransport struct {
	forms         []channel.DescriptionResponse
	describeCalls int
	calls         []wireCall
	results       []json.RawMessage
	describeErr   error
	callErr       error
}

type wireCall struct {
	name string
	args string
}

func (transport *fakeTransport) InvokeWire(_ context.Context, name string, args json.RawMessage) (json.RawMessage, error) {
	transport.calls = append(transport.calls, wireCall{name: name, args: string(args)})
	if transport.callErr != nil {
		return nil, transport.callErr
	}
	if len(transport.results) == 0 {
		return json.RawMessage(`{"state":"accepted"}`), nil
	}
	result := transport.results[0]
	transport.results = transport.results[1:]
	return result, nil
}

func (transport *fakeTransport) DescribeWire(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error) {
	transport.describeCalls++
	if transport.describeErr != nil {
		return channel.DescriptionResponse{}, transport.describeErr
	}
	if len(transport.forms) == 0 {
		return channel.DescriptionResponse{}, nil
	}
	description := transport.forms[0]
	if len(transport.forms) > 1 {
		transport.forms = transport.forms[1:]
	}
	return description, nil
}

func (*fakeTransport) ReceivePush(context.Context) ([]byte, error) { return nil, errors.New("no push") }
func (*fakeTransport) Shutdown() error                             { return nil }

func TestF1RefreshBeforeRejectAllowsFreshExpansion(t *testing.T) {
	oldForm := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"PHASE": {Type: "enum", Options: []string{"SITREP"}},
	}}
	newForm := testForm(fieldspec.Field{Type: "id_ref", Options: []string{"parent-a", "parent-b"}})
	newForm.Fields["PHASE"] = fieldspec.Field{Type: "enum", Options: []string{"SITREP"}}
	transport := &fakeTransport{forms: []channel.DescriptionResponse{
		{SubmitSchema: &oldForm, FormDigest: "same-digest"},
		{SubmitSchema: &newForm, FormDigest: "same-digest"},
	}}
	client, _ := conduct.New(transport)
	var refreshes []Refresh
	registry := New(client, func(refresh Refresh) { refreshes = append(refreshes, refresh) })

	first := invoke(t, registry, `{"headers":{"PHASE":"SITREP"},"form_digest":"same-digest"}`)
	if first.Error != "" || len(transport.calls) != 1 {
		t.Fatalf("initial call = %+v, wire=%#v", first, transport.calls)
	}
	second := invoke(t, registry, `{"headers":{"PHASE":"SITREP","PARENT_DISPATCH_ID":"parent-b"},"form_digest":"same-digest"}`)
	if second.Error != "" || !second.SchemaRefreshed || len(transport.calls) != 2 {
		t.Fatalf("expanded call = %+v, wire=%#v", second, transport.calls)
	}
	if transport.describeCalls != 2 || len(refreshes) != 2 {
		t.Fatalf("describe=%d refreshes=%#v", transport.describeCalls, refreshes)
	}
}

func TestF1DescribeFailureIsTransportErrorAndNoCall(t *testing.T) {
	transport := &fakeTransport{describeErr: errors.New("credential-like transport detail")}
	client, _ := conduct.New(transport)
	result := invoke(t, New(client, nil), `{"headers":{},"form_digest":"digest"}`)
	if result.Code != "transport_error" || len(transport.calls) != 0 || result.Error == "credential-like transport detail" {
		t.Fatalf("result=%+v calls=%#v", result, transport.calls)
	}
}

func TestF2EveryRejectedSubmitRefreshesAndReRenderNormalizes(t *testing.T) {
	form := testForm(fieldspec.Field{Type: "id_ref", Options: []string{"parent-a"}})
	for _, test := range []struct {
		name       string
		outcome    string
		reRendered bool
	}{
		{"volatile contraction", `{"state":"rejected","detail":"PARENT_DISPATCH_ID: enum"}`, false},
		{"stale digest", `{"state":"rejected","relay_id":"r1","detail":"form_digest:re-render"}`, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			transport := &fakeTransport{
				forms:   []channel.DescriptionResponse{{SubmitSchema: &form, FormDigest: "digest"}, {SubmitSchema: &form, FormDigest: "digest"}},
				results: []json.RawMessage{json.RawMessage(test.outcome)},
			}
			client, _ := conduct.New(transport)
			result := invoke(t, New(client, nil), `{"headers":{"PARENT_DISPATCH_ID":"parent-a"},"form_digest":"digest"}`)
			if !result.SchemaRefreshed || transport.describeCalls != 2 {
				t.Fatalf("result=%+v describe=%d", result, transport.describeCalls)
			}
			gotReRender := contains(result.Data, `"class":"re-render"`)
			if gotReRender != test.reRendered {
				t.Fatalf("data=%s re-render=%v want %v", result.Data, gotReRender, test.reRendered)
			}
		})
	}
}

func TestStrictValidationAndCanonicalFacadeBytes(t *testing.T) {
	form := testForm(fieldspec.Field{Type: "id_ref", Options: []string{"parent-a"}})
	transport := &fakeTransport{forms: []channel.DescriptionResponse{{SubmitSchema: &form, FormDigest: "digest"}}}
	client, _ := conduct.New(transport)
	registry := New(client, nil)
	invalid := invoke(t, registry, `{"headers":{},"form_digest":"digest","Body":"case variant"}`)
	if invalid.Code != "P-1.b" || len(transport.calls) != 0 {
		t.Fatalf("invalid=%+v calls=%#v", invalid, transport.calls)
	}
	valid := invoke(t, registry, `{"headers":{"PARENT_DISPATCH_ID":"parent-a"},"to":"m-9.planner","form_digest":"digest"}`)
	if valid.Error != "" || len(transport.calls) != 1 || transport.calls[0].name != "submit" {
		t.Fatalf("valid=%+v calls=%#v", valid, transport.calls)
	}
	var payload fieldspec.SubmitPayload
	if err := json.Unmarshal([]byte(transport.calls[0].args), &payload); err != nil || payload.Envelope.To != "m-9.planner" {
		t.Fatalf("payload=%s err=%v", transport.calls[0].args, err)
	}
}

func TestStaticRelayArgumentsFailBeforeFacadeCall(t *testing.T) {
	transport := &fakeTransport{}
	client, _ := conduct.New(transport)
	registry := New(client, nil)
	for _, invocation := range []executor.Invocation{
		{Identity: executor.Identity{CanonicalToolName: "relay.project"}, Arguments: []byte(`{"view":"everything"}`)},
		{Identity: executor.Identity{CanonicalToolName: "relay.read"}, Arguments: []byte(`{}`)},
	} {
		value, err := registry.Invoke(context.Background(), invocation)
		if err != nil || value.(Result).Code != "schema_invalid" {
			t.Fatalf("value=%+v err=%v", value, err)
		}
	}
	if len(transport.calls) != 0 {
		t.Fatalf("invalid static calls reached facade: %#v", transport.calls)
	}
}

func TestUnknownToolReturnsAXIShapedFailure(t *testing.T) {
	value, err := New(&fakeConductor{}, nil).Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "submit"}})
	result := value.(Result)
	if err != nil || result.Error == "" || result.Code != "unknown_tool" || len(result.Help) == 0 || result.Data != nil {
		t.Fatalf("result=%+v err=%v", result, err)
	}
}

type fakeConductor struct{}

func (*fakeConductor) Relay(context.Context, string, json.RawMessage) (json.RawMessage, error) {
	return nil, nil
}
func (*fakeConductor) Describe(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return channel.DescriptionResponse{}, nil
}

func invoke(t *testing.T, registry *Registry, arguments string) Result {
	t.Helper()
	value, err := registry.Invoke(context.Background(), executor.Invocation{
		Identity:  executor.Identity{CanonicalToolName: "relay.submit"},
		Arguments: []byte(arguments),
	})
	if err != nil {
		t.Fatal(err)
	}
	return value.(Result)
}

func testForm(parent fieldspec.Field) fieldspec.Form {
	parent.ConductorVolatile = true
	parent.DigestExempt = true
	return fieldspec.Form{Fields: map[string]fieldspec.Field{"PARENT_DISPATCH_ID": parent}}
}

func contains(data []byte, fragment string) bool {
	return json.Valid(data) && strings.Contains(string(data), fragment)
}
