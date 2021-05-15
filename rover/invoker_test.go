package rover

import (
	"context"
	"testing"
	"time"
)

type testCmd struct {
	counter int
}

func (c *testCmd) LogPosition(_ bool) {}
func (c *testCmd) Execute() {
	c.counter++
}

func TestExecuteAllWithDelay(t *testing.T) {

	total := 10
	cmd := &testCmd{0}
	ch := make(chan Command)
	delay := 100 * time.Millisecond

	go func() {
		for i := 0; i < total; i++ {
			ch <- cmd
		}

		close(ch)
	}()

	start := time.Now()
	ExecuteAll(context.Background(), ch, delay)
	duration := time.Since(start)

	if total != cmd.counter {
		t.Errorf("Not all commands have run. Expected: %d, got: %d", total, cmd.counter)
	}

	if duration < time.Duration(int(delay)*total) {
		t.Error("Delay not applied between commands")
	}

}
