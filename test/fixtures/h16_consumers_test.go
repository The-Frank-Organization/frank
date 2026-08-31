package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/observe"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

func TestH16StateOnlyConsumersFailClosedOnNonComplete(t *testing.T) {
	pending := engine.Outcome{DecisionState: record.Accepted, PostCommitState: "pending", RelayID: "pending-source"}
	submitter := &h16ConsumerSubmitter{outcome: pending}
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}

	expiry, err := engine.NewExpiryPrompter(st, submitter)
	if err != nil {
		t.Fatalf("new expiry prompter: %v", err)
	}
	expiryDecision := expiry.Prompt(context.Background(), observe.ExpiryRequest{
		Selection:  observe.Selection{Seat: "s12.implementer", CheckID: "read-file", ClaimRef: "expiry", CandidateDigest: "candidate"},
		SoftExpiry: time.Second, HardCeiling: 2 * time.Second,
	})
	if expiryDecision.Action != observe.ExpiryKill {
		t.Fatalf("pending expiry outcome did not fail closed: %+v", expiryDecision)
	}

	approval, err := engine.NewApprovalPrompter(st, submitter)
	if err != nil {
		t.Fatalf("new approval prompter: %v", err)
	}
	approvalDecision := approval.Prompt(context.Background(), observe.ApprovalRequest{Selection: observe.Selection{
		Seat: "s12.implementer", CheckID: "run-suite", ClaimRef: "approval", CandidateDigest: "candidate",
	}})
	if approvalDecision.Allowed || approvalDecision.Scope != observe.ApprovalDenied {
		t.Fatalf("pending approval outcome did not fail closed: %+v", approvalDecision)
	}

	scheduler, err := engine.NewResummonScheduler(st, submitter, func() *tables.T { return tables.New() })
	if err != nil {
		t.Fatalf("new resummon scheduler: %v", err)
	}
	input := engine.ResummonInput{
		Seat: "s12.implementer", DecisionID: "gate-pending", CadenceSlot: "g4-no-response-1",
		Reason: engine.ResummonNoResponse, SummonChannel: engine.SummonLocal,
	}
	before := len(submitter.commands)
	for attempt := 0; attempt < 2; attempt++ {
		if _, err := scheduler.Emit(context.Background(), input); err == nil || !strings.Contains(err.Error(), "resummon emit rejected") {
			t.Fatalf("pending resummon attempt %d err=%v, want retry-later refusal", attempt, err)
		}
	}
	commands := submitter.commands[before:]
	if len(commands) != 2 || commands[0].ContentHash == "" || commands[0].ContentHash != commands[1].ContentHash {
		t.Fatalf("resummon retry content hashes=%q/%q, want stable dedup key", commands[0].ContentHash, commands[1].ContentHash)
	}
}

func TestH16NativeAndMCPForwardBothOutcomeDimensionsByteTransparently(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open seats: %v", err)
	}
	cred, err := mgr.Mint("s12.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("mint seat: %v", err)
	}
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{}}
	const digest = "h16-consumer-form"
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-h16-consumers-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.Marshal(channel.DescriptionResponse{Tools: []string{"submit", "project", "read"}, SubmitSchema: &form, FormDigest: digest})
			},
			Submit: func(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
				var payload fieldspec.SubmitPayload
				if err := json.Unmarshal(args, &payload); err != nil {
					return nil, err
				}
				return h16ConsumerOutcomeBytes(payload.Body), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("serve conductor: %v", err)
	}
	defer func() { _ = srv.Close() }()

	native, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("dial native client: %v", err)
	}
	for _, state := range []string{"pending", "failed", "unknown"} {
		result, err := native.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: record.Record{Body: state}, FormDigest: digest}))
		if err != nil {
			t.Fatalf("native %s call: %v", state, err)
		}
		if want := h16ConsumerOutcomeBytes(state); !bytes.Equal(result, want) {
			t.Fatalf("native %s result=%s, want exact %s", state, result, want)
		}
	}
	_ = native.Close()

	mcpInput := `{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}` + "\n"
	for i, state := range []string{"pending", "failed", "unknown"} {
		mcpInput += mcpSubmitInput(i+2, map[string]any{"headers": map[string]string{}, "body": state, "form_digest": digest})
	}
	stdout, stderr := runFrankMCP(t, ctx, buildFrankMCP(t, ctx), sock, cred.Value, mcpInput)
	if stderr != "" {
		t.Fatalf("MCP stderr=%s", stderr)
	}
	lines := strings.Split(strings.TrimSpace(stdout), "\n")
	if len(lines) != 4 {
		t.Fatalf("MCP output lines=%d, want tools/list + three calls: %s", len(lines), stdout)
	}
	for i, state := range []string{"pending", "failed", "unknown"} {
		var envelope struct {
			Result struct {
				Content []struct {
					Text string `json:"text"`
				} `json:"content"`
			} `json:"result"`
		}
		if err := json.Unmarshal([]byte(lines[i+1]), &envelope); err != nil {
			t.Fatalf("decode MCP %s line: %v", state, err)
		}
		if len(envelope.Result.Content) != 1 || envelope.Result.Content[0].Text != string(h16ConsumerOutcomeBytes(state)) {
			t.Fatalf("MCP %s forwarded=%+v, want exact %s", state, envelope.Result.Content, h16ConsumerOutcomeBytes(state))
		}
	}
}

