package record_test

import (
	"bytes"
	"testing"

	"github.com/jackli/frank/internal/record"
)

func TestSealVerifyChecksumRoundTrip(t *testing.T) {
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID:       "r1",
			DispatchID:    "d1",
			From:          "seat-a.implementer",
			Role:          "implementer",
			DeliveryState: "accepted",
			IntakeID:      "i1",
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "hello"},
		Body:    "body",
		XFields: map[string]string{"X-demo": "value"},
	}

	data, err := record.Seal(rec)
	if err != nil {
		t.Fatalf("Seal: %v", err)
	}
	verified, err := record.Verify(data)
	if err != nil {
		t.Fatalf("Verify sealed record: %v", err)
	}
	if verified.Checksum == "" {
		t.Fatalf("checksum was not populated")
	}
	if verified.Envelope.RelayID != rec.Envelope.RelayID {
		t.Fatalf("relay id = %q, want %q", verified.Envelope.RelayID, rec.Envelope.RelayID)
	}

	tampered := bytes.Replace(data, []byte("hello"), []byte("hullo"), 1)
	if _, err := record.Verify(tampered); err == nil {
		t.Fatalf("Verify accepted a tampered record")
	}
}
