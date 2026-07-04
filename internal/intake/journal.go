package intake

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fsio"
	"github.com/jackli/frank/internal/store"
)

type Cmd struct {
	IntakeID    string          `json:"intake_id,omitempty"`
	Seat        string          `json:"seat"`
	Verb        string          `json:"verb"`
	Payload     json.RawMessage `json:"payload,omitempty"`
	ContentHash string          `json:"content_hash,omitempty"`
}

type Journal struct {
	Root string
	path string
}

func Open(root string) (*Journal, error) {
	path := filepath.Join(root, "journal", "intake.jsonl")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return &Journal{Root: root, path: path}, nil
}

func (j *Journal) Append(cmd Cmd) (string, error) {
	if cmd.ContentHash == "" {
		cmd.ContentHash = hashCmd(cmd)
	}
	entries, err := j.ReadAll()
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if entry.ContentHash == cmd.ContentHash {
			return entry.IntakeID, nil
		}
	}
	cmd.IntakeID = fmt.Sprintf("intake-%06d", len(entries)+1)
	data, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}
	data = append(data, '\n')
	f, err := os.OpenFile(j.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if err := fsio.AppendFsync(f, data); err != nil {
		return "", err
	}
	crashpoint.Hit("post_intake_fsync")
	return cmd.IntakeID, nil
}

func (j *Journal) ReadAll() ([]Cmd, error) {
	f, err := os.Open(j.path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var entries []Cmd
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var cmd Cmd
		if err := json.Unmarshal(scanner.Bytes(), &cmd); err != nil {
			return nil, err
		}
		entries = append(entries, cmd)
	}
	return entries, scanner.Err()
}

func Unconsumed(ctx context.Context, j *Journal, st *store.Store) ([]Cmd, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	entries, err := j.ReadAll()
	if err != nil {
		return nil, err
	}
	records, err := st.Records()
	if err != nil {
		return nil, err
	}
	consumed := map[string]struct{}{}
	for _, rec := range records {
		if rec.Envelope.IntakeID != "" {
			consumed[rec.Envelope.IntakeID] = struct{}{}
		}
	}
	var out []Cmd
	for _, cmd := range entries {
		if _, ok := consumed[cmd.IntakeID]; !ok {
			out = append(out, cmd)
		}
	}
	return out, nil
}

func hashCmd(cmd Cmd) string {
	sum := sha256.Sum256([]byte(cmd.Seat + "\x00" + cmd.Verb + "\x00" + string(cmd.Payload)))
	return hex.EncodeToString(sum[:])
}
