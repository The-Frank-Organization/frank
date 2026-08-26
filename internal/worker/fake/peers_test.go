package fake

import (
	"context"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/worker/executor"
	workerruntime "github.com/jackli/frank/internal/worker/runtime"
)

func TestM10TicketIsOneShotAndWakeIsDuplicateSafe(t *testing.T) {
	peer := NewM10(workerruntime.Assignment{})
	identity := executor.FullIdentity{RunID: "r", TurnID: "t", ToolCallID: "c", Identity: executor.Identity{CanonicalToolName: "read", CanonicalArgsDigest: strings.Repeat("a", 64), TurnEpoch: "1"}}
	authorized, err := peer.Authorize(context.Background(), executor.AuthorizeRequest{Identity: identity})
	if err != nil {
		t.Fatal(err)
	}
	request := executor.ConsumeRequest{TicketID: authorized.TicketID, TurnEpoch: "1", CanonicalToolName: "read", CanonicalArgsDigest: strings.Repeat("a", 64)}
	first, _ := peer.Consume(context.Background(), request)
	second, _ := peer.Consume(context.Background(), request)
	if first.Code != executor.ConsumeOK || second.Code != executor.ConsumeDuplicate {
		t.Fatalf("consume codes = %q, %q", first.Code, second.Code)
	}
	_ = peer.WakeForward(context.Background(), "relay")
	_ = peer.WakeForward(context.Background(), "relay")
	if len(peer.Wakes) != 1 {
		t.Fatalf("wakes = %v", peer.Wakes)
	}
}
