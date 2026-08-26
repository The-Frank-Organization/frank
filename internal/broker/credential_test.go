package broker

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestCredentialCustodyLoadsOnlyFromConfigHomeAndStaysRedacted(t *testing.T) {
	home := t.TempDir()
	if err := os.Chmod(home, 0o700); err != nil {
		t.Fatal(err)
	}
	credentialPath := filepath.Join(home, "seat.credential")
	sentinel := "C09_SENTINEL_SEAT_CREDENTIAL"
	if err := os.WriteFile(filepath.Join(home, BrokerConfigName), []byte(`{"credential_file":"seat.credential"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(credentialPath, []byte(sentinel+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	sink, err := LoadCredential(home)
	if err != nil {
		t.Fatal(err)
	}
	if got := fmt.Sprintf("%v %#v", sink, sink); strings.Contains(got, sentinel) || !strings.Contains(got, "redacted") {
		t.Fatalf("credential rendered on exposed surface: %q", got)
	}
	wire, err := json.Marshal(sink)
	if err != nil || strings.Contains(string(wire), sentinel) || string(wire) != `{}` {
		t.Fatalf("credential reached a wire surface: wire=%s err=%v", wire, err)
	}
	sinkType := reflect.TypeOf(sink)
	for index := 0; index < sinkType.NumMethod(); index++ {
		if method := sinkType.Method(index); method.Name != "String" && method.Name != "GoString" {
			t.Fatalf("credential sink exposes method %s", method.Name)
		}
	}
	used := ""
	if err := sink.bind(func(token []byte) error { used = string(token); return nil }); err != nil || used != sentinel {
		t.Fatalf("credential bind did not receive exact private token: used=%q err=%v", used, err)
	}
}

func TestCredentialCustodyRejectsSymlinkModeAndMultiline(t *testing.T) {
	for _, test := range []struct {
		name string
		mode os.FileMode
		body string
		link bool
	}{
		{name: "mode", mode: 0o644, body: "token\n"},
		{name: "multiline", mode: 0o600, body: "token\nsecond\n"},
		{name: "symlink", mode: 0o600, body: "token\n", link: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			home := t.TempDir()
			if err := os.Chmod(home, 0o700); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(home, BrokerConfigName), []byte(`{"credential_file":"seat.credential"}`), 0o600); err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(home, "seat.credential")
			if test.link {
				target := filepath.Join(t.TempDir(), "target")
				if err := os.WriteFile(target, []byte(test.body), test.mode); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(target, path); err != nil {
					t.Fatal(err)
				}
			} else if err := os.WriteFile(path, []byte(test.body), test.mode); err != nil {
				t.Fatal(err)
			}
			if sink, err := LoadCredential(home); err == nil || sink != nil || strings.Contains(fmt.Sprint(err), test.body) {
				t.Fatalf("invalid credential accepted/leaked: sink=%v err=%v", sink, err)
			}
		})
	}
}
