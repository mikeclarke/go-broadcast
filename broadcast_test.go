package broadcast

import (
	"testing"
	"time"
)

func TestBroadcaster(t *testing.T) {
	// Create testing harness
	done := make(chan Message)

	// Initialize broadcaster
	b := NewBroadcaster()

	// Start listener
	outputc := b.Listen()

	// Func to listen for output
	go func() {
		msg := <-outputc
		done <- msg
	}()

	// Send that shit
	b.Write(1)

	// Wait for results to show up
	select {
	case msg := <- done:
		if msg != 1 {
			t.Fatalf("msg does not match")
		}
	case <- time.After(2 * time.Second):
		t.Fatalf("Messages timed out.")
	}
}
