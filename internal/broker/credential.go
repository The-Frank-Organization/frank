package broker

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	BrokerConfigName  = "broker.json"
	credentialMaxSize = 4 * 1024
)

var ErrCredentialUnavailable = errors.New("broker: credential unavailable")

type CredentialSink struct{ token []byte }

func (*CredentialSink) String() string   { return "broker.CredentialSink{redacted}" }
func (*CredentialSink) GoString() string { return "broker.CredentialSink{redacted}" }

func (sink *CredentialSink) bind(bind func([]byte) error) error {
	if sink == nil || len(sink.token) == 0 || bind == nil {
		return ErrCredentialUnavailable
	}
	return bind(append([]byte(nil), sink.token...))
}

func LoadCredential(configHome string) (*CredentialSink, error) {
	if err := validateConfigHome(configHome); err != nil {
		return nil, err
	}
	configBytes, err := os.ReadFile(filepath.Join(configHome, BrokerConfigName))
	if err != nil {
		return nil, ErrCredentialUnavailable
	}
	decoder := json.NewDecoder(bytes.NewReader(configBytes))
	decoder.DisallowUnknownFields()
	var config struct {
		CredentialFile string `json:"credential_file"`
	}
	if decoder.Decode(&config) != nil || requireEOF(decoder) != nil || config.CredentialFile == "" || filepath.Base(config.CredentialFile) != config.CredentialFile || strings.Contains(config.CredentialFile, string(filepath.Separator)) {
		return nil, ErrCredentialUnavailable
	}
	token, err := readCredential(filepath.Join(configHome, config.CredentialFile))
	if err != nil {
		return nil, err
	}
	return &CredentialSink{token: token}, nil
}

func validateConfigHome(path string) error {
	info, err := os.Lstat(path)
	if err != nil || !info.IsDir() || info.Mode().Perm() != 0o700 || info.Mode()&os.ModeSymlink != 0 {
		return ErrCredentialUnavailable
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != os.Geteuid() {
		return ErrCredentialUnavailable
	}
	return nil
}

func readCredential(path string) ([]byte, error) {
	fd, err := syscall.Open(path, syscall.O_RDONLY|syscall.O_NOFOLLOW|syscall.O_CLOEXEC, 0)
	if err != nil {
		return nil, ErrCredentialUnavailable
	}
	file := os.NewFile(uintptr(fd), "broker-credential")
	if file == nil {
		_ = syscall.Close(fd)
		return nil, ErrCredentialUnavailable
	}
	defer file.Close()
	var before syscall.Stat_t
	if syscall.Fstat(fd, &before) != nil || before.Mode&syscall.S_IFMT != syscall.S_IFREG || before.Mode&0o777 != 0o600 || int(before.Uid) != os.Geteuid() || before.Size <= 0 || before.Size > credentialMaxSize {
		return nil, ErrCredentialUnavailable
	}
	raw, err := io.ReadAll(io.LimitReader(file, credentialMaxSize+1))
	if err != nil || len(raw) == 0 || len(raw) > credentialMaxSize {
		return nil, ErrCredentialUnavailable
	}
	var after syscall.Stat_t
	if syscall.Fstat(fd, &after) != nil || before.Dev != after.Dev || before.Ino != after.Ino || before.Size != after.Size {
		return nil, ErrCredentialUnavailable
	}
	if raw[len(raw)-1] == '\n' {
		raw = raw[:len(raw)-1]
	}
	if len(raw) == 0 {
		return nil, ErrCredentialUnavailable
	}
	for _, character := range raw {
		if character < 0x21 || character > 0x7e {
			return nil, ErrCredentialUnavailable
		}
	}
	return append([]byte(nil), raw...), nil
}
