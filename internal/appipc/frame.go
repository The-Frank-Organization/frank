package appipc

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	ErrEmptyFrame    = errors.New("appipc: frame payload is empty")
	ErrFrameTooLarge = errors.New("appipc: frame exceeds FRAME_MAX")
)

// WriteFrame writes uint32_be(length) followed by the payload. Oversized
// payloads are rejected before any bytes reach the writer.
func WriteFrame(writer io.Writer, payload []byte) error {
	if len(payload) == 0 {
		return ErrEmptyFrame
	}
	if len(payload) > FrameMax {
		return fmt.Errorf("%w: %d > %d", ErrFrameTooLarge, len(payload), FrameMax)
	}
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], uint32(len(payload)))
	if err := writeAll(writer, header[:]); err != nil {
		return fmt.Errorf("appipc: write frame header: %w", err)
	}
	if err := writeAll(writer, payload); err != nil {
		return fmt.Errorf("appipc: write frame payload: %w", err)
	}
	return nil
}

// ReadFrame rejects an oversized declared length before allocating or reading
// its payload.
func ReadFrame(reader io.Reader) ([]byte, error) {
	var header [4]byte
	if _, err := io.ReadFull(reader, header[:]); err != nil {
		return nil, fmt.Errorf("appipc: read frame header: %w", err)
	}
	length := binary.BigEndian.Uint32(header[:])
	if length > FrameMax {
		return nil, fmt.Errorf("%w: %d > %d", ErrFrameTooLarge, length, FrameMax)
	}
	payload := make([]byte, int(length))
	if _, err := io.ReadFull(reader, payload); err != nil {
		return nil, fmt.Errorf("appipc: read frame payload: %w", err)
	}
	return payload, nil
}

func writeAll(writer io.Writer, data []byte) error {
	for len(data) > 0 {
		written, err := writer.Write(data)
		if err != nil {
			return err
		}
		if written <= 0 || written > len(data) {
			return io.ErrShortWrite
		}
		data = data[written:]
	}
	return nil
}
