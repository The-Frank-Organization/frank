package invariants_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

type surfaceCapture struct {
	Family            string
	Channel           string
	Bytes             []byte
	AllowedCredential string
	AllowedEndpoint   string
}

type sinkSite struct {
	PatternID string
	Family    string
	Symbol    string
	Location  string
}

func TestLawPathHygiene(t *testing.T) {
	cat, _ := requireCatalogLaw(t, "TestLawPathHygiene")
	wantFamilies := []string{
		"bounce-reason",
		"process-tool-errors",
		"tool-descriptions-results",
		"rendered-projections",
		"delivery-payloads",
		"seat-mint-accept-reply",
	}
	var gotFamilies []string
	for _, family := range cat.PathHygiene.Families {
		gotFamilies = append(gotFamilies, family.ID)
	}
	if !reflect.DeepEqual(gotFamilies, wantFamilies) {
		t.Fatalf("path-hygiene families = %v, want literal %v", gotFamilies, wantFamilies)
	}
	if len(cat.PathHygiene.Families[5].CarveOuts) != 2 {
		t.Fatalf("seat_mint carve-outs = %v, want exactly credential and endpoint", cat.PathHygiene.Families[5].CarveOuts)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mintRoot := t.TempDir()
	mintReply, mintRead, mintCredential, mintEndpoint := liveSeatMintCapture(t, ctx, mintRoot)
	projectionRoot, projectionBytes := renderedProjectionCapture(t)
	processBytes, processSocket := processErrorCapture(t)
	toolBytes, toolSocket := toolDescriptionCapture(t)
	deliveryBytes, deliverySocket := deliveryCapture(t)

	captures := []surfaceCapture{
		{
			Family: "bounce-reason",
			Bytes: []byte(bounce.Format(fieldspec.Violation{
				Field:  "AUTHORITY",
				Class:  "non-boot-before-active",
				Reason: filepath.Join(mintRoot, "records", "planted"),
			})),
		},
		{Family: "process-tool-errors", Bytes: processBytes},
		{Family: "tool-descriptions-results", Bytes: toolBytes},
		{Family: "rendered-projections", Bytes: projectionBytes},
		{Family: "delivery-payloads", Bytes: deliveryBytes},
		{
			Family:            "seat-mint-accept-reply",
			Channel:           "operator",
			Bytes:             mintReply,
			AllowedCredential: mintCredential,
			AllowedEndpoint:   mintEndpoint,
		},
	}
	forbidden := []string{
		mintRoot,
		projectionRoot,
		processSocket,
		toolSocket,
		deliverySocket,
		mintCredential,
		mintEndpoint,
		"/records/",
		"/staging/",
		"/outbox/",
		"/binding/",
		"registry.json",
		"engine.json",
		"realized_mint_ref",
	}
	if err := scanSurfaceCorpus(cat, captures, forbidden); err != nil {
		t.Fatalf("complete seat-delivered corpus: %v", err)
	}
	for _, forbiddenValue := range []string{mintCredential, mintEndpoint, mintRoot, "realized_mint_ref"} {
		if forbiddenValue != "" && bytes.Contains(mintRead, []byte(forbiddenValue)) {
			t.Fatalf("committed mint read leaked carve-out %q: %s", forbiddenValue, mintRead)
		}
	}

	t.Run("planted leak is rejected inside mint reply", func(t *testing.T) {
		var planted map[string]any
		if err := json.Unmarshal(mintReply, &planted); err != nil {
			t.Fatalf("decode mint reply for planted leak: %v", err)
		}
		planted["detail"] = filepath.Join(mintRoot, "records", "leak.json")
		data, err := json.Marshal(planted)
		if err != nil {
			t.Fatalf("marshal planted mint leak: %v", err)
		}
		bad := append([]surfaceCapture(nil), captures...)
		bad[5].Bytes = data
		if err := scanSurfaceCorpus(cat, bad, forbidden); err == nil {
			t.Fatal("planted mint-reply store path did not fail the scanner")
		}
	})
	t.Run("mint carve-outs are operator-channel only", func(t *testing.T) {
		bad := append([]surfaceCapture(nil), captures...)
		bad[5].Channel = "path-seat.implementer"
		if err := scanSurfaceCorpus(cat, bad, forbidden); err == nil {
			t.Fatal("mint reply carve-outs passed on a non-operator channel")
		}
	})
	t.Run("unregistered family is rejected", func(t *testing.T) {
		bad := append([]surfaceCapture(nil), captures...)
		bad = append(bad, surfaceCapture{Family: "future-seat-surface", Bytes: []byte("safe-looking")})
		if err := scanSurfaceCorpus(cat, bad, forbidden); err == nil {
			t.Fatal("unregistered seat-visible family did not fail the census")
		}
	})

	sites := discoverSinkSites(t, cat)
	if err := validateSinkSites(cat, sites); err != nil {
		t.Fatalf("AST sink census: %v", err)
	}
	t.Run("unregistered sink pattern is rejected", func(t *testing.T) {
		bad := append([]sinkSite(nil), sites...)
		bad = append(bad, sinkSite{
			PatternID: "future-sink-pattern",
			Family:    "future-seat-surface",
			Symbol:    "future.Emit",
			Location:  "synthetic.go:1",
		})
		if err := validateSinkSites(cat, bad); err == nil {
			t.Fatal("unregistered AST sink did not fail the census")
		}
	})
}

func scanSurfaceCorpus(cat catalog, captures []surfaceCapture, forbidden []string) error {
	known := map[string]outputFamily{}
	for _, family := range cat.PathHygiene.Families {
		known[family.ID] = family
	}
	counts := map[string]int{}
	for _, capture := range captures {
		family, ok := known[capture.Family]
		if !ok {
			return fmt.Errorf("unregistered seat-visible family %q", capture.Family)
		}
		counts[capture.Family]++
		data := capture.Bytes
		if capture.Family == "seat-mint-accept-reply" {
			if capture.Channel != "operator" {
				return fmt.Errorf("seat_mint carve-outs reached %q channel", capture.Channel)
			}
			if len(family.CarveOuts) != 2 {
				return fmt.Errorf("seat_mint family has %d carve-outs", len(family.CarveOuts))
			}
			var reply map[string]any
			if err := json.Unmarshal(data, &reply); err != nil {
				return fmt.Errorf("decode seat_mint reply: %w", err)
			}
			credential, _ := reply["credential"].(string)
			endpoint, _ := reply["endpoint"].(string)
			if credential == "" || credential != capture.AllowedCredential {
				return fmt.Errorf("seat_mint credential carve-out mismatch")
			}
			if endpoint == "" || endpoint != capture.AllowedEndpoint {
				return fmt.Errorf("seat_mint endpoint carve-out mismatch")
			}
			delete(reply, "credential")
			delete(reply, "endpoint")
			var err error
			data, err = json.Marshal(reply)
			if err != nil {
				return fmt.Errorf("marshal seat_mint non-carve-out fields: %w", err)
			}
		}
		for _, value := range forbidden {
			if value != "" && bytes.Contains(data, []byte(value)) {
				return fmt.Errorf("%s leaked %q in %s", capture.Family, value, data)
			}
		}
	}
	if len(counts) != len(known) {
		return fmt.Errorf("captured %d families, catalog names %d", len(counts), len(known))
	}
	for id := range known {
		if counts[id] != 1 {
			return fmt.Errorf("family %s capture count = %d, want 1", id, counts[id])
		}
	}
	return nil
}

func discoverSinkSites(t *testing.T, cat catalog) []sinkSite {
	t.Helper()
	type tokenOwner struct {
		pattern sinkPattern
		token   string
	}
	var owners []tokenOwner
	for _, pattern := range cat.PathHygiene.SinkPatterns {
		for _, symbol := range strings.Split(pattern.Symbol, "|") {
			owners = append(owners, tokenOwner{pattern: pattern, token: symbol})
		}
	}
	var sites []sinkSite
	fset := token.NewFileSet()
	for _, dir := range []string{"cmd", "internal"} {
		root := filepath.Join(repoRoot(t), dir)
		err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			file, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				return err
			}
			ast.Inspect(file, func(node ast.Node) bool {
				switch value := node.(type) {
				case *ast.SelectorExpr:
					full := value.Sel.Name
					if ident, ok := value.X.(*ast.Ident); ok {
						full = ident.Name + "." + value.Sel.Name
					}
					for _, owner := range owners {
						if owner.token == full || owner.token == value.Sel.Name {
							sites = append(sites, sinkSite{
								PatternID: owner.pattern.ID,
								Family:    owner.pattern.Family,
								Symbol:    full,
								Location:  relativeLocation(t, fset.Position(value.Pos())),
							})
							break
						}
					}
					return false
				case *ast.Ident:
					for _, owner := range owners {
						if !strings.Contains(owner.token, ".") && owner.token == value.Name {
							sites = append(sites, sinkSite{
								PatternID: owner.pattern.ID,
								Family:    owner.pattern.Family,
								Symbol:    value.Name,
								Location:  relativeLocation(t, fset.Position(value.Pos())),
							})
							break
						}
					}
				}
				return true
			})
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s sink sites: %v", dir, err)
		}
	}
	sort.Slice(sites, func(i, j int) bool {
		if sites[i].PatternID != sites[j].PatternID {
			return sites[i].PatternID < sites[j].PatternID
		}
		if sites[i].Location != sites[j].Location {
			return sites[i].Location < sites[j].Location
		}
		return sites[i].Symbol < sites[j].Symbol
	})
	return sites
}

