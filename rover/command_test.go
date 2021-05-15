package rover

import (
	"context"
	"strings"
	"testing"
)

func TestCommandF(t *testing.T) {

	type scenario struct {
		x, y           int
		dir            string
		expectedX      int
		expectedY      int
		expectedFacing direction
	}

	scenarios := []scenario{
		scenario{1, 1, "n", 1, 2, DIR_N},
		scenario{1, 1, "s", 1, 0, DIR_S},
		scenario{1, 1, "e", 2, 1, DIR_E},
		scenario{1, 1, "w", 0, 1, DIR_W},
	}

	for _, s := range scenarios {
		r, _ := NewRover("test", 1, 1, s.dir)
		c := &cmdF{r}
		c.Execute()

		if r.Facing != s.expectedFacing {
			t.Error("Rover should not change direction for cmdF")
		}

		if r.X != s.expectedX {
			t.Error("Incorrect move on X")
		}

		if r.Y != s.expectedY {
			t.Error("Incorrect move on Y")
		}
	}
}

func TestCommandB(t *testing.T) {

	type scenario struct {
		x, y           int
		dir            string
		expectedX      int
		expectedY      int
		expectedFacing direction
	}

	scenarios := []scenario{
		scenario{1, 1, "n", 1, 0, DIR_N},
		scenario{1, 1, "s", 1, 2, DIR_S},
		scenario{1, 1, "e", 0, 1, DIR_E},
		scenario{1, 1, "w", 2, 1, DIR_W},
	}

	for _, s := range scenarios {
		r, _ := NewRover("test", 1, 1, s.dir)
		c := &cmdB{r}
		c.Execute()

		if r.Facing != s.expectedFacing {
			t.Error("Rover should not change direction for cmdB")
		}

		if r.X != s.expectedX {
			t.Error("Incorrect move on X")
		}

		if r.Y != s.expectedY {
			t.Error("Incorrect move on Y")
		}
	}
}

func TestCommandL(t *testing.T) {

	type scenario struct {
		x, y           int
		dir            string
		expectedX      int
		expectedY      int
		expectedFacing direction
	}

	scenarios := []scenario{
		scenario{1, 1, "n", 1, 1, DIR_W},
		scenario{1, 1, "s", 1, 1, DIR_E},
		scenario{1, 1, "e", 1, 1, DIR_N},
		scenario{1, 1, "w", 1, 1, DIR_S},
	}

	for _, s := range scenarios {
		r, _ := NewRover("test", 1, 1, s.dir)
		c := &cmdL{r}
		c.Execute()

		if r.Facing != s.expectedFacing {
			t.Error("Rover did not turn correctly")
		}

		if r.X != s.expectedX {
			t.Error("Rover should not move on X, Y for cmdL")
		}

		if r.Y != s.expectedY {
			t.Error("Rover should not move on X, Y for cmdL")
		}
	}
}

func TestCommandR(t *testing.T) {

	type scenario struct {
		x, y           int
		dir            string
		expectedX      int
		expectedY      int
		expectedFacing direction
	}

	scenarios := []scenario{
		scenario{1, 1, "n", 1, 1, DIR_E},
		scenario{1, 1, "s", 1, 1, DIR_W},
		scenario{1, 1, "e", 1, 1, DIR_S},
		scenario{1, 1, "w", 1, 1, DIR_N},
	}

	for _, s := range scenarios {
		r, _ := NewRover("test", 1, 1, s.dir)
		c := &cmdR{r}
		c.Execute()

		if r.Facing != s.expectedFacing {
			t.Error("Rover did not turn correctly")
		}

		if r.X != s.expectedX {
			t.Error("Rover should not move on X, Y for cmdR")
		}

		if r.Y != s.expectedY {
			t.Error("Rover should not move on X, Y for cmdR")
		}
	}
}

func TestDecodeCommand(t *testing.T) {

	c := decodeCommand('F', nil)
	if _, ok := c.(*cmdF); !ok {
		t.Errorf("Expected cmdF, got: %T", c)
	}

	c = decodeCommand('B', nil)
	if _, ok := c.(*cmdB); !ok {
		t.Errorf("Expected cmdB, got: %T", c)
	}

	c = decodeCommand('L', nil)
	if _, ok := c.(*cmdL); !ok {
		t.Errorf("Expected cmdL, got: %T", c)
	}

	c = decodeCommand('R', nil)
	if _, ok := c.(*cmdR); !ok {
		t.Errorf("Expected cmdR, got: %T", c)
	}

	c = decodeCommand('T', nil)
	if _, ok := c.(*noCmd); !ok {
		t.Errorf("Expected noCmd, got: %T", c)
	}
}

func TestCommandsStream(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	ch := Commands(ctx, strings.NewReader("FBFF"), nil)

	c := <-ch
	if _, ok := c.(*cmdF); !ok {
		t.Errorf("F: Got the wrong type of command: %T", c)
	}

	cancel()

	if _, ok := <-ch; ok {
		t.Errorf("Chanel should be closed already")
	}

}
