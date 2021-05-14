package rover

import (
	"fmt"
	"strings"
)

type rover struct {
	Name   string
	X, Y   int
	Facing direction
}

func (r *rover) String() string {
	return fmt.Sprintf("%s: [%s] (%d,%d)", r.Name, r.Facing, r.X, r.Y)
}

type direction int

func (d direction) String() string {
	if d == DIR_UNKNOWN {
		return "UNKNOWN"
	}

	return [...]string{"N", "E", "S", "W"}[int(d)]
}

const DIR_UNKNOWN direction = -1

const (
	DIR_N direction = iota
	DIR_E
	DIR_S
	DIR_W
)

func NewRover(name string, x, y int, d string) (*rover, error) {

	dir, err := decodeDirection(d)
	if err != nil {
		return nil, err
	}

	return &rover{name, x, y, dir}, nil
}

func decodeDirection(d string) (direction, error) {

	var dir direction

	switch strings.ToLower(d) {
	case "n":
		dir = DIR_N
	case "s":
		dir = DIR_S
	case "e":
		dir = DIR_E
	case "w":
		dir = DIR_W
	default:
		return DIR_UNKNOWN, fmt.Errorf("unknown direction: %v", d)
	}

	return dir, nil
}