func validateSinkSites(cat catalog, sites []sinkSite) error {
	patterns := map[string]sinkPattern{}
	families := map[string]bool{}
	for _, family := range cat.PathHygiene.Families {
		families[family.ID] = true
	}
	for _, pattern := range cat.PathHygiene.SinkPatterns {
		if !families[pattern.Family] {
			return fmt.Errorf("sink pattern %s maps to unknown family %s", pattern.ID, pattern.Family)
		}
		if _, duplicate := patterns[pattern.ID]; duplicate {
			return fmt.Errorf("duplicate sink pattern %s", pattern.ID)
		}
		patterns[pattern.ID] = pattern
	}
	counts := map[string]int{}
	for _, site := range sites {
		pattern, ok := patterns[site.PatternID]
		if !ok {
			return fmt.Errorf("unregistered sink pattern %s at %s", site.PatternID, site.Location)
		}
		if site.Family != pattern.Family {
			return fmt.Errorf("sink %s family %s, catalog says %s", site.Location, site.Family, pattern.Family)
		}
		counts[site.PatternID]++
	}
	var countMismatches []string
	for _, pattern := range cat.PathHygiene.SinkPatterns {
		if counts[pattern.ID] != pattern.ExpectedSites {
			countMismatches = append(countMismatches, fmt.Sprintf(
				"%s site count = %d, catalog pins %d",
				pattern.ID,
				counts[pattern.ID],
				pattern.ExpectedSites,
			))
		}
	}
	if len(countMismatches) > 0 {
		return errors.New(strings.Join(countMismatches, "; "))
	}
	return nil
}

