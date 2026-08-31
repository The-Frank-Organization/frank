package journal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/The-Frank-Organization/frank/internal/worker/wire"
)

const (
	SessionLogName  = "session.log"
	SessionLockName = "session.lock"
	RunFresh        = "fresh"
	RunResume       = "resume"
)

var (
	ErrFenceHeld        = errors.New("journal writer fence is held")
	ErrResumeLogAbsent  = errors.New("resume journal is absent")
	ErrGenesisIdentity  = errors.New("journal genesis identity fault")
	ErrDescriptorSafety = errors.New("journal descriptor safety fault")
	ErrDurabilityOrder  = errors.New("journal durability order violation")
)

// GenesisCommit is emitted only after the genesis bytes and parent directory
// are durable. The caller maps it to the one-way CTRL-W notification.
type GenesisCommit struct {
	RunID        string
	TurnEpoch    string
	GenerationID string
}

// Config is the turn_open-carried writer configuration for one run.
type Config struct {
	RuntimeDir       string
	RunDisposition   string
	Identity         Identity
	GenerationID     string
	TurnEpoch        string
	OnGenesisDurable func(GenesisCommit) error
}

// Writer owns the held per-run fence and the open journal descriptor.
type Writer struct {
	mu           sync.Mutex
	runtimeDir   string
	logPath      string
	identity     Identity
	generationID string
	turnEpoch    string
	parent       *os.File
	lock         *os.File
	log          *os.File
	fileIdentity objectIdentity
	nextSeq      uint64
	closed       bool
}

type objectIdentity struct {
	device uint64
	inode  uint64
}

// IntakeHighWater admits turn_open frames within one CTRL-W receive
// incarnation. Its zero value is the required UNSET state.
type IntakeHighWater struct {
	set   bool
	value uint64
}

// Admit advances the incarnation-local high-water or drops a duplicate or
// regression before any create/reopen decision.
func (highWater *IntakeHighWater) Admit(seq string) (bool, error) {
	value, err := wire.ParseCounter(seq)
	if err != nil {
		return false, err
	}
	if !highWater.set {
		highWater.set = true
		highWater.value = value
		return true, nil
	}
	if value <= highWater.value {
		return false, nil
	}
	highWater.value = value
	return true, nil
}

// Open acquires the exclusive fence and takes exactly one create or uniform
// resume path. A resume never creates a missing log.
func Open(config Config) (*Writer, error) {
	if config.RuntimeDir == "" || config.GenerationID == "" {
		return nil, errors.New("journal runtime directory or generation_id is empty")
	}
	if !filepath.IsAbs(config.RuntimeDir) {
		return nil, ErrDescriptorSafety
	}
	if config.RunDisposition != RunFresh && config.RunDisposition != RunResume {
		return nil, errors.New("run_disposition is outside the closed enum")
	}
	if err := validateIdentity(config.Identity); err != nil {
		return nil, err
	}
	if _, err := wire.ParseCounter(config.TurnEpoch); err != nil {
		return nil, errors.New("turn_epoch is not a canonical counter")
	}

	parent, err := openPrivateDirectory(config.RuntimeDir)
	if err != nil {
		return nil, err
	}
	writer := &Writer{
		runtimeDir:   config.RuntimeDir,
		logPath:      filepath.Join(config.RuntimeDir, SessionLogName),
		identity:     config.Identity,
		generationID: config.GenerationID,
		turnEpoch:    config.TurnEpoch,
		parent:       parent,
	}
	cleanup := func(openErr error) (*Writer, error) {
		_ = writer.closeUnlocked()
		return nil, openErr
	}

	lock, err := openLock(filepath.Join(config.RuntimeDir, SessionLockName))
	if err != nil {
		return cleanup(err)
	}
	writer.lock = lock
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		if errors.Is(err, syscall.EWOULDBLOCK) || errors.Is(err, syscall.EAGAIN) {
			return cleanup(ErrFenceHeld)
		}
		return cleanup(fmt.Errorf("acquire journal fence: %w", err))
	}

	_, statErr := os.Lstat(writer.logPath)
	present := statErr == nil
	if statErr != nil && !os.IsNotExist(statErr) {
		return cleanup(ErrDescriptorSafety)
	}
	if !present && config.RunDisposition == RunResume {
		return cleanup(ErrResumeLogAbsent)
	}
	if !present {
		if err := writer.createGenesis(config.OnGenesisDurable); err != nil {
			return cleanup(err)
		}
		return writer, nil
	}
	if err := writer.resume(config.OnGenesisDurable); err != nil {
		return cleanup(err)
	}
	return writer, nil
}

