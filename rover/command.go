package rover

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

type Command interface {
	Execute()
}

type noCmd struct{}

func (c *noCmd) Execute() { fmt.Println("NO COMMAND") }

var NoCommand = &noCmd{}

type cmdF struct {
	R *rover
}

func (c *cmdF) Execute() {

	fmt.Printf("F: %s -> ", c.R)

	switch c.R.Facing {
	case DIR_N:
		c.R.Y++
	case DIR_S:
		c.R.Y--
	case DIR_E:
		c.R.X++
	case DIR_W:
		c.R.X--
	}

	fmt.Printf("%s\n", c.R)
}

type cmdB struct {
	R *rover
}
type cmdL struct {
	R *rover
}
type cmdR struct {
	R *rover
}

func (c *cmdB) Execute() {

	fmt.Printf("B: %s -> ", c.R)

	switch c.R.Facing {
	case DIR_N:
		c.R.Y--
	case DIR_S:
		c.R.Y++
	case DIR_E:
		c.R.X--
	case DIR_W:
		c.R.X++
	}

	fmt.Printf("%s\n", c.R)
}

func (c *cmdL) Execute() {

	fmt.Printf("L: %s -> ", c.R)

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

	fmt.Printf("%s\n", c.R)
}

func (c *cmdR) Execute() {

	fmt.Printf("R: %s -> ", c.R)

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

	fmt.Printf("%s\n", c.R)
}

// Commands Receives a string of commands and a pointer to a rover. Returns a read-only channel where it writes the commnds.
// The function returns immediately and launches a separate goroutine to write to the channel. The goroutine exists when the work is done or when the upstream context was canceled. Upon exit it also closes the channel, signaling no more commands will be generated.
func Commands(ctx context.Context, input io.Reader, rvr *rover) <-chan Command {

	ch := make(chan Command)

	go func() {
		defer close(ch)

		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanRunes)

		for scanner.Scan() {
			r := rune(strings.ToUpper(scanner.Text())[0])

			if r == 'X' {
				break
			}

			if r == '\n' {
				continue
			}

			cmd := decodeCommand(r, rvr)

			select {
			case <-ctx.Done():
				return
			case ch <- cmd:
			}
		}
	}()

	return ch
}

func decodeCommand(r rune, rvr *rover) Command {
	switch r {
	case 'F':
		return &cmdF{rvr}
	case 'B':
		return &cmdB{rvr}
	case 'L':
		return &cmdL{rvr}
	case 'R':
		return &cmdR{rvr}
	default:
		return NoCommand
	}
}
