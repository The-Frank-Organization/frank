package channel

import (
	"net"
	"testing"
	"time"
)

func TestPushWriteDropsStalledRecipient(t *testing.T) {
	writer, reader := net.Pipe()
	defer func() { _ = writer.Close() }()
	defer func() { _ = reader.Close() }()
	conn := &serverConn{server: &Server{limit: defaultFrameLimit}, conn: writer}

	done := make(chan error, 1)
	go func() {
		done <- writePushes([]*serverConn{conn}, []byte(`{"kind":"delivery-nudge","relay_id":"relay-stalled"}`))
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("writePushes returned error: %v", err)
		}
	case <-time.After(200 * time.Millisecond):
		_ = reader.Close()
		_ = writer.Close()
		t.Fatalf("writePushes blocked on a stalled recipient")
	}
}
