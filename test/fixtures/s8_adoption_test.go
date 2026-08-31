package fixtures_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS8FXCFG12BlessAdoptsLegacyStoreAndRestartsFullReader(t *testing.T) {
	root, candidates := s8LegacyStoreAndCandidates(t)
	if err := store.BlessS8(root, candidates); err != nil {
		t.Fatalf("BlessS8: %v", err)
	}

	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("full reader Load after bless: %v", err)
	}
	if pinned.Engine.Version != 2 || pinned.Supply == nil {
		t.Fatalf("bless did not consume current governed v2 engine: version=%d supply=%#v", pinned.Engine.Version, pinned.Supply)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if err := st.ValidateGenesis(pinned); err != nil {
		t.Fatalf("ValidateGenesis after bless: %v", err)
	}
	adoption := s8AdoptionRecord(t, st)
	if adoption.Headers["channel"] != "bless" {
		t.Fatalf("adoption channel = %q, want bless", adoption.Headers["channel"])
	}
	if got, err := st.ExpectedConfigDigest(); err != nil || got != pinned.Digest {
		t.Fatalf("ExpectedConfigDigest = %q, %v; want %q", got, err, pinned.Digest)
	}
	assertAdoptionBytesMatchMaterialized(t, root, adoption)

	if err := store.BlessS8(root, candidates); !errors.Is(err, store.ErrStoreAlreadyAdopted) {
		t.Fatalf("BlessS8 rerun = %v, want ErrStoreAlreadyAdopted", err)
	}
}