func TestH16ConsumerCensusRemainsStateOnlyAndHealNudgeDropped(t *testing.T) {
	read := func(path string) []byte {
		t.Helper()
		data, err := os.ReadFile(filepath.Join("..", "..", path))
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		return data
	}
	mainBytes := read("cmd/frank/main.go")
	prompterBytes := read("internal/engine/prompter.go")
	resummonBytes := read("internal/engine/resummon.go")
	mcpBytes := read("cmd/frank-mcp/mcp.go")
	loopBytes := read("internal/engine/loop.go")

	if !bytes.Contains(mainBytes, []byte("out.State != record.Accepted")) || bytes.Count(mainBytes, []byte("deliveryNudgeFrame(out.RelayID)")) != 1 {
		t.Fatal("delivery nudge no longer gates on complete legacy state or gained a heal-refire site")
	}
	for path, source := range map[string][]byte{"prompter.go": prompterBytes, "resummon.go": resummonBytes} {
		if !bytes.Contains(source, []byte("outcome.State != record.Accepted")) {
			t.Fatalf("%s no longer fails closed on the state-only projection", path)
		}
		if bytes.Contains(source, []byte("outcome.DecisionState")) || bytes.Contains(source, []byte("outcome.PostCommitState")) {
			t.Fatalf("%s was migrated to two-dimension reads outside this lane", path)
		}
	}
	reRenderStart := bytes.Index(mcpBytes, []byte("func submitNeedsReRender"))
	reRenderEnd := bytes.Index(mcpBytes[reRenderStart:], []byte("func containsReRender"))
	if reRenderStart < 0 || reRenderEnd < 0 {
		t.Fatal("MCP re-render consumer missing")
	}
	reRender := mcpBytes[reRenderStart : reRenderStart+reRenderEnd]
	if !bytes.Contains(reRender, []byte(`outcome.State != "rejected"`)) || bytes.Contains(reRender, []byte("DecisionState")) || bytes.Contains(reRender, []byte("PostCommitState")) {
		t.Fatalf("MCP re-render consumer changed two-dimension boundary: %s", reRender)
	}
	if !bytes.Contains(mcpBytes, []byte("textToolResult(string(result)")) {
		t.Fatal("MCP raw forwarding path changed")
	}
	if !bytes.Contains(loopBytes, []byte("func (l *Loop) processQuarantine")) || !bytes.Contains(loopBytes, []byte("_ = l.completeTurn()")) {
		t.Fatal("processQuarantine Class-G proof row changed or left the census")
	}
}

type h16ConsumerSubmitter struct {
	outcome  engine.Outcome
	commands []intake.Cmd
}

func (s *h16ConsumerSubmitter) Submit(_ context.Context, cmd intake.Cmd) (<-chan engine.Outcome, string, error) {
	s.commands = append(s.commands, cmd)
	reply := make(chan engine.Outcome, 1)
	reply <- s.outcome
	return reply, cmd.IntakeID, nil
}

func h16ConsumerOutcomeBytes(post string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{"decision_state":"accepted","post_commit_state":%q,"relay_id":"consumer-%s"}`, post, post))
}
