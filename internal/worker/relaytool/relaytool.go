// Package relaytool implements the worker-native conductor tools. It combines
// the m-2 form mapping with the transport-neutral conduct facade and exposes no
// channel, store, configuration, or credential handle.
package relaytool

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/seatclient/formschema"
	"github.com/jackli/frank/internal/worker/executor"
)

type Conductor interface {
	Relay(context.Context, string, json.RawMessage) (json.RawMessage, error)
	Describe(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error)
}

type Refresh struct {
	Phase string `json:"phase"`
	Tier  string `json:"tier"`
}

type RefreshSignal func(Refresh)

// Result keeps data and failures on one structured tool-result channel.
type Result struct {
	Data            json.RawMessage `json:"data,omitempty"`
	Error           string          `json:"error,omitempty"`
	Code            string          `json:"code,omitempty"`
	Help            []string        `json:"help,omitempty"`
	SchemaRefreshed bool            `json:"schema_refreshed,omitempty"`
}

type Registry struct {
	conductor Conductor
	onRefresh RefreshSignal

	mu          sync.Mutex
	submitForm  *fieldspec.Form
	formDigest  string
	schemaPhase string
	schemaTier  string
}

func New(conductor Conductor, onRefresh RefreshSignal) *Registry {
	return &Registry{conductor: conductor, onRefresh: onRefresh}
}

func (registry *Registry) Invoke(ctx context.Context, invocation executor.Invocation) (any, error) {
	if registry == nil || registry.conductor == nil {
		return failure("conductor unavailable", "conductor_unavailable", "reattach the worker seat and retry"), nil
	}
	arguments := json.RawMessage(invocation.Arguments)
	switch invocation.Identity.CanonicalToolName {
	case "relay.submit":
		return registry.submit(ctx, arguments), nil
	case "relay.project":
		if dispositions := formschema.ValidateProjectArguments(arguments); len(dispositions) > 0 {
			return schemaFailure(dispositions), nil
		}
		return registry.call(ctx, "relay.project", arguments, false), nil
	case "relay.read":
		if dispositions := formschema.ValidateReadArguments(arguments); len(dispositions) > 0 {
			return schemaFailure(dispositions), nil
		}
		return registry.call(ctx, "relay.read", arguments, false), nil
	default:
		return failure("unknown relay tool name", "unknown_tool", "use relay.submit, relay.project, or relay.read"), nil
	}
}

func (registry *Registry) submit(ctx context.Context, encoded json.RawMessage) Result {
	arguments, err := formschema.ParseSubmitArguments(encoded)
	if err != nil {
		var mappingError *formschema.MappingError
		code := "schema_invalid"
		if formschema.AsMappingError(err, &mappingError) {
			code = string(mappingError.ID)
		}
		return failure("invalid relay.submit arguments", code, "correct the named argument and retry")
	}
	phase, tier := formschema.DeclaredPhaseTier(arguments)
	refreshed := false
	form, digest, cached := registry.cachedForm(phase, tier)
	if !cached || len(formschema.ValidateSubmitArguments(form, digest, encoded)) > 0 {
		if err := registry.refresh(ctx, phase, tier); err != nil {
			return failure("conductor metadata unavailable", "transport_error", "reattach and fetch the submit schema again")
		}
		refreshed = true
		form, digest, cached = registry.cachedForm(phase, tier)
		if !cached {
			return failure("submit schema unavailable", "protocol_error", "report the missing Describe submit_schema")
		}
		if dispositions := formschema.ValidateSubmitArguments(form, digest, encoded); len(dispositions) > 0 {
			result := schemaFailure(dispositions)
			result.SchemaRefreshed = true
			return result
		}
	}
	payload, err := formschema.SubmitPayloadFromArguments(encoded)
	if err != nil {
		return failure("invalid relay.submit arguments", "schema_invalid", "correct the named argument and retry")
	}
	result, err := registry.conductor.Relay(ctx, "relay.submit", payload)
	if err != nil {
		return failure("conductor call failed", "transport_error", "rediscover the durable mailbox before retrying")
	}
	if submitRejected(result) {
		if err := registry.refresh(ctx, phase, tier); err == nil {
			refreshed = true
		}
		if formschema.SubmitNeedsReRender(result) {
			result = formschema.ReRenderResult(result)
		}
	}
	return Result{Data: append(json.RawMessage(nil), result...), SchemaRefreshed: refreshed}
}

func (registry *Registry) call(ctx context.Context, name string, arguments json.RawMessage, refreshed bool) Result {
	result, err := registry.conductor.Relay(ctx, name, arguments)
	if err != nil {
		return failure("conductor call failed", "transport_error", "reattach the worker seat and retry")
	}
	return Result{Data: append(json.RawMessage(nil), result...), SchemaRefreshed: refreshed}
}

func (registry *Registry) refresh(ctx context.Context, phase, tier string) error {
	description, err := registry.conductor.Describe(ctx, channel.DescribeRequest{Phase: phase, Tier: tier})
	if err != nil {
		return err
	}
	if description.SubmitSchema == nil {
		return errMissingSubmitSchema
	}
	form := *description.SubmitSchema
	registry.mu.Lock()
	registry.submitForm = &form
	registry.formDigest = description.FormDigest
	registry.schemaPhase = phase
	registry.schemaTier = tier
	registry.mu.Unlock()
	if registry.onRefresh != nil {
		registry.onRefresh(Refresh{Phase: phase, Tier: tier})
	}
	return nil
}

func (registry *Registry) cachedForm(phase, tier string) (fieldspec.Form, string, bool) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if registry.submitForm == nil || registry.schemaPhase != phase || registry.schemaTier != tier {
		return fieldspec.Form{}, "", false
	}
	return *registry.submitForm, registry.formDigest, true
}

func failure(message, code string, help ...string) Result {
	return Result{Error: message, Code: code, Help: append([]string(nil), help...)}
}

func schemaFailure(dispositions []formschema.Disposition) Result {
	encoded, _ := json.Marshal(dispositions)
	return failure("arguments do not match the current tool schema: "+string(encoded), "schema_invalid", "fetch the current tool schema and correct the named members")
}

func submitRejected(encoded json.RawMessage) bool {
	var outcome struct {
		State string `json:"state"`
	}
	return json.Unmarshal(encoded, &outcome) == nil && outcome.State == "rejected"
}
