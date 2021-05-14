package main

import (
	"context"
	"fmt"

	"github.com/mehiX/mars-rover/rover"
)

func main() {
	fmt.Println("Mars rover")

	kata, err := rover.NewRover("Kata", 3, 5, "N")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start: %s\n", kata)

	ctx, cancel := context.WithCancel(context.Background())

	ch := rover.Commands(ctx, "FBERLFBLLRFFFF", kata)

	for c := range ch {
		c.Execute()
	}

	fmt.Printf("End: %s\n", kata)

	cancel()
}