func relativeLocation(t *testing.T, pos token.Position) string {
	t.Helper()
	relative, err := filepath.Rel(repoRoot(t), pos.Filename)
	if err != nil {
		return pos.String()
	}
	return fmt.Sprintf("%s:%d", filepath.ToSlash(relative), pos.Line)
}

func processErrorCapture(t *testing.T) ([]byte, string) {
	t.Helper()
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open process-error store: %v", err)
	}
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{}, nil, errors.New(filepath.Join(root, "records", "leak.json"))
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)
	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{
		IntakeID: "intake-path-error",
		Seat:     "s7.implementer",
		Role:     "implementer",
		Verb:     "submit",
		Payload:  json.RawMessage(`{"path":"hidden"}`),
	}, ReplyCh: reply}
	out := waitOutcome(t, reply)

	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s7-error-%d.sock", time.Now().UnixNano()))
	server, err := channel.Serve(sock, channel.ToolSet{
		Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return nil, errors.New(filepath.Join(root, "config", "registry.json"))
		},
	})
	if err != nil {
		t.Fatalf("serve tool-error surface: %v", err)
	}
	defer func() { _ = server.Close() }()
	client, err := channel.Dial(context.Background(), sock)
	if err != nil {
		t.Fatalf("dial tool-error surface: %v", err)
	}
	_, callErr := client.Call(context.Background(), "submit", nil)
	if callErr == nil {
		t.Fatal("path-bearing tool failure unexpectedly succeeded")
	}
	data, err := json.Marshal(map[string]string{
		"loop_reason": out.Reason,
		"tool_error":  callErr.Error(),
		"root_lock": (store.ErrRootLocked{
			ErrorClass:    "root-lock-held",
			HolderPID:     123,
			HolderStarted: "2026-07-10T00:00:00Z",
		}).Error(),
	})
	if err != nil {
		t.Fatalf("marshal process errors: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(sock) })
	return data, sock
}

func toolDescriptionCapture(t *testing.T) ([]byte, string) {
	t.Helper()
	tools := channel.ToolSet{
		Submit:  staticToolResult(`{"state":"accepted"}`),
		Project: staticToolResult(`["relay-safe"]`),
		Read:    staticToolResult(`{"relay_id":"relay-safe"}`),
	}
	client, server := serveToolSurface(t, tools)
	defer func() { _ = server.Close() }()
	described, err := client.DescribeTools(context.Background(), channel.DescribeRequest{Phase: "SITREP", Tier: "small"})
	if err != nil {
		t.Fatalf("describe tool surface: %v", err)
	}
	result, err := client.Call(context.Background(), "project", nil)
	if err != nil {
		t.Fatalf("call described project: %v", err)
	}
	data, err := json.Marshal(map[string]any{"description": described, "result": json.RawMessage(result)})
	if err != nil {
		t.Fatalf("marshal descriptions/results: %v", err)
	}
	return data, ""
}

func renderedProjectionCapture(t *testing.T) (string, []byte) {
	t.Helper()
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open projection-corpus store: %v", err)
	}
	commitRecordDefault(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "relay-path-clean",
			DispatchID:    "path-clean",
			From:          "seat-a.implementer",
			To:            "seat-b.planner",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "path-clean projection"},
		Body:    "path-clean body",
	})
	path := filepath.Join(root, "projections", "relays", "path-clean", "SITREP-implementer-relay-path-clean.md")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read rendered projection corpus: %v", err)
	}
	return root, data
}

