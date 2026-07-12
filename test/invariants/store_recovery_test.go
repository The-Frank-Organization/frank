package invariants_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/store"
)

func TestLawCanonicalWins(t *testing.T) {
	_, _ = requireCatalogLaw(t, "TestLawCanonicalWins")
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open canonical store: %v", err)
	}
	commitRecordDefault(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "relay-canonical",
			DispatchID:    "canonical-law",
			From:          "seat-a.implementer",
			To:            "seat-b.planner",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "canonical truth"},
		Body:    "canonical body",
	})

	indexPath := filepath.Join(root, "projections", "INDEX.md")
	mailboxPath := filepath.Join(root, "mailboxes", "seat-b.planner.jsonl")
	renderPath := filepath.Join(root, "projections", "relays", "canonical-law", "SITREP-implementer-relay-canonical.md")
	if err := os.WriteFile(indexPath, []byte("| forged | SITREP | attacker | seat-b.planner |  | accepted |\n"), 0o644); err != nil {
		t.Fatalf("corrupt index: %v", err)
	}
	if err := os.WriteFile(mailboxPath, []byte("forged\n"), 0o644); err != nil {
		t.Fatalf("corrupt mailbox: %v", err)
	}
	if err := os.WriteFile(renderPath, []byte("forged projection\n"), 0o644); err != nil {
		t.Fatalf("corrupt render: %v", err)
	}

	if err := st.RebuildProjections(); err != nil {
		t.Fatalf("rebuild projections: %v", err)
	}
	got, err := st.Project("seat-b.planner")
	if err != nil {
		t.Fatalf("project rebuilt mailbox: %v", err)
	}
	if want := []string{"relay-canonical"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("rebuilt mailbox = %v, want canonical %v", got, want)
	}
	rec, err := st.Read("relay-canonical")
	if err != nil {
		t.Fatalf("read canonical record: %v", err)
	}
	if rec.Headers["SUBJECT"] != "canonical truth" || rec.Body != "canonical body" {
		t.Fatalf("canonical record changed after projection corruption: %+v", rec)
	}
	rendered, err := os.ReadFile(renderPath)
	if err != nil {
		t.Fatalf("read rebuilt render: %v", err)
	}
	if strings.Contains(string(rendered), "forged projection") ||
		!strings.Contains(string(rendered), "canonical body") {
		t.Fatalf("rebuilt render did not follow canonical record: %s", rendered)
	}
}

func TestLawOnePivotPerMutation(t *testing.T) {
	if os.Getenv("FRANK_S7_PIVOT_CHILD") == "1" {
		root := os.Getenv("FRANK_S7_PIVOT_ROOT")
		st, err := store.Open(root)
		if err != nil {
			panic(err)
		}
		commitRecordDefault(t, st, record.Record{
			Envelope: record.Envelope{
				RelayID:       "law-pivot",
				From:          "s7.implementer",
				Role:          "implementer",
				DeliveryState: record.Accepted,
				IntakeID:      "intake-law-pivot",
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "one pivot"},
		})
		return
	}

	_, _ = requireCatalogLaw(t, "TestLawOnePivotPerMutation")
	t.Run("exactly one canonical rename", func(t *testing.T) {
		root := t.TempDir()
		counter := filepath.Join(t.TempDir(), "renames.log")
		if err := runPivotChild(root, counter, ""); err != nil {
			t.Fatalf("clean pivot child: %v", err)
		}
		data, err := os.ReadFile(counter)
		if err != nil {
			t.Fatalf("read rename counter: %v", err)
		}
		if got := strings.Count(string(data), "records/law-pivot.json\n"); got != 1 {
			t.Fatalf("canonical rename count = %d, want 1; counter:\n%s", got, data)
		}
		if _, err := os.Stat(filepath.Join(root, "records", "law-pivot.json")); err != nil {
			t.Fatalf("canonical record absent after clean pivot: %v", err)
		}
	})
	t.Run("crash before pivot is absent", func(t *testing.T) {
		root := t.TempDir()
		err := runPivotChild(root, filepath.Join(t.TempDir(), "renames.log"), "pre_rename")
		if err == nil {
			t.Fatal("pre-rename child did not crash")
		}
		if _, statErr := os.Stat(filepath.Join(root, "records", "law-pivot.json")); !os.IsNotExist(statErr) {
			t.Fatalf("pre-pivot canonical record exists: %v", statErr)
		}
	})
	t.Run("crash after pivot is present", func(t *testing.T) {
		root := t.TempDir()
		err := runPivotChild(root, filepath.Join(t.TempDir(), "renames.log"), "pre_dir_fsync")
		if err == nil {
			t.Fatal("post-rename child did not crash")
		}
		if _, statErr := os.Stat(filepath.Join(root, "records", "law-pivot.json")); statErr != nil {
			t.Fatalf("post-pivot canonical record absent: %v", statErr)
		}
	})
}

