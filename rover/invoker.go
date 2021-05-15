package rover

import (
	"context"
	"time"
)

// ExecuteAll Listens to a read-only channel for incoming Commands
// It has an option delay (can be 0) that is a sleep between commands. Can be used to simulate a slower process.
// The function returns immediately when the context is cancelled or the incoming channel is closed
func ExecuteAll(ctx context.Context, ch <-chan Command, delay time.Duration) {

	for {
		select {
		case c, ok := <-ch:
			if !ok {
				// channel closed
				return
			}
			c.LogPosition(true)
			c.Execute()
			c.LogPosition(false)
			if delay > 0 {
				// Insert a delay between the commands for easier testing
				// Use a select to stop immediately when an interrupt arrives (not wait for the delay between the commands)
				select {
				case <-time.Tick(delay):
				case <-ctx.Done():
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
