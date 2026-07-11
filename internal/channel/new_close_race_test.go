package channel

import (
	"encoding/json"
	"net"
	"sync"
	"testing"
)

func TestClientCloseIdempotentUnderRace(t *testing.T) {
	for range 100 {
		clientConn, serverConn := net.Pipe()
		client := &Client{
			conn:    clientConn,
			enc:     json.NewEncoder(clientConn),
			pending: map[int64]chan rpcMessage{},
			pushes:  make(chan []byte, 1),
			done:    make(chan struct{}),
			limit:   defaultFrameLimit,
		}
		go client.readLoop()

		start := make(chan struct{})
		var workers sync.WaitGroup
		workers.Add(3)
		go func() {
			defer workers.Done()
			<-start
			_ = serverConn.Close()
		}()
		for range 2 {
			go func() {
				defer workers.Done()
				<-start
				_ = client.Close()
			}()
		}
		close(start)
		workers.Wait()
	}
}
