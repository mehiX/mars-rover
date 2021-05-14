package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/mehiX/mars-rover/rover"
)

var (
	name      string
	x, y      int
	direction string
	command   string
	delay     time.Duration
)

func init() {
	flag.StringVar(&name, "n", "(no name)", "name of the rover")
	flag.IntVar(&x, "x", 0, "start X coordinate")
	flag.IntVar(&y, "y", 0, "start y coordinate")
	flag.StringVar(&direction, "d", "N", "initial pointing direction")
	flag.StringVar(&command, "c", "", "commands to send to the rover")
	flag.DurationVar(&delay, "delay", 0, "use an optional delay between commands for better testing")

	flag.Parse()
}

func main() {

	rvr, err := rover.NewRover(name, x, y, direction)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start: %s\n", rvr)

	ctx, cancel := context.WithCancel(context.Background())

	// graceful shutdown
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		chInt := make(chan os.Signal, 1)
		signal.Notify(chInt, os.Interrupt)

		select {
		case <-chInt:
			// remove ^C from the output to have a clean output
			fmt.Printf("\r")
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()

	ch := rover.Commands(ctx, command, rvr)

	for c := range ch {
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
	}

	fmt.Printf("End: %s\n", rvr)

	cancel()

	// Wait for the graceful shutdown goroutine to also exit clean
	wg.Wait()
}