func (writer *Writer) createGenesis(callback func(GenesisCommit) error) error {
	fd, err := syscall.Open(writer.logPath, syscall.O_CREAT|syscall.O_EXCL|syscall.O_RDWR|syscall.O_NOFOLLOW|syscall.O_CLOEXEC, 0o600)
	if err != nil {
		return fmt.Errorf("exclusive journal create: %w", err)
	}
	writer.log = os.NewFile(uintptr(fd), writer.logPath)
	identity, err := verifyOpenJournal(writer.log)
	if err != nil {
		return err
	}
	writer.fileIdentity = identity
	record := Record{
		Seq: "0", Kind: KindRunOpen, GenerationID: writer.generationID, TSMonotonic: "0",
		Fields: map[string]json.RawMessage{
			"run_id":              marshalString(writer.identity.RunID),
			"run_manifest_digest": marshalString(writer.identity.RunManifestDigest),
			"turn_epoch":          marshalString(writer.turnEpoch),
			"create_auth_id":      marshalString(writer.identity.CreateAuthID),
		},
	}
	finalized, err := FinalizeRecord(record)
	if err != nil {
		return err
	}
	if err := writeRecordLine(writer.log, finalized); err != nil {
		return err
	}
	if err := writer.log.Sync(); err != nil {
		return fmt.Errorf("sync genesis record: %w", err)
	}
	if err := writer.parent.Sync(); err != nil {
		return fmt.Errorf("sync runtime directory: %w", err)
	}
	if err := writer.emitGenesisDurable(callback); err != nil {
		return err
	}
	writer.nextSeq = 1
	return nil
}

func (writer *Writer) resume(callback func(GenesisCommit) error) error {
	fd, err := syscall.Open(writer.logPath, syscall.O_RDWR|syscall.O_NOFOLLOW|syscall.O_CLOEXEC, 0)
	if err != nil {
		return ErrDescriptorSafety
	}
	writer.log = os.NewFile(uintptr(fd), writer.logPath)
	identity, err := verifyOpenJournal(writer.log)
	if err != nil {
		return err
	}
	writer.fileIdentity = identity
	data, err := io.ReadAll(writer.log)
	if err != nil {
		return fmt.Errorf("read journal: %w", err)
	}
	if err := verifyGenesisLine(data, writer.identity); err != nil {
		return err
	}
	// Uniform resume re-establishes file and parent-directory durability before
	// the idempotent genesis notification.
	if err := writer.log.Sync(); err != nil {
		return fmt.Errorf("re-sync genesis: %w", err)
	}
	if err := writer.parent.Sync(); err != nil {
		return fmt.Errorf("re-sync runtime directory: %w", err)
	}
	if err := writer.emitGenesisDurable(callback); err != nil {
		return err
	}
	recovery := Recover(data, writer.identity)
	if recovery.GenesisFault {
		return ErrGenesisIdentity
	}
	if err := recovery.validateForTruncation(); err != nil {
		return err
	}
	if err := writer.log.Truncate(recovery.Boundary.Offset); err != nil {
		return fmt.Errorf("truncate untrusted journal tail: %w", err)
	}
	if err := writer.log.Sync(); err != nil {
		return fmt.Errorf("sync journal truncation: %w", err)
	}
	if _, err := writer.log.Seek(0, io.SeekEnd); err != nil {
		return fmt.Errorf("seek journal append point: %w", err)
	}
	writer.nextSeq = recovery.NextSeq
	return nil
}

func (writer *Writer) emitGenesisDurable(callback func(GenesisCommit) error) error {
	if callback == nil {
		return nil
	}
	return callback(GenesisCommit{
		RunID: writer.identity.RunID, TurnEpoch: writer.turnEpoch, GenerationID: writer.generationID,
	})
}

// Append durably appends one record. Round members remain provisional until a
// later AppendRound marker admits them.
func (writer *Writer) Append(record Record) (Record, error) {
	writer.mu.Lock()
	defer writer.mu.Unlock()
	return writer.appendLocked(record)
}