func deliveryCapture(t *testing.T) ([]byte, string) {
	t.Helper()
	root := t.TempDir()
	manager, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open delivery seat manager: %v", err)
	}
	credential, err := manager.Mint("delivery-seat.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("mint delivery seat: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s7-delivery-%d.sock", time.Now().UnixNano()))
	server, err := channel.ServeAuthenticated(sock, manager, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{Project: staticToolResult(`[]`)}
	})
	if err != nil {
		t.Fatalf("serve delivery surface: %v", err)
	}
	defer func() { _ = server.Close() }()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := channel.DialAuthenticated(ctx, sock, credential.Value)
	if err != nil {
		t.Fatalf("dial delivery surface: %v", err)
	}
	payload := []byte(`{"relay_id":"relay-delivery-safe"}`)
	if err := server.PushTo("delivery-seat.implementer", payload); err != nil {
		t.Fatalf("push delivery payload: %v", err)
	}
	got, err := client.NextPush(ctx)
	if err != nil {
		t.Fatalf("read delivery payload: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(sock) })
	return got, sock
}

func liveSeatMintCapture(t *testing.T, ctx context.Context, root string) ([]byte, []byte, string, string) {
	t.Helper()
	initPinnedStore(t, root)
	manager, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open live-mint seat manager: %v", err)
	}
	operatorCredential, err := manager.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("mint live operator: %v", err)
	}
	binary := buildFrankBinary(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s7-mint-%d.sock", time.Now().UnixNano()))
	command := exec.CommandContext(ctx, binary,
		"-root", root,
		"-socket", sock,
		"-registry", filepath.Join(repoRoot(t), "internal", "fieldspec", "registry.json"),
	)
	var stderr bytes.Buffer
	command.Stderr = &stderr
	if err := command.Start(); err != nil {
		t.Fatalf("start live frank: %v", err)
	}
	t.Cleanup(func() {
		if command.Process != nil {
			_ = command.Process.Kill()
		}
		_ = command.Wait()
		_ = os.Remove(sock)
	})
	waitForSocketPath(t, sock, stderr.String)
	operator, err := channel.DialAuthenticated(ctx, sock, operatorCredential.Value)
	if err != nil {
		t.Fatalf("dial live operator stderr=%s: %v", stderr.String(), err)
	}
	description, err := operator.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "small"})
	if err != nil {
		t.Fatalf("describe live operator: %v", err)
	}
	body, err := json.Marshal(map[string]any{
		"seat":        "path-seat.implementer",
		"role":        "implementer",
		"is_operator": false,
	})
	if err != nil {
		t.Fatalf("marshal live seat_mint body: %v", err)
	}
	payload, err := json.Marshal(fieldspec.SubmitPayload{
		Record: record.Record{
			Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "small",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "path hygiene seat mint",
				"record_kind":     "seat_mint",
			},
			Body: string(body),
		},
		FormDigest: description.FormDigest,
	})
	if err != nil {
		t.Fatalf("marshal live seat_mint payload: %v", err)
	}
	reply, err := operator.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit live seat_mint stderr=%s: %v", stderr.String(), err)
	}
	var outcome struct {
		State      string `json:"state"`
		RelayID    string `json:"relay_id"`
		Credential string `json:"credential"`
		Endpoint   string `json:"endpoint"`
	}
	if err := json.Unmarshal(reply, &outcome); err != nil {
		t.Fatalf("decode live seat_mint reply %s: %v", reply, err)
	}
	if outcome.State != record.Accepted || outcome.RelayID == "" || outcome.Credential == "" || outcome.Endpoint != sock {
		t.Fatalf("live seat_mint outcome = %+v, socket=%s", outcome, sock)
	}
	readArgs, _ := json.Marshal(map[string]string{"relay_id": outcome.RelayID})
	read, err := operator.Call(ctx, "read", readArgs)
	if err != nil {
		t.Fatalf("read live seat_mint record: %v", err)
	}
	return reply, read, outcome.Credential, outcome.Endpoint
}

func buildFrankBinary(t *testing.T, ctx context.Context) string {
	t.Helper()
	binary := filepath.Join(t.TempDir(), "frank")
	command := exec.CommandContext(ctx, "go", "build", "-o", binary, "./cmd/frank")
	command.Dir = repoRoot(t)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("build frank binary: %v\n%s", err, output)
	}
	return binary
}

func waitForSocketPath(t *testing.T, path string, stderr func() string) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(path); err == nil {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatalf("socket %s was not created; stderr=%s", path, stderr())
}
