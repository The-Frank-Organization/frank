// Package conduct is the narrow channel-using conductor client facade. It
// exposes verbs and transport metadata, never raw store or configuration paths.
package conduct

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackli/frank/internal/channel"
)

var ErrUnknownVerb = errors.New("conduct: unknown canonical relay verb")

type Transport interface {
	InvokeWire(context.Context, string, json.RawMessage) (json.RawMessage, error)
	DescribeWire(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error)
	ReceivePush(context.Context) ([]byte, error)
	Shutdown() error
}

type Client struct {
	transport Transport
}

func New(transport Transport) (*Client, error) {
	if transport == nil {
		return nil, errors.New("conduct: transport is absent")
	}
	return &Client{transport: transport}, nil
}

// FromAuthenticated adapts an already-authenticated channel. Authentication
// and credential resolution remain entirely outside this package.
func FromAuthenticated(authenticated *channel.Client) (*Client, error) {
	if authenticated == nil {
		return nil, errors.New("conduct: authenticated channel is absent")
	}
	return New(authenticatedChannel{Client: authenticated})
}

func (client *Client) Call(ctx context.Context, canonicalName string, arguments json.RawMessage) (json.RawMessage, error) {
	return client.relay(ctx, canonicalName, arguments)
}

// Relay is the consumer-facing spelling used by both native and MCP adapters;
// Call remains the locked m-7 seam exposed by the facade.
func (client *Client) Relay(ctx context.Context, canonicalName string, arguments json.RawMessage) (json.RawMessage, error) {
	return client.relay(ctx, canonicalName, arguments)
}

func (client *Client) relay(ctx context.Context, canonicalName string, arguments json.RawMessage) (json.RawMessage, error) {
	var wireName string
	switch canonicalName {
	case "relay.submit":
		wireName = "submit"
	case "relay.project":
		wireName = "project"
	case "relay.read":
		wireName = "read"
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownVerb, canonicalName)
	}
	return client.transport.InvokeWire(ctx, wireName, arguments)
}

func (client *Client) Describe(ctx context.Context, request channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return client.transport.DescribeWire(ctx, request)
}

func (client *Client) NextPush(ctx context.Context) ([]byte, error) {
	return client.transport.ReceivePush(ctx)
}

func (client *Client) Close() error { return client.transport.Shutdown() }

type authenticatedChannel struct{ *channel.Client }

func (transport authenticatedChannel) InvokeWire(ctx context.Context, name string, arguments json.RawMessage) (json.RawMessage, error) {
	return transport.Client.Call(ctx, name, arguments)
}

func (transport authenticatedChannel) DescribeWire(ctx context.Context, request channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return transport.Client.DescribeTools(ctx, request)
}

func (transport authenticatedChannel) ReceivePush(ctx context.Context) ([]byte, error) {
	return transport.Client.NextPush(ctx)
}

func (transport authenticatedChannel) Shutdown() error { return transport.Client.Close() }