func (writer *Writer) appendLocked(record Record) (Record, error) {
	if writer.closed || writer.log == nil {
		return Record{}, errors.New("journal writer is closed")
	}
	if record.Kind == KindRunOpen {
		return Record{}, errors.New("run_open can only be created at genesis")
	}
	if err := writer.verifyDescriptor(); err != nil {
		return Record{}, err
	}
	record = cloneRecord(record)
	record.Seq = wire.FormatCounter(writer.nextSeq)
	record.GenerationID = writer.generationID
	record.TSMonotonic = wire.FormatCounter(writer.nextSeq)
	finalized, err := FinalizeRecord(record)
	if err != nil {
		return Record{}, err
	}
	if err := writeRecordLine(writer.log, finalized); err != nil {
		return Record{}, err
	}
	if err := writer.log.Sync(); err != nil {
		return Record{}, fmt.Errorf("sync appended record: %w", err)
	}
	writer.nextSeq++
	return finalized, nil
}

// RoundCommit is an unforgeable-in-practice proof returned only after all
// content records and their admitting marker have fsync-linearized.
type RoundCommit struct {
	markerSeq   uint64
	toolCallIDs map[string]struct{}
}

// AppendRound durably appends each content record, then its admitting marker.
func (writer *Writer) AppendRound(turnID, roundIndex string, members []Record) (RoundCommit, error) {
	writer.mu.Lock()
	defer writer.mu.Unlock()
	if len(members) == 0 {
		return RoundCommit{}, errors.New("round has no members")
	}
	if _, err := wire.ParseCounter(roundIndex); err != nil {
		return RoundCommit{}, errors.New("round_index is not canonical")
	}
	appended := make([]Record, 0, len(members))
	toolCallIDs := make(map[string]struct{})
	for _, member := range members {
		if !admitEligible(member.Kind) {
			return RoundCommit{}, fmt.Errorf("record kind %q is not round-admit-eligible", member.Kind)
		}
		member.TurnID = turnID
		member.RoundIndex = roundIndex
		written, err := writer.appendLocked(member)
		if err != nil {
			return RoundCommit{}, err
		}
		appended = append(appended, written)
		if member.Kind == KindToolResult {
			if id, err := rawStringValue(member.Fields["tool_call_id"], "tool_call_id"); err == nil {
				toolCallIDs[id] = struct{}{}
			}
		}
	}
	markerSeq := writer.nextSeq
	marker, err := BuildRoundMarker(
		wire.FormatCounter(markerSeq), writer.generationID, turnID, roundIndex,
		wire.FormatCounter(markerSeq), appended,
	)
	if err != nil {
		return RoundCommit{}, err
	}
	if _, err := writer.appendLocked(marker); err != nil {
		return RoundCommit{}, err
	}
	return RoundCommit{markerSeq: markerSeq, toolCallIDs: toolCallIDs}, nil
}

// RequireToolOutcome rejects an outcome unless its exact tool-result identity
// is covered by a marker-durable RoundCommit.
func RequireToolOutcome(commit RoundCommit, toolCallID string) error {
	if commit.markerSeq == 0 || commit.toolCallIDs == nil {
		return ErrDurabilityOrder
	}
	if _, covered := commit.toolCallIDs[toolCallID]; !covered {
		return ErrDurabilityOrder
	}
	return nil
}

// Close releases the fence only after closing the journal descriptor.
func (writer *Writer) Close() error {
	writer.mu.Lock()
	defer writer.mu.Unlock()
	return writer.closeUnlocked()
}

func (writer *Writer) closeUnlocked() error {
	if writer.closed {
		return nil
	}
	writer.closed = true
	var first error
	if writer.log != nil {
		if err := writer.log.Close(); err != nil && first == nil {
			first = err
		}
		writer.log = nil
	}
	if writer.lock != nil {
		if err := syscall.Flock(int(writer.lock.Fd()), syscall.LOCK_UN); err != nil && first == nil {
			first = err
		}
		if err := writer.lock.Close(); err != nil && first == nil {
			first = err
		}
		writer.lock = nil
	}
	if writer.parent != nil {
		if err := writer.parent.Close(); err != nil && first == nil {
			first = err
		}
		writer.parent = nil
	}
	return first
}

