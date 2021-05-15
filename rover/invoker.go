package rover

import (
	"context"
	"time"
)

func ExecuteAll(ctx context.Context, ch <-chan Command, delay time.Duration) {

	for {
		select {
		case c, ok := <-ch:
			if !ok {
				// channel closed
				return
			}
			c.Execute()
			if delay > 0 {
				// Insert a delay between the commands for easier testing
				// Use a select to stop immediately when an interrupt arrives (not wait for the delay between the commands)
				select {
				case <-time.Tick(delay):
				case <-ctx.Done():
					break
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
