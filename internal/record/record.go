package record

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

const (
	Accepted = "accepted"
	Rejected = "rejected"
	Held     = "held"
)

type Envelope struct {
	RelayID       string `json:"relay_id"`
	DispatchID    string `json:"dispatch_id"`
	From          string `json:"from"`
	To            string `json:"to,omitempty"`
	Role          string `json:"role"`
	DeliveryState string `json:"delivery_state"`
	IntakeID      string `json:"intake_id,omitempty"`
	SchemaVersion int    `json:"schema_version"`
}

type Record struct {
	Envelope Envelope          `json:"envelope"`
	Headers  map[string]string `json:"headers"`
	Body     string            `json:"body"`
	XFields  map[string]string `json:"x_fields,omitempty"`
	Checksum string            `json:"checksum"`
}

func Seal(r Record) ([]byte, error) {
	r.Checksum = ""
	sum, err := checksum(r)
	if err != nil {
		return nil, err
	}
	r.Checksum = sum
	return json.MarshalIndent(r, "", "  ")
}

func Verify(data []byte) (Record, error) {
	var r Record
	if err := json.Unmarshal(data, &r); err != nil {
		return Record{}, err
	}
	if r.Checksum == "" {
		return Record{}, errors.New("missing checksum")
	}
	got := r.Checksum
	r.Checksum = ""
	want, err := checksum(r)
	if err != nil {
		return Record{}, err
	}
	if got != want {
		return Record{}, errors.New("checksum mismatch")
	}
	r.Checksum = got
	return r, nil
}

func checksum(r Record) (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}