func TestS8FXCFG12BlessRejectsNonCurrentEngineCandidate(t *testing.T) {
	root, candidates := s8LegacyStoreAndCandidates(t)
	v1 := filepath.Join(t.TempDir(), "engine-v1.json")
	if err := os.WriteFile(v1, []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`), 0o644); err != nil {
		t.Fatalf("write v1 engine: %v", err)
	}
	candidates["engine"] = v1
	if err := store.BlessS8(root, candidates); err == nil || !strings.Contains(err.Error(), "current engine v2 with governed supply required") {
		t.Fatalf("BlessS8 v1/no-supply candidate = %v, want current-v2 governed-supply rejection", err)
	}
	if _, err := os.Stat(filepath.Join(root, "records", "s8-adoption.json")); !os.IsNotExist(err) {
		t.Fatalf("adoption record exists after candidate rejection: %v", err)
	}
}

func TestS8FXCFG12BlessRejectsObserveTrueEngineCandidate(t *testing.T) {
	root, candidates := s8LegacyStoreAndCandidates(t)
	if err := os.WriteFile(candidates["engine"], fixtureEngineConfigV2(t, true), 0o644); err != nil {
		t.Fatalf("write observe-active engine candidate: %v", err)
	}

	err := store.BlessS8(root, candidates)
	if err == nil || !strings.Contains(err.Error(), "optional layer observe must be false") {
		t.Fatalf("BlessS8 observe-active candidate = %v, want optional-layer refusal", err)
	}
	if _, err := os.Stat(filepath.Join(root, "records", "s8-adoption.json")); !os.IsNotExist(err) {
		t.Fatalf("adoption record exists after observe-active refusal: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "config", "catalog", "catalog.json")); !os.IsNotExist(err) {
		t.Fatalf("catalog projection exists after observe-active refusal: %v", err)
	}
}

func TestS8FXCFG12BlessCrashWindowsConvergeBeforeFullLoad(t *testing.T) {
	if os.Getenv("FRANK_S8_BLESS_CHILD") == "1" {
		err := store.BlessS8(os.Getenv("FRANK_S8_BLESS_ROOT"), map[string]string{
			"engine": os.Getenv("FRANK_S8_BLESS_ENGINE"), "catalog": os.Getenv("FRANK_S8_BLESS_CATALOG"),
		})
		if err != nil {
			panic(err)
		}
		return
	}
	for _, crash := range []string{"pre_rename", "post_rename", "pre_projection_write:3"} {
		t.Run(crash, func(t *testing.T) {
			root, candidates := s8LegacyStoreAndCandidates(t)
			cmd := exec.Command(os.Args[0], "-test.run=^TestS8FXCFG12BlessCrashWindowsConvergeBeforeFullLoad$")
			cmd.Env = append(os.Environ(),
				"FRANK_S8_BLESS_CHILD=1", "FRANK_S8_BLESS_ROOT="+root,
				"FRANK_S8_BLESS_ENGINE="+candidates["engine"], "FRANK_S8_BLESS_CATALOG="+candidates["catalog"],
				"FRANK_TEST_CRASHPOINT="+crash,
			)
			err := cmd.Run()
			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				t.Fatalf("bless child error = %T %v, want SIGKILL", err, err)
			}
			status, ok := exitErr.Sys().(syscall.WaitStatus)
			if !ok || status.Signal() != syscall.SIGKILL {
				t.Fatalf("bless child status = %v, want SIGKILL", exitErr.Sys())
			}

			if crash == "pre_rename" {
				err = store.BlessS8(root, candidates)
				if err != nil {
					t.Fatalf("bless rerun after pre-pivot crash: %v", err)
				}
			} else {
				ctx, cancel := context.WithCancel(context.Background())
				socket := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s8-%d-%d.sock", os.Getpid(), time.Now().UnixNano()))
				t.Cleanup(func() { _ = os.Remove(socket) })
				serve, stderr := startFrank(t, ctx, buildFrank(t, ctx), root, socket)
				done := make(chan error, 1)
				go func() { done <- serve.Wait() }()
				deadline := time.Now().Add(5 * time.Second)
				ready := false
				for !ready && time.Now().Before(deadline) {
					select {
					case err := <-done:
						cancel()
						t.Fatalf("normal serve exited before readiness after %s: %v\n%s", crash, err, stderr.String())
					default:
						if _, err := os.Stat(socket); err == nil {
							ready = true
							break
						}
						time.Sleep(20 * time.Millisecond)
					}
				}
				if !ready {
					cancel()
					<-done
					t.Fatalf("normal serve did not reach readiness after %s: %s", crash, stderr.String())
				}
				cancel()
				<-done
				if data := stderr.String(); strings.Contains(data, "config-load") || strings.Contains(data, "store-not-adopted") || strings.Contains(data, "digest-mismatch") {
					t.Fatalf("normal serve failed before readiness after %s: %s", crash, data)
				}
			}
			pinned, err := config.Load(store.StoreRootConfigPaths(root))
			if err != nil {
				t.Fatalf("full phase-0 load after recovery: %v", err)
			}
			st, err := store.Open(root)
			if err != nil {
				t.Fatalf("Open recovered: %v", err)
			}
			if err := st.ValidateGenesis(pinned); err != nil {
				t.Fatalf("ValidateGenesis recovered: %v", err)
			}
			assertAdoptionBytesMatchMaterialized(t, root, s8AdoptionRecord(t, st))
		})
	}
}

func TestS8FXCFG12ProductionBlessCommandAcceptsDesignOrdering(t *testing.T) {
	root, candidates := s8LegacyStoreAndCandidates(t)
	cmd := exec.Command(buildFrank(t, context.Background()), "bless-s8", "-root", root, "-engine-config", candidates["engine"], "-catalog", candidates["catalog"])
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("frank bless-s8: %v\n%s", err, out)
	}
	if _, err := config.Load(store.StoreRootConfigPaths(root)); err != nil {
		t.Fatalf("full load after production bless: %v", err)
	}
}

func TestS8FXCFG13And14AdoptionBodyContractAndSingularReplay(t *testing.T) {
	valid := adoptionFixtureRecord(t, []adoptionFixtureMember{
		{Name: "catalog", Bytes: []byte("catalog-bytes")},
		{Name: "engine", Bytes: []byte("engine-bytes")},
	})
	intents, err := store.ConfigChangeIntentsStrict(valid)
	if err != nil {
		t.Fatalf("valid adoption intents: %v", err)
	}
	var gotNames []string
	for _, intent := range intents {
		if intent.Kind == store.IntentConfig {
			gotNames = append(gotNames, filepath.Base(intent.Path))
		}
	}
	if want := []string{"catalog.json", "engine.json"}; !reflect.DeepEqual(gotNames, want) {
		t.Fatalf("adoption config intent order = %v, want %v", gotNames, want)
	}

	badBodies := map[string]string{
		"missing":               adoptionFixtureBody(t, []adoptionFixtureMember{{Name: "catalog", Bytes: []byte("x")}}),
		"duplicate":             adoptionFixtureBody(t, []adoptionFixtureMember{{Name: "catalog", Bytes: []byte("x")}, {Name: "catalog", Bytes: []byte("y")}}),
		"extra":                 adoptionFixtureBody(t, []adoptionFixtureMember{{Name: "catalog", Bytes: []byte("x")}, {Name: "engine", Bytes: []byte("y")}, {Name: "extra", Bytes: []byte("z")}}),
		"misordered":            adoptionFixtureBody(t, []adoptionFixtureMember{{Name: "engine", Bytes: []byte("y")}, {Name: "catalog", Bytes: []byte("x")}}),
		"reserved-adoption":     `{"members":[{"name":"adoption","bytes_b64":"eA=="},{"name":"engine","bytes_b64":"eQ=="}]}`,
		"reserved-fieldspec":    `{"members":[{"name":"catalog","bytes_b64":"eA=="},{"name":"fieldspec","bytes_b64":"eQ=="}]}`,
		"malformed-base64":      `{"members":[{"name":"catalog","bytes_b64":"***"},{"name":"engine","bytes_b64":"eQ=="}]}`,
		"noncanonical-base64":   `{"members":[{"name":"catalog","bytes_b64":"eA"},{"name":"engine","bytes_b64":"eQ=="}]}`,
		"unknown-field":         `{"members":[{"name":"catalog","bytes_b64":"eA==","extra":true},{"name":"engine","bytes_b64":"eQ=="}]}`,
		"duplicate-members-key": `{"members":[{"name":"catalog","bytes_b64":"eA=="},{"name":"engine","bytes_b64":"eQ=="}],"members":[{"name":"catalog","bytes_b64":"eA=="},{"name":"engine","bytes_b64":"eQ=="}]}`,
		"duplicate-name-key":    `{"members":[{"name":"catalog","name":"catalog","bytes_b64":"eA=="},{"name":"engine","bytes_b64":"eQ=="}]}`,
		"duplicate-bytes-key":   `{"members":[{"name":"catalog","bytes_b64":"eA==","bytes_b64":"eA=="},{"name":"engine","bytes_b64":"eQ=="}]}`,
	}
	for name, body := range badBodies {
		t.Run(name, func(t *testing.T) {
			rec := valid
			rec.Body = body
			if _, err := store.ConfigChangeIntentsStrict(rec); err == nil {
				t.Fatal("invalid adoption body accepted")
			}
		})
	}

	singular := record.Record{
		Envelope: record.Envelope{RelayID: "singular-replay", From: "operator", Role: "operator", DeliveryState: record.Accepted},
		Headers:  map[string]string{"PHASE": "SITREP", "member": "engine"}, Body: "singular-byte-exact",
	}
	singularIntents, err := store.ConfigChangeIntentsStrict(singular)
	if err != nil {
		t.Fatalf("singular replay: %v", err)
	}
	if len(singularIntents) < 2 || singularIntents[1].Kind != store.IntentConfig || string(singularIntents[1].Payload) != singular.Body {
		t.Fatalf("singular arm changed: %#v", singularIntents)
	}

	root, _ := s8LegacyStoreAndCandidates(t)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open legacy adoption submit store: %v", err)
	}
	operator := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	adoption, _ := s5SubmitConfigChange(t, st, loadS5Registry(t), operator, record.Record{Headers: map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
		"SUBJECT": "live adoption must reject", "record_kind": "config_change", "member": "adoption", "new_digest": "offline-only",
	}, Body: valid.Body})
	if adoption.Envelope.DeliveryState != record.Rejected || !strings.Contains(adoption.Body, "offline-bless-only") {
		t.Fatalf("pre-adoption member:adoption = %s %q, want typed offline-bless-only reject", adoption.Envelope.DeliveryState, adoption.Body)
	}
}

func TestS8FXCFG15CatalogChangeIsStateAware(t *testing.T) {
	root, candidates := s8LegacyStoreAndCandidates(t)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open legacy: %v", err)
	}
	reg := loadS5Registry(t)
	operator := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	catalog, err := os.ReadFile(candidates["catalog"])
	if err != nil {
		t.Fatalf("read catalog candidate: %v", err)
	}
	pre, _ := s5SubmitConfigChange(t, st, reg, operator, s8ConfigChangeRecord(t, root, "catalog", catalog))
	if pre.Envelope.DeliveryState != record.Rejected || !strings.Contains(pre.Body, "not-adopted") {
		t.Fatalf("pre-adoption catalog = %s %q, want typed not-adopted reject", pre.Envelope.DeliveryState, pre.Body)
	}

	if err := store.BlessS8(root, candidates); err != nil {
		t.Fatalf("BlessS8: %v", err)
	}
	post, intents := s5SubmitConfigChange(t, st, reg, operator, s8ConfigChangeRecord(t, root, "catalog", catalog))
	if post.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("post-adoption catalog = %s %q, want accepted", post.Envelope.DeliveryState, post.Body)
	}
	var configs []store.Intent
	for _, intent := range intents {
		if intent.Kind == store.IntentConfig {
			configs = append(configs, intent)
		}
	}
	if len(configs) != 1 || filepath.Base(configs[0].Path) != "catalog.json" || !bytes.Equal(configs[0].Payload, catalog) {
		t.Fatalf("post-adoption catalog intents = %#v, want one byte-exact catalog intent", configs)
	}
}

type adoptionFixtureMember struct {
	Name  string
	Bytes []byte
}

func s8LegacyStoreAndCandidates(t *testing.T) (string, map[string]string) {
	t.Helper()
	candidates := s8HistoricalConfigSources(t, false)
	if err := os.WriteFile(candidates["engine"], fixtureEngineConfigV2(t, false), 0o644); err != nil {
		t.Fatalf("write v2 adoption engine: %v", err)
	}
	legacyEngine := filepath.Join(t.TempDir(), "engine.json")
	if err := os.WriteFile(legacyEngine, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write legacy engine: %v", err)
	}
	root := t.TempDir()
	if err := store.Init(root, map[string]string{"fieldspec": candidates["fieldspec"], "engine": legacyEngine}); err != nil {
		t.Fatalf("Init legacy store: %v", err)
	}
	return root, map[string]string{"engine": candidates["engine"], "catalog": candidates["catalog"]}
}

func s8AdoptionRecord(t *testing.T, st *store.Store) record.Record {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	for _, rec := range records {
		if rec.Headers["record_kind"] == "config_change" && rec.Headers["member"] == "adoption" {
			return rec
		}
	}
	t.Fatal("adoption record missing")
	return record.Record{}
}

func assertAdoptionBytesMatchMaterialized(t *testing.T, root string, rec record.Record) {
	t.Helper()
	var body struct {
		Members []struct {
			Name     string `json:"name"`
			BytesB64 string `json:"bytes_b64"`
		} `json:"members"`
	}
	if err := json.Unmarshal([]byte(rec.Body), &body); err != nil {
		t.Fatalf("decode adoption record: %v", err)
	}
	for _, member := range body.Members {
		decoded, err := base64.StdEncoding.DecodeString(member.BytesB64)
		if err != nil {
			t.Fatalf("decode %s: %v", member.Name, err)
		}
		path := store.StoreRootConfigPaths(root)[member.Name]
		materialized, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read materialized %s: %v", member.Name, err)
		}
		if !bytes.Equal(decoded, materialized) {
			t.Fatalf("%s adoption bytes differ from materialized member", member.Name)
		}
	}
}

func adoptionFixtureRecord(t *testing.T, members []adoptionFixtureMember) record.Record {
	t.Helper()
	return record.Record{
		Envelope: record.Envelope{RelayID: "adoption-fixture", From: "operator", Role: "operator", DeliveryState: record.Accepted},
		Headers:  map[string]string{"PHASE": "SITREP", "member": "adoption"},
		Body:     adoptionFixtureBody(t, members),
	}
}

func adoptionFixtureBody(t *testing.T, members []adoptionFixtureMember) string {
	t.Helper()
	body := struct {
		Members []struct {
			Name     string `json:"name"`
			BytesB64 string `json:"bytes_b64"`
		} `json:"members"`
	}{}
	for _, member := range members {
		body.Members = append(body.Members, struct {
			Name     string `json:"name"`
			BytesB64 string `json:"bytes_b64"`
		}{Name: member.Name, BytesB64: base64.StdEncoding.EncodeToString(member.Bytes)})
	}
	data, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal adoption fixture: %v", err)
	}
	return string(data)
}
