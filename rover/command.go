package rover

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
)

type Command interface {
	// LogPosition Print out the coordinates and facing direction. A bool argument if it is the starting position
	LogPosition(bool)
	Execute()
}

type noCmd struct{}

func (c *noCmd) LogPosition(_ bool) { fmt.Println("N/A") }
func (c *noCmd) Execute()           { fmt.Println("NO COMMAND") }

var NoCommand = &noCmd{}

type cmdF struct {
	Factor int
	R      *rover
}
type cmdB struct {
	Factor int
	R      *rover
}
type cmdL struct {
	R *rover
}
type cmdR struct {
	R *rover
}

func (c *cmdF) LogPosition(isStart bool) {
	logPosition(isStart, "F", c.R)
}

// Execute Apply command F
func (c *cmdF) Execute() {

	switch c.R.Facing {
	case DIR_N:
		c.R.Y += c.Factor
	case DIR_S:
		c.R.Y -= c.Factor
	case DIR_E:
		c.R.X += c.Factor
	case DIR_W:
		c.R.X -= c.Factor
	}

}

func (c *cmdB) LogPosition(isStart bool) {
	logPosition(isStart, "B", c.R)
}

// Execute Apply the command B
func (c *cmdB) Execute() {

	switch c.R.Facing {
	case DIR_N:
		c.R.Y -= c.Factor
	case DIR_S:
		c.R.Y += c.Factor
	case DIR_E:
		c.R.X -= c.Factor
	case DIR_W:
		c.R.X += c.Factor
	}

}

func (c *cmdL) LogPosition(isStart bool) {
	logPosition(isStart, "L", c.R)
}

func (c *cmdL) Execute() {

	switch c.R.Facing {
	case DIR_N:
		c.R.Facing = DIR_W
	case DIR_S:
		c.R.Facing = DIR_E
	case DIR_E:
		c.R.Facing = DIR_N
	case DIR_W:
		c.R.Facing = DIR_S
	}

}

func (c *cmdR) LogPosition(isStart bool) {
	logPosition(isStart, "R", c.R)
}

func (c *cmdR) Execute() {

	switch c.R.Facing {
	case DIR_N:
		c.R.Facing = DIR_E
	case DIR_S:
		c.R.Facing = DIR_W
	case DIR_E:
		c.R.Facing = DIR_S
	case DIR_W:
		c.R.Facing = DIR_N
	}

}

// Commands Receives a string of commands and a pointer to a rover. Returns a read-only channel where it writes the commnds.
// The function returns immediately and launches a separate goroutine to write to the channel. The goroutine exists when the work is done or when the upstream context was canceled. Upon exit it also closes the channel, signaling no more commands will be generated.
func Commands(ctx context.Context, input io.Reader, rvr *rover) <-chan Command {

	ch := make(chan Command)

	go func() {
		defer close(ch)

		scanner := bufio.NewScanner(input)
		scanner.Split(onCommand)

		for scanner.Scan() {
			txt := scanner.Text()

			if txt == "X" {
				break
			}

			cmd := NewCommand(txt, rvr)

			select {
			case <-ctx.Done():
				return
			case ch <- cmd:
			}
		}
	}()

	return ch
}

// NewCommand Decodes string of form `10F` where the last rune is the command identifier and the preceding number is the multiplication factor
func NewCommand(txt string, rvr *rover) Command {

	r := rune(txt[len(txt)-1])
	factor, _ := strconv.Atoi(txt[:len(txt)-1])

	switch r {
	case 'F':
		if factor == 0 {
			// set default factor to 1 instead of 0
			factor = 1
		}
		return &cmdF{factor, rvr}
	case 'B':
		if factor == 0 {
			// set default factor to 1 instead of 0
			factor = 1
		}
		return &cmdB{factor, rvr}
	case 'L':
		return &cmdL{rvr}
	case 'R':
		return &cmdR{rvr}
	default:
		return NoCommand
	}
}
