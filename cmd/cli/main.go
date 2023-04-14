package main

import (
	"fmt"
	"github.com/nacho692/live-free-or-die-jugging/pkg/iterative"
	"github.com/nacho692/live-free-or-die-jugging/pkg/models"
	"log"
)

func main() {
	s, err := iterative.Solve(models.State{
		X: models.Jug{
			Capacity: 5,
			Amount:   0,
		},
		Y: models.Jug{
			Capacity: 3,
			Amount:   0,
		},
	}, 4)
	if err != nil {
		log.Fatal(err)
	}

	// TODO Fix this ugly solution representation
	for _, step := range s.Steps {
		fmt.Printf("(%d, %d) \t %s \n",
			step.State.X.Amount, step.State.X.Amount, step.Action)
	}
}
