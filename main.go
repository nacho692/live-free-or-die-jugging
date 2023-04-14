package main

import (
	"fmt"
	"github.com/nacho692/live-free-or-die-jugging/iterative"
	"log"
)

func main() {
	s, err := iterative.Solve(5, 3, 4)
	if err != nil {
		log.Fatal(err)
	}

	// TODO Fix this ugly solution representation
	for i, step := range s.Steps {
		if i == 0 {
			fmt.Printf("(%d, %d)\n", step.Y.Amount, step.X.Amount)
		} else {
			fmt.Printf(s.Action[i-1])
			fmt.Printf("(%d, %d)\n", step.Y.Amount, step.X.Amount)
		}
	}
}
