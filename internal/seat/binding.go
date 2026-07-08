package seat

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/jackli/frank/internal/fsio"
	"github.com/jackli/frank/internal/record"
)

var (
	ErrSeatAlreadyBound = errors.New("seat already bound")
	ErrReservedSeatName = errors.New("reserved seat name")
)

type Cred struct {
	Value string
}

type SeatMeta struct {
	Name           string `json:"name"`
	Role           string `json:"role"`
	IsOperator     bool   `json:"is_operator"`
	AuthGeneration string `json:"-"`
}

type Manager struct {
	root  string
	table bindingTable
	mu    sync.Mutex
}

type bindingTable struct {
	Seats map[string]binding `json:"seats"`
}

type binding struct {
	Credential      string   `json:"credential"`
	Meta            SeatMeta `json:"meta"`
	RealizedMintRef string   `json:"realized_mint_ref,omitempty"`
}

func Open(root string) (*Manager, error) {
	if err := os.MkdirAll(filepath.Join(root, "binding"), 0o700); err != nil {
		return nil, err
	}
	m := &Manager{root: root, table: bindingTable{Seats: map[string]binding{}}}
	data, err := os.ReadFile(filepath.Join(root, "binding", "seats.json"))
	if os.IsNotExist(err) {
		return m, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &m.table); err != nil {
		return nil, err
	}
	if m.table.Seats == nil {
		m.table.Seats = map[string]binding{}
	}
	return m, nil
}

func (m *Manager) Mint(name, role string, isOperator bool) (Cred, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if name == "system" {
		return Cred{}, ErrReservedSeatName
	}
	if _, exists := m.table.Seats[name]; exists {
		return Cred{}, ErrSeatAlreadyBound
	}
	value, err := randomCredential()
	if err != nil {
		return Cred{}, err
	}
	m.table.Seats[name] = binding{
		Credential: value,
		Meta:       SeatMeta{Name: name, Role: role, IsOperator: isOperator},
	}
	if err := m.persist(); err != nil {
		delete(m.table.Seats, name)
		return Cred{}, err
	}
	return Cred{Value: value}, nil
}

func (m *Manager) MintOrReplace(name, role string, isOperator bool, realizedMintRef string) (Cred, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if name == "system" {
		return Cred{}, ErrReservedSeatName
	}
	value, err := randomCredential()
	if err != nil {
		return Cred{}, err
	}
	previous, hadPrevious := m.table.Seats[name]
	m.table.Seats[name] = binding{
		Credential:      value,
		Meta:            SeatMeta{Name: name, Role: role, IsOperator: isOperator},
		RealizedMintRef: realizedMintRef,
	}
	if err := m.persist(); err != nil {
		if hadPrevious {
			m.table.Seats[name] = previous
		} else {
			delete(m.table.Seats, name)
		}
		return Cred{}, err
	}
	return Cred{Value: value}, nil
}

func (m *Manager) Resolve(credential string) (SeatMeta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, binding := range m.table.Seats {
		if subtle.ConstantTimeCompare([]byte(binding.Credential), []byte(credential)) == 1 {
			return binding.Meta, true
		}
	}
	return SeatMeta{}, false
}

func (m *Manager) CredentialsFor(name string) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.table.Seats[name]; ok {
		return 1
	}
	return 0
}

func (m *Manager) RealizedMintRef(name string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	binding, ok := m.table.Seats[name]
	if !ok {
		return "", false
	}
	return binding.RealizedMintRef, true
}

func (m *Manager) Seats() []SeatMeta {
	m.mu.Lock()
	defer m.mu.Unlock()
	seats := make([]SeatMeta, 0, len(m.table.Seats))
	for _, binding := range m.table.Seats {
		seats = append(seats, binding.Meta)
	}
	return seats
}

func Stamp(rec record.Record, meta SeatMeta) record.Record {
	rec.Envelope.From = meta.Name
	rec.Envelope.Role = meta.Role
	return rec
}

func (m *Manager) persist() error {
	data, err := json.MarshalIndent(m.table, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(m.root, "binding"), 0o700); err != nil {
		return err
	}
	if err := fsio.WriteFileAtomic(m.root, filepath.Join("binding", "seats.json"), data); err != nil {
		return err
	}
	return os.Chmod(filepath.Join(m.root, "binding", "seats.json"), 0o600)
}

func randomCredential() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
