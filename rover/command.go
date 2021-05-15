package rover

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Command interface {
	Execute()
}

type noCmd struct{}

func (c *noCmd) Execute() { fmt.Println("NO COMMAND") }

var NoCommand = &noCmd{}

type cmdF struct {
	Factor int
	R      *rover
}

// Execute Apply command F
func (c *cmdF) Execute() {

	fmt.Printf("F: %s -> ", c.R)

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

	fmt.Printf("%s\n", c.R)
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

// Execute Apply the command B
// TODO: Implement the factor
func (c *cmdB) Execute() {

	fmt.Printf("B: %s -> ", c.R)

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
		scanner.Split(onCommand)

		for scanner.Scan() {
			txt := scanner.Text()

			if txt == "X" {
				break
			}

			cmd := decodeCommand(txt, rvr)

			select {
			case <-ctx.Done():
				return
			case ch <- cmd:
			}
		}
	}()

	return ch
}

// decodeCommand Decodes string of form `10F` where the last rune is the command identifier and the preceding number is the multiplication factor
func decodeCommand(txt string, rvr *rover) Command {

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

// onCommand A SplitFunc that filters out bad commands. It returns tokens for only valid commands
func onCommand(data []byte, atEOF bool) (advance int, token []byte, err error) {

	data = []byte(strings.ToUpper(string(data)))
	var start int

	for i := 0; i < len(data); i++ {
		if isValidCommand(data[i]) {
			return i + 1, data[start : i+1], nil
		}

		// if it is not a size for the command
		if !(data[i] >= '0' && data[i] <= '9') {
			start = i + 1
		}
	}

	// no command found yet, try with longer input if not at the end already
	if !atEOF {
		return 0, nil, nil
	}

	return 0, nil, io.EOF
}

func isValidCommand(data byte) bool {
	return data == 'F' || data == 'B' || data == 'R' || data == 'L' || data == 'X'
}
