package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/mehiX/mars-rover/rover"
)

var (
	name      string
	x, y      int
	direction string
	command   string
)

func init() {
	flag.StringVar(&name, "n", "(no name)", "name of the rover")
	flag.IntVar(&x, "x", 0, "start X coordinate")
	flag.IntVar(&y, "y", 0, "start y coordinate")
	flag.StringVar(&direction, "d", "N", "initial pointing direction")
	flag.StringVar(&command, "c", "", "commands to send to the rover")

	flag.Parse()
}

func main() {

	rvr, err := rover.NewRover(name, x, y, direction)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start: %s\n", rvr)

	ctx, cancel := context.WithCancel(context.Background())

	ch := rover.Commands(ctx, command, rvr)

	for c := range ch {
		c.Execute()
	}

	fmt.Printf("End: %s\n", rvr)

	cancel()
}
