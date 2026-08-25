package appipc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"
)

func TestFrameRoundTripUsesUint32BigEndianLength(t *testing.T) {
	var wire bytes.Buffer
	payload := []byte(`{"v":1}`)
	if err := WriteFrame(&wire, payload); err != nil {
		t.Fatalf("WriteFrame: %v", err)
	}
	wantWire := append([]byte{0, 0, 0, byte(len(payload))}, payload...)
	if !bytes.Equal(wire.Bytes(), wantWire) {
		t.Fatalf("wire = %x, want %x", wire.Bytes(), wantWire)
	}
	got, err := ReadFrame(&wire)
	if err != nil {
		t.Fatalf("ReadFrame: %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("payload = %q, want %q", got, payload)
	}
}

func TestFrameAcceptsExactMaximum(t *testing.T) {
	payload := bytes.Repeat([]byte{'x'}, FrameMax)
	var wire bytes.Buffer
	if err := WriteFrame(&wire, payload); err != nil {
		t.Fatalf("WriteFrame(exact maximum): %v", err)
	}
	got, err := ReadFrame(&wire)
	if err != nil {
		t.Fatalf("ReadFrame(exact maximum): %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("exact-maximum payload changed")
	}
}

func TestFrameRejectsOversizeBeforeWritingOrAllocatingPayload(t *testing.T) {
	var wire bytes.Buffer
	if err := WriteFrame(&wire, make([]byte, FrameMax+1)); !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("WriteFrame(over maximum) error = %v, want ErrFrameTooLarge", err)
	}
	if wire.Len() != 0 {
		t.Fatalf("WriteFrame(over maximum) wrote %d bytes", wire.Len())
	}

	var header [4]byte
	binary.BigEndian.PutUint32(header[:], uint32(FrameMax+1))
	if _, err := ReadFrame(bytes.NewReader(header[:])); !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("ReadFrame(over maximum header) error = %v, want ErrFrameTooLarge", err)
	}
}

func TestFrameReportsTruncation(t *testing.T) {
	for _, wire := range [][]byte{
		{0, 0, 0},
		{0, 0, 0, 2, 'x'},
	} {
		if _, err := ReadFrame(bytes.NewReader(wire)); !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatalf("ReadFrame(%x) error = %v, want io.ErrUnexpectedEOF", wire, err)
		}
	}
}
