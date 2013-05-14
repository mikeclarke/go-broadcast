package broadcast

import (
	"log"
	"testing"
	"time"
)

func TestBroadcaster(t *testing.T) {
	// Create testing harness
	done := make(chan bool)

	// Initialize broadcaster
	b := NewBroadcaster()

	// Start listener
	outputc := b.Listen()

	// Func to listen for output
	go func() {
		msg := <-outputc
		log.Print(msg)
		done <- true
	}()

	// Send that shit
	b.Write(1)

	// Wait for results to show up
	select {
	case <- done:
		return
	case <- time.After(2 * time.Second):
		t.Fatalf("Messages timed out.")
	}
}