func TestLawRebuildBeforeOpen(t *testing.T) {
	_, _ = requireCatalogLaw(t, "TestLawRebuildBeforeOpen")
	root := t.TempDir()
	pinned := initPinnedStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open recovery store: %v", err)
	}
	commitRecordDefault(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "ready-after-rebuild",
			From:          "seat-a.implementer",
			To:            "seat-b.planner",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "ready after rebuild"},
	})
	if err := os.Remove(filepath.Join(root, "projections", "INDEX.md")); err != nil {
		t.Fatalf("remove derived index: %v", err)
	}
	torn := filepath.Join(root, "staging", "records", "torn.tmp")
	if err := os.MkdirAll(filepath.Dir(torn), 0o755); err != nil {
		t.Fatalf("mkdir torn staging: %v", err)
	}
	if err := os.WriteFile(torn, []byte("torn"), 0o644); err != nil {
		t.Fatalf("write torn staging: %v", err)
	}

	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("open recovery journal: %v", err)
	}
	if _, err := intake.NewWriter[engine.Outcome](journal, pinned.Engine, nil); !errors.Is(err, intake.ErrWriterNotReady) {
		t.Fatalf("pre-recovery writer error = %v, want ErrWriterNotReady", err)
	}
	tools := channel.ToolSet{
		Submit:  staticToolResult(`{"state":"accepted"}`),
		Project: staticToolResult(`[]`),
		Read:    staticToolResult(`{"relay_id":"ready-after-rebuild"}`),
	}
	before := channel.FullSurface(nil, tools)
	beforeClient, beforeServer := serveToolSurface(t, before)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	beforeNames, err := beforeClient.ListTools(ctx)
	if err != nil {
		t.Fatalf("list pre-recovery tools: %v", err)
	}
	if len(beforeNames) != 0 {
		t.Fatalf("pre-recovery tools = %v, want empty", beforeNames)
	}
	if _, err := beforeClient.Call(ctx, "submit", nil); err == nil {
		t.Fatal("pre-recovery submit unexpectedly available")
	}
	_ = beforeServer.Close()

	result, err := frankrecover.Run(root, pinned)
	if err != nil {
		t.Fatalf("run recovery: %v", err)
	}
	if result.Ready == nil || result.Diag != nil {
		t.Fatalf("recovery result = %+v, want Ready", result)
	}
	if _, err := os.Stat(torn); !os.IsNotExist(err) {
		t.Fatalf("torn staging survived recovery: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "projections", "INDEX.md")); err != nil {
		t.Fatalf("projection not rebuilt before Ready: %v", err)
	}

	afterClient, afterServer := serveToolSurface(t, channel.FullSurface(result.Ready, tools))
	defer func() { _ = afterServer.Close() }()
	afterNames, err := afterClient.ListTools(ctx)
	if err != nil {
		t.Fatalf("list post-recovery tools: %v", err)
	}
	if want := []string{"submit", "project", "read"}; !reflect.DeepEqual(afterNames, want) {
		t.Fatalf("post-recovery tools = %v, want %v", afterNames, want)
	}
	if _, err := afterClient.Call(ctx, "submit", nil); err != nil {
		t.Fatalf("post-recovery submit unavailable: %v", err)
	}
}

func commitRecordDefault(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func runPivotChild(root, counter, crash string) error {
	cmd := exec.Command(os.Args[0], "-test.run=^TestLawOnePivotPerMutation$", "-test.count=1")
	cmd.Env = append(os.Environ(),
		"FRANK_S7_PIVOT_CHILD=1",
		"FRANK_S7_PIVOT_ROOT="+root,
		"FRANK_TEST_RENAME_COUNTER="+counter,
	)
	if crash != "" {
		cmd.Env = append(cmd.Env, "FRANK_TEST_CRASHPOINT="+crash)
	}
	return cmd.Run()
}

func initPinnedStore(t *testing.T, root string) *config.Pinned {
	t.Helper()
	sourceDir := t.TempDir()
	enginePath := filepath.Join(sourceDir, "engine.json")
	if err := os.WriteFile(enginePath, []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	if err := store.Init(root, map[string]string{
		"catalog":   filepath.Join(repoRoot(t), "test", "invariants", "catalog.v1.json"),
		"engine":    enginePath,
		"fieldspec": filepath.Join(repoRoot(t), "internal", "fieldspec", "registry.json"),
	}); err != nil {
		t.Fatalf("init pinned store: %v", err)
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load pinned config: %v", err)
	}
	return pinned
}

func serveToolSurface(t *testing.T, tools channel.ToolSet) (*channel.Client, *channel.Server) {
	t.Helper()
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s7-ready-%d.sock", time.Now().UnixNano()))
	server, err := channel.Serve(sock, tools)
	if err != nil {
		t.Fatalf("serve tool surface: %v", err)
	}
	client, err := channel.Dial(context.Background(), sock)
	if err != nil {
		_ = server.Close()
		t.Fatalf("dial tool surface: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(sock) })
	return client, server
}
