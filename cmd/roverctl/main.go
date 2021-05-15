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

	// used to syncronize with the gracefulShutdown goroutine
	var wg sync.WaitGroup
	wg.Add(1)

	// graceful shutdown
	go gracefulShutdown(ctx, &wg, cancel)

	ch := rover.Commands(ctx, command, rvr)

	rover.ExecuteAll(ctx, ch, delay)

	fmt.Printf("End: %s\n", rvr)

	cancel()

	// Wait for the graceful shutdown goroutine to also exit clean
	wg.Wait()
}

func gracefulShutdown(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc) {
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
}
