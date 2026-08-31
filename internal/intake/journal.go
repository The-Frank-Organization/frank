package intake

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/crashpoint"
	"github.com/The-Frank-Organization/frank/internal/fsio"
	"github.com/The-Frank-Organization/frank/internal/store"
)

const DefaultSegmentRotateBytes int64 = 4 * 1024 * 1024

var ErrLegacyJournal = errors.New("legacy intake journal layout unsupported")

type Cmd struct {
	IntakeID       string          `json:"intake_id,omitempty"`
	Seat           string          `json:"seat"`
	Role           string          `json:"role,omitempty"`
	IsOperator     bool            `json:"is_operator,omitempty"`
	AuthGeneration string          `json:"auth_generation,omitempty"`
	Verb           string          `json:"verb"`
	Payload        json.RawMessage `json:"payload,omitempty"`
	ContentHash    string          `json:"content_hash,omitempty"`
}

type Journal struct {
	Root        string
	dir         string
	rotateBytes int64
}

type Segment struct {
	Seq     int
	Path    string
	Entries []Cmd
}

type segmentHeader struct {
	SegmentHeader bool `json:"segment_header"`
	HighWater     int  `json:"high_water"`
}

func Open(root string) (*Journal, error) {
	return OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: DefaultSegmentRotateBytes})
}

func OpenWithConfig(root string, cfg config.EngineConfig) (*Journal, error) {
	if _, err := os.Stat(filepath.Join(root, "journal", "intake.jsonl")); err == nil {
		return nil, ErrLegacyJournal
	} else if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	dir := filepath.Join(root, "journal", "intake")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	rotateBytes := cfg.SegmentRotateBytes
	if rotateBytes <= 0 {
		rotateBytes = DefaultSegmentRotateBytes
	}
	return &Journal{Root: root, dir: dir, rotateBytes: rotateBytes}, nil
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
	next, err := j.nextIntakeNumber()
	if err != nil {
		return "", err
	}
	cmd.IntakeID = fmt.Sprintf("intake-%06d", next)
	data, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}
	data = append(data, '\n')
	seq, path, err := j.activeSegment()
	if err != nil {
		return "", err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return "", err
	}
	if err := fsio.AppendFsync(f, data); err != nil {
		_ = f.Close()
		return "", err
	}
	crashpoint.Hit("post_intake_fsync")
	if err := f.Close(); err != nil {
		return "", err
	}
	if err := j.rotateAfterAppend(seq, path); err != nil {
		return "", err
	}
	return cmd.IntakeID, nil
}

func (j *Journal) ReadAll() ([]Cmd, error) {
	segments, err := j.Segments()
	if err != nil {
		return nil, err
	}
	var entries []Cmd
	for _, segment := range segments {
		entries = append(entries, segment.Entries...)
	}
	return entries, nil
}

func (j *Journal) Segments() ([]Segment, error) {
	seqs, err := segmentSeqs(j.dir)
	if err != nil {
		return nil, err
	}
	segments := make([]Segment, 0, len(seqs))
	for _, seq := range seqs {
		path := segmentPath(j.dir, seq)
		entries, err := readSegment(path)
		if err != nil {
			return nil, err
		}
		segments = append(segments, Segment{Seq: seq, Path: path, Entries: entries})
	}
	return segments, nil
}

func (j *Journal) activeSegment() (int, string, error) {
	seqs, err := segmentSeqs(j.dir)
	if err != nil {
		return 0, "", err
	}
	if len(seqs) == 0 {
		path := segmentPath(j.dir, 1)
		return 1, path, j.ensureSegmentHeader(path)
	}
	seq := seqs[len(seqs)-1]
	path := segmentPath(j.dir, seq)
	info, err := os.Stat(path)
	if err != nil {
		return 0, "", err
	}
	if info.Size() == 0 {
		if err := j.ensureSegmentHeader(path); err != nil {
			return 0, "", err
		}
		info, err = os.Stat(path)
		if err != nil {
			return 0, "", err
		}
	}
	if info.Size() > j.rotateBytes {
		next := seq + 1
		nextPath := segmentPath(j.dir, next)
		if err := j.createEmptySegment(nextPath); err != nil {
			return 0, "", err
		}
		return next, nextPath, nil
	}
	return seq, path, nil
}

func (j *Journal) rotateAfterAppend(seq int, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.Size() <= j.rotateBytes {
		return nil
	}
	nextPath := segmentPath(j.dir, seq+1)
	if _, err := os.Stat(nextPath); err == nil {
		return nil
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	crashpoint.Hit("pre_segment_rotate")
	if err := j.createEmptySegment(nextPath); err != nil {
		return err
	}
	crashpoint.Hit("post_segment_rotate")
	return nil
}

func segmentSeqs(dir string) ([]int, error) {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var seqs []int
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".jsonl") {
			continue
		}
		seq, err := strconv.Atoi(strings.TrimSuffix(entry.Name(), ".jsonl"))
		if err != nil {
			continue
		}
		seqs = append(seqs, seq)
	}
	sort.Ints(seqs)
	return seqs, nil
}

func readSegment(path string) ([]Cmd, error) {
	entries, _, err := readSegmentWithHighWater(path)
	return entries, err
}

func readSegmentWithHighWater(path string) ([]Cmd, int, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	var entries []Cmd
	var highWater int
	lines := bytes.Split(data, []byte("\n"))
	completeTail := bytes.HasSuffix(data, []byte("\n"))
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		var probe segmentHeader
		if err := json.Unmarshal(line, &probe); err == nil && probe.SegmentHeader {
			if probe.HighWater > highWater {
				highWater = probe.HighWater
			}
			continue
		}
		var cmd Cmd
		if err := json.Unmarshal(line, &cmd); err != nil {
			if i == len(lines)-1 && !completeTail {
				break
			}
			return nil, 0, err
		}
		if n := intakeNumber(cmd.IntakeID); n > highWater {
			highWater = n
		}
		entries = append(entries, cmd)
	}
	return entries, highWater, nil
}

func segmentPath(dir string, seq int) string {
	return filepath.Join(dir, fmt.Sprintf("%06d.jsonl", seq))
}

func (j *Journal) createEmptySegment(path string) error {
	return j.ensureSegmentHeader(path)
}

func (j *Journal) ensureSegmentHeader(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	info, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return err
	}
	if info.Size() == 0 {
		highWater, err := j.highWater()
		if err != nil {
			_ = f.Close()
			return err
		}
		data, err := json.Marshal(segmentHeader{SegmentHeader: true, HighWater: highWater})
		if err != nil {
			_ = f.Close()
			return err
		}
		data = append(data, '\n')
		if err := fsio.AppendFsync(f, data); err != nil {
			_ = f.Close()
			return err
		}
	}
	return f.Close()
}

func (j *Journal) nextIntakeNumber() (int, error) {
	highWater, err := j.highWater()
	if err != nil {
		return 0, err
	}
	return highWater + 1, nil
}

func (j *Journal) highWater() (int, error) {
	seqs, err := segmentSeqs(j.dir)
	if err != nil {
		return 0, err
	}
	var highWater int
	for _, seq := range seqs {
		_, segmentHighWater, err := readSegmentWithHighWater(segmentPath(j.dir, seq))
		if err != nil {
			return 0, err
		}
		if segmentHighWater > highWater {
			highWater = segmentHighWater
		}
	}
	return highWater, nil
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