func (writer *Writer) verifyDescriptor() error {
	if err := verifyPrivateDirectoryFile(writer.parent); err != nil {
		return err
	}
	current, err := verifyOpenJournal(writer.log)
	if err != nil {
		return err
	}
	if current != writer.fileIdentity {
		return ErrDescriptorSafety
	}
	fd, err := syscall.Open(writer.logPath, syscall.O_RDWR|syscall.O_NOFOLLOW|syscall.O_CLOEXEC, 0)
	if err != nil {
		return ErrDescriptorSafety
	}
	pathFile := os.NewFile(uintptr(fd), writer.logPath)
	pathIdentity, verifyErr := verifyOpenJournal(pathFile)
	closeErr := pathFile.Close()
	if verifyErr != nil {
		return verifyErr
	}
	if closeErr != nil {
		return closeErr
	}
	if pathIdentity != writer.fileIdentity {
		return ErrDescriptorSafety
	}
	return nil
}

func openPrivateDirectory(path string) (*os.File, error) {
	fd, err := syscall.Open(path, syscall.O_RDONLY|syscall.O_DIRECTORY|syscall.O_NOFOLLOW|syscall.O_CLOEXEC, 0)
	if err != nil {
		return nil, ErrDescriptorSafety
	}
	directory := os.NewFile(uintptr(fd), path)
	if err := verifyPrivateDirectoryFile(directory); err != nil {
		_ = directory.Close()
		return nil, err
	}
	return directory, nil
}

func verifyPrivateDirectoryFile(directory *os.File) error {
	var stat syscall.Stat_t
	if err := syscall.Fstat(int(directory.Fd()), &stat); err != nil {
		return ErrDescriptorSafety
	}
	if stat.Mode&syscall.S_IFMT != syscall.S_IFDIR || stat.Uid != uint32(os.Geteuid()) || stat.Mode&0o777 != 0o700 {
		return ErrDescriptorSafety
	}
	return nil
}

func openLock(path string) (*os.File, error) {
	fd, err := syscall.Open(path, syscall.O_CREAT|syscall.O_RDWR|syscall.O_NOFOLLOW|syscall.O_CLOEXEC, 0o600)
	if err != nil {
		return nil, ErrDescriptorSafety
	}
	file := os.NewFile(uintptr(fd), path)
	var stat syscall.Stat_t
	if err := syscall.Fstat(fd, &stat); err != nil {
		_ = file.Close()
		return nil, ErrDescriptorSafety
	}
	if err := validateJournalStat(&stat, uint32(os.Geteuid())); err != nil {
		_ = file.Close()
		return nil, err
	}
	return file, nil
}

func verifyOpenJournal(file *os.File) (objectIdentity, error) {
	var stat syscall.Stat_t
	if err := syscall.Fstat(int(file.Fd()), &stat); err != nil {
		return objectIdentity{}, ErrDescriptorSafety
	}
	if err := validateJournalStat(&stat, uint32(os.Geteuid())); err != nil {
		return objectIdentity{}, err
	}
	return objectIdentity{device: uint64(stat.Dev), inode: uint64(stat.Ino)}, nil
}

func validateJournalStat(stat *syscall.Stat_t, euid uint32) error {
	if stat.Mode&syscall.S_IFMT != syscall.S_IFREG || stat.Uid != euid || stat.Mode&0o777 != 0o600 || stat.Nlink != 1 {
		return ErrDescriptorSafety
	}
	return nil
}

func verifyGenesisLine(data []byte, identity Identity) error {
	newline := bytes.IndexByte(data, '\n')
	if newline < 0 {
		return ErrGenesisIdentity
	}
	record, err := DecodeRecord(data[:newline])
	if err != nil || record.Kind != KindRunOpen || !matchesGenesis(record, identity) {
		return ErrGenesisIdentity
	}
	return nil
}

func writeRecordLine(file *os.File, record Record) error {
	encoded, err := MarshalRecord(record)
	if err != nil {
		return err
	}
	line := append(encoded, '\n')
	for len(line) > 0 {
		written, err := file.Write(line)
		if err != nil {
			return fmt.Errorf("append journal record: %w", err)
		}
		if written == 0 {
			return io.ErrShortWrite
		}
		line = line[written:]
	}
	return nil
}

func cloneRecord(record Record) Record {
	cloned := record
	cloned.Fields = make(map[string]json.RawMessage, len(record.Fields))
	for member, raw := range record.Fields {
		cloned.Fields[member] = append(json.RawMessage(nil), raw...)
	}
	return cloned
}
